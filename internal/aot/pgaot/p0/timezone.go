package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_timezone_abbreviations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
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
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
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
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
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
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v496 int32
	_ = v496
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v18 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(13782), v18, int32(1024), int32(8192))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v33 = int32(4464496)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(128)
	v40 = F_palloc(m, int32(3072))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v40
	v43 = int32(0)
	v49 = F_ParseTzFile(m, v13, v43, v21+int32(12), v21+int32(8), v43)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L123
	}
L8:
	;
	if v49 < int32(0) {
		v483 = v18
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v57 = v49<<(uint(int32(4))%32) | int32(8)
	if int32(0) < v49 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = v57
	v63 = v4
	goto L13
L11:
	;
	v144 = v57
	goto L12
L12:
	;
	v154 = F_guc_malloc(m, v144)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L37
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v53+v63*int32(24))+4))
	if v75 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v144 = v138
	goto L12
L15:
	;
	if v75&int32(3) == int32(0) {
		v99 = v75
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v138 = v62
	goto L17
L17:
	;
	v140 = v63 + int32(1)
	if v140 != v49 {
		v62 = v138
		v63 = v140
		goto L13
	} else {
		goto L35
	}
L18:
	;
	v138 = (v132+int32(12))&int32(-8) + v62
	goto L17
L19:
	;
	v132 = v124 - v75
	goto L18
L20:
	;
	v103 = v99
	goto L29
L21:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v83 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v132 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v88 = v75
	goto L25
L25:
	;
	v92 = v88 + int32(1)
	if v92&int32(3) == int32(0) {
		v99 = v92
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v124 = v92
	goto L19
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v97 != 0 {
		v88 = v92
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v112 = int32(-2139062144)
	if (int32(16843008)-v109|v109)&v112 == v112 {
		v103 = v103 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v118 = v103
	goto L32
L31:
	;
	goto L30
L32:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 != 0 {
		v118 = v118 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v124 = v118
	goto L19
L34:
	;
	goto L33
L35:
	;
	goto L14
L36:
	;
	if v154 != 0 {
		v483 = v154
		goto L7
	} else {
		goto L120
	}
L37:
	;
	if v154 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v144
	if v49 <= int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v168 = int32(0)
	v171 = v57
	goto L40
L40:
	;
	v179 = v154 + int32(8) + v168<<(uint(int32(4))%32)
	v182 = v53 + v168*int32(24)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	goto L45
L41:
	;
	goto L36
L42:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v299 != 0 {
		goto L75
	} else {
		goto L76
	}
L43:
	;
	v296 = F_strlen(m, v285)
	mBase = m.M
	goto L42
L45:
	;
	goto L46
L46:
	;
	v190 = int32(10)
	if (v179^v183)&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v289)
	goto L43
L48:
	;
	v270 = v265
	v271 = v266
	v272 = v267
	goto L70
L49:
	;
	if v260 == int32(0) {
		v285 = v258
		v286 = v259
		goto L47
	} else {
		goto L69
	}
L50:
	;
	v258 = v183
	v259 = v179
	v260 = v190
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v183&int32(3) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v227 == int32(0) {
		v285 = v224
		v286 = v225
		goto L47
	} else {
		goto L62
	}
L54:
	;
	v224 = v183
	v225 = v179
	v226 = v190
	v227 = int32(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v203 = v183
	v204 = v179
	v205 = v190
	goto L57
L57:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v207)
	if v207 == int32(0) {
		v265 = v203
		v266 = v204
		v267 = v205
		goto L48
	} else {
		goto L59
	}
L58:
	;
	v224 = v218
	v225 = v212
	v226 = v214
	v227 = v216
	goto L53
L59:
	;
	v211 = int32(1)
	v212 = v204 + v211
	v214 = v205 - v211
	v215 = int32(0)
	v216 = base.B2i32(v214 != v215)
	v218 = v203 + v211
	if v218&int32(3) == v215 {
		v224 = v218
		v225 = v212
		v226 = v214
		v227 = v216
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v214 != 0 {
		v203 = v218
		v204 = v212
		v205 = v214
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v230 == int32(0) {
		v258 = v224
		v259 = v225
		v260 = v226
		goto L49
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(v226) < base.Ui32(int32(4)) {
		v258 = v224
		v259 = v225
		v260 = v226
		goto L49
	} else {
		goto L64
	}
L64:
	;
	v236 = v224
	v237 = v225
	v238 = v226
	goto L65
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v244 = int32(-2139062144)
	if (int32(16843008)-v241|v241)&v244 != v244 {
		v265 = v236
		v266 = v237
		v267 = v238
		goto L48
	} else {
		goto L67
	}
L66:
	;
	v258 = v252
	v259 = v250
	v260 = v254
	goto L49
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v241
	v249 = int32(4)
	v250 = v237 + v249
	v252 = v236 + v249
	v254 = v238 - v249
	if base.Ui32(int32(3)) < base.Ui32(v254) {
		v236 = v252
		v237 = v250
		v238 = v254
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v265 = v258
	v266 = v259
	v267 = v260
	goto L48
L70:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v274)
	if v274 == int32(0) {
		v285 = v270
		v286 = v271
		goto L47
	} else {
		goto L72
	}
L71:
	;
	v285 = v281
	v286 = v279
	goto L47
L72:
	;
	v278 = int32(1)
	v279 = v271 + v278
	v281 = v270 + v278
	v283 = v272 - v278
	if v283 != 0 {
		v270 = v281
		v271 = v279
		v272 = v283
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v456 = v168 + int32(1)
	if v456 != v49 {
		v168 = v456
		v171 = v453
		goto L40
	} else {
		goto L119
	}
L75:
	;
	v300 = v154 + v171
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = int32(0)
	v304 = v300 + int32(4)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if (v305^v304)&int32(3) != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L77
L77:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+12)))
	if v448 != 0 {
		goto L116
	} else {
		goto L117
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+12)) = v171
	v381 = int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+11)) = uint8(v381)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v383&int32(3) == int32(0) {
		v407 = v383
		goto L101
	} else {
		goto L102
	}
L79:
	;
	goto L78
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v360))) = uint8(v359)
	if v359&int32(255) == int32(0) {
		goto L79
	} else {
		goto L95
	}
L81:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v358 = v305
	v359 = v311
	v360 = v304
	goto L80
L82:
	;
	goto L83
L83:
	;
	if v305&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v315 = v305
	v317 = v304
	goto L87
L85:
	;
	v329 = v305
	v331 = v304
	goto L86
L86:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v336 = int32(-2139062144)
	if (int32(16843008)-v333|v333)&v336 != v336 {
		v358 = v329
		v359 = v333
		v360 = v331
		goto L80
	} else {
		goto L91
	}
L87:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v318)
	if v318 == int32(0) {
		goto L79
	} else {
		goto L89
	}
L88:
	;
	v329 = v325
	v331 = v323
	goto L86
L89:
	;
	v322 = int32(1)
	v323 = v317 + v322
	v325 = v315 + v322
	if v325&int32(3) != 0 {
		v315 = v325
		v317 = v323
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v341 = v329
	v342 = v333
	v343 = v331
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343))) = v342
	v345 = int32(4)
	v346 = v343 + v345
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v349 = v341 + v345
	v353 = int32(-2139062144)
	if (v347|(int32(16843008)-v347))&v353 == v353 {
		v341 = v349
		v342 = v347
		v343 = v346
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v358 = v349
	v359 = v347
	v360 = v346
	goto L80
L94:
	;
	goto L93
L95:
	;
	v367 = v358
	v369 = v360
	goto L96
L96:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)) = uint8(v370)
	v372 = int32(1)
	if v370 != 0 {
		v367 = v367 + v372
		v369 = v369 + v372
		goto L96
	} else {
		goto L98
	}
L97:
	;
	goto L79
L98:
	;
	goto L97
L99:
	;
	v453 = (v440+int32(12))&int32(-8) + v171
	goto L74
L100:
	;
	v440 = v432 - v383
	goto L99
L101:
	;
	v411 = v407
	goto L110
L102:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	if v391 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v440 = int32(0)
	goto L99
L104:
	;
	goto L105
L105:
	;
	v396 = v383
	goto L106
L106:
	;
	v400 = v396 + int32(1)
	if v400&int32(3) == int32(0) {
		v407 = v400
		goto L101
	} else {
		goto L108
	}
L107:
	;
	v432 = v400
	goto L100
L108:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v405 != 0 {
		v396 = v400
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v420 = int32(-2139062144)
	if (int32(16843008)-v417|v417)&v420 == v420 {
		v411 = v411 + int32(4)
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v426 = v411
	goto L113
L112:
	;
	goto L111
L113:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v430 != 0 {
		v426 = v426 + int32(1)
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v432 = v426
	goto L100
L115:
	;
	goto L114
L116:
	;
	v449 = int32(6)
	goto L118
L117:
	;
	v449 = int32(5)
	goto L118
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+11)) = uint8(v449)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+12)) = v451
	v453 = v171
	goto L74
L119:
	;
	goto L41
L120:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v472
	goto L121
L121:
	;
	v478 = F_format_elog_string(m, int32(13575), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v478
	v483 = int32(0)
	goto L7
L123:
	;
	m.G0 = v21 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v483
	return base.B2i32(v483 != int32(0))
}
