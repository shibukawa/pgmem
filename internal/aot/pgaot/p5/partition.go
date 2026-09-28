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
		*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(51539607556)
		v26 = F_hash_create(m, int32(_a_F_CreatePartitionDirectory_1), int64(256), v8, int32(1064))
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
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	v67 = F_MakeSingleTupleTableSlot(m, v59, int32(_a_F_ExecInitPartitionDispatchInfo_1))
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
	if v85 < v87 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v120 = v83 << (uint(int32(2)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120+v121))) = v47
	if l3 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	if v87 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v91
	v95 = F_palloc_mul(m, v91, v91)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v87 << (uint(int32(1)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v109 = F_repalloc(m, v106, v87<<(uint(int32(3))%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v95
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v100 = F_palloc_mul(m, int32(4), v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v100
	goto L23
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v116 = F_repalloc(m, v112, v113<<(uint(int32(2))%32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v116
	goto L23
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitPartitionDispatchInfo[1])) = v27
	return v47
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v120))) = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v131 = F_palloc0(m, int32(216))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(394)
	v135 = int32(0)
	F_InitResultRelInfo(m, v131, v38, v135, l5, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v139+v120))) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l3+l4<<(uint(int32(2))%32))+24)) = v83
	goto L32
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
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v479 int32
	_ = v479
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
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
	v686 = m.ExcPending
	if v686 != 0 {
		goto L19
	} else {
		goto L133
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
	v107 = int64(0)
	if v99 == v105 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if int32(0) < v312 {
		goto L67
	} else {
		goto L68
	}
L24:
	;
	v152 = base.B2i32(l1 <= v151)
	if l1 <= v151 {
		v302 = v105
		goto L23
	} else {
		goto L39
	}
L25:
	;
	v151 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v112 = v99 + int32(8)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v113 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v151 = base.I32_popcnt(v116)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v119 = v113 << (uint(int32(2)) % 32)
	if v119 <= int32(7) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v151 = base.I32_wrap_i64(v146)
	goto L24
L32:
	;
	if v119 == int32(0) {
		v146 = v107
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v143 = F_pg_popcount_optimized(m, v112, v119)
	mBase = m.M
	v146 = v143
	goto L31
L35:
	;
	v124 = v119
	v125 = v112
	v126 = v107
	goto L36
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+3)))
	v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+uint32(_c_F_ExecInitPartitionExecPruning[0]))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+2)))
	v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v129)+uint32(_c_F_ExecInitPartitionExecPruning[0]))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
	v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_ExecInitPartitionExecPruning[0]))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v134 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v133)+uint32(_c_F_ExecInitPartitionExecPruning[0]))))
	v138 = v128 + (v130 + (v132 + (v126 + v134)))
	v139 = int32(4)
	v142 = v124 - v139
	if v142 != 0 {
		v124 = v142
		v125 = v125 + v139
		v126 = v138
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v146 = v138
	goto L31
L38:
	;
	goto L37
L39:
	;
	v154 = F_palloc0_mul(m, int32(4), l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	if v99 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v212 < int32(0) {
		v302 = v154
		goto L23
	} else {
		goto L52
	}
L42:
	;
	v212 = base.I32_ctz(v198) | v199<<(uint(int32(5))%32)
	goto L41
L43:
	;
	v212 = int32(-2)
	goto L41
L44:
	;
	v163 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v166 <= v163 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v169 = v99 + int32(8)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v176 = v173 & int32(-1)
	if v176 != 0 {
		v198 = v176
		v199 = v163
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v177 = int32(1)
	if v177 == v166 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v181 = v177
	goto L48
L48:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v169+v181<<(uint(int32(2))%32))))
	if v188 != 0 {
		v198 = v188
		v199 = v181
		goto L42
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v190 = v181 + int32(1)
	if v190 != v166 {
		v181 = v190
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v217 = int32(1)
	v218 = v212
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154+v218<<(uint(int32(2))%32)))) = v217
	if v99 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v302 = v154
	goto L23
L55:
	;
	if int32(0) <= v293 {
		v217 = v217 + int32(1)
		v218 = v293
		goto L53
	} else {
		goto L66
	}
L56:
	;
	v293 = base.I32_ctz(v279) | v280<<(uint(int32(5))%32)
	goto L55
L57:
	;
	v293 = int32(-2)
	goto L55
L58:
	;
	v244 = v218 + int32(1)
	v246 = int32(base.Ui32(v244) >> (uint(int32(5)) % 32))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v247 <= v246 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v250 = v99 + int32(8)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v246<<(uint(int32(2))%32))))
	v257 = v254 & (int32(-1) << (uint(v244) % 32))
	if v257 != 0 {
		v279 = v257
		v280 = v246
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v259 = v246 + int32(1)
	if v259 == v247 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v262 = v259
	goto L62
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v250+v262<<(uint(int32(2))%32))))
	if v269 != 0 {
		v279 = v269
		v280 = v262
		goto L56
	} else {
		goto L64
	}
L63:
	;
	goto L57
L64:
	;
	v271 = v262 + int32(1)
	if v271 != v247 {
		v262 = v271
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L54
L67:
	;
	v318 = v312
	v327 = v6
	goto L70
L68:
	;
	goto L69
L69:
	;
	if l1 <= v151 {
		goto L21
	} else {
		goto L101
	}
L70:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(24)+v327<<(uint(int32(2))%32))))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v339 = v337 - int32(1)
	if int32(0) <= v339 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v343 = v336 + int32(4)
	v346 = v339
	goto L75
L73:
	;
	v463 = v318
	goto L74
L74:
	;
	v479 = v327 + int32(1)
	if v479 < v463 {
		v318 = v463
		v327 = v479
		goto L70
	} else {
		goto L100
	}
L75:
	;
	v362 = v343 + v346*int32(120)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)+28))
	if v364 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v463 = v461
	goto L74
L77:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v366 = F_RelationGetPartitionKey(m, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L19
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if l1 <= v151 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v104)+76))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v370 = F_PartitionDirectoryLookup(m, v368, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L19
	} else {
		goto L81
	}
L81:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v362)+28))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	F_InitPartitionPruneContext(m, v362+int32(76), v374, v370, v366, l0, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L19
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	if int32(0) < v346 {
		v346 = v346 - int32(1)
		goto L75
	} else {
		goto L99
	}
L84:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v362)+20))
	F_bms_free(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	v383 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v362)+20)) = v383
	if v363 <= v383 {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v390 = int32(0)
	goto L87
L87:
	;
	v405 = v390 << (uint(int32(2)) % 32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	v407 = v405 + v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	if int32(0) <= v408 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L83
L89:
	;
	v439 = v390 + int32(1)
	if v439 != v363 {
		v390 = v439
		goto L87
	} else {
		goto L98
	}
L90:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v362)+20))
	v434 = F_bms_add_member(m, v433, v390)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L19
	} else {
		goto L97
	}
L91:
	;
	v413 = v302 + v408<<(uint(int32(2))%32)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v414 - int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	if int32(0) < v418 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421+v405)))
	if v423 < int32(0) {
		goto L89
	} else {
		goto L95
	}
L94:
	;
	goto L89
L95:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v343+v423*int32(120))+20))
	if v429 == int32(0) {
		goto L89
	} else {
		goto L96
	}
L96:
	;
	goto L90
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+20)) = v434
	goto L89
L98:
	;
	goto L88
L99:
	;
	goto L76
L100:
	;
	goto L71
L101:
	;
	v497 = int32(0)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v498 == v497 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	if int32(0) <= v555 {
		goto L113
	} else {
		goto L114
	}
L103:
	;
	v555 = base.I32_ctz(v541) | v542<<(uint(int32(5))%32)
	goto L102
L104:
	;
	v555 = int32(-2)
	goto L102
L105:
	;
	v506 = int32(0)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v509 <= v506 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v512 = v498 + int32(8)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v519 = v516 & int32(-1)
	if v519 != 0 {
		v541 = v519
		v542 = v506
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v520 = int32(1)
	if v520 == v509 {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v524 = v520
	goto L109
L109:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v512+v524<<(uint(int32(2))%32))))
	if v531 != 0 {
		v541 = v531
		v542 = v524
		goto L103
	} else {
		goto L111
	}
L110:
	;
	goto L104
L111:
	;
	v533 = v524 + int32(1)
	if v533 != v509 {
		v524 = v533
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v559 = v497
	v560 = v555
	goto L116
L114:
	;
	v642 = v497
	goto L115
L115:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	F_bms_free(m, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L19
	} else {
		goto L131
	}
L116:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v302+v560<<(uint(int32(2))%32))))
	v580 = F_bms_add_member(m, v559, v577-int32(1))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L19
	} else {
		goto L118
	}
L117:
	;
	v642 = v580
	goto L115
L118:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v582 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	if int32(0) <= v638 {
		v559 = v580
		v560 = v638
		goto L116
	} else {
		goto L130
	}
L120:
	;
	v638 = base.I32_ctz(v624) | v625<<(uint(int32(5))%32)
	goto L119
L121:
	;
	v638 = int32(-2)
	goto L119
L122:
	;
	v589 = v560 + int32(1)
	v591 = int32(base.Ui32(v589) >> (uint(int32(5)) % 32))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v592 <= v591 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v595 = v582 + int32(8)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v595+v591<<(uint(int32(2))%32))))
	v602 = v599 & (int32(-1) << (uint(v589) % 32))
	if v602 != 0 {
		v624 = v602
		v625 = v591
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v604 = v591 + int32(1)
	if v604 == v592 {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v607 = v604
	goto L126
L126:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v595+v607<<(uint(int32(2))%32))))
	if v614 != 0 {
		v624 = v614
		v625 = v607
		goto L120
	} else {
		goto L128
	}
L127:
	;
	goto L121
L128:
	;
	v616 = v607 + int32(1)
	if v616 != v592 {
		v607 = v616
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	goto L117
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v642
	F_pfree(m, v302)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	goto L21
L133:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v688 = F_bmsToString(m, v687)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	v690 = F_bmsToString(m, l3)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v690
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v688
	F_errmsg_internal(m, int32(_a_F_ExecInitPartitionExecPruning_0), v19)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ExecInitPartitionExecPruning_1), int32(2006), int32(_a_F_ExecInitPartitionExecPruning_2))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L19
	} else {
		goto L137
	}
L137:
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
	var v18 int64
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
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
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v9 = m.G0
	v11 = v9 - int32(384)
	m.G0 = v11
	v15 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v18 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
		v20 = F_SearchSysCacheCopy(m, int32(57), v18, int64(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				v23 = v11 + int32(112)
				v24 = int32(0)
				base.MemoryFill(m, v23, v24, int32(272))
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
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v11)+49)) = uint8(v53)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+376)) = base.I64_extend_i32_u(v49)
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
						v62 = F_heap_modify_tuple(m, v20, v57, v23, v11-int32(-64), v11+int32(16))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
							v67 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+131)) = uint8(v67)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+119)))
							if v70 != int32(114) {
							} else {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+126)))
								if v73 != int32(1) {
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
									v79 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v76+v77)+126)) = uint8(v79)
								}
							}
							F_CatalogTupleUpdate(m, v15, v62+int32(4), v62)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_pfree(m, v62)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									F_relation_close(m, v15, int32(3))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
										if v91 == int32(1) {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											F_update_default_partition_oid(m, v94, v95)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = F_RelationGetPartitionDesc(m, l1, int32(1))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														v103 = int32(0)
														if v101 == v103 {
															v119 = v103
														} else {
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
															if v107 == int32(0) {
																v119 = v103
															} else {
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
																if v110 == int32(-1) {
																	v119 = v103
																} else {
																	v113 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v110<<(uint(int32(2))%32))))
																	v119 = v117
																}
															}
														}
														if v119 != 0 {
															F_CacheInvalidateRelcacheByRelid(m, v119)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return
															} else {
																F_CacheInvalidateRelcache(m, l1)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	m.G0 = v11 + int32(384)
																	return
																}
															}
														} else {
															F_CacheInvalidateRelcache(m, l1)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v11 + int32(384)
																return
															}
														}
													}
												}
											}
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = F_RelationGetPartitionDesc(m, l1, int32(1))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = int32(0)
													if v101 == v103 {
														v119 = v103
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
														if v107 == int32(0) {
															v119 = v103
														} else {
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
															if v110 == int32(-1) {
																v119 = v103
															} else {
																v113 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v110<<(uint(int32(2))%32))))
																v119 = v117
															}
														}
													}
													if v119 != 0 {
														F_CacheInvalidateRelcacheByRelid(m, v119)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CacheInvalidateRelcache(m, l1)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v11 + int32(384)
																return
															}
														}
													} else {
														F_CacheInvalidateRelcache(m, l1)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															m.G0 = v11 + int32(384)
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
				v130 = m.ExcPending
				if v130 != 0 {
					return
				} else {
					v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v131
					F_errmsg_internal(m, int32(_a_F_StorePartitionBound_0), v11)
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_StorePartitionBound_1), int32(_a_F_StorePartitionBound_2), int32(_a_F_StorePartitionBound_3))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
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
	F_errfinish(m, int32(_a_F_adjust_partition_colnos_using_map_1), int32(1847), int32(_a_F_adjust_partition_colnos_using_map_2))
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
	F_errfinish(m, int32(_a_F_make_partition_op_expr_3), int32(3840), int32(_a_F_make_partition_op_expr_4))
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
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
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
	v133 = int32(-1)
	goto L3
L3:
	;
	return v133
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
	v133 = v119
	goto L3
L6:
	;
	if v119 < v122 {
		v30 = v122
		v31 = v123
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v119 = v31
	v122 = v39 - int32(1)
	v123 = v31
	goto L6
L8:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v107)
	return v39
L9:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v105)
	v119 = v39
	v122 = v30
	v123 = v39
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
	if int32(0) < v84 {
		goto L7
	} else {
		goto L20
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1+v64)))
	v75 = v58 << (uint(int32(3)) % 32)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v49+v75)))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(l4+v75)))
	v80 = F_FunctionCall2Coll(m, l0+v58*int32(28), v73, v77, v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v84 = base.I32_wrap_i64(v80)
	if v84 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v88 = v58 + int32(1)
	if v88 == l3 {
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
	v58 = v88
	goto L11
L20:
	;
	goto L9
L21:
	;
	goto L5
}
