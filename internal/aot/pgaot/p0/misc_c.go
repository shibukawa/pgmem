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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+256)))
	if v9 != 0 {
		v10 = int32(6617)
	} else {
		v10 = int32(6618)
	}
	*(*int32)(unsafe.Add(mBase, _consts[1097])) = v10
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
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
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
	var v253 int32
	_ = v253
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
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	v3 = int32(0)
	v5 = l0 & int32(32639)
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
	return v317 & int32(65535)
L2:
	;
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314))))
	v317 = v315
	goto L1
L3:
	;
	if v5 != int32(8530) {
		v317 = v3
		goto L1
	} else {
		goto L96
	}
L4:
	;
	v314 = int32(2211000)
	goto L2
L5:
	;
	v314 = int32(2210996)
	goto L2
L6:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1075])))
	v317 = v308
	goto L1
L7:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1076])))
	v317 = v306
	goto L1
L8:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1077])))
	v317 = v304
	goto L1
L9:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1078])))
	v317 = v302
	goto L1
L10:
	;
	v180 = int32(47)
	v182 = int32(23)
	v183 = int32(0)
	goto L64
L11:
	;
	v49 = int32(24)
	v51 = int32(12)
	v52 = int32(0)
	goto L30
L12:
	;
	switch l1 - int32(246) {
	case 0:
		goto L13
	case 1:
		goto L14
	default:
		v317 = v3
		goto L1
	}
L13:
	;
	if v5 <= int32(17485) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	switch v5 - int32(8483) {
	case 0:
		v314 = int32(2210992)
		goto L2
	case 1:
		goto L5
	case 2, 3, 4, 5, 6:
		v317 = v3
		goto L1
	case 7:
		goto L4
	default:
		goto L3
	}
L15:
	;
	if v5 == int32(11357) {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v5 <= int32(20303) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v5 == int32(15742) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	if v5 != int32(17207) {
		v317 = v3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1079])))
	v317 = v24
	goto L1
L21:
	;
	if v5 == int32(17486) {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v5 == int32(20304) {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	if v5 != int32(19292) {
		v317 = v3
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1080])))
	v317 = v32
	goto L1
L26:
	;
	if v5 != int32(20554) {
		v317 = v3
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1081])))
	v317 = v38
	goto L1
L28:
	;
	v317 = v167 & int32(65535)
	goto L1
L29:
	;
	goto L28
L30:
	;
	v57 = v51 << (uint(int32(2)) % 32)
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+uint32(_consts[1082]))))
	v60 = base.B2i32(base.Ui32(v5) < base.Ui32(v59))
	if base.Ui32(v5) < base.Ui32(v59) {
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
	if base.Ui32(v5) < base.Ui32(v59) {
		goto L55
	} else {
		goto L56
	}
L34:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+uint32(_consts[1083]))))
	if base.Ui32(v61) <= base.Ui32(v5) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+uint32(_consts[1084]))))
	if v63 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v68 = v5 - v59&int32(65280)
	if base.Ui32(int32(41280)) <= base.Ui32(v5) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v71 = int32(255)
	v72 = v5 & v71
	v74 = v59 & v71
	v84 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v74))
	if base.Ui32(int32(160)) < base.Ui32(v74) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v111 = int32(255)
	v122 = v63 & v111
	if base.Ui32(int32(160)) < base.Ui32(v122) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v85 = int32(0)
	goto L42
L41:
	;
	v85 = int32(-34)
	goto L42
L42:
	;
	if base.Ui32(int32(160)) < base.Ui32(v74) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v88 = int32(34)
	goto L45
L44:
	;
	v88 = int32(0)
	goto L45
L45:
	;
	if base.Ui32(int32(160)) < base.Ui32(v72) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v91 = v85
	goto L48
L47:
	;
	v91 = v88
	goto L48
L48:
	;
	v96 = int32(33)
	v97 = v72 - v74 + v68>>(uint(int32(8))%32)*int32(157) + v91 + v63&int32(255) - v96
	v98 = int32(94)
	v99 = base.I32_div_s(v97, v98)
	v167 = v97 - v99*v98 + v63&int32(65280) + v99<<(uint(int32(8))%32) + v96
	goto L29
L49:
	;
	v128 = int32(65438)
	goto L51
L50:
	;
	v128 = int32(65472)
	goto L51
L51:
	;
	v129 = v5&v111 - v59&v111 + int32(base.Ui32(v68)>>(uint(int32(8))%32))*int32(94) + v122 + v128
	v131 = int32(157)
	v132 = base.I32_div_s(base.I32_extend16_s(v129), v131)
	v135 = v129 - v132*v131
	if int32(62) < base.I32_extend16_s(v135) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v147 = int32(98)
	goto L54
L53:
	;
	v147 = int32(64)
	goto L54
L54:
	;
	v167 = v135 + v63&int32(65280) + v132<<(uint(int32(8))%32) + v147
	goto L29
L55:
	;
	v151 = v52
	goto L57
L56:
	;
	v151 = v51 + int32(1)
	goto L57
L57:
	;
	if base.Ui32(v5) < base.Ui32(v59) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = v51 - int32(1)
	goto L60
L59:
	;
	v154 = v49
	goto L60
L60:
	;
	if v151 <= v154 {
		v49 = v154
		v51 = (v151 + v154) >> (uint(int32(1)) % 32)
		v52 = v151
		goto L30
	} else {
		goto L61
	}
L61:
	;
	goto L32
L62:
	;
	v317 = v298 & int32(65535)
	goto L1
L63:
	;
	goto L62
L64:
	;
	v188 = v182 << (uint(int32(2)) % 32)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_consts[1085]))))
	v191 = base.B2i32(base.Ui32(v5) < base.Ui32(v190))
	if base.Ui32(v5) < base.Ui32(v190) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v298 = int32(0)
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
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_consts[1086]))))
	if base.Ui32(v192) <= base.Ui32(v5) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_consts[1087]))))
	if v194 == int32(0) {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v199 = v5 - v190&int32(65280)
	if base.Ui32(int32(41280)) <= base.Ui32(v5) {
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
	v253 = v194 & v242
	if base.Ui32(int32(160)) < base.Ui32(v253) {
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
	v298 = v228 - v230*v229 + v194&int32(65280) + v230<<(uint(int32(8))%32) + v227
	goto L63
L83:
	;
	v259 = int32(65438)
	goto L85
L84:
	;
	v259 = int32(65472)
	goto L85
L85:
	;
	v260 = v5&v242 - v190&v242 + int32(base.Ui32(v199)>>(uint(int32(8))%32))*int32(94) + v253 + v259
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
	v298 = v266 + v194&int32(65280) + v263<<(uint(int32(8))%32) + v278
	goto L63
L89:
	;
	v282 = v183
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
		v183 = v282
		goto L64
	} else {
		goto L95
	}
L95:
	;
	goto L66
L96:
	;
	v314 = int32(2211004)
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
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
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
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
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
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v584 int32
	_ = v584
	v7 = int32(0)
	if l3 == v7 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v584
L2:
	;
	v584 = int32(0)
	goto L1
L3:
	;
	if v45&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v12&int32(2) == int32(0) {
		v45 = v12
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
		v584 = v19
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
		v45 = v20
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v584 = v19
	goto L1
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29&int32(72) == int32(8) {
		v45 = v29
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
		v45 = v34
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
	if v34&int32(64) != 0 {
		v584 = v37
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v34&int32(17) == int32(0) {
		v584 = v37
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v45 = v34
	goto L3
L20:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v310&int32(256) != 0 {
		goto L94
	} else {
		goto L95
	}
L21:
	;
	if (l0^l4)&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	goto L23
L23:
	;
	if l5 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v128 = l4 + l1 - int32(base.Ui32(v123)>>(uint(int32(10))%32))&int32(16383)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if (v129^v128)&int32(3) != 0 {
		goto L48
	} else {
		goto L49
	}
L25:
	;
	goto L24
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v101)
	if v101&int32(255) == int32(0) {
		goto L25
	} else {
		goto L41
	}
L27:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v100 = l0
	v101 = v53
	v102 = l4
	goto L26
L28:
	;
	goto L29
L29:
	;
	if l0&int32(3) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v57 = l0
	v59 = l4
	goto L33
L31:
	;
	v71 = l0
	v73 = l4
	goto L32
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 != v78 {
		v100 = v71
		v101 = v75
		v102 = v73
		goto L26
	} else {
		goto L37
	}
L33:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v60)
	if v60 == int32(0) {
		goto L25
	} else {
		goto L35
	}
L34:
	;
	v71 = v67
	v73 = v65
	goto L32
L35:
	;
	v64 = int32(1)
	v65 = v59 + v64
	v67 = v57 + v64
	if v67&int32(3) != 0 {
		v57 = v67
		v59 = v65
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v83 = v71
	v84 = v75
	v85 = v73
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v84
	v87 = int32(4)
	v88 = v85 + v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v91 = v83 + v87
	v95 = int32(-2139062144)
	if (v89|(int32(16843008)-v89))&v95 == v95 {
		v83 = v91
		v84 = v89
		v85 = v88
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v100 = v91
	v101 = v89
	v102 = v88
	goto L26
L40:
	;
	goto L39
L41:
	;
	v109 = v100
	v111 = v102
	goto L42
L42:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)) = uint8(v112)
	v114 = int32(1)
	if v112 != 0 {
		v109 = v109 + v114
		v111 = v111 + v114
		goto L42
	} else {
		goto L44
	}
L43:
	;
	goto L25
L44:
	;
	goto L43
L45:
	;
	if l5 == int32(0) {
		goto L20
	} else {
		goto L66
	}
L46:
	;
	goto L45
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v183)
	if v183&int32(255) == int32(0) {
		goto L46
	} else {
		goto L62
	}
L48:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v182 = v129
	v183 = v135
	v184 = v128
	goto L47
L49:
	;
	goto L50
L50:
	;
	if v129&int32(3) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v139 = v129
	v141 = v128
	goto L54
L52:
	;
	v153 = v129
	v155 = v128
	goto L53
L53:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v160 = int32(-2139062144)
	if (int32(16843008)-v157|v157)&v160 != v160 {
		v182 = v153
		v183 = v157
		v184 = v155
		goto L47
	} else {
		goto L58
	}
L54:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v142)
	if v142 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L55:
	;
	v153 = v149
	v155 = v147
	goto L53
L56:
	;
	v146 = int32(1)
	v147 = v141 + v146
	v149 = v139 + v146
	if v149&int32(3) != 0 {
		v139 = v149
		v141 = v147
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v165 = v153
	v166 = v157
	v167 = v155
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v166
	v169 = int32(4)
	v170 = v167 + v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v173 = v165 + v169
	v177 = int32(-2139062144)
	if (v171|(int32(16843008)-v171))&v177 == v177 {
		v165 = v173
		v166 = v171
		v167 = v170
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v182 = v173
	v183 = v171
	v184 = v170
	goto L47
L61:
	;
	goto L60
L62:
	;
	v191 = v182
	v193 = v184
	goto L63
L63:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)) = uint8(v194)
	v196 = int32(1)
	if v194 != 0 {
		v191 = v191 + v196
		v193 = v193 + v196
		goto L63
	} else {
		goto L65
	}
L64:
	;
	goto L46
L65:
	;
	goto L64
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l1 - int32(base.Ui32(v206)>>(uint(int32(10))%32))&int32(16383)
	goto L20
L67:
	;
	if (v225^l4)&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v225 = v215
	goto L67
L69:
	;
	goto L70
L70:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v217 = F_strlen(m, v216)
	mBase = m.M
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if base.Ui32(v217+v218) <= base.Ui32(int32(base.Ui32(v45)>>(uint(int32(10))%32))&int32(16383)) {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v225 = v216
	goto L67
L72:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v306 = F_strlen(m, l4)
	mBase = m.M
	v308 = F_strcpy(m, v306+l4, l0+int32(base.Ui32(v300)>>(uint(int32(10))%32))&int32(16383))
	mBase = m.M
	goto L93
L73:
	;
	goto L72
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v279)
	if v279&int32(255) == int32(0) {
		goto L73
	} else {
		goto L89
	}
L75:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v278 = v225
	v279 = v231
	v280 = l4
	goto L74
L76:
	;
	goto L77
L77:
	;
	if v225&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v235 = v225
	v237 = l4
	goto L81
L79:
	;
	v249 = v225
	v251 = l4
	goto L80
L80:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v256 = int32(-2139062144)
	if (int32(16843008)-v253|v253)&v256 != v256 {
		v278 = v249
		v279 = v253
		v280 = v251
		goto L74
	} else {
		goto L85
	}
L81:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v238)
	if v238 == int32(0) {
		goto L73
	} else {
		goto L83
	}
L82:
	;
	v249 = v245
	v251 = v243
	goto L80
L83:
	;
	v242 = int32(1)
	v243 = v237 + v242
	v245 = v235 + v242
	if v245&int32(3) != 0 {
		v235 = v245
		v237 = v243
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v261 = v249
	v262 = v253
	v263 = v251
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v262
	v265 = int32(4)
	v266 = v263 + v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v269 = v261 + v265
	v273 = int32(-2139062144)
	if (v267|(int32(16843008)-v267))&v273 == v273 {
		v261 = v269
		v262 = v267
		v263 = v266
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v278 = v269
	v279 = v267
	v280 = v266
	goto L74
L88:
	;
	goto L87
L89:
	;
	v287 = v278
	v289 = v280
	goto L90
L90:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+1)) = uint8(v290)
	v292 = int32(1)
	if v290 != 0 {
		v287 = v287 + v292
		v289 = v289 + v292
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L73
L92:
	;
	goto L91
L93:
	;
	goto L20
L94:
	;
	return l4
L95:
	;
	goto L96
L96:
	;
	if v310&int32(512) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v316 = int32(0)
	v317 = m.G0
	v318 = int32(16)
	v319 = v317 - v318
	m.G0 = v319
	v322 = l2 + v318
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v324 != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	goto L99
L99:
	;
	v552 = F_strlen(m, l4)
	mBase = m.M
	v557 = F_palloc(m, v552<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L106
	} else {
		goto L159
	}
L100:
	;
	if v530 == int32(0) {
		goto L2
	} else {
		goto L158
	}
L101:
	;
	v325 = l4
	v330 = v316
	goto L104
L102:
	;
	v347 = v316
	goto L103
L103:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v355 = int32(base.Ui32(v351)>>(uint(int32(1))%32)) & int32(65535)
	if v347 < v355 {
		v530 = v7
		goto L110
	} else {
		goto L111
	}
L104:
	;
	v335 = v330 + int32(1)
	v336 = F_pg_mblen_cstr(m, v325)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v347 = v335
	goto L103
L106:
	;
	return int32(0)
L107:
	;
	v340 = v336 + v325
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v341 != 0 {
		v325 = v340
		v330 = v335
		goto L104
	} else {
		goto L108
	}
L108:
	;
	goto L105
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L106
	} else {
		goto L155
	}
L110:
	;
	m.G0 = v319 + int32(16)
	goto L100
L111:
	;
	if v351&int32(1) == int32(0) {
		v384 = l4
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if v323 != 0 {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v361 = v347 - v355
	if v361 <= int32(0) {
		v384 = l4
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v364 = v361
	v368 = l4
	goto L115
L115:
	;
	v373 = F_pg_mblen_cstr(m, v368)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L106
	} else {
		goto L117
	}
L116:
	;
	v384 = v375
	goto L112
L117:
	;
	v375 = v373 + v368
	v376 = int32(1)
	if base.Ui32(v376) < base.Ui32(v364) {
		v364 = v364 - v376
		v368 = v375
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v391 = v323
	v393 = v384
	goto L122
L120:
	;
	goto L121
L121:
	;
	v530 = int32(1)
	goto L110
L122:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	switch v398&int32(3) - int32(1) {
	case 0:
		goto L126
	case 1:
		goto L125
	default:
		goto L109
	}
L123:
	;
	goto L121
L124:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	v509 = F_pg_mblen_cstr(m, v393)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L106
	} else {
		goto L153
	}
L125:
	;
	v451 = F_pg_mblen_cstr(m, v393)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L106
	} else {
		goto L140
	}
L126:
	;
	v403 = F_pg_mblen_cstr(m, v393)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L106
	} else {
		goto L127
	}
L127:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+8)))
	if v405 == int32(0) {
		v530 = v7
		goto L110
	} else {
		goto L128
	}
L128:
	;
	v415 = v391 + int32(8)
	goto L129
L129:
	;
	v419 = F_pg_mblen_cstr(m, v415)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L106
	} else {
		goto L131
	}
L130:
	;
	v530 = v7
	goto L110
L131:
	;
	if v403 == v419 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v422 = v403
	goto L135
L133:
	;
	goto L134
L134:
	;
	v449 = v415 + v419
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	if v450 != 0 {
		v415 = v449
		goto L129
	} else {
		goto L139
	}
L135:
	;
	if v422 == int32(0) {
		goto L124
	} else {
		goto L137
	}
L136:
	;
	goto L134
L137:
	;
	v434 = v422 - int32(1)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v434))))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434+v393))))
	if v436 == v438 {
		v422 = v434
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	goto L130
L140:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+8)))
	if v453 == int32(0) {
		goto L124
	} else {
		goto L141
	}
L141:
	;
	v463 = v391 + int32(8)
	goto L142
L142:
	;
	v467 = F_pg_mblen_cstr(m, v463)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L106
	} else {
		goto L144
	}
L143:
	;
	goto L124
L144:
	;
	if v451 == v467 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v470 = v451
	goto L148
L146:
	;
	goto L147
L147:
	;
	v497 = v463 + v467
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	if v498 != 0 {
		v463 = v497
		goto L142
	} else {
		goto L152
	}
L148:
	;
	if v470 == int32(0) {
		v530 = v7
		goto L110
	} else {
		goto L150
	}
L149:
	;
	goto L147
L150:
	;
	v482 = v470 - int32(1)
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v482))))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482+v393))))
	if v484 == v486 {
		v470 = v482
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L143
L153:
	;
	if v508 != 0 {
		v391 = v508
		v393 = v509 + v393
		goto L122
	} else {
		goto L154
	}
L154:
	;
	goto L123
L155:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v538 & int32(3)
	F_errmsg_internal(m, int32(480890), v319)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L106
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(489085), int32(245), int32(345306))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L106
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	v584 = l4
	goto L1
L159:
	;
	v559 = F_pg_mb2wchar_with_len(m, l4, v557, v552)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L106
	} else {
		goto L160
	}
L160:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v562 = int32(0)
	v565 = F_pg_regexec(m, v561, v557, v559, v562, v562, v562)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L106
	} else {
		goto L161
	}
L161:
	;
	F_pfree(m, v557)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L106
	} else {
		goto L162
	}
L162:
	;
	if v565 == int32(0) {
		v584 = l4
		goto L1
	} else {
		goto L163
	}
L163:
	;
	goto L2
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
	var v41 int32
	_ = v41
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
	var v110 int32
	_ = v110
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
	var v161 int32
	_ = v161
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v576 int32
	_ = v576
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v668 int32
	_ = v668
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v760 int32
	_ = v760
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
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
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L5
	} else {
		goto L218
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L5
	} else {
		goto L214
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L5
	} else {
		goto L211
	}
L4:
	;
	v692 = F_table_open(m, int32(2606), int32(2))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L94
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
	v41 = v4
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
	v63 = F_lappend_oid(m, v41, v62)
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
		v41 = v63
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
	v110 = v4
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_sequence_close(m, v77, int32(3))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L5
	} else {
		goto L93
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v110<<(uint(int32(2))%32))))
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
	v642 = v110 + int32(1)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v642 < v643 {
		v110 = v642
		goto L23
	} else {
		goto L92
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
		v161 = v130
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v168 = v161
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
		v161 = v154
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v161 = v154
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
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	if v383 != 0 {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	v206 = v202 & int32(3)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v209 = v207 - int32(2)
	v210 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v202) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v221 = v210
	v233 = int32(0)
	goto L56
L54:
	;
	v303 = v210
	goto L55
L55:
	;
	if v206 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v238 = int32(1)
	v239 = v221 << (uint(v238) % 32)
	v241 = v24 + int32(688)
	v244 = v24 + int32(768)
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v239))))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v246<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v239+v241))) = uint16(v250)
	v253 = v239 | int32(2)
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v253))))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v260<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v253+v241))) = uint16(v264)
	v266 = int32(4)
	v267 = v239 | v266
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v267))))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v274<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v267+v241))) = uint16(v278)
	v281 = v239 | int32(6)
	v288 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v281))))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v288<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v281+v241))) = uint16(v292)
	v295 = v221 + v266
	v297 = v233 + v266
	if v297 != v202&int32(2147483644) {
		v221 = v295
		v233 = v297
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v303 = v295
	goto L55
L58:
	;
	goto L57
L59:
	;
	v326 = v303
	v327 = v210
	goto L60
L60:
	;
	v343 = int32(1)
	v344 = v326 << (uint(v343) % 32)
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(768)+v344))))
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v351<<(uint(v343)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v344+(v24+int32(688))))) = uint16(v355)
	v360 = v327 + v343
	if v360 != v206 {
		v326 = v326 + v343
		v327 = v360
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L51
L62:
	;
	goto L61
L63:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v128)+96))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v128)+80))
	F_GetForeignKeyCheckTriggers(m, v77, v384, v385, v386, v24+int32(172), v24+int32(92))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v86 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v469 = F_palloc0(m, int32(108))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L5
	} else {
		goto L79
	}
L68:
	;
	v395 = int32(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v396 <= v395 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v405 = v395
	goto L70
L70:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v422+v405<<(uint(int32(2))%32))))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	v434 = F_tryAttachPartitionForeignKey(m, l0, v426, l2, v116, v427, v24+int32(688), v24+int32(624), v24+int32(432), v400, v399, v77)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L72
	}
L71:
	;
	F_sequence_close(m, v173, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L77
	}
L72:
	;
	if v434 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v439 = v405 + int32(1)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v439 < v440 {
		v405 = v439
		goto L70
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L71
L76:
	;
	goto L67
L77:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	goto L27
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v469))) = int64(438086664353)
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+12)) = uint8(v473)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+104)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+13)) = uint8(v475)
	v479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v469)+80)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v469)+72)) = v479
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+86)) = uint8(v483)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+87)) = uint8(v485)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+100)) = v479
	*(*int64)(unsafe.Add(mBase, uint32(v469)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+88)) = uint8(v487)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+15)) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+14)) = uint8(v493)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+16)) = uint8(v497)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v479 < v499 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v469)+76))
	v508 = int32(0)
	v510 = v502
	goto L83
L81:
	;
	v557 = v499
	goto L82
L82:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v128)+88))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+107)))
	F_addFkConstraint(m, v24+int32(96), int32(1), v128+int32(4), v469, l2, v173, v576, v116, v557, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v587, v24+int32(560), int32(0), v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L5
	} else {
		goto L88
	}
L83:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(688)+v508<<(uint(int32(1))%32)))))
	v541 = F_makeString(m, v525+v526<<(uint(int32(4))%32)+v535*int32(100)-int32(76))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L85
	}
L84:
	;
	v557 = v548
	goto L82
L85:
	;
	v543 = F_lappend(m, v510, v541)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+76)) = v543
	v547 = v508 + int32(1)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v547 < v548 {
		v508 = v547
		v510 = v543
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	F_addFkRecurseReferencing(m, l0, v469, l2, v173, v576, v594, v597, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v608, v24+int32(560), int32(0), int32(8), v613, v614, v591)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_sequence_close(m, v173, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	goto L27
L92:
	;
	goto L24
L93:
	;
	goto L4
L94:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v24+int32(768), int32(13), int32(3), int32(184), v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	F_ScanKeyInit(m, v24+int32(816), int32(4), int32(3), int32(61), int32(102))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v710 = int32(0)
	v717 = F_systable_beginscan(m, v692, v710, int32(1), v710, int32(2), v24+int32(768))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v719 = F_systable_getnext(m, v717)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	if v719 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v725 = v719
	v731 = v710
	goto L102
L100:
	;
	v760 = v710
	goto L101
L101:
	;
	F_systable_endscan(m, v717)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L5
	} else {
		goto L107
	}
L102:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v725)+16))
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742)+22)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v742+v743)))
	v746 = F_lappend_oid(m, v731, v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L104
	}
L103:
	;
	v760 = v746
	goto L101
L104:
	;
	v748 = F_systable_getnext(m, v717)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	if v748 != 0 {
		v725 = v748
		v731 = v746
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	F_sequence_close(m, v692, int32(2))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v778 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v783 = F_build_attrmap_by_name(m, v780, v781, int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	if v760 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_sequence_close(m, v778, int32(3))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L5
	} else {
		goto L210
	}
L112:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v787 <= int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v809 = int32(0)
	goto L114
L114:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v760)+12))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v813+v809<<(uint(int32(2))%32))))
	v818 = F_SearchSysCache1(m, int32(19), v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L118
	}
L115:
	;
	goto L111
L116:
	;
	F_ReleaseCatCache(m, v818)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L5
	} else {
		goto L208
	}
L117:
	;
	v1334 = int32(0)
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+8))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+107)))
	F_addFkConstraint(m, v24+int32(864), v1334, v1335, v1019, v865, l2, v1124, v817, v1336, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v1347, v24+int32(96), v1334, v1351)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L205
	}
L118:
	;
	if v818 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+22)))
	v822 = v820 + v821
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+92))
	v824 = int32(0)
	if v760 == v824 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L121
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L5
	} else {
		goto L202
	}
L122:
	;
	if v862 != 0 {
		goto L116
	} else {
		goto L135
	}
L123:
	;
	v862 = int32(0)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v830 <= int32(0) {
		v855 = v824
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v862 = v855
	goto L122
L127:
	;
	v833 = int32(0)
	if v833 < v830 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v836 = v830
	goto L130
L129:
	;
	v836 = v833
	goto L130
L130:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v760)+12))
	v839 = int32(0)
	goto L131
L131:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v837+v839<<(uint(int32(2))%32))))
	v848 = base.B2i32(v847 == v823)
	if v847 == v823 {
		v855 = v848
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v855 = v848
	goto L126
L133:
	;
	v850 = v839 + int32(1)
	if v850 != v836 {
		v839 = v850
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v822)+80))
	v865 = F_table_open(m, v863, int32(6))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v822)+88))
	F_DeconstructFkConstraintRow(m, v818, v24+int32(764), v24+int32(688), v24+int32(560), v24+int32(432), v24+int32(304), v24+int32(176), v24+int32(172), v24+int32(96))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v886 <= int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v1019 = F_palloc0(m, int32(108))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L5
	} else {
		goto L147
	}
L139:
	;
	v889 = int32(1)
	v891 = int32(0)
	if v886 != v889 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v901 = v891
	v904 = int32(0)
	goto L143
L141:
	;
	v961 = v891
	goto L142
L142:
	;
	if v886&v889 == int32(0) {
		goto L138
	} else {
		goto L146
	}
L143:
	;
	v918 = int32(1)
	v919 = v901 << (uint(v918) % 32)
	v921 = v24 + int32(624)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v925 = v24 + int32(560)
	v927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v925+v919))))
	v931 = int32(2)
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v923+v927<<(uint(v918)%32)-v931))))
	*(*uint16)(unsafe.Add(mBase, uint32(v919+v921))) = uint16(v933)
	v936 = v919 | v931
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v944 = int32(*(*int16)(unsafe.Add(mBase, uint32(v925+v936))))
	v950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v940+v944<<(uint(v918)%32)-v931))))
	*(*uint16)(unsafe.Add(mBase, uint32(v936+v921))) = uint16(v950)
	v953 = v901 + v931
	v955 = v904 + v931
	if v955 != v886&int32(2147483646) {
		v901 = v953
		v904 = v955
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v961 = v953
	goto L142
L145:
	;
	goto L144
L146:
	;
	v980 = int32(1)
	v981 = v961 << (uint(v980) % 32)
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v989 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(560)+v981))))
	v995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985+v989<<(uint(v980)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v981+(v24+int32(624))))) = uint16(v995)
	goto L138
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+8)) = v822 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v1019))) = int64(438086664353)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+12)) = uint8(v1026)
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+104)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+13)) = uint8(v1028)
	v1032 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+80)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+72)) = v1032
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+86)) = uint8(v1036)
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+87)) = uint8(v1038)
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+100)) = v1032
	*(*int64)(unsafe.Add(mBase, uint32(v1019)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+88)) = uint8(v1040)
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+15)) = uint8(v1032)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+14)) = uint8(v1046)
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+16)) = uint8(v1050)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v1032 < v1052 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+76))
	v1061 = int32(0)
	v1063 = v1055
	goto L151
L149:
	;
	goto L150
L150:
	;
	v1124 = F_index_get_partition(m, l2, v867)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L5
	} else {
		goto L156
	}
L151:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v865)+52))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	v1088 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(688)+v1061<<(uint(int32(1))%32)))))
	v1094 = F_makeString(m, v1078+v1079<<(uint(int32(4))%32)+v1088*int32(100)-int32(76))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	v1096 = F_lappend(m, v1063, v1094)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+76)) = v1096
	v1100 = v1061 + int32(1)
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v1100 < v1101 {
		v1061 = v1100
		v1063 = v1096
		goto L151
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	if v1124 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1126 = int32(0)
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+75)))
	if v1128 != int32(1) {
		v1315 = v1126
		v1322 = v1126
		goto L117
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L5
	} else {
		goto L199
	}
L160:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v822)+80))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v822)+96))
	F_ScanKeyInit(m, v24+int32(864), int32(11), int32(3), int32(184), v817)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	v1140 = int32(0)
	v1142 = int32(1)
	v1147 = F_systable_beginscan(m, v778, int32(2699), v1142, v1140, v1142, v24+int32(864))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L5
	} else {
		goto L164
	}
L162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L5
	} else {
		goto L196
	}
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L5
	} else {
		goto L193
	}
L164:
	;
	v1149 = F_systable_getnext(m, v1147)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	if v1149 == int32(0) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v1153 = v1140
	v1157 = v1149
	v1164 = v1126
	goto L167
L167:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+16))
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1174)+22)))
	v1176 = v1174 + v1175
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+84))
	if v1177 != v1131 {
		v1209 = v1153
		v1211 = v1164
		goto L171
	} else {
		goto L172
	}
L168:
	;
	goto L163
L169:
	;
	v1226 = F_systable_getnext(m, v1147)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L5
	} else {
		goto L191
	}
L170:
	;
	F_systable_endscan(m, v1147)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L190
	}
L171:
	;
	v1213 = F_systable_getnext(m, v1147)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L5
	} else {
		goto L186
	}
L172:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+4))
	if v1179 != v1132 {
		v1209 = v1153
		v1211 = v1164
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+76))
	v1184 = v1181 - int32(1644)
	if base.Ui32(v1184) <= base.Ui32(int32(11)) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v1192 != int32(1) {
		v1209 = v1153
		v1211 = v1164
		goto L171
	} else {
		goto L178
	}
L175:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1184<<(uint(int32(2))%32))+uint32(_consts[296])))
	v1192 = v1191
	goto L177
L176:
	;
	v1192 = int32(0)
	goto L177
L177:
	;
	goto L174
L178:
	;
	v1195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1176)+80)))
	if v1195&int32(8) != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v1204 == int32(0) {
		goto L169
	} else {
		goto L184
	}
L180:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	v1204 = v1198
	v1205 = v1153
	goto L179
L181:
	;
	goto L182
L182:
	;
	if v1195&int32(16) == int32(0) {
		v1204 = v1164
		v1205 = v1153
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	v1204 = v1164
	v1205 = v1203
	goto L179
L184:
	;
	if v1205 != 0 {
		v1220 = v1205
		v1221 = v1204
		goto L170
	} else {
		goto L185
	}
L185:
	;
	v1209 = int32(0)
	v1211 = v1204
	goto L171
L186:
	;
	if v1213 != 0 {
		v1153 = v1209
		v1157 = v1213
		v1164 = v1211
		goto L167
	} else {
		goto L187
	}
L187:
	;
	if v1211 == int32(0) {
		goto L163
	} else {
		goto L188
	}
L188:
	;
	if v1209 == int32(0) {
		goto L162
	} else {
		goto L189
	}
L189:
	;
	v1220 = v1209
	v1221 = v1211
	goto L170
L190:
	;
	v1315 = v1220
	v1322 = v1221
	goto L117
L191:
	;
	if v1226 != 0 {
		v1153 = v1205
		v1157 = v1226
		v1164 = int32(0)
		goto L167
	} else {
		goto L192
	}
L192:
	;
	goto L168
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v817
	F_errmsg_internal(m, int32(40559), v24+int32(32))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(489510), int32(12118), int32(133495))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v817
	F_errmsg_internal(m, int32(40628), v24+int32(48))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(489510), int32(12121), int32(133495))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
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
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v1283 + int32(4)
	F_errmsg_internal(m, int32(181996), v24+int32(16))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(489510), int32(11387), int32(458638))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v817
	F_errmsg_internal(m, int32(40796), v24)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(489510), int32(11319), int32(458638))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v24)+868))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+107)))
	F_addFkRecurseReferenced(m, v1019, v865, l2, v1124, v1354, v1355, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v1366, v24+int32(96), v1322, v1315, v1369)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	F_sequence_close(m, v865, int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	goto L116
L208:
	;
	v1399 = v809 + int32(1)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v1399 < v1400 {
		v809 = v1399
		goto L114
	} else {
		goto L209
	}
L209:
	;
	goto L115
L210:
	;
	m.G0 = v24 + int32(912)
	return
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v116
	F_errmsg_internal(m, int32(40796), v24+int32(80))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(489510), int32(11537), int32(333716))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	F_errmsg(m, int32(164236), int32(0))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(489510), int32(11490), int32(333716))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v1469 = F_get_constraint_name(m, v1468)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1467 + int32(4)
	F_errmsg(m, int32(672195), v24-int32(-64))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(489510), int32(11475), int32(333716))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[585]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[586]))
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
	F_errmsg_internal(m, int32(385862), int32(0))
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
	F_errfinish(m, int32(494708), int32(2892), int32(385934))
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v9 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[631]))
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	F_s_lock(m, v11, int32(494044), int32(238), int32(235731))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v27 = v22 + v24*int32(640)
	v29 = v27 + int32(84)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	if v30 == int32(0) {
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
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v62
	*(*int32)(unsafe.Add(mBase, _consts[631])) = v62
	goto L3
L10:
	;
	if v47 == int32(-1) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
	v46 = v30
	v47 = v41
	goto L10
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v30 != int32(-1) {
		v41 = v36
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v41 = v33
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v46 = v40
	v47 = v36
	goto L10
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
	goto L9
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
	goto L17
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
	goto L17
L21:
	;
	F_s_lock(m, l0, int32(494044), int32(317), int32(77574))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v79 == int32(-1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
L26:
	;
	goto L27
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v89 = v86 + v79*int32(640)
	v91 = v89 + int32(84)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	if v93 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v92 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v102 = v97
	goto L28
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86+v93*int32(640))+84)) = v92
	v102 = v93
	goto L28
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = int64(0)
	v117 = v86 + v79*int32(640)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v118 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v102
	goto L32
L34:
	;
	goto L35
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v108+v92*int32(640))+88)) = v102
	goto L32
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = v123 + v9*int32(640) + int32(84)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v129 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	if v117 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
	goto L38
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L39
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v129
	v137 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v129*int32(640))+84)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(-1)
	goto L39
L43:
	;
	F_SetLatch(m, v117+int32(20))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v118 != int32(-1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	goto L50
L48:
	;
	goto L49
L49:
	;
	return
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v165 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	F_s_lock(m, l0, int32(494044), int32(351), int32(77574))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v173 == int32(-1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v216 = v215 + v9*int32(640)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+88))
	if v217 != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v208 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v182 = v179 + v173*int32(640)
	v184 = v182 + int32(84)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+88))
	if v186 == int32(-1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v185 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v195 = v190
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179+v186*int32(640))+84)) = v185
	v195 = v186
	goto L60
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = int64(0)
	v208 = v182
	goto L56
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v195
	goto L64
L66:
	;
	goto L67
L67:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v201+v185*int32(640))+88)) = v195
	goto L64
L68:
	;
	v222 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v222
	if v208 == v222 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v221 = int32(1)
	goto L68
L70:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+84))
	if v218 != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v221 = int32(0)
	goto L68
L72:
	;
	if v221 != 0 {
		goto L50
	} else {
		goto L76
	}
L73:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v208 == v227 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_SetLatch(m, v208+int32(20))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L51
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v9 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[631]))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
		if v12 != 0 {
			F_s_lock(m, v11, int32(494044), int32(238), int32(235731))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[94]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v24 = *(*int32)(unsafe.Add(mBase, _consts[249]))
				v27 = v22 + v24*int32(640)
				v29 = v27 + int32(84)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
				if v30 == int32(0) {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v33 == int32(0) {
					} else {
						v41 = v33
						*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
						v46 = v30
						v47 = v41
						if v47 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, _consts[94]))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
						}
						*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v30 != int32(-1) {
						v41 = v36
						*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
						v46 = v30
						v47 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						v46 = v40
						v47 = v36
					}
					if v47 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _consts[94]))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
					}
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[631])) = l0
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
				if v70 != 0 {
					F_s_lock(m, l0, int32(494044), int32(75), int32(235699))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _consts[94]))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v85 = v80 + v9*int32(640) + int32(84)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v86 == int32(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
							v94 = *(*int32)(unsafe.Add(mBase, _consts[94]))
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
							*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
							*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
						return
					}
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _consts[94]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
					v85 = v80 + v9*int32(640) + int32(84)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v86 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
						v94 = *(*int32)(unsafe.Add(mBase, _consts[94]))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
						*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[94]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v24 = *(*int32)(unsafe.Add(mBase, _consts[249]))
			v27 = v22 + v24*int32(640)
			v29 = v27 + int32(84)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
			if v30 == int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				if v33 == int32(0) {
				} else {
					v41 = v33
					*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
					v46 = v30
					v47 = v41
					if v47 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _consts[94]))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
					}
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				if v30 != int32(-1) {
					v41 = v36
					*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
					v46 = v30
					v47 = v41
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v46 = v40
					v47 = v36
				}
				if v47 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, _consts[94]))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
				}
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[631])) = l0
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
			if v70 != 0 {
				F_s_lock(m, l0, int32(494044), int32(75), int32(235699))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _consts[94]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
					v85 = v80 + v9*int32(640) + int32(84)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v86 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
						v94 = *(*int32)(unsafe.Add(mBase, _consts[94]))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
						*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					return
				}
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, _consts[94]))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
				v85 = v80 + v9*int32(640) + int32(84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v86 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
					v94 = *(*int32)(unsafe.Add(mBase, _consts[94]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
					*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[631])) = l0
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
		if v70 != 0 {
			F_s_lock(m, l0, int32(494044), int32(75), int32(235699))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, _consts[94]))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
				v85 = v80 + v9*int32(640) + int32(84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v86 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
					v94 = *(*int32)(unsafe.Add(mBase, _consts[94]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
					*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				return
			}
		} else {
			v79 = *(*int32)(unsafe.Add(mBase, _consts[94]))
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
			v85 = v80 + v9*int32(640) + int32(84)
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v86 == int32(-1) {
				*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
				v94 = *(*int32)(unsafe.Add(mBase, _consts[94]))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
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
	v13 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l0^int32(-1))<<(uint(int32(2))%32))))
	v176 = base.B2i32(v19 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v24 = *(*int32)(unsafe.Add(mBase, _consts[563]))
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
	v71 = int32(4396144)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[564]))
	if l0 == v28 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v71 = int32(4396152)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	if l0 == v32 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = int32(4396160)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[566]))
	if l0 == v36 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(4396168)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	if l0 == v40 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = int32(4396176)
	goto L5
L19:
	;
	goto L20
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[568]))
	if l0 == v44 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(4396184)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[569]))
	if l0 == v48 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = int32(4396192)
	goto L5
L25:
	;
	goto L26
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[570]))
	if l0 == v52 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v71 = int32(4396200)
	goto L5
L28:
	;
	goto L29
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[571]))
	if v56 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[572]))
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
	v78 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	v88 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(227285)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(490163)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v101 = v76 + v88 - int32(40)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = int32(4194304)
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
	v132 = int32(4095884)
	v133 = *(*int32)(unsafe.Add(mBase, _consts[573]))
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
	v118 = int32(4194304)
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
	if v124&int32(262143) == int32(1) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[573])) = v150
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
	v164 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v12 == v2 {
		v101 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v101
L2:
	;
	if v12 == int32(4388736) {
		v101 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v12
	v28 = v2
	goto L4
L4:
	;
	if int32(base.Ui32(l0&int32(64))>>(uint(int32(6))%32))^int32(base.Ui32(l0&int32(2))>>(uint(int32(1))%32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v101 = v94
	goto L1
L6:
	;
	v56 = v27 - int32(12)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if int32(base.Ui32(l0)>>(uint(v57)%32))&int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v35 = v27 - int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27-int32(16))))
	v43 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v41<<(uint(int32(2))%32))+44))
	goto L9
L9:
	;
	if base.B2i32(v47 == int32(3)) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(6)
	goto L6
L11:
	;
	v63 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v94 = v28
	goto L13
L13:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v96 != int32(4388736) {
		v27 = v96
		v28 = v94
		goto L4
	} else {
		goto L25
	}
L14:
	;
	return int32(0)
L15:
	;
	if v63 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if base.Ui32(v67) <= base.Ui32(int32(17)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v94 = v28 + int32(1)
	goto L13
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v27-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
	F_errmsg_internal(m, int32(331696), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L23
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32))+uint32(_consts[474])))
	v77 = v76
	goto L22
L21:
	;
	v77 = int32(364347)
	goto L22
L22:
	;
	goto L19
L23:
	;
	F_errfinish(m, int32(490240), int32(3942), int32(278768))
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v18 = F_LWLockAcquire(m, v14+int32(512), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v22 <= int32(0) {
			v107 = v2
		} else {
			v25 = int32(1)
			v28 = v12 + int32(36)
			v30 = *(*int32)(unsafe.Add(mBase, _consts[603]))
			v31 = int32(0)
			if v22 != v25 {
				v38 = v31
				v39 = v2
				v40 = int32(0)
				for {
					v49 = v28 + v38<<(uint(int32(2))%32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v53 = v30 + v50*int32(640)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
					if v54 == int32(0) {
						v61 = v39
					} else {
						if l0 != 0 {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
							if v57 != l0 {
								v61 = v39
							} else {
								v61 = v39 + int32(1)
							}
						} else {
							v61 = v39 + int32(1)
						}
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v65 = v30 + v62*int32(640)
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
					v77 = v40 + v74
					if v77 != v22&int32(2147483646) {
						v38 = v75
						v39 = v73
						v40 = v77
						continue
					} else {
						break
					}
					break
				}
				v80 = v75
				v81 = v73
			} else {
				v80 = v31
				v81 = v2
			}
			if v22&v25 == int32(0) {
				v107 = v81
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v28+v80<<(uint(int32(2))%32))))
				v97 = v30 + v94*int32(640)
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
				if v98 == int32(0) {
					v107 = v81
				} else {
					if l0 != 0 {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+60))
						if v101 != l0 {
							v107 = v81
						} else {
							v107 = v81 + int32(1)
						}
					} else {
						v107 = v81 + int32(1)
					}
				}
			}
		}
		v116 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		F_LWLockRelease(m, v116+int32(512))
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return int32(0)
		} else {
			return v107
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v1 = l0
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L23
	} else {
		goto L66
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L23
	} else {
		goto L62
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L23
	} else {
		goto L58
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
	v197 = m.ExcPending
	if v197 != 0 {
		goto L23
	} else {
		goto L55
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
	v27 = *(*int32)(unsafe.Add(mBase, _consts[226]))
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	v39 = base.B2i32(v37 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v39)
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
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[517])))
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
	F_errmsg_internal(m, int32(509411), v15+int32(32))
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
	F_errfinish(m, int32(492600), int32(574), int32(62092))
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
	v89 = int32(4480304)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	if v94 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(230088)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(993)
	v102 = int32(4473208)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v15 + int32(68)
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v90
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v127 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v122
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
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+116)) = uint8(v158)
	v162 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L23
	} else {
		goto L48
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
	F_s_lock(m, v18, int32(492600), int32(601), int32(62092))
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
	goto L47
L47:
	;
	goto L38
L48:
	;
	if v162 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
	F_errmsg(m, int32(677100), v15+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L23
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	m.G0 = v15 + int32(96)
	return v87
L52:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v171)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v170)
	v174 = int64(32)
	v175 = int64(base.Ui64(v171) >> (uint(v174) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v175)
	v178 = int64(base.Ui64(v170) >> (uint(v174) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v178)
	F_errdetail(m, int32(636133), v15)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(492600), int32(617), int32(62092))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	F_errmsg_internal(m, int32(85477), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(492600), int32(517), int32(62092))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(332739), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(492600), int32(523), int32(62092))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v23
	F_errmsg(m, int32(358212), v15-int32(-64))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(492600), int32(534), int32(62092))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L23
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
	F_errcode(m, int32(325))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v23
	F_errmsg(m, int32(332797), v15+int32(48))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	F_errdetail(m, int32(587537), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	F_errhint(m, int32(558990), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(492600), int32(547), int32(62092))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateDirAndVersionFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v7 = m.G0
	v9 = v7 - int32(1136)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(544483)
	v18 = F_pg_sprintf(m, v9+int32(96), int32(727428), v9+int32(80))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v22 = F_mkdir(m, l0, v21)
	mBase = m.M
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L60
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L56
	}
L5:
	;
	if v22 < int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(524627)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	v40 = F_pg_snprintf(m, v9+int32(112), int32(1024), int32(175616), v9+int32(48))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v28 != int32(20) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v45 = F_OpenTransientFile(m, v9+int32(112), int32(193))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v45 < int32(0) {
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
	v62 = v45
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(167772226)
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0)
	v72 = int32(3)
	v73 = F_write(m, v62, v9+int32(96), v72)
	mBase = m.M
	if v73 != v72 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v52 != int32(20) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v58 = F_OpenTransientFile(m, v9+int32(112), int32(513))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v58 < int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v62 = v58
	goto L15
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v77 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v102 = int32(4095964)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104
	v107 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(167772225)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _consts[162])))
	if v112 != int32(1) {
		v126 = v104
		goto L32
	} else {
		goto L33
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(51)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(112)
	F_errmsg(m, int32(295466), v9+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(489260), int32(507), int32(385977))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L48
	}
L31:
	;
	if v126 == int32(0) {
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
	v117 = F_fsync(m, v62)
	mBase = m.M
	if v117 != int32(-1) {
		v126 = v117
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(-1)
	goto L32
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v121 == int32(27) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _consts[39])))
	if v132 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v135 = F_errstart(m, v133, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v133 = int32(21)
	goto L42
L41:
	;
	v133 = int32(23)
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v135 == int32(0) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(112)
	F_errmsg(m, int32(296731), v9+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(489260), int32(515), int32(385977))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L30
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = int32(0)
	v161 = F_CloseTransientFile(m, v62)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	v165 = int32(4474964)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v167 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	m.G0 = v9 + int32(1136)
	return
L53:
	;
	F_XLogRegisterData(m, v9+int32(88), int32(8))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v182 = F_XLogInsert(m, int32(4), int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v184 = int32(4474964)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v186 - int32(1)
	goto L52
L56:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l0
	F_errmsg(m, int32(293776), v9-int32(-64))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(489260), int32(478), int32(385977))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
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
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(112)
	F_errmsg(m, int32(296462), v9)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(489260), int32(495), int32(385977))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
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
		if v7 == int32(1) {
			if v21 != 0 {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
				if v42 != 0 {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
					m.T0[v71].(func(*base.Module, int32))(m, v15)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v77 = v15
						return v77
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
					if v44 != 0 {
						F_ExecReScan(m, v43)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
							v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 != 0 {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
									if v50&int32(2) == int32(0) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										F_tuplestore_select_read_pointer(m, v9, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_tuplestore_puttupleslot(m, v9, v48)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
												m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v77 = v15
													return v77
												}
											}
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
										return int32(0)
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
									return int32(0)
								}
							}
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
						v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 != 0 {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
								if v50&int32(2) == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									F_tuplestore_select_read_pointer(m, v9, v60)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_tuplestore_puttupleslot(m, v9, v48)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
											m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v77 = v15
												return v77
											}
										}
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
									return int32(0)
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
								return int32(0)
							}
						}
					}
				}
			} else {
				v33 = int32(1)
				v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 != 0 {
						v77 = v15
						return v77
					} else {
						if v7 != int32(1) {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
							m.T0[v71].(func(*base.Module, int32))(m, v15)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v77 = v15
								return v77
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
							if v42 != 0 {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = v15
									return v77
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
								if v44 != 0 {
									F_ExecReScan(m, v43)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
												if v50&int32(2) == int32(0) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v60)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v48)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
															m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return int32(0)
															} else {
																v77 = v15
																return v77
															}
														}
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										}
									}
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
									v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 != 0 {
											v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
											if v50&int32(2) == int32(0) {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												F_tuplestore_select_read_pointer(m, v9, v60)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													F_tuplestore_puttupleslot(m, v9, v48)
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return int32(0)
													} else {
														v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
														m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															v77 = v15
															return v77
														}
													}
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v56 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v21 == int32(0) {
				if v21 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
					if v42 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
						m.T0[v71].(func(*base.Module, int32))(m, v15)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v77 = v15
							return v77
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
						if v44 != 0 {
							F_ExecReScan(m, v43)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
								v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 != 0 {
										v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
										if v50&int32(2) == int32(0) {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
											F_tuplestore_select_read_pointer(m, v9, v60)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												F_tuplestore_puttupleslot(m, v9, v48)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
													m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														v77 = v15
														return v77
													}
												}
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v56 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
											return int32(0)
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
										return int32(0)
									}
								}
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
							v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 != 0 {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
									if v50&int32(2) == int32(0) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										F_tuplestore_select_read_pointer(m, v9, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_tuplestore_puttupleslot(m, v9, v48)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
												m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v77 = v15
													return v77
												}
											}
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
										return int32(0)
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
									return int32(0)
								}
							}
						}
					}
				} else {
					v33 = int32(1)
					v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v77 = v15
							return v77
						} else {
							if v7 != int32(1) {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = v15
									return v77
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
								if v42 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
									m.T0[v71].(func(*base.Module, int32))(m, v15)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v77 = v15
										return v77
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
									if v44 != 0 {
										F_ExecReScan(m, v43)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
											v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 != 0 {
													v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
													if v50&int32(2) == int32(0) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
														F_tuplestore_select_read_pointer(m, v9, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_tuplestore_puttupleslot(m, v9, v48)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int32(0)
																} else {
																	v77 = v15
																	return v77
																}
															}
														}
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
														return int32(0)
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
												if v50&int32(2) == int32(0) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v60)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v48)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
															m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return int32(0)
															} else {
																v77 = v15
																return v77
															}
														}
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+136)))
				if v27 != 0 {
					v33 = int32(1)
					v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v77 = v15
							return v77
						} else {
							if v7 != int32(1) {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = v15
									return v77
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
								if v42 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
									m.T0[v71].(func(*base.Module, int32))(m, v15)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v77 = v15
										return v77
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
									if v44 != 0 {
										F_ExecReScan(m, v43)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
											v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 != 0 {
													v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
													if v50&int32(2) == int32(0) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
														F_tuplestore_select_read_pointer(m, v9, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_tuplestore_puttupleslot(m, v9, v48)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int32(0)
																} else {
																	v77 = v15
																	return v77
																}
															}
														}
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
														return int32(0)
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
												if v50&int32(2) == int32(0) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v60)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v48)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
															m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return int32(0)
															} else {
																v77 = v15
																return v77
															}
														}
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v28 = int32(0)
					v30 = F_tuplestore_advance(m, v9, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v33 = int32(1)
							v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 != 0 {
									v77 = v15
									return v77
								} else {
									if v7 != int32(1) {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
										m.T0[v71].(func(*base.Module, int32))(m, v15)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v77 = v15
											return v77
										}
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
										if v42 != 0 {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
											m.T0[v71].(func(*base.Module, int32))(m, v15)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v77 = v15
												return v77
											}
										} else {
											v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
											if v44 != 0 {
												F_ExecReScan(m, v43)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
													v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														if v48 != 0 {
															v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
															if v50&int32(2) == int32(0) {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
																F_tuplestore_select_read_pointer(m, v9, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return int32(0)
																} else {
																	F_tuplestore_puttupleslot(m, v9, v48)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return int32(0)
																	} else {
																		v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																		m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																		mBase = m.M
																		v68 = m.ExcPending
																		if v68 != 0 {
																			return int32(0)
																		} else {
																			v77 = v15
																			return v77
																		}
																	}
																}
															} else {
																v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
																v56 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
																return int32(0)
															}
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															v56 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
															return int32(0)
														}
													}
												}
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
												v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													if v48 != 0 {
														v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
														if v50&int32(2) == int32(0) {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
															F_tuplestore_select_read_pointer(m, v9, v60)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return int32(0)
															} else {
																F_tuplestore_puttupleslot(m, v9, v48)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return int32(0)
																} else {
																	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																	m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																	mBase = m.M
																	v68 = m.ExcPending
																	if v68 != 0 {
																		return int32(0)
																	} else {
																		v77 = v15
																		return v77
																	}
																}
															}
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															v56 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
															return int32(0)
														}
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
														return int32(0)
													}
												}
											}
										}
									}
								}
							}
						} else {
							v77 = v28
							return v77
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
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(28)
	} else {
		v18 = m.Wasi_snapshot_preview1.Clock_time_get(m, l0, int64(1), v8+int32(24))
		mBase = m.M
		if v18 == int32(0) {
			v25 = int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[166])) = v18
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
	var v8 float64
	_ = v8
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v84 float64
	_ = v84
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
	var v91 float64
	_ = v91
	var v97 float64
	_ = v97
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v127 float64
	_ = v127
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
	var v134 float64
	_ = v134
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v163 float64
	_ = v163
	var v174 float64
	_ = v174
	var v179 float64
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 float64
	_ = v242
	var v243 float64
	_ = v243
	var v245 int32
	_ = v245
	var v262 float64
	_ = v262
	var v266 float64
	_ = v266
	var v268 int32
	_ = v268
	var v274 float64
	_ = v274
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v286 int32
	_ = v286
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v295 float64
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 float64
	_ = v306
	var v313 int32
	_ = v313
	var v316 float64
	_ = v316
	var v319 float64
	_ = v319
	var v329 float64
	_ = v329
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v338 float64
	_ = v338
	var v339 float64
	_ = v339
	var v340 int32
	_ = v340
	var v341 float64
	_ = v341
	var v343 int32
	_ = v343
	var v360 float64
	_ = v360
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v370 float64
	_ = v370
	var v377 int32
	_ = v377
	var v381 float64
	_ = v381
	var v383 float64
	_ = v383
	var v384 float64
	_ = v384
	var v386 float64
	_ = v386
	var v390 float64
	_ = v390
	var v394 float64
	_ = v394
	var v404 float64
	_ = v404
	var v425 float64
	_ = v425
	var v437 float64
	_ = v437
	var v447 float64
	_ = v447
	var v449 float64
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 float64
	_ = v469
	var v470 int32
	_ = v470
	var v473 float64
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 float64
	_ = v480
	var v486 float64
	_ = v486
	var v489 float64
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 float64
	_ = v499
	var v500 float64
	_ = v500
	var v509 float64
	_ = v509
	var v520 float64
	_ = v520
	var v525 float64
	_ = v525
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 float64
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 float64
	_ = v583
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 float64
	_ = v588
	var v589 float64
	_ = v589
	var v591 int32
	_ = v591
	var v608 float64
	_ = v608
	var v612 float64
	_ = v612
	var v614 int32
	_ = v614
	var v620 float64
	_ = v620
	var v623 float64
	_ = v623
	var v624 float64
	_ = v624
	var v632 int32
	_ = v632
	var v638 float64
	_ = v638
	var v640 float64
	_ = v640
	var v641 float64
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 float64
	_ = v652
	var v659 int32
	_ = v659
	var v662 float64
	_ = v662
	var v665 float64
	_ = v665
	var v675 float64
	_ = v675
	var v680 int32
	_ = v680
	var v681 float64
	_ = v681
	var v684 float64
	_ = v684
	var v685 float64
	_ = v685
	var v686 int32
	_ = v686
	var v687 float64
	_ = v687
	var v689 int32
	_ = v689
	var v706 float64
	_ = v706
	var v710 float64
	_ = v710
	var v711 float64
	_ = v711
	var v716 float64
	_ = v716
	var v723 int32
	_ = v723
	var v727 float64
	_ = v727
	var v729 float64
	_ = v729
	var v730 float64
	_ = v730
	var v732 float64
	_ = v732
	var v736 float64
	_ = v736
	var v740 float64
	_ = v740
	var v750 float64
	_ = v750
	var v771 float64
	_ = v771
	var v781 float64
	_ = v781
	var v793 float64
	_ = v793
	v8 = float64(0)
	v21 = l4 - int32(1)
	v34 = int32(-1)
	v35 = v21
	goto L1
L1:
	;
	v44 = base.I32_div_s(v34+v35+int32(1), int32(2))
	v48 = F_range_cmp_bounds(m, l0, l3+v44<<(uint(int32(3))%32), l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v54 < int32(0) {
		v793 = v8
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v53 = base.B2i32(v48 <= int32(0))
	if v48 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v54 = v44
	goto L7
L6:
	;
	v54 = v34
	goto L7
L7:
	;
	if v48 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v35
	goto L10
L9:
	;
	v57 = v44 - int32(1)
	goto L10
L10:
	;
	if v54 < v57 {
		v34 = v54
		v35 = v57
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	return v793
L13:
	;
	v62 = l0 + int32(268)
	v64 = l4 - int32(2)
	if v54 < v64 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = v54
	goto L16
L15:
	;
	v66 = v64
	goto L16
L16:
	;
	v69 = l3 + v66<<(uint(int32(3))%32)
	v72 = F_get_position(m, l0, l1, v69, v69+int32(8))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v74 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v66 < int32(0) {
		v793 = v8
		goto L12
	} else {
		goto L38
	}
L19:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v102 != int32(1) {
		v112 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L18
	} else {
		goto L34
	}
L22:
	;
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L24
L23:
	;
	v80 = float64(1)
	goto L24
L24:
	;
	if v79 != 0 {
		v112 = v80
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v81 == int32(0) {
		v112 = v80
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v84 = float64(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v89 = F_FunctionCall2Coll(m, v62, v86, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v91)&int64(9223372036854775807)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v97 = v84
	goto L30
L29:
	;
	v97 = v91
	goto L30
L30:
	;
	if base.F64_lt(v91, float64(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v100 = v84
	goto L33
L32:
	;
	v100 = v97
	goto L33
L33:
	;
	v112 = v100
	goto L18
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v107 == v108 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v110 = float64(0)
	goto L37
L36:
	;
	v110 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L37
L37:
	;
	v112 = v110
	goto L18
L38:
	;
	v116 = base.F64_convert_i32_u(v21)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	if v117 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v163 = float64(0)
	if base.F64_lt(v154, v163) != 0 {
		v425 = v163
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v122 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v145 != int32(1) {
		v154 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L39
	} else {
		goto L55
	}
L43:
	;
	v123 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L44:
	;
	v123 = float64(1)
	goto L45
L45:
	;
	if v122 != 0 {
		v154 = v123
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v124 == int32(0) {
		v154 = v123
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v127 = float64(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v132 = F_FunctionCall2Coll(m, v62, v129, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v132)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v134)&int64(9223372036854775807)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v140 = v127
	goto L51
L50:
	;
	v140 = v134
	goto L51
L51:
	;
	if base.F64_lt(v134, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v143 = v127
	goto L54
L53:
	;
	v143 = v140
	goto L54
L54:
	;
	v154 = v143
	goto L39
L55:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+6)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v150 == v151 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v153 = float64(0)
	goto L58
L57:
	;
	v153 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L58
L58:
	;
	v154 = v153
	goto L39
L59:
	;
	v437 = base.F64_add(base.F64_div(base.F64_mul(v72, base.F64_sub(float64(1), v425)), v116), float64(0))
	if v66 == int32(0) {
		v793 = v437
		goto L12
	} else {
		goto L135
	}
L60:
	;
	goto L59
L61:
	;
	v174 = float64(1)
	v179 = base.F64_abs(v154)
	if int32(0)&base.F64_eq(v179, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v425 = v174
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v185 = l6 - int32(1)
	if v185 < int32(0) {
		v425 = v174
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v189 = v185
	v193 = int32(-1)
	goto L64
L64:
	;
	v210 = int32(2)
	v211 = base.I32_div_s(v189+v193+int32(1), v210)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l5+v211<<(uint(v210)%32))))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v215)))
	if base.F64_gt(v112, v216) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v185 <= v226 {
		v425 = v174
		goto L60
	} else {
		goto L77
	}
L66:
	;
	if v226 < v224 {
		v189 = v224
		v193 = v226
		goto L64
	} else {
		goto L76
	}
L67:
	;
	v224 = v189
	v226 = v211
	goto L66
L68:
	;
	goto L69
L69:
	;
	v221 = int32(0) & base.F64_ge(v112, v216)
	if v221 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v222 = v189
	goto L72
L71:
	;
	v222 = v211 - int32(1)
	goto L72
L72:
	;
	if v221 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v223 = v211
	goto L75
L74:
	;
	v223 = v193
	goto L75
L75:
	;
	v224 = v222
	v226 = v223
	goto L66
L76:
	;
	goto L65
L77:
	;
	if v226 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v277 = base.F64_convert_i32_s(v185)
	v278 = base.F64_div(base.F64_add(v274, base.F64_convert_i32_u(v268)), v277)
	if base.F64_eq(v112, v154) != 0 {
		v425 = v278
		goto L60
	} else {
		goto L95
	}
L79:
	;
	v268 = int32(0)
	v274 = float64(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v235 = l5 + v226<<(uint(int32(2))%32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v236)))
	v238 = base.F64_abs(v237)
	v239 = math.Float64frombits(uint64(0x7ff0000000000000))
	v240 = base.F64_eq(v238, v239)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v241)))
	v243 = base.F64_abs(v242)
	v245 = base.F64_eq(v243, v239)
	if v245 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v240 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	if v240 != 0 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	if base.F64_eq(base.F64_abs(v112), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v268 = v226
		v274 = float64(0.5)
		goto L78
	} else {
		goto L85
	}
L85:
	;
	v268 = v226
	v274 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v237, v112), base.F64_sub(v237, v242)))
	goto L78
L86:
	;
	if base.F64_eq(v238, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	if v245 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v268 = v226
	v274 = float64(1)
	goto L78
L89:
	;
	v262 = float64(0)
	goto L91
L90:
	;
	v262 = float64(0.5)
	goto L91
L91:
	;
	if base.F64_eq(v243, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v266 = v262
	goto L94
L93:
	;
	v266 = float64(0.5)
	goto L94
L94:
	;
	v268 = v226
	v274 = v266
	goto L78
L95:
	;
	if v185 <= v268 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v390 = float64(0)
	v394 = base.F64_div(base.F64_add(v386, base.F64_convert_i32_u(v377)), v277)
	if base.F64_gt(v381, v390)|base.F64_gt(v394, v390) != 0 {
		goto L128
	} else {
		goto L129
	}
L97:
	;
	v377 = v268
	v381 = v278
	v383 = v112
	v384 = v163
	v386 = v163
	goto L96
L98:
	;
	goto L99
L99:
	;
	v286 = v268
	v292 = v278
	v294 = v163
	v295 = v112
	goto L101
L100:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l5+v286<<(uint(int32(2))%32))))
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v334)))
	if base.F64_eq(v306, v335) != 0 {
		goto L111
	} else {
		goto L112
	}
L101:
	;
	v300 = int32(1)
	v301 = v286 + v300
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l5+v301<<(uint(int32(2))%32))))
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v305)))
	if base.B2i32(base.F64_ge(v154, v306) == int32(0))|int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v377 = v185
	v381 = v319
	v383 = v306
	v384 = v329
	v386 = v163
	goto L96
L103:
	;
	v313 = base.F64_lt(v306, v154)
	goto L105
L104:
	;
	v313 = v300
	goto L105
L105:
	;
	if v313 == int32(0) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v316 = float64(0)
	v319 = base.F64_div(base.F64_convert_i32_u(v286), v277)
	if base.F64_gt(v292, v316)|base.F64_gt(v319, v316) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v329 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v292, v319), float64(0.5)), base.F64_sub(v306, v295)), v294)
	goto L109
L108:
	;
	v329 = v294
	goto L109
L109:
	;
	if v185 != v301 {
		v286 = v301
		v292 = v319
		v294 = v329
		v295 = v306
		goto L101
	} else {
		goto L110
	}
L110:
	;
	goto L102
L111:
	;
	v370 = float64(0)
	goto L113
L112:
	;
	v338 = base.F64_abs(v306)
	v339 = math.Float64frombits(uint64(0x7ff0000000000000))
	v340 = base.F64_eq(v338, v339)
	v341 = base.F64_abs(v335)
	v343 = base.F64_eq(v341, v339)
	if v343 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v377 = v286
	v381 = v292
	v383 = v295
	v384 = v294
	v386 = v370
	goto L96
L114:
	;
	v370 = v365
	goto L113
L115:
	;
	if v340 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	if v340 != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	if base.F64_eq(base.F64_abs(v154), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v365 = float64(0.5)
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v365 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v306, v154), base.F64_sub(v306, v335)))
	goto L114
L119:
	;
	if base.F64_eq(v338, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v343 == int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v365 = float64(1)
	goto L114
L122:
	;
	v360 = float64(0)
	goto L124
L123:
	;
	v360 = float64(0.5)
	goto L124
L124:
	;
	if base.F64_eq(v341, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v364 = v360
	goto L127
L126:
	;
	v364 = float64(0.5)
	goto L127
L127:
	;
	v365 = v364
	goto L114
L128:
	;
	v404 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v381, v394), float64(0.5)), base.F64_sub(v154, v383)), v384)
	goto L130
L129:
	;
	v404 = v384
	goto L130
L130:
	;
	if base.F64_eq(v179, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if base.F64_eq(base.F64_abs(v404), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v425 = float64(0.5)
		goto L60
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v425 = base.F64_div(v404, base.F64_sub(v154, v112))
	goto L60
L134:
	;
	goto L133
L135:
	;
	v447 = v154
	v449 = v437
	v454 = v66
	goto L136
L136:
	;
	v459 = v454 - int32(1)
	v462 = l3 + v459<<(uint(int32(3))%32)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+4)))
	if v463 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v793 = v781
	goto L12
L138:
	;
	v509 = float64(0)
	if base.F64_lt(v500, v509) != 0 {
		v771 = v509
		goto L159
	} else {
		goto L160
	}
L139:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v468 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L141
L141:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v491 != int32(1) {
		v500 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L138
	} else {
		goto L154
	}
L142:
	;
	v469 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L144
L143:
	;
	v469 = float64(1)
	goto L144
L144:
	;
	if v468 != 0 {
		v500 = v469
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v470 == int32(0) {
		v500 = v469
		goto L138
	} else {
		goto L146
	}
L146:
	;
	v473 = float64(1)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v478 = F_FunctionCall2Coll(m, v62, v475, v476, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	v480 = *(*float64)(unsafe.Add(mBase, uint32(v478)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v480)&int64(9223372036854775807)) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v486 = v473
	goto L150
L149:
	;
	v486 = v480
	goto L150
L150:
	;
	if base.F64_lt(v480, float64(0)) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v489 = v473
	goto L153
L152:
	;
	v489 = v486
	goto L153
L153:
	;
	v500 = v489
	goto L138
L154:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+6)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v496 == v497 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v499 = float64(0)
	goto L157
L156:
	;
	v499 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L157
L157:
	;
	v500 = v499
	goto L138
L158:
	;
	v781 = base.F64_add(v449, base.F64_div(base.F64_sub(float64(1), v771), v116))
	if base.Ui32(int32(1)) < base.Ui32(v454) {
		v447 = v500
		v449 = v781
		v454 = v459
		goto L136
	} else {
		goto L234
	}
L159:
	;
	goto L158
L160:
	;
	v520 = float64(1)
	v525 = base.F64_abs(v500)
	if int32(0)&base.F64_eq(v525, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v771 = v520
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v531 = l6 - int32(1)
	if v531 < int32(0) {
		v771 = v520
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v535 = v531
	v539 = int32(-1)
	goto L163
L163:
	;
	v556 = int32(2)
	v557 = base.I32_div_s(v535+v539+int32(1), v556)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l5+v557<<(uint(v556)%32))))
	v562 = *(*float64)(unsafe.Add(mBase, uint32(v561)))
	if base.F64_gt(v447, v562) != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	if v531 <= v572 {
		v771 = v520
		goto L159
	} else {
		goto L176
	}
L165:
	;
	if v572 < v570 {
		v535 = v570
		v539 = v572
		goto L163
	} else {
		goto L175
	}
L166:
	;
	v570 = v535
	v572 = v557
	goto L165
L167:
	;
	goto L168
L168:
	;
	v567 = int32(0) & base.F64_ge(v447, v562)
	if v567 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v568 = v535
	goto L171
L170:
	;
	v568 = v557 - int32(1)
	goto L171
L171:
	;
	if v567 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v569 = v557
	goto L174
L173:
	;
	v569 = v539
	goto L174
L174:
	;
	v570 = v568
	v572 = v569
	goto L165
L175:
	;
	goto L164
L176:
	;
	if v572 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v623 = base.F64_convert_i32_s(v531)
	v624 = base.F64_div(base.F64_add(v620, base.F64_convert_i32_u(v614)), v623)
	if base.F64_eq(v447, v500) != 0 {
		v771 = v624
		goto L159
	} else {
		goto L194
	}
L178:
	;
	v614 = int32(0)
	v620 = float64(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v581 = l5 + v572<<(uint(int32(2))%32)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v583 = *(*float64)(unsafe.Add(mBase, uint32(v582)))
	v584 = base.F64_abs(v583)
	v585 = math.Float64frombits(uint64(0x7ff0000000000000))
	v586 = base.F64_eq(v584, v585)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v588 = *(*float64)(unsafe.Add(mBase, uint32(v587)))
	v589 = base.F64_abs(v588)
	v591 = base.F64_eq(v589, v585)
	if v591 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v586 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	if v586 != 0 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	if base.F64_eq(base.F64_abs(v447), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v614 = v572
		v620 = float64(0.5)
		goto L177
	} else {
		goto L184
	}
L184:
	;
	v614 = v572
	v620 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v583, v447), base.F64_sub(v583, v588)))
	goto L177
L185:
	;
	if base.F64_eq(v584, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	if v591 == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v614 = v572
	v620 = float64(1)
	goto L177
L188:
	;
	v608 = float64(0)
	goto L190
L189:
	;
	v608 = float64(0.5)
	goto L190
L190:
	;
	if base.F64_eq(v589, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v612 = v608
	goto L193
L192:
	;
	v612 = float64(0.5)
	goto L193
L193:
	;
	v614 = v572
	v620 = v612
	goto L177
L194:
	;
	if v531 <= v614 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v736 = float64(0)
	v740 = base.F64_div(base.F64_add(v732, base.F64_convert_i32_u(v723)), v623)
	if base.F64_gt(v727, v736)|base.F64_gt(v740, v736) != 0 {
		goto L227
	} else {
		goto L228
	}
L196:
	;
	v723 = v614
	v727 = v624
	v729 = v447
	v730 = v509
	v732 = v509
	goto L195
L197:
	;
	goto L198
L198:
	;
	v632 = v614
	v638 = v624
	v640 = v509
	v641 = v447
	goto L200
L199:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l5+v632<<(uint(int32(2))%32))))
	v681 = *(*float64)(unsafe.Add(mBase, uint32(v680)))
	if base.F64_eq(v652, v681) != 0 {
		goto L210
	} else {
		goto L211
	}
L200:
	;
	v646 = int32(1)
	v647 = v632 + v646
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l5+v647<<(uint(int32(2))%32))))
	v652 = *(*float64)(unsafe.Add(mBase, uint32(v651)))
	if base.B2i32(base.F64_ge(v500, v652) == int32(0))|int32(1) != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v723 = v531
	v727 = v665
	v729 = v652
	v730 = v675
	v732 = v509
	goto L195
L202:
	;
	v659 = base.F64_lt(v652, v500)
	goto L204
L203:
	;
	v659 = v646
	goto L204
L204:
	;
	if v659 == int32(0) {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v662 = float64(0)
	v665 = base.F64_div(base.F64_convert_i32_u(v632), v623)
	if base.F64_gt(v638, v662)|base.F64_gt(v665, v662) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v675 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v638, v665), float64(0.5)), base.F64_sub(v652, v641)), v640)
	goto L208
L207:
	;
	v675 = v640
	goto L208
L208:
	;
	if v531 != v647 {
		v632 = v647
		v638 = v665
		v640 = v675
		v641 = v652
		goto L200
	} else {
		goto L209
	}
L209:
	;
	goto L201
L210:
	;
	v716 = float64(0)
	goto L212
L211:
	;
	v684 = base.F64_abs(v652)
	v685 = math.Float64frombits(uint64(0x7ff0000000000000))
	v686 = base.F64_eq(v684, v685)
	v687 = base.F64_abs(v681)
	v689 = base.F64_eq(v687, v685)
	if v689 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v723 = v632
	v727 = v638
	v729 = v641
	v730 = v640
	v732 = v716
	goto L195
L213:
	;
	v716 = v711
	goto L212
L214:
	;
	if v686 != 0 {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	if v686 != 0 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	if base.F64_eq(base.F64_abs(v500), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v711 = float64(0.5)
		goto L213
	} else {
		goto L217
	}
L217:
	;
	v711 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v652, v500), base.F64_sub(v652, v681)))
	goto L213
L218:
	;
	if base.F64_eq(v684, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	if v689 == int32(0) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v711 = float64(1)
	goto L213
L221:
	;
	v706 = float64(0)
	goto L223
L222:
	;
	v706 = float64(0.5)
	goto L223
L223:
	;
	if base.F64_eq(v687, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v710 = v706
	goto L226
L225:
	;
	v710 = float64(0.5)
	goto L226
L226:
	;
	v711 = v710
	goto L213
L227:
	;
	v750 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v727, v740), float64(0.5)), base.F64_sub(v500, v729)), v730)
	goto L229
L228:
	;
	v750 = v730
	goto L229
L229:
	;
	if base.F64_eq(v525, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if base.F64_eq(base.F64_abs(v750), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v771 = float64(0.5)
		goto L159
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v771 = base.F64_div(v750, base.F64_sub(v500, v447))
	goto L159
L233:
	;
	goto L232
L234:
	;
	goto L137
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
	var v56 float64
	_ = v56
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
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
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v221 int32
	_ = v221
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
	var v275 int32
	_ = v275
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
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
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
	var v465 float64
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 float64
	_ = v479
	var v482 int32
	_ = v482
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
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
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
	var v629 int32
	_ = v629
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
	var v651 float64
	_ = v651
	var v652 int32
	_ = v652
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v681 float64
	_ = v681
	var v682 int32
	_ = v682
	var v684 float64
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v698 float64
	_ = v698
	var v700 float64
	_ = v700
	var v719 float64
	_ = v719
	var v721 float64
	_ = v721
	var v725 float64
	_ = v725
	var v727 float64
	_ = v727
	var v729 float64
	_ = v729
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v755 float64
	_ = v755
	var v756 float64
	_ = v756
	var v764 float64
	_ = v764
	var v768 float64
	_ = v768
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
		goto L123
	} else {
		goto L124
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
	v56 = float64(1)
	v59 = l7
	v72 = v12
	goto L8
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v72<<(uint(int32(2))%32))))
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
		goto L118
	} else {
		goto L119
	}
L10:
	;
	v499 = v72 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v499 < v500 {
		v56 = v479
		v59 = v482
		v72 = v499
		goto L8
	} else {
		goto L117
	}
L11:
	;
	if l7 == v59 {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v32&int32(-2) != int32(4) {
		goto L25
	} else {
		goto L26
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
		v479 = v56
		v482 = v59
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
	if v94 == int32(0) {
		v479 = v56
		v482 = v59
		goto L10
	} else {
		goto L23
	}
L23:
	;
	if v47 == int32(0) {
		v479 = v56
		v482 = v59
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v152 = int32(0)
	goto L11
L25:
	;
	v152 = int32(0)
	goto L11
L26:
	;
	goto L27
L27:
	;
	if v42 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v149 != int32(1) {
		v479 = v56
		v482 = v59
		goto L10
	} else {
		goto L44
	}
L29:
	;
	v149 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v111 = int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v112 <= v111 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v115 = v111
	goto L34
L33:
	;
	v115 = v112
	goto L34
L34:
	;
	v118 = int32(0)
	v120 = v118
	v121 = v118
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(8)+v120<<(uint(int32(2))%32))))
	if v129 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v149 = v142
	goto L28
L37:
	;
	goto L36
L38:
	;
	v130 = int32(2)
	if v121 != 0 {
		v142 = v130
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v135 = v121
	goto L40
L40:
	;
	v138 = v120 + int32(1)
	if v138 != v115 {
		v120 = v138
		v121 = v135
		goto L35
	} else {
		goto L43
	}
L41:
	;
	v131 = int32(1)
	if base.Ui32(v131) < base.Ui32(base.I32_popcnt(v129)) {
		v142 = v130
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v135 = v131
	goto L40
L43:
	;
	v142 = v135
	goto L37
L44:
	;
	v152 = int32(1)
	goto L11
L45:
	;
	v154 = F_list_copy(m, v59)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	v156 = v59
	goto L47
L47:
	;
	if v156 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v156 = v154
	goto L47
L49:
	;
	v159 = int32(0)
	v161 = F_list_concat(m, v159, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v166 = v79 + int32(288)
	v167 = int32(0)
	v180 = v156
	v185 = v167
	v186 = v167
	goto L53
L52:
	;
	v479 = v56
	v482 = v161
	goto L10
L53:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v185 < v196 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v367 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if int32(0) < v198 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v361 = v180
	v367 = v186
	goto L57
L57:
	;
	goto L54
L58:
	;
	if v332 != 0 {
		v180 = v332
		v185 = v324 + int32(1)
		v186 = v338
		goto L53
	} else {
		goto L87
	}
L59:
	;
	v317 = F_list_delete_nth_cell(m, v180, v185)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L85
	}
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v185<<(uint(int32(2))%32))))
	v210 = int32(0)
	v221 = v198
	goto L63
L61:
	;
	goto L62
L62:
	;
	v324 = v185
	v332 = v180
	v338 = v186
	goto L58
L63:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v205)+60))
	if v234 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L62
L65:
	;
	v286 = v210 + int32(1)
	if v286 < v284 {
		v210 = v286
		v221 = v284
		goto L63
	} else {
		goto L84
	}
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v166+v210<<(uint(int32(2))%32))))
	if v238 != v234 {
		v284 = v221
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(544)+v210<<(uint(int32(2))%32))))
	v244 = int32(0)
	if v243 == v244 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L59
L70:
	;
	if v282 != 0 {
		goto L59
	} else {
		goto L83
	}
L71:
	;
	v282 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v250 <= int32(0) {
		v275 = v244
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v282 = v275
	goto L70
L75:
	;
	v253 = int32(0)
	if v253 < v250 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v256 = v250
	goto L78
L77:
	;
	v256 = v253
	goto L78
L78:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v259 = int32(0)
	goto L79
L79:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v257+v259<<(uint(int32(2))%32))))
	v268 = base.B2i32(v267 == v205)
	if v267 == v205 {
		v275 = v268
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v275 = v268
	goto L74
L81:
	;
	v270 = v259 + int32(1)
	if v270 != v256 {
		v259 = v270
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v284 = v283
	goto L65
L84:
	;
	goto L64
L85:
	;
	v319 = F_lappend(m, v186, v205)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	v324 = v185 - int32(1)
	v332 = v317
	v338 = v319
	goto L58
L87:
	;
	v361 = v332
	v367 = v338
	goto L57
L88:
	;
	v380 = F_list_concat(m, v361, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v79)+284))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v79)+272))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v382 != v383+(v384-v385) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v479 = v56
	v482 = v380
	goto L10
L92:
	;
	v389 = F_list_concat(m, v361, v367)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v392 = F_find_base_rel(m, l0, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L13
	} else {
		goto L96
	}
L95:
	;
	v479 = v56
	v482 = v389
	goto L10
L96:
	;
	v394 = *(*float64)(unsafe.Add(mBase, uint32(v392)+120))
	v395 = float64(1)
	if base.F64_gt(v394, v395) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v398 = v394
	goto L99
L98:
	;
	v398 = v395
	goto L99
L99:
	;
	if v152 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v392)+16))
	v401 = v399
	goto L102
L101:
	;
	v401 = float64(1)
	goto L102
L102:
	;
	v403 = base.F64_mul(v56, base.F64_div(v401, v398))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v404 <= int32(0) {
		v479 = v403
		v482 = v361
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v407 <= int32(0) {
		v479 = v403
		v482 = v361
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v416 = int32(0)
	v421 = v403
	goto L105
L105:
	;
	v441 = v416 << (uint(int32(2)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v166+v441)))
	if v443 == int32(0) {
		v465 = v421
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v479 = v465
	v482 = v361
	goto L10
L107:
	;
	v468 = v416 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v468 < v469 {
		v416 = v468
		v421 = v465
		goto L105
	} else {
		goto L116
	}
L108:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+40)))
	if v446 != int32(1) {
		v465 = v421
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v441+(v79+int32(416)))))
	v451 = int32(0)
	v453 = F_ec_search_derived_clause_for_ems(m, l0, v443, v450, v451, v451)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	if v453 == int32(0) {
		v465 = v421
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v458 = F_clause_selectivity(m, l0, v453, int32(0), v32, l6)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	if base.F64_gt(v458, float64(0)) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v463 = base.F64_div(v421, v458)
	goto L115
L114:
	;
	v463 = v421
	goto L115
L115:
	;
	v465 = v463
	goto L107
L116:
	;
	goto L106
L117:
	;
	goto L9
L118:
	;
	v516 = v482
	v518 = v508
	goto L1
L119:
	;
	v505 = float64(1)
	if base.F64_gt(v479, v505) != 0 {
		v508 = v505
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v516 = v482
	v518 = v479
	goto L1
L121:
	;
	switch v32 {
	case 0:
		goto L159
	case 1:
		goto L164
	case 2:
		goto L163
	default:
		goto L160
	case 4:
		goto L162
	case 5:
		goto L161
	}
L122:
	;
	v681 = F_clauselist_selectivity(m, l0, v667, int32(0), v32, l6)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L13
	} else {
		goto L154
	}
L123:
	;
	v540 = int32(0)
	if v516 == v540 {
		v666 = v540
		v667 = v540
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v651 = F_clauselist_selectivity(m, l0, v516, int32(0), v32, l6)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L13
	} else {
		goto L153
	}
L126:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v544 <= int32(0) {
		v666 = v540
		v667 = v540
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v551 = int32(0)
	v561 = v540
	v562 = v540
	goto L128
L128:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v575+v551<<(uint(int32(2))%32))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+8)))
	if v580 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v666 = v643
	v667 = v644
	goto L122
L130:
	;
	v646 = v551 + int32(1)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v646 < v647 {
		v551 = v646
		v561 = v643
		v562 = v644
		goto L128
	} else {
		goto L152
	}
L131:
	;
	v641 = F_lappend(m, v562, v579)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L151
	}
L132:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v579)+32))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v585 = int32(0)
	if v583 == v585 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	goto L134
L134:
	;
	v639 = F_lappend(m, v561, v579)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L150
	}
L135:
	;
	if v638 != 0 {
		goto L131
	} else {
		goto L149
	}
L136:
	;
	v638 = int32(1)
	goto L135
L137:
	;
	goto L138
L138:
	;
	if v584 == int32(0) {
		v629 = v585
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v638 = v629
	goto L135
L140:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v595 < v594 {
		v629 = v585
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v597 = int32(1)
	if v594 <= v597 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v600 = v597
	goto L144
L143:
	;
	v600 = v594
	goto L144
L144:
	;
	v601 = int32(8)
	v606 = int32(0)
	goto L145
L145:
	;
	v613 = v606 << (uint(int32(2)) % 32)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v583+v601+v613)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613+(v584+v601))))
	v620 = v615 & (v617 ^ int32(-1))
	v622 = base.B2i32(v620 == int32(0))
	if v620 != 0 {
		v629 = v622
		goto L139
	} else {
		goto L147
	}
L146:
	;
	v629 = v622
	goto L139
L147:
	;
	v624 = v606 + int32(1)
	if v624 != v600 {
		v606 = v624
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	goto L134
L150:
	;
	v643 = v639
	v644 = v562
	goto L130
L151:
	;
	v643 = v561
	v644 = v641
	goto L130
L152:
	;
	goto L129
L153:
	;
	v698 = v651
	v700 = float64(0)
	goto L121
L154:
	;
	v684 = F_clauselist_selectivity(m, l0, v666, int32(0), v32, l6)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	F_list_free(m, v667)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	F_list_free(m, v666)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	v698 = v681
	v700 = v684
	goto L121
L158:
	;
	v756 = float64(1e+100)
	if base.F64_gt(v755, v756) != 0 {
		v768 = v756
		goto L177
	} else {
		goto L178
	}
L159:
	;
	v755 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v698)
	goto L158
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L13
	} else {
		goto L174
	}
L161:
	;
	v755 = base.F64_mul(v700, base.F64_mul(l4, base.F64_sub(float64(1), base.F64_mul(v518, v698))))
	goto L158
L162:
	;
	v755 = base.F64_mul(base.F64_mul(l4, v518), v698)
	goto L158
L163:
	;
	v725 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v698)
	if base.F64_gt(l4, v725) != 0 {
		goto L168
	} else {
		goto L169
	}
L164:
	;
	v719 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v698)
	if base.F64_gt(l4, v719) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v721 = l4
	goto L167
L166:
	;
	v721 = v719
	goto L167
L167:
	;
	v755 = base.F64_mul(v700, v721)
	goto L158
L168:
	;
	v727 = l4
	goto L170
L169:
	;
	v727 = v725
	goto L170
L170:
	;
	if base.F64_gt(l5, v727) != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v729 = l5
	goto L173
L172:
	;
	v729 = v727
	goto L173
L173:
	;
	v755 = base.F64_mul(v700, v729)
	goto L158
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v32
	F_errmsg_internal(m, int32(480285), v30)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(493365), int32(5627), int32(351063))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L13
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	m.G0 = v30 + int32(16)
	return v768
L178:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v755)&int64(9223372036854775807)) {
		v768 = v756
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v764 = float64(1)
	if base.F64_le(v755, v764) != 0 {
		v768 = v764
		goto L177
	} else {
		goto L180
	}
L180:
	;
	v768 = base.F64_nearest(v755)
	goto L177
}
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 float32
	_ = v31
	var v37 float64
	_ = v37
	var v40 float32
	_ = v40
	var v46 float64
	_ = v46
	var v50 float64
	_ = v50
	var v54 float32
	_ = v54
	var v60 float64
	_ = v60
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v70 float32
	_ = v70
	var v76 float64
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var __phi573 int32
	_ = __phi573
	var v575 int32
	_ = v575
	var __phi575 int32
	_ = __phi575
	var v577 int32
	_ = v577
	var __phi577 int32
	_ = __phi577
	var v580 int32
	_ = v580
	var __phi580 int32
	_ = __phi580
	var v582 int32
	_ = v582
	var __phi582 int32
	_ = __phi582
	var v583 int32
	_ = v583
	var __phi583 int32
	_ = __phi583
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 float64
	_ = v673
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 float64
	_ = v693
	var v694 float64
	_ = v694
	var v695 float64
	_ = v695
	var v711 int32
	_ = v711
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1197 int32
	_ = v1197
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1243 int32
	_ = v1243
	var v1244 float64
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1265 float64
	_ = v1265
	var v1272 int32
	_ = v1272
	var v1278 float64
	_ = v1278
	var v1279 float64
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1302 float64
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1328 float64
	_ = v1328
	var v1338 float64
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1465 float64
	_ = v1465
	var v1466 float64
	_ = v1466
	var v1486 float64
	_ = v1486
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1513 int32
	_ = v1513
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1575 float64
	_ = v1575
	var v1594 float64
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1605 float64
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1616 float64
	_ = v1616
	var v1621 float64
	_ = v1621
	var v1627 float64
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1658 float32
	_ = v1658
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 + int32(-64)
	m.G0 = v27
	v31 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	if base.F32_ge(v31, float32(0)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L20
	} else {
		goto L245
	}
L2:
	;
	if base.F32_gt(v31, float32(1)) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v37 = float64(0.10000000149011612)
	goto L4
L4:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = base.F64_div(float64(1), v37)
	v40 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.F32_ge(v40, float32(0)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v37 = base.F64_promote_f32(v31)
	goto L4
L6:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = base.F64_div(float64(1), v50)
	v54 = *(*float32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F32_ge(v54, float32(0)) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v50 = float64(0.20000000298023224)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v46 = base.F64_promote_f32(v40)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v46
	if base.F32_gt(v40, float32(1)) != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v50 = v46
	goto L6
L11:
	;
	v65 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = base.F64_div(v65, v64)
	v70 = *(*float32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.F32_ge(v70, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v64 = float64(0.4000000059604645)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v60 = base.F64_promote_f32(v54)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v60
	if base.F32_gt(v54, float32(1)) != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v64 = v60
	goto L11
L16:
	;
	if base.F32_gt(v70, float32(1)) != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v76 = v65
	goto L18
L18:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+40)) = base.F64_div(v65, v76)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = l2
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v83 = F_palloc0(m, v80*int32(32776))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v76 = base.F64_promote_f32(v70)
	goto L18
L20:
	;
	return float32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v91 = F_palloc(m, v88*int32(48))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v93 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	m.G0 = v27 - int32(-64)
	return v1658
L24:
	;
	F_pg_qsort(m, v498, v496, int32(12), int32(1533))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L20
	} else {
		goto L94
	}
L25:
	;
	v96 = int32(8)
	v99 = l1 + v96
	v104 = l2
	v106 = v88 << (uint(int32(2)) % 32)
	v108 = v5
	v110 = v91
	v114 = v5
	goto L28
L26:
	;
	v522 = v83
	v530 = v91
	goto L27
L27:
	;
	F_pfree(m, v530)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L20
	} else {
		goto L92
	}
L28:
	;
	v128 = l2 + v96 + v114*int32(12)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v129 != int32(1) {
		v494 = v106
		v496 = v108
		v498 = v110
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if int32(0) < v496 {
		goto L24
	} else {
		goto L91
	}
L30:
	;
	v515 = v114 + int32(1)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v515 < v517 {
		v104 = v516
		v106 = v494
		v108 = v496
		v110 = v498
		v114 = v515
		goto L28
	} else {
		goto L90
	}
L31:
	;
	v133 = v25 + int32(-4)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(0)
	v144 = l1 + int32(8)
	v147 = v144 + v140<<(uint(int32(2))%32)
	if base.Ui32(v147) <= base.Ui32(v144) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v286 == int32(0) {
		v494 = v106
		v496 = v108
		v498 = v110
		goto L30
	} else {
		goto L62
	}
L33:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)))
	if v214 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	v208 = v144
	v209 = v147
	v210 = v147
	goto L33
L35:
	;
	goto L36
L36:
	;
	v155 = v144
	v157 = v147
	goto L37
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v162 = int32(12)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v172 = int32(2)
	v179 = base.I32_div_s((v157-v155)>>(uint(v172)%32), v172)
	v182 = v155 + v179<<(uint(v172)%32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v191 = int32(0)
	v192 = F_tsCompareString(m, v104+int32(8)+v161*v162+int32(base.Ui32(v165)>>(uint(v162)%32)), v165&int32(4095), v144+v171<<(uint(v172)%32)+int32(base.Ui32(v183)>>(uint(v162)%32)), int32(base.Ui32(v183)>>(uint(int32(1))%32))&int32(2047), v191)
	mBase = m.M
	if v192 == v191 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v208 = v201
	v209 = v182
	v210 = v202
	goto L33
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(1)
	v208 = v155
	v209 = v182
	v210 = v182
	goto L33
L40:
	;
	goto L41
L41:
	;
	v200 = base.B2i32(int32(0) < v192)
	if int32(0) < v192 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v201 = v182 + int32(4)
	goto L44
L43:
	;
	v201 = v155
	goto L44
L44:
	;
	if int32(0) < v192 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v202 = v157
	goto L47
L46:
	;
	v202 = v182
	goto L47
L47:
	;
	if base.Ui32(v201) < base.Ui32(v202) {
		v155 = v201
		v157 = v202
		goto L37
	} else {
		goto L48
	}
L48:
	;
	goto L38
L49:
	;
	v282 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v282 < v283 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(0)
	if base.Ui32(v208) < base.Ui32(v210) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v220 = v209
	goto L53
L52:
	;
	v220 = v210
	goto L53
L53:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v144+v221<<(uint(int32(2))%32)) <= base.Ui32(v220) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v232 = v220
	v233 = v221
	goto L55
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v239 = int32(12)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v255 = int32(1)
	v260 = F_tsCompareString(m, v104+int32(8)+v238*v239+int32(base.Ui32(v242)>>(uint(v239)%32)), v242&int32(4095), v144+v233<<(uint(int32(2))%32)+int32(base.Ui32(v251)>>(uint(v239)%32)), int32(base.Ui32(v251)>>(uint(v255)%32))&int32(2047), v255)
	mBase = m.M
	if v260 != 0 {
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L49
L57:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v261 + int32(1)
	v266 = v232 + int32(4)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v266) < base.Ui32(v144+v267<<(uint(int32(2))%32)) {
		v232 = v266
		v233 = v267
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v286 = v210
	goto L61
L60:
	;
	v286 = v282
	goto L61
L61:
	;
	goto L32
L62:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v289 <= int32(0) {
		v494 = v106
		v496 = v108
		v498 = v110
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v292 = v289
	v294 = v286
	v296 = v106
	v298 = v108
	v300 = v110
	goto L64
L64:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	if v316&int32(1) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v494 = v379
	v496 = v466
	v498 = v383
	goto L30
L66:
	;
	if v335 != 0 {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v348 = v296
	v352 = v300
	goto L73
L68:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v323 = int32(1)
	v334 = v99 + v319<<(uint(int32(2))%32) + (int32(base.Ui32(v316)>>(uint(v323)%32))&int32(2047)+int32(base.Ui32(v316)>>(uint(int32(12))%32))+v323)&int32(4194302)
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334))))
	v336 = v298 + v335
	if v296 <= v336 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v339 = v294 + int32(4)
	if (v339-v286)>>(uint(int32(2))%32) < v292 {
		v294 = v339
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v379 = v296
	v383 = v300
	goto L66
L72:
	;
	v494 = v296
	v496 = v298
	v498 = v300
	goto L30
L73:
	;
	v370 = F_repalloc(m, v352, v348*int32(24))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L20
	} else {
		goto L75
	}
L74:
	;
	v379 = v373
	v383 = v370
	goto L66
L75:
	;
	v373 = v348 << (uint(int32(1)) % 32)
	if v373 <= v336 {
		v348 = v373
		v352 = v370
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v400 = v334 + int32(2)
	v402 = int32(0)
	v408 = v298
	goto L80
L78:
	;
	v460 = v292
	v466 = v298
	goto L79
L79:
	;
	v485 = v294 + int32(4)
	if (v485-v286)>>(uint(int32(2))%32) < v460 {
		v292 = v460
		v294 = v485
		v296 = v379
		v298 = v466
		v300 = v383
		goto L64
	} else {
		goto L89
	}
L80:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	if v426 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	v460 = v459
	v466 = v453
	goto L79
L82:
	;
	v457 = v402 + int32(1)
	if v457 != v335 {
		v402 = v457
		v408 = v453
		goto L80
	} else {
		goto L88
	}
L83:
	;
	v447 = v383 + v408*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v447)+4)) = v294
	*(*uint16)(unsafe.Add(mBase, uint32(v447)+8)) = uint16(v444)
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = v128
	v453 = v408 + int32(1)
	goto L82
L84:
	;
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v402<<(uint(int32(1))%32)))))
	v444 = v432
	goto L83
L85:
	;
	goto L86
L86:
	;
	v433 = int32(1)
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v402<<(uint(v433)%32)))))
	if int32(base.Ui32(v426)>>(uint(int32(base.Ui32(v436)>>(uint(int32(14))%32)))%32))&v433 == int32(0) {
		v453 = v408
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v444 = v436
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
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v522 = v521
	v530 = v498
	goto L27
L92:
	;
	F_pfree(m, v522)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	v1658 = float32(0)
	goto L23
L94:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v498)+8)))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v559 = F_palloc(m, v556<<(uint(int32(2))%32))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v561
	if v496 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v671 = int32(0)
	v673 = float64(0)
	v687 = v671
	v689 = v671
	v693 = v673
	v694 = float64(0)
	v695 = v673
	goto L108
L97:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+8)) = uint16(v555)
	v566 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+4)) = uint16(v566)
	*(*int32)(unsafe.Add(mBase, uint32(v498))) = v559
	v670 = v566
	goto L96
L98:
	;
	goto L99
L99:
	;
	__phi573 = v498
	__phi575 = v498
	__phi577 = v498 + int32(12)
	__phi580 = int32(1)
	__phi582 = v559
	__phi583 = v555
	v573 = __phi573
	v575 = __phi575
	v577 = __phi577
	v580 = __phi580
	v582 = __phi582
	v583 = __phi583
	goto L100
L100:
	;
	v598 = v575 + int32(20)
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598))))
	v600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+8)))
	if v599 != v600 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v628)+8)) = uint16(v630)
	*(*uint16)(unsafe.Add(mBase, uint32(v628)+4)) = uint16(v631)
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v629
	v642 = int32(12)
	v645 = base.I32_div_s(v628-v498+v642, v642)
	v670 = v645
	goto L96
L102:
	;
	v632 = int32(12)
	v633 = v577 + v632
	v636 = base.I32_div_s(v633-v498, v632)
	if v636 < v496 {
		__phi573 = v628
		__phi575 = v577
		__phi577 = v633
		__phi580 = v631
		__phi582 = v629
		__phi583 = v630
		v573 = __phi573
		v575 = __phi575
		v577 = __phi577
		v580 = __phi580
		v582 = __phi582
		v583 = __phi583
		goto L100
	} else {
		goto L107
	}
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+8)) = uint16(v583)
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+4)) = uint16(v580)
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v582
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598))))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	v621 = F_palloc(m, v618<<(uint(int32(2))%32))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L20
	} else {
		goto L106
	}
L104:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v575)+16))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v602 != v603 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v582+base.I32_extend16_s(v580)<<(uint(int32(2))%32)))) = v609
	v628 = v573
	v629 = v582
	v630 = v583
	v631 = v580 + int32(1)
	goto L102
L106:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v623
	v628 = v573 + int32(12)
	v629 = v621
	v630 = v616
	v631 = int32(1)
	goto L102
L107:
	;
	goto L101
L108:
	;
	v711 = v687
	goto L110
L109:
	;
	if l3&int32(1) == int32(0) {
		v1486 = v693
		goto L202
	} else {
		goto L203
	}
L110:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L20
	} else {
		goto L112
	}
L111:
	;
	goto L109
L112:
	;
	v725 = int32(0)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v726)+4))
	if v725 < v727 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v732 = v725
	goto L116
L114:
	;
	goto L115
L115:
	;
	v797 = int32(12)
	v798 = v711 * v797
	v800 = base.I32_div_s(v798, v797)
	if v670 <= v800 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v755 = v732 * int32(32776)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v758 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v755+v756))) = uint8(v758)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v760+v755)+1)) = uint8(v758)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v764+v755)+4)) = v758
	v769 = v732 + int32(1)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	if v769 < v771 {
		v732 = v769
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L115
L118:
	;
	goto L117
L119:
	;
	goto L111
L120:
	;
	v804 = v498 + v711*int32(12)
	v805 = v804
	v815 = v798
	goto L121
L121:
	;
	v829 = int32(*(*int16)(unsafe.Add(mBase, uint32(v805)+4)))
	if int32(0) < v829 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L119
L123:
	;
	v835 = int32(0)
	goto L126
L124:
	;
	goto L125
L125:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v964 = F_TS_execute(m, v957+int32(8), v25+int32(-56), int32(0), int32(1534))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L20
	} else {
		goto L145
	}
L126:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857+v835<<(uint(int32(2))%32))))
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	if v862 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L125
L128:
	;
	v930 = v835 + int32(1)
	v931 = int32(*(*int16)(unsafe.Add(mBase, uint32(v805)+4)))
	if v930 < v931 {
		v835 = v930
		goto L126
	} else {
		goto L144
	}
L129:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v871 = base.I32_div_s(v861-v866-int32(8), int32(12))
	v874 = v865 + v871*int32(32776)
	v875 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v875)
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	if v878 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874)+4)) = v922
	goto L128
L131:
	;
	if v877&int32(1) != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805)+8)))
	v892 = v874 + int32(8)
	v895 = int32(1)
	v898 = v877 & v895
	if v898 != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v885 = int32(32766)
	goto L136
L135:
	;
	v885 = int32(0)
	goto L136
L136:
	;
	v887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v874+v885)+8)) = uint16(v887)
	v922 = int32(1)
	goto L130
L137:
	;
	v899 = int32(16384) - v878
	goto L139
L138:
	;
	v899 = v878 - v895
	goto L139
L139:
	;
	v903 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v892+v899<<(uint(int32(1))%32)))))
	if (v890^v903)&int32(16383) == int32(0) {
		goto L128
	} else {
		goto L140
	}
L140:
	;
	if v898 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v911 = int32(16383) - v878
	goto L143
L142:
	;
	v911 = v878
	goto L143
L143:
	;
	v912 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v892+v911<<(uint(v912)%32)))) = uint16(v890)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	v922 = v916 + v912
	goto L130
L144:
	;
	goto L127
L145:
	;
	if v964 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805)+8)))
	v968 = v966 & int32(16383)
	if v968 == int32(0) {
		goto L119
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v1344 = int32(12)
	v1345 = v805 + v1344
	v1346 = v1345 - v498
	v1348 = base.I32_div_s(v1346, v1344)
	if v1348 < v670 {
		v805 = v1345
		v815 = v1346
		goto L121
	} else {
		goto L201
	}
L149:
	;
	v971 = int32(0)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	if v971 < v973 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v978 = v971
	goto L153
L151:
	;
	goto L152
L152:
	;
	v1043 = v498 + v815
	if base.Ui32(v1043) < base.Ui32(v804) {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	v1001 = v978 * int32(32776)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1004 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1001+v1002))) = uint8(v1004)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1008 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1006+v1001)+1)) = uint8(v1008)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1010+v1001)+4)) = v1004
	v1015 = v978 + v1008
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+4))
	if v1015 < v1017 {
		v978 = v1015
		goto L153
	} else {
		goto L155
	}
L154:
	;
	goto L152
L155:
	;
	goto L154
L156:
	;
	v1243 = base.I32_div_s(v1049-v498, int32(12))
	v1244 = float64(0)
	if base.Ui32(v1049) <= base.Ui32(v805) {
		goto L189
	} else {
		goto L190
	}
L157:
	;
	v711 = v711 + int32(1)
	goto L110
L158:
	;
	v1049 = v1043
	goto L159
L159:
	;
	v1069 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1049)+4)))
	if int32(0) < v1069 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+8)))
	v1213 = v1211 & int32(16383)
	if base.Ui32(v1213) <= base.Ui32(v968) {
		goto L156
	} else {
		goto L188
	}
L161:
	;
	v1075 = int32(0)
	goto L164
L162:
	;
	goto L163
L163:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v1204 = F_TS_execute(m, v1197+int32(8), v25+int32(-56), int32(0), int32(1534))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L20
	} else {
		goto L183
	}
L164:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1049)))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1097+v1075<<(uint(int32(2))%32))))
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101))))
	if v1102 != int32(1) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L163
L166:
	;
	v1170 = v1075 + int32(1)
	v1171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1049)+4)))
	if v1170 < v1171 {
		v1075 = v1170
		goto L164
	} else {
		goto L182
	}
L167:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v1111 = base.I32_div_s(v1101-v1106-int32(8), int32(12))
	v1114 = v1105 + v1111*int32(32776)
	v1115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1114))) = uint8(v1115)
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114)+1)))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	if v1118 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1114)+4)) = v1162
	goto L166
L169:
	;
	if v1117&int32(1) != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v1130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+8)))
	v1132 = v1114 + int32(8)
	v1135 = int32(1)
	v1138 = v1117 & v1135
	if v1138 != 0 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1125 = int32(32766)
	goto L174
L173:
	;
	v1125 = int32(0)
	goto L174
L174:
	;
	v1127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1114+v1125)+8)) = uint16(v1127)
	v1162 = int32(1)
	goto L168
L175:
	;
	v1139 = int32(16384) - v1118
	goto L177
L176:
	;
	v1139 = v1118 - v1135
	goto L177
L177:
	;
	v1143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132+v1139<<(uint(int32(1))%32)))))
	if (v1130^v1143)&int32(16383) == int32(0) {
		goto L166
	} else {
		goto L178
	}
L178:
	;
	if v1138 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1151 = int32(16383) - v1118
	goto L181
L180:
	;
	v1151 = v1118
	goto L181
L181:
	;
	v1152 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1132+v1151<<(uint(v1152)%32)))) = uint16(v1130)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	v1162 = v1156 + v1152
	goto L168
L182:
	;
	goto L165
L183:
	;
	if v1204 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1209 = v1049 - int32(12)
	if base.Ui32(v804) <= base.Ui32(v1209) {
		v1049 = v1209
		goto L159
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	goto L160
L187:
	;
	goto L157
L188:
	;
	goto L157
L189:
	;
	v1248 = v1049
	v1265 = v1244
	goto L192
L190:
	;
	v1302 = v1244
	goto L191
L191:
	;
	v1307 = v805 - v1049
	v1309 = base.I32_div_s(v1307, int32(12))
	v1315 = base.I32_div_s(v1307, int32(24))
	v1317 = v968 - (v1309 + v1213)
	if v1317 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v1272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+8)))
	v1278 = *(*float64)(unsafe.Add(mBase, uint32(v25+int32(-48)+int32(base.Ui32(v1272)>>(uint(int32(11))%32))&int32(24))))
	v1279 = base.F64_add(v1265, v1278)
	v1281 = v1248 + int32(12)
	if base.Ui32(v1281) <= base.Ui32(v805) {
		v1248 = v1281
		v1265 = v1279
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v1302 = v1279
	goto L191
L194:
	;
	goto L193
L195:
	;
	v1320 = v1315
	goto L197
L196:
	;
	v1320 = v1317
	goto L197
L197:
	;
	v1328 = base.F64_mul(base.F64_convert_i32_u(v1213+v968), float64(0.5))
	if v689 <= int32(0) {
		v1338 = v695
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1339 = int32(1)
	v687 = v1243 + v1339
	v689 = v689 + v1339
	v693 = base.F64_add(v693, base.F64_div(base.F64_div(base.F64_convert_i32_s(v1309+int32(1)), v1302), base.F64_convert_i32_s(v1320+int32(1))))
	v694 = v1328
	v695 = v1338
	goto L108
L199:
	;
	if base.F64_gt(v1328, v694) == int32(0) {
		v1338 = v695
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1338 = base.F64_add(v695, base.F64_div(float64(1), base.F64_sub(v1328, v694)))
	goto L198
L201:
	;
	goto L122
L202:
	;
	if l3&int32(2) == int32(0) {
		v1575 = v1486
		goto L217
	} else {
		goto L218
	}
L203:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1378 <= int32(0) {
		v1486 = v693
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1384 = v99 + v1378<<(uint(int32(2))%32)
	if base.Ui32(v99) < base.Ui32(v1384) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1389 = v99
	v1391 = int32(0)
	goto L208
L206:
	;
	v1465 = float64(1)
	goto L207
L207:
	;
	v1466 = F_log(m, v1465)
	mBase = m.M
	v1486 = base.F64_div(v693, v1466)
	goto L202
L208:
	;
	v1411 = int32(1)
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1389)))
	if v1412&v1411 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1465 = base.F64_convert_i32_s(v1434 + int32(1))
	goto L207
L210:
	;
	v1415 = int32(1)
	v1428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1384+(int32(base.Ui32(v1412)>>(uint(v1415)%32))&int32(2047)+int32(base.Ui32(v1412)>>(uint(int32(12))%32))+v1415)&int32(4194302)))))
	if base.Ui32(v1428) <= base.Ui32(v1415) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1433 = v1411
	goto L212
L212:
	;
	v1434 = v1433 + v1391
	v1436 = v1389 + int32(4)
	if base.Ui32(v1436) < base.Ui32(v1384) {
		v1389 = v1436
		v1391 = v1434
		goto L208
	} else {
		goto L216
	}
L213:
	;
	v1431 = v1415
	goto L215
L214:
	;
	v1431 = v1428
	goto L215
L215:
	;
	v1433 = v1431
	goto L212
L216:
	;
	goto L209
L217:
	;
	if l3&int32(4) == int32(0) {
		v1594 = v1575
		goto L230
	} else {
		goto L231
	}
L218:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1499 = v99 + v1496<<(uint(int32(2))%32)
	if base.Ui32(v1499) <= base.Ui32(v99) {
		v1575 = v1486
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1502 = int32(0)
	v1513 = v99
	goto L220
L220:
	;
	v1526 = int32(1)
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1513)))
	if v1527&v1526 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v1549 <= int32(0) {
		v1575 = v1486
		goto L217
	} else {
		goto L229
	}
L222:
	;
	v1530 = int32(1)
	v1543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1499+(int32(base.Ui32(v1527)>>(uint(v1530)%32))&int32(2047)+int32(base.Ui32(v1527)>>(uint(int32(12))%32))+v1530)&int32(4194302)))))
	if base.Ui32(v1543) <= base.Ui32(v1530) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v1548 = v1526
	goto L224
L224:
	;
	v1549 = v1548 + v1502
	v1551 = v1513 + int32(4)
	if base.Ui32(v1551) < base.Ui32(v1499) {
		v1502 = v1549
		v1513 = v1551
		goto L220
	} else {
		goto L228
	}
L225:
	;
	v1546 = v1530
	goto L227
L226:
	;
	v1546 = v1543
	goto L227
L227:
	;
	v1548 = v1546
	goto L224
L228:
	;
	goto L221
L229:
	;
	v1575 = base.F64_div(v1486, base.F64_convert_i32_u(v1549))
	goto L217
L230:
	;
	if l3&int32(8) == int32(0) {
		v1605 = v1594
		goto L234
	} else {
		goto L235
	}
L231:
	;
	if v689 <= int32(0) {
		v1594 = v1575
		goto L230
	} else {
		goto L232
	}
L232:
	;
	if base.F64_gt(v695, float64(0)) == int32(0) {
		v1594 = v1575
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v1594 = base.F64_div(v1575, base.F64_div(base.F64_convert_i32_u(v689), v695))
	goto L230
L234:
	;
	if l3&int32(16) == int32(0) {
		v1621 = v1605
		goto L237
	} else {
		goto L238
	}
L235:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1599 <= int32(0) {
		v1605 = v1594
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1605 = base.F64_div(v1594, base.F64_convert_i32_u(v1599))
	goto L234
L237:
	;
	if l3&int32(32) != 0 {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1610 <= int32(0) {
		v1621 = v1605
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v1616 = F_log(m, base.F64_convert_i32_s(v1610+int32(1)))
	mBase = m.M
	v1621 = base.F64_div(v1605, base.F64_div(v1616, float64(0.6931471805599453)))
	goto L237
L240:
	;
	v1627 = base.F64_div(v1621, base.F64_add(v1621, float64(1)))
	goto L242
L241:
	;
	v1627 = v1621
	goto L242
L242:
	;
	F_pfree(m, v498)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	F_pfree(m, v1630)
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	v1658 = base.F32_demote_f64(v1627)
	goto L23
L245:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	F_errmsg(m, int32(398252), int32(0))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L20
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(492630), int32(876), int32(459000))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cancel_prior_stmt_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	v11 = *(*int32)(unsafe.Add(mBase, _consts[301]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[302]))
	v16 = v11 + v13*int32(20)
	v18 = v16 + int32(16)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
	if v85 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v59 = int32(4480304)
	v60 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v63 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v63
	v66 = F_palloc0(m, int32(36))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v30 = int32(0)
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25+v30<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != l0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v48 = v30 + int32(1)
	if v22 != v48 {
		v30 = v48
		goto L5
	} else {
		goto L11
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 != l1 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v44 != int32(1) {
		v80 = v39
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
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = l0
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v71 = F_lappend(m, v70, v66)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v71
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v60
	v80 = v66
	goto L1
L15:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v173)
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v177
	return
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v97 = v94
	v99 = v95
	goto L22
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v94 = v88
	v95 = v89
	goto L17
L19:
	;
	goto L20
L20:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v91 == v90 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v94 = v91
	v95 = v90
	goto L17
L22:
	;
	if v99 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L15
L24:
	;
	v107 = v99
	goto L26
L25:
	;
	v107 = v97 + int32(16)
	goto L26
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if base.Ui32(v107) < base.Ui32(v108) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v113 = v107
	goto L30
L28:
	;
	goto L29
L29:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v163 != 0 {
		v97 = v163
		v99 = int32(0)
		goto L22
	} else {
		goto L42
	}
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v122 = v113 + v119&int32(134217727)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v123 != l0 {
		goto L15
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v125&int32(3) != l2 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	if v125&int32(28) != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v119&int32(1073741823) | int32(-2147483648)
	v138 = v119 & int32(939524096)
	if v138 == int32(134217728) {
		v149 = int32(24)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = v149 + v113
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if base.Ui32(v150) < base.Ui32(v151) {
		v113 = v150
		goto L30
	} else {
		goto L41
	}
L36:
	;
	if v138 == int32(268435456) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v145 = int32(12)
	goto L39
L38:
	;
	v145 = int32(4)
	goto L39
L39:
	;
	if v138 != int32(805306368) {
		v149 = v145
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(16)
	goto L35
L41:
	;
	goto L31
L42:
	;
	goto L23
}
func F_canonicalize_ec_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
					v35 = v7
					v38 = F_exprCollation(m, l0)
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v38 == l2 {
							v52 = l0
							return v52
						} else {
							v41 = F_exprTypmod(m, l0)
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v35
								v45 = v41
								v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
						v43 = l1
						v45 = int32(-1)
						v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = v50
							return v52
						}
					} else {
						v35 = l1
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
							v43 = l1
							v45 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v35 = l1
							v38 = F_exprCollation(m, l0)
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == l2 {
									v52 = l0
									return v52
								} else {
									v41 = F_exprTypmod(m, l0)
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = v35
										v45 = v41
										v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
						v35 = v7
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
					v35 = v7
					v38 = F_exprCollation(m, l0)
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v38 == l2 {
							v52 = l0
							return v52
						} else {
							v41 = F_exprTypmod(m, l0)
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v35
								v45 = v41
								v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
							v43 = l1
							v45 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v35 = l1
							v38 = F_exprCollation(m, l0)
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == l2 {
									v52 = l0
									return v52
								} else {
									v41 = F_exprTypmod(m, l0)
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = v35
										v45 = v41
										v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
						v35 = v7
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
			if base.Ui32(l1-int32(5077)) < base.Ui32(int32(4)) {
				v35 = v7
				v38 = F_exprCollation(m, l0)
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v38 == l2 {
						v52 = l0
						return v52
					} else {
						v41 = F_exprTypmod(m, l0)
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = v35
							v45 = v41
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
				if base.Ui32(l1-int32(4537)) < base.Ui32(int32(2)) {
					v35 = v7
					v38 = F_exprCollation(m, l0)
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v38 == l2 {
							v52 = l0
							return v52
						} else {
							v41 = F_exprTypmod(m, l0)
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v35
								v45 = v41
								v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
					if l1 == int32(3831) {
						v35 = v7
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
							v43 = l1
							v45 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v35 = l1
							v38 = F_exprCollation(m, l0)
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == l2 {
									v52 = l0
									return v52
								} else {
									v41 = F_exprTypmod(m, l0)
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = v35
										v45 = v41
										v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_consts[1137])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1138]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_consts[1138]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1139]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_consts[1140])))
	v52 = v50 >> (uint(v26) % 32)
	v54 = v50 & v12
	if base.Ui32(v54) <= base.Ui32(int32(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v52&(int32(0)-(l1^v54)) + l0
L4:
	;
	goto L5
L5:
	;
	v64 = v52 & int32(255)
	if v64 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v71 = v64
	v72 = int32(base.Ui32(v52) >> (uint(int32(8)) % 32))
	goto L7
L7:
	;
	v78 = int32(1)
	v79 = int32(base.Ui32(v71) >> (uint(v78) % 32))
	v80 = v79 + v72
	v82 = v80 << (uint(v78) % 32)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1141]))))
	if v85 == v13 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1142]))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_consts[1140])))
	v94 = v92 & int32(255)
	if base.Ui32(v94) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v110 = base.B2i32(base.Ui32(v13) < base.Ui32(v85))
	if base.Ui32(v13) < base.Ui32(v85) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	return (int32(0)-(l1^v94))&(v92>>(uint(int32(8))%32)) + l0
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
	v107 = int32(-1)
	goto L17
L16:
	;
	v107 = int32(1)
	goto L17
L17:
	;
	return v107 + l0
L18:
	;
	v111 = v72
	goto L20
L19:
	;
	v111 = v80
	goto L20
L20:
	;
	if base.Ui32(v13) < base.Ui32(v85) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = v79
	goto L23
L22:
	;
	v113 = v71 - v79
	goto L23
L23:
	;
	if v113 != 0 {
		v71 = v113
		v72 = v111
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v942 int32
	_ = v942
	v5 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
		v942 = v27
		goto L10
	case 2:
		goto L15
	case 6:
		goto L17
	case 21:
		v907 = v5
		goto L11
	default:
		goto L18
	}
L10:
	;
	return v942
L11:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v915 <= int32(0) {
		goto L307
	} else {
		goto L308
	}
L12:
	;
	F_pfree(m, v416)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L4
	} else {
		goto L306
	}
L13:
	;
	v634 = (l3 - l2) >> (uint(int32(2)) % 32)
	v635 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v634) < base.Ui32(v635) {
		goto L230
	} else {
		goto L231
	}
L14:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v629 = F_cdissect(m, l0, v628, l2, l3)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L4
	} else {
		goto L229
	}
L15:
	;
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	if v390&int32(2) == int32(0) {
		goto L13
	} else {
		goto L151
	}
L16:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v323 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v101 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v100<<(uint(v101)%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+24))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v106&v101 != 0 {
		goto L37
	} else {
		goto L38
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
		v942 = v27
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
		v907 = v35
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
	v907 = base.B2i32(l2 != l3) | base.B2i32(base.I32_extend16_s(v44) < base.I32_extend16_s(v45))
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
	v907 = base.B2i32(v45 != int32(0))
	goto L11
L26:
	;
	goto L27
L27:
	;
	v58 = (l3 - l2) >> (uint(int32(2)) % 32)
	v59 = v46 - v41
	v60 = base.I32_div_u_s(v58, v59)
	if v58-v60*v59 != 0 {
		v907 = v35
		goto L11
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v60) < base.Ui32(v45) {
		v907 = v35
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(v44 != int32(256))&base.B2i32(base.Ui32(v44) < base.Ui32(v60)) != 0 {
		v907 = v35
		goto L11
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v58) < base.Ui32(v59) {
		v907 = int32(0)
		goto L11
	} else {
		goto L31
	}
L31:
	;
	v71 = int32(2)
	v80 = l2
	v82 = v60
	goto L32
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+420))
	v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v64+v41<<(uint(v71)%32), v80, v59)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v907 = v94
	goto L11
L34:
	;
	v94 = base.B2i32(v91 != int32(0))
	if v91 != 0 {
		v907 = v94
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v97 = v82 - int32(1)
	if v97 != 0 {
		v80 = v80 + v59<<(uint(v71)%32)
		v82 = v97
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	if v104 != 0 {
		v136 = v104
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v104 != 0 {
		v244 = v104
		goto L85
	} else {
		goto L86
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v137 != 0 {
		v942 = v137
		goto L10
	} else {
		goto L47
	}
L41:
	;
	v109 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v116 = F_newdfa(m, l0, v99+int32(36), v112+int32(72), v109)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v116 == int32(0) {
		v136 = v109
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v120 == int32(98) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = v123
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+64)) = uint16(v125)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+66)) = uint16(v127)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v130<<(uint(int32(2))%32)))) = v116
	v136 = v116
	goto L40
L47:
	;
	v138 = F_getsubdfa(m, l0, v105)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v140 != 0 {
		v942 = v140
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v141 = int32(0)
	v143 = F_shortest(m, l0, v136, l2, l2, l3, v141, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v145 != 0 {
		v942 = v145
		goto L10
	} else {
		goto L51
	}
L51:
	;
	if v143 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(1)
L53:
	;
	goto L54
L54:
	;
	v157 = v143
	goto L55
L55:
	;
	v164 = F_longest(m, l0, v138, v157, l3, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L57
	}
L56:
	;
	v942 = int32(1)
	goto L10
L57:
	;
	if v164 == l3 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v167 = F_cdissect(m, l0, v99, l2, v157)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v206 != 0 {
		v942 = v206
		goto L10
	} else {
		goto L80
	}
L61:
	;
	if v167 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v171 = F_cdissect(m, l0, v105, v157, l3)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	v202 = v167
	goto L64
L64:
	;
	if v202 != int32(1) {
		v942 = v202
		goto L10
	} else {
		goto L79
	}
L65:
	;
	if v171 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v907 = int32(0)
	goto L11
L67:
	;
	goto L68
L68:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v177 <= int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v202 = v171
	goto L64
L70:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v193 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v180) <= base.Ui32(v177) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v183 = v177 << (uint(int32(3)) % 32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v186 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v188+v183)+4)) = v186
	goto L70
L73:
	;
	v195 = v193
	goto L76
L74:
	;
	goto L75
L75:
	;
	goto L69
L76:
	;
	F_zaptreesubs(m, l0, v195)
	mBase = m.M
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v195)+24))
	if v198 != 0 {
		v195 = v198
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	goto L77
L79:
	;
	goto L60
L80:
	;
	if l3 == v157 {
		v942 = int32(1)
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v211 = int32(0)
	v213 = F_shortest(m, l0, v136, l2, v157+int32(4), l3, v211, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v215 != 0 {
		v942 = v215
		goto L10
	} else {
		goto L83
	}
L83:
	;
	if v213 != 0 {
		v157 = v213
		goto L55
	} else {
		goto L84
	}
L84:
	;
	goto L56
L85:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v245 != 0 {
		v942 = v245
		goto L10
	} else {
		goto L92
	}
L86:
	;
	v217 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = F_newdfa(m, l0, v99+int32(36), v220+int32(72), v217)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	if v224 == int32(0) {
		v244 = v217
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v228 == int32(98) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+60)) = v231
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+64)) = uint16(v233)
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+66)) = uint16(v235)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v237+v238<<(uint(int32(2))%32)))) = v224
	v244 = v224
	goto L85
L92:
	;
	v246 = F_getsubdfa(m, l0, v105)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v248 != 0 {
		v942 = v248
		goto L10
	} else {
		goto L94
	}
L94:
	;
	v250 = F_longest(m, l0, v244, l2, l3, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v252 != 0 {
		v942 = v252
		goto L10
	} else {
		goto L96
	}
L96:
	;
	if v250 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	return int32(1)
L98:
	;
	goto L99
L99:
	;
	v264 = v250
	goto L100
L100:
	;
	v271 = F_longest(m, l0, v246, v264, l3, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	v942 = int32(1)
	goto L10
L102:
	;
	if v271 == l3 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v274 = F_cdissect(m, l0, v99, l2, v264)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v313 != 0 {
		v942 = v313
		goto L10
	} else {
		goto L125
	}
L106:
	;
	if v274 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v278 = F_cdissect(m, l0, v105, v264, l3)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	v309 = v274
	goto L109
L109:
	;
	if v309 != int32(1) {
		v942 = v309
		goto L10
	} else {
		goto L124
	}
L110:
	;
	if v278 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v907 = int32(0)
	goto L11
L112:
	;
	goto L113
L113:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v284 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v309 = v278
	goto L109
L115:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v300 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v287) <= base.Ui32(v284) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v290 = v284 << (uint(int32(3)) % 32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v293 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v290+v291))) = v293
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v295+v290)+4)) = v293
	goto L115
L118:
	;
	v302 = v300
	goto L121
L119:
	;
	goto L120
L120:
	;
	goto L114
L121:
	;
	F_zaptreesubs(m, l0, v302)
	mBase = m.M
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v302)+24))
	if v305 != 0 {
		v302 = v305
		goto L121
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	goto L122
L124:
	;
	goto L105
L125:
	;
	if l2 == v264 {
		v942 = int32(1)
		goto L10
	} else {
		goto L126
	}
L126:
	;
	v319 = F_longest(m, l0, v244, l2, v264-int32(4), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v321 != 0 {
		v942 = v321
		goto L10
	} else {
		goto L128
	}
L128:
	;
	if v319 != 0 {
		v264 = v319
		goto L100
	} else {
		goto L129
	}
L129:
	;
	goto L101
L130:
	;
	return int32(1)
L131:
	;
	goto L132
L132:
	;
	v334 = v323
	goto L133
L133:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341+v342<<(uint(int32(2))%32))))
	if v346 != 0 {
		v374 = v346
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v942 = int32(1)
	goto L10
L135:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v375 != 0 {
		v942 = v375
		goto L10
	} else {
		goto L142
	}
L136:
	;
	v347 = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v354 = F_newdfa(m, l0, v334+int32(36), v350+int32(72), v347)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v354 == int32(0) {
		v374 = v347
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v358 == int32(98) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+60)) = v361
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v354)+64)) = uint16(v363)
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v354)+66)) = uint16(v365)
	goto L141
L140:
	;
	goto L141
L141:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v367+v368<<(uint(int32(2))%32)))) = v354
	v374 = v354
	goto L135
L142:
	;
	v377 = F_longest(m, l0, v374, l2, l3, int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	if v377 == l3 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v380 = F_cdissect(m, l0, v334, l2, l3)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v385 != 0 {
		v942 = v385
		goto L10
	} else {
		goto L149
	}
L147:
	;
	if v380 != int32(1) {
		v907 = v380
		goto L11
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v334)+24))
	if v387 != 0 {
		v334 = v387
		goto L133
	} else {
		goto L150
	}
L150:
	;
	goto L134
L151:
	;
	if v388 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if l2 == l3 {
		v907 = v5
		goto L11
	} else {
		goto L155
	}
L153:
	;
	v399 = v388
	goto L154
L154:
	;
	v402 = (l3 - l2) >> (uint(int32(2)) % 32)
	v403 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v402) < base.Ui32(v403) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v399 = int32(1)
	goto L154
L156:
	;
	v405 = v402
	goto L158
L157:
	;
	v405 = v403
	goto L158
L158:
	;
	if v403 == int32(256) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v408 = v402
	goto L161
L160:
	;
	v408 = v405
	goto L161
L161:
	;
	if base.Ui32(v399) < base.Ui32(v408) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v410 = v408
	goto L164
L163:
	;
	v410 = v399
	goto L164
L164:
	;
	v411 = int32(2)
	v416 = F_palloc_extended(m, v410<<(uint(v411)%32)+int32(4), v411)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	if v416 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	return int32(12)
L167:
	;
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = l2
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v424 = F_getsubdfa(m, l0, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v426 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	v432 = int32(1)
	v433 = l2
	v435 = v5
	goto L171
L171:
	;
	v442 = v432 - int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v416+v442<<(uint(int32(2))%32))))
	if l3 == v433 {
		v457 = v433
		goto L173
	} else {
		goto L174
	}
L172:
	;
	F_pfree(m, v416)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L228
	}
L173:
	;
	if base.Ui32(v432) < base.Ui32(v410) {
		goto L180
	} else {
		goto L181
	}
L174:
	;
	if v433 != v446 {
		v457 = v433
		goto L173
	} else {
		goto L175
	}
L175:
	;
	if base.Ui32(v432) < base.Ui32(v399) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	if (l3-v433)>>(uint(int32(2))%32) <= v399-v432 {
		v457 = v433
		goto L173
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v457 = v433 + int32(4)
	goto L173
L179:
	;
	goto L178
L180:
	;
	v462 = v457
	goto L182
L181:
	;
	v462 = l3
	goto L182
L182:
	;
	v463 = int32(0)
	v465 = F_shortest(m, l0, v424, v446, v462, l3, v463, v463)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416+v432<<(uint(int32(2))%32)))) = v465
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v468 != 0 {
		goto L12
	} else {
		goto L184
	}
L184:
	;
	if v465 == int32(0) {
		v542 = v435
		goto L189
	} else {
		goto L190
	}
L185:
	;
	goto L172
L186:
	;
	if int32(0) < v600 {
		v432 = v600
		v433 = v601
		v435 = v603
		goto L171
	} else {
		goto L227
	}
L187:
	;
	v576 = v563
	goto L221
L188:
	;
	if v556 <= int32(0) {
		goto L185
	} else {
		goto L220
	}
L189:
	;
	v552 = v542
	v556 = v442
	goto L188
L190:
	;
	if v435 < v432 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v472 = v435
	goto L193
L192:
	;
	v472 = v442
	goto L193
L193:
	;
	if l3 != v465 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if base.Ui32(v410) <= base.Ui32(v432) {
		v542 = v472
		goto L189
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if base.Ui32(v432) < base.Ui32(v399) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v600 = v432 + int32(1)
	v601 = v465
	v603 = v472
	goto L186
L198:
	;
	v563 = v432
	v568 = v472
	goto L187
L199:
	;
	goto L200
L200:
	;
	v485 = v472
	goto L202
L201:
	;
	F_pfree(m, v416)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L219
	}
L202:
	;
	v492 = v485 + int32(1)
	if v432 < v492 {
		goto L201
	} else {
		goto L204
	}
L203:
	;
	if v530 == int32(1) {
		v552 = v485
		v556 = v492
		goto L188
	} else {
		goto L217
	}
L204:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if v496 <= int32(0) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v522 = int32(2)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v416+v485<<(uint(v522)%32))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v416+v492<<(uint(v522)%32))))
	v530 = F_cdissect(m, l0, v521, v525, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L215
	}
L206:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	if v512 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v499) <= base.Ui32(v496) {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v502 = v496 << (uint(int32(3)) % 32)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v505 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v502+v503))) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v507+v502)+4)) = v505
	goto L206
L209:
	;
	v514 = v512
	goto L212
L210:
	;
	goto L211
L211:
	;
	goto L205
L212:
	;
	F_zaptreesubs(m, l0, v514)
	mBase = m.M
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	if v517 != 0 {
		v514 = v517
		goto L212
	} else {
		goto L214
	}
L213:
	;
	goto L211
L214:
	;
	goto L213
L215:
	;
	if v530 == int32(0) {
		v485 = v492
		goto L202
	} else {
		goto L216
	}
L216:
	;
	goto L203
L217:
	;
	F_pfree(m, v416)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	return v530
L219:
	;
	v907 = int32(0)
	goto L11
L220:
	;
	v563 = v556
	v568 = v552
	goto L187
L221:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v416+v576<<(uint(int32(2))%32))))
	if base.Ui32(v588) < base.Ui32(l3) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L185
L223:
	;
	v600 = v576
	v601 = v588 + int32(4)
	v603 = v568
	goto L186
L224:
	;
	goto L225
L225:
	;
	v592 = int32(1)
	if v592 < v576 {
		v576 = v576 - v592
		goto L221
	} else {
		goto L226
	}
L226:
	;
	goto L222
L227:
	;
	goto L185
L228:
	;
	return int32(1)
L229:
	;
	v907 = v629
	goto L11
L230:
	;
	v637 = v634
	goto L232
L231:
	;
	v637 = v635
	goto L232
L232:
	;
	if v635 == int32(256) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v640 = v634
	goto L235
L234:
	;
	v640 = v637
	goto L235
L235:
	;
	v641 = int32(1)
	if v388 <= v641 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v644 = v641
	goto L238
L237:
	;
	v644 = v388
	goto L238
L238:
	;
	if base.Ui32(v644) < base.Ui32(v640) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v646 = v640
	goto L241
L240:
	;
	v646 = v644
	goto L241
L241:
	;
	v647 = int32(2)
	v652 = F_palloc_extended(m, v646<<(uint(v647)%32)+int32(4), v647)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	if v652 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	return int32(12)
L244:
	;
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652))) = l2
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v660 = F_getsubdfa(m, l0, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v662 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	F_pfree(m, v652)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L4
	} else {
		goto L305
	}
L248:
	;
	v667 = int32(1)
	v671 = l3
	v672 = v5
	goto L249
L249:
	;
	v676 = int32(2)
	v680 = v667 - int32(1)
	v683 = v652 + v680<<(uint(v676)%32)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	v686 = F_longest(m, l0, v660, v684, v671, int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L251
	}
L250:
	;
	F_pfree(m, v652)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L4
	} else {
		goto L304
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652+v667<<(uint(v676)%32)))) = v686
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v689 != 0 {
		goto L247
	} else {
		goto L252
	}
L252:
	;
	if v686 == int32(0) {
		v771 = v672
		goto L257
	} else {
		goto L258
	}
L253:
	;
	goto L250
L254:
	;
	if int32(0) < v839 {
		v667 = v839
		v671 = v843
		v672 = v844
		goto L249
	} else {
		goto L303
	}
L255:
	;
	v805 = v792
	goto L293
L256:
	;
	if v777 <= int32(0) {
		goto L253
	} else {
		goto L292
	}
L257:
	;
	v777 = v680
	v782 = v771
	goto L256
L258:
	;
	if v672 < v667 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v693 = v672
	goto L261
L260:
	;
	v693 = v680
	goto L261
L261:
	;
	if l3 != v686 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v792 = v667
	v797 = v693
	goto L255
L263:
	;
	if base.Ui32(v646) <= base.Ui32(v667) {
		v771 = v693
		goto L257
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	if base.Ui32(v667) < base.Ui32(v644) {
		goto L262
	} else {
		goto L272
	}
L266:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	if v696 == v686 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	if base.Ui32(v644) <= base.Ui32(v667) {
		goto L262
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v839 = v667 + int32(1)
	v843 = l3
	v844 = v693
	goto L254
L270:
	;
	if v644-v667 < (l3-v686)>>(uint(int32(2))%32) {
		goto L262
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	v716 = v693
	goto L274
L273:
	;
	F_pfree(m, v652)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L291
	}
L274:
	;
	v721 = v716 + int32(1)
	if v667 < v721 {
		goto L273
	} else {
		goto L276
	}
L275:
	;
	if v759 == int32(1) {
		v777 = v721
		v782 = v716
		goto L256
	} else {
		goto L289
	}
L276:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	if v725 <= int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v751 = int32(2)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v652+v716<<(uint(v751)%32))))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v652+v721<<(uint(v751)%32))))
	v759 = F_cdissect(m, l0, v750, v754, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L4
	} else {
		goto L287
	}
L278:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	if v741 != 0 {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v728) <= base.Ui32(v725) {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v731 = v725 << (uint(int32(3)) % 32)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v734 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v731+v732))) = v734
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v736+v731)+4)) = v734
	goto L278
L281:
	;
	v743 = v741
	goto L284
L282:
	;
	goto L283
L283:
	;
	goto L277
L284:
	;
	F_zaptreesubs(m, l0, v743)
	mBase = m.M
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v743)+24))
	if v746 != 0 {
		v743 = v746
		goto L284
	} else {
		goto L286
	}
L285:
	;
	goto L283
L286:
	;
	goto L285
L287:
	;
	if v759 == int32(0) {
		v716 = v721
		goto L274
	} else {
		goto L288
	}
L288:
	;
	goto L275
L289:
	;
	F_pfree(m, v652)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	return v759
L291:
	;
	v907 = int32(0)
	goto L11
L292:
	;
	v792 = v777
	v797 = v782
	goto L255
L293:
	;
	v815 = v652 + v805<<(uint(int32(2))%32)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815-int32(4))))
	if base.Ui32(v816) <= base.Ui32(v819) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	goto L253
L295:
	;
	v831 = int32(1)
	if v831 < v805 {
		v805 = v805 - v831
		goto L293
	} else {
		goto L302
	}
L296:
	;
	v822 = v816 - int32(4)
	if base.Ui32(v819) < base.Ui32(v822) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v839 = v805
	v843 = v822
	v844 = v797
	goto L254
L298:
	;
	goto L299
L299:
	;
	if base.Ui32(v644) <= base.Ui32(v805) {
		goto L295
	} else {
		goto L300
	}
L300:
	;
	if v644-v805 < (l3-v819)>>(uint(int32(2))%32) {
		goto L295
	} else {
		goto L301
	}
L301:
	;
	v839 = v805
	v843 = v822
	v844 = v797
	goto L254
L302:
	;
	goto L294
L303:
	;
	goto L253
L304:
	;
	v866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v907 = base.B2i32(l2 != l3) | base.B2i32(v866 != int32(0))
	goto L11
L305:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v907 = v885
	goto L11
L306:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v907 = v901
	goto L11
L307:
	;
	return v907
L308:
	;
	goto L309
L309:
	;
	if v907 != 0 {
		v942 = v907
		goto L10
	} else {
		goto L310
	}
L310:
	;
	v919 = int32(0)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v920) <= base.Ui32(v915) {
		v942 = v919
		goto L10
	} else {
		goto L311
	}
L311:
	;
	v923 = v915 << (uint(int32(3)) % 32)
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v928 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v923+v924))) = (l2 - v926) >> (uint(v928) % 32)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v931+v923)+4)) = (l3 - v933) >> (uint(v928) % 32)
	v942 = v919
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(5))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = base.I32_extend8_s(v5)
		if v11 < int32(0) {
			v14 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v14)
			v16 = int32(7)
			v18 = int32(48)
			v19 = v5&v16 | v18
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v19)
			v26 = int32(base.Ui32(v5)>>(uint(int32(3))%32))&v16 | v18
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v26)
			v35 = int32(base.Ui32(v5&int32(192))>>(uint(int32(6))%32)) | v18
			v36 = int32(92)
		} else {
			v35 = int32(0)
			v36 = v11
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v36)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v35)
		return v7
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
		v16 = *(*int32)(unsafe.Add(mBase, _consts[4]))
		v17 = v16
	} else {
		v17 = l1
	}
	if base.Ui32(l0) < base.Ui32(int32(16384)) {
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
											v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[122])))
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
													v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[965])))
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
																	F_errmsg(m, int32(696839), v11)
																	mBase = m.M
																	v69 = m.ExcPending
																	if v69 != 0 {
																		return int32(0)
																	} else {
																		if v40 != 0 {
																			F_errhint(m, int32(635203), int32(0))
																			mBase = m.M
																			v73 = m.ExcPending
																			if v73 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(488901), int32(129), int32(150423))
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
																			F_errfinish(m, int32(488901), int32(129), int32(150423))
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
											v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[965])))
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
															F_errmsg(m, int32(696839), v11)
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return int32(0)
															} else {
																if v40 != 0 {
																	F_errhint(m, int32(635203), int32(0))
																	mBase = m.M
																	v73 = m.ExcPending
																	if v73 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(488901), int32(129), int32(150423))
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
																	F_errfinish(m, int32(488901), int32(129), int32(150423))
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[166]))
		*(*int32)(unsafe.Add(mBase, _consts[189])) = v8
		v14 = F_format_elog_string(m, int32(625790), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[555])) = v14
			return base.B2i32(v4 != int32(0))
		}
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_check_labels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(544475))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L22
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(544475))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L17
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
	m.G0 = v9 + int32(32)
	return
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		v35 = v15
		v36 = v16
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v36-v35 != 0 {
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v15 != v16 {
		v35 = v15
		v36 = v16
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v20 = l0
	v21 = l1
	goto L11
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v24
		v36 = v25
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v35 = v24
	v36 = v25
	goto L8
L13:
	;
	v28 = int32(1)
	if v24 == v25 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L5
L16:
	;
	return
L17:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg(m, int32(313776), v9)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v55 = F_plpgsql_scanner_errposition(m, l2, l3)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(26863), int32(3886), int32(151813))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errmsg(m, int32(691675), v9+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v81 = F_plpgsql_scanner_errposition(m, l2, l3)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(26863), int32(3893), int32(151813))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
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
	v5 = F_check_slru_buffers(m, int32(133786), l0)
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
	F_appendStringInfoString(m, v11, int32(665562))
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
	v27 = *(*int32)(unsafe.Add(mBase, _consts[293]))
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
	v40 = F_MakeSingleTupleTableSlot(m, v38, int32(1597084))
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
	F_errmsg_plural(m, int32(221399), int32(221446), v110, v7)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(489337), int32(501), int32(142270))
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
	F_errmsg(m, int32(200184), v7+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(489337), int32(467), int32(142270))
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
						F_errmsg(m, int32(435822), v5+int32(16))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(578849)
							F_errdetail_internal(m, int32(204631), v5)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(492525), int32(5709), int32(377214))
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
	var v14 int32
	_ = v14
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v148 int32
	_ = v148
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
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
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
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
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
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(208)
	return v406
L2:
	;
	v65 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[368]))
	if v67 == v65 {
		v365 = v4
		v374 = v65
		goto L22
	} else {
		goto L23
	}
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v22 == int32(0) {
		v41 = v21
		v42 = v22
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	if v42-v41 == int32(0) {
		v406 = v4
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v21 != v22 {
		v41 = v21
		v42 = v22
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v26 = l1
	v27 = l2
	goto L11
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v30
		v42 = v31
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v41 = v30
	v42 = v31
	goto L8
L13:
	;
	v34 = int32(1)
	if v30 == v31 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v46 = int32(-1)
	v49 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	if v49 == int32(0) {
		v406 = v46
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg(m, int32(322551), v16)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(495117), int32(2988), int32(236204))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v406 = v46
	goto L1
L21:
	;
	v406 = int32(0) - (v365 ^ int32(1))
	goto L1
L22:
	;
	if v374 != 0 {
		goto L21
	} else {
		goto L107
	}
L23:
	;
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v71 <= v70 {
		v365 = v4
		v374 = v70
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v82 = v4
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v82<<(uint(int32(2))%32))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v96 == int32(0) {
		v115 = v95
		v116 = v96
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v365 = v349
	v374 = int32(0)
	goto L22
L27:
	;
	v357 = v82 + int32(1)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v357 < v358 {
		v82 = v357
		goto L25
	} else {
		goto L106
	}
L28:
	;
	v349 = int32(0)
	goto L27
L29:
	;
	if v116-v115 != 0 {
		goto L28
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v95 != v96 {
		v115 = v95
		v116 = v96
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v100 = v92
	v101 = l0
	goto L33
L33:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v104
		v116 = v105
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v115 = v104
	v116 = v105
	goto L30
L35:
	;
	v108 = int32(1)
	if v104 == v105 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v119 = F_get_role_oid(m, l1, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v122 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v333|v335 == int32(0) {
		v349 = v333
		goto L27
	} else {
		goto L105
	}
L40:
	;
	v123 = F_strlen(m, l2)
	mBase = m.M
	v128 = F_palloc(m, v123<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v300 == int32(0) {
		v319 = v299
		v320 = v300
		goto L95
	} else {
		goto L96
	}
L43:
	;
	v130 = F_strlen(m, l2)
	mBase = m.M
	v131 = F_pg_mb2wchar_with_len(m, l2, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v138 = F_pg_regexec(m, v133, v128, v131, int32(0), int32(2), v16+int32(192))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	F_pfree(m, v128)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	if v138 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v138 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+4)))
	if v182 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	v333 = int32(0)
	v335 = base.B2i32(v138 != int32(1))
	goto L39
L51:
	;
	v148 = F_pg_regerror(m, v138, v16+int32(80))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	v152 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	if v152 == int32(0) {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v160 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v16 + int32(80)
	F_errmsg(m, int32(202643), v16-int32(-64))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(495117), int32(2861), int32(236184))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v277
	v281 = int32(0)
	v285 = F_list_make1_impl(m, int32(1), v16+int32(44))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L16
	} else {
		goto L89
	}
L59:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v187 == int32(43) {
		v277 = v181
		v278 = int32(0)
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v191 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	if v192 != 0 {
		v277 = v181
		v278 = v191
		goto L58
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v195 = F_strstr(m, v193, int32(548603))
	mBase = m.M
	if v195 == int32(0) {
		v277 = v181
		v278 = v191
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	if v198 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v201 = int32(1)
	v202 = int32(0)
	v205 = F_errstart(m, int32(15), v202)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L16
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v230 = F_strlen(m, v193)
	mBase = m.M
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v16)+204))
	v236 = F_palloc0(m, v230+(v198^int32(-1))+v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L16
	} else {
		goto L73
	}
L68:
	;
	if v205 == int32(0) {
		v333 = v202
		v335 = v201
		goto L39
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v213 + int32(1)
	F_errmsg(m, int32(691432), v16+int32(48))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(495117), int32(2885), int32(236184))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	v333 = v202
	v335 = v201
	goto L39
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v240 = v195 - v239
	if v240 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v16)+204))
	v247 = v246 - v244
	if v247 != 0 {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v241 = F__emscripten_memcpy_bulkmem(m, v236, v239, v240)
	mBase = m.M
	v242 = v241
	goto L77
L76:
	;
	v242 = v236
	goto L77
L77:
	;
	goto L74
L78:
	;
	v252 = F_strlen(m, v242)
	mBase = m.M
	v254 = F_strcpy(m, v252+v242, v195+int32(2))
	mBase = m.M
	goto L82
L79:
	;
	v248 = F__emscripten_memcpy_bulkmem(m, v242+v240, l2+v244, v247)
	mBase = m.M
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	v255 = F_strlen(m, v242)
	mBase = m.M
	v258 = F_palloc0(m, v255+int32(13))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = int32(0)
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+4)) = uint8(v262)
	v266 = v258 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v266
	v269 = v255 + v262
	if v269 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	F_pfree(m, v242)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L16
	} else {
		goto L88
	}
L85:
	;
	v270 = F__emscripten_memcpy_bulkmem(m, v266, v242, v269)
	mBase = m.M
	goto L87
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	v277 = v258
	v278 = v262
	goto L58
L89:
	;
	v287 = F_check_role_2(m, l1, v119, v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L16
	} else {
		goto L90
	}
L90:
	;
	if v278 == int32(0) {
		v333 = v287
		v335 = v281
		goto L39
	} else {
		goto L91
	}
L91:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	if v291 == int32(0) {
		v333 = v287
		v335 = v281
		goto L39
	} else {
		goto L92
	}
L92:
	;
	F_pg_regfree(m, v291)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L16
	} else {
		goto L93
	}
L93:
	;
	v333 = v287
	v335 = v281
	goto L39
L94:
	;
	if v320-v319 != 0 {
		goto L28
	} else {
		goto L102
	}
L95:
	;
	goto L94
L96:
	;
	if v299 != v300 {
		v319 = v299
		v320 = v300
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v304 = v296
	v305 = l2
	goto L98
L98:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+1)))
	if v309 == int32(0) {
		v319 = v308
		v320 = v309
		goto L95
	} else {
		goto L100
	}
L99:
	;
	v319 = v308
	v320 = v309
	goto L95
L100:
	;
	v312 = int32(1)
	if v308 == v309 {
		v304 = v304 + v312
		v305 = v305 + v312
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v322
	v329 = F_list_make1_impl(m, int32(1), v16+int32(40))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	v331 = F_check_role_2(m, l1, v119, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L16
	} else {
		goto L104
	}
L104:
	;
	v333 = v331
	v335 = int32(0)
	goto L39
L105:
	;
	v365 = v333
	v374 = v335 | (v333 ^ int32(1))
	goto L22
L106:
	;
	goto L26
L107:
	;
	if v365 != 0 {
		goto L21
	} else {
		goto L108
	}
L108:
	;
	v377 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	if v377 == int32(0) {
		goto L21
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
	F_errmsg(m, int32(679240), v16+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L16
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(495117), int32(3008), int32(236204))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L16
	} else {
		goto L112
	}
L112:
	;
	goto L21
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
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v476 int32
	_ = v476
	v6 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v476
L2:
	;
	v40 = int32(2)
	v45 = int32(9)
	v46 = v6
	v49 = v6
	goto L7
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v22 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v476 = v6
	goto L1
L6:
	;
	goto L5
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)))
	if v55 != int32(109) {
		v442 = v40
		v447 = v45
		v451 = v49
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v476 = v451
	goto L1
L9:
	;
	v453 = v46 + int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v453 < v454 {
		v40 = v442
		v45 = v447
		v46 = v453
		v49 = v451
		goto L7
	} else {
		goto L108
	}
L10:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+8)))
	if v58 != l1 {
		v442 = v40
		v447 = v45
		v451 = v49
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v60 = int32(0)
	if base.B2i32(l4 <= int32(0)) == v60 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v73 = v60
	v74 = v60
	v77 = int32(0)
	goto L15
L13:
	;
	v296 = v60
	v297 = v60
	goto L14
L14:
	;
	v309 = int32(0)
	if v296 == v309 {
		goto L60
	} else {
		goto L61
	}
L15:
	;
	v87 = v77 << (uint(int32(2)) % 32)
	v88 = l2 + v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v296 = v272
	v297 = v273
	goto L14
L17:
	;
	v286 = v77 + int32(1)
	if v286 != l4 {
		v73 = v272
		v74 = v273
		v77 = v286
		goto L15
	} else {
		goto L58
	}
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3+v87)))
	if v93 == int32(0) {
		v272 = v73
		v273 = v74
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v97 = int32(0)
	if v89 == v97 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	if v150 == int32(0) {
		v272 = v73
		v273 = v74
		goto L17
	} else {
		goto L36
	}
L23:
	;
	v150 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v96 == int32(0) {
		v141 = v97
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v150 = v141
	goto L22
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v107 < v106 {
		v141 = v97
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v109 = int32(1)
	if v106 <= v109 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v112 = v109
	goto L31
L30:
	;
	v112 = v106
	goto L31
L31:
	;
	v113 = int32(8)
	v118 = int32(0)
	goto L32
L32:
	;
	v125 = v118 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v89+v113+v125)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v96+v113))))
	v132 = v127 & (v129 ^ int32(-1))
	v134 = base.B2i32(v132 == int32(0))
	if v132 != 0 {
		v141 = v134
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v141 = v134
	goto L26
L34:
	;
	v136 = v118 + int32(1)
	if v136 != v112 {
		v118 = v136
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3+v87)))
	if v154 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v260 = F_bms_add_members(m, v73, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L48
	} else {
		goto L56
	}
L38:
	;
	v248 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v158 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v160 <= v158 {
		v248 = v158
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v173 = v158
	v177 = v158
	goto L42
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	if v184 == int32(0) {
		v272 = v73
		v273 = v74
		goto L17
	} else {
		goto L44
	}
L43:
	;
	v248 = v232
	goto L37
L44:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v187 <= int32(0) {
		v272 = v73
		v273 = v74
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v177<<(uint(int32(2))%32))))
	v202 = int32(0)
	goto L46
L46:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217+v202<<(uint(int32(2))%32))))
	v222 = F_equal(m, v221, v194)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v232 = F_bms_add_member(m, v173, v202)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L48
	} else {
		goto L54
	}
L48:
	;
	return int32(0)
L49:
	;
	if v222 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v229 = v202 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v229 < v230 {
		v202 = v229
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
	v272 = v73
	v273 = v74
	goto L17
L54:
	;
	v235 = v177 + int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v235 < v236 {
		v173 = v232
		v177 = v235
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	v262 = F_bms_add_members(m, v74, v248)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L48
	} else {
		goto L57
	}
L57:
	;
	v272 = v260
	v273 = v262
	goto L17
L58:
	;
	goto L16
L59:
	;
	v345 = int32(0)
	if v297 == v345 {
		goto L73
	} else {
		goto L74
	}
L60:
	;
	v344 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v316 = int32(1)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	if v317 <= v316 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v320 = v316
	goto L65
L64:
	;
	v320 = v317
	goto L65
L65:
	;
	v324 = int32(0)
	v326 = v309
	goto L66
L66:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v296+int32(8)+v324<<(uint(int32(2))%32))))
	if v332 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v344 = v335
	goto L59
L68:
	;
	v335 = v326 + base.I32_popcnt(v332)
	goto L70
L69:
	;
	v335 = v326
	goto L70
L70:
	;
	v337 = v324 + int32(1)
	if v337 != v320 {
		v324 = v337
		v326 = v335
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L67
L72:
	;
	F_bms_free(m, v296)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L48
	} else {
		goto L85
	}
L73:
	;
	v380 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v352 = int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v353 <= v352 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v356 = v352
	goto L78
L77:
	;
	v356 = v353
	goto L78
L78:
	;
	v360 = int32(0)
	v362 = v345
	goto L79
L79:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v297+int32(8)+v360<<(uint(int32(2))%32))))
	if v368 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v380 = v371
	goto L72
L81:
	;
	v371 = v362 + base.I32_popcnt(v368)
	goto L83
L82:
	;
	v371 = v362
	goto L83
L83:
	;
	v373 = v360 + int32(1)
	if v373 != v356 {
		v360 = v373
		v362 = v371
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	F_bms_free(m, v297)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L48
	} else {
		goto L86
	}
L86:
	;
	v385 = v380 + v344
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v387 = int32(0)
	if v386 == v387 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	if v423 != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v422 = int32(0)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v394 = int32(1)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v395 <= v394 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v398 = v394
	goto L93
L92:
	;
	v398 = v395
	goto L93
L93:
	;
	v402 = int32(0)
	v404 = v387
	goto L94
L94:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v386+int32(8)+v402<<(uint(int32(2))%32))))
	if v410 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v422 = v413
	goto L87
L96:
	;
	v413 = v404 + base.I32_popcnt(v410)
	goto L98
L97:
	;
	v413 = v404
	goto L98
L98:
	;
	v415 = v402 + int32(1)
	if v415 != v398 {
		v402 = v415
		v404 = v413
		goto L94
	} else {
		goto L99
	}
L99:
	;
	goto L95
L100:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v426 = v424
	goto L102
L101:
	;
	v426 = int32(0)
	goto L102
L102:
	;
	v427 = v426 + v422
	if v385 <= v40 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v385 != v40 {
		v442 = v40
		v447 = v45
		v451 = v49
		goto L9
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v442 = v385
	v447 = v427
	v451 = v54
	goto L9
L106:
	;
	if v45 <= v427 {
		v442 = v40
		v447 = v45
		v451 = v49
		goto L9
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
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
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
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
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v93 == int32(-1) {
		v386 = int32(0)
		goto L31
	} else {
		goto L32
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
	if v64 == v65 {
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
	if v65 <= int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v70 = v20
	goto L24
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v75 = F_bms_is_member(m, v70, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L3
L26:
	;
	if v75 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v79+v70)+20)) = uint8(v81)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v84 = v70 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v84 < v85 {
		v70 = v84
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	F_LWLockRelease(m, v6)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L94
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v93
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v101 = v99
	goto L34
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v101
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v245 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+(v6+int32(20))))))
	if v106 != int32(1) {
		goto L33
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(-1)
	F_LWLockRelease(m, v6)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L64
	}
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v109 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v233
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v233 != v236 {
		v101 = v233
		goto L34
	} else {
		goto L63
	}
L38:
	;
	if int32(0) <= v165 {
		v233 = v165
		goto L37
	} else {
		goto L49
	}
L39:
	;
	v165 = base.I32_ctz(v151) | v152<<(uint(int32(5))%32)
	goto L38
L40:
	;
	v165 = int32(-2)
	goto L38
L41:
	;
	v116 = v101 + int32(1)
	v118 = base.I32_div_s(v116, int32(32))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v119 <= v118 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v122 = v109 + int32(8)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v118<<(uint(int32(2))%32))))
	v129 = v126 & (int32(-1) << (uint(v116) % 32))
	if v129 != 0 {
		v151 = v129
		v152 = v118
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v131 = v118 + int32(1)
	if v131 == v119 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v134 = v131
	goto L45
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122+v134<<(uint(int32(2))%32))))
	if v141 != 0 {
		v151 = v141
		v152 = v134
		goto L39
	} else {
		goto L47
	}
L46:
	;
	goto L40
L47:
	;
	v143 = v134 + int32(1)
	if v143 != v119 {
		v134 = v143
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v168 <= v169 {
		v233 = v168
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v171 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if int32(0) <= v229 {
		v233 = v229
		goto L37
	} else {
		goto L62
	}
L52:
	;
	v229 = base.I32_ctz(v215) | v216<<(uint(int32(5))%32)
	goto L51
L53:
	;
	v229 = int32(-2)
	goto L51
L54:
	;
	v180 = v169 - int32(1) + int32(1)
	v182 = base.I32_div_s(v180, int32(32))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v183 <= v182 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v186 = v171 + int32(8)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v182<<(uint(int32(2))%32))))
	v193 = v190 & (int32(-1) << (uint(v180) % 32))
	if v193 != 0 {
		v215 = v193
		v216 = v182
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v195 = v182 + int32(1)
	if v195 == v183 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v198 = v195
	goto L58
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v186+v198<<(uint(int32(2))%32))))
	if v205 != 0 {
		v215 = v205
		v216 = v198
		goto L52
	} else {
		goto L60
	}
L59:
	;
	goto L53
L60:
	;
	v207 = v198 + int32(1)
	if v207 != v183 {
		v198 = v207
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v233 = v232
	goto L37
L63:
	;
	goto L35
L64:
	;
	return int32(0)
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v302
	if v302 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L66:
	;
	v302 = base.I32_ctz(v288) | v289<<(uint(int32(5))%32)
	goto L65
L67:
	;
	v302 = int32(-2)
	goto L65
L68:
	;
	v253 = v246 + int32(1)
	v255 = base.I32_div_s(v253, int32(32))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v256 <= v255 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v259 = v245 + int32(8)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259+v255<<(uint(int32(2))%32))))
	v266 = v263 & (int32(-1) << (uint(v253) % 32))
	if v266 != 0 {
		v288 = v266
		v289 = v255
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v268 = v255 + int32(1)
	if v268 == v256 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v271 = v268
	goto L72
L72:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v259+v271<<(uint(int32(2))%32))))
	if v278 != 0 {
		v288 = v278
		v289 = v271
		goto L66
	} else {
		goto L74
	}
L73:
	;
	goto L67
L74:
	;
	v280 = v271 + int32(1)
	if v280 != v256 {
		v271 = v280
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v307 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	goto L78
L78:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v374 <= v373 {
		v386 = int32(1)
		goto L31
	} else {
		goto L93
	}
L79:
	;
	if v366 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L80:
	;
	v366 = base.I32_ctz(v352) | v353<<(uint(int32(5))%32)
	goto L79
L81:
	;
	v366 = int32(-2)
	goto L79
L82:
	;
	v317 = v308 - int32(1) + int32(1)
	v319 = base.I32_div_s(v317, int32(32))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	if v320 <= v319 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v323 = v307 + int32(8)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323+v319<<(uint(int32(2))%32))))
	v330 = v327 & (int32(-1) << (uint(v317) % 32))
	if v330 != 0 {
		v352 = v330
		v353 = v319
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v332 = v319 + int32(1)
	if v332 == v320 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v335 = v332
	goto L86
L86:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v323+v335<<(uint(int32(2))%32))))
	if v342 != 0 {
		v352 = v342
		v353 = v335
		goto L80
	} else {
		goto L88
	}
L87:
	;
	goto L81
L88:
	;
	v344 = v335 + int32(1)
	if v344 != v320 {
		v335 = v344
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v369 = int32(-1)
	goto L92
L91:
	;
	v369 = v366
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v369
	goto L78
L93:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v378 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v376+v373)+20)) = uint8(v378)
	v386 = v378
	goto L31
L94:
	;
	return v386
}
func F_choose_next_subplan_locally(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	v2 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v8 != 0 {
		v189 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v189
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v9 != int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v15 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v17 = int32(0)
	v19 = F_ExecFindMatchingSubPlans(m, v16, v17, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v23)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v19
	goto L3
L9:
	;
	if v171 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L10:
	;
	if v27 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	goto L12
L12:
	;
	if v27 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L13:
	;
	v171 = v87
	goto L9
L14:
	;
	v87 = base.I32_ctz(v73) | v74<<(uint(int32(5))%32)
	goto L13
L15:
	;
	v87 = int32(-2)
	goto L13
L16:
	;
	v38 = v9 + int32(1)
	v40 = base.I32_div_s(v38, int32(32))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v41 <= v40 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = v27 + int32(8)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v51 = v48 & (int32(-1) << (uint(v38) % 32))
	if v51 != 0 {
		v73 = v51
		v74 = v40
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v53 = v40 + int32(1)
	if v53 == v41 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v56 = v53
	goto L20
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44+v56<<(uint(int32(2))%32))))
	if v63 != 0 {
		v73 = v63
		v74 = v56
		goto L14
	} else {
		goto L22
	}
L21:
	;
	goto L15
L22:
	;
	v65 = v56 + int32(1)
	if v65 != v41 {
		v56 = v65
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v171 = v163
	goto L9
L25:
	;
	v163 = base.I32_clz(v146) | v144<<(uint(int32(5))%32) ^ int32(31)
	goto L24
L26:
	;
	v163 = int32(-2)
	goto L24
L27:
	;
	if v9 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	if v9 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v97 = v94 << (uint(int32(5)) % 32)
	goto L31
L30:
	;
	v97 = v9
	goto L31
L31:
	;
	v99 = v97 - int32(1)
	if v99 < int32(-31) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v103 = v27 + int32(8)
	v105 = base.I32_div_s(v99, int32(32))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+v105<<(uint(int32(2))%32))))
	v117 = v109 & int32(base.Ui32(int32(-1))>>(uint(v105<<(uint(int32(5))%32)-v99+int32(31))%32))
	if v117 != 0 {
		v144 = v105
		v146 = v117
		goto L25
	} else {
		goto L33
	}
L33:
	;
	if v99 < int32(32) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v121 = v105
	goto L35
L35:
	;
	v128 = v121 - int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v103+v128<<(uint(int32(2))%32))))
	if v132 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L26
L37:
	;
	v144 = v128
	v146 = v132
	goto L25
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(int32(1)) < base.Ui32(v121) {
		v121 = v128
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v174 <= int32(0) {
		v189 = v2
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v171
	v189 = int32(1)
	goto L1
L44:
	;
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(v177)
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v98 int32
	_ = v98
	var v104 int64
	_ = v104
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v179 int32
	_ = v179
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
		goto L88
	}
L8:
	;
	v24 = v17
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 == v29 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v27 != 0 {
		v24 = v27
		goto L9
	} else {
		goto L87
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
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
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v37 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L47
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+4)))
	if v45 < int32(0) {
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
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v50 = v48 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v50) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if int32(1)<<(uint(v50)%32)&int32(163841) == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v59 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	if v60 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v72 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v45*int32(24))+12)) = v68
	v72 = v68
	goto L26
L28:
	;
	goto L29
L29:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v70
	v72 = v70
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+36)) = v60
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+32)) = int64(0)
	goto L21
L33:
	;
	if v78 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v78
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v78
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v79
	goto L39
L38:
	;
	goto L39
L39:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v85 - int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v90 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v96 = v37 + int32(8)
	if v89 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v89
	goto L40
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v89
	goto L40
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v90
	goto L46
L45:
	;
	goto L46
L46:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v98 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v104
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v37
	goto L20
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v118 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(-1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v199 != 0 {
		goto L80
	} else {
		goto L81
	}
L49:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118)+4)))
	if v126 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L47
L53:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v160 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v131 = v129 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v131) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	if int32(1)<<(uint(v131)%32)&int32(163841) == int32(0) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v140 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	if v141 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v153 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v126*int32(24))+12)) = v149
	v153 = v149
	goto L58
L60:
	;
	goto L61
L61:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+32)) = v151
	v153 = v151
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+36)) = v141
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118)+32)) = int64(0)
	goto L53
L65:
	;
	if v159 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v159
	goto L65
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v159
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = v160
	goto L71
L70:
	;
	goto L71
L71:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = v166 - int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	if v171 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v177 = v118 + int32(8)
	if v170 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+16)) = v170
	goto L72
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v170
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+28)) = v171
	goto L78
L77:
	;
	goto L78
L78:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v179 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	v185 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v177)+16)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v177))) = v185
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v118
	goto L52
L79:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v198 != 0 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+32)) = v198
	goto L79
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v198
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v24
	goto L12
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+28)) = v202
	goto L83
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v202
	goto L83
L87:
	;
	goto L10
L88:
	;
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v226 = v223
	v227 = v224
	goto L92
L90:
	;
	v235 = v223
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v235
	goto L3
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v226
	v232 = v226 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	if v233 != 0 {
		v226 = v232
		v227 = v233
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v235 = v232
	goto L91
L94:
	;
	goto L93
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	if v36 != v16&int32(65535) {
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
	if v55 != v16&int32(65535) {
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
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 float64
	_ = v68
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
				v45 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
				v86 = int32(0)
				return v86
			} else {
				v49 = F_palloc(m, int32(16))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = F_lseg_interpt_line(m, v49, v12, v11)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v68 = float64(0)
							if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v86 = v49
							} else {
								v76 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
								v86 = int32(0)
							}
							return v86
						} else {
							v54 = F_line_closept_point(m, int32(0), v11, v12)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v57 = F_line_closept_point(m, int32(0), v11, v14)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									v59 = base.F64_lt(v54, v57)
									if v59 != 0 {
										v60 = v54
									} else {
										v60 = v57
									}
									if v49 == int32(0) {
										v68 = v60
									} else {
										if v59 != 0 {
											v63 = v12
										} else {
											v63 = v14
										}
										v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
										*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
										v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
										v68 = v60
									}
									if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v86 = v49
									} else {
										v76 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
										v86 = int32(0)
									}
									return v86
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
					v45 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
					v86 = int32(0)
					return v86
				} else {
					v49 = F_palloc(m, int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = F_lseg_interpt_line(m, v49, v12, v11)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 != 0 {
								v68 = float64(0)
								if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v86 = v49
								} else {
									v76 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
									v86 = int32(0)
								}
								return v86
							} else {
								v54 = F_line_closept_point(m, int32(0), v11, v12)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v57 = F_line_closept_point(m, int32(0), v11, v14)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = base.F64_lt(v54, v57)
										if v59 != 0 {
											v60 = v54
										} else {
											v60 = v57
										}
										if v49 == int32(0) {
											v68 = v60
										} else {
											if v59 != 0 {
												v63 = v12
											} else {
												v63 = v14
											}
											v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
											*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
											v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
											v68 = v60
										}
										if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v86 = v49
										} else {
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											v86 = int32(0)
										}
										return v86
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
					v89 = m.ExcPending
					if v89 != 0 {
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
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							v86 = int32(0)
							return v86
						} else {
							v49 = F_palloc(m, int32(16))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = F_lseg_interpt_line(m, v49, v12, v11)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v51 != 0 {
										v68 = float64(0)
										if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v86 = v49
										} else {
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											v86 = int32(0)
										}
										return v86
									} else {
										v54 = F_line_closept_point(m, int32(0), v11, v12)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v57 = F_line_closept_point(m, int32(0), v11, v14)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												v59 = base.F64_lt(v54, v57)
												if v59 != 0 {
													v60 = v54
												} else {
													v60 = v57
												}
												if v49 == int32(0) {
													v68 = v60
												} else {
													if v59 != 0 {
														v63 = v12
													} else {
														v63 = v14
													}
													v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
													*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
													v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
													v68 = v60
												}
												if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
													v86 = v49
												} else {
													v76 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
													v86 = int32(0)
												}
												return v86
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
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v41 = v30
							if base.F64_eq(v41, v15) != 0 {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v86 = int32(0)
								return v86
							} else {
								v49 = F_palloc(m, int32(16))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = F_lseg_interpt_line(m, v49, v12, v11)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 != 0 {
											v68 = float64(0)
											if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
												v86 = v49
											} else {
												v76 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
												v86 = int32(0)
											}
											return v86
										} else {
											v54 = F_line_closept_point(m, int32(0), v11, v12)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												v57 = F_line_closept_point(m, int32(0), v11, v14)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													v59 = base.F64_lt(v54, v57)
													if v59 != 0 {
														v60 = v54
													} else {
														v60 = v57
													}
													if v49 == int32(0) {
														v68 = v60
													} else {
														if v59 != 0 {
															v63 = v12
														} else {
															v63 = v14
														}
														v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
														*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
														v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
														v68 = v60
													}
													if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
														v86 = v49
													} else {
														v76 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
														v86 = int32(0)
													}
													return v86
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
	v6 = int32(4374672)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[1136]))
	v8 = F_close(m, v7)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1136])) = int32(-1)
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
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
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
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
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v308 int32
	_ = v308
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
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 float64
	_ = v516
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
	var v530 int32
	_ = v530
	var v531 float64
	_ = v531
	var v532 float64
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 float64
	_ = v540
	var v541 float64
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 float64
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 float64
	_ = v557
	var v560 int32
	_ = v560
	var v561 float64
	_ = v561
	var v567 float64
	_ = v567
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v598 int32
	_ = v598
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v745 int32
	_ = v745
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 float64
	_ = v771
	var v772 float64
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 float64
	_ = v823
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v919 int32
	_ = v919
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(400)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v33 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[31]))
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
		goto L18
	} else {
		goto L19
	}
L7:
	;
	goto L6
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v43 != int32(1) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(4474964)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v49 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v48 + v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v52 + v49
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = v28
	v59 = v39 + int32(232)
	if v59&int32(3) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v86 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v85 + v86
	v89 = int32(4474964)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v91 - v86
	goto L7
L11:
	;
	v65 = v39 + int32(392)
	if base.Ui32(v65) <= base.Ui32(v59) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v82 = F___memset(m, v59, int32(0), int32(160))
	mBase = m.M
	goto L10
L14:
	;
	v69 = v39 + int32(236)
	if base.Ui32(v69) < base.Ui32(v65) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = v65
	goto L17
L16:
	;
	v71 = v69
	goto L17
L17:
	;
	v79 = F___memset(m, v59, int32(0), (v71-v39-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L10
L18:
	;
	v100 = int64(1)
	goto L20
L19:
	;
	v100 = int64(2)
	goto L20
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v103 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(124)))) = v139
	v142 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(120)))) = v142
	goto L25
L22:
	;
	goto L21
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v107 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v110 = int32(4474964)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v112 + v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v116 + v113
	*(*int64)(unsafe.Add(mBase, uint32(v103+int32(0))+232)) = v100
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v124 + v113
	v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v113
	goto L22
L25:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+80))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
	*(*int32)(unsafe.Add(mBase, _consts[122])) = v146 | int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v145
	goto L26
L26:
	;
	v154 = int32(4478344)
	v156 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v158 = v156 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v158
	goto L27
L27:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v29&int32(2) != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L4
	} else {
		goto L211
	}
L30:
	;
	F_errmsg(m, int32(142624), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L209
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L4
	} else {
		goto L205
	}
L32:
	;
	F_AtEOXact_GUC(m, int32(0), v158)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L198
	}
L33:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+118)))
	if v224 != int32(116) {
		goto L66
	} else {
		goto L67
	}
L34:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+117)))
	if v220 == int32(1) {
		goto L31
	} else {
		goto L65
	}
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v26)+124))
	v164 = F_pg_class_aclcheck(m, v28, v162, int64(16384))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if l1 != 0 {
		goto L34
	} else {
		goto L64
	}
L38:
	;
	if v164 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v168 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+118)))
	if v187 != int32(116) {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	if v168 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v170 = F_get_rel_name(m, v28)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v170
	F_errmsg(m, int32(103262), v26+int32(112))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(490230), int32(1752), int32(260698))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L32
L50:
	;
	if l1 == int32(0) {
		v223 = v186
		goto L33
	} else {
		goto L54
	}
L51:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v190 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	goto L32
L54:
	;
	v197 = int32(0)
	v200 = F_SearchSysCacheExists(m, int32(57), l1, v197, v197, v197)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v200 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v207&int32(4) == int32(0) {
		goto L34
	} else {
		goto L60
	}
L59:
	;
	goto L32
L60:
	;
	v212 = F_get_index_isclustered(m, l1)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v212 != 0 {
		goto L34
	} else {
		goto L62
	}
L62:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	goto L32
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v223 = v217
	goto L33
L65:
	;
	v223 = v219
	goto L33
L66:
	;
	if l1 != 0 {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v227 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	if l1 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(142782), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(490230), int32(425), int32(304647))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v246 = int32(520026)
	goto L76
L75:
	;
	v246 = int32(525931)
	goto L76
L76:
	;
	F_CheckTableNotInUse(m, l0, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	if l1 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_check_index_is_clusterable(m, l0, l1, int32(8))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	v256 = int32(0)
	goto L80
L80:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+119)))
	if v258 != int32(109) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v253 = F_index_open(m, l1, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v256 = v253
	goto L80
L83:
	;
	F_TransferPredicateLocksToHeapRelation(m, l0)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L87
	}
L84:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+129)))
	if v261 != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	goto L32
L87:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+92))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v268)+84))
	if v256 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	F_mark_index_clustered(m, l0, v271, int32(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	v276 = v268
	goto L90
L90:
	;
	v277 = int32(*(*int8)(unsafe.Add(mBase, uint32(v276)+118)))
	v279 = int32(1)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v280) < base.Ui32(int32(12000)) {
		v289 = v279
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v276 = v275
	goto L90
L92:
	;
	v291 = F_make_new_heap(m, v267, v269, v270, v277, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L96
	}
L93:
	;
	goto L92
L94:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+68))
	if v284 == int32(99) {
		v289 = v279
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v287 = F_isTempToastNamespace(m, v284)
	mBase = m.M
	v289 = v287
	goto L93
L96:
	;
	v294 = F_table_open(m, v291, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v296 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+312)) = v296
	*(*int64)(unsafe.Add(mBase, uint32(v26)+304)) = v296
	*(*int64)(unsafe.Add(mBase, uint32(v26)+296)) = v296
	F_getrusage(m, v26+int32(144))
	mBase = m.M
	F___gettimeofday(m, v26+int32(128))
	mBase = m.M
	goto L98
L98:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+68))
	v310 = F_get_namespace_name(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+112))
	if v313 == int32(0) {
		v330 = v4
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v331 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+392)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+384)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+376)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+368)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+360)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+352)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+344)) = v331
	v349 = F_vacuum_get_cutoffs(m, l0, v26+int32(344), v26+int32(320))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L105
	}
L101:
	;
	F_LockRelationOid(m, v313, int32(8))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+112))
	if v320 == int32(0) {
		v330 = v4
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v294)+48))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+112))
	if v324 == int32(0) {
		v330 = v4
		goto L100
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+264)) = v320
	v330 = int32(1)
	goto L100
L105:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+136))
	if v352 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+140))
	if v374 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v26)+336))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v352))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v355)) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if v367 == int32(0) {
		goto L106
	} else {
		goto L112
	}
L109:
	;
	v367 = base.B2i32(base.Ui32(v355) < base.Ui32(v352))
	goto L108
L110:
	;
	goto L111
L111:
	;
	v367 = int32(base.Ui32(v355-v352) >> (uint(int32(31)) % 32))
	goto L108
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+336)) = v352
	goto L106
L113:
	;
	if v29&int32(1) != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v26)+340))
	goto L115
L115:
	;
	if int32(base.Ui32(v377-v374)>>(uint(int32(31))%32)) == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+340)) = v374
	goto L113
L117:
	;
	v386 = int32(17)
	goto L119
L118:
	;
	v386 = int32(13)
	goto L119
L119:
	;
	if v256 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v26)+328))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+124))
	m.T0[v757].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v294, v256, v726, v745, v26+int32(336), v26+int32(340), v26+int32(312), v26+int32(304), v26+int32(296))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L168
	}
L121:
	;
	F_errfinish(m, int32(490230), v700, int32(499067))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L167
	}
L122:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v256)+48))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+84))
	if v388 == int32(403) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	goto L124
L124:
	;
	v676 = int32(0)
	v678 = F_errstart(m, v386, v676)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L164
	}
L125:
	;
	v660 = F_errstart(m, v386, int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L161
	}
L126:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	v393 = m.G0
	v395 = v393 - int32(112)
	m.G0 = v395
	v397 = int32(1)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	if v399 != v397 {
		v598 = v397
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v636 = int32(0)
	v638 = F_errstart(m, v386, v636)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L4
	} else {
		goto L158
	}
L129:
	;
	m.G0 = v395 + int32(112)
	if v598 != 0 {
		goto L125
	} else {
		goto L157
	}
L130:
	;
	v403 = F_palloc0(m, int32(168))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v403))) = int64(4294967363)
	v408 = F_palloc0(m, int32(92))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = int32(266)
	v413 = F_palloc0(m, int32(384))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v413)+8)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v413)+4)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = int32(267)
	v422 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+344)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v413)+280)) = v422
	v427 = F_palloc0(m, int32(8))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v395)+20)) = v427
	v436 = F_list_make1_impl(m, int32(1), v395+int32(12))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+84)) = v436
	v440 = F_palloc0(m, int32(136))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v442 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v440)+24)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v440)+16)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v440)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(101)
	v449 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+124)) = uint16(v449)
	v451 = int32(29184)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+20)) = uint16(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v395)+16)) = v440
	v458 = F_list_make1_impl(m, v442, v395+int32(8))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+52)) = v458
	v463 = F_addRTEPermissionInfo(m, v403+int32(56), v440)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_setup_simple_rel_arrays(m, v413)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v469 = F_build_simple_rel(m, v413, int32(1), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v469)+108))
	if v471 == int32(0) {
		v598 = v397
		goto L129
	} else {
		goto L141
	}
L141:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v474 <= int32(0) {
		v598 = v397
		goto L129
	} else {
		goto L142
	}
L142:
	;
	v477 = int32(0)
	if v477 < v474 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v481 = v474
	goto L145
L144:
	;
	v481 = v477
	goto L145
L145:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v487 = v477
	goto L146
L146:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v482+v487<<(uint(int32(2))%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v392 != v510 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v469)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v469)+16)) = v516
	v519 = F_get_relation_data_width(m, v391, int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L152
	}
L148:
	;
	v512 = int32(1)
	v514 = v487 + v512
	if v481 != v514 {
		v487 = v514
		goto L146
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	v598 = v512
	goto L129
L152:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v469)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v519
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v469)+116))
	*(*float64)(unsafe.Add(mBase, uint32(v413)+288)) = base.F64_convert_i32_u(v523)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v509)+84))
	F_cost_qual_eval(m, v395+int32(96), v528, v413)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v531 = *(*float64)(unsafe.Add(mBase, uint32(v395)+104))
	v532 = *(*float64)(unsafe.Add(mBase, uint32(v395)+96))
	v534 = v395 + int32(24)
	v535 = int32(0)
	v537 = F_create_seqscan_path(m, v413, v469, v535, v535)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v537)+40))
	v540 = *(*float64)(unsafe.Add(mBase, uint32(v537)+56))
	v541 = *(*float64)(unsafe.Add(mBase, uint32(v469)+120))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v469)+28))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+32))
	v544 = base.F64_add(v532, v531)
	v547 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	v550 = m.G0
	v551 = int32(16)
	v552 = v550 - v551
	m.G0 = v552
	F_cost_tuplesort(m, v552+int32(8), v552, v541, v543, base.F64_add(v544, v544), v547, float64(-1))
	mBase = m.M
	v557 = *(*float64)(unsafe.Add(mBase, uint32(v552)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+32)) = v541
	v560 = int32(*(*uint8)(unsafe.Add(mBase, _consts[255])))
	v561 = base.F64_add(v540, v557)
	*(*float64)(unsafe.Add(mBase, uint32(v534)+48)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v534)+40)) = v539 + (v560 ^ int32(1))
	v567 = *(*float64)(unsafe.Add(mBase, uint32(v552)))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+56)) = base.F64_add(v561, v567)
	m.G0 = v552 + v551
	goto L155
L155:
	;
	v573 = int32(0)
	v582 = F_create_index_path(m, v413, v509, v573, v573, v573, v573, int32(1), v573, v573, float64(1), v573)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v584 = *(*float64)(unsafe.Add(mBase, uint32(v395)+80))
	v585 = *(*float64)(unsafe.Add(mBase, uint32(v582)+56))
	v598 = base.F64_lt(v584, v585)
	goto L129
L157:
	;
	goto L128
L158:
	;
	if v638 == int32(0) {
		v726 = v636
		goto L120
	} else {
		goto L159
	}
L159:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v256)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v310
	v645 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v643 + v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v642 + v645
	F_errmsg(m, int32(689155), v26+int32(96))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v700 = int32(964)
	v717 = int32(0)
	goto L121
L161:
	;
	if v660 == int32(0) {
		v726 = int32(1)
		goto L120
	} else {
		goto L162
	}
L162:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v664 + int32(4)
	F_errmsg(m, int32(78238), v26+int32(80))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v700 = int32(969)
	v717 = int32(1)
	goto L121
L164:
	;
	if v678 == int32(0) {
		v726 = v676
		goto L120
	} else {
		goto L165
	}
L165:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v682 + int32(4)
	F_errmsg(m, int32(671130), v26-int32(-64))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v700 = int32(974)
	v717 = int32(0)
	goto L121
L167:
	;
	v726 = v717
	goto L120
L168:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v26)+340))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v26)+336))
	v762 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+264)) = v762
	v765 = F_RelationGetNumberOfBlocksInFork(m, v294, v762)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v768 = F_errstart(m, v386, int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	if v768 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v771 = *(*float64)(unsafe.Add(mBase, uint32(v26)+304))
	v772 = *(*float64)(unsafe.Add(mBase, uint32(v26)+312))
	v774 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v810 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L179
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v774
	*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v772
	*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v770 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v310
	F_errmsg(m, int32(169380), v26+int32(32))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v788 = *(*float64)(unsafe.Add(mBase, uint32(v26)+296))
	v791 = F_pg_rusage_show(m, v26+int32(128))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v791
	*(*float64)(unsafe.Add(mBase, uint32(v26)+16)) = v788
	F_errdetail(m, int32(586023), v26+int32(16))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(490230), int32(1007), int32(499067))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	goto L173
L179:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v294)+56))
	v815 = F_SearchSysCacheCopy(m, int32(57), v813, int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	if v815 == int32(0) {
		goto L29
	} else {
		goto L181
	}
L181:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819)+22)))
	v821 = v819 + v820
	*(*int32)(unsafe.Add(mBase, uint32(v821)+96)) = v765
	v823 = *(*float64)(unsafe.Add(mBase, uint32(v26)+312))
	*(*float32)(unsafe.Add(mBase, uint32(v821)+100)) = base.F32_demote_f64(v823)
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v826 != int32(1259) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	F_pfree(m, v815)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L188
	}
L183:
	;
	F_CatalogTupleUpdate(m, v810, v815+int32(4), v815)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_CacheInvalidateRelcacheByTuple(m, v815)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L187
	}
L186:
	;
	goto L182
L187:
	;
	goto L182
L188:
	;
	F_sequence_close(m, v810, int32(3))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	F_sequence_close(m, l0, int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	if v256 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_relation_close(m, v256, int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	F_sequence_close(m, v294, int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L196
	}
L195:
	;
	goto L194
L196:
	;
	F_finish_heap_swap(m, v267, v291, v289, v330, int32(0), int32(1), v761, v760, v277)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	goto L32
L198:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v26)+124))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
	*(*int32)(unsafe.Add(mBase, _consts[122])) = v882
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v881
	goto L199
L199:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v889 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	m.G0 = v26 + int32(400)
	return
L201:
	;
	goto L200
L202:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v893 != int32(1) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v889)+220))
	if v896 == int32(0) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v899 = int32(4474964)
	v901 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v902 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v901 + v902
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	*(*int32)(unsafe.Add(mBase, uint32(v889))) = v905 + v902
	v909 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v889)+220)) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v889)+224)) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v889))) = v905 + int32(2)
	v919 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v919 - v902
	goto L201
L205:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	F_errmsg(m, int32(324198), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(490230), int32(410), int32(304647))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errfinish(m, int32(490230), int32(421), int32(304647))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
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
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v294)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v955
	F_errmsg_internal(m, int32(46032), v26)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(490230), int32(1016), int32(499067))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v6) <= base.Ui32(v5) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v10 = base.B2i32(base.Ui32(v8) < base.Ui32(v5))
	} else {
		v10 = int32(-1)
	}
	return v10
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
	var v32 int32
	_ = v32
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
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
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v8 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v250
L2:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v187 == int32(0) {
		goto L62
	} else {
		goto L63
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
	v174 = int32(0)
	v176 = F_hash_search(m, v8, l0, v174, v174)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L59
	} else {
		goto L60
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
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v66 {
		goto L23
	} else {
		goto L24
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
	v57 = v19 + int32(1)
	if v57 != v11 {
		v19 = v57
		goto L9
	} else {
		goto L22
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v32 == int32(0) {
		v51 = v31
		v52 = v32
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v52-v51 != 0 {
		goto L11
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	if v31 != v32 {
		v51 = v31
		v52 = v32
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = v26
	v37 = l0
	goto L17
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v40
		v52 = v41
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v51 = v40
	v52 = v41
	goto L14
L19:
	;
	v44 = int32(1)
	if v40 == v41 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return int32(0)
L22:
	;
	goto L10
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v74 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v121 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v69+v74<<(uint(int32(2))%32))))
	if v81 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v112 = v74 + int32(1)
	if v112 != v66 {
		v74 = v112
		goto L26
	} else {
		goto L39
	}
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v87 == int32(0) {
		v106 = v86
		v107 = v87
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v107-v106 != 0 {
		goto L28
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	if v86 != v87 {
		v106 = v86
		v107 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v91 = v81
	v92 = l0
	goto L34
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v95
		v107 = v96
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v106 = v95
	v107 = v96
	goto L31
L36:
	;
	v99 = int32(1)
	if v95 == v96 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	return int32(0)
L39:
	;
	goto L27
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v124 <= int32(0) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v127 = int32(0)
	if v127 < v124 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v131 = v124
	goto L44
L43:
	;
	v131 = v127
	goto L44
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v136 = v127
	goto L45
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132+v136<<(uint(int32(2))%32))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v147 == int32(0) {
		v166 = v146
		v167 = v147
		goto L48
	} else {
		goto L49
	}
L46:
	;
	return int32(0)
L47:
	;
	if v167-v166 != 0 {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	goto L47
L49:
	;
	if v146 != v147 {
		v166 = v146
		v167 = v147
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v151 = v143
	v152 = l0
	goto L51
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v156 == int32(0) {
		v166 = v155
		v167 = v156
		goto L48
	} else {
		goto L53
	}
L52:
	;
	v166 = v155
	v167 = v156
	goto L48
L53:
	;
	v159 = int32(1)
	if v155 == v156 {
		v151 = v151 + v159
		v152 = v152 + v159
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v170 = v136 + int32(1)
	if v131 != v170 {
		v136 = v170
		goto L45
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L46
L58:
	;
	goto L2
L59:
	;
	return int32(0)
L60:
	;
	if v176 != 0 {
		v250 = v4
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L2
L62:
	;
	return int32(1)
L63:
	;
	goto L64
L64:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v193 <= int32(0) {
		v250 = int32(1)
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v196 = int32(0)
	if v196 < v193 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v199 = v193
	goto L68
L67:
	;
	v199 = v196
	goto L68
L68:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v205 = int32(0)
	goto L69
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200+v205<<(uint(int32(2))%32))))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v216 == int32(0) {
		v235 = v215
		v236 = v216
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v250 = v239
	goto L1
L71:
	;
	v238 = int32(0)
	v239 = base.B2i32(v237 != v238)
	if v237 == v238 {
		v250 = v239
		goto L1
	} else {
		goto L79
	}
L72:
	;
	v237 = v236 - v235
	goto L71
L73:
	;
	if v215 != v216 {
		v235 = v215
		v236 = v216
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v220 = v212
	v221 = l0
	goto L75
L75:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	if v225 == int32(0) {
		v235 = v224
		v236 = v225
		goto L72
	} else {
		goto L77
	}
L76:
	;
	v235 = v224
	v236 = v225
	goto L72
L77:
	;
	v228 = int32(1)
	if v224 == v225 {
		v220 = v220 + v228
		v221 = v221 + v228
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v243 = v205 + int32(1)
	if v243 != v199 {
		v205 = v243
		goto L69
	} else {
		goto L80
	}
L80:
	;
	goto L70
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
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
			v35 = *(*int32)(unsafe.Add(mBase, uint32((v29^v23)<<(uint(int32(2))%32))+uint32(_consts[1099])))
			v38 = int32(16)
			v43 = int32(24)
			v46 = v35 ^ (v24<<(uint(v27)%32)&int32(65280) | v25<<(uint(v38)%32)&int32(16711680) | v26<<(uint(v43)%32))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v63 int32
	_ = v63
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
	if v38 < v6 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v11 = int32(4)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v13&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v26 = int32(1)
	if v8&v26 != 0 {
		v38 = int32(base.Ui32(v8)>>(uint(v26)%32)) - v26
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v22 = v11
	goto L7
L6:
	;
	v22 = base.B2i32(v13 == int32(18)) << (uint(v11) % 32)
	goto L7
L7:
	;
	if v13 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v11
	goto L10
L9:
	;
	v25 = v22
	goto L10
L10:
	;
	v38 = v25
	goto L1
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(1)
L13:
	;
	goto L14
L14:
	;
	if v6 < v38 {
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(1)
	if v8&v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = v46
	goto L20
L19:
	;
	v50 = int32(4)
	goto L20
L20:
	;
	v51 = v7 + v50
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
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v58 = v45
	v59 = v51
	v60 = v6
	v61 = v57
	goto L29
L26:
	;
	v83 = v51
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
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v61 != v63 {
		v78 = v59
		v80 = v61
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
	if v63 == int32(0) {
		v78 = v59
		v80 = v61
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v68 = v60 - int32(1)
	if v68 == int32(0) {
		v78 = v59
		v80 = v61
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v71 = int32(1)
	v72 = v59 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v73 != 0 {
		v58 = v58 + v71
		v59 = v72
		v60 = v68
		v61 = v73
		goto L29
	} else {
		goto L34
	}
L34:
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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
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
	v40 = int32(49152)
	v41 = v19 & v40
	if v41 != v40 {
		if v41 != int32(32768) {
			v52 = v41
		} else {
			v52 = v19 << (uint(int32(1)) % 32) & int32(16384)
		}
	} else {
		v52 = v19 & int32(61440)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v52
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v54
	v63 = base.B2i32(v20 < v54)
	if v20 < v54 {
		v64 = int32(base.Ui32(v19)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v64 = v19 & int32(16383)
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
	v97 = int32(49152)
	v98 = v76 & v97
	if v98 != v97 {
		if v98 != int32(32768) {
			v109 = v98
		} else {
			v109 = v76 << (uint(int32(1)) % 32) & int32(16384)
		}
	} else {
		v109 = v76 & int32(61440)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v109
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v111
	v120 = base.B2i32(v77 < v111)
	if v77 < v111 {
		v121 = int32(base.Ui32(v76)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v121 = v76 & int32(16383)
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
	v154 = int32(49152)
	v155 = v133 & v154
	if v155 != v154 {
		if v155 != int32(32768) {
			v166 = v155
		} else {
			v166 = v133 << (uint(int32(1)) % 32) & int32(16384)
		}
	} else {
		v166 = v133 & int32(61440)
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
		v183 = v133 & int32(16383)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v183
	v186 = v12 + int32(8)
	F_sub_var(m, v186, v12+int32(56), v186)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		return
	} else {
		v194 = v12 + int32(32)
		F_sub_var(m, v194, v12+int32(56), v194)
		mBase = m.M
		v200 = m.ExcPending
		if v200 != 0 {
			return
		} else {
			v202 = v12 + int32(8)
			v205 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			F_mul_var(m, v202, l3, v202, v205+v206)
			mBase = m.M
			v209 = m.ExcPending
			if v209 != 0 {
				return
			} else {
				v214 = int32(0)
				F_div_var(m, v12+int32(8), v12+int32(32), l4, v214, v214, int32(1))
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return
				} else {
					F_add_var(m, l4, int32(1717352), l4)
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return
					} else {
						v222 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						if v222 != 0 {
							F_pfree(m, v222)
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
								return
							} else {
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								if v225 != 0 {
									F_pfree(m, v225)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
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
							v225 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							if v225 != 0 {
								F_pfree(m, v225)
								mBase = m.M
								v227 = m.ExcPending
								if v227 != 0 {
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
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 float64
	_ = v439
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 float64
	_ = v453
	var v455 float32
	_ = v455
	var v458 float64
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 float64
	_ = v564
	var v568 float64
	_ = v568
	var v575 float64
	_ = v575
	var v576 float64
	_ = v576
	var v578 float64
	_ = v578
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v588 float64
	_ = v588
	var v590 float64
	_ = v590
	var v623 float32
	_ = v623
	var v625 float64
	_ = v625
	var v631 float32
	_ = v631
	var v633 float32
	_ = v633
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v825 float32
	_ = v825
	var v826 float64
	_ = v826
	var v827 float32
	_ = v827
	var v829 float64
	_ = v829
	var v835 int32
	_ = v835
	var v841 float64
	_ = v841
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v867 int32
	_ = v867
	var v869 float64
	_ = v869
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 float64
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 float64
	_ = v907
	var v925 int32
	_ = v925
	var v927 float64
	_ = v927
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v944 float64
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v956 float64
	_ = v956
	var v968 float64
	_ = v968
	var v977 int32
	_ = v977
	var v982 float64
	_ = v982
	var v994 float64
	_ = v994
	var v995 float64
	_ = v995
	var v999 float64
	_ = v999
	var v1002 float64
	_ = v1002
	var v1005 float64
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 float64
	_ = v1009
	var v1013 float64
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 float64
	_ = v1019
	var v1021 float64
	_ = v1021
	var v1027 float64
	_ = v1027
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1064 int32
	_ = v1064
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
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
	v40 = int32(65535)
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
	v72 = v5
	v78 = v5
	v80 = v5
	v81 = v5
	v83 = v5
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
	if int32(0) < v432 {
		goto L74
	} else {
		goto L75
	}
L14:
	;
	v99 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v78, v31+int32(3))
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
	v446 = v78 + int32(1)
	if v446 != l2 {
		v72 = v423
		v78 = v446
		v80 = v431
		v81 = v432
		v83 = v434
		v88 = v439
		goto L12
	} else {
		goto L71
	}
L17:
	;
	v423 = v72
	v431 = v80 + int32(1)
	v432 = v81
	v434 = v83
	v439 = v88
	goto L16
L18:
	;
	goto L19
L19:
	;
	v107 = v81 + int32(1)
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v156 = int32(0)
	if v156 < v72 {
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
		v154 = v99
		v155 = v88
		goto L20
	} else {
		goto L38
	}
L24:
	;
	v136 = base.F64_add(v88, base.F64_convert_i32_u(v134))
	v137 = F_toast_raw_datum_size(m, v99)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L33
	}
L25:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if base.Ui32((v112-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v134 = int32(6)
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v126 = int32(1)
	if v108&v126 != 0 {
		v134 = int32(base.Ui32(v108) >> (uint(v126) % 32))
		goto L24
	} else {
		goto L32
	}
L28:
	;
	v119 = int32(18)
	if v112&int32(255) == v119 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v125 = v119
	goto L31
L30:
	;
	v125 = int32(2)
	goto L31
L31:
	;
	v134 = v125
	goto L24
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v134 = int32(base.Ui32(v130) >> (uint(int32(2)) % 32))
	goto L24
L33:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v137) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v423 = v72
	v431 = v80
	v432 = v107
	v434 = v83 + int32(1)
	v439 = v136
	goto L16
L35:
	;
	goto L36
L36:
	;
	v143 = F_pg_detoast_datum(m, v99)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v154 = v143
	v155 = v136
	goto L20
L38:
	;
	v147 = F_strlen(m, v99)
	mBase = m.M
	v154 = v99
	v155 = base.F64_add(v88, base.F64_convert_i32_u(v147+int32(1)))
	goto L20
L39:
	;
	if v210 < v234 {
		goto L68
	} else {
		goto L69
	}
L40:
	;
	if v210 == v233+v72-int32(2) {
		goto L39
	} else {
		goto L64
	}
L41:
	;
	v303 = int32(3)
	v305 = v57 + v236<<(uint(v303)%32)
	v308 = v57 + v234<<(uint(v303)%32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v308-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v311
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v308-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = v315
	v319 = v234 - int32(2)
	v321 = v236
	goto L40
L42:
	;
	v248 = v192 + int32(4)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v249 + int32(1)
	if v163 == int32(0) {
		v423 = v72
		v431 = v80
		v432 = v107
		v434 = v83
		v439 = v155
		goto L16
	} else {
		goto L59
	}
L43:
	;
	v163 = v156
	v164 = v72
	goto L46
L44:
	;
	v210 = v72
	goto L45
L45:
	;
	v233 = base.B2i32(v72 < v54)
	v234 = v72 + v233
	v236 = v234 - int32(1)
	if v236 <= v210 {
		goto L39
	} else {
		goto L57
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v192 = v57 + v163<<(uint(int32(3))%32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v194 = F_FunctionCall2Coll(m, v31+int32(4), v189, v154, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L48
	}
L47:
	;
	v210 = v201
	goto L45
L48:
	;
	if v194 != 0 {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	if v163 < v164 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v197 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v201 = v164
	goto L52
L52:
	;
	v203 = v163 + int32(1)
	if v203 != v72 {
		v163 = v203
		v164 = v201
		goto L46
	} else {
		goto L56
	}
L53:
	;
	v200 = v163
	goto L55
L54:
	;
	v200 = v164
	goto L55
L55:
	;
	v201 = v200
	goto L52
L56:
	;
	goto L47
L57:
	;
	if (v72-v233+v210)&int32(1) == int32(0) {
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v319 = v236
	v321 = v234
	goto L40
L59:
	;
	v259 = v163
	goto L60
L60:
	;
	v285 = v57 + v259<<(uint(int32(3))%32)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v288 = v285 - int32(4)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v286 <= v289 {
		v423 = v72
		v431 = v80
		v432 = v107
		v434 = v83
		v439 = v155
		goto L16
	} else {
		goto L62
	}
L61:
	;
	v423 = v72
	v431 = v80
	v432 = v107
	v434 = v83
	v439 = v155
	goto L16
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+4)) = v289
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v294 = v285 - int32(8)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v286
	v299 = int32(1)
	if v299 < v259 {
		v259 = v259 - v299
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v327 = v319
	v330 = v321
	goto L65
L65:
	;
	v351 = int32(3)
	v353 = v57 + v327<<(uint(v351)%32)
	v356 = v57 + v330<<(uint(v351)%32)
	v357 = int32(16)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356-v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v359
	v361 = int32(12)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v356-v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+4)) = v363
	v366 = v327 - int32(1)
	v369 = v57 + v366<<(uint(v351)%32)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v353-v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v372
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v353-v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v376
	v379 = v327 - int32(2)
	if v210 < v379 {
		v327 = v379
		v330 = v366
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
	v412 = v57 + v210<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v412)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v154
	goto L70
L69:
	;
	goto L70
L70:
	;
	v423 = v234
	v431 = v80
	v432 = v107
	v434 = v83
	v439 = v155
	goto L16
L71:
	;
	goto L13
L72:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v623
	v625 = base.F64_promote_f32(v623)
	if base.F64_gt(v625, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L111
	} else {
		goto L112
	}
L73:
	;
	if v54 <= v423 {
		goto L97
	} else {
		goto L98
	}
L74:
	;
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v450)
	v453 = base.F64_convert_i32_s(l2)
	v455 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v431), v453))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v455
	if v45 != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	if v431 <= int32(0) {
		goto L10
	} else {
		goto L93
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v467
	if int32(0) < v423 {
		goto L84
	} else {
		goto L85
	}
L78:
	;
	v458 = base.F64_div(v439, base.F64_convert_i32_u(v432))
	if base.F64_lt(base.F64_abs(v458), float64(2.147483648e+09)) != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v465 = int32(*(*int16)(unsafe.Add(mBase, uint32(v464)+76)))
	v467 = v465
	goto L77
L81:
	;
	v462 = base.I32_trunc_f64_s(v458)
	v467 = v462
	goto L77
L82:
	;
	goto L83
L83:
	;
	v467 = int32(-2147483648)
	goto L77
L84:
	;
	v471 = int32(0)
	v477 = v471
	v478 = v471
	goto L88
L85:
	;
	goto L86
L86:
	;
	v623 = base.F32_neg(base.F32_sub(float32(1), v455))
	goto L72
L87:
	;
	if v477 != 0 {
		v557 = v477
		v558 = v478
		goto L73
	} else {
		goto L92
	}
L88:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v57+v477<<(uint(int32(3))%32))+4))
	if v504 == int32(1) {
		goto L87
	} else {
		goto L90
	}
L89:
	;
	v557 = v423
	v558 = v507
	goto L73
L90:
	;
	v507 = v478 + v504
	v509 = v477 + int32(1)
	if v509 != v423 {
		v477 = v509
		v478 = v507
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	goto L86
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v546)
	v548 = int32(0)
	if v45 == v548 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v552 = int32(*(*int16)(unsafe.Add(mBase, uint32(v551)+76)))
	v553 = v552
	goto L96
L95:
	;
	v553 = v548
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v553
	goto L10
L97:
	;
	v562 = v432 - v558
	v563 = v562 + v557
	v564 = float64(0)
	v568 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v455)))
	if base.F64_gt(v568, v564) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	if v434 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if v557 != v423 {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v623 = base.F32_convert_i32_s(v423)
	goto L72
L101:
	;
	if base.F64_lt(v584, v585) != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v584 = v564
	v585 = base.F64_convert_i32_s(v563)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v575 = base.F64_convert_i32_s(l2 - v431)
	v576 = base.F64_convert_i32_s(v563)
	v578 = base.F64_convert_i32_s(v562)
	v584 = base.F64_div(base.F64_mul(v575, v576), base.F64_add(base.F64_sub(v575, v578), base.F64_div(base.F64_mul(v575, v578), v568)))
	v585 = v576
	goto L101
L105:
	;
	v588 = v585
	goto L107
L106:
	;
	v588 = v584
	goto L107
L107:
	;
	if base.F64_gt(v588, v568) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v590 = v568
	goto L110
L109:
	;
	v590 = v588
	goto L110
L110:
	;
	v623 = base.F32_demote_f64(base.F64_floor(base.F64_add(v590, float64(0.5))))
	goto L72
L111:
	;
	v631 = base.F32_demote_f64(base.F64_div(base.F64_neg(v625), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v631
	v633 = v631
	goto L113
L112:
	;
	v633 = v623
	goto L113
L113:
	;
	if v54 <= v423 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v1064 <= int32(0) {
		goto L10
	} else {
		goto L172
	}
L115:
	;
	if v49 < v423 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	if v434 != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	if base.F32_gt(v633, float32(0)) == int32(0) {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	if v423 <= v49 {
		v1064 = v423
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L115
L120:
	;
	v641 = v49
	goto L122
L121:
	;
	v641 = v423
	goto L122
L122:
	;
	if v641 <= int32(0) {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	v645 = v641 & int32(3)
	v649 = F_palloc(m, v641<<(uint(int32(2))%32))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	v651 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v641) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v661 = v651
	v665 = int32(0)
	goto L128
L126:
	;
	v732 = v651
	goto L127
L127:
	;
	if v645 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v685 = int32(2)
	v688 = int32(3)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v57+v661<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v661<<(uint(v685)%32)))) = v691
	v694 = v661 | int32(1)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v57+v694<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v694<<(uint(v685)%32)))) = v701
	v704 = v661 | v685
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v57+v704<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v704<<(uint(v685)%32)))) = v711
	v714 = v661 | v688
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v57+v714<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v714<<(uint(v685)%32)))) = v721
	v723 = int32(4)
	v724 = v661 + v723
	v726 = v665 + v723
	if v726 != v641&int32(2147483644) {
		v661 = v724
		v665 = v726
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v732 = v724
	goto L127
L130:
	;
	goto L129
L131:
	;
	v760 = v732
	v762 = int32(0)
	goto L134
L132:
	;
	goto L133
L133:
	;
	v825 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v826 = base.F64_promote_f32(v825)
	v827 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v829 = float64(0)
	v835 = int32(0)
	v841 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v841) != 0 {
		v1041 = v641
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v57+v760<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v760<<(uint(int32(2))%32)))) = v790
	v792 = int32(1)
	v795 = v762 + v792
	if v795 != v645 {
		v760 = v760 + v792
		v762 = v795
		goto L134
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	goto L135
L137:
	;
	v1064 = v1041
	goto L114
L138:
	;
	goto L137
L139:
	;
	if base.F64_le(l3, float64(1)) != 0 {
		v1041 = v641
		goto L138
	} else {
		goto L140
	}
L140:
	;
	if base.Ui32(v641) < base.Ui32(int32(2)) {
		v956 = v829
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if base.F64_lt(v826, float64(0)) != 0 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v852 = v641 - int32(1)
	v853 = int32(3)
	v854 = v852 & v853
	if base.Ui32(v641-int32(2)) < base.Ui32(v853) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v854 == int32(0) {
		v956 = v907
		goto L141
	} else {
		goto L150
	}
L144:
	;
	v905 = int32(0)
	v907 = v829
	goto L143
L145:
	;
	goto L146
L146:
	;
	v867 = int32(0)
	v869 = v829
	v878 = v835
	goto L147
L147:
	;
	v883 = v649 + v867<<(uint(int32(2))%32)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v883)+8))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v883)+12))
	v895 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v869, base.F64_convert_i32_s(v884)), base.F64_convert_i32_s(v887)), base.F64_convert_i32_s(v890)), base.F64_convert_i32_s(v893))
	v896 = int32(4)
	v897 = v867 + v896
	v899 = v878 + v896
	if v899 != v852&int32(-4) {
		v867 = v897
		v869 = v895
		v878 = v899
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v905 = v897
	v907 = v895
	goto L143
L149:
	;
	goto L148
L150:
	;
	v925 = v905
	v927 = v907
	v935 = v835
	goto L151
L151:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v649+v925<<(uint(int32(2))%32))))
	v944 = base.F64_add(v927, base.F64_convert_i32_s(v942))
	v945 = int32(1)
	v948 = v935 + v945
	if v948 != v854 {
		v925 = v925 + v945
		v927 = v944
		v935 = v948
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v956 = v944
	goto L141
L153:
	;
	goto L152
L154:
	;
	v968 = base.F64_mul(l3, base.F64_neg(v826))
	goto L156
L155:
	;
	v968 = v826
	goto L156
L156:
	;
	v977 = v641
	v982 = v956
	goto L157
L157:
	;
	v994 = float64(1)
	v995 = float64(0)
	v999 = base.F64_sub(base.F64_sub(v994, base.F64_div(v982, v841)), base.F64_promote_f32(v827))
	if base.F64_lt(v999, v995) != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1041 = int32(0)
	goto L138
L159:
	;
	v1002 = v995
	goto L161
L160:
	;
	v1002 = v999
	goto L161
L161:
	;
	if base.F64_gt(v1002, float64(1)) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1005 = v994
	goto L164
L163:
	;
	v1005 = v1002
	goto L164
L164:
	;
	v1007 = v977 - int32(1)
	v1009 = base.F64_sub(v968, base.F64_convert_i32_u(v1007))
	if base.F64_gt(v1009, float64(1)) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1013 = base.F64_div(v1005, v1009)
	goto L167
L166:
	;
	v1013 = v1005
	goto L167
L167:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v649+v1007<<(uint(int32(2))%32))))
	v1019 = base.F64_convert_i32_s(v1018)
	v1021 = base.F64_div(base.F64_mul(l3, v1019), v841)
	v1027 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v841), base.F64_mul(base.F64_mul(v1021, v841), base.F64_sub(l3, v1021))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v1013, v841), base.F64_add(v1027, v1027)), float64(0.5)), v1019) != 0 {
		v1041 = v977
		goto L138
	} else {
		goto L168
	}
L168:
	;
	if v1007 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v649-int32(8)+v977<<(uint(int32(2))%32))))
	v977 = v1007
	v982 = base.F64_sub(v982, base.F64_convert_i32_s(v1036))
	goto L157
L170:
	;
	goto L171
L171:
	;
	goto L158
L172:
	;
	v1089 = int32(4480304)
	v1090 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1092
	v1095 = v1064 << (uint(int32(2)) % 32)
	v1096 = F_palloc(m, v1095)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	v1098 = F_palloc(m, v1095)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	v1104 = int32(0)
	goto L175
L175:
	;
	v1129 = v1104 << (uint(int32(2)) % 32)
	v1133 = v57 + v1104<<(uint(int32(3))%32)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135)+78)))
	v1137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1135)+76)))
	v1138 = F_datumCopy(m, v1134, v1136, v1137)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L7
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1090
	v1152 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1152)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1098
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v1064
	goto L10
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1096+v1129))) = v1138
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v1129+v1098))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1142), v453))
	v1148 = v1104 + int32(1)
	if v1148 != v1064 {
		v1104 = v1148
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
}
func F_compute_scalar_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 float64
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 float64
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 float64
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 float64
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 float64
	_ = v199
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 float64
	_ = v244
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var __phi288 int32
	_ = __phi288
	var v291 int32
	_ = v291
	var __phi291 int32
	_ = __phi291
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v369 int32
	_ = v369
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v439 float64
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 float64
	_ = v446
	var v448 float32
	_ = v448
	var v451 float64
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v470 float64
	_ = v470
	var v474 float64
	_ = v474
	var v481 float64
	_ = v481
	var v482 float64
	_ = v482
	var v486 float64
	_ = v486
	var v492 float64
	_ = v492
	var v493 float64
	_ = v493
	var v496 float64
	_ = v496
	var v498 float64
	_ = v498
	var v508 float32
	_ = v508
	var v510 float64
	_ = v510
	var v516 float32
	_ = v516
	var v518 float32
	_ = v518
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
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
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v728 float32
	_ = v728
	var v729 float64
	_ = v729
	var v730 float32
	_ = v730
	var v732 float64
	_ = v732
	var v738 int32
	_ = v738
	var v744 float64
	_ = v744
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v770 int32
	_ = v770
	var v772 float64
	_ = v772
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 float64
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v810 float64
	_ = v810
	var v828 int32
	_ = v828
	var v830 float64
	_ = v830
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v847 float64
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v859 float64
	_ = v859
	var v871 float64
	_ = v871
	var v880 int32
	_ = v880
	var v885 float64
	_ = v885
	var v897 float64
	_ = v897
	var v898 float64
	_ = v898
	var v902 float64
	_ = v902
	var v905 float64
	_ = v905
	var v908 float64
	_ = v908
	var v910 int32
	_ = v910
	var v912 float64
	_ = v912
	var v916 float64
	_ = v916
	var v921 int32
	_ = v921
	var v922 float64
	_ = v922
	var v924 float64
	_ = v924
	var v930 float64
	_ = v930
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v962 int32
	_ = v962
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1339 int32
	_ = v1339
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1474 int32
	_ = v1474
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 float64
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 float64
	_ = v1511
	var v1513 float64
	_ = v1513
	var v1515 float64
	_ = v1515
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1549 int32
	_ = v1549
	var v1554 float32
	_ = v1554
	var v1557 float64
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	v5 = int32(0)
	v26 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(48)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+78)))
	if v38 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+76)))
	v44 = int32(65535)
	v49 = base.B2i32(v41&v44 == v44)
	v50 = base.B2i32(v41 < int32(0))
	goto L3
L2:
	;
	v49 = v5
	v50 = v5
	goto L3
L3:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v55 = F_palloc(m, l2<<(uint(int32(3))%32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v59 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v63 = F_palloc(m, v51<<(uint(int32(3))%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = v65
	v67 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+36)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v35)+28)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v35)+20)) = v67
	v74 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v76
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v65)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	F_PrepareSortSupportFromOrderingOp(m, v80, v35+int32(12))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
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
	m.G0 = v35 + int32(48)
	return
L10:
	;
	v91 = v5
	v96 = v5
	v100 = v5
	v106 = v5
	v107 = v5
	v112 = v26
	goto L11
L11:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	if int32(0) < v194 {
		goto L39
	} else {
		goto L40
	}
L13:
	;
	v124 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v91, v35+int32(4))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v126 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v201 = v91 + int32(1)
	if v201 != l2 {
		v91 = v201
		v96 = v194
		v100 = v196
		v106 = v197
		v107 = v198
		v112 = v199
		goto L11
	} else {
		goto L38
	}
L16:
	;
	v194 = v96
	v196 = v100 + int32(1)
	v197 = v106
	v198 = v107
	v199 = v112
	goto L15
L17:
	;
	goto L18
L18:
	;
	v132 = v106 + int32(1)
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v183 = v55 + v96<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v59+v96<<(uint(int32(2))%32)))) = v96
	v194 = v96 + int32(1)
	v196 = v100
	v197 = v132
	v198 = v107
	v199 = v180
	goto L15
L20:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v133 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	if v50 == int32(0) {
		v177 = v124
		v180 = v112
		goto L19
	} else {
		goto L37
	}
L23:
	;
	v161 = base.F64_add(v112, base.F64_convert_i32_u(v159))
	v162 = F_toast_raw_datum_size(m, v124)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	if base.Ui32((v137-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v159 = int32(6)
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v151 = int32(1)
	if v133&v151 != 0 {
		v159 = int32(base.Ui32(v133) >> (uint(v151) % 32))
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v144 = int32(18)
	if v137&int32(255) == v144 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v150 = v144
	goto L30
L29:
	;
	v150 = int32(2)
	goto L30
L30:
	;
	v159 = v150
	goto L23
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v159 = int32(base.Ui32(v155) >> (uint(int32(2)) % 32))
	goto L23
L32:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v162) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v194 = v96
	v196 = v100
	v197 = v132
	v198 = v107 + int32(1)
	v199 = v161
	goto L15
L34:
	;
	goto L35
L35:
	;
	v168 = F_pg_detoast_datum(m, v124)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v177 = v168
	v180 = v161
	goto L19
L37:
	;
	v172 = F_strlen(m, v124)
	mBase = m.M
	v177 = v124
	v180 = base.F64_add(v112, base.F64_convert_i32_u(v172+int32(1)))
	goto L19
L38:
	;
	goto L12
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v35 + int32(12)
	F_qsort_interruptible(m, v55, v194, int32(8), int32(509), v35+int32(4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if int32(0) < v197 {
		goto L230
	} else {
		goto L231
	}
L42:
	;
	v219 = int32(0)
	v232 = v5
	v233 = v5
	v234 = v5
	v239 = v5
	v244 = v26
	goto L43
L43:
	;
	v251 = v232 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v55+v219<<(uint(int32(3))%32))+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v59+v256<<(uint(int32(2))%32))))
	if v256 != v262 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v443)
	v446 = base.F64_convert_i32_s(l2)
	v448 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v196), v446))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v448
	if v50 != 0 {
		goto L64
	} else {
		goto L65
	}
L45:
	;
	v421 = v233
	v422 = v234
	v427 = v239
	v438 = v251
	goto L47
L46:
	;
	if v251 < int32(2) {
		v388 = v233
		v394 = v239
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v439 = base.F64_add(base.F64_mul(base.F64_convert_i32_u(v219), base.F64_convert_i32_s(v256)), v244)
	v441 = v219 + int32(1)
	if v441 != v194 {
		v219 = v441
		v232 = v438
		v233 = v421
		v234 = v422
		v239 = v427
		v244 = v439
		goto L43
	} else {
		goto L62
	}
L48:
	;
	v421 = v388
	v422 = v234 + int32(1)
	v427 = v394
	v438 = int32(0)
	goto L47
L49:
	;
	v269 = v239 + int32(1)
	v270 = base.B2i32(v233 < v51)
	if v270 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v63-int32(8)+v233<<(uint(int32(3))%32))))
	if v251 <= v276 {
		v388 = v233
		v394 = v269
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v278 = v270 + v233
	v280 = v278 - int32(1)
	if v280 <= int32(0) {
		v340 = v280
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v369 = v63 + v340<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v219 - v232
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v251
	v388 = v278
	v394 = v269
	goto L48
L55:
	;
	__phi288 = v280
	__phi291 = v278
	v288 = __phi288
	v291 = __phi291
	goto L56
L56:
	;
	v317 = v63 + v291<<(uint(int32(3))%32)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317-int32(16))))
	if v251 <= v320 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v340 = int32(0)
	goto L54
L58:
	;
	v340 = v288
	goto L54
L59:
	;
	goto L60
L60:
	;
	v324 = v63 + v288<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v320
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+4)) = v328
	v330 = int32(1)
	if v330 < v288 {
		__phi288 = v288 - v330
		__phi291 = v288
		v288 = __phi288
		v291 = __phi291
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	goto L44
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v460
	if v427 == int32(0) {
		v508 = base.F32_neg(base.F32_sub(float32(1), v448))
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v451 = base.F64_div(v199, base.F64_convert_i32_s(v197))
	if base.F64_lt(base.F64_abs(v451), float64(2.147483648e+09)) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v458 = int32(*(*int16)(unsafe.Add(mBase, uint32(v457)+76)))
	v460 = v458
	goto L63
L67:
	;
	v455 = base.I32_trunc_f64_s(v451)
	v460 = v455
	goto L63
L68:
	;
	goto L69
L69:
	;
	v460 = int32(-2147483648)
	goto L63
L70:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v508
	v510 = base.F64_promote_f32(v508)
	if base.F64_gt(v510, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L85
	} else {
		goto L86
	}
L71:
	;
	if v198 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v469 = v422 + v198
	v470 = float64(0)
	v474 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v448)))
	if base.F64_gt(v474, v470) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if v422 != v427 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v508 = base.F32_convert_i32_s(v427)
	goto L70
L75:
	;
	if base.F64_lt(v492, v493) != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v492 = v470
	v493 = base.F64_convert_i32_s(v469)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v481 = base.F64_convert_i32_s(l2 - v196)
	v482 = base.F64_convert_i32_s(v469)
	v486 = base.F64_convert_i32_s(v198 - v427 + v422)
	v492 = base.F64_div(base.F64_mul(v481, v482), base.F64_add(base.F64_sub(v481, v486), base.F64_div(base.F64_mul(v481, v486), v474)))
	v493 = v482
	goto L75
L79:
	;
	v496 = v493
	goto L81
L80:
	;
	v496 = v492
	goto L81
L81:
	;
	if base.F64_gt(v496, v474) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v498 = v474
	goto L84
L83:
	;
	v498 = v496
	goto L84
L84:
	;
	v508 = base.F32_demote_f64(base.F64_floor(base.F64_add(v498, float64(0.5))))
	goto L70
L85:
	;
	v516 = base.F32_demote_f64(base.F64_div(base.F64_neg(v510), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v516
	v518 = v516
	goto L87
L86:
	;
	v518 = v508
	goto L87
L87:
	;
	if v421 != v422 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v1114 = v422 - v1080
	if v51 < v1114 {
		goto L154
	} else {
		goto L155
	}
L89:
	;
	v993 = int32(0)
	if v962 <= v993 {
		v1080 = v962
		v1111 = v993
		goto L88
	} else {
		goto L147
	}
L90:
	;
	if v51 < v421 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	if v198 != 0 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	if base.F32_gt(v518, float32(0)) == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	if v422 <= v51 {
		v962 = v422
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v528 = v51
	goto L97
L96:
	;
	v528 = v421
	goto L97
L97:
	;
	if v528 <= int32(0) {
		v1080 = v528
		v1111 = int32(0)
		goto L88
	} else {
		goto L98
	}
L98:
	;
	v532 = v528 & int32(3)
	v536 = F_palloc(m, v528<<(uint(int32(2))%32))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v538 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v528) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v548 = v538
	v552 = int32(0)
	goto L103
L101:
	;
	v623 = v538
	goto L102
L102:
	;
	if v532 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v576 = int32(2)
	v579 = int32(3)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v63+v548<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v548<<(uint(v576)%32)))) = v582
	v585 = v548 | int32(1)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v63+v585<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v585<<(uint(v576)%32)))) = v592
	v595 = v548 | v576
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v63+v595<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v595<<(uint(v576)%32)))) = v602
	v605 = v548 | v579
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v63+v605<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v605<<(uint(v576)%32)))) = v612
	v614 = int32(4)
	v615 = v548 + v614
	v617 = v552 + v614
	if v617 != v528&int32(2147483644) {
		v548 = v615
		v552 = v617
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v623 = v615
	goto L102
L105:
	;
	goto L104
L106:
	;
	v655 = v623
	v658 = int32(0)
	goto L109
L107:
	;
	goto L108
L108:
	;
	v728 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v729 = base.F64_promote_f32(v728)
	v730 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v732 = float64(0)
	v738 = int32(0)
	v744 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v744) != 0 {
		v944 = v528
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v63+v655<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v655<<(uint(int32(2))%32)))) = v689
	v691 = int32(1)
	v694 = v658 + v691
	if v694 != v532 {
		v655 = v655 + v691
		v658 = v694
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L108
L111:
	;
	goto L110
L112:
	;
	v962 = v944
	goto L89
L113:
	;
	goto L112
L114:
	;
	if base.F64_le(l3, float64(1)) != 0 {
		v944 = v528
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v528) < base.Ui32(int32(2)) {
		v859 = v732
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if base.F64_lt(v729, float64(0)) != 0 {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v755 = v528 - int32(1)
	v756 = int32(3)
	v757 = v755 & v756
	if base.Ui32(v528-int32(2)) < base.Ui32(v756) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v757 == int32(0) {
		v859 = v810
		goto L116
	} else {
		goto L125
	}
L119:
	;
	v808 = int32(0)
	v810 = v732
	goto L118
L120:
	;
	goto L121
L121:
	;
	v770 = int32(0)
	v772 = v732
	v781 = v738
	goto L122
L122:
	;
	v786 = v536 + v770<<(uint(int32(2))%32)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v786)+4))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v786)+8))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v786)+12))
	v798 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v772, base.F64_convert_i32_s(v787)), base.F64_convert_i32_s(v790)), base.F64_convert_i32_s(v793)), base.F64_convert_i32_s(v796))
	v799 = int32(4)
	v800 = v770 + v799
	v802 = v781 + v799
	if v802 != v755&int32(-4) {
		v770 = v800
		v772 = v798
		v781 = v802
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v808 = v800
	v810 = v798
	goto L118
L124:
	;
	goto L123
L125:
	;
	v828 = v808
	v830 = v810
	v838 = v738
	goto L126
L126:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v536+v828<<(uint(int32(2))%32))))
	v847 = base.F64_add(v830, base.F64_convert_i32_s(v845))
	v848 = int32(1)
	v851 = v838 + v848
	if v851 != v757 {
		v828 = v828 + v848
		v830 = v847
		v838 = v851
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v859 = v847
	goto L116
L128:
	;
	goto L127
L129:
	;
	v871 = base.F64_mul(l3, base.F64_neg(v729))
	goto L131
L130:
	;
	v871 = v729
	goto L131
L131:
	;
	v880 = v528
	v885 = v859
	goto L132
L132:
	;
	v897 = float64(1)
	v898 = float64(0)
	v902 = base.F64_sub(base.F64_sub(v897, base.F64_div(v885, v744)), base.F64_promote_f32(v730))
	if base.F64_lt(v902, v898) != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v944 = int32(0)
	goto L113
L134:
	;
	v905 = v898
	goto L136
L135:
	;
	v905 = v902
	goto L136
L136:
	;
	if base.F64_gt(v905, float64(1)) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v908 = v897
	goto L139
L138:
	;
	v908 = v905
	goto L139
L139:
	;
	v910 = v880 - int32(1)
	v912 = base.F64_sub(v871, base.F64_convert_i32_u(v910))
	if base.F64_gt(v912, float64(1)) != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v916 = base.F64_div(v908, v912)
	goto L142
L141:
	;
	v916 = v908
	goto L142
L142:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v536+v910<<(uint(int32(2))%32))))
	v922 = base.F64_convert_i32_s(v921)
	v924 = base.F64_div(base.F64_mul(l3, v922), v744)
	v930 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v744), base.F64_mul(base.F64_mul(v924, v744), base.F64_sub(l3, v924))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v916, v744), base.F64_add(v930, v930)), float64(0.5)), v922) != 0 {
		v944 = v880
		goto L113
	} else {
		goto L143
	}
L143:
	;
	if v910 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v536-int32(8)+v880<<(uint(int32(2))%32))))
	v880 = v910
	v885 = base.F64_sub(v885, base.F64_convert_i32_s(v939))
	goto L132
L145:
	;
	goto L146
L146:
	;
	goto L133
L147:
	;
	v997 = int32(4480304)
	v998 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1000
	v1003 = v962 << (uint(int32(2)) % 32)
	v1004 = F_palloc(m, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v1006 = F_palloc(m, v1003)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v1012 = int32(0)
	goto L150
L150:
	;
	v1041 = v1012 << (uint(int32(2)) % 32)
	v1043 = int32(3)
	v1045 = v63 + v1012<<(uint(v1043)%32)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1046<<(uint(v1043)%32))))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+78)))
	v1053 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1051)+76)))
	v1054 = F_datumCopy(m, v1050, v1052, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L4
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v998
	v1068 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1068)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1070
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1006
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1004
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v962
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v962
	v1080 = v962
	v1111 = v1068
	goto L88
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1004+v1041))) = v1054
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	*(*float32)(unsafe.Add(mBase, uint32(v1041+v1006))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1058), v446))
	v1064 = v1012 + int32(1)
	if v1064 != v962 {
		v1012 = v1064
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v1116 = v51 + int32(1)
	goto L156
L155:
	;
	v1116 = v1114
	goto L156
L156:
	;
	if int32(2) <= v1116 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1119 = int32(0)
	F_qsort_interruptible(m, v63, v1080, int32(8), int32(510), v1119)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	v1474 = v1111
	goto L159
L159:
	;
	if v194 == int32(1) {
		goto L9
	} else {
		goto L228
	}
L160:
	;
	if v1111 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1125 = int32(0)
	v1131 = v1125
	v1134 = v1125
	v1135 = v1119
	goto L164
L162:
	;
	v1339 = v194
	goto L163
L163:
	;
	v1364 = int32(4480304)
	v1365 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1367
	v1369 = int32(1)
	v1370 = v1339 - v1369
	v1372 = v1116 - v1369
	v1373 = base.I32_div_s(v1370, v1372)
	if v1116 <= v1369 {
		goto L217
	} else {
		goto L218
	}
L164:
	;
	if v1080 <= v1135 {
		v1169 = v194
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v1339 = v1327
	goto L163
L166:
	;
	if v1330 < v194 {
		v1131 = v1330
		v1134 = v1327
		v1135 = v1328
		goto L164
	} else {
		goto L216
	}
L167:
	;
	v1171 = int32(3)
	v1173 = v55 + v1134<<(uint(v1171)%32)
	v1176 = v55 + v1131<<(uint(v1171)%32)
	v1177 = v1169 - v1131
	v1179 = v1177 << (uint(v1171) % 32)
	if v1173 == v1176 {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v1162 = v63 + v1135<<(uint(int32(3))%32)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	if v1131 < v1163 {
		v1169 = v1163
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1162)))
	v1327 = v1134
	v1328 = v1135 + int32(1)
	v1330 = v1167 + v1163
	goto L166
L170:
	;
	v1327 = v1177 + v1134
	v1328 = v1135
	v1330 = v1169
	goto L166
L171:
	;
	goto L170
L172:
	;
	v1183 = v1173 + v1179
	if base.Ui32(v1176-v1183) <= base.Ui32(int32(0)-v1179<<(uint(int32(1))%32)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1190 = F___memcpy(m, v1173, v1176, v1179)
	mBase = m.M
	goto L170
L174:
	;
	goto L175
L175:
	;
	v1193 = (v1173 ^ v1176) & int32(3)
	if base.Ui32(v1173) < base.Ui32(v1176) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	if v1295 == int32(0) {
		goto L171
	} else {
		goto L212
	}
L177:
	;
	if base.Ui32(v1273) <= base.Ui32(int32(3)) {
		v1294 = v1272
		v1295 = v1273
		v1296 = v1274
		goto L176
	} else {
		goto L208
	}
L178:
	;
	if v1193 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	if v1193 != 0 {
		v1255 = v1179
		goto L191
	} else {
		goto L192
	}
L181:
	;
	v1294 = v1176
	v1295 = v1179
	v1296 = v1173
	goto L176
L182:
	;
	goto L183
L183:
	;
	if v1173&int32(3) == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1272 = v1176
	v1273 = v1179
	v1274 = v1173
	goto L177
L185:
	;
	goto L186
L186:
	;
	v1200 = v1176
	v1201 = v1179
	v1202 = v1173
	goto L187
L187:
	;
	if v1201 == int32(0) {
		goto L171
	} else {
		goto L189
	}
L188:
	;
	v1272 = v1209
	v1273 = v1211
	v1274 = v1213
	goto L177
L189:
	;
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1202))) = uint8(v1206)
	v1208 = int32(1)
	v1209 = v1200 + v1208
	v1211 = v1201 - v1208
	v1213 = v1202 + v1208
	if v1213&int32(3) != 0 {
		v1200 = v1209
		v1201 = v1211
		v1202 = v1213
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	if v1255 == int32(0) {
		goto L171
	} else {
		goto L204
	}
L192:
	;
	if v1183&int32(3) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1220 = v1179
	goto L196
L194:
	;
	v1235 = v1179
	goto L195
L195:
	;
	if base.Ui32(v1235) <= base.Ui32(int32(3)) {
		v1255 = v1235
		goto L191
	} else {
		goto L200
	}
L196:
	;
	if v1220 == int32(0) {
		goto L171
	} else {
		goto L198
	}
L197:
	;
	v1235 = v1226
	goto L195
L198:
	;
	v1226 = v1220 - int32(1)
	v1227 = v1173 + v1226
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176+v1226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1227))) = uint8(v1229)
	if v1227&int32(3) != 0 {
		v1220 = v1226
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v1242 = v1235
	goto L201
L201:
	;
	v1246 = v1242 - int32(4)
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1176+v1246)))
	*(*int32)(unsafe.Add(mBase, uint32(v1173+v1246))) = v1249
	if base.Ui32(int32(3)) < base.Ui32(v1246) {
		v1242 = v1246
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v1255 = v1246
	goto L191
L203:
	;
	goto L202
L204:
	;
	v1262 = v1255
	goto L205
L205:
	;
	v1266 = v1262 - int32(1)
	v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176+v1266))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1173+v1266))) = uint8(v1269)
	if v1266 != 0 {
		v1262 = v1266
		goto L205
	} else {
		goto L207
	}
L206:
	;
	goto L171
L207:
	;
	goto L206
L208:
	;
	v1279 = v1272
	v1280 = v1273
	v1281 = v1274
	goto L209
L209:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1279)))
	*(*int32)(unsafe.Add(mBase, uint32(v1281))) = v1283
	v1285 = int32(4)
	v1286 = v1279 + v1285
	v1288 = v1281 + v1285
	v1290 = v1280 - v1285
	if base.Ui32(int32(3)) < base.Ui32(v1290) {
		v1279 = v1286
		v1280 = v1290
		v1281 = v1288
		goto L209
	} else {
		goto L211
	}
L210:
	;
	v1294 = v1286
	v1295 = v1290
	v1296 = v1288
	goto L176
L211:
	;
	goto L210
L212:
	;
	v1301 = v1294
	v1302 = v1295
	v1303 = v1296
	goto L213
L213:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1303))) = uint8(v1305)
	v1307 = int32(1)
	v1312 = v1302 - v1307
	if v1312 != 0 {
		v1301 = v1301 + v1307
		v1302 = v1312
		v1303 = v1303 + v1307
		goto L213
	} else {
		goto L215
	}
L214:
	;
	goto L171
L215:
	;
	goto L214
L216:
	;
	goto L165
L217:
	;
	v1379 = v1369
	goto L219
L218:
	;
	v1379 = v1116
	goto L219
L219:
	;
	v1382 = F_palloc(m, v1116<<(uint(int32(2))%32))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	v1384 = int32(0)
	v1388 = v1384
	v1391 = v1384
	v1393 = v1384
	goto L221
L221:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1393<<(uint(int32(3))%32))))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426)+78)))
	v1428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1426)+76)))
	v1429 = F_datumCopy(m, v1425, v1427, v1428)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L4
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1365
	v1444 = int32(1)
	v1447 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1111<<(uint(v1444)%32))+52)) = uint16(v1447)
	v1451 = l0 + v1111<<(uint(v1447)%32)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1451-int32(-64)))) = v1454
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+164)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+84)) = v1456
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+144)) = v1116
	v1474 = v1111 + v1444
	goto L159
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1382+v1391<<(uint(int32(2))%32)))) = v1429
	v1432 = v1388 + (v1370 - v1373*v1372)
	v1433 = base.B2i32(v1372 <= v1432)
	if v1372 <= v1432 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1437 = v1372
	goto L226
L225:
	;
	v1437 = int32(0)
	goto L226
L226:
	;
	v1440 = v1391 + int32(1)
	if v1440 != v1379 {
		v1388 = v1432 - v1437
		v1391 = v1440
		v1393 = v1433 + (v1393 + v1373)
		goto L221
	} else {
		goto L227
	}
L227:
	;
	goto L222
L228:
	;
	v1496 = int32(4480304)
	v1497 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1499
	v1502 = F_palloc(m, int32(4))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1497
	v1506 = base.F64_convert_i32_u(v194)
	v1508 = int32(1)
	v1511 = base.F64_mul(v1506, base.F64_convert_i32_u(v194-v1508))
	v1513 = base.F64_mul(v1511, float64(0.5))
	v1515 = base.F64_mul(v1513, base.F64_neg(v1513))
	*(*float32)(unsafe.Add(mBase, uint32(v1502))) = base.F32_demote_f64(base.F64_div(base.F64_add(base.F64_mul(v1506, v439), v1515), base.F64_add(base.F64_mul(v1506, base.F64_div(base.F64_mul(v1511, base.F64_convert_i32_s(v194<<(uint(v1508)%32)-v1508)), float64(6))), v1515)))
	v1533 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1474<<(uint(v1508)%32))+52)) = uint16(v1533)
	v1537 = l0 + v1474<<(uint(int32(2))%32)
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1537-int32(-64)))) = v1540
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+124)) = v1502
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+84)) = v1542
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+104)) = v1508
	goto L9
L230:
	;
	v1549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1549)
	v1554 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v196), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v1554
	if v50 != 0 {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	goto L232
L232:
	;
	if v196 <= int32(0) {
		goto L9
	} else {
		goto L240
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1566
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v1554))
	goto L9
L234:
	;
	v1557 = base.F64_div(v199, base.F64_convert_i32_u(v197))
	if base.F64_lt(base.F64_abs(v1557), float64(2.147483648e+09)) != 0 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1564 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1563)+76)))
	v1566 = v1564
	goto L233
L237:
	;
	v1561 = base.I32_trunc_f64_s(v1557)
	v1566 = v1561
	goto L233
L238:
	;
	goto L239
L239:
	;
	v1566 = int32(-2147483648)
	goto L233
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v1576 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1576)
	v1578 = int32(0)
	if v50 == v1578 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1581)+76)))
	v1583 = v1582
	goto L243
L242:
	;
	v1583 = v1578
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1583
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
														F_errmsg(m, int32(105805), int32(0))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494963), int32(1076), int32(311031))
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
															F_errmsg(m, int32(105805), int32(0))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(494963), int32(1076), int32(311031))
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
																F_errmsg(m, int32(60116), int32(0))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(494963), int32(1081), int32(311031))
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
																	F_errmsg(m, int32(60116), int32(0))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(494963), int32(1081), int32(311031))
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
																		v75 = int32(4480304)
																		v76 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																		v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																		*(*int32)(unsafe.Add(mBase, _consts[0])) = v79
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
																						v99 = int32(4480304)
																						v100 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																						*(*int32)(unsafe.Add(mBase, _consts[0])) = v79
																						v109 = *(*int32)(unsafe.Add(mBase, _consts[326]))
																						v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																						mBase = m.M
																						v111 = m.ExcPending
																						if v111 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
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
																									*(*int32)(unsafe.Add(mBase, _consts[0])) = v76
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
																v72 = F_pstrdup(m, int32(4101))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	v74 = v72
																	v75 = int32(4480304)
																	v76 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																	v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v79
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
																					v99 = int32(4480304)
																					v100 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																					*(*int32)(unsafe.Add(mBase, _consts[0])) = v79
																					v109 = *(*int32)(unsafe.Add(mBase, _consts[326]))
																					v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																					mBase = m.M
																					v111 = m.ExcPending
																					if v111 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
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
																								*(*int32)(unsafe.Add(mBase, _consts[0])) = v76
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v79 int32
	_ = v79
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
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
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var __phi277 int32
	_ = __phi277
	var v285 int32
	_ = v285
	var __phi285 int32
	_ = __phi285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v313 int32
	_ = v313
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v593 int32
	_ = v593
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1046 int32
	_ = v1046
	var v1053 int32
	_ = v1053
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1079 int32
	_ = v1079
	var v1091 int32
	_ = v1091
	var v1104 int32
	_ = v1104
	var v1118 int32
	_ = v1118
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
		v1104 = v9
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
	if base.Ui32(v1104) < base.Ui32(l1) {
		goto L299
	} else {
		goto L300
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
	v49 = l4
	v50 = v28
	goto L11
L11:
	;
	v55 = l2 + v47
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 == int32(0) {
		v1104 = v43
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v1104 = v1079
	goto L6
L13:
	;
	if base.I32_extend8_s(v56) < int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if l4 != int32(1) {
		v140 = v49
		v141 = v50
		goto L33
	} else {
		goto L34
	}
L15:
	;
	if v56&int32(224) == int32(192) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v117 = v56
	goto L17
L17:
	;
	v118 = int32(1)
	if base.Ui32(v117) < base.Ui32(int32(128)) {
		v131 = v117
		v132 = v118
		v133 = v118
		goto L14
	} else {
		goto L28
	}
L18:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v55))))
	v117 = v113&int32(63) | v110
	goto L17
L19:
	;
	v110 = v56 << (uint(int32(6)) % 32) & int32(1984)
	v111 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v56&int32(240) == int32(224) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v110 = v56<<(uint(int32(12))%32)&int32(61440) | v79&int32(63)<<(uint(int32(6))%32)
	v111 = int32(2)
	goto L18
L23:
	;
	goto L24
L24:
	;
	if v56&int32(248) != int32(240) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v131 = int32(-1)
	v132 = int32(0)
	v133 = int32(4)
	goto L14
L26:
	;
	goto L27
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v98 = int32(63)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	v110 = v56<<(uint(int32(18))%32)&int32(1835008) | v97&v98<<(uint(int32(12))%32) | v103&v98<<(uint(int32(6))%32)
	v111 = int32(3)
	goto L18
L28:
	;
	v122 = int32(0)
	if base.Ui32(v117) < base.Ui32(int32(2048)) {
		v131 = v117
		v132 = v122
		v133 = int32(2)
		goto L14
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v117) < base.Ui32(int32(65536)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v130 = int32(3)
	goto L32
L31:
	;
	v130 = int32(4)
	goto L32
L32:
	;
	v131 = v117
	v132 = v122
	v133 = v130
	goto L14
L33:
	;
	if v132 != 0 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	if v47 != v50 {
		v140 = int32(0)
		v141 = v50
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v140 = v33
	v141 = v138
	goto L33
L37:
	;
	v1091 = v47 + v133
	if l3 < int32(0) {
		v43 = v1079
		v47 = v1091
		v49 = v140
		v50 = v141
		goto L11
	} else {
		goto L297
	}
L38:
	;
	v968 = v43
	v970 = int32(0)
	goto L273
L39:
	;
	v874 = l0 + v43
	if base.Ui32(v850) <= base.Ui32(int32(2047)) {
		goto L267
	} else {
		goto L268
	}
L40:
	;
	v870 = v43 + int32(1)
	if base.Ui32(l1) < base.Ui32(v870) {
		goto L264
	} else {
		goto L265
	}
L41:
	;
	v864 = v43 + v133
	if base.Ui32(l1) < base.Ui32(v864) {
		goto L257
	} else {
		goto L258
	}
L42:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	if base.Ui32(v850) < base.Ui32(int32(128)) {
		goto L40
	} else {
		goto L249
	}
L43:
	;
	v142 = int32(2)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(v142)%32))+uint32(_consts[1014])))
	v849 = v146 + v131<<(uint(v142)%32) + int32(4)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v152 = int32(0)
	if base.Ui32(v131) < base.Ui32(int32(1416)) {
		v239 = v131
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v246 == int32(0) {
		goto L41
	} else {
		goto L102
	}
L47:
	;
	goto L46
L48:
	;
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239<<(uint(int32(1))%32))+uint32(_consts[1015]))))
	v246 = v244
	goto L47
L49:
	;
	if base.Ui32(v131) <= base.Ui32(int32(43967)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(v131) <= base.Ui32(int32(8580)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v131) < base.Ui32(int32(64256)) {
		v246 = v152
		goto L47
	} else {
		goto L77
	}
L53:
	;
	if base.Ui32(v131-int32(4256)) <= base.Ui32(int32(95)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v131) < base.Ui32(int32(9398)) {
		v246 = v152
		goto L47
	} else {
		goto L64
	}
L56:
	;
	v239 = v131 - int32(2840)
	goto L48
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v131) < base.Ui32(int32(5024)) {
		v246 = v152
		goto L47
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v131) <= base.Ui32(int32(5117)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v239 = v131 - int32(3512)
	goto L48
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v131) < base.Ui32(int32(7296)) {
		v246 = v152
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v239 = v131 - int32(5690)
	goto L48
L64:
	;
	if base.Ui32(v131) <= base.Ui32(int32(11565)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if base.Ui32(v131) <= base.Ui32(int32(9449)) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if base.Ui32(v131) < base.Ui32(int32(42560)) {
		v246 = v152
		goto L47
	} else {
		goto L72
	}
L68:
	;
	v239 = v131 - int32(6507)
	goto L48
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v131) < base.Ui32(int32(11264)) {
		v246 = v152
		goto L47
	} else {
		goto L71
	}
L71:
	;
	v239 = v131 - int32(8321)
	goto L48
L72:
	;
	if base.Ui32(v131) <= base.Ui32(int32(42998)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v239 = v131 - int32(39315)
	goto L48
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v131) < base.Ui32(int32(43859)) {
		v246 = v152
		goto L47
	} else {
		goto L76
	}
L76:
	;
	v239 = v131 - int32(40175)
	goto L48
L77:
	;
	if base.Ui32(v131) <= base.Ui32(int32(68997)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if base.Ui32(v131) <= base.Ui32(int32(65370)) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(v131) < base.Ui32(int32(71840)) {
		v246 = v152
		goto L47
	} else {
		goto L93
	}
L81:
	;
	if base.Ui32(v131) <= base.Ui32(int32(64279)) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(v131) < base.Ui32(int32(66560)) {
		v246 = v152
		goto L47
	} else {
		goto L88
	}
L84:
	;
	v239 = v131 - int32(60463)
	goto L48
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v131) < base.Ui32(int32(65313)) {
		v246 = v152
		goto L47
	} else {
		goto L87
	}
L87:
	;
	v239 = v131 - int32(61496)
	goto L48
L88:
	;
	if base.Ui32(v131) <= base.Ui32(int32(67004)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v239 = v131 - int32(62685)
	goto L48
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(v131) < base.Ui32(int32(68736)) {
		v246 = v152
		goto L47
	} else {
		goto L92
	}
L92:
	;
	v239 = v131 - int32(64416)
	goto L48
L93:
	;
	if base.Ui32(v131) <= base.Ui32(int32(93823)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if base.Ui32(v131) <= base.Ui32(int32(71903)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(int32(67)) < base.Ui32(v131-int32(125184)) {
		v246 = v152
		goto L47
	} else {
		goto L101
	}
L97:
	;
	v239 = v131 - int32(67258)
	goto L48
L98:
	;
	goto L99
L99:
	;
	if base.Ui32(v131) < base.Ui32(int32(93760)) {
		v246 = v152
		goto L47
	} else {
		goto L100
	}
L100:
	;
	v239 = v131 - int32(89114)
	goto L48
L101:
	;
	v239 = v131 - int32(120474)
	goto L48
L102:
	;
	if l5 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v820 = int32(2)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(v820)%32))+uint32(_consts[1014])))
	v849 = v824 + v246<<(uint(v820)%32)
	goto L42
L104:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_consts[1016]))))
	if v253 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v257 = v253 * int32(52)
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+uint32(_consts[1017]))))
	switch v260 {
	case 0:
		goto L38
	case 1:
		goto L106
	default:
		goto L103
	}
L106:
	;
	if v47 == int32(0) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v264 = v47 - int32(1)
	if v264 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if l3 == v47 {
		goto L38
	} else {
		goto L178
	}
L109:
	;
	__phi277 = v264
	__phi285 = v47
	v277 = __phi277
	v285 = __phi285
	goto L110
L110:
	;
	v288 = l2 + v277
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v290 = base.I32_extend8_s(v289)
	if int32(-64) <= v290 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if base.Ui32(int32(127)) < base.Ui32(v351) {
		goto L144
	} else {
		goto L145
	}
L112:
	;
	goto L111
L113:
	;
	if int32(0) <= v290 {
		v351 = v289
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	if int32(0) < v277 {
		__phi277 = v277 - int32(1)
		__phi285 = v277
		v277 = __phi277
		v285 = __phi285
		goto L110
	} else {
		goto L141
	}
L116:
	;
	if base.Ui32(int32(127)) < base.Ui32(v351) {
		goto L127
	} else {
		goto L128
	}
L117:
	;
	if v289&int32(224) == int32(192) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v288))))
	v351 = v346&int32(63) | v343
	goto L116
L119:
	;
	v343 = v289 << (uint(int32(6)) % 32) & int32(1984)
	v344 = int32(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	if v289&int32(240) == int32(224) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v285))))
	v343 = v289<<(uint(int32(12))%32)&int32(61440) | v313&int32(63)<<(uint(int32(6))%32)
	v344 = int32(2)
	goto L118
L123:
	;
	goto L124
L124:
	;
	if v289&int32(248) != int32(240) {
		v351 = int32(-1)
		goto L116
	} else {
		goto L125
	}
L125:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v285))))
	v331 = int32(63)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+2)))
	v343 = v289<<(uint(int32(18))%32)&int32(1835008) | v330&v331<<(uint(int32(12))%32) | v336&v331<<(uint(int32(6))%32)
	v344 = int32(3)
	goto L118
L126:
	;
	if v400 == int32(0) {
		goto L112
	} else {
		goto L140
	}
L127:
	;
	v360 = int32(0)
	v361 = int32(505)
	goto L130
L128:
	;
	goto L129
L129:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351<<(uint(int32(1))%32))+uint32(_consts[501]))))
	v400 = int32(base.Ui32(v390&int32(16)) >> (uint(int32(4)) % 32))
	goto L126
L130:
	;
	v366 = base.I32_div_s(v360+v361, int32(2))
	v368 = v366 << (uint(int32(3)) % 32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_consts[1018])))
	if base.Ui32(v371) < base.Ui32(v351) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v400 = int32(0)
	goto L126
L132:
	;
	if v382 <= v383 {
		v360 = v382
		v361 = v383
		goto L130
	} else {
		goto L139
	}
L133:
	;
	v382 = v366 + int32(1)
	v383 = v361
	goto L132
L134:
	;
	goto L135
L135:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_consts[1019])))
	if base.Ui32(v377) <= base.Ui32(v351) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v400 = int32(1)
	goto L126
L137:
	;
	goto L138
L138:
	;
	v382 = v360
	v383 = v366 - int32(1)
	goto L132
L139:
	;
	goto L131
L140:
	;
	goto L115
L141:
	;
	goto L108
L142:
	;
	if v518 == int32(0) {
		goto L103
	} else {
		goto L177
	}
L143:
	;
	v459 = int32(689)
	v460 = int32(0)
	goto L157
L144:
	;
	v417 = int32(3367)
	v418 = int32(0)
	goto L147
L145:
	;
	goto L146
L146:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351<<(uint(int32(1))%32))+uint32(_consts[501]))))
	v518 = int32(base.Ui32(v449&int32(8)) >> (uint(int32(3)) % 32))
	goto L142
L147:
	;
	v423 = base.I32_div_s(v417+v418, int32(2))
	v425 = v423 * int32(12)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v425)+uint32(_consts[505])))
	if base.Ui32(v428) < base.Ui32(v351) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+uint32(_consts[504]))))
	if v441 != int32(3) {
		goto L143
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	if v439 <= v438 {
		v417 = v438
		v418 = v439
		goto L147
	} else {
		goto L155
	}
L151:
	;
	v438 = v417
	v439 = v423 + int32(1)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v425)+uint32(_consts[506])))
	if base.Ui32(v434) <= base.Ui32(v351) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v438 = v423 - int32(1)
	v439 = v418
	goto L150
L155:
	;
	goto L143
L156:
	;
	v518 = int32(1)
	goto L142
L157:
	;
	v465 = base.I32_div_s(v459+v460, int32(2))
	v467 = v465 << (uint(int32(3)) % 32)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_consts[1020])))
	if base.Ui32(v470) < base.Ui32(v351) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v487 = int32(655)
	v488 = int32(0)
	goto L167
L159:
	;
	if v482 <= v481 {
		v459 = v481
		v460 = v482
		goto L157
	} else {
		goto L166
	}
L160:
	;
	v481 = v459
	v482 = v465 + int32(1)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_consts[1021])))
	if base.Ui32(v476) <= base.Ui32(v351) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v518 = int32(1)
	goto L142
L164:
	;
	goto L165
L165:
	;
	v481 = v465 - int32(1)
	v482 = v460
	goto L159
L166:
	;
	goto L158
L167:
	;
	v493 = base.I32_div_s(v487+v488, int32(2))
	v495 = v493 << (uint(int32(3)) % 32)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v495)+uint32(_consts[1022])))
	if base.Ui32(v498) < base.Ui32(v351) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v518 = int32(0)
	goto L142
L169:
	;
	if v510 <= v509 {
		v487 = v509
		v488 = v510
		goto L167
	} else {
		goto L176
	}
L170:
	;
	v509 = v487
	v510 = v493 + int32(1)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v495)+uint32(_consts[1023])))
	if base.Ui32(v504) <= base.Ui32(v351) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v518 = int32(1)
	goto L142
L174:
	;
	goto L175
L175:
	;
	v509 = v493 - int32(1)
	v510 = v488
	goto L169
L176:
	;
	goto L168
L177:
	;
	goto L108
L178:
	;
	v544 = v47 + int32(1)
	if base.Ui32(l3) <= base.Ui32(v544) {
		goto L38
	} else {
		goto L179
	}
L179:
	;
	v554 = v544
	goto L180
L180:
	;
	v567 = l2 + v554
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	if v568 == int32(0) {
		goto L38
	} else {
		goto L182
	}
L181:
	;
	if base.Ui32(int32(127)) < base.Ui32(v630) {
		goto L215
	} else {
		goto L216
	}
L182:
	;
	v571 = base.I32_extend8_s(v568)
	if int32(-64) <= v571 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	goto L181
L184:
	;
	if int32(0) <= v571 {
		v630 = v568
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	v685 = v554 + int32(1)
	if v685 != l3 {
		v554 = v685
		goto L180
	} else {
		goto L212
	}
L187:
	;
	if base.Ui32(int32(127)) < base.Ui32(v630) {
		goto L198
	} else {
		goto L199
	}
L188:
	;
	if v568&int32(224) == int32(192) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623+v567))))
	v630 = v625&int32(63) | v622
	goto L187
L190:
	;
	v622 = v568 << (uint(int32(6)) % 32) & int32(1984)
	v623 = int32(1)
	goto L189
L191:
	;
	goto L192
L192:
	;
	if v568&int32(240) == int32(224) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+1)))
	v622 = v568<<(uint(int32(12))%32)&int32(61440) | v593&int32(63)<<(uint(int32(6))%32)
	v623 = int32(2)
	goto L189
L194:
	;
	goto L195
L195:
	;
	if v568&int32(248) != int32(240) {
		v630 = int32(-1)
		goto L187
	} else {
		goto L196
	}
L196:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+1)))
	v610 = int32(63)
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	v622 = v568<<(uint(int32(18))%32)&int32(1835008) | v609&v610<<(uint(int32(12))%32) | v615&v610<<(uint(int32(6))%32)
	v623 = int32(3)
	goto L189
L197:
	;
	if v679 == int32(0) {
		goto L183
	} else {
		goto L211
	}
L198:
	;
	v639 = int32(0)
	v640 = int32(505)
	goto L201
L199:
	;
	goto L200
L200:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630<<(uint(int32(1))%32))+uint32(_consts[501]))))
	v679 = int32(base.Ui32(v669&int32(16)) >> (uint(int32(4)) % 32))
	goto L197
L201:
	;
	v645 = base.I32_div_s(v639+v640, int32(2))
	v647 = v645 << (uint(int32(3)) % 32)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v647)+uint32(_consts[1018])))
	if base.Ui32(v650) < base.Ui32(v630) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v679 = int32(0)
	goto L197
L203:
	;
	if v661 <= v662 {
		v639 = v661
		v640 = v662
		goto L201
	} else {
		goto L210
	}
L204:
	;
	v661 = v645 + int32(1)
	v662 = v640
	goto L203
L205:
	;
	goto L206
L206:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v647)+uint32(_consts[1019])))
	if base.Ui32(v656) <= base.Ui32(v630) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v679 = int32(1)
	goto L197
L208:
	;
	goto L209
L209:
	;
	v661 = v639
	v662 = v645 - int32(1)
	goto L203
L210:
	;
	goto L202
L211:
	;
	goto L186
L212:
	;
	goto L38
L213:
	;
	if v796 == int32(0) {
		goto L38
	} else {
		goto L248
	}
L214:
	;
	v737 = int32(689)
	v738 = int32(0)
	goto L228
L215:
	;
	v695 = int32(3367)
	v696 = int32(0)
	goto L218
L216:
	;
	goto L217
L217:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630<<(uint(int32(1))%32))+uint32(_consts[501]))))
	v796 = int32(base.Ui32(v727&int32(8)) >> (uint(int32(3)) % 32))
	goto L213
L218:
	;
	v701 = base.I32_div_s(v695+v696, int32(2))
	v703 = v701 * int32(12)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703)+uint32(_consts[505])))
	if base.Ui32(v706) < base.Ui32(v630) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+uint32(_consts[504]))))
	if v719 != int32(3) {
		goto L214
	} else {
		goto L227
	}
L220:
	;
	goto L219
L221:
	;
	if v717 <= v716 {
		v695 = v716
		v696 = v717
		goto L218
	} else {
		goto L226
	}
L222:
	;
	v716 = v695
	v717 = v701 + int32(1)
	goto L221
L223:
	;
	goto L224
L224:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v703)+uint32(_consts[506])))
	if base.Ui32(v712) <= base.Ui32(v630) {
		goto L220
	} else {
		goto L225
	}
L225:
	;
	v716 = v701 - int32(1)
	v717 = v696
	goto L221
L226:
	;
	goto L214
L227:
	;
	v796 = int32(1)
	goto L213
L228:
	;
	v743 = base.I32_div_s(v737+v738, int32(2))
	v745 = v743 << (uint(int32(3)) % 32)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[1020])))
	if base.Ui32(v748) < base.Ui32(v630) {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	v765 = int32(655)
	v766 = int32(0)
	goto L238
L230:
	;
	if v760 <= v759 {
		v737 = v759
		v738 = v760
		goto L228
	} else {
		goto L237
	}
L231:
	;
	v759 = v737
	v760 = v743 + int32(1)
	goto L230
L232:
	;
	goto L233
L233:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[1021])))
	if base.Ui32(v754) <= base.Ui32(v630) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v796 = int32(1)
	goto L213
L235:
	;
	goto L236
L236:
	;
	v759 = v743 - int32(1)
	v760 = v738
	goto L230
L237:
	;
	goto L229
L238:
	;
	v771 = base.I32_div_s(v765+v766, int32(2))
	v773 = v771 << (uint(int32(3)) % 32)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v773)+uint32(_consts[1022])))
	if base.Ui32(v776) < base.Ui32(v630) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v796 = int32(0)
	goto L213
L240:
	;
	if v788 <= v787 {
		v765 = v787
		v766 = v788
		goto L238
	} else {
		goto L247
	}
L241:
	;
	v787 = v765
	v788 = v771 + int32(1)
	goto L240
L242:
	;
	goto L243
L243:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v773)+uint32(_consts[1023])))
	if base.Ui32(v782) <= base.Ui32(v630) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v796 = int32(1)
	goto L213
L245:
	;
	goto L246
L246:
	;
	v787 = v771 - int32(1)
	v788 = v766
	goto L240
L247:
	;
	goto L239
L248:
	;
	goto L103
L249:
	;
	if base.Ui32(v850) < base.Ui32(int32(65536)) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v858 = int32(3)
	goto L252
L251:
	;
	v858 = int32(4)
	goto L252
L252:
	;
	if base.Ui32(v850) < base.Ui32(int32(2048)) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v861 = int32(2)
	goto L255
L254:
	;
	v861 = v858
	goto L255
L255:
	;
	v862 = v861 + v43
	if base.Ui32(v862) <= base.Ui32(l1) {
		goto L39
	} else {
		goto L256
	}
L256:
	;
	v1079 = v862
	goto L37
L257:
	;
	v1079 = v864
	goto L37
L258:
	;
	goto L259
L259:
	;
	if v133 != 0 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1079 = v864
	goto L37
L261:
	;
	v867 = F__emscripten_memcpy_bulkmem(m, l0+v43, v55, v133)
	mBase = m.M
	goto L263
L262:
	;
	goto L263
L263:
	;
	goto L260
L264:
	;
	v1079 = v870
	goto L37
L265:
	;
	goto L266
L266:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v43))) = uint8(v850)
	v1079 = v870
	goto L37
L267:
	;
	v880 = v850&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)) = uint8(v880)
	v885 = int32(base.Ui32(v850)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v885)
	v1079 = v862
	goto L37
L268:
	;
	goto L269
L269:
	;
	if base.Ui32(v850) <= base.Ui32(int32(65535)) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v889 = int32(63)
	v891 = int32(128)
	v892 = v850&v889 | v891
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+2)) = uint8(v892)
	v897 = int32(base.Ui32(v850)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v897)
	v904 = int32(base.Ui32(v850)>>(uint(int32(6))%32))&v889 | v891
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)) = uint8(v904)
	v1079 = v862
	goto L37
L271:
	;
	goto L272
L272:
	;
	v906 = int32(63)
	v908 = int32(128)
	v909 = v850&v906 | v908
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+3)) = uint8(v909)
	v916 = int32(base.Ui32(v850)>>(uint(int32(6))%32))&v906 | v908
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+2)) = uint8(v916)
	v923 = int32(base.Ui32(v850)>>(uint(int32(12))%32))&v906 | v908
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)) = uint8(v923)
	v930 = int32(base.Ui32(v850)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v930)
	v1079 = v862
	goto L37
L273:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v257+v140*int32(12)+int32(1888740)+v970<<(uint(int32(2))%32))))
	if v983 == int32(0) {
		v1079 = v968
		goto L37
	} else {
		goto L275
	}
L274:
	;
	v1079 = v1065
	goto L37
L275:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v983) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v1067 = v970 + int32(1)
	if v1067 != int32(3) {
		v968 = v1065
		v970 = v1067
		goto L273
	} else {
		goto L296
	}
L277:
	;
	v1065 = v997
	goto L276
L278:
	;
	if base.Ui32(v983) <= base.Ui32(int32(65535)) {
		goto L293
	} else {
		goto L294
	}
L279:
	;
	if base.Ui32(v983) < base.Ui32(int32(65536)) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	v1013 = v968 + int32(1)
	if base.Ui32(l1) < base.Ui32(v1013) {
		goto L290
	} else {
		goto L291
	}
L282:
	;
	v993 = int32(3)
	goto L284
L283:
	;
	v993 = int32(4)
	goto L284
L284:
	;
	if base.Ui32(v983) < base.Ui32(int32(2048)) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v996 = int32(2)
	goto L287
L286:
	;
	v996 = v993
	goto L287
L287:
	;
	v997 = v996 + v968
	if base.Ui32(l1) < base.Ui32(v997) {
		goto L277
	} else {
		goto L288
	}
L288:
	;
	v999 = l0 + v968
	if base.Ui32(int32(2047)) < base.Ui32(v983) {
		goto L278
	} else {
		goto L289
	}
L289:
	;
	v1005 = v983&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+1)) = uint8(v1005)
	v1010 = int32(base.Ui32(v983)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1010)
	goto L277
L290:
	;
	v1065 = v1013
	goto L276
L291:
	;
	goto L292
L292:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v968))) = uint8(v983)
	v1065 = v1013
	goto L276
L293:
	;
	v1019 = int32(63)
	v1021 = int32(128)
	v1022 = v983&v1019 | v1021
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+2)) = uint8(v1022)
	v1027 = int32(base.Ui32(v983)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1027)
	v1034 = int32(base.Ui32(v983)>>(uint(int32(6))%32))&v1019 | v1021
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+1)) = uint8(v1034)
	goto L277
L294:
	;
	goto L295
L295:
	;
	v1036 = int32(63)
	v1038 = int32(128)
	v1039 = v983&v1036 | v1038
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+3)) = uint8(v1039)
	v1046 = int32(base.Ui32(v983)>>(uint(int32(6))%32))&v1036 | v1038
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+2)) = uint8(v1046)
	v1053 = int32(base.Ui32(v983)>>(uint(int32(12))%32))&v1036 | v1038
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+1)) = uint8(v1053)
	v1060 = int32(base.Ui32(v983)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1060)
	goto L277
L296:
	;
	goto L274
L297:
	;
	if base.Ui32(v1091) < base.Ui32(l3) {
		v43 = v1079
		v47 = v1091
		v49 = v140
		v50 = v141
		goto L11
	} else {
		goto L298
	}
L298:
	;
	goto L12
L299:
	;
	v1118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1104))) = uint8(v1118)
	goto L301
L300:
	;
	goto L301
L301:
	;
	return v1104
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
	F_MemoryContextReset(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(4480304)
		v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
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
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
					v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
					v32 = v30 + int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v32
					v37 = *(*int32)(unsafe.Add(mBase, _consts[31]))
					if v37 == int32(0) {
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
						if v41 != int32(1) {
						} else {
							v44 = int32(4474964)
							v46 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							v47 = int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[7])) = v46 + v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v50 + v47
							*(*int64)(unsafe.Add(mBase, uint32(v37+int32(16))+232)) = v32
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v58 + v47
							v64 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							*(*int32)(unsafe.Add(mBase, _consts[7])) = v64 - v47
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
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
				v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v32 = v30 + int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v32
				v37 = *(*int32)(unsafe.Add(mBase, _consts[31]))
				if v37 == int32(0) {
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
					if v41 != int32(1) {
					} else {
						v44 = int32(4474964)
						v46 = *(*int32)(unsafe.Add(mBase, _consts[7]))
						v47 = int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[7])) = v46 + v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v50 + v47
						*(*int64)(unsafe.Add(mBase, uint32(v37+int32(16))+232)) = v32
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v58 + v47
						v64 = *(*int32)(unsafe.Add(mBase, _consts[7]))
						*(*int32)(unsafe.Add(mBase, _consts[7])) = v64 - v47
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_ArrayGetNItems(m, v5, l0+int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if int32(0) < v8 {
			v17 = v8<<(uint(int32(2))%32) + int32(24)
			v18 = F_palloc0(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v8
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(23)
				*(*int64)(unsafe.Add(mBase, uint32(v18)+4)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v17 << (uint(int32(2)) % 32)
				v34 = v18
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				v43 = v34
				v44 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 != 0 {
					v54 = v46
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = (v47<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v57 = v8 << (uint(int32(2)) % 32)
				if v57 != 0 {
					v58 = F__emscripten_memcpy_bulkmem(m, v43+v44, v54+l0, v57)
					mBase = m.M
				} else {
				}
				return v43
			}
		} else {
			v31 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				if v33 != 0 {
					v43 = v31
					v44 = v33
				} else {
					v34 = v31
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v43 = v34
					v44 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 != 0 {
					v54 = v46
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = (v47<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v57 = v8 << (uint(int32(2)) % 32)
				if v57 != 0 {
					v58 = F__emscripten_memcpy_bulkmem(m, v43+v44, v54+l0, v57)
					mBase = m.M
				} else {
				}
				return v43
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
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v6 = F_palloc0(m, int32(40))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(277)
		v13 = l0 + int32(8)
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v25 = F_list_copy(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if v28 != 0 {
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
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v34 != 0 {
						v39 = F__emscripten_memcpy_bulkmem(m, v35, v38, v34)
						mBase = m.M
					} else {
					}
					return v6
				}
			} else {
				return v6
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l2&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(l2&int64(9223090561878065152))>>(uint(v10)%64)))|base.I32_wrap_i64(int64(base.Ui64(l3)>>(uint(v10)%64)))&int32(32768))<<(uint(v10)%64)
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
	var v62 float64
	_ = v62
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
	var v81 float64
	_ = v81
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
	v88 = *(*float64)(unsafe.Add(mBase, _consts[380]))
	v89 = float64(0)
	v90 = base.F64_add(v85, v89)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v90, base.F64_add(base.F64_mul(v86, base.F64_add(v81, v88)), v89))
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l1
	if v17 == int32(0) {
		v62 = v7
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
	v81 = v73
	v85 = v74
	goto L1
L5:
	;
	v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v81 = base.F64_add(v62, v67)
	v85 = base.F64_add(v66, v69)
	goto L1
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= int32(0) {
		v62 = v7
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
	v62 = v54
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
	var v107 float64
	_ = v107
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
	var v150 float64
	_ = v150
	var v154 float64
	_ = v154
	var v157 float64
	_ = v157
	var v162 int32
	_ = v162
	var v165 float64
	_ = v165
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
	v116 = *(*float64)(unsafe.Add(mBase, _consts[380]))
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
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
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
	v107 = v94
	v110 = v95
	goto L6
L10:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v107 = base.F64_add(v89, v90)
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
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _consts[381])))
	if v127 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v157 = v120
	goto L19
L19:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _consts[382])))
	v165 = base.F64_add(base.F64_add(v107, float64(0)), v121)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v162 ^ int32(1)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(v31, base.F64_convert_i32_u(v30)), base.F64_add(v165, v157))
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
	if base.F64_gt(v142, v141) != 0 {
		v154 = v141
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
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v154
	v157 = base.F64_div(v120, v139)
	goto L19
L27:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v142)&int64(9223372036854775807)) {
		v154 = v141
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v150 = float64(1)
	if base.F64_le(v142, v150) != 0 {
		v154 = v150
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v154 = base.F64_nearest(v142)
	goto L26
}
func F_count_nulls(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
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
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v214 int32
	_ = v214
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == v4 {
		v24 = v4
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
		if v16 == int32(0) {
			v24 = v4
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v19 != int32(15) {
				v24 = v4
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+13)))
				v24 = v22
			}
		}
	}
	if v24&int32(1) == int32(0) {
		v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v29 <= int32(0) {
			v192 = v4
			v195 = v29
		} else {
			v33 = v29 & int32(3)
			v35 = l0 + int32(24)
			v36 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v29) {
				v41 = v36
				v44 = v4
				v49 = v4
				for {
					v54 = v35 + v41<<(uint(int32(3))%32)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+8)))
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)))
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+24)))
					v62 = v44 + v55 + v57 + v59 + v61
					v63 = int32(4)
					v64 = v41 + v63
					v66 = v49 + v63
					if v66 != v29&int32(32764) {
						v41 = v64
						v44 = v62
						v49 = v66
						continue
					} else {
						break
					}
					break
				}
				v68 = v64
				v71 = v62
			} else {
				v68 = v36
				v71 = v4
			}
			if v33 == int32(0) {
				v192 = v71
				v195 = v29
			} else {
				v81 = v68
				v84 = v71
				v88 = v4
				for {
					v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v81<<(uint(int32(3))%32)))))
					v96 = v84 + v95
					v97 = int32(1)
					v100 = v88 + v97
					if v100 != v33 {
						v81 = v81 + v97
						v84 = v96
						v88 = v100
						continue
					} else {
						break
					}
					break
				}
				v192 = v96
				v195 = v29
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v195
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v192
		v214 = int32(1)
		return v214
	} else {
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v103 != 0 {
			v214 = int32(0)
			return v214
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v105 = F_pg_detoast_datum(m, v104)
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
				v111 = v105 + int32(16)
				v112 = F_ArrayGetNItems(m, v109, v111)
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int32(0)
				} else {
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
					if v114 == int32(0) {
						v192 = v4
						v195 = v112
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						v120 = v111 + v117<<(uint(int32(3))%32)
						if v120 == int32(0) {
							v192 = v4
							v195 = v112
						} else {
							if v112 <= int32(0) {
								v192 = v4
								v195 = v112
							} else {
								v126 = int32(1)
								if v112 == v126 {
									v174 = v4
									v176 = v120
									v179 = int32(1)
								} else {
									v133 = int32(1)
									v136 = v4
									v137 = v120
									v140 = v4
									for {
										v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
										v149 = int32(1)
										v151 = v133 << (uint(v149) % 32)
										v153 = base.B2i32(v151 == int32(256))
										if v151 == int32(256) {
											v154 = v149
										} else {
											v154 = v151
										}
										v155 = v153 + v137
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
										v160 = v136 + base.B2i32(v133&v144 == int32(0)) + base.B2i32(v154&v156 == int32(0))
										v161 = int32(1)
										v163 = v154 << (uint(v161) % 32)
										v165 = base.B2i32(v163 == int32(256))
										if v163 == int32(256) {
											v166 = v161
										} else {
											v166 = v163
										}
										v167 = v155 + v165
										v169 = v140 + int32(2)
										if v169 != v112&int32(2147483646) {
											v133 = v166
											v136 = v160
											v137 = v167
											v140 = v169
											continue
										} else {
											break
										}
										break
									}
									v174 = v160
									v176 = v167
									v179 = v166
								}
								if v112&v126 == int32(0) {
									v192 = v174
									v195 = v112
								} else {
									v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
									v192 = v174 + base.B2i32(v179&v184 == int32(0))
									v195 = v112
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v195
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v192
					v214 = int32(1)
					return v214
				}
			}
		}
	}
}
func F_create_drop_transactional_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
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
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v3 = l2
	v4 = l3
	v9 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	v14 = F_MemoryContextAlloc(m, v12, int32(28))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[765]))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v10 {
				v39 = v17
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)) = uint32(v46)
				v49 = v39 + int32(8)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v50 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v49
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v49
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
				v61 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v64 + int32(1)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[76]))
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
					v32 = int32(4463536)
					v33 = *(*int32)(unsafe.Add(mBase, _consts[765]))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
					*(*int32)(unsafe.Add(mBase, _consts[765])) = v23
					v39 = v23
					*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
					v46 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)) = uint32(v46)
					v49 = v39 + int32(8)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					if v50 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v49
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v49
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
					v61 = v14 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v61
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v64 + int32(1)
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[76]))
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
				v32 = int32(4463536)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[765]))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
				*(*int32)(unsafe.Add(mBase, _consts[765])) = v23
				v39 = v23
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)) = uint32(v46)
				v49 = v39 + int32(8)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v50 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v49
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v49
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
				v61 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v64 + int32(1)
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
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
	v21 = int32(0)
	goto L7
L4:
	;
	goto L5
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v61 != 0 {
		goto L18
	} else {
		goto L19
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
	if v27 <= v21 {
		v33 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = v29 + v21<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v51
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+24)) = uint16(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+26)) = uint8(v55)
	v21 = v21 + int32(1)
	goto L7
L13:
	;
	goto L6
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v36 <= v21 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v33 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v43 = v40 + v21<<(uint(int32(2))%32)
	if v43 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	return v6
L21:
	;
	F_errmsg_internal(m, int32(169816), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(491722), int32(372), int32(280455))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
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
	var v22 int32
	_ = v22
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
	var v42 int32
	_ = v42
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	v3 = l2
	v4 = int32(0)
	v8 = int32(16)
	if l1 == v4 {
		v42 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = F_palloc(m, v42)
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
		v42 = v8
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
	v22 = v4
	v24 = v8
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v22<<(uint(int32(2))%32))))
	v30 = F_strlen(m, v29)
	mBase = m.M
	v32 = int32(1)
	v33 = v30 + v24 + v32
	v35 = v22 + v32
	if v35 != v17 {
		v22 = v35
		v24 = v33
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v42 = v33
	goto L1
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
	v52 = l0
	goto L13
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v96
	if l1 == int32(0) {
		goto L28
	} else {
		goto L29
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(0)
	return v44
L29:
	;
	goto L30
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v106 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = int32(0)
	v117 = v44 + int32(16)
	goto L34
L32:
	;
	goto L33
L33:
	;
	return v44
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v115<<(uint(int32(2))%32))))
	if (v123^v117)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	goto L33
L36:
	;
	v198 = F_strlen(m, v123)
	mBase = m.M
	v200 = int32(1)
	v203 = v115 + v200
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v203 < v204 {
		v115 = v203
		v117 = v117 + v198 + v200
		goto L34
	} else {
		goto L57
	}
L37:
	;
	goto L36
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v177)
	if v177&int32(255) == int32(0) {
		goto L37
	} else {
		goto L53
	}
L39:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v176 = v123
	v177 = v129
	v178 = v117
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v123&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v133 = v123
	v135 = v117
	goto L45
L43:
	;
	v147 = v123
	v149 = v117
	goto L44
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v154 = int32(-2139062144)
	if (int32(16843008)-v151|v151)&v154 != v154 {
		v176 = v147
		v177 = v151
		v178 = v149
		goto L38
	} else {
		goto L49
	}
L45:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v136)
	if v136 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	v147 = v143
	v149 = v141
	goto L44
L47:
	;
	v140 = int32(1)
	v141 = v135 + v140
	v143 = v133 + v140
	if v143&int32(3) != 0 {
		v133 = v143
		v135 = v141
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v159 = v147
	v160 = v151
	v161 = v149
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v160
	v163 = int32(4)
	v164 = v161 + v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v167 = v159 + v163
	v171 = int32(-2139062144)
	if (v165|(int32(16843008)-v165))&v171 == v171 {
		v159 = v167
		v160 = v165
		v161 = v164
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v176 = v167
	v177 = v165
	v178 = v164
	goto L38
L52:
	;
	goto L51
L53:
	;
	v185 = v176
	v187 = v178
	goto L54
L54:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)) = uint8(v188)
	v190 = int32(1)
	if v188 != 0 {
		v185 = v185 + v190
		v187 = v187 + v190
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L37
L56:
	;
	goto L55
L57:
	;
	goto L35
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	v3 = l2
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v12
		v90 = v11
		v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
		if v96 != 0 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
			*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			if v105 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			if v112 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v116 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
			v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
			if v123 < int32(0) {
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v128 = v126 - int32(97)
				if base.Ui32(int32(17)) < base.Ui32(v128) {
				} else {
					if int32(1)<<(uint(v128)%32)&int32(163841) == int32(0) {
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						if v137 != 0 {
						} else {
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
							v145 = v140 + v123*int32(24) + int32(12)
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
							if v146 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
								v149 = v148
							} else {
								v149 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
							*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
						}
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
					v53 = v16 << (uint(int32(1)) % 32)
					v54 = v34
					v57 = int32(1024)
					if base.Ui32(v57) <= base.Ui32(v53) {
						v60 = v57
					} else {
						v60 = v53
					}
					v64 = v60*int32(40) | int32(8)
					v66 = F_palloc_extended(m, v64, int32(2))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						if v66 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(101)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							if v74 != 0 {
								v76 = v74
							} else {
								v76 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v76
							v90 = int32(0)
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+136))
							*(*int32)(unsafe.Add(mBase, uint32(v68)+136)) = v79 + v64
							*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v60
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66))) = v83
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v66
							v90 = v66 + int32(8)
						}
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
						if v96 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
							*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
							*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							if v105 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
							v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							if v112 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							v116 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
							v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
							if v123 < int32(0) {
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v128 = v126 - int32(97)
								if base.Ui32(int32(17)) < base.Ui32(v128) {
								} else {
									if int32(1)<<(uint(v128)%32)&int32(163841) == int32(0) {
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
										if v137 != 0 {
										} else {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
											v145 = v140 + v123*int32(24) + int32(12)
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											if v146 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
												v149 = v148
											} else {
												v149 = int32(0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
											*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
										}
									}
								}
							}
						}
						return
					}
				} else {
					v40 = v34
					v41 = v35
					*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
					if v46 != 0 {
						v48 = v46
					} else {
						v48 = int32(19)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v48
					v90 = int32(0)
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
					if v96 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v105 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v112 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v116 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
						v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
						if v123 < int32(0) {
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
							v128 = v126 - int32(97)
							if base.Ui32(int32(17)) < base.Ui32(v128) {
							} else {
								if int32(1)<<(uint(v128)%32)&int32(163841) == int32(0) {
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
									if v137 != 0 {
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
										v145 = v140 + v123*int32(24) + int32(12)
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										if v146 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											v149 = v148
										} else {
											v149 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
										*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
									}
								}
							}
						}
					}
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
				v90 = v14 + v15*int32(40) + int32(8)
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
				if v96 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v105 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v112 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v116 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
					v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
					if v123 < int32(0) {
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						v128 = v126 - int32(97)
						if base.Ui32(int32(17)) < base.Ui32(v128) {
						} else {
							if int32(1)<<(uint(v128)%32)&int32(163841) == int32(0) {
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v137 != 0 {
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
									v145 = v140 + v123*int32(24) + int32(12)
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
									if v146 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										v149 = v148
									} else {
										v149 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
									*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
								}
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
				v40 = v27
				v41 = v29
				*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
				if v46 != 0 {
					v48 = v46
				} else {
					v48 = int32(19)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v48
				v90 = int32(0)
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
				if v96 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v105 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v112 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v116 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
					v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
					if v123 < int32(0) {
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						v128 = v126 - int32(97)
						if base.Ui32(int32(17)) < base.Ui32(v128) {
						} else {
							if int32(1)<<(uint(v128)%32)&int32(163841) == int32(0) {
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v137 != 0 {
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
									v145 = v140 + v123*int32(24) + int32(12)
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
									if v146 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										v149 = v148
									} else {
										v149 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
									*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
								}
							}
						}
					}
				}
				return
			} else {
				v53 = int32(64)
				v54 = v27
				v57 = int32(1024)
				if base.Ui32(v57) <= base.Ui32(v53) {
					v60 = v57
				} else {
					v60 = v53
				}
				v64 = v60*int32(40) | int32(8)
				v66 = F_palloc_extended(m, v64, int32(2))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					if v66 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(101)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
						if v74 != 0 {
							v76 = v74
						} else {
							v76 = int32(12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v76
						v90 = int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+136))
						*(*int32)(unsafe.Add(mBase, uint32(v68)+136)) = v79 + v64
						*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v60
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v66))) = v83
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v66
						v90 = v66 + int32(8)
					}
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
					if v96 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v90))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = l3
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v101
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v105 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v90
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v108
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v112 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v90
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v90
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v116 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v115 + v116
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v119 + v116
						v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
						if v123 < int32(0) {
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
							v128 = v126 - int32(97)
							if base.Ui32(int32(17)) < base.Ui32(v128) {
							} else {
								if int32(1)<<(uint(v128)%32)&int32(163841) == int32(0) {
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
									if v137 != 0 {
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
										v145 = v140 + v123*int32(24) + int32(12)
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										if v146 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v146)+36)) = v90
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											v149 = v148
										} else {
											v149 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v149
										*(*int32)(unsafe.Add(mBase, uint32(v145))) = v90
									}
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
	var v17 int32
	_ = v17
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
			v17 = F__emscripten_memcpy_bulkmem(m, v8+int32(4), l0, v5)
			mBase = m.M
		} else {
		}
		return v8
	}
}
