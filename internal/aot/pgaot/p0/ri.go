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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v915 int32
	_ = v915
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v963 int32
	_ = v963
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1053 int32
	_ = v1053
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
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
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L1
	} else {
		goto L220
	}
L5:
	;
	F_relation_close(m, v28, int32(2))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L1
	} else {
		goto L219
	}
L6:
	;
	v1190 = int32(1)
	v1194 = F_ri_PerformCheck(m, v24, v19+int32(984), v1171, v28, v31, v30, int32(0), l1^v1190, v1190, int32(5))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
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
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L135
	}
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L91
	}
L10:
	;
	v443 = int32(6)
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
	v443 = int32(5)
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
	v417 = v46
	goto L20
L20:
	;
	v430 = int32(5)
	v433 = int32(0)
	v437 = F_ri_PerformCheck(m, v24, v19+int32(968), v417, v28, v31, v30, v433, v433, int32(1), v430)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
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
	v101 = v100 + v99
	v102 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)) = uint8(v96)
	v111 = v101 + int32(1)
	v114 = v104 + int32(4)
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
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v127 != int32(34) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v144 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v111)+1)) = uint16(v144)
	if v55 == int32(112) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v139)
	v111 = v140
	v114 = v114 + int32(1)
	goto L31
L35:
	;
	if v127 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v134 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)) = uint8(v134)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v139 = v136
	v140 = v111 + int32(2)
	goto L34
L38:
	;
	v139 = v127
	v140 = v111 + int32(1)
	goto L34
L39:
	;
	v150 = int32(_a_F_ri_restrict_0)
	goto L41
L40:
	;
	v150 = int32(_a_F_ri_restrict_1)
	goto L41
L41:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v151 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if int32(0) < v241 {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+v154<<(uint(int32(1))%32))+170)))
	v159 = F_attnumAttName(m, v31, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v19)+228)) = v19 + int32(992)
	F_appendStringInfo(m, v19+int32(416), int32(_a_F_ri_restrict_2), v19+int32(224))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L56
	}
L46:
	;
	v161 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v161)
	v165 = v159
	v168 = v19 + int32(704)
	goto L47
L47:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v181 != int32(34) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v198 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v168)+1)) = uint16(v198)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+212)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v19)+216)) = v19 + int32(992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = v19 + int32(704)
	F_appendStringInfo(m, v19+int32(416), int32(_a_F_ri_restrict_3), v19+int32(208))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L55
	}
L49:
	;
	goto L48
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v193)
	v165 = v165 + int32(1)
	v168 = v194
	goto L47
L51:
	;
	if v181 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v188 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)) = uint8(v188)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v193 = v190
	v194 = v168 + int32(2)
	goto L50
L54:
	;
	v193 = v181
	v194 = v168 + int32(1)
	goto L50
L55:
	;
	goto L42
L56:
	;
	goto L42
L57:
	;
	v255 = int32(0)
	v258 = int32(_a_F_ri_restrict_4)
	goto L60
L58:
	;
	goto L59
L59:
	;
	v364 = v19 + int32(416)
	F_appendStringInfoString(m, v364, int32(_a_F_ri_restrict_5))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L76
	}
L60:
	;
	v268 = v24 + int32(172) + v255<<(uint(int32(1))%32)
	v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v268))))
	v270 = F_attnumTypeId(m, v31, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	v272 = int32(*(*int16)(unsafe.Add(mBase, uint32(v268))))
	v273 = F_attnumAttName(m, v31, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v275 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v275)
	v279 = v273
	v282 = v19 + int32(704)
	goto L64
L64:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v295 != int32(34) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v312 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+1)) = uint16(v312)
	v315 = v255 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v315
	v318 = v19 + int32(272)
	v322 = F_pg_sprintf(m, v318, int32(_a_F_ri_restrict_6), v19+int32(192))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v307)
	v279 = v279 + int32(1)
	v282 = v308
	goto L64
L68:
	;
	if v295 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v302 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)) = uint8(v302)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v307 = v304
	v308 = v282 + int32(2)
	goto L67
L71:
	;
	v307 = v295
	v308 = v282 + int32(1)
	goto L67
L72:
	;
	v325 = v255 << (uint(int32(2)) % 32)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(428)+v325)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v258
	v330 = v19 + int32(416)
	F_appendStringInfo(m, v330, int32(_a_F_ri_restrict_7), v19+int32(176))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_generate_operator_clause(m, v330, v19+int32(704), v270, v327, v318, v270)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(560)+v325))) = v270
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v315 < v345 {
		v255 = v315
		v258 = int32(_a_F_ri_restrict_8)
		goto L60
	} else {
		goto L75
	}
L75:
	;
	goto L61
L76:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v368 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v373 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+v369<<(uint(int32(1))%32))+234)))
	v374 = F_attnumTypeId(m, v28, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v19)+416))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v412 = F_ri_PlanCheck(m, v406, v407, v19+int32(560), v19+int32(968), v28, v31)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L86
	}
L80:
	;
	F_appendStringInfoString(m, v364, int32(_a_F_ri_restrict_9))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v379
	v382 = v19 + int32(272)
	v386 = F_pg_sprintf(m, v382, int32(_a_F_ri_restrict_6), v19+int32(160))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v24)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = int32(_a_F_ri_restrict_0)
	F_appendStringInfo(m, v364, int32(_a_F_ri_restrict_7), v19+int32(144))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_generate_operator_clause(m, v364, v382, v374, v388, int32(_a_F_ri_restrict_10), int32(_a_F_ri_restrict_11))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_appendStringInfoString(m, v364, int32(_a_F_ri_restrict_12))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L79
L86:
	;
	v417 = v412
	goto L20
L87:
	;
	v439 = F_SPI_finish(m)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v439 != int32(2) {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	if v437 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	v443 = v430
	goto L9
L91:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+988)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v19)+984)) = v462
	v467 = F_ri_FetchPreparedPlan(m, v19+int32(984))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v467 != 0 {
		v1171 = v467
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_initStringInfo(m, v19+int32(968))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+119)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v473)+68))
	v476 = F_get_namespace_name(m, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v478 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v478)
	v482 = v476
	v485 = v19 + int32(704)
	goto L96
L96:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v498 != int32(34) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v515 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v485)+1)) = uint16(v515)
	v518 = v19 + int32(704)
	v519 = F_strlen(m, v518)
	mBase = m.M
	v520 = v519 + v518
	v521 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v520))) = uint8(v521)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+1)) = uint8(v515)
	v530 = v520 + int32(1)
	v533 = v523 + int32(4)
	goto L104
L98:
	;
	goto L97
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v511))) = uint8(v510)
	v482 = v482 + int32(1)
	v485 = v511
	goto L96
L100:
	;
	if v498 == int32(0) {
		goto L98
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v505 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v485)+1)) = uint8(v505)
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	v510 = v507
	v511 = v485 + int32(2)
	goto L99
L103:
	;
	v510 = v498
	v511 = v485 + int32(1)
	goto L99
L104:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	if v546 != int32(34) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v563 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v530)+1)) = uint16(v563)
	if v474 == int32(112) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v559))) = uint8(v558)
	v530 = v559
	v533 = v533 + int32(1)
	goto L104
L108:
	;
	if v546 == int32(0) {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v553 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v530)+1)) = uint8(v553)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	v558 = v555
	v559 = v530 + int32(2)
	goto L107
L111:
	;
	v558 = v546
	v559 = v530 + int32(1)
	goto L107
L112:
	;
	v569 = int32(_a_F_ri_restrict_0)
	goto L114
L113:
	;
	v569 = int32(_a_F_ri_restrict_1)
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = v19 + int32(704)
	F_appendStringInfo(m, v19+int32(968), int32(_a_F_ri_restrict_2), v19+int32(128))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v581 <= int32(0) {
		v709 = v581
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v597 = int32(0)
	v600 = int32(_a_F_ri_restrict_4)
	goto L117
L117:
	;
	v609 = v597 << (uint(int32(1)) % 32)
	v611 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(172)+v609))))
	v612 = F_attnumTypeId(m, v31, v611)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v614 = v609 + (v24 + int32(236))
	v615 = int32(*(*int16)(unsafe.Add(mBase, uint32(v614))))
	v616 = F_attnumTypeId(m, v28, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v618 = int32(*(*int16)(unsafe.Add(mBase, uint32(v614))))
	v619 = F_attnumAttName(m, v28, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v621 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v621)
	v625 = v619
	v628 = v19 + int32(560)
	goto L122
L122:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	if v641 != int32(34) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v689))) = uint8(v688)
	v625 = v625 + int32(1)
	v628 = v689
	goto L122
L125:
	;
	if v641 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v683 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v628)+1)) = uint8(v683)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	v688 = v685
	v689 = v628 + int32(2)
	goto L124
L128:
	;
	v646 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v628)+1)) = uint16(v646)
	v649 = v597 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v649
	v652 = v19 + int32(400)
	v656 = F_pg_sprintf(m, v652, int32(_a_F_ri_restrict_6), v19+int32(112))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v688 = v641
	v689 = v628 + int32(1)
	goto L124
L131:
	;
	v659 = v597 << (uint(int32(2)) % 32)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(300)+v659)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v600
	v664 = v19 + int32(968)
	F_appendStringInfo(m, v664, int32(_a_F_ri_restrict_7), v19+int32(96))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_generate_operator_clause(m, v664, v652, v612, v661, v19+int32(560), v616)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(272)+v659))) = v612
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v649 < v679 {
		v597 = v649
		v600 = int32(_a_F_ri_restrict_8)
		goto L117
	} else {
		goto L134
	}
L134:
	;
	v709 = v679
	goto L7
L135:
	;
	F_errmsg_internal(m, int32(_a_F_ri_restrict_13), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ri_restrict_14), int32(625), int32(_a_F_ri_restrict_15))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
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
	F_appendStringInfoString(m, v19+int32(968), int32(_a_F_ri_restrict_5))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L214
	}
L139:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v724&int32(1) == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v730 = v24 + int32(172)
	v736 = int32(*(*int16)(unsafe.Add(mBase, uint32(v730+v709<<(uint(int32(1))%32)-int32(2)))))
	v737 = F_attnumTypeId(m, v31, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v740 = v24 + int32(236)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v747 = int32(*(*int16)(unsafe.Add(mBase, uint32(v740+v741<<(uint(int32(1))%32)-int32(2)))))
	v748 = F_attnumTypeId(m, v28, v747)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+119)))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v758 = int32(*(*int16)(unsafe.Add(mBase, uint32(v740+v752<<(uint(int32(1))%32)-int32(2)))))
	v759 = F_attnumAttName(m, v28, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v761 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v761)
	v765 = v759
	v768 = v19 + int32(560)
	goto L144
L144:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	if v781 != int32(34) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v798 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v768)+1)) = uint16(v798)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v800
	v803 = v19 + int32(400)
	v807 = F_pg_sprintf(m, v803, int32(_a_F_ri_restrict_6), v19+int32(80))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L152
	}
L146:
	;
	goto L145
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v794))) = uint8(v793)
	v765 = v765 + int32(1)
	v768 = v794
	goto L144
L148:
	;
	if v781 == int32(0) {
		goto L146
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v788 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v768)+1)) = uint8(v788)
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	v793 = v790
	v794 = v768 + int32(2)
	goto L147
L151:
	;
	v793 = v781
	v794 = v768 + int32(1)
	goto L147
L152:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(_a_F_ri_restrict_16))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v815 = v19 + int32(256)
	F_initStringInfo(m, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_appendStringInfoChar(m, v815, int32(40))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v24)+692))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(_a_F_ri_restrict_0)
	F_appendStringInfo(m, v815, int32(_a_F_ri_restrict_7), v19-int32(-64))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_generate_operator_clause(m, v815, v19+int32(560), v748, v821, v803, v737)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_appendStringInfoChar(m, v815, int32(41))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v837 = v19 + int32(240)
	F_initStringInfo(m, v837)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_appendStringInfoString(m, v837, int32(_a_F_ri_restrict_17))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v849 = int32(*(*int16)(unsafe.Add(mBase, uint32(v730+v843<<(uint(int32(1))%32)-int32(2)))))
	v850 = F_attnumAttName(m, v31, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v852 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+416)) = uint8(v852)
	v856 = v850
	v859 = v19 + int32(416)
	goto L162
L162:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	if v872 != int32(34) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v889 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v859)+1)) = uint16(v889)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+68))
	v893 = F_get_namespace_name(m, v892)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L170
	}
L164:
	;
	goto L163
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v885))) = uint8(v884)
	v856 = v856 + int32(1)
	v859 = v885
	goto L162
L166:
	;
	if v872 == int32(0) {
		goto L164
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v879 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v859)+1)) = uint8(v879)
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	v884 = v881
	v885 = v859 + int32(2)
	goto L165
L169:
	;
	v884 = v872
	v885 = v859 + int32(1)
	goto L165
L170:
	;
	v895 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+992)) = uint8(v895)
	v899 = v893
	v902 = v19 + int32(992)
	goto L171
L171:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899))))
	if v915 != int32(34) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v932 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v902)+1)) = uint16(v932)
	v935 = v19 + int32(992)
	v936 = F_strlen(m, v935)
	mBase = m.M
	v937 = v936 + v935
	v938 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v938)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)) = uint8(v932)
	v947 = v937 + int32(1)
	v950 = v940 + int32(4)
	goto L179
L173:
	;
	goto L172
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v928))) = uint8(v927)
	v899 = v899 + int32(1)
	v902 = v928
	goto L171
L175:
	;
	if v915 == int32(0) {
		goto L173
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v922 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)) = uint8(v922)
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899))))
	v927 = v924
	v928 = v902 + int32(2)
	goto L174
L178:
	;
	v927 = v915
	v928 = v902 + int32(1)
	goto L174
L179:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950))))
	if v963 != int32(34) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v980 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v947)+1)) = uint16(v980)
	if v751 == int32(112) {
		goto L187
	} else {
		goto L188
	}
L181:
	;
	goto L180
L182:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v976))) = uint8(v975)
	v947 = v976
	v950 = v950 + int32(1)
	goto L179
L183:
	;
	if v963 == int32(0) {
		goto L181
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v970 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v947)+1)) = uint8(v970)
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950))))
	v975 = v972
	v976 = v947 + int32(2)
	goto L182
L186:
	;
	v975 = v963
	v976 = v947 + int32(1)
	goto L182
L187:
	;
	v986 = int32(_a_F_ri_restrict_0)
	goto L189
L188:
	;
	v986 = int32(_a_F_ri_restrict_1)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v19 + int32(992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v19 + int32(416)
	F_appendStringInfo(m, v19+int32(240), int32(_a_F_ri_restrict_18), v19+int32(48))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if int32(0) < v1001 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1013 = int32(0)
	v1016 = int32(_a_F_ri_restrict_4)
	goto L194
L192:
	;
	goto L193
L193:
	;
	F_appendStringInfoString(m, v19+int32(240), int32(_a_F_ri_restrict_19))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L210
	}
L194:
	;
	v1026 = v730 + v1013<<(uint(int32(1))%32)
	v1027 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1026))))
	v1028 = F_attnumTypeId(m, v31, v1027)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L196
	}
L195:
	;
	goto L193
L196:
	;
	v1030 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1026))))
	v1031 = F_attnumAttName(m, v31, v1030)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v1033 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v1033)
	v1037 = v1031
	v1040 = v19 + int32(560)
	goto L198
L198:
	;
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037))))
	if v1053 != int32(34) {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v1070 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v1040)+1)) = uint16(v1070)
	v1073 = v1013 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1073
	v1076 = v19 + int32(400)
	v1080 = F_pg_sprintf(m, v1076, int32(_a_F_ri_restrict_6), v19+int32(32))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L206
	}
L200:
	;
	goto L199
L201:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1066))) = uint8(v1065)
	v1037 = v1037 + int32(1)
	v1040 = v1066
	goto L198
L202:
	;
	if v1053 == int32(0) {
		goto L200
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1060 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1040)+1)) = uint8(v1060)
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037))))
	v1065 = v1062
	v1066 = v1040 + int32(2)
	goto L201
L205:
	;
	v1065 = v1053
	v1066 = v1040 + int32(1)
	goto L201
L206:
	;
	v1083 = v1013 << (uint(int32(2)) % 32)
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(428)+v1083)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1016
	v1088 = v19 + int32(240)
	F_appendStringInfo(m, v1088, int32(_a_F_ri_restrict_7), v19+int32(16))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_generate_operator_clause(m, v1088, v1076, v1028, v1085, v19+int32(560), v1028)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(272)+v1083))) = v1028
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v1073 < v1103 {
		v1013 = v1073
		v1016 = int32(_a_F_ri_restrict_8)
		goto L194
	} else {
		goto L209
	}
L209:
	;
	goto L195
L210:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v24)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_ri_restrict_0)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v19)+240))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
	v1132 = v19 + int32(968)
	F_appendStringInfo(m, v1132, int32(_a_F_ri_restrict_7), v19)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_generate_operator_clause(m, v1132, v1130, v748, v1126, v1129, int32(_a_F_ri_restrict_11))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_appendStringInfoString(m, v1132, int32(_a_F_ri_restrict_20))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	goto L138
L214:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v19)+968))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v1169 = F_ri_PlanCheck(m, v1163, v1164, v19+int32(272), v19+int32(984), v28, v31)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1171 = v1169
	goto L6
L216:
	;
	v1196 = F_SPI_finish(m)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	if v1196 != int32(2) {
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
	F_errmsg_internal(m, int32(_a_F_ri_restrict_13), int32(0))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_ri_restrict_14), int32(901), int32(_a_F_ri_restrict_21))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
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
