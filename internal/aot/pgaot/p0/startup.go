package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
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
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1086 int32
	_ = v1086
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	F_pq_startmsgread(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(-1)
	v22 = v14 + int32(44)
	v24 = F_pq_getbytes(m, v22, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L323
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L319
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L315
	}
L6:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L310
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L305
	}
L8:
	;
	m.G0 = v14 + int32(48)
	return v1146
L9:
	;
	if v24 == int32(-1) {
		v1146 = v20
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v31 = F_pq_getbytes(m, v22|int32(1), int32(3))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v31 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l1|l2 != 0 {
		v1146 = v20
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v55 = int32(16711935)
	v63 = base.I32_rotr(v54&v55, int32(8)) | base.I32_rotr(v54, int32(24))&v55
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v63 - int32(4)
	if base.Ui32(v63-int32(_a_F_ProcessStartupPacket_0)) <= base.Ui32(int32(-9998)) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v38 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v38 == int32(0) {
		v1146 = v20
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_1), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(540), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v1146 = v20
	goto L8
L21:
	;
	v71 = int32(-1)
	v74 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v92 = F_palloc(m, v63-int32(3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L24:
	;
	if v74 == int32(0) {
		v1146 = v71
		goto L8
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_4), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(552), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v1146 = v71
	goto L8
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v94))) = uint8(v96)
	v98 = int32(-1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v100 = F_pq_getbytes(m, v92, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v100 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v106 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[0])) = uint8(v123)
	goto L39
L34:
	;
	if v106 == int32(0) {
		v1146 = v98
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_1), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(568), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v1146 = v98
	goto L8
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v126 = int32(16711935)
	v134 = base.I32_rotr(v125&v126, int32(8)) | base.I32_rotr(v125, int32(24))&v126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v134
	if v125 == int32(773247492) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if base.Ui32(v138) <= base.Ui32(int32(7)) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if l1|base.B2i32(v125 != int32(790024708)) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L43:
	;
	v141 = int32(-1)
	v144 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v161 = v138 - int32(8)
	if v161 < int32(257) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	if v144 == int32(0) {
		v1146 = v141
		goto L8
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_5), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(896), int32(_a_F_ProcessStartupPacket_6))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v1146 = v141
	goto L8
L51:
	;
	v165 = v161
	goto L53
L52:
	;
	v165 = int32(0)
	goto L53
L53:
	;
	if v165 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v168 = int32(-1)
	v171 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v188 = int32(16711935)
	v190 = int32(8)
	v196 = base.I32_rotr(v187&v188, v190) | base.I32_rotr(v187, int32(24))&v188
	v198 = v92 + v190
	v200 = m.G0
	v202 = v200 - int32(48)
	m.G0 = v202
	if v196 != 0 {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	if v171 == int32(0) {
		v1146 = v168
		goto L8
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_7), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(904), int32(_a_F_ProcessStartupPacket_6))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v1146 = v168
	goto L8
L62:
	;
	m.G0 = v202 + int32(48)
	v1146 = int32(-1)
	goto L8
L63:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v161 == v299 {
		goto L90
	} else {
		goto L91
	}
L64:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[1]))
	if int32(0) < v205+int32(38) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v286 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L85
	}
L67:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[2]))
	v212 = v211
	v214 = v205
	v218 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	v269 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L81
	}
L70:
	;
	v225 = v212 + v218<<(uint(int32(7))%32)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	if v196 == v226 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v229 = v225 + int32(8)
	v231 = v225 + int32(104)
	v234 = base.AtomicRmwXchg32(m, v229, int32(96), int32(1))
	if v234 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v249 = v212
	v250 = v214
	goto L74
L74:
	;
	v252 = v218 + int32(1)
	if v252 < v250+int32(38) {
		v212 = v249
		v214 = v250
		v218 = v252
		goto L70
	} else {
		goto L80
	}
L75:
	;
	F_s_lock(m, v231, int32(_a_F_ProcessStartupPacket_8), int32(754), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v240 == v196 {
		goto L63
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	v242 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v231))), uint32(v242))
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[1]))
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[2]))
	v249 = v248
	v250 = v246
	goto L74
L80:
	;
	goto L71
L81:
	;
	if v269 == int32(0) {
		goto L62
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+32)) = v196
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_10), v202+int32(32))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(798), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L62
L85:
	;
	if v286 == int32(0) {
		goto L62
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_11), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(733), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L62
L89:
	;
	v445 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L114
	}
L90:
	;
	v302 = v225 + int32(16)
	v303 = int32(0)
	if v161 == v303 {
		v408 = v303
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v429 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v231))), uint32(v429))
	goto L89
L93:
	;
	v409 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v229)+96)), uint32(v409))
	if v408 != 0 {
		goto L89
	} else {
		goto L107
	}
L94:
	;
	v308 = v161 & int32(3)
	if base.Ui32(v161) < base.Ui32(int32(4)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v408 = base.B2i32(v387 != int32(0))
	goto L93
L96:
	;
	v362 = v351
	v363 = v352
	v365 = v354
	v368 = v303
	goto L104
L97:
	;
	v351 = v302
	v352 = v198
	v354 = int32(0)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v315 = v302
	v316 = v198
	v318 = int32(0)
	v325 = v4
	goto L100
L100:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+2)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+2)))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+3)))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+3)))
	v341 = v318 | (v326 ^ v327) | (v330 ^ v331) | (v334 ^ v335) | (v338 ^ v339)
	v342 = int32(4)
	v343 = v316 + v342
	v345 = v315 + v342
	v347 = v325 + v342
	if v347 != v161&int32(-4) {
		v315 = v345
		v316 = v343
		v318 = v341
		v325 = v347
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if v308 == int32(0) {
		v387 = v341
		goto L95
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v351 = v345
	v352 = v343
	v354 = v341
	goto L96
L104:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	v376 = v365 | (v373 ^ v374)
	v377 = int32(1)
	v382 = v368 + v377
	if v382 != v308 {
		v362 = v362 + v377
		v363 = v363 + v377
		v365 = v376
		v368 = v382
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v387 = v376
	goto L95
L106:
	;
	goto L105
L107:
	;
	v414 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v414 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v196
	F_errmsg_internal(m, int32(_a_F_ProcessStartupPacket_12), v202)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v428 = F_pgmem_kill(m, int32(0)-v196, int32(2))
	mBase = m.M
	goto L62
L112:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(772), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	if v445 == int32(0) {
		goto L62
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v196
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_13), v202+int32(16))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(789), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L62
L118:
	;
	v480 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)) = uint8(v480)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[3])))
	if v483 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	if l2|base.B2i32(v125 != int32(806801924)) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L121:
	;
	goto L128
L122:
	;
	v488 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v488 == int32(0) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_14), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(612), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L121
L127:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[4]))
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[5]))
	goto L138
L128:
	;
	v515 = F_secure_write(m, l0, v14+int32(43), int32(1))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L130
	}
L129:
	;
	v523 = int32(-1)
	v526 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L133
	}
L130:
	;
	if v515 == int32(1) {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[6]))
	if v520 == int32(27) {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	if v526 == int32(0) {
		v1146 = v523
		goto L8
	} else {
		goto L134
	}
L134:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_15), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(621), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v1146 = v523
	goto L8
L138:
	;
	if int32(0) < v542-v544 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)))
	v552 = F_ProcessStartupPacket(m, l0, int32(1), base.B2i32(v549 == int32(83)))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v1146 = v552
	goto L8
L141:
	;
	v559 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)) = uint8(v559)
	v562 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[3])))
	if v562 != int32(1) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L143
L143:
	;
	v634 = int32(_a_F_ProcessStartupPacket_16)
	if base.Ui32(v634) <= base.Ui32(v134) {
		goto L164
	} else {
		goto L165
	}
L144:
	;
	goto L151
L145:
	;
	v567 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	if v567 == int32(0) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_17), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(666), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	goto L144
L150:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[4]))
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[5]))
	goto L161
L151:
	;
	v594 = F_secure_write(m, l0, v14+int32(43), int32(1))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L153
	}
L152:
	;
	v602 = int32(-1)
	v605 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L156
	}
L153:
	;
	if v594 == int32(1) {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[6]))
	if v599 == int32(27) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	if v605 == int32(0) {
		v1146 = v602
		goto L8
	} else {
		goto L157
	}
L157:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_18), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(675), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v1146 = v602
	goto L8
L161:
	;
	if int32(0) < v621-v623 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)))
	v631 = F_ProcessStartupPacket(m, l0, base.B2i32(v627 == int32(71)), int32(1))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v1146 = v631
	goto L8
L164:
	;
	v637 = v634
	goto L166
L165:
	;
	v637 = v134
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[7])) = v637
	if base.Ui32(v134-int32(_a_F_ProcessStartupPacket_19)) <= base.Ui32(int32(-65537)) {
		goto L5
	} else {
		goto L167
	}
L167:
	;
	v643 = int32(_a_F_ProcessStartupPacket_20)
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[8]))
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[8])) = v647
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = int32(0)
	v651 = int32(4)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v652 < int32(5) {
		v967 = v651
		v972 = v652
		v973 = v4
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if v967 != v972-int32(1) {
		goto L4
	} else {
		goto L267
	}
L169:
	;
	v656 = v651
	v661 = v652
	v662 = v4
	goto L170
L170:
	;
	v666 = v656 + v92
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	if v667 == int32(0) {
		v967 = v656
		v972 = v661
		v973 = v662
		goto L168
	} else {
		goto L172
	}
L171:
	;
	v967 = v963
	v972 = v964
	v973 = v959
	goto L168
L172:
	;
	v670 = F_strlen(m, v666)
	mBase = m.M
	v673 = v670 + v656 + int32(1)
	if v661 <= v673 {
		v967 = v656
		v972 = v661
		v973 = v662
		goto L168
	} else {
		goto L173
	}
L173:
	;
	v675 = v92 + v673
	v676 = int32(_a_F_ProcessStartupPacket_21)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[10])))
	if base.B2i32(v679 == int32(0))|base.B2i32(v679 != v682) != 0 {
		v700 = v679
		v701 = v682
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v960 = F_strlen(m, v675)
	mBase = m.M
	v963 = v960 + v673 + int32(1)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v963 < v964 {
		v656 = v963
		v661 = v964
		v662 = v959
		goto L170
	} else {
		goto L266
	}
L175:
	;
	if v700-v701 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L176:
	;
	goto L175
L177:
	;
	v685 = v666
	v686 = v676
	goto L178
L178:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+1)))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685)+1)))
	if v690 == int32(0) {
		v700 = v690
		v701 = v689
		goto L176
	} else {
		goto L180
	}
L179:
	;
	v700 = v690
	v701 = v689
	goto L176
L180:
	;
	v693 = int32(1)
	if v690 == v689 {
		v685 = v685 + v693
		v686 = v686 + v693
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v705 = F_pstrdup(m, v675)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v708 = int32(_a_F_ProcessStartupPacket_22)
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[11])))
	if base.B2i32(v711 == int32(0))|base.B2i32(v711 != v714) != 0 {
		v732 = v711
		v733 = v714
		goto L187
	} else {
		goto L188
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v705
	v959 = v662
	goto L174
L186:
	;
	if v732-v733 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L187:
	;
	goto L186
L188:
	;
	v717 = v666
	v718 = v708
	goto L189
L189:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718)+1)))
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717)+1)))
	if v722 == int32(0) {
		v732 = v722
		v733 = v721
		goto L187
	} else {
		goto L191
	}
L190:
	;
	v732 = v722
	v733 = v721
	goto L187
L191:
	;
	v725 = int32(1)
	if v722 == v721 {
		v717 = v717 + v725
		v718 = v718 + v725
		goto L189
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	v737 = F_pstrdup(m, v675)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v740 = int32(_a_F_ProcessStartupPacket_23)
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v746 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[12])))
	if base.B2i32(v743 == int32(0))|base.B2i32(v743 != v746) != 0 {
		v764 = v743
		v765 = v746
		goto L198
	} else {
		goto L199
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v737
	v959 = v662
	goto L174
L197:
	;
	if v764-v765 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L198:
	;
	goto L197
L199:
	;
	v749 = v666
	v750 = v740
	goto L200
L200:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+1)))
	if v754 == int32(0) {
		v764 = v754
		v765 = v753
		goto L198
	} else {
		goto L202
	}
L201:
	;
	v764 = v754
	v765 = v753
	goto L198
L202:
	;
	v757 = int32(1)
	if v754 == v753 {
		v749 = v749 + v757
		v750 = v750 + v757
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v769 = F_pstrdup(m, v675)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v772 = int32(_a_F_ProcessStartupPacket_24)
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v778 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[13])))
	if base.B2i32(v775 == int32(0))|base.B2i32(v775 != v778) != 0 {
		v796 = v775
		v797 = v778
		goto L209
	} else {
		goto L210
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v769
	v959 = v662
	goto L174
L208:
	;
	if v796-v797 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L209:
	;
	goto L208
L210:
	;
	v781 = v666
	v782 = v772
	goto L211
L211:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782)+1)))
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+1)))
	if v786 == int32(0) {
		v796 = v786
		v797 = v785
		goto L209
	} else {
		goto L213
	}
L212:
	;
	v796 = v786
	v797 = v785
	goto L209
L213:
	;
	v789 = int32(1)
	if v786 == v785 {
		v781 = v781 + v789
		v782 = v782 + v789
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v801 = int32(_a_F_ProcessStartupPacket_21)
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	v807 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[10])))
	if base.B2i32(v804 == int32(0))|base.B2i32(v804 != v807) != 0 {
		v825 = v804
		v826 = v807
		goto L219
	} else {
		goto L220
	}
L216:
	;
	goto L217
L217:
	;
	v863 = int32(_a_F_ProcessStartupPacket_25)
	goto L237
L218:
	;
	if v825-v826 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L219:
	;
	goto L218
L220:
	;
	v810 = v675
	v811 = v801
	goto L221
L221:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811)+1)))
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810)+1)))
	if v815 == int32(0) {
		v825 = v815
		v826 = v814
		goto L219
	} else {
		goto L223
	}
L222:
	;
	v825 = v815
	v826 = v814
	goto L219
L223:
	;
	v818 = int32(1)
	if v815 == v814 {
		v810 = v810 + v818
		v811 = v811 + v818
		goto L221
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	v831 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[14])) = uint8(v831)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[15])) = uint8(v831)
	v959 = v662
	goto L174
L226:
	;
	goto L227
L227:
	;
	v837 = F_strlen(m, v675)
	mBase = m.M
	v838 = F_parse_bool_with_len(m, v675, v837, int32(_a_F_ProcessStartupPacket_26))
	mBase = m.M
	goto L228
L228:
	;
	if v838 != 0 {
		v959 = v662
		goto L174
	} else {
		goto L229
	}
L229:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_ProcessStartupPacket_24)
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_27), v14+int32(32))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errhint(m, int32(_a_F_ProcessStartupPacket_28), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(784), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	if v901-v902 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L237:
	;
	goto L238
L238:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	if v870 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v871 = v666
	v872 = v863
	v873 = int32(5)
	v874 = v870
	goto L243
L240:
	;
	v897 = v863
	v901 = int32(0)
	goto L241
L241:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897))))
	goto L235
L242:
	;
	v897 = v892
	v901 = v894
	goto L241
L243:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872))))
	if base.B2i32(v874 != v876)|base.B2i32(v876 == int32(0)) != 0 {
		v892 = v872
		v894 = v874
		goto L242
	} else {
		goto L245
	}
L244:
	;
	v892 = v886
	v894 = int32(0)
	goto L242
L245:
	;
	v882 = v873 - int32(1)
	if v882 == int32(0) {
		v892 = v872
		v894 = v874
		goto L242
	} else {
		goto L246
	}
L246:
	;
	v885 = int32(1)
	v886 = v872 + v885
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871)+1)))
	if v887 != 0 {
		v871 = v871 + v885
		v872 = v886
		v873 = v882
		v874 = v887
		goto L243
	} else {
		goto L247
	}
L247:
	;
	goto L244
L248:
	;
	v912 = F_pstrdup(m, v666)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v917 = F_pstrdup(m, v666)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	v914 = F_lappend(m, v662, v912)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v959 = v914
	goto L174
L253:
	;
	v919 = F_lappend(m, v916, v917)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v919
	v922 = F_pstrdup(m, v675)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v924 = F_lappend(m, v919, v922)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v924
	v927 = int32(_a_F_ProcessStartupPacket_29)
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v933 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[16])))
	if base.B2i32(v930 == int32(0))|base.B2i32(v930 != v933) != 0 {
		v951 = v930
		v952 = v933
		goto L258
	} else {
		goto L259
	}
L257:
	;
	if v951-v952 != 0 {
		v959 = v662
		goto L174
	} else {
		goto L264
	}
L258:
	;
	goto L257
L259:
	;
	v936 = v666
	v937 = v927
	goto L260
L260:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)))
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936)+1)))
	if v941 == int32(0) {
		v951 = v941
		v952 = v940
		goto L258
	} else {
		goto L262
	}
L261:
	;
	v951 = v941
	v952 = v940
	goto L258
L262:
	;
	v944 = int32(1)
	if v941 == v940 {
		v936 = v936 + v944
		v937 = v937 + v944
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v955 = F_pg_clean_ascii(m, v675, int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v955
	v959 = v662
	goto L174
L266:
	;
	goto L171
L267:
	;
	v980 = int32(0)
	if base.B2i32(v973 == v980)&base.B2i32(base.Ui32(v134&int32(_a_F_ProcessStartupPacket_30)) <= base.Ui32(int32(2))) == v980 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v990 = m.G0
	v992 = v990 - int32(16)
	m.G0 = v992
	F_pq_beginmessage(m, v992, int32(118))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v1101 == int32(0) {
		goto L3
	} else {
		goto L285
	}
L271:
	;
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[7]))
	F_enlargeStringInfo(m, v992, int32(4))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v992)+4))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	v1007 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v1002+v1003))) = base.I32_rotr(v998, int32(24))&v1007 | base.I32_rotr(v998&v1007, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v992)+4)) = v1002 + int32(4)
	if v973 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	F_pq_endmessage(m, v992)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L284
	}
L274:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v973)+4))
	F_enlargeStringInfo(m, v992, int32(4))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	F_enlargeStringInfo(m, v992, int32(4))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L283
	}
L277:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v992)+4))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	v1027 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v1022+v1023))) = base.I32_rotr(v1018, int32(24))&v1027 | base.I32_rotr(v1018&v1027, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v992)+4)) = v1022 + int32(4)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v973)+4))
	if v1038 <= int32(0) {
		goto L273
	} else {
		goto L278
	}
L278:
	;
	v1043 = int32(0)
	goto L279
L279:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v973)+12))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1052+v1043<<(uint(int32(2))%32))))
	F_pq_sendstring(m, v992, v1056)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L281
	}
L280:
	;
	goto L273
L281:
	;
	v1060 = v1043 + int32(1)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v973)+4))
	if v1060 < v1061 {
		v1043 = v1060
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v992)+4))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	*(*int32)(unsafe.Add(mBase, uint32(v1066+v1067))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v992)+4)) = v1066 + int32(4)
	goto L273
L284:
	;
	m.G0 = v992 + int32(16)
	goto L270
L285:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101))))
	if v1104 == int32(0) {
		goto L3
	} else {
		goto L286
	}
L286:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	if v1107 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v1113 = F_strlen(m, v1112)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v1113) {
		goto L293
	} else {
		goto L294
	}
L288:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107))))
	if v1108 != 0 {
		v1112 = v1107
		goto L287
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v1109 = F_pstrdup(m, v1101)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L292
	}
L291:
	;
	goto L290
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v1109
	v1112 = v1109
	goto L287
L293:
	;
	v1116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1112)+63)) = uint8(v1116)
	goto L295
L294:
	;
	goto L295
L295:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	v1119 = F_strlen(m, v1118)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v1119) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1118)+63)) = uint8(v1122)
	goto L298
L297:
	;
	goto L298
L298:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[15])))
	if v1129 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1130 = int32(6)
	goto L301
L300:
	;
	v1130 = int32(1)
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[17])) = v1130
	if v1129 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[8])) = v644
	v1146 = int32(0)
	goto L8
L303:
	;
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[14])))
	if v1135&int32(1) != 0 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	v1139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1138))) = uint8(v1139)
	goto L302
L305:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_31), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errdetail(m, int32(_a_F_ProcessStartupPacket_32), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(640), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_33), int32(0))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errdetail(m, int32(_a_F_ProcessStartupPacket_32), int32(0))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(694), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(12884901891)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v134 & int32(_a_F_ProcessStartupPacket_30)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(base.Ui32(v134) >> (uint(int32(16)) % 32))
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_34), v14)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(725), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_35), int32(0))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(825), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_36), int32(0))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(842), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_has_startup_progress_timeout_expired(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v31 int64
	_ = v31
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_has_startup_progress_timeout_expired[0]))
	if v11 != 0 {
		v15 = m.G0
		v16 = int32(16)
		v17 = v15 - v16
		m.G0 = v17
		F_gettimeofday(m, v17)
		mBase = m.M
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
		m.G0 = v17 + v16
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_has_startup_progress_timeout_expired[1]))
		v38 = v21 + v20*int64(1000000) - int64(946684800000000) - v31
		if v38 <= int64(0) {
			v50 = int32(0)
			v51 = int32(0)
		} else {
			v42 = int64(1000000)
			v43 = base.I64_div_u_s(v38, v42)
			v50 = base.I32_wrap_i64(v43)
			v51 = base.I32_wrap_i64(v38 - v43*v42)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12)))) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v51
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v54
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
		*(*int32)(unsafe.Add(mBase, _c_F_has_startup_progress_timeout_expired[0])) = int32(0)
	} else {
	}
	m.G0 = v8 + int32(16)
	return base.B2i32(v11 != int32(0))
}
