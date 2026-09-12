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
	v10 = int32(4536272)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = l0
	v15 = F_palloc(m, int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(51539607556)
		v26 = F_hash_create(m, int32(13330), int32(256), v8, int32(1064))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v26
			*(*int32)(unsafe.Add(mBase, _consts[10])) = v11
			m.G0 = v8 + int32(48)
			return v15
		}
	}
}
func F_DestroyPartitionDirectory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_hash_seq_init(m, v5+int32(12), v9)
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
	v14 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v5 + int32(32)
	return
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
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
	v23 = F_hash_seq_search(m, v5+int32(12))
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
		v16 = v23
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[320]))
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
	v26 = int32(4536272)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v29
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v83 = F__emscripten_memset_bulkmem(m, v47+int32(24), base.I32_extend8_s(int32(255)), v79<<(uint(int32(2))%32))
	mBase = m.M
	goto L20
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
	v67 = F_MakeSingleTupleTableSlot(m, v59, int32(1632148))
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
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v86 = v84 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v88 <= v86 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120+v84<<(uint(int32(2))%32)))) = v47
	if l3 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v118
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(4)
	v95 = F_palloc(m, int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v88 << (uint(int32(1)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v109 = F_repalloc(m, v106, v88<<(uint(int32(3))%32))
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
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v101 = F_palloc(m, v98<<(uint(int32(2))%32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v118 = v101
	goto L24
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
	v118 = v116
	goto L24
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v27
	return v47
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v84<<(uint(int32(2))%32)))) = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v134 = F_palloc0(m, int32(216))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(388)
	v138 = int32(0)
	F_InitResultRelInfo(m, v134, v38, v138, l5, v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v143 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v142+v84<<(uint(v143)%32)))) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l3+l4<<(uint(v143)%32))+24)) = v84
	goto L32
}
func F_ExecInitPartitionExecPruning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
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
	var v76 int32
	_ = v76
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
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
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
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
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v483 int32
	_ = v483
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v644 int32
	_ = v644
	var v652 int32
	_ = v652
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = l2 << (uint(int32(2)) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23+v26)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v37 = base.B2i32(l3|v29 == v6)
	if l3 == v6 {
		v76 = v37
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v76 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	goto L1
L3:
	;
	if v29 == int32(0) {
		v76 = v37
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v43 != v44 {
		v76 = int32(0)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v46 = int32(1)
	if v43 <= v46 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v49 = v46
	goto L8
L7:
	;
	v49 = v43
	goto L8
L8:
	;
	v50 = int32(8)
	v55 = int32(0)
	goto L9
L9:
	;
	v63 = v55 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3+v50+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+(v29+v50))))
	v68 = base.B2i32(v65 == v67)
	if v67 != v65 {
		v76 = v68
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v76 = v68
	goto L2
L11:
	;
	v71 = v55 + int32(1)
	if v71 != v49 {
		v55 = v71
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81+v23)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+16)))
	if v84 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L20
	} else {
		goto L132
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v101
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+17)))
	if v103 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+l2<<(uint(int32(2))%32))))
	v101 = v92
	goto L16
L18:
	;
	goto L19
L19:
	;
	v93 = int32(0)
	v97 = F_bms_add_range(m, v93, v93, l1-int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v101 = v97
	goto L16
L22:
	;
	m.G0 = v20 + int32(16)
	return v83
L23:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v107 = int32(0)
	if v101 == v107 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if int32(0) < v307 {
		goto L66
	} else {
		goto L67
	}
L25:
	;
	v144 = base.B2i32(l1 <= v143)
	if l1 <= v143 {
		v297 = v107
		goto L24
	} else {
		goto L38
	}
L26:
	;
	v143 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v115 = int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v116 <= v115 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v119 = v115
	goto L31
L30:
	;
	v119 = v116
	goto L31
L31:
	;
	v123 = int32(0)
	v125 = v107
	goto L32
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(8)+v123<<(uint(int32(2))%32))))
	if v131 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v143 = v134
	goto L25
L34:
	;
	v134 = v125 + base.I32_popcnt(v131)
	goto L36
L35:
	;
	v134 = v125
	goto L36
L36:
	;
	v136 = v123 + int32(1)
	if v136 != v119 {
		v123 = v136
		v125 = v134
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v147 = F_palloc0(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v101 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v205 < int32(0) {
		v297 = v147
		goto L24
	} else {
		goto L51
	}
L41:
	;
	v205 = base.I32_ctz(v191) | v192<<(uint(int32(5))%32)
	goto L40
L42:
	;
	v205 = int32(-2)
	goto L40
L43:
	;
	v158 = base.I32_div_s(int32(0), int32(32))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v159 <= v158 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v162 = v101 + int32(8)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v158<<(uint(int32(2))%32))))
	v169 = v166 & int32(-1)
	if v169 != 0 {
		v191 = v169
		v192 = v158
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v171 = v158 + int32(1)
	if v171 == v159 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v174 = v171
	goto L47
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v162+v174<<(uint(int32(2))%32))))
	if v181 != 0 {
		v191 = v181
		v192 = v174
		goto L41
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v183 = v174 + int32(1)
	if v183 != v159 {
		v174 = v183
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v211 = v205
	v214 = int32(1)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147+v211<<(uint(int32(2))%32)))) = v214
	if v101 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v297 = v147
	goto L24
L54:
	;
	if int32(0) <= v287 {
		v211 = v287
		v214 = v214 + int32(1)
		goto L52
	} else {
		goto L65
	}
L55:
	;
	v287 = base.I32_ctz(v273) | v274<<(uint(int32(5))%32)
	goto L54
L56:
	;
	v287 = int32(-2)
	goto L54
L57:
	;
	v238 = v211 + int32(1)
	v240 = base.I32_div_s(v238, int32(32))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v241 <= v240 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v244 = v101 + int32(8)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v240<<(uint(int32(2))%32))))
	v251 = v248 & (int32(-1) << (uint(v238) % 32))
	if v251 != 0 {
		v273 = v251
		v274 = v240
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v253 = v240 + int32(1)
	if v253 == v241 {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v256 = v253
	goto L61
L61:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v244+v256<<(uint(int32(2))%32))))
	if v263 != 0 {
		v273 = v263
		v274 = v256
		goto L55
	} else {
		goto L63
	}
L62:
	;
	goto L56
L63:
	;
	v265 = v256 + int32(1)
	if v265 != v241 {
		v256 = v265
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L53
L66:
	;
	v315 = v307
	v322 = v6
	goto L69
L67:
	;
	goto L68
L68:
	;
	if l1 <= v143 {
		goto L22
	} else {
		goto L100
	}
L69:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(24)+v322<<(uint(int32(2))%32))))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v335 = v333 - int32(1)
	if int32(0) <= v335 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L68
L71:
	;
	v344 = v335
	goto L74
L72:
	;
	v468 = v315
	goto L73
L73:
	;
	v483 = v322 + int32(1)
	if v483 < v468 {
		v315 = v468
		v322 = v483
		goto L69
	} else {
		goto L99
	}
L74:
	;
	v361 = v332 + int32(4) + v344*int32(120)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	if v363 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v468 = v464
	goto L73
L76:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v365 = F_RelationGetPartitionKey(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L20
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if l1 <= v143 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v106)+76))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v369 = F_PartitionDirectoryLookup(m, v367, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	F_InitPartitionPruneContext(m, v361+int32(76), v373, v369, v365, l0, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L20
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	if int32(0) < v344 {
		v344 = v344 - int32(1)
		goto L74
	} else {
		goto L98
	}
L83:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	F_bms_free(m, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L20
	} else {
		goto L84
	}
L84:
	;
	v382 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+20)) = v382
	if v362 <= v382 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v389 = int32(0)
	goto L86
L86:
	;
	v405 = v389 << (uint(int32(2)) % 32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v407 = v405 + v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	if int32(0) <= v408 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L82
L88:
	;
	v441 = v389 + int32(1)
	if v441 != v362 {
		v389 = v441
		goto L86
	} else {
		goto L97
	}
L89:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v435 = F_bms_add_member(m, v434, v389)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L20
	} else {
		goto L96
	}
L90:
	;
	v413 = v297 + v408<<(uint(int32(2))%32)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v414 - int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	if int32(0) < v418 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421+v405)))
	if v423 < int32(0) {
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L88
L94:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v332+int32(24)+v423*int32(120))))
	if v429 == int32(0) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	goto L89
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+20)) = v435
	goto L88
L97:
	;
	goto L87
L98:
	;
	goto L75
L99:
	;
	goto L70
L100:
	;
	v502 = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v503 == v502 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	if int32(0) <= v560 {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	v560 = base.I32_ctz(v546) | v547<<(uint(int32(5))%32)
	goto L101
L103:
	;
	v560 = int32(-2)
	goto L101
L104:
	;
	v513 = base.I32_div_s(int32(0), int32(32))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v503)+4))
	if v514 <= v513 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v517 = v503 + int32(8)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v513<<(uint(int32(2))%32))))
	v524 = v521 & int32(-1)
	if v524 != 0 {
		v546 = v524
		v547 = v513
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v526 = v513 + int32(1)
	if v526 == v514 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v529 = v526
	goto L108
L108:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v517+v529<<(uint(int32(2))%32))))
	if v536 != 0 {
		v546 = v536
		v547 = v529
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L103
L110:
	;
	v538 = v529 + int32(1)
	if v538 != v514 {
		v529 = v538
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v565 = v560
	v568 = v502
	goto L115
L113:
	;
	v652 = v502
	goto L114
L114:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	F_bms_free(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L20
	} else {
		goto L130
	}
L115:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v297+v565<<(uint(int32(2))%32))))
	v586 = F_bms_add_member(m, v568, v583-int32(1))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L20
	} else {
		goto L117
	}
L116:
	;
	v652 = v586
	goto L114
L117:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v588 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	if int32(0) <= v644 {
		v565 = v644
		v568 = v586
		goto L115
	} else {
		goto L129
	}
L119:
	;
	v644 = base.I32_ctz(v630) | v631<<(uint(int32(5))%32)
	goto L118
L120:
	;
	v644 = int32(-2)
	goto L118
L121:
	;
	v595 = v565 + int32(1)
	v597 = base.I32_div_s(v595, int32(32))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v598 <= v597 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v601 = v588 + int32(8)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v601+v597<<(uint(int32(2))%32))))
	v608 = v605 & (int32(-1) << (uint(v595) % 32))
	if v608 != 0 {
		v630 = v608
		v631 = v597
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v610 = v597 + int32(1)
	if v610 == v598 {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v613 = v610
	goto L125
L125:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v601+v613<<(uint(int32(2))%32))))
	if v620 != 0 {
		v630 = v620
		v631 = v613
		goto L119
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v622 = v613 + int32(1)
	if v622 != v598 {
		v613 = v622
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	goto L116
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v652
	F_pfree(m, v297)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L20
	} else {
		goto L131
	}
L131:
	;
	goto L22
L132:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v696 = F_bmsToString(m, v695)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L20
	} else {
		goto L133
	}
L133:
	;
	v698 = F_bmsToString(m, l3)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L20
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v696
	F_errmsg_internal(m, int32(180349), v20)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L20
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(507856), int32(1898), int32(342016))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L20
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_StorePartitionBound(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(256)
	m.G0 = v10
	v14 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v19 = F_SearchSysCacheCopy(m, int32(57), v17, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v26 = F__emscripten_memset_bulkmem(m, v10+int32(112), base.I32_extend8_s(int32(0)), int32(136))
				mBase = m.M
				v27 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+96)) = uint16(v27)
				v29 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v29
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+48)) = uint16(v27)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v29
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
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+97)) = uint8(v51)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+244)) = v49
						v54 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+49)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
						v63 = F_heap_modify_tuple(m, v19, v56, v10+int32(112), v10-int32(-64), v10+int32(16))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+22)))
							v68 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v65+v66)+131)) = uint8(v68)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
							if v71 != int32(114) {
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+126)))
								if v74 != int32(1) {
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
									v80 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v77+v78)+126)) = uint8(v80)
								}
							}
							F_CatalogTupleUpdate(m, v14, v63+int32(4), v63)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								F_pfree(m, v63)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_sequence_close(m, v14, int32(3))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
										if v92 == int32(1) {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											F_update_default_partition_oid(m, v95, v96)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													v102 = F_RelationGetPartitionDesc(m, l1, int32(1))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														v104 = int32(0)
														if v102 == v104 {
															v120 = v104
														} else {
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
															if v108 == int32(0) {
																v120 = v104
															} else {
																v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+32))
																if v111 == int32(-1) {
																	v120 = v104
																} else {
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v111<<(uint(int32(2))%32))))
																	v120 = v118
																}
															}
														}
														if v120 != 0 {
															F_CacheInvalidateRelcacheByRelid(m, v120)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return
															} else {
																F_CacheInvalidateRelcache(m, l1)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(256)
																	return
																}
															}
														} else {
															F_CacheInvalidateRelcache(m, l1)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return
															} else {
																m.G0 = v10 + int32(256)
																return
															}
														}
													}
												}
											}
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												v102 = F_RelationGetPartitionDesc(m, l1, int32(1))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
													v104 = int32(0)
													if v102 == v104 {
														v120 = v104
													} else {
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
														if v108 == int32(0) {
															v120 = v104
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+32))
															if v111 == int32(-1) {
																v120 = v104
															} else {
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v111<<(uint(int32(2))%32))))
																v120 = v118
															}
														}
													}
													if v120 != 0 {
														F_CacheInvalidateRelcacheByRelid(m, v120)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return
														} else {
															F_CacheInvalidateRelcache(m, l1)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return
															} else {
																m.G0 = v10 + int32(256)
																return
															}
														}
													} else {
														F_CacheInvalidateRelcache(m, l1)
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															m.G0 = v10 + int32(256)
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
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v132
					F_errmsg_internal(m, int32(47234), v10)
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						F_errfinish(m, int32(507612), int32(4069), int32(434470))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
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
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v62
L2:
	;
	v17 = v3
	v19 = v3
	goto L7
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v12 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v62 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22+v19<<(uint(int32(2))%32)))))
	if v26 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v62 = v54
	goto L1
L9:
	;
	v54 = F_lappend_int(m, v17, v37)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L18
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v29 < v26 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(1))%32)-int32(2)))))
	if v37 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
	F_errmsg_internal(m, int32(76582), v10)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(507856), int32(1739), int32(243403))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	v57 = v19 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v57 < v58 {
		v17 = v54
		v19 = v57
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
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
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
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
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v107 - int32(108) {
	case 0:
		goto L31
	default:
		v239 = v106
		goto L28
	case 6:
		goto L30
	}
L2:
	;
	v101 = F_makeRelabelType(m, l3, v33, int32(-1), v98, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L27
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v89 == int32(7) {
		v104 = l3
		goto L1
	} else {
		goto L26
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28+v16)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31+v16)))
	if v30 == v33 {
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
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L23
	}
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v55 == int32(7) {
		v104 = l3
		goto L1
	} else {
		goto L21
	}
L10:
	;
	if v33 == int32(2249) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v33 <= int32(3830) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	switch v33 - int32(2277) {
	case 0, 6:
		goto L9
	case 1, 2, 3, 4, 5:
		goto L3
	default:
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v33-int32(5077)) < base.Ui32(int32(4)) {
		goto L9
	} else {
		goto L18
	}
L15:
	;
	if v33 == int32(2776) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v33 == int32(3500) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L3
L18:
	;
	if base.Ui32(v33-int32(4537)) < base.Ui32(int32(2)) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	if v33 != int32(3831) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v59 = l1 << (uint(int32(2)) % 32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59+v60)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63+v59)))
	if v62 == v65 {
		v104 = l3
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v98 = v62
	goto L2
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71+v16)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74+v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg_internal(m, int32(40594), v13)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(505654), int32(3848), int32(213488))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+l1<<(uint(int32(2))%32))))
	v98 = v96
	goto L2
L27:
	;
	v104 = v101
	goto L1
L28:
	;
	m.G0 = v13 + int32(32)
	return v239
L29:
	;
	if int32(1) < v110 {
		goto L50
	} else {
		goto L51
	}
L30:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+l1<<(uint(int32(2))%32))))
	v214 = F_make_opclause(m, v24, v104, l4, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L49
	}
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v110 < int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v176 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+l1<<(uint(int32(2))%32))))
	v118 = F_get_element_type(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v118 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v121 = F_palloc0(m, int32(36))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = int32(35)
	v126 = l1 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126+v127)))
	v130 = F_get_array_type(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v130
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133+v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137+v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+32)) = int32(-1)
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+20)) = uint8(v142)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v139
	v147 = F_palloc0(m, int32(36))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = int32(20)
	v152 = F_get_opcode(m, v24)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v154 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+20)) = uint8(v154)
	*(*int64)(unsafe.Add(mBase, uint32(v147)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v152
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159+v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v121
	v171 = F_list_make2_impl(m, v13+int32(20), v13+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+28)) = v171
	v239 = v147
	goto L28
L41:
	;
	v218 = int32(0)
	goto L29
L42:
	;
	goto L43
L43:
	;
	v183 = int32(0)
	v186 = v106
	goto L44
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v192 = int32(2)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v186<<(uint(v192)%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196+l1<<(uint(v192)%32))))
	v201 = F_make_opclause(m, v24, v104, v195, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	v218 = v203
	goto L29
L46:
	;
	v203 = F_lappend(m, v183, v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v206 = v186 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v206 < v207 {
		v183 = v203
		v186 = v206
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v239 = v214
	goto L28
L50:
	;
	v230 = F_makeBoolExpr(m, int32(1), v218, int32(-1))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v239 = v233
	goto L28
L53:
	;
	v239 = v230
	goto L28
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
