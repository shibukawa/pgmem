package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessConfigFileInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var __phi280 int32
	_ = __phi280
	var v281 int32
	_ = v281
	var __phi281 int32
	_ = __phi281
	var v286 int32
	_ = v286
	var __phi286 int32
	_ = __phi286
	var v287 int32
	_ = v287
	var __phi287 int32
	_ = __phi287
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int64
	_ = v815
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int64
	_ = v901
	var v902 int64
	_ = v902
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v942 int32
	_ = v942
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v984 int32
	_ = v984
	var v996 int32
	_ = v996
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(160)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+152)) = v4
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[0]))
	v32 = F_ParseConfigFile(m, v23, int32(1), v4, v4, v4, l2, v16+int32(156), v16+int32(152))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L246
	}
L2:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	m.G0 = v16 + int32(160)
	return v984
L3:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L239
	}
L4:
	;
	return int32(0)
L5:
	;
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v942 = v23
	goto L3
L7:
	;
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[1]))
	if v39 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v137 = v16 + int32(132)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[2]))
	F_hash_seq_init(m, v137, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L36
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v109 = int32(_a_F_ProcessConfigFileInternal_0)
	v111 = int32(0)
	v119 = F_ParseConfigFile(m, v109, v111, v111, v111, v111, l2, v16+int32(156), v16+int32(152))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L34
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v16)+152)) = v88
	goto L9
L14:
	;
	v47 = v42
	v49 = v4
	goto L17
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = int32(0)
	goto L2
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
	if v56 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v88 != 0 {
		goto L13
	} else {
		goto L33
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v60 = int32(_a_F_ProcessConfigFileInternal_1)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[3])))
	if base.B2i32(v63 == int32(0))|base.B2i32(v63 != v66) != 0 {
		v84 = v63
		v85 = v66
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v88 = v49
	goto L21
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v89 != 0 {
		v47 = v89
		v49 = v88
		goto L17
	} else {
		goto L32
	}
L22:
	;
	if v84-v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v69 = v59
	v70 = v60
	goto L25
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v74 == int32(0) {
		v84 = v74
		v85 = v73
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v84 = v74
	v85 = v73
	goto L23
L27:
	;
	v77 = int32(1)
	if v74 == v73 {
		v69 = v69 + v77
		v70 = v70 + v77
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v87 = v49
	goto L31
L30:
	;
	v87 = v47
	goto L31
L31:
	;
	v88 = v87
	goto L21
L32:
	;
	goto L18
L33:
	;
	goto L16
L34:
	;
	if v119 == int32(0) {
		v942 = v109
		goto L3
	} else {
		goto L35
	}
L35:
	;
	goto L9
L36:
	;
	v142 = F_hash_seq_search(m, v137)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v142 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v148 = v142
	goto L41
L39:
	;
	goto L40
L40:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	if v179 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+28)) = v158 & int32(-2)
	v164 = F_hash_seq_search(m, v16+int32(132))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	if v164 != 0 {
		v148 = v164
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v497 = int32(0)
	v499 = v16 + int32(132)
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[2]))
	F_hash_seq_init(m, v499, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L121
	}
L46:
	;
	v491 = v23
	goto L45
L47:
	;
	goto L48
L48:
	;
	v188 = v179
	v189 = v23
	v194 = v4
	goto L49
L49:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+20)))
	if v195 != 0 {
		v477 = v189
		v482 = v194
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v482 != 0 {
		v942 = v477
		goto L3
	} else {
		goto L120
	}
L51:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
	if v483 != 0 {
		v188 = v483
		v189 = v477
		v194 = v482
		goto L49
	} else {
		goto L119
	}
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v199 = F_find_option(m, v196, int32(0), int32(1), l2)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v199 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199)+28))
	if v201&int32(1) == int32(0) {
		v258 = v201
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v272 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+28)) = v258 | int32(1)
	v477 = v189
	v482 = v194
	goto L51
L58:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	if v206 == v188 {
		v258 = v201
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v212 = v206
	goto L60
L60:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+20)))
	if v221 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v199)+28))
	v258 = v254
	goto L57
L62:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v212)+24))
	if v252 != v188 {
		v212 = v252
		goto L60
	} else {
		goto L72
	}
L63:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if base.B2i32(v226 == int32(0))|base.B2i32(v226 != v229) != 0 {
		v247 = v226
		v248 = v229
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v247-v248 != 0 {
		goto L62
	} else {
		goto L71
	}
L65:
	;
	goto L64
L66:
	;
	v232 = v222
	v233 = v223
	goto L67
L67:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	if v237 == int32(0) {
		v247 = v237
		v248 = v236
		goto L65
	} else {
		goto L69
	}
L68:
	;
	v247 = v237
	v248 = v236
	goto L65
L69:
	;
	v240 = int32(1)
	if v237 == v236 {
		v232 = v232 + v240
		v233 = v233 + v240
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+20)) = uint8(v250)
	goto L62
L72:
	;
	goto L61
L73:
	;
	v441 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L111
	}
L74:
	;
	__phi280 = v271
	__phi281 = v272
	__phi286 = int32(0)
	__phi287 = int32(1)
	v280 = __phi280
	v281 = __phi281
	v286 = __phi286
	v287 = __phi287
	goto L75
L75:
	;
	v291 = base.B2i32(v281 != int32(46))
	if v291 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	if v420&v291 != 0 {
		v477 = v189
		v482 = v194
		goto L51
	} else {
		goto L110
	}
L77:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	if v425 != 0 {
		__phi280 = v280 + int32(1)
		__phi281 = v425
		__phi286 = v420
		__phi287 = base.B2i32(v281 == int32(46))
		v280 = __phi280
		v281 = __phi281
		v286 = __phi286
		v287 = __phi287
		goto L75
	} else {
		goto L109
	}
L78:
	;
	if v287 == int32(0) {
		v420 = int32(1)
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v298 = base.I32_extend8_s(v281)
	goto L86
L81:
	;
	goto L73
L82:
	;
	if v404|base.B2i32(v298 < int32(0)) != 0 {
		v420 = v286
		goto L77
	} else {
		goto L107
	}
L83:
	;
	v404 = int32(0)
	goto L82
L84:
	;
	v382 = v375
	v384 = v377
	goto L101
L85:
	;
	if base.B2i32(v321 != v322) == int32(0) {
		goto L83
	} else {
		goto L92
	}
L86:
	;
	v313 = int32(_a_F_ProcessConfigFileInternal_2)
	v315 = int32(54)
	goto L87
L87:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if v318 == v298&int32(255) {
		v375 = v313
		v377 = v315
		goto L84
	} else {
		goto L89
	}
L88:
	;
	goto L85
L89:
	;
	v320 = int32(1)
	v321 = v315 - v320
	v322 = int32(0)
	v325 = v313 + v320
	if v325&int32(3) == v322 {
		goto L85
	} else {
		goto L90
	}
L90:
	;
	if v321 != 0 {
		v313 = v325
		v315 = v321
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v338 = v298 & int32(255)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if base.B2i32(v338 == v339)|base.B2i32(base.Ui32(v321) < base.Ui32(int32(4))) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v348 = v325
	v350 = v321
	goto L96
L94:
	;
	v368 = v325
	v370 = v321
	goto L95
L95:
	;
	if v370 == int32(0) {
		goto L83
	} else {
		goto L100
	}
L96:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v355 = v354 ^ v338*int32(16843009)
	v358 = int32(-2139062144)
	if (int32(16843008)-v355|v355)&v358 != v358 {
		v375 = v348
		v377 = v350
		goto L84
	} else {
		goto L98
	}
L97:
	;
	v368 = v363
	v370 = v365
	goto L95
L98:
	;
	v362 = int32(4)
	v363 = v348 + v362
	v365 = v350 - v362
	if base.Ui32(int32(3)) < base.Ui32(v365) {
		v348 = v363
		v350 = v365
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v375 = v368
	v377 = v370
	goto L84
L101:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v298&int32(255) == v387 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L83
L103:
	;
	v404 = v382
	goto L82
L104:
	;
	goto L105
L105:
	;
	v389 = int32(1)
	v392 = v384 - v389
	if v392 != 0 {
		v382 = v382 + v389
		v384 = v392
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v298))%64)&int64(287948969894477825) == int64(0))|(v287|base.B2i32(base.Ui32(int32(63)) < base.Ui32(v281))) != 0 {
		goto L73
	} else {
		goto L108
	}
L108:
	;
	v420 = v286
	goto L77
L109:
	;
	goto L76
L110:
	;
	goto L73
L111:
	;
	if v441 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v465 = F_pstrdup(m, int32(_a_F_ProcessConfigFileInternal_3))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L118
	}
L115:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v446
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_4), v16+int32(112))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(430), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v465
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v477 = v468
	v482 = int32(1)
	goto L51
L119:
	;
	goto L50
L120:
	;
	v491 = v477
	goto L45
L121:
	;
	v504 = F_hash_seq_search(m, v499)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	if v504 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v510 = v504
	v515 = v497
	goto L126
L124:
	;
	v692 = v497
	goto L125
L125:
	;
	v698 = base.B2i32(l0 == int32(2)) & l1
	if v698 != 0 {
		goto L169
	} else {
		goto L170
	}
L126:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+36))
	if v520 != int32(3) {
		v675 = v515
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v692 = v675
	goto L125
L128:
	;
	v681 = F_hash_seq_search(m, v16+int32(132))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L167
	}
L129:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v519)+28))
	if v523&int32(1) != 0 {
		v675 = v515
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	if base.Ui32(v526) <= base.Ui32(int32(1)) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+28)) = v523 | int32(2)
	v533 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	if l1 == int32(0) {
		v675 = v515
		goto L128
	} else {
		goto L148
	}
L134:
	;
	if v533 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v550
	v555 = F_psprintf(m, int32(_a_F_ProcessConfigFileInternal_7), v16+int32(48))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L141
	}
L138:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v538
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_7), v16-int32(-64))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(469), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v558 = v16 + int32(156)
	v560 = v16 + int32(152)
	v562 = F_palloc(m, int32(28))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v562))) = int64(0)
	v566 = F_pstrdup(m, v555)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562)+8)) = v566
	v569 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+24)) = v569
	v571 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v562)+20)) = uint16(v571)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+16)) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v562)+12)) = v569
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if v577 == v569 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v562
	v675 = int32(1)
	goto L128
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = v562
	goto L144
L146:
	;
	goto L147
L147:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+24)) = v562
	goto L144
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+36)) = int32(0)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v519)+32))
	if v589 == int32(3) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v519)+64))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v519)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v592)+4)) = v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v519)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v595
	*(*int32)(unsafe.Add(mBase, uint32(v519)+32)) = int32(0)
	goto L151
L150:
	;
	goto L151
L151:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v519)+56))
	if v600 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v605 = v600
	goto L155
L153:
	;
	goto L154
L154:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v636 = int32(0)
	v644 = F_set_config_with_handle(m, v635, v636, v636, l0, v636, int32(10), v636, int32(1), v636, v636)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L161
	}
L155:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	if v614 == int32(3) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L154
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v605)+12)) = int32(0)
	goto L159
L158:
	;
	goto L159
L159:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	if v619 != 0 {
		v605 = v619
		goto L155
	} else {
		goto L160
	}
L160:
	;
	goto L156
L161:
	;
	if base.B2i32(l0 != int32(2))|base.B2i32(v644 <= int32(0)) != 0 {
		v675 = v515
		goto L128
	} else {
		goto L162
	}
L162:
	;
	v650 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	if v650 == int32(0) {
		v675 = v515
		goto L128
	} else {
		goto L164
	}
L164:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v654
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_8), v16+int32(80))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(505), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v675 = v515
	goto L128
L167:
	;
	if v681 != 0 {
		v510 = v681
		v515 = v675
		goto L126
	} else {
		goto L168
	}
L168:
	;
	goto L127
L169:
	;
	F_InitializeGUCOptionsFromEnvironment(m)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L4
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	if v727 != 0 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	v702 = int32(0)
	v704 = int32(1)
	v711 = F_set_config_with_handle(m, int32(_a_F_ProcessConfigFileInternal_9), v702, int32(_a_F_ProcessConfigFileInternal_10), v704, v704, int32(10), v702, v704, v702, v702)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[4]))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	goto L174
L174:
	;
	v719 = int32(1)
	v721 = int32(0)
	v725 = F_set_config_with_handle(m, int32(_a_F_ProcessConfigFileInternal_11), int32(0), v717, int32(4), v719, int32(10), v721, v719, v721, v721)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	goto L171
L176:
	;
	v734 = v727
	v737 = v491
	v739 = v692
	goto L179
L177:
	;
	v884 = v491
	v886 = v692
	goto L178
L178:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L230
	}
L179:
	;
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+20)))
	if v743 != 0 {
		v871 = v737
		v872 = v739
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v884 = v871
	v886 = v872
	goto L178
L181:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v734)+24))
	if v876 != 0 {
		v734 = v876
		v737 = v871
		v739 = v872
		goto L179
	} else {
		goto L229
	}
L182:
	;
	v744 = int32(0)
	v746 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[5])))
	if (base.B2i32(v698 == int32(0))|v746)&int32(1) == v744 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v754 = F_GetConfigOption(m, v752, int32(1))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	v761 = v744
	goto L185
L185:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v763 = int32(0)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v734)+4))
	v770 = F_set_config_with_handle(m, v762, v763, v764, l0, int32(3), int32(10), v763, l1, v763, v763)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L193
	}
L186:
	;
	if v754 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v757 = v754
	goto L189
L188:
	;
	v757 = int32(_a_F_ProcessConfigFileInternal_12)
	goto L189
L189:
	;
	v758 = F_pstrdup(m, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v761 = v758
	goto L185
L191:
	;
	if v761 == int32(0) {
		v871 = v860
		v872 = v861
		goto L181
	} else {
		goto L227
	}
L192:
	;
	v834 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v734)+21)) = uint8(v834)
	if l1 == int32(0) {
		v860 = v737
		v861 = v739
		goto L191
	} else {
		goto L216
	}
L193:
	;
	if int32(0) < v770 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if v761 == int32(0) {
		goto L192
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if v770 != 0 {
		goto L192
	} else {
		goto L214
	}
L197:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v778 = F_GetConfigOption(m, v776, int32(1))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	if v778 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v781 = v778
	goto L201
L200:
	;
	v781 = int32(_a_F_ProcessConfigFileInternal_12)
	goto L201
L201:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761))))
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	if base.B2i32(v784 == int32(0))|base.B2i32(v784 != v787) != 0 {
		v805 = v784
		v806 = v787
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v805-v806 == int32(0) {
		goto L192
	} else {
		goto L209
	}
L203:
	;
	goto L202
L204:
	;
	v790 = v761
	v791 = v781
	goto L205
L205:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791)+1)))
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790)+1)))
	if v795 == int32(0) {
		v805 = v795
		v806 = v794
		goto L203
	} else {
		goto L207
	}
L206:
	;
	v805 = v795
	v806 = v794
	goto L203
L207:
	;
	v798 = int32(1)
	if v795 == v794 {
		v790 = v790 + v798
		v791 = v791 + v798
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v811 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	if v811 == int32(0) {
		goto L192
	} else {
		goto L211
	}
L211:
	;
	v815 = *(*int64)(unsafe.Add(mBase, uint32(v734)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v815
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_13), v16+int32(32))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(570), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	goto L192
L214:
	;
	v828 = F_pstrdup(m, int32(_a_F_ProcessConfigFileInternal_14))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v734)+8)) = v828
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v734)+12))
	v860 = v831
	v861 = int32(1)
	goto L191
L216:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v734)+16))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v734)+12))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[5])))
	if v846 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v847 = int32(12)
	goto L219
L218:
	;
	v847 = int32(15)
	goto L219
L219:
	;
	v848 = F_find_option(m, v840, int32(1), int32(0), v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	if v848 == int32(0) {
		v860 = v737
		v861 = v739
		goto L191
	} else {
		goto L221
	}
L221:
	;
	v852 = F_guc_strdup(m, v847, v839)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L4
	} else {
		goto L222
	}
L222:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v848)+84))
	if v854 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_pfree(m, v854)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v848)+88)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v848)+84)) = v852
	v860 = v737
	v861 = v739
	goto L191
L226:
	;
	goto L225
L227:
	;
	F_pfree(m, v761)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v871 = v860
	v872 = v861
	goto L181
L229:
	;
	goto L180
L230:
	;
	v896 = m.G0
	v897 = int32(16)
	v898 = v896 - v897
	m.G0 = v898
	F_gettimeofday(m, v898)
	mBase = m.M
	v901 = *(*int64)(unsafe.Add(mBase, uint32(v898)))
	v902 = int64(*(*int32)(unsafe.Add(mBase, uint32(v898)+8)))
	m.G0 = v898 + v897
	goto L231
L231:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ProcessConfigFileInternal[6])) = v902 + v901*int64(1000000) - int64(946684800000000)
	if v886 == int32(0) {
		goto L2
	} else {
		goto L232
	}
L232:
	;
	if l0 == int32(1) {
		v996 = v884
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v917 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	if v917 == int32(0) {
		goto L2
	} else {
		goto L235
	}
L235:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v884
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_15), v16+int32(16))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(617), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	goto L2
L239:
	;
	if l0 == int32(1) {
		v996 = v942
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v953 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	if v953 == int32(0) {
		goto L2
	} else {
		goto L242
	}
L242:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v942
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_16), v16+int32(96))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(622), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	goto L2
L246:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v996
	F_errmsg(m, int32(_a_F_ProcessConfigFileInternal_17), v16)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L4
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_ProcessConfigFileInternal_5), int32(612), int32(_a_F_ProcessConfigFileInternal_6))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_show_config_by_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_text_to_cstring(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		v9 = F_GetConfigOptionByName(m, v3, v7, v7)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_cstring_to_text(m, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
