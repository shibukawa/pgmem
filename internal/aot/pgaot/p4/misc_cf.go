package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cfunc_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
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
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
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
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
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
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
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
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
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
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v615 int32
	_ = v615
	v4 = int32(24)
	v10 = int32(-1636608408)
	if l0&int32(3) != 0 {
		v119 = l0
		v120 = v4
		v121 = v10
		v122 = v10
		v123 = v10
		for {
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
			v126 = v125 + v122
			v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
			v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
			v130 = v129 + v123
			v132 = int32(4)
			v134 = v127 + v121 - v130 ^ base.I32_rotl(v130, v132)
			v138 = v126 - v134 ^ base.I32_rotl(v134, int32(6))
			v139 = v130 + v126
			v140 = v134 + v139
			v141 = v138 + v140
			v145 = v139 - v138 ^ base.I32_rotl(v138, int32(8))
			v149 = v140 - v145 ^ base.I32_rotl(v145, int32(16))
			v153 = v141 - v149 ^ base.I32_rotl(v149, int32(19))
			v154 = v145 + v141
			v155 = v149 + v154
			v156 = v153 + v155
			v160 = v154 - v153 ^ base.I32_rotl(v153, v132)
			v161 = int32(12)
			v162 = v119 + v161
			v164 = v120 - v161
			if base.Ui32(int32(11)) < base.Ui32(v164) {
				v119 = v162
				v120 = v164
				v121 = v155
				v122 = v156
				v123 = v160
				continue
			} else {
				break
			}
			break
		}
		switch v164 - int32(1) {
		case 0:
			v230 = v155
			v231 = v156
			v232 = v160
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 1:
			v223 = v155
			v224 = v156
			v225 = v160
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 2:
			v216 = v155
			v217 = v156
			v218 = v160
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 3:
			v210 = v156
			v211 = v160
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 4:
			v206 = v156
			v207 = v160
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 5:
			v200 = v156
			v201 = v160
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 6:
			v194 = v156
			v195 = v160
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 7:
			v189 = v160
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 8:
			v184 = v160
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 9:
			v179 = v160
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+10)))
			v179 = v175<<(uint(int32(24))%32) + v160
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		default:
			v237 = v155
			v238 = v156
			v239 = v160
		}
	} else {
		v17 = l0
		v18 = v4
		v19 = v10
		v20 = v10
		v21 = v10
		for {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v24 = v23 + v20
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v28 = v27 + v21
			v30 = int32(4)
			v32 = v25 + v19 - v28 ^ base.I32_rotl(v28, v30)
			v36 = v24 - v32 ^ base.I32_rotl(v32, int32(6))
			v37 = v28 + v24
			v38 = v32 + v37
			v39 = v36 + v38
			v43 = v37 - v36 ^ base.I32_rotl(v36, int32(8))
			v47 = v38 - v43 ^ base.I32_rotl(v43, int32(16))
			v51 = v39 - v47 ^ base.I32_rotl(v47, int32(19))
			v52 = v43 + v39
			v53 = v47 + v52
			v54 = v51 + v53
			v58 = v52 - v51 ^ base.I32_rotl(v51, v30)
			v59 = int32(12)
			v60 = v17 + v59
			v62 = v18 - v59
			if base.Ui32(int32(11)) < base.Ui32(v62) {
				v17 = v60
				v18 = v62
				v19 = v53
				v20 = v54
				v21 = v58
				continue
			} else {
				break
			}
			break
		}
		switch v62 - int32(1) {
		case 0:
			v116 = v53
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
			v237 = v116 + v117
			v238 = v54
			v239 = v58
		case 1:
			v111 = v53
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
			v237 = v116 + v117
			v238 = v54
			v239 = v58
		case 2:
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
			v111 = v107<<(uint(int32(16))%32) + v53
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
			v237 = v116 + v117
			v238 = v54
			v239 = v58
		case 3:
			v104 = v54
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 4:
			v101 = v54
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 5:
			v96 = v54
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 6:
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+6)))
			v96 = v92<<(uint(int32(16))%32) + v54
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 7:
			v87 = v58
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		case 8:
			v82 = v58
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		case 9:
			v77 = v58
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		case 10:
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)))
			v77 = v73<<(uint(int32(24))%32) + v58
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		default:
			v237 = v53
			v238 = v54
			v239 = v58
		}
	}
	v242 = int32(14)
	v244 = v238 ^ v239 - base.I32_rotl(v238, v242)
	v248 = v244 ^ v237 - base.I32_rotl(v244, int32(11))
	v252 = v248 ^ v238 - base.I32_rotl(v248, int32(25))
	v256 = v252 ^ v244 - base.I32_rotl(v252, int32(16))
	v260 = v256 ^ v248 - base.I32_rotl(v256, int32(4))
	v264 = v260 ^ v252 - base.I32_rotl(v260, v242)
	v268 = v264 ^ v256 - base.I32_rotl(v264, int32(24))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v269 {
		v273 = l0 + int32(28)
		v275 = v269 << (uint(int32(2)) % 32)
		v281 = v275 - int32(1636608432)
		if v273&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v275) {
				v390 = v273
				v391 = v275
				v392 = v281
				v393 = v281
				v394 = v281
				for {
					v396 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
					v397 = v396 + v393
					v398 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
					v400 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
					v401 = v400 + v394
					v403 = int32(4)
					v405 = v398 + v392 - v401 ^ base.I32_rotl(v401, v403)
					v409 = v397 - v405 ^ base.I32_rotl(v405, int32(6))
					v410 = v401 + v397
					v411 = v405 + v410
					v412 = v409 + v411
					v416 = v410 - v409 ^ base.I32_rotl(v409, int32(8))
					v420 = v411 - v416 ^ base.I32_rotl(v416, int32(16))
					v424 = v412 - v420 ^ base.I32_rotl(v420, int32(19))
					v425 = v416 + v412
					v426 = v420 + v425
					v427 = v424 + v426
					v431 = v425 - v424 ^ base.I32_rotl(v424, v403)
					v432 = int32(12)
					v433 = v390 + v432
					v435 = v391 - v432
					if base.Ui32(int32(11)) < base.Ui32(v435) {
						v390 = v433
						v391 = v435
						v392 = v426
						v393 = v427
						v394 = v431
						continue
					} else {
						break
					}
					break
				}
				v438 = v433
				v439 = v435
				v440 = v426
				v441 = v427
				v442 = v431
			} else {
				v438 = v273
				v439 = v275
				v440 = v281
				v441 = v281
				v442 = v281
			}
			switch v439 - int32(1) {
			case 0:
				v501 = v440
				v502 = v441
				v503 = v442
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 1:
				v494 = v440
				v495 = v441
				v496 = v442
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 2:
				v487 = v440
				v488 = v441
				v489 = v442
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 3:
				v481 = v441
				v482 = v442
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 4:
				v477 = v441
				v478 = v442
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 5:
				v471 = v441
				v472 = v442
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 6:
				v465 = v441
				v466 = v442
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 7:
				v460 = v442
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 8:
				v455 = v442
				v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
				v460 = v456<<(uint(int32(8))%32) + v455
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 9:
				v450 = v442
				v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+9)))
				v455 = v451<<(uint(int32(16))%32) + v450
				v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
				v460 = v456<<(uint(int32(8))%32) + v455
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 10:
				v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+10)))
				v450 = v446<<(uint(int32(24))%32) + v442
				v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+9)))
				v455 = v451<<(uint(int32(16))%32) + v450
				v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
				v460 = v456<<(uint(int32(8))%32) + v455
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			default:
				v508 = v440
				v509 = v441
				v510 = v442
			}
		} else {
			if base.Ui32(v275) < base.Ui32(int32(12)) {
				v336 = v273
				v337 = v275
				v338 = v281
				v339 = v281
				v340 = v281
			} else {
				v288 = v273
				v289 = v275
				v290 = v281
				v291 = v281
				v292 = v281
				for {
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
					v295 = v294 + v291
					v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
					v298 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
					v299 = v298 + v292
					v301 = int32(4)
					v303 = v296 + v290 - v299 ^ base.I32_rotl(v299, v301)
					v307 = v295 - v303 ^ base.I32_rotl(v303, int32(6))
					v308 = v299 + v295
					v309 = v303 + v308
					v310 = v307 + v309
					v314 = v308 - v307 ^ base.I32_rotl(v307, int32(8))
					v318 = v309 - v314 ^ base.I32_rotl(v314, int32(16))
					v322 = v310 - v318 ^ base.I32_rotl(v318, int32(19))
					v323 = v314 + v310
					v324 = v318 + v323
					v325 = v322 + v324
					v329 = v323 - v322 ^ base.I32_rotl(v322, v301)
					v330 = int32(12)
					v331 = v288 + v330
					v333 = v289 - v330
					if base.Ui32(int32(11)) < base.Ui32(v333) {
						v288 = v331
						v289 = v333
						v290 = v324
						v291 = v325
						v292 = v329
						continue
					} else {
						break
					}
					break
				}
				v336 = v331
				v337 = v333
				v338 = v324
				v339 = v325
				v340 = v329
			}
			switch v337 - int32(1) {
			case 0:
				v387 = v338
				v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
				v508 = v387 + v388
				v509 = v339
				v510 = v340
			case 1:
				v382 = v338
				v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
				v387 = v383<<(uint(int32(8))%32) + v382
				v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
				v508 = v387 + v388
				v509 = v339
				v510 = v340
			case 2:
				v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+2)))
				v382 = v378<<(uint(int32(16))%32) + v338
				v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
				v387 = v383<<(uint(int32(8))%32) + v382
				v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
				v508 = v387 + v388
				v509 = v339
				v510 = v340
			case 3:
				v375 = v339
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 4:
				v372 = v339
				v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
				v375 = v372 + v373
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 5:
				v367 = v339
				v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+5)))
				v372 = v368<<(uint(int32(8))%32) + v367
				v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
				v375 = v372 + v373
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 6:
				v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+6)))
				v367 = v363<<(uint(int32(16))%32) + v339
				v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+5)))
				v372 = v368<<(uint(int32(8))%32) + v367
				v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
				v375 = v372 + v373
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 7:
				v358 = v340
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			case 8:
				v353 = v340
				v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
				v358 = v354<<(uint(int32(8))%32) + v353
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			case 9:
				v348 = v340
				v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
				v353 = v349<<(uint(int32(16))%32) + v348
				v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
				v358 = v354<<(uint(int32(8))%32) + v353
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			case 10:
				v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+10)))
				v348 = v344<<(uint(int32(24))%32) + v340
				v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
				v353 = v349<<(uint(int32(16))%32) + v348
				v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
				v358 = v354<<(uint(int32(8))%32) + v353
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			default:
				v508 = v338
				v509 = v339
				v510 = v340
			}
		}
		v513 = int32(14)
		v515 = v509 ^ v510 - base.I32_rotl(v509, v513)
		v519 = v515 ^ v508 - base.I32_rotl(v515, int32(11))
		v523 = v519 ^ v509 - base.I32_rotl(v519, int32(25))
		v527 = v523 ^ v515 - base.I32_rotl(v523, int32(16))
		v531 = v527 ^ v519 - base.I32_rotl(v527, int32(4))
		v535 = v531 ^ v523 - base.I32_rotl(v531, v513)
		v549 = v535 ^ v527 - base.I32_rotl(v535, int32(24)) + (v268<<(uint(int32(6))%32) + int32(base.Ui32(v268)>>(uint(int32(2))%32))) - int32(1640531527) ^ v268
	} else {
		v549 = v268
	}
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v550 != 0 {
		v551 = int32(0)
		v555 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
		v556 = F_hash_bytes_uint32(m, v555)
		mBase = m.M
		v557 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
		v558 = F_hash_bytes_uint32(m, v557)
		mBase = m.M
		v559 = int32(1640531527)
		v560 = v556 - v559
		v569 = v558 + v560<<(uint(int32(6))%32) + int32(base.Ui32(v560)>>(uint(int32(2))%32)) - v559 ^ v560
		v570 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
		if v551 < v570 {
			v576 = v569
			v577 = v551
			v578 = v570
			for {
				v586 = *(*int32)(unsafe.Add(mBase, uint32(v550+int32(88)+v578<<(uint(int32(4))%32)+v577*int32(100))))
				v587 = F_hash_bytes_uint32(m, v586)
				mBase = m.M
				v596 = v587 + (v576<<(uint(int32(6))%32) + int32(base.Ui32(v576)>>(uint(int32(2))%32))) - int32(1640531527) ^ v576
				v598 = v577 + int32(1)
				v599 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
				if v598 < v599 {
					v576 = v596
					v577 = v598
					v578 = v599
					continue
				} else {
					break
				}
				break
			}
			v602 = v596
		} else {
			v602 = v569
		}
		v615 = v602 + (v549<<(uint(int32(6))%32) + int32(base.Ui32(v549)>>(uint(int32(2))%32))) - int32(1640531527) ^ v549
	} else {
		v615 = v549
	}
	return v615
}
