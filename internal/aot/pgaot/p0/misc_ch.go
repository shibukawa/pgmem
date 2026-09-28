package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckAffix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v516 int32
	_ = v516
	if l3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v516
L2:
	;
	v516 = int32(0)
	goto L1
L3:
	;
	v48 = int32(0)
	v52 = int32(base.Ui32(v47)>>(uint(int32(10))%32)) & int32(_a_F_CheckAffix_0)
	if base.Ui32(l1) < base.Ui32(v52) {
		v516 = v48
		goto L1
	} else {
		goto L19
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v13&int32(2) == int32(0) {
		v47 = v13
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
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v21&int32(64) != 0 {
		v516 = v20
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
	if v21&int32(5) != int32(1) {
		v47 = v21
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v516 = v20
	goto L1
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v30&int32(72) == int32(8) {
		v47 = v30
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(l3) < base.Ui32(int32(8)) {
		v47 = v35
		goto L3
	} else {
		goto L17
	}
L16:
	;
	goto L2
L17:
	;
	v38 = int32(0)
	if base.B2i32(v35&int32(17) == v38)|v35&int32(64) != 0 {
		v516 = v38
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v47 = v35
	goto L3
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v55 = F_strlen(m, v54)
	mBase = m.M
	v56 = l1 - v52
	if base.Ui32(int32(511)) < base.Ui32(v55+v56) {
		v516 = v48
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v47&int32(1) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v227&int32(256) != 0 {
		goto L78
	} else {
		goto L79
	}
L22:
	;
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if l5 != 0 {
		goto L50
	} else {
		goto L51
	}
L25:
	;
	base.MemoryCopy(m, l4, l0, v56)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v63 = v56 + l4
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if (v64^v63)&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if l5 == int32(0) {
		goto L21
	} else {
		goto L49
	}
L29:
	;
	goto L28
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v118)
	if v118&int32(255) == int32(0) {
		goto L29
	} else {
		goto L45
	}
L31:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v117 = v64
	v118 = v70
	v119 = v63
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v64&int32(3) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v74 = v64
	v76 = v63
	goto L37
L35:
	;
	v88 = v64
	v90 = v63
	goto L36
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v95 = int32(-2139062144)
	if (int32(16843008)-v92|v92)&v95 != v95 {
		v117 = v88
		v118 = v92
		v119 = v90
		goto L30
	} else {
		goto L41
	}
L37:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v77)
	if v77 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L38:
	;
	v88 = v84
	v90 = v82
	goto L36
L39:
	;
	v81 = int32(1)
	v82 = v76 + v81
	v84 = v74 + v81
	if v84&int32(3) != 0 {
		v74 = v84
		v76 = v82
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v100 = v88
	v101 = v92
	v102 = v90
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v101
	v104 = int32(4)
	v105 = v102 + v104
	v107 = v100 + v104
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v112 = int32(-2139062144)
	if (int32(16843008)-v109|v109)&v112 == v112 {
		v100 = v107
		v101 = v109
		v102 = v105
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v117 = v107
	v118 = v109
	v119 = v105
	goto L30
L44:
	;
	goto L43
L45:
	;
	v126 = v117
	v128 = v119
	goto L46
L46:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v129)
	v131 = int32(1)
	if v129 != 0 {
		v126 = v126 + v131
		v128 = v128 + v131
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L29
L48:
	;
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v56
	goto L21
L50:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if base.Ui32(v142+v55) <= base.Ui32(v52) {
		v516 = v48
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v55 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	base.MemoryCopy(m, l4, v54, v55)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v146 = l4 + v55
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v152 = l0 + int32(base.Ui32(v147)>>(uint(int32(10))%32))&int32(_a_F_CheckAffix_0)
	if (v152^v146)&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L21
L58:
	;
	goto L57
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v206)
	if v206&int32(255) == int32(0) {
		goto L58
	} else {
		goto L74
	}
L60:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v205 = v152
	v206 = v158
	v207 = v146
	goto L59
L61:
	;
	goto L62
L62:
	;
	if v152&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v162 = v152
	v164 = v146
	goto L66
L64:
	;
	v176 = v152
	v178 = v146
	goto L65
L65:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v183 = int32(-2139062144)
	if (int32(16843008)-v180|v180)&v183 != v183 {
		v205 = v176
		v206 = v180
		v207 = v178
		goto L59
	} else {
		goto L70
	}
L66:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v165)
	if v165 == int32(0) {
		goto L58
	} else {
		goto L68
	}
L67:
	;
	v176 = v172
	v178 = v170
	goto L65
L68:
	;
	v169 = int32(1)
	v170 = v164 + v169
	v172 = v162 + v169
	if v172&int32(3) != 0 {
		v162 = v172
		v164 = v170
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v188 = v176
	v189 = v180
	v190 = v178
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v189
	v192 = int32(4)
	v193 = v190 + v192
	v195 = v188 + v192
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v200 = int32(-2139062144)
	if (int32(16843008)-v197|v197)&v200 == v200 {
		v188 = v195
		v189 = v197
		v190 = v193
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v205 = v195
	v206 = v197
	v207 = v193
	goto L59
L73:
	;
	goto L72
L74:
	;
	v214 = v205
	v216 = v207
	goto L75
L75:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)) = uint8(v217)
	v219 = int32(1)
	if v217 != 0 {
		v214 = v214 + v219
		v216 = v216 + v219
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L58
L77:
	;
	goto L76
L78:
	;
	return l4
L79:
	;
	goto L80
L80:
	;
	if v227&int32(512) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v233 = int32(0)
	v235 = m.G0
	v236 = int32(16)
	v237 = v235 - v236
	m.G0 = v237
	v240 = l2 + v236
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v242 != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	goto L83
L83:
	;
	v485 = F_strlen(m, l4)
	mBase = m.M
	v488 = F_palloc_mul(m, int32(4), v485+int32(1))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L90
	} else {
		goto L143
	}
L84:
	;
	if v460 == int32(0) {
		goto L2
	} else {
		goto L142
	}
L85:
	;
	v247 = l4
	v248 = v233
	goto L88
L86:
	;
	v266 = v233
	goto L87
L87:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v275 = int32(base.Ui32(v271)>>(uint(int32(1))%32)) & int32(_a_F_CheckAffix_1)
	if v266 < v275 {
		v460 = v233
		goto L94
	} else {
		goto L95
	}
L88:
	;
	v254 = v248 + int32(1)
	v255 = F_pg_mblen_cstr(m, v247)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v266 = v254
	goto L87
L90:
	;
	return int32(0)
L91:
	;
	v259 = v255 + v247
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v260 != 0 {
		v247 = v259
		v248 = v254
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L90
	} else {
		goto L139
	}
L94:
	;
	m.G0 = v237 + int32(16)
	goto L84
L95:
	;
	if v271&int32(1) == int32(0) {
		v302 = l4
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v241 != 0 {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v281 = v266 - v275
	if v281 <= int32(0) {
		v302 = l4
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v285 = l4
	v288 = v281
	goto L99
L99:
	;
	v294 = F_pg_mblen_cstr(m, v285)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L90
	} else {
		goto L101
	}
L100:
	;
	v302 = v296
	goto L96
L101:
	;
	v296 = v294 + v285
	v297 = int32(1)
	if base.Ui32(v297) < base.Ui32(v288) {
		v285 = v296
		v288 = v288 - v297
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v312 = v302
	v313 = v241
	goto L106
L104:
	;
	goto L105
L105:
	;
	v460 = int32(1)
	goto L94
L106:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	switch v321&int32(3) - int32(1) {
	case 0:
		goto L110
	case 1:
		goto L109
	default:
		goto L93
	}
L107:
	;
	goto L105
L108:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v439 = F_pg_mblen_cstr(m, v312)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L90
	} else {
		goto L137
	}
L109:
	;
	v377 = F_pg_mblen_cstr(m, v312)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L90
	} else {
		goto L124
	}
L110:
	;
	v326 = F_pg_mblen_cstr(m, v312)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L90
	} else {
		goto L111
	}
L111:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+8)))
	if v328 == int32(0) {
		v460 = v233
		goto L94
	} else {
		goto L112
	}
L112:
	;
	v338 = v313 + int32(8)
	goto L113
L113:
	;
	v343 = F_pg_mblen_cstr(m, v338)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L90
	} else {
		goto L115
	}
L114:
	;
	v460 = v233
	goto L94
L115:
	;
	if v326 == v343 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v350 = v326
	goto L119
L117:
	;
	goto L118
L118:
	;
	v375 = v338 + v343
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	if v376 != 0 {
		v338 = v375
		goto L113
	} else {
		goto L123
	}
L119:
	;
	if v350 == int32(0) {
		goto L108
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	v359 = v350 - int32(1)
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338+v359))))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312+v359))))
	if v361 == v363 {
		v350 = v359
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	goto L114
L124:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+8)))
	if v379 == int32(0) {
		goto L108
	} else {
		goto L125
	}
L125:
	;
	v389 = v313 + int32(8)
	goto L126
L126:
	;
	v394 = F_pg_mblen_cstr(m, v389)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L90
	} else {
		goto L128
	}
L127:
	;
	goto L108
L128:
	;
	if v377 == v394 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v401 = v377
	goto L132
L130:
	;
	goto L131
L131:
	;
	v426 = v389 + v394
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v427 != 0 {
		v389 = v426
		goto L126
	} else {
		goto L136
	}
L132:
	;
	if v401 == int32(0) {
		v460 = v233
		goto L94
	} else {
		goto L134
	}
L133:
	;
	goto L131
L134:
	;
	v410 = v401 - int32(1)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+v410))))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312+v410))))
	if v412 == v414 {
		v401 = v410
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	goto L127
L137:
	;
	if v438 != 0 {
		v312 = v439 + v312
		v313 = v438
		goto L106
	} else {
		goto L138
	}
L138:
	;
	goto L107
L139:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v470 & int32(3)
	F_errmsg_internal(m, int32(_a_F_CheckAffix_2), v237)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L90
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_CheckAffix_3), int32(245), int32(_a_F_CheckAffix_4))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L90
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	v516 = l4
	goto L1
L143:
	;
	v490 = F_pg_mb2wchar_with_len(m, l4, v488, v485)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L90
	} else {
		goto L144
	}
L144:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v493 = int32(0)
	v496 = F_pg_regexec(m, v492, v488, v490, v493, v493, v493)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L90
	} else {
		goto L145
	}
L145:
	;
	F_pfree(m, v488)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L90
	} else {
		goto L146
	}
L146:
	;
	if v496 == int32(0) {
		v516 = l4
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L2
}
func F_CheckDim_3(m *base.Module, l0 int32) {
	var v10 int32
	_ = v10
	Fn14208(m, l0, int32(105), int32(_a_F_CheckDim_3_0), int32(_a_F_CheckDim_3_1), int32(_a_F_CheckDim_3_2), int32(100), int32(_a_F_CheckDim_3_3), int32(_a_F_CheckDim_3_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_charge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(base.Ui32(v3) <= base.Ui32(v2)))
}
func F_charlt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(base.Ui32(v2) < base.Ui32(v3)))
}
func F_charout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	v6 = F_palloc(m, int32(5))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v4 < int32(0) {
			v12 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v12)
			v14 = int32(7)
			v16 = int32(48)
			v17 = v4&v14 | v16
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v17)
			v24 = int32(base.Ui32(v4)>>(uint(int32(3))%32))&v14 | v16
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v24)
			v33 = int32(92)
			v34 = int32(base.Ui32(v4&int32(192))>>(uint(int32(6))%32)) | v16
		} else {
			v33 = v4
			v34 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v34)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v33)
		return base.I64_extend_i32_u(v6)
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
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
		v83 = v4
		m.G0 = v11 + int32(16)
		return v83
	} else {
		v22 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				v83 = v4
				m.G0 = v11 + int32(16)
				return v83
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
				v30 = v28 + v29
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+128)))
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+127)))
				F_ReleaseCatCache(m, v22)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v32 != int32(1) {
						v83 = v4
						m.G0 = v11 + int32(16)
						return v83
					} else {
						v37 = int32(1)
						v38 = F_has_bypassrls_privilege(m, v17)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 != 0 {
								v83 = v37
								m.G0 = v11 + int32(16)
								return v83
							} else {
								v41 = F_object_ownercheck(m, int32(1259), l0, v17)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									if v41 != 0 {
										if v31&int32(1) == int32(0) {
											v83 = v37
											m.G0 = v11 + int32(16)
											return v83
										} else {
											v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_enable_rls[1])))
											if int32(base.Ui32(v48&int32(4))>>(uint(int32(2))%32)) != 0 {
												v83 = v37
												m.G0 = v11 + int32(16)
												return v83
											} else {
												v53 = int32(2)
												if l2 != 0 {
													v83 = v53
													m.G0 = v11 + int32(16)
													return v83
												} else {
													v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_enable_rls[2])))
													if v55&int32(1) != 0 {
														v83 = v53
														m.G0 = v11 + int32(16)
														return v83
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16797828))
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v65 = F_get_rel_name(m, l0)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v65
																	F_errmsg(m, int32(_a_F_check_enable_rls_1), v11)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return int32(0)
																	} else {
																		if v41 != 0 {
																			F_errhint(m, int32(_a_F_check_enable_rls_2), int32(0))
																			mBase = m.M
																			v74 = m.ExcPending
																			if v74 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_check_enable_rls_3), int32(129), int32(_a_F_check_enable_rls_4))
																				mBase = m.M
																				v79 = m.ExcPending
																				if v79 != 0 {
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
																			v79 = m.ExcPending
																			if v79 != 0 {
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
										v53 = int32(2)
										if l2 != 0 {
											v83 = v53
											m.G0 = v11 + int32(16)
											return v83
										} else {
											v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_enable_rls[2])))
											if v55&int32(1) != 0 {
												v83 = v53
												m.G0 = v11 + int32(16)
												return v83
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16797828))
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return int32(0)
													} else {
														v65 = F_get_rel_name(m, l0)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v65
															F_errmsg(m, int32(_a_F_check_enable_rls_1), v11)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																if v41 != 0 {
																	F_errhint(m, int32(_a_F_check_enable_rls_2), int32(0))
																	mBase = m.M
																	v74 = m.ExcPending
																	if v74 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_check_enable_rls_3), int32(129), int32(_a_F_check_enable_rls_4))
																		mBase = m.M
																		v79 = m.ExcPending
																		if v79 != 0 {
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
																	v79 = m.ExcPending
																	if v79 != 0 {
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
func F_check_functional_grouping(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v357 int32
	_ = v357
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = m.G0
	v24 = v22 - int32(80)
	m.G0 = v24
	v27 = v20 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v6
	v32 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v357
L2:
	;
	if v165 == int32(0) {
		v357 = v6
		goto L1
	} else {
		goto L43
	}
L3:
	;
	return int32(0)
L4:
	;
	v37 = v24 + int32(16)
	F_ScanKeyInit(m, v37, int32(9), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v45 = int32(1)
	v48 = F_systable_beginscan(m, v32, int32(2665), v45, int32(0), v45, v37)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L40
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L37
	}
L8:
	;
	F_systable_endscan(m, v48)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L35
	}
L9:
	;
	v50 = F_systable_getnext(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v50 == int32(0) {
		v165 = v6
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v54 = v50
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v73 = v71 + v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+72)))
	if v74 == int32(112) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v165 = v6
	goto L8
L14:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+73)))
	if v77&int32(1) != 0 {
		v165 = v6
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v158 = F_systable_getnext(m, v48)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L33
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v83 = F_heap_getattr_3(m, v54, v80, v24+int32(15))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+15)))
	if v85 == int32(1) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v89 = F_pg_detoast_datum(m, base.I32_wrap_i64(v83))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v91 != int32(1) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v94 < int32(0) {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v97 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v98 != int32(21) {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v101 = int32(0)
	if v94 == v101 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+22)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153+v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v156
	v165 = v141
	goto L8
L26:
	;
	v141 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v113 = int32(0)
	v115 = v101
	goto L29
L29:
	;
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89+int32(24)+v115<<(uint(int32(1))%32)))))
	v131 = F_bms_add_member(m, v113, v128+int32(7))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	v141 = v131
	goto L25
L31:
	;
	v134 = v115 + int32(1)
	if v134 != v94 {
		v113 = v131
		v115 = v134
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if v158 != 0 {
		v54 = v158
		goto L12
	} else {
		goto L34
	}
L34:
	;
	goto L13
L35:
	;
	F_relation_close(m, v32, int32(1))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	m.G0 = v24 + int32(80)
	goto L2
L37:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+22)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v192
	F_errmsg_internal(m, int32(_a_F_check_functional_grouping_0), v24)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_check_functional_grouping_1), int32(1504), int32(_a_F_check_functional_grouping_2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errmsg_internal(m, int32(_a_F_check_functional_grouping_3), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_check_functional_grouping_1), int32(1511), int32(_a_F_check_functional_grouping_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	if l3 == int32(0) {
		v275 = v6
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v280 = int32(0)
	if v165 == v280 {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v220 <= int32(0) {
		v275 = v6
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v224 = int32(0)
	v236 = v6
	goto L47
L47:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241+v224<<(uint(int32(2))%32))))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v246 != int32(6) {
		v258 = v236
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v275 = v258
	goto L44
L49:
	;
	v260 = v224 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v260 < v261 {
		v224 = v260
		v236 = v258
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v249 != l1 {
		v258 = v236
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)+28))
	if v251 != l2 {
		v258 = v236
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+8)))
	v256 = F_bms_add_member(m, v236, v253+int32(7))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v258 = v256
	goto L49
L54:
	;
	goto L48
L55:
	;
	if v333 == int32(0) {
		v357 = v6
		goto L1
	} else {
		goto L69
	}
L56:
	;
	v333 = int32(1)
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v275 == int32(0) {
		v326 = v280
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v333 = v326
	goto L55
L60:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v290 < v289 {
		v326 = v280
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v292 = int32(1)
	if v289 <= v292 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v295 = v292
	goto L64
L63:
	;
	v295 = v289
	goto L64
L64:
	;
	v296 = int32(8)
	v301 = int32(0)
	goto L65
L65:
	;
	v308 = v301 << (uint(int32(2)) % 32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v165+v296+v308)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v275+v296+v308)))
	v315 = v310 & (v312 ^ int32(-1))
	v317 = base.B2i32(v315 == int32(0))
	if v315 != 0 {
		v326 = v317
		goto L59
	} else {
		goto L67
	}
L66:
	;
	v326 = v317
	goto L59
L67:
	;
	v319 = v301 + int32(1)
	if v319 != v295 {
		v301 = v319
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v338 = F_lappend_oid(m, v336, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v338
	v357 = int32(1)
	goto L1
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
	F_errfinish(m, int32(_a_F_check_labels_2), int32(3897), int32(_a_F_check_labels_3))
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
	F_errfinish(m, int32(_a_F_check_labels_2), int32(3904), int32(_a_F_check_labels_3))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
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
	var v65 int32
	_ = v65
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = int32(25)
	v12 = v5 + int32(-16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_check_publications_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_GetPublicationsStr(m, l1, v12, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v12, int32(41))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v31 = m.T0[v30].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v24, int32(1), v5+int32(-20))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	F_pfree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v36 == int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = F_list_copy(m, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v43 = F_MakeSingleTupleTableSlot(m, v41, int32(_a_F_check_publications_1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v48 = F_tuplestore_gettupleslot(m, v45, int32(1), int32(0), v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = v39
	goto L17
L15:
	;
	v79 = v39
	goto L16
L16:
	;
	F_ExecDropSingleTupleTableSlot(m, v43)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	if v54 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v79 = v68
	goto L16
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	m.T0[v59].(func(*base.Module, int32, int32))(m, v43, int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = F_text_to_cstring(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v66 = F_makeString(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v68 = F_list_delete(m, v50, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	m.T0[v71].(func(*base.Module, int32))(m, v43)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v77 = F_tuplestore_gettupleslot(m, v74, int32(1), int32(0), v43)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v77 != 0 {
		v50 = v68
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v85 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_pfree(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v88 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	F_tuplestore_end(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v91 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	F_FreeTupleDesc(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v31)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v79 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	m.G0 = v7 - int32(-64)
	return
L44:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v98 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v102 = v5 + int32(-36)
	F_initStringInfo(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_GetPublicationsStr(m, v79, v102, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v110 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v110 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v118
	F_errmsg_plural(m, int32(_a_F_check_publications_2), int32(_a_F_check_publications_3), v117, v7)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_check_publications_4), int32(609), int32(_a_F_check_publications_5))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L43
L53:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v138
	F_errmsg(m, int32(_a_F_check_publications_6), v5+int32(-48))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_check_publications_4), int32(573), int32(_a_F_check_publications_5))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v134 int32
	_ = v134
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v493 int32
	_ = v493
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(208)
	m.G0 = v18
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(208)
	return v493
L2:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_check_usermap[0]))
	if v69 == int32(0) {
		v450 = v4
		v452 = v4
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v20 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v23 == int32(0))|base.B2i32(v23 != v26) != 0 {
		v44 = v23
		v45 = v26
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	if v44-v45 == int32(0) {
		v493 = v4
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v29 = l1
	v30 = l2
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v34 == int32(0) {
		v44 = v34
		v45 = v33
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v44 = v34
	v45 = v33
	goto L8
L12:
	;
	v37 = int32(1)
	if v34 == v33 {
		v29 = v29 + v37
		v30 = v30 + v37
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v49 = int32(-1)
	v52 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v52 == int32(0) {
		v493 = v49
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
	F_errmsg(m, int32(_a_F_check_usermap_0), v18)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2813), int32(_a_F_check_usermap_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v493 = v49
	goto L1
L20:
	;
	if v450|v452 != 0 {
		goto L126
	} else {
		goto L127
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v72 <= int32(0) {
		v450 = v4
		v452 = v4
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v83 = v4
	goto L23
L23:
	;
	v90 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v83<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v100 == v90)|base.B2i32(v100 != v103) != 0 {
		v121 = v100
		v122 = v103
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v444 = int32(0)
	v450 = v444
	v452 = v444
	goto L20
L25:
	;
	if v428|v430 != 0 {
		v450 = v428
		v452 = v430
		goto L20
	} else {
		goto L124
	}
L26:
	;
	if v121-v122 != 0 {
		v428 = v90
		v430 = v90
		goto L25
	} else {
		goto L33
	}
L27:
	;
	goto L26
L28:
	;
	v106 = v97
	v107 = l0
	goto L29
L29:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v111
		v122 = v110
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v121 = v111
	v122 = v110
	goto L27
L31:
	;
	v114 = int32(1)
	if v111 == v110 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v125 = F_get_role_oid(m, l1, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	if v128 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v185
	v420 = F_list_make1_impl(m, int32(1), v18+int32(40))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L15
	} else {
		goto L122
	}
L36:
	;
	v129 = F_strlen(m, l2)
	mBase = m.M
	v134 = F_palloc(m, v129<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L15
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v378 == int32(0))|base.B2i32(v378 != v381) != 0 {
		v399 = v378
		v400 = v381
		goto L113
	} else {
		goto L114
	}
L39:
	;
	v136 = F_strlen(m, l2)
	mBase = m.M
	v137 = F_pg_mb2wchar_with_len(m, l2, v134, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	v144 = F_pg_regexec(m, v139, v134, v137, int32(0), int32(2), v18+int32(192))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	F_pfree(m, v134)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	if v144 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v144 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)))
	if v186 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L46:
	;
	v428 = int32(0)
	v430 = base.B2i32(v144 != int32(1))
	goto L25
L47:
	;
	v153 = v18 + int32(80)
	F_pg_regerror(m, v144, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	v158 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	if v158 == int32(0) {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v166 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v153
	F_errmsg(m, int32(_a_F_check_usermap_3), v18-int32(-64))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2668), int32(_a_F_check_usermap_4))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v190 == int32(43) {
		goto L35
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	if v193 != 0 {
		goto L35
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v196 = F_strstr(m, v194, int32(_a_F_check_usermap_5))
	mBase = m.M
	if v196 == int32(0) {
		goto L35
	} else {
		goto L59
	}
L59:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v18)+200))
	if v199 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v202 = int32(1)
	v203 = int32(0)
	v206 = F_errstart(m, int32(15), v203)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L15
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
	v233 = v232 - v199
	v240 = v194
	v242 = v196
	goto L68
L63:
	;
	if v206 == int32(0) {
		v428 = v203
		v430 = v202
		goto L25
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v214 + int32(1)
	F_errmsg(m, int32(_a_F_check_usermap_6), v18+int32(48))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2695), int32(_a_F_check_usermap_4))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	v428 = v203
	v430 = v202
	goto L25
L68:
	;
	v251 = F_strlen(m, v240)
	mBase = m.M
	v253 = F_palloc(m, v251+(v233-int32(1)))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L15
	} else {
		goto L70
	}
L69:
	;
	v343 = int32(0)
	v344 = F_strlen(m, v253)
	mBase = m.M
	v347 = F_palloc0(m, v344+int32(13))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L15
	} else {
		goto L103
	}
L70:
	;
	v255 = v242 - v240
	if v255 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v253, v240, v255)
	goto L73
L72:
	;
	goto L73
L73:
	;
	v257 = v253 + v255
	if v233 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	base.MemoryCopy(m, v257, l2+v199, v233)
	goto L76
L75:
	;
	goto L76
L76:
	;
	v259 = v233 + v257
	v261 = v242 + int32(2)
	if (v261^v259)&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v337 != v240 {
		goto L98
	} else {
		goto L99
	}
L78:
	;
	goto L77
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v316))) = uint8(v315)
	if v315&int32(255) == int32(0) {
		goto L78
	} else {
		goto L94
	}
L80:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v314 = v261
	v315 = v267
	v316 = v259
	goto L79
L81:
	;
	goto L82
L82:
	;
	if v261&int32(3) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v271 = v261
	v273 = v259
	goto L86
L84:
	;
	v285 = v261
	v287 = v259
	goto L85
L85:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v292 = int32(-2139062144)
	if (int32(16843008)-v289|v289)&v292 != v292 {
		v314 = v285
		v315 = v289
		v316 = v287
		goto L79
	} else {
		goto L90
	}
L86:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v274)
	if v274 == int32(0) {
		goto L78
	} else {
		goto L88
	}
L87:
	;
	v285 = v281
	v287 = v279
	goto L85
L88:
	;
	v278 = int32(1)
	v279 = v273 + v278
	v281 = v271 + v278
	if v281&int32(3) != 0 {
		v271 = v281
		v273 = v279
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v297 = v285
	v298 = v289
	v299 = v287
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v298
	v301 = int32(4)
	v302 = v299 + v301
	v304 = v297 + v301
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v309 = int32(-2139062144)
	if (int32(16843008)-v306|v306)&v309 == v309 {
		v297 = v304
		v298 = v306
		v299 = v302
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v314 = v304
	v315 = v306
	v316 = v302
	goto L79
L93:
	;
	goto L92
L94:
	;
	v323 = v314
	v325 = v316
	goto L95
L95:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v325)+1)) = uint8(v326)
	v328 = int32(1)
	if v326 != 0 {
		v323 = v323 + v328
		v325 = v325 + v328
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L78
L97:
	;
	goto L96
L98:
	;
	F_pfree(m, v240)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L15
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v342 = F_strstr(m, v259, int32(_a_F_check_usermap_5))
	mBase = m.M
	if v342 != 0 {
		v240 = v253
		v242 = v342
		goto L68
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	goto L69
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+8)) = int32(0)
	v351 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v347)+4)) = uint8(v351)
	v354 = v347 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v354
	v357 = v344 + v351
	if v357 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	base.MemoryCopy(m, v354, v253, v357)
	goto L106
L105:
	;
	goto L106
L106:
	;
	F_pfree(m, v253)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v347
	v366 = F_list_make1_impl(m, int32(1), v18+int32(44))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	v368 = F_check_role_2(m, l1, v125, v366)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v370 == int32(0) {
		v428 = v368
		v430 = v343
		goto L25
	} else {
		goto L110
	}
L110:
	;
	F_pg_regfree(m, v370)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L15
	} else {
		goto L111
	}
L111:
	;
	v428 = v368
	v430 = v343
	goto L25
L112:
	;
	if v399-v400 != 0 {
		v428 = v90
		v430 = v90
		goto L25
	} else {
		goto L119
	}
L113:
	;
	goto L112
L114:
	;
	v384 = v375
	v385 = l2
	goto L115
L115:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+1)))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)))
	if v389 == int32(0) {
		v399 = v389
		v400 = v388
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v399 = v389
	v400 = v388
	goto L113
L117:
	;
	v392 = int32(1)
	if v389 == v388 {
		v384 = v384 + v392
		v385 = v385 + v392
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v402
	v408 = F_list_make1_impl(m, int32(1), v18+int32(36))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L15
	} else {
		goto L120
	}
L120:
	;
	v410 = F_check_role_2(m, l1, v125, v408)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	v428 = v410
	v430 = v90
	goto L25
L122:
	;
	v422 = F_check_role_2(m, l1, v125, v420)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	v428 = v422
	v430 = int32(0)
	goto L25
L124:
	;
	v441 = v83 + int32(1)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v441 < v442 {
		v83 = v441
		goto L23
	} else {
		goto L125
	}
L125:
	;
	goto L24
L126:
	;
	v493 = int32(0) - (v450 ^ int32(1))
	goto L1
L127:
	;
	v464 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	if v464 == int32(0) {
		goto L126
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
	F_errmsg(m, int32(_a_F_check_usermap_7), v18+int32(16))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L15
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_check_usermap_1), int32(2833), int32(_a_F_check_usermap_2))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L15
	} else {
		goto L131
	}
L131:
	;
	goto L126
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
	var v314 int64
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v353 int64
	_ = v353
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v383 int64
	_ = v383
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v399 int64
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int64
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v436 int32
	_ = v436
	var v437 int64
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v443 int64
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v451 int64
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v512 int32
	_ = v512
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
	v512 = v6
	goto L6
L6:
	;
	return v512
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v51<<(uint(int32(2))%32))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+16)))
	if v59 != int32(109) {
		v480 = v46
		v484 = v50
		v487 = v53
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v512 = v487
	goto L6
L9:
	;
	v489 = v51 + int32(1)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v489 < v490 {
		v46 = v480
		v50 = v484
		v51 = v489
		v53 = v487
		goto L7
	} else {
		goto L110
	}
L10:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+8)))
	if v62 != l1 {
		v480 = v46
		v484 = v50
		v487 = v53
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
	v314 = int64(0)
	if v301 == int32(0) {
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
	v360 = int64(0)
	if v302 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L60:
	;
	v358 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v319 = v301 + int32(8)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v320 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v358 = base.I32_popcnt(v323)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v326 = v320 << (uint(int32(2)) % 32)
	if v326 <= int32(7) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v358 = base.I32_wrap_i64(v353)
	goto L59
L67:
	;
	if v326 == int32(0) {
		v353 = v314
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v350 = F_pg_popcount_optimized(m, v319, v326)
	mBase = m.M
	v353 = v350
	goto L66
L70:
	;
	v331 = v326
	v332 = v319
	v333 = v314
	goto L71
L71:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+3)))
	v335 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_choose_best_statistics[0]))))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+2)))
	v337 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v336)+uint32(_c_F_choose_best_statistics[0]))))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	v339 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v338)+uint32(_c_F_choose_best_statistics[0]))))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v341 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v340)+uint32(_c_F_choose_best_statistics[0]))))
	v345 = v335 + (v337 + (v339 + (v333 + v341)))
	v346 = int32(4)
	v349 = v331 - v346
	if v349 != 0 {
		v331 = v349
		v332 = v332 + v346
		v333 = v345
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v353 = v345
	goto L66
L73:
	;
	goto L72
L74:
	;
	F_bms_free(m, v301)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L48
	} else {
		goto L89
	}
L75:
	;
	v404 = int32(0)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v365 = v302 + int32(8)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	if v366 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v404 = base.I32_popcnt(v369)
	goto L74
L79:
	;
	goto L80
L80:
	;
	v372 = v366 << (uint(int32(2)) % 32)
	if v372 <= int32(7) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v404 = base.I32_wrap_i64(v399)
	goto L74
L82:
	;
	if v372 == int32(0) {
		v399 = v360
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v396 = F_pg_popcount_optimized(m, v365, v372)
	mBase = m.M
	v399 = v396
	goto L81
L85:
	;
	v377 = v372
	v378 = v365
	v379 = v360
	goto L86
L86:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+3)))
	v381 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v380)+uint32(_c_F_choose_best_statistics[0]))))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+2)))
	v383 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v382)+uint32(_c_F_choose_best_statistics[0]))))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	v385 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v384)+uint32(_c_F_choose_best_statistics[0]))))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v387 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v386)+uint32(_c_F_choose_best_statistics[0]))))
	v391 = v381 + (v383 + (v385 + (v379 + v387)))
	v392 = int32(4)
	v395 = v377 - v392
	if v395 != 0 {
		v377 = v395
		v378 = v378 + v392
		v379 = v391
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v399 = v391
	goto L81
L88:
	;
	goto L87
L89:
	;
	F_bms_free(m, v302)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L48
	} else {
		goto L90
	}
L90:
	;
	v409 = v404 + v358
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v412 = int64(0)
	if v410 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	if v458 != 0 {
		goto L106
	} else {
		goto L107
	}
L92:
	;
	v456 = int32(0)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v417 = v410 + int32(8)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v418 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v456 = base.I32_popcnt(v421)
	goto L91
L96:
	;
	goto L97
L97:
	;
	v424 = v418 << (uint(int32(2)) % 32)
	if v424 <= int32(7) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v456 = base.I32_wrap_i64(v451)
	goto L91
L99:
	;
	if v424 == int32(0) {
		v451 = v412
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v448 = F_pg_popcount_optimized(m, v417, v424)
	mBase = m.M
	v451 = v448
	goto L98
L102:
	;
	v429 = v424
	v430 = v417
	v431 = v412
	goto L103
L103:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+3)))
	v433 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v432)+uint32(_c_F_choose_best_statistics[0]))))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+2)))
	v435 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_choose_best_statistics[0]))))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+1)))
	v437 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v436)+uint32(_c_F_choose_best_statistics[0]))))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v439 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v438)+uint32(_c_F_choose_best_statistics[0]))))
	v443 = v433 + (v435 + (v437 + (v431 + v439)))
	v444 = int32(4)
	v447 = v429 - v444
	if v447 != 0 {
		v429 = v447
		v430 = v430 + v444
		v431 = v443
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v451 = v443
	goto L98
L105:
	;
	goto L104
L106:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	v461 = v459
	goto L108
L107:
	;
	v461 = int32(0)
	goto L108
L108:
	;
	v462 = v461 + v456
	if (base.B2i32(v409 != v46)|base.B2i32(v50 <= v462))&base.B2i32(v409 <= v46) != 0 {
		v480 = v46
		v484 = v50
		v487 = v53
		goto L9
	} else {
		goto L109
	}
L109:
	;
	v480 = v409
	v484 = v462
	v487 = v58
	goto L9
L110:
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
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v104 == int32(-1) {
		v397 = int32(0)
		goto L32
	} else {
		goto L33
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
	v30 = int64(0)
	if v24 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if base.B2i32(v74 == v75)|base.B2i32(v75 <= int32(0)) != 0 {
		goto L3
	} else {
		goto L24
	}
L10:
	;
	v74 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = v24 + int32(8)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v36 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v74 = base.I32_popcnt(v39)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v42 = v36 << (uint(int32(2)) % 32)
	if v42 <= int32(7) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v74 = base.I32_wrap_i64(v69)
	goto L9
L17:
	;
	if v42 == int32(0) {
		v69 = v30
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v66 = F_pg_popcount_optimized(m, v35, v42)
	mBase = m.M
	v69 = v66
	goto L16
L20:
	;
	v47 = v42
	v48 = v35
	v49 = v30
	goto L21
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+3)))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_choose_next_subplan_for_worker[0]))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
	v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_choose_next_subplan_for_worker[0]))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_choose_next_subplan_for_worker[0]))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_choose_next_subplan_for_worker[0]))))
	v61 = v51 + (v53 + (v55 + (v49 + v57)))
	v62 = int32(4)
	v65 = v47 - v62
	if v65 != 0 {
		v47 = v65
		v48 = v48 + v62
		v49 = v61
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v69 = v61
	goto L16
L23:
	;
	goto L22
L24:
	;
	v81 = v20
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v86 = F_bms_is_member(m, v81, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L3
L27:
	;
	if v86 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v92 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v81)+20)) = uint8(v92)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v95 = v81 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v95 < v96 {
		v81 = v95
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	F_LWLockRelease(m, v6)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L95
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v104
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v112 = v110
	goto L35
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v112
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v256 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+(v6+int32(20))))))
	if v117 != int32(1) {
		goto L34
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(-1)
	F_LWLockRelease(m, v6)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L65
	}
L37:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v120 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v244
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v244 != v247 {
		v112 = v244
		goto L35
	} else {
		goto L64
	}
L39:
	;
	if int32(0) <= v176 {
		v244 = v176
		goto L38
	} else {
		goto L50
	}
L40:
	;
	v176 = base.I32_ctz(v162) | v163<<(uint(int32(5))%32)
	goto L39
L41:
	;
	v176 = int32(-2)
	goto L39
L42:
	;
	v127 = v112 + int32(1)
	v129 = int32(base.Ui32(v127) >> (uint(int32(5)) % 32))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v130 <= v129 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v133 = v120 + int32(8)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v129<<(uint(int32(2))%32))))
	v140 = v137 & (int32(-1) << (uint(v127) % 32))
	if v140 != 0 {
		v162 = v140
		v163 = v129
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v142 = v129 + int32(1)
	if v142 == v130 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v145 = v142
	goto L46
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v133+v145<<(uint(int32(2))%32))))
	if v152 != 0 {
		v162 = v152
		v163 = v145
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v154 = v145 + int32(1)
	if v154 != v130 {
		v145 = v154
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v179 <= v180 {
		v244 = v179
		goto L38
	} else {
		goto L51
	}
L51:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v182 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if int32(0) <= v240 {
		v244 = v240
		goto L38
	} else {
		goto L63
	}
L53:
	;
	v240 = base.I32_ctz(v226) | v227<<(uint(int32(5))%32)
	goto L52
L54:
	;
	v240 = int32(-2)
	goto L52
L55:
	;
	v191 = v180 - int32(1) + int32(1)
	v193 = int32(base.Ui32(v191) >> (uint(int32(5)) % 32))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v194 <= v193 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v197 = v182 + int32(8)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v193<<(uint(int32(2))%32))))
	v204 = v201 & (int32(-1) << (uint(v191) % 32))
	if v204 != 0 {
		v226 = v204
		v227 = v193
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v206 = v193 + int32(1)
	if v206 == v194 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v209 = v206
	goto L59
L59:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v197+v209<<(uint(int32(2))%32))))
	if v216 != 0 {
		v226 = v216
		v227 = v209
		goto L53
	} else {
		goto L61
	}
L60:
	;
	goto L54
L61:
	;
	v218 = v209 + int32(1)
	if v218 != v194 {
		v209 = v218
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v244 = v243
	goto L38
L64:
	;
	goto L36
L65:
	;
	return int32(0)
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v313
	if v313 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v313 = base.I32_ctz(v299) | v300<<(uint(int32(5))%32)
	goto L66
L68:
	;
	v313 = int32(-2)
	goto L66
L69:
	;
	v264 = v257 + int32(1)
	v266 = int32(base.Ui32(v264) >> (uint(int32(5)) % 32))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v267 <= v266 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v270 = v256 + int32(8)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v266<<(uint(int32(2))%32))))
	v277 = v274 & (int32(-1) << (uint(v264) % 32))
	if v277 != 0 {
		v299 = v277
		v300 = v266
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v279 = v266 + int32(1)
	if v279 == v267 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v282 = v279
	goto L73
L73:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v270+v282<<(uint(int32(2))%32))))
	if v289 != 0 {
		v299 = v289
		v300 = v282
		goto L67
	} else {
		goto L75
	}
L74:
	;
	goto L68
L75:
	;
	v291 = v282 + int32(1)
	if v291 != v267 {
		v282 = v291
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v318 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	goto L79
L79:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v385 <= v384 {
		v397 = int32(1)
		goto L32
	} else {
		goto L94
	}
L80:
	;
	if v377 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v377 = base.I32_ctz(v363) | v364<<(uint(int32(5))%32)
	goto L80
L82:
	;
	v377 = int32(-2)
	goto L80
L83:
	;
	v328 = v319 - int32(1) + int32(1)
	v330 = int32(base.Ui32(v328) >> (uint(int32(5)) % 32))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v331 <= v330 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v334 = v318 + int32(8)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v334+v330<<(uint(int32(2))%32))))
	v341 = v338 & (int32(-1) << (uint(v328) % 32))
	if v341 != 0 {
		v363 = v341
		v364 = v330
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v343 = v330 + int32(1)
	if v343 == v331 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v346 = v343
	goto L87
L87:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v334+v346<<(uint(int32(2))%32))))
	if v353 != 0 {
		v363 = v353
		v364 = v346
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L82
L89:
	;
	v355 = v346 + int32(1)
	if v355 != v331 {
		v346 = v355
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v380 = int32(-1)
	goto L93
L92:
	;
	v380 = v377
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v380
	goto L79
L94:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v387+v384)+20)) = uint8(v389)
	v397 = v389
	goto L32
L95:
	;
	return v397
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v7 != 0 {
		v178 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v178
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
	if v161 < int32(0) {
		goto L39
	} else {
		goto L40
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
	v161 = v86
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
	v39 = int32(base.Ui32(v37) >> (uint(int32(5)) % 32))
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
	v161 = v154
	goto L9
L25:
	;
	v154 = base.I32_clz(v139) | v138<<(uint(int32(5))%32) ^ int32(31)
	goto L24
L26:
	;
	v154 = int32(-2)
	goto L24
L27:
	;
	v93 = v26 + int32(8)
	if v8 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v99 = v96 << (uint(int32(5)) % 32)
	goto L30
L29:
	;
	v99 = v8
	goto L30
L30:
	;
	v101 = v99 - int32(1)
	v103 = int32(base.Ui32(v101) >> (uint(int32(5)) % 32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v93+v103<<(uint(int32(2))%32))))
	v108 = int32(-1)
	v112 = v107 & int32(base.Ui32(v108)>>(uint(v101^v108)%32))
	if v112 != 0 {
		v138 = v103
		v139 = v112
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v103 == int32(0) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v117 = v103
	goto L33
L33:
	;
	v122 = v117 - int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v93+v122<<(uint(int32(2))%32))))
	if v126 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L26
L35:
	;
	v138 = v122
	v139 = v126
	goto L25
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(int32(1)) < base.Ui32(v117) {
		v117 = v122
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v164 <= int32(0) {
		v178 = v2
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v161
	v178 = int32(1)
	goto L1
L42:
	;
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(v167)
	return int32(0)
}
