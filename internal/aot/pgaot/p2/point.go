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
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
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
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
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
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v312 int64
	_ = v312
	var v314 int32
	_ = v314
	var v319 int64
	_ = v319
	var v334 int64
	_ = v334
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int64
	_ = v353
	var v354 int32
	_ = v354
	var v356 int64
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int64
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int64
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int64
	_ = v657
	var v659 int64
	_ = v659
	var v662 int32
	_ = v662
	var v663 int64
	_ = v663
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int64
	_ = v676
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int64
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v718 int64
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
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
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v866 int32
	_ = v866
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v967 int64
	_ = v967
	var v972 float64
	_ = v972
	var v974 int32
	_ = v974
	var v976 float64
	_ = v976
	var v983 float64
	_ = v983
	var v988 int64
	_ = v988
	var v989 int64
	_ = v989
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int64
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int64
	_ = v1002
	var v1004 int64
	_ = v1004
	var v1005 int64
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int64
	_ = v1009
	var v1010 int64
	_ = v1010
	var v1014 int64
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1024 int64
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1038 int64
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1130 int32
	_ = v1130
	v17 = m.G0
	v19 = v17 - int32(1152)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L10
	} else {
		goto L226
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L10
	} else {
		goto L223
	}
L3:
	;
	v43 = int32(0)
	v44 = int32(_a_F_CreateCheckPoint_0)
	base.MemoryFill(m, v44, v43, int32(80))
	v52 = m.G0
	v53 = int32(16)
	v54 = v52 - v53
	m.G0 = v54
	F_gettimeofday(m, v54)
	mBase = m.M
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	v58 = int64(*(*int32)(unsafe.Add(mBase, uint32(v54)+8)))
	m.G0 = v54 + v53
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
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])) = uint8(v34)
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
	*(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[2])) = v58 + v57*int64(1000000) - int64(946684800000000)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v72 = int32(_a_F_CreateCheckPoint_1)
	v74 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])))
	v75 = int32(1)
	v76 = v74 + v75
	*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])) = uint16(v76)
	v78 = int32(_a_F_CreateCheckPoint_2)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v80 + v75
	v85 = l0 & int32(3)
	if v85 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v91 = F_LWLockAcquire(m, v87+int32(1152), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	base.MemoryFill(m, v19+int32(24), int32(0), int32(88))
	v113 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v113
	if v85 != 0 {
		v216 = v43
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = int32(3)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[7]))
	F_update_controlfile(m, v98, v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v102+int32(1152))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v216
	v231 = F_GetLastImportantRecPtr(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L41
	}
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	if v116 <= int32(0) {
		v216 = v43
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[9]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[10]))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v130 = F_LWLockAcquire(m, v126+int32(384), int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v136+int32(384))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v146 = F_LWLockAcquire(m, v142+int32(512), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if int32(0) < v148 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v153 = v134
	v154 = v148
	v155 = int32(0)
	goto L27
L25:
	;
	v194 = v134
	goto L26
L26:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v209+int32(512))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L10
	} else {
		goto L40
	}
L27:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v122+v155<<(uint(int32(2))%32))))
	if base.Ui32(int32(3)) <= base.Ui32(v170) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v194 = v187
	goto L26
L29:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v153))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v170)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v187 = v153
	v188 = v154
	goto L31
L31:
	;
	v190 = v155 + int32(1)
	if v190 < v188 {
		v153 = v187
		v154 = v188
		v155 = v190
		goto L27
	} else {
		goto L39
	}
L32:
	;
	if v184 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v184 = base.B2i32(base.Ui32(v170) < base.Ui32(v153))
	goto L32
L34:
	;
	goto L35
L35:
	;
	v184 = int32(base.Ui32(v170-v153) >> (uint(int32(31)) % 32))
	goto L32
L36:
	;
	v185 = v170
	goto L38
L37:
	;
	v185 = v153
	goto L38
L38:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v187 = v185
	v188 = v186
	goto L31
L39:
	;
	goto L28
L40:
	;
	v216 = v194
	goto L18
L41:
	;
	if l0&int32(11) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	m.G0 = v19 + int32(1152)
	return v1087
L43:
	;
	if v42 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	v235 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)+32))
	if v231 != v238 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v240 = int32(_a_F_CreateCheckPoint_2)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v242 - int32(1)
	v248 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	if v248 == int32(0) {
		v1087 = v235
		goto L42
	} else {
		goto L47
	}
L47:
	;
	F_errmsg_internal(m, int32(_a_F_CreateCheckPoint_3), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_5), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v1087 = v235
	goto L42
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v279
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L10
	} else {
		goto L54
	}
L51:
	;
	v264 = int32(_a_F_CreateCheckPoint_7)
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12])) = int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+312))
	v279 = v273
	v280 = v265
	goto L50
L52:
	;
	goto L53
L53:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v277
	v279 = v277
	v280 = int32(0)
	goto L50
L54:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+160)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+40)) = uint8(v284)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v287
	if v85 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+440)) = int32(1)
	if v363 != 0 {
		goto L73
	} else {
		goto L74
	}
L56:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v291 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[13])))
	v292 = base.I64_div_u_s(v289, v291)
	v294 = v289 - v292*v291
	if base.Ui64(v294) <= base.Ui64(int64(8151)) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L10
	} else {
		goto L69
	}
L59:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14]))
	v319 = v292*base.I64_extend_i32_s(v314) + v312&int64(4294967295)
	if v319&int64(8191) != int64(0) {
		v334 = v319
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v312 = v294 + int64(40)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v300 = v294 - int64(8152)
	v301 = int64(8168)
	v302 = base.I64_div_u_s(v300, v301)
	v312 = v300 - v302*v301 + v302<<(uint(int64(13))%64) + int64(8216)
	goto L59
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v334
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v337)+152)) = v334
	*(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15])) = v334
	F_WALInsertLockRelease(m)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L10
	} else {
		goto L68
	}
L64:
	;
	if v319&base.I64_extend_i32_s(v314-int32(1)) == int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v334 = v319 | int64(40)
	goto L63
L66:
	;
	goto L67
L67:
	;
	v334 = v319 | int64(24)
	goto L63
L68:
	;
	goto L55
L69:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	F_XLogRegisterData(m, int32(_a_F_CreateCheckPoint_8), int32(4))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	v353 = F_XLogInsert(m, int32(0), int32(224))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	v356 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v356
	goto L55
L73:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	F_s_lock(m, v367+int32(440), int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_9), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L10
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v377)+200)) = v375
	v382 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[16])))
	if v382 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	F_LogCheckpointStart(m, l0, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v85 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(_a_F_CreateCheckPoint_10)
	if v42 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v413 = F_LWLockAcquire(m, v409+int32(384), int32(1))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L91
	}
L84:
	;
	v392 = int32(_a_F_CreateCheckPoint_11)
	goto L86
L85:
	;
	v392 = int32(_a_F_CreateCheckPoint_12)
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v392
	if l0&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v398 = int32(_a_F_CreateCheckPoint_13)
	goto L89
L88:
	;
	v398 = int32(_a_F_CreateCheckPoint_11)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v398
	v401 = v19 + int32(112)
	v404 = F_pg_snprintf(m, v401, int32(128), int32(_a_F_CreateCheckPoint_14), v19)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v406 = F_strlen(m, v401)
	mBase = m.M
	goto L83
L91:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v416)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v416)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v421
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v424+int32(384))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v434 = F_LWLockAcquire(m, v430+int32(_a_F_CreateCheckPoint_15), int32(1))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v438
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v440
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v443+int32(_a_F_CreateCheckPoint_15))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v453 = F_LWLockAcquire(m, v449+int32(256), int32(1))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v457
	if v85 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v461 + v457
	goto L98
L97:
	;
	goto L98
L98:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v465+int32(256))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v483 = F_LWLockAcquire(m, v479+int32(1664), int32(1))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[17]))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(60)))) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(76)))) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(80)))) = v493
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v496+int32(1664))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	v501 = int32(_a_F_CreateCheckPoint_2)
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	v504 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v503 - v504
	v510 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), v504)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v512 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	goto L106
L104:
	;
	goto L105
L105:
	;
	F_pfree(m, v510)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L10
	} else {
		goto L112
	}
L106:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = int32(134217738)
	F_pg_usleep(m, int32(_a_F_CreateCheckPoint_16))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = int32(0)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v546 = F_HaveVirtualXIDsDelayingChkpt(m, v510, v544, int32(1))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if v546 != 0 {
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	F_CheckPointGuts(m, v566, l0)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v572 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), int32(2))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v574 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	goto L118
L116:
	;
	goto L117
L117:
	;
	F_pfree(m, v572)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L10
	} else {
		goto L124
	}
L118:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L10
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v596))) = int32(134217737)
	F_pg_usleep(m, int32(_a_F_CreateCheckPoint_16))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v603))) = int32(0)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v608 = F_HaveVirtualXIDsDelayingChkpt(m, v572, v606, int32(2))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
	;
	if v608 != 0 {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	goto L119
L124:
	;
	if v85 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v634 = int32(_a_F_CreateCheckPoint_2)
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v636 + int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L10
	} else {
		goto L129
	}
L126:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	if v629 <= int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v632 = F_LogStandbySnapshot(m)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	F_XLogRegisterData(m, v19+int32(24), int32(88))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L10
	} else {
		goto L130
	}
L130:
	;
	v647 = int32(0)
	if v85 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v650 = v647
	goto L133
L132:
	;
	v650 = int32(16)
	goto L133
L133:
	;
	v651 = F_XLogInsert(m, v647, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	F_XLogFlush(m, v651)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L10
	} else {
		goto L135
	}
L135:
	;
	if v85 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12])) = v280
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v659 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	if v657 != v659 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	v663 = *(*int64)(unsafe.Add(mBase, uint32(v662)+40))
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v669 = F_LWLockAcquire(m, v665+int32(1152), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L10
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	if v85 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v672)+16)) = int32(1)
	goto L143
L142:
	;
	goto L143
L143:
	;
	v676 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	*(*int64)(unsafe.Add(mBase, uint32(v672)+32)) = v676
	base.MemoryCopy(m, v672+int32(40), v19+int32(24), int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v672)+144)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v672)+136)) = int64(0)
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v692 = *(*int64)(unsafe.Add(mBase, uint32(v691)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v689)+128)) = v692
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[7]))
	F_update_controlfile(m, v695, v689)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L10
	} else {
		goto L144
	}
L144:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v699+int32(1152))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L10
	} else {
		goto L145
	}
L145:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v705)+440)) = int32(1)
	if v706 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	F_s_lock(m, v710+int32(440), int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_17), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L10
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v721 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v720)+440)) = v721
	*(*int64)(unsafe.Add(mBase, uint32(v720)+208)) = v718
	v724 = int32(_a_F_CreateCheckPoint_2)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v726 - int32(1)
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	if v731 == v721 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v763 = int32(0)
	v765 = m.G0
	v767 = v765 - int32(1040)
	m.G0 = v767
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	if v770 != 0 {
		goto L158
	} else {
		goto L159
	}
L151:
	;
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v739 = F_LWLockAcquire(m, v735+int32(_a_F_CreateCheckPoint_18), int32(1))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+20))
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v745+int32(_a_F_CreateCheckPoint_18))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	if v743 == int32(-1) {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[9]))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	F_SetLatch(m, v754+v743*int32(640)+int32(20))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	goto L150
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21])) = v948
	m.G0 = v767 + int32(1040)
	v967 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	if v663 != int64(0) {
		goto L192
	} else {
		goto L193
	}
L157:
	;
	v887 = int32(0)
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)+12))
	v893 = (v794 - v890) >> (uint(int32(2)) % 32)
	if v887 < v893 {
		goto L184
	} else {
		goto L185
	}
L158:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	if int32(0) < v771 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v884 = int32(0)
	goto L160
L160:
	;
	F_list_free_deep(m, v884)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L10
	} else {
		goto L183
	}
L161:
	;
	v775 = v763
	v777 = int32(10)
	goto L164
L162:
	;
	goto L163
L163:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v884 = v866
	goto L160
L164:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v770)+12))
	v794 = v791 + v775<<(uint(int32(2))%32)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795)+26)))
	if v796 != 0 {
		v843 = v777
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L163
L166:
	;
	v846 = v775 + int32(1)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	if v846 < v847 {
		v775 = v846
		v777 = v843
		goto L164
	} else {
		goto L182
	}
L167:
	;
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v795)+24)))
	v799 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])))
	if v797 == v799 {
		goto L157
	} else {
		goto L168
	}
L168:
	;
	v802 = v767 + int32(16)
	v803 = int32(*(*int16)(unsafe.Add(mBase, uint32(v795))))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v803*int32(12))+uint32(_c_F_CreateCheckPoint[22])))
	v809 = m.T0[v808].(func(*base.Module, int32, int32) int32)(m, v795, v802)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L10
	} else {
		goto L170
	}
L169:
	;
	v834 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v795)+26)) = uint8(v834)
	if v834 < v777 {
		goto L178
	} else {
		goto L179
	}
L170:
	;
	if int32(0) <= v809 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	if v814 == int32(44) {
		goto L169
	} else {
		goto L172
	}
L172:
	;
	v819 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	if v819 == int32(0) {
		goto L169
	} else {
		goto L174
	}
L174:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L10
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v802
	F_errmsg(m, int32(_a_F_CreateCheckPoint_19), v767)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_20), int32(243), int32(_a_F_CreateCheckPoint_21))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	goto L169
L178:
	;
	v843 = v777 - int32(1)
	goto L166
L179:
	;
	goto L180
L180:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L10
	} else {
		goto L181
	}
L181:
	;
	v843 = int32(10)
	goto L166
L182:
	;
	goto L165
L183:
	;
	v948 = v763
	goto L156
L184:
	;
	v896 = v887
	goto L187
L185:
	;
	v942 = v889
	goto L186
L186:
	;
	v943 = F_list_delete_first_n(m, v942, v893)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L10
	} else {
		goto L191
	}
L187:
	;
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+12))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v914+v896<<(uint(int32(2))%32))))
	F_pfree(m, v918)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L10
	} else {
		goto L189
	}
L188:
	;
	v925 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v942 = v925
	goto L186
L189:
	;
	v922 = v896 + int32(1)
	if v922 != v893 {
		v896 = v922
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v948 = v943
	goto L156
L192:
	;
	v972 = base.F64_convert_i64_u(v967 - v663)
	*(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[24])) = v972
	v974 = int32(_a_F_CreateCheckPoint_22)
	v976 = *(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[25]))
	if base.F64_gt(v972, v976) != 0 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	v988 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14])))
	v989 = base.I64_div_u_s(v967, v988)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v989
	v992 = v19 + int32(16)
	F_KeepLogSeg(m, v651, v992)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L10
	} else {
		goto L198
	}
L195:
	;
	v983 = v972
	goto L197
L196:
	;
	v983 = base.F64_add(base.F64_mul(v976, float64(0.9)), base.F64_mul(v972, float64(0.1)))
	goto L197
L197:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[25])) = v983
	goto L194
L198:
	;
	v996 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v997 = int32(0)
	v999 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v996, v997, v997)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	if v999 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1002 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	v1004 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14])))
	v1005 = base.I64_div_u_s(v1002, v1004)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1005
	F_KeepLogSeg(m, v651, v992)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L10
	} else {
		goto L203
	}
L201:
	;
	v1010 = v996
	goto L202
L202:
	;
	v1014 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	F_RemoveOldXlogFiles(m, v1010-int64(1), v1014, v651, v1015)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L10
	} else {
		goto L204
	}
L203:
	;
	v1009 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1010 = v1009
	goto L202
L204:
	;
	if v85 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])))
	if v1063 == int32(1) {
		goto L215
	} else {
		goto L216
	}
L206:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+320)))
	if v1020 != int32(1) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1024 = v651 - int64(1)
	v1026 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14]))
	if base.Ui64(v1024&base.I64_extend_i32_s(v1026-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v1026), float64(0.75))))) {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	v1038 = base.I64_div_u_s(v1024, base.I64_extend_i32_s(v1026))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v1046 = F_XLogFileInitInternal(m, v1038+int64(1), v1041, v19+int32(1151), v19+int32(112))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L10
	} else {
		goto L209
	}
L209:
	;
	if int32(0) <= v1046 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1050 = F_close(m, v1046)
	mBase = m.M
	goto L212
L211:
	;
	goto L212
L212:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1151)))
	if v1051 != int32(1) {
		goto L205
	} else {
		goto L213
	}
L213:
	;
	v1054 = int32(_a_F_CreateCheckPoint_23)
	v1056 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[26]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[26])) = v1056 + int32(1)
	goto L205
L214:
	;
	F_LogCheckpointEnd(m, int32(0))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L10
	} else {
		goto L221
	}
L215:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+316))
	v1071 = base.B2i32(v1069 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])) = uint8(v1071)
	if v1069 != int32(2) {
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v1074 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L10
	} else {
		goto L219
	}
L218:
	;
	goto L217
L219:
	;
	F_TruncateSUBTRANS(m, v1074)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L10
	} else {
		goto L220
	}
L220:
	;
	goto L214
L221:
	;
	v1082 = int32(1)
	if v85 == int32(0) {
		v1087 = v1082
		goto L42
	} else {
		goto L222
	}
L222:
	;
	v1087 = v1082
	goto L42
L223:
	;
	F_errmsg_internal(m, int32(_a_F_CreateCheckPoint_24), int32(0))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_25), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	F_errmsg(m, int32(_a_F_CreateCheckPoint_26), int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_27), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v7
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v5
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
	var v17 float64
	_ = v17
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v83 int32
	_ = v83
	var v86 float64
	_ = v86
	var v92 int32
	_ = v92
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v109 float64
	_ = v109
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v15 = base.F64_mul(v13, v14)
	v16 = base.F64_abs(v15)
	v17 = math.Float64frombits(uint64(0x7ff0000000000000))
	v21 = base.F64_eq(base.F64_abs(v13), v17)
	if base.B2i32(base.F64_ne(v16, v17)|v21 == int32(0))&base.F64_ne(base.F64_abs(v14), v17) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v29 = float64(0)
		if base.B2i32(base.F64_eq(v13, v29)|base.F64_ne(v15, v29) == int32(0))&base.F64_ne(v14, v29) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v40 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v41 = base.F64_mul(v39, v40)
			v42 = base.F64_abs(v41)
			v43 = math.Float64frombits(uint64(0x7ff0000000000000))
			v53 = base.F64_ne(base.F64_abs(v40), v43)
			if base.B2i32(base.F64_ne(v42, v43)|base.F64_eq(base.F64_abs(v39), v43) == int32(0))&v53 != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v55 = float64(0)
				v56 = base.F64_ne(v40, v55)
				if v56&base.B2i32(base.F64_eq(v39, v55)|base.F64_ne(v41, v55) == int32(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v65 = math.Float64frombits(uint64(0x7ff0000000000000))
					v67 = base.F64_sub(v15, v41)
					if base.B2i32(base.F64_eq(v16, v65)|base.F64_ne(base.F64_abs(v67), v65) == int32(0))&base.F64_ne(v42, v65) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v77 = base.F64_mul(v13, v40)
						v78 = base.F64_abs(v77)
						if v53 != 0 {
							v83 = base.F64_ne(v78, math.Float64frombits(uint64(0x7ff0000000000000))) | v21
						} else {
							v83 = int32(1)
						}
						if v83 == int32(0) {
							F_float_overflow_error(m)
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v86 = float64(0)
							if v56 != 0 {
								v92 = base.F64_eq(v13, v86) | base.F64_ne(v77, v86)
							} else {
								v92 = int32(1)
							}
							if v92 == int32(0) {
								F_float_underflow_error(m)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v96 = math.Float64frombits(uint64(0x7ff0000000000000))
								v98 = base.F64_mul(v14, v39)
								v99 = base.F64_abs(v98)
								if base.B2i32(base.F64_eq(base.F64_abs(v14), v96)|base.F64_ne(v99, v96) == int32(0))&base.F64_ne(base.F64_abs(v39), v96) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v109 = float64(0)
									if base.B2i32(base.F64_eq(v14, v109)|base.F64_ne(v98, v109) == int32(0))&base.F64_ne(v39, v109) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v119 = math.Float64frombits(uint64(0x7ff0000000000000))
										v121 = base.F64_add(v98, v77)
										if base.F64_eq(v99, v119)|base.F64_ne(base.F64_abs(v121), v119)|base.F64_eq(v78, v119) != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v121
											*(*float64)(unsafe.Add(mBase, uint32(l0))) = v67
											return
										} else {
											F_float_overflow_error(m)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 float64
	_ = v74
	var v78 int64
	_ = v78
	var v79 float64
	_ = v79
	var v82 int64
	_ = v82
	var v91 int32
	_ = v91
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
			v65 = base.B2i32(base.Ui64(v21) < base.Ui64(int64(9218868437227405313)))
			v67 = int32(0)
			if base.B2i32(v65 == v67)|base.F64_ne(v12, v18) != 0 {
				v91 = v67
			} else {
				v74 = v22
				v78 = v25
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v78) {
					v91 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v82))
				} else {
					v91 = base.B2i32(base.Ui64(v82) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v79, v74)
				}
			}
		} else {
			v30 = int32(0)
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
				v91 = v30
			} else {
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v12, v18)), float64(1e-06)) == int32(0))&base.F64_ne(v12, v18) != 0 {
						v91 = v30
					} else {
						v91 = base.F64_eq(v22, v33) | base.F64_le(base.F64_abs(base.F64_sub(v22, v33)), float64(1e-06))
					}
				} else {
					v65 = int32(1)
					v67 = int32(0)
					if base.B2i32(v65 == v67)|base.F64_ne(v12, v18) != 0 {
						v91 = v67
					} else {
						v74 = v22
						v78 = v25
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v78) {
							v91 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v82))
						} else {
							v91 = base.B2i32(base.Ui64(v82) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v79, v74)
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if base.Ui64(v41&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v91 = int32(0)
		} else {
			v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v74 = v46
			v78 = base.I64_reinterpret_f64(v46) & int64(9223372036854775807)
			v79 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v78) {
				v91 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v82))
			} else {
				v91 = base.B2i32(base.Ui64(v82) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v79, v74)
			}
		}
	}
	return v91 ^ int32(1)
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
