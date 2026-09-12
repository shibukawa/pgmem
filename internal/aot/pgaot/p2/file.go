package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileSetCreate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
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
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
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
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
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
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
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
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int64
	_ = v603
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	v7 = m.G0
	v9 = v7 - int32(4144)
	m.G0 = v9
	v14 = l0 + int32(12)
	v15 = F_strlen(m, l1)
	mBase = m.M
	v21 = v15 - int32(1636608432)
	if l1&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v15) {
			v130 = l1
			v131 = v15
			v132 = v21
			v133 = v21
			v134 = v21
			for {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
				v137 = v136 + v133
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
				v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
				v141 = v140 + v134
				v143 = int32(4)
				v145 = v138 + v132 - v141 ^ base.I32_rotl(v141, v143)
				v149 = v137 - v145 ^ base.I32_rotl(v145, int32(6))
				v150 = v141 + v137
				v151 = v145 + v150
				v152 = v149 + v151
				v156 = v150 - v149 ^ base.I32_rotl(v149, int32(8))
				v160 = v151 - v156 ^ base.I32_rotl(v156, int32(16))
				v164 = v152 - v160 ^ base.I32_rotl(v160, int32(19))
				v165 = v156 + v152
				v166 = v160 + v165
				v167 = v164 + v166
				v171 = v165 - v164 ^ base.I32_rotl(v164, v143)
				v172 = int32(12)
				v173 = v130 + v172
				v175 = v131 - v172
				if base.Ui32(int32(11)) < base.Ui32(v175) {
					v130 = v173
					v131 = v175
					v132 = v166
					v133 = v167
					v134 = v171
					continue
				} else {
					break
				}
				break
			}
			v178 = v173
			v179 = v175
			v180 = v166
			v181 = v167
			v182 = v171
		} else {
			v178 = l1
			v179 = v15
			v180 = v21
			v181 = v21
			v182 = v21
		}
		switch v179 - int32(1) {
		case 0:
			v241 = v180
			v242 = v181
			v243 = v182
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 1:
			v234 = v180
			v235 = v181
			v236 = v182
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 2:
			v227 = v180
			v228 = v181
			v229 = v182
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 3:
			v221 = v181
			v222 = v182
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 4:
			v217 = v181
			v218 = v182
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 5:
			v211 = v181
			v212 = v182
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
			v217 = v213<<(uint(int32(8))%32) + v211
			v218 = v212
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 6:
			v205 = v181
			v206 = v182
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
			v211 = v207<<(uint(int32(16))%32) + v205
			v212 = v206
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
			v217 = v213<<(uint(int32(8))%32) + v211
			v218 = v212
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 7:
			v200 = v182
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
			v205 = v201<<(uint(int32(24))%32) + v181
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
			v211 = v207<<(uint(int32(16))%32) + v205
			v212 = v206
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
			v217 = v213<<(uint(int32(8))%32) + v211
			v218 = v212
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 8:
			v195 = v182
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
			v200 = v196<<(uint(int32(8))%32) + v195
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
			v205 = v201<<(uint(int32(24))%32) + v181
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
			v211 = v207<<(uint(int32(16))%32) + v205
			v212 = v206
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
			v217 = v213<<(uint(int32(8))%32) + v211
			v218 = v212
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 9:
			v190 = v182
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
			v195 = v191<<(uint(int32(16))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
			v200 = v196<<(uint(int32(8))%32) + v195
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
			v205 = v201<<(uint(int32(24))%32) + v181
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
			v211 = v207<<(uint(int32(16))%32) + v205
			v212 = v206
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
			v217 = v213<<(uint(int32(8))%32) + v211
			v218 = v212
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		case 10:
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
			v190 = v186<<(uint(int32(24))%32) + v182
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
			v195 = v191<<(uint(int32(16))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
			v200 = v196<<(uint(int32(8))%32) + v195
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
			v205 = v201<<(uint(int32(24))%32) + v181
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
			v211 = v207<<(uint(int32(16))%32) + v205
			v212 = v206
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
			v217 = v213<<(uint(int32(8))%32) + v211
			v218 = v212
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
			v221 = v217 + v219
			v222 = v218
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
			v227 = v223<<(uint(int32(24))%32) + v180
			v228 = v221
			v229 = v222
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
			v234 = v230<<(uint(int32(16))%32) + v227
			v235 = v228
			v236 = v229
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
			v241 = v237<<(uint(int32(8))%32) + v234
			v242 = v235
			v243 = v236
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
			v248 = v241 + v244
			v249 = v242
			v250 = v243
		default:
			v248 = v180
			v249 = v181
			v250 = v182
		}
	} else {
		if base.Ui32(v15) < base.Ui32(int32(12)) {
			v76 = l1
			v77 = v15
			v78 = v21
			v79 = v21
			v80 = v21
		} else {
			v28 = l1
			v29 = v15
			v30 = v21
			v31 = v21
			v32 = v21
			for {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v35 = v34 + v31
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
				v39 = v38 + v32
				v41 = int32(4)
				v43 = v36 + v30 - v39 ^ base.I32_rotl(v39, v41)
				v47 = v35 - v43 ^ base.I32_rotl(v43, int32(6))
				v48 = v39 + v35
				v49 = v43 + v48
				v50 = v47 + v49
				v54 = v48 - v47 ^ base.I32_rotl(v47, int32(8))
				v58 = v49 - v54 ^ base.I32_rotl(v54, int32(16))
				v62 = v50 - v58 ^ base.I32_rotl(v58, int32(19))
				v63 = v54 + v50
				v64 = v58 + v63
				v65 = v62 + v64
				v69 = v63 - v62 ^ base.I32_rotl(v62, v41)
				v70 = int32(12)
				v71 = v28 + v70
				v73 = v29 - v70
				if base.Ui32(int32(11)) < base.Ui32(v73) {
					v28 = v71
					v29 = v73
					v30 = v64
					v31 = v65
					v32 = v69
					continue
				} else {
					break
				}
				break
			}
			v76 = v71
			v77 = v73
			v78 = v64
			v79 = v65
			v80 = v69
		}
		switch v77 - int32(1) {
		case 0:
			v127 = v78
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
			v248 = v127 + v128
			v249 = v79
			v250 = v80
		case 1:
			v122 = v78
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
			v127 = v123<<(uint(int32(8))%32) + v122
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
			v248 = v127 + v128
			v249 = v79
			v250 = v80
		case 2:
			v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
			v122 = v118<<(uint(int32(16))%32) + v78
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
			v127 = v123<<(uint(int32(8))%32) + v122
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
			v248 = v127 + v128
			v249 = v79
			v250 = v80
		case 3:
			v115 = v79
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v248 = v116 + v78
			v249 = v115
			v250 = v80
		case 4:
			v112 = v79
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
			v115 = v112 + v113
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v248 = v116 + v78
			v249 = v115
			v250 = v80
		case 5:
			v107 = v79
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
			v112 = v108<<(uint(int32(8))%32) + v107
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
			v115 = v112 + v113
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v248 = v116 + v78
			v249 = v115
			v250 = v80
		case 6:
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
			v107 = v103<<(uint(int32(16))%32) + v79
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
			v112 = v108<<(uint(int32(8))%32) + v107
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
			v115 = v112 + v113
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v248 = v116 + v78
			v249 = v115
			v250 = v80
		case 7:
			v98 = v80
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
			v248 = v99 + v78
			v249 = v101 + v79
			v250 = v98
		case 8:
			v93 = v80
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
			v98 = v94<<(uint(int32(8))%32) + v93
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
			v248 = v99 + v78
			v249 = v101 + v79
			v250 = v98
		case 9:
			v88 = v80
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
			v93 = v89<<(uint(int32(16))%32) + v88
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
			v98 = v94<<(uint(int32(8))%32) + v93
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
			v248 = v99 + v78
			v249 = v101 + v79
			v250 = v98
		case 10:
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)))
			v88 = v84<<(uint(int32(24))%32) + v80
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
			v93 = v89<<(uint(int32(16))%32) + v88
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
			v98 = v94<<(uint(int32(8))%32) + v93
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
			v248 = v99 + v78
			v249 = v101 + v79
			v250 = v98
		default:
			v248 = v78
			v249 = v79
			v250 = v80
		}
	}
	v253 = int32(14)
	v255 = v249 ^ v250 - base.I32_rotl(v249, v253)
	v259 = v255 ^ v248 - base.I32_rotl(v255, int32(11))
	v263 = v259 ^ v249 - base.I32_rotl(v259, int32(25))
	v267 = v263 ^ v255 - base.I32_rotl(v263, int32(16))
	v271 = v267 ^ v259 - base.I32_rotl(v267, int32(4))
	v275 = v271 ^ v263 - base.I32_rotl(v271, v253)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v281 = base.I32_rem_u_s(v275^v267-base.I32_rotl(v275, int32(24)), v280)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v14+v281<<(uint(int32(2))%32))))
	F_TempTablespacePath(m, v9+int32(3120), v285)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		return int32(0)
	} else {
		v290 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(235464)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v290
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(3120)
		v303 = F_pg_snprintf(m, v9+int32(2096), int32(1024), int32(106036), v9+int32(32))
		mBase = m.M
		v304 = m.ExcPending
		if v304 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(2096)
			v315 = F_pg_snprintf(m, v9+int32(1072), int32(1024), int32(177111), v9+int32(16))
			mBase = m.M
			v316 = m.ExcPending
			if v316 != 0 {
				return int32(0)
			} else {
				v320 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(0))
				mBase = m.M
				v321 = m.ExcPending
				if v321 != 0 {
					return int32(0)
				} else {
					if v320 <= int32(0) {
						v326 = F_strlen(m, l1)
						mBase = m.M
						v332 = v326 - int32(1636608432)
						if l1&int32(3) != 0 {
							if base.Ui32(int32(11)) < base.Ui32(v326) {
								v441 = l1
								v442 = v326
								v443 = v332
								v444 = v332
								v445 = v332
								for {
									v447 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
									v448 = v447 + v444
									v449 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
									v451 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
									v452 = v451 + v445
									v454 = int32(4)
									v456 = v449 + v443 - v452 ^ base.I32_rotl(v452, v454)
									v460 = v448 - v456 ^ base.I32_rotl(v456, int32(6))
									v461 = v452 + v448
									v462 = v456 + v461
									v463 = v460 + v462
									v467 = v461 - v460 ^ base.I32_rotl(v460, int32(8))
									v471 = v462 - v467 ^ base.I32_rotl(v467, int32(16))
									v475 = v463 - v471 ^ base.I32_rotl(v471, int32(19))
									v476 = v467 + v463
									v477 = v471 + v476
									v478 = v475 + v477
									v482 = v476 - v475 ^ base.I32_rotl(v475, v454)
									v483 = int32(12)
									v484 = v441 + v483
									v486 = v442 - v483
									if base.Ui32(int32(11)) < base.Ui32(v486) {
										v441 = v484
										v442 = v486
										v443 = v477
										v444 = v478
										v445 = v482
										continue
									} else {
										break
									}
									break
								}
								v489 = v484
								v490 = v486
								v491 = v477
								v492 = v478
								v493 = v482
							} else {
								v489 = l1
								v490 = v326
								v491 = v332
								v492 = v332
								v493 = v332
							}
							switch v490 - int32(1) {
							case 0:
								v552 = v491
								v553 = v492
								v554 = v493
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 1:
								v545 = v491
								v546 = v492
								v547 = v493
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 2:
								v538 = v491
								v539 = v492
								v540 = v493
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 3:
								v532 = v492
								v533 = v493
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 4:
								v528 = v492
								v529 = v493
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 5:
								v522 = v492
								v523 = v493
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
								v528 = v524<<(uint(int32(8))%32) + v522
								v529 = v523
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 6:
								v516 = v492
								v517 = v493
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+6)))
								v522 = v518<<(uint(int32(16))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
								v528 = v524<<(uint(int32(8))%32) + v522
								v529 = v523
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 7:
								v511 = v493
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+7)))
								v516 = v512<<(uint(int32(24))%32) + v492
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+6)))
								v522 = v518<<(uint(int32(16))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
								v528 = v524<<(uint(int32(8))%32) + v522
								v529 = v523
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 8:
								v506 = v493
								v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+8)))
								v511 = v507<<(uint(int32(8))%32) + v506
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+7)))
								v516 = v512<<(uint(int32(24))%32) + v492
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+6)))
								v522 = v518<<(uint(int32(16))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
								v528 = v524<<(uint(int32(8))%32) + v522
								v529 = v523
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 9:
								v501 = v493
								v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+9)))
								v506 = v502<<(uint(int32(16))%32) + v501
								v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+8)))
								v511 = v507<<(uint(int32(8))%32) + v506
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+7)))
								v516 = v512<<(uint(int32(24))%32) + v492
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+6)))
								v522 = v518<<(uint(int32(16))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
								v528 = v524<<(uint(int32(8))%32) + v522
								v529 = v523
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							case 10:
								v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+10)))
								v501 = v497<<(uint(int32(24))%32) + v493
								v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+9)))
								v506 = v502<<(uint(int32(16))%32) + v501
								v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+8)))
								v511 = v507<<(uint(int32(8))%32) + v506
								v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+7)))
								v516 = v512<<(uint(int32(24))%32) + v492
								v517 = v511
								v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+6)))
								v522 = v518<<(uint(int32(16))%32) + v516
								v523 = v517
								v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
								v528 = v524<<(uint(int32(8))%32) + v522
								v529 = v523
								v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
								v532 = v528 + v530
								v533 = v529
								v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
								v538 = v534<<(uint(int32(24))%32) + v491
								v539 = v532
								v540 = v533
								v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
								v545 = v541<<(uint(int32(16))%32) + v538
								v546 = v539
								v547 = v540
								v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
								v552 = v548<<(uint(int32(8))%32) + v545
								v553 = v546
								v554 = v547
								v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
								v559 = v552 + v555
								v560 = v553
								v561 = v554
							default:
								v559 = v491
								v560 = v492
								v561 = v493
							}
						} else {
							if base.Ui32(v326) < base.Ui32(int32(12)) {
								v387 = l1
								v388 = v326
								v389 = v332
								v390 = v332
								v391 = v332
							} else {
								v339 = l1
								v340 = v326
								v341 = v332
								v342 = v332
								v343 = v332
								for {
									v345 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
									v346 = v345 + v342
									v347 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
									v349 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
									v350 = v349 + v343
									v352 = int32(4)
									v354 = v347 + v341 - v350 ^ base.I32_rotl(v350, v352)
									v358 = v346 - v354 ^ base.I32_rotl(v354, int32(6))
									v359 = v350 + v346
									v360 = v354 + v359
									v361 = v358 + v360
									v365 = v359 - v358 ^ base.I32_rotl(v358, int32(8))
									v369 = v360 - v365 ^ base.I32_rotl(v365, int32(16))
									v373 = v361 - v369 ^ base.I32_rotl(v369, int32(19))
									v374 = v365 + v361
									v375 = v369 + v374
									v376 = v373 + v375
									v380 = v374 - v373 ^ base.I32_rotl(v373, v352)
									v381 = int32(12)
									v382 = v339 + v381
									v384 = v340 - v381
									if base.Ui32(int32(11)) < base.Ui32(v384) {
										v339 = v382
										v340 = v384
										v341 = v375
										v342 = v376
										v343 = v380
										continue
									} else {
										break
									}
									break
								}
								v387 = v382
								v388 = v384
								v389 = v375
								v390 = v376
								v391 = v380
							}
							switch v388 - int32(1) {
							case 0:
								v438 = v389
								v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
								v559 = v438 + v439
								v560 = v390
								v561 = v391
							case 1:
								v433 = v389
								v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
								v438 = v434<<(uint(int32(8))%32) + v433
								v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
								v559 = v438 + v439
								v560 = v390
								v561 = v391
							case 2:
								v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+2)))
								v433 = v429<<(uint(int32(16))%32) + v389
								v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
								v438 = v434<<(uint(int32(8))%32) + v433
								v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
								v559 = v438 + v439
								v560 = v390
								v561 = v391
							case 3:
								v426 = v390
								v427 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v559 = v427 + v389
								v560 = v426
								v561 = v391
							case 4:
								v423 = v390
								v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+4)))
								v426 = v423 + v424
								v427 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v559 = v427 + v389
								v560 = v426
								v561 = v391
							case 5:
								v418 = v390
								v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+5)))
								v423 = v419<<(uint(int32(8))%32) + v418
								v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+4)))
								v426 = v423 + v424
								v427 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v559 = v427 + v389
								v560 = v426
								v561 = v391
							case 6:
								v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+6)))
								v418 = v414<<(uint(int32(16))%32) + v390
								v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+5)))
								v423 = v419<<(uint(int32(8))%32) + v418
								v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+4)))
								v426 = v423 + v424
								v427 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v559 = v427 + v389
								v560 = v426
								v561 = v391
							case 7:
								v409 = v391
								v410 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v412 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
								v559 = v410 + v389
								v560 = v412 + v390
								v561 = v409
							case 8:
								v404 = v391
								v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+8)))
								v409 = v405<<(uint(int32(8))%32) + v404
								v410 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v412 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
								v559 = v410 + v389
								v560 = v412 + v390
								v561 = v409
							case 9:
								v399 = v391
								v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+9)))
								v404 = v400<<(uint(int32(16))%32) + v399
								v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+8)))
								v409 = v405<<(uint(int32(8))%32) + v404
								v410 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v412 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
								v559 = v410 + v389
								v560 = v412 + v390
								v561 = v409
							case 10:
								v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+10)))
								v399 = v395<<(uint(int32(24))%32) + v391
								v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+9)))
								v404 = v400<<(uint(int32(16))%32) + v399
								v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+8)))
								v409 = v405<<(uint(int32(8))%32) + v404
								v410 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
								v412 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
								v559 = v410 + v389
								v560 = v412 + v390
								v561 = v409
							default:
								v559 = v389
								v560 = v390
								v561 = v391
							}
						}
						v564 = int32(14)
						v566 = v560 ^ v561 - base.I32_rotl(v560, v564)
						v570 = v566 ^ v559 - base.I32_rotl(v566, int32(11))
						v574 = v570 ^ v560 - base.I32_rotl(v570, int32(25))
						v578 = v574 ^ v566 - base.I32_rotl(v574, int32(16))
						v582 = v578 ^ v570 - base.I32_rotl(v578, int32(4))
						v586 = v582 ^ v574 - base.I32_rotl(v582, v564)
						v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v592 = base.I32_rem_u_s(v586^v578-base.I32_rotl(v586, int32(24)), v591)
						v596 = *(*int32)(unsafe.Add(mBase, uint32(v14+v592<<(uint(int32(2))%32))))
						F_TempTablespacePath(m, v9+int32(2096), v596)
						mBase = m.M
						v598 = m.ExcPending
						if v598 != 0 {
							return int32(0)
						} else {
							F_TempTablespacePath(m, v9+int32(3120), v596)
							mBase = m.M
							v602 = m.ExcPending
							if v602 != 0 {
								return int32(0)
							} else {
								v603 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(235464)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v603
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(3120)
								v614 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(106036), v9)
								mBase = m.M
								v615 = m.ExcPending
								if v615 != 0 {
									return int32(0)
								} else {
									v617 = v9 + int32(2096)
									v618 = m.G0
									v620 = v618 - int32(32)
									m.G0 = v620
									v623 = v9 + int32(48)
									v625 = *(*int32)(unsafe.Add(mBase, _consts[387]))
									v626 = F_mkdir(m, v623, v625)
									mBase = m.M
									if int32(0) <= v626 {
										m.G0 = v620 + int32(32)
										v689 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
										mBase = m.M
										v690 = m.ExcPending
										if v690 != 0 {
											return int32(0)
										} else {
											v693 = v689
											m.G0 = v9 + int32(4144)
											return v693
										}
									} else {
										v630 = *(*int32)(unsafe.Add(mBase, _consts[159]))
										if v630 == int32(20) {
											m.G0 = v620 + int32(32)
											v689 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
											mBase = m.M
											v690 = m.ExcPending
											if v690 != 0 {
												return int32(0)
											} else {
												v693 = v689
												m.G0 = v9 + int32(4144)
												return v693
											}
										} else {
											v634 = *(*int32)(unsafe.Add(mBase, _consts[387]))
											v635 = F_mkdir(m, v617, v634)
											mBase = m.M
											if v635 < int32(0) {
												v639 = *(*int32)(unsafe.Add(mBase, _consts[159]))
												if v639 != int32(20) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v657 = m.ExcPending
													if v657 != 0 {
														return int32(0)
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v659 = m.ExcPending
														if v659 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v620)+16)) = v617
															F_errmsg(m, int32(296437), v620+int32(16))
															mBase = m.M
															v665 = m.ExcPending
															if v665 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(500025), int32(1685), int32(213457))
																mBase = m.M
																v670 = m.ExcPending
																if v670 != 0 {
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
													v643 = *(*int32)(unsafe.Add(mBase, _consts[387]))
													v644 = F_mkdir(m, v623, v643)
													mBase = m.M
													if int32(0) <= v644 {
														m.G0 = v620 + int32(32)
														v689 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
														mBase = m.M
														v690 = m.ExcPending
														if v690 != 0 {
															return int32(0)
														} else {
															v693 = v689
															m.G0 = v9 + int32(4144)
															return v693
														}
													} else {
														v648 = *(*int32)(unsafe.Add(mBase, _consts[159]))
														if v648 != int32(20) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v674 = m.ExcPending
															if v674 != 0 {
																return int32(0)
															} else {
																F_errcode_for_file_access(m)
																mBase = m.M
																v676 = m.ExcPending
																if v676 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v623
																	F_errmsg(m, int32(296391), v620)
																	mBase = m.M
																	v680 = m.ExcPending
																	if v680 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(500025), int32(1692), int32(213457))
																		mBase = m.M
																		v685 = m.ExcPending
																		if v685 != 0 {
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
															m.G0 = v620 + int32(32)
															v689 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
															mBase = m.M
															v690 = m.ExcPending
															if v690 != 0 {
																return int32(0)
															} else {
																v693 = v689
																m.G0 = v9 + int32(4144)
																return v693
															}
														}
													}
												}
											} else {
												v643 = *(*int32)(unsafe.Add(mBase, _consts[387]))
												v644 = F_mkdir(m, v623, v643)
												mBase = m.M
												if int32(0) <= v644 {
													m.G0 = v620 + int32(32)
													v689 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
													mBase = m.M
													v690 = m.ExcPending
													if v690 != 0 {
														return int32(0)
													} else {
														v693 = v689
														m.G0 = v9 + int32(4144)
														return v693
													}
												} else {
													v648 = *(*int32)(unsafe.Add(mBase, _consts[159]))
													if v648 != int32(20) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v674 = m.ExcPending
														if v674 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v676 = m.ExcPending
															if v676 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v620))) = v623
																F_errmsg(m, int32(296391), v620)
																mBase = m.M
																v680 = m.ExcPending
																if v680 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(500025), int32(1692), int32(213457))
																	mBase = m.M
																	v685 = m.ExcPending
																	if v685 != 0 {
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
														m.G0 = v620 + int32(32)
														v689 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
														mBase = m.M
														v690 = m.ExcPending
														if v690 != 0 {
															return int32(0)
														} else {
															v693 = v689
															m.G0 = v9 + int32(4144)
															return v693
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
						v693 = v320
						m.G0 = v9 + int32(4144)
						return v693
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[354]))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v13
	v15 = int32(4431648)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16
	v22 = base.I32_rem_u_s(v16+int32(1), int32(2147483647))
	*(*int32)(unsafe.Add(mBase, _consts[764])) = v22
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v27 = l0 + int32(12)
		v28 = int32(8)
		v30 = *(*int32)(unsafe.Add(mBase, _consts[409]))
		if v28 < v30 {
			v33 = v28
		} else {
			v33 = v30
		}
		if int32(0) < v33 {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[410]))
			if base.Ui32(int32(4)) <= base.Ui32(v33) {
				v44 = v2
				v50 = v2
				for {
					v54 = v44 << (uint(int32(2)) % 32)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v54))) = v57
					v59 = int32(4)
					v60 = v54 | v59
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v37+v60)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v60))) = v63
					v66 = v54 | int32(8)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v37+v66)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v66))) = v69
					v72 = v54 | int32(12)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v72+v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v72))) = v75
					v78 = v44 + v59
					v80 = v50 + v59
					if v80 != v33&int32(2147483644) {
						v44 = v78
						v50 = v80
						continue
					} else {
						break
					}
					break
				}
				v84 = v78
			} else {
				v84 = v2
			}
			v94 = v33 & int32(3)
			if v94 != 0 {
				v97 = v84
				v104 = v2
				for {
					v107 = v97 << (uint(int32(2)) % 32)
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v107))) = v110
					v112 = int32(1)
					v115 = v104 + v112
					if v115 != v94 {
						v97 = v97 + v112
						v104 = v115
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v140 = v33
		} else {
			v140 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140
		if v140 != 0 {
			if int32(0) < v140 {
				v145 = v140
				v151 = v2
				for {
					v157 = v27 + v151<<(uint(int32(2))%32)
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
					if v158 == int32(0) {
						v162 = *(*int32)(unsafe.Add(mBase, _consts[173]))
						*(*int32)(unsafe.Add(mBase, uint32(v157))) = v162
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v165 = v164
					} else {
						v165 = v145
					}
					v167 = v151 + int32(1)
					if v167 < v165 {
						v145 = v165
						v151 = v167
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return
		} else {
			v181 = *(*int32)(unsafe.Add(mBase, _consts[173]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
			return
		}
	}
}
