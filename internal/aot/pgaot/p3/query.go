package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineQueryRewrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
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
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
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
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	v14 = m.G0
	v16 = v14 - int32(240)
	m.G0 = v16
	v19 = F_table_open(m, l2, int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	if l5 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L10
	} else {
		goto L202
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L10
	} else {
		goto L198
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L10
	} else {
		goto L194
	}
L5:
	;
	v628 = int32(555736)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, _consts[779])))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v632 == int32(0) {
		v651 = v631
		v652 = v632
		goto L182
	} else {
		goto L183
	}
L6:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v576 <= int32(0) {
		goto L5
	} else {
		goto L170
	}
L7:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+119)))
	switch v227 - int32(109) {
	case 0, 9:
		goto L70
	default:
		goto L71
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L66
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L61
	}
L10:
	;
	return
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+119)))
	v24 = v22 - int32(109)
	if base.Ui32(int32(9)) < base.Ui32(v24) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if int32(1)<<(uint(v24)%32)&int32(553) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[501])))
	if v34 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if base.Ui32(v39) < base.Ui32(int32(12000)) {
		v48 = v38
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v52 = F_object_ownercheck(m, int32(1259), l2, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L22
	}
L17:
	;
	if v48 != 0 {
		goto L8
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	if v43 == int32(99) {
		v48 = v38
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v46 = F_isTempToastNamespace(m, v43)
	mBase = m.M
	v48 = v46
	goto L18
L21:
	;
	goto L16
L22:
	;
	if v52 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57)+119)))
	switch v58 - int32(73) {
	case 0, 32:
		goto L32
	default:
		v68 = int32(41)
		goto L27
	case 10:
		goto L31
	case 29:
		goto L28
	case 36:
		goto L29
	case 45:
		goto L30
	}
L24:
	;
	goto L25
L25:
	;
	if l7 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	F_aclcheck_error(m, int32(2), v70, v71+int32(4))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L33
	}
L27:
	;
	v70 = v68
	goto L26
L28:
	;
	v68 = int32(18)
	goto L27
L29:
	;
	v70 = int32(23)
	goto L26
L30:
	;
	v70 = int32(51)
	goto L26
L31:
	;
	v70 = int32(37)
	goto L26
L32:
	;
	v70 = int32(20)
	goto L26
L33:
	;
	goto L25
L34:
	;
	v76 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v76 < v77 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if l4 == int32(1) {
		goto L7
	} else {
		goto L60
	}
L37:
	;
	v89 = v76
	goto L40
L38:
	;
	goto L39
L39:
	;
	if l4 != int32(1) {
		goto L6
	} else {
		goto L59
	}
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v89<<(uint(int32(2))%32))))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+32))
	if v98 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v149 = v89 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v149 < v150 {
		v89 = v149
		goto L40
	} else {
		goto L58
	}
L43:
	;
	v102 = F_getInsertSelectQuery(m, v97, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	if v102 != v97 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+32))
	switch v105 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L42
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L53
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(468879), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(683248), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(525170), int32(293), int32(368349))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(468730), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errhint(m, int32(683279), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(525170), int32(298), int32(368349))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
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
	goto L41
L59:
	;
	goto L7
L60:
	;
	goto L5
L61:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v176 + int32(4)
	F_errmsg(m, int32(174315), v16)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v183)+119)))
	F_errdetail_relkind_not_supported(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(525170), int32(263), int32(368349))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v199 + int32(4)
	F_errmsg(m, int32(345099), v16+int32(96))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(525170), int32(269), int32(368349))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	if l7 != 0 {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v237 + int32(4)
	F_errmsg(m, int32(174393), v16+int32(16))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v247 = int32(*(*int8)(unsafe.Add(mBase, uint32(v246)+119)))
	F_errdetail_relkind_not_supported(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(525170), int32(314), int32(368349))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L10
	} else {
		goto L166
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L10
	} else {
		goto L162
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L158
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L10
	} else {
		goto L154
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L10
	} else {
		goto L150
	}
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if int32(2) <= v255 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L10
	} else {
		goto L145
	}
L85:
	;
	if l5 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v262 != int32(1) {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+42)))
	if v265 == int32(1) {
		goto L79
	} else {
		goto L88
	}
L88:
	;
	if l3 != 0 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261)+76))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	F_checkRuleResultList(m, v268, v269, int32(1), base.B2i32(v227 != int32(109)))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	if l6 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v340 = int32(105)
	v341 = int32(555736)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, _consts[779])))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v345 == int32(0) {
		v364 = v344
		v365 = v345
		goto L106
	} else {
		goto L107
	}
L92:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	if v275 == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v278 <= int32(0) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v292 = int32(0)
	goto L95
L95:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v281+v292<<(uint(int32(2))%32))))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if v300 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L10
	} else {
		goto L101
	}
L97:
	;
	v304 = v292 + int32(1)
	if v278 != v304 {
		v292 = v304
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L91
L101:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v313 + int32(4)
	F_errmsg(m, int32(33963), v16-int32(-64))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(525170), int32(385), int32(368349))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	if v365-v364 == int32(0) {
		v726 = l1
		v737 = v340
		goto L1
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	if v344 != v345 {
		v364 = v344
		v365 = v345
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v349 = l1
	v350 = v341
	goto L109
L109:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	if v354 == int32(0) {
		v364 = v353
		v365 = v354
		goto L106
	} else {
		goto L111
	}
L110:
	;
	v364 = v353
	v365 = v354
	goto L106
L111:
	;
	v357 = int32(1)
	if v353 == v354 {
		v349 = v349 + v357
		v350 = v350 + v357
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v369 = int32(548531)
	goto L116
L114:
	;
	if v406-v407 != 0 {
		goto L77
	} else {
		goto L128
	}
L116:
	;
	goto L117
L117:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v376 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v377 = l1
	v378 = v369
	v379 = int32(4)
	v380 = v376
	goto L122
L119:
	;
	v402 = v369
	v406 = int32(0)
	goto L120
L120:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	goto L114
L121:
	;
	v402 = v397
	v406 = v399
	goto L120
L122:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v380 != v382 {
		v397 = v378
		v399 = v380
		goto L121
	} else {
		goto L124
	}
L123:
	;
	v397 = v391
	v399 = int32(0)
	goto L121
L124:
	;
	if v382 == int32(0) {
		v397 = v378
		v399 = v380
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v387 = v379 - int32(1)
	if v387 == int32(0) {
		v397 = v378
		v399 = v380
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v390 = int32(1)
	v391 = v378 + v390
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+1)))
	if v392 != 0 {
		v377 = v377 + v390
		v378 = v391
		v379 = v387
		v380 = v392
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	v415 = int32(4)
	v416 = l1 + v415
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v419 = v417 + v415
	goto L131
L129:
	;
	if v456-v457 != 0 {
		goto L77
	} else {
		goto L143
	}
L131:
	;
	goto L132
L132:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	if v426 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v427 = v416
	v428 = v419
	v429 = int32(56)
	v430 = v426
	goto L137
L134:
	;
	v452 = v419
	v456 = int32(0)
	goto L135
L135:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	goto L129
L136:
	;
	v452 = v447
	v456 = v449
	goto L135
L137:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	if v430 != v432 {
		v447 = v428
		v449 = v430
		goto L136
	} else {
		goto L139
	}
L138:
	;
	v447 = v441
	v449 = int32(0)
	goto L136
L139:
	;
	if v432 == int32(0) {
		v447 = v428
		v449 = v430
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v437 = v429 - int32(1)
	if v437 == int32(0) {
		v447 = v428
		v449 = v430
		goto L136
	} else {
		goto L141
	}
L141:
	;
	v440 = int32(1)
	v441 = v428 + v440
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	if v442 != 0 {
		v427 = v427 + v440
		v428 = v441
		v429 = v437
		v430 = v442
		goto L137
	} else {
		goto L142
	}
L142:
	;
	goto L138
L143:
	;
	v466 = F_pstrdup(m, int32(555736))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L10
	} else {
		goto L144
	}
L144:
	;
	v726 = v466
	v737 = v340
	goto L1
L145:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L10
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(468827), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	F_errhint(m, int32(683198), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(525170), int32(323), int32(368349))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(468770), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(525170), int32(331), int32(368349))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(549492), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L10
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(525170), int32(341), int32(368349))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L10
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(562743), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L10
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(525170), int32(349), int32(368349))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L10
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	F_errmsg(m, int32(548994), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L10
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(525170), int32(357), int32(368349))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = int32(555736)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v560 + int32(4)
	F_errmsg(m, int32(760336), v16+int32(48))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L10
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(525170), int32(410), int32(368349))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	v589 = int32(0)
	v592 = int32(0)
	goto L171
L171:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v593+v589<<(uint(int32(2))%32))))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+96))
	if v598 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L5
L173:
	;
	if v592&int32(1) != 0 {
		goto L4
	} else {
		goto L176
	}
L174:
	;
	v609 = v592
	goto L175
L175:
	;
	v611 = v589 + int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v611 < v612 {
		v589 = v611
		v592 = v609
		goto L171
	} else {
		goto L180
	}
L176:
	;
	if l3 != 0 {
		goto L3
	} else {
		goto L177
	}
L177:
	;
	if l5 == int32(0) {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v604 = int32(0)
	F_checkRuleResultList(m, v598, v603, v604, v604)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L10
	} else {
		goto L179
	}
L179:
	;
	v609 = int32(1)
	goto L175
L180:
	;
	goto L172
L181:
	;
	if v652-v651 != 0 {
		v726 = l1
		v737 = int32(97)
		goto L1
	} else {
		goto L189
	}
L182:
	;
	goto L181
L183:
	;
	if v631 != v632 {
		v651 = v631
		v652 = v632
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v636 = l1
	v637 = v628
	goto L185
L185:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+1)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636)+1)))
	if v641 == int32(0) {
		v651 = v640
		v652 = v641
		goto L182
	} else {
		goto L187
	}
L186:
	;
	v651 = v640
	v652 = v641
	goto L182
L187:
	;
	v644 = int32(1)
	if v640 == v641 {
		v636 = v636 + v644
		v637 = v637 + v644
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = int32(555736)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v661 + int32(4)
	F_errmsg(m, int32(760374), v16+int32(80))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L10
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(525170), int32(460), int32(368349))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L10
	} else {
		goto L195
	}
L195:
	;
	F_errmsg(m, int32(402417), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(525170), int32(435), int32(368349))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	F_errmsg(m, int32(174260), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L10
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(525170), int32(440), int32(368349))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	F_errmsg(m, int32(174524), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(525170), int32(444), int32(368349))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L10
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v944 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v942
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2618)
	F_sequence_close(m, v19, v944)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L10
	} else {
		goto L267
	}
L207:
	;
	v740 = int32(0)
	if l7 == v740 {
		v942 = v740
		goto L206
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v744 = F_nodeToString(m, l3)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L10
	} else {
		goto L211
	}
L210:
	;
	goto L209
L211:
	;
	v746 = F_nodeToString(m, l7)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L10
	} else {
		goto L212
	}
L212:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+200)) = int64(0)
	v753 = F_strncpy(m, v16+int32(136), v726, int32(64))
	mBase = m.M
	v754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v753)+63)) = uint8(v754)
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+228)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v16)+224)) = int32(79)
	v759 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+220)) = (l4<<(uint(v759)%32) + int32(805306368)) >> (uint(v759) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+216)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v16 + int32(136)
	v770 = F_cstring_to_text(m, v744)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L10
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+232)) = v770
	v773 = F_cstring_to_text(m, v746)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L10
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+236)) = v773
	v778 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L10
	} else {
		goto L216
	}
L216:
	;
	v781 = F_SearchSysCache2(m, int32(60), l2, v726)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L10
	} else {
		goto L219
	}
L217:
	;
	v849 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = int32(2618)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = int32(1259)
	F_recordDependencyOn(m, v16+int32(120), v16+int32(108), v737)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L10
	} else {
		goto L238
	}
L218:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v778)+52))
	v829 = F_heap_modify_tuple(m, v781, v822, v16+int32(208), v16+int32(200), v16+int32(120))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L10
	} else {
		goto L233
	}
L219:
	;
	if v781 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = int64(72340168543043584)
	if l6 != 0 {
		goto L218
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v808 = F_GetNewOidWithIndex(m, v778, int32(2692), int32(1))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L10
	} else {
		goto L229
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	v792 = F_get_rel_name(m, l2)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v726
	F_errmsg(m, int32(124704), v16+int32(32))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(525170), int32(105), int32(402713))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+208)) = v808
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v778)+52))
	v816 = F_heap_form_tuple(m, v811, v16+int32(208), v16+int32(200))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	F_CatalogTupleInsert(m, v778, v816)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	F_pfree(m, v816)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	v848 = v808
	goto L217
L233:
	;
	F_CatalogTupleUpdate(m, v778, v829+int32(4), v829)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	F_ReleaseCatCache(m, v781)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v829)+16))
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837)+22)))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v837+v838)))
	F_pfree(m, v829)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L10
	} else {
		goto L236
	}
L236:
	;
	v845 = F_deleteDependencyRecordsFor(m, int32(2618), v840, int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L237
	}
L237:
	;
	v848 = v840
	goto L217
L238:
	;
	F_recordDependencyOnExpr(m, v16+int32(120), l7, int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L10
	} else {
		goto L239
	}
L239:
	;
	if l3 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	v875 = F_getInsertSelectQuery(m, v873, int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	if v881 != 0 {
		goto L245
	} else {
		goto L246
	}
L243:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v875)+52))
	F_recordDependencyOnExpr(m, v16+int32(120), l3, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v883 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2618), v848, v883, v883)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L10
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	F_sequence_close(m, v778, int32(3))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L10
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	v890 = m.G0
	v892 = v890 - int32(16)
	m.G0 = v892
	v896 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	v900 = F_SearchSysCacheCopy(m, int32(57), l2, int32(0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L10
	} else {
		goto L252
	}
L251:
	;
	v942 = v848
	goto L206
L252:
	;
	if v900 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v900)+16))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+22)))
	v904 = v902 + v903
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+124)))
	if v905 != int32(1) {
		goto L257
	} else {
		goto L258
	}
L254:
	;
	goto L255
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L10
	} else {
		goto L264
	}
L256:
	;
	F_pfree(m, v900)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L10
	} else {
		goto L262
	}
L257:
	;
	v908 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v904)+124)) = uint8(v908)
	F_CatalogTupleUpdate(m, v896, v900+int32(4), v900)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L10
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	F_CacheInvalidateRelcacheByTuple(m, v900)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L10
	} else {
		goto L261
	}
L260:
	;
	goto L256
L261:
	;
	goto L256
L262:
	;
	F_sequence_close(m, v896, int32(3))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L10
	} else {
		goto L263
	}
L263:
	;
	m.G0 = v892 + int32(16)
	goto L251
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892))) = l2
	F_errmsg_internal(m, int32(50136), v892)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L10
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(517751), int32(65), int32(122687))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L10
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	m.G0 = v16 + int32(240)
	return
}
func F_assign_query_collations(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v5 = F_query_tree_walker_impl(m, l1, int32(482), l0, int32(10))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_extract_query_dependencies_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v157 == int32(0)) & v155
L2:
	;
	v155 = v3
	v157 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != int32(67) {
		v84 = l0
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+44)))
	if v103 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L6:
	;
	F_fix_expr_common(m, l1, v84)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L13
	} else {
		goto L42
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v97 = l0
	v101 = v3
	goto L5
L9:
	;
	goto L10
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 != int32(213) {
		v50 = v15
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v57 = base.B2i32(v16 == int32(213))
	v59 = v50
	goto L28
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v20 = F_extract_query_dependencies_walker(m, v19, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v24 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v155 = v3
	v157 = int32(1)
	goto L1
L16:
	;
	goto L17
L17:
	;
	v28 = v24
	goto L18
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v34 != int32(67) {
		v84 = v28
		goto L6
	} else {
		goto L20
	}
L19:
	;
	v155 = v3
	v157 = int32(1)
	goto L1
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v37 != int32(6) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = v28
	v101 = int32(1)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != int32(213) {
		v50 = v41
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = F_extract_query_dependencies_walker(m, v45, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v49 != 0 {
		v28 = v49
		goto L18
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	if v83 != 0 {
		v97 = v83
		v101 = v57
		goto L5
	} else {
		goto L41
	}
L28:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch v61 - int32(241) {
	case 0:
		goto L33
	case 1:
		goto L32
	default:
		goto L34
	}
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v59 = v81
	goto L28
L31:
	;
	v83 = v79
	goto L27
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v76 == int32(6) {
		v80 = v75
		goto L30
	} else {
		goto L40
	}
L33:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 == int32(6) {
		v80 = v71
		goto L30
	} else {
		goto L39
	}
L34:
	;
	if v61 != int32(201) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v83 = int32(0)
	goto L27
L36:
	;
	goto L37
L37:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 != int32(6) {
		v79 = v67
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v80 = v67
	goto L30
L39:
	;
	v79 = v71
	goto L31
L40:
	;
	v79 = v75
	goto L31
L41:
	;
	v155 = v3
	v157 = v57
	goto L1
L42:
	;
	v95 = F_expression_tree_walker_impl(m, v84, int32(838), l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v155 = v95
	v157 = base.B2i32(v9 == int32(67))
	goto L1
L44:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+81)) = uint8(v107)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+52))
	if v109 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v151 = F_query_tree_walker_impl(m, v97, int32(838), l1, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L13
	} else {
		goto L59
	}
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v117 = v3
	goto L50
L50:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v117<<(uint(int32(2))%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	switch v126 {
	case 0:
		goto L55
	case 1, 7:
		goto L54
	default:
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	v140 = v117 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v140 < v141 {
		v117 = v140
		goto L50
	} else {
		goto L58
	}
L53:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+56))
	v134 = F_lappend_oid(m, v133, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L57
	}
L54:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v128 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v131 = v127
	goto L53
L56:
	;
	v131 = v128
	goto L53
L57:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+56)) = v134
	goto L52
L58:
	;
	goto L51
L59:
	;
	v155 = v151
	v157 = v101
	goto L1
}
func F_query_to_xml_and_xmlschema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_text_to_cstring(m, v13)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v21 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_text_to_cstring(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(0)
						v30 = F_SPI_prepare(m, v17, v28, v28)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							if v30 != 0 {
								v32 = F_SPI_cursor_open(m, v30)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									if v32 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v17
											F_errmsg_internal(m, int32(477573), v10+int32(16))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523639), int32(3161), int32(533268))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+92))
										v37 = int32(0)
										v39 = base.B2i32(v19 != v37)
										v40 = F_map_sql_table_to_xmlschema(m, v36, v37, v39, v23)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = F_strlen(m, v40)
											mBase = m.M
											v44 = v42 + int32(1)
											v45 = F_SPI_palloc(m, v44)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return int32(0)
											} else {
												if v44 != 0 {
													v47 = F__emscripten_memcpy_bulkmem(m, v45, v40, v44)
													mBase = m.M
													v48 = v47
												} else {
													v48 = v45
												}
												F_SPI_cursor_close(m, v32)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int32(0)
												} else {
													v51 = F_SPI_finish(m)
													mBase = m.M
													v52 = m.ExcPending
													if v52 != 0 {
														return int32(0)
													} else {
														v54 = F_query_to_xml_internal(m, v17, int32(0), v48, v39, v23)
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return int32(0)
														} else {
															v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
															v58 = F_cstring_to_text_with_len(m, v56, v57)
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(32)
																return v58
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
									F_errmsg_internal(m, int32(477602), v10)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523639), int32(3158), int32(533268))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
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
}
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
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
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
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
	var v240 int32
	_ = v240
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
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	v5 = int32(0)
	if l3&int32(64) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v25 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v24, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L11
	}
L2:
	;
	v23 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = F_palloc(m, int32(168))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L8
L7:
	;
	v23 = v16
	goto L1
L8:
	;
	v21 = F__emscripten_memcpy_bulkmem(m, v16, l0, int32(168))
	mBase = m.M
	goto L10
L10:
	;
	goto L7
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+152))
	v29 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v28, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v33 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v37 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v36, l2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	v41 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v40, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v45 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v44, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v48, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)+144))
	v53 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v52, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+112))
	v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
	v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v60, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	v65 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v64, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v65
	if l3&int32(128) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	if l3&int32(2) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	v71 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v70, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+116))
	if v86 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v23)+116))
	v75 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v74, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+124))
	v79 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v78, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+124)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v23)+120))
	v83 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v82, l2)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+120)) = v83
	goto L22
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v150
	goto L22
L31:
	;
	v97 = v5
	v98 = v5
	goto L36
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if int32(0) < v87 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v150 = v5
	goto L30
L35:
	;
	goto L34
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v97<<(uint(int32(2))%32))))
	v108 = F_palloc(m, int32(56))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	v150 = v136
	goto L30
L38:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+48)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v106)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+40)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v106)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+32)) = v114
	v116 = int32(24)
	v117 = v108 + v116
	v119 = v106 + v116
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v120
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v106)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+16)) = v122
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v106)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v106)))
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v129 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v128, l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v106)+28))
	v133 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v132, l2)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+28)) = v133
	v136 = F_lappend(m, v98, v108)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v139 = v97 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v139 < v140 {
		v97 = v139
		v98 = v136
		goto L36
	} else {
		goto L42
	}
L42:
	;
	goto L37
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v176
	v178 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	if v180 != 0 {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	v172 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v167, l2)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v174 = F_copyObjectImpl(m, v167)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L48
	}
L47:
	;
	v176 = v172
	goto L43
L48:
	;
	v176 = v174
	goto L43
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v274
	return v23
L50:
	;
	v196 = v178
	v197 = v178
	goto L55
L51:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if int32(0) < v181 {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v274 = v178
	goto L49
L54:
	;
	goto L53
L55:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v197<<(uint(int32(2))%32))))
	v208 = F_palloc(m, int32(136))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	v274 = v262
	goto L49
L57:
	;
	goto L59
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	switch v213 {
	case 0:
		goto L69
	case 1:
		goto L68
	case 2:
		goto L67
	case 3:
		goto L66
	case 4:
		goto L65
	case 5:
		goto L64
	default:
		goto L62
	case 9:
		goto L63
	}
L59:
	;
	v211 = F__emscripten_memcpy_bulkmem(m, v208, v206, int32(136))
	mBase = m.M
	goto L61
L61:
	;
	goto L58
L62:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v206)+128))
	v259 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v258, l2)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L89
	}
L63:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v206)+120))
	if l3&int32(256) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v206)+80))
	v245 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v244, l2)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L83
	}
L65:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v206)+76))
	v241 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v240, l2)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L82
	}
L66:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v206)+68))
	v237 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v236, l2)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L81
	}
L67:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v206)+52))
	if l3&int32(4) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	if l3&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	v215 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v214, l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+32)) = v215
	goto L62
L71:
	;
	v221 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v218, l2)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v224 = F_copyObjectImpl(m, v218)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+36)) = v221
	goto L62
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+36)) = v224
	goto L62
L76:
	;
	v230 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v227, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v233 = F_copyObjectImpl(m, v227)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+52)) = v230
	goto L62
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+52)) = v233
	goto L62
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+68)) = v237
	goto L62
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+76)) = v241
	goto L62
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+80)) = v245
	goto L62
L84:
	;
	v251 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v248, l2)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v254 = F_copyObjectImpl(m, v248)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+120)) = v251
	goto L62
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+120)) = v254
	goto L62
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+128)) = v259
	v262 = F_lappend(m, v196, v211)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	v265 = v197 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v265 < v266 {
		v196 = v262
		v197 = v265
		goto L55
	} else {
		goto L91
	}
L91:
	;
	goto L56
}
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v11 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v10, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v11 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v16 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v15, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v19 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v18, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v19 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v22 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v21, l2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v22 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v25 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v24, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v25 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v28 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v27, l2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v31 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v30, l2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v31 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v34 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v33, l2)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v34 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v37 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v36, l2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v37 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v40 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v43 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v42, l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v43 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if l3&int32(128) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if l3&int32(2) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v48 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v47, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v59 == int32(0) {
		goto L25
	} else {
		goto L37
	}
L29:
	;
	if v48 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v51 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v50, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v51 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v54 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v53, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v54 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v57 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v62 <= int32(0) {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v70 = int32(0)
	goto L39
L39:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v70<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	v79 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v78, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L1
L41:
	;
	goto L40
L42:
	;
	if v79 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v82 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v81, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v82 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v85 = v70 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v85 < v86 {
		v70 = v85
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L25
L47:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v101 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v100, l2)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if l3&int32(8) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v101 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v108 = F_range_table_walker_impl(m, v107, l1, l2, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	return int32(0)
L55:
	;
	if v108 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
}
