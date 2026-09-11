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
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
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
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v482 int32
	_ = v482
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
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
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
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
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
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
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
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
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
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v863 int32
	_ = v863
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v962 int32
	_ = v962
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
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
	v231 = F_CreateParallelContext(m, int32(261622), l3)
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
	if v245&int32(3) == int32(0) {
		v269 = v245
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v307 = F_add_size(m, v244, v302&int32(-32)+int32(32))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L57
	}
L41:
	;
	v302 = v294 - v245
	goto L40
L42:
	;
	v273 = v269
	goto L51
L43:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v253 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v302 = int32(0)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v258 = v245
	goto L47
L47:
	;
	v262 = v258 + int32(1)
	if v262&int32(3) == int32(0) {
		v269 = v262
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v294 = v262
	goto L41
L49:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v267 != 0 {
		v258 = v262
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v282 = int32(-2139062144)
	if (int32(16843008)-v279|v279)&v282 == v282 {
		v273 = v273 + int32(4)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v288 = v273
	goto L54
L53:
	;
	goto L52
L54:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v292 != 0 {
		v288 = v288 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v294 = v288
	goto L41
L56:
	;
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v307
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v312 = F_add_size(m, v310, int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	if v228&int32(3) == int32(0) {
		v339 = v228
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v377 = F_add_size(m, v315, v372&int32(-32)+int32(32))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L76
	}
L60:
	;
	v372 = v364 - v228
	goto L59
L61:
	;
	v343 = v339
	goto L70
L62:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v323 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v372 = int32(0)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v328 = v228
	goto L66
L66:
	;
	v332 = v328 + int32(1)
	if v332&int32(3) == int32(0) {
		v339 = v332
		goto L61
	} else {
		goto L68
	}
L67:
	;
	v364 = v332
	goto L60
L68:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v337 != 0 {
		v328 = v332
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v352 = int32(-2139062144)
	if (int32(16843008)-v349|v349)&v352 == v352 {
		v343 = v343 + int32(4)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v358 = v343
	goto L73
L72:
	;
	goto L71
L73:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v362 != 0 {
		v358 = v358 + int32(1)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v364 = v358
	goto L60
L75:
	;
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v377
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v382 = F_add_size(m, v380, int32(1))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v382
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v386 = m.G0
	v388 = v386 - int32(16)
	m.G0 = v388
	v390 = int32(4)
	if v385 == int32(0) {
		v482 = v390
		goto L78
	} else {
		goto L79
	}
L78:
	;
	m.G0 = v388 + int32(16)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v499 = F_add_size(m, v494, (v482+int32(31))&int32(-32))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L98
	}
L79:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v385)+28))
	if v393 <= int32(0) {
		v482 = v390
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v412 = v390
	v414 = v6
	goto L81
L81:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v421 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v482 = v462
	goto L78
L83:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	v435 = F_add_size(m, v412, int32(4))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L88
	}
L84:
	;
	v427 = m.T0[v421].(func(*base.Module, int32, int32, int32, int32) int32)(m, v385, v414+int32(1), int32(0), v388+int32(4))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v432 = v385 + int32(32) + v414*int32(12)
	goto L83
L87:
	;
	v432 = v427
	goto L83
L88:
	;
	v438 = F_add_size(m, v435, int32(2))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v433 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+4)))
	v460 = F_datumEstimateSpace(m, v456, v457, v454&int32(1), v455)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L95
	}
L91:
	;
	F_get_typlenbyval(m, v433, v388+int32(2), v388+int32(1))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v448 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v388)+1)) = uint8(v448)
	v451 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v388)+2)) = uint16(v451)
	v454 = v448
	v455 = v451
	goto L90
L94:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+1)))
	v447 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388)+2)))
	v454 = v446
	v455 = v447
	goto L90
L95:
	;
	v462 = F_add_size(m, v438, v460)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v465 = v414 + int32(1)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v385)+28))
	if v465 < v466 {
		v412 = v462
		v414 = v465
		goto L81
	} else {
		goto L97
	}
L97:
	;
	goto L82
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v499
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v504 = F_add_size(m, v502, int32(1))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v504
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v510 = F_mul_size(m, int32(128), v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	v516 = F_add_size(m, v507, (v510+int32(31))&int32(-32))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v516
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v521 = F_add_size(m, v519, int32(1))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v521
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v527 = F_mul_size(m, int32(32), v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v533 = F_add_size(m, v524, (v527+int32(31))&int32(-32))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v533
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v538 = F_add_size(m, v536, int32(1))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v538
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v544 = F_mul_size(m, int32(65536), v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v550 = F_add_size(m, v541, (v544+int32(31))&int32(-32))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v550
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v555 = F_add_size(m, v553, int32(1))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v555
	v558 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v231
	v564 = F_ExecParallelEstimate(m, l0, v26+int32(24))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v566 = int32(1)
	v567 = v302 + v566
	v569 = v372 + v566
	v570 = int32(0)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v571 == v570 {
		v622 = v6
		v623 = v6
		v624 = v570
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v627 = F_add_size(m, v625, int32(4096))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L4
	} else {
		goto L119
	}
L111:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v576 = F_mul_size(m, v575, l3)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v578 = F_mul_size(m, int32(416), v576)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v586 = (v575<<(uint(int32(2))%32) + int32(23)) & int32(-8)
	v587 = v578 + v586
	v592 = F_add_size(m, v580, (v587+int32(31))&int32(-32))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v592
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v597 = F_add_size(m, v595, int32(1))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v597
	v600 = int32(0)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v601 == v600 {
		v622 = v587
		v623 = v586
		v624 = v600
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v606 = l3 * int32(48)
	v611 = F_add_size(m, v604, (v606+int32(39))&int32(-32))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v611
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v616 = F_add_size(m, v614, int32(1))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v616
	v622 = v587
	v623 = v586
	v624 = v606 | int32(8)
	goto L110
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v627
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v632 = F_add_size(m, v630, int32(1))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v632
	F_InitializeParallelDSM(m, v231)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v639 = F_shm_toc_allocate(m, v637, int32(24))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v639))) = l4
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+12)) = v644
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+16)) = v646
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v648, int64(-2305843009213693951), v639)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v653 = F_shm_toc_allocate(m, v652, v567)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v567 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v658, int64(-2305843009213693944), v657)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L129
	}
L126:
	;
	v656 = F__emscripten_memcpy_bulkmem(m, v653, v655, v567)
	mBase = m.M
	v657 = v656
	goto L128
L127:
	;
	v657 = v653
	goto L128
L128:
	;
	goto L125
L129:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v663 = F_shm_toc_allocate(m, v662, v569)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	if v569 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v667, int64(-2305843009213693950), v666)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L135
	}
L132:
	;
	v665 = F__emscripten_memcpy_bulkmem(m, v663, v228, v569)
	mBase = m.M
	v666 = v665
	goto L134
L133:
	;
	v666 = v663
	goto L134
L134:
	;
	goto L131
L135:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v672 = F_shm_toc_allocate(m, v671, v482)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v672
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v675, int64(-2305843009213693949), v672)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v681 = v26 + int32(8)
	v682 = int32(0)
	v683 = m.G0
	v685 = v683 - int32(16)
	m.G0 = v685
	if v679 == v682 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	m.G0 = v685 + int32(16)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v813 = F_mul_size(m, int32(128), v812)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L4
	} else {
		goto L160
	}
L139:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = int32(0)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v692 + int32(4)
	goto L138
L140:
	;
	goto L141
L141:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v679)+28))
	v698 = int32(0)
	if v698 < v697 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v701 = v697
	goto L144
L143:
	;
	v701 = v698
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696))) = v701
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v703 + int32(4)
	if v697 <= int32(0) {
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v727 = v682
	goto L146
L146:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	if v734 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L138
L148:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v745)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v751 = v749 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v751
	v753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v745)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v751))) = uint16(v753)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v755 + int32(2)
	if v747 != 0 {
		goto L154
	} else {
		goto L155
	}
L149:
	;
	v740 = m.T0[v734].(func(*base.Module, int32, int32, int32, int32) int32)(m, v679, v727+int32(1), int32(0), v685+int32(4))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v745 = v679 + int32(32) + v727*int32(12)
	goto L148
L152:
	;
	v745 = v740
	goto L148
L153:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745)+4)))
	F_datumSerialize(m, v775, v776, v773&int32(1), v774, v681)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L158
	}
L154:
	;
	F_get_typlenbyval(m, v747, v685+int32(2), v685+int32(1))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L4
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v767 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+1)) = uint8(v767)
	v770 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+2)) = uint16(v770)
	v773 = v767
	v774 = v770
	goto L153
L157:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685)+1)))
	v766 = int32(*(*int16)(unsafe.Add(mBase, uint32(v685)+2)))
	v773 = v765
	v774 = v766
	goto L153
L158:
	;
	v782 = v727 + int32(1)
	if v697 != v782 {
		v727 = v782
		goto L146
	} else {
		goto L159
	}
L159:
	;
	goto L147
L160:
	;
	v815 = F_shm_toc_allocate(m, v810, v813)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v817, int64(-2305843009213693948), v815)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v815
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v825 = F_mul_size(m, int32(32), v824)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v827 = F_shm_toc_allocate(m, v822, v825)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v829, int64(-2305843009213693942), v827)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v827
	v835 = F_ExecParallelSetupTupleQueues(m, v231, int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v837 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v837
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v835
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v840 == v837 {
		v962 = v558
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v976 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L168:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v844 = F_shm_toc_allocate(m, v843, v622)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v844)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v844)+4)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v846
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v844)+12)) = v850
	if int32(0) < l3*v850 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v863 = int32(0)
	goto L173
L171:
	;
	goto L172
L172:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v929, int64(-2305843009213693946), v844)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L177
	}
L173:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v881 = int32(416)
	v887 = F__emscripten_memset_bulkmem(m, v844+v623+v863*v881, base.I32_extend8_s(int32(0)), v881)
	mBase = m.M
	goto L175
L174:
	;
	goto L172
L175:
	;
	v888 = int32(1)
	v889 = v880 & v888
	*(*uint8)(unsafe.Add(mBase, uint32(v887))) = uint8(v889)
	v894 = int32(base.Ui32(v880)>>(uint(int32(3))%32)) & v888
	*(*uint8)(unsafe.Add(mBase, uint32(v887)+2)) = uint8(v894)
	v899 = int32(base.Ui32(v880)>>(uint(v888)%32)) & v888
	*(*uint8)(unsafe.Add(mBase, uint32(v887)+1)) = uint8(v899)
	v902 = v863 + v888
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v902 < v903*l3 {
		v863 = v902
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v844
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v934 == int32(0) {
		v962 = v844
		goto L167
	} else {
		goto L178
	}
L178:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v938 = F_shm_toc_allocate(m, v937, v624)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938))) = l3
	v947 = F__emscripten_memset_bulkmem(m, v938+int32(8), base.I32_extend8_s(int32(0)), l3*int32(48))
	mBase = m.M
	goto L180
L180:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v948, int64(-2305843009213693943), v938)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v938
	v962 = v844
	goto L167
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v962
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v231
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = v1004
	v1008 = F_ExecParallelInitializeDSM(m, l0, v26+int32(12))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L4
	} else {
		goto L189
	}
L183:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v981 = F_shm_toc_allocate(m, v979, int32(4096))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v983, int64(-2305843009213693945), v981)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	v990 = F_dsa_create_in_place_ext(m, v981, int32(4096), int32(71), v989)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v990
	if l2 == int32(0) {
		goto L182
	} else {
		goto L187
	}
L187:
	;
	v995 = F_SerializeParamExecParams(m, l1, l2, v990)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v639)+8)) = v995
	goto L182
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+172)) = int32(0)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v1012 != v1013 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L4
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	m.G0 = v26 + int32(32)
	return v37
L193:
	;
	F_errmsg_internal(m, int32(159631), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(466885), int32(877), int32(266353))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
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
	v84 = int32(4323544)
	v86 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v86 + v87
	v90 = int32(4323552)
	v92 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, _consts[368])) = v92 + v93
	v96 = int32(4323560)
	v98 = *(*int64)(unsafe.Add(mBase, _consts[369]))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, _consts[369])) = v98 + v99
	v102 = int32(4323568)
	v104 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, _consts[370])) = v104 + v105
	v108 = int32(4323576)
	v110 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v79)+32))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v110 + v111
	v114 = int32(4323584)
	v116 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v79)+40))
	*(*int64)(unsafe.Add(mBase, _consts[371])) = v116 + v117
	v120 = int32(4323592)
	v122 = *(*int64)(unsafe.Add(mBase, _consts[372]))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	*(*int64)(unsafe.Add(mBase, _consts[372])) = v122 + v123
	v126 = int32(4323600)
	v128 = *(*int64)(unsafe.Add(mBase, _consts[373]))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v79)+56))
	*(*int64)(unsafe.Add(mBase, _consts[373])) = v128 + v129
	v132 = int32(4323608)
	v134 = *(*int64)(unsafe.Add(mBase, _consts[374]))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v79)+64))
	*(*int64)(unsafe.Add(mBase, _consts[374])) = v134 + v135
	v138 = int32(4323616)
	v140 = *(*int64)(unsafe.Add(mBase, _consts[375]))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v79)+72))
	*(*int64)(unsafe.Add(mBase, _consts[375])) = v140 + v141
	v144 = int32(4323624)
	v146 = *(*int64)(unsafe.Add(mBase, _consts[376]))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v79)+80))
	*(*int64)(unsafe.Add(mBase, _consts[376])) = v146 + v147
	v150 = int32(4323632)
	v152 = *(*int64)(unsafe.Add(mBase, _consts[377]))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v79)+88))
	*(*int64)(unsafe.Add(mBase, _consts[377])) = v152 + v153
	v156 = int32(4323640)
	v158 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v79)+96))
	*(*int64)(unsafe.Add(mBase, _consts[378])) = v158 + v159
	v162 = int32(4323648)
	v164 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v79)+104))
	*(*int64)(unsafe.Add(mBase, _consts[379])) = v164 + v165
	v168 = int32(4323656)
	v170 = *(*int64)(unsafe.Add(mBase, _consts[380]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v79)+112))
	*(*int64)(unsafe.Add(mBase, _consts[380])) = v170 + v171
	v174 = int32(4323664)
	v176 = *(*int64)(unsafe.Add(mBase, _consts[381]))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v79)+120))
	*(*int64)(unsafe.Add(mBase, _consts[381])) = v176 + v177
	v180 = int32(4323688)
	v182 = *(*int64)(unsafe.Add(mBase, _consts[14]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int64)(unsafe.Add(mBase, _consts[14])) = v182 + v183
	v186 = int32(4323672)
	v188 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, _consts[18])) = v188 + v189
	v192 = int32(4323680)
	v194 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, _consts[16])) = v194 + v195
	v198 = int32(4323696)
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
	F_errmsg_internal(m, int32(207766), int32(0))
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
	F_errfinish(m, int32(466868), int32(1009), int32(261408))
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
	v56 = int32(4419940)
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(583)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v153
	v158 = int32(4418184)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(24)
	v171 = F___memcpy(m, int32(4323704), int32(4323544), int32(128))
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
	F_BufferUsageAccumDiff(m, v268, int32(4323704))
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
