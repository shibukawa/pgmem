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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
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
		goto L106
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
		v42 = int32(_a_F_transformWindowFuncCall_0)
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
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v94 != 0 {
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
	v42 = int32(_a_F_transformWindowFuncCall_1)
	goto L10
L12:
	;
	v42 = int32(_a_F_transformWindowFuncCall_2)
	goto L10
L13:
	;
	v42 = int32(_a_F_transformWindowFuncCall_3)
	goto L10
L14:
	;
	v42 = int32(_a_F_transformWindowFuncCall_4)
	goto L10
L15:
	;
	v42 = int32(_a_F_transformWindowFuncCall_5)
	goto L10
L16:
	;
	v42 = int32(_a_F_transformWindowFuncCall_6)
	goto L10
L17:
	;
	v42 = int32(_a_F_transformWindowFuncCall_7)
	goto L10
L18:
	;
	v42 = int32(_a_F_transformWindowFuncCall_8)
	goto L10
L19:
	;
	v42 = int32(_a_F_transformWindowFuncCall_9)
	goto L10
L20:
	;
	v42 = int32(_a_F_transformWindowFuncCall_10)
	goto L10
L21:
	;
	v42 = int32(_a_F_transformWindowFuncCall_11)
	goto L10
L22:
	;
	v42 = int32(_a_F_transformWindowFuncCall_12)
	goto L10
L23:
	;
	v42 = int32(_a_F_transformWindowFuncCall_13)
	goto L10
L24:
	;
	v42 = int32(_a_F_transformWindowFuncCall_14)
	goto L10
L25:
	;
	v42 = int32(_a_F_transformWindowFuncCall_15)
	goto L10
L26:
	;
	v42 = int32(_a_F_transformWindowFuncCall_16)
	goto L10
L27:
	;
	v42 = int32(_a_F_transformWindowFuncCall_17)
	goto L10
L28:
	;
	F_errcode(m, int32(_a_F_transformWindowFuncCall_18))
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
	F_errmsg_internal(m, int32(_a_F_transformWindowFuncCall_19), v13+int32(16))
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
	F_errfinish(m, int32(_a_F_transformWindowFuncCall_20), int32(1039), int32(_a_F_transformWindowFuncCall_21))
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
	F_errcode(m, int32(_a_F_transformWindowFuncCall_18))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v78
	F_errmsg(m, int32(_a_F_transformWindowFuncCall_22), v13+int32(32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L39
	}
L36:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_c_F_transformWindowFuncCall[0])))
	v78 = v76
	goto L38
L37:
	;
	v78 = int32(_a_F_transformWindowFuncCall_23)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_parser_errposition(m, l0, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_transformWindowFuncCall_20), int32(1046), int32(_a_F_transformWindowFuncCall_21))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	if v93 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v93 != 0 {
		goto L71
	} else {
		goto L72
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L66
	}
L47:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v97 <= int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v100 = int32(0)
	if v100 < v97 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v103 = v97
	goto L51
L50:
	;
	v103 = v100
	goto L51
L51:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v111 = v4
	goto L52
L52:
	;
	v116 = v111 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v104+v111<<(uint(int32(2))%32))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v121 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L46
L54:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if base.B2i32(v124 == int32(0))|base.B2i32(v124 != v127) != 0 {
		v145 = v124
		v146 = v127
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	if v103 != v116 {
		v111 = v116
		goto L52
	} else {
		goto L65
	}
L57:
	;
	if v145-v146 == int32(0) {
		v295 = v116
		goto L42
	} else {
		goto L64
	}
L58:
	;
	goto L57
L59:
	;
	v130 = v121
	v131 = v94
	goto L60
L60:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v135 == int32(0) {
		v145 = v135
		v146 = v134
		goto L58
	} else {
		goto L62
	}
L61:
	;
	v145 = v135
	v146 = v134
	goto L58
L62:
	;
	v138 = int32(1)
	if v135 == v134 {
		v130 = v130 + v138
		v131 = v131 + v138
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L56
L65:
	;
	goto L53
L66:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v168
	F_errmsg(m, int32(_a_F_transformWindowFuncCall_24), v13)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_parser_errposition(m, l0, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_transformWindowFuncCall_20), int32(1079), int32(_a_F_transformWindowFuncCall_21))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if int32(0) < v182 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v285 = int32(0)
	goto L73
L73:
	;
	v286 = F_lappend(m, v285, l2)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L104
	}
L74:
	;
	v191 = v4
	goto L77
L75:
	;
	goto L76
L76:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v285 = v273
	goto L73
L77:
	;
	v196 = v191 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198+v191<<(uint(int32(2))%32))))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	if v203 != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	goto L76
L79:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v196 < v261 {
		v191 = v196
		goto L77
	} else {
		goto L103
	}
L80:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v236 = F_equal(m, v234, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L94
	}
L81:
	;
	if v197 == int32(0) {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v197 != 0 {
		goto L79
	} else {
		goto L93
	}
L84:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if base.B2i32(v208 == int32(0))|base.B2i32(v208 != v211) != 0 {
		v229 = v208
		v230 = v211
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v229-v230 == int32(0) {
		goto L80
	} else {
		goto L92
	}
L86:
	;
	goto L85
L87:
	;
	v214 = v203
	v215 = v197
	goto L88
L88:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	if v219 == int32(0) {
		v229 = v219
		v230 = v218
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v229 = v219
	v230 = v218
	goto L86
L90:
	;
	v222 = int32(1)
	if v219 == v218 {
		v214 = v214 + v222
		v215 = v215 + v222
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	goto L79
L93:
	;
	goto L80
L94:
	;
	if v236 == int32(0) {
		goto L79
	} else {
		goto L95
	}
L95:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v202)+16))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v242 = F_equal(m, v240, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	if v242 == int32(0) {
		goto L79
	} else {
		goto L97
	}
L97:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v202)+20))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v246 != v247 {
		goto L79
	} else {
		goto L98
	}
L98:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v202)+24))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v251 = F_equal(m, v249, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	if v251 == int32(0) {
		goto L79
	} else {
		goto L100
	}
L100:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v257 = F_equal(m, v255, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v257 == int32(0) {
		goto L79
	} else {
		goto L102
	}
L102:
	;
	v295 = v196
	goto L42
L103:
	;
	goto L78
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v286
	if v286 == int32(0) {
		v295 = int32(0)
		goto L42
	} else {
		goto L105
	}
L105:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v295 = v291
	goto L42
L106:
	;
	F_errcode(m, int32(_a_F_transformWindowFuncCall_18))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_transformWindowFuncCall_25), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v320 = F_locate_windowfunc(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_parser_errposition(m, l0, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_transformWindowFuncCall_20), int32(898), int32(_a_F_transformWindowFuncCall_21))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
