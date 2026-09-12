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
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int64
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v624 int32
	_ = v624
	var v625 int64
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int64
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int64
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v717 int64
	_ = v717
	var v719 int64
	_ = v719
	var v722 int32
	_ = v722
	var v723 int64
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v736 int64
	_ = v736
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v779 int64
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v928 int32
	_ = v928
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1014 int32
	_ = v1014
	var v1029 int64
	_ = v1029
	var v1034 float64
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 float64
	_ = v1038
	var v1045 float64
	_ = v1045
	var v1050 int64
	_ = v1050
	var v1051 int64
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1058 int64
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int64
	_ = v1064
	var v1066 int64
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1074 int64
	_ = v1074
	var v1078 int64
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1089 int64
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1098 float64
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1110 int64
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
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
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L11
	} else {
		goto L248
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L11
	} else {
		goto L245
	}
L3:
	;
	v43 = int32(0)
	v48 = F__emscripten_memset_bulkmem(m, int32(4350888), base.I32_extend8_s(v43), int32(80))
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
	v73 = int32(4378716)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, _consts[205])))
	v76 = int32(1)
	v77 = v75 + v76
	*(*uint16)(unsafe.Add(mBase, _consts[205])) = uint16(v77)
	v79 = int32(4449876)
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
	return v1162
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
	v242 = int32(4449876)
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
		v1162 = v237
		goto L44
	} else {
		goto L49
	}
L49:
	;
	F_errmsg_internal(m, int32(376499), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(481224), int32(7019), int32(85942))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v1162 = v237
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
	v266 = int32(4074036)
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
	F_XLogRegisterData(m, int32(4074004), int32(4))
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
	F_s_lock(m, v369+int32(440), int32(481224), int32(7113), int32(85942))
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(85211)
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
	v468 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v472 = F_LWLockAcquire(m, v468+int32(384), int32(1))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L11
	} else {
		goto L110
	}
L86:
	;
	v394 = int32(722455)
	goto L88
L87:
	;
	v394 = int32(696082)
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
	v400 = int32(704363)
	goto L91
L90:
	;
	v400 = int32(722455)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v400
	v406 = F_pg_snprintf(m, v19+int32(112), int32(128), int32(168237), v19)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	v409 = v19 + int32(112)
	if v409&int32(3) == int32(0) {
		v433 = v409
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L85
L94:
	;
	goto L93
L95:
	;
	v437 = v433
	goto L104
L96:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if v417 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	goto L93
L98:
	;
	goto L99
L99:
	;
	v422 = v409
	goto L100
L100:
	;
	v426 = v422 + int32(1)
	if v426&int32(3) == int32(0) {
		v433 = v426
		goto L95
	} else {
		goto L102
	}
L101:
	;
	goto L94
L102:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v431 != 0 {
		v422 = v426
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v446 = int32(-2139062144)
	if (int32(16843008)-v443|v443)&v446 == v446 {
		v437 = v437 + int32(4)
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v452 = v437
	goto L107
L106:
	;
	goto L105
L107:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v456 != 0 {
		v452 = v452 + int32(1)
		goto L107
	} else {
		goto L109
	}
L108:
	;
	goto L94
L109:
	;
	goto L108
L110:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v475)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v478
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v480
	v483 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v483+int32(384))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v493 = F_LWLockAcquire(m, v489+int32(4992), int32(1))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v497
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v496)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v499
	v502 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v502+int32(4992))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v512 = F_LWLockAcquire(m, v508+int32(256), int32(1))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v516
	if v86 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v520 + v516
	goto L117
L116:
	;
	goto L117
L117:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v524+int32(256))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v542 = F_LWLockAcquire(m, v538+int32(1664), int32(1))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	v545 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(60)))) = v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v548
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v545)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(76)))) = v550
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v545)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(80)))) = v552
	v555 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v555+int32(1664))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	v560 = int32(4449876)
	v562 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v563 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v562 - v563
	v569 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), v563)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v571 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	goto L125
L123:
	;
	goto L124
L124:
	;
	F_pfree(m, v569)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L131
	}
L125:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L11
	} else {
		goto L127
	}
L126:
	;
	goto L124
L127:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = int32(134217738)
	F_pg_usleep(m, int32(10000))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v600))) = int32(0)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v605 = F_HaveVirtualXIDsDelayingChkpt(m, v569, v603, int32(1))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	if v605 != 0 {
		goto L125
	} else {
		goto L130
	}
L130:
	;
	goto L126
L131:
	;
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	F_CheckPointGuts(m, v625, l0)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	v631 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), int32(2))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v633 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	goto L137
L135:
	;
	goto L136
L136:
	;
	F_pfree(m, v631)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L11
	} else {
		goto L143
	}
L137:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L11
	} else {
		goto L139
	}
L138:
	;
	goto L136
L139:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v655))) = int32(134217737)
	F_pg_usleep(m, int32(10000))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L11
	} else {
		goto L140
	}
L140:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v662))) = int32(0)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v667 = F_HaveVirtualXIDsDelayingChkpt(m, v631, v665, int32(2))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L11
	} else {
		goto L141
	}
L141:
	;
	if v667 != 0 {
		goto L137
	} else {
		goto L142
	}
L142:
	;
	goto L138
L143:
	;
	if v86 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v693 = int32(4449876)
	v695 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v695 + int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L11
	} else {
		goto L148
	}
L145:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v688 <= int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v691 = F_LogStandbySnapshot(m)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	F_XLogRegisterData(m, v19+int32(24), int32(88))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L11
	} else {
		goto L149
	}
L149:
	;
	v706 = int32(0)
	v711 = F_XLogInsert(m, v706, base.B2i32(v86 == v706)<<(uint(int32(4))%32))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L11
	} else {
		goto L150
	}
L150:
	;
	F_XLogFlush(m, v711)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L11
	} else {
		goto L151
	}
L151:
	;
	if v86 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, _consts[180])) = v282
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v719 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	if v717 != v719 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v722 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v722)+40))
	v725 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v729 = F_LWLockAcquire(m, v725+int32(1152), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L11
	} else {
		goto L156
	}
L155:
	;
	goto L154
L156:
	;
	v732 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	if v86 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+16)) = int32(1)
	goto L159
L158:
	;
	goto L159
L159:
	;
	v736 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	*(*int64)(unsafe.Add(mBase, uint32(v732)+32)) = v736
	goto L161
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+144)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v732)+136)) = int64(0)
	v750 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v752 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v752)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v750)+128)) = v753
	v756 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	F_update_controlfile(m, v756, v750)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L11
	} else {
		goto L164
	}
L161:
	;
	v743 = F__emscripten_memcpy_bulkmem(m, v732+int32(40), v19+int32(24), int32(88))
	mBase = m.M
	goto L163
L163:
	;
	goto L160
L164:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v760+int32(1152))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L11
	} else {
		goto L165
	}
L165:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v766)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+440)) = int32(1)
	if v767 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_s_lock(m, v771+int32(440), int32(481224), int32(7312), int32(85942))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L11
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v779 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	v781 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v782 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v781)+440)) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v781)+208)) = v779
	v785 = int32(4449876)
	v787 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v787 - int32(1)
	v792 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	if v792 == v782 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L168
L170:
	;
	v824 = int32(0)
	v826 = m.G0
	v828 = v826 - int32(1040)
	m.G0 = v828
	v831 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	if v831 != 0 {
		goto L178
	} else {
		goto L179
	}
L171:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v800 = F_LWLockAcquire(m, v796+int32(6272), int32(1))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L11
	} else {
		goto L172
	}
L172:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+20))
	v806 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v806+int32(6272))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L11
	} else {
		goto L173
	}
L173:
	;
	if v804 == int32(-1) {
		goto L170
	} else {
		goto L174
	}
L174:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	F_SetLatch(m, v815+v804*int32(640)+int32(20))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L11
	} else {
		goto L175
	}
L175:
	;
	goto L170
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[211])) = v1014
	m.G0 = v828 + int32(1040)
	v1029 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	if v723 != int64(0) {
		goto L210
	} else {
		goto L211
	}
L177:
	;
	v949 = int32(0)
	v951 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	v955 = (v855 - v952) >> (uint(int32(2)) % 32)
	if v949 < v955 {
		goto L202
	} else {
		goto L203
	}
L178:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if int32(0) < v832 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v946 = int32(0)
	goto L180
L180:
	;
	F_list_free_deep(m, v946)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L11
	} else {
		goto L201
	}
L181:
	;
	v837 = v824
	v842 = int32(10)
	goto L184
L182:
	;
	goto L183
L183:
	;
	v928 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v946 = v928
	goto L180
L184:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	v855 = v852 + v837<<(uint(int32(2))%32)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+26)))
	if v857 != 0 {
		v906 = v842
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L183
L186:
	;
	v908 = v837 + int32(1)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v908 < v909 {
		v837 = v908
		v842 = v906
		goto L184
	} else {
		goto L200
	}
L187:
	;
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v856)+24)))
	v860 = int32(*(*uint16)(unsafe.Add(mBase, _consts[205])))
	if v858 == v860 {
		goto L177
	} else {
		goto L188
	}
L188:
	;
	v864 = int32(*(*int16)(unsafe.Add(mBase, uint32(v856))))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v864*int32(12))+uint32(_consts[212])))
	v870 = m.T0[v869].(func(*base.Module, int32, int32) int32)(m, v856, v828+int32(16))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L11
	} else {
		goto L190
	}
L189:
	;
	v897 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v856)+26)) = uint8(v897)
	v900 = v842 - v897
	if int32(0) < v900 {
		v906 = v900
		goto L186
	} else {
		goto L198
	}
L190:
	;
	if int32(0) <= v870 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v875 == int32(44) {
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v880 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	if v880 == int32(0) {
		goto L189
	} else {
		goto L194
	}
L194:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L11
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828))) = v828 + int32(16)
	F_errmsg(m, int32(288466), v828)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L11
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(483056), int32(243), int32(85222))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L11
	} else {
		goto L197
	}
L197:
	;
	goto L189
L198:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L11
	} else {
		goto L199
	}
L199:
	;
	v906 = int32(10)
	goto L186
L200:
	;
	goto L185
L201:
	;
	v1014 = v824
	goto L176
L202:
	;
	v959 = v949
	goto L205
L203:
	;
	v1004 = v951
	goto L204
L204:
	;
	v1005 = F_list_delete_first_n(m, v1004, v955)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L11
	} else {
		goto L209
	}
L205:
	;
	v975 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v975)+12))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v976+v959<<(uint(int32(2))%32))))
	F_pfree(m, v980)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L11
	} else {
		goto L207
	}
L206:
	;
	v987 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v1004 = v987
	goto L204
L207:
	;
	v984 = v959 + int32(1)
	if v984 != v955 {
		v959 = v984
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v1014 = v1005
	goto L176
L210:
	;
	v1034 = base.F64_convert_i64_u(v1029 - v723)
	*(*float64)(unsafe.Add(mBase, _consts[213])) = v1034
	v1036 = int32(4351000)
	v1038 = *(*float64)(unsafe.Add(mBase, _consts[214]))
	if base.F64_gt(v1034, v1038) != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	goto L212
L212:
	;
	v1050 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	v1051 = base.I64_div_u_s(v1029, v1050)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1051
	F_KeepLogSeg(m, v711, v19+int32(16))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L11
	} else {
		goto L216
	}
L213:
	;
	v1045 = v1034
	goto L215
L214:
	;
	v1045 = base.F64_add(base.F64_mul(v1038, float64(0.9)), base.F64_mul(v1034, float64(0.1)))
	goto L215
L215:
	;
	*(*float64)(unsafe.Add(mBase, _consts[214])) = v1045
	goto L212
L216:
	;
	v1058 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1059 = int32(0)
	v1061 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v1058, v1059, v1059)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L11
	} else {
		goto L217
	}
L217:
	;
	if v1061 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1064 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	v1066 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	v1067 = base.I64_div_u_s(v1064, v1066)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1067
	F_KeepLogSeg(m, v711, v19+int32(16))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L11
	} else {
		goto L221
	}
L219:
	;
	v1074 = v1058
	goto L220
L220:
	;
	v1078 = *(*int64)(unsafe.Add(mBase, _consts[207]))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	F_RemoveOldXlogFiles(m, v1074-int64(1), v1078, v711, v1079)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L11
	} else {
		goto L222
	}
L221:
	;
	v1073 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1074 = v1073
	goto L220
L222:
	;
	if v86 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v1137 == int32(1) {
		goto L237
	} else {
		goto L238
	}
L224:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+320)))
	if v1084 != int32(1) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v1089 = v711 - int64(1)
	v1091 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v1098 = base.F64_mul(base.F64_convert_i32_s(v1091), float64(0.75))
	if base.F64_lt(v1098, float64(4.294967296e+09))&base.F64_ge(v1098, float64(0)) != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	if base.Ui64(v1089&base.I64_extend_i32_s(v1091-int32(1))) < base.Ui64(base.I64_extend_i32_u(v1106)) {
		goto L223
	} else {
		goto L230
	}
L227:
	;
	v1104 = base.I32_trunc_f64_u(v1098)
	v1106 = v1104
	goto L226
L228:
	;
	goto L229
L229:
	;
	v1106 = int32(0)
	goto L226
L230:
	;
	v1110 = base.I64_div_u_s(v1089, base.I64_extend_i32_s(v1091))
	v1117 = F_XLogFileInitInternal(m, v1110+int64(1), v1087, v19+int32(1151), v19+int32(112))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L11
	} else {
		goto L231
	}
L231:
	;
	if int32(0) <= v1117 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1121 = F_close(m, v1117)
	mBase = m.M
	goto L234
L233:
	;
	goto L234
L234:
	;
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1151)))
	if v1122 != int32(1) {
		goto L223
	} else {
		goto L235
	}
L235:
	;
	v1125 = int32(4350936)
	v1127 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	*(*int32)(unsafe.Add(mBase, _consts[215])) = v1127 + int32(1)
	goto L223
L236:
	;
	F_LogCheckpointEnd(m, int32(0))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L11
	} else {
		goto L243
	}
L237:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+316))
	v1145 = base.B2i32(v1143 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v1145)
	if v1143 != int32(2) {
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1148 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L11
	} else {
		goto L241
	}
L240:
	;
	goto L239
L241:
	;
	F_TruncateSUBTRANS(m, v1148)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L11
	} else {
		goto L242
	}
L242:
	;
	goto L236
L243:
	;
	v1156 = int32(1)
	if v86 == int32(0) {
		v1162 = v1156
		goto L44
	} else {
		goto L244
	}
L244:
	;
	v1162 = v1156
	goto L44
L245:
	;
	F_errmsg_internal(m, int32(13419), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L11
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(481224), int32(6954), int32(85942))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L11
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errmsg(m, int32(235149), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L11
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(481224), int32(7281), int32(85942))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L11
	} else {
		goto L250
	}
L250:
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
