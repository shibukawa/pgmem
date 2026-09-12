package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
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
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
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
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1098 int32
	_ = v1098
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	v12 = m.G0
	v14 = v12 + int32(-64)
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
	v24 = F_pq_getbytes(m, v12+int32(-20), int32(1))
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
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L334
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L330
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L326
	}
L6:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L321
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L316
	}
L8:
	;
	m.G0 = v14 - int32(-64)
	return v1153
L9:
	;
	if v24 == int32(-1) {
		v1153 = v20
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = F_pq_getbytes(m, v12+int32(-20)|int32(1), int32(3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v33 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l1 != 0 {
		v1153 = v20
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v56 = int32(24)
	v58 = int32(65280)
	v60 = int32(8)
	v70 = v55<<(uint(v56)%32) | v55&v58<<(uint(v60)%32) | (int32(base.Ui32(v55)>>(uint(v60)%32))&v58 | int32(base.Ui32(v55)>>(uint(v56)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v70 - int32(4)
	if base.Ui32(v70-int32(10005)) <= base.Ui32(int32(-9998)) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	if l2 != 0 {
		v1153 = v20
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v39 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v39 == int32(0) {
		v1153 = v20
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(107970), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(495925), int32(540), int32(108111))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v1153 = v20
	goto L8
L22:
	;
	v78 = int32(-1)
	v81 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v99 = F_palloc(m, v70-int32(3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	if v81 == int32(0) {
		v1153 = v78
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(107937), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(495925), int32(552), int32(108111))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v1153 = v78
	goto L8
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99+v101))) = uint8(v103)
	v105 = int32(-1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v107 = F_pq_getbytes(m, v99, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v107 == int32(-1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v113 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[646])) = uint8(v130)
	goto L40
L35:
	;
	if v113 == int32(0) {
		v1153 = v105
		goto L8
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(107970), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(495925), int32(568), int32(108111))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v1153 = v105
	goto L8
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v133 = int32(24)
	v135 = int32(65280)
	v137 = int32(8)
	v147 = v132<<(uint(v133)%32) | v132&v135<<(uint(v137)%32) | (int32(base.Ui32(v132)>>(uint(v137)%32))&v135 | int32(base.Ui32(v132)>>(uint(v133)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v147
	if v132 != int32(806801924) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v640 = int32(0)
	v642 = int32(196610)
	if base.Ui32(v642) <= base.Ui32(v147) {
		goto L168
	} else {
		goto L169
	}
L42:
	;
	if v132 != int32(790024708) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if l2 != 0 {
		goto L41
	} else {
		goto L147
	}
L45:
	;
	if v132 != int32(773247492) {
		goto L41
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l1 != 0 {
		goto L41
	} else {
		goto L126
	}
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if base.Ui32(v155) <= base.Ui32(int32(7)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v158 = int32(-1)
	v161 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v178 = v155 - int32(8)
	if v178 < int32(257) {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	if v161 == int32(0) {
		v1153 = v158
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(107845), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(495925), int32(896), int32(108057))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v1153 = v158
	goto L8
L57:
	;
	v182 = v178
	goto L59
L58:
	;
	v182 = int32(0)
	goto L59
L59:
	;
	if v182 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v185 = int32(-1)
	v188 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v205 = int32(24)
	v207 = int32(65280)
	v209 = int32(8)
	v219 = v204<<(uint(v205)%32) | v204&v207<<(uint(v209)%32) | (int32(base.Ui32(v204)>>(uint(v209)%32))&v207 | int32(base.Ui32(v204)>>(uint(v205)%32)))
	v221 = v99 + v209
	v223 = m.G0
	v225 = v223 - int32(48)
	m.G0 = v225
	if v219 != 0 {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	if v188 == int32(0) {
		v1153 = v185
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(107791), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(495925), int32(904), int32(108057))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v1153 = v185
	goto L8
L68:
	;
	m.G0 = v225 + int32(48)
	v1153 = int32(-1)
	goto L8
L69:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v178 == v319 {
		goto L96
	} else {
		goto L97
	}
L70:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	if int32(0) < v228+int32(38) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v306 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L91
	}
L73:
	;
	v233 = v228
	v235 = int32(0)
	goto L76
L74:
	;
	goto L75
L75:
	;
	v289 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L87
	}
L76:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	v248 = v245 + v235<<(uint(int32(7))%32)
	v250 = v248 + int32(8)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v219 == v251 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v250)+96)) = int32(1)
	v257 = v248 + int32(104)
	if v253 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v269 = v233
	goto L80
L80:
	;
	v272 = v235 + int32(1)
	if v272 < v269+int32(38) {
		v233 = v269
		v235 = v272
		goto L76
	} else {
		goto L86
	}
L81:
	;
	F_s_lock(m, v257, int32(497863), int32(754), int32(77246))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v263 == v219 {
		goto L69
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v269 = v268
	goto L80
L86:
	;
	goto L77
L87:
	;
	if v289 == int32(0) {
		goto L68
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+32)) = v219
	F_errmsg(m, int32(128531), v225+int32(32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(497863), int32(798), int32(77246))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L68
L91:
	;
	if v306 == int32(0) {
		goto L68
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(570665), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(497863), int32(733), int32(77246))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	goto L68
L95:
	;
	v462 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L122
	}
L96:
	;
	v322 = v248 + int32(16)
	v323 = int32(0)
	if v178 == v323 {
		v426 = v323
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(0)
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+96)) = int32(0)
	if v426 != 0 {
		goto L95
	} else {
		goto L114
	}
L100:
	;
	v328 = v178 & int32(3)
	if base.Ui32(v178) < base.Ui32(int32(4)) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v328 != 0 {
		goto L108
	} else {
		goto L109
	}
L102:
	;
	v369 = v322
	v370 = v221
	v372 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v335 = v322
	v336 = v221
	v338 = int32(0)
	v344 = int32(0)
	goto L105
L105:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+2)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+3)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+3)))
	v361 = v338 | (v346 ^ v347) | (v350 ^ v351) | (v354 ^ v355) | (v358 ^ v359)
	v362 = int32(4)
	v363 = v336 + v362
	v365 = v335 + v362
	v367 = v344 + v362
	if v367 != v178&int32(-4) {
		v335 = v365
		v336 = v363
		v338 = v361
		v344 = v367
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v369 = v365
	v370 = v363
	v372 = v361
	goto L101
L107:
	;
	goto L106
L108:
	;
	v380 = v369
	v381 = v370
	v383 = v372
	v388 = v323
	goto L111
L109:
	;
	v405 = v372
	goto L110
L110:
	;
	v426 = base.B2i32(v405 != int32(0))
	goto L99
L111:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v394 = v383 | (v391 ^ v392)
	v395 = int32(1)
	v400 = v388 + v395
	if v400 != v328 {
		v380 = v380 + v395
		v381 = v381 + v395
		v383 = v394
		v388 = v400
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v405 = v394
	goto L110
L113:
	;
	goto L112
L114:
	;
	v431 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v431 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v219
	F_errmsg_internal(m, int32(469473), v225)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v445 = F_kill(m, int32(0)-v219, int32(2))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	F_errfinish(m, int32(497863), int32(772), int32(77246))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	goto L68
L122:
	;
	if v462 == int32(0) {
		goto L68
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v219
	F_errmsg(m, int32(469430), v225+int32(16))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(497863), int32(789), int32(77246))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L68
L126:
	;
	v492 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)) = uint8(v492)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _consts[647])))
	if v495 != int32(1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	goto L134
L128:
	;
	v500 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v500 == int32(0) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(447856), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(495925), int32(612), int32(108111))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L127
L133:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	v556 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	goto L144
L134:
	;
	v527 = F_secure_write(m, l0, v12+int32(-16), int32(1))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	v535 = int32(-1)
	v538 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L139
	}
L136:
	;
	if v527 == int32(1) {
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v532 == int32(27) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	if v538 == int32(0) {
		v1153 = v535
		goto L8
	} else {
		goto L140
	}
L140:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(294565), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(495925), int32(621), int32(108111))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v1153 = v535
	goto L8
L144:
	;
	if int32(0) < v554-v556 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v564 = F_ProcessStartupPacket(m, l0, int32(1), base.B2i32(v561 == int32(83)))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v1153 = v564
	goto L8
L147:
	;
	v566 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)) = uint8(v566)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, _consts[647])))
	if v569 != int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	goto L155
L149:
	;
	v574 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	if v574 == int32(0) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(447876), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(495925), int32(666), int32(108111))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	goto L148
L154:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	v630 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	goto L165
L155:
	;
	v601 = F_secure_write(m, l0, v12+int32(-16), int32(1))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v609 = int32(-1)
	v612 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L160
	}
L157:
	;
	if v601 == int32(1) {
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v606 == int32(27) {
		goto L155
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	if v612 == int32(0) {
		v1153 = v609
		goto L8
	} else {
		goto L161
	}
L161:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(294609), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(495925), int32(675), int32(108111))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v1153 = v609
	goto L8
L165:
	;
	if int32(0) < v628-v630 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v638 = F_ProcessStartupPacket(m, l0, base.B2i32(v634 == int32(71)), int32(1))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v1153 = v638
	goto L8
L168:
	;
	v645 = v642
	goto L170
L169:
	;
	v645 = v147
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v645
	if base.Ui32(v147-int32(262144)) <= base.Ui32(int32(-65537)) {
		goto L5
	} else {
		goto L171
	}
L171:
	;
	v651 = int32(4515600)
	v652 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v655 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v655
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = int32(0)
	v659 = int32(4)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v660 < int32(5) {
		v968 = v659
		v969 = v640
		v974 = v660
		goto L172
	} else {
		goto L173
	}
L172:
	;
	if v968 != v974-int32(1) {
		goto L4
	} else {
		goto L278
	}
L173:
	;
	v664 = v659
	v665 = v640
	v670 = v660
	goto L174
L174:
	;
	v674 = v664 + v99
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v675 == int32(0) {
		v968 = v664
		v969 = v665
		v974 = v670
		goto L172
	} else {
		goto L176
	}
L175:
	;
	v968 = v964
	v969 = v959
	v974 = v965
	goto L172
L176:
	;
	v678 = F_strlen(m, v674)
	mBase = m.M
	v681 = v678 + v664 + int32(1)
	if v670 <= v681 {
		v968 = v664
		v969 = v665
		v974 = v670
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v683 = v99 + v681
	v684 = int32(362720)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, _consts[649])))
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v688 == int32(0) {
		v707 = v687
		v708 = v688
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v961 = F_strlen(m, v683)
	mBase = m.M
	v964 = v961 + v681 + int32(1)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v964 < v965 {
		v664 = v964
		v665 = v959
		v670 = v965
		goto L174
	} else {
		goto L277
	}
L179:
	;
	if v708-v707 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L180:
	;
	goto L179
L181:
	;
	if v687 != v688 {
		v707 = v687
		v708 = v688
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v692 = v674
	v693 = v684
	goto L183
L183:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693)+1)))
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692)+1)))
	if v697 == int32(0) {
		v707 = v696
		v708 = v697
		goto L180
	} else {
		goto L185
	}
L184:
	;
	v707 = v696
	v708 = v697
	goto L180
L185:
	;
	v700 = int32(1)
	if v696 == v697 {
		v692 = v692 + v700
		v693 = v693 + v700
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v712 = F_pstrdup(m, v683)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v715 = int32(217760)
	v718 = int32(*(*uint8)(unsafe.Add(mBase, _consts[650])))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v719 == int32(0) {
		v738 = v718
		v739 = v719
		goto L192
	} else {
		goto L193
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v712
	v959 = v665
	goto L178
L191:
	;
	if v739-v738 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L192:
	;
	goto L191
L193:
	;
	if v718 != v719 {
		v738 = v718
		v739 = v719
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v723 = v674
	v724 = v715
	goto L195
L195:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724)+1)))
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+1)))
	if v728 == int32(0) {
		v738 = v727
		v739 = v728
		goto L192
	} else {
		goto L197
	}
L196:
	;
	v738 = v727
	v739 = v728
	goto L192
L197:
	;
	v731 = int32(1)
	if v727 == v728 {
		v723 = v723 + v731
		v724 = v724 + v731
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	v743 = F_pstrdup(m, v683)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v746 = int32(138291)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, _consts[651])))
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v750 == int32(0) {
		v769 = v749
		v770 = v750
		goto L204
	} else {
		goto L205
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v743
	v959 = v665
	goto L178
L203:
	;
	if v770-v769 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L204:
	;
	goto L203
L205:
	;
	if v749 != v750 {
		v769 = v749
		v770 = v750
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v754 = v674
	v755 = v746
	goto L207
L207:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+1)))
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v754)+1)))
	if v759 == int32(0) {
		v769 = v758
		v770 = v759
		goto L204
	} else {
		goto L209
	}
L208:
	;
	v769 = v758
	v770 = v759
	goto L204
L209:
	;
	v762 = int32(1)
	if v758 == v759 {
		v754 = v754 + v762
		v755 = v755 + v762
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v774 = F_pstrdup(m, v683)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v777 = int32(266713)
	v780 = int32(*(*uint8)(unsafe.Add(mBase, _consts[652])))
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v781 == int32(0) {
		v800 = v780
		v801 = v781
		goto L216
	} else {
		goto L217
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v774
	v959 = v665
	goto L178
L215:
	;
	if v801-v800 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L216:
	;
	goto L215
L217:
	;
	if v780 != v781 {
		v800 = v780
		v801 = v781
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v785 = v674
	v786 = v777
	goto L219
L219:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+1)))
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785)+1)))
	if v790 == int32(0) {
		v800 = v789
		v801 = v790
		goto L216
	} else {
		goto L221
	}
L220:
	;
	v800 = v789
	v801 = v790
	goto L216
L221:
	;
	v793 = int32(1)
	if v789 == v790 {
		v785 = v785 + v793
		v786 = v786 + v793
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v805 = int32(362720)
	v808 = int32(*(*uint8)(unsafe.Add(mBase, _consts[649])))
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683))))
	if v809 == int32(0) {
		v828 = v808
		v829 = v809
		goto L227
	} else {
		goto L228
	}
L224:
	;
	goto L225
L225:
	;
	v866 = int32(655817)
	goto L246
L226:
	;
	if v829-v828 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L227:
	;
	goto L226
L228:
	;
	if v808 != v809 {
		v828 = v808
		v829 = v809
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v813 = v683
	v814 = v805
	goto L230
L230:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v814)+1)))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813)+1)))
	if v818 == int32(0) {
		v828 = v817
		v829 = v818
		goto L227
	} else {
		goto L232
	}
L231:
	;
	v828 = v817
	v829 = v818
	goto L227
L232:
	;
	v821 = int32(1)
	if v817 == v818 {
		v813 = v813 + v821
		v814 = v814 + v821
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v834 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[653])) = uint8(v834)
	*(*uint8)(unsafe.Add(mBase, _consts[654])) = uint8(v834)
	v959 = v665
	goto L178
L235:
	;
	goto L236
L236:
	;
	v840 = F_strlen(m, v683)
	mBase = m.M
	v841 = F_parse_bool_with_len(m, v683, v840, int32(4426016))
	mBase = m.M
	goto L237
L237:
	;
	if v841 != 0 {
		v959 = v665
		goto L178
	} else {
		goto L238
	}
L238:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v683
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(266713)
	F_errmsg(m, int32(728018), v12+int32(-32))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errhint(m, int32(668958), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(495925), int32(784), int32(108111))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	if v903-v904 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L246:
	;
	goto L247
L247:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v873 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v874 = v674
	v875 = v866
	v876 = int32(5)
	v877 = v873
	goto L252
L249:
	;
	v899 = v866
	v903 = int32(0)
	goto L250
L250:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899))))
	goto L244
L251:
	;
	v899 = v894
	v903 = v896
	goto L250
L252:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	if v877 != v879 {
		v894 = v875
		v896 = v877
		goto L251
	} else {
		goto L254
	}
L253:
	;
	v894 = v888
	v896 = int32(0)
	goto L251
L254:
	;
	if v879 == int32(0) {
		v894 = v875
		v896 = v877
		goto L251
	} else {
		goto L255
	}
L255:
	;
	v884 = v876 - int32(1)
	if v884 == int32(0) {
		v894 = v875
		v896 = v877
		goto L251
	} else {
		goto L256
	}
L256:
	;
	v887 = int32(1)
	v888 = v875 + v887
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)))
	if v889 != 0 {
		v874 = v874 + v887
		v875 = v888
		v876 = v884
		v877 = v889
		goto L252
	} else {
		goto L257
	}
L257:
	;
	goto L253
L258:
	;
	v914 = F_pstrdup(m, v674)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v919 = F_pstrdup(m, v674)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L263
	}
L261:
	;
	v916 = F_lappend(m, v665, v914)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v959 = v916
	goto L178
L263:
	;
	v921 = F_lappend(m, v918, v919)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v921
	v924 = F_pstrdup(m, v683)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v926 = F_lappend(m, v921, v924)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v926
	v929 = int32(379219)
	v932 = int32(*(*uint8)(unsafe.Add(mBase, _consts[655])))
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v933 == int32(0) {
		v952 = v932
		v953 = v933
		goto L268
	} else {
		goto L269
	}
L267:
	;
	if v953-v952 != 0 {
		v959 = v665
		goto L178
	} else {
		goto L275
	}
L268:
	;
	goto L267
L269:
	;
	if v932 != v933 {
		v952 = v932
		v953 = v933
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v937 = v674
	v938 = v929
	goto L271
L271:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938)+1)))
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)))
	if v942 == int32(0) {
		v952 = v941
		v953 = v942
		goto L268
	} else {
		goto L273
	}
L272:
	;
	v952 = v941
	v953 = v942
	goto L268
L273:
	;
	v945 = int32(1)
	if v941 == v942 {
		v937 = v937 + v945
		v938 = v938 + v945
		goto L271
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	v956 = F_pg_clean_ascii(m, v683, int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v956
	v959 = v665
	goto L178
L277:
	;
	goto L175
L278:
	;
	v981 = int32(0)
	if base.B2i32(v969 == v981)&base.B2i32(base.Ui32(v147&int32(65535)) <= base.Ui32(int32(2))) == v981 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	F_pq_beginmessage(m, v12+int32(-16), int32(118))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v1110 == int32(0) {
		goto L3
	} else {
		goto L296
	}
L282:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	F_enlargeStringInfo(m, v12+int32(-16), int32(4))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v1006 = int32(24)
	v1008 = int32(65280)
	v1010 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1003+v1004))) = v997<<(uint(v1006)%32) | v997&v1008<<(uint(v1010)%32) | (int32(base.Ui32(v997)>>(uint(v1010)%32))&v1008 | int32(base.Ui32(v997)>>(uint(v1006)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v1003 + int32(4)
	if v969 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	v1026 = v1025
	goto L286
L285:
	;
	v1026 = int32(0)
	goto L286
L286:
	;
	F_enlargeStringInfo(m, v12+int32(-16), int32(4))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v1035 = int32(24)
	v1037 = int32(65280)
	v1039 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1032+v1033))) = v1026<<(uint(v1035)%32) | v1026&v1037<<(uint(v1039)%32) | (int32(base.Ui32(v1026)>>(uint(v1039)%32))&v1037 | int32(base.Ui32(v1026)>>(uint(v1035)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v1032 + int32(4)
	if v969 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	F_pq_endmessage(m, v12+int32(-16))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L295
	}
L289:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	if v1056 <= int32(0) {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1063 = int32(0)
	goto L291
L291:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v969)+12))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v1063<<(uint(int32(2))%32))))
	F_pq_sendstring(m, v12+int32(-16), v1077)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L293
	}
L292:
	;
	goto L288
L293:
	;
	v1081 = v1063 + int32(1)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	if v1081 < v1082 {
		v1063 = v1081
		goto L291
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	goto L281
L296:
	;
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110))))
	if v1113 == int32(0) {
		goto L3
	} else {
		goto L297
	}
L297:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	if v1116 != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1122 = F_strlen(m, v1121)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v1122) {
		goto L304
	} else {
		goto L305
	}
L299:
	;
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116))))
	if v1117 != 0 {
		v1121 = v1116
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v1118 = F_pstrdup(m, v1110)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L303
	}
L302:
	;
	goto L301
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v1118
	v1121 = v1118
	goto L298
L304:
	;
	v1125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1121)+63)) = uint8(v1125)
	goto L306
L305:
	;
	goto L306
L306:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	v1128 = F_strlen(m, v1127)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v1128) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1127)+63)) = uint8(v1131)
	goto L309
L308:
	;
	goto L309
L309:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, _consts[654])))
	if v1138 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1139 = int32(6)
	goto L312
L311:
	;
	v1139 = int32(1)
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, _consts[172])) = v1139
	if v1138 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v652
	v1153 = int32(0)
	goto L8
L314:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, _consts[653])))
	if v1144 != 0 {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	v1146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1145))) = uint8(v1146)
	goto L313
L316:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errmsg(m, int32(77180), int32(0))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errdetail(m, int32(624309), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(495925), int32(640), int32(108111))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	F_errmsg(m, int32(76977), int32(0))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	F_errdetail(m, int32(624309), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(495925), int32(694), int32(108111))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(12884901891)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v147 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(base.Ui32(v147) >> (uint(int32(16)) % 32))
	F_errmsg(m, int32(39222), v14)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(495925), int32(725), int32(108111))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	F_errmsg(m, int32(348713), int32(0))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(495925), int32(825), int32(108111))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	F_errmsg(m, int32(107885), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(495925), int32(842), int32(108111))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	if v11 != 0 {
		v15 = m.G0
		v16 = int32(16)
		v17 = v15 - v16
		m.G0 = v17
		F___gettimeofday(m, v17)
		mBase = m.M
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
		m.G0 = v17 + v16
		v31 = *(*int64)(unsafe.Add(mBase, _consts[476]))
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
		*(*int32)(unsafe.Add(mBase, _consts[475])) = int32(0)
	} else {
	}
	m.G0 = v8 + int32(16)
	return base.B2i32(v11 != int32(0))
}
