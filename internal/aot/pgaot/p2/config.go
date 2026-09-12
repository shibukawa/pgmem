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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
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
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int64
	_ = v832
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int64
	_ = v919
	var v920 int64
	_ = v920
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v4
	v23 = int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[1230]))
	v34 = F_ParseConfigFile(m, v25, v23, v4, v4, v4, l2, v17+int32(156), v17+int32(152))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L255
	}
L2:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	m.G0 = v17 + int32(160)
	return v1002
L3:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L240
	}
L4:
	;
	return int32(0)
L5:
	;
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v934 = v4
	v938 = v23
	v939 = v25
	goto L3
L7:
	;
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
	F_hash_seq_init(m, v17+int32(132), v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L37
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v112 = int32(346142)
	v114 = int32(0)
	v122 = F_ParseConfigFile(m, v112, v114, v114, v114, v114, l2, v17+int32(156), v17+int32(152))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L35
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v90
	goto L9
L14:
	;
	v48 = v44
	v49 = v4
	goto L17
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = int32(0)
	goto L2
L17:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)))
	if v59 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v90 != 0 {
		goto L13
	} else {
		goto L34
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v63 = int32(13315)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1232])))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v67 == int32(0) {
		v86 = v66
		v87 = v67
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v90 = v49
	goto L21
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v91 != 0 {
		v48 = v91
		v49 = v90
		goto L17
	} else {
		goto L33
	}
L22:
	;
	if v87-v86 != 0 {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	goto L22
L24:
	;
	if v66 != v67 {
		v86 = v66
		v87 = v67
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v71 = v62
	v72 = v63
	goto L26
L26:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v75
		v87 = v76
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v86 = v75
	v87 = v76
	goto L23
L28:
	;
	v79 = int32(1)
	if v75 == v76 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v89 = v49
	goto L32
L31:
	;
	v89 = v48
	goto L32
L32:
	;
	v90 = v89
	goto L21
L33:
	;
	goto L18
L34:
	;
	goto L16
L35:
	;
	if v122 == int32(0) {
		v934 = v4
		v938 = v23
		v939 = v112
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L9
L37:
	;
	v148 = F_hash_seq_search(m, v17+int32(132))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v148 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v153 = v148
	goto L42
L40:
	;
	goto L41
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	if v187 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = v165 & int32(-2)
	v171 = F_hash_seq_search(m, v17+int32(132))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	if v171 != 0 {
		v153 = v171
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v511 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
	F_hash_seq_init(m, v17+int32(132), v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L127
	}
L47:
	;
	v505 = v25
	goto L46
L48:
	;
	goto L49
L49:
	;
	v194 = v187
	v198 = v25
	v203 = v4
	goto L50
L50:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+20)))
	if v204 != 0 {
		v488 = v198
		v493 = v203
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v493 != 0 {
		v934 = int32(0)
		v938 = int32(1)
		v939 = v488
		goto L3
	} else {
		goto L126
	}
L52:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	if v494 != 0 {
		v194 = v494
		v198 = v488
		v203 = v493
		goto L50
	} else {
		goto L125
	}
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v208 = F_find_option(m, v205, int32(0), int32(1), l2)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v208 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)+28))
	if v210&int32(1) == int32(0) {
		v270 = v210
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v282 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+28)) = v270 | int32(1)
	v488 = v198
	v493 = v203
	goto L52
L59:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	if v215 == v194 {
		v270 = v210
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v220 = v215
	goto L61
L61:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+20)))
	if v231 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v208)+28))
	v270 = v263
	goto L58
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v220)+24))
	if v261 != v194 {
		v220 = v261
		goto L61
	} else {
		goto L74
	}
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v237 == int32(0) {
		v256 = v236
		v257 = v237
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v257-v256 != 0 {
		goto L63
	} else {
		goto L73
	}
L66:
	;
	goto L65
L67:
	;
	if v236 != v237 {
		v256 = v236
		v257 = v237
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v241 = v232
	v242 = v233
	goto L69
L69:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	if v246 == int32(0) {
		v256 = v245
		v257 = v246
		goto L66
	} else {
		goto L71
	}
L70:
	;
	v256 = v245
	v257 = v246
	goto L66
L71:
	;
	v249 = int32(1)
	if v245 == v246 {
		v241 = v241 + v249
		v242 = v242 + v249
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+20)) = uint8(v259)
	goto L63
L74:
	;
	goto L62
L75:
	;
	v451 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L117
	}
L76:
	;
	v290 = v282
	v293 = v281
	v296 = int32(0)
	v297 = int32(1)
	goto L77
L77:
	;
	v302 = v290 & int32(255)
	v304 = base.B2i32(v302 != int32(46))
	if v304 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v428&v304 != 0 {
		v488 = v198
		v493 = v203
		goto L52
	} else {
		goto L116
	}
L79:
	;
	v433 = v293 + int32(1)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	if v434 != 0 {
		v290 = v434
		v293 = v433
		v296 = v428
		v297 = base.B2i32(v302 == int32(46))
		goto L77
	} else {
		goto L115
	}
L80:
	;
	if v297 == int32(0) {
		v428 = int32(1)
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v311 = base.I32_extend8_s(v290)
	goto L88
L83:
	;
	goto L75
L84:
	;
	if v311 < int32(0) {
		v428 = v296
		goto L79
	} else {
		goto L110
	}
L85:
	;
	v415 = int32(0)
	goto L84
L86:
	;
	v393 = v386
	v395 = v388
	goto L104
L87:
	;
	if base.B2i32(v333 != v334) == int32(0) {
		goto L85
	} else {
		goto L95
	}
L88:
	;
	goto L89
L89:
	;
	v325 = int32(520191)
	v327 = int32(54)
	goto L90
L90:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if v330 == v311&int32(255) {
		v386 = v325
		v388 = v327
		goto L86
	} else {
		goto L92
	}
L91:
	;
	goto L87
L92:
	;
	v332 = int32(1)
	v333 = v327 - v332
	v334 = int32(0)
	v337 = v325 + v332
	if v337&int32(3) == v334 {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	if v333 != 0 {
		v325 = v337
		v327 = v333
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	if v349 == v311&int32(255) {
		v379 = v337
		v381 = v333
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v381 == int32(0) {
		goto L85
	} else {
		goto L103
	}
L97:
	;
	if base.Ui32(v333) < base.Ui32(int32(4)) {
		v379 = v337
		v381 = v333
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v359 = v337
	v361 = v333
	goto L99
L99:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v366 = v365 ^ v311&int32(255)*int32(16843009)
	v369 = int32(-2139062144)
	if (int32(16843008)-v366|v366)&v369 != v369 {
		v386 = v359
		v388 = v361
		goto L86
	} else {
		goto L101
	}
L100:
	;
	v379 = v374
	v381 = v376
	goto L96
L101:
	;
	v373 = int32(4)
	v374 = v359 + v373
	v376 = v361 - v373
	if base.Ui32(int32(3)) < base.Ui32(v376) {
		v359 = v374
		v361 = v376
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v386 = v379
	v388 = v381
	goto L86
L104:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	if v311&int32(255) == v398 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L85
L106:
	;
	v415 = v393
	goto L84
L107:
	;
	goto L108
L108:
	;
	v400 = int32(1)
	v403 = v395 - v400
	if v403 != 0 {
		v393 = v393 + v400
		v395 = v403
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	if v415 != 0 {
		v428 = v296
		goto L79
	} else {
		goto L111
	}
L111:
	;
	if v297 != 0 {
		goto L75
	} else {
		goto L112
	}
L112:
	;
	if base.Ui32(int32(63)) < base.Ui32(v302) {
		goto L75
	} else {
		goto L113
	}
L113:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v311))%64)&int64(287948969894477825) == int64(0) {
		goto L75
	} else {
		goto L114
	}
L114:
	;
	v428 = v296
	goto L79
L115:
	;
	goto L78
L116:
	;
	goto L75
L117:
	;
	if v451 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v475 = F_pstrdup(m, int32(221124))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L124
	}
L121:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v456
	F_errmsg(m, int32(487487), v17+int32(112))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(512044), int32(430), int32(319427))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	goto L120
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = v475
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v488 = v478
	v493 = int32(1)
	goto L52
L125:
	;
	goto L51
L126:
	;
	v505 = v488
	goto L46
L127:
	;
	v520 = F_hash_seq_search(m, v17+int32(132))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	if v520 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v525 = v520
	v529 = v511
	goto L132
L130:
	;
	v710 = v511
	goto L131
L131:
	;
	v719 = base.B2i32(l0 == int32(2)) & l1
	if v719 != 0 {
		goto L176
	} else {
		goto L177
	}
L132:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+36))
	if v537 != int32(3) {
		v692 = v529
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v710 = v692
	goto L131
L134:
	;
	v701 = F_hash_seq_search(m, v17+int32(132))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L174
	}
L135:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536)+28))
	if v540&int32(1) != 0 {
		v692 = v529
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if base.Ui32(v543) <= base.Ui32(int32(1)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+28)) = v540 | int32(2)
	v550 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if l1 == int32(0) {
		v692 = v529
		goto L134
	} else {
		goto L154
	}
L140:
	;
	if v550 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v567
	v572 = F_psprintf(m, int32(219162), v17-int32(-64))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L4
	} else {
		goto L147
	}
L144:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v555
	F_errmsg(m, int32(219162), v17+int32(80))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(512044), int32(469), int32(319427))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v575 = v17 + int32(156)
	v577 = v17 + int32(152)
	v579 = F_palloc(m, int32(28))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v579))) = int64(0)
	v583 = F_pstrdup(m, v572)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v579)+8)) = v583
	v586 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v579)+24)) = v586
	v588 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v579)+20)) = uint16(v588)
	*(*int32)(unsafe.Add(mBase, uint32(v579)+16)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v579)+12)) = v586
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	if v594 == v586 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v692 = int32(1)
	goto L134
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = v579
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v579
	goto L150
L152:
	;
	goto L153
L153:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	*(*int32)(unsafe.Add(mBase, uint32(v599)+24)) = v579
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v579
	goto L150
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+36)) = int32(0)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v536)+32))
	if v607 == int32(3) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v536)+64))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v536)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v610)+4)) = v611
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v536)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v536)+32)) = int32(0)
	goto L157
L156:
	;
	goto L157
L157:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v536)+56))
	if v618 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v622 = v618
	goto L161
L159:
	;
	goto L160
L160:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v654 = int32(0)
	v662 = F_set_config_with_handle(m, v653, v654, v654, l0, v654, int32(10), v654, int32(1), v654, v654)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L167
	}
L161:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	if v633 == int32(3) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L160
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v622)+12)) = int32(0)
	goto L165
L164:
	;
	goto L165
L165:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	if v638 != 0 {
		v622 = v638
		goto L161
	} else {
		goto L166
	}
L166:
	;
	goto L162
L167:
	;
	if l0 != int32(2) {
		v692 = v529
		goto L134
	} else {
		goto L168
	}
L168:
	;
	if v662 <= int32(0) {
		v692 = v529
		goto L134
	} else {
		goto L169
	}
L169:
	;
	v669 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	if v669 == int32(0) {
		v692 = v529
		goto L134
	} else {
		goto L171
	}
L171:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v673
	F_errmsg(m, int32(100346), v17+int32(96))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(512044), int32(505), int32(319427))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v692 = v529
	goto L134
L174:
	;
	if v701 != 0 {
		v525 = v701
		v529 = v692
		goto L132
	} else {
		goto L175
	}
L175:
	;
	goto L133
L176:
	;
	F_InitializeGUCOptionsFromEnvironment(m)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	if v748 != 0 {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v723 = int32(0)
	v725 = int32(1)
	v732 = F_set_config_with_handle(m, int32(146489), v723, int32(100489), v725, v725, int32(10), v723, v725, v723, v723)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)))
	goto L181
L181:
	;
	v740 = int32(1)
	v742 = int32(0)
	v746 = F_set_config_with_handle(m, int32(343200), int32(0), v738, int32(4), v740, int32(10), v742, v740, v742, v742)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	goto L178
L183:
	;
	v752 = v748
	v756 = v710
	v757 = v505
	goto L186
L184:
	;
	v901 = v710
	v902 = v505
	goto L185
L185:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L238
	}
L186:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+20)))
	if v763 != 0 {
		v888 = v756
		v889 = v757
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v901 = v888
	v902 = v889
	goto L185
L188:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v752)+24))
	if v893 != 0 {
		v752 = v893
		v756 = v888
		v757 = v889
		goto L186
	} else {
		goto L237
	}
L189:
	;
	v764 = int32(0)
	if v719 == v764 {
		v779 = v764
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v752)))
	v781 = int32(0)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	v788 = F_set_config_with_handle(m, v780, v781, v782, l0, int32(3), int32(10), v781, l1, v781, v781)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L4
	} else {
		goto L200
	}
L191:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
	if v769 != 0 {
		v779 = int32(0)
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v752)))
	v772 = F_GetConfigOption(m, v770, int32(1))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	if v772 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v775 = v772
	goto L196
L195:
	;
	v775 = int32(771673)
	goto L196
L196:
	;
	v776 = F_pstrdup(m, v775)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	v779 = v776
	goto L190
L198:
	;
	if v779 == int32(0) {
		v888 = v877
		v889 = v878
		goto L188
	} else {
		goto L235
	}
L199:
	;
	v851 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v752)+21)) = uint8(v851)
	if l1 == int32(0) {
		v877 = v756
		v878 = v757
		goto L198
	} else {
		goto L224
	}
L200:
	;
	if int32(0) < v788 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	if v779 == int32(0) {
		goto L199
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if v788 != 0 {
		goto L199
	} else {
		goto L222
	}
L204:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v752)))
	v796 = F_GetConfigOption(m, v794, int32(1))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	if v796 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v799 = v796
	goto L208
L207:
	;
	v799 = int32(771673)
	goto L208
L208:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799))))
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	if v803 == int32(0) {
		v822 = v802
		v823 = v803
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if v823-v822 == int32(0) {
		goto L199
	} else {
		goto L217
	}
L210:
	;
	goto L209
L211:
	;
	if v802 != v803 {
		v822 = v802
		v823 = v803
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v807 = v779
	v808 = v799
	goto L213
L213:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808)+1)))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807)+1)))
	if v812 == int32(0) {
		v822 = v811
		v823 = v812
		goto L210
	} else {
		goto L215
	}
L214:
	;
	v822 = v811
	v823 = v812
	goto L210
L215:
	;
	v815 = int32(1)
	if v811 == v812 {
		v807 = v807 + v815
		v808 = v808 + v815
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v828 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	if v828 == int32(0) {
		goto L199
	} else {
		goto L219
	}
L219:
	;
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v752)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v832
	F_errmsg(m, int32(716257), v17+int32(48))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(512044), int32(570), int32(319427))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L221
	}
L221:
	;
	goto L199
L222:
	;
	v845 = F_pstrdup(m, int32(466469))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v752)+8)) = v845
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v877 = int32(1)
	v878 = v848
	goto L198
L224:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v752)+16))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v752)))
	v863 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
	if v863 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v864 = int32(12)
	goto L227
L226:
	;
	v864 = int32(15)
	goto L227
L227:
	;
	v865 = F_find_option(m, v857, int32(1), int32(0), v864)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	if v865 == int32(0) {
		v877 = v756
		v878 = v757
		goto L198
	} else {
		goto L229
	}
L229:
	;
	v869 = F_guc_strdup(m, v864, v856)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L230
	}
L230:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v865)+84))
	if v871 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	F_pfree(m, v871)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L4
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865)+88)) = v855
	*(*int32)(unsafe.Add(mBase, uint32(v865)+84)) = v869
	v877 = v756
	v878 = v757
	goto L198
L234:
	;
	goto L233
L235:
	;
	F_pfree(m, v779)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	v888 = v877
	v889 = v878
	goto L188
L237:
	;
	goto L187
L238:
	;
	v914 = m.G0
	v915 = int32(16)
	v916 = v914 - v915
	m.G0 = v916
	F___gettimeofday(m, v916)
	mBase = m.M
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v916)))
	v920 = int64(*(*int32)(unsafe.Add(mBase, uint32(v916)+8)))
	m.G0 = v916 + v915
	goto L239
L239:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1233])) = v920 + v919*int64(1000000) - int64(946684800000000)
	v934 = int32(1)
	v938 = v901
	v939 = v902
	goto L3
L240:
	;
	if v938&int32(1) == int32(0) {
		goto L2
	} else {
		goto L241
	}
L241:
	;
	if l0 == int32(1) {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v954 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	if v934 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	if v954 == int32(0) {
		goto L2
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	if v954 == int32(0) {
		goto L2
	} else {
		goto L251
	}
L247:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L4
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v939
	F_errmsg(m, int32(466396), v17+int32(16))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(512044), int32(617), int32(319427))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	goto L2
L251:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v939
	F_errmsg(m, int32(466331), v17+int32(32))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(512044), int32(622), int32(319427))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	goto L2
L255:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v939
	F_errmsg(m, int32(134512), v17)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(512044), int32(612), int32(319427))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
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
