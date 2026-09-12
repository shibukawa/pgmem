package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NameListToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	F_initStringInfo(m, v8+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	m.G0 = v8 + int32(32)
	return v76
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v23 = int32(0)
	goto L6
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v23<<(uint(int32(2))%32))))
	if v23&int32(1073741823) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	F_appendStringInfoChar(m, v8+int32(16), int32(46))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v38 != int32(77) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v68 = v23 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 < v69 {
		v23 = v68
		goto L6
	} else {
		goto L24
	}
L13:
	;
	if v38 == int32(468) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_appendStringInfoChar(m, v8+int32(16), int32(42))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	F_appendStringInfoString(m, v8+int32(16), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L12
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v52
	F_errmsg_internal(m, int32(462327), v8)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(480234), int32(3617), int32(317000))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L12
L24:
	;
	goto L7
}
func F_get_name_for_var_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
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
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v663 int32
	_ = v663
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
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
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
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	v14 = m.G0
	v16 = v14 - int32(256)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != int32(8) {
		goto L15
	} else {
		goto L16
	}
L1:
	;
	m.G0 = v16 + int32(256)
	return v881
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L22
	} else {
		goto L280
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L22
	} else {
		goto L277
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L22
	} else {
		goto L274
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L22
	} else {
		goto L271
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L22
	} else {
		goto L268
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L22
	} else {
		goto L265
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L22
	} else {
		goto L262
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L22
	} else {
		goto L259
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L22
	} else {
		goto L256
	}
L11:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v16)+248))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v16)+252))
	goto L245
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v69 = v68 + l2
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v70 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v57 = F_get_expr_result_tupdesc(m, l0, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L22
	} else {
		goto L27
	}
L14:
	;
	if v47 != int32(6) {
		goto L13
	} else {
		goto L25
	}
L15:
	;
	if v18 != int32(36) {
		v47 = v18
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v42 = F_find_param_referent(m, l0, l3, v16+int32(252), v16+int32(248))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if l1 <= int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v25 == int32(0) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 < l1 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+l1<<(uint(int32(2))%32)-int32(4))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v881 = v37
	goto L1
L22:
	;
	return int32(0)
L23:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = v46
	goto L14
L25:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v51 == int32(2249) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v881 = v57 + v59<<(uint(int32(4))%32) + l1*int32(100) - int32(76)
	goto L1
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v73 = v71
	goto L30
L29:
	;
	v73 = int32(0)
	goto L30
L30:
	;
	if v73 <= v69 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v69<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+252)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v81 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89+l0))))
	if v88 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = v86
	v89 = int32(8)
	goto L32
L34:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+40))
	if v84 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v88 = v81
	v89 = int32(40)
	goto L32
L36:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	switch v302 - int32(1) {
	case 0:
		goto L120
	case 1:
		goto L119
	default:
		v684 = l0
		goto L117
	case 5:
		goto L118
	}
L37:
	;
	switch v88 + int32(3) {
	case 0:
		goto L46
	case 1:
		goto L48
	case 2:
		goto L47
	default:
		goto L45
	}
L38:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v94 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v97 = v95
	goto L41
L40:
	;
	v97 = int32(0)
	goto L41
L41:
	;
	if v97 < v88 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+v88<<(uint(int32(2))%32)-int32(4))))
	if v91 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v106 = F_get_rte_attribute_name(m, v105, l1)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v881 = v106
	goto L1
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L22
	} else {
		goto L114
	}
L46:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v79)+64))
	if v241 == int32(0) {
		goto L45
	} else {
		goto L98
	}
L47:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v79)+60))
	if v172 == int32(0) {
		goto L45
	} else {
		goto L71
	}
L48:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v79)+56))
	if v111 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	if v111 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v151 == int32(0) {
		goto L9
	} else {
		goto L63
	}
L51:
	;
	goto L50
L52:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v117 <= int32(0) {
		v151 = int32(0)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v151 = int32(0)
	goto L51
L55:
	;
	v120 = int32(0)
	if v120 < v117 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v123 = v117
	goto L58
L57:
	;
	v123 = v120
	goto L58
L58:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v128 = int32(0)
	goto L59
L59:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v124+v128<<(uint(int32(2))%32))))
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+8)))
	if v137 == v91&int32(65535) {
		v151 = v136
		goto L51
	} else {
		goto L61
	}
L60:
	;
	goto L54
L61:
	;
	v140 = v128 + int32(1)
	if v140 != v123 {
		v128 = v140
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	F_push_child_plan(m, v79, v155, v16+int32(168))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v161 = F_get_name_for_var_field(m, v160, l1, l2, l3)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	v164 = F_list_delete_first(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L66
	}
L66:
	;
	goto L68
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+44)) = v164
	v881 = v161
	goto L1
L68:
	;
	v169 = F__emscripten_memcpy_bulkmem(m, v79, v16+int32(168), int32(80))
	mBase = m.M
	goto L70
L70:
	;
	goto L67
L71:
	;
	if v172 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	if v212 == int32(0) {
		goto L8
	} else {
		goto L85
	}
L73:
	;
	goto L72
L74:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v178 <= int32(0) {
		v212 = int32(0)
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v212 = int32(0)
	goto L73
L77:
	;
	v181 = int32(0)
	if v181 < v178 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v184 = v178
	goto L80
L79:
	;
	v184 = v181
	goto L80
L80:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v189 = int32(0)
	goto L81
L81:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v185+v189<<(uint(int32(2))%32))))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+8)))
	if v198 == v91&int32(65535) {
		v212 = v197
		goto L73
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v201 = v189 + int32(1)
	if v201 != v184 {
		v189 = v201
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v79)+52))
	goto L87
L86:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v79)+40))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	v224 = F_lcons(m, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L22
	} else {
		goto L90
	}
L87:
	;
	v220 = F__emscripten_memcpy_bulkmem(m, v16+int32(168), v79, int32(80))
	mBase = m.M
	goto L89
L89:
	;
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v224
	F_set_deparse_plan(m, v79, v216)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L22
	} else {
		goto L91
	}
L91:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v230 = F_get_name_for_var_field(m, v229, l1, l2, l3)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L22
	} else {
		goto L92
	}
L92:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	v233 = F_list_delete_first(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	goto L95
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v233
	v881 = v230
	goto L1
L95:
	;
	v238 = F__emscripten_memcpy_bulkmem(m, v79, v16+int32(168), int32(80))
	mBase = m.M
	goto L97
L97:
	;
	goto L94
L98:
	;
	if v241 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if v281 == int32(0) {
		goto L7
	} else {
		goto L112
	}
L100:
	;
	goto L99
L101:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	if v247 <= int32(0) {
		v281 = int32(0)
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v281 = int32(0)
	goto L100
L104:
	;
	v250 = int32(0)
	if v250 < v247 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v253 = v247
	goto L107
L106:
	;
	v253 = v250
	goto L107
L107:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	v258 = int32(0)
	goto L108
L108:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v254+v258<<(uint(int32(2))%32))))
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+8)))
	if v267 == v91&int32(65535) {
		v281 = v266
		goto L100
	} else {
		goto L110
	}
L109:
	;
	goto L103
L110:
	;
	v270 = v258 + int32(1)
	if v270 != v253 {
		v258 = v270
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v286 = F_get_name_for_var_field(m, v285, l1, l2, l3)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L22
	} else {
		goto L113
	}
L113:
	;
	v881 = v286
	goto L1
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v88
	F_errmsg_internal(m, int32(464229), v16)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L22
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(474851), int32(8172), int32(414688))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L22
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v698 = F_get_expr_result_tupdesc(m, v684, int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L22
	} else {
		goto L243
	}
L118:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v105)+88))
	v460 = v459 + v69
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if base.Ui32(v461) <= base.Ui32(v460) {
		goto L173
	} else {
		goto L174
	}
L119:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v105)+52))
	if v444 == int32(0) {
		goto L4
	} else {
		goto L170
	}
L120:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v105)+36))
	if v305 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+76))
	if v306 != 0 {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	goto L123
L123:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v79)+52))
	if v372 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L124:
	;
	if v344 == int32(0) {
		goto L6
	} else {
		goto L137
	}
L125:
	;
	goto L124
L126:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v310 <= int32(0) {
		v344 = int32(0)
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v344 = int32(0)
	goto L125
L129:
	;
	v313 = int32(0)
	if v313 < v310 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v316 = v310
	goto L132
L131:
	;
	v316 = v313
	goto L132
L132:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v306)+12))
	v321 = int32(0)
	goto L133
L133:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v317+v321<<(uint(int32(2))%32))))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+8)))
	if v330 == v91&int32(65535) {
		v344 = v329
		goto L125
	} else {
		goto L135
	}
L134:
	;
	goto L128
L135:
	;
	v333 = v321 + int32(1)
	if v333 != v316 {
		v321 = v333
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+26)))
	if v348 == int32(1) {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	if v352 != int32(6) {
		v684 = v351
		goto L117
	} else {
		goto L139
	}
L139:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v356 = F_list_copy_tail(m, v355, v69)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L22
	} else {
		goto L140
	}
L140:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v105)+36))
	F_set_deparse_for_query(m, v16+int32(168), v360, v356)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L22
	} else {
		goto L141
	}
L141:
	;
	v365 = F_lcons(m, v16+int32(168), v356)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L22
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v365
	v369 = F_get_name_for_var_field(m, v351, l1, int32(0), l3)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L22
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v355
	v881 = v369
	goto L1
L144:
	;
	v376 = F_palloc(m, int32(32))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L22
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v79)+60))
	if v385 != 0 {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l1
	v383 = F_pg_snprintf(m, v376, int32(32), int32(448034), v16+int32(112))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L22
	} else {
		goto L148
	}
L148:
	;
	v881 = v376
	goto L1
L149:
	;
	if v423 == int32(0) {
		goto L5
	} else {
		goto L162
	}
L150:
	;
	goto L149
L151:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v389 <= int32(0) {
		v423 = int32(0)
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v423 = int32(0)
	goto L150
L154:
	;
	v392 = int32(0)
	if v392 < v389 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v395 = v389
	goto L157
L156:
	;
	v395 = v392
	goto L157
L157:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v385)+12))
	v400 = int32(0)
	goto L158
L158:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v396+v400<<(uint(int32(2))%32))))
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v408)+8)))
	if v409 == v91&int32(65535) {
		v423 = v408
		goto L150
	} else {
		goto L160
	}
L159:
	;
	goto L153
L160:
	;
	v412 = v400 + int32(1)
	if v412 != v395 {
		v400 = v412
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v79)+52))
	F_push_child_plan(m, v79, v427, v16+int32(168))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L22
	} else {
		goto L163
	}
L163:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v433 = F_get_name_for_var_field(m, v432, l1, l2, l3)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L22
	} else {
		goto L164
	}
L164:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	v436 = F_list_delete_first(m, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L22
	} else {
		goto L165
	}
L165:
	;
	goto L167
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441)+44)) = v436
	v881 = v433
	goto L1
L167:
	;
	v441 = F__emscripten_memcpy_bulkmem(m, v79, v16+int32(168), int32(80))
	mBase = m.M
	goto L169
L169:
	;
	goto L166
L170:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v447+v91<<(uint(int32(2))%32)-int32(4))))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v454 != int32(6) {
		v684 = v453
		goto L117
	} else {
		goto L171
	}
L171:
	;
	v457 = F_get_name_for_var_field(m, v453, l1, v69, l3)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L22
	} else {
		goto L172
	}
L172:
	;
	v881 = v457
	goto L1
L173:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v79)+52))
	if v612 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L174:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v75+v460<<(uint(int32(2))%32))))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+16))
	if v467 == int32(0) {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	if v470 <= int32(0) {
		goto L173
	} else {
		goto L176
	}
L176:
	;
	v473 = int32(0)
	if v473 < v470 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v477 = v470
	goto L179
L178:
	;
	v477 = v473
	goto L179
L179:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v105)+84))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v467)+12))
	v480 = v473
	goto L180
L180:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v479+v480<<(uint(int32(2))%32))))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	if v501 == int32(0) {
		v520 = v500
		v521 = v501
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v496)+16))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v529 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L182:
	;
	if v521-v520 != 0 {
		goto L190
	} else {
		goto L191
	}
L183:
	;
	goto L182
L184:
	;
	if v500 != v501 {
		v520 = v500
		v521 = v501
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v505 = v497
	v506 = v478
	goto L186
L186:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+1)))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+1)))
	if v510 == int32(0) {
		v520 = v509
		v521 = v510
		goto L183
	} else {
		goto L188
	}
L187:
	;
	v520 = v509
	v521 = v510
	goto L183
L188:
	;
	v513 = int32(1)
	if v509 == v510 {
		v505 = v505 + v513
		v506 = v506 + v513
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v524 = v480 + int32(1)
	if v477 != v524 {
		v480 = v524
		goto L180
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L181
L193:
	;
	goto L173
L194:
	;
	v532 = int32(76)
	goto L196
L195:
	;
	v532 = int32(96)
	goto L196
L196:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v526+v532)))
	if v534 != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	if v572 == int32(0) {
		goto L3
	} else {
		goto L210
	}
L198:
	;
	goto L197
L199:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v538 <= int32(0) {
		v572 = int32(0)
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v572 = int32(0)
	goto L198
L202:
	;
	v541 = int32(0)
	if v541 < v538 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v544 = v538
	goto L205
L204:
	;
	v544 = v541
	goto L205
L205:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	v549 = int32(0)
	goto L206
L206:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v545+v549<<(uint(int32(2))%32))))
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+8)))
	if v558 == v91&int32(65535) {
		v572 = v557
		goto L198
	} else {
		goto L208
	}
L207:
	;
	goto L201
L208:
	;
	v561 = v549 + int32(1)
	if v561 != v544 {
		v549 = v561
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+26)))
	if v576 == int32(1) {
		goto L3
	} else {
		goto L211
	}
L211:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	if v580 != int32(6) {
		v684 = v579
		goto L117
	} else {
		goto L212
	}
L212:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v586 = F_list_copy_tail(m, v585, v460)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L22
	} else {
		goto L213
	}
L213:
	;
	F_set_deparse_for_query(m, v16+int32(168), v526, v586)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L22
	} else {
		goto L214
	}
L214:
	;
	v592 = F_lcons(m, v16+int32(168), v586)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L22
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v592
	v596 = F_get_name_for_var_field(m, v579, l1, int32(0), l3)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L22
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v585
	v881 = v596
	goto L1
L217:
	;
	v616 = F_palloc(m, int32(32))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L22
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v79)+60))
	if v625 != 0 {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l1
	v623 = F_pg_snprintf(m, v616, int32(32), int32(448034), v16-int32(-64))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L22
	} else {
		goto L221
	}
L221:
	;
	v881 = v616
	goto L1
L222:
	;
	if v663 == int32(0) {
		goto L2
	} else {
		goto L235
	}
L223:
	;
	goto L222
L224:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v629 <= int32(0) {
		v663 = int32(0)
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v663 = int32(0)
	goto L223
L227:
	;
	v632 = int32(0)
	if v632 < v629 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v635 = v629
	goto L230
L229:
	;
	v635 = v632
	goto L230
L230:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v625)+12))
	v640 = int32(0)
	goto L231
L231:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v636+v640<<(uint(int32(2))%32))))
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v648)+8)))
	if v649 == v91&int32(65535) {
		v663 = v648
		goto L223
	} else {
		goto L233
	}
L232:
	;
	goto L226
L233:
	;
	v652 = v640 + int32(1)
	if v652 != v635 {
		v640 = v652
		goto L231
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v79)+52))
	F_push_child_plan(m, v79, v667, v16+int32(168))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L22
	} else {
		goto L236
	}
L236:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v663)+4))
	v673 = F_get_name_for_var_field(m, v672, l1, l2, l3)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L22
	} else {
		goto L237
	}
L237:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	v676 = F_list_delete_first(m, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L22
	} else {
		goto L238
	}
L238:
	;
	goto L240
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v681)+44)) = v676
	v881 = v673
	goto L1
L240:
	;
	v681 = F__emscripten_memcpy_bulkmem(m, v79, v16+int32(168), int32(80))
	mBase = m.M
	goto L242
L242:
	;
	goto L239
L243:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v698)))
	v881 = v698 + v700<<(uint(int32(4))%32) + l1*int32(100) - int32(76)
	goto L1
L244:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713)+44))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+12))
	v724 = F_list_copy_tail(m, v717, (v709-v718)>>(uint(int32(2))%32)+int32(1))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L22
	} else {
		goto L248
	}
L245:
	;
	v715 = F__emscripten_memcpy_bulkmem(m, v16+int32(168), v713, int32(80))
	mBase = m.M
	goto L247
L247:
	;
	goto L244
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713)+44)) = v724
	F_set_deparse_plan(m, v713, v710)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L22
	} else {
		goto L249
	}
L249:
	;
	v730 = F_get_name_for_var_field(m, v42, l1, int32(0), l3)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L22
	} else {
		goto L250
	}
L250:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v713)+44))
	F_list_free(m, v732)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L22
	} else {
		goto L251
	}
L251:
	;
	goto L253
L252:
	;
	v881 = v730
	goto L1
L253:
	;
	v738 = F__emscripten_memcpy_bulkmem(m, v713, v16+int32(168), int32(80))
	mBase = m.M
	goto L255
L255:
	;
	goto L252
L256:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+164)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v744
	F_errmsg_internal(m, int32(450726), v16+int32(160))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L22
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(474851), int32(8082), int32(414688))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L22
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v91
	F_errmsg_internal(m, int32(464042), v16+int32(16))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L22
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(474851), int32(8124), int32(414688))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L22
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v91
	F_errmsg_internal(m, int32(464079), v16+int32(32))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L22
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(474851), int32(8143), int32(414688))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L22
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v91
	F_errmsg_internal(m, int32(464005), v16+int32(48))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L22
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(474851), int32(8161), int32(414688))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L22
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v806)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v807
	F_errmsg_internal(m, int32(456757), v16+int32(144))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L22
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(474851), int32(8214), int32(414688))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L22
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v91
	F_errmsg_internal(m, int32(463969), v16+int32(128))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L22
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(474851), int32(8282), int32(414688))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L22
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	F_errmsg_internal(m, int32(394041), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L22
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(474851), int32(8297), int32(414688))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L22
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v853
	F_errmsg_internal(m, int32(456796), v16+int32(96))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L22
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(474851), int32(8350), int32(414688))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L22
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v91
	F_errmsg_internal(m, int32(463969), v16+int32(80))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L22
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(474851), int32(8414), int32(414688))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L22
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
