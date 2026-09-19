package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v50
L2:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v39 != v40 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v50 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 != 0 {
		v50 = v3
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	return int32(0)
L17:
	;
	goto L18
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v46 = F_equal(m, v44, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v50 = v46
	goto L1
}
func F_get_extension_control_directories(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	v9 = m.G0
	v11 = v9 - int32(1056)
	m.G0 = v11
	v14 = v11 + int32(32)
	F_get_share_path(m, v14)
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
	v23 = F_psprintf(m, int32(_a_F_get_extension_control_directories_0), v11+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_get_extension_control_directories[0]))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(1056)
	return v220
L5:
	;
	v31 = F_lappend(m, int32(0), v23)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v33 = F_pstrdup(m, v26)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v220 = v31
	goto L4
L9:
	;
	v35 = v33
	v38 = int32(0)
	goto L10
L10:
	;
	v44 = Fn13900(m, v35, int32(58))
	mBase = m.M
	goto L13
L11:
	;
	v220 = v213
	goto L4
L12:
	;
	v51 = v49 + int32(1)
	v52 = F_palloc(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = F_strlen(m, v35)
	mBase = m.M
	v49 = v47
	goto L12
L15:
	;
	goto L16
L16:
	;
	v49 = v44 - v35
	goto L12
L17:
	;
	if v51 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v172 = int32(_a_F_get_extension_control_directories_1)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_extension_control_directories[1])))
	if base.B2i32(v175 == int32(0))|base.B2i32(v175 != v178) != 0 {
		v196 = v175
		v197 = v178
		goto L51
	} else {
		goto L52
	}
L19:
	;
	v169 = F_strlen(m, v165)
	mBase = m.M
	goto L18
L20:
	;
	v165 = v35
	goto L19
L21:
	;
	goto L22
L22:
	;
	v59 = v51 - int32(1)
	if (v52^v35)&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v162)
	v165 = v158
	goto L19
L24:
	;
	v143 = v138
	v144 = v139
	v145 = v140
	goto L45
L25:
	;
	if v133 == int32(0) {
		v158 = v131
		v159 = v132
		goto L23
	} else {
		goto L44
	}
L26:
	;
	v131 = v35
	v132 = v52
	v133 = v59
	goto L25
L27:
	;
	goto L28
L28:
	;
	v63 = int32(0)
	if base.B2i32(v35&int32(3) == v63)|base.B2i32(v59 == v63) == v63 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v99 == int32(0) {
		v158 = v96
		v159 = v97
		goto L23
	} else {
		goto L38
	}
L30:
	;
	v75 = v35
	v76 = v52
	v77 = v59
	goto L33
L31:
	;
	goto L32
L32:
	;
	v96 = v35
	v97 = v52
	v98 = v59
	v99 = base.B2i32(v59 != v63)
	goto L29
L33:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v79)
	if v79 == int32(0) {
		v138 = v75
		v139 = v76
		v140 = v77
		goto L24
	} else {
		goto L35
	}
L34:
	;
	v96 = v90
	v97 = v84
	v98 = v86
	v99 = v88
	goto L29
L35:
	;
	v83 = int32(1)
	v84 = v76 + v83
	v86 = v77 - v83
	v87 = int32(0)
	v88 = base.B2i32(v86 != v87)
	v90 = v75 + v83
	if v90&int32(3) == v87 {
		v96 = v90
		v97 = v84
		v98 = v86
		v99 = v88
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v86 != 0 {
		v75 = v90
		v76 = v84
		v77 = v86
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if base.B2i32(v102 == int32(0))|base.B2i32(base.Ui32(v98) < base.Ui32(int32(4))) != 0 {
		v131 = v96
		v132 = v97
		v133 = v98
		goto L25
	} else {
		goto L39
	}
L39:
	;
	v109 = v96
	v110 = v97
	v111 = v98
	goto L40
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v117 = int32(-2139062144)
	if (int32(16843008)-v114|v114)&v117 != v117 {
		v138 = v109
		v139 = v110
		v140 = v111
		goto L24
	} else {
		goto L42
	}
L41:
	;
	v131 = v125
	v132 = v123
	v133 = v127
	goto L25
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v114
	v122 = int32(4)
	v123 = v110 + v122
	v125 = v109 + v122
	v127 = v111 - v122
	if base.Ui32(int32(3)) < base.Ui32(v127) {
		v109 = v125
		v110 = v123
		v111 = v127
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v138 = v131
	v139 = v132
	v140 = v133
	goto L24
L45:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v147)
	if v147 == int32(0) {
		v158 = v143
		v159 = v144
		goto L23
	} else {
		goto L47
	}
L46:
	;
	v158 = v154
	v159 = v152
	goto L23
L47:
	;
	v151 = int32(1)
	v152 = v144 + v151
	v154 = v143 + v151
	v156 = v145 - v151
	if v156 != 0 {
		v143 = v154
		v144 = v152
		v145 = v156
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	F_pfree(m, v52)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	if v196-v197 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v181 = v52
	v182 = v172
	goto L53
L53:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	if v186 == int32(0) {
		v196 = v186
		v197 = v185
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v196 = v186
	v197 = v185
	goto L51
L55:
	;
	v189 = int32(1)
	if v186 == v185 {
		v181 = v181 + v189
		v182 = v182 + v189
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v202 = F_substitute_path_macro(m, v52, int32(_a_F_get_extension_control_directories_1), v23)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v52
	v206 = F_psprintf(m, int32(_a_F_get_extension_control_directories_0), v11)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v208 = v202
	goto L49
L61:
	;
	v208 = v206
	goto L49
L62:
	;
	F_canonicalize_path_enc(m, v208)
	mBase = m.M
	v213 = F_lappend(m, v38, v208)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v49))))
	if v216 != 0 {
		v35 = v35 + v51
		v38 = v213
		goto L10
	} else {
		goto L64
	}
L64:
	;
	goto L11
}
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v859 int32
	_ = v859
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(240)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v3
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L12
	} else {
		goto L259
	}
L2:
	;
	v201 = F_strlen(m, v192)
	mBase = m.M
	v204 = F_pnstrdup(m, v192, v201-int32(10))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L56
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v56
	v77 = F_psprintf(m, int32(_a_F_parse_extension_control_file_0), v14+int32(208))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L26
	}
L4:
	;
	if v69 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v57 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L8:
	;
	v42 = F_palloc(m, int32(1024))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L19
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_pstrdup(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(47) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v40 = v24
	goto L8
L14:
	;
	v29 = F_pstrdup(m, v20)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v31
	v37 = F_psprintf(m, int32(_a_F_parse_extension_control_file_1), v14+int32(192))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v40 = v29
	goto L8
L18:
	;
	v40 = v37
	goto L8
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+184)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+180)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v40
	v52 = F_pg_snprintf(m, v42, int32(1024), int32(_a_F_parse_extension_control_file_2), v14+int32(176))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_pfree(m, v40)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v69 = v42
	goto L4
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v57
	v65 = F_psprintf(m, int32(_a_F_parse_extension_control_file_3), v14+int32(224))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v69 = v65
	goto L4
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v192 = v72
	v195 = v69
	goto L2
L25:
	;
	if v130 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L26:
	;
	v79 = F_get_extension_control_directories(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v81 = int32(0)
	v82 = m.G0
	v84 = v82 - int32(32)
	m.G0 = v84
	if v79 == v81 {
		v130 = v81
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L44
	}
L29:
	;
	m.G0 = v84 + int32(32)
	goto L25
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v88 <= int32(0) {
		v130 = v81
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v99 = v3
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v99<<(uint(int32(2))%32))))
	v107 = F_pstrdup(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v130 = v81
	goto L29
L34:
	;
	F_canonicalize_path_enc(m, v107)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110 != int32(47) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107
	v116 = F_psprintf(m, int32(_a_F_parse_extension_control_file_1), v84)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v118 = F_pg_file_exists(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v118 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v130 = v116
	goto L29
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v107)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	F_pfree(m, v116)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v125 = v99 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v125 < v126 {
		v99 = v125
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = int32(_a_F_parse_extension_control_file_4)
	F_errmsg(m, int32(_a_F_parse_extension_control_file_5), v84+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(4047), int32(_a_F_parse_extension_control_file_7))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v166 = F_strlen(m, v130)
	mBase = m.M
	v173 = v166 + int32(1)
	goto L51
L49:
	;
	v187 = F_pnstrdup(m, v130, v185-v130)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L55
	}
L50:
	;
	goto L49
L51:
	;
	v175 = int32(0)
	if v173 == v175 {
		v185 = v175
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v185 = v180
	goto L50
L53:
	;
	v179 = v173 - int32(1)
	v180 = v130 + v179
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != int32(47) {
		v173 = v179
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v187
	v192 = v187
	v195 = v130
	goto L2
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v204
	v208 = F_AllocateFile(m, v195, int32(_a_F_parse_extension_control_file_8))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v195)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L12
	} else {
		goto L258
	}
L58:
	;
	if v208 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if l1 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v239 = F_ParseConfigFp(m, v208, v195, int32(0), int32(21), v14+int32(236), v14+int32(232))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L70
	}
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_parse_extension_control_file[0]))
	if v213 == int32(44) {
		goto L57
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v195
	F_errmsg(m, int32(_a_F_parse_extension_control_file_9), v14+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(696), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v241 = F_FreeFile(m, v208)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	if v243 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v256 = v243
	goto L75
L73:
	;
	v822 = int32(0)
	goto L74
L74:
	;
	F_FreeConfigVariables(m, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L12
	} else {
		goto L251
	}
L75:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v266 = int32(_a_F_parse_extension_control_file_11)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[1])))
	if base.B2i32(v269 == int32(0))|base.B2i32(v269 != v272) != 0 {
		v290 = v269
		v291 = v272
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	v822 = v809
	goto L74
L77:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v256)+24))
	if v808 != 0 {
		v256 = v808
		goto L75
	} else {
		goto L250
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L12
	} else {
		goto L246
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L12
	} else {
		goto L242
	}
L80:
	;
	if v290-v291 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	goto L80
L82:
	;
	v275 = v265
	v276 = v266
	goto L83
L83:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	if v280 == int32(0) {
		v290 = v280
		v291 = v279
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v290 = v280
	v291 = v279
	goto L81
L85:
	;
	v283 = int32(1)
	if v280 == v279 {
		v275 = v275 + v283
		v276 = v276 + v283
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	if l1 != 0 {
		goto L79
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v299 = int32(_a_F_parse_extension_control_file_12)
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[2])))
	if base.B2i32(v302 == int32(0))|base.B2i32(v302 != v305) != 0 {
		v323 = v302
		v324 = v305
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v296 = F_pstrdup(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v296
	goto L77
L92:
	;
	if v323-v324 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	goto L92
L94:
	;
	v308 = v265
	v309 = v299
	goto L95
L95:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v313 == int32(0) {
		v323 = v313
		v324 = v312
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v323 = v313
	v324 = v312
	goto L93
L97:
	;
	v316 = int32(1)
	if v313 == v312 {
		v308 = v308 + v316
		v309 = v309 + v316
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if l1 != 0 {
		goto L78
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v332 = int32(_a_F_parse_extension_control_file_13)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[3])))
	if base.B2i32(v335 == int32(0))|base.B2i32(v335 != v338) != 0 {
		v356 = v335
		v357 = v338
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v329 = F_pstrdup(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v329
	goto L77
L104:
	;
	if v356-v357 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	goto L104
L106:
	;
	v341 = v265
	v342 = v332
	goto L107
L107:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if v346 == int32(0) {
		v356 = v346
		v357 = v345
		goto L105
	} else {
		goto L109
	}
L108:
	;
	v356 = v346
	v357 = v345
	goto L105
L109:
	;
	v349 = int32(1)
	if v346 == v345 {
		v341 = v341 + v349
		v342 = v342 + v349
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v362 = F_pstrdup(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v365 = int32(_a_F_parse_extension_control_file_14)
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[4])))
	if base.B2i32(v368 == int32(0))|base.B2i32(v368 != v371) != 0 {
		v389 = v368
		v390 = v371
		goto L116
	} else {
		goto L117
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v362
	goto L77
L115:
	;
	if v389-v390 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L116:
	;
	goto L115
L117:
	;
	v374 = v265
	v375 = v365
	goto L118
L118:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+1)))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	if v379 == int32(0) {
		v389 = v379
		v390 = v378
		goto L116
	} else {
		goto L120
	}
L119:
	;
	v389 = v379
	v390 = v378
	goto L116
L120:
	;
	v382 = int32(1)
	if v379 == v378 {
		v374 = v374 + v382
		v375 = v375 + v382
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v395 = F_pstrdup(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L12
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v398 = int32(_a_F_parse_extension_control_file_15)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[5])))
	if base.B2i32(v401 == int32(0))|base.B2i32(v401 != v404) != 0 {
		v422 = v401
		v423 = v404
		goto L127
	} else {
		goto L128
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v395
	goto L77
L126:
	;
	if v422-v423 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	goto L126
L128:
	;
	v407 = v265
	v408 = v398
	goto L129
L129:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+1)))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+1)))
	if v412 == int32(0) {
		v422 = v412
		v423 = v411
		goto L127
	} else {
		goto L131
	}
L130:
	;
	v422 = v412
	v423 = v411
	goto L127
L131:
	;
	v415 = int32(1)
	if v412 == v411 {
		v407 = v407 + v415
		v408 = v408 + v415
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v428 = F_pstrdup(m, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L12
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v431 = int32(_a_F_parse_extension_control_file_16)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[6])))
	if base.B2i32(v434 == int32(0))|base.B2i32(v434 != v437) != 0 {
		v455 = v434
		v456 = v437
		goto L138
	} else {
		goto L139
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v428
	goto L77
L137:
	;
	if v455-v456 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	v440 = v265
	v441 = v431
	goto L140
L140:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)))
	if v445 == int32(0) {
		v455 = v445
		v456 = v444
		goto L138
	} else {
		goto L142
	}
L141:
	;
	v455 = v445
	v456 = v444
	goto L138
L142:
	;
	v448 = int32(1)
	if v445 == v444 {
		v440 = v440 + v448
		v441 = v441 + v448
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v461 = F_strlen(m, v460)
	mBase = m.M
	v462 = F_parse_bool_with_len(m, v460, v461, l0+int32(32))
	mBase = m.M
	goto L147
L145:
	;
	goto L146
L146:
	;
	v482 = int32(_a_F_parse_extension_control_file_17)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[7])))
	if base.B2i32(v485 == int32(0))|base.B2i32(v485 != v488) != 0 {
		v506 = v485
		v507 = v488
		goto L154
	} else {
		goto L155
	}
L147:
	;
	if v462 != 0 {
		goto L77
	} else {
		goto L148
	}
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L12
	} else {
		goto L150
	}
L150:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v470
	F_errmsg(m, int32(_a_F_parse_extension_control_file_18), v14-int32(-64))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(751), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	if v506-v507 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	goto L153
L155:
	;
	v491 = v265
	v492 = v482
	goto L156
L156:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+1)))
	if v496 == int32(0) {
		v506 = v496
		v507 = v495
		goto L154
	} else {
		goto L158
	}
L157:
	;
	v506 = v496
	v507 = v495
	goto L154
L158:
	;
	v499 = int32(1)
	if v496 == v495 {
		v491 = v491 + v499
		v492 = v492 + v499
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v512 = F_strlen(m, v511)
	mBase = m.M
	v513 = F_parse_bool_with_len(m, v511, v512, l0+int32(33))
	mBase = m.M
	goto L163
L161:
	;
	goto L162
L162:
	;
	v533 = int32(_a_F_parse_extension_control_file_19)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v539 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[8])))
	if base.B2i32(v536 == int32(0))|base.B2i32(v536 != v539) != 0 {
		v557 = v536
		v558 = v539
		goto L170
	} else {
		goto L171
	}
L163:
	;
	if v513 != 0 {
		goto L77
	} else {
		goto L164
	}
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v521
	F_errmsg(m, int32(_a_F_parse_extension_control_file_18), v14+int32(80))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(759), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L12
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	if v557-v558 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L170:
	;
	goto L169
L171:
	;
	v542 = v265
	v543 = v533
	goto L172
L172:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+1)))
	if v547 == int32(0) {
		v557 = v547
		v558 = v546
		goto L170
	} else {
		goto L174
	}
L173:
	;
	v557 = v547
	v558 = v546
	goto L170
L174:
	;
	v550 = int32(1)
	if v547 == v546 {
		v542 = v542 + v550
		v543 = v543 + v550
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v563 = F_strlen(m, v562)
	mBase = m.M
	v564 = F_parse_bool_with_len(m, v562, v563, l0+int32(34))
	mBase = m.M
	goto L179
L177:
	;
	goto L178
L178:
	;
	v584 = int32(_a_F_parse_extension_control_file_20)
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v590 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[9])))
	if base.B2i32(v587 == int32(0))|base.B2i32(v587 != v590) != 0 {
		v608 = v587
		v609 = v590
		goto L186
	} else {
		goto L187
	}
L179:
	;
	if v564 != 0 {
		goto L77
	} else {
		goto L180
	}
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L12
	} else {
		goto L181
	}
L181:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L12
	} else {
		goto L182
	}
L182:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v572
	F_errmsg(m, int32(_a_F_parse_extension_control_file_18), v14+int32(96))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L12
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(767), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L12
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	if v608-v609 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L186:
	;
	goto L185
L187:
	;
	v593 = v265
	v594 = v584
	goto L188
L188:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+1)))
	if v598 == int32(0) {
		v608 = v598
		v609 = v597
		goto L186
	} else {
		goto L190
	}
L189:
	;
	v608 = v598
	v609 = v597
	goto L186
L190:
	;
	v601 = int32(1)
	if v598 == v597 {
		v593 = v593 + v601
		v594 = v594 + v601
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v615 = F_pg_char_to_encoding_private(m, v613)
	mBase = m.M
	if base.Ui32(int32(35)) <= base.Ui32(v615) {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v641 = int32(_a_F_parse_extension_control_file_21)
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[10])))
	if base.B2i32(v644 == int32(0))|base.B2i32(v644 != v647) != 0 {
		v665 = v644
		v666 = v647
		goto L205
	} else {
		goto L206
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v618
	if int32(0) <= v618 {
		goto L77
	} else {
		goto L199
	}
L196:
	;
	v618 = int32(-1)
	goto L198
L197:
	;
	v618 = v615
	goto L198
L198:
	;
	goto L195
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L12
	} else {
		goto L200
	}
L200:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L12
	} else {
		goto L201
	}
L201:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v629
	F_errmsg(m, int32(_a_F_parse_extension_control_file_22), v14+int32(112))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L12
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(776), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L12
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	if v665-v666 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L205:
	;
	goto L204
L206:
	;
	v650 = v265
	v651 = v641
	goto L207
L207:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+1)))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650)+1)))
	if v655 == int32(0) {
		v665 = v655
		v666 = v654
		goto L205
	} else {
		goto L209
	}
L208:
	;
	v665 = v655
	v666 = v654
	goto L205
L209:
	;
	v658 = int32(1)
	if v655 == v654 {
		v650 = v650 + v658
		v651 = v651 + v658
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v671 = F_pstrdup(m, v670)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L12
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v695 = int32(_a_F_parse_extension_control_file_23)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[11])))
	if base.B2i32(v698 == int32(0))|base.B2i32(v698 != v701) != 0 {
		v719 = v698
		v720 = v701
		goto L222
	} else {
		goto L223
	}
L214:
	;
	v674 = F_SplitIdentifierString(m, v671, int32(44), l0+int32(40))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L12
	} else {
		goto L215
	}
L215:
	;
	if v674 != 0 {
		goto L77
	} else {
		goto L216
	}
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L12
	} else {
		goto L217
	}
L217:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L12
	} else {
		goto L218
	}
L218:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v683
	F_errmsg(m, int32(_a_F_parse_extension_control_file_24), v14+int32(128))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L12
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(790), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L12
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	if v719-v720 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L222:
	;
	goto L221
L223:
	;
	v704 = v265
	v705 = v695
	goto L224
L224:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705)+1)))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+1)))
	if v709 == int32(0) {
		v719 = v709
		v720 = v708
		goto L222
	} else {
		goto L226
	}
L225:
	;
	v719 = v709
	v720 = v708
	goto L222
L226:
	;
	v712 = int32(1)
	if v709 == v708 {
		v704 = v704 + v712
		v705 = v705 + v712
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v725 = F_pstrdup(m, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L12
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L238
	}
L231:
	;
	v728 = F_SplitIdentifierString(m, v725, int32(44), l0+int32(44))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L12
	} else {
		goto L232
	}
L232:
	;
	if v728 != 0 {
		goto L77
	} else {
		goto L233
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L12
	} else {
		goto L234
	}
L234:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L12
	} else {
		goto L235
	}
L235:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v737
	F_errmsg(m, int32(_a_F_parse_extension_control_file_24), v14+int32(144))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L12
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(805), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L12
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L12
	} else {
		goto L239
	}
L239:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v756
	F_errmsg(m, int32(_a_F_parse_extension_control_file_25), v14+int32(160))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L12
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(812), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L12
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L12
	} else {
		goto L243
	}
L243:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v776
	F_errmsg(m, int32(_a_F_parse_extension_control_file_26), v14+int32(32))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L12
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(719), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v795
	F_errmsg(m, int32(_a_F_parse_extension_control_file_26), v14+int32(48))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L12
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(729), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L12
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	goto L76
L251:
	;
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v825 != int32(1) {
		goto L57
	} else {
		goto L252
	}
L252:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v828 == int32(0) {
		goto L57
	} else {
		goto L253
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L12
	} else {
		goto L254
	}
L254:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L12
	} else {
		goto L255
	}
L255:
	;
	F_errmsg(m, int32(_a_F_parse_extension_control_file_27), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L12
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(820), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L12
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	m.G0 = v14 + int32(240)
	return
L259:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L12
	} else {
		goto L260
	}
L260:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v881
	F_errmsg(m, int32(_a_F_parse_extension_control_file_28), v14)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L12
	} else {
		goto L261
	}
L261:
	;
	F_errhint(m, int32(_a_F_parse_extension_control_file_29), int32(0))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L12
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(673), int32(_a_F_parse_extension_control_file_10))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L12
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
