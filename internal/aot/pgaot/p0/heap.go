package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleHeaderGetDatum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v3&int32(4) == int32(0) {
		v24 = l0
		return v24
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = F_lookup_rowtype_tupdesc(m, v8, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = F_toast_flatten_tuple_to_datum(m, l0, int32(base.Ui32(v14)>>(uint(int32(2))%32)), v10)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				if v19 < int32(0) {
					v24 = v17
					return v24
				} else {
					F_DecrTupleDescRefCount(m, v10)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = v17
						return v24
					}
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
	var v33 int32
	_ = v33
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
			if base.B2i32(base.Ui32(v18) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v19)) == v17 {
				v31 = base.B2i32(base.Ui32(v19) < base.Ui32(l1))
			} else {
				v31 = int32(base.Ui32(v19-l1) >> (uint(int32(31)) % 32))
			}
			if v31 != 0 {
				v32 = v17
			} else {
				v32 = v18
			}
			v33 = v32
		} else {
			v33 = v11
		}
		m.G0 = v7 + int32(16)
		return v33
	}
}
func F_heap_create_with_catalog(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32 {
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v334 int32
	_ = v334
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
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
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L157
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L153
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L149
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L146
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L141
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
	v44 = int32(0)
	v46 = F_GetSysCacheOid(m, int32(81), l0, l1, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
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
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L137
	}
L13:
	;
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = F_moveArrayTypeName(m, v46, l0, l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	if v48 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v55 = base.B2i32(l2 != int32(1664))
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if v55 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v56 = int32(0)
	if l3 != 0 {
		v105 = l3
		v106 = v56
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_LockRelationOid(m, v105, int32(8))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[0])))
	if v58 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = F_GetNewRelFileNumber(m, l2, v33, l11)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
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
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[1]))
	if v64 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[2]))
	if v78 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[1])) = v68
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[3]))
	if v71 == v68 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[3])) = int32(0)
	v105 = v64
	v106 = v71
	goto L23
L32:
	;
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[2])) = v82
	v85 = l10 - int32(83)
	if base.B2i32(base.Ui32(int32(31)) < base.Ui32(v85))|base.B2i32(int32(1)<<(uint(v85)%32)&int32(-2076180479) == v82) != 0 {
		v105 = v78
		v106 = v56
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[4]))
	if v96 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[4])) = int32(0)
	v105 = v78
	v106 = v96
	goto L23
L35:
	;
	v105 = v103
	v106 = v56
	goto L23
L36:
	;
	if l16 == int32(0) {
		v132 = v22
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+104)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v228)+96)) = int64(-4647714815446351872)
	if l10 == int32(83) {
		goto L56
	} else {
		goto L57
	}
L38:
	;
	v138 = F_heap_create(m, l0, l1, l2, v105, v106, l7, l8, l10, l11, l12, l13, l17, v29+int32(52), v29+int32(48), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
		v132 = v22
		goto L38
	case 19, 26, 29, 31, 35:
		goto L40
	}
L40:
	;
	v130 = F_get_user_default_acl(m, int32(41), l6, l1)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v116 = F_get_user_default_acl(m, int32(37), l6, l1)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v124 = F_heap_create(m, l0, l1, l2, v105, v106, l7, l8, int32(83), l11, l12, l13, l17, v29+int32(52), v29+int32(48), int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+132)) = l19
	v221 = int32(0)
	v224 = v124
	v225 = v116
	goto L37
L44:
	;
	v132 = v130
	goto L38
L45:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+132)) = l19
	v142 = int32(0)
	switch l10 - int32(73) {
	case 0, 10:
		v221 = v142
		v224 = v138
		v225 = v132
		goto L37
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L46
	default:
		goto L47
	}
L46:
	;
	v152 = int32(0)
	v164 = F_AssignTypeArrayOid(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	switch l10 - int32(105) {
	case 0, 11:
		v221 = v142
		v224 = v138
		v225 = v132
		goto L37
	default:
		goto L46
	}
L48:
	;
	v166 = int32(0)
	F_TypeCreate(m, v29+int32(68), l4, l0, l1, v105, l10, l6, int32(-1), int32(99), int32(67), v152, int32(44), int32(2290), int32(2291), int32(2402), int32(2403), v152, v152, v152, v152, v152, v152, v164, v166, v166, v166, v166, int32(100), int32(120), int32(-1), v166, v166, v166)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if l20 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l20)+8)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(l20)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(l20))) = v179
	goto L52
L51:
	;
	goto L52
L52:
	;
	v187 = F_makeArrayTypeName(m, l0, l1)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v189 = int32(0)
	v191 = int32(-1)
	F_TypeCreate(m, v29+int32(68), v164, v187, l1, v189, v189, l6, v191, int32(98), int32(65), v189, int32(44), int32(750), int32(751), int32(2400), int32(2401), v189, v189, int32(3816), int32(_a_F_heap_create_with_catalog_0), v178, int32(1), v189, v189, v189, v189, v189, int32(100), int32(120), v191, v189, v189, v189)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v187)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v221 = v178
	v224 = v138
	v225 = v132
	goto L37
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v228)+96)) = int64(4575657221408423937)
	goto L58
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+140)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v228)+136)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v228)+80)) = l6
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+131)) = uint8(v240)
	*(*int32)(unsafe.Add(mBase, uint32(v228)+76)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v228)+72)) = v221
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
	if v221 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v246 = v221
	goto L61
L60:
	;
	v246 = int32(2249)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+4)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = int32(-1)
	F_InsertPgClassTuple(m, v33, v224, v105, v225, l15)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v257 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v260 = F_CatalogOpenIndexes(m, v257)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_InsertPgAttributeTuples(m, v257, v253, v105, int32(0), v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if int32(0) < v254 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v267 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v363 = l10 - int32(99)
	v364 = int32(0)
	if base.B2i32(v363 == v364)|base.B2i32(v363 == int32(19)) == v364 {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v295 = v267 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(1247)
	v307 = v253 + v293<<(uint(int32(4))%32) + v267*int32(100)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v308
	v313 = v29 + int32(68)
	v315 = v29 + int32(56)
	F_recordDependencyOn(m, v313, v315, int32(110))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v307)+116))
	v320 = int32(0)
	if base.B2i32(v319 == v320)|base.B2i32(v319 == int32(100)) == v320 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(3456)
	F_recordDependencyOn(m, v313, v315, int32(110))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v295 != v254 {
		v267 = v295
		goto L69
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	goto L70
L77:
	;
	v373 = F_CreateTupleDesc(m, int32(6), int32(_a_F_heap_create_with_catalog_1))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_CatalogCloseIndexes(m, v260)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	F_InsertPgAttributeTuples(m, v257, v373, v105, int32(0), v260)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_FreeTupleDesc(m, v373)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	F_relation_close(m, v257, int32(3))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[5]))
	if v387 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[6]))
	if v457 != 0 {
		goto L104
	} else {
		goto L105
	}
L86:
	;
	v391 = l10 - int32(99)
	if base.B2i32(v391 == int32(0))|base.B2i32(v391 == int32(17)) != 0 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v105
	v400 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v400
	F_recordDependencyOnOwner(m, v400, v105, l6)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_recordDependencyOnNewAcl(m, int32(1259), v105, l6, v225)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_recordDependencyOnCurrentExtension(m, v29+int32(68), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v413 = F_new_object_addresses(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(2615)
	v421 = v29 + int32(56)
	F_add_exact_object_address(m, v421, v413)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if l5 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(1247)
	F_add_exact_object_address(m, v421, v413)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	switch l10 - int32(109) {
	case 0, 5:
		goto L98
	default:
		goto L99
	}
L96:
	;
	goto L95
L97:
	;
	F_record_object_address_dependencies(m, v29+int32(68), v413, int32(110))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L102
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(2601)
	F_add_exact_object_address(m, v29+int32(56), v413)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	if base.B2i32(l7 == int32(0))|base.B2i32(l10 != int32(112)) != 0 {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	goto L97
L102:
	;
	F_free_object_addresses(m, v413)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L85
L104:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v105, int32(0), l18)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if l9 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l14) {
		goto L126
	} else {
		goto L127
	}
L109:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v466 <= int32(0) {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v469 = int32(0)
	v473 = v469
	v488 = v469
	goto L112
L112:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l9)+12))
	v498 = int32(2)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v497+v473<<(uint(v498)%32))))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	switch v502 - v498 {
	case 0:
		goto L115
	default:
		goto L116
	case 3:
		goto L117
	}
L113:
	;
	if v540 <= int32(0) {
		goto L108
	} else {
		goto L124
	}
L114:
	;
	v542 = v473 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v542 < v543 {
		v473 = v542
		v488 = v540
		goto L112
	} else {
		goto L123
	}
L115:
	;
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v501)+12)))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v501)+16))
	v537 = F_StoreAttrDefault(m, v224, v535, v536, l18)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L122
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501)+8))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v501)+16))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+20)))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+21)))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+22)))
	v514 = int32(*(*int16)(unsafe.Add(mBase, uint32(v501)+24)))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+26)))
	v516 = F_StoreRelCheck(m, v224, v505, v506, v507, (v508^int32(-1))&int32(1), v513, v514, v515, l18)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v516
	v540 = v488 + int32(1)
	goto L114
L119:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v525
	F_errmsg_internal(m, int32(_a_F_heap_create_with_catalog_2), v29)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(2346), int32(_a_F_heap_create_with_catalog_4))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v537
	v540 = v488
	goto L114
L123:
	;
	goto L113
L124:
	;
	F_SetRelationNumChecks(m, v224, v540)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L108
L126:
	;
	v578 = l14
	goto L128
L127:
	;
	v578 = int32(0)
	goto L128
L128:
	;
	if v578 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v579 = int32(_a_F_heap_create_with_catalog_5)
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[7]))
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[7])) = v583
	v586 = F_palloc(m, int32(16))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	F_relation_close(m, v224, int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L135
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+4)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = v105
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[9]))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v586)+8)) = v592
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[10]))
	v598 = F_lcons(m, v586, v597)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[7])) = v580
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create_with_catalog[10])) = v598
	goto L131
L135:
	;
	F_relation_close(m, v33, int32(3))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	m.G0 = v29 + int32(80)
	return v105
L137:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = l0
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_6), v29+int32(32))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1179), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(_a_F_heap_create_with_catalog_8))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l0
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_9), v29+int32(16))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errhint(m, int32(_a_F_heap_create_with_catalog_10), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1198), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errmsg_internal(m, int32(_a_F_heap_create_with_catalog_11), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1205), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_12), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1236), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_13), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1247), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_heap_create_with_catalog_14), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_heap_create_with_catalog_3), int32(1257), int32(_a_F_heap_create_with_catalog_7))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
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
										F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(240), int32(0))
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
									F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(240), int32(0))
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
									F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(240), int32(0))
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
								F_array_desc(m, l0, v11+int32(12), int32(4), v139, int32(240), int32(0))
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
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
	var v96 int32
	_ = v96
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
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	v19 = m.G0
	v21 = v19 - int32(272)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = int32(1)
	v30 = F_toast_open_indexes(m, l0, v24, v21+int32(268), v21+int32(108))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v21+int32(112), int32(1), int32(3), int32(184), l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = int32(1996)
	v40 = base.I32_div_u_s(l3, v39)
	v41 = int32(1)
	v44 = base.I32_div_u_s(l2-v41, v39)
	v47 = l3 + l4 - v41
	v49 = base.I32_div_u_s(v47, v39)
	if base.B2i32(v44 == v49)&base.B2i32(base.Ui32(l3) <= base.Ui32(int32(1995))) != 0 {
		v77 = v24
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v30<<(uint(int32(2))%32))))
	v83 = F_get_toast_snapshot(m)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L18
	}
L5:
	;
	v55 = v21 + int32(160)
	if v40 == v49 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v57 = int32(2)
	F_ScanKeyInit(m, v55, v57, int32(3), int32(65), v40)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_ScanKeyInit(m, v55, int32(2), int32(4), int32(150), v40)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v77 = v57
	goto L4
L10:
	;
	v70 = int32(2)
	F_ScanKeyInit(m, v21+int32(208), v70, v70, int32(149), v49)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v77 = int32(3)
	goto L4
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L90
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L86
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L82
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L78
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L75
	}
L17:
	;
	if v234 != v49+int32(1) {
		goto L14
	} else {
		goto L72
	}
L18:
	;
	v87 = F_systable_beginscan_ordered(m, l0, v82, v83, v77, v21+int32(112))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v90 = F_systable_getnext_ordered(m, v87, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v90 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v234 = v40
	goto L17
L22:
	;
	goto L23
L23:
	;
	v96 = v21 + int32(107)
	v97 = F_fastgetattr_1(m, v90, int32(2), v23, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v100 = F_fastgetattr_1(m, v90, int32(3), v23, v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	if v40 != v97 {
		v414 = v40
		v415 = v97
		goto L12
	} else {
		goto L31
	}
L26:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v102&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = int32(1)
	if v102&v105 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v114 = int32(4)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v120 = v114
	v121 = int32(base.Ui32(v115)>>(uint(int32(2))%32)) - v114
	goto L25
L30:
	;
	v110 = int32(1)
	v120 = v105
	v121 = int32(base.Ui32(v102)>>(uint(v110)%32)) - v110
	goto L25
L31:
	;
	if base.Ui32(v49) < base.Ui32(v40) {
		v301 = v40
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v127 = v44*int32(-1996) + l2
	if base.Ui32(v40) < base.Ui32(v44) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v129 = int32(1996)
	goto L35
L34:
	;
	v129 = v127
	goto L35
L35:
	;
	if v121 != v129 {
		v362 = v121
		v366 = v40
		v367 = v129
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v133 = l5 - l3 + int32(4)
	v136 = v47 - v49*int32(1996)
	if v40 == v49 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = v136
	goto L39
L38:
	;
	v140 = v121 - int32(1)
	goto L39
L39:
	;
	v142 = v40 * int32(1996)
	v143 = l3 - v142
	v146 = v140 - v143 + int32(1)
	if v146 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryCopy(m, v133+v142+v143, v100+v120+v143, v146)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v152 = int32(1)
	v153 = v40 + v152
	v155 = F_systable_getnext_ordered(m, v87, v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v155 == int32(0) {
		v234 = v153
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v165 = v155
	v168 = v153
	goto L45
L45:
	;
	v181 = v21 + int32(107)
	v182 = F_fastgetattr_1(m, v165, int32(2), v23, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v234 = v223
	goto L17
L47:
	;
	v185 = F_fastgetattr_1(m, v165, int32(3), v23, v181)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	if v168 != v182 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v187&int32(3) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v195 = int32(4)
	v207 = int32(base.Ui32(v192)>>(uint(int32(2))%32)) - v195
	v208 = v195
	goto L48
L51:
	;
	goto L52
L52:
	;
	if v187&int32(1) == int32(0) {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v202 = int32(1)
	v207 = int32(base.Ui32(v187)>>(uint(v202)%32)) - v202
	v208 = v202
	goto L48
L54:
	;
	v414 = v168
	v415 = v182
	goto L12
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v49) < base.Ui32(v168) {
		v301 = v168
		goto L15
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v168) < base.Ui32(v44) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v213 = int32(1996)
	goto L60
L59:
	;
	v213 = v127
	goto L60
L60:
	;
	if v213 != v207 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v362 = v207
	v366 = v168
	v367 = v213
	goto L13
L62:
	;
	goto L63
L63:
	;
	if v168 == v49 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v216 = v136 + int32(1)
	goto L66
L65:
	;
	v216 = v207
	goto L66
L66:
	;
	if v216 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	base.MemoryCopy(m, v133+v168*int32(1996), v185+v208, v216)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v222 = int32(1)
	v223 = v168 + v222
	v225 = F_systable_getnext_ordered(m, v87, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v225 != 0 {
		v165 = v225
		v168 = v223
		goto L45
	} else {
		goto L71
	}
L71:
	;
	goto L46
L72:
	;
	F_systable_endscan_ordered(m, v87)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	F_toast_close_indexes(m, v250, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	m.G0 = v21 + int32(272)
	return
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v279 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_0), v21+int32(96))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(729), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
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
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v319 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v301
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_3), v21+int32(16))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(749), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
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
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v344 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_4), v21)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(786), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
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
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v383 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v44 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v362
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_5), v21+int32(48))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(758), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v415
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v431 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_fetch_toast_slice_6), v21+int32(80))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_heap_fetch_toast_slice_1), int32(742), int32(_a_F_heap_fetch_toast_slice_2))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
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
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v68 int32
	_ = v68
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
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
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
	v189 = m.ExcPending
	if v189 != 0 {
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
	v25 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	v68 = v4
	v70 = v4
	v72 = int32(24)
	goto L4
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v25))))
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
	v68 = int32(1)
	v70 = int32(128)
	v72 = (int32(base.Ui32(v16+int32(7))>>(uint(int32(3))%32)) + int32(30)) & int32(536870904)
	goto L4
L11:
	;
	goto L12
L12:
	;
	v47 = v25 + int32(1)
	if v47 != v16 {
		v25 = v47
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
	if v68 != 0 {
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
	if v68 != 0 {
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
	v137 = int32(0)
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
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1+v137<<(uint(int32(2))%32))))
	v161 = v159
	goto L33
L32:
	;
	v161 = int32(0)
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
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v137))))
	v165 = v163
	goto L36
L35:
	;
	v165 = int32(1)
	goto L36
L36:
	;
	F_fill_val(m, l0+int32(20)+v137<<(uint(int32(4))%32), v151, v14+int32(20), v14+int32(28), v80+int32(44), v161, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v169 = v137 + int32(1)
	if v169 != v111 {
		v137 = v169
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
	v192 = m.ExcPending
	if v192 != 0 {
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
	v198 = m.ExcPending
	if v198 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_heap_form_tuple_3), int32(1134), int32(_a_F_heap_form_tuple_4))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
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
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13924(m, l0, l1, l2, l3, int32(_a_F_heap_getattr_1_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+188))
	if v6 == int32(_a_F_heap_getnext_0) {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_heap_getnext[0]))
		if v10 != 0 {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_getnext[1])))
			if v12&int32(1) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_heap_getnext_1), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_heap_getnext_2), int32(1362), int32(_a_F_heap_getnext_3))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				if v19&int32(1) != 0 {
					F_heapgettup_pagemode(m, l0, int32(1), v18, v17)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						if v30 == int32(0) {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
							if v38 == int32(0) {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
								if v41 != int32(1) {
									return l0 - int32(-64)
								} else {
									F_pgstat_assoc_relation(m, v37)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
										v48 = v47
										v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
										return l0 - int32(-64)
									}
								}
							} else {
								v48 = v38
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
								return l0 - int32(-64)
							}
						}
					}
				} else {
					F_heapgettup(m, l0, int32(1), v18, v17)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						if v30 == int32(0) {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
							if v38 == int32(0) {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
								if v41 != int32(1) {
									return l0 - int32(-64)
								} else {
									F_pgstat_assoc_relation(m, v37)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
										v48 = v47
										v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
										return l0 - int32(-64)
									}
								}
							} else {
								v48 = v38
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
								return l0 - int32(-64)
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
			if v19&int32(1) != 0 {
				F_heapgettup_pagemode(m, l0, int32(1), v18, v17)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					if v30 == int32(0) {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
						if v38 == int32(0) {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
							if v41 != int32(1) {
								return l0 - int32(-64)
							} else {
								F_pgstat_assoc_relation(m, v37)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
									v48 = v47
									v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
									return l0 - int32(-64)
								}
							}
						} else {
							v48 = v38
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
							return l0 - int32(-64)
						}
					}
				}
			} else {
				F_heapgettup(m, l0, int32(1), v18, v17)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					if v30 == int32(0) {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
						if v38 == int32(0) {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
							if v41 != int32(1) {
								return l0 - int32(-64)
							} else {
								F_pgstat_assoc_relation(m, v37)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
									v48 = v47
									v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
									return l0 - int32(-64)
								}
							}
						} else {
							v48 = v38
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49 + int64(1)
							return l0 - int32(-64)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_heap_getnext_4), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_heap_getnext_2), int32(1352), int32(_a_F_heap_getnext_3))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v204 int32
	_ = v204
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	v7 = l6
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
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
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_heap_hot_search_buffer[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(l2^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L1
L3:
	;
	goto L4
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_heap_hot_search_buffer[1]))
	v42 = v36 + l2<<(uint(int32(13))%32) + int32(-8192)
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
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v44 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v23 + int32(16)
	return v310
L9:
	;
	v310 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v62 = v7
	v63 = v44
	v66 = v8
	v70 = v8
	v71 = v7 ^ int32(1)
	goto L12
L12:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+12)))
	if base.Ui32(v76) < base.Ui32(int32(25)) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v310 = v303
	goto L8
L14:
	;
	v310 = int32(0)
	goto L8
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(int32(base.Ui32(v76+int32(_a_F_heap_hot_search_buffer_0))>>(uint(int32(2))%32))&int32(_a_F_heap_hot_search_buffer_1)) < base.Ui32(v63) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v310 = int32(0)
	goto L8
L18:
	;
	goto L19
L19:
	;
	v90 = v42 + int32(20) + v63<<(uint(int32(2))%32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v95 = int32(base.Ui32(v91)>>(uint(int32(15))%32)) & int32(3)
	if v95 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v303 = int32(0)
	if v290 != 0 {
		v62 = v303
		v63 = v290
		v66 = v293
		v70 = v297
		v71 = v298
		goto L12
	} else {
		goto L83
	}
L21:
	;
	if (v62^int32(-1)|base.B2i32(v95 != int32(2)))&int32(1) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v110 = v42 + v91&int32(_a_F_heap_hot_search_buffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(base.Ui32(v112) >> (uint(int32(17)) % 32))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+8)) = uint16(v63)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)) = uint16(v48)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)) = uint16(v49)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v116
	if v62&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v310 = int32(0)
	goto L8
L25:
	;
	goto L26
L26:
	;
	v290 = v91 & int32(_a_F_heap_hot_search_buffer_2)
	v293 = v66
	v297 = v70
	v298 = v71
	goto L20
L27:
	;
	v123 = int32(0)
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+18)))
	if v124 < v123 {
		v310 = v123
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v66 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	if v71&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+20)))
	v131 = int32(768)
	if v130&v131 != v131 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v137 = v135
	goto L35
L34:
	;
	v137 = int32(2)
	goto L35
L35:
	;
	if v137 == v66 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v310 = int32(0)
	goto L8
L37:
	;
	if l5 == int32(0) {
		v214 = v70
		goto L48
	} else {
		goto L49
	}
L38:
	;
	v142 = F_HeapTupleSatisfiesVisibility(m, l4, l3, l2)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	F_HeapCheckForSerializableConflictOut(m, v142, l1, l4, l2, l3)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v142 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v63)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+20)))
	v153 = int32(768)
	if v152&v153 != v153 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v159 = v157
	goto L45
L44:
	;
	v159 = int32(2)
	goto L45
L45:
	;
	F_PredicateLockTID(m, l1, l4+int32(4), l3, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v162 = int32(1)
	if l5 == int32(0) {
		v310 = v162
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v165)
	v310 = v162
	goto L8
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+19)))
	if v216&int32(64) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L49:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v170 != int32(1) {
		v214 = v70
		goto L48
	} else {
		goto L50
	}
L50:
	;
	if v70 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v175 = F_GlobalVisHorizonKindForRel(m, l1)
	mBase = m.M
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175<<(uint(int32(2))%32))+uint32(_c_F_heap_hot_search_buffer[2])))
	goto L54
L52:
	;
	v179 = v70
	goto L53
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+20)))
	if v182&int32(256) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v179 = v178
	goto L53
L55:
	;
	if v208 != 0 {
		v214 = v179
		goto L48
	} else {
		goto L63
	}
L56:
	;
	v208 = int32(base.Ui32(v182&int32(512)) >> (uint(int32(9)) % 32))
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v182&int32(2048)|base.B2i32(v182&int32(_a_F_heap_hot_search_buffer_3) != int32(1024))|base.B2i32(v182&int32(_a_F_heap_hot_search_buffer_4) == int32(64)) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v207 = int32(0)
	goto L61
L60:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v205 = F_GlobalVisTestIsRemovableXid(m, v179, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L39
	} else {
		goto L62
	}
L61:
	;
	v208 = v207
	goto L55
L62:
	;
	v207 = v205
	goto L61
L63:
	;
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v209)
	v214 = v179
	goto L48
L64:
	;
	v310 = int32(0)
	goto L8
L65:
	;
	goto L66
L66:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+20)))
	if v222&int32(2048) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v310 = int32(0)
	goto L8
L68:
	;
	goto L69
L69:
	;
	if v222&int32(768) == int32(512) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v310 = int32(0)
	goto L8
L71:
	;
	goto L72
L72:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+16)))
	v233 = int32(0)
	if v222&int32(_a_F_heap_hot_search_buffer_5) != int32(_a_F_heap_hot_search_buffer_6) {
		v290 = v232
		v293 = v231
		v297 = v214
		v298 = v233
		goto L20
	} else {
		goto L73
	}
L73:
	;
	v241 = F_GetMultiXactIdMembers(m, v231, v23+int32(12), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L39
	} else {
		goto L74
	}
L74:
	;
	v243 = int32(0)
	if v241 <= v243 {
		v290 = v232
		v293 = v243
		v297 = v214
		v298 = v233
		goto L20
	} else {
		goto L75
	}
L75:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v254 = int32(0)
	goto L78
L76:
	;
	F_pfree(m, v247)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L39
	} else {
		goto L82
	}
L77:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v280 = v278
	goto L76
L78:
	;
	v270 = v247 + v254<<(uint(int32(3))%32)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v271) {
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v280 = int32(0)
	goto L76
L80:
	;
	v275 = v254 + int32(1)
	if v275 != v241 {
		v254 = v275
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v290 = v232
	v293 = v280
	v297 = v214
	v298 = v233
	goto L20
L83:
	;
	goto L13
}
func F_heap_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	v21 = F_heap_prepare_insert(m, l0, l1, v17, l2, l3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = int32(0)
	v29 = F_RelationGetBufferForTuple(m, l0, v23, v24, l3, l4, v15+int32(12), v24, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = int32(_a_F_heap_insert_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_insert[0])) = v37 + int32(1)
	v42 = l3 & int32(16)
	F_RelationPutHeapTuple(m, v29, v21, int32(base.Ui32(v42)>>(uint(int32(4))%32)))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v29 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_MarkBufferDirty(m, v29)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L15
	}
L8:
	;
	v77 = v74 & int32(_a_F_heap_insert_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+10)) = uint16(v77)
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v86 = F_visibilitymap_clear(m, v79|v80<<(uint(int32(16))%32), v84, int32(3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[1]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v29^int32(-1))<<(uint(int32(2))%32))))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+10)))
	if v57&int32(4) != 0 {
		v74 = v57
		v75 = v56
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[3]))
	v64 = v61 + v29<<(uint(int32(13))%32)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64-int32(_a_F_heap_insert_6)))))
	if v67&int32(4) == int32(0) {
		v91 = v6
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v91 = v6
	goto L7
L13:
	;
	v74 = v67
	v75 = v64 + int32(-8192)
	goto L8
L14:
	;
	v91 = int32(1)
	goto L7
L15:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+118)))
	if v95 != int32(112) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v261 = int32(_a_F_heap_insert_0)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_insert[0])) = v263 - int32(1)
	F_UnlockReleaseBuffer(m, v29)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L67
	}
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[2]))
	if v99 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v102 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v29 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v103 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v99 < int32(2) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[1]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v29^int32(-1))<<(uint(int32(2))%32))))
	v121 = v113
	goto L23
L25:
	;
	goto L26
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[3]))
	v121 = v115 + v29<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L27:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+8)))
	if v144 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L29
L29:
	;
	if base.B2i32(base.Ui32(v124) < base.Ui32(int32(_a_F_heap_insert_4))) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v129 == int32(0) {
		goto L27
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_log_heap_new_cid(m, l0, v21)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L36
	}
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
	switch v133 - int32(109) {
	case 0, 5:
		goto L34
	default:
		goto L27
	}
L34:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+104)))
	if v136 != int32(1) {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	goto L27
L37:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)))
	v158 = base.B2i32(base.Ui32(int32(24)) < base.Ui32(v149)) & base.B2i32((v149+int32(_a_F_heap_insert_2))&int32(_a_F_heap_insert_3) == int32(4))
	if v158 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v163 = int32(0)
	v164 = v6
	goto L39
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+8)) = uint16(v144)
	if v42 != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v159 = int32(6)
	goto L42
L41:
	;
	v159 = int32(0)
	goto L42
L42:
	;
	if v158 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v162 = int32(-128)
	goto L45
L44:
	;
	v162 = int32(0)
	goto L45
L45:
	;
	v163 = v159
	v164 = v162
	goto L39
L46:
	;
	v168 = v91 | int32(4)
	goto L48
L47:
	;
	v168 = v91
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)) = uint8(v168)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_heap_insert[2]))
	if v171 < int32(2) {
		v204 = v163
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L60
	}
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+118)))
	if v175 != int32(112) {
		v204 = v163
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+119)))
	if v178 == int32(102) {
		v204 = v163
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L53
L53:
	;
	if base.B2i32(base.Ui32(v181) < base.Ui32(int32(_a_F_heap_insert_4)))|l3&int32(8) != 0 {
		v204 = v163
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v188 = v168 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)) = uint8(v188)
	v191 = v163 | int32(16)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+68))
	if v193 != int32(99) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v198 == int32(0) {
		v204 = v191
		goto L49
	} else {
		goto L59
	}
L56:
	;
	v196 = F_isTempToastNamespace(m, v193)
	mBase = m.M
	v198 = v196
	goto L58
L57:
	;
	v198 = int32(1)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v202 = v168 | int32(24)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)) = uint8(v202)
	v204 = v191
	goto L49
L60:
	;
	F_XLogRegisterData(m, v15+int32(8), int32(3))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v214)
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)) = uint16(v216)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+6)) = uint8(v218)
	F_XLogRegisterBuffer(m, int32(0), v29, v204|int32(8))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_XLogRegisterBufData(m, int32(0), v15+int32(2), int32(5))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v233 = int32(23)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_XLogRegisterBufData(m, int32(0), v232+v233, v235-v233)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v241 = int32(_a_F_heap_insert_5)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[4])))
	v244 = v243 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_insert[4])) = uint8(v244)
	goto L65
L65:
	;
	v249 = F_XLogInsert(m, int32(10), v164&int32(255))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = base.I64_rotr(v249, int64(32))
	goto L16
L67:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v269 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_ReleaseBuffer(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	F_CacheInvalidateHeapTuple(m, l0, v21, int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	F_pgstat_count_heap_insert(m, l0, int64(1))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if l1 != v21 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v279)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v281
	F_pfree(m, v21)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	m.G0 = v15 + int32(16)
	return
L77:
	;
	goto L76
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
func F_heap_reloptions(m *base.Module, l0 int32, l1 int32) {
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
		v25 = F_build_reloptions(m, l1, v20, v20, int32(128), int32(_a_F_heap_reloptions_0), int32(24))
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
		v10 = F_build_reloptions(m, l1, int32(1), int32(2), int32(128), int32(_a_F_heap_reloptions_0), int32(24))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 == int32(0) {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = int64(-4616189618054758400)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(100)
				return
			}
		}
	}
}
func F_heap_tableam_handler(m *base.Module, l0 int32) int32 {
	return int32(_a_F_heap_tableam_handler_0)
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
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v331 int32
	_ = v331
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v664 int32
	_ = v664
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v981 int64
	_ = v981
	var v983 int64
	_ = v983
	var v985 int64
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
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
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1086 int32
	_ = v1086
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
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
	v291 = base.I32_div_s(v21+int32(7), int32(8))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	if v297&int32(4) != 0 {
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
	v94 = v55 + v76<<(uint(int32(4))%32) + v70*int32(100)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+105)))
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v78)+9)) = uint8(v95)
	v98 = v94 + int32(20)
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
	v270 = v70 + int32(1)
	if v270 != v56 {
		v70 = v270
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
	v101 = v70 << (uint(int32(2)) % 32)
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
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+v70<<(uint(int32(2))%32))))
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
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v260 = v259 + v78
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+8)))
	v263 = v261 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+8)) = uint8(v263)
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
		v242 = v165
		v243 = v196
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v256+v78)+4)) = v255
	goto L12
L37:
	;
	v245 = int32(1)
	if v243&v245 != 0 {
		v255 = int32(base.Ui32(v243) >> (uint(v245) % 32))
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
	v211 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v210+v70<<(uint(v211)%32)))) = v209
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v216 = v215 + v78
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+8)))
	v219 = v217 | v211
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+8)) = uint8(v219)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)))
	v223 = v221 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)) = uint8(v223)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v225 != int32(1) {
		v242 = v209
		v243 = v225
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
	v229 = int32(18)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v231 == v229 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v234 = v229
	goto L48
L47:
	;
	v234 = int32(2)
	goto L48
L48:
	;
	if base.Ui32((v231-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v241 = int32(6)
	goto L51
L50:
	;
	v241 = v234
	goto L51
L51:
	;
	v255 = v241
	goto L36
L52:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v255 = int32(base.Ui32(v249) >> (uint(int32(2)) % 32))
	goto L36
L53:
	;
	goto L11
L54:
	;
	v300 = (v291 + int32(30)) & int32(-8)
	goto L56
L55:
	;
	v300 = int32(24)
	goto L56
L56:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v301 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v304 = v302
	goto L59
L58:
	;
	v304 = int32(2032)
	goto L59
L59:
	;
	v306 = l3 & int32(-17)
	v311 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v494 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L106
	}
L61:
	;
	v313 = v304 - v300
	if base.Ui32(v311) <= base.Ui32(v313) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	goto L63
L63:
	;
	v331 = v18 + int32(4)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+52))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v345 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L60
L65:
	;
	if v428 < int32(0) {
		goto L60
	} else {
		goto L93
	}
L66:
	;
	v428 = int32(-1)
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
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v331)+24))
	v362 = int32(0)
	v364 = int32(24)
	v365 = int32(-1)
	goto L72
L72:
	;
	v373 = v355 + v362*int32(12)
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+8)))
	if int32(48)&v374 != 0 {
		v410 = v364
		v411 = v365
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v428 = v411
	goto L65
L74:
	;
	v414 = v362 + int32(1)
	if v414 != v345 {
		v362 = v414
		v364 = v410
		v365 = v411
		goto L72
	} else {
		goto L92
	}
L75:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	v377 = int32(2)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376+v362<<(uint(v377)%32))))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if v381&int32(3) == v377 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v389 = int32(1)
	goto L78
L77:
	;
	v389 = int32(0)
	goto L78
L78:
	;
	if base.B2i32(v381 == int32(1))|v389 != 0 {
		v410 = v364
		v411 = v365
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v345<<(uint(int32(4))%32)+v362*int32(100))+104)))
	goto L82
L80:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	v406 = base.B2i32(v364 < v405)
	if v364 < v405 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	goto L83
L83:
	;
	v398 = v394 - int32(101)
	if base.B2i32(v398 == int32(0))|base.B2i32(v398 == int32(19)) != 0 {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v410 = v364
	v411 = v365
	goto L74
L86:
	;
	v407 = v362
	goto L88
L87:
	;
	v407 = v365
	goto L88
L88:
	;
	if v364 < v405 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v408 = v405
	goto L91
L90:
	;
	v408 = v364
	goto L91
L91:
	;
	v410 = v408
	v411 = v407
	goto L74
L92:
	;
	goto L73
L93:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v431<<(uint(int32(4))%32)+v428*int32(100))+104)))
	if v438 == int32(120) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(32)+v428*int32(12))+4))
	if base.Ui32(v458) <= base.Ui32(v313) {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	F_toast_tuple_try_compression(m, v331, v428)
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
	v443 = int32(32)
	v447 = v18 + v443 + v428*int32(12)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+8)))
	v450 = v448 | v443
	*(*uint8)(unsafe.Add(mBase, uint32(v447)+8)) = uint8(v450)
	goto L94
L98:
	;
	goto L94
L99:
	;
	v472 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+112))
	if v461 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_toast_tuple_externalize(m, v18+int32(4), v428, v306)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	if base.Ui32(v313) < base.Ui32(v472) {
		goto L63
	} else {
		goto L104
	}
L104:
	;
	goto L64
L105:
	;
	v645 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L144
	}
L106:
	;
	if base.Ui32(v494) <= base.Ui32(v313) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	goto L108
L108:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+112))
	if v513 == int32(0) {
		goto L105
	} else {
		goto L110
	}
L109:
	;
	goto L105
L110:
	;
	v517 = v18 + int32(4)
	v518 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+52))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	if v531 <= v518 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v614 < int32(0) {
		goto L105
	} else {
		goto L139
	}
L112:
	;
	v614 = int32(-1)
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
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v517)+24))
	v548 = int32(0)
	v550 = int32(24)
	v551 = int32(-1)
	goto L118
L118:
	;
	v559 = v541 + v548*int32(12)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+8)))
	if int32(16)&v560 != 0 {
		v596 = v550
		v597 = v551
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v614 = v597
	goto L111
L120:
	;
	v600 = v548 + int32(1)
	if v600 != v531 {
		v548 = v600
		v550 = v596
		v551 = v597
		goto L118
	} else {
		goto L138
	}
L121:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	v563 = int32(2)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562+v548<<(uint(v563)%32))))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	if v567&int32(3) == v563 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v575 = v518
	goto L124
L123:
	;
	v575 = int32(0)
	goto L124
L124:
	;
	if base.B2i32(v567 == int32(1))|v575 != 0 {
		v596 = v550
		v597 = v551
		goto L120
	} else {
		goto L125
	}
L125:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530+v531<<(uint(int32(4))%32)+v548*int32(100))+104)))
	goto L128
L126:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	v592 = base.B2i32(v550 < v591)
	if v550 < v591 {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	goto L129
L129:
	;
	v584 = v580 - int32(101)
	if base.B2i32(v584 == int32(0))|base.B2i32(v584 == int32(19)) != 0 {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v596 = v550
	v597 = v551
	goto L120
L132:
	;
	v593 = v548
	goto L134
L133:
	;
	v593 = v551
	goto L134
L134:
	;
	if v550 < v591 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v594 = v591
	goto L137
L136:
	;
	v594 = v550
	goto L137
L137:
	;
	v596 = v594
	v597 = v593
	goto L120
L138:
	;
	goto L119
L139:
	;
	F_toast_tuple_externalize(m, v517, v614, v306)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v623 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if base.Ui32(v313) < base.Ui32(v623) {
		goto L108
	} else {
		goto L142
	}
L142:
	;
	goto L109
L143:
	;
	v792 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L181
	}
L144:
	;
	if base.Ui32(v645) <= base.Ui32(v313) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	goto L146
L146:
	;
	v664 = v18 + int32(4)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)+52))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	if v678 <= int32(0) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L143
L148:
	;
	if v761 < int32(0) {
		goto L143
	} else {
		goto L176
	}
L149:
	;
	v761 = int32(-1)
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
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v664)+24))
	v695 = int32(0)
	v697 = int32(24)
	v698 = int32(-1)
	goto L155
L155:
	;
	v706 = v688 + v695*int32(12)
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+8)))
	if int32(48)&v707 != 0 {
		v743 = v697
		v744 = v698
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v761 = v744
	goto L148
L157:
	;
	v747 = v695 + int32(1)
	if v747 != v678 {
		v695 = v747
		v697 = v743
		v698 = v744
		goto L155
	} else {
		goto L175
	}
L158:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v664)+4))
	v710 = int32(2)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709+v695<<(uint(v710)%32))))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v714&int32(3) == v710 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v722 = int32(1)
	goto L161
L160:
	;
	v722 = int32(0)
	goto L161
L161:
	;
	if base.B2i32(v714 == int32(1))|v722 != 0 {
		v743 = v697
		v744 = v698
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v678<<(uint(int32(4))%32)+v695*int32(100))+104)))
	goto L164
L163:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v739 = base.B2i32(v697 < v738)
	if v697 < v738 {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	if v727 != int32(109) {
		v743 = v697
		v744 = v698
		goto L157
	} else {
		goto L167
	}
L167:
	;
	goto L163
L169:
	;
	v740 = v695
	goto L171
L170:
	;
	v740 = v698
	goto L171
L171:
	;
	if v697 < v738 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v741 = v738
	goto L174
L173:
	;
	v741 = v697
	goto L174
L174:
	;
	v743 = v741
	v744 = v740
	goto L157
L175:
	;
	goto L156
L176:
	;
	F_toast_tuple_try_compression(m, v664, v761)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v770 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	if base.Ui32(v313) < base.Ui32(v770) {
		goto L146
	} else {
		goto L179
	}
L179:
	;
	goto L147
L180:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	if v941&int32(8) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L181:
	;
	v795 = int32(_a_F_heap_toast_insert_or_update_6) - v300
	if base.Ui32(v792) <= base.Ui32(v795) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	goto L183
L183:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+112))
	if v813 == int32(0) {
		goto L180
	} else {
		goto L185
	}
L184:
	;
	goto L180
L185:
	;
	v817 = v18 + int32(4)
	v818 = int32(0)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)+52))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v830)))
	if v831 <= v818 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	if v914 < int32(0) {
		goto L180
	} else {
		goto L214
	}
L187:
	;
	v914 = int32(-1)
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
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v817)+24))
	v848 = int32(0)
	v850 = int32(24)
	v851 = int32(-1)
	goto L193
L193:
	;
	v859 = v841 + v848*int32(12)
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+8)))
	if int32(16)&v860 != 0 {
		v896 = v850
		v897 = v851
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v914 = v897
	goto L186
L195:
	;
	v900 = v848 + int32(1)
	if v900 != v831 {
		v848 = v900
		v850 = v896
		v851 = v897
		goto L193
	} else {
		goto L213
	}
L196:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	v863 = int32(2)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v862+v848<<(uint(v863)%32))))
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	if v867&int32(3) == v863 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v875 = v818
	goto L199
L198:
	;
	v875 = int32(0)
	goto L199
L199:
	;
	if base.B2i32(v867 == int32(1))|v875 != 0 {
		v896 = v850
		v897 = v851
		goto L195
	} else {
		goto L200
	}
L200:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830+v831<<(uint(int32(4))%32)+v848*int32(100))+104)))
	goto L202
L201:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v859)+4))
	v892 = base.B2i32(v850 < v891)
	if v850 < v891 {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	if v880 != int32(109) {
		v896 = v850
		v897 = v851
		goto L195
	} else {
		goto L205
	}
L205:
	;
	goto L201
L207:
	;
	v893 = v848
	goto L209
L208:
	;
	v893 = v851
	goto L209
L209:
	;
	if v850 < v891 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v894 = v891
	goto L212
L211:
	;
	v894 = v850
	goto L212
L212:
	;
	v896 = v894
	v897 = v893
	goto L195
L213:
	;
	goto L194
L214:
	;
	F_toast_tuple_externalize(m, v817, v914, v306)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v923 = F_heap_compute_data_size(m, v20, v18+int32(_a_F_heap_toast_insert_or_update_1), v18+int32(_a_F_heap_toast_insert_or_update_2))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	if base.Ui32(v795) < base.Ui32(v923) {
		goto L183
	} else {
		goto L217
	}
L217:
	;
	goto L184
L218:
	;
	v1014 = v18 + int32(4)
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+52))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+20)))
	v1021 = int32(0)
	if base.B2i32(v1018&int32(2) == v1021)|base.B2i32(v1017 <= v1021) != 0 {
		goto L228
	} else {
		goto L229
	}
L219:
	;
	v1008 = l1
	goto L218
L220:
	;
	goto L221
L221:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v951 = base.I32_div_s(v21+int32(7), int32(8))
	if v941&int32(4) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v959 = (v951 + int32(30)) & int32(-8)
	goto L224
L223:
	;
	v959 = int32(24)
	goto L224
L224:
	;
	v961 = v18 + int32(_a_F_heap_toast_insert_or_update_1)
	v963 = v18 + int32(_a_F_heap_toast_insert_or_update_2)
	v964 = F_heap_compute_data_size(m, v20, v961, v963)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v966 = v964 + v959
	v969 = F_palloc0(m, v966+int32(24))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v969))) = v966
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v969)+4)) = v972
	v974 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v969)+8)) = uint16(v974)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v978 = v969 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v969)+16)) = v978
	*(*int32)(unsafe.Add(mBase, uint32(v969)+12)) = v976
	v981 = *(*int64)(unsafe.Add(mBase, uint32(v946)+15))
	*(*int64)(unsafe.Add(mBase, uint32(v969)+39)) = v981
	v983 = *(*int64)(unsafe.Add(mBase, uint32(v946)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v969)+32)) = v983
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v946)))
	*(*int64)(unsafe.Add(mBase, uint32(v969)+24)) = v985
	*(*uint8)(unsafe.Add(mBase, uint32(v969)+46)) = uint8(v959)
	v988 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v969)+42)))
	v991 = v988&int32(_a_F_heap_toast_insert_or_update_7) | v21
	*(*uint16)(unsafe.Add(mBase, uint32(v969)+42)) = uint16(v991)
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	F_heap_fill_tuple(m, v20, v961, v963, v959+v978, v969+int32(44), (v969+int32(47))&(v998<<(uint(int32(29))%32)>>(uint(int32(31))%32)))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v1008 = v969
	goto L218
L228:
	;
	v1075 = v1018
	goto L230
L229:
	;
	v1027 = int32(0)
	goto L231
L230:
	;
	v1078 = int32(0)
	if base.B2i32(v1075&int32(1) == v1078)|base.B2i32(v1017 <= v1078) == v1078 {
		goto L238
	} else {
		goto L239
	}
L231:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+24))
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1027*int32(12))+8)))
	if v1046&int32(2) != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+20)))
	v1075 = v1059
	goto L230
L233:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1049+v1027<<(uint(int32(2))%32))))
	F_pfree(m, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1057 = v1027 + int32(1)
	if v1057 != v1017 {
		v1027 = v1057
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
	v1086 = int32(0)
	goto L241
L239:
	;
	goto L240
L240:
	;
	m.G0 = v18 + int32(_a_F_heap_toast_insert_or_update_0)
	return v1008
L241:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+24))
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101+v1086*int32(12))+8)))
	if v1105&int32(1) != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	goto L240
L243:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+12))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1109+v1086<<(uint(int32(2))%32))))
	F_toast_delete_datum(m, v1113, int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1118 = v1086 + int32(1)
	if v1118 != v1017 {
		v1086 = v1118
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
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
		v25 = F_SearchSysCache1(m, int32(57), l0)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v25 != 0 {
				v31 = F_SysCacheGetAttr(m, int32(57), v25, int32(33), v17+int32(47))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
					if l3 == int32(116) {
						v37 = F_LookupCreationNamespace(m, int32(_a_F_make_new_heap_0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v41 = v37
							if v33 != 0 {
								v43 = int32(0)
							} else {
								v43 = v31
							}
							*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
							v51 = F_pg_snprintf(m, v17+int32(48), int32(64), int32(_a_F_make_new_heap_1), v17+int32(32))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+119)))
								switch v56 - int32(83) {
								case 0, 22, 26, 31, 33:
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+88))
									v62 = base.B2i32(v59 == int32(0))
								default:
									v62 = int32(0)
								}
								v65 = int32(0)
								v73 = int32(1)
								v76 = F_heap_create_with_catalog(m, v17+int32(48), v41, l1, v65, v65, v65, v54, l2, v23, v65, int32(114), l3, v65, v62, v65, v43, v65, v73, v73, l0, v65)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v25)
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
												v85 = F_SearchSysCache1(m, int32(57), v83)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													if v85 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v83
															F_errmsg_internal(m, int32(_a_F_make_new_heap_2), v17+int32(16))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_make_new_heap_3), int32(806), int32(_a_F_make_new_heap_4))
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = F_SysCacheGetAttr(m, int32(57), v85, int32(33), v17+int32(47))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
															if v96 != 0 {
																v97 = int32(0)
															} else {
																v97 = v94
															}
															v98 = F_table_open(m, v76, l4)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																v100 = int32(0)
																v103 = F_create_toast_table(m, v98, v100, v100, v97, l4, v100, v83)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v98, int32(0))
																	mBase = m.M
																	v107 = m.ExcPending
																	if v107 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v85)
																		mBase = m.M
																		v109 = m.ExcPending
																		if v109 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v19, int32(0))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
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
												v115 = m.ExcPending
												if v115 != 0 {
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
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
						v41 = v40
						if v33 != 0 {
							v43 = int32(0)
						} else {
							v43 = v31
						}
						*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
						v51 = F_pg_snprintf(m, v17+int32(48), int32(64), int32(_a_F_make_new_heap_1), v17+int32(32))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+119)))
							switch v56 - int32(83) {
							case 0, 22, 26, 31, 33:
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+88))
								v62 = base.B2i32(v59 == int32(0))
							default:
								v62 = int32(0)
							}
							v65 = int32(0)
							v73 = int32(1)
							v76 = F_heap_create_with_catalog(m, v17+int32(48), v41, l1, v65, v65, v65, v54, l2, v23, v65, int32(114), l3, v65, v62, v65, v43, v65, v73, v73, l0, v65)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v25)
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
											v85 = F_SearchSysCache1(m, int32(57), v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												if v85 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v83
														F_errmsg_internal(m, int32(_a_F_make_new_heap_2), v17+int32(16))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_make_new_heap_3), int32(806), int32(_a_F_make_new_heap_4))
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = F_SysCacheGetAttr(m, int32(57), v85, int32(33), v17+int32(47))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
														if v96 != 0 {
															v97 = int32(0)
														} else {
															v97 = v94
														}
														v98 = F_table_open(m, v76, l4)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															v100 = int32(0)
															v103 = F_create_toast_table(m, v98, v100, v100, v97, l4, v100, v83)
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v98, int32(0))
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v85)
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v19, int32(0))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
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
											v115 = m.ExcPending
											if v115 != 0 {
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
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
					F_errmsg_internal(m, int32(_a_F_make_new_heap_2), v17)
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_make_new_heap_3), int32(733), int32(_a_F_make_new_heap_4))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
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
