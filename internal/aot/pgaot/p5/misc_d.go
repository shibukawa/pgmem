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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	v2 = l1
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1049]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1050]))
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
	v127 = v14
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
	v118 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, _consts[1050])) = v118
	v127 = v118
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
	v34 = v3
	goto L9
L7:
	;
	v73 = v3
	goto L8
L8:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v37 = v28 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1051])))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+2032))
	v42 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+2032)) = v41 >> (uint(v42) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1052])))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+2032)) = v48 >> (uint(v42) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1053])))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+2032)) = v55 >> (uint(v42) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[1054])))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+2032)) = v62 >> (uint(v42) % 32)
	v66 = int32(4)
	v67 = v28 + v66
	v69 = v34 + v66
	if v69 != v12&int32(2147483644) {
		v28 = v67
		v34 = v69
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v73 = v67
	goto L8
L11:
	;
	goto L10
L12:
	;
	v85 = v73
	v89 = int32(0)
	goto L13
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85<<(uint(int32(2))%32))+uint32(_consts[1051])))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+2032))
	v99 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+2032)) = v98 >> (uint(v99) % 32)
	v105 = v89 + v99
	if v105 != v20 {
		v85 = v85 + v99
		v89 = v105
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
		goto L134
	} else {
		goto L135
	}
L17:
	;
	v518 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+2029)) = uint8(v518)
	v521 = v510 + int32(1872)
	goto L105
L18:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	v364 = F_MemoryContextAllocZero(m, v362, int32(2036))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L68
	} else {
		goto L69
	}
L19:
	;
	v136 = int32(0)
	goto L21
L20:
	;
	v347 = v127 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1050])) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v148)+2032)) = v347
	return v148
L21:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v136<<(uint(int32(2))%32))+uint32(_consts[1051])))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2029)))
	if v149 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if int32(2147483646) <= v127 {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	v182 = v136 + int32(1)
	if v182 != v12 {
		v136 = v182
		goto L21
	} else {
		goto L35
	}
L24:
	;
	v153 = v148 + int32(1872)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 == int32(0) {
		v176 = v156
		v177 = v157
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v177-v176 != 0 {
		goto L23
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	if v156 != v157 {
		v176 = v156
		v177 = v157
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v161 = v153
	v162 = l0
	goto L29
L29:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v166 == int32(0) {
		v176 = v165
		v177 = v166
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v176 = v165
	v177 = v166
	goto L26
L31:
	;
	v169 = int32(1)
	if v165 == v166 {
		v161 = v161 + v169
		v162 = v162 + v169
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2028)))
	if v179 == v2 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	goto L23
L35:
	;
	goto L22
L36:
	;
	v187 = v12 & int32(3)
	v188 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v12) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v12 < int32(20) {
		goto L18
	} else {
		goto L51
	}
L39:
	;
	v197 = v188
	v203 = int32(0)
	goto L42
L40:
	;
	v242 = v188
	goto L41
L41:
	;
	if v187 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v206 = v197 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_consts[1051])))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+2032))
	v211 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+2032)) = v210 >> (uint(v211) % 32)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_consts[1052])))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+2032)) = v217 >> (uint(v211) % 32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_consts[1053])))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+2032)) = v224 >> (uint(v211) % 32)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_consts[1054])))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+2032))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+2032)) = v231 >> (uint(v211) % 32)
	v235 = int32(4)
	v236 = v197 + v235
	v238 = v203 + v235
	if v238 != v12&int32(2147483644) {
		v197 = v236
		v203 = v238
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v242 = v236
	goto L41
L44:
	;
	goto L43
L45:
	;
	v252 = v242
	v256 = v188
	goto L48
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1050])) = int32(1073741823)
	goto L38
L48:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v252<<(uint(int32(2))%32))+uint32(_consts[1051])))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+2032))
	v266 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+2032)) = v265 >> (uint(v266) % 32)
	v272 = v256 + v266
	if v272 != v187 {
		v252 = v252 + v266
		v256 = v272
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
L51:
	;
	v299 = int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, _consts[1051]))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+2029)))
	if v302 != v299 {
		v510 = v301
		goto L17
	} else {
		goto L52
	}
L52:
	;
	v307 = v301
	v308 = v299
	goto L53
L53:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v308<<(uint(int32(2))%32))+uint32(_consts[1051])))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+2029)))
	if v320 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v510 = v319
	goto L17
L56:
	;
	goto L57
L57:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+2032))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v307)+2032))
	if v323 < v324 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v326 = v319
	goto L60
L59:
	;
	v326 = v307
	goto L60
L60:
	;
	v328 = v308 + int32(1)
	if v328 == int32(20) {
		v510 = v326
		goto L17
	} else {
		goto L61
	}
L61:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v328<<(uint(int32(2))%32))+uint32(_consts[1051])))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2029)))
	if v336 != int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v510 = v335
	goto L17
L63:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+2032))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v326)+2032))
	if v339 < v340 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v342 = v335
	goto L67
L66:
	;
	v342 = v326
	goto L67
L67:
	;
	v307 = v342
	v308 = v308 + int32(2)
	goto L53
L68:
	;
	return int32(0)
L69:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _consts[1049]))
	*(*int32)(unsafe.Add(mBase, uint32(v369<<(uint(int32(2))%32))+uint32(_consts[1051]))) = v364
	v375 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+2029)) = uint8(v375)
	v378 = v364 + int32(1872)
	goto L73
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+2028)) = uint8(v2)
	v495 = int32(4533280)
	v497 = *(*int32)(unsafe.Add(mBase, _consts[1050]))
	v498 = int32(1)
	v499 = v497 + v498
	*(*int32)(unsafe.Add(mBase, _consts[1050])) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v364)+2032)) = v499
	v502 = int32(4533184)
	v504 = *(*int32)(unsafe.Add(mBase, _consts[1049]))
	*(*int32)(unsafe.Add(mBase, _consts[1049])) = v504 + v498
	v646 = v364
	goto L16
L71:
	;
	v491 = F_strlen(m, v480)
	mBase = m.M
	goto L70
L73:
	;
	goto L74
L74:
	;
	v385 = int32(155)
	if (v378^l0)&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v484 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v481))) = uint8(v484)
	goto L71
L76:
	;
	v465 = v460
	v466 = v461
	v467 = v462
	goto L98
L77:
	;
	if v455 == int32(0) {
		v480 = v453
		v481 = v454
		goto L75
	} else {
		goto L97
	}
L78:
	;
	v453 = l0
	v454 = v378
	v455 = v385
	goto L77
L79:
	;
	goto L80
L80:
	;
	if l0&int32(3) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v422 == int32(0) {
		v480 = v419
		v481 = v420
		goto L75
	} else {
		goto L90
	}
L82:
	;
	v419 = l0
	v420 = v378
	v421 = v385
	v422 = int32(1)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v398 = l0
	v399 = v378
	v400 = v385
	goto L85
L85:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v402)
	if v402 == int32(0) {
		v460 = v398
		v461 = v399
		v462 = v400
		goto L76
	} else {
		goto L87
	}
L86:
	;
	v419 = v413
	v420 = v407
	v421 = v409
	v422 = v411
	goto L81
L87:
	;
	v406 = int32(1)
	v407 = v399 + v406
	v409 = v400 - v406
	v410 = int32(0)
	v411 = base.B2i32(v409 != v410)
	v413 = v398 + v406
	if v413&int32(3) == v410 {
		v419 = v413
		v420 = v407
		v421 = v409
		v422 = v411
		goto L81
	} else {
		goto L88
	}
L88:
	;
	if v409 != 0 {
		v398 = v413
		v399 = v407
		v400 = v409
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	if v425 == int32(0) {
		v453 = v419
		v454 = v420
		v455 = v421
		goto L77
	} else {
		goto L91
	}
L91:
	;
	if base.Ui32(v421) < base.Ui32(int32(4)) {
		v453 = v419
		v454 = v420
		v455 = v421
		goto L77
	} else {
		goto L92
	}
L92:
	;
	v431 = v419
	v432 = v420
	v433 = v421
	goto L93
L93:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v439 = int32(-2139062144)
	if (int32(16843008)-v436|v436)&v439 != v439 {
		v460 = v431
		v461 = v432
		v462 = v433
		goto L76
	} else {
		goto L95
	}
L94:
	;
	v453 = v447
	v454 = v445
	v455 = v449
	goto L77
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432))) = v436
	v444 = int32(4)
	v445 = v432 + v444
	v447 = v431 + v444
	v449 = v433 - v444
	if base.Ui32(int32(3)) < base.Ui32(v449) {
		v431 = v447
		v432 = v445
		v433 = v449
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v460 = v453
	v461 = v454
	v462 = v455
	goto L76
L98:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	*(*uint8)(unsafe.Add(mBase, uint32(v466))) = uint8(v469)
	if v469 == int32(0) {
		v480 = v465
		v481 = v466
		goto L75
	} else {
		goto L100
	}
L99:
	;
	v480 = v476
	v481 = v474
	goto L75
L100:
	;
	v473 = int32(1)
	v474 = v466 + v473
	v476 = v465 + v473
	v478 = v467 - v473
	if v478 != 0 {
		v465 = v476
		v466 = v474
		v467 = v478
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v637 = int32(4533280)
	v639 = *(*int32)(unsafe.Add(mBase, _consts[1050]))
	v641 = v639 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1050])) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v510)+2032)) = v641
	v646 = v510
	goto L16
L103:
	;
	v634 = F_strlen(m, v623)
	mBase = m.M
	goto L102
L105:
	;
	goto L106
L106:
	;
	v528 = int32(155)
	if (v521^l0)&int32(3) != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v627 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v627)
	goto L103
L108:
	;
	v608 = v603
	v609 = v604
	v610 = v605
	goto L130
L109:
	;
	if v598 == int32(0) {
		v623 = v596
		v624 = v597
		goto L107
	} else {
		goto L129
	}
L110:
	;
	v596 = l0
	v597 = v521
	v598 = v528
	goto L109
L111:
	;
	goto L112
L112:
	;
	if l0&int32(3) == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if v565 == int32(0) {
		v623 = v562
		v624 = v563
		goto L107
	} else {
		goto L122
	}
L114:
	;
	v562 = l0
	v563 = v521
	v564 = v528
	v565 = int32(1)
	goto L113
L115:
	;
	goto L116
L116:
	;
	v541 = l0
	v542 = v521
	v543 = v528
	goto L117
L117:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	*(*uint8)(unsafe.Add(mBase, uint32(v542))) = uint8(v545)
	if v545 == int32(0) {
		v603 = v541
		v604 = v542
		v605 = v543
		goto L108
	} else {
		goto L119
	}
L118:
	;
	v562 = v556
	v563 = v550
	v564 = v552
	v565 = v554
	goto L113
L119:
	;
	v549 = int32(1)
	v550 = v542 + v549
	v552 = v543 - v549
	v553 = int32(0)
	v554 = base.B2i32(v552 != v553)
	v556 = v541 + v549
	if v556&int32(3) == v553 {
		v562 = v556
		v563 = v550
		v564 = v552
		v565 = v554
		goto L113
	} else {
		goto L120
	}
L120:
	;
	if v552 != 0 {
		v541 = v556
		v542 = v550
		v543 = v552
		goto L117
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	if v568 == int32(0) {
		v596 = v562
		v597 = v563
		v598 = v564
		goto L109
	} else {
		goto L123
	}
L123:
	;
	if base.Ui32(v564) < base.Ui32(int32(4)) {
		v596 = v562
		v597 = v563
		v598 = v564
		goto L109
	} else {
		goto L124
	}
L124:
	;
	v574 = v562
	v575 = v563
	v576 = v564
	goto L125
L125:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v582 = int32(-2139062144)
	if (int32(16843008)-v579|v579)&v582 != v582 {
		v603 = v574
		v604 = v575
		v605 = v576
		goto L108
	} else {
		goto L127
	}
L126:
	;
	v596 = v590
	v597 = v588
	v598 = v592
	goto L109
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = v579
	v587 = int32(4)
	v588 = v575 + v587
	v590 = v574 + v587
	v592 = v576 - v587
	if base.Ui32(int32(3)) < base.Ui32(v592) {
		v574 = v590
		v575 = v588
		v576 = v592
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v603 = v596
	v604 = v597
	v605 = v598
	goto L108
L130:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v612)
	if v612 == int32(0) {
		v623 = v608
		v624 = v609
		goto L107
	} else {
		goto L132
	}
L131:
	;
	v623 = v619
	v624 = v617
	goto L107
L132:
	;
	v616 = int32(1)
	v617 = v609 + v616
	v619 = v608 + v616
	v621 = v610 - v616
	if v621 != 0 {
		v608 = v619
		v609 = v617
		v610 = v621
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v659 = int32(5)
	goto L136
L135:
	;
	v659 = int32(1)
	goto L136
L136:
	;
	F_parse_format(m, v646, l0, int32(1690976), int32(1690240), int32(1690368), v659, int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L68
	} else {
		goto L137
	}
L137:
	;
	v663 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v646)+2029)) = uint8(v663)
	return v646
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
	var v92 int32
	_ = v92
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
	var v112 int32
	_ = v112
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
	var v189 int32
	_ = v189
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
	var v213 float64
	_ = v213
	var v216 float64
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
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
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
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
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
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
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(-1)
	v17 = int32(684283)
	v21 = m.G0
	v23 = v21 - int32(32)
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v24
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1029])))
	if v32 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v671
L2:
	;
	if v100 != l0 {
		v671 = v16
		goto L1
	} else {
		goto L23
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
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1030])))
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
		v92 = l1
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v100 = v92 - l1
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
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v92 = v88
	goto L15
L19:
	;
	v92 = v71
	goto L15
L20:
	;
	goto L21
L21:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v88 = v71 + int32(1)
	if v86 != 0 {
		v71 = v88
		v72 = v86
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v102 = int32(46)
	v103 = F___strchrnul(m, l1, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v105 == v102 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v399 = int32(31744)
	if l2&v399 == v399 {
		v671 = v16
		goto L1
	} else {
		goto L116
	}
L25:
	;
	if v109 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v109 = v103
	goto L28
L27:
	;
	v109 = int32(0)
	goto L28
L28:
	;
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v109
	v112 = v109 + int32(1)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v113 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if l0 < int32(6) {
		v397 = l0
		goto L24
	} else {
		goto L65
	}
L32:
	;
	v114 = int32(573635)
	v118 = m.G0
	v120 = v118 - int32(32)
	v121 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v120)+24)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v120)+16)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v120)+8)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v120))) = v121
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v129 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v213 = float64(0)
	goto L34
L34:
	;
	v216 = base.F64_nearest(base.F64_mul(v213, float64(1e+06)))
	if base.F64_lt(base.F64_abs(v216), float64(2.147483648e+09)) != 0 {
		goto L62
	} else {
		goto L63
	}
L35:
	;
	v198 = F_strlen(m, v112)
	mBase = m.M
	if v197 != v198 {
		v671 = v16
		goto L1
	} else {
		goto L56
	}
L36:
	;
	v197 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v133 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v137 = v112
	goto L42
L40:
	;
	goto L41
L41:
	;
	v147 = v114
	v148 = v129
	goto L45
L42:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v143 == v129 {
		v137 = v137 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v197 = v137 - v112
	goto L35
L44:
	;
	goto L43
L45:
	;
	v155 = v120 + int32(base.Ui32(v148)>>(uint(int32(3))%32))&int32(28)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156 | v157<<(uint(v148)%32)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v161 != 0 {
		v147 = v147 + v157
		v148 = v161
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v164 == int32(0) {
		v189 = v112
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v197 = v189 - v112
	goto L35
L49:
	;
	v168 = v112
	v169 = v164
	goto L50
L50:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(base.Ui32(v169)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v177)>>(uint(v169)%32))&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v189 = v185
	goto L48
L52:
	;
	v189 = v168
	goto L48
L53:
	;
	goto L54
L54:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v185 = v168 + int32(1)
	if v183 != 0 {
		v168 = v185
		v169 = v183
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(0)
	v205 = F_strtod(m, v109, v14+int32(12))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v210 != 0 {
		v671 = v16
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v212 != 0 {
		v671 = v16
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v213 = v205
	goto L34
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v222
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v224)
	v226 = F_strlen(m, l1)
	mBase = m.M
	v397 = v226
	goto L24
L62:
	;
	v220 = base.I32_trunc_f64_s(v216)
	v222 = v220
	goto L61
L63:
	;
	goto L64
L64:
	;
	v222 = int32(-2147483648)
	goto L61
L65:
	;
	v229 = int32(14)
	if l2&v229 == v229 {
		v397 = l0
		goto L24
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(14)
	v237 = l0 + l1 - int32(2)
	v241 = v237
	goto L68
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v285
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v287)
	v290 = l0 - int32(4)
	v291 = l1 + v290
	v295 = v291
	goto L84
L68:
	;
	v246 = v241 + int32(1)
	v247 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241))))
	v248 = F___isspace(m, v247)
	mBase = m.M
	if v248 != 0 {
		v241 = v246
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v249 = int32(1)
	switch v247&int32(255) - int32(43) {
	case 0:
		v255 = v249
		goto L72
	default:
		v257 = v247
		v258 = v241
		v259 = v249
		goto L71
	case 2:
		goto L73
	}
L70:
	;
	goto L69
L71:
	;
	v260 = int32(0)
	v262 = v257 - int32(48)
	if base.Ui32(v262) <= base.Ui32(int32(9)) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v246))))
	v257 = v256
	v258 = v246
	v259 = v255
	goto L71
L73:
	;
	v255 = int32(0)
	goto L72
L74:
	;
	v265 = v260
	v266 = v262
	v267 = v258
	goto L77
L75:
	;
	v279 = v260
	goto L76
L76:
	;
	if v259 != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v269 = int32(10)
	v271 = v265*v269 - v266
	v272 = int32(*(*int8)(unsafe.Add(mBase, uint32(v267)+1)))
	v276 = v272 - int32(48)
	if base.Ui32(v276) < base.Ui32(v269) {
		v265 = v271
		v266 = v276
		v267 = v267 + int32(1)
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v279 = v271
	goto L76
L79:
	;
	goto L78
L80:
	;
	v285 = int32(0) - v279
	goto L82
L81:
	;
	v285 = v279
	goto L82
L82:
	;
	goto L67
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v339
	v341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v291))) = uint8(v341)
	v346 = l1
	goto L100
L84:
	;
	v300 = v295 + int32(1)
	v301 = int32(*(*int8)(unsafe.Add(mBase, uint32(v295))))
	v302 = F___isspace(m, v301)
	mBase = m.M
	if v302 != 0 {
		v295 = v300
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v303 = int32(1)
	switch v301&int32(255) - int32(43) {
	case 0:
		v309 = v303
		goto L88
	default:
		v311 = v301
		v312 = v295
		v313 = v303
		goto L87
	case 2:
		goto L89
	}
L86:
	;
	goto L85
L87:
	;
	v314 = int32(0)
	v316 = v311 - int32(48)
	if base.Ui32(v316) <= base.Ui32(int32(9)) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v310 = int32(*(*int8)(unsafe.Add(mBase, uint32(v300))))
	v311 = v310
	v312 = v300
	v313 = v309
	goto L87
L89:
	;
	v309 = int32(0)
	goto L88
L90:
	;
	v319 = v314
	v320 = v316
	v321 = v312
	goto L93
L91:
	;
	v333 = v314
	goto L92
L92:
	;
	if v313 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v323 = int32(10)
	v325 = v319*v323 - v320
	v326 = int32(*(*int8)(unsafe.Add(mBase, uint32(v321)+1)))
	v330 = v326 - int32(48)
	if base.Ui32(v330) < base.Ui32(v323) {
		v319 = v325
		v320 = v330
		v321 = v321 + int32(1)
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v333 = v325
	goto L92
L95:
	;
	goto L94
L96:
	;
	v339 = int32(0) - v333
	goto L98
L97:
	;
	v339 = v333
	goto L98
L98:
	;
	goto L83
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v390
	v392 = int32(2)
	if v290 != v392 {
		v671 = v392
		goto L1
	} else {
		goto L115
	}
L100:
	;
	v351 = v346 + int32(1)
	v352 = int32(*(*int8)(unsafe.Add(mBase, uint32(v346))))
	v353 = F___isspace(m, v352)
	mBase = m.M
	if v353 != 0 {
		v346 = v351
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v354 = int32(1)
	switch v352&int32(255) - int32(43) {
	case 0:
		v360 = v354
		goto L104
	default:
		v362 = v352
		v363 = v346
		v364 = v354
		goto L103
	case 2:
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	v365 = int32(0)
	v367 = v362 - int32(48)
	if base.Ui32(v367) <= base.Ui32(int32(9)) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v361 = int32(*(*int8)(unsafe.Add(mBase, uint32(v351))))
	v362 = v361
	v363 = v351
	v364 = v360
	goto L103
L105:
	;
	v360 = int32(0)
	goto L104
L106:
	;
	v370 = v365
	v371 = v367
	v372 = v363
	goto L109
L107:
	;
	v384 = v365
	goto L108
L108:
	;
	if v364 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v374 = int32(10)
	v376 = v370*v374 - v371
	v377 = int32(*(*int8)(unsafe.Add(mBase, uint32(v372)+1)))
	v381 = v377 - int32(48)
	if base.Ui32(v381) < base.Ui32(v374) {
		v370 = v376
		v371 = v381
		v372 = v372 + int32(1)
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v384 = v376
	goto L108
L111:
	;
	goto L110
L112:
	;
	v390 = int32(0) - v384
	goto L114
L113:
	;
	v390 = v384
	goto L114
L114:
	;
	goto L99
L115:
	;
	v395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v395)
	v671 = v392
	goto L1
L116:
	;
	switch v397 - int32(4) {
	case 0:
		goto L118
	default:
		v671 = v16
		goto L1
	case 2:
		goto L119
	}
L117:
	;
	v671 = int32(3)
	goto L1
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(31744)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v567 = l1 + int32(2)
	v571 = v567
	goto L169
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(31744)
	v408 = l1 + int32(4)
	v412 = v408
	goto L121
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v456
	v458 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v408))) = uint8(v458)
	v461 = l1 + int32(2)
	v465 = v461
	goto L137
L121:
	;
	v417 = v412 + int32(1)
	v418 = int32(*(*int8)(unsafe.Add(mBase, uint32(v412))))
	v419 = F___isspace(m, v418)
	mBase = m.M
	if v419 != 0 {
		v412 = v417
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v420 = int32(1)
	switch v418&int32(255) - int32(43) {
	case 0:
		v426 = v420
		goto L125
	default:
		v428 = v418
		v429 = v412
		v430 = v420
		goto L124
	case 2:
		goto L126
	}
L123:
	;
	goto L122
L124:
	;
	v431 = int32(0)
	v433 = v428 - int32(48)
	if base.Ui32(v433) <= base.Ui32(int32(9)) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v427 = int32(*(*int8)(unsafe.Add(mBase, uint32(v417))))
	v428 = v427
	v429 = v417
	v430 = v426
	goto L124
L126:
	;
	v426 = int32(0)
	goto L125
L127:
	;
	v436 = v431
	v437 = v433
	v438 = v429
	goto L130
L128:
	;
	v450 = v431
	goto L129
L129:
	;
	if v430 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v440 = int32(10)
	v442 = v436*v440 - v437
	v443 = int32(*(*int8)(unsafe.Add(mBase, uint32(v438)+1)))
	v447 = v443 - int32(48)
	if base.Ui32(v447) < base.Ui32(v440) {
		v436 = v442
		v437 = v447
		v438 = v438 + int32(1)
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v450 = v442
	goto L129
L132:
	;
	goto L131
L133:
	;
	v456 = int32(0) - v450
	goto L135
L134:
	;
	v456 = v450
	goto L135
L135:
	;
	goto L120
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v509
	v511 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v511)
	v516 = l1
	goto L153
L137:
	;
	v470 = v465 + int32(1)
	v471 = int32(*(*int8)(unsafe.Add(mBase, uint32(v465))))
	v472 = F___isspace(m, v471)
	mBase = m.M
	if v472 != 0 {
		v465 = v470
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v473 = int32(1)
	switch v471&int32(255) - int32(43) {
	case 0:
		v479 = v473
		goto L141
	default:
		v481 = v471
		v482 = v465
		v483 = v473
		goto L140
	case 2:
		goto L142
	}
L139:
	;
	goto L138
L140:
	;
	v484 = int32(0)
	v486 = v481 - int32(48)
	if base.Ui32(v486) <= base.Ui32(int32(9)) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v480 = int32(*(*int8)(unsafe.Add(mBase, uint32(v470))))
	v481 = v480
	v482 = v470
	v483 = v479
	goto L140
L142:
	;
	v479 = int32(0)
	goto L141
L143:
	;
	v489 = v484
	v490 = v486
	v491 = v482
	goto L146
L144:
	;
	v503 = v484
	goto L145
L145:
	;
	if v483 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v493 = int32(10)
	v495 = v489*v493 - v490
	v496 = int32(*(*int8)(unsafe.Add(mBase, uint32(v491)+1)))
	v500 = v496 - int32(48)
	if base.Ui32(v500) < base.Ui32(v493) {
		v489 = v495
		v490 = v500
		v491 = v491 + int32(1)
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v503 = v495
	goto L145
L148:
	;
	goto L147
L149:
	;
	v509 = int32(0) - v503
	goto L151
L150:
	;
	v509 = v503
	goto L151
L151:
	;
	goto L136
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v560
	goto L117
L153:
	;
	v521 = v516 + int32(1)
	v522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516))))
	v523 = F___isspace(m, v522)
	mBase = m.M
	if v523 != 0 {
		v516 = v521
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v524 = int32(1)
	switch v522&int32(255) - int32(43) {
	case 0:
		v530 = v524
		goto L157
	default:
		v532 = v522
		v533 = v516
		v534 = v524
		goto L156
	case 2:
		goto L158
	}
L155:
	;
	goto L154
L156:
	;
	v535 = int32(0)
	v537 = v532 - int32(48)
	if base.Ui32(v537) <= base.Ui32(int32(9)) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v521))))
	v532 = v531
	v533 = v521
	v534 = v530
	goto L156
L158:
	;
	v530 = int32(0)
	goto L157
L159:
	;
	v540 = v535
	v541 = v537
	v542 = v533
	goto L162
L160:
	;
	v554 = v535
	goto L161
L161:
	;
	if v534 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v544 = int32(10)
	v546 = v540*v544 - v541
	v547 = int32(*(*int8)(unsafe.Add(mBase, uint32(v542)+1)))
	v551 = v547 - int32(48)
	if base.Ui32(v551) < base.Ui32(v544) {
		v540 = v546
		v541 = v551
		v542 = v542 + int32(1)
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v554 = v546
	goto L161
L164:
	;
	goto L163
L165:
	;
	v560 = int32(0) - v554
	goto L167
L166:
	;
	v560 = v554
	goto L167
L167:
	;
	goto L152
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v615
	v617 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v567))) = uint8(v617)
	v622 = l1
	goto L185
L169:
	;
	v576 = v571 + int32(1)
	v577 = int32(*(*int8)(unsafe.Add(mBase, uint32(v571))))
	v578 = F___isspace(m, v577)
	mBase = m.M
	if v578 != 0 {
		v571 = v576
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v579 = int32(1)
	switch v577&int32(255) - int32(43) {
	case 0:
		v585 = v579
		goto L173
	default:
		v587 = v577
		v588 = v571
		v589 = v579
		goto L172
	case 2:
		goto L174
	}
L171:
	;
	goto L170
L172:
	;
	v590 = int32(0)
	v592 = v587 - int32(48)
	if base.Ui32(v592) <= base.Ui32(int32(9)) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v586 = int32(*(*int8)(unsafe.Add(mBase, uint32(v576))))
	v587 = v586
	v588 = v576
	v589 = v585
	goto L172
L174:
	;
	v585 = int32(0)
	goto L173
L175:
	;
	v595 = v590
	v596 = v592
	v597 = v588
	goto L178
L176:
	;
	v609 = v590
	goto L177
L177:
	;
	if v589 != 0 {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v599 = int32(10)
	v601 = v595*v599 - v596
	v602 = int32(*(*int8)(unsafe.Add(mBase, uint32(v597)+1)))
	v606 = v602 - int32(48)
	if base.Ui32(v606) < base.Ui32(v599) {
		v595 = v601
		v596 = v606
		v597 = v597 + int32(1)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v609 = v601
	goto L177
L180:
	;
	goto L179
L181:
	;
	v615 = int32(0) - v609
	goto L183
L182:
	;
	v615 = v609
	goto L183
L183:
	;
	goto L168
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v666
	goto L117
L185:
	;
	v627 = v622 + int32(1)
	v628 = int32(*(*int8)(unsafe.Add(mBase, uint32(v622))))
	v629 = F___isspace(m, v628)
	mBase = m.M
	if v629 != 0 {
		v622 = v627
		goto L185
	} else {
		goto L187
	}
L186:
	;
	v630 = int32(1)
	switch v628&int32(255) - int32(43) {
	case 0:
		v636 = v630
		goto L189
	default:
		v638 = v628
		v639 = v622
		v640 = v630
		goto L188
	case 2:
		goto L190
	}
L187:
	;
	goto L186
L188:
	;
	v641 = int32(0)
	v643 = v638 - int32(48)
	if base.Ui32(v643) <= base.Ui32(int32(9)) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	v637 = int32(*(*int8)(unsafe.Add(mBase, uint32(v627))))
	v638 = v637
	v639 = v627
	v640 = v636
	goto L188
L190:
	;
	v636 = int32(0)
	goto L189
L191:
	;
	v646 = v641
	v647 = v643
	v648 = v639
	goto L194
L192:
	;
	v660 = v641
	goto L193
L193:
	;
	if v640 != 0 {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v650 = int32(10)
	v652 = v646*v650 - v647
	v653 = int32(*(*int8)(unsafe.Add(mBase, uint32(v648)+1)))
	v657 = v653 - int32(48)
	if base.Ui32(v657) < base.Ui32(v650) {
		v646 = v652
		v647 = v657
		v648 = v648 + int32(1)
		goto L194
	} else {
		goto L196
	}
L195:
	;
	v660 = v652
	goto L193
L196:
	;
	goto L195
L197:
	;
	v666 = int32(0) - v660
	goto L199
L198:
	;
	v666 = v660
	goto L199
L199:
	;
	goto L184
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	F_ResourceOwnerForget(m, v4, l0, int32(786256))
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[776]))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[777]))
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
	v53 = *(*int32)(unsafe.Add(mBase, _consts[777]))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v17 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v13+int32(32), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(1)
	v34 = F_systable_beginscan(m, v17, int32(2680), v29, int32(0), v29, v13+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L34
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v36 = F_systable_getnext(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = v36
	v47 = v5
	goto L11
L9:
	;
	v82 = v5
	goto L10
L10:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L23
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v52 = v50 + v51
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v82 = v72
	goto L10
L13:
	;
	v73 = F_systable_getnext(m, v34)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v53 != l1 {
		v72 = v47
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+12)))
	if base.B2i32(l2 == int32(0))&(v57&int32(1)) != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if (v57|(l2^int32(1)))&int32(1) == int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_CatalogTupleDelete(m, v17, v40+int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v72 = int32(1)
	goto L13
L21:
	;
	if v73 != 0 {
		v40 = v73
		v47 = v72
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L12
L23:
	;
	F_sequence_close(m, v17, int32(3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v13 + int32(80)
	return v82
L25:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
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
	v102 = l3
	goto L29
L28:
	;
	v102 = int32(274366)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v102
	F_errmsg(m, int32(731444), v13+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errdetail(m, int32(648196), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(638995), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(513076), int32(595), int32(400569))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
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
	v128 = m.ExcPending
	if v128 != 0 {
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
	v130 = l3
	goto L38
L37:
	;
	v130 = int32(274366)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v130
	F_errmsg(m, int32(731473), v13)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errdetail(m, int32(648158), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(513076), int32(601), int32(400569))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v10 = F_DetermineTimeZoneOffsetInternal(m, l0, l1, v6+int32(8))
	m.G0 = v6 + v5
	return v10
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
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
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
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
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
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
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
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1&int32(4304) == int32(4224) {
		v250 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v250 & int32(1)
L2:
	;
	v19 = int32(12)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2*v19)+uint32(_consts[36])))
	v35 = F_GetMultiXactIdMembers(m, l0, v13+v19, int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32))|base.B2i32(l1&int32(4176) == int32(64)))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v35 < int32(0) {
		v250 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v35 == int32(0) {
		v237 = v5
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L3
	} else {
		goto L74
	}
L7:
	;
	v45 = int32(0)
	v48 = v5
	goto L8
L8:
	;
	if v48&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v237 = v229
	goto L6
L10:
	;
	v56 = int32(1)
	if l3 == int32(0) {
		v237 = v56
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v61 = int32(3)
	v62 = v45 << (uint(v61) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v64 = v62 + v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32))+uint32(_consts[37])))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70*int32(12))+uint32(_consts[36])))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if base.Ui32(v76) < base.Ui32(v61) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v59 != 0 {
		v237 = v56
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v231 = v45 + int32(1)
	if v231 != v35 {
		v45 = v231
		v48 = v229
		goto L8
	} else {
		goto L73
	}
L16:
	;
	if v196 != 0 {
		goto L56
	} else {
		goto L57
	}
L17:
	;
	v196 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v87 == v76 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v196 = int32(1)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v91 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v196 = v188
	goto L16
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v95 == int32(0) {
		v188 = int32(0)
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	v159 = int32(0)
	v161 = v91 - int32(1)
	goto L46
L27:
	;
	v100 = v95
	goto L28
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	if v105 == int32(4) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v188 = int32(0)
	goto L23
L30:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v100)+80))
	if v152 != 0 {
		v100 = v152
		goto L28
	} else {
		goto L45
	}
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v108 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v111 = int32(1)
	if v76 == v108 {
		v188 = v111
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	v115 = v113 - int32(1)
	if v115 < int32(0) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v120 = int32(0)
	v122 = v115
	goto L35
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	v128 = int32(2)
	v129 = base.I32_div_s(v122-v120, v128)
	v130 = v129 + v120
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v126+v130<<(uint(v128)%32))))
	if v134 == v76 {
		v188 = v111
		goto L23
	} else {
		goto L37
	}
L36:
	;
	goto L30
L37:
	;
	v138 = F_TransactionIdPrecedes(m, v134, v76)
	mBase = m.M
	if v138 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v139 = v130 + int32(1)
	goto L40
L39:
	;
	v139 = v120
	goto L40
L40:
	;
	if v138 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v142 = v122
	goto L43
L42:
	;
	v142 = v130 - int32(1)
	goto L43
L43:
	;
	if v139 <= v142 {
		v120 = v139
		v122 = v142
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
	v166 = int32(2)
	v167 = base.I32_div_s(v161-v159, v166)
	v168 = v167 + v159
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v157+v168<<(uint(v166)%32))))
	v173 = base.B2i32(v172 == v76)
	if v172 == v76 {
		v188 = v173
		goto L23
	} else {
		goto L48
	}
L47:
	;
	v188 = v173
	goto L23
L48:
	;
	v176 = base.B2i32(base.Ui32(v172) < base.Ui32(v76))
	if base.Ui32(v172) < base.Ui32(v76) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v177 = v168 + int32(1)
	goto L51
L50:
	;
	v177 = v159
	goto L51
L51:
	;
	if base.Ui32(v172) < base.Ui32(v76) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v180 = v161
	goto L54
L53:
	;
	v180 = v168 - int32(1)
	goto L54
L54:
	;
	if v177 <= v180 {
		v159 = v177
		v161 = v180
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
	v201 = int32(1)
	if v48&v201 != 0 {
		v229 = v201
		goto L15
	} else {
		goto L62
	}
L59:
	;
	v229 = v48
	goto L15
L60:
	;
	goto L61
L61:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v199)
	v229 = v48
	goto L15
L62:
	;
	v204 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v75<<(uint(int32(2))%32))+uint32(_consts[42])))
	goto L63
L63:
	;
	if int32(base.Ui32(v209)>>(uint(v23)%32))&int32(1) == int32(0) {
		v229 = v204
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215+v62)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v217) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v229 = int32(1)
	goto L15
L66:
	;
	v220 = F_TransactionIdDidAbort(m, v76)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v224 = F_TransactionIdIsInProgress(m, v76)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	if v220 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v229 = v204
	goto L15
L71:
	;
	if v224 == int32(0) {
		v229 = v204
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
	v250 = v237
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
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
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v24 = int32(0)
	v25 = v8
	goto L3
L5:
	;
	goto L6
L6:
	;
	v16 = int32(1)
	F_ScanKeyInit(m, v8, v16, int32(3), int32(184), l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v24 = v16
	v25 = v8 + int32(48)
	goto L3
L8:
	;
	F_ScanKeyInit(m, v25, int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v33 = v24
	goto L10
L10:
	;
	v34 = F_table_beginscan_catalog(m, v12, v33, v8)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v33 = v24 + int32(1)
	goto L10
L12:
	;
	v36 = F_heap_getnext(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v36
	goto L17
L15:
	;
	goto L16
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	m.T0[v56].(func(*base.Module, int32))(m, v34)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	F_CatalogTupleDelete(m, v12, v40+int32(4))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v47 = F_heap_getnext(m, v34)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v47 != 0 {
		v40 = v47
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	F_sequence_close(m, v12, int32(3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
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
	var v127 int64
	_ = v127
	var v132 int32
	_ = v132
	var v145 float64
	_ = v145
	var v156 float64
	_ = v156
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v175 float64
	_ = v175
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v187 float64
	_ = v187
	var v193 float64
	_ = v193
	var v198 float64
	_ = v198
	var v202 float64
	_ = v202
	var v206 float64
	_ = v206
	var v212 float64
	_ = v212
	var v220 int64
	_ = v220
	var v225 int32
	_ = v225
	var v248 float64
	_ = v248
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v262 float64
	_ = v262
	var v267 float64
	_ = v267
	var v271 float64
	_ = v271
	var v280 float64
	_ = v280
	var v289 float64
	_ = v289
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v302 float64
	_ = v302
	var v306 float64
	_ = v306
	var v314 int64
	_ = v314
	var v319 int32
	_ = v319
	var v332 float64
	_ = v332
	var v343 float64
	_ = v343
	var v355 float64
	_ = v355
	var v356 float64
	_ = v356
	var v357 float64
	_ = v357
	var v362 float64
	_ = v362
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v369 float64
	_ = v369
	var v374 float64
	_ = v374
	var v380 float64
	_ = v380
	var v385 float64
	_ = v385
	var v389 float64
	_ = v389
	var v393 float64
	_ = v393
	var v399 float64
	_ = v399
	var v402 float64
	_ = v402
	var v406 float64
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_abs(v12)
	if base.Ui64(base.I64_reinterpret_f64(v13)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1042])))
		if v18 == int32(0) {
			F_init_degree_constants(m)
			mBase = m.M
		} else {
		}
		if base.F64_gt(v13, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v418 = m.ExcPending
			if v418 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v421 = m.ExcPending
				if v421 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418208), int32(0))
					mBase = m.M
					v425 = m.ExcPending
					if v425 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(2164), int32(442206))
						mBase = m.M
						v430 = m.ExcPending
						if v430 != 0 {
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
					v119 = *(*float64)(unsafe.Add(mBase, _consts[1043]))
					v402 = base.F64_mul(base.F64_div(v115, v119), float64(30))
				} else {
					v127 = base.I64_reinterpret_f64(v12)
					v132 = base.I32_wrap_i64(int64(base.Ui64(v127)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v132) {
						if base.I32_wrap_i64(v127)|(v132-int32(1072693248)) == int32(0) {
							if int64(0) <= v127 {
								v145 = float64(0)
							} else {
								v145 = float64(3.141592653589793)
							}
							v202 = v145
						} else {
							v202 = base.F64_div(float64(0), base.F64_sub(v12, v12))
						}
					} else {
						if base.Ui32(v132) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v132) < base.Ui32(int32(1012924417)) {
								v198 = float64(1.5707963267948966)
								v202 = v198
							} else {
								v156 = F_R(m, base.F64_mul(v12, v12))
								mBase = m.M
								v202 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v12, v156)), v12), float64(1.5707963267948966))
							}
						} else {
							if v127 < int64(0) {
								v168 = base.F64_mul(base.F64_add(v12, float64(1)), float64(0.5))
								v169 = base.F64_sqrt(v168)
								v170 = F_R(m, v168)
								mBase = m.M
								v175 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v169, base.F64_add(base.F64_mul(v169, v170), float64(-6.123233995736766e-17))))
								v202 = base.F64_add(v175, v175)
							} else {
								v180 = base.F64_mul(base.F64_sub(float64(1), v12), float64(0.5))
								v181 = base.F64_sqrt(v180)
								v182 = F_R(m, v180)
								mBase = m.M
								v187 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v181) & int64(-4294967296))
								v193 = base.F64_add(base.F64_add(base.F64_mul(v181, v182), base.F64_div(base.F64_sub(v180, base.F64_mul(v187, v187)), base.F64_add(v181, v187))), v187)
								v198 = base.F64_add(v193, v193)
								v202 = v198
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v202
					v206 = *(*float64)(unsafe.Add(mBase, _consts[1044]))
					v402 = base.F64_add(base.F64_mul(base.F64_div(v202, v206), float64(-60)), float64(90))
				}
			} else {
				v212 = base.F64_neg(v12)
				if base.F64_ge(v12, float64(-0.5)) != 0 {
					v220 = base.I64_reinterpret_f64(v212)
					v225 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v225) {
						if base.I32_wrap_i64(v220)|(v225-int32(1072693248)) == int32(0) {
							v302 = base.F64_add(base.F64_mul(v212, float64(1.5707963267948966)), float64(7.52316384526264e-37))
						} else {
							v302 = base.F64_div(float64(0), base.F64_sub(v212, v212))
						}
					} else {
						if base.Ui32(v225) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v225+int32(-1048576)) < base.Ui32(int32(1044381696)) {
								v294 = v212
								v302 = v294
							} else {
								v248 = F_R(m, base.F64_mul(v212, v212))
								mBase = m.M
								v302 = base.F64_add(base.F64_mul(v212, v248), v212)
							}
						} else {
							v255 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v212)), float64(0.5))
							v256 = base.F64_sqrt(v255)
							v257 = F_R(m, v255)
							mBase = m.M
							if base.Ui32(int32(1072640819)) <= base.Ui32(v225) {
								v262 = base.F64_add(base.F64_mul(v256, v257), v256)
								v289 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v262, v262), float64(-6.123233995736766e-17)))
							} else {
								v267 = float64(0.7853981633974483)
								v271 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v256) & int64(-4294967296))
								v280 = base.F64_div(base.F64_sub(v255, base.F64_mul(v271, v271)), base.F64_add(v256, v271))
								v289 = base.F64_add(base.F64_sub(base.F64_sub(v267, base.F64_add(v271, v271)), base.F64_sub(base.F64_mul(base.F64_add(v256, v256), v257), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v280, v280)))), v267)
							}
							if v220 < int64(0) {
								v293 = base.F64_neg(v289)
							} else {
								v293 = v289
							}
							v294 = v293
							v302 = v294
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v302
					v306 = *(*float64)(unsafe.Add(mBase, _consts[1043]))
					v399 = base.F64_mul(base.F64_div(v302, v306), float64(30))
				} else {
					v314 = base.I64_reinterpret_f64(v212)
					v319 = base.I32_wrap_i64(int64(base.Ui64(v314)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v319) {
						if base.I32_wrap_i64(v314)|(v319-int32(1072693248)) == int32(0) {
							if int64(0) <= v314 {
								v332 = float64(0)
							} else {
								v332 = float64(3.141592653589793)
							}
							v389 = v332
						} else {
							v389 = base.F64_div(float64(0), base.F64_sub(v212, v212))
						}
					} else {
						if base.Ui32(v319) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v319) < base.Ui32(int32(1012924417)) {
								v385 = float64(1.5707963267948966)
								v389 = v385
							} else {
								v343 = F_R(m, base.F64_mul(v212, v212))
								mBase = m.M
								v389 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v212, v343)), v212), float64(1.5707963267948966))
							}
						} else {
							if v314 < int64(0) {
								v355 = base.F64_mul(base.F64_add(v212, float64(1)), float64(0.5))
								v356 = base.F64_sqrt(v355)
								v357 = F_R(m, v355)
								mBase = m.M
								v362 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v356, base.F64_add(base.F64_mul(v356, v357), float64(-6.123233995736766e-17))))
								v389 = base.F64_add(v362, v362)
							} else {
								v367 = base.F64_mul(base.F64_sub(float64(1), v212), float64(0.5))
								v368 = base.F64_sqrt(v367)
								v369 = F_R(m, v367)
								mBase = m.M
								v374 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v368) & int64(-4294967296))
								v380 = base.F64_add(base.F64_add(base.F64_mul(v368, v369), base.F64_div(base.F64_sub(v367, base.F64_mul(v374, v374)), base.F64_add(v368, v374))), v374)
								v385 = base.F64_add(v380, v380)
								v389 = v385
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v389
					v393 = *(*float64)(unsafe.Add(mBase, _consts[1044]))
					v399 = base.F64_add(base.F64_mul(base.F64_div(v389, v393), float64(-60)), float64(90))
				}
				v402 = base.F64_neg(v399)
			}
			if base.F64_eq(base.F64_abs(v402), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v432 = m.ExcPending
				if v432 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v406 = v402
				v407 = F_Float8GetDatum(m, v406)
				mBase = m.M
				v410 = m.ExcPending
				if v410 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v407
				}
			}
		}
	} else {
		v406 = math.Float64frombits(uint64(0x7ff8000000000000))
		v407 = F_Float8GetDatum(m, v406)
		mBase = m.M
		v410 = m.ExcPending
		if v410 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v407
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
		v10 = int32(4800)
	} else {
		v10 = int32(4799)
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
	v31 = base.I32_div_s((v26+l1)*int32(7834), int32(256))
	v34 = l2 + v11*int32(365) + v16 + v19 + v22 + v31 - int32(32167)
	v43 = int32(4799) + l0
	v48 = base.I32_div_s(v43, int32(4))
	v51 = base.I32_div_s(v43, int32(-100))
	v54 = base.I32_div_s(v43, int32(400))
	v63 = base.I32_div_s(int32(109676), int32(256))
	v66 = int32(4) + v43*int32(365) + v48 + v51 + v54 + v63 - int32(32167)
	v67 = int32(1)
	v71 = int32(7)
	v72 = base.I32_rem_s(v66-v67+v67, v71)
	if v72 < int32(0) {
		v77 = v72 + v71
	} else {
		v77 = v72
	}
	if v34 < v66-v77 {
		v90 = int32(4799) + (l0 - int32(1))
		v95 = base.I32_div_s(v90, int32(4))
		v98 = base.I32_div_s(v90, int32(-100))
		v101 = base.I32_div_s(v90, int32(400))
		v110 = base.I32_div_s(int32(109676), int32(256))
		v113 = int32(4) + v90*int32(365) + v95 + v98 + v101 + v110 - int32(32167)
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
		v141 = l0 + int32(4800)
		v146 = base.I32_div_s(v141, int32(4))
		v149 = base.I32_div_s(v141, int32(-100))
		v152 = base.I32_div_s(v141, int32(400))
		v161 = base.I32_div_s(int32(109676), int32(256))
		v164 = int32(4) + v141*int32(365) + v146 + v149 + v152 + v161 - int32(32167)
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
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
			if base.Ui32(v12) <= base.Ui32(int32(155)) {
				v24 = F_DCH_cache_fetch(m, v8, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_DCH_to_char(m, v24, l2, l0, v17, l3)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
			} else {
				v30 = F_palloc(m, v14+int32(12))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_parse_format(m, v30, v8, int32(1690976), int32(1690240), int32(1690368), int32(1), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_DCH_to_char(m, v30, l2, l0, v17, l3)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v30)
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
		*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(0)
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
					F_errmsg(m, int32(418208), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(1898), int32(144166))
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
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
	v10 = l0 - int32(44032)
	if base.Ui32(v10) <= base.Ui32(int32(11171)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v15 = int32(2)
		v18 = int32(65535)
		v19 = v10 & v18
		v20 = int32(588)
		v21 = base.I32_div_u_s(v19, v20)
		*(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(v15)%32)))) = v21 | int32(4352)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v26 = int32(1)
		v27 = v25 + v26
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v27
		v37 = int32(28)
		v38 = base.I32_div_u_s((v10-v21*v20)&v18, v37)
		*(*int32)(unsafe.Add(mBase, uint32(v13+v27<<(uint(v15)%32)))) = v38 + int32(4449)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v44 = v42 + v26
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v44
		v47 = base.I32_rem_u_s(v19, v37)
		if v47 == int32(0) {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13+v44<<(uint(int32(2))%32)))) = v47 + int32(4519)
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
			return
		}
	} else {
		v57 = l0 & int32(255)
		v58 = int32(8)
		v63 = int32(base.Ui32(l0<<(uint(v58)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
		v69 = int32(base.Ui32(int32(base.Ui32(l0)>>(uint(v58)%32))&int32(65280)) >> (uint(v58) % 32))
		v71 = int32(base.Ui32(l0) >> (uint(int32(24)) % 32))
		v72 = int32(8191)
		v83 = int32(13687)
		v84 = base.I32_rem_u_s(v57+(v63+(v69+v71*v72)*v72)*v72+int32(402620417), v83)
		v85 = int32(1)
		v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84<<(uint(v85)%32))+uint32(_consts[1324]))))
		v90 = int32(257)
		v100 = base.I32_rem_u_s(((v71*v90+v69)*v90+v63)*v90+v57, v83)
		v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100<<(uint(v85)%32))+uint32(_consts[1324]))))
		v106 = v89 + v105
		if base.Ui32(int32(6842)) < base.Ui32(v106) {
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
			return
		} else {
			v110 = v106 << (uint(int32(3)) % 32)
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[1325])))
			if l0 != v113 {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
				v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
				return
			} else {
				v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[1326]))))
				v117 = v115 & int32(31)
				if v117 == int32(0) {
					v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)))) = l0
					v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v140 + int32(1)
					return
				} else {
					if l1 != 0 {
						v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[1327]))))
						if v115&int32(64) != 0 {
							v147 = int32(4632544)
							*(*int32)(unsafe.Add(mBase, _consts[1328])) = v144
							v155 = int32(1)
							v156 = v147
						} else {
							v155 = v117
							v156 = v144<<(uint(int32(2))%32) + int32(2090560)
						}
						v158 = int32(0)
						for {
							v169 = *(*int32)(unsafe.Add(mBase, uint32(v156+v158<<(uint(int32(2))%32))))
							v175 = v169 - int32(44032)
							if base.Ui32(v175) <= base.Ui32(int32(11171)) {
								v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v180 = int32(2)
								v183 = int32(65535)
								v184 = v175 & v183
								v185 = int32(588)
								v186 = base.I32_div_u_s(v184, v185)
								*(*int32)(unsafe.Add(mBase, uint32(v178+v179<<(uint(v180)%32)))) = v186 | int32(4352)
								v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v191 = int32(1)
								v192 = v190 + v191
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v192
								v202 = int32(28)
								v203 = base.I32_div_u_s((v175-v186*v185)&v183, v202)
								*(*int32)(unsafe.Add(mBase, uint32(v178+v192<<(uint(v180)%32)))) = v203 + int32(4449)
								v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v209 = v207 + v191
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v209
								v212 = base.I32_rem_u_s(v184, v202)
								if v212 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v178+v209<<(uint(int32(2))%32)))) = v212 + int32(4519)
									v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
								}
							} else {
								v222 = v169 & int32(255)
								v223 = int32(8)
								v228 = int32(base.Ui32(v169<<(uint(v223)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
								v234 = int32(base.Ui32(int32(base.Ui32(v169)>>(uint(v223)%32))&int32(65280)) >> (uint(v223) % 32))
								v236 = int32(base.Ui32(v169) >> (uint(int32(24)) % 32))
								v237 = int32(8191)
								v248 = int32(13687)
								v249 = base.I32_rem_u_s(v222+(v228+(v234+v236*v237)*v237)*v237+int32(402620417), v248)
								v250 = int32(1)
								v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249<<(uint(v250)%32))+uint32(_consts[1324]))))
								v255 = int32(257)
								v265 = base.I32_rem_u_s(((v236*v255+v234)*v255+v228)*v255+v222, v248)
								v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v265<<(uint(v250)%32))+uint32(_consts[1324]))))
								v271 = v254 + v270
								if base.Ui32(int32(6842)) < base.Ui32(v271) {
									v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
									v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
								} else {
									v275 = v271 << (uint(int32(3)) % 32)
									v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1325])))
									if v169 != v278 {
										v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
										v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
									} else {
										v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1326]))))
										v282 = v280 & int32(31)
										if v282 == int32(0) {
											v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
											v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
										} else {
											if l1 != 0 {
												v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1327]))))
												if v280&int32(64) != 0 {
													v312 = int32(4632544)
													*(*int32)(unsafe.Add(mBase, _consts[1328])) = v309
													v320 = int32(1)
													v321 = v312
												} else {
													v320 = v282
													v321 = v309<<(uint(int32(2))%32) + int32(2090560)
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
												if v280&int32(32) == int32(0) {
													v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1327]))))
													if v280&int32(64) != 0 {
														v312 = int32(4632544)
														*(*int32)(unsafe.Add(mBase, _consts[1328])) = v309
														v320 = int32(1)
														v321 = v312
													} else {
														v320 = v282
														v321 = v309<<(uint(int32(2))%32) + int32(2090560)
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
						if v115&int32(32) == int32(0) {
							v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[1327]))))
							if v115&int32(64) != 0 {
								v147 = int32(4632544)
								*(*int32)(unsafe.Add(mBase, _consts[1328])) = v144
								v155 = int32(1)
								v156 = v147
							} else {
								v155 = v117
								v156 = v144<<(uint(int32(2))%32) + int32(2090560)
							}
							v158 = int32(0)
							for {
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v156+v158<<(uint(int32(2))%32))))
								v175 = v169 - int32(44032)
								if base.Ui32(v175) <= base.Ui32(int32(11171)) {
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v180 = int32(2)
									v183 = int32(65535)
									v184 = v175 & v183
									v185 = int32(588)
									v186 = base.I32_div_u_s(v184, v185)
									*(*int32)(unsafe.Add(mBase, uint32(v178+v179<<(uint(v180)%32)))) = v186 | int32(4352)
									v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v191 = int32(1)
									v192 = v190 + v191
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v192
									v202 = int32(28)
									v203 = base.I32_div_u_s((v175-v186*v185)&v183, v202)
									*(*int32)(unsafe.Add(mBase, uint32(v178+v192<<(uint(v180)%32)))) = v203 + int32(4449)
									v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v209 = v207 + v191
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v209
									v212 = base.I32_rem_u_s(v184, v202)
									if v212 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v178+v209<<(uint(int32(2))%32)))) = v212 + int32(4519)
										v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
									}
								} else {
									v222 = v169 & int32(255)
									v223 = int32(8)
									v228 = int32(base.Ui32(v169<<(uint(v223)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
									v234 = int32(base.Ui32(int32(base.Ui32(v169)>>(uint(v223)%32))&int32(65280)) >> (uint(v223) % 32))
									v236 = int32(base.Ui32(v169) >> (uint(int32(24)) % 32))
									v237 = int32(8191)
									v248 = int32(13687)
									v249 = base.I32_rem_u_s(v222+(v228+(v234+v236*v237)*v237)*v237+int32(402620417), v248)
									v250 = int32(1)
									v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249<<(uint(v250)%32))+uint32(_consts[1324]))))
									v255 = int32(257)
									v265 = base.I32_rem_u_s(((v236*v255+v234)*v255+v228)*v255+v222, v248)
									v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v265<<(uint(v250)%32))+uint32(_consts[1324]))))
									v271 = v254 + v270
									if base.Ui32(int32(6842)) < base.Ui32(v271) {
										v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
										v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
									} else {
										v275 = v271 << (uint(int32(3)) % 32)
										v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1325])))
										if v169 != v278 {
											v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
											v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
										} else {
											v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1326]))))
											v282 = v280 & int32(31)
											if v282 == int32(0) {
												v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = v169
												v305 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v305 + int32(1)
											} else {
												if l1 != 0 {
													v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1327]))))
													if v280&int32(64) != 0 {
														v312 = int32(4632544)
														*(*int32)(unsafe.Add(mBase, _consts[1328])) = v309
														v320 = int32(1)
														v321 = v312
													} else {
														v320 = v282
														v321 = v309<<(uint(int32(2))%32) + int32(2090560)
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
													if v280&int32(32) == int32(0) {
														v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1327]))))
														if v280&int32(64) != 0 {
															v312 = int32(4632544)
															*(*int32)(unsafe.Add(mBase, _consts[1328])) = v309
															v320 = int32(1)
															v321 = v312
														} else {
															v320 = v282
															v321 = v309<<(uint(int32(2))%32) + int32(2090560)
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
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
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
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
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
	v55 = F_palloc(m, v52+int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	v25 = int32(4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v27&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v40 = int32(1)
	if v22&v40 != 0 {
		v52 = int32(base.Ui32(v22)>>(uint(v40)%32)) - v40
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v36 = v25
	goto L9
L8:
	;
	v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
	goto L9
L9:
	;
	if v27 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v25
	goto L12
L11:
	;
	v39 = v36
	goto L12
L12:
	;
	v52 = v39
	goto L1
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v57 = int32(1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v59&v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	F_pfree(m, v55)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L100
	}
L16:
	;
	v62 = v57
	goto L18
L17:
	;
	v62 = int32(4)
	goto L18
L18:
	;
	v63 = v18 + v62
	v64 = v63 + v52
	if base.Ui32(v64) <= base.Ui32(v63) {
		v318 = v2
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v68 = v63
	v71 = v2
	v72 = v2
	v73 = v2
	v75 = v2
	goto L20
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	switch v72 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L37
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
		goto L39
	}
L21:
	;
	switch v283 {
	case 0:
		v318 = v284
		goto L15
	default:
		goto L92
	case 7:
		goto L91
	}
L22:
	;
	v287 = v280 + int32(1)
	if base.Ui32(v287) < base.Ui32(v64) {
		v68 = v287
		v71 = v282
		v72 = v283
		v73 = v284
		v75 = v285
		goto L20
	} else {
		goto L90
	}
L23:
	;
	v280 = v274
	v282 = v276
	v283 = v277
	v284 = v278
	v285 = v75
	goto L22
L24:
	;
	v274 = v68
	v276 = v71 + int32(1)
	v277 = int32(5)
	v278 = v73
	goto L23
L25:
	;
	v274 = v68
	v276 = v71
	v277 = v269
	v278 = v73
	goto L23
L26:
	;
	v269 = int32(4)
	goto L25
L27:
	;
	v269 = int32(3)
	goto L25
L28:
	;
	v280 = v260
	v282 = v71
	v283 = v262
	v284 = v73
	v285 = v71
	goto L22
L29:
	;
	v260 = v68
	v262 = int32(6)
	goto L28
L30:
	;
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23, 35:
		goto L87
	default:
		goto L86
	}
L31:
	;
	if v81 == int32(34) {
		goto L75
	} else {
		goto L76
	}
L32:
	;
	if v81 != int32(92) {
		goto L60
	} else {
		goto L61
	}
L33:
	;
	v152 = int32(5)
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		goto L26
	default:
		goto L55
	case 25:
		goto L29
	case 30:
		v280 = v68
		v282 = v71
		v283 = v152
		v284 = v73
		v285 = v71
		goto L22
	case 60:
		goto L56
	}
L34:
	;
	if base.Ui32(v81-int32(9)) < base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L47
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v274 = v68
	v276 = v71 + int32(1)
	v277 = int32(2)
	v278 = v73
	goto L23
L36:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v115)
	v274 = v68
	v276 = v71 + int32(1)
	v277 = int32(3)
	v278 = v73
	goto L23
L37:
	;
	if v81 != int32(34) {
		goto L35
	} else {
		goto L44
	}
L38:
	;
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		goto L36
	default:
		goto L42
	case 52:
		goto L43
	}
L39:
	;
	switch v81 - int32(9) {
	case 0, 1, 2, 3, 4, 23, 35:
		v280 = v68
		v282 = v71
		v283 = int32(0)
		v284 = v73
		v285 = v75
		goto L22
	default:
		goto L40
	case 25:
		goto L41
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v81)
	v274 = v68
	v276 = v55 + int32(1)
	v277 = int32(1)
	v278 = v73
	goto L23
L41:
	;
	v274 = v68
	v276 = v55
	v277 = int32(2)
	v278 = v73
	goto L23
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v98 = int32(1)
	v274 = v68
	v276 = v71 + v98
	v277 = v98
	v278 = v73
	goto L23
L43:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v92)
	v274 = v68
	v276 = v71 + int32(1)
	v277 = int32(4)
	v278 = v73
	goto L23
L44:
	;
	v104 = v68 + int32(1)
	if base.Ui32(v64) <= base.Ui32(v104) {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v106 != int32(34) {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v109 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v109)
	v274 = v104
	v276 = v71 + int32(1)
	v277 = int32(2)
	v278 = v73
	goto L23
L47:
	;
	if v81 == int32(32) {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	if v81 == int32(61) {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v139 = F_text_to_cstring(m, v18)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v139
	F_errmsg(m, int32(753025), v16+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(514289), int32(1708), int32(79154))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v280 = v68
	v282 = v71 + int32(1)
	v283 = int32(7)
	v284 = v73
	v285 = v71
	goto L22
L56:
	;
	v156 = v68 + int32(1)
	if base.Ui32(v64) <= base.Ui32(v156) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v158 != int32(39) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v260 = v156
	v262 = v152
	goto L28
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	goto L24
L60:
	;
	if v81 != int32(39) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v198 = v68 + int32(1)
	if base.Ui32(v64) <= base.Ui32(v198) {
		goto L72
	} else {
		goto L73
	}
L63:
	;
	v171 = v68 + int32(1)
	if base.Ui32(v64) <= base.Ui32(v171) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v181)
	v186 = F_pstrdup(m, v55)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L67
	}
L65:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v173 != int32(39) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v176 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v176)
	v274 = v171
	v276 = v71 + int32(1)
	v277 = int32(5)
	v278 = v73
	goto L23
L67:
	;
	v188 = F_pstrdup(m, v75)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v190 = F_makeString(m, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v193 = F_makeDefElem(m, v186, v190, int32(-1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v195 = F_lappend(m, v73, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v274 = v68
	v276 = v71 + int32(1)
	v277 = v181
	v278 = v195
	goto L23
L72:
	;
	v208 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v208)
	goto L24
L73:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v200 != int32(92) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v203 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v203)
	v274 = v198
	v276 = v71 + int32(1)
	v277 = int32(5)
	v278 = v73
	goto L23
L75:
	;
	v214 = v68 + int32(1)
	if base.Ui32(v64) <= base.Ui32(v214) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v274 = v68
	v276 = v71 + int32(1)
	v277 = int32(6)
	v278 = v73
	goto L23
L78:
	;
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v224)
	v229 = F_pstrdup(m, v55)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L81
	}
L79:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v216 != int32(34) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v219 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v219)
	v274 = v214
	v276 = v71 + int32(1)
	v277 = int32(6)
	v278 = v73
	goto L23
L81:
	;
	v231 = F_pstrdup(m, v75)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v233 = F_makeString(m, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v236 = F_makeDefElem(m, v229, v233, int32(-1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v238 = F_lappend(m, v73, v236)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v274 = v68
	v276 = v71 + int32(1)
	v277 = v224
	v278 = v238
	goto L23
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v81)
	v274 = v68
	v276 = v71 + int32(1)
	v277 = int32(7)
	v278 = v73
	goto L23
L87:
	;
	v246 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v246)
	v251 = F_buildDefItem(m, v55, v75)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	v253 = F_lappend(m, v73, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v274 = v68
	v276 = v71 + int32(1)
	v277 = v246
	v278 = v253
	goto L23
L90:
	;
	goto L21
L91:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v307)
	v309 = F_buildDefItem(m, v55, v285)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L2
	} else {
		goto L98
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v296 = F_text_to_cstring(m, v18)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v296
	F_errmsg(m, int32(753025), v16)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(514289), int32(1823), int32(79154))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	v311 = F_lappend(m, v284, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v318 = v311
	goto L15
L100:
	;
	m.G0 = v16 + int32(32)
	return v318
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 float64
	_ = v61
	var v66 float64
	_ = v66
	var v73 float64
	_ = v73
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v95 int32
	_ = v95
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v107 int32
	_ = v107
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v142 int32
	_ = v142
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v157 float64
	_ = v157
	var v158 float64
	_ = v158
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	var v166 float64
	_ = v166
	var v168 float64
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 float64
	_ = v184
	var v192 float64
	_ = v192
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v247 float64
	_ = v247
	var v263 float64
	_ = v263
	var v269 float64
	_ = v269
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v324 float64
	_ = v324
	var v341 float64
	_ = v341
	var v350 float64
	_ = v350
	var v351 float64
	_ = v351
	var v352 float64
	_ = v352
	var v355 float64
	_ = v355
	var v373 float64
	_ = v373
	var v374 float64
	_ = v374
	var v388 float64
	_ = v388
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	v2 = float64(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
		v388 = v12
		v397 = F_Float8GetDatum(m, v388)
		mBase = m.M
		v400 = m.ExcPending
		if v400 != 0 {
			return int32(0)
		} else {
			return v397
		}
	} else {
		if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_lt(v12, float64(0)) == int32(0) {
				v388 = v12
				v397 = F_Float8GetDatum(m, v388)
				mBase = m.M
				v400 = m.ExcPending
				if v400 != 0 {
					return int32(0)
				} else {
					return v397
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v415 = m.ExcPending
				if v415 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(0)
			v30 = base.I64_reinterpret_f64(v12)
			v35 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(2146435072)) <= base.Ui32(v35) {
				v373 = base.F64_add(v12, math.Float64frombits(uint64(0x7ff0000000000000)))
			} else {
				if base.Ui32(v35) <= base.Ui32(int32(1016070143)) {
					v373 = base.F64_div(float64(1), v12)
				} else {
					if base.F64_ne(v12, base.F64_floor(v12)) != 0 {
						if base.Ui32(int32(1080492032)) <= base.Ui32(v35) {
							v66 = float64(0.5)
							if base.F64_eq(base.F64_floor(base.F64_mul(v12, v66)), base.F64_mul(base.F64_floor(v12), v66)) != 0 {
								v73 = float64(0)
							} else {
								v73 = math.Float64frombits(uint64(0x8000000000000000))
							}
							if v30 < int64(0) {
								v373 = v73
							} else {
								v373 = base.F64_mul(v12, float64(8.98846567431158e+307))
							}
						} else {
							v78 = base.F64_abs(v12)
							v79 = float64(5.52468004077673)
							v80 = base.F64_add(v78, v79)
							v82 = float64(-5.52468004077673)
							if base.F64_gt(v78, v79) != 0 {
								v89 = base.F64_add(base.F64_sub(v80, v78), v82)
							} else {
								v89 = base.F64_sub(base.F64_add(v80, v82), v78)
							}
							v91 = base.F64_add(v78, float64(-0.5))
							if base.F64_lt(v78, float64(8)) != 0 {
								v95 = int32(12)
								v98 = v2
								v99 = v2
								for {
									v107 = v95 << (uint(int32(3)) % 32)
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1045])))
									v111 = base.F64_add(base.F64_mul(v99, v78), v110)
									v115 = *(*float64)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1046])))
									v116 = base.F64_add(base.F64_mul(v98, v78), v115)
									if v95 != 0 {
										v95 = v95 - int32(1)
										v98 = v116
										v99 = v111
										continue
									} else {
										break
									}
									break
								}
								v148 = v116
								v149 = v111
							} else {
								v122 = v2
								v123 = v2
								v127 = int32(0)
								for {
									v131 = v127 << (uint(int32(3)) % 32)
									v134 = *(*float64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[1045])))
									v135 = base.F64_add(base.F64_div(v123, v78), v134)
									v139 = *(*float64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[1046])))
									v140 = base.F64_add(base.F64_div(v122, v78), v139)
									v142 = v127 + int32(1)
									if v142 != int32(13) {
										v122 = v140
										v123 = v135
										v127 = v142
										continue
									} else {
										break
									}
									break
								}
								v148 = v140
								v149 = v135
							}
							v157 = F_exp(m, base.F64_neg(v80))
							mBase = m.M
							v158 = base.F64_mul(base.F64_div(v148, v149), v157)
							if base.F64_lt(v12, float64(0)) != 0 {
								v163 = base.F64_mul(v78, float64(0.5))
								v165 = base.F64_sub(v163, base.F64_floor(v163))
								v166 = base.F64_add(v165, v165)
								v168 = base.F64_mul(v166, float64(4))
								if base.F64_lt(base.F64_abs(v168), float64(2.147483648e+09)) != 0 {
									v172 = base.I32_trunc_f64_s(v168)
									v174 = v172
								} else {
									v174 = int32(-2147483648)
								}
								v175 = int32(1)
								v178 = base.I32_div_s(v174+v175, int32(2))
								v184 = base.F64_mul(base.F64_sub(v166, base.F64_mul(base.F64_convert_i32_s(v178), float64(0.5))), float64(3.141592653589793))
								switch v178 - v175 {
								case 0:
									v231 = float64(1)
									v232 = base.F64_mul(v184, v184)
									v234 = base.F64_mul(v232, float64(0.5))
									v235 = base.F64_sub(v231, v234)
									v247 = base.F64_mul(v232, v232)
									v341 = base.F64_add(v235, base.F64_add(base.F64_sub(base.F64_sub(v231, v235), v234), base.F64_sub(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v247, v247), base.F64_add(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v184, float64(0)))))
								case 1:
									v263 = base.F64_neg(v184)
									v269 = base.F64_mul(v263, v263)
									v341 = base.F64_add(base.F64_mul(base.F64_mul(v263, v269), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(base.F64_mul(v269, base.F64_mul(v269, v269)), base.F64_add(base.F64_mul(v269, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v269, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v263)
								case 2:
									v308 = float64(1)
									v309 = base.F64_mul(v184, v184)
									v311 = base.F64_mul(v309, float64(0.5))
									v312 = base.F64_sub(v308, v311)
									v324 = base.F64_mul(v309, v309)
									v341 = base.F64_neg(base.F64_add(v312, base.F64_add(base.F64_sub(base.F64_sub(v308, v312), v311), base.F64_sub(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v324, v324), base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v184, float64(0))))))
								default:
									v192 = base.F64_mul(v184, v184)
									v341 = base.F64_add(base.F64_mul(base.F64_mul(v184, v192), base.F64_add(base.F64_mul(v192, base.F64_add(base.F64_mul(base.F64_mul(v192, base.F64_mul(v192, v192)), base.F64_add(base.F64_mul(v192, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v192, base.F64_add(base.F64_mul(v192, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v184)
								}
								v350 = base.F64_div(float64(-3.141592653589793), base.F64_mul(v158, base.F64_mul(v78, v341)))
								v351 = base.F64_neg(v89)
								v352 = base.F64_neg(v91)
							} else {
								v350 = v158
								v351 = v89
								v352 = v91
							}
							v355 = F_pow(m, v80, base.F64_mul(v352, float64(0.5)))
							mBase = m.M
							v373 = base.F64_mul(v355, base.F64_mul(v355, base.F64_add(v350, base.F64_div(base.F64_mul(base.F64_mul(v351, float64(6.02468004077673)), v350), v80))))
						}
					} else {
						if v30 < int64(0) {
							v373 = math.Float64frombits(uint64(0x7ff8000000000000))
						} else {
							if base.F64_le(v12, float64(23)) == int32(0) {
								if base.Ui32(int32(1080492032)) <= base.Ui32(v35) {
									v66 = float64(0.5)
									if base.F64_eq(base.F64_floor(base.F64_mul(v12, v66)), base.F64_mul(base.F64_floor(v12), v66)) != 0 {
										v73 = float64(0)
									} else {
										v73 = math.Float64frombits(uint64(0x8000000000000000))
									}
									if v30 < int64(0) {
										v373 = v73
									} else {
										v373 = base.F64_mul(v12, float64(8.98846567431158e+307))
									}
								} else {
									v78 = base.F64_abs(v12)
									v79 = float64(5.52468004077673)
									v80 = base.F64_add(v78, v79)
									v82 = float64(-5.52468004077673)
									if base.F64_gt(v78, v79) != 0 {
										v89 = base.F64_add(base.F64_sub(v80, v78), v82)
									} else {
										v89 = base.F64_sub(base.F64_add(v80, v82), v78)
									}
									v91 = base.F64_add(v78, float64(-0.5))
									if base.F64_lt(v78, float64(8)) != 0 {
										v95 = int32(12)
										v98 = v2
										v99 = v2
										for {
											v107 = v95 << (uint(int32(3)) % 32)
											v110 = *(*float64)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1045])))
											v111 = base.F64_add(base.F64_mul(v99, v78), v110)
											v115 = *(*float64)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1046])))
											v116 = base.F64_add(base.F64_mul(v98, v78), v115)
											if v95 != 0 {
												v95 = v95 - int32(1)
												v98 = v116
												v99 = v111
												continue
											} else {
												break
											}
											break
										}
										v148 = v116
										v149 = v111
									} else {
										v122 = v2
										v123 = v2
										v127 = int32(0)
										for {
											v131 = v127 << (uint(int32(3)) % 32)
											v134 = *(*float64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[1045])))
											v135 = base.F64_add(base.F64_div(v123, v78), v134)
											v139 = *(*float64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[1046])))
											v140 = base.F64_add(base.F64_div(v122, v78), v139)
											v142 = v127 + int32(1)
											if v142 != int32(13) {
												v122 = v140
												v123 = v135
												v127 = v142
												continue
											} else {
												break
											}
											break
										}
										v148 = v140
										v149 = v135
									}
									v157 = F_exp(m, base.F64_neg(v80))
									mBase = m.M
									v158 = base.F64_mul(base.F64_div(v148, v149), v157)
									if base.F64_lt(v12, float64(0)) != 0 {
										v163 = base.F64_mul(v78, float64(0.5))
										v165 = base.F64_sub(v163, base.F64_floor(v163))
										v166 = base.F64_add(v165, v165)
										v168 = base.F64_mul(v166, float64(4))
										if base.F64_lt(base.F64_abs(v168), float64(2.147483648e+09)) != 0 {
											v172 = base.I32_trunc_f64_s(v168)
											v174 = v172
										} else {
											v174 = int32(-2147483648)
										}
										v175 = int32(1)
										v178 = base.I32_div_s(v174+v175, int32(2))
										v184 = base.F64_mul(base.F64_sub(v166, base.F64_mul(base.F64_convert_i32_s(v178), float64(0.5))), float64(3.141592653589793))
										switch v178 - v175 {
										case 0:
											v231 = float64(1)
											v232 = base.F64_mul(v184, v184)
											v234 = base.F64_mul(v232, float64(0.5))
											v235 = base.F64_sub(v231, v234)
											v247 = base.F64_mul(v232, v232)
											v341 = base.F64_add(v235, base.F64_add(base.F64_sub(base.F64_sub(v231, v235), v234), base.F64_sub(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v247, v247), base.F64_add(base.F64_mul(v232, base.F64_add(base.F64_mul(v232, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v184, float64(0)))))
										case 1:
											v263 = base.F64_neg(v184)
											v269 = base.F64_mul(v263, v263)
											v341 = base.F64_add(base.F64_mul(base.F64_mul(v263, v269), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(base.F64_mul(v269, base.F64_mul(v269, v269)), base.F64_add(base.F64_mul(v269, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v269, base.F64_add(base.F64_mul(v269, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v263)
										case 2:
											v308 = float64(1)
											v309 = base.F64_mul(v184, v184)
											v311 = base.F64_mul(v309, float64(0.5))
											v312 = base.F64_sub(v308, v311)
											v324 = base.F64_mul(v309, v309)
											v341 = base.F64_neg(base.F64_add(v312, base.F64_add(base.F64_sub(base.F64_sub(v308, v312), v311), base.F64_sub(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v324, v324), base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v309, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v184, float64(0))))))
										default:
											v192 = base.F64_mul(v184, v184)
											v341 = base.F64_add(base.F64_mul(base.F64_mul(v184, v192), base.F64_add(base.F64_mul(v192, base.F64_add(base.F64_mul(base.F64_mul(v192, base.F64_mul(v192, v192)), base.F64_add(base.F64_mul(v192, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v192, base.F64_add(base.F64_mul(v192, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v184)
										}
										v350 = base.F64_div(float64(-3.141592653589793), base.F64_mul(v158, base.F64_mul(v78, v341)))
										v351 = base.F64_neg(v89)
										v352 = base.F64_neg(v91)
									} else {
										v350 = v158
										v351 = v89
										v352 = v91
									}
									v355 = F_pow(m, v80, base.F64_mul(v352, float64(0.5)))
									mBase = m.M
									v373 = base.F64_mul(v355, base.F64_mul(v355, base.F64_add(v350, base.F64_div(base.F64_mul(base.F64_mul(v351, float64(6.02468004077673)), v350), v80))))
								}
							} else {
								if base.F64_lt(base.F64_abs(v12), float64(2.147483648e+09)) != 0 {
									v54 = base.I32_trunc_f64_s(v12)
									v56 = v54
								} else {
									v56 = int32(-2147483648)
								}
								v61 = *(*float64)(unsafe.Add(mBase, uint32(v56<<(uint(int32(3))%32))+uint32(_consts[1047])))
								v373 = v61
							}
						}
					}
				}
			}
			v374 = base.F64_abs(v373)
			if base.F64_ne(v374, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v374)) < base.Ui64(int64(9218868437227405313))) == int32(0) {
				if base.F64_ne(v373, float64(0)) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v415 = m.ExcPending
					if v415 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					F_float_underflow_error(m)
					mBase = m.M
					v403 = m.ExcPending
					if v403 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				if base.F64_eq(v373, float64(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v403 = m.ExcPending
					if v403 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v388 = v373
					v397 = F_Float8GetDatum(m, v388)
					mBase = m.M
					v400 = m.ExcPending
					if v400 != 0 {
						return int32(0)
					} else {
						return v397
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
		F_ResourceOwnerForget(m, v7, v4, int32(4427496))
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v73 int32
	_ = v73
	var __phi73 int32
	_ = __phi73
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _consts[473])) = v2
	v20 = v2
	goto L3
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L30
	} else {
		goto L32
	}
L2:
	;
	v164 = int32(-1)
	goto L1
L3:
	;
	v26 = l0 + v20<<(uint(int32(3))%32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = v27 * int32(40)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[475]))))
	if v32 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if int32(0) < v128 {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v37 <= v35 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+4)))
	if v116 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v41 = v35
	goto L9
L9:
	;
	v49 = v41 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[1287])))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v27 != v53 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v61 <= v41 {
		v164 = v41
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v56 = v41 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v56 < v58 {
		v41 = v56
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[1287])))
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)) = uint8(v64)
	v67 = v41 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v67 < v69 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	__phi72 = v41
	__phi73 = v67
	v72 = __phi72
	v73 = __phi73
	goto L19
L17:
	;
	goto L18
L18:
	;
	v102 = int32(4547960)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	*(*int32)(unsafe.Add(mBase, _consts[474])) = v104 - int32(1)
	goto L7
L19:
	;
	v79 = int32(2)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v73<<(uint(v79)%32))+uint32(_consts[1287])))
	*(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(v79)%32))+uint32(_consts[1287]))) = v87
	v90 = v73 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	if v90 < v92 {
		__phi72 = v73
		__phi73 = v90
		v72 = __phi72
		v73 = __phi73
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
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[481]))) = uint8(v121)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v124 = v20 + int32(1)
	if v124 != int32(2) {
		v20 = v124
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L4
L26:
	;
	v134 = m.G0
	v135 = int32(16)
	v136 = v134 - v135
	m.G0 = v136
	F___gettimeofday(m, v136)
	mBase = m.M
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
	m.G0 = v136 + v135
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
	F_schedule_alarm(m, v140+v139*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v164
	v177 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v177 - int32(1)
	F_errmsg_internal(m, int32(484647), v11)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(511735), int32(143), int32(28627))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
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
	var v109 int32
	_ = v109
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
	var v193 int32
	_ = v193
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
	var v266 int32
	_ = v266
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
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
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
	var v650 int32
	_ = v650
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
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
	v910 = m.ExcPending
	if v910 != 0 {
		goto L8
	} else {
		goto L254
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
	v54 = v14
	goto L5
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32))))
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
	v877 = v54 + int32(1)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v877 < v878 {
		v54 = v877
		goto L5
	} else {
		goto L253
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
		v109 = v65
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v118 = v109
	goto L10
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v75 < v74 {
		v109 = v65
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
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+(l5+v81))))
	v100 = v95 & (v97 ^ int32(-1))
	v102 = base.B2i32(v100 == int32(0))
	if v100 != 0 {
		v109 = v102
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v109 = v102
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
		v193 = v149
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v202 = v193
	goto L31
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v159 < v158 {
		v193 = v149
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
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+(v148+v165))))
	v184 = v179 & (v181 ^ int32(-1))
	v186 = base.B2i32(v184 == int32(0))
	if v184 != 0 {
		v193 = v186
		goto L35
	} else {
		goto L43
	}
L42:
	;
	v193 = v186
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
	F_errmsg_internal(m, int32(433135), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(516724), int32(2603), int32(161312))
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
	if v299 == v301 {
		v342 = v301
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
		v266 = v222
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v275 = v266
	goto L54
L59:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v232 < v231 {
		v266 = v222
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
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+(l6+v238))))
	v257 = v252 & (v254 ^ int32(-1))
	v259 = base.B2i32(v257 == int32(0))
	if v257 != 0 {
		v266 = v259
		goto L58
	} else {
		goto L66
	}
L65:
	;
	v266 = v259
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
	v490 = F_make_restrictinfo(m, l0, v62, v342^int32(1), l10, l11, v300, l4, v477, l8, l7)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L130
	}
L82:
	;
	if v342 != 0 {
		goto L96
	} else {
		goto L97
	}
L83:
	;
	goto L82
L84:
	;
	if l7 == int32(0) {
		v342 = v301
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v310 < v311 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v313 = v310
	goto L88
L87:
	;
	v313 = v311
	goto L88
L88:
	;
	if v313 <= int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v316 = int32(1)
	goto L91
L90:
	;
	v316 = v313
	goto L91
L91:
	;
	v317 = int32(8)
	v322 = int32(0)
	goto L92
L92:
	;
	v329 = v322 << (uint(int32(2)) % 32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l7+v317+v329)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329+(v299+v317))))
	v334 = v331 & v333
	v336 = base.B2i32(v334 != int32(0))
	if v334 != 0 {
		v342 = v336
		goto L83
	} else {
		goto L94
	}
L93:
	;
	v342 = v336
	goto L83
L94:
	;
	v338 = v322 + int32(1)
	if v338 != v316 {
		v322 = v338
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	if l12 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v353 = int32(0)
	if v62 == v353 {
		v382 = v353
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v477 = l6
	v487 = int32(0)
	goto L81
L100:
	;
	goto L101
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l12)))
	v350 = F_lappend(m, v349, v62)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v350
	goto L7
L103:
	;
	if v382 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L104:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	switch v356 - int32(52) {
	case 0:
		goto L107
	case 1:
		goto L106
	default:
		v382 = v353
		goto L103
	}
L105:
	;
	v382 = int32(0)
	goto L103
L106:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v368 != int32(4) {
		goto L105
	} else {
		goto L113
	}
L107:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v359 != 0 {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+12)))
	if v360 != 0 {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v361 == int32(0) {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if v364 != int32(6) {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	if v367 != 0 {
		goto L105
	} else {
		goto L112
	}
L112:
	;
	v382 = v361
	goto L103
L113:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v371 == int32(0) {
		goto L105
	} else {
		goto L114
	}
L114:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	if v374 != int32(6) {
		goto L105
	} else {
		goto L115
	}
L115:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v371)+28))
	if v377 == int32(0) {
		v382 = v371
		goto L103
	} else {
		goto L116
	}
L116:
	;
	goto L105
L117:
	;
	v477 = v299
	v487 = l9
	goto L81
L118:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v382)+24))
	if v385 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v388 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v391 = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v392 <= v391 {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v408 = v391
	v417 = v392
	goto L122
L122:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v388)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419+v408<<(uint(int32(2))%32))))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+20))
	if v424 != int32(5) {
		v435 = v417
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L117
L124:
	;
	v437 = v408 + int32(1)
	if v437 < v435 {
		v408 = v437
		v417 = v435
		goto L122
	} else {
		goto L129
	}
L125:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423)+24))
	if v427 == int32(0) {
		v435 = v417
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v382)+24))
	v431 = F_bms_is_member(m, v427, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	if v431 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	v435 = v433
	goto L124
L129:
	;
	goto L123
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v490
	if v477 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v539 == int32(2) {
		goto L147
	} else {
		goto L148
	}
L132:
	;
	v539 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v501 = int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	if v502 <= v501 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v505 = v501
	goto L137
L136:
	;
	v505 = v502
	goto L137
L137:
	;
	v508 = int32(0)
	v510 = v508
	v511 = v508
	goto L138
L138:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v477+int32(8)+v510<<(uint(int32(2))%32))))
	if v519 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v539 = v532
	goto L131
L140:
	;
	goto L139
L141:
	;
	v520 = int32(2)
	if v511 != 0 {
		v532 = v520
		goto L140
	} else {
		goto L144
	}
L142:
	;
	v525 = v511
	goto L143
L143:
	;
	v528 = v510 + int32(1)
	if v528 != v505 {
		v510 = v528
		v511 = v525
		goto L138
	} else {
		goto L146
	}
L144:
	;
	v521 = int32(1)
	if base.Ui32(v521) < base.Ui32(base.I32_popcnt(v519)) {
		v532 = v520
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v525 = v521
	goto L143
L146:
	;
	v532 = v525
	goto L140
L147:
	;
	v543 = F_pull_var_clause(m, v62, int32(26))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L8
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+10)))
	if v554 != 0 {
		goto L157
	} else {
		goto L158
	}
L150:
	;
	if l11 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v546 = F_bms_intersect(m, v477, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L8
	} else {
		goto L154
	}
L152:
	;
	v548 = v477
	goto L153
L153:
	;
	F_add_vars_to_targetlist(m, l0, v543, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L8
	} else {
		goto L155
	}
L154:
	;
	v548 = v546
	goto L153
L155:
	;
	F_list_free(m, v543)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	goto L149
L157:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v490)+96))
	if v583 == int32(0) {
		v849 = v490
		goto L169
	} else {
		goto L170
	}
L158:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	if v555 == int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v558 != int32(17) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v555)+28))
	if v561 == int32(0) {
		goto L157
	} else {
		goto L161
	}
L161:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v564 != int32(2) {
		goto L157
	} else {
		goto L162
	}
L162:
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
		goto L163
	}
L163:
	;
	v572 = F_op_mergejoinable(m, v567, v570)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L8
	} else {
		goto L164
	}
L164:
	;
	if v572 == int32(0) {
		goto L157
	} else {
		goto L165
	}
L165:
	;
	v576 = F_contain_volatile_functions(m, v490)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L8
	} else {
		goto L166
	}
L166:
	;
	if v576 != 0 {
		goto L157
	} else {
		goto L167
	}
L167:
	;
	v578 = F_get_mergejoin_opfamilies(m, v567)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L8
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+96)) = v578
	goto L157
L169:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L8
	} else {
		goto L252
	}
L170:
	;
	if v487 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v589 = F_process_equivalence(m, l0, v27+int32(12), v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L8
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v342 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	if v589 != 0 {
		goto L7
	} else {
		goto L175
	}
L175:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+96))
	if v592 == int32(0) {
		v849 = v591
		goto L169
	} else {
		goto L176
	}
L176:
	;
	F_initialize_mergeclause_eclasses(m, l0, v591)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L8
	} else {
		goto L177
	}
L177:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v849 = v597
	goto L169
L178:
	;
	F_initialize_mergeclause_eclasses(m, l0, v490)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L8
	} else {
		goto L251
	}
L179:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+9)))
	if v600 != int32(1) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	F_initialize_mergeclause_eclasses(m, l0, v490)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L8
	} else {
		goto L181
	}
L181:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v490)+44))
	v606 = int32(0)
	if v605 == v606 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v490)+48))
	v720 = int32(0)
	if v719 == v720 {
		goto L217
	} else {
		goto L218
	}
L183:
	;
	if v659 == int32(0) {
		goto L182
	} else {
		goto L197
	}
L184:
	;
	v659 = int32(1)
	goto L183
L185:
	;
	goto L186
L186:
	;
	if l7 == int32(0) {
		v650 = v606
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v659 = v650
	goto L183
L188:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v616 < v615 {
		v650 = v606
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v618 = int32(1)
	if v615 <= v618 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v621 = v618
	goto L192
L191:
	;
	v621 = v615
	goto L192
L192:
	;
	v622 = int32(8)
	v627 = int32(0)
	goto L193
L193:
	;
	v634 = v627 << (uint(int32(2)) % 32)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v605+v622+v634)))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v634+(l7+v622))))
	v641 = v636 & (v638 ^ int32(-1))
	v643 = base.B2i32(v641 == int32(0))
	if v641 != 0 {
		v650 = v643
		goto L187
	} else {
		goto L195
	}
L194:
	;
	v650 = v643
	goto L187
L195:
	;
	v645 = v627 + int32(1)
	if v645 != v621 {
		v627 = v645
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v490)+48))
	v663 = int32(0)
	if v662 == v663 {
		v704 = v663
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v704 != 0 {
		goto L182
	} else {
		goto L212
	}
L199:
	;
	goto L198
L200:
	;
	if l7 == int32(0) {
		v704 = v663
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v662)+4))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v672 < v673 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v675 = v672
	goto L204
L203:
	;
	v675 = v673
	goto L204
L204:
	;
	if v675 <= int32(1) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v678 = int32(1)
	goto L207
L206:
	;
	v678 = v675
	goto L207
L207:
	;
	v679 = int32(8)
	v684 = int32(0)
	goto L208
L208:
	;
	v691 = v684 << (uint(int32(2)) % 32)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l7+v679+v691)))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v691+(v662+v679))))
	v696 = v693 & v695
	v698 = base.B2i32(v696 != int32(0))
	if v696 != 0 {
		v704 = v698
		goto L199
	} else {
		goto L210
	}
L209:
	;
	v704 = v698
	goto L199
L210:
	;
	v700 = v684 + int32(1)
	if v700 != v678 {
		v684 = v700
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v709 = F_palloc0(m, int32(12))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L8
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v709)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v709)+4)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v709))) = int32(321)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v716 = F_lappend(m, v715, v709)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L8
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v716
	goto L7
L215:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v833 != int32(2) {
		v849 = v490
		goto L169
	} else {
		goto L248
	}
L216:
	;
	if v773 == int32(0) {
		goto L215
	} else {
		goto L230
	}
L217:
	;
	v773 = int32(1)
	goto L216
L218:
	;
	goto L219
L219:
	;
	if l7 == int32(0) {
		v764 = v720
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v773 = v764
	goto L216
L221:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v730 < v729 {
		v764 = v720
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v732 = int32(1)
	if v729 <= v732 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v735 = v732
	goto L225
L224:
	;
	v735 = v729
	goto L225
L225:
	;
	v736 = int32(8)
	v741 = int32(0)
	goto L226
L226:
	;
	v748 = v741 << (uint(int32(2)) % 32)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v719+v736+v748)))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v748+(l7+v736))))
	v755 = v750 & (v752 ^ int32(-1))
	v757 = base.B2i32(v755 == int32(0))
	if v755 != 0 {
		v764 = v757
		goto L220
	} else {
		goto L228
	}
L227:
	;
	v764 = v757
	goto L220
L228:
	;
	v759 = v741 + int32(1)
	if v759 != v735 {
		v741 = v759
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v490)+44))
	v777 = int32(0)
	if v776 == v777 {
		v818 = v777
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v818 != 0 {
		goto L215
	} else {
		goto L245
	}
L232:
	;
	goto L231
L233:
	;
	if l7 == int32(0) {
		v818 = v777
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v786 < v787 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v789 = v786
	goto L237
L236:
	;
	v789 = v787
	goto L237
L237:
	;
	if v789 <= int32(1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v792 = int32(1)
	goto L240
L239:
	;
	v792 = v789
	goto L240
L240:
	;
	v793 = int32(8)
	v798 = int32(0)
	goto L241
L241:
	;
	v805 = v798 << (uint(int32(2)) % 32)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l7+v793+v805)))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v805+(v776+v793))))
	v810 = v807 & v809
	v812 = base.B2i32(v810 != int32(0))
	if v810 != 0 {
		v818 = v812
		goto L232
	} else {
		goto L243
	}
L242:
	;
	v818 = v812
	goto L232
L243:
	;
	v814 = v798 + int32(1)
	if v814 != v792 {
		v798 = v814
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v823 = F_palloc0(m, int32(12))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L8
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v823)+4)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v823))) = int32(321)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v830 = F_lappend(m, v829, v823)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L8
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v830
	goto L7
L248:
	;
	v837 = F_palloc0(m, int32(12))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L8
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v837)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v837)+4)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v837))) = int32(321)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v844 = F_lappend(m, v843, v837)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L8
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v844
	goto L7
L251:
	;
	v849 = v490
	goto L169
L252:
	;
	goto L7
L253:
	;
	goto L6
L254:
	;
	F_errmsg_internal(m, int32(150876), int32(0))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L8
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(516724), int32(2611), int32(161312))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L8
	} else {
		goto L256
	}
L256:
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
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
	var v155 int32
	_ = v155
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
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
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
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
	v319 = m.ExcPending
	if v319 != 0 {
		goto L24
	} else {
		goto L93
	}
L4:
	;
	m.G0 = v9 + int32(16)
	return
L5:
	;
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v67 = int32(0)
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
	v34 = int32(-1)
	goto L13
L12:
	;
	v67 = v59
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
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12)))) = v51
	v59 = int32(1)
	goto L12
L15:
	;
	if int32(0) <= v34 {
		v59 = v14
		goto L12
	} else {
		goto L18
	}
L16:
	;
	v51 = v34
	goto L17
L17:
	;
	v53 = v31 + int32(1)
	if v53 != v26 {
		v31 = v53
		v34 = v51
		goto L13
	} else {
		goto L20
	}
L18:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v41)) {
		v59 = v14
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v51 = base.I32_ctz(v41) | v31<<(uint(int32(5))%32)
	goto L17
L20:
	;
	goto L14
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v69 = F_find_base_rel(m, l0, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v118 != 0 {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	return
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v68<<(uint(int32(2))%32))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+20)))
	if v76 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v69)+184))
	v110 = F_lappend(m, v109, v106)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L24
	} else {
		goto L37
	}
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+21)))
	if v79 != int32(112) {
		v106 = l1
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v82 = F_restriction_is_always_true(m, l0, l1)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L24
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v82 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v84 = F_restriction_is_always_false(m, l0, l1)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	if v84 == int32(0) {
		v106 = l1
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v90 = int32(0)
	v92 = F_makeBoolConst(m, v90, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v102 = F_make_restrictinfo(m, l0, v92, v94, v95, v96, v97, int32(0), v99, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v88
	v106 = v102
	goto L26
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+184)) = v110
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v69)+208))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if base.Ui32(v113) < base.Ui32(v114) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = v113
	goto L40
L39:
	;
	v116 = v114
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+208)) = v116
	goto L4
L41:
	;
	F_check_memoizable(m, l1)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L24
	} else {
		goto L52
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v119 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v122 != int32(17) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	if v125 == int32(0) {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v128 != int32(2) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v134 = F_exprType(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L24
	} else {
		goto L47
	}
L47:
	;
	v136 = F_op_hashjoinable(m, v131, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L24
	} else {
		goto L48
	}
L48:
	;
	if v136 == int32(0) {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v140 = F_contain_volatile_functions(m, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L24
	} else {
		goto L50
	}
L50:
	;
	if v140 != 0 {
		goto L41
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v131
	goto L41
L52:
	;
	v147 = F_restriction_is_always_true(m, l0, l1)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L24
	} else {
		goto L54
	}
L53:
	;
	goto L4
L54:
	;
	if v147 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v149 = F_restriction_is_always_false(m, l0, l1)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L24
	} else {
		goto L56
	}
L56:
	;
	if v149 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v153 = int32(0)
	v155 = F_makeBoolConst(m, v153, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L24
	} else {
		goto L60
	}
L58:
	;
	v169 = l1
	goto L59
L59:
	;
	if v11 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v165 = F_make_restrictinfo(m, l0, v155, v157, v158, v159, v160, int32(0), v162, v163, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L24
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+56)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v151
	v169 = v165
	goto L59
L62:
	;
	if v228 < int32(0) {
		goto L53
	} else {
		goto L73
	}
L63:
	;
	v228 = base.I32_ctz(v214) | v215<<(uint(int32(5))%32)
	goto L62
L64:
	;
	v228 = int32(-2)
	goto L62
L65:
	;
	v181 = base.I32_div_s(int32(0), int32(32))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v182 <= v181 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v185 = v11 + int32(8)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v181<<(uint(int32(2))%32))))
	v192 = v189 & int32(-1)
	if v192 != 0 {
		v214 = v192
		v215 = v181
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v194 = v181 + int32(1)
	if v194 == v182 {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v197 = v194
	goto L69
L69:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v185+v197<<(uint(int32(2))%32))))
	if v204 != 0 {
		v214 = v204
		v215 = v197
		goto L63
	} else {
		goto L71
	}
L70:
	;
	goto L64
L71:
	;
	v206 = v197 + int32(1)
	if v206 != v182 {
		v197 = v206
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v233 = v228
	goto L74
L74:
	;
	v237 = F_find_base_rel_ignore_join(m, l0, v233)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L24
	} else {
		goto L76
	}
L75:
	;
	goto L53
L76:
	;
	if v237 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v237)+212))
	v240 = F_lappend(m, v239, v169)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L24
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v11 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+212)) = v240
	goto L79
L81:
	;
	if int32(0) <= v298 {
		v233 = v298
		goto L74
	} else {
		goto L92
	}
L82:
	;
	v298 = base.I32_ctz(v284) | v285<<(uint(int32(5))%32)
	goto L81
L83:
	;
	v298 = int32(-2)
	goto L81
L84:
	;
	v249 = v233 + int32(1)
	v251 = base.I32_div_s(v249, int32(32))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v252 <= v251 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v255 = v11 + int32(8)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v251<<(uint(int32(2))%32))))
	v262 = v259 & (int32(-1) << (uint(v249) % 32))
	if v262 != 0 {
		v284 = v262
		v285 = v251
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v264 = v251 + int32(1)
	if v264 == v252 {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v267 = v264
	goto L88
L88:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255+v267<<(uint(int32(2))%32))))
	if v274 != 0 {
		v284 = v274
		v285 = v267
		goto L82
	} else {
		goto L90
	}
L89:
	;
	goto L83
L90:
	;
	v276 = v267 + int32(1)
	if v276 != v252 {
		v267 = v276
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	goto L75
L93:
	;
	F_errmsg_internal(m, int32(373269), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(516724), int32(3276), int32(161280))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_div_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v28 int64
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v155 int32
	_ = v155
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v197 int64
	_ = v197
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v211 int64
	_ = v211
	var v218 int64
	_ = v218
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v243 int64
	_ = v243
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v287 int64
	_ = v287
	var v294 int64
	_ = v294
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v347 int64
	_ = v347
	var v360 int64
	_ = v360
	var v361 int64
	_ = v361
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v370 int32
	_ = v370
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
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
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v634 int32
	_ = v634
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v682 int32
	_ = v682
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int64
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
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
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int64
	_ = v871
	var v872 int64
	_ = v872
	var v874 int64
	_ = v874
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int64
	_ = v886
	var v889 int64
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int64
	_ = v940
	var v943 int64
	_ = v943
	var v959 int32
	_ = v959
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int64
	_ = v985
	var v987 int64
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v996 int64
	_ = v996
	var v998 int64
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1141 int32
	_ = v1141
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1193 float64
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 float64
	_ = v1199
	var v1204 float64
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1209 int64
	_ = v1209
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1240 int64
	_ = v1240
	var v1241 int64
	_ = v1241
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int64
	_ = v1254
	var v1255 float64
	_ = v1255
	var v1257 float64
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1280 int64
	_ = v1280
	var v1283 int64
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1294 int32
	_ = v1294
	var v1316 int64
	_ = v1316
	var v1324 int32
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1329 int64
	_ = v1329
	var v1332 int64
	_ = v1332
	var v1334 int64
	_ = v1334
	var v1342 int64
	_ = v1342
	var v1346 int64
	_ = v1346
	var v1347 int64
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1352 int64
	_ = v1352
	var v1354 int64
	_ = v1354
	var v1383 int64
	_ = v1383
	var v1384 int64
	_ = v1384
	var v1385 int64
	_ = v1385
	var v1388 float64
	_ = v1388
	var v1389 int64
	_ = v1389
	var v1395 float64
	_ = v1395
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1413 int64
	_ = v1413
	var v1421 int32
	_ = v1421
	var v1443 int64
	_ = v1443
	var v1444 int64
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1449 int64
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1502 int64
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 int64
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int64
	_ = v1515
	var v1519 int64
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1534 int32
	_ = v1534
	var v1564 int32
	_ = v1564
	var v1565 int64
	_ = v1565
	var v1569 int64
	_ = v1569
	var v1574 int64
	_ = v1574
	var v1575 int64
	_ = v1575
	var v1603 int64
	_ = v1603
	var v1604 int64
	_ = v1604
	var v1605 int64
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1612 int64
	_ = v1612
	var v1624 int32
	_ = v1624
	var v1656 int32
	_ = v1656
	var v1659 int64
	_ = v1659
	var v1672 int32
	_ = v1672
	var v1694 int64
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1702 int64
	_ = v1702
	var v1703 int64
	_ = v1703
	var v1706 int64
	_ = v1706
	var v1709 int64
	_ = v1709
	var v1711 int64
	_ = v1711
	var v1719 int64
	_ = v1719
	var v1723 int64
	_ = v1723
	var v1724 int64
	_ = v1724
	var v1731 int64
	_ = v1731
	var v1763 int64
	_ = v1763
	var v1773 int32
	_ = v1773
	var v1803 int64
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1818 int32
	_ = v1818
	var v1849 int64
	_ = v1849
	var v1853 int64
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1860 int64
	_ = v1860
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1893 int64
	_ = v1893
	var v1902 int32
	_ = v1902
	var v1903 int64
	_ = v1903
	var v1907 int64
	_ = v1907
	var v1909 int64
	_ = v1909
	var v1914 int64
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1921 int64
	_ = v1921
	var v1925 int64
	_ = v1925
	var v1929 int64
	_ = v1929
	var v1934 int64
	_ = v1934
	var v1937 int64
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1949 int32
	_ = v1949
	var v1970 int64
	_ = v1970
	var v1979 int32
	_ = v1979
	var v1980 int64
	_ = v1980
	var v1984 int64
	_ = v1984
	var v1986 int64
	_ = v1986
	var v1991 int64
	_ = v1991
	var v1996 int64
	_ = v1996
	var v1997 int64
	_ = v1997
	var v2025 int64
	_ = v2025
	var v2027 int64
	_ = v2027
	var v2033 int64
	_ = v2033
	var v2034 int64
	_ = v2034
	var v2036 int64
	_ = v2036
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2082 int32
	_ = v2082
	var v2103 int64
	_ = v2103
	var v2112 int32
	_ = v2112
	var v2116 int64
	_ = v2116
	var v2117 int64
	_ = v2117
	var v2119 int64
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int64
	_ = v2124
	var v2126 int64
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2131 int64
	_ = v2131
	var v2132 int64
	_ = v2132
	var v2134 int64
	_ = v2134
	var v2136 int64
	_ = v2136
	var v2140 int64
	_ = v2140
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2215 int32
	_ = v2215
	var v2232 int64
	_ = v2232
	var v2239 int32
	_ = v2239
	var v2243 int64
	_ = v2243
	var v2244 int64
	_ = v2244
	var v2247 int64
	_ = v2247
	var v2250 int64
	_ = v2250
	var v2252 int64
	_ = v2252
	var v2260 int64
	_ = v2260
	var v2264 int64
	_ = v2264
	var v2265 int64
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2314 int32
	_ = v2314
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2378 int32
	_ = v2378
	var v2382 int32
	_ = v2382
	var v2388 int32
	_ = v2388
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2445 int32
	_ = v2445
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2488 int32
	_ = v2488
	var v2494 int32
	_ = v2494
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2541 int32
	_ = v2541
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2614 int32
	_ = v2614
	var v2620 int32
	_ = v2620
	v28 = int64(0)
	v35 = m.G0
	v37 = v35 - int32(48)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v39 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v37 + int32(48)
	return
L2:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v2180 != 0 {
		goto L287
	} else {
		goto L288
	}
L3:
	;
	v2075 = v1656 - int32(8)
	v2082 = v1015
	v2103 = v1731
	goto L280
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L17
	} else {
		goto L276
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L17
	} else {
		goto L272
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42))))
	if v43 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v39 <= int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = base.I32_extend16_s(v43)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v39 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v39) <= base.Ui32(int32(4)) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+2)))
	v58 = v54 + v48*int32(10000)
	v59 = v49 - int32(1)
	goto L13
L12:
	;
	v58 = v48
	v59 = v49
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v62 == int32(16384) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v65 = int32(0) - v58
	goto L16
L15:
	;
	v65 = v58
	goto L16
L16:
	;
	F_div_var_int(m, l0, v65, v59, l2, l3, l4)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v71 = int64(*(*int16)(unsafe.Add(mBase, uint32(v42)+4)))
	v72 = int64(*(*int16)(unsafe.Add(mBase, uint32(v42)+2)))
	v75 = int64(10000)
	v80 = v71 + (v72+base.I64_extend16_s(base.I64_extend_i32_u(v43))*v75)*v75
	if v39 != int32(3) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L22:
	;
	v83 = int64(*(*int16)(unsafe.Add(mBase, uint32(v42)+6)))
	v87 = v83 + v80*int64(10000)
	goto L24
L23:
	;
	v87 = v80
	goto L24
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v90 == int32(16384) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v93 = int64(0) - v87
	goto L27
L26:
	;
	v93 = v87
	goto L27
L27:
	;
	if v93 == int64(0) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v68 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v98 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v110 = int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = v111 + (v39 + (v112 ^ int32(-1)))
	v120 = base.I32_div_s(l3+int32(3), int32(4))
	v123 = v116 + v120 + v110
	if v123 <= v110 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_pfree(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
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
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v104
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v104
	goto L1
L35:
	;
	goto L34
L36:
	;
	v126 = v110
	goto L38
L37:
	;
	v126 = v123
	goto L38
L38:
	;
	v127 = v126 + l4
	v132 = F_palloc(m, v127<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	v134 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v134)
	v137 = v132 + int32(2)
	v144 = v87 >> (uint(int64(63)) % 64)
	v146 = v87 ^ v144 - v144
	if base.Ui64(int64(1844674407370956)) <= base.Ui64(v146) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if base.B2i32(v109 == v134)^base.B2i32(int64(0) < v93) != 0 {
		goto L61
	} else {
		goto L62
	}
L41:
	;
	if v127 <= int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v127 <= int32(0) {
		goto L40
	} else {
		goto L54
	}
L44:
	;
	v155 = int32(0)
	v183 = int64(0)
	v184 = v28
	goto L45
L45:
	;
	v190 = v37 + int32(32)
	v191 = int64(10000)
	v192 = int64(0)
	v197 = int64(32)
	v200 = int64(base.Ui64(v183) >> (uint(v197) % 64))
	v203 = int64(4294967295)
	v206 = v183 & v203
	v207 = v191 * v206
	v211 = int64(base.Ui64(v207)>>(uint(v197)%64)) + v191*v200
	v218 = v206*v192 + v211&v203
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v183*v192 + v184*v191 + v192*v200 + int64(base.Ui64(v211)>>(uint(v197)%64)) + int64(base.Ui64(v218)>>(uint(v197)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v207&v203 | v218<<(uint(v197)%64)
	goto L47
L46:
	;
	goto L40
L47:
	;
	v231 = v37 + int32(16)
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v37)+32))
	if v155 < v68 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v237 = int32(*(*int16)(unsafe.Add(mBase, uint32(v108+v155<<(uint(int32(1))%32)))))
	v238 = v237
	goto L50
L49:
	;
	v238 = int32(0)
	goto L50
L50:
	;
	v239 = base.I64_extend_i32_s(v238)
	v240 = v232 + v239
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(40))))
	v247 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v240) < base.Ui64(v232))) + (v243 + v239>>(uint(int64(63))%64))
	v249 = m.G0
	v250 = int32(16)
	v251 = v249 - v250
	m.G0 = v251
	F___udivmodti4(m, v251, v240, v247, v146, int64(0))
	mBase = m.M
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v251)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v231))) = v255
	m.G0 = v251 + v250
	goto L51
L51:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v137+v155<<(uint(int32(1))%32)))) = uint16(v265)
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v37)+24))
	v273 = int64(32)
	v274 = int64(base.Ui64(v146) >> (uint(v273) % 64))
	v276 = int64(base.Ui64(v265) >> (uint(v273) % 64))
	v279 = int64(4294967295)
	v280 = v146 & v279
	v282 = v265 & v279
	v283 = v280 * v282
	v287 = int64(base.Ui64(v283)>>(uint(v273)%64)) + v280*v276
	v294 = v282*v274 + v287&v279
	*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v265*int64(0) + v267*v146 + v274*v276 + int64(base.Ui64(v287)>>(uint(v273)%64)) + int64(base.Ui64(v294)>>(uint(v273)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v37))) = v283&v279 | v294<<(uint(v273)%64)
	goto L52
L52:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v313 = v155 + int32(1)
	if v313 != v127 {
		v155 = v313
		v183 = v240 - v307
		v184 = v247 - v305 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v240) < base.Ui64(v307)))
		goto L45
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	v318 = int32(0)
	v347 = v28
	goto L55
L55:
	;
	if v318 < v68 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L40
L57:
	;
	v360 = int64(*(*int16)(unsafe.Add(mBase, uint32(v108+v318<<(uint(int32(1))%32)))))
	v361 = v360
	goto L59
L58:
	;
	v361 = int64(0)
	goto L59
L59:
	;
	v364 = v361 + v347*int64(10000)
	v365 = base.I64_div_u_s(v364, v146)
	*(*uint16)(unsafe.Add(mBase, uint32(v137+v318<<(uint(int32(1))%32)))) = uint16(v365)
	v370 = v318 + int32(1)
	if v370 != v127 {
		v318 = v370
		v347 = v364 - v146*v365
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	v408 = int32(16384)
	goto L63
L62:
	;
	v408 = int32(0)
	goto L63
L63:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v409 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L17
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v127
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v116
	if l4 != 0 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v713
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v719
	goto L1
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v713 = int32(0)
	v719 = v682
	goto L68
L70:
	;
	if int32(0) < v579 {
		goto L105
	} else {
		goto L106
	}
L71:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v579 = v578
	v581 = v577
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v426 = l3 + v423<<(uint(int32(2))%32)
	if v426+int32(4) < int32(0) {
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
	v545 = l3 + v116<<(uint(int32(2))%32)
	if v545+int32(4) <= int32(0) {
		v682 = v137
		goto L69
	} else {
		goto L102
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
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v437 = l3 & int32(3)
	v441 = base.I32_div_s(v426+int32(7), int32(4))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v442 <= v441 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	goto L75
L80:
	;
	if int32(0) <= v508 {
		goto L79
	} else {
		goto L101
	}
L81:
	;
	v488 = v482
	goto L95
L82:
	;
	v455 = int32(1)
	v456 = v441 - v455
	v459 = v435 + v456<<(uint(v455)%32)
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v459))))
	v461 = int32(2)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v437<<(uint(v461)%32))+uint32(_consts[1070])))
	v466 = base.I32_rem_s(v460, v465)
	v467 = v460 - v466
	*(*uint16)(unsafe.Add(mBase, uint32(v459))) = uint16(v467)
	v470 = base.I32_div_s(v465, v461)
	if v466 < v470 {
		v508 = v456
		goto L80
	} else {
		goto L90
	}
L83:
	;
	if v437 == int32(0) {
		goto L79
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v441
	if v437 != 0 {
		goto L82
	} else {
		goto L88
	}
L86:
	;
	if v441 != v442 {
		goto L79
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v441
	goto L82
L88:
	;
	v452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v435+v441<<(uint(int32(1))%32)))))
	if v452 <= int32(4999) {
		v508 = v441
		goto L80
	} else {
		goto L89
	}
L89:
	;
	v482 = v441
	goto L81
L90:
	;
	v473 = v465 + base.I32_extend16_s(v467)
	if int32(9999) < v473 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v478 = v473 + int32(55536)
	goto L93
L92:
	;
	v478 = v473
	goto L93
L93:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v459))) = uint16(v478)
	if v473 < int32(10000) {
		v508 = v456
		goto L80
	} else {
		goto L94
	}
L94:
	;
	v482 = v456
	goto L81
L95:
	;
	v494 = int32(1)
	v495 = v488 - v494
	v498 = v435 + v495<<(uint(v494)%32)
	v501 = int32(*(*int16)(unsafe.Add(mBase, uint32(v498))))
	v503 = base.B2i32(int32(9998) < v501)
	if int32(9998) < v501 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v508 = v495
	goto L80
L97:
	;
	v504 = int32(-9999)
	goto L99
L98:
	;
	v504 = v494
	goto L99
L99:
	;
	v505 = v504 + v501
	*(*uint16)(unsafe.Add(mBase, uint32(v498))) = uint16(v505)
	if int32(9998) < v501 {
		v488 = v495
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v516 - int32(2)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v521 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v520 + v521
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v524 + v521
	goto L79
L102:
	;
	v553 = base.I32_div_s(v545+int32(7), int32(4))
	if v126 < v553 {
		goto L71
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v553
	v557 = l3 & int32(3)
	if v557 == int32(0) {
		v579 = v553
		v581 = v137
		goto L70
	} else {
		goto L104
	}
L104:
	;
	v563 = int32(2)
	v564 = v137 + v553<<(uint(int32(1))%32) - v563
	v565 = int32(*(*int16)(unsafe.Add(mBase, uint32(v564))))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v557<<(uint(v563)%32))+uint32(_consts[1070])))
	v571 = base.I32_rem_s(v565, v570)
	v572 = v565 - v571
	*(*uint16)(unsafe.Add(mBase, uint32(v564))) = uint16(v572)
	goto L71
L105:
	;
	v587 = v579
	v593 = v581
	goto L109
L106:
	;
	goto L107
L107:
	;
	if v579 != 0 {
		v713 = v579
		v719 = v581
		goto L68
	} else {
		goto L117
	}
L108:
	;
	v634 = v587
	goto L113
L109:
	;
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v593))))
	if v621 != 0 {
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v682 = v581 + v579<<(uint(int32(1))%32)
	goto L69
L111:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v623 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v622 - v623
	if v623 < v587 {
		v587 = v587 - v623
		v593 = v593 + int32(2)
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v671 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v593-int32(2)+v634<<(uint(int32(1))%32)))))
	if v671 != 0 {
		v713 = v634
		v719 = v593
		goto L68
	} else {
		goto L115
	}
L114:
	;
	v682 = v593
	goto L69
L115:
	;
	v672 = int32(1)
	if v672 < v634 {
		v634 = v634 - v672
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v682 = v581
	goto L69
L118:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v751 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v761 = int32(1)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v764 = v762 - v763
	v768 = base.I32_div_s(l3+int32(3), int32(4))
	v771 = v764 + v768 + int32(2)
	if v771 <= v761 {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	F_pfree(m, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
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
	v757 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v757
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v757
	goto L1
L124:
	;
	goto L123
L125:
	;
	v774 = v761
	goto L127
L126:
	;
	v774 = v771
	goto L127
L127:
	;
	v775 = v774 + l4
	v780 = l5 | base.B2i32(base.Ui32(v39) < base.Ui32(int32(13)))
	if v780 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v781 = v775
	goto L130
L129:
	;
	v781 = v775 + int32(4)
	goto L130
L130:
	;
	v782 = int32(1)
	v783 = v781 + v782
	v784 = int32(2)
	v785 = base.I32_div_s(v783, v784)
	v789 = base.I32_div_s(v39+v782, v784)
	v793 = base.I32_div_s(v68+v782, v784)
	if v780 == v782 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v809 = v804 << (uint(int32(3)) % 32)
	v815 = F_palloc(m, v809+v803<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L17
	} else {
		goto L144
	}
L132:
	;
	v796 = v789 + v785
	if v793 < v796 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	if v789 < v785 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v798 = v793
	goto L137
L136:
	;
	v798 = v796
	goto L137
L137:
	;
	v803 = v789
	v804 = v796
	v805 = v798
	goto L131
L138:
	;
	v800 = v789
	goto L140
L139:
	;
	v800 = v785
	goto L140
L140:
	;
	if v793 < v785 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v802 = v793
	goto L143
L142:
	;
	v802 = v785
	goto L143
L143:
	;
	v803 = v800
	v804 = v785
	v805 = v802
	goto L131
L144:
	;
	v818 = v805 - int32(1)
	if v818 <= int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v980 = v815 + v809
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v985 = int64(*(*int16)(unsafe.Add(mBase, uint32(v981+v959<<(uint(int32(2))%32)))))
	v987 = v985 * int64(10000)
	v988 = int32(1)
	v991 = v959<<(uint(v988)%32) | v988
	if v991 < v68 {
		goto L156
	} else {
		goto L157
	}
L146:
	;
	v959 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v824 = int32(0)
	if v805 != int32(2) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v836 = v824
	v839 = int32(0)
	goto L152
L150:
	;
	v903 = v824
	goto L151
L151:
	;
	if v818&int32(1) == int32(0) {
		v959 = v818
		goto L145
	} else {
		goto L155
	}
L152:
	;
	v864 = int32(3)
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v868 = int32(2)
	v870 = v867 + v836<<(uint(v868)%32)
	v871 = int64(*(*int16)(unsafe.Add(mBase, uint32(v870))))
	v872 = int64(10000)
	v874 = int64(*(*int16)(unsafe.Add(mBase, uint32(v870)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v815+v836<<(uint(v864)%32)))) = v871*v872 + v874
	v878 = v836 | int32(1)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v885 = v882 + v878<<(uint(v868)%32)
	v886 = int64(*(*int16)(unsafe.Add(mBase, uint32(v885))))
	v889 = int64(*(*int16)(unsafe.Add(mBase, uint32(v885)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v815+v878<<(uint(v864)%32)))) = v886*v872 + v889
	v893 = v836 + v868
	v895 = v839 + v868
	if v895 != v818&int32(2147483646) {
		v836 = v893
		v839 = v895
		goto L152
	} else {
		goto L154
	}
L153:
	;
	v903 = v893
	goto L151
L154:
	;
	goto L153
L155:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v939 = v936 + v903<<(uint(int32(2))%32)
	v940 = int64(*(*int16)(unsafe.Add(mBase, uint32(v939))))
	v943 = int64(*(*int16)(unsafe.Add(mBase, uint32(v939)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v815+v903<<(uint(int32(3))%32)))) = v940*int64(10000) + v943
	v959 = v818
	goto L145
L156:
	;
	v996 = int64(*(*int16)(unsafe.Add(mBase, uint32(v981+v991<<(uint(int32(1))%32)))))
	v998 = v987 + v996
	goto L158
L157:
	;
	v998 = v987
	goto L158
L158:
	;
	v999 = int32(8)
	v1000 = v980 + v999
	v1001 = int32(3)
	v1003 = v815 + v959<<(uint(v1001)%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1003))) = v998
	v1005 = int32(0)
	v1013 = F__emscripten_memset_bulkmem(m, v1003+v999, base.I32_extend8_s(v1005), (v804-v959)<<(uint(v1001)%32))
	mBase = m.M
	goto L159
L159:
	;
	v1015 = v803 - int32(1)
	if v803 < int32(2) {
		v1141 = v1005
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1170 = v1141 << (uint(int32(2)) % 32)
	v1172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1168+v1170))))
	v1174 = v1172 * int32(10000)
	v1178 = int32(1)
	v1181 = v1141<<(uint(v1178)%32) | v1178
	if v1181 < v39 {
		goto L169
	} else {
		goto L170
	}
L161:
	;
	v1020 = int32(0)
	if v803 != int32(2) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1032 = v1020
	v1035 = int32(0)
	goto L165
L163:
	;
	v1093 = v1020
	goto L164
L164:
	;
	if v1015&int32(1) == int32(0) {
		v1141 = v1015
		goto L160
	} else {
		goto L168
	}
L165:
	;
	v1060 = int32(2)
	v1061 = v1032 << (uint(v1060) % 32)
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1064 = v1063 + v1061
	v1065 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1064))))
	v1066 = int32(10000)
	v1068 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1064)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1000+v1061))) = v1065*v1066 + v1068
	v1072 = v1061 | int32(4)
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1075 = v1074 + v1072
	v1076 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1075))))
	v1079 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1075)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1000+v1072))) = v1076*v1066 + v1079
	v1083 = v1032 + v1060
	v1085 = v1035 + v1060
	if v1085 != v1015&int32(-2) {
		v1032 = v1083
		v1035 = v1085
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v1093 = v1083
	goto L164
L167:
	;
	goto L166
L168:
	;
	v1124 = v1093 << (uint(int32(2)) % 32)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1127 = v1126 + v1124
	v1128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1127))))
	v1131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1127)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1000+v1124))) = v1128*int32(10000) + v1131
	v1141 = v1015
	goto L160
L169:
	;
	v1186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1168+v1181<<(uint(int32(1))%32)))))
	v1188 = v1174 + v1186
	goto L171
L170:
	;
	v1188 = v1174
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1000+v1170))) = v1188
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	v1193 = base.F64_mul(base.F64_convert_i32_s(v1190), float64(1e+08))
	if int32(2) <= v803 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v980)+12))
	v1199 = base.F64_add(v1193, base.F64_convert_i32_s(v1196))
	goto L174
L173:
	;
	v1199 = v1193
	goto L174
L174:
	;
	if int32(2) <= v783 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v1204 = base.F64_div(float64(1), v1199)
	v1206 = v804 - int32(1)
	v1209 = *(*int64)(unsafe.Add(mBase, uint32(v815)))
	v1225 = v804
	v1226 = int32(0)
	v1240 = int64(1)
	v1241 = v1209
	goto L178
L176:
	;
	v1624 = int32(0)
	goto L177
L177:
	;
	if v780 == int32(0) {
		goto L2
	} else {
		goto L226
	}
L178:
	;
	v1250 = v1226 + int32(1)
	v1251 = int32(3)
	v1253 = v815 + v1250<<(uint(v1251)%32)
	v1254 = *(*int64)(unsafe.Add(mBase, uint32(v1253)))
	v1255 = base.F64_convert_i64_s(v1254)
	v1257 = base.F64_mul(v1204, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v1241), float64(1e+08)), v1255))
	v1264 = v815 + v1226<<(uint(v1251)%32)
	if base.F64_lt(base.F64_abs(v1257), float64(2.147483648e+09)) != 0 {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v1624 = v785
	goto L177
L180:
	;
	v1612 = v1606 + v1605*int64(100000000)
	*(*int64)(unsafe.Add(mBase, uint32(v1253))) = v1612
	*(*int64)(unsafe.Add(mBase, uint32(v1264))) = v1603
	if v785 != v1250 {
		v1225 = v1225 - int32(1)
		v1226 = v1250
		v1240 = v1604
		v1241 = v1612
		goto L178
	} else {
		goto L225
	}
L181:
	;
	v1271 = v1270 - base.B2i32(base.F64_ge(v1257, float64(0)) == int32(0))
	if v1271 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v1268 = base.I32_trunc_f64_s(v1257)
	v1270 = v1268
	goto L181
L183:
	;
	goto L184
L184:
	;
	v1270 = int32(-2147483648)
	goto L181
L185:
	;
	v1603 = int64(0)
	v1604 = v1240
	v1605 = v1241
	v1606 = v1254
	goto L180
L186:
	;
	goto L187
L187:
	;
	v1276 = v1271 >> (uint(int32(31)) % 32)
	v1280 = v1240 + base.I64_extend_i32_u(v1271^v1276-v1276)
	if v1280 < int64(92233720369) {
		v1421 = v1271
		v1443 = v1280
		v1444 = v1241
		v1445 = v1254
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1449 = base.I64_extend_i32_s(v1421)
	v1450 = v804 - v1226
	if v803 < v1450 {
		goto L209
	} else {
		goto L210
	}
L189:
	;
	v1283 = int64(0)
	v1284 = v1226 + (v803 - int32(2))
	if v1284 < v1206 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1286 = v1284
	goto L192
L191:
	;
	v1286 = v1206
	goto L192
L192:
	;
	if v1226 < v1286 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1294 = v1286
	v1316 = v1283
	goto L196
L194:
	;
	v1383 = v1283
	v1384 = v1241
	v1385 = v1254
	v1388 = v1255
	goto L195
L195:
	;
	v1389 = v1383 + v1384
	*(*int64)(unsafe.Add(mBase, uint32(v1264))) = v1389
	v1395 = base.F64_mul(v1204, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v1389), float64(1e+08)), v1388))
	if base.F64_lt(base.F64_abs(v1395), float64(2.147483648e+09)) != 0 {
		goto L205
	} else {
		goto L206
	}
L196:
	;
	v1324 = v815 + v1294<<(uint(int32(3))%32)
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(v1324)))
	v1326 = v1325 + v1316
	if v1326 < int64(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v1352 = *(*int64)(unsafe.Add(mBase, uint32(v1253)))
	v1354 = *(*int64)(unsafe.Add(mBase, uint32(v1264)))
	v1383 = v1347
	v1384 = v1354
	v1385 = v1352
	v1388 = base.F64_convert_i64_s(v1352)
	goto L195
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1324))) = v1346
	v1350 = v1294 - int32(1)
	if v1226 < v1350 {
		v1294 = v1350
		v1316 = v1347
		goto L196
	} else {
		goto L203
	}
L199:
	;
	v1329 = int64(-1)
	v1332 = base.I64_div_u_s(v1326^v1329, int64(100000000))
	v1334 = v1332 ^ v1329
	v1346 = v1334*int64(-100000000) + v1326
	v1347 = v1334
	goto L198
L200:
	;
	goto L201
L201:
	;
	if base.Ui64(v1326) < base.Ui64(int64(100000000)) {
		v1346 = v1326
		v1347 = int64(0)
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v1342 = base.I64_div_u_s(v1326, int64(100000000))
	v1346 = v1342*int64(-100000000) + v1326
	v1347 = v1342
	goto L198
L203:
	;
	goto L197
L204:
	;
	v1406 = v1405 - base.B2i32(base.F64_ge(v1395, float64(0)) == int32(0))
	v1408 = v1406 >> (uint(int32(31)) % 32)
	v1413 = base.I64_extend_i32_u(v1406 ^ v1408 - v1408 + int32(1))
	if v1406 != 0 {
		v1421 = v1406
		v1443 = v1413
		v1444 = v1389
		v1445 = v1385
		goto L188
	} else {
		goto L208
	}
L205:
	;
	v1403 = base.I32_trunc_f64_s(v1395)
	v1405 = v1403
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1405 = int32(-2147483648)
	goto L204
L208:
	;
	v1603 = int64(0)
	v1604 = v1413
	v1605 = v1389
	v1606 = v1385
	goto L180
L209:
	;
	v1452 = v803
	goto L211
L210:
	;
	v1452 = v1450
	goto L211
L211:
	;
	if v1452 <= int32(0) {
		v1603 = v1449
		v1604 = v1443
		v1605 = v1444
		v1606 = v1445
		goto L180
	} else {
		goto L212
	}
L212:
	;
	if v803 < v1225 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1456 = v803
	goto L215
L214:
	;
	v1456 = v1225
	goto L215
L215:
	;
	v1457 = int32(1)
	v1459 = int32(0)
	if v1456 != v1457 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1471 = v1459
	v1472 = int32(0)
	goto L219
L217:
	;
	v1534 = v1459
	goto L218
L218:
	;
	if v1456&v1457 != 0 {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1499 = int32(3)
	v1501 = v1264 + v1471<<(uint(v1499)%32)
	v1502 = *(*int64)(unsafe.Add(mBase, uint32(v1501)))
	v1503 = int32(2)
	v1506 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1471<<(uint(v1503)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1501))) = v1502 - v1506*v1449
	v1511 = v1471 | int32(1)
	v1514 = v1264 + v1511<<(uint(v1499)%32)
	v1515 = *(*int64)(unsafe.Add(mBase, uint32(v1514)))
	v1519 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1511<<(uint(v1503)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1514))) = v1515 - v1519*v1449
	v1524 = v1471 + v1503
	v1526 = v1472 + v1503
	if v1526 != v1456&int32(-2) {
		v1471 = v1524
		v1472 = v1526
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v1534 = v1524
	goto L218
L221:
	;
	goto L220
L222:
	;
	v1564 = v1264 + v1534<<(uint(int32(3))%32)
	v1565 = *(*int64)(unsafe.Add(mBase, uint32(v1564)))
	v1569 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1534<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1564))) = v1565 - v1569*v1449
	goto L224
L223:
	;
	goto L224
L224:
	;
	v1574 = *(*int64)(unsafe.Add(mBase, uint32(v1253)))
	v1575 = *(*int64)(unsafe.Add(mBase, uint32(v1264)))
	v1603 = v1449
	v1604 = v1443
	v1605 = v1575
	v1606 = v1574
	goto L180
L225:
	;
	goto L179
L226:
	;
	v1656 = v815 + v1624<<(uint(int32(3))%32)
	if v803 <= int32(1) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1773 = v1656 - int32(8)
	v1803 = v1763
	goto L240
L228:
	;
	v1659 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1656))) = v1659
	v1763 = v1659
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1672 = v803 - int32(2)
	v1694 = int64(0)
	goto L231
L231:
	;
	v1701 = v1656 + v1672<<(uint(int32(3))%32)
	v1702 = *(*int64)(unsafe.Add(mBase, uint32(v1701)))
	v1703 = v1702 + v1694
	if v1703 < int64(0) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1656))) = v1724
	v1731 = int64(0)
	if v1724 < v1731 {
		goto L3
	} else {
		goto L239
	}
L233:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1701)+8)) = v1723
	if int32(0) < v1672 {
		v1672 = v1672 - int32(1)
		v1694 = v1724
		goto L231
	} else {
		goto L238
	}
L234:
	;
	v1706 = int64(-1)
	v1709 = base.I64_div_u_s(v1703^v1706, int64(100000000))
	v1711 = v1709 ^ v1706
	v1723 = v1711*int64(-100000000) + v1703
	v1724 = v1711
	goto L233
L235:
	;
	goto L236
L236:
	;
	if base.Ui64(v1703) < base.Ui64(int64(100000000)) {
		v1723 = v1703
		v1724 = int64(0)
		goto L233
	} else {
		goto L237
	}
L237:
	;
	v1719 = base.I64_div_u_s(v1703, int64(100000000))
	v1723 = v1719*int64(-100000000) + v1703
	v1724 = v1719
	goto L233
L238:
	;
	goto L232
L239:
	;
	v1763 = v1724
	goto L227
L240:
	;
	v1808 = int32(0)
	if v803 <= v1808 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v2033 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000))))
	v2034 = v2025 + v2027 - v2033
	*(*int64)(unsafe.Add(mBase, uint32(v1656))) = v2034
	v2036 = *(*int64)(unsafe.Add(mBase, uint32(v1773)))
	*(*int64)(unsafe.Add(mBase, uint32(v1773))) = v2036 + int64(1)
	v1803 = v2034
	goto L240
L243:
	;
	v2025 = int64(0)
	v2027 = v1803
	goto L242
L244:
	;
	goto L245
L245:
	;
	v1818 = v1808
	goto L246
L246:
	;
	v1849 = *(*int64)(unsafe.Add(mBase, uint32(v1656+v1818<<(uint(int32(3))%32))))
	v1853 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1818<<(uint(int32(2))%32)))))
	if v1849 < v1853 {
		goto L2
	} else {
		goto L248
	}
L247:
	;
	v1860 = int64(0)
	if v803 < int32(2) {
		v2025 = v1860
		v2027 = v1803
		goto L242
	} else {
		goto L253
	}
L248:
	;
	if v1849 <= v1853 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1857 = v1818 + int32(1)
	if v1857 < v803 {
		v1818 = v1857
		goto L246
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	goto L247
L252:
	;
	goto L251
L253:
	;
	if v803 != int32(2) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1872 = v1015
	v1873 = int32(0)
	v1893 = v1860
	goto L257
L255:
	;
	v1949 = v1015
	v1970 = v1860
	goto L256
L256:
	;
	if v1015&int32(1) != 0 {
		goto L266
	} else {
		goto L267
	}
L257:
	;
	v1902 = v1656 + v1872<<(uint(int32(3))%32)
	v1903 = *(*int64)(unsafe.Add(mBase, uint32(v1902)))
	v1907 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1872<<(uint(int32(2))%32)))))
	v1909 = v1903 - v1907 + v1893
	if v1909 < int64(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1949 = v1939
	v1970 = v1937
	goto L256
L259:
	;
	v1914 = v1909 + int64(100000000)
	goto L261
L260:
	;
	v1914 = v1909
	goto L261
L261:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1902))) = v1914
	v1917 = v1872 - int32(1)
	v1920 = v1656 + v1917<<(uint(int32(3))%32)
	v1921 = *(*int64)(unsafe.Add(mBase, uint32(v1920)))
	v1925 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1917<<(uint(int32(2))%32)))))
	v1929 = v1921 - v1925 + v1909>>(uint(int64(63))%64)
	if v1929 < int64(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1934 = v1929 + int64(100000000)
	goto L264
L263:
	;
	v1934 = v1929
	goto L264
L264:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1920))) = v1934
	v1937 = v1929 >> (uint(int64(63)) % 64)
	v1938 = int32(2)
	v1939 = v1872 - v1938
	v1941 = v1873 + v1938
	if v1941 != v1015&int32(-2) {
		v1872 = v1939
		v1873 = v1941
		v1893 = v1937
		goto L257
	} else {
		goto L265
	}
L265:
	;
	goto L258
L266:
	;
	v1979 = v1656 + v1949<<(uint(int32(3))%32)
	v1980 = *(*int64)(unsafe.Add(mBase, uint32(v1979)))
	v1984 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v1949<<(uint(int32(2))%32)))))
	v1986 = v1980 - v1984 + v1970
	if v1986 < int64(0) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	v1996 = v1970
	goto L268
L268:
	;
	v1997 = *(*int64)(unsafe.Add(mBase, uint32(v1656)))
	v2025 = v1996
	v2027 = v1997
	goto L242
L269:
	;
	v1991 = v1986 + int64(100000000)
	goto L271
L270:
	;
	v1991 = v1986
	goto L271
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1979))) = v1991
	v1996 = v1986 >> (uint(int64(63)) % 64)
	goto L268
L272:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L17
	} else {
		goto L273
	}
L273:
	;
	F_errmsg(m, int32(249746), int32(0))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L17
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(520760), int32(9400), int32(239169))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L17
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L17
	} else {
		goto L277
	}
L277:
	;
	F_errmsg(m, int32(249746), int32(0))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L17
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(520760), int32(10040), int32(580880))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L17
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	v2112 = v1656 + v2082<<(uint(int32(3))%32)
	v2116 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000+v2082<<(uint(int32(2))%32)))))
	v2117 = *(*int64)(unsafe.Add(mBase, uint32(v2112)))
	v2119 = v2116 + (v2117 + v2103)
	v2123 = base.B2i32(int64(99999999) < v2119)
	if int64(99999999) < v2119 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	goto L2
L282:
	;
	v2124 = v2119 - int64(100000000)
	goto L284
L283:
	;
	v2124 = v2119
	goto L284
L284:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2112))) = v2124
	v2126 = base.I64_extend_i32_u(v2123)
	v2127 = int32(1)
	if base.Ui32(v2127) < base.Ui32(v2082) {
		v2082 = v2082 - v2127
		v2103 = v2126
		goto L280
	} else {
		goto L285
	}
L285:
	;
	v2131 = *(*int64)(unsafe.Add(mBase, uint32(v1656)))
	v2132 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1000))))
	v2134 = v2131 + (v2126 + v2132)
	*(*int64)(unsafe.Add(mBase, uint32(v1656))) = v2134
	v2136 = *(*int64)(unsafe.Add(mBase, uint32(v2075)))
	*(*int64)(unsafe.Add(mBase, uint32(v2075))) = v2136 - int64(1)
	v2140 = int64(0)
	if v2134 < v2140 {
		v2082 = v1015
		v2103 = v2140
		goto L280
	} else {
		goto L286
	}
L286:
	;
	goto L281
L287:
	;
	F_pfree(m, v2180)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L17
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v2184 = v764 + int32(1)
	v2187 = int32(2)
	v2191 = F_palloc(m, v785<<(uint(v2187)%32)|v2187)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L17
	} else {
		goto L291
	}
L290:
	;
	goto L289
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v2191
	v2194 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2191))) = uint16(v2194)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v785 << (uint(int32(1)) % 32)
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2198 = int32(2)
	v2199 = v2197 + v2198
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2199
	if v2198 <= v783 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v2215 = v785
	v2232 = int64(0)
	goto L295
L293:
	;
	goto L294
L294:
	;
	F_pfree(m, v815)
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L17
	} else {
		goto L303
	}
L295:
	;
	v2239 = v2215 - int32(1)
	v2243 = *(*int64)(unsafe.Add(mBase, uint32(v815+v2239<<(uint(int32(3))%32))))
	v2244 = v2243 + v2232
	if v2244 < int64(0) {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	goto L294
L297:
	;
	v2268 = v2199 + v2239<<(uint(int32(2))%32)
	v2269 = base.I32_wrap_i64(v2264)
	v2270 = int32(10000)
	v2271 = base.I32_div_u_s(v2269, v2270)
	*(*uint16)(unsafe.Add(mBase, uint32(v2268))) = uint16(v2271)
	v2275 = v2269 - v2271*v2270
	*(*uint16)(unsafe.Add(mBase, uint32(v2268)+2)) = uint16(v2275)
	if int32(1) < v2215 {
		v2215 = v2239
		v2232 = v2265
		goto L295
	} else {
		goto L302
	}
L298:
	;
	v2247 = int64(-1)
	v2250 = base.I64_div_u_s(v2244^v2247, int64(100000000))
	v2252 = v2250 ^ v2247
	v2264 = v2252*int64(-100000000) + v2244
	v2265 = v2252
	goto L297
L299:
	;
	goto L300
L300:
	;
	if base.Ui64(v2244) < base.Ui64(int64(100000000)) {
		v2264 = v2244
		v2265 = int64(0)
		goto L297
	} else {
		goto L301
	}
L301:
	;
	v2260 = base.I64_div_u_s(v2244, int64(100000000))
	v2264 = v2260*int64(-100000000) + v2244
	v2265 = v2260
	goto L297
L302:
	;
	goto L296
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = base.B2i32(v806 != v807) << (uint(int32(14)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2184
	if l4 != 0 {
		goto L308
	} else {
		goto L309
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2620
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2614
	goto L1
L305:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v2614 = v2577
	v2620 = int32(0)
	goto L304
L306:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if int32(0) < v2481 {
		goto L345
	} else {
		goto L346
	}
L307:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2481 = v2479
	goto L306
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2326 = l3 + v2323<<(uint(int32(2))%32)
	if v2326+int32(4) < int32(0) {
		goto L312
	} else {
		goto L313
	}
L309:
	;
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v2445 = l3 + v2184<<(uint(int32(2))%32)
	if v2445+int32(4) <= int32(0) {
		goto L338
	} else {
		goto L339
	}
L311:
	;
	goto L307
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	goto L311
L313:
	;
	goto L314
L314:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2337 = l3 & int32(3)
	v2341 = base.I32_div_s(v2326+int32(7), int32(4))
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2342 <= v2341 {
		goto L319
	} else {
		goto L320
	}
L315:
	;
	goto L311
L316:
	;
	if int32(0) <= v2408 {
		goto L315
	} else {
		goto L337
	}
L317:
	;
	v2388 = v2382
	goto L331
L318:
	;
	v2355 = int32(1)
	v2356 = v2341 - v2355
	v2359 = v2335 + v2356<<(uint(v2355)%32)
	v2360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2359))))
	v2361 = int32(2)
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v2337<<(uint(v2361)%32))+uint32(_consts[1070])))
	v2366 = base.I32_rem_s(v2360, v2365)
	v2367 = v2360 - v2366
	*(*uint16)(unsafe.Add(mBase, uint32(v2359))) = uint16(v2367)
	v2370 = base.I32_div_s(v2365, v2361)
	if v2366 < v2370 {
		v2408 = v2356
		goto L316
	} else {
		goto L326
	}
L319:
	;
	if v2337 == int32(0) {
		goto L315
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2341
	if v2337 != 0 {
		goto L318
	} else {
		goto L324
	}
L322:
	;
	if v2341 != v2342 {
		goto L315
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2341
	goto L318
L324:
	;
	v2352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2335+v2341<<(uint(int32(1))%32)))))
	if v2352 <= int32(4999) {
		v2408 = v2341
		goto L316
	} else {
		goto L325
	}
L325:
	;
	v2382 = v2341
	goto L317
L326:
	;
	v2373 = v2365 + base.I32_extend16_s(v2367)
	if int32(9999) < v2373 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2378 = v2373 + int32(55536)
	goto L329
L328:
	;
	v2378 = v2373
	goto L329
L329:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2359))) = uint16(v2378)
	if v2373 < int32(10000) {
		v2408 = v2356
		goto L316
	} else {
		goto L330
	}
L330:
	;
	v2382 = v2356
	goto L317
L331:
	;
	v2394 = int32(1)
	v2395 = v2388 - v2394
	v2398 = v2335 + v2395<<(uint(v2394)%32)
	v2401 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2398))))
	v2403 = base.B2i32(int32(9998) < v2401)
	if int32(9998) < v2401 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v2408 = v2395
	goto L316
L333:
	;
	v2404 = int32(-9999)
	goto L335
L334:
	;
	v2404 = v2394
	goto L335
L335:
	;
	v2405 = v2404 + v2401
	*(*uint16)(unsafe.Add(mBase, uint32(v2398))) = uint16(v2405)
	if int32(9998) < v2401 {
		v2388 = v2395
		goto L331
	} else {
		goto L336
	}
L336:
	;
	goto L332
L337:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2416 - int32(2)
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2421 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2420 + v2421
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2424 + v2421
	goto L315
L338:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2577 = v2450
	goto L305
L339:
	;
	goto L340
L340:
	;
	v2454 = base.I32_div_s(v2445+int32(7), int32(4))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2455 < v2454 {
		v2481 = v2455
		goto L306
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2454
	v2459 = l3 & int32(3)
	if v2459 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v2481 = v2454
	goto L306
L343:
	;
	goto L344
L344:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2466 = int32(2)
	v2467 = v2462 + v2454<<(uint(int32(1))%32) - v2466
	v2468 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2467))))
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2459<<(uint(v2466)%32))+uint32(_consts[1070])))
	v2474 = base.I32_rem_s(v2468, v2473)
	v2475 = v2468 - v2474
	*(*uint16)(unsafe.Add(mBase, uint32(v2467))) = uint16(v2475)
	goto L307
L345:
	;
	v2488 = v2482
	v2494 = v2481
	goto L349
L346:
	;
	goto L347
L347:
	;
	if v2481 != 0 {
		v2614 = v2482
		v2620 = v2481
		goto L304
	} else {
		goto L357
	}
L348:
	;
	v2541 = v2494
	goto L353
L349:
	;
	v2522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2488))))
	if v2522 != 0 {
		goto L348
	} else {
		goto L351
	}
L350:
	;
	v2577 = v2482 + v2481<<(uint(int32(1))%32)
	goto L305
L351:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2524 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2523 - v2524
	if v2524 < v2494 {
		v2488 = v2488 + int32(2)
		v2494 = v2494 - v2524
		goto L349
	} else {
		goto L352
	}
L352:
	;
	goto L350
L353:
	;
	v2572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2488-int32(2)+v2541<<(uint(int32(1))%32)))))
	if v2572 != 0 {
		v2614 = v2488
		v2620 = v2541
		goto L304
	} else {
		goto L355
	}
L354:
	;
	v2577 = v2488
	goto L305
L355:
	;
	v2573 = int32(1)
	if v2573 < v2541 {
		v2541 = v2541 - v2573
		goto L353
	} else {
		goto L356
	}
L356:
	;
	goto L354
L357:
	;
	v2577 = v2482
	goto L305
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
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
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int64
	_ = v452
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
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
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1638 int32
	_ = v1638
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
	var v2001 int32
	_ = v2001
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2079 int32
	_ = v2079
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int32
	_ = v2103
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2198 int32
	_ = v2198
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2274 int32
	_ = v2274
	var v2278 int32
	_ = v2278
	var v2283 int32
	_ = v2283
	var v2287 int32
	_ = v2287
	var v2293 int32
	_ = v2293
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2315 int32
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
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
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L41
	} else {
		goto L735
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L41
	} else {
		goto L731
	}
L3:
	;
	m.G0 = v15 + int32(96)
	return
L4:
	;
	v2299 = F_get_object_catcache_oid(m, v17)
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L41
	} else {
		goto L712
	}
L5:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2216 = m.G0
	v2218 = v2216 - int32(32)
	m.G0 = v2218
	v2222 = F_table_open(m, int32(1255), int32(3))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L41
	} else {
		goto L686
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L41
	} else {
		goto L683
	}
L7:
	;
	if v17 == int32(2328) {
		goto L4
	} else {
		goto L682
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L41
	} else {
		goto L679
	}
L9:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2142 = m.G0
	v2144 = v2142 - int32(16)
	m.G0 = v2144
	v2148 = F_table_open(m, int32(6104), int32(3))
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L41
	} else {
		goto L663
	}
L10:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2052 = m.G0
	v2054 = v2052 - int32(16)
	m.G0 = v2054
	v2058 = F_table_open(m, int32(6106), int32(3))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L41
	} else {
		goto L639
	}
L11:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1963 = m.G0
	v1965 = v1963 - int32(16)
	m.G0 = v1965
	v1969 = F_table_open(m, int32(6237), int32(3))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L41
	} else {
		goto L615
	}
L12:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1904 = m.G0
	v1906 = v1904 + int32(-64)
	m.G0 = v1906
	v1909 = *(*int32)(unsafe.Add(mBase, _consts[222]))
	if v1909 != v1903 {
		goto L597
	} else {
		goto L598
	}
L13:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1825 = m.G0
	v1827 = v1825 + int32(-64)
	m.G0 = v1827
	v1831 = F_table_open(m, int32(3602), int32(3))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L41
	} else {
		goto L572
	}
L14:
	;
	if v17 != int32(3381) {
		goto L6
	} else {
		goto L540
	}
L15:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1594 = m.G0
	v1596 = v1594 - int32(96)
	m.G0 = v1596
	v1600 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L41
	} else {
		goto L502
	}
L16:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1494 = m.G0
	v1496 = v1494 - int32(80)
	m.G0 = v1496
	v1500 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L41
	} else {
		goto L472
	}
L17:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1419 = m.G0
	v1421 = v1419 - int32(32)
	m.G0 = v1421
	v1425 = F_table_open(m, int32(2617), int32(3))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L41
	} else {
		goto L446
	}
L18:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1323 = m.G0
	v1325 = v1323 + int32(-64)
	m.G0 = v1325
	v1329 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L41
	} else {
		goto L419
	}
L19:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1224 = m.G0
	v1226 = v1224 - int32(80)
	m.G0 = v1226
	v1230 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L41
	} else {
		goto L394
	}
L20:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1084 = m.G0
	v1086 = v1084 + int32(-64)
	m.G0 = v1086
	v1090 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L41
	} else {
		goto L353
	}
L21:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v890 = m.G0
	v892 = v890 - int32(16)
	m.G0 = v892
	v896 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L41
	} else {
		goto L307
	}
L22:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v180 = F_get_rel_relkind(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
	F_ScanKeyInit(m, v45+int32(48), int32(1), int32(3), int32(184), v42)
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
	v64 = F_systable_beginscan(m, v49, int32(3257), v59, int32(0), v59, v45+int32(48))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
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
	v149 = m.ExcPending
	if v149 != 0 {
		goto L41
	} else {
		goto L74
	}
L46:
	;
	v66 = F_systable_getnext(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	if v66 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68+v69)+68))
	v73 = F_table_open(m, v71, int32(8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
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
	v136 = m.ExcPending
	if v136 != 0 {
		goto L41
	} else {
		goto L71
	}
L51:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[223])))
	if v101 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L41
	} else {
		goto L54
	}
L53:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+119)))
	switch v76 - int32(112) {
	case 0, 2:
		goto L51
	default:
		goto L52
	}
L54:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L41
	} else {
		goto L55
	}
L55:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v86 + int32(4)
	F_errmsg(m, int32(411055), v45+int32(16))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L41
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(511555), int32(374), int32(482804))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
	v105 = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v73)+56))
	if base.Ui32(v106) < base.Ui32(int32(12000)) {
		v115 = v105
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	F_CatalogTupleDelete(m, v49, v66+int32(4))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L41
	} else {
		goto L66
	}
L61:
	;
	if v115 != 0 {
		goto L45
	} else {
		goto L65
	}
L62:
	;
	goto L61
L63:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+68))
	if v110 == int32(99) {
		v115 = v105
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v113 = F_isTempToastNamespace(m, v110)
	mBase = m.M
	v115 = v113
	goto L62
L65:
	;
	goto L60
L66:
	;
	F_systable_endscan(m, v64)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L41
	} else {
		goto L67
	}
L67:
	;
	F_CacheInvalidateRelcache(m, v73)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L41
	} else {
		goto L68
	}
L68:
	;
	F_sequence_close(m, v73, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L41
	} else {
		goto L69
	}
L69:
	;
	F_sequence_close(m, v49, int32(3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	F_errmsg_internal(m, int32(42291), v45)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(511555), int32(358), int32(482804))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
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
	v152 = m.ExcPending
	if v152 != 0 {
		goto L41
	} else {
		goto L75
	}
L75:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = v153 + int32(4)
	F_errmsg(m, int32(341295), v45+int32(32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L41
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(511555), int32(380), int32(482804))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	switch v17 - int32(6100) {
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
	switch v17 - int32(6237) {
	case 0:
		goto L11
	default:
		goto L6
	case 6:
		goto L8
	}
L83:
	;
	if v180 != int32(83) {
		goto L3
	} else {
		goto L295
	}
L84:
	;
	if v180&int32(-33) == int32(73) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v190 = int32(base.Ui32(l1&int32(2)) >> (uint(int32(1)) % 32))
	v195 = m.G0
	v197 = v195 - int32(96)
	m.G0 = v197
	v200 = F_SearchSysCache1(m, int32(34), v186)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L41
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v438 != 0 {
		goto L175
	} else {
		goto L176
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L41
	} else {
		goto L172
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L41
	} else {
		goto L168
	}
L90:
	;
	if v200 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+22)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202+v203)+4))
	F_ReleaseCatCache(m, v200)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
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
	v396 = m.ExcPending
	if v396 != 0 {
		goto L41
	} else {
		goto L165
	}
L94:
	;
	v208 = int32(4)
	if int32(base.Ui32(l1&int32(32))>>(uint(int32(5))%32)) != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v211 = v208
	goto L97
L96:
	;
	v211 = int32(8)
	goto L97
L97:
	;
	if v190 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v212 = v208
	goto L100
L99:
	;
	v212 = v211
	goto L100
L100:
	;
	v213 = F_table_open(m, v205, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L41
	} else {
		goto L101
	}
L101:
	;
	v215 = F_index_open(m, v186, v212)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L41
	} else {
		goto L102
	}
L102:
	;
	F_CheckTableNotInUse(m, v215, int32(531284))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L41
	} else {
		goto L103
	}
L103:
	;
	if v190 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v316)+48))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+119)))
	switch v321 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L136
	default:
		goto L135
	}
L105:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v221 != 0 {
		goto L89
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_TransferPredicateLocksToHeapRelation(m, v215)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L41
	} else {
		goto L134
	}
L108:
	;
	F_index_set_state_flags(m, v186, int32(2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L41
	} else {
		goto L109
	}
L109:
	;
	F_CacheInvalidateRelcache(m, v213)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L41
	} else {
		goto L110
	}
L110:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v213)+60))
	v229 = v197 + int32(72)
	*(*int64)(unsafe.Add(mBase, uint32(v229))) = int64(72057594037927936)
	*(*uint32)(unsafe.Add(mBase, uint32(v197)+68)) = uint32(v227)
	*(*int64)(unsafe.Add(mBase, uint32(v197)+88)) = v227
	v235 = int64(base.Ui64(v227) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v197)+64)) = uint32(v235)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v215)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+80)) = v237
	F_sequence_close(m, v213, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L41
	} else {
		goto L111
	}
L111:
	;
	F_relation_close(m, v215, int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L41
	} else {
		goto L112
	}
L112:
	;
	F_LockRelationIdForSession(m, v197+int32(88), int32(4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L41
	} else {
		goto L113
	}
L113:
	;
	F_LockRelationIdForSession(m, v197+int32(80), int32(4))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L41
	} else {
		goto L114
	}
L114:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L41
	} else {
		goto L115
	}
L115:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L41
	} else {
		goto L116
	}
L116:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L41
	} else {
		goto L117
	}
L117:
	;
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+56)) = v261
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v197)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+48)) = v263
	F_WaitForLockers(m, v197+int32(48), int32(8))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L41
	} else {
		goto L118
	}
L118:
	;
	v270 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L41
	} else {
		goto L119
	}
L119:
	;
	F_PushActiveSnapshot(m, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L41
	} else {
		goto L120
	}
L120:
	;
	v275 = F_table_open(m, v205, int32(4))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L41
	} else {
		goto L121
	}
L121:
	;
	v278 = F_index_open(m, v186, int32(4))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L41
	} else {
		goto L122
	}
L122:
	;
	F_TransferPredicateLocksToHeapRelation(m, v278)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L41
	} else {
		goto L123
	}
L123:
	;
	F_index_set_state_flags(m, v186, int32(3))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L41
	} else {
		goto L124
	}
L124:
	;
	F_CacheInvalidateRelcache(m, v275)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L41
	} else {
		goto L125
	}
L125:
	;
	F_sequence_close(m, v275, int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L41
	} else {
		goto L126
	}
L126:
	;
	F_relation_close(m, v278, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L41
	} else {
		goto L127
	}
L127:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L41
	} else {
		goto L128
	}
L128:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L41
	} else {
		goto L129
	}
L129:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L41
	} else {
		goto L130
	}
L130:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+40)) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v197)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+32)) = v301
	F_WaitForLockers(m, v197+int32(32), int32(8))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L41
	} else {
		goto L131
	}
L131:
	;
	v309 = F_table_open(m, v205, int32(4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L41
	} else {
		goto L132
	}
L132:
	;
	v312 = F_index_open(m, v186, int32(8))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L41
	} else {
		goto L133
	}
L133:
	;
	v316 = v312
	v317 = v309
	goto L104
L134:
	;
	v316 = v215
	v317 = v213
	goto L104
L135:
	;
	F_pgstat_drop_relation(m, v316)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L41
	} else {
		goto L138
	}
L136:
	;
	F_RelationDropStorage(m, v316)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L41
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	F_relation_close(m, v316, int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L41
	} else {
		goto L139
	}
L139:
	;
	F_RelationForgetRelation(m, v186)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L41
	} else {
		goto L140
	}
L140:
	;
	v333 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L41
	} else {
		goto L141
	}
L141:
	;
	F_PushActiveSnapshot(m, v333)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L41
	} else {
		goto L142
	}
L142:
	;
	v339 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L41
	} else {
		goto L143
	}
L143:
	;
	v342 = F_SearchSysCache1(m, int32(34), v186)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L41
	} else {
		goto L144
	}
L144:
	;
	if v342 == int32(0) {
		goto L88
	} else {
		goto L145
	}
L145:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v339)+52))
	v348 = F_heap_attisnull(m, v342, int32(20), v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L41
	} else {
		goto L146
	}
L146:
	;
	F_CatalogTupleDelete(m, v339, v342+int32(4))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L41
	} else {
		goto L147
	}
L147:
	;
	F_ReleaseCatCache(m, v342)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L41
	} else {
		goto L148
	}
L148:
	;
	F_sequence_close(m, v339, int32(3))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L41
	} else {
		goto L149
	}
L149:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L41
	} else {
		goto L150
	}
L150:
	;
	if v348 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	F_RemoveStatistics(m, v186, int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L41
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	F_DeleteAttributeTuples(m, v186)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L41
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	F_DeleteRelationTuple(m, v186)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L41
	} else {
		goto L156
	}
L156:
	;
	v370 = int32(0)
	v373 = F_DeleteInheritsTuple(m, v186, v370, v370, v370)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L41
	} else {
		goto L157
	}
L157:
	;
	F_CacheInvalidateRelcache(m, v317)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L41
	} else {
		goto L158
	}
L158:
	;
	F_sequence_close(m, v317, int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L41
	} else {
		goto L159
	}
L159:
	;
	if v190 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	F_UnlockRelationIdForSession(m, v197+int32(88), int32(4))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L41
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	m.G0 = v197 + int32(96)
	goto L83
L163:
	;
	F_UnlockRelationIdForSession(m, v197+int32(80), int32(4))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L41
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v186
	F_errmsg_internal(m, int32(42447), v197)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L41
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(511623), int32(3594), int32(274776))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	v412 = m.ExcPending
	if v412 != 0 {
		goto L41
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(267596), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L41
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(511623), int32(2222), int32(244300))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v186
	F_errmsg_internal(m, int32(42447), v197+int32(16))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L41
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(511623), int32(2353), int32(244300))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
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
	v439 = base.I32_extend16_s(v438)
	v440 = m.G0
	v442 = v440 - int32(272)
	m.G0 = v442
	v449 = F__emscripten_memset_bulkmem(m, v442+int32(96), base.I32_extend8_s(int32(0)), int32(100))
	mBase = m.M
	goto L178
L176:
	;
	goto L177
L177:
	;
	v553 = m.G0
	v555 = v553 - int32(80)
	m.G0 = v555
	v558 = F_SearchSysCache1(m, int32(57), v437)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L41
	} else {
		goto L197
	}
L178:
	;
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+88)) = uint8(v450)
	v452 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v442)+80)) = v452
	*(*int64)(unsafe.Add(mBase, uint32(v442)+72)) = v452
	*(*int64)(unsafe.Add(mBase, uint32(v442)+64)) = v452
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+56)) = uint8(v450)
	*(*int64)(unsafe.Add(mBase, uint32(v442)+48)) = v452
	*(*int64)(unsafe.Add(mBase, uint32(v442)+40)) = v452
	*(*int64)(unsafe.Add(mBase, uint32(v442)+32)) = v452
	v467 = F_relation_open(m, v437, int32(8))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L41
	} else {
		goto L179
	}
L179:
	;
	v471 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L41
	} else {
		goto L180
	}
L180:
	;
	v474 = F_SearchSysCacheCopy(m, int32(7), v437, v439)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L41
	} else {
		goto L181
	}
L181:
	;
	if v474 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L41
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v474)+16))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+22)))
	v494 = v492 + v493
	v495 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+86)) = uint8(v495)
	*(*int32)(unsafe.Add(mBase, uint32(v494)+68)) = v495
	v499 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v494)+90)) = uint16(v499)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+16)) = v439
	v508 = F_pg_snprintf(m, v442+int32(208), int32(64), int32(684976), v442+int32(16))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L41
	} else {
		goto L188
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+4)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v439
	F_errmsg_internal(m, int32(49690), v442)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L41
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(516025), int32(1709), int32(483070))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L41
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
	v515 = F_strncpy(m, v494+int32(4), v442+int32(208), int32(64))
	mBase = m.M
	v516 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v515)+63)) = uint8(v516)
	goto L189
L189:
	;
	v518 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+88)) = uint8(v518)
	v520 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+56)) = uint8(v520)
	v522 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+85)) = v522
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+84)) = uint8(v520)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+52)) = v522
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v471)+52))
	v535 = F_heap_modify_tuple(m, v474, v528, v442+int32(96), v442-int32(-64), v442+int32(32))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L41
	} else {
		goto L190
	}
L190:
	;
	F_CatalogTupleUpdate(m, v471, v535+int32(4), v535)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L41
	} else {
		goto L191
	}
L191:
	;
	F_sequence_close(m, v471, int32(3))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L41
	} else {
		goto L192
	}
L192:
	;
	F_RemoveStatistics(m, v437, v439)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L41
	} else {
		goto L193
	}
L193:
	;
	F_relation_close(m, v467, int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L41
	} else {
		goto L194
	}
L194:
	;
	m.G0 = v442 + int32(272)
	goto L83
L195:
	;
	goto L83
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L41
	} else {
		goto L292
	}
L197:
	;
	if v558 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v558)+16))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+22)))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560+v561)+131)))
	if v563 != int32(1) {
		v581 = int32(0)
		v582 = int32(0)
		goto L201
	} else {
		goto L202
	}
L199:
	;
	goto L200
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L41
	} else {
		goto L289
	}
L201:
	;
	F_ReleaseCatCache(m, v558)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L41
	} else {
		goto L213
	}
L202:
	;
	v567 = F_get_partition_parent(m, v437, int32(1))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L41
	} else {
		goto L203
	}
L203:
	;
	F_LockRelationOid(m, v567, int32(8))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L41
	} else {
		goto L204
	}
L204:
	;
	v572 = F_get_default_partition_oid(m, v567)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L41
	} else {
		goto L205
	}
L205:
	;
	if v572 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v581 = int32(0)
	v582 = v567
	goto L201
L207:
	;
	goto L208
L208:
	;
	if v437 == v572 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v581 = v437
	v582 = v567
	goto L201
L210:
	;
	goto L211
L211:
	;
	F_LockRelationOid(m, v572, int32(8))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L41
	} else {
		goto L212
	}
L212:
	;
	v581 = v572
	v582 = v567
	goto L201
L213:
	;
	v586 = F_relation_open(m, v437, int32(8))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L41
	} else {
		goto L214
	}
L214:
	;
	F_CheckTableNotInUse(m, v586, int32(562026))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L41
	} else {
		goto L215
	}
L215:
	;
	F_CheckTableForSerializableConflictIn(m, v586)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L41
	} else {
		goto L216
	}
L216:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v586)+48))
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+119)))
	if v594 == int32(102) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v599 = F_table_open(m, int32(3118), int32(3))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L41
	} else {
		goto L220
	}
L218:
	;
	v619 = v594
	goto L219
L219:
	;
	if v619&int32(255) == int32(112) {
		goto L226
	} else {
		goto L227
	}
L220:
	;
	v602 = F_SearchSysCache1(m, int32(33), v437)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L41
	} else {
		goto L221
	}
L221:
	;
	if v602 == int32(0) {
		goto L196
	} else {
		goto L222
	}
L222:
	;
	F_CatalogTupleDelete(m, v599, v602+int32(4))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L41
	} else {
		goto L223
	}
L223:
	;
	F_ReleaseCatCache(m, v602)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L41
	} else {
		goto L224
	}
L224:
	;
	F_sequence_close(m, v599, int32(3))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L41
	} else {
		goto L225
	}
L225:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v586)+48))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+119)))
	v619 = v616
	goto L219
L226:
	;
	v624 = m.G0
	v626 = v624 - int32(16)
	m.G0 = v626
	v630 = F_table_open(m, int32(3350), int32(3))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L41
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	if v437 == v581 {
		goto L240
	} else {
		goto L241
	}
L229:
	;
	v633 = F_SearchSysCache1(m, int32(45), v437)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L41
	} else {
		goto L230
	}
L230:
	;
	if v633 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L41
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	F_CatalogTupleDelete(m, v630, v633+int32(4))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L41
	} else {
		goto L237
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v437
	F_errmsg_internal(m, int32(49538), v626)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L41
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(516025), int32(4032), int32(483418))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L41
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_ReleaseCatCache(m, v633)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L41
	} else {
		goto L238
	}
L238:
	;
	F_sequence_close(m, v630, int32(3))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L41
	} else {
		goto L239
	}
L239:
	;
	m.G0 = v626 + int32(16)
	goto L228
L240:
	;
	F_update_default_partition_oid(m, v582, int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L41
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v586)+48))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+119)))
	switch v670 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L245
	default:
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	F_pgstat_drop_relation(m, v586)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L41
	} else {
		goto L247
	}
L245:
	;
	F_RelationDropStorage(m, v586)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L41
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	F_relation_close(m, v586, int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L41
	} else {
		goto L248
	}
L248:
	;
	F_RemoveSubscriptionRel(m, int32(0), v437)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L41
	} else {
		goto L249
	}
L249:
	;
	v684 = *(*int32)(unsafe.Add(mBase, _consts[224]))
	if v684 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	F_RelationForgetRelation(m, v437)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L41
	} else {
		goto L263
	}
L251:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	if v687 <= int32(0) {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v690 = int32(0)
	if v690 < v687 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v693 = v687
	goto L255
L254:
	;
	v693 = v690
	goto L255
L255:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v684)+12))
	v700 = int32(0)
	goto L256
L256:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v694+v700<<(uint(int32(2))%32))))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	if v437 != v712 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+8))
	goto L262
L258:
	;
	v715 = v700 + int32(1)
	if v693 != v715 {
		v700 = v715
		goto L256
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	goto L257
L261:
	;
	goto L250
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711)+12)) = v719
	goto L250
L263:
	;
	v737 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L41
	} else {
		goto L264
	}
L264:
	;
	F_ScanKeyInit(m, v555+int32(32), int32(1), int32(3), int32(184), v437)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L41
	} else {
		goto L265
	}
L265:
	;
	v747 = int32(1)
	v752 = F_systable_beginscan(m, v737, int32(2680), v747, int32(0), v747, v555+int32(32))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L41
	} else {
		goto L266
	}
L266:
	;
	v754 = F_systable_getnext(m, v752)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L41
	} else {
		goto L267
	}
L267:
	;
	if v754 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v760 = v754
	goto L271
L269:
	;
	goto L270
L270:
	;
	F_systable_endscan(m, v752)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L41
	} else {
		goto L276
	}
L271:
	;
	F_CatalogTupleDelete(m, v737, v760+int32(4))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L41
	} else {
		goto L273
	}
L272:
	;
	goto L270
L273:
	;
	v772 = F_systable_getnext(m, v752)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L41
	} else {
		goto L274
	}
L274:
	;
	if v772 != 0 {
		v760 = v772
		goto L271
	} else {
		goto L275
	}
L275:
	;
	goto L272
L276:
	;
	F_sequence_close(m, v737, int32(3))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L41
	} else {
		goto L277
	}
L277:
	;
	F_RemoveStatistics(m, v437, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L41
	} else {
		goto L278
	}
L278:
	;
	F_DeleteAttributeTuples(m, v437)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L41
	} else {
		goto L279
	}
L279:
	;
	F_DeleteRelationTuple(m, v437)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L41
	} else {
		goto L280
	}
L280:
	;
	if v582 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	if v581 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	m.G0 = v555 + int32(80)
	goto L195
L284:
	;
	F_CacheInvalidateRelcacheByRelid(m, v582)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L41
	} else {
		goto L288
	}
L285:
	;
	if v437 == v581 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	F_CacheInvalidateRelcacheByRelid(m, v581)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L41
	} else {
		goto L287
	}
L287:
	;
	goto L284
L288:
	;
	goto L283
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v437
	F_errmsg_internal(m, int32(49502), v555)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L41
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(516025), int32(1803), int32(341236))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L41
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+16)) = v437
	F_errmsg_internal(m, int32(56256), v555+int32(16))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L41
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(516025), int32(1857), int32(341236))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L41
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v851 = m.G0
	v853 = v851 - int32(16)
	m.G0 = v853
	v857 = F_table_open(m, int32(2224), int32(3))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L41
	} else {
		goto L296
	}
L296:
	;
	v860 = F_SearchSysCache1(m, int32(61), v850)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L41
	} else {
		goto L297
	}
L297:
	;
	if v860 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L41
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	F_CatalogTupleDelete(m, v857, v860+int32(4))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L41
	} else {
		goto L304
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v853))) = v850
	F_errmsg_internal(m, int32(57190), v853)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L41
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(519990), int32(579), int32(400955))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L41
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	F_ReleaseCatCache(m, v860)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L41
	} else {
		goto L305
	}
L305:
	;
	F_sequence_close(m, v857, int32(3))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L41
	} else {
		goto L306
	}
L306:
	;
	m.G0 = v853 + int32(16)
	goto L3
L307:
	;
	v899 = F_SearchSysCache1(m, int32(82), v889)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L41
	} else {
		goto L309
	}
L308:
	;
	goto L3
L309:
	;
	if v899 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	F_CatalogTupleDelete(m, v896, v899+int32(4))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L41
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L41
	} else {
		goto L350
	}
L313:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v899)+16))
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905)+22)))
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905+v906)+79)))
	if v908 == int32(101) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v911 = m.G0
	v913 = v911 - int32(48)
	m.G0 = v913
	v917 = F_table_open(m, int32(3501), int32(3))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L41
	} else {
		goto L317
	}
L315:
	;
	v986 = v908
	goto L316
L316:
	;
	if v986&int32(255) == int32(114) {
		goto L331
	} else {
		goto L332
	}
L317:
	;
	F_ScanKeyInit(m, v913, int32(2), int32(3), int32(184), v889)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L41
	} else {
		goto L318
	}
L318:
	;
	v925 = int32(1)
	v928 = F_systable_beginscan(m, v917, int32(3503), v925, int32(0), v925, v913)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L41
	} else {
		goto L319
	}
L319:
	;
	v930 = F_systable_getnext(m, v928)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L41
	} else {
		goto L320
	}
L320:
	;
	if v930 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v932 = v930
	goto L324
L322:
	;
	goto L323
L323:
	;
	F_systable_endscan(m, v928)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L41
	} else {
		goto L329
	}
L324:
	;
	F_CatalogTupleDelete(m, v917, v932+int32(4))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L41
	} else {
		goto L326
	}
L325:
	;
	goto L323
L326:
	;
	v948 = F_systable_getnext(m, v928)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L41
	} else {
		goto L327
	}
L327:
	;
	if v948 != 0 {
		v932 = v948
		goto L324
	} else {
		goto L328
	}
L328:
	;
	goto L325
L329:
	;
	F_sequence_close(m, v917, int32(3))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L41
	} else {
		goto L330
	}
L330:
	;
	m.G0 = v913 + int32(48)
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v899)+16))
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+22)))
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970+v971)+79)))
	v986 = v973
	goto L316
L331:
	;
	v991 = m.G0
	v993 = v991 - int32(48)
	m.G0 = v993
	v997 = F_table_open(m, int32(3541), int32(3))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L41
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	F_ReleaseCatCache(m, v899)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L41
	} else {
		goto L348
	}
L334:
	;
	F_ScanKeyInit(m, v993, int32(1), int32(3), int32(184), v889)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L41
	} else {
		goto L335
	}
L335:
	;
	v1005 = int32(1)
	v1008 = F_systable_beginscan(m, v997, int32(3542), v1005, int32(0), v1005, v993)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L41
	} else {
		goto L336
	}
L336:
	;
	v1010 = F_systable_getnext(m, v1008)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L41
	} else {
		goto L337
	}
L337:
	;
	if v1010 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1014 = v1010
	goto L341
L339:
	;
	goto L340
L340:
	;
	F_systable_endscan(m, v1008)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L41
	} else {
		goto L346
	}
L341:
	;
	F_CatalogTupleDelete(m, v997, v1014+int32(4))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L41
	} else {
		goto L343
	}
L342:
	;
	goto L340
L343:
	;
	v1028 = F_systable_getnext(m, v1008)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L41
	} else {
		goto L344
	}
L344:
	;
	if v1028 != 0 {
		v1014 = v1028
		goto L341
	} else {
		goto L345
	}
L345:
	;
	goto L342
L346:
	;
	F_sequence_close(m, v997, int32(3))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L41
	} else {
		goto L347
	}
L347:
	;
	m.G0 = v993 + int32(48)
	goto L333
L348:
	;
	F_sequence_close(m, v896, int32(3))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L41
	} else {
		goto L349
	}
L349:
	;
	m.G0 = v892 + int32(16)
	goto L308
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892))) = v889
	F_errmsg_internal(m, int32(54674), v892)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L41
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(514319), int32(666), int32(483090))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L41
	} else {
		goto L352
	}
L352:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L353:
	;
	v1093 = F_SearchSysCache1(m, int32(19), v1083)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L41
	} else {
		goto L358
	}
L354:
	;
	goto L3
L355:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L41
	} else {
		goto L391
	}
L356:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L41
	} else {
		goto L388
	}
L357:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L41
	} else {
		goto L385
	}
L358:
	;
	if v1093 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+16))
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1095)+22)))
	v1097 = v1095 + v1096
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+80))
	if v1098 != 0 {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	goto L361
L361:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L41
	} else {
		goto L382
	}
L362:
	;
	F_CatalogTupleDelete(m, v1090, v1093+int32(4))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L41
	} else {
		goto L379
	}
L363:
	;
	v1100 = F_table_open(m, v1098, int32(8))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L41
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+84))
	if v1141 == int32(0) {
		goto L355
	} else {
		goto L378
	}
L366:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097)+72)))
	if v1102 == int32(99) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1107 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L41
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_sequence_close(m, v1100, int32(0))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L41
	} else {
		goto L377
	}
L370:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+80))
	v1112 = F_SearchSysCacheCopy(m, int32(57), v1110, int32(0))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L41
	} else {
		goto L371
	}
L371:
	;
	if v1112 == int32(0) {
		goto L357
	} else {
		goto L372
	}
L372:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+16))
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116)+22)))
	v1118 = v1116 + v1117
	v1119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1118)+122)))
	if v1119 == int32(0) {
		goto L356
	} else {
		goto L373
	}
L373:
	;
	v1123 = v1119 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1118)+122)) = uint16(v1123)
	F_CatalogTupleUpdate(m, v1107, v1112+int32(4), v1112)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L41
	} else {
		goto L374
	}
L374:
	;
	F_pfree(m, v1112)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L41
	} else {
		goto L375
	}
L375:
	;
	F_sequence_close(m, v1107, int32(3))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L41
	} else {
		goto L376
	}
L376:
	;
	goto L369
L377:
	;
	goto L362
L378:
	;
	goto L362
L379:
	;
	F_ReleaseCatCache(m, v1093)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L41
	} else {
		goto L380
	}
L380:
	;
	F_sequence_close(m, v1090, int32(3))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L41
	} else {
		goto L381
	}
L381:
	;
	m.G0 = v1086 - int32(-64)
	goto L354
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1086))) = v1083
	F_errmsg_internal(m, int32(43359), v1086)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L41
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(512484), int32(922), int32(482821))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L41
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1086)+32)) = v1178
	F_errmsg_internal(m, int32(49502), v1084+int32(-32))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L41
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(512484), int32(954), int32(482821))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L41
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1086)+48)) = v1194 + int32(4)
	F_errmsg_internal(m, int32(593002), v1084+int32(-16))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L41
	} else {
		goto L389
	}
L389:
	;
	F_errfinish(m, int32(512484), int32(959), int32(482821))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L41
	} else {
		goto L390
	}
L390:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1086)+16)) = v1083
	F_errmsg_internal(m, int32(383716), v1084+int32(-48))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L41
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(512484), int32(982), int32(482821))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L41
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	F_ScanKeyInit(m, v1226+int32(32), int32(1), int32(3), int32(184), v1223)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L41
	} else {
		goto L395
	}
L395:
	;
	v1240 = int32(1)
	v1245 = F_systable_beginscan(m, v1230, int32(2657), v1240, int32(0), v1240, v1226+int32(32))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L41
	} else {
		goto L398
	}
L396:
	;
	goto L3
L397:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L41
	} else {
		goto L416
	}
L398:
	;
	v1247 = F_systable_getnext(m, v1245)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L41
	} else {
		goto L399
	}
L399:
	;
	if v1247 != 0 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+16))
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249)+22)))
	v1251 = v1249 + v1250
	v1252 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1251)+8)))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+4))
	v1255 = F_relation_open(m, v1253, int32(8))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L41
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L41
	} else {
		goto L413
	}
L403:
	;
	F_CatalogTupleDelete(m, v1230, v1247+int32(4))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L41
	} else {
		goto L404
	}
L404:
	;
	F_systable_endscan(m, v1245)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L41
	} else {
		goto L405
	}
L405:
	;
	F_sequence_close(m, v1230, int32(3))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L41
	} else {
		goto L406
	}
L406:
	;
	v1268 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L41
	} else {
		goto L407
	}
L407:
	;
	v1271 = F_SearchSysCacheCopy(m, int32(7), v1253, v1252)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L41
	} else {
		goto L408
	}
L408:
	;
	if v1271 == int32(0) {
		goto L397
	} else {
		goto L409
	}
L409:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+16))
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275)+22)))
	v1278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1275+v1276)+87)) = uint8(v1278)
	F_CatalogTupleUpdate(m, v1268, v1271+int32(4), v1271)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L41
	} else {
		goto L410
	}
L410:
	;
	F_sequence_close(m, v1268, int32(3))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L41
	} else {
		goto L411
	}
L411:
	;
	F_relation_close(m, v1255, int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L41
	} else {
		goto L412
	}
L412:
	;
	m.G0 = v1226 + int32(80)
	goto L396
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226))) = v1223
	F_errmsg_internal(m, int32(53244), v1226)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L41
	} else {
		goto L414
	}
L414:
	;
	F_errfinish(m, int32(518742), int32(232), int32(482863))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L41
	} else {
		goto L415
	}
L415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+20)) = v1253
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+16)) = v1252
	F_errmsg_internal(m, int32(49690), v1226+int32(16))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L41
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(518742), int32(254), int32(482863))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L41
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	v1333 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L41
	} else {
		goto L420
	}
L420:
	;
	F_ScanKeyInit(m, v1323+int32(-48), int32(1), int32(3), int32(184), v1322)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L41
	} else {
		goto L421
	}
L421:
	;
	v1343 = int32(1)
	v1348 = F_systable_beginscan(m, v1329, int32(2996), v1343, int32(0), v1343, v1323+int32(-48))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L41
	} else {
		goto L423
	}
L422:
	;
	goto L3
L423:
	;
	v1350 = F_systable_getnext(m, v1348)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L41
	} else {
		goto L424
	}
L424:
	;
	if v1350 != 0 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	F_CatalogTupleDelete(m, v1329, v1350+int32(4))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L41
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L41
	} else {
		goto L442
	}
L428:
	;
	F_systable_endscan(m, v1348)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L41
	} else {
		goto L429
	}
L429:
	;
	F_ScanKeyInit(m, v1323+int32(-48), int32(1), int32(3), int32(184), v1322)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L41
	} else {
		goto L430
	}
L430:
	;
	v1366 = int32(1)
	v1371 = F_systable_beginscan(m, v1333, int32(2683), v1366, int32(0), v1366, v1323+int32(-48))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L41
	} else {
		goto L431
	}
L431:
	;
	goto L432
L432:
	;
	v1385 = F_systable_getnext(m, v1371)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L41
	} else {
		goto L434
	}
L433:
	;
	F_systable_endscan(m, v1371)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L41
	} else {
		goto L439
	}
L434:
	;
	if v1385 != 0 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	F_CatalogTupleDelete(m, v1333, v1385+int32(4))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L41
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	goto L433
L438:
	;
	goto L432
L439:
	;
	F_sequence_close(m, v1333, int32(3))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L41
	} else {
		goto L440
	}
L440:
	;
	F_sequence_close(m, v1329, int32(3))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L41
	} else {
		goto L441
	}
L441:
	;
	m.G0 = v1325 - int32(-64)
	goto L422
L442:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L41
	} else {
		goto L443
	}
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1325))) = v1322
	F_errmsg(m, int32(74155), v1325)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L41
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(512862), int32(125), int32(244385))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L41
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	v1428 = F_SearchSysCache1(m, int32(40), v1418)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L41
	} else {
		goto L449
	}
L447:
	;
	goto L3
L448:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L41
	} else {
		goto L469
	}
L449:
	;
	if v1428 != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+16))
	v1431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1430)+22)))
	v1432 = v1430 + v1431
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+92))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+96))
	if v1433|v1434 == int32(0) {
		v1452 = v1428
		goto L453
	} else {
		goto L454
	}
L451:
	;
	goto L452
L452:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L41
	} else {
		goto L466
	}
L453:
	;
	F_CatalogTupleDelete(m, v1425, v1452+int32(4))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L41
	} else {
		goto L463
	}
L454:
	;
	F_OperatorUpd(m, v1418, v1433, v1434, int32(1))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L41
	} else {
		goto L455
	}
L455:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+92))
	if v1441 != v1418 {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+96))
	if v1418 != v1443 {
		v1452 = v1428
		goto L453
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	F_ReleaseCatCache(m, v1428)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L41
	} else {
		goto L460
	}
L459:
	;
	goto L458
L460:
	;
	v1448 = F_SearchSysCache1(m, int32(40), v1418)
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L41
	} else {
		goto L461
	}
L461:
	;
	if v1448 == int32(0) {
		goto L448
	} else {
		goto L462
	}
L462:
	;
	v1452 = v1448
	goto L453
L463:
	;
	F_ReleaseCatCache(m, v1452)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L41
	} else {
		goto L464
	}
L464:
	;
	F_sequence_close(m, v1425, int32(3))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L41
	} else {
		goto L465
	}
L465:
	;
	m.G0 = v1421 + int32(32)
	goto L447
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1421))) = v1418
	F_errmsg_internal(m, int32(46326), v1421)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L41
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(514131), int32(456), int32(482921))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L41
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1421)+16)) = v1418
	F_errmsg_internal(m, int32(46326), v1421+int32(16))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L41
	} else {
		goto L470
	}
L470:
	;
	F_errfinish(m, int32(514131), int32(473), int32(482921))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L41
	} else {
		goto L471
	}
L471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L472:
	;
	F_ScanKeyInit(m, v1496+int32(32), int32(1), int32(3), int32(184), v1493)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L41
	} else {
		goto L473
	}
L473:
	;
	v1510 = int32(1)
	v1515 = F_systable_beginscan(m, v1500, int32(2692), v1510, int32(0), v1510, v1496+int32(32))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L41
	} else {
		goto L476
	}
L474:
	;
	goto L3
L475:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L41
	} else {
		goto L498
	}
L476:
	;
	v1517 = F_systable_getnext(m, v1515)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L41
	} else {
		goto L477
	}
L477:
	;
	if v1517 != 0 {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+16))
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519)+22)))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1519+v1520)+68))
	v1524 = F_table_open(m, v1522, int32(8))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L41
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L41
	} else {
		goto L495
	}
L481:
	;
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, _consts[223])))
	if v1527 == int32(0) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v1531 = int32(1)
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1524)+56))
	if base.Ui32(v1532) < base.Ui32(int32(12000)) {
		v1541 = v1531
		goto L486
	} else {
		goto L487
	}
L483:
	;
	goto L484
L484:
	;
	F_CatalogTupleDelete(m, v1500, v1517+int32(4))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L41
	} else {
		goto L490
	}
L485:
	;
	if v1541 != 0 {
		goto L475
	} else {
		goto L489
	}
L486:
	;
	goto L485
L487:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1524)+48))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+68))
	if v1536 == int32(99) {
		v1541 = v1531
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v1539 = F_isTempToastNamespace(m, v1536)
	mBase = m.M
	v1541 = v1539
	goto L486
L489:
	;
	goto L484
L490:
	;
	F_systable_endscan(m, v1515)
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L41
	} else {
		goto L491
	}
L491:
	;
	F_sequence_close(m, v1500, int32(3))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L41
	} else {
		goto L492
	}
L492:
	;
	F_CacheInvalidateRelcache(m, v1524)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L41
	} else {
		goto L493
	}
L493:
	;
	F_sequence_close(m, v1524, int32(0))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L41
	} else {
		goto L494
	}
L494:
	;
	m.G0 = v1496 + int32(80)
	goto L474
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1496))) = v1493
	F_errmsg_internal(m, int32(56051), v1496)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L41
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(518838), int32(61), int32(483105))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L41
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L41
	} else {
		goto L499
	}
L499:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1524)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+16)) = v1579 + int32(4)
	F_errmsg(m, int32(341295), v1496+int32(16))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L41
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(518838), int32(75), int32(483105))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L41
	} else {
		goto L501
	}
L501:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L502:
	;
	F_ScanKeyInit(m, v1596+int32(48), int32(1), int32(3), int32(184), v1593)
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L41
	} else {
		goto L503
	}
L503:
	;
	v1610 = int32(1)
	v1615 = F_systable_beginscan(m, v1600, int32(2702), v1610, int32(0), v1610, v1596+int32(48))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L41
	} else {
		goto L507
	}
L504:
	;
	goto L3
L505:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L41
	} else {
		goto L536
	}
L506:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L41
	} else {
		goto L531
	}
L507:
	;
	v1617 = F_systable_getnext(m, v1615)
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L41
	} else {
		goto L508
	}
L508:
	;
	if v1617 != 0 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+16))
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619)+22)))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1619+v1620)+4))
	v1624 = F_table_open(m, v1622, int32(8))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L41
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L41
	} else {
		goto L528
	}
L512:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+48))
	v1627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+119)))
	v1629 = v1627 - int32(102)
	v1638 = (v1629<<(uint(int32(7))%32) | int32(base.Ui32(v1629&int32(254))>>(uint(int32(1))%32))) & int32(255)
	if base.Ui32(int32(8)) < base.Ui32(v1638) {
		goto L506
	} else {
		goto L513
	}
L513:
	;
	if int32(1)<<(uint(v1638)%32)&int32(353) == int32(0) {
		goto L506
	} else {
		goto L514
	}
L514:
	;
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, _consts[223])))
	if v1648 == int32(0) {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	v1652 = int32(1)
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+56))
	if base.Ui32(v1653) < base.Ui32(int32(12000)) {
		v1662 = v1652
		goto L519
	} else {
		goto L520
	}
L516:
	;
	goto L517
L517:
	;
	F_CatalogTupleDelete(m, v1600, v1617+int32(4))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L41
	} else {
		goto L523
	}
L518:
	;
	if v1662 != 0 {
		goto L505
	} else {
		goto L522
	}
L519:
	;
	goto L518
L520:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+48))
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+68))
	if v1657 == int32(99) {
		v1662 = v1652
		goto L519
	} else {
		goto L521
	}
L521:
	;
	v1660 = F_isTempToastNamespace(m, v1657)
	mBase = m.M
	v1662 = v1660
	goto L519
L522:
	;
	goto L517
L523:
	;
	F_systable_endscan(m, v1615)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L41
	} else {
		goto L524
	}
L524:
	;
	F_sequence_close(m, v1600, int32(3))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L41
	} else {
		goto L525
	}
L525:
	;
	F_CacheInvalidateRelcache(m, v1624)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L41
	} else {
		goto L526
	}
L526:
	;
	F_sequence_close(m, v1624, int32(0))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L41
	} else {
		goto L527
	}
L527:
	;
	m.G0 = v1596 + int32(96)
	goto L504
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1596))) = v1593
	F_errmsg_internal(m, int32(47119), v1596)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L41
	} else {
		goto L529
	}
L529:
	;
	F_errfinish(m, int32(515472), int32(1316), int32(482940))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L41
	} else {
		goto L530
	}
L530:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L531:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L41
	} else {
		goto L532
	}
L532:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1596)+16)) = v1700 + int32(4)
	F_errmsg(m, int32(142100), v1596+int32(16))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L41
	} else {
		goto L533
	}
L533:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+48))
	v1710 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1709)+119)))
	F_errdetail_relkind_not_supported(m, v1710)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L41
	} else {
		goto L534
	}
L534:
	;
	F_errfinish(m, int32(515472), int32(1333), int32(482940))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L41
	} else {
		goto L535
	}
L535:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L536:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L41
	} else {
		goto L537
	}
L537:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1596)+32)) = v1725 + int32(4)
	F_errmsg(m, int32(341295), v1596+int32(32))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L41
	} else {
		goto L538
	}
L538:
	;
	F_errfinish(m, int32(515472), int32(1339), int32(482940))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L41
	} else {
		goto L539
	}
L539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L540:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1742 = m.G0
	v1744 = v1742 - int32(16)
	m.G0 = v1744
	v1748 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L41
	} else {
		goto L541
	}
L541:
	;
	v1751 = F_SearchSysCache1(m, int32(64), v1741)
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L41
	} else {
		goto L543
	}
L542:
	;
	goto L3
L543:
	;
	if v1751 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+16))
	v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1753)+22)))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1753+v1754)+4))
	v1758 = F_table_open(m, v1756, int32(4))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L41
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L41
	} else {
		goto L569
	}
L547:
	;
	v1762 = F_table_open(m, int32(3429), int32(3))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L41
	} else {
		goto L548
	}
L548:
	;
	v1766 = F_SearchSysCache2(m, int32(62), v1741, int32(1))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L41
	} else {
		goto L549
	}
L549:
	;
	if v1766 != 0 {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	F_CatalogTupleDelete(m, v1762, v1766+int32(4))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L41
	} else {
		goto L553
	}
L551:
	;
	goto L552
L552:
	;
	F_sequence_close(m, v1762, int32(3))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L41
	} else {
		goto L555
	}
L553:
	;
	F_ReleaseCatCache(m, v1766)
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L41
	} else {
		goto L554
	}
L554:
	;
	goto L552
L555:
	;
	v1779 = F_table_open(m, int32(3429), int32(3))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L41
	} else {
		goto L556
	}
L556:
	;
	v1783 = F_SearchSysCache2(m, int32(62), v1741, int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L41
	} else {
		goto L557
	}
L557:
	;
	if v1783 != 0 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	F_CatalogTupleDelete(m, v1779, v1783+int32(4))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L41
	} else {
		goto L561
	}
L559:
	;
	goto L560
L560:
	;
	F_sequence_close(m, v1779, int32(3))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L41
	} else {
		goto L563
	}
L561:
	;
	F_ReleaseCatCache(m, v1783)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L41
	} else {
		goto L562
	}
L562:
	;
	goto L560
L563:
	;
	F_CacheInvalidateRelcacheByRelid(m, v1756)
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L41
	} else {
		goto L564
	}
L564:
	;
	F_CatalogTupleDelete(m, v1748, v1751+int32(4))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L41
	} else {
		goto L565
	}
L565:
	;
	F_ReleaseCatCache(m, v1751)
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L41
	} else {
		goto L566
	}
L566:
	;
	F_sequence_close(m, v1758, int32(0))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L41
	} else {
		goto L567
	}
L567:
	;
	F_sequence_close(m, v1748, int32(3))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L41
	} else {
		goto L568
	}
L568:
	;
	m.G0 = v1744 + int32(16)
	goto L542
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1744))) = v1741
	F_errmsg_internal(m, int32(45079), v1744)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L41
	} else {
		goto L570
	}
L570:
	;
	F_errfinish(m, int32(514105), int32(803), int32(482900))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L41
	} else {
		goto L571
	}
L571:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L572:
	;
	v1834 = F_SearchSysCache1(m, int32(74), v1824)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L41
	} else {
		goto L574
	}
L573:
	;
	goto L3
L574:
	;
	if v1834 != 0 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	F_CatalogTupleDelete(m, v1831, v1834+int32(4))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L41
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L41
	} else {
		goto L593
	}
L578:
	;
	F_ReleaseCatCache(m, v1834)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L41
	} else {
		goto L579
	}
L579:
	;
	F_sequence_close(m, v1831, int32(3))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L41
	} else {
		goto L580
	}
L580:
	;
	v1847 = F_table_open(m, int32(3603), int32(3))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L41
	} else {
		goto L581
	}
L581:
	;
	F_ScanKeyInit(m, v1825+int32(-48), int32(1), int32(3), int32(184), v1824)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L41
	} else {
		goto L582
	}
L582:
	;
	v1857 = int32(1)
	v1862 = F_systable_beginscan(m, v1847, int32(3609), v1857, int32(0), v1857, v1825+int32(-48))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L41
	} else {
		goto L583
	}
L583:
	;
	goto L584
L584:
	;
	v1876 = F_systable_getnext(m, v1862)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L41
	} else {
		goto L586
	}
L585:
	;
	F_systable_endscan(m, v1862)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L41
	} else {
		goto L591
	}
L586:
	;
	if v1876 != 0 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	F_CatalogTupleDelete(m, v1847, v1876+int32(4))
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L41
	} else {
		goto L590
	}
L588:
	;
	goto L589
L589:
	;
	goto L585
L590:
	;
	goto L584
L591:
	;
	F_sequence_close(m, v1847, int32(3))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L41
	} else {
		goto L592
	}
L592:
	;
	m.G0 = v1827 - int32(-64)
	goto L573
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1827))) = v1824
	F_errmsg_internal(m, int32(41852), v1827)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L41
	} else {
		goto L594
	}
L594:
	;
	F_errfinish(m, int32(514289), int32(1123), int32(482977))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L41
	} else {
		goto L595
	}
L595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L596:
	;
	goto L3
L597:
	;
	v1913 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L41
	} else {
		goto L600
	}
L598:
	;
	goto L599
L599:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L41
	} else {
		goto L610
	}
L600:
	;
	F_ScanKeyInit(m, v1904+int32(-48), int32(1), int32(3), int32(184), v1903)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L41
	} else {
		goto L601
	}
L601:
	;
	v1923 = int32(1)
	v1928 = F_systable_beginscan(m, v1913, int32(3080), v1923, int32(0), v1923, v1904+int32(-48))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L41
	} else {
		goto L602
	}
L602:
	;
	v1930 = F_systable_getnext(m, v1928)
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L41
	} else {
		goto L603
	}
L603:
	;
	if v1930 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	F_CatalogTupleDelete(m, v1913, v1930+int32(4))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L41
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	F_systable_endscan(m, v1928)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L41
	} else {
		goto L608
	}
L607:
	;
	goto L606
L608:
	;
	F_sequence_close(m, v1913, int32(3))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L41
	} else {
		goto L609
	}
L609:
	;
	m.G0 = v1906 - int32(-64)
	goto L596
L610:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L41
	} else {
		goto L611
	}
L611:
	;
	v1951 = F_get_extension_name(m, v1903)
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L41
	} else {
		goto L612
	}
L612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1906))) = v1951
	F_errmsg(m, int32(474422), v1906)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L41
	} else {
		goto L613
	}
L613:
	;
	F_errfinish(m, int32(516462), int32(2302), int32(483025))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L41
	} else {
		goto L614
	}
L614:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L615:
	;
	v1972 = F_SearchSysCache1(m, int32(49), v1962)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L41
	} else {
		goto L617
	}
L616:
	;
	goto L3
L617:
	;
	if v1972 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+16))
	v1975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1974)+22)))
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1974+v1975)+8))
	v1979 = F_GetSchemaPublicationRelations(m, v1977, int32(2))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L41
	} else {
		goto L622
	}
L619:
	;
	goto L620
L620:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L41
	} else {
		goto L636
	}
L621:
	;
	F_CatalogTupleDelete(m, v1969, v1972+int32(4))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L41
	} else {
		goto L633
	}
L622:
	;
	if v1979 == int32(0) {
		goto L621
	} else {
		goto L623
	}
L623:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1979)+4))
	if v1983 <= int32(4095) {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	if v1983 <= int32(0) {
		goto L621
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L41
	} else {
		goto L632
	}
L627:
	;
	v1989 = int32(0)
	goto L628
L628:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1979)+12))
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v2001+v1989<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v2005)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L41
	} else {
		goto L630
	}
L629:
	;
	goto L621
L630:
	;
	v2009 = v1989 + int32(1)
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1979)+4))
	if v2009 < v2010 {
		v1989 = v2009
		goto L628
	} else {
		goto L631
	}
L631:
	;
	goto L629
L632:
	;
	goto L621
L633:
	;
	F_ReleaseCatCache(m, v1972)
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L41
	} else {
		goto L634
	}
L634:
	;
	F_sequence_close(m, v1969, int32(3))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L41
	} else {
		goto L635
	}
L635:
	;
	m.G0 = v1965 + int32(16)
	goto L616
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1965))) = v1962
	F_errmsg_internal(m, int32(60160), v1965)
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L41
	} else {
		goto L637
	}
L637:
	;
	F_errfinish(m, int32(514207), int32(1635), int32(483127))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L41
	} else {
		goto L638
	}
L638:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L639:
	;
	v2061 = F_SearchSysCache1(m, int32(52), v2051)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L41
	} else {
		goto L641
	}
L640:
	;
	goto L3
L641:
	;
	if v2061 != 0 {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+16))
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+22)))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2065+v2066)+8))
	v2069 = F_GetPubPartitionOptionRelations(m, int32(0), int32(2), v2068)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L41
	} else {
		goto L646
	}
L643:
	;
	goto L644
L644:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L41
	} else {
		goto L660
	}
L645:
	;
	F_CatalogTupleDelete(m, v2058, v2061+int32(4))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L41
	} else {
		goto L657
	}
L646:
	;
	if v2069 == int32(0) {
		goto L645
	} else {
		goto L647
	}
L647:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+4))
	if v2073 <= int32(4095) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	if v2073 <= int32(0) {
		goto L645
	} else {
		goto L651
	}
L649:
	;
	goto L650
L650:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L41
	} else {
		goto L656
	}
L651:
	;
	v2079 = int32(0)
	goto L652
L652:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+12))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2091+v2079<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v2095)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L41
	} else {
		goto L654
	}
L653:
	;
	goto L645
L654:
	;
	v2099 = v2079 + int32(1)
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+4))
	if v2099 < v2100 {
		v2079 = v2099
		goto L652
	} else {
		goto L655
	}
L655:
	;
	goto L653
L656:
	;
	goto L645
L657:
	;
	F_ReleaseCatCache(m, v2061)
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L41
	} else {
		goto L658
	}
L658:
	;
	F_sequence_close(m, v2058, int32(3))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L41
	} else {
		goto L659
	}
L659:
	;
	m.G0 = v2054 + int32(16)
	goto L640
L660:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2054))) = v2051
	F_errmsg_internal(m, int32(56211), v2054)
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L41
	} else {
		goto L661
	}
L661:
	;
	F_errfinish(m, int32(514207), int32(1566), int32(483045))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L41
	} else {
		goto L662
	}
L662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L663:
	;
	v2151 = F_SearchSysCache1(m, int32(51), v2141)
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L41
	} else {
		goto L665
	}
L664:
	;
	goto L3
L665:
	;
	if v2151 != 0 {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+16))
	v2154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2153)+22)))
	v2156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2153+v2154)+72)))
	if v2156 == int32(1) {
		goto L669
	} else {
		goto L670
	}
L667:
	;
	goto L668
L668:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L41
	} else {
		goto L676
	}
L669:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L41
	} else {
		goto L672
	}
L670:
	;
	goto L671
L671:
	;
	F_CatalogTupleDelete(m, v2148, v2151+int32(4))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L41
	} else {
		goto L673
	}
L672:
	;
	goto L671
L673:
	;
	F_ReleaseCatCache(m, v2151)
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L41
	} else {
		goto L674
	}
L674:
	;
	F_sequence_close(m, v2148, int32(3))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L41
	} else {
		goto L675
	}
L675:
	;
	m.G0 = v2144 + int32(16)
	goto L664
L676:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2144))) = v2141
	F_errmsg_internal(m, int32(49980), v2144)
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L41
	} else {
		goto L677
	}
L677:
	;
	F_errfinish(m, int32(514207), int32(1604), int32(483003))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L41
	} else {
		goto L678
	}
L678:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L679:
	;
	F_errmsg_internal(m, int32(261513), int32(0))
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L41
	} else {
		goto L680
	}
L680:
	;
	F_errfinish(m, int32(511542), int32(1478), int32(261549))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L41
	} else {
		goto L681
	}
L681:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L682:
	;
	goto L6
L683:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v2205
	F_errmsg_internal(m, int32(62593), v15)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L41
	} else {
		goto L684
	}
L684:
	;
	F_errfinish(m, int32(511542), int32(1482), int32(261549))
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L41
	} else {
		goto L685
	}
L685:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L686:
	;
	v2225 = F_SearchSysCache1(m, int32(47), v2215)
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L41
	} else {
		goto L689
	}
L687:
	;
	goto L3
L688:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L41
	} else {
		goto L709
	}
L689:
	;
	if v2225 != 0 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2225)+16))
	v2228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2227)+22)))
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2227+v2228)+96)))
	F_CatalogTupleDelete(m, v2222, v2225+int32(4))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L41
	} else {
		goto L693
	}
L691:
	;
	goto L692
L692:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L41
	} else {
		goto L706
	}
L693:
	;
	F_ReleaseCatCache(m, v2225)
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L41
	} else {
		goto L694
	}
L694:
	;
	F_sequence_close(m, v2222, int32(3))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L41
	} else {
		goto L695
	}
L695:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	F_pgstat_drop_transactional(m, int32(3), v2242, base.I64_extend_i32_u(v2215))
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L41
	} else {
		goto L696
	}
L696:
	;
	if v2230 == int32(97) {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	v2250 = F_table_open(m, int32(2600), int32(3))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L41
	} else {
		goto L700
	}
L698:
	;
	goto L699
L699:
	;
	m.G0 = v2218 + int32(32)
	goto L687
L700:
	;
	v2253 = F_SearchSysCache1(m, int32(0), v2215)
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L41
	} else {
		goto L701
	}
L701:
	;
	if v2253 == int32(0) {
		goto L688
	} else {
		goto L702
	}
L702:
	;
	F_CatalogTupleDelete(m, v2250, v2253+int32(4))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L41
	} else {
		goto L703
	}
L703:
	;
	F_ReleaseCatCache(m, v2253)
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L41
	} else {
		goto L704
	}
L704:
	;
	F_sequence_close(m, v2250, int32(3))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L41
	} else {
		goto L705
	}
L705:
	;
	goto L699
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2218))) = v2215
	F_errmsg_internal(m, int32(47929), v2218)
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L41
	} else {
		goto L707
	}
L707:
	;
	F_errfinish(m, int32(514176), int32(1324), int32(482958))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L41
	} else {
		goto L708
	}
L708:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L709:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2218)+16)) = v2215
	F_errmsg_internal(m, int32(47870), v2218+int32(16))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L41
	} else {
		goto L710
	}
L710:
	;
	F_errfinish(m, int32(514176), int32(1345), int32(482958))
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L41
	} else {
		goto L711
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L712:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2303 = F_table_open(m, v2301, int32(3))
	mBase = m.M
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L41
	} else {
		goto L713
	}
L713:
	;
	if int32(0) <= v2299 {
		goto L715
	} else {
		goto L716
	}
L714:
	;
	F_sequence_close(m, v2303, int32(3))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L41
	} else {
		goto L730
	}
L715:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2308 = F_SearchSysCache1(m, v2299, v2307)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L41
	} else {
		goto L718
	}
L716:
	;
	goto L717
L717:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2321 = F_get_object_attnum_oid(m, v2320)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L41
	} else {
		goto L722
	}
L718:
	;
	if v2308 == int32(0) {
		goto L2
	} else {
		goto L719
	}
L719:
	;
	F_CatalogTupleDelete(m, v2303, v2308+int32(4))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L41
	} else {
		goto L720
	}
L720:
	;
	F_ReleaseCatCache(m, v2308)
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L41
	} else {
		goto L721
	}
L721:
	;
	goto L714
L722:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v15+int32(48), v2321, int32(3), int32(184), v2325)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L41
	} else {
		goto L723
	}
L723:
	;
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2329 = F_get_object_oid_index(m, v2328)
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L41
	} else {
		goto L724
	}
L724:
	;
	v2331 = int32(1)
	v2336 = F_systable_beginscan(m, v2303, v2329, v2331, int32(0), v2331, v15+int32(48))
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L41
	} else {
		goto L725
	}
L725:
	;
	v2338 = F_systable_getnext(m, v2336)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L41
	} else {
		goto L726
	}
L726:
	;
	if v2338 == int32(0) {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	F_CatalogTupleDelete(m, v2303, v2338+int32(4))
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L41
	} else {
		goto L728
	}
L728:
	;
	F_systable_endscan(m, v2336)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L41
	} else {
		goto L729
	}
L729:
	;
	goto L714
L730:
	;
	goto L3
L731:
	;
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2373 = F_get_object_class_descr(m, v2372)
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L41
	} else {
		goto L732
	}
L732:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v2375
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v2373
	F_errmsg_internal(m, int32(45831), v15+int32(16))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L41
	} else {
		goto L733
	}
L733:
	;
	F_errfinish(m, int32(511542), int32(1207), int32(482885))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L41
	} else {
		goto L734
	}
L734:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L735:
	;
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2393 = F_get_object_class_descr(m, v2392)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L41
	} else {
		goto L736
	}
L736:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v2395
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v2393
	F_errmsg_internal(m, int32(45800), v15+int32(32))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L41
	} else {
		goto L737
	}
L737:
	;
	F_errfinish(m, int32(511542), int32(1230), int32(482885))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L41
	} else {
		goto L738
	}
L738:
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var __phi340 int32
	_ = __phi340
	var v341 int32
	_ = v341
	var __phi341 int32
	_ = __phi341
	var v349 int32
	_ = v349
	var __phi349 int32
	_ = __phi349
	var v351 int32
	_ = v351
	var __phi351 int32
	_ = __phi351
	var v356 int32
	_ = v356
	var __phi356 int32
	_ = __phi356
	var v359 int32
	_ = v359
	var __phi359 int32
	_ = __phi359
	var v363 int32
	_ = v363
	var __phi363 int32
	_ = __phi363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int64
	_ = v490
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int64
	_ = v502
	var v504 int32
	_ = v504
	var v505 int64
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int64
	_ = v533
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 float64
	_ = v709
	var v710 int64
	_ = v710
	var v718 int32
	_ = v718
	var v729 int32
	_ = v729
	var v732 float64
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v742 int64
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v970 int32
	_ = v970
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1035 int32
	_ = v1035
	var v1036 int64
	_ = v1036
	var v1048 int32
	_ = v1048
	var v1049 float64
	_ = v1049
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	v4 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(1360)
	m.G0 = v33
	v36 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v38 = l1
	v39 = l2
	v44 = v4
	v60 = v4
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
	m.G0 = v33 + int32(1360)
	return
L3:
	;
	goto L2
L4:
	;
	if v60 != 0 {
		goto L37
	} else {
		goto L38
	}
L5:
	;
	v178 = v38
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v67 == int32(0) {
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
	F_dostr(m, v38, v160-v38, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	goto L9
L11:
	;
	v150 = v145
	goto L29
L12:
	;
	v145 = v137
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
	v84 = v73
	goto L19
L17:
	;
	v97 = v73
	goto L18
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v106 = int32(-2139062144)
	if (int32(16843008)-v103|v103)&v106 != v106 {
		v137 = v97
		goto L12
	} else {
		goto L24
	}
L19:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 == int32(0) {
		v160 = v84
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v97 = v94
	goto L18
L21:
	;
	if int32(37) == v89 {
		v160 = v84
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v94 = v84 + int32(1)
	if v94&int32(3) != 0 {
		v84 = v94
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v112 = v97
	v114 = v103
	goto L25
L25:
	;
	v118 = v114 ^ int32(623191333)
	v121 = int32(-2139062144)
	if (int32(16843008)-v118|v118)&v121 != v121 {
		v137 = v112
		goto L12
	} else {
		goto L27
	}
L26:
	;
	v145 = v127
	goto L11
L27:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v127 = v112 + int32(4)
	v131 = int32(-2139062144)
	if (v125|(int32(16843008)-v125))&v131 == v131 {
		v112 = v127
		v114 = v125
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v152 == int32(0) {
		v160 = v150
		goto L10
	} else {
		goto L31
	}
L30:
	;
	v160 = v150
	goto L10
L31:
	;
	if v152 != int32(37) {
		v150 = v150 + int32(1)
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	return
L34:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v174 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v175 == int32(0) {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v178 = v160
	goto L4
L37:
	;
	v179 = v60
	goto L39
L38:
	;
	v179 = v178
	goto L39
L39:
	;
	v181 = v178 + int32(1)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v182 != int32(115) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v185 = int32(0)
	v199 = v181
	v200 = v39
	v202 = v182
	v204 = v185
	v205 = v44
	v207 = v185
	v208 = v185
	v209 = v185
	v211 = v185
	v213 = v185
	v214 = v185
	v215 = v185
	v216 = v185
	v218 = v185
	v219 = v185
	v222 = v185
	v223 = v185
	goto L43
L41:
	;
	goto L42
L42:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v1227 != 0 {
		goto L416
	} else {
		goto L417
	}
L43:
	;
	v228 = int32(1)
	v230 = v199 + v228
	v231 = base.I32_extend8_s(v202)
	switch v202&int32(255) - int32(36) {
	case 0:
		goto L63
	case 1:
		goto L53
	default:
		goto L49
	case 3, 68:
		v1163 = v214
		v1164 = v223
		goto L47
	case 6:
		goto L64
	case 7:
		goto L68
	case 9:
		v1198 = v200
		v1199 = v228
		v1202 = v204
		v1203 = v205
		v1205 = v207
		v1209 = v211
		v1211 = v213
		v1212 = v214
		v1213 = v215
		v1214 = v216
		v1216 = v218
		v1217 = v219
		v1220 = v222
		v1221 = v223
		goto L45
	case 10:
		goto L65
	case 12:
		goto L67
	case 13, 14, 15, 16, 17, 18, 19, 20, 21:
		v240 = v213
		goto L66
	case 33, 35, 65, 66, 67:
		goto L55
	case 52, 75, 81, 84:
		goto L59
	case 63:
		goto L58
	case 64, 69:
		goto L60
	case 72:
		goto L62
	case 73:
		goto L54
	case 76:
		goto L56
	case 79:
		goto L57
	case 86:
		goto L61
	}
L45:
	;
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v199 = v230
	v200 = v1198
	v202 = v1226
	v204 = v1202
	v205 = v1203
	v207 = v1205
	v208 = v1202
	v209 = v1199
	v211 = v1209
	v213 = v1211
	v214 = v1212
	v215 = v1213
	v216 = v1214
	v218 = v1216
	v219 = v1217
	v222 = v1220
	v223 = v1221
	goto L43
L46:
	;
	v1198 = v200
	v1199 = v209
	v1202 = v1195
	v1203 = v1172
	v1205 = v1174
	v1209 = v1178
	v1211 = v1180
	v1212 = v1181
	v1213 = v1182
	v1214 = v1183
	v1216 = v218
	v1217 = v1186
	v1220 = v1189
	v1221 = v1190
	goto L45
L47:
	;
	v1172 = v205
	v1174 = v207
	v1178 = v211
	v1180 = v213
	v1181 = v1163
	v1182 = v215
	v1183 = v216
	v1186 = v219
	v1189 = v222
	v1190 = v1164
	v1195 = v208
	goto L46
L48:
	;
	v1198 = v200 + int32(4)
	v1199 = v1157
	v1202 = v1158
	v1203 = int32(0)
	v1205 = v1159
	v1209 = v1160
	v1211 = v213
	v1212 = v214
	v1213 = v1161
	v1214 = v250
	v1216 = v218
	v1217 = v219
	v1220 = v222
	v1221 = v223
	goto L45
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(28)
	v1155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v1155)
	goto L3
L50:
	;
	v1096 = int32(1)
	v1097 = int32(0)
	if v222 == v1097 {
		goto L404
	} else {
		goto L405
	}
L51:
	;
	if v294 <= int32(0) {
		goto L50
	} else {
		goto L393
	}
L52:
	;
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v970 == int32(0) {
		v38 = v230
		v39 = v960
		v44 = v205
		v60 = v179
		goto L1
	} else {
		goto L392
	}
L53:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v924 == int32(0) {
		v952 = v923
		goto L381
	} else {
		goto L382
	}
L54:
	;
	v918 = F_pg_strerror_r(m, v36, v33+int32(320))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L33
	} else {
		goto L379
	}
L55:
	;
	v698 = (v200 + int32(7)) & int32(-8)
	v700 = v698 + int32(8)
	v707 = v205 & int32(1)
	if v707 != 0 {
		goto L296
	} else {
		goto L297
	}
L56:
	;
	v673 = v205 & int32(1)
	if v673 != 0 {
		goto L282
	} else {
		goto L283
	}
L57:
	;
	if v207 != 0 {
		goto L242
	} else {
		goto L243
	}
L58:
	;
	v547 = v205 & int32(1)
	if v547 != 0 {
		goto L206
	} else {
		goto L207
	}
L59:
	;
	if v207 != 0 {
		goto L181
	} else {
		goto L182
	}
L60:
	;
	if v207 != 0 {
		goto L156
	} else {
		goto L157
	}
L61:
	;
	v1163 = v214
	v1164 = int32(1)
	goto L47
L62:
	;
	if v223 != 0 {
		goto L153
	} else {
		goto L154
	}
L63:
	;
	if v205&int32(1) != 0 {
		goto L50
	} else {
		goto L87
	}
L64:
	;
	v250 = int32(1)
	v251 = int32(0)
	if v205&v250 != 0 {
		goto L75
	} else {
		goto L76
	}
L65:
	;
	if v216 != 0 {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	v1172 = v205
	v1174 = v207
	v1178 = v211
	v1180 = v240
	v1181 = v214
	v1182 = v215
	v1183 = v216
	v1186 = v219
	v1189 = v222
	v1190 = v223
	v1195 = v208*int32(10) + v231 - int32(48)
	goto L46
L67:
	;
	if v207|v208 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v1163 = int32(1)
	v1164 = v223
	goto L47
L69:
	;
	v239 = v213
	goto L71
L70:
	;
	v239 = int32(48)
	goto L71
L71:
	;
	v240 = v239
	goto L66
L72:
	;
	v247 = v211
	goto L74
L73:
	;
	v247 = v208
	goto L74
L74:
	;
	v248 = int32(0)
	v1172 = v205
	v1174 = int32(1)
	v1178 = v247
	v1180 = v213
	v1181 = v214
	v1182 = v215
	v1183 = v248
	v1186 = v219
	v1189 = v222
	v1190 = v223
	v1195 = v248
	goto L46
L75:
	;
	v254 = int32(1)
	v1198 = v200
	v1199 = v209
	v1202 = v251
	v1203 = v254
	v1205 = v207
	v1209 = v211
	v1211 = v213
	v1212 = v214
	v1213 = v215
	v1214 = v250
	v1216 = v218
	v1217 = v219
	v1220 = v254
	v1221 = v223
	goto L45
L76:
	;
	goto L77
L77:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v207 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v259 = int32(0)
	v261 = base.B2i32(v259 <= v258)
	if v259 <= v258 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v264 = v258 >> (uint(int32(31)) % 32)
	v267 = int32(0)
	if v258 < v267 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v262 = v258
	goto L83
L82:
	;
	v262 = v259
	goto L83
L83:
	;
	v1157 = v209
	v1158 = v251
	v1159 = v261
	v1160 = v211
	v1161 = v262
	goto L48
L84:
	;
	v271 = int32(1)
	goto L86
L85:
	;
	v271 = v209
	goto L86
L86:
	;
	v1157 = v271
	v1158 = v267
	v1159 = int32(0)
	v1160 = v258 ^ v264 - v264
	v1161 = v215
	goto L48
L87:
	;
	v275 = int32(0)
	v281 = F__emscripten_memset_bulkmem(m, v33+int32(320), base.I32_extend8_s(v275), int32(128))
	mBase = m.M
	goto L88
L88:
	;
	v285 = v179
	v294 = v275
	goto L89
L89:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v312 != int32(37) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L49
L91:
	;
	if v312 == int32(0) {
		goto L51
	} else {
		goto L94
	}
L92:
	;
	v329 = v285
	goto L93
L93:
	;
	v332 = int32(0)
	__phi340 = v329 + int32(1)
	__phi341 = v332
	__phi349 = v294
	__phi351 = v332
	__phi356 = v332
	__phi359 = v332
	__phi363 = v332
	v340 = __phi340
	v341 = __phi341
	v349 = __phi349
	v351 = __phi351
	v356 = __phi356
	v359 = __phi359
	v363 = __phi363
	goto L100
L94:
	;
	v319 = int32(37)
	v320 = F___strchrnul(m, v285+int32(1), v319)
	mBase = m.M
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	if v322 == v319 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v326 == int32(0) {
		goto L51
	} else {
		goto L99
	}
L96:
	;
	v326 = v320
	goto L98
L97:
	;
	v326 = int32(0)
	goto L98
L98:
	;
	goto L95
L99:
	;
	v329 = v326
	goto L93
L100:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	v369 = v340 + int32(1)
	v370 = int32(0)
	switch v367 - int32(36) {
	case 0:
		goto L111
	case 1, 73:
		v465 = v349
		goto L103
	default:
		goto L49
	case 3, 7, 9, 68:
		v404 = v359
		goto L108
	case 6:
		goto L102
	case 10:
		__phi340 = v369
		__phi341 = v370
		v340 = __phi340
		v341 = __phi341
		goto L100
	case 12, 13, 14, 15, 16, 17, 18, 19, 20, 21:
		goto L112
	case 33, 35, 65, 66, 67:
		goto L104
	case 52, 64, 69, 75, 81, 84:
		goto L107
	case 63:
		goto L106
	case 72:
		goto L110
	case 76, 79:
		goto L105
	case 86:
		goto L109
	}
L101:
	;
	goto L90
L102:
	;
	v470 = int32(1)
	if v356&v470 == int32(0) {
		__phi340 = v369
		__phi341 = v370
		__phi356 = v470
		v340 = __phi340
		v341 = __phi341
		v356 = __phi356
		goto L100
	} else {
		goto L152
	}
L103:
	;
	if v356&int32(1) == int32(0) {
		v285 = v369
		v294 = v465
		goto L89
	} else {
		goto L151
	}
L104:
	;
	if v351 == int32(0) {
		goto L49
	} else {
		goto L146
	}
L105:
	;
	if v351 == int32(0) {
		goto L49
	} else {
		goto L141
	}
L106:
	;
	if v351 == int32(0) {
		goto L49
	} else {
		goto L136
	}
L107:
	;
	if v351 == int32(0) {
		goto L49
	} else {
		goto L122
	}
L108:
	;
	__phi340 = v369
	__phi359 = v404
	v340 = __phi340
	v359 = __phi359
	goto L100
L109:
	;
	v404 = int32(1)
	goto L108
L110:
	;
	if v359 != 0 {
		goto L119
	} else {
		goto L120
	}
L111:
	;
	if base.Ui32(v341-int32(32)) < base.Ui32(int32(-31)) {
		goto L49
	} else {
		goto L113
	}
L112:
	;
	__phi340 = v369
	__phi341 = v341*int32(10) + v367 - int32(48)
	v340 = __phi340
	v341 = __phi341
	goto L100
L113:
	;
	v384 = int32(0)
	if v356&int32(1) == v384 {
		__phi340 = v369
		__phi341 = v370
		__phi351 = v341
		__phi356 = v384
		v340 = __phi340
		v341 = __phi341
		v351 = __phi351
		v356 = __phi356
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v391 = v33 + int32(320) + v341<<(uint(int32(2))%32)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	if base.Ui32(int32(1)) < base.Ui32(v392) {
		goto L49
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = int32(1)
	if v341 < v349 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v398 = v349
	goto L118
L117:
	;
	v398 = v341
	goto L118
L118:
	;
	__phi340 = v369
	__phi341 = int32(0)
	__phi349 = v398
	__phi356 = v384
	v340 = __phi340
	v341 = __phi341
	v349 = __phi349
	v356 = __phi356
	goto L100
L119:
	;
	v401 = int32(1)
	goto L121
L120:
	;
	v401 = v363
	goto L121
L121:
	;
	__phi340 = v369
	__phi359 = int32(1)
	__phi363 = v401
	v340 = __phi340
	v359 = __phi359
	v363 = __phi363
	goto L100
L122:
	;
	v409 = int32(2)
	v411 = v33 + int32(320) + v351<<(uint(v409)%32)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	if v359 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v417 = v409
	goto L125
L124:
	;
	v417 = int32(1)
	goto L125
L125:
	;
	if v363 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v418 = int32(3)
	goto L128
L127:
	;
	v418 = v417
	goto L128
L128:
	;
	if v412 != v418 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v420 = v412
	goto L131
L130:
	;
	v420 = int32(0)
	goto L131
L131:
	;
	if v420 != 0 {
		goto L49
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v418
	if v351 < v349 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v423 = v349
	goto L135
L134:
	;
	v423 = v351
	goto L135
L135:
	;
	v465 = v423
	goto L103
L136:
	;
	v430 = v33 + int32(320) + v351<<(uint(int32(2))%32)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if base.Ui32(int32(1)) < base.Ui32(v431) {
		goto L49
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = int32(1)
	if v351 < v349 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v437 = v349
	goto L140
L139:
	;
	v437 = v351
	goto L140
L140:
	;
	v465 = v437
	goto L103
L141:
	;
	v444 = v33 + int32(320) + v351<<(uint(int32(2))%32)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	switch v445 {
	case 0, 5:
		goto L142
	default:
		goto L49
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = int32(5)
	if v351 < v349 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v449 = v349
	goto L145
L144:
	;
	v449 = v351
	goto L145
L145:
	;
	v465 = v449
	goto L103
L146:
	;
	v456 = v33 + int32(320) + v351<<(uint(int32(2))%32)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	switch v457 {
	case 0, 4:
		goto L147
	default:
		goto L49
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = int32(4)
	if v351 < v349 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v461 = v349
	goto L150
L149:
	;
	v461 = v351
	goto L150
L150:
	;
	v465 = v461
	goto L103
L151:
	;
	goto L49
L152:
	;
	goto L101
L153:
	;
	v476 = int32(1)
	goto L155
L154:
	;
	v476 = v219
	goto L155
L155:
	;
	v1172 = v205
	v1174 = v207
	v1178 = v211
	v1180 = v213
	v1181 = v214
	v1182 = v215
	v1183 = v216
	v1186 = v476
	v1189 = v222
	v1190 = int32(1)
	v1195 = v208
	goto L46
L156:
	;
	v479 = v208
	goto L158
L157:
	;
	v479 = v215
	goto L158
L158:
	;
	if v216 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v480 = v215
	goto L161
L160:
	;
	v480 = v479
	goto L161
L161:
	;
	if v207 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v481 = v211
	goto L164
L163:
	;
	v481 = v208
	goto L164
L164:
	;
	if v216 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v482 = v211
	goto L167
L166:
	;
	v482 = v481
	goto L167
L167:
	;
	if v205&int32(1) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v489 = v33 + int32(48) + v218<<(uint(int32(3))%32)
	if v219 != 0 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	if v219 != 0 {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v489)))
	F_fmtint(m, v490, v231, v214, v209, v482, v213, v480, v207, l0)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L33
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v493 = int64(*(*int32)(unsafe.Add(mBase, uint32(v489))))
	F_fmtint(m, v493, v231, v214, v209, v482, v213, v480, v207, l0)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L33
	} else {
		goto L175
	}
L174:
	;
	v960 = v200
	goto L52
L175:
	;
	v960 = v200
	goto L52
L176:
	;
	v499 = (v200 + int32(7)) & int32(-8)
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v499)))
	F_fmtint(m, v502, v231, v214, v209, v482, v213, v480, v207, l0)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L33
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v505 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200))))
	F_fmtint(m, v505, v231, v214, v209, v482, v213, v480, v207, l0)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L33
	} else {
		goto L180
	}
L179:
	;
	v960 = v499 + int32(8)
	goto L52
L180:
	;
	v960 = v200 + int32(4)
	goto L52
L181:
	;
	v510 = v208
	goto L183
L182:
	;
	v510 = v215
	goto L183
L183:
	;
	if v216 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v511 = v215
	goto L186
L185:
	;
	v511 = v510
	goto L186
L186:
	;
	if v207 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v512 = v211
	goto L189
L188:
	;
	v512 = v208
	goto L189
L189:
	;
	if v216 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v513 = v211
	goto L192
L191:
	;
	v513 = v512
	goto L192
L192:
	;
	if v205&int32(1) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v520 = v33 + int32(48) + v218<<(uint(int32(3))%32)
	if v219 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L195
L195:
	;
	if v219 != 0 {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v520)))
	F_fmtint(m, v521, v231, v214, v209, v513, v213, v511, v207, l0)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L33
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v524 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v520))))
	F_fmtint(m, v524, v231, v214, v209, v513, v213, v511, v207, l0)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L33
	} else {
		goto L200
	}
L199:
	;
	v960 = v200
	goto L52
L200:
	;
	v960 = v200
	goto L52
L201:
	;
	v530 = (v200 + int32(7)) & int32(-8)
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v530)))
	F_fmtint(m, v533, v231, v214, v209, v513, v213, v511, v207, l0)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L33
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v536 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v200))))
	F_fmtint(m, v536, v231, v214, v209, v513, v213, v511, v207, l0)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L33
	} else {
		goto L205
	}
L204:
	;
	v960 = v530 + int32(8)
	goto L52
L205:
	;
	v960 = v200 + int32(4)
	goto L52
L206:
	;
	v548 = v33 + int32(48) + v218<<(uint(int32(3))%32)
	goto L208
L207:
	;
	v548 = v200
	goto L208
L208:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	if v207 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v551 = v211
	goto L211
L210:
	;
	v551 = v208
	goto L211
L211:
	;
	if v216 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v552 = v211
	goto L214
L213:
	;
	v552 = v551
	goto L214
L214:
	;
	v554 = v552 - int32(1)
	v555 = int32(0)
	if v555 < v554 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v558 = v554
	goto L217
L216:
	;
	v558 = v555
	goto L217
L217:
	;
	if v209 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v560 = int32(0) - v558
	goto L220
L219:
	;
	v560 = v558
	goto L220
L220:
	;
	if int32(0) < v560 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	F_dopr_outchmulti(m, int32(32), v560, l0)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L33
	} else {
		goto L224
	}
L222:
	;
	v567 = v560
	goto L223
L223:
	;
	if v547 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v567 = int32(0)
	goto L223
L225:
	;
	v570 = int32(0)
	goto L227
L226:
	;
	v570 = int32(4)
	goto L227
L227:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v572 == int32(0) {
		v600 = v571
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v610 = v200 + v570
	if int32(0) <= v567 {
		v960 = v610
		goto L52
	} else {
		goto L240
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v600 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v600))) = uint8(v549)
	goto L228
L230:
	;
	if base.Ui32(v571) < base.Ui32(v572) {
		v600 = v571
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v576 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v579 + int32(1)
	goto L228
L233:
	;
	goto L234
L234:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v583 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v600 = v599
	goto L229
L236:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v571 == v584 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v587 = v571 - v584
	v588 = F_fwrite(m, v584, int32(1), v587, v576)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L33
	} else {
		goto L238
	}
L238:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v588 + v590
	if v587 == v588 {
		goto L235
	} else {
		goto L239
	}
L239:
	;
	v594 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v594)
	goto L235
L240:
	;
	F_dopr_outchmulti(m, int32(32), int32(0)-v567, l0)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L33
	} else {
		goto L241
	}
L241:
	;
	v960 = v610
	goto L52
L242:
	;
	v618 = v211
	goto L244
L243:
	;
	v618 = v208
	goto L244
L244:
	;
	if v216 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v619 = v211
	goto L247
L246:
	;
	v619 = v618
	goto L247
L247:
	;
	v626 = v205 & int32(1)
	if v626 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v627 = v33 + int32(48) + v218<<(uint(int32(3))%32)
	goto L250
L249:
	;
	v627 = v200
	goto L250
L250:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	if v628 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v630 = v628
	goto L253
L252:
	;
	v630 = int32(701724)
	goto L253
L253:
	;
	if v626 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v633 = int32(0)
	goto L256
L255:
	;
	v633 = int32(4)
	goto L256
L256:
	;
	if v207 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v642 = v200 + v633
	v643 = int32(0)
	v644 = v619 - v641
	if v643 < v644 {
		goto L268
	} else {
		goto L269
	}
L258:
	;
	if v216 != 0 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	goto L260
L260:
	;
	v640 = F_strlen(m, v630)
	mBase = m.M
	v641 = v640
	goto L257
L261:
	;
	v634 = v215
	goto L263
L262:
	;
	v634 = v208
	goto L263
L263:
	;
	v637 = F_memchr(m, v630, int32(0), v634)
	mBase = m.M
	if v637 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v641 = v639
	goto L257
L265:
	;
	v639 = v637 - v630
	goto L267
L266:
	;
	v639 = v634
	goto L267
L267:
	;
	goto L264
L268:
	;
	v648 = v644
	goto L270
L269:
	;
	v648 = v643
	goto L270
L270:
	;
	if v209 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v650 = v643 - v648
	goto L273
L272:
	;
	v650 = v648
	goto L273
L273:
	;
	if int32(0) < v650 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	F_dopr_outchmulti(m, int32(32), v650, l0)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L33
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	F_dostr(m, v630, v641, l0)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L33
	} else {
		goto L279
	}
L277:
	;
	F_dostr(m, v630, v641, l0)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L33
	} else {
		goto L278
	}
L278:
	;
	v960 = v642
	goto L52
L279:
	;
	if int32(0) <= v650 {
		v960 = v642
		goto L52
	} else {
		goto L280
	}
L280:
	;
	F_dopr_outchmulti(m, int32(32), int32(0)-v650, l0)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L33
	} else {
		goto L281
	}
L281:
	;
	v960 = v642
	goto L52
L282:
	;
	v674 = v33 + int32(48) + v218<<(uint(int32(3))%32)
	goto L284
L283:
	;
	v674 = v200
	goto L284
L284:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v675
	if v673 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v679 = int32(0)
	goto L287
L286:
	;
	v679 = int32(4)
	goto L287
L287:
	;
	v680 = v200 + v679
	v685 = F_snprintf(m, v33+int32(320), int32(64), int32(249397), v33)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L33
	} else {
		goto L288
	}
L288:
	;
	if v685 < int32(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v689 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v689)
	v960 = v680
	goto L52
L290:
	;
	goto L291
L291:
	;
	F_dostr(m, v33+int32(320), v685, l0)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L33
	} else {
		goto L292
	}
L292:
	;
	v960 = v680
	goto L52
L293:
	;
	if v707 != 0 {
		goto L376
	} else {
		goto L377
	}
L294:
	;
	v907 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v907)
	goto L293
L295:
	;
	if v207 != 0 {
		goto L331
	} else {
		goto L332
	}
L296:
	;
	v708 = v33 + int32(48) + v218<<(uint(int32(3))%32)
	goto L298
L297:
	;
	v708 = v698
	goto L298
L298:
	;
	v709 = *(*float64)(unsafe.Add(mBase, uint32(v708)))
	v710 = base.I64_reinterpret_f64(v709)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v710&int64(9223372036854775807)) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+320)) = int32(5136718)
	v718 = int32(0)
	v795 = v718
	v796 = int32(3)
	v797 = v718
	goto L295
L300:
	;
	goto L301
L301:
	;
	if base.F64_lt(v709, float64(0)) != 0 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	if base.F64_eq(base.F64_abs(v732), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L309
	} else {
		goto L310
	}
L303:
	;
	v732 = base.F64_neg(v709)
	v733 = int32(45)
	goto L302
L304:
	;
	if base.B2i32(v710 != int64(0))&base.F64_eq(v709, float64(0)) != 0 {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	if v214 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v729 = int32(43)
	goto L308
L307:
	;
	v729 = int32(0)
	goto L308
L308:
	;
	v732 = v709
	v733 = v729
	goto L302
L309:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1330])))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+328)) = uint8(v739)
	v742 = *(*int64)(unsafe.Add(mBase, _consts[1331]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+320)) = v742
	v795 = int32(0)
	v796 = int32(8)
	v797 = v733
	goto L295
L310:
	;
	goto L311
L311:
	;
	if v207 != 0 {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	if v792 < int32(0) {
		goto L294
	} else {
		goto L330
	}
L313:
	;
	v745 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+1356)) = uint8(v745)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+1355)) = uint8(v202)
	v748 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+1354)) = uint8(v748)
	v750 = int32(11813)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+1352)) = uint16(v750)
	*(*float64)(unsafe.Add(mBase, uint32(v33)+40)) = v732
	if v207 != 0 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	goto L315
L315:
	;
	v774 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+1354)) = uint8(v774)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+1353)) = uint8(v202)
	v778 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+1352)) = uint8(v778)
	*(*float64)(unsafe.Add(mBase, uint32(v33)+16)) = v732
	v788 = F_snprintf(m, v33+int32(320), int32(1024), v33+int32(1352), v33+int32(16))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L33
	} else {
		goto L329
	}
L316:
	;
	v754 = v208
	goto L318
L317:
	;
	v754 = v215
	goto L318
L318:
	;
	if v216 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v755 = v215
	goto L321
L320:
	;
	v755 = v754
	goto L321
L321:
	;
	v756 = int32(0)
	if v756 < v755 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v759 = v755
	goto L324
L323:
	;
	v759 = v756
	goto L324
L324:
	;
	if int32(350) <= v759 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v762 = int32(350)
	goto L327
L326:
	;
	v762 = v759
	goto L327
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v762
	v772 = F_snprintf(m, v33+int32(320), int32(1024), v33+int32(1352), v33+int32(32))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L33
	} else {
		goto L328
	}
L328:
	;
	v790 = v759 - v762
	v792 = v772
	goto L312
L329:
	;
	v790 = v774
	v792 = v788
	goto L312
L330:
	;
	v795 = v790
	v796 = v792
	v797 = v733
	goto L295
L331:
	;
	v800 = v211
	goto L333
L332:
	;
	v800 = v208
	goto L333
L333:
	;
	if v216 != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v801 = v211
	goto L336
L335:
	;
	v801 = v800
	goto L336
L336:
	;
	v803 = v801 - (v795 + v796)
	v804 = int32(0)
	if v804 < v803 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v807 = v803
	goto L339
L338:
	;
	v807 = v804
	goto L339
L339:
	;
	if v209 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v809 = int32(0) - v807
	goto L342
L341:
	;
	v809 = v807
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+316)) = v809
	F_leading_pad(m, v213, v797, v33+int32(316), l0)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L33
	} else {
		goto L343
	}
L343:
	;
	if int32(0) < v795 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v33)+316))
	if int32(0) <= v898 {
		goto L293
	} else {
		goto L371
	}
L345:
	;
	v818 = v33 + int32(320)
	v822 = F_strlen(m, v818)
	mBase = m.M
	v829 = v822 + int32(1)
	goto L351
L346:
	;
	goto L347
L347:
	;
	F_dostr(m, v33+int32(320), v796, l0)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L33
	} else {
		goto L370
	}
L348:
	;
	F_dostr(m, v33+int32(320), v796, l0)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L33
	} else {
		goto L368
	}
L349:
	;
	if v841 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L350:
	;
	goto L349
L351:
	;
	v831 = int32(0)
	if v829 == v831 {
		v841 = v831
		goto L350
	} else {
		goto L353
	}
L352:
	;
	v841 = v836
	goto L350
L353:
	;
	v835 = v829 - int32(1)
	v836 = v818 + v835
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836))))
	if v837 != int32(101) {
		v829 = v835
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v845 = v33 + int32(320)
	v849 = F_strlen(m, v845)
	mBase = m.M
	v856 = v849 + int32(1)
	goto L360
L356:
	;
	v871 = v841
	goto L357
L357:
	;
	v873 = v33 + int32(320)
	v876 = v871 - v873
	F_dostr(m, v873, v876, l0)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L33
	} else {
		goto L365
	}
L358:
	;
	if v868 == int32(0) {
		goto L348
	} else {
		goto L364
	}
L359:
	;
	goto L358
L360:
	;
	v858 = int32(0)
	if v856 == v858 {
		v868 = v858
		goto L359
	} else {
		goto L362
	}
L361:
	;
	v868 = v863
	goto L359
L362:
	;
	v862 = v856 - int32(1)
	v863 = v845 + v862
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	if v864 != int32(69) {
		v856 = v862
		goto L360
	} else {
		goto L363
	}
L363:
	;
	goto L361
L364:
	;
	v871 = v868
	goto L357
L365:
	;
	F_dopr_outchmulti(m, int32(48), v795, l0)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L33
	} else {
		goto L366
	}
L366:
	;
	F_dostr(m, v871, v796-v876, l0)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L33
	} else {
		goto L367
	}
L367:
	;
	goto L344
L368:
	;
	F_dopr_outchmulti(m, int32(48), v795, l0)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L33
	} else {
		goto L369
	}
L369:
	;
	goto L344
L370:
	;
	goto L344
L371:
	;
	F_dopr_outchmulti(m, int32(32), int32(0)-v898, l0)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L33
	} else {
		goto L372
	}
L372:
	;
	if v707 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v906 = v200
	goto L375
L374:
	;
	v906 = v700
	goto L375
L375:
	;
	v960 = v906
	goto L52
L376:
	;
	v915 = v200
	goto L378
L377:
	;
	v915 = v700
	goto L378
L378:
	;
	v960 = v915
	goto L52
L379:
	;
	v920 = F_strlen(m, v918)
	mBase = m.M
	F_dostr(m, v918, v920, l0)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L33
	} else {
		goto L380
	}
L380:
	;
	v960 = v200
	goto L52
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v952 + int32(1)
	v958 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v952))) = uint8(v958)
	v960 = v200
	goto L52
L382:
	;
	if base.Ui32(v923) < base.Ui32(v924) {
		v952 = v923
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v928 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v931 + int32(1)
	v960 = v200
	goto L52
L385:
	;
	goto L386
L386:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v935 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v952 = v951
	goto L381
L388:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v923 == v936 {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v939 = v923 - v936
	v940 = F_fwrite(m, v936, int32(1), v939, v928)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L33
	} else {
		goto L390
	}
L390:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v940 + v942
	if v939 == v940 {
		goto L387
	} else {
		goto L391
	}
L391:
	;
	v946 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v946)
	goto L387
L392:
	;
	goto L3
L393:
	;
	v980 = int32(1)
	v981 = v200
	goto L394
L394:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(320)+v980<<(uint(int32(2))%32))))
	switch v1012 {
	case 0:
		goto L49
	case 1:
		goto L402
	case 2:
		goto L401
	case 3:
		goto L400
	case 4:
		goto L399
	case 5:
		goto L398
	default:
		v1062 = v981
		goto L396
	}
L395:
	;
	goto L50
L396:
	;
	v1064 = v980 + int32(1)
	if v1064 <= v294 {
		v980 = v1064
		v981 = v1062
		goto L394
	} else {
		goto L403
	}
L397:
	;
	v1062 = v981 + int32(4)
	goto L396
L398:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(48)+v980<<(uint(int32(3))%32)))) = v1058
	goto L397
L399:
	;
	v1048 = (v981 + int32(7)) & int32(-8)
	v1049 = *(*float64)(unsafe.Add(mBase, uint32(v1048)))
	*(*float64)(unsafe.Add(mBase, uint32(v33+int32(48)+v980<<(uint(int32(3))%32)))) = v1049
	v1062 = v1048 + int32(8)
	goto L396
L400:
	;
	v1035 = (v981 + int32(7)) & int32(-8)
	v1036 = *(*int64)(unsafe.Add(mBase, uint32(v1035)))
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(48)+v980<<(uint(int32(3))%32)))) = v1036
	v1062 = v1035 + int32(8)
	goto L396
L401:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(48)+v980<<(uint(int32(3))%32)))) = v1025
	goto L397
L402:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(48)+v980<<(uint(int32(3))%32)))) = v1018
	goto L397
L403:
	;
	goto L395
L404:
	;
	v1198 = v200
	v1199 = v209
	v1202 = int32(0)
	v1203 = v1096
	v1205 = v207
	v1209 = v211
	v1211 = v213
	v1212 = v214
	v1213 = v215
	v1214 = v216
	v1216 = v208
	v1217 = v219
	v1220 = v1097
	v1221 = v223
	goto L45
L405:
	;
	goto L406
L406:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(48)+v208<<(uint(int32(3))%32))))
	if v207 != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1107 = int32(0)
	v1109 = base.B2i32(v1107 <= v1106)
	if v1107 <= v1106 {
		goto L410
	} else {
		goto L411
	}
L408:
	;
	goto L409
L409:
	;
	v1113 = v1106 >> (uint(int32(31)) % 32)
	if v1106 < int32(0) {
		goto L413
	} else {
		goto L414
	}
L410:
	;
	v1110 = v1106
	goto L412
L411:
	;
	v1110 = v1107
	goto L412
L412:
	;
	v1172 = v1096
	v1174 = v1109
	v1178 = v211
	v1180 = v213
	v1181 = v214
	v1182 = v1110
	v1183 = v216
	v1186 = v219
	v1189 = v1097
	v1190 = v223
	v1195 = int32(0)
	goto L46
L413:
	;
	v1119 = int32(1)
	goto L415
L414:
	;
	v1119 = v209
	goto L415
L415:
	;
	v1120 = int32(0)
	v1198 = v200
	v1199 = v1119
	v1202 = v1120
	v1203 = v1096
	v1205 = v1120
	v1209 = v1106 ^ v1113 - v1113
	v1211 = v213
	v1212 = v214
	v1213 = v215
	v1214 = v216
	v1216 = v218
	v1217 = v219
	v1220 = v1097
	v1221 = v223
	goto L45
L416:
	;
	v1229 = v1227
	goto L418
L417:
	;
	v1229 = int32(701724)
	goto L418
L418:
	;
	v1230 = F_strlen(m, v1229)
	mBase = m.M
	F_dostr(m, v1229, v1230, l0)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L33
	} else {
		goto L419
	}
L419:
	;
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1237 == int32(0) {
		v38 = v178 + int32(2)
		v39 = v39 + int32(4)
		v60 = v179
		goto L1
	} else {
		goto L420
	}
L420:
	;
	goto L3
}
func F_dutch_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
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
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v1109
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= v9 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	if v72 == v7 {
		v102 = v7
		goto L31
	} else {
		goto L32
	}
L4:
	;
	goto L3
L5:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v81
	goto L2
L6:
	;
	if v72 <= v71 {
		goto L4
	} else {
		goto L30
	}
L7:
	;
	v32 = F_find_among(m, l0, int32(4231184), int32(11))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v71 = v9
	v72 = v15
	goto L6
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v9))))
	v20 = int32(224)
	if v19&v20 != v20 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if int32(1)<<(uint(v19)%32)&int32(340306450) != 0 {
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v36
	switch v32 - int32(1) {
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = v36
	v72 = v70
	goto L6
L15:
	;
	v66 = F_slice_from_s(m, l0, int32(1), int32(2209251))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L28
	}
L16:
	;
	v60 = F_slice_from_s(m, l0, int32(1), int32(2209250))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v54 = F_slice_from_s(m, l0, int32(1), int32(2209249))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v48 = F_slice_from_s(m, l0, int32(1), int32(2209248))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v42 = F_slice_from_s(m, l0, int32(1), int32(2209247))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v42 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v1109 = v42
	goto L1
L22:
	;
	if int32(0) <= v48 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v1109 = v48
	goto L1
L24:
	;
	if int32(0) <= v54 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v1109 = v54
	goto L1
L26:
	;
	if int32(0) <= v60 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v1109 = v60
	goto L1
L28:
	;
	if int32(0) <= v66 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v1109 = v66
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71 + int32(1)
	goto L5
L31:
	;
	v106 = v102
	goto L36
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v7))))
	if v87 != int32(121) {
		v102 = v7
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v90 = int32(1)
	v91 = v7 + v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
	v96 = F_slice_from_s(m, l0, v90, int32(2209252))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v96 < int32(0) {
		v1109 = v96
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = v100
	goto L31
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v118 < v117 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v259
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v270 = v268 + int32(3)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v271 < v270 {
		goto L85
	} else {
		goto L86
	}
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v161 != 0 {
		v259 = v162
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v120 = v117
	goto L41
L40:
	;
	v120 = v118
	goto L41
L41:
	;
	goto L43
L42:
	;
	v161 = v157
	goto L38
L43:
	;
	if v117 == v120 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v157 = int32(0)
	goto L42
L45:
	;
	v161 = int32(-1)
	goto L38
L46:
	;
	goto L47
L47:
	;
	v132 = int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v117))))
	if int32(232) < v135 {
		v157 = v132
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v137 = v135 - int32(97)
	if v137 < int32(0) {
		v157 = v132
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v137)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v143)>>(uint(v137&int32(7))%32))&int32(1) == int32(0) {
		v157 = v132
		goto L42
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117 + int32(1)
	goto L51
L51:
	;
	goto L44
L52:
	;
	if v106 < v259 {
		goto L82
	} else {
		goto L83
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v163
	if v163 == v162 {
		v231 = v162
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	goto L36
L55:
	;
	v251 = F_slice_from_s(m, l0, int32(1), int32(2209281))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L80
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	if v163 == v231 {
		goto L74
	} else {
		goto L75
	}
L57:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v163))))
	if v168 != int32(105) {
		v231 = v162
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v172 = v163 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v184 < v172 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v227 == int32(0) {
		goto L55
	} else {
		goto L73
	}
L60:
	;
	v186 = v172
	goto L62
L61:
	;
	v186 = v184
	goto L62
L62:
	;
	goto L64
L63:
	;
	v227 = v223
	goto L59
L64:
	;
	if v172 == v186 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v223 = int32(0)
	goto L63
L66:
	;
	v227 = int32(-1)
	goto L59
L67:
	;
	goto L68
L68:
	;
	v198 = int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v172))))
	if int32(232) < v201 {
		v223 = v198
		goto L63
	} else {
		goto L69
	}
L69:
	;
	v203 = v201 - int32(97)
	if v203 < int32(0) {
		v223 = v198
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v203)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v209)>>(uint(v203&int32(7))%32))&int32(1) == int32(0) {
		v223 = v198
		goto L63
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163 + int32(2)
	goto L72
L72:
	;
	goto L65
L73:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v231 = v230
	goto L56
L74:
	;
	v259 = v163
	goto L52
L75:
	;
	goto L76
L76:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v163))))
	if v236 != int32(121) {
		v259 = v231
		goto L52
	} else {
		goto L77
	}
L77:
	;
	v239 = int32(1)
	v240 = v163 + v239
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v240
	v245 = F_slice_from_s(m, l0, v239, int32(2209282))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if int32(0) <= v245 {
		goto L54
	} else {
		goto L79
	}
L79:
	;
	v1109 = v245
	goto L1
L80:
	;
	if v251 < int32(0) {
		v1109 = v251
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L54
L82:
	;
	v262 = v106 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v262
	v106 = v262
	goto L36
L83:
	;
	goto L84
L84:
	;
	goto L37
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v508
	if v508 <= v7 {
		v614 = v270
		goto L154
	} else {
		goto L155
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v268
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v283 < v268 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v323 < int32(0) {
		goto L85
	} else {
		goto L102
	}
L88:
	;
	v285 = v268
	goto L90
L89:
	;
	v285 = v283
	goto L90
L90:
	;
	v292 = v268
	goto L92
L91:
	;
	v323 = v303
	goto L87
L92:
	;
	if v292 == v285 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v323 = int32(-1)
	goto L87
L95:
	;
	goto L96
L96:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v292))))
	if int32(232) < v298 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v315 = v292 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v315
	v292 = v315
	goto L92
L98:
	;
	v300 = v298 - int32(97)
	if v300 < int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v303 = int32(1)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v300)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v307)>>(uint(v300&int32(7))%32))&v303 != 0 {
		goto L91
	} else {
		goto L100
	}
L100:
	;
	goto L97
L102:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v327 = v326 + v323
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v327
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v338 < v327 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v381 < int32(0) {
		goto L85
	} else {
		goto L117
	}
L104:
	;
	v340 = v327
	goto L106
L105:
	;
	v340 = v338
	goto L106
L106:
	;
	v347 = v327
	goto L108
L107:
	;
	v381 = int32(1)
	goto L103
L108:
	;
	if v347 == v340 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v381 = int32(-1)
	goto L103
L111:
	;
	goto L112
L112:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v347))))
	if int32(232) < v355 {
		goto L107
	} else {
		goto L113
	}
L113:
	;
	v357 = v355 - int32(97)
	if v357 < int32(0) {
		goto L107
	} else {
		goto L114
	}
L114:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v357)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v363)>>(uint(v357&int32(7))%32))&int32(1) == int32(0) {
		goto L107
	} else {
		goto L115
	}
L115:
	;
	v372 = v347 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v372
	v347 = v372
	goto L108
L117:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v385 = v384 + v381
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	if v388 < v385 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v390 = v385
	goto L120
L119:
	;
	v390 = v388
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = v390
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v400 < v399 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v440 < int32(0) {
		goto L85
	} else {
		goto L136
	}
L122:
	;
	v402 = v399
	goto L124
L123:
	;
	v402 = v400
	goto L124
L124:
	;
	v409 = v399
	goto L126
L125:
	;
	v440 = v420
	goto L121
L126:
	;
	if v409 == v402 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v440 = int32(-1)
	goto L121
L129:
	;
	goto L130
L130:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v409))))
	if int32(232) < v415 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v432 = v409 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v432
	v409 = v432
	goto L126
L132:
	;
	v417 = v415 - int32(97)
	if v417 < int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v420 = int32(1)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v417)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v424)>>(uint(v417&int32(7))%32))&v420 != 0 {
		goto L125
	} else {
		goto L134
	}
L134:
	;
	goto L131
L136:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v444 = v443 + v440
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v444
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v455 < v444 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v498 < int32(0) {
		goto L85
	} else {
		goto L151
	}
L138:
	;
	v457 = v444
	goto L140
L139:
	;
	v457 = v455
	goto L140
L140:
	;
	v464 = v444
	goto L142
L141:
	;
	v498 = int32(1)
	goto L137
L142:
	;
	if v464 == v457 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v498 = int32(-1)
	goto L137
L145:
	;
	goto L146
L146:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470+v464))))
	if int32(232) < v472 {
		goto L141
	} else {
		goto L147
	}
L147:
	;
	v474 = v472 - int32(97)
	if v474 < int32(0) {
		goto L141
	} else {
		goto L148
	}
L148:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v474)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v480)>>(uint(v474&int32(7))%32))&int32(1) == int32(0) {
		goto L141
	} else {
		goto L149
	}
L149:
	;
	v489 = v464 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v489
	v464 = v489
	goto L142
L151:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v502 + v498
	goto L85
L152:
	;
	v626 = F_r_e_ending_1(m, l0)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L12
	} else {
		goto L193
	}
L153:
	;
	if v548 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L154:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v615
	v623 = v615
	v624 = v614
	v625 = v615
	goto L152
L155:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v508-int32(1)))))
	if v516&int32(224) != int32(96) {
		v614 = v270
		goto L154
	} else {
		goto L156
	}
L156:
	;
	if int32(1)<<(uint(v516)%32)&int32(540704) == int32(0) {
		v614 = v270
		goto L154
	} else {
		goto L157
	}
L157:
	;
	v529 = F_find_among_b(m, l0, int32(4231408), int32(5))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L12
	} else {
		goto L158
	}
L158:
	;
	if v529 == int32(0) {
		v614 = v270
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v533
	switch v529 - int32(1) {
	case 0:
		goto L162
	case 1:
		goto L161
	case 2:
		goto L160
	default:
		v614 = v270
		goto L154
	}
L160:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+8))
	if v533 < v555 {
		goto L172
	} else {
		goto L173
	}
L161:
	;
	v548 = F_r_en_ending_1(m, l0)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L12
	} else {
		goto L168
	}
L162:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	if v533 < v538 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v614 = int32(0)
	goto L154
L164:
	;
	goto L165
L165:
	;
	v543 = F_slice_from_s(m, l0, int32(4), int32(2209293))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	if v543 < int32(0) {
		v1109 = v543
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v614 = int32(1)
	goto L154
L168:
	;
	if v548 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v553 = int32(base.Ui32(v548) >> (uint(int32(31)) % 32))
	goto L171
L170:
	;
	v553 = int32(2)
	goto L171
L171:
	;
	switch v553 {
	case 0, 2:
		v614 = v548
		goto L154
	default:
		goto L153
	}
L172:
	;
	v614 = int32(0)
	goto L154
L173:
	;
	goto L174
L174:
	;
	v558 = int32(1)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L177
L175:
	;
	if v607 != 0 {
		v614 = v558
		goto L154
	} else {
		goto L187
	}
L176:
	;
	v607 = v604
	goto L175
L177:
	;
	if v566 <= v567 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v604 = int32(0)
	goto L176
L179:
	;
	v607 = int32(-1)
	goto L175
L180:
	;
	goto L181
L181:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578+v566-int32(1)))))
	if int32(232) < v582 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566 - int32(1)
	goto L186
L183:
	;
	v584 = v582 - int32(97)
	if v584 < int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v587 = int32(1)
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v584)>>(uint(int32(3))%32)))+uint32(_consts[1427]))))
	if int32(base.Ui32(v591)>>(uint(v584&int32(7))%32))&v587 != 0 {
		v604 = v587
		goto L176
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	goto L178
L187:
	;
	v608 = F_slice_del(m, l0)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L12
	} else {
		goto L188
	}
L188:
	;
	if v608 < int32(0) {
		v1109 = v608
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v614 = v558
	goto L154
L190:
	;
	return v548
L191:
	;
	goto L192
L192:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v623 = v620
	v624 = v548
	v625 = v621
	goto L152
L193:
	;
	if v626 < int32(0) {
		v1109 = v626
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v630 = v623 - v625
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v632 = v630 + v631
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v632
	v635 = int32(4)
	v637 = int32(0)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v632-v640 < v635 {
		v650 = v637
		goto L202
	} else {
		goto L203
	}
L195:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1059
	v1062 = v1059
	goto L332
L196:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L300
L197:
	;
	if v909 < int32(0) {
		v1109 = v909
		goto L1
	} else {
		goto L297
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v712
	v717 = v712 - int32(1)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v717 <= v718 {
		goto L230
	} else {
		goto L231
	}
L199:
	;
	if v693 < int32(0) {
		v909 = v698
		goto L197
	} else {
		goto L228
	}
L200:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v707 = v706 + v630
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v707
	v712 = v707
	v713 = v704
	v714 = v706
	goto L198
L201:
	;
	if v650 == int32(0) {
		v704 = v624
		goto L200
	} else {
		goto L205
	}
L202:
	;
	goto L201
L203:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v646 = F_memcmp(m, v643+v632-v635, int32(2209329), v635)
	mBase = m.M
	if v646 != 0 {
		v650 = v637
		goto L202
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v632 - v635
	v650 = int32(1)
	goto L202
L205:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v653
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	if v653 < v656 {
		v704 = v624
		goto L200
	} else {
		goto L206
	}
L206:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v658 < v653 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660+v653-int32(1)))))
	if v664 == int32(99) {
		v704 = v624
		goto L200
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v667 = F_slice_del(m, l0)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L12
	} else {
		goto L211
	}
L210:
	;
	goto L209
L211:
	;
	if v667 < int32(0) {
		v1109 = v667
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v671
	v673 = int32(2)
	v675 = int32(0)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v671-v678 < v673 {
		v688 = v675
		goto L214
	} else {
		goto L215
	}
L213:
	;
	if v688 == int32(0) {
		v704 = v624
		goto L200
	} else {
		goto L217
	}
L214:
	;
	goto L213
L215:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v684 = F_memcmp(m, v681+v671-v673, int32(2209333), v673)
	mBase = m.M
	if v684 != 0 {
		v688 = v675
		goto L214
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v671 - v673
	v688 = int32(1)
	goto L214
L217:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v691
	v693 = F_r_en_ending_1(m, l0)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L12
	} else {
		goto L218
	}
L218:
	;
	v696 = base.B2i32(v693 < int32(0))
	if v693 < int32(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v697 = v693
	goto L221
L220:
	;
	v697 = v624
	goto L221
L221:
	;
	if v693 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v698 = v697
	goto L224
L223:
	;
	v698 = v624
	goto L224
L224:
	;
	if v693 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v702 = int32(base.Ui32(v693) >> (uint(int32(31)) % 32))
	goto L227
L226:
	;
	v702 = int32(4)
	goto L227
L227:
	;
	switch v702 {
	case 0, 4:
		v704 = v698
		goto L200
	default:
		goto L199
	}
L228:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v712 = v709
	v713 = v698
	v714 = v710
	goto L198
L229:
	;
	v898 = int32(0)
	v899 = base.B2i32(v868 < v898)
	if v899 == v898 {
		goto L196
	} else {
		goto L290
	}
L230:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v894 + (v712 - v714)
	goto L196
L231:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720+v717))))
	if v722&int32(224) != int32(96) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	if int32(1)<<(uint(v722)%32)&int32(264336) == int32(0) {
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v735 = F_find_among_b(m, l0, int32(4231520), int32(6))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L12
	} else {
		goto L234
	}
L234:
	;
	if v735 == int32(0) {
		goto L230
	} else {
		goto L235
	}
L235:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v739
	switch v735 - int32(1) {
	case 0:
		goto L240
	case 1:
		goto L239
	case 2:
		goto L238
	case 3:
		goto L237
	case 4:
		goto L236
	default:
		goto L230
	}
L236:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)+4))
	if v739 < v882 {
		goto L230
	} else {
		goto L286
	}
L237:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	if v739 < v875 {
		goto L230
	} else {
		goto L283
	}
L238:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	if v739 < v862 {
		goto L230
	} else {
		goto L276
	}
L239:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	if v739 < v846 {
		goto L230
	} else {
		goto L269
	}
L240:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	if v739 < v744 {
		goto L230
	} else {
		goto L241
	}
L241:
	;
	v746 = F_slice_del(m, l0)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L12
	} else {
		goto L242
	}
L242:
	;
	if v746 < int32(0) {
		v1109 = v746
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v750
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v753 = int32(2)
	v755 = int32(0)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v750-v758 < v753 {
		v768 = v755
		goto L246
	} else {
		goto L247
	}
L244:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v792 = v790 + (v750 - v752)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v792
	v794 = int32(0)
	v797 = v792 - int32(1)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v797 <= v798 {
		v840 = v794
		goto L257
	} else {
		goto L258
	}
L245:
	;
	if v768 == int32(0) {
		goto L244
	} else {
		goto L249
	}
L246:
	;
	goto L245
L247:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v764 = F_memcmp(m, v761+v750-v753, int32(2209335), v753)
	mBase = m.M
	if v764 != 0 {
		v768 = v755
		goto L246
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v750 - v753
	v768 = int32(1)
	goto L246
L249:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v771
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)+4))
	if v771 < v774 {
		goto L244
	} else {
		goto L250
	}
L250:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v776 < v771 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778+v771-int32(1)))))
	if v782 == int32(101) {
		goto L244
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v785 = F_slice_del(m, l0)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L12
	} else {
		goto L255
	}
L254:
	;
	goto L253
L255:
	;
	if int32(0) <= v785 {
		goto L230
	} else {
		goto L256
	}
L256:
	;
	v1109 = v785
	goto L1
L257:
	;
	if int32(0) <= v840 {
		goto L230
	} else {
		goto L268
	}
L258:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800+v797))))
	if v802&int32(224) != int32(96) {
		v840 = v794
		goto L257
	} else {
		goto L259
	}
L259:
	;
	if int32(1)<<(uint(v802)%32)&int32(1050640) == int32(0) {
		v840 = v794
		goto L257
	} else {
		goto L260
	}
L260:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v816 = F_find_among_b(m, l0, int32(4231728), int32(3))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L12
	} else {
		goto L261
	}
L261:
	;
	if v816 == int32(0) {
		v840 = v794
		goto L257
	} else {
		goto L262
	}
L262:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v822 = v820 + (v792 - v813)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v822
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v822
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v822 <= v825 {
		v840 = v794
		goto L257
	} else {
		goto L263
	}
L263:
	;
	v827 = int32(1)
	v828 = v822 - v827
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v828
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v828
	v832 = F_slice_del(m, l0)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L12
	} else {
		goto L264
	}
L264:
	;
	if int32(0) <= v832 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v839 = v827
	goto L267
L266:
	;
	v839 = v832 >> (uint(int32(31)) % 32) & v832
	goto L267
L267:
	;
	v840 = v839
	goto L257
L268:
	;
	v1109 = v840
	goto L1
L269:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v848 < v739 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850+v739-int32(1)))))
	if v854 == int32(101) {
		goto L230
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v857 = F_slice_del(m, l0)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L12
	} else {
		goto L274
	}
L273:
	;
	goto L272
L274:
	;
	if int32(0) <= v857 {
		goto L230
	} else {
		goto L275
	}
L275:
	;
	v1109 = v857
	goto L1
L276:
	;
	v864 = F_slice_del(m, l0)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L12
	} else {
		goto L277
	}
L277:
	;
	if v864 < int32(0) {
		v1109 = v864
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v868 = F_r_e_ending_1(m, l0)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L12
	} else {
		goto L279
	}
L279:
	;
	if v868 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v873 = int32(base.Ui32(v868) >> (uint(int32(31)) % 32))
	goto L282
L281:
	;
	v873 = int32(6)
	goto L282
L282:
	;
	switch v873 {
	case 0, 6:
		goto L230
	default:
		goto L229
	}
L283:
	;
	v877 = F_slice_del(m, l0)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L12
	} else {
		goto L284
	}
L284:
	;
	if int32(0) <= v877 {
		goto L230
	} else {
		goto L285
	}
L285:
	;
	v1109 = v877
	goto L1
L286:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v881)+12))
	if v884 == int32(0) {
		goto L230
	} else {
		goto L287
	}
L287:
	;
	v887 = F_slice_del(m, l0)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L12
	} else {
		goto L288
	}
L288:
	;
	if v887 < int32(0) {
		v1109 = v887
		goto L1
	} else {
		goto L289
	}
L289:
	;
	goto L230
L290:
	;
	if v868 < v898 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v902 = v868
	goto L293
L292:
	;
	v902 = v713
	goto L293
L293:
	;
	if v868 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v903 = v902
	goto L296
L295:
	;
	v903 = v713
	goto L296
L296:
	;
	v909 = v903
	goto L197
L297:
	;
	goto L195
L298:
	;
	if v964 != 0 {
		goto L195
	} else {
		goto L310
	}
L299:
	;
	v964 = v961
	goto L298
L300:
	;
	if v923 <= v924 {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v961 = int32(0)
	goto L299
L302:
	;
	v964 = int32(-1)
	goto L298
L303:
	;
	goto L304
L304:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935+v923-int32(1)))))
	if int32(232) < v939 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v923 - int32(1)
	goto L309
L306:
	;
	v941 = v939 - int32(73)
	if v941 < int32(0) {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v944 = int32(1)
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v941)>>(uint(int32(3))%32)))+uint32(_consts[1428]))))
	if int32(base.Ui32(v948)>>(uint(v941&int32(7))%32))&v944 != 0 {
		v961 = v944
		goto L299
	} else {
		goto L308
	}
L308:
	;
	goto L305
L309:
	;
	goto L301
L310:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v967 = v965 - int32(1)
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v967 <= v968 {
		goto L195
	} else {
		goto L311
	}
L311:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970+v967))))
	if v972&int32(224) != int32(96) {
		goto L195
	} else {
		goto L312
	}
L312:
	;
	if int32(1)<<(uint(v972)%32)&int32(2129954) == int32(0) {
		goto L195
	} else {
		goto L313
	}
L313:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v986 = F_find_among_b(m, l0, int32(4231648), int32(4))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L12
	} else {
		goto L314
	}
L314:
	;
	if v986 == int32(0) {
		goto L195
	} else {
		goto L315
	}
L315:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L318
L316:
	;
	if v1038 != 0 {
		goto L195
	} else {
		goto L328
	}
L317:
	;
	v1038 = v1035
	goto L316
L318:
	;
	if v997 <= v998 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1035 = int32(0)
	goto L317
L320:
	;
	v1038 = int32(-1)
	goto L316
L321:
	;
	goto L322
L322:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009+v997-int32(1)))))
	if int32(232) < v1013 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v997 - int32(1)
	goto L327
L324:
	;
	v1015 = v1013 - int32(97)
	if v1015 < int32(0) {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v1018 = int32(1)
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1015)>>(uint(int32(3))%32)))+uint32(_consts[1426]))))
	if int32(base.Ui32(v1022)>>(uint(v1015&int32(7))%32))&v1018 != 0 {
		v1035 = v1018
		goto L317
	} else {
		goto L326
	}
L326:
	;
	goto L323
L327:
	;
	goto L319
L328:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1041 = v1039 + (v965 - v983)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1041
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1041
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1041 <= v1044 {
		goto L195
	} else {
		goto L329
	}
L329:
	;
	v1047 = v1041 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1047
	v1050 = F_slice_del(m, l0)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L12
	} else {
		goto L330
	}
L330:
	;
	if v1050 < int32(0) {
		v1109 = v1050
		goto L1
	} else {
		goto L331
	}
L331:
	;
	goto L195
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1062
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1068 <= v1062 {
		goto L337
	} else {
		goto L338
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1059
	v1109 = int32(1)
	goto L1
L334:
	;
	goto L333
L335:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1062 = v1105
	goto L332
L336:
	;
	if v1098 <= v1097 {
		goto L334
	} else {
		goto L348
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1062
	v1097 = v1062
	v1098 = v1068
	goto L336
L338:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070+v1062))))
	switch v1072 - int32(73) {
	case 0, 16:
		goto L339
	default:
		goto L337
	}
L339:
	;
	v1077 = F_find_among(m, l0, int32(4231792), int32(3))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L12
	} else {
		goto L340
	}
L340:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1079
	switch v1077 - int32(1) {
	case 0:
		goto L342
	case 1:
		goto L341
	case 2:
		goto L343
	default:
		goto L335
	}
L341:
	;
	v1092 = F_slice_from_s(m, l0, int32(1), int32(2209414))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L12
	} else {
		goto L346
	}
L342:
	;
	v1086 = F_slice_from_s(m, l0, int32(1), int32(2209413))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L12
	} else {
		goto L344
	}
L343:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1097 = v1079
	v1098 = v1083
	goto L336
L344:
	;
	if int32(0) <= v1086 {
		goto L335
	} else {
		goto L345
	}
L345:
	;
	v1109 = v1086
	goto L1
L346:
	;
	if int32(0) <= v1092 {
		goto L335
	} else {
		goto L347
	}
L347:
	;
	v1109 = v1092
	goto L1
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1097 + int32(1)
	goto L335
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
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
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
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
	if v26 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v230 == int32(0) {
		goto L3
	} else {
		goto L76
	}
L6:
	;
	v230 = v2
	goto L5
L7:
	;
	goto L8
L8:
	;
	v32 = v2
	v35 = v2
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v32<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = int32(351314)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1462])))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 == int32(0) {
		v67 = v47
		v68 = v48
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L72
	}
L11:
	;
	goto L10
L12:
	;
	v198 = v32 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v198 < v199 {
		v32 = v198
		v35 = v196
		goto L9
	} else {
		goto L71
	}
L13:
	;
	if v68-v67 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	goto L13
L15:
	;
	if v47 != v48 {
		v67 = v47
		v68 = v48
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v52 = v43
	v53 = v44
	goto L17
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v56
		v68 = v57
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v67 = v56
	v68 = v57
	goto L14
L19:
	;
	v60 = int32(1)
	if v56 == v57 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v72 = F_defGetBoolean(m, v42)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v75 = int32(351305)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1463])))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v79 == int32(0) {
		v98 = v78
		v99 = v79
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v72)
	v196 = v35
	goto L12
L25:
	;
	if v99-v98 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	goto L25
L27:
	;
	if v78 != v79 {
		v98 = v78
		v99 = v79
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v83 = v43
	v84 = v75
	goto L29
L29:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v87
		v99 = v88
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v98 = v87
	v99 = v88
	goto L26
L31:
	;
	v91 = int32(1)
	if v87 == v88 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v103 = F_defGetBoolean(m, v42)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v106 = int32(158420)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1464])))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v110 == int32(0) {
		v129 = v109
		v130 = v110
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)) = uint8(v103)
	v196 = v35
	goto L12
L37:
	;
	if v130-v129 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	goto L37
L39:
	;
	if v109 != v110 {
		v129 = v109
		v130 = v110
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v114 = v43
	v115 = v106
	goto L41
L41:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v119 == int32(0) {
		v129 = v118
		v130 = v119
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v129 = v118
	v130 = v119
	goto L38
L43:
	;
	v122 = int32(1)
	if v118 == v119 {
		v114 = v114 + v122
		v115 = v115 + v122
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v134 = F_defGetBoolean(m, v42)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v137 = int32(158407)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1465])))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v141 == int32(0) {
		v160 = v140
		v161 = v141
		goto L50
	} else {
		goto L51
	}
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+10)) = uint8(v134)
	v196 = v35
	goto L12
L49:
	;
	if v161-v160 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	goto L49
L51:
	;
	if v140 != v141 {
		v160 = v140
		v161 = v141
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v145 = v43
	v146 = v137
	goto L53
L53:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v150 == int32(0) {
		v160 = v149
		v161 = v150
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v160 = v149
	v161 = v150
	goto L50
L55:
	;
	v153 = int32(1)
	if v149 == v150 {
		v145 = v145 + v153
		v146 = v146 + v153
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v165 = F_defGetBoolean(m, v42)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v168 = int32(172477)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1466])))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v172 == int32(0) {
		v191 = v171
		v192 = v172
		goto L62
	} else {
		goto L63
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)) = uint8(v165)
	v196 = v35
	goto L12
L61:
	;
	if v192-v191 != 0 {
		goto L11
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	if v171 != v172 {
		v191 = v171
		v192 = v172
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v176 = v43
	v177 = v168
	goto L65
L65:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	if v181 == int32(0) {
		v191 = v180
		v192 = v181
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v191 = v180
	v192 = v181
	goto L62
L67:
	;
	v184 = int32(1)
	if v180 == v181 {
		v176 = v176 + v184
		v177 = v177 + v184
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v194 = F_defGetString(m, v42)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v196 = v194
	goto L12
L71:
	;
	v230 = v196
	goto L5
L72:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v208
	F_errmsg(m, int32(753658), v12+int32(16))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(516174), int32(191), int32(106116))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v238 = F_get_tsearch_config_filename(m, v230, int32(172477))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	F_pfree(m, v238)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L129
	}
L78:
	;
	F_tsearch_readline_end(m, v12+int32(36))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L128
	}
L79:
	;
	v240 = F_tsearch_readline_begin(m, v12+int32(36), v238)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v240 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v244 = F_tsearch_readline(m, v12+int32(36))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L124
	}
L84:
	;
	if v244 == int32(0) {
		goto L78
	} else {
		goto L85
	}
L85:
	;
	v250 = v244
	v252 = int32(0)
	goto L86
L86:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v258 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	F_tsearch_readline_end(m, v12+int32(36))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L121
	}
L88:
	;
	v259 = F_strlen(m, v250)
	mBase = m.M
	v261 = F_str_tolower(m, v250, v259, int32(100))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v332 = v252
	goto L90
L90:
	;
	v340 = F_tsearch_readline(m, v12+int32(36))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L119
	}
L91:
	;
	F_pfree(m, v250)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v266 = v261
	v268 = v252
	goto L93
L93:
	;
	v276 = F_find_word(m, v266, v12+int32(32))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	F_pfree(m, v261)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L118
	}
L95:
	;
	if v276 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v278 == v268 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v323 = v268
	goto L98
L98:
	;
	goto L94
L99:
	;
	if v268 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if v266 != v261 {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	v285 = int32(16)
	goto L104
L103:
	;
	v285 = v268 << (uint(int32(1)) % 32)
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v285
	v288 = v285 << (uint(int32(3)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v289 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v294
	goto L101
L106:
	;
	v290 = F_repalloc(m, v289, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v292 = F_palloc(m, v288)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	v294 = v290
	goto L105
L110:
	;
	v294 = v292
	goto L105
L111:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+10)))
	if v321 != 0 {
		v266 = v317
		v268 = v318
		goto L93
	} else {
		goto L117
	}
L112:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v303 = F_pnstrdup(m, v276, v301-v276)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
	if v299 != 0 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v317 = v300
	v318 = v268
	goto L111
L115:
	;
	v306 = v268 << (uint(int32(3)) % 32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v306+v307))) = v303
	v310 = F_pstrdup(m, v261)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v312+v306)+4)) = v310
	v317 = v301
	v318 = v268 + int32(1)
	goto L111
L117:
	;
	v323 = v318
	goto L98
L118:
	;
	v332 = v323
	goto L90
L119:
	;
	if v340 != 0 {
		v250 = v340
		v252 = v332
		goto L86
	} else {
		goto L120
	}
L120:
	;
	goto L87
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v332
	if v332 < int32(2) {
		goto L77
	} else {
		goto L122
	}
L122:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_pg_qsort(m, v349, v332, int32(8), int32(7542))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L77
L124:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v238
	F_errmsg(m, int32(310507), v12)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(516174), int32(89), int32(17774))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
	goto L77
L129:
	;
	goto L3
}
