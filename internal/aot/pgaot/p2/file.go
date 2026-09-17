package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileSetCreate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
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
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
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
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
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
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int64
	_ = v595
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	v9 = m.G0
	v11 = v9 - int32(_a_F_FileSetCreate_0)
	m.G0 = v11
	v14 = v11 + int32(3120)
	v16 = l0 + int32(12)
	v17 = F_strlen(m, l1)
	mBase = m.M
	v23 = v17 - int32(1636608432)
	if l1&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v17) {
			v132 = l1
			v133 = v17
			v134 = v23
			v135 = v23
			v136 = v23
			for {
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
				v139 = v138 + v135
				v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
				v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
				v143 = v142 + v136
				v145 = int32(4)
				v147 = v140 + v134 - v143 ^ base.I32_rotl(v143, v145)
				v151 = v139 - v147 ^ base.I32_rotl(v147, int32(6))
				v152 = v143 + v139
				v153 = v147 + v152
				v154 = v151 + v153
				v158 = v152 - v151 ^ base.I32_rotl(v151, int32(8))
				v162 = v153 - v158 ^ base.I32_rotl(v158, int32(16))
				v166 = v154 - v162 ^ base.I32_rotl(v162, int32(19))
				v167 = v158 + v154
				v168 = v162 + v167
				v169 = v166 + v168
				v173 = v167 - v166 ^ base.I32_rotl(v166, v145)
				v174 = int32(12)
				v175 = v132 + v174
				v177 = v133 - v174
				if base.Ui32(int32(11)) < base.Ui32(v177) {
					v132 = v175
					v133 = v177
					v134 = v168
					v135 = v169
					v136 = v173
					continue
				} else {
					break
				}
				break
			}
			v180 = v175
			v181 = v177
			v182 = v168
			v183 = v169
			v184 = v173
		} else {
			v180 = l1
			v181 = v17
			v182 = v23
			v183 = v23
			v184 = v23
		}
		switch v181 - int32(1) {
		case 0:
			v243 = v182
			v244 = v183
			v245 = v184
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 1:
			v236 = v182
			v237 = v183
			v238 = v184
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 2:
			v229 = v182
			v230 = v183
			v231 = v184
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 3:
			v223 = v183
			v224 = v184
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 4:
			v219 = v183
			v220 = v184
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 5:
			v213 = v183
			v214 = v184
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+5)))
			v219 = v215<<(uint(int32(8))%32) + v213
			v220 = v214
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 6:
			v207 = v183
			v208 = v184
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+6)))
			v213 = v209<<(uint(int32(16))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+5)))
			v219 = v215<<(uint(int32(8))%32) + v213
			v220 = v214
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 7:
			v202 = v184
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+7)))
			v207 = v203<<(uint(int32(24))%32) + v183
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+6)))
			v213 = v209<<(uint(int32(16))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+5)))
			v219 = v215<<(uint(int32(8))%32) + v213
			v220 = v214
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 8:
			v197 = v184
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+8)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+7)))
			v207 = v203<<(uint(int32(24))%32) + v183
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+6)))
			v213 = v209<<(uint(int32(16))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+5)))
			v219 = v215<<(uint(int32(8))%32) + v213
			v220 = v214
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 9:
			v192 = v184
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+9)))
			v197 = v193<<(uint(int32(16))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+8)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+7)))
			v207 = v203<<(uint(int32(24))%32) + v183
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+6)))
			v213 = v209<<(uint(int32(16))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+5)))
			v219 = v215<<(uint(int32(8))%32) + v213
			v220 = v214
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		case 10:
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+10)))
			v192 = v188<<(uint(int32(24))%32) + v184
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+9)))
			v197 = v193<<(uint(int32(16))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+8)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+7)))
			v207 = v203<<(uint(int32(24))%32) + v183
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+6)))
			v213 = v209<<(uint(int32(16))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+5)))
			v219 = v215<<(uint(int32(8))%32) + v213
			v220 = v214
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+4)))
			v223 = v219 + v221
			v224 = v220
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
			v229 = v225<<(uint(int32(24))%32) + v182
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
			v236 = v232<<(uint(int32(16))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
			v243 = v239<<(uint(int32(8))%32) + v236
			v244 = v237
			v245 = v238
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
			v250 = v243 + v246
			v251 = v244
			v252 = v245
		default:
			v250 = v182
			v251 = v183
			v252 = v184
		}
	} else {
		if base.Ui32(v17) < base.Ui32(int32(12)) {
			v78 = l1
			v79 = v17
			v80 = v23
			v81 = v23
			v82 = v23
		} else {
			v30 = l1
			v31 = v17
			v32 = v23
			v33 = v23
			v34 = v23
			for {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				v37 = v36 + v33
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
				v41 = v40 + v34
				v43 = int32(4)
				v45 = v38 + v32 - v41 ^ base.I32_rotl(v41, v43)
				v49 = v37 - v45 ^ base.I32_rotl(v45, int32(6))
				v50 = v41 + v37
				v51 = v45 + v50
				v52 = v49 + v51
				v56 = v50 - v49 ^ base.I32_rotl(v49, int32(8))
				v60 = v51 - v56 ^ base.I32_rotl(v56, int32(16))
				v64 = v52 - v60 ^ base.I32_rotl(v60, int32(19))
				v65 = v56 + v52
				v66 = v60 + v65
				v67 = v64 + v66
				v71 = v65 - v64 ^ base.I32_rotl(v64, v43)
				v72 = int32(12)
				v73 = v30 + v72
				v75 = v31 - v72
				if base.Ui32(int32(11)) < base.Ui32(v75) {
					v30 = v73
					v31 = v75
					v32 = v66
					v33 = v67
					v34 = v71
					continue
				} else {
					break
				}
				break
			}
			v78 = v73
			v79 = v75
			v80 = v66
			v81 = v67
			v82 = v71
		}
		switch v79 - int32(1) {
		case 0:
			v129 = v80
			v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
			v250 = v129 + v130
			v251 = v81
			v252 = v82
		case 1:
			v124 = v80
			v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
			v129 = v125<<(uint(int32(8))%32) + v124
			v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
			v250 = v129 + v130
			v251 = v81
			v252 = v82
		case 2:
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+2)))
			v124 = v120<<(uint(int32(16))%32) + v80
			v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
			v129 = v125<<(uint(int32(8))%32) + v124
			v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
			v250 = v129 + v130
			v251 = v81
			v252 = v82
		case 3:
			v117 = v81
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v250 = v118 + v80
			v251 = v117
			v252 = v82
		case 4:
			v114 = v81
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)))
			v117 = v114 + v115
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v250 = v118 + v80
			v251 = v117
			v252 = v82
		case 5:
			v109 = v81
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)))
			v117 = v114 + v115
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v250 = v118 + v80
			v251 = v117
			v252 = v82
		case 6:
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
			v109 = v105<<(uint(int32(16))%32) + v81
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)))
			v117 = v114 + v115
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v250 = v118 + v80
			v251 = v117
			v252 = v82
		case 7:
			v100 = v82
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
			v250 = v101 + v80
			v251 = v103 + v81
			v252 = v100
		case 8:
			v95 = v82
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
			v250 = v101 + v80
			v251 = v103 + v81
			v252 = v100
		case 9:
			v90 = v82
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+9)))
			v95 = v91<<(uint(int32(16))%32) + v90
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
			v250 = v101 + v80
			v251 = v103 + v81
			v252 = v100
		case 10:
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+10)))
			v90 = v86<<(uint(int32(24))%32) + v82
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+9)))
			v95 = v91<<(uint(int32(16))%32) + v90
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
			v250 = v101 + v80
			v251 = v103 + v81
			v252 = v100
		default:
			v250 = v80
			v251 = v81
			v252 = v82
		}
	}
	v255 = int32(14)
	v257 = v251 ^ v252 - base.I32_rotl(v251, v255)
	v261 = v257 ^ v250 - base.I32_rotl(v257, int32(11))
	v265 = v261 ^ v251 - base.I32_rotl(v261, int32(25))
	v269 = v265 ^ v257 - base.I32_rotl(v265, int32(16))
	v273 = v269 ^ v261 - base.I32_rotl(v269, int32(4))
	v277 = v273 ^ v265 - base.I32_rotl(v273, v255)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = base.I32_rem_u_s(v277^v269-base.I32_rotl(v277, int32(24)), v282)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v16+v283<<(uint(int32(2))%32))))
	F_TempTablespacePath(m, v14, v287)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		return int32(0)
	} else {
		v292 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(_a_F_FileSetCreate_1)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v292
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
		v298 = v11 + int32(2096)
		v303 = F_pg_snprintf(m, v298, int32(1024), int32(_a_F_FileSetCreate_2), v11+int32(32))
		mBase = m.M
		v304 = m.ExcPending
		if v304 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v298
			v308 = v11 + int32(1072)
			v313 = F_pg_snprintf(m, v308, int32(1024), int32(_a_F_FileSetCreate_3), v11+int32(16))
			mBase = m.M
			v314 = m.ExcPending
			if v314 != 0 {
				return int32(0)
			} else {
				v316 = F_PathNameCreateTemporaryFile(m, v308, int32(0))
				mBase = m.M
				v317 = m.ExcPending
				if v317 != 0 {
					return int32(0)
				} else {
					if v316 <= int32(0) {
						v320 = F_strlen(m, l1)
						mBase = m.M
						v326 = v320 - int32(1636608432)
						if l1&int32(3) != 0 {
							if base.Ui32(int32(11)) < base.Ui32(v320) {
								v435 = l1
								v436 = v320
								v437 = v326
								v438 = v326
								v439 = v326
								for {
									v441 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
									v442 = v441 + v438
									v443 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
									v445 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
									v446 = v445 + v439
									v448 = int32(4)
									v450 = v443 + v437 - v446 ^ base.I32_rotl(v446, v448)
									v454 = v442 - v450 ^ base.I32_rotl(v450, int32(6))
									v455 = v446 + v442
									v456 = v450 + v455
									v457 = v454 + v456
									v461 = v455 - v454 ^ base.I32_rotl(v454, int32(8))
									v465 = v456 - v461 ^ base.I32_rotl(v461, int32(16))
									v469 = v457 - v465 ^ base.I32_rotl(v465, int32(19))
									v470 = v461 + v457
									v471 = v465 + v470
									v472 = v469 + v471
									v476 = v470 - v469 ^ base.I32_rotl(v469, v448)
									v477 = int32(12)
									v478 = v435 + v477
									v480 = v436 - v477
									if base.Ui32(int32(11)) < base.Ui32(v480) {
										v435 = v478
										v436 = v480
										v437 = v471
										v438 = v472
										v439 = v476
										continue
									} else {
										break
									}
									break
								}
								v483 = v478
								v484 = v480
								v485 = v471
								v486 = v472
								v487 = v476
							} else {
								v483 = l1
								v484 = v320
								v485 = v326
								v486 = v326
								v487 = v326
							}
							switch v484 - int32(1) {
							case 0:
								v546 = v485
								v547 = v486
								v548 = v487
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 1:
								v539 = v485
								v540 = v486
								v541 = v487
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 2:
								v532 = v485
								v533 = v486
								v534 = v487
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 3:
								v526 = v486
								v527 = v487
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 4:
								v522 = v486
								v523 = v487
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 5:
								v516 = v486
								v517 = v487
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+5)))
								v522 = v518<<(uint(int32(8))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 6:
								v510 = v486
								v511 = v487
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+6)))
								v516 = v512<<(uint(int32(16))%32) + v510
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+5)))
								v522 = v518<<(uint(int32(8))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 7:
								v505 = v487
								v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+7)))
								v510 = v506<<(uint(int32(24))%32) + v486
								v511 = v505
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+6)))
								v516 = v512<<(uint(int32(16))%32) + v510
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+5)))
								v522 = v518<<(uint(int32(8))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 8:
								v500 = v487
								v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+8)))
								v505 = v501<<(uint(int32(8))%32) + v500
								v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+7)))
								v510 = v506<<(uint(int32(24))%32) + v486
								v511 = v505
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+6)))
								v516 = v512<<(uint(int32(16))%32) + v510
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+5)))
								v522 = v518<<(uint(int32(8))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 9:
								v495 = v487
								v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+9)))
								v500 = v496<<(uint(int32(16))%32) + v495
								v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+8)))
								v505 = v501<<(uint(int32(8))%32) + v500
								v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+7)))
								v510 = v506<<(uint(int32(24))%32) + v486
								v511 = v505
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+6)))
								v516 = v512<<(uint(int32(16))%32) + v510
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+5)))
								v522 = v518<<(uint(int32(8))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							case 10:
								v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+10)))
								v495 = v491<<(uint(int32(24))%32) + v487
								v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+9)))
								v500 = v496<<(uint(int32(16))%32) + v495
								v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+8)))
								v505 = v501<<(uint(int32(8))%32) + v500
								v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+7)))
								v510 = v506<<(uint(int32(24))%32) + v486
								v511 = v505
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+6)))
								v516 = v512<<(uint(int32(16))%32) + v510
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+5)))
								v522 = v518<<(uint(int32(8))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
								v526 = v522 + v524
								v527 = v523
								v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+3)))
								v532 = v528<<(uint(int32(24))%32) + v485
								v533 = v526
								v534 = v527
								v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)))
								v539 = v535<<(uint(int32(16))%32) + v532
								v540 = v533
								v541 = v534
								v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
								v546 = v542<<(uint(int32(8))%32) + v539
								v547 = v540
								v548 = v541
								v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
								v553 = v546 + v549
								v554 = v547
								v555 = v548
							default:
								v553 = v485
								v554 = v486
								v555 = v487
							}
						} else {
							if base.Ui32(v320) < base.Ui32(int32(12)) {
								v381 = l1
								v382 = v320
								v383 = v326
								v384 = v326
								v385 = v326
							} else {
								v333 = l1
								v334 = v320
								v335 = v326
								v336 = v326
								v337 = v326
								for {
									v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
									v340 = v339 + v336
									v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
									v343 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
									v344 = v343 + v337
									v346 = int32(4)
									v348 = v341 + v335 - v344 ^ base.I32_rotl(v344, v346)
									v352 = v340 - v348 ^ base.I32_rotl(v348, int32(6))
									v353 = v344 + v340
									v354 = v348 + v353
									v355 = v352 + v354
									v359 = v353 - v352 ^ base.I32_rotl(v352, int32(8))
									v363 = v354 - v359 ^ base.I32_rotl(v359, int32(16))
									v367 = v355 - v363 ^ base.I32_rotl(v363, int32(19))
									v368 = v359 + v355
									v369 = v363 + v368
									v370 = v367 + v369
									v374 = v368 - v367 ^ base.I32_rotl(v367, v346)
									v375 = int32(12)
									v376 = v333 + v375
									v378 = v334 - v375
									if base.Ui32(int32(11)) < base.Ui32(v378) {
										v333 = v376
										v334 = v378
										v335 = v369
										v336 = v370
										v337 = v374
										continue
									} else {
										break
									}
									break
								}
								v381 = v376
								v382 = v378
								v383 = v369
								v384 = v370
								v385 = v374
							}
							switch v382 - int32(1) {
							case 0:
								v432 = v383
								v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
								v553 = v432 + v433
								v554 = v384
								v555 = v385
							case 1:
								v427 = v383
								v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)))
								v432 = v428<<(uint(int32(8))%32) + v427
								v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
								v553 = v432 + v433
								v554 = v384
								v555 = v385
							case 2:
								v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+2)))
								v427 = v423<<(uint(int32(16))%32) + v383
								v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)))
								v432 = v428<<(uint(int32(8))%32) + v427
								v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
								v553 = v432 + v433
								v554 = v384
								v555 = v385
							case 3:
								v420 = v384
								v421 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v553 = v421 + v383
								v554 = v420
								v555 = v385
							case 4:
								v417 = v384
								v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+4)))
								v420 = v417 + v418
								v421 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v553 = v421 + v383
								v554 = v420
								v555 = v385
							case 5:
								v412 = v384
								v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+5)))
								v417 = v413<<(uint(int32(8))%32) + v412
								v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+4)))
								v420 = v417 + v418
								v421 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v553 = v421 + v383
								v554 = v420
								v555 = v385
							case 6:
								v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+6)))
								v412 = v408<<(uint(int32(16))%32) + v384
								v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+5)))
								v417 = v413<<(uint(int32(8))%32) + v412
								v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+4)))
								v420 = v417 + v418
								v421 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v553 = v421 + v383
								v554 = v420
								v555 = v385
							case 7:
								v403 = v385
								v404 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v406 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
								v553 = v404 + v383
								v554 = v406 + v384
								v555 = v403
							case 8:
								v398 = v385
								v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+8)))
								v403 = v399<<(uint(int32(8))%32) + v398
								v404 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v406 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
								v553 = v404 + v383
								v554 = v406 + v384
								v555 = v403
							case 9:
								v393 = v385
								v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+9)))
								v398 = v394<<(uint(int32(16))%32) + v393
								v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+8)))
								v403 = v399<<(uint(int32(8))%32) + v398
								v404 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v406 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
								v553 = v404 + v383
								v554 = v406 + v384
								v555 = v403
							case 10:
								v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+10)))
								v393 = v389<<(uint(int32(24))%32) + v385
								v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+9)))
								v398 = v394<<(uint(int32(16))%32) + v393
								v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+8)))
								v403 = v399<<(uint(int32(8))%32) + v398
								v404 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
								v406 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
								v553 = v404 + v383
								v554 = v406 + v384
								v555 = v403
							default:
								v553 = v383
								v554 = v384
								v555 = v385
							}
						}
						v558 = int32(14)
						v560 = v554 ^ v555 - base.I32_rotl(v554, v558)
						v564 = v560 ^ v553 - base.I32_rotl(v560, int32(11))
						v568 = v564 ^ v554 - base.I32_rotl(v564, int32(25))
						v572 = v568 ^ v560 - base.I32_rotl(v568, int32(16))
						v576 = v572 ^ v564 - base.I32_rotl(v572, int32(4))
						v580 = v576 ^ v568 - base.I32_rotl(v576, v558)
						v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v586 = base.I32_rem_u_s(v580^v572-base.I32_rotl(v580, int32(24)), v585)
						v590 = *(*int32)(unsafe.Add(mBase, uint32(v16+v586<<(uint(int32(2))%32))))
						F_TempTablespacePath(m, v298, v590)
						mBase = m.M
						v592 = m.ExcPending
						if v592 != 0 {
							return int32(0)
						} else {
							F_TempTablespacePath(m, v14, v590)
							mBase = m.M
							v594 = m.ExcPending
							if v594 != 0 {
								return int32(0)
							} else {
								v595 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_FileSetCreate_1)
								*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v595
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
								v601 = v11 + int32(48)
								v604 = F_pg_snprintf(m, v601, int32(1024), int32(_a_F_FileSetCreate_2), v11)
								mBase = m.M
								v605 = m.ExcPending
								if v605 != 0 {
									return int32(0)
								} else {
									v606 = m.G0
									v608 = v606 - int32(32)
									m.G0 = v608
									v611 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[0]))
									v612 = F_mkdir(m, v601, v611)
									mBase = m.M
									if int32(0) <= v612 {
										m.G0 = v608 + int32(32)
										v675 = F_PathNameCreateTemporaryFile(m, v11+int32(1072), int32(1))
										mBase = m.M
										v676 = m.ExcPending
										if v676 != 0 {
											return int32(0)
										} else {
											v679 = v675
											m.G0 = v11 + int32(_a_F_FileSetCreate_0)
											return v679
										}
									} else {
										v616 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[1]))
										if v616 == int32(20) {
											m.G0 = v608 + int32(32)
											v675 = F_PathNameCreateTemporaryFile(m, v11+int32(1072), int32(1))
											mBase = m.M
											v676 = m.ExcPending
											if v676 != 0 {
												return int32(0)
											} else {
												v679 = v675
												m.G0 = v11 + int32(_a_F_FileSetCreate_0)
												return v679
											}
										} else {
											v620 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[0]))
											v621 = F_mkdir(m, v298, v620)
											mBase = m.M
											if v621 < int32(0) {
												v625 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[1]))
												if v625 != int32(20) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v643 = m.ExcPending
													if v643 != 0 {
														return int32(0)
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v645 = m.ExcPending
														if v645 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v608)+16)) = v298
															F_errmsg(m, int32(_a_F_FileSetCreate_4), v608+int32(16))
															mBase = m.M
															v651 = m.ExcPending
															if v651 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_FileSetCreate_5), int32(1685), int32(_a_F_FileSetCreate_6))
																mBase = m.M
																v656 = m.ExcPending
																if v656 != 0 {
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
													v629 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[0]))
													v630 = F_mkdir(m, v601, v629)
													mBase = m.M
													if int32(0) <= v630 {
														m.G0 = v608 + int32(32)
														v675 = F_PathNameCreateTemporaryFile(m, v11+int32(1072), int32(1))
														mBase = m.M
														v676 = m.ExcPending
														if v676 != 0 {
															return int32(0)
														} else {
															v679 = v675
															m.G0 = v11 + int32(_a_F_FileSetCreate_0)
															return v679
														}
													} else {
														v634 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[1]))
														if v634 != int32(20) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v660 = m.ExcPending
															if v660 != 0 {
																return int32(0)
															} else {
																F_errcode_for_file_access(m)
																mBase = m.M
																v662 = m.ExcPending
																if v662 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v608))) = v601
																	F_errmsg(m, int32(_a_F_FileSetCreate_7), v608)
																	mBase = m.M
																	v666 = m.ExcPending
																	if v666 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_FileSetCreate_5), int32(1692), int32(_a_F_FileSetCreate_6))
																		mBase = m.M
																		v671 = m.ExcPending
																		if v671 != 0 {
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
															m.G0 = v608 + int32(32)
															v675 = F_PathNameCreateTemporaryFile(m, v11+int32(1072), int32(1))
															mBase = m.M
															v676 = m.ExcPending
															if v676 != 0 {
																return int32(0)
															} else {
																v679 = v675
																m.G0 = v11 + int32(_a_F_FileSetCreate_0)
																return v679
															}
														}
													}
												}
											} else {
												v629 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[0]))
												v630 = F_mkdir(m, v601, v629)
												mBase = m.M
												if int32(0) <= v630 {
													m.G0 = v608 + int32(32)
													v675 = F_PathNameCreateTemporaryFile(m, v11+int32(1072), int32(1))
													mBase = m.M
													v676 = m.ExcPending
													if v676 != 0 {
														return int32(0)
													} else {
														v679 = v675
														m.G0 = v11 + int32(_a_F_FileSetCreate_0)
														return v679
													}
												} else {
													v634 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetCreate[1]))
													if v634 != int32(20) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v660 = m.ExcPending
														if v660 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v662 = m.ExcPending
															if v662 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v608))) = v601
																F_errmsg(m, int32(_a_F_FileSetCreate_7), v608)
																mBase = m.M
																v666 = m.ExcPending
																if v666 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_FileSetCreate_5), int32(1692), int32(_a_F_FileSetCreate_6))
																	mBase = m.M
																	v671 = m.ExcPending
																	if v671 != 0 {
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
														m.G0 = v608 + int32(32)
														v675 = F_PathNameCreateTemporaryFile(m, v11+int32(1072), int32(1))
														mBase = m.M
														v676 = m.ExcPending
														if v676 != 0 {
															return int32(0)
														} else {
															v679 = v675
															m.G0 = v11 + int32(_a_F_FileSetCreate_0)
															return v679
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
						v679 = v316
						m.G0 = v11 + int32(_a_F_FileSetCreate_0)
						return v679
					}
				}
			}
		}
	}
}
func F_FileSetInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v225 int32
	_ = v225
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
	v16 = int32(_a_F_FileSetInit_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
	v23 = base.I32_rem_u_s(v17+int32(1), int32(2147483647))
	*(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[1])) = v23
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v28 = l0 + int32(12)
		v29 = int32(8)
		v31 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[2]))
		if v29 < v31 {
			v34 = v29
		} else {
			v34 = v31
		}
		if int32(0) < v34 {
			v38 = v34 & int32(3)
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[3]))
			if base.Ui32(int32(4)) <= base.Ui32(v34) {
				v49 = v2
				v54 = v2
				for {
					v58 = v49 << (uint(int32(2)) % 32)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v40+v58)))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v58))) = v61
					v63 = int32(4)
					v64 = v58 | v63
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v40+v64)))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v64))) = v67
					v70 = v58 | int32(8)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v40+v70)))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v70))) = v73
					v76 = v58 | int32(12)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v40+v76)))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v76))) = v79
					v82 = v49 + v63
					v84 = v54 + v63
					if v84 != v34&int32(2147483644) {
						v49 = v82
						v54 = v84
						continue
					} else {
						break
					}
					break
				}
				if v38 == int32(0) {
				} else {
					v92 = v82
					v104 = v92
					v110 = v2
					for {
						v113 = v104 << (uint(int32(2)) % 32)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v40+v113)))
						*(*int32)(unsafe.Add(mBase, uint32(v28+v113))) = v116
						v118 = int32(1)
						v121 = v110 + v118
						if v121 != v38 {
							v104 = v104 + v118
							v110 = v121
							continue
						} else {
							break
						}
						break
					}
				}
			} else {
				v92 = v2
				v104 = v92
				v110 = v2
				for {
					v113 = v104 << (uint(int32(2)) % 32)
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v40+v113)))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v113))) = v116
					v118 = int32(1)
					v121 = v110 + v118
					if v121 != v38 {
						v104 = v104 + v118
						v110 = v121
						continue
					} else {
						break
					}
					break
				}
			}
			v148 = v34
		} else {
			v148 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v148
		if v148 != 0 {
			if v148 <= int32(0) {
			} else {
				v152 = int32(0)
				if v148 != int32(1) {
					v159 = v152
					v170 = v2
					for {
						v173 = v28 + v159<<(uint(int32(2))%32)
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
						if v174 == int32(0) {
							v178 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v173))) = v178
						} else {
						}
						v180 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v180 == int32(0) {
							v184 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v184
						} else {
						}
						v186 = int32(2)
						v187 = v159 + v186
						v189 = v170 + v186
						if v189 != v148&int32(2147483646) {
							v159 = v187
							v170 = v189
							continue
						} else {
							break
						}
						break
					}
					if v148&int32(1) == int32(0) {
					} else {
						v193 = v187
						v207 = v28 + v193<<(uint(int32(2))%32)
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
						if v208 != 0 {
						} else {
							v210 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v207))) = v210
						}
					}
				} else {
					v193 = v152
					v207 = v28 + v193<<(uint(int32(2))%32)
					v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
					if v208 != 0 {
					} else {
						v210 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v207))) = v210
					}
				}
			}
			return
		} else {
			v225 = *(*int32)(unsafe.Add(mBase, _c_F_FileSetInit[4]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v225
			return
		}
	}
}
