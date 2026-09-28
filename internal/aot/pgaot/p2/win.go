package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinGetFuncArgInFrame(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v213 int64
	_ = v213
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v234 int64
	_ = v234
	var v242 int64
	_ = v242
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v342 int32
	_ = v342
	var v351 int64
	_ = v351
	v6 = int64(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+408))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v21 == int32(3) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 - int32(-64)
	return v351
L2:
	;
	v342 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v342)
	v351 = int64(0)
	goto L1
L3:
	;
	v311 = F_window_gettupleslot(m, l0, v307, v19)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L12
	} else {
		goto L117
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L114
	}
L5:
	;
	switch l2 {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	default:
		goto L4
	}
L6:
	;
	goto L7
L7:
	;
	switch l2 {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	default:
		goto L42
	}
L8:
	;
	v57 = l1 >> (uint(int32(31)) % 32)
	v59 = l1 ^ v57 - v57
	v62 = int32(0)
	v66 = v53
	v67 = v6
	goto L20
L9:
	;
	if int32(0) < l1 {
		goto L2
	} else {
		goto L18
	}
L10:
	;
	if l1 < int32(0) {
		goto L2
	} else {
		goto L16
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int64(0)
L13:
	;
	F_errmsg_internal(m, int32(_a_F_WinGetFuncArgInFrame_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_WinGetFuncArgInFrame_1), int32(3433), int32(_a_F_WinGetFuncArgInFrame_2))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	F_update_frameheadpos(m, v18)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	v53 = v43
	v54 = v43
	v55 = int64(1)
	goto L8
L18:
	;
	F_update_frametailpos(m, v18)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
	v53 = v49 - int64(1)
	v54 = v6
	v55 = int64(-1)
	goto L8
L20:
	;
	if v66 < int64(0) {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	if l3 == int32(0) {
		v351 = v115
		goto L1
	} else {
		goto L38
	}
L22:
	;
	v77 = F_row_is_in_frame(m, l0, v66, v19, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	if v114 <= v59 {
		v62 = v114
		v66 = v66 + v55
		v67 = v115
		goto L20
	} else {
		goto L37
	}
L24:
	;
	v81 = F_get_notnull_info(m, l0, v66)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L28
	}
L25:
	;
	switch v77 + int32(1) {
	case 0:
		goto L2
	case 1:
		v114 = v62
		v115 = v67
		goto L23
	default:
		goto L24
	}
L26:
	;
	v101 = v62 + int32(1)
	if v101 <= v59 {
		v114 = v101
		v115 = v67
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v83 = F_window_gettupleslot(m, l0, v66, v19)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	switch v81 {
	case 0:
		goto L27
	case 1:
		v114 = v62
		v115 = v67
		goto L23
	default:
		goto L26
	}
L29:
	;
	if v83 == int32(0) {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
	v92 = m.T0[v91].(func(*base.Module, int32, int32, int32) int64)(m, v90, v20, l4)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	F_put_notnull_info(m, l0, v66, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	v114 = v62 + (v94 ^ int32(1))
	v115 = v92
	goto L23
L33:
	;
	v103 = F_window_gettupleslot(m, l0, v66, v19)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v103 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v112 = m.T0[v111].(func(*base.Module, int32, int32, int32) int64)(m, v110, v20, l4)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v114 = v101
	v115 = v112
	goto L23
L37:
	;
	goto L21
L38:
	;
	F_WinSetMarkPosition(m, l0, v54)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	v351 = v115
	goto L1
L40:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
	F_update_frameheadpos(m, v18)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L12
	} else {
		goto L112
	}
L41:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
	v307 = v142 + base.I64_extend_i32_u(base.B2i32(v278 <= v142)&base.B2i32(v140 <= v278))
	v308 = v142
	goto L3
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L109
	}
L43:
	;
	if int32(0) < l1 {
		goto L2
	} else {
		goto L80
	}
L44:
	;
	if l1 < int32(0) {
		goto L2
	} else {
		goto L49
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_WinGetFuncArgInFrame_0), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_WinGetFuncArgInFrame_1), int32(4048), int32(_a_F_WinGetFuncArgInFrame_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_update_frameheadpos(m, v18)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	v142 = v140 + base.I64_extend_i32_u(l1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
	switch int32(base.Ui32(v143)>>(uint(int32(15))%32)) & int32(7) {
	case 0:
		v307 = v142
		v308 = v142
		goto L3
	case 1:
		goto L41
	case 2:
		goto L53
	default:
		goto L51
	case 4:
		goto L52
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L77
	}
L52:
	;
	F_update_grouptailpos(m, v18)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L64
	}
L53:
	;
	F_update_grouptailpos(m, v18)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
	if v142 < v150 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v307 = v142
	v308 = v142
	goto L3
L56:
	;
	goto L57
L57:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v18)+360))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	if v152 <= v153 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v307 = v142
	v308 = v142
	goto L3
L59:
	;
	goto L60
L60:
	;
	if v153 < v150 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v157 = v150
	goto L63
L62:
	;
	v157 = v153
	goto L63
L63:
	;
	v307 = v142 + v152 - v157
	v308 = v142
	goto L3
L64:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
	if v142 < v161 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v307 = v142
	v308 = v142
	goto L3
L66:
	;
	goto L67
L67:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v18)+360))
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	if v163 <= v164 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v307 = v142
	v308 = v142
	goto L3
L69:
	;
	goto L70
L70:
	;
	if v164 < v161 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v167 = v161
	goto L73
L72:
	;
	v167 = v164
	goto L73
L73:
	;
	if v167 == v142 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
	v307 = v169
	v308 = v142
	goto L3
L75:
	;
	goto L76
L76:
	;
	v307 = v142 + v163 + (v167 ^ int64(-1))
	v308 = v142
	goto L3
L77:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v178
	F_errmsg_internal(m, int32(_a_F_WinGetFuncArgInFrame_4), v14+int32(-32))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_WinGetFuncArgInFrame_1), int32(_a_F_WinGetFuncArgInFrame_5), int32(_a_F_WinGetFuncArgInFrame_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
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
	F_update_frametailpos(m, v18)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
	v198 = v194 + base.I64_extend_i32_s(l1) - int64(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
	switch int32(base.Ui32(v199)>>(uint(int32(15))%32)) & int32(7) {
	case 0:
		v307 = v198
		v308 = v198
		goto L3
	case 1:
		goto L40
	case 2:
		goto L84
	default:
		goto L82
	case 4:
		goto L83
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
	} else {
		goto L106
	}
L83:
	;
	F_update_grouptailpos(m, v18)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L94
	}
L84:
	;
	F_update_grouptailpos(m, v18)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v18)+360))
	if v206 <= v198 {
		v217 = v198
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_update_frameheadpos(m, v18)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L92
	}
L87:
	;
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
	if v209 <= v208 {
		v217 = v198
		goto L86
	} else {
		goto L88
	}
L88:
	;
	if v206 < v209 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v213 = v206
	goto L91
L90:
	;
	v213 = v209
	goto L91
L91:
	;
	v217 = v198 + v208 - v213
	goto L86
L92:
	;
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	if v220 <= v217 {
		v307 = v217
		v308 = v220
		goto L3
	} else {
		goto L93
	}
L93:
	;
	goto L2
L94:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v18)+360))
	if v224 <= v198 {
		v242 = v198
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_update_frameheadpos(m, v18)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L104
	}
L96:
	;
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
	if v227 <= v226 {
		v242 = v198
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v224 < v227 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v230 = v224
	goto L100
L99:
	;
	v230 = v227
	goto L100
L100:
	;
	if v230-int64(1) == v198 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
	v242 = v234
	goto L95
L102:
	;
	goto L103
L103:
	;
	v242 = v198 + v226 - v230 + int64(1)
	goto L95
L104:
	;
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	if v245 <= v242 {
		v307 = v242
		v308 = v245
		goto L3
	} else {
		goto L105
	}
L105:
	;
	goto L2
L106:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v251
	F_errmsg_internal(m, int32(_a_F_WinGetFuncArgInFrame_4), v14+int32(-16))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_WinGetFuncArgInFrame_1), int32(_a_F_WinGetFuncArgInFrame_6), int32(_a_F_WinGetFuncArgInFrame_3))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg_internal(m, int32(_a_F_WinGetFuncArgInFrame_7), v14+int32(-48))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_WinGetFuncArgInFrame_1), int32(_a_F_WinGetFuncArgInFrame_8), int32(_a_F_WinGetFuncArgInFrame_3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L12
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
	v291 = v198 - base.I64_extend_i32_u(base.B2i32(v284 < v194)&base.B2i32(v198 <= v284))
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
	if v292 <= v291 {
		v307 = v291
		v308 = v292
		goto L3
	} else {
		goto L113
	}
L113:
	;
	goto L2
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	F_errmsg_internal(m, int32(_a_F_WinGetFuncArgInFrame_7), v16)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_WinGetFuncArgInFrame_1), int32(3455), int32(_a_F_WinGetFuncArgInFrame_2))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	if v311 == int32(0) {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v316 = F_row_is_in_frame(m, l0, v307, v19, int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	if v316 <= int32(0) {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	if l3 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_WinSetMarkPosition(m, l0, v308)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L12
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+24))
	v327 = m.T0[v326].(func(*base.Module, int32, int32, int32) int64)(m, v325, v20, l4)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L12
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v351 = v327
	goto L1
}
func F_WinSetMarkPosition(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 <= l1 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		F_tuplestore_select_read_pointer(m, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			if v12 < l1 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
				v17 = F_tuplestore_skiptuples(m, v14, l1-v12, int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					F_tuplestore_select_read_pointer(m, v20, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						if v24 < l1 {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
							v29 = F_tuplestore_skiptuples(m, v26, l1-v24, int32(1))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
								return
							}
						} else {
							return
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_tuplestore_select_read_pointer(m, v20, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					if v24 < l1 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
						v29 = F_tuplestore_skiptuples(m, v26, l1-v24, int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
							return
						}
					} else {
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_WinSetMarkPosition_0), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_WinSetMarkPosition_1), int32(3769), int32(_a_F_WinSetMarkPosition_2))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
