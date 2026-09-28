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
	v48 = v30 + v45*int32(768)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+336))
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
	v54 = int32(0)
	if base.B2i32(v53 == v54)|base.B2i32(l1 <= v54) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
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
func F_handle_pm_shutdown_request_signal(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
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
	v17 = int32(0)
	v20 = base.AtomicRmwOr32(m, v17, int32(_a_F_handle_pm_shutdown_request_signal_1), v17)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v21 != 0 {
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
	v7 = int32(_a_F_handle_pm_shutdown_request_signal_2)
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
	v24 = int32(0)
	v27 = base.AtomicRmwOr32(m, v24, int32(_a_F_handle_pm_shutdown_request_signal_1), v24)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v28 == v24 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v31 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[2]))
	if v35 == v31 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[3]))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v65 = F_pgmem_kill(m, v31, int32(23))
	mBase = m.M
	goto L6
L13:
	;
	m.G0 = v39 + int32(16)
	goto L5
L14:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+15)) = uint8(v45)
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[4]))
	v53 = F_write(m, v49, v39+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v53 {
		goto L13
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[5]))
	if v57 == int32(27) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v12 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(l0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 == int32(0) {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v19)+74)))
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = v21
						return v25 & int32(1)
					}
				}
			}
		} else {
			v25 = int32(1)
			return v25 & int32(1)
		}
	}
}
func F_has_superclass(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14309(m, l0, int32(2680), int32(1), int32(2611))
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
	Fn14310(m, l0, l1, l2, l3, int32(405))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_hashbool(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
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
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = int32(711645284)
	v12 = base.B2i32(v2 != int64(0)) - int32(1636608428) ^ v9 - int32(1455628627)
	v17 = v12 ^ int32(-1636608428) - base.I32_rotl(v12, int32(25))
	v22 = v17 ^ v9 - base.I32_rotl(v17, int32(16))
	v26 = v22 ^ v12 - base.I32_rotl(v22, int32(4))
	v30 = v26 ^ v17 - base.I32_rotl(v26, int32(14))
	return base.I64_extend_i32_u(v30 ^ v22 - base.I32_rotl(v30, int32(24)))
}
func F_hashbucketcleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v122 int32
	_ = v122
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 float64
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v185 float64
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 float64
	_ = v209
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v400 int32
	_ = v400
	var v401 int64
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
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v835 int64
	_ = v835
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v838 int32
	_ = v838
	var v870 int64
	_ = v870
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v890 int64
	_ = v890
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v981 int32
	_ = v981
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1445 int32
	_ = v1445
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1540 int32
	_ = v1540
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1553 int64
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int64
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1588 int64
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1613 int32
	_ = v1613
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1627 int32
	_ = v1627
	var v1629 int64
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1639 int32
	_ = v1639
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1653 int32
	_ = v1653
	var v1658 int32
	_ = v1658
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1672 int32
	_ = v1672
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1759 int32
	_ = v1759
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1873 int32
	_ = v1873
	var v1885 int32
	_ = v1885
	var v1892 int32
	_ = v1892
	var v1921 int32
	_ = v1921
	v32 = m.G0
	v34 = v32 - int32(_a_F_hashbucketcleanup_0)
	m.G0 = v34
	v54 = l2
	v62 = l3
	v65 = int32(0)
	goto L1
L1:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v70 = int32(0)
	v71 = base.B2i32(v70 <= v54)
	if v71 == v70 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+16)))
	v91 = v90 + v89
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v92) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75+(v54^int32(-1))<<(uint(int32(2))%32))))
	v89 = v81
	goto L5
L7:
	;
	goto L8
L8:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v89 = v83 + v54<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L9:
	;
	if v319 != int32(-1) {
		goto L67
	} else {
		goto L68
	}
L10:
	;
	v122 = int32(1)
	v136 = int32(0)
	goto L15
L11:
	;
	v96 = v92 + int32(_a_F_hashbucketcleanup_1)
	if v96&int32(_a_F_hashbucketcleanup_2) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v319 = v100
	v329 = v65
	goto L9
L14:
	;
	goto L13
L15:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(20)+v122<<(uint(int32(2))%32))))
	v146 = v89 + v143&int32(_a_F_hashbucketcleanup_3)
	if l11 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v189 <= int32(0) {
		v319 = v193
		v329 = v65
		goto L9
	} else {
		goto L37
	}
L17:
	;
	if v122 != int32(base.Ui32(v96)>>(uint(int32(2))%32))&int32(_a_F_hashbucketcleanup_4) {
		v122 = v122 + int32(1)
		v136 = v189
		goto L15
	} else {
		goto L36
	}
L18:
	;
	if l9 == int32(0) {
		v189 = v136
		goto L17
	} else {
		goto L35
	}
L19:
	;
	v177 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(16)+v136<<(uint(v177)%32)))) = uint16(v122)
	v189 = v136 + v177
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
	v149 = m.T0[l11].(func(*base.Module, int32, int32) int32)(m, v146, l12)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v149 == int32(0) {
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
	v155 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v155, float64(1))
	goto L19
L25:
	;
	v163 = int32(*(*int16)(unsafe.Add(mBase, uint32(v146)+6)))
	if int32(0) <= v163 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v170 = v168 & l6
	if base.Ui32(v170) <= base.Ui32(l5) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v166 = int32(8)
	goto L29
L28:
	;
	v166 = int32(16)
	goto L29
L29:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v146+v166)))
	goto L26
L30:
	;
	if v172&v170 == l1 {
		goto L18
	} else {
		goto L34
	}
L31:
	;
	v172 = int32(-1)
	goto L33
L32:
	;
	v172 = l7
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L19
L35:
	;
	v185 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v185, float64(1))
	v189 = v136
	goto L17
L36:
	;
	goto L16
L37:
	;
	v196 = int32(0)
	v197 = int32(_a_F_hashbucketcleanup_5)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v199 + int32(1)
	F_PageIndexMultiDelete(m, v89, v34+int32(16), v189)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if l8 == int32(0) {
		v223 = v196
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_MarkBufferDirty(m, v54)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L43
	}
L40:
	;
	v209 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	if base.F64_gt(v209, float64(0)) == int32(0) {
		v223 = v196
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+12)))
	if v214&int32(128) == int32(0) {
		v223 = v196
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v220 = v214 & int32(_a_F_hashbucketcleanup_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+12)) = uint16(v220)
	v223 = int32(1)
	goto L39
L43:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+118)))
	if v228 != int32(112) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v71 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L45:
	;
	v269 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L3
	} else {
		goto L61
	}
L46:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v232 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v235 != 0 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+14)) = uint8(v223)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+15)) = uint8(base.B2i32(l2 == v54))
	F_XLogBeginInsert(m)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v236 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	F_XLogRegisterData(m, v34+int32(14), int32(2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+15)))
	if v247 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(42))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	F_XLogRegisterBuffer(m, int32(1), v54, int32(8))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v258 = int32(1)
	F_XLogRegisterBufData(m, v258, v34+int32(16), v189<<(uint(v258)%32))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v267 = F_XLogInsert(m, int32(12), int32(144))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	v271 = v267
	goto L44
L61:
	;
	v271 = v269
	goto L44
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = base.I64_rotl(v271, int64(32))
	v293 = int32(_a_F_hashbucketcleanup_5)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	v296 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v295 - v296
	v319 = v193
	v329 = v296
	goto L9
L63:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275+(v54^int32(-1))<<(uint(int32(2))%32))))
	v289 = v281
	goto L62
L64:
	;
	goto L65
L65:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v289 = v283 + v54<<(uint(int32(13))%32) + int32(-8192)
	goto L62
L66:
	;
	v54 = v334
	v62 = v319
	v65 = v329
	goto L1
L67:
	;
	v334 = F__hash_getbuf_with_strategy(m, l0, v319, int32(1), l4)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L3
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if l2 != v54 {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	if l3 == v62 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_UnlockBuffer(m, v54)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_UnlockReleaseBuffer(m, v54)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L75
	}
L74:
	;
	goto L66
L75:
	;
	goto L66
L76:
	;
	F_UnlockReleaseBuffer(m, v54)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if l10 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	F_LockBufferInternal(m, l2, int32(3))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	if l2 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	goto L83
L83:
	;
	if v329 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L84:
	;
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364)+16)))
	v366 = int32(_a_F_hashbucketcleanup_5)
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v368 + int32(1)
	v372 = v365 + v364
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372)+12)))
	v375 = v373 & int32(_a_F_hashbucketcleanup_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v372)+12)) = uint16(v375)
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L3
	} else {
		goto L88
	}
L85:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v350+(l2^int32(-1))<<(uint(int32(2))%32))))
	v364 = v356
	goto L84
L86:
	;
	goto L87
L87:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v364 = v358 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L84
L88:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+118)))
	if v380 != int32(112) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v364))) = base.I64_rotl(v401, int64(32))
	v405 = int32(_a_F_hashbucketcleanup_5)
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v407 - int32(1)
	goto L83
L90:
	;
	v399 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L3
	} else {
		goto L100
	}
L91:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v384 <= int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v387 != 0 {
		goto L90
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L97
	}
L95:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v388 != 0 {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v397 = F_XLogInsert(m, int32(12), int32(160))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v401 = v397
	goto L89
L100:
	;
	v401 = v399
	goto L89
L101:
	;
	m.G0 = v1921 + int32(_a_F_hashbucketcleanup_0)
	return
L102:
	;
	F_UnlockBuffer(m, l2)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L3
	} else {
		goto L421
	}
L103:
	;
	v415 = F_IsBufferCleanupOK(m, l2)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	if v415 == int32(0) {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v419 = int32(0)
	v420 = m.G0
	v422 = v420 + int32(-8192)
	m.G0 = v422
	if l2 < v419 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	m.G0 = v1873 - int32(-8192)
	v1921 = v1885
	goto L101
L107:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441)+16)))
	v443 = v442 + v441
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v444 != int32(-1) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v427+(l2^int32(-1))<<(uint(int32(2))%32))))
	v441 = v433
	goto L107
L109:
	;
	goto L110
L110:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v441 = v435 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L107
L111:
	;
	v453 = v444
	v459 = v419
	goto L114
L112:
	;
	goto L113
L113:
	;
	F_UnlockBuffer(m, l2)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L3
	} else {
		goto L420
	}
L114:
	;
	if v459 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v506 = l2
	v507 = l3
	v510 = l4
	v511 = l3
	v513 = l2
	v514 = v453
	v515 = l0
	v518 = v481
	v522 = v422
	v524 = v441
	v525 = v500
	v526 = v502
	v533 = v443
	v534 = v34
	goto L126
L116:
	;
	F_UnlockReleaseBuffer(m, v459)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v481 = F__hash_getbuf_with_strategy(m, l0, v453, int32(1), l4)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L3
	} else {
		goto L121
	}
L119:
	;
	goto L118
L120:
	;
	v501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v500)+16)))
	v502 = v501 + v500
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	if v503 != int32(-1) {
		v453 = v503
		v459 = v481
		goto L114
	} else {
		goto L125
	}
L121:
	;
	if v481 < int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486+(v481^int32(-1))<<(uint(int32(2))%32))))
	v500 = v492
	goto L120
L123:
	;
	goto L124
L124:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v500 = v494 + v481<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L125:
	;
	goto L115
L126:
	;
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v525)+12)))
	if base.Ui32(v537) < base.Ui32(int32(25)) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	F_UnlockReleaseBuffer(m, v1088)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L3
	} else {
		goto L419
	}
L128:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v1121 = v522 + int32(2464)
	v1123 = v522 + int32(16)
	v1127 = v1101 & int32(_a_F_hashbucketcleanup_4)
	v1129 = m.G0
	v1131 = v1129 - int32(32)
	m.G0 = v1131
	F__hash_checkpage(m, v515, v518, int32(1))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L3
	} else {
		goto L227
	}
L129:
	;
	v1088 = v506
	v1089 = v507
	v1101 = int32(0)
	v1106 = v524
	v1115 = v533
	goto L128
L130:
	;
	goto L131
L131:
	;
	v544 = int32(base.Ui32(v537+int32(_a_F_hashbucketcleanup_1)) >> (uint(int32(2)) % 32))
	if v544&int32(_a_F_hashbucketcleanup_4) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v1088 = v506
	v1089 = v507
	v1101 = int32(0)
	v1106 = v524
	v1115 = v533
	goto L128
L133:
	;
	goto L134
L134:
	;
	v560 = v506
	v561 = v507
	v564 = v544
	v576 = v524
	v585 = v533
	goto L135
L135:
	;
	v590 = int32(0)
	v595 = v560
	v596 = v561
	v601 = v590
	v608 = v590
	v613 = v576
	v617 = int32(1)
	v621 = v590
	v622 = v585
	goto L137
L136:
	;
	v1088 = v698
	v1089 = v693
	v1101 = v970
	v1106 = v967
	v1115 = v969
	goto L128
L137:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v525+int32(20)+v617&int32(_a_F_hashbucketcleanup_4)<<(uint(int32(2))%32))))
	v632 = int32(_a_F_hashbucketcleanup_8)
	if v631&v632 != v632 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v1077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v525)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1077) {
		goto L223
	} else {
		goto L224
	}
L139:
	;
	goto L138
L140:
	;
	v638 = v525 + v631&int32(_a_F_hashbucketcleanup_3)
	v639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v638)+6)))
	v645 = (v639&int32(_a_F_hashbucketcleanup_9) + int32(7)) & int32(_a_F_hashbucketcleanup_10)
	v646 = v595
	v647 = v596
	v652 = v601
	v659 = v608
	v664 = v613
	v672 = v621
	v673 = v622
	goto L144
L141:
	;
	v1041 = v595
	v1042 = v596
	v1047 = v601
	v1054 = v608
	v1059 = v613
	v1067 = v621
	v1068 = v622
	goto L142
L142:
	;
	v1073 = v617 + int32(1)
	if base.Ui32(v1073&int32(_a_F_hashbucketcleanup_4)) <= base.Ui32(v564&int32(_a_F_hashbucketcleanup_4)) {
		v595 = v1041
		v596 = v1042
		v601 = v1047
		v608 = v1054
		v613 = v1059
		v617 = v1073
		v621 = v1067
		v622 = v1068
		goto L137
	} else {
		goto L222
	}
L143:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v522+int32(_a_F_hashbucketcleanup_11)+v652&int32(_a_F_hashbucketcleanup_4)<<(uint(int32(1))%32)))) = uint16(v617)
	v1025 = F_CopyIndexTuple(m, v638)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L3
	} else {
		goto L221
	}
L144:
	;
	v678 = v659 & int32(_a_F_hashbucketcleanup_4)
	v681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+14)))
	v682 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+12)))
	v683 = v681 - v682
	v685 = (v678 + int32(1)) << (uint(int32(2)) % 32)
	if v685 <= v683 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v981 = v970
	goto L217
L146:
	;
	v690 = v645 + v672
	if base.Ui32(v690) <= base.Ui32(v689) {
		goto L143
	} else {
		goto L150
	}
L147:
	;
	v689 = v683 - v685
	goto L149
L148:
	;
	v689 = int32(0)
	goto L149
L149:
	;
	goto L146
L150:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	if v514 != v693 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v696 = F__hash_getbuf_with_strategy(m, v515, v693, int32(1), v510)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L3
	} else {
		goto L154
	}
L152:
	;
	v698 = int32(0)
	goto L153
L153:
	;
	if v678 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v698 = v696
	goto L153
L155:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v515)+48))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+118)))
	if v700 != int32(112) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	if v647 == v511 {
		goto L203
	} else {
		goto L204
	}
L158:
	;
	v714 = int32(_a_F_hashbucketcleanup_5)
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v716 + int32(1)
	F__hash_pgaddmultitup(m, v515, v646, v522+int32(2464), v522+int32(16), v678)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L3
	} else {
		goto L166
	}
L159:
	;
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v704 <= int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v515)+32))
	if v707 != 0 {
		goto L158
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_XLogEnsureRecordSpace(m, int32(0), v678+int32(3))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L3
	} else {
		goto L165
	}
L163:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v515)+40))
	if v708 != 0 {
		goto L158
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	goto L158
L166:
	;
	F_MarkBufferDirty(m, v646)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	v731 = v652 & int32(_a_F_hashbucketcleanup_4)
	F_PageIndexMultiDelete(m, v525, v522+int32(_a_F_hashbucketcleanup_11), v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	F_MarkBufferDirty(m, v518)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v515)+48))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+118)))
	if v737 != int32(112) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v646 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L171:
	;
	v837 = F_XLogGetFakeLSN(m, v515)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L3
	} else {
		goto L193
	}
L172:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v741 <= int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v515)+32))
	if v744 != 0 {
		goto L171
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v522)+12)) = uint16(v659)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+14)) = uint8(base.B2i32(v646 == v513))
	F_XLogBeginInsert(m)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L3
	} else {
		goto L178
	}
L176:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v515)+40))
	if v745 != 0 {
		goto L171
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	F_XLogRegisterData(m, v522+int32(12), int32(3))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L3
	} else {
		goto L179
	}
L179:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+14)))
	if v756 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	F_XLogRegisterBuffer(m, int32(0), v513, int32(42))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L3
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	F_XLogRegisterBuffer(m, int32(1), v646, int32(8))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L3
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	v767 = int32(1)
	F_XLogRegisterBufData(m, v767, v522+int32(16), v678<<(uint(v767)%32))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L3
	} else {
		goto L185
	}
L185:
	;
	v781 = int32(0)
	goto L186
L186:
	;
	v808 = v781 << (uint(int32(2)) % 32)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808+(v522+int32(2464)))))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v522+int32(832)+v808)))
	F_XLogRegisterBufData(m, int32(1), v812, v816)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L3
	} else {
		goto L188
	}
L187:
	;
	F_XLogRegisterBuffer(m, int32(2), v518, int32(8))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L3
	} else {
		goto L190
	}
L188:
	;
	v820 = v781 + int32(1)
	if v820 != v678 {
		v781 = v820
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	F_XLogRegisterBufData(m, int32(2), v522+int32(_a_F_hashbucketcleanup_11), v731<<(uint(int32(1))%32))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L3
	} else {
		goto L191
	}
L191:
	;
	v835 = F_XLogInsert(m, int32(12), int32(112))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	v870 = v835
	goto L170
L193:
	;
	v870 = v837
	goto L170
L194:
	;
	v890 = base.I64_rotl(v870, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v888))) = v890
	if v518 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v874+(v646^int32(-1))<<(uint(int32(2))%32))))
	v888 = v880
	goto L194
L196:
	;
	goto L197
L197:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v888 = v882 + v646<<(uint(int32(13))%32) + int32(-8192)
	goto L194
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v903))) = v890
	v905 = int32(_a_F_hashbucketcleanup_5)
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v907 - int32(1)
	goto L157
L199:
	;
	v895 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v895+(v518^int32(-1))<<(uint(int32(2))%32))))
	v903 = v897
	goto L198
L200:
	;
	goto L201
L201:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v903 = v899 + v518<<(uint(int32(13))%32) + int32(-8192)
	goto L198
L202:
	;
	if v693 == v514 {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	F_UnlockBuffer(m, v646)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L3
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_UnlockReleaseBuffer(m, v646)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L3
	} else {
		goto L207
	}
L206:
	;
	goto L202
L207:
	;
	goto L202
L208:
	;
	F_UnlockReleaseBuffer(m, v518)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L3
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if v698 < int32(0) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1873 = v522
	v1885 = v534
	goto L106
L212:
	;
	v968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v967)+16)))
	v969 = v968 + v967
	v970 = int32(0)
	if v678 == v970 {
		v646 = v698
		v647 = v693
		v652 = v970
		v659 = v970
		v664 = v967
		v672 = v970
		v673 = v969
		goto L144
	} else {
		goto L216
	}
L213:
	;
	v953 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v953+(v698^int32(-1))<<(uint(int32(2))%32))))
	v967 = v959
	goto L212
L214:
	;
	goto L215
L215:
	;
	v961 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v967 = v961 + v698<<(uint(int32(13))%32) + int32(-8192)
	goto L212
L216:
	;
	goto L145
L217:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v522+int32(2464)+v981<<(uint(int32(2))%32))))
	F_pfree(m, v1011)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L3
	} else {
		goto L219
	}
L218:
	;
	goto L139
L219:
	;
	v1015 = v981 + int32(1)
	if v1015 != v678 {
		v981 = v1015
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v1028 = v678 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1028+(v522+int32(832))))) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v522+int32(2464)+v1028))) = v1025
	v1037 = int32(1)
	v1041 = v646
	v1042 = v647
	v1047 = v652 + v1037
	v1054 = v659 + v1037
	v1059 = v664
	v1067 = v690
	v1068 = v673
	goto L142
L222:
	;
	v1088 = v1041
	v1089 = v1042
	v1101 = v1054
	v1106 = v1059
	v1115 = v1068
	goto L128
L223:
	;
	v1085 = int32(base.Ui32(v1077+int32(_a_F_hashbucketcleanup_1)) >> (uint(int32(2)) % 32))
	goto L225
L224:
	;
	v1085 = int32(0)
	goto L225
L225:
	;
	if v1085&int32(_a_F_hashbucketcleanup_4) != 0 {
		v560 = v698
		v561 = v693
		v564 = v1085
		v576 = v967
		v585 = v969
		goto L135
	} else {
		goto L226
	}
L226:
	;
	goto L136
L227:
	;
	if v518 < int32(0) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	if v518 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L229:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[4]))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1139+(v518^int32(-1))*int32(56))+16))
	v1154 = v1145
	goto L228
L230:
	;
	goto L231
L231:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[5]))
	v1148 = int32(56)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1147+v518*v1148-v1148)+16))
	v1154 = v1153
	goto L228
L232:
	;
	v1173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1172)+16)))
	v1174 = v1173 + v1172
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1174)))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1088 < int32(0) {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1158+(v518^int32(-1))<<(uint(int32(2))%32))))
	v1172 = v1164
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1172 = v1166 + v518<<(uint(int32(13))%32) + int32(-8192)
	goto L232
L236:
	;
	if v1175 == int32(-1) {
		v1203 = int32(0)
		goto L240
	} else {
		goto L241
	}
L237:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[4]))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1180+(v1088^int32(-1))*int32(56))+16))
	v1195 = v1186
	goto L236
L238:
	;
	goto L239
L239:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[5]))
	v1189 = int32(56)
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1188+v1088*v1189-v1189)+16))
	v1195 = v1194
	goto L236
L240:
	;
	if v1176 != int32(-1) {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	if v1175 == v1195 {
		v1203 = v1088
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1201 = F__hash_getbuf_with_strategy(m, v515, v1175, int32(3), v510)
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L3
	} else {
		goto L243
	}
L243:
	;
	v1203 = v1201
	goto L240
L244:
	;
	v1207 = F__hash_getbuf_with_strategy(m, v515, v1176, int32(1), v510)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L3
	} else {
		goto L247
	}
L245:
	;
	v1209 = int32(0)
	goto L246
L246:
	;
	v1213 = F__hash_getbuf(m, v515, int32(0), int32(1), int32(8))
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L3
	} else {
		goto L249
	}
L247:
	;
	v1209 = v1207
	goto L246
L248:
	;
	v1235 = F__hash_ovflblkno_to_bitno(m, v1232+int32(24), v1154)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L3
	} else {
		goto L253
	}
L249:
	;
	if v1213 < int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1218+(v1213^int32(-1))<<(uint(int32(2))%32))))
	v1232 = v1224
	goto L248
L251:
	;
	goto L252
L252:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1232 = v1226 + v1213<<(uint(int32(13))%32) + int32(-8192)
	goto L248
L253:
	;
	v1237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1232)+46)))
	v1238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1232)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v1131)+28)) = v1235 & (v1238<<(uint(int32(3))%32) - int32(1))
	v1245 = int32(base.Ui32(v1235) >> (uint(v1237) % 32))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+68))
	if base.Ui32(v1245) < base.Ui32(v1246) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	if v1127 != 0 {
		goto L401
	} else {
		goto L402
	}
L255:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1232+v1245<<(uint(int32(2))%32))+468))
	F_UnlockBuffer(m, v1213)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L3
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L3
	} else {
		goto L398
	}
L258:
	;
	v1256 = F__hash_getbuf(m, v515, v1251, int32(3), int32(4))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L3
	} else {
		goto L260
	}
L259:
	;
	F_LockBufferInternal(m, v1213, int32(3))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L3
	} else {
		goto L264
	}
L260:
	;
	if v1256 < int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1261+(v1256^int32(-1))<<(uint(int32(2))%32))))
	v1275 = v1267
	goto L259
L262:
	;
	goto L263
L263:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1275 = v1269 + v1256<<(uint(int32(13))%32) + int32(-8192)
	goto L259
L264:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v515)+48))
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279)+118)))
	if v1280 != int32(112) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1294 = int32(_a_F_hashbucketcleanup_5)
	v1296 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v1296 + int32(1)
	if v1127 != 0 {
		goto L273
	} else {
		goto L274
	}
L266:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v1284 <= int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v515)+32))
	if v1287 != 0 {
		goto L265
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_XLogEnsureRecordSpace(m, int32(6), v1127+int32(4))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L3
	} else {
		goto L272
	}
L270:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v515)+40))
	if v1288 != 0 {
		goto L265
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	goto L265
L273:
	;
	F__hash_pgaddmultitup(m, v515, v1088, v1121, v1123, v1127)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L3
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	F_PageInit(m, v1172, int32(_a_F_hashbucketcleanup_12), int32(16))
	mBase = m.M
	goto L278
L276:
	;
	F_MarkBufferDirty(m, v1088)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L3
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v1307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1172)+16)))
	v1308 = v1172 + v1307
	*(*int64)(unsafe.Add(mBase, uint32(v1308)+8)) = int64(-36028792723996673)
	*(*int64)(unsafe.Add(mBase, uint32(v1308))) = int64(-1)
	F_MarkBufferDirty(m, v518)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L3
	} else {
		goto L279
	}
L279:
	;
	if v1203 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	if v1203 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L281:
	;
	goto L282
L282:
	;
	if v1209 != 0 {
		goto L288
	} else {
		goto L289
	}
L283:
	;
	v1333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1332)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1333+v1332)+4)) = v1176
	F_MarkBufferDirty(m, v1203)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L3
	} else {
		goto L287
	}
L284:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1318+(v1203^int32(-1))<<(uint(int32(2))%32))))
	v1332 = v1324
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1332 = v1326 + v1203<<(uint(int32(13))%32) + int32(-8192)
	goto L283
L287:
	;
	goto L282
L288:
	;
	if v1209 < int32(0) {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	goto L290
L290:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+28))
	v1367 = base.I32_div_s(v1365, int32(32))
	v1370 = v1275 + int32(24) + v1367<<(uint(int32(2))%32)
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1370)))
	*(*int32)(unsafe.Add(mBase, uint32(v1370))) = v1371 & base.I32_rotl(int32(-2), v1365)
	F_MarkBufferDirty(m, v1256)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L3
	} else {
		goto L296
	}
L291:
	;
	v1357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1356)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1357+v1356))) = v1175
	F_MarkBufferDirty(m, v1209)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L3
	} else {
		goto L295
	}
L292:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1342+(v1209^int32(-1))<<(uint(int32(2))%32))))
	v1356 = v1348
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1356 = v1350 + v1209<<(uint(int32(13))%32) + int32(-8192)
	goto L291
L295:
	;
	goto L290
L296:
	;
	v1379 = v1232 - int32(-64)
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+64))
	v1381 = base.B2i32(base.Ui32(v1380) <= base.Ui32(v1235))
	if v1381 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1379))) = v1235
	F_MarkBufferDirty(m, v1213)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L3
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1387 = base.B2i32(v1088 == v1203)
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v515)+48))
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391)+118)))
	if v1392 != int32(112) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	goto L299
L301:
	;
	if v1387|base.B2i32(v1127 != int32(0)) != 0 {
		goto L348
	} else {
		goto L349
	}
L302:
	;
	v1555 = F_XLogGetFakeLSN(m, v515)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L3
	} else {
		goto L347
	}
L303:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v1396 <= int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v515)+32))
	if v1399 != 0 {
		goto L302
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1131)+27)) = uint8(v1387)
	v1402 = base.B2i32(v1088 == v513)
	*(*uint8)(unsafe.Add(mBase, uint32(v1131)+26)) = uint8(v1402)
	*(*uint16)(unsafe.Add(mBase, uint32(v1131)+24)) = uint16(v1127)
	*(*int32)(unsafe.Add(mBase, uint32(v1131)+20)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v1131)+16)) = v1175
	F_XLogBeginInsert(m)
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L3
	} else {
		goto L309
	}
L307:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v515)+40))
	if v1400 != 0 {
		goto L302
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	F_XLogRegisterData(m, v1131+int32(16), int32(12))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L3
	} else {
		goto L310
	}
L310:
	;
	if v1402 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	F_XLogRegisterBuffer(m, int32(0), v513, int32(42))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L3
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	if v1127 != 0 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	goto L313
L315:
	;
	F_XLogRegisterBuffer(m, int32(2), v518, int32(8))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L3
	} else {
		goto L330
	}
L316:
	;
	F_XLogRegisterBuffer(m, int32(1), v1088, int32(8))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L3
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	if base.B2i32(v1402 == int32(0))&base.B2i32(v1088 != v1203) != 0 {
		goto L315
	} else {
		goto L325
	}
L319:
	;
	v1424 = int32(1)
	F_XLogRegisterBufData(m, v1424, v1123, v1127<<(uint(v1424)%32))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	v1445 = int32(0)
	goto L321
L321:
	;
	v1463 = v1445 << (uint(int32(2)) % 32)
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1121+v1463)))
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1463+(v522+int32(832)))))
	F_XLogRegisterBufData(m, int32(1), v1465, v1467)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L3
	} else {
		goto L323
	}
L322:
	;
	goto L315
L323:
	;
	v1471 = v1445 + int32(1)
	if v1471 != v1127 {
		v1445 = v1471
		goto L321
	} else {
		goto L324
	}
L324:
	;
	goto L322
L325:
	;
	if v1088 == v1203 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1480 = int32(8)
	goto L328
L327:
	;
	v1480 = int32(40)
	goto L328
L328:
	;
	F_XLogRegisterBuffer(m, int32(1), v1088, v1480)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L3
	} else {
		goto L329
	}
L329:
	;
	goto L315
L330:
	;
	v1518 = int32(0)
	if base.B2i32(v1203 == v1518)|v1387 == v1518 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	F_XLogRegisterBuffer(m, int32(3), v1203, int32(8))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L3
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	if v1209 != 0 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	goto L333
L335:
	;
	F_XLogRegisterBuffer(m, int32(4), v1209, int32(8))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L3
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	F_XLogRegisterBuffer(m, int32(5), v1256, int32(8))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L3
	} else {
		goto L339
	}
L338:
	;
	goto L337
L339:
	;
	F_XLogRegisterBufData(m, int32(5), v1131+int32(28), int32(4))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L3
	} else {
		goto L340
	}
L340:
	;
	if v1381 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	F_XLogRegisterBuffer(m, int32(6), v1213, int32(8))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L3
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1553 = F_XLogInsert(m, int32(12), int32(128))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L3
	} else {
		goto L346
	}
L344:
	;
	F_XLogRegisterBufData(m, int32(6), v1379, int32(4))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L3
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	v1588 = v1553
	goto L301
L347:
	;
	v1588 = v1555
	goto L301
L348:
	;
	if v1088 < int32(0) {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	goto L350
L350:
	;
	if v518 < int32(0) {
		goto L356
	} else {
		goto L357
	}
L351:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1606))) = base.I64_rotl(v1588, int64(32))
	goto L350
L352:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1592+(v1088^int32(-1))<<(uint(int32(2))%32))))
	v1606 = v1598
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1606 = v1600 + v1088<<(uint(int32(13))%32) + int32(-8192)
	goto L351
L355:
	;
	v1629 = base.I64_rotl(v1588, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v1627))) = v1629
	v1631 = int32(0)
	if base.B2i32(v1203 == v1631)|v1387 == v1631 {
		goto L359
	} else {
		goto L360
	}
L356:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1613+(v518^int32(-1))<<(uint(int32(2))%32))))
	v1627 = v1619
	goto L355
L357:
	;
	goto L358
L358:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1627 = v1621 + v518<<(uint(int32(13))%32) + int32(-8192)
	goto L355
L359:
	;
	if v1203 < int32(0) {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	goto L361
L361:
	;
	if v1209 != 0 {
		goto L366
	} else {
		goto L367
	}
L362:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1653))) = v1629
	goto L361
L363:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1639+(v1203^int32(-1))<<(uint(int32(2))%32))))
	v1653 = v1645
	goto L362
L364:
	;
	goto L365
L365:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1653 = v1647 + v1203<<(uint(int32(13))%32) + int32(-8192)
	goto L362
L366:
	;
	if v1209 < int32(0) {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	goto L368
L368:
	;
	if v1256 < int32(0) {
		goto L374
	} else {
		goto L375
	}
L369:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1672))) = v1629
	goto L368
L370:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1658+(v1209^int32(-1))<<(uint(int32(2))%32))))
	v1672 = v1664
	goto L369
L371:
	;
	goto L372
L372:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1672 = v1666 + v1209<<(uint(int32(13))%32) + int32(-8192)
	goto L369
L373:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1691))) = v1629
	if v1381 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L374:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1677+(v1256^int32(-1))<<(uint(int32(2))%32))))
	v1691 = v1683
	goto L373
L375:
	;
	goto L376
L376:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1691 = v1685 + v1256<<(uint(int32(13))%32) + int32(-8192)
	goto L373
L377:
	;
	if v1213 < int32(0) {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	goto L379
L379:
	;
	v1714 = int32(_a_F_hashbucketcleanup_5)
	v1716 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v1716 - int32(1)
	v1720 = int32(0)
	if base.B2i32(v1203 == v1720)|base.B2i32(v1175 == v1195) == v1720 {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1712))) = v1629
	goto L379
L381:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1698+(v1213^int32(-1))<<(uint(int32(2))%32))))
	v1712 = v1704
	goto L380
L382:
	;
	goto L383
L383:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1712 = v1706 + v1213<<(uint(int32(13))%32) + int32(-8192)
	goto L380
L384:
	;
	F_UnlockReleaseBuffer(m, v1203)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L3
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	if v518 != 0 {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	goto L386
L388:
	;
	F_UnlockReleaseBuffer(m, v518)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L3
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	if v1209 != 0 {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	goto L390
L392:
	;
	F_UnlockReleaseBuffer(m, v1209)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L3
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	F_UnlockReleaseBuffer(m, v1256)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L3
	} else {
		goto L396
	}
L395:
	;
	goto L394
L396:
	;
	F_UnlockReleaseBuffer(m, v1213)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L3
	} else {
		goto L397
	}
L397:
	;
	m.G0 = v1131 + int32(32)
	goto L254
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1131))) = v1235
	F_errmsg_internal(m, int32(_a_F_hashbucketcleanup_13), v1131)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L3
	} else {
		goto L399
	}
L399:
	;
	F_errfinish(m, int32(_a_F_hashbucketcleanup_14), int32(568), int32(_a_F_hashbucketcleanup_15))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L3
	} else {
		goto L400
	}
L400:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L401:
	;
	v1759 = int32(0)
	goto L404
L402:
	;
	goto L403
L403:
	;
	if v1089 == v1119 {
		goto L409
	} else {
		goto L410
	}
L404:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v522+int32(2464)+v1759<<(uint(int32(2))%32))))
	F_pfree(m, v1789)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L3
	} else {
		goto L406
	}
L405:
	;
	goto L403
L406:
	;
	v1793 = v1759 + int32(1)
	if v1793 != v1127 {
		v1759 = v1793
		goto L404
	} else {
		goto L407
	}
L407:
	;
	goto L405
L408:
	;
	goto L127
L409:
	;
	if v1089 != v511 {
		goto L408
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	v1831 = F__hash_getbuf_with_strategy(m, v515, v1119, int32(1), v510)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L3
	} else {
		goto L415
	}
L412:
	;
	F_UnlockBuffer(m, v1088)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L3
	} else {
		goto L413
	}
L413:
	;
	v1873 = v522
	v1885 = v534
	goto L106
L414:
	;
	v1851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1850)+16)))
	v506 = v1088
	v507 = v1089
	v514 = v1119
	v518 = v1831
	v524 = v1106
	v525 = v1850
	v526 = v1851 + v1850
	v533 = v1115
	goto L126
L415:
	;
	if v1831 < int32(0) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1836+(v1831^int32(-1))<<(uint(int32(2))%32))))
	v1850 = v1842
	goto L414
L417:
	;
	goto L418
L418:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1850 = v1844 + v1831<<(uint(int32(13))%32) + int32(-8192)
	goto L414
L419:
	;
	v1873 = v522
	v1885 = v534
	goto L106
L420:
	;
	v1873 = v422
	v1885 = v34
	goto L106
L421:
	;
	v1921 = v34
	goto L101
}
func F_hashfloat8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int64
	_ = v287
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	v11 = v8 & int64(9223372036854775807)
	if v11 == int64(0) {
		v287 = int64(0)
	} else {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v11) {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(9221120237041090560)
		} else {
		}
		v20 = v6 + int32(8)
		v27 = int32(-1636608424)
		if v20&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v247 = v27
				v248 = v27
				v249 = v27
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 1:
				v240 = v27
				v241 = v27
				v242 = v27
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 2:
				v233 = v27
				v234 = v27
				v235 = v27
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 3:
				v227 = v27
				v228 = v27
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 4:
				v223 = v27
				v224 = v27
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 5:
				v217 = v27
				v218 = v27
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v223 = v219<<(uint(int32(8))%32) + v217
				v224 = v218
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 6:
				v211 = v27
				v212 = v27
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
				v217 = v213<<(uint(int32(16))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v223 = v219<<(uint(int32(8))%32) + v217
				v224 = v218
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 7:
				v206 = v27
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)))
				v211 = v207<<(uint(int32(24))%32) + v27
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
				v217 = v213<<(uint(int32(16))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v223 = v219<<(uint(int32(8))%32) + v217
				v224 = v218
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 8:
				v201 = v27
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)))
				v211 = v207<<(uint(int32(24))%32) + v27
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
				v217 = v213<<(uint(int32(16))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v223 = v219<<(uint(int32(8))%32) + v217
				v224 = v218
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 9:
				v196 = v27
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
				v201 = v197<<(uint(int32(16))%32) + v196
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)))
				v211 = v207<<(uint(int32(24))%32) + v27
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
				v217 = v213<<(uint(int32(16))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v223 = v219<<(uint(int32(8))%32) + v217
				v224 = v218
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			case 10:
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)))
				v196 = v192<<(uint(int32(24))%32) + v27
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
				v201 = v197<<(uint(int32(16))%32) + v196
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)))
				v211 = v207<<(uint(int32(24))%32) + v27
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
				v217 = v213<<(uint(int32(16))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v223 = v219<<(uint(int32(8))%32) + v217
				v224 = v218
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v227 = v223 + v225
				v228 = v224
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v233 = v229<<(uint(int32(24))%32) + v27
				v234 = v227
				v235 = v228
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v240 = v236<<(uint(int32(16))%32) + v233
				v241 = v234
				v242 = v235
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v247 = v243<<(uint(int32(8))%32) + v240
				v248 = v241
				v249 = v242
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v247 + v250
				v255 = v248
				v256 = v249
			default:
				v254 = v27
				v255 = v27
				v256 = v27
			}
		} else {
			switch int32(7) {
			case 0:
				v133 = v27
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v133 + v134
				v255 = v27
				v256 = v27
			case 1:
				v128 = v27
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v133 = v129<<(uint(int32(8))%32) + v128
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v133 + v134
				v255 = v27
				v256 = v27
			case 2:
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v128 = v124<<(uint(int32(16))%32) + v27
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v133 = v129<<(uint(int32(8))%32) + v128
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v254 = v133 + v134
				v255 = v27
				v256 = v27
			case 3:
				v121 = v27
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v254 = v122 + v27
				v255 = v121
				v256 = v27
			case 4:
				v118 = v27
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v121 = v118 + v119
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v254 = v122 + v27
				v255 = v121
				v256 = v27
			case 5:
				v113 = v27
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v118 = v114<<(uint(int32(8))%32) + v113
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v121 = v118 + v119
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v254 = v122 + v27
				v255 = v121
				v256 = v27
			case 6:
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
				v113 = v109<<(uint(int32(16))%32) + v27
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
				v118 = v114<<(uint(int32(8))%32) + v113
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v121 = v118 + v119
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v254 = v122 + v27
				v255 = v121
				v256 = v27
			case 7:
				v104 = v27
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v254 = v105 + v27
				v255 = v107 + v27
				v256 = v104
			case 8:
				v99 = v27
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
				v104 = v100<<(uint(int32(8))%32) + v99
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v254 = v105 + v27
				v255 = v107 + v27
				v256 = v104
			case 9:
				v94 = v27
				v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
				v99 = v95<<(uint(int32(16))%32) + v94
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
				v104 = v100<<(uint(int32(8))%32) + v99
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v254 = v105 + v27
				v255 = v107 + v27
				v256 = v104
			case 10:
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)))
				v94 = v90<<(uint(int32(24))%32) + v27
				v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
				v99 = v95<<(uint(int32(16))%32) + v94
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
				v104 = v100<<(uint(int32(8))%32) + v99
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v254 = v105 + v27
				v255 = v107 + v27
				v256 = v104
			default:
				v254 = v27
				v255 = v27
				v256 = v27
			}
		}
		v259 = int32(14)
		v261 = v255 ^ v256 - base.I32_rotl(v255, v259)
		v265 = v261 ^ v254 - base.I32_rotl(v261, int32(11))
		v269 = v265 ^ v255 - base.I32_rotl(v265, int32(25))
		v273 = v269 ^ v261 - base.I32_rotl(v269, int32(16))
		v277 = v273 ^ v265 - base.I32_rotl(v273, int32(4))
		v281 = v277 ^ v269 - base.I32_rotl(v277, v259)
		v287 = base.I64_extend_i32_u(v281 ^ v273 - base.I32_rotl(v281, int32(24)))
	}
	m.G0 = v6 + int32(16)
	return v287
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
			v42 = F__hash_next(m, l0, l1)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				return v42
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			if v20 == int32(0) {
				v25 = F_palloc_mul(m, int32(4), int32(408))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v25
					v28 = v25
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					if int32(407) < v29 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v29 + int32(1)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32)))) = v38
					}
					v42 = F__hash_next(m, l0, l1)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						return v42
					}
				}
			} else {
				v28 = v20
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				if int32(407) < v29 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v29 + int32(1)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32)))) = v38
				}
				v42 = F__hash_next(m, l0, l1)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					return v42
				}
			}
		}
	}
}
func F_hashhandler(m *base.Module, l0 int32) int64 {
	return int64(829784)
}
func F_hashint2extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 == int64(0) {
		v10 = int32(-1636608428)
		v49 = v10
		v50 = v10
		v53 = int32(0)
	} else {
		v13 = base.I32_wrap_i64(v3)
		v15 = v13 + int32(1021750440)
		v20 = base.I32_wrap_i64(int64(base.Ui64(v3)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v26 = v13 - v20 - int32(1636608428) ^ base.I32_rotl(v20, int32(6))
		v30 = v15 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v20 + v15
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v37 + v33
		v47 = v41 + v46
		v49 = v47
		v50 = v46
		v53 = v33 - v41 ^ base.I32_rotl(v41, int32(4)) ^ v47
	}
	v54 = int32(14)
	v56 = v53 - base.I32_rotl(v49, v54)
	v61 = v56 ^ (v2 + v50) - base.I32_rotl(v56, int32(11))
	v65 = v49 ^ v61 - base.I32_rotl(v61, int32(25))
	v69 = v65 ^ v56 - base.I32_rotl(v65, int32(16))
	v73 = v69 ^ v61 - base.I32_rotl(v69, int32(4))
	v77 = v73 ^ v65 - base.I32_rotl(v73, v54)
	return base.I64_extend_i32_u(v77)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v77^v69-base.I32_rotl(v77, int32(24)))
}
func F_hashint8extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v11 int64
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v11 == int64(0) {
		v18 = int32(-1636608428)
		v57 = v18
		v58 = v18
		v61 = int32(0)
	} else {
		v21 = base.I32_wrap_i64(v11)
		v23 = v21 + int32(1021750440)
		v28 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v34 = v21 - v28 - int32(1636608428) ^ base.I32_rotl(v28, int32(6))
		v38 = v23 - v34 ^ base.I32_rotl(v34, int32(8))
		v39 = v28 + v23
		v40 = v34 + v39
		v41 = v38 + v40
		v45 = v39 - v38 ^ base.I32_rotl(v38, int32(16))
		v49 = v40 - v45 ^ base.I32_rotl(v45, int32(19))
		v54 = v45 + v41
		v55 = v49 + v54
		v57 = v55
		v58 = v54
		v61 = v41 - v49 ^ base.I32_rotl(v49, int32(4)) ^ v55
	}
	v62 = int32(14)
	v64 = v61 - base.I32_rotl(v57, v62)
	v69 = v64 ^ (base.I32_wrap_i64(v3>>(uint(int64(63))%64)^int64(base.Ui64(v3)>>(uint(int64(32))%64))^v3) + v58) - base.I32_rotl(v64, int32(11))
	v73 = v57 ^ v69 - base.I32_rotl(v69, int32(25))
	v77 = v73 ^ v64 - base.I32_rotl(v73, int32(16))
	v81 = v77 ^ v69 - base.I32_rotl(v77, int32(4))
	v85 = v81 ^ v73 - base.I32_rotl(v81, v62)
	return base.I64_extend_i32_u(v85)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v85^v77-base.I32_rotl(v85, int32(24)))
}
func F_hashmacaddr(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return base.I64_extend_i32_u(v263 ^ v255 - base.I32_rotl(v263, int32(24)))
}
func F_hashmacaddr8(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return base.I64_extend_i32_u(v263 ^ v255 - base.I32_rotl(v263, int32(24)))
}
func F_hashoptions(m *base.Module, l0 int64, l1 int32) int32 {
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
					v32 = v28 * int32(56)
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
						v32 = v28 * int32(56)
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
							v32 = v28 * int32(56)
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
func F_hashtextextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
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
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
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
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
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
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int64
	_ = v415
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v732 int32
	_ = v732
	var v738 int64
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 != 0 {
			v16 = F_pg_newlocale_from_collation(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v18 = int32(1)
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				if v20&v18 != 0 {
					v23 = v18
				} else {
					v23 = int32(4)
				}
				v24 = v11 + v23
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v25 == int32(1) {
					if v20 == int32(1) {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
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
						v56 = v43
					} else {
						v44 = int32(1)
						if v20&v44 != 0 {
							v56 = int32(base.Ui32(v20)>>(uint(v44)%32)) - v44
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v63 = v56 - int32(1636608432)
					if v57 == int64(0) {
						v100 = v63
						v102 = v63
						v104 = v63
					} else {
						v67 = v63 + base.I32_wrap_i64(v57)
						v68 = v67 + v63
						v72 = int32(4)
						v74 = base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(32))%64))) ^ base.I32_rotl(v63, v72)
						v78 = v67 - v74 ^ base.I32_rotl(v74, int32(6))
						v82 = v68 - v78 ^ base.I32_rotl(v78, int32(8))
						v83 = v68 + v74
						v84 = v78 + v83
						v85 = v82 + v84
						v89 = v83 - v82 ^ base.I32_rotl(v82, int32(16))
						v93 = v84 - v89 ^ base.I32_rotl(v89, int32(19))
						v98 = v85 + v89
						v100 = v98
						v102 = v85 - v93 ^ base.I32_rotl(v93, v72)
						v104 = v93 + v98
					}
					if v24&int32(3) != 0 {
						if base.Ui32(int32(11)) < base.Ui32(v56) {
							v109 = v24
							v110 = v56
							v112 = v100
							v113 = v104
							v114 = v102
							for {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
								v117 = v116 + v113
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v121 = v120 + v114
								v123 = int32(4)
								v125 = v118 + v112 - v121 ^ base.I32_rotl(v121, v123)
								v129 = v117 - v125 ^ base.I32_rotl(v125, int32(6))
								v130 = v121 + v117
								v131 = v125 + v130
								v132 = v129 + v131
								v136 = v130 - v129 ^ base.I32_rotl(v129, int32(8))
								v140 = v131 - v136 ^ base.I32_rotl(v136, int32(16))
								v144 = v132 - v140 ^ base.I32_rotl(v140, int32(19))
								v145 = v136 + v132
								v146 = v140 + v145
								v147 = v144 + v146
								v151 = v145 - v144 ^ base.I32_rotl(v144, v123)
								v152 = int32(12)
								v153 = v109 + v152
								v155 = v110 - v152
								if base.Ui32(int32(11)) < base.Ui32(v155) {
									v109 = v153
									v110 = v155
									v112 = v146
									v113 = v147
									v114 = v151
									continue
								} else {
									break
								}
								break
							}
							v158 = v153
							v159 = v155
							v161 = v146
							v162 = v147
							v163 = v151
						} else {
							v158 = v24
							v159 = v56
							v161 = v100
							v162 = v104
							v163 = v102
						}
						switch v159 - int32(1) {
						case 0:
							v328 = v161
							v329 = v162
							v330 = v163
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 1:
							v321 = v161
							v322 = v162
							v323 = v163
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 2:
							v314 = v161
							v315 = v162
							v316 = v163
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 3:
							v308 = v162
							v309 = v163
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 4:
							v304 = v162
							v305 = v163
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 5:
							v298 = v162
							v299 = v163
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 6:
							v292 = v162
							v293 = v163
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 7:
							v287 = v163
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 8:
							v282 = v163
							v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
							v287 = v283<<(uint(int32(8))%32) + v282
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 9:
							v277 = v163
							v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+9)))
							v282 = v278<<(uint(int32(16))%32) + v277
							v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
							v287 = v283<<(uint(int32(8))%32) + v282
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 10:
							v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+10)))
							v277 = v273<<(uint(int32(24))%32) + v163
							v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+9)))
							v282 = v278<<(uint(int32(16))%32) + v277
							v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
							v287 = v283<<(uint(int32(8))%32) + v282
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						default:
							v336 = v161
							v337 = v162
							v338 = v163
						}
					} else {
						if base.Ui32(int32(12)) <= base.Ui32(v56) {
							v169 = v24
							v170 = v56
							v172 = v100
							v173 = v104
							v174 = v102
							for {
								v176 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
								v177 = v176 + v173
								v178 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
								v181 = v180 + v174
								v183 = int32(4)
								v185 = v178 + v172 - v181 ^ base.I32_rotl(v181, v183)
								v189 = v177 - v185 ^ base.I32_rotl(v185, int32(6))
								v190 = v181 + v177
								v191 = v185 + v190
								v192 = v189 + v191
								v196 = v190 - v189 ^ base.I32_rotl(v189, int32(8))
								v200 = v191 - v196 ^ base.I32_rotl(v196, int32(16))
								v204 = v192 - v200 ^ base.I32_rotl(v200, int32(19))
								v205 = v196 + v192
								v206 = v200 + v205
								v207 = v204 + v206
								v211 = v205 - v204 ^ base.I32_rotl(v204, v183)
								v212 = int32(12)
								v213 = v169 + v212
								v215 = v170 - v212
								if base.Ui32(int32(11)) < base.Ui32(v215) {
									v169 = v213
									v170 = v215
									v172 = v206
									v173 = v207
									v174 = v211
									continue
								} else {
									break
								}
								break
							}
							v218 = v213
							v219 = v215
							v221 = v206
							v222 = v207
							v223 = v211
						} else {
							v218 = v24
							v219 = v56
							v221 = v100
							v222 = v104
							v223 = v102
						}
						switch v219 - int32(1) {
						case 0:
							v270 = v221
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
							v336 = v270 + v271
							v337 = v222
							v338 = v223
						case 1:
							v265 = v221
							v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
							v270 = v266<<(uint(int32(8))%32) + v265
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
							v336 = v270 + v271
							v337 = v222
							v338 = v223
						case 2:
							v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+2)))
							v265 = v261<<(uint(int32(16))%32) + v221
							v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
							v270 = v266<<(uint(int32(8))%32) + v265
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
							v336 = v270 + v271
							v337 = v222
							v338 = v223
						case 3:
							v258 = v222
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 4:
							v255 = v222
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							v258 = v255 + v256
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 5:
							v250 = v222
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+5)))
							v255 = v251<<(uint(int32(8))%32) + v250
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							v258 = v255 + v256
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 6:
							v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+6)))
							v250 = v246<<(uint(int32(16))%32) + v222
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+5)))
							v255 = v251<<(uint(int32(8))%32) + v250
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							v258 = v255 + v256
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 7:
							v241 = v223
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						case 8:
							v236 = v223
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
							v241 = v237<<(uint(int32(8))%32) + v236
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						case 9:
							v231 = v223
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)))
							v236 = v232<<(uint(int32(16))%32) + v231
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
							v241 = v237<<(uint(int32(8))%32) + v236
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						case 10:
							v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+10)))
							v231 = v227<<(uint(int32(24))%32) + v223
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)))
							v236 = v232<<(uint(int32(16))%32) + v231
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
							v241 = v237<<(uint(int32(8))%32) + v236
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						default:
							v336 = v221
							v337 = v222
							v338 = v223
						}
					}
					v341 = int32(14)
					v343 = v337 ^ v338 - base.I32_rotl(v337, v341)
					v347 = v343 ^ v336 - base.I32_rotl(v343, int32(11))
					v351 = v347 ^ v337 - base.I32_rotl(v347, int32(25))
					v355 = v351 ^ v343 - base.I32_rotl(v351, int32(16))
					v359 = v355 ^ v347 - base.I32_rotl(v355, int32(4))
					v363 = v359 ^ v351 - base.I32_rotl(v359, v341)
					v738 = base.I64_extend_i32_u(v363)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v363^v355-base.I32_rotl(v363, int32(24)))
					v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v739 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v742 = m.ExcPending
						if v742 != 0 {
							return int64(0)
						} else {
							return v738
						}
					} else {
						return v738
					}
				} else {
					v373 = int32(0)
					if v20 == int32(1) {
						v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
						if v380 == int32(18) {
							v383 = int32(16)
						} else {
							v383 = int32(0)
						}
						if base.Ui32((v380-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v390 = int32(4)
						} else {
							v390 = v383
						}
						v403 = v390
					} else {
						v391 = int32(1)
						if v20&v391 != 0 {
							v403 = int32(base.Ui32(v20)>>(uint(v391)%32)) - v391
						} else {
							v397 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v403 = int32(base.Ui32(v397)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v404 = F_pg_strnxfrm(m, v373, v373, v24, v403, v16)
					mBase = m.M
					v405 = m.ExcPending
					if v405 != 0 {
						return int64(0)
					} else {
						v407 = v404 + int32(1)
						v408 = F_palloc(m, v407)
						mBase = m.M
						v409 = m.ExcPending
						if v409 != 0 {
							return int64(0)
						} else {
							v410 = F_pg_strnxfrm(m, v408, v407, v24, v403, v16)
							mBase = m.M
							v411 = m.ExcPending
							if v411 != 0 {
								return int64(0)
							} else {
								if base.Ui32(v404) < base.Ui32(v410) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v767 = m.ExcPending
									if v767 != 0 {
										return int64(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_hashtextextended_0), int32(0))
										mBase = m.M
										v771 = m.ExcPending
										if v771 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_hashtextextended_1), int32(361), int32(_a_F_hashtextextended_2))
											mBase = m.M
											v776 = m.ExcPending
											if v776 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v414 = v410 + int32(1)
									v415 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
									v421 = v414 - int32(1636608432)
									if v415 == int64(0) {
										v458 = v421
										v460 = v421
										v462 = v421
									} else {
										v425 = v421 + base.I32_wrap_i64(v415)
										v426 = v425 + v421
										v430 = int32(4)
										v432 = base.I32_wrap_i64(int64(base.Ui64(v415)>>(uint(int64(32))%64))) ^ base.I32_rotl(v421, v430)
										v436 = v425 - v432 ^ base.I32_rotl(v432, int32(6))
										v440 = v426 - v436 ^ base.I32_rotl(v436, int32(8))
										v441 = v426 + v432
										v442 = v436 + v441
										v443 = v440 + v442
										v447 = v441 - v440 ^ base.I32_rotl(v440, int32(16))
										v451 = v442 - v447 ^ base.I32_rotl(v447, int32(19))
										v456 = v443 + v447
										v458 = v456
										v460 = v443 - v451 ^ base.I32_rotl(v451, v430)
										v462 = v451 + v456
									}
									if v408&int32(3) != 0 {
										if base.Ui32(int32(11)) < base.Ui32(v414) {
											v467 = v408
											v468 = v414
											v470 = v458
											v471 = v462
											v472 = v460
											for {
												v474 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
												v475 = v474 + v471
												v476 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
												v478 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
												v479 = v478 + v472
												v481 = int32(4)
												v483 = v476 + v470 - v479 ^ base.I32_rotl(v479, v481)
												v487 = v475 - v483 ^ base.I32_rotl(v483, int32(6))
												v488 = v479 + v475
												v489 = v483 + v488
												v490 = v487 + v489
												v494 = v488 - v487 ^ base.I32_rotl(v487, int32(8))
												v498 = v489 - v494 ^ base.I32_rotl(v494, int32(16))
												v502 = v490 - v498 ^ base.I32_rotl(v498, int32(19))
												v503 = v494 + v490
												v504 = v498 + v503
												v505 = v502 + v504
												v509 = v503 - v502 ^ base.I32_rotl(v502, v481)
												v510 = int32(12)
												v511 = v467 + v510
												v513 = v468 - v510
												if base.Ui32(int32(11)) < base.Ui32(v513) {
													v467 = v511
													v468 = v513
													v470 = v504
													v471 = v505
													v472 = v509
													continue
												} else {
													break
												}
												break
											}
											v516 = v511
											v517 = v513
											v519 = v504
											v520 = v505
											v521 = v509
										} else {
											v516 = v408
											v517 = v414
											v519 = v458
											v520 = v462
											v521 = v460
										}
										switch v517 - int32(1) {
										case 0:
											v686 = v519
											v687 = v520
											v688 = v521
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 1:
											v679 = v519
											v680 = v520
											v681 = v521
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 2:
											v672 = v519
											v673 = v520
											v674 = v521
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 3:
											v666 = v520
											v667 = v521
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 4:
											v662 = v520
											v663 = v521
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 5:
											v656 = v520
											v657 = v521
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 6:
											v650 = v520
											v651 = v521
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 7:
											v645 = v521
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 8:
											v640 = v521
											v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
											v645 = v641<<(uint(int32(8))%32) + v640
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 9:
											v635 = v521
											v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+9)))
											v640 = v636<<(uint(int32(16))%32) + v635
											v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
											v645 = v641<<(uint(int32(8))%32) + v640
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 10:
											v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+10)))
											v635 = v631<<(uint(int32(24))%32) + v521
											v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+9)))
											v640 = v636<<(uint(int32(16))%32) + v635
											v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
											v645 = v641<<(uint(int32(8))%32) + v640
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										default:
											v694 = v519
											v695 = v520
											v696 = v521
										}
									} else {
										if base.Ui32(int32(12)) <= base.Ui32(v414) {
											v527 = v408
											v528 = v414
											v530 = v458
											v531 = v462
											v532 = v460
											for {
												v534 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
												v535 = v534 + v531
												v536 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
												v538 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
												v539 = v538 + v532
												v541 = int32(4)
												v543 = v536 + v530 - v539 ^ base.I32_rotl(v539, v541)
												v547 = v535 - v543 ^ base.I32_rotl(v543, int32(6))
												v548 = v539 + v535
												v549 = v543 + v548
												v550 = v547 + v549
												v554 = v548 - v547 ^ base.I32_rotl(v547, int32(8))
												v558 = v549 - v554 ^ base.I32_rotl(v554, int32(16))
												v562 = v550 - v558 ^ base.I32_rotl(v558, int32(19))
												v563 = v554 + v550
												v564 = v558 + v563
												v565 = v562 + v564
												v569 = v563 - v562 ^ base.I32_rotl(v562, v541)
												v570 = int32(12)
												v571 = v527 + v570
												v573 = v528 - v570
												if base.Ui32(int32(11)) < base.Ui32(v573) {
													v527 = v571
													v528 = v573
													v530 = v564
													v531 = v565
													v532 = v569
													continue
												} else {
													break
												}
												break
											}
											v576 = v571
											v577 = v573
											v579 = v564
											v580 = v565
											v581 = v569
										} else {
											v576 = v408
											v577 = v414
											v579 = v458
											v580 = v462
											v581 = v460
										}
										switch v577 - int32(1) {
										case 0:
											v628 = v579
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
											v694 = v628 + v629
											v695 = v580
											v696 = v581
										case 1:
											v623 = v579
											v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
											v628 = v624<<(uint(int32(8))%32) + v623
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
											v694 = v628 + v629
											v695 = v580
											v696 = v581
										case 2:
											v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+2)))
											v623 = v619<<(uint(int32(16))%32) + v579
											v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
											v628 = v624<<(uint(int32(8))%32) + v623
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
											v694 = v628 + v629
											v695 = v580
											v696 = v581
										case 3:
											v616 = v580
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 4:
											v613 = v580
											v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+4)))
											v616 = v613 + v614
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 5:
											v608 = v580
											v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+5)))
											v613 = v609<<(uint(int32(8))%32) + v608
											v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+4)))
											v616 = v613 + v614
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 6:
											v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+6)))
											v608 = v604<<(uint(int32(16))%32) + v580
											v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+5)))
											v613 = v609<<(uint(int32(8))%32) + v608
											v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+4)))
											v616 = v613 + v614
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 7:
											v599 = v581
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										case 8:
											v594 = v581
											v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+8)))
											v599 = v595<<(uint(int32(8))%32) + v594
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										case 9:
											v589 = v581
											v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+9)))
											v594 = v590<<(uint(int32(16))%32) + v589
											v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+8)))
											v599 = v595<<(uint(int32(8))%32) + v594
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										case 10:
											v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+10)))
											v589 = v585<<(uint(int32(24))%32) + v581
											v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+9)))
											v594 = v590<<(uint(int32(16))%32) + v589
											v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+8)))
											v599 = v595<<(uint(int32(8))%32) + v594
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										default:
											v694 = v579
											v695 = v580
											v696 = v581
										}
									}
									v699 = int32(14)
									v701 = v695 ^ v696 - base.I32_rotl(v695, v699)
									v705 = v701 ^ v694 - base.I32_rotl(v701, int32(11))
									v709 = v705 ^ v695 - base.I32_rotl(v705, int32(25))
									v713 = v709 ^ v701 - base.I32_rotl(v709, int32(16))
									v717 = v713 ^ v705 - base.I32_rotl(v713, int32(4))
									v721 = v717 ^ v709 - base.I32_rotl(v717, v699)
									F_pfree(m, v408)
									mBase = m.M
									v732 = m.ExcPending
									if v732 != 0 {
										return int64(0)
									} else {
										v738 = base.I64_extend_i32_u(v721)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v721^v713-base.I32_rotl(v721, int32(24)))
										v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v739 != v11 {
											F_pfree(m, v11)
											mBase = m.M
											v742 = m.ExcPending
											if v742 != 0 {
												return int64(0)
											} else {
												return v738
											}
										} else {
											return v738
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
			v747 = m.ExcPending
			if v747 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v750 = m.ExcPending
				if v750 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_hashtextextended_3), int32(0))
					mBase = m.M
					v754 = m.ExcPending
					if v754 != 0 {
						return int64(0)
					} else {
						F_errhint(m, int32(_a_F_hashtextextended_4), int32(0))
						mBase = m.M
						v758 = m.ExcPending
						if v758 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_hashtextextended_1), int32(336), int32(_a_F_hashtextextended_2))
							mBase = m.M
							v763 = m.ExcPending
							if v763 != 0 {
								return int64(0)
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
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
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v593 int32
	_ = v593
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(208)
	m.G0 = v18
	v22 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	if int32(0) < v245 {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	return int32(0)
L3:
	;
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+80))
	v31 = F_get_opfamily_name(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L49
	}
L7:
	;
	v35 = base.I64_extend_i32_u(v30)
	v36 = int64(0)
	v38 = F_SearchSysCacheList(m, int32(4), int32(1), v35, v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(1)
	v43 = int64(0)
	v45 = F_SearchSysCacheList(m, int32(5), v40, v35, v43, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v47 <= int32(0) {
		v233 = v40
		v238 = v2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v55 = v40
	v56 = v2
	v60 = v2
	goto L11
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v45-int32(-64)+v56<<(uint(int32(2))%32))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v73 = v71 + v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	if v74 == v75 {
		v104 = v55
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v233 = v210
	v238 = v211
	goto L1
L13:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+16)))
	switch v105 - int32(1) {
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
	v77 = int32(0)
	v80 = F_errstart(m, int32(17), v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v80 == int32(0) {
		v104 = v77
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v88 = F_format_procedure(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_1), v18+int32(192))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(91), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v104 = v77
	goto L13
L21:
	;
	v214 = v56 + int32(1)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v214 < v215 {
		v55 = v210
		v56 = v214
		v60 = v211
		goto L11
	} else {
		goto L48
	}
L22:
	;
	v180 = int32(0)
	v183 = F_errstart(m, int32(17), v180)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L42
	}
L23:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+16)))
	v170 = int32(1)
	if base.Ui32(v170) < base.Ui32((v169-v170)&int32(_a_F_hashvalidate_4)) {
		v210 = v104
		v211 = v60
		goto L21
	} else {
		goto L40
	}
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v156
	v159 = int32(1)
	v164 = F_check_amproc_signature(m, v155, int32(23), v159, v159, v159, v18+int32(160))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L38
	}
L25:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v144 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v143
	v149 = int32(2)
	v153 = F_check_amproc_signature(m, v142, v144, int32(1), v149, v149, v18+int32(176))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L36
	}
L26:
	;
	v113 = int32(0)
	v116 = F_errstart(m, int32(17), v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L30
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v109 = F_check_amoptsproc_signature(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v109 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	if v116 == int32(0) {
		v210 = v113
		v211 = v60
		goto L21
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v124 = F_format_procedure(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_5), v18+int32(128))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(115), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v210 = v113
	v211 = v60
	goto L21
L36:
	;
	if v153 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	if v164 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v177 = F_list_append_unique_oid(m, v60, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v210 = v104
	v211 = v177
	goto L21
L42:
	;
	if v183 == int32(0) {
		v210 = v180
		v211 = v60
		goto L21
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v191 = F_format_procedure(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+156)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v18)+152)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v18)+148)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_6), v18+int32(144))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(127), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v210 = v180
	v211 = v60
	goto L21
L48:
	;
	goto L12
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(_a_F_hashvalidate_7), v18)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(60), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	v254 = v233
	v255 = int32(0)
	goto L55
L53:
	;
	v490 = v233
	goto L54
L54:
	;
	v502 = F_identify_opfamily_groups(m, v38, v45)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L2
	} else {
		goto L126
	}
L55:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(-64)+v255<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+72))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+22)))
	v272 = v270 + v271
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272)+16)))
	if v273 == int32(1) {
		v306 = v254
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v490 = v482
	goto L54
L57:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+18)))
	if v307 == int32(115) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v276 = int32(0)
	v279 = F_errstart(m, int32(17), v276)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	if v279 == int32(0) {
		v306 = v276
		goto L57
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v287 = F_format_operator(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v289 = int32(*(*int16)(unsafe.Add(mBase, uint32(v272)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_8), v18+int32(112))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(153), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v306 = v276
	goto L57
L65:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v345 = F_check_amop_signature(m, v341, int32(16), v343, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L77
	}
L66:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	if v310 == int32(0) {
		v340 = v306
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v313 = int32(0)
	v316 = F_errstart(m, int32(17), v313)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v316 == int32(0) {
		v340 = v313
		goto L65
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v324 = F_format_operator(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_9), v18+int32(96))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(165), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v340 = v313
	goto L65
L76:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v376 = int32(0)
	if v238 == v376 {
		goto L87
	} else {
		goto L88
	}
L77:
	;
	if v345 != 0 {
		v374 = v340
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v347 = int32(0)
	v350 = F_errstart(m, int32(17), v347)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v350 == int32(0) {
		v374 = v347
		goto L76
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v358 = F_format_operator(m, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_10), v18+int32(80))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(178), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v374 = v347
	goto L76
L85:
	;
	v484 = v255 + int32(1)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	if v484 < v485 {
		v254 = v482
		v255 = v484
		goto L55
	} else {
		goto L122
	}
L86:
	;
	if v414 != 0 {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v414 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v382 <= int32(0) {
		v408 = v376
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v414 = v408
	goto L86
L91:
	;
	v385 = int32(0)
	if v385 < v382 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v388 = v382
	goto L94
L93:
	;
	v388 = v385
	goto L94
L94:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v391 = int32(0)
	goto L95
L95:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v389+v391<<(uint(int32(2))%32))))
	v400 = base.B2i32(v399 == v375)
	if v399 == v375 {
		v408 = v400
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v408 = v400
	goto L90
L97:
	;
	v402 = v391 + int32(1)
	if v402 != v388 {
		v391 = v402
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v416 = int32(0)
	if v238 == v416 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	v455 = int32(0)
	v458 = F_errstart(m, int32(17), v455)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L116
	}
L102:
	;
	if v454 != 0 {
		v482 = v374
		goto L85
	} else {
		goto L115
	}
L103:
	;
	v454 = int32(0)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v422 <= int32(0) {
		v448 = v416
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v454 = v448
	goto L102
L107:
	;
	v425 = int32(0)
	if v425 < v422 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v428 = v422
	goto L110
L109:
	;
	v428 = v425
	goto L110
L110:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v431 = int32(0)
	goto L111
L111:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v429+v431<<(uint(int32(2))%32))))
	v440 = base.B2i32(v439 == v415)
	if v439 == v415 {
		v448 = v440
		goto L106
	} else {
		goto L113
	}
L112:
	;
	v448 = v440
	goto L106
L113:
	;
	v442 = v431 + int32(1)
	if v442 != v428 {
		v431 = v442
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
	if v458 == int32(0) {
		v482 = v455
		goto L85
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v466 = F_format_operator(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_11), v18-int32(-64))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(190), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v482 = v455
	goto L85
L122:
	;
	goto L56
L123:
	;
	F_ReleaseCatCacheList(m, v45)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L2
	} else {
		goto L169
	}
L124:
	;
	if v238 != 0 {
		goto L160
	} else {
		goto L161
	}
L125:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v656 = v640
	v668 = v652
	goto L124
L126:
	;
	if v502 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	if int32(0) < v504 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v611 = int32(0)
	v614 = F_errstart(m, int32(17), v611)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L152
	}
L130:
	;
	v507 = int32(0)
	v511 = v507
	v512 = v490
	v513 = v507
	goto L133
L131:
	;
	v581 = v490
	v593 = int32(1)
	goto L132
L132:
	;
	if v593 == int32(0) {
		v640 = v581
		goto L125
	} else {
		goto L151
	}
L133:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v502)+12))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524+v513<<(uint(int32(2))%32))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	if v29 == v529 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v581 = v569
	v593 = base.B2i32(v534 == int32(0))
	goto L132
L135:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	if v531 == v29 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v534 = v511
	goto L137
L137:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v528)+8))
	if v535 == int64(2) {
		v569 = v512
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v533 = v528
	goto L140
L139:
	;
	v533 = v511
	goto L140
L140:
	;
	v534 = v533
	goto L137
L141:
	;
	v572 = v513 + int32(1)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	if v572 < v573 {
		v511 = v534
		v512 = v569
		v513 = v572
		goto L133
	} else {
		goto L150
	}
L142:
	;
	v538 = int32(0)
	v541 = F_errstart(m, int32(17), v538)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	if v541 == int32(0) {
		v569 = v538
		goto L141
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v549 = F_format_type_be(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	v552 = F_format_type_be(m, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L2
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_12), v18+int32(48))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(219), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	v569 = v538
	goto L141
L150:
	;
	goto L134
L151:
	;
	goto L129
L152:
	;
	if v614 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L2
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v634 = int32(0)
	if v502 == v634 {
		v656 = v611
		v668 = v634
		goto L124
	} else {
		goto L159
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v28 + int32(8)
	F_errmsg(m, int32(_a_F_hashvalidate_13), v18+int32(32))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(231), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	goto L155
L159:
	;
	v640 = v611
	goto L125
L160:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v673 = v669 * v669
	goto L162
L161:
	;
	v673 = int32(0)
	goto L162
L162:
	;
	if v668 == v673 {
		v698 = v656
		goto L123
	} else {
		goto L163
	}
L163:
	;
	v675 = int32(0)
	v678 = F_errstart(m, int32(17), v675)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	if v678 == int32(0) {
		v698 = v675
		goto L123
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v31
	F_errmsg(m, int32(_a_F_hashvalidate_14), v18+int32(16))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(247), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	v698 = v675
	goto L123
L169:
	;
	F_ReleaseCatCacheList(m, v38)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	F_ReleaseCatCache(m, v22)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	m.G0 = v18 + int32(208)
	return v698
}
func F_hashvarlenaextended(m *base.Module, l0 int32) int64 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
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
	var v203 int32
	_ = v203
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
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v15 = v13 & v11
		if v15 != 0 {
			v16 = v11
		} else {
			v16 = int32(4)
		}
		v17 = v7 + v16
		if v13 == int32(1) {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if v23 == int32(18) {
				v26 = int32(16)
			} else {
				v26 = int32(0)
			}
			if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v33 = int32(4)
			} else {
				v33 = v26
			}
			v44 = v33
		} else {
			v34 = int32(1)
			if v15 != 0 {
				v44 = int32(base.Ui32(v13)>>(uint(v34)%32)) - v34
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v51 = v44 - int32(1636608432)
		if v45 == int64(0) {
			v88 = v51
			v90 = v51
			v92 = v51
		} else {
			v55 = v51 + base.I32_wrap_i64(v45)
			v56 = v55 + v51
			v60 = int32(4)
			v62 = base.I32_wrap_i64(int64(base.Ui64(v45)>>(uint(int64(32))%64))) ^ base.I32_rotl(v51, v60)
			v66 = v55 - v62 ^ base.I32_rotl(v62, int32(6))
			v70 = v56 - v66 ^ base.I32_rotl(v66, int32(8))
			v71 = v56 + v62
			v72 = v66 + v71
			v73 = v70 + v72
			v77 = v71 - v70 ^ base.I32_rotl(v70, int32(16))
			v81 = v72 - v77 ^ base.I32_rotl(v77, int32(19))
			v86 = v73 + v77
			v88 = v86
			v90 = v73 - v81 ^ base.I32_rotl(v81, v60)
			v92 = v81 + v86
		}
		if v17&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v44) {
				v97 = v17
				v98 = v44
				v100 = v88
				v101 = v92
				v102 = v90
				for {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
					v105 = v104 + v101
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
					v109 = v108 + v102
					v111 = int32(4)
					v113 = v106 + v100 - v109 ^ base.I32_rotl(v109, v111)
					v117 = v105 - v113 ^ base.I32_rotl(v113, int32(6))
					v118 = v109 + v105
					v119 = v113 + v118
					v120 = v117 + v119
					v124 = v118 - v117 ^ base.I32_rotl(v117, int32(8))
					v128 = v119 - v124 ^ base.I32_rotl(v124, int32(16))
					v132 = v120 - v128 ^ base.I32_rotl(v128, int32(19))
					v133 = v124 + v120
					v134 = v128 + v133
					v135 = v132 + v134
					v139 = v133 - v132 ^ base.I32_rotl(v132, v111)
					v140 = int32(12)
					v141 = v97 + v140
					v143 = v98 - v140
					if base.Ui32(int32(11)) < base.Ui32(v143) {
						v97 = v141
						v98 = v143
						v100 = v134
						v101 = v135
						v102 = v139
						continue
					} else {
						break
					}
					break
				}
				v146 = v141
				v147 = v143
				v149 = v134
				v150 = v135
				v151 = v139
			} else {
				v146 = v17
				v147 = v44
				v149 = v88
				v150 = v92
				v151 = v90
			}
			switch v147 - int32(1) {
			case 0:
				v316 = v149
				v317 = v150
				v318 = v151
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 1:
				v309 = v149
				v310 = v150
				v311 = v151
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 2:
				v302 = v149
				v303 = v150
				v304 = v151
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 3:
				v296 = v150
				v297 = v151
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 4:
				v292 = v150
				v293 = v151
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 5:
				v286 = v150
				v287 = v151
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+5)))
				v292 = v288<<(uint(int32(8))%32) + v286
				v293 = v287
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 6:
				v280 = v150
				v281 = v151
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+6)))
				v286 = v282<<(uint(int32(16))%32) + v280
				v287 = v281
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+5)))
				v292 = v288<<(uint(int32(8))%32) + v286
				v293 = v287
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 7:
				v275 = v151
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+7)))
				v280 = v276<<(uint(int32(24))%32) + v150
				v281 = v275
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+6)))
				v286 = v282<<(uint(int32(16))%32) + v280
				v287 = v281
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+5)))
				v292 = v288<<(uint(int32(8))%32) + v286
				v293 = v287
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 8:
				v270 = v151
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+8)))
				v275 = v271<<(uint(int32(8))%32) + v270
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+7)))
				v280 = v276<<(uint(int32(24))%32) + v150
				v281 = v275
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+6)))
				v286 = v282<<(uint(int32(16))%32) + v280
				v287 = v281
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+5)))
				v292 = v288<<(uint(int32(8))%32) + v286
				v293 = v287
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 9:
				v265 = v151
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+9)))
				v270 = v266<<(uint(int32(16))%32) + v265
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+8)))
				v275 = v271<<(uint(int32(8))%32) + v270
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+7)))
				v280 = v276<<(uint(int32(24))%32) + v150
				v281 = v275
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+6)))
				v286 = v282<<(uint(int32(16))%32) + v280
				v287 = v281
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+5)))
				v292 = v288<<(uint(int32(8))%32) + v286
				v293 = v287
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			case 10:
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+10)))
				v265 = v261<<(uint(int32(24))%32) + v151
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+9)))
				v270 = v266<<(uint(int32(16))%32) + v265
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+8)))
				v275 = v271<<(uint(int32(8))%32) + v270
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+7)))
				v280 = v276<<(uint(int32(24))%32) + v150
				v281 = v275
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+6)))
				v286 = v282<<(uint(int32(16))%32) + v280
				v287 = v281
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+5)))
				v292 = v288<<(uint(int32(8))%32) + v286
				v293 = v287
				v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
				v296 = v292 + v294
				v297 = v293
				v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
				v302 = v298<<(uint(int32(24))%32) + v149
				v303 = v296
				v304 = v297
				v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
				v309 = v305<<(uint(int32(16))%32) + v302
				v310 = v303
				v311 = v304
				v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
				v316 = v312<<(uint(int32(8))%32) + v309
				v317 = v310
				v318 = v311
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
				v324 = v316 + v319
				v325 = v317
				v326 = v318
			default:
				v324 = v149
				v325 = v150
				v326 = v151
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v44) {
				v157 = v17
				v158 = v44
				v160 = v88
				v161 = v92
				v162 = v90
				for {
					v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
					v165 = v164 + v161
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
					v169 = v168 + v162
					v171 = int32(4)
					v173 = v166 + v160 - v169 ^ base.I32_rotl(v169, v171)
					v177 = v165 - v173 ^ base.I32_rotl(v173, int32(6))
					v178 = v169 + v165
					v179 = v173 + v178
					v180 = v177 + v179
					v184 = v178 - v177 ^ base.I32_rotl(v177, int32(8))
					v188 = v179 - v184 ^ base.I32_rotl(v184, int32(16))
					v192 = v180 - v188 ^ base.I32_rotl(v188, int32(19))
					v193 = v184 + v180
					v194 = v188 + v193
					v195 = v192 + v194
					v199 = v193 - v192 ^ base.I32_rotl(v192, v171)
					v200 = int32(12)
					v201 = v157 + v200
					v203 = v158 - v200
					if base.Ui32(int32(11)) < base.Ui32(v203) {
						v157 = v201
						v158 = v203
						v160 = v194
						v161 = v195
						v162 = v199
						continue
					} else {
						break
					}
					break
				}
				v206 = v201
				v207 = v203
				v209 = v194
				v210 = v195
				v211 = v199
			} else {
				v206 = v17
				v207 = v44
				v209 = v88
				v210 = v92
				v211 = v90
			}
			switch v207 - int32(1) {
			case 0:
				v258 = v209
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v324 = v258 + v259
				v325 = v210
				v326 = v211
			case 1:
				v253 = v209
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v258 = v254<<(uint(int32(8))%32) + v253
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v324 = v258 + v259
				v325 = v210
				v326 = v211
			case 2:
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v253 = v249<<(uint(int32(16))%32) + v209
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v258 = v254<<(uint(int32(8))%32) + v253
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v324 = v258 + v259
				v325 = v210
				v326 = v211
			case 3:
				v246 = v210
				v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v324 = v247 + v209
				v325 = v246
				v326 = v211
			case 4:
				v243 = v210
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v246 = v243 + v244
				v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v324 = v247 + v209
				v325 = v246
				v326 = v211
			case 5:
				v238 = v210
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v246 = v243 + v244
				v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v324 = v247 + v209
				v325 = v246
				v326 = v211
			case 6:
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
				v238 = v234<<(uint(int32(16))%32) + v210
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v246 = v243 + v244
				v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v324 = v247 + v209
				v325 = v246
				v326 = v211
			case 7:
				v229 = v211
				v230 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
				v324 = v230 + v209
				v325 = v232 + v210
				v326 = v229
			case 8:
				v224 = v211
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
				v229 = v225<<(uint(int32(8))%32) + v224
				v230 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
				v324 = v230 + v209
				v325 = v232 + v210
				v326 = v229
			case 9:
				v219 = v211
				v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+9)))
				v224 = v220<<(uint(int32(16))%32) + v219
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
				v229 = v225<<(uint(int32(8))%32) + v224
				v230 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
				v324 = v230 + v209
				v325 = v232 + v210
				v326 = v229
			case 10:
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
				v219 = v215<<(uint(int32(24))%32) + v211
				v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+9)))
				v224 = v220<<(uint(int32(16))%32) + v219
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
				v229 = v225<<(uint(int32(8))%32) + v224
				v230 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
				v324 = v230 + v209
				v325 = v232 + v210
				v326 = v229
			default:
				v324 = v209
				v325 = v210
				v326 = v211
			}
		}
		v329 = int32(14)
		v331 = v325 ^ v326 - base.I32_rotl(v325, v329)
		v335 = v331 ^ v324 - base.I32_rotl(v331, int32(11))
		v339 = v335 ^ v325 - base.I32_rotl(v335, int32(25))
		v343 = v339 ^ v331 - base.I32_rotl(v339, int32(16))
		v347 = v343 ^ v335 - base.I32_rotl(v343, int32(4))
		v351 = v347 ^ v339 - base.I32_rotl(v347, v329)
		v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v361 != v7 {
			F_pfree(m, v7)
			mBase = m.M
			v364 = m.ExcPending
			if v364 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v351)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v351^v343-base.I32_rotl(v351, int32(24)))
			}
		} else {
			return base.I64_extend_i32_u(v351)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v351^v343-base.I32_rotl(v351, int32(24)))
		}
	}
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
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v187 int64
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
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
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
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v234
	v238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v238
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v238)
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v83 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = base.I32_rotr(v102, int32(16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
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
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(116)+v126<<(uint(int32(1))%32)))))
	v140 = v68 + int32(20) + v137<<(uint(int32(2))%32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v68 + v141&int32(_a_F_heapgettup_pagemode_0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v137)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(base.Ui32(v146) >> (uint(int32(17)) % 32))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v211
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
	v180 = F_heap_getattr_1(m, l0+int32(68), v177, v157, v19+int32(15))
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
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v162)+48))
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
	if v187 == int64(0) {
		goto L34
	} else {
		goto L44
	}
L44:
	;
	v194 = v164 - int32(1)
	if v194 != 0 {
		v162 = v162 + int32(56)
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
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(293), int32(_a_F_hex_decode_safe_3))
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
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(298), int32(_a_F_hex_decode_safe_3))
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
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(303), int32(_a_F_hex_decode_safe_3))
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
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
func F_hnswhandler(m *base.Module, l0 int32) int64 {
	return int64(4171604)
}
func F_hnswoptions(m *base.Module, l0 int64, l1 int32) int32 {
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
func F_hypothetical_percent_rank_final(m *base.Module, l0 int32) int64 {
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
	var v26 int64
	_ = v26
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_hypothetical_rank_common(m, l0, int32(-1), v7+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		if v16 == int64(0) {
			v26 = int64(0)
		} else {
			v26 = base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v12-int64(1)), base.F64_convert_i64_s(v16)))
		}
		m.G0 = v7 + int32(16)
		return v26
	}
}
func F_hypothetical_rank_final(m *base.Module, l0 int32) int64 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = F_hypothetical_rank_common(m, l0, int32(-1), v6+int32(8))
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
