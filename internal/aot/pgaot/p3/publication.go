package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_parse_publication_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
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
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	v9 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v9)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v9)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v9)
	v28 = int32(110)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v28)
	if l1 == v9 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L23
	} else {
		goto L141
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L23
	} else {
		goto L137
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L23
	} else {
		goto L133
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L23
	} else {
		goto L129
	}
L5:
	;
	F_errorConflictingDefElem(m, v52, l0)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L23
	} else {
		goto L128
	}
L6:
	;
	m.G0 = v16 + int32(80)
	return
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v47 = v9
	goto L9
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v47<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v54 = int32(_a_F_parse_publication_options_0)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[0])))
	if base.B2i32(v57 == int32(0))|base.B2i32(v57 != v60) != 0 {
		v78 = v57
		v79 = v60
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L6
L11:
	;
	v429 = v47 + int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v429 < v430 {
		v47 = v429
		goto L9
	} else {
		goto L127
	}
L12:
	;
	if v78-v79 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v63 = v53
	v64 = v54
	goto L15
L15:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v68 == int32(0) {
		v78 = v68
		v79 = v67
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v78 = v68
	v79 = v67
	goto L13
L17:
	;
	v71 = int32(1)
	if v68 == v67 {
		v63 = v63 + v71
		v64 = v64 + v71
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v83 == int32(1) {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v252 = int32(_a_F_parse_publication_options_1)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[1])))
	if base.B2i32(v255 == int32(0))|base.B2i32(v255 != v258) != 0 {
		v276 = v255
		v277 = v258
		goto L73
	} else {
		goto L74
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v88)
	v90 = F_defGetString(m, v52)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	v92 = F_pstrdup(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v97 = F_SplitIdentifierString(m, v92, int32(44), v16+int32(76))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if v97 == int32(0) {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if v101 == int32(0) {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v104 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v105 <= v104 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v119 = v104
	goto L30
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v119<<(uint(int32(2))%32))))
	v126 = int32(_a_F_parse_publication_options_2)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[2])))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L11
L32:
	;
	v249 = v119 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v249 < v250 {
		v119 = v249
		goto L30
	} else {
		goto L71
	}
L33:
	;
	if v150-v151 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	goto L33
L35:
	;
	v135 = v125
	v136 = v126
	goto L36
L36:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v150 = v140
	v151 = v139
	goto L34
L38:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v155)
	goto L32
L41:
	;
	goto L42
L42:
	;
	v157 = int32(_a_F_parse_publication_options_3)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[3])))
	if base.B2i32(v160 == int32(0))|base.B2i32(v160 != v163) != 0 {
		v181 = v160
		v182 = v163
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v181-v182 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	v166 = v125
	v167 = v157
	goto L46
L46:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	if v171 == int32(0) {
		v181 = v171
		v182 = v170
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v181 = v171
	v182 = v170
	goto L44
L48:
	;
	v174 = int32(1)
	if v171 == v170 {
		v166 = v166 + v174
		v167 = v167 + v174
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v186)
	goto L32
L51:
	;
	goto L52
L52:
	;
	v188 = int32(_a_F_parse_publication_options_4)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[4])))
	if base.B2i32(v191 == int32(0))|base.B2i32(v191 != v194) != 0 {
		v212 = v191
		v213 = v194
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v212-v213 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	v197 = v125
	v198 = v188
	goto L56
L56:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v202 == int32(0) {
		v212 = v202
		v213 = v201
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v212 = v202
	v213 = v201
	goto L54
L58:
	;
	v205 = int32(1)
	if v202 == v201 {
		v197 = v197 + v205
		v198 = v198 + v205
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v217)
	goto L32
L61:
	;
	goto L62
L62:
	;
	v219 = int32(_a_F_parse_publication_options_5)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[5])))
	if base.B2i32(v222 == int32(0))|base.B2i32(v222 != v225) != 0 {
		v243 = v222
		v244 = v225
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v243-v244 != 0 {
		goto L3
	} else {
		goto L70
	}
L64:
	;
	goto L63
L65:
	;
	v228 = v125
	v229 = v219
	goto L66
L66:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v233 == int32(0) {
		v243 = v233
		v244 = v232
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v243 = v233
	v244 = v232
	goto L64
L68:
	;
	v236 = int32(1)
	if v233 == v232 {
		v228 = v228 + v236
		v229 = v229 + v236
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v246)
	goto L32
L71:
	;
	goto L31
L72:
	;
	if v276-v277 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	goto L72
L74:
	;
	v261 = v53
	v262 = v252
	goto L75
L75:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v266 == int32(0) {
		v276 = v266
		v277 = v265
		goto L73
	} else {
		goto L77
	}
L76:
	;
	v276 = v266
	v277 = v265
	goto L73
L77:
	;
	v269 = int32(1)
	if v266 == v265 {
		v261 = v261 + v269
		v262 = v262 + v269
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v281 == int32(1) {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v289 = int32(_a_F_parse_publication_options_6)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_publication_options[6])))
	if base.B2i32(v292 == int32(0))|base.B2i32(v292 != v295) != 0 {
		v313 = v292
		v314 = v295
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v284 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v284)
	v286 = F_defGetBoolean(m, v52)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v286)
	goto L11
L84:
	;
	if v313-v314 != 0 {
		goto L2
	} else {
		goto L91
	}
L85:
	;
	goto L84
L86:
	;
	v298 = v53
	v299 = v289
	goto L87
L87:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+1)))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)))
	if v303 == int32(0) {
		v313 = v303
		v314 = v302
		goto L85
	} else {
		goto L89
	}
L88:
	;
	v313 = v303
	v314 = v302
	goto L85
L89:
	;
	v306 = int32(1)
	if v303 == v302 {
		v298 = v298 + v306
		v299 = v299 + v306
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v316 == int32(1) {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	v319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v319)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v321 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v506 = int32(_a_F_parse_publication_options_7)
	goto L1
L94:
	;
	goto L95
L95:
	;
	v325 = F_defGetString(m, v52)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L23
	} else {
		goto L96
	}
L96:
	;
	v330 = v325
	v331 = int32(_a_F_parse_publication_options_8)
	goto L98
L97:
	;
	if v368 != 0 {
		goto L110
	} else {
		goto L111
	}
L98:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v334 == v335 {
		v357 = v334
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v368 = int32(0)
	goto L97
L100:
	;
	v359 = int32(1)
	if v357 != 0 {
		v330 = v330 + v359
		v331 = v331 + v359
		goto L98
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v334-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v345 = v334 | int32(32)
	goto L104
L103:
	;
	v345 = v334
	goto L104
L104:
	;
	if base.Ui32((v335-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v354 = v335 | int32(32)
	goto L107
L106:
	;
	v354 = v335
	goto L107
L107:
	;
	if v345 == v354 {
		v357 = v345
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v368 = v345 - v354
	goto L97
L109:
	;
	goto L99
L110:
	;
	v372 = v325
	v373 = int32(_a_F_parse_publication_options_9)
	goto L114
L111:
	;
	v413 = int32(110)
	goto L112
L112:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v413)
	goto L11
L113:
	;
	if v410 != 0 {
		v506 = v325
		goto L1
	} else {
		goto L126
	}
L114:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v376 == v377 {
		v399 = v376
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v410 = int32(0)
	goto L113
L116:
	;
	v401 = int32(1)
	if v399 != 0 {
		v372 = v372 + v401
		v373 = v373 + v401
		goto L114
	} else {
		goto L125
	}
L117:
	;
	if base.Ui32((v376-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v387 = v376 | int32(32)
	goto L120
L119:
	;
	v387 = v376
	goto L120
L120:
	;
	if base.Ui32((v377-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v396 = v377 | int32(32)
	goto L123
L122:
	;
	v396 = v377
	goto L123
L123:
	;
	if v387 == v396 {
		v399 = v387
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v410 = v387 - v396
	goto L113
L125:
	;
	goto L115
L126:
	;
	v413 = int32(115)
	goto L112
L127:
	;
	goto L10
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L23
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_parse_publication_options_0)
	F_errmsg(m, int32(_a_F_parse_publication_options_10), v16+int32(16))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L23
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_parse_publication_options_11), int32(136), int32(_a_F_parse_publication_options_12))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L23
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L23
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_parse_publication_options_0)
	F_errmsg(m, int32(_a_F_parse_publication_options_13), v16)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L23
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_parse_publication_options_11), int32(155), int32(_a_F_parse_publication_options_12))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L23
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L23
	} else {
		goto L138
	}
L138:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v494
	F_errmsg(m, int32(_a_F_parse_publication_options_14), v16-int32(-64))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L23
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_parse_publication_options_11), int32(175), int32(_a_F_parse_publication_options_12))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L23
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L23
	} else {
		goto L142
	}
L142:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v514
	F_errmsg(m, int32(_a_F_parse_publication_options_15), v16+int32(48))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L23
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(_a_F_parse_publication_options_9)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_parse_publication_options_8)
	F_errdetail(m, int32(_a_F_parse_publication_options_16), v16+int32(32))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L23
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_parse_publication_options_11), int32(2139), int32(_a_F_parse_publication_options_17))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L23
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
