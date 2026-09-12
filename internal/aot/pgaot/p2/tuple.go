package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescTruncatedCopy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v10 = F_palloc(m, l1*int32(116)+int32(20))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(-4294965047)
	v19 = int32(20)
	v20 = v10 + v19
	v21 = int32(4)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = l1 * int32(100)
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if int32(0) < v34 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v32 = F__emscripten_memcpy_bulkmem(m, v20+l1<<(uint(v21)%32), l0+v24<<(uint(v21)%32)+v19, v31)
	mBase = m.M
	goto L6
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v39 = int32(0)
	v41 = v34
	goto L10
L8:
	;
	goto L9
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v66
	return v10
L10:
	;
	v48 = v20 + v41<<(uint(int32(4))%32) + v39*int32(100)
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+90)) = uint8(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+86)) = v49
	F_populate_compact_attribute(m, v10, v39)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v56 = v39 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v56 < v57 {
		v39 = v56
		v41 = v57
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_CreateTupleQueueDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(24))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(11)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(780)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(781)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(782)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(783)
		return v4
	}
}
func F_ExecInitScanTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	v6 = F_MakeTupleTableSlot(m, l2, l3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v9 = F_lappend(m, v8, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v9
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+96)) = uint8(base.B2i32(l2 != int32(0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v6
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+100)) = uint8(v17)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = l3
			return
		}
	}
}
func F_GetTupleTransactionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v15 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l0, int32(-2), v8+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
		v20 = int32(0)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[643])))
		if v22 == v20 {
			v25 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v25)
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
			v31 = v20
			m.G0 = v8 + int32(16)
			return v31
		} else {
			v29 = F_TransactionIdGetCommitTsData(m, v15, l3, l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = v29
				m.G0 = v8 + int32(16)
				return v31
			}
		}
	}
}
func F_TupleDescGetAttInMetadata(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	v12 = m.G0
	v13 = int32(16)
	v14 = v12 - v13
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = F_palloc(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 != int32(2249) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	v33 = F_palloc0(m, v16*int32(28))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) <= v25 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	F_assign_record_type_typmod(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v36 = v16 << (uint(int32(2)) % 32)
	v37 = F_palloc0(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = F_palloc0(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v16 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v33
	m.G0 = v14 + int32(16)
	return v18
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = l0 + int32(20) + v57<<(uint(int32(4))%32) + v47*int32(100)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+91)))
	if v64 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v71 = v47 << (uint(int32(2)) % 32)
	F_getTypeInputInfo(m, v67, v14+int32(12), v37+v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v86 = v47 + int32(1)
	if v86 != v16 {
		v47 = v86
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_fmgr_info(m, v75, v33+v47*int32(28))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v39+v71))) = v82
	goto L17
L20:
	;
	goto L14
}
func F_TupleDescGetDefault(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(0)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v16 = v11
	goto L8
L5:
	;
	v41 = v11
	goto L6
L6:
	;
	return v41
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v32 = F_stringToNode(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v23 = v13 + v16<<(uint(int32(3))%32)
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if v24 == l1&int32(65535) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v27 = v16 + int32(1)
	if v27 != v12 {
		v16 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	v41 = v32
	goto L6
}
func F_TupleDescInitEntryCollation(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v4<<(uint(int32(4))%32)+l1*int32(100))+16)) = l2
	return
}
func F_tuple_data_split(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v24 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = v2
	goto L3
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v35 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v32 = v28
	goto L3
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v39 = F_pg_detoast_datum_packed(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v43 = v2
	goto L8
L8:
	;
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(6) <= v44 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v41 = F_text_to_cstring(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v43 = v41
	goto L8
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v50 = base.B2i32(v47 != int32(0))
	goto L13
L12:
	;
	v50 = v2
	goto L13
L13:
	;
	v51 = F_superuser(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L21
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L168
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L164
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L160
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L156
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L151
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L147
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L143
	}
L21:
	;
	if v51 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v32 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L139
	}
L25:
	;
	m.G0 = v21 - int32(-64)
	return v431
L26:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
	v431 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v59 = v34 & int32(1)
	if v59 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v139 = F_relation_open(m, v23, int32(1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L48
	}
L30:
	;
	if v43 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v43 != 0 {
		goto L17
	} else {
		goto L47
	}
L33:
	;
	v67 = (v33&int32(2047) + int32(7)) & int32(4088)
	v68 = F_strlen(m, v43)
	mBase = m.M
	if v67 != v68 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v72 = F_palloc(m, v67|int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v67 == int32(0) {
		v129 = v72
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v77 = int32(0)
	v84 = v2
	goto L37
L37:
	;
	v95 = v77 + v43
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v96&int32(254) != int32(48) {
		goto L18
	} else {
		goto L39
	}
L38:
	;
	v129 = v72
	goto L29
L39:
	;
	v104 = v77 & int32(7)
	if v104 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v108 = base.I32_extend8_s(v84)
	goto L42
L41:
	;
	v108 = int32(0)
	goto L42
L42:
	;
	v109 = (v96-int32(48))<<(uint(v104)%32) | v108
	if v104 == int32(7) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(base.Ui32(v77)>>(uint(int32(3))%32))))) = uint8(v109)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v117 = v77 + int32(1)
	if v117 != v67 {
		v77 = v117
		v84 = v109
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v129 = v2
	goto L29
L48:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+52))
	v144 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v146 = F_initArrayResult(m, int32(17), v144, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+119)))
	if v150 != int32(83) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+84))
	if v153 != int32(2) {
		goto L16
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v157 = v33 & int32(2047)
	if v157 <= v148 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	if v403 != v413 {
		goto L14
	} else {
		goto L134
	}
L55:
	;
	v187 = v32 + int32(4)
	v189 = v162 & int32(65535)
	v194 = int32(0)
	v196 = v194
	v198 = v194
	v203 = v146
	goto L64
L56:
	;
	v162 = int32(base.Ui32(v137)>>(uint(int32(2))%32)) - int32(4)
	if v148 != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	v403 = int32(0)
	v408 = v146
	v413 = v162 & int32(65535)
	goto L54
L60:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(218525), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(518937), int32(343), int32(326112))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	if base.Ui32(v157) <= base.Ui32(v196) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v403 = v392
	v408 = v396
	v413 = v189
	goto L54
L66:
	;
	v399 = v196 + int32(1)
	if v399 != v148 {
		v196 = v399
		v198 = v392
		v203 = v396
		goto L64
	} else {
		goto L133
	}
L67:
	;
	v380 = v378 + v308
	v384 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v385 = F_accumArrayResult(m, v203, v332, int32(0), int32(17), v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L130
	}
L68:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v378 = int32(base.Ui32(v375) >> (uint(int32(2)) % 32))
	goto L67
L69:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v373 = F_accumArrayResult(m, v203, int32(0), int32(1), int32(17), v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L129
	}
L70:
	;
	if v59 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+int32(base.Ui32(v196)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v218)>>(uint(v196&int32(7))%32))&int32(1) == int32(0) {
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v228 = v141 + int32(20) + v196<<(uint(int32(4))%32)
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+4)))
	v231 = base.B2i32(v229 != int32(65535))
	if v231 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	if v189 < v311+v308 {
		goto L15
	} else {
		goto L101
	}
L76:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v187))))
	if v235 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+12)))
	v308 = (v198 + v299 - int32(1)) & (int32(0) - v299)
	v311 = base.I32_extend16_s(v229)
	goto L75
L79:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+12)))
	v244 = (v198 + v238 - int32(1)) & (int32(0) - v238)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v244))))
	v247 = v246
	v248 = v244
	goto L81
L80:
	;
	v247 = v235
	v248 = v198
	goto L81
L81:
	;
	v249 = v248 + v187
	v251 = v247 & int32(255)
	if v251 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	switch v254 - int32(1) {
	case 0, 17:
		goto L85
	default:
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v292 = int32(1)
	if v247&v292 != 0 {
		v308 = v248
		v311 = int32(base.Ui32(v251) >> (uint(v292) % 32))
		goto L75
	} else {
		goto L100
	}
L85:
	;
	v277 = int32(6)
	v279 = int32(18)
	if v254 == v279 {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v196
	F_errmsg(m, int32(497489), v21)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(518937), int32(383), int32(326112))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v283 = v279
	goto L93
L92:
	;
	v283 = int32(2)
	goto L93
L93:
	;
	if v254&int32(254) == int32(2) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v288 = v277
	goto L96
L95:
	;
	v288 = v283
	goto L96
L96:
	;
	if v254 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v291 = v277
	goto L99
L98:
	;
	v291 = v288
	goto L99
L99:
	;
	v308 = v248
	v311 = v291
	goto L75
L100:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v308 = v248
	v311 = int32(base.Ui32(v296) >> (uint(int32(2)) % 32))
	goto L75
L101:
	;
	if v231|(v50^int32(1)) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v334 = int32(*(*int16)(unsafe.Add(mBase, uint32(v228)+4)))
	if int32(0) < v334 {
		v378 = v334
		goto L67
	} else {
		goto L112
	}
L103:
	;
	v318 = F_pg_detoast_datum_copy(m, v308+v187)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v321 = v311 + int32(4)
	v322 = F_palloc(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	v332 = v318
	goto L102
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v321 << (uint(int32(2)) % 32)
	if v311 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v332 = v322
	goto L102
L109:
	;
	v330 = F__emscripten_memcpy_bulkmem(m, v322+int32(4), v308+v187, v311)
	mBase = m.M
	goto L111
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	v337 = v308 + v187
	if v334 == int32(-1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	if v340 == int32(1) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v365 = F_strlen(m, v337)
	mBase = m.M
	v378 = v365 + int32(1)
	goto L67
L116:
	;
	v343 = int32(6)
	v345 = int32(18)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	if v347 == v345 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	if v340&int32(1) == int32(0) {
		goto L68
	} else {
		goto L128
	}
L119:
	;
	v350 = v345
	goto L121
L120:
	;
	v350 = int32(2)
	goto L121
L121:
	;
	if v347&int32(254) == int32(2) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v355 = v343
	goto L124
L123:
	;
	v355 = v350
	goto L124
L124:
	;
	if v347 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v358 = v343
	goto L127
L126:
	;
	v358 = v355
	goto L127
L127:
	;
	v378 = v358
	goto L67
L128:
	;
	v378 = int32(base.Ui32(v340) >> (uint(int32(1)) % 32))
	goto L67
L129:
	;
	v392 = v198
	v396 = v373
	goto L66
L130:
	;
	if v332 == int32(0) {
		v392 = v380
		v396 = v385
		goto L66
	} else {
		goto L131
	}
L131:
	;
	F_pfree(m, v332)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v392 = v380
	v396 = v385
	goto L66
L133:
	;
	goto L65
L134:
	;
	F_relation_close(m, v139, int32(1))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v425 = F_makeArrayResult(m, v408, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	if v129 == int32(0) {
		v431 = v425
		goto L25
	} else {
		goto L137
	}
L137:
	;
	F_pfree(m, v129)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v431 = v425
	goto L25
L139:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(149801), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(518937), int32(460), int32(109258))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(558586), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(518937), int32(478), int32(109258))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v68
	F_errmsg(m, int32(39153), v19+int32(-16))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(518937), int32(485), int32(109258))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v523 = F_pg_mblen_cstr(m, v95)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v523
	F_errmsg(m, int32(347283), v19+int32(-32))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(518937), int32(104), int32(131885))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v548 = F_strlen(m, v43)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v548
	F_errmsg(m, int32(344503), v19+int32(-48))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(518937), int32(496), int32(109258))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(465291), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(518937), int32(338), int32(326112))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	F_errmsg(m, int32(530902), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(518937), int32(396), int32(326112))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(530259), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(518937), int32(420), int32(326112))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
