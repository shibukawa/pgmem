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
	var v46 int32
	_ = v46
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
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
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
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
	v503 = m.ExcPending
	if v503 != 0 {
		goto L24
	} else {
		goto L148
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L24
	} else {
		goto L144
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L24
	} else {
		goto L140
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L24
	} else {
		goto L136
	}
L5:
	;
	F_errorConflictingDefElem(m, v52, l0)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L24
	} else {
		goto L135
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
	v46 = v9
	goto L9
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v46<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v54 = int32(320771)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[477])))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v58 == int32(0) {
		v77 = v57
		v78 = v58
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L6
L11:
	;
	v422 = v46 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v422 < v423 {
		v46 = v422
		goto L9
	} else {
		goto L134
	}
L12:
	;
	if v78-v77 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v57 != v58 {
		v77 = v57
		v78 = v58
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v62 = v53
	v63 = v54
	goto L16
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v66
		v78 = v67
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v77 = v66
	v78 = v67
	goto L13
L18:
	;
	v70 = int32(1)
	if v66 == v67 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v82 == int32(1) {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v247 = int32(84060)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, _consts[478])))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v251 == int32(0) {
		v270 = v250
		v271 = v251
		goto L78
	} else {
		goto L79
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
	v89 = F_defGetString(m, v52)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	v91 = F_pstrdup(m, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v96 = F_SplitIdentifierString(m, v91, int32(44), v16+int32(76))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if v96 == int32(0) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if v100 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v103 = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v104 <= v103 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v116 = v103
	goto L31
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v116<<(uint(int32(2))%32))))
	v125 = int32(81504)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, _consts[479])))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v129 == int32(0) {
		v148 = v128
		v149 = v129
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L11
L33:
	;
	v244 = v116 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v244 < v245 {
		v116 = v244
		goto L31
	} else {
		goto L76
	}
L34:
	;
	if v149-v148 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	goto L34
L36:
	;
	if v128 != v129 {
		v148 = v128
		v149 = v129
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v133 = v124
	v134 = v125
	goto L38
L38:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v138 == int32(0) {
		v148 = v137
		v149 = v138
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v148 = v137
	v149 = v138
	goto L35
L40:
	;
	v141 = int32(1)
	if v137 == v138 {
		v133 = v133 + v141
		v134 = v134 + v141
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v153)
	goto L33
L43:
	;
	goto L44
L44:
	;
	v155 = int32(354215)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _consts[480])))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v159 == int32(0) {
		v178 = v158
		v179 = v159
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v179-v178 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	goto L45
L47:
	;
	if v158 != v159 {
		v178 = v158
		v179 = v159
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v163 = v124
	v164 = v155
	goto L49
L49:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v168 == int32(0) {
		v178 = v167
		v179 = v168
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v178 = v167
	v179 = v168
	goto L46
L51:
	;
	v171 = int32(1)
	if v167 == v168 {
		v163 = v163 + v171
		v164 = v164 + v171
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v183)
	goto L33
L54:
	;
	goto L55
L55:
	;
	v185 = int32(349218)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _consts[481])))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v189 == int32(0) {
		v208 = v188
		v209 = v189
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v209-v208 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	goto L56
L58:
	;
	if v188 != v189 {
		v208 = v188
		v209 = v189
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v193 = v124
	v194 = v185
	goto L60
L60:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	if v198 == int32(0) {
		v208 = v197
		v209 = v198
		goto L57
	} else {
		goto L62
	}
L61:
	;
	v208 = v197
	v209 = v198
	goto L57
L62:
	;
	v201 = int32(1)
	if v197 == v198 {
		v193 = v193 + v201
		v194 = v194 + v201
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v213)
	goto L33
L65:
	;
	goto L66
L66:
	;
	v215 = int32(355250)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, _consts[482])))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v219 == int32(0) {
		v238 = v218
		v239 = v219
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v239-v238 != 0 {
		goto L3
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	if v218 != v219 {
		v238 = v218
		v239 = v219
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v223 = v124
	v224 = v215
	goto L71
L71:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v228 == int32(0) {
		v238 = v227
		v239 = v228
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v238 = v227
	v239 = v228
	goto L68
L73:
	;
	v231 = int32(1)
	if v227 == v228 {
		v223 = v223 + v231
		v224 = v224 + v231
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v241 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v241)
	goto L33
L76:
	;
	goto L32
L77:
	;
	if v271-v270 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	goto L77
L79:
	;
	if v250 != v251 {
		v270 = v250
		v271 = v251
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v255 = v53
	v256 = v247
	goto L81
L81:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	if v260 == int32(0) {
		v270 = v259
		v271 = v260
		goto L78
	} else {
		goto L83
	}
L82:
	;
	v270 = v259
	v271 = v260
	goto L78
L83:
	;
	v263 = int32(1)
	if v259 == v260 {
		v255 = v255 + v263
		v256 = v256 + v263
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v275 == int32(1) {
		goto L5
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v283 = int32(146634)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, _consts[483])))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v287 == int32(0) {
		v306 = v286
		v307 = v287
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v278 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v278)
	v280 = F_defGetBoolean(m, v52)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L24
	} else {
		goto L89
	}
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v280)
	goto L11
L90:
	;
	if v307-v306 != 0 {
		goto L2
	} else {
		goto L98
	}
L91:
	;
	goto L90
L92:
	;
	if v286 != v287 {
		v306 = v286
		v307 = v287
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v291 = v53
	v292 = v283
	goto L94
L94:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+1)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v296 == int32(0) {
		v306 = v295
		v307 = v296
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v306 = v295
	v307 = v296
	goto L91
L96:
	;
	v299 = int32(1)
	if v295 == v296 {
		v291 = v291 + v299
		v292 = v292 + v299
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v309 == int32(1) {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v314 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v499 = int32(738681)
	goto L1
L101:
	;
	goto L102
L102:
	;
	v318 = F_defGetString(m, v52)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L24
	} else {
		goto L103
	}
L103:
	;
	v323 = v318
	v324 = int32(370343)
	goto L105
L104:
	;
	if v361 != 0 {
		goto L117
	} else {
		goto L118
	}
L105:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v327 == v328 {
		v350 = v327
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v361 = int32(0)
	goto L104
L107:
	;
	v352 = int32(1)
	if v350 != 0 {
		v323 = v323 + v352
		v324 = v324 + v352
		goto L105
	} else {
		goto L116
	}
L108:
	;
	if base.Ui32((v327-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v338 = v327 | int32(32)
	goto L111
L110:
	;
	v338 = v327
	goto L111
L111:
	;
	if base.Ui32((v328-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v347 = v328 | int32(32)
	goto L114
L113:
	;
	v347 = v328
	goto L114
L114:
	;
	if v338 == v347 {
		v350 = v338
		goto L107
	} else {
		goto L115
	}
L115:
	;
	v361 = v338 - v347
	goto L104
L116:
	;
	goto L106
L117:
	;
	v365 = v318
	v366 = int32(447601)
	goto L121
L118:
	;
	v406 = int32(110)
	goto L119
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v406)
	goto L11
L120:
	;
	if v403 != 0 {
		v499 = v318
		goto L1
	} else {
		goto L133
	}
L121:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v369 == v370 {
		v392 = v369
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v403 = int32(0)
	goto L120
L123:
	;
	v394 = int32(1)
	if v392 != 0 {
		v365 = v365 + v394
		v366 = v366 + v394
		goto L121
	} else {
		goto L132
	}
L124:
	;
	if base.Ui32((v369-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v380 = v369 | int32(32)
	goto L127
L126:
	;
	v380 = v369
	goto L127
L127:
	;
	if base.Ui32((v370-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v389 = v370 | int32(32)
	goto L130
L129:
	;
	v389 = v370
	goto L130
L130:
	;
	if v380 == v389 {
		v392 = v380
		goto L123
	} else {
		goto L131
	}
L131:
	;
	v403 = v380 - v389
	goto L120
L132:
	;
	goto L122
L133:
	;
	v406 = int32(115)
	goto L119
L134:
	;
	goto L10
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L24
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(320771)
	F_errmsg(m, int32(682330), v16+int32(16))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L24
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(491199), int32(136), int32(136529))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L24
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L24
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(320771)
	F_errmsg(m, int32(709230), v16)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(491199), int32(155), int32(136529))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L24
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L24
	} else {
		goto L145
	}
L145:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v487
	F_errmsg(m, int32(706987), v16-int32(-64))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L24
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(491199), int32(175), int32(136529))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L24
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L24
	} else {
		goto L149
	}
L149:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v507
	F_errmsg(m, int32(709135), v16+int32(48))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L24
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(447601)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(370343)
	F_errdetail(m, int32(649304), v16+int32(32))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(491199), int32(2139), int32(246519))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L24
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
