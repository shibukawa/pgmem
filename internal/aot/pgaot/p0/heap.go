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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v13 = F_HeapTupleSatisfiesVacuumHorizon(m, l0, l2, v7+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(2) {
			v19 = int32(0)
			v20 = int32(2)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if base.B2i32(base.Ui32(v20) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v21)) == v19 {
				v33 = base.B2i32(base.Ui32(v21) < base.Ui32(l1))
			} else {
				v33 = int32(base.Ui32(v21-l1) >> (uint(int32(31)) % 32))
			}
			if v33 != 0 {
				v34 = v19
			} else {
				v34 = v20
			}
			v35 = v34
		} else {
			v35 = v13
		}
		m.G0 = v7 + int32(16)
		return v35
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
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v336 int32
	_ = v336
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
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
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L158
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L154
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L150
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L147
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L142
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
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L138
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
		v104 = l3
		v105 = v56
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_LockRelationOid(m, v104, int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
	}
L24:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
	if v58 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v102 = F_GetNewRelFileNumber(m, l2, v33, l11)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L36
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
	v64 = *(*int32)(unsafe.Add(mBase, _consts[201]))
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
	v78 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	if v78 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[201])) = v68
	v71 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	if v71 == v68 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[203])) = int32(0)
	v104 = v64
	v105 = v71
	goto L23
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[202])) = int32(0)
	v85 = l10 - int32(83)
	if base.Ui32(int32(31)) < base.Ui32(v85) {
		v104 = v78
		v105 = v56
		goto L23
	} else {
		goto L33
	}
L33:
	;
	if int32(1)<<(uint(v85)%32)&int32(-2076180479) == int32(0) {
		v104 = v78
		v105 = v56
		goto L23
	} else {
		goto L34
	}
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if v95 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[204])) = int32(0)
	v104 = v78
	v105 = v95
	goto L23
L36:
	;
	v104 = v102
	v105 = v56
	goto L23
L37:
	;
	if l16 == int32(0) {
		v131 = v22
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v227)+104)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v227)+96)) = int64(-4647714815446351872)
	if l10 == int32(83) {
		goto L57
	} else {
		goto L58
	}
L39:
	;
	v137 = F_heap_create(m, l0, l1, l2, v104, v105, l7, l8, l10, l11, l12, l13, l17, v29+int32(52), v29+int32(48), int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L46
	}
L40:
	;
	switch l10 - int32(83) {
	case 0:
		goto L42
	default:
		v131 = v22
		goto L39
	case 19, 26, 29, 31, 35:
		goto L41
	}
L41:
	;
	v129 = F_get_user_default_acl(m, int32(41), l6, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v115 = F_get_user_default_acl(m, int32(37), l6, l1)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v123 = F_heap_create(m, l0, l1, l2, v104, v105, l7, l8, int32(83), l11, l12, l13, l17, v29+int32(52), v29+int32(48), int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+132)) = l19
	v220 = int32(0)
	v222 = v123
	v224 = v115
	goto L38
L45:
	;
	v131 = v129
	goto L39
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+132)) = l19
	v141 = int32(0)
	switch l10 - int32(73) {
	case 0, 10:
		v220 = v141
		v222 = v137
		v224 = v131
		goto L38
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L47
	default:
		goto L48
	}
L47:
	;
	v151 = int32(0)
	v163 = F_AssignTypeArrayOid(m)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	switch l10 - int32(105) {
	case 0, 11:
		v220 = v141
		v222 = v137
		v224 = v131
		goto L38
	default:
		goto L47
	}
L49:
	;
	v165 = int32(0)
	F_TypeCreate(m, v29+int32(68), l4, l0, l1, v104, l10, l6, int32(-1), int32(99), int32(67), v151, int32(44), int32(2290), int32(2291), int32(2402), int32(2403), v151, v151, v151, v151, v151, v151, v163, v165, v165, v165, v165, int32(100), int32(120), int32(-1), v165, v165, v165)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if l20 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l20)+8)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(l20)+4)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(l20))) = v178
	goto L53
L52:
	;
	goto L53
L53:
	;
	v186 = F_makeArrayTypeName(m, l0, l1)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v188 = int32(0)
	v190 = int32(-1)
	F_TypeCreate(m, v29+int32(68), v163, v186, l1, v188, v188, l6, v190, int32(98), int32(65), v188, int32(44), int32(750), int32(751), int32(2400), int32(2401), v188, v188, int32(3816), int32(6179), v177, int32(1), v188, v188, v188, v188, v188, int32(100), int32(120), v190, v188, v188, v188)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_pfree(m, v186)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v220 = v177
	v222 = v137
	v224 = v131
	goto L38
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v227)+96)) = int64(4575657221408423937)
	goto L59
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+140)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v227)+136)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v227)+80)) = l6
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+131)) = uint8(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v227)+76)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v227)+72)) = v220
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v222)+52))
	if v220 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v245 = v220
	goto L62
L61:
	;
	v245 = int32(2249)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v222)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+8)) = int32(-1)
	F_InsertPgClassTuple(m, v33, v222, v104, v224, l15)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v222)+52))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v256 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v259 = F_CatalogOpenIndexes(m, v256)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_InsertPgAttributeTuples(m, v256, v252, v104, int32(0), v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if int32(0) < v253 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v268 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	switch l10 - int32(99) {
	case 0, 19:
		goto L78
	default:
		goto L79
	}
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v296 = v268 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(1247)
	v308 = v252 + int32(20) + v294<<(uint(int32(4))%32) + v268*int32(100)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v309
	F_recordDependencyOn(m, v29+int32(68), v29+int32(56), int32(110))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v308)+96))
	if v320 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v296 != v253 {
		v268 = v296
		goto L70
	} else {
		goto L77
	}
L74:
	;
	if v320 == int32(100) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(3456)
	F_recordDependencyOn(m, v29+int32(68), v29+int32(56), int32(110))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	goto L71
L78:
	;
	F_CatalogCloseIndexes(m, v259)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L83
	}
L79:
	;
	v368 = F_CreateTupleDesc(m, int32(6), int32(765440))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_InsertPgAttributeTuples(m, v256, v368, v104, int32(0), v259)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_FreeTupleDesc(m, v368)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	F_sequence_close(m, v256, int32(3))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v382 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v447 != 0 {
		goto L105
	} else {
		goto L106
	}
L86:
	;
	switch l10 - int32(99) {
	case 0, 17:
		goto L85
	default:
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v104
	v390 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v390
	F_recordDependencyOnOwner(m, v390, v104, l6)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_recordDependencyOnNewAcl(m, int32(1259), v104, l6, v224)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_recordDependencyOnCurrentExtension(m, v29+int32(68), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v403 = F_new_object_addresses(m)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(2615)
	F_add_exact_object_address(m, v29+int32(56), v403)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
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
	F_add_exact_object_address(m, v29+int32(56), v403)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
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
	F_record_object_address_dependencies(m, v29+int32(68), v403, int32(110))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L103
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(2601)
	F_add_exact_object_address(m, v29+int32(56), v403)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L102
	}
L99:
	;
	if l7 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	if l10 != int32(112) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	goto L97
L103:
	;
	F_free_object_addresses(m, v403)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L85
L105:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v104, int32(0), l18)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if l9 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l14) {
		goto L127
	} else {
		goto L128
	}
L110:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v456 <= int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v459 = int32(0)
	v463 = v459
	v478 = v459
	goto L113
L113:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l9)+12))
	v488 = int32(2)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+v463<<(uint(v488)%32))))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	switch v492 - v488 {
	case 0:
		goto L116
	default:
		goto L117
	case 3:
		goto L118
	}
L114:
	;
	if v530 <= int32(0) {
		goto L109
	} else {
		goto L125
	}
L115:
	;
	v532 = v463 + int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v532 < v533 {
		v463 = v532
		v478 = v530
		goto L113
	} else {
		goto L124
	}
L116:
	;
	v525 = int32(*(*int16)(unsafe.Add(mBase, uint32(v491)+12)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v491)+16))
	v527 = F_StoreAttrDefault(m, v222, v525, v526, l18)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L123
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v491)+8))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v491)+16))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+20)))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+21)))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+22)))
	v504 = int32(*(*int16)(unsafe.Add(mBase, uint32(v491)+24)))
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+26)))
	v506 = F_StoreRelCheck(m, v222, v495, v496, v497, (v498^int32(-1))&int32(1), v503, v504, v505, l18)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v506
	v530 = v478 + int32(1)
	goto L115
L120:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v515
	F_errmsg_internal(m, int32(484600), v29)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(496174), int32(2346), int32(120363))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v527
	v530 = v478
	goto L115
L124:
	;
	goto L114
L125:
	;
	F_SetRelationNumChecks(m, v222, v530)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L109
L127:
	;
	v568 = l14
	goto L129
L128:
	;
	v568 = int32(0)
	goto L129
L129:
	;
	if v568 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v569 = int32(4515488)
	v570 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v573 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v573
	v576 = F_palloc(m, int32(16))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_sequence_close(m, v222, int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L136
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v576)+4)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v104
	v581 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+8))
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v576)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v576)+8)) = v582
	v587 = *(*int32)(unsafe.Add(mBase, _consts[208]))
	v588 = F_lcons(m, v576, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v570
	*(*int32)(unsafe.Add(mBase, _consts[208])) = v588
	goto L132
L136:
	;
	F_sequence_close(m, v33, int32(3))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	m.G0 = v29 + int32(80)
	return v104
L138:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = l0
	F_errmsg(m, int32(116708), v29+int32(32))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(496174), int32(1179), int32(327365))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l0
	F_errmsg(m, int32(116985), v29+int32(16))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errhint(m, int32(633712), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(496174), int32(1198), int32(327365))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(419122), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(496174), int32(1205), int32(327365))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(412743), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(496174), int32(1236), int32(327365))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(413058), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(496174), int32(1247), int32(327365))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(412810), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(496174), int32(1257), int32(327365))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
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
		F_appendStringInfo(m, l0, int32(746079), v8+int32(32))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v35, int32(123833))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v39
				F_appendStringInfo(m, l0, int32(510564), v8+int32(16))
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
		F_appendStringInfo(m, l0, int32(746051), v8-int32(-64))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v55, int32(123829))
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
				F_appendStringInfo(m, l0, int32(58477), v8+int32(48))
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
		F_appendStringInfoString(m, l0, int32(507952))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return
		} else {
			if v94&int32(1) != 0 {
				F_appendStringInfoString(m, l0, int32(746163))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					if v94&int32(2) != 0 {
						F_appendStringInfoString(m, l0, int32(746111))
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
								F_appendStringInfo(m, l0, int32(57922), v8+int32(112))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l0, int32(547338))
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
							F_appendStringInfo(m, l0, int32(57922), v8+int32(112))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l0, int32(547338))
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
					F_appendStringInfoString(m, l0, int32(746111))
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
							F_appendStringInfo(m, l0, int32(57922), v8+int32(112))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l0, int32(547338))
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
						F_appendStringInfo(m, l0, int32(57922), v8+int32(112))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l0, int32(547338))
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
		F_appendStringInfo(m, l0, int32(746051), v8+int32(96))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v79, int32(123829))
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
				F_appendStringInfo(m, l0, int32(58477), v8+int32(80))
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
		F_appendStringInfo(m, l0, int32(58512), v8+int32(128))
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
		F_appendStringInfo(m, l0, int32(746079), v8+int32(160))
		mBase = m.M
		v159 = m.ExcPending
		if v159 != 0 {
			return
		} else {
			v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
			F_infobits_desc(m, l0, v160, int32(123833))
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return
			} else {
				v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v164
				F_appendStringInfo(m, l0, int32(510564), v8+int32(144))
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
		F_appendStringInfo(m, l0, int32(58512), v8+int32(176))
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
		F_appendStringInfo(m, l0, int32(510530), v8)
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v342 int32
	_ = v342
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
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
	v85 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v85 != 0 {
		goto L19
	} else {
		goto L20
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
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L112
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L108
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L104
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L100
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L97
	}
L17:
	;
	if v274 != v49+int32(1) {
		goto L14
	} else {
		goto L87
	}
L18:
	;
	if v100 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v100 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	v90 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	if v90 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v100 = base.B2i32(v88 != int32(0))
	goto L18
L23:
	;
	if v88 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v95 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v100 = int32(0)
	goto L18
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v119 = F_systable_beginscan_ordered(m, l0, v82, int32(4173960), v77, v21+int32(112))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	F_errmsg_internal(m, int32(86778), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(494178), int32(653), int32(86394))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v122 = F_systable_getnext_ordered(m, v119, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v122 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v274 = v40
	goto L17
L35:
	;
	goto L36
L36:
	;
	v129 = F_fastgetattr_1(m, v122, int32(2), v23, v21+int32(107))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v134 = F_fastgetattr_1(m, v122, int32(3), v23, v21+int32(107))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	if v40 != v129 {
		v503 = v40
		v504 = v129
		goto L12
	} else {
		goto L44
	}
L39:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v136&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v139 = int32(1)
	if v136&v139 == int32(0) {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v148 = int32(4)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v154 = v148
	v155 = int32(base.Ui32(v149)>>(uint(int32(2))%32)) - v148
	goto L38
L43:
	;
	v144 = int32(1)
	v154 = v139
	v155 = int32(base.Ui32(v136)>>(uint(v144)%32)) - v144
	goto L38
L44:
	;
	if base.Ui32(v49) < base.Ui32(v40) {
		v390 = v40
		goto L15
	} else {
		goto L45
	}
L45:
	;
	v161 = v44*int32(-1996) + l2
	if base.Ui32(v40) < base.Ui32(v44) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v163 = int32(1996)
	goto L48
L47:
	;
	v163 = v161
	goto L48
L48:
	;
	if v155 != v163 {
		v451 = v155
		v455 = v40
		v456 = v163
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v165 = int32(1996)
	v166 = v40 * v165
	v167 = l3 - v166
	v170 = l5 - l3 + int32(4)
	v179 = v47 - v49*v165
	if v40 == v49 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v183 = v179
	goto L52
L51:
	;
	v183 = v155 - int32(1)
	goto L52
L52:
	;
	v186 = v183 - v167 + int32(1)
	if v186 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v189 = int32(1)
	v190 = v40 + v189
	v192 = F_systable_getnext_ordered(m, v119, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	v187 = F__emscripten_memcpy_bulkmem(m, v167+(v170+v166), v134+v154+v167, v186)
	mBase = m.M
	goto L56
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	if v192 == int32(0) {
		v274 = v190
		goto L17
	} else {
		goto L58
	}
L58:
	;
	v202 = v192
	v205 = v190
	goto L59
L59:
	;
	v219 = F_fastgetattr_1(m, v202, int32(2), v23, v21+int32(107))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v274 = v263
	goto L17
L61:
	;
	v224 = F_fastgetattr_1(m, v202, int32(3), v23, v21+int32(107))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	if v205 != v219 {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v226&int32(3) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v234 = int32(4)
	v246 = int32(base.Ui32(v231)>>(uint(int32(2))%32)) - v234
	v247 = v234
	goto L62
L65:
	;
	goto L66
L66:
	;
	if v226&int32(1) == int32(0) {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v241 = int32(1)
	v246 = int32(base.Ui32(v226)>>(uint(v241)%32)) - v241
	v247 = v241
	goto L62
L68:
	;
	v503 = v205
	v504 = v219
	goto L12
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v49) < base.Ui32(v205) {
		v390 = v205
		goto L15
	} else {
		goto L71
	}
L71:
	;
	if base.Ui32(v205) < base.Ui32(v44) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v252 = int32(1996)
	goto L74
L73:
	;
	v252 = v161
	goto L74
L74:
	;
	if v252 != v246 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v451 = v246
	v455 = v205
	v456 = v252
	goto L13
L76:
	;
	goto L77
L77:
	;
	if v205 == v49 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v259 = v179 + int32(1)
	goto L80
L79:
	;
	v259 = v246
	goto L80
L80:
	;
	if v259 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v262 = int32(1)
	v263 = v205 + v262
	v265 = F_systable_getnext_ordered(m, v119, v262)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L85
	}
L82:
	;
	v260 = F__emscripten_memcpy_bulkmem(m, v170+v205*int32(1996), v247+v224, v259)
	mBase = m.M
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	if v265 != 0 {
		v202 = v265
		v205 = v263
		goto L59
	} else {
		goto L86
	}
L86:
	;
	goto L60
L87:
	;
	F_systable_endscan_ordered(m, v119)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	v291 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	if v291 < v292 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v295 = v291
	goto L92
L90:
	;
	goto L91
L91:
	;
	F_pfree(m, v290)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L96
	}
L92:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v290+v295<<(uint(int32(2))%32))))
	F_relation_close(m, v316, int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	v321 = v295 + int32(1)
	if v321 != v292 {
		v295 = v321
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	m.G0 = v21 + int32(272)
	return
L97:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v368 + int32(4)
	F_errmsg_internal(m, int32(184797), v21+int32(96))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(492917), int32(729), int32(418270))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v408 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v390
	F_errmsg_internal(m, int32(184979), v21+int32(16))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(492917), int32(749), int32(418270))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v433 + int32(4)
	F_errmsg_internal(m, int32(184848), v21)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(492917), int32(786), int32(418270))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v472 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v44 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v451
	F_errmsg_internal(m, int32(184897), v21+int32(48))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(492917), int32(758), int32(418270))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v520 + int32(4)
	F_errmsg_internal(m, int32(185053), v21+int32(80))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(492917), int32(742), int32(418270))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 <= int32(1664) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v15 <= int32(0) {
		v56 = v4
		v58 = int32(24)
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L12
	} else {
		goto L37
	}
L4:
	;
	v59 = F_heap_compute_data_size(m, l0, l1, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v24 = v4
	goto L6
L6:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v24))))
	if v32 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = v4
	v58 = int32(24)
	goto L4
L8:
	;
	v39 = base.I32_div_s(v15+int32(7), int32(8))
	v56 = int32(1)
	v58 = (v39 + int32(30)) & int32(-8)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v45 = v24 + int32(1)
	if v45 != v15 {
		v24 = v45
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	v63 = v59 + v58
	v66 = F_palloc0(m, v63+int32(24))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v68
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+8)) = uint16(v68)
	v72 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v63
	v76 = v66 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v63 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v76
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+40)) = uint16(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v83
	v90 = v66 + int32(42)
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90))))
	v94 = v91&int32(63488) | v15
	*(*uint16)(unsafe.Add(mBase, uint32(v90))) = uint16(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+46)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v58 + v76
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v56 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v103 = v66 + int32(47)
	goto L17
L16:
	;
	v103 = v68
	goto L17
L17:
	;
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = base.B2i32(v103 != v104) << (uint(int32(7)) % 32)
	if v103 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v112 = v103 - int32(1)
	goto L20
L19:
	;
	v112 = v104
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v112
	v115 = v66 + int32(44)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v118 = v116 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(v115))) = uint16(v118)
	if int32(0) < v99 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v128 = int32(0)
	goto L24
L22:
	;
	goto L23
L23:
	;
	m.G0 = v13 + int32(32)
	return v66
L24:
	;
	v135 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v142 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v143 = v13 + int32(24)
	goto L28
L27:
	;
	v143 = v135
	goto L28
L28:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1+v128<<(uint(int32(2))%32))))
	v152 = v151
	goto L31
L30:
	;
	v152 = v135
	goto L31
L31:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v128))))
	v156 = v154
	goto L34
L33:
	;
	v156 = int32(1)
	goto L34
L34:
	;
	F_fill_val(m, l0+int32(20)+v128<<(uint(int32(4))%32), v143, v13+int32(20), v13+int32(28), v115, v152, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v160 = v128 + int32(1)
	if v160 != v99 {
		v128 = v160
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(1664)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15
	F_errmsg(m, int32(679004), v13)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(499074), int32(1134), int32(383757))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if int32(0) < l1 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+18)))
		if base.Ui32(v15&int32(2047)) < base.Ui32(l1) {
			v19 = F_getmissingattr(m, l2, l1, l3)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v93 = v19
				m.G0 = v10 + int32(16)
				return v93
			}
		} else {
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v23)
			v25 = int32(1)
			v26 = l1 - v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
			if v28&v25 == v23 {
				v37 = l2 + v26<<(uint(int32(4))%32) + int32(20)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				if int32(0) <= v38 {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v43 = v27 + v41 + v38
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+6)))
					if v44 != int32(1) {
						v93 = v43
						m.G0 = v10 + int32(16)
						return v93
					} else {
						v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+4)))
						switch v47&int32(65535) - int32(1) {
						case 0:
							v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v43))))
							v93 = v52
							m.G0 = v10 + int32(16)
							return v93
						case 1:
							v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
							v93 = v53
							m.G0 = v10 + int32(16)
							return v93
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v47
								F_errmsg_internal(m, int32(483438), v10)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(326693), int32(70), int32(67821))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v93 = v54
							m.G0 = v10 + int32(16)
							return v93
						}
					}
				} else {
					v68 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v93 = v68
						m.G0 = v10 + int32(16)
						return v93
					}
				}
			} else {
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(base.Ui32(v26)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v73)>>(uint(v26&int32(7))%32))&int32(1) == int32(0) {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v81)
					v93 = int32(0)
					m.G0 = v10 + int32(16)
					return v93
				} else {
					v84 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v93 = v84
						m.G0 = v10 + int32(16)
						return v93
					}
				}
			}
		}
	} else {
		v86 = F_heap_getsysattr(m, l0, l1, l3)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return int32(0)
		} else {
			v93 = v86
			m.G0 = v10 + int32(16)
			return v93
		}
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
	if v6 == int32(758388) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[25]))
		if v10 != 0 {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[26])))
			if v12&int32(1) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(336177), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497384), int32(1362), int32(63453))
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
				F_errmsg_internal(m, int32(444125), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497384), int32(1352), int32(63453))
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(l2^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L1
L3:
	;
	goto L4
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	return v309
L9:
	;
	v309 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v62 = v7
	v65 = v8
	v66 = v44
	v69 = v8
	v70 = v7 ^ int32(1)
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
	v309 = v302
	goto L8
L14:
	;
	v309 = int32(0)
	goto L8
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(int32(base.Ui32(v76+int32(262120))>>(uint(int32(2))%32))&int32(65535)) < base.Ui32(v66) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v309 = int32(0)
	goto L8
L18:
	;
	goto L19
L19:
	;
	v88 = int32(0)
	v93 = v66<<(uint(int32(2))%32) + (v42 + int32(24)) - int32(4)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	switch int32(base.Ui32(v94)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L21
	case 1:
		goto L22
	default:
		v309 = v88
		goto L8
	}
L20:
	;
	v302 = int32(0)
	if v292 != 0 {
		v62 = v302
		v65 = v291
		v66 = v292
		v69 = v295
		v70 = v296
		goto L12
	} else {
		goto L81
	}
L21:
	;
	v109 = v42 + v94&int32(32767)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(base.Ui32(v111) >> (uint(int32(17)) % 32))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+8)) = uint16(v66)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)) = uint16(v48)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)) = uint16(v49)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v115
	if v62&int32(1) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v62&int32(1) == int32(0) {
		v309 = v88
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v291 = v65
	v292 = v94 & int32(32767)
	v295 = v69
	v296 = v70
	goto L20
L24:
	;
	v122 = int32(0)
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v109)+18)))
	if v123 < v122 {
		v309 = v122
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v65 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	if v70&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+20)))
	v130 = int32(768)
	if v129&v130 != v130 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v136 = v134
	goto L32
L31:
	;
	v136 = int32(2)
	goto L32
L32:
	;
	if v136 == v65 {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v309 = int32(0)
	goto L8
L34:
	;
	if l5 == int32(0) {
		v213 = v69
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v141 = F_HeapTupleSatisfiesVisibility(m, l4, l3, l2)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return int32(0)
L37:
	;
	F_HeapCheckForSerializableConflictOut(m, v141, l1, l4, l2, l3)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if v141 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v66)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+20)))
	v152 = int32(768)
	if v151&v152 != v152 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v158 = v156
	goto L42
L41:
	;
	v158 = int32(2)
	goto L42
L42:
	;
	F_PredicateLockTID(m, l1, l4+int32(4), l3, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v161 = int32(1)
	if l5 == int32(0) {
		v309 = v161
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v164)
	v309 = v161
	goto L8
L45:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+19)))
	if v215&int32(64) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L46:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v169 != int32(1) {
		v213 = v69
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if v69 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v174 = F_GlobalVisHorizonKindForRel(m, l1)
	mBase = m.M
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174<<(uint(int32(2))%32))+uint32(_consts[27])))
	goto L51
L49:
	;
	v180 = v69
	goto L50
L50:
	;
	v181 = int32(0)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+20)))
	if v183&int32(256) == v181 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v180 = v179
	goto L50
L52:
	;
	if v207 != 0 {
		v213 = v180
		goto L45
	} else {
		goto L61
	}
L53:
	;
	v207 = int32(base.Ui32(v183&int32(512)) >> (uint(int32(9)) % 32))
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v183&int32(2048) != 0 {
		v205 = v181
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v207 = v205
	goto L52
L57:
	;
	if v183&int32(5248) != int32(1024) {
		v205 = v181
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if v183&int32(4176) == int32(64) {
		v205 = v181
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v203 = F_GlobalVisTestIsRemovableXid(m, v180, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L36
	} else {
		goto L60
	}
L60:
	;
	v205 = v203
	goto L56
L61:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v208)
	v213 = v180
	goto L45
L62:
	;
	v309 = int32(0)
	goto L8
L63:
	;
	goto L64
L64:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+20)))
	if v221&int32(2048) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v309 = int32(0)
	goto L8
L66:
	;
	goto L67
L67:
	;
	if v221&int32(768) == int32(512) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v309 = int32(0)
	goto L8
L69:
	;
	goto L70
L70:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+16)))
	v232 = int32(0)
	if v221&int32(4224) != int32(4096) {
		v291 = v230
		v292 = v231
		v295 = v213
		v296 = v232
		goto L20
	} else {
		goto L71
	}
L71:
	;
	v240 = F_GetMultiXactIdMembers(m, v230, v23+int32(12), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L36
	} else {
		goto L72
	}
L72:
	;
	v242 = int32(0)
	if v240 <= v242 {
		v291 = v242
		v292 = v231
		v295 = v213
		v296 = v232
		goto L20
	} else {
		goto L73
	}
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v253 = int32(0)
	goto L76
L74:
	;
	F_pfree(m, v246)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L36
	} else {
		goto L80
	}
L75:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v279 = v277
	goto L74
L76:
	;
	v269 = v246 + v253<<(uint(int32(3))%32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v270) {
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v279 = int32(0)
	goto L74
L78:
	;
	v274 = v253 + int32(1)
	if v274 != v240 {
		v253 = v274
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v291 = v279
	v292 = v231
	v295 = v213
	v296 = v232
	goto L20
L81:
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
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
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
	v35 = int32(4510148)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v37 + int32(1)
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
	v77 = v74 & int32(65531)
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
	v50 = *(*int32)(unsafe.Add(mBase, _consts[5]))
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
	v61 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v64 = v61 + v29<<(uint(int32(13))%32)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64-int32(8182)))))
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
	v261 = int32(4510148)
	v263 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v263 - int32(1)
	F_UnlockReleaseBuffer(m, v29)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L68
	}
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v107 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v29^int32(-1))<<(uint(int32(2))%32))))
	v121 = v113
	goto L23
L25:
	;
	goto L26
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	if base.B2i32(base.Ui32(v124) < base.Ui32(int32(12000))) == int32(0) {
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
	v158 = base.B2i32(base.Ui32(int32(24)) < base.Ui32(v149)) & base.B2i32((v149+int32(262120))&int32(262140) == int32(4))
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
	v171 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
		goto L61
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
	if l3&int32(8) != 0 {
		v204 = v163
		goto L49
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32(v181) < base.Ui32(int32(12000)) {
		v204 = v163
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v187 = v168 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)) = uint8(v187)
	v190 = v163 | int32(16)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+68))
	if v194 != int32(99) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v198 == int32(0) {
		v204 = v190
		goto L49
	} else {
		goto L60
	}
L57:
	;
	v197 = F_isTempToastNamespace(m, v194)
	mBase = m.M
	v198 = v197
	goto L59
L58:
	;
	v198 = int32(1)
	goto L59
L59:
	;
	goto L56
L60:
	;
	v202 = v168 | int32(24)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+10)) = uint8(v202)
	v204 = v190
	goto L49
L61:
	;
	F_XLogRegisterData(m, v15+int32(8), int32(3))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
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
		goto L63
	}
L63:
	;
	F_XLogRegisterBufData(m, int32(0), v15+int32(2), int32(5))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
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
		goto L65
	}
L65:
	;
	v241 = int32(4411332)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	v244 = v243 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[28])) = uint8(v244)
	goto L66
L66:
	;
	v249 = F_XLogInsert(m, int32(10), v164&int32(255))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = base.I64_rotr(v249, int64(32))
	goto L16
L68:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v269 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_ReleaseBuffer(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_CacheInvalidateHeapTuple(m, l0, v21, int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_pgstat_count_heap_insert(m, l0, int64(1))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if l1 != v21 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v279
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v281)
	F_pfree(m, v21)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	m.G0 = v15 + int32(16)
	return
L78:
	;
	goto L77
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
		v25 = F_build_reloptions(m, l1, v20, v20, int32(128), int32(757504), int32(24))
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
		v10 = F_build_reloptions(m, l1, int32(1), int32(2), int32(128), int32(757504), int32(24))
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
	return int32(758388)
}
func F_heap_toast_insert_or_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
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
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
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
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v732 int32
	_ = v732
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1045 int64
	_ = v1045
	var v1047 int64
	_ = v1047
	var v1049 int64
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1094 int32
	_ = v1094
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1151 int32
	_ = v1151
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(35232)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_heap_deform_tuple(m, l1, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
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
	F_heap_deform_tuple(m, l2, v21, v19+int32(19232), v19+int32(32032))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v42 = v5
	v43 = int32(0)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v19 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v19 + int32(33632)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v19 + int32(25632)
	v57 = v19 + int32(4)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)) = uint8(v61)
	if v61 < v60 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v42 = v19 + int32(32032)
	v43 = v19 + int32(19232)
	goto L5
L7:
	;
	v77 = v5
	goto L10
L8:
	;
	goto L9
L9:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v358&int32(4) != 0 {
		goto L75
	} else {
		goto L76
	}
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v85 = v77 * int32(12)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+v86)+8)) = uint8(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v90+v85))) = v88
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v101 = v59 + int32(20) + v83<<(uint(int32(4))%32) + v77*int32(100)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+85)))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v85)+9)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v104 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L9
L12:
	;
	v339 = v77 + int32(1)
	if v339 != v60 {
		v77 = v339
		goto L10
	} else {
		goto L74
	}
L13:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v77))))
	if v228 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L14:
	;
	v106 = v77 << (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106+v107)))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+72)))
	if v110 != int32(65535) {
		v223 = v109
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218+v77<<(uint(int32(2))%32))))
	v223 = v222
	goto L13
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v77))))
	if v115 != 0 {
		v223 = v109
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v104+v106)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v118 != int32(1) {
		v223 = v109
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v121 != int32(18) {
		v223 = v109
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v77))))
	if v126 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v213 = v210 + v85 + int32(8)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v216 = v214 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v216)
	goto L12
L22:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v201 = v198 + v85 + int32(8)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v203 = int32(1)
	v204 = v202 | v203
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v204)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)))
	v208 = v206 | v203
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)) = uint8(v208)
	v223 = v109
	goto L13
L23:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v127 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v130 != int32(18) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v133 = int32(18)
	goto L29
L26:
	;
	if v195 == int32(0) {
		goto L21
	} else {
		goto L44
	}
L27:
	;
	v195 = int32(0)
	goto L26
L28:
	;
	v169 = v164
	v170 = v165
	v171 = v166
	goto L38
L29:
	;
	if (v117|v109)&int32(3) != 0 {
		v164 = v117
		v165 = v109
		v166 = v133
		goto L28
	} else {
		goto L32
	}
L31:
	;
	if v154 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v141 = v117
	v142 = v109
	v143 = v133
	goto L33
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v146 != v147 {
		v164 = v141
		v165 = v142
		v166 = v143
		goto L28
	} else {
		goto L35
	}
L34:
	;
	goto L31
L35:
	;
	v149 = int32(4)
	v150 = v142 + v149
	v152 = v141 + v149
	v154 = v143 - v149
	if base.Ui32(int32(3)) < base.Ui32(v154) {
		v141 = v152
		v142 = v150
		v143 = v154
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v164 = v152
	v165 = v150
	v166 = v154
	goto L28
L38:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v174 == v175 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v195 = v174 - v175
	goto L26
L40:
	;
	v177 = int32(1)
	v182 = v171 - v177
	if v182 != 0 {
		v169 = v169 + v177
		v170 = v170 + v177
		v171 = v182
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L27
L44:
	;
	goto L22
L45:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v234 = v231 + v85 + int32(8)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	v237 = v235 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v237)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)))
	v241 = v239 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)) = uint8(v241)
	goto L12
L46:
	;
	goto L47
L47:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+72)))
	if v243 == int32(65535) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+84)))
	if v246 == int32(112) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v328 = v325 + v85 + int32(8)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	v331 = v329 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v328))) = uint8(v331)
	goto L12
L51:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v252 = v249 + v85 + int32(8)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v255 = v253 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v255)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v258 != int32(1) {
		v309 = v223
		v310 = v258
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v322+v85)+4)) = v321
	goto L12
L55:
	;
	v311 = int32(1)
	if v310&v311 != 0 {
		v321 = int32(base.Ui32(v310) >> (uint(v311) % 32))
		goto L54
	} else {
		goto L73
	}
L56:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v261+v85))) = v223
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+84)))
	if v264 == int32(112) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v273 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v272+v77<<(uint(v273)%32)))) = v271
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v280 = v277 + v85 + int32(8)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	v283 = v281 | v273
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v283)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)))
	v287 = v285 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)) = uint8(v287)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v289 != int32(1) {
		v309 = v271
		v310 = v289
		goto L55
	} else {
		goto L63
	}
L58:
	;
	v267 = F_detoast_attr(m, v223)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v269 = F_detoast_external_attr(m, v223)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v271 = v267
	goto L57
L62:
	;
	v271 = v269
	goto L57
L63:
	;
	v292 = int32(6)
	v294 = int32(18)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	if v296 == v294 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v299 = v294
	goto L66
L65:
	;
	v299 = int32(2)
	goto L66
L66:
	;
	if v296&int32(254) == int32(2) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v304 = v292
	goto L69
L68:
	;
	v304 = v299
	goto L69
L69:
	;
	if v296 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v307 = v292
	goto L72
L71:
	;
	v307 = v304
	goto L72
L72:
	;
	v321 = v307
	goto L54
L73:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v321 = int32(base.Ui32(v315) >> (uint(int32(2)) % 32))
	goto L54
L74:
	;
	goto L11
L75:
	;
	v364 = base.I32_div_s(v22+int32(7), int32(8))
	v369 = (v364 + int32(30)) & int32(-8)
	goto L77
L76:
	;
	v369 = int32(24)
	goto L77
L77:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v370 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+8))
	v373 = v371
	goto L80
L79:
	;
	v373 = int32(2032)
	goto L80
L80:
	;
	v375 = l3 & int32(-17)
	v380 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v562 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L127
	}
L82:
	;
	v382 = v373 - v369
	if base.Ui32(v380) <= base.Ui32(v382) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	goto L84
L84:
	;
	v403 = v19 + int32(4)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+52))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	if v417 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L81
L86:
	;
	if v495 < int32(0) {
		goto L81
	} else {
		goto L114
	}
L87:
	;
	v495 = int32(-1)
	goto L86
L88:
	;
	goto L89
L89:
	;
	goto L90
L90:
	;
	goto L92
L92:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v403)+24))
	v436 = int32(0)
	v439 = int32(24)
	v440 = int32(-1)
	goto L93
L93:
	;
	v447 = v429 + v436*int32(12)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+8)))
	if int32(48)&v448 != 0 {
		v478 = v439
		v479 = v440
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v495 = v479
	goto L86
L95:
	;
	v481 = v436 + int32(1)
	if v481 != v417 {
		v436 = v481
		v439 = v478
		v440 = v479
		goto L93
	} else {
		goto L113
	}
L96:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v450+v436<<(uint(int32(2))%32))))
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	if v455 == int32(1) {
		v478 = v439
		v479 = v440
		goto L95
	} else {
		goto L97
	}
L97:
	;
	goto L98
L98:
	;
	goto L100
L100:
	;
	if v455&int32(3) == int32(2) {
		v478 = v439
		v479 = v440
		goto L95
	} else {
		goto L101
	}
L101:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v417<<(uint(int32(4))%32)+int32(20)+v436*int32(100))+84)))
	goto L104
L102:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v473 = base.B2i32(v439 < v472)
	if v439 < v472 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	goto L105
L105:
	;
	switch v467 - int32(101) {
	case 0, 19:
		goto L102
	default:
		v478 = v439
		v479 = v440
		goto L95
	}
L107:
	;
	v474 = v472
	goto L109
L108:
	;
	v474 = v439
	goto L109
L109:
	;
	if v439 < v472 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v475 = v436
	goto L112
L111:
	;
	v475 = v440
	goto L112
L112:
	;
	v478 = v474
	v479 = v475
	goto L95
L113:
	;
	goto L94
L114:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(104)+v498<<(uint(int32(4))%32)+v495*int32(100)))))
	if v505 == int32(120) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v495*int32(12)+v19)+36))
	if base.Ui32(v525) <= base.Ui32(v382) {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	F_toast_tuple_try_compression(m, v19+int32(4), v495)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v516 = v495*int32(12) + v19 + int32(40)
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	v519 = v517 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v516))) = uint8(v519)
	goto L115
L119:
	;
	goto L115
L120:
	;
	v539 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L124
	}
L121:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)+112))
	if v528 == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	F_toast_tuple_externalize(m, v19+int32(4), v495, v375)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L120
L124:
	;
	if base.Ui32(v382) < base.Ui32(v539) {
		goto L84
	} else {
		goto L125
	}
L125:
	;
	goto L85
L126:
	;
	v712 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L165
	}
L127:
	;
	if base.Ui32(v562) <= base.Ui32(v382) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	goto L129
L129:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+112))
	if v582 == int32(0) {
		goto L126
	} else {
		goto L131
	}
L130:
	;
	goto L126
L131:
	;
	v586 = v19 + int32(4)
	v587 = int32(0)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v598)+52))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	if v600 <= v587 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v678 < int32(0) {
		goto L126
	} else {
		goto L160
	}
L133:
	;
	v678 = int32(-1)
	goto L132
L134:
	;
	goto L135
L135:
	;
	goto L137
L137:
	;
	goto L138
L138:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v586)+24))
	v619 = int32(0)
	v622 = int32(24)
	v623 = int32(-1)
	goto L139
L139:
	;
	v630 = v612 + v619*int32(12)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+8)))
	if int32(16)&v631 != 0 {
		v661 = v622
		v662 = v623
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v678 = v662
	goto L132
L141:
	;
	v664 = v619 + int32(1)
	if v664 != v600 {
		v619 = v664
		v622 = v661
		v623 = v662
		goto L139
	} else {
		goto L159
	}
L142:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v633+v619<<(uint(int32(2))%32))))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	if v638 == int32(1) {
		v661 = v622
		v662 = v623
		goto L141
	} else {
		goto L143
	}
L143:
	;
	if v638&int32(3) == int32(2) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v646 = v587
	goto L146
L145:
	;
	v646 = int32(0)
	goto L146
L146:
	;
	if v646 != 0 {
		v661 = v622
		v662 = v623
		goto L141
	} else {
		goto L147
	}
L147:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v600<<(uint(int32(4))%32)+int32(20)+v619*int32(100))+84)))
	goto L150
L148:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v656 = base.B2i32(v622 < v655)
	if v622 < v655 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	goto L151
L151:
	;
	switch v650 - int32(101) {
	case 0, 19:
		goto L148
	default:
		v661 = v622
		v662 = v623
		goto L141
	}
L153:
	;
	v657 = v655
	goto L155
L154:
	;
	v657 = v622
	goto L155
L155:
	;
	if v622 < v655 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v658 = v619
	goto L158
L157:
	;
	v658 = v623
	goto L158
L158:
	;
	v661 = v657
	v662 = v658
	goto L141
L159:
	;
	goto L140
L160:
	;
	F_toast_tuple_externalize(m, v19+int32(4), v678, v375)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v689 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if base.Ui32(v382) < base.Ui32(v689) {
		goto L129
	} else {
		goto L163
	}
L163:
	;
	goto L130
L164:
	;
	v858 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L202
	}
L165:
	;
	if base.Ui32(v712) <= base.Ui32(v382) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	goto L167
L167:
	;
	v732 = v19 + int32(4)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+52))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	if v746 <= int32(0) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L164
L169:
	;
	if v824 < int32(0) {
		goto L164
	} else {
		goto L197
	}
L170:
	;
	v824 = int32(-1)
	goto L169
L171:
	;
	goto L172
L172:
	;
	goto L173
L173:
	;
	goto L175
L175:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v732)+24))
	v765 = int32(0)
	v768 = int32(24)
	v769 = int32(-1)
	goto L176
L176:
	;
	v776 = v758 + v765*int32(12)
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776)+8)))
	if int32(48)&v777 != 0 {
		v807 = v768
		v808 = v769
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v824 = v808
	goto L169
L178:
	;
	v810 = v765 + int32(1)
	if v810 != v746 {
		v765 = v810
		v768 = v807
		v769 = v808
		goto L176
	} else {
		goto L196
	}
L179:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v779+v765<<(uint(int32(2))%32))))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783))))
	if v784 == int32(1) {
		v807 = v768
		v808 = v769
		goto L178
	} else {
		goto L180
	}
L180:
	;
	goto L181
L181:
	;
	goto L183
L183:
	;
	if v784&int32(3) == int32(2) {
		v807 = v768
		v808 = v769
		goto L178
	} else {
		goto L184
	}
L184:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745+v746<<(uint(int32(4))%32)+int32(20)+v765*int32(100))+84)))
	goto L186
L185:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v802 = base.B2i32(v768 < v801)
	if v768 < v801 {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	if v796 != int32(109) {
		v807 = v768
		v808 = v769
		goto L178
	} else {
		goto L189
	}
L189:
	;
	goto L185
L190:
	;
	v803 = v801
	goto L192
L191:
	;
	v803 = v768
	goto L192
L192:
	;
	if v768 < v801 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v804 = v765
	goto L195
L194:
	;
	v804 = v769
	goto L195
L195:
	;
	v807 = v803
	v808 = v804
	goto L178
L196:
	;
	goto L177
L197:
	;
	F_toast_tuple_try_compression(m, v19+int32(4), v824)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v835 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	if base.Ui32(v382) < base.Ui32(v835) {
		goto L167
	} else {
		goto L200
	}
L200:
	;
	goto L168
L201:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v1006&int32(8) == int32(0) {
		goto L240
	} else {
		goto L241
	}
L202:
	;
	v861 = int32(8160) - v369
	if base.Ui32(v858) <= base.Ui32(v861) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	goto L204
L204:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)+112))
	if v880 == int32(0) {
		goto L201
	} else {
		goto L206
	}
L205:
	;
	goto L201
L206:
	;
	v884 = v19 + int32(4)
	v885 = int32(0)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v884)))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+52))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)))
	if v898 <= v885 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	if v976 < int32(0) {
		goto L201
	} else {
		goto L235
	}
L208:
	;
	v976 = int32(-1)
	goto L207
L209:
	;
	goto L210
L210:
	;
	goto L212
L212:
	;
	goto L213
L213:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v884)+24))
	v917 = int32(0)
	v920 = int32(24)
	v921 = int32(-1)
	goto L214
L214:
	;
	v928 = v910 + v917*int32(12)
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928)+8)))
	if int32(16)&v929 != 0 {
		v959 = v920
		v960 = v921
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v976 = v960
	goto L207
L216:
	;
	v962 = v917 + int32(1)
	if v962 != v898 {
		v917 = v962
		v920 = v959
		v921 = v960
		goto L214
	} else {
		goto L234
	}
L217:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v884)+4))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v931+v917<<(uint(int32(2))%32))))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	if v936 == int32(1) {
		v959 = v920
		v960 = v921
		goto L216
	} else {
		goto L218
	}
L218:
	;
	if v936&int32(3) == int32(2) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v944 = v885
	goto L221
L220:
	;
	v944 = int32(0)
	goto L221
L221:
	;
	if v944 != 0 {
		v959 = v920
		v960 = v921
		goto L216
	} else {
		goto L222
	}
L222:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897+v898<<(uint(int32(4))%32)+int32(20)+v917*int32(100))+84)))
	goto L224
L223:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v928)+4))
	v954 = base.B2i32(v920 < v953)
	if v920 < v953 {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	if v948 != int32(109) {
		v959 = v920
		v960 = v921
		goto L216
	} else {
		goto L227
	}
L227:
	;
	goto L223
L228:
	;
	v955 = v953
	goto L230
L229:
	;
	v955 = v920
	goto L230
L230:
	;
	if v920 < v953 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v956 = v917
	goto L233
L232:
	;
	v956 = v921
	goto L233
L233:
	;
	v959 = v955
	v960 = v956
	goto L216
L234:
	;
	goto L215
L235:
	;
	F_toast_tuple_externalize(m, v19+int32(4), v976, v375)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v987 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	if base.Ui32(v861) < base.Ui32(v987) {
		goto L204
	} else {
		goto L238
	}
L238:
	;
	goto L205
L239:
	;
	v1082 = v19 + int32(4)
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+52))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+20)))
	if v1086&int32(2) == int32(0) {
		v1128 = v1086
		goto L249
	} else {
		goto L250
	}
L240:
	;
	v1077 = l1
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1006&int32(4) != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1018 = base.I32_div_s(v22+int32(7), int32(8))
	v1023 = (v1018 + int32(30)) & int32(-8)
	goto L245
L244:
	;
	v1023 = int32(24)
	goto L245
L245:
	;
	v1028 = F_heap_compute_data_size(m, v21, v19+int32(25632), v19+int32(33632))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v1030 = v1028 + v1023
	v1033 = F_palloc0(m, v1030+int32(24))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v1030
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+4)) = v1036
	v1038 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1033)+8)) = uint16(v1038)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1042 = v1033 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+16)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+12)) = v1040
	v1045 = *(*int64)(unsafe.Add(mBase, uint32(v1011)+15))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+39)) = v1045
	v1047 = *(*int64)(unsafe.Add(mBase, uint32(v1011)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+32)) = v1047
	v1049 = *(*int64)(unsafe.Add(mBase, uint32(v1011)))
	*(*int64)(unsafe.Add(mBase, uint32(v1042))) = v1049
	*(*uint8)(unsafe.Add(mBase, uint32(v1033)+46)) = uint8(v1023)
	v1053 = v1033 + int32(42)
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053))))
	v1057 = v1054&int32(63488) | v22
	*(*uint16)(unsafe.Add(mBase, uint32(v1053))) = uint16(v1057)
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	F_heap_fill_tuple(m, v21, v19+int32(25632), v19+int32(33632), v1023+v1042, v1033+int32(44), (v1033+int32(47))&(v1068<<(uint(int32(29))%32)>>(uint(int32(31))%32)))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	v1077 = v1033
	goto L239
L249:
	;
	if v1128&int32(1) == int32(0) {
		goto L259
	} else {
		goto L260
	}
L250:
	;
	if v1085 <= int32(0) {
		v1128 = v1086
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v1094 = int32(0)
	goto L252
L252:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+24))
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110+v1094*int32(12))+8)))
	if v1114&int32(2) != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+20)))
	v1128 = v1127
	goto L249
L254:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1094<<(uint(int32(2))%32))))
	F_pfree(m, v1121)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1125 = v1094 + int32(1)
	if v1125 != v1085 {
		v1094 = v1125
		goto L252
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	goto L253
L259:
	;
	m.G0 = v19 + int32(35232)
	return v1077
L260:
	;
	if v1085 <= int32(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1151 = int32(0)
	goto L262
L262:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+24))
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1167+v1151*int32(12))+8)))
	if v1171&int32(1) != 0 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	goto L259
L264:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+12))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1175+v1151<<(uint(int32(2))%32))))
	F_toast_delete_datum(m, v1179, int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1184 = v1151 + int32(1)
	if v1184 != v1085 {
		v1151 = v1184
		goto L262
	} else {
		goto L268
	}
L267:
	;
	goto L266
L268:
	;
	goto L263
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
						v37 = F_LookupCreationNamespace(m, int32(235617))
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
							v51 = F_pg_snprintf(m, v17+int32(48), int32(64), int32(38532), v17+int32(32))
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
															F_errmsg_internal(m, int32(46291), v17+int32(16))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(495341), int32(806), int32(238827))
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
																	F_sequence_close(m, v98, int32(0))
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
																			F_sequence_close(m, v19, int32(0))
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
												F_sequence_close(m, v19, int32(0))
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
						v51 = F_pg_snprintf(m, v17+int32(48), int32(64), int32(38532), v17+int32(32))
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
														F_errmsg_internal(m, int32(46291), v17+int32(16))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(495341), int32(806), int32(238827))
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
																F_sequence_close(m, v98, int32(0))
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
																		F_sequence_close(m, v19, int32(0))
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
											F_sequence_close(m, v19, int32(0))
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
					F_errmsg_internal(m, int32(46291), v17)
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495341), int32(733), int32(238827))
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
