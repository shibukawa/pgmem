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
	var v69 int32
	_ = v69
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 float64
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 float64
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 float64
	_ = v371
	var v378 float64
	_ = v378
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
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
	v276 = m.ExcPending
	if v276 != 0 {
		goto L93
	} else {
		goto L94
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
		v69 = v23
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v76 = v69
	goto L7
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v33 < v32 {
		v69 = v23
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22+v39+v51)))
	v58 = v53 & (v55 ^ int32(-1))
	v60 = base.B2i32(v58 == int32(0))
	if v58 != 0 {
		v69 = v60
		goto L11
	} else {
		goto L19
	}
L18:
	;
	v69 = v60
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
	if base.B2i32(v82 == v85)|base.B2i32(v84 == v85) != 0 {
		v130 = v85
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v130 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L24:
	;
	goto L23
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v95 < v96 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = v95
	goto L28
L27:
	;
	v98 = v96
	goto L28
L28:
	;
	if v98 <= int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = int32(1)
	goto L31
L30:
	;
	v101 = v98
	goto L31
L31:
	;
	v102 = int32(8)
	v107 = int32(0)
	goto L32
L32:
	;
	v114 = v107 << (uint(int32(2)) % 32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v84+v102+v114)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v82+v102+v114)))
	v119 = v116 & v118
	v121 = base.B2i32(v119 != int32(0))
	if v119 != 0 {
		v130 = v121
		goto L24
	} else {
		goto L34
	}
L33:
	;
	v130 = v121
	goto L24
L34:
	;
	v123 = v107 + int32(1)
	if v123 != v101 {
		v107 = v123
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v136 = int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v137 == int32(0) {
		v264 = v136
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v270 == int32(0) {
		goto L1
	} else {
		goto L92
	}
L38:
	;
	v270 = v264
	goto L37
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)+228))
	v142 = F_bms_overlap(m, v140, v141)
	mBase = m.M
	if v142 == int32(0) {
		v264 = v136
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v145 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v146 - int32(279) {
	case 0, 1:
		goto L41
	default:
		v264 = v145
		goto L38
	case 3:
		goto L51
	case 4:
		goto L50
	case 5:
		goto L49
	case 9:
		goto L48
	case 10:
		goto L47
	case 11:
		goto L45
	case 14:
		goto L44
	case 15:
		goto L43
	case 17:
		goto L42
	case 19, 20, 21:
		goto L46
	}
L41:
	;
	v264 = int32(1)
	goto L38
L42:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v254 = F_path_is_reparameterizable_by_child(m, v253, v133)
	mBase = m.M
	if v254 == int32(0) {
		v264 = v145
		goto L38
	} else {
		goto L91
	}
L43:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v252 = F_path_is_reparameterizable_by_child(m, v251, v133)
	mBase = m.M
	if v252 != 0 {
		goto L41
	} else {
		goto L90
	}
L44:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v250 = F_path_is_reparameterizable_by_child(m, v249, v133)
	mBase = m.M
	if v250 != 0 {
		goto L41
	} else {
		goto L89
	}
L45:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v227 == int32(0) {
		goto L41
	} else {
		goto L81
	}
L46:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v222 = F_path_is_reparameterizable_by_child(m, v221, v133)
	mBase = m.M
	if v222 == int32(0) {
		v264 = v145
		goto L38
	} else {
		goto L79
	}
L47:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v199 == int32(0) {
		goto L41
	} else {
		goto L71
	}
L48:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v195 == int32(0) {
		goto L41
	} else {
		goto L69
	}
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v173 == int32(0) {
		goto L41
	} else {
		goto L61
	}
L50:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v151 == int32(0) {
		goto L41
	} else {
		goto L53
	}
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v150 = F_path_is_reparameterizable_by_child(m, v149, v133)
	mBase = m.M
	if v150 != 0 {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v264 = v145
	goto L38
L53:
	;
	v154 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v155 <= v154 {
		goto L41
	} else {
		goto L54
	}
L54:
	;
	v158 = v154
	goto L55
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v158<<(uint(int32(2))%32))))
	v167 = F_path_is_reparameterizable_by_child(m, v166, v133)
	mBase = m.M
	if v167 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v270 = int32(0)
	goto L37
L57:
	;
	v169 = v158 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v169 < v170 {
		v158 = v169
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L41
L61:
	;
	v176 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v177 <= v176 {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	v180 = v176
	goto L63
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v180<<(uint(int32(2))%32))))
	v189 = F_path_is_reparameterizable_by_child(m, v188, v133)
	mBase = m.M
	if v189 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v270 = int32(0)
	goto L37
L65:
	;
	v191 = v180 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v191 < v192 {
		v180 = v191
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L41
L69:
	;
	v198 = F_path_is_reparameterizable_by_child(m, v195, v133)
	mBase = m.M
	if v198 != 0 {
		goto L41
	} else {
		goto L70
	}
L70:
	;
	v264 = v145
	goto L38
L71:
	;
	v202 = int32(0)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v203 <= v202 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	v206 = v202
	goto L73
L73:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210+v206<<(uint(int32(2))%32))))
	v215 = F_path_is_reparameterizable_by_child(m, v214, v133)
	mBase = m.M
	if v215 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v270 = int32(0)
	goto L37
L75:
	;
	v217 = v206 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v217 < v218 {
		v206 = v217
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L41
L79:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v226 = F_path_is_reparameterizable_by_child(m, v225, v133)
	mBase = m.M
	if v226 != 0 {
		goto L41
	} else {
		goto L80
	}
L80:
	;
	v264 = v145
	goto L38
L81:
	;
	v230 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v231 <= v230 {
		goto L41
	} else {
		goto L82
	}
L82:
	;
	v234 = v230
	goto L83
L83:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238+v234<<(uint(int32(2))%32))))
	v243 = F_path_is_reparameterizable_by_child(m, v242, v133)
	mBase = m.M
	if v243 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v270 = int32(0)
	goto L37
L85:
	;
	v245 = v234 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v245 < v246 {
		v234 = v245
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	goto L41
L89:
	;
	v264 = v145
	goto L38
L90:
	;
	v264 = v145
	goto L38
L91:
	;
	goto L41
L92:
	;
	goto L2
L93:
	;
	return
L94:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v279 = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v285 == v279 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	if v425 == int32(0) {
		goto L1
	} else {
		goto L136
	}
L96:
	;
	v425 = v410
	goto L95
L97:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v342 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L98:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	if v288 <= int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v299 = v279
	goto L100
L100:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v299<<(uint(int32(2))%32))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+64))
	v309 = F_compare_pathkeys(m, l4, v308)
	mBase = m.M
	if v309 == int32(3) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L97
L102:
	;
	v329 = v299 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	if v329 < v330 {
		v299 = v329
		goto L100
	} else {
		goto L107
	}
L103:
	;
	v312 = *(*float64)(unsafe.Add(mBase, uint32(v307)+56))
	v316 = int32(0)
	v320 = base.B2i32(base.F64_gt(v278, base.F64_mul(v312, float64(1.01))) == v316) | base.B2i32(v309 == int32(1))
	if v320 == v316 {
		v410 = v320
		goto L96
	} else {
		goto L104
	}
L104:
	;
	if v309 == int32(2) {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	if base.F64_lt(base.F64_mul(v278, float64(1.01)), v312) != 0 {
		v410 = v320
		goto L96
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	goto L101
L108:
	;
	v425 = int32(1)
	goto L95
L109:
	;
	goto L110
L110:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	if v346 <= int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v425 = int32(1)
	goto L95
L112:
	;
	goto L113
L113:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v358 = int32(0)
	goto L114
L114:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v358<<(uint(int32(2))%32))))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+40))
	if v277 != v367 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v410 = v400
	goto L96
L116:
	;
	if v351 != 0 {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	if v367 <= v277 {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v371 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	if base.F64_le(v278, base.F64_mul(v371, float64(1.01))) == int32(0) {
		goto L116
	} else {
		goto L121
	}
L120:
	;
	v425 = int32(1)
	goto L95
L121:
	;
	v425 = int32(1)
	goto L95
L122:
	;
	v400 = int32(1)
	v402 = v358 + v400
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	if v402 < v403 {
		v358 = v402
		goto L114
	} else {
		goto L135
	}
L123:
	;
	v378 = *(*float64)(unsafe.Add(mBase, uint32(v366)+48))
	if base.F64_gt(v278, base.F64_mul(v378, float64(1.01))) == int32(0) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	if v385 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	v388 = int32(0)
	goto L129
L128:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v366)+64))
	v388 = v387
	goto L129
L129:
	;
	v389 = F_compare_pathkeys(m, l4, v388)
	mBase = m.M
	if v389&int32(-3) != 0 {
		goto L122
	} else {
		goto L130
	}
L130:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	if v393 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v396 = v394
	goto L133
L132:
	;
	v396 = int32(0)
	goto L133
L133:
	;
	v397 = F_bms_equal(m, int32(0), v396)
	mBase = m.M
	if v397 != 0 {
		v410 = int32(0)
		goto L96
	} else {
		goto L134
	}
L134:
	;
	goto L122
L135:
	;
	goto L115
L136:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v430 = F_create_nestloop_path(m, l0, l1, l5, v13, l6, l2, l3, v428, l4, int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L93
	} else {
		goto L137
	}
L137:
	;
	F_add_partial_path(m, l1, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L93
	} else {
		goto L138
	}
L138:
	;
	goto L1
}
