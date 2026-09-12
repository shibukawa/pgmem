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
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
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
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
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
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
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
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
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
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v699 int32
	_ = v699
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
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
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
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v752 int32
	_ = v752
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v851 int32
	_ = v851
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
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
	v232 = F_CreateParallelContext(m, int32(172078), int32(292227), l3)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v237 = F_add_size(m, v235, int32(32))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v242 = F_add_size(m, v240, int32(1))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v242
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v247 = F_strlen(m, v246)
	mBase = m.M
	v252 = F_add_size(m, v245, v247&int32(-32)+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v252
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v257 = F_add_size(m, v255, int32(1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v257
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v261 = F_strlen(m, v228)
	mBase = m.M
	v266 = F_add_size(m, v260, v261&int32(-32)+int32(32))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v266
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v271 = F_add_size(m, v269, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v275 = m.G0
	v277 = v275 - int32(16)
	m.G0 = v277
	v279 = int32(4)
	if v274 == int32(0) {
		v371 = v279
		goto L44
	} else {
		goto L45
	}
L44:
	;
	m.G0 = v277 + int32(16)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v388 = F_add_size(m, v383, (v371+int32(31))&int32(-32))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L64
	}
L45:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if v282 <= int32(0) {
		v371 = v279
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v301 = v279
	v303 = v6
	goto L47
L47:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v310 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v371 = v351
	goto L44
L49:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v324 = F_add_size(m, v301, int32(4))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L54
	}
L50:
	;
	v316 = m.T0[v310].(func(*base.Module, int32, int32, int32, int32) int32)(m, v274, v303+int32(1), int32(0), v277+int32(4))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v321 = v274 + int32(32) + v303*int32(12)
	goto L49
L53:
	;
	v321 = v316
	goto L49
L54:
	;
	v327 = F_add_size(m, v324, int32(2))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v322 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+4)))
	v349 = F_datumEstimateSpace(m, v345, v346, v343&int32(1), v344)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L61
	}
L57:
	;
	F_get_typlenbyval(m, v322, v277+int32(2), v277+int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v337 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)) = uint8(v337)
	v340 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+2)) = uint16(v340)
	v343 = v337
	v344 = v340
	goto L56
L60:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	v336 = int32(*(*int16)(unsafe.Add(mBase, uint32(v277)+2)))
	v343 = v335
	v344 = v336
	goto L56
L61:
	;
	v351 = F_add_size(m, v327, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v354 = v303 + int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if v354 < v355 {
		v301 = v351
		v303 = v354
		goto L47
	} else {
		goto L63
	}
L63:
	;
	goto L48
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v388
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v393 = F_add_size(m, v391, int32(1))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v393
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v399 = F_mul_size(m, int32(128), v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v405 = F_add_size(m, v396, (v399+int32(31))&int32(-32))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v405
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v410 = F_add_size(m, v408, int32(1))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v410
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v416 = F_mul_size(m, int32(32), v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v422 = F_add_size(m, v413, (v416+int32(31))&int32(-32))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v422
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v427 = F_add_size(m, v425, int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v427
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v433 = F_mul_size(m, int32(65536), v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v439 = F_add_size(m, v430, (v433+int32(31))&int32(-32))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v439
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v444 = F_add_size(m, v442, int32(1))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v444
	v447 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v232
	v453 = F_ExecParallelEstimate(m, l0, v26+int32(24))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v455 = int32(1)
	v456 = v247 + v455
	v458 = v261 + v455
	v459 = int32(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v460 == v459 {
		v511 = v6
		v512 = v6
		v513 = v459
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v516 = F_add_size(m, v514, int32(4096))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L85
	}
L77:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v465 = F_mul_size(m, v464, l3)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v467 = F_mul_size(m, int32(416), v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v475 = (v464<<(uint(int32(2))%32) + int32(23)) & int32(-8)
	v476 = v467 + v475
	v481 = F_add_size(m, v469, (v476+int32(31))&int32(-32))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v481
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v486 = F_add_size(m, v484, int32(1))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v486
	v489 = int32(0)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v490 == v489 {
		v511 = v476
		v512 = v475
		v513 = v489
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	v495 = l3 * int32(48)
	v500 = F_add_size(m, v493, (v495+int32(39))&int32(-32))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v500
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v505 = F_add_size(m, v503, int32(1))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v505
	v511 = v476
	v512 = v475
	v513 = v495 | int32(8)
	goto L76
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v516
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	v521 = F_add_size(m, v519, int32(1))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v521
	F_InitializeParallelDSM(m, v232)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v528 = F_shm_toc_allocate(m, v526, int32(24))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v528))) = l4
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+12)) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+16)) = v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v537, int64(-2305843009213693951), v528)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v542 = F_shm_toc_allocate(m, v541, v456)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v456 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v547, int64(-2305843009213693944), v546)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L95
	}
L92:
	;
	v545 = F__emscripten_memcpy_bulkmem(m, v542, v544, v456)
	mBase = m.M
	v546 = v545
	goto L94
L93:
	;
	v546 = v542
	goto L94
L94:
	;
	goto L91
L95:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v552 = F_shm_toc_allocate(m, v551, v458)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	if v458 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v556, int64(-2305843009213693950), v555)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L101
	}
L98:
	;
	v554 = F__emscripten_memcpy_bulkmem(m, v552, v228, v458)
	mBase = m.M
	v555 = v554
	goto L100
L99:
	;
	v555 = v552
	goto L100
L100:
	;
	goto L97
L101:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v561 = F_shm_toc_allocate(m, v560, v371)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v561
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v564, int64(-2305843009213693949), v561)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v570 = v26 + int32(8)
	v571 = int32(0)
	v572 = m.G0
	v574 = v572 - int32(16)
	m.G0 = v574
	if v568 == v571 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	m.G0 = v574 + int32(16)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v702 = F_mul_size(m, int32(128), v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L126
	}
L105:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v578))) = int32(0)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v581 + int32(4)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v568)+28))
	v587 = int32(0)
	if v587 < v586 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v590 = v586
	goto L110
L109:
	;
	v590 = v587
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = v590
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v592 + int32(4)
	if v586 <= int32(0) {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v616 = v571
	goto L112
L112:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v623 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L104
L114:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v634)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v640 = v638 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v640
	v642 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v634)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v640))) = uint16(v642)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v644 + int32(2)
	if v636 != 0 {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	v629 = m.T0[v623].(func(*base.Module, int32, int32, int32, int32) int32)(m, v568, v616+int32(1), int32(0), v574+int32(4))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v634 = v568 + int32(32) + v616*int32(12)
	goto L114
L118:
	;
	v634 = v629
	goto L114
L119:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+4)))
	F_datumSerialize(m, v664, v665, v662&int32(1), v663, v570)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L4
	} else {
		goto L124
	}
L120:
	;
	F_get_typlenbyval(m, v636, v574+int32(2), v574+int32(1))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v656 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v574)+1)) = uint8(v656)
	v659 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v574)+2)) = uint16(v659)
	v662 = v656
	v663 = v659
	goto L119
L123:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+1)))
	v655 = int32(*(*int16)(unsafe.Add(mBase, uint32(v574)+2)))
	v662 = v654
	v663 = v655
	goto L119
L124:
	;
	v671 = v616 + int32(1)
	if v586 != v671 {
		v616 = v671
		goto L112
	} else {
		goto L125
	}
L125:
	;
	goto L113
L126:
	;
	v704 = F_shm_toc_allocate(m, v699, v702)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v706, int64(-2305843009213693948), v704)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v704
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v714 = F_mul_size(m, int32(32), v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v716 = F_shm_toc_allocate(m, v711, v714)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v718, int64(-2305843009213693942), v716)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v716
	v724 = F_ExecParallelSetupTupleQueues(m, v232, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v726 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v724
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v729 == v726 {
		v851 = v447
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v232)+44))
	if v865 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L134:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v733 = F_shm_toc_allocate(m, v732, v511)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v733)+4)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = v735
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+12)) = v739
	if int32(0) < l3*v739 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v752 = int32(0)
	goto L139
L137:
	;
	goto L138
L138:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v818, int64(-2305843009213693946), v733)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L143
	}
L139:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v770 = int32(416)
	v776 = F__emscripten_memset_bulkmem(m, v733+v512+v752*v770, base.I32_extend8_s(int32(0)), v770)
	mBase = m.M
	goto L141
L140:
	;
	goto L138
L141:
	;
	v777 = int32(1)
	v778 = v769 & v777
	*(*uint8)(unsafe.Add(mBase, uint32(v776))) = uint8(v778)
	v783 = int32(base.Ui32(v769)>>(uint(int32(3))%32)) & v777
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+2)) = uint8(v783)
	v788 = int32(base.Ui32(v769)>>(uint(v777)%32)) & v777
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+1)) = uint8(v788)
	v791 = v752 + v777
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v791 < v792*l3 {
		v752 = v791
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v733
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v823 == int32(0) {
		v851 = v733
		goto L133
	} else {
		goto L144
	}
L144:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v827 = F_shm_toc_allocate(m, v826, v513)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v827))) = l3
	v836 = F__emscripten_memset_bulkmem(m, v827+int32(8), base.I32_extend8_s(int32(0)), l3*int32(48))
	mBase = m.M
	goto L146
L146:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v837, int64(-2305843009213693943), v827)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v827
	v851 = v733
	goto L133
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v851
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v232
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = v893
	v897 = F_ExecParallelInitializeDSM(m, l0, v26+int32(12))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L4
	} else {
		goto L155
	}
L149:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v870 = F_shm_toc_allocate(m, v868, int32(4096))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	F_shm_toc_insert(m, v872, int64(-2305843009213693945), v870)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v232)+44))
	v879 = F_dsa_create_in_place_ext(m, v870, int32(4096), int32(71), v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v879
	if l2 == int32(0) {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v884 = F_SerializeParamExecParams(m, l1, l2, v879)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v884
	*(*int32)(unsafe.Add(mBase, uint32(v528)+8)) = v884
	goto L148
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = int32(0)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v901 != v902 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
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
	F_errmsg_internal(m, int32(182034), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(523317), int32(877), int32(297678))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
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
	v84 = int32(4460328)
	v86 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v86 + v87
	v90 = int32(4460336)
	v92 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, _consts[368])) = v92 + v93
	v96 = int32(4460344)
	v98 = *(*int64)(unsafe.Add(mBase, _consts[369]))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, _consts[369])) = v98 + v99
	v102 = int32(4460352)
	v104 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, _consts[370])) = v104 + v105
	v108 = int32(4460360)
	v110 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v79)+32))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v110 + v111
	v114 = int32(4460368)
	v116 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v79)+40))
	*(*int64)(unsafe.Add(mBase, _consts[371])) = v116 + v117
	v120 = int32(4460376)
	v122 = *(*int64)(unsafe.Add(mBase, _consts[372]))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	*(*int64)(unsafe.Add(mBase, _consts[372])) = v122 + v123
	v126 = int32(4460384)
	v128 = *(*int64)(unsafe.Add(mBase, _consts[373]))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v79)+56))
	*(*int64)(unsafe.Add(mBase, _consts[373])) = v128 + v129
	v132 = int32(4460392)
	v134 = *(*int64)(unsafe.Add(mBase, _consts[374]))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v79)+64))
	*(*int64)(unsafe.Add(mBase, _consts[374])) = v134 + v135
	v138 = int32(4460400)
	v140 = *(*int64)(unsafe.Add(mBase, _consts[375]))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v79)+72))
	*(*int64)(unsafe.Add(mBase, _consts[375])) = v140 + v141
	v144 = int32(4460408)
	v146 = *(*int64)(unsafe.Add(mBase, _consts[376]))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v79)+80))
	*(*int64)(unsafe.Add(mBase, _consts[376])) = v146 + v147
	v150 = int32(4460416)
	v152 = *(*int64)(unsafe.Add(mBase, _consts[377]))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v79)+88))
	*(*int64)(unsafe.Add(mBase, _consts[377])) = v152 + v153
	v156 = int32(4460424)
	v158 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v79)+96))
	*(*int64)(unsafe.Add(mBase, _consts[378])) = v158 + v159
	v162 = int32(4460432)
	v164 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v79)+104))
	*(*int64)(unsafe.Add(mBase, _consts[379])) = v164 + v165
	v168 = int32(4460440)
	v170 = *(*int64)(unsafe.Add(mBase, _consts[380]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v79)+112))
	*(*int64)(unsafe.Add(mBase, _consts[380])) = v170 + v171
	v174 = int32(4460448)
	v176 = *(*int64)(unsafe.Add(mBase, _consts[381]))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v79)+120))
	*(*int64)(unsafe.Add(mBase, _consts[381])) = v176 + v177
	v180 = int32(4460472)
	v182 = *(*int64)(unsafe.Add(mBase, _consts[14]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int64)(unsafe.Add(mBase, _consts[14])) = v182 + v183
	v186 = int32(4460456)
	v188 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, _consts[18])) = v188 + v189
	v192 = int32(4460464)
	v194 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, _consts[16])) = v194 + v195
	v198 = int32(4460480)
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
	F_errmsg_internal(m, int32(232443), int32(0))
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
	F_errfinish(m, int32(523300), int32(1009), int32(291970))
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
	v56 = int32(4556756)
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
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v29 + int32(36)
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v29 + int32(40)
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v124
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v124
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
	v158 = int32(4555000)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(24)
	v171 = F___memcpy(m, int32(4460488), int32(4460328), int32(128))
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
	v189 = *(*int32)(unsafe.Add(mBase, _consts[413]))
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
	v243 = *(*int32)(unsafe.Add(mBase, _consts[413]))
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
	F_BufferUsageAccumDiff(m, v268, int32(4460488))
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
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _consts[419])))
	if v310 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v315 = *(*int64)(unsafe.Add(mBase, _consts[420]))
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
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v327
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
