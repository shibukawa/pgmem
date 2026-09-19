package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	v6 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_backend_io_op[0]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v8))|base.B2i32(int32(1)<<(uint(v8)%32)&int32(_a_F_pgstat_count_backend_io_op_0) == v6) == v6 {
		v27 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pgstat_count_backend_io_op[1])))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pgstat_count_backend_io_op[1]))) = v28 + base.I64_extend_i32_u(l3)
		v32 = *(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pgstat_count_backend_io_op[2])))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pgstat_count_backend_io_op[2]))) = v32 + l4
		v36 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_backend_io_op[3])) = uint8(v36)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_backend_io_op[4])) = uint8(v36)
	} else {
	}
	return
}
func F_pgstat_drop_transactional(m *base.Module, l0 int32, l1 int32, l2 int64) {
	var v6 int32
	_ = v6
	F_create_drop_transactional_internal(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_fetch_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = int32(0)
	v6 = F_pgstat_get_entry_ref(m, l0, l1, l2, v4, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v11 != 0 {
				v12 = v6
			} else {
				v12 = int32(0)
			}
			v14 = v12
		} else {
			v14 = int32(0)
		}
		return v14
	}
}
func F_pgstat_fetch_stat_checkpointer(m *base.Module) int32 {
	var v5 int32
	_ = v5
	F_pgstat_snapshot_fixed(m, int32(9))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(_a_F_pgstat_fetch_stat_checkpointer_0)
	}
}
func F_pgstat_get_beentry_by_proc_number(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(432)
	m.G0 = v5
	F_pgstat_read_current_status(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+408)) = l0
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_beentry_by_proc_number[0]))
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_beentry_by_proc_number[1]))
		v18 = F_bsearch(m, v5, v13, v15, int32(432), int32(1196))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(432)
			return v18
		}
	}
}
func F_pgstat_get_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v265 int64
	_ = v265
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v275 int64
	_ = v275
	var v281 int64
	_ = v281
	var v287 int64
	_ = v287
	var v292 int64
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v323 int64
	_ = v323
	var v326 int32
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v344 int64
	_ = v344
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int64
	_ = v359
	var v369 int64
	_ = v369
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v432 int64
	_ = v432
	var v438 int64
	_ = v438
	var v443 int64
	_ = v443
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	var v490 int64
	_ = v490
	var v495 int64
	_ = v495
	var v497 int64
	_ = v497
	var v502 int64
	_ = v502
	var v508 int64
	_ = v508
	var v513 int64
	_ = v513
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v549 int64
	_ = v549
	var v551 int64
	_ = v551
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int64
	_ = v612
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v631 int64
	_ = v631
	var v633 int64
	_ = v633
	var v634 int64
	_ = v634
	var v640 int32
	_ = v640
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v645 int64
	_ = v645
	var v646 int64
	_ = v646
	var v647 int64
	_ = v647
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v659 int64
	_ = v659
	var v665 int64
	_ = v665
	var v670 int64
	_ = v670
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int64
	_ = v715
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int64
	_ = v771
	var v773 int64
	_ = v773
	var v775 int64
	_ = v775
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v802 int64
	_ = v802
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v863 int32
	_ = v863
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v995 int64
	_ = v995
	var v997 int64
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int64
	_ = v1110
	var v1112 int64
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1143 int32
	_ = v1143
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int64
	_ = v1173
	var v1175 int64
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1205 int32
	_ = v1205
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	v6 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l0
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[0]))
	if v27 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[1]))
	v37 = F_AllocSetContextCreateInternal(m, v32, int32(_a_F_pgstat_get_entry_ref_0), int32(0), int32(1024), int32(_a_F_pgstat_get_entry_ref_1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[2]))
	if v43 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[0])) = v37
	goto L3
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[1]))
	v53 = F_AllocSetContextCreateInternal(m, v48, int32(_a_F_pgstat_get_entry_ref_2), int32(0), int32(1024), int32(_a_F_pgstat_get_entry_ref_1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v56 = v43
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	if v58 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[2])) = v53
	v56 = v53
	goto L8
L10:
	;
	v62 = F_MemoryContextAllocZero(m, v56, int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if l4 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = v56
	v69 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_get_entry_ref_3), int32(5))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v62)+12)) = int64(987842478335)
	*(*int64)(unsafe.Add(mBase, uint32(v62))) = int64(256)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v69
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3])) = v62
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[4]))
	v81 = int64(0)
	v84 = base.AtomicRmwCmpxchg64(m, v80, int32(16), v81, v81)
	*(*uint32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[5])) = uint32(v84)
	goto L12
L15:
	;
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v88)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v90 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	if v92 == v90 {
		v244 = v90
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v262
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v261
	v265 = int64(23)
	v268 = int64(2388976653695081527)
	v269 = (v261 ^ int64(base.Ui64(v261)>>(uint(v265)%64))) * v268
	v270 = int64(47)
	v275 = int64(-8645972361240307355)
	v281 = (v262 ^ int64(base.Ui64(v262)>>(uint(v265)%64))) * v268
	v287 = ((v269^int64(base.Ui64(v269)>>(uint(v270)%64))^int64(-9208349263878056368))*v275 ^ int64(base.Ui64(v281)>>(uint(v270)%64)) ^ v281) * v275
	v292 = (int64(base.Ui64(v287)>>(uint(v265)%64)) ^ v287) * v268
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v311 = base.B2i32(base.Ui32(v300) < base.Ui32(v301))
	goto L42
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[4]))
	v97 = int64(0)
	v100 = base.AtomicRmwCmpxchg64(m, v96, int32(16), v97, v97)
	v102 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[5])))
	if v100 == v102 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	v244 = v105
	goto L18
L21:
	;
	goto L22
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[4]))
	v108 = int64(0)
	v111 = base.AtomicRmwCmpxchg64(m, v107, int32(16), v108, v108)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
	if v115 == v108 {
		v154 = int32(-1)
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v168 = v114
	v174 = v6
	v175 = v154
	goto L29
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	v125 = int32(0)
	goto L25
L25:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v125*int32(24))+16)))
	if v141 != int32(1) {
		v154 = v125
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v154 = int32(-1)
	goto L23
L27:
	;
	v145 = v125 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v145)) < base.Ui64(v115) {
		v125 = v145
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v192 = v174
	v193 = v175
	v195 = v174
	goto L32
L30:
	;
	*(*uint32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[5])) = uint32(v111)
	v244 = v168
	goto L18
L31:
	;
	goto L30
L32:
	;
	if v195&int32(1) != 0 {
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+16)))
	if v223 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v206 = int32(1)
	v207 = v193 - v206
	v211 = base.B2i32(v205&(v207^v154) == int32(0))
	v212 = v211 | v192
	v215 = v205 & v207
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	v217 = v193*int32(24) + v216
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+16)))
	if v218 != v206 {
		v192 = v212
		v193 = v215
		v195 = v211
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	if v226 == v227 {
		v174 = v212
		v175 = v215
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	if v229 != 0 {
		v174 = v212
		v175 = v215
		goto L29
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v217)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v230
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v217)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v232
	F_pgstat_release_entry_ref(m, v21+int32(48), v221, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	v168 = v240
	v174 = v212
	v175 = v215
	goto L29
L42:
	;
	if v311 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L4
	} else {
		goto L193
	}
L44:
	;
	goto L43
L45:
	;
	v1205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v1205
	v311 = v1205
	goto L42
L46:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_delete_entry(m, v1159, v916)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L4
	} else {
		goto L187
	}
L47:
	;
	m.G0 = v21 + int32(96)
	return v1143
L48:
	;
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	v905 = v21 - int32(-64)
	v907 = F_dshash_find(m, v903, v905, int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L136
	}
L49:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[0]))
	v876 = F_MemoryContextAlloc(m, v874, int32(24))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L4
	} else {
		goto L134
	}
L50:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v848 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v847 + v848
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+16)) = uint8(v848)
	*(*int64)(unsafe.Add(mBase, uint32(v837)+8)) = v262
	*(*int64)(unsafe.Add(mBase, uint32(v837))) = v261
	v863 = v837
	goto L49
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L131
	}
L52:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	if v323 == int64(4294967296) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v244)+20))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v604 = v603 & base.I32_wrap_i64(int64(base.Ui64(v292)>>(uint(v270)%64))^v292-int64(base.Ui64(v292)>>(uint(int64(32))%64)))
	v607 = v602 + v604*int32(24)
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+16)))
	if v608 == int32(0) {
		v837 = v607
		goto L50
	} else {
		goto L96
	}
L55:
	;
	v326 = int32(0)
	v328 = int64(2)
	v330 = v323 << (uint(int64(1)) % 64)
	if base.Ui64(v330) <= base.Ui64(v328) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v311 = int32(1)
	goto L42
L57:
	;
	v333 = v328
	goto L59
L58:
	;
	v333 = v330
	goto L59
L59:
	;
	v334 = int64(1)
	if v333&(v333-v334) == int64(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v344 = v333
	goto L62
L61:
	;
	v344 = v334 << (uint(int64(64)-base.I64_clz(v333)) % 64)
	goto L62
L62:
	;
	if base.Ui64(v344*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v244)+20))
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v244)+24))
	v356 = F_MemoryContextAllocExtended(m, v351, base.I32_wrap_i64(v344)*int32(24), int32(5))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	goto L44
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+20)) = v356
	v359 = int64(1)
	if v344&(v344-v359) == int64(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v369 = v344
	goto L69
L68:
	;
	v369 = v359 << (uint(int64(64)-base.I64_clz(v344)) % 64)
	goto L69
L69:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v369*int64(24)) {
		goto L44
	} else {
		goto L70
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = v369
	v377 = base.I32_wrap_i64(v369) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v377
	if v369 == int64(4294967296) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v386 = int32(-85899346)
	goto L73
L72:
	;
	v386 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v369), float64(0.9)))
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v386
	if v350 != int64(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v403 = v326
	goto L78
L75:
	;
	goto L76
L76:
	;
	F_pfree(m, v349)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L95
	}
L77:
	;
	v467 = v326
	v473 = v459
	goto L83
L78:
	;
	v410 = v349 + v403*int32(24)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+16)))
	if v411 != int32(1) {
		v459 = v403
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v459 = int32(0)
	goto L77
L80:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v410)))
	v415 = int64(23)
	v418 = int64(2388976653695081527)
	v419 = (int64(base.Ui64(v414)>>(uint(v415)%64)) ^ v414) * v418
	v420 = int64(47)
	v425 = int64(-8645972361240307355)
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v410)+8))
	v432 = (int64(base.Ui64(v427)>>(uint(v415)%64)) ^ v427) * v418
	v438 = ((v419^int64(base.Ui64(v419)>>(uint(v420)%64))^int64(-9208349263878056368))*v425 ^ int64(base.Ui64(v432)>>(uint(v420)%64)) ^ v432) * v425
	v443 = (int64(base.Ui64(v438)>>(uint(v415)%64)) ^ v438) * v418
	if v377&base.I32_wrap_i64(int64(base.Ui64(v443)>>(uint(v420)%64))^v443-int64(base.Ui64(v443)>>(uint(int64(32))%64))) == v403 {
		v459 = v403
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v454 = v403 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v454)) < base.Ui64(v350) {
		v403 = v454
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v480 = v349 + v473*int32(24)
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+16)))
	if v481 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L76
L85:
	;
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v480)))
	v485 = int64(23)
	v488 = int64(2388976653695081527)
	v489 = (int64(base.Ui64(v484)>>(uint(v485)%64)) ^ v484) * v488
	v490 = int64(47)
	v495 = int64(-8645972361240307355)
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v480)+8))
	v502 = (int64(base.Ui64(v497)>>(uint(v485)%64)) ^ v497) * v488
	v508 = ((v489^int64(base.Ui64(v489)>>(uint(v490)%64))^int64(-9208349263878056368))*v495 ^ int64(base.Ui64(v502)>>(uint(v490)%64)) ^ v502) * v495
	v513 = (int64(base.Ui64(v508)>>(uint(v485)%64)) ^ v508) * v488
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v530 = base.I32_wrap_i64(int64(base.Ui64(v513)>>(uint(v490)%64)) ^ v513 - int64(base.Ui64(v513)>>(uint(int64(32))%64)))
	goto L88
L86:
	;
	goto L87
L87:
	;
	v572 = v473 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v572)) < base.Ui64(v350) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v540 = v530 & v521
	v545 = v356 + v540*int32(24)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+16)))
	if v546 != 0 {
		v530 = v540 + int32(1)
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v547 = *(*int64)(unsafe.Add(mBase, uint32(v480)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v545)+16)) = v547
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v480)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v545)+8)) = v549
	v551 = *(*int64)(unsafe.Add(mBase, uint32(v480)))
	*(*int64)(unsafe.Add(mBase, uint32(v545))) = v551
	goto L87
L90:
	;
	goto L89
L91:
	;
	v576 = v572
	goto L93
L92:
	;
	v576 = int32(0)
	goto L93
L93:
	;
	v578 = v467 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v578)) < base.Ui64(v350) {
		v467 = v578
		v473 = v576
		goto L83
	} else {
		goto L94
	}
L94:
	;
	goto L84
L95:
	;
	goto L56
L96:
	;
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v21)+80))
	v620 = v604
	v621 = v607
	v623 = int32(0)
	goto L97
L97:
	;
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v621)+8))
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v21)+88))
	if v631^v612|(v633^v634) != int64(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v621)+20))
	if v812 == int32(0) {
		v863 = v621
		goto L49
	} else {
		goto L129
	}
L99:
	;
	v640 = v620 + int32(1)
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
	v642 = int64(23)
	v645 = int64(2388976653695081527)
	v646 = (int64(base.Ui64(v641)>>(uint(v642)%64)) ^ v641) * v645
	v647 = int64(47)
	v652 = int64(-8645972361240307355)
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v621)+8))
	v659 = (int64(base.Ui64(v654)>>(uint(v642)%64)) ^ v654) * v645
	v665 = ((v646^int64(base.Ui64(v646)>>(uint(v647)%64))^int64(-9208349263878056368))*v652 ^ int64(base.Ui64(v659)>>(uint(v647)%64)) ^ v659) * v652
	v670 = (int64(base.Ui64(v665)>>(uint(v642)%64)) ^ v665) * v645
	v678 = v603 & base.I32_wrap_i64(int64(base.Ui64(v670)>>(uint(v647)%64))^v670-int64(base.Ui64(v670)>>(uint(int64(32))%64)))
	if base.Ui32(v620) < base.Ui32(v678) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	goto L98
L102:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v682 = v620 + v680
	goto L104
L103:
	;
	v682 = v620
	goto L104
L104:
	;
	if base.Ui32(v682-v678) < base.Ui32(v623) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v685 = v640 & v603
	v688 = v602 + v685*int32(24)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+16)))
	if v689 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	v797 = v623 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v797) {
		goto L124
	} else {
		goto L125
	}
L108:
	;
	v696 = int32(0)
	v701 = v685
	goto L111
L109:
	;
	v736 = v688
	v737 = v685
	goto L110
L110:
	;
	if v620 != v737 {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	v710 = v696 + int32(1)
	if int32(151) <= v710 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v736 = v725
	v737 = v722
	goto L110
L113:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v715 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v713), base.F64_convert_i64_u(v715)), float64(0.1)) != 0 {
		goto L45
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v722 = (v701 + int32(1)) & v603
	v725 = v602 + v722*int32(24)
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725)+16)))
	if v726 != 0 {
		v696 = v710
		v701 = v722
		goto L111
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	goto L112
L118:
	;
	v755 = v736
	v756 = v737
	goto L121
L119:
	;
	goto L120
L120:
	;
	v837 = v621
	goto L50
L121:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v767 = v764 & (v756 - int32(1))
	v770 = v602 + v767*int32(24)
	v771 = *(*int64)(unsafe.Add(mBase, uint32(v770)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v755)+16)) = v771
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v770)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v755)+8)) = v773
	v775 = *(*int64)(unsafe.Add(mBase, uint32(v770)))
	*(*int64)(unsafe.Add(mBase, uint32(v755))) = v775
	if v620 != v767 {
		v755 = v770
		v756 = v767
		goto L121
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	goto L122
L124:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v800), base.F64_convert_i64_u(v802)), float64(0.1)) != 0 {
		goto L45
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v807 = v640 & v603
	v810 = v602 + v807*int32(24)
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810)+16)))
	if v811 != 0 {
		v620 = v807
		v621 = v810
		v623 = v797
		goto L97
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v837 = v810
	goto L50
L129:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v815 != 0 {
		v1143 = v812
		goto L47
	} else {
		goto L130
	}
L130:
	;
	v891 = v812
	goto L48
L131:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_get_entry_ref_4), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_pgstat_get_entry_ref_5), int32(630), int32(_a_F_pgstat_get_entry_ref_6))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v863)+20)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v876)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v876))) = int64(0)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v863)+20))
	v891 = v883
	goto L48
L135:
	;
	if v992 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L136:
	;
	if v907|base.B2i32(l3 == int32(0)) != 0 {
		v992 = v907
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	v916 = F_dshash_find_or_insert(m, v913, v905, v21+int32(80))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+80)))
	if v918 != 0 {
		v992 = v916
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v919 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v916)+20)) = v919
	v921 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v916)+16)) = uint8(v921)
	*(*int32)(unsafe.Add(mBase, uint32(v916)+24)) = v921
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	if base.Ui32(l0-v919) <= base.Ui32(int32(11)) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	v958 = F_dsa_allocate_extended(m, v926, v956, int32(6))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L147
	}
L141:
	;
	v955 = l0*int32(72) + int32(_a_F_pgstat_get_entry_ref_7)
	goto L140
L142:
	;
	goto L143
L143:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v953 = int32(0)
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v955 = v953
	goto L140
L145:
	;
	v941 = int32(0)
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[8]))
	if v943 == v941 {
		v953 = v941
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v943+l0<<(uint(int32(2))%32)-int32(96))))
	v953 = v951
	goto L144
L147:
	;
	if v958 == int32(0) {
		goto L46
	} else {
		goto L148
	}
L148:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	v964 = F_dsa_get_address(m, v963, v958)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v964))) = int32(-559038737)
	*(*int32)(unsafe.Add(mBase, uint32(v916)+28)) = v958
	v970 = v964 + int32(4)
	v971 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v970))) = uint16(v971)
	*(*int32)(unsafe.Add(mBase, uint32(v970)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v970)+8)) = int64(-1)
	goto L150
L150:
	;
	v979 = base.AtomicRmwAdd32(m, v916, int32(20), int32(1))
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v981, v916)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v916
	*(*int32)(unsafe.Add(mBase, uint32(v891)+4)) = v964
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v916)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v891)+8)) = v986
	if l4 == int32(0) {
		v1143 = v891
		goto L47
	} else {
		goto L152
	}
L152:
	;
	v990 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v990)
	v1143 = v891
	goto L47
L153:
	;
	v995 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v995
	v997 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v997
	F_pgstat_release_entry_ref(m, v21, v891, int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L4
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v1003 = int32(0)
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+16)))
	if base.B2i32(l3 == v1003)|base.B2i32(v1005&int32(1) == v1003) == v1003 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v1143 = int32(0)
	goto L47
L157:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v992)+28))
	v1016 = F_dsa_get_address(m, v1014, v1015)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	if v1005&int32(1) != 0 {
		goto L180
	} else {
		goto L181
	}
L160:
	;
	v1018 = int32(1)
	v1020 = base.AtomicRmwAdd32(m, v992, int32(20), v1018)
	v1023 = base.AtomicRmwAdd32(m, v992, int32(24), v1018)
	v1024 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v992)+16)) = uint8(v1024)
	if base.Ui32(l0-v1018) <= base.Ui32(int32(11)) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+16))
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		goto L169
	} else {
		goto L170
	}
L162:
	;
	v1054 = l0*int32(72) + int32(_a_F_pgstat_get_entry_ref_7)
	goto L161
L163:
	;
	goto L164
L164:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1052 = int32(0)
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1054 = v1052
	goto L161
L166:
	;
	v1040 = int32(0)
	v1042 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[8]))
	if v1042 == v1040 {
		v1052 = v1040
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1042+l0<<(uint(int32(2))%32)-int32(96))))
	v1052 = v1050
	goto L165
L168:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+20))
	if v1085 != 0 {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	v1084 = l0*int32(72) + int32(_a_F_pgstat_get_entry_ref_7)
	goto L168
L170:
	;
	goto L171
L171:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1082 = int32(0)
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v1084 = v1082
	goto L168
L173:
	;
	v1070 = int32(0)
	v1072 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[8]))
	if v1072 == v1070 {
		v1082 = v1070
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1072+l0<<(uint(int32(2))%32)-int32(96))))
	v1082 = v1080
	goto L172
L175:
	;
	base.MemoryFill(m, v1055+v1016, int32(0), v1085)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v1091 = base.AtomicRmwAdd32(m, v992, int32(20), int32(1))
	v1093 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v1093, v992)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v891)+4)) = v1016
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v992)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v891)+8)) = v1098
	if l4 == int32(0) {
		v1143 = v891
		goto L47
	} else {
		goto L179
	}
L179:
	;
	v1102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1102)
	v1143 = v891
	goto L47
L180:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v1107, v992)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L4
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v992)+28))
	v1123 = F_dsa_get_address(m, v1121, v1122)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L4
	} else {
		goto L185
	}
L183:
	;
	v1110 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v1110
	v1112 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v1112
	F_pgstat_release_entry_ref(m, v21+int32(16), v891, int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v1143 = int32(0)
	goto L47
L185:
	;
	v1127 = base.AtomicRmwAdd32(m, v992, int32(20), int32(1))
	v1129 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v1129, v992)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v891)+4)) = v1123
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v992)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v891)+8)) = v1134
	v1143 = v891
	goto L47
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	F_errcode(m, int32(_a_F_pgstat_get_entry_ref_8))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_pgstat_get_entry_ref_9), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v1173 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v1173
	v1175 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v1175
	F_errdetail(m, int32(_a_F_pgstat_get_entry_ref_10), v21+int32(32))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_pgstat_get_entry_ref_11), int32(537), int32(_a_F_pgstat_get_entry_ref_12))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_get_entry_ref_13), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_pgstat_get_entry_ref_5), int32(327), int32(_a_F_pgstat_get_entry_ref_14))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_get_slru_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(7)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_slru_name[0])))
		v8 = v6
	} else {
		v8 = int32(0)
	}
	return v8
}
func F_pgstat_get_transactional_drops(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_transactional_drops[0]))
	if v9 == v3 {
		v65 = v3
		return v65
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		v15 = F_palloc(m, v12<<(uint(int32(4))%32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v20 == int32(0) {
				v65 = v3
			} else {
				v24 = v9 + int32(8)
				if v20 == v24 {
					v65 = v3
				} else {
					v29 = v20
					v30 = v3
					for {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29-int32(4)))))
						if l0 != 0 {
							if v35&int32(1) == int32(0) {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v47 = v44 + v30<<(uint(int32(4))%32)
								v49 = v29 - int32(20)
								v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v50
								v52 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v52
								v57 = v30 + int32(1)
							} else {
								v57 = v30
							}
						} else {
							if v35&int32(1) == int32(0) {
								v57 = v30
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v47 = v44 + v30<<(uint(int32(4))%32)
								v49 = v29 - int32(20)
								v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v50
								v52 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v52
								v57 = v30 + int32(1)
							}
						}
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						if v59 != v24 {
							v29 = v59
							v30 = v57
							continue
						} else {
							break
						}
						break
					}
					v65 = v57
				}
			}
			return v65
		}
	}
}
func F_pgstat_get_wait_event(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v124 = v2
		m.G0 = v6 + int32(16)
		return v124
	} else {
		v10 = int32(_a_F_pgstat_get_wait_event_0)
		switch int32(base.Ui32(l0-int32(16777216)) >> (uint(int32(24)) % 32)) {
		case 0:
			v16 = l0 & int32(_a_F_pgstat_get_wait_event_1)
			if base.Ui32(v16) <= base.Ui32(int32(94)) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[0])))
				v41 = v21
			} else {
				v25 = (v16 - int32(95)) & int32(_a_F_pgstat_get_wait_event_1)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_wait_event[1]))
				if v25 < v27 {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_wait_event[2]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v25<<(uint(int32(2))%32))))
					if v34 != 0 {
						v36 = v34
					} else {
						v36 = int32(_a_F_pgstat_get_wait_event_2)
					}
					v39 = v36
				} else {
					v39 = int32(_a_F_pgstat_get_wait_event_2)
				}
				v41 = v39
			}
			v124 = v41
			m.G0 = v6 + int32(16)
			return v124
		default:
			v124 = v10
			m.G0 = v6 + int32(16)
			return v124
		case 2:
			v43 = l0 & int32(_a_F_pgstat_get_wait_event_1)
			if base.Ui32(v43) <= base.Ui32(int32(11)) {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v43<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[3])))
				v50 = v48
			} else {
				v50 = int32(_a_F_pgstat_get_wait_event_3)
			}
			v124 = v50
			m.G0 = v6 + int32(16)
			return v124
		case 3:
			if l0 == int32(67108864) {
				v87 = int32(_a_F_pgstat_get_wait_event_4)
			} else {
				v87 = int32(_a_F_pgstat_get_wait_event_0)
			}
			v124 = v87
			m.G0 = v6 + int32(16)
			return v124
		case 4:
			v89 = l0 - int32(83886080)
			if base.Ui32(int32(18)) <= base.Ui32(v89) {
				v124 = v10
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[4])))
				v124 = v94
			}
			m.G0 = v6 + int32(16)
			return v124
		case 5:
			v96 = l0 - int32(100663296)
			if base.Ui32(int32(9)) <= base.Ui32(v96) {
				v124 = v10
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v96<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[5])))
				v124 = v101
			}
			m.G0 = v6 + int32(16)
			return v124
		case 6, 10:
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
			if l0 == int32(117440512) {
				v124 = int32(_a_F_pgstat_get_wait_event_5)
				m.G0 = v6 + int32(16)
				return v124
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_wait_event[6]))
				v60 = F_LWLockAcquire(m, v56+int32(_a_F_pgstat_get_wait_event_6), int32(1))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_wait_event[7]))
					v71 = F_hash_search(m, v65, v6+int32(12), int32(0), v6+int32(11))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_wait_event[6]))
						F_LWLockRelease(m, v74+int32(_a_F_pgstat_get_wait_event_6))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							if v71 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									v133 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v133
									F_errmsg_internal(m, int32(_a_F_pgstat_get_wait_event_7), v6)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pgstat_get_wait_event_8), int32(295), int32(_a_F_pgstat_get_wait_event_9))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v124 = v71 + int32(4)
								m.G0 = v6 + int32(16)
								return v124
							}
						}
					}
				}
			}
		case 7:
			v103 = l0 + int32(-134217728)
			if base.Ui32(int32(57)) <= base.Ui32(v103) {
				v124 = v10
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[8])))
				v124 = v108
			}
			m.G0 = v6 + int32(16)
			return v124
		case 8:
			v110 = l0 - int32(150994944)
			if base.Ui32(int32(10)) <= base.Ui32(v110) {
				v124 = v10
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v110<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[9])))
				v124 = v115
			}
			m.G0 = v6 + int32(16)
			return v124
		case 9:
			v117 = l0 - int32(167772160)
			if base.Ui32(int32(81)) <= base.Ui32(v117) {
				v124 = v10
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v117<<(uint(int32(2))%32))+uint32(_c_F_pgstat_get_wait_event[10])))
				v124 = v122
			}
			m.G0 = v6 + int32(16)
			return v124
		}
	}
}
func F_pgstat_io_snapshot_cb(m *base.Module) {
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
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_snapshot_cb[0]))
	v8 = v6 + int32(608)
	v10 = F_LWLockAcquire(m, v8, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_snapshot_cb[0]))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+896))
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_io_snapshot_cb[1])) = v15
	base.MemoryCopy(m, int32(_a_F_pgstat_io_snapshot_cb_0), v6+int32(904), int32(2880))
	F_LWLockRelease(m, v8)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(1)
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_io_snapshot_cb[0]))
	v35 = v30 + v25<<(uint(int32(4))%32) + int32(608)
	v37 = F_LWLockAcquire(m, v35, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v39 = int32(2880)
	v40 = v25 * v39
	base.MemoryCopy(m, v40+int32(_a_F_pgstat_io_snapshot_cb_0), v30+v40+int32(904), v39)
	F_LWLockRelease(m, v35)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v51 = v25 + int32(1)
	if v51 != int32(18) {
		v25 = v51
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pgstat_progress_update_param(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_param[0]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_progress_update_param[1])))
		if v9&int32(1) == int32(0) {
		} else {
			v14 = int32(_a_F_pgstat_progress_update_param_0)
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_param[2]))
			v17 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_param[2])) = v16 + v17
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20 + v17
			*(*int64)(unsafe.Add(mBase, uint32(v5+l0<<(uint(int32(3))%32))+232)) = l1
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v28 + v17
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_param[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_update_param[2])) = v34 - v17
		}
	}
	return
}
func F_pgstat_report_archiver(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_archiver[0]))
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	F_gettimeofday(m, v13)
	mBase = m.M
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+8)))
	m.G0 = v13 + v12
	v26 = int32(_a_F_pgstat_report_archiver_0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_archiver[1]))
	v29 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_archiver[1])) = v28 + v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v32 + v29
	if l1 != 0 {
		v38 = int32(112)
	} else {
		v38 = int32(48)
	}
	v39 = v7 + v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v40 + int64(1)
	if l1 != 0 {
		v46 = int32(120)
	} else {
		v46 = int32(56)
	}
	v47 = v7 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+40)) = uint8(v48)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v58
	if l1 != 0 {
		v62 = int32(168)
	} else {
		v62 = int32(104)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7+v62))) = v17 + v16*int64(1000000) - int64(946684800000000)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v65 + v66
	v69 = int32(_a_F_pgstat_report_archiver_0)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_archiver[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_archiver[1])) = v71 - v66
	return
}
func F_pgstat_report_subscription_error(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	v4 = int32(0)
	v7 = F_pgstat_prep_pending_entry(m, int32(5), v4, base.I64_extend_i32_u(l0), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if l1 != 0 {
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v10 + int64(1)
			return
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v14 + int64(1)
			return
		}
	}
}
func F_pgstat_report_tempfile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_tempfile[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_tempfile[1]))
		v12 = F_pgstat_prep_pending_entry(m, int32(1), v9, int64(0), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v15 + base.I64_extend_i32_u(l0)
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v19 + int64(1)
			return
		}
	} else {
		return
	}
}
func F_pgstat_snapshot_fixed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[0])))
	if v7 != 0 {
		v9 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[1])) = v9
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[2])) = v9
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[3])) = v9
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[4])) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[5])) = v18
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[6])) = v18
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[7]))
		if v27 != 0 {
			F_MemoryContextDelete(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[7])) = int32(0)
				F_pgstat_clear_backend_activity_snapshot(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[0])) = uint8(v36)
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[8]))
					if v40 == int32(2) {
						F_pgstat_build_snapshot(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							return
						}
					} else {
						if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
							v73 = l0
							v74 = l0*int32(72) + int32(_a_F_pgstat_snapshot_fixed_0)
							v76 = int32(_a_F_pgstat_snapshot_fixed_1)
						} else {
							v57 = l0 - int32(24)
							if base.Ui32(int32(8)) < base.Ui32(v57) {
								v72 = int32(0)
							} else {
								v60 = int32(0)
								v62 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[9]))
								if v62 == v60 {
									v72 = v60
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v62+l0<<(uint(int32(2))%32)-int32(96))))
									v72 = v70
								}
							}
							v73 = v57
							v74 = v72
							v76 = int32(_a_F_pgstat_snapshot_fixed_2)
						}
						v77 = v73 + v76
						if v40 == int32(0) {
							v80 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
							m.T0[v83].(func(*base.Module))(m)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
								return
							}
						} else {
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
							if v82 != 0 {
								return
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
								m.T0[v83].(func(*base.Module))(m)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
									return
								}
							}
						}
					}
				}
			}
		} else {
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[0])) = uint8(v36)
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[8]))
				if v40 == int32(2) {
					F_pgstat_build_snapshot(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						return
					}
				} else {
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v73 = l0
						v74 = l0*int32(72) + int32(_a_F_pgstat_snapshot_fixed_0)
						v76 = int32(_a_F_pgstat_snapshot_fixed_1)
					} else {
						v57 = l0 - int32(24)
						if base.Ui32(int32(8)) < base.Ui32(v57) {
							v72 = int32(0)
						} else {
							v60 = int32(0)
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[9]))
							if v62 == v60 {
								v72 = v60
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v62+l0<<(uint(int32(2))%32)-int32(96))))
								v72 = v70
							}
						}
						v73 = v57
						v74 = v72
						v76 = int32(_a_F_pgstat_snapshot_fixed_2)
					}
					v77 = v73 + v76
					if v40 == int32(0) {
						v80 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
						m.T0[v83].(func(*base.Module))(m)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v86 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
							return
						}
					} else {
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
						if v82 != 0 {
							return
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
							m.T0[v83].(func(*base.Module))(m)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[8]))
		if v40 == int32(2) {
			F_pgstat_build_snapshot(m)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				return
			}
		} else {
			if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
				v73 = l0
				v74 = l0*int32(72) + int32(_a_F_pgstat_snapshot_fixed_0)
				v76 = int32(_a_F_pgstat_snapshot_fixed_1)
			} else {
				v57 = l0 - int32(24)
				if base.Ui32(int32(8)) < base.Ui32(v57) {
					v72 = int32(0)
				} else {
					v60 = int32(0)
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_snapshot_fixed[9]))
					if v62 == v60 {
						v72 = v60
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v62+l0<<(uint(int32(2))%32)-int32(96))))
						v72 = v70
					}
				}
				v73 = v57
				v74 = v72
				v76 = int32(_a_F_pgstat_snapshot_fixed_2)
			}
			v77 = v73 + v76
			if v40 == int32(0) {
				v80 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
				m.T0[v83].(func(*base.Module))(m)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					v86 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
					return
				}
			} else {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
				if v82 != 0 {
					return
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
					m.T0[v83].(func(*base.Module))(m)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
						return
					}
				}
			}
		}
	}
}
func F_pgstat_twophase_postcommit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_twophase_postcommit[0]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+52)))
	if v11 != 0 {
		v12 = int32(0)
	} else {
		v12 = v10
	}
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v16 = F_pgstat_prep_pending_entry(m, int32(2), v12, base.I64_extend_i32_u(v13), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v11)
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v13
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v21 + v22
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v25 + v26
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v29 + v30
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+53)))
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v33)
		if v33 == int32(0) {
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+88))
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v18)+96))
			v47 = v37
			v48 = v38
		} else {
			v40 = v18 + int32(88)
			v41 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = v41
			*(*int64)(unsafe.Add(mBase, uint32(v40))) = v41
			v47 = int64(0)
			v48 = v41
		}
		v49 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v49 - v50 + v47
		v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		v55 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v54 + v55 + v48
		v59 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
		v60 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		v61 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v62 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v59 + (v60 + (v61 + v62))
		return
	}
}
func F_pgstat_wal_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_wal_snapshot_cb[0]))
	v6 = v4 + int32(_a_F_pgstat_wal_snapshot_cb_0)
	v8 = F_LWLockAcquire(m, v6, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_pgstat_wal_snapshot_cb[1])))
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_snapshot_cb[2])) = v11
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_pgstat_wal_snapshot_cb[3])))
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_snapshot_cb[4])) = v14
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_pgstat_wal_snapshot_cb[5])))
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_snapshot_cb[6])) = v17
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_pgstat_wal_snapshot_cb[7])))
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_snapshot_cb[8])) = v20
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_pgstat_wal_snapshot_cb[9])))
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_wal_snapshot_cb[10])) = v23
		F_LWLockRelease(m, v6)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	}
}
