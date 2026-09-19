package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
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
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
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
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int64
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int64
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int64
	_ = v675
	var v685 int64
	_ = v685
	var v688 int32
	_ = v688
	var v691 int64
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int64
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v909 int32
	_ = v909
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v1010 int64
	_ = v1010
	var v1015 float64
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 float64
	_ = v1019
	var v1026 float64
	_ = v1026
	var v1031 int64
	_ = v1031
	var v1032 int64
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int64
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int64
	_ = v1045
	var v1047 int64
	_ = v1047
	var v1048 int64
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int64
	_ = v1052
	var v1053 int64
	_ = v1053
	var v1057 int64
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1067 int64
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1081 int64
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1130 int32
	_ = v1130
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
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
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L10
	} else {
		goto L237
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L10
	} else {
		goto L234
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
	v154 = int32(0)
	v155 = v148
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
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v122+v154<<(uint(int32(2))%32))))
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
	v188 = v155
	goto L31
L31:
	;
	v190 = v154 + int32(1)
	if v190 < v188 {
		v153 = v187
		v154 = v190
		v155 = v188
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
	return v1130
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
		v1130 = v235
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
	v1130 = v235
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
	v365 = base.AtomicRmwXchg32(m, v362, int32(440), int32(1))
	if v365 != 0 {
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
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v376)+200)) = v377
	v379 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v376)+440)), uint32(v379))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[16])))
	if v383 == int32(1) {
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
	v388 = m.ExcPending
	if v388 != 0 {
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
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v414 = F_LWLockAcquire(m, v410+int32(384), int32(1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L10
	} else {
		goto L91
	}
L84:
	;
	v393 = int32(_a_F_CreateCheckPoint_11)
	goto L86
L85:
	;
	v393 = int32(_a_F_CreateCheckPoint_12)
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v393
	if l0&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v399 = int32(_a_F_CreateCheckPoint_13)
	goto L89
L88:
	;
	v399 = int32(_a_F_CreateCheckPoint_11)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v399
	v402 = v19 + int32(112)
	v405 = F_pg_snprintf(m, v402, int32(128), int32(_a_F_CreateCheckPoint_14), v19)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v407 = F_strlen(m, v402)
	mBase = m.M
	goto L83
L91:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v417)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v417)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v422
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v425+int32(384))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v435 = F_LWLockAcquire(m, v431+int32(_a_F_CreateCheckPoint_15), int32(1))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v439
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v438)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v441
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v444+int32(_a_F_CreateCheckPoint_15))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v454 = F_LWLockAcquire(m, v450+int32(256), int32(1))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v458
	if v85 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v462 + v458
	goto L98
L97:
	;
	goto L98
L98:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v466+int32(256))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v484 = F_LWLockAcquire(m, v480+int32(1664), int32(1))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[17]))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(60)))) = v488
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(76)))) = v492
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v487)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(80)))) = v494
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v497+int32(1664))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	v502 = int32(_a_F_CreateCheckPoint_2)
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	v505 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v504 - v505
	v511 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), v505)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v513 {
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
	F_pfree(m, v511)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L10
	} else {
		goto L111
	}
L106:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L10
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	v534 = int32(_a_F_CreateCheckPoint_16)
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v535))) = int32(134217738)
	F_pg_usleep(m, int32(_a_F_CreateCheckPoint_17))
	mBase = m.M
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = int32(0)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v546 = F_HaveVirtualXIDsDelayingChkpt(m, v511, v544, int32(1))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	if v546 != 0 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	goto L107
L111:
	;
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	F_CheckPointGuts(m, v566, l0)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	v572 = F_GetVirtualXIDsDelayingChkpt(m, v19+int32(12), int32(2))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if int32(0) < v574 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	goto L117
L115:
	;
	goto L116
L116:
	;
	F_pfree(m, v572)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L10
	} else {
		goto L122
	}
L117:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L10
	} else {
		goto L119
	}
L118:
	;
	goto L116
L119:
	;
	v595 = int32(_a_F_CreateCheckPoint_16)
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v596))) = int32(134217737)
	F_pg_usleep(m, int32(_a_F_CreateCheckPoint_17))
	mBase = m.M
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v602))) = int32(0)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v607 = F_HaveVirtualXIDsDelayingChkpt(m, v572, v605, int32(2))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	if v607 != 0 {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	if v85 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v633 = int32(_a_F_CreateCheckPoint_2)
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v635 + int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L10
	} else {
		goto L127
	}
L124:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	if v628 <= int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v631 = F_LogStandbySnapshot(m)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	F_XLogRegisterData(m, v19+int32(24), int32(88))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	v646 = int32(0)
	if v85 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v649 = v646
	goto L131
L130:
	;
	v649 = int32(16)
	goto L131
L131:
	;
	v650 = F_XLogInsert(m, v646, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	F_XLogFlush(m, v650)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	if v85 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12])) = v280
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v658 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	if v656 != v658 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v661)+40))
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v668 = F_LWLockAcquire(m, v664+int32(1152), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L10
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	if v85 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+16)) = int32(1)
	goto L141
L140:
	;
	goto L141
L141:
	;
	v675 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	*(*int64)(unsafe.Add(mBase, uint32(v671)+32)) = v675
	base.MemoryCopy(m, v671+int32(40), v19+int32(24), int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+144)) = int32(0)
	v685 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v671)+136)) = v685
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v691 = base.AtomicRmwOr64(m, v688, int32(240), v685)
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v693)+128)) = v691
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[7]))
	F_update_controlfile(m, v696, v693)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v700+int32(1152))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L10
	} else {
		goto L143
	}
L143:
	;
	v706 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v709 = base.AtomicRmwXchg32(m, v706, int32(440), int32(1))
	if v709 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	F_s_lock(m, v711+int32(440), int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_18), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L10
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v721 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v720)+208)) = v721
	v723 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v720)+440)), uint32(v723))
	v726 = int32(_a_F_CreateCheckPoint_2)
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v728 - int32(1)
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	if v733 == v723 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L146
L148:
	;
	v806 = int32(0)
	v808 = m.G0
	v810 = v808 - int32(1040)
	m.G0 = v810
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	if v813 != 0 {
		goto L169
	} else {
		goto L170
	}
L149:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v741 = F_LWLockAcquire(m, v737+int32(_a_F_CreateCheckPoint_19), int32(1))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+20))
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v747+int32(_a_F_CreateCheckPoint_19))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	if v745 == int32(-1) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[9]))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v761 = v756 + v745*int32(640) + int32(20)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	if v762 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L148
L154:
	;
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = int32(1)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v761)+4))
	if v765 == int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	if v768 == int32(0) {
		goto L154
	} else {
		goto L157
	}
L157:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[22]))
	if v772 == v768 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v774 = m.G0
	v776 = v774 - int32(16)
	m.G0 = v776
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	if v779 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L160
L160:
	;
	v802 = F_pgmem_kill(m, v768, int32(23))
	mBase = m.M
	goto L154
L161:
	;
	m.G0 = v776 + int32(16)
	goto L153
L162:
	;
	v782 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+15)) = uint8(v782)
	goto L163
L163:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[24]))
	v790 = F_write(m, v786, v776+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v790 {
		goto L161
	} else {
		goto L165
	}
L164:
	;
	goto L161
L165:
	;
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[25]))
	if v794 == int32(27) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21])) = v992
	m.G0 = v810 + int32(1040)
	v1010 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	if v662 != int64(0) {
		goto L203
	} else {
		goto L204
	}
L168:
	;
	v930 = int32(0)
	v932 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v932)+12))
	v936 = (v837 - v933) >> (uint(int32(2)) % 32)
	if v930 < v936 {
		goto L195
	} else {
		goto L196
	}
L169:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if int32(0) < v814 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v927 = int32(0)
	goto L171
L171:
	;
	F_list_free_deep(m, v927)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L10
	} else {
		goto L194
	}
L172:
	;
	v818 = v806
	v820 = int32(10)
	goto L175
L173:
	;
	goto L174
L174:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v927 = v909
	goto L171
L175:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v813)+12))
	v837 = v834 + v818<<(uint(int32(2))%32)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838)+26)))
	if v839 != 0 {
		v886 = v820
		goto L177
	} else {
		goto L178
	}
L176:
	;
	goto L174
L177:
	;
	v889 = v818 + int32(1)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if v889 < v890 {
		v818 = v889
		v820 = v886
		goto L175
	} else {
		goto L193
	}
L178:
	;
	v840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838)+24)))
	v842 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])))
	if v840 == v842 {
		goto L168
	} else {
		goto L179
	}
L179:
	;
	v845 = v810 + int32(16)
	v846 = int32(*(*int16)(unsafe.Add(mBase, uint32(v838))))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v846*int32(12))+uint32(_c_F_CreateCheckPoint[26])))
	v852 = m.T0[v851].(func(*base.Module, int32, int32) int32)(m, v838, v845)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L10
	} else {
		goto L181
	}
L180:
	;
	v877 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v838)+26)) = uint8(v877)
	if v877 < v820 {
		goto L189
	} else {
		goto L190
	}
L181:
	;
	if int32(0) <= v852 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[25]))
	if v857 == int32(44) {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v862 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L10
	} else {
		goto L184
	}
L184:
	;
	if v862 == int32(0) {
		goto L180
	} else {
		goto L185
	}
L185:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v845
	F_errmsg(m, int32(_a_F_CreateCheckPoint_20), v810)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_21), int32(243), int32(_a_F_CreateCheckPoint_22))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	goto L180
L189:
	;
	v886 = v820 - int32(1)
	goto L177
L190:
	;
	goto L191
L191:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L10
	} else {
		goto L192
	}
L192:
	;
	v886 = int32(10)
	goto L177
L193:
	;
	goto L176
L194:
	;
	v992 = v806
	goto L167
L195:
	;
	v939 = v930
	goto L198
L196:
	;
	v985 = v932
	goto L197
L197:
	;
	v986 = F_list_delete_first_n(m, v985, v936)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L10
	} else {
		goto L202
	}
L198:
	;
	v956 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)+12))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v957+v939<<(uint(int32(2))%32))))
	F_pfree(m, v961)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L10
	} else {
		goto L200
	}
L199:
	;
	v968 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	v985 = v968
	goto L197
L200:
	;
	v965 = v939 + int32(1)
	if v965 != v936 {
		v939 = v965
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v992 = v986
	goto L167
L203:
	;
	v1015 = base.F64_convert_i64_u(v1010 - v662)
	*(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[27])) = v1015
	v1017 = int32(_a_F_CreateCheckPoint_23)
	v1019 = *(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[28]))
	if base.F64_gt(v1015, v1019) != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v1031 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14])))
	v1032 = base.I64_div_u_s(v1010, v1031)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1032
	v1035 = v19 + int32(16)
	F_KeepLogSeg(m, v650, v1035)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L10
	} else {
		goto L209
	}
L206:
	;
	v1026 = v1015
	goto L208
L207:
	;
	v1026 = base.F64_add(base.F64_mul(v1019, float64(0.9)), base.F64_mul(v1015, float64(0.1)))
	goto L208
L208:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[28])) = v1026
	goto L205
L209:
	;
	v1039 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1040 = int32(0)
	v1042 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v1039, v1040, v1040)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L10
	} else {
		goto L210
	}
L210:
	;
	if v1042 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1045 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	v1047 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14])))
	v1048 = base.I64_div_u_s(v1045, v1047)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1048
	F_KeepLogSeg(m, v650, v1035)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L10
	} else {
		goto L214
	}
L212:
	;
	v1053 = v1039
	goto L213
L213:
	;
	v1057 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15]))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	F_RemoveOldXlogFiles(m, v1053-int64(1), v1057, v650, v1058)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L10
	} else {
		goto L215
	}
L214:
	;
	v1052 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v1053 = v1052
	goto L213
L215:
	;
	if v85 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])))
	if v1106 == int32(1) {
		goto L226
	} else {
		goto L227
	}
L217:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+320)))
	if v1063 != int32(1) {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v1067 = v650 - int64(1)
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14]))
	if base.Ui64(v1067&base.I64_extend_i32_s(v1069-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v1069), float64(0.75))))) {
		goto L216
	} else {
		goto L219
	}
L219:
	;
	v1081 = base.I64_div_u_s(v1067, base.I64_extend_i32_s(v1069))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v1089 = F_XLogFileInitInternal(m, v1081+int64(1), v1084, v19+int32(1151), v19+int32(112))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L10
	} else {
		goto L220
	}
L220:
	;
	if int32(0) <= v1089 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1093 = F_close(m, v1089)
	mBase = m.M
	goto L223
L222:
	;
	goto L223
L223:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1151)))
	if v1094 != int32(1) {
		goto L216
	} else {
		goto L224
	}
L224:
	;
	v1097 = int32(_a_F_CreateCheckPoint_24)
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[29]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[29])) = v1099 + int32(1)
	goto L216
L225:
	;
	F_LogCheckpointEnd(m, int32(0))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L10
	} else {
		goto L232
	}
L226:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+316))
	v1114 = base.B2i32(v1112 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])) = uint8(v1114)
	if v1112 != int32(2) {
		goto L225
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1117 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L10
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	F_TruncateSUBTRANS(m, v1117)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	goto L225
L232:
	;
	v1125 = int32(1)
	if v85 == int32(0) {
		v1130 = v1125
		goto L42
	} else {
		goto L233
	}
L233:
	;
	v1130 = v1125
	goto L42
L234:
	;
	F_errmsg_internal(m, int32(_a_F_CreateCheckPoint_25), int32(0))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_26), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L10
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_errmsg(m, int32(_a_F_CreateCheckPoint_27), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L10
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_28), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L10
	} else {
		goto L239
	}
L239:
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
