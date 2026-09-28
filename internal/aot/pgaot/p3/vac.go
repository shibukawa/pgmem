package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_update_datfrozenxid(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
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
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v289 int64
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v304 int64
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
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
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int64
	_ = v516
	var v521 int32
	_ = v521
	var __phi521 int32
	_ = __phi521
	var v527 int32
	_ = v527
	var __phi527 int32
	_ = __phi527
	var v528 int32
	_ = v528
	var __phi528 int32
	_ = __phi528
	var v530 int32
	_ = v530
	var __phi530 int32
	_ = __phi530
	var v537 int64
	_ = v537
	var __phi537 int64
	_ = __phi537
	var v538 int64
	_ = v538
	var __phi538 int64
	_ = __phi538
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int64
	_ = v556
	var v557 int64
	_ = v557
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int64
	_ = v571
	var v572 int64
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int64
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int64
	_ = v624
	var v627 int32
	_ = v627
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int64
	_ = v645
	var v646 int64
	_ = v646
	var v652 int32
	_ = v652
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int64
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v741 int64
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int64
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int64
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v779 int64
	_ = v779
	var v780 int32
	_ = v780
	var v782 int64
	_ = v782
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int64
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int64
	_ = v843
	var v844 int64
	_ = v844
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int64
	_ = v866
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v909 int64
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v937 int64
	_ = v937
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	v1 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(16908288)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+4)) = int64(0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v34
	v39 = F_LockAcquire(m, v27, int32(7), v1, v1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	m.G0 = v27 + int32(16)
	v45 = F_GetOldestNonRemovableTransactionId(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v47 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v49 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v51 = base.I32_wrap_i64(v49)
	v52 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v56 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v23 + int32(96)
	return
L8:
	;
	F_systable_endscan(m, v63)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L271
	}
L9:
	;
	v58 = int32(0)
	v63 = F_systable_beginscan(m, v56, v58, v58, v58, v58, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v65 = F_systable_getnext(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v65 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v69 = v65
	v72 = v47
	v73 = v45
	goto L15
L13:
	;
	v143 = v47
	v144 = v45
	goto L14
L14:
	;
	F_systable_endscan(m, v63)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L41
	}
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
	v91 = v89 + v90
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+119)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+136))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+140))
	v96 = v92 - int32(109)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v96))|base.B2i32(int32(1)<<(uint(v96)%32)&int32(161) == int32(0)) != 0 {
		v136 = v72
		v137 = v73
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v143 = v136
	v144 = v137
	goto L14
L17:
	;
	v138 = F_systable_getnext(m, v63)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L39
	}
L18:
	;
	if v93 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if base.B2i32(base.Ui32(v51) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v93) < base.Ui32(int32(3))) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v126 = v73
	goto L21
L21:
	;
	if v94 == int32(0) {
		v136 = v72
		v137 = v126
		goto L17
	} else {
		goto L34
	}
L22:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v93))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v73)) != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	if int32(0) <= v51-v93 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v51) < base.Ui32(v93) {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	goto L8
L27:
	;
	goto L22
L28:
	;
	v124 = int32(base.Ui32(v93-v73) >> (uint(int32(31)) % 32))
	goto L30
L29:
	;
	v124 = base.B2i32(base.Ui32(v93) < base.Ui32(v73))
	goto L30
L30:
	;
	if v124 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v125 = v93
	goto L33
L32:
	;
	v125 = v73
	goto L33
L33:
	;
	v126 = v125
	goto L21
L34:
	;
	if v52-v94 < int32(0) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v94-v72 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v135 = v94
	goto L38
L37:
	;
	v135 = v72
	goto L38
L38:
	;
	v136 = v135
	v137 = v126
	goto L17
L39:
	;
	if v138 != 0 {
		v69 = v138
		v72 = v136
		v73 = v137
		goto L15
	} else {
		goto L40
	}
L40:
	;
	goto L16
L41:
	;
	F_relation_close(m, v56, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v167 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v170 = v23 + int32(32)
	v175 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0])))
	F_ScanKeyInit(m, v170, int32(1), int32(3), int32(184), v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_systable_inplace_update_begin(m, v167, int32(2672), v170, v23+int32(92), v23+int32(28))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	if v185 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v186 = int32(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v189 = v187 + v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+84))
	if v144 != v190 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L48
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L268
	}
L49:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v189)+88))
	v224 = int32(0)
	if base.B2i32(v219 == v143)|base.B2i32(v224 <= v219-v143|(v52-v219)) == v224 {
		goto L68
	} else {
		goto L69
	}
L50:
	;
	v192 = int32(3)
	if base.B2i32(base.Ui32(v144) < base.Ui32(v192))|base.B2i32(base.Ui32(v190) < base.Ui32(v192)) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v216 = v186
	goto L52
L52:
	;
	v217 = v144
	v218 = v216
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+84)) = v144
	v216 = int32(1)
	goto L52
L54:
	;
	v203 = int32(3)
	if base.B2i32(base.Ui32(v51) < base.Ui32(v203))|base.B2i32(base.Ui32(v190) < base.Ui32(v203)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	if int32(0) <= v190-v144 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v190) < base.Ui32(v144) {
		goto L53
	} else {
		goto L59
	}
L58:
	;
	goto L53
L59:
	;
	goto L54
L60:
	;
	if v51-v190 < int32(0) {
		goto L53
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v190) <= base.Ui32(v51) {
		v217 = v190
		v218 = v186
		goto L49
	} else {
		goto L64
	}
L63:
	;
	v217 = v190
	v218 = v186
	goto L49
L64:
	;
	goto L53
L65:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	F_pfree(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L74
	}
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	F_systable_inplace_update_cancel(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L73
	}
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	F_systable_inplace_update_finish(m, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L72
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+88)) = v143
	v232 = v143
	goto L67
L69:
	;
	goto L70
L70:
	;
	if v218 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v232 = v219
	goto L67
L72:
	;
	v242 = v232
	v243 = int32(1)
	goto L65
L73:
	;
	v242 = v219
	v243 = int32(0)
	goto L65
L74:
	;
	F_relation_close(m, v167, int32(3))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v243 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v252 = int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v258 = F_LWLockAcquire(m, v254+int32(384), v252)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v304 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L91
	}
L79:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[2]))
	v262 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v261)+36)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v261)+20))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v267+int32(384))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if base.B2i32(v263 == int32(0))|base.B2i32(base.Ui32(v265) < base.Ui32(int32(3))) != 0 {
		v296 = v252
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v296 == int32(0) {
		goto L7
	} else {
		goto L90
	}
L82:
	;
	v277 = int32(3)
	if base.B2i32(base.Ui32(v264) < base.Ui32(v277))|base.B2i32(base.Ui32(v263) < base.Ui32(v277)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v289 = int64(0)
	v292 = F_SearchSysCacheExists(m, int32(21), v262, v289, v289, v289)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L89
	}
L84:
	;
	if v264-v263 < int32(0) {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v263) <= base.Ui32(v264) {
		v296 = v252
		goto L81
	} else {
		goto L88
	}
L87:
	;
	v296 = v252
	goto L81
L88:
	;
	goto L83
L89:
	;
	v296 = v292 ^ int32(1)
	goto L81
L90:
	;
	goto L78
L91:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v311 = F_LWLockAcquire(m, v307+int32(_a_F_vac_update_datfrozenxid_0), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	v317 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+188))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	m.T0[v455].(func(*base.Module, int32))(m, v321)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L139
	}
L94:
	;
	v319 = int32(0)
	v321 = F_table_beginscan_catalog(m, v317, v319, v319)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v323 = F_heap_getnext(m, v321)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v323 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v434 = v217
	v436 = v242
	v437 = v314
	v438 = v314
	v447 = v1
	v448 = v1
	goto L93
L98:
	;
	goto L99
L99:
	;
	v327 = base.I32_wrap_i64(v304)
	v330 = v323
	v331 = v217
	v333 = v242
	v334 = v314
	v335 = v314
	v344 = v1
	v345 = v1
	goto L100
L100:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v330)+16))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+22)))
	v352 = v350 + v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+84))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v352)+88))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v352)+80))
	goto L103
L101:
	;
	v434 = v425
	v436 = v426
	v437 = v427
	v438 = v428
	v447 = v429
	v448 = v430
	goto L93
L102:
	;
	v431 = F_heap_getnext(m, v321)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L137
	}
L103:
	;
	if v355 == int32(-2) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v360 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if base.B2i32(base.Ui32(v51) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v353) < base.Ui32(int32(3))) == int32(0) {
		goto L114
	} else {
		goto L115
	}
L107:
	;
	if v360 == int32(0) {
		v425 = v331
		v426 = v333
		v427 = v334
		v428 = v335
		v429 = v344
		v430 = v345
		goto L102
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v352 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_1), v23+int32(16))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_2), int32(1899), int32(_a_F_vac_update_datfrozenxid_3))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v425 = v331
	v426 = v333
	v427 = v334
	v428 = v335
	v429 = v344
	v430 = v345
	goto L102
L111:
	;
	v391 = int32(3)
	if base.B2i32(base.Ui32(v327) < base.Ui32(v391))|base.B2i32(base.Ui32(v353) < base.Ui32(v391)) == int32(0) {
		goto L123
	} else {
		goto L124
	}
L112:
	;
	v390 = int32(1)
	goto L111
L113:
	;
	if int32(0) <= v52-v354 {
		v390 = v344
		goto L111
	} else {
		goto L119
	}
L114:
	;
	if int32(0) <= v51-v353 {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if base.Ui32(v51) < base.Ui32(v353) {
		goto L112
	} else {
		goto L118
	}
L117:
	;
	goto L112
L118:
	;
	goto L113
L119:
	;
	goto L112
L120:
	;
	if v354-v333 < int32(0) {
		goto L134
	} else {
		goto L135
	}
L121:
	;
	v415 = v331
	v416 = v334
	v417 = int32(1)
	goto L120
L122:
	;
	v402 = int32(3)
	if base.B2i32(base.Ui32(v331) < base.Ui32(v402))|base.B2i32(base.Ui32(v353) < base.Ui32(v402)) == int32(0) {
		goto L129
	} else {
		goto L130
	}
L123:
	;
	if int32(0) <= v327-v353 {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if base.Ui32(v327) < base.Ui32(v353) {
		goto L121
	} else {
		goto L127
	}
L126:
	;
	goto L121
L127:
	;
	goto L122
L128:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v415 = v353
	v416 = v413
	v417 = v345
	goto L120
L129:
	;
	if v353-v331 < int32(0) {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if base.Ui32(v353) < base.Ui32(v331) {
		goto L128
	} else {
		goto L133
	}
L132:
	;
	v415 = v331
	v416 = v334
	v417 = v345
	goto L120
L133:
	;
	v415 = v331
	v416 = v334
	v417 = v345
	goto L120
L134:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v422 = v354
	v423 = v421
	goto L136
L135:
	;
	v422 = v333
	v423 = v335
	goto L136
L136:
	;
	v425 = v415
	v426 = v422
	v427 = v416
	v428 = v423
	v429 = v390
	v430 = v417
	goto L102
L137:
	;
	if v431 != 0 {
		v330 = v431
		v331 = v425
		v333 = v426
		v334 = v427
		v335 = v428
		v344 = v429
		v345 = v430
		goto L100
	} else {
		goto L138
	}
L138:
	;
	goto L101
L139:
	;
	F_relation_close(m, v317, int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if v448 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v1000+int32(_a_F_vac_update_datfrozenxid_0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L267
	}
L142:
	;
	v463 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if v447 != 0 {
		goto L141
	} else {
		goto L150
	}
L145:
	;
	if v463 == int32(0) {
		goto L141
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_4), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v473 = F_errdetail(m, int32(_a_F_vac_update_datfrozenxid_5), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_2), int32(1945), int32(_a_F_vac_update_datfrozenxid_3))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	goto L141
L150:
	;
	v480 = int32(0)
	v482 = m.G0
	v484 = v482 - int32(16)
	m.G0 = v484
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v491 = F_LWLockAcquire(m, v487+int32(_a_F_vac_update_datfrozenxid_6), int32(1))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v498 = F_LWLockAcquire(m, v494+int32(3456), int32(1))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[3]))
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v501)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v484))) = v502
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v501)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+8)) = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v501)+8))
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v501)))
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v509+int32(3456))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
	if base.B2i32(v506 == v514)&base.B2i32(v516 == v507) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v674+int32(_a_F_vac_update_datfrozenxid_6))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L186
	}
L155:
	;
	__phi521 = v514
	__phi527 = int32(-1)
	__phi528 = v480
	__phi530 = v480
	__phi537 = int64(-1)
	__phi538 = v516
	v521 = __phi521
	v527 = __phi527
	v528 = __phi528
	v530 = __phi530
	v537 = __phi537
	v538 = __phi538
	goto L156
L156:
	;
	if v537 != v538 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v591 < int32(0) {
		goto L154
	} else {
		goto L181
	}
L158:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	if int32(0) <= v527 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v591 = v527
	v592 = v528
	v593 = v530
	goto L160
L160:
	;
	v594 = v521 + v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+8))
	v596 = int32(3)
	v602 = int32(0)
	if base.B2i32(base.Ui32(v595) < base.Ui32(v596))|base.B2i32(base.Ui32(v434) < base.Ui32(v596))|base.B2i32(v602 <= v595-v434) == v602 {
		goto L170
	} else {
		goto L171
	}
L161:
	;
	if v528 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v566 = v528
	v568 = v543
	goto L163
L163:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+28))
	v571 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[5])))
	v572 = base.I64_rem_s(v538, v571)
	v578 = F_LWLockAcquire(m, v569+base.I32_wrap_i64(v572)<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L168
	}
L164:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v543)+12))
	v548 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v546+v527))) = uint8(v548)
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v552 = v551
	goto L166
L165:
	;
	v552 = v543
	goto L166
L166:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v552)+28))
	v556 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[5])))
	v557 = base.I64_rem_s(v537, v556)
	F_LWLockRelease(m, v554+base.I32_wrap_i64(v557)<<(uint(int32(7))%32))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v566 = int32(0)
	v568 = v565
	goto L163
L168:
	;
	v582 = F_SimpleLruReadPage(m, int32(_a_F_vac_update_datfrozenxid_7), v538, int32(1), v484)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v586+v582<<(uint(int32(2))%32))))
	v591 = v582
	v592 = v566
	v593 = v590
	goto L160
L170:
	;
	v609 = F_TransactionIdDidCommit(m, v595)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	v614 = v592
	goto L172
L172:
	;
	v615 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	v618 = v616 + v617
	v622 = base.B2i32(base.Ui32(v618-int32(_a_F_vac_update_datfrozenxid_8)) < base.Ui32(int32(-8193)))
	v624 = v615 + base.I64_extend_i32_u(v622)
	*(*int64)(unsafe.Add(mBase, uint32(v484))) = v624
	if base.Ui32(v618-int32(_a_F_vac_update_datfrozenxid_8)) < base.Ui32(int32(-8193)) {
		goto L177
	} else {
		goto L178
	}
L173:
	;
	if v609 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v611 = int32(2)
	goto L176
L175:
	;
	v611 = int32(0)
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594)+8)) = v611
	v614 = int32(1)
	goto L172
L177:
	;
	v627 = int32(0)
	goto L179
L178:
	;
	v627 = v618
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v484)+8)) = v627
	if base.B2i32(v627 != v506)|base.B2i32(v624 != v507) != 0 {
		__phi521 = v627
		__phi527 = v591
		__phi528 = v614
		__phi530 = v593
		__phi537 = v538
		__phi538 = v624
		v521 = __phi521
		v527 = __phi527
		v528 = __phi528
		v530 = __phi530
		v537 = __phi537
		v538 = __phi538
		goto L156
	} else {
		goto L180
	}
L180:
	;
	goto L157
L181:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	if v614 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v638 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v636+v591))) = uint8(v638)
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v642 = v641
	goto L184
L183:
	;
	v642 = v635
	goto L184
L184:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+28))
	v645 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[5])))
	v646 = base.I64_rem_s(v538, v645)
	F_LWLockRelease(m, v643+base.I32_wrap_i64(v646)<<(uint(int32(7))%32))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L154
L186:
	;
	m.G0 = v484 + int32(16)
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v687 = F_LWLockAcquire(m, v683+int32(_a_F_vac_update_datfrozenxid_9), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[2]))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+40))
	if v691 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v707+int32(_a_F_vac_update_datfrozenxid_9))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L196
	}
L189:
	;
	v694 = int32(3)
	if base.B2i32(base.Ui32(v434) < base.Ui32(v694))|base.B2i32(base.Ui32(v691) < base.Ui32(v694)) == int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v690)+40)) = v434
	goto L188
L191:
	;
	if v691-v434 < int32(0) {
		goto L190
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	if base.Ui32(v434) <= base.Ui32(v691) {
		goto L188
	} else {
		goto L195
	}
L194:
	;
	goto L188
L195:
	;
	goto L190
L196:
	;
	v712 = m.G0
	v714 = v712 - int32(32)
	m.G0 = v714
	*(*int64)(unsafe.Add(mBase, uint32(v714)+8)) = base.I64_extend_i32_u(int32(base.Ui32(v434) >> (uint(int32(15)) % 32)))
	v724 = F_SlruScanDirectory(m, int32(_a_F_vac_update_datfrozenxid_10), int32(297), v714+int32(8))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	if v724 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	F_AdvanceOldestClogXid(m, v434)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v750 = int32(32)
	m.G0 = v714 + v750
	v753 = m.G0
	v755 = v753 - v750
	m.G0 = v755
	v758 = base.I32_div_u_s(v434, int32(819))
	*(*int64)(unsafe.Add(mBase, uint32(v755)+8)) = base.I64_extend_i32_u(v758)
	v765 = F_SlruScanDirectory(m, int32(_a_F_vac_update_datfrozenxid_11), int32(297), v755+int32(8))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L207
	}
L201:
	;
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v714)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v714)+28)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v714)+24)) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v714)+16)) = v728
	F_XLogBeginInsert(m)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v734 = int32(16)
	F_XLogRegisterData(m, v714+v734, v734)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v741 = F_XLogInsert(m, int32(3), int32(16))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_XLogFlush(m, v741)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v746 = *(*int64)(unsafe.Add(mBase, uint32(v714)+8))
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_10), v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	goto L200
L207:
	;
	if v765 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v755)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v755)+24)) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v755)+16)) = v767
	F_XLogBeginInsert(m)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	m.G0 = v755 + int32(32)
	v789 = m.G0
	v791 = v789 - int32(80)
	m.G0 = v791
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v798 = F_LWLockAcquire(m, v794+int32(_a_F_vac_update_datfrozenxid_12), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L215
	}
L211:
	;
	F_XLogRegisterData(m, v755+int32(16), int32(12))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v779 = F_XLogInsert(m, int32(18), int32(16))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v755)+8))
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_11), v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	goto L210
L215:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v805 = F_LWLockAcquire(m, v801+int32(1664), int32(1))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[6]))
	v809 = *(*int64)(unsafe.Add(mBase, uint32(v808)+8))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v808)+20))
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v813+int32(1664))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	if v436-v811 <= int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v967+int32(_a_F_vac_update_datfrozenxid_12))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L264
	}
L219:
	;
	goto L218
L220:
	;
	goto L221
L221:
	;
	if v436 != v810 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v824 = F_find_multixact_start(m, v436, v791+int32(56))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	v844 = v809
	goto L224
L224:
	;
	if v844 == int64(0) {
		goto L235
	} else {
		goto L236
	}
L225:
	;
	if v824 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v830 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v791)+56))
	v844 = v843
	goto L224
L229:
	;
	if v830 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791)+48)) = v436
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_13), v791+int32(48))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	goto L218
L233:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_14), int32(2732), int32(_a_F_vac_update_datfrozenxid_15))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	v849 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v862 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L244
	}
L238:
	;
	if v849 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = v436
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_16), v791)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	goto L218
L242:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_14), int32(2749), int32(_a_F_vac_update_datfrozenxid_15))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	if v862 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v791)+32)) = v844
	v866 = base.I64_div_u_s(v844, int64(1636))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+40)) = int64(base.Ui64(v866) >> (uint(int64(5)) % 64))
	*(*int32)(unsafe.Add(mBase, uint32(v791)+16)) = v436
	*(*int64)(unsafe.Add(mBase, uint32(v791)+24)) = base.I64_extend_i32_u(int32(base.Ui32(v436) >> (uint(int32(15)) % 32)))
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_17), v791+int32(16))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v885 = int32(_a_F_vac_update_datfrozenxid_18)
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7]))
	v888 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7])) = v887 + v888
	v892 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[8]))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v892)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v892)+336)) = v893 | v888
	*(*int64)(unsafe.Add(mBase, uint32(v791)+72)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v791)+68)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v791)+64)) = v438
	F_XLogBeginInsert(m)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L250
	}
L248:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_14), int32(2760), int32(_a_F_vac_update_datfrozenxid_15))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	F_XLogRegisterData(m, v791-int32(-64), int32(16))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v909 = F_XLogInsert(m, int32(6), int32(48))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_XLogFlush(m, v909)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v918 = F_LWLockAcquire(m, v914+int32(1664), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v921 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v921)+32)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v921)+24)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v921)+20)) = v436
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v926+int32(1664))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	if v844 != int64(1) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v937 = base.I64_div_u_s(v844-int64(1), int64(1636))
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_19), v937)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v942 = int32(1)
	if v436 == v942 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	goto L258
L260:
	;
	v948 = int32(_a_F_vac_update_datfrozenxid_20)
	goto L262
L261:
	;
	v948 = int32(base.Ui32(v436-v942) >> (uint(int32(10)) % 32))
	goto L262
L262:
	;
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_21), base.I64_extend_i32_u(v948))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v953 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[8]))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v953)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v953)+336)) = v954 & int32(-2)
	v958 = int32(_a_F_vac_update_datfrozenxid_18)
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7])) = v960 - int32(1)
	goto L218
L264:
	;
	m.G0 = v791 + int32(80)
	F_SetTransactionIdLimit(m, v434, v437)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_SetMultiXactIdLimit(m, v436, v438)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	goto L141
L267:
	;
	goto L7
L268:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v1010
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_22), v23)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_2), int32(1768), int32(_a_F_vac_update_datfrozenxid_23))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
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
	F_relation_close(m, v56, int32(1))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	goto L7
}
func F_vac_update_relstats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 float32
	_ = v48
	var v49 float32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	v17 = m.G0
	v19 = v17 - int32(128)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = v19 - int32(-64)
	F_ScanKeyInit(m, v27, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v21))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_systable_inplace_update_begin(m, v24, int32(2662), v27, v19+int32(60), v19+int32(56))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v41 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+96))
	if l1 != v45 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L75
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = l1
	goto L10
L9:
	;
	goto L10
L10:
	;
	v48 = base.F32_demote_f64(l2)
	v49 = *(*float32)(unsafe.Add(mBase, uint32(v44)+100))
	if base.F32_eq(v48, v49) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = base.B2i32(l1 != v45)
	goto L13
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v44)+100)) = v48
	v54 = int32(1)
	goto L13
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+104))
	if l3 != v55 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+104)) = l3
	v59 = int32(1)
	goto L16
L15:
	;
	v59 = v54
	goto L16
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+108))
	if l4 != v60 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = l4
	v64 = int32(1)
	goto L19
L18:
	;
	v64 = v59
	goto L19
L19:
	;
	if l10 != 0 {
		v89 = v64
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v44)+136))
	if l8 != 0 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	if l5 != 0 {
		v73 = v64
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+124)))
	if v74 != int32(1) {
		v81 = v73
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+116)))
	if v65&int32(1) == int32(0) {
		v73 = v64
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+116)) = uint8(v70)
	v73 = int32(1)
	goto L22
L25:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+125)))
	if v82 != int32(1) {
		v89 = v81
		goto L20
	} else {
		goto L28
	}
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v77 != 0 {
		v81 = v73
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+124)) = uint8(v78)
	v81 = int32(1)
	goto L25
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v85 != 0 {
		v89 = v81
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+125)) = uint8(v86)
	v89 = int32(1)
	goto L20
L30:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v91)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v93 = int32(0)
	if base.B2i32(v90 == l6)|base.B2i32(base.Ui32(l6) < base.Ui32(int32(3))) != 0 {
		v122 = v89
		v123 = v93
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v44)+140))
	if l9 != 0 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v101 = int32(0)
	v103 = base.B2i32(base.Ui32(int32(2)) < base.Ui32(v90)) & base.B2i32(v101 <= v90-l6)
	if v103 == v101 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+136)) = l6
	if l8 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v106 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v108 = base.I32_wrap_i64(v106)
	if base.Ui32(v108) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	if int32(0) <= v108-v90 {
		v122 = v89
		v123 = v93
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v122 = int32(1)
	v123 = v103
	goto L33
L41:
	;
	goto L42
L42:
	;
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v119)
	v122 = v119
	v123 = v103
	goto L33
L43:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v127)
	goto L45
L44:
	;
	goto L45
L45:
	;
	if base.B2i32(l7 == int32(0))|base.B2i32(v126 == l7) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	F_relation_close(m, v24, int32(3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L60
	}
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	F_systable_inplace_update_cancel(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L59
	}
L48:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_systable_inplace_update_finish(m, v154, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L58
	}
L49:
	;
	v149 = int32(0)
	if v122 == v149 {
		goto L47
	} else {
		goto L57
	}
L50:
	;
	v133 = v126 - l7
	if int32(0) <= v133 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v136 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v141 = int32(0)
	v142 = base.B2i32(v141 <= v133)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+140)) = l7
	if l9 == v141 {
		v153 = v142
		goto L48
	} else {
		goto L56
	}
L54:
	;
	if int32(0) <= v136-v126 {
		goto L49
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v146)
	v153 = v142
	goto L48
L57:
	;
	v153 = v149
	goto L48
L58:
	;
	v162 = v153
	goto L46
L59:
	;
	v162 = v149
	goto L46
L60:
	;
	if v123 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v162 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	v170 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v170 == int32(0) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v177 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vac_update_relstats_3), v19+int32(32))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_vac_update_relstats_1), int32(1585), int32(_a_F_vac_update_relstats_2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	m.G0 = v19 + int32(128)
	return
L69:
	;
	v198 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v198 == int32(0) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v205 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vac_update_relstats_0), v19+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_vac_update_relstats_1), int32(1591), int32(_a_F_vac_update_relstats_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L68
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	F_errmsg_internal(m, int32(_a_F_vac_update_relstats_4), v19)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_vac_update_relstats_1), int32(1464), int32(_a_F_vac_update_relstats_2))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
