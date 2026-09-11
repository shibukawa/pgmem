package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_try_partial_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 float64
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 float64
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 float64
	_ = v370
	var v377 float64
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v15 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(96)
	return
L2:
	;
	F_initial_cost_nestloop(m, l0, v13, l5, l2, l3, l6)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L94
	} else {
		goto L95
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+228))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v20
	goto L6
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v22 = v21
	goto L6
L6:
	;
	v23 = int32(0)
	if v18 == v23 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v76 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L8:
	;
	v76 = int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	if v22 == int32(0) {
		v67 = v23
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v76 = v67
	goto L7
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v33 < v32 {
		v67 = v23
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v35 = int32(1)
	if v32 <= v35 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = v35
	goto L16
L15:
	;
	v38 = v32
	goto L16
L16:
	;
	v39 = int32(8)
	v44 = int32(0)
	goto L17
L17:
	;
	v51 = v44 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18+v39+v51)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v22+v39))))
	v58 = v53 & (v55 ^ int32(-1))
	v60 = base.B2i32(v58 == int32(0))
	if v58 != 0 {
		v67 = v60
		goto L11
	} else {
		goto L19
	}
L18:
	;
	v67 = v60
	goto L11
L19:
	;
	v62 = v44 + int32(1)
	if v62 != v38 {
		v44 = v62
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v79 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+228))
	v85 = int32(0)
	if v82 == v85 {
		v126 = v85
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v126 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L24:
	;
	goto L23
L25:
	;
	if v84 == int32(0) {
		v126 = v85
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v94 < v95 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v97 = v94
	goto L29
L28:
	;
	v97 = v95
	goto L29
L29:
	;
	if v97 <= int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v100 = int32(1)
	goto L32
L31:
	;
	v100 = v97
	goto L32
L32:
	;
	v101 = int32(8)
	v106 = int32(0)
	goto L33
L33:
	;
	v113 = v106 << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84+v101+v113)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+(v82+v101))))
	v118 = v115 & v117
	v120 = base.B2i32(v118 != int32(0))
	if v118 != 0 {
		v126 = v120
		goto L24
	} else {
		goto L35
	}
L34:
	;
	v126 = v120
	goto L24
L35:
	;
	v122 = v106 + int32(1)
	if v122 != v100 {
		v106 = v122
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v135 = int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v136 == int32(0) {
		v263 = v135
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v269 == int32(0) {
		goto L1
	} else {
		goto L93
	}
L39:
	;
	v269 = v263
	goto L38
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)+228))
	v141 = F_bms_overlap(m, v139, v140)
	mBase = m.M
	if v141 == int32(0) {
		v263 = v135
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v144 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v145 - int32(279) {
	case 0, 1:
		goto L42
	default:
		v263 = v144
		goto L39
	case 3:
		goto L52
	case 4:
		goto L51
	case 5:
		goto L50
	case 9:
		goto L49
	case 10:
		goto L48
	case 11:
		goto L46
	case 14:
		goto L45
	case 15:
		goto L44
	case 17:
		goto L43
	case 19, 20, 21:
		goto L47
	}
L42:
	;
	v263 = int32(1)
	goto L39
L43:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v253 = F_path_is_reparameterizable_by_child(m, v252, v132)
	mBase = m.M
	if v253 == int32(0) {
		v263 = v144
		goto L39
	} else {
		goto L92
	}
L44:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v251 = F_path_is_reparameterizable_by_child(m, v250, v132)
	mBase = m.M
	if v251 != 0 {
		goto L42
	} else {
		goto L91
	}
L45:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v249 = F_path_is_reparameterizable_by_child(m, v248, v132)
	mBase = m.M
	if v249 != 0 {
		goto L42
	} else {
		goto L90
	}
L46:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v226 == int32(0) {
		goto L42
	} else {
		goto L82
	}
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v221 = F_path_is_reparameterizable_by_child(m, v220, v132)
	mBase = m.M
	if v221 == int32(0) {
		v263 = v144
		goto L39
	} else {
		goto L80
	}
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v198 == int32(0) {
		goto L42
	} else {
		goto L72
	}
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v194 == int32(0) {
		goto L42
	} else {
		goto L70
	}
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v172 == int32(0) {
		goto L42
	} else {
		goto L62
	}
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v150 == int32(0) {
		goto L42
	} else {
		goto L54
	}
L52:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v149 = F_path_is_reparameterizable_by_child(m, v148, v132)
	mBase = m.M
	if v149 != 0 {
		goto L42
	} else {
		goto L53
	}
L53:
	;
	v263 = v144
	goto L39
L54:
	;
	v153 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v154 <= v153 {
		goto L42
	} else {
		goto L55
	}
L55:
	;
	v157 = v153
	goto L56
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v157<<(uint(int32(2))%32))))
	v166 = F_path_is_reparameterizable_by_child(m, v165, v132)
	mBase = m.M
	if v166 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v269 = int32(0)
	goto L38
L58:
	;
	v168 = v157 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v168 < v169 {
		v157 = v168
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L42
L62:
	;
	v175 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v176 <= v175 {
		goto L42
	} else {
		goto L63
	}
L63:
	;
	v179 = v175
	goto L64
L64:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v179<<(uint(int32(2))%32))))
	v188 = F_path_is_reparameterizable_by_child(m, v187, v132)
	mBase = m.M
	if v188 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v269 = int32(0)
	goto L38
L66:
	;
	v190 = v179 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v190 < v191 {
		v179 = v190
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	goto L42
L70:
	;
	v197 = F_path_is_reparameterizable_by_child(m, v194, v132)
	mBase = m.M
	if v197 != 0 {
		goto L42
	} else {
		goto L71
	}
L71:
	;
	v263 = v144
	goto L39
L72:
	;
	v201 = int32(0)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v202 <= v201 {
		goto L42
	} else {
		goto L73
	}
L73:
	;
	v205 = v201
	goto L74
L74:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v205<<(uint(int32(2))%32))))
	v214 = F_path_is_reparameterizable_by_child(m, v213, v132)
	mBase = m.M
	if v214 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v269 = int32(0)
	goto L38
L76:
	;
	v216 = v205 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v216 < v217 {
		v205 = v216
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L42
L80:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v225 = F_path_is_reparameterizable_by_child(m, v224, v132)
	mBase = m.M
	if v225 != 0 {
		goto L42
	} else {
		goto L81
	}
L81:
	;
	v263 = v144
	goto L39
L82:
	;
	v229 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v230 <= v229 {
		goto L42
	} else {
		goto L83
	}
L83:
	;
	v233 = v229
	goto L84
L84:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v233<<(uint(int32(2))%32))))
	v242 = F_path_is_reparameterizable_by_child(m, v241, v132)
	mBase = m.M
	if v242 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v269 = int32(0)
	goto L38
L86:
	;
	v244 = v233 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v244 < v245 {
		v233 = v244
		goto L84
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	goto L42
L90:
	;
	v263 = v144
	goto L39
L91:
	;
	v263 = v144
	goto L39
L92:
	;
	goto L42
L93:
	;
	goto L2
L94:
	;
	return
L95:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v278 = int32(0)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v284 == v278 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	if v422 == int32(0) {
		goto L1
	} else {
		goto L137
	}
L97:
	;
	v422 = v407
	goto L96
L98:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v341 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L99:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v287 <= int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v296 = v278
	goto L101
L101:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302+v296<<(uint(int32(2))%32))))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+64))
	v308 = F_compare_pathkeys(m, l4, v307)
	mBase = m.M
	if v308 == int32(3) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L98
L103:
	;
	v328 = v296 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v328 < v329 {
		v296 = v328
		goto L101
	} else {
		goto L108
	}
L104:
	;
	v311 = *(*float64)(unsafe.Add(mBase, uint32(v306)+56))
	v315 = int32(0)
	v319 = base.B2i32(base.F64_gt(v277, base.F64_mul(v311, float64(1.01))) == v315) | base.B2i32(v308 == int32(1))
	if v319 == v315 {
		v407 = v319
		goto L97
	} else {
		goto L105
	}
L105:
	;
	if v308 == int32(2) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	if base.F64_lt(base.F64_mul(v277, float64(1.01)), v311) != 0 {
		v407 = v319
		goto L97
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	goto L102
L109:
	;
	v422 = int32(1)
	goto L96
L110:
	;
	goto L111
L111:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v345 <= int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v422 = int32(1)
	goto L96
L113:
	;
	goto L114
L114:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v357 = int32(0)
	goto L115
L115:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361+v357<<(uint(int32(2))%32))))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+40))
	if v276 != v366 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v407 = v397
	goto L97
L117:
	;
	if v350 != 0 {
		goto L124
	} else {
		goto L125
	}
L118:
	;
	if v366 <= v276 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v370 = *(*float64)(unsafe.Add(mBase, uint32(v365)+56))
	if base.F64_le(v277, base.F64_mul(v370, float64(1.01))) == int32(0) {
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v422 = int32(1)
	goto L96
L122:
	;
	v422 = int32(1)
	goto L96
L123:
	;
	v397 = int32(1)
	v399 = v357 + v397
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v399 < v400 {
		v357 = v399
		goto L115
	} else {
		goto L136
	}
L124:
	;
	v377 = *(*float64)(unsafe.Add(mBase, uint32(v365)+48))
	if base.F64_gt(v277, base.F64_mul(v377, float64(1.01))) == int32(0) {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v383 = int32(0)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
	if v384 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L126
L128:
	;
	v386 = v383
	goto L130
L129:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v365)+64))
	v386 = v385
	goto L130
L130:
	;
	v387 = F_compare_pathkeys(m, l4, v386)
	mBase = m.M
	if v387&int32(-3) != 0 {
		goto L123
	} else {
		goto L131
	}
L131:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
	if v391 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	v393 = v392
	goto L134
L133:
	;
	v393 = v383
	goto L134
L134:
	;
	v394 = F_bms_equal(m, int32(0), v393)
	mBase = m.M
	if v394 != 0 {
		v407 = v383
		goto L97
	} else {
		goto L135
	}
L135:
	;
	goto L123
L136:
	;
	goto L116
L137:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v427 = F_create_nestloop_path(m, l0, l1, l5, v13, l6, l2, l3, v425, l4, int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L94
	} else {
		goto L138
	}
L138:
	;
	F_add_partial_path(m, l1, v427)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L94
	} else {
		goto L139
	}
L139:
	;
	goto L1
}
