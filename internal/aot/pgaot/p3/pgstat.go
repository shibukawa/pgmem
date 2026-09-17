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
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
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
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v270 int64
	_ = v270
	var v276 int64
	_ = v276
	var v282 int64
	_ = v282
	var v287 int64
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v339 int64
	_ = v339
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int64
	_ = v354
	var v364 int64
	_ = v364
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v413 int64
	_ = v413
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v427 int64
	_ = v427
	var v433 int64
	_ = v433
	var v438 int64
	_ = v438
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v480 int64
	_ = v480
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v497 int64
	_ = v497
	var v503 int64
	_ = v503
	var v508 int64
	_ = v508
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int64
	_ = v607
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v626 int64
	_ = v626
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v637 int64
	_ = v637
	var v640 int64
	_ = v640
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v647 int64
	_ = v647
	var v649 int64
	_ = v649
	var v654 int64
	_ = v654
	var v660 int64
	_ = v660
	var v665 int64
	_ = v665
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int64
	_ = v710
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int64
	_ = v766
	var v768 int64
	_ = v768
	var v770 int64
	_ = v770
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v797 int64
	_ = v797
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v858 int32
	_ = v858
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v886 int32
	_ = v886
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int64
	_ = v991
	var v993 int64
	_ = v993
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int64
	_ = v1109
	var v1111 int64
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
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
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[4]))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+16)) = v80
	*(*uint32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[5])) = uint32(v80)
	goto L12
L15:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v87)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v89 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	if v91 == v89 {
		v239 = v89
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v257
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v256
	v260 = int64(23)
	v263 = int64(2388976653695081527)
	v264 = (v256 ^ int64(base.Ui64(v256)>>(uint(v260)%64))) * v263
	v265 = int64(47)
	v270 = int64(-8645972361240307355)
	v276 = (v257 ^ int64(base.Ui64(v257)>>(uint(v260)%64))) * v263
	v282 = ((v264^int64(base.Ui64(v264)>>(uint(v265)%64))^int64(-9208349263878056368))*v270 ^ int64(base.Ui64(v276)>>(uint(v265)%64)) ^ v276) * v270
	v287 = (int64(base.Ui64(v282)>>(uint(v260)%64)) ^ v282) * v263
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	v306 = base.B2i32(base.Ui32(v295) < base.Ui32(v296))
	goto L42
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[4]))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v95)+16)) = v96
	v99 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[5])))
	if v99 == v96 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	v239 = v102
	goto L18
L21:
	;
	goto L22
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[4]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v104)+16)) = v105
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	if v110 == int64(0) {
		v149 = int32(-1)
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v163 = v109
	v169 = v6
	v170 = v149
	goto L29
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	v120 = int32(0)
	goto L25
L25:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v120*int32(24))+16)))
	if v136 != int32(1) {
		v149 = v120
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v149 = int32(-1)
	goto L23
L27:
	;
	v140 = v120 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v140)) < base.Ui64(v110) {
		v120 = v140
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v187 = v169
	v188 = v170
	v190 = v169
	goto L32
L30:
	;
	*(*uint32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[5])) = uint32(v105)
	v239 = v163
	goto L18
L31:
	;
	goto L30
L32:
	;
	if v190&int32(1) != 0 {
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+16)))
	if v218 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v201 = int32(1)
	v202 = v188 - v201
	v206 = base.B2i32(v200&(v202^v149) == int32(0))
	v207 = v206 | v187
	v210 = v200 & v202
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	v212 = v188*int32(24) + v211
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+16)))
	if v213 != v201 {
		v187 = v207
		v188 = v210
		v190 = v206
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if v221 == v222 {
		v169 = v207
		v170 = v210
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v224 != 0 {
		v169 = v207
		v170 = v210
		goto L29
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v212)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v225
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v212)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v227
	F_pgstat_release_entry_ref(m, v21+int32(48), v216, int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[3]))
	v163 = v235
	v169 = v207
	v170 = v210
	goto L29
L42:
	;
	if v306 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = v1205
	v306 = v1205
	goto L42
L46:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_delete_entry(m, v1159, v911)
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
	v898 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	v900 = v21 - int32(-64)
	v902 = F_dshash_find(m, v898, v900, int32(0))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L136
	}
L49:
	;
	v869 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[0]))
	v871 = F_MemoryContextAlloc(m, v869, int32(24))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L4
	} else {
		goto L134
	}
L50:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v843 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+8)) = v842 + v843
	*(*uint8)(unsafe.Add(mBase, uint32(v832)+16)) = uint8(v843)
	*(*int64)(unsafe.Add(mBase, uint32(v832)+8)) = v257
	*(*int64)(unsafe.Add(mBase, uint32(v832))) = v256
	v858 = v832
	goto L49
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L4
	} else {
		goto L131
	}
L52:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v239)))
	if v318 == int64(4294967296) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v599 = v598 & base.I32_wrap_i64(int64(base.Ui64(v287)>>(uint(v265)%64))^v287-int64(base.Ui64(v287)>>(uint(int64(32))%64)))
	v602 = v597 + v599*int32(24)
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+16)))
	if v603 == int32(0) {
		v832 = v602
		goto L50
	} else {
		goto L96
	}
L55:
	;
	v321 = int32(0)
	v323 = int64(2)
	v325 = v318 << (uint(int64(1)) % 64)
	if base.Ui64(v325) <= base.Ui64(v323) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v306 = int32(1)
	goto L42
L57:
	;
	v328 = v323
	goto L59
L58:
	;
	v328 = v325
	goto L59
L59:
	;
	v329 = int64(1)
	if v328&(v328-v329) == int64(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v339 = v328
	goto L62
L61:
	;
	v339 = v329 << (uint(int64(64)-base.I64_clz(v328)) % 64)
	goto L62
L62:
	;
	if base.Ui64(v339*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v239)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v239)+24))
	v351 = F_MemoryContextAllocExtended(m, v346, base.I32_wrap_i64(v339)*int32(24), int32(5))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = v351
	v354 = int64(1)
	if v339&(v339-v354) == int64(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v364 = v339
	goto L69
L68:
	;
	v364 = v354 << (uint(int64(64)-base.I64_clz(v339)) % 64)
	goto L69
L69:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v364*int64(24)) {
		goto L44
	} else {
		goto L70
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v364
	v372 = base.I32_wrap_i64(v364) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = v372
	if v364 == int64(4294967296) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v381 = int32(-85899346)
	goto L73
L72:
	;
	v381 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v364), float64(0.9)))
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = v381
	if v345 != int64(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v398 = v321
	goto L78
L75:
	;
	goto L76
L76:
	;
	F_pfree(m, v344)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L95
	}
L77:
	;
	v462 = v321
	v468 = v454
	goto L83
L78:
	;
	v405 = v344 + v398*int32(24)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+16)))
	if v406 != int32(1) {
		v454 = v398
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v454 = int32(0)
	goto L77
L80:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v405)))
	v410 = int64(23)
	v413 = int64(2388976653695081527)
	v414 = (int64(base.Ui64(v409)>>(uint(v410)%64)) ^ v409) * v413
	v415 = int64(47)
	v420 = int64(-8645972361240307355)
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v405)+8))
	v427 = (int64(base.Ui64(v422)>>(uint(v410)%64)) ^ v422) * v413
	v433 = ((v414^int64(base.Ui64(v414)>>(uint(v415)%64))^int64(-9208349263878056368))*v420 ^ int64(base.Ui64(v427)>>(uint(v415)%64)) ^ v427) * v420
	v438 = (int64(base.Ui64(v433)>>(uint(v410)%64)) ^ v433) * v413
	if v372&base.I32_wrap_i64(int64(base.Ui64(v438)>>(uint(v415)%64))^v438-int64(base.Ui64(v438)>>(uint(int64(32))%64))) == v398 {
		v454 = v398
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v449 = v398 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v449)) < base.Ui64(v345) {
		v398 = v449
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v475 = v344 + v468*int32(24)
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+16)))
	if v476 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L76
L85:
	;
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v475)))
	v480 = int64(23)
	v483 = int64(2388976653695081527)
	v484 = (int64(base.Ui64(v479)>>(uint(v480)%64)) ^ v479) * v483
	v485 = int64(47)
	v490 = int64(-8645972361240307355)
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v475)+8))
	v497 = (int64(base.Ui64(v492)>>(uint(v480)%64)) ^ v492) * v483
	v503 = ((v484^int64(base.Ui64(v484)>>(uint(v485)%64))^int64(-9208349263878056368))*v490 ^ int64(base.Ui64(v497)>>(uint(v485)%64)) ^ v497) * v490
	v508 = (int64(base.Ui64(v503)>>(uint(v480)%64)) ^ v503) * v483
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v525 = base.I32_wrap_i64(int64(base.Ui64(v508)>>(uint(v485)%64)) ^ v508 - int64(base.Ui64(v508)>>(uint(int64(32))%64)))
	goto L88
L86:
	;
	goto L87
L87:
	;
	v567 = v468 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v567)) < base.Ui64(v345) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v535 = v525 & v516
	v540 = v351 + v535*int32(24)
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+16)))
	if v541 != 0 {
		v525 = v535 + int32(1)
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v475)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v540)+16)) = v542
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v475)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v540)+8)) = v544
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v475)))
	*(*int64)(unsafe.Add(mBase, uint32(v540))) = v546
	goto L87
L90:
	;
	goto L89
L91:
	;
	v571 = v567
	goto L93
L92:
	;
	v571 = int32(0)
	goto L93
L93:
	;
	v573 = v462 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v573)) < base.Ui64(v345) {
		v462 = v573
		v468 = v571
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
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v21)+80))
	v615 = v599
	v616 = v602
	v618 = int32(0)
	goto L97
L97:
	;
	v626 = *(*int64)(unsafe.Add(mBase, uint32(v616)))
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v616)+8))
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v21)+88))
	if v626^v607|(v628^v629) != int64(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v616)+20))
	if v807 == int32(0) {
		v858 = v616
		goto L49
	} else {
		goto L129
	}
L99:
	;
	v635 = v615 + int32(1)
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v616)))
	v637 = int64(23)
	v640 = int64(2388976653695081527)
	v641 = (int64(base.Ui64(v636)>>(uint(v637)%64)) ^ v636) * v640
	v642 = int64(47)
	v647 = int64(-8645972361240307355)
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v616)+8))
	v654 = (int64(base.Ui64(v649)>>(uint(v637)%64)) ^ v649) * v640
	v660 = ((v641^int64(base.Ui64(v641)>>(uint(v642)%64))^int64(-9208349263878056368))*v647 ^ int64(base.Ui64(v654)>>(uint(v642)%64)) ^ v654) * v647
	v665 = (int64(base.Ui64(v660)>>(uint(v637)%64)) ^ v660) * v640
	v673 = v598 & base.I32_wrap_i64(int64(base.Ui64(v665)>>(uint(v642)%64))^v665-int64(base.Ui64(v665)>>(uint(int64(32))%64)))
	if base.Ui32(v615) < base.Ui32(v673) {
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
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v677 = v615 + v675
	goto L104
L103:
	;
	v677 = v615
	goto L104
L104:
	;
	if base.Ui32(v677-v673) < base.Ui32(v618) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v680 = v635 & v598
	v683 = v597 + v680*int32(24)
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+16)))
	if v684 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	v792 = v618 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v792) {
		goto L124
	} else {
		goto L125
	}
L108:
	;
	v691 = int32(0)
	v696 = v680
	goto L111
L109:
	;
	v731 = v683
	v732 = v680
	goto L110
L110:
	;
	if v615 != v732 {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	v705 = v691 + int32(1)
	if int32(151) <= v705 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v731 = v720
	v732 = v717
	goto L110
L113:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v239)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v708), base.F64_convert_i64_u(v710)), float64(0.1)) != 0 {
		goto L45
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v717 = (v696 + int32(1)) & v598
	v720 = v597 + v717*int32(24)
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720)+16)))
	if v721 != 0 {
		v691 = v705
		v696 = v717
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
	v750 = v731
	v751 = v732
	goto L121
L119:
	;
	goto L120
L120:
	;
	v832 = v616
	goto L50
L121:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v762 = v759 & (v751 - int32(1))
	v765 = v597 + v762*int32(24)
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v765)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v750)+16)) = v766
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v765)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v750)+8)) = v768
	v770 = *(*int64)(unsafe.Add(mBase, uint32(v765)))
	*(*int64)(unsafe.Add(mBase, uint32(v750))) = v770
	if v615 != v762 {
		v750 = v765
		v751 = v762
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
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v239)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v795), base.F64_convert_i64_u(v797)), float64(0.1)) != 0 {
		goto L45
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v802 = v635 & v598
	v805 = v597 + v802*int32(24)
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805)+16)))
	if v806 != 0 {
		v615 = v802
		v616 = v805
		v618 = v792
		goto L97
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v832 = v805
	goto L50
L129:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v807)+4))
	if v810 != 0 {
		v1143 = v807
		goto L47
	} else {
		goto L130
	}
L130:
	;
	v886 = v807
	goto L48
L131:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_get_entry_ref_4), int32(0))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_pgstat_get_entry_ref_5), int32(630), int32(_a_F_pgstat_get_entry_ref_6))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v858)+20)) = v871
	*(*int32)(unsafe.Add(mBase, uint32(v871)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v871))) = int64(0)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v858)+20))
	v886 = v878
	goto L48
L135:
	;
	if v988 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L136:
	;
	if v902|base.B2i32(l3 == int32(0)) != 0 {
		v988 = v902
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	v911 = F_dshash_find_or_insert(m, v908, v900, v21+int32(80))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+80)))
	if v913 != 0 {
		v988 = v911
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v914 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v911)+20)) = v914
	v916 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v911)+16)) = uint8(v916)
	*(*int32)(unsafe.Add(mBase, uint32(v911)+24)) = v916
	v921 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	if base.Ui32(l0-v914) <= base.Ui32(int32(11)) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	v953 = F_dsa_allocate_extended(m, v921, v951, int32(6))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L4
	} else {
		goto L147
	}
L141:
	;
	v950 = l0*int32(72) + int32(_a_F_pgstat_get_entry_ref_7)
	goto L140
L142:
	;
	goto L143
L143:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v948 = int32(0)
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v950 = v948
	goto L140
L145:
	;
	v936 = int32(0)
	v938 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[8]))
	if v938 == v936 {
		v948 = v936
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v938+l0<<(uint(int32(2))%32)-int32(96))))
	v948 = v946
	goto L144
L147:
	;
	if v953 == int32(0) {
		goto L46
	} else {
		goto L148
	}
L148:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	v959 = F_dsa_get_address(m, v958, v953)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v959))) = int32(-559038737)
	*(*int32)(unsafe.Add(mBase, uint32(v911)+28)) = v953
	v965 = v959 + int32(4)
	v966 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v965))) = uint16(v966)
	*(*int32)(unsafe.Add(mBase, uint32(v965)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v965)+8)) = int64(-1)
	goto L150
L150:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v911)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v911)+20)) = v972 + int32(1)
	v977 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v977, v911)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v886))) = v911
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = v959
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v911)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v886)+8)) = v982
	if l4 == int32(0) {
		v1143 = v886
		goto L47
	} else {
		goto L152
	}
L152:
	;
	v986 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v986)
	v1143 = v886
	goto L47
L153:
	;
	v991 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v991
	v993 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v993
	F_pgstat_release_entry_ref(m, v21, v886, int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L4
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v999 = int32(0)
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988)+16)))
	if base.B2i32(l3 == v999)|base.B2i32(v1001&int32(1) == v999) == v999 {
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
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v988)+28))
	v1012 = F_dsa_get_address(m, v1010, v1011)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	if v1001&int32(1) != 0 {
		goto L180
	} else {
		goto L181
	}
L160:
	;
	v1014 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v988)+16)) = uint8(v1014)
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	v1017 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v988)+20)) = v1016 + v1017
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v988)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+24)) = v1020 + v1017
	if base.Ui32(l0-v1017) <= base.Ui32(int32(11)) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+16))
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		goto L169
	} else {
		goto L170
	}
L162:
	;
	v1052 = l0*int32(72) + int32(_a_F_pgstat_get_entry_ref_7)
	goto L161
L163:
	;
	goto L164
L164:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1050 = int32(0)
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1052 = v1050
	goto L161
L166:
	;
	v1038 = int32(0)
	v1040 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[8]))
	if v1040 == v1038 {
		v1050 = v1038
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1040+l0<<(uint(int32(2))%32)-int32(96))))
	v1050 = v1048
	goto L165
L168:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+20))
	if v1083 != 0 {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	v1082 = l0*int32(72) + int32(_a_F_pgstat_get_entry_ref_7)
	goto L168
L170:
	;
	goto L171
L171:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1080 = int32(0)
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v1082 = v1080
	goto L168
L173:
	;
	v1068 = int32(0)
	v1070 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[8]))
	if v1070 == v1068 {
		v1080 = v1068
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1070+l0<<(uint(int32(2))%32)-int32(96))))
	v1080 = v1078
	goto L172
L175:
	;
	base.MemoryFill(m, v1053+v1012, int32(0), v1083)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+20)) = v1087 + int32(1)
	v1092 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v1092, v988)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v886))) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = v1012
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v988)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v886)+8)) = v1097
	if l4 == int32(0) {
		v1143 = v886
		goto L47
	} else {
		goto L179
	}
L179:
	;
	v1101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1101)
	v1143 = v886
	goto L47
L180:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v1106, v988)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L4
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[7]))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v988)+28))
	v1122 = F_dsa_get_address(m, v1120, v1121)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L4
	} else {
		goto L185
	}
L183:
	;
	v1109 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v1109
	v1111 = *(*int64)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v1111
	F_pgstat_release_entry_ref(m, v21+int32(16), v886, int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
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
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+20)) = v1124 + int32(1)
	v1129 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_entry_ref[6]))
	F_dshash_release_lock(m, v1129, v988)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v886))) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = v1122
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v988)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v886)+8)) = v1134
	v1143 = v886
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
