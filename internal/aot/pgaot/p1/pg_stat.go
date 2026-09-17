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
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_force_next_flush[0])) = uint8(v3)
	return int32(0)
}
func F_pg_stat_get_activity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int64
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v351 int32
	_ = v351
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int64
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int64
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int64
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int64
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v634 int32
	_ = v634
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v769 int32
	_ = v769
	var v791 int32
	_ = v791
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1085 int32
	_ = v1085
	var v1087 int64
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	v21 = m.G0
	v23 = v21 - int32(448)
	m.G0 = v23
	v25 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v30 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = v33
	goto L5
L4:
	;
	v34 = int32(-1)
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v25 <= int32(0) {
		v1149 = v23
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v1149 + int32(448)
	return int32(0)
L8:
	;
	v53 = v23
	v60 = v34
	v61 = int32(1)
	v62 = v25
	v63 = v35
	v64 = v23 + int32(307)
	v65 = v23 + int32(305)
	v66 = v23 + int32(288) | int32(6)
	v67 = base.B2i32(v34 == int32(-1))
	goto L9
L9:
	;
	base.MemoryFill(m, v53+int32(320), int32(0), int32(124))
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+311)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v53)+304)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v53)+296)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v53)+288)) = v77
	v85 = F_pgstat_get_local_beentry_by_index(m, v61)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v1149 = v53
	goto L7
L11:
	;
	if v67 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v1146 = v61 + int32(1)
	if v1146 <= v62 {
		v61 = v1146
		goto L9
	} else {
		goto L278
	}
L13:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v89 != v60 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+48))
	if v91 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+324)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
	if v97 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+320)) = v91
	goto L17
L19:
	;
	goto L20
L20:
	;
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+288)) = uint8(v93)
	goto L17
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v85)+212))
	if v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+328)) = v97
	goto L21
L23:
	;
	goto L24
L24:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+290)) = uint8(v99)
	goto L21
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v85)+412))
	if v107 != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v102 = F_cstring_to_text(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+291)) = uint8(v105)
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+332)) = v102
	goto L25
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v85)+416))
	if v111 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+380)) = v107
	goto L30
L32:
	;
	goto L33
L33:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+303)) = uint8(v109)
	goto L30
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[0]))
	v118 = F_has_privs_of_role(m, v116, int32(3375))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L45
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+384)) = v111
	goto L34
L36:
	;
	goto L37
L37:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+304)) = uint8(v113)
	goto L34
L38:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
	F_tuplestore_putvalues(m, v1115, v1116, v53+int32(320), v53+int32(288))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L276
	}
L39:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v813 == int32(5) {
		goto L204
	} else {
		goto L205
	}
L40:
	;
	v791 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+300)) = uint16(v791)
	goto L39
L41:
	;
	v769 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+302)) = uint8(v769)
	goto L40
L42:
	;
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v462))))
	switch v657 - int32(1) {
	case 0:
		goto L173
	case 1, 9:
		goto L174
	default:
		goto L41
	}
L43:
	;
	v494 = v468 + v462
	v496 = v464 & int32(-4)
	v498 = v496 - int32(28)
	if base.Ui32(v494) < base.Ui32(v498) {
		goto L152
	} else {
		goto L153
	}
L44:
	;
	v481 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L151
	}
L45:
	;
	if v118 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[0]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
	v125 = F_has_privs_of_role(m, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v85)+208))
	switch v129 {
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
	if v125 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v85)+216))
	v161 = F_pgstat_clip_activity(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L67
	}
L52:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+292)) = uint8(v158)
	goto L51
L53:
	;
	v155 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L66
	}
L54:
	;
	v151 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L65
	}
L55:
	;
	v147 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L64
	}
L56:
	;
	v143 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_4))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L63
	}
L57:
	;
	v139 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_5))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	v135 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_6))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v131 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_activity_7))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v131
	goto L51
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v135
	goto L51
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v139
	goto L51
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v143
	goto L51
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v147
	goto L51
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v151
	goto L51
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v155
	goto L51
L67:
	;
	v163 = F_cstring_to_text(m, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+340)) = v163
	F_pfree(m, v161)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+317)) = uint8(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v171 = F_BackendPidGetProc(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	if v401 != 0 {
		goto L120
	} else {
		goto L121
	}
L71:
	;
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+294)) = uint8(v399)
	v401 = v379
	goto L70
L72:
	;
	if v171 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v175 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v176 == int32(1) {
		v379 = v175
		goto L71
	} else {
		goto L76
	}
L74:
	;
	v219 = v171
	goto L75
L75:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+548))
	if v220 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L76:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v180 = int32(0)
	if v179 == v180 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v215 == int32(0) {
		v379 = v175
		goto L71
	} else {
		goto L88
	}
L78:
	;
	v215 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[1]))
	v192 = v180
	goto L83
L81:
	;
	v215 = v209
	goto L77
L82:
	;
	v209 = v199 + int32(640)
	goto L81
L83:
	;
	v195 = v192 * int32(640)
	v196 = v188 + v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	if v197 == v179 {
		v209 = v196
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v215 = int32(0)
	goto L77
L85:
	;
	v199 = v188 + v195
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+684))
	if v200 == v179 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v203 = v192 + int32(2)
	if v203 != int32(38) {
		v192 = v203
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v219 = v215
	goto L75
L89:
	;
	v236 = F_pgstat_get_wait_event(m, v220)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L96
	}
L90:
	;
	v235 = int32(0)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v225 = v220 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v225) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v235 = int32(_a_F_pg_stat_get_activity_8)
	goto L89
L94:
	;
	goto L95
L95:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pg_stat_get_activity[2])))
	v235 = v233
	goto L89
L96:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v219)+616))
	if v238 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v235 == int32(0) {
		v379 = v236
		goto L71
	} else {
		goto L117
	}
L98:
	;
	v351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+317)) = uint8(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+436)) = v334
	goto L97
L99:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v239 != v240 {
		v334 = v239
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v243 != int32(5) {
		goto L97
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	v253 = F_LWLockAcquire(m, v249+int32(_a_F_pg_stat_get_activity_9), int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v255 = int32(-1)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[4]))
	if v257 <= int32(0) {
		v306 = v255
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	F_LWLockRelease(m, v324+int32(_a_F_pg_stat_get_activity_9))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L115
	}
L106:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[5]))
	v268 = int32(0)
	goto L107
L107:
	;
	v286 = v261 + int32(16) + v268*int32(112)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+16)))
	if v287 != int32(1) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v306 = v255
	goto L105
L109:
	;
	v301 = v268 + int32(1)
	if v301 != v257 {
		v268 = v301
		goto L107
	} else {
		goto L114
	}
L110:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	if v290 != int32(3) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v286)+20))
	if v293 == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v293)+44))
	if v246 != v296 {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v286)+64))
	v306 = v298
	goto L105
L114:
	;
	goto L108
L115:
	;
	if v306 == int32(-1) {
		goto L97
	} else {
		goto L116
	}
L116:
	;
	v334 = v306
	goto L98
L117:
	;
	v376 = F_cstring_to_text(m, v235)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+344)) = v376
	v401 = v236
	goto L70
L119:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
	if v426 == int64(0) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v421 = F_cstring_to_text(m, v401)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v424 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+295)) = uint8(v424)
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+348)) = v421
	goto L119
L124:
	;
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
	if v437 != int64(0) {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v435 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+296)) = uint8(v435)
	goto L124
L126:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v429 == int32(6) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v432 = F_Int64GetDatum(m, v426)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+352)) = v432
	goto L124
L129:
	;
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	if v445 != int64(0) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v440 = F_Int64GetDatum(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+297)) = uint8(v443)
	goto L129
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+356)) = v440
	goto L129
L134:
	;
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
	if v453 != int64(0) {
		goto L140
	} else {
		goto L141
	}
L135:
	;
	v448 = F_Int64GetDatum(m, v445)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+298)) = uint8(v451)
	goto L134
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+360)) = v448
	goto L134
L139:
	;
	v462 = v85 + int32(56)
	v464 = v85 + int32(188)
	v468 = (int32(-56) - v85) & int32(3)
	if v468 == int32(0) {
		goto L43
	} else {
		goto L144
	}
L140:
	;
	v456 = F_Int64GetDatum(m, v453)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v459 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+299)) = uint8(v459)
	goto L139
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+364)) = v456
	goto L139
L144:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v471 != 0 {
		goto L42
	} else {
		goto L145
	}
L145:
	;
	if v468 == int32(1) {
		goto L43
	} else {
		goto L146
	}
L146:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+57)))
	if v474 != 0 {
		goto L42
	} else {
		goto L147
	}
L147:
	;
	if v468 == int32(2) {
		goto L43
	} else {
		goto L148
	}
L148:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+58)))
	if v477 != 0 {
		goto L42
	} else {
		goto L149
	}
L149:
	;
	if v468 == int32(3) {
		goto L43
	} else {
		goto L150
	}
L150:
	;
	goto L42
L151:
	;
	v483 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+292)) = uint8(v483)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+340)) = v481
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+8)) = uint8(v483)
	v488 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v488
	*(*int64)(unsafe.Add(mBase, uint32(v65)+6)) = v488
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v488
	goto L38
L152:
	;
	v500 = v494
	v503 = v468
	goto L155
L153:
	;
	v542 = v468
	goto L154
L154:
	;
	v559 = v542 + v462
	if base.Ui32(v559) < base.Ui32(v496) {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v500)+28))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v500)+24))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v500)+20))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v500)+16))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v500)+12))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v500)+8))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	if v520|(v521|(v522|(v523|(v524|(v525|(v526|v527)))))) != 0 {
		goto L42
	} else {
		goto L157
	}
L156:
	;
	v542 = v536
	goto L154
L157:
	;
	v536 = v503 + int32(32)
	v537 = v462 + v536
	if base.Ui32(v537) < base.Ui32(v498) {
		v500 = v537
		v503 = v536
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v561 = v559
	v564 = v542
	goto L162
L160:
	;
	v589 = v542
	goto L161
L161:
	;
	v606 = int32(132)
	if base.Ui32(v589) <= base.Ui32(v606) {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	if v581 != 0 {
		goto L42
	} else {
		goto L164
	}
L163:
	;
	v589 = v583
	goto L161
L164:
	;
	v583 = v564 + int32(4)
	v584 = v462 + v583
	if base.Ui32(v584) < base.Ui32(v496) {
		v561 = v584
		v564 = v583
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v609 = v606
	goto L168
L167:
	;
	v609 = v589
	goto L168
L168:
	;
	v613 = v589
	goto L169
L169:
	;
	if v609 == v613 {
		goto L41
	} else {
		goto L171
	}
L170:
	;
	goto L42
L171:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613+v462))))
	if v634 == int32(0) {
		v613 = v613 + int32(1)
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+376)) = int32(-1)
	goto L40
L174:
	;
	v660 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+32)) = uint8(v660)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v660)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v85)+184))
	v665 = int32(32)
	v666 = v53 + v665
	v670 = F_pg_getnameinfo_all(m, v462, v664, v666, int32(255), v53, v665, int32(3))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	if v670 != 0 {
		goto L41
	} else {
		goto L176
	}
L176:
	;
	v672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v462))))
	if v672 != int32(10) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v684 = F_DirectFunctionCall1Coll(m, int32(1466), int32(0), v666)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L181
	}
L178:
	;
	goto L177
L179:
	;
	v676 = F_strchr(m, v666, int32(37))
	mBase = m.M
	if v676 == int32(0) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v679 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v679)
	goto L178
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+368)) = v684
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	if v687 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v701 = v53
	goto L188
L183:
	;
	v696 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+301)) = uint8(v696)
	goto L182
L184:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687))))
	if v690 == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v693 = F_cstring_to_text(m, v687)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+372)) = v693
	goto L182
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+376)) = v745
	goto L39
L188:
	;
	v706 = v701 + int32(1)
	v707 = int32(*(*int8)(unsafe.Add(mBase, uint32(v701))))
	v708 = F___isspace(m, v707)
	mBase = m.M
	if v708 != 0 {
		v701 = v706
		goto L188
	} else {
		goto L190
	}
L189:
	;
	v709 = int32(1)
	switch v707&int32(255) - int32(43) {
	case 0:
		v715 = v709
		goto L192
	default:
		v717 = v707
		v718 = v701
		v719 = v709
		goto L191
	case 2:
		goto L193
	}
L190:
	;
	goto L189
L191:
	;
	v720 = int32(0)
	v722 = v717 - int32(48)
	if base.Ui32(v722) <= base.Ui32(int32(9)) {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v716 = int32(*(*int8)(unsafe.Add(mBase, uint32(v706))))
	v717 = v716
	v718 = v706
	v719 = v715
	goto L191
L193:
	;
	v715 = int32(0)
	goto L192
L194:
	;
	v725 = v720
	v726 = v722
	v727 = v718
	goto L197
L195:
	;
	v739 = v720
	goto L196
L196:
	;
	if v719 != 0 {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	v729 = int32(10)
	v731 = v725*v729 - v726
	v732 = int32(*(*int8)(unsafe.Add(mBase, uint32(v727)+1)))
	v736 = v732 - int32(48)
	if base.Ui32(v736) < base.Ui32(v729) {
		v725 = v731
		v726 = v736
		v727 = v727 + int32(1)
		goto L197
	} else {
		goto L199
	}
L198:
	;
	v739 = v731
	goto L196
L199:
	;
	goto L198
L200:
	;
	v745 = int32(0) - v739
	goto L202
L201:
	;
	v745 = v739
	goto L202
L202:
	;
	goto L187
L203:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+192)))
	if v1005 == int32(1) {
		goto L248
	} else {
		goto L249
	}
L204:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v817 = int32(0)
	v819 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	v823 = F_LWLockAcquire(m, v819+int32(_a_F_pg_stat_get_activity_10), int32(1))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	if base.Ui32(v813) <= base.Ui32(int32(17)) {
		goto L243
	} else {
		goto L244
	}
L207:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[6]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	if v827 <= int32(0) {
		v944 = v817
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v965 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_activity[3]))
	F_LWLockRelease(m, v965+int32(_a_F_pg_stat_get_activity_10))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L237
	}
L209:
	;
	v832 = v817
	goto L210
L210:
	;
	v854 = v826 + int32(16) + v832*int32(1480)
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	v856 = int32(0)
	if base.B2i32(v855 <= v856)|base.B2i32(v816 != v855) == v856 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v944 = int32(0)
	goto L208
L212:
	;
	v862 = int32(_a_F_pg_stat_get_activity_11)
	v865 = v854 + int32(112)
	if (v865^v862)&int32(3) != 0 {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	goto L214
L214:
	;
	v941 = v832 + int32(1)
	if v941 != v827 {
		v832 = v941
		goto L210
	} else {
		goto L236
	}
L215:
	;
	v944 = v862
	goto L208
L216:
	;
	goto L215
L217:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v920))) = uint8(v919)
	if v919&int32(255) == int32(0) {
		goto L216
	} else {
		goto L232
	}
L218:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865))))
	v918 = v865
	v919 = v871
	v920 = v862
	goto L217
L219:
	;
	goto L220
L220:
	;
	if v865&int32(3) != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v875 = v865
	v877 = v862
	goto L224
L222:
	;
	v889 = v865
	v891 = v862
	goto L223
L223:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	v896 = int32(-2139062144)
	if (int32(16843008)-v893|v893)&v896 != v896 {
		v918 = v889
		v919 = v893
		v920 = v891
		goto L217
	} else {
		goto L228
	}
L224:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	*(*uint8)(unsafe.Add(mBase, uint32(v877))) = uint8(v878)
	if v878 == int32(0) {
		goto L216
	} else {
		goto L226
	}
L225:
	;
	v889 = v885
	v891 = v883
	goto L223
L226:
	;
	v882 = int32(1)
	v883 = v877 + v882
	v885 = v875 + v882
	if v885&int32(3) != 0 {
		v875 = v885
		v877 = v883
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v901 = v889
	v902 = v893
	v903 = v891
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v903))) = v902
	v905 = int32(4)
	v906 = v903 + v905
	v908 = v901 + v905
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	v913 = int32(-2139062144)
	if (int32(16843008)-v910|v910)&v913 == v913 {
		v901 = v908
		v902 = v910
		v903 = v906
		goto L229
	} else {
		goto L231
	}
L230:
	;
	v918 = v908
	v919 = v910
	v920 = v906
	goto L217
L231:
	;
	goto L230
L232:
	;
	v927 = v918
	v929 = v920
	goto L233
L233:
	;
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v929)+1)) = uint8(v930)
	v932 = int32(1)
	if v930 != 0 {
		v927 = v927 + v932
		v929 = v929 + v932
		goto L233
	} else {
		goto L235
	}
L234:
	;
	goto L216
L235:
	;
	goto L234
L236:
	;
	goto L211
L237:
	;
	if v944 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v970 = F_cstring_to_text(m, v944)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v973 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+305)) = uint8(v973)
	goto L203
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+388)) = v970
	goto L203
L242:
	;
	v982 = F_cstring_to_text(m, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L246
	}
L243:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v813<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_activity[7])))
	v981 = v979
	goto L245
L244:
	;
	v981 = int32(_a_F_pg_stat_get_activity_12)
	goto L245
L245:
	;
	goto L242
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+388)) = v982
	goto L203
L247:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+200)))
	if v1064 == int32(1) {
		goto L268
	} else {
		goto L269
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+392)) = int32(1)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v85)+196))
	v1013 = F_cstring_to_text(m, v1010+int32(4))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+392)) = int32(0)
	v1059 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v1059)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(16843009)
	goto L247
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+396)) = v1013
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v85)+196))
	v1019 = F_cstring_to_text(m, v1016+int32(68))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+400)) = v1019
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v85)+196))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1022)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+404)) = v1023
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022)+132)))
	if v1025 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+196)))
	if v1035 != 0 {
		goto L259
	} else {
		goto L260
	}
L254:
	;
	v1028 = F_cstring_to_text(m, v1022+int32(132))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1032 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+310)) = uint8(v1032)
	v1034 = v1022
	goto L253
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+408)) = v1028
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v85)+196))
	v1034 = v1031
	goto L253
L258:
	;
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048)+260)))
	if v1049 != 0 {
		goto L263
	} else {
		goto L264
	}
L259:
	;
	v1037 = int32(0)
	v1042 = F_DirectFunctionCall3Coll(m, int32(408), v1037, v1034+int32(196), v1037, int32(-1))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1046 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+311)) = uint8(v1046)
	v1048 = v1034
	goto L258
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+412)) = v1042
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v85)+196))
	v1048 = v1045
	goto L258
L263:
	;
	v1052 = F_cstring_to_text(m, v1048+int32(260))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1055 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+312)) = uint8(v1055)
	goto L247
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+416)) = v1052
	goto L247
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+432)) = v1085
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v85)+392))
	if v1087 == int64(0) {
		goto L272
	} else {
		goto L273
	}
L268:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v85)+204))
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+64)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+420)) = v1068
	v1070 = F_cstring_to_text(m, v1067)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1077 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+314)) = uint8(v1077)
	v1079 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+420)) = v1079
	*(*int32)(unsafe.Add(mBase, uint32(v53)+428)) = v1079
	v1085 = v1079
	goto L267
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+424)) = v1070
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v85)+204))
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073)+65)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+428)) = v1074
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073)+66)))
	v1085 = v1076
	goto L267
L272:
	;
	v1090 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+318)) = uint8(v1090)
	goto L38
L273:
	;
	goto L274
L274:
	;
	v1092 = F_Int64GetDatum(m, v1087)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+440)) = v1092
	goto L38
L276:
	;
	if v67 == int32(0) {
		v1149 = v53
		goto L7
	} else {
		goto L277
	}
L277:
	;
	goto L12
L278:
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity_start[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
					if v21 == int64(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						return int32(0)
					} else {
						v24 = F_Int64GetDatum(m, v21)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v24
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity_start[0]))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v19 = F_has_privs_of_role(m, v17, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						if v19 != 0 {
							v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
							if v21 == int64(0) {
								v28 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
								return int32(0)
							} else {
								v24 = F_Int64GetDatum(m, v21)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return int32(0)
								} else {
									return v24
								}
							}
						} else {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_start[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+16))
					if v21 == int64(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						return int32(0)
					} else {
						v24 = F_Int64GetDatum(m, v21)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v24
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_start[0]))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v19 = F_has_privs_of_role(m, v17, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						if v19 != 0 {
							v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+16))
							if v21 == int64(0) {
								v28 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
								return int32(0)
							} else {
								v24 = F_Int64GetDatum(m, v21)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return int32(0)
								} else {
									return v24
								}
							}
						} else {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
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
func F_pg_stat_statements_1_10(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(6), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
