package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	v128 = l4 + l1 - int32(base.Ui32(v123)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)
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
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l1 - int32(base.Ui32(v206)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)
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
	if base.Ui32(v217+v218) <= base.Ui32(int32(base.Ui32(v45)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)) {
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
	v308 = F_strcpy(m, v306+l4, l0+int32(base.Ui32(v300)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0))
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
	v355 = int32(base.Ui32(v351)>>(uint(int32(1))%32)) & int32(_a_F_CheckAffix_1)
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
	F_errmsg_internal(m, int32(_a_F_CheckAffix_2), v319)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L106
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_CheckAffix_3), int32(245), int32(_a_F_CheckAffix_4))
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
func F_CheckDim_3(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if int32(0) < l0 {
		if base.Ui32(int32(_a_F_CheckDim_3_0)) <= base.Ui32(l0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_CheckDim_3_1)
					F_errmsg(m, int32(_a_F_CheckDim_3_2), v5)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckDim_3_3), int32(105), int32(_a_F_CheckDim_3_4))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_CheckDim_3_5), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckDim_3_3), int32(100), int32(_a_F_CheckDim_3_4))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
	F_errstart_cold(m, int32(21), int32(_a_F_check_labels_0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L22
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_check_labels_0))
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
	F_errmsg(m, int32(_a_F_check_labels_1), v9)
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
	F_errfinish(m, int32(_a_F_check_labels_2), int32(3886), int32(_a_F_check_labels_3))
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
	F_errmsg(m, int32(_a_F_check_labels_4), v9+int32(16))
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
	F_errfinish(m, int32(_a_F_check_labels_2), int32(3893), int32(_a_F_check_labels_3))
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
	v40 = F_MakeSingleTupleTableSlot(m, v38, int32(_a_F_check_publications_1))
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
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_check_usermap[0]))
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
	F_errmsg(m, int32(_a_F_check_usermap_0), v16)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2988), int32(_a_F_check_usermap_2))
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
	F_errmsg(m, int32(_a_F_check_usermap_3), v16-int32(-64))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2861), int32(_a_F_check_usermap_4))
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
	v195 = F_strstr(m, v193, int32(_a_F_check_usermap_5))
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
	F_errmsg(m, int32(_a_F_check_usermap_6), v16+int32(48))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2885), int32(_a_F_check_usermap_4))
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
	F_errmsg(m, int32(_a_F_check_usermap_7), v16+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L16
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(3008), int32(_a_F_check_usermap_2))
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
