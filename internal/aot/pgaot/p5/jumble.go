package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
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
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
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
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
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
	var v498 int32
	_ = v498
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
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
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v380 = v12
	} else {
		v13 = int32(4)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v15-int32(1021)) < base.Ui32(v13) {
			v25 = v15
			v27 = v13
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v32 = int32(1024)
					v39 = int32(-1636607408)
					if v14&int32(3) != 0 {
						v85 = v14
						v86 = v32
						v88 = v39
						v89 = v39
						v90 = v39
						for {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
							v93 = v92 + v89
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							v97 = v96 + v90
							v99 = int32(4)
							v101 = v94 + v88 - v97 ^ base.I32_rotl(v97, v99)
							v105 = v93 - v101 ^ base.I32_rotl(v101, int32(6))
							v106 = v97 + v93
							v107 = v101 + v106
							v108 = v105 + v107
							v112 = v106 - v105 ^ base.I32_rotl(v105, int32(8))
							v116 = v107 - v112 ^ base.I32_rotl(v112, int32(16))
							v120 = v108 - v116 ^ base.I32_rotl(v116, int32(19))
							v121 = v112 + v108
							v122 = v116 + v121
							v123 = v120 + v122
							v127 = v121 - v120 ^ base.I32_rotl(v120, v99)
							v128 = int32(12)
							v129 = v85 + v128
							v131 = v86 - v128
							if base.Ui32(int32(11)) < base.Ui32(v131) {
								v85 = v129
								v86 = v131
								v88 = v122
								v89 = v123
								v90 = v127
								continue
							} else {
								break
							}
							break
						}
						switch v131 - int32(1) {
						case 0:
							v304 = v122
							v305 = v123
							v306 = v127
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 1:
							v297 = v122
							v298 = v123
							v299 = v127
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 2:
							v290 = v122
							v291 = v123
							v292 = v127
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 3:
							v284 = v123
							v285 = v127
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 4:
							v280 = v123
							v281 = v127
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 5:
							v274 = v123
							v275 = v127
							v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
							v280 = v276<<(uint(int32(8))%32) + v274
							v281 = v275
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 6:
							v268 = v123
							v269 = v127
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
							v274 = v270<<(uint(int32(16))%32) + v268
							v275 = v269
							v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
							v280 = v276<<(uint(int32(8))%32) + v274
							v281 = v275
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 7:
							v263 = v127
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
							v268 = v264<<(uint(int32(24))%32) + v123
							v269 = v263
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
							v274 = v270<<(uint(int32(16))%32) + v268
							v275 = v269
							v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
							v280 = v276<<(uint(int32(8))%32) + v274
							v281 = v275
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 8:
							v258 = v127
							v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
							v263 = v259<<(uint(int32(8))%32) + v258
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
							v268 = v264<<(uint(int32(24))%32) + v123
							v269 = v263
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
							v274 = v270<<(uint(int32(16))%32) + v268
							v275 = v269
							v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
							v280 = v276<<(uint(int32(8))%32) + v274
							v281 = v275
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 9:
							v253 = v127
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
							v258 = v254<<(uint(int32(16))%32) + v253
							v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
							v263 = v259<<(uint(int32(8))%32) + v258
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
							v268 = v264<<(uint(int32(24))%32) + v123
							v269 = v263
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
							v274 = v270<<(uint(int32(16))%32) + v268
							v275 = v269
							v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
							v280 = v276<<(uint(int32(8))%32) + v274
							v281 = v275
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						case 10:
							v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+10)))
							v253 = v249<<(uint(int32(24))%32) + v127
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
							v258 = v254<<(uint(int32(16))%32) + v253
							v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
							v263 = v259<<(uint(int32(8))%32) + v258
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
							v268 = v264<<(uint(int32(24))%32) + v123
							v269 = v263
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
							v274 = v270<<(uint(int32(16))%32) + v268
							v275 = v269
							v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
							v280 = v276<<(uint(int32(8))%32) + v274
							v281 = v275
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
							v284 = v280 + v282
							v285 = v281
							v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
							v290 = v286<<(uint(int32(24))%32) + v122
							v291 = v284
							v292 = v285
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
							v297 = v293<<(uint(int32(16))%32) + v290
							v298 = v291
							v299 = v292
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
							v304 = v300<<(uint(int32(8))%32) + v297
							v305 = v298
							v306 = v299
							v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
							v312 = v304 + v307
							v313 = v305
							v314 = v306
						default:
							v312 = v122
							v313 = v123
							v314 = v127
						}
					} else {
						v145 = v14
						v146 = v32
						v148 = v39
						v149 = v39
						v150 = v39
						for {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
							v153 = v152 + v149
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
							v157 = v156 + v150
							v159 = int32(4)
							v161 = v154 + v148 - v157 ^ base.I32_rotl(v157, v159)
							v165 = v153 - v161 ^ base.I32_rotl(v161, int32(6))
							v166 = v157 + v153
							v167 = v161 + v166
							v168 = v165 + v167
							v172 = v166 - v165 ^ base.I32_rotl(v165, int32(8))
							v176 = v167 - v172 ^ base.I32_rotl(v172, int32(16))
							v180 = v168 - v176 ^ base.I32_rotl(v176, int32(19))
							v181 = v172 + v168
							v182 = v176 + v181
							v183 = v180 + v182
							v187 = v181 - v180 ^ base.I32_rotl(v180, v159)
							v188 = int32(12)
							v189 = v145 + v188
							v191 = v146 - v188
							if base.Ui32(int32(11)) < base.Ui32(v191) {
								v145 = v189
								v146 = v191
								v148 = v182
								v149 = v183
								v150 = v187
								continue
							} else {
								break
							}
							break
						}
						switch v191 - int32(1) {
						case 0:
							v246 = v182
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
							v312 = v246 + v247
							v313 = v183
							v314 = v187
						case 1:
							v241 = v182
							v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
							v246 = v242<<(uint(int32(8))%32) + v241
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
							v312 = v246 + v247
							v313 = v183
							v314 = v187
						case 2:
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+2)))
							v241 = v237<<(uint(int32(16))%32) + v182
							v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
							v246 = v242<<(uint(int32(8))%32) + v241
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
							v312 = v246 + v247
							v313 = v183
							v314 = v187
						case 3:
							v234 = v183
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v312 = v235 + v182
							v313 = v234
							v314 = v187
						case 4:
							v231 = v183
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
							v234 = v231 + v232
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v312 = v235 + v182
							v313 = v234
							v314 = v187
						case 5:
							v226 = v183
							v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+5)))
							v231 = v227<<(uint(int32(8))%32) + v226
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
							v234 = v231 + v232
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v312 = v235 + v182
							v313 = v234
							v314 = v187
						case 6:
							v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+6)))
							v226 = v222<<(uint(int32(16))%32) + v183
							v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+5)))
							v231 = v227<<(uint(int32(8))%32) + v226
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
							v234 = v231 + v232
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v312 = v235 + v182
							v313 = v234
							v314 = v187
						case 7:
							v217 = v187
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v220 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
							v312 = v218 + v182
							v313 = v220 + v183
							v314 = v217
						case 8:
							v212 = v187
							v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)))
							v217 = v213<<(uint(int32(8))%32) + v212
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v220 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
							v312 = v218 + v182
							v313 = v220 + v183
							v314 = v217
						case 9:
							v207 = v187
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+9)))
							v212 = v208<<(uint(int32(16))%32) + v207
							v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)))
							v217 = v213<<(uint(int32(8))%32) + v212
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v220 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
							v312 = v218 + v182
							v313 = v220 + v183
							v314 = v217
						case 10:
							v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+10)))
							v207 = v203<<(uint(int32(24))%32) + v187
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+9)))
							v212 = v208<<(uint(int32(16))%32) + v207
							v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)))
							v217 = v213<<(uint(int32(8))%32) + v212
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
							v220 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
							v312 = v218 + v182
							v313 = v220 + v183
							v314 = v217
						default:
							v312 = v182
							v313 = v183
							v314 = v187
						}
					}
					v317 = int32(14)
					v319 = v313 ^ v314 - base.I32_rotl(v313, v317)
					v323 = v319 ^ v312 - base.I32_rotl(v319, int32(11))
					v327 = v323 ^ v313 - base.I32_rotl(v323, int32(25))
					v331 = v327 ^ v319 - base.I32_rotl(v327, int32(16))
					v335 = v331 ^ v323 - base.I32_rotl(v331, int32(4))
					v339 = v335 ^ v327 - base.I32_rotl(v335, v317)
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = base.I64_extend_i32_u(v339)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v339^v331-base.I32_rotl(v339, int32(24)))
					v351 = int32(8)
				} else {
					v351 = v25
				}
				v353 = int32(1024) - v351
				if base.Ui32(v27) < base.Ui32(v353) {
					v355 = v27
				} else {
					v355 = v353
				}
				if v355 != 0 {
					base.MemoryCopy(m, v351+v14, v29, v355)
				} else {
				}
				v359 = v351 + v355
				v360 = v27 - v355
				if v360 != 0 {
					v25 = v359
					v27 = v360
					v29 = v355 + v29
					continue
				} else {
					break
				}
				break
			}
			v369 = v359
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15+v14))) = v9
			v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v369 = v363 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369
		v380 = v369
	}
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(1024)-v380) < base.Ui32(l2) {
		v390 = l1
		v391 = l2
		v392 = v380
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v392) {
				v399 = int32(1024)
				v406 = int32(-1636607408)
				if v385&int32(3) != 0 {
					v452 = v385
					v453 = v399
					v455 = v406
					v456 = v406
					v457 = v406
					for {
						v459 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
						v460 = v459 + v456
						v461 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
						v463 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
						v464 = v463 + v457
						v466 = int32(4)
						v468 = v461 + v455 - v464 ^ base.I32_rotl(v464, v466)
						v472 = v460 - v468 ^ base.I32_rotl(v468, int32(6))
						v473 = v464 + v460
						v474 = v468 + v473
						v475 = v472 + v474
						v479 = v473 - v472 ^ base.I32_rotl(v472, int32(8))
						v483 = v474 - v479 ^ base.I32_rotl(v479, int32(16))
						v487 = v475 - v483 ^ base.I32_rotl(v483, int32(19))
						v488 = v479 + v475
						v489 = v483 + v488
						v490 = v487 + v489
						v494 = v488 - v487 ^ base.I32_rotl(v487, v466)
						v495 = int32(12)
						v496 = v452 + v495
						v498 = v453 - v495
						if base.Ui32(int32(11)) < base.Ui32(v498) {
							v452 = v496
							v453 = v498
							v455 = v489
							v456 = v490
							v457 = v494
							continue
						} else {
							break
						}
						break
					}
					switch v498 - int32(1) {
					case 0:
						v671 = v489
						v672 = v490
						v673 = v494
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 1:
						v664 = v489
						v665 = v490
						v666 = v494
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 2:
						v657 = v489
						v658 = v490
						v659 = v494
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 3:
						v651 = v490
						v652 = v494
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 4:
						v647 = v490
						v648 = v494
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 5:
						v641 = v490
						v642 = v494
						v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
						v647 = v643<<(uint(int32(8))%32) + v641
						v648 = v642
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 6:
						v635 = v490
						v636 = v494
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
						v641 = v637<<(uint(int32(16))%32) + v635
						v642 = v636
						v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
						v647 = v643<<(uint(int32(8))%32) + v641
						v648 = v642
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 7:
						v630 = v494
						v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+7)))
						v635 = v631<<(uint(int32(24))%32) + v490
						v636 = v630
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
						v641 = v637<<(uint(int32(16))%32) + v635
						v642 = v636
						v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
						v647 = v643<<(uint(int32(8))%32) + v641
						v648 = v642
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 8:
						v625 = v494
						v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
						v630 = v626<<(uint(int32(8))%32) + v625
						v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+7)))
						v635 = v631<<(uint(int32(24))%32) + v490
						v636 = v630
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
						v641 = v637<<(uint(int32(16))%32) + v635
						v642 = v636
						v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
						v647 = v643<<(uint(int32(8))%32) + v641
						v648 = v642
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 9:
						v620 = v494
						v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+9)))
						v625 = v621<<(uint(int32(16))%32) + v620
						v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
						v630 = v626<<(uint(int32(8))%32) + v625
						v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+7)))
						v635 = v631<<(uint(int32(24))%32) + v490
						v636 = v630
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
						v641 = v637<<(uint(int32(16))%32) + v635
						v642 = v636
						v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
						v647 = v643<<(uint(int32(8))%32) + v641
						v648 = v642
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					case 10:
						v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+10)))
						v620 = v616<<(uint(int32(24))%32) + v494
						v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+9)))
						v625 = v621<<(uint(int32(16))%32) + v620
						v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
						v630 = v626<<(uint(int32(8))%32) + v625
						v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+7)))
						v635 = v631<<(uint(int32(24))%32) + v490
						v636 = v630
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
						v641 = v637<<(uint(int32(16))%32) + v635
						v642 = v636
						v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
						v647 = v643<<(uint(int32(8))%32) + v641
						v648 = v642
						v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
						v651 = v647 + v649
						v652 = v648
						v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
						v657 = v653<<(uint(int32(24))%32) + v489
						v658 = v651
						v659 = v652
						v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
						v664 = v660<<(uint(int32(16))%32) + v657
						v665 = v658
						v666 = v659
						v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
						v671 = v667<<(uint(int32(8))%32) + v664
						v672 = v665
						v673 = v666
						v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
						v679 = v671 + v674
						v680 = v672
						v681 = v673
					default:
						v679 = v489
						v680 = v490
						v681 = v494
					}
				} else {
					v512 = v385
					v513 = v399
					v515 = v406
					v516 = v406
					v517 = v406
					for {
						v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
						v520 = v519 + v516
						v521 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
						v523 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
						v524 = v523 + v517
						v526 = int32(4)
						v528 = v521 + v515 - v524 ^ base.I32_rotl(v524, v526)
						v532 = v520 - v528 ^ base.I32_rotl(v528, int32(6))
						v533 = v524 + v520
						v534 = v528 + v533
						v535 = v532 + v534
						v539 = v533 - v532 ^ base.I32_rotl(v532, int32(8))
						v543 = v534 - v539 ^ base.I32_rotl(v539, int32(16))
						v547 = v535 - v543 ^ base.I32_rotl(v543, int32(19))
						v548 = v539 + v535
						v549 = v543 + v548
						v550 = v547 + v549
						v554 = v548 - v547 ^ base.I32_rotl(v547, v526)
						v555 = int32(12)
						v556 = v512 + v555
						v558 = v513 - v555
						if base.Ui32(int32(11)) < base.Ui32(v558) {
							v512 = v556
							v513 = v558
							v515 = v549
							v516 = v550
							v517 = v554
							continue
						} else {
							break
						}
						break
					}
					switch v558 - int32(1) {
					case 0:
						v613 = v549
						v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
						v679 = v613 + v614
						v680 = v550
						v681 = v554
					case 1:
						v608 = v549
						v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
						v613 = v609<<(uint(int32(8))%32) + v608
						v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
						v679 = v613 + v614
						v680 = v550
						v681 = v554
					case 2:
						v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+2)))
						v608 = v604<<(uint(int32(16))%32) + v549
						v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
						v613 = v609<<(uint(int32(8))%32) + v608
						v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
						v679 = v613 + v614
						v680 = v550
						v681 = v554
					case 3:
						v601 = v550
						v602 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v679 = v602 + v549
						v680 = v601
						v681 = v554
					case 4:
						v598 = v550
						v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
						v601 = v598 + v599
						v602 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v679 = v602 + v549
						v680 = v601
						v681 = v554
					case 5:
						v593 = v550
						v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+5)))
						v598 = v594<<(uint(int32(8))%32) + v593
						v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
						v601 = v598 + v599
						v602 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v679 = v602 + v549
						v680 = v601
						v681 = v554
					case 6:
						v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+6)))
						v593 = v589<<(uint(int32(16))%32) + v550
						v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+5)))
						v598 = v594<<(uint(int32(8))%32) + v593
						v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
						v601 = v598 + v599
						v602 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v679 = v602 + v549
						v680 = v601
						v681 = v554
					case 7:
						v584 = v554
						v585 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
						v679 = v585 + v549
						v680 = v587 + v550
						v681 = v584
					case 8:
						v579 = v554
						v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+8)))
						v584 = v580<<(uint(int32(8))%32) + v579
						v585 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
						v679 = v585 + v549
						v680 = v587 + v550
						v681 = v584
					case 9:
						v574 = v554
						v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+9)))
						v579 = v575<<(uint(int32(16))%32) + v574
						v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+8)))
						v584 = v580<<(uint(int32(8))%32) + v579
						v585 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
						v679 = v585 + v549
						v680 = v587 + v550
						v681 = v584
					case 10:
						v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+10)))
						v574 = v570<<(uint(int32(24))%32) + v554
						v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+9)))
						v579 = v575<<(uint(int32(16))%32) + v574
						v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+8)))
						v584 = v580<<(uint(int32(8))%32) + v579
						v585 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
						v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
						v679 = v585 + v549
						v680 = v587 + v550
						v681 = v584
					default:
						v679 = v549
						v680 = v550
						v681 = v554
					}
				}
				v684 = int32(14)
				v686 = v680 ^ v681 - base.I32_rotl(v680, v684)
				v690 = v686 ^ v679 - base.I32_rotl(v686, int32(11))
				v694 = v690 ^ v680 - base.I32_rotl(v690, int32(25))
				v698 = v694 ^ v686 - base.I32_rotl(v694, int32(16))
				v702 = v698 ^ v690 - base.I32_rotl(v698, int32(4))
				v706 = v702 ^ v694 - base.I32_rotl(v702, v684)
				*(*int64)(unsafe.Add(mBase, uint32(v385))) = base.I64_extend_i32_u(v706)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v706^v698-base.I32_rotl(v706, int32(24)))
				v718 = int32(8)
			} else {
				v718 = v392
			}
			v720 = int32(1024) - v718
			if base.Ui32(v391) < base.Ui32(v720) {
				v722 = v391
			} else {
				v722 = v720
			}
			if v722 != 0 {
				base.MemoryCopy(m, v718+v385, v390, v722)
			} else {
			}
			v726 = v718 + v722
			v727 = v391 - v722
			if v727 != 0 {
				v390 = v390 + v722
				v391 = v727
				v392 = v726
				continue
			} else {
				break
			}
			break
		}
		v735 = v726
	} else {
		if l2 != 0 {
			base.MemoryCopy(m, v380+v385, l1, l2)
		} else {
		}
		v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v735 = v730 + l2
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v735
	return
}
func F_JumbleQuery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
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
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
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
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
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
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
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
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v729 int64
	_ = v729
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_palloc(m, int32(1024))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
			v23 = F_palloc(m, int32(384))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v25)
				F__jumbleNode(m, v10, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
					if v34 != 0 {
						v35 = int32(4)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						if base.Ui32(v37-int32(1021)) < base.Ui32(v35) {
							v46 = v37
							v48 = v35
							v50 = v10 + int32(28)
							for {
								if base.Ui32(int32(1024)) <= base.Ui32(v46) {
									v54 = int32(1024)
									v61 = int32(-1636607408)
									if v36&int32(3) != 0 {
										v107 = v36
										v108 = v54
										v110 = v61
										v111 = v61
										v112 = v61
										for {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
											v115 = v114 + v111
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
											v119 = v118 + v112
											v121 = int32(4)
											v123 = v116 + v110 - v119 ^ base.I32_rotl(v119, v121)
											v127 = v115 - v123 ^ base.I32_rotl(v123, int32(6))
											v128 = v119 + v115
											v129 = v123 + v128
											v130 = v127 + v129
											v134 = v128 - v127 ^ base.I32_rotl(v127, int32(8))
											v138 = v129 - v134 ^ base.I32_rotl(v134, int32(16))
											v142 = v130 - v138 ^ base.I32_rotl(v138, int32(19))
											v143 = v134 + v130
											v144 = v138 + v143
											v145 = v142 + v144
											v149 = v143 - v142 ^ base.I32_rotl(v142, v121)
											v150 = int32(12)
											v151 = v107 + v150
											v153 = v108 - v150
											if base.Ui32(int32(11)) < base.Ui32(v153) {
												v107 = v151
												v108 = v153
												v110 = v144
												v111 = v145
												v112 = v149
												continue
											} else {
												break
											}
											break
										}
										switch v153 - int32(1) {
										case 0:
											v326 = v144
											v327 = v145
											v328 = v149
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 1:
											v319 = v144
											v320 = v145
											v321 = v149
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 2:
											v312 = v144
											v313 = v145
											v314 = v149
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 3:
											v306 = v145
											v307 = v149
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 4:
											v302 = v145
											v303 = v149
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 5:
											v296 = v145
											v297 = v149
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
											v302 = v298<<(uint(int32(8))%32) + v296
											v303 = v297
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 6:
											v290 = v145
											v291 = v149
											v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
											v296 = v292<<(uint(int32(16))%32) + v290
											v297 = v291
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
											v302 = v298<<(uint(int32(8))%32) + v296
											v303 = v297
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 7:
											v285 = v149
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
											v290 = v286<<(uint(int32(24))%32) + v145
											v291 = v285
											v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
											v296 = v292<<(uint(int32(16))%32) + v290
											v297 = v291
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
											v302 = v298<<(uint(int32(8))%32) + v296
											v303 = v297
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 8:
											v280 = v149
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
											v285 = v281<<(uint(int32(8))%32) + v280
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
											v290 = v286<<(uint(int32(24))%32) + v145
											v291 = v285
											v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
											v296 = v292<<(uint(int32(16))%32) + v290
											v297 = v291
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
											v302 = v298<<(uint(int32(8))%32) + v296
											v303 = v297
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 9:
											v275 = v149
											v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+9)))
											v280 = v276<<(uint(int32(16))%32) + v275
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
											v285 = v281<<(uint(int32(8))%32) + v280
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
											v290 = v286<<(uint(int32(24))%32) + v145
											v291 = v285
											v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
											v296 = v292<<(uint(int32(16))%32) + v290
											v297 = v291
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
											v302 = v298<<(uint(int32(8))%32) + v296
											v303 = v297
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										case 10:
											v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+10)))
											v275 = v271<<(uint(int32(24))%32) + v149
											v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+9)))
											v280 = v276<<(uint(int32(16))%32) + v275
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
											v285 = v281<<(uint(int32(8))%32) + v280
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
											v290 = v286<<(uint(int32(24))%32) + v145
											v291 = v285
											v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
											v296 = v292<<(uint(int32(16))%32) + v290
											v297 = v291
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
											v302 = v298<<(uint(int32(8))%32) + v296
											v303 = v297
											v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
											v306 = v302 + v304
											v307 = v303
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
											v312 = v308<<(uint(int32(24))%32) + v144
											v313 = v306
											v314 = v307
											v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
											v319 = v315<<(uint(int32(16))%32) + v312
											v320 = v313
											v321 = v314
											v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
											v326 = v322<<(uint(int32(8))%32) + v319
											v327 = v320
											v328 = v321
											v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
											v334 = v326 + v329
											v335 = v327
											v336 = v328
										default:
											v334 = v144
											v335 = v145
											v336 = v149
										}
									} else {
										v167 = v36
										v168 = v54
										v170 = v61
										v171 = v61
										v172 = v61
										for {
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
											v175 = v174 + v171
											v176 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
											v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
											v179 = v178 + v172
											v181 = int32(4)
											v183 = v176 + v170 - v179 ^ base.I32_rotl(v179, v181)
											v187 = v175 - v183 ^ base.I32_rotl(v183, int32(6))
											v188 = v179 + v175
											v189 = v183 + v188
											v190 = v187 + v189
											v194 = v188 - v187 ^ base.I32_rotl(v187, int32(8))
											v198 = v189 - v194 ^ base.I32_rotl(v194, int32(16))
											v202 = v190 - v198 ^ base.I32_rotl(v198, int32(19))
											v203 = v194 + v190
											v204 = v198 + v203
											v205 = v202 + v204
											v209 = v203 - v202 ^ base.I32_rotl(v202, v181)
											v210 = int32(12)
											v211 = v167 + v210
											v213 = v168 - v210
											if base.Ui32(int32(11)) < base.Ui32(v213) {
												v167 = v211
												v168 = v213
												v170 = v204
												v171 = v205
												v172 = v209
												continue
											} else {
												break
											}
											break
										}
										switch v213 - int32(1) {
										case 0:
											v268 = v204
											v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
											v334 = v268 + v269
											v335 = v205
											v336 = v209
										case 1:
											v263 = v204
											v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
											v268 = v264<<(uint(int32(8))%32) + v263
											v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
											v334 = v268 + v269
											v335 = v205
											v336 = v209
										case 2:
											v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+2)))
											v263 = v259<<(uint(int32(16))%32) + v204
											v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
											v268 = v264<<(uint(int32(8))%32) + v263
											v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
											v334 = v268 + v269
											v335 = v205
											v336 = v209
										case 3:
											v256 = v205
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v334 = v257 + v204
											v335 = v256
											v336 = v209
										case 4:
											v253 = v205
											v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
											v256 = v253 + v254
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v334 = v257 + v204
											v335 = v256
											v336 = v209
										case 5:
											v248 = v205
											v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
											v253 = v249<<(uint(int32(8))%32) + v248
											v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
											v256 = v253 + v254
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v334 = v257 + v204
											v335 = v256
											v336 = v209
										case 6:
											v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+6)))
											v248 = v244<<(uint(int32(16))%32) + v205
											v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
											v253 = v249<<(uint(int32(8))%32) + v248
											v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
											v256 = v253 + v254
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v334 = v257 + v204
											v335 = v256
											v336 = v209
										case 7:
											v239 = v209
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
											v334 = v240 + v204
											v335 = v242 + v205
											v336 = v239
										case 8:
											v234 = v209
											v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
											v239 = v235<<(uint(int32(8))%32) + v234
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
											v334 = v240 + v204
											v335 = v242 + v205
											v336 = v239
										case 9:
											v229 = v209
											v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+9)))
											v234 = v230<<(uint(int32(16))%32) + v229
											v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
											v239 = v235<<(uint(int32(8))%32) + v234
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
											v334 = v240 + v204
											v335 = v242 + v205
											v336 = v239
										case 10:
											v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+10)))
											v229 = v225<<(uint(int32(24))%32) + v209
											v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+9)))
											v234 = v230<<(uint(int32(16))%32) + v229
											v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
											v239 = v235<<(uint(int32(8))%32) + v234
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
											v334 = v240 + v204
											v335 = v242 + v205
											v336 = v239
										default:
											v334 = v204
											v335 = v205
											v336 = v209
										}
									}
									v339 = int32(14)
									v341 = v335 ^ v336 - base.I32_rotl(v335, v339)
									v345 = v341 ^ v334 - base.I32_rotl(v341, int32(11))
									v349 = v345 ^ v335 - base.I32_rotl(v345, int32(25))
									v353 = v349 ^ v341 - base.I32_rotl(v349, int32(16))
									v357 = v353 ^ v345 - base.I32_rotl(v353, int32(4))
									v361 = v357 ^ v349 - base.I32_rotl(v357, v339)
									*(*int64)(unsafe.Add(mBase, uint32(v36))) = base.I64_extend_i32_u(v361)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v361^v353-base.I32_rotl(v361, int32(24)))
									v373 = int32(8)
								} else {
									v373 = v46
								}
								v375 = int32(1024) - v373
								if base.Ui32(v48) < base.Ui32(v375) {
									v377 = v48
								} else {
									v377 = v375
								}
								if v377 != 0 {
									base.MemoryCopy(m, v373+v36, v50, v377)
								} else {
								}
								v381 = v373 + v377
								v382 = v48 - v377
								if v382 != 0 {
									v46 = v381
									v48 = v382
									v50 = v377 + v50
									continue
								} else {
									break
								}
								break
							}
							v390 = v381
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v37+v36))) = v34
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v390 = v385 + int32(4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v390
					} else {
					}
					v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
					if v407 == int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
					} else {
					}
					v412 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v413 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v420 = v413 - int32(1636608432)
					if v412&int32(3) != 0 {
						if base.Ui32(int32(11)) < base.Ui32(v413) {
							v466 = v412
							v467 = v413
							v469 = v420
							v470 = v420
							v471 = v420
							for {
								v473 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
								v474 = v473 + v470
								v475 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
								v477 = *(*int32)(unsafe.Add(mBase, uint32(v466)+8))
								v478 = v477 + v471
								v480 = int32(4)
								v482 = v475 + v469 - v478 ^ base.I32_rotl(v478, v480)
								v486 = v474 - v482 ^ base.I32_rotl(v482, int32(6))
								v487 = v478 + v474
								v488 = v482 + v487
								v489 = v486 + v488
								v493 = v487 - v486 ^ base.I32_rotl(v486, int32(8))
								v497 = v488 - v493 ^ base.I32_rotl(v493, int32(16))
								v501 = v489 - v497 ^ base.I32_rotl(v497, int32(19))
								v502 = v493 + v489
								v503 = v497 + v502
								v504 = v501 + v503
								v508 = v502 - v501 ^ base.I32_rotl(v501, v480)
								v509 = int32(12)
								v510 = v466 + v509
								v512 = v467 - v509
								if base.Ui32(int32(11)) < base.Ui32(v512) {
									v466 = v510
									v467 = v512
									v469 = v503
									v470 = v504
									v471 = v508
									continue
								} else {
									break
								}
								break
							}
							v515 = v510
							v516 = v512
							v518 = v503
							v519 = v504
							v520 = v508
						} else {
							v515 = v412
							v516 = v413
							v518 = v420
							v519 = v420
							v520 = v420
						}
						switch v516 - int32(1) {
						case 0:
							v685 = v518
							v686 = v519
							v687 = v520
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 1:
							v678 = v518
							v679 = v519
							v680 = v520
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 2:
							v671 = v518
							v672 = v519
							v673 = v520
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 3:
							v665 = v519
							v666 = v520
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 4:
							v661 = v519
							v662 = v520
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 5:
							v655 = v519
							v656 = v520
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+5)))
							v661 = v657<<(uint(int32(8))%32) + v655
							v662 = v656
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 6:
							v649 = v519
							v650 = v520
							v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+6)))
							v655 = v651<<(uint(int32(16))%32) + v649
							v656 = v650
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+5)))
							v661 = v657<<(uint(int32(8))%32) + v655
							v662 = v656
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 7:
							v644 = v520
							v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+7)))
							v649 = v645<<(uint(int32(24))%32) + v519
							v650 = v644
							v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+6)))
							v655 = v651<<(uint(int32(16))%32) + v649
							v656 = v650
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+5)))
							v661 = v657<<(uint(int32(8))%32) + v655
							v662 = v656
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 8:
							v639 = v520
							v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+8)))
							v644 = v640<<(uint(int32(8))%32) + v639
							v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+7)))
							v649 = v645<<(uint(int32(24))%32) + v519
							v650 = v644
							v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+6)))
							v655 = v651<<(uint(int32(16))%32) + v649
							v656 = v650
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+5)))
							v661 = v657<<(uint(int32(8))%32) + v655
							v662 = v656
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 9:
							v634 = v520
							v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+9)))
							v639 = v635<<(uint(int32(16))%32) + v634
							v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+8)))
							v644 = v640<<(uint(int32(8))%32) + v639
							v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+7)))
							v649 = v645<<(uint(int32(24))%32) + v519
							v650 = v644
							v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+6)))
							v655 = v651<<(uint(int32(16))%32) + v649
							v656 = v650
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+5)))
							v661 = v657<<(uint(int32(8))%32) + v655
							v662 = v656
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						case 10:
							v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+10)))
							v634 = v630<<(uint(int32(24))%32) + v520
							v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+9)))
							v639 = v635<<(uint(int32(16))%32) + v634
							v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+8)))
							v644 = v640<<(uint(int32(8))%32) + v639
							v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+7)))
							v649 = v645<<(uint(int32(24))%32) + v519
							v650 = v644
							v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+6)))
							v655 = v651<<(uint(int32(16))%32) + v649
							v656 = v650
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+5)))
							v661 = v657<<(uint(int32(8))%32) + v655
							v662 = v656
							v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+4)))
							v665 = v661 + v663
							v666 = v662
							v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
							v671 = v667<<(uint(int32(24))%32) + v518
							v672 = v665
							v673 = v666
							v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
							v678 = v674<<(uint(int32(16))%32) + v671
							v679 = v672
							v680 = v673
							v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
							v685 = v681<<(uint(int32(8))%32) + v678
							v686 = v679
							v687 = v680
							v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
							v693 = v685 + v688
							v694 = v686
							v695 = v687
						default:
							v693 = v518
							v694 = v519
							v695 = v520
						}
					} else {
						if base.Ui32(int32(12)) <= base.Ui32(v413) {
							v526 = v412
							v527 = v413
							v529 = v420
							v530 = v420
							v531 = v420
							for {
								v533 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
								v534 = v533 + v530
								v535 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
								v537 = *(*int32)(unsafe.Add(mBase, uint32(v526)+8))
								v538 = v537 + v531
								v540 = int32(4)
								v542 = v535 + v529 - v538 ^ base.I32_rotl(v538, v540)
								v546 = v534 - v542 ^ base.I32_rotl(v542, int32(6))
								v547 = v538 + v534
								v548 = v542 + v547
								v549 = v546 + v548
								v553 = v547 - v546 ^ base.I32_rotl(v546, int32(8))
								v557 = v548 - v553 ^ base.I32_rotl(v553, int32(16))
								v561 = v549 - v557 ^ base.I32_rotl(v557, int32(19))
								v562 = v553 + v549
								v563 = v557 + v562
								v564 = v561 + v563
								v568 = v562 - v561 ^ base.I32_rotl(v561, v540)
								v569 = int32(12)
								v570 = v526 + v569
								v572 = v527 - v569
								if base.Ui32(int32(11)) < base.Ui32(v572) {
									v526 = v570
									v527 = v572
									v529 = v563
									v530 = v564
									v531 = v568
									continue
								} else {
									break
								}
								break
							}
							v575 = v570
							v576 = v572
							v578 = v563
							v579 = v564
							v580 = v568
						} else {
							v575 = v412
							v576 = v413
							v578 = v420
							v579 = v420
							v580 = v420
						}
						switch v576 - int32(1) {
						case 0:
							v627 = v578
							v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
							v693 = v627 + v628
							v694 = v579
							v695 = v580
						case 1:
							v622 = v578
							v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+1)))
							v627 = v623<<(uint(int32(8))%32) + v622
							v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
							v693 = v627 + v628
							v694 = v579
							v695 = v580
						case 2:
							v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+2)))
							v622 = v618<<(uint(int32(16))%32) + v578
							v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+1)))
							v627 = v623<<(uint(int32(8))%32) + v622
							v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
							v693 = v627 + v628
							v694 = v579
							v695 = v580
						case 3:
							v615 = v579
							v616 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v693 = v616 + v578
							v694 = v615
							v695 = v580
						case 4:
							v612 = v579
							v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+4)))
							v615 = v612 + v613
							v616 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v693 = v616 + v578
							v694 = v615
							v695 = v580
						case 5:
							v607 = v579
							v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+5)))
							v612 = v608<<(uint(int32(8))%32) + v607
							v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+4)))
							v615 = v612 + v613
							v616 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v693 = v616 + v578
							v694 = v615
							v695 = v580
						case 6:
							v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+6)))
							v607 = v603<<(uint(int32(16))%32) + v579
							v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+5)))
							v612 = v608<<(uint(int32(8))%32) + v607
							v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+4)))
							v615 = v612 + v613
							v616 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v693 = v616 + v578
							v694 = v615
							v695 = v580
						case 7:
							v598 = v580
							v599 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v601 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
							v693 = v599 + v578
							v694 = v601 + v579
							v695 = v598
						case 8:
							v593 = v580
							v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+8)))
							v598 = v594<<(uint(int32(8))%32) + v593
							v599 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v601 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
							v693 = v599 + v578
							v694 = v601 + v579
							v695 = v598
						case 9:
							v588 = v580
							v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+9)))
							v593 = v589<<(uint(int32(16))%32) + v588
							v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+8)))
							v598 = v594<<(uint(int32(8))%32) + v593
							v599 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v601 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
							v693 = v599 + v578
							v694 = v601 + v579
							v695 = v598
						case 10:
							v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+10)))
							v588 = v584<<(uint(int32(24))%32) + v580
							v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+9)))
							v593 = v589<<(uint(int32(16))%32) + v588
							v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+8)))
							v598 = v594<<(uint(int32(8))%32) + v593
							v599 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
							v601 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
							v693 = v599 + v578
							v694 = v601 + v579
							v695 = v598
						default:
							v693 = v578
							v694 = v579
							v695 = v580
						}
					}
					v698 = int32(14)
					v700 = v694 ^ v695 - base.I32_rotl(v694, v698)
					v704 = v700 ^ v693 - base.I32_rotl(v700, int32(11))
					v708 = v704 ^ v694 - base.I32_rotl(v704, int32(25))
					v712 = v708 ^ v700 - base.I32_rotl(v708, int32(16))
					v716 = v712 ^ v704 - base.I32_rotl(v712, int32(4))
					v720 = v716 ^ v708 - base.I32_rotl(v716, v698)
					v729 = base.I64_extend_i32_u(v720)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v720^v712-base.I32_rotl(v720, int32(24)))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v729
					if v729 == int64(0) {
						v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v735 != 0 {
							v736 = int64(2)
						} else {
							v736 = int64(1)
						}
						*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v736
					} else {
					}
					return v10
				}
			}
		}
	}
}
func F__jumbleAlterUserMappingStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 != 0 {
			v8 = F_strlen(m, v7)
			mBase = m.M
			v10 = v8 + int32(1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v73 = v19
			} else {
				v20 = int32(4)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v22-int32(1021)) < base.Ui32(v20) {
					v32 = v22
					v34 = v20
					v36 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v32) {
							v41 = F_hash_bytes_extended(m, v21, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v21))) = v41
							v44 = int32(8)
						} else {
							v44 = v32
						}
						v46 = int32(1024) - v44
						if base.Ui32(v34) < base.Ui32(v46) {
							v48 = v34
						} else {
							v48 = v46
						}
						if v48 != 0 {
							base.MemoryCopy(m, v44+v21, v36, v48)
						} else {
						}
						v52 = v44 + v48
						v53 = v34 - v48
						if v53 != 0 {
							v32 = v52
							v34 = v53
							v36 = v48 + v36
							continue
						} else {
							break
						}
						break
					}
					v62 = v52
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v21))) = v16
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v62 = v56 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
				v73 = v62
			}
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(int32(1024)-v73) < base.Ui32(v10) {
				v83 = v7
				v84 = v10
				v85 = v73
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v85) {
						v94 = F_hash_bytes_extended(m, v78, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v78))) = v94
						v97 = int32(8)
					} else {
						v97 = v85
					}
					v99 = int32(1024) - v97
					if base.Ui32(v84) < base.Ui32(v99) {
						v101 = v84
					} else {
						v101 = v99
					}
					if v101 != 0 {
						base.MemoryCopy(m, v97+v78, v83, v101)
					} else {
					}
					v105 = v97 + v101
					v106 = v84 - v101
					if v106 != 0 {
						v83 = v83 + v101
						v84 = v106
						v85 = v105
						continue
					} else {
						break
					}
					break
				}
				v114 = v105
			} else {
				if v10 != 0 {
					base.MemoryCopy(m, v73+v78, v7, v10)
				} else {
				}
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v114 = v109 + v10
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
		} else {
			v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v120 + int32(1)
		}
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		F__jumbleNode(m, l0, v124)
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleArrayCoerceExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
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
	var v109 int32
	_ = v109
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = l1 + int32(12)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v70 = v19
			} else {
				v20 = int32(4)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v22-int32(1021)) < base.Ui32(v20) {
					v31 = v22
					v33 = v20
					v35 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v31) {
							v40 = F_hash_bytes_extended(m, v21, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v21))) = v40
							v43 = int32(8)
						} else {
							v43 = v31
						}
						v45 = int32(1024) - v43
						if base.Ui32(v33) < base.Ui32(v45) {
							v47 = v33
						} else {
							v47 = v45
						}
						if v47 != 0 {
							base.MemoryCopy(m, v43+v21, v35, v47)
						} else {
						}
						v51 = v43 + v47
						v52 = v33 - v47
						if v52 != 0 {
							v31 = v51
							v33 = v52
							v35 = v47 + v35
							continue
						} else {
							break
						}
						break
					}
					v60 = v51
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v21))) = v16
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v60 = v55 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
				v70 = v60
			}
			v75 = int32(4)
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(v70-int32(1021)) < base.Ui32(v75) {
				v82 = v10
				v83 = v70
				v85 = v75
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v83) {
						v92 = F_hash_bytes_extended(m, v76, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v92
						v95 = int32(8)
					} else {
						v95 = v83
					}
					v97 = int32(1024) - v95
					if base.Ui32(v85) < base.Ui32(v97) {
						v99 = v85
					} else {
						v99 = v97
					}
					if v99 != 0 {
						base.MemoryCopy(m, v95+v76, v82, v99)
					} else {
					}
					v103 = v95 + v99
					v104 = v85 - v99
					if v104 != 0 {
						v82 = v82 + v99
						v83 = v103
						v85 = v104
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
			} else {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				*(*int32)(unsafe.Add(mBase, uint32(v70+v76))) = v107
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109 + int32(4)
			}
			return
		}
	}
}
func F__jumbleBoolExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v4 = l1 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = v13
	} else {
		v14 = int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v16-int32(1021)) < base.Ui32(v14) {
			v25 = v16
			v27 = v14
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v34 = F_hash_bytes_extended(m, v15, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v34
					v37 = int32(8)
				} else {
					v37 = v25
				}
				v39 = int32(1024) - v37
				if base.Ui32(v27) < base.Ui32(v39) {
					v41 = v27
				} else {
					v41 = v39
				}
				if v41 != 0 {
					base.MemoryCopy(m, v37+v15, v29, v41)
				} else {
				}
				v45 = v37 + v41
				v46 = v27 - v41
				if v46 != 0 {
					v25 = v45
					v27 = v46
					v29 = v41 + v29
					continue
				} else {
					break
				}
				break
			}
			v54 = v45
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v15))) = v10
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v49 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
		v64 = v54
	}
	v69 = int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v64-int32(1021)) < base.Ui32(v69) {
		v76 = v4
		v77 = v64
		v79 = v69
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v77) {
				v86 = F_hash_bytes_extended(m, v70, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v70))) = v86
				v89 = int32(8)
			} else {
				v89 = v77
			}
			v91 = int32(1024) - v89
			if base.Ui32(v79) < base.Ui32(v91) {
				v93 = v79
			} else {
				v93 = v91
			}
			if v93 != 0 {
				base.MemoryCopy(m, v89+v70, v76, v93)
			} else {
			}
			v97 = v89 + v93
			v98 = v79 - v93
			if v98 != 0 {
				v76 = v76 + v93
				v77 = v97
				v79 = v98
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(v64+v70))) = v101
		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(4)
	}
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		return
	} else {
		return
	}
}
func F__jumbleCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int64
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		v7 = v5 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v29 = v19
				v31 = v17
				v33 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v29) {
						v38 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v38
						v41 = int32(8)
					} else {
						v41 = v29
					}
					v43 = int32(1024) - v41
					if base.Ui32(v31) < base.Ui32(v43) {
						v45 = v31
					} else {
						v45 = v43
					}
					if v45 != 0 {
						base.MemoryCopy(m, v41+v18, v33, v45)
					} else {
					}
					v49 = v41 + v45
					v50 = v31 - v45
					if v50 != 0 {
						v29 = v49
						v31 = v50
						v33 = v45 + v33
						continue
					} else {
						break
					}
					break
				}
				v59 = v49
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v59 = v53 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
			v70 = v59
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v70) < base.Ui32(v7) {
			v80 = v4
			v81 = v7
			v82 = v70
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v82) {
					v91 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v91
					v94 = int32(8)
				} else {
					v94 = v82
				}
				v96 = int32(1024) - v94
				if base.Ui32(v81) < base.Ui32(v96) {
					v98 = v81
				} else {
					v98 = v96
				}
				if v98 != 0 {
					base.MemoryCopy(m, v94+v75, v80, v98)
				} else {
				}
				v102 = v94 + v98
				v103 = v81 - v98
				if v103 != 0 {
					v80 = v80 + v98
					v81 = v103
					v82 = v102
					continue
				} else {
					break
				}
				break
			}
			v111 = v102
		} else {
			if v7 != 0 {
				base.MemoryCopy(m, v70+v75, v4, v7)
			} else {
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v111 = v106 + v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v117 + int32(1)
	}
	v122 = l1 + int32(8)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v128 == int32(0) {
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v182 = v131
	} else {
		v132 = int32(4)
		v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v134-int32(1021)) < base.Ui32(v132) {
			v143 = v134
			v144 = v132
			v147 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v143) {
					v152 = F_hash_bytes_extended(m, v133, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v133))) = v152
					v155 = int32(8)
				} else {
					v155 = v143
				}
				v157 = int32(1024) - v155
				if base.Ui32(v144) < base.Ui32(v157) {
					v159 = v144
				} else {
					v159 = v157
				}
				if v159 != 0 {
					base.MemoryCopy(m, v155+v133, v147, v159)
				} else {
				}
				v163 = v155 + v159
				v164 = v144 - v159
				if v164 != 0 {
					v143 = v163
					v144 = v164
					v147 = v159 + v147
					continue
				} else {
					break
				}
				break
			}
			v172 = v163
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v134+v133))) = v128
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v172 = v167 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
		v182 = v172
	}
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v182 != int32(1024) {
		v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
		*(*uint8)(unsafe.Add(mBase, uint32(v182+v187))) = uint8(v191)
		v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v193 + int32(1)
	} else {
		v199 = F_hash_bytes_extended(m, v187, int32(1024), int64(0))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v187))) = v199
		v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
		*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)) = uint8(v201)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
	}
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F__jumbleNode(m, l0, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		return
	} else {
		return
	}
}
func F__jumbleFieldStore(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleFuncExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v4 = l1 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = v13
	} else {
		v14 = int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v16-int32(1021)) < base.Ui32(v14) {
			v25 = v16
			v27 = v14
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v34 = F_hash_bytes_extended(m, v15, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v34
					v37 = int32(8)
				} else {
					v37 = v25
				}
				v39 = int32(1024) - v37
				if base.Ui32(v27) < base.Ui32(v39) {
					v41 = v27
				} else {
					v41 = v39
				}
				if v41 != 0 {
					base.MemoryCopy(m, v37+v15, v29, v41)
				} else {
				}
				v45 = v37 + v41
				v46 = v27 - v41
				if v46 != 0 {
					v25 = v45
					v27 = v46
					v29 = v41 + v29
					continue
				} else {
					break
				}
				break
			}
			v54 = v45
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v15))) = v10
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v49 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
		v64 = v54
	}
	v69 = int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v64-int32(1021)) < base.Ui32(v69) {
		v76 = v4
		v77 = v64
		v79 = v69
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v77) {
				v86 = F_hash_bytes_extended(m, v70, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v70))) = v86
				v89 = int32(8)
			} else {
				v89 = v77
			}
			v91 = int32(1024) - v89
			if base.Ui32(v79) < base.Ui32(v91) {
				v93 = v79
			} else {
				v93 = v91
			}
			if v93 != 0 {
				base.MemoryCopy(m, v89+v70, v76, v93)
			} else {
			}
			v97 = v89 + v93
			v98 = v79 - v93
			if v98 != 0 {
				v76 = v76 + v93
				v77 = v97
				v79 = v98
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(v64+v70))) = v101
		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(4)
	}
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F__jumbleNode(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		return
	} else {
		return
	}
}
func F__jumbleJsonObjectConstructor(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = l1 + int32(12)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v70 = v19
			} else {
				v20 = int32(4)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v22-int32(1021)) < base.Ui32(v20) {
					v31 = v22
					v32 = v20
					v35 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v31) {
							v40 = F_hash_bytes_extended(m, v21, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v21))) = v40
							v43 = int32(8)
						} else {
							v43 = v31
						}
						v45 = int32(1024) - v43
						if base.Ui32(v32) < base.Ui32(v45) {
							v47 = v32
						} else {
							v47 = v45
						}
						if v47 != 0 {
							base.MemoryCopy(m, v43+v21, v35, v47)
						} else {
						}
						v51 = v43 + v47
						v52 = v32 - v47
						if v52 != 0 {
							v31 = v51
							v32 = v52
							v35 = v47 + v35
							continue
						} else {
							break
						}
						break
					}
					v60 = v51
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v21))) = v16
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v60 = v55 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
				v70 = v60
			}
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v70 != int32(1024) {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				*(*uint8)(unsafe.Add(mBase, uint32(v70+v75))) = uint8(v79)
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 + int32(1)
			} else {
				v87 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v75))) = v87
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)) = uint8(v89)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
			}
			v94 = l1 + int32(13)
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v100 == int32(0) {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v154 = v103
			} else {
				v104 = int32(4)
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v106-int32(1021)) < base.Ui32(v104) {
					v115 = v106
					v116 = v104
					v119 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v115) {
							v124 = F_hash_bytes_extended(m, v105, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v105))) = v124
							v127 = int32(8)
						} else {
							v127 = v115
						}
						v129 = int32(1024) - v127
						if base.Ui32(v116) < base.Ui32(v129) {
							v131 = v116
						} else {
							v131 = v129
						}
						if v131 != 0 {
							base.MemoryCopy(m, v127+v105, v119, v131)
						} else {
						}
						v135 = v127 + v131
						v136 = v116 - v131
						if v136 != 0 {
							v115 = v135
							v116 = v136
							v119 = v131 + v119
							continue
						} else {
							break
						}
						break
					}
					v144 = v135
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v106+v105))) = v100
					v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v144 = v139 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v144
				v154 = v144
			}
			v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v154 != int32(1024) {
				v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
				*(*uint8)(unsafe.Add(mBase, uint32(v154+v159))) = uint8(v163)
				v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165 + int32(1)
			} else {
				v171 = F_hash_bytes_extended(m, v159, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v159))) = v171
				v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
				*(*uint8)(unsafe.Add(mBase, uint32(v159)+8)) = uint8(v173)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
			}
			return
		}
	}
}
func F__jumbleWithClause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v7 = l1 + int32(8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v67 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v28 = v19
				v29 = v17
				v32 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v28) {
						v37 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v37
						v40 = int32(8)
					} else {
						v40 = v28
					}
					v42 = int32(1024) - v40
					if base.Ui32(v29) < base.Ui32(v42) {
						v44 = v29
					} else {
						v44 = v42
					}
					if v44 != 0 {
						base.MemoryCopy(m, v40+v18, v32, v44)
					} else {
					}
					v48 = v40 + v44
					v49 = v29 - v44
					if v49 != 0 {
						v28 = v48
						v29 = v49
						v32 = v44 + v32
						continue
					} else {
						break
					}
					break
				}
				v57 = v48
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v57 = v52 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
			v67 = v57
		}
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v67 != int32(1024) {
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			*(*uint8)(unsafe.Add(mBase, uint32(v67+v72))) = uint8(v76)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v78 + int32(1)
		} else {
			v84 = F_hash_bytes_extended(m, v72, int32(1024), int64(0))
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v72))) = v84
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)) = uint8(v86)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
		}
		return
	}
}
