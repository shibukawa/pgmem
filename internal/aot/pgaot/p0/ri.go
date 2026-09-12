package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ri_restrict(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v245 int32
	_ = v245
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
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
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v857 int32
	_ = v857
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v1002 int32
	_ = v1002
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1027 int32
	_ = v1027
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1134 int32
	_ = v1134
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1267 int32
	_ = v1267
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	v17 = m.G0
	v19 = v17 - int32(1264)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_ri_FetchConstraintInfo(m, v21, v22, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	v28 = F_table_open(m, v26, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L1
	} else {
		goto L220
	}
L5:
	;
	F_sequence_close(m, v28, int32(2))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L219
	}
L6:
	;
	v1239 = int32(1)
	v1243 = F_ri_PerformCheck(m, v24, v19+int32(984), v1220, v28, v31, v30, int32(0), l1^v1239, v1239, int32(5))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L216
	}
L7:
	;
	if l1 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L135
	}
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L91
	}
L10:
	;
	v460 = int32(6)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v460 = int32(5)
	goto L9
L14:
	;
	goto L15
L15:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+972)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+968)) = v40
	v46 = F_ri_FetchPreparedPlan(m, v19+int32(968))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v46 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_initStringInfo(m, v19+int32(416))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v434 = v46
	goto L20
L20:
	;
	v447 = int32(5)
	v450 = int32(0)
	v454 = F_ri_PerformCheck(m, v24, v19+int32(968), v434, v28, v31, v30, v450, v450, int32(1), v447)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L87
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+119)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
	v57 = F_get_namespace_name(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v59 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+992)) = uint8(v59)
	v63 = v57
	v66 = v19 + int32(992)
	goto L23
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v79 != int32(34) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v96 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+1)) = uint16(v96)
	v99 = v19 + int32(992)
	v100 = F_strlen(m, v99)
	mBase = m.M
	v103 = v100 + v99
	v104 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v104)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v108 = v103 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v96)
	v113 = v108
	v116 = v106 + int32(4)
	goto L31
L25:
	;
	goto L24
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v91)
	v63 = v63 + int32(1)
	v66 = v92
	goto L23
L27:
	;
	if v79 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v86 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v91 = v88
	v92 = v66 + int32(2)
	goto L26
L30:
	;
	v91 = v79
	v92 = v66 + int32(1)
	goto L26
L31:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v129 != int32(34) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v146 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v113)+1)) = uint16(v146)
	if v55&int32(255) == int32(112) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v141)
	v113 = v142
	v116 = v116 + int32(1)
	goto L31
L35:
	;
	if v129 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v136 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)) = uint8(v136)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v141 = v138
	v142 = v113 + int32(2)
	goto L34
L38:
	;
	v141 = v129
	v142 = v113 + int32(1)
	goto L34
L39:
	;
	v154 = int32(759461)
	goto L41
L40:
	;
	v154 = int32(745885)
	goto L41
L41:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v155 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if int32(0) < v245 {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v158<<(uint(int32(1))%32)+v24)+170)))
	v163 = F_attnumAttName(m, v31, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v19)+228)) = v19 + int32(992)
	F_appendStringInfo(m, v19+int32(416), int32(30266), v19+int32(224))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L56
	}
L46:
	;
	v165 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v165)
	v169 = v163
	v172 = v19 + int32(704)
	goto L47
L47:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v185 != int32(34) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v202 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+1)) = uint16(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+212)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v19)+216)) = v19 + int32(992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = v19 + int32(704)
	F_appendStringInfo(m, v19+int32(416), int32(30224), v19+int32(208))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L55
	}
L49:
	;
	goto L48
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v197)
	v169 = v169 + int32(1)
	v172 = v198
	goto L47
L51:
	;
	if v185 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v192 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)) = uint8(v192)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v197 = v194
	v198 = v172 + int32(2)
	goto L50
L54:
	;
	v197 = v185
	v198 = v172 + int32(1)
	goto L50
L55:
	;
	goto L42
L56:
	;
	goto L42
L57:
	;
	v260 = int32(0)
	v264 = int32(541031)
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_appendStringInfoString(m, v19+int32(416), int32(30287))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L76
	}
L60:
	;
	v272 = v24 + int32(172) + v260<<(uint(int32(1))%32)
	v273 = int32(*(*int16)(unsafe.Add(mBase, uint32(v272))))
	v274 = F_attnumTypeId(m, v31, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	v276 = int32(*(*int16)(unsafe.Add(mBase, uint32(v272))))
	v277 = F_attnumAttName(m, v31, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v279 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v279)
	v283 = v277
	v286 = v19 + int32(704)
	goto L64
L64:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v299 != int32(34) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v316 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+1)) = uint16(v316)
	v319 = v260 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v319
	v326 = F_pg_sprintf(m, v19+int32(272), int32(468463), v19+int32(192))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v312))) = uint8(v311)
	v283 = v283 + int32(1)
	v286 = v312
	goto L64
L68:
	;
	if v299 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v306 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)) = uint8(v306)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	v311 = v308
	v312 = v286 + int32(2)
	goto L67
L71:
	;
	v311 = v299
	v312 = v286 + int32(1)
	goto L67
L72:
	;
	v329 = v260 << (uint(int32(2)) % 32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(428)+v329)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v264
	F_appendStringInfo(m, v19+int32(416), int32(740376), v19+int32(176))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_generate_operator_clause(m, v19+int32(416), v19+int32(704), v274, v331, v19+int32(272), v274)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(560)+v329))) = v274
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v319 < v353 {
		v260 = v319
		v264 = int32(544263)
		goto L60
	} else {
		goto L75
	}
L75:
	;
	goto L61
L76:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v376 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v377<<(uint(int32(1))%32)+v24)+234)))
	v382 = F_attnumTypeId(m, v28, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v19)+416))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v429 = F_ri_PlanCheck(m, v423, v424, v19+int32(560), v19+int32(968), v28, v31)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L86
	}
L80:
	;
	F_appendStringInfoString(m, v19+int32(416), int32(746800))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v389
	v396 = F_pg_sprintf(m, v19+int32(272), int32(468463), v19+int32(160))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v24)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = int32(759461)
	F_appendStringInfo(m, v19+int32(416), int32(740376), v19+int32(144))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_generate_operator_clause(m, v19+int32(416), v19+int32(272), v382, v398, int32(338623), int32(4537))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_appendStringInfoString(m, v19+int32(416), int32(677701))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L79
L86:
	;
	v434 = v429
	goto L20
L87:
	;
	v456 = F_SPI_finish(m)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v456 != int32(2) {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	if v454 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	v460 = v447
	goto L9
L91:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+988)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v19)+984)) = v479
	v484 = F_ri_FetchPreparedPlan(m, v19+int32(984))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v484 != 0 {
		v1220 = v484
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_initStringInfo(m, v19+int32(968))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+119)))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v490)+68))
	v493 = F_get_namespace_name(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v495 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v495)
	v499 = v493
	v502 = v19 + int32(704)
	goto L96
L96:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	if v515 != int32(34) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v532 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v502)+1)) = uint16(v532)
	v535 = v19 + int32(704)
	v536 = F_strlen(m, v535)
	mBase = m.M
	v539 = v536 + v535
	v540 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v539))) = uint8(v540)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v544 = v539 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v532)
	v549 = v544
	v552 = v542 + int32(4)
	goto L104
L98:
	;
	goto L97
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v528))) = uint8(v527)
	v499 = v499 + int32(1)
	v502 = v528
	goto L96
L100:
	;
	if v515 == int32(0) {
		goto L98
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v522 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+1)) = uint8(v522)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v527 = v524
	v528 = v502 + int32(2)
	goto L99
L103:
	;
	v527 = v515
	v528 = v502 + int32(1)
	goto L99
L104:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	if v565 != int32(34) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v582 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v549)+1)) = uint16(v582)
	if v491&int32(255) == int32(112) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v578))) = uint8(v577)
	v549 = v578
	v552 = v552 + int32(1)
	goto L104
L108:
	;
	if v565 == int32(0) {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v572 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+1)) = uint8(v572)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	v577 = v574
	v578 = v549 + int32(2)
	goto L107
L111:
	;
	v577 = v565
	v578 = v549 + int32(1)
	goto L107
L112:
	;
	v590 = int32(759461)
	goto L114
L113:
	;
	v590 = int32(745885)
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = v19 + int32(704)
	F_appendStringInfo(m, v19+int32(968), int32(30266), v19+int32(128))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v602 <= int32(0) {
		v734 = v602
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v619 = int32(0)
	v623 = int32(541031)
	goto L117
L117:
	;
	v630 = v619 << (uint(int32(1)) % 32)
	v632 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(172)+v630))))
	v633 = F_attnumTypeId(m, v31, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v635 = v630 + (v24 + int32(236))
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(v635))))
	v637 = F_attnumTypeId(m, v28, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v639 = int32(*(*int16)(unsafe.Add(mBase, uint32(v635))))
	v640 = F_attnumAttName(m, v28, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v642 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v642)
	v646 = v640
	v649 = v19 + int32(560)
	goto L122
L122:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	if v662 != int32(34) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v714))) = uint8(v713)
	v646 = v646 + int32(1)
	v649 = v714
	goto L122
L125:
	;
	if v662 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v708 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+1)) = uint8(v708)
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	v713 = v710
	v714 = v649 + int32(2)
	goto L124
L128:
	;
	v667 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v649)+1)) = uint16(v667)
	v670 = v619 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v670
	v677 = F_pg_sprintf(m, v19+int32(400), int32(468463), v19+int32(112))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v713 = v662
	v714 = v649 + int32(1)
	goto L124
L131:
	;
	v680 = v619 << (uint(int32(2)) % 32)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(300)+v680)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v623
	F_appendStringInfo(m, v19+int32(968), int32(740376), v19+int32(96))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_generate_operator_clause(m, v19+int32(968), v19+int32(400), v633, v682, v19+int32(560), v637)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(272)+v680))) = v633
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v670 < v704 {
		v619 = v670
		v623 = int32(544263)
		goto L117
	} else {
		goto L134
	}
L134:
	;
	v734 = v704
	goto L7
L135:
	;
	F_errmsg_internal(m, int32(455680), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(495198), int32(625), int32(326823))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(30287))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L214
	}
L139:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v749&int32(1) == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v755 = v24 + int32(172)
	v761 = int32(*(*int16)(unsafe.Add(mBase, uint32(v755+v734<<(uint(int32(1))%32)-int32(2)))))
	v762 = F_attnumTypeId(m, v31, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v765 = v24 + int32(236)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v772 = int32(*(*int16)(unsafe.Add(mBase, uint32(v765+v766<<(uint(int32(1))%32)-int32(2)))))
	v773 = F_attnumTypeId(m, v28, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+119)))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v783 = int32(*(*int16)(unsafe.Add(mBase, uint32(v777<<(uint(int32(1))%32)+v765-int32(2)))))
	v784 = F_attnumAttName(m, v28, v783)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v786 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v786)
	v790 = v784
	v793 = v19 + int32(560)
	goto L144
L144:
	;
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790))))
	if v806 != int32(34) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v823 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v793)+1)) = uint16(v823)
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v825
	v832 = F_pg_sprintf(m, v19+int32(400), int32(468463), v19+int32(80))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L152
	}
L146:
	;
	goto L145
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v819))) = uint8(v818)
	v790 = v790 + int32(1)
	v793 = v819
	goto L144
L148:
	;
	if v806 == int32(0) {
		goto L146
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v813 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v793)+1)) = uint8(v813)
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790))))
	v818 = v815
	v819 = v793 + int32(2)
	goto L147
L151:
	;
	v818 = v806
	v819 = v793 + int32(1)
	goto L147
L152:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(687509))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_initStringInfo(m, v19+int32(256))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_appendStringInfoChar(m, v19+int32(256), int32(40))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v24)+692))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(759461)
	F_appendStringInfo(m, v19+int32(256), int32(740376), v19-int32(-64))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_generate_operator_clause(m, v19+int32(256), v19+int32(560), v773, v848, v19+int32(400), v762)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_appendStringInfoChar(m, v19+int32(256), int32(41))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_initStringInfo(m, v19+int32(240))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_appendStringInfoString(m, v19+int32(240), int32(746680))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v886 = int32(*(*int16)(unsafe.Add(mBase, uint32(v880<<(uint(int32(1))%32)+v755-int32(2)))))
	v887 = F_attnumAttName(m, v31, v886)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v889 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+416)) = uint8(v889)
	v893 = v887
	v896 = v19 + int32(416)
	goto L162
L162:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	if v909 != int32(34) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v926 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v896)+1)) = uint16(v926)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v928)+68))
	v930 = F_get_namespace_name(m, v929)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L170
	}
L164:
	;
	goto L163
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v922))) = uint8(v921)
	v893 = v893 + int32(1)
	v896 = v922
	goto L162
L166:
	;
	if v909 == int32(0) {
		goto L164
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v916 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v896)+1)) = uint8(v916)
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	v921 = v918
	v922 = v896 + int32(2)
	goto L165
L169:
	;
	v921 = v909
	v922 = v896 + int32(1)
	goto L165
L170:
	;
	v932 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+992)) = uint8(v932)
	v936 = v930
	v939 = v19 + int32(992)
	goto L171
L171:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936))))
	if v952 != int32(34) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v969 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v939)+1)) = uint16(v969)
	v972 = v19 + int32(992)
	v973 = F_strlen(m, v972)
	mBase = m.M
	v976 = v973 + v972
	v977 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v976))) = uint8(v977)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v981 = v976 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v981))) = uint8(v969)
	v986 = v981
	v989 = v979 + int32(4)
	goto L179
L173:
	;
	goto L172
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v965))) = uint8(v964)
	v936 = v936 + int32(1)
	v939 = v965
	goto L171
L175:
	;
	if v952 == int32(0) {
		goto L173
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v959 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v939)+1)) = uint8(v959)
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936))))
	v964 = v961
	v965 = v939 + int32(2)
	goto L174
L178:
	;
	v964 = v952
	v965 = v939 + int32(1)
	goto L174
L179:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989))))
	if v1002 != int32(34) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v1019 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v986)+1)) = uint16(v1019)
	if v776&int32(255) == int32(112) {
		goto L187
	} else {
		goto L188
	}
L181:
	;
	goto L180
L182:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1015))) = uint8(v1014)
	v986 = v1015
	v989 = v989 + int32(1)
	goto L179
L183:
	;
	if v1002 == int32(0) {
		goto L181
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v1009 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)) = uint8(v1009)
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989))))
	v1014 = v1011
	v1015 = v986 + int32(2)
	goto L182
L186:
	;
	v1014 = v1002
	v1015 = v986 + int32(1)
	goto L182
L187:
	;
	v1027 = int32(759461)
	goto L189
L188:
	;
	v1027 = int32(745885)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v19 + int32(992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v19 + int32(416)
	F_appendStringInfo(m, v19+int32(240), int32(27045), v19+int32(48))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if int32(0) < v1042 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1055 = int32(0)
	v1059 = int32(541031)
	goto L194
L192:
	;
	goto L193
L193:
	;
	F_appendStringInfoString(m, v19+int32(240), int32(684574))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L210
	}
L194:
	;
	v1067 = v755 + v1055<<(uint(int32(1))%32)
	v1068 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1067))))
	v1069 = F_attnumTypeId(m, v31, v1068)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L196
	}
L195:
	;
	goto L193
L196:
	;
	v1071 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1067))))
	v1072 = F_attnumAttName(m, v31, v1071)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v1074 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v1074)
	v1078 = v1072
	v1081 = v19 + int32(560)
	goto L198
L198:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078))))
	if v1094 != int32(34) {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v1111 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v1081)+1)) = uint16(v1111)
	v1114 = v1055 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1114
	v1121 = F_pg_sprintf(m, v19+int32(400), int32(468463), v19+int32(32))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L206
	}
L200:
	;
	goto L199
L201:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107))) = uint8(v1106)
	v1078 = v1078 + int32(1)
	v1081 = v1107
	goto L198
L202:
	;
	if v1094 == int32(0) {
		goto L200
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1101 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1081)+1)) = uint8(v1101)
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078))))
	v1106 = v1103
	v1107 = v1081 + int32(2)
	goto L201
L205:
	;
	v1106 = v1094
	v1107 = v1081 + int32(1)
	goto L201
L206:
	;
	v1124 = v1055 << (uint(int32(2)) % 32)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(428)+v1124)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1059
	F_appendStringInfo(m, v19+int32(240), int32(740376), v19+int32(16))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_generate_operator_clause(m, v19+int32(240), v19+int32(400), v1069, v1126, v19+int32(560), v1069)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(272)+v1124))) = v1069
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v1114 < v1148 {
		v1055 = v1114
		v1059 = int32(544263)
		goto L194
	} else {
		goto L209
	}
L209:
	;
	goto L195
L210:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v24)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(759461)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v19)+240))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
	F_appendStringInfo(m, v19+int32(968), int32(740376), v19)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_generate_operator_clause(m, v19+int32(968), v1175, v773, v1171, v1174, int32(4537))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(678864))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	goto L138
L214:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v19)+968))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v1218 = F_ri_PlanCheck(m, v1212, v1213, v19+int32(272), v19+int32(984), v28, v31)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1220 = v1218
	goto L6
L216:
	;
	v1245 = F_SPI_finish(m)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	if v1245 != int32(2) {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	goto L5
L219:
	;
	m.G0 = v19 + int32(1264)
	return
L220:
	;
	F_errmsg_internal(m, int32(455680), int32(0))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(495198), int32(901), int32(109607))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
