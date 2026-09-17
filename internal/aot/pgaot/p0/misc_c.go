package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CMPTRGM_CHOOSE(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CMPTRGM_CHOOSE[0]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+256)))
	if v9 != 0 {
		v10 = int32(_a_F_CMPTRGM_CHOOSE_0)
	} else {
		v10 = int32(_a_F_CMPTRGM_CHOOSE_1)
	}
	*(*int32)(unsafe.Add(mBase, _c_F_CMPTRGM_CHOOSE[1])) = v10
	v12 = m.T0[v10].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_CNStoBIG5(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	v3 = int32(0)
	v5 = l0 & int32(_a_F_CNStoBIG5_0)
	switch l1 - int32(149) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L12
	}
L1:
	;
	return v319 & int32(_a_F_CNStoBIG5_1)
L2:
	;
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316))))
	v319 = v317
	goto L1
L3:
	;
	if v5 != int32(_a_F_CNStoBIG5_2) {
		v319 = v3
		goto L1
	} else {
		goto L96
	}
L4:
	;
	v316 = int32(_a_F_CNStoBIG5_3)
	goto L2
L5:
	;
	v316 = int32(_a_F_CNStoBIG5_4)
	goto L2
L6:
	;
	v310 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[0])))
	v319 = v310
	goto L1
L7:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[1])))
	v319 = v308
	goto L1
L8:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[2])))
	v319 = v306
	goto L1
L9:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[3])))
	v319 = v304
	goto L1
L10:
	;
	v180 = int32(47)
	v182 = int32(23)
	v186 = int32(0)
	goto L64
L11:
	;
	v47 = int32(24)
	v49 = int32(12)
	v53 = int32(0)
	goto L30
L12:
	;
	switch l1 - int32(246) {
	case 0:
		goto L13
	case 1:
		goto L14
	default:
		v319 = v3
		goto L1
	}
L13:
	;
	if base.Ui32(v5) <= base.Ui32(int32(_a_F_CNStoBIG5_5)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	switch v5 - int32(_a_F_CNStoBIG5_6) {
	case 0:
		v316 = int32(_a_F_CNStoBIG5_7)
		goto L2
	case 1:
		goto L5
	case 2, 3, 4, 5, 6:
		v319 = v3
		goto L1
	case 7:
		goto L4
	default:
		goto L3
	}
L15:
	;
	if v5 == int32(_a_F_CNStoBIG5_8) {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(v5) <= base.Ui32(int32(_a_F_CNStoBIG5_9)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v5 == int32(_a_F_CNStoBIG5_10) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	if v5 != int32(_a_F_CNStoBIG5_11) {
		v319 = v3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[4])))
	v319 = v22
	goto L1
L21:
	;
	if v5 == int32(_a_F_CNStoBIG5_12) {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v5 == int32(_a_F_CNStoBIG5_13) {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	if v5 != int32(_a_F_CNStoBIG5_14) {
		v319 = v3
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[5])))
	v319 = v30
	goto L1
L26:
	;
	if v5 != int32(_a_F_CNStoBIG5_15) {
		v319 = v3
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CNStoBIG5[6])))
	v319 = v36
	goto L1
L28:
	;
	v319 = v167 & int32(_a_F_CNStoBIG5_1)
	goto L1
L29:
	;
	goto L28
L30:
	;
	v55 = v49 << (uint(int32(2)) % 32)
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_CNStoBIG5[7]))))
	v58 = base.B2i32(base.Ui32(v5) < base.Ui32(v57))
	if base.Ui32(v5) < base.Ui32(v57) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v167 = int32(0)
	goto L29
L32:
	;
	goto L31
L33:
	;
	if base.Ui32(v5) < base.Ui32(v57) {
		goto L55
	} else {
		goto L56
	}
L34:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_CNStoBIG5[8]))))
	if base.Ui32(v59) <= base.Ui32(v5) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_CNStoBIG5[9]))))
	if v61 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v66 = v5 - v57&int32(_a_F_CNStoBIG5_16)
	if base.Ui32(int32(_a_F_CNStoBIG5_17)) <= base.Ui32(v5) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v69 = int32(255)
	v70 = v5 & v69
	v72 = v57 & v69
	v82 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v72))
	if base.Ui32(int32(160)) < base.Ui32(v72) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v109 = int32(255)
	v110 = v61 & v109
	if base.Ui32(int32(160)) < base.Ui32(v110) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v83 = int32(0)
	goto L42
L41:
	;
	v83 = int32(-34)
	goto L42
L42:
	;
	if base.Ui32(int32(160)) < base.Ui32(v72) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v86 = int32(34)
	goto L45
L44:
	;
	v86 = int32(0)
	goto L45
L45:
	;
	if base.Ui32(int32(160)) < base.Ui32(v70) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v89 = v83
	goto L48
L47:
	;
	v89 = v86
	goto L48
L48:
	;
	v94 = int32(33)
	v95 = v70 - v72 + v66>>(uint(int32(8))%32)*int32(157) + v89 + v61&int32(255) - v94
	v96 = int32(94)
	v97 = base.I32_div_s(v95, v96)
	v167 = v95 - v97*v96 + v61&int32(_a_F_CNStoBIG5_16) + v97<<(uint(int32(8))%32) + v94
	goto L29
L49:
	;
	v126 = int32(_a_F_CNStoBIG5_18)
	goto L51
L50:
	;
	v126 = int32(_a_F_CNStoBIG5_19)
	goto L51
L51:
	;
	v127 = v110 + (v5&v109 - v57&v109 + int32(base.Ui32(v66)>>(uint(int32(8))%32))*int32(94)) + v126
	v129 = int32(157)
	v130 = base.I32_div_s(base.I32_extend16_s(v127), v129)
	v133 = v127 - v130*v129
	if int32(62) < base.I32_extend16_s(v133) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v145 = int32(98)
	goto L54
L53:
	;
	v145 = int32(64)
	goto L54
L54:
	;
	v167 = v133 + v61&int32(_a_F_CNStoBIG5_16) + v130<<(uint(int32(8))%32) + v145
	goto L29
L55:
	;
	v149 = v53
	goto L57
L56:
	;
	v149 = v49 + int32(1)
	goto L57
L57:
	;
	if base.Ui32(v5) < base.Ui32(v57) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v152 = v49 - int32(1)
	goto L60
L59:
	;
	v152 = v47
	goto L60
L60:
	;
	if v149 <= v152 {
		v47 = v152
		v49 = (v149 + v152) >> (uint(int32(1)) % 32)
		v53 = v149
		goto L30
	} else {
		goto L61
	}
L61:
	;
	goto L32
L62:
	;
	v319 = v300 & int32(_a_F_CNStoBIG5_1)
	goto L1
L63:
	;
	goto L62
L64:
	;
	v188 = v182 << (uint(int32(2)) % 32)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_CNStoBIG5[10]))))
	v191 = base.B2i32(base.Ui32(v5) < base.Ui32(v190))
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v300 = int32(0)
	goto L63
L66:
	;
	goto L65
L67:
	;
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L89
	} else {
		goto L90
	}
L68:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_CNStoBIG5[11]))))
	if base.Ui32(v192) <= base.Ui32(v5) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_CNStoBIG5[12]))))
	if v194 == int32(0) {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v199 = v5 - v190&int32(_a_F_CNStoBIG5_16)
	if base.Ui32(int32(_a_F_CNStoBIG5_17)) <= base.Ui32(v5) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v202 = int32(255)
	v203 = v5 & v202
	v205 = v190 & v202
	v215 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v205))
	if base.Ui32(int32(160)) < base.Ui32(v205) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v242 = int32(255)
	v243 = v194 & v242
	if base.Ui32(int32(160)) < base.Ui32(v243) {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v216 = int32(0)
	goto L76
L75:
	;
	v216 = int32(-34)
	goto L76
L76:
	;
	if base.Ui32(int32(160)) < base.Ui32(v205) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v219 = int32(34)
	goto L79
L78:
	;
	v219 = int32(0)
	goto L79
L79:
	;
	if base.Ui32(int32(160)) < base.Ui32(v203) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v222 = v216
	goto L82
L81:
	;
	v222 = v219
	goto L82
L82:
	;
	v227 = int32(33)
	v228 = v203 - v205 + v199>>(uint(int32(8))%32)*int32(157) + v222 + v194&int32(255) - v227
	v229 = int32(94)
	v230 = base.I32_div_s(v228, v229)
	v300 = v228 - v230*v229 + v194&int32(_a_F_CNStoBIG5_16) + v230<<(uint(int32(8))%32) + v227
	goto L63
L83:
	;
	v259 = int32(_a_F_CNStoBIG5_18)
	goto L85
L84:
	;
	v259 = int32(_a_F_CNStoBIG5_19)
	goto L85
L85:
	;
	v260 = v243 + (v5&v242 - v190&v242 + int32(base.Ui32(v199)>>(uint(int32(8))%32))*int32(94)) + v259
	v262 = int32(157)
	v263 = base.I32_div_s(base.I32_extend16_s(v260), v262)
	v266 = v260 - v263*v262
	if int32(62) < base.I32_extend16_s(v266) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v278 = int32(98)
	goto L88
L87:
	;
	v278 = int32(64)
	goto L88
L88:
	;
	v300 = v266 + v194&int32(_a_F_CNStoBIG5_16) + v263<<(uint(int32(8))%32) + v278
	goto L63
L89:
	;
	v282 = v186
	goto L91
L90:
	;
	v282 = v182 + int32(1)
	goto L91
L91:
	;
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v285 = v182 - int32(1)
	goto L94
L93:
	;
	v285 = v180
	goto L94
L94:
	;
	if v282 <= v285 {
		v180 = v285
		v182 = (v282 + v285) >> (uint(int32(1)) % 32)
		v186 = v282
		goto L64
	} else {
		goto L95
	}
L95:
	;
	goto L66
L96:
	;
	v316 = int32(_a_F_CNStoBIG5_20)
	goto L2
}
func F_CheckAffix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v585 int32
	_ = v585
	v7 = int32(0)
	if l3 == v7 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v585
L2:
	;
	v585 = int32(0)
	goto L1
L3:
	;
	if v46&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v12&int32(2) == int32(0) {
		v46 = v12
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l3&int32(2) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20&int32(64) != 0 {
		v585 = v19
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if l3&int32(4) != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v20&int32(5) != int32(1) {
		v46 = v20
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v585 = v19
	goto L1
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29&int32(72) == int32(8) {
		v46 = v29
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(l3) < base.Ui32(int32(8)) {
		v46 = v34
		goto L3
	} else {
		goto L17
	}
L16:
	;
	goto L2
L17:
	;
	v37 = int32(0)
	if base.B2i32(v34&int32(17) == v37)|v34&int32(64) != 0 {
		v585 = v37
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v46 = v34
	goto L3
L19:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v311&int32(256) != 0 {
		goto L93
	} else {
		goto L94
	}
L20:
	;
	if (l0^l4)&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L22
L22:
	;
	if l5 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L23:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v129 = l4 + l1 - int32(base.Ui32(v124)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if (v130^v129)&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L24:
	;
	goto L23
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v102)
	if v102&int32(255) == int32(0) {
		goto L24
	} else {
		goto L40
	}
L26:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v101 = l0
	v102 = v54
	v103 = l4
	goto L25
L27:
	;
	goto L28
L28:
	;
	if l0&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v58 = l0
	v60 = l4
	goto L32
L30:
	;
	v72 = l0
	v74 = l4
	goto L31
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v79 = int32(-2139062144)
	if (int32(16843008)-v76|v76)&v79 != v79 {
		v101 = v72
		v102 = v76
		v103 = v74
		goto L25
	} else {
		goto L36
	}
L32:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v61)
	if v61 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L33:
	;
	v72 = v68
	v74 = v66
	goto L31
L34:
	;
	v65 = int32(1)
	v66 = v60 + v65
	v68 = v58 + v65
	if v68&int32(3) != 0 {
		v58 = v68
		v60 = v66
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v84 = v72
	v85 = v76
	v86 = v74
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v85
	v88 = int32(4)
	v89 = v86 + v88
	v91 = v84 + v88
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v96 = int32(-2139062144)
	if (int32(16843008)-v93|v93)&v96 == v96 {
		v84 = v91
		v85 = v93
		v86 = v89
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v101 = v91
	v102 = v93
	v103 = v89
	goto L25
L39:
	;
	goto L38
L40:
	;
	v110 = v101
	v112 = v103
	goto L41
L41:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)) = uint8(v113)
	v115 = int32(1)
	if v113 != 0 {
		v110 = v110 + v115
		v112 = v112 + v115
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L24
L43:
	;
	goto L42
L44:
	;
	if l5 == int32(0) {
		goto L19
	} else {
		goto L65
	}
L45:
	;
	goto L44
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v184)
	if v184&int32(255) == int32(0) {
		goto L45
	} else {
		goto L61
	}
L47:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v183 = v130
	v184 = v136
	v185 = v129
	goto L46
L48:
	;
	goto L49
L49:
	;
	if v130&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v140 = v130
	v142 = v129
	goto L53
L51:
	;
	v154 = v130
	v156 = v129
	goto L52
L52:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 != v161 {
		v183 = v154
		v184 = v158
		v185 = v156
		goto L46
	} else {
		goto L57
	}
L53:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v143)
	if v143 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L54:
	;
	v154 = v150
	v156 = v148
	goto L52
L55:
	;
	v147 = int32(1)
	v148 = v142 + v147
	v150 = v140 + v147
	if v150&int32(3) != 0 {
		v140 = v150
		v142 = v148
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v166 = v154
	v167 = v158
	v168 = v156
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v167
	v170 = int32(4)
	v171 = v168 + v170
	v173 = v166 + v170
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v178 = int32(-2139062144)
	if (int32(16843008)-v175|v175)&v178 == v178 {
		v166 = v173
		v167 = v175
		v168 = v171
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v183 = v173
	v184 = v175
	v185 = v171
	goto L46
L60:
	;
	goto L59
L61:
	;
	v192 = v183
	v194 = v185
	goto L62
L62:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)) = uint8(v195)
	v197 = int32(1)
	if v195 != 0 {
		v192 = v192 + v197
		v194 = v194 + v197
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L45
L64:
	;
	goto L63
L65:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l1 - int32(base.Ui32(v207)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)
	goto L19
L66:
	;
	if (v226^l4)&int32(3) != 0 {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v226 = v216
	goto L66
L68:
	;
	goto L69
L69:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v218 = F_strlen(m, v217)
	mBase = m.M
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if base.Ui32(v218+v219) <= base.Ui32(int32(base.Ui32(v46)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)) {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v226 = v217
	goto L66
L71:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v307 = F_strlen(m, l4)
	mBase = m.M
	v309 = F_strcpy(m, v307+l4, l0+int32(base.Ui32(v301)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0))
	mBase = m.M
	goto L92
L72:
	;
	goto L71
L73:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v281))) = uint8(v280)
	if v280&int32(255) == int32(0) {
		goto L72
	} else {
		goto L88
	}
L74:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v279 = v226
	v280 = v232
	v281 = l4
	goto L73
L75:
	;
	goto L76
L76:
	;
	if v226&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v236 = v226
	v238 = l4
	goto L80
L78:
	;
	v250 = v226
	v252 = l4
	goto L79
L79:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v257 = int32(-2139062144)
	if (int32(16843008)-v254|v254)&v257 != v257 {
		v279 = v250
		v280 = v254
		v281 = v252
		goto L73
	} else {
		goto L84
	}
L80:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	*(*uint8)(unsafe.Add(mBase, uint32(v238))) = uint8(v239)
	if v239 == int32(0) {
		goto L72
	} else {
		goto L82
	}
L81:
	;
	v250 = v246
	v252 = v244
	goto L79
L82:
	;
	v243 = int32(1)
	v244 = v238 + v243
	v246 = v236 + v243
	if v246&int32(3) != 0 {
		v236 = v246
		v238 = v244
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v262 = v250
	v263 = v254
	v264 = v252
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v263
	v266 = int32(4)
	v267 = v264 + v266
	v269 = v262 + v266
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v274 = int32(-2139062144)
	if (int32(16843008)-v271|v271)&v274 == v274 {
		v262 = v269
		v263 = v271
		v264 = v267
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v279 = v269
	v280 = v271
	v281 = v267
	goto L73
L87:
	;
	goto L86
L88:
	;
	v288 = v279
	v290 = v281
	goto L89
L89:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v291)
	v293 = int32(1)
	if v291 != 0 {
		v288 = v288 + v293
		v290 = v290 + v293
		goto L89
	} else {
		goto L91
	}
L90:
	;
	goto L72
L91:
	;
	goto L90
L92:
	;
	goto L19
L93:
	;
	return l4
L94:
	;
	goto L95
L95:
	;
	if v311&int32(512) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v317 = int32(0)
	v318 = m.G0
	v319 = int32(16)
	v320 = v318 - v319
	m.G0 = v320
	v323 = l2 + v319
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v325 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	goto L98
L98:
	;
	v553 = F_strlen(m, l4)
	mBase = m.M
	v558 = F_palloc(m, v553<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L105
	} else {
		goto L158
	}
L99:
	;
	if v531 == int32(0) {
		goto L2
	} else {
		goto L157
	}
L100:
	;
	v326 = l4
	v331 = v317
	goto L103
L101:
	;
	v348 = v317
	goto L102
L102:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v356 = int32(base.Ui32(v352)>>(uint(int32(1))%32)) & int32(_a_F_CheckAffix_1)
	if v348 < v356 {
		v531 = v7
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v336 = v331 + int32(1)
	v337 = F_pg_mblen_cstr(m, v326)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v348 = v336
	goto L102
L105:
	;
	return int32(0)
L106:
	;
	v341 = v337 + v326
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v342 != 0 {
		v326 = v341
		v331 = v336
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L105
	} else {
		goto L154
	}
L109:
	;
	m.G0 = v320 + int32(16)
	goto L99
L110:
	;
	if v352&int32(1) == int32(0) {
		v385 = l4
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v324 != 0 {
		goto L118
	} else {
		goto L119
	}
L112:
	;
	v362 = v348 - v356
	if v362 <= int32(0) {
		v385 = l4
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v365 = v362
	v369 = l4
	goto L114
L114:
	;
	v374 = F_pg_mblen_cstr(m, v369)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L105
	} else {
		goto L116
	}
L115:
	;
	v385 = v376
	goto L111
L116:
	;
	v376 = v374 + v369
	v377 = int32(1)
	if base.Ui32(v377) < base.Ui32(v365) {
		v365 = v365 - v377
		v369 = v376
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v392 = v324
	v394 = v385
	goto L121
L119:
	;
	goto L120
L120:
	;
	v531 = int32(1)
	goto L109
L121:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	switch v399&int32(3) - int32(1) {
	case 0:
		goto L125
	case 1:
		goto L124
	default:
		goto L108
	}
L122:
	;
	goto L120
L123:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v510 = F_pg_mblen_cstr(m, v394)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L105
	} else {
		goto L152
	}
L124:
	;
	v452 = F_pg_mblen_cstr(m, v394)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L105
	} else {
		goto L139
	}
L125:
	;
	v404 = F_pg_mblen_cstr(m, v394)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L105
	} else {
		goto L126
	}
L126:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+8)))
	if v406 == int32(0) {
		v531 = v7
		goto L109
	} else {
		goto L127
	}
L127:
	;
	v416 = v392 + int32(8)
	goto L128
L128:
	;
	v420 = F_pg_mblen_cstr(m, v416)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L105
	} else {
		goto L130
	}
L129:
	;
	v531 = v7
	goto L109
L130:
	;
	if v404 == v420 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v423 = v404
	goto L134
L132:
	;
	goto L133
L133:
	;
	v450 = v416 + v420
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	if v451 != 0 {
		v416 = v450
		goto L128
	} else {
		goto L138
	}
L134:
	;
	if v423 == int32(0) {
		goto L123
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	v435 = v423 - int32(1)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v435))))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v394))))
	if v437 == v439 {
		v423 = v435
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	goto L129
L139:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+8)))
	if v454 == int32(0) {
		goto L123
	} else {
		goto L140
	}
L140:
	;
	v464 = v392 + int32(8)
	goto L141
L141:
	;
	v468 = F_pg_mblen_cstr(m, v464)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L105
	} else {
		goto L143
	}
L142:
	;
	goto L123
L143:
	;
	if v452 == v468 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v471 = v452
	goto L147
L145:
	;
	goto L146
L146:
	;
	v498 = v464 + v468
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if v499 != 0 {
		v464 = v498
		goto L141
	} else {
		goto L151
	}
L147:
	;
	if v471 == int32(0) {
		v531 = v7
		goto L109
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	v483 = v471 - int32(1)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464+v483))))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v394))))
	if v485 == v487 {
		v471 = v483
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	goto L142
L152:
	;
	if v509 != 0 {
		v392 = v509
		v394 = v510 + v394
		goto L121
	} else {
		goto L153
	}
L153:
	;
	goto L122
L154:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v539 & int32(3)
	F_errmsg_internal(m, int32(_a_F_CheckAffix_2), v320)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L105
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_CheckAffix_3), int32(245), int32(_a_F_CheckAffix_4))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L105
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	v585 = l4
	goto L1
L158:
	;
	v560 = F_pg_mb2wchar_with_len(m, l4, v558, v553)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L105
	} else {
		goto L159
	}
L159:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v563 = int32(0)
	v566 = F_pg_regexec(m, v562, v558, v560, v563, v563, v563)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L105
	} else {
		goto L160
	}
L160:
	;
	F_pfree(m, v558)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L105
	} else {
		goto L161
	}
L161:
	;
	if v566 == int32(0) {
		v585 = l4
		goto L1
	} else {
		goto L162
	}
L162:
	;
	goto L2
}
func F_CheckDim_3(m *base.Module, l0 int32) {
	var v10 int32
	_ = v10
	Fn13823(m, l0, int32(105), int32(_a_F_CheckDim_3_0), int32(_a_F_CheckDim_3_1), int32(_a_F_CheckDim_3_2), int32(100), int32(_a_F_CheckDim_3_3), int32(_a_F_CheckDim_3_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_CloneForeignKeyConstraints(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v603 int32
	_ = v603
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1340 int32
	_ = v1340
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(912)
	m.G0 = v24
	v26 = F_RelationGetFKeyList(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L5
	} else {
		goto L215
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L5
	} else {
		goto L211
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L5
	} else {
		goto L208
	}
L4:
	;
	v627 = F_table_open(m, int32(2606), int32(2))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L5
	} else {
		goto L91
	}
L5:
	;
	return
L6:
	;
	if v26 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v30 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = v4
	v42 = v4
	goto L9
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v37<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v59 == v60 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	if v63 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v63 = F_lappend_oid(m, v42, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v66 = v37 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v66 < v67 {
		v37 = v66
		v42 = v63
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
	if v72 == int32(102) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v77 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v82 = F_build_attrmap_by_name(m, v79, v80, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v84 = F_RelationGetFKeyList(m, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v86 = F_copyObjectImpl(m, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if int32(0) < v88 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v101 = v4
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_relation_close(m, v77, int32(3))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L5
	} else {
		goto L90
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v101<<(uint(int32(2))%32))))
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+172)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v117
	v122 = F_SearchSysCache1(m, int32(19), v116)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v122 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+22)))
	v128 = v126 + v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+92))
	v130 = int32(0)
	if v63 == v130 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v577 = v101 + int32(1)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v577 < v578 {
		v101 = v577
		goto L23
	} else {
		goto L89
	}
L28:
	;
	if v168 != 0 {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v168 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v136 <= int32(0) {
		v162 = v130
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v168 = v162
	goto L28
L33:
	;
	v139 = int32(0)
	if v139 < v136 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v142 = v136
	goto L36
L35:
	;
	v142 = v139
	goto L36
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v145 = int32(0)
	goto L37
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143+v145<<(uint(int32(2))%32))))
	v154 = base.B2i32(v153 == v129)
	if v153 == v129 {
		v162 = v154
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v162 = v154
	goto L32
L39:
	;
	v156 = v145 + int32(1)
	if v156 != v142 {
		v145 = v156
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v128)+96))
	v173 = F_table_open(m, v171, int32(6))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	goto L27
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+119)))
	if v176 == int32(112) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
	v182 = F_find_all_inheritors(m, v179, int32(6), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_DeconstructFkConstraintRow(m, v122, v24+int32(864), v24+int32(768), v24+int32(624), v24+int32(432), v24+int32(304), v24+int32(176), v24+int32(764), v24+int32(560))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v202 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	if v330 != 0 {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v206 = int32(0)
	if v202 != int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v218 = v206
	v219 = int32(0)
	goto L56
L54:
	;
	v276 = v206
	goto L55
L55:
	;
	v293 = int32(1)
	v294 = v276 << (uint(v293) % 32)
	v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(768)+v294))))
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205+v301<<(uint(v293)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v294+(v24+int32(688))))) = uint16(v307)
	goto L51
L56:
	;
	v235 = int32(1)
	v236 = v218 << (uint(v235) % 32)
	v238 = v24 + int32(688)
	v241 = v24 + int32(768)
	v243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241+v236))))
	v247 = int32(2)
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205+v243<<(uint(v235)%32)-v247))))
	*(*uint16)(unsafe.Add(mBase, uint32(v236+v238))) = uint16(v249)
	v252 = v236 | v247
	v257 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241+v252))))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205+v257<<(uint(v235)%32)-v247))))
	*(*uint16)(unsafe.Add(mBase, uint32(v238+v252))) = uint16(v263)
	v266 = v218 + v247
	v268 = v219 + v247
	if v268 != v202&int32(2147483646) {
		v218 = v266
		v219 = v268
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v202&int32(1) == int32(0) {
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v276 = v266
	goto L55
L60:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v128)+96))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v128)+80))
	F_GetForeignKeyCheckTriggers(m, v77, v331, v332, v333, v24+int32(172), v24+int32(92))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v86 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v416 = F_palloc0(m, int32(108))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L76
	}
L65:
	;
	v342 = int32(0)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v343 <= v342 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v352 = v342
	goto L67
L67:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v369+v352<<(uint(int32(2))%32))))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	v381 = F_tryAttachPartitionForeignKey(m, l0, v373, l2, v116, v374, v24+int32(688), v24+int32(624), v24+int32(432), v347, v346, v77)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L69
	}
L68:
	;
	F_relation_close(m, v173, int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L74
	}
L69:
	;
	if v381 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v386 = v352 + int32(1)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v386 < v387 {
		v352 = v386
		goto L67
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	goto L68
L73:
	;
	goto L64
L74:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	goto L27
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v416))) = int64(438086664353)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+12)) = uint8(v420)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v416)+104)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+13)) = uint8(v422)
	v426 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+80)) = v426
	*(*int32)(unsafe.Add(mBase, uint32(v416)+72)) = v426
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+86)) = uint8(v430)
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+87)) = uint8(v432)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v416)+100)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v416)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+88)) = uint8(v434)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+15)) = uint8(v426)
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+14)) = uint8(v440)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)) = uint8(v444)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v426 < v446 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v416)+76))
	v455 = int32(0)
	v459 = v449
	goto L80
L78:
	;
	v504 = v446
	goto L79
L79:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v128)+88))
	v525 = v24 + int32(624)
	v527 = v24 + int32(688)
	v529 = v24 + int32(432)
	v531 = v24 + int32(304)
	v533 = v24 + int32(176)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v536 = v24 + int32(560)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+107)))
	F_addFkConstraint(m, v24+int32(96), int32(1), v128+int32(4), v416, l2, v173, v523, v116, v504, v525, v527, v529, v531, v533, v534, v536, int32(0), v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L5
	} else {
		goto L85
	}
L80:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(688)+v455<<(uint(int32(1))%32)))))
	v488 = F_makeString(m, v472+v473<<(uint(int32(4))%32)+v482*int32(100)-int32(76))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L5
	} else {
		goto L82
	}
L81:
	;
	v504 = v495
	goto L79
L82:
	;
	v490 = F_lappend(m, v459, v488)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+76)) = v490
	v494 = v455 + int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v494 < v495 {
		v455 = v494
		v459 = v490
		goto L80
	} else {
		goto L84
	}
L84:
	;
	goto L81
L85:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	F_addFkRecurseReferencing(m, l0, v416, l2, v173, v523, v541, v544, v525, v527, v529, v531, v533, v545, v536, int32(0), int32(8), v548, v549, v538)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_relation_close(m, v173, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	goto L27
L89:
	;
	goto L24
L90:
	;
	goto L4
L91:
	;
	v630 = v24 + int32(768)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v630, int32(13), int32(3), int32(184), v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	F_ScanKeyInit(m, v24+int32(816), int32(4), int32(3), int32(61), int32(102))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v645 = int32(0)
	v650 = F_systable_beginscan(m, v627, v645, int32(1), v645, int32(2), v630)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v652 = F_systable_getnext(m, v650)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	if v652 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v658 = v652
	v659 = v645
	goto L99
L97:
	;
	v688 = v645
	goto L98
L98:
	;
	F_systable_endscan(m, v650)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L104
	}
L99:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v658)+16))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+22)))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v675+v676)))
	v679 = F_lappend_oid(m, v659, v678)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L5
	} else {
		goto L101
	}
L100:
	;
	v688 = v679
	goto L98
L101:
	;
	v681 = F_systable_getnext(m, v650)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	if v681 != 0 {
		v658 = v681
		v659 = v679
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	F_relation_close(m, v627, int32(2))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v711 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v716 = F_build_attrmap_by_name(m, v713, v714, int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v688 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_relation_close(m, v711, int32(3))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L5
	} else {
		goto L207
	}
L109:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	if v720 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v738 = int32(0)
	goto L111
L111:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v688)+12))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746+v738<<(uint(int32(2))%32))))
	v751 = F_SearchSysCache1(m, int32(19), v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L115
	}
L112:
	;
	goto L108
L113:
	;
	F_ReleaseCatCache(m, v751)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L5
	} else {
		goto L205
	}
L114:
	;
	v1261 = int32(0)
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v950)+8))
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v1265 = v24 + int32(624)
	v1267 = v24 + int32(688)
	v1269 = v24 + int32(432)
	v1271 = v24 + int32(304)
	v1273 = v24 + int32(176)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v1276 = v24 + int32(96)
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+107)))
	F_addFkConstraint(m, v24+int32(864), v1261, v1262, v950, v798, l2, v1055, v750, v1263, v1265, v1267, v1269, v1271, v1273, v1274, v1276, v1261, v1278)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L5
	} else {
		goto L202
	}
L115:
	;
	if v751 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v751)+16))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+22)))
	v755 = v753 + v754
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)+92))
	v757 = int32(0)
	if v688 == v757 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	goto L118
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L199
	}
L119:
	;
	if v795 != 0 {
		goto L113
	} else {
		goto L132
	}
L120:
	;
	v795 = int32(0)
	goto L119
L121:
	;
	goto L122
L122:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	if v763 <= int32(0) {
		v789 = v757
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v795 = v789
	goto L119
L124:
	;
	v766 = int32(0)
	if v766 < v763 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v769 = v763
	goto L127
L126:
	;
	v769 = v766
	goto L127
L127:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v688)+12))
	v772 = int32(0)
	goto L128
L128:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v770+v772<<(uint(int32(2))%32))))
	v781 = base.B2i32(v780 == v756)
	if v780 == v756 {
		v789 = v781
		goto L123
	} else {
		goto L130
	}
L129:
	;
	v789 = v781
	goto L123
L130:
	;
	v783 = v772 + int32(1)
	if v783 != v769 {
		v772 = v783
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v755)+80))
	v798 = F_table_open(m, v796, int32(6))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v755)+88))
	F_DeconstructFkConstraintRow(m, v751, v24+int32(764), v24+int32(688), v24+int32(560), v24+int32(432), v24+int32(304), v24+int32(176), v24+int32(172), v24+int32(96))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v819 <= int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v950 = F_palloc0(m, int32(108))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L5
	} else {
		goto L144
	}
L136:
	;
	v822 = int32(0)
	if v819 != int32(1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v834 = v822
	v837 = int32(0)
	goto L140
L138:
	;
	v894 = v822
	goto L139
L139:
	;
	v911 = int32(1)
	v912 = v894 << (uint(v911) % 32)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	v920 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(560)+v912))))
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916+v920<<(uint(v911)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v912+(v24+int32(624))))) = uint16(v926)
	goto L135
L140:
	;
	v851 = int32(1)
	v852 = v834 << (uint(v851) % 32)
	v854 = v24 + int32(624)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	v858 = v24 + int32(560)
	v860 = int32(*(*int16)(unsafe.Add(mBase, uint32(v858+v852))))
	v864 = int32(2)
	v866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v856+v860<<(uint(v851)%32)-v864))))
	*(*uint16)(unsafe.Add(mBase, uint32(v852+v854))) = uint16(v866)
	v869 = v852 | v864
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	v875 = int32(*(*int16)(unsafe.Add(mBase, uint32(v858+v869))))
	v881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v871+v875<<(uint(v851)%32)-v864))))
	*(*uint16)(unsafe.Add(mBase, uint32(v854+v869))) = uint16(v881)
	v884 = v834 + v864
	v886 = v837 + v864
	if v886 != v819&int32(2147483646) {
		v834 = v884
		v837 = v886
		goto L140
	} else {
		goto L142
	}
L141:
	;
	if v819&int32(1) == int32(0) {
		goto L135
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v894 = v884
	goto L139
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v950)+8)) = v755 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v950))) = int64(438086664353)
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+12)) = uint8(v957)
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v950)+104)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+13)) = uint8(v959)
	v963 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v950)+80)) = v963
	*(*int32)(unsafe.Add(mBase, uint32(v950)+72)) = v963
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+86)) = uint8(v967)
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+87)) = uint8(v969)
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v950)+100)) = v963
	*(*int64)(unsafe.Add(mBase, uint32(v950)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+88)) = uint8(v971)
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+15)) = uint8(v963)
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+14)) = uint8(v977)
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v950)+16)) = uint8(v981)
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v963 < v983 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v950)+76))
	v992 = int32(0)
	v996 = v986
	goto L148
L146:
	;
	goto L147
L147:
	;
	v1055 = F_index_get_partition(m, l2, v800)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L153
	}
L148:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v798)+52))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1009)))
	v1019 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(688)+v992<<(uint(int32(1))%32)))))
	v1025 = F_makeString(m, v1009+v1010<<(uint(int32(4))%32)+v1019*int32(100)-int32(76))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L5
	} else {
		goto L150
	}
L149:
	;
	goto L147
L150:
	;
	v1027 = F_lappend(m, v996, v1025)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v950)+76)) = v1027
	v1031 = v992 + int32(1)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v1031 < v1032 {
		v992 = v1031
		v996 = v1027
		goto L148
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	if v1055 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1057 = int32(0)
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+75)))
	if v1059 != int32(1) {
		v1242 = v1057
		v1247 = v1057
		goto L114
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L5
	} else {
		goto L196
	}
L157:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v755)+80))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v755)+96))
	v1065 = v24 + int32(864)
	F_ScanKeyInit(m, v1065, int32(11), int32(3), int32(184), v750)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L5
	} else {
		goto L158
	}
L158:
	;
	v1071 = int32(0)
	v1073 = int32(1)
	v1076 = F_systable_beginscan(m, v711, int32(2699), v1073, v1071, v1073, v1065)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L5
	} else {
		goto L161
	}
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L5
	} else {
		goto L193
	}
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L5
	} else {
		goto L190
	}
L161:
	;
	v1078 = F_systable_getnext(m, v1076)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	if v1078 == int32(0) {
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v1086 = v1078
	v1091 = v1057
	v1092 = v1071
	goto L164
L164:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+16))
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103)+22)))
	v1105 = v1103 + v1104
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+84))
	if v1106 != v1062 {
		v1137 = v1091
		v1138 = v1092
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L160
L166:
	;
	v1153 = F_systable_getnext(m, v1076)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L5
	} else {
		goto L188
	}
L167:
	;
	F_systable_endscan(m, v1076)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L5
	} else {
		goto L187
	}
L168:
	;
	v1140 = F_systable_getnext(m, v1076)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L5
	} else {
		goto L183
	}
L169:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+4))
	if v1108 != v1063 {
		v1137 = v1091
		v1138 = v1092
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+76))
	v1112 = v1110 - int32(1644)
	if base.Ui32(v1112) <= base.Ui32(int32(11)) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v1119 != int32(1) {
		v1137 = v1091
		v1138 = v1092
		goto L168
	} else {
		goto L175
	}
L172:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1112<<(uint(int32(2))%32))+uint32(_c_F_CloneForeignKeyConstraints[0])))
	v1119 = v1117
	goto L174
L173:
	;
	v1119 = int32(0)
	goto L174
L174:
	;
	goto L171
L175:
	;
	v1122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+80)))
	if v1122&int32(8) != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	if v1131 == int32(0) {
		goto L166
	} else {
		goto L181
	}
L177:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1105)))
	v1131 = v1125
	v1132 = v1092
	goto L176
L178:
	;
	goto L179
L179:
	;
	if v1122&int32(16) == int32(0) {
		v1131 = v1091
		v1132 = v1092
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1105)))
	v1131 = v1091
	v1132 = v1130
	goto L176
L181:
	;
	if v1132 != 0 {
		v1146 = v1132
		v1147 = v1131
		goto L167
	} else {
		goto L182
	}
L182:
	;
	v1137 = v1131
	v1138 = int32(0)
	goto L168
L183:
	;
	if v1140 != 0 {
		v1086 = v1140
		v1091 = v1137
		v1092 = v1138
		goto L164
	} else {
		goto L184
	}
L184:
	;
	if v1137 == int32(0) {
		goto L160
	} else {
		goto L185
	}
L185:
	;
	if v1138 == int32(0) {
		goto L159
	} else {
		goto L186
	}
L186:
	;
	v1146 = v1138
	v1147 = v1137
	goto L167
L187:
	;
	v1242 = v1146
	v1247 = v1147
	goto L114
L188:
	;
	if v1153 != 0 {
		v1086 = v1153
		v1091 = int32(0)
		v1092 = v1132
		goto L164
	} else {
		goto L189
	}
L189:
	;
	goto L165
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v750
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_0), v24+int32(32))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L5
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_2), int32(_a_F_CloneForeignKeyConstraints_3))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v750
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_4), v24+int32(48))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_5), int32(_a_F_CloneForeignKeyConstraints_3))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v800
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v1210 + int32(4)
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_6), v24+int32(16))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_7), int32(_a_F_CloneForeignKeyConstraints_8))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v750
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_9), v24)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_10), int32(_a_F_CloneForeignKeyConstraints_8))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v24)+868))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+107)))
	F_addFkRecurseReferenced(m, v950, v798, l2, v1055, v1281, v1282, v1265, v1267, v1269, v1271, v1273, v1283, v1276, v1247, v1242, v1284)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	F_relation_close(m, v798, int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	goto L113
L205:
	;
	v1314 = v738 + int32(1)
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	if v1314 < v1315 {
		v738 = v1314
		goto L111
	} else {
		goto L206
	}
L206:
	;
	goto L112
L207:
	;
	m.G0 = v24 + int32(912)
	return
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v116
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_9), v24+int32(80))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_11), int32(_a_F_CloneForeignKeyConstraints_12))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	F_errmsg(m, int32(_a_F_CloneForeignKeyConstraints_13), int32(0))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_14), int32(_a_F_CloneForeignKeyConstraints_12))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L215:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v1384 = F_get_constraint_name(m, v1383)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v1384
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1382 + int32(4)
	F_errmsg(m, int32(_a_F_CloneForeignKeyConstraints_15), v24-int32(-64))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_16), int32(_a_F_CloneForeignKeyConstraints_12))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CloseTransientFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CloseTransientFile[0]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CloseTransientFile[1]))
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v40 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v19 = v12 + v14*int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != int32(3) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v14 {
		v14 = v14 - int32(1)
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v23 != l0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = F_FreeDesc(m, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v25
L11:
	;
	goto L5
L12:
	;
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CloseTransientFile_0), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_pgaio_closing_fd(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(_a_F_CloseTransientFile_1), int32(2892), int32(_a_F_CloseTransientFile_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v53 = F_close(m, l0)
	mBase = m.M
	return v53
}
func F_ConditionVariableBroadcast(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[1]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v69 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	F_s_lock(m, v11, int32(_a_F_ConditionVariableBroadcast_0), int32(238), int32(_a_F_ConditionVariableBroadcast_1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[0]))
	v27 = v22 + v24*int32(640)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	if v28 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[1])) = v60
	goto L3
L10:
	;
	if v45 == int32(-1) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v28*int32(640))+84)) = v39
	v44 = v28
	v45 = v39
	goto L10
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	if v28 != int32(-1) {
		v39 = v34
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v39 = v31
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	v44 = v38
	v45 = v34
	goto L10
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+84)) = int64(0)
	goto L9
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v44
	goto L17
L19:
	;
	goto L20
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v51+v45*int32(640))+88)) = v44
	goto L17
L21:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableBroadcast_0), int32(317), int32(_a_F_ConditionVariableBroadcast_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v77 != int32(-1) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	F_SetLatch(m, v85+int32(20))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L70
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
	F_SetLatch(m, v85+int32(20))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L42
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+88)) = v118
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v118*int32(640))+84)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v117)+84)) = int32(-1)
	goto L27
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v85 = v82 + v77*int32(640)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+88))
	if v87 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
L32:
	;
	if v86 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v86
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+88))
	v96 = v91
	goto L32
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+v87*int32(640))+84)) = v86
	v96 = v87
	goto L32
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v85)+84)) = int64(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v109 == int32(-1) {
		goto L26
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96
	goto L36
L38:
	;
	goto L39
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v102+v86*int32(640))+88)) = v96
	goto L36
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v117 = v114 + v9*int32(640)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v118 != int32(-1) {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117)+84)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L27
L42:
	;
	goto L43
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v152 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L25
L45:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableBroadcast_0), int32(351), int32(_a_F_ConditionVariableBroadcast_2))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v160 == int32(-1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = v199 + v9*int32(640)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+88))
	if v201 != 0 {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v194 = int32(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v169 = v166 + v160*int32(640)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+84))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)+88))
	if v171 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v170 == int32(-1) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+88))
	v180 = v175
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166+v171*int32(640))+84)) = v170
	v180 = v171
	goto L53
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169)+84)) = int64(0)
	v194 = v169
	goto L49
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v180
	goto L57
L59:
	;
	goto L60
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int32)(unsafe.Add(mBase, uint32(v186+v170*int32(640))+88)) = v180
	goto L57
L61:
	;
	v206 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v206
	if v194 == v206 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v205 = int32(1)
	goto L61
L63:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+84))
	if v202 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v205 = int32(0)
	goto L61
L65:
	;
	if v205 != 0 {
		goto L43
	} else {
		goto L69
	}
L66:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[3]))
	if v194 == v211 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	F_SetLatch(m, v194+int32(20))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	goto L44
L70:
	;
	goto L25
}
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1]))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
		if v12 != 0 {
			F_s_lock(m, v11, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(238), int32(_a_F_ConditionVariablePrepareToSleep_1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
				v27 = v22 + v24*int32(640)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
				if v28 == int32(0) {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
					if v31 == int32(0) {
					} else {
						v39 = v31
						*(*int32)(unsafe.Add(mBase, uint32(v22+v28*int32(640))+84)) = v39
						v44 = v28
						v45 = v39
						if v45 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v44
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							*(*int32)(unsafe.Add(mBase, uint32(v51+v45*int32(640))+88)) = v44
						}
						*(*int64)(unsafe.Add(mBase, uint32(v27)+84)) = int64(0)
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
					if v28 != int32(-1) {
						v39 = v34
						*(*int32)(unsafe.Add(mBase, uint32(v22+v28*int32(640))+84)) = v39
						v44 = v28
						v45 = v39
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v34
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
						v44 = v38
						v45 = v34
					}
					if v45 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v44
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						*(*int32)(unsafe.Add(mBase, uint32(v51+v45*int32(640))+88)) = v44
					}
					*(*int64)(unsafe.Add(mBase, uint32(v27)+84)) = int64(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
				if v68 != 0 {
					F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(75), int32(_a_F_ConditionVariablePrepareToSleep_2))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
						v81 = v78 + v9*int32(640)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v82 == int32(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(v81)+84)) = int64(-1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v81)+88)) = v82
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
							*(*int32)(unsafe.Add(mBase, uint32(v91+v82*int32(640))+84)) = v9
							*(*int32)(unsafe.Add(mBase, uint32(v81)+84)) = int32(-1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
						return
					}
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
					v81 = v78 + v9*int32(640)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v82 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v81)+84)) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v81)+88)) = v82
						v90 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						*(*int32)(unsafe.Add(mBase, uint32(v91+v82*int32(640))+84)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v81)+84)) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
			v27 = v22 + v24*int32(640)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
			if v28 == int32(0) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
				if v31 == int32(0) {
				} else {
					v39 = v31
					*(*int32)(unsafe.Add(mBase, uint32(v22+v28*int32(640))+84)) = v39
					v44 = v28
					v45 = v39
					if v45 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v44
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						*(*int32)(unsafe.Add(mBase, uint32(v51+v45*int32(640))+88)) = v44
					}
					*(*int64)(unsafe.Add(mBase, uint32(v27)+84)) = int64(0)
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
				if v28 != int32(-1) {
					v39 = v34
					*(*int32)(unsafe.Add(mBase, uint32(v22+v28*int32(640))+84)) = v39
					v44 = v28
					v45 = v39
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v34
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
					v44 = v38
					v45 = v34
				}
				if v45 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v44
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(v51+v45*int32(640))+88)) = v44
				}
				*(*int64)(unsafe.Add(mBase, uint32(v27)+84)) = int64(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
			if v68 != 0 {
				F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(75), int32(_a_F_ConditionVariablePrepareToSleep_2))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
					v81 = v78 + v9*int32(640)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v82 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v81)+84)) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v81)+88)) = v82
						v90 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						*(*int32)(unsafe.Add(mBase, uint32(v91+v82*int32(640))+84)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v81)+84)) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					return
				}
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
				v81 = v78 + v9*int32(640)
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v82 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v81)+84)) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v81)+88)) = v82
					v90 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
					*(*int32)(unsafe.Add(mBase, uint32(v91+v82*int32(640))+84)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v81)+84)) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
		if v68 != 0 {
			F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(75), int32(_a_F_ConditionVariablePrepareToSleep_2))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
				v81 = v78 + v9*int32(640)
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v82 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v81)+84)) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v81)+88)) = v82
					v90 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
					*(*int32)(unsafe.Add(mBase, uint32(v91+v82*int32(640))+84)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v81)+84)) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				return
			}
		} else {
			v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
			v81 = v78 + v9*int32(640)
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v82 == int32(-1) {
				*(*int64)(unsafe.Add(mBase, uint32(v81)+84)) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v81)+88)) = v82
				v90 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				*(*int32)(unsafe.Add(mBase, uint32(v91+v82*int32(640))+84)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v81)+84)) = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
			return
		}
	}
}
func F_ConditionalLockBufferForCleanup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 < v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v176
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l0^int32(-1))<<(uint(int32(2))%32))))
	v176 = base.B2i32(v19 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[1]))
	if l0 == v24 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 != int32(1) {
		v176 = v2
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[2]))
	if l0 == v28 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[3]))
	if l0 == v32 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_2)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[4]))
	if l0 == v36 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_3)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[5]))
	if l0 == v40 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_4)
	goto L5
L19:
	;
	goto L20
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[6]))
	if l0 == v44 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_5)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[7]))
	if l0 == v48 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_6)
	goto L5
L25:
	;
	goto L26
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[8]))
	if l0 == v52 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_7)
	goto L5
L28:
	;
	goto L29
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[9]))
	if v56 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[10]))
	v63 = int32(0)
	v65 = F_hash_search(m, v60, v8+int32(8), v63, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	if v65 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v71 = v65
	goto L5
L34:
	;
	v76 = l0 << (uint(int32(6)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[11]))
	v83 = F_LWLockConditionalAcquire(m, v76+v78-int32(16), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v83 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_ConditionalLockBufferForCleanup_8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_ConditionalLockBufferForCleanup_9)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_ConditionalLockBufferForCleanup_10)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v101 = v76 + v88 - int32(40)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = int32(_a_F_ConditionalLockBufferForCleanup_11)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102 | v103
	if v102&v103 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	goto L40
L38:
	;
	v124 = v102
	goto L39
L39:
	;
	v132 = int32(_a_F_ConditionalLockBufferForCleanup_12)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[12]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v135 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v124 = v117
	goto L39
L42:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v118 = int32(_a_F_ConditionalLockBufferForCleanup_11)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v117 | v118
	if v117&v118 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v124&int32(_a_F_ConditionalLockBufferForCleanup_13) == int32(1) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[12])) = v150
	goto L45
L47:
	;
	if int32(999) < v133 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v133 < int32(11) {
		goto L45
	} else {
		goto L54
	}
L50:
	;
	v140 = int32(900)
	if v140 <= v133 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v143 = v140
	goto L53
L52:
	;
	v143 = v133
	goto L53
L53:
	;
	v150 = v143 + int32(100)
	goto L46
L54:
	;
	v150 = v133 - int32(1)
	goto L46
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v124 & int32(-4456447)
	v176 = int32(1)
	goto L1
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v124 & int32(-4194305)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[11]))
	F_LWLockRelease(m, v164+l0<<(uint(int32(6))%32)-int32(16))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	v176 = int32(0)
	goto L1
}
func F_CountChildren(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CountChildren[0]))
	if base.B2i32(v12 == v2)|base.B2i32(v12 == int32(_a_F_CountChildren_0)) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = v12
	v33 = v2
	goto L4
L2:
	;
	v103 = v2
	goto L3
L3:
	;
	m.G0 = v9 + int32(16)
	return v103
L4:
	;
	if int32(base.Ui32(l0&int32(64))>>(uint(int32(6))%32))^int32(base.Ui32(l0&int32(2))>>(uint(int32(1))%32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v103 = v95
	goto L3
L6:
	;
	v59 = v30 - int32(12)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if int32(base.Ui32(l0)>>(uint(v60)%32))&int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v38 = v30 - int32(12)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30-int32(16))))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CountChildren[1]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32))+44))
	goto L9
L9:
	;
	if base.B2i32(v50 == int32(3)) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(6)
	goto L6
L11:
	;
	v66 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v95 = v33
	goto L13
L13:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v96 != int32(_a_F_CountChildren_0) {
		v30 = v96
		v33 = v95
		goto L4
	} else {
		goto L25
	}
L14:
	;
	return int32(0)
L15:
	;
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if base.Ui32(v70) <= base.Ui32(int32(17)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v95 = v33 + int32(1)
	goto L13
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v30-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
	F_errmsg_internal(m, int32(_a_F_CountChildren_1), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L23
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_CountChildren[2])))
	v77 = v75
	goto L22
L21:
	;
	v77 = int32(_a_F_CountChildren_2)
	goto L22
L22:
	;
	goto L19
L23:
	;
	F_errfinish(m, int32(_a_F_CountChildren_3), int32(3942), int32(_a_F_CountChildren_4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	goto L5
}
func F_CountDBBackends(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[1]))
	v18 = F_LWLockAcquire(m, v14+int32(512), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v22 <= int32(0) {
			v108 = v2
		} else {
			v26 = v12 + int32(36)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[2]))
			v29 = int32(0)
			if v22 != int32(1) {
				v38 = v29
				v39 = int32(0)
				v40 = v2
				for {
					v49 = v26 + v38<<(uint(int32(2))%32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v53 = v28 + v50*int32(640)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
					if v54 == int32(0) {
						v61 = v40
					} else {
						if l0 != 0 {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
							if v57 != l0 {
								v61 = v40
							} else {
								v61 = v40 + int32(1)
							}
						} else {
							v61 = v40 + int32(1)
						}
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v65 = v28 + v62*int32(640)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+44))
					if v66 == int32(0) {
						v73 = v61
					} else {
						if l0 != 0 {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+60))
							if v69 != l0 {
								v73 = v61
							} else {
								v73 = v61 + int32(1)
							}
						} else {
							v73 = v61 + int32(1)
						}
					}
					v74 = int32(2)
					v75 = v38 + v74
					v77 = v39 + v74
					if v77 != v22&int32(2147483646) {
						v38 = v75
						v39 = v77
						v40 = v73
						continue
					} else {
						break
					}
					break
				}
				if v22&int32(1) == int32(0) {
					v108 = v73
				} else {
					v82 = v75
					v84 = v73
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v26+v82<<(uint(int32(2))%32))))
					v97 = v28 + v94*int32(640)
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
					if v98 == int32(0) {
						v108 = v84
					} else {
						if l0 != 0 {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+60))
							if v101 != l0 {
								v108 = v84
							} else {
								v108 = v84 + int32(1)
							}
						} else {
							v108 = v84 + int32(1)
						}
					}
				}
			} else {
				v82 = v29
				v84 = v2
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v26+v82<<(uint(int32(2))%32))))
				v97 = v28 + v94*int32(640)
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
				if v98 == int32(0) {
					v108 = v84
				} else {
					if l0 != 0 {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+60))
						if v101 != l0 {
							v108 = v84
						} else {
							v108 = v84 + int32(1)
						}
					} else {
						v108 = v84 + int32(1)
					}
				}
			}
		}
		v116 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[1]))
		F_LWLockRelease(m, v116+int32(512))
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return int32(0)
		} else {
			return v108
		}
	}
}
func F_CreateDecodingContext(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v1 = l0
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[0]))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L23
	} else {
		goto L65
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L23
	} else {
		goto L61
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L23
	} else {
		goto L57
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L23
	} else {
		goto L54
	}
L7:
	;
	v23 = v18 + int32(24)
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[1]))
	if v19 != v27 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[2])))
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	if v1 == int64(0) {
		v83 = v51
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v41 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	v39 = base.B2i32(v37 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[2])) = uint8(v39)
	v41 = v39
	goto L16
L15:
	;
	v41 = int32(0)
	goto L16
L16:
	;
	goto L13
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+201)))
	if v44 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[6])))
	if v48 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v84 = int32(0)
	v87 = F_StartupDecodingContext(m, l1, v83, v84, v84, l2, v84, l3, l4, l5, l6)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L23
	} else {
		goto L30
	}
L21:
	;
	if base.Ui64(v51) <= base.Ui64(v1) {
		v83 = v1
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v57 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)) = uint32(v61)
	v63 = int64(32)
	v64 = int64(base.Ui64(v61) >> (uint(v63) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v64)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v1)
	v68 = int64(base.Ui64(v1) >> (uint(v63) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+32)) = uint32(v68)
	F_errmsg_internal(m, int32(_a_F_CreateDecodingContext_9), v15+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v83 = v81
	goto L20
L28:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(574), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v89 = int32(_a_F_CreateDecodingContext_4)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4])) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	if v94 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(_a_F_CreateDecodingContext_5)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(993)
	v102 = int32(_a_F_CreateDecodingContext_6)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5])) = v15 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v15 + int32(80)
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+164)) = uint8(v112)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+147)) = uint8(v112)
	m.T0[v94].(func(*base.Module, int32, int32, int32))(m, v87, v87+int32(108), v112)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4])) = v90
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v127 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5])) = v122
	goto L33
L35:
	;
	v130 = int32(1)
	goto L37
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+146)))
	v130 = v129
	goto L37
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+145)))
	v132 = v130 & v131
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+145)) = uint8(v132)
	if v132 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+116)) = uint8(v157)
	v161 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L23
	} else {
		goto L47
	}
L39:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v136 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1)
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_s_lock(m, v18, int32(_a_F_CreateDecodingContext_1), int32(601), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L23
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v83
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)) = uint8(v146)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v83
	goto L38
L47:
	;
	if v161 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_7), v15+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L23
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v15 + int32(96)
	return v87
L51:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v170)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v169)
	v173 = int64(32)
	v174 = int64(base.Ui64(v170) >> (uint(v173) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v174)
	v177 = int64(base.Ui64(v169) >> (uint(v173) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v177)
	F_errdetail(m, int32(_a_F_CreateDecodingContext_8), v15)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(617), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	F_errmsg_internal(m, int32(_a_F_CreateDecodingContext_10), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(517), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_CreateDecodingContext_0), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(523), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L23
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_3), v15-int32(-64))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(534), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_11), v15+int32(48))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	F_errdetail(m, int32(_a_F_CreateDecodingContext_12), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	F_errhint(m, int32(_a_F_CreateDecodingContext_13), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(547), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateDirAndVersionFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v8 = m.G0
	v10 = v8 - int32(1136)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(_a_F_CreateDirAndVersionFile_0)
	v19 = F_pg_sprintf(m, v10+int32(96), int32(_a_F_CreateDirAndVersionFile_1), v10+int32(80))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[0]))
	v23 = F_mkdir(m, l0, v22)
	mBase = m.M
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L60
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L56
	}
L5:
	;
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(_a_F_CreateDirAndVersionFile_2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l0
	v36 = v10 + int32(112)
	v41 = F_pg_snprintf(m, v36, int32(1024), int32(_a_F_CreateDirAndVersionFile_3), v10+int32(48))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v29 != int32(20) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v44 = F_OpenTransientFile(m, v36, int32(193))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v44 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l3 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v59 = v44
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(167772226)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1])) = int32(0)
	v69 = int32(3)
	v70 = F_write(m, v59, v10+int32(96), v69)
	mBase = m.M
	if v70 != v69 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v51 != int32(20) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v55 = F_OpenTransientFile(m, v36, int32(513))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v55 < int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v59 = v55
	goto L15
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v74 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v99 = int32(_a_F_CreateDirAndVersionFile_4)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	v101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v101
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = int32(167772225)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[3])))
	if v109 != int32(1) {
		v123 = v101
		goto L32
	} else {
		goto L33
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1])) = int32(51)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_5), v10+int32(32))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(507), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
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
	F_fsync_fname(m, l0, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L48
	}
L31:
	;
	if v123 == int32(0) {
		goto L30
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	goto L34
L34:
	;
	v114 = F_fsync(m, v59)
	mBase = m.M
	if v114 != int32(-1) {
		v123 = v114
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v123 = int32(-1)
	goto L32
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v118 == int32(27) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[4])))
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v132 = F_errstart(m, v130, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v130 = int32(21)
	goto L42
L41:
	;
	v130 = int32(23)
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v132 == int32(0) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_8), v10+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(515), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L30
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(0)
	v158 = F_CloseTransientFile(m, v59)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if l3 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v162 = int32(_a_F_CreateDirAndVersionFile_9)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5])) = v164 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	m.G0 = v10 + int32(1136)
	return
L53:
	;
	F_XLogRegisterData(m, v10+int32(88), int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v179 = F_XLogInsert(m, int32(4), int32(16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v181 = int32(_a_F_CreateDirAndVersionFile_9)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5])) = v183 - int32(1)
	goto L52
L56:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l0
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_10), v10-int32(-64))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(478), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_11), v10)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(495), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CteScanNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+132))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	F_tuplestore_select_read_pointer(m, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v17*int32(24))+4)))
		v22 = int32(0)
		if base.B2i32(v21 == v22)|base.B2i32(v7 == int32(1)) == v22 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+136)))
			if v30 != 0 {
				v36 = int32(1)
				v39 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v36), v36, v15)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 != 0 {
						v78 = v15
						return v78
					} else {
						if v7 != int32(1) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							m.T0[v74].(func(*base.Module, int32))(m, v15)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v78 = v15
								return v78
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
							if v45 != 0 {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
								m.T0[v74].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v78 = v15
									return v78
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
								if v47 != 0 {
									F_ExecReScan(m, v46)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
										v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											if v51 != 0 {
												v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
												if v53&int32(2) == int32(0) {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v63)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v51)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
															m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int32(0)
															} else {
																v78 = v15
																return v78
															}
														}
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v59 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
													return int32(0)
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
									v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 != 0 {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
											if v53&int32(2) == int32(0) {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												F_tuplestore_select_read_pointer(m, v9, v63)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_tuplestore_puttupleslot(m, v9, v51)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
														m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int32(0)
														} else {
															v78 = v15
															return v78
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v31 = int32(0)
				v33 = F_tuplestore_advance(m, v9, v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v33 != 0 {
						v36 = int32(1)
						v39 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v36), v36, v15)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 != 0 {
								v78 = v15
								return v78
							} else {
								if v7 != int32(1) {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
									m.T0[v74].(func(*base.Module, int32))(m, v15)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = v15
										return v78
									}
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
									if v45 != 0 {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
										m.T0[v74].(func(*base.Module, int32))(m, v15)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v78 = v15
											return v78
										}
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
										if v47 != 0 {
											F_ExecReScan(m, v46)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
												v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int32(0)
												} else {
													if v51 != 0 {
														v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
														if v53&int32(2) == int32(0) {
															v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
															F_tuplestore_select_read_pointer(m, v9, v63)
															mBase = m.M
															v65 = m.ExcPending
															if v65 != 0 {
																return int32(0)
															} else {
																F_tuplestore_puttupleslot(m, v9, v51)
																mBase = m.M
																v67 = m.ExcPending
																if v67 != 0 {
																	return int32(0)
																} else {
																	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
																	m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
																	mBase = m.M
																	v71 = m.ExcPending
																	if v71 != 0 {
																		return int32(0)
																	} else {
																		v78 = v15
																		return v78
																	}
																}
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															v59 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
															return int32(0)
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v59 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
														return int32(0)
													}
												}
											}
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
											v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												if v51 != 0 {
													v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
													if v53&int32(2) == int32(0) {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
														F_tuplestore_select_read_pointer(m, v9, v63)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															F_tuplestore_puttupleslot(m, v9, v51)
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
																m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
																mBase = m.M
																v71 = m.ExcPending
																if v71 != 0 {
																	return int32(0)
																} else {
																	v78 = v15
																	return v78
																}
															}
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v59 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
														return int32(0)
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v59 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v78 = v31
						return v78
					}
				}
			}
		} else {
			if v21 != 0 {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
				if v45 != 0 {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
					m.T0[v74].(func(*base.Module, int32))(m, v15)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v78 = v15
						return v78
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
					if v47 != 0 {
						F_ExecReScan(m, v46)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
							v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v51 != 0 {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
									if v53&int32(2) == int32(0) {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										F_tuplestore_select_read_pointer(m, v9, v63)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											F_tuplestore_puttupleslot(m, v9, v51)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
												m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													v78 = v15
													return v78
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v59 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
										return int32(0)
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
									return int32(0)
								}
							}
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
						v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 != 0 {
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
								if v53&int32(2) == int32(0) {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									F_tuplestore_select_read_pointer(m, v9, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										F_tuplestore_puttupleslot(m, v9, v51)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
											m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v78 = v15
												return v78
											}
										}
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
									return int32(0)
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
								return int32(0)
							}
						}
					}
				}
			} else {
				v36 = int32(1)
				v39 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v36), v36, v15)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 != 0 {
						v78 = v15
						return v78
					} else {
						if v7 != int32(1) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							m.T0[v74].(func(*base.Module, int32))(m, v15)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v78 = v15
								return v78
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
							if v45 != 0 {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
								m.T0[v74].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v78 = v15
									return v78
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
								if v47 != 0 {
									F_ExecReScan(m, v46)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
										v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											if v51 != 0 {
												v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
												if v53&int32(2) == int32(0) {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v63)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v51)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
															m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int32(0)
															} else {
																v78 = v15
																return v78
															}
														}
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v59 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
													return int32(0)
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
									v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 != 0 {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
											if v53&int32(2) == int32(0) {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												F_tuplestore_select_read_pointer(m, v9, v63)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_tuplestore_puttupleslot(m, v9, v51)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
														m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int32(0)
														} else {
															v78 = v15
															return v78
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
											return int32(0)
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
func F___clock_gettime(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if base.Ui32(int32(4)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F___clock_gettime[0])) = int32(28)
	} else {
		v18 = m.Wasi_snapshot_preview1.Clock_time_get(m, l0, int64(1), v8+int32(24))
		mBase = m.M
		if v18 == int32(0) {
			v25 = int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F___clock_gettime[0])) = v18
			v25 = int32(-1)
		}
		if v25 != 0 {
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			v28 = v8 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
			v31 = int64(1000000000)
			v32 = base.I64_div_u_s(v26, v31)
			*(*int64)(unsafe.Add(mBase, uint32(v28))) = v32
			v36 = v26 - v32*v31
			*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)) = uint32(v36)
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v40
		}
	}
	m.G0 = v8 + int32(32)
	return
}
func F_calc_hist_selectivity_contains(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v79 int32
	_ = v79
	var v82 float64
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v123 float64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 float64
	_ = v149
	var v151 float64
	_ = v151
	var v153 int32
	_ = v153
	var v158 float64
	_ = v158
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 float64
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v236 int32
	_ = v236
	var v250 int32
	_ = v250
	var v259 float64
	_ = v259
	var v263 float64
	_ = v263
	var v265 int32
	_ = v265
	var v271 float64
	_ = v271
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v283 int32
	_ = v283
	var v288 float64
	_ = v288
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 float64
	_ = v302
	var v309 float64
	_ = v309
	var v312 float64
	_ = v312
	var v322 float64
	_ = v322
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v334 float64
	_ = v334
	var v336 int32
	_ = v336
	var v350 int32
	_ = v350
	var v359 float64
	_ = v359
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v369 float64
	_ = v369
	var v376 int32
	_ = v376
	var v379 float64
	_ = v379
	var v381 float64
	_ = v381
	var v382 float64
	_ = v382
	var v384 float64
	_ = v384
	var v388 float64
	_ = v388
	var v392 float64
	_ = v392
	var v402 float64
	_ = v402
	var v422 float64
	_ = v422
	var v434 float64
	_ = v434
	var v439 int32
	_ = v439
	var v442 float64
	_ = v442
	var v443 float64
	_ = v443
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 float64
	_ = v462
	var v463 int32
	_ = v463
	var v466 float64
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 float64
	_ = v473
	var v479 float64
	_ = v479
	var v482 float64
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 float64
	_ = v492
	var v493 float64
	_ = v493
	var v496 int32
	_ = v496
	var v501 float64
	_ = v501
	var v512 float64
	_ = v512
	var v514 float64
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 float64
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 float64
	_ = v571
	var v572 float64
	_ = v572
	var v573 float64
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 float64
	_ = v576
	var v577 float64
	_ = v577
	var v579 int32
	_ = v579
	var v593 int32
	_ = v593
	var v602 float64
	_ = v602
	var v606 float64
	_ = v606
	var v608 int32
	_ = v608
	var v614 float64
	_ = v614
	var v617 float64
	_ = v617
	var v618 float64
	_ = v618
	var v626 int32
	_ = v626
	var v631 float64
	_ = v631
	var v633 float64
	_ = v633
	var v634 float64
	_ = v634
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 float64
	_ = v645
	var v652 float64
	_ = v652
	var v655 float64
	_ = v655
	var v665 float64
	_ = v665
	var v670 int32
	_ = v670
	var v671 float64
	_ = v671
	var v674 float64
	_ = v674
	var v675 float64
	_ = v675
	var v676 int32
	_ = v676
	var v677 float64
	_ = v677
	var v679 int32
	_ = v679
	var v693 int32
	_ = v693
	var v702 float64
	_ = v702
	var v706 float64
	_ = v706
	var v707 float64
	_ = v707
	var v712 float64
	_ = v712
	var v719 int32
	_ = v719
	var v722 float64
	_ = v722
	var v724 float64
	_ = v724
	var v725 float64
	_ = v725
	var v727 float64
	_ = v727
	var v731 float64
	_ = v731
	var v735 float64
	_ = v735
	var v745 float64
	_ = v745
	var v765 float64
	_ = v765
	var v775 float64
	_ = v775
	var v786 float64
	_ = v786
	v19 = l4 - int32(1)
	v31 = v19
	v32 = int32(-1)
	goto L1
L1:
	;
	v40 = base.I32_div_s(v31+v32+int32(1), int32(2))
	v44 = F_range_cmp_bounds(m, l0, l3+v40<<(uint(int32(3))%32), l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v50 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v49 = base.B2i32(v44 <= int32(0))
	if v44 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v50 = v40
	goto L7
L6:
	;
	v50 = v32
	goto L7
L7:
	;
	if v44 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = v31
	goto L10
L9:
	;
	v53 = v40 - int32(1)
	goto L10
L10:
	;
	if v50 < v53 {
		v31 = v53
		v32 = v50
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	return float64(0)
L13:
	;
	goto L14
L14:
	;
	v60 = l0 + int32(268)
	v62 = l4 - int32(2)
	if base.Ui32(v50) < base.Ui32(v62) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = v50
	goto L17
L16:
	;
	v64 = v62
	goto L17
L17:
	;
	v67 = l3 + v64<<(uint(int32(3))%32)
	v70 = F_get_position(m, l0, l1, v67, v67+int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v72 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v112 = base.F64_convert_i32_u(v19)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)))
	if v113 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v100 != int32(1) {
		v110 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L19
	} else {
		goto L35
	}
L23:
	;
	v78 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L25
L24:
	;
	v78 = float64(1)
	goto L25
L25:
	;
	if v77 != 0 {
		v110 = v78
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v79 == int32(0) {
		v110 = v78
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v82 = float64(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v87 = F_FunctionCall2Coll(m, v60, v84, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v87)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v89)&int64(9223372036854775807)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v95 = v82
	goto L31
L30:
	;
	v95 = v89
	goto L31
L31:
	;
	if base.F64_lt(v89, float64(0)) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = v82
	goto L34
L33:
	;
	v98 = v95
	goto L34
L34:
	;
	v110 = v98
	goto L19
L35:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v105 == v106 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v108 = float64(0)
	goto L38
L37:
	;
	v108 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L38
L38:
	;
	v110 = v108
	goto L19
L39:
	;
	v153 = int32(0)
	v158 = float64(0)
	if base.F64_lt(v151, v158) != 0 {
		v422 = v158
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v118 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v141 != int32(1) {
		v151 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L39
	} else {
		goto L55
	}
L43:
	;
	v119 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L44:
	;
	v119 = float64(1)
	goto L45
L45:
	;
	if v118 != 0 {
		v151 = v119
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v120 == int32(0) {
		v151 = v119
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v123 = float64(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v128 = F_FunctionCall2Coll(m, v60, v125, v126, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v130)&int64(9223372036854775807)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v136 = v123
	goto L51
L50:
	;
	v136 = v130
	goto L51
L51:
	;
	if base.F64_lt(v130, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v139 = v123
	goto L54
L53:
	;
	v139 = v136
	goto L54
L54:
	;
	v151 = v139
	goto L39
L55:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+6)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v146 == v147 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v149 = float64(0)
	goto L58
L57:
	;
	v149 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L58
L58:
	;
	v151 = v149
	goto L39
L59:
	;
	v434 = base.F64_add(base.F64_div(base.F64_mul(v70, base.F64_sub(float64(1), v422)), v112), float64(0))
	if v64 != 0 {
		goto L131
	} else {
		goto L132
	}
L60:
	;
	goto L59
L61:
	;
	v169 = float64(1)
	v171 = base.F64_abs(v151)
	if base.F64_eq(v171, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v174 = v153
	goto L64
L63:
	;
	v174 = int32(0)
	goto L64
L64:
	;
	if v174 != 0 {
		v422 = v169
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v177 = l6 - int32(1)
	if v177 < int32(0) {
		v422 = v169
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v181 = v177
	v185 = int32(-1)
	goto L67
L67:
	;
	v201 = int32(2)
	v202 = base.I32_div_s(v181+v185+int32(1), v201)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l5+v202<<(uint(v201)%32))))
	v207 = *(*float64)(unsafe.Add(mBase, uint32(v206)))
	if base.F64_gt(v110, v207) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v177 <= v217 {
		v422 = v169
		goto L60
	} else {
		goto L80
	}
L69:
	;
	if v217 < v215 {
		v181 = v215
		v185 = v217
		goto L67
	} else {
		goto L79
	}
L70:
	;
	v215 = v181
	v217 = v202
	goto L69
L71:
	;
	goto L72
L72:
	;
	v212 = v153 & base.F64_ge(v110, v207)
	if v212 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v213 = v181
	goto L75
L74:
	;
	v213 = v202 - int32(1)
	goto L75
L75:
	;
	if v212 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v214 = v202
	goto L78
L77:
	;
	v214 = v185
	goto L78
L78:
	;
	v215 = v213
	v217 = v214
	goto L69
L79:
	;
	goto L68
L80:
	;
	if v217 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v274 = base.F64_convert_i32_u(v177)
	v275 = base.F64_div(base.F64_add(v271, base.F64_convert_i32_u(v265)), v274)
	if base.F64_eq(v110, v151) != 0 {
		v422 = v275
		goto L60
	} else {
		goto L96
	}
L82:
	;
	v265 = int32(0)
	v271 = float64(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v226 = l5 + v217<<(uint(int32(2))%32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	v229 = base.F64_abs(v228)
	v230 = math.Float64frombits(uint64(0x7ff0000000000000))
	v231 = base.F64_eq(v229, v230)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v232)))
	v234 = base.F64_abs(v233)
	v236 = base.F64_eq(v234, v230)
	if v231|v236 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.F64_eq(base.F64_abs(v110), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v265 = v217
		v271 = float64(0.5)
		goto L81
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v250 = int32(0)
	if v231|base.B2i32(v236 == v250) == v250 {
		v265 = v217
		v271 = float64(1)
		goto L81
	} else {
		goto L89
	}
L88:
	;
	v265 = v217
	v271 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v228, v110), base.F64_sub(v228, v233)))
	goto L81
L89:
	;
	if base.F64_eq(v229, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v259 = float64(0)
	goto L92
L91:
	;
	v259 = float64(0.5)
	goto L92
L92:
	;
	if base.F64_eq(v234, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v263 = v259
	goto L95
L94:
	;
	v263 = float64(0.5)
	goto L95
L95:
	;
	v265 = v217
	v271 = v263
	goto L81
L96:
	;
	if v177 <= v265 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v388 = float64(0)
	v392 = base.F64_div(base.F64_add(v384, base.F64_convert_i32_u(v376)), v274)
	if base.F64_gt(v381, v388)|base.F64_gt(v392, v388) != 0 {
		goto L124
	} else {
		goto L125
	}
L98:
	;
	v376 = v265
	v379 = v110
	v381 = v275
	v382 = v158
	v384 = v158
	goto L97
L99:
	;
	goto L100
L100:
	;
	v283 = v265
	v288 = v275
	v290 = v158
	v291 = v110
	goto L102
L101:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l5+v283<<(uint(int32(2))%32))))
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v327)))
	if base.F64_eq(v302, v328) != 0 {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v297 = v283 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l5+v297<<(uint(int32(2))%32))))
	v302 = *(*float64)(unsafe.Add(mBase, uint32(v301)))
	if base.F64_lt(v302, v151)|v153&base.F64_ge(v151, v302) == int32(0) {
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v376 = v177
	v379 = v302
	v381 = v312
	v382 = v322
	v384 = v158
	goto L97
L104:
	;
	v309 = float64(0)
	v312 = base.F64_div(base.F64_convert_i32_u(v283), v274)
	if base.F64_gt(v288, v309)|base.F64_gt(v312, v309) != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v322 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v288, v312), float64(0.5)), base.F64_sub(v302, v291)), v290)
	goto L107
L106:
	;
	v322 = v290
	goto L107
L107:
	;
	if v177 != v297 {
		v283 = v297
		v288 = v312
		v290 = v322
		v291 = v302
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L103
L109:
	;
	v369 = float64(0)
	goto L111
L110:
	;
	v331 = base.F64_abs(v302)
	v332 = math.Float64frombits(uint64(0x7ff0000000000000))
	v333 = base.F64_eq(v331, v332)
	v334 = base.F64_abs(v328)
	v336 = base.F64_eq(v334, v332)
	if v333|v336 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v376 = v283
	v379 = v291
	v381 = v288
	v382 = v290
	v384 = v369
	goto L97
L112:
	;
	v369 = v364
	goto L111
L113:
	;
	if base.F64_eq(base.F64_abs(v151), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v364 = float64(0.5)
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v350 = int32(0)
	if v333|base.B2i32(v336 == v350) == v350 {
		v364 = float64(1)
		goto L112
	} else {
		goto L117
	}
L116:
	;
	v364 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v302, v151), base.F64_sub(v302, v328)))
	goto L112
L117:
	;
	if base.F64_eq(v331, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v359 = float64(0)
	goto L120
L119:
	;
	v359 = float64(0.5)
	goto L120
L120:
	;
	if base.F64_eq(v334, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v363 = v359
	goto L123
L122:
	;
	v363 = float64(0.5)
	goto L123
L123:
	;
	v364 = v363
	goto L112
L124:
	;
	v402 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v381, v392), float64(0.5)), base.F64_sub(v151, v379)), v382)
	goto L126
L125:
	;
	v402 = v382
	goto L126
L126:
	;
	if base.F64_eq(v171, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if base.F64_eq(base.F64_abs(v402), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v422 = float64(0.5)
		goto L60
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v422 = base.F64_div(v402, base.F64_sub(v151, v110))
	goto L60
L130:
	;
	goto L129
L131:
	;
	v439 = v64
	v442 = v151
	v443 = v434
	goto L134
L132:
	;
	v786 = v434
	goto L133
L133:
	;
	return v786
L134:
	;
	v452 = v439 - int32(1)
	v455 = l3 + v452<<(uint(int32(3))%32)
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+4)))
	if v456 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v786 = v775
	goto L133
L136:
	;
	v496 = int32(0)
	v501 = float64(0)
	if base.F64_lt(v493, v501) != 0 {
		v765 = v501
		goto L157
	} else {
		goto L158
	}
L137:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v461 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v484 != int32(1) {
		v493 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L136
	} else {
		goto L152
	}
L140:
	;
	v462 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L142
L141:
	;
	v462 = float64(1)
	goto L142
L142:
	;
	if v461 != 0 {
		v493 = v462
		goto L136
	} else {
		goto L143
	}
L143:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v463 == int32(0) {
		v493 = v462
		goto L136
	} else {
		goto L144
	}
L144:
	;
	v466 = float64(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v471 = F_FunctionCall2Coll(m, v60, v468, v469, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	v473 = *(*float64)(unsafe.Add(mBase, uint32(v471)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v473)&int64(9223372036854775807)) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v479 = v466
	goto L148
L147:
	;
	v479 = v473
	goto L148
L148:
	;
	if base.F64_lt(v473, float64(0)) != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v482 = v466
	goto L151
L150:
	;
	v482 = v479
	goto L151
L151:
	;
	v493 = v482
	goto L136
L152:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+6)))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v489 == v490 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v492 = float64(0)
	goto L155
L154:
	;
	v492 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L155
L155:
	;
	v493 = v492
	goto L136
L156:
	;
	v775 = base.F64_add(v443, base.F64_div(base.F64_sub(float64(1), v765), v112))
	if base.Ui32(int32(1)) < base.Ui32(v439) {
		v439 = v452
		v442 = v493
		v443 = v775
		goto L134
	} else {
		goto L228
	}
L157:
	;
	goto L156
L158:
	;
	v512 = float64(1)
	v514 = base.F64_abs(v493)
	if base.F64_eq(v514, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v517 = v496
	goto L161
L160:
	;
	v517 = int32(0)
	goto L161
L161:
	;
	if v517 != 0 {
		v765 = v512
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v520 = l6 - int32(1)
	if v520 < int32(0) {
		v765 = v512
		goto L157
	} else {
		goto L163
	}
L163:
	;
	v524 = v520
	v528 = int32(-1)
	goto L164
L164:
	;
	v544 = int32(2)
	v545 = base.I32_div_s(v524+v528+int32(1), v544)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l5+v545<<(uint(v544)%32))))
	v550 = *(*float64)(unsafe.Add(mBase, uint32(v549)))
	if base.F64_gt(v442, v550) != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	if v520 <= v560 {
		v765 = v512
		goto L157
	} else {
		goto L177
	}
L166:
	;
	if v560 < v558 {
		v524 = v558
		v528 = v560
		goto L164
	} else {
		goto L176
	}
L167:
	;
	v558 = v524
	v560 = v545
	goto L166
L168:
	;
	goto L169
L169:
	;
	v555 = v496 & base.F64_ge(v442, v550)
	if v555 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v556 = v524
	goto L172
L171:
	;
	v556 = v545 - int32(1)
	goto L172
L172:
	;
	if v555 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v557 = v545
	goto L175
L174:
	;
	v557 = v528
	goto L175
L175:
	;
	v558 = v556
	v560 = v557
	goto L166
L176:
	;
	goto L165
L177:
	;
	if v560 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v617 = base.F64_convert_i32_u(v520)
	v618 = base.F64_div(base.F64_add(v614, base.F64_convert_i32_u(v608)), v617)
	if base.F64_eq(v442, v493) != 0 {
		v765 = v618
		goto L157
	} else {
		goto L193
	}
L179:
	;
	v608 = int32(0)
	v614 = float64(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v569 = l5 + v560<<(uint(int32(2))%32)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v570)))
	v572 = base.F64_abs(v571)
	v573 = math.Float64frombits(uint64(0x7ff0000000000000))
	v574 = base.F64_eq(v572, v573)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v576 = *(*float64)(unsafe.Add(mBase, uint32(v575)))
	v577 = base.F64_abs(v576)
	v579 = base.F64_eq(v577, v573)
	if v574|v579 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if base.F64_eq(base.F64_abs(v442), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v608 = v560
		v614 = float64(0.5)
		goto L178
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v593 = int32(0)
	if v574|base.B2i32(v579 == v593) == v593 {
		v608 = v560
		v614 = float64(1)
		goto L178
	} else {
		goto L186
	}
L185:
	;
	v608 = v560
	v614 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v571, v442), base.F64_sub(v571, v576)))
	goto L178
L186:
	;
	if base.F64_eq(v572, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v602 = float64(0)
	goto L189
L188:
	;
	v602 = float64(0.5)
	goto L189
L189:
	;
	if base.F64_eq(v577, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v606 = v602
	goto L192
L191:
	;
	v606 = float64(0.5)
	goto L192
L192:
	;
	v608 = v560
	v614 = v606
	goto L178
L193:
	;
	if v520 <= v608 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v731 = float64(0)
	v735 = base.F64_div(base.F64_add(v727, base.F64_convert_i32_u(v719)), v617)
	if base.F64_gt(v724, v731)|base.F64_gt(v735, v731) != 0 {
		goto L221
	} else {
		goto L222
	}
L195:
	;
	v719 = v608
	v722 = v442
	v724 = v618
	v725 = v501
	v727 = v501
	goto L194
L196:
	;
	goto L197
L197:
	;
	v626 = v608
	v631 = v618
	v633 = v501
	v634 = v442
	goto L199
L198:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l5+v626<<(uint(int32(2))%32))))
	v671 = *(*float64)(unsafe.Add(mBase, uint32(v670)))
	if base.F64_eq(v645, v671) != 0 {
		goto L206
	} else {
		goto L207
	}
L199:
	;
	v640 = v626 + int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l5+v640<<(uint(int32(2))%32))))
	v645 = *(*float64)(unsafe.Add(mBase, uint32(v644)))
	if base.F64_lt(v645, v493)|v496&base.F64_ge(v493, v645) == int32(0) {
		goto L198
	} else {
		goto L201
	}
L200:
	;
	v719 = v520
	v722 = v645
	v724 = v655
	v725 = v665
	v727 = v501
	goto L194
L201:
	;
	v652 = float64(0)
	v655 = base.F64_div(base.F64_convert_i32_u(v626), v617)
	if base.F64_gt(v631, v652)|base.F64_gt(v655, v652) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v665 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v631, v655), float64(0.5)), base.F64_sub(v645, v634)), v633)
	goto L204
L203:
	;
	v665 = v633
	goto L204
L204:
	;
	if v520 != v640 {
		v626 = v640
		v631 = v655
		v633 = v665
		v634 = v645
		goto L199
	} else {
		goto L205
	}
L205:
	;
	goto L200
L206:
	;
	v712 = float64(0)
	goto L208
L207:
	;
	v674 = base.F64_abs(v645)
	v675 = math.Float64frombits(uint64(0x7ff0000000000000))
	v676 = base.F64_eq(v674, v675)
	v677 = base.F64_abs(v671)
	v679 = base.F64_eq(v677, v675)
	if v676|v679 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	v719 = v626
	v722 = v634
	v724 = v631
	v725 = v633
	v727 = v712
	goto L194
L209:
	;
	v712 = v707
	goto L208
L210:
	;
	if base.F64_eq(base.F64_abs(v493), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v707 = float64(0.5)
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v693 = int32(0)
	if v676|base.B2i32(v679 == v693) == v693 {
		v707 = float64(1)
		goto L209
	} else {
		goto L214
	}
L213:
	;
	v707 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v645, v493), base.F64_sub(v645, v671)))
	goto L209
L214:
	;
	if base.F64_eq(v674, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v702 = float64(0)
	goto L217
L216:
	;
	v702 = float64(0.5)
	goto L217
L217:
	;
	if base.F64_eq(v677, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v706 = v702
	goto L220
L219:
	;
	v706 = float64(0.5)
	goto L220
L220:
	;
	v707 = v706
	goto L209
L221:
	;
	v745 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v724, v735), float64(0.5)), base.F64_sub(v493, v722)), v725)
	goto L223
L222:
	;
	v745 = v725
	goto L223
L223:
	;
	if base.F64_eq(v514, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	if base.F64_eq(base.F64_abs(v745), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v765 = float64(0.5)
		goto L157
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v765 = base.F64_div(v745, base.F64_sub(v493, v442))
	goto L157
L227:
	;
	goto L226
L228:
	;
	goto L135
}
func F_calc_joinrel_size_estimate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 float64
	_ = v56
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 float64
	_ = v394
	var v395 float64
	_ = v395
	var v398 float64
	_ = v398
	var v399 float64
	_ = v399
	var v401 float64
	_ = v401
	var v403 float64
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v421 float64
	_ = v421
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 float64
	_ = v458
	var v459 int32
	_ = v459
	var v463 float64
	_ = v463
	var v464 float64
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 float64
	_ = v479
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 float64
	_ = v502
	var v505 float64
	_ = v505
	var v508 float64
	_ = v508
	var v516 int32
	_ = v516
	var v518 float64
	_ = v518
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v662 int32
	_ = v662
	var v677 float64
	_ = v677
	var v678 int32
	_ = v678
	var v680 float64
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 float64
	_ = v687
	var v688 int32
	_ = v688
	var v697 float64
	_ = v697
	var v699 float64
	_ = v699
	var v718 float64
	_ = v718
	var v720 float64
	_ = v720
	var v724 float64
	_ = v724
	var v726 float64
	_ = v726
	var v728 float64
	_ = v728
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v756 float64
	_ = v756
	var v760 float64
	_ = v760
	var v769 float64
	_ = v769
	var v773 float64
	_ = v773
	v12 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v33 == v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(1)<<(uint(v32)%32)&int32(174) != 0 {
		goto L121
	} else {
		goto L122
	}
L2:
	;
	v516 = l7
	v518 = float64(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v38 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v516 = l7
	v518 = float64(1)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v47 = base.B2i32(v32&int32(-2) != int32(4))
	v50 = l7
	v56 = float64(1)
	v71 = v12
	goto L8
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v71<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = F_bms_is_member(m, v80, v43)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v502 = float64(0)
	if base.F64_lt(v479, v502) != 0 {
		v508 = v502
		goto L117
	} else {
		goto L118
	}
L10:
	;
	v499 = v71 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v499 < v500 {
		v50 = v473
		v56 = v479
		v71 = v499
		goto L8
	} else {
		goto L116
	}
L11:
	;
	if v50 == l7 {
		goto L44
	} else {
		goto L45
	}
L12:
	;
	if v32&int32(-2) != int32(4) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	return float64(0)
L14:
	;
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v86 = F_bms_is_member(m, v85, v42)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v89 = F_bms_is_member(m, v88, v43)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	if v86 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v89 == int32(0) {
		v473 = v50
		v479 = v56
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v94 = F_bms_is_member(m, v93, v42)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v96 = int32(0)
	if base.B2i32(v94 == v96)|base.B2i32(v47 == v96) != 0 {
		v473 = v50
		v479 = v56
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v152 = int32(0)
	goto L11
L24:
	;
	v152 = int32(0)
	goto L11
L25:
	;
	goto L26
L26:
	;
	v104 = int32(0)
	if v42 == v104 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v149 != int32(1) {
		v473 = v50
		v479 = v56
		goto L10
	} else {
		goto L43
	}
L28:
	;
	v149 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v112 = int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v113 <= v112 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = v112
	goto L33
L32:
	;
	v116 = v113
	goto L33
L33:
	;
	v120 = int32(0)
	v122 = v104
	goto L34
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(8)+v120<<(uint(int32(2))%32))))
	if v129 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v149 = v141
	goto L27
L36:
	;
	goto L35
L37:
	;
	v130 = int32(2)
	if v122 != 0 {
		v141 = v130
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v136 = v122
	goto L39
L39:
	;
	v138 = v120 + int32(1)
	if v138 != v116 {
		v120 = v138
		v122 = v136
		goto L34
	} else {
		goto L42
	}
L40:
	;
	v131 = int32(1)
	if base.Ui32(v131) < base.Ui32(base.I32_popcnt(v129)) {
		v141 = v130
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v136 = v131
	goto L39
L42:
	;
	v141 = v136
	goto L36
L43:
	;
	v152 = int32(1)
	goto L11
L44:
	;
	v154 = F_list_copy(m, v50)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L47
	}
L45:
	;
	v156 = v50
	goto L46
L46:
	;
	if v156 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v156 = v154
	goto L46
L48:
	;
	v159 = int32(0)
	v161 = F_list_concat(m, v159, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v166 = v79 + int32(288)
	v167 = int32(0)
	v171 = v156
	v181 = v167
	v184 = v167
	goto L52
L51:
	;
	v473 = v161
	v479 = v56
	goto L10
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v181 < v196 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v365 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L54:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if int32(0) < v198 {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v352 = v171
	v365 = v184
	goto L56
L56:
	;
	goto L53
L57:
	;
	if v323 != 0 {
		v171 = v323
		v181 = v324 + int32(1)
		v184 = v336
		goto L52
	} else {
		goto L86
	}
L58:
	;
	v317 = F_list_delete_nth_cell(m, v171, v181)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L84
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v181<<(uint(int32(2))%32))))
	v210 = int32(0)
	v220 = v198
	goto L62
L60:
	;
	goto L61
L61:
	;
	v323 = v171
	v324 = v181
	v336 = v184
	goto L57
L62:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v205)+60))
	if v234 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L61
L64:
	;
	v286 = v210 + int32(1)
	if v286 < v284 {
		v210 = v286
		v220 = v284
		goto L62
	} else {
		goto L83
	}
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v166+v210<<(uint(int32(2))%32))))
	if v238 != v234 {
		v284 = v220
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(544)+v210<<(uint(int32(2))%32))))
	v244 = int32(0)
	if v243 == v244 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L58
L69:
	;
	if v282 != 0 {
		goto L58
	} else {
		goto L82
	}
L70:
	;
	v282 = int32(0)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v250 <= int32(0) {
		v276 = v244
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v282 = v276
	goto L69
L74:
	;
	v253 = int32(0)
	if v253 < v250 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v256 = v250
	goto L77
L76:
	;
	v256 = v253
	goto L77
L77:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v259 = int32(0)
	goto L78
L78:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v257+v259<<(uint(int32(2))%32))))
	v268 = base.B2i32(v267 == v205)
	if v267 == v205 {
		v276 = v268
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v276 = v268
	goto L73
L80:
	;
	v270 = v259 + int32(1)
	if v270 != v256 {
		v259 = v270
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v284 = v283
	goto L64
L83:
	;
	goto L63
L84:
	;
	v319 = F_lappend(m, v184, v205)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	v323 = v317
	v324 = v181 - int32(1)
	v336 = v319
	goto L57
L86:
	;
	v352 = v323
	v365 = v336
	goto L56
L87:
	;
	v380 = F_list_concat(m, v352, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v79)+284))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v79)+272))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v382 != v383+(v384-v385) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v473 = v380
	v479 = v56
	goto L10
L91:
	;
	v389 = F_list_concat(m, v352, v365)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v392 = F_find_base_rel(m, l0, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L13
	} else {
		goto L95
	}
L94:
	;
	v473 = v389
	v479 = v56
	goto L10
L95:
	;
	v394 = *(*float64)(unsafe.Add(mBase, uint32(v392)+120))
	v395 = float64(1)
	if base.F64_gt(v394, v395) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v398 = v394
	goto L98
L97:
	;
	v398 = v395
	goto L98
L98:
	;
	if v152 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v392)+16))
	v401 = v399
	goto L101
L100:
	;
	v401 = float64(1)
	goto L101
L101:
	;
	v403 = base.F64_mul(v56, base.F64_div(v401, v398))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v404 <= int32(0) {
		v473 = v352
		v479 = v403
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v407 <= int32(0) {
		v473 = v352
		v479 = v403
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v416 = int32(0)
	v421 = v403
	goto L104
L104:
	;
	v441 = v416 << (uint(int32(2)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v166+v441)))
	if v443 == int32(0) {
		v464 = v421
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v473 = v352
	v479 = v464
	goto L10
L106:
	;
	v468 = v416 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v468 < v469 {
		v416 = v468
		v421 = v464
		goto L104
	} else {
		goto L115
	}
L107:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+40)))
	if v446 != int32(1) {
		v464 = v421
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(416)+v441)))
	v451 = int32(0)
	v453 = F_ec_search_derived_clause_for_ems(m, l0, v443, v450, v451, v451)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	if v453 == int32(0) {
		v464 = v421
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v458 = F_clause_selectivity(m, l0, v453, int32(0), v32, l6)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	if base.F64_gt(v458, float64(0)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v463 = base.F64_div(v421, v458)
	goto L114
L113:
	;
	v463 = v421
	goto L114
L114:
	;
	v464 = v463
	goto L106
L115:
	;
	goto L105
L116:
	;
	goto L9
L117:
	;
	v516 = v473
	v518 = v508
	goto L1
L118:
	;
	v505 = float64(1)
	if base.F64_gt(v479, v505) != 0 {
		v508 = v505
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v516 = v473
	v518 = v479
	goto L1
L120:
	;
	switch v32 {
	case 0:
		goto L158
	case 1:
		goto L163
	case 2:
		goto L162
	default:
		goto L159
	case 4:
		goto L161
	case 5:
		goto L160
	}
L121:
	;
	v540 = int32(0)
	if v516 == v540 {
		v651 = v540
		v662 = v540
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v687 = F_clauselist_selectivity(m, l0, v516, int32(0), v32, l6)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L13
	} else {
		goto L156
	}
L124:
	;
	v677 = F_clauselist_selectivity(m, l0, v662, int32(0), v32, l6)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L13
	} else {
		goto L152
	}
L125:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v544 <= int32(0) {
		v651 = v540
		v662 = v540
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v550 = v540
	v551 = int32(0)
	v561 = v540
	goto L127
L127:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v575+v551<<(uint(int32(2))%32))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+8)))
	if v580 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v651 = v643
	v662 = v644
	goto L124
L129:
	;
	v646 = v551 + int32(1)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v646 < v647 {
		v550 = v643
		v551 = v646
		v561 = v644
		goto L127
	} else {
		goto L151
	}
L130:
	;
	v641 = F_lappend(m, v561, v579)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L150
	}
L131:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v579)+32))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v585 = int32(0)
	if v583 == v585 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	goto L133
L133:
	;
	v639 = F_lappend(m, v550, v579)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L149
	}
L134:
	;
	if v638 != 0 {
		goto L130
	} else {
		goto L148
	}
L135:
	;
	v638 = int32(1)
	goto L134
L136:
	;
	goto L137
L137:
	;
	if v584 == int32(0) {
		v631 = v585
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v638 = v631
	goto L134
L139:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v595 < v594 {
		v631 = v585
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v597 = int32(1)
	if v594 <= v597 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v600 = v597
	goto L143
L142:
	;
	v600 = v594
	goto L143
L143:
	;
	v601 = int32(8)
	v606 = int32(0)
	goto L144
L144:
	;
	v613 = v606 << (uint(int32(2)) % 32)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v583+v601+v613)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v584+v601+v613)))
	v620 = v615 & (v617 ^ int32(-1))
	v622 = base.B2i32(v620 == int32(0))
	if v620 != 0 {
		v631 = v622
		goto L138
	} else {
		goto L146
	}
L145:
	;
	v631 = v622
	goto L138
L146:
	;
	v624 = v606 + int32(1)
	if v624 != v600 {
		v606 = v624
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	goto L133
L149:
	;
	v643 = v639
	v644 = v561
	goto L129
L150:
	;
	v643 = v550
	v644 = v641
	goto L129
L151:
	;
	goto L128
L152:
	;
	v680 = F_clauselist_selectivity(m, l0, v651, int32(0), v32, l6)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L13
	} else {
		goto L153
	}
L153:
	;
	F_list_free(m, v662)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	F_list_free(m, v651)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	v697 = v677
	v699 = v680
	goto L120
L156:
	;
	v697 = v687
	v699 = float64(0)
	goto L120
L157:
	;
	m.G0 = v30 + int32(16)
	v760 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v756)&int64(9223372036854775807)))|base.F64_gt(v756, v760) != 0 {
		v773 = v760
		goto L176
	} else {
		goto L177
	}
L158:
	;
	v756 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v697)
	goto L157
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L13
	} else {
		goto L173
	}
L160:
	;
	v756 = base.F64_mul(v699, base.F64_mul(l4, base.F64_sub(float64(1), base.F64_mul(v518, v697))))
	goto L157
L161:
	;
	v756 = base.F64_mul(base.F64_mul(l4, v518), v697)
	goto L157
L162:
	;
	v724 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v697)
	if base.F64_gt(l4, v724) != 0 {
		goto L167
	} else {
		goto L168
	}
L163:
	;
	v718 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v697)
	if base.F64_gt(l4, v718) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v720 = l4
	goto L166
L165:
	;
	v720 = v718
	goto L166
L166:
	;
	v756 = base.F64_mul(v699, v720)
	goto L157
L167:
	;
	v726 = l4
	goto L169
L168:
	;
	v726 = v724
	goto L169
L169:
	;
	if base.F64_lt(v726, l5) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v728 = l5
	goto L172
L171:
	;
	v728 = v726
	goto L172
L172:
	;
	v756 = base.F64_mul(v699, v728)
	goto L157
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v32
	F_errmsg_internal(m, int32(_a_F_calc_joinrel_size_estimate_0), v30)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L13
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_calc_joinrel_size_estimate_1), int32(_a_F_calc_joinrel_size_estimate_2), int32(_a_F_calc_joinrel_size_estimate_3))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	return v773
L177:
	;
	v769 = float64(1)
	if base.F64_le(v756, v769) != 0 {
		v773 = v769
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v773 = base.F64_nearest(v756)
	goto L176
}
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 float32
	_ = v32
	var v39 float64
	_ = v39
	var v42 float32
	_ = v42
	var v48 float64
	_ = v48
	var v52 float64
	_ = v52
	var v56 float32
	_ = v56
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v71 float32
	_ = v71
	var v78 float64
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v485 int32
	_ = v485
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var __phi590 int32
	_ = __phi590
	var v592 int32
	_ = v592
	var __phi592 int32
	_ = __phi592
	var v594 int32
	_ = v594
	var __phi594 int32
	_ = __phi594
	var v596 int32
	_ = v596
	var __phi596 int32
	_ = __phi596
	var v597 int32
	_ = v597
	var __phi597 int32
	_ = __phi597
	var v599 int32
	_ = v599
	var __phi599 int32
	_ = __phi599
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 float64
	_ = v690
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v712 float64
	_ = v712
	var v713 float64
	_ = v713
	var v714 float64
	_ = v714
	var v719 int32
	_ = v719
	var v729 int32
	_ = v729
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1082 int32
	_ = v1082
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1199 int32
	_ = v1199
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1282 float64
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1306 float64
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1318 float64
	_ = v1318
	var v1319 float64
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1345 float64
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1370 float64
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1383 float64
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1489 float64
	_ = v1489
	var v1511 float64
	_ = v1511
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1540 int32
	_ = v1540
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1605 float64
	_ = v1605
	var v1616 int32
	_ = v1616
	var v1626 float64
	_ = v1626
	var v1631 int32
	_ = v1631
	var v1637 float64
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1648 float64
	_ = v1648
	var v1653 float64
	_ = v1653
	var v1659 float64
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1692 float32
	_ = v1692
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v32 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	if base.F32_ge(v32, float32(0)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L20
	} else {
		goto L244
	}
L2:
	;
	if base.F32_gt(v32, float32(1)) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v39 = float64(0.10000000149011612)
	goto L4
L4:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+16)) = base.F64_div(float64(1), v39)
	v42 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.F32_ge(v42, float32(0)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v39 = base.F64_promote_f32(v32)
	goto L4
L6:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = base.F64_div(float64(1), v52)
	v56 = *(*float32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F32_ge(v56, float32(0)) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v52 = float64(0.20000000298023224)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v48 = base.F64_promote_f32(v42)
	*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v48
	if base.F32_gt(v42, float32(1)) != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v52 = v48
	goto L6
L11:
	;
	v67 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = base.F64_div(v67, v66)
	v71 = *(*float32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.F32_ge(v71, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v66 = float64(0.4000000059604645)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v62 = base.F64_promote_f32(v56)
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v62
	if base.F32_gt(v56, float32(1)) != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v66 = v62
	goto L11
L16:
	;
	if base.F32_gt(v71, float32(1)) != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v78 = float64(1)
	goto L18
L18:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = base.F64_div(v67, v78)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l2
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v85 = F_palloc0(m, v82*int32(_a_F_calc_rank_cd_0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v78 = base.F64_promote_f32(v71)
	goto L18
L20:
	;
	return float32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v85
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v93 = F_palloc(m, v90*int32(48))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v95 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	m.G0 = v29 - int32(-64)
	return v1692
L24:
	;
	F_pg_qsort(m, v515, v517, int32(12), int32(1517))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L20
	} else {
		goto L94
	}
L25:
	;
	v98 = int32(8)
	v101 = l1 + v98
	v106 = l2
	v108 = v90 << (uint(int32(2)) % 32)
	v112 = v93
	v113 = v5
	v114 = v5
	goto L28
L26:
	;
	v541 = v85
	v549 = v93
	goto L27
L27:
	;
	F_pfree(m, v549)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L20
	} else {
		goto L92
	}
L28:
	;
	v132 = l2 + v98 + v113*int32(12)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v133 != int32(1) {
		v511 = v108
		v515 = v112
		v517 = v114
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if int32(0) < v517 {
		goto L24
	} else {
		goto L91
	}
L30:
	;
	v534 = v113 + int32(1)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if v534 < v536 {
		v106 = v535
		v108 = v511
		v112 = v515
		v113 = v534
		v114 = v517
		goto L28
	} else {
		goto L90
	}
L31:
	;
	v137 = v27 + int32(-4)
	v138 = int32(0)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v138
	v148 = l1 + int32(8)
	v151 = v148 + v144<<(uint(int32(2))%32)
	if v144 <= v138 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v291 == int32(0) {
		v511 = v108
		v515 = v112
		v517 = v114
		goto L30
	} else {
		goto L62
	}
L33:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+2)))
	if v219 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	v213 = v151
	v214 = v148
	v215 = v151
	goto L33
L35:
	;
	goto L36
L36:
	;
	v161 = v148
	v162 = v151
	goto L37
L37:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v167 = int32(12)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v177 = int32(2)
	v184 = base.I32_div_s((v162-v161)>>(uint(v177)%32), v177)
	v187 = v161 + v184<<(uint(v177)%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v196 = int32(0)
	v197 = F_tsCompareString(m, v106+int32(8)+v166*v167+int32(base.Ui32(v170)>>(uint(v167)%32)), v170&int32(4095), v148+v176<<(uint(v177)%32)+int32(base.Ui32(v188)>>(uint(v167)%32)), int32(base.Ui32(v188)>>(uint(int32(1))%32))&int32(2047), v196)
	mBase = m.M
	if v197 == v196 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v213 = v187
	v214 = v206
	v215 = v207
	goto L33
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = int32(1)
	v213 = v187
	v214 = v161
	v215 = v187
	goto L33
L40:
	;
	goto L41
L41:
	;
	v205 = base.B2i32(int32(0) < v197)
	if int32(0) < v197 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v206 = v187 + int32(4)
	goto L44
L43:
	;
	v206 = v161
	goto L44
L44:
	;
	if int32(0) < v197 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v207 = v162
	goto L47
L46:
	;
	v207 = v187
	goto L47
L47:
	;
	if base.Ui32(v206) < base.Ui32(v207) {
		v161 = v206
		v162 = v207
		goto L37
	} else {
		goto L48
	}
L48:
	;
	goto L38
L49:
	;
	v287 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v287 < v288 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = int32(0)
	if base.Ui32(v214) < base.Ui32(v215) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v225 = v213
	goto L53
L52:
	;
	v225 = v215
	goto L53
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v148+v226<<(uint(int32(2))%32)) <= base.Ui32(v225) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v237 = v226
	v238 = v225
	goto L55
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v244 = int32(12)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v260 = int32(1)
	v265 = F_tsCompareString(m, v106+int32(8)+v243*v244+int32(base.Ui32(v247)>>(uint(v244)%32)), v247&int32(4095), v148+v237<<(uint(int32(2))%32)+int32(base.Ui32(v256)>>(uint(v244)%32)), int32(base.Ui32(v256)>>(uint(v260)%32))&int32(2047), v260)
	mBase = m.M
	if v265 != 0 {
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L49
L57:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v266 + int32(1)
	v271 = v238 + int32(4)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v271) < base.Ui32(v148+v272<<(uint(int32(2))%32)) {
		v237 = v272
		v238 = v271
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v291 = v215
	goto L61
L60:
	;
	v291 = v287
	goto L61
L61:
	;
	goto L32
L62:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if v294 <= int32(0) {
		v511 = v108
		v515 = v112
		v517 = v114
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v297 = v294
	v299 = v291
	v301 = v108
	v305 = v112
	v307 = v114
	goto L64
L64:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v323&int32(1) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v511 = v388
	v515 = v392
	v517 = v485
	goto L30
L66:
	;
	if v342 != 0 {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v355 = v301
	v359 = v305
	goto L73
L68:
	;
	v326 = int32(1)
	v336 = (int32(base.Ui32(v323)>>(uint(v326)%32))&int32(2047) + int32(base.Ui32(v323)>>(uint(int32(12))%32)) + v326) & int32(_a_F_calc_rank_cd_1)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v339 = v337 << (uint(int32(2)) % 32)
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336+(v101+v339)))))
	v343 = v307 + v342
	if v301 <= v343 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v346 = v299 + int32(4)
	if (v346-v291)>>(uint(int32(2))%32) < v297 {
		v299 = v346
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v388 = v301
	v392 = v305
	goto L66
L72:
	;
	v511 = v301
	v515 = v305
	v517 = v307
	goto L30
L73:
	;
	v379 = F_repalloc(m, v359, v355*int32(24))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L20
	} else {
		goto L75
	}
L74:
	;
	v388 = v382
	v392 = v379
	goto L66
L75:
	;
	v382 = v355 << (uint(int32(1)) % 32)
	if v382 <= v343 {
		v355 = v382
		v359 = v379
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v413 = l1 + v339 + v336 + int32(10)
	v415 = int32(0)
	v425 = v307
	goto L80
L78:
	;
	v475 = v297
	v485 = v307
	goto L79
L79:
	;
	v502 = v299 + int32(4)
	if (v502-v291)>>(uint(int32(2))%32) < v475 {
		v297 = v475
		v299 = v502
		v301 = v388
		v305 = v392
		v307 = v485
		goto L64
	} else {
		goto L89
	}
L80:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v441 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v475 = v474
	v485 = v469
	goto L79
L82:
	;
	v472 = v415 + int32(1)
	if v472 != v342 {
		v415 = v472
		v425 = v469
		goto L80
	} else {
		goto L88
	}
L83:
	;
	v462 = v392 + v425*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v299
	*(*uint16)(unsafe.Add(mBase, uint32(v462)+8)) = uint16(v459)
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v132
	v469 = v425 + int32(1)
	goto L82
L84:
	;
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413+v415<<(uint(int32(1))%32)))))
	v459 = v447
	goto L83
L85:
	;
	goto L86
L86:
	;
	v448 = int32(1)
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413+v415<<(uint(v448)%32)))))
	if int32(base.Ui32(v441)>>(uint(int32(base.Ui32(v451)>>(uint(int32(14))%32)))%32))&v448 == int32(0) {
		v469 = v425
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v459 = v451
	goto L83
L88:
	;
	goto L81
L89:
	;
	goto L65
L90:
	;
	goto L29
L91:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v541 = v540
	v549 = v515
	goto L27
L92:
	;
	F_pfree(m, v541)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	v1692 = float32(0)
	goto L23
L94:
	;
	v576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515)+8)))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v580 = F_palloc(m, v577<<(uint(int32(2))%32))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v580))) = v582
	if v517 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v655)+8)) = uint16(v662)
	*(*uint16)(unsafe.Add(mBase, uint32(v655)+4)) = uint16(v661)
	*(*int32)(unsafe.Add(mBase, uint32(v655))) = v664
	v685 = int32(12)
	v686 = v655 - v515 + v685
	v688 = base.I32_div_s(v686, v685)
	v690 = float64(0)
	v701 = int32(0)
	v708 = v5
	v712 = v690
	v713 = v690
	v714 = float64(0)
	goto L108
L97:
	;
	v655 = v515
	v661 = int32(1)
	v662 = v576
	v664 = v580
	goto L96
L98:
	;
	goto L99
L99:
	;
	__phi590 = v515
	__phi592 = v515
	__phi594 = v515 + int32(12)
	__phi596 = int32(1)
	__phi597 = v576
	__phi599 = v580
	v590 = __phi590
	v592 = __phi592
	v594 = __phi594
	v596 = __phi596
	v597 = __phi597
	v599 = __phi599
	goto L100
L100:
	;
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592)+20)))
	v617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592)+8)))
	if v616 != v617 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v655 = v645
	v661 = v648
	v662 = v646
	v664 = v647
	goto L96
L102:
	;
	v649 = int32(12)
	v650 = v594 + v649
	v653 = base.I32_div_s(v650-v515, v649)
	if v653 < v517 {
		__phi590 = v645
		__phi592 = v594
		__phi594 = v650
		__phi596 = v648
		__phi597 = v646
		__phi599 = v647
		v590 = __phi590
		v592 = __phi592
		v594 = __phi594
		v596 = __phi596
		v597 = __phi597
		v599 = __phi599
		goto L100
	} else {
		goto L107
	}
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v590)+8)) = uint16(v597)
	*(*uint16)(unsafe.Add(mBase, uint32(v590)+4)) = uint16(v596)
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v599
	v633 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592)+20)))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+4))
	v638 = F_palloc(m, v635<<(uint(int32(2))%32))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L20
	} else {
		goto L106
	}
L104:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v592)+16))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	if v619 != v620 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v592)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v599+base.I32_extend16_s(v596)<<(uint(int32(2))%32)))) = v626
	v645 = v590
	v646 = v597
	v647 = v599
	v648 = v596 + int32(1)
	goto L102
L106:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v592)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v640
	v645 = v590 + int32(12)
	v646 = v633
	v647 = v638
	v648 = int32(1)
	goto L102
L107:
	;
	goto L101
L108:
	;
	if v688 < v701 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if l3&int32(1) == int32(0) {
		v1511 = v712
		goto L205
	} else {
		goto L206
	}
L110:
	;
	v719 = v701
	goto L112
L111:
	;
	v719 = v688
	goto L112
L112:
	;
	v729 = v701
	goto L113
L113:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L20
	} else {
		goto L115
	}
L114:
	;
	goto L109
L115:
	;
	v748 = int32(0)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	if v748 < v750 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v755 = v748
	goto L119
L117:
	;
	goto L118
L118:
	;
	if v729 == v719 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v780 = v755 * int32(_a_F_calc_rank_cd_0)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v783 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v780+v781))) = uint8(v783)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v785+v780)+1)) = uint8(v783)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v789+v780)+4)) = v783
	v794 = v755 + int32(1)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v794 < v796 {
		v755 = v794
		goto L119
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	goto L120
L122:
	;
	goto L114
L123:
	;
	v826 = v729 * int32(12)
	v827 = v515 + v826
	v828 = v827
	v835 = v826
	goto L124
L124:
	;
	v854 = int32(*(*int16)(unsafe.Add(mBase, uint32(v828)+4)))
	if int32(0) < v854 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L122
L126:
	;
	v860 = int32(0)
	goto L129
L127:
	;
	goto L128
L128:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v993 = F_TS_execute(m, v986+int32(8), v27+int32(-56), int32(0), int32(1518))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L20
	} else {
		goto L148
	}
L129:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v884+v860<<(uint(int32(2))%32))))
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	if v889 != int32(1) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L128
L131:
	;
	v957 = v860 + int32(1)
	v958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v828)+4)))
	if v957 < v958 {
		v860 = v957
		goto L129
	} else {
		goto L147
	}
L132:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v898 = base.I32_div_s(v888-v893-int32(8), int32(12))
	v901 = v892 + v898*int32(_a_F_calc_rank_cd_0)
	v902 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v901))) = uint8(v902)
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901)+1)))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	if v905 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v901)+4)) = v949
	goto L131
L134:
	;
	if v904&int32(1) != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v917 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828)+8)))
	v919 = v901 + int32(8)
	v922 = int32(1)
	v925 = v904 & v922
	if v925 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v912 = int32(_a_F_calc_rank_cd_2)
	goto L139
L138:
	;
	v912 = int32(0)
	goto L139
L139:
	;
	v914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v901+v912)+8)) = uint16(v914)
	v949 = int32(1)
	goto L133
L140:
	;
	v926 = int32(_a_F_calc_rank_cd_3) - v905
	goto L142
L141:
	;
	v926 = v905 - v922
	goto L142
L142:
	;
	v930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v919+v926<<(uint(int32(1))%32)))))
	if (v917^v930)&int32(_a_F_calc_rank_cd_4) == int32(0) {
		goto L131
	} else {
		goto L143
	}
L143:
	;
	if v925 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v938 = int32(_a_F_calc_rank_cd_4) - v905
	goto L146
L145:
	;
	v938 = v905
	goto L146
L146:
	;
	v939 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v919+v938<<(uint(v939)%32)))) = uint16(v917)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	v949 = v943 + v939
	goto L133
L147:
	;
	goto L130
L148:
	;
	if v993 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828)+8)))
	v997 = v995 & int32(_a_F_calc_rank_cd_4)
	if v997 == int32(0) {
		goto L122
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v1393 = v828 + int32(12)
	v1394 = v1393 - v515
	if v1394 < v686 {
		v828 = v1393
		v835 = v1394
		goto L124
	} else {
		goto L204
	}
L152:
	;
	v1000 = int32(0)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if v1000 < v1002 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1007 = v1000
	goto L156
L154:
	;
	goto L155
L155:
	;
	if v835 < v826 {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	v1032 = v1007 * int32(_a_F_calc_rank_cd_0)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1035 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1032+v1033))) = uint8(v1035)
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1039 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1037+v1032)+1)) = uint8(v1039)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1041+v1032)+4)) = v1035
	v1046 = v1007 + v1039
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1046 < v1048 {
		v1007 = v1046
		goto L156
	} else {
		goto L158
	}
L157:
	;
	goto L155
L158:
	;
	goto L157
L159:
	;
	v1282 = float64(0)
	if base.Ui32(v1082) <= base.Ui32(v828) {
		goto L192
	} else {
		goto L193
	}
L160:
	;
	v729 = v729 + int32(1)
	goto L113
L161:
	;
	v1082 = v835 + v515
	goto L162
L162:
	;
	v1104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if int32(0) < v1104 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+8)))
	v1252 = v1250 & int32(_a_F_calc_rank_cd_4)
	if base.Ui32(v1252) <= base.Ui32(v997) {
		goto L159
	} else {
		goto L191
	}
L164:
	;
	v1110 = int32(0)
	goto L167
L165:
	;
	goto L166
L166:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1243 = F_TS_execute(m, v1236+int32(8), v27+int32(-56), int32(0), int32(1518))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L20
	} else {
		goto L186
	}
L167:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1134+v1110<<(uint(int32(2))%32))))
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138))))
	if v1139 != int32(1) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v1207 = v1110 + int32(1)
	v1208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if v1207 < v1208 {
		v1110 = v1207
		goto L167
	} else {
		goto L185
	}
L170:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1148 = base.I32_div_s(v1138-v1143-int32(8), int32(12))
	v1151 = v1142 + v1148*int32(_a_F_calc_rank_cd_0)
	v1152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1151))) = uint8(v1152)
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+1)))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	if v1155 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+4)) = v1199
	goto L169
L172:
	;
	if v1154&int32(1) != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+8)))
	v1169 = v1151 + int32(8)
	v1172 = int32(1)
	v1175 = v1154 & v1172
	if v1175 != 0 {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v1162 = int32(_a_F_calc_rank_cd_2)
	goto L177
L176:
	;
	v1162 = int32(0)
	goto L177
L177:
	;
	v1164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1151+v1162)+8)) = uint16(v1164)
	v1199 = int32(1)
	goto L171
L178:
	;
	v1176 = int32(_a_F_calc_rank_cd_3) - v1155
	goto L180
L179:
	;
	v1176 = v1155 - v1172
	goto L180
L180:
	;
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1169+v1176<<(uint(int32(1))%32)))))
	if (v1167^v1180)&int32(_a_F_calc_rank_cd_4) == int32(0) {
		goto L169
	} else {
		goto L181
	}
L181:
	;
	if v1175 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1188 = int32(_a_F_calc_rank_cd_4) - v1155
	goto L184
L183:
	;
	v1188 = v1155
	goto L184
L184:
	;
	v1189 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1169+v1188<<(uint(v1189)%32)))) = uint16(v1167)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	v1199 = v1193 + v1189
	goto L171
L185:
	;
	goto L168
L186:
	;
	if v1243 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1248 = v1082 - int32(12)
	if base.Ui32(v827) <= base.Ui32(v1248) {
		v1082 = v1248
		goto L162
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	goto L163
L190:
	;
	goto L160
L191:
	;
	goto L160
L192:
	;
	v1286 = v1082
	v1306 = v1282
	goto L195
L193:
	;
	v1345 = v1282
	goto L194
L194:
	;
	v1349 = v828 - v1082
	v1351 = base.I32_div_s(v1349, int32(12))
	v1357 = base.I32_div_s(v1349, int32(24))
	v1359 = v997 - (v1351 + v1252)
	if v1359 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1286)+8)))
	v1318 = *(*float64)(unsafe.Add(mBase, uint32(v27+int32(-48)+int32(base.Ui32(v1312)>>(uint(int32(11))%32))&int32(24))))
	v1319 = base.F64_add(v1306, v1318)
	v1321 = v1286 + int32(12)
	if base.Ui32(v1321) <= base.Ui32(v828) {
		v1286 = v1321
		v1306 = v1319
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v1345 = v1319
	goto L194
L197:
	;
	goto L196
L198:
	;
	v1362 = v1357
	goto L200
L199:
	;
	v1362 = v1359
	goto L200
L200:
	;
	v1370 = base.F64_mul(base.F64_convert_i32_u(v1252+v997), float64(0.5))
	v1372 = int32(0)
	if base.B2i32(base.F64_lt(v714, v1370) == v1372)|base.B2i32(v708 <= v1372) == v1372 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1383 = base.F64_add(v713, base.F64_div(float64(1), base.F64_sub(v1370, v714)))
	goto L203
L202:
	;
	v1383 = v713
	goto L203
L203:
	;
	v1386 = base.I32_div_s(v1082-v515, int32(12))
	v1387 = int32(1)
	v701 = v1386 + v1387
	v708 = v708 + v1387
	v712 = base.F64_add(v712, base.F64_div(base.F64_div(base.F64_convert_i32_s(v1351+int32(1)), v1345), base.F64_convert_i32_s(v1362+int32(1))))
	v713 = v1383
	v714 = v1370
	goto L108
L204:
	;
	goto L125
L205:
	;
	if l3&int32(2) == int32(0) {
		v1605 = v1511
		goto L217
	} else {
		goto L218
	}
L206:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1426 <= int32(0) {
		v1511 = v712
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1431 = v101 + v1426<<(uint(int32(2))%32)
	v1435 = v101
	v1437 = int32(0)
	goto L208
L208:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1435)))
	if v1459&int32(1) != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1489 = F_log(m, base.F64_convert_i32_s(v1482+int32(1)))
	mBase = m.M
	v1511 = base.F64_div(v712, v1489)
	goto L205
L210:
	;
	v1462 = int32(1)
	v1475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1431+(int32(base.Ui32(v1459)>>(uint(v1462)%32))&int32(2047)+int32(base.Ui32(v1459)>>(uint(int32(12))%32))+v1462)&int32(_a_F_calc_rank_cd_1)))))
	if base.Ui32(v1475) <= base.Ui32(v1462) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1481 = int32(1)
	goto L212
L212:
	;
	v1482 = v1481 + v1437
	v1484 = v1435 + int32(4)
	if base.Ui32(v1484) < base.Ui32(v1431) {
		v1435 = v1484
		v1437 = v1482
		goto L208
	} else {
		goto L216
	}
L213:
	;
	v1478 = v1462
	goto L215
L214:
	;
	v1478 = v1475
	goto L215
L215:
	;
	v1481 = v1478
	goto L212
L216:
	;
	goto L209
L217:
	;
	v1616 = int32(0)
	if base.B2i32(base.F64_gt(v713, float64(0)) == v1616)|(base.B2i32(l3&int32(4) == v1616)|base.B2i32(v708 <= v1616)) != 0 {
		goto L230
	} else {
		goto L231
	}
L218:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1521 <= int32(0) {
		v1605 = v1511
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1526 = v101 + v1521<<(uint(int32(2))%32)
	v1528 = int32(0)
	v1540 = v101
	goto L220
L220:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1540)))
	if v1554&int32(1) != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v1577 <= int32(0) {
		v1605 = v1511
		goto L217
	} else {
		goto L229
	}
L222:
	;
	v1557 = int32(1)
	v1570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526+(int32(base.Ui32(v1554)>>(uint(v1557)%32))&int32(2047)+int32(base.Ui32(v1554)>>(uint(int32(12))%32))+v1557)&int32(_a_F_calc_rank_cd_1)))))
	if base.Ui32(v1570) <= base.Ui32(v1557) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v1576 = int32(1)
	goto L224
L224:
	;
	v1577 = v1576 + v1528
	v1579 = v1540 + int32(4)
	if base.Ui32(v1579) < base.Ui32(v1526) {
		v1528 = v1577
		v1540 = v1579
		goto L220
	} else {
		goto L228
	}
L225:
	;
	v1573 = v1557
	goto L227
L226:
	;
	v1573 = v1570
	goto L227
L227:
	;
	v1576 = v1573
	goto L224
L228:
	;
	goto L221
L229:
	;
	v1605 = base.F64_div(v1511, base.F64_convert_i32_u(v1577))
	goto L217
L230:
	;
	v1626 = v1605
	goto L232
L231:
	;
	v1626 = base.F64_div(v1605, base.F64_div(base.F64_convert_i32_u(v708), v713))
	goto L232
L232:
	;
	if l3&int32(8) == int32(0) {
		v1637 = v1626
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if l3&int32(16) == int32(0) {
		v1653 = v1637
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1631 <= int32(0) {
		v1637 = v1626
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1637 = base.F64_div(v1626, base.F64_convert_i32_u(v1631))
	goto L233
L236:
	;
	if l3&int32(32) != 0 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1642 <= int32(0) {
		v1653 = v1637
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1648 = F_log(m, base.F64_convert_i32_s(v1642+int32(1)))
	mBase = m.M
	v1653 = base.F64_div(v1637, base.F64_div(v1648, float64(0.6931471805599453)))
	goto L236
L239:
	;
	v1659 = base.F64_div(v1653, base.F64_add(v1653, float64(1)))
	goto L241
L240:
	;
	v1659 = v1653
	goto L241
L241:
	;
	F_pfree(m, v515)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	F_pfree(m, v1662)
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	v1692 = base.F32_demote_f64(v1659)
	goto L23
L244:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	F_errmsg(m, int32(_a_F_calc_rank_cd_5), int32(0))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_calc_rank_cd_6), int32(876), int32(_a_F_calc_rank_cd_7))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L20
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cancel_prior_stmt_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[0]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[1]))
	v15 = v10 + v12*int32(20)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+10)))
	if v79 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v54 = int32(_a_F_cancel_prior_stmt_triggers_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2])) = v58
	v61 = F_palloc0(m, int32(36))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v27 = int32(0)
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22+v27<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 != l0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v44 = v27 + int32(1)
	if v19 != v44 {
		v27 = v44
		goto L5
	} else {
		goto L11
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 != l1 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	if v40 != int32(1) {
		v75 = v35
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	goto L6
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = l0
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v66 = F_lappend(m, v65, v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v66
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2])) = v55
	v75 = v61
	goto L1
L15:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+10)) = uint8(v163)
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v167
	return
L16:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v82 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v91 = v88
	v93 = v89
	goto L22
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	v88 = v82
	v89 = v83
	goto L17
L19:
	;
	goto L20
L20:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v85 == v84 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v88 = v85
	v89 = v84
	goto L17
L22:
	;
	if v93 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L15
L24:
	;
	v100 = v93
	goto L26
L25:
	;
	v100 = v91 + int32(16)
	goto L26
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if base.Ui32(v100) < base.Ui32(v101) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = v100
	goto L30
L28:
	;
	goto L29
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v154 != 0 {
		v91 = v154
		v93 = int32(0)
		goto L22
	} else {
		goto L41
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v114 = v106 + v111&int32(134217727)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v115 != l0 {
		goto L15
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if base.B2i32(v117&int32(3) != l2)|v117&int32(28) != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v111&int32(1073741823) | int32(-2147483648)
	v131 = v111 & int32(939524096)
	if v131 == int32(134217728) {
		v141 = int32(24)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v142 = v141 + v106
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if base.Ui32(v142) < base.Ui32(v143) {
		v106 = v142
		goto L30
	} else {
		goto L40
	}
L35:
	;
	if v131 != int32(805306368) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v131 == int32(268435456) {
		v141 = int32(12)
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v141 = int32(16)
	goto L34
L39:
	;
	v141 = int32(4)
	goto L34
L40:
	;
	goto L31
L41:
	;
	goto L23
}
func F_canonicalize_ec_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v7 = F_exprType(m, l0)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if l1 <= int32(3830) {
			if l1 <= int32(2775) {
				switch l1 - int32(2277) {
				case 0, 6:
					v37 = v7
					v39 = F_exprCollation(m, l0)
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == l2 {
							v52 = l0
							return v52
						} else {
							v42 = F_exprTypmod(m, l0)
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v37
								v46 = v42
								v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = v50
									return v52
								}
							}
						}
					}
				case 1, 2, 3, 4, 5:
					if l1 != v7 {
						v44 = l1
						v46 = int32(-1)
						v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = v50
							return v52
						}
					} else {
						v37 = l1
						v39 = F_exprCollation(m, l0)
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v52 = l0
								return v52
							} else {
								v42 = F_exprTypmod(m, l0)
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = v37
									v46 = v42
									v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = v50
										return v52
									}
								}
							}
						}
					}
				default:
					if l1 != int32(2249) {
						if l1 != v7 {
							v44 = l1
							v46 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v37 = l1
							v39 = F_exprCollation(m, l0)
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 == l2 {
									v52 = l0
									return v52
								} else {
									v42 = F_exprTypmod(m, l0)
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v44 = v37
										v46 = v42
										v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = v50
											return v52
										}
									}
								}
							}
						}
					} else {
						v37 = v7
						v39 = F_exprCollation(m, l0)
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v52 = l0
								return v52
							} else {
								v42 = F_exprTypmod(m, l0)
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = v37
									v46 = v42
									v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = v50
										return v52
									}
								}
							}
						}
					}
				}
			} else {
				if l1 == int32(2776) {
					v37 = v7
					v39 = F_exprCollation(m, l0)
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == l2 {
							v52 = l0
							return v52
						} else {
							v42 = F_exprTypmod(m, l0)
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v37
								v46 = v42
								v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = v50
									return v52
								}
							}
						}
					}
				} else {
					if l1 != int32(3500) {
						if l1 != v7 {
							v44 = l1
							v46 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v37 = l1
							v39 = F_exprCollation(m, l0)
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 == l2 {
									v52 = l0
									return v52
								} else {
									v42 = F_exprTypmod(m, l0)
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v44 = v37
										v46 = v42
										v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = v50
											return v52
										}
									}
								}
							}
						}
					} else {
						v37 = v7
						v39 = F_exprCollation(m, l0)
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v52 = l0
								return v52
							} else {
								v42 = F_exprTypmod(m, l0)
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = v37
									v46 = v42
									v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = v50
										return v52
									}
								}
							}
						}
					}
				}
			}
		} else {
			if base.B2i32(base.Ui32(l1-int32(_a_F_canonicalize_ec_expression_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(_a_F_canonicalize_ec_expression_1)) < base.Ui32(int32(2)))|base.B2i32(l1 == int32(3831)) != 0 {
				v37 = v7
				v39 = F_exprCollation(m, l0)
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 == l2 {
						v52 = l0
						return v52
					} else {
						v42 = F_exprTypmod(m, l0)
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = v37
							v46 = v42
							v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						}
					}
				}
			} else {
				if l1 != v7 {
					v44 = l1
					v46 = int32(-1)
					v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = v50
						return v52
					}
				} else {
					v37 = l1
					v39 = F_exprCollation(m, l0)
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == l2 {
							v52 = l0
							return v52
						} else {
							v42 = F_exprTypmod(m, l0)
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v37
								v46 = v42
								v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = v50
									return v52
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = F_find_duplicate_ors(m, l0, l1)
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_casemap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	if base.Ui32(int32(_a_F_casemap_0)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v23 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_casemap[0])))
	v24 = int32(8)
	v25 = int32(base.Ui32(l0) >> (uint(v24) % 32))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_casemap[1]))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v26*int32(86))+uint32(_c_F_casemap[1]))))
	v35 = base.I32_rem_u_s(int32(base.Ui32(v23*v30)>>(uint(int32(11))%32)), int32(6))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_casemap[2]))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(v21)%32)+v38<<(uint(v21)%32))+uint32(_c_F_casemap[3])))
	v44 = v42 >> (uint(v24) % 32)
	v46 = v42 & v12
	if base.Ui32(v46) <= base.Ui32(int32(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v44&(int32(0)-(l1^v46)) + l0
L4:
	;
	goto L5
L5:
	;
	v56 = v44 & int32(255)
	if v56 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v63 = int32(base.Ui32(v44) >> (uint(int32(8)) % 32))
	v64 = v56
	goto L7
L7:
	;
	v70 = int32(1)
	v71 = int32(base.Ui32(v64) >> (uint(v70) % 32))
	v72 = v71 + v63
	v74 = v72 << (uint(v70) % 32)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_casemap[4]))))
	if v75 == v13 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_casemap[5]))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_c_F_casemap[3])))
	v84 = v82 & int32(255)
	if base.Ui32(v84) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v100 = base.B2i32(base.Ui32(v13) < base.Ui32(v75))
	if base.Ui32(v13) < base.Ui32(v75) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	return (int32(0)-(l1^v84))&(v82>>(uint(int32(8))%32)) + l0
L13:
	;
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v97 = int32(-1)
	goto L17
L16:
	;
	v97 = int32(1)
	goto L17
L17:
	;
	return v97 + l0
L18:
	;
	v101 = v63
	goto L20
L19:
	;
	v101 = v72
	goto L20
L20:
	;
	if base.Ui32(v13) < base.Ui32(v75) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v103 = v71
	goto L23
L22:
	;
	v103 = v64 - v71
	goto L23
L23:
	;
	if v103 != 0 {
		v63 = v101
		v64 = v103
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L8
}
func F_cashlarger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v7 < v5 {
		v9 = v5
	} else {
		v9 = v7
	}
	v10 = F_Int64GetDatum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_cdissect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
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
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v705 int32
	_ = v705
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v808 int32
	_ = v808
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v836 int32
	_ = v836
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v948 int32
	_ = v948
	v5 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_cdissect[0]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = m.T0[v22].(func(*base.Module) int32)(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(19)
L8:
	;
	goto L9
L9:
	;
	v27 = int32(15)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v28 - int32(40) {
	case 0:
		goto L14
	case 1, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		v948 = v27
		goto L10
	case 2:
		goto L15
	case 6:
		goto L17
	case 21:
		v911 = v5
		goto L11
	default:
		goto L18
	}
L10:
	;
	return v948
L11:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v920 <= int32(0) {
		goto L297
	} else {
		goto L298
	}
L12:
	;
	F_pfree(m, v415)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L4
	} else {
		goto L296
	}
L13:
	;
	v636 = (l3 - l2) >> (uint(int32(2)) % 32)
	v637 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v636) < base.Ui32(v637) {
		goto L225
	} else {
		goto L226
	}
L14:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v631 = F_cdissect(m, l0, v630, l2, l3)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L224
	}
L15:
	;
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+1)))
	if v389&int32(2) == int32(0) {
		goto L13
	} else {
		goto L150
	}
L16:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v322 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = int32(2)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+v99<<(uint(v100)%32))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v105&v100 != 0 {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	if v28 == int32(124) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if v28 != int32(98) {
		v948 = v27
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v35 = int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = v36 + v37<<(uint(int32(3))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 == int32(-1) {
		v911 = v35
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v41 == v46 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v911 = base.B2i32(l2 != l3) | base.B2i32(v44 < v45)
	goto L11
L23:
	;
	goto L24
L24:
	;
	if l2 == l3 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v911 = base.B2i32(v45 != int32(0))
	goto L11
L26:
	;
	goto L27
L27:
	;
	v56 = (l3 - l2) >> (uint(int32(2)) % 32)
	v57 = v46 - v41
	v58 = base.I32_div_u_s(v56, v57)
	if v56-v58*v57|base.B2i32(base.Ui32(v58) < base.Ui32(v45)) != 0 {
		v911 = v35
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(v44 != int32(256))&base.B2i32(base.Ui32(v44) < base.Ui32(v58)) != 0 {
		v911 = v35
		goto L11
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v56) < base.Ui32(v57) {
		v911 = int32(0)
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v70 = int32(2)
	v80 = l2
	v81 = v58
	goto L31
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+420))
	v90 = m.T0[v89].(func(*base.Module, int32, int32, int32) int32)(m, v63+v41<<(uint(v70)%32), v80, v57)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	v911 = v93
	goto L11
L33:
	;
	v93 = base.B2i32(v90 != int32(0))
	if v90 != 0 {
		v911 = v93
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v96 = v81 - int32(1)
	if v96 != 0 {
		v80 = v80 + v57<<(uint(v70)%32)
		v81 = v96
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v103 != 0 {
		v135 = v103
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v103 != 0 {
		v243 = v103
		goto L84
	} else {
		goto L85
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v136 != 0 {
		v948 = v136
		goto L10
	} else {
		goto L46
	}
L40:
	;
	v108 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = F_newdfa(m, l0, v98+int32(36), v111+int32(72), v108)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v115 == int32(0) {
		v135 = v108
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v119 == int32(98) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+60)) = v122
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+64)) = uint16(v124)
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+66)) = uint16(v126)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v128+v129<<(uint(int32(2))%32)))) = v115
	v135 = v115
	goto L39
L46:
	;
	v137 = F_getsubdfa(m, l0, v104)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v139 != 0 {
		v948 = v139
		goto L10
	} else {
		goto L48
	}
L48:
	;
	v140 = int32(0)
	v142 = F_shortest(m, l0, v135, l2, l2, l3, v140, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v144 != 0 {
		v948 = v144
		goto L10
	} else {
		goto L50
	}
L50:
	;
	if v142 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return int32(1)
L52:
	;
	goto L53
L53:
	;
	v156 = v142
	goto L54
L54:
	;
	v163 = F_longest(m, l0, v137, v156, l3, int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	v948 = int32(1)
	goto L10
L56:
	;
	if v163 == l3 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v166 = F_cdissect(m, l0, v98, l2, v156)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v205 != 0 {
		v948 = v205
		goto L10
	} else {
		goto L79
	}
L60:
	;
	if v166 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v170 = F_cdissect(m, l0, v104, v156, l3)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	v201 = v166
	goto L63
L63:
	;
	if v201 != int32(1) {
		v948 = v201
		goto L10
	} else {
		goto L78
	}
L64:
	;
	if v170 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v911 = int32(0)
	goto L11
L66:
	;
	goto L67
L67:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v176 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v201 = v170
	goto L63
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v192 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v179) <= base.Ui32(v176) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v182 = v176 << (uint(int32(3)) % 32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v185 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v182+v183))) = v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v187+v182)+4)) = v185
	goto L69
L72:
	;
	v194 = v192
	goto L75
L73:
	;
	goto L74
L74:
	;
	goto L68
L75:
	;
	F_zaptreesubs(m, l0, v194)
	mBase = m.M
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	if v197 != 0 {
		v194 = v197
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L74
L77:
	;
	goto L76
L78:
	;
	goto L59
L79:
	;
	if l3 == v156 {
		v948 = int32(1)
		goto L10
	} else {
		goto L80
	}
L80:
	;
	v210 = int32(0)
	v212 = F_shortest(m, l0, v135, l2, v156+int32(4), l3, v210, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v214 != 0 {
		v948 = v214
		goto L10
	} else {
		goto L82
	}
L82:
	;
	if v212 != 0 {
		v156 = v212
		goto L54
	} else {
		goto L83
	}
L83:
	;
	goto L55
L84:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v244 != 0 {
		v948 = v244
		goto L10
	} else {
		goto L91
	}
L85:
	;
	v216 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v223 = F_newdfa(m, l0, v98+int32(36), v219+int32(72), v216)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v223 == int32(0) {
		v243 = v216
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v227 == int32(98) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+60)) = v230
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v223)+64)) = uint16(v232)
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v223)+66)) = uint16(v234)
	goto L90
L89:
	;
	goto L90
L90:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v237<<(uint(int32(2))%32)))) = v223
	v243 = v223
	goto L84
L91:
	;
	v245 = F_getsubdfa(m, l0, v104)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v247 != 0 {
		v948 = v247
		goto L10
	} else {
		goto L93
	}
L93:
	;
	v249 = F_longest(m, l0, v243, l2, l3, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v251 != 0 {
		v948 = v251
		goto L10
	} else {
		goto L95
	}
L95:
	;
	if v249 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	return int32(1)
L97:
	;
	goto L98
L98:
	;
	v263 = v249
	goto L99
L99:
	;
	v270 = F_longest(m, l0, v245, v263, l3, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L101
	}
L100:
	;
	v948 = int32(1)
	goto L10
L101:
	;
	if v270 == l3 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v273 = F_cdissect(m, l0, v98, l2, v263)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v312 != 0 {
		v948 = v312
		goto L10
	} else {
		goto L124
	}
L105:
	;
	if v273 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v277 = F_cdissect(m, l0, v104, v263, l3)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	v308 = v273
	goto L108
L108:
	;
	if v308 != int32(1) {
		v948 = v308
		goto L10
	} else {
		goto L123
	}
L109:
	;
	if v277 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v911 = int32(0)
	goto L11
L111:
	;
	goto L112
L112:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v283 <= int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v308 = v277
	goto L108
L114:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v299 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v286) <= base.Ui32(v283) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v289 = v283 << (uint(int32(3)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v292 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v289+v290))) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v294+v289)+4)) = v292
	goto L114
L117:
	;
	v301 = v299
	goto L120
L118:
	;
	goto L119
L119:
	;
	goto L113
L120:
	;
	F_zaptreesubs(m, l0, v301)
	mBase = m.M
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+24))
	if v304 != 0 {
		v301 = v304
		goto L120
	} else {
		goto L122
	}
L121:
	;
	goto L119
L122:
	;
	goto L121
L123:
	;
	goto L104
L124:
	;
	if l2 == v263 {
		v948 = int32(1)
		goto L10
	} else {
		goto L125
	}
L125:
	;
	v318 = F_longest(m, l0, v243, l2, v263-int32(4), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v320 != 0 {
		v948 = v320
		goto L10
	} else {
		goto L127
	}
L127:
	;
	if v318 != 0 {
		v263 = v318
		goto L99
	} else {
		goto L128
	}
L128:
	;
	goto L100
L129:
	;
	return int32(1)
L130:
	;
	goto L131
L131:
	;
	v333 = v322
	goto L132
L132:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341<<(uint(int32(2))%32))))
	if v345 != 0 {
		v372 = v345
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v948 = int32(1)
	goto L10
L134:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v374 != 0 {
		v948 = v374
		goto L10
	} else {
		goto L141
	}
L135:
	;
	v346 = int32(0)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v353 = F_newdfa(m, l0, v333+int32(36), v349+int32(72), v346)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	if v353 == int32(0) {
		v372 = v346
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v357 == int32(98) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+60)) = v360
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v353)+64)) = uint16(v362)
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v353)+66)) = uint16(v364)
	goto L140
L139:
	;
	goto L140
L140:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v366+v367<<(uint(int32(2))%32)))) = v353
	v372 = v353
	goto L134
L141:
	;
	v376 = F_longest(m, l0, v372, l2, l3, int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	if v376 == l3 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v379 = F_cdissect(m, l0, v333, l2, l3)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v384 != 0 {
		v948 = v384
		goto L10
	} else {
		goto L148
	}
L146:
	;
	if v379 != int32(1) {
		v911 = v379
		goto L11
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v333)+24))
	if v386 != 0 {
		v333 = v386
		goto L132
	} else {
		goto L149
	}
L149:
	;
	goto L133
L150:
	;
	if v387 <= int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if l2 == l3 {
		v911 = v5
		goto L11
	} else {
		goto L154
	}
L152:
	;
	v398 = v387
	goto L153
L153:
	;
	v401 = (l3 - l2) >> (uint(int32(2)) % 32)
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v401) < base.Ui32(v402) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v398 = int32(1)
	goto L153
L155:
	;
	v404 = v401
	goto L157
L156:
	;
	v404 = v402
	goto L157
L157:
	;
	if v402 == int32(256) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v407 = v401
	goto L160
L159:
	;
	v407 = v404
	goto L160
L160:
	;
	if base.Ui32(v398) < base.Ui32(v407) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v409 = v407
	goto L163
L162:
	;
	v409 = v398
	goto L163
L163:
	;
	v410 = int32(2)
	v415 = F_palloc_extended(m, v409<<(uint(v410)%32)+int32(4), v410)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	if v415 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	return int32(12)
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415))) = l2
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v423 = F_getsubdfa(m, l0, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v425 != 0 {
		goto L12
	} else {
		goto L169
	}
L169:
	;
	v431 = l2
	v432 = int32(1)
	v434 = v5
	goto L170
L170:
	;
	v445 = int32(2)
	v451 = v432 - int32(1)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v415+v451<<(uint(v445)%32))))
	if base.B2i32(base.Ui32(v432) < base.Ui32(v398))&base.B2i32((l3-v431)>>(uint(v445)%32) <= v398-v432)|(base.B2i32(l3 == v431)|base.B2i32(v455 != v431)) != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	F_pfree(m, v415)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L4
	} else {
		goto L223
	}
L172:
	;
	v459 = v431
	goto L174
L173:
	;
	v459 = v431 + int32(4)
	goto L174
L174:
	;
	if base.Ui32(v432) < base.Ui32(v409) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v464 = v459
	goto L177
L176:
	;
	v464 = l3
	goto L177
L177:
	;
	v465 = int32(0)
	v467 = F_shortest(m, l0, v423, v455, v464, l3, v465, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415+v432<<(uint(int32(2))%32)))) = v467
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v470 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	if v467 == int32(0) {
		v544 = v434
		goto L184
	} else {
		goto L185
	}
L180:
	;
	goto L171
L181:
	;
	if int32(0) < v603 {
		v431 = v602
		v432 = v603
		v434 = v605
		goto L170
	} else {
		goto L222
	}
L182:
	;
	v579 = v566
	goto L216
L183:
	;
	if v558 <= int32(0) {
		goto L180
	} else {
		goto L215
	}
L184:
	;
	v551 = v544
	v558 = v451
	goto L183
L185:
	;
	if v434 < v451 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v474 = v434
	goto L188
L187:
	;
	v474 = v451
	goto L188
L188:
	;
	if l3 != v467 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if base.Ui32(v409) <= base.Ui32(v432) {
		v544 = v474
		goto L184
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	if base.Ui32(v432) < base.Ui32(v398) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v602 = v467
	v603 = v432 + int32(1)
	v605 = v474
	goto L181
L193:
	;
	v566 = v432
	v567 = v474
	goto L182
L194:
	;
	goto L195
L195:
	;
	v487 = v474
	goto L197
L196:
	;
	F_pfree(m, v415)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L214
	}
L197:
	;
	v494 = v487 + int32(1)
	if v432 < v494 {
		goto L196
	} else {
		goto L199
	}
L198:
	;
	if v532 == int32(1) {
		v551 = v487
		v558 = v494
		goto L183
	} else {
		goto L212
	}
L199:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	if v498 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v524 = int32(2)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v415+v487<<(uint(v524)%32))))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v415+v494<<(uint(v524)%32))))
	v532 = F_cdissect(m, l0, v523, v527, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L210
	}
L201:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v496)+20))
	if v514 != 0 {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v501) <= base.Ui32(v498) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v504 = v498 << (uint(int32(3)) % 32)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v507 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v504+v505))) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v509+v504)+4)) = v507
	goto L201
L204:
	;
	v516 = v514
	goto L207
L205:
	;
	goto L206
L206:
	;
	goto L200
L207:
	;
	F_zaptreesubs(m, l0, v516)
	mBase = m.M
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v516)+24))
	if v519 != 0 {
		v516 = v519
		goto L207
	} else {
		goto L209
	}
L208:
	;
	goto L206
L209:
	;
	goto L208
L210:
	;
	if v532 == int32(0) {
		v487 = v494
		goto L197
	} else {
		goto L211
	}
L211:
	;
	goto L198
L212:
	;
	F_pfree(m, v415)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	return v532
L214:
	;
	v911 = int32(0)
	goto L11
L215:
	;
	v566 = v558
	v567 = v551
	goto L182
L216:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v415+v579<<(uint(int32(2))%32))))
	if base.Ui32(v590) < base.Ui32(l3) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L180
L218:
	;
	v602 = v590 + int32(4)
	v603 = v579
	v605 = v567
	goto L181
L219:
	;
	goto L220
L220:
	;
	v594 = int32(1)
	if v594 < v579 {
		v579 = v579 - v594
		goto L216
	} else {
		goto L221
	}
L221:
	;
	goto L217
L222:
	;
	goto L180
L223:
	;
	return int32(1)
L224:
	;
	v911 = v631
	goto L11
L225:
	;
	v639 = v636
	goto L227
L226:
	;
	v639 = v637
	goto L227
L227:
	;
	if v637 == int32(256) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v642 = v636
	goto L230
L229:
	;
	v642 = v639
	goto L230
L230:
	;
	v643 = int32(1)
	if v387 <= v643 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v646 = v643
	goto L233
L232:
	;
	v646 = v387
	goto L233
L233:
	;
	if base.Ui32(v646) < base.Ui32(v642) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v648 = v642
	goto L236
L235:
	;
	v648 = v646
	goto L236
L236:
	;
	v649 = int32(2)
	v654 = F_palloc_extended(m, v648<<(uint(v649)%32)+int32(4), v649)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	if v654 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	return int32(12)
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v654))) = l2
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v662 = F_getsubdfa(m, l0, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v664 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	F_pfree(m, v654)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L4
	} else {
		goto L295
	}
L243:
	;
	v670 = int32(1)
	v671 = v5
	v673 = l3
	goto L244
L244:
	;
	v678 = int32(2)
	v682 = v670 - int32(1)
	v685 = v654 + v682<<(uint(v678)%32)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	v688 = F_longest(m, l0, v662, v686, v673, int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L246
	}
L245:
	;
	F_pfree(m, v654)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L4
	} else {
		goto L294
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v654+v670<<(uint(v678)%32)))) = v688
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v691 != 0 {
		goto L242
	} else {
		goto L247
	}
L247:
	;
	if v688 == int32(0) {
		v775 = v671
		goto L252
	} else {
		goto L253
	}
L248:
	;
	goto L245
L249:
	;
	if int32(0) < v845 {
		v670 = v845
		v671 = v846
		v673 = v848
		goto L244
	} else {
		goto L293
	}
L250:
	;
	v808 = v795
	goto L284
L251:
	;
	if v780 <= int32(0) {
		goto L248
	} else {
		goto L283
	}
L252:
	;
	v780 = v682
	v783 = v775
	goto L251
L253:
	;
	if v671 < v682 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v695 = v671
	goto L256
L255:
	;
	v695 = v682
	goto L256
L256:
	;
	if l3 != v688 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v795 = v670
	v798 = v695
	goto L250
L258:
	;
	if base.Ui32(v648) <= base.Ui32(v670) {
		v775 = v695
		goto L252
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	if base.Ui32(v670) < base.Ui32(v646) {
		goto L257
	} else {
		goto L263
	}
L261:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	if (base.B2i32(v646-v670 < (l3-v688)>>(uint(int32(2))%32))|base.B2i32(base.Ui32(v646) <= base.Ui32(v670)))&base.B2i32(v705 == v688) != 0 {
		goto L257
	} else {
		goto L262
	}
L262:
	;
	v845 = v670 + int32(1)
	v846 = v695
	v848 = l3
	goto L249
L263:
	;
	v717 = v695
	goto L265
L264:
	;
	F_pfree(m, v654)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L282
	}
L265:
	;
	v725 = v717 + int32(1)
	if v670 < v725 {
		goto L264
	} else {
		goto L267
	}
L266:
	;
	if v763 == int32(1) {
		v780 = v725
		v783 = v717
		goto L251
	} else {
		goto L280
	}
L267:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	if v729 <= int32(0) {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v755 = int32(2)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v654+v717<<(uint(v755)%32))))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v654+v725<<(uint(v755)%32))))
	v763 = F_cdissect(m, l0, v754, v758, v762)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L4
	} else {
		goto L278
	}
L269:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v727)+20))
	if v745 != 0 {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v732) <= base.Ui32(v729) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v735 = v729 << (uint(int32(3)) % 32)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v738 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v735+v736))) = v738
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v735)+4)) = v738
	goto L269
L272:
	;
	v747 = v745
	goto L275
L273:
	;
	goto L274
L274:
	;
	goto L268
L275:
	;
	F_zaptreesubs(m, l0, v747)
	mBase = m.M
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v747)+24))
	if v750 != 0 {
		v747 = v750
		goto L275
	} else {
		goto L277
	}
L276:
	;
	goto L274
L277:
	;
	goto L276
L278:
	;
	if v763 == int32(0) {
		v717 = v725
		goto L265
	} else {
		goto L279
	}
L279:
	;
	goto L266
L280:
	;
	F_pfree(m, v654)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	return v763
L282:
	;
	v911 = int32(0)
	goto L11
L283:
	;
	v795 = v780
	v798 = v783
	goto L250
L284:
	;
	v819 = v654 + v808<<(uint(int32(2))%32)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819-int32(4))))
	if base.Ui32(v820) <= base.Ui32(v823) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	goto L248
L286:
	;
	v836 = int32(1)
	if v836 < v808 {
		v808 = v808 - v836
		goto L284
	} else {
		goto L292
	}
L287:
	;
	v826 = v820 - int32(4)
	if base.Ui32(v823) < base.Ui32(v826) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v845 = v808
	v846 = v798
	v848 = v826
	goto L249
L289:
	;
	goto L290
L290:
	;
	if base.B2i32(v646-v808 < (l3-v823)>>(uint(int32(2))%32))|base.B2i32(base.Ui32(v646) <= base.Ui32(v808)) != 0 {
		goto L286
	} else {
		goto L291
	}
L291:
	;
	v845 = v808
	v846 = v798
	v848 = v826
	goto L249
L292:
	;
	goto L285
L293:
	;
	goto L248
L294:
	;
	v871 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v911 = base.B2i32(l2 != l3) | base.B2i32(v871 != int32(0))
	goto L11
L295:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v911 = v890
	goto L11
L296:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v911 = v906
	goto L11
L297:
	;
	return v911
L298:
	;
	goto L299
L299:
	;
	if v911 != 0 {
		v948 = v911
		goto L10
	} else {
		goto L300
	}
L300:
	;
	v924 = int32(0)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v925) <= base.Ui32(v920) {
		v948 = v924
		goto L10
	} else {
		goto L301
	}
L301:
	;
	v928 = v920 << (uint(int32(3)) % 32)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v933 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v928+v929))) = (l2 - v931) >> (uint(v933) % 32)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v936+v928)+4)) = (l3 - v938) >> (uint(v933) % 32)
	v948 = v924
	goto L10
}
func F_charge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(base.Ui32(v3) <= base.Ui32(v2))
}
func F_charlt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(base.Ui32(v2) < base.Ui32(v3))
}
func F_charout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(5))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = base.I32_extend8_s(v4)
		if v10 < int32(0) {
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v13)
			v15 = int32(7)
			v17 = int32(48)
			v18 = v4&v15 | v17
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v18)
			v25 = int32(base.Ui32(v4)>>(uint(int32(3))%32))&v15 | v17
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v25)
			v35 = int32(92)
			v36 = int32(base.Ui32(v4&int32(192))>>(uint(int32(6))%32)) | v17
		} else {
			v35 = v10
			v36 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v36)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v35)
		return v6
	}
}
func F_check_enable_rls(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == v4 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_check_enable_rls[0]))
		v17 = v16
	} else {
		v17 = l1
	}
	if base.Ui32(l0) < base.Ui32(int32(_a_F_check_enable_rls_0)) {
		v82 = v4
		m.G0 = v11 + int32(16)
		return v82
	} else {
		v21 = F_SearchSysCache1(m, int32(57), l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				v82 = v4
				m.G0 = v11 + int32(16)
				return v82
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
				v29 = v27 + v28
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+128)))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)))
				F_ReleaseCatCache(m, v21)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v31 != int32(1) {
						v82 = v4
						m.G0 = v11 + int32(16)
						return v82
					} else {
						v36 = int32(1)
						v37 = F_has_bypassrls_privilege(m, v17)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 != 0 {
								v82 = v36
								m.G0 = v11 + int32(16)
								return v82
							} else {
								v40 = F_object_ownercheck(m, int32(1259), l0, v17)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									if v40 != 0 {
										if v30&int32(1) == int32(0) {
											v82 = v36
											m.G0 = v11 + int32(16)
											return v82
										} else {
											v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_enable_rls[1])))
											if int32(base.Ui32(v47&int32(4))>>(uint(int32(2))%32)) != 0 {
												v82 = v36
												m.G0 = v11 + int32(16)
												return v82
											} else {
												v52 = int32(2)
												if l2 != 0 {
													v82 = v52
													m.G0 = v11 + int32(16)
													return v82
												} else {
													v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_enable_rls[2])))
													if v54&int32(1) != 0 {
														v82 = v52
														m.G0 = v11 + int32(16)
														return v82
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16797828))
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return int32(0)
															} else {
																v64 = F_get_rel_name(m, l0)
																mBase = m.M
																v65 = m.ExcPending
																if v65 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
																	F_errmsg(m, int32(_a_F_check_enable_rls_1), v11)
																	mBase = m.M
																	v69 = m.ExcPending
																	if v69 != 0 {
																		return int32(0)
																	} else {
																		if v40 != 0 {
																			F_errhint(m, int32(_a_F_check_enable_rls_2), int32(0))
																			mBase = m.M
																			v73 = m.ExcPending
																			if v73 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_check_enable_rls_3), int32(129), int32(_a_F_check_enable_rls_4))
																				mBase = m.M
																				v78 = m.ExcPending
																				if v78 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		} else {
																			F_errfinish(m, int32(_a_F_check_enable_rls_3), int32(129), int32(_a_F_check_enable_rls_4))
																			mBase = m.M
																			v78 = m.ExcPending
																			if v78 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
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
									} else {
										v52 = int32(2)
										if l2 != 0 {
											v82 = v52
											m.G0 = v11 + int32(16)
											return v82
										} else {
											v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_enable_rls[2])))
											if v54&int32(1) != 0 {
												v82 = v52
												m.G0 = v11 + int32(16)
												return v82
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16797828))
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int32(0)
													} else {
														v64 = F_get_rel_name(m, l0)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
															F_errmsg(m, int32(_a_F_check_enable_rls_1), v11)
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return int32(0)
															} else {
																if v40 != 0 {
																	F_errhint(m, int32(_a_F_check_enable_rls_2), int32(0))
																	mBase = m.M
																	v73 = m.ExcPending
																	if v73 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_check_enable_rls_3), int32(129), int32(_a_F_check_enable_rls_4))
																		mBase = m.M
																		v78 = m.ExcPending
																		if v78 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																} else {
																	F_errfinish(m, int32(_a_F_check_enable_rls_3), int32(129), int32(_a_F_check_enable_rls_4))
																	mBase = m.M
																	v78 = m.ExcPending
																	if v78 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
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
						}
					}
				}
			}
		}
	}
}
func F_check_io_max_concurrency(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_io_max_concurrency[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_io_max_concurrency[1])) = v8
		v14 = F_format_elog_string(m, int32(_a_F_check_io_max_concurrency_0), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_io_max_concurrency[2])) = v14
			return base.B2i32(v4 != int32(0))
		}
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_check_labels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_check_labels_0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L21
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_check_labels_0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	if l0 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v8 + int32(32)
	return
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v14 == int32(0))|base.B2i32(v14 != v17) != 0 {
		v35 = v14
		v36 = v17
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v35-v36 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v20 = l0
	v21 = l1
	goto L10
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v25
		v36 = v24
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v35 = v25
	v36 = v24
	goto L8
L12:
	;
	v28 = int32(1)
	if v25 == v24 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L5
L15:
	;
	return
L16:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	F_errmsg(m, int32(_a_F_check_labels_1), v8)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v52 = F_plpgsql_scanner_errposition(m, l2, l3)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_check_labels_2), int32(3886), int32(_a_F_check_labels_3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F_errmsg(m, int32(_a_F_check_labels_4), v8+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v73 = F_plpgsql_scanner_errposition(m, l2, l3)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_check_labels_2), int32(3893), int32(_a_F_check_labels_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_memoizable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v7 == int32(0) {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			if v10 != int32(17) {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
				if v13 == int32(0) {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					if v16 != int32(2) {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						v21 = F_exprType(m, v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v24 = F_lookup_type_cache(m, v21, int32(17))
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
								if v26 == int32(0) {
								} else {
									v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
									if v29 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v29
									}
								}
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
								v37 = F_exprType(m, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									if v37 != v21 {
										v41 = F_lookup_type_cache(m, v37, int32(17))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											v43 = v41
											v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
											if v44 == int32(0) {
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
												if v47 == int32(0) {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v47
												}
											}
											return
										}
									} else {
										v43 = v24
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
										if v44 == int32(0) {
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
											if v47 == int32(0) {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v47
											}
										}
										return
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
func F_check_notify_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_notify_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_check_publications(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(25)
	v11 = F_makeStringInfo(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_appendStringInfoString(m, v11, int32(_a_F_check_publications_0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_GetPublicationsStr(m, l1, v11, int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v11, int32(41))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	v29 = m.T0[v28].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v22, int32(1), v7+int32(28))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_free_attrmap(m, v11)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v33 == int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v36 = F_list_copy(m, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v40 = F_MakeTupleTableSlot(m, v38, int32(_a_F_check_publications_1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v45 = F_tuplestore_gettupleslot(m, v42, int32(1), int32(0), v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = v36
	goto L17
L15:
	;
	v75 = v36
	goto L16
L16:
	;
	F_ExecDropSingleTupleTableSlot(m, v40)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+6)))
	if v51 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v75 = v63
	goto L16
L19:
	;
	F_slot_getsomeattrs_int(m, v40, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = F_text_to_cstring(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v61 = F_makeString(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v63 = F_list_delete(m, v48, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	m.T0[v66].(func(*base.Module, int32))(m, v40)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v72 = F_tuplestore_gettupleslot(m, v69, int32(1), int32(0), v40)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v72 != 0 {
		v48 = v63
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v80 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_pfree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v83 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	F_tuplestore_end(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v86 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	F_FreeTupleDesc(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v29)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v75 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	m.G0 = v7 + int32(32)
	return
L44:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v93 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v96 = F_makeStringInfo(m)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_GetPublicationsStr(m, v75, v96, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v103 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v103 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v111
	F_errmsg_plural(m, int32(_a_F_check_publications_2), int32(_a_F_check_publications_3), v110, v7)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_check_publications_4), int32(501), int32(_a_F_check_publications_5))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L43
L53:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v131
	F_errmsg(m, int32(_a_F_check_publications_6), v7+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_check_publications_4), int32(467), int32(_a_F_check_publications_5))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_rolespec_name(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	if l0 == int32(0) {
		m.G0 = v5 + int32(32)
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v9 != 0 {
			m.G0 = v5 + int32(32)
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v11 = int32(0)
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			if v12 != int32(112) {
				v21 = v11
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
				if v15 != int32(103) {
					v21 = v11
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
					v21 = base.B2i32(v18 == int32(95))
				}
			}
			if v21 == int32(0) {
				m.G0 = v5 + int32(32)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errcode(m, int32(151818372))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v31
						F_errmsg(m, int32(_a_F_check_rolespec_name_0), v5+int32(16))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_check_rolespec_name_1)
							F_errdetail_internal(m, int32(_a_F_check_rolespec_name_2), v5)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_rolespec_name_3), int32(_a_F_check_rolespec_name_4), int32(_a_F_check_rolespec_name_5))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
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
			}
		}
	}
}
func F_check_usermap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
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
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
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
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v386 int32
	_ = v386
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(208)
	return v386
L2:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_check_usermap[0]))
	if v66 == int32(0) {
		v345 = v4
		v348 = v4
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v20 == int32(0))|base.B2i32(v20 != v23) != 0 {
		v41 = v20
		v42 = v23
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	if v41-v42 == int32(0) {
		v386 = v4
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v26 = l1
	v27 = l2
	goto L10
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v31
		v42 = v30
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v41 = v31
	v42 = v30
	goto L8
L12:
	;
	v34 = int32(1)
	if v31 == v30 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v46 = int32(-1)
	v49 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v49 == int32(0) {
		v386 = v46
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg(m, int32(_a_F_check_usermap_0), v15)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2988), int32(_a_F_check_usermap_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v386 = v46
	goto L1
L20:
	;
	if v345|v348 != 0 {
		goto L98
	} else {
		goto L99
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v69 <= int32(0) {
		v345 = v4
		v348 = v4
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v80 = v4
	goto L23
L23:
	;
	v84 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v80<<(uint(int32(2))%32))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v94 == v84)|base.B2i32(v94 != v97) != 0 {
		v115 = v94
		v116 = v97
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v340 = int32(0)
	v345 = v340
	v348 = v340
	goto L20
L25:
	;
	if v329|v331 != 0 {
		v345 = v329
		v348 = v331
		goto L20
	} else {
		goto L96
	}
L26:
	;
	if v115-v116 != 0 {
		v329 = v84
		v331 = v84
		goto L25
	} else {
		goto L33
	}
L27:
	;
	goto L26
L28:
	;
	v100 = v91
	v101 = l0
	goto L29
L29:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v105
		v116 = v104
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v115 = v105
	v116 = v104
	goto L27
L31:
	;
	v108 = int32(1)
	if v105 == v104 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v119 = F_get_role_oid(m, l1, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v122 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v123 = F_strlen(m, l2)
	mBase = m.M
	v128 = F_palloc(m, v123<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L15
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v295 == int32(0))|base.B2i32(v295 != v298) != 0 {
		v316 = v295
		v317 = v298
		goto L87
	} else {
		goto L88
	}
L38:
	;
	v130 = F_strlen(m, l2)
	mBase = m.M
	v131 = F_pg_mb2wchar_with_len(m, l2, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v138 = F_pg_regexec(m, v133, v128, v131, int32(0), int32(2), v15+int32(192))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	F_pfree(m, v128)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	if v138 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v138 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
	if v180 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v329 = int32(0)
	v331 = base.B2i32(v138 != int32(1))
	goto L25
L46:
	;
	v147 = v15 + int32(80)
	F_pg_regerror(m, v138, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v152 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	if v152 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v160 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v147
	F_errmsg(m, int32(_a_F_check_usermap_3), v15-int32(-64))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2861), int32(_a_F_check_usermap_4))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	v237 = F_strlen(m, v188)
	mBase = m.M
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	v243 = F_palloc0(m, v237+(v193^int32(-1))+v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L15
	} else {
		goto L69
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v179
	v233 = F_list_make1_impl(m, int32(1), v15+int32(40))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L15
	} else {
		goto L67
	}
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v184 == int32(43) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v187 != 0 {
		goto L54
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v190 = F_strstr(m, v188, int32(_a_F_check_usermap_5))
	mBase = m.M
	if v190 == int32(0) {
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	if int32(0) <= v193 {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v196 = int32(1)
	v197 = int32(0)
	v200 = F_errstart(m, int32(15), v197)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	if v200 == int32(0) {
		v329 = v197
		v331 = v196
		goto L25
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v208 + int32(1)
	F_errmsg(m, int32(_a_F_check_usermap_6), v15+int32(48))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2885), int32(_a_F_check_usermap_4))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	v329 = v197
	v331 = v196
	goto L25
L67:
	;
	v235 = F_check_role_2(m, l1, v119, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L15
	} else {
		goto L68
	}
L68:
	;
	v329 = v235
	v331 = int32(0)
	goto L25
L69:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v247 = v190 - v246
	if v247 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	base.MemoryCopy(m, v243, v246, v247)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	v251 = v249 - v250
	if v251 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryCopy(m, v243+v247, l2+v250, v251)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v255 = int32(0)
	v258 = F_strlen(m, v243)
	mBase = m.M
	v260 = F_strcpy(m, v258+v243, v190+int32(2))
	mBase = m.M
	goto L76
L76:
	;
	v261 = F_strlen(m, v243)
	mBase = m.M
	v264 = F_palloc0(m, v261+int32(13))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = int32(0)
	v268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+4)) = uint8(v268)
	v271 = v264 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v271
	v274 = v261 + v268
	if v274 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	base.MemoryCopy(m, v271, v243, v274)
	goto L80
L79:
	;
	goto L80
L80:
	;
	F_pfree(m, v243)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v264
	v283 = F_list_make1_impl(m, int32(1), v15+int32(44))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	v285 = F_check_role_2(m, l1, v119, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	if v287 == int32(0) {
		v329 = v285
		v331 = v255
		goto L25
	} else {
		goto L84
	}
L84:
	;
	F_pg_regfree(m, v287)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	v329 = v285
	v331 = v255
	goto L25
L86:
	;
	if v316-v317 != 0 {
		v329 = v84
		v331 = v84
		goto L25
	} else {
		goto L93
	}
L87:
	;
	goto L86
L88:
	;
	v301 = v292
	v302 = l2
	goto L89
L89:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v306 == int32(0) {
		v316 = v306
		v317 = v305
		goto L87
	} else {
		goto L91
	}
L90:
	;
	v316 = v306
	v317 = v305
	goto L87
L91:
	;
	v309 = int32(1)
	if v306 == v305 {
		v301 = v301 + v309
		v302 = v302 + v309
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v319
	v325 = F_list_make1_impl(m, int32(1), v15+int32(36))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L15
	} else {
		goto L94
	}
L94:
	;
	v327 = F_check_role_2(m, l1, v119, v325)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L15
	} else {
		goto L95
	}
L95:
	;
	v329 = v327
	v331 = v84
	goto L25
L96:
	;
	v337 = v80 + int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v337 < v338 {
		v80 = v337
		goto L23
	} else {
		goto L97
	}
L97:
	;
	goto L24
L98:
	;
	v386 = int32(0) - (v345 ^ int32(1))
	goto L1
L99:
	;
	v357 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L15
	} else {
		goto L100
	}
L100:
	;
	if v357 == int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(_a_F_check_usermap_7), v15+int32(16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L15
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(3008), int32(_a_F_check_usermap_2))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	goto L98
}
func F_chooseNextStatEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v17 = l3
	goto L1
L1:
	;
	v27 = int32(base.Ui32(v17+l4) >> (uint(int32(1)) % 32))
	v28 = base.B2i32(v17 == v27)
	if v17 == v27 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return
L3:
	;
	v40 = v27 + int32(1)
	if v40 == l4 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v31 = int32(base.Ui32(v17+v27) >> (uint(int32(1)) % 32))
	if base.Ui32(v31) < base.Ui32(l5) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v33 = v31 - l5
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v34) <= base.Ui32(v33) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	F_insertStatEntry(m, l0, l1, l2, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	goto L3
L9:
	;
	if v28 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v44 = int32(base.Ui32(v27+(l4+int32(1))) >> (uint(int32(1)) % 32))
	if base.Ui32(v44) < base.Ui32(l5) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v46 = v44 - l5
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v47) <= base.Ui32(v46) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_insertStatEntry(m, l0, l1, l2, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	F_chooseNextStatEntry(m, l0, l1, l2, v17, v27, l5)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v40 != l4 {
		v17 = v40
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	goto L2
}
func F_choose_best_statistics(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
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
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v482 int32
	_ = v482
	v6 = int32(0)
	if l0 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v26 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v46 = int32(2)
	v50 = int32(9)
	v51 = v6
	v53 = v6
	goto L7
L5:
	;
	v482 = v6
	goto L6
L6:
	;
	return v482
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v51<<(uint(int32(2))%32))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+16)))
	if v59 != int32(109) {
		v450 = v46
		v454 = v50
		v457 = v53
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v482 = v457
	goto L6
L9:
	;
	v459 = v51 + int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v459 < v460 {
		v46 = v450
		v50 = v454
		v51 = v459
		v53 = v457
		goto L7
	} else {
		goto L104
	}
L10:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+8)))
	if v62 != l1 {
		v450 = v46
		v454 = v50
		v457 = v53
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v64 = int32(0)
	if base.B2i32(l4 <= int32(0)) == v64 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v78 = v64
	v79 = v64
	v83 = int32(0)
	goto L15
L13:
	;
	v301 = v64
	v302 = v64
	goto L14
L14:
	;
	v313 = int32(0)
	if v301 == v313 {
		goto L60
	} else {
		goto L61
	}
L15:
	;
	v91 = v83 << (uint(int32(2)) % 32)
	v92 = l2 + v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v301 = v277
	v302 = v278
	goto L14
L17:
	;
	v290 = v83 + int32(1)
	if v290 != l4 {
		v78 = v277
		v79 = v278
		v83 = v290
		goto L15
	} else {
		goto L58
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3+v91)))
	if v97 == int32(0) {
		v277 = v78
		v278 = v79
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v101 = int32(0)
	if v93 == v101 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	if v154 == int32(0) {
		v277 = v78
		v278 = v79
		goto L17
	} else {
		goto L36
	}
L23:
	;
	v154 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v100 == int32(0) {
		v147 = v101
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v154 = v147
	goto L22
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v111 < v110 {
		v147 = v101
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v113 = int32(1)
	if v110 <= v113 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v116 = v113
	goto L31
L30:
	;
	v116 = v110
	goto L31
L31:
	;
	v117 = int32(8)
	v122 = int32(0)
	goto L32
L32:
	;
	v129 = v122 << (uint(int32(2)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v93+v117+v129)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v100+v117+v129)))
	v136 = v131 & (v133 ^ int32(-1))
	v138 = base.B2i32(v136 == int32(0))
	if v136 != 0 {
		v147 = v138
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v147 = v138
	goto L26
L34:
	;
	v140 = v122 + int32(1)
	if v140 != v116 {
		v122 = v140
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l3+v91)))
	if v158 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v264 = F_bms_add_members(m, v78, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L48
	} else {
		goto L56
	}
L38:
	;
	v253 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v162 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v164 <= v162 {
		v253 = v162
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v178 = v162
	v183 = v162
	goto L42
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	if v188 == int32(0) {
		v277 = v78
		v278 = v79
		goto L17
	} else {
		goto L44
	}
L43:
	;
	v253 = v236
	goto L37
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v191 <= int32(0) {
		v277 = v78
		v278 = v79
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v183<<(uint(int32(2))%32))))
	v212 = int32(0)
	goto L46
L46:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221+v212<<(uint(int32(2))%32))))
	v226 = F_equal(m, v225, v198)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v236 = F_bms_add_member(m, v178, v212)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L48
	} else {
		goto L54
	}
L48:
	;
	return int32(0)
L49:
	;
	if v226 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v233 = v212 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v233 < v234 {
		v212 = v233
		goto L46
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L47
L53:
	;
	v277 = v78
	v278 = v79
	goto L17
L54:
	;
	v239 = v183 + int32(1)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v239 < v240 {
		v178 = v236
		v183 = v239
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	v266 = F_bms_add_members(m, v79, v253)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L48
	} else {
		goto L57
	}
L57:
	;
	v277 = v264
	v278 = v266
	goto L17
L58:
	;
	goto L16
L59:
	;
	v349 = int32(0)
	if v302 == v349 {
		goto L73
	} else {
		goto L74
	}
L60:
	;
	v348 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v320 = int32(1)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v321 <= v320 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v324 = v320
	goto L65
L64:
	;
	v324 = v321
	goto L65
L65:
	;
	v328 = int32(0)
	v330 = v313
	goto L66
L66:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v301+int32(8)+v328<<(uint(int32(2))%32))))
	if v336 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v348 = v339
	goto L59
L68:
	;
	v339 = v330 + base.I32_popcnt(v336)
	goto L70
L69:
	;
	v339 = v330
	goto L70
L70:
	;
	v341 = v328 + int32(1)
	if v341 != v324 {
		v328 = v341
		v330 = v339
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L67
L72:
	;
	F_bms_free(m, v301)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L48
	} else {
		goto L85
	}
L73:
	;
	v384 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v356 = int32(1)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	if v357 <= v356 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v360 = v356
	goto L78
L77:
	;
	v360 = v357
	goto L78
L78:
	;
	v364 = int32(0)
	v366 = v349
	goto L79
L79:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v302+int32(8)+v364<<(uint(int32(2))%32))))
	if v372 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v384 = v375
	goto L72
L81:
	;
	v375 = v366 + base.I32_popcnt(v372)
	goto L83
L82:
	;
	v375 = v366
	goto L83
L83:
	;
	v377 = v364 + int32(1)
	if v377 != v360 {
		v364 = v377
		v366 = v375
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	F_bms_free(m, v302)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L48
	} else {
		goto L86
	}
L86:
	;
	v389 = v384 + v348
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v391 = int32(0)
	if v390 == v391 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	if v428 != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v426 = int32(0)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v398 = int32(1)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v399 <= v398 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v402 = v398
	goto L93
L92:
	;
	v402 = v399
	goto L93
L93:
	;
	v406 = int32(0)
	v408 = v391
	goto L94
L94:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v390+int32(8)+v406<<(uint(int32(2))%32))))
	if v414 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v426 = v417
	goto L87
L96:
	;
	v417 = v408 + base.I32_popcnt(v414)
	goto L98
L97:
	;
	v417 = v408
	goto L98
L98:
	;
	v419 = v406 + int32(1)
	if v419 != v402 {
		v406 = v419
		v408 = v417
		goto L94
	} else {
		goto L99
	}
L99:
	;
	goto L95
L100:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v431 = v429
	goto L102
L101:
	;
	v431 = int32(0)
	goto L102
L102:
	;
	v432 = v431 + v426
	if (base.B2i32(v389 != v46)|base.B2i32(v50 <= v432))&base.B2i32(v389 <= v46) != 0 {
		v450 = v46
		v454 = v50
		v457 = v53
		goto L9
	} else {
		goto L103
	}
L103:
	;
	v450 = v389
	v454 = v432
	v457 = v58
	goto L9
L104:
	;
	goto L8
}
func F_choose_next_subplan_for_worker(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v8 = F_LWLockAcquire(m, v6, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v12 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v94 == int32(-1) {
		v387 = int32(0)
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v12)+20)) = uint8(v17)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v19 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v24 = F_ExecFindMatchingSubPlans(m, v21, v20, v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v24
	v29 = int32(0)
	if v24 == v29 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if base.B2i32(v64 == v65)|base.B2i32(v65 <= int32(0)) != 0 {
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v64 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v37 <= v36 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = v36
	goto L15
L14:
	;
	v40 = v37
	goto L15
L15:
	;
	v44 = int32(0)
	v46 = v29
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v44<<(uint(int32(2))%32))))
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v64 = v55
	goto L9
L18:
	;
	v55 = v46 + base.I32_popcnt(v52)
	goto L20
L19:
	;
	v55 = v46
	goto L20
L20:
	;
	v57 = v44 + int32(1)
	if v57 != v40 {
		v44 = v57
		v46 = v55
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v71 = v20
	goto L23
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v76 = F_bms_is_member(m, v71, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L3
L25:
	;
	if v76 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v71)+20)) = uint8(v82)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v85 = v71 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v85 < v86 {
		v71 = v85
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	F_LWLockRelease(m, v6)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L93
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v94
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v102 = v100
	goto L33
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v102
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v246 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L33:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+(v6+int32(20))))))
	if v107 != int32(1) {
		goto L32
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(-1)
	F_LWLockRelease(m, v6)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L63
	}
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v110 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v234
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v234 != v237 {
		v102 = v234
		goto L33
	} else {
		goto L62
	}
L37:
	;
	if int32(0) <= v166 {
		v234 = v166
		goto L36
	} else {
		goto L48
	}
L38:
	;
	v166 = base.I32_ctz(v152) | v153<<(uint(int32(5))%32)
	goto L37
L39:
	;
	v166 = int32(-2)
	goto L37
L40:
	;
	v117 = v102 + int32(1)
	v119 = base.I32_div_s(v117, int32(32))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v120 <= v119 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v123 = v110 + int32(8)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v119<<(uint(int32(2))%32))))
	v130 = v127 & (int32(-1) << (uint(v117) % 32))
	if v130 != 0 {
		v152 = v130
		v153 = v119
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v132 = v119 + int32(1)
	if v132 == v120 {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v135 = v132
	goto L44
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v123+v135<<(uint(int32(2))%32))))
	if v142 != 0 {
		v152 = v142
		v153 = v135
		goto L38
	} else {
		goto L46
	}
L45:
	;
	goto L39
L46:
	;
	v144 = v135 + int32(1)
	if v144 != v120 {
		v135 = v144
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v169 <= v170 {
		v234 = v169
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v172 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if int32(0) <= v230 {
		v234 = v230
		goto L36
	} else {
		goto L61
	}
L51:
	;
	v230 = base.I32_ctz(v216) | v217<<(uint(int32(5))%32)
	goto L50
L52:
	;
	v230 = int32(-2)
	goto L50
L53:
	;
	v181 = v170 - int32(1) + int32(1)
	v183 = base.I32_div_s(v181, int32(32))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v184 <= v183 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v187 = v172 + int32(8)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187+v183<<(uint(int32(2))%32))))
	v194 = v191 & (int32(-1) << (uint(v181) % 32))
	if v194 != 0 {
		v216 = v194
		v217 = v183
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v196 = v183 + int32(1)
	if v196 == v184 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v199 = v196
	goto L57
L57:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v187+v199<<(uint(int32(2))%32))))
	if v206 != 0 {
		v216 = v206
		v217 = v199
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v208 = v199 + int32(1)
	if v208 != v184 {
		v199 = v208
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v234 = v233
	goto L36
L62:
	;
	goto L34
L63:
	;
	return int32(0)
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v303
	if v303 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L65:
	;
	v303 = base.I32_ctz(v289) | v290<<(uint(int32(5))%32)
	goto L64
L66:
	;
	v303 = int32(-2)
	goto L64
L67:
	;
	v254 = v247 + int32(1)
	v256 = base.I32_div_s(v254, int32(32))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v257 <= v256 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v260 = v246 + int32(8)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260+v256<<(uint(int32(2))%32))))
	v267 = v264 & (int32(-1) << (uint(v254) % 32))
	if v267 != 0 {
		v289 = v267
		v290 = v256
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v269 = v256 + int32(1)
	if v269 == v257 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v272 = v269
	goto L71
L71:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v260+v272<<(uint(int32(2))%32))))
	if v279 != 0 {
		v289 = v279
		v290 = v272
		goto L65
	} else {
		goto L73
	}
L72:
	;
	goto L66
L73:
	;
	v281 = v272 + int32(1)
	if v281 != v257 {
		v272 = v281
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v308 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	goto L77
L77:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v375 <= v374 {
		v387 = int32(1)
		goto L30
	} else {
		goto L92
	}
L78:
	;
	if v367 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L79:
	;
	v367 = base.I32_ctz(v353) | v354<<(uint(int32(5))%32)
	goto L78
L80:
	;
	v367 = int32(-2)
	goto L78
L81:
	;
	v318 = v309 - int32(1) + int32(1)
	v320 = base.I32_div_s(v318, int32(32))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v321 <= v320 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v324 = v308 + int32(8)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324+v320<<(uint(int32(2))%32))))
	v331 = v328 & (int32(-1) << (uint(v318) % 32))
	if v331 != 0 {
		v353 = v331
		v354 = v320
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v333 = v320 + int32(1)
	if v333 == v321 {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v336 = v333
	goto L85
L85:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v324+v336<<(uint(int32(2))%32))))
	if v343 != 0 {
		v353 = v343
		v354 = v336
		goto L79
	} else {
		goto L87
	}
L86:
	;
	goto L80
L87:
	;
	v345 = v336 + int32(1)
	if v345 != v321 {
		v336 = v345
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v370 = int32(-1)
	goto L91
L90:
	;
	v370 = v367
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v370
	goto L77
L92:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v379 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v377+v374)+20)) = uint8(v379)
	v387 = v379
	goto L30
L93:
	;
	return v387
}
func F_choose_next_subplan_locally(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v7 != 0 {
		v183 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v183
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v8 != int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v28 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v11 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v14 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v16 = int32(0)
	v18 = F_ExecFindMatchingSubPlans(m, v15, v16, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v18
	goto L3
L9:
	;
	if v166 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L10:
	;
	if v26 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	goto L12
L12:
	;
	v87 = int32(0)
	if base.B2i32(v26 == v87)|base.B2i32(v8 == v87) != 0 {
		goto L26
	} else {
		goto L27
	}
L13:
	;
	v166 = v86
	goto L9
L14:
	;
	v86 = base.I32_ctz(v72) | v73<<(uint(int32(5))%32)
	goto L13
L15:
	;
	v86 = int32(-2)
	goto L13
L16:
	;
	v37 = v8 + int32(1)
	v39 = base.I32_div_s(v37, int32(32))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v40 <= v39 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v43 = v26 + int32(8)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v39<<(uint(int32(2))%32))))
	v50 = v47 & (int32(-1) << (uint(v37) % 32))
	if v50 != 0 {
		v72 = v50
		v73 = v39
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v52 = v39 + int32(1)
	if v52 == v40 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v55 = v52
	goto L20
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43+v55<<(uint(int32(2))%32))))
	if v62 != 0 {
		v72 = v62
		v73 = v55
		goto L14
	} else {
		goto L22
	}
L21:
	;
	goto L15
L22:
	;
	v64 = v55 + int32(1)
	if v64 != v40 {
		v55 = v64
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v166 = v159
	goto L9
L25:
	;
	v159 = base.I32_clz(v144) | v142<<(uint(int32(5))%32) ^ int32(31)
	goto L24
L26:
	;
	v159 = int32(-2)
	goto L24
L27:
	;
	if v8 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v97 = v94 << (uint(int32(5)) % 32)
	goto L30
L29:
	;
	v97 = v8
	goto L30
L30:
	;
	v99 = v97 - int32(1)
	if v99 < int32(-31) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v103 = v26 + int32(8)
	v105 = base.I32_div_s(v99, int32(32))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+v105<<(uint(int32(2))%32))))
	v117 = v109 & int32(base.Ui32(int32(-1))>>(uint(v105<<(uint(int32(5))%32)-v99+int32(31))%32))
	if v117 != 0 {
		v142 = v105
		v144 = v117
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if v99 < int32(32) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v121 = v105
	goto L34
L34:
	;
	v127 = v121 - int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v103+v127<<(uint(int32(2))%32))))
	if v131 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L26
L36:
	;
	v142 = v127
	v144 = v131
	goto L25
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(1)) < base.Ui32(v121) {
		v121 = v127
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v169 <= int32(0) {
		v183 = v2
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v166
	v183 = int32(1)
	goto L1
L43:
	;
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(v172)
	return int32(0)
}
func F_cleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_markreachable(m, l0, v10, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_markcanreach(m, l0, v13, v14, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_cleartraverse(m, l0, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L86
	}
L8:
	;
	v22 = v17
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v26 != 0 {
		goto L7
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 == v29 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v27 != 0 {
		v22 = v27
		goto L9
	} else {
		goto L85
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
	if v31 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v37 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L46
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+4)))
	if v44 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L15
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v79 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v49 = v47 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v49))|base.B2i32(int32(1)<<(uint(v49)%32)&int32(_a_F_cleanup_0) == int32(0)) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v59 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	if v60 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v72 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v44*int32(24))+12)) = v68
	v72 = v68
	goto L25
L27:
	;
	goto L28
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v70
	v72 = v70
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+36)) = v60
	goto L31
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+32)) = int64(0)
	goto L21
L32:
	;
	if v78 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v78
	goto L32
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v78
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v79
	goto L38
L37:
	;
	goto L38
L38:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v85 - int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v90 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v89 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v89
	goto L39
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v89
	goto L39
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v90
	goto L45
L44:
	;
	goto L45
L45:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v96 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
	v103 = v37 + int32(8)
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v103)+16)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v103)+8)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v103))) = v104
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v37
	goto L20
L46:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v118 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(-1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v199 != 0 {
		goto L78
	} else {
		goto L79
	}
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118)+4)))
	if v125 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L46
L52:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v160 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v130 = v128 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v130))|base.B2i32(int32(1)<<(uint(v130)%32)&int32(_a_F_cleanup_0) == int32(0)) != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v140 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	if v141 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v153 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v125*int32(24))+12)) = v149
	v153 = v149
	goto L56
L58:
	;
	goto L59
L59:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+32)) = v151
	v153 = v151
	goto L56
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+36)) = v141
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118)+32)) = int64(0)
	goto L52
L63:
	;
	if v159 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+20)) = v159
	goto L63
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v159
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = v160
	goto L69
L68:
	;
	goto L69
L69:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v166 - int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	if v171 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v170 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v170
	goto L70
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v170
	goto L70
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+28)) = v171
	goto L76
L75:
	;
	goto L76
L76:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v177 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	v184 = v118 + int32(8)
	v185 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v184)+16)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v184)+8)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = v185
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v118
	goto L51
L77:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v198 != 0 {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+32)) = v198
	goto L77
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v198
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v22
	goto L12
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+28)) = v202
	goto L81
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v202
	goto L81
L85:
	;
	goto L10
L86:
	;
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v226 = v223
	v227 = v224
	goto L90
L88:
	;
	v235 = v223
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v235
	goto L3
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v226
	v232 = v226 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	if v233 != 0 {
		v226 = v232
		v227 = v233
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v235 = v232
	goto L89
L92:
	;
	goto L91
}
func F_cloneouts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_cloneouts[0]))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v21 <= v22 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v78 != 0 {
		v14 = v78
		goto L4
	} else {
		goto L33
	}
L12:
	;
	F_createarc(m, l0, l4, v16, l2, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L32
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v24 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v43 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L16:
	;
	v28 = v24
	goto L17
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v34 != l3 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L12
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v42 != 0 {
		v28 = v42
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if v36 != v16&int32(_a_F_cloneouts_0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v40 == l4 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	goto L18
L24:
	;
	v47 = v43
	goto L25
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	if v53 != l2 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L12
L27:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v61 != 0 {
		v47 = v61
		goto L25
	} else {
		goto L31
	}
L28:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)))
	if v55 != v16&int32(_a_F_cloneouts_0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v59 == l4 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	goto L26
L32:
	;
	goto L11
L33:
	;
	goto L5
}
func F_close_ls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v41 float64
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 float64
	_ = v66
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = v12 + int32(16)
	v15 = F_point_sl(m, v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v20 = base.F64_abs(v19)
		if base.F64_le(v20, float64(1e-06)) != 0 {
			v41 = float64(0)
			if base.F64_eq(v41, v15) != 0 {
				v78 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
				v84 = int32(0)
				return v84
			} else {
				v47 = F_palloc(m, int32(16))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = F_lseg_interpt_line(m, v47, v12, v11)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 != 0 {
							v66 = float64(0)
							if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v84 = v47
							} else {
								v78 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
								v84 = int32(0)
							}
							return v84
						} else {
							v52 = F_line_closept_point(m, int32(0), v11, v12)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v55 = F_line_closept_point(m, int32(0), v11, v14)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v57 = base.F64_lt(v52, v55)
									if v57 != 0 {
										v58 = v52
									} else {
										v58 = v55
									}
									if v47 == int32(0) {
										v66 = v58
									} else {
										if v57 != 0 {
											v61 = v12
										} else {
											v61 = v14
										}
										v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v62
										v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
										*(*int64)(unsafe.Add(mBase, uint32(v47))) = v64
										v66 = v58
									}
									if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v84 = v47
									} else {
										v78 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
										v84 = int32(0)
									}
									return v84
								}
							}
						}
					}
				}
			}
		} else {
			v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v25 = base.F64_abs(v24)
			if base.F64_le(v25, float64(1e-06)) != 0 {
				v41 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(v41, v15) != 0 {
					v78 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
					v84 = int32(0)
					return v84
				} else {
					v47 = F_palloc(m, int32(16))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_lseg_interpt_line(m, v47, v12, v11)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v66 = float64(0)
								if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v84 = v47
								} else {
									v78 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
									v84 = int32(0)
								}
								return v84
							} else {
								v52 = F_line_closept_point(m, int32(0), v11, v12)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v55 = F_line_closept_point(m, int32(0), v11, v14)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v57 = base.F64_lt(v52, v55)
										if v57 != 0 {
											v58 = v52
										} else {
											v58 = v55
										}
										if v47 == int32(0) {
											v66 = v58
										} else {
											if v57 != 0 {
												v61 = v12
											} else {
												v61 = v14
											}
											v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v62
											v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
											*(*int64)(unsafe.Add(mBase, uint32(v47))) = v64
											v66 = v58
										}
										if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v84 = v47
										} else {
											v78 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
											v84 = int32(0)
										}
										return v84
									}
								}
							}
						}
					}
				}
			} else {
				v30 = base.F64_div(v19, base.F64_neg(v24))
				v32 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v30), v32)&base.F64_ne(v20, v32) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if base.F64_ne(v30, float64(0)) != 0 {
						v41 = v30
						if base.F64_eq(v41, v15) != 0 {
							v78 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
							v84 = int32(0)
							return v84
						} else {
							v47 = F_palloc(m, int32(16))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v49 = F_lseg_interpt_line(m, v47, v12, v11)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									if v49 != 0 {
										v66 = float64(0)
										if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v84 = v47
										} else {
											v78 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
											v84 = int32(0)
										}
										return v84
									} else {
										v52 = F_line_closept_point(m, int32(0), v11, v12)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v55 = F_line_closept_point(m, int32(0), v11, v14)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												v57 = base.F64_lt(v52, v55)
												if v57 != 0 {
													v58 = v52
												} else {
													v58 = v55
												}
												if v47 == int32(0) {
													v66 = v58
												} else {
													if v57 != 0 {
														v61 = v12
													} else {
														v61 = v14
													}
													v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v62
													v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
													*(*int64)(unsafe.Add(mBase, uint32(v47))) = v64
													v66 = v58
												}
												if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
													v84 = v47
												} else {
													v78 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
													v84 = int32(0)
												}
												return v84
											}
										}
									}
								}
							}
						}
					} else {
						if base.F64_ne(v25, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v41 = v30
							if base.F64_eq(v41, v15) != 0 {
								v78 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
								v84 = int32(0)
								return v84
							} else {
								v47 = F_palloc(m, int32(16))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = F_lseg_interpt_line(m, v47, v12, v11)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										if v49 != 0 {
											v66 = float64(0)
											if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
												v84 = v47
											} else {
												v78 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
												v84 = int32(0)
											}
											return v84
										} else {
											v52 = F_line_closept_point(m, int32(0), v11, v12)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												v55 = F_line_closept_point(m, int32(0), v11, v14)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return int32(0)
												} else {
													v57 = base.F64_lt(v52, v55)
													if v57 != 0 {
														v58 = v52
													} else {
														v58 = v55
													}
													if v47 == int32(0) {
														v66 = v58
													} else {
														if v57 != 0 {
															v61 = v12
														} else {
															v61 = v14
														}
														v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v62
														v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
														*(*int64)(unsafe.Add(mBase, uint32(v47))) = v64
														v66 = v58
													}
													if base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
														v84 = v47
													} else {
														v78 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
														v84 = int32(0)
													}
													return v84
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
}
func F_close_sb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_box_closept_lseg(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v22 = int32(0)
			} else {
				v22 = v8
			}
			return v22
		}
	}
}
func F_closelog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = m.G0
	v3 = int32(16)
	v4 = v2 - v3
	m.G0 = v4
	v6 = int32(_a_F_closelog_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_closelog[0]))
	v8 = F_close(m, v7)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_closelog[0])) = int32(-1)
	m.G0 = v4 + v3
	return
}
func F_cluster_rel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 float64
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 float64
	_ = v511
	var v512 float64
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 float64
	_ = v520
	var v521 float64
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 float64
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 float64
	_ = v537
	var v540 int32
	_ = v540
	var v541 float64
	_ = v541
	var v547 float64
	_ = v547
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 float64
	_ = v564
	var v565 float64
	_ = v565
	var v576 int32
	_ = v576
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 float64
	_ = v751
	var v752 float64
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	var v768 float64
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 float64
	_ = v803
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(400)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[0]))
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[1]))
	if v39 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L6
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[2])))
	if v43&int32(1) == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = int32(_a_F_cluster_rel_0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v50 + v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v54 + v51
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = v28
	base.MemoryFill(m, v39+int32(232), int32(0), int32(160))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v65 + v51
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v71 - v51
	goto L7
L10:
	;
	v78 = int64(1)
	goto L12
L11:
	;
	v78 = int64(2)
	goto L12
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[1]))
	if v81 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(124)))) = v119
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(120)))) = v122
	goto L17
L14:
	;
	goto L13
L15:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[2])))
	if v85&int32(1) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v90 = int32(_a_F_cluster_rel_0)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	v93 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v92 + v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v96 + v93
	*(*int64)(unsafe.Add(mBase, uint32(v81+int32(0))+232)) = v78
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v104 + v93
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v110 - v93
	goto L14
L17:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+80))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[5])) = v126 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[4])) = v125
	goto L18
L18:
	;
	v134 = int32(_a_F_cluster_rel_1)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[6]))
	v138 = v136 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[6])) = v138
	goto L19
L19:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v29&int32(2) != 0 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L203
	}
L22:
	;
	F_errmsg(m, int32(_a_F_cluster_rel_2), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L4
	} else {
		goto L201
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L4
	} else {
		goto L197
	}
L24:
	;
	F_AtEOXact_GUC(m, int32(0), v138)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L4
	} else {
		goto L190
	}
L25:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+118)))
	if v204 != int32(116) {
		goto L58
	} else {
		goto L59
	}
L26:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+117)))
	if v200 == int32(1) {
		goto L23
	} else {
		goto L57
	}
L27:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v26)+124))
	v144 = F_pg_class_aclcheck(m, v28, v142, int64(16384))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l1 != 0 {
		goto L26
	} else {
		goto L56
	}
L30:
	;
	if v144 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v148 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+118)))
	if v167 != int32(116) {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	if v148 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = F_get_rel_name(m, v28)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L41
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v150
	F_errmsg(m, int32(_a_F_cluster_rel_3), v26+int32(112))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(1752), int32(_a_F_cluster_rel_5))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	goto L24
L42:
	;
	if l1 == int32(0) {
		v203 = v166
		goto L25
	} else {
		goto L46
	}
L43:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v170 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	goto L24
L46:
	;
	v177 = int32(0)
	v180 = F_SearchSysCacheExists(m, int32(57), l1, v177, v177, v177)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v180 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v187&int32(4) == int32(0) {
		goto L26
	} else {
		goto L52
	}
L51:
	;
	goto L24
L52:
	;
	v192 = F_get_index_isclustered(m, l1)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v192 != 0 {
		goto L26
	} else {
		goto L54
	}
L54:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L24
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v203 = v197
	goto L25
L57:
	;
	v203 = v199
	goto L25
L58:
	;
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v207 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_cluster_rel_6), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(425), int32(_a_F_cluster_rel_7))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v226 = int32(_a_F_cluster_rel_8)
	goto L68
L67:
	;
	v226 = int32(_a_F_cluster_rel_9)
	goto L68
L68:
	;
	F_CheckTableNotInUse(m, l0, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	if l1 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_check_index_is_clusterable(m, l0, l1, int32(8))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	v236 = int32(0)
	goto L72
L72:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+119)))
	if v238 != int32(109) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v233 = F_index_open(m, l1, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v236 = v233
	goto L72
L75:
	;
	F_TransferPredicateLocksToHeapRelation(m, l0)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L79
	}
L76:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+129)))
	if v241 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L24
L79:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+92))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v248)+84))
	if v236 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v236)+56))
	F_mark_index_clustered(m, l0, v251, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	v256 = v248
	goto L82
L82:
	;
	v257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256)+118)))
	v259 = int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v260) < base.Ui32(int32(_a_F_cluster_rel_10)) {
		v269 = v259
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v256 = v255
	goto L82
L84:
	;
	v271 = F_make_new_heap(m, v247, v249, v250, v257, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+68))
	if v264 == int32(99) {
		v269 = v259
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v267 = F_isTempToastNamespace(m, v264)
	mBase = m.M
	v269 = v267
	goto L85
L88:
	;
	v274 = F_table_open(m, v271, int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v276 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+312)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v26)+304)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v26)+296)) = v276
	F_getrusage(m, v26+int32(144))
	mBase = m.M
	F_gettimeofday(m, v26+int32(128))
	mBase = m.M
	goto L90
L90:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+68))
	v290 = F_get_namespace_name(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+112))
	if v293 == int32(0) {
		v310 = v4
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v311 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+392)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v26)+384)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v26)+376)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v26)+368)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v26)+360)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v26)+352)) = v311
	*(*int64)(unsafe.Add(mBase, uint32(v26)+344)) = v311
	v329 = F_vacuum_get_cutoffs(m, l0, v26+int32(344), v26+int32(320))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L97
	}
L93:
	;
	F_LockRelationOid(m, v293, int32(8))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+112))
	if v300 == int32(0) {
		v310 = v4
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v274)+48))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+112))
	if v304 == int32(0) {
		v310 = v4
		goto L92
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+264)) = v300
	v310 = int32(1)
	goto L92
L97:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+136))
	if v332 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+140))
	if v354 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L99:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v26)+336))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v332))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v335)) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v347 == int32(0) {
		goto L98
	} else {
		goto L104
	}
L101:
	;
	v347 = base.B2i32(base.Ui32(v335) < base.Ui32(v332))
	goto L100
L102:
	;
	goto L103
L103:
	;
	v347 = int32(base.Ui32(v335-v332) >> (uint(int32(31)) % 32))
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+336)) = v332
	goto L98
L105:
	;
	if v29&int32(1) != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v26)+340))
	goto L107
L107:
	;
	if int32(base.Ui32(v357-v354)>>(uint(int32(31))%32)) == int32(0) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+340)) = v354
	goto L105
L109:
	;
	v366 = int32(17)
	goto L111
L110:
	;
	v366 = int32(13)
	goto L111
L111:
	;
	if v236 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v26)+328))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)+124))
	m.T0[v737].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v274, v236, v706, v725, v26+int32(336), v26+int32(340), v26+int32(312), v26+int32(304), v26+int32(296))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L160
	}
L113:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), v675, int32(_a_F_cluster_rel_11))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L159
	}
L114:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+84))
	if v368 == int32(403) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	goto L116
L116:
	;
	v656 = int32(0)
	v658 = F_errstart(m, v366, v656)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L156
	}
L117:
	;
	v640 = F_errstart(m, v366, int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L153
	}
L118:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v236)+56))
	v373 = m.G0
	v375 = v373 - int32(112)
	m.G0 = v375
	v377 = int32(1)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[7])))
	if v379 != v377 {
		v576 = v377
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v616 = int32(0)
	v618 = F_errstart(m, v366, v616)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L150
	}
L121:
	;
	m.G0 = v375 + int32(112)
	if v576 != 0 {
		goto L117
	} else {
		goto L149
	}
L122:
	;
	v383 = F_palloc0(m, int32(168))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v383))) = int64(4294967363)
	v388 = F_palloc0(m, int32(92))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = int32(266)
	v393 = F_palloc0(m, int32(384))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v393)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v393)+8)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v393)+4)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = int32(267)
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v393)+344)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v393)+280)) = v402
	v407 = F_palloc0(m, int32(8))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v375)+12)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v375)+20)) = v407
	v416 = F_list_make1_impl(m, int32(1), v375+int32(12))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v393)+84)) = v416
	v420 = F_palloc0(m, int32(136))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v422 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v420)+24)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v420)+16)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v420)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = int32(101)
	v429 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v420)+124)) = uint16(v429)
	v431 = int32(_a_F_cluster_rel_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v420)+20)) = uint16(v431)
	*(*int32)(unsafe.Add(mBase, uint32(v375)+8)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v375)+16)) = v420
	v438 = F_list_make1_impl(m, v422, v375+int32(8))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+52)) = v438
	v443 = F_addRTEPermissionInfo(m, v383+int32(56), v420)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_setup_simple_rel_arrays(m, v393)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v449 = F_build_simple_rel(m, v393, int32(1), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v449)+108))
	if v451 == int32(0) {
		v576 = v377
		goto L121
	} else {
		goto L133
	}
L133:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v454 <= int32(0) {
		v576 = v377
		goto L121
	} else {
		goto L134
	}
L134:
	;
	v457 = int32(0)
	if v457 < v454 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v461 = v454
	goto L137
L136:
	;
	v461 = v457
	goto L137
L137:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	v467 = v457
	goto L138
L138:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v462+v467<<(uint(int32(2))%32))))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v372 != v490 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v496 = *(*float64)(unsafe.Add(mBase, uint32(v449)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v449)+16)) = v496
	v499 = F_get_relation_data_width(m, v371, int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L144
	}
L140:
	;
	v492 = int32(1)
	v494 = v467 + v492
	if v461 != v494 {
		v467 = v494
		goto L138
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	v576 = v492
	goto L121
L144:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v449)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+32)) = v499
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v449)+116))
	*(*float64)(unsafe.Add(mBase, uint32(v393)+288)) = base.F64_convert_i32_u(v503)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v489)+84))
	F_cost_qual_eval(m, v375+int32(96), v508, v393)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v511 = *(*float64)(unsafe.Add(mBase, uint32(v375)+104))
	v512 = *(*float64)(unsafe.Add(mBase, uint32(v375)+96))
	v514 = v375 + int32(24)
	v515 = int32(0)
	v517 = F_create_seqscan_path(m, v393, v449, v515, v515)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v517)+40))
	v520 = *(*float64)(unsafe.Add(mBase, uint32(v517)+56))
	v521 = *(*float64)(unsafe.Add(mBase, uint32(v449)+120))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v449)+28))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+32))
	v524 = base.F64_add(v512, v511)
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[9]))
	v530 = m.G0
	v531 = int32(16)
	v532 = v530 - v531
	m.G0 = v532
	F_cost_tuplesort(m, v532+int32(8), v532, v521, v523, base.F64_add(v524, v524), v527, float64(-1))
	mBase = m.M
	v537 = *(*float64)(unsafe.Add(mBase, uint32(v532)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v514)+32)) = v521
	v540 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[10])))
	v541 = base.F64_add(v520, v537)
	*(*float64)(unsafe.Add(mBase, uint32(v514)+48)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v514)+40)) = v519 + (v540 ^ int32(1))
	v547 = *(*float64)(unsafe.Add(mBase, uint32(v532)))
	*(*float64)(unsafe.Add(mBase, uint32(v514)+56)) = base.F64_add(v541, v547)
	m.G0 = v532 + v531
	goto L147
L147:
	;
	v553 = int32(0)
	v562 = F_create_index_path(m, v393, v489, v553, v553, v553, v553, int32(1), v553, v553, float64(1), v553)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v564 = *(*float64)(unsafe.Add(mBase, uint32(v375)+80))
	v565 = *(*float64)(unsafe.Add(mBase, uint32(v562)+56))
	v576 = base.F64_lt(v564, v565)
	goto L121
L149:
	;
	goto L120
L150:
	;
	if v618 == int32(0) {
		v706 = v616
		goto L112
	} else {
		goto L151
	}
L151:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v290
	v625 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v623 + v625
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v622 + v625
	F_errmsg(m, int32(_a_F_cluster_rel_13), v26+int32(96))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v675 = int32(964)
	v697 = int32(0)
	goto L113
L153:
	;
	if v640 == int32(0) {
		v706 = int32(1)
		goto L112
	} else {
		goto L154
	}
L154:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v644 + int32(4)
	F_errmsg(m, int32(_a_F_cluster_rel_14), v26+int32(80))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v675 = int32(969)
	v697 = int32(1)
	goto L113
L156:
	;
	if v658 == int32(0) {
		v706 = v656
		goto L112
	} else {
		goto L157
	}
L157:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v662 + int32(4)
	F_errmsg(m, int32(_a_F_cluster_rel_15), v26-int32(-64))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	v675 = int32(974)
	v697 = int32(0)
	goto L113
L159:
	;
	v706 = v697
	goto L112
L160:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v26)+340))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v26)+336))
	v742 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+264)) = v742
	v745 = F_RelationGetNumberOfBlocksInFork(m, v274, v742)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v748 = F_errstart(m, v366, int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	if v748 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v751 = *(*float64)(unsafe.Add(mBase, uint32(v26)+304))
	v752 = *(*float64)(unsafe.Add(mBase, uint32(v26)+312))
	v754 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v790 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L4
	} else {
		goto L171
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v754
	*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v752
	*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v750 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v290
	F_errmsg(m, int32(_a_F_cluster_rel_16), v26+int32(32))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v768 = *(*float64)(unsafe.Add(mBase, uint32(v26)+296))
	v771 = F_pg_rusage_show(m, v26+int32(128))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v771
	*(*float64)(unsafe.Add(mBase, uint32(v26)+16)) = v768
	F_errdetail(m, int32(_a_F_cluster_rel_17), v26+int32(16))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(1007), int32(_a_F_cluster_rel_11))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	goto L165
L171:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v274)+56))
	v795 = F_SearchSysCacheCopy(m, int32(57), v793, int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	if v795 == int32(0) {
		goto L21
	} else {
		goto L173
	}
L173:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v795)+16))
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799)+22)))
	v801 = v799 + v800
	*(*int32)(unsafe.Add(mBase, uint32(v801)+96)) = v745
	v803 = *(*float64)(unsafe.Add(mBase, uint32(v26)+312))
	*(*float32)(unsafe.Add(mBase, uint32(v801)+100)) = base.F32_demote_f64(v803)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v806 != int32(1259) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	F_pfree(m, v795)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L180
	}
L175:
	;
	F_CatalogTupleUpdate(m, v790, v795+int32(4), v795)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L4
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	F_CacheInvalidateRelcacheByTuple(m, v795)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L4
	} else {
		goto L179
	}
L178:
	;
	goto L174
L179:
	;
	goto L174
L180:
	;
	F_relation_close(m, v790, int32(3))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	F_relation_close(m, l0, int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	if v236 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	F_relation_close(m, v236, int32(0))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	F_relation_close(m, v274, int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	F_finish_heap_swap(m, v247, v271, v269, v310, int32(0), int32(1), v741, v740, v257)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	goto L24
L190:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v26)+124))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[5])) = v862
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[4])) = v861
	goto L191
L191:
	;
	v869 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[1]))
	if v869 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	m.G0 = v26 + int32(400)
	return
L193:
	;
	goto L192
L194:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[2])))
	if v873&int32(1) == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v869)+220))
	if v878 == int32(0) {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v881 = int32(_a_F_cluster_rel_0)
	v883 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	v884 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v883 + v884
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v869)))
	*(*int32)(unsafe.Add(mBase, uint32(v869))) = v887 + v884
	v891 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v869)+220)) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v869)+224)) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v869))) = v887 + int32(2)
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v901 - v884
	goto L193
L197:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	F_errmsg(m, int32(_a_F_cluster_rel_18), int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(410), int32(_a_F_cluster_rel_7))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(421), int32(_a_F_cluster_rel_7))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v274)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v937
	F_errmsg_internal(m, int32(_a_F_cluster_rel_19), v26)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(1016), int32(_a_F_cluster_rel_11))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	v5 = l2 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v67
L2:
	;
	v67 = int32(0)
	goto L1
L3:
	;
	v41 = v36
	v42 = v37
	v43 = v38
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v36 = l0
		v37 = l1
		v38 = v5
		goto L3
	} else {
		goto L7
	}
L5:
	;
	v29 = l0
	v30 = l1
	v31 = v5
	goto L6
L6:
	;
	if v31 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v13 = l0
	v14 = l1
	v15 = v5
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18 != v19 {
		v36 = v13
		v37 = v14
		v38 = v15
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v29 = v24
	v30 = v22
	v31 = v26
	goto L6
L10:
	;
	v21 = int32(4)
	v22 = v14 + v21
	v24 = v13 + v21
	v26 = v15 - v21
	if base.Ui32(int32(3)) < base.Ui32(v26) {
		v13 = v24
		v14 = v22
		v15 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = v29
	v37 = v30
	v38 = v31
	goto L3
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 == v47 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v67 = v46 - v47
	goto L1
L15:
	;
	v49 = int32(1)
	v54 = v43 - v49
	if v54 != 0 {
		v41 = v41 + v49
		v42 = v42 + v49
		v43 = v54
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_cmpEntries(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = int32(1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v8 == v6 {
		if v7&int32(1) != 0 {
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v25)
			v28 = int32(0)
		} else {
			v28 = v6
		}
		return v28
	} else {
		if v7&int32(1) != 0 {
			v28 = int32(-1)
			return v28
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v20 = F_FunctionCall2Coll(m, v16, v17, v18, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 != 0 {
					v28 = v20
				} else {
					v25 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v25)
					v28 = int32(0)
				}
				return v28
			}
		}
	}
}
func F_cmp_lbestatus(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+408))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+408))
	return v3 - v4
}
func F_cntsize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v7 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34 + v35&int32(4095) + int32(1)
	goto L3
L7:
	;
	v22 = int32(0)
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v22<<(uint(int32(2))%32))))
	F_cntsize(m, v27, l1, l2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	v31 = v22 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 < v32 {
		v22 = v31
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
}
func F_codepoint_range_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v4) <= base.Ui32(v3) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v9 = base.B2i32(base.Ui32(v6) < base.Ui32(v3))
	} else {
		v9 = int32(-1)
	}
	return v9
}
func F_colname_is_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
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
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v8 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v253
L2:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v190 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v11 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v177 = int32(0)
	v179 = F_hash_search(m, v8, l0, v177, v177)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L56
	} else {
		goto L57
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v19 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v67 {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14+v19<<(uint(int32(2))%32))))
	if v26 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v58 = v19 + int32(1)
	if v58 != v11 {
		v19 = v58
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v31 == int32(0))|base.B2i32(v31 != v34) != 0 {
		v52 = v31
		v53 = v34
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v52-v53 != 0 {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v37 = v26
	v38 = l0
	goto L16
L16:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v42
		v53 = v41
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v52 = v42
	v53 = v41
	goto L14
L18:
	;
	v45 = int32(1)
	if v42 == v41 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(0)
L21:
	;
	goto L10
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v123 == int32(0) {
		goto L2
	} else {
		goto L38
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v70+v75<<(uint(int32(2))%32))))
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v114 = v75 + int32(1)
	if v114 != v67 {
		v75 = v114
		goto L25
	} else {
		goto L37
	}
L28:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v87 == int32(0))|base.B2i32(v87 != v90) != 0 {
		v108 = v87
		v109 = v90
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v108-v109 != 0 {
		goto L27
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v93 = v82
	v94 = l0
	goto L32
L32:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v98 == int32(0) {
		v108 = v98
		v109 = v97
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v108 = v98
	v109 = v97
	goto L30
L34:
	;
	v101 = int32(1)
	if v98 == v97 {
		v93 = v93 + v101
		v94 = v94 + v101
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	return int32(0)
L37:
	;
	goto L26
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v126 <= int32(0) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v129 = int32(0)
	if v129 < v126 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v133 = v126
	goto L42
L41:
	;
	v133 = v129
	goto L42
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v138 = v129
	goto L43
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134+v138<<(uint(int32(2))%32))))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v148 == int32(0))|base.B2i32(v148 != v151) != 0 {
		v169 = v148
		v170 = v151
		goto L46
	} else {
		goto L47
	}
L44:
	;
	return int32(0)
L45:
	;
	if v169-v170 != 0 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	goto L45
L47:
	;
	v154 = v145
	v155 = l0
	goto L48
L48:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v159 == int32(0) {
		v169 = v159
		v170 = v158
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v169 = v159
	v170 = v158
	goto L46
L50:
	;
	v162 = int32(1)
	if v159 == v158 {
		v154 = v154 + v162
		v155 = v155 + v162
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v173 = v138 + int32(1)
	if v133 != v173 {
		v138 = v173
		goto L43
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L44
L55:
	;
	goto L2
L56:
	;
	return int32(0)
L57:
	;
	if v179 != 0 {
		v253 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L2
L59:
	;
	return int32(1)
L60:
	;
	goto L61
L61:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v196 <= int32(0) {
		v253 = int32(1)
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v202 = v196
	goto L65
L64:
	;
	v202 = v199
	goto L65
L65:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v208 = int32(0)
	goto L66
L66:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203+v208<<(uint(int32(2))%32))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v218 == int32(0))|base.B2i32(v218 != v221) != 0 {
		v239 = v218
		v240 = v221
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v253 = v243
	goto L1
L68:
	;
	v242 = int32(0)
	v243 = base.B2i32(v241 != v242)
	if v241 == v242 {
		v253 = v243
		goto L1
	} else {
		goto L75
	}
L69:
	;
	v241 = v239 - v240
	goto L68
L70:
	;
	v224 = v215
	v225 = l0
	goto L71
L71:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v229 == int32(0) {
		v239 = v229
		v240 = v228
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v239 = v229
	v240 = v228
	goto L69
L73:
	;
	v232 = int32(1)
	if v229 == v228 {
		v224 = v224 + v232
		v225 = v225 + v232
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v247 = v208 + int32(1)
	if v247 != v202 {
		v208 = v247
		goto L66
	} else {
		goto L76
	}
L76:
	;
	goto L67
}
func F_comp_trgm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_comp_trgm[0]))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_compact_trigram(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	v9 = int32(255)
	switch l2 {
	case 0:
		v61 = l2
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v61)
		v69 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v69)
		return
	default:
		v20 = l1
		v21 = l2
		v23 = v9
		v24 = v9
		v25 = v9
		v26 = v9
		for {
			v27 = int32(8)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v35 = *(*int32)(unsafe.Add(mBase, uint32((v29^v23)<<(uint(int32(2))%32))+uint32(_c_F_compact_trigram[0])))
			v38 = int32(16)
			v43 = int32(24)
			v46 = v35 ^ (v24<<(uint(v27)%32)&int32(_a_F_compact_trigram_0) | v25<<(uint(v38)%32)&int32(16711680) | v26<<(uint(v43)%32))
			v53 = int32(1)
			v56 = v21 - v53
			if v56 != 0 {
				v20 = v20 + v53
				v21 = v56
				v23 = int32(base.Ui32(v46) >> (uint(v43) % 32))
				v24 = v35
				v25 = int32(base.Ui32(v46) >> (uint(v27) % 32))
				v26 = int32(base.Ui32(v46) >> (uint(v38) % 32))
				continue
			} else {
				break
			}
			break
		}
		v61 = v46 ^ int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v61)
		v69 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v69)
		return
	case 3:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v15)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v17)
		return
	}
}
func F_compare_lexeme_textfreq(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v37 < v6 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v14 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = int32(1)
	if v8&v25 != 0 {
		v37 = int32(base.Ui32(v8)>>(uint(v25)%32)) - v25
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v17 = int32(16)
	goto L7
L6:
	;
	v17 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v14-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = int32(4)
	goto L10
L9:
	;
	v24 = v17
	goto L10
L10:
	;
	v37 = v24
	goto L1
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v37 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(1)
L13:
	;
	goto L14
L14:
	;
	if v6 < v37 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(-1)
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = int32(1)
	if v8&v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v45
	goto L20
L19:
	;
	v49 = int32(4)
	goto L20
L20:
	;
	v50 = v7 + v49
	if v6 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v95
L22:
	;
	v95 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = v44
	v58 = v50
	v59 = v6
	v60 = v56
	goto L29
L26:
	;
	v83 = v50
	v87 = int32(0)
	goto L27
L27:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v95 = v87 - v88
	goto L21
L28:
	;
	v83 = v78
	v87 = v80
	goto L27
L29:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if base.B2i32(v60 != v62)|base.B2i32(v62 == int32(0)) != 0 {
		v78 = v58
		v80 = v60
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v78 = v72
	v80 = int32(0)
	goto L28
L31:
	;
	v68 = v59 - int32(1)
	if v68 == int32(0) {
		v78 = v58
		v80 = v60
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v71 = int32(1)
	v72 = v58 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v73 != 0 {
		v57 = v57 + v71
		v58 = v72
		v59 = v68
		v60 = v73
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
}
func F_compare_scalars_simple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, v6, v7, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 < int32(0) {
			v16 = int32(1)
		} else {
			v16 = int32(0) - v9
		}
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if v17 != 0 {
			v18 = v16
		} else {
			v18 = v9
		}
		return v18
	}
}
func F_compute_bucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = base.I32_extend16_s(v19)
	v22 = base.B2i32(int32(0) <= v20)
	if int32(0) <= v20 {
		v23 = int32(-8)
	} else {
		v23 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v14)>>(uint(int32(2))%32))+v23) >> (uint(int32(1)) % 32))
	if int32(0) <= v20 {
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
		v38 = v28
	} else {
		v38 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v38
	v40 = int32(_a_F_compute_bucket_0)
	v41 = v19 & v40
	if v41 != v40 {
		if v41 != int32(_a_F_compute_bucket_1) {
			v52 = v41
		} else {
			v52 = v19 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v52 = v19 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v52
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v54
	v63 = base.B2i32(v20 < v54)
	if v20 < v54 {
		v64 = int32(base.Ui32(v19)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v64 = v19 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v64
	if v20 < v54 {
		v68 = int32(6)
	} else {
		v68 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l1 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v77 = base.I32_extend16_s(v76)
	v79 = base.B2i32(int32(0) <= v77)
	if int32(0) <= v77 {
		v80 = int32(-8)
	} else {
		v80 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v71)>>(uint(int32(2))%32))+v80) >> (uint(int32(1)) % 32))
	if int32(0) <= v77 {
		v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
		v95 = v85
	} else {
		v95 = v76<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v76&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v95
	v97 = int32(_a_F_compute_bucket_0)
	v98 = v76 & v97
	if v98 != v97 {
		if v98 != int32(_a_F_compute_bucket_1) {
			v109 = v98
		} else {
			v109 = v76 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v109 = v76 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v109
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v111
	v120 = base.B2i32(v77 < v111)
	if v77 < v111 {
		v121 = int32(base.Ui32(v76)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v121 = v76 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v121
	if v77 < v111 {
		v125 = int32(6)
	} else {
		v125 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l2 + v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v134 = base.I32_extend16_s(v133)
	v136 = base.B2i32(int32(0) <= v134)
	if int32(0) <= v134 {
		v137 = int32(-8)
	} else {
		v137 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(base.Ui32(int32(base.Ui32(v128)>>(uint(int32(2))%32))+v137) >> (uint(int32(1)) % 32))
	if int32(0) <= v134 {
		v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		v152 = v142
	} else {
		v152 = v133<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v133&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v152
	v154 = int32(_a_F_compute_bucket_0)
	v155 = v133 & v154
	if v155 != v154 {
		if v155 != int32(_a_F_compute_bucket_1) {
			v166 = v155
		} else {
			v166 = v133 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v166 = v133 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v166
	v168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v168
	v173 = base.B2i32(v134 < v168)
	if v134 < v168 {
		v174 = int32(6)
	} else {
		v174 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l0 + v174
	if v134 < v168 {
		v183 = int32(base.Ui32(v133)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v183 = v133 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v183
	v186 = v12 + int32(8)
	v188 = v12 + int32(56)
	F_sub_var(m, v186, v188, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		return
	} else {
		v192 = v12 + int32(32)
		F_sub_var(m, v192, v188, v192)
		mBase = m.M
		v194 = m.ExcPending
		if v194 != 0 {
			return
		} else {
			v195 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v196 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			F_mul_var(m, v186, l3, v186, v195+v196)
			mBase = m.M
			v199 = m.ExcPending
			if v199 != 0 {
				return
			} else {
				v200 = int32(0)
				F_div_var(m, v186, v192, l4, v200, v200, int32(1))
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return
				} else {
					F_add_var(m, l4, int32(_a_F_compute_bucket_5), l4)
					mBase = m.M
					v207 = m.ExcPending
					if v207 != 0 {
						return
					} else {
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						if v208 != 0 {
							F_pfree(m, v208)
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								if v211 != 0 {
									F_pfree(m, v211)
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return
									} else {
										m.G0 = v12 + int32(80)
										return
									}
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							}
						} else {
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							if v211 != 0 {
								F_pfree(m, v211)
								mBase = m.M
								v213 = m.ExcPending
								if v213 != 0 {
									return
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							} else {
								m.G0 = v12 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_compute_distinct_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 float64
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
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
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 float64
	_ = v434
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 float64
	_ = v448
	var v450 float32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 float64
	_ = v558
	var v562 float64
	_ = v562
	var v569 float64
	_ = v569
	var v570 float64
	_ = v570
	var v572 float64
	_ = v572
	var v578 float64
	_ = v578
	var v579 float64
	_ = v579
	var v582 float64
	_ = v582
	var v584 float64
	_ = v584
	var v617 float32
	_ = v617
	var v619 float64
	_ = v619
	var v625 float32
	_ = v625
	var v627 float32
	_ = v627
	var v630 int32
	_ = v630
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v828 float32
	_ = v828
	var v829 float64
	_ = v829
	var v830 float32
	_ = v830
	var v832 float64
	_ = v832
	var v838 int32
	_ = v838
	var v844 float64
	_ = v844
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v870 int32
	_ = v870
	var v873 float64
	_ = v873
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 float64
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v913 float64
	_ = v913
	var v928 int32
	_ = v928
	var v931 float64
	_ = v931
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v947 float64
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v960 float64
	_ = v960
	var v971 float64
	_ = v971
	var v978 int32
	_ = v978
	var v984 float64
	_ = v984
	var v995 float64
	_ = v995
	var v996 float64
	_ = v996
	var v1000 float64
	_ = v1000
	var v1003 float64
	_ = v1003
	var v1006 float64
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 float64
	_ = v1010
	var v1014 float64
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1020 float64
	_ = v1020
	var v1022 float64
	_ = v1022
	var v1028 float64
	_ = v1028
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1066 int32
	_ = v1066
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	v5 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(32)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
	if v34 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+76)))
	v40 = int32(_a_F_compute_distinct_stats_0)
	v45 = base.B2i32(v37 < int32(0))
	v46 = base.B2i32(v37&v40 == v40)
	goto L3
L2:
	;
	v45 = v5
	v46 = v5
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v48 = int32(10)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = v49 << (uint(int32(1)) % 32)
	if v51 <= v48 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v54 = v48
	goto L6
L5:
	;
	v54 = v51
	goto L6
L6:
	;
	v57 = F_palloc(m, v54<<(uint(int32(3))%32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	F_fmgr_info(m, v59, v31+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l2 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v31 + int32(32)
	return
L11:
	;
	v71 = v5
	v77 = v5
	v79 = v5
	v83 = v5
	v87 = v5
	v88 = float64(0)
	goto L12
L12:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	if int32(0) < v425 {
		goto L74
	} else {
		goto L75
	}
L14:
	;
	v99 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v87, v31+int32(3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
	if v101 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v441 = v87 + int32(1)
	if v441 != l2 {
		v71 = v417
		v77 = v423
		v79 = v425
		v83 = v429
		v87 = v441
		v88 = v434
		goto L12
	} else {
		goto L71
	}
L17:
	;
	v417 = v71
	v423 = v77 + int32(1)
	v425 = v79
	v429 = v83
	v434 = v88
	goto L16
L18:
	;
	goto L19
L19:
	;
	v107 = v79 + int32(1)
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v153 = int32(0)
	if v153 < v71 {
		goto L43
	} else {
		goto L44
	}
L21:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v108 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	if v45 == int32(0) {
		v151 = v99
		v152 = v88
		goto L20
	} else {
		goto L38
	}
L24:
	;
	v134 = base.F64_add(v88, base.F64_convert_i32_u(v132))
	v135 = F_toast_raw_datum_size(m, v99)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L33
	}
L25:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if base.Ui32((v112-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v132 = int32(6)
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v124 = int32(1)
	if v108&v124 != 0 {
		v132 = int32(base.Ui32(v108) >> (uint(v124) % 32))
		goto L24
	} else {
		goto L32
	}
L28:
	;
	v119 = int32(18)
	if v112 == v119 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v123 = v119
	goto L31
L30:
	;
	v123 = int32(2)
	goto L31
L31:
	;
	v132 = v123
	goto L24
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v132 = int32(base.Ui32(v128) >> (uint(int32(2)) % 32))
	goto L24
L33:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v135) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v417 = v71
	v423 = v77
	v425 = v107
	v429 = v83 + int32(1)
	v434 = v134
	goto L16
L35:
	;
	goto L36
L36:
	;
	v141 = F_pg_detoast_datum(m, v99)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v151 = v141
	v152 = v134
	goto L20
L38:
	;
	v145 = F_strlen(m, v99)
	mBase = m.M
	v151 = v99
	v152 = base.F64_add(v88, base.F64_convert_i32_u(v145+int32(1)))
	goto L20
L39:
	;
	if v209 < v231 {
		goto L68
	} else {
		goto L69
	}
L40:
	;
	if v209 == v230+v71-int32(2) {
		goto L39
	} else {
		goto L64
	}
L41:
	;
	v298 = int32(3)
	v300 = v57 + v233<<(uint(v298)%32)
	v303 = v57 + v231<<(uint(v298)%32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = v306
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v303-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v310
	v314 = v231 - int32(2)
	v316 = v233
	goto L40
L42:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v244 + int32(1)
	if v160 == int32(0) {
		v417 = v71
		v423 = v77
		v425 = v107
		v429 = v83
		v434 = v152
		goto L16
	} else {
		goto L59
	}
L43:
	;
	v160 = v153
	v163 = v71
	goto L46
L44:
	;
	v209 = v71
	goto L45
L45:
	;
	v230 = base.B2i32(v71 < v54)
	v231 = v71 + v230
	v233 = v231 - int32(1)
	if v233 <= v209 {
		goto L39
	} else {
		goto L57
	}
L46:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v189 = v57 + v160<<(uint(int32(3))%32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v191 = F_FunctionCall2Coll(m, v31+int32(4), v186, v151, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L48
	}
L47:
	;
	v209 = v198
	goto L45
L48:
	;
	if v191 != 0 {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	if v160 < v163 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v194 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v198 = v163
	goto L52
L52:
	;
	v200 = v160 + int32(1)
	if v200 != v71 {
		v160 = v200
		v163 = v198
		goto L46
	} else {
		goto L56
	}
L53:
	;
	v197 = v160
	goto L55
L54:
	;
	v197 = v163
	goto L55
L55:
	;
	v198 = v197
	goto L52
L56:
	;
	goto L47
L57:
	;
	if (v71-v230+v209)&int32(1) == int32(0) {
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v314 = v233
	v316 = v231
	goto L40
L59:
	;
	v254 = v160
	goto L60
L60:
	;
	v280 = v57 + v254<<(uint(int32(3))%32)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v283 = v280 - int32(4)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v281 <= v284 {
		v417 = v71
		v423 = v77
		v425 = v107
		v429 = v83
		v434 = v152
		goto L16
	} else {
		goto L62
	}
L61:
	;
	v417 = v71
	v423 = v77
	v425 = v107
	v429 = v83
	v434 = v152
	goto L16
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v284
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v289 = v280 - int32(8)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v281
	v294 = int32(1)
	if v294 < v254 {
		v254 = v254 - v294
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v322 = v314
	v328 = v316
	goto L65
L65:
	;
	v346 = int32(3)
	v348 = v57 + v322<<(uint(v346)%32)
	v351 = v57 + v328<<(uint(v346)%32)
	v352 = int32(16)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351-v352)))
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v354
	v356 = int32(12)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v351-v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v348)+4)) = v358
	v361 = v322 - int32(1)
	v364 = v57 + v361<<(uint(v346)%32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v348-v352)))
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v367
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v348-v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v371
	v374 = v322 - int32(2)
	if v209 < v374 {
		v322 = v374
		v328 = v361
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L39
L67:
	;
	goto L66
L68:
	;
	v407 = v57 + v209<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v151
	goto L70
L69:
	;
	goto L70
L70:
	;
	v417 = v231
	v423 = v77
	v425 = v107
	v429 = v83
	v434 = v152
	goto L16
L71:
	;
	goto L13
L72:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v617
	v619 = base.F64_promote_f32(v617)
	if base.F64_gt(v619, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L104
	} else {
		goto L105
	}
L73:
	;
	if base.B2i32(v54 <= v417)|v429|base.B2i32(v547 != v417) == int32(0) {
		v617 = base.F32_convert_i32_s(v417)
		goto L72
	} else {
		goto L93
	}
L74:
	;
	v445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v445)
	v448 = base.F64_convert_i32_s(l2)
	v450 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v423), v448))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v450
	if v45 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v423 <= int32(0) {
		goto L10
	} else {
		goto L89
	}
L77:
	;
	v457 = base.I32_trunc_sat_f64_s(base.F64_div(v434, base.F64_convert_i32_u(v425)))
	goto L79
L78:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v456 = int32(*(*int16)(unsafe.Add(mBase, uint32(v455)+76)))
	v457 = v456
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v457
	if int32(0) < v417 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v461 = int32(0)
	v467 = v461
	v470 = v461
	goto L84
L81:
	;
	goto L82
L82:
	;
	v617 = base.F32_neg(base.F32_sub(float32(1), v450))
	goto L72
L83:
	;
	if v467 != 0 {
		v547 = v467
		v548 = v470
		goto L73
	} else {
		goto L88
	}
L84:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v57+v467<<(uint(int32(3))%32))+4))
	if v494 == int32(1) {
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v547 = v417
	v548 = v497
	goto L73
L86:
	;
	v497 = v494 + v470
	v499 = v467 + int32(1)
	if v499 != v417 {
		v467 = v499
		v470 = v497
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	goto L82
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v536)
	v538 = int32(0)
	if v45 == v538 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v542 = int32(*(*int16)(unsafe.Add(mBase, uint32(v541)+76)))
	v543 = v542
	goto L92
L91:
	;
	v543 = v538
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v543
	goto L10
L93:
	;
	v556 = v425 - v548
	v557 = v547 + v556
	v558 = float64(0)
	v562 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v450)))
	if base.F64_gt(v562, v558) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if base.F64_lt(v578, v579) != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v578 = v558
	v579 = base.F64_convert_i32_s(v557)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v569 = base.F64_convert_i32_s(l2 - v423)
	v570 = base.F64_convert_i32_s(v557)
	v572 = base.F64_convert_i32_s(v556)
	v578 = base.F64_div(base.F64_mul(v569, v570), base.F64_add(base.F64_sub(v569, v572), base.F64_div(base.F64_mul(v569, v572), v562)))
	v579 = v570
	goto L94
L98:
	;
	v582 = v579
	goto L100
L99:
	;
	v582 = v578
	goto L100
L100:
	;
	if base.F64_gt(v582, v562) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v584 = v562
	goto L103
L102:
	;
	v584 = v582
	goto L103
L103:
	;
	v617 = base.F32_demote_f64(base.F64_floor(base.F64_add(v584, float64(0.5))))
	goto L72
L104:
	;
	v625 = base.F32_demote_f64(base.F64_div(base.F64_neg(v619), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v625
	v627 = v625
	goto L106
L105:
	;
	v627 = v617
	goto L106
L106:
	;
	v630 = int32(0)
	if base.B2i32(base.B2i32(base.F32_gt(v627, float32(0)) == v630)|(base.B2i32(v54 <= v417)|v429) == v630)&base.B2i32(v417 <= v49) == v630 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v49 < v417 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v1066 = v417
	goto L109
L109:
	;
	if v1066 <= int32(0) {
		goto L10
	} else {
		goto L159
	}
L110:
	;
	v642 = v49
	goto L112
L111:
	;
	v642 = v417
	goto L112
L112:
	;
	if v642 <= int32(0) {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v646 = v642 & int32(3)
	v650 = F_palloc(m, v642<<(uint(int32(2))%32))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	v652 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v642) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v828 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v829 = base.F64_promote_f32(v828)
	v830 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v832 = float64(0)
	v838 = int32(0)
	v844 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v844)|base.F64_le(l3, float64(1)) != 0 {
		v1044 = v642
		goto L127
	} else {
		goto L128
	}
L116:
	;
	v662 = v652
	v666 = int32(0)
	goto L119
L117:
	;
	v735 = v652
	goto L118
L118:
	;
	v763 = v735
	v764 = int32(0)
	goto L123
L119:
	;
	v686 = int32(2)
	v689 = int32(3)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v57+v662<<(uint(v689)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v650+v662<<(uint(v686)%32)))) = v692
	v695 = v662 | int32(1)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v57+v695<<(uint(v689)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v650+v695<<(uint(v686)%32)))) = v702
	v705 = v662 | v686
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v57+v705<<(uint(v689)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v650+v705<<(uint(v686)%32)))) = v712
	v715 = v662 | v689
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v57+v715<<(uint(v689)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v650+v715<<(uint(v686)%32)))) = v722
	v724 = int32(4)
	v725 = v662 + v724
	v727 = v666 + v724
	if v727 != v642&int32(2147483644) {
		v662 = v725
		v666 = v727
		goto L119
	} else {
		goto L121
	}
L120:
	;
	if v646 == int32(0) {
		goto L115
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v735 = v725
	goto L118
L123:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v57+v763<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v650+v763<<(uint(int32(2))%32)))) = v793
	v795 = int32(1)
	v798 = v764 + v795
	if v798 != v646 {
		v763 = v763 + v795
		v764 = v798
		goto L123
	} else {
		goto L125
	}
L124:
	;
	goto L115
L125:
	;
	goto L124
L126:
	;
	v1066 = v1044
	goto L109
L127:
	;
	goto L126
L128:
	;
	if base.Ui32(v642) < base.Ui32(int32(2)) {
		v960 = v832
		goto L129
	} else {
		goto L130
	}
L129:
	;
	if base.F64_lt(v829, float64(0)) != 0 {
		goto L141
	} else {
		goto L142
	}
L130:
	;
	v856 = v642 - int32(1)
	v857 = int32(3)
	v858 = v856 & v857
	v859 = int32(0)
	if base.Ui32(v857) <= base.Ui32(v642-int32(2)) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v870 = v859
	v873 = v832
	v882 = v838
	goto L134
L132:
	;
	v910 = v859
	v913 = v832
	goto L133
L133:
	;
	v928 = v910
	v931 = v913
	v941 = v838
	goto L138
L134:
	;
	v886 = v650 + v870<<(uint(int32(2))%32)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v886)+4))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v886)+8))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v886)+12))
	v898 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v873, base.F64_convert_i32_s(v887)), base.F64_convert_i32_s(v890)), base.F64_convert_i32_s(v893)), base.F64_convert_i32_s(v896))
	v899 = int32(4)
	v900 = v870 + v899
	v902 = v882 + v899
	if v902 != v856&int32(-4) {
		v870 = v900
		v873 = v898
		v882 = v902
		goto L134
	} else {
		goto L136
	}
L135:
	;
	if v858 == int32(0) {
		v960 = v898
		goto L129
	} else {
		goto L137
	}
L136:
	;
	goto L135
L137:
	;
	v910 = v900
	v913 = v898
	goto L133
L138:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v650+v928<<(uint(int32(2))%32))))
	v947 = base.F64_add(v931, base.F64_convert_i32_s(v945))
	v948 = int32(1)
	v951 = v941 + v948
	if v951 != v858 {
		v928 = v928 + v948
		v931 = v947
		v941 = v951
		goto L138
	} else {
		goto L140
	}
L139:
	;
	v960 = v947
	goto L129
L140:
	;
	goto L139
L141:
	;
	v971 = base.F64_mul(l3, base.F64_neg(v829))
	goto L143
L142:
	;
	v971 = v829
	goto L143
L143:
	;
	v978 = v642
	v984 = v960
	goto L144
L144:
	;
	v995 = float64(1)
	v996 = float64(0)
	v1000 = base.F64_sub(base.F64_sub(v995, base.F64_div(v984, v844)), base.F64_promote_f32(v830))
	if base.F64_lt(v1000, v996) != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v1044 = int32(0)
	goto L127
L146:
	;
	v1003 = v996
	goto L148
L147:
	;
	v1003 = v1000
	goto L148
L148:
	;
	if base.F64_gt(v1003, float64(1)) != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v1006 = v995
	goto L151
L150:
	;
	v1006 = v1003
	goto L151
L151:
	;
	v1008 = v978 - int32(1)
	v1010 = base.F64_sub(v971, base.F64_convert_i32_u(v1008))
	if base.F64_gt(v1010, float64(1)) != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v1014 = base.F64_div(v1006, v1010)
	goto L154
L153:
	;
	v1014 = v1006
	goto L154
L154:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v650+v1008<<(uint(int32(2))%32))))
	v1020 = base.F64_convert_i32_s(v1019)
	v1022 = base.F64_div(base.F64_mul(l3, v1020), v844)
	v1028 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v844), base.F64_mul(base.F64_mul(v1022, v844), base.F64_sub(l3, v1022))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v1014, v844), base.F64_add(v1028, v1028)), float64(0.5)), v1020) != 0 {
		v1044 = v978
		goto L127
	} else {
		goto L155
	}
L155:
	;
	if v1008 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v650+v978<<(uint(int32(2))%32)-int32(8))))
	v978 = v1008
	v984 = base.F64_sub(v984, base.F64_convert_i32_s(v1039))
	goto L144
L157:
	;
	goto L158
L158:
	;
	goto L145
L159:
	;
	v1092 = int32(_a_F_compute_distinct_stats_1)
	v1093 = *(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0]))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1095
	v1098 = v1066 << (uint(int32(2)) % 32)
	v1099 = F_palloc(m, v1098)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L7
	} else {
		goto L160
	}
L160:
	;
	v1101 = F_palloc(m, v1098)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	v1107 = int32(0)
	goto L162
L162:
	;
	v1132 = v1107 << (uint(int32(2)) % 32)
	v1136 = v57 + v1107<<(uint(int32(3))%32)
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1136)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138)+78)))
	v1140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1138)+76)))
	v1141 = F_datumCopy(m, v1137, v1139, v1140)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L7
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1093
	v1155 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1155)
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1101
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v1066
	goto L10
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1099+v1132))) = v1141
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v1101+v1132))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1145), v448))
	v1151 = v1107 + int32(1)
	if v1151 != v1066 {
		v1107 = v1151
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
}
func F_compute_scalar_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 float64
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 float64
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 float64
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 float64
	_ = v175
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v237 float64
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var __phi280 int32
	_ = __phi280
	var v283 int32
	_ = v283
	var __phi283 int32
	_ = __phi283
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v360 int32
	_ = v360
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	var v428 float64
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 float64
	_ = v435
	var v437 float32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v456 int32
	_ = v456
	var v457 float64
	_ = v457
	var v461 float64
	_ = v461
	var v468 float64
	_ = v468
	var v469 float64
	_ = v469
	var v473 float64
	_ = v473
	var v479 float64
	_ = v479
	var v480 float64
	_ = v480
	var v483 float64
	_ = v483
	var v485 float64
	_ = v485
	var v495 float32
	_ = v495
	var v497 float64
	_ = v497
	var v503 float32
	_ = v503
	var v505 float32
	_ = v505
	var v508 int32
	_ = v508
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v717 float32
	_ = v717
	var v718 float64
	_ = v718
	var v719 float32
	_ = v719
	var v721 float64
	_ = v721
	var v727 int32
	_ = v727
	var v733 float64
	_ = v733
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v759 int32
	_ = v759
	var v762 float64
	_ = v762
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v787 float64
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v802 float64
	_ = v802
	var v817 int32
	_ = v817
	var v820 float64
	_ = v820
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 float64
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v849 float64
	_ = v849
	var v860 float64
	_ = v860
	var v867 int32
	_ = v867
	var v873 float64
	_ = v873
	var v884 float64
	_ = v884
	var v885 float64
	_ = v885
	var v889 float64
	_ = v889
	var v892 float64
	_ = v892
	var v895 float64
	_ = v895
	var v897 int32
	_ = v897
	var v899 float64
	_ = v899
	var v903 float64
	_ = v903
	var v908 int32
	_ = v908
	var v909 float64
	_ = v909
	var v911 float64
	_ = v911
	var v917 float64
	_ = v917
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v951 int32
	_ = v951
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1180 int32
	_ = v1180
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1302 int32
	_ = v1302
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 float64
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1348 float64
	_ = v1348
	var v1350 float64
	_ = v1350
	var v1352 float64
	_ = v1352
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1384 int32
	_ = v1384
	var v1389 float32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	v5 = int32(0)
	v25 = float64(0)
	v32 = m.G0
	v34 = v32 - int32(48)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+78)))
	if v37 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+76)))
	v43 = int32(_a_F_compute_scalar_stats_0)
	v48 = base.B2i32(v40&v43 == v43)
	v49 = base.B2i32(v40 < int32(0))
	goto L3
L2:
	;
	v48 = v5
	v49 = v5
	goto L3
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v54 = F_palloc(m, l2<<(uint(int32(3))%32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v58 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v62 = F_palloc(m, v50<<(uint(int32(3))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v64
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+36)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v34)+28)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v34)+20)) = v66
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)) = uint8(v64)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	F_PrepareSortSupportFromOrderingOp(m, v79, v34+int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if l2 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v34 + int32(48)
	return
L10:
	;
	v91 = v5
	v95 = v5
	v97 = v5
	v103 = v5
	v107 = v5
	v110 = v25
	goto L11
L11:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	if int32(0) < v190 {
		goto L39
	} else {
		goto L40
	}
L13:
	;
	v122 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v91, v34+int32(4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v124 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v195 = v91 + int32(1)
	if v195 != l2 {
		v91 = v195
		v95 = v189
		v97 = v190
		v103 = v191
		v107 = v192
		v110 = v193
		goto L11
	} else {
		goto L38
	}
L16:
	;
	v189 = v95 + int32(1)
	v190 = v97
	v191 = v103
	v192 = v107
	v193 = v110
	goto L15
L17:
	;
	goto L18
L18:
	;
	v130 = v107 + int32(1)
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v178 = v54 + v97<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v58+v97<<(uint(int32(2))%32)))) = v97
	v189 = v95
	v190 = v97 + int32(1)
	v191 = v103
	v192 = v130
	v193 = v175
	goto L15
L20:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v131 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	if v49 == int32(0) {
		v173 = v122
		v175 = v110
		goto L19
	} else {
		goto L37
	}
L23:
	;
	v157 = base.F64_add(v110, base.F64_convert_i32_u(v155))
	v158 = F_toast_raw_datum_size(m, v122)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if base.Ui32((v135-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v155 = int32(6)
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v147 = int32(1)
	if v131&v147 != 0 {
		v155 = int32(base.Ui32(v131) >> (uint(v147) % 32))
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v142 = int32(18)
	if v135 == v142 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v146 = v142
	goto L30
L29:
	;
	v146 = int32(2)
	goto L30
L30:
	;
	v155 = v146
	goto L23
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v155 = int32(base.Ui32(v151) >> (uint(int32(2)) % 32))
	goto L23
L32:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v158) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v189 = v95
	v190 = v97
	v191 = v103 + int32(1)
	v192 = v130
	v193 = v157
	goto L15
L34:
	;
	goto L35
L35:
	;
	v164 = F_pg_detoast_datum(m, v122)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v173 = v164
	v175 = v157
	goto L19
L37:
	;
	v168 = F_strlen(m, v122)
	mBase = m.M
	v173 = v122
	v175 = base.F64_add(v110, base.F64_convert_i32_u(v168+int32(1)))
	goto L19
L38:
	;
	goto L12
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v34 + int32(12)
	F_qsort_interruptible(m, v54, v190, int32(8), int32(509), v34+int32(4))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if int32(0) < v192 {
		goto L175
	} else {
		goto L176
	}
L42:
	;
	v211 = int32(0)
	v222 = v5
	v223 = v5
	v224 = v5
	v228 = v5
	v237 = v25
	goto L43
L43:
	;
	v242 = v222 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v54+v211<<(uint(int32(3))%32))+4))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v58+v247<<(uint(int32(2))%32))))
	if v247 != v253 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v432 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v432)
	v435 = base.F64_convert_i32_s(l2)
	v437 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v189), v435))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v437
	if v49 != 0 {
		goto L61
	} else {
		goto L62
	}
L45:
	;
	v409 = v223
	v410 = v224
	v414 = v228
	v427 = v242
	goto L47
L46:
	;
	if v242 < int32(2) {
		v378 = v224
		v382 = v228
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v428 = base.F64_add(base.F64_mul(base.F64_convert_i32_u(v211), base.F64_convert_i32_s(v247)), v237)
	v430 = v211 + int32(1)
	if v430 != v190 {
		v211 = v430
		v222 = v427
		v223 = v409
		v224 = v410
		v228 = v414
		v237 = v428
		goto L43
	} else {
		goto L60
	}
L48:
	;
	v409 = v223 + int32(1)
	v410 = v378
	v414 = v382
	v427 = int32(0)
	goto L47
L49:
	;
	v260 = v228 + int32(1)
	v261 = base.B2i32(v224 < v50)
	if v261 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v62+v224<<(uint(int32(3))%32)-int32(8))))
	if v242 <= v269 {
		v378 = v224
		v382 = v260
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v271 = v261 + v224
	v273 = v271 - int32(1)
	if v273 <= int32(0) {
		v331 = v273
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v360 = v62 + v331<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+4)) = v211 - v222
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v242
	v378 = v271
	v382 = v260
	goto L48
L55:
	;
	__phi280 = v273
	__phi283 = v271
	v280 = __phi280
	v283 = __phi283
	goto L56
L56:
	;
	v309 = v62 + v283<<(uint(int32(3))%32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309-int32(16))))
	if v242 <= v312 {
		v331 = v280
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v331 = int32(0)
	goto L54
L58:
	;
	v316 = v62 + v280<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v312
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v309-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+4)) = v320
	v322 = int32(1)
	if v322 < v280 {
		__phi280 = v280 - v322
		__phi283 = v280
		v280 = __phi280
		v283 = __phi283
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L44
L61:
	;
	v444 = base.I32_trunc_sat_f64_s(base.F64_div(v193, base.F64_convert_i32_s(v192)))
	goto L63
L62:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v442)+76)))
	v444 = v443
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v444
	if v414 == int32(0) {
		v495 = base.F32_neg(base.F32_sub(float32(1), v437))
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v495
	v497 = base.F64_promote_f32(v495)
	if base.F64_gt(v497, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	if v191|base.B2i32(v409 != v414) == int32(0) {
		v495 = base.F32_convert_i32_s(v414)
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v456 = v409 + v191
	v457 = float64(0)
	v461 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v437)))
	if base.F64_gt(v461, v457) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if base.F64_lt(v479, v480) != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v479 = v457
	v480 = base.F64_convert_i32_s(v456)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v468 = base.F64_convert_i32_s(l2 - v189)
	v469 = base.F64_convert_i32_s(v456)
	v473 = base.F64_convert_i32_s(v191 - v414 + v409)
	v479 = base.F64_div(base.F64_mul(v468, v469), base.F64_add(base.F64_sub(v468, v473), base.F64_div(base.F64_mul(v468, v473), v461)))
	v480 = v469
	goto L67
L71:
	;
	v483 = v480
	goto L73
L72:
	;
	v483 = v479
	goto L73
L73:
	;
	if base.F64_gt(v483, v461) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v485 = v461
	goto L76
L75:
	;
	v485 = v483
	goto L76
L76:
	;
	v495 = base.F32_demote_f64(base.F64_floor(base.F64_add(v485, float64(0.5))))
	goto L64
L77:
	;
	v503 = base.F32_demote_f64(base.F64_div(base.F64_neg(v497), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v503
	v505 = v503
	goto L79
L78:
	;
	v505 = v495
	goto L79
L79:
	;
	v508 = int32(0)
	if base.B2i32(base.F32_gt(v505, float32(0)) == v508)|(base.B2i32(v409 != v410)|v191) == v508 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v1100 = v409 - v1067
	if v50 < v1100 {
		goto L142
	} else {
		goto L143
	}
L81:
	;
	v981 = int32(0)
	if v951 <= v981 {
		v1067 = v951
		v1097 = v981
		goto L80
	} else {
		goto L135
	}
L82:
	;
	if v409 <= v50 {
		v951 = v409
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v50 < v410 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v519 = v50
	goto L88
L87:
	;
	v519 = v410
	goto L88
L88:
	;
	if v519 <= int32(0) {
		v1067 = v519
		v1097 = int32(0)
		goto L80
	} else {
		goto L89
	}
L89:
	;
	v523 = v519 & int32(3)
	v527 = F_palloc(m, v519<<(uint(int32(2))%32))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v529 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v519) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v717 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v718 = base.F64_promote_f32(v717)
	v719 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v721 = float64(0)
	v727 = int32(0)
	v733 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v733)|base.F64_le(l3, float64(1)) != 0 {
		v933 = v519
		goto L103
	} else {
		goto L104
	}
L92:
	;
	v540 = v529
	v542 = int32(0)
	goto L95
L93:
	;
	v616 = v529
	goto L94
L94:
	;
	v647 = v616
	v648 = int32(0)
	goto L99
L95:
	;
	v566 = int32(2)
	v569 = int32(3)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v62+v540<<(uint(v569)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v527+v540<<(uint(v566)%32)))) = v572
	v575 = v540 | int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v62+v575<<(uint(v569)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v527+v575<<(uint(v566)%32)))) = v582
	v585 = v540 | v566
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v62+v585<<(uint(v569)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v527+v585<<(uint(v566)%32)))) = v592
	v595 = v540 | v569
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v62+v595<<(uint(v569)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v527+v595<<(uint(v566)%32)))) = v602
	v604 = int32(4)
	v605 = v540 + v604
	v607 = v542 + v604
	if v607 != v519&int32(2147483644) {
		v540 = v605
		v542 = v607
		goto L95
	} else {
		goto L97
	}
L96:
	;
	if v523 == int32(0) {
		goto L91
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	v616 = v605
	goto L94
L99:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v62+v647<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v527+v647<<(uint(int32(2))%32)))) = v679
	v681 = int32(1)
	v684 = v648 + v681
	if v684 != v523 {
		v647 = v647 + v681
		v648 = v684
		goto L99
	} else {
		goto L101
	}
L100:
	;
	goto L91
L101:
	;
	goto L100
L102:
	;
	v951 = v933
	goto L81
L103:
	;
	goto L102
L104:
	;
	if base.Ui32(v519) < base.Ui32(int32(2)) {
		v849 = v721
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if base.F64_lt(v718, float64(0)) != 0 {
		goto L117
	} else {
		goto L118
	}
L106:
	;
	v745 = v519 - int32(1)
	v746 = int32(3)
	v747 = v745 & v746
	v748 = int32(0)
	if base.Ui32(v746) <= base.Ui32(v519-int32(2)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v759 = v748
	v762 = v721
	v771 = v727
	goto L110
L108:
	;
	v799 = v748
	v802 = v721
	goto L109
L109:
	;
	v817 = v799
	v820 = v802
	v830 = v727
	goto L114
L110:
	;
	v775 = v527 + v759<<(uint(int32(2))%32)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v775)+8))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v775)+12))
	v787 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v762, base.F64_convert_i32_s(v776)), base.F64_convert_i32_s(v779)), base.F64_convert_i32_s(v782)), base.F64_convert_i32_s(v785))
	v788 = int32(4)
	v789 = v759 + v788
	v791 = v771 + v788
	if v791 != v745&int32(-4) {
		v759 = v789
		v762 = v787
		v771 = v791
		goto L110
	} else {
		goto L112
	}
L111:
	;
	if v747 == int32(0) {
		v849 = v787
		goto L105
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v799 = v789
	v802 = v787
	goto L109
L114:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v527+v817<<(uint(int32(2))%32))))
	v836 = base.F64_add(v820, base.F64_convert_i32_s(v834))
	v837 = int32(1)
	v840 = v830 + v837
	if v840 != v747 {
		v817 = v817 + v837
		v820 = v836
		v830 = v840
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v849 = v836
	goto L105
L116:
	;
	goto L115
L117:
	;
	v860 = base.F64_mul(l3, base.F64_neg(v718))
	goto L119
L118:
	;
	v860 = v718
	goto L119
L119:
	;
	v867 = v519
	v873 = v849
	goto L120
L120:
	;
	v884 = float64(1)
	v885 = float64(0)
	v889 = base.F64_sub(base.F64_sub(v884, base.F64_div(v873, v733)), base.F64_promote_f32(v719))
	if base.F64_lt(v889, v885) != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v933 = int32(0)
	goto L103
L122:
	;
	v892 = v885
	goto L124
L123:
	;
	v892 = v889
	goto L124
L124:
	;
	if base.F64_gt(v892, float64(1)) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v895 = v884
	goto L127
L126:
	;
	v895 = v892
	goto L127
L127:
	;
	v897 = v867 - int32(1)
	v899 = base.F64_sub(v860, base.F64_convert_i32_u(v897))
	if base.F64_gt(v899, float64(1)) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v903 = base.F64_div(v895, v899)
	goto L130
L129:
	;
	v903 = v895
	goto L130
L130:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v527+v897<<(uint(int32(2))%32))))
	v909 = base.F64_convert_i32_s(v908)
	v911 = base.F64_div(base.F64_mul(l3, v909), v733)
	v917 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v733), base.F64_mul(base.F64_mul(v911, v733), base.F64_sub(l3, v911))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v903, v733), base.F64_add(v917, v917)), float64(0.5)), v909) != 0 {
		v933 = v867
		goto L103
	} else {
		goto L131
	}
L131:
	;
	if v897 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v527+v867<<(uint(int32(2))%32)-int32(8))))
	v867 = v897
	v873 = base.F64_sub(v873, base.F64_convert_i32_s(v928))
	goto L120
L133:
	;
	goto L134
L134:
	;
	goto L121
L135:
	;
	v985 = int32(_a_F_compute_scalar_stats_1)
	v986 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v988
	v991 = v951 << (uint(int32(2)) % 32)
	v992 = F_palloc(m, v991)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v994 = F_palloc(m, v991)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v1001 = int32(0)
	goto L138
L138:
	;
	v1028 = v1001 << (uint(int32(2)) % 32)
	v1030 = int32(3)
	v1032 = v62 + v1001<<(uint(v1030)%32)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+4))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v54+v1033<<(uint(v1030)%32))))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038)+78)))
	v1040 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1038)+76)))
	v1041 = F_datumCopy(m, v1037, v1039, v1040)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L4
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v986
	v1055 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1055)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1057
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v994
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1060
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v951
	v1067 = v951
	v1097 = v1055
	goto L80
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v992+v1028))) = v1041
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	*(*float32)(unsafe.Add(mBase, uint32(v994+v1028))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1045), v435))
	v1051 = v1001 + int32(1)
	if v1051 != v951 {
		v1001 = v1051
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v1102 = v50 + int32(1)
	goto L144
L143:
	;
	v1102 = v1100
	goto L144
L144:
	;
	if int32(2) <= v1102 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v1105 = int32(0)
	F_qsort_interruptible(m, v62, v1067, int32(8), int32(510), v1105)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L4
	} else {
		goto L148
	}
L146:
	;
	v1302 = v1097
	goto L147
L147:
	;
	if v190 == int32(1) {
		goto L9
	} else {
		goto L173
	}
L148:
	;
	if v1097 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v1111 = int32(0)
	v1118 = v1105
	v1119 = v1111
	v1120 = v1111
	goto L152
L150:
	;
	v1180 = v190
	goto L151
L151:
	;
	v1205 = int32(_a_F_compute_scalar_stats_1)
	v1206 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1208
	v1210 = int32(1)
	v1211 = v1180 - v1210
	v1213 = v1102 - v1210
	v1214 = base.I32_div_s(v1211, v1213)
	if v1102 <= v1210 {
		goto L162
	} else {
		goto L163
	}
L152:
	;
	if v1067 <= v1120 {
		v1154 = v190
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v1180 = v1168
	goto L151
L154:
	;
	if v1172 < v190 {
		v1118 = v1172
		v1119 = v1168
		v1120 = v1169
		goto L152
	} else {
		goto L161
	}
L155:
	;
	v1156 = v1154 - v1118
	v1158 = v1156 << (uint(int32(3)) % 32)
	if v1158 != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v1147 = v62 + v1120<<(uint(int32(3))%32)
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	if v1118 < v1148 {
		v1154 = v1148
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1147)))
	v1168 = v1119
	v1169 = v1120 + int32(1)
	v1172 = v1152 + v1148
	goto L154
L158:
	;
	v1159 = int32(3)
	base.MemoryCopy(m, v54+v1119<<(uint(v1159)%32), v54+v1118<<(uint(v1159)%32), v1158)
	goto L160
L159:
	;
	goto L160
L160:
	;
	v1168 = v1119 + v1156
	v1169 = v1120
	v1172 = v1154
	goto L154
L161:
	;
	goto L153
L162:
	;
	v1220 = v1210
	goto L164
L163:
	;
	v1220 = v1102
	goto L164
L164:
	;
	v1223 = F_palloc(m, v1102<<(uint(int32(2))%32))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v1225 = int32(0)
	v1229 = v1225
	v1233 = v1225
	v1234 = v1225
	goto L166
L166:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v54+v1234<<(uint(int32(3))%32))))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266)+78)))
	v1268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1266)+76)))
	v1269 = F_datumCopy(m, v1265, v1267, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L4
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1206
	v1284 = int32(1)
	v1287 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1097<<(uint(v1284)%32))+52)) = uint16(v1287)
	v1291 = l0 + v1097<<(uint(v1287)%32)
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+64)) = v1292
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+164)) = v1223
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+84)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+144)) = v1102
	v1302 = v1097 + v1284
	goto L147
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1223+v1233<<(uint(int32(2))%32)))) = v1269
	v1272 = v1229 + (v1211 - v1214*v1213)
	v1273 = base.B2i32(v1213 <= v1272)
	if v1213 <= v1272 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1277 = v1213
	goto L171
L170:
	;
	v1277 = int32(0)
	goto L171
L171:
	;
	v1280 = v1233 + int32(1)
	if v1280 != v1220 {
		v1229 = v1272 - v1277
		v1233 = v1280
		v1234 = v1273 + (v1234 + v1214)
		goto L166
	} else {
		goto L172
	}
L172:
	;
	goto L167
L173:
	;
	v1333 = int32(_a_F_compute_scalar_stats_1)
	v1334 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1336
	v1339 = F_palloc(m, int32(4))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1334
	v1343 = base.F64_convert_i32_u(v190)
	v1345 = int32(1)
	v1348 = base.F64_mul(v1343, base.F64_convert_i32_u(v190-v1345))
	v1350 = base.F64_mul(v1348, float64(0.5))
	v1352 = base.F64_mul(v1350, base.F64_neg(v1350))
	*(*float32)(unsafe.Add(mBase, uint32(v1339))) = base.F32_demote_f64(base.F64_div(base.F64_add(base.F64_mul(v1343, v428), v1352), base.F64_add(base.F64_mul(v1343, base.F64_div(base.F64_mul(v1348, base.F64_convert_i32_s(v190<<(uint(v1345)%32)-v1345)), float64(6))), v1352)))
	v1370 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1302<<(uint(v1345)%32))+52)) = uint16(v1370)
	v1374 = l0 + v1302<<(uint(int32(2))%32)
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+64)) = v1375
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+124)) = v1339
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+84)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+104)) = v1345
	goto L9
L175:
	;
	v1384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1384)
	v1389 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v189), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v1389
	if v49 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	if v189 <= int32(0) {
		goto L9
	} else {
		goto L181
	}
L178:
	;
	v1396 = base.I32_trunc_sat_f64_s(base.F64_div(v193, base.F64_convert_i32_u(v192)))
	goto L180
L179:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1394)+76)))
	v1396 = v1395
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1396
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v1389))
	goto L9
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v1406 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1406)
	v1408 = int32(0)
	if v49 == v1408 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1412 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1411)+76)))
	v1413 = v1412
	goto L184
L183:
	;
	v1413 = v1408
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1413
	goto L9
}
func F_connectby_text_serial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = F_text_to_cstring(m, v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v29 = F_pg_detoast_datum_packed(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = F_text_to_cstring(m, v29)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v34 = F_pg_detoast_datum_packed(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = F_text_to_cstring(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v39 = F_pg_detoast_datum_packed(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = F_text_to_cstring(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v44 = F_pg_detoast_datum_packed(m, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = F_text_to_cstring(m, v44)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v48 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_connectby_text_serial_0), int32(0))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1076), int32(_a_F_connectby_text_serial_2))
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
												if v51 != int32(383) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_connectby_text_serial_0), int32(0))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1076), int32(_a_F_connectby_text_serial_2))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
													if v54&int32(2) == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_connectby_text_serial_3), int32(0))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1081), int32(_a_F_connectby_text_serial_2))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
														if v59 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(1088))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_connectby_text_serial_3), int32(0))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1081), int32(_a_F_connectby_text_serial_2))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														} else {
															v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
															v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															if v63 == int32(7) {
																v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																v67 = F_pg_detoast_datum_packed(m, v66)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int32(0)
																} else {
																	v69 = F_text_to_cstring(m, v67)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return int32(0)
																	} else {
																		v74 = v69
																		v75 = int32(_a_F_connectby_text_serial_4)
																		v76 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																		v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																		*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																		v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
																		v82 = F_CreateTupleDescCopy(m, v81)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return int32(0)
																		} else {
																			v85 = base.B2i32(v63 == int32(7))
																			F_validateConnectbyTupleDesc(m, v82, v85, int32(1))
																			mBase = m.M
																			v88 = m.ExcPending
																			if v88 != 0 {
																				return int32(0)
																			} else {
																				v89 = F_TupleDescGetAttInMetadata(m, v82)
																				mBase = m.M
																				v90 = m.ExcPending
																				if v90 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(2)
																					v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
																					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(1)
																					F_SPI_connect_ext(m, int32(0))
																					mBase = m.M
																					v98 = m.ExcPending
																					if v98 != 0 {
																						return int32(0)
																					} else {
																						v99 = int32(_a_F_connectby_text_serial_4)
																						v100 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																						v109 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[1]))
																						v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																						mBase = m.M
																						v111 = m.ExcPending
																						if v111 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v100
																							F_build_tuplestore_recursively(m, v31, v36, v26, v41, v74, v46, v46, int32(0), v19+int32(12), v62, v85, int32(1), v89, v110)
																							mBase = m.M
																							v119 = m.ExcPending
																							if v119 != 0 {
																								return int32(0)
																							} else {
																								v120 = F_SPI_finish(m)
																								mBase = m.M
																								v121 = m.ExcPending
																								if v121 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v82
																									*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v110
																									*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v76
																									m.G0 = v19 + int32(16)
																									return int32(0)
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v72 = F_pstrdup(m, int32(_a_F_connectby_text_serial_5))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	v74 = v72
																	v75 = int32(_a_F_connectby_text_serial_4)
																	v76 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																	v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																	*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
																	v82 = F_CreateTupleDescCopy(m, v81)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int32(0)
																	} else {
																		v85 = base.B2i32(v63 == int32(7))
																		F_validateConnectbyTupleDesc(m, v82, v85, int32(1))
																		mBase = m.M
																		v88 = m.ExcPending
																		if v88 != 0 {
																			return int32(0)
																		} else {
																			v89 = F_TupleDescGetAttInMetadata(m, v82)
																			mBase = m.M
																			v90 = m.ExcPending
																			if v90 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(2)
																				v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(1)
																				F_SPI_connect_ext(m, int32(0))
																				mBase = m.M
																				v98 = m.ExcPending
																				if v98 != 0 {
																					return int32(0)
																				} else {
																					v99 = int32(_a_F_connectby_text_serial_4)
																					v100 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																					*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																					v109 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[1]))
																					v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																					mBase = m.M
																					v111 = m.ExcPending
																					if v111 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v100
																						F_build_tuplestore_recursively(m, v31, v36, v26, v41, v74, v46, v46, int32(0), v19+int32(12), v62, v85, int32(1), v89, v110)
																						mBase = m.M
																						v119 = m.ExcPending
																						if v119 != 0 {
																							return int32(0)
																						} else {
																							v120 = F_SPI_finish(m)
																							mBase = m.M
																							v121 = m.ExcPending
																							if v121 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v82
																								*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v110
																								*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v76
																								m.G0 = v19 + int32(16)
																								return int32(0)
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
func F_contains_multiexpr_param(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(8) {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			return base.B2i32(v10 == int32(3))
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(1056), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_contsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.001))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_convert_case(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v81 int32
	_ = v81
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var __phi308 int32
	_ = __phi308
	var v311 int32
	_ = v311
	var __phi311 int32
	_ = __phi311
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v346 int32
	_ = v346
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v528 int32
	_ = v528
	var v538 int32
	_ = v538
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v577 int32
	_ = v577
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v921 int32
	_ = v921
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1017 int32
	_ = v1017
	var v1031 int32
	_ = v1031
	var v1043 int32
	_ = v1043
	var v1057 int32
	_ = v1057
	v9 = int32(0)
	if l4 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = v9
	goto L3
L3:
	;
	if l3 == int32(0) {
		v1043 = v9
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = v24
	goto L3
L6:
	;
	if base.Ui32(v1043) < base.Ui32(l1) {
		goto L268
	} else {
		goto L269
	}
L7:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = int32(1)
	goto L10
L9:
	;
	v33 = int32(2)
	goto L10
L10:
	;
	v43 = v9
	v47 = v9
	v48 = l4
	v50 = v28
	goto L11
L11:
	;
	v55 = l2 + v47
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 == int32(0) {
		v1043 = v43
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v1043 = v1017
	goto L6
L13:
	;
	v59 = int32(1)
	if int32(0) <= base.I32_extend8_s(v56) {
		v131 = v56
		v132 = v59
		v134 = v59
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if l4 != int32(1) {
		v141 = v48
		v142 = v50
		goto L31
	} else {
		goto L32
	}
L15:
	;
	if v56&int32(224) == int32(192) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v55))))
	v118 = v115&int32(63) | v112
	if base.Ui32(v112) < base.Ui32(int32(128)) {
		v131 = v118
		v132 = v59
		v134 = int32(1)
		goto L14
	} else {
		goto L26
	}
L17:
	;
	v112 = v56 << (uint(int32(6)) % 32) & int32(1984)
	v113 = int32(1)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v56&int32(240) == int32(224) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v112 = v56<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v81&int32(63)<<(uint(int32(6))%32)
	v113 = int32(2)
	goto L16
L21:
	;
	goto L22
L22:
	;
	if v56&int32(248) != int32(240) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v131 = int32(-1)
	v132 = int32(0)
	v134 = int32(4)
	goto L14
L24:
	;
	goto L25
L25:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v100 = int32(63)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	v112 = v56<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v99&v100<<(uint(int32(12))%32) | v105&v100<<(uint(int32(6))%32)
	v113 = int32(3)
	goto L16
L26:
	;
	v122 = int32(0)
	if base.Ui32(v112) < base.Ui32(int32(2048)) {
		v131 = v118
		v132 = v122
		v134 = int32(2)
		goto L14
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v112) < base.Ui32(int32(_a_F_convert_case_2)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v130 = int32(3)
	goto L30
L29:
	;
	v130 = int32(4)
	goto L30
L30:
	;
	v131 = v118
	v132 = v122
	v134 = v130
	goto L14
L31:
	;
	if v132 != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v47 != v50 {
		v141 = int32(0)
		v142 = v50
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v139 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v141 = v33
	v142 = v139
	goto L31
L35:
	;
	v1031 = v47 + v134
	if base.B2i32(l3 < int32(0))|base.B2i32(base.Ui32(v1031) < base.Ui32(l3)) != 0 {
		v43 = v1017
		v47 = v1031
		v48 = v141
		v50 = v142
		goto L11
	} else {
		goto L267
	}
L36:
	;
	v906 = v43
	v907 = int32(0)
	goto L243
L37:
	;
	v812 = l0 + v43
	if base.Ui32(v787) <= base.Ui32(int32(2047)) {
		goto L237
	} else {
		goto L238
	}
L38:
	;
	v808 = v43 + int32(1)
	if base.Ui32(l1) < base.Ui32(v808) {
		goto L234
	} else {
		goto L235
	}
L39:
	;
	v801 = v43 + v134
	if base.Ui32(l1) < base.Ui32(v801) {
		goto L228
	} else {
		goto L229
	}
L40:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	if base.Ui32(v787) < base.Ui32(int32(128)) {
		goto L38
	} else {
		goto L220
	}
L41:
	;
	v143 = int32(2)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141<<(uint(v143)%32))+uint32(_c_F_convert_case[0])))
	v786 = v145 + v131<<(uint(v143)%32) + int32(4)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v151 = int32(0)
	if base.Ui32(v131) <= base.Ui32(int32(1415)) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v283 == int32(0) {
		goto L39
	} else {
		goto L101
	}
L45:
	;
	goto L44
L46:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[1]))))
	v283 = v156
	goto L45
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_3)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_4)) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_5)) {
		v283 = v151
		goto L45
	} else {
		goto L76
	}
L52:
	;
	if base.Ui32(v131-int32(_a_F_convert_case_6)) <= base.Ui32(int32(95)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_7)) {
		v283 = v151
		goto L45
	} else {
		goto L63
	}
L55:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[2]))))
	v283 = v169
	goto L45
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_8)) {
		v283 = v151
		goto L45
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_9)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[3]))))
	v283 = v178
	goto L45
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_10)) {
		v283 = v151
		goto L45
	} else {
		goto L62
	}
L62:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[4]))))
	v283 = v185
	goto L45
L63:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_11)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_12)) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_13)) {
		v283 = v151
		goto L45
	} else {
		goto L71
	}
L67:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[5]))))
	v283 = v196
	goto L45
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_14)) {
		v283 = v151
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[6]))))
	v283 = v203
	goto L45
L71:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_15)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[7]))))
	v283 = v212
	goto L45
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_16)) {
		v283 = v151
		goto L45
	} else {
		goto L75
	}
L75:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[8]))))
	v283 = v219
	goto L45
L76:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_17)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_18)) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_19)) {
		v283 = v151
		goto L45
	} else {
		goto L92
	}
L80:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_20)) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_21)) {
		v283 = v151
		goto L45
	} else {
		goto L87
	}
L83:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[9]))))
	v283 = v232
	goto L45
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_22)) {
		v283 = v151
		goto L45
	} else {
		goto L86
	}
L86:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[10]))))
	v283 = v239
	goto L45
L87:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_23)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[11]))))
	v283 = v248
	goto L45
L89:
	;
	goto L90
L90:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_24)) {
		v283 = v151
		goto L45
	} else {
		goto L91
	}
L91:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[12]))))
	v283 = v255
	goto L45
L92:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_25)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_26)) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(67)) < base.Ui32(v131-int32(_a_F_convert_case_27)) {
		v283 = v151
		goto L45
	} else {
		goto L100
	}
L96:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[13]))))
	v283 = v266
	goto L45
L97:
	;
	goto L98
L98:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_28)) {
		v283 = v151
		goto L45
	} else {
		goto L99
	}
L99:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[14]))))
	v283 = v273
	goto L45
L100:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131<<(uint(int32(1))%32))+uint32(_c_F_convert_case[15]))))
	v283 = v282
	goto L45
L101:
	;
	if l5 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v759 = int32(2)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v141<<(uint(v759)%32))+uint32(_c_F_convert_case[0])))
	v786 = v761 + v283<<(uint(v759)%32)
	goto L40
L103:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_convert_case[16]))))
	if v288 == int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v292 = v288 * int32(52)
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292)+uint32(_c_F_convert_case[17]))))
	switch v293 {
	case 0:
		goto L36
	case 1:
		goto L105
	default:
		goto L102
	}
L105:
	;
	if v47 == int32(0) {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v297 = v47 - int32(1)
	if v297 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if l3 == v47 {
		goto L36
	} else {
		goto L162
	}
L108:
	;
	__phi308 = v47
	__phi311 = v297
	v308 = __phi308
	v311 = __phi311
	goto L109
L109:
	;
	v321 = l2 + v311
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	v323 = base.I32_extend8_s(v322)
	if v323 < int32(-64) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L107
L111:
	;
	if int32(0) < v311 {
		__phi308 = v311
		__phi311 = v311 - int32(1)
		v308 = __phi308
		v311 = __phi311
		goto L109
	} else {
		goto L161
	}
L112:
	;
	if int32(0) <= v323 {
		v384 = v322
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v390 = Fn13969(m, v384, int32(4), int32(16), int32(_a_F_convert_case_29), int32(_a_F_convert_case_30), int32(505))
	mBase = m.M
	goto L123
L114:
	;
	if v322&int32(224) == int32(192) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377+v321))))
	v384 = v379&int32(63) | v376
	goto L113
L116:
	;
	v376 = v322 << (uint(int32(6)) % 32) & int32(1984)
	v377 = int32(1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	if v322&int32(240) == int32(224) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v308))))
	v376 = v322<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v346&int32(63)<<(uint(int32(6))%32)
	v377 = int32(2)
	goto L115
L120:
	;
	goto L121
L121:
	;
	if v322&int32(248) != int32(240) {
		v384 = int32(-1)
		goto L113
	} else {
		goto L122
	}
L122:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v308))))
	v364 = int32(63)
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+2)))
	v376 = v322<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v363&v364<<(uint(int32(12))%32) | v369&v364<<(uint(int32(6))%32)
	v377 = int32(3)
	goto L115
L123:
	;
	if v390 != 0 {
		goto L111
	} else {
		goto L124
	}
L124:
	;
	if base.Ui32(int32(127)) < base.Ui32(v384) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	if v498 != 0 {
		goto L107
	} else {
		goto L160
	}
L126:
	;
	v439 = int32(689)
	v440 = int32(0)
	goto L140
L127:
	;
	v399 = int32(3367)
	v400 = int32(0)
	goto L130
L128:
	;
	goto L129
L129:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384<<(uint(int32(1))%32))+uint32(_c_F_convert_case[18]))))
	v498 = int32(base.Ui32(v429&int32(8)) >> (uint(int32(3)) % 32))
	goto L125
L130:
	;
	v405 = base.I32_div_s(v399+v400, int32(2))
	v407 = v405 * int32(12)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v407)+uint32(_c_F_convert_case[19])))
	if base.Ui32(v410) < base.Ui32(v384) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+uint32(_c_F_convert_case[20]))))
	if v423 != int32(3) {
		goto L126
	} else {
		goto L139
	}
L132:
	;
	goto L131
L133:
	;
	if v421 <= v420 {
		v399 = v420
		v400 = v421
		goto L130
	} else {
		goto L138
	}
L134:
	;
	v420 = v399
	v421 = v405 + int32(1)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v407)+uint32(_c_F_convert_case[21])))
	if base.Ui32(v416) <= base.Ui32(v384) {
		goto L132
	} else {
		goto L137
	}
L137:
	;
	v420 = v405 - int32(1)
	v421 = v400
	goto L133
L138:
	;
	goto L126
L139:
	;
	v498 = int32(1)
	goto L125
L140:
	;
	v445 = base.I32_div_s(v439+v440, int32(2))
	v447 = v445 << (uint(int32(3)) % 32)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v447)+uint32(_c_F_convert_case[22])))
	if base.Ui32(v450) < base.Ui32(v384) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v467 = int32(655)
	v468 = int32(0)
	goto L150
L142:
	;
	if v462 <= v461 {
		v439 = v461
		v440 = v462
		goto L140
	} else {
		goto L149
	}
L143:
	;
	v461 = v439
	v462 = v445 + int32(1)
	goto L142
L144:
	;
	goto L145
L145:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v447)+uint32(_c_F_convert_case[23])))
	if base.Ui32(v456) <= base.Ui32(v384) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v498 = int32(1)
	goto L125
L147:
	;
	goto L148
L148:
	;
	v461 = v445 - int32(1)
	v462 = v440
	goto L142
L149:
	;
	goto L141
L150:
	;
	v473 = base.I32_div_s(v467+v468, int32(2))
	v475 = v473 << (uint(int32(3)) % 32)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475)+uint32(_c_F_convert_case[24])))
	if base.Ui32(v478) < base.Ui32(v384) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v498 = int32(0)
	goto L125
L152:
	;
	if v490 <= v489 {
		v467 = v489
		v468 = v490
		goto L150
	} else {
		goto L159
	}
L153:
	;
	v489 = v467
	v490 = v473 + int32(1)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v475)+uint32(_c_F_convert_case[25])))
	if base.Ui32(v484) <= base.Ui32(v384) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v498 = int32(1)
	goto L125
L157:
	;
	goto L158
L158:
	;
	v489 = v473 - int32(1)
	v490 = v468
	goto L152
L159:
	;
	goto L151
L160:
	;
	goto L102
L161:
	;
	goto L110
L162:
	;
	v528 = v47 + int32(1)
	if base.Ui32(l3) <= base.Ui32(v528) {
		goto L36
	} else {
		goto L163
	}
L163:
	;
	v538 = v528
	goto L164
L164:
	;
	v551 = l2 + v538
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v552 == int32(0) {
		goto L36
	} else {
		goto L166
	}
L165:
	;
	if base.Ui32(int32(127)) < base.Ui32(v614) {
		goto L186
	} else {
		goto L187
	}
L166:
	;
	v555 = base.I32_extend8_s(v552)
	if int32(-64) <= v555 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L165
L168:
	;
	if int32(0) <= v555 {
		v614 = v552
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	v626 = v538 + int32(1)
	if v626 != l3 {
		v538 = v626
		goto L164
	} else {
		goto L183
	}
L171:
	;
	v620 = Fn13969(m, v614, int32(4), int32(16), int32(_a_F_convert_case_29), int32(_a_F_convert_case_30), int32(505))
	mBase = m.M
	goto L181
L172:
	;
	if v552&int32(224) == int32(192) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607+v551))))
	v614 = v609&int32(63) | v606
	goto L171
L174:
	;
	v606 = v552 << (uint(int32(6)) % 32) & int32(1984)
	v607 = int32(1)
	goto L173
L175:
	;
	goto L176
L176:
	;
	if v552&int32(240) == int32(224) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+1)))
	v606 = v552<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v577&int32(63)<<(uint(int32(6))%32)
	v607 = int32(2)
	goto L173
L178:
	;
	goto L179
L179:
	;
	if v552&int32(248) != int32(240) {
		v614 = int32(-1)
		goto L171
	} else {
		goto L180
	}
L180:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+1)))
	v594 = int32(63)
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+2)))
	v606 = v552<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v593&v594<<(uint(int32(12))%32) | v599&v594<<(uint(int32(6))%32)
	v607 = int32(3)
	goto L173
L181:
	;
	if v620 == int32(0) {
		goto L167
	} else {
		goto L182
	}
L182:
	;
	goto L170
L183:
	;
	goto L36
L184:
	;
	if v735 == int32(0) {
		goto L36
	} else {
		goto L219
	}
L185:
	;
	v676 = int32(689)
	v677 = int32(0)
	goto L199
L186:
	;
	v636 = int32(3367)
	v637 = int32(0)
	goto L189
L187:
	;
	goto L188
L188:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614<<(uint(int32(1))%32))+uint32(_c_F_convert_case[18]))))
	v735 = int32(base.Ui32(v666&int32(8)) >> (uint(int32(3)) % 32))
	goto L184
L189:
	;
	v642 = base.I32_div_s(v636+v637, int32(2))
	v644 = v642 * int32(12)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v644)+uint32(_c_F_convert_case[19])))
	if base.Ui32(v647) < base.Ui32(v614) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+uint32(_c_F_convert_case[20]))))
	if v660 != int32(3) {
		goto L185
	} else {
		goto L198
	}
L191:
	;
	goto L190
L192:
	;
	if v658 <= v657 {
		v636 = v657
		v637 = v658
		goto L189
	} else {
		goto L197
	}
L193:
	;
	v657 = v636
	v658 = v642 + int32(1)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v644)+uint32(_c_F_convert_case[21])))
	if base.Ui32(v653) <= base.Ui32(v614) {
		goto L191
	} else {
		goto L196
	}
L196:
	;
	v657 = v642 - int32(1)
	v658 = v637
	goto L192
L197:
	;
	goto L185
L198:
	;
	v735 = int32(1)
	goto L184
L199:
	;
	v682 = base.I32_div_s(v676+v677, int32(2))
	v684 = v682 << (uint(int32(3)) % 32)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v684)+uint32(_c_F_convert_case[22])))
	if base.Ui32(v687) < base.Ui32(v614) {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v704 = int32(655)
	v705 = int32(0)
	goto L209
L201:
	;
	if v699 <= v698 {
		v676 = v698
		v677 = v699
		goto L199
	} else {
		goto L208
	}
L202:
	;
	v698 = v676
	v699 = v682 + int32(1)
	goto L201
L203:
	;
	goto L204
L204:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v684)+uint32(_c_F_convert_case[23])))
	if base.Ui32(v693) <= base.Ui32(v614) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v735 = int32(1)
	goto L184
L206:
	;
	goto L207
L207:
	;
	v698 = v682 - int32(1)
	v699 = v677
	goto L201
L208:
	;
	goto L200
L209:
	;
	v710 = base.I32_div_s(v704+v705, int32(2))
	v712 = v710 << (uint(int32(3)) % 32)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712)+uint32(_c_F_convert_case[24])))
	if base.Ui32(v715) < base.Ui32(v614) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	v735 = int32(0)
	goto L184
L211:
	;
	if v727 <= v726 {
		v704 = v726
		v705 = v727
		goto L209
	} else {
		goto L218
	}
L212:
	;
	v726 = v704
	v727 = v710 + int32(1)
	goto L211
L213:
	;
	goto L214
L214:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v712)+uint32(_c_F_convert_case[25])))
	if base.Ui32(v721) <= base.Ui32(v614) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v735 = int32(1)
	goto L184
L216:
	;
	goto L217
L217:
	;
	v726 = v710 - int32(1)
	v727 = v705
	goto L211
L218:
	;
	goto L210
L219:
	;
	goto L102
L220:
	;
	if base.Ui32(v787) < base.Ui32(int32(_a_F_convert_case_2)) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v795 = int32(3)
	goto L223
L222:
	;
	v795 = int32(4)
	goto L223
L223:
	;
	if base.Ui32(v787) < base.Ui32(int32(2048)) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v798 = int32(2)
	goto L226
L225:
	;
	v798 = v795
	goto L226
L226:
	;
	v799 = v798 + v43
	if base.Ui32(v799) <= base.Ui32(l1) {
		goto L37
	} else {
		goto L227
	}
L227:
	;
	v1017 = v799
	goto L35
L228:
	;
	v1017 = v801
	goto L35
L229:
	;
	goto L230
L230:
	;
	if v134 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1017 = v801
	goto L35
L232:
	;
	goto L233
L233:
	;
	base.MemoryCopy(m, l0+v43, v55, v134)
	v1017 = v801
	goto L35
L234:
	;
	v1017 = v808
	goto L35
L235:
	;
	goto L236
L236:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v43))) = uint8(v787)
	v1017 = v808
	goto L35
L237:
	;
	v818 = v787&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v812)+1)) = uint8(v818)
	v823 = int32(base.Ui32(v787)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v812))) = uint8(v823)
	v1017 = v799
	goto L35
L238:
	;
	goto L239
L239:
	;
	if base.Ui32(v787) <= base.Ui32(int32(_a_F_convert_case_31)) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v827 = int32(63)
	v829 = int32(128)
	v830 = v787&v827 | v829
	*(*uint8)(unsafe.Add(mBase, uint32(v812)+2)) = uint8(v830)
	v835 = int32(base.Ui32(v787)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v812))) = uint8(v835)
	v842 = int32(base.Ui32(v787)>>(uint(int32(6))%32))&v827 | v829
	*(*uint8)(unsafe.Add(mBase, uint32(v812)+1)) = uint8(v842)
	v1017 = v799
	goto L35
L241:
	;
	goto L242
L242:
	;
	v844 = int32(63)
	v846 = int32(128)
	v847 = v787&v844 | v846
	*(*uint8)(unsafe.Add(mBase, uint32(v812)+3)) = uint8(v847)
	v854 = int32(base.Ui32(v787)>>(uint(int32(6))%32))&v844 | v846
	*(*uint8)(unsafe.Add(mBase, uint32(v812)+2)) = uint8(v854)
	v861 = int32(base.Ui32(v787)>>(uint(int32(12))%32))&v844 | v846
	*(*uint8)(unsafe.Add(mBase, uint32(v812)+1)) = uint8(v861)
	v868 = int32(base.Ui32(v787)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v812))) = uint8(v868)
	v1017 = v799
	goto L35
L243:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v141*int32(12)+v292+int32(_a_F_convert_case_32)+v907<<(uint(int32(2))%32))))
	if v921 == int32(0) {
		v1017 = v906
		goto L35
	} else {
		goto L245
	}
L244:
	;
	v1017 = v1003
	goto L35
L245:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v921) {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v1005 = v907 + int32(1)
	if v1005 != int32(3) {
		v906 = v1003
		v907 = v1005
		goto L243
	} else {
		goto L266
	}
L247:
	;
	v1003 = v935
	goto L246
L248:
	;
	if base.Ui32(v921) <= base.Ui32(int32(_a_F_convert_case_31)) {
		goto L263
	} else {
		goto L264
	}
L249:
	;
	if base.Ui32(v921) < base.Ui32(int32(_a_F_convert_case_2)) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	v951 = v906 + int32(1)
	if base.Ui32(l1) < base.Ui32(v951) {
		goto L260
	} else {
		goto L261
	}
L252:
	;
	v931 = int32(3)
	goto L254
L253:
	;
	v931 = int32(4)
	goto L254
L254:
	;
	if base.Ui32(v921) < base.Ui32(int32(2048)) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v934 = int32(2)
	goto L257
L256:
	;
	v934 = v931
	goto L257
L257:
	;
	v935 = v934 + v906
	if base.Ui32(l1) < base.Ui32(v935) {
		goto L247
	} else {
		goto L258
	}
L258:
	;
	v937 = l0 + v906
	if base.Ui32(int32(2047)) < base.Ui32(v921) {
		goto L248
	} else {
		goto L259
	}
L259:
	;
	v943 = v921&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)) = uint8(v943)
	v948 = int32(base.Ui32(v921)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v948)
	goto L247
L260:
	;
	v1003 = v951
	goto L246
L261:
	;
	goto L262
L262:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v906))) = uint8(v921)
	v1003 = v951
	goto L246
L263:
	;
	v957 = int32(63)
	v959 = int32(128)
	v960 = v921&v957 | v959
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+2)) = uint8(v960)
	v965 = int32(base.Ui32(v921)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v965)
	v972 = int32(base.Ui32(v921)>>(uint(int32(6))%32))&v957 | v959
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)) = uint8(v972)
	goto L247
L264:
	;
	goto L265
L265:
	;
	v974 = int32(63)
	v976 = int32(128)
	v977 = v921&v974 | v976
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+3)) = uint8(v977)
	v984 = int32(base.Ui32(v921)>>(uint(int32(6))%32))&v974 | v976
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+2)) = uint8(v984)
	v991 = int32(base.Ui32(v921)>>(uint(int32(12))%32))&v974 | v976
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)) = uint8(v991)
	v998 = int32(base.Ui32(v921)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v998)
	goto L247
L266:
	;
	goto L244
L267:
	;
	goto L12
L268:
	;
	v1057 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1043))) = uint8(v1057)
	goto L270
L269:
	;
	goto L270
L270:
	;
	return v1043
}
func F_copy_dest_receive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
	F_MemoryContextReset(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a_F_copy_dest_receive_0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
		*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		if v20 < v19 {
			F_slot_getsomeattrs_int(m, l0, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				m.T0[v25].(func(*base.Module, int32, int32))(m, v7, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v14
					v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
					v32 = v30 + int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v32
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[1]))
					if v37 == int32(0) {
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copy_dest_receive[2])))
						if v41&int32(1) == int32(0) {
						} else {
							v46 = int32(_a_F_copy_dest_receive_1)
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
							v49 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v48 + v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v52 + v49
							*(*int64)(unsafe.Add(mBase, uint32(v37+int32(16))+232)) = v32
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v60 + v49
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v66 - v49
						}
					}
					return int32(1)
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			m.T0[v25].(func(*base.Module, int32, int32))(m, v7, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v14
				v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v32 = v30 + int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v32
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[1]))
				if v37 == int32(0) {
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copy_dest_receive[2])))
					if v41&int32(1) == int32(0) {
					} else {
						v46 = int32(_a_F_copy_dest_receive_1)
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
						v49 = int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v48 + v49
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v52 + v49
						*(*int64)(unsafe.Add(mBase, uint32(v37+int32(16))+232)) = v32
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v60 + v49
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v66 - v49
					}
				}
				return int32(1)
			}
		}
	}
}
func F_copy_intArrayType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = F_ArrayGetNItemsSafe(m, v6, l0+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if int32(0) < v9 {
			v18 = v9<<(uint(int32(2))%32) + int32(24)
			v19 = F_palloc0(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(23)
				*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v18 << (uint(int32(2)) % 32)
				v35 = v19
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v44 = v35
				v45 = (v37<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 == int32(0) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				} else {
					v56 = v46
				}
				v58 = v9 << (uint(int32(2)) % 32)
				if v58 != 0 {
					base.MemoryCopy(m, v44+v45, l0+v56, v58)
				} else {
				}
				return v44
			}
		} else {
			v32 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
				if v34 != 0 {
					v44 = v32
					v45 = v34
				} else {
					v35 = v32
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					v44 = v35
					v45 = (v37<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 == int32(0) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				} else {
					v56 = v46
				}
				v58 = v9 << (uint(int32(2)) % 32)
				if v58 != 0 {
					base.MemoryCopy(m, v44+v45, l0+v56, v58)
				} else {
				}
				return v44
			}
		}
	}
}
func F_copy_pathtarget(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v6 = F_palloc0(m, int32(40))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(277)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v20
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = F_list_copy(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v23
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v26 == int32(0) {
				return v6
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v29 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v34 = v30 << (uint(int32(2)) % 32)
				} else {
					v34 = int32(0)
				}
				v35 = F_palloc(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v35
					if v34 == int32(0) {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						base.MemoryCopy(m, v35, v40, v34)
					}
					return v6
				}
			}
		}
	}
}
func F_copysignl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l2&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(l2&int64(9223090561878065152))>>(uint(v10)%64)))|base.I32_wrap_i64(int64(base.Ui64(l3)>>(uint(v10)%64)))&int32(_a_F_copysignl_0))<<(uint(v10)%64)
	return
}
func F_cost_resultscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	v7 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v88 = *(*float64)(unsafe.Add(mBase, _c_F_cost_resultscan[0]))
	v89 = float64(0)
	v90 = base.F64_add(v85, v89)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v90, base.F64_add(base.F64_mul(v86, base.F64_add(v82, v88)), v89))
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v18
	if v17 == int32(0) {
		v63 = v7
		v66 = float64(0)
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v71 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v71
	v73 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v82 = v73
	v85 = v74
	goto L1
L5:
	;
	v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v82 = base.F64_add(v63, v67)
	v85 = base.F64_add(v66, v69)
	goto L1
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= int32(0) {
		v63 = v7
		v66 = float64(0)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v34 = int32(0)
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v34<<(uint(int32(2))%32))))
	v48 = F_cost_qual_eval_walker(m, v45, v13+int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v63 = v54
	v66 = v55
	goto L5
L10:
	;
	return
L11:
	;
	v51 = v34 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v51 < v52 {
		v34 = v51
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
func F_cost_seqscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v82 float64
	_ = v82
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v106 float64
	_ = v106
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v118 float64
	_ = v118
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v133 float64
	_ = v133
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v151 float64
	_ = v151
	var v155 float64
	_ = v155
	var v158 float64
	_ = v158
	var v163 int32
	_ = v163
	var v166 float64
	_ = v166
	v8 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = l3 + int32(8)
	goto L3
L2:
	;
	v23 = l2 + int32(16)
	goto L3
L3:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	F_get_tablespace_page_costs(m, v26, int32(0), v17)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v111)+24))
	v113 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v116 = *(*float64)(unsafe.Add(mBase, _c_F_cost_seqscan[0]))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v120 = base.F64_add(base.F64_mul(v112, v113), base.F64_mul(base.F64_add(v110, v116), v118))
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v111)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) < v122 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v33
	if v32 == int32(0) {
		v82 = v8
		v89 = float64(0)
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v106 = v94
	v110 = v95
	goto L6
L10:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v106 = base.F64_add(v89, v90)
	v110 = base.F64_add(v82, v92)
	goto L6
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v42 <= int32(0) {
		v82 = v8
		v89 = float64(0)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(0)
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v49<<(uint(int32(2))%32))))
	v67 = F_cost_qual_eval_walker(m, v64, v17+int32(8))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v82 = v73
	v89 = v74
	goto L10
L15:
	;
	v70 = v49 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v70 < v71 {
		v49 = v70
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v125 = base.F64_convert_i32_u(v122)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_seqscan[1])))
	if v127 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v158 = v120
	goto L19
L19:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_seqscan[2])))
	v166 = base.F64_add(base.F64_add(v106, float64(0)), v121)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v163 ^ int32(1)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(v31, base.F64_convert_i32_u(v30)), base.F64_add(v166, v158))
	m.G0 = v17 + int32(32)
	return
L20:
	;
	v133 = base.F64_add(base.F64_mul(v125, float64(-0.3)), float64(1))
	if base.F64_gt(v133, float64(0)) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v139 = v125
	goto L22
L22:
	;
	v141 = float64(1e+100)
	v142 = base.F64_div(v113, v139)
	if base.F64_gt(v142, v141)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v142)&int64(9223372036854775807))) != 0 {
		v155 = v141
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v137 = v133
	goto L25
L24:
	;
	v137 = math.Float64frombits(uint64(0x8000000000000000))
	goto L25
L25:
	;
	v139 = base.F64_add(v137, v125)
	goto L22
L26:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v155
	v158 = base.F64_div(v120, v139)
	goto L19
L27:
	;
	v151 = float64(1)
	if base.F64_le(v142, v151) != 0 {
		v155 = v151
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v155 = base.F64_nearest(v142)
	goto L26
}
func F_count_nulls(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v234 int32
	_ = v234
	v4 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == v4 {
		v27 = v4
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		if v19 == int32(0) {
			v27 = v4
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v22 != int32(15) {
				v27 = v4
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+13)))
				v27 = v25
			}
		}
	}
	if v27&int32(1) == int32(0) {
		v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v32 <= int32(0) {
			v207 = v32
			v209 = v4
		} else {
			if base.Ui32(int32(4)) <= base.Ui32(v32) {
				v42 = v4
				v45 = v4
				v48 = v4
				for {
					v55 = l0 + v42<<(uint(int32(3))%32)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+32)))
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+40)))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+48)))
					v63 = v45 + v56 + v58 + v60 + v62
					v64 = int32(4)
					v65 = v42 + v64
					v67 = v48 + v64
					if v67 != v32&int32(_a_F_count_nulls_0) {
						v42 = v65
						v45 = v63
						v48 = v67
						continue
					} else {
						break
					}
					break
				}
				if v32&int32(3) == int32(0) {
					v207 = v32
					v209 = v63
				} else {
					v76 = v65
					v79 = v63
					v92 = v76
					v95 = v79
					v97 = v4
					for {
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92<<(uint(int32(3))%32))+24)))
						v107 = v95 + v106
						v108 = int32(1)
						v111 = v97 + v108
						if v111 != v32&int32(3) {
							v92 = v92 + v108
							v95 = v107
							v97 = v111
							continue
						} else {
							break
						}
						break
					}
					v207 = v32
					v209 = v107
				}
			} else {
				v76 = v4
				v79 = v4
				v92 = v76
				v95 = v79
				v97 = v4
				for {
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92<<(uint(int32(3))%32))+24)))
					v107 = v95 + v106
					v108 = int32(1)
					v111 = v97 + v108
					if v111 != v32&int32(3) {
						v92 = v92 + v108
						v95 = v107
						v97 = v111
						continue
					} else {
						break
					}
					break
				}
				v207 = v32
				v209 = v107
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v207
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v209
		v234 = int32(1)
		return v234
	} else {
		v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v114 != 0 {
			v234 = int32(0)
			return v234
		} else {
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v116 = F_pg_detoast_datum(m, v115)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
				v122 = v116 + int32(16)
				v123 = F_ArrayGetNItemsSafe(m, v120, v122)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
					if v125 == int32(0) {
						v207 = v123
						v209 = v4
					} else {
						v128 = int32(1)
						if v123 <= int32(0) {
							v207 = v123
							v209 = v4
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
							v134 = v122 + v131<<(uint(int32(3))%32)
							if v123 != int32(1) {
								v141 = v134
								v144 = v128
								v147 = v4
								v149 = v4
								for {
									v155 = int32(1)
									v158 = v144 << (uint(v155) % 32)
									v160 = base.B2i32(v158 == int32(256))
									if v158 == int32(256) {
										v161 = v155
									} else {
										v161 = v158
									}
									v163 = v161 << (uint(int32(1)) % 32)
									v165 = base.B2i32(v163 == int32(256))
									if v163 == int32(256) {
										v166 = v155
									} else {
										v166 = v163
									}
									v167 = v160 + v141
									v168 = v165 + v167
									v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
									v171 = int32(0)
									v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
									v178 = base.B2i32(v161&v169 == v171) + (v147 + base.B2i32(v144&v173 == v171))
									v180 = v149 + int32(2)
									if v180 != v123&int32(2147483646) {
										v141 = v168
										v144 = v166
										v147 = v178
										v149 = v180
										continue
									} else {
										break
									}
									break
								}
								if v123&int32(1) == int32(0) {
									v207 = v123
									v209 = v178
								} else {
									v184 = v168
									v187 = v166
									v190 = v178
									v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
									v207 = v123
									v209 = v190 + base.B2i32(v187&v198 == int32(0))
								}
							} else {
								v184 = v134
								v187 = v128
								v190 = v4
								v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
								v207 = v123
								v209 = v190 + base.B2i32(v187&v198 == int32(0))
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v207
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v209
					v234 = int32(1)
					return v234
				}
			}
		}
	}
}
func F_create_drop_transactional_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v4 = l3
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
	v14 = F_MemoryContextAlloc(m, v12, int32(28))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v10 {
				v39 = v17
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = v39 + int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v47 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v46
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v55
				v58 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v58
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v61 + int32(1)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
				v23 = F_MemoryContextAlloc(m, v21, int32(24))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v10
					v29 = v23 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v29
					v32 = int32(_a_F_create_drop_transactional_internal_0)
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
					*(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2])) = v23
					v39 = v23
					*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
					v46 = v39 + int32(8)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					if v47 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v46
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v55
					v58 = v14 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v58
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v61 + int32(1)
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
			v23 = F_MemoryContextAlloc(m, v21, int32(24))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v10
				v29 = v23 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v29
				v32 = int32(_a_F_create_drop_transactional_internal_0)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
				*(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2])) = v23
				v39 = v23
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = v39 + int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v47 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v46
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v55
				v58 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v58
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v61 + int32(1)
				return
			}
		}
	}
}
func F_create_plan(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = int64(0)
	v6 = F_create_plan_recurse(m, l0, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v10 != int32(333) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v22 = int32(0)
	goto L7
L4:
	;
	goto L5
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	goto L5
L7:
	;
	v23 = int32(0)
	if v13 == v23 {
		v33 = v23
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v14 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= v22 {
		v33 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = v29 + v22<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+v22<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v52
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)) = uint16(v54)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+26)) = uint8(v56)
	v22 = v22 + int32(1)
	goto L7
L13:
	;
	goto L6
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.B2i32(v33 == int32(0))|base.B2i32(v38 <= v22) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v41 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	return v6
L20:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_0), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_create_plan_1), int32(372), int32(_a_F_create_plan_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_syncrep_config(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v3 = l2
	v4 = int32(0)
	v8 = int32(16)
	if l1 == v4 {
		v41 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = F_palloc(m, v41)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= int32(0) {
		v41 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v11
	goto L6
L5:
	;
	v17 = v14
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = v8
	v24 = v4
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v24<<(uint(int32(2))%32))))
	v30 = F_strlen(m, v29)
	mBase = m.M
	v32 = int32(1)
	v33 = v30 + v23 + v32
	v35 = v24 + v32
	if v35 != v17 {
		v23 = v33
		v24 = v35
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v41 = v33
	goto L1
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v41
	v52 = l0
	goto L13
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v96
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v57 = v52 + int32(1)
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52))))
	v59 = F___isspace(m, v58)
	mBase = m.M
	if v59 != 0 {
		v52 = v57
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v60 = int32(1)
	switch v58&int32(255) - int32(43) {
	case 0:
		v66 = v60
		goto L17
	default:
		v68 = v58
		v69 = v52
		v70 = v60
		goto L16
	case 2:
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v71 = int32(0)
	v73 = v68 - int32(48)
	if base.Ui32(v73) <= base.Ui32(int32(9)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
	v68 = v67
	v69 = v57
	v70 = v66
	goto L16
L18:
	;
	v66 = int32(0)
	goto L17
L19:
	;
	v76 = v71
	v77 = v73
	v78 = v69
	goto L22
L20:
	;
	v90 = v71
	goto L21
L21:
	;
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v80 = int32(10)
	v82 = v76*v80 - v77
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+1)))
	v87 = v83 - int32(48)
	if base.Ui32(v87) < base.Ui32(v80) {
		v76 = v82
		v77 = v87
		v78 = v78 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v90 = v82
	goto L21
L24:
	;
	goto L23
L25:
	;
	v96 = int32(0) - v90
	goto L27
L26:
	;
	v96 = v90
	goto L27
L27:
	;
	goto L12
L28:
	;
	return v44
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v101 <= int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(0)
	goto L28
L32:
	;
	v111 = v44 + int32(16)
	v112 = int32(0)
	goto L33
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v112<<(uint(int32(2))%32))))
	if (v118^v111)&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	goto L28
L35:
	;
	v193 = F_strlen(m, v118)
	mBase = m.M
	v195 = int32(1)
	v198 = v112 + v195
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v198 < v199 {
		v111 = v111 + v193 + v195
		v112 = v198
		goto L33
	} else {
		goto L56
	}
L36:
	;
	goto L35
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v172)
	if v172&int32(255) == int32(0) {
		goto L36
	} else {
		goto L52
	}
L38:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v171 = v118
	v172 = v124
	v173 = v111
	goto L37
L39:
	;
	goto L40
L40:
	;
	if v118&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v128 = v118
	v130 = v111
	goto L44
L42:
	;
	v142 = v118
	v144 = v111
	goto L43
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v149 = int32(-2139062144)
	if (int32(16843008)-v146|v146)&v149 != v149 {
		v171 = v142
		v172 = v146
		v173 = v144
		goto L37
	} else {
		goto L48
	}
L44:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v131)
	if v131 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L45:
	;
	v142 = v138
	v144 = v136
	goto L43
L46:
	;
	v135 = int32(1)
	v136 = v130 + v135
	v138 = v128 + v135
	if v138&int32(3) != 0 {
		v128 = v138
		v130 = v136
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v154 = v142
	v155 = v146
	v156 = v144
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155
	v158 = int32(4)
	v159 = v156 + v158
	v161 = v154 + v158
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v166 = int32(-2139062144)
	if (int32(16843008)-v163|v163)&v166 == v166 {
		v154 = v161
		v155 = v163
		v156 = v159
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v171 = v161
	v172 = v163
	v173 = v159
	goto L37
L51:
	;
	goto L50
L52:
	;
	v180 = v171
	v182 = v173
	goto L53
L53:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)) = uint8(v183)
	v185 = int32(1)
	if v183 != 0 {
		v180 = v180 + v185
		v182 = v182 + v185
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L36
L55:
	;
	goto L54
L56:
	;
	goto L34
}
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v3 = l2
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v12
		v88 = v11
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
		if v94 != 0 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
			*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			if v103 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			if v110 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v114 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
			v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
			if v121 < int32(0) {
			} else {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
				v126 = v124 - int32(97)
				if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
				} else {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					if v136 != 0 {
					} else {
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
						v142 = v139 + v121*int32(24)
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
						if v143 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
							v146 = v145
						} else {
							v146 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
						*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
					}
				}
			}
		}
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if base.Ui32(v16) <= base.Ui32(v15) {
				v34 = l0 + int32(76)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+136))
				if base.Ui32(v36) < base.Ui32(int32(98000000)) {
					v50 = int32(1024)
					v52 = v16 << (uint(int32(1)) % 32)
					if base.Ui32(v50) <= base.Ui32(v52) {
						v55 = v50
					} else {
						v55 = v52
					}
					v57 = v34
					v58 = v55
					v62 = v58*int32(40) + int32(8)
					v64 = F_palloc_extended(m, v62, int32(2))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						if v64 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(101)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
							if v72 != 0 {
								v74 = v72
							} else {
								v74 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v74
							v88 = int32(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+136))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+136)) = v77 + v62
							*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v58
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64))) = v81
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
							v88 = v64 + int32(8)
						}
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
						if v94 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
							*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
							*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							if v103 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							if v110 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							v114 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
							v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
							if v121 < int32(0) {
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
								v126 = v124 - int32(97)
								if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
									if v136 != 0 {
									} else {
										v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
										v142 = v139 + v121*int32(24)
										v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
										if v143 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
											v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
											v146 = v145
										} else {
											v146 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
										*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
									}
								}
							}
						}
						return
					}
				} else {
					v39 = v35
					v40 = v34
					*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(101)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
					if v45 != 0 {
						v47 = v45
					} else {
						v47 = int32(19)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v47
					v88 = int32(0)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
					if v94 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v103 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v110 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v114 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
						v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
						if v121 < int32(0) {
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
							v126 = v124 - int32(97)
							if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v136 != 0 {
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
									v142 = v139 + v121*int32(24)
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									if v143 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
										v146 = v145
									} else {
										v146 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
									*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
								}
							}
						}
					}
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
				v88 = v14 + v15*int32(40) + int32(8)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
				if v94 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v103 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v110 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v114 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
					v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
					if v121 < int32(0) {
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v126 = v124 - int32(97)
						if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
							if v136 != 0 {
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
								v142 = v139 + v121*int32(24)
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
								if v143 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									v146 = v145
								} else {
									v146 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
								*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
							}
						}
					}
				}
				return
			}
		} else {
			v27 = l0 + int32(76)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+136))
			if base.Ui32(int32(97999999)) < base.Ui32(v30) {
				v39 = v29
				v40 = v27
				*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(101)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
				if v45 != 0 {
					v47 = v45
				} else {
					v47 = int32(19)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v47
				v88 = int32(0)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
				if v94 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v103 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v110 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v114 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
					v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
					if v121 < int32(0) {
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v126 = v124 - int32(97)
						if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
							if v136 != 0 {
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
								v142 = v139 + v121*int32(24)
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
								if v143 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									v146 = v145
								} else {
									v146 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
								*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
							}
						}
					}
				}
				return
			} else {
				v57 = v27
				v58 = int32(64)
				v62 = v58*int32(40) + int32(8)
				v64 = F_palloc_extended(m, v62, int32(2))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v64 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(101)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
						if v72 != 0 {
							v74 = v72
						} else {
							v74 = int32(12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v74
						v88 = int32(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+136))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+136)) = v77 + v62
						*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v58
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = v81
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
						v88 = v64 + int32(8)
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
					if v94 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v103 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v110 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v114 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
						v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
						if v121 < int32(0) {
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
							v126 = v124 - int32(97)
							if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v136 != 0 {
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
									v142 = v139 + v121*int32(24)
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									if v143 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
										v146 = v145
									} else {
										v146 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
									*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
								}
							}
						}
					}
					return
				}
			}
		}
	}
}
func F_cstring_to_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v5 = F_strlen(m, l0)
	mBase = m.M
	v7 = v5 + int32(4)
	v8 = F_palloc(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v7 << (uint(int32(2)) % 32)
		if v5 != 0 {
			base.MemoryCopy(m, v8+int32(4), l0, v5)
		} else {
		}
		return v8
	}
}
