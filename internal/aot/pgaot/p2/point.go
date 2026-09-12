package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CreateCheckPoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
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
	var v197 int32
	_ = v197
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
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
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v321 int64
	_ = v321
	var v336 int64
	_ = v336
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int64
	_ = v355
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v568 int32
	_ = v568
	var v569 int64
	_ = v569
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int64
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int64
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int64
	_ = v661
	var v663 int64
	_ = v663
	var v666 int32
	_ = v666
	var v667 int64
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int64
	_ = v680
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int64
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v872 int32
	_ = v872
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v958 int32
	_ = v958
	var v973 int64
	_ = v973
	var v978 float64
	_ = v978
	var v980 int32
	_ = v980
	var v982 float64
	_ = v982
	var v989 float64
	_ = v989
	var v994 int64
	_ = v994
	var v995 int64
	_ = v995
	var v1000 int32
	_ = v1000
	var v1002 int64
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int64
	_ = v1008
	var v1010 int64
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1017 int64
	_ = v1017
	var v1018 int64
	_ = v1018
	var v1022 int64
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int64
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1042 float64
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1054 int64
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	v17 = m.G0
	v19 = v17 - int32(1152)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L11
	} else {
		goto L231
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L11
	} else {
		goto L228
	}
L3:
	;
	v43 = int32(0)
	v48 = F__emscripten_memset_bulkmem(m, int32(4375976), base.I32_extend8_s(v43), int32(80))
	mBase = m.M
	goto L9
L4:
	;
	v42 = base.B2i32(l0&int32(2) == int32(0))
	goto L3
L5:
	;
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+316))
	v33 = int32(2)
	v34 = base.B2i32(v32 != v33)
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v34)
	v37 = l0 & v33
	v39 = base.B2i32(v37 == int32(0))
	if v37 != 0 {
		v42 = v39
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v32 != v33 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v42 = v39
	goto L3
L9:
	;
	v53 = m.G0
	v54 = int32(16)
	v55 = v53 - v54
	m.G0 = v55
	F___gettimeofday(m, v55)
	mBase = m.M
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v55)+8)))
	m.G0 = v55 + v54
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, _consts[204])) = v59 + v58*int64(1000000) - int64(946684800000000)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v73 = int32(4403804)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, _consts[205])))
	v76 = int32(1)
	v77 = v75 + v76
	*(*uint16)(unsafe.Add(mBase, _consts[205])) = uint16(v77)
	v79 = int32(4474964)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v81 + v76
	v86 = l0 & int32(3)
	if v86 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v92 = F_LWLockAcquire(m, v88+int32(1152), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v114 = F__emscripten_memset_bulkmem(m, v19+int32(24), base.I32_extend8_s(int32(0)), int32(88))
	mBase = m.M
	goto L19
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = int32(3)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	F_update_controlfile(m, v99, v95)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v103+int32(1152))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v115 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v115
	if v86 != 0 {
		v219 = v43
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v219
	v233 = F_GetLastImportantRecPtr(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L11
	} else {
		goto L43
	}
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v118 <= int32(0) {
		v219 = v43
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v126 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v128 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v132 = F_LWLockAcquire(m, v128+int32(384), int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v138 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v138+int32(384))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v148 = F_LWLockAcquire(m, v144+int32(512), int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if int32(0) < v150 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v156 = v136
	v157 = int32(0)
	v160 = v150
	goto L29
L27:
	;
	v197 = v136
	goto L28
L28:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v211+int32(512))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L42
	}
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v124+v157<<(uint(int32(2))%32))))
	if base.Ui32(int32(3)) <= base.Ui32(v172) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v197 = v189
	goto L28
L31:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v156))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v172)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v189 = v156
	v190 = v160
	goto L33
L33:
	;
	v192 = v157 + int32(1)
	if v192 < v190 {
		v156 = v189
		v157 = v192
		v160 = v190
		goto L29
	} else {
		goto L41
	}
L34:
	;
	if v186 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v186 = base.B2i32(base.Ui32(v172) < base.Ui32(v156))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v186 = int32(base.Ui32(v172-v156) >> (uint(int32(31)) % 32))
	goto L34
L38:
	;
	v187 = v172
	goto L40
L39:
	;
	v187 = v156
	goto L40
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v189 = v187
	v190 = v188
	goto L33
L41:
	;
	goto L30
L42:
	;
	v219 = v197
	goto L20
L43:
	;
	if l0&int32(11) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	m.G0 = v19 + int32(1152)
	return v1106
L45:
	;
	if v42 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	v237 = int32(0)
	v239 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v239)+32))
	if v233 != v240 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v242 = int32(4474964)
	v244 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v244 - int32(1)
	v250 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	if v250 == int32(0) {
		v1106 = v237
		goto L44
	} else {
		goto L49
	}
L49:
	;
	F_errmsg_internal(m, int32(386737), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(493113), int32(7019), int32(88502))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v1106 = v237
	goto L44
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v281
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L11
	} else {
		goto L56
	}
L53:
	;
	v266 = int32(4095044)
	v267 = *(*int32)(unsafe.Add(mBase, _consts[180]))
	*(*int32)(unsafe.Add(mBase, _consts[180])) = int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)+312))
	v281 = v275
	v282 = v267
	goto L52
L54:
	;
	goto L55
L55:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v279
	v281 = v279
	v282 = int32(0)
	goto L52
L56:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+160)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+40)) = uint8(v286)
	v289 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v289
	if v86 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+440)) = int32(1)
	if v365 != 0 {
		goto L75
	} else {
		goto L76
	}
L58:
	;
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v293 = int64(*(*int32)(unsafe.Add(mBase, _consts[206])))
	v294 = base.I64_div_u_s(v291, v293)
	v296 = v291 - v294*v293
	if base.Ui64(v296) <= base.Ui64(int64(8151)) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L71
	}
L61:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v321 = v294*base.I64_extend_i32_s(v316) + v314&int64(4294967295)
	if v321&int64(8191) != int64(0) {
		v336 = v321
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v314 = v296 + int64(40)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v302 = v296 - int64(8152)
	v303 = int64(8168)
	v304 = base.I64_div_u_s(v302, v303)
	v314 = v302 - v304*v303 + v304<<(uint(int64(13))%64) + int64(8216)
	goto L61
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v336
	v339 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	*(*int64)(unsafe.Add(mBase, uint32(v339)+152)) = v336
	*(*int64)(unsafe.Add(mBase, _consts[207])) = v336
	F_WALInsertLockRelease(m)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L70
	}
L66:
	;
	if v321&base.I64_extend_i32_s(v316-int32(1)) == int64(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v336 = v321 | int64(40)
	goto L65
L68:
	;
	goto L69
L69:
	;
	v336 = v321 | int64(24)
	goto L65
L70:
	;
	goto L57
L71:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	F_XLogRegisterData(m, int32(4095012), int32(4))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v355 = F_XLogInsert(m, int32(0), int32(224))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	v358 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v358
	goto L57
L75:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_s_lock(m, v369+int32(440), int32(493113), int32(7113), int32(88502))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L11
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v379 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	*(*int32)(unsafe.Add(mBase, uint32(v379)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v379)+200)) = v377
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _consts[208])))
	if v384 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_LogCheckpointStart(m, l0, int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v86 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(87771)
	if v42 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v416 = F_LWLockAcquire(m, v412+int32(384), int32(1))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L11
	} else {
		goto L93
	}
L86:
	;
	v394 = int32(736510)
	goto L88
L87:
	;
	v394 = int32(710111)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v394
	if l0&int32(1) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v400 = int32(718392)
	goto L91
L90:
	;
	v400 = int32(736510)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v400
	v406 = F_pg_snprintf(m, v19+int32(112), int32(128), int32(174189), v19)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	v410 = F_strlen(m, v19+int32(112))
	mBase = m.M
	goto L85
L93:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v419)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v419)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v424
	v427 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v427+int32(384))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v437 = F_LWLockAcquire(m, v433+int32(4992), int32(1))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v440)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v443
	v446 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v446+int32(4992))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v456 = F_LWLockAcquire(m, v452+int32(256), int32(1))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v460
	if v86 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v464 + v460
	goto L100
L99:
	;
	goto L100
L100:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v468+int32(256))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v486 = F_LWLockAcquire(m, v482+int32(1664), int32(1))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(60)))) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v492
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(76)))) = v494
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v489)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(80)))) = v496
	v499 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v499+int32(1664))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	v504 = int32(4474964)
	v506 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v507 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v506 - v507
	v513 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), v507)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v515 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	goto L108
L106:
	;
	goto L107
L107:
	;
	F_pfree(m, v513)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L11
	} else {
		goto L114
	}
L108:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L11
	} else {
		goto L110
	}
L109:
	;
	goto L107
L110:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = int32(134217738)
	F_pg_usleep(m, int32(10000))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = int32(0)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v549 = F_HaveVirtualXIDsDelayingChkpt(m, v513, v547, int32(1))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	if v549 != 0 {
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v569 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	F_CheckPointGuts(m, v569, l0)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	v575 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), int32(2))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v577 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	goto L120
L118:
	;
	goto L119
L119:
	;
	F_pfree(m, v575)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L11
	} else {
		goto L126
	}
L120:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L11
	} else {
		goto L122
	}
L121:
	;
	goto L119
L122:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v599))) = int32(134217737)
	F_pg_usleep(m, int32(10000))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = int32(0)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v611 = F_HaveVirtualXIDsDelayingChkpt(m, v575, v609, int32(2))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	if v611 != 0 {
		goto L120
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	if v86 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v637 = int32(4474964)
	v639 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v639 + int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L11
	} else {
		goto L131
	}
L128:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v632 <= int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v635 = F_LogStandbySnapshot(m)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	F_XLogRegisterData(m, v19+int32(24), int32(88))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	v650 = int32(0)
	v655 = F_XLogInsert(m, v650, base.B2i32(v86 == v650)<<(uint(int32(4))%32))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	F_XLogFlush(m, v655)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L11
	} else {
		goto L134
	}
L134:
	;
	if v86 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _consts[180])) = v282
	v661 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v663 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	if v661 != v663 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v667 = *(*int64)(unsafe.Add(mBase, uint32(v666)+40))
	v669 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v673 = F_LWLockAcquire(m, v669+int32(1152), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L11
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	if v86 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v676)+16)) = int32(1)
	goto L142
L141:
	;
	goto L142
L142:
	;
	v680 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	*(*int64)(unsafe.Add(mBase, uint32(v676)+32)) = v680
	goto L144
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v676)+144)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v676)+136)) = int64(0)
	v694 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v696 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v696)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v694)+128)) = v697
	v700 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	F_update_controlfile(m, v700, v694)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L11
	} else {
		goto L147
	}
L144:
	;
	v687 = F__emscripten_memcpy_bulkmem(m, v676+int32(40), v19+int32(24), int32(88))
	mBase = m.M
	goto L146
L146:
	;
	goto L143
L147:
	;
	v704 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v704+int32(1152))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L11
	} else {
		goto L148
	}
L148:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v710)+440)) = int32(1)
	if v711 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_s_lock(m, v715+int32(440), int32(493113), int32(7312), int32(88502))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L11
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	v725 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v726 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v725)+440)) = v726
	*(*int64)(unsafe.Add(mBase, uint32(v725)+208)) = v723
	v729 = int32(4474964)
	v731 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v731 - int32(1)
	v736 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	if v736 == v726 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v768 = int32(0)
	v770 = m.G0
	v772 = v770 - int32(1040)
	m.G0 = v772
	v775 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	if v775 != 0 {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v744 = F_LWLockAcquire(m, v740+int32(6272), int32(1))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L11
	} else {
		goto L155
	}
L155:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+20))
	v750 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v750+int32(6272))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L11
	} else {
		goto L156
	}
L156:
	;
	if v748 == int32(-1) {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	F_SetLatch(m, v759+v748*int32(640)+int32(20))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L11
	} else {
		goto L158
	}
L158:
	;
	goto L153
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[211])) = v958
	m.G0 = v772 + int32(1040)
	v973 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	if v667 != int64(0) {
		goto L193
	} else {
		goto L194
	}
L160:
	;
	v893 = int32(0)
	v895 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v895)+12))
	v899 = (v799 - v896) >> (uint(int32(2)) % 32)
	if v893 < v899 {
		goto L185
	} else {
		goto L186
	}
L161:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	if int32(0) < v776 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v890 = int32(0)
	goto L163
L163:
	;
	F_list_free_deep(m, v890)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L11
	} else {
		goto L184
	}
L164:
	;
	v781 = v768
	v786 = int32(10)
	goto L167
L165:
	;
	goto L166
L166:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v890 = v872
	goto L163
L167:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v775)+12))
	v799 = v796 + v781<<(uint(int32(2))%32)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800)+26)))
	if v801 != 0 {
		v850 = v786
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v852 = v781 + int32(1)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	if v852 < v853 {
		v781 = v852
		v786 = v850
		goto L167
	} else {
		goto L183
	}
L170:
	;
	v802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v800)+24)))
	v804 = int32(*(*uint16)(unsafe.Add(mBase, _consts[205])))
	if v802 == v804 {
		goto L160
	} else {
		goto L171
	}
L171:
	;
	v808 = int32(*(*int16)(unsafe.Add(mBase, uint32(v800))))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v808*int32(12))+uint32(_consts[212])))
	v814 = m.T0[v813].(func(*base.Module, int32, int32) int32)(m, v800, v772+int32(16))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L11
	} else {
		goto L173
	}
L172:
	;
	v841 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v800)+26)) = uint8(v841)
	v844 = v786 - v841
	if int32(0) < v844 {
		v850 = v844
		goto L169
	} else {
		goto L181
	}
L173:
	;
	if int32(0) <= v814 {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v819 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v819 == int32(44) {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v824 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L11
	} else {
		goto L176
	}
L176:
	;
	if v824 == int32(0) {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L11
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v772 + int32(16)
	F_errmsg(m, int32(296337), v772)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L11
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(494945), int32(243), int32(87782))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L11
	} else {
		goto L180
	}
L180:
	;
	goto L172
L181:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L11
	} else {
		goto L182
	}
L182:
	;
	v850 = int32(10)
	goto L169
L183:
	;
	goto L168
L184:
	;
	v958 = v768
	goto L159
L185:
	;
	v903 = v893
	goto L188
L186:
	;
	v948 = v895
	goto L187
L187:
	;
	v949 = F_list_delete_first_n(m, v948, v899)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L11
	} else {
		goto L192
	}
L188:
	;
	v919 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)+12))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v920+v903<<(uint(int32(2))%32))))
	F_pfree(m, v924)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L11
	} else {
		goto L190
	}
L189:
	;
	v931 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v948 = v931
	goto L187
L190:
	;
	v928 = v903 + int32(1)
	if v928 != v899 {
		v903 = v928
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v958 = v949
	goto L159
L193:
	;
	v978 = base.F64_convert_i64_u(v973 - v667)
	*(*float64)(unsafe.Add(mBase, _consts[213])) = v978
	v980 = int32(4376088)
	v982 = *(*float64)(unsafe.Add(mBase, _consts[214]))
	if base.F64_gt(v978, v982) != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L195
L195:
	;
	v994 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	v995 = base.I64_div_u_s(v973, v994)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v995
	F_KeepLogSeg(m, v655, v19+int32(16))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L11
	} else {
		goto L199
	}
L196:
	;
	v989 = v978
	goto L198
L197:
	;
	v989 = base.F64_add(base.F64_mul(v982, float64(0.9)), base.F64_mul(v978, float64(0.1)))
	goto L198
L198:
	;
	*(*float64)(unsafe.Add(mBase, _consts[214])) = v989
	goto L195
L199:
	;
	v1002 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1003 = int32(0)
	v1005 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v1002, v1003, v1003)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L11
	} else {
		goto L200
	}
L200:
	;
	if v1005 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1008 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	v1010 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	v1011 = base.I64_div_u_s(v1008, v1010)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1011
	F_KeepLogSeg(m, v655, v19+int32(16))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L11
	} else {
		goto L204
	}
L202:
	;
	v1018 = v1002
	goto L203
L203:
	;
	v1022 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	F_RemoveOldXlogFiles(m, v1018-int64(1), v1022, v655, v1023)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L11
	} else {
		goto L205
	}
L204:
	;
	v1017 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1018 = v1017
	goto L203
L205:
	;
	if v86 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v1081 == int32(1) {
		goto L220
	} else {
		goto L221
	}
L207:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027)+320)))
	if v1028 != int32(1) {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v1033 = v655 - int64(1)
	v1035 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v1042 = base.F64_mul(base.F64_convert_i32_s(v1035), float64(0.75))
	if base.F64_lt(v1042, float64(4.294967296e+09))&base.F64_ge(v1042, float64(0)) != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if base.Ui64(v1033&base.I64_extend_i32_s(v1035-int32(1))) < base.Ui64(base.I64_extend_i32_u(v1050)) {
		goto L206
	} else {
		goto L213
	}
L210:
	;
	v1048 = base.I32_trunc_f64_u(v1042)
	v1050 = v1048
	goto L209
L211:
	;
	goto L212
L212:
	;
	v1050 = int32(0)
	goto L209
L213:
	;
	v1054 = base.I64_div_u_s(v1033, base.I64_extend_i32_s(v1035))
	v1061 = F_XLogFileInitInternal(m, v1054+int64(1), v1031, v19+int32(1151), v19+int32(112))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L11
	} else {
		goto L214
	}
L214:
	;
	if int32(0) <= v1061 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1065 = F_close(m, v1061)
	mBase = m.M
	goto L217
L216:
	;
	goto L217
L217:
	;
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1151)))
	if v1066 != int32(1) {
		goto L206
	} else {
		goto L218
	}
L218:
	;
	v1069 = int32(4376024)
	v1071 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	*(*int32)(unsafe.Add(mBase, _consts[215])) = v1071 + int32(1)
	goto L206
L219:
	;
	F_LogCheckpointEnd(m, int32(0))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L11
	} else {
		goto L226
	}
L220:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+316))
	v1089 = base.B2i32(v1087 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v1089)
	if v1087 != int32(2) {
		goto L219
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1092 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L11
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	F_TruncateSUBTRANS(m, v1092)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L11
	} else {
		goto L225
	}
L225:
	;
	goto L219
L226:
	;
	v1100 = int32(1)
	if v86 == int32(0) {
		v1106 = v1100
		goto L44
	} else {
		goto L227
	}
L227:
	;
	v1106 = v1100
	goto L44
L228:
	;
	F_errmsg_internal(m, int32(14258), int32(0))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L11
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(493113), int32(6954), int32(88502))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L11
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	F_errmsg(m, int32(241963), int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L11
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(493113), int32(7281), int32(88502))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L11
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_construct_point(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v5
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v7
		return v9
	}
}
func F_point_above(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_point_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_point_dt(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Float8GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_point_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_point_mul_point(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v49 float64
	_ = v49
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v89 float64
	_ = v89
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v15 = base.F64_mul(v13, v14)
	v16 = base.F64_abs(v15)
	if base.F64_ne(v16, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v89
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v49
	return
L2:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L43
	} else {
		goto L45
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L43
	} else {
		goto L44
	}
L4:
	;
	if base.F64_ne(v15, float64(0)) != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if base.F64_ne(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v33 = base.F64_mul(v31, v32)
	v34 = base.F64_abs(v33)
	if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if base.F64_eq(v13, float64(0)) != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if base.F64_ne(v14, float64(0)) != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if base.F64_ne(v33, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if base.F64_eq(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.F64_ne(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v49 = base.F64_sub(v15, v33)
	if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if base.F64_eq(v31, float64(0)) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.F64_ne(v32, float64(0)) != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v57 = base.F64_mul(v13, v32)
	v58 = base.F64_abs(v57)
	if base.F64_ne(v58, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if base.F64_eq(v16, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if base.F64_ne(v57, float64(0)) != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.F64_ne(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v73 = base.F64_mul(v14, v31)
	v74 = base.F64_abs(v73)
	if base.F64_ne(v74, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if base.F64_eq(v13, float64(0)) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if base.F64_ne(v32, float64(0)) != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	if base.F64_ne(v73, float64(0)) != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.F64_ne(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v89 = base.F64_add(v73, v57)
	if base.F64_ne(base.F64_abs(v89), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L40
	}
L37:
	;
	if base.F64_eq(v14, float64(0)) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if base.F64_ne(v31, float64(0)) != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	if base.F64_eq(v74, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if base.F64_eq(v58, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L3
L43:
	;
	return
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 float64
	_ = v22
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v41 int64
	_ = v41
	var v46 float64
	_ = v46
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 float64
	_ = v72
	var v76 int64
	_ = v76
	var v77 float64
	_ = v77
	var v80 int64
	_ = v80
	var v89 int32
	_ = v89
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
		v20 = int64(9223372036854775807)
		v21 = base.I64_reinterpret_f64(v18) & v20
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
		v25 = base.I64_reinterpret_f64(v22) & v20
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v25) {
			v64 = base.B2i32(base.Ui64(v21) < base.Ui64(int64(9218868437227405313)))
			v66 = int32(0)
			if base.F64_ne(v12, v18) != 0 {
				v89 = v66
			} else {
				if v64 == int32(0) {
					v89 = v66
				} else {
					v72 = v22
					v76 = v25
					v77 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v80 = base.I64_reinterpret_f64(v77) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
						v89 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v80))
					} else {
						v89 = base.B2i32(base.Ui64(v80) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v77, v72)
					}
				}
			}
		} else {
			v30 = int32(0)
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
				v89 = v30
			} else {
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.F64_ne(v12, v18) != 0 {
						if base.F64_le(base.F64_abs(base.F64_sub(v12, v18)), float64(1e-06)) == int32(0) {
							v89 = v30
						} else {
							v89 = base.F64_eq(v22, v33) | base.F64_le(base.F64_abs(base.F64_sub(v22, v33)), float64(1e-06))
						}
					} else {
						v89 = base.F64_eq(v22, v33) | base.F64_le(base.F64_abs(base.F64_sub(v22, v33)), float64(1e-06))
					}
				} else {
					v64 = int32(1)
					v66 = int32(0)
					if base.F64_ne(v12, v18) != 0 {
						v89 = v66
					} else {
						if v64 == int32(0) {
							v89 = v66
						} else {
							v72 = v22
							v76 = v25
							v77 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v80 = base.I64_reinterpret_f64(v77) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
								v89 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v80))
							} else {
								v89 = base.B2i32(base.Ui64(v80) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v77, v72)
							}
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if base.Ui64(v41&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v89 = int32(0)
		} else {
			v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v72 = v46
			v76 = base.I64_reinterpret_f64(v46) & int64(9223372036854775807)
			v77 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v80 = base.I64_reinterpret_f64(v77) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
				v89 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v80))
			} else {
				v89 = base.B2i32(base.Ui64(v80) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v77, v72)
			}
		}
	}
	return v89 ^ int32(1)
}
func F_point_slope(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_point_sl(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Float8GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
