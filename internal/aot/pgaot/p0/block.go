package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BlockRefTableGetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
	v23 = int32(-1636608416)
	if v9&int32(3) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v284 = (v277 ^ v269 - base.I32_rotl(v277, int32(24))) & v283
	v287 = v282 + v284*int32(40)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+20)))
	if v288 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L2:
	;
	v255 = int32(14)
	v257 = v251 ^ v252 - base.I32_rotl(v251, v255)
	v261 = v257 ^ v250 - base.I32_rotl(v257, int32(11))
	v265 = v261 ^ v251 - base.I32_rotl(v261, int32(25))
	v269 = v265 ^ v257 - base.I32_rotl(v265, int32(16))
	v273 = v269 ^ v261 - base.I32_rotl(v269, int32(4))
	v277 = v273 ^ v265 - base.I32_rotl(v273, v255)
	goto L1
L3:
	;
	switch v177 - int32(1) {
	case 0:
		v243 = v168
		v244 = v169
		v245 = v173
		goto L30
	case 1:
		v236 = v168
		v237 = v169
		v238 = v173
		goto L31
	case 2:
		v229 = v168
		v230 = v169
		v231 = v173
		goto L32
	case 3:
		v223 = v169
		v224 = v173
		goto L33
	case 4:
		v219 = v169
		v220 = v173
		goto L34
	case 5:
		v213 = v169
		v214 = v173
		goto L35
	case 6:
		v207 = v169
		v208 = v173
		goto L36
	case 7:
		v202 = v173
		goto L37
	case 8:
		v197 = v173
		goto L38
	case 9:
		v192 = v173
		goto L39
	case 10:
		goto L40
	default:
		v250 = v168
		v251 = v169
		v252 = v173
		goto L2
	}
L4:
	;
	v132 = v9
	v133 = v8
	v134 = v23
	v135 = v23
	v136 = v23
	goto L27
L5:
	;
	goto L4
L6:
	;
	goto L7
L7:
	;
	goto L11
L9:
	;
	switch v75 - int32(1) {
	case 0:
		v129 = v66
		goto L16
	case 1:
		v124 = v66
		goto L17
	case 2:
		goto L18
	case 3:
		v117 = v67
		goto L19
	case 4:
		v114 = v67
		goto L20
	case 5:
		v109 = v67
		goto L21
	case 6:
		goto L22
	case 7:
		v100 = v71
		goto L23
	case 8:
		v95 = v71
		goto L24
	case 9:
		v90 = v71
		goto L25
	case 10:
		goto L26
	default:
		v250 = v66
		v251 = v67
		v252 = v71
		goto L2
	}
L11:
	;
	goto L12
L12:
	;
	v30 = v9
	v31 = v8
	v32 = v23
	v33 = v23
	v34 = v23
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v37 = v36 + v33
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v41 = v40 + v34
	v43 = int32(4)
	v45 = v38 + v32 - v41 ^ base.I32_rotl(v41, v43)
	v49 = v37 - v45 ^ base.I32_rotl(v45, int32(6))
	v50 = v41 + v37
	v51 = v45 + v50
	v52 = v49 + v51
	v56 = v50 - v49 ^ base.I32_rotl(v49, int32(8))
	v60 = v51 - v56 ^ base.I32_rotl(v56, int32(16))
	v64 = v52 - v60 ^ base.I32_rotl(v60, int32(19))
	v65 = v56 + v52
	v66 = v60 + v65
	v67 = v64 + v66
	v71 = v65 - v64 ^ base.I32_rotl(v64, v43)
	v72 = int32(12)
	v73 = v30 + v72
	v75 = v31 - v72
	if base.Ui32(int32(11)) < base.Ui32(v75) {
		v30 = v73
		v31 = v75
		v32 = v66
		v33 = v67
		v34 = v71
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L9
L15:
	;
	goto L14
L16:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v250 = v129 + v130
	v251 = v67
	v252 = v71
	goto L2
L17:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v129 = v125<<(uint(int32(8))%32) + v124
	goto L16
L18:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+2)))
	v124 = v120<<(uint(int32(16))%32) + v66
	goto L17
L19:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v250 = v118 + v66
	v251 = v117
	v252 = v71
	goto L2
L20:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
	v117 = v114 + v115
	goto L19
L21:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+5)))
	v114 = v110<<(uint(int32(8))%32) + v109
	goto L20
L22:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+6)))
	v109 = v105<<(uint(int32(16))%32) + v67
	goto L21
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v250 = v101 + v66
	v251 = v103 + v67
	v252 = v100
	goto L2
L24:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
	v100 = v96<<(uint(int32(8))%32) + v95
	goto L23
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+9)))
	v95 = v91<<(uint(int32(16))%32) + v90
	goto L24
L26:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+10)))
	v90 = v86<<(uint(int32(24))%32) + v71
	goto L25
L27:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v139 = v138 + v135
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v143 = v142 + v136
	v145 = int32(4)
	v147 = v140 + v134 - v143 ^ base.I32_rotl(v143, v145)
	v151 = v139 - v147 ^ base.I32_rotl(v147, int32(6))
	v152 = v143 + v139
	v153 = v147 + v152
	v154 = v151 + v153
	v158 = v152 - v151 ^ base.I32_rotl(v151, int32(8))
	v162 = v153 - v158 ^ base.I32_rotl(v158, int32(16))
	v166 = v154 - v162 ^ base.I32_rotl(v162, int32(19))
	v167 = v158 + v154
	v168 = v162 + v167
	v169 = v166 + v168
	v173 = v167 - v166 ^ base.I32_rotl(v166, v145)
	v174 = int32(12)
	v175 = v132 + v174
	v177 = v133 - v174
	if base.Ui32(int32(11)) < base.Ui32(v177) {
		v132 = v175
		v133 = v177
		v134 = v168
		v135 = v169
		v136 = v173
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	goto L28
L30:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v250 = v243 + v246
	v251 = v244
	v252 = v245
	goto L2
L31:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	v243 = v239<<(uint(int32(8))%32) + v236
	v244 = v237
	v245 = v238
	goto L30
L32:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	v236 = v232<<(uint(int32(16))%32) + v229
	v237 = v230
	v238 = v231
	goto L31
L33:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
	v229 = v225<<(uint(int32(24))%32) + v168
	v230 = v223
	v231 = v224
	goto L32
L34:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
	v223 = v219 + v221
	v224 = v220
	goto L33
L35:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
	v219 = v215<<(uint(int32(8))%32) + v213
	v220 = v214
	goto L34
L36:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
	v213 = v209<<(uint(int32(16))%32) + v207
	v214 = v208
	goto L35
L37:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+7)))
	v207 = v203<<(uint(int32(24))%32) + v169
	v208 = v202
	goto L36
L38:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
	v202 = v198<<(uint(int32(8))%32) + v197
	goto L37
L39:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+9)))
	v197 = v193<<(uint(int32(16))%32) + v192
	goto L38
L40:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+10)))
	v192 = v188<<(uint(int32(24))%32) + v173
	goto L39
L41:
	;
	m.G0 = v9 + int32(16)
	return v380
L42:
	;
	v380 = int32(0)
	goto L41
L43:
	;
	v291 = v284
	v292 = v287
	goto L44
L44:
	;
	v297 = int32(16)
	goto L49
L45:
	;
	if v292 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L46:
	;
	if v359 != 0 {
		goto L64
	} else {
		goto L65
	}
L47:
	;
	v359 = int32(0)
	goto L46
L48:
	;
	v333 = v328
	v334 = v329
	v335 = v330
	goto L58
L49:
	;
	if (v292|v9)&int32(3) != 0 {
		v328 = v292
		v329 = v9
		v330 = v297
		goto L48
	} else {
		goto L52
	}
L51:
	;
	if v318 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L52:
	;
	v305 = v292
	v306 = v9
	v307 = v297
	goto L53
L53:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v310 != v311 {
		v328 = v305
		v329 = v306
		v330 = v307
		goto L48
	} else {
		goto L55
	}
L54:
	;
	goto L51
L55:
	;
	v313 = int32(4)
	v314 = v306 + v313
	v316 = v305 + v313
	v318 = v307 - v313
	if base.Ui32(int32(3)) < base.Ui32(v318) {
		v305 = v316
		v306 = v314
		v307 = v318
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v328 = v316
	v329 = v314
	v330 = v318
	goto L48
L58:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v338 == v339 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v359 = v338 - v339
	goto L46
L60:
	;
	v341 = int32(1)
	v346 = v335 - v341
	if v346 != 0 {
		v333 = v333 + v341
		v334 = v334 + v341
		v335 = v346
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L47
L64:
	;
	v362 = (v291 + int32(1)) & v283
	v365 = v282 + v362*int32(40)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+20)))
	if v366 != 0 {
		v291 = v362
		v292 = v365
		goto L44
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L45
L67:
	;
	goto L42
L68:
	;
	v380 = int32(0)
	goto L41
L69:
	;
	goto L70
L70:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v370
	v380 = v292
	goto L41
}
func F_BlockRefTableRead(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if int32(0) < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = l0 + int32(8)
	v17 = l1
	v18 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(32)
	return
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[997])))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[998])))
	if v24 < v23 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v76 {
		v17 = v75
		v18 = v76
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v27 = v23 - v24
	if v18 < v27 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(65536)) <= base.Ui32(v18) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v29 = v18
	goto L12
L11:
	;
	v29 = v27
	goto L12
L12:
	;
	if v29 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[999])))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[998])))
	v35 = m.Env.Pgmem_crc32c(m, v32, v15+v33, v29)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[999]))) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[998])))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[998]))) = v37 + v29
	v75 = v31 + v29
	v76 = v18 - v29
	goto L6
L14:
	;
	v30 = F__emscripten_memcpy_bulkmem(m, v17, v24+v15, v29)
	mBase = m.M
	v31 = v30
	goto L16
L15:
	;
	v31 = v17
	goto L16
L16:
	;
	goto L13
L17:
	;
	v46 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, v42, v17, v18)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v61 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, v42, v15, int32(65536))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L24
	}
L20:
	;
	return
L21:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[999])))
	v49 = m.Env.Pgmem_crc32c(m, v48, v17, v46)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[999]))) = v49
	v51 = v18 - v46
	v52 = v17 + v46
	if v46 != 0 {
		v75 = v52
		v76 = v51
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[487])))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1000])))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[488])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55
	m.T0[v54].(func(*base.Module, int32, int32, int32))(m, v53, int32(19020), v10)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v75 = v52
	v76 = v51
	goto L6
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[998]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[997]))) = v61
	if v61 != 0 {
		v75 = v17
		v76 = v18
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[487])))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1000])))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[488])))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v68
	m.T0[v67].(func(*base.Module, int32, int32, int32))(m, v66, int32(19020), v10+int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v75 = v17
	v76 = v18
	goto L6
L27:
	;
	goto L5
}
func F_BlockRefTableSetLimitBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = v14 + int32(40)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l2
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v20
	v31 = F_blockreftable_insert(m, v23, v14+int32(8), v14+int32(31))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return
	} else {
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+31)))
		if v33 == int32(0) {
			v36 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v36
			*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = l3
			*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = v36
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
			if base.Ui32(v41) <= base.Ui32(l3) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = l3
				v45 = int32(base.Ui32(l3) >> (uint(int32(16)) % 32))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
				if base.Ui32(v46) <= base.Ui32(v45) {
				} else {
					v49 = v45 + int32(1)
					if base.Ui32(v49) < base.Ui32(v46) {
						v53 = v49
						for {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
							v63 = int32(1)
							v66 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v62+v53<<(uint(v63)%32)))) = uint16(v66)
							v69 = v53 + v63
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
							if base.Ui32(v69) < base.Ui32(v70) {
								v53 = v69
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
					v86 = v83 + v45<<(uint(int32(1))%32)
					v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
					if v87 != 0 {
						v89 = l3 & int32(65535)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v45<<(uint(int32(2))%32))))
						if v87 == int32(4096) {
							if l3&int32(1) != 0 {
								v149 = v94 + int32(base.Ui32(v89)>>(uint(int32(3))%32))&int32(8190)
								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149))))
								v155 = v150 & base.I32_rotl(int32(-2), l3&int32(15))
								*(*uint16)(unsafe.Add(mBase, uint32(v149))) = uint16(v155)
								v159 = v89 + int32(1)
							} else {
								v159 = v89
							}
							if v89 == int32(65535) {
							} else {
								v163 = v159
								for {
									v173 = int32(3)
									v175 = int32(536870910)
									v177 = v94 + int32(base.Ui32(v163)>>(uint(v173)%32))&v175
									v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177))))
									v179 = int32(-2)
									v180 = int32(15)
									v183 = v178 & base.I32_rotl(v179, v163&v180)
									*(*uint16)(unsafe.Add(mBase, uint32(v177))) = uint16(v183)
									v186 = v163 + int32(1)
									v191 = v94 + int32(base.Ui32(v186)>>(uint(v173)%32))&v175
									v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191))))
									v197 = v192 & base.I32_rotl(v179, v186&v180)
									*(*uint16)(unsafe.Add(mBase, uint32(v191))) = uint16(v197)
									v200 = v163 + int32(2)
									if v200 != int32(65536) {
										v163 = v200
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v97 = int32(0)
							v101 = v97
							v103 = v97
							v107 = v83
							for {
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+v101<<(uint(int32(1))%32)))))
								if base.Ui32(v113) < base.Ui32(v89) {
									v115 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v94+v103<<(uint(v115)%32)))) = uint16(v113)
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
									v122 = v103 + v115
									v123 = v119
								} else {
									v122 = v103
									v123 = v107
								}
								v124 = int32(1)
								v125 = v101 + v124
								v128 = v123 + v45<<(uint(v124)%32)
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
								if base.Ui32(v125) < base.Ui32(v129) {
									v101 = v125
									v103 = v122
									v107 = v123
									continue
								} else {
									break
								}
								break
							}
							v135 = v122
							v140 = v128
							*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v135)
						}
					} else {
						v135 = v87
						v140 = v86
						*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v135)
					}
				}
			}
		}
		m.G0 = v14 + int32(48)
		return
	}
}
