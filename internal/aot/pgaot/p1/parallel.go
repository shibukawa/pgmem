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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
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
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v612 int32
	_ = v612
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v757 int32
	_ = v757
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v856 int32
	_ = v856
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = F_copyObjectImpl(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v109 = F_palloc0(m, int32(120))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 <= int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v56 = v6
	goto L12
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v56<<(uint(int32(2))%32))))
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+26)) = uint8(v79)
	v82 = v56 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v82 < v83 {
		v56 = v82
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
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(4294967630)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitParallelPlan[0]))
	if v115 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = v120
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitParallelPlan[0]))
	if v123 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v120 = int64(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v115)+392))
	v120 = v119
	goto L16
L20:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v123)+400))
	v126 = v124
	goto L22
L21:
	;
	v126 = int64(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+40)) = v43
	v128 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+32)) = uint16(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = int32(_a_F_ExecInitParallelPlan_0)
	*(*int64)(unsafe.Add(mBase, uint32(v109)+16)) = v126
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+48)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+52)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+56)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v109)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+68))
	if v146 == v128 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v216 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v109)+76)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+84)) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+88)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v109)+60)) = v222
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+112)) = int64(-1)
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+100)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = v227
	v234 = F_nodeToString(m, v109)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L36
	}
L24:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v149 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v158 = int32(0)
	v160 = v6
	goto L26
L26:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v158<<(uint(int32(2))%32))))
	if v180 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L23
L28:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+37)))
	if v182 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v185 = int32(0)
	goto L30
L30:
	;
	v186 = F_lappend(m, v160, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L34
	}
L31:
	;
	v183 = v180
	goto L33
L32:
	;
	v183 = int32(0)
	goto L33
L33:
	;
	v185 = v183
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+68)) = v186
	v190 = v158 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v190 < v191 {
		v158 = v190
		v160 = v186
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	v238 = F_CreateParallelContext(m, int32(_a_F_ExecInitParallelPlan_1), int32(_a_F_ExecInitParallelPlan_2), l3)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v238
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v243 = F_add_size(m, v241, int32(32))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v243
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v248 = F_add_size(m, v246, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v248
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v253 = F_strlen(m, v252)
	mBase = m.M
	v258 = F_add_size(m, v251, v253&int32(-32)+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v258
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v263 = F_add_size(m, v261, int32(1))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v263
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v267 = F_strlen(m, v234)
	mBase = m.M
	v272 = F_add_size(m, v266, v267&int32(-32)+int32(32))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v272
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v277 = F_add_size(m, v275, int32(1))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v277
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v281 = m.G0
	v283 = v281 - int32(32)
	m.G0 = v283
	v285 = int32(4)
	if v280 == int32(0) {
		v369 = v285
		goto L44
	} else {
		goto L45
	}
L44:
	;
	m.G0 = v283 + int32(32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v394 = F_add_size(m, v389, (v369+int32(31))&int32(-32))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L64
	}
L45:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v280)+28))
	if v288 <= int32(0) {
		v369 = v285
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v299 = v285
	v301 = v6
	goto L47
L47:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	if v316 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v369 = v357
	goto L44
L49:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v330 = F_add_size(m, v299, int32(4))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L54
	}
L50:
	;
	v322 = m.T0[v316].(func(*base.Module, int32, int32, int32, int32) int32)(m, v280, v301+int32(1), int32(0), v283+int32(16))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v327 = v280 + int32(32) + v301<<(uint(int32(4))%32)
	goto L49
L53:
	;
	v327 = v322
	goto L49
L54:
	;
	v333 = F_add_size(m, v330, int32(2))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v328 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v327)))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+8)))
	v355 = F_datumEstimateSpace(m, v351, v352, v349&int32(1), v350)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L61
	}
L57:
	;
	F_get_typlenbyval(m, v328, v283+int32(14), v283+int32(13))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+13)) = uint8(v343)
	v346 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v283)+14)) = uint16(v346)
	v349 = v343
	v350 = v346
	goto L56
L60:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+13)))
	v342 = int32(*(*int16)(unsafe.Add(mBase, uint32(v283)+14)))
	v349 = v341
	v350 = v342
	goto L56
L61:
	;
	v357 = F_add_size(m, v333, v355)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v360 = v301 + int32(1)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v280)+28))
	if v360 < v361 {
		v299 = v357
		v301 = v360
		goto L47
	} else {
		goto L63
	}
L63:
	;
	goto L48
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v394
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v399 = F_add_size(m, v397, int32(1))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v399
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v405 = F_mul_size(m, int32(128), v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v411 = F_add_size(m, v402, (v405+int32(31))&int32(-32))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v411
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v416 = F_add_size(m, v414, int32(1))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v416
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v422 = F_mul_size(m, int32(40), v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v428 = F_add_size(m, v419, (v422+int32(31))&int32(-32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v428
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v433 = F_add_size(m, v431, int32(1))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v433
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v439 = F_mul_size(m, int32(_a_F_ExecInitParallelPlan_0), v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v445 = F_add_size(m, v436, (v439+int32(31))&int32(-32))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v445
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v450 = F_add_size(m, v448, int32(1))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v238
	v458 = F_ExecParallelEstimate(m, l0, v26+int32(24))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v461 = v253 + int32(1)
	v462 = int32(0)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v463 == v462 {
		v515 = v6
		v516 = v6
		v517 = v462
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v519 = v267 + int32(1)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v522 = F_add_size(m, v520, int32(_a_F_ExecInitParallelPlan_3))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L85
	}
L77:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v468 = F_mul_size(m, v467, l3)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v470 = F_mul_size(m, int32(440), v468)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v478 = (v467<<(uint(int32(2))%32) + int32(23)) & int32(-8)
	v479 = v470 + v478
	v484 = F_add_size(m, v472, (v479+int32(31))&int32(-32))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v484
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v489 = F_add_size(m, v487, int32(1))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v489
	v492 = int32(0)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v493 == v492 {
		v515 = v478
		v516 = v479
		v517 = v492
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v498 = l3 * int32(48)
	v503 = F_add_size(m, v496, (v498+int32(39))&int32(-32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v503
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v508 = F_add_size(m, v506, int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v508
	v515 = v478
	v516 = v479
	v517 = v498 | int32(8)
	goto L76
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v522
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v527 = F_add_size(m, v525, int32(1))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v527
	F_InitializeParallelDSM(m, v238)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v534 = F_shm_toc_allocate(m, v532, int32(24))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v534))) = l4
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = v539
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+16)) = v541
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v543, int64(-2305843009213693951), v534)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v548 = F_shm_toc_allocate(m, v547, v461)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	if v461 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	base.MemoryCopy(m, v548, v550, v461)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v552, int64(-2305843009213693944), v548)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v557 = F_shm_toc_allocate(m, v556, v519)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	if v519 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	base.MemoryCopy(m, v557, v234, v519)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v560, int64(-2305843009213693950), v557)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v565 = F_shm_toc_allocate(m, v564, v369)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v565
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v568, int64(-2305843009213693949), v565)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v574 = v26 + int32(8)
	v575 = int32(0)
	v576 = m.G0
	v578 = v576 - int32(32)
	m.G0 = v578
	if v572 == v575 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	m.G0 = v578 + int32(32)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v706 = F_mul_size(m, int32(128), v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L124
	}
L103:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = int32(0)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = v585 + int32(4)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v591 = int32(0)
	if v591 < v590 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v594 = v590
	goto L108
L107:
	;
	v594 = v591
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v594
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = v596 + int32(4)
	if v590 <= int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	v612 = v575
	goto L110
L110:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	if v627 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L102
L112:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v638)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v639))) = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = v642 + int32(4)
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v638)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v642)+4)) = uint16(v646)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = v648 + int32(2)
	if v640 != 0 {
		goto L118
	} else {
		goto L119
	}
L113:
	;
	v633 = m.T0[v627].(func(*base.Module, int32, int32, int32, int32) int32)(m, v572, v612+int32(1), int32(0), v578+int32(16))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v638 = v572 + int32(32) + v612<<(uint(int32(4))%32)
	goto L112
L116:
	;
	v638 = v633
	goto L112
L117:
	;
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v638)))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638)+8)))
	F_datumSerialize(m, v668, v669, v666&int32(1), v667, v574)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L4
	} else {
		goto L122
	}
L118:
	;
	F_get_typlenbyval(m, v640, v578+int32(14), v578+int32(13))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v660 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v578)+13)) = uint8(v660)
	v663 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v578)+14)) = uint16(v663)
	v666 = v660
	v667 = v663
	goto L117
L121:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+13)))
	v659 = int32(*(*int16)(unsafe.Add(mBase, uint32(v578)+14)))
	v666 = v658
	v667 = v659
	goto L117
L122:
	;
	v675 = v612 + int32(1)
	if v590 != v675 {
		v612 = v675
		goto L110
	} else {
		goto L123
	}
L123:
	;
	goto L111
L124:
	;
	v708 = F_shm_toc_allocate(m, v703, v706)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v710, int64(-2305843009213693948), v708)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v708
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v718 = F_mul_size(m, int32(40), v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v720 = F_shm_toc_allocate(m, v715, v718)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v722, int64(-2305843009213693942), v720)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v720
	v728 = F_ExecParallelSetupTupleQueues(m, v238, int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v730 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v728
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v733 == v730 {
		v856 = v230
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	if v869 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L132:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v737 = F_shm_toc_allocate(m, v736, v516)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v739
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+12)) = v743
	if int32(0) < l3*v743 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v757 = int32(0)
	goto L137
L135:
	;
	goto L136
L136:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v823, int64(-2305843009213693946), v737)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L140
	}
L137:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v774 = int32(440)
	v776 = v737 + v515 + v757*v774
	v777 = int32(0)
	base.MemoryFill(m, v776, v777, v774)
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+360)) = uint8(v777)
	v782 = int32(1)
	v783 = v773 & v782
	*(*uint8)(unsafe.Add(mBase, uint32(v776))) = uint8(v783)
	v788 = int32(base.Ui32(v773)>>(uint(int32(3))%32)) & v782
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+2)) = uint8(v788)
	v793 = int32(base.Ui32(v773)>>(uint(v782)%32)) & v782
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+1)) = uint8(v793)
	v796 = v757 + v782
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v796 < v797*l3 {
		v757 = v796
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L136
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v737
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v828 == int32(0) {
		v856 = v737
		goto L131
	} else {
		goto L141
	}
L141:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v832 = F_shm_toc_allocate(m, v831, v517)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832))) = l3
	v836 = l3 * int32(48)
	if v836 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	base.MemoryFill(m, v832+int32(8), int32(0), v836)
	goto L145
L144:
	;
	goto L145
L145:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v841, int64(-2305843009213693943), v832)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v832
	v856 = v737
	goto L131
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v856
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v238
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = v898
	v902 = F_ExecParallelInitializeDSM(m, l0, v26+int32(12))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L154
	}
L148:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	v874 = F_shm_toc_allocate(m, v872, int32(_a_F_ExecInitParallelPlan_3))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	F_shm_toc_insert(m, v876, int64(-2305843009213693945), v874)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v883 = F_dsa_create_in_place_ext(m, v874, int32(_a_F_ExecInitParallelPlan_3), int32(75), v882)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v883
	if l2 == int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v888 = F_SerializeParamExecParams(m, l1, l2, v883)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v888
	goto L147
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = int32(0)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v906 != v907 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L4
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	m.G0 = v26 + int32(32)
	return v37
L158:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInitParallelPlan_4), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_ExecInitParallelPlan_5), int32(931), int32(_a_F_ExecInitParallelPlan_6))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
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
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
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
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v216)
	goto L3
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = v76 + v74<<(uint(int32(7))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = v80 + v74*int32(40)
	v84 = int32(_a_F_ExecParallelFinish_0)
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[0]))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[0])) = v86 + v87
	v90 = int32(_a_F_ExecParallelFinish_1)
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[1]))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[1])) = v92 + v93
	v96 = int32(_a_F_ExecParallelFinish_2)
	v98 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[2]))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[2])) = v98 + v99
	v102 = int32(_a_F_ExecParallelFinish_3)
	v104 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[3]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[3])) = v104 + v105
	v108 = int32(_a_F_ExecParallelFinish_4)
	v110 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[4]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v79)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[4])) = v110 + v111
	v114 = int32(_a_F_ExecParallelFinish_5)
	v116 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[5]))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v79)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[5])) = v116 + v117
	v120 = int32(_a_F_ExecParallelFinish_6)
	v122 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[6]))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[6])) = v122 + v123
	v126 = int32(_a_F_ExecParallelFinish_7)
	v128 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[7]))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v79)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[7])) = v128 + v129
	v132 = int32(_a_F_ExecParallelFinish_8)
	v134 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[8]))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v79)+64))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[8])) = v134 + v135
	v138 = int32(_a_F_ExecParallelFinish_9)
	v140 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[9]))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v79)+72))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[9])) = v140 + v141
	v144 = int32(_a_F_ExecParallelFinish_10)
	v146 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[10]))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v79)+80))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[10])) = v146 + v147
	v150 = int32(_a_F_ExecParallelFinish_11)
	v152 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[11]))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v79)+88))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[11])) = v152 + v153
	v156 = int32(_a_F_ExecParallelFinish_12)
	v158 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[12]))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v79)+96))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[12])) = v158 + v159
	v162 = int32(_a_F_ExecParallelFinish_13)
	v164 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[13]))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v79)+104))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[13])) = v164 + v165
	v168 = int32(_a_F_ExecParallelFinish_14)
	v170 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[14]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v79)+112))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[14])) = v170 + v171
	v174 = int32(_a_F_ExecParallelFinish_15)
	v176 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[15]))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v79)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[15])) = v176 + v177
	v180 = int32(_a_F_ExecParallelFinish_16)
	v182 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[16]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[16])) = v182 + v183
	v186 = int32(_a_F_ExecParallelFinish_17)
	v188 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[17]))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[17])) = v188 + v189
	v192 = int32(_a_F_ExecParallelFinish_18)
	v194 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[18]))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[18])) = v194 + v195
	v198 = int32(_a_F_ExecParallelFinish_19)
	v200 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[19]))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v83)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[19])) = v200 + v201
	v204 = int32(_a_F_ExecParallelFinish_20)
	v206 = *(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[20]))
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v83)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_ExecParallelFinish[20])) = v206 + v207
	goto L33
L32:
	;
	goto L30
L33:
	;
	v211 = v74 + int32(1)
	if v211 != v8 {
		v74 = v211
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
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v193 int64
	_ = v193
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int64
	_ = v288
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v329 int32
	_ = v329
	var v334 int64
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
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
	F_errmsg_internal(m, int32(_a_F_parallel_vacuum_main_0), int32(0))
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
	F_errfinish(m, int32(_a_F_parallel_vacuum_main_1), int32(1225), int32(_a_F_parallel_vacuum_main_2))
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
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[0])) = v34
	F_pgstat_report_activity(m, int32(3), v34)
	mBase = m.M
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[1]))
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v89 = F_table_open(m, v87, int32(4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[2])))
	if v47&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v43)+392))
	if int32(1)&base.B2i32(v54 != int64(0)) != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v58 = int32(_a_F_parallel_vacuum_main_3)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[3]))
	v61 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[3])) = v60 + v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v64 + v61
	v68 = int32(0)
	v70 = int32(_a_F_parallel_vacuum_main_4)
	v71 = base.AtomicRmwOr32(m, v68, v70, v68)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+392)) = v39
	v76 = base.AtomicRmwOr32(m, v68, v70, v68)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v77 + v61
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[3])) = v83 - v61
	goto L11
L15:
	;
	F_vac_open_indexes(m, v89, int32(3), v12+int32(16), v12+int32(20))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	if int32(0) < v98 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[4])) = v98
	goto L19
L18:
	;
	goto L19
L19:
	;
	v105 = F_shm_toc_lookup(m, l1, int64(5), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v110 = F_palloc0(m, int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v112 = F_dsa_attach(m, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v115 = F_palloc0(m, int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = v112
	v118 = F_dsa_get_address(m, v112, v108)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v115
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+72)))
	if v123 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v29
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[5])) = v29 + int32(36)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[6])) = v29 + int32(40)
	v144 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[7])) = v144
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[8])) = v144
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v105
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v110
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	v157 = F_get_namespace_name(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[9])) = v29 + int32(80)
	F_parallel_vacuum_update_shared_delay_params(m)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v157
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v163 = F_pstrdup(m, v160+int32(4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v163
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v172 = F_GetAccessStrategyWithSize(m, v169<<(uint(int32(3))%32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(630)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v172
	v177 = int32(_a_F_parallel_vacuum_main_5)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[10])) = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(24)
	base.MemoryCopy(m, int32(_a_F_parallel_vacuum_main_6), int32(_a_F_parallel_vacuum_main_7), int32(128))
	v193 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[12])) = v193
	v197 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[14])) = v197
	v201 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[16])) = v201
	v205 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[18])) = v205
	v209 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[19]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[20])) = v209
	goto L34
L34:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[6]))
	if v212 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v215 = base.AtomicRmwAdd32(m, v212, int32(0), int32(1))
	goto L37
L36:
	;
	goto L37
L37:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v219 = base.AtomicRmwAdd32(m, v216, int32(44), int32(1))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v219 < v220 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v224 = v219
	goto L41
L39:
	;
	goto L40
L40:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[6]))
	if v263 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v234 = v231 + v224*int32(48)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+4)))
	if v235 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v224<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, v12+int32(24), v244, v234)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v250 = base.AtomicRmwAdd32(m, v247, int32(44), int32(1))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v250 < v251 {
		v224 = v250
		goto L41
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	goto L42
L48:
	;
	v266 = base.AtomicRmwSub32(m, v263, int32(0), int32(1))
	goto L50
L49:
	;
	goto L50
L50:
	;
	v269 = F_shm_toc_lookup(m, l1, int64(3), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v273 = F_shm_toc_lookup(m, l1, int64(4), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[21]))
	v279 = v269 + v276<<(uint(int32(7))%32)
	v282 = v273 + v276*int32(40)
	base.MemoryFill(m, v279, int32(0), int32(128))
	F_BufferUsageAccumDiff(m, v279, int32(_a_F_parallel_vacuum_main_6))
	mBase = m.M
	v288 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v282)+32)) = v288
	*(*int64)(unsafe.Add(mBase, uint32(v282)+24)) = v288
	*(*int64)(unsafe.Add(mBase, uint32(v282)+16)) = v288
	*(*int64)(unsafe.Add(mBase, uint32(v282)+8)) = v288
	*(*int64)(unsafe.Add(mBase, uint32(v282))) = v288
	v299 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[15]))
	v301 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+16)) = v299 - v301
	v305 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[19]))
	v307 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[20]))
	*(*int64)(unsafe.Add(mBase, uint32(v282))) = v305 - v307
	v311 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[17]))
	v313 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[18]))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+8)) = v311 - v313
	v317 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[13]))
	v319 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+24)) = v317 - v319
	v323 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[11]))
	v325 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+32)) = v323 - v325
	goto L53
L53:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[22])))
	if v329 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v334 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[23]))
	F_pgstat_progress_parallel_incr_param(m, int32(10), v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	F_pfree(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	F_dsa_detach(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v110)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[10])) = v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_vac_close_indexes(m, v348, v349, int32(3))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_relation_close(m, v89, int32(4))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	F_bms_free(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+72)))
	if v359 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[9])) = int32(0)
	goto L66
L65:
	;
	goto L66
L66:
	;
	m.G0 = v12 + int32(96)
	return
}
