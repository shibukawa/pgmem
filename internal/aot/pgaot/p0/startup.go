package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
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
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
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
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1082 int32
	_ = v1082
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
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
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L324
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L1
	} else {
		goto L320
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L316
	}
L6:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L311
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L306
	}
L8:
	;
	m.G0 = v14 + int32(48)
	return v1142
L9:
	;
	if v24 == int32(-1) {
		v1142 = v20
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
		v1142 = v20
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
		v1142 = v20
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
	v1142 = v20
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
		v1142 = v71
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
	v1142 = v71
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
		v1142 = v98
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
	v1142 = v98
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
		goto L119
	} else {
		goto L120
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
		v1142 = v141
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
	v1142 = v141
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
		v1142 = v168
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
	v1142 = v168
	goto L8
L62:
	;
	m.G0 = v202 + int32(48)
	v1142 = int32(-1)
	goto L8
L63:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v161 == v296 {
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
	v283 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L85
	}
L67:
	;
	v210 = int32(0)
	v211 = v205
	goto L70
L68:
	;
	goto L69
L69:
	;
	v266 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L81
	}
L70:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[2]))
	v225 = v222 + v210<<(uint(int32(7))%32)
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
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+96)) = int32(1)
	v234 = v225 + int32(104)
	if v230 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v246 = v211
	goto L74
L74:
	;
	v249 = v210 + int32(1)
	if v249 < v246+int32(38) {
		v210 = v249
		v211 = v246
		goto L70
	} else {
		goto L80
	}
L75:
	;
	F_s_lock(m, v234, int32(_a_F_ProcessStartupPacket_8), int32(754), int32(_a_F_ProcessStartupPacket_9))
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
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[1]))
	v246 = v245
	goto L74
L80:
	;
	goto L71
L81:
	;
	if v266 == int32(0) {
		goto L62
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+32)) = v196
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_10), v202+int32(32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(798), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L62
L85:
	;
	if v283 == int32(0) {
		goto L62
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_11), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(733), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L62
L89:
	;
	v441 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L115
	}
L90:
	;
	v299 = v225 + int32(16)
	v300 = int32(0)
	if v161 == v300 {
		v405 = v300
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(0)
	goto L89
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+96)) = int32(0)
	if v405 != 0 {
		goto L89
	} else {
		goto L107
	}
L94:
	;
	v305 = v161 & int32(3)
	if base.Ui32(v161) < base.Ui32(int32(4)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v405 = base.B2i32(v384 != int32(0))
	goto L93
L96:
	;
	v359 = v348
	v361 = v350
	v362 = v351
	v364 = v300
	goto L104
L97:
	;
	v348 = v299
	v350 = v198
	v351 = int32(0)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v312 = v299
	v314 = v198
	v315 = int32(0)
	v322 = v4
	goto L100
L100:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+2)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+2)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+3)))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+3)))
	v338 = v315 | (v323 ^ v324) | (v327 ^ v328) | (v331 ^ v332) | (v335 ^ v336)
	v339 = int32(4)
	v340 = v314 + v339
	v342 = v312 + v339
	v344 = v322 + v339
	if v344 != v161&int32(-4) {
		v312 = v342
		v314 = v340
		v315 = v338
		v322 = v344
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if v305 == int32(0) {
		v384 = v338
		goto L95
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v348 = v342
	v350 = v340
	v351 = v338
	goto L96
L104:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	v373 = v362 | (v370 ^ v371)
	v374 = int32(1)
	v379 = v364 + v374
	if v379 != v305 {
		v359 = v359 + v374
		v361 = v361 + v374
		v362 = v373
		v364 = v379
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v384 = v373
	goto L95
L106:
	;
	goto L105
L107:
	;
	v410 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v410 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v196
	F_errmsg_internal(m, int32(_a_F_ProcessStartupPacket_12), v202)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v424 = F_kill(m, int32(0)-v196, int32(2))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(772), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	goto L62
L115:
	;
	if v441 == int32(0) {
		goto L62
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v196
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_13), v202+int32(16))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_8), int32(789), int32(_a_F_ProcessStartupPacket_9))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L62
L119:
	;
	v476 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)) = uint8(v476)
	v479 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[3])))
	if v479 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	if l2|base.B2i32(v125 != int32(806801924)) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L122:
	;
	goto L129
L123:
	;
	v484 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	if v484 == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_14), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(612), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L122
L128:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[4]))
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[5]))
	goto L139
L129:
	;
	v511 = F_secure_write(m, l0, v14+int32(43), int32(1))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	v519 = int32(-1)
	v522 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L134
	}
L131:
	;
	if v511 == int32(1) {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[6]))
	if v516 == int32(27) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	if v522 == int32(0) {
		v1142 = v519
		goto L8
	} else {
		goto L135
	}
L135:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_15), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(621), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v1142 = v519
	goto L8
L139:
	;
	if int32(0) < v538-v540 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)))
	v548 = F_ProcessStartupPacket(m, l0, int32(1), base.B2i32(v545 == int32(83)))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v1142 = v548
	goto L8
L142:
	;
	v555 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)) = uint8(v555)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[3])))
	if v558 != int32(1) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v630 = int32(_a_F_ProcessStartupPacket_16)
	if base.Ui32(v630) <= base.Ui32(v134) {
		goto L165
	} else {
		goto L166
	}
L145:
	;
	goto L152
L146:
	;
	v563 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v563 == int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_17), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(666), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L145
L151:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[4]))
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[5]))
	goto L162
L152:
	;
	v590 = F_secure_write(m, l0, v14+int32(43), int32(1))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L154
	}
L153:
	;
	v598 = int32(-1)
	v601 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L157
	}
L154:
	;
	if v590 == int32(1) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[6]))
	if v595 == int32(27) {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	goto L153
L157:
	;
	if v601 == int32(0) {
		v1142 = v598
		goto L8
	} else {
		goto L158
	}
L158:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_18), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(675), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v1142 = v598
	goto L8
L162:
	;
	if int32(0) < v617-v619 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+43)))
	v627 = F_ProcessStartupPacket(m, l0, base.B2i32(v623 == int32(71)), int32(1))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v1142 = v627
	goto L8
L165:
	;
	v633 = v630
	goto L167
L166:
	;
	v633 = v134
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[7])) = v633
	if base.Ui32(v134-int32(_a_F_ProcessStartupPacket_19)) <= base.Ui32(int32(-65537)) {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	v639 = int32(_a_F_ProcessStartupPacket_20)
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[8]))
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[8])) = v643
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = int32(0)
	v647 = int32(4)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v648 < int32(5) {
		v963 = v647
		v967 = v648
		v969 = v4
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v963 != v967-int32(1) {
		goto L4
	} else {
		goto L268
	}
L170:
	;
	v652 = v647
	v656 = v648
	v658 = v4
	goto L171
L171:
	;
	v662 = v652 + v92
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	if v663 == int32(0) {
		v963 = v652
		v967 = v656
		v969 = v658
		goto L169
	} else {
		goto L173
	}
L172:
	;
	v963 = v959
	v967 = v960
	v969 = v955
	goto L169
L173:
	;
	v666 = F_strlen(m, v662)
	mBase = m.M
	v669 = v666 + v652 + int32(1)
	if v656 <= v669 {
		v963 = v652
		v967 = v656
		v969 = v658
		goto L169
	} else {
		goto L174
	}
L174:
	;
	v671 = v92 + v669
	v672 = int32(_a_F_ProcessStartupPacket_21)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v678 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[10])))
	if base.B2i32(v675 == int32(0))|base.B2i32(v675 != v678) != 0 {
		v696 = v675
		v697 = v678
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v956 = F_strlen(m, v671)
	mBase = m.M
	v959 = v956 + v669 + int32(1)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v959 < v960 {
		v652 = v959
		v656 = v960
		v658 = v955
		goto L171
	} else {
		goto L267
	}
L176:
	;
	if v696-v697 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L177:
	;
	goto L176
L178:
	;
	v681 = v662
	v682 = v672
	goto L179
L179:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+1)))
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+1)))
	if v686 == int32(0) {
		v696 = v686
		v697 = v685
		goto L177
	} else {
		goto L181
	}
L180:
	;
	v696 = v686
	v697 = v685
	goto L177
L181:
	;
	v689 = int32(1)
	if v686 == v685 {
		v681 = v681 + v689
		v682 = v682 + v689
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v701 = F_pstrdup(m, v671)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v704 = int32(_a_F_ProcessStartupPacket_22)
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v710 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[11])))
	if base.B2i32(v707 == int32(0))|base.B2i32(v707 != v710) != 0 {
		v728 = v707
		v729 = v710
		goto L188
	} else {
		goto L189
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v701
	v955 = v658
	goto L175
L187:
	;
	if v728-v729 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	goto L187
L189:
	;
	v713 = v662
	v714 = v704
	goto L190
L190:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714)+1)))
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)))
	if v718 == int32(0) {
		v728 = v718
		v729 = v717
		goto L188
	} else {
		goto L192
	}
L191:
	;
	v728 = v718
	v729 = v717
	goto L188
L192:
	;
	v721 = int32(1)
	if v718 == v717 {
		v713 = v713 + v721
		v714 = v714 + v721
		goto L190
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	v733 = F_pstrdup(m, v671)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v736 = int32(_a_F_ProcessStartupPacket_23)
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v742 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[12])))
	if base.B2i32(v739 == int32(0))|base.B2i32(v739 != v742) != 0 {
		v760 = v739
		v761 = v742
		goto L199
	} else {
		goto L200
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v733
	v955 = v658
	goto L175
L198:
	;
	if v760-v761 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L199:
	;
	goto L198
L200:
	;
	v745 = v662
	v746 = v736
	goto L201
L201:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+1)))
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745)+1)))
	if v750 == int32(0) {
		v760 = v750
		v761 = v749
		goto L199
	} else {
		goto L203
	}
L202:
	;
	v760 = v750
	v761 = v749
	goto L199
L203:
	;
	v753 = int32(1)
	if v750 == v749 {
		v745 = v745 + v753
		v746 = v746 + v753
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v765 = F_pstrdup(m, v671)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v768 = int32(_a_F_ProcessStartupPacket_24)
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v774 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[13])))
	if base.B2i32(v771 == int32(0))|base.B2i32(v771 != v774) != 0 {
		v792 = v771
		v793 = v774
		goto L210
	} else {
		goto L211
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v765
	v955 = v658
	goto L175
L209:
	;
	if v792-v793 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L210:
	;
	goto L209
L211:
	;
	v777 = v662
	v778 = v768
	goto L212
L212:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+1)))
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+1)))
	if v782 == int32(0) {
		v792 = v782
		v793 = v781
		goto L210
	} else {
		goto L214
	}
L213:
	;
	v792 = v782
	v793 = v781
	goto L210
L214:
	;
	v785 = int32(1)
	if v782 == v781 {
		v777 = v777 + v785
		v778 = v778 + v785
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v797 = int32(_a_F_ProcessStartupPacket_21)
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671))))
	v803 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[10])))
	if base.B2i32(v800 == int32(0))|base.B2i32(v800 != v803) != 0 {
		v821 = v800
		v822 = v803
		goto L220
	} else {
		goto L221
	}
L217:
	;
	goto L218
L218:
	;
	v859 = int32(_a_F_ProcessStartupPacket_25)
	goto L238
L219:
	;
	if v821-v822 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L220:
	;
	goto L219
L221:
	;
	v806 = v671
	v807 = v797
	goto L222
L222:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807)+1)))
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806)+1)))
	if v811 == int32(0) {
		v821 = v811
		v822 = v810
		goto L220
	} else {
		goto L224
	}
L223:
	;
	v821 = v811
	v822 = v810
	goto L220
L224:
	;
	v814 = int32(1)
	if v811 == v810 {
		v806 = v806 + v814
		v807 = v807 + v814
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v827 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[14])) = uint8(v827)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[15])) = uint8(v827)
	v955 = v658
	goto L175
L227:
	;
	goto L228
L228:
	;
	v833 = F_strlen(m, v671)
	mBase = m.M
	v834 = F_parse_bool_with_len(m, v671, v833, int32(_a_F_ProcessStartupPacket_26))
	mBase = m.M
	goto L229
L229:
	;
	if v834 != 0 {
		v955 = v658
		goto L175
	} else {
		goto L230
	}
L230:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_ProcessStartupPacket_24)
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_27), v14+int32(32))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errhint(m, int32(_a_F_ProcessStartupPacket_28), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(784), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	if v897-v898 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L238:
	;
	goto L239
L239:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	if v866 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v867 = v662
	v868 = v859
	v869 = int32(5)
	v870 = v866
	goto L244
L241:
	;
	v893 = v859
	v897 = int32(0)
	goto L242
L242:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	goto L236
L243:
	;
	v893 = v888
	v897 = v890
	goto L242
L244:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868))))
	if base.B2i32(v870 != v872)|base.B2i32(v872 == int32(0)) != 0 {
		v888 = v868
		v890 = v870
		goto L243
	} else {
		goto L246
	}
L245:
	;
	v888 = v882
	v890 = int32(0)
	goto L243
L246:
	;
	v878 = v869 - int32(1)
	if v878 == int32(0) {
		v888 = v868
		v890 = v870
		goto L243
	} else {
		goto L247
	}
L247:
	;
	v881 = int32(1)
	v882 = v868 + v881
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+1)))
	if v883 != 0 {
		v867 = v867 + v881
		v868 = v882
		v869 = v878
		v870 = v883
		goto L244
	} else {
		goto L248
	}
L248:
	;
	goto L245
L249:
	;
	v908 = F_pstrdup(m, v662)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v913 = F_pstrdup(m, v662)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	v910 = F_lappend(m, v658, v908)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v955 = v910
	goto L175
L254:
	;
	v915 = F_lappend(m, v912, v913)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v915
	v918 = F_pstrdup(m, v671)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v920 = F_lappend(m, v915, v918)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v920
	v923 = int32(_a_F_ProcessStartupPacket_29)
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v929 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[16])))
	if base.B2i32(v926 == int32(0))|base.B2i32(v926 != v929) != 0 {
		v947 = v926
		v948 = v929
		goto L259
	} else {
		goto L260
	}
L258:
	;
	if v947-v948 != 0 {
		v955 = v658
		goto L175
	} else {
		goto L265
	}
L259:
	;
	goto L258
L260:
	;
	v932 = v662
	v933 = v923
	goto L261
L261:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933)+1)))
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932)+1)))
	if v937 == int32(0) {
		v947 = v937
		v948 = v936
		goto L259
	} else {
		goto L263
	}
L262:
	;
	v947 = v937
	v948 = v936
	goto L259
L263:
	;
	v940 = int32(1)
	if v937 == v936 {
		v932 = v932 + v940
		v933 = v933 + v940
		goto L261
	} else {
		goto L264
	}
L264:
	;
	goto L262
L265:
	;
	v951 = F_pg_clean_ascii(m, v671, int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v951
	v955 = v658
	goto L175
L267:
	;
	goto L172
L268:
	;
	v976 = int32(0)
	if base.B2i32(v969 == v976)&base.B2i32(base.Ui32(v134&int32(_a_F_ProcessStartupPacket_30)) <= base.Ui32(int32(2))) == v976 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v986 = m.G0
	v988 = v986 - int32(16)
	m.G0 = v988
	F_pq_beginmessage(m, v988, int32(118))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v1097 == int32(0) {
		goto L3
	} else {
		goto L286
	}
L272:
	;
	v994 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[7]))
	F_enlargeStringInfo(m, v988, int32(4))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	v1003 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v998+v999))) = base.I32_rotr(v994, int32(24))&v1003 | base.I32_rotr(v994&v1003, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+4)) = v998 + int32(4)
	if v969 != 0 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	F_pq_endmessage(m, v988)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L285
	}
L275:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	F_enlargeStringInfo(m, v988, int32(4))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	F_enlargeStringInfo(m, v988, int32(4))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L284
	}
L278:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	v1023 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v1018+v1019))) = base.I32_rotr(v1014, int32(24))&v1023 | base.I32_rotr(v1014&v1023, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+4)) = v1018 + int32(4)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	if v1034 <= int32(0) {
		goto L274
	} else {
		goto L279
	}
L279:
	;
	v1039 = int32(0)
	goto L280
L280:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v969)+12))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1048+v1039<<(uint(int32(2))%32))))
	F_pq_sendstring(m, v988, v1052)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L282
	}
L281:
	;
	goto L274
L282:
	;
	v1056 = v1039 + int32(1)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	if v1056 < v1057 {
		v1039 = v1056
		goto L280
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	*(*int32)(unsafe.Add(mBase, uint32(v1062+v1063))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v988)+4)) = v1062 + int32(4)
	goto L274
L285:
	;
	m.G0 = v988 + int32(16)
	goto L271
L286:
	;
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097))))
	if v1100 == int32(0) {
		goto L3
	} else {
		goto L287
	}
L287:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	if v1103 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v1109 = F_strlen(m, v1108)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v1109) {
		goto L294
	} else {
		goto L295
	}
L289:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103))))
	if v1104 != 0 {
		v1108 = v1103
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1105 = F_pstrdup(m, v1097)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L293
	}
L292:
	;
	goto L291
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v1105
	v1108 = v1105
	goto L288
L294:
	;
	v1112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1108)+63)) = uint8(v1112)
	goto L296
L295:
	;
	goto L296
L296:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	v1115 = F_strlen(m, v1114)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v1115) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1114)+63)) = uint8(v1118)
	goto L299
L298:
	;
	goto L299
L299:
	;
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[15])))
	if v1125 != 0 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1126 = int32(6)
	goto L302
L301:
	;
	v1126 = int32(1)
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[17])) = v1126
	if v1125 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[8])) = v640
	v1142 = int32(0)
	goto L8
L304:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessStartupPacket[14])))
	if v1131&int32(1) != 0 {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	v1135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1134))) = uint8(v1135)
	goto L303
L306:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_31), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errdetail(m, int32(_a_F_ProcessStartupPacket_32), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(640), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_33), int32(0))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errdetail(m, int32(_a_F_ProcessStartupPacket_32), int32(0))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(694), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(12884901891)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v134 & int32(_a_F_ProcessStartupPacket_30)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(base.Ui32(v134) >> (uint(int32(16)) % 32))
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_34), v14)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(725), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_35), int32(0))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(825), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F_errmsg(m, int32(_a_F_ProcessStartupPacket_36), int32(0))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_ProcessStartupPacket_2), int32(842), int32(_a_F_ProcessStartupPacket_3))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
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
