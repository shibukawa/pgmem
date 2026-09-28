package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_force_next_flush(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_force_next_flush[0])) = uint8(v3)
	return int64(0)
}
func F_pg_stat_get_activity(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v82 int64
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v372 int32
	_ = v372
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int64
	_ = v462
	var v466 int32
	_ = v466
	var v468 int64
	_ = v468
	var v472 int32
	_ = v472
	var v474 int64
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int64
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v659 int32
	_ = v659
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v710 int64
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v798 int32
	_ = v798
	var v821 int32
	_ = v821
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1082 int64
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int64
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int64
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int64
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int64
	_ = v1123
	var v1128 int64
	_ = v1128
	var v1130 int64
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	v22 = m.G0
	v24 = v22 - int32(576)
	m.G0 = v24
	v26 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v31 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v35 = v34
	goto L5
L4:
	;
	v35 = int32(-1)
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v26 <= int32(0) {
		v1192 = v24
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v1192 + int32(576)
	return int64(0)
L8:
	;
	v57 = v24
	v64 = v35
	v65 = int32(1)
	v66 = v26
	v67 = v36
	v68 = v24 + int32(307)
	v69 = v24 + int32(305)
	v70 = v24 + int32(288) | int32(6)
	v71 = base.B2i32(v35 == int32(-1))
	v76 = base.I64_extend_i32_u(v24 + int32(32))
	goto L9
L9:
	;
	base.MemoryFill(m, v57+int32(320), int32(0), int32(248))
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+311)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v57)+304)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v57)+296)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v57)+288)) = v82
	v90 = F_pgstat_get_local_beentry_by_index(m, v65)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v1192 = v57
	goto L7
L11:
	;
	if v71 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v1189 = v65 + int32(1)
	if v1189 <= v66 {
		v65 = v1189
		goto L9
	} else {
		goto L273
	}
L13:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v94 != v64 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+48))
	if v96 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v101 = int64(*(*int32)(unsafe.Add(mBase, uint32(v90)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+328)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v90)+52))
	if v103 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+320)) = base.I64_extend_i32_u(v96)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+288)) = uint8(v99)
	goto L17
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+212))
	if v108 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+336)) = base.I64_extend_i32_u(v103)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+290)) = uint8(v106)
	goto L21
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+412))
	if v115 != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v109 = F_cstring_to_text(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+291)) = uint8(v113)
	goto L25
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+344)) = base.I64_extend_i32_u(v109)
	goto L25
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v90)+416))
	if v120 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+440)) = base.I64_extend_i32_u(v115)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v118 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+303)) = uint8(v118)
	goto L30
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[0]))
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
	*(*int64)(unsafe.Add(mBase, uint32(v57)+448)) = base.I64_extend_i32_u(v120)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+304)) = uint8(v123)
	goto L34
L38:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	F_tuplestore_putvalues(m, v1157, v1158, v57+int32(320), v57+int32(288))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L271
	}
L39:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v844 == int32(5) {
		goto L200
	} else {
		goto L201
	}
L40:
	;
	v821 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+300)) = uint16(v821)
	goto L39
L41:
	;
	v798 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+302)) = uint8(v798)
	goto L40
L42:
	;
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v481))))
	switch v683 - int32(1) {
	case 0:
		goto L169
	case 1, 9:
		goto L170
	default:
		goto L41
	}
L43:
	;
	v514 = v487 + v481
	v516 = v483 & int32(-4)
	v518 = v516 - int32(28)
	if base.Ui32(v514) < base.Ui32(v518) {
		goto L148
	} else {
		goto L149
	}
L44:
	;
	v500 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L147
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
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[0]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v90)+52))
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v90)+208))
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
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v90)+216))
	v178 = F_pgstat_clip_activity(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L67
	}
L52:
	;
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+292)) = uint8(v175)
	goto L51
L53:
	;
	v171 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L66
	}
L54:
	;
	v166 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L65
	}
L55:
	;
	v161 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L64
	}
L56:
	;
	v156 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L63
	}
L57:
	;
	v151 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	v146 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_6))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v141 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_7))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v141)
	goto L51
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v146)
	goto L51
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v151)
	goto L51
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v156)
	goto L51
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v161)
	goto L51
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v166)
	goto L51
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+352)) = base.I64_extend_i32_u(v171)
	goto L51
L67:
	;
	v180 = F_cstring_to_text(m, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+360)) = base.I64_extend_i32_u(v180)
	F_pfree(m, v178)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+317)) = uint8(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v189 = F_BackendPidGetProc(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	if v426 != 0 {
		goto L120
	} else {
		goto L121
	}
L71:
	;
	v424 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+294)) = uint8(v424)
	v426 = v403
	goto L70
L72:
	;
	if v189 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v193 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v194 == int32(1) {
		v403 = v193
		goto L71
	} else {
		goto L76
	}
L74:
	;
	v237 = v189
	goto L75
L75:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+648))
	if v238 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L76:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v198 = int32(0)
	if v197 == v198 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v233 == int32(0) {
		v403 = v193
		goto L71
	} else {
		goto L88
	}
L78:
	;
	v233 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[1]))
	v210 = v198
	goto L83
L81:
	;
	v233 = v227
	goto L77
L82:
	;
	v227 = v217 + int32(768)
	goto L81
L83:
	;
	v213 = v210 * int32(768)
	v214 = v206 + v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	if v215 == v197 {
		v227 = v214
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v233 = int32(0)
	goto L77
L85:
	;
	v217 = v206 + v213
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+780))
	if v218 == v197 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v221 = v210 + int32(2)
	if v221 != int32(38) {
		v210 = v221
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v237 = v233
	goto L75
L89:
	;
	v254 = F_pgstat_get_wait_event(m, v238)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L96
	}
L90:
	;
	v253 = int32(0)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v243 = v238 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v243) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v253 = int32(_a_F_pg_stat_get_activity_8)
	goto L89
L94:
	;
	goto L95
L95:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v243)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pg_stat_get_activity[2])))
	v253 = v251
	goto L89
L96:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v237)+364))
	if v256 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v253 == int32(0) {
		v403 = v254
		goto L71
	} else {
		goto L117
	}
L98:
	;
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+317)) = uint8(v372)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+552)) = base.I64_extend_i32_s(v354)
	goto L97
L99:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v257 != v258 {
		v354 = v257
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v261 != int32(5) {
		goto L97
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	v271 = F_LWLockAcquire(m, v267+int32(_a_F_pg_stat_get_activity_9), int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v273 = int32(-1)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[4]))
	if v275 <= int32(0) {
		v325 = v273
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	F_LWLockRelease(m, v344+int32(_a_F_pg_stat_get_activity_9))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L115
	}
L106:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[5]))
	v286 = int32(0)
	goto L107
L107:
	;
	v305 = v279 + int32(16) + v286<<(uint(int32(7))%32)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+16)))
	if v306 != int32(1) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v325 = v273
	goto L105
L109:
	;
	v320 = v286 + int32(1)
	if v320 != v275 {
		v286 = v320
		goto L107
	} else {
		goto L114
	}
L110:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v309 != int32(4) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v305)+20))
	if v312 == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	if v264 != v315 {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v305)+64))
	v325 = v317
	goto L105
L114:
	;
	goto L108
L115:
	;
	if v325 == int32(-1) {
		goto L97
	} else {
		goto L116
	}
L116:
	;
	v354 = v325
	goto L98
L117:
	;
	v399 = F_cstring_to_text(m, v253)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+368)) = base.I64_extend_i32_u(v399)
	v426 = v254
	goto L70
L119:
	;
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v90)+24))
	if v453 == int64(0) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v447 = F_cstring_to_text(m, v426)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+295)) = uint8(v451)
	goto L119
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+376)) = base.I64_extend_i32_u(v447)
	goto L119
L124:
	;
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v90)+32))
	if v462 != int64(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+296)) = uint8(v460)
	goto L124
L126:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v456 == int32(6) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+384)) = v453
	goto L124
L128:
	;
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v90)+16))
	if v468 != int64(0) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+392)) = v462
	goto L128
L130:
	;
	goto L131
L131:
	;
	v466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+297)) = uint8(v466)
	goto L128
L132:
	;
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v90)+40))
	if v474 != int64(0) {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+400)) = v468
	goto L132
L134:
	;
	goto L135
L135:
	;
	v472 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+298)) = uint8(v472)
	goto L132
L136:
	;
	v481 = v90 + int32(56)
	v483 = v90 + int32(188)
	v487 = (int32(-56) - v90) & int32(3)
	if v487 == int32(0) {
		goto L43
	} else {
		goto L140
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+408)) = v474
	goto L136
L138:
	;
	goto L139
L139:
	;
	v478 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+299)) = uint8(v478)
	goto L136
L140:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v490 != 0 {
		goto L42
	} else {
		goto L141
	}
L141:
	;
	if v487 == int32(1) {
		goto L43
	} else {
		goto L142
	}
L142:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+57)))
	if v493 != 0 {
		goto L42
	} else {
		goto L143
	}
L143:
	;
	if v487 == int32(2) {
		goto L43
	} else {
		goto L144
	}
L144:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+58)))
	if v496 != 0 {
		goto L42
	} else {
		goto L145
	}
L145:
	;
	if v487 == int32(3) {
		goto L43
	} else {
		goto L146
	}
L146:
	;
	goto L42
L147:
	;
	v502 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+292)) = uint8(v502)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+360)) = base.I64_extend_i32_u(v500)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v502)
	v508 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v69)+6)) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v508
	goto L38
L148:
	;
	v520 = v514
	v523 = v487
	goto L151
L149:
	;
	v563 = v487
	goto L150
L150:
	;
	v581 = v563 + v481
	if base.Ui32(v581) < base.Ui32(v516) {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v520)+28))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v520)+24))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v520)+20))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v520)+16))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v520)+12))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v520)+8))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	if v541|(v542|(v543|(v544|(v545|(v546|(v547|v548)))))) != 0 {
		goto L42
	} else {
		goto L153
	}
L152:
	;
	v563 = v557
	goto L150
L153:
	;
	v557 = v523 + int32(32)
	v558 = v481 + v557
	if base.Ui32(v558) < base.Ui32(v518) {
		v520 = v558
		v523 = v557
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v583 = v581
	v586 = v563
	goto L158
L156:
	;
	v612 = v563
	goto L157
L157:
	;
	v630 = int32(132)
	if base.Ui32(v612) <= base.Ui32(v630) {
		goto L162
	} else {
		goto L163
	}
L158:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	if v604 != 0 {
		goto L42
	} else {
		goto L160
	}
L159:
	;
	v612 = v606
	goto L157
L160:
	;
	v606 = v586 + int32(4)
	v607 = v481 + v606
	if base.Ui32(v607) < base.Ui32(v516) {
		v583 = v607
		v586 = v606
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v633 = v630
	goto L164
L163:
	;
	v633 = v612
	goto L164
L164:
	;
	v637 = v612
	goto L165
L165:
	;
	if v633 == v637 {
		goto L41
	} else {
		goto L167
	}
L166:
	;
	goto L42
L167:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v481))))
	if v659 == int32(0) {
		v637 = v637 + int32(1)
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+432)) = int64(-1)
	goto L40
L170:
	;
	v686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+32)) = uint8(v686)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v686)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v90)+184))
	v691 = int32(32)
	v692 = v57 + v691
	v696 = F_pg_getnameinfo_all(m, v481, v690, v692, int32(255), v57, v691, int32(3))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	if v696 != 0 {
		goto L41
	} else {
		goto L172
	}
L172:
	;
	v698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v481))))
	if v698 != int32(10) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v710 = F_DirectFunctionCall1Coll(m, int32(1678), int32(0), v76)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L177
	}
L174:
	;
	goto L173
L175:
	;
	v702 = F_strchr(m, v692, int32(37))
	mBase = m.M
	if v702 == int32(0) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v705 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v702))) = uint8(v705)
	goto L174
L177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+416)) = v710
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	if v713 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v728 = v57
	goto L184
L179:
	;
	v723 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+301)) = uint8(v723)
	goto L178
L180:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v716 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v719 = F_cstring_to_text(m, v713)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+424)) = base.I64_extend_i32_u(v719)
	goto L178
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+432)) = base.I64_extend_i32_s(v772)
	goto L39
L184:
	;
	v733 = v728 + int32(1)
	v734 = int32(*(*int8)(unsafe.Add(mBase, uint32(v728))))
	v735 = F___isspace(m, v734)
	mBase = m.M
	if v735 != 0 {
		v728 = v733
		goto L184
	} else {
		goto L186
	}
L185:
	;
	v736 = int32(1)
	switch v734&int32(255) - int32(43) {
	case 0:
		v742 = v736
		goto L188
	default:
		v744 = v734
		v745 = v728
		v746 = v736
		goto L187
	case 2:
		goto L189
	}
L186:
	;
	goto L185
L187:
	;
	v747 = int32(0)
	v749 = v744 - int32(48)
	if base.Ui32(v749) <= base.Ui32(int32(9)) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	v743 = int32(*(*int8)(unsafe.Add(mBase, uint32(v733))))
	v744 = v743
	v745 = v733
	v746 = v742
	goto L187
L189:
	;
	v742 = int32(0)
	goto L188
L190:
	;
	v752 = v747
	v753 = v749
	v754 = v745
	goto L193
L191:
	;
	v766 = v747
	goto L192
L192:
	;
	if v746 != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v756 = int32(10)
	v758 = v752*v756 - v753
	v759 = int32(*(*int8)(unsafe.Add(mBase, uint32(v754)+1)))
	v763 = v759 - int32(48)
	if base.Ui32(v763) < base.Ui32(v756) {
		v752 = v758
		v753 = v763
		v754 = v754 + int32(1)
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v766 = v758
	goto L192
L195:
	;
	goto L194
L196:
	;
	v772 = int32(0) - v766
	goto L198
L197:
	;
	v772 = v766
	goto L198
L198:
	;
	goto L183
L199:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+192)))
	if v1041 == int32(1) {
		goto L244
	} else {
		goto L245
	}
L200:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v848 = int32(0)
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	v854 = F_LWLockAcquire(m, v850+int32(_a_F_pg_stat_get_activity_10), int32(1))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v844) <= base.Ui32(int32(17)) {
		goto L239
	} else {
		goto L240
	}
L203:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[6]))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v857)))
	if v858 <= int32(0) {
		v976 = v848
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	F_LWLockRelease(m, v998+int32(_a_F_pg_stat_get_activity_10))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L233
	}
L205:
	;
	v863 = v848
	goto L206
L206:
	;
	v886 = v857 + int32(16) + v863*int32(1488)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)+4))
	v888 = int32(0)
	if base.B2i32(v887 <= v888)|base.B2i32(v847 != v887) == v888 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v976 = int32(0)
	goto L204
L208:
	;
	v894 = int32(_a_F_pg_stat_get_activity_11)
	v897 = v886 + int32(112)
	if (v897^v894)&int32(3) != 0 {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	goto L210
L210:
	;
	v973 = v863 + int32(1)
	if v973 != v858 {
		v863 = v973
		goto L206
	} else {
		goto L232
	}
L211:
	;
	v976 = v894
	goto L204
L212:
	;
	goto L211
L213:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v952))) = uint8(v951)
	if v951&int32(255) == int32(0) {
		goto L212
	} else {
		goto L228
	}
L214:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897))))
	v950 = v897
	v951 = v903
	v952 = v894
	goto L213
L215:
	;
	goto L216
L216:
	;
	if v897&int32(3) != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v907 = v897
	v909 = v894
	goto L220
L218:
	;
	v921 = v897
	v923 = v894
	goto L219
L219:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	v928 = int32(-2139062144)
	if (int32(16843008)-v925|v925)&v928 != v928 {
		v950 = v921
		v951 = v925
		v952 = v923
		goto L213
	} else {
		goto L224
	}
L220:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907))))
	*(*uint8)(unsafe.Add(mBase, uint32(v909))) = uint8(v910)
	if v910 == int32(0) {
		goto L212
	} else {
		goto L222
	}
L221:
	;
	v921 = v917
	v923 = v915
	goto L219
L222:
	;
	v914 = int32(1)
	v915 = v909 + v914
	v917 = v907 + v914
	if v917&int32(3) != 0 {
		v907 = v917
		v909 = v915
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v933 = v921
	v934 = v925
	v935 = v923
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v935))) = v934
	v937 = int32(4)
	v938 = v935 + v937
	v940 = v933 + v937
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	v945 = int32(-2139062144)
	if (int32(16843008)-v942|v942)&v945 == v945 {
		v933 = v940
		v934 = v942
		v935 = v938
		goto L225
	} else {
		goto L227
	}
L226:
	;
	v950 = v940
	v951 = v942
	v952 = v938
	goto L213
L227:
	;
	goto L226
L228:
	;
	v959 = v950
	v961 = v952
	goto L229
L229:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v961)+1)) = uint8(v962)
	v964 = int32(1)
	if v962 != 0 {
		v959 = v959 + v964
		v961 = v961 + v964
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
	if v976 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1003 = F_cstring_to_text(m, v976)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1007 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+305)) = uint8(v1007)
	goto L199
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+456)) = base.I64_extend_i32_u(v1003)
	goto L199
L238:
	;
	v1016 = F_cstring_to_text(m, v1015)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L1
	} else {
		goto L242
	}
L239:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v844<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_activity[7])))
	v1015 = v1013
	goto L241
L240:
	;
	v1015 = int32(_a_F_pg_stat_get_activity_12)
	goto L241
L241:
	;
	goto L238
L242:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+456)) = base.I64_extend_i32_u(v1016)
	goto L199
L243:
	;
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+200)))
	if v1105 == int32(1) {
		goto L264
	} else {
		goto L265
	}
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+464)) = int64(1)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v90)+196))
	v1049 = F_cstring_to_text(m, v1046+int32(4))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+464)) = int64(0)
	v1100 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)) = uint16(v1100)
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(16843009)
	goto L243
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+472)) = base.I64_extend_i32_u(v1049)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v90)+196))
	v1056 = F_cstring_to_text(m, v1053+int32(68))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+480)) = base.I64_extend_i32_u(v1056)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v90)+196))
	v1061 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1060))))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+488)) = v1061
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060)+132)))
	if v1063 != 0 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073)+196)))
	if v1074 != 0 {
		goto L255
	} else {
		goto L256
	}
L250:
	;
	v1066 = F_cstring_to_text(m, v1060+int32(132))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1071 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+310)) = uint8(v1071)
	v1073 = v1060
	goto L249
L253:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+496)) = base.I64_extend_i32_u(v1066)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v90)+196))
	v1073 = v1070
	goto L249
L254:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1088)+260)))
	if v1089 != 0 {
		goto L259
	} else {
		goto L260
	}
L255:
	;
	v1082 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v1073+int32(196)), int64(0), int64(-1))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1086 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+311)) = uint8(v1086)
	v1088 = v1073
	goto L254
L258:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+504)) = v1082
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v90)+196))
	v1088 = v1085
	goto L254
L259:
	;
	v1092 = F_cstring_to_text(m, v1088+int32(260))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1096 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+312)) = uint8(v1096)
	goto L243
L262:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+512)) = base.I64_extend_i32_u(v1092)
	goto L243
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+544)) = v1128
	v1130 = *(*int64)(unsafe.Add(mBase, uint32(v90)+392))
	if v1130 == int64(0) {
		goto L268
	} else {
		goto L269
	}
L264:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+204))
	v1109 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+64)))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+520)) = v1109
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v90)+204))
	v1112 = F_cstring_to_text(m, v1111)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+314)) = uint8(v1121)
	v1123 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+520)) = v1123
	*(*int64)(unsafe.Add(mBase, uint32(v57)+536)) = v1123
	v1128 = v1123
	goto L263
L267:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+528)) = base.I64_extend_i32_u(v1112)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v90)+204))
	v1117 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1116)+65)))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+536)) = v1117
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v90)+204))
	v1120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1119)+66)))
	v1128 = v1120
	goto L263
L268:
	;
	v1133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+318)) = uint8(v1133)
	goto L38
L269:
	;
	goto L270
L270:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57)+560)) = v1130
	goto L38
L271:
	;
	if v71 == int32(0) {
		v1192 = v57
		goto L7
	} else {
		goto L272
	}
L272:
	;
	goto L12
L273:
	;
	goto L10
}
func F_pg_stat_get_backend_activity_start(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v29 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
			v32 = int64(0)
			return v32
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity_start[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				if v14 == int32(0) {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity_start[0]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v21 = F_has_privs_of_role(m, v19, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						if v21 == int32(0) {
							v29 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
							v32 = int64(0)
						} else {
							v25 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
							if v25 != int64(0) {
								v32 = v25
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
								v32 = int64(0)
							}
						}
						return v32
					}
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
					if v25 != int64(0) {
						v32 = v25
					} else {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
						v32 = int64(0)
					}
					return v32
				}
			}
		}
	}
}
func F_pg_stat_get_backend_start(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v29 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
			v32 = int64(0)
			return v32
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_start[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				if v14 == int32(0) {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_start[0]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v21 = F_has_privs_of_role(m, v19, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						if v21 == int32(0) {
							v29 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
							v32 = int64(0)
						} else {
							v25 = *(*int64)(unsafe.Add(mBase, uint32(v5)+16))
							if v25 != int64(0) {
								v32 = v25
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
								v32 = int64(0)
							}
						}
						return v32
					}
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(v5)+16))
					if v25 != int64(0) {
						v32 = v25
					} else {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
						v32 = int64(0)
					}
					return v32
				}
			}
		}
	}
}
func F_pg_stat_get_checkpointer_num_performed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+16))
		return v6
	}
}
func F_pg_stat_get_checkpointer_num_requested(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+8))
		return v6
	}
}
func F_pg_stat_get_checkpointer_restartpoints_performed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+40))
		return v6
	}
}
func F_pg_stat_get_db_blocks_hit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
			return v11
		}
	}
}
func F_pg_stat_get_db_conflict_startup_deadlock(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+120))
			return v11
		}
	}
}
func F_pg_stat_get_db_deadlocks(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+144))
			return v11
		}
	}
}
func F_pg_stat_get_db_sessions_abandoned(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+216))
			return v11
		}
	}
}
func F_pg_stat_get_db_sessions_fatal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+224))
			return v11
		}
	}
}
func F_pg_stat_get_db_temp_files(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+128))
			return v11
		}
	}
}
func F_pg_stat_get_db_xact_commit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
			return v11
		}
	}
}
func F_pg_stat_get_function_self_time(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pgstat_fetch_stat_funcentry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v14), float64(1000)))
		}
	}
}
func F_pg_stat_get_function_stat_reset_time(m *base.Module, l0 int32) int64 {
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
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_funcentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_ins_since_vacuum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+96))
			return v11
		}
	}
}
func F_pg_stat_get_lastscan(m *base.Module, l0 int32) int64 {
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
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_total_analyze_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+200))
			return base.I64_reinterpret_f64(base.F64_convert_i64_s(v11))
		}
	}
}
func F_pg_stat_get_total_autovacuum_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+192))
			return base.I64_reinterpret_f64(base.F64_convert_i64_s(v11))
		}
	}
}
func F_pg_stat_get_tuples_newpage_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+64))
			return v11
		}
	}
}
func F_pg_stat_get_tuples_returned(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
			return v11
		}
	}
}
func F_pg_stat_get_xact_function_calls(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_find_funcstat_entry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			return v14
		}
	}
}
func F_pg_stat_get_xact_tuples_inserted(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+40))
			return v11
		}
	}
}
func F_pg_stat_statements_1_10(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(6), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
