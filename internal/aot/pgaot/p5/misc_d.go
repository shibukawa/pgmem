package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DCH_cache_fetch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
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
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	v2 = l1
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1]))
	if int32(2147483646) <= v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v116 = v14
	goto L3
L3:
	;
	if v12 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v108 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1])) = v108
	v116 = v108
	goto L3
L5:
	;
	v20 = v12 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v12) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = v3
	v35 = v3
	goto L9
L7:
	;
	v67 = v3
	goto L8
L8:
	;
	v77 = v67
	v80 = int32(0)
	goto L13
L9:
	;
	v37 = v28 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_DCH_cache_fetch[2])))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+2032))
	v40 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+2032)) = v39 >> (uint(v40) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_DCH_cache_fetch[3])))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2032)) = v44 >> (uint(v40) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_DCH_cache_fetch[4])))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+2032)) = v49 >> (uint(v40) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_DCH_cache_fetch[5])))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+2032)) = v54 >> (uint(v40) % 32)
	v58 = int32(4)
	v59 = v28 + v58
	v61 = v35 + v58
	if v61 != v12&int32(2147483644) {
		v28 = v59
		v35 = v61
		goto L9
	} else {
		goto L11
	}
L10:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v67 = v59
	goto L8
L13:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(int32(2))%32))+uint32(_c_F_DCH_cache_fetch[2])))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+2032))
	v89 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+2032)) = v88 >> (uint(v89) % 32)
	v95 = v80 + v89
	if v95 != v20 {
		v77 = v77 + v89
		v80 = v95
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L4
L15:
	;
	goto L14
L16:
	;
	if v2 != 0 {
		goto L116
	} else {
		goto L117
	}
L17:
	;
	v394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v386)+2029)) = uint8(v394)
	v397 = v386 + int32(1872)
	goto L88
L18:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[6]))
	v236 = F_MemoryContextAllocZero(m, v234, int32(2036))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L52
	} else {
		goto L53
	}
L19:
	;
	v126 = int32(0)
	goto L21
L20:
	;
	v219 = v116 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1])) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v136)+2032)) = v219
	return v136
L21:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v126<<(uint(int32(2))%32))+uint32(_c_F_DCH_cache_fetch[2])))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+2029)))
	if v137 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v12 < int32(20) {
		goto L18
	} else {
		goto L35
	}
L23:
	;
	v171 = v126 + int32(1)
	if v171 != v12 {
		v126 = v171
		goto L21
	} else {
		goto L34
	}
L24:
	;
	v141 = v136 + int32(1872)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v144 == int32(0))|base.B2i32(v144 != v147) != 0 {
		v165 = v144
		v166 = v147
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v165-v166 != 0 {
		goto L23
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v150 = v141
	v151 = l0
	goto L28
L28:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	if v155 == int32(0) {
		v165 = v155
		v166 = v154
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v165 = v155
	v166 = v154
	goto L26
L30:
	;
	v158 = int32(1)
	if v155 == v154 {
		v150 = v150 + v158
		v151 = v151 + v158
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+2028)))
	if v168 == v2 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	goto L23
L34:
	;
	goto L22
L35:
	;
	v175 = int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[2]))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+2029)))
	if v178 != v175 {
		v386 = v177
		goto L17
	} else {
		goto L36
	}
L36:
	;
	v183 = v177
	v186 = v175
	goto L37
L37:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v186<<(uint(int32(2))%32))+uint32(_c_F_DCH_cache_fetch[2])))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2029)))
	if v194 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v386 = v193
	goto L17
L40:
	;
	goto L41
L41:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193)+2032))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2032))
	if v197 < v198 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v200 = v193
	goto L44
L43:
	;
	v200 = v183
	goto L44
L44:
	;
	v202 = v186 + int32(1)
	if v202 == int32(20) {
		v386 = v200
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v202<<(uint(int32(2))%32))+uint32(_c_F_DCH_cache_fetch[2])))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2029)))
	if v208 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v386 = v207
	goto L17
L47:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+2032))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200)+2032))
	if v211 < v212 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v214 = v207
	goto L51
L50:
	;
	v214 = v200
	goto L51
L51:
	;
	v183 = v214
	v186 = v186 + int32(2)
	goto L37
L52:
	;
	return int32(0)
L53:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v241<<(uint(int32(2))%32))+uint32(_c_F_DCH_cache_fetch[2]))) = v236
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+2029)) = uint8(v247)
	v250 = v236 + int32(1872)
	goto L57
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+2028)) = uint8(v2)
	v371 = int32(_a_F_DCH_cache_fetch_0)
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1]))
	v374 = int32(1)
	v375 = v373 + v374
	*(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1])) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v236)+2032)) = v375
	v378 = int32(_a_F_DCH_cache_fetch_1)
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[0])) = v380 + v374
	v526 = v236
	goto L16
L55:
	;
	v367 = F_strlen(m, v356)
	mBase = m.M
	goto L54
L57:
	;
	goto L58
L58:
	;
	v257 = int32(155)
	if (v250^l0)&int32(3) != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v360)
	goto L55
L60:
	;
	v341 = v336
	v342 = v337
	v343 = v338
	goto L81
L61:
	;
	if v331 == int32(0) {
		v356 = v329
		v357 = v330
		goto L59
	} else {
		goto L80
	}
L62:
	;
	v329 = l0
	v330 = v250
	v331 = v257
	goto L61
L63:
	;
	goto L64
L64:
	;
	v261 = int32(0)
	if base.B2i32(l0&int32(3) == v261)|int32(0) == v261 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v297 == int32(0) {
		v356 = v294
		v357 = v295
		goto L59
	} else {
		goto L74
	}
L66:
	;
	v273 = l0
	v274 = v250
	v275 = v257
	goto L69
L67:
	;
	goto L68
L68:
	;
	v294 = l0
	v295 = v250
	v296 = v257
	v297 = int32(1)
	goto L65
L69:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v277)
	if v277 == int32(0) {
		v336 = v273
		v337 = v274
		v338 = v275
		goto L60
	} else {
		goto L71
	}
L70:
	;
	v294 = v288
	v295 = v282
	v296 = v284
	v297 = v286
	goto L65
L71:
	;
	v281 = int32(1)
	v282 = v274 + v281
	v284 = v275 - v281
	v285 = int32(0)
	v286 = base.B2i32(v284 != v285)
	v288 = v273 + v281
	if v288&int32(3) == v285 {
		v294 = v288
		v295 = v282
		v296 = v284
		v297 = v286
		goto L65
	} else {
		goto L72
	}
L72:
	;
	if v284 != 0 {
		v273 = v288
		v274 = v282
		v275 = v284
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if base.B2i32(v300 == int32(0))|base.B2i32(base.Ui32(v296) < base.Ui32(int32(4))) != 0 {
		v329 = v294
		v330 = v295
		v331 = v296
		goto L61
	} else {
		goto L75
	}
L75:
	;
	v307 = v294
	v308 = v295
	v309 = v296
	goto L76
L76:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v315 = int32(-2139062144)
	if (int32(16843008)-v312|v312)&v315 != v315 {
		v336 = v307
		v337 = v308
		v338 = v309
		goto L60
	} else {
		goto L78
	}
L77:
	;
	v329 = v323
	v330 = v321
	v331 = v325
	goto L61
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v312
	v320 = int32(4)
	v321 = v308 + v320
	v323 = v307 + v320
	v325 = v309 - v320
	if base.Ui32(int32(3)) < base.Ui32(v325) {
		v307 = v323
		v308 = v321
		v309 = v325
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v336 = v329
	v337 = v330
	v338 = v331
	goto L60
L81:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v345)
	if v345 == int32(0) {
		v356 = v341
		v357 = v342
		goto L59
	} else {
		goto L83
	}
L82:
	;
	v356 = v352
	v357 = v350
	goto L59
L83:
	;
	v349 = int32(1)
	v350 = v342 + v349
	v352 = v341 + v349
	v354 = v343 - v349
	if v354 != 0 {
		v341 = v352
		v342 = v350
		v343 = v354
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v517 = int32(_a_F_DCH_cache_fetch_0)
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1]))
	v521 = v519 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_DCH_cache_fetch[1])) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v386)+2032)) = v521
	v526 = v386
	goto L16
L86:
	;
	v514 = F_strlen(m, v503)
	mBase = m.M
	goto L85
L88:
	;
	goto L89
L89:
	;
	v404 = int32(155)
	if (v397^l0)&int32(3) != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v507 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v504))) = uint8(v507)
	goto L86
L91:
	;
	v488 = v483
	v489 = v484
	v490 = v485
	goto L112
L92:
	;
	if v478 == int32(0) {
		v503 = v476
		v504 = v477
		goto L90
	} else {
		goto L111
	}
L93:
	;
	v476 = l0
	v477 = v397
	v478 = v404
	goto L92
L94:
	;
	goto L95
L95:
	;
	v408 = int32(0)
	if base.B2i32(l0&int32(3) == v408)|int32(0) == v408 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v444 == int32(0) {
		v503 = v441
		v504 = v442
		goto L90
	} else {
		goto L105
	}
L97:
	;
	v420 = l0
	v421 = v397
	v422 = v404
	goto L100
L98:
	;
	goto L99
L99:
	;
	v441 = l0
	v442 = v397
	v443 = v404
	v444 = int32(1)
	goto L96
L100:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	*(*uint8)(unsafe.Add(mBase, uint32(v421))) = uint8(v424)
	if v424 == int32(0) {
		v483 = v420
		v484 = v421
		v485 = v422
		goto L91
	} else {
		goto L102
	}
L101:
	;
	v441 = v435
	v442 = v429
	v443 = v431
	v444 = v433
	goto L96
L102:
	;
	v428 = int32(1)
	v429 = v421 + v428
	v431 = v422 - v428
	v432 = int32(0)
	v433 = base.B2i32(v431 != v432)
	v435 = v420 + v428
	if v435&int32(3) == v432 {
		v441 = v435
		v442 = v429
		v443 = v431
		v444 = v433
		goto L96
	} else {
		goto L103
	}
L103:
	;
	if v431 != 0 {
		v420 = v435
		v421 = v429
		v422 = v431
		goto L100
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441))))
	if base.B2i32(v447 == int32(0))|base.B2i32(base.Ui32(v443) < base.Ui32(int32(4))) != 0 {
		v476 = v441
		v477 = v442
		v478 = v443
		goto L92
	} else {
		goto L106
	}
L106:
	;
	v454 = v441
	v455 = v442
	v456 = v443
	goto L107
L107:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v462 = int32(-2139062144)
	if (int32(16843008)-v459|v459)&v462 != v462 {
		v483 = v454
		v484 = v455
		v485 = v456
		goto L91
	} else {
		goto L109
	}
L108:
	;
	v476 = v470
	v477 = v468
	v478 = v472
	goto L92
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = v459
	v467 = int32(4)
	v468 = v455 + v467
	v470 = v454 + v467
	v472 = v456 - v467
	if base.Ui32(int32(3)) < base.Ui32(v472) {
		v454 = v470
		v455 = v468
		v456 = v472
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v483 = v476
	v484 = v477
	v485 = v478
	goto L91
L112:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	*(*uint8)(unsafe.Add(mBase, uint32(v489))) = uint8(v492)
	if v492 == int32(0) {
		v503 = v488
		v504 = v489
		goto L90
	} else {
		goto L114
	}
L113:
	;
	v503 = v499
	v504 = v497
	goto L90
L114:
	;
	v496 = int32(1)
	v497 = v489 + v496
	v499 = v488 + v496
	v501 = v490 - v496
	if v501 != 0 {
		v488 = v499
		v489 = v497
		v490 = v501
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v539 = int32(5)
	goto L118
L117:
	;
	v539 = int32(1)
	goto L118
L118:
	;
	F_parse_format(m, v526, l0, int32(_a_F_DCH_cache_fetch_2), int32(_a_F_DCH_cache_fetch_3), int32(_a_F_DCH_cache_fetch_4), v539, int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L52
	} else {
		goto L119
	}
L119:
	;
	v543 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+2029)) = uint8(v543)
	return v526
}
func F_DecodeNumberField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 float64
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 float64
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
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
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
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
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(-1)
	v17 = int32(_a_F_DecodeNumberField_0)
	v21 = m.G0
	v23 = v21 - int32(32)
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v24
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeNumberField[0])))
	if v32 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v567
L2:
	;
	if v100 != l0 {
		v567 = v16
		goto L1
	} else {
		goto L21
	}
L3:
	;
	v100 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeNumberField[1])))
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = l1
	goto L9
L7:
	;
	goto L8
L8:
	;
	v50 = v17
	v51 = v32
	goto L12
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v46 == v32 {
		v40 = v40 + int32(1)
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v100 = v40 - l1
	goto L2
L11:
	;
	goto L10
L12:
	;
	v58 = v23 + int32(base.Ui32(v51)>>(uint(int32(3))%32))&int32(28)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59 | v60<<(uint(v51)%32)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v64 != 0 {
		v50 = v50 + v60
		v51 = v64
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v67 == int32(0) {
		v90 = l1
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v100 = v90 - l1
	goto L2
L16:
	;
	v71 = l1
	v72 = v67
	goto L17
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(base.Ui32(v72)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v80)>>(uint(v72)%32))&int32(1) == int32(0) {
		v90 = v71
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v90 = v88
	goto L15
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v88 = v71 + int32(1)
	if v86 != 0 {
		v71 = v88
		v72 = v86
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v102 = int32(46)
	v103 = F___strchrnul(m, l1, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v105 == v102 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v396 = int32(_a_F_DecodeNumberField_1)
	if l2&v396 == v396 {
		v567 = v16
		goto L1
	} else {
		goto L107
	}
L23:
	;
	if v109 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v109 = v103
	goto L26
L25:
	;
	v109 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v111 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v223 = int32(14)
	if base.B2i32(l2&v223 == v223)|base.B2i32(l0 < int32(6)) != 0 {
		v394 = l0
		goto L22
	} else {
		goto L57
	}
L30:
	;
	v113 = v109 + int32(1)
	v114 = int32(_a_F_DecodeNumberField_2)
	v118 = m.G0
	v120 = v118 - int32(32)
	v121 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v120)+24)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v120)+16)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v120)+8)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v120))) = v121
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeNumberField[2])))
	if v129 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v214 = float64(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v214, float64(1e+06))))
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v220)
	v222 = F_strlen(m, l1)
	mBase = m.M
	v394 = v222
	goto L22
L33:
	;
	v198 = F_strlen(m, v113)
	mBase = m.M
	if v197 != v198 {
		v567 = v16
		goto L1
	} else {
		goto L52
	}
L34:
	;
	v197 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeNumberField[3])))
	if v133 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v137 = v113
	goto L40
L38:
	;
	goto L39
L39:
	;
	v147 = v114
	v148 = v129
	goto L43
L40:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v143 == v129 {
		v137 = v137 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v197 = v137 - v113
	goto L33
L42:
	;
	goto L41
L43:
	;
	v155 = v120 + int32(base.Ui32(v148)>>(uint(int32(3))%32))&int32(28)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156 | v157<<(uint(v148)%32)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v161 != 0 {
		v147 = v147 + v157
		v148 = v161
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v164 == int32(0) {
		v187 = v113
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v197 = v187 - v113
	goto L33
L47:
	;
	v168 = v113
	v169 = v164
	goto L48
L48:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(base.Ui32(v169)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v177)>>(uint(v169)%32))&int32(1) == int32(0) {
		v187 = v168
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v187 = v185
	goto L46
L50:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v185 = v168 + int32(1)
	if v183 != 0 {
		v168 = v185
		v169 = v183
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeNumberField[4])) = int32(0)
	v205 = F_strtod(m, v109, v14+int32(12))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v210 != 0 {
		v567 = v16
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeNumberField[4]))
	if v212 != 0 {
		v567 = v16
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v214 = v205
	goto L32
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(14)
	v234 = l0 + l1 - int32(2)
	v238 = v234
	goto L59
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v282
	v284 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v284)
	v287 = l0 - int32(4)
	v288 = l1 + v287
	v292 = v288
	goto L75
L59:
	;
	v243 = v238 + int32(1)
	v244 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238))))
	v245 = F___isspace(m, v244)
	mBase = m.M
	if v245 != 0 {
		v238 = v243
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v246 = int32(1)
	switch v244&int32(255) - int32(43) {
	case 0:
		v252 = v246
		goto L63
	default:
		v254 = v244
		v255 = v238
		v256 = v246
		goto L62
	case 2:
		goto L64
	}
L61:
	;
	goto L60
L62:
	;
	v257 = int32(0)
	v259 = v254 - int32(48)
	if base.Ui32(v259) <= base.Ui32(int32(9)) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v253 = int32(*(*int8)(unsafe.Add(mBase, uint32(v243))))
	v254 = v253
	v255 = v243
	v256 = v252
	goto L62
L64:
	;
	v252 = int32(0)
	goto L63
L65:
	;
	v262 = v257
	v263 = v259
	v264 = v255
	goto L68
L66:
	;
	v276 = v257
	goto L67
L67:
	;
	if v256 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v266 = int32(10)
	v268 = v262*v266 - v263
	v269 = int32(*(*int8)(unsafe.Add(mBase, uint32(v264)+1)))
	v273 = v269 - int32(48)
	if base.Ui32(v273) < base.Ui32(v266) {
		v262 = v268
		v263 = v273
		v264 = v264 + int32(1)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v276 = v268
	goto L67
L70:
	;
	goto L69
L71:
	;
	v282 = int32(0) - v276
	goto L73
L72:
	;
	v282 = v276
	goto L73
L73:
	;
	goto L58
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v336
	v338 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v288))) = uint8(v338)
	v343 = l1
	goto L91
L75:
	;
	v297 = v292 + int32(1)
	v298 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292))))
	v299 = F___isspace(m, v298)
	mBase = m.M
	if v299 != 0 {
		v292 = v297
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v300 = int32(1)
	switch v298&int32(255) - int32(43) {
	case 0:
		v306 = v300
		goto L79
	default:
		v308 = v298
		v309 = v292
		v310 = v300
		goto L78
	case 2:
		goto L80
	}
L77:
	;
	goto L76
L78:
	;
	v311 = int32(0)
	v313 = v308 - int32(48)
	if base.Ui32(v313) <= base.Ui32(int32(9)) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v297))))
	v308 = v307
	v309 = v297
	v310 = v306
	goto L78
L80:
	;
	v306 = int32(0)
	goto L79
L81:
	;
	v316 = v311
	v317 = v313
	v318 = v309
	goto L84
L82:
	;
	v330 = v311
	goto L83
L83:
	;
	if v310 != 0 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v320 = int32(10)
	v322 = v316*v320 - v317
	v323 = int32(*(*int8)(unsafe.Add(mBase, uint32(v318)+1)))
	v327 = v323 - int32(48)
	if base.Ui32(v327) < base.Ui32(v320) {
		v316 = v322
		v317 = v327
		v318 = v318 + int32(1)
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v330 = v322
	goto L83
L86:
	;
	goto L85
L87:
	;
	v336 = int32(0) - v330
	goto L89
L88:
	;
	v336 = v330
	goto L89
L89:
	;
	goto L74
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v387
	v389 = int32(2)
	if v287 != v389 {
		v567 = v389
		goto L1
	} else {
		goto L106
	}
L91:
	;
	v348 = v343 + int32(1)
	v349 = int32(*(*int8)(unsafe.Add(mBase, uint32(v343))))
	v350 = F___isspace(m, v349)
	mBase = m.M
	if v350 != 0 {
		v343 = v348
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v351 = int32(1)
	switch v349&int32(255) - int32(43) {
	case 0:
		v357 = v351
		goto L95
	default:
		v359 = v349
		v360 = v343
		v361 = v351
		goto L94
	case 2:
		goto L96
	}
L93:
	;
	goto L92
L94:
	;
	v362 = int32(0)
	v364 = v359 - int32(48)
	if base.Ui32(v364) <= base.Ui32(int32(9)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v358 = int32(*(*int8)(unsafe.Add(mBase, uint32(v348))))
	v359 = v358
	v360 = v348
	v361 = v357
	goto L94
L96:
	;
	v357 = int32(0)
	goto L95
L97:
	;
	v367 = v362
	v368 = v364
	v369 = v360
	goto L100
L98:
	;
	v381 = v362
	goto L99
L99:
	;
	if v361 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v371 = int32(10)
	v373 = v367*v371 - v368
	v374 = int32(*(*int8)(unsafe.Add(mBase, uint32(v369)+1)))
	v378 = v374 - int32(48)
	if base.Ui32(v378) < base.Ui32(v371) {
		v367 = v373
		v368 = v378
		v369 = v369 + int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v381 = v373
	goto L99
L102:
	;
	goto L101
L103:
	;
	v387 = int32(0) - v381
	goto L105
L104:
	;
	v387 = v381
	goto L105
L105:
	;
	goto L90
L106:
	;
	v392 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v392)
	v567 = v389
	goto L1
L107:
	;
	switch v394 - int32(4) {
	case 0:
		goto L109
	default:
		v567 = v16
		goto L1
	case 2:
		goto L110
	}
L108:
	;
	v466 = l1 + int32(2)
	goto L128
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_DecodeNumberField_1)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_DecodeNumberField_1)
	v409 = l1 + int32(4)
	goto L112
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v453
	v455 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v455)
	goto L108
L112:
	;
	v414 = v409 + int32(1)
	v415 = int32(*(*int8)(unsafe.Add(mBase, uint32(v409))))
	v416 = F___isspace(m, v415)
	mBase = m.M
	if v416 != 0 {
		v409 = v414
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v417 = int32(1)
	switch v415&int32(255) - int32(43) {
	case 0:
		v423 = v417
		goto L116
	default:
		v425 = v415
		v426 = v409
		v427 = v417
		goto L115
	case 2:
		goto L117
	}
L114:
	;
	goto L113
L115:
	;
	v428 = int32(0)
	v430 = v425 - int32(48)
	if base.Ui32(v430) <= base.Ui32(int32(9)) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v424 = int32(*(*int8)(unsafe.Add(mBase, uint32(v414))))
	v425 = v424
	v426 = v414
	v427 = v423
	goto L115
L117:
	;
	v423 = int32(0)
	goto L116
L118:
	;
	v433 = v428
	v434 = v430
	v435 = v426
	goto L121
L119:
	;
	v447 = v428
	goto L120
L120:
	;
	if v427 != 0 {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v437 = int32(10)
	v439 = v433*v437 - v434
	v440 = int32(*(*int8)(unsafe.Add(mBase, uint32(v435)+1)))
	v444 = v440 - int32(48)
	if base.Ui32(v444) < base.Ui32(v437) {
		v433 = v439
		v434 = v444
		v435 = v435 + int32(1)
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v447 = v439
	goto L120
L123:
	;
	goto L122
L124:
	;
	v453 = int32(0) - v447
	goto L126
L125:
	;
	v453 = v447
	goto L126
L126:
	;
	goto L111
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v510
	v512 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v512)
	v517 = l1
	goto L144
L128:
	;
	v471 = v466 + int32(1)
	v472 = int32(*(*int8)(unsafe.Add(mBase, uint32(v466))))
	v473 = F___isspace(m, v472)
	mBase = m.M
	if v473 != 0 {
		v466 = v471
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v474 = int32(1)
	switch v472&int32(255) - int32(43) {
	case 0:
		v480 = v474
		goto L132
	default:
		v482 = v472
		v483 = v466
		v484 = v474
		goto L131
	case 2:
		goto L133
	}
L130:
	;
	goto L129
L131:
	;
	v485 = int32(0)
	v487 = v482 - int32(48)
	if base.Ui32(v487) <= base.Ui32(int32(9)) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v481 = int32(*(*int8)(unsafe.Add(mBase, uint32(v471))))
	v482 = v481
	v483 = v471
	v484 = v480
	goto L131
L133:
	;
	v480 = int32(0)
	goto L132
L134:
	;
	v490 = v485
	v491 = v487
	v492 = v483
	goto L137
L135:
	;
	v504 = v485
	goto L136
L136:
	;
	if v484 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v494 = int32(10)
	v496 = v490*v494 - v491
	v497 = int32(*(*int8)(unsafe.Add(mBase, uint32(v492)+1)))
	v501 = v497 - int32(48)
	if base.Ui32(v501) < base.Ui32(v494) {
		v490 = v496
		v491 = v501
		v492 = v492 + int32(1)
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v504 = v496
	goto L136
L139:
	;
	goto L138
L140:
	;
	v510 = int32(0) - v504
	goto L142
L141:
	;
	v510 = v504
	goto L142
L142:
	;
	goto L127
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v561
	v567 = int32(3)
	goto L1
L144:
	;
	v522 = v517 + int32(1)
	v523 = int32(*(*int8)(unsafe.Add(mBase, uint32(v517))))
	v524 = F___isspace(m, v523)
	mBase = m.M
	if v524 != 0 {
		v517 = v522
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v525 = int32(1)
	switch v523&int32(255) - int32(43) {
	case 0:
		v531 = v525
		goto L148
	default:
		v533 = v523
		v534 = v517
		v535 = v525
		goto L147
	case 2:
		goto L149
	}
L146:
	;
	goto L145
L147:
	;
	v536 = int32(0)
	v538 = v533 - int32(48)
	if base.Ui32(v538) <= base.Ui32(int32(9)) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v532 = int32(*(*int8)(unsafe.Add(mBase, uint32(v522))))
	v533 = v532
	v534 = v522
	v535 = v531
	goto L147
L149:
	;
	v531 = int32(0)
	goto L148
L150:
	;
	v541 = v536
	v542 = v538
	v543 = v534
	goto L153
L151:
	;
	v555 = v536
	goto L152
L152:
	;
	if v535 != 0 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v545 = int32(10)
	v547 = v541*v545 - v542
	v548 = int32(*(*int8)(unsafe.Add(mBase, uint32(v543)+1)))
	v552 = v548 - int32(48)
	if base.Ui32(v552) < base.Ui32(v545) {
		v541 = v547
		v542 = v552
		v543 = v543 + int32(1)
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v555 = v547
	goto L152
L155:
	;
	goto L154
L156:
	;
	v561 = int32(0) - v555
	goto L158
L157:
	;
	v561 = v555
	goto L158
L158:
	;
	goto L143
}
func F_DecrTupleDescRefCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_DecrTupleDescRefCount[0]))
	F_ResourceOwnerForget(m, v4, l0, int32(_a_F_DecrTupleDescRefCount_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = v8 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
		if v10 == int32(0) {
			F_FreeTupleDesc(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_DecrementParentLocks(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return
L3:
	;
	goto L2
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v26
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_DecrementParentLocks[0]))
	v33 = F_get_hash_value(m, v32, v7)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v20 == int32(-1) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v25 = v24
	goto L4
L8:
	;
	v25 = int32(-1)
	goto L4
L9:
	;
	return
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_DecrementParentLocks[1]))
	v37 = int32(0)
	v39 = F_hash_search_with_hash_value(m, v36, v7, v33, v37, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v39 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v45 = v43 - int32(1)
	v46 = int32(0)
	v48 = base.B2i32(v46 < v45)
	if v46 < v45 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v45
	goto L15
L14:
	;
	v49 = v46
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v49
	if v46 < v45 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+16)))
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_DecrementParentLocks[1]))
	v56 = F_hash_search_with_hash_value(m, v53, v7, v33, int32(2), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L1
}
func F_DeleteInheritsTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v16 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = v12 + int32(32)
	F_ScanKeyInit(m, v21, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = int32(1)
	v31 = F_systable_beginscan(m, v16, int32(2680), v28, int32(0), v28, v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v33 = F_systable_getnext(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = v33
	v43 = v5
	goto L11
L9:
	;
	v77 = v5
	goto L10
L10:
	;
	F_systable_endscan(m, v31)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v77 = v66
	goto L10
L13:
	;
	v67 = F_systable_getnext(m, v31)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 != l1 {
		v66 = v43
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+12)))
	v50 = int32(1)
	v51 = v49 ^ v50
	if l2|v51&v50 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if l2&v51 == int32(1) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_simple_heap_delete(m, v16, v35+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v66 = int32(1)
	goto L13
L21:
	;
	if v67 != 0 {
		v35 = v67
		v43 = v66
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L12
L23:
	;
	F_relation_close(m, v16, int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v12 + int32(80)
	return v77
L25:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if l3 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v95 = l3
	goto L29
L28:
	;
	v95 = int32(_a_F_DeleteInheritsTuple_0)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v95
	F_errmsg(m, int32(_a_F_DeleteInheritsTuple_1), v12+int32(16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errdetail(m, int32(_a_F_DeleteInheritsTuple_2), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(_a_F_DeleteInheritsTuple_3), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_DeleteInheritsTuple_4), int32(595), int32(_a_F_DeleteInheritsTuple_5))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if l3 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = l3
	goto L38
L37:
	;
	v123 = int32(_a_F_DeleteInheritsTuple_0)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v123
	F_errmsg(m, int32(_a_F_DeleteInheritsTuple_6), v12)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errdetail(m, int32(_a_F_DeleteInheritsTuple_7), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_DeleteInheritsTuple_4), int32(601), int32(_a_F_DeleteInheritsTuple_5))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DetermineTimeZoneOffset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int64
	_ = v81
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = v6 + int32(8)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v21 <= int32(-4713) {
		if v21 != int32(-4713) {
			v156 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
			v169 = v156
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if int32(10) < v26 {
				v37 = v26
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v41 = int32(60)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v52 = base.B2i32(int32(2) < v37)
				if int32(2) < v37 {
					v53 = int32(_a_F_DetermineTimeZoneOffset_0)
				} else {
					v53 = int32(_a_F_DetermineTimeZoneOffset_1)
				}
				v54 = v53 + v21
				v62 = base.I32_div_u_s(v54, int32(100))
				v65 = base.I32_div_u_s(v54, int32(400))
				if int32(2) < v37 {
					v69 = int32(1)
				} else {
					v69 = int32(13)
				}
				v74 = base.I32_div_s((v69+v37)*int32(_a_F_DetermineTimeZoneOffset_2), int32(256))
				v77 = v48 + v54*int32(365) + int32(base.Ui32(v54)>>(uint(int32(2))%32)) - v62 + v65 + v74 - int32(_a_F_DetermineTimeZoneOffset_3)
				v81 = base.I64_extend_i32_s(v38+(v39+v40*v41)*v41) + base.I64_extend_i32_s(v77)*int64(86400)
				if base.B2i32(int32(0) < v77)&base.B2i32(v81 < int64(0)) != 0 {
					v156 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
					v169 = v156
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v81 - int64(86400)
					v100 = F_pg_next_dst_boundary(m, v19+int32(24), v19+int32(12), v19+int32(4), v19+int32(16), v19+int32(8), v19, l1)
					mBase = m.M
					if v100 < int32(0) {
						v156 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
						v169 = v156
					} else {
						if v100 == int32(0) {
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v105
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v81 - base.I64_extend_i32_s(v107)
							v169 = int32(0) - v107
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							v115 = v81 - base.I64_extend_i32_s(v113)
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
							v120 = v81 - base.I64_extend_i32_s(v118)
							if base.B2i32(v116 <= v115)|base.B2i32(v116 <= v120) == int32(0) {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v125
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v115
								v169 = int32(0) - v113
							} else {
								if base.B2i32(v120 < v116)|base.B2i32(v115 <= v116) == int32(0) {
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v135
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v120
									v169 = int32(0) - v118
								} else {
									if v113 < v118 {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v141
										*(*int64)(unsafe.Add(mBase, uint32(v9))) = v115
										v169 = int32(0) - v113
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v146
										*(*int64)(unsafe.Add(mBase, uint32(v9))) = v120
										v169 = int32(0) - v118
									}
								}
							}
						}
					}
				}
			} else {
				v156 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
				*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
				v169 = v156
			}
		}
	} else {
		if v21 <= int32(_a_F_DetermineTimeZoneOffset_4) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v37 = v31
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = int32(60)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v52 = base.B2i32(int32(2) < v37)
			if int32(2) < v37 {
				v53 = int32(_a_F_DetermineTimeZoneOffset_0)
			} else {
				v53 = int32(_a_F_DetermineTimeZoneOffset_1)
			}
			v54 = v53 + v21
			v62 = base.I32_div_u_s(v54, int32(100))
			v65 = base.I32_div_u_s(v54, int32(400))
			if int32(2) < v37 {
				v69 = int32(1)
			} else {
				v69 = int32(13)
			}
			v74 = base.I32_div_s((v69+v37)*int32(_a_F_DetermineTimeZoneOffset_2), int32(256))
			v77 = v48 + v54*int32(365) + int32(base.Ui32(v54)>>(uint(int32(2))%32)) - v62 + v65 + v74 - int32(_a_F_DetermineTimeZoneOffset_3)
			v81 = base.I64_extend_i32_s(v38+(v39+v40*v41)*v41) + base.I64_extend_i32_s(v77)*int64(86400)
			if base.B2i32(int32(0) < v77)&base.B2i32(v81 < int64(0)) != 0 {
				v156 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
				*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
				v169 = v156
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v81 - int64(86400)
				v100 = F_pg_next_dst_boundary(m, v19+int32(24), v19+int32(12), v19+int32(4), v19+int32(16), v19+int32(8), v19, l1)
				mBase = m.M
				if v100 < int32(0) {
					v156 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
					v169 = v156
				} else {
					if v100 == int32(0) {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v105
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = v81 - base.I64_extend_i32_s(v107)
						v169 = int32(0) - v107
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						v115 = v81 - base.I64_extend_i32_s(v113)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
						v120 = v81 - base.I64_extend_i32_s(v118)
						if base.B2i32(v116 <= v115)|base.B2i32(v116 <= v120) == int32(0) {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v125
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v115
							v169 = int32(0) - v113
						} else {
							if base.B2i32(v120 < v116)|base.B2i32(v115 <= v116) == int32(0) {
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v135
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v120
								v169 = int32(0) - v118
							} else {
								if v113 < v118 {
									v141 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v141
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v115
									v169 = int32(0) - v113
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v146
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v120
									v169 = int32(0) - v118
								}
							}
						}
					}
				}
			}
		} else {
			if v21 != int32(_a_F_DetermineTimeZoneOffset_5) {
				v156 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
				*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
				v169 = v156
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if int32(5) < v34 {
					v156 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
					v169 = v156
				} else {
					v37 = v34
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v41 = int32(60)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v52 = base.B2i32(int32(2) < v37)
					if int32(2) < v37 {
						v53 = int32(_a_F_DetermineTimeZoneOffset_0)
					} else {
						v53 = int32(_a_F_DetermineTimeZoneOffset_1)
					}
					v54 = v53 + v21
					v62 = base.I32_div_u_s(v54, int32(100))
					v65 = base.I32_div_u_s(v54, int32(400))
					if int32(2) < v37 {
						v69 = int32(1)
					} else {
						v69 = int32(13)
					}
					v74 = base.I32_div_s((v69+v37)*int32(_a_F_DetermineTimeZoneOffset_2), int32(256))
					v77 = v48 + v54*int32(365) + int32(base.Ui32(v54)>>(uint(int32(2))%32)) - v62 + v65 + v74 - int32(_a_F_DetermineTimeZoneOffset_3)
					v81 = base.I64_extend_i32_s(v38+(v39+v40*v41)*v41) + base.I64_extend_i32_s(v77)*int64(86400)
					if base.B2i32(int32(0) < v77)&base.B2i32(v81 < int64(0)) != 0 {
						v156 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
						v169 = v156
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v81 - int64(86400)
						v100 = F_pg_next_dst_boundary(m, v19+int32(24), v19+int32(12), v19+int32(4), v19+int32(16), v19+int32(8), v19, l1)
						mBase = m.M
						if v100 < int32(0) {
							v156 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
							v169 = v156
						} else {
							if v100 == int32(0) {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v105
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v81 - base.I64_extend_i32_s(v107)
								v169 = int32(0) - v107
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
								v115 = v81 - base.I64_extend_i32_s(v113)
								v116 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
								v120 = v81 - base.I64_extend_i32_s(v118)
								if base.B2i32(v116 <= v115)|base.B2i32(v116 <= v120) == int32(0) {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v125
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v115
									v169 = int32(0) - v113
								} else {
									if base.B2i32(v120 < v116)|base.B2i32(v115 <= v116) == int32(0) {
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v135
										*(*int64)(unsafe.Add(mBase, uint32(v9))) = v120
										v169 = int32(0) - v118
									} else {
										if v113 < v118 {
											v141 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v141
											*(*int64)(unsafe.Add(mBase, uint32(v9))) = v115
											v169 = int32(0) - v113
										} else {
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v146
											*(*int64)(unsafe.Add(mBase, uint32(v9))) = v120
											v169 = int32(0) - v118
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
	m.G0 = v19 + int32(32)
	m.G0 = v6 + int32(16)
	return v169
}
func F_DoesMultiXactIdConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
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
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1&int32(_a_F_DoesMultiXactIdConflict_0) == int32(_a_F_DoesMultiXactIdConflict_1) {
		v246 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v246 & int32(1)
L2:
	;
	v19 = int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2*v19)+uint32(_c_F_DoesMultiXactIdConflict[0])))
	v33 = F_GetMultiXactIdMembers(m, l0, v13+v19, int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32))|base.B2i32(l1&int32(_a_F_DoesMultiXactIdConflict_2) == int32(64)))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v33 < int32(0) {
		v246 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v33 == int32(0) {
		v233 = v5
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L3
	} else {
		goto L74
	}
L7:
	;
	v43 = int32(0)
	v46 = v5
	goto L8
L8:
	;
	if v46&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v233 = v225
	goto L6
L10:
	;
	v54 = int32(1)
	if l3 == int32(0) {
		v233 = v54
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v59 = int32(3)
	v60 = v43 << (uint(v59) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v62 = v60 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63<<(uint(int32(2))%32))+uint32(_c_F_DoesMultiXactIdConflict[1])))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66*int32(12))+uint32(_c_F_DoesMultiXactIdConflict[0])))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if base.Ui32(v72) < base.Ui32(v59) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v57 != 0 {
		v233 = v54
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v227 = v43 + int32(1)
	if v227 != v33 {
		v43 = v227
		v46 = v225
		goto L8
	} else {
		goto L73
	}
L16:
	;
	if v192 != 0 {
		goto L56
	} else {
		goto L57
	}
L17:
	;
	v192 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_DoesMultiXactIdConflict[2]))
	if v83 == v72 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v192 = int32(1)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_DoesMultiXactIdConflict[3]))
	if v87 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v192 = v184
	goto L16
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_DoesMultiXactIdConflict[4]))
	if v91 == int32(0) {
		v184 = int32(0)
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_DoesMultiXactIdConflict[5]))
	v155 = int32(0)
	v157 = v87 - int32(1)
	goto L46
L27:
	;
	v96 = v91
	goto L28
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v101 == int32(4) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v184 = int32(0)
	goto L23
L30:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v96)+80))
	if v148 != 0 {
		v96 = v148
		goto L28
	} else {
		goto L45
	}
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v104 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v107 = int32(1)
	if v72 == v104 {
		v184 = v107
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	v111 = v109 - int32(1)
	if v111 < int32(0) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v116 = int32(0)
	v118 = v111
	goto L35
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v96)+48))
	v124 = int32(2)
	v125 = base.I32_div_s(v118-v116, v124)
	v126 = v125 + v116
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122+v126<<(uint(v124)%32))))
	if v130 == v72 {
		v184 = v107
		goto L23
	} else {
		goto L37
	}
L36:
	;
	goto L30
L37:
	;
	v134 = F_TransactionIdPrecedes(m, v130, v72)
	mBase = m.M
	if v134 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = v126 + int32(1)
	goto L40
L39:
	;
	v135 = v116
	goto L40
L40:
	;
	if v134 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v138 = v118
	goto L43
L42:
	;
	v138 = v126 - int32(1)
	goto L43
L43:
	;
	if v135 <= v138 {
		v116 = v135
		v118 = v138
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	goto L29
L46:
	;
	v162 = int32(2)
	v163 = base.I32_div_s(v157-v155, v162)
	v164 = v163 + v155
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v153+v164<<(uint(v162)%32))))
	v169 = base.B2i32(v168 == v72)
	if v168 == v72 {
		v184 = v169
		goto L23
	} else {
		goto L48
	}
L47:
	;
	v184 = v169
	goto L23
L48:
	;
	v172 = base.B2i32(base.Ui32(v168) < base.Ui32(v72))
	if base.Ui32(v168) < base.Ui32(v72) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v173 = v164 + int32(1)
	goto L51
L50:
	;
	v173 = v155
	goto L51
L51:
	;
	if base.Ui32(v168) < base.Ui32(v72) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v176 = v157
	goto L54
L53:
	;
	v176 = v164 - int32(1)
	goto L54
L54:
	;
	if v173 <= v176 {
		v155 = v173
		v157 = v176
		goto L46
	} else {
		goto L55
	}
L55:
	;
	goto L47
L56:
	;
	if l3 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v197 = int32(1)
	if v46&v197 != 0 {
		v225 = v197
		goto L15
	} else {
		goto L62
	}
L59:
	;
	v225 = v46
	goto L15
L60:
	;
	goto L61
L61:
	;
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v195)
	v225 = v46
	goto L15
L62:
	;
	v200 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_c_F_DoesMultiXactIdConflict[6])))
	goto L63
L63:
	;
	if int32(base.Ui32(v205)>>(uint(v21)%32))&int32(1) == int32(0) {
		v225 = v200
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211+v60)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v213) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v225 = int32(1)
	goto L15
L66:
	;
	v216 = F_TransactionIdDidAbort(m, v72)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L3
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v220 = F_TransactionIdIsInProgress(m, v72)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	if v216 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v225 = v200
	goto L15
L71:
	;
	if v220 == int32(0) {
		v225 = v200
		goto L15
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	goto L9
L74:
	;
	v246 = v233
	goto L1
}
func F_DropSetting(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v12 = F_table_open(m, int32(2964), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = int32(1)
	F_ScanKeyInit(m, v8, v14, int32(3), int32(184), l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v22 = int32(0)
	v23 = v8
	goto L5
L5:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v22 = v14
	v23 = v8 + int32(48)
	goto L5
L7:
	;
	F_ScanKeyInit(m, v23, int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v31 = v22
	goto L9
L9:
	;
	v32 = F_table_beginscan_catalog(m, v12, v31, v8)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v31 = v22 + int32(1)
	goto L9
L11:
	;
	v34 = F_heap_getnext(m, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = v34
	goto L16
L14:
	;
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+188))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	m.T0[v54].(func(*base.Module, int32))(m, v32)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	F_simple_heap_delete(m, v12, v39+int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v45 = F_heap_getnext(m, v32)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v45 != 0 {
		v39 = v45
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	F_relation_close(m, v12, int32(3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v8 + int32(96)
	return
}
func F_danish_ISO_8859_1_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(1), int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_dasind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v18 int32
	_ = v18
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v61 float64
	_ = v61
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v75 float64
	_ = v75
	var v80 float64
	_ = v80
	var v84 float64
	_ = v84
	var v93 float64
	_ = v93
	var v102 float64
	_ = v102
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v115 float64
	_ = v115
	var v119 float64
	_ = v119
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v144 float64
	_ = v144
	var v155 float64
	_ = v155
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v174 float64
	_ = v174
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v186 float64
	_ = v186
	var v192 float64
	_ = v192
	var v196 float64
	_ = v196
	var v199 float64
	_ = v199
	var v203 float64
	_ = v203
	var v209 float64
	_ = v209
	var v217 int64
	_ = v217
	var v222 int32
	_ = v222
	var v245 float64
	_ = v245
	var v252 float64
	_ = v252
	var v253 float64
	_ = v253
	var v254 float64
	_ = v254
	var v259 float64
	_ = v259
	var v264 float64
	_ = v264
	var v268 float64
	_ = v268
	var v277 float64
	_ = v277
	var v286 float64
	_ = v286
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v310 int64
	_ = v310
	var v315 int32
	_ = v315
	var v328 float64
	_ = v328
	var v339 float64
	_ = v339
	var v351 float64
	_ = v351
	var v352 float64
	_ = v352
	var v353 float64
	_ = v353
	var v358 float64
	_ = v358
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v370 float64
	_ = v370
	var v376 float64
	_ = v376
	var v380 float64
	_ = v380
	var v383 float64
	_ = v383
	var v387 float64
	_ = v387
	var v393 float64
	_ = v393
	var v396 float64
	_ = v396
	var v400 float64
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_abs(v12)
	if base.Ui64(base.I64_reinterpret_f64(v13)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dasind[0])))
		if v18 == int32(0) {
			F_init_degree_constants(m)
			mBase = m.M
		} else {
		}
		if base.F64_gt(v13, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v412 = m.ExcPending
			if v412 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v415 = m.ExcPending
				if v415 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dasind_0), int32(0))
					mBase = m.M
					v419 = m.ExcPending
					if v419 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dasind_1), int32(2164), int32(_a_F_dasind_2))
						mBase = m.M
						v424 = m.ExcPending
						if v424 != 0 {
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
			if base.F64_ge(v12, float64(0)) != 0 {
				if base.F64_le(v12, float64(0.5)) != 0 {
					v33 = base.I64_reinterpret_f64(v12)
					v38 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v38) {
						if base.I32_wrap_i64(v33)|(v38-int32(1072693248)) == int32(0) {
							v115 = base.F64_add(base.F64_mul(v12, float64(1.5707963267948966)), float64(7.52316384526264e-37))
						} else {
							v115 = base.F64_div(float64(0), base.F64_sub(v12, v12))
						}
					} else {
						if base.Ui32(v38) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v38+int32(-1048576)) < base.Ui32(int32(1044381696)) {
								v107 = v12
								v115 = v107
							} else {
								v61 = F_R(m, base.F64_mul(v12, v12))
								mBase = m.M
								v115 = base.F64_add(base.F64_mul(v12, v61), v12)
							}
						} else {
							v68 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v12)), float64(0.5))
							v69 = base.F64_sqrt(v68)
							v70 = F_R(m, v68)
							mBase = m.M
							if base.Ui32(int32(1072640819)) <= base.Ui32(v38) {
								v75 = base.F64_add(base.F64_mul(v69, v70), v69)
								v102 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v75, v75), float64(-6.123233995736766e-17)))
							} else {
								v80 = float64(0.7853981633974483)
								v84 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v69) & int64(-4294967296))
								v93 = base.F64_div(base.F64_sub(v68, base.F64_mul(v84, v84)), base.F64_add(v69, v84))
								v102 = base.F64_add(base.F64_sub(base.F64_sub(v80, base.F64_add(v84, v84)), base.F64_sub(base.F64_mul(base.F64_add(v69, v69), v70), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v93, v93)))), v80)
							}
							if v33 < int64(0) {
								v106 = base.F64_neg(v102)
							} else {
								v106 = v102
							}
							v107 = v106
							v115 = v107
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v115
					v119 = *(*float64)(unsafe.Add(mBase, _c_F_dasind[1]))
					v396 = base.F64_mul(base.F64_div(v115, v119), float64(30))
				} else {
					v126 = base.I64_reinterpret_f64(v12)
					v131 = base.I32_wrap_i64(int64(base.Ui64(v126)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v131) {
						if base.I32_wrap_i64(v126)|(v131-int32(1072693248)) == int32(0) {
							if int64(0) <= v126 {
								v144 = float64(0)
							} else {
								v144 = float64(3.141592653589793)
							}
							v199 = v144
						} else {
							v199 = base.F64_div(float64(0), base.F64_sub(v12, v12))
						}
					} else {
						if base.Ui32(v131) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v131) < base.Ui32(int32(1012924417)) {
								v196 = float64(1.5707963267948966)
								v199 = v196
							} else {
								v155 = F_R(m, base.F64_mul(v12, v12))
								mBase = m.M
								v199 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v12, v155)), v12), float64(1.5707963267948966))
							}
						} else {
							if v126 < int64(0) {
								v167 = base.F64_mul(base.F64_add(v12, float64(1)), float64(0.5))
								v168 = base.F64_sqrt(v167)
								v169 = F_R(m, v167)
								mBase = m.M
								v174 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v168, base.F64_add(base.F64_mul(v168, v169), float64(-6.123233995736766e-17))))
								v199 = base.F64_add(v174, v174)
							} else {
								v179 = base.F64_mul(base.F64_sub(float64(1), v12), float64(0.5))
								v180 = base.F64_sqrt(v179)
								v181 = F_R(m, v179)
								mBase = m.M
								v186 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v180) & int64(-4294967296))
								v192 = base.F64_add(base.F64_add(base.F64_mul(v180, v181), base.F64_div(base.F64_sub(v179, base.F64_mul(v186, v186)), base.F64_add(v180, v186))), v186)
								v196 = base.F64_add(v192, v192)
								v199 = v196
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v199
					v203 = *(*float64)(unsafe.Add(mBase, _c_F_dasind[2]))
					v396 = base.F64_add(base.F64_mul(base.F64_div(v199, v203), float64(-60)), float64(90))
				}
			} else {
				v209 = base.F64_neg(v12)
				if base.F64_ge(v12, float64(-0.5)) != 0 {
					v217 = base.I64_reinterpret_f64(v209)
					v222 = base.I32_wrap_i64(int64(base.Ui64(v217)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v222) {
						if base.I32_wrap_i64(v217)|(v222-int32(1072693248)) == int32(0) {
							v299 = base.F64_add(base.F64_mul(v209, float64(1.5707963267948966)), float64(7.52316384526264e-37))
						} else {
							v299 = base.F64_div(float64(0), base.F64_sub(v209, v209))
						}
					} else {
						if base.Ui32(v222) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v222+int32(-1048576)) < base.Ui32(int32(1044381696)) {
								v291 = v209
								v299 = v291
							} else {
								v245 = F_R(m, base.F64_mul(v209, v209))
								mBase = m.M
								v299 = base.F64_add(base.F64_mul(v209, v245), v209)
							}
						} else {
							v252 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v209)), float64(0.5))
							v253 = base.F64_sqrt(v252)
							v254 = F_R(m, v252)
							mBase = m.M
							if base.Ui32(int32(1072640819)) <= base.Ui32(v222) {
								v259 = base.F64_add(base.F64_mul(v253, v254), v253)
								v286 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v259, v259), float64(-6.123233995736766e-17)))
							} else {
								v264 = float64(0.7853981633974483)
								v268 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v253) & int64(-4294967296))
								v277 = base.F64_div(base.F64_sub(v252, base.F64_mul(v268, v268)), base.F64_add(v253, v268))
								v286 = base.F64_add(base.F64_sub(base.F64_sub(v264, base.F64_add(v268, v268)), base.F64_sub(base.F64_mul(base.F64_add(v253, v253), v254), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v277, v277)))), v264)
							}
							if v217 < int64(0) {
								v290 = base.F64_neg(v286)
							} else {
								v290 = v286
							}
							v291 = v290
							v299 = v291
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v299
					v303 = *(*float64)(unsafe.Add(mBase, _c_F_dasind[1]))
					v393 = base.F64_mul(base.F64_div(v299, v303), float64(30))
				} else {
					v310 = base.I64_reinterpret_f64(v209)
					v315 = base.I32_wrap_i64(int64(base.Ui64(v310)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v315) {
						if base.I32_wrap_i64(v310)|(v315-int32(1072693248)) == int32(0) {
							if int64(0) <= v310 {
								v328 = float64(0)
							} else {
								v328 = float64(3.141592653589793)
							}
							v383 = v328
						} else {
							v383 = base.F64_div(float64(0), base.F64_sub(v209, v209))
						}
					} else {
						if base.Ui32(v315) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v315) < base.Ui32(int32(1012924417)) {
								v380 = float64(1.5707963267948966)
								v383 = v380
							} else {
								v339 = F_R(m, base.F64_mul(v209, v209))
								mBase = m.M
								v383 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v209, v339)), v209), float64(1.5707963267948966))
							}
						} else {
							if v310 < int64(0) {
								v351 = base.F64_mul(base.F64_add(v209, float64(1)), float64(0.5))
								v352 = base.F64_sqrt(v351)
								v353 = F_R(m, v351)
								mBase = m.M
								v358 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v352, base.F64_add(base.F64_mul(v352, v353), float64(-6.123233995736766e-17))))
								v383 = base.F64_add(v358, v358)
							} else {
								v363 = base.F64_mul(base.F64_sub(float64(1), v209), float64(0.5))
								v364 = base.F64_sqrt(v363)
								v365 = F_R(m, v363)
								mBase = m.M
								v370 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v364) & int64(-4294967296))
								v376 = base.F64_add(base.F64_add(base.F64_mul(v364, v365), base.F64_div(base.F64_sub(v363, base.F64_mul(v370, v370)), base.F64_add(v364, v370))), v370)
								v380 = base.F64_add(v376, v376)
								v383 = v380
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v383
					v387 = *(*float64)(unsafe.Add(mBase, _c_F_dasind[2]))
					v393 = base.F64_add(base.F64_mul(base.F64_div(v383, v387), float64(-60)), float64(90))
				}
				v396 = base.F64_neg(v393)
			}
			if base.F64_eq(base.F64_abs(v396), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v426 = m.ExcPending
				if v426 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v400 = v396
				v401 = F_Float8GetDatum(m, v400)
				mBase = m.M
				v404 = m.ExcPending
				if v404 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v401
				}
			}
		}
	} else {
		v400 = math.Float64frombits(uint64(0x7ff8000000000000))
		v401 = F_Float8GetDatum(m, v400)
		mBase = m.M
		v404 = m.ExcPending
		if v404 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v401
		}
	}
}
func F_date2isoweek(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
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
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	v9 = base.B2i32(int32(2) < l1)
	if int32(2) < l1 {
		v10 = int32(_a_F_date2isoweek_0)
	} else {
		v10 = int32(_a_F_date2isoweek_1)
	}
	v11 = v10 + l0
	v16 = base.I32_div_s(v11, int32(4))
	v19 = base.I32_div_s(v11, int32(-100))
	v22 = base.I32_div_s(v11, int32(400))
	if int32(2) < l1 {
		v26 = int32(1)
	} else {
		v26 = int32(13)
	}
	v31 = base.I32_div_s((v26+l1)*int32(_a_F_date2isoweek_2), int32(256))
	v34 = l2 + v11*int32(365) + v16 + v19 + v22 + v31 - int32(_a_F_date2isoweek_3)
	v43 = int32(_a_F_date2isoweek_1) + l0
	v48 = base.I32_div_s(v43, int32(4))
	v51 = base.I32_div_s(v43, int32(-100))
	v54 = base.I32_div_s(v43, int32(400))
	v63 = base.I32_div_s(int32(_a_F_date2isoweek_4), int32(256))
	v66 = int32(4) + v43*int32(365) + v48 + v51 + v54 + v63 - int32(_a_F_date2isoweek_3)
	v67 = int32(1)
	v71 = int32(7)
	v72 = base.I32_rem_s(v66-v67+v67, v71)
	if v72 < int32(0) {
		v77 = v72 + v71
	} else {
		v77 = v72
	}
	if v34 < v66-v77 {
		v90 = int32(_a_F_date2isoweek_1) + (l0 - int32(1))
		v95 = base.I32_div_s(v90, int32(4))
		v98 = base.I32_div_s(v90, int32(-100))
		v101 = base.I32_div_s(v90, int32(400))
		v110 = base.I32_div_s(int32(_a_F_date2isoweek_4), int32(256))
		v113 = int32(4) + v90*int32(365) + v95 + v98 + v101 + v110 - int32(_a_F_date2isoweek_3)
		v114 = int32(1)
		v118 = int32(7)
		v119 = base.I32_rem_s(v113-v114+v114, v118)
		if v119 < int32(0) {
			v124 = v119 + v118
		} else {
			v124 = v119
		}
		v125 = v113
		v126 = v124
	} else {
		v125 = v66
		v126 = v77
	}
	v128 = v126 - v125 + v34
	if int32(357) <= v128 {
		v141 = l0 + int32(_a_F_date2isoweek_0)
		v146 = base.I32_div_s(v141, int32(4))
		v149 = base.I32_div_s(v141, int32(-100))
		v152 = base.I32_div_s(v141, int32(400))
		v161 = base.I32_div_s(int32(_a_F_date2isoweek_4), int32(256))
		v164 = int32(4) + v141*int32(365) + v146 + v149 + v152 + v161 - int32(_a_F_date2isoweek_3)
		v165 = int32(1)
		v169 = int32(7)
		v170 = base.I32_rem_s(v164-v165+v165, v169)
		if v170 < int32(0) {
			v175 = v170 + v169
		} else {
			v175 = v170
		}
		v176 = v164 - v175
		if v34 < v176 {
			v179 = v128
		} else {
			v179 = v34 - v176
		}
		v181 = v179
	} else {
		v181 = v128
	}
	v183 = base.I32_div_s(v181, int32(7))
	return v183 + int32(1)
}
func F_datetime_to_char_body(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v8 = F_text_to_cstring(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_strlen(m, v8)
		mBase = m.M
		v14 = v12 * int32(12)
		v17 = F_palloc(m, v14|int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v19)
			if base.Ui32(int32(156)) <= base.Ui32(v12) {
				v25 = F_palloc(m, v14+int32(12))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_parse_format(m, v25, v8, int32(_a_F_datetime_to_char_body_0), int32(_a_F_datetime_to_char_body_1), int32(_a_F_datetime_to_char_body_2), int32(1), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_DCH_to_char(m, v25, l2, l0, v17, l3)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v25)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v8)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = F_cstring_to_text(m, v17)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v17)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											return v46
										}
									}
								}
							}
						}
					}
				}
			} else {
				v39 = F_DCH_cache_fetch(m, v8, int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_DCH_to_char(m, v39, l2, l0, v17, l3)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v8)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = F_cstring_to_text(m, v17)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v17)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									return v46
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_db_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_dceil(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_ceil(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_dcos(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 float64
	_ = v35
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _c_F_dcos[0])) = int32(0)
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dcos_0), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dcos_1), int32(1898), int32(_a_F_dcos_2))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
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
			v19 = m.G0
			v21 = v19 - int32(16)
			m.G0 = v21
			v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v4))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v28) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v28) < base.Ui32(int32(1044816030)) {
					v57 = float64(1)
				} else {
					v35 = F___cos(m, v4, float64(0))
					mBase = m.M
					v57 = v35
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v28) {
					v57 = base.F64_sub(v4, v4)
				} else {
					v39 = F___rem_pio2(m, v4, v21)
					mBase = m.M
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
					switch v39&int32(3) - int32(1) {
					case 0:
						v48 = F___sin(m, v41, v40, int32(1))
						mBase = m.M
						v57 = base.F64_neg(v48)
					case 1:
						v50 = F___cos(m, v41, v40)
						mBase = m.M
						v57 = base.F64_neg(v50)
					case 2:
						v53 = F___sin(m, v41, v40, int32(1))
						mBase = m.M
						v57 = v53
					default:
						v46 = F___cos(m, v41, v40)
						mBase = m.M
						v57 = v46
					}
				}
			}
			m.G0 = v21 + int32(16)
			v62 = v57
			v63 = F_Float8GetDatum(m, v62)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				return v63
			}
		}
	} else {
		v62 = math.Float64frombits(uint64(0x7ff8000000000000))
		v63 = F_Float8GetDatum(m, v62)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			return v63
		}
	}
}
func F_dec_lex_level(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v3&int32(4) == int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = int32(0)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v28 - int32(1)
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(2))%32))))
		if v14 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = int32(0)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v28 - int32(1)
			return
		} else {
			F_pfree(m, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v28 - int32(1)
				return
			}
		}
	}
}
func F_decompose_code(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v356 int32
	_ = v356
	v10 = l0 - int32(_a_F_decompose_code_0)
	if base.Ui32(v10) <= base.Ui32(int32(_a_F_decompose_code_1)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v15 = int32(2)
		v18 = int32(_a_F_decompose_code_2)
		v19 = v10 & v18
		v20 = int32(588)
		v21 = base.I32_div_u_s(v19, v20)
		*(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(v15)%32)))) = v21 | int32(_a_F_decompose_code_3)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v26 = int32(1)
		v27 = v25 + v26
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v27
		v37 = int32(28)
		v38 = base.I32_div_u_s((v10-v21*v20)&v18, v37)
		*(*int32)(unsafe.Add(mBase, uint32(v13+v27<<(uint(v15)%32)))) = v38 + int32(_a_F_decompose_code_4)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v44 = v42 + v26
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v44
		v47 = base.I32_rem_u_s(v19, v37)
		if v47 == int32(0) {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13+v44<<(uint(int32(2))%32)))) = v47 + int32(_a_F_decompose_code_5)
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
			return
		}
	} else {
		v56 = int32(16711935)
		v58 = int32(8)
		v59 = base.I32_rotr(l0&v56, v58)
		v60 = int32(24)
		v61 = base.I32_rotr(l0, v60)
		v63 = int32(255)
		v64 = (v59 | v61) & v63
		v65 = int32(_a_F_decompose_code_6)
		v70 = int32(base.Ui32(v59)>>(uint(v58)%32)) & v63
		v77 = int32(base.Ui32(v61&v56) >> (uint(int32(16)) % 32))
		v82 = int32(base.Ui32(v59) >> (uint(v60) % 32))
		v86 = int32(_a_F_decompose_code_7)
		v87 = base.I32_rem_u_s(((v64*v65+v70)*v65+v77)*v65+v82+int32(402620417), v86)
		v88 = int32(1)
		v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87<<(uint(v88)%32))+uint32(_c_F_decompose_code[0]))))
		v91 = int32(257)
		v101 = base.I32_rem_u_s(((v64*v91+v70)*v91+v77)*v91+v82, v86)
		v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101<<(uint(v88)%32))+uint32(_c_F_decompose_code[0]))))
		v105 = v90 + v104
		if base.Ui32(int32(_a_F_decompose_code_8)) < base.Ui32(v105) {
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
			return
		} else {
			v109 = v105 << (uint(int32(3)) % 32)
			v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_decompose_code[1])))
			if l0 != v110 {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
				v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
				return
			} else {
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_decompose_code[2]))))
				v116 = v114 & int32(31)
				if v116 == int32(0) {
					v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
					v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
					return
				} else {
					if l1|base.B2i32(v114&int32(32) == int32(0)) != 0 {
						v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_decompose_code[3]))))
						if v114&int32(64) != 0 {
							v147 = int32(_a_F_decompose_code_9)
							*(*int32)(unsafe.Add(mBase, _c_F_decompose_code[4])) = v144
							v155 = int32(1)
							v156 = v147
						} else {
							v155 = v116
							v156 = v144<<(uint(int32(2))%32) + int32(_a_F_decompose_code_10)
						}
						v158 = int32(0)
						for {
							v169 = *(*int32)(unsafe.Add(mBase, uint32(v156+v158<<(uint(int32(2))%32))))
							v175 = v169 - int32(_a_F_decompose_code_0)
							if base.Ui32(v175) <= base.Ui32(int32(_a_F_decompose_code_1)) {
								v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v180 = int32(2)
								v183 = int32(_a_F_decompose_code_2)
								v184 = v175 & v183
								v185 = int32(588)
								v186 = base.I32_div_u_s(v184, v185)
								*(*int32)(unsafe.Add(mBase, uint32(v178+v179<<(uint(v180)%32)))) = v186 | int32(_a_F_decompose_code_3)
								v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v191 = int32(1)
								v192 = v190 + v191
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v192
								v202 = int32(28)
								v203 = base.I32_div_u_s((v175-v186*v185)&v183, v202)
								*(*int32)(unsafe.Add(mBase, uint32(v178+v192<<(uint(v180)%32)))) = v203 + int32(_a_F_decompose_code_4)
								v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v209 = v207 + v191
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v209
								v212 = base.I32_rem_u_s(v184, v202)
								if v212 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v178+v209<<(uint(int32(2))%32)))) = v212 + int32(_a_F_decompose_code_5)
									v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
								}
							} else {
								v221 = int32(16711935)
								v223 = int32(8)
								v224 = base.I32_rotr(v169&v221, v223)
								v225 = int32(24)
								v226 = base.I32_rotr(v169, v225)
								v228 = int32(255)
								v229 = (v224 | v226) & v228
								v230 = int32(_a_F_decompose_code_6)
								v235 = int32(base.Ui32(v224)>>(uint(v223)%32)) & v228
								v242 = int32(base.Ui32(v226&v221) >> (uint(int32(16)) % 32))
								v247 = int32(base.Ui32(v224) >> (uint(v225) % 32))
								v251 = int32(_a_F_decompose_code_7)
								v252 = base.I32_rem_u_s(((v229*v230+v235)*v230+v242)*v230+v247+int32(402620417), v251)
								v253 = int32(1)
								v255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252<<(uint(v253)%32))+uint32(_c_F_decompose_code[0]))))
								v256 = int32(257)
								v266 = base.I32_rem_u_s(((v229*v256+v235)*v256+v242)*v256+v247, v251)
								v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v266<<(uint(v253)%32))+uint32(_c_F_decompose_code[0]))))
								v270 = v255 + v269
								if base.Ui32(int32(_a_F_decompose_code_8)) < base.Ui32(v270) {
									v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
									v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
								} else {
									v274 = v270 << (uint(int32(3)) % 32)
									v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+uint32(_c_F_decompose_code[1])))
									if v169 != v275 {
										v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
										v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
									} else {
										v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+uint32(_c_F_decompose_code[2]))))
										v281 = v279 & int32(31)
										if v281 == int32(0) {
											v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
											v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
										} else {
											if l1|base.B2i32(v279&int32(32) == int32(0)) != 0 {
												v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274)+uint32(_c_F_decompose_code[3]))))
												if v279&int32(64) != 0 {
													v312 = int32(_a_F_decompose_code_9)
													*(*int32)(unsafe.Add(mBase, _c_F_decompose_code[4])) = v309
													v320 = int32(1)
													v321 = v312
												} else {
													v320 = v281
													v321 = v309<<(uint(int32(2))%32) + int32(_a_F_decompose_code_10)
												}
												v323 = int32(0)
												for {
													v334 = *(*int32)(unsafe.Add(mBase, uint32(v321+v323<<(uint(int32(2))%32))))
													F_decompose_code(m, v334, l1, l2, l3)
													mBase = m.M
													v337 = v323 + int32(1)
													if v337 != v320 {
														v323 = v337
														continue
													} else {
														break
													}
													break
												}
											} else {
												v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
												v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
											}
										}
									}
								}
							}
							v356 = v158 + int32(1)
							if v356 != v155 {
								v158 = v356
								continue
							} else {
								break
							}
							break
						}
						return
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
						return
					}
				}
			}
		}
	}
}
func F_defined(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_defined(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_degrees(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v7 = base.F64_div(v5, float64(0.017453292519943295))
	v9 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v7), v9)&base.F64_ne(base.F64_abs(v5), v9) == int32(0) {
		v17 = float64(0)
		if base.F64_eq(v7, v17)&base.F64_ne(v5, v17) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v22 = F_Float8GetDatum(m, v7)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_delete(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_delete(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_deserialize_deflist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = F_palloc(m, v51+int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L14
	}
L2:
	;
	return int32(0)
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v28 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v39 = int32(1)
	if v22&v39 != 0 {
		v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v31 = int32(16)
	goto L9
L8:
	;
	v31 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(4)
	goto L12
L11:
	;
	v38 = v31
	goto L12
L12:
	;
	v51 = v38
	goto L1
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	if v51 <= int32(0) {
		v317 = v2
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_pfree(m, v54)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L2
	} else {
		goto L101
	}
L16:
	;
	v58 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v60&v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = v58
	goto L19
L18:
	;
	v63 = int32(4)
	goto L19
L19:
	;
	v64 = v18 + v63
	v65 = v51 + v64
	v68 = v64
	v70 = v2
	v71 = v2
	v73 = v2
	v75 = v2
	goto L20
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	switch v70 - int32(1) {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	case 5:
		goto L31
	case 6:
		goto L30
	default:
		goto L37
	}
L21:
	;
	switch v279 {
	case 0:
		v317 = v282
		goto L15
	default:
		goto L93
	case 7:
		goto L92
	}
L22:
	;
	v284 = v278 + int32(1)
	if base.Ui32(v284) < base.Ui32(v65) {
		v68 = v284
		v70 = v279
		v71 = v280
		v73 = v281
		v75 = v282
		goto L20
	} else {
		goto L91
	}
L23:
	;
	v278 = v273
	v279 = v274
	v280 = v275
	v281 = v73
	v282 = v277
	goto L22
L24:
	;
	v273 = v68
	v274 = int32(5)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L25:
	;
	v273 = v68
	v274 = v268
	v275 = v71
	v277 = v75
	goto L23
L26:
	;
	v268 = int32(4)
	goto L25
L27:
	;
	v268 = int32(3)
	goto L25
L28:
	;
	v278 = v259
	v279 = v260
	v280 = v71
	v281 = v71
	v282 = v75
	goto L22
L29:
	;
	v259 = v68
	v260 = int32(6)
	goto L28
L30:
	;
	v246 = v71 + int32(1)
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23, 35:
		goto L88
	default:
		goto L87
	}
L31:
	;
	if v81 == int32(34) {
		goto L76
	} else {
		goto L77
	}
L32:
	;
	if v81 != int32(92) {
		goto L61
	} else {
		goto L62
	}
L33:
	;
	v153 = int32(5)
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		goto L26
	default:
		goto L56
	case 25:
		goto L29
	case 30:
		v278 = v68
		v279 = v153
		v280 = v71
		v281 = v71
		v282 = v75
		goto L22
	case 60:
		goto L57
	}
L34:
	;
	if base.B2i32(v81 == int32(32))|base.B2i32(base.Ui32(v81-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L27
	} else {
		goto L49
	}
L35:
	;
	if v81 == int32(34) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v91 = v71 + int32(1)
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		goto L42
	default:
		goto L40
	case 52:
		goto L41
	}
L37:
	;
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23, 35:
		v278 = v68
		v279 = int32(0)
		v280 = v71
		v281 = v73
		v282 = v75
		goto L22
	default:
		goto L38
	case 25:
		goto L39
	}
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v81)
	v273 = v68
	v274 = int32(1)
	v275 = v54 + int32(1)
	v277 = v75
	goto L23
L39:
	;
	v273 = v68
	v274 = int32(2)
	v275 = v54
	v277 = v75
	goto L23
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v273 = v68
	v274 = int32(1)
	v275 = v91
	v277 = v75
	goto L23
L41:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v97)
	v273 = v68
	v274 = int32(4)
	v275 = v91
	v277 = v75
	goto L23
L42:
	;
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v94)
	v273 = v68
	v274 = int32(3)
	v275 = v91
	v277 = v75
	goto L23
L43:
	;
	v105 = v68 + int32(1)
	if base.Ui32(v65) <= base.Ui32(v105) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v273 = v68
	v274 = int32(2)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L46:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v115)
	v273 = v68
	v274 = int32(3)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L47:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v107 != int32(34) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v110 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v110)
	v273 = v105
	v274 = int32(2)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L49:
	;
	if v81 == int32(61) {
		goto L26
	} else {
		goto L50
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v140 = F_text_to_cstring(m, v18)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v140
	F_errmsg(m, int32(_a_F_deserialize_deflist_0), v16+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_deserialize_deflist_1), int32(1708), int32(_a_F_deserialize_deflist_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v278 = v68
	v279 = int32(7)
	v280 = v71 + int32(1)
	v281 = v71
	v282 = v75
	goto L22
L57:
	;
	v157 = v68 + int32(1)
	if base.Ui32(v65) <= base.Ui32(v157) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v159 != int32(39) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v259 = v157
	v260 = v153
	goto L28
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	goto L24
L61:
	;
	if v81 != int32(39) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v199 = v68 + int32(1)
	if base.Ui32(v65) <= base.Ui32(v199) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	v172 = v68 + int32(1)
	if base.Ui32(v65) <= base.Ui32(v172) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v182)
	v187 = F_pstrdup(m, v54)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L68
	}
L66:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v174 != int32(39) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v177 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v177)
	v273 = v172
	v274 = int32(5)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L68:
	;
	v189 = F_pstrdup(m, v73)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v191 = F_makeString(m, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v194 = F_makeDefElem(m, v187, v191, int32(-1))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v196 = F_lappend(m, v75, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v273 = v68
	v274 = v182
	v275 = v71 + int32(1)
	v277 = v196
	goto L23
L73:
	;
	v209 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v209)
	goto L24
L74:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v201 != int32(92) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v204 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v204)
	v273 = v199
	v274 = int32(5)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L76:
	;
	v215 = v68 + int32(1)
	if base.Ui32(v65) <= base.Ui32(v215) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v273 = v68
	v274 = int32(6)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L79:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v225)
	v230 = F_pstrdup(m, v54)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v217 != int32(34) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v220 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v220)
	v273 = v215
	v274 = int32(6)
	v275 = v71 + int32(1)
	v277 = v75
	goto L23
L82:
	;
	v232 = F_pstrdup(m, v73)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v234 = F_makeString(m, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v237 = F_makeDefElem(m, v230, v234, int32(-1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v239 = F_lappend(m, v75, v237)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v273 = v68
	v274 = v225
	v275 = v71 + int32(1)
	v277 = v239
	goto L23
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v273 = v68
	v274 = int32(7)
	v275 = v246
	v277 = v75
	goto L23
L88:
	;
	v249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v249)
	v252 = F_buildDefItem(m, v54, v73)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v254 = F_lappend(m, v75, v252)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v273 = v68
	v274 = v249
	v275 = v246
	v277 = v254
	goto L23
L91:
	;
	goto L21
L92:
	;
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v304)
	v306 = F_buildDefItem(m, v54, v281)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L2
	} else {
		goto L99
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v293 = F_text_to_cstring(m, v18)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v293
	F_errmsg(m, int32(_a_F_deserialize_deflist_0), v16)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_deserialize_deflist_1), int32(1823), int32(_a_F_deserialize_deflist_2))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v308 = F_lappend(m, v282, v306)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	v317 = v308
	goto L15
L101:
	;
	m.G0 = v16 + int32(32)
	return v317
}
func F_detoast_attr_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
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
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v306 int32
	_ = v306
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	if v4 <= l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = l0
	v21 = l2
	goto L7
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L36
	} else {
		goto L152
	}
L4:
	;
	m.G0 = v15 - int32(-64)
	return v464
L5:
	;
	v421 = int32(1)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	if v423&v421 != 0 {
		goto L134
	} else {
		goto L135
	}
L6:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v101&int32(3) != int32(2) {
		goto L42
	} else {
		goto L43
	}
L7:
	;
	v32 = base.B2i32(v21 < int32(0))
	if v21 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v48&int32(254) != int32(2) {
		v100 = v19
		goto L6
	} else {
		goto L40
	}
L9:
	;
	v41 = v21
	v44 = int32(-1)
	goto L11
L10:
	;
	v35 = l1 + v21
	v37 = v32 ^ base.B2i32(v35 < l1)
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v45 != int32(1) {
		v100 = v19
		goto L6
	} else {
		goto L18
	}
L12:
	;
	v38 = int32(-1)
	goto L14
L13:
	;
	v38 = v21
	goto L14
L14:
	;
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = int32(-1)
	goto L17
L16:
	;
	v40 = v35
	goto L17
L17:
	;
	v41 = v38
	v44 = v40
	goto L11
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v48 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L8
L20:
	;
	if v48 != int32(18) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v19)+2))
	v19 = v89
	v21 = v41
	goto L7
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+6))
	v55 = v53 & int32(1073741823)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+2))
	if base.Ui32(v55) < base.Ui32(v56-int32(4)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if int32(0) <= v44 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v87 = F_toast_fetch_datum_slice(m, v19, l1, v41)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L36
	} else {
		goto L39
	}
L27:
	;
	if base.Ui32(v53) <= base.Ui32(int32(1073741823)) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v85 = F_toast_fetch_datum(m, v19)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L36
	} else {
		goto L38
	}
L30:
	;
	v71 = base.I64_div_s(base.I64_extend_i32_s(v44)*int64(9)+int64(7), int64(8))
	v73 = v71 + int64(2)
	v74 = base.I64_extend_i32_s(v55)
	if v73 < v74 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v80 = v55
	goto L32
L32:
	;
	v81 = F_toast_fetch_datum_slice(m, v19, int32(0), v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v76 = v73
	goto L35
L34:
	;
	v76 = v74
	goto L35
L35:
	;
	v80 = base.I32_wrap_i64(v76)
	goto L32
L36:
	;
	return int32(0)
L37:
	;
	v100 = v81
	goto L6
L38:
	;
	v100 = v85
	goto L6
L39:
	;
	v464 = v87
	goto L4
L40:
	;
	v94 = F_detoast_external_attr(m, v19)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v100 = v94
	goto L6
L42:
	;
	v419 = v100
	goto L5
L43:
	;
	goto L44
L44:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if int32(0) <= v44 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v19 == v100 {
		v419 = v415
		goto L5
	} else {
		goto L131
	}
L46:
	;
	v411 = F_pglz_decompress_datum(m, v100)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L36
	} else {
		goto L130
	}
L47:
	;
	v110 = int32(base.Ui32(v106) >> (uint(int32(30)) % 32))
	if base.Ui32(v106&int32(1073741823)) <= base.Ui32(v44) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v392 = int32(base.Ui32(v106) >> (uint(int32(30)) % 32))
	switch v392 {
	case 0:
		goto L46
	case 1:
		goto L125
	default:
		goto L124
	}
L50:
	;
	switch v110 {
	case 0:
		goto L46
	case 1:
		goto L54
	default:
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	switch v110 {
	case 0:
		goto L61
	case 1:
		goto L60
	default:
		goto L59
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L36
	} else {
		goto L56
	}
L54:
	;
	v114 = F_lz4_decompress_datum(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L36
	} else {
		goto L55
	}
L55:
	;
	v415 = v114
	goto L45
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v110
	F_errmsg_internal(m, int32(_a_F_detoast_attr_slice_0), v13+int32(-32))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L36
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_detoast_attr_slice_1), int32(489), int32(_a_F_detoast_attr_slice_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L36
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L36
	} else {
		goto L121
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L36
	} else {
		goto L116
	}
L61:
	;
	v133 = F_palloc(m, v44+int32(4))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L36
	} else {
		goto L62
	}
L62:
	;
	v135 = int32(8)
	v136 = v100 + v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v141 = int32(base.Ui32(v137)>>(uint(int32(2))%32)) - v135
	v143 = v133 + int32(4)
	v144 = int32(0)
	v152 = v143 + v44
	v153 = v136 + v141
	if base.B2i32(v141 <= v144)|base.B2i32(v44 <= v144) == v144 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if v328 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L64:
	;
	goto L63
L65:
	;
	goto L106
L66:
	;
	v161 = v136
	v164 = v143
	goto L69
L67:
	;
	goto L68
L68:
	;
	v306 = v143
	goto L65
L69:
	;
	v174 = v161 + int32(1)
	if base.Ui32(v153) <= base.Ui32(v174) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v306 = v292
	goto L65
L71:
	;
	if base.Ui32(v153) <= base.Ui32(v289) {
		v306 = v292
		goto L65
	} else {
		goto L103
	}
L72:
	;
	v289 = v174
	v292 = v164
	goto L71
L73:
	;
	goto L74
L74:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v179 = v174
	v181 = v164
	v187 = v176
	v188 = int32(0)
	goto L75
L75:
	;
	if v187&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v289 = v265
	v292 = v277
	goto L71
L77:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v188))|base.B2i32(base.Ui32(v153) <= base.Ui32(v265)) != 0 {
		v289 = v265
		v292 = v277
		goto L71
	} else {
		goto L101
	}
L78:
	;
	v192 = int32(-1)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v197 = v193&int32(15) + int32(3)
	if v197 != int32(18) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v259)
	v261 = int32(1)
	v265 = v179 + v261
	v277 = v181 + v261
	goto L77
L81:
	;
	v207 = v197
	v208 = v179 + int32(2)
	goto L83
L82:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+2)))
	v207 = v202 + int32(18)
	v208 = v179 + int32(3)
	goto L83
L83:
	;
	if base.Ui32(v153) < base.Ui32(v208) {
		v328 = v192
		goto L64
	} else {
		goto L84
	}
L84:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v215 = v210 | v193<<(uint(int32(4))%32)&int32(3840)
	if base.B2i32(v215 == int32(0))|base.B2i32(v181-v143 < v215) != 0 {
		v328 = v192
		goto L64
	} else {
		goto L85
	}
L85:
	;
	v221 = v152 - v181
	if v207 < v221 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v223 = v207
	goto L88
L87:
	;
	v223 = v221
	goto L88
L88:
	;
	if v215 < v223 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v226 = v215
	v228 = v181
	v230 = v223
	goto L92
L90:
	;
	v245 = v215
	v247 = v181
	v249 = v223
	goto L91
L91:
	;
	if v249 != 0 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	if v226 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v245 = v242
	v247 = v239
	v249 = v240
	goto L91
L94:
	;
	base.MemoryCopy(m, v228, v228-v226, v226)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v239 = v226 + v228
	v240 = v230 - v226
	v242 = v226 << (uint(int32(1)) % 32)
	if v242 < v240 {
		v226 = v242
		v228 = v239
		v230 = v240
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L93
L98:
	;
	base.MemoryCopy(m, v247, v247-v245, v249)
	goto L100
L99:
	;
	goto L100
L100:
	;
	v265 = v208
	v277 = v247 + v249
	goto L77
L101:
	;
	v282 = int32(1)
	if base.Ui32(v277) < base.Ui32(v152) {
		v179 = v265
		v181 = v277
		v187 = int32(base.Ui32(v187&int32(254)) >> (uint(v282) % 32))
		v188 = v188 + v282
		goto L75
	} else {
		goto L102
	}
L102:
	;
	goto L76
L103:
	;
	if base.Ui32(v292) < base.Ui32(v152) {
		v161 = v289
		v164 = v292
		goto L69
	} else {
		goto L104
	}
L104:
	;
	goto L70
L106:
	;
	goto L107
L107:
	;
	v328 = v306 - v143
	goto L64
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L36
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v328<<(uint(int32(2))%32) + int32(16)
	v415 = v133
	goto L45
L112:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L36
	} else {
		goto L113
	}
L113:
	;
	F_errmsg_internal(m, int32(_a_F_detoast_attr_slice_3), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L36
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_detoast_attr_slice_4), int32(126), int32(_a_F_detoast_attr_slice_5))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L36
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L36
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_detoast_attr_slice_6), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L36
	} else {
		goto L118
	}
L118:
	;
	F_errdetail(m, int32(_a_F_detoast_attr_slice_7), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L36
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_detoast_attr_slice_4), int32(218), int32(_a_F_detoast_attr_slice_8))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L36
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v110
	F_errmsg_internal(m, int32(_a_F_detoast_attr_slice_0), v13+int32(-48))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L36
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_detoast_attr_slice_1), int32(532), int32(_a_F_detoast_attr_slice_9))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L36
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L36
	} else {
		goto L127
	}
L125:
	;
	v393 = F_lz4_decompress_datum(m)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L36
	} else {
		goto L126
	}
L126:
	;
	v415 = v393
	goto L45
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v392
	F_errmsg_internal(m, int32(_a_F_detoast_attr_slice_0), v13+int32(-16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L36
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_detoast_attr_slice_1), int32(489), int32(_a_F_detoast_attr_slice_2))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L36
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v415 = v411
	goto L45
L131:
	;
	F_pfree(m, v100)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L36
	} else {
		goto L132
	}
L132:
	;
	v419 = v415
	goto L5
L133:
	;
	if l1 < v437 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v426 = int32(1)
	v436 = v421
	v437 = int32(base.Ui32(v423)>>(uint(v426)%32)) - v426
	goto L133
L135:
	;
	goto L136
L136:
	;
	v430 = int32(4)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v436 = v430
	v437 = int32(base.Ui32(v431)>>(uint(int32(2))%32)) - v430
	goto L133
L137:
	;
	v439 = v437 - l1
	if v437 < v44 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v446 = int32(0)
	v447 = v4
	goto L139
L139:
	;
	v449 = v446 + int32(4)
	v450 = F_palloc(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L36
	} else {
		goto L146
	}
L140:
	;
	v441 = v439
	goto L142
L141:
	;
	v441 = v41
	goto L142
L142:
	;
	if v41 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v444 = v439
	goto L145
L144:
	;
	v444 = v441
	goto L145
L145:
	;
	v446 = v444
	v447 = l1
	goto L139
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = v449 << (uint(int32(2)) % 32)
	if v446 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	base.MemoryCopy(m, v450+int32(4), v419+v436+v447, v446)
	goto L149
L148:
	;
	goto L149
L149:
	;
	if v19 == v419 {
		v464 = v450
		goto L4
	} else {
		goto L150
	}
L150:
	;
	F_pfree(m, v419)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L36
	} else {
		goto L151
	}
L151:
	;
	v464 = v450
	goto L4
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg_internal(m, int32(_a_F_detoast_attr_slice_10), v15)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L36
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_detoast_attr_slice_1), int32(215), int32(_a_F_detoast_attr_slice_11))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L36
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dgamma(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v56 float64
	_ = v56
	var v61 float64
	_ = v61
	var v68 float64
	_ = v68
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v90 int32
	_ = v90
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v102 int32
	_ = v102
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 float64
	_ = v166
	var v174 float64
	_ = v174
	var v213 float64
	_ = v213
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v229 float64
	_ = v229
	var v245 float64
	_ = v245
	var v251 float64
	_ = v251
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v306 float64
	_ = v306
	var v323 float64
	_ = v323
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v336 float64
	_ = v336
	var v354 float64
	_ = v354
	var v355 float64
	_ = v355
	var v369 float64
	_ = v369
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	v2 = float64(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
		v369 = v12
		v378 = F_Float8GetDatum(m, v369)
		mBase = m.M
		v381 = m.ExcPending
		if v381 != 0 {
			return int32(0)
		} else {
			return v378
		}
	} else {
		if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_lt(v12, float64(0)) == int32(0) {
				v369 = v12
				v378 = F_Float8GetDatum(m, v369)
				mBase = m.M
				v381 = m.ExcPending
				if v381 != 0 {
					return int32(0)
				} else {
					return v378
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_dgamma[0])) = int32(0)
			v30 = base.I64_reinterpret_f64(v12)
			v35 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(2146435072)) <= base.Ui32(v35) {
				v354 = base.F64_add(v12, math.Float64frombits(uint64(0x7ff0000000000000)))
			} else {
				if base.Ui32(v35) <= base.Ui32(int32(1016070143)) {
					v354 = base.F64_div(float64(1), v12)
				} else {
					if base.F64_ne(v12, base.F64_floor(v12)) != 0 {
						if base.Ui32(int32(1080492032)) <= base.Ui32(v35) {
							v61 = float64(0.5)
							if base.F64_eq(base.F64_floor(base.F64_mul(v12, v61)), base.F64_mul(base.F64_floor(v12), v61)) != 0 {
								v68 = float64(0)
							} else {
								v68 = math.Float64frombits(uint64(0x8000000000000000))
							}
							if v30 < int64(0) {
								v354 = v68
							} else {
								v354 = base.F64_mul(v12, float64(8.98846567431158e+307))
							}
						} else {
							v73 = base.F64_abs(v12)
							v74 = float64(5.52468004077673)
							v75 = base.F64_add(v73, v74)
							v77 = float64(-5.52468004077673)
							if base.F64_gt(v73, v74) != 0 {
								v84 = base.F64_add(base.F64_sub(v75, v73), v77)
							} else {
								v84 = base.F64_sub(base.F64_add(v75, v77), v73)
							}
							v86 = base.F64_add(v73, float64(-0.5))
							if base.F64_lt(v73, float64(8)) != 0 {
								v90 = int32(12)
								v93 = v2
								v95 = v2
								for {
									v102 = v90 << (uint(int32(3)) % 32)
									v103 = *(*float64)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_dgamma[1])))
									v104 = base.F64_add(base.F64_mul(v95, v73), v103)
									v106 = *(*float64)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_dgamma[2])))
									v107 = base.F64_add(base.F64_mul(v93, v73), v106)
									if v90 != 0 {
										v90 = v90 - int32(1)
										v93 = v107
										v95 = v104
										continue
									} else {
										break
									}
									break
								}
								v135 = v107
								v137 = v104
							} else {
								v113 = v2
								v115 = v2
								v118 = int32(0)
								for {
									v122 = v118 << (uint(int32(3)) % 32)
									v123 = *(*float64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_dgamma[1])))
									v124 = base.F64_add(base.F64_div(v115, v73), v123)
									v126 = *(*float64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_dgamma[2])))
									v127 = base.F64_add(base.F64_div(v113, v73), v126)
									v129 = v118 + int32(1)
									if v129 != int32(13) {
										v113 = v127
										v115 = v124
										v118 = v129
										continue
									} else {
										break
									}
									break
								}
								v135 = v127
								v137 = v124
							}
							v144 = F_exp(m, base.F64_neg(v75))
							mBase = m.M
							v145 = base.F64_mul(base.F64_div(v135, v137), v144)
							if base.F64_lt(v12, float64(0)) != 0 {
								v149 = float64(0.5)
								v150 = base.F64_mul(v73, v149)
								v152 = base.F64_sub(v150, base.F64_floor(v150))
								v153 = base.F64_add(v152, v152)
								v157 = int32(1)
								v160 = base.I32_div_s(base.I32_trunc_sat_f64_s(base.F64_mul(v153, float64(4)))+v157, int32(2))
								v166 = base.F64_mul(base.F64_sub(v153, base.F64_mul(base.F64_convert_i32_s(v160), v149)), float64(3.141592653589793))
								switch v160 - v157 {
								case 0:
									v213 = float64(1)
									v214 = base.F64_mul(v166, v166)
									v216 = base.F64_mul(v214, float64(0.5))
									v217 = base.F64_sub(v213, v216)
									v229 = base.F64_mul(v214, v214)
									v323 = base.F64_add(v217, base.F64_add(base.F64_sub(base.F64_sub(v213, v217), v216), base.F64_sub(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v229, v229), base.F64_add(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v166, float64(0)))))
								case 1:
									v245 = base.F64_neg(v166)
									v251 = base.F64_mul(v245, v245)
									v323 = base.F64_add(base.F64_mul(base.F64_mul(v245, v251), base.F64_add(base.F64_mul(v251, base.F64_add(base.F64_mul(base.F64_mul(v251, base.F64_mul(v251, v251)), base.F64_add(base.F64_mul(v251, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v251, base.F64_add(base.F64_mul(v251, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v245)
								case 2:
									v290 = float64(1)
									v291 = base.F64_mul(v166, v166)
									v293 = base.F64_mul(v291, float64(0.5))
									v294 = base.F64_sub(v290, v293)
									v306 = base.F64_mul(v291, v291)
									v323 = base.F64_neg(base.F64_add(v294, base.F64_add(base.F64_sub(base.F64_sub(v290, v294), v293), base.F64_sub(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v306, v306), base.F64_add(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v166, float64(0))))))
								default:
									v174 = base.F64_mul(v166, v166)
									v323 = base.F64_add(base.F64_mul(base.F64_mul(v166, v174), base.F64_add(base.F64_mul(v174, base.F64_add(base.F64_mul(base.F64_mul(v174, base.F64_mul(v174, v174)), base.F64_add(base.F64_mul(v174, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v174, base.F64_add(base.F64_mul(v174, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v166)
								}
								v331 = base.F64_div(float64(-3.141592653589793), base.F64_mul(v145, base.F64_mul(v73, v323)))
								v332 = base.F64_neg(v84)
								v333 = base.F64_neg(v86)
							} else {
								v331 = v145
								v332 = v84
								v333 = v86
							}
							v336 = F_pow(m, v75, base.F64_mul(v333, float64(0.5)))
							mBase = m.M
							v354 = base.F64_mul(v336, base.F64_mul(v336, base.F64_add(v331, base.F64_div(base.F64_mul(base.F64_mul(v332, float64(6.02468004077673)), v331), v75))))
						}
					} else {
						if v30 < int64(0) {
							v354 = math.Float64frombits(uint64(0x7ff8000000000000))
						} else {
							if base.F64_le(v12, float64(23)) == int32(0) {
								if base.Ui32(int32(1080492032)) <= base.Ui32(v35) {
									v61 = float64(0.5)
									if base.F64_eq(base.F64_floor(base.F64_mul(v12, v61)), base.F64_mul(base.F64_floor(v12), v61)) != 0 {
										v68 = float64(0)
									} else {
										v68 = math.Float64frombits(uint64(0x8000000000000000))
									}
									if v30 < int64(0) {
										v354 = v68
									} else {
										v354 = base.F64_mul(v12, float64(8.98846567431158e+307))
									}
								} else {
									v73 = base.F64_abs(v12)
									v74 = float64(5.52468004077673)
									v75 = base.F64_add(v73, v74)
									v77 = float64(-5.52468004077673)
									if base.F64_gt(v73, v74) != 0 {
										v84 = base.F64_add(base.F64_sub(v75, v73), v77)
									} else {
										v84 = base.F64_sub(base.F64_add(v75, v77), v73)
									}
									v86 = base.F64_add(v73, float64(-0.5))
									if base.F64_lt(v73, float64(8)) != 0 {
										v90 = int32(12)
										v93 = v2
										v95 = v2
										for {
											v102 = v90 << (uint(int32(3)) % 32)
											v103 = *(*float64)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_dgamma[1])))
											v104 = base.F64_add(base.F64_mul(v95, v73), v103)
											v106 = *(*float64)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_dgamma[2])))
											v107 = base.F64_add(base.F64_mul(v93, v73), v106)
											if v90 != 0 {
												v90 = v90 - int32(1)
												v93 = v107
												v95 = v104
												continue
											} else {
												break
											}
											break
										}
										v135 = v107
										v137 = v104
									} else {
										v113 = v2
										v115 = v2
										v118 = int32(0)
										for {
											v122 = v118 << (uint(int32(3)) % 32)
											v123 = *(*float64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_dgamma[1])))
											v124 = base.F64_add(base.F64_div(v115, v73), v123)
											v126 = *(*float64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_dgamma[2])))
											v127 = base.F64_add(base.F64_div(v113, v73), v126)
											v129 = v118 + int32(1)
											if v129 != int32(13) {
												v113 = v127
												v115 = v124
												v118 = v129
												continue
											} else {
												break
											}
											break
										}
										v135 = v127
										v137 = v124
									}
									v144 = F_exp(m, base.F64_neg(v75))
									mBase = m.M
									v145 = base.F64_mul(base.F64_div(v135, v137), v144)
									if base.F64_lt(v12, float64(0)) != 0 {
										v149 = float64(0.5)
										v150 = base.F64_mul(v73, v149)
										v152 = base.F64_sub(v150, base.F64_floor(v150))
										v153 = base.F64_add(v152, v152)
										v157 = int32(1)
										v160 = base.I32_div_s(base.I32_trunc_sat_f64_s(base.F64_mul(v153, float64(4)))+v157, int32(2))
										v166 = base.F64_mul(base.F64_sub(v153, base.F64_mul(base.F64_convert_i32_s(v160), v149)), float64(3.141592653589793))
										switch v160 - v157 {
										case 0:
											v213 = float64(1)
											v214 = base.F64_mul(v166, v166)
											v216 = base.F64_mul(v214, float64(0.5))
											v217 = base.F64_sub(v213, v216)
											v229 = base.F64_mul(v214, v214)
											v323 = base.F64_add(v217, base.F64_add(base.F64_sub(base.F64_sub(v213, v217), v216), base.F64_sub(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v229, v229), base.F64_add(base.F64_mul(v214, base.F64_add(base.F64_mul(v214, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v166, float64(0)))))
										case 1:
											v245 = base.F64_neg(v166)
											v251 = base.F64_mul(v245, v245)
											v323 = base.F64_add(base.F64_mul(base.F64_mul(v245, v251), base.F64_add(base.F64_mul(v251, base.F64_add(base.F64_mul(base.F64_mul(v251, base.F64_mul(v251, v251)), base.F64_add(base.F64_mul(v251, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v251, base.F64_add(base.F64_mul(v251, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v245)
										case 2:
											v290 = float64(1)
											v291 = base.F64_mul(v166, v166)
											v293 = base.F64_mul(v291, float64(0.5))
											v294 = base.F64_sub(v290, v293)
											v306 = base.F64_mul(v291, v291)
											v323 = base.F64_neg(base.F64_add(v294, base.F64_add(base.F64_sub(base.F64_sub(v290, v294), v293), base.F64_sub(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v306, v306), base.F64_add(base.F64_mul(v291, base.F64_add(base.F64_mul(v291, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v166, float64(0))))))
										default:
											v174 = base.F64_mul(v166, v166)
											v323 = base.F64_add(base.F64_mul(base.F64_mul(v166, v174), base.F64_add(base.F64_mul(v174, base.F64_add(base.F64_mul(base.F64_mul(v174, base.F64_mul(v174, v174)), base.F64_add(base.F64_mul(v174, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v174, base.F64_add(base.F64_mul(v174, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v166)
										}
										v331 = base.F64_div(float64(-3.141592653589793), base.F64_mul(v145, base.F64_mul(v73, v323)))
										v332 = base.F64_neg(v84)
										v333 = base.F64_neg(v86)
									} else {
										v331 = v145
										v332 = v84
										v333 = v86
									}
									v336 = F_pow(m, v75, base.F64_mul(v333, float64(0.5)))
									mBase = m.M
									v354 = base.F64_mul(v336, base.F64_mul(v336, base.F64_add(v331, base.F64_div(base.F64_mul(base.F64_mul(v332, float64(6.02468004077673)), v331), v75))))
								}
							} else {
								v56 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_trunc_sat_f64_s(v12)<<(uint(int32(3))%32))+uint32(_c_F_dgamma[3])))
								v354 = v56
							}
						}
					}
				}
			}
			v355 = base.F64_abs(v354)
			if base.F64_ne(v355, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v355)) < base.Ui64(int64(9218868437227405313))) == int32(0) {
				if base.F64_ne(v354, float64(0)) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					F_float_underflow_error(m)
					mBase = m.M
					v384 = m.ExcPending
					if v384 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				if base.F64_eq(v354, float64(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v384 = m.ExcPending
					if v384 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v369 = v354
					v378 = F_Float8GetDatum(m, v369)
					mBase = m.M
					v381 = m.ExcPending
					if v381 != 0 {
						return int32(0)
					} else {
						return v378
					}
				}
			}
		}
	}
}
func F_digest_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	m.Env.Pgmem_hash_free(m, v5)
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v7 != 0 {
		F_ResourceOwnerForget(m, v7, v4, int32(_a_F_digest_free_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_pfree(m, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_disable_timeouts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var __phi70 int32
	_ = __phi70
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[0])) = v2
	v20 = v2
	goto L3
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L30
	} else {
		goto L32
	}
L2:
	;
	v160 = int32(-1)
	goto L1
L3:
	;
	v26 = l0 + v20<<(uint(int32(3))%32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = v27 * int32(40)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_disable_timeouts[1]))))
	if v30 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	if int32(0) < v124 {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	if v35 <= v33 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+4)))
	if v112 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v39 = v33
	goto L9
L9:
	;
	v47 = v39 << (uint(int32(2)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_disable_timeouts[3])))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v27 != v49 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	if v57 <= v39 {
		v160 = v39
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v52 = v39 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	if v52 < v54 {
		v39 = v52
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L2
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_disable_timeouts[3])))
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)) = uint8(v62)
	v65 = v39 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	if v65 < v67 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	__phi70 = v39
	__phi72 = v65
	v70 = __phi70
	v72 = __phi72
	goto L19
L17:
	;
	goto L18
L18:
	;
	v98 = int32(_a_F_disable_timeouts_0)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2])) = v100 - int32(1)
	goto L7
L19:
	;
	v77 = int32(2)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(v77)%32))+uint32(_c_F_disable_timeouts[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(v77)%32))+uint32(_c_F_disable_timeouts[3]))) = v83
	v86 = v72 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	if v86 < v88 {
		__phi70 = v72
		__phi72 = v86
		v70 = __phi70
		v72 = __phi72
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_disable_timeouts[4]))) = uint8(v117)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v120 = v20 + int32(1)
	if v120 != int32(2) {
		v20 = v120
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L4
L26:
	;
	v130 = m.G0
	v131 = int32(16)
	v132 = v130 - v131
	m.G0 = v132
	F_gettimeofday(m, v132)
	mBase = m.M
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v136 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+8)))
	m.G0 = v132 + v131
	goto L29
L27:
	;
	goto L28
L28:
	;
	m.G0 = v11 + int32(16)
	return
L29:
	;
	F_schedule_alarm(m, v136+v135*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return
L31:
	;
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v160
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeouts[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v173 - int32(1)
	F_errmsg_internal(m, int32(_a_F_disable_timeouts_1), v11)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_disable_timeouts_2), int32(143), int32(_a_F_disable_timeouts_3))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dispatch_compare_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v4) < base.Ui32(v5) {
		v8 = int32(-1)
	} else {
		v8 = base.B2i32(base.Ui32(v5) < base.Ui32(v4))
	}
	return v8
}
func F_distribute_quals_to_rels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	v14 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	if l1 == v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L8
	} else {
		goto L251
	}
L2:
	;
	m.G0 = v27 + int32(16)
	return
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v55 = v14
	goto L5
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v55<<(uint(int32(2))%32))))
	v63 = F_pull_varnos(m, l0, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v879 = v55 + int32(1)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v879 < v880 {
		v55 = v879
		goto L5
	} else {
		goto L250
	}
L8:
	;
	return
L9:
	;
	v65 = int32(0)
	if v63 == v65 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v118 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v118 = int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	if l5 == int32(0) {
		v111 = v65
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v118 = v111
	goto L10
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v75 < v74 {
		v111 = v65
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v77 = int32(1)
	if v74 <= v77 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v80 = v77
	goto L19
L18:
	;
	v80 = v74
	goto L19
L19:
	;
	v81 = int32(8)
	v86 = int32(0)
	goto L20
L20:
	;
	v93 = v86 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v63+v81+v93)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l5+v81+v93)))
	v100 = v95 & (v97 ^ int32(-1))
	v102 = base.B2i32(v100 == int32(0))
	if v100 != 0 {
		v111 = v102
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v111 = v102
	goto L14
L22:
	;
	v104 = v86 + int32(1)
	if v104 != v80 {
		v86 = v104
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v134 = l2
	goto L28
L25:
	;
	goto L26
L26:
	;
	if l6 != 0 {
		goto L51
	} else {
		goto L52
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L47
	}
L28:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	if v145 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	v206 = F_lappend(m, v205, v62)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L46
	}
L30:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
	v149 = int32(0)
	if v63 == v149 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v202 == int32(0) {
		v134 = v145
		goto L28
	} else {
		goto L45
	}
L32:
	;
	v202 = int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	if v148 == int32(0) {
		v195 = v149
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v202 = v195
	goto L31
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v159 < v158 {
		v195 = v149
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v161 = int32(1)
	if v158 <= v161 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v164 = v161
	goto L40
L39:
	;
	v164 = v158
	goto L40
L40:
	;
	v165 = int32(8)
	v170 = int32(0)
	goto L41
L41:
	;
	v177 = v170 << (uint(int32(2)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v63+v165+v177)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v148+v165+v177)))
	v184 = v179 & (v181 ^ int32(-1))
	v186 = base.B2i32(v184 == int32(0))
	if v184 != 0 {
		v195 = v186
		goto L35
	} else {
		goto L43
	}
L42:
	;
	v195 = v186
	goto L35
L43:
	;
	v188 = v170 + int32(1)
	if v188 != v164 {
		v170 = v188
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L29
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+40)) = v206
	goto L7
L47:
	;
	F_errmsg_internal(m, int32(_a_F_distribute_quals_to_rels_0), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_distribute_quals_to_rels_1), int32(2603), int32(_a_F_distribute_quals_to_rels_2))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v301 = int32(0)
	if base.B2i32(v299 == v301)|base.B2i32(l7 == v301) != 0 {
		v346 = v301
		goto L83
	} else {
		goto L84
	}
L51:
	;
	v222 = int32(0)
	if v63 == v222 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v281 = int32(0)
	if v63 != 0 {
		v299 = v63
		v300 = v281
		goto L50
	} else {
		goto L71
	}
L54:
	;
	if v275 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L55:
	;
	v275 = int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	if l6 == int32(0) {
		v268 = v222
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v275 = v268
	goto L54
L59:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v232 < v231 {
		v268 = v222
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v234 = int32(1)
	if v231 <= v234 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v237 = v234
	goto L63
L62:
	;
	v237 = v231
	goto L63
L63:
	;
	v238 = int32(8)
	v243 = int32(0)
	goto L64
L64:
	;
	v250 = v243 << (uint(int32(2)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v63+v238+v250)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l6+v238+v250)))
	v257 = v252 & (v254 ^ int32(-1))
	v259 = base.B2i32(v257 == int32(0))
	if v257 != 0 {
		v268 = v259
		goto L58
	} else {
		goto L66
	}
L65:
	;
	v268 = v259
	goto L58
L66:
	;
	v261 = v243 + int32(1)
	if v261 != v237 {
		v243 = v261
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v278 = int32(0)
	if v63 != 0 {
		v299 = v63
		v300 = v278
		goto L50
	} else {
		goto L69
	}
L69:
	;
	v279 = F_bms_copy(m, l6)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v299 = v279
	v300 = v278
	goto L50
L71:
	;
	v282 = F_contain_volatile_functions(m, v62)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	if v282 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v284 = F_bms_copy(m, l5)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	if v287 == v290 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v299 = v284
	v300 = v281
	goto L50
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v293 = v292
	goto L79
L78:
	;
	v293 = l5
	goto L79
L79:
	;
	v294 = F_bms_copy(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)) = uint8(v296)
	v299 = v294
	v300 = int32(1)
	goto L50
L81:
	;
	v491 = F_make_restrictinfo(m, l0, v62, v346^int32(1), l10, l11, v300, l4, v478, l8, l7)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L8
	} else {
		goto L129
	}
L82:
	;
	if v346 != 0 {
		goto L95
	} else {
		goto L96
	}
L83:
	;
	goto L82
L84:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v311 < v312 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v314 = v311
	goto L87
L86:
	;
	v314 = v312
	goto L87
L87:
	;
	if v314 <= int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v317 = int32(1)
	goto L90
L89:
	;
	v317 = v314
	goto L90
L90:
	;
	v318 = int32(8)
	v323 = int32(0)
	goto L91
L91:
	;
	v330 = v323 << (uint(int32(2)) % 32)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l7+v318+v330)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v299+v318+v330)))
	v335 = v332 & v334
	v337 = base.B2i32(v335 != int32(0))
	if v335 != 0 {
		v346 = v337
		goto L83
	} else {
		goto L93
	}
L92:
	;
	v346 = v337
	goto L83
L93:
	;
	v339 = v323 + int32(1)
	if v339 != v317 {
		v323 = v339
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	if l12 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v354 = int32(0)
	if v62 == v354 {
		v383 = v354
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v478 = l6
	v488 = int32(0)
	goto L81
L99:
	;
	goto L100
L100:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l12)))
	v351 = F_lappend(m, v350, v62)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v351
	goto L7
L102:
	;
	if v383 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	switch v357 - int32(52) {
	case 0:
		goto L106
	case 1:
		goto L105
	default:
		v383 = v354
		goto L102
	}
L104:
	;
	v383 = int32(0)
	goto L102
L105:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v369 != int32(4) {
		goto L104
	} else {
		goto L112
	}
L106:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v360 != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+12)))
	if v361 != 0 {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v362 == int32(0) {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	if v365 != int32(6) {
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v362)+28))
	if v368 != 0 {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v383 = v362
	goto L102
L112:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v372 == int32(0) {
		goto L104
	} else {
		goto L113
	}
L113:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	if v375 != int32(6) {
		goto L104
	} else {
		goto L114
	}
L114:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372)+28))
	if v378 == int32(0) {
		v383 = v372
		goto L102
	} else {
		goto L115
	}
L115:
	;
	goto L104
L116:
	;
	v478 = v299
	v488 = l9
	goto L81
L117:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	if v386 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v389 == int32(0) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v392 = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v393 <= v392 {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v409 = v392
	v419 = v393
	goto L121
L121:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v409<<(uint(int32(2))%32))))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+20))
	if v425 != int32(5) {
		v436 = v419
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L116
L123:
	;
	v438 = v409 + int32(1)
	if v438 < v436 {
		v409 = v438
		v419 = v436
		goto L121
	} else {
		goto L128
	}
L124:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424)+24))
	if v428 == int32(0) {
		v436 = v419
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	v432 = F_bms_is_member(m, v428, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	if v432 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v436 = v434
	goto L123
L128:
	;
	goto L122
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v491
	v494 = int32(0)
	if v478 == v494 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v539 == int32(2) {
		goto L146
	} else {
		goto L147
	}
L131:
	;
	v539 = int32(0)
	goto L130
L132:
	;
	goto L133
L133:
	;
	v502 = int32(1)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if v503 <= v502 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v506 = v502
	goto L136
L135:
	;
	v506 = v503
	goto L136
L136:
	;
	v510 = int32(0)
	v512 = v494
	goto L137
L137:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v478+int32(8)+v510<<(uint(int32(2))%32))))
	if v519 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v539 = v531
	goto L130
L139:
	;
	goto L138
L140:
	;
	v520 = int32(2)
	if v512 != 0 {
		v531 = v520
		goto L139
	} else {
		goto L143
	}
L141:
	;
	v526 = v512
	goto L142
L142:
	;
	v528 = v510 + int32(1)
	if v528 != v506 {
		v510 = v528
		v512 = v526
		goto L137
	} else {
		goto L145
	}
L143:
	;
	v521 = int32(1)
	if base.Ui32(v521) < base.Ui32(base.I32_popcnt(v519)) {
		v531 = v520
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v526 = v521
	goto L142
L145:
	;
	v531 = v526
	goto L139
L146:
	;
	v543 = F_pull_var_clause(m, v62, int32(26))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L8
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+10)))
	if v554 != 0 {
		goto L156
	} else {
		goto L157
	}
L149:
	;
	if l11 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v546 = F_bms_intersect(m, v478, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L8
	} else {
		goto L153
	}
L151:
	;
	v548 = v478
	goto L152
L152:
	;
	F_add_vars_to_targetlist(m, l0, v543, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L8
	} else {
		goto L154
	}
L153:
	;
	v548 = v546
	goto L152
L154:
	;
	F_list_free(m, v543)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L8
	} else {
		goto L155
	}
L155:
	;
	goto L148
L156:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v491)+96))
	if v583 == int32(0) {
		v851 = v491
		goto L168
	} else {
		goto L169
	}
L157:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	if v555 == int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v558 != int32(17) {
		goto L156
	} else {
		goto L159
	}
L159:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v555)+28))
	if v561 == int32(0) {
		goto L156
	} else {
		goto L160
	}
L160:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v564 != int32(2) {
		goto L156
	} else {
		goto L161
	}
L161:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v570 = F_exprType(m, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	v572 = F_op_mergejoinable(m, v567, v570)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	if v572 == int32(0) {
		goto L156
	} else {
		goto L164
	}
L164:
	;
	v576 = F_contain_volatile_functions(m, v491)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L8
	} else {
		goto L165
	}
L165:
	;
	if v576 != 0 {
		goto L156
	} else {
		goto L166
	}
L166:
	;
	v578 = F_get_mergejoin_opfamilies(m, v567)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+96)) = v578
	goto L156
L168:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v851)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L8
	} else {
		goto L249
	}
L169:
	;
	if v488 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v589 = F_process_equivalence(m, l0, v27+int32(12), v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L8
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v346 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L173:
	;
	if v589 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+96))
	if v592 == int32(0) {
		v851 = v591
		goto L168
	} else {
		goto L175
	}
L175:
	;
	F_initialize_mergeclause_eclasses(m, l0, v591)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L8
	} else {
		goto L176
	}
L176:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v851 = v597
	goto L168
L177:
	;
	F_initialize_mergeclause_eclasses(m, l0, v491)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L8
	} else {
		goto L248
	}
L178:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+9)))
	if v600 != int32(1) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	F_initialize_mergeclause_eclasses(m, l0, v491)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L8
	} else {
		goto L180
	}
L180:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v491)+44))
	v606 = int32(0)
	if v605 == v606 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v491)+48))
	v721 = int32(0)
	if v720 == v721 {
		goto L215
	} else {
		goto L216
	}
L182:
	;
	if v659 == int32(0) {
		goto L181
	} else {
		goto L196
	}
L183:
	;
	v659 = int32(1)
	goto L182
L184:
	;
	goto L185
L185:
	;
	if l7 == int32(0) {
		v652 = v606
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v659 = v652
	goto L182
L187:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v616 < v615 {
		v652 = v606
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v618 = int32(1)
	if v615 <= v618 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v621 = v618
	goto L191
L190:
	;
	v621 = v615
	goto L191
L191:
	;
	v622 = int32(8)
	v627 = int32(0)
	goto L192
L192:
	;
	v634 = v627 << (uint(int32(2)) % 32)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v605+v622+v634)))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l7+v622+v634)))
	v641 = v636 & (v638 ^ int32(-1))
	v643 = base.B2i32(v641 == int32(0))
	if v641 != 0 {
		v652 = v643
		goto L186
	} else {
		goto L194
	}
L193:
	;
	v652 = v643
	goto L186
L194:
	;
	v645 = v627 + int32(1)
	if v645 != v621 {
		v627 = v645
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v491)+48))
	v663 = int32(0)
	if base.B2i32(v662 == v663)|base.B2i32(l7 == v663) != 0 {
		v708 = v663
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v708 != 0 {
		goto L181
	} else {
		goto L210
	}
L198:
	;
	goto L197
L199:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v662)+4))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v673 < v674 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v676 = v673
	goto L202
L201:
	;
	v676 = v674
	goto L202
L202:
	;
	if v676 <= int32(1) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v679 = int32(1)
	goto L205
L204:
	;
	v679 = v676
	goto L205
L205:
	;
	v680 = int32(8)
	v685 = int32(0)
	goto L206
L206:
	;
	v692 = v685 << (uint(int32(2)) % 32)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l7+v680+v692)))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v662+v680+v692)))
	v697 = v694 & v696
	v699 = base.B2i32(v697 != int32(0))
	if v697 != 0 {
		v708 = v699
		goto L198
	} else {
		goto L208
	}
L207:
	;
	v708 = v699
	goto L198
L208:
	;
	v701 = v685 + int32(1)
	if v701 != v679 {
		v685 = v701
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v710 = F_palloc0(m, int32(12))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L8
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v710))) = int32(321)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v717 = F_lappend(m, v716, v710)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L8
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v717
	goto L7
L213:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v835 != int32(2) {
		v851 = v491
		goto L168
	} else {
		goto L245
	}
L214:
	;
	if v774 == int32(0) {
		goto L213
	} else {
		goto L228
	}
L215:
	;
	v774 = int32(1)
	goto L214
L216:
	;
	goto L217
L217:
	;
	if l7 == int32(0) {
		v767 = v721
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v774 = v767
	goto L214
L219:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v720)+4))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v731 < v730 {
		v767 = v721
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v733 = int32(1)
	if v730 <= v733 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v736 = v733
	goto L223
L222:
	;
	v736 = v730
	goto L223
L223:
	;
	v737 = int32(8)
	v742 = int32(0)
	goto L224
L224:
	;
	v749 = v742 << (uint(int32(2)) % 32)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v720+v737+v749)))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l7+v737+v749)))
	v756 = v751 & (v753 ^ int32(-1))
	v758 = base.B2i32(v756 == int32(0))
	if v756 != 0 {
		v767 = v758
		goto L218
	} else {
		goto L226
	}
L225:
	;
	v767 = v758
	goto L218
L226:
	;
	v760 = v742 + int32(1)
	if v760 != v736 {
		v742 = v760
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v491)+44))
	v778 = int32(0)
	if base.B2i32(v777 == v778)|base.B2i32(l7 == v778) != 0 {
		v823 = v778
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v823 != 0 {
		goto L213
	} else {
		goto L242
	}
L230:
	;
	goto L229
L231:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v777)+4))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v788 < v789 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v791 = v788
	goto L234
L233:
	;
	v791 = v789
	goto L234
L234:
	;
	if v791 <= int32(1) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v794 = int32(1)
	goto L237
L236:
	;
	v794 = v791
	goto L237
L237:
	;
	v795 = int32(8)
	v800 = int32(0)
	goto L238
L238:
	;
	v807 = v800 << (uint(int32(2)) % 32)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l7+v795+v807)))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v777+v795+v807)))
	v812 = v809 & v811
	v814 = base.B2i32(v812 != int32(0))
	if v812 != 0 {
		v823 = v814
		goto L230
	} else {
		goto L240
	}
L239:
	;
	v823 = v814
	goto L230
L240:
	;
	v816 = v800 + int32(1)
	if v816 != v794 {
		v800 = v816
		goto L238
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v825 = F_palloc0(m, int32(12))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L8
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v825)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v825)+4)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v825))) = int32(321)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v832 = F_lappend(m, v831, v825)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L8
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v832
	goto L7
L245:
	;
	v839 = F_palloc0(m, int32(12))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L8
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v839)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v839)+4)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v839))) = int32(321)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v846 = F_lappend(m, v845, v839)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L8
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v846
	goto L7
L248:
	;
	v851 = v491
	goto L168
L249:
	;
	goto L7
L250:
	;
	goto L6
L251:
	;
	F_errmsg_internal(m, int32(_a_F_distribute_quals_to_rels_3), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L8
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_distribute_quals_to_rels_1), int32(2611), int32(_a_F_distribute_quals_to_rels_2))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L8
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_distribute_restrictinfo_to_rels(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
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
	var v234 int32
	_ = v234
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
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(0)
	if v11 == v14 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L23
	} else {
		goto L92
	}
L4:
	;
	m.G0 = v9 + int32(16)
	return
L5:
	;
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v68 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v22 = int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v23 <= v22 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v22
	goto L11
L10:
	;
	v26 = v23
	goto L11
L11:
	;
	v31 = int32(0)
	v33 = int32(-1)
	goto L13
L12:
	;
	v68 = v60
	goto L5
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)+v31<<(uint(int32(2))%32))))
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12)))) = v52
	v60 = int32(1)
	goto L12
L15:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v41)))|base.B2i32(int32(0) <= v33) != 0 {
		v60 = v14
		goto L12
	} else {
		goto L18
	}
L16:
	;
	v52 = v33
	goto L17
L17:
	;
	v54 = v31 + int32(1)
	if v54 != v26 {
		v31 = v54
		v33 = v52
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v52 = base.I32_ctz(v41) | v31<<(uint(int32(5))%32)
	goto L17
L19:
	;
	goto L14
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v70 = F_find_base_rel(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v119 != 0 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	return
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v69<<(uint(int32(2))%32))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+20)))
	if v77 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v70)+184))
	v111 = F_lappend(m, v110, v107)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L23
	} else {
		goto L36
	}
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+21)))
	if v80 != int32(112) {
		v107 = l1
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v83 = F_restriction_is_always_true(m, l0, l1)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v83 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v85 = F_restriction_is_always_false(m, l0, l1)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	if v85 == int32(0) {
		v107 = l1
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v91 = int32(0)
	v93 = F_makeBoolConst(m, v91, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L23
	} else {
		goto L34
	}
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v103 = F_make_restrictinfo(m, l0, v93, v95, v96, v97, v98, int32(0), v100, v101, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L23
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+56)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v89
	v107 = v103
	goto L25
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+184)) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	if base.Ui32(v114) < base.Ui32(v115) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v117 = v114
	goto L39
L38:
	;
	v117 = v115
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+208)) = v117
	goto L4
L40:
	;
	F_check_memoizable(m, l1)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L23
	} else {
		goto L51
	}
L41:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v120 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v123 != int32(17) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	if v126 == int32(0) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v129 != int32(2) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = F_exprType(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v137 = F_op_hashjoinable(m, v132, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	if v137 == int32(0) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v141 = F_contain_volatile_functions(m, l1)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	if v141 != 0 {
		goto L40
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v132
	goto L40
L51:
	;
	v148 = F_restriction_is_always_true(m, l0, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L53
	}
L52:
	;
	goto L4
L53:
	;
	if v148 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v150 = F_restriction_is_always_false(m, l0, l1)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	if v150 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v154 = int32(0)
	v156 = F_makeBoolConst(m, v154, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L23
	} else {
		goto L59
	}
L57:
	;
	v170 = l1
	goto L58
L58:
	;
	if v11 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v166 = F_make_restrictinfo(m, l0, v156, v158, v159, v160, v161, int32(0), v163, v164, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+56)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v152
	v170 = v166
	goto L58
L61:
	;
	if v229 < int32(0) {
		goto L52
	} else {
		goto L72
	}
L62:
	;
	v229 = base.I32_ctz(v215) | v216<<(uint(int32(5))%32)
	goto L61
L63:
	;
	v229 = int32(-2)
	goto L61
L64:
	;
	v182 = base.I32_div_s(int32(0), int32(32))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v183 <= v182 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v186 = v11 + int32(8)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v182<<(uint(int32(2))%32))))
	v193 = v190 & int32(-1)
	if v193 != 0 {
		v215 = v193
		v216 = v182
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v195 = v182 + int32(1)
	if v195 == v183 {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v198 = v195
	goto L68
L68:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v186+v198<<(uint(int32(2))%32))))
	if v205 != 0 {
		v215 = v205
		v216 = v198
		goto L62
	} else {
		goto L70
	}
L69:
	;
	goto L63
L70:
	;
	v207 = v198 + int32(1)
	if v207 != v183 {
		v198 = v207
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v234 = v229
	goto L73
L73:
	;
	v238 = F_find_base_rel_ignore_join(m, l0, v234)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L23
	} else {
		goto L75
	}
L74:
	;
	goto L52
L75:
	;
	if v238 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v238)+212))
	v241 = F_lappend(m, v240, v170)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L23
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v11 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+212)) = v241
	goto L78
L80:
	;
	if int32(0) <= v299 {
		v234 = v299
		goto L73
	} else {
		goto L91
	}
L81:
	;
	v299 = base.I32_ctz(v285) | v286<<(uint(int32(5))%32)
	goto L80
L82:
	;
	v299 = int32(-2)
	goto L80
L83:
	;
	v250 = v234 + int32(1)
	v252 = base.I32_div_s(v250, int32(32))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v253 <= v252 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v256 = v11 + int32(8)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v252<<(uint(int32(2))%32))))
	v263 = v260 & (int32(-1) << (uint(v250) % 32))
	if v263 != 0 {
		v285 = v263
		v286 = v252
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v265 = v252 + int32(1)
	if v265 == v253 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v268 = v265
	goto L87
L87:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256+v268<<(uint(int32(2))%32))))
	if v275 != 0 {
		v285 = v275
		v286 = v268
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L82
L89:
	;
	v277 = v268 + int32(1)
	if v277 != v253 {
		v268 = v277
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L74
L92:
	;
	F_errmsg_internal(m, int32(_a_F_distribute_restrictinfo_to_rels_0), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L23
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_distribute_restrictinfo_to_rels_1), int32(3276), int32(_a_F_distribute_restrictinfo_to_rels_2))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L23
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_div_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v153 int32
	_ = v153
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v196 int64
	_ = v196
	var v199 int64
	_ = v199
	var v202 int64
	_ = v202
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v210 int64
	_ = v210
	var v217 int64
	_ = v217
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v286 int64
	_ = v286
	var v293 int64
	_ = v293
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v325 int64
	_ = v325
	var v353 int32
	_ = v353
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v368 int32
	_ = v368
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
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
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v678 int32
	_ = v678
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int64
	_ = v756
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v830 int32
	_ = v830
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int64
	_ = v867
	var v868 int64
	_ = v868
	var v870 int64
	_ = v870
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int64
	_ = v882
	var v885 int64
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v900 int32
	_ = v900
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int64
	_ = v937
	var v940 int64
	_ = v940
	var v955 int32
	_ = v955
	var v978 int32
	_ = v978
	var v982 int64
	_ = v982
	var v984 int64
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v993 int64
	_ = v993
	var v995 int64
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1028 int32
	_ = v1028
	var v1047 int32
	_ = v1047
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1092 int32
	_ = v1092
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1144 int32
	_ = v1144
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1191 float64
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 float64
	_ = v1197
	var v1202 float64
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1207 int64
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1216 int64
	_ = v1216
	var v1218 int64
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int64
	_ = v1256
	var v1257 float64
	_ = v1257
	var v1259 float64
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1274 int64
	_ = v1274
	var v1277 int64
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1288 int32
	_ = v1288
	var v1291 int64
	_ = v1291
	var v1320 int32
	_ = v1320
	var v1321 int64
	_ = v1321
	var v1322 int64
	_ = v1322
	var v1325 int64
	_ = v1325
	var v1328 int64
	_ = v1328
	var v1330 int64
	_ = v1330
	var v1338 int64
	_ = v1338
	var v1342 int64
	_ = v1342
	var v1343 int64
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1348 int64
	_ = v1348
	var v1349 int64
	_ = v1349
	var v1357 int64
	_ = v1357
	var v1358 int64
	_ = v1358
	var v1359 int64
	_ = v1359
	var v1384 float64
	_ = v1384
	var v1386 int64
	_ = v1386
	var v1392 float64
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1405 int64
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1414 int64
	_ = v1414
	var v1415 int64
	_ = v1415
	var v1416 int64
	_ = v1416
	var v1443 int64
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1464 int32
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int64
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1501 int64
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int64
	_ = v1510
	var v1514 int64
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1530 int32
	_ = v1530
	var v1562 int32
	_ = v1562
	var v1563 int64
	_ = v1563
	var v1567 int64
	_ = v1567
	var v1606 int64
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1614 int64
	_ = v1614
	var v1615 int64
	_ = v1615
	var v1616 int64
	_ = v1616
	var v1617 int64
	_ = v1617
	var v1645 int64
	_ = v1645
	var v1656 int32
	_ = v1656
	var v1690 int32
	_ = v1690
	var v1693 int64
	_ = v1693
	var v1707 int64
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1736 int32
	_ = v1736
	var v1737 int64
	_ = v1737
	var v1738 int64
	_ = v1738
	var v1741 int64
	_ = v1741
	var v1744 int64
	_ = v1744
	var v1746 int64
	_ = v1746
	var v1754 int64
	_ = v1754
	var v1758 int64
	_ = v1758
	var v1759 int64
	_ = v1759
	var v1766 int64
	_ = v1766
	var v1777 int64
	_ = v1777
	var v1809 int32
	_ = v1809
	var v1818 int64
	_ = v1818
	var v1845 int32
	_ = v1845
	var v1854 int32
	_ = v1854
	var v1887 int64
	_ = v1887
	var v1891 int64
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1898 int64
	_ = v1898
	var v1909 int32
	_ = v1909
	var v1910 int64
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1941 int32
	_ = v1941
	var v1942 int64
	_ = v1942
	var v1946 int64
	_ = v1946
	var v1948 int64
	_ = v1948
	var v1953 int64
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1960 int64
	_ = v1960
	var v1964 int64
	_ = v1964
	var v1968 int64
	_ = v1968
	var v1973 int64
	_ = v1973
	var v1976 int64
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1989 int32
	_ = v1989
	var v1990 int64
	_ = v1990
	var v2021 int32
	_ = v2021
	var v2022 int64
	_ = v2022
	var v2026 int64
	_ = v2026
	var v2028 int64
	_ = v2028
	var v2033 int64
	_ = v2033
	var v2043 int64
	_ = v2043
	var v2072 int64
	_ = v2072
	var v2079 int64
	_ = v2079
	var v2081 int64
	_ = v2081
	var v2109 int64
	_ = v2109
	var v2110 int64
	_ = v2110
	var v2112 int64
	_ = v2112
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2144 int32
	_ = v2144
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2157 int32
	_ = v2157
	var v2158 int64
	_ = v2158
	var v2189 int32
	_ = v2189
	var v2193 int64
	_ = v2193
	var v2194 int64
	_ = v2194
	var v2196 int64
	_ = v2196
	var v2200 int32
	_ = v2200
	var v2201 int64
	_ = v2201
	var v2203 int64
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2208 int64
	_ = v2208
	var v2209 int64
	_ = v2209
	var v2211 int64
	_ = v2211
	var v2213 int64
	_ = v2213
	var v2217 int64
	_ = v2217
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2292 int64
	_ = v2292
	var v2300 int32
	_ = v2300
	var v2319 int32
	_ = v2319
	var v2323 int64
	_ = v2323
	var v2324 int64
	_ = v2324
	var v2327 int64
	_ = v2327
	var v2330 int64
	_ = v2330
	var v2332 int64
	_ = v2332
	var v2340 int64
	_ = v2340
	var v2344 int64
	_ = v2344
	var v2345 int64
	_ = v2345
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2355 int32
	_ = v2355
	var v2395 int32
	_ = v2395
	var v2404 int32
	_ = v2404
	var v2407 int32
	_ = v2407
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2468 int32
	_ = v2468
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2504 int32
	_ = v2504
	var v2525 int32
	_ = v2525
	var v2530 int32
	_ = v2530
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2558 int32
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2568 int32
	_ = v2568
	var v2573 int32
	_ = v2573
	var v2603 int32
	_ = v2603
	var v2609 int32
	_ = v2609
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2659 int32
	_ = v2659
	var v2697 int32
	_ = v2697
	var v2702 int32
	_ = v2702
	v7 = int64(0)
	v12 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(48)
	m.G0 = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v40 == v12 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v38 + int32(48)
	return
L2:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v2255 != 0 {
		goto L280
	} else {
		goto L281
	}
L3:
	;
	v2151 = v1690 - int32(8)
	v2157 = v1012
	v2158 = v1766
	goto L273
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L17
	} else {
		goto L269
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L17
	} else {
		goto L265
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	if v44 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v40 <= int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = base.I32_extend16_s(v44)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v40) <= base.Ui32(int32(4)) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v59 = v55 + v49*int32(_a_F_div_var_0)
	v60 = v50 - int32(1)
	goto L13
L12:
	;
	v59 = v49
	v60 = v50
	goto L13
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v63 == int32(_a_F_div_var_1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = int32(0) - v59
	goto L16
L15:
	;
	v66 = v59
	goto L16
L16:
	;
	F_div_var_int(m, l0, v66, v60, l2, l3, l4)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	goto L1
L19:
	;
	v72 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v73 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v76 = int64(10000)
	v81 = v72 + (v73+base.I64_extend16_s(base.I64_extend_i32_u(v44))*v76)*v76
	if v40 != int32(3) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if v69 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L22:
	;
	v84 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v88 = v84 + v81*int64(10000)
	goto L24
L23:
	;
	v88 = v81
	goto L24
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v91 == int32(_a_F_div_var_1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = int64(0) - v88
	goto L27
L26:
	;
	v94 = v88
	goto L27
L27:
	;
	if v94 == int64(0) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v69 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v99 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v111 = int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v117 = v112 + (v40 + (v113 ^ int32(-1)))
	v121 = base.I32_div_s(l3+int32(3), int32(4))
	v124 = v117 + v121 + v111
	if v124 <= v111 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_pfree(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v105 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v105
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v105
	goto L1
L35:
	;
	goto L34
L36:
	;
	v127 = v111
	goto L38
L37:
	;
	v127 = v124
	goto L38
L38:
	;
	v128 = v127 + l4
	v133 = F_palloc(m, v128<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	v135 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v133))) = uint16(v135)
	v138 = v133 + int32(2)
	v145 = v88 >> (uint(int64(63)) % 64)
	v147 = v88 ^ v145 - v145
	if base.Ui64(int64(1844674407370956)) <= base.Ui64(v147) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v405 != 0 {
		goto L61
	} else {
		goto L62
	}
L41:
	;
	if v128 <= int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v128 <= int32(0) {
		goto L40
	} else {
		goto L54
	}
L44:
	;
	v153 = int32(0)
	v161 = v7
	v162 = v7
	goto L45
L45:
	;
	v189 = v38 + int32(32)
	v190 = int64(10000)
	v191 = int64(0)
	v196 = int64(32)
	v199 = int64(base.Ui64(v162) >> (uint(v196) % 64))
	v202 = int64(4294967295)
	v205 = v162 & v202
	v206 = v190 * v205
	v210 = int64(base.Ui64(v206)>>(uint(v196)%64)) + v190*v199
	v217 = v205*v191 + v210&v202
	*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v162*v191 + v161*v190 + v191*v199 + int64(base.Ui64(v210)>>(uint(v196)%64)) + int64(base.Ui64(v217)>>(uint(v196)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v189))) = v206&v202 | v217<<(uint(v196)%64)
	goto L47
L46:
	;
	goto L40
L47:
	;
	v229 = v38 + int32(16)
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v38)+32))
	if v153 < v69 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v109+v153<<(uint(int32(1))%32)))))
	v237 = v235
	goto L50
L49:
	;
	v237 = int32(0)
	goto L50
L50:
	;
	v238 = base.I64_extend_i32_s(v237)
	v239 = v230 + v238
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v38)+40))
	v246 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v230))) + (v242 + v238>>(uint(int64(63))%64))
	v248 = m.G0
	v249 = int32(16)
	v250 = v248 - v249
	m.G0 = v250
	F___udivmodti4(m, v250, v239, v246, v147, int64(0))
	mBase = m.M
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v250)+8))
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
	*(*int64)(unsafe.Add(mBase, uint32(v229))) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v229)+8)) = v254
	m.G0 = v250 + v249
	goto L51
L51:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v138+v153<<(uint(int32(1))%32)))) = uint16(v264)
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
	v272 = int64(32)
	v273 = int64(base.Ui64(v147) >> (uint(v272) % 64))
	v275 = int64(base.Ui64(v264) >> (uint(v272) % 64))
	v278 = int64(4294967295)
	v279 = v147 & v278
	v281 = v264 & v278
	v282 = v279 * v281
	v286 = int64(base.Ui64(v282)>>(uint(v272)%64)) + v279*v275
	v293 = v281*v273 + v286&v278
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v264*int64(0) + v266*v147 + v273*v275 + int64(base.Ui64(v286)>>(uint(v272)%64)) + int64(base.Ui64(v293)>>(uint(v272)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v282&v278 | v293<<(uint(v272)%64)
	goto L52
L52:
	;
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	v312 = v153 + int32(1)
	if v312 != v128 {
		v153 = v312
		v161 = v246 - v304 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v306)))
		v162 = v239 - v306
		goto L45
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	v317 = int32(0)
	v325 = v7
	goto L55
L55:
	;
	v353 = v317 << (uint(int32(1)) % 32)
	if v317 < v69 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L40
L57:
	;
	v357 = int64(*(*int16)(unsafe.Add(mBase, uint32(v353+v109))))
	v359 = v357
	goto L59
L58:
	;
	v359 = int64(0)
	goto L59
L59:
	;
	v362 = v359 + v325*int64(10000)
	v363 = base.I64_div_u_s(v362, v147)
	*(*uint16)(unsafe.Add(mBase, uint32(v138+v353))) = uint16(v363)
	v368 = v317 + int32(1)
	if v368 != v128 {
		v317 = v368
		v325 = v362 - v363*v147
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	F_pfree(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L17
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v128
	if base.B2i32(v110 == v135)^base.B2i32(int64(0) < v94) != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v413 = int32(_a_F_div_var_1)
	goto L67
L66:
	;
	v413 = int32(0)
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v117
	if l4 != 0 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v711
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v716
	goto L1
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v711 = int32(0)
	v716 = v678
	goto L68
L70:
	;
	if int32(0) < v574 {
		goto L104
	} else {
		goto L105
	}
L71:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v574 = v573
	v576 = v572
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v425 = l3 + v422<<(uint(int32(2))%32)
	if v425+int32(4) < int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v543 = l3 + v117<<(uint(int32(2))%32)
	if v543+int32(4) <= int32(0) {
		v678 = v138
		goto L69
	} else {
		goto L101
	}
L75:
	;
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v436 = l3 & int32(3)
	v440 = base.I32_div_s(v425+int32(7), int32(4))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v441 <= v440 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	goto L75
L80:
	;
	if int32(0) <= v506 {
		goto L79
	} else {
		goto L100
	}
L81:
	;
	v486 = v480
	goto L94
L82:
	;
	v455 = int32(1)
	v456 = v440 - v455
	v459 = v434 + v456<<(uint(v455)%32)
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v459))))
	v461 = int32(2)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v436<<(uint(v461)%32))+uint32(_c_F_div_var[0])))
	v464 = base.I32_rem_s(v460, v463)
	v465 = v460 - v464
	*(*uint16)(unsafe.Add(mBase, uint32(v459))) = uint16(v465)
	v468 = base.I32_div_s(v463, v461)
	if v464 < v468 {
		v506 = v456
		goto L80
	} else {
		goto L89
	}
L83:
	;
	if base.B2i32(v436 == int32(0))|base.B2i32(v440 != v441) != 0 {
		goto L79
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v440
	if v436 != 0 {
		goto L82
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v440
	goto L82
L87:
	;
	v452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v434+v440<<(uint(int32(1))%32)))))
	if v452 <= int32(_a_F_div_var_2) {
		v506 = v440
		goto L80
	} else {
		goto L88
	}
L88:
	;
	v480 = v440
	goto L81
L89:
	;
	v471 = v463 + base.I32_extend16_s(v465)
	if int32(_a_F_div_var_3) < v471 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v476 = v471 + int32(_a_F_div_var_4)
	goto L92
L91:
	;
	v476 = v471
	goto L92
L92:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v459))) = uint16(v476)
	if v471 < int32(_a_F_div_var_0) {
		v506 = v456
		goto L80
	} else {
		goto L93
	}
L93:
	;
	v480 = v456
	goto L81
L94:
	;
	v492 = int32(1)
	v493 = v486 - v492
	v496 = v434 + v493<<(uint(v492)%32)
	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v496))))
	v501 = base.B2i32(int32(_a_F_div_var_5) < v499)
	if int32(_a_F_div_var_5) < v499 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v506 = v493
	goto L80
L96:
	;
	v502 = int32(-9999)
	goto L98
L97:
	;
	v502 = v492
	goto L98
L98:
	;
	v503 = v502 + v499
	*(*uint16)(unsafe.Add(mBase, uint32(v496))) = uint16(v503)
	if int32(_a_F_div_var_5) < v499 {
		v486 = v493
		goto L94
	} else {
		goto L99
	}
L99:
	;
	goto L95
L100:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v514 - int32(2)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v519 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v518 + v519
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v522 + v519
	goto L79
L101:
	;
	v551 = base.I32_div_s(v543+int32(7), int32(4))
	if v127 < v551 {
		goto L71
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v551
	v555 = l3 & int32(3)
	if v555 == int32(0) {
		v574 = v551
		v576 = v138
		goto L70
	} else {
		goto L103
	}
L103:
	;
	v561 = int32(2)
	v562 = v138 + v551<<(uint(int32(1))%32) - v561
	v563 = int32(*(*int16)(unsafe.Add(mBase, uint32(v562))))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v555<<(uint(v561)%32))+uint32(_c_F_div_var[0])))
	v567 = base.I32_rem_s(v563, v566)
	v568 = v563 - v567
	*(*uint16)(unsafe.Add(mBase, uint32(v562))) = uint16(v568)
	goto L71
L104:
	;
	v582 = v574
	v587 = v576
	goto L107
L105:
	;
	goto L106
L106:
	;
	if v574 != 0 {
		v711 = v574
		v716 = v576
		goto L68
	} else {
		goto L117
	}
L107:
	;
	v617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587))))
	if v617 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v678 = v576 + v574<<(uint(int32(1))%32)
	goto L69
L109:
	;
	v618 = v582
	goto L112
L110:
	;
	goto L111
L111:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v664 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v663 - v664
	if v664 < v582 {
		v582 = v582 - v664
		v587 = v587 + int32(2)
		goto L107
	} else {
		goto L116
	}
L112:
	;
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587+v618<<(uint(int32(1))%32)-int32(2)))))
	if v658 != 0 {
		v711 = v618
		v716 = v587
		goto L68
	} else {
		goto L114
	}
L114:
	;
	v659 = int32(1)
	if v659 < v618 {
		v618 = v618 - v659
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v678 = v587
	goto L69
L116:
	;
	goto L108
L117:
	;
	v678 = v576
	goto L69
L118:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v750 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v764 = l5 | base.B2i32(base.Ui32(v40) < base.Ui32(int32(13)))
	if v764 != 0 {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	F_pfree(m, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L17
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v756 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v756
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v756
	goto L1
L124:
	;
	goto L123
L125:
	;
	v765 = int32(1)
	goto L127
L126:
	;
	v765 = int32(5)
	goto L127
L127:
	;
	v767 = int32(1)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v770 = v768 - v769
	v774 = base.I32_div_s(l3+int32(3), int32(4))
	v777 = v770 + v774 + int32(2)
	if v777 <= v767 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v780 = v767
	goto L130
L129:
	;
	v780 = v777
	goto L130
L130:
	;
	v781 = v765 + l4 + v780
	v782 = int32(2)
	v783 = base.I32_div_s(v781, v782)
	v784 = int32(1)
	v787 = base.I32_div_s(v40+v784, v782)
	v791 = base.I32_div_s(v69+v784, v782)
	if v764 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v805 = v800 << (uint(int32(3)) % 32)
	v811 = F_palloc(m, v805+v799<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L17
	} else {
		goto L144
	}
L132:
	;
	v792 = v787 + v783
	if v791 < v792 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	if v787 < v783 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v794 = v791
	goto L137
L136:
	;
	v794 = v792
	goto L137
L137:
	;
	v799 = v787
	v800 = v792
	v801 = v794
	goto L131
L138:
	;
	v796 = v787
	goto L140
L139:
	;
	v796 = v783
	goto L140
L140:
	;
	if v791 < v783 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v798 = v791
	goto L143
L142:
	;
	v798 = v783
	goto L143
L143:
	;
	v799 = v796
	v800 = v783
	v801 = v798
	goto L131
L144:
	;
	v814 = v801 - int32(1)
	if v814 <= int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v982 = int64(*(*int16)(unsafe.Add(mBase, uint32(v978+v955<<(uint(int32(2))%32)))))
	v984 = v982 * int64(10000)
	v985 = int32(1)
	v988 = v955<<(uint(v985)%32) | v985
	if v988 < v69 {
		goto L156
	} else {
		goto L157
	}
L146:
	;
	v955 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v818 = int32(0)
	if v801 != int32(2) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v830 = v818
	v849 = v12
	goto L152
L150:
	;
	v900 = v818
	goto L151
L151:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v936 = v933 + v900<<(uint(int32(2))%32)
	v937 = int64(*(*int16)(unsafe.Add(mBase, uint32(v936))))
	v940 = int64(*(*int16)(unsafe.Add(mBase, uint32(v936)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v811+v900<<(uint(int32(3))%32)))) = v937*int64(10000) + v940
	v955 = v814
	goto L145
L152:
	;
	v860 = int32(3)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v864 = int32(2)
	v866 = v863 + v830<<(uint(v864)%32)
	v867 = int64(*(*int16)(unsafe.Add(mBase, uint32(v866))))
	v868 = int64(10000)
	v870 = int64(*(*int16)(unsafe.Add(mBase, uint32(v866)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v811+v830<<(uint(v860)%32)))) = v867*v868 + v870
	v874 = v830 | int32(1)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v881 = v878 + v874<<(uint(v864)%32)
	v882 = int64(*(*int16)(unsafe.Add(mBase, uint32(v881))))
	v885 = int64(*(*int16)(unsafe.Add(mBase, uint32(v881)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v811+v874<<(uint(v860)%32)))) = v882*v868 + v885
	v889 = v830 + v864
	v891 = v849 + v864
	if v891 != v814&int32(2147483646) {
		v830 = v889
		v849 = v891
		goto L152
	} else {
		goto L154
	}
L153:
	;
	if v814&int32(1) == int32(0) {
		v955 = v814
		goto L145
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	v900 = v889
	goto L151
L156:
	;
	v993 = int64(*(*int16)(unsafe.Add(mBase, uint32(v978+v988<<(uint(int32(1))%32)))))
	v995 = v984 + v993
	goto L158
L157:
	;
	v995 = v984
	goto L158
L158:
	;
	v996 = int32(3)
	v998 = v811 + v955<<(uint(v996)%32)
	*(*int64)(unsafe.Add(mBase, uint32(v998))) = v995
	v1003 = (v800 - v955) << (uint(v996) % 32)
	if v1003 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	base.MemoryFill(m, v998+int32(8), int32(0), v1003)
	goto L161
L160:
	;
	goto L161
L161:
	;
	v1008 = v805 + v811
	v1010 = v1008 + int32(8)
	v1012 = v799 - int32(1)
	if v799 < int32(2) {
		v1144 = int32(0)
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1170 = v1144 << (uint(int32(2)) % 32)
	v1172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1168+v1170))))
	v1174 = v1172 * int32(_a_F_div_var_0)
	v1176 = int32(1)
	v1179 = v1144<<(uint(v1176)%32) | v1176
	if base.Ui32(v1179) < base.Ui32(v40) {
		goto L171
	} else {
		goto L172
	}
L163:
	;
	v1015 = int32(0)
	if v799 != int32(2) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1028 = v1015
	v1047 = int32(0)
	goto L167
L165:
	;
	v1092 = v1015
	goto L166
L166:
	;
	v1123 = v1092 << (uint(int32(2)) % 32)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1126 = v1125 + v1123
	v1127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1126))))
	v1130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1126)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1010+v1123))) = v1127*int32(_a_F_div_var_0) + v1130
	v1144 = v1012
	goto L162
L167:
	;
	v1058 = int32(2)
	v1059 = v1028 << (uint(v1058) % 32)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1062 = v1061 + v1059
	v1063 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1062))))
	v1064 = int32(_a_F_div_var_0)
	v1066 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1062)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1010+v1059))) = v1063*v1064 + v1066
	v1070 = v1059 | int32(4)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1073 = v1072 + v1070
	v1074 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1073))))
	v1077 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1073)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1010+v1070))) = v1074*v1064 + v1077
	v1081 = v1028 + v1058
	v1083 = v1047 + v1058
	if v1083 != v1012&int32(-2) {
		v1028 = v1081
		v1047 = v1083
		goto L167
	} else {
		goto L169
	}
L168:
	;
	if v1012&int32(1) == int32(0) {
		v1144 = v1012
		goto L162
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v1092 = v1081
	goto L166
L171:
	;
	v1184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1168+v1179<<(uint(int32(1))%32)))))
	v1186 = v1174 + v1184
	goto L173
L172:
	;
	v1186 = v1174
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1170+v1010))) = v1186
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1010)))
	v1191 = base.F64_mul(base.F64_convert_i32_s(v1188), float64(1e+08))
	if int32(2) <= v799 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+12))
	v1197 = base.F64_add(v1191, base.F64_convert_i32_s(v1194))
	goto L176
L175:
	;
	v1197 = v1191
	goto L176
L176:
	;
	if int32(2) <= v781 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v1202 = base.F64_div(float64(1), v1197)
	v1204 = v800 - int32(1)
	v1207 = *(*int64)(unsafe.Add(mBase, uint32(v811)))
	v1211 = int32(0)
	v1216 = v1207
	v1218 = int64(1)
	v1222 = v800
	goto L180
L178:
	;
	v1656 = int32(0)
	goto L179
L179:
	;
	if v764 == int32(0) {
		goto L2
	} else {
		goto L220
	}
L180:
	;
	v1245 = int32(3)
	v1247 = v811 + v1211<<(uint(v1245)%32)
	v1252 = v1211 + int32(1)
	v1255 = v811 + v1252<<(uint(v1245)%32)
	v1256 = *(*int64)(unsafe.Add(mBase, uint32(v1255)))
	v1257 = base.F64_convert_i64_s(v1256)
	v1259 = base.F64_mul(v1202, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v1216), float64(1e+08)), v1257))
	v1263 = int32(0)
	v1265 = base.I32_trunc_sat_f64_s(v1259) - base.B2i32(base.F64_ge(v1259, float64(0)) == v1263)
	if v1265 == v1263 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v1656 = v783
	goto L179
L182:
	;
	v1645 = v1615 + v1614*int64(100000000)
	*(*int64)(unsafe.Add(mBase, uint32(v1255))) = v1645
	*(*int64)(unsafe.Add(mBase, uint32(v1247))) = v1617
	if v1252 != v783 {
		v1211 = v1252
		v1216 = v1645
		v1218 = v1616
		v1222 = v1222 - int32(1)
		goto L180
	} else {
		goto L219
	}
L183:
	;
	v1614 = v1216
	v1615 = v1256
	v1616 = v1218
	v1617 = int64(0)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v1270 = v1265 >> (uint(int32(31)) % 32)
	v1274 = v1218 + base.I64_extend_i32_u(v1265^v1270-v1270)
	if int64(92233720369) <= v1274 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1277 = int64(0)
	v1279 = v1211 + (v799 - int32(2))
	if v1279 < v1204 {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v1413 = v1265
	v1414 = v1216
	v1415 = v1256
	v1416 = v1274
	goto L188
L188:
	;
	v1443 = base.I64_extend_i32_s(v1413)
	v1444 = v800 - v1211
	if v799 < v1444 {
		goto L204
	} else {
		goto L205
	}
L189:
	;
	v1281 = v1279
	goto L191
L190:
	;
	v1281 = v1204
	goto L191
L191:
	;
	if v1211 < v1281 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1288 = v1281
	v1291 = v1277
	goto L195
L193:
	;
	v1357 = v1216
	v1358 = v1256
	v1359 = v1277
	v1384 = v1257
	goto L194
L194:
	;
	v1386 = v1357 + v1359
	*(*int64)(unsafe.Add(mBase, uint32(v1247))) = v1386
	v1392 = base.F64_mul(v1202, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v1386), float64(1e+08)), v1384))
	v1396 = int32(0)
	v1398 = base.I32_trunc_sat_f64_s(v1392) - base.B2i32(base.F64_ge(v1392, float64(0)) == v1396)
	v1400 = v1398 >> (uint(int32(31)) % 32)
	v1405 = base.I64_extend_i32_u(v1398 ^ v1400 - v1400 + int32(1))
	if v1398 == v1396 {
		v1614 = v1386
		v1615 = v1358
		v1616 = v1405
		v1617 = v1277
		goto L182
	} else {
		goto L203
	}
L195:
	;
	v1320 = v811 + v1288<<(uint(int32(3))%32)
	v1321 = *(*int64)(unsafe.Add(mBase, uint32(v1320)))
	v1322 = v1321 + v1291
	if v1322 < int64(0) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v1348 = *(*int64)(unsafe.Add(mBase, uint32(v1247)))
	v1349 = *(*int64)(unsafe.Add(mBase, uint32(v1255)))
	v1357 = v1348
	v1358 = v1349
	v1359 = v1343
	v1384 = base.F64_convert_i64_s(v1349)
	goto L194
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1320))) = v1342
	v1346 = v1288 - int32(1)
	if v1211 < v1346 {
		v1288 = v1346
		v1291 = v1343
		goto L195
	} else {
		goto L202
	}
L198:
	;
	v1325 = int64(-1)
	v1328 = base.I64_div_u_s(v1322^v1325, int64(100000000))
	v1330 = v1328 ^ v1325
	v1342 = v1330*int64(-100000000) + v1322
	v1343 = v1330
	goto L197
L199:
	;
	goto L200
L200:
	;
	if base.Ui64(v1322) < base.Ui64(int64(100000000)) {
		v1342 = v1322
		v1343 = int64(0)
		goto L197
	} else {
		goto L201
	}
L201:
	;
	v1338 = base.I64_div_u_s(v1322, int64(100000000))
	v1342 = v1338*int64(-100000000) + v1322
	v1343 = v1338
	goto L197
L202:
	;
	goto L196
L203:
	;
	v1413 = v1398
	v1414 = v1386
	v1415 = v1358
	v1416 = v1405
	goto L188
L204:
	;
	v1446 = v799
	goto L206
L205:
	;
	v1446 = v1444
	goto L206
L206:
	;
	if v1446 <= int32(0) {
		v1614 = v1414
		v1615 = v1415
		v1616 = v1416
		v1617 = v1443
		goto L182
	} else {
		goto L207
	}
L207:
	;
	v1449 = int32(0)
	if v799 < v1222 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v1606 = *(*int64)(unsafe.Add(mBase, uint32(v1255)))
	v1607 = *(*int64)(unsafe.Add(mBase, uint32(v1247)))
	v1614 = v1607
	v1615 = v1606
	v1616 = v1416
	v1617 = v1443
	goto L182
L209:
	;
	v1451 = v799
	goto L211
L210:
	;
	v1451 = v1222
	goto L211
L211:
	;
	if v1451 != int32(1) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1464 = v1449
	v1470 = int32(0)
	goto L215
L213:
	;
	v1530 = v1449
	goto L214
L214:
	;
	v1562 = v1247 + v1530<<(uint(int32(3))%32)
	v1563 = *(*int64)(unsafe.Add(mBase, uint32(v1562)))
	v1567 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1530<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1562))) = v1563 - v1567*v1443
	goto L208
L215:
	;
	v1494 = int32(3)
	v1496 = v1247 + v1464<<(uint(v1494)%32)
	v1497 = *(*int64)(unsafe.Add(mBase, uint32(v1496)))
	v1498 = int32(2)
	v1501 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1464<<(uint(v1498)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1496))) = v1497 - v1501*v1443
	v1506 = v1464 | int32(1)
	v1509 = v1247 + v1506<<(uint(v1494)%32)
	v1510 = *(*int64)(unsafe.Add(mBase, uint32(v1509)))
	v1514 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1506<<(uint(v1498)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1509))) = v1510 - v1514*v1443
	v1519 = v1464 + v1498
	v1521 = v1470 + v1498
	if v1521 != v1451&int32(-2) {
		v1464 = v1519
		v1470 = v1521
		goto L215
	} else {
		goto L217
	}
L216:
	;
	if v1451&int32(1) == int32(0) {
		goto L208
	} else {
		goto L218
	}
L217:
	;
	goto L216
L218:
	;
	v1530 = v1519
	goto L214
L219:
	;
	goto L181
L220:
	;
	v1690 = v811 + v1656<<(uint(int32(3))%32)
	if v799 <= int32(1) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v1809 = v1690 - int32(8)
	v1818 = v1777
	goto L234
L222:
	;
	v1693 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1690))) = v1693
	v1777 = v1693
	goto L221
L223:
	;
	goto L224
L224:
	;
	v1707 = int64(0)
	v1710 = v799 - int32(2)
	goto L225
L225:
	;
	v1736 = v1690 + v1710<<(uint(int32(3))%32)
	v1737 = *(*int64)(unsafe.Add(mBase, uint32(v1736)))
	v1738 = v1737 + v1707
	if v1738 < int64(0) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1690))) = v1759
	v1766 = int64(0)
	if v1759 < v1766 {
		goto L3
	} else {
		goto L233
	}
L227:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1736)+8)) = v1758
	if int32(0) < v1710 {
		v1707 = v1759
		v1710 = v1710 - int32(1)
		goto L225
	} else {
		goto L232
	}
L228:
	;
	v1741 = int64(-1)
	v1744 = base.I64_div_u_s(v1738^v1741, int64(100000000))
	v1746 = v1744 ^ v1741
	v1758 = v1746*int64(-100000000) + v1738
	v1759 = v1746
	goto L227
L229:
	;
	goto L230
L230:
	;
	if base.Ui64(v1738) < base.Ui64(int64(100000000)) {
		v1758 = v1738
		v1759 = int64(0)
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v1754 = base.I64_div_u_s(v1738, int64(100000000))
	v1758 = v1754*int64(-100000000) + v1738
	v1759 = v1754
	goto L227
L232:
	;
	goto L226
L233:
	;
	v1777 = v1759
	goto L221
L234:
	;
	v1845 = int32(0)
	if v799 <= v1845 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v2109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010))))
	v2110 = v2079 + v2081 - v2109
	*(*int64)(unsafe.Add(mBase, uint32(v1690))) = v2110
	v2112 = *(*int64)(unsafe.Add(mBase, uint32(v1809)))
	*(*int64)(unsafe.Add(mBase, uint32(v1809))) = v2112 + int64(1)
	v1818 = v2110
	goto L234
L237:
	;
	v2079 = int64(0)
	v2081 = v1818
	goto L236
L238:
	;
	goto L239
L239:
	;
	v1854 = v1845
	goto L240
L240:
	;
	v1887 = *(*int64)(unsafe.Add(mBase, uint32(v1690+v1854<<(uint(int32(3))%32))))
	v1891 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1854<<(uint(int32(2))%32)))))
	if v1887 < v1891 {
		goto L2
	} else {
		goto L242
	}
L241:
	;
	v1898 = int64(0)
	if v799 < int32(2) {
		v2079 = v1898
		v2081 = v1818
		goto L236
	} else {
		goto L247
	}
L242:
	;
	if v1887 <= v1891 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1895 = v1854 + int32(1)
	if v1895 < v799 {
		v1854 = v1895
		goto L240
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	goto L241
L246:
	;
	goto L245
L247:
	;
	if v799 != int32(2) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v2072 = *(*int64)(unsafe.Add(mBase, uint32(v1690)))
	v2079 = v2043
	v2081 = v2072
	goto L236
L249:
	;
	v1909 = v1012
	v1910 = v1898
	v1915 = int32(0)
	goto L252
L250:
	;
	v1989 = v1012
	v1990 = v1898
	goto L251
L251:
	;
	v2021 = v1690 + v1989<<(uint(int32(3))%32)
	v2022 = *(*int64)(unsafe.Add(mBase, uint32(v2021)))
	v2026 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1989<<(uint(int32(2))%32)))))
	v2028 = v2022 - v2026 + v1990
	if v2028 < int64(0) {
		goto L262
	} else {
		goto L263
	}
L252:
	;
	v1941 = v1690 + v1909<<(uint(int32(3))%32)
	v1942 = *(*int64)(unsafe.Add(mBase, uint32(v1941)))
	v1946 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1909<<(uint(int32(2))%32)))))
	v1948 = v1942 - v1946 + v1910
	if v1948 < int64(0) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	if v1012&int32(1) == int32(0) {
		v2043 = v1976
		goto L248
	} else {
		goto L261
	}
L254:
	;
	v1953 = v1948 + int64(100000000)
	goto L256
L255:
	;
	v1953 = v1948
	goto L256
L256:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1941))) = v1953
	v1956 = v1909 - int32(1)
	v1959 = v1690 + v1956<<(uint(int32(3))%32)
	v1960 = *(*int64)(unsafe.Add(mBase, uint32(v1959)))
	v1964 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v1956<<(uint(int32(2))%32)))))
	v1968 = v1960 - v1964 + v1948>>(uint(int64(63))%64)
	if v1968 < int64(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1973 = v1968 + int64(100000000)
	goto L259
L258:
	;
	v1973 = v1968
	goto L259
L259:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1959))) = v1973
	v1976 = v1968 >> (uint(int64(63)) % 64)
	v1977 = int32(2)
	v1978 = v1909 - v1977
	v1980 = v1915 + v1977
	if v1980 != v1012&int32(-2) {
		v1909 = v1978
		v1910 = v1976
		v1915 = v1980
		goto L252
	} else {
		goto L260
	}
L260:
	;
	goto L253
L261:
	;
	v1989 = v1978
	v1990 = v1976
	goto L251
L262:
	;
	v2033 = v2028 + int64(100000000)
	goto L264
L263:
	;
	v2033 = v2028
	goto L264
L264:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2021))) = v2033
	v2043 = v2028 >> (uint(int64(63)) % 64)
	goto L248
L265:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L17
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_div_var_6), int32(0))
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L17
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_div_var_7), int32(_a_F_div_var_8), int32(_a_F_div_var_9))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L17
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L17
	} else {
		goto L270
	}
L270:
	;
	F_errmsg(m, int32(_a_F_div_var_6), int32(0))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L17
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_div_var_7), int32(_a_F_div_var_10), int32(_a_F_div_var_11))
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L17
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	v2189 = v1690 + v2157<<(uint(int32(3))%32)
	v2193 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010+v2157<<(uint(int32(2))%32)))))
	v2194 = *(*int64)(unsafe.Add(mBase, uint32(v2189)))
	v2196 = v2193 + (v2194 + v2158)
	v2200 = base.B2i32(int64(99999999) < v2196)
	if int64(99999999) < v2196 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	goto L2
L275:
	;
	v2201 = v2196 - int64(100000000)
	goto L277
L276:
	;
	v2201 = v2196
	goto L277
L277:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2189))) = v2201
	v2203 = base.I64_extend_i32_u(v2200)
	v2204 = int32(1)
	if base.Ui32(v2204) < base.Ui32(v2157) {
		v2157 = v2157 - v2204
		v2158 = v2203
		goto L273
	} else {
		goto L278
	}
L278:
	;
	v2208 = *(*int64)(unsafe.Add(mBase, uint32(v1690)))
	v2209 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1010))))
	v2211 = v2208 + (v2203 + v2209)
	*(*int64)(unsafe.Add(mBase, uint32(v1690))) = v2211
	v2213 = *(*int64)(unsafe.Add(mBase, uint32(v2151)))
	*(*int64)(unsafe.Add(mBase, uint32(v2151))) = v2213 - int64(1)
	v2217 = int64(0)
	if v2211 < v2217 {
		v2157 = v1012
		v2158 = v2217
		goto L273
	} else {
		goto L279
	}
L279:
	;
	goto L274
L280:
	;
	F_pfree(m, v2255)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L17
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	if v802 != v803 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	goto L282
L284:
	;
	v2263 = int32(_a_F_div_var_1)
	goto L286
L285:
	;
	v2263 = int32(0)
	goto L286
L286:
	;
	v2265 = v770 + int32(1)
	v2266 = int32(2)
	v2270 = F_palloc(m, v783<<(uint(v2266)%32)|v2266)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L17
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v2270
	v2273 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2270))) = uint16(v2273)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v783 << (uint(int32(1)) % 32)
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2277 = int32(2)
	v2278 = v2276 + v2277
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2278
	if v2277 <= v781 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v2292 = int64(0)
	v2300 = v783
	goto L291
L289:
	;
	goto L290
L290:
	;
	F_pfree(m, v811)
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L17
	} else {
		goto L299
	}
L291:
	;
	v2319 = v2300 - int32(1)
	v2323 = *(*int64)(unsafe.Add(mBase, uint32(v811+v2319<<(uint(int32(3))%32))))
	v2324 = v2323 + v2292
	if v2324 < int64(0) {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	goto L290
L293:
	;
	v2348 = v2278 + v2319<<(uint(int32(2))%32)
	v2349 = base.I32_wrap_i64(v2344)
	v2350 = int32(_a_F_div_var_0)
	v2351 = base.I32_div_u_s(v2349, v2350)
	*(*uint16)(unsafe.Add(mBase, uint32(v2348))) = uint16(v2351)
	v2355 = v2349 - v2351*v2350
	*(*uint16)(unsafe.Add(mBase, uint32(v2348)+2)) = uint16(v2355)
	if int32(1) < v2300 {
		v2292 = v2345
		v2300 = v2319
		goto L291
	} else {
		goto L298
	}
L294:
	;
	v2327 = int64(-1)
	v2330 = base.I64_div_u_s(v2324^v2327, int64(100000000))
	v2332 = v2330 ^ v2327
	v2344 = v2332*int64(-100000000) + v2324
	v2345 = v2332
	goto L293
L295:
	;
	goto L296
L296:
	;
	if base.Ui64(v2324) < base.Ui64(int64(100000000)) {
		v2344 = v2324
		v2345 = int64(0)
		goto L293
	} else {
		goto L297
	}
L297:
	;
	v2340 = base.I64_div_u_s(v2324, int64(100000000))
	v2344 = v2340*int64(-100000000) + v2324
	v2345 = v2340
	goto L293
L298:
	;
	goto L292
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2265
	if l4 != 0 {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2702
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2697
	goto L1
L301:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v2697 = v2659
	v2702 = int32(0)
	goto L300
L302:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if int32(0) < v2561 {
		goto L340
	} else {
		goto L341
	}
L303:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2561 = v2558
	goto L302
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2407 = l3 + v2404<<(uint(int32(2))%32)
	if v2407+int32(4) < int32(0) {
		goto L308
	} else {
		goto L309
	}
L305:
	;
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v2525 = l3 + v2265<<(uint(int32(2))%32)
	if v2525+int32(4) <= int32(0) {
		goto L333
	} else {
		goto L334
	}
L307:
	;
	goto L303
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	goto L307
L309:
	;
	goto L310
L310:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2418 = l3 & int32(3)
	v2422 = base.I32_div_s(v2407+int32(7), int32(4))
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2423 <= v2422 {
		goto L315
	} else {
		goto L316
	}
L311:
	;
	goto L307
L312:
	;
	if int32(0) <= v2488 {
		goto L311
	} else {
		goto L332
	}
L313:
	;
	v2468 = v2462
	goto L326
L314:
	;
	v2437 = int32(1)
	v2438 = v2422 - v2437
	v2441 = v2416 + v2438<<(uint(v2437)%32)
	v2442 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2441))))
	v2443 = int32(2)
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(v2418<<(uint(v2443)%32))+uint32(_c_F_div_var[0])))
	v2446 = base.I32_rem_s(v2442, v2445)
	v2447 = v2442 - v2446
	*(*uint16)(unsafe.Add(mBase, uint32(v2441))) = uint16(v2447)
	v2450 = base.I32_div_s(v2445, v2443)
	if v2446 < v2450 {
		v2488 = v2438
		goto L312
	} else {
		goto L321
	}
L315:
	;
	if base.B2i32(v2418 == int32(0))|base.B2i32(v2422 != v2423) != 0 {
		goto L311
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2422
	if v2418 != 0 {
		goto L314
	} else {
		goto L319
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2422
	goto L314
L319:
	;
	v2434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2416+v2422<<(uint(int32(1))%32)))))
	if v2434 <= int32(_a_F_div_var_2) {
		v2488 = v2422
		goto L312
	} else {
		goto L320
	}
L320:
	;
	v2462 = v2422
	goto L313
L321:
	;
	v2453 = v2445 + base.I32_extend16_s(v2447)
	if int32(_a_F_div_var_3) < v2453 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v2458 = v2453 + int32(_a_F_div_var_4)
	goto L324
L323:
	;
	v2458 = v2453
	goto L324
L324:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2441))) = uint16(v2458)
	if v2453 < int32(_a_F_div_var_0) {
		v2488 = v2438
		goto L312
	} else {
		goto L325
	}
L325:
	;
	v2462 = v2438
	goto L313
L326:
	;
	v2474 = int32(1)
	v2475 = v2468 - v2474
	v2478 = v2416 + v2475<<(uint(v2474)%32)
	v2481 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2478))))
	v2483 = base.B2i32(int32(_a_F_div_var_5) < v2481)
	if int32(_a_F_div_var_5) < v2481 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v2488 = v2475
	goto L312
L328:
	;
	v2484 = int32(-9999)
	goto L330
L329:
	;
	v2484 = v2474
	goto L330
L330:
	;
	v2485 = v2484 + v2481
	*(*uint16)(unsafe.Add(mBase, uint32(v2478))) = uint16(v2485)
	if int32(_a_F_div_var_5) < v2481 {
		v2468 = v2475
		goto L326
	} else {
		goto L331
	}
L331:
	;
	goto L327
L332:
	;
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2496 - int32(2)
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2501 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2500 + v2501
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2504 + v2501
	goto L311
L333:
	;
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2659 = v2530
	goto L301
L334:
	;
	goto L335
L335:
	;
	v2534 = base.I32_div_s(v2525+int32(7), int32(4))
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2535 < v2534 {
		v2561 = v2535
		goto L302
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2534
	v2539 = l3 & int32(3)
	if v2539 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v2561 = v2534
	goto L302
L338:
	;
	goto L339
L339:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2546 = int32(2)
	v2547 = v2542 + v2534<<(uint(int32(1))%32) - v2546
	v2548 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2547))))
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(v2539<<(uint(v2546)%32))+uint32(_c_F_div_var[0])))
	v2552 = base.I32_rem_s(v2548, v2551)
	v2553 = v2548 - v2552
	*(*uint16)(unsafe.Add(mBase, uint32(v2547))) = uint16(v2553)
	goto L303
L340:
	;
	v2568 = v2562
	v2573 = v2561
	goto L343
L341:
	;
	goto L342
L342:
	;
	if v2561 != 0 {
		v2697 = v2562
		v2702 = v2561
		goto L300
	} else {
		goto L353
	}
L343:
	;
	v2603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2568))))
	if v2603 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v2659 = v2562 + v2561<<(uint(int32(1))%32)
	goto L301
L345:
	;
	v2609 = v2573
	goto L348
L346:
	;
	goto L347
L347:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2650 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2649 - v2650
	if v2650 < v2573 {
		v2568 = v2568 + int32(2)
		v2573 = v2573 - v2650
		goto L343
	} else {
		goto L352
	}
L348:
	;
	v2644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2568+v2609<<(uint(int32(1))%32)-int32(2)))))
	if v2644 != 0 {
		v2697 = v2568
		v2702 = v2609
		goto L300
	} else {
		goto L350
	}
L350:
	;
	v2645 = int32(1)
	if v2645 < v2609 {
		v2609 = v2609 - v2645
		goto L348
	} else {
		goto L351
	}
L351:
	;
	v2659 = v2568
	goto L301
L352:
	;
	goto L344
L353:
	;
	v2659 = v2562
	goto L301
}
func F_doDeletion(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
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
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
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
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int64
	_ = v446
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
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
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1866 int32
	_ = v1866
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1967 int32
	_ = v1967
	var v1979 int32
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2057 int32
	_ = v2057
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2176 int32
	_ = v2176
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2223 int32
	_ = v2223
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2271 int32
	_ = v2271
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2328 int32
	_ = v2328
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2358 int32
	_ = v2358
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2378 int32
	_ = v2378
	var v2383 int32
	_ = v2383
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 <= int32(2752) {
		goto L23
	} else {
		goto L24
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L41
	} else {
		goto L733
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L41
	} else {
		goto L729
	}
L3:
	;
	m.G0 = v15 + int32(96)
	return
L4:
	;
	v2277 = F_get_object_catcache_oid(m, v17)
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L41
	} else {
		goto L710
	}
L5:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2194 = m.G0
	v2196 = v2194 - int32(32)
	m.G0 = v2196
	v2200 = F_table_open(m, int32(1255), int32(3))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L41
	} else {
		goto L684
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L41
	} else {
		goto L681
	}
L7:
	;
	if v17 == int32(2328) {
		goto L4
	} else {
		goto L680
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L41
	} else {
		goto L677
	}
L9:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2120 = m.G0
	v2122 = v2120 - int32(16)
	m.G0 = v2122
	v2126 = F_table_open(m, int32(_a_F_doDeletion_0), int32(3))
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L41
	} else {
		goto L661
	}
L10:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2030 = m.G0
	v2032 = v2030 - int32(16)
	m.G0 = v2032
	v2036 = F_table_open(m, int32(_a_F_doDeletion_1), int32(3))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L41
	} else {
		goto L637
	}
L11:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1941 = m.G0
	v1943 = v1941 - int32(16)
	m.G0 = v1943
	v1947 = F_table_open(m, int32(_a_F_doDeletion_2), int32(3))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L41
	} else {
		goto L613
	}
L12:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1884 = m.G0
	v1886 = v1884 + int32(-64)
	m.G0 = v1886
	v1889 = *(*int32)(unsafe.Add(mBase, _c_F_doDeletion[0]))
	if v1889 != v1883 {
		goto L595
	} else {
		goto L596
	}
L13:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1807 = m.G0
	v1809 = v1807 + int32(-64)
	m.G0 = v1809
	v1813 = F_table_open(m, int32(3602), int32(3))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L41
	} else {
		goto L570
	}
L14:
	;
	if v17 != int32(3381) {
		goto L6
	} else {
		goto L538
	}
L15:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1577 = m.G0
	v1579 = v1577 - int32(96)
	m.G0 = v1579
	v1583 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L41
	} else {
		goto L501
	}
L16:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1479 = m.G0
	v1481 = v1479 - int32(80)
	m.G0 = v1481
	v1485 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L41
	} else {
		goto L471
	}
L17:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1404 = m.G0
	v1406 = v1404 - int32(32)
	m.G0 = v1406
	v1410 = F_table_open(m, int32(2617), int32(3))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L41
	} else {
		goto L445
	}
L18:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1314 = m.G0
	v1316 = v1314 + int32(-64)
	m.G0 = v1316
	v1320 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L41
	} else {
		goto L418
	}
L19:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1217 = m.G0
	v1219 = v1217 - int32(80)
	m.G0 = v1219
	v1223 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L41
	} else {
		goto L393
	}
L20:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1077 = m.G0
	v1079 = v1077 + int32(-64)
	m.G0 = v1079
	v1083 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L41
	} else {
		goto L352
	}
L21:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v883 = m.G0
	v885 = v883 - int32(16)
	m.G0 = v885
	v889 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L41
	} else {
		goto L306
	}
L22:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v178 = F_get_rel_relkind(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L41
	} else {
		goto L84
	}
L23:
	;
	if v17 <= int32(2327) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v17 <= int32(3575) {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	switch v17 - int32(1213) {
	case 0, 47, 49:
		goto L8
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
		goto L6
	case 34:
		goto L21
	case 42:
		goto L5
	case 46:
		goto L22
	case 48:
		goto L4
	default:
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	switch v17 - int32(2601) {
	case 0, 1, 2, 4, 6, 11, 14, 15:
		goto L4
	case 3:
		goto L19
	case 5:
		goto L20
	case 7, 8, 9, 10, 13, 18:
		goto L6
	case 12:
		goto L18
	case 16:
		goto L17
	case 17:
		goto L16
	case 19:
		goto L15
	default:
		goto L7
	}
L29:
	;
	if base.Ui32(v17-int32(1417)) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v17 != int32(826) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L4
L32:
	;
	if v17 <= int32(3380) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v17 <= int32(3763) {
		goto L78
	} else {
		goto L79
	}
L35:
	;
	if v17 == int32(2753) {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	switch v17 - int32(3456) {
	case 0, 10:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L6
	default:
		goto L14
	}
L38:
	;
	if v17 == int32(3079) {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	if v17 != int32(3256) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = m.G0
	v45 = v43 - int32(96)
	m.G0 = v45
	v49 = F_table_open(m, int32(3256), int32(3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return
L42:
	;
	v52 = v45 + int32(48)
	F_ScanKeyInit(m, v52, int32(1), int32(3), int32(184), v42)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v59 = int32(1)
	v62 = F_systable_beginscan(m, v49, int32(3257), v59, int32(0), v59, v52)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L3
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L41
	} else {
		goto L74
	}
L46:
	;
	v64 = F_systable_getnext(m, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	if v64 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v67)+68))
	v71 = F_table_open(m, v69, int32(8))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L41
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L41
	} else {
		goto L71
	}
L51:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_doDeletion[1])))
	if v99 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L41
	} else {
		goto L54
	}
L53:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+119)))
	switch v74 - int32(112) {
	case 0, 2:
		goto L51
	default:
		goto L52
	}
L54:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L41
	} else {
		goto L55
	}
L55:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v84 + int32(4)
	F_errmsg(m, int32(_a_F_doDeletion_3), v45+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L41
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_doDeletion_4), int32(374), int32(_a_F_doDeletion_5))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L41
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
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	if base.Ui32(v104) < base.Ui32(int32(_a_F_doDeletion_6)) {
		v113 = v103
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	F_simple_heap_delete(m, v49, v64+int32(4))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L41
	} else {
		goto L66
	}
L61:
	;
	if v113 != 0 {
		goto L45
	} else {
		goto L65
	}
L62:
	;
	goto L61
L63:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+68))
	if v108 == int32(99) {
		v113 = v103
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v111 = F_isTempToastNamespace(m, v108)
	mBase = m.M
	v113 = v111
	goto L62
L65:
	;
	goto L60
L66:
	;
	F_systable_endscan(m, v62)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L41
	} else {
		goto L67
	}
L67:
	;
	F_CacheInvalidateRelcache(m, v71)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L41
	} else {
		goto L68
	}
L68:
	;
	F_relation_close(m, v71, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L41
	} else {
		goto L69
	}
L69:
	;
	F_relation_close(m, v49, int32(3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L41
	} else {
		goto L70
	}
L70:
	;
	m.G0 = v45 + int32(96)
	goto L44
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v42
	F_errmsg_internal(m, int32(_a_F_doDeletion_7), v45)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_doDeletion_4), int32(358), int32(_a_F_doDeletion_5))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L41
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
	F_errcode(m, int32(16797828))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L41
	} else {
		goto L75
	}
L75:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = v151 + int32(4)
	F_errmsg(m, int32(_a_F_doDeletion_8), v45+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L41
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_doDeletion_4), int32(380), int32(_a_F_doDeletion_5))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L41
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	switch v17 - int32(3576) {
	case 0, 24, 25:
		goto L4
	default:
		goto L6
	case 26:
		goto L13
	}
L79:
	;
	goto L80
L80:
	;
	switch v17 - int32(_a_F_doDeletion_9) {
	case 0:
		goto L8
	case 1, 2, 3, 5:
		goto L6
	case 4:
		goto L9
	case 6:
		goto L10
	default:
		goto L81
	}
L81:
	;
	if v17 == int32(3764) {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	switch v17 - int32(_a_F_doDeletion_2) {
	case 0:
		goto L11
	default:
		goto L6
	case 6:
		goto L8
	}
L83:
	;
	if v178 != int32(83) {
		goto L3
	} else {
		goto L294
	}
L84:
	;
	if v178&int32(-33) == int32(73) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v188 = int32(base.Ui32(l1&int32(2)) >> (uint(int32(1)) % 32))
	v193 = m.G0
	v195 = v193 - int32(96)
	m.G0 = v195
	v198 = F_SearchSysCache1(m, int32(34), v184)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L41
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v433 != 0 {
		goto L175
	} else {
		goto L176
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L41
	} else {
		goto L172
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L41
	} else {
		goto L168
	}
L90:
	;
	if v198 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+22)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v200+v201)+4))
	F_ReleaseCatCache(m, v198)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L41
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L41
	} else {
		goto L165
	}
L94:
	;
	v206 = int32(4)
	if int32(base.Ui32(l1&int32(32))>>(uint(int32(5))%32)) != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v209 = v206
	goto L97
L96:
	;
	v209 = int32(8)
	goto L97
L97:
	;
	if v188 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v210 = v206
	goto L100
L99:
	;
	v210 = v209
	goto L100
L100:
	;
	v211 = F_table_open(m, v203, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L41
	} else {
		goto L101
	}
L101:
	;
	v213 = F_index_open(m, v184, v210)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L41
	} else {
		goto L102
	}
L102:
	;
	F_CheckTableNotInUse(m, v213, int32(_a_F_doDeletion_10))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L41
	} else {
		goto L103
	}
L103:
	;
	if v188 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)+48))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+119)))
	switch v316 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L136
	default:
		goto L135
	}
L105:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_doDeletion[2]))
	if v219 != 0 {
		goto L89
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_TransferPredicateLocksToHeapRelation(m, v213)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L41
	} else {
		goto L134
	}
L108:
	;
	F_index_set_state_flags(m, v184, int32(2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L41
	} else {
		goto L109
	}
L109:
	;
	F_CacheInvalidateRelcache(m, v211)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L41
	} else {
		goto L110
	}
L110:
	;
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v211)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v195)+72)) = int64(72057594037927936)
	*(*uint32)(unsafe.Add(mBase, uint32(v195)+68)) = uint32(v225)
	*(*int64)(unsafe.Add(mBase, uint32(v195)+88)) = v225
	v231 = int64(base.Ui64(v225) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v195)+64)) = uint32(v231)
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v213)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v195)+80)) = v233
	F_relation_close(m, v211, int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L41
	} else {
		goto L111
	}
L111:
	;
	F_relation_close(m, v213, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L41
	} else {
		goto L112
	}
L112:
	;
	F_LockRelationIdForSession(m, v195+int32(88), int32(4))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L41
	} else {
		goto L113
	}
L113:
	;
	F_LockRelationIdForSession(m, v195+int32(80), int32(4))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L41
	} else {
		goto L114
	}
L114:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L41
	} else {
		goto L115
	}
L115:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L41
	} else {
		goto L116
	}
L116:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L41
	} else {
		goto L117
	}
L117:
	;
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v195)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v195)+56)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v195)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v195)+48)) = v259
	F_WaitForLockers(m, v195+int32(48), int32(8))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L41
	} else {
		goto L118
	}
L118:
	;
	v266 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L41
	} else {
		goto L119
	}
L119:
	;
	F_PushActiveSnapshot(m, v266)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L41
	} else {
		goto L120
	}
L120:
	;
	v271 = F_table_open(m, v203, int32(4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L41
	} else {
		goto L121
	}
L121:
	;
	v274 = F_index_open(m, v184, int32(4))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L41
	} else {
		goto L122
	}
L122:
	;
	F_TransferPredicateLocksToHeapRelation(m, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L41
	} else {
		goto L123
	}
L123:
	;
	F_index_set_state_flags(m, v184, int32(3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L41
	} else {
		goto L124
	}
L124:
	;
	F_CacheInvalidateRelcache(m, v271)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L41
	} else {
		goto L125
	}
L125:
	;
	F_relation_close(m, v271, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L41
	} else {
		goto L126
	}
L126:
	;
	F_relation_close(m, v274, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L41
	} else {
		goto L127
	}
L127:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L41
	} else {
		goto L128
	}
L128:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L41
	} else {
		goto L129
	}
L129:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L41
	} else {
		goto L130
	}
L130:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v195)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v195)+40)) = v295
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v195)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v195)+32)) = v297
	F_WaitForLockers(m, v195+int32(32), int32(8))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L41
	} else {
		goto L131
	}
L131:
	;
	v305 = F_table_open(m, v203, int32(4))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L41
	} else {
		goto L132
	}
L132:
	;
	v308 = F_index_open(m, v184, int32(8))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L41
	} else {
		goto L133
	}
L133:
	;
	v312 = v308
	v313 = v305
	goto L104
L134:
	;
	v312 = v213
	v313 = v211
	goto L104
L135:
	;
	F_pgstat_drop_relation(m, v312)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L41
	} else {
		goto L138
	}
L136:
	;
	F_RelationDropStorage(m, v312)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L41
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	F_relation_close(m, v312, int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L41
	} else {
		goto L139
	}
L139:
	;
	F_RelationForgetRelation(m, v184)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L41
	} else {
		goto L140
	}
L140:
	;
	v328 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L41
	} else {
		goto L141
	}
L141:
	;
	F_PushActiveSnapshot(m, v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L41
	} else {
		goto L142
	}
L142:
	;
	v334 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L41
	} else {
		goto L143
	}
L143:
	;
	v337 = F_SearchSysCache1(m, int32(34), v184)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L41
	} else {
		goto L144
	}
L144:
	;
	if v337 == int32(0) {
		goto L88
	} else {
		goto L145
	}
L145:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v334)+52))
	v343 = F_heap_attisnull(m, v337, int32(20), v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L41
	} else {
		goto L146
	}
L146:
	;
	F_simple_heap_delete(m, v334, v337+int32(4))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L41
	} else {
		goto L147
	}
L147:
	;
	F_ReleaseCatCache(m, v337)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L41
	} else {
		goto L148
	}
L148:
	;
	F_relation_close(m, v334, int32(3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L41
	} else {
		goto L149
	}
L149:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L41
	} else {
		goto L150
	}
L150:
	;
	if v343 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	F_RemoveStatistics(m, v184, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L41
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	F_DeleteAttributeTuples(m, v184)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L41
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	F_DeleteRelationTuple(m, v184)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L41
	} else {
		goto L156
	}
L156:
	;
	v365 = int32(0)
	v368 = F_DeleteInheritsTuple(m, v184, v365, v365, v365)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L41
	} else {
		goto L157
	}
L157:
	;
	F_CacheInvalidateRelcache(m, v313)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L41
	} else {
		goto L158
	}
L158:
	;
	F_relation_close(m, v313, int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L41
	} else {
		goto L159
	}
L159:
	;
	if v188 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	F_UnlockRelationIdForSession(m, v195+int32(88), int32(4))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L41
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	m.G0 = v195 + int32(96)
	goto L83
L163:
	;
	F_UnlockRelationIdForSession(m, v195+int32(80), int32(4))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L41
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v184
	F_errmsg_internal(m, int32(_a_F_doDeletion_11), v195)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L41
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_doDeletion_12), int32(3594), int32(_a_F_doDeletion_13))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L41
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L41
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(_a_F_doDeletion_14), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L41
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_doDeletion_12), int32(2222), int32(_a_F_doDeletion_15))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L41
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v184
	F_errmsg_internal(m, int32(_a_F_doDeletion_11), v195+int32(16))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L41
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_doDeletion_12), int32(2353), int32(_a_F_doDeletion_15))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L41
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	v434 = base.I32_extend16_s(v433)
	v435 = m.G0
	v437 = v435 - int32(272)
	m.G0 = v437
	v441 = int32(0)
	base.MemoryFill(m, v437+int32(96), v441, int32(100))
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+88)) = uint8(v441)
	v446 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v437)+80)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v437)+72)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v437)+64)) = v446
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+56)) = uint8(v441)
	*(*int64)(unsafe.Add(mBase, uint32(v437)+48)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v437)+40)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v437)+32)) = v446
	v461 = F_relation_open(m, v432, int32(8))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L41
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v545 = m.G0
	v547 = v545 - int32(80)
	m.G0 = v547
	v550 = F_SearchSysCache1(m, int32(57), v432)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L41
	} else {
		goto L196
	}
L178:
	;
	v465 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L41
	} else {
		goto L179
	}
L179:
	;
	v468 = F_SearchSysCacheCopy(m, int32(7), v432, v434)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L41
	} else {
		goto L180
	}
L180:
	;
	if v468 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L41
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v468)+16))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+22)))
	v488 = v486 + v487
	v489 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+86)) = uint8(v489)
	*(*int32)(unsafe.Add(mBase, uint32(v488)+68)) = v489
	v493 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v488)+90)) = uint16(v493)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+16)) = v434
	v497 = v437 + int32(208)
	v502 = F_pg_snprintf(m, v497, int32(64), int32(_a_F_doDeletion_16), v437+int32(16))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L41
	} else {
		goto L187
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v434
	F_errmsg_internal(m, int32(_a_F_doDeletion_17), v437)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L41
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_doDeletion_18), int32(1709), int32(_a_F_doDeletion_19))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L41
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	v507 = F_strncpy(m, v488+int32(4), v497, int32(64))
	mBase = m.M
	v508 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v507)+63)) = uint8(v508)
	goto L188
L188:
	;
	v510 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+88)) = uint8(v510)
	v512 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+56)) = uint8(v512)
	v514 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+85)) = v514
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+84)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+52)) = v514
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v465)+52))
	v527 = F_heap_modify_tuple(m, v468, v520, v437+int32(96), v437-int32(-64), v437+int32(32))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L41
	} else {
		goto L189
	}
L189:
	;
	F_CatalogTupleUpdate(m, v465, v527+int32(4), v527)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L41
	} else {
		goto L190
	}
L190:
	;
	F_relation_close(m, v465, int32(3))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L41
	} else {
		goto L191
	}
L191:
	;
	F_RemoveStatistics(m, v432, v434)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L41
	} else {
		goto L192
	}
L192:
	;
	F_relation_close(m, v461, int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L41
	} else {
		goto L193
	}
L193:
	;
	m.G0 = v437 + int32(272)
	goto L83
L194:
	;
	goto L83
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L41
	} else {
		goto L291
	}
L196:
	;
	if v550 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+22)))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552+v553)+131)))
	if v555 != int32(1) {
		v573 = int32(0)
		v574 = int32(0)
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L41
	} else {
		goto L288
	}
L200:
	;
	F_ReleaseCatCache(m, v550)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L41
	} else {
		goto L212
	}
L201:
	;
	v559 = F_get_partition_parent(m, v432, int32(1))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L41
	} else {
		goto L202
	}
L202:
	;
	F_LockRelationOid(m, v559, int32(8))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L41
	} else {
		goto L203
	}
L203:
	;
	v564 = F_get_default_partition_oid(m, v559)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L41
	} else {
		goto L204
	}
L204:
	;
	if v564 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v573 = int32(0)
	v574 = v559
	goto L200
L206:
	;
	goto L207
L207:
	;
	if v432 == v564 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v573 = v432
	v574 = v559
	goto L200
L209:
	;
	goto L210
L210:
	;
	F_LockRelationOid(m, v564, int32(8))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L41
	} else {
		goto L211
	}
L211:
	;
	v573 = v564
	v574 = v559
	goto L200
L212:
	;
	v578 = F_relation_open(m, v432, int32(8))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L41
	} else {
		goto L213
	}
L213:
	;
	F_CheckTableNotInUse(m, v578, int32(_a_F_doDeletion_20))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L41
	} else {
		goto L214
	}
L214:
	;
	F_CheckTableForSerializableConflictIn(m, v578)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L41
	} else {
		goto L215
	}
L215:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v578)+48))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+119)))
	if v586 == int32(102) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v591 = F_table_open(m, int32(3118), int32(3))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L41
	} else {
		goto L219
	}
L217:
	;
	v611 = v586
	goto L218
L218:
	;
	if v611&int32(255) == int32(112) {
		goto L225
	} else {
		goto L226
	}
L219:
	;
	v594 = F_SearchSysCache1(m, int32(33), v432)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L41
	} else {
		goto L220
	}
L220:
	;
	if v594 == int32(0) {
		goto L195
	} else {
		goto L221
	}
L221:
	;
	F_simple_heap_delete(m, v591, v594+int32(4))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L41
	} else {
		goto L222
	}
L222:
	;
	F_ReleaseCatCache(m, v594)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L41
	} else {
		goto L223
	}
L223:
	;
	F_relation_close(m, v591, int32(3))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L41
	} else {
		goto L224
	}
L224:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v578)+48))
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+119)))
	v611 = v608
	goto L218
L225:
	;
	v616 = m.G0
	v618 = v616 - int32(16)
	m.G0 = v618
	v622 = F_table_open(m, int32(3350), int32(3))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L41
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if v432 == v573 {
		goto L239
	} else {
		goto L240
	}
L228:
	;
	v625 = F_SearchSysCache1(m, int32(45), v432)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L41
	} else {
		goto L229
	}
L229:
	;
	if v625 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L41
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	F_simple_heap_delete(m, v622, v625+int32(4))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L41
	} else {
		goto L236
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v618))) = v432
	F_errmsg_internal(m, int32(_a_F_doDeletion_21), v618)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L41
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_doDeletion_18), int32(4032), int32(_a_F_doDeletion_22))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L41
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_ReleaseCatCache(m, v625)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L41
	} else {
		goto L237
	}
L237:
	;
	F_relation_close(m, v622, int32(3))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L41
	} else {
		goto L238
	}
L238:
	;
	m.G0 = v618 + int32(16)
	goto L227
L239:
	;
	F_update_default_partition_oid(m, v574, int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L41
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v578)+48))
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+119)))
	switch v662 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L244
	default:
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	F_pgstat_drop_relation(m, v578)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L41
	} else {
		goto L246
	}
L244:
	;
	F_RelationDropStorage(m, v578)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L41
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	F_relation_close(m, v578, int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L41
	} else {
		goto L247
	}
L247:
	;
	F_RemoveSubscriptionRel(m, int32(0), v432)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L41
	} else {
		goto L248
	}
L248:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_doDeletion[3]))
	if v676 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	F_RelationForgetRelation(m, v432)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L41
	} else {
		goto L262
	}
L250:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v679 <= int32(0) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v682 = int32(0)
	if v682 < v679 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v685 = v679
	goto L254
L253:
	;
	v685 = v682
	goto L254
L254:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v676)+12))
	v692 = int32(0)
	goto L255
L255:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v686+v692<<(uint(int32(2))%32))))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v703)))
	if v432 != v704 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_doDeletion[4]))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+8))
	goto L261
L257:
	;
	v707 = v692 + int32(1)
	if v685 != v707 {
		v692 = v707
		goto L255
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	goto L256
L260:
	;
	goto L249
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v703)+12)) = v711
	goto L249
L262:
	;
	v729 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L41
	} else {
		goto L263
	}
L263:
	;
	v732 = v547 + int32(32)
	F_ScanKeyInit(m, v732, int32(1), int32(3), int32(184), v432)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L41
	} else {
		goto L264
	}
L264:
	;
	v739 = int32(1)
	v742 = F_systable_beginscan(m, v729, int32(2680), v739, int32(0), v739, v732)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L41
	} else {
		goto L265
	}
L265:
	;
	v744 = F_systable_getnext(m, v742)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L41
	} else {
		goto L266
	}
L266:
	;
	if v744 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v750 = v744
	goto L270
L268:
	;
	goto L269
L269:
	;
	F_systable_endscan(m, v742)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L41
	} else {
		goto L275
	}
L270:
	;
	F_simple_heap_delete(m, v729, v750+int32(4))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L41
	} else {
		goto L272
	}
L271:
	;
	goto L269
L272:
	;
	v762 = F_systable_getnext(m, v742)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L41
	} else {
		goto L273
	}
L273:
	;
	if v762 != 0 {
		v750 = v762
		goto L270
	} else {
		goto L274
	}
L274:
	;
	goto L271
L275:
	;
	F_relation_close(m, v729, int32(3))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L41
	} else {
		goto L276
	}
L276:
	;
	F_RemoveStatistics(m, v432, int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L41
	} else {
		goto L277
	}
L277:
	;
	F_DeleteAttributeTuples(m, v432)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L41
	} else {
		goto L278
	}
L278:
	;
	F_DeleteRelationTuple(m, v432)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L41
	} else {
		goto L279
	}
L279:
	;
	if v574 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v788 = int32(0)
	if base.B2i32(v573 == v788)|base.B2i32(v432 == v573) == v788 {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	goto L282
L282:
	;
	m.G0 = v547 + int32(80)
	goto L194
L283:
	;
	F_CacheInvalidateRelcacheByRelid(m, v573)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L41
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	F_CacheInvalidateRelcacheByRelid(m, v574)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L41
	} else {
		goto L287
	}
L286:
	;
	goto L285
L287:
	;
	goto L282
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v432
	F_errmsg_internal(m, int32(_a_F_doDeletion_23), v547)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L41
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_doDeletion_18), int32(1803), int32(_a_F_doDeletion_24))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L41
	} else {
		goto L290
	}
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+16)) = v432
	F_errmsg_internal(m, int32(_a_F_doDeletion_25), v547+int32(16))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L41
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_doDeletion_18), int32(1857), int32(_a_F_doDeletion_24))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L41
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v844 = m.G0
	v846 = v844 - int32(16)
	m.G0 = v846
	v850 = F_table_open(m, int32(2224), int32(3))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L41
	} else {
		goto L295
	}
L295:
	;
	v853 = F_SearchSysCache1(m, int32(61), v843)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L41
	} else {
		goto L296
	}
L296:
	;
	if v853 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L41
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	F_simple_heap_delete(m, v850, v853+int32(4))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L41
	} else {
		goto L303
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v846))) = v843
	F_errmsg_internal(m, int32(_a_F_doDeletion_26), v846)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L41
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_doDeletion_27), int32(579), int32(_a_F_doDeletion_28))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L41
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_ReleaseCatCache(m, v853)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L41
	} else {
		goto L304
	}
L304:
	;
	F_relation_close(m, v850, int32(3))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L41
	} else {
		goto L305
	}
L305:
	;
	m.G0 = v846 + int32(16)
	goto L3
L306:
	;
	v892 = F_SearchSysCache1(m, int32(82), v882)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L41
	} else {
		goto L308
	}
L307:
	;
	goto L3
L308:
	;
	if v892 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	F_simple_heap_delete(m, v889, v892+int32(4))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L41
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L41
	} else {
		goto L349
	}
L312:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v892)+16))
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+22)))
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898+v899)+79)))
	if v901 == int32(101) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v904 = m.G0
	v906 = v904 - int32(48)
	m.G0 = v906
	v910 = F_table_open(m, int32(3501), int32(3))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L41
	} else {
		goto L316
	}
L314:
	;
	v979 = v901
	goto L315
L315:
	;
	if v979&int32(255) == int32(114) {
		goto L330
	} else {
		goto L331
	}
L316:
	;
	F_ScanKeyInit(m, v906, int32(2), int32(3), int32(184), v882)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L41
	} else {
		goto L317
	}
L317:
	;
	v918 = int32(1)
	v921 = F_systable_beginscan(m, v910, int32(3503), v918, int32(0), v918, v906)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L41
	} else {
		goto L318
	}
L318:
	;
	v923 = F_systable_getnext(m, v921)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L41
	} else {
		goto L319
	}
L319:
	;
	if v923 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v925 = v923
	goto L323
L321:
	;
	goto L322
L322:
	;
	F_systable_endscan(m, v921)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L41
	} else {
		goto L328
	}
L323:
	;
	F_simple_heap_delete(m, v910, v925+int32(4))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L41
	} else {
		goto L325
	}
L324:
	;
	goto L322
L325:
	;
	v941 = F_systable_getnext(m, v921)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L41
	} else {
		goto L326
	}
L326:
	;
	if v941 != 0 {
		v925 = v941
		goto L323
	} else {
		goto L327
	}
L327:
	;
	goto L324
L328:
	;
	F_relation_close(m, v910, int32(3))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L41
	} else {
		goto L329
	}
L329:
	;
	m.G0 = v906 + int32(48)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v892)+16))
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+22)))
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963+v964)+79)))
	v979 = v966
	goto L315
L330:
	;
	v984 = m.G0
	v986 = v984 - int32(48)
	m.G0 = v986
	v990 = F_table_open(m, int32(3541), int32(3))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L41
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	F_ReleaseCatCache(m, v892)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L41
	} else {
		goto L347
	}
L333:
	;
	F_ScanKeyInit(m, v986, int32(1), int32(3), int32(184), v882)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L41
	} else {
		goto L334
	}
L334:
	;
	v998 = int32(1)
	v1001 = F_systable_beginscan(m, v990, int32(3542), v998, int32(0), v998, v986)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L41
	} else {
		goto L335
	}
L335:
	;
	v1003 = F_systable_getnext(m, v1001)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L41
	} else {
		goto L336
	}
L336:
	;
	if v1003 != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1007 = v1003
	goto L340
L338:
	;
	goto L339
L339:
	;
	F_systable_endscan(m, v1001)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L41
	} else {
		goto L345
	}
L340:
	;
	F_simple_heap_delete(m, v990, v1007+int32(4))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L41
	} else {
		goto L342
	}
L341:
	;
	goto L339
L342:
	;
	v1021 = F_systable_getnext(m, v1001)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L41
	} else {
		goto L343
	}
L343:
	;
	if v1021 != 0 {
		v1007 = v1021
		goto L340
	} else {
		goto L344
	}
L344:
	;
	goto L341
L345:
	;
	F_relation_close(m, v990, int32(3))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L41
	} else {
		goto L346
	}
L346:
	;
	m.G0 = v986 + int32(48)
	goto L332
L347:
	;
	F_relation_close(m, v889, int32(3))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L41
	} else {
		goto L348
	}
L348:
	;
	m.G0 = v885 + int32(16)
	goto L307
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v885))) = v882
	F_errmsg_internal(m, int32(_a_F_doDeletion_29), v885)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L41
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_doDeletion_30), int32(666), int32(_a_F_doDeletion_31))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L41
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	v1086 = F_SearchSysCache1(m, int32(19), v1076)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L41
	} else {
		goto L357
	}
L353:
	;
	goto L3
L354:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L41
	} else {
		goto L390
	}
L355:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L41
	} else {
		goto L387
	}
L356:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L41
	} else {
		goto L384
	}
L357:
	;
	if v1086 != 0 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+16))
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1088)+22)))
	v1090 = v1088 + v1089
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+80))
	if v1091 != 0 {
		goto L362
	} else {
		goto L363
	}
L359:
	;
	goto L360
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L41
	} else {
		goto L381
	}
L361:
	;
	F_simple_heap_delete(m, v1083, v1086+int32(4))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L41
	} else {
		goto L378
	}
L362:
	;
	v1093 = F_table_open(m, v1091, int32(8))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L41
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+84))
	if v1134 == int32(0) {
		goto L354
	} else {
		goto L377
	}
L365:
	;
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090)+72)))
	if v1095 == int32(99) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1100 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L41
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	F_relation_close(m, v1093, int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L41
	} else {
		goto L376
	}
L369:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+80))
	v1105 = F_SearchSysCacheCopy(m, int32(57), v1103, int32(0))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L41
	} else {
		goto L370
	}
L370:
	;
	if v1105 == int32(0) {
		goto L356
	} else {
		goto L371
	}
L371:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+16))
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109)+22)))
	v1111 = v1109 + v1110
	v1112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1111)+122)))
	if v1112 == int32(0) {
		goto L355
	} else {
		goto L372
	}
L372:
	;
	v1116 = v1112 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1111)+122)) = uint16(v1116)
	F_CatalogTupleUpdate(m, v1100, v1105+int32(4), v1105)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L41
	} else {
		goto L373
	}
L373:
	;
	F_pfree(m, v1105)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L41
	} else {
		goto L374
	}
L374:
	;
	F_relation_close(m, v1100, int32(3))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L41
	} else {
		goto L375
	}
L375:
	;
	goto L368
L376:
	;
	goto L361
L377:
	;
	goto L361
L378:
	;
	F_ReleaseCatCache(m, v1086)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L41
	} else {
		goto L379
	}
L379:
	;
	F_relation_close(m, v1083, int32(3))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L41
	} else {
		goto L380
	}
L380:
	;
	m.G0 = v1079 - int32(-64)
	goto L353
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079))) = v1076
	F_errmsg_internal(m, int32(_a_F_doDeletion_32), v1079)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L41
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_doDeletion_33), int32(922), int32(_a_F_doDeletion_34))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L41
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+32)) = v1171
	F_errmsg_internal(m, int32(_a_F_doDeletion_23), v1077+int32(-32))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L41
	} else {
		goto L385
	}
L385:
	;
	F_errfinish(m, int32(_a_F_doDeletion_33), int32(954), int32(_a_F_doDeletion_34))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L41
	} else {
		goto L386
	}
L386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L387:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+48)) = v1187 + int32(4)
	F_errmsg_internal(m, int32(_a_F_doDeletion_35), v1077+int32(-16))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L41
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(_a_F_doDeletion_33), int32(959), int32(_a_F_doDeletion_34))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L41
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+16)) = v1076
	F_errmsg_internal(m, int32(_a_F_doDeletion_36), v1077+int32(-48))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L41
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_doDeletion_33), int32(982), int32(_a_F_doDeletion_34))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L41
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	v1226 = v1219 + int32(32)
	F_ScanKeyInit(m, v1226, int32(1), int32(3), int32(184), v1216)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L41
	} else {
		goto L394
	}
L394:
	;
	v1233 = int32(1)
	v1236 = F_systable_beginscan(m, v1223, int32(2657), v1233, int32(0), v1233, v1226)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L41
	} else {
		goto L397
	}
L395:
	;
	goto L3
L396:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L41
	} else {
		goto L415
	}
L397:
	;
	v1238 = F_systable_getnext(m, v1236)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L41
	} else {
		goto L398
	}
L398:
	;
	if v1238 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+16))
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240)+22)))
	v1242 = v1240 + v1241
	v1243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1242)+8)))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+4))
	v1246 = F_relation_open(m, v1244, int32(8))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L41
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L41
	} else {
		goto L412
	}
L402:
	;
	F_simple_heap_delete(m, v1223, v1238+int32(4))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L41
	} else {
		goto L403
	}
L403:
	;
	F_systable_endscan(m, v1236)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L41
	} else {
		goto L404
	}
L404:
	;
	F_relation_close(m, v1223, int32(3))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L41
	} else {
		goto L405
	}
L405:
	;
	v1259 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L41
	} else {
		goto L406
	}
L406:
	;
	v1262 = F_SearchSysCacheCopy(m, int32(7), v1244, v1243)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L41
	} else {
		goto L407
	}
L407:
	;
	if v1262 == int32(0) {
		goto L396
	} else {
		goto L408
	}
L408:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+16))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266)+22)))
	v1269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1266+v1267)+87)) = uint8(v1269)
	F_CatalogTupleUpdate(m, v1259, v1262+int32(4), v1262)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L41
	} else {
		goto L409
	}
L409:
	;
	F_relation_close(m, v1259, int32(3))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L41
	} else {
		goto L410
	}
L410:
	;
	F_relation_close(m, v1246, int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L41
	} else {
		goto L411
	}
L411:
	;
	m.G0 = v1219 + int32(80)
	goto L395
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1219))) = v1216
	F_errmsg_internal(m, int32(_a_F_doDeletion_37), v1219)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L41
	} else {
		goto L413
	}
L413:
	;
	F_errfinish(m, int32(_a_F_doDeletion_38), int32(232), int32(_a_F_doDeletion_39))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L41
	} else {
		goto L414
	}
L414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1219)+20)) = v1244
	*(*int32)(unsafe.Add(mBase, uint32(v1219)+16)) = v1243
	F_errmsg_internal(m, int32(_a_F_doDeletion_17), v1219+int32(16))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L41
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_doDeletion_38), int32(254), int32(_a_F_doDeletion_39))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L41
	} else {
		goto L417
	}
L417:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L418:
	;
	v1324 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L41
	} else {
		goto L419
	}
L419:
	;
	v1327 = v1314 + int32(-48)
	F_ScanKeyInit(m, v1327, int32(1), int32(3), int32(184), v1313)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L41
	} else {
		goto L420
	}
L420:
	;
	v1334 = int32(1)
	v1337 = F_systable_beginscan(m, v1320, int32(2996), v1334, int32(0), v1334, v1327)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L41
	} else {
		goto L422
	}
L421:
	;
	goto L3
L422:
	;
	v1339 = F_systable_getnext(m, v1337)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L41
	} else {
		goto L423
	}
L423:
	;
	if v1339 != 0 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	F_simple_heap_delete(m, v1320, v1339+int32(4))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L41
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L41
	} else {
		goto L441
	}
L427:
	;
	F_systable_endscan(m, v1337)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L41
	} else {
		goto L428
	}
L428:
	;
	F_ScanKeyInit(m, v1327, int32(1), int32(3), int32(184), v1313)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L41
	} else {
		goto L429
	}
L429:
	;
	v1353 = int32(1)
	v1356 = F_systable_beginscan(m, v1324, int32(2683), v1353, int32(0), v1353, v1327)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L41
	} else {
		goto L430
	}
L430:
	;
	goto L431
L431:
	;
	v1370 = F_systable_getnext(m, v1356)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L41
	} else {
		goto L433
	}
L432:
	;
	F_systable_endscan(m, v1356)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L41
	} else {
		goto L438
	}
L433:
	;
	if v1370 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	F_simple_heap_delete(m, v1324, v1370+int32(4))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L41
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	goto L432
L437:
	;
	goto L431
L438:
	;
	F_relation_close(m, v1324, int32(3))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L41
	} else {
		goto L439
	}
L439:
	;
	F_relation_close(m, v1320, int32(3))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L41
	} else {
		goto L440
	}
L440:
	;
	m.G0 = v1316 - int32(-64)
	goto L421
L441:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L41
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1313
	F_errmsg(m, int32(_a_F_doDeletion_40), v1316)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L41
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_doDeletion_41), int32(125), int32(_a_F_doDeletion_42))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L41
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	v1413 = F_SearchSysCache1(m, int32(40), v1403)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L41
	} else {
		goto L448
	}
L446:
	;
	goto L3
L447:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L41
	} else {
		goto L468
	}
L448:
	;
	if v1413 != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+16))
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415)+22)))
	v1417 = v1415 + v1416
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+92))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+96))
	if v1418|v1419 == int32(0) {
		v1437 = v1413
		goto L452
	} else {
		goto L453
	}
L450:
	;
	goto L451
L451:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L41
	} else {
		goto L465
	}
L452:
	;
	F_simple_heap_delete(m, v1410, v1437+int32(4))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L41
	} else {
		goto L462
	}
L453:
	;
	F_OperatorUpd(m, v1403, v1418, v1419, int32(1))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L41
	} else {
		goto L454
	}
L454:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+92))
	if v1426 != v1403 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+96))
	if v1403 != v1428 {
		v1437 = v1413
		goto L452
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	F_ReleaseCatCache(m, v1413)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L41
	} else {
		goto L459
	}
L458:
	;
	goto L457
L459:
	;
	v1433 = F_SearchSysCache1(m, int32(40), v1403)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L41
	} else {
		goto L460
	}
L460:
	;
	if v1433 == int32(0) {
		goto L447
	} else {
		goto L461
	}
L461:
	;
	v1437 = v1433
	goto L452
L462:
	;
	F_ReleaseCatCache(m, v1437)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L41
	} else {
		goto L463
	}
L463:
	;
	F_relation_close(m, v1410, int32(3))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L41
	} else {
		goto L464
	}
L464:
	;
	m.G0 = v1406 + int32(32)
	goto L446
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1406))) = v1403
	F_errmsg_internal(m, int32(_a_F_doDeletion_43), v1406)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L41
	} else {
		goto L466
	}
L466:
	;
	F_errfinish(m, int32(_a_F_doDeletion_44), int32(456), int32(_a_F_doDeletion_45))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L41
	} else {
		goto L467
	}
L467:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1406)+16)) = v1403
	F_errmsg_internal(m, int32(_a_F_doDeletion_43), v1406+int32(16))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L41
	} else {
		goto L469
	}
L469:
	;
	F_errfinish(m, int32(_a_F_doDeletion_44), int32(473), int32(_a_F_doDeletion_45))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L41
	} else {
		goto L470
	}
L470:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L471:
	;
	v1488 = v1481 + int32(32)
	F_ScanKeyInit(m, v1488, int32(1), int32(3), int32(184), v1478)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L41
	} else {
		goto L472
	}
L472:
	;
	v1495 = int32(1)
	v1498 = F_systable_beginscan(m, v1485, int32(2692), v1495, int32(0), v1495, v1488)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L41
	} else {
		goto L475
	}
L473:
	;
	goto L3
L474:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L41
	} else {
		goto L497
	}
L475:
	;
	v1500 = F_systable_getnext(m, v1498)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L41
	} else {
		goto L476
	}
L476:
	;
	if v1500 != 0 {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+16))
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1502)+22)))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1502+v1503)+68))
	v1507 = F_table_open(m, v1505, int32(8))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L41
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L41
	} else {
		goto L494
	}
L480:
	;
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_doDeletion[1])))
	if v1510 == int32(0) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v1514 = int32(1)
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+56))
	if base.Ui32(v1515) < base.Ui32(int32(_a_F_doDeletion_6)) {
		v1524 = v1514
		goto L485
	} else {
		goto L486
	}
L482:
	;
	goto L483
L483:
	;
	F_simple_heap_delete(m, v1485, v1500+int32(4))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L41
	} else {
		goto L489
	}
L484:
	;
	if v1524 != 0 {
		goto L474
	} else {
		goto L488
	}
L485:
	;
	goto L484
L486:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+48))
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+68))
	if v1519 == int32(99) {
		v1524 = v1514
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v1522 = F_isTempToastNamespace(m, v1519)
	mBase = m.M
	v1524 = v1522
	goto L485
L488:
	;
	goto L483
L489:
	;
	F_systable_endscan(m, v1498)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L41
	} else {
		goto L490
	}
L490:
	;
	F_relation_close(m, v1485, int32(3))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L41
	} else {
		goto L491
	}
L491:
	;
	F_CacheInvalidateRelcache(m, v1507)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L41
	} else {
		goto L492
	}
L492:
	;
	F_relation_close(m, v1507, int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L41
	} else {
		goto L493
	}
L493:
	;
	m.G0 = v1481 + int32(80)
	goto L473
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1481))) = v1478
	F_errmsg_internal(m, int32(_a_F_doDeletion_46), v1481)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L41
	} else {
		goto L495
	}
L495:
	;
	F_errfinish(m, int32(_a_F_doDeletion_47), int32(61), int32(_a_F_doDeletion_48))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L41
	} else {
		goto L496
	}
L496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L497:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L41
	} else {
		goto L498
	}
L498:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1481)+16)) = v1562 + int32(4)
	F_errmsg(m, int32(_a_F_doDeletion_8), v1481+int32(16))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L41
	} else {
		goto L499
	}
L499:
	;
	F_errfinish(m, int32(_a_F_doDeletion_47), int32(75), int32(_a_F_doDeletion_48))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L41
	} else {
		goto L500
	}
L500:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L501:
	;
	v1586 = v1579 + int32(48)
	F_ScanKeyInit(m, v1586, int32(1), int32(3), int32(184), v1576)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L41
	} else {
		goto L502
	}
L502:
	;
	v1593 = int32(1)
	v1596 = F_systable_beginscan(m, v1583, int32(2702), v1593, int32(0), v1593, v1586)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L41
	} else {
		goto L506
	}
L503:
	;
	goto L3
L504:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L41
	} else {
		goto L534
	}
L505:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L41
	} else {
		goto L529
	}
L506:
	;
	v1598 = F_systable_getnext(m, v1596)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L41
	} else {
		goto L507
	}
L507:
	;
	if v1598 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+16))
	v1601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1600)+22)))
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1600+v1601)+4))
	v1605 = F_table_open(m, v1603, int32(8))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L41
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L41
	} else {
		goto L526
	}
L511:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+48))
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607)+119)))
	v1610 = v1608 - int32(102)
	v1615 = int32(1)
	v1619 = (v1610<<(uint(int32(7))%32) | int32(base.Ui32(v1610&int32(254))>>(uint(v1615)%32))) & int32(255)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v1619))|base.B2i32(v1615<<(uint(v1619)%32)&int32(353) == int32(0)) != 0 {
		goto L505
	} else {
		goto L512
	}
L512:
	;
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_doDeletion[1])))
	if v1630 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v1634 = int32(1)
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+56))
	if base.Ui32(v1635) < base.Ui32(int32(_a_F_doDeletion_6)) {
		v1644 = v1634
		goto L517
	} else {
		goto L518
	}
L514:
	;
	goto L515
L515:
	;
	F_simple_heap_delete(m, v1583, v1598+int32(4))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L41
	} else {
		goto L521
	}
L516:
	;
	if v1644 != 0 {
		goto L504
	} else {
		goto L520
	}
L517:
	;
	goto L516
L518:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+48))
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+68))
	if v1639 == int32(99) {
		v1644 = v1634
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v1642 = F_isTempToastNamespace(m, v1639)
	mBase = m.M
	v1644 = v1642
	goto L517
L520:
	;
	goto L515
L521:
	;
	F_systable_endscan(m, v1596)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L41
	} else {
		goto L522
	}
L522:
	;
	F_relation_close(m, v1583, int32(3))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L41
	} else {
		goto L523
	}
L523:
	;
	F_CacheInvalidateRelcache(m, v1605)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L41
	} else {
		goto L524
	}
L524:
	;
	F_relation_close(m, v1605, int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L41
	} else {
		goto L525
	}
L525:
	;
	m.G0 = v1579 + int32(96)
	goto L503
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1579))) = v1576
	F_errmsg_internal(m, int32(_a_F_doDeletion_49), v1579)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L41
	} else {
		goto L527
	}
L527:
	;
	F_errfinish(m, int32(_a_F_doDeletion_50), int32(1316), int32(_a_F_doDeletion_51))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L41
	} else {
		goto L528
	}
L528:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L529:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L41
	} else {
		goto L530
	}
L530:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+16)) = v1682 + int32(4)
	F_errmsg(m, int32(_a_F_doDeletion_52), v1579+int32(16))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L41
	} else {
		goto L531
	}
L531:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+48))
	v1692 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1691)+119)))
	F_errdetail_relkind_not_supported(m, v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L41
	} else {
		goto L532
	}
L532:
	;
	F_errfinish(m, int32(_a_F_doDeletion_50), int32(1333), int32(_a_F_doDeletion_51))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L41
	} else {
		goto L533
	}
L533:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L534:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L41
	} else {
		goto L535
	}
L535:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+32)) = v1707 + int32(4)
	F_errmsg(m, int32(_a_F_doDeletion_8), v1579+int32(32))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L41
	} else {
		goto L536
	}
L536:
	;
	F_errfinish(m, int32(_a_F_doDeletion_50), int32(1339), int32(_a_F_doDeletion_51))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L41
	} else {
		goto L537
	}
L537:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L538:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1724 = m.G0
	v1726 = v1724 - int32(16)
	m.G0 = v1726
	v1730 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L41
	} else {
		goto L539
	}
L539:
	;
	v1733 = F_SearchSysCache1(m, int32(64), v1723)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L41
	} else {
		goto L541
	}
L540:
	;
	goto L3
L541:
	;
	if v1733 != 0 {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+16))
	v1736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1735)+22)))
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1735+v1736)+4))
	v1740 = F_table_open(m, v1738, int32(4))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L41
	} else {
		goto L545
	}
L543:
	;
	goto L544
L544:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L41
	} else {
		goto L567
	}
L545:
	;
	v1744 = F_table_open(m, int32(3429), int32(3))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L41
	} else {
		goto L546
	}
L546:
	;
	v1748 = F_SearchSysCache2(m, int32(62), v1723, int32(1))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L41
	} else {
		goto L547
	}
L547:
	;
	if v1748 != 0 {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	F_simple_heap_delete(m, v1744, v1748+int32(4))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L41
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	F_relation_close(m, v1744, int32(3))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L41
	} else {
		goto L553
	}
L551:
	;
	F_ReleaseCatCache(m, v1748)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L41
	} else {
		goto L552
	}
L552:
	;
	goto L550
L553:
	;
	v1761 = F_table_open(m, int32(3429), int32(3))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L41
	} else {
		goto L554
	}
L554:
	;
	v1765 = F_SearchSysCache2(m, int32(62), v1723, int32(0))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L41
	} else {
		goto L555
	}
L555:
	;
	if v1765 != 0 {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	F_simple_heap_delete(m, v1761, v1765+int32(4))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L41
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	F_relation_close(m, v1761, int32(3))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L41
	} else {
		goto L561
	}
L559:
	;
	F_ReleaseCatCache(m, v1765)
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L41
	} else {
		goto L560
	}
L560:
	;
	goto L558
L561:
	;
	F_CacheInvalidateRelcacheByRelid(m, v1738)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L41
	} else {
		goto L562
	}
L562:
	;
	F_simple_heap_delete(m, v1730, v1733+int32(4))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L41
	} else {
		goto L563
	}
L563:
	;
	F_ReleaseCatCache(m, v1733)
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L41
	} else {
		goto L564
	}
L564:
	;
	F_relation_close(m, v1740, int32(0))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L41
	} else {
		goto L565
	}
L565:
	;
	F_relation_close(m, v1730, int32(3))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L41
	} else {
		goto L566
	}
L566:
	;
	m.G0 = v1726 + int32(16)
	goto L540
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1726))) = v1723
	F_errmsg_internal(m, int32(_a_F_doDeletion_53), v1726)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L41
	} else {
		goto L568
	}
L568:
	;
	F_errfinish(m, int32(_a_F_doDeletion_54), int32(803), int32(_a_F_doDeletion_55))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L41
	} else {
		goto L569
	}
L569:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L570:
	;
	v1816 = F_SearchSysCache1(m, int32(74), v1806)
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L41
	} else {
		goto L572
	}
L571:
	;
	goto L3
L572:
	;
	if v1816 != 0 {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	F_simple_heap_delete(m, v1813, v1816+int32(4))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L41
	} else {
		goto L576
	}
L574:
	;
	goto L575
L575:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L41
	} else {
		goto L591
	}
L576:
	;
	F_ReleaseCatCache(m, v1816)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L41
	} else {
		goto L577
	}
L577:
	;
	F_relation_close(m, v1813, int32(3))
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L41
	} else {
		goto L578
	}
L578:
	;
	v1829 = F_table_open(m, int32(3603), int32(3))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L41
	} else {
		goto L579
	}
L579:
	;
	v1832 = v1807 + int32(-48)
	F_ScanKeyInit(m, v1832, int32(1), int32(3), int32(184), v1806)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L41
	} else {
		goto L580
	}
L580:
	;
	v1839 = int32(1)
	v1842 = F_systable_beginscan(m, v1829, int32(3609), v1839, int32(0), v1839, v1832)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L41
	} else {
		goto L581
	}
L581:
	;
	goto L582
L582:
	;
	v1856 = F_systable_getnext(m, v1842)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L41
	} else {
		goto L584
	}
L583:
	;
	F_systable_endscan(m, v1842)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L41
	} else {
		goto L589
	}
L584:
	;
	if v1856 != 0 {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	F_simple_heap_delete(m, v1829, v1856+int32(4))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L41
	} else {
		goto L588
	}
L586:
	;
	goto L587
L587:
	;
	goto L583
L588:
	;
	goto L582
L589:
	;
	F_relation_close(m, v1829, int32(3))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L41
	} else {
		goto L590
	}
L590:
	;
	m.G0 = v1809 - int32(-64)
	goto L571
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1809))) = v1806
	F_errmsg_internal(m, int32(_a_F_doDeletion_56), v1809)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L41
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_doDeletion_57), int32(1123), int32(_a_F_doDeletion_58))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L41
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	goto L3
L595:
	;
	v1893 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L41
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L41
	} else {
		goto L608
	}
L598:
	;
	v1896 = v1884 + int32(-48)
	F_ScanKeyInit(m, v1896, int32(1), int32(3), int32(184), v1883)
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L41
	} else {
		goto L599
	}
L599:
	;
	v1903 = int32(1)
	v1906 = F_systable_beginscan(m, v1893, int32(3080), v1903, int32(0), v1903, v1896)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L41
	} else {
		goto L600
	}
L600:
	;
	v1908 = F_systable_getnext(m, v1906)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L41
	} else {
		goto L601
	}
L601:
	;
	if v1908 != 0 {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	F_simple_heap_delete(m, v1893, v1908+int32(4))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L41
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	F_systable_endscan(m, v1906)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L41
	} else {
		goto L606
	}
L605:
	;
	goto L604
L606:
	;
	F_relation_close(m, v1893, int32(3))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L41
	} else {
		goto L607
	}
L607:
	;
	m.G0 = v1886 - int32(-64)
	goto L594
L608:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L41
	} else {
		goto L609
	}
L609:
	;
	v1929 = F_get_extension_name(m, v1883)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L41
	} else {
		goto L610
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1886))) = v1929
	F_errmsg(m, int32(_a_F_doDeletion_59), v1886)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L41
	} else {
		goto L611
	}
L611:
	;
	F_errfinish(m, int32(_a_F_doDeletion_60), int32(2302), int32(_a_F_doDeletion_61))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L41
	} else {
		goto L612
	}
L612:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L613:
	;
	v1950 = F_SearchSysCache1(m, int32(49), v1940)
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L41
	} else {
		goto L615
	}
L614:
	;
	goto L3
L615:
	;
	if v1950 != 0 {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+16))
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1952)+22)))
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1952+v1953)+8))
	v1957 = F_GetSchemaPublicationRelations(m, v1955, int32(2))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L41
	} else {
		goto L620
	}
L617:
	;
	goto L618
L618:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L41
	} else {
		goto L634
	}
L619:
	;
	F_simple_heap_delete(m, v1947, v1950+int32(4))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L41
	} else {
		goto L631
	}
L620:
	;
	if v1957 == int32(0) {
		goto L619
	} else {
		goto L621
	}
L621:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1957)+4))
	if v1961 <= int32(4095) {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	if v1961 <= int32(0) {
		goto L619
	} else {
		goto L625
	}
L623:
	;
	goto L624
L624:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L41
	} else {
		goto L630
	}
L625:
	;
	v1967 = int32(0)
	goto L626
L626:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1957)+12))
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1979+v1967<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v1983)
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L41
	} else {
		goto L628
	}
L627:
	;
	goto L619
L628:
	;
	v1987 = v1967 + int32(1)
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1957)+4))
	if v1987 < v1988 {
		v1967 = v1987
		goto L626
	} else {
		goto L629
	}
L629:
	;
	goto L627
L630:
	;
	goto L619
L631:
	;
	F_ReleaseCatCache(m, v1950)
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L41
	} else {
		goto L632
	}
L632:
	;
	F_relation_close(m, v1947, int32(3))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L41
	} else {
		goto L633
	}
L633:
	;
	m.G0 = v1943 + int32(16)
	goto L614
L634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1943))) = v1940
	F_errmsg_internal(m, int32(_a_F_doDeletion_62), v1943)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L41
	} else {
		goto L635
	}
L635:
	;
	F_errfinish(m, int32(_a_F_doDeletion_63), int32(1635), int32(_a_F_doDeletion_64))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L41
	} else {
		goto L636
	}
L636:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L637:
	;
	v2039 = F_SearchSysCache1(m, int32(52), v2029)
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L41
	} else {
		goto L639
	}
L638:
	;
	goto L3
L639:
	;
	if v2039 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2039)+16))
	v2044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2043)+22)))
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2043+v2044)+8))
	v2047 = F_GetPubPartitionOptionRelations(m, int32(0), int32(2), v2046)
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L41
	} else {
		goto L644
	}
L641:
	;
	goto L642
L642:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L41
	} else {
		goto L658
	}
L643:
	;
	F_simple_heap_delete(m, v2036, v2039+int32(4))
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L41
	} else {
		goto L655
	}
L644:
	;
	if v2047 == int32(0) {
		goto L643
	} else {
		goto L645
	}
L645:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+4))
	if v2051 <= int32(4095) {
		goto L646
	} else {
		goto L647
	}
L646:
	;
	if v2051 <= int32(0) {
		goto L643
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L41
	} else {
		goto L654
	}
L649:
	;
	v2057 = int32(0)
	goto L650
L650:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+12))
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2069+v2057<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v2073)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L41
	} else {
		goto L652
	}
L651:
	;
	goto L643
L652:
	;
	v2077 = v2057 + int32(1)
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+4))
	if v2077 < v2078 {
		v2057 = v2077
		goto L650
	} else {
		goto L653
	}
L653:
	;
	goto L651
L654:
	;
	goto L643
L655:
	;
	F_ReleaseCatCache(m, v2039)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L41
	} else {
		goto L656
	}
L656:
	;
	F_relation_close(m, v2036, int32(3))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L41
	} else {
		goto L657
	}
L657:
	;
	m.G0 = v2032 + int32(16)
	goto L638
L658:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2032))) = v2029
	F_errmsg_internal(m, int32(_a_F_doDeletion_65), v2032)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L41
	} else {
		goto L659
	}
L659:
	;
	F_errfinish(m, int32(_a_F_doDeletion_63), int32(1566), int32(_a_F_doDeletion_66))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L41
	} else {
		goto L660
	}
L660:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L661:
	;
	v2129 = F_SearchSysCache1(m, int32(51), v2119)
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L41
	} else {
		goto L663
	}
L662:
	;
	goto L3
L663:
	;
	if v2129 != 0 {
		goto L664
	} else {
		goto L665
	}
L664:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2129)+16))
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2131)+22)))
	v2134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2131+v2132)+72)))
	if v2134 == int32(1) {
		goto L667
	} else {
		goto L668
	}
L665:
	;
	goto L666
L666:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L41
	} else {
		goto L674
	}
L667:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L41
	} else {
		goto L670
	}
L668:
	;
	goto L669
L669:
	;
	F_simple_heap_delete(m, v2126, v2129+int32(4))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L41
	} else {
		goto L671
	}
L670:
	;
	goto L669
L671:
	;
	F_ReleaseCatCache(m, v2129)
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L41
	} else {
		goto L672
	}
L672:
	;
	F_relation_close(m, v2126, int32(3))
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L41
	} else {
		goto L673
	}
L673:
	;
	m.G0 = v2122 + int32(16)
	goto L662
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2122))) = v2119
	F_errmsg_internal(m, int32(_a_F_doDeletion_67), v2122)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L41
	} else {
		goto L675
	}
L675:
	;
	F_errfinish(m, int32(_a_F_doDeletion_63), int32(1604), int32(_a_F_doDeletion_68))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L41
	} else {
		goto L676
	}
L676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L677:
	;
	F_errmsg_internal(m, int32(_a_F_doDeletion_69), int32(0))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L41
	} else {
		goto L678
	}
L678:
	;
	F_errfinish(m, int32(_a_F_doDeletion_70), int32(1478), int32(_a_F_doDeletion_71))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L41
	} else {
		goto L679
	}
L679:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L680:
	;
	goto L6
L681:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v2183
	F_errmsg_internal(m, int32(_a_F_doDeletion_72), v15)
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L41
	} else {
		goto L682
	}
L682:
	;
	F_errfinish(m, int32(_a_F_doDeletion_70), int32(1482), int32(_a_F_doDeletion_71))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L41
	} else {
		goto L683
	}
L683:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L684:
	;
	v2203 = F_SearchSysCache1(m, int32(47), v2193)
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L41
	} else {
		goto L687
	}
L685:
	;
	goto L3
L686:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L41
	} else {
		goto L707
	}
L687:
	;
	if v2203 != 0 {
		goto L688
	} else {
		goto L689
	}
L688:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v2203)+16))
	v2206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2205)+22)))
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2205+v2206)+96)))
	F_simple_heap_delete(m, v2200, v2203+int32(4))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L41
	} else {
		goto L691
	}
L689:
	;
	goto L690
L690:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L41
	} else {
		goto L704
	}
L691:
	;
	F_ReleaseCatCache(m, v2203)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L41
	} else {
		goto L692
	}
L692:
	;
	F_relation_close(m, v2200, int32(3))
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		goto L41
	} else {
		goto L693
	}
L693:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, _c_F_doDeletion[5]))
	F_pgstat_drop_transactional(m, int32(3), v2220, base.I64_extend_i32_u(v2193))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L41
	} else {
		goto L694
	}
L694:
	;
	if v2208 == int32(97) {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	v2228 = F_table_open(m, int32(2600), int32(3))
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L41
	} else {
		goto L698
	}
L696:
	;
	goto L697
L697:
	;
	m.G0 = v2196 + int32(32)
	goto L685
L698:
	;
	v2231 = F_SearchSysCache1(m, int32(0), v2193)
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L41
	} else {
		goto L699
	}
L699:
	;
	if v2231 == int32(0) {
		goto L686
	} else {
		goto L700
	}
L700:
	;
	F_simple_heap_delete(m, v2228, v2231+int32(4))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L41
	} else {
		goto L701
	}
L701:
	;
	F_ReleaseCatCache(m, v2231)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L41
	} else {
		goto L702
	}
L702:
	;
	F_relation_close(m, v2228, int32(3))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L41
	} else {
		goto L703
	}
L703:
	;
	goto L697
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2196))) = v2193
	F_errmsg_internal(m, int32(_a_F_doDeletion_73), v2196)
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L41
	} else {
		goto L705
	}
L705:
	;
	F_errfinish(m, int32(_a_F_doDeletion_74), int32(1324), int32(_a_F_doDeletion_75))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L41
	} else {
		goto L706
	}
L706:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2196)+16)) = v2193
	F_errmsg_internal(m, int32(_a_F_doDeletion_76), v2196+int32(16))
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L41
	} else {
		goto L708
	}
L708:
	;
	F_errfinish(m, int32(_a_F_doDeletion_74), int32(1345), int32(_a_F_doDeletion_75))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L41
	} else {
		goto L709
	}
L709:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L710:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2281 = F_table_open(m, v2279, int32(3))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L41
	} else {
		goto L711
	}
L711:
	;
	if int32(0) <= v2277 {
		goto L713
	} else {
		goto L714
	}
L712:
	;
	F_relation_close(m, v2281, int32(3))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L41
	} else {
		goto L728
	}
L713:
	;
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2286 = F_SearchSysCache1(m, v2277, v2285)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L41
	} else {
		goto L716
	}
L714:
	;
	goto L715
L715:
	;
	v2297 = v15 + int32(48)
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2299 = F_get_object_attnum_oid(m, v2298)
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L41
	} else {
		goto L720
	}
L716:
	;
	if v2286 == int32(0) {
		goto L2
	} else {
		goto L717
	}
L717:
	;
	F_simple_heap_delete(m, v2281, v2286+int32(4))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L41
	} else {
		goto L718
	}
L718:
	;
	F_ReleaseCatCache(m, v2286)
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L41
	} else {
		goto L719
	}
L719:
	;
	goto L712
L720:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v2297, v2299, int32(3), int32(184), v2303)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L41
	} else {
		goto L721
	}
L721:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2307 = F_get_object_oid_index(m, v2306)
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L41
	} else {
		goto L722
	}
L722:
	;
	v2309 = int32(1)
	v2312 = F_systable_beginscan(m, v2281, v2307, v2309, int32(0), v2309, v2297)
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L41
	} else {
		goto L723
	}
L723:
	;
	v2314 = F_systable_getnext(m, v2312)
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L41
	} else {
		goto L724
	}
L724:
	;
	if v2314 == int32(0) {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	F_simple_heap_delete(m, v2281, v2314+int32(4))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L41
	} else {
		goto L726
	}
L726:
	;
	F_systable_endscan(m, v2312)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L41
	} else {
		goto L727
	}
L727:
	;
	goto L712
L728:
	;
	goto L3
L729:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2349 = F_get_object_class_descr(m, v2348)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L41
	} else {
		goto L730
	}
L730:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v2349
	F_errmsg_internal(m, int32(_a_F_doDeletion_77), v15+int32(16))
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L41
	} else {
		goto L731
	}
L731:
	;
	F_errfinish(m, int32(_a_F_doDeletion_70), int32(1207), int32(_a_F_doDeletion_78))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L41
	} else {
		goto L732
	}
L732:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L733:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2369 = F_get_object_class_descr(m, v2368)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L41
	} else {
		goto L734
	}
L734:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v2371
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v2369
	F_errmsg_internal(m, int32(_a_F_doDeletion_79), v15+int32(32))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L41
	} else {
		goto L735
	}
L735:
	;
	F_errfinish(m, int32(_a_F_doDeletion_70), int32(1230), int32(_a_F_doDeletion_78))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L41
	} else {
		goto L736
	}
L736:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dopr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v39 int32
	_ = v39
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
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
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
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v426 int32
	_ = v426
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 float64
	_ = v837
	var v838 int64
	_ = v838
	var v846 int32
	_ = v846
	var v850 float64
	_ = v850
	var v860 int32
	_ = v860
	var v861 float64
	_ = v861
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v871 int64
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1040 int32
	_ = v1040
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1189 int32
	_ = v1189
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1207 int32
	_ = v1207
	var v1208 int64
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int64
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1220 int64
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int64
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1238 int32
	_ = v1238
	var v1239 int64
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1242 int64
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1251 int64
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int64
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1316 int32
	_ = v1316
	var v1325 int32
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1338 int32
	_ = v1338
	var v1339 float64
	_ = v1339
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1534 int32
	_ = v1534
	v4 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(1360)
	m.G0 = v33
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_dopr[0]))
	v37 = l0
	v38 = l1
	v39 = l2
	v42 = v33
	v49 = v4
	v60 = v4
	v64 = v36
	goto L1
L1:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v67 == int32(37) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	m.G0 = v1534 + int32(1360)
	return
L3:
	;
	goto L2
L4:
	;
	if v60 != 0 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v177 = v38
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v67 == int32(0) {
		v1534 = v42
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v73 = v38 + int32(1)
	goto L13
L9:
	;
	F_dostr(m, v38, v159-v38, v37)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	goto L9
L11:
	;
	v149 = v144
	goto L28
L12:
	;
	v144 = v136
	goto L11
L13:
	;
	if v73&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v82 = v73
	goto L19
L17:
	;
	v96 = v73
	goto L18
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v105 = int32(-2139062144)
	if (int32(16843008)-v102|v102)&v105 != v105 {
		v136 = v96
		goto L12
	} else {
		goto L23
	}
L19:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.B2i32(v87 == int32(0))|base.B2i32(int32(37) == v87) != 0 {
		v159 = v82
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v96 = v93
	goto L18
L21:
	;
	v93 = v82 + int32(1)
	if v93&int32(3) != 0 {
		v82 = v93
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v111 = v96
	v113 = v102
	goto L24
L24:
	;
	v117 = v113 ^ int32(623191333)
	v120 = int32(-2139062144)
	if (int32(16843008)-v117|v117)&v120 != v120 {
		v136 = v111
		goto L12
	} else {
		goto L26
	}
L25:
	;
	v144 = v126
	goto L11
L26:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v126 = v111 + int32(4)
	v130 = int32(-2139062144)
	if (v124|(int32(16843008)-v124))&v130 == v130 {
		v111 = v126
		v113 = v124
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v151 == int32(0) {
		v159 = v149
		goto L10
	} else {
		goto L30
	}
L29:
	;
	v159 = v149
	goto L10
L30:
	;
	if v151 != int32(37) {
		v149 = v149 + int32(1)
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	return
L33:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+20)))
	if v173 != 0 {
		v1534 = v42
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v174 == int32(0) {
		v1534 = v42
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v177 = v159
	goto L4
L36:
	;
	v178 = v60
	goto L38
L37:
	;
	v178 = v177
	goto L38
L38:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v179 != int32(115) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v184 = int32(0)
	v197 = v37
	v198 = v177 + int32(1)
	v199 = v39
	v201 = v179
	v202 = v42
	v203 = v184
	v204 = v184
	v206 = v184
	v208 = v184
	v209 = v49
	v211 = v184
	v212 = v184
	v213 = v184
	v214 = v184
	v215 = v184
	v217 = v184
	v218 = v184
	v220 = v178
	v221 = v184
	v222 = v184
	v224 = v64
	goto L42
L40:
	;
	goto L41
L41:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v1516 != 0 {
		goto L427
	} else {
		goto L428
	}
L42:
	;
	v227 = int32(1)
	v229 = v198 + v227
	v230 = base.I32_extend8_s(v201)
	switch v201&int32(255) - int32(36) {
	case 0:
		goto L62
	case 1:
		goto L61
	default:
		goto L48
	case 3, 68:
		v1452 = v214
		v1453 = v222
		goto L46
	case 6:
		goto L63
	case 7:
		goto L67
	case 9:
		v1487 = v199
		v1488 = v227
		v1492 = v204
		v1494 = v206
		v1497 = v209
		v1499 = v211
		v1500 = v212
		v1501 = v213
		v1502 = v214
		v1503 = v215
		v1505 = v217
		v1506 = v218
		v1509 = v221
		v1510 = v222
		goto L44
	case 10:
		goto L64
	case 12:
		goto L66
	case 13, 14, 15, 16, 17, 18, 19, 20, 21:
		v239 = v213
		goto L65
	case 33, 35, 65, 66, 67:
		goto L59
	case 52, 75, 81, 84:
		goto L55
	case 63:
		goto L56
	case 64, 69:
		goto L54
	case 72:
		goto L51
	case 73:
		goto L60
	case 76:
		goto L58
	case 79:
		goto L57
	case 86:
		goto L52
	}
L44:
	;
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v198 = v229
	v199 = v1487
	v201 = v1515
	v203 = v1488
	v204 = v1492
	v206 = v1494
	v208 = v1492
	v209 = v1497
	v211 = v1499
	v212 = v1500
	v213 = v1501
	v214 = v1502
	v215 = v1503
	v217 = v1505
	v218 = v1506
	v221 = v1509
	v222 = v1510
	goto L42
L45:
	;
	v1487 = v199
	v1488 = v203
	v1492 = v1484
	v1494 = v1463
	v1497 = v1466
	v1499 = v1468
	v1500 = v1469
	v1501 = v1470
	v1502 = v1471
	v1503 = v1472
	v1505 = v217
	v1506 = v1475
	v1509 = v1478
	v1510 = v1479
	goto L44
L46:
	;
	v1463 = v206
	v1466 = v209
	v1468 = v211
	v1469 = v212
	v1470 = v213
	v1471 = v1452
	v1472 = v215
	v1475 = v218
	v1478 = v221
	v1479 = v1453
	v1484 = v208
	goto L45
L47:
	;
	v1487 = v199 + int32(4)
	v1488 = v1447
	v1492 = v250
	v1494 = v1448
	v1497 = int32(0)
	v1499 = v1449
	v1500 = v249
	v1501 = v213
	v1502 = v214
	v1503 = v1450
	v1505 = v217
	v1506 = v218
	v1509 = v221
	v1510 = v222
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dopr[0])) = int32(28)
	v1445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v1445)
	v1534 = v202
	goto L3
L49:
	;
	v1386 = int32(1)
	v1387 = int32(0)
	if v221 == v1387 {
		goto L415
	} else {
		goto L416
	}
L50:
	;
	if v292 <= int32(0) {
		goto L49
	} else {
		goto L407
	}
L51:
	;
	if v222 != 0 {
		goto L404
	} else {
		goto L405
	}
L52:
	;
	v1452 = v214
	v1453 = int32(1)
	goto L46
L53:
	;
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)))
	if v1270 == int32(0) {
		v37 = v197
		v38 = v229
		v39 = v1259
		v42 = v202
		v49 = v209
		v60 = v220
		v64 = v224
		goto L1
	} else {
		goto L403
	}
L54:
	;
	if v206 != 0 {
		goto L378
	} else {
		goto L379
	}
L55:
	;
	if v206 != 0 {
		goto L353
	} else {
		goto L354
	}
L56:
	;
	v1124 = v209 & int32(1)
	if v1124 != 0 {
		goto L317
	} else {
		goto L318
	}
L57:
	;
	if v206 != 0 {
		goto L277
	} else {
		goto L278
	}
L58:
	;
	v1047 = v209 & int32(1)
	if v1047 != 0 {
		goto L266
	} else {
		goto L267
	}
L59:
	;
	v826 = (v199 + int32(7)) & int32(-8)
	v828 = v826 + int32(8)
	v835 = v209 & int32(1)
	if v835 != 0 {
		goto L184
	} else {
		goto L185
	}
L60:
	;
	v818 = F_pg_strerror_r(m, v224, v202+int32(320))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L32
	} else {
		goto L179
	}
L61:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v778 = int32(0)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if base.B2i32(v777 == v778)|base.B2i32(base.Ui32(v780) < base.Ui32(v777)) == v778 {
		goto L168
	} else {
		goto L169
	}
L62:
	;
	if v209&int32(1) != 0 {
		goto L49
	} else {
		goto L86
	}
L63:
	;
	v249 = int32(1)
	v250 = int32(0)
	if v209&v249 != 0 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	if v212 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v1463 = v206
	v1466 = v209
	v1468 = v211
	v1469 = v212
	v1470 = v239
	v1471 = v214
	v1472 = v215
	v1475 = v218
	v1478 = v221
	v1479 = v222
	v1484 = v208*int32(10) + v230 - int32(48)
	goto L45
L66:
	;
	if v206|v208 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v1452 = int32(1)
	v1453 = v222
	goto L46
L68:
	;
	v238 = v213
	goto L70
L69:
	;
	v238 = int32(48)
	goto L70
L70:
	;
	v239 = v238
	goto L65
L71:
	;
	v246 = v211
	goto L73
L72:
	;
	v246 = v208
	goto L73
L73:
	;
	v247 = int32(0)
	v1463 = int32(1)
	v1466 = v209
	v1468 = v246
	v1469 = v247
	v1470 = v213
	v1471 = v214
	v1472 = v215
	v1475 = v218
	v1478 = v221
	v1479 = v222
	v1484 = v247
	goto L45
L74:
	;
	v253 = int32(1)
	v1487 = v199
	v1488 = v203
	v1492 = v250
	v1494 = v206
	v1497 = v253
	v1499 = v211
	v1500 = v249
	v1501 = v213
	v1502 = v214
	v1503 = v215
	v1505 = v217
	v1506 = v218
	v1509 = v253
	v1510 = v222
	goto L44
L75:
	;
	goto L76
L76:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v206 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v258 = int32(0)
	v260 = base.B2i32(v258 <= v257)
	if v258 <= v257 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v263 = v257 >> (uint(int32(31)) % 32)
	if v257 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v261 = v257
	goto L82
L81:
	;
	v261 = v258
	goto L82
L82:
	;
	v1447 = v203
	v1448 = v260
	v1449 = v211
	v1450 = v261
	goto L47
L83:
	;
	v269 = int32(1)
	goto L85
L84:
	;
	v269 = v203
	goto L85
L85:
	;
	v1447 = v269
	v1448 = int32(0)
	v1449 = v257 ^ v263 - v263
	v1450 = v215
	goto L47
L86:
	;
	v273 = int32(0)
	base.MemoryFill(m, v202+int32(320), v273, int32(128))
	v282 = v220
	v292 = v273
	goto L87
L87:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v309 != int32(37) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L48
L89:
	;
	if v309 == int32(0) {
		goto L50
	} else {
		goto L92
	}
L90:
	;
	v326 = v282
	goto L91
L91:
	;
	v329 = int32(0)
	v335 = v326 + int32(1)
	v336 = v329
	v342 = v329
	v345 = v292
	v354 = v329
	goto L98
L92:
	;
	v316 = int32(37)
	v317 = F___strchrnul(m, v282+int32(1), v316)
	mBase = m.M
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	if v319 == v316 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v323 == int32(0) {
		goto L50
	} else {
		goto L97
	}
L94:
	;
	v323 = v317
	goto L96
L95:
	;
	v323 = int32(0)
	goto L96
L96:
	;
	goto L93
L97:
	;
	v326 = v323
	goto L91
L98:
	;
	v366 = v335
	v367 = v336
	v370 = int32(0)
	v385 = v354
	goto L100
L99:
	;
	goto L88
L100:
	;
	v396 = v366
	v412 = v367
	v415 = v385
	goto L102
L101:
	;
	if base.Ui32(int32(-31)) <= base.Ui32(v370-int32(32)) {
		v335 = v456
		v336 = v412
		v342 = v370
		v354 = v415
		goto L98
	} else {
		goto L167
	}
L102:
	;
	v426 = v396
	goto L111
L103:
	;
	goto L101
L104:
	;
	goto L103
L105:
	;
	if v412 != 0 {
		goto L164
	} else {
		goto L165
	}
L106:
	;
	if v342 == int32(0) {
		goto L48
	} else {
		goto L149
	}
L107:
	;
	if v342 == int32(0) {
		goto L48
	} else {
		goto L143
	}
L108:
	;
	if v342 == int32(0) {
		goto L48
	} else {
		goto L137
	}
L109:
	;
	if v342 == int32(0) {
		goto L48
	} else {
		goto L131
	}
L110:
	;
	v469 = v456
	v470 = v412
	v473 = v457
	v488 = v415
	goto L114
L111:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	v454 = int32(1)
	v456 = v426 + v454
	v457 = int32(0)
	switch v453 - int32(36) {
	case 0:
		goto L104
	case 1, 73:
		v282 = v456
		v292 = v345
		goto L87
	default:
		goto L48
	case 3, 7, 9, 68:
		v426 = v456
		goto L111
	case 6:
		goto L110
	case 10:
		v366 = v456
		v367 = v412
		v370 = v457
		v385 = v415
		goto L100
	case 12, 13, 14, 15, 16, 17, 18, 19, 20, 21:
		goto L113
	case 33, 35, 65, 66, 67:
		v589 = v456
		v598 = v457
		goto L109
	case 52, 64, 69, 75, 81, 84:
		v723 = v456
		v724 = v412
		v732 = v457
		v742 = v415
		goto L106
	case 63:
		v677 = v456
		v686 = v457
		goto L107
	case 72:
		goto L105
	case 76, 79:
		v633 = v456
		v642 = v457
		goto L108
	case 86:
		v396 = v456
		v412 = v454
		goto L102
	}
L112:
	;
	v366 = v456
	v367 = v412
	v370 = v370*int32(10) + v453 - int32(48)
	v385 = v415
	goto L100
L113:
	;
	goto L112
L114:
	;
	v499 = v469
	v508 = v470
	v518 = v488
	goto L117
L116:
	;
	v469 = v559
	v470 = v508
	v473 = v473*int32(10) + v556 - int32(48)
	v488 = v518
	goto L114
L117:
	;
	v529 = v499
	goto L120
L118:
	;
	if base.Ui32(v473-int32(32)) < base.Ui32(int32(-31)) {
		goto L48
	} else {
		goto L126
	}
L119:
	;
	goto L118
L120:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	v557 = int32(1)
	v559 = v529 + v557
	switch v556 - int32(36) {
	case 0:
		goto L119
	default:
		goto L48
	case 3, 7, 9, 68:
		v529 = v559
		goto L120
	case 10:
		v469 = v559
		v470 = v508
		v473 = int32(0)
		v488 = v518
		goto L114
	case 12, 13, 14, 15, 16, 17, 18, 19, 20, 21:
		goto L116
	case 33, 35, 65, 66, 67:
		v589 = v559
		v598 = v557
		goto L109
	case 52, 64, 69, 75, 81, 84:
		v723 = v559
		v724 = v508
		v732 = v557
		v742 = v518
		goto L106
	case 63:
		v677 = v559
		v686 = v557
		goto L107
	case 72:
		goto L122
	case 76, 79:
		v633 = v559
		v642 = v557
		goto L108
	case 86:
		v499 = v559
		v508 = v557
		goto L117
	}
L121:
	;
	if v508 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	v564 = int32(1)
	goto L125
L124:
	;
	v564 = v518
	goto L125
L125:
	;
	v499 = v559
	v508 = v557
	v518 = v564
	goto L117
L126:
	;
	v573 = v202 + int32(320) + v473<<(uint(int32(2))%32)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	if base.Ui32(int32(1)) < base.Ui32(v574) {
		goto L48
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = int32(1)
	if v473 < v345 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v580 = v345
	goto L130
L129:
	;
	v580 = v473
	goto L130
L130:
	;
	v335 = v559
	v336 = v508
	v345 = v580
	v354 = v518
	goto L98
L131:
	;
	v622 = v202 + int32(320) + v342<<(uint(int32(2))%32)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	switch v623 {
	case 0, 4:
		goto L132
	default:
		goto L48
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = int32(4)
	if v342 < v345 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v627 = v345
	goto L135
L134:
	;
	v627 = v342
	goto L135
L135:
	;
	if v598 == int32(0) {
		v282 = v589
		v292 = v627
		goto L87
	} else {
		goto L136
	}
L136:
	;
	goto L48
L137:
	;
	v666 = v202 + int32(320) + v342<<(uint(int32(2))%32)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)))
	switch v667 {
	case 0, 5:
		goto L138
	default:
		goto L48
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v666))) = int32(5)
	if v342 < v345 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v671 = v345
	goto L141
L140:
	;
	v671 = v342
	goto L141
L141:
	;
	if v642 == int32(0) {
		v282 = v633
		v292 = v671
		goto L87
	} else {
		goto L142
	}
L142:
	;
	goto L48
L143:
	;
	v710 = v202 + int32(320) + v342<<(uint(int32(2))%32)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	if base.Ui32(int32(1)) < base.Ui32(v711) {
		goto L48
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710))) = int32(1)
	if v342 < v345 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v717 = v345
	goto L147
L146:
	;
	v717 = v342
	goto L147
L147:
	;
	if v686 == int32(0) {
		v282 = v677
		v292 = v717
		goto L87
	} else {
		goto L148
	}
L148:
	;
	goto L48
L149:
	;
	v754 = int32(2)
	v756 = v202 + int32(320) + v342<<(uint(v754)%32)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	if v724 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v762 = v754
	goto L152
L151:
	;
	v762 = int32(1)
	goto L152
L152:
	;
	if v742 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v763 = int32(3)
	goto L155
L154:
	;
	v763 = v762
	goto L155
L155:
	;
	if v757 != v763 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v765 = v757
	goto L158
L157:
	;
	v765 = int32(0)
	goto L158
L158:
	;
	if v765 != 0 {
		goto L48
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v756))) = v763
	if v342 < v345 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v768 = v345
	goto L162
L161:
	;
	v768 = v342
	goto L162
L162:
	;
	if v732 == int32(0) {
		v282 = v723
		v292 = v768
		goto L87
	} else {
		goto L163
	}
L163:
	;
	goto L48
L164:
	;
	v772 = int32(1)
	goto L166
L165:
	;
	v772 = v415
	goto L166
L166:
	;
	v396 = v456
	v412 = v454
	v415 = v772
	goto L102
L167:
	;
	goto L99
L168:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	if v785 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v808 = v780
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v808 + int32(1)
	v814 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v808))) = uint8(v814)
	v1259 = v199
	goto L53
L171:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v788 + int32(1)
	v1259 = v199
	goto L53
L172:
	;
	goto L173
L173:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)))
	if v792 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v808 = v807
	goto L170
L175:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v780 == v793 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v796 = v780 - v793
	v797 = F_fwrite(m, v793, int32(1), v796, v785)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L32
	} else {
		goto L177
	}
L177:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v797 + v799
	if v796 == v797 {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v803 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v803)
	goto L174
L179:
	;
	v820 = F_strlen(m, v818)
	mBase = m.M
	F_dostr(m, v818, v820, v197)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L32
	} else {
		goto L180
	}
L180:
	;
	v1259 = v199
	goto L53
L181:
	;
	if v835 != 0 {
		goto L263
	} else {
		goto L264
	}
L182:
	;
	v1032 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v1032)
	goto L181
L183:
	;
	if v206 != 0 {
		goto L218
	} else {
		goto L219
	}
L184:
	;
	v836 = v202 + int32(48) + v217<<(uint(int32(3))%32)
	goto L186
L185:
	;
	v836 = v826
	goto L186
L186:
	;
	v837 = *(*float64)(unsafe.Add(mBase, uint32(v836)))
	v838 = base.I64_reinterpret_f64(v837)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v838&int64(9223372036854775807)) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+320)) = int32(_a_F_dopr_0)
	v846 = int32(0)
	v924 = v846
	v925 = int32(3)
	v926 = v846
	goto L183
L188:
	;
	goto L189
L189:
	;
	v850 = float64(0)
	if base.B2i32(v838 != int64(0))&base.F64_eq(v837, v850)|base.F64_lt(v837, v850) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v861 = base.F64_neg(v837)
	v862 = int32(45)
	goto L192
L191:
	;
	if v214 != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if base.F64_eq(base.F64_abs(v861), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v860 = int32(43)
	goto L195
L194:
	;
	v860 = int32(0)
	goto L195
L195:
	;
	v861 = v837
	v862 = v860
	goto L192
L196:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dopr[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+328)) = uint8(v868)
	v871 = *(*int64)(unsafe.Add(mBase, _c_F_dopr[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v202)+320)) = v871
	v924 = int32(0)
	v925 = int32(8)
	v926 = v862
	goto L183
L197:
	;
	goto L198
L198:
	;
	if v206 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	if v921 < int32(0) {
		goto L182
	} else {
		goto L217
	}
L200:
	;
	v874 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1356)) = uint8(v874)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1355)) = uint8(v201)
	v877 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1354)) = uint8(v877)
	v879 = int32(_a_F_dopr_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v202)+1352)) = uint16(v879)
	*(*float64)(unsafe.Add(mBase, uint32(v202)+40)) = v861
	if v206 != 0 {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L202
L202:
	;
	v903 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1354)) = uint8(v903)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1353)) = uint8(v201)
	v907 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1352)) = uint8(v907)
	*(*float64)(unsafe.Add(mBase, uint32(v202)+16)) = v861
	v917 = F_snprintf(m, v202+int32(320), int32(1024), v202+int32(1352), v202+int32(16))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L32
	} else {
		goto L216
	}
L203:
	;
	v883 = v208
	goto L205
L204:
	;
	v883 = v215
	goto L205
L205:
	;
	if v212 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v884 = v215
	goto L208
L207:
	;
	v884 = v883
	goto L208
L208:
	;
	v885 = int32(0)
	if v885 < v884 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v888 = v884
	goto L211
L210:
	;
	v888 = v885
	goto L211
L211:
	;
	if int32(350) <= v888 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v891 = int32(350)
	goto L214
L213:
	;
	v891 = v888
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+32)) = v891
	v901 = F_snprintf(m, v202+int32(320), int32(1024), v202+int32(1352), v202+int32(32))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L32
	} else {
		goto L215
	}
L215:
	;
	v919 = v888 - v891
	v921 = v901
	goto L199
L216:
	;
	v919 = v903
	v921 = v917
	goto L199
L217:
	;
	v924 = v919
	v925 = v921
	v926 = v862
	goto L183
L218:
	;
	v929 = v211
	goto L220
L219:
	;
	v929 = v208
	goto L220
L220:
	;
	if v212 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v930 = v211
	goto L223
L222:
	;
	v930 = v929
	goto L223
L223:
	;
	v932 = v930 - (v924 + v925)
	v933 = int32(0)
	if v933 < v932 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v936 = v932
	goto L226
L225:
	;
	v936 = v933
	goto L226
L226:
	;
	if v203 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v938 = int32(0) - v936
	goto L229
L228:
	;
	v938 = v936
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+316)) = v938
	F_leading_pad(m, v213, v926, v202+int32(316), v197)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L32
	} else {
		goto L230
	}
L230:
	;
	if int32(0) < v924 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v202)+316))
	if int32(0) <= v1023 {
		goto L181
	} else {
		goto L258
	}
L232:
	;
	v947 = v202 + int32(320)
	v951 = F_strlen(m, v947)
	mBase = m.M
	v958 = v951 + int32(1)
	goto L238
L233:
	;
	goto L234
L234:
	;
	F_dostr(m, v202+int32(320), v925, v197)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L32
	} else {
		goto L257
	}
L235:
	;
	F_dostr(m, v202+int32(320), v925, v197)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L32
	} else {
		goto L255
	}
L236:
	;
	if v970 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L237:
	;
	goto L236
L238:
	;
	v960 = int32(0)
	if v958 == v960 {
		v970 = v960
		goto L237
	} else {
		goto L240
	}
L239:
	;
	v970 = v965
	goto L237
L240:
	;
	v964 = v958 - int32(1)
	v965 = v947 + v964
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965))))
	if v966 != int32(101) {
		v958 = v964
		goto L238
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v976 = F_strlen(m, v947)
	mBase = m.M
	v983 = v976 + int32(1)
	goto L247
L243:
	;
	v998 = v970
	goto L244
L244:
	;
	v1000 = v202 + int32(320)
	v1001 = v998 - v1000
	F_dostr(m, v1000, v1001, v197)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L32
	} else {
		goto L252
	}
L245:
	;
	if v995 == int32(0) {
		goto L235
	} else {
		goto L251
	}
L246:
	;
	goto L245
L247:
	;
	v985 = int32(0)
	if v983 == v985 {
		v995 = v985
		goto L246
	} else {
		goto L249
	}
L248:
	;
	v995 = v990
	goto L246
L249:
	;
	v989 = v983 - int32(1)
	v990 = v947 + v989
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	if v991 != int32(69) {
		v983 = v989
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v998 = v995
	goto L244
L252:
	;
	F_dopr_outchmulti(m, int32(48), v924, v197)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L32
	} else {
		goto L253
	}
L253:
	;
	F_dostr(m, v998, v925-v1001, v197)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L32
	} else {
		goto L254
	}
L254:
	;
	goto L231
L255:
	;
	F_dopr_outchmulti(m, int32(48), v924, v197)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L32
	} else {
		goto L256
	}
L256:
	;
	goto L231
L257:
	;
	goto L231
L258:
	;
	F_dopr_outchmulti(m, int32(32), int32(0)-v1023, v197)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L32
	} else {
		goto L259
	}
L259:
	;
	if v835 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1031 = v199
	goto L262
L261:
	;
	v1031 = v828
	goto L262
L262:
	;
	v1259 = v1031
	goto L53
L263:
	;
	v1040 = v199
	goto L265
L264:
	;
	v1040 = v828
	goto L265
L265:
	;
	v1259 = v1040
	goto L53
L266:
	;
	v1048 = v202 + int32(48) + v217<<(uint(int32(3))%32)
	goto L268
L267:
	;
	v1048 = v199
	goto L268
L268:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v1049
	if v1047 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1053 = int32(0)
	goto L271
L270:
	;
	v1053 = int32(4)
	goto L271
L271:
	;
	v1054 = v199 + v1053
	v1059 = F_snprintf(m, v202+int32(320), int32(64), int32(_a_F_dopr_2), v202)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L32
	} else {
		goto L272
	}
L272:
	;
	if v1059 < int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1063 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v1063)
	v1259 = v1054
	goto L53
L274:
	;
	goto L275
L275:
	;
	F_dostr(m, v202+int32(320), v1059, v197)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L32
	} else {
		goto L276
	}
L276:
	;
	v1259 = v1054
	goto L53
L277:
	;
	v1069 = v211
	goto L279
L278:
	;
	v1069 = v208
	goto L279
L279:
	;
	if v212 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1070 = v211
	goto L282
L281:
	;
	v1070 = v1069
	goto L282
L282:
	;
	v1077 = v209 & int32(1)
	if v1077 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1078 = v202 + int32(48) + v217<<(uint(int32(3))%32)
	goto L285
L284:
	;
	v1078 = v199
	goto L285
L285:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	if v1079 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1081 = v1079
	goto L288
L287:
	;
	v1081 = int32(_a_F_dopr_3)
	goto L288
L288:
	;
	if v1077 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1084 = int32(0)
	goto L291
L290:
	;
	v1084 = int32(4)
	goto L291
L291:
	;
	if v206 != 0 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1093 = v199 + v1084
	v1094 = int32(0)
	v1095 = v1070 - v1092
	if v1094 < v1095 {
		goto L303
	} else {
		goto L304
	}
L293:
	;
	if v212 != 0 {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	goto L295
L295:
	;
	v1091 = F_strlen(m, v1081)
	mBase = m.M
	v1092 = v1091
	goto L292
L296:
	;
	v1085 = v215
	goto L298
L297:
	;
	v1085 = v208
	goto L298
L298:
	;
	v1088 = F_memchr(m, v1081, int32(0), v1085)
	mBase = m.M
	if v1088 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1092 = v1090
	goto L292
L300:
	;
	v1090 = v1088 - v1081
	goto L302
L301:
	;
	v1090 = v1085
	goto L302
L302:
	;
	goto L299
L303:
	;
	v1099 = v1095
	goto L305
L304:
	;
	v1099 = v1094
	goto L305
L305:
	;
	if v203 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1101 = v1094 - v1099
	goto L308
L307:
	;
	v1101 = v1099
	goto L308
L308:
	;
	if int32(0) < v1101 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	F_dopr_outchmulti(m, int32(32), v1101, v197)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L32
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	F_dostr(m, v1081, v1092, v197)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L32
	} else {
		goto L314
	}
L312:
	;
	F_dostr(m, v1081, v1092, v197)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L32
	} else {
		goto L313
	}
L313:
	;
	v1259 = v1093
	goto L53
L314:
	;
	if int32(0) <= v1101 {
		v1259 = v1093
		goto L53
	} else {
		goto L315
	}
L315:
	;
	F_dopr_outchmulti(m, int32(32), int32(0)-v1101, v197)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L32
	} else {
		goto L316
	}
L316:
	;
	v1259 = v1093
	goto L53
L317:
	;
	v1125 = v202 + int32(48) + v217<<(uint(int32(3))%32)
	goto L319
L318:
	;
	v1125 = v199
	goto L319
L319:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1125)))
	if v206 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1128 = v211
	goto L322
L321:
	;
	v1128 = v208
	goto L322
L322:
	;
	if v212 != 0 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1129 = v211
	goto L325
L324:
	;
	v1129 = v1128
	goto L325
L325:
	;
	v1131 = v1129 - int32(1)
	v1132 = int32(0)
	if v1132 < v1131 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1135 = v1131
	goto L328
L327:
	;
	v1135 = v1132
	goto L328
L328:
	;
	if v203 != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1137 = int32(0) - v1135
	goto L331
L330:
	;
	v1137 = v1135
	goto L331
L331:
	;
	if int32(0) < v1137 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	F_dopr_outchmulti(m, int32(32), v1137, v197)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L32
	} else {
		goto L335
	}
L333:
	;
	v1144 = v1137
	goto L334
L334:
	;
	if v1124 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	v1144 = int32(0)
	goto L334
L336:
	;
	v1147 = int32(0)
	goto L338
L337:
	;
	v1147 = int32(4)
	goto L338
L338:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v1149 = int32(0)
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if base.B2i32(v1148 == v1149)|base.B2i32(base.Ui32(v1151) < base.Ui32(v1148)) == v1149 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v1189 = v199 + v1147
	if int32(0) <= v1144 {
		v1259 = v1189
		goto L53
	} else {
		goto L351
	}
L340:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	if v1156 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v1179 = v1151
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v1179 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1179))) = uint8(v1126)
	goto L339
L343:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v1159 + int32(1)
	goto L339
L344:
	;
	goto L345
L345:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)))
	if v1163 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v1179 = v1178
	goto L342
L347:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v1151 == v1164 {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1167 = v1151 - v1164
	v1168 = F_fwrite(m, v1164, int32(1), v1167, v1156)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L32
	} else {
		goto L349
	}
L349:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v1168 + v1170
	if v1167 == v1168 {
		goto L346
	} else {
		goto L350
	}
L350:
	;
	v1174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v1174)
	goto L346
L351:
	;
	F_dopr_outchmulti(m, int32(32), int32(0)-v1144, v197)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L32
	} else {
		goto L352
	}
L352:
	;
	v1259 = v1189
	goto L53
L353:
	;
	v1197 = v208
	goto L355
L354:
	;
	v1197 = v215
	goto L355
L355:
	;
	if v212 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1198 = v215
	goto L358
L357:
	;
	v1198 = v1197
	goto L358
L358:
	;
	if v206 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1199 = v211
	goto L361
L360:
	;
	v1199 = v208
	goto L361
L361:
	;
	if v212 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1200 = v211
	goto L364
L363:
	;
	v1200 = v1199
	goto L364
L364:
	;
	if v209&int32(1) != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1207 = v202 + int32(48) + v217<<(uint(int32(3))%32)
	if v218 != 0 {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	goto L367
L367:
	;
	if v218 != 0 {
		goto L373
	} else {
		goto L374
	}
L368:
	;
	v1208 = *(*int64)(unsafe.Add(mBase, uint32(v1207)))
	F_fmtint(m, v1208, v230, v214, v203, v1200, v213, v1198, v206, v197)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L32
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v1211 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1207))))
	F_fmtint(m, v1211, v230, v214, v203, v1200, v213, v1198, v206, v197)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L32
	} else {
		goto L372
	}
L371:
	;
	v1259 = v199
	goto L53
L372:
	;
	v1259 = v199
	goto L53
L373:
	;
	v1217 = (v199 + int32(7)) & int32(-8)
	v1220 = *(*int64)(unsafe.Add(mBase, uint32(v1217)))
	F_fmtint(m, v1220, v230, v214, v203, v1200, v213, v1198, v206, v197)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L32
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	v1223 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199))))
	F_fmtint(m, v1223, v230, v214, v203, v1200, v213, v1198, v206, v197)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L32
	} else {
		goto L377
	}
L376:
	;
	v1259 = v1217 + int32(8)
	goto L53
L377:
	;
	v1259 = v199 + int32(4)
	goto L53
L378:
	;
	v1228 = v208
	goto L380
L379:
	;
	v1228 = v215
	goto L380
L380:
	;
	if v212 != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1229 = v215
	goto L383
L382:
	;
	v1229 = v1228
	goto L383
L383:
	;
	if v206 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1230 = v211
	goto L386
L385:
	;
	v1230 = v208
	goto L386
L386:
	;
	if v212 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1231 = v211
	goto L389
L388:
	;
	v1231 = v1230
	goto L389
L389:
	;
	if v209&int32(1) != 0 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1238 = v202 + int32(48) + v217<<(uint(int32(3))%32)
	if v218 != 0 {
		goto L393
	} else {
		goto L394
	}
L391:
	;
	goto L392
L392:
	;
	if v218 != 0 {
		goto L398
	} else {
		goto L399
	}
L393:
	;
	v1239 = *(*int64)(unsafe.Add(mBase, uint32(v1238)))
	F_fmtint(m, v1239, v230, v214, v203, v1231, v213, v1229, v206, v197)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L32
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1242 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1238))))
	F_fmtint(m, v1242, v230, v214, v203, v1231, v213, v1229, v206, v197)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L32
	} else {
		goto L397
	}
L396:
	;
	v1259 = v199
	goto L53
L397:
	;
	v1259 = v199
	goto L53
L398:
	;
	v1248 = (v199 + int32(7)) & int32(-8)
	v1251 = *(*int64)(unsafe.Add(mBase, uint32(v1248)))
	F_fmtint(m, v1251, v230, v214, v203, v1231, v213, v1229, v206, v197)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L32
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v1254 = int64(*(*int32)(unsafe.Add(mBase, uint32(v199))))
	F_fmtint(m, v1254, v230, v214, v203, v1231, v213, v1229, v206, v197)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L32
	} else {
		goto L402
	}
L401:
	;
	v1259 = v1248 + int32(8)
	goto L53
L402:
	;
	v1259 = v199 + int32(4)
	goto L53
L403:
	;
	v1534 = v202
	goto L3
L404:
	;
	v1275 = int32(1)
	goto L406
L405:
	;
	v1275 = v218
	goto L406
L406:
	;
	v1463 = v206
	v1466 = v209
	v1468 = v211
	v1469 = v212
	v1470 = v213
	v1471 = v214
	v1472 = v215
	v1475 = v1275
	v1478 = v221
	v1479 = int32(1)
	v1484 = v208
	goto L45
L407:
	;
	v1284 = int32(1)
	v1285 = v199
	goto L408
L408:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(320)+v1284<<(uint(int32(2))%32))))
	switch v1316 {
	case 0:
		goto L48
	case 1, 2, 5:
		goto L411
	case 3:
		goto L413
	case 4:
		goto L412
	default:
		v1352 = v1285
		goto L410
	}
L409:
	;
	goto L49
L410:
	;
	v1354 = v1284 + int32(1)
	if v1354 <= v292 {
		v1284 = v1354
		v1285 = v1352
		goto L408
	} else {
		goto L414
	}
L411:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1285)))
	*(*int32)(unsafe.Add(mBase, uint32(v202+int32(48)+v1284<<(uint(int32(3))%32)))) = v1348
	v1352 = v1285 + int32(4)
	goto L410
L412:
	;
	v1338 = (v1285 + int32(7)) & int32(-8)
	v1339 = *(*float64)(unsafe.Add(mBase, uint32(v1338)))
	*(*float64)(unsafe.Add(mBase, uint32(v202+int32(48)+v1284<<(uint(int32(3))%32)))) = v1339
	v1352 = v1338 + int32(8)
	goto L410
L413:
	;
	v1325 = (v1285 + int32(7)) & int32(-8)
	v1326 = *(*int64)(unsafe.Add(mBase, uint32(v1325)))
	*(*int64)(unsafe.Add(mBase, uint32(v202+int32(48)+v1284<<(uint(int32(3))%32)))) = v1326
	v1352 = v1325 + int32(8)
	goto L410
L414:
	;
	goto L409
L415:
	;
	v1487 = v199
	v1488 = v203
	v1492 = int32(0)
	v1494 = v206
	v1497 = v1386
	v1499 = v211
	v1500 = v212
	v1501 = v213
	v1502 = v214
	v1503 = v215
	v1505 = v208
	v1506 = v218
	v1509 = v1387
	v1510 = v222
	goto L44
L416:
	;
	goto L417
L417:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(48)+v208<<(uint(int32(3))%32))))
	if v206 != 0 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1397 = int32(0)
	v1399 = base.B2i32(v1397 <= v1396)
	if v1397 <= v1396 {
		goto L421
	} else {
		goto L422
	}
L419:
	;
	goto L420
L420:
	;
	v1403 = v1396 >> (uint(int32(31)) % 32)
	if v1396 < int32(0) {
		goto L424
	} else {
		goto L425
	}
L421:
	;
	v1400 = v1396
	goto L423
L422:
	;
	v1400 = v1397
	goto L423
L423:
	;
	v1463 = v1399
	v1466 = v1386
	v1468 = v211
	v1469 = v212
	v1470 = v213
	v1471 = v214
	v1472 = v1400
	v1475 = v218
	v1478 = v1387
	v1479 = v222
	v1484 = int32(0)
	goto L45
L424:
	;
	v1409 = int32(1)
	goto L426
L425:
	;
	v1409 = v203
	goto L426
L426:
	;
	v1410 = int32(0)
	v1487 = v199
	v1488 = v1409
	v1492 = v1410
	v1494 = v1410
	v1497 = v1386
	v1499 = v1396 ^ v1403 - v1403
	v1500 = v212
	v1501 = v213
	v1502 = v214
	v1503 = v215
	v1505 = v217
	v1506 = v218
	v1509 = v1387
	v1510 = v222
	goto L44
L427:
	;
	v1518 = v1516
	goto L429
L428:
	;
	v1518 = int32(_a_F_dopr_3)
	goto L429
L429:
	;
	v1519 = F_strlen(m, v1518)
	mBase = m.M
	F_dostr(m, v1518, v1519, v37)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L32
	} else {
		goto L430
	}
L430:
	;
	v1526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+20)))
	if v1526 == int32(0) {
		v38 = v177 + int32(2)
		v39 = v39 + int32(4)
		v60 = v178
		goto L1
	} else {
		goto L431
	}
L431:
	;
	v1534 = v42
	goto L3
}
func F_dutch_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v840 int32
	_ = v840
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v6
	goto L2
L1:
	;
	return v1034
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 <= v8 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	if v71 == v6 {
		v100 = v6
		goto L31
	} else {
		goto L32
	}
L4:
	;
	goto L3
L5:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v79
	goto L2
L6:
	;
	if v71 <= v69 {
		goto L4
	} else {
		goto L30
	}
L7:
	;
	v30 = F_find_among(m, l0, int32(_a_F_dutch_ISO_8859_1_stem_0), int32(11))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
	v69 = v8
	v71 = v13
	goto L6
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v8))))
	v18 = int32(224)
	if v17&v18 != v18 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if int32(1)<<(uint(v17)%32)&int32(340306450) != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v34
	switch v30 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	default:
		goto L5
	}
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v69 = v34
	v71 = v68
	goto L6
L15:
	;
	v64 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L28
	}
L16:
	;
	v58 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v52 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v46 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v40 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_5))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v40 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v1034 = v40
	goto L1
L22:
	;
	if int32(0) <= v46 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v1034 = v46
	goto L1
L24:
	;
	if int32(0) <= v52 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v1034 = v52
	goto L1
L26:
	;
	if int32(0) <= v58 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v1034 = v58
	goto L1
L28:
	;
	if int32(0) <= v64 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v1034 = v64
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69 + int32(1)
	goto L5
L31:
	;
	v103 = v100
	goto L36
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v6))))
	if v85 != int32(121) {
		v100 = v6
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v88 = int32(1)
	v89 = v6 + v88
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v89
	v94 = F_slice_from_s(m, l0, v88, int32(_a_F_dutch_ISO_8859_1_stem_6))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v94 < int32(0) {
		v1034 = v94
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = v98
	goto L31
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v115 < v114 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v262)+8)) = v256
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v267 = v265 + int32(3)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v268 < v267 {
		v504 = v265
		goto L85
	} else {
		goto L86
	}
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v158 != 0 {
		v256 = v159
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v117 = v114
	goto L41
L40:
	;
	v117 = v115
	goto L41
L41:
	;
	goto L43
L42:
	;
	v158 = v153
	goto L38
L43:
	;
	if v114 == v117 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v153 = int32(0)
	goto L42
L45:
	;
	v158 = int32(-1)
	goto L38
L46:
	;
	goto L47
L47:
	;
	v129 = int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v114))))
	if int32(232) < v132 {
		v153 = v129
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v134 = v132 - int32(97)
	if v134 < int32(0) {
		v153 = v129
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v140)>>(uint(v134&int32(7))%32))&int32(1) == int32(0) {
		v153 = v129
		goto L42
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114 + int32(1)
	goto L51
L51:
	;
	goto L44
L52:
	;
	if v103 < v256 {
		goto L82
	} else {
		goto L83
	}
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v160
	if v160 == v159 {
		v228 = v159
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	goto L36
L55:
	;
	v248 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_7))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L80
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160
	if v160 == v228 {
		goto L74
	} else {
		goto L75
	}
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v160))))
	if v165 != int32(105) {
		v228 = v159
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v169 = v160 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v169
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v181 < v169 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v224 == int32(0) {
		goto L55
	} else {
		goto L73
	}
L60:
	;
	v183 = v169
	goto L62
L61:
	;
	v183 = v181
	goto L62
L62:
	;
	goto L64
L63:
	;
	v224 = v219
	goto L59
L64:
	;
	if v169 == v183 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v219 = int32(0)
	goto L63
L66:
	;
	v224 = int32(-1)
	goto L59
L67:
	;
	goto L68
L68:
	;
	v195 = int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v169))))
	if int32(232) < v198 {
		v219 = v195
		goto L63
	} else {
		goto L69
	}
L69:
	;
	v200 = v198 - int32(97)
	if v200 < int32(0) {
		v219 = v195
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v200)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v206)>>(uint(v200&int32(7))%32))&int32(1) == int32(0) {
		v219 = v195
		goto L63
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160 + int32(2)
	goto L72
L72:
	;
	goto L65
L73:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v228 = v227
	goto L56
L74:
	;
	v256 = v160
	goto L52
L75:
	;
	goto L76
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v160))))
	if v233 != int32(121) {
		v256 = v228
		goto L52
	} else {
		goto L77
	}
L77:
	;
	v236 = int32(1)
	v237 = v160 + v236
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v237
	v242 = F_slice_from_s(m, l0, v236, int32(_a_F_dutch_ISO_8859_1_stem_8))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if int32(0) <= v242 {
		goto L54
	} else {
		goto L79
	}
L79:
	;
	v1034 = v242
	goto L1
L80:
	;
	if v248 < int32(0) {
		v1034 = v248
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L54
L82:
	;
	v259 = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259
	v103 = v259
	goto L36
L83:
	;
	goto L84
L84:
	;
	goto L37
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v506
	if v506 <= v6 {
		v611 = v504
		goto L152
	} else {
		goto L153
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v267
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v265
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v280 < v265 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v320 < int32(0) {
		v504 = v265
		goto L85
	} else {
		goto L102
	}
L88:
	;
	v282 = v265
	goto L90
L89:
	;
	v282 = v280
	goto L90
L90:
	;
	v289 = v265
	goto L92
L91:
	;
	v320 = v300
	goto L87
L92:
	;
	if v289 == v282 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v320 = int32(-1)
	goto L87
L95:
	;
	goto L96
L96:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v289))))
	if int32(232) < v295 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v312 = v289 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v312
	v289 = v312
	goto L92
L98:
	;
	v297 = v295 - int32(97)
	if v297 < int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v300 = int32(1)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v297)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v304)>>(uint(v297&int32(7))%32))&v300 != 0 {
		goto L91
	} else {
		goto L100
	}
L100:
	;
	goto L97
L102:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v324 = v323 + v320
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v335 < v324 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v378 < int32(0) {
		v504 = v265
		goto L85
	} else {
		goto L117
	}
L104:
	;
	v337 = v324
	goto L106
L105:
	;
	v337 = v335
	goto L106
L106:
	;
	v343 = v324
	goto L108
L107:
	;
	v378 = int32(1)
	goto L103
L108:
	;
	if v343 == v337 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v378 = int32(-1)
	goto L103
L111:
	;
	goto L112
L112:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v343))))
	if int32(232) < v352 {
		goto L107
	} else {
		goto L113
	}
L113:
	;
	v354 = v352 - int32(97)
	if v354 < int32(0) {
		goto L107
	} else {
		goto L114
	}
L114:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v354)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v360)>>(uint(v354&int32(7))%32))&int32(1) == int32(0) {
		goto L107
	} else {
		goto L115
	}
L115:
	;
	v369 = v343 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369
	v343 = v369
	goto L108
L117:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v382 = v381 + v378
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v385 < v382 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v387 = v382
	goto L120
L119:
	;
	v387 = v385
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+8)) = v387
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v397 < v396 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v437 < int32(0) {
		v504 = v385
		goto L85
	} else {
		goto L136
	}
L122:
	;
	v399 = v396
	goto L124
L123:
	;
	v399 = v397
	goto L124
L124:
	;
	v406 = v396
	goto L126
L125:
	;
	v437 = v417
	goto L121
L126:
	;
	if v406 == v399 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v437 = int32(-1)
	goto L121
L129:
	;
	goto L130
L130:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v406))))
	if int32(232) < v412 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v429 = v406 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429
	v406 = v429
	goto L126
L132:
	;
	v414 = v412 - int32(97)
	if v414 < int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v417 = int32(1)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v414)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v421)>>(uint(v414&int32(7))%32))&v417 != 0 {
		goto L125
	} else {
		goto L134
	}
L134:
	;
	goto L131
L136:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v441 = v440 + v437
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v441
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v452 < v441 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v495 < int32(0) {
		v504 = v385
		goto L85
	} else {
		goto L151
	}
L138:
	;
	v454 = v441
	goto L140
L139:
	;
	v454 = v452
	goto L140
L140:
	;
	v460 = v441
	goto L142
L141:
	;
	v495 = int32(1)
	goto L137
L142:
	;
	if v460 == v454 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v495 = int32(-1)
	goto L137
L145:
	;
	goto L146
L146:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467+v460))))
	if int32(232) < v469 {
		goto L141
	} else {
		goto L147
	}
L147:
	;
	v471 = v469 - int32(97)
	if v471 < int32(0) {
		goto L141
	} else {
		goto L148
	}
L148:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v471)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v477)>>(uint(v471&int32(7))%32))&int32(1) == int32(0) {
		goto L141
	} else {
		goto L149
	}
L149:
	;
	v486 = v460 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v486
	v460 = v486
	goto L142
L151:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+4)) = v499 + v495
	v504 = v385
	goto L85
L152:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v612
	v614 = F_r_e_ending_1(m, l0)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L12
	} else {
		goto L185
	}
L153:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v512 = int32(1)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510+v506-v512))))
	if base.B2i32(v514&int32(224) != int32(96))|base.B2i32(v512<<(uint(v514)%32)&int32(_a_F_dutch_ISO_8859_1_stem_9) == int32(0)) != 0 {
		v611 = v504
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v528 = F_find_among_b(m, l0, int32(_a_F_dutch_ISO_8859_1_stem_10), int32(5))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	if v528 == int32(0) {
		v611 = v504
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v532
	switch v528 - int32(1) {
	case 0:
		goto L159
	case 1:
		goto L158
	case 2:
		goto L157
	default:
		v611 = v504
		goto L152
	}
L157:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+8))
	if v532 < v552 {
		goto L167
	} else {
		goto L168
	}
L158:
	;
	v547 = F_r_en_ending_1(m, l0)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L12
	} else {
		goto L165
	}
L159:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	if v532 < v537 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v611 = int32(0)
	goto L152
L161:
	;
	goto L162
L162:
	;
	v542 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_ISO_8859_1_stem_11))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L12
	} else {
		goto L163
	}
L163:
	;
	if v542 < int32(0) {
		v1034 = v542
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v611 = int32(1)
	goto L152
L165:
	;
	if int32(0) <= v547 {
		v611 = v547
		goto L152
	} else {
		goto L166
	}
L166:
	;
	v1034 = v547
	goto L1
L167:
	;
	v611 = int32(0)
	goto L152
L168:
	;
	goto L169
L169:
	;
	v555 = int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L172
L170:
	;
	if v604 != 0 {
		v611 = v555
		goto L152
	} else {
		goto L182
	}
L171:
	;
	v604 = v601
	goto L170
L172:
	;
	if v563 <= v564 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v601 = int32(0)
	goto L171
L174:
	;
	v604 = int32(-1)
	goto L170
L175:
	;
	goto L176
L176:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v563-int32(1)))))
	if int32(232) < v579 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v563 - int32(1)
	goto L181
L178:
	;
	v581 = v579 - int32(97)
	if v581 < int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v584 = int32(1)
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v581)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v588)>>(uint(v581&int32(7))%32))&v584 != 0 {
		v601 = v584
		goto L171
	} else {
		goto L180
	}
L180:
	;
	goto L177
L181:
	;
	goto L173
L182:
	;
	v605 = F_slice_del(m, l0)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L12
	} else {
		goto L183
	}
L183:
	;
	if v605 < int32(0) {
		v1034 = v605
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v611 = v555
	goto L152
L185:
	;
	if v614 < int32(0) {
		v1034 = v614
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v618
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v618
	v621 = int32(4)
	v623 = int32(0)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v618-v626 < v621 {
		v636 = v623
		goto L194
	} else {
		goto L195
	}
L187:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v975
	v978 = v975
	goto L302
L188:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v830
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L271
L189:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	if v715 < v818 {
		goto L188
	} else {
		goto L265
	}
L190:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	if v715 < v811 {
		goto L188
	} else {
		goto L262
	}
L191:
	;
	if int32(0) <= v805 {
		goto L187
	} else {
		goto L261
	}
L192:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v688
	v692 = v688 - int32(1)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v692 <= v693 {
		goto L188
	} else {
		goto L218
	}
L193:
	;
	if v636 == int32(0) {
		v687 = v611
		goto L192
	} else {
		goto L197
	}
L194:
	;
	goto L193
L195:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v632 = F_memcmp(m, v629+v618-v621, int32(_a_F_dutch_ISO_8859_1_stem_12), v621)
	mBase = m.M
	if v632 != 0 {
		v636 = v623
		goto L194
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v618 - v621
	v636 = int32(1)
	goto L194
L197:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v639 < v642 {
		v687 = v611
		goto L192
	} else {
		goto L198
	}
L198:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v644 < v639 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v639-int32(1)))))
	if v650 == int32(99) {
		v687 = v611
		goto L192
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v653 = F_slice_del(m, l0)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L12
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	if v653 < int32(0) {
		v1034 = v653
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v657
	v659 = int32(2)
	v661 = int32(0)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v657-v664 < v659 {
		v674 = v661
		goto L206
	} else {
		goto L207
	}
L205:
	;
	if v674 == int32(0) {
		v687 = v611
		goto L192
	} else {
		goto L209
	}
L206:
	;
	goto L205
L207:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v670 = F_memcmp(m, v667+v657-v659, int32(_a_F_dutch_ISO_8859_1_stem_13), v659)
	mBase = m.M
	if v670 != 0 {
		v674 = v661
		goto L206
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v657 - v659
	v674 = int32(1)
	goto L206
L209:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v677
	v679 = F_r_en_ending_1(m, l0)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L12
	} else {
		goto L210
	}
L210:
	;
	v682 = base.B2i32(v679 < int32(0))
	if v679 < int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v683 = v679
	goto L213
L212:
	;
	v683 = v611
	goto L213
L213:
	;
	if v679 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v684 = v683
	goto L216
L215:
	;
	v684 = v611
	goto L216
L216:
	;
	if v679 < int32(0) {
		v805 = v684
		goto L191
	} else {
		goto L217
	}
L217:
	;
	v687 = v684
	goto L192
L218:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695+v692))))
	if base.B2i32(v697&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v697)%32)&int32(_a_F_dutch_ISO_8859_1_stem_14) == int32(0)) != 0 {
		goto L188
	} else {
		goto L219
	}
L219:
	;
	v711 = F_find_among_b(m, l0, int32(_a_F_dutch_ISO_8859_1_stem_15), int32(6))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L12
	} else {
		goto L220
	}
L220:
	;
	if v711 == int32(0) {
		goto L188
	} else {
		goto L221
	}
L221:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v715
	switch v711 - int32(1) {
	case 0:
		goto L224
	case 1:
		goto L223
	case 2:
		goto L222
	case 3:
		goto L190
	case 4:
		goto L189
	default:
		goto L188
	}
L222:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	if v715 < v791 {
		goto L188
	} else {
		goto L250
	}
L223:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+4))
	if v715 < v775 {
		goto L188
	} else {
		goto L243
	}
L224:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v715 < v720 {
		goto L188
	} else {
		goto L225
	}
L225:
	;
	v722 = F_slice_del(m, l0)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L12
	} else {
		goto L226
	}
L226:
	;
	if v722 < int32(0) {
		v1034 = v722
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v726
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v729 = int32(2)
	v731 = int32(0)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v726-v734 < v729 {
		v744 = v731
		goto L230
	} else {
		goto L231
	}
L228:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v766 + (v726 - v728)
	v770 = F_r_undouble_1(m, l0)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L12
	} else {
		goto L241
	}
L229:
	;
	if v744 == int32(0) {
		goto L228
	} else {
		goto L233
	}
L230:
	;
	goto L229
L231:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v740 = F_memcmp(m, v737+v726-v729, int32(_a_F_dutch_ISO_8859_1_stem_16), v729)
	mBase = m.M
	if v740 != 0 {
		v744 = v731
		goto L230
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v726 - v729
	v744 = int32(1)
	goto L230
L233:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	if v747 < v750 {
		goto L228
	} else {
		goto L234
	}
L234:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v752 < v747 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v754+v747-int32(1)))))
	if v758 == int32(101) {
		goto L228
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v761 = F_slice_del(m, l0)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L12
	} else {
		goto L239
	}
L238:
	;
	goto L237
L239:
	;
	if int32(0) <= v761 {
		goto L188
	} else {
		goto L240
	}
L240:
	;
	v1034 = v761
	goto L1
L241:
	;
	if int32(0) <= v770 {
		goto L188
	} else {
		goto L242
	}
L242:
	;
	v1034 = v770
	goto L1
L243:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v777 < v715 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779+v715-int32(1)))))
	if v783 == int32(101) {
		goto L188
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v786 = F_slice_del(m, l0)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L12
	} else {
		goto L248
	}
L247:
	;
	goto L246
L248:
	;
	if int32(0) <= v786 {
		goto L188
	} else {
		goto L249
	}
L249:
	;
	v1034 = v786
	goto L1
L250:
	;
	v793 = F_slice_del(m, l0)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L12
	} else {
		goto L251
	}
L251:
	;
	if v793 < int32(0) {
		v1034 = v793
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v797 = F_r_e_ending_1(m, l0)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L12
	} else {
		goto L253
	}
L253:
	;
	if int32(0) <= v797 {
		goto L188
	} else {
		goto L254
	}
L254:
	;
	if v797 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v803 = v797
	goto L257
L256:
	;
	v803 = v687
	goto L257
L257:
	;
	if v797 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v804 = v803
	goto L260
L259:
	;
	v804 = v687
	goto L260
L260:
	;
	v805 = v804
	goto L191
L261:
	;
	v1034 = v805
	goto L1
L262:
	;
	v813 = F_slice_del(m, l0)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L12
	} else {
		goto L263
	}
L263:
	;
	if int32(0) <= v813 {
		goto L188
	} else {
		goto L264
	}
L264:
	;
	v1034 = v813
	goto L1
L265:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v817)+12))
	if v820 == int32(0) {
		goto L188
	} else {
		goto L266
	}
L266:
	;
	v823 = F_slice_del(m, l0)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L12
	} else {
		goto L267
	}
L267:
	;
	if v823 < int32(0) {
		v1034 = v823
		goto L1
	} else {
		goto L268
	}
L268:
	;
	goto L188
L269:
	;
	if v880 != 0 {
		goto L187
	} else {
		goto L281
	}
L270:
	;
	v880 = v877
	goto L269
L271:
	;
	if v830 <= v840 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v877 = int32(0)
	goto L270
L273:
	;
	v880 = int32(-1)
	goto L269
L274:
	;
	goto L275
L275:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851+v830-int32(1)))))
	if int32(232) < v855 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v830 - int32(1)
	goto L280
L277:
	;
	v857 = v855 - int32(73)
	if v857 < int32(0) {
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v860 = int32(1)
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v857)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v864)>>(uint(v857&int32(7))%32))&v860 != 0 {
		v877 = v860
		goto L270
	} else {
		goto L279
	}
L279:
	;
	goto L276
L280:
	;
	goto L272
L281:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v883 = v881 - int32(1)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v883 <= v884 {
		goto L187
	} else {
		goto L282
	}
L282:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886+v883))))
	if base.B2i32(v888&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v888)%32)&int32(_a_F_dutch_ISO_8859_1_stem_17) == int32(0)) != 0 {
		goto L187
	} else {
		goto L283
	}
L283:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v903 = F_find_among_b(m, l0, int32(_a_F_dutch_ISO_8859_1_stem_18), int32(4))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L12
	} else {
		goto L284
	}
L284:
	;
	if v903 == int32(0) {
		goto L187
	} else {
		goto L285
	}
L285:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L288
L286:
	;
	if v955 != 0 {
		goto L187
	} else {
		goto L298
	}
L287:
	;
	v955 = v952
	goto L286
L288:
	;
	if v914 <= v915 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v952 = int32(0)
	goto L287
L290:
	;
	v955 = int32(-1)
	goto L286
L291:
	;
	goto L292
L292:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926+v914-int32(1)))))
	if int32(232) < v930 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v914 - int32(1)
	goto L297
L294:
	;
	v932 = v930 - int32(97)
	if v932 < int32(0) {
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v935 = int32(1)
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v932)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v939)>>(uint(v932&int32(7))%32))&v935 != 0 {
		v952 = v935
		goto L287
	} else {
		goto L296
	}
L296:
	;
	goto L293
L297:
	;
	goto L289
L298:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v958 = v956 + (v881 - v900)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v958
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v958 <= v961 {
		goto L187
	} else {
		goto L299
	}
L299:
	;
	v964 = v958 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964
	v967 = F_slice_del(m, l0)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L12
	} else {
		goto L300
	}
L300:
	;
	if v967 < int32(0) {
		v1034 = v967
		goto L1
	} else {
		goto L301
	}
L301:
	;
	goto L187
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v978
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v983 <= v978 {
		goto L307
	} else {
		goto L308
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v975
	v1034 = int32(1)
	goto L1
L304:
	;
	goto L303
L305:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v978 = v1030
	goto L302
L306:
	;
	if v1021 <= v1020 {
		goto L304
	} else {
		goto L318
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v978
	v1020 = v978
	v1021 = v983
	goto L306
L308:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985+v978))))
	v989 = v987 - int32(73)
	v990 = int32(0)
	if base.B2i32(v989 == v990)|base.B2i32(v989 == int32(16)) == v990 {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v999 = F_find_among(m, l0, int32(_a_F_dutch_ISO_8859_1_stem_19), int32(3))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L12
	} else {
		goto L310
	}
L310:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1001
	switch v999 - int32(1) {
	case 0:
		goto L312
	case 1:
		goto L311
	case 2:
		goto L313
	default:
		goto L305
	}
L311:
	;
	v1014 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_20))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L12
	} else {
		goto L316
	}
L312:
	;
	v1008 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_ISO_8859_1_stem_21))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L12
	} else {
		goto L314
	}
L313:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1020 = v1001
	v1021 = v1005
	goto L306
L314:
	;
	if int32(0) <= v1008 {
		goto L305
	} else {
		goto L315
	}
L315:
	;
	v1034 = v1008
	goto L1
L316:
	;
	if int32(0) <= v1014 {
		goto L305
	} else {
		goto L317
	}
L317:
	;
	v1034 = v1014
	goto L1
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1020 + int32(1)
	goto L305
}
func F_dxsyn_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v385 int32
	_ = v385
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_palloc0(m, int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(16777473)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(0)
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v12 + int32(80)
	return v16
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if int32(0) < v26 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	F_pfree(m, v220)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L122
	}
L6:
	;
	F_tsearch_readline_end(m, v12+int32(36))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L121
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L117
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L113
	}
L9:
	;
	v31 = v2
	v34 = v2
	goto L12
L10:
	;
	v211 = v2
	goto L11
L11:
	;
	if v211 == int32(0) {
		goto L3
	} else {
		goto L69
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v31<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = int32(_a_F_dxsyn_init_0)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dxsyn_init[0])))
	if base.B2i32(v47 == int32(0))|base.B2i32(v47 != v50) != 0 {
		v68 = v47
		v69 = v50
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v211 = v201
	goto L11
L14:
	;
	v203 = v31 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v203 < v204 {
		v31 = v203
		v34 = v201
		goto L12
	} else {
		goto L68
	}
L15:
	;
	if v68-v69 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	v53 = v43
	v54 = v44
	goto L18
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v58 == int32(0) {
		v68 = v58
		v69 = v57
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v68 = v58
	v69 = v57
	goto L16
L20:
	;
	v61 = int32(1)
	if v58 == v57 {
		v53 = v53 + v61
		v54 = v54 + v61
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v73 = F_defGetBoolean(m, v42)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v76 = int32(_a_F_dxsyn_init_1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dxsyn_init[1])))
	if base.B2i32(v79 == int32(0))|base.B2i32(v79 != v82) != 0 {
		v100 = v79
		v101 = v82
		goto L27
	} else {
		goto L28
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v73)
	v201 = v34
	goto L14
L26:
	;
	if v100-v101 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v85 = v43
	v86 = v76
	goto L29
L29:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v90 == int32(0) {
		v100 = v90
		v101 = v89
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v100 = v90
	v101 = v89
	goto L27
L31:
	;
	v93 = int32(1)
	if v90 == v89 {
		v85 = v85 + v93
		v86 = v86 + v93
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v105 = F_defGetBoolean(m, v42)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v108 = int32(_a_F_dxsyn_init_2)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dxsyn_init[2])))
	if base.B2i32(v111 == int32(0))|base.B2i32(v111 != v114) != 0 {
		v132 = v111
		v133 = v114
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)) = uint8(v105)
	v201 = v34
	goto L14
L37:
	;
	if v132-v133 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	v117 = v43
	v118 = v108
	goto L40
L40:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v122 == int32(0) {
		v132 = v122
		v133 = v121
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v132 = v122
	v133 = v121
	goto L38
L42:
	;
	v125 = int32(1)
	if v122 == v121 {
		v117 = v117 + v125
		v118 = v118 + v125
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v137 = F_defGetBoolean(m, v42)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v140 = int32(_a_F_dxsyn_init_3)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dxsyn_init[3])))
	if base.B2i32(v143 == int32(0))|base.B2i32(v143 != v146) != 0 {
		v164 = v143
		v165 = v146
		goto L49
	} else {
		goto L50
	}
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+10)) = uint8(v137)
	v201 = v34
	goto L14
L48:
	;
	if v164-v165 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	goto L48
L50:
	;
	v149 = v43
	v150 = v140
	goto L51
L51:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	if v154 == int32(0) {
		v164 = v154
		v165 = v153
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v164 = v154
	v165 = v153
	goto L49
L53:
	;
	v157 = int32(1)
	if v154 == v153 {
		v149 = v149 + v157
		v150 = v150 + v157
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v169 = F_defGetBoolean(m, v42)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v172 = int32(_a_F_dxsyn_init_4)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dxsyn_init[4])))
	if base.B2i32(v175 == int32(0))|base.B2i32(v175 != v178) != 0 {
		v196 = v175
		v197 = v178
		goto L60
	} else {
		goto L61
	}
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)) = uint8(v169)
	v201 = v34
	goto L14
L59:
	;
	if v196-v197 != 0 {
		goto L8
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v181 = v43
	v182 = v172
	goto L62
L62:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	if v186 == int32(0) {
		v196 = v186
		v197 = v185
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v196 = v186
	v197 = v185
	goto L60
L64:
	;
	v189 = int32(1)
	if v186 == v185 {
		v181 = v181 + v189
		v182 = v182 + v189
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v199 = F_defGetString(m, v42)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v201 = v199
	goto L14
L68:
	;
	goto L13
L69:
	;
	v218 = v12 + int32(36)
	v220 = F_get_tsearch_config_filename(m, v211, int32(_a_F_dxsyn_init_4))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v222 = F_tsearch_readline_begin(m, v218, v220)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v222 == int32(0) {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v226 = F_tsearch_readline(m, v218)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v226 == int32(0) {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v231 = v226
	v233 = int32(0)
	goto L75
L75:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v240 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	F_tsearch_readline_end(m, v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L110
	}
L77:
	;
	v241 = F_strlen(m, v231)
	mBase = m.M
	v243 = F_str_tolower(m, v231, v241, int32(100))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	v313 = v233
	goto L79
L79:
	;
	v321 = v12 + int32(36)
	v322 = F_tsearch_readline(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L108
	}
L80:
	;
	F_pfree(m, v231)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v247 = v243
	v249 = v233
	goto L82
L82:
	;
	v258 = F_find_word(m, v247, v12+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	F_pfree(m, v243)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L107
	}
L84:
	;
	if v258 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v260 == v249 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v305 = v249
	goto L87
L87:
	;
	goto L83
L88:
	;
	if v249 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v247 != v243 {
		goto L101
	} else {
		goto L102
	}
L91:
	;
	v267 = int32(16)
	goto L93
L92:
	;
	v267 = v249 << (uint(int32(1)) % 32)
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v267
	v270 = v267 << (uint(int32(3)) % 32)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v271 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v276
	goto L90
L95:
	;
	v272 = F_repalloc(m, v271, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v274 = F_palloc(m, v270)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	v276 = v272
	goto L94
L99:
	;
	v276 = v274
	goto L94
L100:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+10)))
	if v303 != 0 {
		v247 = v299
		v249 = v300
		goto L82
	} else {
		goto L106
	}
L101:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v285 = F_pnstrdup(m, v258, v283-v258)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
	if v281 != 0 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v299 = v282
	v300 = v249
	goto L100
L104:
	;
	v288 = v249 << (uint(int32(3)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v288+v289))) = v285
	v292 = F_pstrdup(m, v243)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v294+v288)+4)) = v292
	v299 = v283
	v300 = v249 + int32(1)
	goto L100
L106:
	;
	v305 = v300
	goto L87
L107:
	;
	v313 = v305
	goto L79
L108:
	;
	if v322 != 0 {
		v231 = v322
		v233 = v313
		goto L75
	} else {
		goto L109
	}
L109:
	;
	goto L76
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v313
	if v313 < int32(2) {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_pg_qsort(m, v329, v313, int32(8), int32(_a_F_dxsyn_init_5))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L5
L113:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v341
	F_errmsg(m, int32(_a_F_dxsyn_init_6), v12+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_dxsyn_init_7), int32(191), int32(_a_F_dxsyn_init_8))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v220
	F_errmsg(m, int32(_a_F_dxsyn_init_9), v12)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_dxsyn_init_7), int32(89), int32(_a_F_dxsyn_init_10))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
	goto L5
L122:
	;
	goto L3
}
