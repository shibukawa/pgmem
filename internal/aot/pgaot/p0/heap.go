package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleCleanMoved(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v6 = int32(1)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if base.Ui32(v7) < base.Ui32(int32(_a_F_HeapTupleCleanMoved_0)) {
		v184 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L46
	} else {
		goto L65
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L46
	} else {
		goto L62
	}
L3:
	;
	return v184
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v10) < base.Ui32(int32(3)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v142 != 0 {
		goto L2
	} else {
		goto L45
	}
L6:
	;
	v142 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleCleanMoved[0]))
	if v22 == v10 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v142 = int32(1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleCleanMoved[1]))
	if v26 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v142 = v132
	goto L5
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleCleanMoved[2]))
	if v30 == int32(0) {
		v132 = int32(0)
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleCleanMoved[3]))
	v102 = int32(0)
	v105 = v26 - int32(1)
	goto L35
L16:
	;
	v35 = v30
	goto L17
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v41 == int32(4) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v132 = int32(0)
	goto L12
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v35)+80))
	if v95 != 0 {
		v35 = v95
		goto L17
	} else {
		goto L34
	}
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v44 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v47 = int32(1)
	if v10 == v44 {
		v132 = v47
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v51 = v49 - int32(1)
	if v51 < int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	v57 = int32(0)
	v60 = v51
	goto L24
L24:
	;
	v65 = int32(2)
	v66 = base.I32_div_s(v60-v57, v65)
	v67 = v66 + v57
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54+v67<<(uint(v65)%32))))
	if v71 == v10 {
		v132 = v47
		goto L12
	} else {
		goto L26
	}
L25:
	;
	goto L19
L26:
	;
	v80 = base.B2i32(v71-v10 < int32(0)) | base.B2i32(base.Ui32(v71) < base.Ui32(int32(3)))
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = v67 + int32(1)
	goto L29
L28:
	;
	v81 = v57
	goto L29
L29:
	;
	if v80 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v84 = v60
	goto L32
L31:
	;
	v84 = v67 - int32(1)
	goto L32
L32:
	;
	if v81 <= v84 {
		v57 = v81
		v60 = v84
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	goto L18
L35:
	;
	v110 = int32(2)
	v111 = base.I32_div_s(v105-v102, v110)
	v112 = v111 + v102
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v100+v112<<(uint(v110)%32))))
	v117 = base.B2i32(v116 == v10)
	if v116 == v10 {
		v132 = v117
		goto L12
	} else {
		goto L37
	}
L36:
	;
	v132 = v117
	goto L12
L37:
	;
	v120 = base.B2i32(base.Ui32(v116) < base.Ui32(v10))
	if base.Ui32(v116) < base.Ui32(v10) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v121 = v112 + int32(1)
	goto L40
L39:
	;
	v121 = v102
	goto L40
L40:
	;
	if base.Ui32(v116) < base.Ui32(v10) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v124 = v105
	goto L43
L42:
	;
	v124 = v112 - int32(1)
	goto L43
L43:
	;
	if v121 <= v124 {
		v102 = v121
		v105 = v124
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v143 = F_TransactionIdIsInProgress(m, v10)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	if v143 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v148 = l0 + int32(20)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148))))
	if v149&int32(_a_F_HeapTupleCleanMoved_0) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	F_BufferSetHintBits16(m, v148, v175&int32(_a_F_HeapTupleCleanMoved_1), l1)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L46
	} else {
		goto L61
	}
L50:
	;
	v175 = v173
	v177 = int32(0)
	goto L49
L51:
	;
	v152 = F_TransactionIdDidCommit(m, v10)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L46
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if int32(0) <= base.I32_extend16_s(v149) {
		v184 = v6
		goto L3
	} else {
		goto L56
	}
L54:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148))))
	if v152 != 0 {
		v173 = v154 | int32(512)
		goto L50
	} else {
		goto L55
	}
L55:
	;
	v175 = v154 | int32(256)
	v177 = int32(1)
	goto L49
L56:
	;
	v163 = F_TransactionIdDidCommit(m, v10)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L46
	} else {
		goto L57
	}
L57:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148))))
	if v163 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = v165 | int32(256)
	v177 = int32(1)
	goto L49
L59:
	;
	goto L60
L60:
	;
	v173 = v165 | int32(512)
	goto L50
L61:
	;
	v184 = v177
	goto L3
L62:
	;
	F_errmsg_internal(m, int32(_a_F_HeapTupleCleanMoved_2), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L46
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_HeapTupleCleanMoved_3), int32(243), int32(_a_F_HeapTupleCleanMoved_4))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L46
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_HeapTupleCleanMoved_5), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L46
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_HeapTupleCleanMoved_3), int32(246), int32(_a_F_HeapTupleCleanMoved_4))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L46
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_HeapTupleHeaderGetDatum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v4&int32(4) == int32(0) {
		return base.I64_extend_i32_u(l0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = F_lookup_rowtype_tupdesc(m, v11, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = F_toast_flatten_tuple_to_datum(m, l0, int32(base.Ui32(v17)>>(uint(int32(2))%32)), v13)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if int32(0) <= v22 {
					F_DecrTupleDescRefCount(m, v13)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						return v20
					}
				} else {
					return v20
				}
			}
		}
	}
}
func F_HeapTupleSatisfiesVacuum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_HeapTupleSatisfiesVacuumHorizon(m, l0, l2, v7+int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(2) {
			v17 = int32(0)
			v18 = int32(2)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if base.B2i32(base.Ui32(v18) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == v17 {
				v31 = base.B2i32(base.Ui32(v19) < base.Ui32(l1))
			} else {
				v31 = int32(base.Ui32(v19-l1) >> (uint(int32(31)) % 32))
			}
			if v31 != 0 {
				v32 = v17
			} else {
				v32 = v18
			}
			v34 = v32
		} else {
			v34 = v11
		}
		m.G0 = v7 + int32(16)
		return v34
	}
}
func F_heap_create_with_catalog(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int64, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	v22 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(80)
	m.G0 = v29
	v33 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_CheckAttributeNamesTypes(m, l8, l10, l17)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = F_get_relname_relid(m, l0, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L159
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L155
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L151
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L148
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L143
	}
L9:
	;
	if v39 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = int64(0)
	v48 = F_GetSysCacheOid(m, int32(81), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1), v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L139
	}
L13:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = F_moveArrayTypeName(m, v48, l0, l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l12 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v50 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v57 = base.B2i32(l2 != int32(1664))
	goto L21
L20:
	;
	v57 = int32(0)
	goto L21
L21:
	;
	if v57 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v58 = int32(0)
	if l3 != 0 {
		v107 = l3
		v108 = v58
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_LockRelationOid(m, v107, int32(8))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[0])))
	if v60 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v105 = F_GetNewRelFileNumber(m, l2, v33, l11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L26:
	;
	if l10 == int32(116) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[1]))
	if v66 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[2]))
	if v80 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[1])) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[3]))
	if v73 == v70 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[3])) = int32(0)
	v107 = v66
	v108 = v73
	goto L23
L32:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[2])) = v84
	v87 = l10 - int32(83)
	if base.B2i32(base.Ui32(int32(31)) < base.Ui32(v87))|base.B2i32(int32(1)<<(uint(v87)%32)&int32(-2076180479) == v84) != 0 {
		v107 = v80
		v108 = v58
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[4]))
	if v98 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[4])) = int32(0)
	v107 = v80
	v108 = v98
	goto L23
L35:
	;
	v107 = v105
	v108 = v58
	goto L23
L36:
	;
	if l16 == int32(0) {
		v134 = v22
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v230)+104)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+96)) = int64(-4647714815446351872)
	if l10 == int32(83) {
		goto L56
	} else {
		goto L57
	}
L38:
	;
	v140 = F_heap_create(m, l0, l1, l2, v107, v108, l7, l8, l10, l11, l12, l13, l17, v29+int32(52), v29+int32(48), int32(1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	switch l10 - int32(83) {
	case 0:
		goto L41
	default:
		v134 = v22
		goto L38
	case 19, 26, 29, 31, 35:
		goto L40
	}
L40:
	;
	v132 = F_get_user_default_acl(m, int32(42), l6, l1)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v118 = F_get_user_default_acl(m, int32(38), l6, l1)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v126 = F_heap_create(m, l0, l1, l2, v107, v108, l7, l8, int32(83), l11, l12, l13, l17, v29+int32(52), v29+int32(48), int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+132)) = l19
	v223 = int32(0)
	v226 = v126
	v227 = v118
	goto L37
L44:
	;
	v134 = v132
	goto L38
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+132)) = l19
	v144 = int32(0)
	switch l10 - int32(73) {
	case 0, 10:
		v223 = v144
		v226 = v140
		v227 = v134
		goto L37
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L46
	default:
		goto L47
	}
L46:
	;
	v154 = int32(0)
	v166 = F_AssignTypeArrayOid(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	switch l10 - int32(105) {
	case 0, 11:
		v223 = v144
		v226 = v140
		v227 = v134
		goto L37
	default:
		goto L46
	}
L48:
	;
	v168 = int32(0)
	F_TypeCreate(m, v29+int32(68), l4, l0, l1, v107, l10, l6, int32(-1), int32(99), int32(67), v154, int32(44), int32(2290), int32(2291), int32(2402), int32(2403), v154, v154, v154, v154, v154, v154, v166, v168, v168, v168, v168, int32(100), int32(120), int32(-1), v168, v168, v168)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if l20 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l20)+8)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(l20)+4)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(l20))) = v181
	goto L52
L51:
	;
	goto L52
L52:
	;
	v189 = F_makeArrayTypeName(m, l0, l1)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v191 = int32(0)
	v193 = int32(-1)
	F_TypeCreate(m, v29+int32(68), v166, v189, l1, v191, v191, l6, v193, int32(98), int32(65), v191, int32(44), int32(750), int32(751), int32(2400), int32(2401), v191, v191, int32(3816), int32(_a_F_heap_create_with_catalog_0), v180, int32(1), v191, v191, v191, v191, v191, int32(100), int32(120), v193, v191, v191, v191)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v189)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v223 = v180
	v226 = v140
	v227 = v134
	goto L37
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v230)+96)) = int64(4575657221408423937)
	goto L58
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+140)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v230)+136)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v230)+80)) = l6
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+131)) = uint8(v242)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+76)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v230)+72)) = v223
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v226)+52))
	if v223 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v248 = v223
	goto L61
L60:
	;
	v248 = int32(2249)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v226)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v250)+8)) = int32(-1)
	F_InsertPgClassTuple(m, v33, v226, v107, base.I64_extend_i32_u(v227), l15)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v226)+52))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v260 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v263 = F_CatalogOpenIndexes(m, v260)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_InsertPgAttributeTuples(m, v260, v256, v107, int32(0), v263)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if int32(0) < v257 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v270 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v375 = l10 - int32(99)
	v376 = int32(0)
	if base.B2i32(v375 == v376)|base.B2i32(v375 == int32(19)) == v376 {
		goto L79
	} else {
		goto L80
	}
L69:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v302 = v256 + v296<<(uint(int32(3))%32) + v270*int32(100)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+119)))
	if v303 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L68
L71:
	;
	if v343 != v257 {
		v270 = v343
		goto L69
	} else {
		goto L78
	}
L72:
	;
	v343 = v270 + int32(1)
	goto L71
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(1247)
	v314 = v270 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v314
	v317 = v302 + int32(28)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v318
	v323 = v29 + int32(68)
	v325 = v29 + int32(56)
	F_recordDependencyOn(m, v323, v325, int32(110))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v317)+96))
	if base.B2i32(v329 == int32(0))|base.B2i32(v329 == int32(100)) != 0 {
		v343 = v314
		goto L71
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(3456)
	F_recordDependencyOn(m, v323, v325, int32(110))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v343 = v314
	goto L71
L78:
	;
	goto L70
L79:
	;
	v385 = F_CreateTupleDesc(m, int32(6), int32(_a_F_heap_create_with_catalog_1))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_CatalogCloseIndexes(m, v263)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L85
	}
L82:
	;
	F_InsertPgAttributeTuples(m, v260, v385, v107, int32(0), v263)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_FreeTupleDesc(m, v385)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L81
L85:
	;
	F_relation_close(m, v260, int32(3))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[5]))
	if v399 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[6]))
	if v469 != 0 {
		goto L106
	} else {
		goto L107
	}
L88:
	;
	v403 = l10 - int32(99)
	if base.B2i32(v403 == int32(0))|base.B2i32(v403 == int32(17)) != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v107
	v412 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v412
	F_recordDependencyOnOwner(m, v412, v107, l6)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_recordDependencyOnNewAcl(m, int32(1259), v107, l6, v227)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_recordDependencyOnCurrentExtension(m, v29+int32(68), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v425 = F_new_object_addresses(m)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(2615)
	v433 = v29 + int32(56)
	F_add_exact_object_address(m, v433, v425)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if l5 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(1247)
	F_add_exact_object_address(m, v433, v425)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	switch l10 - int32(109) {
	case 0, 5:
		goto L100
	default:
		goto L101
	}
L98:
	;
	goto L97
L99:
	;
	F_record_object_address_dependencies(m, v29+int32(68), v425, int32(110))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L104
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(2601)
	F_add_exact_object_address(m, v29+int32(56), v425)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	if base.B2i32(l7 == int32(0))|base.B2i32(l10 != int32(112)) != 0 {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	goto L99
L104:
	;
	F_free_object_addresses(m, v425)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L87
L106:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v107, int32(0), l18)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if l9 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l14) {
		goto L128
	} else {
		goto L129
	}
L111:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v478 <= int32(0) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v481 = int32(0)
	v485 = v481
	v500 = v481
	goto L114
L114:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l9)+12))
	v510 = int32(2)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509+v485<<(uint(v510)%32))))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	switch v514 - v510 {
	case 0:
		goto L117
	default:
		goto L118
	case 3:
		goto L119
	}
L115:
	;
	if v552 <= int32(0) {
		goto L110
	} else {
		goto L126
	}
L116:
	;
	v554 = v485 + int32(1)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v554 < v555 {
		v485 = v554
		v500 = v552
		goto L114
	} else {
		goto L125
	}
L117:
	;
	v547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v513)+12)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	v549 = F_StoreAttrDefault(m, v226, v547, v548, l18)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L124
	}
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+20)))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+21)))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+22)))
	v526 = int32(*(*int16)(unsafe.Add(mBase, uint32(v513)+24)))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+26)))
	v528 = F_StoreRelCheck(m, v226, v517, v518, v519, (v520^int32(-1))&int32(1), v525, v526, v527, l18)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v528
	v552 = v500 + int32(1)
	goto L116
L121:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v537
	F_errmsg_internal(m, int32(_a_F_heap_create_with_catalog_2), v29)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(2367), int32(_a_F_heap_create_with_catalog_4))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v549
	v552 = v500
	goto L116
L125:
	;
	goto L115
L126:
	;
	F_SetRelationNumChecks(m, v226, v552)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L110
L128:
	;
	v590 = l14
	goto L130
L129:
	;
	v590 = int32(0)
	goto L130
L130:
	;
	if v590 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v591 = int32(_a_F_heap_create_with_catalog_5)
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[7]))
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[7])) = v595
	v598 = F_palloc(m, int32(16))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	F_relation_close(m, v226, int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L137
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+4)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v598))) = v107
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[9]))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+8))
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v598)+8)) = v604
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[10]))
	v610 = F_lcons(m, v598, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[7])) = v592
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[10])) = v610
	goto L133
L137:
	;
	F_relation_close(m, v33, int32(3))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	m.G0 = v29 + int32(80)
	return v107
L139:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = l0
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_6), v29+int32(32))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1197), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(_a_F_heap_create_with_catalog_8))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l0
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_9), v29+int32(16))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errhint(m, int32(_a_F_heap_create_with_catalog_10), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1216), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errmsg_internal(m, int32(_a_F_heap_create_with_catalog_11), int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1223), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_12), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1254), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_13), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1265), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_14), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1275), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	switch int32(base.Ui32(v12)>>(uint(int32(4))%32))&int32(7) - int32(1) {
	case 0:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v26
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_0), v8+int32(32))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v35, int32(_a_F_heap_desc_1))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v39
				F_appendStringInfo(m, l0, int32(_a_F_heap_desc_2), v8+int32(16))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					m.G0 = v8 + int32(192)
					return
				}
			}
		}
	case 1:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v47
		*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v46
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_3), v8-int32(-64))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v55, int32(_a_F_heap_desc_4))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v60
				*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v59
				F_appendStringInfo(m, l0, int32(_a_F_heap_desc_5), v8+int32(48))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					m.G0 = v8 + int32(192)
					return
				}
			}
		}
	case 2:
		v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)))
		F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_6))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return
		} else {
			if v94&int32(1) != 0 {
				F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_7))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					if v94&int32(2) != 0 {
						F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_8))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v109-int32(1)))))
							if v113 == int32(32) {
								v117 = v109 - int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
								v120 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v108+v117))) = uint8(v120)
							} else {
							}
							F_appendStringInfoChar(m, l0, int32(93))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v126
								F_appendStringInfo(m, l0, int32(_a_F_heap_desc_9), v8+int32(112))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_10))
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
										F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(248), int32(0))
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return
										} else {
											m.G0 = v8 + int32(192)
											return
										}
									}
								}
							}
						}
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v109-int32(1)))))
						if v113 == int32(32) {
							v117 = v109 - int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
							v120 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v108+v117))) = uint8(v120)
						} else {
						}
						F_appendStringInfoChar(m, l0, int32(93))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v126
							F_appendStringInfo(m, l0, int32(_a_F_heap_desc_9), v8+int32(112))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_10))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(248), int32(0))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return
									} else {
										m.G0 = v8 + int32(192)
										return
									}
								}
							}
						}
					}
				}
			} else {
				if v94&int32(2) != 0 {
					F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_8))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v109-int32(1)))))
						if v113 == int32(32) {
							v117 = v109 - int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
							v120 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v108+v117))) = uint8(v120)
						} else {
						}
						F_appendStringInfoChar(m, l0, int32(93))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v126
							F_appendStringInfo(m, l0, int32(_a_F_heap_desc_9), v8+int32(112))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_10))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(248), int32(0))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return
									} else {
										m.G0 = v8 + int32(192)
										return
									}
								}
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v109-int32(1)))))
					if v113 == int32(32) {
						v117 = v109 - int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
						v120 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v108+v117))) = uint8(v120)
					} else {
					}
					F_appendStringInfoChar(m, l0, int32(93))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v126
						F_appendStringInfo(m, l0, int32(_a_F_heap_desc_9), v8+int32(112))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_heap_desc_10))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
								F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(248), int32(0))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return
								} else {
									m.G0 = v8 + int32(192)
									return
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v71
		*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v70
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_3), v8+int32(96))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v79, int32(_a_F_heap_desc_4))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v85
				*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v83
				F_appendStringInfo(m, l0, int32(_a_F_heap_desc_5), v8+int32(80))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					m.G0 = v8 + int32(192)
					return
				}
			}
		}
	case 4:
		v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = v144
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_11), v8+int32(128))
		mBase = m.M
		v150 = m.ExcPending
		if v150 != 0 {
			return
		} else {
			m.G0 = v8 + int32(192)
			return
		}
	case 5:
		v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v152
		*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = v151
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_0), v8+int32(160))
		mBase = m.M
		v159 = m.ExcPending
		if v159 != 0 {
			return
		} else {
			v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v160, int32(_a_F_heap_desc_1))
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return
			} else {
				v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v164
				F_appendStringInfo(m, l0, int32(_a_F_heap_desc_2), v8+int32(144))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return
				} else {
					m.G0 = v8 + int32(192)
					return
				}
			}
		}
	case 6:
		v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v171
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_11), v8+int32(176))
		mBase = m.M
		v177 = m.ExcPending
		if v177 != 0 {
			return
		} else {
			v178 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v181 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v182 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
			F_standby_desc_invalidations(m, l0, v178, v11+int32(20), v181, v182, v183)
			mBase = m.M
			v185 = m.ExcPending
			if v185 != 0 {
				return
			} else {
				m.G0 = v8 + int32(192)
				return
			}
		}
	default:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
		F_appendStringInfo(m, l0, int32(_a_F_heap_desc_12), v8)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			m.G0 = v8 + int32(192)
			return
		}
	}
}
func F_heap_fetch_toast_slice(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	v20 = m.G0
	v22 = v20 - int32(288)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = int32(1)
	v31 = F_toast_open_indexes(m, l0, v25, v22+int32(284), v22+int32(108))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v22+int32(112), int32(1), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = int32(1996)
	v42 = base.I32_div_u_s(l3, v41)
	v43 = int32(1)
	v46 = base.I32_div_u_s(l2-v43, v41)
	v49 = l3 + l4 - v43
	v51 = base.I32_div_u_s(v49, v41)
	if base.B2i32(v46 == v51)&base.B2i32(base.Ui32(l3) <= base.Ui32(int32(1995))) != 0 {
		v81 = v25
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83+v31<<(uint(int32(2))%32))))
	v88 = F_get_toast_snapshot(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L5:
	;
	v56 = base.I64_extend_i32_u(v42)
	v58 = v22 + int32(168)
	if v42 == v51 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v60 = int32(2)
	F_ScanKeyInit(m, v58, v60, int32(3), int32(65), v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_ScanKeyInit(m, v58, int32(2), int32(4), int32(150), v56)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v81 = v60
	goto L4
L10:
	;
	v73 = int32(2)
	F_ScanKeyInit(m, v22+int32(224), v73, v73, int32(149), base.I64_extend_i32_u(v51))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v81 = int32(3)
	goto L4
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L90
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L86
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L82
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L78
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L75
	}
L17:
	;
	if v244 != v51+int32(1) {
		goto L14
	} else {
		goto L72
	}
L18:
	;
	v92 = F_systable_beginscan_ordered(m, l0, v87, v88, v81, v22+int32(112))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v95 = F_systable_getnext_ordered(m, v92, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v95 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v244 = v42
	goto L17
L22:
	;
	goto L23
L23:
	;
	v101 = v22 + int32(107)
	v102 = F_fastgetattr_1(m, v95, int32(2), v24, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v105 = F_fastgetattr_1(m, v95, int32(3), v24, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v128 = base.I32_wrap_i64(v102)
	if v42 != v128 {
		v428 = v42
		v430 = v128
		goto L12
	} else {
		goto L31
	}
L26:
	;
	v107 = base.I32_wrap_i64(v105)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v111 = int32(1)
	if v108&v111 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v120 = int32(4)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v126 = v120
	v127 = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v120
	goto L25
L30:
	;
	v116 = int32(1)
	v126 = v111
	v127 = int32(base.Ui32(v108)>>(uint(v116)%32)) - v116
	goto L25
L31:
	;
	if base.Ui32(v51) < base.Ui32(v42) {
		v313 = v42
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v134 = v46*int32(-1996) + l2
	if base.Ui32(v42) < base.Ui32(v46) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v136 = int32(1996)
	goto L35
L34:
	;
	v136 = v134
	goto L35
L35:
	;
	if v127 != v136 {
		v375 = v127
		v379 = v42
		v381 = v136
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v140 = l5 - l3 + int32(4)
	v143 = v49 - v51*int32(1996)
	if v42 == v51 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v147 = v143
	goto L39
L38:
	;
	v147 = v127 - int32(1)
	goto L39
L39:
	;
	v149 = v42 * int32(1996)
	v150 = l3 - v149
	v153 = v147 - v150 + int32(1)
	if v153 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryCopy(m, v140+v149+v150, v107+v126+v150, v153)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v159 = int32(1)
	v160 = v42 + v159
	v162 = F_systable_getnext_ordered(m, v92, v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v162 == int32(0) {
		v244 = v160
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v172 = v162
	v175 = v160
	goto L45
L45:
	;
	v189 = v22 + int32(107)
	v190 = F_fastgetattr_1(m, v172, int32(2), v24, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v244 = v233
	goto L17
L47:
	;
	v193 = F_fastgetattr_1(m, v172, int32(3), v24, v189)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v218 = base.I32_wrap_i64(v190)
	if v218 != v175 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v195 = base.I32_wrap_i64(v193)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v196&int32(3) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v204 = int32(4)
	v216 = int32(base.Ui32(v201)>>(uint(int32(2))%32)) - v204
	v217 = v204
	goto L48
L51:
	;
	goto L52
L52:
	;
	if v196&int32(1) == int32(0) {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v211 = int32(1)
	v216 = int32(base.Ui32(v196)>>(uint(v211)%32)) - v211
	v217 = v211
	goto L48
L54:
	;
	v428 = v175
	v430 = v218
	goto L12
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v51) < base.Ui32(v175) {
		v313 = v175
		goto L15
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v175) < base.Ui32(v46) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v223 = int32(1996)
	goto L60
L59:
	;
	v223 = v134
	goto L60
L60:
	;
	if v223 != v216 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v375 = v216
	v379 = v175
	v381 = v223
	goto L13
L62:
	;
	goto L63
L63:
	;
	if v175 == v51 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v226 = v143 + int32(1)
	goto L66
L65:
	;
	v226 = v216
	goto L66
L66:
	;
	if v226 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	base.MemoryCopy(m, v140+v175*int32(1996), v195+v217, v226)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v232 = int32(1)
	v233 = v175 + v232
	v235 = F_systable_getnext_ordered(m, v92, v232)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v235 != 0 {
		v172 = v235
		v175 = v233
		goto L45
	} else {
		goto L71
	}
L71:
	;
	goto L46
L72:
	;
	F_systable_endscan_ordered(m, v92)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v22)+108))
	F_toast_close_indexes(m, v261, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	m.G0 = v22 + int32(288)
	return
L75:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v291 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_0), v22+int32(96))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(729), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v332 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v313
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_3), v22+int32(16))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(749), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v357 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_4), v22)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(786), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v397 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v46 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v375
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_5), v22+int32(48))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(758), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v22)+92)) = v446 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_6), v22+int32(80))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(742), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 <= int32(1664) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < v16 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L14
	} else {
		goto L39
	}
L4:
	;
	v73 = F_heap_compute_data_size(m, l0, l1, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v26 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	v69 = v4
	v70 = v4
	v72 = int32(24)
	goto L4
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v26))))
	if v33 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v69 = int32(1)
	v70 = int32(128)
	v72 = (int32(base.Ui32(v16+int32(7))>>(uint(int32(3))%32)) + int32(30)) & int32(536870904)
	goto L4
L11:
	;
	goto L12
L12:
	;
	v47 = v26 + int32(1)
	if v47 != v16 {
		v26 = v47
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	return int32(0)
L15:
	;
	v77 = v73 + v72
	v80 = F_palloc0(m, v77+int32(24))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v77 << (uint(int32(2)) % 32)
	v85 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v85
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+8)) = uint16(v85)
	v89 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v77
	v93 = v80 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+32)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+40)) = uint16(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+36)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v97
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+42)))
	v106 = v103&int32(_a_F_heap_form_tuple_0) | v16
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+42)) = uint16(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+46)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v72 + v93
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v70
	if v69 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v116 = v80 + int32(47)
	goto L19
L18:
	;
	v116 = v85
	goto L19
L19:
	;
	if v69 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v120 = v116 - int32(1)
	goto L22
L21:
	;
	v120 = int32(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v120
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+44)))
	v124 = v122 & int32(_a_F_heap_form_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+44)) = uint16(v124)
	if int32(0) < v111 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v138 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	m.G0 = v14 + int32(32)
	return v80
L26:
	;
	v145 = v138 << (uint(int32(3)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v150 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v151 = v14 + int32(24)
	goto L30
L29:
	;
	v151 = int32(0)
	goto L30
L30:
	;
	if l1 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(l1+v145)))
	v159 = v157
	goto L33
L32:
	;
	v159 = int64(0)
	goto L33
L33:
	;
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v138))))
	v163 = v161
	goto L36
L35:
	;
	v163 = int32(1)
	goto L36
L36:
	;
	F_fill_val(m, v145+(l0+int32(28)), v151, v14+int32(20), v14+int32(28), v80+int32(44), v159, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v167 = v138 + int32(1)
	if v167 != v111 {
		v138 = v167
		goto L26
	} else {
		goto L38
	}
L38:
	;
	goto L27
L39:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(1664)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg(m, int32(_a_F_heap_form_tuple_2), v14)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_heap_form_tuple_3), int32(1042), int32(_a_F_heap_form_tuple_4))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14311(m, l0, l1, l2, l3, int32(_a_F_heap_getattr_1_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_heap_getnext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+188))
	if v6 == int32(_a_F_heap_getnext_0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
		if v11&int32(1) != 0 {
			F_heapgettup_pagemode(m, l0, int32(1), v10, v9)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if v22 == int32(0) {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
					if v30 == int32(0) {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+268)))
						if v33 != int32(1) {
							return l0 + int32(68)
						} else {
							F_pgstat_assoc_relation(m, v29)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+272))
								v40 = v39
								v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v41 + int64(1)
								return l0 + int32(68)
							}
						}
					} else {
						v40 = v30
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v41 + int64(1)
						return l0 + int32(68)
					}
				}
			}
		} else {
			F_heapgettup(m, l0, int32(1), v10, v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if v22 == int32(0) {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
					if v30 == int32(0) {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+268)))
						if v33 != int32(1) {
							return l0 + int32(68)
						} else {
							F_pgstat_assoc_relation(m, v29)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+272))
								v40 = v39
								v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v41 + int64(1)
								return l0 + int32(68)
							}
						}
					} else {
						v40 = v30
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v41 + int64(1)
						return l0 + int32(68)
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_heap_getnext_1), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_heap_getnext_2), int32(1450), int32(_a_F_heap_getnext_3))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
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
func F_heap_hot_search_buffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
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
	var v240 int32
	_ = v240
	v7 = l6
	v8 = int32(0)
	if l2 < v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l5 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_heap_hot_search_buffer[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+(l2^int32(-1))<<(uint(int32(2))%32))))
	v37 = v29
	goto L1
L3:
	;
	goto L4
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_heap_hot_search_buffer[1]))
	v37 = v31 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v39 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v56 = v7
	v58 = v39
	v61 = v7 ^ int32(1)
	v62 = v8
	v63 = v8
	goto L10
L10:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)))
	if base.B2i32(base.Ui32(v69) < base.Ui32(int32(25)))|base.B2i32(base.Ui32(int32(base.Ui32(v69+int32(_a_F_heap_hot_search_buffer_0))>>(uint(int32(2))%32))&int32(_a_F_heap_hot_search_buffer_1)) < base.Ui32(v58)) != 0 {
		goto L8
	} else {
		goto L12
	}
L11:
	;
	return v240
L12:
	;
	v82 = v37 + int32(20) + v58<<(uint(int32(2))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v87 = int32(base.Ui32(v83)>>(uint(int32(15))%32)) & int32(3)
	if v87 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L11
L14:
	;
	v238 = int32(0)
	if v233 != 0 {
		v56 = v238
		v58 = v233
		v61 = v235
		v62 = v236
		v63 = v237
		goto L10
	} else {
		goto L62
	}
L15:
	;
	if (v56^int32(-1)|base.B2i32(v87 != int32(2)))&int32(1) != 0 {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v101 = v37 + v83&int32(_a_F_heap_hot_search_buffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(base.Ui32(v103) >> (uint(int32(17)) % 32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+8)) = uint16(v58)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)) = uint16(v42)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)) = uint16(v43)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v107
	if v56&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v233 = v83 & int32(_a_F_heap_hot_search_buffer_2)
	v235 = v61
	v236 = v62
	v237 = v63
	goto L14
L19:
	;
	v114 = int32(0)
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+18)))
	if v115 < v114 {
		v240 = v114
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v62 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	if v61&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+20)))
	v122 = int32(768)
	if v121&v122 != v122 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v128 = v126
	goto L27
L26:
	;
	v128 = int32(2)
	goto L27
L27:
	;
	if v128 == v62 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	if l5 == int32(0) {
		v206 = v63
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v132 = F_HeapTupleSatisfiesVisibility(m, l4, l3, l2)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	F_HeapCheckForSerializableConflictOut(m, v132, l1, l4, l2, l3)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v132 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v58)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+20)))
	v143 = int32(768)
	if v142&v143 != v143 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v149 = v147
	goto L37
L36:
	;
	v149 = int32(2)
	goto L37
L37:
	;
	F_PredicateLockTID(m, l1, l4+int32(4), l3, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if l5 == int32(0) {
		v240 = int32(1)
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v155)
	return int32(1)
L40:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)))
	if v208&int32(64) == int32(0) {
		goto L8
	} else {
		goto L56
	}
L41:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v162 != int32(1) {
		v206 = v63
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if v63 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v167 = F_GlobalVisHorizonKindForRel(m, l1)
	mBase = m.M
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v167<<(uint(int32(2))%32))+uint32(_c_F_heap_hot_search_buffer[2])))
	goto L46
L44:
	;
	v171 = v63
	goto L45
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+20)))
	if v174&int32(256) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v171 = v170
	goto L45
L47:
	;
	if v200 != 0 {
		v206 = v171
		goto L40
	} else {
		goto L55
	}
L48:
	;
	v200 = int32(base.Ui32(v174&int32(512)) >> (uint(int32(9)) % 32))
	goto L47
L49:
	;
	goto L50
L50:
	;
	if v174&int32(2048)|base.B2i32(v174&int32(_a_F_heap_hot_search_buffer_3) != int32(1024))|base.B2i32(v174&int32(_a_F_heap_hot_search_buffer_4) == int32(64)) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v199 = int32(0)
	goto L53
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v197 = F_GlobalVisTestIsRemovableXid(m, v171, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L31
	} else {
		goto L54
	}
L53:
	;
	v200 = v199
	goto L47
L54:
	;
	v199 = v197
	goto L53
L55:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v201)
	v206 = v171
	goto L40
L56:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+20)))
	if v213&int32(2048)|base.B2i32(v213&int32(768) == int32(512)) != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+16)))
	if v213&int32(_a_F_heap_hot_search_buffer_5) == int32(_a_F_heap_hot_search_buffer_6) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v227 = F_HeapTupleGetUpdateXid(m, v207)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L31
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v233 = v221
	v235 = int32(0)
	v236 = v229
	v237 = v206
	goto L14
L61:
	;
	v233 = v221
	v235 = int32(0)
	v236 = v227
	v237 = v206
	goto L14
L62:
	;
	v240 = v238
	goto L13
}
func F_heap_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(0)
	v24 = F_heap_prepare_insert(m, l0, l1, v20, l2, l3)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v27 = int32(0)
	v32 = F_RelationGetBufferForTuple(m, l0, v26, v27, l3, l4, v18+int32(12), v27, v27)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v32 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+(v32^int32(-1))<<(uint(int32(2))%32))))
	v51 = v43
	goto L3
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[1]))
	v51 = v45 + v32<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L9:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+10)))
	v58 = v56 & int32(4)
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_LockBufferInternal(m, v59, int32(3))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v63 = int32(_a_F_heap_insert_0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_insert[2])) = v65 + int32(1)
	v70 = l3 & int32(16)
	F_RelationPutHeapTuple(m, v32, v24, int32(base.Ui32(v70)>>(uint(int32(4))%32)))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	if v58 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v82 = F_visibilitymap_clear(m, v75|v76<<(uint(int32(16))%32), v80, int32(3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v88 = int32(0)
	goto L17
L17:
	;
	if l3&int32(4)|base.B2i32(base.Ui32(v20) < base.Ui32(int32(3))) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+10)))
	v86 = v84 & int32(_a_F_heap_insert_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+10)) = uint16(v86)
	v88 = v82
	goto L17
L19:
	;
	F_MarkBufferDirty(m, v32)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v95 = int32(0)
	if base.B2i32(base.Ui32(v94) < base.Ui32(int32(3)))|base.B2i32(v95 <= v20-v94) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v102 = v94
	goto L23
L22:
	;
	v102 = v95
	goto L23
L23:
	;
	if v102 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v20
	goto L19
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+118)))
	if v108 != int32(112) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v312 = int32(_a_F_heap_insert_0)
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_insert[2])) = v314 - int32(1)
	F_UnlockReleaseBuffer(m, v32)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L92
	}
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[3]))
	if v112 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v148 = int32(0)
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)))
	if v150 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L38
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v115 != 0 {
		goto L26
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v112 != int32(1) {
		goto L29
	} else {
		goto L36
	}
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v116 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[4])))
	if v118 == int32(1) {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[4])))
	if v124&int32(1) == int32(0) {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	if base.B2i32(base.Ui32(v129) < base.Ui32(int32(_a_F_heap_insert_2))) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v134 == int32(0) {
		goto L28
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_log_heap_new_cid(m, l0, v24)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+119)))
	switch v138 - int32(109) {
	case 0, 5:
		goto L43
	default:
		goto L28
	}
L43:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+112)))
	if v141 != int32(1) {
		goto L28
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	goto L28
L46:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+12)))
	v164 = base.B2i32(base.Ui32(int32(24)) < base.Ui32(v155)) & base.B2i32((v155+int32(_a_F_heap_insert_3))&int32(_a_F_heap_insert_4) == int32(4))
	if v164 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v169 = v148
	v170 = v148
	goto L48
L48:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+8)) = uint16(v150)
	v173 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	if v70 != 0 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v165 = int32(6)
	goto L51
L50:
	;
	v165 = int32(0)
	goto L51
L51:
	;
	if v164 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v168 = int32(-128)
	goto L54
L53:
	;
	v168 = int32(0)
	goto L54
L54:
	;
	v169 = v165
	v170 = v168
	goto L48
L55:
	;
	v176 = v173 | int32(4)
	goto L57
L56:
	;
	v176 = v173
	goto L57
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)) = uint8(v176)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[3]))
	if v179 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L77
	}
L59:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[4])))
	if v183&int32(1) == int32(0) {
		v222 = v169
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+118)))
	if v189 != int32(112) {
		v222 = v169
		goto L58
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	if v179 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v194 != 0 {
		v222 = v169
		goto L58
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+119)))
	if v196 == int32(102) {
		v222 = v169
		goto L58
	} else {
		goto L69
	}
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v195 != 0 {
		v222 = v169
		goto L58
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L70
L70:
	;
	if base.B2i32(base.Ui32(v199) < base.Ui32(int32(_a_F_heap_insert_2)))|l3&int32(8) != 0 {
		v222 = v169
		goto L58
	} else {
		goto L71
	}
L71:
	;
	v206 = v176 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)) = uint8(v206)
	v209 = v169 | int32(16)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+68))
	if v211 != int32(99) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v216 == int32(0) {
		v222 = v209
		goto L58
	} else {
		goto L76
	}
L73:
	;
	v214 = F_isTempToastNamespace(m, v211)
	mBase = m.M
	v216 = v214
	goto L75
L74:
	;
	v216 = int32(1)
	goto L75
L75:
	;
	goto L72
L76:
	;
	v220 = v176 | int32(24)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)) = uint8(v220)
	v222 = v209
	goto L58
L77:
	;
	F_XLogRegisterData(m, v18+int32(8), int32(3))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+2)) = uint16(v232)
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v234)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+6)) = uint8(v236)
	F_XLogRegisterBuffer(m, int32(0), v32, v222|int32(8))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_XLogRegisterBufData(m, int32(0), v18+int32(2), int32(5))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v251 = int32(23)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	F_XLogRegisterBufData(m, int32(0), v250+v251, v253-v251)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v259 = int32(_a_F_heap_insert_5)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[5])))
	v262 = v261 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[5])) = uint8(v262)
	goto L82
L82:
	;
	if v88 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_XLogRegisterBuffer(m, int32(1), v265, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v300 = F_XLogInsert(m, int32(10), v170&int32(255))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v272 = F_XLogInsert(m, int32(10), v170&int32(255))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v275 = base.I64_rotl(v272, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v277 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[0]))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v281+(v277^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v287))) = v275
	goto L26
L89:
	;
	goto L90
L90:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v290+v277<<(uint(int32(13))%32))+uint32(_c_F_heap_insert[6]))) = v275
	goto L26
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = base.I64_rotl(v300, int64(32))
	goto L26
L92:
	;
	if v58 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_UnlockBuffer(m, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v323 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	F_ReleaseBuffer(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_CacheInvalidateHeapTuple(m, l0, v24, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	F_pgstat_count_heap_insert(m, l0, int64(1))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if l1 != v24 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v333)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v335
	F_pfree(m, v24)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	m.G0 = v18 + int32(16)
	return
L106:
	;
	goto L105
}
func F_heap_log_freeze_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v6) < base.Ui32(v7) {
		return int32(-1)
	} else {
		v11 = int32(1)
		if base.Ui32(v7) < base.Ui32(v6) {
			v40 = v11
			return v40
		} else {
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			if base.Ui32(v13) < base.Ui32(v14) {
				return int32(-1)
			} else {
				if base.Ui32(v14) < base.Ui32(v13) {
					v40 = v11
					return v40
				} else {
					v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
					v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
					if base.Ui32(v19) < base.Ui32(v20) {
						return int32(-1)
					} else {
						if base.Ui32(v20) < base.Ui32(v19) {
							v40 = v11
							return v40
						} else {
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
							v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
							if base.Ui32(v25) < base.Ui32(v26) {
								return int32(-1)
							} else {
								if base.Ui32(v26) < base.Ui32(v25) {
									v40 = v11
								} else {
									v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
									v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
									if base.Ui32(v32) < base.Ui32(v33) {
										v40 = int32(-1)
									} else {
										v40 = base.B2i32(base.Ui32(v33) < base.Ui32(v32))
									}
								}
								return v40
							}
						}
					}
				}
			}
		}
	}
}
func F_heap_reloptions(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	switch l0 - int32(109) {
	case 0, 5:
		v20 = int32(1)
		v25 = F_build_reloptions(m, l1, v20, v20, int32(136), int32(_a_F_heap_reloptions_0), int32(26))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	default:
		return
	case 7:
		v10 = F_build_reloptions(m, l1, int32(1), int32(2), int32(136), int32(_a_F_heap_reloptions_0), int32(26))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 == int32(0) {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = int64(-4616189618054758400)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(100)
				return
			}
		}
	}
}
func F_heap_tableam_handler(m *base.Module, l0 int32) int64 {
	return int64(830052)
}
func F_heap_toast_insert_or_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v665 int32
	_ = v665
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v982 int64
	_ = v982
	var v984 int64
	_ = v984
	var v986 int64
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1087 int32
	_ = v1087
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1114 int64
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(_a_F_heap_toast_insert_or_update_0)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	F_heap_deform_tuple(m, l1, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = v18 + int32(_a_F_heap_toast_insert_or_update_3)
	v33 = v18 + int32(_a_F_heap_toast_insert_or_update_4)
	F_heap_deform_tuple(m, l2, v20, v31, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v38 = v5
	v39 = int32(0)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v18 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v18 + int32(_a_F_heap_toast_insert_or_update_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v18 + int32(_a_F_heap_toast_insert_or_update_1)
	v53 = v18 + int32(4)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)) = uint8(v57)
	if v57 < v56 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v38 = v33
	v39 = v31
	goto L5
L7:
	;
	v70 = v5
	goto L10
L8:
	;
	goto L9
L9:
	;
	v292 = base.I32_div_s(v21+int32(7), int32(8))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	if v298&int32(4) != 0 {
		goto L54
	} else {
		goto L55
	}
L10:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v78 = v70 * int32(12)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v78+v79)+8)) = uint8(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v83+v78))) = v81
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v94 = v55 + v76<<(uint(int32(3))%32) + v70*int32(100)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+113)))
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v78)+9)) = uint8(v95)
	v98 = v94 + int32(28)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v99 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L9
L12:
	;
	v271 = v70 + int32(1)
	if v271 != v56 {
		v70 = v271
		goto L10
	} else {
		goto L53
	}
L13:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v70))))
	if v170 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v101 = v70 << (uint(int32(3)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101+v102)))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+72)))
	if v105 != int32(_a_F_heap_toast_insert_or_update_5) {
		v165 = v104
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+v70<<(uint(int32(3))%32))))
	v165 = v164
	goto L13
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v70))))
	if v110 != 0 {
		v165 = v104
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99+v101)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v113 != int32(1) {
		v165 = v104
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	if v116 != int32(18) {
		v165 = v104
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v70))))
	if v121 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v155 = v154 + v78
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)))
	v158 = v156 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)) = uint8(v158)
	goto L12
L22:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v145 = v144 + v78
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+8)))
	v147 = int32(1)
	v148 = v146 | v147
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+8)) = uint8(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)))
	v152 = v150 | v147
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)) = uint8(v152)
	v165 = v104
	goto L13
L23:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v122 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v125 != int32(18) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+16)))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v112)))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v112)+8))
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v104)+8))
	if base.I64_extend_i32_u(v128^v129)&int64(65535)|(v134^v135|(v137^v138)) == int64(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v174 = v173 + v78
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
	v177 = v175 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)) = uint8(v177)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)))
	v181 = v179 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)) = uint8(v181)
	goto L12
L28:
	;
	goto L29
L29:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+72)))
	if v183 == int32(_a_F_heap_toast_insert_or_update_5) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+84)))
	if v186 == int32(112) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v261 = v260 + v78
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+8)))
	v264 = v262 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+8)) = uint8(v264)
	goto L12
L33:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v190 = v189 + v78
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)))
	v193 = v191 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)) = uint8(v193)
	goto L35
L34:
	;
	goto L35
L35:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v196 != int32(1) {
		v243 = v165
		v244 = v196
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v257+v78)+4)) = v256
	goto L12
L37:
	;
	v246 = int32(1)
	if v244&v246 != 0 {
		v256 = int32(base.Ui32(v244) >> (uint(v246) % 32))
		goto L36
	} else {
		goto L52
	}
L38:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v199+v78))) = v165
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+84)))
	if v202 == int32(112) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v210+v70<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v209)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v217 = v216 + v78
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)))
	v220 = v218 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)) = uint8(v220)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)))
	v224 = v222 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)) = uint8(v224)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v226 != int32(1) {
		v243 = v209
		v244 = v226
		goto L37
	} else {
		goto L45
	}
L40:
	;
	v205 = F_detoast_attr(m, v165)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v207 = F_detoast_external_attr(m, v165)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v209 = v205
	goto L39
L44:
	;
	v209 = v207
	goto L39
L45:
	;
	v230 = int32(18)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v232 == v230 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v235 = v230
	goto L48
L47:
	;
	v235 = int32(2)
	goto L48
L48:
	;
	if base.Ui32((v232-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v242 = int32(6)
	goto L51
L50:
	;
	v242 = v235
	goto L51
L51:
	;
	v256 = v242
	goto L36
L52:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v256 = int32(base.Ui32(v250) >> (uint(int32(2)) % 32))
	goto L36
L53:
	;
	goto L11
L54:
	;
	v301 = (v292 + int32(30)) & int32(-8)
	goto L56
L55:
	;
	v301 = int32(24)
	goto L56
L56:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v302 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+8))
	v305 = v303
	goto L59
L58:
	;
	v305 = int32(2032)
	goto L59
L59:
	;
	v307 = l3 & int32(-17)
	v312 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v495 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L106
	}
L61:
	;
	v314 = v305 - v301
	if base.Ui32(v312) <= base.Ui32(v314) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	goto L63
L63:
	;
	v332 = v18 + int32(4)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+52))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v346 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L60
L65:
	;
	if v429 < int32(0) {
		goto L60
	} else {
		goto L93
	}
L66:
	;
	v429 = int32(-1)
	goto L65
L67:
	;
	goto L68
L68:
	;
	goto L69
L69:
	;
	goto L71
L71:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v332)+24))
	v363 = int32(0)
	v365 = int32(24)
	v366 = int32(-1)
	goto L72
L72:
	;
	v374 = v356 + v363*int32(12)
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+8)))
	if int32(48)&v375 != 0 {
		v411 = v365
		v412 = v366
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v429 = v412
	goto L65
L74:
	;
	v415 = v363 + int32(1)
	if v415 != v346 {
		v363 = v415
		v365 = v411
		v366 = v412
		goto L72
	} else {
		goto L92
	}
L75:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v378 = int32(3)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v377+v363<<(uint(v378)%32))))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if v382&v378 == int32(2) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v390 = int32(1)
	goto L78
L77:
	;
	v390 = int32(0)
	goto L78
L78:
	;
	if base.B2i32(v382 == int32(1))|v390 != 0 {
		v411 = v365
		v412 = v366
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345+v346<<(uint(int32(3))%32)+v363*int32(100))+112)))
	goto L82
L80:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v407 = base.B2i32(v365 < v406)
	if v365 < v406 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	goto L83
L83:
	;
	v399 = v395 - int32(101)
	if base.B2i32(v399 == int32(0))|base.B2i32(v399 == int32(19)) != 0 {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v411 = v365
	v412 = v366
	goto L74
L86:
	;
	v408 = v363
	goto L88
L87:
	;
	v408 = v366
	goto L88
L88:
	;
	if v365 < v406 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v409 = v406
	goto L91
L90:
	;
	v409 = v365
	goto L91
L91:
	;
	v411 = v409
	v412 = v408
	goto L74
L92:
	;
	goto L73
L93:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v432<<(uint(int32(3))%32)+v429*int32(100))+112)))
	if v439 == int32(120) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(32)+v429*int32(12))+4))
	if base.Ui32(v459) <= base.Ui32(v314) {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	F_toast_tuple_try_compression(m, v332, v429)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v444 = int32(32)
	v448 = v18 + v444 + v429*int32(12)
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)))
	v451 = v449 | v444
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)) = uint8(v451)
	goto L94
L98:
	;
	goto L94
L99:
	;
	v473 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+112))
	if v462 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_toast_tuple_externalize(m, v18+int32(4), v429, v307)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	if base.Ui32(v314) < base.Ui32(v473) {
		goto L63
	} else {
		goto L104
	}
L104:
	;
	goto L64
L105:
	;
	v646 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L144
	}
L106:
	;
	if base.Ui32(v495) <= base.Ui32(v314) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	goto L108
L108:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+112))
	if v514 == int32(0) {
		goto L105
	} else {
		goto L110
	}
L109:
	;
	goto L105
L110:
	;
	v518 = v18 + int32(4)
	v519 = int32(0)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+52))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	if v532 <= v519 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v615 < int32(0) {
		goto L105
	} else {
		goto L139
	}
L112:
	;
	v615 = int32(-1)
	goto L111
L113:
	;
	goto L114
L114:
	;
	goto L116
L116:
	;
	goto L117
L117:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v518)+24))
	v549 = int32(0)
	v551 = int32(24)
	v552 = int32(-1)
	goto L118
L118:
	;
	v560 = v542 + v549*int32(12)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+8)))
	if int32(16)&v561 != 0 {
		v597 = v551
		v598 = v552
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v615 = v598
	goto L111
L120:
	;
	v601 = v549 + int32(1)
	if v601 != v532 {
		v549 = v601
		v551 = v597
		v552 = v598
		goto L118
	} else {
		goto L138
	}
L121:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v564 = int32(3)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563+v549<<(uint(v564)%32))))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	if v568&v564 == int32(2) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v576 = v519
	goto L124
L123:
	;
	v576 = int32(0)
	goto L124
L124:
	;
	if base.B2i32(v568 == int32(1))|v576 != 0 {
		v597 = v551
		v598 = v552
		goto L120
	} else {
		goto L125
	}
L125:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531+v532<<(uint(int32(3))%32)+v549*int32(100))+112)))
	goto L128
L126:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v593 = base.B2i32(v551 < v592)
	if v551 < v592 {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	goto L129
L129:
	;
	v585 = v581 - int32(101)
	if base.B2i32(v585 == int32(0))|base.B2i32(v585 == int32(19)) != 0 {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v597 = v551
	v598 = v552
	goto L120
L132:
	;
	v594 = v549
	goto L134
L133:
	;
	v594 = v552
	goto L134
L134:
	;
	if v551 < v592 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v595 = v592
	goto L137
L136:
	;
	v595 = v551
	goto L137
L137:
	;
	v597 = v595
	v598 = v594
	goto L120
L138:
	;
	goto L119
L139:
	;
	F_toast_tuple_externalize(m, v518, v615, v307)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v624 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if base.Ui32(v314) < base.Ui32(v624) {
		goto L108
	} else {
		goto L142
	}
L142:
	;
	goto L109
L143:
	;
	v793 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L181
	}
L144:
	;
	if base.Ui32(v646) <= base.Ui32(v314) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	goto L146
L146:
	;
	v665 = v18 + int32(4)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v677)+52))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if v679 <= int32(0) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L143
L148:
	;
	if v762 < int32(0) {
		goto L143
	} else {
		goto L176
	}
L149:
	;
	v762 = int32(-1)
	goto L148
L150:
	;
	goto L151
L151:
	;
	goto L152
L152:
	;
	goto L154
L154:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v665)+24))
	v696 = int32(0)
	v698 = int32(24)
	v699 = int32(-1)
	goto L155
L155:
	;
	v707 = v689 + v696*int32(12)
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+8)))
	if int32(48)&v708 != 0 {
		v744 = v698
		v745 = v699
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v762 = v745
	goto L148
L157:
	;
	v748 = v696 + int32(1)
	if v748 != v679 {
		v696 = v748
		v698 = v744
		v699 = v745
		goto L155
	} else {
		goto L175
	}
L158:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	v711 = int32(3)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v710+v696<<(uint(v711)%32))))
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714))))
	if v715&v711 == int32(2) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v723 = int32(1)
	goto L161
L160:
	;
	v723 = int32(0)
	goto L161
L161:
	;
	if base.B2i32(v715 == int32(1))|v723 != 0 {
		v744 = v698
		v745 = v699
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678+v679<<(uint(int32(3))%32)+v696*int32(100))+112)))
	goto L164
L163:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v740 = base.B2i32(v698 < v739)
	if v698 < v739 {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	if v728 != int32(109) {
		v744 = v698
		v745 = v699
		goto L157
	} else {
		goto L167
	}
L167:
	;
	goto L163
L169:
	;
	v741 = v696
	goto L171
L170:
	;
	v741 = v699
	goto L171
L171:
	;
	if v698 < v739 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v742 = v739
	goto L174
L173:
	;
	v742 = v698
	goto L174
L174:
	;
	v744 = v742
	v745 = v741
	goto L157
L175:
	;
	goto L156
L176:
	;
	F_toast_tuple_try_compression(m, v665, v762)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v771 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	if base.Ui32(v314) < base.Ui32(v771) {
		goto L146
	} else {
		goto L179
	}
L179:
	;
	goto L147
L180:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	if v942&int32(8) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L181:
	;
	v796 = int32(_a_F_heap_toast_insert_or_update_6) - v301
	if base.Ui32(v793) <= base.Ui32(v796) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	goto L183
L183:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)+112))
	if v814 == int32(0) {
		goto L180
	} else {
		goto L185
	}
L184:
	;
	goto L180
L185:
	;
	v818 = v18 + int32(4)
	v819 = int32(0)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v830)+52))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	if v832 <= v819 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	if v915 < int32(0) {
		goto L180
	} else {
		goto L214
	}
L187:
	;
	v915 = int32(-1)
	goto L186
L188:
	;
	goto L189
L189:
	;
	goto L191
L191:
	;
	goto L192
L192:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v818)+24))
	v849 = int32(0)
	v851 = int32(24)
	v852 = int32(-1)
	goto L193
L193:
	;
	v860 = v842 + v849*int32(12)
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860)+8)))
	if int32(16)&v861 != 0 {
		v897 = v851
		v898 = v852
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v915 = v898
	goto L186
L195:
	;
	v901 = v849 + int32(1)
	if v901 != v832 {
		v849 = v901
		v851 = v897
		v852 = v898
		goto L193
	} else {
		goto L213
	}
L196:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v818)+4))
	v864 = int32(3)
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v863+v849<<(uint(v864)%32))))
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	if v868&v864 == int32(2) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v876 = v819
	goto L199
L198:
	;
	v876 = int32(0)
	goto L199
L199:
	;
	if base.B2i32(v868 == int32(1))|v876 != 0 {
		v897 = v851
		v898 = v852
		goto L195
	} else {
		goto L200
	}
L200:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831+v832<<(uint(int32(3))%32)+v849*int32(100))+112)))
	goto L202
L201:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	v893 = base.B2i32(v851 < v892)
	if v851 < v892 {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	if v881 != int32(109) {
		v897 = v851
		v898 = v852
		goto L195
	} else {
		goto L205
	}
L205:
	;
	goto L201
L207:
	;
	v894 = v849
	goto L209
L208:
	;
	v894 = v852
	goto L209
L209:
	;
	if v851 < v892 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v895 = v892
	goto L212
L211:
	;
	v895 = v851
	goto L212
L212:
	;
	v897 = v895
	v898 = v894
	goto L195
L213:
	;
	goto L194
L214:
	;
	F_toast_tuple_externalize(m, v818, v915, v307)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v924 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	if base.Ui32(v796) < base.Ui32(v924) {
		goto L183
	} else {
		goto L217
	}
L217:
	;
	goto L184
L218:
	;
	v1015 = v18 + int32(4)
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+52))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+20)))
	v1022 = int32(0)
	if base.B2i32(v1019&int32(2) == v1022)|base.B2i32(v1018 <= v1022) != 0 {
		goto L228
	} else {
		goto L229
	}
L219:
	;
	v1009 = l1
	goto L218
L220:
	;
	goto L221
L221:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v952 = base.I32_div_s(v21+int32(7), int32(8))
	if v942&int32(4) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v960 = (v952 + int32(30)) & int32(-8)
	goto L224
L223:
	;
	v960 = int32(24)
	goto L224
L224:
	;
	v962 = v18 + int32(_a_F_heap_toast_insert_or_update_1)
	v964 = v18 + int32(_a_F_heap_toast_insert_or_update_2)
	v965 = F_heap_compute_data_size(m, v20, v962, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v967 = v965 + v960
	v970 = F_palloc0(m, v967+int32(24))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v970))) = v967
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v970)+4)) = v973
	v975 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v970)+8)) = uint16(v975)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v979 = v970 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v970)+16)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v970)+12)) = v977
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v947)+15))
	*(*int64)(unsafe.Add(mBase, uint32(v970)+39)) = v982
	v984 = *(*int64)(unsafe.Add(mBase, uint32(v947)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v970)+32)) = v984
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v947)))
	*(*int64)(unsafe.Add(mBase, uint32(v970)+24)) = v986
	*(*uint8)(unsafe.Add(mBase, uint32(v970)+46)) = uint8(v960)
	v989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v970)+42)))
	v992 = v989&int32(_a_F_heap_toast_insert_or_update_7) | v21
	*(*uint16)(unsafe.Add(mBase, uint32(v970)+42)) = uint16(v992)
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	F_heap_fill_tuple(m, v20, v962, v964, v960+v979, v970+int32(44), (v970+int32(47))&(v999<<(uint(int32(29))%32)>>(uint(int32(31))%32)))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v1009 = v970
	goto L218
L228:
	;
	v1076 = v1019
	goto L230
L229:
	;
	v1028 = int32(0)
	goto L231
L230:
	;
	v1079 = int32(0)
	if base.B2i32(v1076&int32(1) == v1079)|base.B2i32(v1018 <= v1079) == v1079 {
		goto L238
	} else {
		goto L239
	}
L231:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+24))
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043+v1028*int32(12))+8)))
	if v1047&int32(2) != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+20)))
	v1076 = v1060
	goto L230
L233:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+4))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1050+v1028<<(uint(int32(3))%32))))
	F_pfree(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1058 = v1028 + int32(1)
	if v1058 != v1018 {
		v1028 = v1058
		goto L231
	} else {
		goto L237
	}
L236:
	;
	goto L235
L237:
	;
	goto L232
L238:
	;
	v1087 = int32(0)
	goto L241
L239:
	;
	goto L240
L240:
	;
	m.G0 = v18 + int32(_a_F_heap_toast_insert_or_update_0)
	return v1009
L241:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+24))
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102+v1087*int32(12))+8)))
	if v1106&int32(1) != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	goto L240
L243:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+12))
	v1114 = *(*int64)(unsafe.Add(mBase, uint32(v1110+v1087<<(uint(int32(3))%32))))
	F_toast_delete_datum(m, v1114, int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1119 = v1087 + int32(1)
	if v1119 != v1018 {
		v1087 = v1119
		goto L241
	} else {
		goto L247
	}
L246:
	;
	goto L245
L247:
	;
	goto L242
}
func F_make_new_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
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
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	v19 = F_table_open(m, l0, l4)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
		v26 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v26 != 0 {
				v32 = F_SysCacheGetAttr(m, int32(57), v26, int32(33), v17+int32(47))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
					if l3 == int32(116) {
						v38 = F_LookupCreationNamespace(m, int32(_a_F_make_new_heap_0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v42 = v38
							if v34 != 0 {
								v44 = int64(0)
							} else {
								v44 = v32
							}
							*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
							v52 = F_pg_snprintf(m, v17+int32(48), int32(64), int32(_a_F_make_new_heap_1), v17+int32(32))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+80))
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+119)))
								switch v56 - int32(83) {
								case 0, 22, 26, 31, 33:
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+88))
									v62 = base.B2i32(v59 == int32(0))
								default:
									v62 = int32(0)
								}
								v65 = int32(0)
								v73 = int32(1)
								v76 = F_heap_create_with_catalog(m, v17+int32(48), v42, l1, v65, v65, v65, v55, l2, v23, v65, int32(114), l3, v65, v62, v65, v44, v65, v73, v73, l0, v65)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v26)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+112))
											if v83 != 0 {
												v86 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v83))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if v86 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v83
															F_errmsg_internal(m, int32(_a_F_make_new_heap_2), v17+int32(16))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_make_new_heap_3), int32(1350), int32(_a_F_make_new_heap_4))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v95 = F_SysCacheGetAttr(m, int32(57), v86, int32(33), v17+int32(47))
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
															if v97 != 0 {
																v98 = int64(0)
															} else {
																v98 = v95
															}
															v99 = F_table_open(m, v76, l4)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int32(0)
															} else {
																v101 = int32(0)
																v104 = F_create_toast_table(m, v99, v101, v101, v98, l4, v101, v83)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v99, int32(0))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v86)
																		mBase = m.M
																		v110 = m.ExcPending
																		if v110 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v19, int32(0))
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v17 + int32(112)
																				return v76
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												F_relation_close(m, v19, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													m.G0 = v17 + int32(112)
													return v76
												}
											}
										}
									}
								}
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
						v42 = v41
						if v34 != 0 {
							v44 = int64(0)
						} else {
							v44 = v32
						}
						*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
						v52 = F_pg_snprintf(m, v17+int32(48), int32(64), int32(_a_F_make_new_heap_1), v17+int32(32))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+80))
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+119)))
							switch v56 - int32(83) {
							case 0, 22, 26, 31, 33:
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+88))
								v62 = base.B2i32(v59 == int32(0))
							default:
								v62 = int32(0)
							}
							v65 = int32(0)
							v73 = int32(1)
							v76 = F_heap_create_with_catalog(m, v17+int32(48), v42, l1, v65, v65, v65, v55, l2, v23, v65, int32(114), l3, v65, v62, v65, v44, v65, v73, v73, l0, v65)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v26)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_CommandCounterIncrement(m)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+112))
										if v83 != 0 {
											v86 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v83))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v86 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v83
														F_errmsg_internal(m, int32(_a_F_make_new_heap_2), v17+int32(16))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_make_new_heap_3), int32(1350), int32(_a_F_make_new_heap_4))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v95 = F_SysCacheGetAttr(m, int32(57), v86, int32(33), v17+int32(47))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
														if v97 != 0 {
															v98 = int64(0)
														} else {
															v98 = v95
														}
														v99 = F_table_open(m, v76, l4)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int32(0)
														} else {
															v101 = int32(0)
															v104 = F_create_toast_table(m, v99, v101, v101, v98, l4, v101, v83)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v99, int32(0))
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v86)
																	mBase = m.M
																	v110 = m.ExcPending
																	if v110 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v19, int32(0))
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v17 + int32(112)
																			return v76
																		}
																	}
																}
															}
														}
													}
												}
											}
										} else {
											F_relation_close(m, v19, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												m.G0 = v17 + int32(112)
												return v76
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
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
					F_errmsg_internal(m, int32(_a_F_make_new_heap_2), v17)
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_make_new_heap_3), int32(1277), int32(_a_F_make_new_heap_4))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
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
