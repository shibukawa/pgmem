package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GetNewRelFileNumber(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	switch l2 - int32(112) {
	case 0, 5:
		v37 = int32(-1)
		goto L1
	default:
		goto L3
	case 4:
		goto L2
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewRelFileNumber[0]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewRelFileNumber[1]))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewRelFileNumber[2]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewRelFileNumber[3]))
	if v32 == int32(-1) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	F_errmsg_internal(m, int32(_a_F_GetNewRelFileNumber_0), v9)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(_a_F_GetNewRelFileNumber_1), int32(581), int32(_a_F_GetNewRelFileNumber_2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L8:
	;
	v35 = v30
	goto L10
L9:
	;
	v35 = v32
	goto L10
L10:
	;
	v37 = v35
	goto L1
L11:
	;
	v43 = l0
	goto L13
L12:
	;
	v43 = v42
	goto L13
L13:
	;
	if v43 != int32(1664) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = v39
	goto L16
L15:
	;
	v46 = int32(0)
	goto L16
L16:
	;
	goto L17
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewRelFileNumber[4]))
	if v54 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	m.G0 = v9 + int32(80)
	return v65
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l1 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	F_GetRelationPath(m, v9+int32(8), v46, v43, v65, v37, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L29
	}
L24:
	;
	v61 = F_GetNewOidWithIndex(m, l1, int32(2662), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v63 = F_GetNewObjectId(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	v65 = v61
	goto L23
L28:
	;
	v65 = v63
	goto L23
L29:
	;
	v71 = int32(0)
	v72 = F_access(m, v9+int32(8), v71)
	mBase = m.M
	if v72 == v71 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L18
}
func F_ScanRelIsReadOnly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+72))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	v8 = F_bms_is_member(m, v4, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v18 = int32(0)
			return v18
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
			v14 = F_bms_is_member(m, v4, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v18 = v14 ^ int32(1)
				return v18
			}
		}
	}
}
func F_StoreRelCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	v10 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = F_nodeToString(m, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = F_pull_var_clause(m, l2, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if l7 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	if v29 == int32(0) {
		v177 = v10
		v181 = v10
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v33 <= int32(0) {
		v177 = v33
		v181 = v10
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v38 = F_palloc(m, v33<<(uint(int32(1))%32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v41 <= v40 {
		v177 = v40
		v181 = v38
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v53 = v40
	v59 = v41
	v60 = v10
	goto L9
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v60<<(uint(int32(2))%32))))
	v68 = int32(0)
	if v53 <= v68 {
		v110 = v68
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v177 = v155
	v181 = v38
	goto L3
L11:
	;
	v166 = v60 + int32(1)
	if v166 < v161 {
		v53 = v155
		v59 = v161
		v60 = v166
		goto L9
	} else {
		goto L20
	}
L12:
	;
	v138 = int32(1)
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v38+v53<<(uint(v138)%32)))) = uint16(v141)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v155 = v53 + v138
	v161 = v145
	goto L11
L13:
	;
	if v53 != v110 {
		v155 = v53
		v161 = v59
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	v83 = v68
	goto L15
L15:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+v83<<(uint(int32(1))%32)))))
	if v94 == v71 {
		v110 = v83
		goto L13
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v97 = v83 + int32(1)
	if v97 != v53 {
		v83 = v97
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L12
L20:
	;
	goto L10
L21:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v187)+68))
	v214 = int32(0)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v226 = int32(32)
	v233 = F_CreateConstraintEntry(m, l1, v212, int32(99), v214, v214, l3, l4, v214, v217, v181, v177, v177, v214, v214, v214, v214, v214, v214, v214, v214, v226, v226, v214, v214, v226, v214, l2, v24, l5, l6, l7, v214, l8)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+119)))
	if v190 != int32(112) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v200 + int32(4)
	F_errmsg(m, int32(_a_F_StoreRelCheck_0), v22)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_StoreRelCheck_1), int32(2224), int32(_a_F_StoreRelCheck_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_pfree(m, v24)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v22 + int32(16)
	return v233
}
func F_get_rel_persistence(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_rel_persistence_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_rel_persistence_1), int32(2398), int32(_a_F_get_rel_persistence_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29+v30)+118)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_get_rel_relispartition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v10)+131)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v12
				return v15 & int32(1)
			}
		} else {
			v15 = int32(0)
			return v15 & int32(1)
		}
	}
}
func F_make_rel_from_joinlist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v195 int32
	_ = v195
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
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 float64
	_ = v404
	var v409 int64
	_ = v409
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v427 int64
	_ = v427
	var v432 int64
	_ = v432
	var v437 int64
	_ = v437
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 float64
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v471 float64
	_ = v471
	var v473 int64
	_ = v473
	var v493 float64
	_ = v493
	var v498 float64
	_ = v498
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 float64
	_ = v509
	var v510 float64
	_ = v510
	var v512 float64
	_ = v512
	var v513 float64
	_ = v513
	var v516 float64
	_ = v516
	var v519 float64
	_ = v519
	var v523 float64
	_ = v523
	var v526 float64
	_ = v526
	var v530 float64
	_ = v530
	var v532 int64
	_ = v532
	var v537 int32
	_ = v537
	var v538 float64
	_ = v538
	var v541 float64
	_ = v541
	var v542 int64
	_ = v542
	var v545 int64
	_ = v545
	var v554 float64
	_ = v554
	var v556 float64
	_ = v556
	var v560 float64
	_ = v560
	var v561 float64
	_ = v561
	var v562 float64
	_ = v562
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 float64
	_ = v571
	var v574 int32
	_ = v574
	var v578 float64
	_ = v578
	var v579 float64
	_ = v579
	var v580 float64
	_ = v580
	var v589 float64
	_ = v589
	var v592 float64
	_ = v592
	var v595 float64
	_ = v595
	var v602 float64
	_ = v602
	var v603 float64
	_ = v603
	var v614 float64
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v841 int64
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v904 int64
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 float64
	_ = v911
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1016 int32
	_ = v1016
	var v1043 float64
	_ = v1043
	var v1045 float64
	_ = v1045
	var v1047 float64
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 float64
	_ = v1049
	var v1087 float64
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1097 float64
	_ = v1097
	var v1099 float64
	_ = v1099
	var v1103 float64
	_ = v1103
	var v1108 float64
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1114 float64
	_ = v1114
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1161 float64
	_ = v1161
	var v1163 float64
	_ = v1163
	var v1167 float64
	_ = v1167
	var v1172 float64
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1223 int32
	_ = v1223
	var v1224 float64
	_ = v1224
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1271 float64
	_ = v1271
	var v1273 float64
	_ = v1273
	var v1277 float64
	_ = v1277
	var v1282 float64
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1340 int32
	_ = v1340
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1396 int32
	_ = v1396
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1433 int64
	_ = v1433
	var v1435 int64
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1451 int32
	_ = v1451
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1507 int32
	_ = v1507
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1544 int64
	_ = v1544
	var v1546 int64
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1611 int32
	_ = v1611
	var v1612 int64
	_ = v1612
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1630 int32
	_ = v1630
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1750 int32
	_ = v1750
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1810 int32
	_ = v1810
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1854 int32
	_ = v1854
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1960 int32
	_ = v1960
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2004 int32
	_ = v2004
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2109 int32
	_ = v2109
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2304 int32
	_ = v2304
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2453 int64
	_ = v2453
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2485 int32
	_ = v2485
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2518 int32
	_ = v2518
	var v2519 int64
	_ = v2519
	var v2521 int64
	_ = v2521
	var v2523 int64
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
	var v2566 int32
	_ = v2566
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2585 int32
	_ = v2585
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2630 int32
	_ = v2630
	var v2640 int32
	_ = v2640
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2690 int32
	_ = v2690
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2726 int64
	_ = v2726
	var v2728 int64
	_ = v2728
	var v2730 int64
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2780 int32
	_ = v2780
	var v2788 int32
	_ = v2788
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2800 int32
	_ = v2800
	var v2804 int32
	_ = v2804
	var v2806 int32
	_ = v2806
	var v2810 int32
	_ = v2810
	var v2815 int64
	_ = v2815
	var v2817 int32
	_ = v2817
	var v2824 int32
	_ = v2824
	var v2859 int32
	_ = v2859
	var v2863 int32
	_ = v2863
	var v2866 int32
	_ = v2866
	var v2868 int32
	_ = v2868
	var v2870 int32
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2925 int32
	_ = v2925
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2933 int32
	_ = v2933
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2943 int32
	_ = v2943
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2953 int32
	_ = v2953
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2997 int32
	_ = v2997
	var v2998 int32
	_ = v2998
	var v3003 int32
	_ = v3003
	var v3007 int32
	_ = v3007
	var v3010 int32
	_ = v3010
	var v3046 int32
	_ = v3046
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3056 int32
	_ = v3056
	var v3061 int64
	_ = v3061
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3103 int32
	_ = v3103
	var v3104 int32
	_ = v3104
	var v3107 int32
	_ = v3107
	var v3111 int32
	_ = v3111
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3131 int32
	_ = v3131
	var v3169 int32
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3193 int32
	_ = v3193
	var v3198 int64
	_ = v3198
	var v3200 int32
	_ = v3200
	var v3202 int32
	_ = v3202
	var v3241 int32
	_ = v3241
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3261 int32
	_ = v3261
	var v3296 int32
	_ = v3296
	var v3299 int32
	_ = v3299
	var v3303 int32
	_ = v3303
	var v3344 int32
	_ = v3344
	var v3348 int32
	_ = v3348
	var v3353 int32
	_ = v3353
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3362 int32
	_ = v3362
	var v3367 int64
	_ = v3367
	var v3372 int32
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3381 int32
	_ = v3381
	var v3422 int32
	_ = v3422
	var v3426 int32
	_ = v3426
	var v3431 int32
	_ = v3431
	var v3434 int32
	_ = v3434
	var v3470 int32
	_ = v3470
	var v3477 int32
	_ = v3477
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3525 int64
	_ = v3525
	var v3527 int64
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3532 int32
	_ = v3532
	var v3534 float64
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3536 int32
	_ = v3536
	var v3537 int32
	_ = v3537
	var v3540 int32
	_ = v3540
	var v3543 int32
	_ = v3543
	var v3548 float64
	_ = v3548
	var v3555 int32
	_ = v3555
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3562 int32
	_ = v3562
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3599 float64
	_ = v3599
	var v3603 int32
	_ = v3603
	var v3604 float64
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3619 float64
	_ = v3619
	var v3633 int32
	_ = v3633
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3645 int32
	_ = v3645
	var v3649 int32
	_ = v3649
	var v3650 int32
	_ = v3650
	var v3651 int32
	_ = v3651
	var v3660 int32
	_ = v3660
	var v3668 int32
	_ = v3668
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3686 int32
	_ = v3686
	var v3688 int32
	_ = v3688
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3696 int32
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3702 int32
	_ = v3702
	var v3704 int32
	_ = v3704
	var v3707 int32
	_ = v3707
	var v3709 int32
	_ = v3709
	var v3716 int32
	_ = v3716
	var v3725 int32
	_ = v3725
	var v3730 int32
	_ = v3730
	var v3732 int32
	_ = v3732
	var v3733 int32
	_ = v3733
	var v3735 int32
	_ = v3735
	var v3737 int32
	_ = v3737
	var v3739 int32
	_ = v3739
	var v3742 int32
	_ = v3742
	var v3753 int64
	_ = v3753
	var v3755 int64
	_ = v3755
	var v3757 int32
	_ = v3757
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3761 int32
	_ = v3761
	var v3764 int32
	_ = v3764
	var v3766 int32
	_ = v3766
	var v3767 int64
	_ = v3767
	var v3769 int64
	_ = v3769
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3814 int32
	_ = v3814
	var v3815 int32
	_ = v3815
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3818 int64
	_ = v3818
	var v3820 int64
	_ = v3820
	var v3823 int32
	_ = v3823
	var v3824 int32
	_ = v3824
	var v3825 int64
	_ = v3825
	var v3827 int64
	_ = v3827
	var v3829 int64
	_ = v3829
	var v3831 int64
	_ = v3831
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3878 int32
	_ = v3878
	var v3917 int32
	_ = v3917
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3921 int32
	_ = v3921
	var v3927 int32
	_ = v3927
	var v3931 int32
	_ = v3931
	var v3936 int32
	_ = v3936
	var v3938 int32
	_ = v3938
	var v3940 int32
	_ = v3940
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3945 int32
	_ = v3945
	var v3952 int32
	_ = v3952
	var v3988 int32
	_ = v3988
	var v3990 int32
	_ = v3990
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3995 int32
	_ = v3995
	var v4033 int32
	_ = v4033
	var v4035 int32
	_ = v4035
	var v4037 int32
	_ = v4037
	var v4039 int32
	_ = v4039
	var v4042 int32
	_ = v4042
	var v4046 int32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4054 int32
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4064 int32
	_ = v4064
	var v4098 int32
	_ = v4098
	var v4099 int32
	_ = v4099
	var v4101 int32
	_ = v4101
	var v4104 int32
	_ = v4104
	var v4105 int32
	_ = v4105
	var v4108 int32
	_ = v4108
	var v4110 int32
	_ = v4110
	var v4111 int32
	_ = v4111
	var v4114 int32
	_ = v4114
	var v4128 int32
	_ = v4128
	var v4154 int32
	_ = v4154
	var v4158 int32
	_ = v4158
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4164 int32
	_ = v4164
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4167 int32
	_ = v4167
	var v4170 int32
	_ = v4170
	var v4176 int32
	_ = v4176
	var v4178 int32
	_ = v4178
	var v4179 int32
	_ = v4179
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4185 int32
	_ = v4185
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4201 int32
	_ = v4201
	var v4204 int32
	_ = v4204
	var v4211 int32
	_ = v4211
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4217 int32
	_ = v4217
	var v4218 int32
	_ = v4218
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4225 int32
	_ = v4225
	var v4226 int32
	_ = v4226
	var v4227 int32
	_ = v4227
	var v4228 int32
	_ = v4228
	var v4229 int32
	_ = v4229
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4233 int32
	_ = v4233
	var v4236 int32
	_ = v4236
	var v4237 int32
	_ = v4237
	var v4247 int32
	_ = v4247
	var v4255 int32
	_ = v4255
	var v4258 int32
	_ = v4258
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4273 int32
	_ = v4273
	var v4306 int32
	_ = v4306
	var v4307 int32
	_ = v4307
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4326 int32
	_ = v4326
	var v4329 int32
	_ = v4329
	var v4330 int32
	_ = v4330
	var v4335 int32
	_ = v4335
	var v4342 int32
	_ = v4342
	var v4344 int32
	_ = v4344
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4349 int32
	_ = v4349
	var v4351 int32
	_ = v4351
	var v4358 int32
	_ = v4358
	var v4359 int32
	_ = v4359
	var v4360 int32
	_ = v4360
	var v4363 int32
	_ = v4363
	var v4364 int32
	_ = v4364
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4370 int32
	_ = v4370
	var v4371 int32
	_ = v4371
	var v4373 int32
	_ = v4373
	var v4376 int32
	_ = v4376
	var v4377 int32
	_ = v4377
	var v4384 int32
	_ = v4384
	var v4417 int32
	_ = v4417
	var v4421 int32
	_ = v4421
	var v4422 int32
	_ = v4422
	var v4423 int32
	_ = v4423
	var v4424 int32
	_ = v4424
	var v4434 int32
	_ = v4434
	var v4435 int32
	_ = v4435
	var v4437 int32
	_ = v4437
	var v4440 int32
	_ = v4440
	var v4441 int32
	_ = v4441
	var v4446 int32
	_ = v4446
	var v4453 int32
	_ = v4453
	var v4455 int32
	_ = v4455
	var v4457 int32
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4460 int32
	_ = v4460
	var v4462 int32
	_ = v4462
	var v4469 int32
	_ = v4469
	var v4472 int32
	_ = v4472
	var v4473 int32
	_ = v4473
	var v4475 int32
	_ = v4475
	var v4476 int32
	_ = v4476
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4556 int32
	_ = v4556
	var v4557 int32
	_ = v4557
	var v4562 int32
	_ = v4562
	var v4568 int32
	_ = v4568
	var v4600 int32
	_ = v4600
	var v4603 int32
	_ = v4603
	var v4621 int32
	_ = v4621
	var v4647 int32
	_ = v4647
	var v4651 int32
	_ = v4651
	var v4652 int32
	_ = v4652
	var v4653 int32
	_ = v4653
	var v4657 int32
	_ = v4657
	var v4658 int32
	_ = v4658
	var v4659 int32
	_ = v4659
	var v4660 int32
	_ = v4660
	var v4663 int32
	_ = v4663
	var v4669 int32
	_ = v4669
	var v4671 int32
	_ = v4671
	var v4672 int32
	_ = v4672
	var v4676 int32
	_ = v4676
	var v4677 int32
	_ = v4677
	var v4678 int32
	_ = v4678
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4686 int32
	_ = v4686
	var v4687 int32
	_ = v4687
	var v4694 int32
	_ = v4694
	var v4697 int32
	_ = v4697
	var v4704 int32
	_ = v4704
	var v4706 int32
	_ = v4706
	var v4707 int32
	_ = v4707
	var v4710 int32
	_ = v4710
	var v4711 int32
	_ = v4711
	var v4714 int32
	_ = v4714
	var v4715 int32
	_ = v4715
	var v4716 int32
	_ = v4716
	var v4717 int32
	_ = v4717
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4720 int32
	_ = v4720
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4723 int32
	_ = v4723
	var v4724 int32
	_ = v4724
	var v4725 int32
	_ = v4725
	var v4726 int32
	_ = v4726
	var v4729 int32
	_ = v4729
	var v4730 int32
	_ = v4730
	var v4740 int32
	_ = v4740
	var v4748 int32
	_ = v4748
	var v4751 int32
	_ = v4751
	var v4758 int32
	_ = v4758
	var v4759 int32
	_ = v4759
	var v4765 int32
	_ = v4765
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4803 int32
	_ = v4803
	var v4804 int32
	_ = v4804
	var v4805 int32
	_ = v4805
	var v4815 int32
	_ = v4815
	var v4816 int32
	_ = v4816
	var v4818 int32
	_ = v4818
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4827 int32
	_ = v4827
	var v4834 int32
	_ = v4834
	var v4836 int32
	_ = v4836
	var v4838 int32
	_ = v4838
	var v4839 int32
	_ = v4839
	var v4841 int32
	_ = v4841
	var v4843 int32
	_ = v4843
	var v4850 int32
	_ = v4850
	var v4851 int32
	_ = v4851
	var v4852 int32
	_ = v4852
	var v4855 int32
	_ = v4855
	var v4856 int32
	_ = v4856
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4862 int32
	_ = v4862
	var v4863 int32
	_ = v4863
	var v4903 int32
	_ = v4903
	var v4904 int32
	_ = v4904
	var v4944 int32
	_ = v4944
	var v4945 int32
	_ = v4945
	var v4984 int32
	_ = v4984
	var v4985 int32
	_ = v4985
	var v4986 int32
	_ = v4986
	var v5001 int32
	_ = v5001
	var v5027 int32
	_ = v5027
	var v5030 int32
	_ = v5030
	var v5033 int32
	_ = v5033
	var v5037 int32
	_ = v5037
	var v5043 int32
	_ = v5043
	var v5076 int32
	_ = v5076
	var v5080 int32
	_ = v5080
	var v5081 int32
	_ = v5081
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5093 int32
	_ = v5093
	var v5094 int32
	_ = v5094
	var v5096 int32
	_ = v5096
	var v5099 int32
	_ = v5099
	var v5100 int32
	_ = v5100
	var v5105 int32
	_ = v5105
	var v5112 int32
	_ = v5112
	var v5114 int32
	_ = v5114
	var v5116 int32
	_ = v5116
	var v5117 int32
	_ = v5117
	var v5119 int32
	_ = v5119
	var v5121 int32
	_ = v5121
	var v5128 int32
	_ = v5128
	var v5131 int32
	_ = v5131
	var v5132 int32
	_ = v5132
	var v5134 int32
	_ = v5134
	var v5135 int32
	_ = v5135
	var v5175 int32
	_ = v5175
	var v5176 int32
	_ = v5176
	var v5215 int32
	_ = v5215
	var v5253 int32
	_ = v5253
	var v5254 int32
	_ = v5254
	var v5258 int32
	_ = v5258
	var v5262 int32
	_ = v5262
	var v5267 int32
	_ = v5267
	var v5308 int32
	_ = v5308
	var v5312 int32
	_ = v5312
	var v5315 int32
	_ = v5315
	var v5316 int32
	_ = v5316
	var v5320 int32
	_ = v5320
	var v5356 int32
	_ = v5356
	var v5360 int32
	_ = v5360
	var v5361 int32
	_ = v5361
	var v5362 int32
	_ = v5362
	var v5363 int32
	_ = v5363
	var v5377 int32
	_ = v5377
	var v5378 int32
	_ = v5378
	var v5380 int32
	_ = v5380
	var v5383 int32
	_ = v5383
	var v5384 int32
	_ = v5384
	var v5389 int32
	_ = v5389
	var v5397 int32
	_ = v5397
	var v5399 int32
	_ = v5399
	var v5401 int32
	_ = v5401
	var v5402 int32
	_ = v5402
	var v5405 int32
	_ = v5405
	var v5409 int32
	_ = v5409
	var v5415 int32
	_ = v5415
	var v5420 int32
	_ = v5420
	var v5422 int32
	_ = v5422
	var v5423 int32
	_ = v5423
	var v5424 int32
	_ = v5424
	var v5432 int32
	_ = v5432
	var v5434 int32
	_ = v5434
	var v5436 int32
	_ = v5436
	var v5437 int32
	_ = v5437
	var v5477 int32
	_ = v5477
	var v5516 int32
	_ = v5516
	var v5520 int32
	_ = v5520
	var v5526 int32
	_ = v5526
	var v5530 int32
	_ = v5530
	var v5535 int32
	_ = v5535
	var v5536 int32
	_ = v5536
	var v5537 int32
	_ = v5537
	var v5546 int32
	_ = v5546
	v3 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(16)
	m.G0 = v40
	if l1 == v3 {
		v5546 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v40 + int32(16)
	return v5546
L2:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v44 <= int32(0) {
		v5546 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v50 = v3
	v51 = v3
	goto L4
L4:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v50<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v44 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v116 = F_lappend(m, v51, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L13
	} else {
		goto L19
	}
L7:
	;
	if v89 == int32(63) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v113 = F_make_rel_from_joinlist(m, l0, v88)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L18
	}
L10:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v95 = F_find_base_rel(m, l0, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v115 = v95
	goto L6
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v103
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_0), v40)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_1), int32(3878), int32(_a_F_make_rel_from_joinlist_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
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
	v115 = v113
	goto L6
L19:
	;
	v119 = v50 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v119 < v120 {
		v50 = v119
		v51 = v116
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L5
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v5546 = v125
	goto L1
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v116
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[0]))
	if v128 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v129 = m.T0[v128].(func(*base.Module, int32, int32, int32) int32)(m, l0, v44, v116)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[1])))
	if v132 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v5546 = v129
	goto L1
L28:
	;
	v4046 = m.G0
	v4048 = v4046 - int32(16)
	m.G0 = v4048
	v4054 = F_palloc0(m, v44<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		goto L13
	} else {
		goto L458
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[2]))
	if v44 < v136 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v138 = int32(0)
	v140 = m.G0
	v142 = v140 - int32(48)
	m.G0 = v142
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	if v145 < v138 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[4]))
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[5]))
	if int32(0) < v152 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v351 = v145
	goto L33
L33:
	;
	F_SetPlannerInfoExtensionState(m, l0, v351, v142+int32(24))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L59
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3])) = v348
	v351 = v348
	goto L33
L35:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[6]))
	if v278 <= v241 {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v155 = v138
	goto L39
L37:
	;
	goto L38
L38:
	;
	if v150 != 0 {
		v241 = v152
		v242 = v150
		goto L35
	} else {
		goto L50
	}
L39:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v150+v155<<(uint(int32(2))%32))))
	v196 = int32(_a_F_make_rel_from_joinlist_3)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[7])))
	if base.B2i32(v199 == int32(0))|base.B2i32(v199 != v202) != 0 {
		v220 = v199
		v221 = v202
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v220-v221 == int32(0) {
		v348 = v155
		goto L34
	} else {
		goto L48
	}
L42:
	;
	goto L41
L43:
	;
	v205 = v195
	v206 = v196
	goto L44
L44:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v210 == int32(0) {
		v220 = v210
		v221 = v209
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v220 = v210
	v221 = v209
	goto L42
L46:
	;
	v213 = int32(1)
	if v210 == v209 {
		v205 = v205 + v213
		v206 = v206 + v213
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v226 = v155 + int32(1)
	if v226 != v152 {
		v155 = v226
		goto L39
	} else {
		goto L49
	}
L49:
	;
	v241 = v152
	v242 = v150
	goto L35
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[6])) = int32(16)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[8]))
	v235 = F_MemoryContextAlloc(m, v233, int32(64))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[4])) = v235
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[5]))
	v241 = v239
	v242 = v235
	goto L35
L52:
	;
	v280 = int32(1)
	v283 = v241 + v280
	if v283&v241 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v300 = v241
	v301 = v242
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301+v300<<(uint(int32(2))%32)))) = int32(_a_F_make_rel_from_joinlist_3)
	*(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[5])) = v300 + int32(1)
	v348 = v300
	goto L34
L55:
	;
	v288 = v280 << (uint(int32(32)-base.I32_clz(v283)) % 32)
	goto L57
L56:
	;
	v288 = v283
	goto L57
L57:
	;
	v291 = F_repalloc(m, v242, v288<<(uint(int32(2))%32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[6])) = v288
	*(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[4])) = v291
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[5]))
	v300 = v298
	v301 = v291
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = v116
	v392 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+339)) = uint8(v392)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v394+v396<<(uint(int32(2))%32))))
	v402 = v400 + int32(8)
	v404 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[9]))
	v409 = base.I64_trunc_sat_f64_s(base.F64_mul(v404, float64(4.503599627370495e+15)))
	v411 = v409 + int64(4354685564936845354)
	v412 = int64(30)
	v415 = int64(-4658895280553007687)
	v416 = (int64(base.Ui64(v411)>>(uint(v412)%64)) ^ v411) * v415
	v417 = int64(27)
	v420 = int64(-7723592293110705685)
	v421 = (int64(base.Ui64(v416)>>(uint(v417)%64)) ^ v416) * v420
	v422 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v402)+8)) = int64(base.Ui64(v421)>>(uint(v422)%64)) ^ v421
	v427 = v409 - int64(7046029254386353131)
	v432 = (int64(base.Ui64(v427)>>(uint(v412)%64)) ^ v427) * v415
	v437 = (int64(base.Ui64(v432)>>(uint(v417)%64)) ^ v432) * v420
	*(*int64)(unsafe.Add(mBase, uint32(v402))) = int64(base.Ui64(v437)>>(uint(v422)%64)) ^ v437
	goto L60
L60:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[10]))
	if int32(1) < v443 {
		v625 = v443
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[11]))
	v637 = F_palloc(m, int32(12))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L13
	} else {
		goto L103
	}
L62:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[12]))
	v450 = base.F64_add(base.F64_convert_i32_s(v44), float64(1))
	goto L64
L63:
	;
	v616 = v447 * int32(50)
	if base.F64_gt(v614, base.F64_convert_i32_s(v616)) != 0 {
		v625 = v616
		goto L61
	} else {
		goto L101
	}
L64:
	;
	v456 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v450))>>(uint(int64(52))%64))) & int32(2047)
	v461 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	goto L65
L65:
	;
	goto L66
L66:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v461) <= base.Ui32(v456-v461) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v471 = base.F64_add(v450, float64(1))
	if base.Ui32(v456) < base.Ui32(v461) {
		v614 = v471
		goto L63
	} else {
		goto L70
	}
L68:
	;
	v505 = v456
	goto L69
L69:
	;
	v509 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[13]))
	v510 = base.F64_add(v450, v509)
	v512 = base.F64_sub(v450, base.F64_sub(v510, v509))
	v513 = base.F64_mul(v512, v512)
	v516 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[14]))
	v519 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[15]))
	v523 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[16]))
	v526 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[17]))
	v530 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[18]))
	v532 = base.I64_reinterpret_f64(v510)
	v537 = base.I32_wrap_i64(v532) << (uint(int32(4)) % 32) & int32(2032)
	v538 = *(*float64)(unsafe.Add(mBase, uint32(v537)+uint32(_c_F_make_rel_from_joinlist[19])))
	v541 = base.F64_add(base.F64_mul(base.F64_mul(v513, v513), base.F64_add(base.F64_mul(v512, v516), v519)), base.F64_add(base.F64_mul(v513, base.F64_add(base.F64_mul(v512, v523), v526)), base.F64_add(base.F64_mul(v512, v530), v538)))
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v537)+uint32(_c_F_make_rel_from_joinlist[20])))
	v545 = v542 + v532<<(uint(int64(45))%64)
	if v505 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L70:
	;
	v473 = base.I64_reinterpret_f64(v450)
	goto L72
L71:
	;
	if base.Ui64(v473<<(uint(int64(1))%64)) <= base.Ui64(int64(-9143996093422370816)) {
		goto L83
	} else {
		goto L84
	}
L72:
	;
	if base.Ui32(v456) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	if v473 == int64(-4503599627370496) {
		v614 = float64(0)
		goto L63
	} else {
		goto L74
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(math.Float64frombits(uint64(0x7ff0000000000000))))>>(uint(int64(52))%64)))) <= base.Ui32(v456) {
		v614 = v471
		goto L63
	} else {
		goto L76
	}
L76:
	;
	if int64(0) <= v473 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v493 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
	mBase = m.M
	goto L80
L78:
	;
	goto L79
L79:
	;
	if base.Ui64(v473) < base.Ui64(int64(-4570929321408987136)) {
		goto L71
	} else {
		goto L81
	}
L80:
	;
	v614 = v493
	goto L63
L81:
	;
	v498 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
	mBase = m.M
	goto L82
L82:
	;
	v614 = v498
	goto L63
L83:
	;
	v504 = v456
	goto L85
L84:
	;
	v504 = int32(0)
	goto L85
L85:
	;
	v505 = v504
	goto L69
L86:
	;
	if v532&int64(2147483648) == int64(0) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L88
L88:
	;
	v603 = base.F64_reinterpret_i64(v545)
	v614 = base.F64_add(base.F64_mul(v603, v541), v603)
	goto L63
L89:
	;
	v614 = v602
	goto L63
L90:
	;
	v554 = base.F64_reinterpret_i64(v545 - int64(4503599627370496))
	v556 = base.F64_add(base.F64_mul(v554, v541), v554)
	v602 = base.F64_add(v556, v556)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v560 = base.F64_reinterpret_i64(v545 + int64(4602678819172646912))
	v561 = base.F64_mul(v560, v541)
	v562 = base.F64_add(v561, v560)
	if base.F64_lt(v562, float64(1)) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v566 = m.G0
	v568 = v566 - int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v568)+8)) = int64(4503599627370496)
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v568)+8))
	goto L96
L94:
	;
	v595 = v562
	goto L95
L95:
	;
	v602 = base.F64_mul(v595, float64(2.2250738585072014e-308))
	goto L89
L96:
	;
	v574 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v574-int32(16))+8)) = base.F64_mul(v571, float64(2.2250738585072014e-308))
	goto L97
L97:
	;
	v578 = float64(0)
	v579 = float64(1)
	v580 = base.F64_add(v562, v579)
	v589 = base.F64_add(base.F64_add(v580, base.F64_add(base.F64_add(v561, base.F64_sub(v560, v562)), base.F64_add(v562, base.F64_sub(v579, v580)))), float64(-1))
	if base.F64_eq(v589, v578) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v592 = v578
	goto L100
L99:
	;
	v592 = v589
	goto L100
L100:
	;
	v595 = v592
	goto L95
L101:
	;
	v620 = v447 * int32(10)
	if base.F64_lt(v614, base.F64_convert_i32_s(v620)) != 0 {
		v625 = v620
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v625 = base.I32_trunc_sat_f64_s(base.F64_ceil(v614))
	goto L61
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637)+8)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v637)+4)) = v625
	v642 = F_palloc_mul(m, int32(24), v625)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637))) = v642
	if int32(0) < v625 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v651 = int32(0)
	goto L108
L106:
	;
	goto L107
L107:
	;
	v734 = int32(0)
	v735 = m.G0
	v737 = v735 - int32(16)
	m.G0 = v737
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v734 < v739 {
		goto L114
	} else {
		goto L115
	}
L108:
	;
	v691 = F_palloc_mul(m, int32(4), v44+int32(1))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L13
	} else {
		goto L110
	}
L109:
	;
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642+v651*int32(24)))) = v691
	v695 = v651 + int32(1)
	if v695 != v625 {
		v651 = v695
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	F_pg_qsort(m, v979, v980, int32(24), int32(867))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L13
	} else {
		goto L142
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L13
	} else {
		goto L139
	}
L114:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v746 = v138
	v747 = v734
	goto L117
L115:
	;
	goto L116
L116:
	;
	m.G0 = v737 + int32(16)
	goto L112
L117:
	;
	v781 = v747 * int32(24)
	v782 = v742 + v781
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	if v785 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	F_geqo_eval(m, v737, l0, v897, v898)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L13
	} else {
		goto L129
	}
L120:
	;
	v788 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v783))) = v788
	if v785 == v788 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v798 = int32(1)
	goto L122
L122:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v830+v832<<(uint(int32(2))%32))))
	v841 = F_pg_prng_uint64_range(m, v836+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v798))
	mBase = m.M
	v842 = base.I32_wrap_i64(v841)
	goto L124
L123:
	;
	goto L119
L124:
	;
	if v842 != v798 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v844 = int32(2)
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v783+v842<<(uint(v844)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v783+v798<<(uint(v844)%32)))) = v850
	goto L127
L126:
	;
	goto L127
L127:
	;
	v856 = v798 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v783+v842<<(uint(int32(2))%32)))) = v856
	if v856 != v785 {
		v798 = v856
		goto L122
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	v901 = v896 + v781
	v902 = *(*int64)(unsafe.Add(mBase, uint32(v737)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+16)) = v902
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v737)))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+8)) = v904
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v907 = v906 + v781
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+8))
	if v908 == int32(2147483647) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v923 < v924 {
		v746 = v922
		v747 = v923
		goto L117
	} else {
		goto L138
	}
L131:
	;
	v919 = v746 + int32(1)
	if v747 != 0 {
		v922 = v919
		v923 = v747
		goto L130
	} else {
		goto L136
	}
L132:
	;
	v911 = *(*float64)(unsafe.Add(mBase, uint32(v907)+16))
	if base.F64_lt(v911, float64(1.7976931348623157e+308)) == int32(0) {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v922 = v746
	v923 = v747 + int32(1)
	goto L130
L135:
	;
	goto L134
L136:
	;
	if int32(_a_F_make_rel_from_joinlist_4) <= v919 {
		goto L113
	} else {
		goto L137
	}
L137:
	;
	v922 = v919
	v923 = v747
	goto L130
L138:
	;
	goto L118
L139:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_5), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_6), int32(115), int32(_a_F_make_rel_from_joinlist_7))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L13
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
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v986 = F_alloc_chromo(m, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v989 = F_alloc_chromo(m, v988)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v995 = F_palloc_mul(m, int32(24), v992+int32(1))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	if int32(0) < v635 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v999 = v635
	goto L148
L147:
	;
	v999 = v625
	goto L148
L148:
	;
	if int32(0) < v999 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v1003 = v986 + int32(8)
	v1016 = int32(0)
	goto L152
L150:
	;
	goto L151
L151:
	;
	v3917 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v3918 = *(*int32)(unsafe.Add(mBase, uint32(v3917)))
	v3919 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v3920 = F_gimme_tree(m, l0, v3918, v3919)
	mBase = m.M
	v3921 = m.ExcPending
	if v3921 != 0 {
		goto L13
	} else {
		goto L438
	}
L152:
	;
	v1043 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[21]))
	v1045 = base.F64_add(v1043, float64(-1))
	v1047 = base.F64_mul(v1045, float64(4))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v1049 = base.F64_convert_i32_s(v1048)
	goto L154
L153:
	;
	goto L151
L154:
	;
	v1087 = base.F64_mul(v1043, v1043)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1088+v1090<<(uint(int32(2))%32))))
	v1097 = F_pg_prng_double(m, v1094+int32(8))
	mBase = m.M
	goto L156
L155:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v1114 = base.F64_convert_i32_s(v1113)
	goto L161
L156:
	;
	v1099 = base.F64_sub(v1087, base.F64_mul(v1047, v1097))
	if base.F64_gt(v1099, float64(0)) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1103 = base.F64_sqrt(v1099)
	goto L159
L158:
	;
	v1103 = v1099
	goto L159
L159:
	;
	v1108 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v1043, v1103), v1049), float64(0.5)), v1045)
	if base.F64_lt(v1108, float64(0))|base.F64_ge(v1108, v1049) != 0 {
		goto L154
	} else {
		goto L160
	}
L160:
	;
	goto L155
L161:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v1154 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1154<<(uint(int32(2))%32))))
	v1161 = F_pg_prng_double(m, v1158+int32(8))
	mBase = m.M
	goto L163
L162:
	;
	v1177 = base.I32_trunc_sat_f64_s(v1172)
	v1178 = base.I32_trunc_sat_f64_s(v1108)
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if base.B2i32(v1177 != v1178)|base.B2i32(v1180 < int32(2)) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L163:
	;
	v1163 = base.F64_sub(v1087, base.F64_mul(v1047, v1161))
	if base.F64_gt(v1163, float64(0)) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1167 = base.F64_sqrt(v1163)
	goto L166
L165:
	;
	v1167 = v1163
	goto L166
L166:
	;
	v1172 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v1043, v1167), v1114), float64(0.5)), v1045)
	if base.F64_lt(v1172, float64(0))|base.F64_le(v1114, v1172) != 0 {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	goto L162
L168:
	;
	goto L171
L169:
	;
	v1289 = v1177
	goto L170
L170:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v1329 = v1326 + v1178*int32(24)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v1331 = int32(0)
	if v1330 <= v1331 {
		goto L182
	} else {
		goto L183
	}
L171:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v1224 = base.F64_convert_i32_s(v1223)
	goto L173
L172:
	;
	v1289 = v1287
	goto L170
L173:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v1264 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1262+v1264<<(uint(int32(2))%32))))
	v1271 = F_pg_prng_double(m, v1268+int32(8))
	mBase = m.M
	goto L175
L174:
	;
	v1287 = base.I32_trunc_sat_f64_s(v1282)
	if v1287 == v1178 {
		goto L171
	} else {
		goto L180
	}
L175:
	;
	v1273 = base.F64_sub(v1087, base.F64_mul(v1047, v1271))
	if base.F64_gt(v1273, float64(0)) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1277 = base.F64_sqrt(v1273)
	goto L178
L177:
	;
	v1277 = v1273
	goto L178
L178:
	;
	v1282 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v1043, v1277), v1224), float64(0.5)), v1045)
	if base.F64_lt(v1282, float64(0))|base.F64_ge(v1282, v1224) != 0 {
		goto L173
	} else {
		goto L179
	}
L179:
	;
	goto L174
L180:
	;
	goto L172
L181:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v1440 = v1437 + v1289*int32(24)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v1442 = int32(0)
	if v1441 <= v1442 {
		goto L195
	} else {
		goto L196
	}
L182:
	;
	v1433 = *(*int64)(unsafe.Add(mBase, uint32(v1329)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v986)+16)) = v1433
	v1435 = *(*int64)(unsafe.Add(mBase, uint32(v1329)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v986)+8)) = v1435
	goto L181
L183:
	;
	v1340 = v1330 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1330) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1348 = v1331
	v1352 = v1331
	goto L187
L185:
	;
	v1396 = v1331
	goto L186
L186:
	;
	v1405 = v1396
	v1410 = v1331
	goto L191
L187:
	;
	v1355 = v1348 << (uint(int32(2)) % 32)
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1358+v1355)))
	*(*int32)(unsafe.Add(mBase, uint32(v1355+v1356))) = v1360
	v1362 = int32(4)
	v1363 = v1355 | v1362
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1366+v1363)))
	*(*int32)(unsafe.Add(mBase, uint32(v1363+v1364))) = v1368
	v1371 = v1355 | int32(8)
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1374+v1371)))
	*(*int32)(unsafe.Add(mBase, uint32(v1371+v1372))) = v1376
	v1379 = v1355 | int32(12)
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1382+v1379)))
	*(*int32)(unsafe.Add(mBase, uint32(v1379+v1380))) = v1384
	v1387 = v1348 + v1362
	v1389 = v1352 + v1362
	if v1389 != v1330&int32(2147483644) {
		v1348 = v1387
		v1352 = v1389
		goto L187
	} else {
		goto L189
	}
L188:
	;
	if v1340 == int32(0) {
		goto L182
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	v1396 = v1387
	goto L186
L191:
	;
	v1412 = v1405 << (uint(int32(2)) % 32)
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1412)))
	*(*int32)(unsafe.Add(mBase, uint32(v1412+v1413))) = v1417
	v1419 = int32(1)
	v1422 = v1410 + v1419
	if v1422 != v1340 {
		v1405 = v1405 + v1419
		v1410 = v1422
		goto L191
	} else {
		goto L193
	}
L192:
	;
	goto L182
L193:
	;
	goto L192
L194:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1550 = int32(0)
	v1551 = int32(1)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	if base.B2i32(v1552 <= v1550) == v1550 {
		goto L207
	} else {
		goto L208
	}
L195:
	;
	v1544 = *(*int64)(unsafe.Add(mBase, uint32(v1440)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v989)+16)) = v1544
	v1546 = *(*int64)(unsafe.Add(mBase, uint32(v1440)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v989)+8)) = v1546
	goto L194
L196:
	;
	v1451 = v1441 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1441) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1459 = v1442
	v1463 = v1442
	goto L200
L198:
	;
	v1507 = v1442
	goto L199
L199:
	;
	v1516 = v1507
	v1521 = v1442
	goto L204
L200:
	;
	v1466 = v1459 << (uint(int32(2)) % 32)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1469+v1466)))
	*(*int32)(unsafe.Add(mBase, uint32(v1466+v1467))) = v1471
	v1473 = int32(4)
	v1474 = v1466 | v1473
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1477+v1474)))
	*(*int32)(unsafe.Add(mBase, uint32(v1474+v1475))) = v1479
	v1482 = v1466 | int32(8)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1485+v1482)))
	*(*int32)(unsafe.Add(mBase, uint32(v1482+v1483))) = v1487
	v1490 = v1466 | int32(12)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1493+v1490)))
	*(*int32)(unsafe.Add(mBase, uint32(v1490+v1491))) = v1495
	v1498 = v1459 + v1473
	v1500 = v1463 + v1473
	if v1500 != v1441&int32(2147483644) {
		v1459 = v1498
		v1463 = v1500
		goto L200
	} else {
		goto L202
	}
L201:
	;
	if v1451 == int32(0) {
		goto L195
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	v1507 = v1498
	goto L199
L204:
	;
	v1523 = v1516 << (uint(int32(2)) % 32)
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1526+v1523)))
	*(*int32)(unsafe.Add(mBase, uint32(v1523+v1524))) = v1528
	v1530 = int32(1)
	v1533 = v1521 + v1530
	if v1533 != v1451 {
		v1516 = v1516 + v1530
		v1521 = v1533
		goto L204
	} else {
		goto L206
	}
L205:
	;
	goto L195
L206:
	;
	goto L205
L207:
	;
	v1557 = int32(2)
	v1559 = v1552 + int32(1)
	if v1559 <= v1557 {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	goto L209
L209:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v2437 = m.G0
	v2439 = v2437 - int32(32)
	m.G0 = v2439
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v2444 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2442+v2444<<(uint(int32(2))%32))))
	v2453 = F_pg_prng_uint64_range(m, v2448+int32(8), base.I64_extend_i32_s(int32(1)), base.I64_extend_i32_s(v2435))
	mBase = m.M
	goto L266
L210:
	;
	v1562 = v1557
	goto L212
L211:
	;
	v1562 = v1559
	goto L212
L212:
	;
	v1564 = v1562 - int32(1)
	v1566 = v1564 & int32(3)
	if int32(5) <= v1559 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1750 = int32(0)
	goto L224
L214:
	;
	v1574 = int32(0)
	v1575 = v1551
	goto L217
L215:
	;
	v1630 = v1551
	goto L216
L216:
	;
	v1665 = int32(0)
	v1668 = v1630
	goto L221
L217:
	;
	v1611 = v995 + v1575*int32(24)
	v1612 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+88)) = v1612
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+64)) = v1612
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+40)) = v1612
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+16)) = v1612
	v1620 = int32(4)
	v1621 = v1575 + v1620
	v1623 = v1574 + v1620
	if v1623 != v1564&int32(-4) {
		v1574 = v1623
		v1575 = v1621
		goto L217
	} else {
		goto L219
	}
L218:
	;
	if v1566 == int32(0) {
		goto L213
	} else {
		goto L220
	}
L219:
	;
	goto L218
L220:
	;
	v1630 = v1621
	goto L216
L221:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v995+v1668*int32(24))+16)) = int64(0)
	v1707 = int32(1)
	v1710 = v1665 + v1707
	if v1710 != v1566 {
		v1665 = v1710
		v1668 = v1668 + v1707
		goto L221
	} else {
		goto L223
	}
L222:
	;
	goto L213
L223:
	;
	goto L222
L224:
	;
	v1788 = v1750 + int32(1)
	if v1788 != v1552 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L209
L226:
	;
	v1791 = v1788
	goto L228
L227:
	;
	v1791 = int32(0)
	goto L228
L228:
	;
	v1792 = int32(2)
	v1793 = v1791 << (uint(v1792) % 32)
	v1794 = v1548 + v1793
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1794)))
	v1796 = int32(0)
	v1798 = v1750 << (uint(v1792) % 32)
	v1799 = v1548 + v1798
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1799)))
	v1803 = v995 + v1800*int32(24)
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+16))
	if v1804 <= v1796 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1948 = int32(0)
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1799)))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v1794)))
	v1953 = v995 + v1950*int32(24)
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+16))
	if v1954 <= v1948 {
		goto L239
	} else {
		goto L240
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1803+v1804<<(uint(int32(2))%32)))) = v1795
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+16))
	v1902 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1803)+16)) = v1901 + v1902
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1803)+20)) = v1905 + v1902
	goto L229
L231:
	;
	v1810 = v1796
	goto L232
L232:
	;
	v1846 = v1803 + v1810<<(uint(int32(2))%32)
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)))
	v1849 = v1847 >> (uint(int32(31)) % 32)
	if v1795 != v1847^v1849-v1849 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1846))) = int32(0) - v1795
	goto L229
L234:
	;
	v1854 = v1810 + int32(1)
	if v1804 != v1854 {
		v1810 = v1854
		goto L232
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	goto L233
L237:
	;
	goto L230
L238:
	;
	v2095 = int32(0)
	v2096 = v1549 + v1793
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2096)))
	v2098 = v1549 + v1798
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2098)))
	v2102 = v995 + v2099*int32(24)
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+16))
	if v2103 <= v2095 {
		goto L248
	} else {
		goto L249
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1953+v1954<<(uint(int32(2))%32)))) = v1949
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+16))
	v2051 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1953)+16)) = v2050 + v2051
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1953)+20)) = v2054 + v2051
	goto L238
L240:
	;
	v1960 = v1948
	goto L241
L241:
	;
	v1996 = v1953 + v1960<<(uint(int32(2))%32)
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1996)))
	v1999 = v1997 >> (uint(int32(31)) % 32)
	if v1949 != v1997^v1999-v1999 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1996))) = int32(0) - v1949
	goto L238
L243:
	;
	v2004 = v1960 + int32(1)
	if v1954 != v2004 {
		v1960 = v2004
		goto L241
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	goto L242
L246:
	;
	goto L239
L247:
	;
	v2248 = int32(0)
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v2098)))
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2096)))
	v2253 = v995 + v2250*int32(24)
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v2253)+16))
	if v2254 <= v2248 {
		goto L257
	} else {
		goto L258
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102+v2103<<(uint(int32(2))%32)))) = v2097
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+16))
	v2201 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2102)+16)) = v2200 + v2201
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2102)+20)) = v2204 + v2201
	goto L247
L249:
	;
	v2109 = v2095
	goto L250
L250:
	;
	v2145 = v2102 + v2109<<(uint(int32(2))%32)
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2145)))
	v2148 = v2146 >> (uint(int32(31)) % 32)
	if v2097 != v2146^v2148-v2148 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2145))) = int32(0) - v2097
	goto L247
L252:
	;
	v2153 = v2109 + int32(1)
	if v2103 != v2153 {
		v2109 = v2153
		goto L250
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	goto L251
L255:
	;
	goto L248
L256:
	;
	if v1788 != v1552 {
		v1750 = v1788
		goto L224
	} else {
		goto L265
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2253+v2254<<(uint(int32(2))%32)))) = v2249
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(v2253)+16))
	v2351 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2253)+16)) = v2350 + v2351
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2253)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2253)+20)) = v2354 + v2351
	goto L256
L258:
	;
	v2260 = v2248
	goto L259
L259:
	;
	v2296 = v2253 + v2260<<(uint(int32(2))%32)
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)))
	v2299 = v2297 >> (uint(int32(31)) % 32)
	if v2249 != v2297^v2299-v2299 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2296))) = int32(0) - v2249
	goto L256
L261:
	;
	v2304 = v2260 + int32(1)
	if v2254 != v2304 {
		v2260 = v2304
		goto L259
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	goto L260
L264:
	;
	goto L257
L265:
	;
	goto L225
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2434))) = base.I32_wrap_i64(v2453)
	if int32(2) <= v2435 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v2458 = int32(2)
	v2460 = v2435 + int32(1)
	if v2460 <= v2458 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	m.G0 = v2439 + int32(32)
	v3521 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v3522 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	F_geqo_eval(m, v142+int32(8), l0, v3521, v3522)
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L13
	} else {
		goto L384
	}
L270:
	;
	v2463 = v2458
	goto L272
L271:
	;
	v2463 = v2460
	goto L272
L272:
	;
	v2464 = int32(1)
	v2465 = v2463 - v2464
	v2485 = v2464
	goto L273
L273:
	;
	v2512 = v2434 + v2485<<(uint(int32(2))%32)
	v2514 = v2512 - int32(4)
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2514)))
	v2518 = v995 + v2515*int32(24)
	v2519 = *(*int64)(unsafe.Add(mBase, uint32(v2518)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2439)+24)) = v2519
	v2521 = *(*int64)(unsafe.Add(mBase, uint32(v2518)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2439)+16)) = v2521
	v2523 = *(*int64)(unsafe.Add(mBase, uint32(v2518)))
	*(*int64)(unsafe.Add(mBase, uint32(v2439)+8)) = v2523
	v2525 = int32(0)
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+28))
	if v2525 < v2526 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	goto L269
L275:
	;
	v2530 = v2525
	goto L278
L276:
	;
	v2690 = v2515
	goto L277
L277:
	;
	v2722 = v995 + v2690*int32(24)
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(v2722)+20))
	if int32(0) < v2723 {
		goto L293
	} else {
		goto L294
	}
L278:
	;
	v2566 = int32(0)
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v2439+int32(8)+v2530<<(uint(int32(2))%32))))
	v2574 = v2572 >> (uint(int32(31)) % 32)
	v2579 = v995 + (v2572^v2574-v2574)*int32(24)
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v2579)+20))
	if v2580 <= v2566 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v2514)))
	v2690 = v2682
	goto L277
L280:
	;
	v2680 = v2530 + int32(1)
	if v2680 != v2526 {
		v2530 = v2680
		goto L278
	} else {
		goto L288
	}
L281:
	;
	v2585 = v2566
	goto L282
L282:
	;
	v2622 = v2579 + v2585<<(uint(int32(2))%32)
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v2622)))
	v2625 = v2623 >> (uint(int32(31)) % 32)
	if v2515 != v2623^v2625-v2625 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2579)+20)) = v2580 - int32(1)
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(v2579+v2580<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v2622))) = v2640
	goto L280
L284:
	;
	v2630 = v2585 + int32(1)
	if v2580 != v2630 {
		v2585 = v2630
		goto L282
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	goto L283
L287:
	;
	goto L280
L288:
	;
	goto L279
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2512))) = v3434
	v3470 = *(*int32)(unsafe.Add(mBase, uint32(v2514)))
	*(*int32)(unsafe.Add(mBase, uint32(v995+v3470*int32(24))+20)) = int32(-1)
	v3477 = v2485 + int32(1)
	if v3477 != v2435 {
		v2485 = v3477
		goto L273
	} else {
		goto L383
	}
L290:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L13
	} else {
		goto L380
	}
L291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L13
	} else {
		goto L377
	}
L292:
	;
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v3358 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v3362 = *(*int32)(unsafe.Add(mBase, uint32(v3356+v3358<<(uint(int32(2))%32))))
	v3367 = F_pg_prng_uint64_range(m, v3362+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(int32(-2)))
	mBase = m.M
	goto L376
L293:
	;
	v2726 = *(*int64)(unsafe.Add(mBase, uint32(v2722)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2439)+24)) = v2726
	v2728 = *(*int64)(unsafe.Add(mBase, uint32(v2722)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2439)+16)) = v2728
	v2730 = *(*int64)(unsafe.Add(mBase, uint32(v2722)))
	*(*int64)(unsafe.Add(mBase, uint32(v2439)+8)) = v2730
	v2732 = int32(0)
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+28))
	if v2733 <= v2732 {
		goto L292
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v2872 = int32(0)
	v2874 = int32(1)
	if base.B2i32(v2460 < int32(3)) == v2872 {
		goto L317
	} else {
		goto L318
	}
L296:
	;
	v2740 = v2732
	v2741 = int32(5)
	v2745 = int32(-1)
	goto L297
L297:
	;
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2439+int32(8)+v2740<<(uint(int32(2))%32))))
	if v2780 < int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v2800 = int32(0)
	v2804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v2806 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2804+v2806<<(uint(int32(2))%32))))
	v2815 = F_pg_prng_uint64_range(m, v2810+int32(8), base.I64_extend_i32_s(v2800), base.I64_extend_i32_s(v2796-int32(1)))
	mBase = m.M
	goto L308
L299:
	;
	v3434 = int32(0) - v2780
	goto L289
L300:
	;
	goto L301
L301:
	;
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v995+v2780*int32(24))+20))
	if v2788 < v2741 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	v2798 = v2740 + int32(1)
	if v2798 != v2733 {
		v2740 = v2798
		v2741 = v2795
		v2745 = v2796
		goto L297
	} else {
		goto L307
	}
L303:
	;
	v2795 = v2788
	v2796 = int32(1)
	goto L302
L304:
	;
	goto L305
L305:
	;
	if v2745 == int32(-1) {
		goto L291
	} else {
		goto L306
	}
L306:
	;
	v2795 = v2741
	v2796 = v2745 + base.B2i32(v2788 == v2741)
	goto L302
L307:
	;
	goto L298
L308:
	;
	v2817 = v2800
	v2824 = v2796
	goto L309
L309:
	;
	v2859 = *(*int32)(unsafe.Add(mBase, uint32(v2439+int32(8)+v2817<<(uint(int32(2))%32))))
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v995+v2859*int32(24))+20))
	if v2795 == v2863 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	goto L290
L311:
	;
	v2866 = v2824 - int32(1)
	if v2866 == base.I32_wrap_i64(v2815) {
		v3434 = v2859
		goto L289
	} else {
		goto L314
	}
L312:
	;
	v2868 = v2824
	goto L313
L313:
	;
	v2870 = v2817 + int32(1)
	if v2870 != v2733 {
		v2817 = v2870
		v2824 = v2868
		goto L309
	} else {
		goto L315
	}
L314:
	;
	v2868 = v2866
	goto L313
L315:
	;
	goto L310
L316:
	;
	if v3010 != 0 {
		goto L338
	} else {
		goto L339
	}
L317:
	;
	v2879 = v2872
	v2880 = v2872
	v2881 = v2874
	v2882 = v2872
	goto L320
L318:
	;
	v2958 = v2872
	v2960 = v2872
	v2963 = v2874
	goto L319
L319:
	;
	if v2963 == v2690 {
		goto L330
	} else {
		goto L331
	}
L320:
	;
	if v2881 == v2690 {
		v2929 = v2879
		v2930 = v2882
		goto L322
	} else {
		goto L323
	}
L321:
	;
	if v2465&v2464 == int32(0) {
		v3007 = v2947
		v3010 = v2948
		goto L316
	} else {
		goto L329
	}
L322:
	;
	v2933 = v2881 + int32(1)
	if v2933 == v2690 {
		v2947 = v2929
		v2948 = v2930
		goto L325
	} else {
		goto L326
	}
L323:
	;
	v2919 = v995 + v2881*int32(24)
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2919)+20))
	if v2920 == int32(-1) {
		v2929 = v2879
		v2930 = v2882
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2919)+16))
	v2929 = v2879 + int32(1)
	v2930 = v2882 + base.B2i32(v2925 == int32(4))
	goto L322
L325:
	;
	v2950 = int32(2)
	v2951 = v2881 + v2950
	v2953 = v2880 + v2950
	if v2953 != v2465&int32(-2) {
		v2879 = v2947
		v2880 = v2953
		v2881 = v2951
		v2882 = v2948
		goto L320
	} else {
		goto L328
	}
L326:
	;
	v2937 = v995 + v2933*int32(24)
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2937)+20))
	if v2938 == int32(-1) {
		v2947 = v2929
		v2948 = v2930
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(v2937)+16))
	v2947 = v2929 + int32(1)
	v2948 = v2930 + base.B2i32(v2943 == int32(4))
	goto L325
L328:
	;
	goto L321
L329:
	;
	v2958 = v2947
	v2960 = v2948
	v2963 = v2951
	goto L319
L330:
	;
	v3007 = v2958
	v3010 = v2960
	goto L316
L331:
	;
	goto L332
L332:
	;
	v2997 = v995 + v2963*int32(24)
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(v2997)+20))
	if v2998 == int32(-1) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v3007 = v2958
	v3010 = v2960
	goto L316
L334:
	;
	goto L335
L335:
	;
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v2997)+16))
	v3007 = v2958 + int32(1)
	v3010 = v2960 + base.B2i32(v3003 == int32(4))
	goto L316
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L13
	} else {
		goto L373
	}
L337:
	;
	F_errmsg_internal(m, v3296, int32(0))
	mBase = m.M
	v3299 = m.ExcPending
	if v3299 != 0 {
		goto L13
	} else {
		goto L371
	}
L338:
	;
	v3046 = int32(1)
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v3052 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v3056 = *(*int32)(unsafe.Add(mBase, uint32(v3050+v3052<<(uint(int32(2))%32))))
	v3061 = F_pg_prng_uint64_range(m, v3056+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v3010-v3046))
	mBase = m.M
	goto L341
L339:
	;
	goto L340
L340:
	;
	if v3007 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L341:
	;
	v3065 = v3046
	v3066 = v3010
	goto L342
L342:
	;
	if v3065 == v2690 {
		v3114 = v3066
		goto L344
	} else {
		goto L345
	}
L343:
	;
	v3120 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L13
	} else {
		goto L350
	}
L344:
	;
	v3116 = v3065 + int32(1)
	if v3116 <= v2435 {
		v3065 = v3116
		v3066 = v3114
		goto L342
	} else {
		goto L349
	}
L345:
	;
	v3103 = v995 + v3065*int32(24)
	v3104 = *(*int32)(unsafe.Add(mBase, uint32(v3103)+20))
	if v3104 == int32(-1) {
		v3114 = v3066
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v3103)+16))
	if v3107 != int32(4) {
		v3114 = v3066
		goto L344
	} else {
		goto L347
	}
L347:
	;
	v3111 = v3066 - int32(1)
	if base.I32_wrap_i64(v3061) == v3111 {
		v3434 = v3065
		goto L289
	} else {
		goto L348
	}
L348:
	;
	v3114 = v3111
	goto L344
L349:
	;
	goto L343
L350:
	;
	if v3120 == int32(0) {
		goto L336
	} else {
		goto L351
	}
L351:
	;
	v3261 = int32(431)
	v3296 = int32(_a_F_make_rel_from_joinlist_8)
	goto L337
L352:
	;
	v3131 = int32(1)
	goto L355
L353:
	;
	goto L354
L354:
	;
	v3183 = int32(1)
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v3189 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v3193 = *(*int32)(unsafe.Add(mBase, uint32(v3187+v3189<<(uint(int32(2))%32))))
	v3198 = F_pg_prng_uint64_range(m, v3193+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v3007-v3183))
	mBase = m.M
	goto L361
L355:
	;
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v995+v3131*int32(24))+20))
	if int32(0) <= v3169 {
		v3434 = v3131
		goto L289
	} else {
		goto L357
	}
L356:
	;
	v3177 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3178 = m.ExcPending
	if v3178 != 0 {
		goto L13
	} else {
		goto L359
	}
L357:
	;
	v3173 = v3131 + int32(1)
	if v3173 <= v2435 {
		v3131 = v3173
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	if v3177 == int32(0) {
		goto L336
	} else {
		goto L360
	}
L360:
	;
	v3261 = int32(470)
	v3296 = int32(_a_F_make_rel_from_joinlist_9)
	goto L337
L361:
	;
	v3200 = v3007
	v3202 = v3183
	goto L362
L362:
	;
	if v3202 == v2690 {
		v3247 = v3200
		goto L364
	} else {
		goto L365
	}
L363:
	;
	v3253 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		goto L13
	} else {
		goto L369
	}
L364:
	;
	v3249 = v3202 + int32(1)
	if v3249 <= v2435 {
		v3200 = v3247
		v3202 = v3249
		goto L362
	} else {
		goto L368
	}
L365:
	;
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v995+v3202*int32(24))+20))
	if v3241 == int32(-1) {
		v3247 = v3200
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v3245 = v3200 - int32(1)
	if base.I32_wrap_i64(v3198) == v3245 {
		v3434 = v3202
		goto L289
	} else {
		goto L367
	}
L367:
	;
	v3247 = v3245
	goto L364
L368:
	;
	goto L363
L369:
	;
	if v3253 == int32(0) {
		goto L336
	} else {
		goto L370
	}
L370:
	;
	v3261 = int32(452)
	v3296 = int32(_a_F_make_rel_from_joinlist_10)
	goto L337
L371:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_11), v3261, int32(_a_F_make_rel_from_joinlist_12))
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		goto L13
	} else {
		goto L372
	}
L372:
	;
	goto L336
L373:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_13), int32(0))
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		goto L13
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_11), int32(475), int32(_a_F_make_rel_from_joinlist_12))
	mBase = m.M
	v3353 = m.ExcPending
	if v3353 != 0 {
		goto L13
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	goto L290
L377:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_14), int32(0))
	mBase = m.M
	v3376 = m.ExcPending
	if v3376 != 0 {
		goto L13
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_11), int32(345), int32(_a_F_make_rel_from_joinlist_15))
	mBase = m.M
	v3381 = m.ExcPending
	if v3381 != 0 {
		goto L13
	} else {
		goto L379
	}
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L380:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_16), int32(0))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L13
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_11), int32(370), int32(_a_F_make_rel_from_joinlist_15))
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L13
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	goto L274
L384:
	;
	v3525 = *(*int64)(unsafe.Add(mBase, uint32(v142)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1003)+8)) = v3525
	v3527 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1003))) = v3527
	v3530 = m.G0
	v3532 = v3530 - int32(48)
	m.G0 = v3532
	v3534 = *(*float64)(unsafe.Add(mBase, uint32(v986)+16))
	v3535 = *(*int32)(unsafe.Add(mBase, uint32(v986)+8))
	v3536 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v3537 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v3540 = v3536 + v3537*int32(24)
	v3543 = *(*int32)(unsafe.Add(mBase, uint32(v3540-int32(16))))
	if v3535 != v3543 {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	m.G0 = v3532 + int32(48)
	v3878 = v1016 + int32(1)
	if v3878 != v999 {
		v1016 = v3878
		goto L152
	} else {
		goto L437
	}
L386:
	;
	v3555 = base.I32_div_s(v3537, int32(2))
	v3558 = int32(0)
	v3559 = v3537 - int32(1)
	v3562 = v3555
	goto L392
L387:
	;
	if v3535 < v3543 {
		goto L386
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	v3548 = *(*float64)(unsafe.Add(mBase, uint32(v3540-int32(8))))
	if base.F64_le(v3534, v3548) == int32(0) {
		goto L385
	} else {
		goto L391
	}
L390:
	;
	goto L385
L391:
	;
	goto L386
L392:
	;
	v3595 = v3536 + v3558*int32(24)
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v3595)+8))
	if v3596 != v3535 {
		goto L397
	} else {
		goto L398
	}
L393:
	;
	v3649 = v3540 - int32(24)
	v3650 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v3651 = int32(0)
	if v3650 <= v3651 {
		goto L421
	} else {
		goto L422
	}
L394:
	;
	if v3640 == int32(-1) {
		v3558 = v3641
		v3559 = v3642
		v3562 = v3645
		goto L392
	} else {
		goto L419
	}
L395:
	;
	v3640 = v3558
	v3641 = v3558
	v3642 = v3559
	v3645 = v3562
	goto L394
L396:
	;
	v3603 = v3536 + v3562*int32(24)
	v3604 = *(*float64)(unsafe.Add(mBase, uint32(v3603)+16))
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(v3603)+8))
	if base.B2i32(v3535 != v3605)|base.F64_ne(v3534, v3604) == int32(0) {
		goto L402
	} else {
		goto L403
	}
L397:
	;
	if v3596 <= v3535 {
		goto L396
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v3599 = *(*float64)(unsafe.Add(mBase, uint32(v3595)+16))
	if base.F64_le(v3534, v3599) != 0 {
		goto L395
	} else {
		goto L401
	}
L400:
	;
	goto L395
L401:
	;
	goto L396
L402:
	;
	v3640 = v3562
	v3641 = v3558
	v3642 = v3559
	v3645 = v3562
	goto L394
L403:
	;
	goto L404
L404:
	;
	v3613 = v3536 + v3559*int32(24)
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+8))
	if v3614 == v3535 {
		goto L407
	} else {
		goto L408
	}
L405:
	;
	if v3605 != v3535 {
		goto L414
	} else {
		goto L415
	}
L406:
	;
	v3640 = v3559
	v3641 = v3558
	v3642 = v3559
	v3645 = v3562
	goto L394
L407:
	;
	v3619 = *(*float64)(unsafe.Add(mBase, uint32(v3613)+16))
	if base.B2i32(v3559-v3558 < int32(2))|base.F64_eq(v3534, v3619) != 0 {
		goto L406
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	if int32(2) <= v3559-v3558 {
		goto L405
	} else {
		goto L411
	}
L410:
	;
	goto L405
L411:
	;
	goto L406
L412:
	;
	v3638 = base.I32_div_s(v3559-v3562, int32(2))
	v3640 = int32(-1)
	v3641 = v3562
	v3642 = v3559
	v3645 = v3562 + v3638
	goto L394
L413:
	;
	v3633 = base.I32_div_s(v3562-v3558, int32(2))
	v3640 = int32(-1)
	v3641 = v3558
	v3642 = v3562
	v3645 = v3633 + v3558
	goto L394
L414:
	;
	if v3535 < v3605 {
		goto L413
	} else {
		goto L417
	}
L415:
	;
	goto L416
L416:
	;
	if base.F64_lt(v3534, v3604) == int32(0) {
		goto L412
	} else {
		goto L418
	}
L417:
	;
	goto L412
L418:
	;
	goto L413
L419:
	;
	goto L393
L420:
	;
	v3757 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v3759 = int32(24)
	v3761 = v3757 + v3758*v3759
	v3764 = *(*int32)(unsafe.Add(mBase, uint32(v3761-v3759)))
	v3766 = v3761 - int32(16)
	v3767 = *(*int64)(unsafe.Add(mBase, uint32(v3766)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3532)+40)) = v3767
	v3769 = *(*int64)(unsafe.Add(mBase, uint32(v3766)))
	*(*int64)(unsafe.Add(mBase, uint32(v3532)+32)) = v3769
	if v3758 <= v3640 {
		goto L385
	} else {
		goto L433
	}
L421:
	;
	v3753 = *(*int64)(unsafe.Add(mBase, uint32(v986)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3649)+16)) = v3753
	v3755 = *(*int64)(unsafe.Add(mBase, uint32(v986)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3649)+8)) = v3755
	goto L420
L422:
	;
	v3660 = v3650 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3650) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v3668 = v3651
	v3672 = v3651
	goto L426
L424:
	;
	v3716 = v3651
	goto L425
L425:
	;
	v3725 = v3716
	v3730 = v3651
	goto L430
L426:
	;
	v3675 = v3668 << (uint(int32(2)) % 32)
	v3676 = *(*int32)(unsafe.Add(mBase, uint32(v3649)))
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v3680 = *(*int32)(unsafe.Add(mBase, uint32(v3678+v3675)))
	*(*int32)(unsafe.Add(mBase, uint32(v3675+v3676))) = v3680
	v3682 = int32(4)
	v3683 = v3675 | v3682
	v3684 = *(*int32)(unsafe.Add(mBase, uint32(v3649)))
	v3686 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v3688 = *(*int32)(unsafe.Add(mBase, uint32(v3686+v3683)))
	*(*int32)(unsafe.Add(mBase, uint32(v3683+v3684))) = v3688
	v3691 = v3675 | int32(8)
	v3692 = *(*int32)(unsafe.Add(mBase, uint32(v3649)))
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v3696 = *(*int32)(unsafe.Add(mBase, uint32(v3694+v3691)))
	*(*int32)(unsafe.Add(mBase, uint32(v3691+v3692))) = v3696
	v3699 = v3675 | int32(12)
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(v3649)))
	v3702 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v3704 = *(*int32)(unsafe.Add(mBase, uint32(v3702+v3699)))
	*(*int32)(unsafe.Add(mBase, uint32(v3699+v3700))) = v3704
	v3707 = v3668 + v3682
	v3709 = v3672 + v3682
	if v3709 != v3650&int32(2147483644) {
		v3668 = v3707
		v3672 = v3709
		goto L426
	} else {
		goto L428
	}
L427:
	;
	if v3660 == int32(0) {
		goto L421
	} else {
		goto L429
	}
L428:
	;
	goto L427
L429:
	;
	v3716 = v3707
	goto L425
L430:
	;
	v3732 = v3725 << (uint(int32(2)) % 32)
	v3733 = *(*int32)(unsafe.Add(mBase, uint32(v3649)))
	v3735 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	v3737 = *(*int32)(unsafe.Add(mBase, uint32(v3735+v3732)))
	*(*int32)(unsafe.Add(mBase, uint32(v3732+v3733))) = v3737
	v3739 = int32(1)
	v3742 = v3730 + v3739
	if v3742 != v3660 {
		v3725 = v3725 + v3739
		v3730 = v3742
		goto L430
	} else {
		goto L432
	}
L431:
	;
	goto L421
L432:
	;
	goto L431
L433:
	;
	v3773 = v3532 + int32(32)
	v3775 = v3532 + int32(12)
	v3777 = v3640
	v3778 = v3764
	goto L434
L434:
	;
	v3814 = v3777 * int32(24)
	v3815 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v3816 = v3814 + v3815
	v3817 = *(*int32)(unsafe.Add(mBase, uint32(v3816)))
	v3818 = *(*int64)(unsafe.Add(mBase, uint32(v3816)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3775)+8)) = v3818
	v3820 = *(*int64)(unsafe.Add(mBase, uint32(v3816)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3775))) = v3820
	*(*int32)(unsafe.Add(mBase, uint32(v3816))) = v3778
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v3824 = v3823 + v3814
	v3825 = *(*int64)(unsafe.Add(mBase, uint32(v3773)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3824)+16)) = v3825
	v3827 = *(*int64)(unsafe.Add(mBase, uint32(v3773)))
	*(*int64)(unsafe.Add(mBase, uint32(v3824)+8)) = v3827
	v3829 = *(*int64)(unsafe.Add(mBase, uint32(v3775)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3773)+8)) = v3829
	v3831 = *(*int64)(unsafe.Add(mBase, uint32(v3775)))
	*(*int64)(unsafe.Add(mBase, uint32(v3773))) = v3831
	v3834 = v3777 + int32(1)
	v3835 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v3834 < v3835 {
		v3777 = v3834
		v3778 = v3817
		goto L434
	} else {
		goto L436
	}
L435:
	;
	goto L385
L436:
	;
	goto L435
L437:
	;
	goto L153
L438:
	;
	if v3920 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3927 = m.ExcPending
	if v3927 != 0 {
		goto L13
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	F_free_attrmap(m, v986)
	mBase = m.M
	v3938 = m.ExcPending
	if v3938 != 0 {
		goto L13
	} else {
		goto L445
	}
L442:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_5), int32(0))
	mBase = m.M
	v3931 = m.ExcPending
	if v3931 != 0 {
		goto L13
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_17), int32(287), int32(_a_F_make_rel_from_joinlist_3))
	mBase = m.M
	v3936 = m.ExcPending
	if v3936 != 0 {
		goto L13
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	F_free_attrmap(m, v989)
	mBase = m.M
	v3940 = m.ExcPending
	if v3940 != 0 {
		goto L13
	} else {
		goto L446
	}
L446:
	;
	F_pfree(m, v995)
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L13
	} else {
		goto L447
	}
L447:
	;
	v3943 = int32(0)
	v3944 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v3945 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v3943 < v3945 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v3952 = v3943
	goto L451
L449:
	;
	v4033 = v3944
	goto L450
L450:
	;
	F_pfree(m, v4033)
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L13
	} else {
		goto L455
	}
L451:
	;
	v3988 = *(*int32)(unsafe.Add(mBase, uint32(v3944+v3952*int32(24))))
	F_pfree(m, v3988)
	mBase = m.M
	v3990 = m.ExcPending
	if v3990 != 0 {
		goto L13
	} else {
		goto L453
	}
L452:
	;
	v3995 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v4033 = v3995
	goto L450
L453:
	;
	v3992 = v3952 + int32(1)
	v3993 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v3992 < v3993 {
		v3952 = v3992
		goto L451
	} else {
		goto L454
	}
L454:
	;
	goto L452
L455:
	;
	F_pfree(m, v637)
	mBase = m.M
	v4037 = m.ExcPending
	if v4037 != 0 {
		goto L13
	} else {
		goto L456
	}
L456:
	;
	v4039 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	F_SetPlannerInfoExtensionState(m, l0, v4039, int32(0))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L13
	} else {
		goto L457
	}
L457:
	;
	m.G0 = v142 + int32(48)
	v5546 = v3920
	goto L1
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v4054
	*(*int32)(unsafe.Add(mBase, uint32(v4054)+4)) = v116
	if int32(2) <= v44 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v4064 = int32(2)
	goto L462
L460:
	;
	goto L461
L461:
	;
	v5516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v5520 = *(*int32)(unsafe.Add(mBase, uint32(v5516+v44<<(uint(int32(2))%32))))
	if v5520 == int32(0) {
		goto L698
	} else {
		goto L699
	}
L462:
	;
	v4098 = int32(0)
	v4099 = m.G0
	v4101 = v4099 - int32(16)
	m.G0 = v4101
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v4064
	v4104 = int32(2)
	v4105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v4108 = v4105 + v4064<<(uint(v4104)%32)
	v4110 = v4108 - int32(4)
	v4111 = *(*int32)(unsafe.Add(mBase, uint32(v4110)))
	if v4111 == v4098 {
		goto L464
	} else {
		goto L465
	}
L463:
	;
	goto L461
L464:
	;
	v4556 = int32(2)
	v4557 = v4064 - v4556
	if v4556 <= v4557 {
		goto L555
	} else {
		goto L556
	}
L465:
	;
	v4114 = *(*int32)(unsafe.Add(mBase, uint32(v4111)+4))
	if v4114 <= int32(0) {
		goto L464
	} else {
		goto L466
	}
L466:
	;
	v4128 = v4098
	goto L467
L467:
	;
	v4154 = *(*int32)(unsafe.Add(mBase, uint32(v4111)+12))
	v4158 = *(*int32)(unsafe.Add(mBase, uint32(v4154+v4128<<(uint(int32(2))%32))))
	v4159 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+228))
	if v4159 != 0 {
		goto L471
	} else {
		goto L472
	}
L468:
	;
	goto L464
L469:
	;
	v4516 = v4128 + int32(1)
	v4517 = *(*int32)(unsafe.Add(mBase, uint32(v4111)+4))
	if v4516 < v4517 {
		v4128 = v4516
		goto L467
	} else {
		goto L554
	}
L470:
	;
	v4373 = *(*int32)(unsafe.Add(mBase, uint32(v4105)+4))
	if v4373 == int32(0) {
		goto L469
	} else {
		goto L532
	}
L471:
	;
	v4258 = *(*int32)(unsafe.Add(mBase, uint32(v4105)+4))
	if v4258 == int32(0) {
		goto L469
	} else {
		goto L502
	}
L472:
	;
	v4160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4158)+232)))
	if v4160 != 0 {
		goto L471
	} else {
		goto L473
	}
L473:
	;
	v4164 = int32(1)
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+72))
	if v4165 != 0 {
		v4247 = v4164
		goto L475
	} else {
		goto L476
	}
L474:
	;
	if v4255 == int32(0) {
		goto L470
	} else {
		goto L501
	}
L475:
	;
	v4255 = v4247
	goto L474
L476:
	;
	v4166 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+112))
	if v4166 != 0 {
		v4247 = v4164
		goto L475
	} else {
		goto L477
	}
L477:
	;
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v4167 == int32(0) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v4201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v4201 == int32(0) {
		goto L487
	} else {
		goto L488
	}
L479:
	;
	v4170 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+4))
	if v4170 <= int32(0) {
		goto L478
	} else {
		goto L480
	}
L480:
	;
	v4176 = int32(0)
	goto L481
L481:
	;
	v4178 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4179 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+12))
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v4179+v4176<<(uint(int32(2))%32))))
	v4184 = *(*int32)(unsafe.Add(mBase, uint32(v4183)+12))
	v4185 = F_bms_is_subset(m, v4178, v4184)
	mBase = m.M
	if v4185 == int32(0) {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	goto L478
L483:
	;
	v4193 = v4176 + int32(1)
	v4194 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+4))
	if v4193 < v4194 {
		v4176 = v4193
		goto L481
	} else {
		goto L486
	}
L484:
	;
	v4188 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4183)+12))
	v4190 = F_bms_equal(m, v4188, v4189)
	mBase = m.M
	if v4190 != 0 {
		goto L483
	} else {
		goto L485
	}
L485:
	;
	v4255 = int32(1)
	goto L474
L486:
	;
	goto L482
L487:
	;
	v4247 = int32(0)
	goto L475
L488:
	;
	v4204 = *(*int32)(unsafe.Add(mBase, uint32(v4201)+4))
	if v4204 <= int32(0) {
		goto L487
	} else {
		goto L489
	}
L489:
	;
	v4211 = int32(0)
	goto L490
L490:
	;
	v4213 = *(*int32)(unsafe.Add(mBase, uint32(v4201)+12))
	v4214 = int32(2)
	v4217 = *(*int32)(unsafe.Add(mBase, uint32(v4213+v4211<<(uint(v4214)%32))))
	v4218 = *(*int32)(unsafe.Add(mBase, uint32(v4217)+20))
	if v4218 == v4214 {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	goto L487
L492:
	;
	v4236 = v4211 + int32(1)
	v4237 = *(*int32)(unsafe.Add(mBase, uint32(v4201)+4))
	if v4236 < v4237 {
		v4211 = v4236
		goto L490
	} else {
		goto L500
	}
L493:
	;
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(v4217)+4))
	v4222 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4223 = F_bms_is_subset(m, v4221, v4222)
	mBase = m.M
	if v4223 != 0 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v4224 = *(*int32)(unsafe.Add(mBase, uint32(v4217)+8))
	v4225 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4226 = F_bms_is_subset(m, v4224, v4225)
	mBase = m.M
	if v4226 != 0 {
		goto L492
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	v4227 = int32(1)
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v4217)+4))
	v4229 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4230 = F_bms_overlap(m, v4228, v4229)
	mBase = m.M
	if v4230 != 0 {
		v4247 = v4227
		goto L475
	} else {
		goto L498
	}
L497:
	;
	goto L496
L498:
	;
	v4231 = *(*int32)(unsafe.Add(mBase, uint32(v4217)+8))
	v4232 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4233 = F_bms_overlap(m, v4231, v4232)
	mBase = m.M
	if v4233 != 0 {
		v4247 = v4227
		goto L475
	} else {
		goto L499
	}
L499:
	;
	goto L492
L500:
	;
	goto L491
L501:
	;
	goto L471
L502:
	;
	if v4064 == int32(2) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v4266 = v4128 + int32(1)
	goto L505
L504:
	;
	v4266 = int32(0)
	goto L505
L505:
	;
	v4267 = *(*int32)(unsafe.Add(mBase, uint32(v4258)+4))
	if v4267 <= v4266 {
		goto L469
	} else {
		goto L506
	}
L506:
	;
	v4273 = v4266
	goto L507
L507:
	;
	v4306 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4307 = *(*int32)(unsafe.Add(mBase, uint32(v4258)+12))
	v4311 = *(*int32)(unsafe.Add(mBase, uint32(v4307+v4273<<(uint(int32(2))%32))))
	v4312 = *(*int32)(unsafe.Add(mBase, uint32(v4311)+8))
	v4313 = int32(0)
	if base.B2i32(v4306 == v4313)|base.B2i32(v4312 == v4313) != 0 {
		v4358 = v4313
		goto L511
	} else {
		goto L512
	}
L508:
	;
	goto L469
L509:
	;
	v4370 = v4273 + int32(1)
	v4371 = *(*int32)(unsafe.Add(mBase, uint32(v4258)+4))
	if v4370 < v4371 {
		v4273 = v4370
		goto L507
	} else {
		goto L531
	}
L510:
	;
	if v4358 != 0 {
		goto L509
	} else {
		goto L523
	}
L511:
	;
	goto L510
L512:
	;
	v4323 = *(*int32)(unsafe.Add(mBase, uint32(v4306)+4))
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v4312)+4))
	if v4323 < v4324 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v4326 = v4323
	goto L515
L514:
	;
	v4326 = v4324
	goto L515
L515:
	;
	if v4326 <= int32(1) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v4329 = int32(1)
	goto L518
L517:
	;
	v4329 = v4326
	goto L518
L518:
	;
	v4330 = int32(8)
	v4335 = int32(0)
	goto L519
L519:
	;
	v4342 = v4335 << (uint(int32(2)) % 32)
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(v4312+v4330+v4342)))
	v4346 = *(*int32)(unsafe.Add(mBase, uint32(v4306+v4330+v4342)))
	v4347 = v4344 & v4346
	v4349 = base.B2i32(v4347 != int32(0))
	if v4347 != 0 {
		v4358 = v4349
		goto L511
	} else {
		goto L521
	}
L520:
	;
	v4358 = v4349
	goto L511
L521:
	;
	v4351 = v4335 + int32(1)
	if v4351 != v4329 {
		v4335 = v4351
		goto L519
	} else {
		goto L522
	}
L522:
	;
	goto L520
L523:
	;
	v4359 = F_have_relevant_joinclause(m, l0, v4158, v4311)
	mBase = m.M
	v4360 = m.ExcPending
	if v4360 != 0 {
		goto L13
	} else {
		goto L524
	}
L524:
	;
	if v4359 == int32(0) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v4363 = F_have_join_order_restriction(m, l0, v4158, v4311)
	mBase = m.M
	v4364 = m.ExcPending
	if v4364 != 0 {
		goto L13
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v4367 = F_make_join_rel(m, l0, v4158, v4311)
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L13
	} else {
		goto L530
	}
L528:
	;
	if v4363 == int32(0) {
		goto L509
	} else {
		goto L529
	}
L529:
	;
	goto L527
L530:
	;
	goto L509
L531:
	;
	goto L508
L532:
	;
	v4376 = int32(0)
	v4377 = *(*int32)(unsafe.Add(mBase, uint32(v4373)+4))
	if v4377 <= v4376 {
		goto L469
	} else {
		goto L533
	}
L533:
	;
	v4384 = v4376
	goto L534
L534:
	;
	v4417 = *(*int32)(unsafe.Add(mBase, uint32(v4373)+12))
	v4421 = *(*int32)(unsafe.Add(mBase, uint32(v4417+v4384<<(uint(int32(2))%32))))
	v4422 = *(*int32)(unsafe.Add(mBase, uint32(v4421)+8))
	v4423 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+8))
	v4424 = int32(0)
	if base.B2i32(v4422 == v4424)|base.B2i32(v4423 == v4424) != 0 {
		v4469 = v4424
		goto L537
	} else {
		goto L538
	}
L535:
	;
	goto L469
L536:
	;
	if v4469 == int32(0) {
		goto L549
	} else {
		goto L550
	}
L537:
	;
	goto L536
L538:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+4))
	v4435 = *(*int32)(unsafe.Add(mBase, uint32(v4423)+4))
	if v4434 < v4435 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v4437 = v4434
	goto L541
L540:
	;
	v4437 = v4435
	goto L541
L541:
	;
	if v4437 <= int32(1) {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v4440 = int32(1)
	goto L544
L543:
	;
	v4440 = v4437
	goto L544
L544:
	;
	v4441 = int32(8)
	v4446 = int32(0)
	goto L545
L545:
	;
	v4453 = v4446 << (uint(int32(2)) % 32)
	v4455 = *(*int32)(unsafe.Add(mBase, uint32(v4423+v4441+v4453)))
	v4457 = *(*int32)(unsafe.Add(mBase, uint32(v4422+v4441+v4453)))
	v4458 = v4455 & v4457
	v4460 = base.B2i32(v4458 != int32(0))
	if v4458 != 0 {
		v4469 = v4460
		goto L537
	} else {
		goto L547
	}
L546:
	;
	v4469 = v4460
	goto L537
L547:
	;
	v4462 = v4446 + int32(1)
	if v4462 != v4440 {
		v4446 = v4462
		goto L545
	} else {
		goto L548
	}
L548:
	;
	goto L546
L549:
	;
	v4472 = F_make_join_rel(m, l0, v4158, v4421)
	mBase = m.M
	v4473 = m.ExcPending
	if v4473 != 0 {
		goto L13
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	v4475 = v4384 + int32(1)
	v4476 = *(*int32)(unsafe.Add(mBase, uint32(v4373)+4))
	if v4475 < v4476 {
		v4384 = v4475
		goto L534
	} else {
		goto L553
	}
L552:
	;
	goto L551
L553:
	;
	goto L535
L554:
	;
	goto L468
L555:
	;
	v4562 = v4104
	v4568 = v4557
	goto L558
L556:
	;
	goto L557
L557:
	;
	v4984 = *(*int32)(unsafe.Add(mBase, uint32(v4108)))
	if v4984 != 0 {
		goto L629
	} else {
		goto L630
	}
L558:
	;
	v4600 = *(*int32)(unsafe.Add(mBase, uint32(v4105+v4562<<(uint(int32(2))%32))))
	if v4600 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L559:
	;
	goto L557
L560:
	;
	v4944 = v4562 + int32(1)
	v4945 = v4064 - v4944
	if v4944 <= v4945 {
		v4562 = v4944
		v4568 = v4945
		goto L558
	} else {
		goto L628
	}
L561:
	;
	v4603 = *(*int32)(unsafe.Add(mBase, uint32(v4600)+4))
	if v4603 <= int32(0) {
		goto L560
	} else {
		goto L562
	}
L562:
	;
	v4621 = int32(0)
	goto L563
L563:
	;
	v4647 = *(*int32)(unsafe.Add(mBase, uint32(v4600)+12))
	v4651 = *(*int32)(unsafe.Add(mBase, uint32(v4647+v4621<<(uint(int32(2))%32))))
	v4652 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+228))
	if v4652 != 0 {
		goto L566
	} else {
		goto L567
	}
L564:
	;
	goto L560
L565:
	;
	v4903 = v4621 + int32(1)
	v4904 = *(*int32)(unsafe.Add(mBase, uint32(v4600)+4))
	if v4903 < v4904 {
		v4621 = v4903
		goto L563
	} else {
		goto L627
	}
L566:
	;
	v4751 = *(*int32)(unsafe.Add(mBase, uint32(v4105+v4568<<(uint(int32(2))%32))))
	if v4751 == int32(0) {
		goto L565
	} else {
		goto L597
	}
L567:
	;
	v4653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4651)+232)))
	if v4653 != 0 {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	v4657 = int32(1)
	v4658 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+72))
	if v4658 != 0 {
		v4740 = v4657
		goto L570
	} else {
		goto L571
	}
L569:
	;
	if v4748 == int32(0) {
		goto L565
	} else {
		goto L596
	}
L570:
	;
	v4748 = v4740
	goto L569
L571:
	;
	v4659 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+112))
	if v4659 != 0 {
		v4740 = v4657
		goto L570
	} else {
		goto L572
	}
L572:
	;
	v4660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v4660 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v4694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v4694 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L574:
	;
	v4663 = *(*int32)(unsafe.Add(mBase, uint32(v4660)+4))
	if v4663 <= int32(0) {
		goto L573
	} else {
		goto L575
	}
L575:
	;
	v4669 = int32(0)
	goto L576
L576:
	;
	v4671 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4672 = *(*int32)(unsafe.Add(mBase, uint32(v4660)+12))
	v4676 = *(*int32)(unsafe.Add(mBase, uint32(v4672+v4669<<(uint(int32(2))%32))))
	v4677 = *(*int32)(unsafe.Add(mBase, uint32(v4676)+12))
	v4678 = F_bms_is_subset(m, v4671, v4677)
	mBase = m.M
	if v4678 == int32(0) {
		goto L578
	} else {
		goto L579
	}
L577:
	;
	goto L573
L578:
	;
	v4686 = v4669 + int32(1)
	v4687 = *(*int32)(unsafe.Add(mBase, uint32(v4660)+4))
	if v4686 < v4687 {
		v4669 = v4686
		goto L576
	} else {
		goto L581
	}
L579:
	;
	v4681 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4682 = *(*int32)(unsafe.Add(mBase, uint32(v4676)+12))
	v4683 = F_bms_equal(m, v4681, v4682)
	mBase = m.M
	if v4683 != 0 {
		goto L578
	} else {
		goto L580
	}
L580:
	;
	v4748 = int32(1)
	goto L569
L581:
	;
	goto L577
L582:
	;
	v4740 = int32(0)
	goto L570
L583:
	;
	v4697 = *(*int32)(unsafe.Add(mBase, uint32(v4694)+4))
	if v4697 <= int32(0) {
		goto L582
	} else {
		goto L584
	}
L584:
	;
	v4704 = int32(0)
	goto L585
L585:
	;
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v4694)+12))
	v4707 = int32(2)
	v4710 = *(*int32)(unsafe.Add(mBase, uint32(v4706+v4704<<(uint(v4707)%32))))
	v4711 = *(*int32)(unsafe.Add(mBase, uint32(v4710)+20))
	if v4711 == v4707 {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	goto L582
L587:
	;
	v4729 = v4704 + int32(1)
	v4730 = *(*int32)(unsafe.Add(mBase, uint32(v4694)+4))
	if v4729 < v4730 {
		v4704 = v4729
		goto L585
	} else {
		goto L595
	}
L588:
	;
	v4714 = *(*int32)(unsafe.Add(mBase, uint32(v4710)+4))
	v4715 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4716 = F_bms_is_subset(m, v4714, v4715)
	mBase = m.M
	if v4716 != 0 {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v4717 = *(*int32)(unsafe.Add(mBase, uint32(v4710)+8))
	v4718 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4719 = F_bms_is_subset(m, v4717, v4718)
	mBase = m.M
	if v4719 != 0 {
		goto L587
	} else {
		goto L592
	}
L590:
	;
	goto L591
L591:
	;
	v4720 = int32(1)
	v4721 = *(*int32)(unsafe.Add(mBase, uint32(v4710)+4))
	v4722 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4723 = F_bms_overlap(m, v4721, v4722)
	mBase = m.M
	if v4723 != 0 {
		v4740 = v4720
		goto L570
	} else {
		goto L593
	}
L592:
	;
	goto L591
L593:
	;
	v4724 = *(*int32)(unsafe.Add(mBase, uint32(v4710)+8))
	v4725 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4726 = F_bms_overlap(m, v4724, v4725)
	mBase = m.M
	if v4726 != 0 {
		v4740 = v4720
		goto L570
	} else {
		goto L594
	}
L594:
	;
	goto L587
L595:
	;
	goto L586
L596:
	;
	goto L566
L597:
	;
	if v4562 == v4568 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v4758 = v4621 + int32(1)
	goto L600
L599:
	;
	v4758 = int32(0)
	goto L600
L600:
	;
	v4759 = *(*int32)(unsafe.Add(mBase, uint32(v4751)+4))
	if v4759 <= v4758 {
		goto L565
	} else {
		goto L601
	}
L601:
	;
	v4765 = v4758
	goto L602
L602:
	;
	v4798 = *(*int32)(unsafe.Add(mBase, uint32(v4651)+8))
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v4751)+12))
	v4803 = *(*int32)(unsafe.Add(mBase, uint32(v4799+v4765<<(uint(int32(2))%32))))
	v4804 = *(*int32)(unsafe.Add(mBase, uint32(v4803)+8))
	v4805 = int32(0)
	if base.B2i32(v4798 == v4805)|base.B2i32(v4804 == v4805) != 0 {
		v4850 = v4805
		goto L606
	} else {
		goto L607
	}
L603:
	;
	goto L565
L604:
	;
	v4862 = v4765 + int32(1)
	v4863 = *(*int32)(unsafe.Add(mBase, uint32(v4751)+4))
	if v4862 < v4863 {
		v4765 = v4862
		goto L602
	} else {
		goto L626
	}
L605:
	;
	if v4850 != 0 {
		goto L604
	} else {
		goto L618
	}
L606:
	;
	goto L605
L607:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v4798)+4))
	v4816 = *(*int32)(unsafe.Add(mBase, uint32(v4804)+4))
	if v4815 < v4816 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v4818 = v4815
	goto L610
L609:
	;
	v4818 = v4816
	goto L610
L610:
	;
	if v4818 <= int32(1) {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v4821 = int32(1)
	goto L613
L612:
	;
	v4821 = v4818
	goto L613
L613:
	;
	v4822 = int32(8)
	v4827 = int32(0)
	goto L614
L614:
	;
	v4834 = v4827 << (uint(int32(2)) % 32)
	v4836 = *(*int32)(unsafe.Add(mBase, uint32(v4804+v4822+v4834)))
	v4838 = *(*int32)(unsafe.Add(mBase, uint32(v4798+v4822+v4834)))
	v4839 = v4836 & v4838
	v4841 = base.B2i32(v4839 != int32(0))
	if v4839 != 0 {
		v4850 = v4841
		goto L606
	} else {
		goto L616
	}
L615:
	;
	v4850 = v4841
	goto L606
L616:
	;
	v4843 = v4827 + int32(1)
	if v4843 != v4821 {
		v4827 = v4843
		goto L614
	} else {
		goto L617
	}
L617:
	;
	goto L615
L618:
	;
	v4851 = F_have_relevant_joinclause(m, l0, v4651, v4803)
	mBase = m.M
	v4852 = m.ExcPending
	if v4852 != 0 {
		goto L13
	} else {
		goto L619
	}
L619:
	;
	if v4851 == int32(0) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v4855 = F_have_join_order_restriction(m, l0, v4651, v4803)
	mBase = m.M
	v4856 = m.ExcPending
	if v4856 != 0 {
		goto L13
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	v4859 = F_make_join_rel(m, l0, v4651, v4803)
	mBase = m.M
	v4860 = m.ExcPending
	if v4860 != 0 {
		goto L13
	} else {
		goto L625
	}
L623:
	;
	if v4855 == int32(0) {
		goto L604
	} else {
		goto L624
	}
L624:
	;
	goto L622
L625:
	;
	goto L604
L626:
	;
	goto L603
L627:
	;
	goto L564
L628:
	;
	goto L559
L629:
	;
	m.G0 = v4101 + int32(16)
	v5308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v5312 = *(*int32)(unsafe.Add(mBase, uint32(v5308+v4064<<(uint(int32(2))%32))))
	if v5312 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L630:
	;
	v4985 = *(*int32)(unsafe.Add(mBase, uint32(v4110)))
	if v4985 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v4986 = *(*int32)(unsafe.Add(mBase, uint32(v4985)+4))
	if int32(0) < v4986 {
		goto L634
	} else {
		goto L635
	}
L632:
	;
	goto L633
L633:
	;
	v5253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v5253 != 0 {
		goto L629
	} else {
		goto L664
	}
L634:
	;
	v5001 = int32(0)
	goto L637
L635:
	;
	goto L636
L636:
	;
	v5215 = *(*int32)(unsafe.Add(mBase, uint32(v4108)))
	if v5215 != 0 {
		goto L629
	} else {
		goto L663
	}
L637:
	;
	v5027 = *(*int32)(unsafe.Add(mBase, uint32(v4105)+4))
	if v5027 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L638:
	;
	goto L636
L639:
	;
	v5175 = v5001 + int32(1)
	v5176 = *(*int32)(unsafe.Add(mBase, uint32(v4985)+4))
	if v5175 < v5176 {
		v5001 = v5175
		goto L637
	} else {
		goto L662
	}
L640:
	;
	v5030 = *(*int32)(unsafe.Add(mBase, uint32(v5027)+4))
	if v5030 <= int32(0) {
		goto L639
	} else {
		goto L641
	}
L641:
	;
	v5033 = *(*int32)(unsafe.Add(mBase, uint32(v4985)+12))
	v5037 = *(*int32)(unsafe.Add(mBase, uint32(v5033+v5001<<(uint(int32(2))%32))))
	v5043 = int32(0)
	goto L642
L642:
	;
	v5076 = *(*int32)(unsafe.Add(mBase, uint32(v5027)+12))
	v5080 = *(*int32)(unsafe.Add(mBase, uint32(v5076+v5043<<(uint(int32(2))%32))))
	v5081 = *(*int32)(unsafe.Add(mBase, uint32(v5080)+8))
	v5082 = *(*int32)(unsafe.Add(mBase, uint32(v5037)+8))
	v5083 = int32(0)
	if base.B2i32(v5081 == v5083)|base.B2i32(v5082 == v5083) != 0 {
		v5128 = v5083
		goto L645
	} else {
		goto L646
	}
L643:
	;
	goto L639
L644:
	;
	if v5128 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L645:
	;
	goto L644
L646:
	;
	v5093 = *(*int32)(unsafe.Add(mBase, uint32(v5081)+4))
	v5094 = *(*int32)(unsafe.Add(mBase, uint32(v5082)+4))
	if v5093 < v5094 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v5096 = v5093
	goto L649
L648:
	;
	v5096 = v5094
	goto L649
L649:
	;
	if v5096 <= int32(1) {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	v5099 = int32(1)
	goto L652
L651:
	;
	v5099 = v5096
	goto L652
L652:
	;
	v5100 = int32(8)
	v5105 = int32(0)
	goto L653
L653:
	;
	v5112 = v5105 << (uint(int32(2)) % 32)
	v5114 = *(*int32)(unsafe.Add(mBase, uint32(v5082+v5100+v5112)))
	v5116 = *(*int32)(unsafe.Add(mBase, uint32(v5081+v5100+v5112)))
	v5117 = v5114 & v5116
	v5119 = base.B2i32(v5117 != int32(0))
	if v5117 != 0 {
		v5128 = v5119
		goto L645
	} else {
		goto L655
	}
L654:
	;
	v5128 = v5119
	goto L645
L655:
	;
	v5121 = v5105 + int32(1)
	if v5121 != v5099 {
		v5105 = v5121
		goto L653
	} else {
		goto L656
	}
L656:
	;
	goto L654
L657:
	;
	v5131 = F_make_join_rel(m, l0, v5037, v5080)
	mBase = m.M
	v5132 = m.ExcPending
	if v5132 != 0 {
		goto L13
	} else {
		goto L660
	}
L658:
	;
	goto L659
L659:
	;
	v5134 = v5043 + int32(1)
	v5135 = *(*int32)(unsafe.Add(mBase, uint32(v5027)+4))
	if v5134 < v5135 {
		v5043 = v5134
		goto L642
	} else {
		goto L661
	}
L660:
	;
	goto L659
L661:
	;
	goto L643
L662:
	;
	goto L638
L663:
	;
	goto L633
L664:
	;
	v5254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+333)))
	if v5254 != 0 {
		goto L629
	} else {
		goto L665
	}
L665:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5258 = m.ExcPending
	if v5258 != 0 {
		goto L13
	} else {
		goto L666
	}
L666:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4101))) = v4064
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_18), v4101)
	mBase = m.M
	v5262 = m.ExcPending
	if v5262 != 0 {
		goto L13
	} else {
		goto L667
	}
L667:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_19), int32(260), int32(_a_F_make_rel_from_joinlist_20))
	mBase = m.M
	v5267 = m.ExcPending
	if v5267 != 0 {
		goto L13
	} else {
		goto L668
	}
L668:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L669:
	;
	v5477 = v4064 + int32(1)
	if v5477 <= v44 {
		v4064 = v5477
		goto L462
	} else {
		goto L697
	}
L670:
	;
	v5315 = int32(0)
	v5316 = *(*int32)(unsafe.Add(mBase, uint32(v5312)+4))
	if v5316 <= v5315 {
		goto L669
	} else {
		goto L671
	}
L671:
	;
	v5320 = v5315
	goto L672
L672:
	;
	v5356 = *(*int32)(unsafe.Add(mBase, uint32(v5312)+12))
	v5360 = *(*int32)(unsafe.Add(mBase, uint32(v5356+v5320<<(uint(int32(2))%32))))
	v5361 = *(*int32)(unsafe.Add(mBase, uint32(v5360)+8))
	v5362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5363 = int32(0)
	if base.B2i32(v5361 == v5363)|base.B2i32(v5362 == v5363) != 0 {
		v5409 = base.B2i32(v5361|v5362 == v5363)
		goto L675
	} else {
		goto L676
	}
L673:
	;
	goto L669
L674:
	;
	F_generate_partitionwise_join_paths(m, l0, v5360)
	mBase = m.M
	v5415 = m.ExcPending
	if v5415 != 0 {
		goto L13
	} else {
		goto L685
	}
L675:
	;
	goto L674
L676:
	;
	v5377 = *(*int32)(unsafe.Add(mBase, uint32(v5361)+4))
	v5378 = *(*int32)(unsafe.Add(mBase, uint32(v5362)+4))
	if v5377 != v5378 {
		v5409 = int32(0)
		goto L675
	} else {
		goto L677
	}
L677:
	;
	v5380 = int32(1)
	if v5377 <= v5380 {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	v5383 = v5380
	goto L680
L679:
	;
	v5383 = v5377
	goto L680
L680:
	;
	v5384 = int32(8)
	v5389 = int32(0)
	goto L681
L681:
	;
	v5397 = v5389 << (uint(int32(2)) % 32)
	v5399 = *(*int32)(unsafe.Add(mBase, uint32(v5361+v5384+v5397)))
	v5401 = *(*int32)(unsafe.Add(mBase, uint32(v5362+v5384+v5397)))
	v5402 = base.B2i32(v5399 == v5401)
	if v5399 != v5401 {
		v5409 = v5402
		goto L675
	} else {
		goto L683
	}
L682:
	;
	v5409 = v5402
	goto L675
L683:
	;
	v5405 = v5389 + int32(1)
	if v5405 != v5383 {
		v5389 = v5405
		goto L681
	} else {
		goto L684
	}
L684:
	;
	goto L682
L685:
	;
	if v5409 == int32(0) {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	F_generate_useful_gather_paths(m, l0, v5360, int32(0))
	mBase = m.M
	v5420 = m.ExcPending
	if v5420 != 0 {
		goto L13
	} else {
		goto L689
	}
L687:
	;
	goto L688
L688:
	;
	F_set_cheapest(m, v5360)
	mBase = m.M
	v5422 = m.ExcPending
	if v5422 != 0 {
		goto L13
	} else {
		goto L690
	}
L689:
	;
	goto L688
L690:
	;
	v5423 = *(*int32)(unsafe.Add(mBase, uint32(v5360)+240))
	v5424 = int32(0)
	if (v5409|base.B2i32(v5423 == v5424))&int32(1) == v5424 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	F_generate_grouped_paths(m, l0, v5423, v5360)
	mBase = m.M
	v5432 = m.ExcPending
	if v5432 != 0 {
		goto L13
	} else {
		goto L694
	}
L692:
	;
	goto L693
L693:
	;
	v5436 = v5320 + int32(1)
	v5437 = *(*int32)(unsafe.Add(mBase, uint32(v5312)+4))
	if v5436 < v5437 {
		v5320 = v5436
		goto L672
	} else {
		goto L696
	}
L694:
	;
	F_set_cheapest(m, v5423)
	mBase = m.M
	v5434 = m.ExcPending
	if v5434 != 0 {
		goto L13
	} else {
		goto L695
	}
L695:
	;
	goto L693
L696:
	;
	goto L673
L697:
	;
	goto L463
L698:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5526 = m.ExcPending
	if v5526 != 0 {
		goto L13
	} else {
		goto L701
	}
L699:
	;
	goto L700
L700:
	;
	v5536 = *(*int32)(unsafe.Add(mBase, uint32(v5520)+12))
	v5537 = *(*int32)(unsafe.Add(mBase, uint32(v5536)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(0)
	m.G0 = v4048 + int32(16)
	v5546 = v5537
	goto L1
L701:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4048))) = v44
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_18), v4048)
	mBase = m.M
	v5530 = m.ExcPending
	if v5530 != 0 {
		goto L13
	} else {
		goto L702
	}
L702:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_1), int32(4042), int32(_a_F_make_rel_from_joinlist_21))
	mBase = m.M
	v5535 = m.ExcPending
	if v5535 != 0 {
		goto L13
	} else {
		goto L703
	}
L703:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_rel_consider_parallel(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	switch v6 {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		v8 = F_get_rel_persistence(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if v8 == int32(116) {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				if v12 != 0 {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v14 = F_func_parallel(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						if v14 != int32(115) {
							return
						} else {
							v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
							v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
							v20 = F_is_parallel_safe(m, l0, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								if v20 == int32(0) {
									return
								} else {
									v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
									if v24 != int32(102) {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
										v72 = F_is_parallel_safe(m, l0, v71)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											if v72 == int32(0) {
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
												v78 = F_is_parallel_safe(m, l0, v77)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													if v78 == int32(0) {
													} else {
														v82 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
													}
													return
												}
											}
										}
									} else {
										v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
										v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+144))
										if v28 == int32(0) {
											return
										} else {
											v31 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
											mBase = m.M
											v32 = m.ExcPending
											if v32 != 0 {
												return
											} else {
												if v31 != 0 {
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
													v72 = F_is_parallel_safe(m, l0, v71)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														if v72 == int32(0) {
															return
														} else {
															v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
															v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
															v78 = F_is_parallel_safe(m, l0, v77)
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return
															} else {
																if v78 == int32(0) {
																} else {
																	v82 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
																}
																return
															}
														}
													}
												} else {
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
					if v24 != int32(102) {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
						v72 = F_is_parallel_safe(m, l0, v71)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							if v72 == int32(0) {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								v78 = F_is_parallel_safe(m, l0, v77)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									if v78 == int32(0) {
									} else {
										v82 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
									}
									return
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+144))
						if v28 == int32(0) {
							return
						} else {
							v31 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								if v31 != 0 {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
									v72 = F_is_parallel_safe(m, l0, v71)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										if v72 == int32(0) {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
											v78 = F_is_parallel_safe(m, l0, v77)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												if v78 == int32(0) {
												} else {
													v82 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
												}
												return
											}
										}
									}
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	case 1:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
		if v34 != 0 {
			v35 = int32(1)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			if v36 != int32(7) {
				v56 = v35
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
				if v39 != int32(1) {
					v56 = v35
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
					if v43 == int32(0) {
						v56 = int32(0)
					} else {
						v46 = int32(1)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v47 != int32(7) {
							v56 = v46
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
							if v50 != 0 {
								v56 = int32(0)
							} else {
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
								if v51 != int64(0) {
									v56 = v46
								} else {
									v56 = int32(0)
								}
							}
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
			if v43 == int32(0) {
				v56 = int32(0)
			} else {
				v46 = int32(1)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47 != int32(7) {
					v56 = v46
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
					if v50 != 0 {
						v56 = int32(0)
					} else {
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
						if v51 != int64(0) {
							v56 = v46
						} else {
							v56 = int32(0)
						}
					}
				}
			}
		}
		if v56 == int32(0) {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
			v72 = F_is_parallel_safe(m, l0, v71)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return
			} else {
				if v72 == int32(0) {
					return
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
					v78 = F_is_parallel_safe(m, l0, v77)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						if v78 == int32(0) {
						} else {
							v82 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
						}
						return
					}
				}
			}
		} else {
			return
		}
	case 2, 4, 6, 7, 9:
		return
	case 3:
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
		v61 = F_is_parallel_safe(m, l0, v60)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			if v61 != 0 {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				v72 = F_is_parallel_safe(m, l0, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					if v72 == int32(0) {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						v78 = F_is_parallel_safe(m, l0, v77)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							if v78 == int32(0) {
							} else {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
							}
							return
						}
					}
				}
			} else {
				return
			}
		}
	case 5:
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
		v64 = F_is_parallel_safe(m, l0, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return
		} else {
			if v64 == int32(0) {
				return
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				v72 = F_is_parallel_safe(m, l0, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					if v72 == int32(0) {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						v78 = F_is_parallel_safe(m, l0, v77)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							if v78 == int32(0) {
							} else {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
							}
							return
						}
					}
				}
			}
		}
	default:
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
		v72 = F_is_parallel_safe(m, l0, v71)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return
		} else {
			if v72 == int32(0) {
				return
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v78 = F_is_parallel_safe(m, l0, v77)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if v78 == int32(0) {
					} else {
						v82 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v82)
					}
					return
				}
			}
		}
	}
}
func F_set_rel_pathlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 float64
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 float64
	_ = v449
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 float64
	_ = v461
	var v462 float64
	_ = v462
	var v463 int32
	_ = v463
	var v464 int64
	_ = v464
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 float64
	_ = v515
	var v516 float64
	_ = v516
	var v534 float64
	_ = v534
	var v542 float64
	_ = v542
	var v543 float64
	_ = v543
	var v545 float64
	_ = v545
	var v547 float64
	_ = v547
	var v548 float64
	_ = v548
	var v566 float64
	_ = v566
	var v574 float64
	_ = v574
	var v575 float64
	_ = v575
	var v577 float64
	_ = v577
	var v578 int32
	_ = v578
	var v579 float64
	_ = v579
	var v580 int64
	_ = v580
	var v585 float64
	_ = v585
	var v586 float64
	_ = v586
	var v590 int32
	_ = v590
	var v591 int64
	_ = v591
	var v596 float64
	_ = v596
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 float64
	_ = v655
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 float64
	_ = v667
	var v668 float64
	_ = v668
	var v669 int32
	_ = v669
	var v670 int64
	_ = v670
	var v679 int32
	_ = v679
	var v688 int32
	_ = v688
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 float64
	_ = v721
	var v722 float64
	_ = v722
	var v740 float64
	_ = v740
	var v748 float64
	_ = v748
	var v749 float64
	_ = v749
	var v751 float64
	_ = v751
	var v753 float64
	_ = v753
	var v754 float64
	_ = v754
	var v772 float64
	_ = v772
	var v780 float64
	_ = v780
	var v781 float64
	_ = v781
	var v783 float64
	_ = v783
	var v784 int32
	_ = v784
	var v785 float64
	_ = v785
	var v786 int64
	_ = v786
	var v791 float64
	_ = v791
	var v792 float64
	_ = v792
	var v796 int32
	_ = v796
	var v797 int64
	_ = v797
	var v802 float64
	_ = v802
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 float64
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 float64
	_ = v845
	var v846 int64
	_ = v846
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 float64
	_ = v896
	var v897 float64
	_ = v897
	var v898 float64
	_ = v898
	var v916 float64
	_ = v916
	var v918 float64
	_ = v918
	var v924 float64
	_ = v924
	var v925 float64
	_ = v925
	var v927 float64
	_ = v927
	var v929 float64
	_ = v929
	var v931 float64
	_ = v931
	var v933 float64
	_ = v933
	var v934 float64
	_ = v934
	var v952 float64
	_ = v952
	var v954 float64
	_ = v954
	var v955 float64
	_ = v955
	var v960 float64
	_ = v960
	var v961 float64
	_ = v961
	var v963 float64
	_ = v963
	var v964 int32
	_ = v964
	var v965 float64
	_ = v965
	var v966 int64
	_ = v966
	var v969 float64
	_ = v969
	var v970 float64
	_ = v970
	var v974 int32
	_ = v974
	var v975 int64
	_ = v975
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int64
	_ = v1013
	var v1015 int64
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1119 float64
	_ = v1119
	var v1120 float64
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1369 int32
	_ = v1369
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1491 int32
	_ = v1491
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1624 int32
	_ = v1624
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1662 int32
	_ = v1662
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1794 int32
	_ = v1794
	var v1798 int32
	_ = v1798
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1877 int32
	_ = v1877
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1908 int32
	_ = v1908
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1958 float64
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1965 int32
	_ = v1965
	var v1968 int32
	_ = v1968
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v2015 int32
	_ = v2015
	var v2021 int32
	_ = v2021
	var v2030 int32
	_ = v2030
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2171 int32
	_ = v2171
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2225 int32
	_ = v2225
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2280 int32
	_ = v2280
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2302 int32
	_ = v2302
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 float64
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2382 int32
	_ = v2382
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2407 int32
	_ = v2407
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2452 int32
	_ = v2452
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2492 int32
	_ = v2492
	var v2494 int32
	_ = v2494
	v5 = int32(0)
	v18 = float64(0)
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v32 == v5 {
		v53 = v5
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[0]))
	if v2420 != 0 {
		goto L428
	} else {
		goto L429
	}
L2:
	;
	if v53 != 0 {
		v2394 = l0
		v2395 = l1
		v2396 = l2
		v2397 = l3
		v2407 = v28
		goto L1
	} else {
		goto L12
	}
L3:
	;
	goto L2
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v36 = v35
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if base.Ui32(int32(2)) <= base.Ui32(v40-int32(303)) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v53 = int32(1)
	goto L3
L7:
	;
	if v40 != int32(293) {
		v53 = v5
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v36 = v39 + int32(72)
	goto L5
L9:
	;
	goto L6
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	if v47 != 0 {
		v53 = v5
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v54 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v57 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	switch v175 {
	case 0:
		goto L51
	case 1, 6, 7, 8:
		v2394 = l0
		v2395 = l1
		v2396 = l2
		v2397 = l3
		v2407 = v28
		goto L1
	default:
		goto L47
	case 3:
		goto L50
	case 4:
		goto L49
	case 5:
		goto L48
	}
L16:
	;
	F_add_paths_to_append_rel(m, l0, l1, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if int32(0) < v63 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	return
L20:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L21:
	;
	v71 = v5
	v72 = v5
	goto L24
L22:
	;
	v153 = v5
	goto L23
L23:
	;
	F_add_paths_to_append_rel(m, l0, l1, v153)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L19
	} else {
		goto L45
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v72<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v96 != l2 {
		v141 = v71
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v153 = v141
	goto L23
L26:
	;
	v145 = v72 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v145 < v146 {
		v71 = v141
		v72 = v145
		goto L24
	} else {
		goto L44
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v100 = v98 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104+v100)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v107 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+26)) = uint8(v110)
	goto L30
L29:
	;
	goto L30
L30:
	;
	F_set_rel_pathlist(m, l0, v103, v98, v106)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v114 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+44))
	if v116 == v114 {
		v137 = v114
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v137 != 0 {
		v141 = v71
		goto L26
	} else {
		goto L42
	}
L33:
	;
	goto L32
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v120 = v119
	goto L35
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if base.Ui32(int32(2)) <= base.Ui32(v124-int32(303)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v137 = int32(1)
	goto L33
L37:
	;
	if v124 != int32(293) {
		v137 = v114
		goto L33
	} else {
		goto L40
	}
L38:
	;
	v120 = v123 + int32(72)
	goto L35
L39:
	;
	goto L36
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v123)+72))
	if v131 != 0 {
		v137 = v114
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v138 = F_lappend(m, v71, v103)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v141 = v138
	goto L26
L44:
	;
	goto L25
L45:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L46:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1009 = m.G0
	v1011 = v1009 - int32(16)
	m.G0 = v1011
	v1013 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v1015 = v1013 & int64(16)
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v1019 = F_TidQualFromRestrictInfoList(m, l0, v1016, l1, v1011+int32(15))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L19
	} else {
		goto L185
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L19
	} else {
		goto L179
	}
L48:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v818 = F_palloc0(m, int32(72))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L19
	} else {
		goto L162
	}
L49:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v612 = F_palloc0(m, int32(72))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L19
	} else {
		goto L137
	}
L50:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+72)))
	if v272 != int32(1) {
		v387 = v5
		goto L85
	} else {
		goto L86
	}
L51:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v176 == int32(102) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	m.T0[v181].(func(*base.Module, int32, int32, int32))(m, l0, l1, v179)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v184 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L55:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v189 = F_palloc0(m, int32(72))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v189))) = int64(1477468750106)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+12)) = v194
	v196 = F_get_baserel_parampathinfo(m, l0, l1, v187)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+20)) = uint8(v198)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+16)) = v196
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+64)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v189)+24)) = v198
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+21)) = uint8(v201)
	F_cost_samplescan(m, v189, l0, l1, v196)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v209) <= base.Ui32(int32(1)) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	F_add_path(m, l1, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L19
	} else {
		goto L84
	}
L61:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v213 = int32(0)
	if v212 == v213 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v263 = F_GetTsmRoutine(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L81
	}
L64:
	;
	if v258 == int32(1) {
		v268 = v189
		goto L60
	} else {
		goto L80
	}
L65:
	;
	v258 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v221 = int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v222 <= v221 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v225 = v221
	goto L70
L69:
	;
	v225 = v222
	goto L70
L70:
	;
	v229 = int32(0)
	v231 = v213
	goto L71
L71:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v212+int32(8)+v229<<(uint(int32(2))%32))))
	if v238 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v258 = v250
	goto L64
L73:
	;
	goto L72
L74:
	;
	v239 = int32(2)
	if v231 != 0 {
		v250 = v239
		goto L73
	} else {
		goto L77
	}
L75:
	;
	v245 = v231
	goto L76
L76:
	;
	v247 = v229 + int32(1)
	if v247 != v225 {
		v229 = v247
		v231 = v245
		goto L71
	} else {
		goto L79
	}
L77:
	;
	v240 = int32(1)
	if base.Ui32(v240) < base.Ui32(base.I32_popcnt(v238)) {
		v250 = v239
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v245 = v240
	goto L76
L79:
	;
	v250 = v245
	goto L73
L80:
	;
	goto L63
L81:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+9)))
	if v265 != 0 {
		v268 = v189
		goto L60
	} else {
		goto L82
	}
L82:
	;
	v266 = F_create_material_path(m, l1, v189)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L83
	}
L83:
	;
	v268 = v266
	goto L60
L84:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L85:
	;
	v407 = F_palloc0(m, int32(72))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L19
	} else {
		goto L112
	}
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v276 == int32(0) {
		v387 = v5
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v279 <= int32(0) {
		v387 = v5
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+90)))
	v290 = v5
	goto L90
L89:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v330 = m.G0
	v332 = v330 - int32(48)
	m.G0 = v332
	v341 = F_get_ordering_op_properties(m, int32(412), v332+int32(44), v332+int32(40), v332+int32(36))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L19
	} else {
		goto L99
	}
L90:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v282+v290<<(uint(int32(2))%32))))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v313 != int32(6) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v387 = int32(0)
	goto L85
L92:
	;
	v325 = v290 + int32(1)
	if v279 != v325 {
		v290 = v325
		goto L90
	} else {
		goto L97
	}
L93:
	;
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312)+8)))
	if v316 != v283 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v318 != v319 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v312)+28))
	if v321 == int32(0) {
		goto L89
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	goto L91
L98:
	;
	v387 = v361
	goto L85
L99:
	;
	if v341 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v332)+44))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v332)+40))
	v345 = F_exprCollation(m, v312)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L19
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L109
	}
L103:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v332)+36))
	v349 = base.B2i32(v347 == int32(5))
	v350 = int32(0)
	v352 = F_make_pathkey_from_sortinfo(m, l0, v312, v343, v344, v345, v349, v349, v350, v328, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L19
	} else {
		goto L104
	}
L104:
	;
	if v352 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332)+12)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v332)+32)) = v352
	v359 = F_list_make1_impl(m, int32(1), v332+int32(12))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L19
	} else {
		goto L108
	}
L106:
	;
	v361 = int32(0)
	goto L107
L107:
	;
	m.G0 = v332 + int32(48)
	goto L98
L108:
	;
	v361 = v359
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332)+16)) = int32(412)
	F_errmsg_internal(m, int32(_a_F_set_rel_pathlist_0), v332+int32(16))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_set_rel_pathlist_1), int32(1016), int32(_a_F_set_rel_pathlist_2))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L19
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
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = int64(1511828488474)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v412
	v414 = F_get_baserel_parampathinfo(m, l0, l1, v271)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L19
	} else {
		goto L113
	}
L113:
	;
	v416 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v407)+20)) = uint8(v416)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+16)) = v414
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+64)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v407)+24)) = v416
	*(*uint8)(unsafe.Add(mBase, uint32(v407)+21)) = uint8(v419)
	v424 = m.G0
	v426 = v424 - int32(32)
	m.G0 = v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v428 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v414 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v442 = v428 + v429<<(uint(int32(2))%32)
	goto L114
L116:
	;
	goto L117
L117:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+52))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v442 = v435 + v436<<(uint(int32(2))%32) - int32(4)
	goto L114
L118:
	;
	v448 = v414 + int32(8)
	goto L120
L119:
	;
	v448 = l1 + int32(16)
	goto L120
L120:
	;
	v449 = *(*float64)(unsafe.Add(mBase, uint32(v448)))
	*(*float64)(unsafe.Add(mBase, uint32(v407)+32)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v443)+68))
	v452 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v426)+16)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v426)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v426)+24)) = v452
	v459 = F_cost_qual_eval_walker(m, v451, v426+int32(8))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L19
	} else {
		goto L121
	}
L121:
	;
	v461 = *(*float64)(unsafe.Add(mBase, uint32(v426)+24))
	v462 = *(*float64)(unsafe.Add(mBase, uint32(v426)+16))
	if v414 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v575 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v577 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	v579 = *(*float64)(unsafe.Add(mBase, uint32(v578)+24))
	v580 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v585 = *(*float64)(unsafe.Add(mBase, uint32(v578)+16))
	v586 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v462, v461), float64(0)), v574), v585)
	*(*float64)(unsafe.Add(mBase, uint32(v407)+48)) = v586
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v407)+24))
	if v590 != 0 {
		goto L133
	} else {
		goto L134
	}
L123:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v414)+16))
	v464 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v426)+16)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v426)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v426)+24)) = v464
	if v463 == int32(0) {
		v534 = v18
		v542 = float64(0)
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v547 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v548 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v566 = v547
	v574 = v548
	goto L122
L126:
	;
	v543 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v545 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v566 = base.F64_add(v534, v543)
	v574 = base.F64_add(v542, v545)
	goto L122
L127:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v473 <= int32(0) {
		v534 = v18
		v542 = float64(0)
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v482 = int32(0)
	goto L129
L129:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v502+v482<<(uint(int32(2))%32))))
	v509 = F_cost_qual_eval_walker(m, v506, v426+int32(8))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L19
	} else {
		goto L131
	}
L130:
	;
	v515 = *(*float64)(unsafe.Add(mBase, uint32(v426)+24))
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v426)+16))
	v534 = v515
	v542 = v516
	goto L126
L131:
	;
	v512 = v482 + int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v512 < v513 {
		v482 = v512
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v591 = int64(-1)
	goto L135
L134:
	;
	v591 = int64(-262145)
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407)+40)) = base.B2i32(v580|v591 != int64(-1))
	v596 = *(*float64)(unsafe.Add(mBase, uint32(v407)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v407)+56)) = base.F64_add(v586, base.F64_add(base.F64_mul(v579, v596), base.F64_add(base.F64_mul(v575, base.F64_add(v566, v577)), float64(0))))
	m.G0 = v426 + int32(32)
	F_add_path(m, l1, v407)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v612)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v612))) = int64(1520418423066)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v612)+12)) = v617
	v619 = F_get_baserel_parampathinfo(m, l0, l1, v610)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L19
	} else {
		goto L138
	}
L138:
	;
	v621 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v612)+20)) = uint8(v621)
	*(*int32)(unsafe.Add(mBase, uint32(v612)+16)) = v619
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v612)+64)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(v612)+24)) = v621
	*(*uint8)(unsafe.Add(mBase, uint32(v612)+21)) = uint8(v624)
	v630 = m.G0
	v632 = v630 - int32(32)
	m.G0 = v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v634 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	if v619 != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v648 = v634 + v635<<(uint(int32(2))%32)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+52))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+12))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v648 = v641 + v642<<(uint(int32(2))%32) - int32(4)
	goto L139
L143:
	;
	v654 = v619 + int32(8)
	goto L145
L144:
	;
	v654 = l1 + int32(16)
	goto L145
L145:
	;
	v655 = *(*float64)(unsafe.Add(mBase, uint32(v654)))
	*(*float64)(unsafe.Add(mBase, uint32(v612)+32)) = v655
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v649)+76))
	v658 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v632)+16)) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v632)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v632)+24)) = v658
	v665 = F_cost_qual_eval_walker(m, v657, v632+int32(8))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L19
	} else {
		goto L146
	}
L146:
	;
	v667 = *(*float64)(unsafe.Add(mBase, uint32(v632)+24))
	v668 = *(*float64)(unsafe.Add(mBase, uint32(v632)+16))
	if v619 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v781 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v783 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v612)+12))
	v785 = *(*float64)(unsafe.Add(mBase, uint32(v784)+24))
	v786 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v791 = *(*float64)(unsafe.Add(mBase, uint32(v784)+16))
	v792 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v668, v667), float64(0)), v780), v791)
	*(*float64)(unsafe.Add(mBase, uint32(v612)+48)) = v792
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v612)+24))
	if v796 != 0 {
		goto L158
	} else {
		goto L159
	}
L148:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	v670 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v632)+16)) = v670
	*(*int32)(unsafe.Add(mBase, uint32(v632)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v632)+24)) = v670
	if v669 == int32(0) {
		v740 = v18
		v748 = float64(0)
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v753 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v754 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v772 = v753
	v780 = v754
	goto L147
L151:
	;
	v749 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v751 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v772 = base.F64_add(v740, v749)
	v780 = base.F64_add(v748, v751)
	goto L147
L152:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v679 <= int32(0) {
		v740 = v18
		v748 = float64(0)
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v688 = int32(0)
	goto L154
L154:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v669)+12))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v708+v688<<(uint(int32(2))%32))))
	v715 = F_cost_qual_eval_walker(m, v712, v632+int32(8))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L19
	} else {
		goto L156
	}
L155:
	;
	v721 = *(*float64)(unsafe.Add(mBase, uint32(v632)+24))
	v722 = *(*float64)(unsafe.Add(mBase, uint32(v632)+16))
	v740 = v721
	v748 = v722
	goto L151
L156:
	;
	v718 = v688 + int32(1)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v718 < v719 {
		v688 = v718
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v797 = int64(-1)
	goto L160
L159:
	;
	v797 = int64(-262145)
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v612)+40)) = base.B2i32(v786|v797 != int64(-1))
	v802 = *(*float64)(unsafe.Add(mBase, uint32(v612)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v612)+56)) = base.F64_add(v792, base.F64_add(base.F64_mul(v785, v802), base.F64_add(base.F64_mul(v781, base.F64_add(v772, v783)), float64(0))))
	m.G0 = v632 + int32(32)
	F_add_path(m, l1, v612)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v818)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v818))) = int64(1516123455770)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+12)) = v823
	v825 = F_get_baserel_parampathinfo(m, l0, l1, v816)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	v827 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v818)+20)) = uint8(v827)
	*(*int32)(unsafe.Add(mBase, uint32(v818)+16)) = v825
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+64)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v818)+24)) = v827
	*(*uint8)(unsafe.Add(mBase, uint32(v818)+21)) = uint8(v830)
	v836 = m.G0
	v838 = v836 - int32(32)
	m.G0 = v838
	if v825 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v961 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v963 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v965 = *(*float64)(unsafe.Add(mBase, uint32(v964)+24))
	v966 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v969 = *(*float64)(unsafe.Add(mBase, uint32(v964)+16))
	v970 = base.F64_add(base.F64_add(v960, float64(0)), v969)
	*(*float64)(unsafe.Add(mBase, uint32(v818)+48)) = v970
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v818)+24))
	if v974 != 0 {
		goto L175
	} else {
		goto L176
	}
L165:
	;
	v840 = *(*float64)(unsafe.Add(mBase, uint32(v825)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v818)+32)) = v840
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v825)+16))
	v843 = int32(0)
	v845 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[2]))
	v846 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v838)+16)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v838)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v838)+24)) = v846
	if v842 == v843 {
		v916 = v18
		v918 = v840
		v924 = float64(0)
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v929 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v818)+32)) = v929
	v931 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v933 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[2]))
	v934 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v952 = v931
	v954 = v929
	v955 = v933
	v960 = v934
	goto L164
L168:
	;
	v925 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v927 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v952 = base.F64_add(v916, v925)
	v954 = v918
	v955 = v845
	v960 = base.F64_add(v924, v927)
	goto L164
L169:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v842)+4))
	if v855 <= int32(0) {
		v916 = v18
		v918 = v840
		v924 = float64(0)
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v863 = v843
	goto L171
L171:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v842)+12))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883+v863<<(uint(int32(2))%32))))
	v890 = F_cost_qual_eval_walker(m, v887, v838+int32(8))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L19
	} else {
		goto L173
	}
L172:
	;
	v896 = *(*float64)(unsafe.Add(mBase, uint32(v818)+32))
	v897 = *(*float64)(unsafe.Add(mBase, uint32(v838)+24))
	v898 = *(*float64)(unsafe.Add(mBase, uint32(v838)+16))
	v916 = v897
	v918 = v896
	v924 = v898
	goto L168
L173:
	;
	v893 = v863 + int32(1)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v842)+4))
	if v893 < v894 {
		v863 = v893
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v975 = int64(-1)
	goto L177
L176:
	;
	v975 = int64(-262145)
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v818)+40)) = base.B2i32(v966|v975 != int64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(v818)+56)) = base.F64_add(v970, base.F64_add(base.F64_mul(v965, v954), base.F64_add(base.F64_mul(v961, base.F64_add(v955, base.F64_add(v952, v963))), float64(0))))
	m.G0 = v838 + int32(32)
	F_add_path(m, l1, v818)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L19
	} else {
		goto L178
	}
L178:
	;
	v2394 = l0
	v2395 = l1
	v2396 = l2
	v2397 = l3
	v2407 = v28
	goto L1
L179:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v998
	F_errmsg_internal(m, int32(_a_F_set_rel_pathlist_3), v28)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L19
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_set_rel_pathlist_4), int32(578), int32(_a_F_set_rel_pathlist_5))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L19
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	m.G0 = v1011 + int32(16)
	if v1276&int32(1) == int32(0) {
		goto L259
	} else {
		goto L260
	}
L183:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+160)))
	if v1046&int32(1) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L184:
	;
	if v1015 == int64(0) {
		v1276 = int32(0)
		goto L182
	} else {
		goto L192
	}
L185:
	;
	if v1019 == int32(0) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011)+15)))
	if (v1023|base.I32_wrap_i64(int64(base.Ui64(v1015)>>(uint(int64(4))%64))))&int32(1) == int32(0) {
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1033 = F_create_tidscan_path(m, l0, l1, v1019, v1032)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L19
	} else {
		goto L188
	}
L188:
	;
	F_add_path(m, l1, v1033)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L19
	} else {
		goto L189
	}
L189:
	;
	if v1023&int32(1) != 0 {
		v1276 = v1023
		goto L182
	} else {
		goto L190
	}
L190:
	;
	if v1015 != int64(0) {
		goto L183
	} else {
		goto L191
	}
L191:
	;
	v1276 = v1023
	goto L182
L192:
	;
	goto L183
L193:
	;
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v1257 == int32(1) {
		goto L253
	} else {
		goto L254
	}
L194:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v1051 == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1054 <= int32(0) {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v1063 = int32(0)
	v1064 = v5
	goto L197
L197:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1083+v1063<<(uint(int32(2))%32))))
	v1088 = F_IsBinaryTidClause(m, v1087, l1)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L19
	} else {
		goto L200
	}
L198:
	;
	if v1100 == int32(0) {
		goto L193
	} else {
		goto L205
	}
L199:
	;
	v1102 = v1063 + int32(1)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1102 < v1103 {
		v1063 = v1102
		v1064 = v1100
		goto L197
	} else {
		goto L204
	}
L200:
	;
	if v1088 == int32(0) {
		v1100 = v1064
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+4))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1092)+4))
	if base.Ui32(int32(3)) < base.Ui32(v1093-int32(2799)) {
		v1100 = v1064
		goto L199
	} else {
		goto L202
	}
L202:
	;
	v1098 = F_lappend(m, v1064, v1087)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L19
	} else {
		goto L203
	}
L203:
	;
	v1100 = v1098
	goto L199
L204:
	;
	goto L198
L205:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1109 = F_create_tidrangescan_path(m, l0, l1, v1100, v1107, int32(0))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	F_add_path(m, l1, v1109)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L19
	} else {
		goto L207
	}
L207:
	;
	if v1107 != 0 {
		goto L193
	} else {
		goto L208
	}
L208:
	;
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1113&int32(1) == int32(0) {
		goto L193
	} else {
		goto L209
	}
L209:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v1119 = base.F64_convert_i32_u(v1118)
	v1120 = float64(-1)
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[3]))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1125 != int32(-1) {
		v1214 = v1125
		goto L212
	} else {
		goto L213
	}
L210:
	;
	if v1224 <= int32(0) {
		goto L193
	} else {
		goto L250
	}
L211:
	;
	goto L210
L212:
	;
	if v1214 < v1122 {
		goto L247
	} else {
		goto L248
	}
L213:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1128 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1145 = int32(0)
	if base.F64_ge(v1119, float64(0)) == v1145 {
		v1177 = v1145
		goto L222
	} else {
		goto L223
	}
L215:
	;
	if base.F64_ge(v1119, float64(0)) != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	if base.F64_lt(v1119, base.F64_convert_i32_s(v1133)) != 0 {
		v1224 = int32(0)
		goto L211
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	if base.F64_ge(v1120, float64(0)) == int32(0) {
		goto L214
	} else {
		goto L220
	}
L219:
	;
	goto L218
L220:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[5]))
	if base.F64_lt(v1120, base.F64_convert_i32_s(v1142)) != 0 {
		v1224 = int32(0)
		goto L211
	} else {
		goto L221
	}
L221:
	;
	goto L214
L222:
	;
	if base.F64_ge(v1120, float64(0)) == int32(0) {
		v1214 = v1177
		goto L212
	} else {
		goto L231
	}
L223:
	;
	v1150 = int32(1)
	v1152 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	if v1152 <= v1150 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1155 = v1150
	goto L226
L225:
	;
	v1155 = v1152
	goto L226
L226:
	;
	v1157 = v1155
	v1161 = int32(1)
	goto L227
L227:
	;
	v1164 = v1157 * int32(3)
	if base.F64_ge(v1119, base.F64_convert_i32_u(v1164)) == int32(0) {
		v1177 = v1161
		goto L222
	} else {
		goto L229
	}
L228:
	;
	v1177 = v1170
	goto L222
L229:
	;
	v1170 = v1161 + int32(1)
	if v1164 < int32(715827883) {
		v1157 = v1164
		v1161 = v1170
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v1183 = int32(1)
	v1185 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[5]))
	if v1185 <= v1183 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1188 = v1183
	goto L234
L233:
	;
	v1188 = v1185
	goto L234
L234:
	;
	v1190 = v1188
	v1195 = int32(1)
	goto L235
L235:
	;
	v1197 = v1190 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v1197), v1120) != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	if v1177 < v1204 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	v1201 = v1195 + int32(1)
	if v1197 < int32(715827883) {
		v1190 = v1197
		v1195 = v1201
		goto L235
	} else {
		goto L240
	}
L238:
	;
	v1204 = v1195
	goto L239
L239:
	;
	goto L236
L240:
	;
	v1204 = v1201
	goto L239
L241:
	;
	v1206 = v1177
	goto L243
L242:
	;
	v1206 = v1204
	goto L243
L243:
	;
	if int32(0) < v1177 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1209 = v1206
	goto L246
L245:
	;
	v1209 = v1204
	goto L246
L246:
	;
	v1214 = v1209
	goto L212
L247:
	;
	v1217 = v1214
	goto L249
L248:
	;
	v1217 = v1122
	goto L249
L249:
	;
	v1224 = v1217
	goto L211
L250:
	;
	v1228 = F_create_tidrangescan_path(m, l0, l1, v1100, int32(0), v1224)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L19
	} else {
		goto L251
	}
L251:
	;
	F_add_partial_path(m, l1, v1228)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L19
	} else {
		goto L252
	}
L252:
	;
	goto L193
L253:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v1263 = F_generate_implied_equalities_for_column(m, l0, l1, int32(874), int32(0), v1262)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L19
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	F_BuildParameterizedTidPaths(m, l0, l1, v1267)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L19
	} else {
		goto L258
	}
L256:
	;
	F_BuildParameterizedTidPaths(m, l0, l1, v1263)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L19
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v1276 = int32(0)
	goto L182
L259:
	;
	v1304 = F_create_seqscan_path(m, l0, l1, v1008, int32(0))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L19
	} else {
		goto L262
	}
L260:
	;
	v2369 = l0
	v2370 = l1
	v2371 = l2
	v2372 = l3
	v2382 = v28
	goto L261
L261:
	;
	v2394 = v2369
	v2395 = v2370
	v2396 = v2371
	v2397 = v2372
	v2407 = v2382
	goto L1
L262:
	;
	F_add_path(m, l1, v1304)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L19
	} else {
		goto L263
	}
L263:
	;
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if base.B2i32(v1308 != int32(1))|v1008 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1423 = m.G0
	v1425 = v1423 - int32(416)
	m.G0 = v1425
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v1427 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L265:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[3]))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1314 != int32(-1) {
		v1369 = v1314
		goto L266
	} else {
		goto L267
	}
L266:
	;
	if v1369 < v1313 {
		goto L276
	} else {
		goto L277
	}
L267:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1321 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	if base.B2i32(v1317 == int32(0))&base.B2i32(base.I64_extend_i32_u(v1323) < base.I64_extend_i32_s(v1321)) != 0 {
		goto L264
	} else {
		goto L268
	}
L268:
	;
	v1327 = int32(1)
	if v1321 <= v1327 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1330 = v1327
	goto L271
L270:
	;
	v1330 = v1321
	goto L271
L271:
	;
	v1336 = v1330
	v1337 = int32(1)
	goto L272
L272:
	;
	v1358 = v1336 * int32(3)
	if base.Ui32(v1323) < base.Ui32(v1358) {
		v1369 = v1337
		goto L266
	} else {
		goto L274
	}
L273:
	;
	v1369 = v1361
	goto L266
L274:
	;
	v1361 = v1337 + int32(1)
	if v1358 < int32(715827883) {
		v1336 = v1358
		v1337 = v1361
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v1390 = v1369
	goto L278
L277:
	;
	v1390 = v1313
	goto L278
L278:
	;
	if v1390 <= int32(0) {
		goto L264
	} else {
		goto L279
	}
L279:
	;
	v1394 = F_create_seqscan_path(m, l0, l1, int32(0), v1390)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L19
	} else {
		goto L280
	}
L280:
	;
	F_add_partial_path(m, l1, v1394)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L19
	} else {
		goto L281
	}
L281:
	;
	goto L264
L282:
	;
	m.G0 = v1425 + int32(416)
	v2369 = l0
	v2370 = l1
	v2371 = l2
	v2372 = l3
	v2382 = v28
	goto L261
L283:
	;
	v1430 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+404)) = v1430
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+400)) = v1430
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+4))
	if v1430 < v1434 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1450 = v5
	v1453 = v5
	goto L287
L285:
	;
	v1908 = v5
	v1924 = int32(0)
	goto L286
L286:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v1927 = F_generate_bitmap_or_paths(m, l0, l1, v1925, int32(0))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L19
	} else {
		goto L346
	}
L287:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+12))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1466+v1453<<(uint(int32(2))%32))))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+88))
	if v1471 != 0 {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+404))
	v1908 = v1877
	v1924 = v1897
	goto L286
L289:
	;
	v1894 = v1453 + int32(1)
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+4))
	if v1894 < v1895 {
		v1450 = v1877
		v1453 = v1894
		goto L287
	} else {
		goto L345
	}
L290:
	;
	v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470)+100)))
	if v1472 != int32(1) {
		v1877 = v1450
		goto L289
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v1477 = int32(0)
	base.MemoryFill(m, v1425+int32(268), v1477, int32(132))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+96))
	if v1480 == v1477 {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	goto L292
L294:
	;
	F_get_index_paths(m, l0, l1, v1470, v1425+int32(268), v1425+int32(404))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L19
	} else {
		goto L301
	}
L295:
	;
	v1483 = int32(0)
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+4))
	if v1484 <= v1483 {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1491 = v1483
	goto L297
L297:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+12))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1512+v1491<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v1516, v1470, v1425+int32(268))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L19
	} else {
		goto L299
	}
L298:
	;
	goto L294
L299:
	;
	v1522 = v1491 + int32(1)
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+4))
	if v1522 < v1523 {
		v1491 = v1522
		goto L297
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	v1558 = int32(0)
	base.MemoryFill(m, v1425+int32(136), v1558, int32(132))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v1561 == v1558 {
		v1624 = v1450
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1642 = int32(0)
	base.MemoryFill(m, v1425+int32(4), v1642, int32(132))
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+12))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1646)+232)))
	if v1647 != int32(1) {
		v1787 = v1642
		goto L318
	} else {
		goto L319
	}
L303:
	;
	v1564 = int32(0)
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+4))
	if v1565 <= v1564 {
		v1624 = v1450
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1572 = v1564
	v1577 = v1450
	goto L305
L305:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+12))
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1593+v1572<<(uint(int32(2))%32))))
	v1598 = F_join_clause_is_movable_to(m, v1597, l1)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L19
	} else {
		goto L307
	}
L306:
	;
	v1624 = v1610
	goto L302
L307:
	;
	if v1598 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1597)+52))
	goto L311
L309:
	;
	v1610 = v1577
	goto L310
L310:
	;
	v1612 = v1572 + int32(1)
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+4))
	if v1612 < v1613 {
		v1572 = v1612
		v1577 = v1610
		goto L305
	} else {
		goto L317
	}
L311:
	;
	if v1600 != int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1603 = F_list_append_unique_ptr(m, v1577, v1597)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L19
	} else {
		goto L315
	}
L313:
	;
	v1605 = v1577
	goto L314
L314:
	;
	F_match_clause_to_index(m, l0, v1597, v1470, v1425+int32(136))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L19
	} else {
		goto L316
	}
L315:
	;
	v1605 = v1603
	goto L314
L316:
	;
	v1610 = v1605
	goto L310
L317:
	;
	goto L306
L318:
	;
	v1788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425)+136)))
	if v1788|v1787&int32(1) == int32(0) {
		v1877 = v1624
		goto L289
	} else {
		goto L332
	}
L319:
	;
	v1650 = int32(0)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+40))
	if v1651 <= v1650 {
		v1787 = v1650
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1662 = v1650
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+412)) = v1662
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+408)) = v1470
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+12))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+112))
	v1686 = F_generate_implied_equalities_for_column(m, l0, v1681, int32(870), v1425+int32(408), v1685)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L19
	} else {
		goto L324
	}
L322:
	;
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425)+4)))
	v1787 = v1761
	goto L318
L323:
	;
	v1758 = v1662 + int32(1)
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+40))
	if v1758 < v1759 {
		v1662 = v1758
		goto L321
	} else {
		goto L331
	}
L324:
	;
	if v1686 == int32(0) {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v1690 = int32(0)
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1686)+4))
	if v1691 <= v1690 {
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v1698 = v1690
	goto L327
L327:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1686)+12))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1719+v1698<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v1723, v1470, v1425+int32(4))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L19
	} else {
		goto L329
	}
L328:
	;
	goto L323
L329:
	;
	v1729 = v1698 + int32(1)
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1686)+4))
	if v1729 < v1730 {
		v1698 = v1729
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	goto L322
L332:
	;
	v1794 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+408)) = v1794
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+40))
	if v1798 <= v1794 {
		v1877 = v1624
		goto L289
	} else {
		goto L333
	}
L333:
	;
	v1805 = v1794
	v1809 = v1794
	goto L334
L334:
	;
	v1828 = v1805 << (uint(int32(2)) % 32)
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1425+int32(140)+v1828)))
	if v1830 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	v1877 = v1624
	goto L289
L336:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1830)+4))
	v1832 = v1831
	goto L338
L337:
	;
	v1832 = int32(0)
	goto L338
L338:
	;
	v1841 = v1832 + v1809
	F_consider_index_join_outer_rels(m, l0, l1, v1470, v1425+int32(268), v1425+int32(136), v1425+int32(4), v1425+int32(400), v1830, v1841, v1425+int32(408))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L19
	} else {
		goto L339
	}
L339:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1425+int32(8)+v1828)))
	if v1847 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	v1850 = v1848
	goto L342
L341:
	;
	v1850 = int32(0)
	goto L342
L342:
	;
	v1859 = v1841 + v1850
	F_consider_index_join_outer_rels(m, l0, l1, v1470, v1425+int32(268), v1425+int32(136), v1425+int32(4), v1425+int32(400), v1847, v1859, v1425+int32(408))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L19
	} else {
		goto L343
	}
L343:
	;
	v1865 = v1805 + int32(1)
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+40))
	if v1865 < v1866 {
		v1805 = v1865
		v1809 = v1859
		goto L334
	} else {
		goto L344
	}
L344:
	;
	goto L335
L345:
	;
	goto L288
L346:
	;
	v1929 = F_list_concat(m, v1924, v1927)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L19
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+404)) = v1929
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v1933 = F_generate_bitmap_or_paths(m, l0, l1, v1908, v1932)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L19
	} else {
		goto L348
	}
L348:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+400))
	v1936 = F_list_concat(m, v1935, v1933)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L19
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+400)) = v1936
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+404))
	if v1939 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	if v1936 == int32(0) {
		goto L282
	} else {
		goto L379
	}
L351:
	;
	v1942 = F_choose_bitmap_and(m, l0, l1, v1939)
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L19
	} else {
		goto L352
	}
L352:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1947 = F_create_bitmap_heap_path(m, l0, l1, v1942, v1944, float64(1), int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L19
	} else {
		goto L353
	}
L353:
	;
	F_add_path(m, l1, v1947)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L19
	} else {
		goto L354
	}
L354:
	;
	v1951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1951 != int32(1) {
		goto L350
	} else {
		goto L355
	}
L355:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v1954 != 0 {
		goto L350
	} else {
		goto L356
	}
L356:
	;
	v1956 = int32(0)
	v1958 = F_compute_bitmap_pages(m, l0, l1, v1942, float64(1), v1956, v1956)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L19
	} else {
		goto L357
	}
L357:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[3]))
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1962 != int32(-1) {
		v2030 = v1962
		goto L359
	} else {
		goto L360
	}
L358:
	;
	goto L350
L359:
	;
	if v2030 < v1961 {
		goto L373
	} else {
		goto L374
	}
L360:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1968 = int32(0)
	if v1965|base.B2i32(base.F64_ge(v1958, float64(0)) == v1968) == v1968 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	if base.F64_lt(v1958, base.F64_convert_i32_s(v1974)) != 0 {
		goto L358
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1977 = int32(0)
	if base.F64_ge(v1958, float64(0)) == v1977 {
		v2030 = v1977
		goto L359
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	v1982 = int32(1)
	v1984 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	if v1984 <= v1982 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1987 = v1982
	goto L368
L367:
	;
	v1987 = v1984
	goto L368
L368:
	;
	v1995 = int32(1)
	v1997 = v1987
	goto L369
L369:
	;
	v2015 = v1997 * int32(3)
	if base.F64_ge(v1958, base.F64_convert_i32_u(v2015)) == int32(0) {
		v2030 = v1995
		goto L359
	} else {
		goto L371
	}
L370:
	;
	v2030 = v2021
	goto L359
L371:
	;
	v2021 = v1995 + int32(1)
	if v2015 < int32(715827883) {
		v1995 = v2021
		v1997 = v2015
		goto L369
	} else {
		goto L372
	}
L372:
	;
	goto L370
L373:
	;
	v2050 = v2030
	goto L375
L374:
	;
	v2050 = v1961
	goto L375
L375:
	;
	if v2050 <= int32(0) {
		goto L358
	} else {
		goto L376
	}
L376:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v2055 = F_create_bitmap_heap_path(m, l0, l1, v1942, v2053, float64(1), v2050)
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L19
	} else {
		goto L377
	}
L377:
	;
	F_add_partial_path(m, l1, v2055)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L19
	} else {
		goto L378
	}
L378:
	;
	goto L358
L379:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+4))
	if v2111 <= int32(0) {
		goto L282
	} else {
		goto L380
	}
L380:
	;
	v2114 = int32(0)
	v2120 = v2114
	v2122 = v2114
	goto L381
L381:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+12))
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2141+v2120<<(uint(int32(2))%32))))
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2145)+16))
	if v2146 != 0 {
		goto L383
	} else {
		goto L384
	}
L382:
	;
	if v2150 == int32(0) {
		goto L282
	} else {
		goto L388
	}
L383:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+4))
	v2149 = v2147
	goto L385
L384:
	;
	v2149 = int32(0)
	goto L385
L385:
	;
	v2150 = F_list_append_unique(m, v2122, v2149)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L19
	} else {
		goto L386
	}
L386:
	;
	v2153 = v2120 + int32(1)
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+4))
	if v2153 < v2154 {
		v2120 = v2153
		v2122 = v2150
		goto L381
	} else {
		goto L387
	}
L387:
	;
	goto L382
L388:
	;
	v2158 = int32(0)
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v2150)+4))
	if v2159 <= v2158 {
		goto L282
	} else {
		goto L389
	}
L389:
	;
	v2171 = v2158
	goto L390
L390:
	;
	v2187 = int32(0)
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+400))
	if v2188 == v2187 {
		v2302 = v2187
		goto L392
	} else {
		goto L393
	}
L391:
	;
	goto L282
L392:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+404))
	v2321 = F_list_concat(m, v2302, v2320)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L19
	} else {
		goto L419
	}
L393:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+4))
	if v2191 <= int32(0) {
		v2302 = v2187
		goto L392
	} else {
		goto L394
	}
L394:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2150)+12))
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v2194+v2171<<(uint(int32(2))%32))))
	v2204 = int32(0)
	v2207 = v2187
	goto L395
L395:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+12))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2225+v2204<<(uint(int32(2))%32))))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2229)+16))
	if v2230 != 0 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v2302 = v2290
	goto L392
L397:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+4))
	v2233 = v2231
	goto L399
L398:
	;
	v2233 = int32(0)
	goto L399
L399:
	;
	v2234 = int32(0)
	if v2233 == v2234 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	if v2287 != 0 {
		goto L414
	} else {
		goto L415
	}
L401:
	;
	v2287 = int32(1)
	goto L400
L402:
	;
	goto L403
L403:
	;
	if v2198 == int32(0) {
		v2280 = v2234
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v2287 = v2280
	goto L400
L405:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2233)+4))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2198)+4))
	if v2244 < v2243 {
		v2280 = v2234
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v2246 = int32(1)
	if v2243 <= v2246 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v2249 = v2246
	goto L409
L408:
	;
	v2249 = v2243
	goto L409
L409:
	;
	v2250 = int32(8)
	v2255 = int32(0)
	goto L410
L410:
	;
	v2262 = v2255 << (uint(int32(2)) % 32)
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2233+v2250+v2262)))
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v2198+v2250+v2262)))
	v2269 = v2264 & (v2266 ^ int32(-1))
	v2271 = base.B2i32(v2269 == int32(0))
	if v2269 != 0 {
		v2280 = v2271
		goto L404
	} else {
		goto L412
	}
L411:
	;
	v2280 = v2271
	goto L404
L412:
	;
	v2273 = v2255 + int32(1)
	if v2273 != v2249 {
		v2255 = v2273
		goto L410
	} else {
		goto L413
	}
L413:
	;
	goto L411
L414:
	;
	v2288 = F_lappend(m, v2207, v2229)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L19
	} else {
		goto L417
	}
L415:
	;
	v2290 = v2207
	goto L416
L416:
	;
	v2292 = v2204 + int32(1)
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+4))
	if v2292 < v2293 {
		v2204 = v2292
		v2207 = v2290
		goto L395
	} else {
		goto L418
	}
L417:
	;
	v2290 = v2288
	goto L416
L418:
	;
	goto L396
L419:
	;
	v2323 = F_choose_bitmap_and(m, l0, l1, v2321)
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L19
	} else {
		goto L420
	}
L420:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+16))
	if v2325 != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+4))
	v2328 = v2326
	goto L423
L422:
	;
	v2328 = int32(0)
	goto L423
L423:
	;
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v2330 = F_get_loop_count(m, l0, v2329, v2328)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L19
	} else {
		goto L424
	}
L424:
	;
	v2333 = F_create_bitmap_heap_path(m, l0, l1, v2323, v2328, v2330, int32(0))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L19
	} else {
		goto L425
	}
L425:
	;
	F_add_path(m, l1, v2333)
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L19
	} else {
		goto L426
	}
L426:
	;
	v2338 = v2171 + int32(1)
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2150)+4))
	if v2338 < v2339 {
		v2171 = v2338
		goto L390
	} else {
		goto L427
	}
L427:
	;
	goto L391
L428:
	;
	m.T0[v2420].(func(*base.Module, int32, int32, int32, int32))(m, v2394, v2395, v2396, v2397)
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L19
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+4))
	if v2423 != 0 {
		goto L432
	} else {
		goto L433
	}
L431:
	;
	goto L430
L432:
	;
	F_set_cheapest(m, v2395)
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L19
	} else {
		goto L447
	}
L433:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+8))
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v2394)+60))
	v2426 = int32(0)
	if base.B2i32(v2424 == v2426)|base.B2i32(v2425 == v2426) != 0 {
		v2472 = base.B2i32(v2424|v2425 == v2426)
		goto L435
	} else {
		goto L436
	}
L434:
	;
	if v2472 != 0 {
		goto L432
	} else {
		goto L445
	}
L435:
	;
	goto L434
L436:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v2424)+4))
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2425)+4))
	if v2440 != v2441 {
		v2472 = int32(0)
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v2443 = int32(1)
	if v2440 <= v2443 {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v2446 = v2443
	goto L440
L439:
	;
	v2446 = v2440
	goto L440
L440:
	;
	v2447 = int32(8)
	v2452 = int32(0)
	goto L441
L441:
	;
	v2460 = v2452 << (uint(int32(2)) % 32)
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2424+v2447+v2460)))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2425+v2447+v2460)))
	v2465 = base.B2i32(v2462 == v2464)
	if v2462 != v2464 {
		v2472 = v2465
		goto L435
	} else {
		goto L443
	}
L442:
	;
	v2472 = v2465
	goto L435
L443:
	;
	v2468 = v2452 + int32(1)
	if v2468 != v2446 {
		v2452 = v2468
		goto L441
	} else {
		goto L444
	}
L444:
	;
	goto L442
L445:
	;
	F_generate_useful_gather_paths(m, v2394, v2395, int32(0))
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L19
	} else {
		goto L446
	}
L446:
	;
	goto L432
L447:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2394)+152))
	if v2482 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	m.G0 = v2407 + int32(16)
	return
L449:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v2394)+156))
	if v2485 == int32(0) {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+240))
	if v2488 == int32(0) {
		goto L448
	} else {
		goto L451
	}
L451:
	;
	F_generate_grouped_paths(m, v2394, v2488, v2395)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L19
	} else {
		goto L452
	}
L452:
	;
	F_set_cheapest(m, v2488)
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L19
	} else {
		goto L453
	}
L453:
	;
	goto L448
}
func F_set_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v24 float64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v682 int32
	_ = v682
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 float64
	_ = v793
	var v794 float64
	_ = v794
	var v795 float64
	_ = v795
	var v796 float64
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 float64
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 float64
	_ = v898
	var v900 float64
	_ = v900
	var v930 int32
	_ = v930
	var v931 float64
	_ = v931
	var v932 float64
	_ = v932
	var v933 float64
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 float64
	_ = v997
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1011 float64
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1027 int32
	_ = v1027
	var v1052 int32
	_ = v1052
	var v1059 float64
	_ = v1059
	var v1094 int32
	_ = v1094
	var v1124 int32
	_ = v1124
	var v1126 int64
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1138 int64
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1167 int64
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1179 int64
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 float64
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1332 int32
	_ = v1332
	var v1340 int32
	_ = v1340
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1568 int32
	_ = v1568
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int64
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1615 int32
	_ = v1615
	var v1640 int32
	_ = v1640
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1682 int64
	_ = v1682
	var v1684 int64
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1696 float64
	_ = v1696
	var v1697 float64
	_ = v1697
	var v1706 float64
	_ = v1706
	var v1710 float64
	_ = v1710
	var v1712 float64
	_ = v1712
	var v1714 float64
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 float64
	_ = v1777
	var v1780 float64
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1804 float64
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1809 float64
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1814 float64
	_ = v1814
	var v1815 int64
	_ = v1815
	var v1822 int32
	_ = v1822
	var v1828 int32
	_ = v1828
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1867 float64
	_ = v1867
	var v1868 float64
	_ = v1868
	var v1870 float64
	_ = v1870
	var v1894 float64
	_ = v1894
	var v1895 float64
	_ = v1895
	var v1900 float64
	_ = v1900
	var v1901 float64
	_ = v1901
	var v1903 float64
	_ = v1903
	var v1905 float64
	_ = v1905
	var v1907 float64
	_ = v1907
	var v1909 float64
	_ = v1909
	var v1910 float64
	_ = v1910
	var v1935 float64
	_ = v1935
	var v1936 float64
	_ = v1936
	var v1937 float64
	_ = v1937
	var v1940 float64
	_ = v1940
	var v1941 float64
	_ = v1941
	var v1942 int64
	_ = v1942
	var v1944 float64
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1949 int64
	_ = v1949
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2075 float64
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2116 int32
	_ = v2116
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2247 float64
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2287 int32
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2303 int32
	_ = v2303
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2319 int32
	_ = v2319
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2335 int32
	_ = v2335
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2361 float64
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2396 int32
	_ = v2396
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 float64
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 float64
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2479 int32
	_ = v2479
	var v2499 int32
	_ = v2499
	var v2505 int32
	_ = v2505
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2549 int32
	_ = v2549
	var v2557 int32
	_ = v2557
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2641 int32
	_ = v2641
	var v2666 int32
	_ = v2666
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2749 int32
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2769 int32
	_ = v2769
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2782 int32
	_ = v2782
	var v2785 float64
	_ = v2785
	var v2786 float64
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2812 int32
	_ = v2812
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2827 int32
	_ = v2827
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2836 int64
	_ = v2836
	var v2840 int32
	_ = v2840
	var v2847 int32
	_ = v2847
	var v2848 int64
	_ = v2848
	var v2850 int32
	_ = v2850
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2862 int32
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2890 int32
	_ = v2890
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2924 int32
	_ = v2924
	var v2928 int32
	_ = v2928
	var v2937 int32
	_ = v2937
	var v2959 int32
	_ = v2959
	var v2962 int32
	_ = v2962
	var v2969 int32
	_ = v2969
	var v2995 int32
	_ = v2995
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3009 int32
	_ = v3009
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3050 int32
	_ = v3050
	var v3057 int32
	_ = v3057
	var v3083 int32
	_ = v3083
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	v5 = int32(0)
	v24 = float64(0)
	v30 = m.G0
	v32 = v30 - int32(208)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v32 + int32(208)
	return
L2:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v68 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v35 = F_relation_excluded_by_constraints(m, l0, l1, l3)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v35 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+200)) = v39
	v41 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v39
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v32)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v55
	v62 = F_create_append_path(m, v39, l1, v32, v39, v52, v39, v39, float64(-1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_add_path(m, l1, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	switch v1157 {
	case 0:
		goto L213
	case 1:
		goto L212
	default:
		goto L210
	case 3:
		goto L204
	case 4:
		goto L205
	case 5:
		goto L206
	case 6:
		goto L207
	case 7:
		goto L208
	case 8:
		goto L209
	}
L13:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_rel_size[0])))
	if v74 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+90)))
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v91 = v89 - v90
	v96 = F_palloc0(m, v91<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L19
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v77 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v78 != int32(112) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81-v82<<(uint(int32(2))%32))))
	if v86 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)) = uint8(v87)
	goto L14
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v98 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v1124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+200)) = v1124
	v1126 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v1126
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v1126
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+32)) = v1124
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v1124
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v1124
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1138 = *(*int64)(unsafe.Add(mBase, uint32(v32)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v1138
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v1140
	v1149 = F_create_append_path(m, v1124, l1, v32+int32(16), v1124, v1137, v1124, v1124, float64(-1))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L4
	} else {
		goto L197
	}
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v101 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v111 = v5
	v126 = v5
	v127 = v24
	v128 = v24
	v129 = v24
	goto L23
L23:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v111<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+184)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v139 != l2 {
		v930 = v126
		v931 = v127
		v932 = v128
		v933 = v129
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v930 == int32(0) {
		goto L20
	} else {
		goto L186
	}
L25:
	;
	v938 = v111 + int32(1)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v938 < v939 {
		v111 = v938
		v126 = v930
		v127 = v931
		v128 = v932
		v129 = v933
		goto L23
	} else {
		goto L185
	}
L26:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141+v142<<(uint(int32(2))%32))))
	v147 = F_find_base_rel(m, l0, v142)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v149 = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	if v151 == v149 {
		v172 = v149
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v172 != 0 {
		v930 = v126
		v931 = v127
		v932 = v128
		v933 = v129
		goto L25
	} else {
		goto L38
	}
L29:
	;
	goto L28
L30:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v155 = v154
	goto L31
L31:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if base.Ui32(int32(2)) <= base.Ui32(v159-int32(303)) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v172 = int32(1)
	goto L29
L33:
	;
	if v159 != int32(293) {
		v172 = v149
		goto L29
	} else {
		goto L36
	}
L34:
	;
	v155 = v158 + int32(72)
	goto L31
L35:
	;
	goto L32
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)+72))
	if v166 != 0 {
		v172 = v149
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v173 = F_relation_excluded_by_constraints(m, l0, v147, v146)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v173 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+200)) = v175
	v177 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v147)+16)) = v177
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v147)+52)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v147)+44)) = v175
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v32)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v191
	v200 = F_create_append_path(m, v175, v147, v32+int32(32), v175, v188, v175, v175, float64(-1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v206 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	F_add_path(m, v147, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_set_cheapest(m, v147)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v930 = v126
	v931 = v127
	v932 = v128
	v933 = v129
	goto L25
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+228)) = v316
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v346 = F_adjust_appendrel_attrs(m, l0, v342, int32(1), v32+int32(184))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L72
	}
L47:
	;
	v316 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v210 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v212 <= v210 {
		v316 = v210
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v219 = v210
	v220 = v210
	goto L51
L51:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v219<<(uint(int32(2))%32))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+28))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v251 = int32(0)
	if base.B2i32(v249 == v251)|base.B2i32(v250 == v251) != 0 {
		v296 = v251
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v316 = v306
	goto L46
L53:
	;
	if v296 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	goto L53
L55:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v261 < v262 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v264 = v261
	goto L58
L57:
	;
	v264 = v262
	goto L58
L58:
	;
	if v264 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v267 = int32(1)
	goto L61
L60:
	;
	v267 = v264
	goto L61
L61:
	;
	v268 = int32(8)
	v273 = int32(0)
	goto L62
L62:
	;
	v280 = v273 << (uint(int32(2)) % 32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v250+v268+v280)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v249+v268+v280)))
	v285 = v282 & v284
	v287 = base.B2i32(v285 != int32(0))
	if v285 != 0 {
		v296 = v287
		goto L54
	} else {
		goto L64
	}
L63:
	;
	v296 = v287
	goto L54
L64:
	;
	v289 = v273 + int32(1)
	if v289 != v267 {
		v273 = v289
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v302 = F_adjust_appendrel_attrs(m, l0, v248, int32(1), v32+int32(184))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	v306 = v220
	goto L68
L68:
	;
	v308 = v219 + int32(1)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v308 < v309 {
		v219 = v308
		v220 = v306
		goto L51
	} else {
		goto L71
	}
L69:
	;
	v304 = F_lappend(m, v220, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v306 = v304
	goto L68
L71:
	;
	goto L52
L72:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v348)+4)) = v346
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v350 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+232)) = uint8(v746)
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)))
	if v748 == int32(1) {
		goto L144
	} else {
		goto L145
	}
L74:
	;
	v354 = int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v355 != 0 {
		v360 = v354
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v32)+184))
	v364 = m.G0
	v366 = v364 - int32(16)
	m.G0 = v366
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v363
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v147)+252))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v371 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	if v360 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L78:
	;
	goto L77
L79:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v356 != 0 {
		v360 = v354
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v360 = base.B2i32(v357 != int32(0))
	goto L78
L81:
	;
	goto L76
L82:
	;
	if int32(0) <= v428 {
		goto L93
	} else {
		goto L94
	}
L83:
	;
	v428 = base.I32_ctz(v414) | v415<<(uint(int32(5))%32)
	goto L82
L84:
	;
	v428 = int32(-2)
	goto L82
L85:
	;
	v379 = int32(0)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v382 <= v379 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v385 = v371 + int32(8)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	v392 = v389 & int32(-1)
	if v392 != 0 {
		v414 = v392
		v415 = v379
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v393 = int32(1)
	if v393 == v382 {
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v397 = v393
	goto L89
L89:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v385+v397<<(uint(int32(2))%32))))
	if v404 != 0 {
		v414 = v404
		v415 = v397
		goto L83
	} else {
		goto L91
	}
L90:
	;
	goto L84
L91:
	;
	v406 = v397 + int32(1)
	if v406 != v382 {
		v397 = v406
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v440 = v428
	goto L96
L94:
	;
	goto L95
L95:
	;
	m.G0 = v366 + int32(16)
	goto L73
L96:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v461+v440<<(uint(int32(2))%32))))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+41)))
	if v466 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L95
L98:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v626 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L99:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v465)+16))
	if v467 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v470 = int32(0)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	if v471 <= v470 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v478 = v470
	goto L102
L102:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v467)+12))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v503+v478<<(uint(int32(2))%32))))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+12)))
	if v508 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L98
L104:
	;
	v594 = v478 + int32(1)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	if v594 < v595 {
		v478 = v594
		goto L102
	} else {
		goto L131
	}
L105:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v510 = int32(0)
	if v509 == v510 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v563 == int32(0) {
		goto L104
	} else {
		goto L120
	}
L107:
	;
	v563 = int32(1)
	goto L106
L108:
	;
	goto L109
L109:
	;
	if v370 == int32(0) {
		v556 = v510
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v563 = v556
	goto L106
L111:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v520 < v519 {
		v556 = v510
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v522 = int32(1)
	if v519 <= v522 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v525 = v522
	goto L115
L114:
	;
	v525 = v519
	goto L115
L115:
	;
	v526 = int32(8)
	v531 = int32(0)
	goto L116
L116:
	;
	v538 = v531 << (uint(int32(2)) % 32)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v509+v526+v538)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v370+v526+v538)))
	v545 = v540 & (v542 ^ int32(-1))
	v547 = base.B2i32(v545 == int32(0))
	if v545 != 0 {
		v556 = v547
		goto L110
	} else {
		goto L118
	}
L117:
	;
	v556 = v547
	goto L110
L118:
	;
	v549 = v531 + int32(1)
	if v549 != v525 {
		v531 = v549
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	if v566 == int32(0) {
		goto L104
	} else {
		goto L121
	}
L121:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v570 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v583 = F_bms_difference(m, v582, v370)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L128
	}
L123:
	;
	v576 = F_adjust_appendrel_attrs(m, l0, v569, int32(1), v366+int32(12))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v147)+248))
	v579 = F_adjust_appendrel_attrs_multilevel(m, l0, v569, v147, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L4
	} else {
		goto L127
	}
L126:
	;
	v581 = v576
	goto L122
L127:
	;
	v581 = v579
	goto L122
L128:
	;
	v585 = F_bms_add_members(m, v583, v369)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v147)+76))
	F_add_child_eq_member(m, l0, v465, v440, v581, v585, v587, v507, v588, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	goto L104
L131:
	;
	goto L103
L132:
	;
	if int32(0) <= v682 {
		v440 = v682
		goto L96
	} else {
		goto L143
	}
L133:
	;
	v682 = base.I32_ctz(v668) | v669<<(uint(int32(5))%32)
	goto L132
L134:
	;
	v682 = int32(-2)
	goto L132
L135:
	;
	v633 = v440 + int32(1)
	v635 = int32(base.Ui32(v633) >> (uint(int32(5)) % 32))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	if v636 <= v635 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v639 = v626 + int32(8)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v639+v635<<(uint(int32(2))%32))))
	v646 = v643 & (int32(-1) << (uint(v633) % 32))
	if v646 != 0 {
		v668 = v646
		v669 = v635
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v648 = v635 + int32(1)
	if v648 == v636 {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v651 = v648
	goto L139
L139:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v639+v651<<(uint(int32(2))%32))))
	if v658 != 0 {
		v668 = v658
		v669 = v651
		goto L133
	} else {
		goto L141
	}
L140:
	;
	goto L134
L141:
	;
	v660 = v651 + int32(1)
	if v660 != v636 {
		v651 = v660
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	goto L97
L144:
	;
	v751 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+233)) = uint8(v751)
	goto L146
L145:
	;
	goto L146
L146:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+94)))
	if v754 != int32(1) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_set_rel_size(m, l0, v147, v142, v146)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L151
	}
L148:
	;
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v757 != int32(1) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	F_set_rel_consider_parallel(m, l0, v147, v146)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	goto L147
L151:
	;
	v764 = int32(0)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	if v766 == v764 {
		v787 = v764
		goto L153
	} else {
		goto L154
	}
L152:
	;
	if v787 != 0 {
		v930 = v126
		v931 = v127
		v932 = v128
		v933 = v129
		goto L25
	} else {
		goto L162
	}
L153:
	;
	goto L152
L154:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v766)+12))
	v770 = v769
	goto L155
L155:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	if base.Ui32(int32(2)) <= base.Ui32(v774-int32(303)) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v787 = int32(1)
	goto L153
L157:
	;
	if v774 != int32(293) {
		v787 = v764
		goto L153
	} else {
		goto L160
	}
L158:
	;
	v770 = v773 + int32(72)
	goto L155
L159:
	;
	goto L156
L160:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v773)+72))
	if v781 != 0 {
		v787 = v764
		goto L153
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+26)))
	if v788 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v791 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v791)
	goto L165
L164:
	;
	goto L165
L165:
	;
	v793 = *(*float64)(unsafe.Add(mBase, uint32(v147)+16))
	v794 = base.F64_add(v127, v793)
	v795 = *(*float64)(unsafe.Add(mBase, uint32(v147)+128))
	v796 = base.F64_add(v128, v795)
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+32))
	v801 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v798), v793), v129)
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v797)+4))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	v810 = int32(0)
	goto L166
L166:
	;
	v835 = int32(0)
	if v804 == v835 {
		v845 = v835
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v846 = int32(1)
	if v802 == int32(0) {
		v930 = v846
		v931 = v794
		v932 = v796
		v933 = v801
		goto L25
	} else {
		goto L171
	}
L169:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v839 <= v810 {
		v845 = int32(0)
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v804)+12))
	v845 = v841 + v810<<(uint(int32(2))%32)
	goto L168
L171:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v802)+4))
	if base.B2i32(v845 == int32(0))|base.B2i32(v851 <= v810) != 0 {
		v930 = v846
		v931 = v794
		v932 = v796
		v933 = v801
		goto L25
	} else {
		goto L172
	}
L172:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v802)+12))
	if v854 == int32(0) {
		v930 = v846
		v931 = v794
		v932 = v796
		v933 = v801
		goto L25
	} else {
		goto L173
	}
L173:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v857)))
	if v858 != int32(6) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v810 = v810 + int32(1)
	goto L166
L175:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	if v861 != l2 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v863 = int32(*(*int16)(unsafe.Add(mBase, uint32(v857)+8)))
	v864 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v854+v810<<(uint(int32(2))%32))))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v869)))
	if v870 != int32(6) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v896 = v96 + (v863-v864)<<(uint(int32(3))%32)
	v898 = *(*float64)(unsafe.Add(mBase, uint32(v147)+16))
	v900 = *(*float64)(unsafe.Add(mBase, uint32(v896)))
	*(*float64)(unsafe.Add(mBase, uint32(v896))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v893), v898), v900)
	goto L174
L178:
	;
	v887 = F_exprType(m, v869)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L4
	} else {
		goto L182
	}
L179:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v869)+4))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v147)+76))
	if v873 != v874 {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v147)+96))
	v877 = int32(*(*int16)(unsafe.Add(mBase, uint32(v869)+8)))
	v878 = int32(*(*int16)(unsafe.Add(mBase, uint32(v147)+88)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v876+(v877-v878)<<(uint(int32(2))%32))))
	if int32(0) < v883 {
		v893 = v883
		goto L177
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	v889 = F_exprTypmod(m, v869)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v891 = F_get_typavgwidth(m, v887, v889)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v893 = v891
	goto L177
L185:
	;
	goto L24
L186:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v931
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v932
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v945)+32)) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v933, v931)))
	v950 = int32(0)
	if v91 < v950 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	F_pfree(m, v96)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L4
	} else {
		goto L196
	}
L188:
	;
	if v89 != v90 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v954 = int32(1)
	v955 = v91 + v954
	v964 = int32(0)
	v965 = v950
	goto L192
L190:
	;
	v1027 = v950
	goto L191
L191:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v1059 = *(*float64)(unsafe.Add(mBase, uint32(v96+v1027<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1052+v1027<<(uint(int32(2))%32)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v1059, v931)))
	goto L187
L192:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v991 = int32(2)
	v994 = int32(3)
	v997 = *(*float64)(unsafe.Add(mBase, uint32(v96+v965<<(uint(v994)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v990+v965<<(uint(v991)%32)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v997, v931)))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v1004 = v965 | int32(1)
	v1011 = *(*float64)(unsafe.Add(mBase, uint32(v96+v1004<<(uint(v994)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1002+v1004<<(uint(v991)%32)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v1011, v931)))
	v1017 = v965 + v991
	v1019 = v964 + v991
	if v1019 != v955&int32(-2) {
		v964 = v1019
		v965 = v1017
		goto L192
	} else {
		goto L194
	}
L193:
	;
	if v955&v954 == int32(0) {
		goto L187
	} else {
		goto L195
	}
L194:
	;
	goto L193
L195:
	;
	v1027 = v1017
	goto L191
L196:
	;
	goto L1
L197:
	;
	F_add_path(m, l1, v1149)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	F_pfree(m, v96)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	goto L1
L201:
	;
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(v32)+184))
	F_pfree(m, v2530)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L4
	} else {
		goto L456
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v2479
	v2505 = v2499
	goto L201
L203:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v32)+180))
	v2479 = v1568
	v2499 = v2469
	goto L202
L204:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2369 != 0 {
		goto L442
	} else {
		goto L443
	}
L205:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+128)) = int64(4636737291354636288)
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L4
	} else {
		goto L440
	}
L206:
	;
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2341 != 0 {
		goto L433
	} else {
		goto L434
	}
L207:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+92)))
	if v1967 != int32(1) {
		goto L368
	} else {
		goto L369
	}
L208:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1761 != 0 {
		goto L339
	} else {
		goto L340
	}
L209:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+128)) = int64(4607182418800017408)
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L4
	} else {
		goto L333
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L4
	} else {
		goto L330
	}
L211:
	;
	v1592 = m.G0
	v1594 = v1592 - int32(32)
	m.G0 = v1594
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(4652007308841189376)
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v1599 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1594)+16)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(v1594)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v1594)+24)) = v1599
	v1605 = v1594 + int32(16)
	if v1598 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L212:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+180)) = int32(0)
	v1231 = F_copyObjectImpl(m, v1228)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L4
	} else {
		goto L228
	}
L213:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	switch v1158 - int32(102) {
	case 0:
		goto L211
	default:
		goto L214
	case 10:
		goto L215
	}
L214:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v1197 != 0 {
		goto L219
	} else {
		goto L220
	}
L215:
	;
	v1161 = m.G0
	v1163 = v1161 - int32(32)
	m.G0 = v1163
	v1165 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+24)) = v1165
	v1167 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1163)+16)) = v1167
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v1167
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+32)) = v1165
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v1165
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v1165
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1179 = *(*int64)(unsafe.Add(mBase, uint32(v1163)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1163))) = v1179
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+8)) = v1181
	v1188 = F_create_append_path(m, v1165, l1, v1163, v1165, v1178, v1165, v1165, float64(-1))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	F_add_path(m, l1, v1188)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	m.G0 = v1163 + int32(32)
	goto L1
L219:
	;
	v1198 = m.G0
	v1200 = v1198 - int32(16)
	m.G0 = v1200
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	F_check_index_predicates(m, l0, l1)
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	F_check_index_predicates(m, l0, l1)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L4
	} else {
		goto L226
	}
L222:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+4))
	v1206 = F_GetTsmRoutine(m, v1205)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+8))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+12))
	m.T0[v1211].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, v1208, v1200+int32(12), v1200)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v1214
	v1216 = *(*float64)(unsafe.Add(mBase, uint32(v1200)))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v1216
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	m.G0 = v1200 + int32(16)
	goto L1
L226:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L4
	} else {
		goto L227
	}
L227:
	;
	goto L1
L228:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+184)) = int64(0)
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+76))
	if v1236 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1236)+4))
	v1241 = v1237 + int32(1)
	goto L231
L230:
	;
	v1241 = int32(1)
	goto L231
L231:
	;
	v1242 = F_palloc0(m, v1241)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+184)) = v1242
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+189)) = uint8(v1245)
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v1247 == int32(0) {
		v2505 = v5
		goto L201
	} else {
		goto L233
	}
L233:
	;
	v1252 = F_subquery_is_pushdown_safe(m, v1231, v1231, v32+int32(184))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	if v1252 == int32(0) {
		v2505 = v5
		goto L201
	} else {
		goto L235
	}
L235:
	;
	v1256 = int32(0)
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v1257 == v1256 {
		v2479 = v5
		v2499 = v1256
		goto L202
	} else {
		goto L236
	}
L236:
	;
	v1260 = int32(0)
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+4))
	if v1261 <= v1260 {
		v2479 = v5
		v2499 = v1260
		goto L202
	} else {
		goto L237
	}
L237:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v32)+184))
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+188)))
	v1266 = int32(1)
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+189)))
	v1275 = v5
	v1280 = v5
	goto L238
L238:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+12))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1300+v1275<<(uint(int32(2))%32))))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+10)))
	if v1305 != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L203
L240:
	;
	v1589 = v1275 + int32(1)
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+4))
	if v1589 < v1590 {
		v1275 = v1589
		v1280 = v1568
		goto L238
	} else {
		goto L313
	}
L241:
	;
	v1557 = F_lappend(m, v1280, v1304)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L4
	} else {
		goto L312
	}
L242:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+4))
	v1307 = F_contain_subplans(m, v1306)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	if v1307 != 0 {
		goto L241
	} else {
		goto L244
	}
L244:
	;
	if v1265&v1266 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1309 = F_contain_volatile_functions(m, v1304)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	if v1268&v1266 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	if v1309 != 0 {
		goto L241
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v1311 = F_contain_leaked_vars(m, v1306)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L4
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1314 = F_pull_var_clause(m, v1306, int32(16))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L4
	} else {
		goto L257
	}
L253:
	;
	if v1311 != 0 {
		goto L241
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+37)))
	if v1496 != 0 {
		goto L294
	} else {
		goto L295
	}
L256:
	;
	F_list_free(m, v1314)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L4
	} else {
		goto L292
	}
L257:
	;
	if v1314 == int32(0) {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+4))
	if v1318 <= int32(0) {
		goto L256
	} else {
		goto L259
	}
L259:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+12))
	v1332 = int32(0)
	v1340 = int32(1)
	goto L261
L260:
	;
	F_list_free(m, v1314)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L4
	} else {
		goto L291
	}
L261:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1321+v1332<<(uint(int32(2))%32))))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)))
	if v1357 != int32(6) {
		goto L260
	} else {
		goto L263
	}
L262:
	;
	F_list_free(m, v1314)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L4
	} else {
		goto L271
	}
L263:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+4))
	if v1360 != l2 {
		goto L260
	} else {
		goto L264
	}
L264:
	;
	v1362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1356)+8)))
	if v1362 == int32(0) {
		goto L260
	} else {
		goto L265
	}
L265:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1362+v1264))))
	if v1366 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	if v1366&int32(23) != 0 {
		goto L260
	} else {
		goto L269
	}
L267:
	;
	v1370 = v1340
	goto L268
L268:
	;
	v1372 = v1332 + int32(1)
	if v1318 != v1372 {
		v1332 = v1372
		v1340 = v1370
		goto L261
	} else {
		goto L270
	}
L269:
	;
	v1370 = int32(2)
	goto L268
L270:
	;
	goto L262
L271:
	;
	if v1370 == int32(1) {
		goto L255
	} else {
		goto L272
	}
L272:
	;
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+37)))
	if v1378 != int32(1) {
		goto L241
	} else {
		goto L273
	}
L273:
	;
	v1381 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+192)) = uint8(v1381)
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	if v1383 != int32(17) {
		goto L241
	} else {
		goto L274
	}
L274:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+28))
	if v1386 == int32(0) {
		goto L241
	} else {
		goto L275
	}
L275:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+4))
	if v1389 != int32(2) {
		goto L241
	} else {
		goto L276
	}
L276:
	;
	F_set_opfuncid(m, v1306)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+8))
	v1395 = F_func_strict(m, v1394)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	if v1395 == int32(0) {
		goto L241
	} else {
		goto L279
	}
L279:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+28))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1399)+12))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1400)))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	if v1402 != int32(6) {
		v1427 = v1400
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+192)))
	if v1457&int32(1) != 0 {
		goto L241
	} else {
		goto L290
	}
L281:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+4))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	if v1430 != int32(6) {
		goto L241
	} else {
		goto L286
	}
L282:
	;
	v1405 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1401)+8)))
	if v1405 <= int32(0) {
		v1427 = v1400
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+76))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+12))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1409+v1405<<(uint(int32(2))%32)-int32(4))))
	v1416 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1415)+8)))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+4))
	v1423 = F_find_window_run_conditions(m, v1231, v1416, v1417, v1306, int32(1), v32+int32(192), v32+int32(180))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	if v1423 != 0 {
		goto L280
	} else {
		goto L285
	}
L285:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+28))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+12))
	v1427 = v1426
	goto L281
L286:
	;
	v1433 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1429)+8)))
	if v1433 <= int32(0) {
		goto L241
	} else {
		goto L287
	}
L287:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+76))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+12))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1437+v1433<<(uint(int32(2))%32)-int32(4))))
	v1444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1443)+8)))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+4))
	v1451 = F_find_window_run_conditions(m, v1231, v1444, v1445, v1306, int32(0), v32+int32(192), v32+int32(180))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L4
	} else {
		goto L288
	}
L288:
	;
	if v1451 == int32(0) {
		goto L241
	} else {
		goto L289
	}
L289:
	;
	goto L280
L290:
	;
	v1568 = v1280
	goto L240
L291:
	;
	goto L241
L292:
	;
	goto L255
L293:
	;
	F_subquery_push_qual(m, v1231, l3, l2, v1306)
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L4
	} else {
		goto L311
	}
L294:
	;
	v1523 = F_expression_has_grouping_conflict(m, v1306, int32(868), v1231)
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L4
	} else {
		goto L309
	}
L295:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+120))
	if v1497 != 0 {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+144))
	if v1498 == int32(0) {
		goto L293
	} else {
		goto L297
	}
L297:
	;
	if v1498 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	if v1518 == int32(0) {
		goto L293
	} else {
		goto L308
	}
L299:
	;
	goto L298
L300:
	;
	v1518 = int32(0)
	goto L299
L301:
	;
	v1504 = v1498
	goto L302
L302:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1504)))
	if v1506 != int32(142) {
		goto L300
	} else {
		goto L304
	}
L303:
	;
	goto L300
L304:
	;
	v1509 = int32(1)
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+32))
	if v1510 != 0 {
		v1518 = v1509
		goto L299
	} else {
		goto L305
	}
L305:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+12))
	v1512 = F_setop_has_grouping(m, v1511)
	mBase = m.M
	if v1512 != 0 {
		v1518 = v1509
		goto L299
	} else {
		goto L306
	}
L306:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+16))
	if v1513 != 0 {
		v1504 = v1513
		goto L302
	} else {
		goto L307
	}
L307:
	;
	goto L303
L308:
	;
	goto L294
L309:
	;
	if v1523 != 0 {
		goto L241
	} else {
		goto L310
	}
L310:
	;
	goto L293
L311:
	;
	v1568 = v1280
	goto L240
L312:
	;
	v1568 = v1557
	goto L240
L313:
	;
	goto L239
L314:
	;
	v1682 = *(*int64)(unsafe.Add(mBase, uint32(v1605)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+216)) = v1682
	v1684 = *(*int64)(unsafe.Add(mBase, uint32(v1605)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+208)) = v1684
	F_set_rel_width(m, l0, l1)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L4
	} else {
		goto L321
	}
L315:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+4))
	if v1608 <= int32(0) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1615 = v5
	goto L317
L317:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+12))
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1640+v1615<<(uint(int32(2))%32))))
	v1647 = F_cost_qual_eval_walker(m, v1644, v1594+int32(8))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L4
	} else {
		goto L319
	}
L318:
	;
	goto L314
L319:
	;
	v1650 = v1615 + int32(1)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+4))
	if v1650 < v1651 {
		v1615 = v1650
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	m.G0 = v1594 + int32(32)
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+4))
	m.T0[v1693].(func(*base.Module, int32, int32, int32))(m, l0, l1, v1691)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L4
	} else {
		goto L322
	}
L322:
	;
	v1696 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v1697 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1696)&int64(9223372036854775807)))|base.F64_gt(v1696, v1697) != 0 {
		v1710 = v1697
		goto L324
	} else {
		goto L325
	}
L323:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v1710
	v1712 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	if base.F64_lt(v1710, v1712) != 0 {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	goto L323
L325:
	;
	v1706 = float64(1)
	if base.F64_le(v1696, v1706) != 0 {
		v1710 = v1706
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1710 = base.F64_nearest(v1696)
	goto L324
L327:
	;
	v1714 = v1712
	goto L329
L328:
	;
	v1714 = v1710
	goto L329
L329:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v1714
	goto L1
L330:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v1720
	F_errmsg_internal(m, int32(_a_F_set_rel_size_0), v32+int32(48))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L4
	} else {
		goto L331
	}
L331:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(504), int32(_a_F_set_rel_size_2))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L4
	} else {
		goto L332
	}
L332:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L333:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1738 = F_palloc0(m, int32(72))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1738)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1738))) = int64(1438814044442)
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1738)+12)) = v1743
	v1745 = F_get_baserel_parampathinfo(m, l0, l1, v1736)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	v1747 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1738)+20)) = uint8(v1747)
	*(*int32)(unsafe.Add(mBase, uint32(v1738)+16)) = v1745
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1738)+64)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v1738)+24)) = v1747
	*(*uint8)(unsafe.Add(mBase, uint32(v1738)+21)) = uint8(v1750)
	F_cost_resultscan(m, v1738, l0, l1, v1745)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L4
	} else {
		goto L336
	}
L336:
	;
	F_add_path(m, l1, v1738)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L4
	} else {
		goto L337
	}
L337:
	;
	goto L1
L338:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1775)))
	v1777 = *(*float64)(unsafe.Add(mBase, uint32(v1776)+112))
	if base.F64_lt(v1777, float64(0)) != 0 {
		goto L342
	} else {
		goto L343
	}
L339:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1775 = v1761 + v1762<<(uint(int32(2))%32)
	goto L338
L340:
	;
	goto L341
L341:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1766)+52))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1767)+12))
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1775 = v1768 + v1769<<(uint(int32(2))%32) - int32(4)
	goto L338
L342:
	;
	v1780 = float64(1000)
	goto L344
L343:
	;
	v1780 = v1777
	goto L344
L344:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v1780
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1786 = F_palloc0(m, int32(72))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1786)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1786))) = int64(1529008357658)
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1786)+12)) = v1791
	v1793 = F_get_baserel_parampathinfo(m, l0, l1, v1784)
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	v1795 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1786)+20)) = uint8(v1795)
	*(*int32)(unsafe.Add(mBase, uint32(v1786)+16)) = v1793
	v1798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1786)+64)) = v1795
	*(*int32)(unsafe.Add(mBase, uint32(v1786)+24)) = v1795
	*(*uint8)(unsafe.Add(mBase, uint32(v1786)+21)) = uint8(v1798)
	v1804 = float64(0)
	v1805 = m.G0
	v1807 = v1805 - int32(32)
	m.G0 = v1807
	if v1793 != 0 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1941 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v1942 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v1944 = base.F64_add(v1940, float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(v1786)+48)) = v1944
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1786)+24))
	if v1948 != 0 {
		goto L359
	} else {
		goto L360
	}
L349:
	;
	v1809 = *(*float64)(unsafe.Add(mBase, uint32(v1793)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1786)+32)) = v1809
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1793)+16))
	v1812 = int32(0)
	v1814 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_size[1]))
	v1815 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1807)+16)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v1807)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v1807)+24)) = v1815
	if v1811 == v1812 {
		v1894 = v1804
		v1895 = v24
		v1900 = v1814
		goto L352
	} else {
		goto L353
	}
L350:
	;
	goto L351
L351:
	;
	v1905 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1786)+32)) = v1905
	v1907 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v1909 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_size[1]))
	v1910 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v1935 = v1907
	v1936 = v1909
	v1937 = v1909
	v1940 = v1910
	goto L348
L352:
	;
	v1901 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v1903 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v1935 = base.F64_add(v1895, v1901)
	v1936 = v1814
	v1937 = v1900
	v1940 = base.F64_add(v1894, v1903)
	goto L348
L353:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+4))
	if v1822 <= int32(0) {
		v1894 = v1804
		v1895 = v24
		v1900 = v1814
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v1828 = v1812
	goto L355
L355:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+12))
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1854+v1828<<(uint(int32(2))%32))))
	v1861 = F_cost_qual_eval_walker(m, v1858, v1807+int32(8))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L4
	} else {
		goto L357
	}
L356:
	;
	v1867 = *(*float64)(unsafe.Add(mBase, uint32(v1807)+24))
	v1868 = *(*float64)(unsafe.Add(mBase, uint32(v1807)+16))
	v1870 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_size[1]))
	v1894 = v1868
	v1895 = v1867
	v1900 = v1870
	goto L352
L357:
	;
	v1864 = v1828 + int32(1)
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+4))
	if v1864 < v1865 {
		v1828 = v1864
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	v1949 = int64(-1)
	goto L361
L360:
	;
	v1949 = int64(-262145)
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1786)+40)) = base.B2i32(v1942|v1949 != int64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(v1786)+56)) = base.F64_add(v1944, base.F64_add(base.F64_mul(v1941, base.F64_add(v1936, base.F64_add(v1935, v1937))), float64(0)))
	m.G0 = v1807 + int32(32)
	F_add_path(m, l1, v1786)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L4
	} else {
		goto L362
	}
L362:
	;
	goto L1
L363:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L4
	} else {
		goto L429
	}
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L4
	} else {
		goto L426
	}
L365:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L4
	} else {
		goto L423
	}
L366:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L4
	} else {
		goto L420
	}
L367:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+4))
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+48))
	if v2103 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L368:
	;
	v1972 = l0
	v1974 = v1966
	goto L371
L369:
	;
	goto L370
L370:
	;
	if v1966 == int32(0) {
		goto L366
	} else {
		goto L378
	}
L371:
	;
	if v1974 == int32(0) {
		goto L367
	} else {
		goto L373
	}
L372:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L4
	} else {
		goto L375
	}
L373:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+16))
	if v2003 != 0 {
		v1972 = v2003
		v1974 = v1974 - int32(1)
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+176)) = v2008
	F_errmsg_internal(m, int32(_a_F_set_rel_size_3), v32+int32(176))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3075), int32(_a_F_set_rel_size_4))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L4
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	v2024 = l0
	v2026 = v1966
	goto L380
L379:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+364))
	if v2072 == int32(0) {
		goto L365
	} else {
		goto L387
	}
L380:
	;
	v2052 = v2026 - int32(1)
	if v2052 == int32(0) {
		goto L379
	} else {
		goto L382
	}
L381:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L4
	} else {
		goto L384
	}
L382:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+16))
	if v2055 != 0 {
		v2024 = v2055
		v2026 = v2052
		goto L380
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v2060
	F_errmsg_internal(m, int32(_a_F_set_rel_size_3), v32+int32(112))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3207), int32(_a_F_set_rel_size_5))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L387:
	;
	v2075 = *(*float64)(unsafe.Add(mBase, uint32(v2072)+32))
	F_set_cte_size_estimates(m, l0, l1, v2075)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v2080 = F_palloc0(m, int32(72))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2080)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2080))) = int64(1533303324954)
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2080)+12)) = v2085
	v2087 = F_get_baserel_parampathinfo(m, l0, l1, v2078)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	v2089 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2080)+20)) = uint8(v2089)
	*(*int32)(unsafe.Add(mBase, uint32(v2080)+16)) = v2087
	v2092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v2080)+64)) = v2089
	*(*int32)(unsafe.Add(mBase, uint32(v2080)+24)) = v2089
	*(*uint8)(unsafe.Add(mBase, uint32(v2080)+21)) = uint8(v2092)
	F_cost_ctescan(m, v2080, l0, l1, v2087)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	F_add_path(m, l1, v2080)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	goto L1
L393:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+84))
	if v2222 == int32(0) {
		goto L364
	} else {
		goto L411
	}
L394:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L4
	} else {
		goto L408
	}
L395:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+4))
	if v2106 <= int32(0) {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+12))
	v2116 = int32(0)
	goto L397
L397:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v2110+v2116<<(uint(int32(2))%32))))
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2144)+4))
	v2148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2145))))
	v2151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109))))
	if base.B2i32(v2148 == int32(0))|base.B2i32(v2148 != v2151) != 0 {
		v2169 = v2148
		v2170 = v2151
		goto L400
	} else {
		goto L401
	}
L398:
	;
	goto L394
L399:
	;
	if v2169-v2170 == int32(0) {
		goto L393
	} else {
		goto L406
	}
L400:
	;
	goto L399
L401:
	;
	v2154 = v2145
	v2155 = v2109
	goto L402
L402:
	;
	v2158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2155)+1)))
	v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2154)+1)))
	if v2159 == int32(0) {
		v2169 = v2159
		v2170 = v2158
		goto L400
	} else {
		goto L404
	}
L403:
	;
	v2169 = v2159
	v2170 = v2158
	goto L400
L404:
	;
	v2162 = int32(1)
	if v2159 == v2158 {
		v2154 = v2154 + v2162
		v2155 = v2155 + v2162
		goto L402
	} else {
		goto L405
	}
L405:
	;
	goto L403
L406:
	;
	v2175 = v2116 + int32(1)
	if v2106 != v2175 {
		v2116 = v2175
		goto L397
	} else {
		goto L407
	}
L407:
	;
	goto L398
L408:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v2210
	F_errmsg_internal(m, int32(_a_F_set_rel_size_6), v32+int32(128))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L4
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3093), int32(_a_F_set_rel_size_4))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L4
	} else {
		goto L410
	}
L410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L411:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2222)+4))
	if v2225 <= v2116 {
		goto L364
	} else {
		goto L412
	}
L412:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2222)+12))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2227+v2116<<(uint(int32(2))%32))))
	if v2231 <= int32(0) {
		goto L363
	} else {
		goto L413
	}
L413:
	;
	v2237 = v2231<<(uint(int32(2))%32) - int32(4)
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+12))
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2239)+12))
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2237+v2240)))
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+8))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2243)+12))
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2244+v2237)))
	v2247 = *(*float64)(unsafe.Add(mBase, uint32(v2246)+24))
	F_set_cte_size_estimates(m, l0, l1, v2247)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L4
	} else {
		goto L414
	}
L414:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+64))
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2246)+44))
	v2252 = F_convert_subquery_pathkeys(m, l0, l1, v2250, v2251)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L4
	} else {
		goto L415
	}
L415:
	;
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v2256 = F_palloc0(m, int32(72))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L4
	} else {
		goto L416
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2256))) = int64(1524713390362)
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+12)) = v2261
	v2263 = F_get_baserel_parampathinfo(m, l0, l1, v2254)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L4
	} else {
		goto L417
	}
L417:
	;
	v2265 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2256)+20)) = uint8(v2265)
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+16)) = v2263
	v2268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+64)) = v2252
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+24)) = v2265
	*(*uint8)(unsafe.Add(mBase, uint32(v2256)+21)) = uint8(v2268)
	F_cost_ctescan(m, v2256, l0, l1, v2263)
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L4
	} else {
		goto L418
	}
L418:
	;
	F_add_path(m, l1, v2256)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L4
	} else {
		goto L419
	}
L419:
	;
	goto L1
L420:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v2281
	F_errmsg_internal(m, int32(_a_F_set_rel_size_3), v32+int32(80))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L4
	} else {
		goto L421
	}
L421:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3200), int32(_a_F_set_rel_size_5))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L4
	} else {
		goto L422
	}
L422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L423:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v2297
	F_errmsg_internal(m, int32(_a_F_set_rel_size_7), v32+int32(96))
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L4
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3211), int32(_a_F_set_rel_size_5))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L426:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v2313
	F_errmsg_internal(m, int32(_a_F_set_rel_size_8), v32+int32(144))
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L4
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3095), int32(_a_F_set_rel_size_4))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L4
	} else {
		goto L428
	}
L428:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L429:
	;
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+160)) = v2329
	F_errmsg_internal(m, int32(_a_F_set_rel_size_9), v32+int32(160))
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L4
	} else {
		goto L430
	}
L430:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3098), int32(_a_F_set_rel_size_4))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L4
	} else {
		goto L431
	}
L431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L432:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2355)))
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2356)+80))
	if v2357 != 0 {
		goto L436
	} else {
		goto L437
	}
L433:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v2355 = v2341 + v2342<<(uint(int32(2))%32)
	goto L432
L434:
	;
	goto L435
L435:
	;
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2346)+52))
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v2347)+12))
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v2355 = v2348 + v2349<<(uint(int32(2))%32) - int32(4)
	goto L432
L436:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+4))
	v2361 = base.F64_convert_i32_s(v2358)
	goto L438
L437:
	;
	v2361 = float64(0)
	goto L438
L438:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v2361
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L4
	} else {
		goto L439
	}
L439:
	;
	goto L1
L440:
	;
	goto L1
L441:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2383)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+128)) = int64(0)
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2384)+68))
	if v2387 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L442:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v2383 = v2369 + v2370<<(uint(int32(2))%32)
	goto L441
L443:
	;
	goto L444
L444:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v2374)+52))
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2375)+12))
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v2383 = v2376 + v2377<<(uint(int32(2))%32) - int32(4)
	goto L441
L445:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L4
	} else {
		goto L455
	}
L446:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2387)+4))
	if v2390 <= int32(0) {
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v2396 = int32(0)
	goto L448
L448:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v2387)+12))
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2423+v2396<<(uint(int32(2))%32))))
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v2427)+4))
	v2429 = F_expression_returns_set_rows(m, l0, v2428)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L4
	} else {
		goto L450
	}
L449:
	;
	goto L445
L450:
	;
	v2431 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	if base.F64_gt(v2429, v2431) != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+128)) = v2429
	goto L453
L452:
	;
	goto L453
L453:
	;
	v2435 = v2396 + int32(1)
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2387)+4))
	if v2435 < v2436 {
		v2396 = v2435
		goto L448
	} else {
		goto L454
	}
L454:
	;
	goto L449
L455:
	;
	goto L1
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+192)) = v2505
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+144))
	if v2534 != 0 {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v2730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+36)))
	if v2730 != 0 {
		v2786 = v24
		goto L494
	} else {
		goto L495
	}
L458:
	;
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+120))
	if v2535 != 0 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v2536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+40)))
	if v2536 != int32(1) {
		goto L457
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2539)+4))
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	F_pull_varattnos(m, v2540, v2541, v32+int32(192))
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		goto L4
	} else {
		goto L463
	}
L462:
	;
	goto L461
L463:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v2546 == int32(0) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v2628 = F_bms_is_member(m, int32(7), v2627)
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L4
	} else {
		goto L471
	}
L465:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v2546)+4))
	if v2549 <= int32(0) {
		goto L464
	} else {
		goto L466
	}
L466:
	;
	v2557 = int32(0)
	goto L467
L467:
	;
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v2546)+12))
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2582+v2557<<(uint(int32(2))%32))))
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2586)+4))
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	F_pull_varattnos(m, v2587, v2588, v32+int32(192))
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L4
	} else {
		goto L469
	}
L468:
	;
	goto L464
L469:
	;
	v2594 = v2557 + int32(1)
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(v2546)+4))
	if v2594 < v2595 {
		v2557 = v2594
		goto L467
	} else {
		goto L470
	}
L470:
	;
	goto L468
L471:
	;
	if v2628 != 0 {
		goto L457
	} else {
		goto L472
	}
L472:
	;
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+76))
	if v2630 == int32(0) {
		goto L457
	} else {
		goto L473
	}
L473:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2630)+4))
	if v2633 <= int32(0) {
		goto L457
	} else {
		goto L474
	}
L474:
	;
	v2641 = int32(0)
	goto L475
L475:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v2630)+12))
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2666+v2641<<(uint(int32(2))%32))))
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v2670)+16))
	if v2671 != 0 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	goto L457
L477:
	;
	v2698 = v2641 + int32(1)
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v2630)+4))
	if v2698 < v2699 {
		v2641 = v2698
		goto L475
	} else {
		goto L493
	}
L478:
	;
	v2672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2670)+26)))
	if v2672 != 0 {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v2670)+4))
	v2674 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2670)+8)))
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v2678 = F_bms_is_member(m, v2674+int32(7), v2677)
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L4
	} else {
		goto L480
	}
L480:
	;
	if v2678 != 0 {
		goto L477
	} else {
		goto L481
	}
L481:
	;
	v2680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+38)))
	if v2680 == int32(1) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2683 = F_expression_returns_set(m, v2673)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L4
	} else {
		goto L485
	}
L483:
	;
	goto L484
L484:
	;
	v2685 = F_contain_volatile_functions(m, v2673)
	mBase = m.M
	v2686 = m.ExcPending
	if v2686 != 0 {
		goto L4
	} else {
		goto L487
	}
L485:
	;
	if v2683 != 0 {
		goto L477
	} else {
		goto L486
	}
L486:
	;
	goto L484
L487:
	;
	if v2685 != 0 {
		goto L477
	} else {
		goto L488
	}
L488:
	;
	v2687 = F_exprType(m, v2673)
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L4
	} else {
		goto L489
	}
L489:
	;
	v2689 = F_exprTypmod(m, v2673)
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L4
	} else {
		goto L490
	}
L490:
	;
	v2691 = F_exprCollation(m, v2673)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L4
	} else {
		goto L491
	}
L491:
	;
	v2693 = F_makeNullConst(m, v2687, v2689, v2691)
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L4
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2670)+4)) = v2693
	goto L477
L493:
	;
	goto L476
L494:
	;
	v2787 = int32(0)
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2789)+4))
	v2792 = F_choose_plan_name(m, v2788, v2790, v2787)
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L4
	} else {
		goto L518
	}
L495:
	;
	v2731 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+100))
	if v2731 != 0 {
		v2786 = v24
		goto L494
	} else {
		goto L496
	}
L496:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+108))
	if v2732 != 0 {
		v2786 = v24
		goto L494
	} else {
		goto L497
	}
L497:
	;
	v2733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+334)))
	if v2733 != 0 {
		v2786 = v24
		goto L494
	} else {
		goto L498
	}
L498:
	;
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+120))
	if v2734 != 0 {
		v2786 = v24
		goto L494
	} else {
		goto L499
	}
L499:
	;
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+124))
	if v2735 != 0 {
		v2786 = v24
		goto L494
	} else {
		goto L500
	}
L500:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2737 = int32(0)
	if v2736 == v2737 {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	if v2782 == int32(2) {
		v2786 = v24
		goto L494
	} else {
		goto L517
	}
L502:
	;
	v2782 = int32(0)
	goto L501
L503:
	;
	goto L504
L504:
	;
	v2745 = int32(1)
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2736)+4))
	if v2746 <= v2745 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v2749 = v2745
	goto L507
L506:
	;
	v2749 = v2746
	goto L507
L507:
	;
	v2753 = int32(0)
	v2755 = v2737
	goto L508
L508:
	;
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2736+int32(8)+v2753<<(uint(int32(2))%32))))
	if v2762 != 0 {
		goto L511
	} else {
		goto L512
	}
L509:
	;
	v2782 = v2774
	goto L501
L510:
	;
	goto L509
L511:
	;
	v2763 = int32(2)
	if v2755 != 0 {
		v2774 = v2763
		goto L510
	} else {
		goto L514
	}
L512:
	;
	v2769 = v2755
	goto L513
L513:
	;
	v2771 = v2753 + int32(1)
	if v2771 != v2749 {
		v2753 = v2771
		v2755 = v2769
		goto L508
	} else {
		goto L516
	}
L514:
	;
	v2764 = int32(1)
	if base.Ui32(v2764) < base.Ui32(base.I32_popcnt(v2762)) {
		v2774 = v2763
		goto L510
	} else {
		goto L515
	}
L515:
	;
	v2769 = v2764
	goto L513
L516:
	;
	v2774 = v2769
	goto L510
L517:
	;
	v2785 = *(*float64)(unsafe.Add(mBase, uint32(l0)+312))
	v2786 = v2785
	goto L494
L518:
	;
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2795 = int32(0)
	v2798 = F_subquery_planner(m, v2794, v1231, v2792, l0, v2795, v2795, v2786, v2795)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L4
	} else {
		goto L519
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+148)) = v2798
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+152)) = v2801
	v2803 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2803
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v2808 = F_fetch_upper_rel(m, v2805, int32(7), v2803)
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L4
	} else {
		goto L520
	}
L520:
	;
	v2810 = int32(0)
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v2808)+44))
	if v2812 == v2810 {
		v2833 = v2810
		goto L522
	} else {
		goto L523
	}
L521:
	;
	if v2833 != 0 {
		goto L531
	} else {
		goto L532
	}
L522:
	;
	goto L521
L523:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v2812)+12))
	v2816 = v2815
	goto L524
L524:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(v2816)))
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2819)))
	if base.Ui32(int32(2)) <= base.Ui32(v2820-int32(303)) {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	v2833 = int32(1)
	goto L522
L526:
	;
	if v2820 != int32(293) {
		v2833 = v2810
		goto L522
	} else {
		goto L529
	}
L527:
	;
	v2816 = v2819 + int32(72)
	goto L524
L528:
	;
	goto L525
L529:
	;
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v2819)+72))
	if v2827 != 0 {
		v2833 = v2810
		goto L522
	} else {
		goto L530
	}
L530:
	;
	goto L528
L531:
	;
	v2834 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+200)) = v2834
	v2836 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v2836
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v2836
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2840)+32)) = v2834
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v2834
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v2834
	v2847 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v2848 = *(*int64)(unsafe.Add(mBase, uint32(v32)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+64)) = v2848
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v2850
	v2859 = F_create_append_path(m, v2834, l1, v32-int32(-64), v2834, v2847, v2834, v2834, float64(-1))
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L4
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	F_set_subquery_size_estimates(m, l0, l1)
	mBase = m.M
	v2866 = m.ExcPending
	if v2866 != 0 {
		goto L4
	} else {
		goto L537
	}
L534:
	;
	F_add_path(m, l1, v2859)
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L4
	} else {
		goto L535
	}
L535:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L4
	} else {
		goto L536
	}
L536:
	;
	goto L1
L537:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2867)+4))
	if v2868 != 0 {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+4))
	v2870 = v2869
	goto L540
L539:
	;
	v2870 = v2787
	goto L540
L540:
	;
	v2871 = int32(0)
	v2873 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+76))
	if v2873 != 0 {
		goto L542
	} else {
		goto L543
	}
L541:
	;
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2808)+44))
	if v2959 == int32(0) {
		goto L558
	} else {
		goto L559
	}
L542:
	;
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+4))
	v2876 = v2874
	goto L544
L543:
	;
	v2876 = int32(0)
	goto L544
L544:
	;
	if v2876 != v2870 {
		v2937 = v2871
		goto L541
	} else {
		goto L545
	}
L545:
	;
	if v2868 == int32(0) {
		goto L546
	} else {
		goto L547
	}
L546:
	;
	v2937 = int32(1)
	goto L541
L547:
	;
	goto L548
L548:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+4))
	v2882 = int32(0)
	if v2882 < v2881 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v2885 = v2881
	goto L551
L550:
	;
	v2885 = v2882
	goto L551
L551:
	;
	v2890 = v2871
	goto L552
L552:
	;
	v2915 = base.B2i32(v2890 == v2885)
	if v2890 == v2885 {
		v2937 = v2915
		goto L541
	} else {
		goto L554
	}
L553:
	;
	v2937 = v2915
	goto L541
L554:
	;
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+12))
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2916+v2890<<(uint(int32(2))%32))))
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2920)))
	if v2921 != int32(6) {
		v2937 = v2915
		goto L541
	} else {
		goto L555
	}
L555:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2920)+4))
	if v2924 != l2 {
		v2937 = v2915
		goto L541
	} else {
		goto L556
	}
L556:
	;
	v2928 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2920)+8)))
	if v2890+int32(1) == v2928 {
		v2890 = v2928
		goto L552
	} else {
		goto L557
	}
L557:
	;
	goto L553
L558:
	;
	v3043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if base.B2i32(v3043 != int32(1))|v1233 != 0 {
		goto L1
	} else {
		goto L568
	}
L559:
	;
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v2959)+4))
	if v2962 <= int32(0) {
		goto L558
	} else {
		goto L560
	}
L560:
	;
	v2969 = int32(0)
	goto L561
L561:
	;
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(v2959)+12))
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v2995+v2969<<(uint(int32(2))%32))))
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v2999)+64))
	v3001 = *(*int32)(unsafe.Add(mBase, uint32(v2999)+12))
	v3002 = F_make_tlist_from_pathtarget(m, v3001)
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L4
	} else {
		goto L563
	}
L562:
	;
	goto L558
L563:
	;
	v3004 = F_convert_subquery_pathkeys(m, l0, l1, v3000, v3002)
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L4
	} else {
		goto L564
	}
L564:
	;
	v3006 = F_create_subqueryscan_path(m, l0, l1, v2999, v2937, v3004, v1233)
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L4
	} else {
		goto L565
	}
L565:
	;
	F_add_path(m, l1, v3006)
	mBase = m.M
	v3009 = m.ExcPending
	if v3009 != 0 {
		goto L4
	} else {
		goto L566
	}
L566:
	;
	v3011 = v2969 + int32(1)
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(v2959)+4))
	if v3011 < v3012 {
		v2969 = v3011
		goto L561
	} else {
		goto L567
	}
L567:
	;
	goto L562
L568:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v2808)+52))
	if v3047 == int32(0) {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(v3047)+4))
	if v3050 <= int32(0) {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	v3057 = int32(0)
	goto L571
L571:
	;
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(v3047)+12))
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v3083+v3057<<(uint(int32(2))%32))))
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v3087)+64))
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v3087)+12))
	v3090 = F_make_tlist_from_pathtarget(m, v3089)
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L4
	} else {
		goto L573
	}
L572:
	;
	goto L1
L573:
	;
	v3092 = F_convert_subquery_pathkeys(m, l0, l1, v3088, v3090)
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		goto L4
	} else {
		goto L574
	}
L574:
	;
	v3095 = F_create_subqueryscan_path(m, l0, l1, v3087, v2937, v3092, int32(0))
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L4
	} else {
		goto L575
	}
L575:
	;
	F_add_partial_path(m, l1, v3095)
	mBase = m.M
	v3098 = m.ExcPending
	if v3098 != 0 {
		goto L4
	} else {
		goto L576
	}
L576:
	;
	v3100 = v3057 + int32(1)
	v3101 = *(*int32)(unsafe.Add(mBase, uint32(v3047)+4))
	if v3100 < v3101 {
		v3057 = v3100
		goto L571
	} else {
		goto L577
	}
L577:
	;
	goto L572
}
