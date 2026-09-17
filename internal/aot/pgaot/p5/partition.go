package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreatePartitionDirectory(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = int32(_a_F_CreatePartitionDirectory_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CreatePartitionDirectory[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreatePartitionDirectory[0])) = l0
	v15 = F_palloc(m, int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(51539607556)
		v26 = F_hash_create(m, int32(_a_F_CreatePartitionDirectory_1), int32(256), v8, int32(1064))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v26
			*(*int32)(unsafe.Add(mBase, _c_F_CreatePartitionDirectory[0])) = v11
			m.G0 = v8 + int32(48)
			return v15
		}
	}
}
func F_DestroyPartitionDirectory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = v6 + int32(12)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_hash_seq_init(m, v9, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = F_hash_seq_search(m, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = v13
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v6 + int32(32)
	return
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_RelationDecrementReferenceCount(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v23 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v23 != 0 {
		v15 = v23
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}
func F_ExecInitPartitionDispatchInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitPartitionDispatchInfo[0]))
	v20 = F_CreatePartitionDirectory(m, v15, base.B2i32(v17 < int32(2)))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v25 = v12
	goto L3
L3:
	;
	v26 = int32(_a_F_ExecInitPartitionDispatchInfo_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitPartitionDispatchInfo[1]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitPartitionDispatchInfo[1])) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+56))
	if v32 != l2 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v20
	v25 = v20
	goto L3
L6:
	;
	v35 = F_table_open(m, l2, int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v38 = v31
	v39 = v25
	goto L8
L8:
	;
	v40 = F_PartitionDirectoryLookup(m, v39, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v38 = v35
	v39 = v37
	goto L8
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v47 = F_palloc(m, v42<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v38
	v50 = F_RelationGetPartitionKey(m, v38)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v50
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v72
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v78 = v76 << (uint(int32(2)) % 32)
	if v78 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v61 = F_build_attrmap_by_name_if_req(m, v58, v59, v56)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v69 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v69
	v72 = v69
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v61
	if v61 == int32(0) {
		v72 = v56
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v67 = F_MakeTupleTableSlot(m, v59, int32(_a_F_ExecInitPartitionDispatchInfo_1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v72 = v67
	goto L13
L20:
	;
	base.MemoryFill(m, v47+int32(24), int32(255), v78)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v85 = v83 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v87 <= v85 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v87 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v120 = v83 << (uint(int32(2)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120+v121))) = v47
	if l3 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v117
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(4)
	v94 = F_palloc(m, int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v87 << (uint(int32(1)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v108 = F_repalloc(m, v105, v87<<(uint(int32(3))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v94
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v100 = F_palloc(m, v97<<(uint(int32(2))%32))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v117 = v100
	goto L26
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v115 = F_repalloc(m, v111, v112<<(uint(int32(2))%32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v117 = v115
	goto L26
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitPartitionDispatchInfo[1])) = v27
	return v47
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v120))) = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v131 = F_palloc0(m, int32(216))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(388)
	v135 = int32(0)
	F_InitResultRelInfo(m, v131, v38, v135, l5, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v139+v120))) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l3+l4<<(uint(int32(2))%32))+24)) = v83
	goto L34
}
func F_ExecInitPartitionExecPruning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v470 int32
	_ = v470
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v22 = l2 << (uint(int32(2)) % 32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22+v25)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if base.B2i32(l3 == v6)|base.B2i32(v28 == v6) != 0 {
		v75 = base.B2i32(l3|v28 == v6)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v75 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	goto L1
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v43 != v44 {
		v75 = int32(0)
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v46 = int32(1)
	if v43 <= v46 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v49 = v46
	goto L7
L6:
	;
	v49 = v43
	goto L7
L7:
	;
	v50 = int32(8)
	v55 = int32(0)
	goto L8
L8:
	;
	v63 = v55 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3+v50+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v28+v50+v63)))
	v68 = base.B2i32(v65 == v67)
	if v65 != v67 {
		v75 = v68
		goto L2
	} else {
		goto L10
	}
L9:
	;
	v75 = v68
	goto L2
L10:
	;
	v71 = v55 + int32(1)
	if v71 != v49 {
		v55 = v71
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81+v22)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+16)))
	if v84 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L19
	} else {
		goto L131
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v99
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+17)))
	if v101 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v22+v88)))
	v99 = v90
	goto L15
L17:
	;
	goto L18
L18:
	;
	v91 = int32(0)
	v95 = F_bms_add_range(m, v91, v91, l1-int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v99 = v95
	goto L15
L21:
	;
	m.G0 = v19 + int32(16)
	return v83
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v105 = int32(0)
	if v99 == v105 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if int32(0) < v303 {
		goto L65
	} else {
		goto L66
	}
L24:
	;
	v142 = base.B2i32(l1 <= v141)
	if l1 <= v141 {
		v293 = v105
		goto L23
	} else {
		goto L37
	}
L25:
	;
	v141 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v113 = int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v114 <= v113 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v117 = v113
	goto L30
L29:
	;
	v117 = v114
	goto L30
L30:
	;
	v121 = int32(0)
	v123 = v105
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v99+int32(8)+v121<<(uint(int32(2))%32))))
	if v129 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v141 = v132
	goto L24
L33:
	;
	v132 = v123 + base.I32_popcnt(v129)
	goto L35
L34:
	;
	v132 = v123
	goto L35
L35:
	;
	v134 = v121 + int32(1)
	if v134 != v117 {
		v121 = v134
		v123 = v132
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v145 = F_palloc0(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	if v99 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v203 < int32(0) {
		v293 = v145
		goto L23
	} else {
		goto L50
	}
L40:
	;
	v203 = base.I32_ctz(v189) | v190<<(uint(int32(5))%32)
	goto L39
L41:
	;
	v203 = int32(-2)
	goto L39
L42:
	;
	v156 = base.I32_div_s(int32(0), int32(32))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v157 <= v156 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v160 = v99 + int32(8)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+v156<<(uint(int32(2))%32))))
	v167 = v164 & int32(-1)
	if v167 != 0 {
		v189 = v167
		v190 = v156
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v169 = v156 + int32(1)
	if v169 == v157 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v172 = v169
	goto L46
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v160+v172<<(uint(int32(2))%32))))
	if v179 != 0 {
		v189 = v179
		v190 = v172
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v181 = v172 + int32(1)
	if v181 != v157 {
		v172 = v181
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v208 = int32(1)
	v209 = v203
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145+v209<<(uint(int32(2))%32)))) = v208
	if v99 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v293 = v145
	goto L23
L53:
	;
	if int32(0) <= v284 {
		v208 = v208 + int32(1)
		v209 = v284
		goto L51
	} else {
		goto L64
	}
L54:
	;
	v284 = base.I32_ctz(v270) | v271<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v284 = int32(-2)
	goto L53
L56:
	;
	v235 = v209 + int32(1)
	v237 = base.I32_div_s(v235, int32(32))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v238 <= v237 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v241 = v99 + int32(8)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241+v237<<(uint(int32(2))%32))))
	v248 = v245 & (int32(-1) << (uint(v235) % 32))
	if v248 != 0 {
		v270 = v248
		v271 = v237
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v250 = v237 + int32(1)
	if v250 == v238 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v253 = v250
	goto L60
L60:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v241+v253<<(uint(int32(2))%32))))
	if v260 != 0 {
		v270 = v260
		v271 = v253
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v262 = v253 + int32(1)
	if v262 != v238 {
		v253 = v262
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L52
L65:
	;
	v309 = v303
	v318 = v6
	goto L68
L66:
	;
	goto L67
L67:
	;
	if l1 <= v141 {
		goto L21
	} else {
		goto L99
	}
L68:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(24)+v318<<(uint(int32(2))%32))))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v330 = v328 - int32(1)
	if int32(0) <= v330 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v334 = v327 + int32(4)
	v337 = v330
	goto L73
L71:
	;
	v454 = v309
	goto L72
L72:
	;
	v470 = v318 + int32(1)
	if v470 < v454 {
		v309 = v454
		v318 = v470
		goto L68
	} else {
		goto L98
	}
L73:
	;
	v353 = v334 + v337*int32(120)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v353)+28))
	if v355 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v454 = v452
	goto L72
L75:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v357 = F_RelationGetPartitionKey(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L19
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if l1 <= v141 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v104)+76))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v361 = F_PartitionDirectoryLookup(m, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v353)+28))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	F_InitPartitionPruneContext(m, v353+int32(76), v365, v361, v357, l0, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if int32(0) < v337 {
		v337 = v337 - int32(1)
		goto L73
	} else {
		goto L97
	}
L82:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v353)+20))
	F_bms_free(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L19
	} else {
		goto L83
	}
L83:
	;
	v374 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v353)+20)) = v374
	if v354 <= v374 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v381 = int32(0)
	goto L85
L85:
	;
	v396 = v381 << (uint(int32(2)) % 32)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	v398 = v396 + v397
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if int32(0) <= v399 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L81
L87:
	;
	v430 = v381 + int32(1)
	if v430 != v354 {
		v381 = v430
		goto L85
	} else {
		goto L96
	}
L88:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v353)+20))
	v425 = F_bms_add_member(m, v424, v381)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L19
	} else {
		goto L95
	}
L89:
	;
	v404 = v293 + v399<<(uint(int32(2))%32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	*(*int32)(unsafe.Add(mBase, uint32(v398))) = v405 - int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	if int32(0) < v409 {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v412+v396)))
	if v414 < int32(0) {
		goto L87
	} else {
		goto L93
	}
L92:
	;
	goto L87
L93:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v334+v414*int32(120))+20))
	if v420 == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L88
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+20)) = v425
	goto L87
L96:
	;
	goto L86
L97:
	;
	goto L74
L98:
	;
	goto L69
L99:
	;
	v488 = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v489 == v488 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	if int32(0) <= v546 {
		goto L111
	} else {
		goto L112
	}
L101:
	;
	v546 = base.I32_ctz(v532) | v533<<(uint(int32(5))%32)
	goto L100
L102:
	;
	v546 = int32(-2)
	goto L100
L103:
	;
	v499 = base.I32_div_s(int32(0), int32(32))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v500 <= v499 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v503 = v489 + int32(8)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v503+v499<<(uint(int32(2))%32))))
	v510 = v507 & int32(-1)
	if v510 != 0 {
		v532 = v510
		v533 = v499
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v512 = v499 + int32(1)
	if v512 == v500 {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v515 = v512
	goto L107
L107:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v503+v515<<(uint(int32(2))%32))))
	if v522 != 0 {
		v532 = v522
		v533 = v515
		goto L101
	} else {
		goto L109
	}
L108:
	;
	goto L102
L109:
	;
	v524 = v515 + int32(1)
	if v524 != v500 {
		v515 = v524
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v550 = v488
	v551 = v546
	goto L114
L112:
	;
	v633 = v488
	goto L113
L113:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	F_bms_free(m, v648)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L19
	} else {
		goto L129
	}
L114:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v293+v551<<(uint(int32(2))%32))))
	v571 = F_bms_add_member(m, v550, v568-int32(1))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L19
	} else {
		goto L116
	}
L115:
	;
	v633 = v571
	goto L113
L116:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v573 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if int32(0) <= v629 {
		v550 = v571
		v551 = v629
		goto L114
	} else {
		goto L128
	}
L118:
	;
	v629 = base.I32_ctz(v615) | v616<<(uint(int32(5))%32)
	goto L117
L119:
	;
	v629 = int32(-2)
	goto L117
L120:
	;
	v580 = v551 + int32(1)
	v582 = base.I32_div_s(v580, int32(32))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v583 <= v582 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v586 = v573 + int32(8)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v586+v582<<(uint(int32(2))%32))))
	v593 = v590 & (int32(-1) << (uint(v580) % 32))
	if v593 != 0 {
		v615 = v593
		v616 = v582
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v595 = v582 + int32(1)
	if v595 == v583 {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v598 = v595
	goto L124
L124:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v586+v598<<(uint(int32(2))%32))))
	if v605 != 0 {
		v615 = v605
		v616 = v598
		goto L118
	} else {
		goto L126
	}
L125:
	;
	goto L119
L126:
	;
	v607 = v598 + int32(1)
	if v607 != v583 {
		v598 = v607
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	goto L115
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v633
	F_pfree(m, v293)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L19
	} else {
		goto L130
	}
L130:
	;
	goto L21
L131:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v679 = F_bmsToString(m, v678)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	v681 = F_bmsToString(m, l3)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L19
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v679
	F_errmsg_internal(m, int32(_a_F_ExecInitPartitionExecPruning_0), v19)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ExecInitPartitionExecPruning_1), int32(1898), int32(_a_F_ExecInitPartitionExecPruning_2))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_StorePartitionBound(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v9 = m.G0
	v11 = v9 - int32(256)
	m.G0 = v11
	v15 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v20 = F_SearchSysCacheCopy(m, int32(57), v18, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				v23 = v11 + int32(112)
				v24 = int32(0)
				base.MemoryFill(m, v23, v24, int32(136))
				*(*uint16)(unsafe.Add(mBase, uint32(v11)+96)) = uint16(v24)
				v29 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v29
				*(*uint16)(unsafe.Add(mBase, uint32(v11)+48)) = uint16(v24)
				v47 = F_nodeToString(m, l2)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v49 = F_cstring_to_text(m, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v11)+97)) = uint8(v51)
						*(*int32)(unsafe.Add(mBase, uint32(v11)+244)) = v49
						v54 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v11)+49)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
						v61 = F_heap_modify_tuple(m, v20, v56, v23, v11-int32(-64), v11+int32(16))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
							v66 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v63+v64)+131)) = uint8(v66)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
							if v69 != int32(114) {
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+126)))
								if v72 != int32(1) {
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v75+v76)+126)) = uint8(v78)
								}
							}
							F_CatalogTupleUpdate(m, v15, v61+int32(4), v61)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								F_pfree(m, v61)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									F_relation_close(m, v15, int32(3))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
										if v90 == int32(1) {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											F_update_default_partition_oid(m, v93, v94)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													v100 = F_RelationGetPartitionDesc(m, l1, int32(1))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														v102 = int32(0)
														if v100 == v102 {
															v118 = v102
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
															if v106 == int32(0) {
																v118 = v102
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+32))
																if v109 == int32(-1) {
																	v118 = v102
																} else {
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v109<<(uint(int32(2))%32))))
																	v118 = v116
																}
															}
														}
														if v118 != 0 {
															F_CacheInvalidateRelcacheByRelid(m, v118)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return
															} else {
																F_CacheInvalidateRelcache(m, l1)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return
																} else {
																	m.G0 = v11 + int32(256)
																	return
																}
															}
														} else {
															F_CacheInvalidateRelcache(m, l1)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return
															} else {
																m.G0 = v11 + int32(256)
																return
															}
														}
													}
												}
											}
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												v100 = F_RelationGetPartitionDesc(m, l1, int32(1))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v102 = int32(0)
													if v100 == v102 {
														v118 = v102
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
														if v106 == int32(0) {
															v118 = v102
														} else {
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+32))
															if v109 == int32(-1) {
																v118 = v102
															} else {
																v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v109<<(uint(int32(2))%32))))
																v118 = v116
															}
														}
													}
													if v118 != 0 {
														F_CacheInvalidateRelcacheByRelid(m, v118)
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return
														} else {
															F_CacheInvalidateRelcache(m, l1)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return
															} else {
																m.G0 = v11 + int32(256)
																return
															}
														}
													} else {
														F_CacheInvalidateRelcache(m, l1)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return
														} else {
															m.G0 = v11 + int32(256)
															return
														}
													}
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
				v129 = m.ExcPending
				if v129 != 0 {
					return
				} else {
					v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v130
					F_errmsg_internal(m, int32(_a_F_StorePartitionBound_0), v11)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_StorePartitionBound_1), int32(4069), int32(_a_F_StorePartitionBound_2))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
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
func F_adjust_partition_colnos_using_map(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		v55 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L13
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v55
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		v55 = v3
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = v3
	v22 = v3
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+v21<<(uint(int32(2))%32)))))
	if v28 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v55 = v42
	goto L2
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 < v28 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33+v28<<(uint(int32(1))%32)-int32(2)))))
	if v39 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v42 = F_lappend_int(m, v22, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = v21 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 < v48 {
		v21 = v47
		v22 = v42
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v28
	F_errmsg_internal(m, int32(_a_F_adjust_partition_colnos_using_map_0), v10)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_adjust_partition_colnos_using_map_1), int32(1739), int32(_a_F_adjust_partition_colnos_using_map_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_partition_op_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v241 int32
	_ = v241
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = l1 << (uint(int32(2)) % 32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20+v16)))
	v24 = F_get_opfamily_member(m, v19, v22, v22, base.I32_extend16_s(l2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v109 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v110 - int32(108) {
	case 0:
		goto L27
	default:
		v241 = v109
		goto L25
	case 6:
		goto L26
	}
L2:
	;
	v104 = F_makeRelabelType(m, l3, v30, int32(-1), v100, int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L24
	}
L3:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v92 == int32(7) {
		v106 = l3
		goto L1
	} else {
		goto L23
	}
L4:
	;
	return int32(0)
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28+v16)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31+v16)))
	if base.B2i32(v30 == v33)|base.B2i32(v30 == int32(2249)) != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L20
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v58 == int32(7) {
		v106 = l3
		goto L1
	} else {
		goto L18
	}
L10:
	;
	if v30 <= int32(3830) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	switch v30 - int32(2277) {
	case 0, 6:
		goto L9
	case 1, 2, 3, 4, 5:
		goto L3
	default:
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.B2i32(base.Ui32(v30-int32(_a_F_make_partition_op_expr_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v30-int32(_a_F_make_partition_op_expr_1)) < base.Ui32(int32(2))) != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	if base.B2i32(v30 == int32(2776))|base.B2i32(v30 == int32(3500)) != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	if v30 != int32(3831) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	v62 = l1 << (uint(int32(2)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66+v62)))
	if v65 == v68 {
		v106 = l3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v100 = v65
	goto L2
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74+v16)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77+v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg_internal(m, int32(_a_F_make_partition_op_expr_2), v13)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_make_partition_op_expr_3), int32(3848), int32(_a_F_make_partition_op_expr_4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+l1<<(uint(int32(2))%32))))
	v100 = v99
	goto L2
L24:
	;
	v106 = v104
	goto L1
L25:
	;
	m.G0 = v13 + int32(32)
	return v241
L26:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229+l1<<(uint(int32(2))%32))))
	v234 = F_make_opclause(m, v24, v106, l4, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L50
	}
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v113 < int32(2) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v178 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	v117 = l1 << (uint(int32(2)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)))
	v121 = F_get_element_type(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v121 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v124 = F_palloc0(m, int32(36))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(35)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v117+v128)))
	v131 = F_get_array_type(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v131
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134+v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138+v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+32)) = int32(-1)
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+20)) = uint8(v143)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v140
	v148 = F_palloc0(m, int32(36))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(20)
	v153 = F_get_opcode(m, v24)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+20)) = uint8(v155)
	*(*int64)(unsafe.Add(mBase, uint32(v148)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+8)) = v153
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160+v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+24)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v124
	v172 = F_list_make2_impl(m, v13+int32(20), v13+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v172
	v241 = v148
	goto L25
L37:
	;
	if int32(2) <= v113 {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v213 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v185 = int32(0)
	v188 = v109
	goto L41
L41:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v194 = int32(2)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v188<<(uint(v194)%32))))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198+l1<<(uint(v194)%32))))
	v203 = F_make_opclause(m, v24, v106, v197, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	v213 = v205
	goto L37
L43:
	;
	v205 = F_lappend(m, v185, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v208 = v188 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v208 < v209 {
		v185 = v205
		v188 = v208
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v225 = F_makeBoolExpr(m, int32(1), v213, int32(-1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v241 = v228
	goto L25
L49:
	;
	v241 = v225
	goto L25
L50:
	;
	v241 = v234
	goto L25
}
func F_partition_range_datum_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v17 = v15 - int32(1)
	if int32(0) <= v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = v17
	v31 = int32(-1)
	goto L4
L2:
	;
	v130 = int32(-1)
	goto L3
L3:
	;
	return v130
L4:
	;
	v34 = int32(0)
	v39 = base.I32_div_s(v30+v31+int32(1), int32(2))
	if l3 <= v34 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v130 = v116
	goto L3
L6:
	;
	if v116 < v119 {
		v30 = v119
		v31 = v120
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v116 = v31
	v119 = v39 - int32(1)
	v120 = v31
	goto L6
L8:
	;
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v104)
	return v39
L9:
	;
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v102)
	v116 = v39
	v119 = v30
	v120 = v39
	goto L6
L10:
	;
	v43 = v39 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47+v43)))
	v58 = v34
	goto L11
L11:
	;
	v64 = v58 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46+v64)))
	switch v66 + int32(1) {
	case 0:
		goto L9
	default:
		goto L13
	case 2:
		goto L7
	}
L12:
	;
	if int32(0) < v78 {
		goto L7
	} else {
		goto L20
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1+v64)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64+v49)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4+v64)))
	v78 = F_FunctionCall2Coll(m, l0+v58*int32(28), v73, v75, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v78 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v85 = v58 + int32(1)
	if v85 == l3 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L12
L19:
	;
	v58 = v85
	goto L11
L20:
	;
	goto L9
L21:
	;
	goto L5
}
