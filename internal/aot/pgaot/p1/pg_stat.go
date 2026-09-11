package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_force_next_flush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[914])) = uint8(v3)
	return int32(0)
}
func F_pg_stat_get_activity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v66 int32
	_ = v66
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v368 int32
	_ = v368
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int64
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int64
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int64
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int64
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int64
	_ = v521
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v601 int32
	_ = v601
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v649 int32
	_ = v649
	var v672 int32
	_ = v672
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v817 int32
	_ = v817
	var v841 int32
	_ = v841
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1145 int64
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	v23 = m.G0
	v25 = v23 - int32(448)
	m.G0 = v25
	v27 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v32 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = v35
	goto L5
L4:
	;
	v36 = int32(-1)
	goto L5
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v27 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v25 + int32(448)
	return int32(0)
L8:
	;
	v44 = v25 + int32(307)
	v46 = v25 + int32(305)
	v50 = v25 + int32(288) | int32(6)
	v66 = int32(1)
	goto L9
L9:
	;
	v85 = F__emscripten_memset_bulkmem(m, v25+int32(320), base.I32_extend8_s(int32(0)), int32(124))
	mBase = m.M
	goto L11
L10:
	;
	goto L7
L11:
	;
	v86 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25+int32(311)))) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v25+int32(304)))) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v25)+296)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v25)+288)) = v86
	v94 = F_pgstat_get_local_beentry_by_index(m, v66)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v36 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	if v101 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v96 == v36 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v99 = v66 + int32(1)
	if v99 <= v27 {
		v66 = v99
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+324)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v94)+52))
	if v107 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+320)) = v101
	goto L17
L19:
	;
	goto L20
L20:
	;
	v103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+288)) = uint8(v103)
	goto L17
L21:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v94)+212))
	if v111 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+328)) = v107
	goto L21
L23:
	;
	goto L24
L24:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+290)) = uint8(v109)
	goto L21
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v94)+412))
	if v117 != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v112 = F_cstring_to_text(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+291)) = uint8(v115)
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+332)) = v112
	goto L25
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v94)+416))
	if v121 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+380)) = v117
	goto L30
L32:
	;
	goto L33
L33:
	;
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+303)) = uint8(v119)
	goto L30
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v128 = F_has_privs_of_role(m, v126, int32(3375))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L45
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+384)) = v121
	goto L34
L36:
	;
	goto L37
L37:
	;
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+304)) = uint8(v123)
	goto L34
L38:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	F_tuplestore_putvalues(m, v1175, v1176, v25+int32(320), v25+int32(288))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L272
	}
L39:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v865 == int32(5) {
		goto L200
	} else {
		goto L201
	}
L40:
	;
	v841 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+300)) = uint16(v841)
	goto L39
L41:
	;
	v817 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+302)) = uint8(v817)
	goto L40
L42:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487))))
	switch v699 - int32(1) {
	case 0:
		goto L169
	case 1, 9:
		goto L170
	default:
		goto L41
	}
L43:
	;
	v532 = v487 + (int32(-56)-v94)&int32(3)
	v534 = v485 & int32(-4)
	v536 = v534 - int32(28)
	if base.Ui32(v532) < base.Ui32(v536) {
		goto L151
	} else {
		goto L152
	}
L44:
	;
	v514 = F_cstring_to_text(m, int32(513958))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L150
	}
L45:
	;
	if v128 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v94)+52))
	v135 = F_has_privs_of_role(m, v133, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v94)+208))
	switch v139 {
	case 0:
		goto L52
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	case 5:
		goto L55
	case 6:
		goto L54
	case 7:
		goto L53
	default:
		goto L51
	}
L49:
	;
	if v135 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v94)+216))
	v171 = F_pgstat_clip_activity(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L67
	}
L52:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+292)) = uint8(v168)
	goto L51
L53:
	;
	v165 = F_cstring_to_text(m, int32(427750))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L66
	}
L54:
	;
	v161 = F_cstring_to_text(m, int32(628141))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L65
	}
L55:
	;
	v157 = F_cstring_to_text(m, int32(286210))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L64
	}
L56:
	;
	v153 = F_cstring_to_text(m, int32(241450))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L63
	}
L57:
	;
	v149 = F_cstring_to_text(m, int32(321698))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	v145 = F_cstring_to_text(m, int32(366772))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v141 = F_cstring_to_text(m, int32(308422))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v141
	goto L51
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v145
	goto L51
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v149
	goto L51
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v153
	goto L51
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v157
	goto L51
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v161
	goto L51
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = v165
	goto L51
L67:
	;
	v173 = F_cstring_to_text(m, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+340)) = v173
	F_pfree(m, v171)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+317)) = uint8(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v181 = F_BackendPidGetProc(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	if v422 != 0 {
		goto L119
	} else {
		goto L120
	}
L71:
	;
	v420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+294)) = uint8(v420)
	v422 = v398
	goto L70
L72:
	;
	if v181 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v185 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v186 == int32(1) {
		v398 = v185
		goto L71
	} else {
		goto L76
	}
L74:
	;
	v228 = v181
	goto L75
L75:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+548))
	if v229 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L76:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v190 = int32(0)
	if v189 == v190 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v224 == int32(0) {
		v398 = v185
		goto L71
	} else {
		goto L87
	}
L78:
	;
	v224 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v199 = v190
	goto L82
L81:
	;
	v224 = v219
	goto L77
L82:
	;
	v204 = v197 + v199*int32(640)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+44))
	if v205 == v189 {
		v219 = v204
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v224 = int32(0)
	goto L77
L84:
	;
	v211 = v197 + (v199|int32(1))*int32(640)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+44))
	if v212 == v189 {
		v219 = v211
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v215 = v199 + int32(2)
	if v215 != int32(38) {
		v199 = v215
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v228 = v224
	goto L75
L88:
	;
	v247 = F_pgstat_get_wait_event(m, v229)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L95
	}
L89:
	;
	v246 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v234 = v229 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v234) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v246 = int32(513571)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v234)>>(uint(int32(22))%32))&int32(1020))+uint32(_consts[1118])))
	v246 = v244
	goto L88
L95:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v228)+616))
	if v249 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	if v246 == int32(0) {
		v398 = v247
		goto L71
	} else {
		goto L116
	}
L97:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+317)) = uint8(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+436)) = v349
	goto L96
L98:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+44))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v250 != v251 {
		v349 = v250
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v254 != int32(5) {
		goto L96
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v260 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v264 = F_LWLockAcquire(m, v260+int32(5504), int32(1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v266 = int32(-1)
	v268 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	if v268 <= int32(0) {
		v322 = v266
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v339+int32(5504))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L114
	}
L105:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v279 = int32(0)
	goto L106
L106:
	;
	v299 = v272 + int32(16) + v279*int32(112)
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+16)))
	if v300 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v322 = v266
	goto L104
L108:
	;
	v314 = v279 + int32(1)
	if v314 != v268 {
		v279 = v314
		goto L106
	} else {
		goto L113
	}
L109:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v303 != int32(3) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	if v306 == int32(0) {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306)+44))
	if v257 != v309 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v299)+64))
	v322 = v311
	goto L104
L113:
	;
	goto L107
L114:
	;
	if v322 == int32(-1) {
		goto L96
	} else {
		goto L115
	}
L115:
	;
	v349 = v322
	goto L97
L116:
	;
	v395 = F_cstring_to_text(m, v246)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+344)) = v395
	v422 = v247
	goto L70
L118:
	;
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v94)+24))
	if v449 == int64(0) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v444 = F_cstring_to_text(m, v422)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v447 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+295)) = uint8(v447)
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+348)) = v444
	goto L118
L123:
	;
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v94)+32))
	if v460 != int64(0) {
		goto L129
	} else {
		goto L130
	}
L124:
	;
	v458 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+296)) = uint8(v458)
	goto L123
L125:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v452 == int32(6) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v455 = F_Int64GetDatum(m, v449)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+352)) = v455
	goto L123
L128:
	;
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v94)+16))
	if v468 != int64(0) {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v463 = F_Int64GetDatum(m, v460)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+297)) = uint8(v466)
	goto L128
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+356)) = v463
	goto L128
L133:
	;
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v94)+40))
	if v476 != int64(0) {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v471 = F_Int64GetDatum(m, v468)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+298)) = uint8(v474)
	goto L133
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+360)) = v471
	goto L133
L138:
	;
	v485 = v94 + int32(188)
	v487 = v94 + int32(56)
	if v487&int32(3) == int32(0) {
		goto L43
	} else {
		goto L143
	}
L139:
	;
	v479 = F_Int64GetDatum(m, v476)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v482 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+299)) = uint8(v482)
	goto L138
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+364)) = v479
	goto L138
L143:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+56)))
	if v492 != 0 {
		goto L42
	} else {
		goto L144
	}
L144:
	;
	v494 = v94 + int32(57)
	if v494&int32(3) == int32(0) {
		goto L43
	} else {
		goto L145
	}
L145:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	if v499 != 0 {
		goto L42
	} else {
		goto L146
	}
L146:
	;
	v501 = v94 + int32(58)
	if v501&int32(3) == int32(0) {
		goto L43
	} else {
		goto L147
	}
L147:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v506 != 0 {
		goto L42
	} else {
		goto L148
	}
L148:
	;
	if (v94-int32(1))&int32(3) == int32(0) {
		goto L43
	} else {
		goto L149
	}
L149:
	;
	goto L42
L150:
	;
	v516 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+292)) = uint8(v516)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+340)) = v514
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+8)) = uint8(v516)
	v521 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v521
	*(*int64)(unsafe.Add(mBase, uint32(v46)+6)) = v521
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v521
	goto L38
L151:
	;
	v538 = v532
	goto L154
L152:
	;
	v578 = v532
	goto L153
L153:
	;
	if base.Ui32(v578) < base.Ui32(v534) {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v560|(v561|(v562|(v563|(v564|(v565|(v566|v567)))))) != 0 {
		goto L42
	} else {
		goto L156
	}
L155:
	;
	v578 = v576
	goto L153
L156:
	;
	v576 = v538 + int32(32)
	if base.Ui32(v576) < base.Ui32(v536) {
		v538 = v576
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v601 = v578
	goto L161
L159:
	;
	v627 = v578
	goto L160
L160:
	;
	v649 = v627
	goto L165
L161:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	if v623 != 0 {
		goto L42
	} else {
		goto L163
	}
L162:
	;
	v627 = v625
	goto L160
L163:
	;
	v625 = v601 + int32(4)
	if base.Ui32(v625) < base.Ui32(v534) {
		v601 = v625
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	if base.Ui32(v485) <= base.Ui32(v649) {
		goto L41
	} else {
		goto L167
	}
L166:
	;
	goto L42
L167:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	if v672 == int32(0) {
		v649 = v649 + int32(1)
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+376)) = int32(-1)
	goto L40
L170:
	;
	v702 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)) = uint8(v702)
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v702)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v94)+184))
	v707 = int32(32)
	v712 = F_pg_getnameinfo_all(m, v487, v706, v25+v707, int32(255), v25, v707, int32(3))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	if v712 != 0 {
		goto L41
	} else {
		goto L172
	}
L172:
	;
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487))))
	if v714 != int32(10) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v730 = F_DirectFunctionCall1Coll(m, int32(1481), int32(0), v25+int32(32))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L177
	}
L174:
	;
	goto L173
L175:
	;
	v720 = F_strchr(m, v25+int32(32), int32(37))
	mBase = m.M
	if v720 == int32(0) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v723 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v720))) = uint8(v723)
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+368)) = v730
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v733 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v747 = v25
	goto L184
L179:
	;
	v742 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+301)) = uint8(v742)
	goto L178
L180:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	if v736 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v739 = F_cstring_to_text(m, v733)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+372)) = v739
	goto L178
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+376)) = v791
	goto L39
L184:
	;
	v752 = v747 + int32(1)
	v753 = int32(*(*int8)(unsafe.Add(mBase, uint32(v747))))
	v754 = F___isspace(m, v753)
	mBase = m.M
	if v754 != 0 {
		v747 = v752
		goto L184
	} else {
		goto L186
	}
L185:
	;
	v755 = int32(1)
	switch v753&int32(255) - int32(43) {
	case 0:
		v761 = v755
		goto L188
	default:
		v763 = v753
		v764 = v747
		v765 = v755
		goto L187
	case 2:
		goto L189
	}
L186:
	;
	goto L185
L187:
	;
	v766 = int32(0)
	v768 = v763 - int32(48)
	if base.Ui32(v768) <= base.Ui32(int32(9)) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	v762 = int32(*(*int8)(unsafe.Add(mBase, uint32(v752))))
	v763 = v762
	v764 = v752
	v765 = v761
	goto L187
L189:
	;
	v761 = int32(0)
	goto L188
L190:
	;
	v771 = v766
	v772 = v768
	v773 = v764
	goto L193
L191:
	;
	v785 = v766
	goto L192
L192:
	;
	if v765 != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v775 = int32(10)
	v777 = v771*v775 - v772
	v778 = int32(*(*int8)(unsafe.Add(mBase, uint32(v773)+1)))
	v782 = v778 - int32(48)
	if base.Ui32(v782) < base.Ui32(v775) {
		v771 = v777
		v772 = v782
		v773 = v773 + int32(1)
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v785 = v777
	goto L192
L195:
	;
	goto L194
L196:
	;
	v791 = int32(0) - v785
	goto L198
L197:
	;
	v791 = v785
	goto L198
L198:
	;
	goto L183
L199:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+192)))
	if v1063 == int32(1) {
		goto L244
	} else {
		goto L245
	}
L200:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v869 = int32(0)
	v871 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v875 = F_LWLockAcquire(m, v871+int32(4224), int32(1))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v865) <= base.Ui32(int32(17)) {
		goto L239
	} else {
		goto L240
	}
L203:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v878)))
	if v879 <= int32(0) {
		v995 = v869
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v1018+int32(4224))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L233
	}
L205:
	;
	v884 = v869
	goto L206
L206:
	;
	v908 = v878 + int32(16) + v884*int32(1480)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)+4))
	if v909 <= int32(0) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v995 = int32(0)
	goto L204
L208:
	;
	v992 = v884 + int32(1)
	if v992 != v879 {
		v884 = v992
		goto L206
	} else {
		goto L232
	}
L209:
	;
	if v909 != v868 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v913 = int32(4333184)
	v916 = v908 + int32(112)
	if (v916^v913)&int32(3) != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	v995 = v913
	goto L204
L212:
	;
	goto L211
L213:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v971))) = uint8(v970)
	if v970&int32(255) == int32(0) {
		goto L212
	} else {
		goto L228
	}
L214:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	v969 = v916
	v970 = v922
	v971 = v913
	goto L213
L215:
	;
	goto L216
L216:
	;
	if v916&int32(3) != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v926 = v916
	v928 = v913
	goto L220
L218:
	;
	v940 = v916
	v942 = v913
	goto L219
L219:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	v947 = int32(-2139062144)
	if (int32(16843008)-v944|v944)&v947 != v947 {
		v969 = v940
		v970 = v944
		v971 = v942
		goto L213
	} else {
		goto L224
	}
L220:
	;
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
	*(*uint8)(unsafe.Add(mBase, uint32(v928))) = uint8(v929)
	if v929 == int32(0) {
		goto L212
	} else {
		goto L222
	}
L221:
	;
	v940 = v936
	v942 = v934
	goto L219
L222:
	;
	v933 = int32(1)
	v934 = v928 + v933
	v936 = v926 + v933
	if v936&int32(3) != 0 {
		v926 = v936
		v928 = v934
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v952 = v940
	v953 = v944
	v954 = v942
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v954))) = v953
	v956 = int32(4)
	v957 = v954 + v956
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v952)+4))
	v960 = v952 + v956
	v964 = int32(-2139062144)
	if (v958|(int32(16843008)-v958))&v964 == v964 {
		v952 = v960
		v953 = v958
		v954 = v957
		goto L225
	} else {
		goto L227
	}
L226:
	;
	v969 = v960
	v970 = v958
	v971 = v957
	goto L213
L227:
	;
	goto L226
L228:
	;
	v978 = v969
	v980 = v971
	goto L229
L229:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v980)+1)) = uint8(v981)
	v983 = int32(1)
	if v981 != 0 {
		v978 = v978 + v983
		v980 = v980 + v983
		goto L229
	} else {
		goto L231
	}
L230:
	;
	goto L212
L231:
	;
	goto L230
L232:
	;
	goto L207
L233:
	;
	if v995 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1023 = F_cstring_to_text(m, v995)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1026 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+305)) = uint8(v1026)
	goto L199
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+388)) = v1023
	goto L199
L238:
	;
	v1038 = F_cstring_to_text(m, v1037)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L242
	}
L239:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v865<<(uint(int32(2))%32))+uint32(_consts[586])))
	v1037 = v1036
	goto L241
L240:
	;
	v1037 = int32(345732)
	goto L241
L241:
	;
	goto L238
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+388)) = v1038
	goto L199
L243:
	;
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+200)))
	if v1122 == int32(1) {
		goto L264
	} else {
		goto L265
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+392)) = int32(1)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v94)+196))
	v1071 = F_cstring_to_text(m, v1068+int32(4))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+392)) = int32(0)
	v1117 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v1117)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(16843009)
	goto L243
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+396)) = v1071
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v94)+196))
	v1077 = F_cstring_to_text(m, v1074+int32(68))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+400)) = v1077
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v94)+196))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+404)) = v1081
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1080)+132)))
	if v1083 != 0 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092)+196)))
	if v1093 != 0 {
		goto L255
	} else {
		goto L256
	}
L250:
	;
	v1086 = F_cstring_to_text(m, v1080+int32(132))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1090 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+310)) = uint8(v1090)
	v1092 = v1080
	goto L249
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+408)) = v1086
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v94)+196))
	v1092 = v1089
	goto L249
L254:
	;
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106)+260)))
	if v1107 != 0 {
		goto L259
	} else {
		goto L260
	}
L255:
	;
	v1095 = int32(0)
	v1100 = F_DirectFunctionCall3Coll(m, int32(408), v1095, v1092+int32(196), v1095, int32(-1))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+311)) = uint8(v1104)
	v1106 = v1092
	goto L254
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+412)) = v1100
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v94)+196))
	v1106 = v1103
	goto L254
L259:
	;
	v1110 = F_cstring_to_text(m, v1106+int32(260))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+312)) = uint8(v1113)
	goto L243
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+416)) = v1110
	goto L243
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+432)) = v1143
	v1145 = *(*int64)(unsafe.Add(mBase, uint32(v94)+392))
	if v1145 == int64(0) {
		goto L268
	} else {
		goto L269
	}
L264:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v94)+204))
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+64)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+420)) = v1126
	v1128 = F_cstring_to_text(m, v1125)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+314)) = uint8(v1135)
	v1137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+420)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v25)+428)) = v1137
	v1143 = v1137
	goto L263
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+424)) = v1128
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v94)+204))
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+65)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+428)) = v1132
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+66)))
	v1143 = v1134
	goto L263
L268:
	;
	v1148 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+318)) = uint8(v1148)
	goto L38
L269:
	;
	goto L270
L270:
	;
	v1150 = F_Int64GetDatum(m, v1145)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+440)) = v1150
	goto L38
L272:
	;
	if v36 != int32(-1) {
		goto L7
	} else {
		goto L273
	}
L273:
	;
	v1186 = v66 + int32(1)
	if v1186 <= v27 {
		v66 = v1186
		goto L9
	} else {
		goto L274
	}
L274:
	;
	goto L10
}
func F_pg_stat_get_backend_activity_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[31]))
			v18 = F_has_privs_of_role(m, v16, int32(3375))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
					if v29 == int64(0) {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						return int32(0)
					} else {
						v36 = F_Int64GetDatum(m, v29)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							return v36
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _consts[31]))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v23 = F_has_privs_of_role(m, v21, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 != 0 {
							v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
							if v29 == int64(0) {
								v32 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
								return int32(0)
							} else {
								v36 = F_Int64GetDatum(m, v29)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									return v36
								}
							}
						} else {
							v25 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_backend_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[31]))
			v18 = F_has_privs_of_role(m, v16, int32(3375))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)+16))
					if v29 == int64(0) {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						return int32(0)
					} else {
						v36 = F_Int64GetDatum(m, v29)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							return v36
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _consts[31]))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v23 = F_has_privs_of_role(m, v21, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 != 0 {
							v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)+16))
							if v29 == int64(0) {
								v32 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
								return int32(0)
							} else {
								v36 = F_Int64GetDatum(m, v29)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									return v36
								}
							}
						} else {
							v25 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_checkpointer_num_performed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+16))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_checkpointer_num_requested(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+8))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_checkpointer_restartpoints_performed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+40))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_db_blocks_hit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_conflict_startup_deadlock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+120))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_deadlocks(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+144))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_sessions_abandoned(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+216))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_sessions_fatal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+224))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_temp_files(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+128))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_xact_commit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_function_self_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pgstat_fetch_stat_funcentry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
			v18 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v14), float64(1000)))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_pg_stat_get_ins_since_vacuum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+96))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_lastscan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			if v9 != int64(0) {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		}
	}
}
func F_pg_stat_get_total_analyze_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Float8GetDatum(m, float64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+200))
			v15 = F_Float8GetDatum(m, base.F64_convert_i64_s(v13))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_stat_get_total_autovacuum_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Float8GetDatum(m, float64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+192))
			v15 = F_Float8GetDatum(m, base.F64_convert_i64_s(v13))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_stat_get_tuples_newpage_updated(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+64))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_tuples_returned(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_xact_function_calls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_find_funcstat_entry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v15 = F_Int64GetDatum(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_inserted(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+40))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
