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
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
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
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
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
	var v362 int32
	_ = v362
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
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
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
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
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
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v607 int32
	_ = v607
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
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
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
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
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v750 int32
	_ = v750
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v846 int32
	_ = v846
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
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
	v109 = F_palloc0(m, int32(104))
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
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(4294967626)
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
	*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = v43
	v128 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+28)) = uint16(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(_a_F_ExecInitParallelPlan_0)
	*(*int64)(unsafe.Add(mBase, uint32(v109)+16)) = v126
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+40)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+48)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+64)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v109)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+52)) = v139
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+64))
	if v146 == v128 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v216 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v109)+76)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v109)+68)) = v216
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v109)+88)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+84)) = v221
	v227 = F_nodeToString(m, v109)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v109)+64)) = v186
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
	v231 = F_CreateParallelContext(m, int32(_a_F_ExecInitParallelPlan_1), int32(_a_F_ExecInitParallelPlan_2), l3)
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
	v260 = F_strlen(m, v227)
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
		v362 = v278
		goto L44
	} else {
		goto L45
	}
L44:
	;
	m.G0 = v276 + int32(16)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v387 = F_add_size(m, v382, (v362+int32(31))&int32(-32))
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
		v362 = v278
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v292 = v278
	v296 = v6
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
	v362 = v350
	goto L44
L49:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	v323 = F_add_size(m, v292, int32(4))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L54
	}
L50:
	;
	v315 = m.T0[v309].(func(*base.Module, int32, int32, int32, int32) int32)(m, v273, v296+int32(1), int32(0), v276+int32(4))
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
	v320 = v273 + int32(32) + v296*int32(12)
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
	v353 = v296 + int32(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v353 < v354 {
		v292 = v350
		v296 = v353
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
	v432 = F_mul_size(m, int32(_a_F_ExecInitParallelPlan_0), v431)
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
	v455 = v246 + int32(1)
	v456 = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v457 == v456 {
		v509 = v6
		v510 = v6
		v511 = v456
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v513 = v260 + int32(1)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v516 = F_add_size(m, v514, int32(_a_F_ExecInitParallelPlan_3))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L85
	}
L77:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v462 = F_mul_size(m, v461, l3)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v464 = F_mul_size(m, int32(416), v462)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v472 = (v461<<(uint(int32(2))%32) + int32(23)) & int32(-8)
	v473 = v464 + v472
	v478 = F_add_size(m, v466, (v473+int32(31))&int32(-32))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v478
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v483 = F_add_size(m, v481, int32(1))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v483
	v486 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v487 == v486 {
		v509 = v472
		v510 = v473
		v511 = v486
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v492 = l3 * int32(48)
	v497 = F_add_size(m, v490, (v492+int32(39))&int32(-32))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v497
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v502 = F_add_size(m, v500, int32(1))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v502
	v509 = v472
	v510 = v473
	v511 = v492 | int32(8)
	goto L76
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v516
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
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
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v521
	F_InitializeParallelDSM(m, v231)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
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
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
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
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v542 = F_shm_toc_allocate(m, v541, v455)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	if v455 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	base.MemoryCopy(m, v542, v544, v455)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v546, int64(-2305843009213693944), v542)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v551 = F_shm_toc_allocate(m, v550, v513)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	if v513 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	base.MemoryCopy(m, v551, v227, v513)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v554, int64(-2305843009213693950), v551)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v559 = F_shm_toc_allocate(m, v558, v362)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v559
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v562, int64(-2305843009213693949), v559)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v568 = v26 + int32(8)
	v569 = int32(0)
	v570 = m.G0
	v572 = v570 - int32(16)
	m.G0 = v572
	if v566 == v569 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	m.G0 = v572 + int32(16)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v700 = F_mul_size(m, int32(128), v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L124
	}
L103:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v579 + int32(4)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v566)+28))
	v585 = int32(0)
	if v585 < v584 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v588 = v584
	goto L108
L107:
	;
	v588 = v585
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v588
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v590 + int32(4)
	if v584 <= int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	v607 = v569
	goto L110
L110:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	if v621 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L102
L112:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v632)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v633))) = v634
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v636 + int32(4)
	v640 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v632)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v636)+4)) = uint16(v640)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v642 + int32(2)
	if v634 != 0 {
		goto L118
	} else {
		goto L119
	}
L113:
	;
	v627 = m.T0[v621].(func(*base.Module, int32, int32, int32, int32) int32)(m, v566, v607+int32(1), int32(0), v572+int32(4))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v632 = v566 + int32(32) + v607*int32(12)
	goto L112
L116:
	;
	v632 = v627
	goto L112
L117:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+4)))
	F_datumSerialize(m, v662, v663, v660&int32(1), v661, v568)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L122
	}
L118:
	;
	F_get_typlenbyval(m, v634, v572+int32(2), v572+int32(1))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v654 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+1)) = uint8(v654)
	v657 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v572)+2)) = uint16(v657)
	v660 = v654
	v661 = v657
	goto L117
L121:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+1)))
	v653 = int32(*(*int16)(unsafe.Add(mBase, uint32(v572)+2)))
	v660 = v652
	v661 = v653
	goto L117
L122:
	;
	v669 = v607 + int32(1)
	if v584 != v669 {
		v607 = v669
		goto L110
	} else {
		goto L123
	}
L123:
	;
	goto L111
L124:
	;
	v702 = F_shm_toc_allocate(m, v697, v700)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v704, int64(-2305843009213693948), v702)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v702
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v712 = F_mul_size(m, int32(32), v711)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v714 = F_shm_toc_allocate(m, v709, v712)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v716, int64(-2305843009213693942), v714)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v714
	v722 = F_ExecParallelSetupTupleQueues(m, v231, int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v724 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v722
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v727 == v724 {
		v846 = v446
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v861 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L132:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v731 = F_shm_toc_allocate(m, v730, v510)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v731)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v731)+4)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v731))) = v733
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = v737
	if int32(0) < l3*v737 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v750 = int32(0)
	goto L137
L135:
	;
	goto L136
L136:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v815, int64(-2305843009213693946), v731)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L140
	}
L137:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v768 = int32(416)
	v770 = v731 + v509 + v750*v768
	base.MemoryFill(m, v770, int32(0), v768)
	v774 = int32(1)
	v775 = v767 & v774
	*(*uint8)(unsafe.Add(mBase, uint32(v770))) = uint8(v775)
	v780 = int32(base.Ui32(v767)>>(uint(int32(3))%32)) & v774
	*(*uint8)(unsafe.Add(mBase, uint32(v770)+2)) = uint8(v780)
	v785 = int32(base.Ui32(v767)>>(uint(v774)%32)) & v774
	*(*uint8)(unsafe.Add(mBase, uint32(v770)+1)) = uint8(v785)
	v788 = v750 + v774
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v788 < v789*l3 {
		v750 = v788
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
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v731
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v820 == int32(0) {
		v846 = v731
		goto L131
	} else {
		goto L141
	}
L141:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v824 = F_shm_toc_allocate(m, v823, v511)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = l3
	v828 = l3 * int32(48)
	if v828 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	base.MemoryFill(m, v824+int32(8), int32(0), v828)
	goto L145
L144:
	;
	goto L145
L145:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v833, int64(-2305843009213693943), v824)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v824
	v846 = v731
	goto L131
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v231
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = v890
	v894 = F_ExecParallelInitializeDSM(m, l0, v26+int32(12))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L4
	} else {
		goto L154
	}
L148:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v866 = F_shm_toc_allocate(m, v864, int32(_a_F_ExecInitParallelPlan_3))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v868, int64(-2305843009213693945), v866)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	v875 = F_dsa_create_in_place_ext(m, v866, int32(_a_F_ExecInitParallelPlan_3), int32(71), v874)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v875
	if l2 == int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v880 = F_SerializeParamExecParams(m, l1, l2, v875)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v528)+8)) = v880
	goto L147
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = int32(0)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v898 != v899 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
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
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_ExecInitParallelPlan_5), int32(877), int32(_a_F_ExecInitParallelPlan_6))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
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
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v184 int64
	_ = v184
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
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
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int64
	_ = v267
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v300 int32
	_ = v300
	var v305 int64
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
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
	F_errfinish(m, int32(_a_F_parallel_vacuum_main_1), int32(1009), int32(_a_F_parallel_vacuum_main_2))
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v81 = F_table_open(m, v79, int32(4))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v43)+392)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v64 + int32(2)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[3])) = v75 - v61
	goto L11
L15:
	;
	F_vac_open_indexes(m, v81, int32(3), v12+int32(16), v12+int32(20))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	if int32(0) < v90 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[4])) = v90
	goto L19
L18:
	;
	goto L19
L19:
	;
	v97 = F_shm_toc_lookup(m, l1, int64(5), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v102 = F_palloc0(m, int32(12))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v104 = F_dsa_attach(m, v99)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v107 = F_palloc0(m, int32(8))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v104
	v110 = F_dsa_get_address(m, v104, v100)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v102)+8)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = v107
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[5])) = v29 + int32(36)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[6])) = v29 + int32(40)
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[7])) = v126
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[8])) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v29
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v81)+48))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+68))
	v140 = F_get_namespace_name(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v140
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v81)+48))
	v146 = F_pstrdup(m, v143+int32(4))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v146
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v155 = F_GetAccessStrategyWithSize(m, v152<<(uint(int32(3))%32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(584)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v155
	v160 = int32(_a_F_parallel_vacuum_main_4)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[9])) = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(24)
	base.MemoryCopy(m, int32(_a_F_parallel_vacuum_main_5), int32(_a_F_parallel_vacuum_main_6), int32(128))
	v176 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[11])) = v176
	v180 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[13])) = v180
	v184 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[15])) = v184
	v188 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[17])) = v188
	goto L29
L29:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[6]))
	if v191 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v194 = base.AtomicRmwAdd32(m, v191, int32(0), int32(1))
	goto L32
L31:
	;
	goto L32
L32:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v198 = base.AtomicRmwAdd32(m, v195, int32(44), int32(1))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v198 < v199 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v203 = v198
	goto L36
L34:
	;
	goto L35
L35:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[6]))
	if v242 != 0 {
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
	v229 = base.AtomicRmwAdd32(m, v226, int32(44), int32(1))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v229 < v230 {
		v203 = v229
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
	v245 = base.AtomicRmwSub32(m, v242, int32(0), int32(1))
	goto L45
L44:
	;
	goto L45
L45:
	;
	v248 = F_shm_toc_lookup(m, l1, int64(3), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v252 = F_shm_toc_lookup(m, l1, int64(4), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[18]))
	v258 = v248 + v255<<(uint(int32(7))%32)
	v261 = v252 + v255<<(uint(int32(5))%32)
	base.MemoryFill(m, v258, int32(0), int32(128))
	F_BufferUsageAccumDiff(m, v258, int32(_a_F_parallel_vacuum_main_5))
	mBase = m.M
	v267 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v261)+24)) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v261)+16)) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v261)+8)) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v267
	v276 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[12]))
	v278 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+16)) = v276 - v278
	v282 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[16]))
	v284 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[17]))
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v282 - v284
	v288 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[14]))
	v290 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+8)) = v288 - v290
	v294 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[10]))
	v296 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+24)) = v294 - v296
	goto L48
L48:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[19])))
	if v300 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v305 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[20]))
	F_pgstat_progress_parallel_incr_param(m, int32(10), v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	F_pfree(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	F_dsa_detach(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v102)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_main[9])) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_vac_close_indexes(m, v319, v320, int32(3))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_relation_close(m, v81, int32(4))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	F_bms_free(m, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
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
