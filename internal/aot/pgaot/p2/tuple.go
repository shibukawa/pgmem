package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescTruncatedCopy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v9 = F_palloc(m, l1*int32(116)+int32(20))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(-4294965047)
	v19 = l1 * int32(100)
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = int32(4)
	v23 = int32(20)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v9+l1<<(uint(v20)%32)+v23, l0+v25<<(uint(v20)%32)+v23, v19)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if int32(0) < v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = int32(0)
	v39 = v32
	goto L9
L7:
	;
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v62
	return v9
L9:
	;
	v45 = v9 + v39<<(uint(int32(4))%32) + v37*int32(100)
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+110)) = uint8(v46)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+106)) = v46
	F_populate_compact_attribute(m, v9, v37)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v53 = v37 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v53 < v54 {
		v37 = v53
		v39 = v54
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
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
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTupleTransactionInfo[0])))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = F_palloc(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 != int32(2249) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	v32 = F_palloc0(m, v15*int32(28))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) <= v24 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	F_assign_record_type_typmod(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v35 = v15 << (uint(int32(2)) % 32)
	v36 = F_palloc0(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v38 = F_palloc0(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v15 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v32
	m.G0 = v13 + int32(16)
	return v17
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = l0 + v53<<(uint(int32(4))%32) + v44*int32(100)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+111)))
	if v60 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v64 = v59 + int32(20)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+68))
	v69 = v44 << (uint(int32(2)) % 32)
	F_getTypeInputInfo(m, v65, v13+int32(12), v36+v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v85 = v44 + int32(1)
	if v85 != v15 {
		v44 = v85
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_fmgr_info(m, v73, v32+v44*int32(28))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v38+v69))) = v80
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
	var v42 int32
	_ = v42
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
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v16 = int32(0)
	goto L8
L5:
	;
	v42 = int32(0)
	goto L6
L6:
	;
	return v42
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
	if v24 == l1&int32(_a_F_TupleDescGetDefault_0) {
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
	v42 = v32
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
	var v130 int32
	_ = v130
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
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
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
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L160
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L156
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L152
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L148
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L143
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L139
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L135
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
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L131
	}
L25:
	;
	m.G0 = v21 - int32(-64)
	return v424
L26:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
	v424 = int32(0)
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
		v130 = v72
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
	v130 = v72
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
	v130 = v2
	goto L29
L48:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+52))
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
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
	if v396 != v403 {
		goto L14
	} else {
		goto L126
	}
L55:
	;
	v183 = v32 + int32(4)
	v185 = v162 & int32(_a_F_tuple_data_split_0)
	v188 = int32(0)
	v192 = v188
	v194 = v188
	v199 = v146
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
	v396 = int32(0)
	v401 = v146
	v403 = v162 & int32(_a_F_tuple_data_split_0)
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
	F_errmsg(m, int32(_a_F_tuple_data_split_1), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(343), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
	if base.Ui32(v157) <= base.Ui32(v192) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v396 = v385
	v401 = v388
	v403 = v185
	goto L54
L66:
	;
	v392 = v192 + int32(1)
	if v392 != v148 {
		v192 = v392
		v194 = v385
		v199 = v388
		goto L64
	} else {
		goto L125
	}
L67:
	;
	v373 = v371 + v302
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
	v378 = F_accumArrayResult(m, v199, v327, int32(0), int32(17), v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L122
	}
L68:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v371 = int32(base.Ui32(v368) >> (uint(int32(2)) % 32))
	goto L67
L69:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
	v366 = F_accumArrayResult(m, v199, int32(0), int32(1), int32(17), v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L121
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
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+int32(base.Ui32(v192)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v214)>>(uint(v192&int32(7))%32))&int32(1) == int32(0) {
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v224 = v141 + int32(20) + v192<<(uint(int32(4))%32)
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+4)))
	v227 = base.B2i32(v225 != int32(_a_F_tuple_data_split_0))
	if v227 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	if v185 < v305+v302 {
		goto L15
	} else {
		goto L99
	}
L76:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v183))))
	if v231 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+12)))
	v302 = (v194 + v293 - int32(1)) & (int32(0) - v293)
	v305 = base.I32_extend16_s(v225)
	goto L75
L79:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+12)))
	v240 = (v194 + v234 - int32(1)) & (int32(0) - v234)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v240))))
	v243 = v242
	v244 = v240
	goto L81
L80:
	;
	v243 = v231
	v244 = v194
	goto L81
L81:
	;
	v246 = v244 + v183
	v248 = v243 & int32(255)
	if v248 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v253 = v251 - int32(1)
	v254 = int32(0)
	if base.B2i32(v253 == v254)|base.B2i32(v253 == int32(17)) == v254 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v286 = int32(1)
	if v243&v286 != 0 {
		v302 = v244
		v305 = int32(base.Ui32(v248) >> (uint(v286) % 32))
		goto L75
	} else {
		goto L98
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v278 = int32(18)
	if v251 == v278 {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v192
	F_errmsg(m, int32(_a_F_tuple_data_split_4), v21)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(383), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v282 = v278
	goto L94
L93:
	;
	v282 = int32(2)
	goto L94
L94:
	;
	if base.Ui32(v251) < base.Ui32(int32(4)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v285 = int32(6)
	goto L97
L96:
	;
	v285 = v282
	goto L97
L97:
	;
	v302 = v244
	v305 = v285
	goto L75
L98:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v302 = v244
	v305 = int32(base.Ui32(v290) >> (uint(int32(2)) % 32))
	goto L75
L99:
	;
	if base.B2i32(v50 == v188)|v227 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v329 = int32(*(*int16)(unsafe.Add(mBase, uint32(v224)+4)))
	if int32(0) < v329 {
		v371 = v329
		goto L67
	} else {
		goto L107
	}
L101:
	;
	v312 = F_pg_detoast_datum_copy(m, v302+v183)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v315 = v305 + int32(4)
	v316 = F_palloc(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L105
	}
L104:
	;
	v327 = v312
	goto L100
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v315 << (uint(int32(2)) % 32)
	if v305 == int32(0) {
		v327 = v316
		goto L100
	} else {
		goto L106
	}
L106:
	;
	base.MemoryCopy(m, v316+int32(4), v302+v183, v305)
	v327 = v316
	goto L100
L107:
	;
	v332 = v302 + v183
	if v329 == int32(-1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v335 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	v358 = F_strlen(m, v332)
	mBase = m.M
	v371 = v358 + int32(1)
	goto L67
L111:
	;
	v339 = int32(18)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	if v341 == v339 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	if v335&int32(1) == int32(0) {
		goto L68
	} else {
		goto L120
	}
L114:
	;
	v344 = v339
	goto L116
L115:
	;
	v344 = int32(2)
	goto L116
L116:
	;
	if base.Ui32((v341-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v351 = int32(6)
	goto L119
L118:
	;
	v351 = v344
	goto L119
L119:
	;
	v371 = v351
	goto L67
L120:
	;
	v371 = int32(base.Ui32(v335) >> (uint(int32(1)) % 32))
	goto L67
L121:
	;
	v385 = v194
	v388 = v366
	goto L66
L122:
	;
	if v327 == int32(0) {
		v385 = v373
		v388 = v378
		goto L66
	} else {
		goto L123
	}
L123:
	;
	F_pfree(m, v327)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v385 = v373
	v388 = v378
	goto L66
L125:
	;
	goto L65
L126:
	;
	F_relation_close(m, v139, int32(1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_tuple_data_split[0]))
	v418 = F_makeArrayResult(m, v401, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	if v130 == int32(0) {
		v424 = v418
		goto L25
	} else {
		goto L129
	}
L129:
	;
	F_pfree(m, v130)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v424 = v418
	goto L25
L131:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_5), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(460), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_7), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(478), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v68
	F_errmsg(m, int32(_a_F_tuple_data_split_8), v19+int32(-16))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(485), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
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
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	v504 = F_pg_mblen_cstr(m, v95)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v504
	F_errmsg(m, int32(_a_F_tuple_data_split_9), v19+int32(-32))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(104), int32(_a_F_tuple_data_split_10))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v525 = F_strlen(m, v43)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v525
	F_errmsg(m, int32(_a_F_tuple_data_split_11), v19+int32(-48))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(496), int32(_a_F_tuple_data_split_6))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_12), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(338), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
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
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_13), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(396), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_tuple_data_split_14), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_tuple_data_split_2), int32(420), int32(_a_F_tuple_data_split_3))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
