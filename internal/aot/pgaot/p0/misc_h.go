package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HaveVirtualXIDsDelayingChkpt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[1]))
	v19 = F_LWLockAcquire(m, v15+int32(512), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v24 <= v23 {
		v109 = v23
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[1]))
	F_LWLockRelease(m, v111+int32(512))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L17
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[2]))
	v35 = int32(0)
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v35<<(uint(int32(2))%32))))
	v48 = v30 + v45*int32(640)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+120))
	if v49&l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v109 = int32(0)
	goto L3
L7:
	;
	v95 = v35 + int32(1)
	if v95 != v24 {
		v35 = v95
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	v54 = int32(0)
	if base.B2i32(v53 == v54)|base.B2i32(l1 <= v54) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v64 = int32(0)
	goto L10
L10:
	;
	v74 = l0 + v64<<(uint(int32(3))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v59 != v75 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	v81 = v64 + int32(1)
	if v81 != l1 {
		v64 = v81
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v53 != v77 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v109 = int32(1)
	goto L3
L15:
	;
	goto L11
L16:
	;
	goto L6
L17:
	;
	return v109
}
func F_handle_pm_shutdown_request_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	switch l0 - int32(2) {
	case 0:
		v7 = int32(_a_F_handle_pm_shutdown_request_signal_0)
		goto L3
	case 1:
		goto L4
	default:
		goto L1
	case 13:
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[1])) = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	goto L2
L4:
	;
	v7 = int32(_a_F_handle_pm_shutdown_request_signal_1)
	goto L3
L5:
	;
	return
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v23 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[2]))
	if v27 == v23 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[3]))
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v57 = F_pgmem_kill(m, v23, int32(23))
	mBase = m.M
	goto L6
L13:
	;
	m.G0 = v31 + int32(16)
	goto L5
L14:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+15)) = uint8(v37)
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[4]))
	v45 = F_write(m, v41, v31+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v45 {
		goto L13
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[5]))
	if v49 == int32(27) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_has_bypassrls_privilege(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v11 = F_SearchSysCache1(m, int32(11), l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					return int32(0)
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v18)+74)))
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = v20
						return v24 & int32(1)
					}
				}
			}
		} else {
			v24 = int32(1)
			return v24 & int32(1)
		}
	}
}
func F_has_superclass(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13922(m, l0, int32(2680), int32(1), int32(2611))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_hashadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	Fn13923(m, l0, l1, l2, l3, int32(405))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_hashbool(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(711645284)
	v12 = base.B2i32(v2 != int32(0)) - int32(1636608428) ^ v9 - int32(1455628627)
	v17 = v12 ^ int32(-1636608428) - base.I32_rotl(v12, int32(25))
	v22 = v17 ^ v9 - base.I32_rotl(v17, int32(16))
	v26 = v22 ^ v12 - base.I32_rotl(v22, int32(4))
	v30 = v26 ^ v17 - base.I32_rotl(v26, int32(14))
	return v30 ^ v22 - base.I32_rotl(v30, int32(24))
}
func F_hashbucketcleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 float64
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v188 float64
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v212 float64
	_ = v212
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int64
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v841 int64
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v861 int64
	_ = v861
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v992 int32
	_ = v992
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1066 int32
	_ = v1066
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1114 int32
	_ = v1114
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1580 int64
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1623 int64
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1638 int32
	_ = v1638
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1652 int32
	_ = v1652
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1673 int32
	_ = v1673
	var v1680 int32
	_ = v1680
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1714 int32
	_ = v1714
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1794 int32
	_ = v1794
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1913 int32
	_ = v1913
	var v1924 int32
	_ = v1924
	var v1933 int32
	_ = v1933
	var v1962 int32
	_ = v1962
	v33 = m.G0
	v35 = v33 - int32(_a_F_hashbucketcleanup_0)
	m.G0 = v35
	v50 = l2
	v62 = l3
	v67 = int32(0)
	goto L1
L1:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v72 = int32(0)
	v73 = base.B2i32(v72 <= v50)
	if v73 == v72 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+16)))
	v93 = v92 + v91
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v94) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77+(v50^int32(-1))<<(uint(int32(2))%32))))
	v91 = v83
	goto L5
L7:
	;
	goto L8
L8:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v91 = v85 + v50<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L9:
	;
	if v324 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L10:
	;
	v127 = int32(1)
	v138 = int32(0)
	goto L15
L11:
	;
	v98 = v94 + int32(_a_F_hashbucketcleanup_1)
	if v98&int32(_a_F_hashbucketcleanup_2) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v324 = v102
	v331 = v67
	goto L9
L14:
	;
	goto L13
L15:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(20)+v127<<(uint(int32(2))%32))))
	v149 = v91 + v146&int32(_a_F_hashbucketcleanup_3)
	if l11 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v192 <= int32(0) {
		v324 = v196
		v331 = v67
		goto L9
	} else {
		goto L37
	}
L17:
	;
	if v127 != int32(base.Ui32(v98)>>(uint(int32(2))%32))&int32(_a_F_hashbucketcleanup_4) {
		v127 = v127 + int32(1)
		v138 = v192
		goto L15
	} else {
		goto L36
	}
L18:
	;
	if l9 == int32(0) {
		v192 = v138
		goto L17
	} else {
		goto L35
	}
L19:
	;
	v180 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(16)+v138<<(uint(v180)%32)))) = uint16(v127)
	v192 = v138 + v180
	goto L17
L20:
	;
	if l10 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v152 = m.T0[l11].(func(*base.Module, int32, int32) int32)(m, v149, l12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v152 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if l8 == int32(0) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v158 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v158, float64(1))
	goto L19
L25:
	;
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v149)+6)))
	if int32(0) <= v166 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v173 = v171 & l6
	if base.Ui32(v173) <= base.Ui32(l5) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v169 = int32(8)
	goto L29
L28:
	;
	v169 = int32(16)
	goto L29
L29:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v149+v169)))
	goto L26
L30:
	;
	if v175&v173 == l1 {
		goto L18
	} else {
		goto L34
	}
L31:
	;
	v175 = int32(-1)
	goto L33
L32:
	;
	v175 = l7
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L19
L35:
	;
	v188 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v188, float64(1))
	v192 = v138
	goto L17
L36:
	;
	goto L16
L37:
	;
	v199 = int32(0)
	v200 = int32(_a_F_hashbucketcleanup_5)
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v202 + int32(1)
	F_PageIndexMultiDelete(m, v91, v35+int32(16), v192)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if l8 == int32(0) {
		v227 = v199
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_MarkBufferDirty(m, v50)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L43
	}
L40:
	;
	v212 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	if base.F64_gt(v212, float64(0)) == int32(0) {
		v227 = v199
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+12)))
	if v217&int32(128) == int32(0) {
		v227 = v199
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v223 = v217 & int32(_a_F_hashbucketcleanup_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+12)) = uint16(v223)
	v227 = int32(1)
	goto L39
L43:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+118)))
	if v231 != int32(112) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v294 = int32(_a_F_hashbucketcleanup_5)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	v297 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v296 - v297
	v324 = v196
	v331 = v297
	goto L9
L45:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v235 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v238 != 0 {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+14)) = uint8(v227)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+15)) = uint8(base.B2i32(l2 == v50))
	F_XLogBeginInsert(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v239 != 0 {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	F_XLogRegisterData(m, v35+int32(14), int32(2))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+15)))
	if v250 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(42))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_XLogRegisterBuffer(m, int32(1), v50, int32(8))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v261 = int32(1)
	F_XLogRegisterBufData(m, v261, v35+int32(16), v192<<(uint(v261)%32))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v270 = F_XLogInsert(m, int32(12), int32(144))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	if v73 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = base.I64_rotr(v270, int64(32))
	goto L44
L61:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275+(v50^int32(-1))<<(uint(int32(2))%32))))
	v289 = v281
	goto L60
L62:
	;
	goto L63
L63:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v289 = v283 + v50<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v50 = v336
	v62 = v324
	v67 = v331
	goto L1
L65:
	;
	v336 = F__hash_getbuf_with_strategy(m, l0, v324, int32(1), l4)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L3
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if l2 != v50 {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	if l3 == v62 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_LockBuffer(m, v50, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L3
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_UnlockReleaseBuffer(m, v50)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L73
	}
L72:
	;
	goto L64
L73:
	;
	goto L64
L74:
	;
	F_UnlockReleaseBuffer(m, v50)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if l10 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	F_LockBuffer(m, l2, int32(2))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if l2 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L81
L81:
	;
	if v331 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L82:
	;
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v367)+16)))
	v369 = int32(_a_F_hashbucketcleanup_5)
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v371 + int32(1)
	v375 = v368 + v367
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+12)))
	v378 = v376 & int32(_a_F_hashbucketcleanup_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v375)+12)) = uint16(v378)
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L3
	} else {
		goto L86
	}
L83:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v353+(l2^int32(-1))<<(uint(int32(2))%32))))
	v367 = v359
	goto L82
L84:
	;
	goto L85
L85:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v367 = v361 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L82
L86:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+118)))
	if v383 != int32(112) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v405 = int32(_a_F_hashbucketcleanup_5)
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v407 - int32(1)
	goto L81
L88:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v387 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v390 != 0 {
		goto L87
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L94
	}
L92:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v391 != 0 {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	v400 = F_XLogInsert(m, int32(12), int32(160))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = base.I64_rotr(v400, int64(32))
	goto L87
L97:
	;
	m.G0 = v1962 + int32(_a_F_hashbucketcleanup_0)
	return
L98:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L3
	} else {
		goto L415
	}
L99:
	;
	v415 = F_IsBufferCleanupOK(m, l2)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	if v415 == int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v419 = int32(0)
	v420 = m.G0
	v422 = v420 + int32(-8192)
	m.G0 = v422
	if l2 < v419 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	m.G0 = v1913 - int32(-8192)
	v1962 = v1924
	goto L97
L103:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441)+16)))
	v443 = v442 + v441
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v444 != int32(-1) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v427+(l2^int32(-1))<<(uint(int32(2))%32))))
	v441 = v433
	goto L103
L105:
	;
	goto L106
L106:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v441 = v435 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v453 = v444
	v459 = v419
	goto L110
L108:
	;
	goto L109
L109:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L3
	} else {
		goto L414
	}
L110:
	;
	if v459 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v507 = l2
	v508 = l3
	v511 = l4
	v512 = l3
	v514 = l2
	v515 = v453
	v516 = l0
	v519 = v482
	v523 = v501
	v524 = v422
	v530 = v441
	v531 = v503
	v534 = v443
	v535 = v35
	goto L122
L112:
	;
	F_UnlockReleaseBuffer(m, v459)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L3
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v482 = F__hash_getbuf_with_strategy(m, l0, v453, int32(1), l4)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L3
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v501)+16)))
	v503 = v502 + v501
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+4))
	if v504 != int32(-1) {
		v453 = v504
		v459 = v482
		goto L110
	} else {
		goto L121
	}
L117:
	;
	if v482 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v487+(v482^int32(-1))<<(uint(int32(2))%32))))
	v501 = v493
	goto L116
L119:
	;
	goto L120
L120:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v501 = v495 + v482<<(uint(int32(13))%32) + int32(-8192)
	goto L116
L121:
	;
	goto L111
L122:
	;
	v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v523)+12)))
	if base.Ui32(v539) < base.Ui32(int32(25)) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	F_UnlockReleaseBuffer(m, v1101)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L3
	} else {
		goto L413
	}
L124:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v1135 = v524 + int32(2464)
	v1137 = v524 + int32(16)
	v1141 = v1114 & int32(_a_F_hashbucketcleanup_4)
	v1143 = m.G0
	v1145 = v1143 - int32(32)
	m.G0 = v1145
	F__hash_checkpage(m, v516, v519, int32(1))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L3
	} else {
		goto L221
	}
L125:
	;
	v1101 = v507
	v1102 = v508
	v1114 = int32(0)
	v1124 = v530
	v1128 = v534
	goto L124
L126:
	;
	goto L127
L127:
	;
	v546 = int32(base.Ui32(v539+int32(_a_F_hashbucketcleanup_1)) >> (uint(int32(2)) % 32))
	if v546&int32(_a_F_hashbucketcleanup_4) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1101 = v507
	v1102 = v508
	v1114 = int32(0)
	v1124 = v530
	v1128 = v534
	goto L124
L129:
	;
	goto L130
L130:
	;
	v562 = v507
	v563 = v508
	v566 = v546
	v583 = v530
	v587 = v534
	goto L131
L131:
	;
	v593 = int32(0)
	v598 = v562
	v599 = v563
	v604 = v593
	v611 = v593
	v618 = int32(1)
	v621 = v583
	v623 = v593
	v625 = v587
	goto L133
L132:
	;
	v1101 = v703
	v1102 = v698
	v1114 = v981
	v1124 = v978
	v1128 = v980
	goto L124
L133:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v523+int32(20)+v618&int32(_a_F_hashbucketcleanup_4)<<(uint(int32(2))%32))))
	v636 = int32(_a_F_hashbucketcleanup_8)
	if v635&v636 != v636 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v523)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1090) {
		goto L217
	} else {
		goto L218
	}
L135:
	;
	goto L134
L136:
	;
	v642 = v523 + v635&int32(_a_F_hashbucketcleanup_3)
	v643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v642)+6)))
	v649 = (v643&int32(_a_F_hashbucketcleanup_9) + int32(7)) & int32(_a_F_hashbucketcleanup_10)
	v650 = v598
	v651 = v599
	v656 = v604
	v663 = v611
	v673 = v621
	v675 = v623
	v677 = v625
	goto L140
L137:
	;
	v1053 = v598
	v1054 = v599
	v1059 = v604
	v1066 = v611
	v1076 = v621
	v1078 = v623
	v1080 = v625
	goto L138
L138:
	;
	v1086 = v618 + int32(1)
	if base.Ui32(v1086&int32(_a_F_hashbucketcleanup_4)) <= base.Ui32(v566&int32(_a_F_hashbucketcleanup_4)) {
		v598 = v1053
		v599 = v1054
		v604 = v1059
		v611 = v1066
		v618 = v1086
		v621 = v1076
		v623 = v1078
		v625 = v1080
		goto L133
	} else {
		goto L216
	}
L139:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v524+int32(_a_F_hashbucketcleanup_11)+v656&int32(_a_F_hashbucketcleanup_4)<<(uint(int32(1))%32)))) = uint16(v618)
	v1037 = F_CopyIndexTuple(m, v642)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L3
	} else {
		goto L215
	}
L140:
	;
	v683 = v663 & int32(_a_F_hashbucketcleanup_4)
	v686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673)+14)))
	v687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673)+12)))
	v688 = v686 - v687
	v690 = (v683 + int32(1)) << (uint(int32(2)) % 32)
	if v690 <= v688 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v992 = v981
	goto L211
L142:
	;
	v695 = v649 + v675
	if base.Ui32(v695) <= base.Ui32(v694) {
		goto L139
	} else {
		goto L146
	}
L143:
	;
	v694 = v688 - v690
	goto L145
L144:
	;
	v694 = int32(0)
	goto L145
L145:
	;
	goto L142
L146:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v677)+4))
	if v515 != v698 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v701 = F__hash_getbuf_with_strategy(m, v516, v698, int32(1), v511)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L3
	} else {
		goto L150
	}
L148:
	;
	v703 = int32(0)
	goto L149
L149:
	;
	if v683 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v703 = v701
	goto L149
L151:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v516)+48))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+118)))
	if v705 != int32(112) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	if v651 == v512 {
		goto L197
	} else {
		goto L198
	}
L154:
	;
	v719 = int32(_a_F_hashbucketcleanup_5)
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v721 + int32(1)
	F__hash_pgaddmultitup(m, v516, v650, v524+int32(2464), v524+int32(16), v683)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L3
	} else {
		goto L162
	}
L155:
	;
	v709 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v709 <= int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v516)+32))
	if v712 != 0 {
		goto L154
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	F_XLogEnsureRecordSpace(m, int32(0), v683+int32(3))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L3
	} else {
		goto L161
	}
L159:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v516)+40))
	if v713 != 0 {
		goto L154
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	goto L154
L162:
	;
	F_MarkBufferDirty(m, v650)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	v736 = v656 & int32(_a_F_hashbucketcleanup_4)
	F_PageIndexMultiDelete(m, v523, v524+int32(_a_F_hashbucketcleanup_11), v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	F_MarkBufferDirty(m, v519)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v516)+48))
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741)+118)))
	if v742 != int32(112) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v914 = int32(_a_F_hashbucketcleanup_5)
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v916 - int32(1)
	goto L153
L167:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v746 <= int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v516)+32))
	if v749 != 0 {
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v524)+12)) = uint16(v663)
	*(*uint8)(unsafe.Add(mBase, uint32(v524)+14)) = uint8(base.B2i32(v650 == v514))
	F_XLogBeginInsert(m)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L3
	} else {
		goto L173
	}
L171:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v516)+40))
	if v750 != 0 {
		goto L166
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	F_XLogRegisterData(m, v524+int32(12), int32(3))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L3
	} else {
		goto L174
	}
L174:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+14)))
	if v761 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	F_XLogRegisterBuffer(m, int32(0), v514, int32(42))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L3
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	F_XLogRegisterBuffer(m, int32(1), v650, int32(8))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L3
	} else {
		goto L179
	}
L178:
	;
	goto L177
L179:
	;
	v772 = int32(1)
	F_XLogRegisterBufData(m, v772, v524+int32(16), v683<<(uint(v772)%32))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L3
	} else {
		goto L180
	}
L180:
	;
	v786 = int32(0)
	goto L181
L181:
	;
	v814 = v786 << (uint(int32(2)) % 32)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v814+(v524+int32(2464)))))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v524+int32(832)+v814)))
	F_XLogRegisterBufData(m, int32(1), v818, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L3
	} else {
		goto L183
	}
L182:
	;
	F_XLogRegisterBuffer(m, int32(2), v519, int32(8))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L3
	} else {
		goto L185
	}
L183:
	;
	v826 = v786 + int32(1)
	if v826 != v683 {
		v786 = v826
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	F_XLogRegisterBufData(m, int32(2), v524+int32(_a_F_hashbucketcleanup_11), v736<<(uint(int32(1))%32))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L3
	} else {
		goto L186
	}
L186:
	;
	v841 = F_XLogInsert(m, int32(12), int32(112))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L3
	} else {
		goto L187
	}
L187:
	;
	if v650 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v861 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v860))) = base.I64_rotr(v841, v861)
	if v519 < int32(0) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v846+(v650^int32(-1))<<(uint(int32(2))%32))))
	v860 = v852
	goto L188
L190:
	;
	goto L191
L191:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v860 = v854 + v650<<(uint(int32(13))%32) + int32(-8192)
	goto L188
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v879)+4)) = base.I32_wrap_i64(v841)
	*(*int32)(unsafe.Add(mBase, uint32(v879))) = base.I32_wrap_i64(int64(base.Ui64(v841) >> (uint(v861) % 64)))
	goto L166
L193:
	;
	v871 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v871+(v519^int32(-1))<<(uint(int32(2))%32))))
	v879 = v873
	goto L192
L194:
	;
	goto L195
L195:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v879 = v875 + v519<<(uint(int32(13))%32) + int32(-8192)
	goto L192
L196:
	;
	if v698 == v515 {
		goto L202
	} else {
		goto L203
	}
L197:
	;
	F_LockBuffer(m, v650, int32(0))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L3
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	F_UnlockReleaseBuffer(m, v650)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L3
	} else {
		goto L201
	}
L200:
	;
	goto L196
L201:
	;
	goto L196
L202:
	;
	F_UnlockReleaseBuffer(m, v519)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L3
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if v703 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v1913 = v524
	v1924 = v535
	goto L102
L206:
	;
	v979 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v978)+16)))
	v980 = v979 + v978
	v981 = int32(0)
	if v683 == v981 {
		v650 = v703
		v651 = v698
		v656 = v981
		v663 = v981
		v673 = v978
		v675 = v981
		v677 = v980
		goto L140
	} else {
		goto L210
	}
L207:
	;
	v964 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v964+(v703^int32(-1))<<(uint(int32(2))%32))))
	v978 = v970
	goto L206
L208:
	;
	goto L209
L209:
	;
	v972 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v978 = v972 + v703<<(uint(int32(13))%32) + int32(-8192)
	goto L206
L210:
	;
	goto L141
L211:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v524+int32(2464)+v992<<(uint(int32(2))%32))))
	F_pfree(m, v1023)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L3
	} else {
		goto L213
	}
L212:
	;
	goto L135
L213:
	;
	v1027 = v992 + int32(1)
	if v1027 != v683 {
		v992 = v1027
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v1040 = v683 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1040+(v524+int32(832))))) = v649
	*(*int32)(unsafe.Add(mBase, uint32(v524+int32(2464)+v1040))) = v1037
	v1049 = int32(1)
	v1053 = v650
	v1054 = v651
	v1059 = v656 + v1049
	v1066 = v663 + v1049
	v1076 = v673
	v1078 = v695
	v1080 = v677
	goto L138
L216:
	;
	v1101 = v1053
	v1102 = v1054
	v1114 = v1066
	v1124 = v1076
	v1128 = v1080
	goto L124
L217:
	;
	v1098 = int32(base.Ui32(v1090+int32(_a_F_hashbucketcleanup_1)) >> (uint(int32(2)) % 32))
	goto L219
L218:
	;
	v1098 = int32(0)
	goto L219
L219:
	;
	if v1098&int32(_a_F_hashbucketcleanup_4) != 0 {
		v562 = v703
		v563 = v698
		v566 = v1098
		v583 = v978
		v587 = v980
		goto L131
	} else {
		goto L220
	}
L220:
	;
	goto L132
L221:
	;
	if v519 < int32(0) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	if v519 < int32(0) {
		goto L227
	} else {
		goto L228
	}
L223:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[4]))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1153+(v519^int32(-1))<<(uint(int32(6))%32))+16))
	v1168 = v1159
	goto L222
L224:
	;
	goto L225
L225:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[5]))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1161+v519<<(uint(int32(6))%32)+int32(-64))+16))
	v1168 = v1167
	goto L222
L226:
	;
	v1187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1186)+16)))
	v1188 = v1187 + v1186
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1188)))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+4))
	if v1101 < int32(0) {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1172+(v519^int32(-1))<<(uint(int32(2))%32))))
	v1186 = v1178
	goto L226
L228:
	;
	goto L229
L229:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1186 = v1180 + v519<<(uint(int32(13))%32) + int32(-8192)
	goto L226
L230:
	;
	if v1189 == int32(-1) {
		v1217 = int32(0)
		goto L234
	} else {
		goto L235
	}
L231:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[4]))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1194+(v1101^int32(-1))<<(uint(int32(6))%32))+16))
	v1209 = v1200
	goto L230
L232:
	;
	goto L233
L233:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[5]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1202+v1101<<(uint(int32(6))%32)+int32(-64))+16))
	v1209 = v1208
	goto L230
L234:
	;
	if v1190 != int32(-1) {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	if v1209 == v1189 {
		v1217 = v1101
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1215 = F__hash_getbuf_with_strategy(m, v516, v1189, int32(3), v511)
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L3
	} else {
		goto L237
	}
L237:
	;
	v1217 = v1215
	goto L234
L238:
	;
	v1221 = F__hash_getbuf_with_strategy(m, v516, v1190, int32(1), v511)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L3
	} else {
		goto L241
	}
L239:
	;
	v1223 = int32(0)
	goto L240
L240:
	;
	v1227 = F__hash_getbuf(m, v516, int32(0), int32(1), int32(8))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L3
	} else {
		goto L243
	}
L241:
	;
	v1223 = v1221
	goto L240
L242:
	;
	v1249 = F__hash_ovflblkno_to_bitno(m, v1246+int32(24), v1168)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L3
	} else {
		goto L247
	}
L243:
	;
	if v1227 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1232+(v1227^int32(-1))<<(uint(int32(2))%32))))
	v1246 = v1238
	goto L242
L245:
	;
	goto L246
L246:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1246 = v1240 + v1227<<(uint(int32(13))%32) + int32(-8192)
	goto L242
L247:
	;
	v1251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1246)+46)))
	v1252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1246)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v1145)+28)) = v1249 & (v1252<<(uint(int32(3))%32) - int32(1))
	v1259 = int32(base.Ui32(v1249) >> (uint(v1251) % 32))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+68))
	if base.Ui32(v1259) < base.Ui32(v1260) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	if v1141 != 0 {
		goto L395
	} else {
		goto L396
	}
L249:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1246+v1259<<(uint(int32(2))%32))+468))
	F_LockBuffer(m, v1227, int32(0))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L3
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L3
	} else {
		goto L392
	}
L252:
	;
	v1271 = F__hash_getbuf(m, v516, v1265, int32(2), int32(4))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L3
	} else {
		goto L254
	}
L253:
	;
	F_LockBuffer(m, v1227, int32(2))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L3
	} else {
		goto L258
	}
L254:
	;
	if v1271 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1276+(v1271^int32(-1))<<(uint(int32(2))%32))))
	v1290 = v1282
	goto L253
L256:
	;
	goto L257
L257:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1290 = v1284 + v1271<<(uint(int32(13))%32) + int32(-8192)
	goto L253
L258:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v516)+48))
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+118)))
	if v1295 != int32(112) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1309 = int32(_a_F_hashbucketcleanup_5)
	v1311 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v1311 + int32(1)
	if v1141 != 0 {
		goto L267
	} else {
		goto L268
	}
L260:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v1299 <= int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v516)+32))
	if v1302 != 0 {
		goto L259
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	F_XLogEnsureRecordSpace(m, int32(6), v1141+int32(4))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L3
	} else {
		goto L266
	}
L264:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v516)+40))
	if v1303 != 0 {
		goto L259
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	goto L259
L267:
	;
	F__hash_pgaddmultitup(m, v516, v1101, v1135, v1137, v1141)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L3
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_PageInit(m, v1186, int32(_a_F_hashbucketcleanup_12), int32(16))
	mBase = m.M
	goto L272
L270:
	;
	F_MarkBufferDirty(m, v1101)
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L3
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	v1322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1186)+16)))
	v1323 = v1186 + v1322
	*(*int64)(unsafe.Add(mBase, uint32(v1323)+8)) = int64(-36028792723996673)
	*(*int64)(unsafe.Add(mBase, uint32(v1323))) = int64(-1)
	F_MarkBufferDirty(m, v519)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L3
	} else {
		goto L273
	}
L273:
	;
	if v1217 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	if v1217 < int32(0) {
		goto L278
	} else {
		goto L279
	}
L275:
	;
	goto L276
L276:
	;
	if v1223 != 0 {
		goto L282
	} else {
		goto L283
	}
L277:
	;
	v1348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1347)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1348+v1347)+4)) = v1190
	F_MarkBufferDirty(m, v1217)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L3
	} else {
		goto L281
	}
L278:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1333+(v1217^int32(-1))<<(uint(int32(2))%32))))
	v1347 = v1339
	goto L277
L279:
	;
	goto L280
L280:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1347 = v1341 + v1217<<(uint(int32(13))%32) + int32(-8192)
	goto L277
L281:
	;
	goto L276
L282:
	;
	if v1223 < int32(0) {
		goto L286
	} else {
		goto L287
	}
L283:
	;
	goto L284
L284:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+28))
	v1382 = base.I32_div_s(v1380, int32(32))
	v1385 = v1290 + int32(24) + v1382<<(uint(int32(2))%32)
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1385)))
	*(*int32)(unsafe.Add(mBase, uint32(v1385))) = v1386 & base.I32_rotl(int32(-2), v1380)
	F_MarkBufferDirty(m, v1271)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L3
	} else {
		goto L290
	}
L285:
	;
	v1372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1371)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1372+v1371))) = v1189
	F_MarkBufferDirty(m, v1223)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L3
	} else {
		goto L289
	}
L286:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1357+(v1223^int32(-1))<<(uint(int32(2))%32))))
	v1371 = v1363
	goto L285
L287:
	;
	goto L288
L288:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1371 = v1365 + v1223<<(uint(int32(13))%32) + int32(-8192)
	goto L285
L289:
	;
	goto L284
L290:
	;
	v1394 = v1246 - int32(-64)
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+64))
	v1396 = base.B2i32(base.Ui32(v1395) <= base.Ui32(v1249))
	if v1396 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1394))) = v1249
	F_MarkBufferDirty(m, v1227)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L3
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v516)+48))
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402)+118)))
	if v1403 != int32(112) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	goto L293
L295:
	;
	v1749 = int32(_a_F_hashbucketcleanup_5)
	v1751 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v1751 - int32(1)
	v1755 = int32(0)
	if base.B2i32(v1217 == v1755)|base.B2i32(v1209 == v1189) == v1755 {
		goto L378
	} else {
		goto L379
	}
L296:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v1407 <= int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v516)+32))
	if v1410 != 0 {
		goto L295
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1145)+24)) = uint16(v1141)
	*(*int32)(unsafe.Add(mBase, uint32(v1145)+20)) = v1190
	*(*int32)(unsafe.Add(mBase, uint32(v1145)+16)) = v1189
	*(*uint8)(unsafe.Add(mBase, uint32(v1145)+27)) = uint8(base.B2i32(v1101 == v1217))
	*(*uint8)(unsafe.Add(mBase, uint32(v1145)+26)) = uint8(base.B2i32(v1101 == v514))
	F_XLogBeginInsert(m)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L3
	} else {
		goto L302
	}
L300:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v516)+40))
	if v1411 != 0 {
		goto L295
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	F_XLogRegisterData(m, v1145+int32(16), int32(12))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L3
	} else {
		goto L303
	}
L303:
	;
	v1426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+26)))
	if v1426 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	F_XLogRegisterBuffer(m, int32(0), v514, int32(42))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L3
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145)+24)))
	if v1433 != 0 {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L306
L308:
	;
	F_XLogRegisterBuffer(m, int32(2), v519, int32(8))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L3
	} else {
		goto L327
	}
L309:
	;
	v1434 = int32(1)
	F_XLogRegisterBuffer(m, v1434, v1101, int32(8))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L3
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+27)))
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+26)))
	if v1492 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L312:
	;
	v1439 = int32(1)
	F_XLogRegisterBufData(m, v1439, v1137, v1141<<(uint(v1439)%32))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L3
	} else {
		goto L313
	}
L313:
	;
	if v1141 == int32(0) {
		v1511 = v1434
		goto L308
	} else {
		goto L314
	}
L314:
	;
	v1453 = int32(0)
	goto L315
L315:
	;
	v1481 = v1453 << (uint(int32(2)) % 32)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1135+v1481)))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1481+(v524+int32(832)))))
	F_XLogRegisterBufData(m, int32(1), v1483, v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L3
	} else {
		goto L317
	}
L316:
	;
	v1511 = v1434
	goto L308
L317:
	;
	v1489 = v1453 + int32(1)
	if v1489 != v1141 {
		v1453 = v1489
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	v1495 = int32(0)
	if v1491&int32(1) == v1495 {
		v1511 = v1495
		goto L308
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1501 = int32(1)
	if v1491&v1501 != 0 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	goto L321
L323:
	;
	v1506 = int32(8)
	goto L325
L324:
	;
	v1506 = int32(40)
	goto L325
L325:
	;
	F_XLogRegisterBuffer(m, v1501, v1101, v1506)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L3
	} else {
		goto L326
	}
L326:
	;
	v1511 = v1491
	goto L308
L327:
	;
	if v1217 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	if v1223 != 0 {
		goto L332
	} else {
		goto L333
	}
L329:
	;
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+27)))
	if v1547&int32(1) != 0 {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	F_XLogRegisterBuffer(m, int32(3), v1217, int32(8))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L3
	} else {
		goto L331
	}
L331:
	;
	goto L328
L332:
	;
	F_XLogRegisterBuffer(m, int32(4), v1223, int32(8))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L3
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	F_XLogRegisterBuffer(m, int32(5), v1271, int32(8))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L3
	} else {
		goto L336
	}
L335:
	;
	goto L334
L336:
	;
	F_XLogRegisterBufData(m, int32(5), v1145+int32(28), int32(4))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L3
	} else {
		goto L337
	}
L337:
	;
	if v1396 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	F_XLogRegisterBuffer(m, int32(6), v1227, int32(8))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L3
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	v1580 = F_XLogInsert(m, int32(12), int32(128))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L3
	} else {
		goto L343
	}
L341:
	;
	F_XLogRegisterBufData(m, int32(6), v1394, int32(4))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L3
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	if v1511&int32(1) != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	if v1101 < int32(0) {
		goto L348
	} else {
		goto L349
	}
L345:
	;
	goto L346
L346:
	;
	if v519 < int32(0) {
		goto L352
	} else {
		goto L353
	}
L347:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1601))) = base.I64_rotr(v1580, int64(32))
	goto L346
L348:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1587+(v1101^int32(-1))<<(uint(int32(2))%32))))
	v1601 = v1593
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1601 = v1595 + v1101<<(uint(int32(13))%32) + int32(-8192)
	goto L347
L351:
	;
	v1623 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1622))) = base.I64_rotr(v1580, v1623)
	v1628 = base.I32_wrap_i64(int64(base.Ui64(v1580) >> (uint(v1623) % 64)))
	v1629 = base.I32_wrap_i64(v1580)
	if v1217 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L352:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1608+(v519^int32(-1))<<(uint(int32(2))%32))))
	v1622 = v1614
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1622 = v1616 + v519<<(uint(int32(13))%32) + int32(-8192)
	goto L351
L355:
	;
	if v1223 != 0 {
		goto L362
	} else {
		goto L363
	}
L356:
	;
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+27)))
	if v1632&int32(1) != 0 {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	if v1217 < int32(0) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1652)+4)) = v1629
	*(*int32)(unsafe.Add(mBase, uint32(v1652))) = v1628
	goto L355
L359:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1638+(v1217^int32(-1))<<(uint(int32(2))%32))))
	v1652 = v1644
	goto L358
L360:
	;
	goto L361
L361:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1652 = v1646 + v1217<<(uint(int32(13))%32) + int32(-8192)
	goto L358
L362:
	;
	if v1223 < int32(0) {
		goto L366
	} else {
		goto L367
	}
L363:
	;
	goto L364
L364:
	;
	if v1271 < int32(0) {
		goto L370
	} else {
		goto L371
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1673)+4)) = v1629
	*(*int32)(unsafe.Add(mBase, uint32(v1673))) = v1628
	goto L364
L366:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1659+(v1223^int32(-1))<<(uint(int32(2))%32))))
	v1673 = v1665
	goto L365
L367:
	;
	goto L368
L368:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1673 = v1667 + v1223<<(uint(int32(13))%32) + int32(-8192)
	goto L365
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+4)) = v1629
	*(*int32)(unsafe.Add(mBase, uint32(v1694))) = v1628
	if base.Ui32(v1395) <= base.Ui32(v1249) {
		goto L295
	} else {
		goto L373
	}
L370:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1680+(v1271^int32(-1))<<(uint(int32(2))%32))))
	v1694 = v1686
	goto L369
L371:
	;
	goto L372
L372:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1694 = v1688 + v1271<<(uint(int32(13))%32) + int32(-8192)
	goto L369
L373:
	;
	if v1227 < int32(0) {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1714)+4)) = v1629
	*(*int32)(unsafe.Add(mBase, uint32(v1714))) = v1628
	goto L295
L375:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1700+(v1227^int32(-1))<<(uint(int32(2))%32))))
	v1714 = v1706
	goto L374
L376:
	;
	goto L377
L377:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1714 = v1708 + v1227<<(uint(int32(13))%32) + int32(-8192)
	goto L374
L378:
	;
	F_UnlockReleaseBuffer(m, v1217)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L3
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	if v519 != 0 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	goto L380
L382:
	;
	F_UnlockReleaseBuffer(m, v519)
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L3
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	if v1223 != 0 {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	goto L384
L386:
	;
	F_UnlockReleaseBuffer(m, v1223)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L3
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	F_UnlockReleaseBuffer(m, v1271)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L3
	} else {
		goto L390
	}
L389:
	;
	goto L388
L390:
	;
	F_UnlockReleaseBuffer(m, v1227)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L3
	} else {
		goto L391
	}
L391:
	;
	m.G0 = v1145 + int32(32)
	goto L248
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1145))) = v1249
	F_errmsg_internal(m, int32(_a_F_hashbucketcleanup_13), v1145)
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L3
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(_a_F_hashbucketcleanup_14), int32(562), int32(_a_F_hashbucketcleanup_15))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L3
	} else {
		goto L394
	}
L394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L395:
	;
	v1794 = int32(0)
	goto L398
L396:
	;
	goto L397
L397:
	;
	if v1102 == v1133 {
		goto L403
	} else {
		goto L404
	}
L398:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v524+int32(2464)+v1794<<(uint(int32(2))%32))))
	F_pfree(m, v1825)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L3
	} else {
		goto L400
	}
L399:
	;
	goto L397
L400:
	;
	v1829 = v1794 + int32(1)
	if v1829 != v1141 {
		v1794 = v1829
		goto L398
	} else {
		goto L401
	}
L401:
	;
	goto L399
L402:
	;
	goto L123
L403:
	;
	if v1102 != v512 {
		goto L402
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1869 = F__hash_getbuf_with_strategy(m, v516, v1133, int32(1), v511)
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L3
	} else {
		goto L409
	}
L406:
	;
	F_LockBuffer(m, v1101, int32(0))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L3
	} else {
		goto L407
	}
L407:
	;
	v1913 = v524
	v1924 = v535
	goto L102
L408:
	;
	v1889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1888)+16)))
	v507 = v1101
	v508 = v1102
	v515 = v1133
	v519 = v1869
	v523 = v1888
	v530 = v1124
	v531 = v1889 + v1888
	v534 = v1128
	goto L122
L409:
	;
	if v1869 < int32(0) {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1874+(v1869^int32(-1))<<(uint(int32(2))%32))))
	v1888 = v1880
	goto L408
L411:
	;
	goto L412
L412:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1888 = v1882 + v1869<<(uint(int32(13))%32) + int32(-8192)
	goto L408
L413:
	;
	v1913 = v524
	v1924 = v535
	goto L102
L414:
	;
	v1913 = v422
	v1924 = v35
	goto L102
L415:
	;
	v1962 = v35
	goto L97
}
func F_hashbyteaextended(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hashvarlenaextended(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_hashfloat8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	if base.F64_ne(v9, float64(0)) != 0 {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v9)&int64(9223372036854775807)) {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(9221120237041090560)
		} else {
		}
		v22 = v6 + int32(8)
		v29 = int32(-1636608424)
		if v22&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v249 = v29
				v250 = v29
				v251 = v29
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 1:
				v242 = v29
				v243 = v29
				v244 = v29
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 2:
				v235 = v29
				v236 = v29
				v237 = v29
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 3:
				v229 = v29
				v230 = v29
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 4:
				v225 = v29
				v226 = v29
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 5:
				v219 = v29
				v220 = v29
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 6:
				v213 = v29
				v214 = v29
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 7:
				v208 = v29
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 8:
				v203 = v29
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v208 = v204<<(uint(int32(8))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 9:
				v198 = v29
				v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v203 = v199<<(uint(int32(16))%32) + v198
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v208 = v204<<(uint(int32(8))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 10:
				v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
				v198 = v194<<(uint(int32(24))%32) + v29
				v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v203 = v199<<(uint(int32(16))%32) + v198
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v208 = v204<<(uint(int32(8))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			default:
				v256 = v29
				v257 = v29
				v258 = v29
			}
		} else {
			switch int32(7) {
			case 0:
				v135 = v29
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v135 + v136
				v257 = v29
				v258 = v29
			case 1:
				v130 = v29
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v135 = v131<<(uint(int32(8))%32) + v130
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v135 + v136
				v257 = v29
				v258 = v29
			case 2:
				v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v130 = v126<<(uint(int32(16))%32) + v29
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v135 = v131<<(uint(int32(8))%32) + v130
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v135 + v136
				v257 = v29
				v258 = v29
			case 3:
				v123 = v29
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 4:
				v120 = v29
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v123 = v120 + v121
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 5:
				v115 = v29
				v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v120 = v116<<(uint(int32(8))%32) + v115
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v123 = v120 + v121
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 6:
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v115 = v111<<(uint(int32(16))%32) + v29
				v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v120 = v116<<(uint(int32(8))%32) + v115
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v123 = v120 + v121
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 7:
				v106 = v29
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			case 8:
				v101 = v29
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v106 = v102<<(uint(int32(8))%32) + v101
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			case 9:
				v96 = v29
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v101 = v97<<(uint(int32(16))%32) + v96
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v106 = v102<<(uint(int32(8))%32) + v101
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			case 10:
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
				v96 = v92<<(uint(int32(24))%32) + v29
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v101 = v97<<(uint(int32(16))%32) + v96
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v106 = v102<<(uint(int32(8))%32) + v101
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			default:
				v256 = v29
				v257 = v29
				v258 = v29
			}
		}
		v261 = int32(14)
		v263 = v257 ^ v258 - base.I32_rotl(v257, v261)
		v267 = v263 ^ v256 - base.I32_rotl(v263, int32(11))
		v271 = v267 ^ v257 - base.I32_rotl(v267, int32(25))
		v275 = v271 ^ v263 - base.I32_rotl(v271, int32(16))
		v279 = v275 ^ v267 - base.I32_rotl(v275, int32(4))
		v283 = v279 ^ v271 - base.I32_rotl(v279, v261)
		v288 = v283 ^ v275 - base.I32_rotl(v283, int32(24))
	} else {
		v288 = int32(0)
	}
	m.G0 = v6 + int32(16)
	return v288
}
func F_hashgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v6)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v9 == int32(-1) {
		v12 = F__hash_first(m, l0, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		if v17 != int32(1) {
			v41 = F__hash_next(m, l0, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				return v41
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			if v20 == int32(0) {
				v24 = F_palloc(m, int32(1632))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					v27 = v24
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					if int32(407) < v28 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v28 + int32(1)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v37
					}
					v41 = F__hash_next(m, l0, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						return v41
					}
				}
			} else {
				v27 = v20
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				if int32(407) < v28 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v28 + int32(1)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v37
				}
				v41 = F__hash_next(m, l0, l1)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					return v41
				}
			}
		}
	}
}
func F_hashhandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v37 int64
	_ = v37
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = v7
		v9 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+14)) = uint16(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+10)) = int32(16842752)
		v13 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+8)) = uint16(v13)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(844429225099702)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+19)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = v7
		v21 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+23)) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(121)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(122)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(123)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(124)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(125)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(126)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(127)
		v37 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+76)) = v37
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(128)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(129)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(130)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(131)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(132)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(133)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(134)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = int32(23)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+28)) = uint16(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+128)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v37
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v37
		*(*int32)(unsafe.Add(mBase, uint32(v3)+132)) = int32(135)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = int32(136)
		return v3
	}
}
func F_hashint2extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		v11 = int32(-1636608428)
		v50 = v11
		v51 = v11
		v54 = int32(0)
	} else {
		v14 = base.I32_wrap_i64(v4)
		v16 = v14 + int32(1021750440)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v27 = v14 - v21 - int32(1636608428) ^ base.I32_rotl(v21, int32(6))
		v31 = v16 - v27 ^ base.I32_rotl(v27, int32(8))
		v32 = v21 + v16
		v33 = v27 + v32
		v34 = v31 + v33
		v38 = v32 - v31 ^ base.I32_rotl(v31, int32(16))
		v42 = v33 - v38 ^ base.I32_rotl(v38, int32(19))
		v47 = v38 + v34
		v48 = v42 + v47
		v50 = v48
		v51 = v47
		v54 = v34 - v42 ^ base.I32_rotl(v42, int32(4)) ^ v48
	}
	v55 = int32(14)
	v57 = v54 - base.I32_rotl(v50, v55)
	v62 = v57 ^ (v2 + v51) - base.I32_rotl(v57, int32(11))
	v66 = v50 ^ v62 - base.I32_rotl(v62, int32(25))
	v70 = v66 ^ v57 - base.I32_rotl(v66, int32(16))
	v74 = v70 ^ v62 - base.I32_rotl(v70, int32(4))
	v78 = v74 ^ v66 - base.I32_rotl(v74, v55)
	v88 = F_Int64GetDatum(m, base.I64_extend_i32_u(v78)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v78^v70-base.I32_rotl(v78, int32(24))))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		return int32(0)
	} else {
		return v88
	}
}
func F_hashint8extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if v13 == int64(0) {
		v20 = int32(-1636608428)
		v59 = v20
		v60 = v20
		v63 = int32(0)
	} else {
		v23 = base.I32_wrap_i64(v13)
		v25 = v23 + int32(1021750440)
		v30 = base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v36 = v23 - v30 - int32(1636608428) ^ base.I32_rotl(v30, int32(6))
		v40 = v25 - v36 ^ base.I32_rotl(v36, int32(8))
		v41 = v30 + v25
		v42 = v36 + v41
		v43 = v40 + v42
		v47 = v41 - v40 ^ base.I32_rotl(v40, int32(16))
		v51 = v42 - v47 ^ base.I32_rotl(v47, int32(19))
		v56 = v47 + v43
		v57 = v51 + v56
		v59 = v57
		v60 = v56
		v63 = v43 - v51 ^ base.I32_rotl(v51, int32(4)) ^ v57
	}
	v64 = int32(14)
	v66 = v63 - base.I32_rotl(v59, v64)
	v71 = v66 ^ (base.I32_wrap_i64(v4>>(uint(int64(63))%64)^int64(base.Ui64(v4)>>(uint(int64(32))%64))^v4) + v60) - base.I32_rotl(v66, int32(11))
	v75 = v59 ^ v71 - base.I32_rotl(v71, int32(25))
	v79 = v75 ^ v66 - base.I32_rotl(v75, int32(16))
	v83 = v79 ^ v71 - base.I32_rotl(v79, int32(4))
	v87 = v83 ^ v75 - base.I32_rotl(v83, v64)
	v97 = F_Int64GetDatum(m, base.I64_extend_i32_u(v87)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v87^v79-base.I32_rotl(v87, int32(24))))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		return int32(0)
	} else {
		return v97
	}
}
func F_hashmacaddr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(-1636608426)
	if v2&int32(3) != 0 {
		switch int32(5) {
		case 0:
			v229 = v9
			v230 = v9
			v231 = v9
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 1:
			v222 = v9
			v223 = v9
			v224 = v9
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 2:
			v215 = v9
			v216 = v9
			v217 = v9
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 3:
			v209 = v9
			v210 = v9
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 4:
			v205 = v9
			v206 = v9
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 5:
			v199 = v9
			v200 = v9
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 6:
			v193 = v9
			v194 = v9
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 7:
			v188 = v9
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 8:
			v183 = v9
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 9:
			v178 = v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v178 = v174<<(uint(int32(24))%32) + v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	} else {
		switch int32(5) {
		case 0:
			v115 = v9
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 1:
			v110 = v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 2:
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v110 = v106<<(uint(int32(16))%32) + v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 3:
			v103 = v9
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 4:
			v100 = v9
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 5:
			v95 = v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 6:
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v95 = v91<<(uint(int32(16))%32) + v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 7:
			v86 = v9
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 8:
			v81 = v9
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 9:
			v76 = v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 10:
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v76 = v72<<(uint(int32(24))%32) + v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	}
	v241 = int32(14)
	v243 = v237 ^ v238 - base.I32_rotl(v237, v241)
	v247 = v243 ^ v236 - base.I32_rotl(v243, int32(11))
	v251 = v247 ^ v237 - base.I32_rotl(v247, int32(25))
	v255 = v251 ^ v243 - base.I32_rotl(v251, int32(16))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(4))
	v263 = v259 ^ v251 - base.I32_rotl(v259, v241)
	return v263 ^ v255 - base.I32_rotl(v263, int32(24))
}
func F_hashmacaddr8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(-1636608424)
	if v2&int32(3) != 0 {
		switch int32(7) {
		case 0:
			v229 = v9
			v230 = v9
			v231 = v9
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 1:
			v222 = v9
			v223 = v9
			v224 = v9
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 2:
			v215 = v9
			v216 = v9
			v217 = v9
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 3:
			v209 = v9
			v210 = v9
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 4:
			v205 = v9
			v206 = v9
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 5:
			v199 = v9
			v200 = v9
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 6:
			v193 = v9
			v194 = v9
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 7:
			v188 = v9
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 8:
			v183 = v9
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 9:
			v178 = v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v178 = v174<<(uint(int32(24))%32) + v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	} else {
		switch int32(7) {
		case 0:
			v115 = v9
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 1:
			v110 = v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 2:
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v110 = v106<<(uint(int32(16))%32) + v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 3:
			v103 = v9
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 4:
			v100 = v9
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 5:
			v95 = v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 6:
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v95 = v91<<(uint(int32(16))%32) + v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 7:
			v86 = v9
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 8:
			v81 = v9
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 9:
			v76 = v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 10:
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v76 = v72<<(uint(int32(24))%32) + v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	}
	v241 = int32(14)
	v243 = v237 ^ v238 - base.I32_rotl(v237, v241)
	v247 = v243 ^ v236 - base.I32_rotl(v243, int32(11))
	v251 = v247 ^ v237 - base.I32_rotl(v247, int32(25))
	v255 = v251 ^ v243 - base.I32_rotl(v251, int32(16))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(4))
	v263 = v259 ^ v251 - base.I32_rotl(v259, v241)
	return v263 ^ v255 - base.I32_rotl(v263, int32(24))
}
func F_hashoptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(8)
	v7 = F_build_reloptions(m, l0, l1, v3, v3, int32(_a_F_hashoptions_0), int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_hashrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	if v8 == int32(-1) {
		F__hash_dropscanbuf(m, v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(-4294967296)
			if l1 == v18 {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v28 <= int32(0) {
				} else {
					v32 = v28 * int32(48)
					if v32 == int32(0) {
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v35, l1, v32)
					}
				}
			}
			v38 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v38)
			return
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		if v11 <= int32(0) {
			F__hash_dropscanbuf(m, v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v18
				*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(-4294967296)
				if l1 == v18 {
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v28 <= int32(0) {
					} else {
						v32 = v28 * int32(48)
						if v32 == int32(0) {
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							base.MemoryCopy(m, v35, l1, v32)
						}
					}
				}
				v38 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v38)
				return
			}
		} else {
			F__hash_kill_items(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F__hash_dropscanbuf(m, v7)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v18
					*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(-4294967296)
					if l1 == v18 {
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v28 <= int32(0) {
						} else {
							v32 = v28 * int32(48)
							if v32 == int32(0) {
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								base.MemoryCopy(m, v35, l1, v32)
							}
						}
					}
					v38 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v38)
					return
				}
			}
		}
	}
}
func F_hashtextextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int64
	_ = v413
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
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
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v14 != 0 {
			v15 = F_pg_newlocale_from_collation(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(1)
				v18 = v10 + v17
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				v23 = v21 & v17
				if v23 != 0 {
					v24 = v18
				} else {
					v24 = v10 + int32(4)
				}
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
				if v25 == int32(1) {
					if v21 == int32(1) {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v33 == int32(18) {
							v36 = int32(16)
						} else {
							v36 = int32(0)
						}
						if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v43 = int32(4)
						} else {
							v43 = v36
						}
						v54 = v43
					} else {
						v44 = int32(1)
						if v23 != 0 {
							v54 = int32(base.Ui32(v21)>>(uint(v44)%32)) - v44
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
					v62 = v54 - int32(1636608432)
					if v56 == int64(0) {
						v99 = v62
						v101 = v62
						v103 = v62
					} else {
						v66 = v62 + base.I32_wrap_i64(v56)
						v67 = v66 + v62
						v71 = int32(4)
						v73 = base.I32_wrap_i64(int64(base.Ui64(v56)>>(uint(int64(32))%64))) ^ base.I32_rotl(v62, v71)
						v77 = v66 - v73 ^ base.I32_rotl(v73, int32(6))
						v81 = v67 - v77 ^ base.I32_rotl(v77, int32(8))
						v82 = v67 + v73
						v83 = v77 + v82
						v84 = v81 + v83
						v88 = v82 - v81 ^ base.I32_rotl(v81, int32(16))
						v92 = v83 - v88 ^ base.I32_rotl(v88, int32(19))
						v97 = v84 + v88
						v99 = v97
						v101 = v84 - v92 ^ base.I32_rotl(v92, v71)
						v103 = v92 + v97
					}
					if v24&int32(3) != 0 {
						if base.Ui32(int32(11)) < base.Ui32(v54) {
							v108 = v24
							v109 = v54
							v111 = v99
							v112 = v103
							v113 = v101
							for {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
								v116 = v115 + v112
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
								v120 = v119 + v113
								v122 = int32(4)
								v124 = v117 + v111 - v120 ^ base.I32_rotl(v120, v122)
								v128 = v116 - v124 ^ base.I32_rotl(v124, int32(6))
								v129 = v120 + v116
								v130 = v124 + v129
								v131 = v128 + v130
								v135 = v129 - v128 ^ base.I32_rotl(v128, int32(8))
								v139 = v130 - v135 ^ base.I32_rotl(v135, int32(16))
								v143 = v131 - v139 ^ base.I32_rotl(v139, int32(19))
								v144 = v135 + v131
								v145 = v139 + v144
								v146 = v143 + v145
								v150 = v144 - v143 ^ base.I32_rotl(v143, v122)
								v151 = int32(12)
								v152 = v108 + v151
								v154 = v109 - v151
								if base.Ui32(int32(11)) < base.Ui32(v154) {
									v108 = v152
									v109 = v154
									v111 = v145
									v112 = v146
									v113 = v150
									continue
								} else {
									break
								}
								break
							}
							v157 = v152
							v158 = v154
							v160 = v145
							v161 = v146
							v162 = v150
						} else {
							v157 = v24
							v158 = v54
							v160 = v99
							v161 = v103
							v162 = v101
						}
						switch v158 - int32(1) {
						case 0:
							v327 = v160
							v328 = v161
							v329 = v162
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 1:
							v320 = v160
							v321 = v161
							v322 = v162
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 2:
							v313 = v160
							v314 = v161
							v315 = v162
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 3:
							v307 = v161
							v308 = v162
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 4:
							v303 = v161
							v304 = v162
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 5:
							v297 = v161
							v298 = v162
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
							v303 = v299<<(uint(int32(8))%32) + v297
							v304 = v298
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 6:
							v291 = v161
							v292 = v162
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
							v297 = v293<<(uint(int32(16))%32) + v291
							v298 = v292
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
							v303 = v299<<(uint(int32(8))%32) + v297
							v304 = v298
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 7:
							v286 = v162
							v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+7)))
							v291 = v287<<(uint(int32(24))%32) + v161
							v292 = v286
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
							v297 = v293<<(uint(int32(16))%32) + v291
							v298 = v292
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
							v303 = v299<<(uint(int32(8))%32) + v297
							v304 = v298
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 8:
							v281 = v162
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+8)))
							v286 = v282<<(uint(int32(8))%32) + v281
							v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+7)))
							v291 = v287<<(uint(int32(24))%32) + v161
							v292 = v286
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
							v297 = v293<<(uint(int32(16))%32) + v291
							v298 = v292
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
							v303 = v299<<(uint(int32(8))%32) + v297
							v304 = v298
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 9:
							v276 = v162
							v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+9)))
							v281 = v277<<(uint(int32(16))%32) + v276
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+8)))
							v286 = v282<<(uint(int32(8))%32) + v281
							v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+7)))
							v291 = v287<<(uint(int32(24))%32) + v161
							v292 = v286
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
							v297 = v293<<(uint(int32(16))%32) + v291
							v298 = v292
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
							v303 = v299<<(uint(int32(8))%32) + v297
							v304 = v298
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						case 10:
							v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+10)))
							v276 = v272<<(uint(int32(24))%32) + v162
							v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+9)))
							v281 = v277<<(uint(int32(16))%32) + v276
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+8)))
							v286 = v282<<(uint(int32(8))%32) + v281
							v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+7)))
							v291 = v287<<(uint(int32(24))%32) + v161
							v292 = v286
							v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
							v297 = v293<<(uint(int32(16))%32) + v291
							v298 = v292
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
							v303 = v299<<(uint(int32(8))%32) + v297
							v304 = v298
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
							v307 = v303 + v305
							v308 = v304
							v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
							v313 = v309<<(uint(int32(24))%32) + v160
							v314 = v307
							v315 = v308
							v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
							v320 = v316<<(uint(int32(16))%32) + v313
							v321 = v314
							v322 = v315
							v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
							v327 = v323<<(uint(int32(8))%32) + v320
							v328 = v321
							v329 = v322
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v335 = v327 + v330
							v336 = v328
							v337 = v329
						default:
							v335 = v160
							v336 = v161
							v337 = v162
						}
					} else {
						if base.Ui32(int32(12)) <= base.Ui32(v54) {
							v168 = v24
							v169 = v54
							v171 = v99
							v172 = v103
							v173 = v101
							for {
								v175 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
								v176 = v175 + v172
								v177 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
								v180 = v179 + v173
								v182 = int32(4)
								v184 = v177 + v171 - v180 ^ base.I32_rotl(v180, v182)
								v188 = v176 - v184 ^ base.I32_rotl(v184, int32(6))
								v189 = v180 + v176
								v190 = v184 + v189
								v191 = v188 + v190
								v195 = v189 - v188 ^ base.I32_rotl(v188, int32(8))
								v199 = v190 - v195 ^ base.I32_rotl(v195, int32(16))
								v203 = v191 - v199 ^ base.I32_rotl(v199, int32(19))
								v204 = v195 + v191
								v205 = v199 + v204
								v206 = v203 + v205
								v210 = v204 - v203 ^ base.I32_rotl(v203, v182)
								v211 = int32(12)
								v212 = v168 + v211
								v214 = v169 - v211
								if base.Ui32(int32(11)) < base.Ui32(v214) {
									v168 = v212
									v169 = v214
									v171 = v205
									v172 = v206
									v173 = v210
									continue
								} else {
									break
								}
								break
							}
							v217 = v212
							v218 = v214
							v220 = v205
							v221 = v206
							v222 = v210
						} else {
							v217 = v24
							v218 = v54
							v220 = v99
							v221 = v103
							v222 = v101
						}
						switch v218 - int32(1) {
						case 0:
							v269 = v220
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
							v335 = v269 + v270
							v336 = v221
							v337 = v222
						case 1:
							v264 = v220
							v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
							v269 = v265<<(uint(int32(8))%32) + v264
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
							v335 = v269 + v270
							v336 = v221
							v337 = v222
						case 2:
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+2)))
							v264 = v260<<(uint(int32(16))%32) + v220
							v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
							v269 = v265<<(uint(int32(8))%32) + v264
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
							v335 = v269 + v270
							v336 = v221
							v337 = v222
						case 3:
							v257 = v221
							v258 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v335 = v258 + v220
							v336 = v257
							v337 = v222
						case 4:
							v254 = v221
							v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)))
							v257 = v254 + v255
							v258 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v335 = v258 + v220
							v336 = v257
							v337 = v222
						case 5:
							v249 = v221
							v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+5)))
							v254 = v250<<(uint(int32(8))%32) + v249
							v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)))
							v257 = v254 + v255
							v258 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v335 = v258 + v220
							v336 = v257
							v337 = v222
						case 6:
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+6)))
							v249 = v245<<(uint(int32(16))%32) + v221
							v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+5)))
							v254 = v250<<(uint(int32(8))%32) + v249
							v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)))
							v257 = v254 + v255
							v258 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v335 = v258 + v220
							v336 = v257
							v337 = v222
						case 7:
							v240 = v222
							v241 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
							v335 = v241 + v220
							v336 = v243 + v221
							v337 = v240
						case 8:
							v235 = v222
							v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)))
							v240 = v236<<(uint(int32(8))%32) + v235
							v241 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
							v335 = v241 + v220
							v336 = v243 + v221
							v337 = v240
						case 9:
							v230 = v222
							v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+9)))
							v235 = v231<<(uint(int32(16))%32) + v230
							v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)))
							v240 = v236<<(uint(int32(8))%32) + v235
							v241 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
							v335 = v241 + v220
							v336 = v243 + v221
							v337 = v240
						case 10:
							v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
							v230 = v226<<(uint(int32(24))%32) + v222
							v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+9)))
							v235 = v231<<(uint(int32(16))%32) + v230
							v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)))
							v240 = v236<<(uint(int32(8))%32) + v235
							v241 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
							v335 = v241 + v220
							v336 = v243 + v221
							v337 = v240
						default:
							v335 = v220
							v336 = v221
							v337 = v222
						}
					}
					v340 = int32(14)
					v342 = v336 ^ v337 - base.I32_rotl(v336, v340)
					v346 = v342 ^ v335 - base.I32_rotl(v342, int32(11))
					v350 = v346 ^ v336 - base.I32_rotl(v346, int32(25))
					v354 = v350 ^ v342 - base.I32_rotl(v350, int32(16))
					v358 = v354 ^ v346 - base.I32_rotl(v354, int32(4))
					v362 = v358 ^ v350 - base.I32_rotl(v358, v340)
					v372 = F_Int64GetDatum(m, base.I64_extend_i32_u(v362)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v362^v354-base.I32_rotl(v362, int32(24))))
					mBase = m.M
					v373 = m.ExcPending
					if v373 != 0 {
						return int32(0)
					} else {
						v733 = v372
						v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v737 != v10 {
							F_pfree(m, v10)
							mBase = m.M
							v740 = m.ExcPending
							if v740 != 0 {
								return int32(0)
							} else {
								return v733
							}
						} else {
							return v733
						}
					}
				} else {
					v374 = int32(0)
					if v21 == int32(1) {
						v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v381 == int32(18) {
							v384 = int32(16)
						} else {
							v384 = int32(0)
						}
						if base.Ui32((v381-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v391 = int32(4)
						} else {
							v391 = v384
						}
						v402 = v391
					} else {
						v392 = int32(1)
						if v23 != 0 {
							v402 = int32(base.Ui32(v21)>>(uint(v392)%32)) - v392
						} else {
							v396 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v402 = int32(base.Ui32(v396)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v403 = F_pg_strnxfrm(m, v374, v374, v24, v402, v15)
					mBase = m.M
					v404 = m.ExcPending
					if v404 != 0 {
						return int32(0)
					} else {
						v406 = v403 + int32(1)
						v407 = F_palloc(m, v406)
						mBase = m.M
						v408 = m.ExcPending
						if v408 != 0 {
							return int32(0)
						} else {
							v409 = F_pg_strnxfrm(m, v407, v406, v24, v402, v15)
							mBase = m.M
							v410 = m.ExcPending
							if v410 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v403) < base.Ui32(v409) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v765 = m.ExcPending
									if v765 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_hashtextextended_0), int32(0))
										mBase = m.M
										v769 = m.ExcPending
										if v769 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_hashtextextended_1), int32(361), int32(_a_F_hashtextextended_2))
											mBase = m.M
											v774 = m.ExcPending
											if v774 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v413 = *(*int64)(unsafe.Add(mBase, uint32(v412)))
									v419 = v406 - int32(1636608432)
									if v413 == int64(0) {
										v456 = v419
										v458 = v419
										v460 = v419
									} else {
										v423 = v419 + base.I32_wrap_i64(v413)
										v424 = v423 + v419
										v428 = int32(4)
										v430 = base.I32_wrap_i64(int64(base.Ui64(v413)>>(uint(int64(32))%64))) ^ base.I32_rotl(v419, v428)
										v434 = v423 - v430 ^ base.I32_rotl(v430, int32(6))
										v438 = v424 - v434 ^ base.I32_rotl(v434, int32(8))
										v439 = v424 + v430
										v440 = v434 + v439
										v441 = v438 + v440
										v445 = v439 - v438 ^ base.I32_rotl(v438, int32(16))
										v449 = v440 - v445 ^ base.I32_rotl(v445, int32(19))
										v454 = v441 + v445
										v456 = v454
										v458 = v441 - v449 ^ base.I32_rotl(v449, v428)
										v460 = v449 + v454
									}
									if v407&int32(3) != 0 {
										if base.Ui32(int32(11)) < base.Ui32(v406) {
											v465 = v407
											v466 = v406
											v468 = v456
											v469 = v460
											v470 = v458
											for {
												v472 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
												v473 = v472 + v469
												v474 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
												v476 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
												v477 = v476 + v470
												v479 = int32(4)
												v481 = v474 + v468 - v477 ^ base.I32_rotl(v477, v479)
												v485 = v473 - v481 ^ base.I32_rotl(v481, int32(6))
												v486 = v477 + v473
												v487 = v481 + v486
												v488 = v485 + v487
												v492 = v486 - v485 ^ base.I32_rotl(v485, int32(8))
												v496 = v487 - v492 ^ base.I32_rotl(v492, int32(16))
												v500 = v488 - v496 ^ base.I32_rotl(v496, int32(19))
												v501 = v492 + v488
												v502 = v496 + v501
												v503 = v500 + v502
												v507 = v501 - v500 ^ base.I32_rotl(v500, v479)
												v508 = int32(12)
												v509 = v465 + v508
												v511 = v466 - v508
												if base.Ui32(int32(11)) < base.Ui32(v511) {
													v465 = v509
													v466 = v511
													v468 = v502
													v469 = v503
													v470 = v507
													continue
												} else {
													break
												}
												break
											}
											v514 = v509
											v515 = v511
											v517 = v502
											v518 = v503
											v519 = v507
										} else {
											v514 = v407
											v515 = v406
											v517 = v456
											v518 = v460
											v519 = v458
										}
										switch v515 - int32(1) {
										case 0:
											v684 = v517
											v685 = v518
											v686 = v519
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 1:
											v677 = v517
											v678 = v518
											v679 = v519
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 2:
											v670 = v517
											v671 = v518
											v672 = v519
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 3:
											v664 = v518
											v665 = v519
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 4:
											v660 = v518
											v661 = v519
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 5:
											v654 = v518
											v655 = v519
											v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
											v660 = v656<<(uint(int32(8))%32) + v654
											v661 = v655
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 6:
											v648 = v518
											v649 = v519
											v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+6)))
											v654 = v650<<(uint(int32(16))%32) + v648
											v655 = v649
											v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
											v660 = v656<<(uint(int32(8))%32) + v654
											v661 = v655
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 7:
											v643 = v519
											v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+7)))
											v648 = v644<<(uint(int32(24))%32) + v518
											v649 = v643
											v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+6)))
											v654 = v650<<(uint(int32(16))%32) + v648
											v655 = v649
											v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
											v660 = v656<<(uint(int32(8))%32) + v654
											v661 = v655
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 8:
											v638 = v519
											v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+8)))
											v643 = v639<<(uint(int32(8))%32) + v638
											v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+7)))
											v648 = v644<<(uint(int32(24))%32) + v518
											v649 = v643
											v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+6)))
											v654 = v650<<(uint(int32(16))%32) + v648
											v655 = v649
											v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
											v660 = v656<<(uint(int32(8))%32) + v654
											v661 = v655
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 9:
											v633 = v519
											v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+9)))
											v638 = v634<<(uint(int32(16))%32) + v633
											v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+8)))
											v643 = v639<<(uint(int32(8))%32) + v638
											v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+7)))
											v648 = v644<<(uint(int32(24))%32) + v518
											v649 = v643
											v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+6)))
											v654 = v650<<(uint(int32(16))%32) + v648
											v655 = v649
											v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
											v660 = v656<<(uint(int32(8))%32) + v654
											v661 = v655
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										case 10:
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+10)))
											v633 = v629<<(uint(int32(24))%32) + v519
											v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+9)))
											v638 = v634<<(uint(int32(16))%32) + v633
											v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+8)))
											v643 = v639<<(uint(int32(8))%32) + v638
											v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+7)))
											v648 = v644<<(uint(int32(24))%32) + v518
											v649 = v643
											v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+6)))
											v654 = v650<<(uint(int32(16))%32) + v648
											v655 = v649
											v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
											v660 = v656<<(uint(int32(8))%32) + v654
											v661 = v655
											v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
											v664 = v660 + v662
											v665 = v661
											v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
											v670 = v666<<(uint(int32(24))%32) + v517
											v671 = v664
											v672 = v665
											v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
											v677 = v673<<(uint(int32(16))%32) + v670
											v678 = v671
											v679 = v672
											v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
											v684 = v680<<(uint(int32(8))%32) + v677
											v685 = v678
											v686 = v679
											v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
											v692 = v684 + v687
											v693 = v685
											v694 = v686
										default:
											v692 = v517
											v693 = v518
											v694 = v519
										}
									} else {
										if base.Ui32(int32(12)) <= base.Ui32(v406) {
											v525 = v407
											v526 = v406
											v528 = v456
											v529 = v460
											v530 = v458
											for {
												v532 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
												v533 = v532 + v529
												v534 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
												v536 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
												v537 = v536 + v530
												v539 = int32(4)
												v541 = v534 + v528 - v537 ^ base.I32_rotl(v537, v539)
												v545 = v533 - v541 ^ base.I32_rotl(v541, int32(6))
												v546 = v537 + v533
												v547 = v541 + v546
												v548 = v545 + v547
												v552 = v546 - v545 ^ base.I32_rotl(v545, int32(8))
												v556 = v547 - v552 ^ base.I32_rotl(v552, int32(16))
												v560 = v548 - v556 ^ base.I32_rotl(v556, int32(19))
												v561 = v552 + v548
												v562 = v556 + v561
												v563 = v560 + v562
												v567 = v561 - v560 ^ base.I32_rotl(v560, v539)
												v568 = int32(12)
												v569 = v525 + v568
												v571 = v526 - v568
												if base.Ui32(int32(11)) < base.Ui32(v571) {
													v525 = v569
													v526 = v571
													v528 = v562
													v529 = v563
													v530 = v567
													continue
												} else {
													break
												}
												break
											}
											v574 = v569
											v575 = v571
											v577 = v562
											v578 = v563
											v579 = v567
										} else {
											v574 = v407
											v575 = v406
											v577 = v456
											v578 = v460
											v579 = v458
										}
										switch v575 - int32(1) {
										case 0:
											v626 = v577
											v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
											v692 = v626 + v627
											v693 = v578
											v694 = v579
										case 1:
											v621 = v577
											v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+1)))
											v626 = v622<<(uint(int32(8))%32) + v621
											v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
											v692 = v626 + v627
											v693 = v578
											v694 = v579
										case 2:
											v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)))
											v621 = v617<<(uint(int32(16))%32) + v577
											v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+1)))
											v626 = v622<<(uint(int32(8))%32) + v621
											v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
											v692 = v626 + v627
											v693 = v578
											v694 = v579
										case 3:
											v614 = v578
											v615 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v692 = v615 + v577
											v693 = v614
											v694 = v579
										case 4:
											v611 = v578
											v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+4)))
											v614 = v611 + v612
											v615 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v692 = v615 + v577
											v693 = v614
											v694 = v579
										case 5:
											v606 = v578
											v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+5)))
											v611 = v607<<(uint(int32(8))%32) + v606
											v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+4)))
											v614 = v611 + v612
											v615 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v692 = v615 + v577
											v693 = v614
											v694 = v579
										case 6:
											v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+6)))
											v606 = v602<<(uint(int32(16))%32) + v578
											v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+5)))
											v611 = v607<<(uint(int32(8))%32) + v606
											v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+4)))
											v614 = v611 + v612
											v615 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v692 = v615 + v577
											v693 = v614
											v694 = v579
										case 7:
											v597 = v579
											v598 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
											v692 = v598 + v577
											v693 = v600 + v578
											v694 = v597
										case 8:
											v592 = v579
											v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+8)))
											v597 = v593<<(uint(int32(8))%32) + v592
											v598 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
											v692 = v598 + v577
											v693 = v600 + v578
											v694 = v597
										case 9:
											v587 = v579
											v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+9)))
											v592 = v588<<(uint(int32(16))%32) + v587
											v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+8)))
											v597 = v593<<(uint(int32(8))%32) + v592
											v598 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
											v692 = v598 + v577
											v693 = v600 + v578
											v694 = v597
										case 10:
											v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+10)))
											v587 = v583<<(uint(int32(24))%32) + v579
											v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+9)))
											v592 = v588<<(uint(int32(16))%32) + v587
											v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+8)))
											v597 = v593<<(uint(int32(8))%32) + v592
											v598 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
											v692 = v598 + v577
											v693 = v600 + v578
											v694 = v597
										default:
											v692 = v577
											v693 = v578
											v694 = v579
										}
									}
									v697 = int32(14)
									v699 = v693 ^ v694 - base.I32_rotl(v693, v697)
									v703 = v699 ^ v692 - base.I32_rotl(v699, int32(11))
									v707 = v703 ^ v693 - base.I32_rotl(v703, int32(25))
									v711 = v707 ^ v699 - base.I32_rotl(v707, int32(16))
									v715 = v711 ^ v703 - base.I32_rotl(v711, int32(4))
									v719 = v715 ^ v707 - base.I32_rotl(v715, v697)
									v729 = F_Int64GetDatum(m, base.I64_extend_i32_u(v719)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v719^v711-base.I32_rotl(v719, int32(24))))
									mBase = m.M
									v730 = m.ExcPending
									if v730 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v407)
										mBase = m.M
										v732 = m.ExcPending
										if v732 != 0 {
											return int32(0)
										} else {
											v733 = v729
											v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v737 != v10 {
												F_pfree(m, v10)
												mBase = m.M
												v740 = m.ExcPending
												if v740 != 0 {
													return int32(0)
												} else {
													return v733
												}
											} else {
												return v733
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v745 = m.ExcPending
			if v745 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v748 = m.ExcPending
				if v748 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_hashtextextended_3), int32(0))
					mBase = m.M
					v752 = m.ExcPending
					if v752 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_hashtextextended_4), int32(0))
						mBase = m.M
						v756 = m.ExcPending
						if v756 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hashtextextended_1), int32(336), int32(_a_F_hashtextextended_2))
							mBase = m.M
							v761 = m.ExcPending
							if v761 != 0 {
								return int32(0)
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
func F_hashvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v57 int32
	_ = v57
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int64
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
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
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v584 int32
	_ = v584
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(208)
	m.G0 = v17
	v20 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if int32(0) < v240 {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	return int32(0)
L3:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v26 = v24 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+80))
	v29 = F_get_opfamily_name(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L49
	}
L7:
	;
	v33 = int32(0)
	v35 = F_SearchSysCacheList(m, int32(4), int32(1), v28, v33, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(1)
	v40 = int32(0)
	v42 = F_SearchSysCacheList(m, int32(5), v37, v28, v40, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v44 <= int32(0) {
		v229 = v37
		v234 = v2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v52 = v37
	v53 = v2
	v57 = v2
	goto L11
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(48)+v53<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v69 = v67 + v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v70 == v71 {
		v100 = v52
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v229 = v206
	v234 = v207
	goto L1
L13:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
	switch v101 - int32(1) {
	case 0:
		goto L24
	case 1:
		goto L25
	case 2:
		goto L27
	default:
		goto L26
	}
L14:
	;
	v73 = int32(0)
	v76 = F_errstart(m, int32(17), v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v76 == int32(0) {
		v100 = v73
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v84 = F_format_procedure(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_1), v17+int32(192))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(91), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v100 = v73
	goto L13
L21:
	;
	v210 = v53 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v210 < v211 {
		v52 = v206
		v53 = v210
		v57 = v207
		goto L11
	} else {
		goto L48
	}
L22:
	;
	v176 = int32(0)
	v179 = F_errstart(m, int32(17), v176)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L42
	}
L23:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
	v166 = int32(1)
	if base.Ui32(v166) < base.Ui32((v165-v166)&int32(_a_F_hashvalidate_4)) {
		v206 = v100
		v207 = v57
		goto L21
	} else {
		goto L40
	}
L24:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v152
	v155 = int32(1)
	v160 = F_check_amproc_signature(m, v151, int32(23), v155, v155, v155, v17+int32(160))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L38
	}
L25:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v140 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v139
	v145 = int32(2)
	v149 = F_check_amproc_signature(m, v138, v140, int32(1), v145, v145, v17+int32(176))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L36
	}
L26:
	;
	v109 = int32(0)
	v112 = F_errstart(m, int32(17), v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L30
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v105 = F_check_amoptsproc_signature(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v105 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	if v112 == int32(0) {
		v206 = v109
		v207 = v57
		goto L21
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v120 = F_format_procedure(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_5), v17+int32(128))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(115), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v206 = v109
	v207 = v57
	goto L21
L36:
	;
	if v149 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	if v160 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v173 = F_list_append_unique_oid(m, v57, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v206 = v100
	v207 = v173
	goto L21
L42:
	;
	if v179 == int32(0) {
		v206 = v176
		v207 = v57
		goto L21
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v187 = F_format_procedure(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v17)+148)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_6), v17+int32(144))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(127), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v206 = v176
	v207 = v57
	goto L21
L48:
	;
	goto L12
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg_internal(m, int32(_a_F_hashvalidate_7), v17)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(60), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v249 = v229
	v250 = int32(0)
	goto L55
L53:
	;
	v484 = v229
	goto L54
L54:
	;
	v495 = F_identify_opfamily_groups(m, v35, v42)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L126
	}
L55:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(48)+v250<<(uint(int32(2))%32))))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+56))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+22)))
	v266 = v264 + v265
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+16)))
	if v267 == int32(1) {
		v300 = v249
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v484 = v476
	goto L54
L57:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+18)))
	if v301 == int32(115) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v270 = int32(0)
	v273 = F_errstart(m, int32(17), v270)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	if v273 == int32(0) {
		v300 = v270
		goto L57
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v281 = F_format_operator(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v266)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_8), v17+int32(112))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(153), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v300 = v270
	goto L57
L65:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v339 = F_check_amop_signature(m, v335, int32(16), v337, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L77
	}
L66:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v266)+28))
	if v304 == int32(0) {
		v334 = v300
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v307 = int32(0)
	v310 = F_errstart(m, int32(17), v307)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v310 == int32(0) {
		v334 = v307
		goto L65
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v318 = F_format_operator(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_9), v17+int32(96))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(165), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v334 = v307
	goto L65
L76:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v370 = int32(0)
	if v234 == v370 {
		goto L87
	} else {
		goto L88
	}
L77:
	;
	if v339 != 0 {
		v368 = v334
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v341 = int32(0)
	v344 = F_errstart(m, int32(17), v341)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v344 == int32(0) {
		v368 = v341
		goto L76
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v352 = F_format_operator(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_10), v17+int32(80))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(178), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v368 = v341
	goto L76
L85:
	;
	v478 = v250 + int32(1)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if v478 < v479 {
		v249 = v476
		v250 = v478
		goto L55
	} else {
		goto L122
	}
L86:
	;
	if v408 != 0 {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v408 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v376 <= int32(0) {
		v402 = v370
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v408 = v402
	goto L86
L91:
	;
	v379 = int32(0)
	if v379 < v376 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v382 = v376
	goto L94
L93:
	;
	v382 = v379
	goto L94
L94:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v385 = int32(0)
	goto L95
L95:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v383+v385<<(uint(int32(2))%32))))
	v394 = base.B2i32(v393 == v369)
	if v393 == v369 {
		v402 = v394
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v402 = v394
	goto L90
L97:
	;
	v396 = v385 + int32(1)
	if v396 != v382 {
		v385 = v396
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v410 = int32(0)
	if v234 == v410 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	v449 = int32(0)
	v452 = F_errstart(m, int32(17), v449)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L2
	} else {
		goto L116
	}
L102:
	;
	if v448 != 0 {
		v476 = v368
		goto L85
	} else {
		goto L115
	}
L103:
	;
	v448 = int32(0)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v416 <= int32(0) {
		v442 = v410
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v448 = v442
	goto L102
L107:
	;
	v419 = int32(0)
	if v419 < v416 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v422 = v416
	goto L110
L109:
	;
	v422 = v419
	goto L110
L110:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v425 = int32(0)
	goto L111
L111:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v423+v425<<(uint(int32(2))%32))))
	v434 = base.B2i32(v433 == v409)
	if v433 == v409 {
		v442 = v434
		goto L106
	} else {
		goto L113
	}
L112:
	;
	v442 = v434
	goto L106
L113:
	;
	v436 = v425 + int32(1)
	if v436 != v422 {
		v425 = v436
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	goto L101
L116:
	;
	if v452 == int32(0) {
		v476 = v449
		goto L85
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v460 = F_format_operator(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_11), v17-int32(-64))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(190), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v476 = v449
	goto L85
L122:
	;
	goto L56
L123:
	;
	F_ReleaseCatCacheList(m, v42)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L2
	} else {
		goto L169
	}
L124:
	;
	if v234 != 0 {
		goto L160
	} else {
		goto L161
	}
L125:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v645 = v630
	v656 = v641
	goto L124
L126:
	;
	if v495 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if int32(0) < v497 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v601 = int32(0)
	v604 = F_errstart(m, int32(17), v601)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L2
	} else {
		goto L152
	}
L130:
	;
	v500 = int32(0)
	v504 = v500
	v505 = v484
	v506 = v500
	goto L133
L131:
	;
	v573 = v484
	v584 = int32(1)
	goto L132
L132:
	;
	if v584 == int32(0) {
		v630 = v573
		goto L125
	} else {
		goto L151
	}
L133:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v516+v506<<(uint(int32(2))%32))))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	if v27 == v521 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v573 = v561
	v584 = base.B2i32(v526 == int32(0))
	goto L132
L135:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v523 == v27 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v526 = v504
	goto L137
L137:
	;
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v520)+8))
	if v527 == int64(2) {
		v561 = v505
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v525 = v520
	goto L140
L139:
	;
	v525 = v504
	goto L140
L140:
	;
	v526 = v525
	goto L137
L141:
	;
	v564 = v506 + int32(1)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v564 < v565 {
		v504 = v526
		v505 = v561
		v506 = v564
		goto L133
	} else {
		goto L150
	}
L142:
	;
	v530 = int32(0)
	v533 = F_errstart(m, int32(17), v530)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	if v533 == int32(0) {
		v561 = v530
		goto L141
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v541 = F_format_type_be(m, v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	v544 = F_format_type_be(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L2
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_12), v17+int32(48))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(219), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	v561 = v530
	goto L141
L150:
	;
	goto L134
L151:
	;
	goto L129
L152:
	;
	if v604 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L2
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v624 = int32(0)
	if v495 == v624 {
		v645 = v601
		v656 = v624
		goto L124
	} else {
		goto L159
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v26 + int32(8)
	F_errmsg(m, int32(_a_F_hashvalidate_13), v17+int32(32))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(231), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	goto L155
L159:
	;
	v630 = v601
	goto L125
L160:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v661 = v657 * v657
	goto L162
L161:
	;
	v661 = int32(0)
	goto L162
L162:
	;
	if v656 == v661 {
		v686 = v645
		goto L123
	} else {
		goto L163
	}
L163:
	;
	v663 = int32(0)
	v666 = F_errstart(m, int32(17), v663)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	if v666 == int32(0) {
		v686 = v663
		goto L123
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_14), v17+int32(16))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(247), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	v686 = v663
	goto L123
L169:
	;
	F_ReleaseCatCacheList(m, v35)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	m.G0 = v17 + int32(208)
	return v686
}
func F_hashvarlenaextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v17 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
		v53 = v45 - int32(1636608432)
		if v47 == int64(0) {
			v90 = v53
			v92 = v53
			v94 = v53
		} else {
			v57 = v53 + base.I32_wrap_i64(v47)
			v58 = v57 + v53
			v62 = int32(4)
			v64 = base.I32_wrap_i64(int64(base.Ui64(v47)>>(uint(int64(32))%64))) ^ base.I32_rotl(v53, v62)
			v68 = v57 - v64 ^ base.I32_rotl(v64, int32(6))
			v72 = v58 - v68 ^ base.I32_rotl(v68, int32(8))
			v73 = v58 + v64
			v74 = v68 + v73
			v75 = v72 + v74
			v79 = v73 - v72 ^ base.I32_rotl(v72, int32(16))
			v83 = v74 - v79 ^ base.I32_rotl(v79, int32(19))
			v88 = v75 + v79
			v90 = v88
			v92 = v75 - v83 ^ base.I32_rotl(v83, v62)
			v94 = v83 + v88
		}
		if v18&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v45) {
				v99 = v18
				v100 = v45
				v102 = v90
				v103 = v94
				v104 = v92
				for {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
					v107 = v106 + v103
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
					v111 = v110 + v104
					v113 = int32(4)
					v115 = v108 + v102 - v111 ^ base.I32_rotl(v111, v113)
					v119 = v107 - v115 ^ base.I32_rotl(v115, int32(6))
					v120 = v111 + v107
					v121 = v115 + v120
					v122 = v119 + v121
					v126 = v120 - v119 ^ base.I32_rotl(v119, int32(8))
					v130 = v121 - v126 ^ base.I32_rotl(v126, int32(16))
					v134 = v122 - v130 ^ base.I32_rotl(v130, int32(19))
					v135 = v126 + v122
					v136 = v130 + v135
					v137 = v134 + v136
					v141 = v135 - v134 ^ base.I32_rotl(v134, v113)
					v142 = int32(12)
					v143 = v99 + v142
					v145 = v100 - v142
					if base.Ui32(int32(11)) < base.Ui32(v145) {
						v99 = v143
						v100 = v145
						v102 = v136
						v103 = v137
						v104 = v141
						continue
					} else {
						break
					}
					break
				}
				v148 = v143
				v149 = v145
				v151 = v136
				v152 = v137
				v153 = v141
			} else {
				v148 = v18
				v149 = v45
				v151 = v90
				v152 = v94
				v153 = v92
			}
			switch v149 - int32(1) {
			case 0:
				v318 = v151
				v319 = v152
				v320 = v153
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 1:
				v311 = v151
				v312 = v152
				v313 = v153
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 2:
				v304 = v151
				v305 = v152
				v306 = v153
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 3:
				v298 = v152
				v299 = v153
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 4:
				v294 = v152
				v295 = v153
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 5:
				v288 = v152
				v289 = v153
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
				v294 = v290<<(uint(int32(8))%32) + v288
				v295 = v289
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 6:
				v282 = v152
				v283 = v153
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+6)))
				v288 = v284<<(uint(int32(16))%32) + v282
				v289 = v283
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
				v294 = v290<<(uint(int32(8))%32) + v288
				v295 = v289
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 7:
				v277 = v153
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+7)))
				v282 = v278<<(uint(int32(24))%32) + v152
				v283 = v277
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+6)))
				v288 = v284<<(uint(int32(16))%32) + v282
				v289 = v283
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
				v294 = v290<<(uint(int32(8))%32) + v288
				v295 = v289
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 8:
				v272 = v153
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+8)))
				v277 = v273<<(uint(int32(8))%32) + v272
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+7)))
				v282 = v278<<(uint(int32(24))%32) + v152
				v283 = v277
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+6)))
				v288 = v284<<(uint(int32(16))%32) + v282
				v289 = v283
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
				v294 = v290<<(uint(int32(8))%32) + v288
				v295 = v289
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 9:
				v267 = v153
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+9)))
				v272 = v268<<(uint(int32(16))%32) + v267
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+8)))
				v277 = v273<<(uint(int32(8))%32) + v272
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+7)))
				v282 = v278<<(uint(int32(24))%32) + v152
				v283 = v277
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+6)))
				v288 = v284<<(uint(int32(16))%32) + v282
				v289 = v283
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
				v294 = v290<<(uint(int32(8))%32) + v288
				v295 = v289
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			case 10:
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+10)))
				v267 = v263<<(uint(int32(24))%32) + v153
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+9)))
				v272 = v268<<(uint(int32(16))%32) + v267
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+8)))
				v277 = v273<<(uint(int32(8))%32) + v272
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+7)))
				v282 = v278<<(uint(int32(24))%32) + v152
				v283 = v277
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+6)))
				v288 = v284<<(uint(int32(16))%32) + v282
				v289 = v283
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
				v294 = v290<<(uint(int32(8))%32) + v288
				v295 = v289
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
				v298 = v294 + v296
				v299 = v295
				v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)))
				v304 = v300<<(uint(int32(24))%32) + v151
				v305 = v298
				v306 = v299
				v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)))
				v311 = v307<<(uint(int32(16))%32) + v304
				v312 = v305
				v313 = v306
				v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
				v318 = v314<<(uint(int32(8))%32) + v311
				v319 = v312
				v320 = v313
				v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
				v326 = v318 + v321
				v327 = v319
				v328 = v320
			default:
				v326 = v151
				v327 = v152
				v328 = v153
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v45) {
				v159 = v18
				v160 = v45
				v162 = v90
				v163 = v94
				v164 = v92
				for {
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
					v167 = v166 + v163
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
					v171 = v170 + v164
					v173 = int32(4)
					v175 = v168 + v162 - v171 ^ base.I32_rotl(v171, v173)
					v179 = v167 - v175 ^ base.I32_rotl(v175, int32(6))
					v180 = v171 + v167
					v181 = v175 + v180
					v182 = v179 + v181
					v186 = v180 - v179 ^ base.I32_rotl(v179, int32(8))
					v190 = v181 - v186 ^ base.I32_rotl(v186, int32(16))
					v194 = v182 - v190 ^ base.I32_rotl(v190, int32(19))
					v195 = v186 + v182
					v196 = v190 + v195
					v197 = v194 + v196
					v201 = v195 - v194 ^ base.I32_rotl(v194, v173)
					v202 = int32(12)
					v203 = v159 + v202
					v205 = v160 - v202
					if base.Ui32(int32(11)) < base.Ui32(v205) {
						v159 = v203
						v160 = v205
						v162 = v196
						v163 = v197
						v164 = v201
						continue
					} else {
						break
					}
					break
				}
				v208 = v203
				v209 = v205
				v211 = v196
				v212 = v197
				v213 = v201
			} else {
				v208 = v18
				v209 = v45
				v211 = v90
				v212 = v94
				v213 = v92
			}
			switch v209 - int32(1) {
			case 0:
				v260 = v211
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v326 = v260 + v261
				v327 = v212
				v328 = v213
			case 1:
				v255 = v211
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v260 = v256<<(uint(int32(8))%32) + v255
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v326 = v260 + v261
				v327 = v212
				v328 = v213
			case 2:
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v255 = v251<<(uint(int32(16))%32) + v211
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v260 = v256<<(uint(int32(8))%32) + v255
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v326 = v260 + v261
				v327 = v212
				v328 = v213
			case 3:
				v248 = v212
				v249 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v326 = v249 + v211
				v327 = v248
				v328 = v213
			case 4:
				v245 = v212
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v248 = v245 + v246
				v249 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v326 = v249 + v211
				v327 = v248
				v328 = v213
			case 5:
				v240 = v212
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v245 = v241<<(uint(int32(8))%32) + v240
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v248 = v245 + v246
				v249 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v326 = v249 + v211
				v327 = v248
				v328 = v213
			case 6:
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
				v240 = v236<<(uint(int32(16))%32) + v212
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v245 = v241<<(uint(int32(8))%32) + v240
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v248 = v245 + v246
				v249 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v326 = v249 + v211
				v327 = v248
				v328 = v213
			case 7:
				v231 = v213
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v234 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
				v326 = v232 + v211
				v327 = v234 + v212
				v328 = v231
			case 8:
				v226 = v213
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
				v231 = v227<<(uint(int32(8))%32) + v226
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v234 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
				v326 = v232 + v211
				v327 = v234 + v212
				v328 = v231
			case 9:
				v221 = v213
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
				v226 = v222<<(uint(int32(16))%32) + v221
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
				v231 = v227<<(uint(int32(8))%32) + v226
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v234 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
				v326 = v232 + v211
				v327 = v234 + v212
				v328 = v231
			case 10:
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
				v221 = v217<<(uint(int32(24))%32) + v213
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
				v226 = v222<<(uint(int32(16))%32) + v221
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
				v231 = v227<<(uint(int32(8))%32) + v226
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
				v234 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
				v326 = v232 + v211
				v327 = v234 + v212
				v328 = v231
			default:
				v326 = v211
				v327 = v212
				v328 = v213
			}
		}
		v331 = int32(14)
		v333 = v327 ^ v328 - base.I32_rotl(v327, v331)
		v337 = v333 ^ v326 - base.I32_rotl(v333, int32(11))
		v341 = v337 ^ v327 - base.I32_rotl(v337, int32(25))
		v345 = v341 ^ v333 - base.I32_rotl(v341, int32(16))
		v349 = v345 ^ v337 - base.I32_rotl(v345, int32(4))
		v353 = v349 ^ v341 - base.I32_rotl(v349, v331)
		v363 = F_Int64GetDatum(m, base.I64_extend_i32_u(v353)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v353^v345-base.I32_rotl(v353, int32(24))))
		mBase = m.M
		v364 = m.ExcPending
		if v364 != 0 {
			return int32(0)
		} else {
			v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v365 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v368 = m.ExcPending
				if v368 != 0 {
					return int32(0)
				} else {
					return v363
				}
			} else {
				return v363
			}
		}
	}
}
func F_have_free_buffer(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_have_free_buffer[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	return int32(base.Ui32(v3^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_heapgettup_pagemode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
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
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v23 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v62 = v57
	v63 = v54
	v65 = v55
	v68 = v56
	goto L12
L2:
	;
	v54 = v46
	v55 = v52
	v56 = v44
	v57 = int32(1)
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v26 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v54 = v5
	v55 = v5
	v56 = v5
	v57 = int32(0)
	goto L1
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v46 = v45 + l1
	if l1 != int32(1) {
		v52 = v45
		goto L2
	} else {
		goto L10
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[0]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(v26^int32(-1))<<(uint(int32(2))%32))))
	v44 = v36
	goto L6
L8:
	;
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[1]))
	v44 = v38 + v26<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v52 = v49 - v46
	goto L2
L11:
	;
	m.G0 = v19 + int32(16)
	return
L12:
	;
	if v62 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v234 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v234
	v238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v238
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v238)
	goto L11
L14:
	;
	goto L13
L15:
	;
	F_heap_fetch_next_buffer(m, l0, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v65 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	return
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v78 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_heap_prepare_pagescan(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v83 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = base.I32_rotr(v102, int32(16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v107 = int32(1)
	if l1 != v107 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[0]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v83^int32(-1))<<(uint(int32(2))%32))))
	v101 = v93
	goto L22
L24:
	;
	goto L25
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[1]))
	v101 = v95 + v83<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L26:
	;
	v112 = v106 - v107
	goto L28
L27:
	;
	v112 = int32(0)
	goto L28
L28:
	;
	v62 = int32(1)
	v63 = v112
	v65 = v106
	v68 = v101
	goto L12
L29:
	;
	v125 = v65
	v126 = v63
	goto L32
L30:
	;
	v224 = v65
	goto L31
L31:
	;
	v62 = int32(0)
	v65 = v224
	goto L12
L32:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v126<<(uint(int32(1))%32)))))
	v140 = v68 + int32(20) + v137<<(uint(int32(2))%32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v68 + v141&int32(_a_F_heapgettup_pagemode_0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v137)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v146) >> (uint(int32(17)) % 32))
	v151 = int32(0)
	if base.B2i32(l3 == v151)|base.B2i32(l2 == v151) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v224 = v216
	goto L31
L34:
	;
	v216 = v125 - int32(1)
	if v216 != 0 {
		v125 = v216
		v126 = l1 + v126
		goto L32
	} else {
		goto L46
	}
L35:
	;
	v211 = v63
	goto L37
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+52))
	v162 = l3
	v164 = l2
	goto L38
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v211
	goto L11
L38:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v174&int32(1) != 0 {
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v211 = v126
	goto L37
L40:
	;
	v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v162)+4)))
	v180 = F_heap_getattr_1(m, l0-int32(-64), v177, v157, v19+int32(15))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v182 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v187 = F_FunctionCall2Coll(m, v162+int32(16), v185, v180, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	if v187 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L44:
	;
	v194 = v164 - int32(1)
	if v194 != 0 {
		v162 = v162 + int32(48)
		v164 = v194
		goto L38
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	goto L33
}
func F_hemdist_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = int32(4)
	v12 = v10 & v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v13&v11 != 0 {
		if v12 != 0 {
			return int32(0)
		} else {
			v19 = l2 << (uint(int32(3)) % 32)
			if l2 <= int32(0) {
				return v19
			} else {
				v122 = l1 + int32(8)
				v123 = v19
				v125 = v122
				v126 = int32(0)
				v129 = int32(0)
				for {
					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
					v135 = int32(1)
					v170 = v126 + v134&v135 + int32(base.Ui32(v134)>>(uint(int32(7))%32)) + int32(base.Ui32(v134)>>(uint(v135)%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(2))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(3))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(4))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(5))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(6))%32))&v135
					v174 = v129 + v135
					if v174 != l2 {
						v125 = v125 + v135
						v126 = v170
						v129 = v174
						continue
					} else {
						break
					}
					break
				}
				return v123 - v170
			}
		}
	} else {
		if v12 != 0 {
			v26 = l2 << (uint(int32(3)) % 32)
			if l2 <= int32(0) {
				return v26
			} else {
				v122 = l0 + int32(8)
				v123 = v26
				v125 = v122
				v126 = int32(0)
				v129 = int32(0)
				for {
					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
					v135 = int32(1)
					v170 = v126 + v134&v135 + int32(base.Ui32(v134)>>(uint(int32(7))%32)) + int32(base.Ui32(v134)>>(uint(v135)%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(2))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(3))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(4))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(5))%32))&v135 + int32(base.Ui32(v134)>>(uint(int32(6))%32))&v135
					v174 = v129 + v135
					if v174 != l2 {
						v125 = v125 + v135
						v126 = v170
						v129 = v174
						continue
					} else {
						break
					}
					break
				}
				return v123 - v170
			}
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v36 = int32(8)
				v37 = l1 + v36
				v39 = l0 + v36
				v40 = int32(0)
				v42 = int32(1)
				v44 = l2 << (uint(int32(3)) % 32)
				if v44 <= v42 {
					v47 = v42
				} else {
					v47 = v44
				}
				if v47 != int32(1) {
					v55 = v40
					v56 = v40
					v57 = int32(0)
					for {
						v65 = int32(base.Ui32(v56) >> (uint(int32(3)) % 32))
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v65))))
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v65))))
						v70 = v67 ^ v69
						v72 = v56 & int32(6)
						v73 = int32(1)
						v82 = int32(base.Ui32(v70)>>(uint(v72|v73)%32))&v73 + (int32(base.Ui32(v70)>>(uint(v72)%32))&v73 + v55)
						v83 = int32(2)
						v84 = v56 + v83
						v86 = v57 + v83
						if v86 != v47&int32(2147483640) {
							v55 = v82
							v56 = v84
							v57 = v86
							continue
						} else {
							break
						}
						break
					}
					if v47&int32(1) == int32(0) {
						v112 = v82
					} else {
						v90 = v82
						v91 = v84
						v100 = int32(base.Ui32(v91) >> (uint(int32(3)) % 32))
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v100))))
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v39))))
						v112 = v90 + int32(base.Ui32(v102^v104)>>(uint(v91&int32(7))%32))&int32(1)
					}
				} else {
					v90 = v40
					v91 = v40
					v100 = int32(base.Ui32(v91) >> (uint(int32(3)) % 32))
					v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v100))))
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v39))))
					v112 = v90 + int32(base.Ui32(v102^v104)>>(uint(v91&int32(7))%32))&int32(1)
				}
				return v112
			}
		}
	}
}
func F_hex_decode_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v149 int64
	_ = v149
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 == int32(0) {
		v134 = l2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v149
L2:
	;
	v149 = base.I64_extend_i32_s(v134 - l2)
	goto L1
L3:
	;
	v17 = l0 + l1
	v18 = l0
	v24 = l2
	goto L4
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v30 = v28 - int32(9)
	v37 = int32(0)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(_a_F_hex_decode_safe_0) == v37) == v37 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v134 = v124
	goto L2
L6:
	;
	v43 = v18 + int32(1)
	if base.Ui32(v43) < base.Ui32(v17) {
		v18 = v43
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(v28) <= base.Ui32(int32(126)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v134 = v24
	goto L2
L10:
	;
	v75 = v18 + int32(1)
	if base.Ui32(v17) <= base.Ui32(v75) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_hex_decode_safe[0]))))
	if int32(0) <= v47 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v51 = F_errsave_start(m, l3)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	return int64(0)
L16:
	;
	if v51 == int32(0) {
		v149 = v10
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v60 = F_pg_mblen_range(m, v18, v17)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v60
	F_errmsg(m, int32(_a_F_hex_decode_safe_1), v13+int32(16))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(239), int32(_a_F_hex_decode_safe_3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v149 = v10
	goto L1
L22:
	;
	v77 = F_errsave_start(m, l3)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if base.Ui32(v93) <= base.Ui32(int32(126)) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	if v77 == int32(0) {
		v149 = v10
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_hex_decode_safe_4), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(244), int32(_a_F_hex_decode_safe_3))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v149 = v10
	goto L1
L30:
	;
	v121 = v96 | v47<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v121)
	v124 = v24 + int32(1)
	v126 = v18 + int32(2)
	if base.Ui32(v126) < base.Ui32(v17) {
		v18 = v126
		v24 = v124
		goto L4
	} else {
		goto L41
	}
L31:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_hex_decode_safe[0]))))
	if int32(0) <= v96 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v100 = F_errsave_start(m, l3)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	if v100 == int32(0) {
		v149 = v10
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v107 = F_pg_mblen_range(m, v75, v17)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v107
	F_errmsg(m, int32(_a_F_hex_decode_safe_1), v13)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(249), int32(_a_F_hex_decode_safe_3))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v149 = v10
	goto L1
L41:
	;
	goto L5
}
func F_hex_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	return base.I64_extend_i32_u(l1) << (uint(int64(1)) % 64)
}
func F_hex_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	if l1 != 0 {
		v6 = l0
		v8 = l2
		for {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			v11 = int32(1)
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10<<(uint(v11)%32))+uint32(_c_F_hex_encode[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v13)
			v18 = v6 + v11
			if base.Ui32(v18) < base.Ui32(l0+l1) {
				v6 = v18
				v8 = v8 + int32(2)
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return base.I64_extend_i32_u(l1) << (uint(int64(1)) % 64)
}
func F_hindi_UTF_8_create_env(m *base.Module) int32 {
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v1 = int32(0)
	v3 = F_SN_create_env(m, v1, v1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_hmac_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = m.T0[v5].(func(*base.Module, int32) int32)(m, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
		m.T0[v9].(func(*base.Module, int32))(m, v8)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v6 != 0 {
				base.MemoryFill(m, v12, int32(0), v6)
			} else {
			}
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v6 != 0 {
				base.MemoryFill(m, v15, int32(0), v6)
			} else {
			}
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			F_pfree(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				F_pfree(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_hnswbuildphasename(m *base.Module, l0 int64) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	if l0 == int64(2) {
		v7 = int32(_a_F_hnswbuildphasename_0)
	} else {
		v7 = int32(0)
	}
	if l0 == int64(1) {
		v10 = int32(_a_F_hnswbuildphasename_1)
	} else {
		v10 = v7
	}
	return v10
}
func F_hnswendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+52))
	F_MemoryContextDelete(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_pfree(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
			return
		}
	}
}
func F_hnswhandler(m *base.Module, l0 int32) int32 {
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v16 = Fn13926(m, l0, int32(_a_F_hnswhandler_0), int32(_a_F_hnswhandler_1), int32(_a_F_hnswhandler_2), int32(_a_F_hnswhandler_3), int32(_a_F_hnswhandler_4), int32(_a_F_hnswhandler_5), int32(_a_F_hnswhandler_6), int32(_a_F_hnswhandler_7), int32(_a_F_hnswhandler_8), int32(_a_F_hnswhandler_9), int32(_a_F_hnswhandler_10), int32(_a_F_hnswhandler_11), int32(_a_F_hnswhandler_12), int64(72057594038124544))
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		return v16
	}
}
func F_hnswoptions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_hnswoptions[0]))
	v8 = F_build_reloptions(m, l0, l1, v4, int32(12), int32(_a_F_hnswoptions_0), int32(2))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_hyphenate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
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
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	v5 = int32(0)
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = l2
	goto L3
L2:
	;
	v18 = v5
	goto L3
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v69 = l3 + v16<<(uint(int32(3))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69-int32(380))))
	v73 = int32(1)
	v76 = int32(base.Ui32(v72+v73) >> (uint(v73) % 32))
	if v76 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v22 = l0
	v23 = l1
	v25 = int32(0)
	v26 = v16
	goto L10
L8:
	;
	v48 = l0
	v63 = int32(1)
	goto L9
L9:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v64)
	return v63
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v26)
	v38 = int32(1)
	v41 = v22 + v38
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v42 != 0 {
		v22 = v41
		v23 = v23 + v38
		v25 = v25 + v38
		v26 = v42
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v48 = v41
	v63 = v25 + int32(2)
	goto L9
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69-int32(384))))
	v86 = v83 - int32(1)
	v87 = v86 + v76
	v90 = l2 + v87<<(uint(int32(3))%32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v96 = l1
	v97 = v92
	v98 = v91
	v100 = v5
	v102 = v87
	v103 = v76
	v104 = v86
	v105 = v5
	v106 = v83 + v72
	goto L16
L16:
	;
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	if v100&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	return int32(0)
L18:
	;
	if v256 != 0 {
		v96 = v250
		v97 = v261
		v98 = v251
		v100 = v253
		v102 = v255
		v103 = v256
		v104 = v257
		v105 = v258
		v106 = v259
		goto L16
	} else {
		goto L54
	}
L19:
	;
	v147 = v100 | base.B2i32(base.I32_extend8_s(v143) < v108)
	v148 = int32(1)
	v151 = base.B2i32(v108 < v144) | v105
	if v147&v148&(v151&v148) != 0 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	v129 = (v126 | v100) & int32(1)
	if v129 != 0 {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	if v108 < v114 {
		v126 = int32(0)
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v105&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v143 = v119
	v144 = v118
	goto L19
L26:
	;
	goto L27
L27:
	;
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	if v108 <= v121 {
		v143 = v120
		v144 = v121
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v126 = base.B2i32(v120 <= v108)
	goto L20
L29:
	;
	v130 = v106
	goto L31
L30:
	;
	v130 = v102
	goto L31
L31:
	;
	if v129 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v131 = v102
	goto L34
L33:
	;
	v131 = v104
	goto L34
L34:
	;
	v134 = int32(base.Ui32(v130-v131) >> (uint(int32(1)) % 32))
	v135 = v134 + v131
	v138 = l2 + v135<<(uint(int32(3))%32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v250 = l1
	v251 = v139
	v253 = v140
	v255 = v135
	v256 = v134
	v257 = v131
	v258 = v140
	v259 = v130
	v261 = v142
	goto L18
L35:
	;
	v246 = int32(2)
	v250 = v164
	v251 = v98 + v246
	v253 = v147
	v255 = v102
	v256 = v103
	v257 = v104
	v258 = v151
	v259 = v106
	v261 = v97 + v246
	goto L18
L36:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l2+v102<<(uint(int32(3))%32))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v183 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v155 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v159 = v98 + int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v160 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v164 = v96 + int32(1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v165 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v155-int32(48))&int32(255)) {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	v250 = v164
	v251 = v159
	v253 = v147
	v255 = v102
	v256 = v103
	v257 = v104
	v258 = v151
	v259 = v106
	v261 = v97 + int32(1)
	goto L18
L42:
	;
	v185 = l0
	v186 = l1
	v188 = v183
	v189 = v182
	v192 = int32(0)
	goto L45
L43:
	;
	v225 = l0
	v226 = l1
	v240 = int32(1)
	goto L44
L44:
	;
	v241 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v241)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)) = uint8(v243)
	return v240
L45:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v200 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v225 = v216
	v226 = v217
	v240 = v221 + int32(1)
	goto L44
L47:
	;
	v201 = int32(45)
	v205 = base.B2i32(v188&int32(255) != v201)
	if v188&int32(255) != v201 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v216 = v185
	v217 = v186
	v221 = v192
	goto L49
L49:
	;
	goto L46
L50:
	;
	v206 = v200
	goto L52
L51:
	;
	v206 = v201
	goto L52
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v206)
	v208 = int32(1)
	v209 = v192 + v208
	v211 = v185 + v208
	v212 = v186 + v205
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	if v213 != 0 {
		v185 = v211
		v186 = v212
		v188 = v213
		v189 = v189 + v208
		v192 = v209
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v216 = v211
	v217 = v212
	v221 = v209
	goto L49
L54:
	;
	goto L17
}
func F_hypothetical_percent_rank_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_hypothetical_rank_common(m, l0, int32(-1), v7+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		if v16 == int64(0) {
			v25 = float64(0)
		} else {
			v25 = base.F64_div(base.F64_convert_i64_s(v12-int64(1)), base.F64_convert_i64_s(v16))
		}
		v26 = F_Float8GetDatum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v26
		}
	}
}
func F_hypothetical_rank_final(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_hypothetical_rank_common(m, l0, int32(-1), v5+int32(8))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Int64GetDatum(m, v10)
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v14
		}
	}
}
