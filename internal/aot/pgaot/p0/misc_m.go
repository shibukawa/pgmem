package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	v6 = int32(0)
	if l3 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v16 != int32(37) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(1)
L4:
	;
	return int32(0)
L5:
	;
	v27 = int32(0)
	if base.B2i32(l1 <= v27)|base.B2i32(l3 <= v27) != 0 {
		v460 = l2
		v461 = l3
		v464 = base.B2i32(int32(0) < l1)
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v503
L7:
	;
	if v464 != 0 {
		v503 = v6
		goto L6
	} else {
		goto L147
	}
L8:
	;
	v32 = l0
	v33 = l1
	v34 = l2
	v35 = l3
	goto L9
L9:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	switch v45 - int32(92) {
	case 0:
		goto L21
	case 1, 2:
		goto L18
	case 3:
		goto L19
	default:
		goto L22
	}
L10:
	;
	v460 = v453
	v461 = v451
	v464 = v449
	goto L7
L11:
	;
	v448 = int32(0)
	v449 = base.B2i32(v448 < v443)
	v450 = int32(1)
	v451 = v445 - v450
	v453 = v444 + v450
	if v443 <= v448 {
		v460 = v453
		v461 = v451
		v464 = v449
		goto L7
	} else {
		goto L145
	}
L12:
	;
	v439 = int32(1)
	v443 = v33 - v439
	v444 = v437
	v445 = v438
	v447 = v32 + v439
	goto L11
L13:
	;
	v430 = F_pg_strncoll(m, v195, v300, v32, v33, l4)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L140
	}
L14:
	;
	v425 = F_pg_strncoll(m, v34, v418-v34, v32, v33, l4)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L139
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L135
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L131
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L127
	}
L18:
	;
	if l4 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L19:
	;
	v139 = F_pg_mblen_with_len(m, v32, v33)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L58
	}
L20:
	;
	v64 = v32
	v65 = v33
	v66 = v34
	v67 = v35
	goto L30
L21:
	;
	if v35 <= int32(1) {
		goto L17
	} else {
		goto L25
	}
L22:
	;
	if v45 != int32(37) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v50 = int32(1)
	if base.Ui32(v35) <= base.Ui32(v50) {
		v503 = v50
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v55 == v56 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v58 = int32(1)
	v437 = v34 + v58
	v438 = v35 - v58
	goto L12
L27:
	;
	goto L28
L28:
	;
	return int32(0)
L29:
	;
	if v65 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	v77 = int32(1)
	v78 = v67 - v77
	v80 = v66 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	switch v81 - int32(92) {
	case 0:
		goto L32
	case 1, 2:
		v103 = v81
		goto L29
	case 3:
		goto L34
	default:
		goto L35
	}
L31:
	;
	if v78 == int32(1) {
		goto L16
	} else {
		goto L42
	}
L32:
	;
	goto L31
L33:
	;
	if base.Ui32(int32(2)) < base.Ui32(v67) {
		v64 = v94
		v65 = v95
		v66 = v80
		v67 = v78
		goto L30
	} else {
		goto L41
	}
L34:
	;
	if v65 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v81 == int32(37) {
		v94 = v64
		v95 = v65
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v103 = v81
	goto L29
L37:
	;
	return int32(-1)
L38:
	;
	goto L39
L39:
	;
	v90 = F_pg_mblen_with_len(m, v64, v65)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v94 = v64 + v90
	v95 = v65 - v90
	goto L33
L41:
	;
	v503 = int32(1)
	goto L6
L42:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
	v103 = v102
	goto L29
L43:
	;
	return int32(-1)
L44:
	;
	goto L45
L45:
	;
	v110 = v64
	v111 = v65
	goto L46
L46:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v103&int32(255) != v123 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	return int32(-1)
L48:
	;
	v131 = F_pg_mblen_with_len(m, v110, v111)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L56
	}
L49:
	;
	if l4 == int32(0) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v128 = F_MB_MatchText(m, v110, v111, v80, v78, l4)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v127 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v128 != 0 {
		v503 = v128
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L48
L56:
	;
	v134 = v111 - v131
	if int32(0) < v134 {
		v110 = v110 + v131
		v111 = v134
		goto L46
	} else {
		goto L57
	}
L57:
	;
	goto L47
L58:
	;
	v443 = v33 - v139
	v444 = v34
	v445 = v35
	v447 = v32 + v139
	goto L11
L59:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v45 == v359 {
		v437 = v34
		v438 = v35
		goto L12
	} else {
		goto L126
	}
L60:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v145 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	if v35 == int32(0) {
		v418 = v34
		goto L14
	} else {
		goto L62
	}
L62:
	;
	v151 = v35
	v153 = v6
	v155 = v34
	goto L66
L63:
	;
	v316 = v33
	v321 = v32
	goto L107
L64:
	;
	v304 = v34
	v305 = v151
	v309 = v155
	v310 = v6
	v311 = v155 - v34
	goto L63
L65:
	;
	v194 = v192 - v34
	v195 = F_palloc(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L77
	}
L66:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	switch v161 - int32(92) {
	case 0:
		goto L69
	case 1, 2:
		v177 = v151
		v178 = v153
		v179 = v155
		goto L68
	case 3:
		goto L70
	default:
		goto L71
	}
L67:
	;
	v184 = int32(1)
	if v178&v184 == int32(0) {
		v418 = v181
		goto L14
	} else {
		goto L76
	}
L68:
	;
	v180 = int32(1)
	v181 = v179 + v180
	v183 = v177 - v180
	if v183 != 0 {
		v151 = v183
		v153 = v178
		v155 = v181
		goto L66
	} else {
		goto L75
	}
L69:
	;
	v171 = v151 - int32(1)
	if v171 == int32(0) {
		goto L15
	} else {
		goto L74
	}
L70:
	;
	if v153&int32(1) == int32(0) {
		goto L64
	} else {
		goto L73
	}
L71:
	;
	if v161 != int32(37) {
		v177 = v151
		v178 = v153
		v179 = v155
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v190 = v151
	v192 = v155
	v193 = v6
	goto L65
L74:
	;
	v174 = int32(1)
	v177 = v171
	v178 = v174
	v179 = v155 + v174
	goto L68
L75:
	;
	goto L67
L76:
	;
	v190 = int32(0)
	v192 = v181
	v193 = v184
	goto L65
L77:
	;
	if base.Ui32(v192) <= base.Ui32(v34) {
		v292 = v195
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v300 = v292 - v195
	if v193 != 0 {
		goto L13
	} else {
		goto L106
	}
L79:
	;
	v199 = v194 & int32(3)
	if v199 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v34-v192) {
		v292 = v232
		goto L78
	} else {
		goto L90
	}
L81:
	;
	v232 = v195
	v233 = v34
	goto L80
L82:
	;
	goto L83
L83:
	;
	v207 = v195
	v208 = v34
	v211 = v6
	goto L84
L84:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v215 != int32(92) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v232 = v221
	v233 = v223
	goto L80
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v215)
	v221 = v207 + int32(1)
	goto L88
L87:
	;
	v221 = v207
	goto L88
L88:
	;
	v222 = int32(1)
	v223 = v208 + v222
	v225 = v211 + v222
	if v225 != v199 {
		v207 = v221
		v208 = v223
		v211 = v225
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v248 = v232
	v249 = v233
	goto L91
L91:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v256 != int32(92) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v292 = v283
	goto L78
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v256)
	v262 = v248 + int32(1)
	goto L95
L94:
	;
	v262 = v248
	goto L95
L95:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	if v263 != int32(92) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v263)
	v269 = v262 + int32(1)
	goto L98
L97:
	;
	v269 = v262
	goto L98
L98:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+2)))
	if v270 != int32(92) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v270)
	v276 = v269 + int32(1)
	goto L101
L100:
	;
	v276 = v269
	goto L101
L101:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+3)))
	if v277 != int32(92) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v277)
	v283 = v276 + int32(1)
	goto L104
L103:
	;
	v283 = v276
	goto L104
L104:
	;
	v285 = v249 + int32(4)
	if v285 != v192 {
		v248 = v283
		v249 = v285
		goto L91
	} else {
		goto L105
	}
L105:
	;
	goto L92
L106:
	;
	v304 = v195
	v305 = v190
	v309 = v192
	v310 = v195
	v311 = v300
	goto L63
L107:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_MB_MatchText[0]))
	if v329 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v333 = F_pg_strncoll(m, v304, v311, v32, v321-v32, l4)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L111
L113:
	;
	if v316 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	if v333 != 0 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v335 = F_MB_MatchText(m, v321, v316, v309, v305, l4)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	if v335 != int32(1) {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	if v310 == int32(0) {
		v503 = int32(1)
		goto L6
	} else {
		goto L118
	}
L118:
	;
	F_pfree(m, v310)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	return int32(1)
L120:
	;
	v348 = int32(0)
	if v310 == v348 {
		v503 = v348
		goto L6
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v355 = F_pg_mblen_with_len(m, v321, v316)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	F_pfree(m, v310)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	return int32(0)
L125:
	;
	v316 = v316 - v355
	v321 = v355 + v321
	goto L107
L126:
	;
	return int32(0)
L127:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	F_errmsg(m, int32(_a_F_MB_MatchText_0), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_MB_MatchText_1), int32(107), int32(_a_F_MB_MatchText_2))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_MB_MatchText_0), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_MB_MatchText_1), int32(169), int32(_a_F_MB_MatchText_2))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_MB_MatchText_0), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_MB_MatchText_1), int32(237), int32(_a_F_MB_MatchText_2))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	return base.B2i32(v425 == int32(0))
L140:
	;
	if v195 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_pfree(m, v195)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	return base.B2i32(v430 == int32(0))
L144:
	;
	goto L143
L145:
	;
	if int32(1) < v445 {
		v32 = v447
		v33 = v443
		v34 = v453
		v35 = v451
		goto L9
	} else {
		goto L146
	}
L146:
	;
	goto L10
L147:
	;
	v471 = int32(1)
	if v461 <= int32(0) {
		v503 = v471
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v476 = v460
	v477 = v461
	goto L149
L149:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v487 != int32(37) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v503 = v471
	goto L6
L151:
	;
	return int32(-1)
L152:
	;
	goto L153
L153:
	;
	v492 = int32(1)
	if v492 < v477 {
		v476 = v476 + v492
		v477 = v477 - v492
		goto L149
	} else {
		goto L154
	}
L154:
	;
	goto L150
}
func F_MJFillInner(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_MemoryContextReset(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v20
		if v11 != 0 {
			v22 = int32(_a_F_MJFillInner_0)
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			*(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0])) = v25
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) int32)(m, v11, v12, v9+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0])) = v23
				if v30 == int32(0) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v64 == int32(0) {
						v73 = v2
					} else {
						v67 = *(*float64)(unsafe.Add(mBase, uint32(v64)+248))
						*(*float64)(unsafe.Add(mBase, uint32(v64)+248)) = base.F64_add(v67, float64(1))
						v73 = v2
					}
					m.G0 = v9 + int32(16)
					return v73
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
					m.T0[v41].(func(*base.Module, int32))(m, v39)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(_a_F_MJFillInner_0)
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0]))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0])) = v47
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
						v53 = m.T0[v52].(func(*base.Module, int32, int32, int32) int32)(m, v37+int32(4), v38, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0])) = v45
							v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
							v59 = v57 & int32(_a_F_MJFillInner_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v59)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							*(*uint16)(unsafe.Add(mBase, uint32(v39)+6)) = uint16(v62)
							v73 = v39
							m.G0 = v9 + int32(16)
							return v73
						}
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
			m.T0[v41].(func(*base.Module, int32))(m, v39)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = int32(_a_F_MJFillInner_0)
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0])) = v47
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
				v53 = m.T0[v52].(func(*base.Module, int32, int32, int32) int32)(m, v37+int32(4), v38, int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_MJFillInner[0])) = v45
					v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
					v59 = v57 & int32(_a_F_MJFillInner_1)
					*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v59)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					*(*uint16)(unsafe.Add(mBase, uint32(v39)+6)) = uint16(v62)
					v73 = v39
					m.G0 = v9 + int32(16)
					return v73
				}
			}
		}
	}
}
func F___math_oflow(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v2 = float64(3.105036184601418e+231)
	if l0 != 0 {
		v4 = base.F64_neg(v2)
	} else {
		v4 = v2
	}
	v5 = F_fp_barrier_1(m, v4)
	return base.F64_mul(v2, v5)
}
func F___multi3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v31 int64
	_ = v31
	v10 = int64(32)
	v11 = int64(base.Ui64(l3) >> (uint(v10) % 64))
	v13 = int64(base.Ui64(l1) >> (uint(v10) % 64))
	v16 = int64(4294967295)
	v17 = l3 & v16
	v19 = l1 & v16
	v20 = v17 * v19
	v24 = int64(base.Ui64(v20)>>(uint(v10)%64)) + v17*v13
	v31 = v19*v11 + v24&v16
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l1*l4 + l2*l3 + v11*v13 + int64(base.Ui64(v24)>>(uint(v10)%64)) + int64(base.Ui64(v31)>>(uint(v10)%64))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v20&v16 | v31<<(uint(v10)%64)
	return
}
func F_makeAConst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 - int32(465) {
	case 0:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = F_palloc0(m, int32(20))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(1997159792712)
			v24 = v18
			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
			v27 = v24
			return v27
		}
	case 1:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = F_palloc0(m, int32(20))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(2001454760008)
			v24 = v9
			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
			v27 = v24
			return v27
		}
	default:
		v27 = int32(0)
		return v27
	}
}
func F_makeNotExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v14 = F_list_make1_impl(m, int32(1), v6+int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = F_makeBoolExpr(m, int32(2), v14, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v18
		}
	}
}
func F_makeNullConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_get_typlenbyval(m, l0, v10+int32(14), v10+int32(13))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+14)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
		v23 = F_palloc0(m, int32(32))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(-1)
			*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)) = uint8(v21)
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v28)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(7)
			m.G0 = v10 + int32(16)
			return v23
		}
	}
}
func F_make_SAOP_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
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
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
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
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = F_get_array_type(m, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			v186 = int32(0)
			m.G0 = v15 + int32(32)
			return v186
		} else {
			if l6 != 0 {
				v25 = F_palloc0(m, int32(36))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(35)
					v30 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)) = uint8(v30)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l5
					v144 = v25
					v154 = F_palloc0(m, int32(36))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(20)
						v159 = F_get_opcode(m, l0)
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = l4
							v162 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v154)+20)) = uint8(v162)
							*(*int64)(unsafe.Add(mBase, uint32(v154)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v159
							*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v144
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v144
							v175 = F_list_make2_impl(m, v15+int32(8), v15+int32(4))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v175
								v186 = v154
								m.G0 = v15 + int32(32)
								return v186
							}
						}
					}
				}
			} else {
				if l5 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1)
					F_get_typlenbyvalalign(m, l2, v15+int32(30), v15+int32(29), v15+int32(28))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v51 = F_palloc(m, int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v54 = F_palloc(m, int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v115 = v51
								v116 = v54
								v119 = int32(-1)
								v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+30)))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)))
								v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+28)))
								v129 = F_construct_md_array(m, v115, v116, int32(1), v15+int32(24), v15+int32(20), l2, v126, v127, v128)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									v131 = int32(0)
									v133 = F_makeConst(m, v17, v119, l3, v119, v129, v131, v131)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v115)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v116)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												F_list_free(m, l5)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													v144 = v133
													v154 = F_palloc0(m, int32(36))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = l0
														*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(20)
														v159 = F_get_opcode(m, l0)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = l4
															v162 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v154)+20)) = uint8(v162)
															*(*int64)(unsafe.Add(mBase, uint32(v154)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v159
															*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v144
															*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v144
															v175 = F_list_make2_impl(m, v15+int32(8), v15+int32(4))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v175
																v186 = v154
																m.G0 = v15 + int32(32)
																return v186
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
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1)
					F_get_typlenbyvalalign(m, l2, v15+int32(30), v15+int32(29), v15+int32(28))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
						v71 = F_palloc(m, v68<<(uint(int32(2))%32))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
							v74 = F_palloc(m, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
								if v76 <= int32(0) {
									v115 = v71
									v116 = v74
								} else {
									v86 = int32(0)
									for {
										v93 = v86 << (uint(int32(2)) % 32)
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v95+v93)))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v71+v93))) = v98
										v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)))
										*(*uint8)(unsafe.Add(mBase, uint32(v86+v74))) = uint8(v101)
										v104 = v86 + int32(1)
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
										if v104 < v105 {
											v86 = v104
											continue
										} else {
											break
										}
										break
									}
									v115 = v71
									v116 = v74
								}
								v119 = int32(-1)
								v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+30)))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)))
								v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+28)))
								v129 = F_construct_md_array(m, v115, v116, int32(1), v15+int32(24), v15+int32(20), l2, v126, v127, v128)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									v131 = int32(0)
									v133 = F_makeConst(m, v17, v119, l3, v119, v129, v131, v131)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v115)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v116)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												F_list_free(m, l5)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													v144 = v133
													v154 = F_palloc0(m, int32(36))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = l0
														*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(20)
														v159 = F_get_opcode(m, l0)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = l4
															v162 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v154)+20)) = uint8(v162)
															*(*int64)(unsafe.Add(mBase, uint32(v154)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v159
															*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v144
															*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v144
															v175 = F_list_make2_impl(m, v15+int32(8), v15+int32(4))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v175
																v186 = v154
																m.G0 = v15 + int32(32)
																return v186
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
func F_make_ands_explicit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	if l0 == int32(0) {
		v6 = F_palloc0(m, int32(32))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(-1)
			v12 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+24)) = uint16(v12)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = int64(4294967297)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(4294967295)
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(68719476743)
			return v6
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v21 == int32(1) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			return v25
		} else {
			v28 = F_palloc0(m, int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = int64(21)
				return v28
			}
		}
	}
}
func F_make_plain_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v196 int64
	_ = v196
	v4 = l3
	v5 = l4
	v6 = l5
	v7 = l6
	v14 = F_palloc0(m, int32(168))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(318)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)) = uint8(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l7
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+9)) = uint8(v29)
	if l7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = F_contain_leaked_vars(m, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v36 = v29
	goto L5
L5:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)) = uint8(v36)
	if l1 == v37 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v36 = v32 ^ int32(1)
	goto L5
L7:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if l8 != 0 {
		goto L36
	} else {
		goto L37
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = int64(0)
	v130 = F_pull_varnos(m, l0, l1)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v42 != int32(17) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 != int32(2) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = F_pull_varnos(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v53
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v57 == v56 {
		v65 = v56
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = F_pull_varnos(m, l0, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 < int32(2) {
		v65 = v56
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v65 = v64
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v70 = F_bms_union(m, v69, v66)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v73 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v76 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v79 = int32(0)
	if base.B2i32(v73 == v79)|base.B2i32(v76 == v79) != 0 {
		v124 = v79
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v124 != 0 {
		goto L7
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v89 < v90 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v92 = v89
	goto L26
L25:
	;
	v92 = v90
	goto L26
L26:
	;
	if v92 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v95 = int32(1)
	goto L29
L28:
	;
	v95 = v92
	goto L29
L29:
	;
	v96 = int32(8)
	v101 = int32(0)
	goto L30
L30:
	;
	v108 = v101 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v76+v96+v108)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v73+v96+v108)))
	v113 = v110 & v112
	v115 = base.B2i32(v113 != int32(0))
	if v113 != 0 {
		v124 = v115
		goto L22
	} else {
		goto L32
	}
L31:
	;
	v124 = v115
	goto L22
L32:
	;
	v117 = v101 + int32(1)
	if v117 != v95 {
		v101 = v117
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+9)) = uint8(v125)
	goto L7
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v130
	goto L7
L36:
	;
	v137 = l8
	goto L38
L37:
	;
	v137 = v136
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v140 = F_bms_difference(m, v136, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v142 = int32(0)
	if v140 == v142 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v177
	F_bms_free(m, v140)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L53
	}
L41:
	;
	v177 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v149 = int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v150 <= v149 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v153 = v149
	goto L46
L45:
	;
	v153 = v150
	goto L46
L46:
	;
	v157 = int32(0)
	v159 = v142
	goto L47
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(8)+v157<<(uint(int32(2))%32))))
	if v165 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v177 = v168
	goto L40
L49:
	;
	v168 = v159 + base.I32_popcnt(v165)
	goto L51
L50:
	;
	v168 = v159
	goto L51
L51:
	;
	v170 = v157 + int32(1)
	if v170 != v153 {
		v157 = v170
		v159 = v168
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v183 = v181 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v183
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v185
	v187 = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v183
	v196 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v196
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v196
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v196
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+120)) = uint8(v185)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+160)) = v196
	*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v14)+144)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v187
	return v14
}
func F_manifest_process_file(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v303 int64
	_ = v303
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v324 int64
	_ = v324
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v349 int64
	_ = v349
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
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
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int64
	_ = v977
	var v979 int64
	_ = v979
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
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
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1378 int64
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1430 int64
	_ = v1430
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1454 int64
	_ = v1454
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1470 int32
	_ = v1470
	var v1485 int32
	_ = v1485
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1549 int32
	_ = v1549
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1568 int32
	_ = v1568
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v18 = F_strlen(m, l1)
	mBase = m.M
	v24 = v18 - int32(1636608432)
	if l1&int32(3) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v286 = base.B2i32(base.Ui32(v283) < base.Ui32(v284))
	goto L41
L2:
	;
	v256 = int32(14)
	v258 = v252 ^ v253 - base.I32_rotl(v252, v256)
	v262 = v258 ^ v251 - base.I32_rotl(v258, int32(11))
	v266 = v262 ^ v252 - base.I32_rotl(v262, int32(25))
	v270 = v266 ^ v258 - base.I32_rotl(v266, int32(16))
	v274 = v270 ^ v262 - base.I32_rotl(v270, int32(4))
	v278 = v274 ^ v266 - base.I32_rotl(v274, v256)
	goto L1
L3:
	;
	switch v182 - int32(1) {
	case 0:
		v244 = v183
		v245 = v184
		v246 = v185
		goto L30
	case 1:
		v237 = v183
		v238 = v184
		v239 = v185
		goto L31
	case 2:
		v230 = v183
		v231 = v184
		v232 = v185
		goto L32
	case 3:
		v224 = v184
		v225 = v185
		goto L33
	case 4:
		v220 = v184
		v221 = v185
		goto L34
	case 5:
		v214 = v184
		v215 = v185
		goto L35
	case 6:
		v208 = v184
		v209 = v185
		goto L36
	case 7:
		v203 = v185
		goto L37
	case 8:
		v198 = v185
		goto L38
	case 9:
		v193 = v185
		goto L39
	case 10:
		goto L40
	default:
		v251 = v183
		v252 = v184
		v253 = v185
		goto L2
	}
L4:
	;
	v133 = l1
	v134 = v18
	v135 = v24
	v136 = v24
	v137 = v24
	goto L27
L5:
	;
	if base.Ui32(int32(11)) < base.Ui32(v18) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(v18) < base.Ui32(int32(12)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v181 = l1
	v182 = v18
	v183 = v24
	v184 = v24
	v185 = v24
	goto L3
L9:
	;
	switch v80 - int32(1) {
	case 0:
		v130 = v81
		goto L16
	case 1:
		v125 = v81
		goto L17
	case 2:
		goto L18
	case 3:
		v118 = v82
		goto L19
	case 4:
		v115 = v82
		goto L20
	case 5:
		v110 = v82
		goto L21
	case 6:
		goto L22
	case 7:
		v101 = v83
		goto L23
	case 8:
		v96 = v83
		goto L24
	case 9:
		v91 = v83
		goto L25
	case 10:
		goto L26
	default:
		v251 = v81
		v252 = v82
		v253 = v83
		goto L2
	}
L10:
	;
	v79 = l1
	v80 = v18
	v81 = v24
	v82 = v24
	v83 = v24
	goto L9
L11:
	;
	goto L12
L12:
	;
	v31 = l1
	v32 = v18
	v33 = v24
	v34 = v24
	v35 = v24
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v38 = v37 + v34
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v42 = v41 + v35
	v44 = int32(4)
	v46 = v39 + v33 - v42 ^ base.I32_rotl(v42, v44)
	v50 = v38 - v46 ^ base.I32_rotl(v46, int32(6))
	v51 = v42 + v38
	v52 = v46 + v51
	v53 = v50 + v52
	v57 = v51 - v50 ^ base.I32_rotl(v50, int32(8))
	v61 = v52 - v57 ^ base.I32_rotl(v57, int32(16))
	v65 = v53 - v61 ^ base.I32_rotl(v61, int32(19))
	v66 = v57 + v53
	v67 = v61 + v66
	v68 = v65 + v67
	v72 = v66 - v65 ^ base.I32_rotl(v65, v44)
	v73 = int32(12)
	v74 = v31 + v73
	v76 = v32 - v73
	if base.Ui32(int32(11)) < base.Ui32(v76) {
		v31 = v74
		v32 = v76
		v33 = v67
		v34 = v68
		v35 = v72
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v79 = v74
	v80 = v76
	v81 = v67
	v82 = v68
	v83 = v72
	goto L9
L15:
	;
	goto L14
L16:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v251 = v130 + v131
	v252 = v82
	v253 = v83
	goto L2
L17:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v130 = v126<<(uint(int32(8))%32) + v125
	goto L16
L18:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+2)))
	v125 = v121<<(uint(int32(16))%32) + v81
	goto L17
L19:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v251 = v119 + v81
	v252 = v118
	v253 = v83
	goto L2
L20:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)))
	v118 = v115 + v116
	goto L19
L21:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)))
	v115 = v111<<(uint(int32(8))%32) + v110
	goto L20
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	v110 = v106<<(uint(int32(16))%32) + v82
	goto L21
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v251 = v102 + v81
	v252 = v104 + v82
	v253 = v101
	goto L2
L24:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	v101 = v97<<(uint(int32(8))%32) + v96
	goto L23
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+9)))
	v96 = v92<<(uint(int32(16))%32) + v91
	goto L24
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+10)))
	v91 = v87<<(uint(int32(24))%32) + v83
	goto L25
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v140 = v139 + v136
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v144 = v143 + v137
	v146 = int32(4)
	v148 = v141 + v135 - v144 ^ base.I32_rotl(v144, v146)
	v152 = v140 - v148 ^ base.I32_rotl(v148, int32(6))
	v153 = v144 + v140
	v154 = v148 + v153
	v155 = v152 + v154
	v159 = v153 - v152 ^ base.I32_rotl(v152, int32(8))
	v163 = v154 - v159 ^ base.I32_rotl(v159, int32(16))
	v167 = v155 - v163 ^ base.I32_rotl(v163, int32(19))
	v168 = v159 + v155
	v169 = v163 + v168
	v170 = v167 + v169
	v174 = v168 - v167 ^ base.I32_rotl(v167, v146)
	v175 = int32(12)
	v176 = v133 + v175
	v178 = v134 - v175
	if base.Ui32(int32(11)) < base.Ui32(v178) {
		v133 = v176
		v134 = v178
		v135 = v169
		v136 = v170
		v137 = v174
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v181 = v176
	v182 = v178
	v183 = v169
	v184 = v170
	v185 = v174
	goto L3
L29:
	;
	goto L28
L30:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v251 = v244 + v247
	v252 = v245
	v253 = v246
	goto L2
L31:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	v244 = v240<<(uint(int32(8))%32) + v237
	v245 = v238
	v246 = v239
	goto L30
L32:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+2)))
	v237 = v233<<(uint(int32(16))%32) + v230
	v238 = v231
	v239 = v232
	goto L31
L33:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+3)))
	v230 = v226<<(uint(int32(24))%32) + v183
	v231 = v224
	v232 = v225
	goto L32
L34:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+4)))
	v224 = v220 + v222
	v225 = v221
	goto L33
L35:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+5)))
	v220 = v216<<(uint(int32(8))%32) + v214
	v221 = v215
	goto L34
L36:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+6)))
	v214 = v210<<(uint(int32(16))%32) + v208
	v215 = v209
	goto L35
L37:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+7)))
	v208 = v204<<(uint(int32(24))%32) + v184
	v209 = v203
	goto L36
L38:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+8)))
	v203 = v199<<(uint(int32(8))%32) + v198
	goto L37
L39:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+9)))
	v198 = v194<<(uint(int32(16))%32) + v193
	goto L38
L40:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+10)))
	v193 = v189<<(uint(int32(24))%32) + v185
	goto L39
L41:
	;
	if v286 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L60
	} else {
		goto L257
	}
L43:
	;
	goto L42
L44:
	;
	v1549 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v1549
	v286 = v1549
	goto L41
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L60
	} else {
		goto L254
	}
L46:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if v303 == int64(4294967296) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v1026 = v1025 & (v278 ^ v270 - base.I32_rotl(v278, int32(24)))
	v1029 = v1024 + v1026<<(uint(int32(4))%32)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1029)))
	if v1030 != 0 {
		goto L172
	} else {
		goto L173
	}
L49:
	;
	v306 = int32(0)
	v308 = int64(2)
	v310 = v303 << (uint(int64(1)) % 64)
	if base.Ui64(v310) <= base.Ui64(v308) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v286 = int32(1)
	goto L41
L51:
	;
	v313 = v308
	goto L53
L52:
	;
	v313 = v310
	goto L53
L53:
	;
	v314 = int64(1)
	if v313&(v313-v314) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v324 = v313
	goto L56
L55:
	;
	v324 = v314 << (uint(int64(64)-base.I64_clz(v313)) % 64)
	goto L56
L56:
	;
	if base.Ui64(v324<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v336 = F_MemoryContextAllocExtended(m, v331, base.I32_wrap_i64(v324)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	goto L43
L60:
	;
	return
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v336
	v339 = int64(1)
	if v324&(v324-v339) == int64(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v349 = v324
	goto L64
L63:
	;
	v349 = v339 << (uint(int64(64)-base.I64_clz(v324)) % 64)
	goto L64
L64:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v349<<(uint(int64(4))%64)) {
		goto L43
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = base.I32_wrap_i64(v349) - int32(1)
	if v349 == int64(4294967296) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v366 = int32(-85899346)
	goto L68
L67:
	;
	v366 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v349), float64(0.9)))
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v366
	if v330 != int64(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v373 = v306
	goto L73
L70:
	;
	goto L71
L71:
	;
	F_pfree(m, v329)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L60
	} else {
		goto L170
	}
L72:
	;
	v670 = v666
	v674 = v306
	goto L118
L73:
	;
	v387 = v329 + v373<<(uint(int32(4))%32)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	if v388 != int32(1) {
		v666 = v373
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v666 = int32(0)
	goto L72
L75:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	v392 = F_strlen(m, v391)
	mBase = m.M
	v398 = v392 - int32(1636608432)
	if v391&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if (v652^v644-base.I32_rotl(v652, int32(24)))&v657 == v373 {
		v666 = v373
		goto L72
	} else {
		goto L116
	}
L77:
	;
	v630 = int32(14)
	v632 = v626 ^ v627 - base.I32_rotl(v626, v630)
	v636 = v632 ^ v625 - base.I32_rotl(v632, int32(11))
	v640 = v636 ^ v626 - base.I32_rotl(v636, int32(25))
	v644 = v640 ^ v632 - base.I32_rotl(v640, int32(16))
	v648 = v644 ^ v636 - base.I32_rotl(v644, int32(4))
	v652 = v648 ^ v640 - base.I32_rotl(v648, v630)
	goto L76
L78:
	;
	switch v556 - int32(1) {
	case 0:
		v618 = v557
		v619 = v558
		v620 = v559
		goto L105
	case 1:
		v611 = v557
		v612 = v558
		v613 = v559
		goto L106
	case 2:
		v604 = v557
		v605 = v558
		v606 = v559
		goto L107
	case 3:
		v598 = v558
		v599 = v559
		goto L108
	case 4:
		v594 = v558
		v595 = v559
		goto L109
	case 5:
		v588 = v558
		v589 = v559
		goto L110
	case 6:
		v582 = v558
		v583 = v559
		goto L111
	case 7:
		v577 = v559
		goto L112
	case 8:
		v572 = v559
		goto L113
	case 9:
		v567 = v559
		goto L114
	case 10:
		goto L115
	default:
		v625 = v557
		v626 = v558
		v627 = v559
		goto L77
	}
L79:
	;
	v507 = v391
	v508 = v392
	v509 = v398
	v510 = v398
	v511 = v398
	goto L102
L80:
	;
	if base.Ui32(int32(11)) < base.Ui32(v392) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(v392) < base.Ui32(int32(12)) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v555 = v391
	v556 = v392
	v557 = v398
	v558 = v398
	v559 = v398
	goto L78
L84:
	;
	switch v454 - int32(1) {
	case 0:
		v504 = v455
		goto L91
	case 1:
		v499 = v455
		goto L92
	case 2:
		goto L93
	case 3:
		v492 = v456
		goto L94
	case 4:
		v489 = v456
		goto L95
	case 5:
		v484 = v456
		goto L96
	case 6:
		goto L97
	case 7:
		v475 = v457
		goto L98
	case 8:
		v470 = v457
		goto L99
	case 9:
		v465 = v457
		goto L100
	case 10:
		goto L101
	default:
		v625 = v455
		v626 = v456
		v627 = v457
		goto L77
	}
L85:
	;
	v453 = v391
	v454 = v392
	v455 = v398
	v456 = v398
	v457 = v398
	goto L84
L86:
	;
	goto L87
L87:
	;
	v405 = v391
	v406 = v392
	v407 = v398
	v408 = v398
	v409 = v398
	goto L88
L88:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v412 = v411 + v408
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	v416 = v415 + v409
	v418 = int32(4)
	v420 = v413 + v407 - v416 ^ base.I32_rotl(v416, v418)
	v424 = v412 - v420 ^ base.I32_rotl(v420, int32(6))
	v425 = v416 + v412
	v426 = v420 + v425
	v427 = v424 + v426
	v431 = v425 - v424 ^ base.I32_rotl(v424, int32(8))
	v435 = v426 - v431 ^ base.I32_rotl(v431, int32(16))
	v439 = v427 - v435 ^ base.I32_rotl(v435, int32(19))
	v440 = v431 + v427
	v441 = v435 + v440
	v442 = v439 + v441
	v446 = v440 - v439 ^ base.I32_rotl(v439, v418)
	v447 = int32(12)
	v448 = v405 + v447
	v450 = v406 - v447
	if base.Ui32(int32(11)) < base.Ui32(v450) {
		v405 = v448
		v406 = v450
		v407 = v441
		v408 = v442
		v409 = v446
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v453 = v448
	v454 = v450
	v455 = v441
	v456 = v442
	v457 = v446
	goto L84
L90:
	;
	goto L89
L91:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	v625 = v504 + v505
	v626 = v456
	v627 = v457
	goto L77
L92:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	v504 = v500<<(uint(int32(8))%32) + v499
	goto L91
L93:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+2)))
	v499 = v495<<(uint(int32(16))%32) + v455
	goto L92
L94:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v625 = v493 + v455
	v626 = v492
	v627 = v457
	goto L77
L95:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+4)))
	v492 = v489 + v490
	goto L94
L96:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+5)))
	v489 = v485<<(uint(int32(8))%32) + v484
	goto L95
L97:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+6)))
	v484 = v480<<(uint(int32(16))%32) + v456
	goto L96
L98:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v625 = v476 + v455
	v626 = v478 + v456
	v627 = v475
	goto L77
L99:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+8)))
	v475 = v471<<(uint(int32(8))%32) + v470
	goto L98
L100:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+9)))
	v470 = v466<<(uint(int32(16))%32) + v465
	goto L99
L101:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+10)))
	v465 = v461<<(uint(int32(24))%32) + v457
	goto L100
L102:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v514 = v513 + v510
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v518 = v517 + v511
	v520 = int32(4)
	v522 = v515 + v509 - v518 ^ base.I32_rotl(v518, v520)
	v526 = v514 - v522 ^ base.I32_rotl(v522, int32(6))
	v527 = v518 + v514
	v528 = v522 + v527
	v529 = v526 + v528
	v533 = v527 - v526 ^ base.I32_rotl(v526, int32(8))
	v537 = v528 - v533 ^ base.I32_rotl(v533, int32(16))
	v541 = v529 - v537 ^ base.I32_rotl(v537, int32(19))
	v542 = v533 + v529
	v543 = v537 + v542
	v544 = v541 + v543
	v548 = v542 - v541 ^ base.I32_rotl(v541, v520)
	v549 = int32(12)
	v550 = v507 + v549
	v552 = v508 - v549
	if base.Ui32(int32(11)) < base.Ui32(v552) {
		v507 = v550
		v508 = v552
		v509 = v543
		v510 = v544
		v511 = v548
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v555 = v550
	v556 = v552
	v557 = v543
	v558 = v544
	v559 = v548
	goto L78
L104:
	;
	goto L103
L105:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	v625 = v618 + v621
	v626 = v619
	v627 = v620
	goto L77
L106:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+1)))
	v618 = v614<<(uint(int32(8))%32) + v611
	v619 = v612
	v620 = v613
	goto L105
L107:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+2)))
	v611 = v607<<(uint(int32(16))%32) + v604
	v612 = v605
	v613 = v606
	goto L106
L108:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+3)))
	v604 = v600<<(uint(int32(24))%32) + v557
	v605 = v598
	v606 = v599
	goto L107
L109:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+4)))
	v598 = v594 + v596
	v599 = v595
	goto L108
L110:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+5)))
	v594 = v590<<(uint(int32(8))%32) + v588
	v595 = v589
	goto L109
L111:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+6)))
	v588 = v584<<(uint(int32(16))%32) + v582
	v589 = v583
	goto L110
L112:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+7)))
	v582 = v578<<(uint(int32(24))%32) + v558
	v583 = v577
	goto L111
L113:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+8)))
	v577 = v573<<(uint(int32(8))%32) + v572
	goto L112
L114:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+9)))
	v572 = v568<<(uint(int32(16))%32) + v567
	goto L113
L115:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+10)))
	v567 = v563<<(uint(int32(24))%32) + v559
	goto L114
L116:
	;
	v661 = v373 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v661)) < base.Ui64(v330) {
		v373 = v661
		goto L73
	} else {
		goto L117
	}
L117:
	;
	goto L74
L118:
	;
	v684 = v329 + v670<<(uint(int32(4))%32)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v684)))
	if v685 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L71
L120:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	v689 = F_strlen(m, v688)
	mBase = m.M
	v695 = v689 - int32(1636608432)
	if v688&int32(3) != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L122
L122:
	;
	v997 = v670 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v997)) < base.Ui64(v330) {
		goto L166
	} else {
		goto L167
	}
L123:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v955 = v949 ^ v941 - base.I32_rotl(v949, int32(24))
	goto L163
L124:
	;
	v927 = int32(14)
	v929 = v923 ^ v924 - base.I32_rotl(v923, v927)
	v933 = v929 ^ v922 - base.I32_rotl(v929, int32(11))
	v937 = v933 ^ v923 - base.I32_rotl(v933, int32(25))
	v941 = v937 ^ v929 - base.I32_rotl(v937, int32(16))
	v945 = v941 ^ v933 - base.I32_rotl(v941, int32(4))
	v949 = v945 ^ v937 - base.I32_rotl(v945, v927)
	goto L123
L125:
	;
	switch v853 - int32(1) {
	case 0:
		v915 = v854
		v916 = v855
		v917 = v856
		goto L152
	case 1:
		v908 = v854
		v909 = v855
		v910 = v856
		goto L153
	case 2:
		v901 = v854
		v902 = v855
		v903 = v856
		goto L154
	case 3:
		v895 = v855
		v896 = v856
		goto L155
	case 4:
		v891 = v855
		v892 = v856
		goto L156
	case 5:
		v885 = v855
		v886 = v856
		goto L157
	case 6:
		v879 = v855
		v880 = v856
		goto L158
	case 7:
		v874 = v856
		goto L159
	case 8:
		v869 = v856
		goto L160
	case 9:
		v864 = v856
		goto L161
	case 10:
		goto L162
	default:
		v922 = v854
		v923 = v855
		v924 = v856
		goto L124
	}
L126:
	;
	v804 = v688
	v805 = v689
	v806 = v695
	v807 = v695
	v808 = v695
	goto L149
L127:
	;
	if base.Ui32(int32(11)) < base.Ui32(v689) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	if base.Ui32(v689) < base.Ui32(int32(12)) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v852 = v688
	v853 = v689
	v854 = v695
	v855 = v695
	v856 = v695
	goto L125
L131:
	;
	switch v751 - int32(1) {
	case 0:
		v801 = v752
		goto L138
	case 1:
		v796 = v752
		goto L139
	case 2:
		goto L140
	case 3:
		v789 = v753
		goto L141
	case 4:
		v786 = v753
		goto L142
	case 5:
		v781 = v753
		goto L143
	case 6:
		goto L144
	case 7:
		v772 = v754
		goto L145
	case 8:
		v767 = v754
		goto L146
	case 9:
		v762 = v754
		goto L147
	case 10:
		goto L148
	default:
		v922 = v752
		v923 = v753
		v924 = v754
		goto L124
	}
L132:
	;
	v750 = v688
	v751 = v689
	v752 = v695
	v753 = v695
	v754 = v695
	goto L131
L133:
	;
	goto L134
L134:
	;
	v702 = v688
	v703 = v689
	v704 = v695
	v705 = v695
	v706 = v695
	goto L135
L135:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	v709 = v708 + v705
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v702)+8))
	v713 = v712 + v706
	v715 = int32(4)
	v717 = v710 + v704 - v713 ^ base.I32_rotl(v713, v715)
	v721 = v709 - v717 ^ base.I32_rotl(v717, int32(6))
	v722 = v713 + v709
	v723 = v717 + v722
	v724 = v721 + v723
	v728 = v722 - v721 ^ base.I32_rotl(v721, int32(8))
	v732 = v723 - v728 ^ base.I32_rotl(v728, int32(16))
	v736 = v724 - v732 ^ base.I32_rotl(v732, int32(19))
	v737 = v728 + v724
	v738 = v732 + v737
	v739 = v736 + v738
	v743 = v737 - v736 ^ base.I32_rotl(v736, v715)
	v744 = int32(12)
	v745 = v702 + v744
	v747 = v703 - v744
	if base.Ui32(int32(11)) < base.Ui32(v747) {
		v702 = v745
		v703 = v747
		v704 = v738
		v705 = v739
		v706 = v743
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v750 = v745
	v751 = v747
	v752 = v738
	v753 = v739
	v754 = v743
	goto L131
L137:
	;
	goto L136
L138:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	v922 = v801 + v802
	v923 = v753
	v924 = v754
	goto L124
L139:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	v801 = v797<<(uint(int32(8))%32) + v796
	goto L138
L140:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+2)))
	v796 = v792<<(uint(int32(16))%32) + v752
	goto L139
L141:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v922 = v790 + v752
	v923 = v789
	v924 = v754
	goto L124
L142:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+4)))
	v789 = v786 + v787
	goto L141
L143:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+5)))
	v786 = v782<<(uint(int32(8))%32) + v781
	goto L142
L144:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+6)))
	v781 = v777<<(uint(int32(16))%32) + v753
	goto L143
L145:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v922 = v773 + v752
	v923 = v775 + v753
	v924 = v772
	goto L124
L146:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+8)))
	v772 = v768<<(uint(int32(8))%32) + v767
	goto L145
L147:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+9)))
	v767 = v763<<(uint(int32(16))%32) + v762
	goto L146
L148:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+10)))
	v762 = v758<<(uint(int32(24))%32) + v754
	goto L147
L149:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	v811 = v810 + v807
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v804)))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v804)+8))
	v815 = v814 + v808
	v817 = int32(4)
	v819 = v812 + v806 - v815 ^ base.I32_rotl(v815, v817)
	v823 = v811 - v819 ^ base.I32_rotl(v819, int32(6))
	v824 = v815 + v811
	v825 = v819 + v824
	v826 = v823 + v825
	v830 = v824 - v823 ^ base.I32_rotl(v823, int32(8))
	v834 = v825 - v830 ^ base.I32_rotl(v830, int32(16))
	v838 = v826 - v834 ^ base.I32_rotl(v834, int32(19))
	v839 = v830 + v826
	v840 = v834 + v839
	v841 = v838 + v840
	v845 = v839 - v838 ^ base.I32_rotl(v838, v817)
	v846 = int32(12)
	v847 = v804 + v846
	v849 = v805 - v846
	if base.Ui32(int32(11)) < base.Ui32(v849) {
		v804 = v847
		v805 = v849
		v806 = v840
		v807 = v841
		v808 = v845
		goto L149
	} else {
		goto L151
	}
L150:
	;
	v852 = v847
	v853 = v849
	v854 = v840
	v855 = v841
	v856 = v845
	goto L125
L151:
	;
	goto L150
L152:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	v922 = v915 + v918
	v923 = v916
	v924 = v917
	goto L124
L153:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+1)))
	v915 = v911<<(uint(int32(8))%32) + v908
	v916 = v909
	v917 = v910
	goto L152
L154:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+2)))
	v908 = v904<<(uint(int32(16))%32) + v901
	v909 = v902
	v910 = v903
	goto L153
L155:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+3)))
	v901 = v897<<(uint(int32(24))%32) + v854
	v902 = v895
	v903 = v896
	goto L154
L156:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+4)))
	v895 = v891 + v893
	v896 = v892
	goto L155
L157:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+5)))
	v891 = v887<<(uint(int32(8))%32) + v885
	v892 = v886
	goto L156
L158:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+6)))
	v885 = v881<<(uint(int32(16))%32) + v879
	v886 = v880
	goto L157
L159:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+7)))
	v879 = v875<<(uint(int32(24))%32) + v855
	v880 = v874
	goto L158
L160:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+8)))
	v874 = v870<<(uint(int32(8))%32) + v869
	goto L159
L161:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+9)))
	v869 = v865<<(uint(int32(16))%32) + v864
	goto L160
L162:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+10)))
	v864 = v860<<(uint(int32(24))%32) + v856
	goto L161
L163:
	;
	v970 = v955 & v954
	v975 = v336 + v970<<(uint(int32(4))%32)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v975)))
	if v976 != 0 {
		v955 = v970 + int32(1)
		goto L163
	} else {
		goto L165
	}
L164:
	;
	v977 = *(*int64)(unsafe.Add(mBase, uint32(v684)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v975)+8)) = v977
	v979 = *(*int64)(unsafe.Add(mBase, uint32(v684)))
	*(*int64)(unsafe.Add(mBase, uint32(v975))) = v979
	goto L122
L165:
	;
	goto L164
L166:
	;
	v1001 = v997
	goto L168
L167:
	;
	v1001 = int32(0)
	goto L168
L168:
	;
	v1003 = v674 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1003)) < base.Ui64(v330) {
		v670 = v1001
		v674 = v1003
		goto L118
	} else {
		goto L169
	}
L169:
	;
	goto L119
L170:
	;
	goto L50
L171:
	;
	return
L172:
	;
	v1036 = v1026
	v1037 = int32(0)
	v1039 = v1029
	goto L176
L173:
	;
	v1485 = v1029
	goto L174
L174:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v1494 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v1493 + v1494
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1485))) = v1494
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+24))
	v1502 = F_MemoryContextStrdup(m, v1501, l1)
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L60
	} else {
		goto L253
	}
L175:
	;
	v1485 = v1470
	goto L174
L176:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+4))
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v1050 == int32(0))|base.B2i32(v1050 != v1053) != 0 {
		v1071 = v1050
		v1072 = v1053
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v1470 = v1461
	goto L175
L178:
	;
	if v1071-v1072 == int32(0) {
		goto L171
	} else {
		goto L185
	}
L179:
	;
	goto L178
L180:
	;
	v1056 = v1047
	v1057 = l1
	goto L181
L181:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+1)))
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+1)))
	if v1061 == int32(0) {
		v1071 = v1061
		v1072 = v1060
		goto L179
	} else {
		goto L183
	}
L182:
	;
	v1071 = v1061
	v1072 = v1060
	goto L179
L183:
	;
	v1064 = int32(1)
	if v1061 == v1060 {
		v1056 = v1056 + v1064
		v1057 = v1057 + v1064
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v1076 = F_strlen(m, v1047)
	mBase = m.M
	v1082 = v1076 - int32(1636608432)
	if v1047&int32(3) != 0 {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v1342 = (v1336 ^ v1328 - base.I32_rotl(v1336, int32(24))) & v1341
	if base.Ui32(v1036) < base.Ui32(v1342) {
		goto L226
	} else {
		goto L227
	}
L187:
	;
	v1314 = int32(14)
	v1316 = v1310 ^ v1311 - base.I32_rotl(v1310, v1314)
	v1320 = v1316 ^ v1309 - base.I32_rotl(v1316, int32(11))
	v1324 = v1320 ^ v1310 - base.I32_rotl(v1320, int32(25))
	v1328 = v1324 ^ v1316 - base.I32_rotl(v1324, int32(16))
	v1332 = v1328 ^ v1320 - base.I32_rotl(v1328, int32(4))
	v1336 = v1332 ^ v1324 - base.I32_rotl(v1332, v1314)
	goto L186
L188:
	;
	switch v1240 - int32(1) {
	case 0:
		v1302 = v1241
		v1303 = v1242
		v1304 = v1243
		goto L215
	case 1:
		v1295 = v1241
		v1296 = v1242
		v1297 = v1243
		goto L216
	case 2:
		v1288 = v1241
		v1289 = v1242
		v1290 = v1243
		goto L217
	case 3:
		v1282 = v1242
		v1283 = v1243
		goto L218
	case 4:
		v1278 = v1242
		v1279 = v1243
		goto L219
	case 5:
		v1272 = v1242
		v1273 = v1243
		goto L220
	case 6:
		v1266 = v1242
		v1267 = v1243
		goto L221
	case 7:
		v1261 = v1243
		goto L222
	case 8:
		v1256 = v1243
		goto L223
	case 9:
		v1251 = v1243
		goto L224
	case 10:
		goto L225
	default:
		v1309 = v1241
		v1310 = v1242
		v1311 = v1243
		goto L187
	}
L189:
	;
	v1191 = v1047
	v1192 = v1076
	v1193 = v1082
	v1194 = v1082
	v1195 = v1082
	goto L212
L190:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1076) {
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if base.Ui32(v1076) < base.Ui32(int32(12)) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	v1239 = v1047
	v1240 = v1076
	v1241 = v1082
	v1242 = v1082
	v1243 = v1082
	goto L188
L194:
	;
	switch v1138 - int32(1) {
	case 0:
		v1188 = v1139
		goto L201
	case 1:
		v1183 = v1139
		goto L202
	case 2:
		goto L203
	case 3:
		v1176 = v1140
		goto L204
	case 4:
		v1173 = v1140
		goto L205
	case 5:
		v1168 = v1140
		goto L206
	case 6:
		goto L207
	case 7:
		v1159 = v1141
		goto L208
	case 8:
		v1154 = v1141
		goto L209
	case 9:
		v1149 = v1141
		goto L210
	case 10:
		goto L211
	default:
		v1309 = v1139
		v1310 = v1140
		v1311 = v1141
		goto L187
	}
L195:
	;
	v1137 = v1047
	v1138 = v1076
	v1139 = v1082
	v1140 = v1082
	v1141 = v1082
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1089 = v1047
	v1090 = v1076
	v1091 = v1082
	v1092 = v1082
	v1093 = v1082
	goto L198
L198:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	v1096 = v1095 + v1092
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1089)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+8))
	v1100 = v1099 + v1093
	v1102 = int32(4)
	v1104 = v1097 + v1091 - v1100 ^ base.I32_rotl(v1100, v1102)
	v1108 = v1096 - v1104 ^ base.I32_rotl(v1104, int32(6))
	v1109 = v1100 + v1096
	v1110 = v1104 + v1109
	v1111 = v1108 + v1110
	v1115 = v1109 - v1108 ^ base.I32_rotl(v1108, int32(8))
	v1119 = v1110 - v1115 ^ base.I32_rotl(v1115, int32(16))
	v1123 = v1111 - v1119 ^ base.I32_rotl(v1119, int32(19))
	v1124 = v1115 + v1111
	v1125 = v1119 + v1124
	v1126 = v1123 + v1125
	v1130 = v1124 - v1123 ^ base.I32_rotl(v1123, v1102)
	v1131 = int32(12)
	v1132 = v1089 + v1131
	v1134 = v1090 - v1131
	if base.Ui32(int32(11)) < base.Ui32(v1134) {
		v1089 = v1132
		v1090 = v1134
		v1091 = v1125
		v1092 = v1126
		v1093 = v1130
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v1137 = v1132
	v1138 = v1134
	v1139 = v1125
	v1140 = v1126
	v1141 = v1130
	goto L194
L200:
	;
	goto L199
L201:
	;
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137))))
	v1309 = v1188 + v1189
	v1310 = v1140
	v1311 = v1141
	goto L187
L202:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+1)))
	v1188 = v1184<<(uint(int32(8))%32) + v1183
	goto L201
L203:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+2)))
	v1183 = v1179<<(uint(int32(16))%32) + v1139
	goto L202
L204:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1137)))
	v1309 = v1177 + v1139
	v1310 = v1176
	v1311 = v1141
	goto L187
L205:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+4)))
	v1176 = v1173 + v1174
	goto L204
L206:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+5)))
	v1173 = v1169<<(uint(int32(8))%32) + v1168
	goto L205
L207:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+6)))
	v1168 = v1164<<(uint(int32(16))%32) + v1140
	goto L206
L208:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1137)))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1137)+4))
	v1309 = v1160 + v1139
	v1310 = v1162 + v1140
	v1311 = v1159
	goto L187
L209:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+8)))
	v1159 = v1155<<(uint(int32(8))%32) + v1154
	goto L208
L210:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+9)))
	v1154 = v1150<<(uint(int32(16))%32) + v1149
	goto L209
L211:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+10)))
	v1149 = v1145<<(uint(int32(24))%32) + v1141
	goto L210
L212:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	v1198 = v1197 + v1194
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1191)))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+8))
	v1202 = v1201 + v1195
	v1204 = int32(4)
	v1206 = v1199 + v1193 - v1202 ^ base.I32_rotl(v1202, v1204)
	v1210 = v1198 - v1206 ^ base.I32_rotl(v1206, int32(6))
	v1211 = v1202 + v1198
	v1212 = v1206 + v1211
	v1213 = v1210 + v1212
	v1217 = v1211 - v1210 ^ base.I32_rotl(v1210, int32(8))
	v1221 = v1212 - v1217 ^ base.I32_rotl(v1217, int32(16))
	v1225 = v1213 - v1221 ^ base.I32_rotl(v1221, int32(19))
	v1226 = v1217 + v1213
	v1227 = v1221 + v1226
	v1228 = v1225 + v1227
	v1232 = v1226 - v1225 ^ base.I32_rotl(v1225, v1204)
	v1233 = int32(12)
	v1234 = v1191 + v1233
	v1236 = v1192 - v1233
	if base.Ui32(int32(11)) < base.Ui32(v1236) {
		v1191 = v1234
		v1192 = v1236
		v1193 = v1227
		v1194 = v1228
		v1195 = v1232
		goto L212
	} else {
		goto L214
	}
L213:
	;
	v1239 = v1234
	v1240 = v1236
	v1241 = v1227
	v1242 = v1228
	v1243 = v1232
	goto L188
L214:
	;
	goto L213
L215:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239))))
	v1309 = v1302 + v1305
	v1310 = v1303
	v1311 = v1304
	goto L187
L216:
	;
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+1)))
	v1302 = v1298<<(uint(int32(8))%32) + v1295
	v1303 = v1296
	v1304 = v1297
	goto L215
L217:
	;
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+2)))
	v1295 = v1291<<(uint(int32(16))%32) + v1288
	v1296 = v1289
	v1297 = v1290
	goto L216
L218:
	;
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+3)))
	v1288 = v1284<<(uint(int32(24))%32) + v1241
	v1289 = v1282
	v1290 = v1283
	goto L217
L219:
	;
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+4)))
	v1282 = v1278 + v1280
	v1283 = v1279
	goto L218
L220:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+5)))
	v1278 = v1274<<(uint(int32(8))%32) + v1272
	v1279 = v1273
	goto L219
L221:
	;
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+6)))
	v1272 = v1268<<(uint(int32(16))%32) + v1266
	v1273 = v1267
	goto L220
L222:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+7)))
	v1266 = v1262<<(uint(int32(24))%32) + v1242
	v1267 = v1261
	goto L221
L223:
	;
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+8)))
	v1261 = v1257<<(uint(int32(8))%32) + v1256
	goto L222
L224:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+9)))
	v1256 = v1252<<(uint(int32(16))%32) + v1251
	goto L223
L225:
	;
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+10)))
	v1251 = v1247<<(uint(int32(24))%32) + v1243
	goto L224
L226:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v1346 = v1036 + v1344
	goto L228
L227:
	;
	v1346 = v1036
	goto L228
L228:
	;
	v1349 = v1341 & (v1036 + int32(1))
	if base.Ui32(v1346-v1342) < base.Ui32(v1037) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1354 = v1024 + v1349<<(uint(int32(4))%32)
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	if v1355 != 0 {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	v1449 = v1037 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1449) {
		goto L248
	} else {
		goto L249
	}
L232:
	;
	v1357 = v1349
	v1360 = int32(0)
	goto L235
L233:
	;
	v1390 = v1349
	v1395 = v1354
	goto L234
L234:
	;
	if v1390 != v1036 {
		goto L242
	} else {
		goto L243
	}
L235:
	;
	v1373 = v1360 + int32(1)
	if int32(151) <= v1373 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1390 = v1385
	v1395 = v1388
	goto L234
L237:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v1378 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1376), base.F64_convert_i64_u(v1378)), float64(0.1)) != 0 {
		goto L44
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1385 = (v1357 + int32(1)) & v1341
	v1388 = v1024 + v1385<<(uint(int32(4))%32)
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)))
	if v1389 != 0 {
		v1357 = v1385
		v1360 = v1373
		goto L235
	} else {
		goto L241
	}
L240:
	;
	goto L239
L241:
	;
	goto L236
L242:
	;
	v1406 = v1390
	v1411 = v1395
	goto L245
L243:
	;
	goto L244
L244:
	;
	v1470 = v1039
	goto L175
L245:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v1424 = v1421 & (v1406 - int32(1))
	v1427 = v1024 + v1424<<(uint(int32(4))%32)
	v1428 = *(*int64)(unsafe.Add(mBase, uint32(v1427)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1411)+8)) = v1428
	v1430 = *(*int64)(unsafe.Add(mBase, uint32(v1427)))
	*(*int64)(unsafe.Add(mBase, uint32(v1411))) = v1430
	if v1424 != v1036 {
		v1406 = v1424
		v1411 = v1427
		goto L245
	} else {
		goto L247
	}
L246:
	;
	goto L244
L247:
	;
	goto L246
L248:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v1454 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1452), base.F64_convert_i64_u(v1454)), float64(0.1)) != 0 {
		goto L44
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1461 = v1024 + v1349<<(uint(int32(4))%32)
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1461)))
	if v1462 != 0 {
		v1036 = v1349
		v1037 = v1449
		v1039 = v1461
		goto L176
	} else {
		goto L252
	}
L251:
	;
	goto L250
L252:
	;
	goto L177
L253:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1485)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+4)) = v1502
	goto L171
L254:
	;
	F_errmsg_internal(m, int32(_a_F_manifest_process_file_0), int32(0))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L60
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_manifest_process_file_1), int32(630), int32(_a_F_manifest_process_file_2))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L60
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errmsg_internal(m, int32(_a_F_manifest_process_file_3), int32(0))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L60
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_manifest_process_file_1), int32(327), int32(_a_F_manifest_process_file_4))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L60
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_markNullableIfNeeded(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 == int32(0) {
		v54 = l0
	} else {
		v12 = v8 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(v8) {
			v18 = l0
			v20 = int32(0)
			for {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v33 = v20 + int32(8)
				if v33 != v8&int32(-8) {
					v18 = v31
					v20 = v33
					continue
				} else {
					break
				}
				break
			}
			if v12 == int32(0) {
				v54 = v31
			} else {
				v37 = v31
				v44 = v37
				v46 = int32(0)
				for {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					v52 = v46 + int32(1)
					if v52 != v12 {
						v44 = v50
						v46 = v52
						continue
					} else {
						break
					}
					break
				}
				v54 = v50
			}
		} else {
			v37 = l0
			v44 = v37
			v46 = int32(0)
			for {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				v52 = v46 + int32(1)
				if v52 != v12 {
					v44 = v50
					v46 = v52
					continue
				} else {
					break
				}
				break
			}
			v54 = v50
		}
	}
	if v7 <= int32(0) {
		return
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
		if v62 == int32(0) {
			return
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
			if v65 < v7 {
				return
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+v7<<(uint(int32(2))%32)-int32(4))))
				if v73 == int32(0) {
					return
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v77 = F_bms_union(m, v76, v73)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v77
						return
					}
				}
			}
		}
	}
}
func F_markcanreach(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = m.T0[v8].(func(*base.Module) int32)(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(101)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v19 != l2 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v17 = v15
	goto L8
L7:
	;
	v17 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v17
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l3
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = v22
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	F_markcanreach(m, l0, v29, l2, l3)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v32 != 0 {
		v26 = v32
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_markst(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v4 = v2 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v4)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 != 0 {
		v7 = v6
		for {
			v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			v10 = v8 | int32(64)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v10)
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
			if v12 != 0 {
				v13 = v12
				for {
					F_markst(m, v13)
					mBase = m.M
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					if v15 != 0 {
						v13 = v15
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
			if v17 != 0 {
				v7 = v17
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_mask_page_lsn_and_checksum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_match_kind(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(v3 == l1)
}
func F_materialize_finished_plan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v42 float64
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v135 float64
	_ = v135
	var v138 int32
	_ = v138
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 float64
	_ = v158
	var v162 float64
	_ = v162
	var v170 float64
	_ = v170
	var v176 float64
	_ = v176
	var v182 float64
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v198 float64
	_ = v198
	var v201 float64
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v12 = F_palloc0(m, int32(72))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(360)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v18
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v19
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
		v42 = float64(0)
		if v29 == v19 {
			v129 = v19
			v135 = v42
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			if v45 <= int32(0) {
				v129 = v19
				v135 = v42
			} else {
				if v45 == int32(1) {
					v102 = int32(0)
					v104 = v19
					v110 = v42
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v102<<(uint(int32(2))%32))))
					v116 = *(*float64)(unsafe.Add(mBase, uint32(v115)+56))
					v117 = *(*float64)(unsafe.Add(mBase, uint32(v115)+64))
					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+38)))
					v129 = v120 ^ int32(1) | v104
					v135 = base.F64_add(v110, base.F64_add(v116, v117))
				} else {
					v51 = int32(0)
					if v51 < v45 {
						v54 = v45
					} else {
						v54 = v51
					}
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
					v64 = int32(0)
					v66 = v19
					v71 = v19
					v72 = v42
					for {
						v73 = int32(2)
						v75 = v59 + v64<<(uint(v73)%32)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v76)+56))
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v76)+64))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
						v82 = *(*float64)(unsafe.Add(mBase, uint32(v81)+56))
						v83 = *(*float64)(unsafe.Add(mBase, uint32(v81)+64))
						v85 = base.F64_add(base.F64_add(v72, base.F64_add(v77, v78)), base.F64_add(v82, v83))
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+38)))
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+38)))
						v91 = base.B2i32(v86&v87 == int32(0)) | v66
						v93 = v64 + v73
						v95 = v71 + v73
						if v95 != v54&int32(2147483646) {
							v64 = v93
							v66 = v91
							v71 = v95
							v72 = v85
							continue
						} else {
							break
						}
						break
					}
					if v54&int32(1) == int32(0) {
						v129 = v91
						v135 = v85
					} else {
						v102 = v93
						v104 = v91
						v110 = v85
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v102<<(uint(int32(2))%32))))
						v116 = *(*float64)(unsafe.Add(mBase, uint32(v115)+56))
						v117 = *(*float64)(unsafe.Add(mBase, uint32(v115)+64))
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+38)))
						v129 = v120 ^ int32(1) | v104
						v135 = base.F64_add(v110, base.F64_add(v116, v117))
					}
				}
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v135
		v138 = v129 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(15)))) = uint8(v138)
		v140 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v141 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		v142 = base.F64_sub(v140, v141)
		*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v142
		v144 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
		v145 = base.F64_sub(v144, v141)
		*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v145
		v147 = int32(24)
		v148 = v9 + v147
		v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v150 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v155 = *(*int32)(unsafe.Add(mBase, _c_F_materialize_finished_plan[0]))
		*(*float64)(unsafe.Add(mBase, uint32(v148)+32)) = v150
		v158 = *(*float64)(unsafe.Add(mBase, _c_F_materialize_finished_plan[1]))
		v162 = base.F64_add(base.F64_mul(base.F64_add(v158, v158), v150), base.F64_sub(v145, v142))
		v170 = base.F64_mul(v150, base.F64_convert_i32_u((v151+int32(7))&int32(-8)+v147))
		if base.F64_gt(v170, base.F64_convert_i32_u(v155<<(uint(int32(10))%32))) != 0 {
			v176 = *(*float64)(unsafe.Add(mBase, _c_F_materialize_finished_plan[2]))
			v182 = base.F64_add(base.F64_mul(v176, base.F64_ceil(base.F64_mul(v170, float64(0.0001220703125)))), v162)
		} else {
			v182 = v162
		}
		v184 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_materialize_finished_plan[3])))
		*(*float64)(unsafe.Add(mBase, uint32(v148)+56)) = base.F64_add(v142, v182)
		*(*float64)(unsafe.Add(mBase, uint32(v148)+48)) = v142
		*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v149 + (v184 ^ int32(1))
		v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v192
		v194 = *(*float64)(unsafe.Add(mBase, uint32(v9)+72))
		v195 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = base.F64_add(v194, v195)
		v198 = *(*float64)(unsafe.Add(mBase, uint32(v9)+80))
		*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = base.F64_add(v195, v198)
		v201 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v201
		v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v204 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+36)) = uint8(v204)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v203
		v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)) = uint8(v207)
		m.G0 = v9 + int32(96)
		return v12
	}
}
func F_maybe_start_bgworkers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v243 int32
	_ = v243
	v1 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[0])) = uint8(v18)
	*(*uint8)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[1])) = uint8(v1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[2])))
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[3]))
	if base.B2i32(v26 == int32(0))|base.B2i32(v26 == int32(_a_F_maybe_start_bgworkers_0)) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = v26
	v40 = v1
	v42 = int64(0)
	goto L4
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v46 = v35 - int32(20)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 != 0 {
		v232 = v40
		v234 = v42
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[0])) = uint8(v243)
	goto L1
L6:
	;
	goto L5
L7:
	;
	if v44 != int32(_a_F_maybe_start_bgworkers_0) {
		v35 = v44
		v40 = v232
		v42 = v234
		goto L4
	} else {
		goto L67
	}
L8:
	;
	v49 = v35 - int32(1480)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35-int32(4)))))
	if v52 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_ForgetBackgroundWorker(m, v49)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v58 = v35 - int32(16)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	if v59 == int64(0) {
		v114 = v42
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v232 = v40
	v234 = v42
	goto L7
L14:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v35-int32(1284))))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[4]))
	switch v119 {
	case 0, 1, 2:
		goto L29
	case 3:
		goto L30
	case 4:
		goto L31
	default:
		v232 = v40
		v234 = v114
		goto L7
	}
L15:
	;
	v63 = v35 - int32(1280)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 == int32(-1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v35-int32(24))))
	F_ForgetBackgroundWorker(m, v49)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v42 == int64(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v69 == int32(0) {
		v232 = v40
		v234 = v42
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v75 = F_kill(m, v69, int32(10))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v232 = v40
	v234 = v42
	goto L7
L22:
	;
	v82 = m.G0
	v83 = int32(16)
	v84 = v82 - v83
	m.G0 = v84
	F_gettimeofday(m, v84)
	mBase = m.M
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
	v88 = int64(*(*int32)(unsafe.Add(mBase, uint32(v84)+8)))
	m.G0 = v84 + v83
	goto L25
L23:
	;
	v99 = v64
	v100 = v42
	v101 = v59
	goto L24
L24:
	;
	goto L26
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	v99 = v97
	v100 = v88 + v87*int64(1000000) - int64(946684800000000)
	v101 = v98
	goto L24
L26:
	;
	if base.I64_extend_i32_s(v99*int32(1000))*int64(1000) <= v100-v101 {
		v114 = v100
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v110 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[1])) = uint8(v110)
	v232 = v40
	v234 = v100
	goto L7
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(0)
	v127 = F_AssignPostmasterChildSlot(m, int32(5))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L37
	}
L29:
	;
	if v117 != 0 {
		v232 = v40
		v234 = v114
		goto L7
	} else {
		goto L34
	}
L30:
	;
	if base.Ui32(v117) < base.Ui32(int32(2)) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	if base.Ui32(v117) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v232 = v40
	v234 = v114
	goto L7
L33:
	;
	v232 = v40
	v234 = v114
	goto L7
L34:
	;
	goto L28
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v169
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_start_bgworkers[5]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1472))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1460))
	*(*int32)(unsafe.Add(mBase, uint32(v212+v213*int32(1480))+20)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1456))
	if v219 != 0 {
		goto L62
	} else {
		goto L63
	}
L36:
	;
	v193 = m.G0
	v194 = int32(16)
	v195 = v193 - v194
	m.G0 = v195
	F_gettimeofday(m, v195)
	mBase = m.M
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	v199 = int64(*(*int32)(unsafe.Add(mBase, uint32(v195)+8)))
	m.G0 = v195 + v194
	goto L61
L37:
	;
	if v127 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v133 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+16)) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v49
	v156 = F_errstart(m, int32(14), v149)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L46
	}
L41:
	;
	if v133 == int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(_a_F_maybe_start_bgworkers_1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_maybe_start_bgworkers_2), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_maybe_start_bgworkers_3), int32(_a_F_maybe_start_bgworkers_4), int32(_a_F_maybe_start_bgworkers_5))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	goto L36
L46:
	;
	if v156 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
	F_errmsg_internal(m, int32(_a_F_maybe_start_bgworkers_6), v15)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v169 = F_postmaster_child_launch(m, int32(5), v49)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L52
	}
L50:
	;
	F_errfinish(m, int32(_a_F_maybe_start_bgworkers_3), int32(_a_F_maybe_start_bgworkers_7), int32(_a_F_maybe_start_bgworkers_5))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v169 != int32(-1) {
		goto L35
	} else {
		goto L53
	}
L53:
	;
	v175 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	if v175 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_errmsg(m, int32(_a_F_maybe_start_bgworkers_8), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v186 = F_ReleasePostmasterChildSlot(m, v127)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L60
	}
L58:
	;
	F_errfinish(m, int32(_a_F_maybe_start_bgworkers_3), int32(_a_F_maybe_start_bgworkers_9), int32(_a_F_maybe_start_bgworkers_5))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L36
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v199 + v198*int64(1000000) - int64(946684800000000)
	goto L6
L62:
	;
	v221 = F_kill(m, v219, int32(10))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v224 = v40 + int32(1)
	if int32(99) < v224 {
		goto L6
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v232 = v224
	v234 = v114
	goto L7
L67:
	;
	goto L1
}
func F_mbtowc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v104 int32
	_ = v104
	if l1 == int32(0) {
		return int32(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v11 = base.I32_extend8_s(v10)
		if int32(0) <= v11 {
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
			} else {
			}
			return base.B2i32(v11 != int32(0))
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_mbtowc[0]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == int32(0) {
				if l0 == int32(0) {
					v104 = int32(1)
					return v104
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11 & int32(_a_F_mbtowc_0)
					return int32(1)
				}
			} else {
				v32 = v10 - int32(194)
				if base.Ui32(int32(50)) < base.Ui32(v32) {
					*(*int32)(unsafe.Add(mBase, _c_F_mbtowc[1])) = int32(25)
					v104 = int32(-1)
					return v104
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v37 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_mbtowc[2])))
					if base.Ui32(int32(7)) < base.Ui32(v37-int32(16)|(v37+v42>>(uint(int32(26))%32))) {
						*(*int32)(unsafe.Add(mBase, _c_F_mbtowc[1])) = int32(25)
						v104 = int32(-1)
						return v104
					} else {
						v53 = v35 - int32(128) | v42<<(uint(int32(6))%32)
						if int32(0) <= v53 {
							if l0 == int32(0) {
								v104 = int32(2)
								return v104
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v53
								return int32(2)
							}
						} else {
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
							v64 = v62 - int32(128)
							if base.Ui32(int32(63)) < base.Ui32(v64) {
								*(*int32)(unsafe.Add(mBase, _c_F_mbtowc[1])) = int32(25)
								v104 = int32(-1)
								return v104
							} else {
								v68 = v53 << (uint(int32(6)) % 32)
								v69 = v64 | v68
								if int32(0) <= v68 {
									if l0 == int32(0) {
										v104 = int32(3)
										return v104
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v69
										return int32(3)
									}
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
									v80 = v78 - int32(128)
									if base.Ui32(int32(63)) < base.Ui32(v80) {
										*(*int32)(unsafe.Add(mBase, _c_F_mbtowc[1])) = int32(25)
										v104 = int32(-1)
										return v104
									} else {
										if l0 == int32(0) {
											v104 = int32(4)
											return v104
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80 | v69<<(uint(int32(6))%32)
											return int32(4)
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
func F_mcv_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 float32
	_ = v106
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v131 int32
	_ = v131
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	v9 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 == int32(0) {
		v140 = v9
		v141 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v141
	m.G0 = v14 + int32(80)
	return v140
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if int32(0) < v74 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l3
	goto L3
L5:
	;
	v58 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L16
	}
L6:
	;
	v33 = v16
	goto L8
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 == int32(0) {
		v140 = v9
		v141 = v9
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v37 = F_get_attstatsslot(m, v14+int32(44), v33, int32(1), int32(0), int32(3))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L13
	}
L9:
	;
	v25 = F_get_func_leakproof(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return float64(0)
L11:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = v31
	goto L8
L13:
	;
	if v37 == int32(0) {
		v140 = v9
		v141 = v9
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)) = uint8(v41)
	v45 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+26)) = uint16(v45)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v41)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
	if l4 == v41 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l3
	goto L3
L16:
	;
	if v58 == int32(0) {
		v140 = v9
		v141 = v9
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v62 = F_get_func_name(m, v22)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v62
	F_errmsg_internal(m, int32(_a_F_mcv_selectivity_0), v14)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_mcv_selectivity_1), int32(_a_F_mcv_selectivity_2), int32(_a_F_mcv_selectivity_3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v140 = v9
	v141 = v9
	goto L1
L21:
	;
	v78 = int32(0)
	v86 = v9
	v87 = v9
	goto L24
L22:
	;
	v125 = v9
	v126 = v9
	goto L23
L23:
	;
	F_free_attstatsslot(m, v14+int32(44))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L38
	}
L24:
	;
	v90 = v78 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90+v91)))
	if l4 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v125 = v111
	v126 = v112
	goto L23
L26:
	;
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v96)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = m.T0[v101].(func(*base.Module, int32) int32)(m, v14+int32(8))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L30
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v93
	goto L26
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v93
	goto L26
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v104+v90)))
	v107 = base.F64_promote_f32(v106)
	if v102 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v109 = base.F64_add(v86, v107)
	goto L33
L32:
	;
	v109 = v86
	goto L33
L33:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
	if v110 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v111 = v86
	goto L36
L35:
	;
	v111 = v109
	goto L36
L36:
	;
	v112 = base.F64_add(v87, v107)
	v114 = v78 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v114 < v115 {
		v78 = v114
		v86 = v111
		v87 = v112
		goto L24
	} else {
		goto L37
	}
L37:
	;
	goto L25
L38:
	;
	v140 = v125
	v141 = v126
	goto L1
}
func F_md5_bytea(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13944(m, l0, int32(_a_F_md5_bytea_0), int32(71))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	v4 = int32(0)
	if base.B2i32(l0&int32(3) == v4)|base.B2i32(l2 == v4) != 0 {
		v34 = l0
		v36 = l2
		v37 = base.B2i32(l2 != v4)
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v86 = v79
	v88 = v81
	goto L19
L3:
	;
	if v37 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v17 = l0
	v19 = l2
	goto L5
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 == l1&int32(255) {
		v79 = v17
		v81 = v19
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v34 = v29
	v36 = v25
	v37 = v27
	goto L3
L7:
	;
	v24 = int32(1)
	v25 = v19 - v24
	v26 = int32(0)
	v27 = base.B2i32(v25 != v26)
	v29 = v17 + v24
	if v29&int32(3) == v26 {
		v34 = v29
		v36 = v25
		v37 = v27
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v25 != 0 {
		v17 = v29
		v19 = v25
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v42 = l1 & int32(255)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if base.B2i32(v42 == v43)|base.B2i32(base.Ui32(v36) < base.Ui32(int32(4))) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = v34
	v54 = v36
	goto L14
L12:
	;
	v72 = v34
	v74 = v36
	goto L13
L13:
	;
	if v74 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v59 = v58 ^ v42*int32(16843009)
	v62 = int32(-2139062144)
	if (int32(16843008)-v59|v59)&v62 != v62 {
		v79 = v52
		v81 = v54
		goto L2
	} else {
		goto L16
	}
L15:
	;
	v72 = v67
	v74 = v69
	goto L13
L16:
	;
	v66 = int32(4)
	v67 = v52 + v66
	v69 = v54 - v66
	if base.Ui32(int32(3)) < base.Ui32(v69) {
		v52 = v67
		v54 = v69
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v79 = v72
	v81 = v74
	goto L2
L19:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if l1&int32(255) == v91 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L1
L21:
	;
	return v86
L22:
	;
	goto L23
L23:
	;
	v94 = int32(1)
	v97 = v88 - v94
	if v97 != 0 {
		v86 = v86 + v94
		v88 = v97
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
}
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = l0 + l2
	if base.Ui32(l1-v7) <= base.Ui32(int32(0)-l2<<(uint(int32(1))%32)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(int32(512)) <= base.Ui32(l2) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v185 = (l0 ^ l1) & int32(3)
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L55
	} else {
		goto L56
	}
L6:
	;
	return
L7:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v20 = l0 + l2
	if (l0^l1)&int32(3) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	base.MemoryCopy(m, l0, l1, l2)
	goto L12
L11:
	;
	goto L12
L12:
	;
	goto L6
L13:
	;
	if base.Ui32(v152) < base.Ui32(v20) {
		goto L47
	} else {
		goto L48
	}
L14:
	;
	if l0&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v20) < base.Ui32(int32(4)) {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	v56 = v20 & int32(-4)
	if base.Ui32(v20) < base.Ui32(int32(64)) {
		v106 = v50
		v107 = v51
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v50 = l1
	v51 = l0
	goto L17
L19:
	;
	goto L20
L20:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v50 = l1
	v51 = l0
	goto L17
L22:
	;
	goto L23
L23:
	;
	v33 = l1
	v34 = l0
	goto L24
L24:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v38)
	v40 = int32(1)
	v41 = v33 + v40
	v43 = v34 + v40
	if v43&int32(3) == int32(0) {
		v50 = v41
		v51 = v43
		goto L17
	} else {
		goto L26
	}
L25:
	;
	v50 = v41
	v51 = v43
	goto L17
L26:
	;
	if base.Ui32(v43) < base.Ui32(v20) {
		v33 = v41
		v34 = v43
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	if base.Ui32(v56) <= base.Ui32(v107) {
		v151 = v106
		v152 = v107
		goto L13
	} else {
		goto L34
	}
L29:
	;
	v60 = v56 + int32(-64)
	if base.Ui32(v60) < base.Ui32(v51) {
		v106 = v50
		v107 = v51
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v63 = v50
	v64 = v51
	goto L31
L31:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v63)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+48)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+52)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v63)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+56)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v63)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+60)) = v98
	v100 = int32(-64)
	v101 = v63 - v100
	v103 = v64 - v100
	if base.Ui32(v103) <= base.Ui32(v60) {
		v63 = v101
		v64 = v103
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v106 = v101
	v107 = v103
	goto L28
L33:
	;
	goto L32
L34:
	;
	v113 = v106
	v114 = v107
	goto L35
L35:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v118
	v120 = int32(4)
	v121 = v113 + v120
	v123 = v114 + v120
	if base.Ui32(v123) < base.Ui32(v56) {
		v113 = v121
		v114 = v123
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v151 = v121
	v152 = v123
	goto L13
L37:
	;
	goto L36
L38:
	;
	v151 = l1
	v152 = l0
	goto L13
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(l2) < base.Ui32(int32(4)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v151 = l1
	v152 = l0
	goto L13
L42:
	;
	goto L43
L43:
	;
	v132 = l1
	v133 = l0
	goto L44
L44:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)) = uint8(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+2)) = uint8(v141)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+3)) = uint8(v143)
	v145 = int32(4)
	v146 = v132 + v145
	v148 = v133 + v145
	if base.Ui32(v148) <= base.Ui32(v20-int32(4)) {
		v132 = v146
		v133 = v148
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v151 = v146
	v152 = v148
	goto L13
L46:
	;
	goto L45
L47:
	;
	v158 = v151
	v159 = v152
	goto L50
L48:
	;
	goto L49
L49:
	;
	goto L6
L50:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v163)
	v165 = int32(1)
	v168 = v159 + v165
	if v168 != v20 {
		v158 = v158 + v165
		v159 = v168
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	goto L51
L53:
	;
	if v287 == int32(0) {
		goto L1
	} else {
		goto L85
	}
L54:
	;
	if base.Ui32(v265) <= base.Ui32(int32(3)) {
		v285 = v263
		v286 = v264
		v287 = v265
		goto L53
	} else {
		goto L81
	}
L55:
	;
	if v185 != 0 {
		v285 = l0
		v286 = l1
		v287 = l2
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v185 != 0 {
		v247 = l2
		goto L64
	} else {
		goto L65
	}
L58:
	;
	if l0&int32(3) == int32(0) {
		v263 = l0
		v264 = l1
		v265 = l2
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v191 = l0
	v192 = l1
	v193 = l2
	goto L60
L60:
	;
	if v193 == int32(0) {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v263 = v205
	v264 = v201
	v265 = v203
	goto L54
L62:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v198)
	v200 = int32(1)
	v201 = v192 + v200
	v203 = v193 - v200
	v205 = v191 + v200
	if v205&int32(3) != 0 {
		v191 = v205
		v192 = v201
		v193 = v203
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	if v247 == int32(0) {
		goto L1
	} else {
		goto L77
	}
L65:
	;
	if v7&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v212 = l2
	goto L69
L67:
	;
	v227 = l2
	goto L68
L68:
	;
	if base.Ui32(v227) <= base.Ui32(int32(3)) {
		v247 = v227
		goto L64
	} else {
		goto L73
	}
L69:
	;
	if v212 == int32(0) {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	v227 = v218
	goto L68
L71:
	;
	v218 = v212 - int32(1)
	v219 = l0 + v218
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v218))))
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v221)
	if v219&int32(3) != 0 {
		v212 = v218
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v234 = v227
	goto L74
L74:
	;
	v238 = v234 - int32(4)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1+v238)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v238))) = v241
	if base.Ui32(int32(3)) < base.Ui32(v238) {
		v234 = v238
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v247 = v238
	goto L64
L76:
	;
	goto L75
L77:
	;
	v254 = v247
	goto L78
L78:
	;
	v258 = v254 - int32(1)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v258))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v258))) = uint8(v261)
	if v258 != 0 {
		v254 = v258
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L1
L80:
	;
	goto L79
L81:
	;
	v270 = v263
	v271 = v264
	v272 = v265
	goto L82
L82:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v275
	v277 = int32(4)
	v278 = v271 + v277
	v280 = v270 + v277
	v282 = v272 - v277
	if base.Ui32(int32(3)) < base.Ui32(v282) {
		v270 = v280
		v271 = v278
		v272 = v282
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v285 = v280
	v286 = v278
	v287 = v282
	goto L53
L84:
	;
	goto L83
L85:
	;
	v292 = v285
	v293 = v286
	v294 = v287
	goto L86
L86:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v297)
	v299 = int32(1)
	v304 = v294 - v299
	if v304 != 0 {
		v292 = v292 + v299
		v293 = v293 + v299
		v294 = v304
		goto L86
	} else {
		goto L88
	}
L87:
	;
	goto L1
L88:
	;
	goto L87
}
func F_minmax_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l1 - int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0+v14<<(uint(int32(2))%32))+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if l2 != v20 {
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = l2
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v23
	} else {
	}
	v39 = v19 + l3*int32(28) - int32(24)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v40 == int32(0) {
		v43 = int32(4)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v14<<(uint(int32(2))%32))))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		v59 = v50 + v51<<(uint(v43)%32) + v14*int32(100) + int32(88)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
		v62 = F_SearchSysCache4(m, v43, v49, v60, l2, base.I32_extend16_s(l3))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			if v62 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v90
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_errmsg_internal(m, int32(_a_F_minmax_get_strategy_procinfo_0), v11)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_minmax_get_strategy_procinfo_1), int32(301), int32(_a_F_minmax_get_strategy_procinfo_2))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v70 = F_SysCacheGetAttrNotNull(m, int32(4), v62, int32(7))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v62)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = F_get_opcode(m, v70)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_fmgr_info_cxt(m, v74, v39, v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return v39
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v39
	}
}
func F_minmax_multi_get_procinfo(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13850(m, l0, l1, int32(_a_F_minmax_multi_get_procinfo_0), int32(2886), int32(_a_F_minmax_multi_get_procinfo_1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_mkVoidAffix(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	v4 = int32(0)
	if l1 != 0 {
		v15 = l2
	} else {
		v15 = v4
	}
	if l1 != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = v18
		v20 = int32(12)
	} else {
		v19 = l2
		v20 = int32(16)
	}
	v22 = F_palloc0(m, int32(16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(3)
		v26 = l0 + v20
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v22
		if v19 <= v15 {
			return
		} else {
			v31 = v19 - v15
			v32 = int32(3)
			v33 = v31 & v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v35 = int32(0)
			v38 = v19 + (v15 ^ int32(-1))
			if base.Ui32(v32) <= base.Ui32(v38) {
				v44 = v15
				v46 = v35
				v55 = v4
				for {
					v58 = v34 + v44*int32(24)
					v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+5)))
					v60 = int32(_a_F_mkVoidAffix_0)
					v62 = int32(0)
					v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+29)))
					v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+53)))
					v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+77)))
					v82 = v46 + base.B2i32(v59&v60 == v62) + base.B2i32(v65&v60 == v62) + base.B2i32(v71&v60 == v62) + base.B2i32(v77&v60 == v62)
					v83 = int32(4)
					v84 = v44 + v83
					v86 = v55 + v83
					if v86 != v31&int32(-4) {
						v44 = v84
						v46 = v82
						v55 = v86
						continue
					} else {
						break
					}
					break
				}
				if v33 == int32(0) {
					v133 = v82
				} else {
					v91 = v84
					v93 = v82
					v104 = v91
					v106 = v93
					v110 = v4
					for {
						v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v104*int32(24))+5)))
						v124 = v106 + base.B2i32(v119&int32(_a_F_mkVoidAffix_0) == int32(0))
						v125 = int32(1)
						v128 = v110 + v125
						if v128 != v33 {
							v104 = v104 + v125
							v106 = v124
							v110 = v128
							continue
						} else {
							break
						}
						break
					}
					v133 = v124
				}
			} else {
				v91 = v15
				v93 = v35
				v104 = v91
				v106 = v93
				v110 = v4
				for {
					v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v104*int32(24))+5)))
					v124 = v106 + base.B2i32(v119&int32(_a_F_mkVoidAffix_0) == int32(0))
					v125 = int32(1)
					v128 = v110 + v125
					if v128 != v33 {
						v104 = v104 + v125
						v106 = v124
						v110 = v128
						continue
					} else {
						break
					}
					break
				}
				v133 = v124
			}
			if v133 == int32(0) {
				return
			} else {
				v146 = v133 << (uint(int32(2)) % 32)
				if base.Ui32(int32(1025)) <= base.Ui32(v146) {
					v149 = F_palloc0(m, v146)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						v168 = v149
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v168
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v172 | v133<<(uint(int32(8))%32)
						if v38 == int32(0) {
							v241 = int32(0)
							v245 = v15
							v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v256 = v253 + v245*int32(24)
							v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+5)))
							if v257&int32(_a_F_mkVoidAffix_0) != 0 {
							} else {
								v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v260+v241<<(uint(int32(2))%32)))) = v256
							}
						} else {
							v184 = int32(0)
							v187 = v184
							v189 = v184
							v191 = v15
							for {
								v200 = v191 * int32(24)
								v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v202 = v200 + v201
								v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+5)))
								if v203&int32(_a_F_mkVoidAffix_0) == int32(0) {
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v208+v187<<(uint(int32(2))%32)))) = v202
									v215 = v187 + int32(1)
								} else {
									v215 = v187
								}
								v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v217 = v216 + v200
								v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+29)))
								if v218&int32(_a_F_mkVoidAffix_0) == int32(0) {
									v223 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v223+v215<<(uint(int32(2))%32)))) = v217 + int32(24)
									v232 = v215 + int32(1)
								} else {
									v232 = v215
								}
								v233 = int32(2)
								v234 = v191 + v233
								v236 = v189 + v233
								if v236 != v31&int32(-2) {
									v187 = v232
									v189 = v236
									v191 = v234
									continue
								} else {
									break
								}
								break
							}
							if v31&int32(1) == int32(0) {
							} else {
								v241 = v232
								v245 = v234
								v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v256 = v253 + v245*int32(24)
								v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+5)))
								if v257&int32(_a_F_mkVoidAffix_0) != 0 {
								} else {
									v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v260+v241<<(uint(int32(2))%32)))) = v256
								}
							}
						}
						return
					}
				} else {
					v154 = (v146 + int32(7)) & int32(4088)
					v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					if base.Ui32(v154) <= base.Ui32(v155) {
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						v162 = v155
						v163 = v157
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v162 - v154
						*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v163 + v154
						v168 = v163
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v168
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v172 | v133<<(uint(int32(8))%32)
						if v38 == int32(0) {
							v241 = int32(0)
							v245 = v15
							v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v256 = v253 + v245*int32(24)
							v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+5)))
							if v257&int32(_a_F_mkVoidAffix_0) != 0 {
							} else {
								v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v260+v241<<(uint(int32(2))%32)))) = v256
							}
						} else {
							v184 = int32(0)
							v187 = v184
							v189 = v184
							v191 = v15
							for {
								v200 = v191 * int32(24)
								v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v202 = v200 + v201
								v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+5)))
								if v203&int32(_a_F_mkVoidAffix_0) == int32(0) {
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v208+v187<<(uint(int32(2))%32)))) = v202
									v215 = v187 + int32(1)
								} else {
									v215 = v187
								}
								v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v217 = v216 + v200
								v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+29)))
								if v218&int32(_a_F_mkVoidAffix_0) == int32(0) {
									v223 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v223+v215<<(uint(int32(2))%32)))) = v217 + int32(24)
									v232 = v215 + int32(1)
								} else {
									v232 = v215
								}
								v233 = int32(2)
								v234 = v191 + v233
								v236 = v189 + v233
								if v236 != v31&int32(-2) {
									v187 = v232
									v189 = v236
									v191 = v234
									continue
								} else {
									break
								}
								break
							}
							if v31&int32(1) == int32(0) {
							} else {
								v241 = v232
								v245 = v234
								v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v256 = v253 + v245*int32(24)
								v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+5)))
								if v257&int32(_a_F_mkVoidAffix_0) != 0 {
								} else {
									v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v260+v241<<(uint(int32(2))%32)))) = v256
								}
							}
						}
						return
					} else {
						v158 = int32(_a_F_mkVoidAffix_1)
						v160 = F_palloc0(m, v158)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							v162 = v158
							v163 = v160
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v162 - v154
							*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v163 + v154
							v168 = v163
							*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v168
							v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v172 | v133<<(uint(int32(8))%32)
							if v38 == int32(0) {
								v241 = int32(0)
								v245 = v15
								v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v256 = v253 + v245*int32(24)
								v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+5)))
								if v257&int32(_a_F_mkVoidAffix_0) != 0 {
								} else {
									v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v260+v241<<(uint(int32(2))%32)))) = v256
								}
							} else {
								v184 = int32(0)
								v187 = v184
								v189 = v184
								v191 = v15
								for {
									v200 = v191 * int32(24)
									v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v202 = v200 + v201
									v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+5)))
									if v203&int32(_a_F_mkVoidAffix_0) == int32(0) {
										v208 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v208+v187<<(uint(int32(2))%32)))) = v202
										v215 = v187 + int32(1)
									} else {
										v215 = v187
									}
									v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v217 = v216 + v200
									v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+29)))
									if v218&int32(_a_F_mkVoidAffix_0) == int32(0) {
										v223 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v223+v215<<(uint(int32(2))%32)))) = v217 + int32(24)
										v232 = v215 + int32(1)
									} else {
										v232 = v215
									}
									v233 = int32(2)
									v234 = v191 + v233
									v236 = v189 + v233
									if v236 != v31&int32(-2) {
										v187 = v232
										v189 = v236
										v191 = v234
										continue
									} else {
										break
									}
									break
								}
								if v31&int32(1) == int32(0) {
								} else {
									v241 = v232
									v245 = v234
									v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v256 = v253 + v245*int32(24)
									v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+5)))
									if v257&int32(_a_F_mkVoidAffix_0) != 0 {
									} else {
										v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v260+v241<<(uint(int32(2))%32)))) = v256
									}
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
func F_mode_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int64
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v2 = int32(0)
	v10 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v2
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v160
L2:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v155)
	v160 = int32(0)
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	if v20 == int64(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v25 = v23 + int32(76)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v30 = F_get_opcode(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v39 = v23
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+58)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	F_fmgr_info_cxt(m, v30, v25, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v39 = v38
	goto L7
L11:
	;
	v51 = int32(0)
	v55 = v51
	v56 = v51
	v58 = v2
	v61 = v2
	v62 = v10
	v63 = v10
	goto L17
L12:
	;
	F_tuplesort_performsort(m, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_tuplesort_rescan(m, v40)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v47)
	goto L11
L16:
	;
	goto L11
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v65 = int32(1)
	v73 = F_tuplesort_getdatum(m, v64, v65, v65, v14+int32(12), v14+int32(11), v14+int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L22
	}
L19:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_mode_final[0]))
	if v139 == int32(0) {
		v55 = v132
		v56 = v133
		v58 = v134
		v61 = v135
		v62 = v136
		v63 = v137
		goto L17
	} else {
		goto L53
	}
L20:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_pfree(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L52
	}
L21:
	;
	if v41&int32(1) != 0 {
		v132 = v55
		v133 = v56
		v134 = v58
		v135 = v61
		v136 = v121
		v137 = v122
		goto L19
	} else {
		goto L51
	}
L22:
	;
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v75 != 0 {
		goto L17
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if (v56|v41)&int32(1) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L26:
	;
	if v62 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = int64(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v132 = v82
	v133 = int32(1)
	v134 = v82
	v135 = v80
	v136 = v79
	v137 = v79
	goto L19
L28:
	;
	goto L29
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v83 != v61 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if (v56|v41)&int32(1) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v87 = F_FunctionCall2Coll(m, v25, v85, v86, v58)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	if v87 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v56 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v121 = v62
	v122 = v63 + int64(1)
	goto L21
L35:
	;
	goto L36
L36:
	;
	v94 = v62 + int64(1)
	if v94 <= v63 {
		v121 = v94
		v122 = v63
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v96 = int32(1)
	if v41&v96 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = v58
	v133 = v96
	v134 = v58
	v135 = v61
	v136 = v94
	v137 = v94
	goto L19
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v55)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v125 = v58
	v126 = v96
	v127 = v94
	v128 = v94
	goto L20
L42:
	;
	F_pfree(m, v58)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v132 = v55
	v133 = int32(0)
	v134 = v111
	v135 = v110
	v136 = int64(1)
	v137 = v63
	goto L19
L45:
	;
	goto L44
L46:
	;
	F_pfree(m, v58)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v63 != int64(0) {
		v160 = v55
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L2
L51:
	;
	v125 = v55
	v126 = v56
	v127 = v121
	v128 = v122
	goto L20
L52:
	;
	v132 = v125
	v133 = v126
	v134 = v58
	v135 = v61
	v136 = v127
	v137 = v128
	goto L19
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v55 = v132
	v56 = v133
	v58 = v134
	v61 = v135
	v62 = v136
	v63 = v137
	goto L17
}
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = int32(0)
	if base.B2i32(l0 == v4)|base.B2i32(l1 == v4) != 0 {
		v19 = int32(0)
		return v19
	} else {
		v13 = base.I64_extend_i32_u(l1) * base.I64_extend_i32_u(l0)
		if base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(32))%64))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_mul_size_0), int32(0))
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_mul_size_1), int32(521), int32(_a_F_mul_size_2))
						v38 = m.ExcPending
						if v38 != 0 {
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
			v19 = base.I32_wrap_i64(v13)
			return v19
		}
	}
}
func F_multirangesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
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
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 float64
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 float64
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v126 float64
	_ = v126
	var v129 float64
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 float64
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
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
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 float32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 float32
	_ = v266
	var v268 int32
	_ = v268
	var v270 float64
	_ = v270
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
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
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
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
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int64
	_ = v393
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v433 float64
	_ = v433
	var v434 int32
	_ = v434
	var v438 float64
	_ = v438
	var v439 int32
	_ = v439
	var v444 float64
	_ = v444
	var v445 int32
	_ = v445
	var v451 float64
	_ = v451
	var v452 int32
	_ = v452
	var v457 float64
	_ = v457
	var v458 int32
	_ = v458
	var v463 float64
	_ = v463
	var v464 int32
	_ = v464
	var v469 float64
	_ = v469
	var v470 int32
	_ = v470
	var v475 float64
	_ = v475
	var v476 int32
	_ = v476
	var v481 float64
	_ = v481
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 float64
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v501 float64
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 float64
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v532 float64
	_ = v532
	var v533 int32
	_ = v533
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v583 float64
	_ = v583
	var v584 int32
	_ = v584
	var v587 float64
	_ = v587
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v617 float64
	_ = v617
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v653 float64
	_ = v653
	var v656 float64
	_ = v656
	var v666 float64
	_ = v666
	var v673 float64
	_ = v673
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v712 float64
	_ = v712
	var v718 float64
	_ = v718
	var v721 float64
	_ = v721
	var v735 float64
	_ = v735
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v771 float64
	_ = v771
	var v774 float64
	_ = v774
	var v790 float64
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 float64
	_ = v795
	var v813 float64
	_ = v813
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	v10 = float64(0)
	v16 = m.G0
	v18 = v16 - int32(176)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	v32 = F_get_restriction_variable(m, v23, v22, v21, v18+int32(44), v18+int32(40), v18+int32(39))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v818 = F_Float8GetDatum(m, v813)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L2
	} else {
		goto L241
	}
L2:
	;
	return int32(0)
L3:
	;
	if v32 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = float64(0.005)
	if v20 <= int32(_a_F_multirangesel_0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	goto L6
L6:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81 != int32(7) {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v813 = v79
	goto L1
L8:
	;
	v79 = v76
	goto L7
L9:
	;
	v76 = float64(0.3333333333333333)
	goto L8
L10:
	;
	v79 = float64(0.01)
	goto L7
L11:
	;
	if base.B2i32(v20 == int32(3585))|base.B2i32(v20 == int32(4035)) != 0 {
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v44 = v20 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v44) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v20-int32(_a_F_multirangesel_1)) < base.Ui32(int32(6)) {
		goto L9
	} else {
		goto L18
	}
L15:
	;
	v48 = int32(1) << (uint(v44) % 32)
	if v48&int32(_a_F_multirangesel_2) != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v48&int32(_a_F_multirangesel_3) == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v76 = v40
	goto L8
L18:
	;
	if base.Ui32(v20-int32(_a_F_multirangesel_4)) < base.Ui32(int32(2)) {
		v76 = v40
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if v20 != int32(_a_F_multirangesel_5) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	goto L10
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v84 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+24)))
	if v130 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v85].(func(*base.Module, int32))(m, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v90 = float64(0.005)
	if v20 <= int32(_a_F_multirangesel_0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	v813 = v129
	goto L1
L30:
	;
	v129 = v126
	goto L29
L31:
	;
	v126 = float64(0.3333333333333333)
	goto L30
L32:
	;
	v129 = float64(0.01)
	goto L29
L33:
	;
	if base.B2i32(v20 == int32(3585))|base.B2i32(v20 == int32(4035)) != 0 {
		goto L31
	} else {
		goto L43
	}
L34:
	;
	v94 = v20 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v94) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(v20-int32(_a_F_multirangesel_1)) < base.Ui32(int32(6)) {
		goto L31
	} else {
		goto L40
	}
L37:
	;
	v98 = int32(1) << (uint(v94) % 32)
	if v98&int32(_a_F_multirangesel_2) != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v98&int32(_a_F_multirangesel_3) == int32(0) {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v126 = v90
	goto L30
L40:
	;
	if base.Ui32(v20-int32(_a_F_multirangesel_4)) < base.Ui32(int32(2)) {
		v126 = v90
		goto L30
	} else {
		goto L41
	}
L41:
	;
	if v20 != int32(_a_F_multirangesel_5) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	goto L31
L43:
	;
	goto L32
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v133 == int32(0) {
		v813 = v10
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+39)))
	if v139 != 0 {
		v149 = v20
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v136].(func(*base.Module, int32))(m, v133)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v813 = v10
	goto L1
L49:
	;
	if v149 <= int32(4034) {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v140 = F_get_commutator(m, v20)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v140 != 0 {
		v149 = v140
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v142 = float64(0.01)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v143 == int32(0) {
		v813 = v142
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v146].(func(*base.Module, int32))(m, v143)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v813 = v142
	goto L1
L55:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v791 != 0 {
		goto L235
	} else {
		goto L236
	}
L56:
	;
	v735 = float64(0.005)
	if v149 <= int32(_a_F_multirangesel_0) {
		goto L225
	} else {
		goto L226
	}
L57:
	;
	if v247 == int32(0) {
		goto L56
	} else {
		goto L83
	}
L58:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v235 != v236 {
		goto L56
	} else {
		goto L80
	}
L59:
	;
	if v149 == int32(3585) {
		goto L56
	} else {
		goto L79
	}
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v211 = F_multirange_get_typcache(m, l0, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L75
	}
L61:
	;
	v153 = v149 - int32(2866)
	if base.Ui32(int32(10)) < base.Ui32(v153) {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	switch v149 - int32(_a_F_multirangesel_1) {
	case 0, 3:
		goto L56
	case 1, 4:
		goto L60
	case 2:
		goto L58
	default:
		goto L72
	}
L64:
	;
	v157 = int32(1) << (uint(v153) % 32)
	if v157&int32(705) != 0 {
		goto L56
	} else {
		goto L65
	}
L65:
	;
	if v157&int32(1042) != 0 {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	if v153 != int32(3) {
		goto L59
	} else {
		goto L67
	}
L67:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v165 = F_multirange_get_typcache(m, l0, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+296))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+200))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v168 != v171 {
		goto L56
	} else {
		goto L69
	}
L69:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v173)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+146)) = uint8(v173)
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+144)) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v175
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+110)) = uint8(v178)
	v183 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+108)) = uint16(v183)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v175
	v192 = F_range_serialize(m, v169, v18+int32(140), v18+int32(104), v178, v178)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v165)+296))
	v200 = F_make_multirange(m, v195, v196, int32(1), v18+int32(32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v244 = v165
	v247 = v200
	goto L57
L72:
	;
	switch v149 - int32(_a_F_multirangesel_4) {
	case 0:
		goto L56
	case 1:
		goto L60
	default:
		goto L73
	}
L73:
	;
	if v149 != int32(4035) {
		goto L58
	} else {
		goto L74
	}
L74:
	;
	goto L60
L75:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+296))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	if v214 != v216 {
		goto L56
	} else {
		goto L76
	}
L76:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)+20))
	v219 = F_pg_detoast_datum(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v219
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v211)+296))
	v227 = F_make_multirange(m, v222, v223, int32(1), v18+int32(32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v244 = v211
	v247 = v227
	goto L57
L79:
	;
	goto L58
L80:
	;
	v238 = F_multirange_get_typcache(m, l0, v235)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+20))
	v242 = F_pg_detoast_datum(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v244 = v238
	v247 = v242
	goto L57
L83:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v250 != 0 {
		goto L93
	} else {
		goto L94
	}
L84:
	;
	v718 = float64(0)
	v721 = base.F64_mul(base.F64_sub(float64(1), v275), v712)
	if base.F64_lt(v721, v718) != 0 {
		v790 = v718
		goto L55
	} else {
		goto L218
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L2
	} else {
		goto L215
	}
L86:
	;
	if v149 == int32(_a_F_multirangesel_5) {
		v712 = v10
		goto L84
	} else {
		goto L213
	}
L87:
	;
	v673 = base.F64_sub(float64(1), v274)
	if base.B2i32(v149 != int32(_a_F_multirangesel_6))&base.B2i32(v149 != int32(2874)) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L88:
	;
	v617 = float64(0.005)
	if v149 <= int32(_a_F_multirangesel_0) {
		goto L200
	} else {
		goto L201
	}
L89:
	;
	F_free_attstatsslot(m, v18+int32(104))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L2
	} else {
		goto L192
	}
L90:
	;
	v583 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v326, v320, int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L191
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L2
	} else {
		goto L188
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L2
	} else {
		goto L185
	}
L93:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+22)))
	v254 = *(*float32)(unsafe.Add(mBase, uint32(v251+v252)+8))
	v256 = v18 + int32(140)
	v260 = F_get_attstatsslot(m, v256, v250, int32(6), int32(0), int32(2))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L96
	}
L94:
	;
	v274 = v10
	v275 = float64(0)
	goto L95
L95:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	if v276 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L96:
	;
	if v260 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v18)+164))
	if v262 != int32(1) {
		goto L92
	} else {
		goto L100
	}
L98:
	;
	v270 = v10
	goto L99
L99:
	;
	v274 = v270
	v275 = base.F64_promote_f32(v254)
	goto L95
L100:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
	v266 = *(*float32)(unsafe.Add(mBase, uint32(v265)))
	F_free_attstatsslot(m, v256)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v270 = base.F64_promote_f32(v266)
	goto L99
L102:
	;
	if v149 <= int32(_a_F_multirangesel_0) {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	goto L104
L104:
	;
	v296 = v18 + int32(44)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v244)+296))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+216))
	v299 = F_statistic_proc_security_check(m, v296, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L2
	} else {
		goto L114
	}
L105:
	;
	if v149 == int32(4035) {
		v712 = v10
		goto L84
	} else {
		goto L113
	}
L106:
	;
	v712 = base.F64_sub(float64(1), v274)
	goto L84
L107:
	;
	v712 = float64(1)
	goto L84
L108:
	;
	v712 = v274
	goto L84
L109:
	;
	switch v149 - int32(2862) {
	case 0, 5, 6, 14, 15:
		v712 = v10
		goto L84
	case 1, 12:
		goto L108
	case 2, 8, 9:
		goto L107
	case 3:
		goto L106
	case 4, 7, 10, 11, 13:
		goto L85
	default:
		goto L105
	}
L110:
	;
	goto L111
L111:
	;
	v284 = v149 - int32(_a_F_multirangesel_7)
	if base.B2i32(base.Ui32(int32(4)) < base.Ui32(v284))|base.B2i32(v284 == int32(2)) != 0 {
		goto L86
	} else {
		goto L112
	}
L112:
	;
	v712 = v10
	goto L84
L113:
	;
	goto L85
L114:
	;
	if v299 == int32(0) {
		goto L88
	} else {
		goto L115
	}
L115:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v297)+272))
	if v303 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v304 = F_statistic_proc_security_check(m, v296, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v308 == int32(0) {
		goto L88
	} else {
		goto L121
	}
L119:
	;
	if v304 == int32(0) {
		goto L88
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v316 = F_get_attstatsslot(m, v18+int32(140), v308, int32(7), int32(0), int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	if v316 == int32(0) {
		goto L88
	} else {
		goto L123
	}
L123:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v18)+156))
	if v320 < int32(2) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_free_attstatsslot(m, v18+int32(140))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L184
	}
L125:
	;
	v325 = v320 << (uint(int32(3)) % 32)
	v326 = F_palloc(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	v328 = F_palloc(m, v325)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L127
	}
L127:
	;
	v330 = int32(0)
	goto L128
L128:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v18)+152))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345+v330<<(uint(int32(2))%32))))
	v350 = F_pg_detoast_datum(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L130
	}
L129:
	;
	v367 = v149 - int32(2870)
	if base.B2i32(base.Ui32(int32(4)) < base.Ui32(v367))|base.B2i32(v367 == int32(2)) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v353 = v330 << (uint(int32(3)) % 32)
	F_range_deserialize(m, v297, v350, v326+v353, v328+v353, v18+int32(79))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+79)))
	if v360 == int32(1) {
		goto L91
	} else {
		goto L132
	}
L132:
	;
	v364 = v330 + int32(1)
	if v364 != v320 {
		v330 = v364
		goto L128
	} else {
		goto L133
	}
L133:
	;
	goto L129
L134:
	;
	v407 = v18 + int32(80)
	F_multirange_get_bounds(m, v297, v247, int32(0), v18+int32(96), v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L2
	} else {
		goto L142
	}
L135:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v375 == int32(0) {
		goto L124
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = int32(0)
	v393 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v393
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v393
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v393
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v393
	goto L134
L138:
	;
	v383 = F_get_attstatsslot(m, v18+int32(104), v375, int32(6), int32(0), int32(1))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	if v383 == int32(0) {
		goto L124
	} else {
		goto L140
	}
L140:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v18)+120))
	if int32(2) <= v388 {
		goto L134
	} else {
		goto L141
	}
L141:
	;
	v587 = float64(-1)
	goto L89
L142:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	F_multirange_get_bounds(m, v297, v247, v410-int32(1), v407, v18+int32(88))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	if v149 <= int32(_a_F_multirangesel_0) {
		goto L157
	} else {
		goto L158
	}
L144:
	;
	v532 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v326, v320, int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L2
	} else {
		goto L183
	}
L145:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L2
	} else {
		goto L180
	}
L146:
	;
	if v149 == int32(4035) {
		goto L90
	} else {
		goto L179
	}
L147:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+92)))
	if v494 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L148:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v18)+120))
	v492 = F_calc_hist_selectivity_contains(m, v297, v18+int32(96), v18+int32(88), v326, v320, v490, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L2
	} else {
		goto L173
	}
L149:
	;
	v475 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v328, v320, int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L171
	}
L150:
	;
	v469 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(88), v328, v320, int32(1))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L170
	}
L151:
	;
	v463 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(88), v326, v320, int32(1))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L2
	} else {
		goto L169
	}
L152:
	;
	v457 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v328, v320, int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L2
	} else {
		goto L168
	}
L153:
	;
	v451 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v326, v320, int32(1))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L2
	} else {
		goto L167
	}
L154:
	;
	v444 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v326, v320, int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L2
	} else {
		goto L166
	}
L155:
	;
	v438 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v326, v320, int32(1))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L2
	} else {
		goto L165
	}
L156:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+100)))
	if v427 != int32(1) {
		goto L147
	} else {
		goto L163
	}
L157:
	;
	switch v149 - int32(2862) {
	case 0:
		goto L144
	case 1:
		goto L155
	case 2:
		goto L153
	case 3:
		goto L154
	case 4, 10, 11, 13:
		goto L145
	case 5, 6, 7:
		goto L149
	case 8, 9:
		goto L148
	case 12:
		goto L156
	case 14, 15:
		goto L150
	default:
		goto L146
	}
L158:
	;
	goto L159
L159:
	;
	switch v149 - int32(_a_F_multirangesel_7) {
	case 0, 1:
		goto L152
	case 2:
		goto L145
	case 3, 4:
		goto L151
	default:
		goto L160
	}
L160:
	;
	if v149 == int32(_a_F_multirangesel_5) {
		goto L90
	} else {
		goto L161
	}
L161:
	;
	if v149 != int32(_a_F_multirangesel_6) {
		goto L145
	} else {
		goto L162
	}
L162:
	;
	goto L156
L163:
	;
	v433 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(88), v328, v320, int32(1))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	v587 = v433
	goto L89
L165:
	;
	v587 = v438
	goto L89
L166:
	;
	v587 = base.F64_sub(float64(1), v444)
	goto L89
L167:
	;
	v587 = base.F64_sub(float64(1), v451)
	goto L89
L168:
	;
	v587 = v457
	goto L89
L169:
	;
	v587 = base.F64_sub(float64(1), v463)
	goto L89
L170:
	;
	v587 = v469
	goto L89
L171:
	;
	v481 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(88), v326, v320, int32(1))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L2
	} else {
		goto L172
	}
L172:
	;
	v587 = base.F64_sub(float64(1), base.F64_add(v475, base.F64_sub(float64(1), v481)))
	goto L89
L173:
	;
	v587 = v492
	goto L89
L174:
	;
	v501 = F_calc_hist_selectivity_scalar(m, v297, v18+int32(96), v326, v320, int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L2
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v18)+120))
	v510 = F_calc_hist_selectivity_contained(m, v297, v18+int32(96), v18+int32(88), v326, v320, v508, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L2
	} else {
		goto L178
	}
L177:
	;
	v587 = base.F64_sub(float64(1), v501)
	goto L89
L178:
	;
	v587 = v510
	goto L89
L179:
	;
	goto L145
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v149
	F_errmsg_internal(m, int32(_a_F_multirangesel_8), v18+int32(16))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_multirangesel_9), int32(690), int32(_a_F_multirangesel_10))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L2
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v587 = v532
	goto L89
L184:
	;
	goto L88
L185:
	;
	F_errmsg_internal(m, int32(_a_F_multirangesel_11), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_multirangesel_9), int32(318), int32(_a_F_multirangesel_12))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errmsg_internal(m, int32(_a_F_multirangesel_13), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L2
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_multirangesel_9), int32(509), int32(_a_F_multirangesel_10))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L2
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	v587 = base.F64_sub(float64(1), v583)
	goto L89
L192:
	;
	F_free_attstatsslot(m, v18+int32(140))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L2
	} else {
		goto L193
	}
L193:
	;
	if base.F64_lt(v587, float64(0)) == int32(0) {
		v666 = v587
		goto L87
	} else {
		goto L194
	}
L194:
	;
	goto L88
L195:
	;
	v666 = v656
	goto L87
L196:
	;
	v656 = v653
	goto L195
L197:
	;
	v653 = float64(0.3333333333333333)
	goto L196
L198:
	;
	v656 = float64(0.01)
	goto L195
L199:
	;
	if base.B2i32(v149 == int32(3585))|base.B2i32(v149 == int32(4035)) != 0 {
		goto L197
	} else {
		goto L209
	}
L200:
	;
	v621 = v149 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v621) {
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v149-int32(_a_F_multirangesel_1)) < base.Ui32(int32(6)) {
		goto L197
	} else {
		goto L206
	}
L203:
	;
	v625 = int32(1) << (uint(v621) % 32)
	if v625&int32(_a_F_multirangesel_2) != 0 {
		goto L197
	} else {
		goto L204
	}
L204:
	;
	if v625&int32(_a_F_multirangesel_3) == int32(0) {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v653 = v617
	goto L196
L206:
	;
	if base.Ui32(v149-int32(_a_F_multirangesel_4)) < base.Ui32(int32(2)) {
		v653 = v617
		goto L196
	} else {
		goto L207
	}
L207:
	;
	if v149 != int32(_a_F_multirangesel_5) {
		goto L198
	} else {
		goto L208
	}
L208:
	;
	goto L197
L209:
	;
	goto L198
L210:
	;
	v712 = base.F64_add(base.F64_mul(v673, v666), v274)
	goto L84
L211:
	;
	goto L212
L212:
	;
	v712 = base.F64_mul(v673, v666)
	goto L84
L213:
	;
	if v149 == int32(_a_F_multirangesel_6) {
		v712 = v274
		goto L84
	} else {
		goto L214
	}
L214:
	;
	goto L85
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v149
	F_errmsg_internal(m, int32(_a_F_multirangesel_14), v18)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L2
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_multirangesel_9), int32(402), int32(_a_F_multirangesel_12))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L2
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
	if base.F64_gt(v721, float64(1)) == int32(0) {
		v790 = v721
		goto L55
	} else {
		goto L219
	}
L219:
	;
	v790 = float64(1)
	goto L55
L220:
	;
	v790 = v774
	goto L55
L221:
	;
	v774 = v771
	goto L220
L222:
	;
	v771 = float64(0.3333333333333333)
	goto L221
L223:
	;
	v774 = float64(0.01)
	goto L220
L224:
	;
	if base.B2i32(v149 == int32(3585))|base.B2i32(v149 == int32(4035)) != 0 {
		goto L222
	} else {
		goto L234
	}
L225:
	;
	v739 = v149 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v739) {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if base.Ui32(v149-int32(_a_F_multirangesel_1)) < base.Ui32(int32(6)) {
		goto L222
	} else {
		goto L231
	}
L228:
	;
	v743 = int32(1) << (uint(v739) % 32)
	if v743&int32(_a_F_multirangesel_2) != 0 {
		goto L222
	} else {
		goto L229
	}
L229:
	;
	if v743&int32(_a_F_multirangesel_3) == int32(0) {
		goto L224
	} else {
		goto L230
	}
L230:
	;
	v771 = v735
	goto L221
L231:
	;
	if base.Ui32(v149-int32(_a_F_multirangesel_4)) < base.Ui32(int32(2)) {
		v771 = v735
		goto L221
	} else {
		goto L232
	}
L232:
	;
	if v149 != int32(_a_F_multirangesel_5) {
		goto L223
	} else {
		goto L233
	}
L233:
	;
	goto L222
L234:
	;
	goto L223
L235:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v792].(func(*base.Module, int32))(m, v791)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L2
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v795 = float64(0)
	if base.F64_lt(v790, v795) != 0 {
		v813 = v795
		goto L1
	} else {
		goto L239
	}
L238:
	;
	goto L237
L239:
	;
	if base.F64_gt(v790, float64(1)) == int32(0) {
		v813 = v790
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v813 = float64(1)
	goto L1
L241:
	;
	m.G0 = v18 + int32(176)
	return v818
}
func F_multixactoffsetssyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(_a_F_multixactoffsetssyncfiletag_0), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
