package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitParallelPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
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
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v751 int32
	_ = v751
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v850 int32
	_ = v850
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	v6 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = v28
	goto L3
L2:
	;
	v29 = F_MakePerTupleExprContext(m, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_ExecSetParamPlanMulti(m, l2, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v33 = v29
	goto L3
L6:
	;
	v37 = F_palloc0(m, int32(44))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = l0
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+32)) = uint8(v40)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = F_copyObjectImpl(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v110 = F_palloc0(m, int32(104))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
	if v46 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v57 = v40
	goto L12
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v57<<(uint(int32(2))%32))))
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+26)) = uint8(v80)
	v83 = v57 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v83 < v84 {
		v57 = v83
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L8
L14:
	;
	goto L13
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = int64(4294967626)
	v116 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v116 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v110)+8)) = v121
	v124 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v124 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v121 = int64(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v116)+392))
	v121 = v120
	goto L16
L20:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v124)+400))
	v127 = v125
	goto L22
L21:
	;
	v127 = int64(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v44
	v129 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+28)) = uint16(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = int32(65536)
	*(*int64)(unsafe.Add(mBase, uint32(v110)+16)) = v127
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+40)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+44)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+48)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+64)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v110)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+52)) = v141
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	if v148 == v129 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v217 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v110)+68)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v110)+76)) = v217
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+96)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v110)+88)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+84)) = v222
	v228 = F_nodeToString(m, v110)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L36
	}
L24:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v151 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v159 = v129
	v162 = v6
	goto L26
L26:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v159<<(uint(int32(2))%32))))
	if v181 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L23
L28:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+37)))
	if v183 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v186 = int32(0)
	goto L30
L30:
	;
	v187 = F_lappend(m, v162, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L34
	}
L31:
	;
	v184 = v181
	goto L33
L32:
	;
	v184 = int32(0)
	goto L33
L33:
	;
	v186 = v184
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+64)) = v187
	v191 = v159 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v191 < v192 {
		v159 = v191
		v162 = v187
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	v231 = F_CreateParallelContext(m, int32(278341), l3)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v231
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v236 = F_add_size(m, v234, int32(32))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v236
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v241 = F_add_size(m, v239, int32(1))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v246 = F_strlen(m, v245)
	mBase = m.M
	v251 = F_add_size(m, v244, v246&int32(-32)+int32(32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v251
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v256 = F_add_size(m, v254, int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v260 = F_strlen(m, v228)
	mBase = m.M
	v265 = F_add_size(m, v259, v260&int32(-32)+int32(32))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v265
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v270 = F_add_size(m, v268, int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v270
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v274 = m.G0
	v276 = v274 - int32(16)
	m.G0 = v276
	v278 = int32(4)
	if v273 == int32(0) {
		v370 = v278
		goto L44
	} else {
		goto L45
	}
L44:
	;
	m.G0 = v276 + int32(16)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v387 = F_add_size(m, v382, (v370+int32(31))&int32(-32))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L64
	}
L45:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v281 <= int32(0) {
		v370 = v278
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v300 = v278
	v302 = v6
	goto L47
L47:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v309 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v370 = v350
	goto L44
L49:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	v323 = F_add_size(m, v300, int32(4))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L54
	}
L50:
	;
	v315 = m.T0[v309].(func(*base.Module, int32, int32, int32, int32) int32)(m, v273, v302+int32(1), int32(0), v276+int32(4))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v320 = v273 + int32(32) + v302*int32(12)
	goto L49
L53:
	;
	v320 = v315
	goto L49
L54:
	;
	v326 = F_add_size(m, v323, int32(2))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v321 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+4)))
	v348 = F_datumEstimateSpace(m, v344, v345, v342&int32(1), v343)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L61
	}
L57:
	;
	F_get_typlenbyval(m, v321, v276+int32(2), v276+int32(1))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)) = uint8(v336)
	v339 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v276)+2)) = uint16(v339)
	v342 = v336
	v343 = v339
	goto L56
L60:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(v276)+2)))
	v342 = v334
	v343 = v335
	goto L56
L61:
	;
	v350 = F_add_size(m, v326, v348)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v353 = v302 + int32(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v353 < v354 {
		v300 = v350
		v302 = v353
		goto L47
	} else {
		goto L63
	}
L63:
	;
	goto L48
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v387
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v392 = F_add_size(m, v390, int32(1))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v392
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v398 = F_mul_size(m, int32(128), v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v404 = F_add_size(m, v395, (v398+int32(31))&int32(-32))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v404
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v409 = F_add_size(m, v407, int32(1))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v409
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v415 = F_mul_size(m, int32(32), v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v421 = F_add_size(m, v412, (v415+int32(31))&int32(-32))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v421
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v426 = F_add_size(m, v424, int32(1))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v426
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v432 = F_mul_size(m, int32(65536), v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v438 = F_add_size(m, v429, (v432+int32(31))&int32(-32))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v438
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v443 = F_add_size(m, v441, int32(1))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v443
	v446 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v231
	v452 = F_ExecParallelEstimate(m, l0, v26+int32(24))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v454 = int32(1)
	v455 = v246 + v454
	v457 = v260 + v454
	v458 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v459 == v458 {
		v510 = v6
		v511 = v6
		v512 = v458
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v515 = F_add_size(m, v513, int32(4096))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L85
	}
L77:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v464 = F_mul_size(m, v463, l3)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v466 = F_mul_size(m, int32(416), v464)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v474 = (v463<<(uint(int32(2))%32) + int32(23)) & int32(-8)
	v475 = v466 + v474
	v480 = F_add_size(m, v468, (v475+int32(31))&int32(-32))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v480
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v485 = F_add_size(m, v483, int32(1))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v485
	v488 = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v489 == v488 {
		v510 = v475
		v511 = v474
		v512 = v488
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v494 = l3 * int32(48)
	v499 = F_add_size(m, v492, (v494+int32(39))&int32(-32))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v499
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v504 = F_add_size(m, v502, int32(1))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v504
	v510 = v475
	v511 = v474
	v512 = v494 | int32(8)
	goto L76
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v515
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v520 = F_add_size(m, v518, int32(1))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v520
	F_InitializeParallelDSM(m, v231)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v527 = F_shm_toc_allocate(m, v525, int32(24))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v527)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v527))) = l4
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v527)+12)) = v532
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v527)+16)) = v534
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v536, int64(-2305843009213693951), v527)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v541 = F_shm_toc_allocate(m, v540, v455)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v455 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v546, int64(-2305843009213693944), v545)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L95
	}
L92:
	;
	v544 = F__emscripten_memcpy_bulkmem(m, v541, v543, v455)
	mBase = m.M
	v545 = v544
	goto L94
L93:
	;
	v545 = v541
	goto L94
L94:
	;
	goto L91
L95:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v551 = F_shm_toc_allocate(m, v550, v457)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	if v457 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v555, int64(-2305843009213693950), v554)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L101
	}
L98:
	;
	v553 = F__emscripten_memcpy_bulkmem(m, v551, v228, v457)
	mBase = m.M
	v554 = v553
	goto L100
L99:
	;
	v554 = v551
	goto L100
L100:
	;
	goto L97
L101:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v560 = F_shm_toc_allocate(m, v559, v370)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v560
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v563, int64(-2305843009213693949), v560)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v569 = v26 + int32(8)
	v570 = int32(0)
	v571 = m.G0
	v573 = v571 - int32(16)
	m.G0 = v573
	if v567 == v570 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	m.G0 = v573 + int32(16)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v701 = F_mul_size(m, int32(128), v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L126
	}
L105:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = int32(0)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v580 + int32(4)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v567)+28))
	v586 = int32(0)
	if v586 < v585 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v589 = v585
	goto L110
L109:
	;
	v589 = v586
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584))) = v589
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v591 + int32(4)
	if v585 <= int32(0) {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v615 = v570
	goto L112
L112:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if v622 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L104
L114:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v633)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v639 = v637 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v639
	v641 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v639))) = uint16(v641)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v643 + int32(2)
	if v635 != 0 {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	v628 = m.T0[v622].(func(*base.Module, int32, int32, int32, int32) int32)(m, v567, v615+int32(1), int32(0), v573+int32(4))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v633 = v567 + int32(32) + v615*int32(12)
	goto L114
L118:
	;
	v633 = v628
	goto L114
L119:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+4)))
	F_datumSerialize(m, v663, v664, v661&int32(1), v662, v569)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L4
	} else {
		goto L124
	}
L120:
	;
	F_get_typlenbyval(m, v635, v573+int32(2), v573+int32(1))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v655 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)) = uint8(v655)
	v658 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+2)) = uint16(v658)
	v661 = v655
	v662 = v658
	goto L119
L123:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)))
	v654 = int32(*(*int16)(unsafe.Add(mBase, uint32(v573)+2)))
	v661 = v653
	v662 = v654
	goto L119
L124:
	;
	v670 = v615 + int32(1)
	if v585 != v670 {
		v615 = v670
		goto L112
	} else {
		goto L125
	}
L125:
	;
	goto L113
L126:
	;
	v703 = F_shm_toc_allocate(m, v698, v701)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v705, int64(-2305843009213693948), v703)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v703
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v713 = F_mul_size(m, int32(32), v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v715 = F_shm_toc_allocate(m, v710, v713)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v717, int64(-2305843009213693942), v715)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v715
	v723 = F_ExecParallelSetupTupleQueues(m, v231, int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v725 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v723
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v728 == v725 {
		v850 = v446
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v864 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L134:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v732 = F_shm_toc_allocate(m, v731, v510)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v732)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v732)+4)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v732))) = v734
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v732)+12)) = v738
	if int32(0) < l3*v738 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v751 = int32(0)
	goto L139
L137:
	;
	goto L138
L138:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v817, int64(-2305843009213693946), v732)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L4
	} else {
		goto L143
	}
L139:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v769 = int32(416)
	v775 = F__emscripten_memset_bulkmem(m, v732+v511+v751*v769, base.I32_extend8_s(int32(0)), v769)
	mBase = m.M
	goto L141
L140:
	;
	goto L138
L141:
	;
	v776 = int32(1)
	v777 = v768 & v776
	*(*uint8)(unsafe.Add(mBase, uint32(v775))) = uint8(v777)
	v782 = int32(base.Ui32(v768)>>(uint(int32(3))%32)) & v776
	*(*uint8)(unsafe.Add(mBase, uint32(v775)+2)) = uint8(v782)
	v787 = int32(base.Ui32(v768)>>(uint(v776)%32)) & v776
	*(*uint8)(unsafe.Add(mBase, uint32(v775)+1)) = uint8(v787)
	v790 = v751 + v776
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v790 < v791*l3 {
		v751 = v790
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v732
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v822 == int32(0) {
		v850 = v732
		goto L133
	} else {
		goto L144
	}
L144:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v826 = F_shm_toc_allocate(m, v825, v512)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v826))) = l3
	v835 = F__emscripten_memset_bulkmem(m, v826+int32(8), base.I32_extend8_s(int32(0)), l3*int32(48))
	mBase = m.M
	goto L146
L146:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v836, int64(-2305843009213693943), v826)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v826
	v850 = v732
	goto L133
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v850
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v231
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = v892
	v896 = F_ExecParallelInitializeDSM(m, l0, v26+int32(12))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L4
	} else {
		goto L155
	}
L149:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v869 = F_shm_toc_allocate(m, v867, int32(4096))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v871, int64(-2305843009213693945), v869)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	v878 = F_dsa_create_in_place_ext(m, v869, int32(4096), int32(71), v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v878
	if l2 == int32(0) {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v883 = F_SerializeParamExecParams(m, l1, l2, v878)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v883
	*(*int32)(unsafe.Add(mBase, uint32(v527)+8)) = v883
	goto L148
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = int32(0)
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v900 != v901 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L4
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	m.G0 = v26 + int32(32)
	return v37
L159:
	;
	F_errmsg_internal(m, int32(171182), int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(497428), int32(877), int32(283565))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecParallelFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	if int32(0) < v8 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v38 != 0 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v14 = int32(0)
	goto L10
L8:
	;
	v30 = v9
	goto L9
L9:
	;
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L15
	}
L10:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	F_shm_mq_detach(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v30 = v26
	goto L9
L12:
	;
	return
L13:
	;
	v24 = v14 + int32(1)
	if v24 != v8 {
		v14 = v24
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	goto L6
L16:
	;
	if int32(0) < v8 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WaitForParallelWorkersToFinish(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L27
	}
L19:
	;
	v43 = int32(0)
	goto L22
L20:
	;
	v59 = v38
	goto L21
L21:
	;
	F_pfree(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L26
	}
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v43<<(uint(int32(2))%32))))
	F_pfree(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L24
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v59 = v55
	goto L21
L24:
	;
	v53 = v43 + int32(1)
	if v53 != v8 {
		v43 = v53
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	goto L18
L27:
	;
	if int32(0) < v8 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v210)
	goto L3
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = v76 + v74<<(uint(int32(7))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = v80 + v74<<(uint(int32(5))%32)
	v84 = int32(4413656)
	v86 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v86 + v87
	v90 = int32(4413664)
	v92 = *(*int64)(unsafe.Add(mBase, _consts[369]))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, _consts[369])) = v92 + v93
	v96 = int32(4413672)
	v98 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, _consts[370])) = v98 + v99
	v102 = int32(4413680)
	v104 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, _consts[371])) = v104 + v105
	v108 = int32(4413688)
	v110 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v79)+32))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v110 + v111
	v114 = int32(4413696)
	v116 = *(*int64)(unsafe.Add(mBase, _consts[372]))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v79)+40))
	*(*int64)(unsafe.Add(mBase, _consts[372])) = v116 + v117
	v120 = int32(4413704)
	v122 = *(*int64)(unsafe.Add(mBase, _consts[373]))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	*(*int64)(unsafe.Add(mBase, _consts[373])) = v122 + v123
	v126 = int32(4413712)
	v128 = *(*int64)(unsafe.Add(mBase, _consts[374]))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v79)+56))
	*(*int64)(unsafe.Add(mBase, _consts[374])) = v128 + v129
	v132 = int32(4413720)
	v134 = *(*int64)(unsafe.Add(mBase, _consts[375]))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v79)+64))
	*(*int64)(unsafe.Add(mBase, _consts[375])) = v134 + v135
	v138 = int32(4413728)
	v140 = *(*int64)(unsafe.Add(mBase, _consts[376]))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v79)+72))
	*(*int64)(unsafe.Add(mBase, _consts[376])) = v140 + v141
	v144 = int32(4413736)
	v146 = *(*int64)(unsafe.Add(mBase, _consts[377]))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v79)+80))
	*(*int64)(unsafe.Add(mBase, _consts[377])) = v146 + v147
	v150 = int32(4413744)
	v152 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v79)+88))
	*(*int64)(unsafe.Add(mBase, _consts[378])) = v152 + v153
	v156 = int32(4413752)
	v158 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v79)+96))
	*(*int64)(unsafe.Add(mBase, _consts[379])) = v158 + v159
	v162 = int32(4413760)
	v164 = *(*int64)(unsafe.Add(mBase, _consts[380]))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v79)+104))
	*(*int64)(unsafe.Add(mBase, _consts[380])) = v164 + v165
	v168 = int32(4413768)
	v170 = *(*int64)(unsafe.Add(mBase, _consts[381]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v79)+112))
	*(*int64)(unsafe.Add(mBase, _consts[381])) = v170 + v171
	v174 = int32(4413776)
	v176 = *(*int64)(unsafe.Add(mBase, _consts[382]))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v79)+120))
	*(*int64)(unsafe.Add(mBase, _consts[382])) = v176 + v177
	v180 = int32(4413800)
	v182 = *(*int64)(unsafe.Add(mBase, _consts[14]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int64)(unsafe.Add(mBase, _consts[14])) = v182 + v183
	v186 = int32(4413784)
	v188 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, _consts[18])) = v188 + v189
	v192 = int32(4413792)
	v194 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, _consts[16])) = v194 + v195
	v198 = int32(4413808)
	v200 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v83)+24))
	*(*int64)(unsafe.Add(mBase, _consts[12])) = v200 + v201
	goto L33
L32:
	;
	goto L30
L33:
	;
	v205 = v74 + int32(1)
	if v205 != v8 {
		v74 = v205
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
}
func F_ExecParallelReinitialize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+152))
	if v8 != 0 {
		v11 = v8
		F_ExecSetParamPlanMulti(m, l2, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			F_ReinitializeParallelDSM(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v19 = F_ExecParallelSetupTupleQueues(m, v17, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v19
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)) = uint8(v21)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
					v30 = F_shm_toc_lookup(m, v27, int64(-2305843009213693951), v21)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						if v32 != 0 {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							F_dsa_free(m, v33, v32)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(0)
								if l2 != 0 {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v39 = F_SerializeParamExecParams(m, v7, l2, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v39
										*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v39
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
											return
										}
									}
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
										return
									}
								}
							}
						} else {
							if l2 != 0 {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v39 = F_SerializeParamExecParams(m, v7, l2, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v39
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
										return
									}
								}
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v9 = F_MakePerTupleExprContext(m, v7)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = v9
			F_ExecSetParamPlanMulti(m, l2, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_ReinitializeParallelDSM(m, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v19 = F_ExecParallelSetupTupleQueues(m, v17, int32(1))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v21 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v19
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)) = uint8(v21)
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
						v30 = F_shm_toc_lookup(m, v27, int64(-2305843009213693951), v21)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							if v32 != 0 {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								F_dsa_free(m, v33, v32)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(0)
									if l2 != 0 {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v39 = F_SerializeParamExecParams(m, v7, l2, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v39
											*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v39
											v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
											v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
												return
											}
										}
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
											return
										}
									}
								}
							} else {
								if l2 != 0 {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v39 = F_SerializeParamExecParams(m, v7, l2, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v39
										*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v39
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
											return
										}
									}
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v44
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v47 = F_ExecParallelReInitializeDSM(m, l0, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
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
func F_parallel_vacuum_main(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v178 int64
	_ = v178
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v310 int32
	_ = v310
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v16 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errmsg_internal(m, int32(220678), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v29 = F_shm_toc_lookup(m, l1, int64(1), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(497411), int32(1009), int32(278127))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v34 = F_shm_toc_lookup(m, l1, int64(2), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v34
	F_pgstat_report_activity(m, int32(3), v34)
	mBase = m.M
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	v43 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v79 = F_table_open(m, v77, int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v47 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v43)+392))
	if int32(1)&base.B2i32(v52 != int64(0)) != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v56 = int32(4510052)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v59 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v58 + v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v62 + v59
	*(*int64)(unsafe.Add(mBase, uint32(v43)+392)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v62 + int32(2)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v73 - v59
	goto L11
L15:
	;
	F_vac_open_indexes(m, v79, int32(3), v12+int32(16), v12+int32(20))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	if int32(0) < v88 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v88
	goto L19
L18:
	;
	goto L19
L19:
	;
	v95 = F_shm_toc_lookup(m, l1, int64(5), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v100 = F_palloc0(m, int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v102 = F_dsa_attach(m, v97)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v105 = F_palloc0(m, int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v102
	v108 = F_dsa_get_address(m, v102, v98)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v105
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v29 + int32(36)
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v29 + int32(40)
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v124
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v29
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+68))
	v138 = F_get_namespace_name(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	v144 = F_pstrdup(m, v141+int32(4))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v144
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v153 = F_GetAccessStrategyWithSize(m, v150<<(uint(int32(3))%32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(584)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v153
	v158 = int32(4508296)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[338]))
	*(*int32)(unsafe.Add(mBase, _consts[338])) = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(24)
	v171 = F___memcpy(m, int32(4413816), int32(4413656), int32(128))
	mBase = m.M
	v174 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	*(*int64)(unsafe.Add(mBase, _consts[13])) = v174
	v178 = *(*int64)(unsafe.Add(mBase, _consts[14]))
	*(*int64)(unsafe.Add(mBase, _consts[15])) = v178
	v182 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	*(*int64)(unsafe.Add(mBase, _consts[17])) = v182
	v186 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	*(*int64)(unsafe.Add(mBase, _consts[19])) = v186
	goto L29
L29:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	if v189 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v190 + int32(1)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+44)) = v195 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v195 < v199 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v203 = v195
	goto L36
L34:
	;
	goto L35
L35:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	if v243 != 0 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v213 = v210 + v203*int32(48)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
	if v214 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v203<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, v12+int32(24), v223, v213)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+44)) = v227 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v227 < v231 {
		v203 = v227
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v244 - int32(1)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v250 = F_shm_toc_lookup(m, l1, int64(3), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v254 = F_shm_toc_lookup(m, l1, int64(4), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v263 = v254 + v257<<(uint(int32(5))%32)
	v268 = F___memset(m, v250+v257<<(uint(int32(7))%32), int32(0), int32(128))
	mBase = m.M
	F_BufferUsageAccumDiff(m, v268, int32(4413816))
	mBase = m.M
	v272 = v263 + int32(24)
	v273 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v272))) = v273
	v276 = v263 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v276))) = v273
	v280 = v263 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v280))) = v273
	*(*int64)(unsafe.Add(mBase, uint32(v263))) = v273
	v286 = *(*int64)(unsafe.Add(mBase, _consts[14]))
	v288 = *(*int64)(unsafe.Add(mBase, _consts[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v276))) = v286 - v288
	v292 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	v294 = *(*int64)(unsafe.Add(mBase, _consts[19]))
	*(*int64)(unsafe.Add(mBase, uint32(v263))) = v292 - v294
	v298 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	v300 = *(*int64)(unsafe.Add(mBase, _consts[17]))
	*(*int64)(unsafe.Add(mBase, uint32(v280))) = v298 - v300
	v304 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	v306 = *(*int64)(unsafe.Add(mBase, _consts[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v272))) = v304 - v306
	goto L48
L48:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _consts[420])))
	if v310 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v315 = *(*int64)(unsafe.Add(mBase, _consts[421]))
	F_pgstat_progress_parallel_incr_param(m, int32(10), v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	F_pfree(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	F_dsa_detach(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v100)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, _consts[338])) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_vac_close_indexes(m, v329, v330, int32(3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_sequence_close(m, v79, int32(4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	F_bms_free(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	m.G0 = v12 + int32(96)
	return
}
func F_parallel_vacuum_reset_dead_items(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_TidStoreDestroy(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+56))
		v9 = F_TidStoreCreateShared(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v9
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v14
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v20
			*(*int64)(unsafe.Add(mBase, uint32(v4)+64)) = int64(0)
			return
		}
	}
}
