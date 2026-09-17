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
	F_errfinish(m, int32(_a_F_StoreRelCheck_1), int32(2203), int32(_a_F_StoreRelCheck_2))
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_rel_persistence_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_rel_persistence_1), int32(2226), int32(_a_F_get_rel_persistence_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28+v29)+118)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_get_rel_relispartition(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9)+131)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14 & int32(1)
			}
		} else {
			v14 = int32(0)
			return v14 & int32(1)
		}
	}
}
func F_make_rel_from_joinlist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 float64
	_ = v148
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v171 int64
	_ = v171
	var v176 int64
	_ = v176
	var v181 int64
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 float64
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v215 float64
	_ = v215
	var v217 int64
	_ = v217
	var v237 float64
	_ = v237
	var v242 float64
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 float64
	_ = v253
	var v254 float64
	_ = v254
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v260 float64
	_ = v260
	var v263 float64
	_ = v263
	var v267 float64
	_ = v267
	var v270 float64
	_ = v270
	var v274 float64
	_ = v274
	var v276 int64
	_ = v276
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v285 float64
	_ = v285
	var v286 int64
	_ = v286
	var v289 int64
	_ = v289
	var v298 float64
	_ = v298
	var v300 float64
	_ = v300
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v306 float64
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 float64
	_ = v315
	var v318 int32
	_ = v318
	var v322 float64
	_ = v322
	var v323 float64
	_ = v323
	var v324 float64
	_ = v324
	var v333 float64
	_ = v333
	var v336 float64
	_ = v336
	var v339 float64
	_ = v339
	var v346 float64
	_ = v346
	var v347 float64
	_ = v347
	var v358 float64
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v569 int32
	_ = v569
	var v574 int64
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 float64
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 float64
	_ = v637
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v737 int32
	_ = v737
	var v761 float64
	_ = v761
	var v763 float64
	_ = v763
	var v765 float64
	_ = v765
	var v766 int32
	_ = v766
	var v767 float64
	_ = v767
	var v804 float64
	_ = v804
	var v805 int32
	_ = v805
	var v808 float64
	_ = v808
	var v810 float64
	_ = v810
	var v814 float64
	_ = v814
	var v819 float64
	_ = v819
	var v824 int32
	_ = v824
	var v825 float64
	_ = v825
	var v862 int32
	_ = v862
	var v865 float64
	_ = v865
	var v867 float64
	_ = v867
	var v871 float64
	_ = v871
	var v876 float64
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v926 int32
	_ = v926
	var v927 float64
	_ = v927
	var v964 int32
	_ = v964
	var v967 float64
	_ = v967
	var v969 float64
	_ = v969
	var v973 float64
	_ = v973
	var v978 float64
	_ = v978
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1035 int32
	_ = v1035
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1091 int32
	_ = v1091
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1128 float64
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1144 int32
	_ = v1144
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1237 float64
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1301 int32
	_ = v1301
	var v1302 int64
	_ = v1302
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1320 int32
	_ = v1320
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1438 int32
	_ = v1438
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1789 int32
	_ = v1789
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1937 int32
	_ = v1937
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2029 int32
	_ = v2029
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2120 int64
	_ = v2120
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2152 int32
	_ = v2152
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2184 int32
	_ = v2184
	var v2185 int64
	_ = v2185
	var v2187 int64
	_ = v2187
	var v2189 int64
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2294 int32
	_ = v2294
	var v2304 int32
	_ = v2304
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2353 int32
	_ = v2353
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2388 int64
	_ = v2388
	var v2390 int64
	_ = v2390
	var v2392 int64
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2441 int32
	_ = v2441
	var v2449 int32
	_ = v2449
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2465 int32
	_ = v2465
	var v2470 int64
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2479 int32
	_ = v2479
	var v2513 int32
	_ = v2513
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2522 int32
	_ = v2522
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2578 int32
	_ = v2578
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2586 int32
	_ = v2586
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2596 int32
	_ = v2596
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2655 int32
	_ = v2655
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2697 int32
	_ = v2697
	var v2701 int32
	_ = v2701
	var v2706 int64
	_ = v2706
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2755 int32
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2775 int32
	_ = v2775
	var v2812 int32
	_ = v2812
	var v2816 int32
	_ = v2816
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2826 int32
	_ = v2826
	var v2830 int32
	_ = v2830
	var v2835 int64
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2877 int32
	_ = v2877
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2897 int32
	_ = v2897
	var v2931 int32
	_ = v2931
	var v2934 int32
	_ = v2934
	var v2938 int32
	_ = v2938
	var v2978 int32
	_ = v2978
	var v2982 int32
	_ = v2982
	var v2987 int32
	_ = v2987
	var v2990 int32
	_ = v2990
	var v2995 int64
	_ = v2995
	var v3000 int32
	_ = v3000
	var v3004 int32
	_ = v3004
	var v3009 int32
	_ = v3009
	var v3049 int32
	_ = v3049
	var v3053 int32
	_ = v3053
	var v3058 int32
	_ = v3058
	var v3061 int32
	_ = v3061
	var v3096 int32
	_ = v3096
	var v3103 int32
	_ = v3103
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3146 float64
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3158 float64
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3164 int32
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3201 float64
	_ = v3201
	var v3206 float64
	_ = v3206
	var v3214 float64
	_ = v3214
	var v3223 int32
	_ = v3223
	var v3228 int32
	_ = v3228
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3252 int32
	_ = v3252
	var v3260 int32
	_ = v3260
	var v3264 int32
	_ = v3264
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3270 int32
	_ = v3270
	var v3272 int32
	_ = v3272
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3280 int32
	_ = v3280
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3294 int32
	_ = v3294
	var v3296 int32
	_ = v3296
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3308 int32
	_ = v3308
	var v3317 int32
	_ = v3317
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3327 int32
	_ = v3327
	var v3329 int32
	_ = v3329
	var v3331 int32
	_ = v3331
	var v3334 int32
	_ = v3334
	var v3345 float64
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3352 int32
	_ = v3352
	var v3355 int32
	_ = v3355
	var v3358 float64
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3386 float64
	_ = v3386
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 float64
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3446 int32
	_ = v3446
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3494 int32
	_ = v3494
	var v3498 int32
	_ = v3498
	var v3503 int32
	_ = v3503
	var v3505 int32
	_ = v3505
	var v3507 int32
	_ = v3507
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3515 int32
	_ = v3515
	var v3554 int32
	_ = v3554
	var v3556 int32
	_ = v3556
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3561 int32
	_ = v3561
	var v3598 int32
	_ = v3598
	var v3600 int32
	_ = v3600
	var v3602 int32
	_ = v3602
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3625 int32
	_ = v3625
	var v3659 int32
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3669 int32
	_ = v3669
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3682 int32
	_ = v3682
	var v3714 int32
	_ = v3714
	var v3718 int32
	_ = v3718
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3724 int32
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3727 int32
	_ = v3727
	var v3730 int32
	_ = v3730
	var v3736 int32
	_ = v3736
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3761 int32
	_ = v3761
	var v3764 int32
	_ = v3764
	var v3771 int32
	_ = v3771
	var v3773 int32
	_ = v3773
	var v3774 int32
	_ = v3774
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3781 int32
	_ = v3781
	var v3782 int32
	_ = v3782
	var v3783 int32
	_ = v3783
	var v3784 int32
	_ = v3784
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3807 int32
	_ = v3807
	var v3815 int32
	_ = v3815
	var v3818 int32
	_ = v3818
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3834 int32
	_ = v3834
	var v3865 int32
	_ = v3865
	var v3866 int32
	_ = v3866
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3872 int32
	_ = v3872
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3885 int32
	_ = v3885
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3894 int32
	_ = v3894
	var v3901 int32
	_ = v3901
	var v3903 int32
	_ = v3903
	var v3905 int32
	_ = v3905
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3910 int32
	_ = v3910
	var v3917 int32
	_ = v3917
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3929 int32
	_ = v3929
	var v3930 int32
	_ = v3930
	var v3932 int32
	_ = v3932
	var v3935 int32
	_ = v3935
	var v3936 int32
	_ = v3936
	var v3944 int32
	_ = v3944
	var v3975 int32
	_ = v3975
	var v3979 int32
	_ = v3979
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3995 int32
	_ = v3995
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4004 int32
	_ = v4004
	var v4011 int32
	_ = v4011
	var v4013 int32
	_ = v4013
	var v4015 int32
	_ = v4015
	var v4016 int32
	_ = v4016
	var v4018 int32
	_ = v4018
	var v4020 int32
	_ = v4020
	var v4027 int32
	_ = v4027
	var v4030 int32
	_ = v4030
	var v4031 int32
	_ = v4031
	var v4033 int32
	_ = v4033
	var v4034 int32
	_ = v4034
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4112 int32
	_ = v4112
	var v4113 int32
	_ = v4113
	var v4119 int32
	_ = v4119
	var v4125 int32
	_ = v4125
	var v4155 int32
	_ = v4155
	var v4158 int32
	_ = v4158
	var v4169 int32
	_ = v4169
	var v4201 int32
	_ = v4201
	var v4205 int32
	_ = v4205
	var v4206 int32
	_ = v4206
	var v4207 int32
	_ = v4207
	var v4211 int32
	_ = v4211
	var v4212 int32
	_ = v4212
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4217 int32
	_ = v4217
	var v4223 int32
	_ = v4223
	var v4225 int32
	_ = v4225
	var v4226 int32
	_ = v4226
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4237 int32
	_ = v4237
	var v4240 int32
	_ = v4240
	var v4241 int32
	_ = v4241
	var v4248 int32
	_ = v4248
	var v4251 int32
	_ = v4251
	var v4258 int32
	_ = v4258
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4270 int32
	_ = v4270
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4274 int32
	_ = v4274
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4279 int32
	_ = v4279
	var v4280 int32
	_ = v4280
	var v4283 int32
	_ = v4283
	var v4284 int32
	_ = v4284
	var v4294 int32
	_ = v4294
	var v4302 int32
	_ = v4302
	var v4305 int32
	_ = v4305
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4320 int32
	_ = v4320
	var v4351 int32
	_ = v4351
	var v4352 int32
	_ = v4352
	var v4356 int32
	_ = v4356
	var v4357 int32
	_ = v4357
	var v4358 int32
	_ = v4358
	var v4368 int32
	_ = v4368
	var v4369 int32
	_ = v4369
	var v4371 int32
	_ = v4371
	var v4374 int32
	_ = v4374
	var v4375 int32
	_ = v4375
	var v4380 int32
	_ = v4380
	var v4387 int32
	_ = v4387
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4392 int32
	_ = v4392
	var v4394 int32
	_ = v4394
	var v4396 int32
	_ = v4396
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4408 int32
	_ = v4408
	var v4409 int32
	_ = v4409
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4415 int32
	_ = v4415
	var v4416 int32
	_ = v4416
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4495 int32
	_ = v4495
	var v4496 int32
	_ = v4496
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4536 int32
	_ = v4536
	var v4544 int32
	_ = v4544
	var v4576 int32
	_ = v4576
	var v4579 int32
	_ = v4579
	var v4582 int32
	_ = v4582
	var v4586 int32
	_ = v4586
	var v4593 int32
	_ = v4593
	var v4624 int32
	_ = v4624
	var v4628 int32
	_ = v4628
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4631 int32
	_ = v4631
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4644 int32
	_ = v4644
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4653 int32
	_ = v4653
	var v4660 int32
	_ = v4660
	var v4662 int32
	_ = v4662
	var v4664 int32
	_ = v4664
	var v4665 int32
	_ = v4665
	var v4667 int32
	_ = v4667
	var v4669 int32
	_ = v4669
	var v4676 int32
	_ = v4676
	var v4679 int32
	_ = v4679
	var v4680 int32
	_ = v4680
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4722 int32
	_ = v4722
	var v4723 int32
	_ = v4723
	var v4761 int32
	_ = v4761
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4803 int32
	_ = v4803
	var v4807 int32
	_ = v4807
	var v4812 int32
	_ = v4812
	var v4852 int32
	_ = v4852
	var v4856 int32
	_ = v4856
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4864 int32
	_ = v4864
	var v4899 int32
	_ = v4899
	var v4903 int32
	_ = v4903
	var v4905 int32
	_ = v4905
	var v4906 int32
	_ = v4906
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4922 int32
	_ = v4922
	var v4923 int32
	_ = v4923
	var v4925 int32
	_ = v4925
	var v4928 int32
	_ = v4928
	var v4929 int32
	_ = v4929
	var v4934 int32
	_ = v4934
	var v4942 int32
	_ = v4942
	var v4944 int32
	_ = v4944
	var v4946 int32
	_ = v4946
	var v4947 int32
	_ = v4947
	var v4950 int32
	_ = v4950
	var v4954 int32
	_ = v4954
	var v4963 int32
	_ = v4963
	var v4965 int32
	_ = v4965
	var v4967 int32
	_ = v4967
	var v4968 int32
	_ = v4968
	var v5007 int32
	_ = v5007
	var v5045 int32
	_ = v5045
	var v5049 int32
	_ = v5049
	var v5055 int32
	_ = v5055
	var v5059 int32
	_ = v5059
	var v5064 int32
	_ = v5064
	var v5065 int32
	_ = v5065
	var v5066 int32
	_ = v5066
	var v5074 int32
	_ = v5074
	v3 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	if l1 == v3 {
		v5074 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v39 + int32(16)
	return v5074
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 <= int32(0) {
		v5074 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = v3
	v51 = v3
	goto L4
L4:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v48<<(uint(int32(2))%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v43 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v114 = F_lappend(m, v51, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L19
	}
L7:
	;
	if v87 == int32(63) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v111 = F_make_rel_from_joinlist(m, l0, v86)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L18
	}
L10:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v93 = F_find_base_rel(m, l0, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
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
	v100 = m.ExcPending
	if v100 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v113 = v93
	goto L6
L15:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v101
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_0), v39)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_1), int32(3393), int32(_a_F_make_rel_from_joinlist_2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	v113 = v111
	goto L6
L19:
	;
	v117 = v48 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v117 < v118 {
		v48 = v117
		v51 = v114
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L5
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v5074 = v123
	goto L1
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v114
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[0]))
	if v126 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v127 = m.T0[v126].(func(*base.Module, int32, int32, int32) int32)(m, l0, v43, v114)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[1])))
	if v130 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v5074 = v127
	goto L1
L28:
	;
	v3608 = m.G0
	v3610 = v3608 - int32(16)
	m.G0 = v3610
	v3616 = F_palloc0(m, v43<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v3617 = m.ExcPending
	if v3617 != 0 {
		goto L13
	} else {
		goto L409
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[2]))
	if v43 < v134 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v136 = m.G0
	v138 = v136 - int32(32)
	m.G0 = v138
	v140 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v138 + v140
	*(*int32)(unsafe.Add(mBase, uint32(v138)+8)) = v114
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v146 = v144 + v140
	v148 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[3]))
	v153 = base.I64_trunc_sat_f64_s(base.F64_mul(v148, float64(4.503599627370495e+15)))
	v155 = v153 + int64(4354685564936845354)
	v156 = int64(30)
	v159 = int64(-4658895280553007687)
	v160 = (int64(base.Ui64(v155)>>(uint(v156)%64)) ^ v155) * v159
	v161 = int64(27)
	v164 = int64(-7723592293110705685)
	v165 = (int64(base.Ui64(v160)>>(uint(v161)%64)) ^ v160) * v164
	v166 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v146)+8)) = int64(base.Ui64(v165)>>(uint(v166)%64)) ^ v165
	v171 = v153 - int64(7046029254386353131)
	v176 = (int64(base.Ui64(v171)>>(uint(v156)%64)) ^ v171) * v159
	v181 = (int64(base.Ui64(v176)>>(uint(v161)%64)) ^ v176) * v164
	*(*int64)(unsafe.Add(mBase, uint32(v146))) = int64(base.Ui64(v181)>>(uint(v166)%64)) ^ v181
	goto L31
L31:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[4]))
	if int32(1) < v187 {
		v372 = v187
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[5]))
	v382 = F_palloc(m, int32(12))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L13
	} else {
		goto L74
	}
L33:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[6]))
	v194 = base.F64_add(base.F64_convert_i32_s(v43), float64(1))
	goto L35
L34:
	;
	v360 = v191 * int32(50)
	if base.F64_gt(v358, base.F64_convert_i32_s(v360)) != 0 {
		v372 = v360
		goto L32
	} else {
		goto L72
	}
L35:
	;
	v200 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v194))>>(uint(int64(52))%64))) & int32(2047)
	v205 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	goto L36
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v205) <= base.Ui32(v200-v205) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v215 = base.F64_add(v194, float64(1))
	if base.Ui32(v200) < base.Ui32(v205) {
		v358 = v215
		goto L34
	} else {
		goto L41
	}
L39:
	;
	v249 = v200
	goto L40
L40:
	;
	v253 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[7]))
	v254 = base.F64_add(v194, v253)
	v256 = base.F64_sub(v194, base.F64_sub(v254, v253))
	v257 = base.F64_mul(v256, v256)
	v260 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[8]))
	v263 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[9]))
	v267 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[10]))
	v270 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[11]))
	v274 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[12]))
	v276 = base.I64_reinterpret_f64(v254)
	v281 = base.I32_wrap_i64(v276) << (uint(int32(4)) % 32) & int32(2032)
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_make_rel_from_joinlist[13])))
	v285 = base.F64_add(base.F64_mul(base.F64_mul(v257, v257), base.F64_add(base.F64_mul(v256, v260), v263)), base.F64_add(base.F64_mul(v257, base.F64_add(base.F64_mul(v256, v267), v270)), base.F64_add(base.F64_mul(v256, v274), v282)))
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_c_F_make_rel_from_joinlist[14])))
	v289 = v286 + v276<<(uint(int64(45))%64)
	if v249 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v217 = base.I64_reinterpret_f64(v194)
	goto L43
L42:
	;
	if base.Ui64(v217<<(uint(int64(1))%64)) <= base.Ui64(int64(-9143996093422370816)) {
		goto L54
	} else {
		goto L55
	}
L43:
	;
	if base.Ui32(v200) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if v217 == int64(-4503599627370496) {
		v358 = float64(0)
		goto L34
	} else {
		goto L45
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(math.Float64frombits(uint64(0x7ff0000000000000))))>>(uint(int64(52))%64)))) <= base.Ui32(v200) {
		v358 = v215
		goto L34
	} else {
		goto L47
	}
L47:
	;
	if int64(0) <= v217 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v237 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
	mBase = m.M
	goto L51
L49:
	;
	goto L50
L50:
	;
	if base.Ui64(v217) < base.Ui64(int64(-4570929321408987136)) {
		goto L42
	} else {
		goto L52
	}
L51:
	;
	v358 = v237
	goto L34
L52:
	;
	v242 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
	mBase = m.M
	goto L53
L53:
	;
	v358 = v242
	goto L34
L54:
	;
	v248 = v200
	goto L56
L55:
	;
	v248 = int32(0)
	goto L56
L56:
	;
	v249 = v248
	goto L40
L57:
	;
	if v276&int64(2147483648) == int64(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	v347 = base.F64_reinterpret_i64(v289)
	v358 = base.F64_add(base.F64_mul(v347, v285), v347)
	goto L34
L60:
	;
	v358 = v346
	goto L34
L61:
	;
	v298 = base.F64_reinterpret_i64(v289 - int64(4503599627370496))
	v300 = base.F64_add(base.F64_mul(v298, v285), v298)
	v346 = base.F64_add(v300, v300)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v304 = base.F64_reinterpret_i64(v289 + int64(4602678819172646912))
	v305 = base.F64_mul(v304, v285)
	v306 = base.F64_add(v305, v304)
	if base.F64_lt(v306, float64(1)) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v310 = m.G0
	v312 = v310 - int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+8)) = int64(4503599627370496)
	v315 = *(*float64)(unsafe.Add(mBase, uint32(v312)+8))
	goto L67
L65:
	;
	v339 = v306
	goto L66
L66:
	;
	v346 = base.F64_mul(v339, float64(2.2250738585072014e-308))
	goto L60
L67:
	;
	v318 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v318-int32(16))+8)) = base.F64_mul(v315, float64(2.2250738585072014e-308))
	goto L68
L68:
	;
	v322 = float64(0)
	v323 = float64(1)
	v324 = base.F64_add(v306, v323)
	v333 = base.F64_add(base.F64_add(v324, base.F64_add(base.F64_add(v305, base.F64_sub(v304, v306)), base.F64_add(v306, base.F64_sub(v323, v324)))), float64(-1))
	if base.F64_eq(v333, v322) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v336 = v322
	goto L71
L70:
	;
	v336 = v333
	goto L71
L71:
	;
	v339 = v336
	goto L66
L72:
	;
	v364 = v191 * int32(10)
	if base.F64_lt(v358, base.F64_convert_i32_s(v364)) != 0 {
		v372 = v364
		goto L32
	} else {
		goto L73
	}
L73:
	;
	v372 = base.I32_trunc_sat_f64_s(base.F64_ceil(v358))
	goto L32
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v382)+8)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v382)+4)) = v372
	v388 = F_palloc(m, v372<<(uint(int32(4))%32))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v382))) = v388
	if int32(0) < v372 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v398 = int32(0)
	goto L79
L77:
	;
	goto L78
L78:
	;
	v479 = int32(0)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v479 < v480 {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	v437 = F_palloc(m, v43<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L13
	} else {
		goto L81
	}
L80:
	;
	goto L78
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388+v398<<(uint(int32(4))%32)))) = v437
	v441 = v398 + int32(1)
	if v441 != v372 {
		v398 = v441
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	F_pg_qsort(m, v699, v700, int32(16), int32(819))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L13
	} else {
		goto L111
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L13
	} else {
		goto L108
	}
L85:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v485 = v479
	v488 = v3
	goto L88
L86:
	;
	goto L87
L87:
	;
	goto L83
L88:
	;
	v521 = v488 << (uint(int32(4)) % 32)
	v522 = v483 + v521
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	if v525 <= int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L87
L90:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v630 = F_geqo_eval(m, l0, v628, v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L13
	} else {
		goto L100
	}
L91:
	;
	v528 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v523))) = v528
	if v525 == v528 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v534 = int32(1)
	goto L93
L93:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v574 = F_pg_prng_uint64_range(m, v569+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v534))
	mBase = m.M
	v575 = base.I32_wrap_i64(v574)
	goto L95
L94:
	;
	goto L90
L95:
	;
	if v575 != v534 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v577 = int32(2)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v523+v575<<(uint(v577)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v523+v534<<(uint(v577)%32)))) = v583
	goto L98
L97:
	;
	goto L98
L98:
	;
	v589 = v534 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v523+v575<<(uint(int32(2))%32)))) = v589
	if v589 != v525 {
		v534 = v589
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L94
L100:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	*(*float64)(unsafe.Add(mBase, uint32(v632+v521)+8)) = v630
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v637 = *(*float64)(unsafe.Add(mBase, uint32(v635+v521)+8))
	if base.F64_lt(v637, float64(1.7976931348623157e+308)) != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v647 < v648 {
		v485 = v646
		v488 = v647
		goto L88
	} else {
		goto L107
	}
L102:
	;
	v646 = v485
	v647 = v488 + int32(1)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v643 = v485 + int32(1)
	if v488 != 0 {
		v646 = v643
		v647 = v488
		goto L101
	} else {
		goto L105
	}
L105:
	;
	if int32(_a_F_make_rel_from_joinlist_3) <= v643 {
		goto L84
	} else {
		goto L106
	}
L106:
	;
	v646 = v643
	v647 = v488
	goto L101
L107:
	;
	goto L89
L108:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_4), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_5), int32(117), int32(_a_F_make_rel_from_joinlist_6))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v706 = F_alloc_chromo(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v709 = F_alloc_chromo(m, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v712 = int32(24)
	v716 = F_palloc(m, v711*v712+v712)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L13
	} else {
		goto L114
	}
L114:
	;
	if int32(0) < v380 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v720 = v380
	goto L117
L116:
	;
	v720 = v372
	goto L117
L117:
	;
	if int32(0) < v720 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v737 = int32(0)
	goto L121
L119:
	;
	goto L120
L120:
	;
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3484)))
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v3487 = F_gimme_tree(m, l0, v3485, v3486)
	mBase = m.M
	v3488 = m.ExcPending
	if v3488 != 0 {
		goto L13
	} else {
		goto L390
	}
L121:
	;
	v761 = *(*float64)(unsafe.Add(mBase, _c_F_make_rel_from_joinlist[15]))
	v763 = base.F64_add(v761, float64(-1))
	v765 = base.F64_mul(v763, float64(4))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v767 = base.F64_convert_i32_s(v766)
	goto L123
L122:
	;
	goto L120
L123:
	;
	v804 = base.F64_mul(v761, v761)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v808 = F_pg_prng_double(m, v805+int32(8))
	mBase = m.M
	goto L125
L124:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v825 = base.F64_convert_i32_s(v824)
	goto L130
L125:
	;
	v810 = base.F64_sub(v804, base.F64_mul(v765, v808))
	if base.F64_gt(v810, float64(0)) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v814 = base.F64_sqrt(v810)
	goto L128
L127:
	;
	v814 = v810
	goto L128
L128:
	;
	v819 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v761, v814), v767), float64(0.5)), v763)
	if base.F64_lt(v819, float64(0))|base.F64_ge(v819, v767) != 0 {
		goto L123
	} else {
		goto L129
	}
L129:
	;
	goto L124
L130:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v865 = F_pg_prng_double(m, v862+int32(8))
	mBase = m.M
	goto L132
L131:
	;
	v881 = base.I32_trunc_sat_f64_s(v876)
	v882 = base.I32_trunc_sat_f64_s(v819)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if base.B2i32(v881 != v882)|base.B2i32(v884 < int32(2)) == int32(0) {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v867 = base.F64_sub(v804, base.F64_mul(v765, v865))
	if base.F64_gt(v867, float64(0)) != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v871 = base.F64_sqrt(v867)
	goto L135
L134:
	;
	v871 = v867
	goto L135
L135:
	;
	v876 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v761, v871), v825), float64(0.5)), v763)
	if base.F64_lt(v876, float64(0))|base.F64_le(v825, v876) != 0 {
		goto L130
	} else {
		goto L136
	}
L136:
	;
	goto L131
L137:
	;
	goto L140
L138:
	;
	v989 = v881
	goto L139
L139:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v1024 = v1021 + v882<<(uint(int32(4))%32)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v1026 = int32(0)
	if v1025 <= v1026 {
		goto L151
	} else {
		goto L152
	}
L140:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v927 = base.F64_convert_i32_s(v926)
	goto L142
L141:
	;
	v989 = v983
	goto L139
L142:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v967 = F_pg_prng_double(m, v964+int32(8))
	mBase = m.M
	goto L144
L143:
	;
	v983 = base.I32_trunc_sat_f64_s(v978)
	if v983 == v882 {
		goto L140
	} else {
		goto L149
	}
L144:
	;
	v969 = base.F64_sub(v804, base.F64_mul(v765, v967))
	if base.F64_gt(v969, float64(0)) != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v973 = base.F64_sqrt(v969)
	goto L147
L146:
	;
	v973 = v969
	goto L147
L147:
	;
	v978 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v761, v973), v927), float64(0.5)), v763)
	if base.F64_lt(v978, float64(0))|base.F64_ge(v978, v927) != 0 {
		goto L142
	} else {
		goto L148
	}
L148:
	;
	goto L143
L149:
	;
	goto L141
L150:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v1133 = v1130 + v989<<(uint(int32(4))%32)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v1135 = int32(0)
	if v1134 <= v1135 {
		goto L164
	} else {
		goto L165
	}
L151:
	;
	v1128 = *(*float64)(unsafe.Add(mBase, uint32(v1024)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v706)+8)) = v1128
	goto L150
L152:
	;
	v1035 = v1025 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1025) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1043 = v1026
	v1047 = v1026
	goto L156
L154:
	;
	v1091 = v1026
	goto L155
L155:
	;
	v1100 = v1091
	v1105 = v1026
	goto L160
L156:
	;
	v1050 = v1043 << (uint(int32(2)) % 32)
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1053+v1050)))
	*(*int32)(unsafe.Add(mBase, uint32(v1050+v1051))) = v1055
	v1057 = int32(4)
	v1058 = v1050 | v1057
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1061+v1058)))
	*(*int32)(unsafe.Add(mBase, uint32(v1058+v1059))) = v1063
	v1066 = v1050 | int32(8)
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1069+v1066)))
	*(*int32)(unsafe.Add(mBase, uint32(v1066+v1067))) = v1071
	v1074 = v1050 | int32(12)
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1077+v1074)))
	*(*int32)(unsafe.Add(mBase, uint32(v1074+v1075))) = v1079
	v1082 = v1043 + v1057
	v1084 = v1047 + v1057
	if v1084 != v1025&int32(2147483644) {
		v1043 = v1082
		v1047 = v1084
		goto L156
	} else {
		goto L158
	}
L157:
	;
	if v1035 == int32(0) {
		goto L151
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	v1091 = v1082
	goto L155
L160:
	;
	v1107 = v1100 << (uint(int32(2)) % 32)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1110+v1107)))
	*(*int32)(unsafe.Add(mBase, uint32(v1107+v1108))) = v1112
	v1114 = int32(1)
	v1117 = v1105 + v1114
	if v1117 != v1035 {
		v1100 = v1100 + v1114
		v1105 = v1117
		goto L160
	} else {
		goto L162
	}
L161:
	;
	goto L151
L162:
	;
	goto L161
L163:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v1241 = int32(0)
	v1242 = int32(1)
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	if base.B2i32(v1243 <= v1241) == v1241 {
		goto L176
	} else {
		goto L177
	}
L164:
	;
	v1237 = *(*float64)(unsafe.Add(mBase, uint32(v1133)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v709)+8)) = v1237
	goto L163
L165:
	;
	v1144 = v1134 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1134) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1152 = v1135
	v1156 = v1135
	goto L169
L167:
	;
	v1200 = v1135
	goto L168
L168:
	;
	v1209 = v1200
	v1214 = v1135
	goto L173
L169:
	;
	v1159 = v1152 << (uint(int32(2)) % 32)
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1162+v1159)))
	*(*int32)(unsafe.Add(mBase, uint32(v1159+v1160))) = v1164
	v1166 = int32(4)
	v1167 = v1159 | v1166
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1167)))
	*(*int32)(unsafe.Add(mBase, uint32(v1167+v1168))) = v1172
	v1175 = v1159 | int32(8)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1178+v1175)))
	*(*int32)(unsafe.Add(mBase, uint32(v1175+v1176))) = v1180
	v1183 = v1159 | int32(12)
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1186+v1183)))
	*(*int32)(unsafe.Add(mBase, uint32(v1183+v1184))) = v1188
	v1191 = v1152 + v1166
	v1193 = v1156 + v1166
	if v1193 != v1134&int32(2147483644) {
		v1152 = v1191
		v1156 = v1193
		goto L169
	} else {
		goto L171
	}
L170:
	;
	if v1144 == int32(0) {
		goto L164
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v1200 = v1191
	goto L168
L173:
	;
	v1216 = v1209 << (uint(int32(2)) % 32)
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1219+v1216)))
	*(*int32)(unsafe.Add(mBase, uint32(v1216+v1217))) = v1221
	v1223 = int32(1)
	v1226 = v1214 + v1223
	if v1226 != v1144 {
		v1209 = v1209 + v1223
		v1214 = v1226
		goto L173
	} else {
		goto L175
	}
L174:
	;
	goto L164
L175:
	;
	goto L174
L176:
	;
	v1248 = int32(2)
	v1250 = v1243 + int32(1)
	if v1250 <= v1248 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v2110 = m.G0
	v2112 = v2110 - int32(32)
	m.G0 = v2112
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2120 = F_pg_prng_uint64_range(m, v2115+int32(8), base.I64_extend_i32_s(int32(1)), base.I64_extend_i32_s(v2108))
	mBase = m.M
	goto L235
L179:
	;
	v1253 = v1248
	goto L181
L180:
	;
	v1253 = v1250
	goto L181
L181:
	;
	v1255 = v1253 - int32(1)
	v1257 = v1255 & int32(3)
	if int32(5) <= v1250 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v1438 = int32(0)
	goto L193
L183:
	;
	v1264 = int32(0)
	v1266 = v1242
	goto L186
L184:
	;
	v1320 = v1242
	goto L185
L185:
	;
	v1355 = int32(0)
	v1357 = v1320
	goto L190
L186:
	;
	v1301 = v716 + v1266*int32(24)
	v1302 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1301)+88)) = v1302
	*(*int64)(unsafe.Add(mBase, uint32(v1301)+64)) = v1302
	*(*int64)(unsafe.Add(mBase, uint32(v1301)+40)) = v1302
	*(*int64)(unsafe.Add(mBase, uint32(v1301)+16)) = v1302
	v1310 = int32(4)
	v1311 = v1266 + v1310
	v1313 = v1264 + v1310
	if v1313 != v1255&int32(-4) {
		v1264 = v1313
		v1266 = v1311
		goto L186
	} else {
		goto L188
	}
L187:
	;
	if v1257 == int32(0) {
		goto L182
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	v1320 = v1311
	goto L185
L190:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v716+v1357*int32(24))+16)) = int64(0)
	v1395 = int32(1)
	v1398 = v1355 + v1395
	if v1398 != v1257 {
		v1355 = v1398
		v1357 = v1357 + v1395
		goto L190
	} else {
		goto L192
	}
L191:
	;
	goto L182
L192:
	;
	goto L191
L193:
	;
	v1474 = v1438 + int32(1)
	if v1474 != v1243 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L178
L195:
	;
	v1477 = v1474
	goto L197
L196:
	;
	v1477 = int32(0)
	goto L197
L197:
	;
	v1478 = int32(2)
	v1479 = v1477 << (uint(v1478) % 32)
	v1480 = v1239 + v1479
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1480)))
	v1482 = int32(0)
	v1484 = v1438 << (uint(v1478) % 32)
	v1485 = v1239 + v1484
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1485)))
	v1489 = v716 + v1486*int32(24)
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+16))
	if v1490 <= v1482 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v1631 = int32(0)
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1485)))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1480)))
	v1636 = v716 + v1633*int32(24)
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+16))
	if v1637 <= v1631 {
		goto L208
	} else {
		goto L209
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1489+v1490<<(uint(int32(2))%32)))) = v1481
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+16))
	v1586 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1489)+16)) = v1585 + v1586
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1489)+20)) = v1589 + v1586
	goto L198
L200:
	;
	v1496 = v1482
	goto L201
L201:
	;
	v1531 = v1489 + v1496<<(uint(int32(2))%32)
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1531)))
	v1534 = v1532 >> (uint(int32(31)) % 32)
	if v1481 != v1532^v1534-v1534 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1531))) = int32(0) - v1481
	goto L198
L203:
	;
	v1539 = v1496 + int32(1)
	if v1490 != v1539 {
		v1496 = v1539
		goto L201
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	goto L202
L206:
	;
	goto L199
L207:
	;
	v1775 = int32(0)
	v1776 = v1240 + v1479
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1776)))
	v1778 = v1240 + v1484
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1778)))
	v1782 = v716 + v1779*int32(24)
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+16))
	if v1783 <= v1775 {
		goto L217
	} else {
		goto L218
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1636+v1637<<(uint(int32(2))%32)))) = v1632
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+16))
	v1732 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1636)+16)) = v1731 + v1732
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1636)+20)) = v1735 + v1732
	goto L207
L209:
	;
	v1643 = v1631
	goto L210
L210:
	;
	v1678 = v1636 + v1643<<(uint(int32(2))%32)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1678)))
	v1681 = v1679 >> (uint(int32(31)) % 32)
	if v1632 != v1679^v1681-v1681 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1678))) = int32(0) - v1632
	goto L207
L212:
	;
	v1686 = v1643 + int32(1)
	if v1637 != v1686 {
		v1643 = v1686
		goto L210
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	goto L211
L215:
	;
	goto L208
L216:
	;
	v1925 = int32(0)
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1778)))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1776)))
	v1930 = v716 + v1927*int32(24)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1930)+16))
	if v1931 <= v1925 {
		goto L226
	} else {
		goto L227
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1782+v1783<<(uint(int32(2))%32)))) = v1777
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+16))
	v1879 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+16)) = v1878 + v1879
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+20)) = v1882 + v1879
	goto L216
L218:
	;
	v1789 = v1775
	goto L219
L219:
	;
	v1824 = v1782 + v1789<<(uint(int32(2))%32)
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	v1827 = v1825 >> (uint(int32(31)) % 32)
	if v1777 != v1825^v1827-v1827 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1824))) = int32(0) - v1777
	goto L216
L221:
	;
	v1832 = v1789 + int32(1)
	if v1783 != v1832 {
		v1789 = v1832
		goto L219
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	goto L220
L224:
	;
	goto L217
L225:
	;
	if v1474 != v1243 {
		v1438 = v1474
		goto L193
	} else {
		goto L234
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1930+v1931<<(uint(int32(2))%32)))) = v1926
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1930)+16))
	v2026 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1930)+16)) = v2025 + v2026
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v1930)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1930)+20)) = v2029 + v2026
	goto L225
L227:
	;
	v1937 = v1925
	goto L228
L228:
	;
	v1972 = v1930 + v1937<<(uint(int32(2))%32)
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1972)))
	v1975 = v1973 >> (uint(int32(31)) % 32)
	if v1926 != v1973^v1975-v1975 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1972))) = int32(0) - v1926
	goto L225
L230:
	;
	v1980 = v1937 + int32(1)
	if v1931 != v1980 {
		v1937 = v1980
		goto L228
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	goto L229
L233:
	;
	goto L226
L234:
	;
	goto L194
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2107))) = base.I32_wrap_i64(v2120)
	if int32(2) <= v2108 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v2125 = int32(2)
	v2127 = v2108 + int32(1)
	if v2127 <= v2125 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L238
L238:
	;
	m.G0 = v2112 + int32(32)
	v3144 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v3146 = F_geqo_eval(m, l0, v3144, v3145)
	mBase = m.M
	v3147 = m.ExcPending
	if v3147 != 0 {
		goto L13
	} else {
		goto L353
	}
L239:
	;
	v2130 = v2125
	goto L241
L240:
	;
	v2130 = v2127
	goto L241
L241:
	;
	v2131 = int32(1)
	v2132 = v2130 - v2131
	v2152 = v2131
	goto L242
L242:
	;
	v2178 = v2107 + v2152<<(uint(int32(2))%32)
	v2180 = v2178 - int32(4)
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v2180)))
	v2184 = v716 + v2181*int32(24)
	v2185 = *(*int64)(unsafe.Add(mBase, uint32(v2184)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+24)) = v2185
	v2187 = *(*int64)(unsafe.Add(mBase, uint32(v2184)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+16)) = v2187
	v2189 = *(*int64)(unsafe.Add(mBase, uint32(v2184)))
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+8)) = v2189
	v2191 = int32(0)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+28))
	if v2191 < v2192 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	goto L238
L244:
	;
	v2196 = v2191
	goto L247
L245:
	;
	v2353 = v2181
	goto L246
L246:
	;
	v2384 = v716 + v2353*int32(24)
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2384)+20))
	if int32(0) < v2385 {
		goto L262
	} else {
		goto L263
	}
L247:
	;
	v2231 = int32(0)
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2112+int32(8)+v2196<<(uint(int32(2))%32))))
	v2239 = v2237 >> (uint(int32(31)) % 32)
	v2244 = v716 + (v2237^v2239-v2239)*int32(24)
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+20))
	if v2245 <= v2231 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2180)))
	v2353 = v2345
	goto L246
L249:
	;
	v2343 = v2196 + int32(1)
	if v2343 != v2192 {
		v2196 = v2343
		goto L247
	} else {
		goto L257
	}
L250:
	;
	v2250 = v2231
	goto L251
L251:
	;
	v2286 = v2244 + v2250<<(uint(int32(2))%32)
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2286)))
	v2289 = v2287 >> (uint(int32(31)) % 32)
	if v2181 != v2287^v2289-v2289 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2244)+20)) = v2245 - int32(1)
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2244+v2245<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v2286))) = v2304
	goto L249
L253:
	;
	v2294 = v2250 + int32(1)
	if v2245 != v2294 {
		v2250 = v2294
		goto L251
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	goto L252
L256:
	;
	goto L249
L257:
	;
	goto L248
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2178))) = v3061
	v3096 = *(*int32)(unsafe.Add(mBase, uint32(v2180)))
	*(*int32)(unsafe.Add(mBase, uint32(v716+v3096*int32(24))+20)) = int32(-1)
	v3103 = v2152 + int32(1)
	if v3103 != v2108 {
		v2152 = v3103
		goto L242
	} else {
		goto L352
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3049 = m.ExcPending
	if v3049 != 0 {
		goto L13
	} else {
		goto L349
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3000 = m.ExcPending
	if v3000 != 0 {
		goto L13
	} else {
		goto L346
	}
L261:
	;
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2995 = F_pg_prng_uint64_range(m, v2990+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(int32(-2)))
	mBase = m.M
	goto L345
L262:
	;
	v2388 = *(*int64)(unsafe.Add(mBase, uint32(v2384)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+24)) = v2388
	v2390 = *(*int64)(unsafe.Add(mBase, uint32(v2384)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+16)) = v2390
	v2392 = *(*int64)(unsafe.Add(mBase, uint32(v2384)))
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+8)) = v2392
	v2394 = int32(0)
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+28))
	if v2395 <= v2394 {
		goto L261
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v2526 = int32(0)
	v2528 = int32(1)
	if base.B2i32(v2127 < int32(3)) == v2526 {
		goto L286
	} else {
		goto L287
	}
L265:
	;
	v2402 = v2394
	v2403 = int32(5)
	v2407 = int32(-1)
	goto L266
L266:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2112+int32(8)+v2402<<(uint(int32(2))%32))))
	if v2441 < int32(0) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v2461 = int32(0)
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2470 = F_pg_prng_uint64_range(m, v2465+int32(8), base.I64_extend_i32_s(v2461), base.I64_extend_i32_s(v2457-int32(1)))
	mBase = m.M
	goto L277
L268:
	;
	v3061 = int32(0) - v2441
	goto L258
L269:
	;
	goto L270
L270:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v716+v2441*int32(24))+20))
	if v2449 < v2403 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v2459 = v2402 + int32(1)
	if v2459 != v2395 {
		v2402 = v2459
		v2403 = v2456
		v2407 = v2457
		goto L266
	} else {
		goto L276
	}
L272:
	;
	v2456 = v2449
	v2457 = int32(1)
	goto L271
L273:
	;
	goto L274
L274:
	;
	if v2407 == int32(-1) {
		goto L260
	} else {
		goto L275
	}
L275:
	;
	v2456 = v2403
	v2457 = v2407 + base.B2i32(v2449 == v2403)
	goto L271
L276:
	;
	goto L267
L277:
	;
	v2472 = v2461
	v2479 = v2457
	goto L278
L278:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v2112+int32(8)+v2472<<(uint(int32(2))%32))))
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(v716+v2513*int32(24))+20))
	if v2456 == v2517 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L259
L280:
	;
	v2520 = v2479 - int32(1)
	if v2520 == base.I32_wrap_i64(v2470) {
		v3061 = v2513
		goto L258
	} else {
		goto L283
	}
L281:
	;
	v2522 = v2479
	goto L282
L282:
	;
	v2524 = v2472 + int32(1)
	if v2524 != v2395 {
		v2472 = v2524
		v2479 = v2522
		goto L278
	} else {
		goto L284
	}
L283:
	;
	v2522 = v2520
	goto L282
L284:
	;
	goto L279
L285:
	;
	if v2662 != 0 {
		goto L307
	} else {
		goto L308
	}
L286:
	;
	v2533 = v2526
	v2534 = v2526
	v2535 = v2528
	v2536 = v2526
	goto L289
L287:
	;
	v2611 = v2526
	v2613 = v2526
	v2614 = v2528
	goto L288
L288:
	;
	if v2614 == v2353 {
		goto L299
	} else {
		goto L300
	}
L289:
	;
	if v2535 == v2353 {
		v2582 = v2533
		v2583 = v2536
		goto L291
	} else {
		goto L292
	}
L290:
	;
	if v2132&v2131 == int32(0) {
		v2659 = v2600
		v2662 = v2601
		goto L285
	} else {
		goto L298
	}
L291:
	;
	v2586 = v2535 + int32(1)
	if v2586 == v2353 {
		v2600 = v2582
		v2601 = v2583
		goto L294
	} else {
		goto L295
	}
L292:
	;
	v2572 = v716 + v2535*int32(24)
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v2572)+20))
	if v2573 == int32(-1) {
		v2582 = v2533
		v2583 = v2536
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(v2572)+16))
	v2582 = v2533 + int32(1)
	v2583 = v2536 + base.B2i32(v2578 == int32(4))
	goto L291
L294:
	;
	v2603 = int32(2)
	v2604 = v2535 + v2603
	v2606 = v2534 + v2603
	if v2606 != v2132&int32(-2) {
		v2533 = v2600
		v2534 = v2606
		v2535 = v2604
		v2536 = v2601
		goto L289
	} else {
		goto L297
	}
L295:
	;
	v2590 = v716 + v2586*int32(24)
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+20))
	if v2591 == int32(-1) {
		v2600 = v2582
		v2601 = v2583
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+16))
	v2600 = v2582 + int32(1)
	v2601 = v2583 + base.B2i32(v2596 == int32(4))
	goto L294
L297:
	;
	goto L290
L298:
	;
	v2611 = v2600
	v2613 = v2601
	v2614 = v2604
	goto L288
L299:
	;
	v2659 = v2611
	v2662 = v2613
	goto L285
L300:
	;
	goto L301
L301:
	;
	v2649 = v716 + v2614*int32(24)
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2649)+20))
	if v2650 == int32(-1) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v2659 = v2611
	v2662 = v2613
	goto L285
L303:
	;
	goto L304
L304:
	;
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v2649)+16))
	v2659 = v2611 + int32(1)
	v2662 = v2613 + base.B2i32(v2655 == int32(4))
	goto L285
L305:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L13
	} else {
		goto L342
	}
L306:
	;
	F_errmsg_internal(m, v2931, int32(0))
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L13
	} else {
		goto L340
	}
L307:
	;
	v2697 = int32(1)
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2706 = F_pg_prng_uint64_range(m, v2701+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v2662-v2697))
	mBase = m.M
	goto L310
L308:
	;
	goto L309
L309:
	;
	if v2659 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L310:
	;
	v2710 = v2697
	v2711 = v2662
	goto L311
L311:
	;
	if v2710 == v2353 {
		v2758 = v2711
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v2764 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L13
	} else {
		goto L319
	}
L313:
	;
	v2760 = v2710 + int32(1)
	if v2760 <= v2108 {
		v2710 = v2760
		v2711 = v2758
		goto L311
	} else {
		goto L318
	}
L314:
	;
	v2747 = v716 + v2710*int32(24)
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2747)+20))
	if v2748 == int32(-1) {
		v2758 = v2711
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2747)+16))
	if v2751 != int32(4) {
		v2758 = v2711
		goto L313
	} else {
		goto L316
	}
L316:
	;
	v2755 = v2711 - int32(1)
	if base.I32_wrap_i64(v2706) == v2755 {
		v3061 = v2710
		goto L258
	} else {
		goto L317
	}
L317:
	;
	v2758 = v2755
	goto L313
L318:
	;
	goto L312
L319:
	;
	if v2764 == int32(0) {
		goto L305
	} else {
		goto L320
	}
L320:
	;
	v2897 = int32(422)
	v2931 = int32(_a_F_make_rel_from_joinlist_7)
	goto L306
L321:
	;
	v2775 = int32(1)
	goto L324
L322:
	;
	goto L323
L323:
	;
	v2826 = int32(1)
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2835 = F_pg_prng_uint64_range(m, v2830+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v2659-v2826))
	mBase = m.M
	goto L330
L324:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v716+v2775*int32(24))+20))
	if int32(0) <= v2812 {
		v3061 = v2775
		goto L258
	} else {
		goto L326
	}
L325:
	;
	v2820 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L13
	} else {
		goto L328
	}
L326:
	;
	v2816 = v2775 + int32(1)
	if v2816 <= v2108 {
		v2775 = v2816
		goto L324
	} else {
		goto L327
	}
L327:
	;
	goto L325
L328:
	;
	if v2820 == int32(0) {
		goto L305
	} else {
		goto L329
	}
L329:
	;
	v2897 = int32(461)
	v2931 = int32(_a_F_make_rel_from_joinlist_8)
	goto L306
L330:
	;
	v2837 = v2659
	v2839 = v2826
	goto L331
L331:
	;
	if v2839 == v2353 {
		v2883 = v2837
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v2889 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2890 = m.ExcPending
	if v2890 != 0 {
		goto L13
	} else {
		goto L338
	}
L333:
	;
	v2885 = v2839 + int32(1)
	if v2885 <= v2108 {
		v2837 = v2883
		v2839 = v2885
		goto L331
	} else {
		goto L337
	}
L334:
	;
	v2877 = *(*int32)(unsafe.Add(mBase, uint32(v716+v2839*int32(24))+20))
	if v2877 == int32(-1) {
		v2883 = v2837
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v2881 = v2837 - int32(1)
	if base.I32_wrap_i64(v2835) == v2881 {
		v3061 = v2839
		goto L258
	} else {
		goto L336
	}
L336:
	;
	v2883 = v2881
	goto L333
L337:
	;
	goto L332
L338:
	;
	if v2889 == int32(0) {
		goto L305
	} else {
		goto L339
	}
L339:
	;
	v2897 = int32(443)
	v2931 = int32(_a_F_make_rel_from_joinlist_9)
	goto L306
L340:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_10), v2897, int32(_a_F_make_rel_from_joinlist_11))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L13
	} else {
		goto L341
	}
L341:
	;
	goto L305
L342:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_12), int32(0))
	mBase = m.M
	v2982 = m.ExcPending
	if v2982 != 0 {
		goto L13
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_10), int32(466), int32(_a_F_make_rel_from_joinlist_11))
	mBase = m.M
	v2987 = m.ExcPending
	if v2987 != 0 {
		goto L13
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	goto L259
L346:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_13), int32(0))
	mBase = m.M
	v3004 = m.ExcPending
	if v3004 != 0 {
		goto L13
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_10), int32(337), int32(_a_F_make_rel_from_joinlist_14))
	mBase = m.M
	v3009 = m.ExcPending
	if v3009 != 0 {
		goto L13
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_15), int32(0))
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L13
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_10), int32(362), int32(_a_F_make_rel_from_joinlist_14))
	mBase = m.M
	v3058 = m.ExcPending
	if v3058 != 0 {
		goto L13
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	goto L243
L353:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v706)+8)) = v3146
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v3154 = v3152 - int32(1)
	v3158 = *(*float64)(unsafe.Add(mBase, uint32(v3151+v3154<<(uint(int32(4))%32))+8))
	if base.F64_gt(v3146, v3158) != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v3446 = v737 + int32(1)
	if v3446 != v720 {
		v737 = v3446
		goto L121
	} else {
		goto L389
	}
L355:
	;
	v3161 = base.I32_div_s(v3152, int32(2))
	v3164 = int32(0)
	v3165 = v3161
	v3166 = v3154
	goto L356
L356:
	;
	v3201 = *(*float64)(unsafe.Add(mBase, uint32(v3151+v3164<<(uint(int32(4))%32))+8))
	if base.F64_ge(v3201, v3146) != 0 {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	v3241 = v3151 + v3152<<(uint(int32(4))%32) - int32(16)
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v3243 = int32(0)
	if v3242 <= v3243 {
		goto L373
	} else {
		goto L374
	}
L358:
	;
	if v3230 == int32(-1) {
		v3164 = v3231
		v3165 = v3234
		v3166 = v3232
		goto L356
	} else {
		goto L371
	}
L359:
	;
	v3230 = v3164
	v3231 = v3164
	v3232 = v3166
	v3234 = v3165
	goto L358
L360:
	;
	goto L361
L361:
	;
	v3206 = *(*float64)(unsafe.Add(mBase, uint32(v3151+v3165<<(uint(int32(4))%32))+8))
	if base.F64_eq(v3206, v3146) != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v3230 = v3165
	v3231 = v3164
	v3232 = v3166
	v3234 = v3165
	goto L358
L363:
	;
	goto L364
L364:
	;
	v3214 = *(*float64)(unsafe.Add(mBase, uint32(v3151+v3166<<(uint(int32(4))%32))+8))
	if base.B2i32(v3166-v3164 < int32(2))|base.F64_eq(v3146, v3214) == int32(0) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	if base.F64_lt(v3146, v3206) != 0 {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	goto L367
L367:
	;
	v3230 = v3166
	v3231 = v3164
	v3232 = v3166
	v3234 = v3165
	goto L358
L368:
	;
	v3223 = base.I32_div_s(v3165-v3164, int32(2))
	v3230 = int32(-1)
	v3231 = v3164
	v3232 = v3165
	v3234 = v3223 + v3164
	goto L358
L369:
	;
	goto L370
L370:
	;
	v3228 = base.I32_div_s(v3166-v3165, int32(2))
	v3230 = int32(-1)
	v3231 = v3165
	v3232 = v3166
	v3234 = v3165 + v3228
	goto L358
L371:
	;
	goto L357
L372:
	;
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v3347 <= v3230 {
		goto L354
	} else {
		goto L385
	}
L373:
	;
	v3345 = *(*float64)(unsafe.Add(mBase, uint32(v706)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v3241)+8)) = v3345
	goto L372
L374:
	;
	v3252 = v3242 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3242) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v3260 = v3243
	v3264 = v3243
	goto L378
L376:
	;
	v3308 = v3243
	goto L377
L377:
	;
	v3317 = v3308
	v3322 = v3243
	goto L382
L378:
	;
	v3267 = v3260 << (uint(int32(2)) % 32)
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v3241)))
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v3272 = *(*int32)(unsafe.Add(mBase, uint32(v3270+v3267)))
	*(*int32)(unsafe.Add(mBase, uint32(v3267+v3268))) = v3272
	v3274 = int32(4)
	v3275 = v3267 | v3274
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3241)))
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v3278+v3275)))
	*(*int32)(unsafe.Add(mBase, uint32(v3275+v3276))) = v3280
	v3283 = v3267 | int32(8)
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v3241)))
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v3286+v3283)))
	*(*int32)(unsafe.Add(mBase, uint32(v3283+v3284))) = v3288
	v3291 = v3267 | int32(12)
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3241)))
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(v3294+v3291)))
	*(*int32)(unsafe.Add(mBase, uint32(v3291+v3292))) = v3296
	v3299 = v3260 + v3274
	v3301 = v3264 + v3274
	if v3301 != v3242&int32(2147483644) {
		v3260 = v3299
		v3264 = v3301
		goto L378
	} else {
		goto L380
	}
L379:
	;
	if v3252 == int32(0) {
		goto L373
	} else {
		goto L381
	}
L380:
	;
	goto L379
L381:
	;
	v3308 = v3299
	goto L377
L382:
	;
	v3324 = v3317 << (uint(int32(2)) % 32)
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v3241)))
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(v3327+v3324)))
	*(*int32)(unsafe.Add(mBase, uint32(v3324+v3325))) = v3329
	v3331 = int32(1)
	v3334 = v3322 + v3331
	if v3334 != v3252 {
		v3317 = v3317 + v3331
		v3322 = v3334
		goto L382
	} else {
		goto L384
	}
L383:
	;
	goto L373
L384:
	;
	goto L383
L385:
	;
	v3349 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3352 = v3349 + v3347<<(uint(int32(4))%32)
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3352-int32(16))))
	v3358 = *(*float64)(unsafe.Add(mBase, uint32(v3352-int32(8))))
	v3360 = v3230
	v3362 = v3355
	v3386 = v3358
	goto L386
L386:
	;
	v3396 = v3360 << (uint(int32(4)) % 32)
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3398 = v3396 + v3397
	v3399 = *(*int32)(unsafe.Add(mBase, uint32(v3398)))
	*(*int32)(unsafe.Add(mBase, uint32(v3398))) = v3362
	v3401 = *(*float64)(unsafe.Add(mBase, uint32(v3398)+8))
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	*(*float64)(unsafe.Add(mBase, uint32(v3402+v3396)+8)) = v3386
	v3406 = v3360 + int32(1)
	v3407 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v3406 < v3407 {
		v3360 = v3406
		v3362 = v3399
		v3386 = v3401
		goto L386
	} else {
		goto L388
	}
L387:
	;
	goto L354
L388:
	;
	goto L387
L389:
	;
	goto L122
L390:
	;
	if v3487 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L13
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	F_free_attrmap(m, v706)
	mBase = m.M
	v3505 = m.ExcPending
	if v3505 != 0 {
		goto L13
	} else {
		goto L397
	}
L394:
	;
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_4), int32(0))
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
		goto L13
	} else {
		goto L395
	}
L395:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_16), int32(275), int32(_a_F_make_rel_from_joinlist_17))
	mBase = m.M
	v3503 = m.ExcPending
	if v3503 != 0 {
		goto L13
	} else {
		goto L396
	}
L396:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L397:
	;
	F_free_attrmap(m, v709)
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L13
	} else {
		goto L398
	}
L398:
	;
	F_pfree(m, v716)
	mBase = m.M
	v3509 = m.ExcPending
	if v3509 != 0 {
		goto L13
	} else {
		goto L399
	}
L399:
	;
	v3510 = int32(0)
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v3510 < v3512 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v3515 = v3510
	goto L403
L401:
	;
	v3598 = v3511
	goto L402
L402:
	;
	F_pfree(m, v3598)
	mBase = m.M
	v3600 = m.ExcPending
	if v3600 != 0 {
		goto L13
	} else {
		goto L407
	}
L403:
	;
	v3554 = *(*int32)(unsafe.Add(mBase, uint32(v3511+v3515<<(uint(int32(4))%32))))
	F_pfree(m, v3554)
	mBase = m.M
	v3556 = m.ExcPending
	if v3556 != 0 {
		goto L13
	} else {
		goto L405
	}
L404:
	;
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3598 = v3561
	goto L402
L405:
	;
	v3558 = v3515 + int32(1)
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v3558 < v3559 {
		v3515 = v3558
		goto L403
	} else {
		goto L406
	}
L406:
	;
	goto L404
L407:
	;
	F_pfree(m, v382)
	mBase = m.M
	v3602 = m.ExcPending
	if v3602 != 0 {
		goto L13
	} else {
		goto L408
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = int32(0)
	m.G0 = v138 + int32(32)
	v5074 = v3487
	goto L1
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v3616
	*(*int32)(unsafe.Add(mBase, uint32(v3616)+4)) = v114
	if int32(2) <= v43 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v3625 = int32(2)
	goto L413
L411:
	;
	goto L412
L412:
	;
	v5045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v5049 = *(*int32)(unsafe.Add(mBase, uint32(v5045+v43<<(uint(int32(2))%32))))
	if v5049 == int32(0) {
		goto L644
	} else {
		goto L645
	}
L413:
	;
	v3659 = int32(0)
	v3660 = m.G0
	v3662 = v3660 - int32(16)
	m.G0 = v3662
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v3625
	v3665 = int32(2)
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v3669 = v3666 + v3625<<(uint(v3665)%32)
	v3671 = v3669 - int32(4)
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(v3671)))
	if v3672 == v3659 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	goto L412
L415:
	;
	v4112 = int32(2)
	v4113 = v3625 - v4112
	if v4112 <= v4113 {
		goto L506
	} else {
		goto L507
	}
L416:
	;
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v3672)+4))
	if v3675 <= int32(0) {
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v3682 = v3659
	goto L418
L418:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v3672)+12))
	v3718 = *(*int32)(unsafe.Add(mBase, uint32(v3714+v3682<<(uint(int32(2))%32))))
	v3719 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+212))
	if v3719 != 0 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	goto L415
L420:
	;
	v4073 = v3682 + int32(1)
	v4074 = *(*int32)(unsafe.Add(mBase, uint32(v3672)+4))
	if v4073 < v4074 {
		v3682 = v4073
		goto L418
	} else {
		goto L505
	}
L421:
	;
	v3932 = *(*int32)(unsafe.Add(mBase, uint32(v3666)+4))
	if v3932 == int32(0) {
		goto L420
	} else {
		goto L483
	}
L422:
	;
	v3818 = *(*int32)(unsafe.Add(mBase, uint32(v3666)+4))
	if v3818 == int32(0) {
		goto L420
	} else {
		goto L453
	}
L423:
	;
	v3720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3718)+216)))
	if v3720 != 0 {
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v3724 = int32(1)
	v3725 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+64))
	if v3725 != 0 {
		v3807 = v3724
		goto L426
	} else {
		goto L427
	}
L425:
	;
	if v3815 == int32(0) {
		goto L421
	} else {
		goto L452
	}
L426:
	;
	v3815 = v3807
	goto L425
L427:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+104))
	if v3726 != 0 {
		v3807 = v3724
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v3727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v3727 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v3761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3761 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L430:
	;
	v3730 = *(*int32)(unsafe.Add(mBase, uint32(v3727)+4))
	if v3730 <= int32(0) {
		goto L429
	} else {
		goto L431
	}
L431:
	;
	v3736 = int32(0)
	goto L432
L432:
	;
	v3738 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(v3727)+12))
	v3743 = *(*int32)(unsafe.Add(mBase, uint32(v3739+v3736<<(uint(int32(2))%32))))
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v3743)+12))
	v3745 = F_bms_is_subset(m, v3738, v3744)
	mBase = m.M
	if v3745 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	goto L429
L434:
	;
	v3753 = v3736 + int32(1)
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3727)+4))
	if v3753 < v3754 {
		v3736 = v3753
		goto L432
	} else {
		goto L437
	}
L435:
	;
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3749 = *(*int32)(unsafe.Add(mBase, uint32(v3743)+12))
	v3750 = F_bms_equal(m, v3748, v3749)
	mBase = m.M
	if v3750 != 0 {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v3815 = int32(1)
	goto L425
L437:
	;
	goto L433
L438:
	;
	v3807 = int32(0)
	goto L426
L439:
	;
	v3764 = *(*int32)(unsafe.Add(mBase, uint32(v3761)+4))
	if v3764 <= int32(0) {
		goto L438
	} else {
		goto L440
	}
L440:
	;
	v3771 = int32(0)
	goto L441
L441:
	;
	v3773 = *(*int32)(unsafe.Add(mBase, uint32(v3761)+12))
	v3774 = int32(2)
	v3777 = *(*int32)(unsafe.Add(mBase, uint32(v3773+v3771<<(uint(v3774)%32))))
	v3778 = *(*int32)(unsafe.Add(mBase, uint32(v3777)+20))
	if v3778 == v3774 {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	goto L438
L443:
	;
	v3796 = v3771 + int32(1)
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v3761)+4))
	if v3796 < v3797 {
		v3771 = v3796
		goto L441
	} else {
		goto L451
	}
L444:
	;
	v3781 = *(*int32)(unsafe.Add(mBase, uint32(v3777)+4))
	v3782 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3783 = F_bms_is_subset(m, v3781, v3782)
	mBase = m.M
	if v3783 != 0 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v3784 = *(*int32)(unsafe.Add(mBase, uint32(v3777)+8))
	v3785 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3786 = F_bms_is_subset(m, v3784, v3785)
	mBase = m.M
	if v3786 != 0 {
		goto L443
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	v3787 = int32(1)
	v3788 = *(*int32)(unsafe.Add(mBase, uint32(v3777)+4))
	v3789 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3790 = F_bms_overlap(m, v3788, v3789)
	mBase = m.M
	if v3790 != 0 {
		v3807 = v3787
		goto L426
	} else {
		goto L449
	}
L448:
	;
	goto L447
L449:
	;
	v3791 = *(*int32)(unsafe.Add(mBase, uint32(v3777)+8))
	v3792 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3793 = F_bms_overlap(m, v3791, v3792)
	mBase = m.M
	if v3793 != 0 {
		v3807 = v3787
		goto L426
	} else {
		goto L450
	}
L450:
	;
	goto L443
L451:
	;
	goto L442
L452:
	;
	goto L422
L453:
	;
	if v3625 == int32(2) {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v3826 = v3682 + int32(1)
	goto L456
L455:
	;
	v3826 = int32(0)
	goto L456
L456:
	;
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(v3818)+4))
	if v3827 <= v3826 {
		goto L420
	} else {
		goto L457
	}
L457:
	;
	v3834 = v3826
	goto L458
L458:
	;
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3866 = *(*int32)(unsafe.Add(mBase, uint32(v3818)+12))
	v3870 = *(*int32)(unsafe.Add(mBase, uint32(v3866+v3834<<(uint(int32(2))%32))))
	v3871 = *(*int32)(unsafe.Add(mBase, uint32(v3870)+8))
	v3872 = int32(0)
	if base.B2i32(v3865 == v3872)|base.B2i32(v3871 == v3872) != 0 {
		v3917 = v3872
		goto L462
	} else {
		goto L463
	}
L459:
	;
	goto L420
L460:
	;
	v3929 = v3834 + int32(1)
	v3930 = *(*int32)(unsafe.Add(mBase, uint32(v3818)+4))
	if v3929 < v3930 {
		v3834 = v3929
		goto L458
	} else {
		goto L482
	}
L461:
	;
	if v3917 != 0 {
		goto L460
	} else {
		goto L474
	}
L462:
	;
	goto L461
L463:
	;
	v3882 = *(*int32)(unsafe.Add(mBase, uint32(v3865)+4))
	v3883 = *(*int32)(unsafe.Add(mBase, uint32(v3871)+4))
	if v3882 < v3883 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v3885 = v3882
	goto L466
L465:
	;
	v3885 = v3883
	goto L466
L466:
	;
	if v3885 <= int32(1) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v3888 = int32(1)
	goto L469
L468:
	;
	v3888 = v3885
	goto L469
L469:
	;
	v3889 = int32(8)
	v3894 = int32(0)
	goto L470
L470:
	;
	v3901 = v3894 << (uint(int32(2)) % 32)
	v3903 = *(*int32)(unsafe.Add(mBase, uint32(v3871+v3889+v3901)))
	v3905 = *(*int32)(unsafe.Add(mBase, uint32(v3865+v3889+v3901)))
	v3906 = v3903 & v3905
	v3908 = base.B2i32(v3906 != int32(0))
	if v3906 != 0 {
		v3917 = v3908
		goto L462
	} else {
		goto L472
	}
L471:
	;
	v3917 = v3908
	goto L462
L472:
	;
	v3910 = v3894 + int32(1)
	if v3910 != v3888 {
		v3894 = v3910
		goto L470
	} else {
		goto L473
	}
L473:
	;
	goto L471
L474:
	;
	v3918 = F_have_relevant_joinclause(m, l0, v3718, v3870)
	mBase = m.M
	v3919 = m.ExcPending
	if v3919 != 0 {
		goto L13
	} else {
		goto L475
	}
L475:
	;
	if v3918 == int32(0) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v3922 = F_have_join_order_restriction(m, l0, v3718, v3870)
	mBase = m.M
	v3923 = m.ExcPending
	if v3923 != 0 {
		goto L13
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v3926 = F_make_join_rel(m, l0, v3718, v3870)
	mBase = m.M
	v3927 = m.ExcPending
	if v3927 != 0 {
		goto L13
	} else {
		goto L481
	}
L479:
	;
	if v3922 == int32(0) {
		goto L460
	} else {
		goto L480
	}
L480:
	;
	goto L478
L481:
	;
	goto L460
L482:
	;
	goto L459
L483:
	;
	v3935 = int32(0)
	v3936 = *(*int32)(unsafe.Add(mBase, uint32(v3932)+4))
	if v3936 <= v3935 {
		goto L420
	} else {
		goto L484
	}
L484:
	;
	v3944 = v3935
	goto L485
L485:
	;
	v3975 = *(*int32)(unsafe.Add(mBase, uint32(v3932)+12))
	v3979 = *(*int32)(unsafe.Add(mBase, uint32(v3975+v3944<<(uint(int32(2))%32))))
	v3980 = *(*int32)(unsafe.Add(mBase, uint32(v3979)+8))
	v3981 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+8))
	v3982 = int32(0)
	if base.B2i32(v3980 == v3982)|base.B2i32(v3981 == v3982) != 0 {
		v4027 = v3982
		goto L488
	} else {
		goto L489
	}
L486:
	;
	goto L420
L487:
	;
	if v4027 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L488:
	;
	goto L487
L489:
	;
	v3992 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+4))
	v3993 = *(*int32)(unsafe.Add(mBase, uint32(v3981)+4))
	if v3992 < v3993 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v3995 = v3992
	goto L492
L491:
	;
	v3995 = v3993
	goto L492
L492:
	;
	if v3995 <= int32(1) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v3998 = int32(1)
	goto L495
L494:
	;
	v3998 = v3995
	goto L495
L495:
	;
	v3999 = int32(8)
	v4004 = int32(0)
	goto L496
L496:
	;
	v4011 = v4004 << (uint(int32(2)) % 32)
	v4013 = *(*int32)(unsafe.Add(mBase, uint32(v3981+v3999+v4011)))
	v4015 = *(*int32)(unsafe.Add(mBase, uint32(v3980+v3999+v4011)))
	v4016 = v4013 & v4015
	v4018 = base.B2i32(v4016 != int32(0))
	if v4016 != 0 {
		v4027 = v4018
		goto L488
	} else {
		goto L498
	}
L497:
	;
	v4027 = v4018
	goto L488
L498:
	;
	v4020 = v4004 + int32(1)
	if v4020 != v3998 {
		v4004 = v4020
		goto L496
	} else {
		goto L499
	}
L499:
	;
	goto L497
L500:
	;
	v4030 = F_make_join_rel(m, l0, v3718, v3979)
	mBase = m.M
	v4031 = m.ExcPending
	if v4031 != 0 {
		goto L13
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	v4033 = v3944 + int32(1)
	v4034 = *(*int32)(unsafe.Add(mBase, uint32(v3932)+4))
	if v4033 < v4034 {
		v3944 = v4033
		goto L485
	} else {
		goto L504
	}
L503:
	;
	goto L502
L504:
	;
	goto L486
L505:
	;
	goto L419
L506:
	;
	v4119 = v4113
	v4125 = v3665
	goto L509
L507:
	;
	goto L508
L508:
	;
	v4534 = *(*int32)(unsafe.Add(mBase, uint32(v3669)))
	if v4534 != 0 {
		goto L580
	} else {
		goto L581
	}
L509:
	;
	v4155 = *(*int32)(unsafe.Add(mBase, uint32(v3666+v4125<<(uint(int32(2))%32))))
	if v4155 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	goto L508
L511:
	;
	v4495 = v4125 + int32(1)
	v4496 = v3625 - v4495
	if v4495 <= v4496 {
		v4119 = v4496
		v4125 = v4495
		goto L509
	} else {
		goto L579
	}
L512:
	;
	v4158 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+4))
	if v4158 <= int32(0) {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v4169 = int32(0)
	goto L514
L514:
	;
	v4201 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+12))
	v4205 = *(*int32)(unsafe.Add(mBase, uint32(v4201+v4169<<(uint(int32(2))%32))))
	v4206 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+212))
	if v4206 != 0 {
		goto L517
	} else {
		goto L518
	}
L515:
	;
	goto L511
L516:
	;
	v4455 = v4169 + int32(1)
	v4456 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+4))
	if v4455 < v4456 {
		v4169 = v4455
		goto L514
	} else {
		goto L578
	}
L517:
	;
	v4305 = *(*int32)(unsafe.Add(mBase, uint32(v3666+v4119<<(uint(int32(2))%32))))
	if v4305 == int32(0) {
		goto L516
	} else {
		goto L548
	}
L518:
	;
	v4207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4205)+216)))
	if v4207 != 0 {
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v4211 = int32(1)
	v4212 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+64))
	if v4212 != 0 {
		v4294 = v4211
		goto L521
	} else {
		goto L522
	}
L520:
	;
	if v4302 == int32(0) {
		goto L516
	} else {
		goto L547
	}
L521:
	;
	v4302 = v4294
	goto L520
L522:
	;
	v4213 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+104))
	if v4213 != 0 {
		v4294 = v4211
		goto L521
	} else {
		goto L523
	}
L523:
	;
	v4214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v4214 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v4248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v4248 == int32(0) {
		goto L533
	} else {
		goto L534
	}
L525:
	;
	v4217 = *(*int32)(unsafe.Add(mBase, uint32(v4214)+4))
	if v4217 <= int32(0) {
		goto L524
	} else {
		goto L526
	}
L526:
	;
	v4223 = int32(0)
	goto L527
L527:
	;
	v4225 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4226 = *(*int32)(unsafe.Add(mBase, uint32(v4214)+12))
	v4230 = *(*int32)(unsafe.Add(mBase, uint32(v4226+v4223<<(uint(int32(2))%32))))
	v4231 = *(*int32)(unsafe.Add(mBase, uint32(v4230)+12))
	v4232 = F_bms_is_subset(m, v4225, v4231)
	mBase = m.M
	if v4232 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L528:
	;
	goto L524
L529:
	;
	v4240 = v4223 + int32(1)
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v4214)+4))
	if v4240 < v4241 {
		v4223 = v4240
		goto L527
	} else {
		goto L532
	}
L530:
	;
	v4235 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4236 = *(*int32)(unsafe.Add(mBase, uint32(v4230)+12))
	v4237 = F_bms_equal(m, v4235, v4236)
	mBase = m.M
	if v4237 != 0 {
		goto L529
	} else {
		goto L531
	}
L531:
	;
	v4302 = int32(1)
	goto L520
L532:
	;
	goto L528
L533:
	;
	v4294 = int32(0)
	goto L521
L534:
	;
	v4251 = *(*int32)(unsafe.Add(mBase, uint32(v4248)+4))
	if v4251 <= int32(0) {
		goto L533
	} else {
		goto L535
	}
L535:
	;
	v4258 = int32(0)
	goto L536
L536:
	;
	v4260 = *(*int32)(unsafe.Add(mBase, uint32(v4248)+12))
	v4261 = int32(2)
	v4264 = *(*int32)(unsafe.Add(mBase, uint32(v4260+v4258<<(uint(v4261)%32))))
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+20))
	if v4265 == v4261 {
		goto L538
	} else {
		goto L539
	}
L537:
	;
	goto L533
L538:
	;
	v4283 = v4258 + int32(1)
	v4284 = *(*int32)(unsafe.Add(mBase, uint32(v4248)+4))
	if v4283 < v4284 {
		v4258 = v4283
		goto L536
	} else {
		goto L546
	}
L539:
	;
	v4268 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+4))
	v4269 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4270 = F_bms_is_subset(m, v4268, v4269)
	mBase = m.M
	if v4270 != 0 {
		goto L540
	} else {
		goto L541
	}
L540:
	;
	v4271 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+8))
	v4272 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4273 = F_bms_is_subset(m, v4271, v4272)
	mBase = m.M
	if v4273 != 0 {
		goto L538
	} else {
		goto L543
	}
L541:
	;
	goto L542
L542:
	;
	v4274 = int32(1)
	v4275 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+4))
	v4276 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4277 = F_bms_overlap(m, v4275, v4276)
	mBase = m.M
	if v4277 != 0 {
		v4294 = v4274
		goto L521
	} else {
		goto L544
	}
L543:
	;
	goto L542
L544:
	;
	v4278 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+8))
	v4279 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4280 = F_bms_overlap(m, v4278, v4279)
	mBase = m.M
	if v4280 != 0 {
		v4294 = v4274
		goto L521
	} else {
		goto L545
	}
L545:
	;
	goto L538
L546:
	;
	goto L537
L547:
	;
	goto L517
L548:
	;
	if v4119 == v4125 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v4312 = v4169 + int32(1)
	goto L551
L550:
	;
	v4312 = int32(0)
	goto L551
L551:
	;
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(v4305)+4))
	if v4313 <= v4312 {
		goto L516
	} else {
		goto L552
	}
L552:
	;
	v4320 = v4312
	goto L553
L553:
	;
	v4351 = *(*int32)(unsafe.Add(mBase, uint32(v4205)+8))
	v4352 = *(*int32)(unsafe.Add(mBase, uint32(v4305)+12))
	v4356 = *(*int32)(unsafe.Add(mBase, uint32(v4352+v4320<<(uint(int32(2))%32))))
	v4357 = *(*int32)(unsafe.Add(mBase, uint32(v4356)+8))
	v4358 = int32(0)
	if base.B2i32(v4351 == v4358)|base.B2i32(v4357 == v4358) != 0 {
		v4403 = v4358
		goto L557
	} else {
		goto L558
	}
L554:
	;
	goto L516
L555:
	;
	v4415 = v4320 + int32(1)
	v4416 = *(*int32)(unsafe.Add(mBase, uint32(v4305)+4))
	if v4415 < v4416 {
		v4320 = v4415
		goto L553
	} else {
		goto L577
	}
L556:
	;
	if v4403 != 0 {
		goto L555
	} else {
		goto L569
	}
L557:
	;
	goto L556
L558:
	;
	v4368 = *(*int32)(unsafe.Add(mBase, uint32(v4351)+4))
	v4369 = *(*int32)(unsafe.Add(mBase, uint32(v4357)+4))
	if v4368 < v4369 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v4371 = v4368
	goto L561
L560:
	;
	v4371 = v4369
	goto L561
L561:
	;
	if v4371 <= int32(1) {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v4374 = int32(1)
	goto L564
L563:
	;
	v4374 = v4371
	goto L564
L564:
	;
	v4375 = int32(8)
	v4380 = int32(0)
	goto L565
L565:
	;
	v4387 = v4380 << (uint(int32(2)) % 32)
	v4389 = *(*int32)(unsafe.Add(mBase, uint32(v4357+v4375+v4387)))
	v4391 = *(*int32)(unsafe.Add(mBase, uint32(v4351+v4375+v4387)))
	v4392 = v4389 & v4391
	v4394 = base.B2i32(v4392 != int32(0))
	if v4392 != 0 {
		v4403 = v4394
		goto L557
	} else {
		goto L567
	}
L566:
	;
	v4403 = v4394
	goto L557
L567:
	;
	v4396 = v4380 + int32(1)
	if v4396 != v4374 {
		v4380 = v4396
		goto L565
	} else {
		goto L568
	}
L568:
	;
	goto L566
L569:
	;
	v4404 = F_have_relevant_joinclause(m, l0, v4205, v4356)
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L13
	} else {
		goto L570
	}
L570:
	;
	if v4404 == int32(0) {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v4408 = F_have_join_order_restriction(m, l0, v4205, v4356)
	mBase = m.M
	v4409 = m.ExcPending
	if v4409 != 0 {
		goto L13
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	v4412 = F_make_join_rel(m, l0, v4205, v4356)
	mBase = m.M
	v4413 = m.ExcPending
	if v4413 != 0 {
		goto L13
	} else {
		goto L576
	}
L574:
	;
	if v4408 == int32(0) {
		goto L555
	} else {
		goto L575
	}
L575:
	;
	goto L573
L576:
	;
	goto L555
L577:
	;
	goto L554
L578:
	;
	goto L515
L579:
	;
	goto L510
L580:
	;
	m.G0 = v3662 + int32(16)
	v4852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v4856 = *(*int32)(unsafe.Add(mBase, uint32(v4852+v3625<<(uint(int32(2))%32))))
	if v4856 == int32(0) {
		goto L620
	} else {
		goto L621
	}
L581:
	;
	v4535 = *(*int32)(unsafe.Add(mBase, uint32(v3671)))
	if v4535 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v4536 = *(*int32)(unsafe.Add(mBase, uint32(v4535)+4))
	if int32(0) < v4536 {
		goto L585
	} else {
		goto L586
	}
L583:
	;
	goto L584
L584:
	;
	v4798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v4798 != 0 {
		goto L580
	} else {
		goto L615
	}
L585:
	;
	v4544 = int32(0)
	goto L588
L586:
	;
	goto L587
L587:
	;
	v4761 = *(*int32)(unsafe.Add(mBase, uint32(v3669)))
	if v4761 != 0 {
		goto L580
	} else {
		goto L614
	}
L588:
	;
	v4576 = *(*int32)(unsafe.Add(mBase, uint32(v3666)+4))
	if v4576 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L589:
	;
	goto L587
L590:
	;
	v4722 = v4544 + int32(1)
	v4723 = *(*int32)(unsafe.Add(mBase, uint32(v4535)+4))
	if v4722 < v4723 {
		v4544 = v4722
		goto L588
	} else {
		goto L613
	}
L591:
	;
	v4579 = *(*int32)(unsafe.Add(mBase, uint32(v4576)+4))
	if v4579 <= int32(0) {
		goto L590
	} else {
		goto L592
	}
L592:
	;
	v4582 = *(*int32)(unsafe.Add(mBase, uint32(v4535)+12))
	v4586 = *(*int32)(unsafe.Add(mBase, uint32(v4582+v4544<<(uint(int32(2))%32))))
	v4593 = int32(0)
	goto L593
L593:
	;
	v4624 = *(*int32)(unsafe.Add(mBase, uint32(v4576)+12))
	v4628 = *(*int32)(unsafe.Add(mBase, uint32(v4624+v4593<<(uint(int32(2))%32))))
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v4628)+8))
	v4630 = *(*int32)(unsafe.Add(mBase, uint32(v4586)+8))
	v4631 = int32(0)
	if base.B2i32(v4629 == v4631)|base.B2i32(v4630 == v4631) != 0 {
		v4676 = v4631
		goto L596
	} else {
		goto L597
	}
L594:
	;
	goto L590
L595:
	;
	if v4676 == int32(0) {
		goto L608
	} else {
		goto L609
	}
L596:
	;
	goto L595
L597:
	;
	v4641 = *(*int32)(unsafe.Add(mBase, uint32(v4629)+4))
	v4642 = *(*int32)(unsafe.Add(mBase, uint32(v4630)+4))
	if v4641 < v4642 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v4644 = v4641
	goto L600
L599:
	;
	v4644 = v4642
	goto L600
L600:
	;
	if v4644 <= int32(1) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v4647 = int32(1)
	goto L603
L602:
	;
	v4647 = v4644
	goto L603
L603:
	;
	v4648 = int32(8)
	v4653 = int32(0)
	goto L604
L604:
	;
	v4660 = v4653 << (uint(int32(2)) % 32)
	v4662 = *(*int32)(unsafe.Add(mBase, uint32(v4630+v4648+v4660)))
	v4664 = *(*int32)(unsafe.Add(mBase, uint32(v4629+v4648+v4660)))
	v4665 = v4662 & v4664
	v4667 = base.B2i32(v4665 != int32(0))
	if v4665 != 0 {
		v4676 = v4667
		goto L596
	} else {
		goto L606
	}
L605:
	;
	v4676 = v4667
	goto L596
L606:
	;
	v4669 = v4653 + int32(1)
	if v4669 != v4647 {
		v4653 = v4669
		goto L604
	} else {
		goto L607
	}
L607:
	;
	goto L605
L608:
	;
	v4679 = F_make_join_rel(m, l0, v4586, v4628)
	mBase = m.M
	v4680 = m.ExcPending
	if v4680 != 0 {
		goto L13
	} else {
		goto L611
	}
L609:
	;
	goto L610
L610:
	;
	v4682 = v4593 + int32(1)
	v4683 = *(*int32)(unsafe.Add(mBase, uint32(v4576)+4))
	if v4682 < v4683 {
		v4593 = v4682
		goto L593
	} else {
		goto L612
	}
L611:
	;
	goto L610
L612:
	;
	goto L594
L613:
	;
	goto L589
L614:
	;
	goto L584
L615:
	;
	v4799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v4799 != 0 {
		goto L580
	} else {
		goto L616
	}
L616:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4803 = m.ExcPending
	if v4803 != 0 {
		goto L13
	} else {
		goto L617
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3662))) = v3625
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_18), v3662)
	mBase = m.M
	v4807 = m.ExcPending
	if v4807 != 0 {
		goto L13
	} else {
		goto L618
	}
L618:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_19), int32(255), int32(_a_F_make_rel_from_joinlist_20))
	mBase = m.M
	v4812 = m.ExcPending
	if v4812 != 0 {
		goto L13
	} else {
		goto L619
	}
L619:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L620:
	;
	v5007 = v3625 + int32(1)
	if v5007 <= v43 {
		v3625 = v5007
		goto L413
	} else {
		goto L643
	}
L621:
	;
	v4859 = int32(0)
	v4860 = *(*int32)(unsafe.Add(mBase, uint32(v4856)+4))
	if v4860 <= v4859 {
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v4864 = v4859
	goto L623
L623:
	;
	v4899 = *(*int32)(unsafe.Add(mBase, uint32(v4856)+12))
	v4903 = *(*int32)(unsafe.Add(mBase, uint32(v4899+v4864<<(uint(int32(2))%32))))
	F_generate_partitionwise_join_paths(m, l0, v4903)
	mBase = m.M
	v4905 = m.ExcPending
	if v4905 != 0 {
		goto L13
	} else {
		goto L625
	}
L624:
	;
	goto L620
L625:
	;
	v4906 = *(*int32)(unsafe.Add(mBase, uint32(v4903)+8))
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v4908 = int32(0)
	if base.B2i32(v4906 == v4908)|base.B2i32(v4907 == v4908) != 0 {
		v4954 = base.B2i32(v4906|v4907 == v4908)
		goto L627
	} else {
		goto L628
	}
L626:
	;
	if v4954 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L627:
	;
	goto L626
L628:
	;
	v4922 = *(*int32)(unsafe.Add(mBase, uint32(v4906)+4))
	v4923 = *(*int32)(unsafe.Add(mBase, uint32(v4907)+4))
	if v4922 != v4923 {
		v4954 = int32(0)
		goto L627
	} else {
		goto L629
	}
L629:
	;
	v4925 = int32(1)
	if v4922 <= v4925 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v4928 = v4925
	goto L632
L631:
	;
	v4928 = v4922
	goto L632
L632:
	;
	v4929 = int32(8)
	v4934 = int32(0)
	goto L633
L633:
	;
	v4942 = v4934 << (uint(int32(2)) % 32)
	v4944 = *(*int32)(unsafe.Add(mBase, uint32(v4906+v4929+v4942)))
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v4907+v4929+v4942)))
	v4947 = base.B2i32(v4944 == v4946)
	if v4944 != v4946 {
		v4954 = v4947
		goto L627
	} else {
		goto L635
	}
L634:
	;
	v4954 = v4947
	goto L627
L635:
	;
	v4950 = v4934 + int32(1)
	if v4950 != v4928 {
		v4934 = v4950
		goto L633
	} else {
		goto L636
	}
L636:
	;
	goto L634
L637:
	;
	F_generate_useful_gather_paths(m, l0, v4903, int32(0))
	mBase = m.M
	v4963 = m.ExcPending
	if v4963 != 0 {
		goto L13
	} else {
		goto L640
	}
L638:
	;
	goto L639
L639:
	;
	F_set_cheapest(m, v4903)
	mBase = m.M
	v4965 = m.ExcPending
	if v4965 != 0 {
		goto L13
	} else {
		goto L641
	}
L640:
	;
	goto L639
L641:
	;
	v4967 = v4864 + int32(1)
	v4968 = *(*int32)(unsafe.Add(mBase, uint32(v4856)+4))
	if v4967 < v4968 {
		v4864 = v4967
		goto L623
	} else {
		goto L642
	}
L642:
	;
	goto L624
L643:
	;
	goto L414
L644:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L13
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	v5065 = *(*int32)(unsafe.Add(mBase, uint32(v5049)+12))
	v5066 = *(*int32)(unsafe.Add(mBase, uint32(v5065)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	m.G0 = v3610 + int32(16)
	v5074 = v5066
	goto L1
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3610))) = v43
	F_errmsg_internal(m, int32(_a_F_make_rel_from_joinlist_18), v3610)
	mBase = m.M
	v5059 = m.ExcPending
	if v5059 != 0 {
		goto L13
	} else {
		goto L648
	}
L648:
	;
	F_errfinish(m, int32(_a_F_make_rel_from_joinlist_1), int32(3533), int32(_a_F_make_rel_from_joinlist_21))
	mBase = m.M
	v5064 = m.ExcPending
	if v5064 != 0 {
		goto L13
	} else {
		goto L649
	}
L649:
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
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
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
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
										v73 = F_is_parallel_safe(m, l0, v72)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											if v73 == int32(0) {
												return
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
												v79 = F_is_parallel_safe(m, l0, v78)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													if v79 == int32(0) {
													} else {
														v83 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
													}
													return
												}
											}
										}
									} else {
										v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
										v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+140))
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
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
													v73 = F_is_parallel_safe(m, l0, v72)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return
													} else {
														if v73 == int32(0) {
															return
														} else {
															v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
															v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
															v79 = F_is_parallel_safe(m, l0, v78)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																if v79 == int32(0) {
																} else {
																	v83 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
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
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
						v73 = F_is_parallel_safe(m, l0, v72)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							if v73 == int32(0) {
								return
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
								v79 = F_is_parallel_safe(m, l0, v78)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									if v79 == int32(0) {
									} else {
										v83 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
									}
									return
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+140))
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
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
									v73 = F_is_parallel_safe(m, l0, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										if v73 == int32(0) {
											return
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
											v79 = F_is_parallel_safe(m, l0, v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												if v79 == int32(0) {
												} else {
													v83 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
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
				v57 = v35
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
				if v39 != int32(1) {
					v57 = v35
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
					if v43 == int32(0) {
						v57 = int32(0)
					} else {
						v46 = int32(1)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v47 != int32(7) {
							v57 = v46
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
							if v50 != 0 {
								v57 = int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
								v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
								if v52 != int64(0) {
									v57 = v46
								} else {
									v57 = int32(0)
								}
							}
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
			if v43 == int32(0) {
				v57 = int32(0)
			} else {
				v46 = int32(1)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47 != int32(7) {
					v57 = v46
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
					if v50 != 0 {
						v57 = int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
						if v52 != int64(0) {
							v57 = v46
						} else {
							v57 = int32(0)
						}
					}
				}
			}
		}
		if v57 == int32(0) {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
			v73 = F_is_parallel_safe(m, l0, v72)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				if v73 == int32(0) {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
					v79 = F_is_parallel_safe(m, l0, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						if v79 == int32(0) {
						} else {
							v83 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
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
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
		v62 = F_is_parallel_safe(m, l0, v61)
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return
		} else {
			if v62 != 0 {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
				v73 = F_is_parallel_safe(m, l0, v72)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					if v73 == int32(0) {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
						v79 = F_is_parallel_safe(m, l0, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							if v79 == int32(0) {
							} else {
								v83 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
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
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
		v65 = F_is_parallel_safe(m, l0, v64)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			if v65 == int32(0) {
				return
			} else {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
				v73 = F_is_parallel_safe(m, l0, v72)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					if v73 == int32(0) {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
						v79 = F_is_parallel_safe(m, l0, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							if v79 == int32(0) {
							} else {
								v83 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
							}
							return
						}
					}
				}
			}
		}
	default:
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
		v73 = F_is_parallel_safe(m, l0, v72)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			if v73 == int32(0) {
				return
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
				v79 = F_is_parallel_safe(m, l0, v78)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					if v79 == int32(0) {
					} else {
						v83 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v83)
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
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
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 float64
	_ = v459
	var v461 int32
	_ = v461
	var v462 int64
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 float64
	_ = v471
	var v472 float64
	_ = v472
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 float64
	_ = v527
	var v528 float64
	_ = v528
	var v546 float64
	_ = v546
	var v556 float64
	_ = v556
	var v557 float64
	_ = v557
	var v559 float64
	_ = v559
	var v561 float64
	_ = v561
	var v562 float64
	_ = v562
	var v580 float64
	_ = v580
	var v590 float64
	_ = v590
	var v591 float64
	_ = v591
	var v593 float64
	_ = v593
	var v594 int32
	_ = v594
	var v595 float64
	_ = v595
	var v596 float64
	_ = v596
	var v600 float64
	_ = v600
	var v603 float64
	_ = v603
	var v605 float64
	_ = v605
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 float64
	_ = v664
	var v666 int32
	_ = v666
	var v667 int64
	_ = v667
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 float64
	_ = v676
	var v677 float64
	_ = v677
	var v678 int32
	_ = v678
	var v679 int64
	_ = v679
	var v688 int32
	_ = v688
	var v697 int32
	_ = v697
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 float64
	_ = v732
	var v733 float64
	_ = v733
	var v751 float64
	_ = v751
	var v761 float64
	_ = v761
	var v762 float64
	_ = v762
	var v764 float64
	_ = v764
	var v766 float64
	_ = v766
	var v767 float64
	_ = v767
	var v785 float64
	_ = v785
	var v795 float64
	_ = v795
	var v796 float64
	_ = v796
	var v798 float64
	_ = v798
	var v799 int32
	_ = v799
	var v800 float64
	_ = v800
	var v801 float64
	_ = v801
	var v805 float64
	_ = v805
	var v808 float64
	_ = v808
	var v810 float64
	_ = v810
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 float64
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 float64
	_ = v853
	var v854 int64
	_ = v854
	var v863 int32
	_ = v863
	var v871 int32
	_ = v871
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 float64
	_ = v906
	var v907 float64
	_ = v907
	var v908 float64
	_ = v908
	var v926 float64
	_ = v926
	var v927 float64
	_ = v927
	var v936 float64
	_ = v936
	var v937 float64
	_ = v937
	var v939 float64
	_ = v939
	var v941 float64
	_ = v941
	var v943 float64
	_ = v943
	var v945 float64
	_ = v945
	var v946 float64
	_ = v946
	var v964 float64
	_ = v964
	var v965 float64
	_ = v965
	var v968 float64
	_ = v968
	var v974 float64
	_ = v974
	var v975 float64
	_ = v975
	var v977 float64
	_ = v977
	var v978 int32
	_ = v978
	var v979 float64
	_ = v979
	var v980 float64
	_ = v980
	var v983 float64
	_ = v983
	var v985 float64
	_ = v985
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1144 float64
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 float64
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 float64
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int64
	_ = v1153
	var v1160 int32
	_ = v1160
	var v1172 int32
	_ = v1172
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 float64
	_ = v1203
	var v1222 float64
	_ = v1222
	var v1231 int32
	_ = v1231
	var v1237 int32
	_ = v1237
	var v1238 float64
	_ = v1238
	var v1239 float64
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int64
	_ = v1241
	var v1250 int32
	_ = v1250
	var v1263 int32
	_ = v1263
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 float64
	_ = v1294
	var v1295 float64
	_ = v1295
	var v1313 float64
	_ = v1313
	var v1323 float64
	_ = v1323
	var v1324 float64
	_ = v1324
	var v1326 float64
	_ = v1326
	var v1328 float64
	_ = v1328
	var v1329 float64
	_ = v1329
	var v1347 float64
	_ = v1347
	var v1357 float64
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 float64
	_ = v1359
	var v1361 float64
	_ = v1361
	var v1362 float64
	_ = v1362
	var v1366 float64
	_ = v1366
	var v1368 float64
	_ = v1368
	var v1370 float64
	_ = v1370
	var v1379 float64
	_ = v1379
	var v1384 float64
	_ = v1384
	var v1397 int32
	_ = v1397
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1472 int32
	_ = v1472
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1566 int32
	_ = v1566
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1694 int32
	_ = v1694
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1779 int32
	_ = v1779
	var v1786 int32
	_ = v1786
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1835 int32
	_ = v1835
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1872 int32
	_ = v1872
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1911 int32
	_ = v1911
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2024 int32
	_ = v2024
	var v2029 int32
	_ = v2029
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2080 int32
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2100 int32
	_ = v2100
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2120 int32
	_ = v2120
	var v2133 int32
	_ = v2133
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2183 float64
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2257 int32
	_ = v2257
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2376 int32
	_ = v2376
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2408 int32
	_ = v2408
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2494 int32
	_ = v2494
	var v2501 int32
	_ = v2501
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2519 int32
	_ = v2519
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2541 int32
	_ = v2541
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2571 float64
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2625 int32
	_ = v2625
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2652 int32
	_ = v2652
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2699 int32
	_ = v2699
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2719 int32
	_ = v2719
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	v5 = int32(0)
	v18 = float64(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v34 == v5 {
		v55 = v5
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[0]))
	if v2667 != 0 {
		goto L403
	} else {
		goto L404
	}
L2:
	;
	if v55 != 0 {
		v2639 = l0
		v2640 = l1
		v2641 = l2
		v2642 = l3
		v2652 = v30
		goto L1
	} else {
		goto L12
	}
L3:
	;
	goto L2
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v38 = v37
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if base.Ui32(int32(2)) <= base.Ui32(v42-int32(301)) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v55 = int32(1)
	goto L3
L7:
	;
	if v42 != int32(290) {
		v55 = v5
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v38 = v41 + int32(72)
	goto L5
L9:
	;
	goto L6
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v49 != 0 {
		v55 = v5
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v56 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v59 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v181 {
	case 0:
		goto L51
	case 1, 6, 7, 8:
		v2639 = l0
		v2640 = l1
		v2641 = l2
		v2642 = l3
		v2652 = v30
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
	v64 = m.ExcPending
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if int32(0) < v65 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	return
L20:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L21:
	;
	v73 = v5
	v74 = v5
	goto L24
L22:
	;
	v157 = v5
	goto L23
L23:
	;
	F_add_paths_to_append_rel(m, l0, l1, v157)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L19
	} else {
		goto L45
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v74<<(uint(int32(2))%32))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v100 != l2 {
		v145 = v73
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v157 = v145
	goto L23
L26:
	;
	v149 = v74 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v149 < v150 {
		v73 = v145
		v74 = v149
		goto L24
	} else {
		goto L44
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v104 = v102 << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104+v105)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108+v104)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v111 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+26)) = uint8(v114)
	goto L30
L29:
	;
	goto L30
L30:
	;
	F_set_rel_pathlist(m, l0, v107, v102, v110)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v118 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
	if v120 == v118 {
		v141 = v118
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v141 != 0 {
		v145 = v73
		goto L26
	} else {
		goto L42
	}
L33:
	;
	goto L32
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v124 = v123
	goto L35
L35:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if base.Ui32(int32(2)) <= base.Ui32(v128-int32(301)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v141 = int32(1)
	goto L33
L37:
	;
	if v128 != int32(290) {
		v141 = v118
		goto L33
	} else {
		goto L40
	}
L38:
	;
	v124 = v127 + int32(72)
	goto L35
L39:
	;
	goto L36
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+72))
	if v135 != 0 {
		v141 = v118
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v142 = F_lappend(m, v73, v107)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v145 = v142
	goto L26
L44:
	;
	goto L25
L45:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L46:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1016 = m.G0
	v1018 = v1016 - int32(16)
	m.G0 = v1018
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v1023 = F_TidQualFromRestrictInfoList(m, l0, v1020, l1, v1018+int32(15))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L19
	} else {
		goto L173
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L19
	} else {
		goto L170
	}
L48:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v826 = F_palloc0(m, int32(72))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L19
	} else {
		goto L156
	}
L49:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v621 = F_palloc0(m, int32(72))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L19
	} else {
		goto L134
	}
L50:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+72)))
	if v278 != int32(1) {
		v395 = v5
		goto L85
	} else {
		goto L86
	}
L51:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v182 == int32(102) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	m.T0[v187].(func(*base.Module, int32, int32, int32))(m, l0, l1, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v190 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L55:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L56:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v195 = F_palloc0(m, int32(72))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v195))) = int64(1460288880919)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = v200
	v202 = F_get_baserel_parampathinfo(m, l0, l1, v193)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+20)) = uint8(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v202
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+64)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v195)+24)) = v204
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+21)) = uint8(v207)
	F_cost_samplescan(m, v195, l0, l1, v202)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v215) <= base.Ui32(int32(1)) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	F_add_path(m, l1, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L19
	} else {
		goto L84
	}
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v219 = int32(0)
	if v218 == v219 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v269 = F_GetTsmRoutine(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L19
	} else {
		goto L81
	}
L64:
	;
	if v264 == int32(1) {
		v274 = v195
		goto L60
	} else {
		goto L80
	}
L65:
	;
	v264 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v227 = int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v228 <= v227 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v231 = v227
	goto L70
L69:
	;
	v231 = v228
	goto L70
L70:
	;
	v235 = int32(0)
	v237 = v219
	goto L71
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v218+int32(8)+v235<<(uint(int32(2))%32))))
	if v244 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v264 = v256
	goto L64
L73:
	;
	goto L72
L74:
	;
	v245 = int32(2)
	if v237 != 0 {
		v256 = v245
		goto L73
	} else {
		goto L77
	}
L75:
	;
	v251 = v237
	goto L76
L76:
	;
	v253 = v235 + int32(1)
	if v253 != v231 {
		v235 = v253
		v237 = v251
		goto L71
	} else {
		goto L79
	}
L77:
	;
	v246 = int32(1)
	if base.Ui32(v246) < base.Ui32(base.I32_popcnt(v244)) {
		v256 = v245
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v251 = v246
	goto L76
L79:
	;
	v256 = v251
	goto L73
L80:
	;
	goto L63
L81:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+9)))
	if v271 != 0 {
		v274 = v195
		goto L60
	} else {
		goto L82
	}
L82:
	;
	v272 = F_create_material_path(m, l1, v195)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L19
	} else {
		goto L83
	}
L83:
	;
	v274 = v272
	goto L60
L84:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L85:
	;
	v417 = F_palloc0(m, int32(72))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L19
	} else {
		goto L112
	}
L86:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v282 == int32(0) {
		v395 = v5
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v285 <= int32(0) {
		v395 = v5
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
	v296 = v5
	goto L90
L89:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v338 = m.G0
	v340 = v338 - int32(48)
	m.G0 = v340
	v349 = F_get_ordering_op_properties(m, int32(412), v340+int32(44), v340+int32(40), v340+int32(36))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L19
	} else {
		goto L99
	}
L90:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v288+v296<<(uint(int32(2))%32))))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v321 != int32(6) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v395 = int32(0)
	goto L85
L92:
	;
	v333 = v296 + int32(1)
	if v285 != v333 {
		v296 = v333
		goto L90
	} else {
		goto L97
	}
L93:
	;
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320)+8)))
	if v324 != v289 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v326 != v327 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v320)+28))
	if v329 == int32(0) {
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
	v395 = v369
	goto L85
L99:
	;
	if v349 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v340)+44))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v340)+40))
	v353 = F_exprCollation(m, v320)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
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
	v376 = m.ExcPending
	if v376 != 0 {
		goto L19
	} else {
		goto L109
	}
L103:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v340)+36))
	v357 = base.B2i32(v355 == int32(5))
	v358 = int32(0)
	v360 = F_make_pathkey_from_sortinfo(m, l0, v320, v351, v352, v353, v357, v357, v358, v336, v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L19
	} else {
		goto L104
	}
L104:
	;
	if v360 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+12)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v340)+32)) = v360
	v367 = F_list_make1_impl(m, int32(1), v340+int32(12))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L108
	}
L106:
	;
	v369 = int32(0)
	goto L107
L107:
	;
	m.G0 = v340 + int32(48)
	goto L98
L108:
	;
	v369 = v367
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+16)) = int32(412)
	F_errmsg_internal(m, int32(_a_F_set_rel_pathlist_0), v340+int32(16))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_set_rel_pathlist_1), int32(1016), int32(_a_F_set_rel_pathlist_2))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v417)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v417))) = int64(1494648619287)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+12)) = v422
	v424 = F_get_baserel_parampathinfo(m, l0, l1, v277)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L19
	} else {
		goto L113
	}
L113:
	;
	v426 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+20)) = uint8(v426)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+16)) = v424
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+64)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v417)+24)) = v426
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+21)) = uint8(v429)
	v434 = m.G0
	v436 = v434 - int32(32)
	m.G0 = v436
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v438 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	if v424 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v452 = v438 + v439<<(uint(int32(2))%32)
	goto L114
L116:
	;
	goto L117
L117:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+52))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v452 = v445 + v446<<(uint(int32(2))%32) - int32(4)
	goto L114
L118:
	;
	v458 = v424 + int32(8)
	goto L120
L119:
	;
	v458 = l1 + int32(16)
	goto L120
L120:
	;
	v459 = *(*float64)(unsafe.Add(mBase, uint32(v458)))
	*(*float64)(unsafe.Add(mBase, uint32(v417)+32)) = v459
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v453)+68))
	v462 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v436)+16)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v436)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v436)+24)) = v462
	v469 = F_cost_qual_eval_walker(m, v461, v436+int32(8))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L19
	} else {
		goto L121
	}
L121:
	;
	v471 = *(*float64)(unsafe.Add(mBase, uint32(v436)+24))
	v472 = *(*float64)(unsafe.Add(mBase, uint32(v436)+16))
	if v424 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v591 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v593 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v595 = *(*float64)(unsafe.Add(mBase, uint32(v594)+24))
	v596 = *(*float64)(unsafe.Add(mBase, uint32(v594)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+40)) = int32(0)
	v600 = float64(0)
	v603 = base.F64_add(v596, base.F64_add(base.F64_add(base.F64_add(v472, v471), v600), v590))
	*(*float64)(unsafe.Add(mBase, uint32(v417)+48)) = v603
	v605 = *(*float64)(unsafe.Add(mBase, uint32(v417)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v417)+56)) = base.F64_add(v603, base.F64_add(base.F64_mul(v595, v605), base.F64_add(base.F64_mul(v591, base.F64_add(v580, v593)), v600)))
	m.G0 = v436 + int32(32)
	F_add_path(m, l1, v417)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L19
	} else {
		goto L133
	}
L123:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v424)+16))
	v474 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v436)+16)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v436)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v436)+24)) = v474
	if v473 == int32(0) {
		v546 = v18
		v556 = float64(0)
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v561 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v562 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v580 = v561
	v590 = v562
	goto L122
L126:
	;
	v557 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v559 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v580 = base.F64_add(v546, v557)
	v590 = base.F64_add(v556, v559)
	goto L122
L127:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v483 <= int32(0) {
		v546 = v18
		v556 = float64(0)
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v492 = int32(0)
	goto L129
L129:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v473)+12))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v514+v492<<(uint(int32(2))%32))))
	v521 = F_cost_qual_eval_walker(m, v518, v436+int32(8))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L19
	} else {
		goto L131
	}
L130:
	;
	v527 = *(*float64)(unsafe.Add(mBase, uint32(v436)+24))
	v528 = *(*float64)(unsafe.Add(mBase, uint32(v436)+16))
	v546 = v527
	v556 = v528
	goto L126
L131:
	;
	v524 = v492 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v524 < v525 {
		v492 = v524
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v621))) = int64(1503238553879)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v621)+12)) = v626
	v628 = F_get_baserel_parampathinfo(m, l0, l1, v619)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	v630 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+20)) = uint8(v630)
	*(*int32)(unsafe.Add(mBase, uint32(v621)+16)) = v628
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v621)+64)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(v621)+24)) = v630
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+21)) = uint8(v633)
	v639 = m.G0
	v641 = v639 - int32(32)
	m.G0 = v641
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v643 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	if v628 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v657 = v643 + v644<<(uint(int32(2))%32)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+52))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)+12))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v657 = v650 + v651<<(uint(int32(2))%32) - int32(4)
	goto L136
L140:
	;
	v663 = v628 + int32(8)
	goto L142
L141:
	;
	v663 = l1 + int32(16)
	goto L142
L142:
	;
	v664 = *(*float64)(unsafe.Add(mBase, uint32(v663)))
	*(*float64)(unsafe.Add(mBase, uint32(v621)+32)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v658)+76))
	v667 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v641)+16)) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v641)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v641)+24)) = v667
	v674 = F_cost_qual_eval_walker(m, v666, v641+int32(8))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	v676 = *(*float64)(unsafe.Add(mBase, uint32(v641)+24))
	v677 = *(*float64)(unsafe.Add(mBase, uint32(v641)+16))
	if v628 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v796 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v798 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v621)+12))
	v800 = *(*float64)(unsafe.Add(mBase, uint32(v799)+24))
	v801 = *(*float64)(unsafe.Add(mBase, uint32(v799)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v621)+40)) = int32(0)
	v805 = float64(0)
	v808 = base.F64_add(v801, base.F64_add(base.F64_add(base.F64_add(v677, v676), v805), v795))
	*(*float64)(unsafe.Add(mBase, uint32(v621)+48)) = v808
	v810 = *(*float64)(unsafe.Add(mBase, uint32(v621)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v621)+56)) = base.F64_add(v808, base.F64_add(base.F64_mul(v800, v810), base.F64_add(base.F64_mul(v796, base.F64_add(v785, v798)), v805)))
	m.G0 = v641 + int32(32)
	F_add_path(m, l1, v621)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L19
	} else {
		goto L155
	}
L145:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v628)+16))
	v679 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v641)+16)) = v679
	*(*int32)(unsafe.Add(mBase, uint32(v641)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v641)+24)) = v679
	if v678 == int32(0) {
		v751 = v18
		v761 = float64(0)
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v766 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v767 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v785 = v766
	v795 = v767
	goto L144
L148:
	;
	v762 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v764 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v785 = base.F64_add(v751, v762)
	v795 = base.F64_add(v761, v764)
	goto L144
L149:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	if v688 <= int32(0) {
		v751 = v18
		v761 = float64(0)
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v697 = int32(0)
	goto L151
L151:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v678)+12))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v719+v697<<(uint(int32(2))%32))))
	v726 = F_cost_qual_eval_walker(m, v723, v641+int32(8))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L19
	} else {
		goto L153
	}
L152:
	;
	v732 = *(*float64)(unsafe.Add(mBase, uint32(v641)+24))
	v733 = *(*float64)(unsafe.Add(mBase, uint32(v641)+16))
	v751 = v732
	v761 = v733
	goto L148
L153:
	;
	v729 = v697 + int32(1)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	if v729 < v730 {
		v697 = v729
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v826)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v826))) = int64(1498943586583)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v826)+12)) = v831
	v833 = F_get_baserel_parampathinfo(m, l0, l1, v824)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	v835 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v826)+20)) = uint8(v835)
	*(*int32)(unsafe.Add(mBase, uint32(v826)+16)) = v833
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v826)+64)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v826)+24)) = v835
	*(*uint8)(unsafe.Add(mBase, uint32(v826)+21)) = uint8(v838)
	v844 = m.G0
	v846 = v844 - int32(32)
	m.G0 = v846
	if v833 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v975 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v977 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v826)+12))
	v979 = *(*float64)(unsafe.Add(mBase, uint32(v978)+24))
	v980 = *(*float64)(unsafe.Add(mBase, uint32(v978)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v826)+40)) = int32(0)
	v983 = float64(0)
	v985 = base.F64_add(v980, base.F64_add(v974, v983))
	*(*float64)(unsafe.Add(mBase, uint32(v826)+48)) = v985
	*(*float64)(unsafe.Add(mBase, uint32(v826)+56)) = base.F64_add(v985, base.F64_add(base.F64_mul(v979, v964), base.F64_add(base.F64_mul(v975, base.F64_add(v968, base.F64_add(v965, v977))), v983)))
	m.G0 = v846 + int32(32)
	F_add_path(m, l1, v826)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L19
	} else {
		goto L169
	}
L159:
	;
	v848 = *(*float64)(unsafe.Add(mBase, uint32(v833)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v826)+32)) = v848
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v833)+16))
	v851 = int32(0)
	v853 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[2]))
	v854 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v846)+16)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v846)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v846)+24)) = v854
	if v850 == v851 {
		v926 = v848
		v927 = v18
		v936 = float64(0)
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v941 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v826)+32)) = v941
	v943 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v945 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[2]))
	v946 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v964 = v941
	v965 = v943
	v968 = v945
	v974 = v946
	goto L158
L162:
	;
	v937 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v939 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v964 = v926
	v965 = base.F64_add(v927, v937)
	v968 = v853
	v974 = base.F64_add(v936, v939)
	goto L158
L163:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v850)+4))
	if v863 <= int32(0) {
		v926 = v848
		v927 = v18
		v936 = float64(0)
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v871 = v851
	goto L165
L165:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v850)+12))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v893+v871<<(uint(int32(2))%32))))
	v900 = F_cost_qual_eval_walker(m, v897, v846+int32(8))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L19
	} else {
		goto L167
	}
L166:
	;
	v906 = *(*float64)(unsafe.Add(mBase, uint32(v826)+32))
	v907 = *(*float64)(unsafe.Add(mBase, uint32(v846)+24))
	v908 = *(*float64)(unsafe.Add(mBase, uint32(v846)+16))
	v926 = v906
	v927 = v907
	v936 = v908
	goto L162
L167:
	;
	v903 = v871 + int32(1)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v850)+4))
	if v903 < v904 {
		v871 = v903
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v2639 = l0
	v2640 = l1
	v2641 = l2
	v2642 = l3
	v2652 = v30
	goto L1
L170:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1005
	F_errmsg_internal(m, int32(_a_F_set_rel_pathlist_3), v30)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L19
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_set_rel_pathlist_4), int32(527), int32(_a_F_set_rel_pathlist_5))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L19
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_rel_pathlist[3])))
	if v1023 != 0 {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	m.G0 = v1018 + int32(16)
	if v1472 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L175:
	;
	v1472 = int32(0)
	goto L174
L176:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+152)))
	if v1052&int32(1) == int32(0) {
		goto L186
	} else {
		goto L187
	}
L177:
	;
	v1027 = int32(1)
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1018)+15)))
	if v1026&v1027|v1030&v1027 == int32(0) {
		goto L175
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	if v1026&int32(1) == int32(0) {
		v1472 = v5
		goto L174
	} else {
		goto L185
	}
L180:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1037 = F_create_tidscan_path(m, l0, l1, v1023, v1036)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L19
	} else {
		goto L181
	}
L181:
	;
	F_add_path(m, l1, v1037)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L19
	} else {
		goto L182
	}
L182:
	;
	if v1030&int32(1) != 0 {
		v1472 = v1027
		goto L174
	} else {
		goto L183
	}
L183:
	;
	v1043 = int32(0)
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_rel_pathlist[3])))
	if v1045 != 0 {
		v1051 = v1043
		goto L176
	} else {
		goto L184
	}
L184:
	;
	v1472 = v1043
	goto L174
L185:
	;
	v1051 = v5
	goto L176
L186:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v1425 == int32(1) {
		goto L228
	} else {
		goto L229
	}
L187:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v1057 == int32(0) {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1060 <= int32(0) {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v1069 = v1051
	v1070 = v5
	goto L190
L190:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1090+v1069<<(uint(int32(2))%32))))
	v1095 = F_IsBinaryTidClause(m, v1094, l1)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L19
	} else {
		goto L193
	}
L191:
	;
	if v1107 == int32(0) {
		goto L186
	} else {
		goto L198
	}
L192:
	;
	v1109 = v1069 + int32(1)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1109 < v1110 {
		v1069 = v1109
		v1070 = v1107
		goto L190
	} else {
		goto L197
	}
L193:
	;
	if v1095 == int32(0) {
		v1107 = v1070
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+4))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+4))
	if base.Ui32(int32(3)) < base.Ui32(v1100-int32(2799)) {
		v1107 = v1070
		goto L192
	} else {
		goto L195
	}
L195:
	;
	v1105 = F_lappend(m, v1070, v1094)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L19
	} else {
		goto L196
	}
L196:
	;
	v1107 = v1105
	goto L192
L197:
	;
	goto L191
L198:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1116 = F_palloc0(m, int32(80))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L19
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1116))) = int64(1486058684702)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+12)) = v1121
	v1123 = F_get_baserel_parampathinfo(m, l0, l1, v1114)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L19
	} else {
		goto L200
	}
L200:
	;
	v1125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1116)+20)) = uint8(v1125)
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+16)) = v1123
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+72)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+64)) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+24)) = v1125
	*(*uint8)(unsafe.Add(mBase, uint32(v1116)+21)) = uint8(v1128)
	v1135 = m.G0
	v1137 = v1135 - int32(48)
	m.G0 = v1137
	if v1123 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1143 = v1123 + int32(8)
	goto L203
L202:
	;
	v1143 = l1 + int32(16)
	goto L203
L203:
	;
	v1144 = *(*float64)(unsafe.Add(mBase, uint32(v1143)))
	*(*float64)(unsafe.Add(mBase, uint32(v1116)+32)) = v1144
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1147 = int32(0)
	v1149 = F_clauselist_selectivity(m, l0, v1107, v1146, v1147, v1147)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L19
	} else {
		goto L204
	}
L204:
	;
	v1151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v1153 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1137)+32)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v1137)+24)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v1137)+40)) = v1153
	if v1107 == int32(0) {
		v1222 = v18
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_get_tablespace_page_costs(m, v1231, v1137+int32(16), v1137+int32(8))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L19
	} else {
		goto L212
	}
L206:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	if v1160 <= int32(0) {
		v1222 = v18
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1172 = v5
	goto L208
L208:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+12))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1190+v1172<<(uint(int32(2))%32))))
	v1197 = F_cost_qual_eval_walker(m, v1194, v1137+int32(24))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L19
	} else {
		goto L210
	}
L209:
	;
	v1203 = *(*float64)(unsafe.Add(mBase, uint32(v1137)+40))
	v1222 = v1203
	goto L205
L210:
	;
	v1200 = v1172 + int32(1)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	if v1200 < v1201 {
		v1172 = v1200
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v1238 = *(*float64)(unsafe.Add(mBase, uint32(v1137)+8))
	v1239 = *(*float64)(unsafe.Add(mBase, uint32(v1137)+16))
	if v1123 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+12))
	v1359 = *(*float64)(unsafe.Add(mBase, uint32(v1358)+24))
	v1361 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_pathlist[1]))
	v1362 = *(*float64)(unsafe.Add(mBase, uint32(v1358)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1116)+40)) = int32(0)
	v1366 = float64(0)
	v1368 = base.F64_add(v1362, base.F64_add(base.F64_add(v1222, v1357), v1366))
	*(*float64)(unsafe.Add(mBase, uint32(v1116)+48)) = v1368
	v1370 = *(*float64)(unsafe.Add(mBase, uint32(v1116)+32))
	v1379 = base.F64_ceil(base.F64_mul(v1149, base.F64_convert_i32_u(v1152)))
	if base.F64_le(v1379, v1366) != 0 {
		goto L224
	} else {
		goto L225
	}
L214:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+16))
	v1241 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1137)+32)) = v1241
	*(*int32)(unsafe.Add(mBase, uint32(v1137)+24)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v1137)+40)) = v1241
	if v1240 == int32(0) {
		v1313 = v18
		v1323 = float64(0)
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v1328 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v1329 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v1347 = v1328
	v1357 = v1329
	goto L213
L217:
	;
	v1324 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v1326 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v1347 = base.F64_add(v1313, v1324)
	v1357 = base.F64_add(v1323, v1326)
	goto L213
L218:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+4))
	if v1250 <= int32(0) {
		v1313 = v18
		v1323 = float64(0)
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1263 = int32(0)
	goto L220
L220:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+12))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1281+v1263<<(uint(int32(2))%32))))
	v1288 = F_cost_qual_eval_walker(m, v1285, v1137+int32(24))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L19
	} else {
		goto L222
	}
L221:
	;
	v1294 = *(*float64)(unsafe.Add(mBase, uint32(v1137)+40))
	v1295 = *(*float64)(unsafe.Add(mBase, uint32(v1137)+32))
	v1313 = v1294
	v1323 = v1295
	goto L217
L222:
	;
	v1291 = v1263 + int32(1)
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+4))
	if v1291 < v1292 {
		v1263 = v1291
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v1384 = v1366
	goto L226
L225:
	;
	v1384 = base.F64_add(v1379, float64(-1))
	goto L226
L226:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1116)+56)) = base.F64_add(v1368, base.F64_add(base.F64_mul(v1359, v1370), base.F64_add(base.F64_mul(base.F64_sub(base.F64_add(v1347, v1361), v1222), base.F64_mul(v1149, v1151)), base.F64_add(base.F64_add(base.F64_mul(v1238, v1384), v1239), float64(0)))))
	m.G0 = v1137 + int32(48)
	F_add_path(m, l1, v1116)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L19
	} else {
		goto L227
	}
L227:
	;
	goto L186
L228:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v1431 = F_generate_implied_equalities_for_column(m, l0, l1, int32(827), int32(0), v1430)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L19
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	F_BuildParameterizedTidPaths(m, l0, l1, v1435)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L19
	} else {
		goto L233
	}
L231:
	;
	F_BuildParameterizedTidPaths(m, l0, l1, v1431)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L19
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	goto L175
L234:
	;
	v1499 = F_create_seqscan_path(m, l0, l1, v1015, int32(0))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L19
	} else {
		goto L237
	}
L235:
	;
	v2612 = l0
	v2613 = l1
	v2614 = l2
	v2615 = l3
	v2625 = v30
	goto L236
L236:
	;
	v2639 = v2612
	v2640 = v2613
	v2641 = v2614
	v2642 = v2615
	v2652 = v2625
	goto L1
L237:
	;
	F_add_path(m, l1, v1499)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L19
	} else {
		goto L238
	}
L238:
	;
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if base.B2i32(v1503 != int32(1))|v1015 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1624 = m.G0
	v1626 = v1624 - int32(416)
	m.G0 = v1626
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v1628 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L240:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	if v1509 != int32(-1) {
		v1566 = v1509
		goto L241
	} else {
		goto L242
	}
L241:
	;
	if v1566 < v1508 {
		goto L251
	} else {
		goto L252
	}
L242:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1516 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[5]))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if base.B2i32(v1512 == int32(0))&base.B2i32(base.I64_extend_i32_u(v1518) < base.I64_extend_i32_s(v1516)) != 0 {
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v1522 = int32(1)
	if v1516 <= v1522 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1525 = v1522
	goto L246
L245:
	;
	v1525 = v1516
	goto L246
L246:
	;
	v1531 = v1525
	v1532 = int32(1)
	goto L247
L247:
	;
	v1555 = v1531 * int32(3)
	if base.Ui32(v1518) < base.Ui32(v1555) {
		v1566 = v1532
		goto L241
	} else {
		goto L249
	}
L248:
	;
	v1566 = v1558
	goto L241
L249:
	;
	v1558 = v1532 + int32(1)
	if v1555 < int32(715827883) {
		v1531 = v1555
		v1532 = v1558
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v1589 = v1566
	goto L253
L252:
	;
	v1589 = v1508
	goto L253
L253:
	;
	if v1589 <= int32(0) {
		goto L239
	} else {
		goto L254
	}
L254:
	;
	v1593 = F_create_seqscan_path(m, l0, l1, int32(0), v1589)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L19
	} else {
		goto L255
	}
L255:
	;
	F_add_partial_path(m, l1, v1593)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L19
	} else {
		goto L256
	}
L256:
	;
	goto L239
L257:
	;
	m.G0 = v1626 + int32(416)
	v2612 = l0
	v2613 = l1
	v2614 = l2
	v2615 = l3
	v2625 = v30
	goto L236
L258:
	;
	v1631 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+404)) = v1631
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+400)) = v1631
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+4))
	if v1631 < v1635 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1653 = v5
	v1654 = v5
	goto L262
L260:
	;
	v2133 = v5
	v2149 = int32(0)
	goto L261
L261:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v2152 = F_generate_bitmap_or_paths(m, l0, l1, v2150, int32(0))
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L19
	} else {
		goto L321
	}
L262:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+12))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1669+v1654<<(uint(int32(2))%32))))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+88))
	if v1674 != 0 {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+404))
	v2133 = v2100
	v2149 = v2120
	goto L261
L264:
	;
	v2117 = v1654 + int32(1)
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+4))
	if v2117 < v2118 {
		v1653 = v2100
		v1654 = v2117
		goto L262
	} else {
		goto L320
	}
L265:
	;
	v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+100)))
	if v1675 != int32(1) {
		v2100 = v1653
		goto L264
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1680 = int32(0)
	base.MemoryFill(m, v1626+int32(268), v1680, int32(132))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+96))
	if v1683 == v1680 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	goto L267
L269:
	;
	F_get_index_paths(m, l0, l1, v1673, v1626+int32(268), v1626+int32(404))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L19
	} else {
		goto L276
	}
L270:
	;
	v1686 = int32(0)
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1683)+4))
	if v1687 <= v1686 {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1694 = v1686
	goto L272
L272:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1683)+12))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1717+v1694<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v1721, v1673, v1626+int32(268))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L19
	} else {
		goto L274
	}
L273:
	;
	goto L269
L274:
	;
	v1727 = v1694 + int32(1)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1683)+4))
	if v1727 < v1728 {
		v1694 = v1727
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v1765 = int32(0)
	base.MemoryFill(m, v1626+int32(136), v1765, int32(132))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v1768 == v1765 {
		v1835 = v1653
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1853 = int32(0)
	base.MemoryFill(m, v1626+int32(4), v1853, int32(132))
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+12))
	v1858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1857)+216)))
	if v1858 != int32(1) {
		v2006 = v1853
		goto L293
	} else {
		goto L294
	}
L278:
	;
	v1771 = int32(0)
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+4))
	if v1772 <= v1771 {
		v1835 = v1653
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1779 = v1771
	v1786 = v1653
	goto L280
L280:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+12))
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1802+v1779<<(uint(int32(2))%32))))
	v1807 = F_join_clause_is_movable_to(m, v1806, l1)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L19
	} else {
		goto L282
	}
L281:
	;
	v1835 = v1819
	goto L277
L282:
	;
	if v1807 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1806)+52))
	goto L286
L284:
	;
	v1819 = v1786
	goto L285
L285:
	;
	v1821 = v1779 + int32(1)
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+4))
	if v1821 < v1822 {
		v1779 = v1821
		v1786 = v1819
		goto L280
	} else {
		goto L292
	}
L286:
	;
	if v1809 != int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1812 = F_list_append_unique_ptr(m, v1786, v1806)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L19
	} else {
		goto L290
	}
L288:
	;
	v1814 = v1786
	goto L289
L289:
	;
	F_match_clause_to_index(m, l0, v1806, v1673, v1626+int32(136))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L19
	} else {
		goto L291
	}
L290:
	;
	v1814 = v1812
	goto L289
L291:
	;
	v1819 = v1814
	goto L285
L292:
	;
	goto L281
L293:
	;
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+136)))
	if v2007|v2006&int32(1) == int32(0) {
		v2100 = v1835
		goto L264
	} else {
		goto L307
	}
L294:
	;
	v1861 = int32(0)
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+40))
	if v1862 <= v1861 {
		v2006 = v1861
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1872 = v1861
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+412)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+408)) = v1673
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+12))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+104))
	v1899 = F_generate_implied_equalities_for_column(m, l0, v1894, int32(821), v1626+int32(408), v1898)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L19
	} else {
		goto L299
	}
L297:
	;
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+4)))
	v2006 = v1978
	goto L293
L298:
	;
	v1975 = v1872 + int32(1)
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+40))
	if v1975 < v1976 {
		v1872 = v1975
		goto L296
	} else {
		goto L306
	}
L299:
	;
	if v1899 == int32(0) {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1903 = int32(0)
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1899)+4))
	if v1904 <= v1903 {
		goto L298
	} else {
		goto L301
	}
L301:
	;
	v1911 = v1903
	goto L302
L302:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1899)+12))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1934+v1911<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v1938, v1673, v1626+int32(4))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L19
	} else {
		goto L304
	}
L303:
	;
	goto L298
L304:
	;
	v1944 = v1911 + int32(1)
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1899)+4))
	if v1944 < v1945 {
		v1911 = v1944
		goto L302
	} else {
		goto L305
	}
L305:
	;
	goto L303
L306:
	;
	goto L297
L307:
	;
	v2013 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+408)) = v2013
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+40))
	if v2017 <= v2013 {
		v2100 = v1835
		goto L264
	} else {
		goto L308
	}
L308:
	;
	v2024 = v2013
	v2029 = v2013
	goto L309
L309:
	;
	v2049 = v2024 << (uint(int32(2)) % 32)
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v1626+int32(140)+v2049)))
	if v2051 != 0 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v2100 = v1835
	goto L264
L311:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2051)+4))
	v2053 = v2052
	goto L313
L312:
	;
	v2053 = int32(0)
	goto L313
L313:
	;
	v2062 = v2053 + v2029
	F_consider_index_join_outer_rels(m, l0, l1, v1673, v1626+int32(268), v1626+int32(136), v1626+int32(4), v1626+int32(400), v2051, v2062, v1626+int32(408))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L19
	} else {
		goto L314
	}
L314:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1626+int32(8)+v2049)))
	if v2068 != 0 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+4))
	v2071 = v2069
	goto L317
L316:
	;
	v2071 = int32(0)
	goto L317
L317:
	;
	v2080 = v2062 + v2071
	F_consider_index_join_outer_rels(m, l0, l1, v1673, v1626+int32(268), v1626+int32(136), v1626+int32(4), v1626+int32(400), v2068, v2080, v1626+int32(408))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L19
	} else {
		goto L318
	}
L318:
	;
	v2086 = v2024 + int32(1)
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+40))
	if v2086 < v2087 {
		v2024 = v2086
		v2029 = v2080
		goto L309
	} else {
		goto L319
	}
L319:
	;
	goto L310
L320:
	;
	goto L263
L321:
	;
	v2154 = F_list_concat(m, v2149, v2152)
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L19
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+404)) = v2154
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v2158 = F_generate_bitmap_or_paths(m, l0, l1, v2133, v2157)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L19
	} else {
		goto L323
	}
L323:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+400))
	v2161 = F_list_concat(m, v2160, v2158)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L19
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+400)) = v2161
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+404))
	if v2164 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	if v2161 == int32(0) {
		goto L257
	} else {
		goto L354
	}
L326:
	;
	v2167 = F_choose_bitmap_and(m, l0, l1, v2164)
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L19
	} else {
		goto L327
	}
L327:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2172 = F_create_bitmap_heap_path(m, l0, l1, v2167, v2169, float64(1), int32(0))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L19
	} else {
		goto L328
	}
L328:
	;
	F_add_path(m, l1, v2172)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L19
	} else {
		goto L329
	}
L329:
	;
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v2176 != int32(1) {
		goto L325
	} else {
		goto L330
	}
L330:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v2179 != 0 {
		goto L325
	} else {
		goto L331
	}
L331:
	;
	v2181 = int32(0)
	v2183 = F_compute_bitmap_pages(m, l0, l1, v2167, float64(1), v2181, v2181)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L19
	} else {
		goto L332
	}
L332:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[4]))
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	if v2187 != int32(-1) {
		v2257 = v2187
		goto L334
	} else {
		goto L335
	}
L333:
	;
	goto L325
L334:
	;
	if v2257 < v2186 {
		goto L348
	} else {
		goto L349
	}
L335:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2193 = int32(0)
	if v2190|base.B2i32(base.F64_ge(v2183, float64(0)) == v2193) == v2193 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[5]))
	if base.F64_lt(v2183, base.F64_convert_i32_s(v2199)) != 0 {
		goto L333
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	v2202 = int32(0)
	if base.F64_ge(v2183, float64(0)) == v2202 {
		v2257 = v2202
		goto L334
	} else {
		goto L340
	}
L339:
	;
	goto L338
L340:
	;
	v2207 = int32(1)
	v2209 = *(*int32)(unsafe.Add(mBase, _c_F_set_rel_pathlist[5]))
	if v2209 <= v2207 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v2212 = v2207
	goto L343
L342:
	;
	v2212 = v2209
	goto L343
L343:
	;
	v2220 = int32(1)
	v2221 = v2212
	goto L344
L344:
	;
	v2242 = v2221 * int32(3)
	if base.F64_ge(v2183, base.F64_convert_i32_u(v2242)) == int32(0) {
		v2257 = v2220
		goto L334
	} else {
		goto L346
	}
L345:
	;
	v2257 = v2248
	goto L334
L346:
	;
	v2248 = v2220 + int32(1)
	if v2242 < int32(715827883) {
		v2220 = v2248
		v2221 = v2242
		goto L344
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	v2279 = v2257
	goto L350
L349:
	;
	v2279 = v2186
	goto L350
L350:
	;
	if v2279 <= int32(0) {
		goto L333
	} else {
		goto L351
	}
L351:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2284 = F_create_bitmap_heap_path(m, l0, l1, v2167, v2282, float64(1), v2279)
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L19
	} else {
		goto L352
	}
L352:
	;
	F_add_partial_path(m, l1, v2284)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L19
	} else {
		goto L353
	}
L353:
	;
	goto L333
L354:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2161)+4))
	if v2344 <= int32(0) {
		goto L257
	} else {
		goto L355
	}
L355:
	;
	v2347 = int32(0)
	v2353 = v2347
	v2355 = v2347
	goto L356
L356:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2161)+12))
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2376+v2353<<(uint(int32(2))%32))))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2380)+16))
	if v2381 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	if v2385 == int32(0) {
		goto L257
	} else {
		goto L363
	}
L358:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2381)+4))
	v2384 = v2382
	goto L360
L359:
	;
	v2384 = int32(0)
	goto L360
L360:
	;
	v2385 = F_list_append_unique(m, v2355, v2384)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L19
	} else {
		goto L361
	}
L361:
	;
	v2388 = v2353 + int32(1)
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2161)+4))
	if v2388 < v2389 {
		v2353 = v2388
		v2355 = v2385
		goto L356
	} else {
		goto L362
	}
L362:
	;
	goto L357
L363:
	;
	v2393 = int32(0)
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	if v2394 <= v2393 {
		goto L257
	} else {
		goto L364
	}
L364:
	;
	v2408 = v2393
	goto L365
L365:
	;
	v2424 = int32(0)
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+400))
	if v2425 == v2424 {
		v2541 = v2424
		goto L367
	} else {
		goto L368
	}
L366:
	;
	goto L257
L367:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+404))
	v2562 = F_list_concat(m, v2541, v2561)
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L19
	} else {
		goto L394
	}
L368:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v2425)+4))
	if v2428 <= int32(0) {
		v2541 = v2424
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+12))
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2431+v2408<<(uint(int32(2))%32))))
	v2441 = int32(0)
	v2444 = v2424
	goto L370
L370:
	;
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2425)+12))
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2464+v2441<<(uint(int32(2))%32))))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2468)+16))
	if v2469 != 0 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	v2541 = v2529
	goto L367
L372:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2469)+4))
	v2472 = v2470
	goto L374
L373:
	;
	v2472 = int32(0)
	goto L374
L374:
	;
	v2473 = int32(0)
	if v2472 == v2473 {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	if v2526 != 0 {
		goto L389
	} else {
		goto L390
	}
L376:
	;
	v2526 = int32(1)
	goto L375
L377:
	;
	goto L378
L378:
	;
	if v2435 == int32(0) {
		v2519 = v2473
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v2526 = v2519
	goto L375
L380:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2472)+4))
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v2435)+4))
	if v2483 < v2482 {
		v2519 = v2473
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2485 = int32(1)
	if v2482 <= v2485 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v2488 = v2485
	goto L384
L383:
	;
	v2488 = v2482
	goto L384
L384:
	;
	v2489 = int32(8)
	v2494 = int32(0)
	goto L385
L385:
	;
	v2501 = v2494 << (uint(int32(2)) % 32)
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v2472+v2489+v2501)))
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(v2435+v2489+v2501)))
	v2508 = v2503 & (v2505 ^ int32(-1))
	v2510 = base.B2i32(v2508 == int32(0))
	if v2508 != 0 {
		v2519 = v2510
		goto L379
	} else {
		goto L387
	}
L386:
	;
	v2519 = v2510
	goto L379
L387:
	;
	v2512 = v2494 + int32(1)
	if v2512 != v2488 {
		v2494 = v2512
		goto L385
	} else {
		goto L388
	}
L388:
	;
	goto L386
L389:
	;
	v2527 = F_lappend(m, v2444, v2468)
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L19
	} else {
		goto L392
	}
L390:
	;
	v2529 = v2444
	goto L391
L391:
	;
	v2531 = v2441 + int32(1)
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v2425)+4))
	if v2531 < v2532 {
		v2441 = v2531
		v2444 = v2529
		goto L370
	} else {
		goto L393
	}
L392:
	;
	v2529 = v2527
	goto L391
L393:
	;
	goto L371
L394:
	;
	v2564 = F_choose_bitmap_and(m, l0, l1, v2562)
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L19
	} else {
		goto L395
	}
L395:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v2564)+16))
	if v2566 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+4))
	v2569 = v2567
	goto L398
L397:
	;
	v2569 = int32(0)
	goto L398
L398:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2571 = F_get_loop_count(m, l0, v2570, v2569)
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L19
	} else {
		goto L399
	}
L399:
	;
	v2574 = F_create_bitmap_heap_path(m, l0, l1, v2564, v2569, v2571, int32(0))
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L19
	} else {
		goto L400
	}
L400:
	;
	F_add_path(m, l1, v2574)
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L19
	} else {
		goto L401
	}
L401:
	;
	v2579 = v2408 + int32(1)
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	if v2579 < v2580 {
		v2408 = v2579
		goto L365
	} else {
		goto L402
	}
L402:
	;
	goto L366
L403:
	;
	m.T0[v2667].(func(*base.Module, int32, int32, int32, int32))(m, v2639, v2640, v2641, v2642)
	mBase = m.M
	v2669 = m.ExcPending
	if v2669 != 0 {
		goto L19
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+4))
	if v2670 != 0 {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	goto L405
L407:
	;
	F_set_cheapest(m, v2640)
	mBase = m.M
	v2728 = m.ExcPending
	if v2728 != 0 {
		goto L19
	} else {
		goto L422
	}
L408:
	;
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+8))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2639)+52))
	v2673 = int32(0)
	if base.B2i32(v2671 == v2673)|base.B2i32(v2672 == v2673) != 0 {
		v2719 = base.B2i32(v2671|v2672 == v2673)
		goto L410
	} else {
		goto L411
	}
L409:
	;
	if v2719 != 0 {
		goto L407
	} else {
		goto L420
	}
L410:
	;
	goto L409
L411:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v2671)+4))
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2672)+4))
	if v2687 != v2688 {
		v2719 = int32(0)
		goto L410
	} else {
		goto L412
	}
L412:
	;
	v2690 = int32(1)
	if v2687 <= v2690 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v2693 = v2690
	goto L415
L414:
	;
	v2693 = v2687
	goto L415
L415:
	;
	v2694 = int32(8)
	v2699 = int32(0)
	goto L416
L416:
	;
	v2707 = v2699 << (uint(int32(2)) % 32)
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2671+v2694+v2707)))
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2672+v2694+v2707)))
	v2712 = base.B2i32(v2709 == v2711)
	if v2709 != v2711 {
		v2719 = v2712
		goto L410
	} else {
		goto L418
	}
L417:
	;
	v2719 = v2712
	goto L410
L418:
	;
	v2715 = v2699 + int32(1)
	if v2715 != v2693 {
		v2699 = v2715
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	F_generate_useful_gather_paths(m, v2639, v2640, int32(0))
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L19
	} else {
		goto L421
	}
L421:
	;
	goto L407
L422:
	;
	m.G0 = v2652 + int32(16)
	return
}
func F_set_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v24 float64
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
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
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
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
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v662 int32
	_ = v662
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 float64
	_ = v771
	var v772 float64
	_ = v772
	var v773 float64
	_ = v773
	var v774 float64
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 float64
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 float64
	_ = v875
	var v877 float64
	_ = v877
	var v907 int32
	_ = v907
	var v908 float64
	_ = v908
	var v909 float64
	_ = v909
	var v910 float64
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 float64
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v986 float64
	_ = v986
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1002 int32
	_ = v1002
	var v1026 int32
	_ = v1026
	var v1033 float64
	_ = v1033
	var v1067 int32
	_ = v1067
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 float64
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1226 int32
	_ = v1226
	var v1238 int32
	_ = v1238
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1281 int32
	_ = v1281
	var v1289 int32
	_ = v1289
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1475 int32
	_ = v1475
	var v1492 int32
	_ = v1492
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 float64
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 float64
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1630 float64
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 float64
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1782 int32
	_ = v1782
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1840 int32
	_ = v1840
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1911 float64
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 float64
	_ = v1958
	var v1961 float64
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1985 float64
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1990 float64
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1995 float64
	_ = v1995
	var v1996 int64
	_ = v1996
	var v2003 int32
	_ = v2003
	var v2009 int32
	_ = v2009
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2047 float64
	_ = v2047
	var v2048 float64
	_ = v2048
	var v2050 float64
	_ = v2050
	var v2074 float64
	_ = v2074
	var v2075 float64
	_ = v2075
	var v2079 float64
	_ = v2079
	var v2080 float64
	_ = v2080
	var v2082 float64
	_ = v2082
	var v2084 float64
	_ = v2084
	var v2086 float64
	_ = v2086
	var v2088 float64
	_ = v2088
	var v2089 float64
	_ = v2089
	var v2114 float64
	_ = v2114
	var v2115 float64
	_ = v2115
	var v2116 float64
	_ = v2116
	var v2118 float64
	_ = v2118
	var v2119 float64
	_ = v2119
	var v2120 float64
	_ = v2120
	var v2121 float64
	_ = v2121
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2186 int64
	_ = v2186
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2202 int32
	_ = v2202
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2267 int64
	_ = v2267
	var v2269 int64
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2281 float64
	_ = v2281
	var v2282 float64
	_ = v2282
	var v2291 float64
	_ = v2291
	var v2295 float64
	_ = v2295
	var v2297 float64
	_ = v2297
	var v2299 float64
	_ = v2299
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2311 int32
	_ = v2311
	var v2316 int32
	_ = v2316
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2327 int32
	_ = v2327
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2343 int32
	_ = v2343
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2359 int32
	_ = v2359
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2382 int32
	_ = v2382
	var v2394 int32
	_ = v2394
	var v2400 int32
	_ = v2400
	var v2424 int32
	_ = v2424
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2451 int32
	_ = v2451
	var v2475 int32
	_ = v2475
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2485 int32
	_ = v2485
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2533 int32
	_ = v2533
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2672 int32
	_ = v2672
	var v2675 float64
	_ = v2675
	var v2676 float64
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2688 int32
	_ = v2688
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2730 int32
	_ = v2730
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2765 int32
	_ = v2765
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2802 int32
	_ = v2802
	var v2812 int32
	_ = v2812
	var v2832 int32
	_ = v2832
	var v2835 int32
	_ = v2835
	var v2842 int32
	_ = v2842
	var v2867 int32
	_ = v2867
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2914 int32
	_ = v2914
	var v2918 int32
	_ = v2918
	var v2921 int32
	_ = v2921
	var v2928 int32
	_ = v2928
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2968 int32
	_ = v2968
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	v5 = int32(0)
	v24 = float64(0)
	v29 = m.G0
	v31 = v29 - int32(144)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v33 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v31 + int32(144)
	return
L2:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v61 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v34 = F_relation_excluded_by_constraints(m, l0, l1, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v34 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v41
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v55 = F_create_append_path(m, v41, l1, v41, v41, v41, v51, v41, v41, float64(-1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_add_path(m, l1, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v1121 {
	case 0:
		goto L218
	case 1:
		goto L217
	default:
		goto L210
	case 3:
		goto L216
	case 4:
		goto L215
	case 5:
		goto L214
	case 6:
		goto L213
	case 7:
		goto L212
	case 8:
		goto L211
	}
L13:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_rel_size[0])))
	if v67 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+82)))
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v84 = v82 - v83
	v89 = F_palloc0(m, v84<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L19
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v70 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v71 != int32(112) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74-v75<<(uint(int32(2))%32))))
	if v79 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+217)) = uint8(v80)
	goto L14
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v91 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v1099 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+32)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v1099
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1113 = F_create_append_path(m, v1099, l1, v1099, v1099, v1099, v1109, v1099, v1099, float64(-1))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L4
	} else {
		goto L198
	}
L21:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v94 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v105 = v5
	v119 = v5
	v120 = v24
	v121 = v24
	v122 = v24
	goto L23
L23:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v105<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v131 != l2 {
		v907 = v119
		v908 = v120
		v909 = v121
		v910 = v122
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v907 == int32(0) {
		goto L20
	} else {
		goto L187
	}
L25:
	;
	v914 = v105 + int32(1)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v914 < v915 {
		v105 = v914
		v119 = v907
		v120 = v908
		v121 = v909
		v122 = v910
		goto L23
	} else {
		goto L186
	}
L26:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(2))%32))))
	v139 = F_find_base_rel(m, l0, v134)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v141 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
	if v143 == v141 {
		v164 = v141
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v164 != 0 {
		v907 = v119
		v908 = v120
		v909 = v121
		v910 = v122
		goto L25
	} else {
		goto L38
	}
L29:
	;
	goto L28
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v147 = v146
	goto L31
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if base.Ui32(int32(2)) <= base.Ui32(v151-int32(301)) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v164 = int32(1)
	goto L29
L33:
	;
	if v151 != int32(290) {
		v164 = v141
		goto L29
	} else {
		goto L36
	}
L34:
	;
	v147 = v150 + int32(72)
	goto L31
L35:
	;
	goto L32
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)+72))
	if v158 != 0 {
		v164 = v141
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v165 = F_relation_excluded_by_constraints(m, l0, v139, v138)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v165 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = int64(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+32)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v139)+40)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v139)+32)) = v170
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v139)+64))
	v184 = F_create_append_path(m, v170, v139, v170, v170, v170, v180, v170, v170, float64(-1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v190 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	F_add_path(m, v139, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_set_cheapest(m, v139)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v907 = v119
	v908 = v120
	v909 = v121
	v910 = v122
	goto L25
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+212)) = v299
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v328 = F_adjust_appendrel_attrs(m, l0, v324, int32(1), v31+int32(128))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L72
	}
L47:
	;
	v299 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v194 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v196 <= v194 {
		v299 = v194
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v203 = v194
	v204 = v194
	goto L51
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v203<<(uint(int32(2))%32))))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v234 = int32(0)
	if base.B2i32(v232 == v234)|base.B2i32(v233 == v234) != 0 {
		v279 = v234
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v299 = v289
	goto L46
L53:
	;
	if v279 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	goto L53
L55:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v244 < v245 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v247 = v244
	goto L58
L57:
	;
	v247 = v245
	goto L58
L58:
	;
	if v247 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v250 = int32(1)
	goto L61
L60:
	;
	v250 = v247
	goto L61
L61:
	;
	v251 = int32(8)
	v256 = int32(0)
	goto L62
L62:
	;
	v263 = v256 << (uint(int32(2)) % 32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v233+v251+v263)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v232+v251+v263)))
	v268 = v265 & v267
	v270 = base.B2i32(v268 != int32(0))
	if v268 != 0 {
		v279 = v270
		goto L54
	} else {
		goto L64
	}
L63:
	;
	v279 = v270
	goto L54
L64:
	;
	v272 = v256 + int32(1)
	if v272 != v250 {
		v256 = v272
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v285 = F_adjust_appendrel_attrs(m, l0, v231, int32(1), v31+int32(128))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	v289 = v204
	goto L68
L68:
	;
	v291 = v203 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v291 < v292 {
		v203 = v291
		v204 = v289
		goto L51
	} else {
		goto L71
	}
L69:
	;
	v287 = F_lappend(m, v204, v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v289 = v287
	goto L68
L71:
	;
	goto L52
L72:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v328
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v332 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+216)) = uint8(v724)
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+217)))
	if v726 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L74:
	;
	v336 = int32(1)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v337 != 0 {
		v343 = v336
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	v347 = m.G0
	v349 = v347 - int32(16)
	m.G0 = v349
	*(*int32)(unsafe.Add(mBase, uint32(v349)+12)) = v346
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v139)+228))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v354 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L77:
	;
	if v343 == int32(0) {
		goto L73
	} else {
		goto L82
	}
L78:
	;
	goto L77
L79:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v338 != 0 {
		v343 = v336
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v339 != 0 {
		v343 = v336
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v343 = base.B2i32(v340 != int32(0))
	goto L78
L82:
	;
	goto L76
L83:
	;
	if int32(0) <= v411 {
		goto L94
	} else {
		goto L95
	}
L84:
	;
	v411 = base.I32_ctz(v397) | v398<<(uint(int32(5))%32)
	goto L83
L85:
	;
	v411 = int32(-2)
	goto L83
L86:
	;
	v364 = base.I32_div_s(int32(0), int32(32))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if v365 <= v364 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v368 = v354 + int32(8)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v368+v364<<(uint(int32(2))%32))))
	v375 = v372 & int32(-1)
	if v375 != 0 {
		v397 = v375
		v398 = v364
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v377 = v364 + int32(1)
	if v377 == v365 {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v380 = v377
	goto L90
L90:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v368+v380<<(uint(int32(2))%32))))
	if v387 != 0 {
		v397 = v387
		v398 = v380
		goto L84
	} else {
		goto L92
	}
L91:
	;
	goto L85
L92:
	;
	v389 = v380 + int32(1)
	if v389 != v365 {
		v380 = v389
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v419 = v411
	goto L97
L95:
	;
	goto L96
L96:
	;
	m.G0 = v349 + int32(16)
	goto L73
L97:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443+v419<<(uint(int32(2))%32))))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+41)))
	if v448 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L96
L99:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v606 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L100:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v447)+16))
	if v449 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v452 = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v453 <= v452 {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v460 = v452
	goto L103
L103:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484+v460<<(uint(int32(2))%32))))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+12)))
	if v489 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L99
L105:
	;
	v575 = v460 + int32(1)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v575 < v576 {
		v460 = v575
		goto L103
	} else {
		goto L132
	}
L106:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v488)+8))
	v491 = int32(0)
	if v490 == v491 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v544 == int32(0) {
		goto L105
	} else {
		goto L121
	}
L108:
	;
	v544 = int32(1)
	goto L107
L109:
	;
	goto L110
L110:
	;
	if v353 == int32(0) {
		v537 = v491
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v544 = v537
	goto L107
L112:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v501 < v500 {
		v537 = v491
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v503 = int32(1)
	if v500 <= v503 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v506 = v503
	goto L116
L115:
	;
	v506 = v500
	goto L116
L116:
	;
	v507 = int32(8)
	v512 = int32(0)
	goto L117
L117:
	;
	v519 = v512 << (uint(int32(2)) % 32)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v490+v507+v519)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v353+v507+v519)))
	v526 = v521 & (v523 ^ int32(-1))
	v528 = base.B2i32(v526 == int32(0))
	if v526 != 0 {
		v537 = v528
		goto L111
	} else {
		goto L119
	}
L118:
	;
	v537 = v528
	goto L111
L119:
	;
	v530 = v512 + int32(1)
	if v530 != v506 {
		v512 = v530
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v488)+8))
	if v547 == int32(0) {
		goto L105
	} else {
		goto L122
	}
L122:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v551 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v488)+8))
	v564 = F_bms_difference(m, v563, v353)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L129
	}
L124:
	;
	v557 = F_adjust_appendrel_attrs(m, l0, v550, int32(1), v349+int32(12))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v139)+224))
	v560 = F_adjust_appendrel_attrs_multilevel(m, l0, v550, v139, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L128
	}
L127:
	;
	v562 = v557
	goto L123
L128:
	;
	v562 = v560
	goto L123
L129:
	;
	v566 = F_bms_add_members(m, v564, v352)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v488)+20))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v488)+16))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v139)+68))
	F_add_child_eq_member(m, l0, v447, v419, v562, v566, v568, v488, v569, v570)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	goto L105
L132:
	;
	goto L104
L133:
	;
	if int32(0) <= v662 {
		v419 = v662
		goto L97
	} else {
		goto L144
	}
L134:
	;
	v662 = base.I32_ctz(v648) | v649<<(uint(int32(5))%32)
	goto L133
L135:
	;
	v662 = int32(-2)
	goto L133
L136:
	;
	v613 = v419 + int32(1)
	v615 = base.I32_div_s(v613, int32(32))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v606)+4))
	if v616 <= v615 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v619 = v606 + int32(8)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619+v615<<(uint(int32(2))%32))))
	v626 = v623 & (int32(-1) << (uint(v613) % 32))
	if v626 != 0 {
		v648 = v626
		v649 = v615
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v628 = v615 + int32(1)
	if v628 == v616 {
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v631 = v628
	goto L140
L140:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v619+v631<<(uint(int32(2))%32))))
	if v638 != 0 {
		v648 = v638
		v649 = v631
		goto L134
	} else {
		goto L142
	}
L141:
	;
	goto L135
L142:
	;
	v640 = v631 + int32(1)
	if v640 != v616 {
		v631 = v640
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	goto L98
L145:
	;
	v729 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+217)) = uint8(v729)
	goto L147
L146:
	;
	goto L147
L147:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+82)))
	if v732 != int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_set_rel_size(m, l0, v139, v134, v138)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L152
	}
L149:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v735 != int32(1) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	F_set_rel_consider_parallel(m, l0, v139, v138)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v742 = int32(0)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
	if v744 == v742 {
		v765 = v742
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v765 != 0 {
		v907 = v119
		v908 = v120
		v909 = v121
		v910 = v122
		goto L25
	} else {
		goto L163
	}
L154:
	;
	goto L153
L155:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v744)+12))
	v748 = v747
	goto L156
L156:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v748)))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	if base.Ui32(int32(2)) <= base.Ui32(v752-int32(301)) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v765 = int32(1)
	goto L154
L158:
	;
	if v752 != int32(290) {
		v765 = v742
		goto L154
	} else {
		goto L161
	}
L159:
	;
	v748 = v751 + int32(72)
	goto L156
L160:
	;
	goto L157
L161:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v751)+72))
	if v759 != 0 {
		v765 = v742
		goto L154
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+26)))
	if v766 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v769 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v769)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v771 = *(*float64)(unsafe.Add(mBase, uint32(v139)+16))
	v772 = base.F64_add(v120, v771)
	v773 = *(*float64)(unsafe.Add(mBase, uint32(v139)+120))
	v774 = base.F64_add(v121, v773)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)+32))
	v779 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v776), v771), v122)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)+4))
	v788 = int32(0)
	goto L167
L167:
	;
	v812 = int32(0)
	if v782 == v812 {
		v822 = v812
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v823 = int32(1)
	if v780 == int32(0) {
		v907 = v823
		v908 = v772
		v909 = v774
		v910 = v779
		goto L25
	} else {
		goto L172
	}
L170:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	if v816 <= v788 {
		v822 = int32(0)
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v782)+12))
	v822 = v818 + v788<<(uint(int32(2))%32)
	goto L169
L172:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	if base.B2i32(v822 == int32(0))|base.B2i32(v828 <= v788) != 0 {
		v907 = v823
		v908 = v772
		v909 = v774
		v910 = v779
		goto L25
	} else {
		goto L173
	}
L173:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v780)+12))
	if v831 == int32(0) {
		v907 = v823
		v908 = v772
		v909 = v774
		v910 = v779
		goto L25
	} else {
		goto L174
	}
L174:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v822)))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	if v835 != int32(6) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v788 = v788 + int32(1)
	goto L167
L176:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	if v838 != l2 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v840 = int32(*(*int16)(unsafe.Add(mBase, uint32(v834)+8)))
	v841 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v831+v788<<(uint(int32(2))%32))))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)))
	if v847 != int32(6) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v873 = v89 + (v840-v841)<<(uint(int32(3))%32)
	v875 = *(*float64)(unsafe.Add(mBase, uint32(v139)+16))
	v877 = *(*float64)(unsafe.Add(mBase, uint32(v873)))
	*(*float64)(unsafe.Add(mBase, uint32(v873))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v870), v875), v877)
	goto L175
L179:
	;
	v864 = F_exprType(m, v846)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L4
	} else {
		goto L183
	}
L180:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v139)+68))
	if v850 != v851 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v139)+88))
	v854 = int32(*(*int16)(unsafe.Add(mBase, uint32(v846)+8)))
	v855 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139)+80)))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v853+(v854-v855)<<(uint(int32(2))%32))))
	if int32(0) < v860 {
		v870 = v860
		goto L178
	} else {
		goto L182
	}
L182:
	;
	goto L179
L183:
	;
	v866 = F_exprTypmod(m, v846)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v868 = F_get_typavgwidth(m, v864, v866)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v870 = v868
	goto L178
L186:
	;
	goto L24
L187:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v908
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v909
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v921)+32)) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v910, v908)))
	v926 = int32(0)
	if v84 < v926 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	F_pfree(m, v89)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L4
	} else {
		goto L197
	}
L189:
	;
	if v82 != v83 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v930 = int32(1)
	v931 = v84 + v930
	v940 = int32(0)
	v941 = v926
	goto L193
L191:
	;
	v1002 = v926
	goto L192
L192:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v1033 = *(*float64)(unsafe.Add(mBase, uint32(v89+v1002<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1026+v1002<<(uint(int32(2))%32)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v1033, v908)))
	goto L188
L193:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v966 = int32(2)
	v969 = int32(3)
	v972 = *(*float64)(unsafe.Add(mBase, uint32(v89+v941<<(uint(v969)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v965+v941<<(uint(v966)%32)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v972, v908)))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v979 = v941 | int32(1)
	v986 = *(*float64)(unsafe.Add(mBase, uint32(v89+v979<<(uint(v969)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v977+v979<<(uint(v966)%32)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v986, v908)))
	v992 = v941 + v966
	v994 = v940 + v966
	if v994 != v931&int32(-2) {
		v940 = v994
		v941 = v992
		goto L193
	} else {
		goto L195
	}
L194:
	;
	if v931&v930 == int32(0) {
		goto L188
	} else {
		goto L196
	}
L195:
	;
	goto L194
L196:
	;
	v1002 = v992
	goto L192
L197:
	;
	goto L1
L198:
	;
	F_add_path(m, l1, v1113)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	F_pfree(m, v89)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	goto L1
L202:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	F_pfree(m, v2424)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L4
	} else {
		goto L436
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+184)) = v2382
	v2400 = v2394
	goto L202
L204:
	;
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v2382 = v1492
	v2394 = v2365
	goto L203
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L4
	} else {
		goto L433
	}
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L4
	} else {
		goto L430
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L4
	} else {
		goto L427
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L4
	} else {
		goto L424
	}
L209:
	;
	v2179 = m.G0
	v2181 = v2179 - int32(32)
	m.G0 = v2181
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(4652007308841189376)
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v2186 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2181)+16)) = v2186
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v2181)+24)) = v2186
	v2192 = v2181 + int32(16)
	if v2185 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L4
	} else {
		goto L405
	}
L211:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(4607182418800017408)
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L4
	} else {
		goto L400
	}
L212:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1942 != 0 {
		goto L379
	} else {
		goto L380
	}
L213:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+92)))
	if v1635 != int32(1) {
		goto L326
	} else {
		goto L327
	}
L214:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1610 != 0 {
		goto L318
	} else {
		goto L319
	}
L215:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(4636737291354636288)
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L4
	} else {
		goto L316
	}
L216:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1508 != 0 {
		goto L302
	} else {
		goto L303
	}
L217:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = int32(0)
	v1182 = F_copyObjectImpl(m, v1179)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L4
	} else {
		goto L233
	}
L218:
	;
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	switch v1122 - int32(102) {
	case 0:
		goto L209
	default:
		goto L219
	case 10:
		goto L220
	}
L219:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v1148 != 0 {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v1128 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+32)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v1128
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1142 = F_create_append_path(m, v1128, l1, v1128, v1128, v1128, v1138, v1128, v1128, float64(-1))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L4
	} else {
		goto L221
	}
L221:
	;
	F_add_path(m, l1, v1142)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L222
	}
L222:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	goto L1
L224:
	;
	v1149 = m.G0
	v1151 = v1149 - int32(16)
	m.G0 = v1151
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	F_check_index_predicates(m, l0, l1)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L4
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	F_check_index_predicates(m, l0, l1)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L4
	} else {
		goto L231
	}
L227:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	v1157 = F_GetTsmRoutine(m, v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+8))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+12))
	m.T0[v1162].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, v1159, v1151+int32(12), v1151)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+116)) = v1165
	v1167 = *(*float64)(unsafe.Add(mBase, uint32(v1151)))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v1167
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L4
	} else {
		goto L230
	}
L230:
	;
	m.G0 = v1151 + int32(16)
	goto L1
L231:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	goto L1
L233:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = int64(0)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+76))
	if v1187 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+4))
	v1192 = v1188 + int32(1)
	goto L236
L235:
	;
	v1192 = int32(1)
	goto L236
L236:
	;
	v1193 = F_palloc0(m, v1192)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1193
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+133)) = uint8(v1196)
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v1198 == int32(0) {
		v2400 = v5
		goto L202
	} else {
		goto L238
	}
L238:
	;
	v1203 = F_subquery_is_pushdown_safe(m, v1182, v1182, v31+int32(128))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	if v1203 == int32(0) {
		v2400 = v5
		goto L202
	} else {
		goto L240
	}
L240:
	;
	v1207 = int32(0)
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v1208 == v1207 {
		v2382 = v5
		v2394 = v1207
		goto L203
	} else {
		goto L241
	}
L241:
	;
	v1211 = int32(0)
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1212 <= v1211 {
		v2382 = v5
		v2394 = v1211
		goto L203
	} else {
		goto L242
	}
L242:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+132)))
	v1217 = int32(1)
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+133)))
	v1226 = v5
	v1238 = v5
	goto L243
L243:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+12))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1250+v1226<<(uint(int32(2))%32))))
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254)+10)))
	if v1255 != 0 {
		goto L248
	} else {
		goto L249
	}
L244:
	;
	goto L204
L245:
	;
	v1505 = v1226 + int32(1)
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1505 < v1506 {
		v1226 = v1505
		v1238 = v1492
		goto L243
	} else {
		goto L300
	}
L246:
	;
	F_subquery_push_qual(m, v1182, l3, l2, v1256)
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L4
	} else {
		goto L299
	}
L247:
	;
	F_list_free(m, v1264)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L4
	} else {
		goto L298
	}
L248:
	;
	v1441 = F_lappend(m, v1238, v1254)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L4
	} else {
		goto L297
	}
L249:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+4))
	v1257 = F_contain_subplans(m, v1256)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	if v1257 != 0 {
		goto L248
	} else {
		goto L251
	}
L251:
	;
	if v1216&v1217 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1259 = F_contain_volatile_functions(m, v1254)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L4
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	if v1219&v1217 != 0 {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	if v1259 != 0 {
		goto L248
	} else {
		goto L256
	}
L256:
	;
	goto L254
L257:
	;
	v1261 = F_contain_leaked_vars(m, v1256)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L4
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1264 = F_pull_var_clause(m, v1256, int32(16))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L4
	} else {
		goto L262
	}
L260:
	;
	if v1261 != 0 {
		goto L248
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	if v1264 == int32(0) {
		goto L247
	} else {
		goto L263
	}
L263:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	if v1268 <= int32(0) {
		goto L247
	} else {
		goto L264
	}
L264:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	v1281 = int32(0)
	v1289 = int32(1)
	goto L266
L265:
	;
	F_list_free(m, v1264)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L4
	} else {
		goto L296
	}
L266:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1271+v1281<<(uint(int32(2))%32))))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	if v1306 != int32(6) {
		goto L265
	} else {
		goto L268
	}
L267:
	;
	F_list_free(m, v1264)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L4
	} else {
		goto L276
	}
L268:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	if v1309 != l2 {
		goto L265
	} else {
		goto L269
	}
L269:
	;
	v1311 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1305)+8)))
	if v1311 == int32(0) {
		goto L265
	} else {
		goto L270
	}
L270:
	;
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311+v1215))))
	if v1315 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	if v1315&int32(23) != 0 {
		goto L265
	} else {
		goto L274
	}
L272:
	;
	v1319 = v1289
	goto L273
L273:
	;
	v1321 = v1281 + int32(1)
	if v1268 != v1321 {
		v1281 = v1321
		v1289 = v1319
		goto L266
	} else {
		goto L275
	}
L274:
	;
	v1319 = int32(2)
	goto L273
L275:
	;
	goto L267
L276:
	;
	if v1319 == int32(1) {
		goto L246
	} else {
		goto L277
	}
L277:
	;
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182)+37)))
	if v1327 != int32(1) {
		goto L248
	} else {
		goto L278
	}
L278:
	;
	v1330 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+140)) = uint8(v1330)
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	if v1332 != int32(17) {
		goto L248
	} else {
		goto L279
	}
L279:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+28))
	if v1335 == int32(0) {
		goto L248
	} else {
		goto L280
	}
L280:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+4))
	if v1338 != int32(2) {
		goto L248
	} else {
		goto L281
	}
L281:
	;
	F_set_opfuncid(m, v1256)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L4
	} else {
		goto L282
	}
L282:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+8))
	v1344 = F_func_strict(m, v1343)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	if v1344 == int32(0) {
		goto L248
	} else {
		goto L284
	}
L284:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+28))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+12))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	if v1351 != int32(6) {
		v1376 = v1349
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+140)))
	if v1406&int32(1) != 0 {
		goto L248
	} else {
		goto L295
	}
L286:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1378)))
	if v1379 != int32(6) {
		goto L248
	} else {
		goto L291
	}
L287:
	;
	v1354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1350)+8)))
	if v1354 <= int32(0) {
		v1376 = v1349
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+76))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+12))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1358+v1354<<(uint(int32(2))%32)-int32(4))))
	v1365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1364)+8)))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+4))
	v1372 = F_find_window_run_conditions(m, v1182, v1365, v1366, v1256, int32(1), v31+int32(140), v31+int32(124))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L4
	} else {
		goto L289
	}
L289:
	;
	if v1372 != 0 {
		goto L285
	} else {
		goto L290
	}
L290:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+28))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+12))
	v1376 = v1375
	goto L286
L291:
	;
	v1382 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1378)+8)))
	if v1382 <= int32(0) {
		goto L248
	} else {
		goto L292
	}
L292:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+76))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+12))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1386+v1382<<(uint(int32(2))%32)-int32(4))))
	v1393 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1392)+8)))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+4))
	v1400 = F_find_window_run_conditions(m, v1182, v1393, v1394, v1256, int32(0), v31+int32(140), v31+int32(124))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L4
	} else {
		goto L293
	}
L293:
	;
	if v1400 == int32(0) {
		goto L248
	} else {
		goto L294
	}
L294:
	;
	goto L285
L295:
	;
	v1492 = v1238
	goto L245
L296:
	;
	goto L248
L297:
	;
	v1492 = v1441
	goto L245
L298:
	;
	goto L246
L299:
	;
	v1492 = v1238
	goto L245
L300:
	;
	goto L244
L301:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1522)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1523)+68))
	if v1526 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L302:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1522 = v1508 + v1509<<(uint(int32(2))%32)
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1513)+52))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+12))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1522 = v1515 + v1516<<(uint(int32(2))%32) - int32(4)
	goto L301
L305:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L4
	} else {
		goto L315
	}
L306:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+4))
	if v1529 <= int32(0) {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v1535 = int32(0)
	goto L308
L308:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+12))
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1561+v1535<<(uint(int32(2))%32))))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	v1567 = F_expression_returns_set_rows(m, l0, v1566)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L4
	} else {
		goto L310
	}
L309:
	;
	goto L305
L310:
	;
	v1569 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	if base.F64_gt(v1567, v1569) != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v1567
	goto L313
L312:
	;
	goto L313
L313:
	;
	v1573 = v1535 + int32(1)
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+4))
	if v1573 < v1574 {
		v1535 = v1573
		goto L308
	} else {
		goto L314
	}
L314:
	;
	goto L309
L315:
	;
	goto L1
L316:
	;
	goto L1
L317:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1624)))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+80))
	if v1626 != 0 {
		goto L321
	} else {
		goto L322
	}
L318:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1624 = v1610 + v1611<<(uint(int32(2))%32)
	goto L317
L319:
	;
	goto L320
L320:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1615)+52))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1616)+12))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1624 = v1617 + v1618<<(uint(int32(2))%32) - int32(4)
	goto L317
L321:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+4))
	v1630 = base.F64_convert_i32_s(v1627)
	goto L323
L322:
	;
	v1630 = float64(0)
	goto L323
L323:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v1630
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L4
	} else {
		goto L324
	}
L324:
	;
	goto L1
L325:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+4))
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+48))
	if v1769 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L326:
	;
	v1640 = l0
	v1642 = v1634
	goto L329
L327:
	;
	goto L328
L328:
	;
	if v1634 == int32(0) {
		goto L208
	} else {
		goto L336
	}
L329:
	;
	if v1642 == int32(0) {
		goto L325
	} else {
		goto L331
	}
L330:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L4
	} else {
		goto L333
	}
L331:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+16))
	if v1670 != 0 {
		v1640 = v1670
		v1642 = v1642 - int32(1)
		goto L329
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1675
	F_errmsg_internal(m, int32(_a_F_set_rel_size_0), v31+int32(112))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(2929), int32(_a_F_set_rel_size_2))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	v1691 = l0
	v1693 = v1634
	goto L338
L337:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+348))
	if v1738 == int32(0) {
		goto L207
	} else {
		goto L345
	}
L338:
	;
	v1718 = v1693 - int32(1)
	if v1718 == int32(0) {
		goto L337
	} else {
		goto L340
	}
L339:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L4
	} else {
		goto L342
	}
L340:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+16))
	if v1721 != 0 {
		v1691 = v1721
		v1693 = v1718
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v1726
	F_errmsg_internal(m, int32(_a_F_set_rel_size_0), v31+int32(48))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3061), int32(_a_F_set_rel_size_3))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L4
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	v1741 = *(*float64)(unsafe.Add(mBase, uint32(v1738)+32))
	F_set_cte_size_estimates(m, l0, l1, v1741)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1746 = F_palloc0(m, int32(72))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1746)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1746))) = int64(1516123455767)
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1746)+12)) = v1751
	v1753 = F_get_baserel_parampathinfo(m, l0, l1, v1744)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L4
	} else {
		goto L348
	}
L348:
	;
	v1755 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1746)+20)) = uint8(v1755)
	*(*int32)(unsafe.Add(mBase, uint32(v1746)+16)) = v1753
	v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1746)+64)) = v1755
	*(*int32)(unsafe.Add(mBase, uint32(v1746)+24)) = v1755
	*(*uint8)(unsafe.Add(mBase, uint32(v1746)+21)) = uint8(v1758)
	F_cost_ctescan(m, v1746, l0, l1, v1753)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L4
	} else {
		goto L349
	}
L349:
	;
	F_add_path(m, l1, v1746)
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	goto L1
L351:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+76))
	if v1886 == int32(0) {
		goto L206
	} else {
		goto L369
	}
L352:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L4
	} else {
		goto L366
	}
L353:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1769)+4))
	if v1772 <= int32(0) {
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1769)+12))
	v1782 = int32(0)
	goto L355
L355:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1776+v1782<<(uint(int32(2))%32))))
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1809)+4))
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1810))))
	v1816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1775))))
	if base.B2i32(v1813 == int32(0))|base.B2i32(v1813 != v1816) != 0 {
		v1834 = v1813
		v1835 = v1816
		goto L358
	} else {
		goto L359
	}
L356:
	;
	goto L352
L357:
	;
	if v1834-v1835 == int32(0) {
		goto L351
	} else {
		goto L364
	}
L358:
	;
	goto L357
L359:
	;
	v1819 = v1810
	v1820 = v1775
	goto L360
L360:
	;
	v1823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820)+1)))
	v1824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819)+1)))
	if v1824 == int32(0) {
		v1834 = v1824
		v1835 = v1823
		goto L358
	} else {
		goto L362
	}
L361:
	;
	v1834 = v1824
	v1835 = v1823
	goto L358
L362:
	;
	v1827 = int32(1)
	if v1824 == v1823 {
		v1819 = v1819 + v1827
		v1820 = v1820 + v1827
		goto L360
	} else {
		goto L363
	}
L363:
	;
	goto L361
L364:
	;
	v1840 = v1782 + int32(1)
	if v1772 != v1840 {
		v1782 = v1840
		goto L355
	} else {
		goto L365
	}
L365:
	;
	goto L356
L366:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1874
	F_errmsg_internal(m, int32(_a_F_set_rel_size_4), v31-int32(-64))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L4
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(2947), int32(_a_F_set_rel_size_2))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1886)+4))
	if v1889 <= v1782 {
		goto L206
	} else {
		goto L370
	}
L370:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1886)+12))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1891+v1782<<(uint(int32(2))%32))))
	if v1895 <= int32(0) {
		goto L205
	} else {
		goto L371
	}
L371:
	;
	v1901 = v1895<<(uint(int32(2))%32) - int32(4)
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1902)+12))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1903)+12))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1901+v1904)))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1902)+8))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1907)+12))
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1908+v1901)))
	v1911 = *(*float64)(unsafe.Add(mBase, uint32(v1910)+24))
	F_set_cte_size_estimates(m, l0, l1, v1911)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L4
	} else {
		goto L372
	}
L372:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+64))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+44))
	v1916 = F_convert_subquery_pathkeys(m, l0, l1, v1914, v1915)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L4
	} else {
		goto L373
	}
L373:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1920 = F_palloc0(m, int32(72))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1920)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1920))) = int64(1507533521175)
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1920)+12)) = v1925
	v1927 = F_get_baserel_parampathinfo(m, l0, l1, v1918)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L4
	} else {
		goto L375
	}
L375:
	;
	v1929 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1920)+20)) = uint8(v1929)
	*(*int32)(unsafe.Add(mBase, uint32(v1920)+16)) = v1927
	v1932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1920)+64)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v1920)+24)) = v1929
	*(*uint8)(unsafe.Add(mBase, uint32(v1920)+21)) = uint8(v1932)
	F_cost_ctescan(m, v1920, l0, l1, v1927)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	F_add_path(m, l1, v1920)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L4
	} else {
		goto L377
	}
L377:
	;
	goto L1
L378:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1956)))
	v1958 = *(*float64)(unsafe.Add(mBase, uint32(v1957)+112))
	if base.F64_lt(v1958, float64(0)) != 0 {
		goto L382
	} else {
		goto L383
	}
L379:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1956 = v1942 + v1943<<(uint(int32(2))%32)
	goto L378
L380:
	;
	goto L381
L381:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1947)+52))
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1948)+12))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1956 = v1949 + v1950<<(uint(int32(2))%32) - int32(4)
	goto L378
L382:
	;
	v1961 = float64(1000)
	goto L384
L383:
	;
	v1961 = v1958
	goto L384
L384:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v1961
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1967 = F_palloc0(m, int32(72))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1967))) = int64(1511828488471)
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+12)) = v1972
	v1974 = F_get_baserel_parampathinfo(m, l0, l1, v1965)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	v1976 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1967)+20)) = uint8(v1976)
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+16)) = v1974
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+64)) = v1976
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+24)) = v1976
	*(*uint8)(unsafe.Add(mBase, uint32(v1967)+21)) = uint8(v1979)
	v1985 = float64(0)
	v1986 = m.G0
	v1988 = v1986 - int32(32)
	m.G0 = v1988
	if v1974 != 0 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	v2119 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v2120 = float64(0)
	v2121 = base.F64_add(v2118, v2120)
	*(*float64)(unsafe.Add(mBase, uint32(v1967)+48)) = v2121
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+40)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v1967)+56)) = base.F64_add(v2121, base.F64_add(base.F64_mul(v2119, base.F64_add(v2115, base.F64_add(v2114, v2116))), v2120))
	m.G0 = v1988 + int32(32)
	F_add_path(m, l1, v1967)
	mBase = m.M
	v2136 = m.ExcPending
	if v2136 != 0 {
		goto L4
	} else {
		goto L399
	}
L389:
	;
	v1990 = *(*float64)(unsafe.Add(mBase, uint32(v1974)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1967)+32)) = v1990
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+16))
	v1993 = int32(0)
	v1995 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_size[1]))
	v1996 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1988)+16)) = v1996
	*(*int32)(unsafe.Add(mBase, uint32(v1988)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v1988)+24)) = v1996
	if v1992 == v1993 {
		v2074 = v1985
		v2075 = v24
		v2079 = v1995
		goto L392
	} else {
		goto L393
	}
L390:
	;
	goto L391
L391:
	;
	v2084 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1967)+32)) = v2084
	v2086 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v2088 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_size[1]))
	v2089 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v2114 = v2086
	v2115 = v2088
	v2116 = v2088
	v2118 = v2089
	goto L388
L392:
	;
	v2080 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v2082 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v2114 = base.F64_add(v2075, v2080)
	v2115 = v1995
	v2116 = v2079
	v2118 = base.F64_add(v2074, v2082)
	goto L388
L393:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	if v2003 <= int32(0) {
		v2074 = v1985
		v2075 = v24
		v2079 = v1995
		goto L392
	} else {
		goto L394
	}
L394:
	;
	v2009 = v1993
	goto L395
L395:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+12))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2034+v2009<<(uint(int32(2))%32))))
	v2041 = F_cost_qual_eval_walker(m, v2038, v1988+int32(8))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L4
	} else {
		goto L397
	}
L396:
	;
	v2047 = *(*float64)(unsafe.Add(mBase, uint32(v1988)+24))
	v2048 = *(*float64)(unsafe.Add(mBase, uint32(v1988)+16))
	v2050 = *(*float64)(unsafe.Add(mBase, _c_F_set_rel_size[1]))
	v2074 = v2048
	v2075 = v2047
	v2079 = v2050
	goto L392
L397:
	;
	v2044 = v2009 + int32(1)
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	if v2044 < v2045 {
		v2009 = v2044
		goto L395
	} else {
		goto L398
	}
L398:
	;
	goto L396
L399:
	;
	goto L1
L400:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2143 = F_palloc0(m, int32(72))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L4
	} else {
		goto L401
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2143))) = int64(1421634175255)
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+12)) = v2148
	v2150 = F_get_baserel_parampathinfo(m, l0, l1, v2141)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L4
	} else {
		goto L402
	}
L402:
	;
	v2152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2143)+20)) = uint8(v2152)
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+16)) = v2150
	v2155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+64)) = v2152
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+24)) = v2152
	*(*uint8)(unsafe.Add(mBase, uint32(v2143)+21)) = uint8(v2155)
	F_cost_resultscan(m, v2143, l0, l1, v2150)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L4
	} else {
		goto L403
	}
L403:
	;
	F_add_path(m, l1, v2143)
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L4
	} else {
		goto L404
	}
L404:
	;
	goto L1
L405:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v2169
	F_errmsg_internal(m, int32(_a_F_set_rel_size_5), v31)
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L4
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(453), int32(_a_F_set_rel_size_6))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L4
	} else {
		goto L407
	}
L407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L408:
	;
	v2267 = *(*int64)(unsafe.Add(mBase, uint32(v2192)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+200)) = v2267
	v2269 = *(*int64)(unsafe.Add(mBase, uint32(v2192)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+192)) = v2269
	F_set_rel_width(m, l0, l1)
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L4
	} else {
		goto L415
	}
L409:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+4))
	if v2195 <= int32(0) {
		goto L408
	} else {
		goto L410
	}
L410:
	;
	v2202 = v5
	goto L411
L411:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+12))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2226+v2202<<(uint(int32(2))%32))))
	v2233 = F_cost_qual_eval_walker(m, v2230, v2181+int32(8))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L4
	} else {
		goto L413
	}
L412:
	;
	goto L408
L413:
	;
	v2236 = v2202 + int32(1)
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+4))
	if v2236 < v2237 {
		v2202 = v2236
		goto L411
	} else {
		goto L414
	}
L414:
	;
	goto L412
L415:
	;
	m.G0 = v2181 + int32(32)
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2277)+4))
	m.T0[v2278].(func(*base.Module, int32, int32, int32))(m, l0, l1, v2276)
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L4
	} else {
		goto L416
	}
L416:
	;
	v2281 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v2282 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v2281)&int64(9223372036854775807)))|base.F64_gt(v2281, v2282) != 0 {
		v2295 = v2282
		goto L418
	} else {
		goto L419
	}
L417:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v2295
	v2297 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	if base.F64_lt(v2295, v2297) != 0 {
		goto L421
	} else {
		goto L422
	}
L418:
	;
	goto L417
L419:
	;
	v2291 = float64(1)
	if base.F64_le(v2281, v2291) != 0 {
		v2295 = v2291
		goto L418
	} else {
		goto L420
	}
L420:
	;
	v2295 = base.F64_nearest(v2281)
	goto L418
L421:
	;
	v2299 = v2297
	goto L423
L422:
	;
	v2299 = v2295
	goto L423
L423:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v2299
	goto L1
L424:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v2305
	F_errmsg_internal(m, int32(_a_F_set_rel_size_0), v31+int32(16))
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3054), int32(_a_F_set_rel_size_3))
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L4
	} else {
		goto L426
	}
L426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L427:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v2321
	F_errmsg_internal(m, int32(_a_F_set_rel_size_7), v31+int32(32))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L4
	} else {
		goto L428
	}
L428:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(3065), int32(_a_F_set_rel_size_3))
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L4
	} else {
		goto L429
	}
L429:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L430:
	;
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v2337
	F_errmsg_internal(m, int32(_a_F_set_rel_size_8), v31+int32(80))
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L4
	} else {
		goto L431
	}
L431:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(2949), int32(_a_F_set_rel_size_2))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L4
	} else {
		goto L432
	}
L432:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L433:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v2353
	F_errmsg_internal(m, int32(_a_F_set_rel_size_9), v31+int32(96))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L4
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_set_rel_size_1), int32(2952), int32(_a_F_set_rel_size_2))
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L4
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v2400
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+144))
	if v2428 != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v2620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1178)+36)))
	if v2620 != 0 {
		v2676 = v24
		goto L474
	} else {
		goto L475
	}
L438:
	;
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+120))
	if v2429 != 0 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182)+40)))
	if v2430 != int32(1) {
		goto L437
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+4))
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v2434, v2435, v31+int32(140))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L4
	} else {
		goto L443
	}
L442:
	;
	goto L441
L443:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v2440 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v2520 = F_bms_is_member(m, int32(7), v2519)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L4
	} else {
		goto L451
	}
L445:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2440)+4))
	if v2443 <= int32(0) {
		goto L444
	} else {
		goto L446
	}
L446:
	;
	v2451 = int32(0)
	goto L447
L447:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2440)+12))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2475+v2451<<(uint(int32(2))%32))))
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(v2479)+4))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v2480, v2481, v31+int32(140))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L4
	} else {
		goto L449
	}
L448:
	;
	goto L444
L449:
	;
	v2487 = v2451 + int32(1)
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2440)+4))
	if v2487 < v2488 {
		v2451 = v2487
		goto L447
	} else {
		goto L450
	}
L450:
	;
	goto L448
L451:
	;
	if v2520 != 0 {
		goto L437
	} else {
		goto L452
	}
L452:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+76))
	if v2522 == int32(0) {
		goto L437
	} else {
		goto L453
	}
L453:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v2522)+4))
	if v2525 <= int32(0) {
		goto L437
	} else {
		goto L454
	}
L454:
	;
	v2533 = int32(0)
	goto L455
L455:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v2522)+12))
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v2557+v2533<<(uint(int32(2))%32))))
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v2561)+16))
	if v2562 != 0 {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	goto L437
L457:
	;
	v2589 = v2533 + int32(1)
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2522)+4))
	if v2589 < v2590 {
		v2533 = v2589
		goto L455
	} else {
		goto L473
	}
L458:
	;
	v2563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+26)))
	if v2563 != 0 {
		goto L457
	} else {
		goto L459
	}
L459:
	;
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(v2561)+4))
	v2565 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2561)+8)))
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v2569 = F_bms_is_member(m, v2565+int32(7), v2568)
	mBase = m.M
	v2570 = m.ExcPending
	if v2570 != 0 {
		goto L4
	} else {
		goto L460
	}
L460:
	;
	if v2569 != 0 {
		goto L457
	} else {
		goto L461
	}
L461:
	;
	v2571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182)+38)))
	if v2571 == int32(1) {
		goto L462
	} else {
		goto L463
	}
L462:
	;
	v2574 = F_expression_returns_set(m, v2564)
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L4
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	v2576 = F_contain_volatile_functions(m, v2564)
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L4
	} else {
		goto L467
	}
L465:
	;
	if v2574 != 0 {
		goto L457
	} else {
		goto L466
	}
L466:
	;
	goto L464
L467:
	;
	if v2576 != 0 {
		goto L457
	} else {
		goto L468
	}
L468:
	;
	v2578 = F_exprType(m, v2564)
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L4
	} else {
		goto L469
	}
L469:
	;
	v2580 = F_exprTypmod(m, v2564)
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L4
	} else {
		goto L470
	}
L470:
	;
	v2582 = F_exprCollation(m, v2564)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L4
	} else {
		goto L471
	}
L471:
	;
	v2584 = F_makeNullConst(m, v2578, v2580, v2582)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L4
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2561)+4)) = v2584
	goto L457
L473:
	;
	goto L456
L474:
	;
	v2677 = int32(0)
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2681 = F_subquery_planner(m, v2678, v1182, l0, v2677, v2676, v2677)
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L4
	} else {
		goto L498
	}
L475:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+100))
	if v2621 != 0 {
		v2676 = v24
		goto L474
	} else {
		goto L476
	}
L476:
	;
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+108))
	if v2622 != 0 {
		v2676 = v24
		goto L474
	} else {
		goto L477
	}
L477:
	;
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+318)))
	if v2623 != 0 {
		v2676 = v24
		goto L474
	} else {
		goto L478
	}
L478:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+120))
	if v2624 != 0 {
		v2676 = v24
		goto L474
	} else {
		goto L479
	}
L479:
	;
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+124))
	if v2625 != 0 {
		v2676 = v24
		goto L474
	} else {
		goto L480
	}
L480:
	;
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v2627 = int32(0)
	if v2626 == v2627 {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	if v2672 == int32(2) {
		v2676 = v24
		goto L474
	} else {
		goto L497
	}
L482:
	;
	v2672 = int32(0)
	goto L481
L483:
	;
	goto L484
L484:
	;
	v2635 = int32(1)
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2626)+4))
	if v2636 <= v2635 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2639 = v2635
	goto L487
L486:
	;
	v2639 = v2636
	goto L487
L487:
	;
	v2643 = int32(0)
	v2645 = v2627
	goto L488
L488:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v2626+int32(8)+v2643<<(uint(int32(2))%32))))
	if v2652 != 0 {
		goto L491
	} else {
		goto L492
	}
L489:
	;
	v2672 = v2664
	goto L481
L490:
	;
	goto L489
L491:
	;
	v2653 = int32(2)
	if v2645 != 0 {
		v2664 = v2653
		goto L490
	} else {
		goto L494
	}
L492:
	;
	v2659 = v2645
	goto L493
L493:
	;
	v2661 = v2643 + int32(1)
	if v2661 != v2639 {
		v2643 = v2661
		v2645 = v2659
		goto L488
	} else {
		goto L496
	}
L494:
	;
	v2654 = int32(1)
	if base.Ui32(v2654) < base.Ui32(base.I32_popcnt(v2652)) {
		v2664 = v2653
		goto L490
	} else {
		goto L495
	}
L495:
	;
	v2659 = v2654
	goto L493
L496:
	;
	v2664 = v2659
	goto L490
L497:
	;
	v2675 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	v2676 = v2675
	goto L474
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v2681
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+144)) = v2684
	v2686 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2686
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v2691 = F_fetch_upper_rel(m, v2688, int32(7), v2686)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L4
	} else {
		goto L499
	}
L499:
	;
	v2693 = int32(0)
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(v2691)+32))
	if v2695 == v2693 {
		v2716 = v2693
		goto L501
	} else {
		goto L502
	}
L500:
	;
	if v2716 != 0 {
		goto L510
	} else {
		goto L511
	}
L501:
	;
	goto L500
L502:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2695)+12))
	v2699 = v2698
	goto L503
L503:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2699)))
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2702)))
	if base.Ui32(int32(2)) <= base.Ui32(v2703-int32(301)) {
		goto L505
	} else {
		goto L506
	}
L504:
	;
	v2716 = int32(1)
	goto L501
L505:
	;
	if v2703 != int32(290) {
		v2716 = v2693
		goto L501
	} else {
		goto L508
	}
L506:
	;
	v2699 = v2702 + int32(72)
	goto L503
L507:
	;
	goto L504
L508:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+72))
	if v2710 != 0 {
		v2716 = v2693
		goto L501
	} else {
		goto L509
	}
L509:
	;
	goto L507
L510:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2720 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2719)+32)) = v2720
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v2720
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v2720
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2734 = F_create_append_path(m, v2720, l1, v2720, v2720, v2720, v2730, v2720, v2720, float64(-1))
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L4
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	F_set_subquery_size_estimates(m, l0, l1)
	mBase = m.M
	v2741 = m.ExcPending
	if v2741 != 0 {
		goto L4
	} else {
		goto L516
	}
L513:
	;
	F_add_path(m, l1, v2734)
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L4
	} else {
		goto L514
	}
L514:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L4
	} else {
		goto L515
	}
L515:
	;
	goto L1
L516:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v2742)+4))
	if v2743 != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(v2743)+4))
	v2745 = v2744
	goto L519
L518:
	;
	v2745 = v2677
	goto L519
L519:
	;
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+76))
	if v2747 != 0 {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v2832 = *(*int32)(unsafe.Add(mBase, uint32(v2691)+32))
	if v2832 == int32(0) {
		goto L537
	} else {
		goto L538
	}
L521:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2747)+4))
	v2750 = v2748
	goto L523
L522:
	;
	v2750 = int32(0)
	goto L523
L523:
	;
	if v2750 != v2745 {
		v2812 = int32(0)
		goto L520
	} else {
		goto L524
	}
L524:
	;
	if v2743 == int32(0) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2812 = int32(1)
	goto L520
L526:
	;
	goto L527
L527:
	;
	v2755 = int32(0)
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2743)+4))
	if v2755 < v2756 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v2760 = v2756
	goto L530
L529:
	;
	v2760 = v2755
	goto L530
L530:
	;
	v2765 = v2755
	goto L531
L531:
	;
	v2789 = base.B2i32(v2765 == v2760)
	if v2765 == v2760 {
		v2812 = v2789
		goto L520
	} else {
		goto L533
	}
L532:
	;
	v2812 = v2789
	goto L520
L533:
	;
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2743)+12))
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2790+v2765<<(uint(int32(2))%32))))
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2794)))
	if v2795 != int32(6) {
		v2812 = v2789
		goto L520
	} else {
		goto L534
	}
L534:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2794)+4))
	if v2798 != l2 {
		v2812 = v2789
		goto L520
	} else {
		goto L535
	}
L535:
	;
	v2802 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2794)+8)))
	if v2765+int32(1) == v2802 {
		v2765 = v2802
		goto L531
	} else {
		goto L536
	}
L536:
	;
	goto L532
L537:
	;
	v2914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if base.B2i32(v2914 != int32(1))|v1184 != 0 {
		goto L1
	} else {
		goto L547
	}
L538:
	;
	v2835 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+4))
	if v2835 <= int32(0) {
		goto L537
	} else {
		goto L539
	}
L539:
	;
	v2842 = int32(0)
	goto L540
L540:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+12))
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v2867+v2842<<(uint(int32(2))%32))))
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+64))
	v2873 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+12))
	v2874 = F_make_tlist_from_pathtarget(m, v2873)
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L4
	} else {
		goto L542
	}
L541:
	;
	goto L537
L542:
	;
	v2876 = F_convert_subquery_pathkeys(m, l0, l1, v2872, v2874)
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L4
	} else {
		goto L543
	}
L543:
	;
	v2878 = F_create_subqueryscan_path(m, l0, l1, v2871, v2812, v2876, v1184)
	mBase = m.M
	v2879 = m.ExcPending
	if v2879 != 0 {
		goto L4
	} else {
		goto L544
	}
L544:
	;
	F_add_path(m, l1, v2878)
	mBase = m.M
	v2881 = m.ExcPending
	if v2881 != 0 {
		goto L4
	} else {
		goto L545
	}
L545:
	;
	v2883 = v2842 + int32(1)
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+4))
	if v2883 < v2884 {
		v2842 = v2883
		goto L540
	} else {
		goto L546
	}
L546:
	;
	goto L541
L547:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v2691)+40))
	if v2918 == int32(0) {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2918)+4))
	if v2921 <= int32(0) {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	v2928 = int32(0)
	goto L550
L550:
	;
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(v2918)+12))
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(v2953+v2928<<(uint(int32(2))%32))))
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v2957)+64))
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2957)+12))
	v2960 = F_make_tlist_from_pathtarget(m, v2959)
	mBase = m.M
	v2961 = m.ExcPending
	if v2961 != 0 {
		goto L4
	} else {
		goto L552
	}
L551:
	;
	goto L1
L552:
	;
	v2962 = F_convert_subquery_pathkeys(m, l0, l1, v2958, v2960)
	mBase = m.M
	v2963 = m.ExcPending
	if v2963 != 0 {
		goto L4
	} else {
		goto L553
	}
L553:
	;
	v2965 = F_create_subqueryscan_path(m, l0, l1, v2957, v2812, v2962, int32(0))
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L4
	} else {
		goto L554
	}
L554:
	;
	F_add_partial_path(m, l1, v2965)
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L4
	} else {
		goto L555
	}
L555:
	;
	v2970 = v2928 + int32(1)
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2918)+4))
	if v2970 < v2971 {
		v2928 = v2970
		goto L550
	} else {
		goto L556
	}
L556:
	;
	goto L551
}
