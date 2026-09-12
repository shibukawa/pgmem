package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_hash_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
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
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
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
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(1)
	v14 = v9 + v13
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v19 = v17 & v13
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v14
	goto L5
L4:
	;
	v20 = v9 + int32(4)
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v17 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v52 = F_str_tolower(m, v20, v50, int32(100))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v25 = int32(4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v27&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v40 = int32(1)
	if v19 != 0 {
		v50 = int32(base.Ui32(v17)>>(uint(v40)%32)) - v40
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v36 = v25
	goto L12
L11:
	;
	v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
	goto L12
L12:
	;
	if v27 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = v25
	goto L15
L14:
	;
	v39 = v36
	goto L15
L15:
	;
	v50 = v39
	goto L6
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	if v52&int32(3) == int32(0) {
		v77 = v52
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v116 = v110 - int32(1636608432)
	if v22 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v110 = v102 - v52
	goto L18
L20:
	;
	v81 = v77
	goto L29
L21:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v61 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v110 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v66 = v52
	goto L25
L25:
	;
	v70 = v66 + int32(1)
	if v70&int32(3) == int32(0) {
		v77 = v70
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v102 = v70
	goto L19
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 != 0 {
		v66 = v70
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 == v90 {
		v81 = v81 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v96 = v81
	goto L32
L31:
	;
	goto L30
L32:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		v96 = v96 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v102 = v96
	goto L19
L34:
	;
	goto L33
L35:
	;
	v426 = F_Int64GetDatum(m, base.I64_extend_i32_u(v416)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v408^v416-base.I32_rotl(v416, int32(24))))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L78
	}
L36:
	;
	if v52&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	v153 = v116
	v155 = v116
	v157 = v116
	goto L36
L38:
	;
	goto L39
L39:
	;
	v120 = v116 + base.I32_wrap_i64(v22)
	v121 = v120 + v116
	v125 = int32(4)
	v127 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(32))%64))) ^ base.I32_rotl(v116, v125)
	v131 = v120 - v127 ^ base.I32_rotl(v127, int32(6))
	v135 = v121 - v131 ^ base.I32_rotl(v131, int32(8))
	v136 = v127 + v121
	v137 = v131 + v136
	v138 = v135 + v137
	v142 = v136 - v135 ^ base.I32_rotl(v135, int32(16))
	v146 = v137 - v142 ^ base.I32_rotl(v142, int32(19))
	v151 = v142 + v138
	v153 = v151
	v155 = v138 - v146 ^ base.I32_rotl(v146, v125)
	v157 = v146 + v151
	goto L36
L40:
	;
	v394 = int32(14)
	v396 = v390 ^ v391 - base.I32_rotl(v390, v394)
	v400 = v396 ^ v389 - base.I32_rotl(v396, int32(11))
	v404 = v400 ^ v390 - base.I32_rotl(v400, int32(25))
	v408 = v404 ^ v396 - base.I32_rotl(v404, int32(16))
	v412 = v408 ^ v400 - base.I32_rotl(v408, int32(4))
	v416 = v412 ^ v404 - base.I32_rotl(v412, v394)
	goto L35
L41:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v389 = v381 + v384
	v390 = v382
	v391 = v383
	goto L40
L42:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	v381 = v377<<(uint(int32(8))%32) + v374
	v382 = v375
	v383 = v376
	goto L41
L43:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+2)))
	v374 = v370<<(uint(int32(16))%32) + v367
	v375 = v368
	v376 = v369
	goto L42
L44:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+3)))
	v367 = v363<<(uint(int32(24))%32) + v214
	v368 = v361
	v369 = v362
	goto L43
L45:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
	v361 = v357 + v359
	v362 = v358
	goto L44
L46:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
	v357 = v353<<(uint(int32(8))%32) + v351
	v358 = v352
	goto L45
L47:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+6)))
	v351 = v347<<(uint(int32(16))%32) + v345
	v352 = v346
	goto L46
L48:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+7)))
	v345 = v341<<(uint(int32(24))%32) + v215
	v346 = v340
	goto L47
L49:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
	v340 = v336<<(uint(int32(8))%32) + v335
	goto L48
L50:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+9)))
	v335 = v331<<(uint(int32(16))%32) + v330
	goto L49
L51:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+10)))
	v330 = v326<<(uint(int32(24))%32) + v216
	goto L50
L52:
	;
	if base.Ui32(int32(11)) < base.Ui32(v110) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v110) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v162 = v52
	v163 = v110
	v165 = v153
	v166 = v157
	v167 = v155
	goto L58
L56:
	;
	v211 = v52
	v212 = v110
	v214 = v153
	v215 = v157
	v216 = v155
	goto L57
L57:
	;
	switch v212 - int32(1) {
	case 0:
		v381 = v214
		v382 = v215
		v383 = v216
		goto L41
	case 1:
		v374 = v214
		v375 = v215
		v376 = v216
		goto L42
	case 2:
		v367 = v214
		v368 = v215
		v369 = v216
		goto L43
	case 3:
		v361 = v215
		v362 = v216
		goto L44
	case 4:
		v357 = v215
		v358 = v216
		goto L45
	case 5:
		v351 = v215
		v352 = v216
		goto L46
	case 6:
		v345 = v215
		v346 = v216
		goto L47
	case 7:
		v340 = v216
		goto L48
	case 8:
		v335 = v216
		goto L49
	case 9:
		v330 = v216
		goto L50
	case 10:
		goto L51
	default:
		v389 = v214
		v390 = v215
		v391 = v216
		goto L40
	}
L58:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v170 = v169 + v166
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v174 = v173 + v167
	v176 = int32(4)
	v178 = v171 + v165 - v174 ^ base.I32_rotl(v174, v176)
	v182 = v170 - v178 ^ base.I32_rotl(v178, int32(6))
	v183 = v174 + v170
	v184 = v178 + v183
	v185 = v182 + v184
	v189 = v183 - v182 ^ base.I32_rotl(v182, int32(8))
	v193 = v184 - v189 ^ base.I32_rotl(v189, int32(16))
	v197 = v185 - v193 ^ base.I32_rotl(v193, int32(19))
	v198 = v189 + v185
	v199 = v193 + v198
	v200 = v197 + v199
	v204 = v198 - v197 ^ base.I32_rotl(v197, v176)
	v205 = int32(12)
	v206 = v162 + v205
	v208 = v163 - v205
	if base.Ui32(int32(11)) < base.Ui32(v208) {
		v162 = v206
		v163 = v208
		v165 = v199
		v166 = v200
		v167 = v204
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v211 = v206
	v212 = v208
	v214 = v199
	v215 = v200
	v216 = v204
	goto L57
L60:
	;
	goto L59
L61:
	;
	v222 = v52
	v223 = v110
	v225 = v153
	v226 = v157
	v227 = v155
	goto L64
L62:
	;
	v271 = v52
	v272 = v110
	v274 = v153
	v275 = v157
	v276 = v155
	goto L63
L63:
	;
	switch v272 - int32(1) {
	case 0:
		v323 = v274
		goto L67
	case 1:
		v318 = v274
		goto L68
	case 2:
		goto L69
	case 3:
		v311 = v275
		goto L70
	case 4:
		v308 = v275
		goto L71
	case 5:
		v303 = v275
		goto L72
	case 6:
		goto L73
	case 7:
		v294 = v276
		goto L74
	case 8:
		v289 = v276
		goto L75
	case 9:
		v284 = v276
		goto L76
	case 10:
		goto L77
	default:
		v389 = v274
		v390 = v275
		v391 = v276
		goto L40
	}
L64:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v230 = v229 + v226
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v234 = v233 + v227
	v236 = int32(4)
	v238 = v231 + v225 - v234 ^ base.I32_rotl(v234, v236)
	v242 = v230 - v238 ^ base.I32_rotl(v238, int32(6))
	v243 = v234 + v230
	v244 = v238 + v243
	v245 = v242 + v244
	v249 = v243 - v242 ^ base.I32_rotl(v242, int32(8))
	v253 = v244 - v249 ^ base.I32_rotl(v249, int32(16))
	v257 = v245 - v253 ^ base.I32_rotl(v253, int32(19))
	v258 = v249 + v245
	v259 = v253 + v258
	v260 = v257 + v259
	v264 = v258 - v257 ^ base.I32_rotl(v257, v236)
	v265 = int32(12)
	v266 = v222 + v265
	v268 = v223 - v265
	if base.Ui32(int32(11)) < base.Ui32(v268) {
		v222 = v266
		v223 = v268
		v225 = v259
		v226 = v260
		v227 = v264
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v271 = v266
	v272 = v268
	v274 = v259
	v275 = v260
	v276 = v264
	goto L63
L66:
	;
	goto L65
L67:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v389 = v323 + v324
	v390 = v275
	v391 = v276
	goto L40
L68:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v323 = v319<<(uint(int32(8))%32) + v318
	goto L67
L69:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+2)))
	v318 = v314<<(uint(int32(16))%32) + v274
	goto L68
L70:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v389 = v312 + v274
	v390 = v311
	v391 = v276
	goto L40
L71:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+4)))
	v311 = v308 + v309
	goto L70
L72:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+5)))
	v308 = v304<<(uint(int32(8))%32) + v303
	goto L71
L73:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+6)))
	v303 = v299<<(uint(int32(16))%32) + v275
	goto L72
L74:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	v389 = v295 + v274
	v390 = v297 + v275
	v391 = v294
	goto L40
L75:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+8)))
	v294 = v290<<(uint(int32(8))%32) + v289
	goto L74
L76:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+9)))
	v289 = v285<<(uint(int32(16))%32) + v284
	goto L75
L77:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+10)))
	v284 = v280<<(uint(int32(24))%32) + v276
	goto L76
L78:
	;
	F_pfree(m, v52)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v430 != v9 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_pfree(m, v9)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	return v426
L83:
	;
	goto L82
}
func F_citext_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_citext_pattern_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v13)
							}
						} else {
							return base.B2i32(int32(0) < v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v13)
						}
					} else {
						return base.B2i32(int32(0) < v13)
					}
				}
			}
		}
	}
}
func F_citext_pattern_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v13 <= int32(0))
							}
						} else {
							return base.B2i32(v13 <= int32(0))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v13 <= int32(0))
						}
					} else {
						return base.B2i32(v13 <= int32(0))
					}
				}
			}
		}
	}
}
