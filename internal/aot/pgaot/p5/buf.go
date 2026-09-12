package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileCreateTemp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v1 = l0
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_OpenTemporaryFile(m, v1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = F_palloc(m, int32(8240))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v14)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
				v19 = *(*int32)(unsafe.Add(mBase, _consts[10]))
				v20 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v19
				*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v20
				v28 = F_palloc(m, int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v9
					*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(0)
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v1)
					return v12
				}
			}
		}
	}
}
func F_BufFileOpenFileSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
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
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
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
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
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
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
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
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int64
	_ = v529
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	v15 = m.G0
	v17 = v15 - int32(1056)
	m.G0 = v17
	v21 = F_palloc(m, int32(64))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = int32(16)
	v34 = int32(0)
	v35 = v21
	goto L3
L3:
	;
	v40 = v34 + int32(1)
	if base.Ui32(v29) < base.Ui32(v40) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v34 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L5:
	;
	v44 = F_repalloc(m, v35, v29<<(uint(int32(3))%32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v48 = v29
	v49 = v35
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v34
	v58 = F_pg_snprintf(m, v17+int32(32), int32(1024), int32(444902), v17+int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v48 = v29 << (uint(int32(1)) % 32)
	v49 = v44
	goto L7
L9:
	;
	v60 = m.G0
	v62 = v60 - int32(2080)
	m.G0 = v62
	v67 = v17 + int32(32)
	if v67&int32(3) == int32(0) {
		v91 = v67
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v130 = v124 - int32(1636608432)
	if v67&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L11:
	;
	v124 = v116 - v67
	goto L10
L12:
	;
	v95 = v91
	goto L21
L13:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v75 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v124 = int32(0)
	goto L10
L15:
	;
	goto L16
L16:
	;
	v80 = v67
	goto L17
L17:
	;
	v84 = v80 + int32(1)
	if v84&int32(3) == int32(0) {
		v91 = v84
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v116 = v84
	goto L11
L19:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 != 0 {
		v80 = v84
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 == v104 {
		v95 = v95 + int32(4)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v110 = v95
	goto L24
L23:
	;
	goto L22
L24:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 != 0 {
		v110 = v110 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v116 = v110
	goto L11
L26:
	;
	goto L25
L27:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v390 = base.I32_rem_u_s(v384^v376-base.I32_rotl(v384, int32(24)), v389)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0+v390<<(uint(int32(2))%32))+12))
	F_TempTablespacePath(m, v62+int32(1056), v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L67
	}
L28:
	;
	v362 = int32(14)
	v364 = v358 ^ v359 - base.I32_rotl(v358, v362)
	v368 = v364 ^ v357 - base.I32_rotl(v364, int32(11))
	v372 = v368 ^ v358 - base.I32_rotl(v368, int32(25))
	v376 = v372 ^ v364 - base.I32_rotl(v372, int32(16))
	v380 = v376 ^ v368 - base.I32_rotl(v376, int32(4))
	v384 = v380 ^ v372 - base.I32_rotl(v380, v362)
	goto L27
L29:
	;
	switch v288 - int32(1) {
	case 0:
		v350 = v289
		v351 = v290
		v352 = v291
		goto L56
	case 1:
		v343 = v289
		v344 = v290
		v345 = v291
		goto L57
	case 2:
		v336 = v289
		v337 = v290
		v338 = v291
		goto L58
	case 3:
		v330 = v290
		v331 = v291
		goto L59
	case 4:
		v326 = v290
		v327 = v291
		goto L60
	case 5:
		v320 = v290
		v321 = v291
		goto L61
	case 6:
		v314 = v290
		v315 = v291
		goto L62
	case 7:
		v309 = v291
		goto L63
	case 8:
		v304 = v291
		goto L64
	case 9:
		v299 = v291
		goto L65
	case 10:
		goto L66
	default:
		v357 = v289
		v358 = v290
		v359 = v291
		goto L28
	}
L30:
	;
	v239 = v67
	v240 = v124
	v241 = v130
	v242 = v130
	v243 = v130
	goto L53
L31:
	;
	if base.Ui32(int32(11)) < base.Ui32(v124) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(v124) < base.Ui32(int32(12)) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v287 = v67
	v288 = v124
	v289 = v130
	v290 = v130
	v291 = v130
	goto L29
L35:
	;
	switch v186 - int32(1) {
	case 0:
		v236 = v187
		goto L42
	case 1:
		v231 = v187
		goto L43
	case 2:
		goto L44
	case 3:
		v224 = v188
		goto L45
	case 4:
		v221 = v188
		goto L46
	case 5:
		v216 = v188
		goto L47
	case 6:
		goto L48
	case 7:
		v207 = v189
		goto L49
	case 8:
		v202 = v189
		goto L50
	case 9:
		v197 = v189
		goto L51
	case 10:
		goto L52
	default:
		v357 = v187
		v358 = v188
		v359 = v189
		goto L28
	}
L36:
	;
	v185 = v67
	v186 = v124
	v187 = v130
	v188 = v130
	v189 = v130
	goto L35
L37:
	;
	goto L38
L38:
	;
	v137 = v67
	v138 = v124
	v139 = v130
	v140 = v130
	v141 = v130
	goto L39
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v144 = v143 + v140
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v148 = v147 + v141
	v150 = int32(4)
	v152 = v145 + v139 - v148 ^ base.I32_rotl(v148, v150)
	v156 = v144 - v152 ^ base.I32_rotl(v152, int32(6))
	v157 = v148 + v144
	v158 = v152 + v157
	v159 = v156 + v158
	v163 = v157 - v156 ^ base.I32_rotl(v156, int32(8))
	v167 = v158 - v163 ^ base.I32_rotl(v163, int32(16))
	v171 = v159 - v167 ^ base.I32_rotl(v167, int32(19))
	v172 = v163 + v159
	v173 = v167 + v172
	v174 = v171 + v173
	v178 = v172 - v171 ^ base.I32_rotl(v171, v150)
	v179 = int32(12)
	v180 = v137 + v179
	v182 = v138 - v179
	if base.Ui32(int32(11)) < base.Ui32(v182) {
		v137 = v180
		v138 = v182
		v139 = v173
		v140 = v174
		v141 = v178
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v185 = v180
	v186 = v182
	v187 = v173
	v188 = v174
	v189 = v178
	goto L35
L41:
	;
	goto L40
L42:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v357 = v236 + v237
	v358 = v188
	v359 = v189
	goto L28
L43:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	v236 = v232<<(uint(int32(8))%32) + v231
	goto L42
L44:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+2)))
	v231 = v227<<(uint(int32(16))%32) + v187
	goto L43
L45:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v357 = v225 + v187
	v358 = v224
	v359 = v189
	goto L28
L46:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)))
	v224 = v221 + v222
	goto L45
L47:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+5)))
	v221 = v217<<(uint(int32(8))%32) + v216
	goto L46
L48:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+6)))
	v216 = v212<<(uint(int32(16))%32) + v188
	goto L47
L49:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v357 = v208 + v187
	v358 = v210 + v188
	v359 = v207
	goto L28
L50:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+8)))
	v207 = v203<<(uint(int32(8))%32) + v202
	goto L49
L51:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+9)))
	v202 = v198<<(uint(int32(16))%32) + v197
	goto L50
L52:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+10)))
	v197 = v193<<(uint(int32(24))%32) + v189
	goto L51
L53:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v246 = v245 + v242
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v250 = v249 + v243
	v252 = int32(4)
	v254 = v247 + v241 - v250 ^ base.I32_rotl(v250, v252)
	v258 = v246 - v254 ^ base.I32_rotl(v254, int32(6))
	v259 = v250 + v246
	v260 = v254 + v259
	v261 = v258 + v260
	v265 = v259 - v258 ^ base.I32_rotl(v258, int32(8))
	v269 = v260 - v265 ^ base.I32_rotl(v265, int32(16))
	v273 = v261 - v269 ^ base.I32_rotl(v269, int32(19))
	v274 = v265 + v261
	v275 = v269 + v274
	v276 = v273 + v275
	v280 = v274 - v273 ^ base.I32_rotl(v273, v252)
	v281 = int32(12)
	v282 = v239 + v281
	v284 = v240 - v281
	if base.Ui32(int32(11)) < base.Ui32(v284) {
		v239 = v282
		v240 = v284
		v241 = v275
		v242 = v276
		v243 = v280
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v287 = v282
	v288 = v284
	v289 = v275
	v290 = v276
	v291 = v280
	goto L29
L55:
	;
	goto L54
L56:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v357 = v350 + v353
	v358 = v351
	v359 = v352
	goto L28
L57:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	v350 = v346<<(uint(int32(8))%32) + v343
	v351 = v344
	v352 = v345
	goto L56
L58:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+2)))
	v343 = v339<<(uint(int32(16))%32) + v336
	v344 = v337
	v345 = v338
	goto L57
L59:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+3)))
	v336 = v332<<(uint(int32(24))%32) + v289
	v337 = v330
	v338 = v331
	goto L58
L60:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+4)))
	v330 = v326 + v328
	v331 = v327
	goto L59
L61:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+5)))
	v326 = v322<<(uint(int32(8))%32) + v320
	v327 = v321
	goto L60
L62:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+6)))
	v320 = v316<<(uint(int32(16))%32) + v314
	v321 = v315
	goto L61
L63:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+7)))
	v314 = v310<<(uint(int32(24))%32) + v290
	v315 = v309
	goto L62
L64:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+8)))
	v309 = v305<<(uint(int32(8))%32) + v304
	goto L63
L65:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+9)))
	v304 = v300<<(uint(int32(16))%32) + v299
	goto L64
L66:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+10)))
	v299 = v295<<(uint(int32(24))%32) + v291
	goto L65
L67:
	;
	v397 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = int32(223358)
	*(*int64)(unsafe.Add(mBase, uint32(v62)+24)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v62 + int32(1056)
	v410 = F_pg_snprintf(m, v62+int32(32), int32(1024), int32(98479), v62+int32(16))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v62 + int32(32)
	v420 = F_pg_snprintf(m, v62+int32(1056), int32(1024), int32(166924), v62)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v423 = v62 + int32(1056)
	v424 = m.G0
	v426 = v424 - int32(16)
	m.G0 = v426
	v429 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_ResourceOwnerEnlarge(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	v434 = F_PathNameOpenFilePerm(m, v423, l2, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	m.G0 = v426 + int32(16)
	m.G0 = v62 + int32(2080)
	*(*int32)(unsafe.Add(mBase, uint32(v49+v34<<(uint(int32(2))%32)))) = v434
	if int32(0) < v434 {
		goto L82
	} else {
		goto L83
	}
L72:
	;
	if v434 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v439 == int32(44) {
		goto L71
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_ResourceOwnerRemember(m, v458, v434, int32(1586192))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L81
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v423
	F_errmsg(m, int32(282768), v426)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(476343), int32(1925), int32(370492))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _consts[553]))
	v466 = v463 + v434*int32(48)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+8)) = v468
	v471 = v466 + int32(4)
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v471))))
	v474 = v472 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v471))) = uint16(v474)
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[647])) = uint8(v477)
	goto L71
L82:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v493 == int32(0) {
		v29 = v48
		v34 = v40
		v35 = v49
		goto L3
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L4
L85:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v29 = v48
	v34 = v40
	v35 = v49
	goto L3
L87:
	;
	m.G0 = v17 + int32(1056)
	return v544
L88:
	;
	F_pfree(m, v49)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v522 = F_palloc(m, int32(8240))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L97
	}
L91:
	;
	if l3 != 0 {
		v544 = int32(0)
		goto L87
	} else {
		goto L92
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(32)
	F_errmsg(m, int32(284874), v17)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(475605), int32(339), int32(101004))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v524 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v522)+8)) = uint16(v524)
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v34
	v528 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v529 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v522)+32)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v522)+24)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v522)+20)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v522)+40)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v522)+12)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+10)) = uint8(base.B2i32(l2 == v524))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+4)) = v49
	v541 = F_pstrdup(m, l1)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+16)) = v541
	v544 = v522
	goto L87
}
