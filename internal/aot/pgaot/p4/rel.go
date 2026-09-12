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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	switch l2 - int32(112) {
	case 0, 5:
		v36 = int32(-1)
		goto L1
	default:
		goto L3
	case 4:
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
	F_errmsg_internal(m, int32(503727), v9)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(499324), int32(581), int32(228968))
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
	v36 = v35
	goto L1
L11:
	;
	v42 = l0
	goto L13
L12:
	;
	v42 = v41
	goto L13
L13:
	;
	if v42 != int32(1664) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v38
	goto L16
L15:
	;
	v45 = int32(0)
	goto L16
L16:
	;
	goto L17
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	m.G0 = v9 + int32(80)
	return v64
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	F_GetRelationPath(m, v9+int32(8), v45, v42, v64, v36, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L29
	}
L24:
	;
	v60 = F_GetNewOidWithIndex(m, l1, int32(2662), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v62 = F_GetNewObjectId(m)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	v64 = v60
	goto L23
L28:
	;
	v64 = v62
	goto L23
L29:
	;
	v70 = int32(0)
	v71 = F_access(m, v9+int32(8), v70)
	mBase = m.M
	if v71 == v70 {
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
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
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
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if l7 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v178 = v10
	v180 = v10
	goto L3
L6:
	;
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v33 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v178 = v33
	v180 = v10
	goto L3
L9:
	;
	goto L10
L10:
	;
	v38 = F_palloc(m, v33<<(uint(int32(1))%32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v40 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v178 = int32(0)
	v180 = v38
	goto L3
L13:
	;
	goto L14
L14:
	;
	v54 = int32(0)
	v59 = v10
	v61 = v40
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
	v69 = int32(0)
	if v54 <= v69 {
		v110 = v69
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v178 = v156
	v180 = v38
	goto L3
L17:
	;
	v167 = v59 + int32(1)
	if v167 < v163 {
		v54 = v156
		v59 = v167
		v61 = v163
		goto L15
	} else {
		goto L26
	}
L18:
	;
	v139 = int32(1)
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v38+v54<<(uint(v139)%32)))) = uint16(v142)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v156 = v54 + v139
	v163 = v146
	goto L17
L19:
	;
	if v54 != v110 {
		v156 = v54
		v163 = v61
		goto L17
	} else {
		goto L25
	}
L20:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+8)))
	v83 = v69
	goto L21
L21:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+v83<<(uint(int32(1))%32)))))
	if v95 == v72 {
		v110 = v83
		goto L19
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v98 = v83 + int32(1)
	if v98 != v54 {
		v83 = v98
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L18
L26:
	;
	goto L16
L27:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v188)+68))
	v215 = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v227 = int32(32)
	v234 = F_CreateConstraintEntry(m, l1, v213, int32(99), v215, v215, l3, l4, v215, v218, v180, v178, v178, v215, v215, v215, v215, v215, v215, v215, v215, v227, v227, v215, v215, v227, v215, l2, v24, l5, l6, l7, v215, l8)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L34
	}
L28:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+119)))
	if v191 != int32(112) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v201 + int32(4)
	F_errmsg(m, int32(722007), v22)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(497201), int32(2203), int32(319126))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_pfree(m, v24)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	m.G0 = v22 + int32(16)
	return v234
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
				F_errmsg_internal(m, int32(46515), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500396), int32(2226), int32(417366))
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 float64
	_ = v154
	var v159 float64
	_ = v159
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v183 int64
	_ = v183
	var v188 int64
	_ = v188
	var v193 int64
	_ = v193
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 float64
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v236 int64
	_ = v236
	var v258 float64
	_ = v258
	var v263 float64
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 float64
	_ = v273
	var v274 float64
	_ = v274
	var v276 float64
	_ = v276
	var v277 float64
	_ = v277
	var v280 float64
	_ = v280
	var v283 float64
	_ = v283
	var v287 float64
	_ = v287
	var v290 float64
	_ = v290
	var v294 float64
	_ = v294
	var v296 int64
	_ = v296
	var v301 int32
	_ = v301
	var v304 float64
	_ = v304
	var v307 float64
	_ = v307
	var v310 int64
	_ = v310
	var v313 int64
	_ = v313
	var v322 float64
	_ = v322
	var v324 float64
	_ = v324
	var v328 float64
	_ = v328
	var v329 float64
	_ = v329
	var v330 float64
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 float64
	_ = v339
	var v342 int32
	_ = v342
	var v346 float64
	_ = v346
	var v347 float64
	_ = v347
	var v348 float64
	_ = v348
	var v357 float64
	_ = v357
	var v360 float64
	_ = v360
	var v363 float64
	_ = v363
	var v370 float64
	_ = v370
	var v371 float64
	_ = v371
	var v382 float64
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 float64
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v433 int32
	_ = v433
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v576 int32
	_ = v576
	var v610 int32
	_ = v610
	var v615 int64
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 float64
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 float64
	_ = v681
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v782 int32
	_ = v782
	var v811 float64
	_ = v811
	var v813 float64
	_ = v813
	var v815 float64
	_ = v815
	var v816 int32
	_ = v816
	var v817 float64
	_ = v817
	var v857 float64
	_ = v857
	var v858 int32
	_ = v858
	var v861 float64
	_ = v861
	var v863 float64
	_ = v863
	var v867 float64
	_ = v867
	var v872 float64
	_ = v872
	var v876 int32
	_ = v876
	var v877 float64
	_ = v877
	var v917 int32
	_ = v917
	var v920 float64
	_ = v920
	var v922 float64
	_ = v922
	var v926 float64
	_ = v926
	var v931 float64
	_ = v931
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v990 int32
	_ = v990
	var v991 float64
	_ = v991
	var v1031 int32
	_ = v1031
	var v1034 float64
	_ = v1034
	var v1036 float64
	_ = v1036
	var v1040 float64
	_ = v1040
	var v1045 float64
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1109 int32
	_ = v1109
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1202 float64
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1218 int32
	_ = v1218
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1311 float64
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1377 int32
	_ = v1377
	var v1378 int64
	_ = v1378
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1397 int32
	_ = v1397
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1522 int32
	_ = v1522
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1585 int32
	_ = v1585
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1741 int32
	_ = v1741
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1896 int32
	_ = v1896
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2053 int32
	_ = v2053
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2098 int32
	_ = v2098
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2150 int32
	_ = v2150
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2249 int64
	_ = v2249
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2285 int32
	_ = v2285
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2319 int64
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int64
	_ = v2323
	var v2325 int64
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2346 int32
	_ = v2346
	var v2370 int32
	_ = v2370
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2389 int32
	_ = v2389
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2444 int32
	_ = v2444
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2493 int32
	_ = v2493
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2534 int64
	_ = v2534
	var v2536 int64
	_ = v2536
	var v2538 int64
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2590 int32
	_ = v2590
	var v2598 int32
	_ = v2598
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2608 int32
	_ = v2608
	var v2610 int32
	_ = v2610
	var v2614 int32
	_ = v2614
	var v2619 int64
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2672 int32
	_ = v2672
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2750 int32
	_ = v2750
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2812 int32
	_ = v2812
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2821 int32
	_ = v2821
	var v2825 int32
	_ = v2825
	var v2830 int64
	_ = v2830
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2902 int32
	_ = v2902
	var v2942 int32
	_ = v2942
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2956 int32
	_ = v2956
	var v2960 int32
	_ = v2960
	var v2965 int64
	_ = v2965
	var v2967 int32
	_ = v2967
	var v2969 int32
	_ = v2969
	var v3010 int32
	_ = v3010
	var v3014 int32
	_ = v3014
	var v3016 int32
	_ = v3016
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3030 int32
	_ = v3030
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3074 int32
	_ = v3074
	var v3117 int32
	_ = v3117
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3129 int32
	_ = v3129
	var v3134 int64
	_ = v3134
	var v3139 int32
	_ = v3139
	var v3143 int32
	_ = v3143
	var v3148 int32
	_ = v3148
	var v3191 int32
	_ = v3191
	var v3195 int32
	_ = v3195
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3242 int32
	_ = v3242
	var v3249 int32
	_ = v3249
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3295 float64
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3307 float64
	_ = v3307
	var v3310 int32
	_ = v3310
	var v3313 int32
	_ = v3313
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3353 float64
	_ = v3353
	var v3358 float64
	_ = v3358
	var v3363 float64
	_ = v3363
	var v3372 int32
	_ = v3372
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3383 int32
	_ = v3383
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3401 int32
	_ = v3401
	var v3409 int32
	_ = v3409
	var v3413 int32
	_ = v3413
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3419 int32
	_ = v3419
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3427 int32
	_ = v3427
	var v3429 int32
	_ = v3429
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3435 int32
	_ = v3435
	var v3437 int32
	_ = v3437
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3443 int32
	_ = v3443
	var v3445 int32
	_ = v3445
	var v3448 int32
	_ = v3448
	var v3450 int32
	_ = v3450
	var v3455 int32
	_ = v3455
	var v3466 int32
	_ = v3466
	var v3471 int32
	_ = v3471
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3476 int32
	_ = v3476
	var v3478 int32
	_ = v3478
	var v3480 int32
	_ = v3480
	var v3483 int32
	_ = v3483
	var v3494 float64
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3498 int32
	_ = v3498
	var v3501 int32
	_ = v3501
	var v3504 int32
	_ = v3504
	var v3507 float64
	_ = v3507
	var v3509 int32
	_ = v3509
	var v3512 int32
	_ = v3512
	var v3538 float64
	_ = v3538
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3553 float64
	_ = v3553
	var v3554 int32
	_ = v3554
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3601 int32
	_ = v3601
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3652 int32
	_ = v3652
	var v3656 int32
	_ = v3656
	var v3661 int32
	_ = v3661
	var v3663 int32
	_ = v3663
	var v3665 int32
	_ = v3665
	var v3667 int32
	_ = v3667
	var v3668 int32
	_ = v3668
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3680 int32
	_ = v3680
	var v3715 int32
	_ = v3715
	var v3717 int32
	_ = v3717
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3762 int32
	_ = v3762
	var v3764 int32
	_ = v3764
	var v3766 int32
	_ = v3766
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3789 int32
	_ = v3789
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3829 int32
	_ = v3829
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3836 int32
	_ = v3836
	var v3838 int32
	_ = v3838
	var v3839 int32
	_ = v3839
	var v3842 int32
	_ = v3842
	var v3846 int32
	_ = v3846
	var v3852 int32
	_ = v3852
	var v3886 int32
	_ = v3886
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3899 int32
	_ = v3899
	var v3902 int32
	_ = v3902
	var v3908 int32
	_ = v3908
	var v3910 int32
	_ = v3910
	var v3911 int32
	_ = v3911
	var v3915 int32
	_ = v3915
	var v3916 int32
	_ = v3916
	var v3917 int32
	_ = v3917
	var v3920 int32
	_ = v3920
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3933 int32
	_ = v3933
	var v3936 int32
	_ = v3936
	var v3943 int32
	_ = v3943
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3949 int32
	_ = v3949
	var v3950 int32
	_ = v3950
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3979 int32
	_ = v3979
	var v3987 int32
	_ = v3987
	var v3990 int32
	_ = v3990
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4004 int32
	_ = v4004
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4045 int32
	_ = v4045
	var v4046 int32
	_ = v4046
	var v4047 int32
	_ = v4047
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4059 int32
	_ = v4059
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4068 int32
	_ = v4068
	var v4075 int32
	_ = v4075
	var v4077 int32
	_ = v4077
	var v4079 int32
	_ = v4079
	var v4080 int32
	_ = v4080
	var v4082 int32
	_ = v4082
	var v4084 int32
	_ = v4084
	var v4088 int32
	_ = v4088
	var v4092 int32
	_ = v4092
	var v4093 int32
	_ = v4093
	var v4096 int32
	_ = v4096
	var v4097 int32
	_ = v4097
	var v4100 int32
	_ = v4100
	var v4101 int32
	_ = v4101
	var v4103 int32
	_ = v4103
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4109 int32
	_ = v4109
	var v4110 int32
	_ = v4110
	var v4116 int32
	_ = v4116
	var v4152 int32
	_ = v4152
	var v4156 int32
	_ = v4156
	var v4157 int32
	_ = v4157
	var v4158 int32
	_ = v4158
	var v4159 int32
	_ = v4159
	var v4168 int32
	_ = v4168
	var v4169 int32
	_ = v4169
	var v4171 int32
	_ = v4171
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4180 int32
	_ = v4180
	var v4187 int32
	_ = v4187
	var v4189 int32
	_ = v4189
	var v4191 int32
	_ = v4191
	var v4192 int32
	_ = v4192
	var v4194 int32
	_ = v4194
	var v4196 int32
	_ = v4196
	var v4200 int32
	_ = v4200
	var v4206 int32
	_ = v4206
	var v4207 int32
	_ = v4207
	var v4209 int32
	_ = v4209
	var v4210 int32
	_ = v4210
	var v4252 int32
	_ = v4252
	var v4253 int32
	_ = v4253
	var v4294 int32
	_ = v4294
	var v4295 int32
	_ = v4295
	var v4307 int32
	_ = v4307
	var v4312 int32
	_ = v4312
	var v4340 int32
	_ = v4340
	var v4343 int32
	_ = v4343
	var v4355 int32
	_ = v4355
	var v4389 int32
	_ = v4389
	var v4393 int32
	_ = v4393
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4399 int32
	_ = v4399
	var v4400 int32
	_ = v4400
	var v4401 int32
	_ = v4401
	var v4402 int32
	_ = v4402
	var v4405 int32
	_ = v4405
	var v4411 int32
	_ = v4411
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4418 int32
	_ = v4418
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4423 int32
	_ = v4423
	var v4424 int32
	_ = v4424
	var v4425 int32
	_ = v4425
	var v4428 int32
	_ = v4428
	var v4429 int32
	_ = v4429
	var v4436 int32
	_ = v4436
	var v4439 int32
	_ = v4439
	var v4446 int32
	_ = v4446
	var v4448 int32
	_ = v4448
	var v4449 int32
	_ = v4449
	var v4452 int32
	_ = v4452
	var v4453 int32
	_ = v4453
	var v4456 int32
	_ = v4456
	var v4457 int32
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4460 int32
	_ = v4460
	var v4461 int32
	_ = v4461
	var v4462 int32
	_ = v4462
	var v4463 int32
	_ = v4463
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4466 int32
	_ = v4466
	var v4467 int32
	_ = v4467
	var v4468 int32
	_ = v4468
	var v4471 int32
	_ = v4471
	var v4472 int32
	_ = v4472
	var v4482 int32
	_ = v4482
	var v4490 int32
	_ = v4490
	var v4493 int32
	_ = v4493
	var v4500 int32
	_ = v4500
	var v4501 int32
	_ = v4501
	var v4506 int32
	_ = v4506
	var v4542 int32
	_ = v4542
	var v4543 int32
	_ = v4543
	var v4547 int32
	_ = v4547
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4558 int32
	_ = v4558
	var v4559 int32
	_ = v4559
	var v4561 int32
	_ = v4561
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4570 int32
	_ = v4570
	var v4577 int32
	_ = v4577
	var v4579 int32
	_ = v4579
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4584 int32
	_ = v4584
	var v4586 int32
	_ = v4586
	var v4590 int32
	_ = v4590
	var v4594 int32
	_ = v4594
	var v4595 int32
	_ = v4595
	var v4598 int32
	_ = v4598
	var v4599 int32
	_ = v4599
	var v4602 int32
	_ = v4602
	var v4603 int32
	_ = v4603
	var v4605 int32
	_ = v4605
	var v4606 int32
	_ = v4606
	var v4648 int32
	_ = v4648
	var v4649 int32
	_ = v4649
	var v4691 int32
	_ = v4691
	var v4692 int32
	_ = v4692
	var v4733 int32
	_ = v4733
	var v4734 int32
	_ = v4734
	var v4735 int32
	_ = v4735
	var v4736 int32
	_ = v4736
	var v4746 int32
	_ = v4746
	var v4780 int32
	_ = v4780
	var v4783 int32
	_ = v4783
	var v4786 int32
	_ = v4786
	var v4790 int32
	_ = v4790
	var v4795 int32
	_ = v4795
	var v4831 int32
	_ = v4831
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4837 int32
	_ = v4837
	var v4838 int32
	_ = v4838
	var v4847 int32
	_ = v4847
	var v4848 int32
	_ = v4848
	var v4850 int32
	_ = v4850
	var v4853 int32
	_ = v4853
	var v4854 int32
	_ = v4854
	var v4859 int32
	_ = v4859
	var v4866 int32
	_ = v4866
	var v4868 int32
	_ = v4868
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4873 int32
	_ = v4873
	var v4875 int32
	_ = v4875
	var v4879 int32
	_ = v4879
	var v4885 int32
	_ = v4885
	var v4886 int32
	_ = v4886
	var v4888 int32
	_ = v4888
	var v4889 int32
	_ = v4889
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4973 int32
	_ = v4973
	var v5013 int32
	_ = v5013
	var v5014 int32
	_ = v5014
	var v5018 int32
	_ = v5018
	var v5022 int32
	_ = v5022
	var v5027 int32
	_ = v5027
	var v5070 int32
	_ = v5070
	var v5074 int32
	_ = v5074
	var v5077 int32
	_ = v5077
	var v5078 int32
	_ = v5078
	var v5082 int32
	_ = v5082
	var v5120 int32
	_ = v5120
	var v5124 int32
	_ = v5124
	var v5126 int32
	_ = v5126
	var v5127 int32
	_ = v5127
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5136 int32
	_ = v5136
	var v5142 int32
	_ = v5142
	var v5143 int32
	_ = v5143
	var v5145 int32
	_ = v5145
	var v5148 int32
	_ = v5148
	var v5149 int32
	_ = v5149
	var v5154 int32
	_ = v5154
	var v5162 int32
	_ = v5162
	var v5164 int32
	_ = v5164
	var v5166 int32
	_ = v5166
	var v5167 int32
	_ = v5167
	var v5170 int32
	_ = v5170
	var v5175 int32
	_ = v5175
	var v5183 int32
	_ = v5183
	var v5185 int32
	_ = v5185
	var v5187 int32
	_ = v5187
	var v5188 int32
	_ = v5188
	var v5230 int32
	_ = v5230
	var v5271 int32
	_ = v5271
	var v5275 int32
	_ = v5275
	var v5281 int32
	_ = v5281
	var v5285 int32
	_ = v5285
	var v5290 int32
	_ = v5290
	var v5291 int32
	_ = v5291
	var v5292 int32
	_ = v5292
	var v5300 int32
	_ = v5300
	v3 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(16)
	m.G0 = v42
	if l1 == v3 {
		v5300 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v42 + int32(16)
	return v5300
L2:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v46 <= int32(0) {
		v5300 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v51 = v3
	v52 = v3
	goto L4
L4:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v51<<(uint(int32(2))%32))))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v46 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v120 = F_lappend(m, v52, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L19
	}
L7:
	;
	if v93 == int32(63) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v117 = F_make_rel_from_joinlist(m, l0, v92)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L18
	}
L10:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v99 = F_find_base_rel(m, l0, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v119 = v99
	goto L6
L15:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v107
	F_errmsg_internal(m, int32(486843), v42)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(495231), int32(3393), int32(73926))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
	v119 = v117
	goto L6
L19:
	;
	v123 = v51 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v123 < v124 {
		v51 = v123
		v52 = v120
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L5
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v5300 = v129
	goto L1
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v120
	v132 = *(*int32)(unsafe.Add(mBase, _consts[618]))
	if v132 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v133 = m.T0[v132].(func(*base.Module, int32, int32, int32) int32)(m, l0, v46, v120)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _consts[619])))
	if v136 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v5300 = v133
	goto L1
L28:
	;
	v3772 = m.G0
	v3774 = v3772 - int32(16)
	m.G0 = v3774
	v3780 = F_palloc0(m, v46<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v3781 = m.ExcPending
	if v3781 != 0 {
		goto L13
	} else {
		goto L431
	}
L29:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[620]))
	if v46 < v140 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v142 = m.G0
	v144 = v142 - int32(32)
	m.G0 = v144
	v146 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v144 + v146
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v120
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v152 = v150 + v146
	v154 = *(*float64)(unsafe.Add(mBase, _consts[621]))
	v159 = base.F64_mul(v154, float64(4.503599627370495e+15))
	if base.F64_lt(base.F64_abs(v159), float64(9.223372036854776e+18)) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[622]))
	if int32(1) < v206 {
		v400 = v206
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v167 = v165 + int64(4354685564936845354)
	v168 = int64(30)
	v171 = int64(-4658895280553007687)
	v172 = (int64(base.Ui64(v167)>>(uint(v168)%64)) ^ v167) * v171
	v173 = int64(27)
	v176 = int64(-7723592293110705685)
	v177 = (int64(base.Ui64(v172)>>(uint(v173)%64)) ^ v172) * v176
	v178 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v152)+8)) = int64(base.Ui64(v177)>>(uint(v178)%64)) ^ v177
	v183 = v165 - int64(7046029254386353131)
	v188 = (int64(base.Ui64(v183)>>(uint(v168)%64)) ^ v183) * v171
	v193 = (int64(base.Ui64(v188)>>(uint(v173)%64)) ^ v188) * v176
	*(*int64)(unsafe.Add(mBase, uint32(v152))) = int64(base.Ui64(v193)>>(uint(v178)%64)) ^ v193
	if v183|v167 == int64(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v163 = base.I64_trunc_f64_s(v159)
	v165 = v163
	goto L32
L34:
	;
	goto L35
L35:
	;
	v165 = int64(-9223372036854775807 - 1)
	goto L32
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v152)+8)) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, uint32(v152))) = int64(6364136223846793005)
	goto L38
L37:
	;
	goto L38
L38:
	;
	goto L31
L39:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[623]))
	v410 = F_palloc(m, int32(12))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L84
	}
L40:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[624]))
	v213 = base.F64_add(base.F64_convert_i32_s(v46), float64(1))
	goto L42
L41:
	;
	v384 = v210 * int32(50)
	if base.F64_gt(v382, base.F64_convert_i32_s(v384)) != 0 {
		v400 = v384
		goto L39
	} else {
		goto L79
	}
L42:
	;
	v219 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v213))>>(uint(int64(52))%64))) & int32(2047)
	v224 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	goto L43
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v224) <= base.Ui32(v219-v224) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if base.Ui32(v219) < base.Ui32(v224) {
		v382 = base.F64_add(v213, float64(1))
		goto L41
	} else {
		goto L48
	}
L46:
	;
	v270 = v219
	goto L47
L47:
	;
	v273 = *(*float64)(unsafe.Add(mBase, _consts[625]))
	v274 = base.F64_add(v213, v273)
	v276 = base.F64_sub(v213, base.F64_sub(v274, v273))
	v277 = base.F64_mul(v276, v276)
	v280 = *(*float64)(unsafe.Add(mBase, _consts[626]))
	v283 = *(*float64)(unsafe.Add(mBase, _consts[627]))
	v287 = *(*float64)(unsafe.Add(mBase, _consts[628]))
	v290 = *(*float64)(unsafe.Add(mBase, _consts[629]))
	v294 = *(*float64)(unsafe.Add(mBase, _consts[630]))
	v296 = base.I64_reinterpret_f64(v274)
	v301 = base.I32_wrap_i64(v296) << (uint(int32(4)) % 32) & int32(2032)
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v301)+uint32(_consts[631])))
	v307 = base.F64_add(base.F64_mul(base.F64_mul(v277, v277), base.F64_add(base.F64_mul(v276, v280), v283)), base.F64_add(base.F64_mul(v277, base.F64_add(base.F64_mul(v276, v287), v290)), base.F64_add(base.F64_mul(v276, v294), v304)))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v301)+uint32(_consts[632])))
	v313 = v310 + v296<<(uint(int64(45))%64)
	if v270 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L48:
	;
	v236 = base.I64_reinterpret_f64(v213)
	goto L50
L49:
	;
	if base.Ui64(v236<<(uint(int64(1))%64)) <= base.Ui64(int64(-9143996093422370816)) {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	if base.Ui32(v219) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	if v236 == int64(-4503599627370496) {
		v382 = float64(0)
		goto L41
	} else {
		goto L52
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(math.Float64frombits(uint64(0x7ff0000000000000))))>>(uint(int64(52))%64)))) <= base.Ui32(v219) {
		v382 = base.F64_add(v213, float64(1))
		goto L41
	} else {
		goto L54
	}
L54:
	;
	if int64(0) <= v236 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v258 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
	mBase = m.M
	goto L58
L56:
	;
	goto L57
L57:
	;
	if base.Ui64(v236) < base.Ui64(int64(-4570929321408987136)) {
		goto L49
	} else {
		goto L59
	}
L58:
	;
	v382 = v258
	goto L41
L59:
	;
	v263 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
	mBase = m.M
	goto L60
L60:
	;
	v382 = v263
	goto L41
L61:
	;
	v269 = v219
	goto L63
L62:
	;
	v269 = int32(0)
	goto L63
L63:
	;
	v270 = v269
	goto L47
L64:
	;
	if v296&int64(2147483648) == int64(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	v371 = base.F64_reinterpret_i64(v313)
	v382 = base.F64_add(base.F64_mul(v371, v307), v371)
	goto L41
L67:
	;
	v382 = v370
	goto L41
L68:
	;
	v322 = base.F64_reinterpret_i64(v313 - int64(4503599627370496))
	v324 = base.F64_add(base.F64_mul(v322, v307), v322)
	v370 = base.F64_add(v324, v324)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v328 = base.F64_reinterpret_i64(v313 + int64(4602678819172646912))
	v329 = base.F64_mul(v328, v307)
	v330 = base.F64_add(v329, v328)
	if base.F64_lt(v330, float64(1)) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v334 = m.G0
	v336 = v334 - int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v336)+8)) = int64(4503599627370496)
	v339 = *(*float64)(unsafe.Add(mBase, uint32(v336)+8))
	goto L74
L72:
	;
	v363 = v330
	goto L73
L73:
	;
	v370 = base.F64_mul(v363, float64(2.2250738585072014e-308))
	goto L67
L74:
	;
	v342 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v342-int32(16))+8)) = base.F64_mul(v339, float64(2.2250738585072014e-308))
	goto L75
L75:
	;
	v346 = float64(0)
	v347 = float64(1)
	v348 = base.F64_add(v330, v347)
	v357 = base.F64_add(base.F64_add(v348, base.F64_add(base.F64_add(v329, base.F64_sub(v328, v330)), base.F64_add(v330, base.F64_sub(v347, v348)))), float64(-1))
	if base.F64_eq(v357, v346) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v360 = v346
	goto L78
L77:
	;
	v360 = v357
	goto L78
L78:
	;
	v363 = v360
	goto L73
L79:
	;
	v388 = v210 * int32(10)
	if base.F64_lt(v382, base.F64_convert_i32_s(v388)) != 0 {
		v400 = v388
		goto L39
	} else {
		goto L80
	}
L80:
	;
	v391 = base.F64_ceil(v382)
	if base.F64_lt(base.F64_abs(v391), float64(2.147483648e+09)) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v395 = base.I32_trunc_f64_s(v391)
	v400 = v395
	goto L39
L82:
	;
	goto L83
L83:
	;
	v400 = int32(-2147483648)
	goto L39
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+8)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v410)+4)) = v400
	v416 = F_palloc(m, v400<<(uint(int32(4))%32))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v416
	if int32(0) < v400 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v433 = int32(0)
	goto L89
L87:
	;
	goto L88
L88:
	;
	v513 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v513 < v515 {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v468 = F_palloc(m, v46<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L13
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416+v433<<(uint(int32(4))%32)))) = v468
	v472 = v433 + int32(1)
	if v472 != v400 {
		v433 = v472
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	F_pg_qsort(m, v746, v747, int32(16), int32(819))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L13
	} else {
		goto L121
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L13
	} else {
		goto L118
	}
L95:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v519 = v513
	v520 = v513
	goto L98
L96:
	;
	goto L97
L97:
	;
	goto L93
L98:
	;
	v559 = v519 << (uint(int32(4)) % 32)
	v560 = v518 + v559
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	if v563 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L97
L100:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v674 = F_geqo_eval(m, l0, v672, v673)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L13
	} else {
		goto L110
	}
L101:
	;
	v566 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = v566
	if v563 == v566 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v576 = int32(1)
	goto L103
L103:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v615 = F_pg_prng_uint64_range(m, v610+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v576))
	mBase = m.M
	v616 = base.I32_wrap_i64(v615)
	goto L105
L104:
	;
	goto L100
L105:
	;
	if v616 != v576 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v618 = int32(2)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v561+v616<<(uint(v618)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v561+v576<<(uint(v618)%32)))) = v624
	goto L108
L107:
	;
	goto L108
L108:
	;
	v630 = v576 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v561+v616<<(uint(int32(2))%32)))) = v630
	if v630 != v563 {
		v576 = v630
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L104
L110:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*float64)(unsafe.Add(mBase, uint32(v676+v559)+8)) = v674
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v681 = *(*float64)(unsafe.Add(mBase, uint32(v679+v559)+8))
	if base.F64_lt(v681, float64(1.7976931348623157e+308)) != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v690 < v692 {
		v519 = v690
		v520 = v691
		goto L98
	} else {
		goto L117
	}
L112:
	;
	v690 = v519 + int32(1)
	v691 = v520
	goto L111
L113:
	;
	goto L114
L114:
	;
	v687 = v520 + int32(1)
	if v519 != 0 {
		v690 = v519
		v691 = v687
		goto L111
	} else {
		goto L115
	}
L115:
	;
	if int32(10000) <= v687 {
		goto L94
	} else {
		goto L116
	}
L116:
	;
	v690 = v519
	v691 = v687
	goto L111
L117:
	;
	goto L99
L118:
	;
	F_errmsg_internal(m, int32(283970), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L13
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(498499), int32(117), int32(302255))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v753 = F_alloc_chromo(m, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v756 = F_alloc_chromo(m, v755)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v759 = int32(24)
	v763 = F_palloc(m, v758*v759+v759)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	if int32(0) < v408 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v767 = v408
	goto L127
L126:
	;
	v767 = v400
	goto L127
L127:
	;
	if int32(0) < v767 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v782 = int32(0)
	goto L131
L129:
	;
	goto L130
L130:
	;
	v3642 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v3643 = *(*int32)(unsafe.Add(mBase, uint32(v3642)))
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v3645 = F_gimme_tree(m, l0, v3643, v3644)
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L13
	} else {
		goto L412
	}
L131:
	;
	v811 = *(*float64)(unsafe.Add(mBase, _consts[633]))
	v813 = base.F64_add(v811, float64(-1))
	v815 = base.F64_mul(v813, float64(4))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v817 = base.F64_convert_i32_s(v816)
	goto L133
L132:
	;
	goto L130
L133:
	;
	v857 = base.F64_mul(v811, v811)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v861 = F_pg_prng_double(m, v858+int32(8))
	mBase = m.M
	goto L135
L134:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v877 = base.F64_convert_i32_s(v876)
	goto L141
L135:
	;
	v863 = base.F64_sub(v857, base.F64_mul(v815, v861))
	if base.F64_gt(v863, float64(0)) != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v867 = base.F64_sqrt(v863)
	goto L138
L137:
	;
	v867 = v863
	goto L138
L138:
	;
	v872 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v811, v867), v817), float64(0.5)), v813)
	if base.F64_lt(v872, float64(0)) != 0 {
		goto L133
	} else {
		goto L139
	}
L139:
	;
	if base.F64_ge(v872, v817) != 0 {
		goto L133
	} else {
		goto L140
	}
L140:
	;
	goto L134
L141:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v920 = F_pg_prng_double(m, v917+int32(8))
	mBase = m.M
	goto L143
L142:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if base.F64_lt(base.F64_abs(v931), float64(2.147483648e+09)) != 0 {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	v922 = base.F64_sub(v857, base.F64_mul(v815, v920))
	if base.F64_gt(v922, float64(0)) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v926 = base.F64_sqrt(v922)
	goto L146
L145:
	;
	v926 = v922
	goto L146
L146:
	;
	v931 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v811, v926), v877), float64(0.5)), v813)
	if base.F64_lt(v931, float64(0)) != 0 {
		goto L141
	} else {
		goto L147
	}
L147:
	;
	if base.F64_ge(v931, v877) != 0 {
		goto L141
	} else {
		goto L148
	}
L148:
	;
	goto L142
L149:
	;
	if base.F64_lt(base.F64_abs(v872), float64(2.147483648e+09)) != 0 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v939 = base.I32_trunc_f64_s(v931)
	v941 = v939
	goto L149
L151:
	;
	goto L152
L152:
	;
	v941 = int32(-2147483648)
	goto L149
L153:
	;
	if v935 < int32(2) {
		v1057 = v941
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v947 = base.I32_trunc_f64_s(v872)
	v949 = v947
	goto L153
L155:
	;
	goto L156
L156:
	;
	v949 = int32(-2147483648)
	goto L153
L157:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v1098 = v1095 + v949<<(uint(int32(4))%32)
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v1100 = int32(0)
	if v1099 <= v1100 {
		goto L176
	} else {
		goto L177
	}
L158:
	;
	if v949 != v941 {
		v1057 = v941
		goto L157
	} else {
		goto L159
	}
L159:
	;
	goto L160
L160:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v991 = base.F64_convert_i32_s(v990)
	goto L162
L161:
	;
	v1057 = v1054
	goto L157
L162:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v1034 = F_pg_prng_double(m, v1031+int32(8))
	mBase = m.M
	goto L164
L163:
	;
	if base.F64_lt(base.F64_abs(v1045), float64(2.147483648e+09)) != 0 {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	v1036 = base.F64_sub(v857, base.F64_mul(v815, v1034))
	if base.F64_gt(v1036, float64(0)) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1040 = base.F64_sqrt(v1036)
	goto L167
L166:
	;
	v1040 = v1036
	goto L167
L167:
	;
	v1045 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_sub(v811, v1040), v991), float64(0.5)), v813)
	if base.F64_lt(v1045, float64(0)) != 0 {
		goto L162
	} else {
		goto L168
	}
L168:
	;
	if base.F64_ge(v1045, v991) != 0 {
		goto L162
	} else {
		goto L169
	}
L169:
	;
	goto L163
L170:
	;
	if v1054 == v949 {
		goto L160
	} else {
		goto L174
	}
L171:
	;
	v1052 = base.I32_trunc_f64_s(v1045)
	v1054 = v1052
	goto L170
L172:
	;
	goto L173
L173:
	;
	v1054 = int32(-2147483648)
	goto L170
L174:
	;
	goto L161
L175:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v1207 = v1204 + v1057<<(uint(int32(4))%32)
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v1209 = int32(0)
	if v1208 <= v1209 {
		goto L189
	} else {
		goto L190
	}
L176:
	;
	v1202 = *(*float64)(unsafe.Add(mBase, uint32(v1098)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v753)+8)) = v1202
	goto L175
L177:
	;
	v1109 = v1099 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1099) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1117 = v1100
	v1121 = v1100
	goto L181
L179:
	;
	v1163 = v1100
	goto L180
L180:
	;
	if v1109 == int32(0) {
		goto L176
	} else {
		goto L184
	}
L181:
	;
	v1124 = v1117 << (uint(int32(2)) % 32)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1127+v1124)))
	*(*int32)(unsafe.Add(mBase, uint32(v1124+v1125))) = v1129
	v1131 = int32(4)
	v1132 = v1124 | v1131
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1135+v1132)))
	*(*int32)(unsafe.Add(mBase, uint32(v1132+v1133))) = v1137
	v1140 = v1124 | int32(8)
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1143+v1140)))
	*(*int32)(unsafe.Add(mBase, uint32(v1140+v1141))) = v1145
	v1148 = v1124 | int32(12)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1151+v1148)))
	*(*int32)(unsafe.Add(mBase, uint32(v1148+v1149))) = v1153
	v1156 = v1117 + v1131
	v1158 = v1121 + v1131
	if v1158 != v1099&int32(2147483644) {
		v1117 = v1156
		v1121 = v1158
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v1163 = v1156
	goto L180
L183:
	;
	goto L182
L184:
	;
	v1174 = v1163
	v1179 = v1100
	goto L185
L185:
	;
	v1181 = v1174 << (uint(int32(2)) % 32)
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1184+v1181)))
	*(*int32)(unsafe.Add(mBase, uint32(v1181+v1182))) = v1186
	v1188 = int32(1)
	v1191 = v1179 + v1188
	if v1191 != v1109 {
		v1174 = v1174 + v1188
		v1179 = v1191
		goto L185
	} else {
		goto L187
	}
L186:
	;
	goto L176
L187:
	;
	goto L186
L188:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v1316 = int32(1)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	if v1317 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L189:
	;
	v1311 = *(*float64)(unsafe.Add(mBase, uint32(v1207)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v756)+8)) = v1311
	goto L188
L190:
	;
	v1218 = v1208 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1208) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1226 = v1209
	v1230 = v1209
	goto L194
L192:
	;
	v1272 = v1209
	goto L193
L193:
	;
	if v1218 == int32(0) {
		goto L189
	} else {
		goto L197
	}
L194:
	;
	v1233 = v1226 << (uint(int32(2)) % 32)
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1233)))
	*(*int32)(unsafe.Add(mBase, uint32(v1233+v1234))) = v1238
	v1240 = int32(4)
	v1241 = v1233 | v1240
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1244+v1241)))
	*(*int32)(unsafe.Add(mBase, uint32(v1241+v1242))) = v1246
	v1249 = v1233 | int32(8)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1252+v1249)))
	*(*int32)(unsafe.Add(mBase, uint32(v1249+v1250))) = v1254
	v1257 = v1233 | int32(12)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1260+v1257)))
	*(*int32)(unsafe.Add(mBase, uint32(v1257+v1258))) = v1262
	v1265 = v1226 + v1240
	v1267 = v1230 + v1240
	if v1267 != v1208&int32(2147483644) {
		v1226 = v1265
		v1230 = v1267
		goto L194
	} else {
		goto L196
	}
L195:
	;
	v1272 = v1265
	goto L193
L196:
	;
	goto L195
L197:
	;
	v1283 = v1272
	v1288 = v1209
	goto L198
L198:
	;
	v1290 = v1283 << (uint(int32(2)) % 32)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1293+v1290)))
	*(*int32)(unsafe.Add(mBase, uint32(v1290+v1291))) = v1295
	v1297 = int32(1)
	v1300 = v1288 + v1297
	if v1300 != v1218 {
		v1283 = v1283 + v1297
		v1288 = v1300
		goto L198
	} else {
		goto L200
	}
L199:
	;
	goto L189
L200:
	;
	goto L199
L201:
	;
	goto L203
L202:
	;
	v1321 = int32(2)
	v1323 = v1317 + int32(1)
	if v1323 <= v1321 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v2239 = m.G0
	v2241 = v2239 - int32(32)
	m.G0 = v2241
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2249 = F_pg_prng_uint64_range(m, v2244+int32(8), base.I64_extend_i32_s(int32(1)), base.I64_extend_i32_s(v2237))
	mBase = m.M
	goto L261
L204:
	;
	v1326 = v1321
	goto L206
L205:
	;
	v1326 = v1323
	goto L206
L206:
	;
	v1328 = v1326 - int32(1)
	v1330 = v1328 & int32(3)
	if int32(5) <= v1323 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1338 = int32(0)
	v1340 = v1316
	goto L210
L208:
	;
	v1397 = v1316
	goto L209
L209:
	;
	if v1330 != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v1377 = v763 + v1340*int32(24)
	v1378 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1377)+16)) = v1378
	*(*int64)(unsafe.Add(mBase, uint32(v1377)+88)) = v1378
	*(*int64)(unsafe.Add(mBase, uint32(v1377-int32(-64)))) = v1378
	*(*int64)(unsafe.Add(mBase, uint32(v1377)+40)) = v1378
	v1388 = int32(4)
	v1389 = v1340 + v1388
	v1391 = v1338 + v1388
	if v1391 != v1328&int32(-4) {
		v1338 = v1391
		v1340 = v1389
		goto L210
	} else {
		goto L212
	}
L211:
	;
	v1397 = v1389
	goto L209
L212:
	;
	goto L211
L213:
	;
	v1433 = int32(0)
	v1437 = v1397
	goto L216
L214:
	;
	goto L215
L215:
	;
	v1522 = int32(0)
	goto L219
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v763+v1437*int32(24))+16)) = int64(0)
	v1477 = int32(1)
	v1480 = v1433 + v1477
	if v1480 != v1330 {
		v1433 = v1480
		v1437 = v1437 + v1477
		goto L216
	} else {
		goto L218
	}
L217:
	;
	goto L215
L218:
	;
	goto L217
L219:
	;
	v1562 = v1522 + int32(1)
	if v1562 != v1317 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	goto L203
L221:
	;
	v1565 = v1562
	goto L223
L222:
	;
	v1565 = int32(0)
	goto L223
L223:
	;
	v1566 = int32(2)
	v1567 = v1565 << (uint(v1566) % 32)
	v1568 = v1313 + v1567
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1568)))
	v1570 = int32(0)
	v1572 = v1522 << (uint(v1566) % 32)
	v1573 = v1313 + v1572
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1573)))
	v1577 = v763 + v1574*int32(24)
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+16))
	if v1578 <= v1570 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1728 = int32(0)
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1573)))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1568)))
	v1733 = v763 + v1730*int32(24)
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+16))
	if v1734 <= v1728 {
		goto L234
	} else {
		goto L235
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1577+v1578<<(uint(int32(2))%32)))) = v1569
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+16))
	v1680 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1577)+16)) = v1679 + v1680
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1577)+20)) = v1683 + v1680
	goto L224
L226:
	;
	v1585 = v1570
	goto L227
L227:
	;
	v1622 = v1577 + v1585<<(uint(int32(2))%32)
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1622)))
	v1625 = v1623 >> (uint(int32(31)) % 32)
	if v1569 != v1623^v1625-v1625 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622))) = int32(0) - v1569
	goto L224
L229:
	;
	v1630 = v1585 + int32(1)
	if v1578 != v1630 {
		v1585 = v1630
		goto L227
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	goto L228
L232:
	;
	goto L225
L233:
	;
	v1881 = int32(0)
	v1882 = v1567 + v1314
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1882)))
	v1884 = v1572 + v1314
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1884)))
	v1888 = v763 + v1885*int32(24)
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+16))
	if v1889 <= v1881 {
		goto L243
	} else {
		goto L244
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1733+v1734<<(uint(int32(2))%32)))) = v1729
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+16))
	v1835 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+16)) = v1834 + v1835
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+20)) = v1838 + v1835
	goto L233
L235:
	;
	v1741 = v1728
	goto L236
L236:
	;
	v1778 = v1733 + v1741<<(uint(int32(2))%32)
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1778)))
	v1781 = v1779 >> (uint(int32(31)) % 32)
	if v1729 != v1779^v1781-v1781 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1778))) = int32(0) - v1729
	goto L233
L238:
	;
	v1786 = v1741 + int32(1)
	if v1734 != v1786 {
		v1741 = v1786
		goto L236
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	goto L237
L241:
	;
	goto L234
L242:
	;
	v2040 = int32(0)
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1884)))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1882)))
	v2045 = v763 + v2042*int32(24)
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2045)+16))
	if v2046 <= v2040 {
		goto L252
	} else {
		goto L253
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1888+v1889<<(uint(int32(2))%32)))) = v1883
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+16))
	v1991 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1888)+16)) = v1990 + v1991
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1888)+20)) = v1994 + v1991
	goto L242
L244:
	;
	v1896 = v1881
	goto L245
L245:
	;
	v1933 = v1888 + v1896<<(uint(int32(2))%32)
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1933)))
	v1936 = v1934 >> (uint(int32(31)) % 32)
	if v1883 != v1934^v1936-v1936 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1933))) = int32(0) - v1883
	goto L242
L247:
	;
	v1941 = v1896 + int32(1)
	if v1889 != v1941 {
		v1896 = v1941
		goto L245
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	goto L246
L250:
	;
	goto L243
L251:
	;
	if v1562 != v1317 {
		v1522 = v1562
		goto L219
	} else {
		goto L260
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2045+v2046<<(uint(int32(2))%32)))) = v2041
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2045)+16))
	v2147 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2045)+16)) = v2146 + v2147
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v2045)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2045)+20)) = v2150 + v2147
	goto L251
L253:
	;
	v2053 = v2040
	goto L254
L254:
	;
	v2090 = v2045 + v2053<<(uint(int32(2))%32)
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2090)))
	v2093 = v2091 >> (uint(int32(31)) % 32)
	if v2041 != v2091^v2093-v2093 {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2090))) = int32(0) - v2041
	goto L251
L256:
	;
	v2098 = v2053 + int32(1)
	if v2046 != v2098 {
		v2053 = v2098
		goto L254
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	goto L255
L259:
	;
	goto L252
L260:
	;
	goto L220
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2236))) = base.I32_wrap_i64(v2249)
	if int32(2) <= v2237 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v2256 = int32(2)
	v2258 = v2237 + int32(1)
	if v2258 <= v2256 {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	goto L264
L264:
	;
	m.G0 = v2241 + int32(32)
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v3295 = F_geqo_eval(m, l0, v3293, v3294)
	mBase = m.M
	v3296 = m.ExcPending
	if v3296 != 0 {
		goto L13
	} else {
		goto L375
	}
L265:
	;
	v2261 = v2256
	goto L267
L266:
	;
	v2261 = v2258
	goto L267
L267:
	;
	v2262 = int32(1)
	v2263 = v2261 - v2262
	v2285 = v2262
	goto L268
L268:
	;
	v2310 = int32(24)
	v2311 = v2241 + v2310
	v2313 = v2285 << (uint(int32(2)) % 32)
	v2314 = v2236 - int32(4) + v2313
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2314)))
	v2318 = v763 + v2315*v2310
	v2319 = *(*int64)(unsafe.Add(mBase, uint32(v2318)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2311))) = v2319
	v2322 = v2241 + int32(16)
	v2323 = *(*int64)(unsafe.Add(mBase, uint32(v2318)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2322))) = v2323
	v2325 = *(*int64)(unsafe.Add(mBase, uint32(v2318)))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+8)) = v2325
	v2327 = int32(0)
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+28))
	if v2327 < v2328 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	goto L264
L270:
	;
	v2346 = v2327
	goto L273
L271:
	;
	v2493 = v2315
	goto L272
L272:
	;
	v2530 = v763 + v2493*int32(24)
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(v2530)+20))
	if int32(0) < v2531 {
		goto L288
	} else {
		goto L289
	}
L273:
	;
	v2370 = int32(0)
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2241+int32(8)+v2346<<(uint(int32(2))%32))))
	v2378 = v2376 >> (uint(int32(31)) % 32)
	v2383 = v763 + (v2376^v2378-v2378)*int32(24)
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2383)+20))
	if v2384 <= v2370 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2314)))
	v2493 = v2488
	goto L272
L275:
	;
	v2486 = v2346 + int32(1)
	if v2486 != v2328 {
		v2346 = v2486
		goto L273
	} else {
		goto L283
	}
L276:
	;
	v2389 = v2370
	goto L277
L277:
	;
	v2428 = v2383 + v2389<<(uint(int32(2))%32)
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2428)))
	v2431 = v2429 >> (uint(int32(31)) % 32)
	if v2315 != v2429^v2431-v2431 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v2439 = v2384 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2383)+20)) = v2439
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2383+v2439<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2428))) = v2444
	goto L275
L279:
	;
	v2436 = v2389 + int32(1)
	if v2384 != v2436 {
		v2389 = v2436
		goto L277
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	goto L278
L282:
	;
	goto L275
L283:
	;
	goto L274
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2236+v2313))) = v3203
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v2314)))
	*(*int32)(unsafe.Add(mBase, uint32(v763+v3242*int32(24))+20)) = int32(-1)
	v3249 = v2285 + int32(1)
	if v3249 != v2237 {
		v2285 = v3249
		goto L268
	} else {
		goto L374
	}
L285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L13
	} else {
		goto L371
	}
L286:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L13
	} else {
		goto L368
	}
L287:
	;
	v3129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v3134 = F_pg_prng_uint64_range(m, v3129+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(int32(-2)))
	mBase = m.M
	goto L367
L288:
	;
	v2534 = *(*int64)(unsafe.Add(mBase, uint32(v2530)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2311))) = v2534
	v2536 = *(*int64)(unsafe.Add(mBase, uint32(v2530)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2322))) = v2536
	v2538 = *(*int64)(unsafe.Add(mBase, uint32(v2530)))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+8)) = v2538
	v2540 = int32(0)
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+28))
	if v2541 <= v2540 {
		goto L287
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v2678 = int32(0)
	v2680 = int32(1)
	if base.B2i32(v2258 < int32(3)) == v2678 {
		goto L311
	} else {
		goto L312
	}
L291:
	;
	v2547 = int32(5)
	v2548 = v2540
	v2550 = int32(-1)
	goto L292
L292:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2241+int32(8)+v2548<<(uint(int32(2))%32))))
	if v2590 < int32(0) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v2610 = int32(0)
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2619 = F_pg_prng_uint64_range(m, v2614+int32(8), base.I64_extend_i32_s(v2610), base.I64_extend_i32_s(v2606-int32(1)))
	mBase = m.M
	goto L303
L294:
	;
	v3203 = int32(0) - v2590
	goto L284
L295:
	;
	goto L296
L296:
	;
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(v763+v2590*int32(24))+20))
	if v2598 < v2547 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v2608 = v2548 + int32(1)
	if v2608 != v2541 {
		v2547 = v2605
		v2548 = v2608
		v2550 = v2606
		goto L292
	} else {
		goto L302
	}
L298:
	;
	v2605 = v2598
	v2606 = int32(1)
	goto L297
L299:
	;
	goto L300
L300:
	;
	if v2550 == int32(-1) {
		goto L286
	} else {
		goto L301
	}
L301:
	;
	v2605 = v2547
	v2606 = v2550 + base.B2i32(v2598 == v2547)
	goto L297
L302:
	;
	goto L293
L303:
	;
	v2621 = v2610
	v2625 = v2606
	goto L304
L304:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2241+int32(8)+v2621<<(uint(int32(2))%32))))
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v763+v2665*int32(24))+20))
	if v2605 == v2669 {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	goto L285
L306:
	;
	v2672 = v2625 - int32(1)
	if v2672 == base.I32_wrap_i64(v2619) {
		v3203 = v2665
		goto L284
	} else {
		goto L309
	}
L307:
	;
	v2674 = v2625
	goto L308
L308:
	;
	v2676 = v2621 + int32(1)
	if v2676 != v2541 {
		v2621 = v2676
		v2625 = v2674
		goto L304
	} else {
		goto L310
	}
L309:
	;
	v2674 = v2672
	goto L308
L310:
	;
	goto L305
L311:
	;
	v2684 = v2678
	v2685 = v2678
	v2686 = v2680
	v2689 = v2678
	goto L314
L312:
	;
	v2762 = v2678
	v2763 = v2678
	v2764 = v2680
	goto L313
L313:
	;
	if v2263&v2262 == int32(0) {
		v2816 = v2762
		v2817 = v2763
		goto L323
	} else {
		goto L324
	}
L314:
	;
	if v2686 == v2493 {
		v2736 = v2684
		v2737 = v2685
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v2762 = v2754
	v2763 = v2755
	v2764 = v2758
	goto L313
L316:
	;
	v2740 = v2686 + int32(1)
	if v2740 == v2493 {
		v2754 = v2736
		v2755 = v2737
		goto L319
	} else {
		goto L320
	}
L317:
	;
	v2726 = v763 + v2686*int32(24)
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(v2726)+20))
	if v2727 == int32(-1) {
		v2736 = v2684
		v2737 = v2685
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v2726)+16))
	v2736 = v2684 + int32(1)
	v2737 = v2685 + base.B2i32(v2732 == int32(4))
	goto L316
L319:
	;
	v2757 = int32(2)
	v2758 = v2686 + v2757
	v2760 = v2689 + v2757
	if v2760 != v2263&int32(-2) {
		v2684 = v2754
		v2685 = v2755
		v2686 = v2758
		v2689 = v2760
		goto L314
	} else {
		goto L322
	}
L320:
	;
	v2744 = v763 + v2740*int32(24)
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(v2744)+20))
	if v2745 == int32(-1) {
		v2754 = v2736
		v2755 = v2737
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2744)+16))
	v2754 = v2736 + int32(1)
	v2755 = v2737 + base.B2i32(v2750 == int32(4))
	goto L319
L322:
	;
	goto L315
L323:
	;
	if v2817 != 0 {
		goto L329
	} else {
		goto L330
	}
L324:
	;
	if v2764 == v2493 {
		v2816 = v2762
		v2817 = v2763
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v2806 = v763 + v2764*int32(24)
	v2807 = *(*int32)(unsafe.Add(mBase, uint32(v2806)+20))
	if v2807 == int32(-1) {
		v2816 = v2762
		v2817 = v2763
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v2806)+16))
	v2816 = v2762 + int32(1)
	v2817 = v2763 + base.B2i32(v2812 == int32(4))
	goto L323
L327:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3117 = m.ExcPending
	if v3117 != 0 {
		goto L13
	} else {
		goto L364
	}
L328:
	;
	F_errmsg_internal(m, v3067, int32(0))
	mBase = m.M
	v3070 = m.ExcPending
	if v3070 != 0 {
		goto L13
	} else {
		goto L362
	}
L329:
	;
	v2821 = int32(1)
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2830 = F_pg_prng_uint64_range(m, v2825+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v2817-v2821))
	mBase = m.M
	goto L332
L330:
	;
	goto L331
L331:
	;
	if v2816 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L332:
	;
	v2833 = v2817
	v2834 = v2821
	goto L333
L333:
	;
	if v2834 == v2493 {
		v2885 = v2833
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v2891 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2892 = m.ExcPending
	if v2892 != 0 {
		goto L13
	} else {
		goto L341
	}
L335:
	;
	v2887 = v2834 + int32(1)
	if v2887 <= v2237 {
		v2833 = v2885
		v2834 = v2887
		goto L333
	} else {
		goto L340
	}
L336:
	;
	v2874 = v763 + v2834*int32(24)
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v2874)+20))
	if v2875 == int32(-1) {
		v2885 = v2833
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v2878 = *(*int32)(unsafe.Add(mBase, uint32(v2874)+16))
	if v2878 != int32(4) {
		v2885 = v2833
		goto L335
	} else {
		goto L338
	}
L338:
	;
	v2882 = v2833 - int32(1)
	if base.I32_wrap_i64(v2830) == v2882 {
		v3203 = v2834
		goto L284
	} else {
		goto L339
	}
L339:
	;
	v2885 = v2882
	goto L335
L340:
	;
	goto L334
L341:
	;
	if v2891 == int32(0) {
		goto L327
	} else {
		goto L342
	}
L342:
	;
	v3030 = int32(422)
	v3067 = int32(561014)
	goto L328
L343:
	;
	v2902 = int32(1)
	goto L346
L344:
	;
	goto L345
L345:
	;
	v2956 = int32(1)
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v2965 = F_pg_prng_uint64_range(m, v2960+int32(8), base.I64_extend_i32_s(int32(0)), base.I64_extend_i32_s(v2816-v2956))
	mBase = m.M
	goto L352
L346:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v763+v2902*int32(24))+20))
	if int32(0) <= v2942 {
		v3203 = v2902
		goto L284
	} else {
		goto L348
	}
L347:
	;
	v2950 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2951 = m.ExcPending
	if v2951 != 0 {
		goto L13
	} else {
		goto L350
	}
L348:
	;
	v2946 = v2902 + int32(1)
	if v2946 <= v2237 {
		v2902 = v2946
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	if v2950 == int32(0) {
		goto L327
	} else {
		goto L351
	}
L351:
	;
	v3030 = int32(461)
	v3067 = int32(89263)
	goto L328
L352:
	;
	v2967 = v2816
	v2969 = v2956
	goto L353
L353:
	;
	if v2969 == v2493 {
		v3016 = v2967
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v3022 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L13
	} else {
		goto L360
	}
L355:
	;
	v3018 = v2969 + int32(1)
	if v3018 <= v2237 {
		v2967 = v3016
		v2969 = v3018
		goto L353
	} else {
		goto L359
	}
L356:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(v763+v2969*int32(24))+20))
	if v3010 == int32(-1) {
		v3016 = v2967
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v3014 = v2967 - int32(1)
	if base.I32_wrap_i64(v2965) == v3014 {
		v3203 = v2969
		goto L284
	} else {
		goto L358
	}
L358:
	;
	v3016 = v3014
	goto L355
L359:
	;
	goto L354
L360:
	;
	if v3022 == int32(0) {
		goto L327
	} else {
		goto L361
	}
L361:
	;
	v3030 = int32(443)
	v3067 = int32(170422)
	goto L328
L362:
	;
	F_errfinish(m, int32(493569), v3030, int32(364175))
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L13
	} else {
		goto L363
	}
L363:
	;
	goto L327
L364:
	;
	F_errmsg_internal(m, int32(425142), int32(0))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L13
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(493569), int32(466), int32(364175))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L13
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	goto L285
L368:
	;
	F_errmsg_internal(m, int32(106619), int32(0))
	mBase = m.M
	v3143 = m.ExcPending
	if v3143 != 0 {
		goto L13
	} else {
		goto L369
	}
L369:
	;
	F_errfinish(m, int32(493569), int32(337), int32(374904))
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L13
	} else {
		goto L370
	}
L370:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L371:
	;
	F_errmsg_internal(m, int32(425156), int32(0))
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L13
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(493569), int32(362), int32(374904))
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L13
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	goto L269
L375:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v753)+8)) = v3295
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v3303 = v3301 - int32(1)
	v3307 = *(*float64)(unsafe.Add(mBase, uint32(v3300+v3303<<(uint(int32(4))%32))+8))
	if base.F64_gt(v3295, v3307) != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v3601 = v782 + int32(1)
	if v3601 != v767 {
		v782 = v3601
		goto L131
	} else {
		goto L411
	}
L377:
	;
	v3310 = base.I32_div_s(v3301, int32(2))
	v3313 = int32(0)
	v3315 = v3310
	v3316 = v3303
	goto L378
L378:
	;
	v3353 = *(*float64)(unsafe.Add(mBase, uint32(v3300+v3313<<(uint(int32(4))%32))+8))
	if base.F64_ge(v3353, v3295) != 0 {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	v3390 = v3300 + v3301<<(uint(int32(4))%32) - int32(16)
	v3391 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v3392 = int32(0)
	if v3391 <= v3392 {
		goto L395
	} else {
		goto L396
	}
L380:
	;
	if v3379 == int32(-1) {
		v3313 = v3380
		v3315 = v3383
		v3316 = v3381
		goto L378
	} else {
		goto L393
	}
L381:
	;
	v3379 = v3313
	v3380 = v3313
	v3381 = v3316
	v3383 = v3315
	goto L380
L382:
	;
	goto L383
L383:
	;
	v3358 = *(*float64)(unsafe.Add(mBase, uint32(v3300+v3315<<(uint(int32(4))%32))+8))
	if base.F64_eq(v3358, v3295) != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v3379 = v3315
	v3380 = v3313
	v3381 = v3316
	v3383 = v3315
	goto L380
L385:
	;
	goto L386
L386:
	;
	v3363 = *(*float64)(unsafe.Add(mBase, uint32(v3300+v3316<<(uint(int32(4))%32))+8))
	if base.F64_eq(v3295, v3363) != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v3379 = v3316
	v3380 = v3313
	v3381 = v3316
	v3383 = v3315
	goto L380
L388:
	;
	if v3316-v3313 < int32(2) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	if base.F64_lt(v3295, v3358) != 0 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v3372 = base.I32_div_s(v3315-v3313, int32(2))
	v3379 = int32(-1)
	v3380 = v3313
	v3381 = v3315
	v3383 = v3372 + v3313
	goto L380
L391:
	;
	goto L392
L392:
	;
	v3377 = base.I32_div_s(v3316-v3315, int32(2))
	v3379 = int32(-1)
	v3380 = v3315
	v3381 = v3316
	v3383 = v3315 + v3377
	goto L380
L393:
	;
	goto L379
L394:
	;
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v3496 <= v3379 {
		goto L376
	} else {
		goto L407
	}
L395:
	;
	v3494 = *(*float64)(unsafe.Add(mBase, uint32(v753)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v3390)+8)) = v3494
	goto L394
L396:
	;
	v3401 = v3391 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3391) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v3409 = v3392
	v3413 = v3392
	goto L400
L398:
	;
	v3455 = v3392
	goto L399
L399:
	;
	if v3401 == int32(0) {
		goto L395
	} else {
		goto L403
	}
L400:
	;
	v3416 = v3409 << (uint(int32(2)) % 32)
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(v3390)))
	v3419 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v3421 = *(*int32)(unsafe.Add(mBase, uint32(v3419+v3416)))
	*(*int32)(unsafe.Add(mBase, uint32(v3416+v3417))) = v3421
	v3423 = int32(4)
	v3424 = v3416 | v3423
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v3390)))
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v3427+v3424)))
	*(*int32)(unsafe.Add(mBase, uint32(v3424+v3425))) = v3429
	v3432 = v3416 | int32(8)
	v3433 = *(*int32)(unsafe.Add(mBase, uint32(v3390)))
	v3435 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(v3435+v3432)))
	*(*int32)(unsafe.Add(mBase, uint32(v3432+v3433))) = v3437
	v3440 = v3416 | int32(12)
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3390)))
	v3443 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3443+v3440)))
	*(*int32)(unsafe.Add(mBase, uint32(v3440+v3441))) = v3445
	v3448 = v3409 + v3423
	v3450 = v3413 + v3423
	if v3450 != v3391&int32(2147483644) {
		v3409 = v3448
		v3413 = v3450
		goto L400
	} else {
		goto L402
	}
L401:
	;
	v3455 = v3448
	goto L399
L402:
	;
	goto L401
L403:
	;
	v3466 = v3455
	v3471 = v3392
	goto L404
L404:
	;
	v3473 = v3466 << (uint(int32(2)) % 32)
	v3474 = *(*int32)(unsafe.Add(mBase, uint32(v3390)))
	v3476 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v3476+v3473)))
	*(*int32)(unsafe.Add(mBase, uint32(v3473+v3474))) = v3478
	v3480 = int32(1)
	v3483 = v3471 + v3480
	if v3483 != v3401 {
		v3466 = v3466 + v3480
		v3471 = v3483
		goto L404
	} else {
		goto L406
	}
L405:
	;
	goto L395
L406:
	;
	goto L405
L407:
	;
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v3501 = v3498 + v3496<<(uint(int32(4))%32)
	v3504 = *(*int32)(unsafe.Add(mBase, uint32(v3501-int32(16))))
	v3507 = *(*float64)(unsafe.Add(mBase, uint32(v3501-int32(8))))
	v3509 = v3379
	v3512 = v3504
	v3538 = v3507
	goto L408
L408:
	;
	v3548 = v3509 << (uint(int32(4)) % 32)
	v3549 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v3550 = v3548 + v3549
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3550)))
	*(*int32)(unsafe.Add(mBase, uint32(v3550))) = v3512
	v3553 = *(*float64)(unsafe.Add(mBase, uint32(v3550)+8))
	v3554 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*float64)(unsafe.Add(mBase, uint32(v3554+v3548)+8)) = v3538
	v3558 = v3509 + int32(1)
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v3558 < v3559 {
		v3509 = v3558
		v3512 = v3551
		v3538 = v3553
		goto L408
	} else {
		goto L410
	}
L409:
	;
	goto L376
L410:
	;
	goto L409
L411:
	;
	goto L132
L412:
	;
	if v3645 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3652 = m.ExcPending
	if v3652 != 0 {
		goto L13
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	F_free_attrmap(m, v753)
	mBase = m.M
	v3663 = m.ExcPending
	if v3663 != 0 {
		goto L13
	} else {
		goto L419
	}
L416:
	;
	F_errmsg_internal(m, int32(283970), int32(0))
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L13
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(497814), int32(275), int32(240924))
	mBase = m.M
	v3661 = m.ExcPending
	if v3661 != 0 {
		goto L13
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	F_free_attrmap(m, v756)
	mBase = m.M
	v3665 = m.ExcPending
	if v3665 != 0 {
		goto L13
	} else {
		goto L420
	}
L420:
	;
	F_pfree(m, v763)
	mBase = m.M
	v3667 = m.ExcPending
	if v3667 != 0 {
		goto L13
	} else {
		goto L421
	}
L421:
	;
	v3668 = int32(0)
	v3669 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v3670 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v3668 < v3670 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v3680 = v3668
	goto L425
L423:
	;
	v3762 = v3669
	goto L424
L424:
	;
	F_pfree(m, v3762)
	mBase = m.M
	v3764 = m.ExcPending
	if v3764 != 0 {
		goto L13
	} else {
		goto L429
	}
L425:
	;
	v3715 = *(*int32)(unsafe.Add(mBase, uint32(v3669+v3680<<(uint(int32(4))%32))))
	F_pfree(m, v3715)
	mBase = m.M
	v3717 = m.ExcPending
	if v3717 != 0 {
		goto L13
	} else {
		goto L427
	}
L426:
	;
	v3722 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v3762 = v3722
	goto L424
L427:
	;
	v3719 = v3680 + int32(1)
	v3720 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v3719 < v3720 {
		v3680 = v3719
		goto L425
	} else {
		goto L428
	}
L428:
	;
	goto L426
L429:
	;
	F_pfree(m, v410)
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L13
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = int32(0)
	m.G0 = v144 + int32(32)
	v5300 = v3645
	goto L1
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v3780)+4)) = v120
	if int32(2) <= v46 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v3789 = int32(2)
	goto L435
L433:
	;
	goto L434
L434:
	;
	v5271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v5275 = *(*int32)(unsafe.Add(mBase, uint32(v5271+v46<<(uint(int32(2))%32))))
	if v5275 == int32(0) {
		goto L671
	} else {
		goto L672
	}
L435:
	;
	v3826 = int32(0)
	v3827 = m.G0
	v3829 = v3827 - int32(16)
	m.G0 = v3829
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v3789
	v3832 = int32(2)
	v3833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v3836 = v3833 + v3789<<(uint(v3832)%32)
	v3838 = v3836 - int32(4)
	v3839 = *(*int32)(unsafe.Add(mBase, uint32(v3838)))
	if v3839 == v3826 {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	goto L434
L437:
	;
	v4294 = int32(2)
	v4295 = v3789 - v4294
	if v4294 <= v4295 {
		goto L530
	} else {
		goto L531
	}
L438:
	;
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v3839)+4))
	if v3842 <= int32(0) {
		goto L437
	} else {
		goto L439
	}
L439:
	;
	v3846 = v3833 + int32(4)
	v3852 = v3826
	goto L440
L440:
	;
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3839)+12))
	v3890 = *(*int32)(unsafe.Add(mBase, uint32(v3886+v3852<<(uint(int32(2))%32))))
	v3891 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+212))
	if v3891 != 0 {
		goto L444
	} else {
		goto L445
	}
L441:
	;
	goto L437
L442:
	;
	v4252 = v3852 + int32(1)
	v4253 = *(*int32)(unsafe.Add(mBase, uint32(v3839)+4))
	if v4252 < v4253 {
		v3852 = v4252
		goto L440
	} else {
		goto L529
	}
L443:
	;
	v4106 = *(*int32)(unsafe.Add(mBase, uint32(v3846)))
	if v4106 == int32(0) {
		goto L442
	} else {
		goto L506
	}
L444:
	;
	v3990 = *(*int32)(unsafe.Add(mBase, uint32(v3846)))
	if v3990 == int32(0) {
		goto L442
	} else {
		goto L475
	}
L445:
	;
	v3892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3890)+216)))
	if v3892 != 0 {
		goto L444
	} else {
		goto L446
	}
L446:
	;
	v3896 = int32(1)
	v3897 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+64))
	if v3897 != 0 {
		v3979 = v3896
		goto L448
	} else {
		goto L449
	}
L447:
	;
	if v3987 == int32(0) {
		goto L443
	} else {
		goto L474
	}
L448:
	;
	v3987 = v3979
	goto L447
L449:
	;
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+104))
	if v3898 != 0 {
		v3979 = v3896
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v3899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v3899 == int32(0) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v3933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3933 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L452:
	;
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v3899)+4))
	if v3902 <= int32(0) {
		goto L451
	} else {
		goto L453
	}
L453:
	;
	v3908 = int32(0)
	goto L454
L454:
	;
	v3910 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v3911 = *(*int32)(unsafe.Add(mBase, uint32(v3899)+12))
	v3915 = *(*int32)(unsafe.Add(mBase, uint32(v3911+v3908<<(uint(int32(2))%32))))
	v3916 = *(*int32)(unsafe.Add(mBase, uint32(v3915)+12))
	v3917 = F_bms_is_subset(m, v3910, v3916)
	mBase = m.M
	if v3917 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L455:
	;
	goto L451
L456:
	;
	v3925 = v3908 + int32(1)
	v3926 = *(*int32)(unsafe.Add(mBase, uint32(v3899)+4))
	if v3925 < v3926 {
		v3908 = v3925
		goto L454
	} else {
		goto L459
	}
L457:
	;
	v3920 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v3921 = *(*int32)(unsafe.Add(mBase, uint32(v3915)+12))
	v3922 = F_bms_equal(m, v3920, v3921)
	mBase = m.M
	if v3922 != 0 {
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v3987 = int32(1)
	goto L447
L459:
	;
	goto L455
L460:
	;
	v3979 = int32(0)
	goto L448
L461:
	;
	v3936 = *(*int32)(unsafe.Add(mBase, uint32(v3933)+4))
	if v3936 <= int32(0) {
		goto L460
	} else {
		goto L462
	}
L462:
	;
	v3943 = int32(0)
	goto L463
L463:
	;
	v3945 = *(*int32)(unsafe.Add(mBase, uint32(v3933)+12))
	v3946 = int32(2)
	v3949 = *(*int32)(unsafe.Add(mBase, uint32(v3945+v3943<<(uint(v3946)%32))))
	v3950 = *(*int32)(unsafe.Add(mBase, uint32(v3949)+20))
	if v3950 == v3946 {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	goto L460
L465:
	;
	v3968 = v3943 + int32(1)
	v3969 = *(*int32)(unsafe.Add(mBase, uint32(v3933)+4))
	if v3968 < v3969 {
		v3943 = v3968
		goto L463
	} else {
		goto L473
	}
L466:
	;
	v3953 = *(*int32)(unsafe.Add(mBase, uint32(v3949)+4))
	v3954 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v3955 = F_bms_is_subset(m, v3953, v3954)
	mBase = m.M
	if v3955 != 0 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v3949)+8))
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v3958 = F_bms_is_subset(m, v3956, v3957)
	mBase = m.M
	if v3958 != 0 {
		goto L465
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	v3959 = int32(1)
	v3960 = *(*int32)(unsafe.Add(mBase, uint32(v3949)+4))
	v3961 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v3962 = F_bms_overlap(m, v3960, v3961)
	mBase = m.M
	if v3962 != 0 {
		v3979 = v3959
		goto L448
	} else {
		goto L471
	}
L470:
	;
	goto L469
L471:
	;
	v3963 = *(*int32)(unsafe.Add(mBase, uint32(v3949)+8))
	v3964 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v3965 = F_bms_overlap(m, v3963, v3964)
	mBase = m.M
	if v3965 != 0 {
		v3979 = v3959
		goto L448
	} else {
		goto L472
	}
L472:
	;
	goto L465
L473:
	;
	goto L464
L474:
	;
	goto L444
L475:
	;
	if v3789 == int32(2) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v3998 = v3852 + int32(1)
	goto L478
L477:
	;
	v3998 = int32(0)
	goto L478
L478:
	;
	v3999 = *(*int32)(unsafe.Add(mBase, uint32(v3990)+4))
	if v3999 <= v3998 {
		goto L442
	} else {
		goto L479
	}
L479:
	;
	v4004 = v3998
	goto L480
L480:
	;
	v4040 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v4041 = *(*int32)(unsafe.Add(mBase, uint32(v3990)+12))
	v4045 = *(*int32)(unsafe.Add(mBase, uint32(v4041+v4004<<(uint(int32(2))%32))))
	v4046 = *(*int32)(unsafe.Add(mBase, uint32(v4045)+8))
	v4047 = int32(0)
	if v4040 == v4047 {
		v4088 = v4047
		goto L484
	} else {
		goto L485
	}
L481:
	;
	goto L442
L482:
	;
	v4103 = v4004 + int32(1)
	v4104 = *(*int32)(unsafe.Add(mBase, uint32(v3990)+4))
	if v4103 < v4104 {
		v4004 = v4103
		goto L480
	} else {
		goto L505
	}
L483:
	;
	if v4088 != 0 {
		goto L482
	} else {
		goto L497
	}
L484:
	;
	goto L483
L485:
	;
	if v4046 == int32(0) {
		v4088 = v4047
		goto L484
	} else {
		goto L486
	}
L486:
	;
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(v4040)+4))
	v4057 = *(*int32)(unsafe.Add(mBase, uint32(v4046)+4))
	if v4056 < v4057 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v4059 = v4056
	goto L489
L488:
	;
	v4059 = v4057
	goto L489
L489:
	;
	if v4059 <= int32(1) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v4062 = int32(1)
	goto L492
L491:
	;
	v4062 = v4059
	goto L492
L492:
	;
	v4063 = int32(8)
	v4068 = int32(0)
	goto L493
L493:
	;
	v4075 = v4068 << (uint(int32(2)) % 32)
	v4077 = *(*int32)(unsafe.Add(mBase, uint32(v4046+v4063+v4075)))
	v4079 = *(*int32)(unsafe.Add(mBase, uint32(v4075+(v4040+v4063))))
	v4080 = v4077 & v4079
	v4082 = base.B2i32(v4080 != int32(0))
	if v4080 != 0 {
		v4088 = v4082
		goto L484
	} else {
		goto L495
	}
L494:
	;
	v4088 = v4082
	goto L484
L495:
	;
	v4084 = v4068 + int32(1)
	if v4084 != v4062 {
		v4068 = v4084
		goto L493
	} else {
		goto L496
	}
L496:
	;
	goto L494
L497:
	;
	v4092 = F_have_relevant_joinclause(m, l0, v3890, v4045)
	mBase = m.M
	v4093 = m.ExcPending
	if v4093 != 0 {
		goto L13
	} else {
		goto L498
	}
L498:
	;
	if v4092 == int32(0) {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v4096 = F_have_join_order_restriction(m, l0, v3890, v4045)
	mBase = m.M
	v4097 = m.ExcPending
	if v4097 != 0 {
		goto L13
	} else {
		goto L502
	}
L500:
	;
	goto L501
L501:
	;
	v4100 = F_make_join_rel(m, l0, v3890, v4045)
	mBase = m.M
	v4101 = m.ExcPending
	if v4101 != 0 {
		goto L13
	} else {
		goto L504
	}
L502:
	;
	if v4096 == int32(0) {
		goto L482
	} else {
		goto L503
	}
L503:
	;
	goto L501
L504:
	;
	goto L482
L505:
	;
	goto L481
L506:
	;
	v4109 = int32(0)
	v4110 = *(*int32)(unsafe.Add(mBase, uint32(v4106)+4))
	if v4110 <= v4109 {
		goto L442
	} else {
		goto L507
	}
L507:
	;
	v4116 = v4109
	goto L508
L508:
	;
	v4152 = *(*int32)(unsafe.Add(mBase, uint32(v4106)+12))
	v4156 = *(*int32)(unsafe.Add(mBase, uint32(v4152+v4116<<(uint(int32(2))%32))))
	v4157 = *(*int32)(unsafe.Add(mBase, uint32(v4156)+8))
	v4158 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+8))
	v4159 = int32(0)
	if v4157 == v4159 {
		v4200 = v4159
		goto L511
	} else {
		goto L512
	}
L509:
	;
	goto L442
L510:
	;
	if v4200 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L511:
	;
	goto L510
L512:
	;
	if v4158 == int32(0) {
		v4200 = v4159
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v4168 = *(*int32)(unsafe.Add(mBase, uint32(v4157)+4))
	v4169 = *(*int32)(unsafe.Add(mBase, uint32(v4158)+4))
	if v4168 < v4169 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v4171 = v4168
	goto L516
L515:
	;
	v4171 = v4169
	goto L516
L516:
	;
	if v4171 <= int32(1) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v4174 = int32(1)
	goto L519
L518:
	;
	v4174 = v4171
	goto L519
L519:
	;
	v4175 = int32(8)
	v4180 = int32(0)
	goto L520
L520:
	;
	v4187 = v4180 << (uint(int32(2)) % 32)
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4158+v4175+v4187)))
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v4187+(v4157+v4175))))
	v4192 = v4189 & v4191
	v4194 = base.B2i32(v4192 != int32(0))
	if v4192 != 0 {
		v4200 = v4194
		goto L511
	} else {
		goto L522
	}
L521:
	;
	v4200 = v4194
	goto L511
L522:
	;
	v4196 = v4180 + int32(1)
	if v4196 != v4174 {
		v4180 = v4196
		goto L520
	} else {
		goto L523
	}
L523:
	;
	goto L521
L524:
	;
	v4206 = F_make_join_rel(m, l0, v3890, v4156)
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L13
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v4209 = v4116 + int32(1)
	v4210 = *(*int32)(unsafe.Add(mBase, uint32(v4106)+4))
	if v4209 < v4210 {
		v4116 = v4209
		goto L508
	} else {
		goto L528
	}
L527:
	;
	goto L526
L528:
	;
	goto L509
L529:
	;
	goto L441
L530:
	;
	v4307 = v3832
	v4312 = v4295
	goto L533
L531:
	;
	goto L532
L532:
	;
	v4733 = *(*int32)(unsafe.Add(mBase, uint32(v3836)))
	if v4733 != 0 {
		goto L605
	} else {
		goto L606
	}
L533:
	;
	v4340 = *(*int32)(unsafe.Add(mBase, uint32(v3833+v4307<<(uint(int32(2))%32))))
	if v4340 == int32(0) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	goto L532
L535:
	;
	v4691 = v4307 + int32(1)
	v4692 = v3789 - v4691
	if v4691 <= v4692 {
		v4307 = v4691
		v4312 = v4692
		goto L533
	} else {
		goto L604
	}
L536:
	;
	v4343 = *(*int32)(unsafe.Add(mBase, uint32(v4340)+4))
	if v4343 <= int32(0) {
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v4355 = int32(0)
	goto L538
L538:
	;
	v4389 = *(*int32)(unsafe.Add(mBase, uint32(v4340)+12))
	v4393 = *(*int32)(unsafe.Add(mBase, uint32(v4389+v4355<<(uint(int32(2))%32))))
	v4394 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+212))
	if v4394 != 0 {
		goto L541
	} else {
		goto L542
	}
L539:
	;
	goto L535
L540:
	;
	v4648 = v4355 + int32(1)
	v4649 = *(*int32)(unsafe.Add(mBase, uint32(v4340)+4))
	if v4648 < v4649 {
		v4355 = v4648
		goto L538
	} else {
		goto L603
	}
L541:
	;
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v3833+v4312<<(uint(int32(2))%32))))
	if v4493 == int32(0) {
		goto L540
	} else {
		goto L572
	}
L542:
	;
	v4395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4393)+216)))
	if v4395 != 0 {
		goto L541
	} else {
		goto L543
	}
L543:
	;
	v4399 = int32(1)
	v4400 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+64))
	if v4400 != 0 {
		v4482 = v4399
		goto L545
	} else {
		goto L546
	}
L544:
	;
	if v4490 == int32(0) {
		goto L540
	} else {
		goto L571
	}
L545:
	;
	v4490 = v4482
	goto L544
L546:
	;
	v4401 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+104))
	if v4401 != 0 {
		v4482 = v4399
		goto L545
	} else {
		goto L547
	}
L547:
	;
	v4402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v4402 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v4436 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L549:
	;
	v4405 = *(*int32)(unsafe.Add(mBase, uint32(v4402)+4))
	if v4405 <= int32(0) {
		goto L548
	} else {
		goto L550
	}
L550:
	;
	v4411 = int32(0)
	goto L551
L551:
	;
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4414 = *(*int32)(unsafe.Add(mBase, uint32(v4402)+12))
	v4418 = *(*int32)(unsafe.Add(mBase, uint32(v4414+v4411<<(uint(int32(2))%32))))
	v4419 = *(*int32)(unsafe.Add(mBase, uint32(v4418)+12))
	v4420 = F_bms_is_subset(m, v4413, v4419)
	mBase = m.M
	if v4420 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L552:
	;
	goto L548
L553:
	;
	v4428 = v4411 + int32(1)
	v4429 = *(*int32)(unsafe.Add(mBase, uint32(v4402)+4))
	if v4428 < v4429 {
		v4411 = v4428
		goto L551
	} else {
		goto L556
	}
L554:
	;
	v4423 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4424 = *(*int32)(unsafe.Add(mBase, uint32(v4418)+12))
	v4425 = F_bms_equal(m, v4423, v4424)
	mBase = m.M
	if v4425 != 0 {
		goto L553
	} else {
		goto L555
	}
L555:
	;
	v4490 = int32(1)
	goto L544
L556:
	;
	goto L552
L557:
	;
	v4482 = int32(0)
	goto L545
L558:
	;
	v4439 = *(*int32)(unsafe.Add(mBase, uint32(v4436)+4))
	if v4439 <= int32(0) {
		goto L557
	} else {
		goto L559
	}
L559:
	;
	v4446 = int32(0)
	goto L560
L560:
	;
	v4448 = *(*int32)(unsafe.Add(mBase, uint32(v4436)+12))
	v4449 = int32(2)
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v4448+v4446<<(uint(v4449)%32))))
	v4453 = *(*int32)(unsafe.Add(mBase, uint32(v4452)+20))
	if v4453 == v4449 {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	goto L557
L562:
	;
	v4471 = v4446 + int32(1)
	v4472 = *(*int32)(unsafe.Add(mBase, uint32(v4436)+4))
	if v4471 < v4472 {
		v4446 = v4471
		goto L560
	} else {
		goto L570
	}
L563:
	;
	v4456 = *(*int32)(unsafe.Add(mBase, uint32(v4452)+4))
	v4457 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4458 = F_bms_is_subset(m, v4456, v4457)
	mBase = m.M
	if v4458 != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v4459 = *(*int32)(unsafe.Add(mBase, uint32(v4452)+8))
	v4460 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4461 = F_bms_is_subset(m, v4459, v4460)
	mBase = m.M
	if v4461 != 0 {
		goto L562
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	v4462 = int32(1)
	v4463 = *(*int32)(unsafe.Add(mBase, uint32(v4452)+4))
	v4464 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4465 = F_bms_overlap(m, v4463, v4464)
	mBase = m.M
	if v4465 != 0 {
		v4482 = v4462
		goto L545
	} else {
		goto L568
	}
L567:
	;
	goto L566
L568:
	;
	v4466 = *(*int32)(unsafe.Add(mBase, uint32(v4452)+8))
	v4467 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4468 = F_bms_overlap(m, v4466, v4467)
	mBase = m.M
	if v4468 != 0 {
		v4482 = v4462
		goto L545
	} else {
		goto L569
	}
L569:
	;
	goto L562
L570:
	;
	goto L561
L571:
	;
	goto L541
L572:
	;
	if v4307 == v4312 {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v4500 = v4355 + int32(1)
	goto L575
L574:
	;
	v4500 = int32(0)
	goto L575
L575:
	;
	v4501 = *(*int32)(unsafe.Add(mBase, uint32(v4493)+4))
	if v4501 <= v4500 {
		goto L540
	} else {
		goto L576
	}
L576:
	;
	v4506 = v4500
	goto L577
L577:
	;
	v4542 = *(*int32)(unsafe.Add(mBase, uint32(v4393)+8))
	v4543 = *(*int32)(unsafe.Add(mBase, uint32(v4493)+12))
	v4547 = *(*int32)(unsafe.Add(mBase, uint32(v4543+v4506<<(uint(int32(2))%32))))
	v4548 = *(*int32)(unsafe.Add(mBase, uint32(v4547)+8))
	v4549 = int32(0)
	if v4542 == v4549 {
		v4590 = v4549
		goto L581
	} else {
		goto L582
	}
L578:
	;
	goto L540
L579:
	;
	v4605 = v4506 + int32(1)
	v4606 = *(*int32)(unsafe.Add(mBase, uint32(v4493)+4))
	if v4605 < v4606 {
		v4506 = v4605
		goto L577
	} else {
		goto L602
	}
L580:
	;
	if v4590 != 0 {
		goto L579
	} else {
		goto L594
	}
L581:
	;
	goto L580
L582:
	;
	if v4548 == int32(0) {
		v4590 = v4549
		goto L581
	} else {
		goto L583
	}
L583:
	;
	v4558 = *(*int32)(unsafe.Add(mBase, uint32(v4542)+4))
	v4559 = *(*int32)(unsafe.Add(mBase, uint32(v4548)+4))
	if v4558 < v4559 {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v4561 = v4558
	goto L586
L585:
	;
	v4561 = v4559
	goto L586
L586:
	;
	if v4561 <= int32(1) {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v4564 = int32(1)
	goto L589
L588:
	;
	v4564 = v4561
	goto L589
L589:
	;
	v4565 = int32(8)
	v4570 = int32(0)
	goto L590
L590:
	;
	v4577 = v4570 << (uint(int32(2)) % 32)
	v4579 = *(*int32)(unsafe.Add(mBase, uint32(v4548+v4565+v4577)))
	v4581 = *(*int32)(unsafe.Add(mBase, uint32(v4577+(v4542+v4565))))
	v4582 = v4579 & v4581
	v4584 = base.B2i32(v4582 != int32(0))
	if v4582 != 0 {
		v4590 = v4584
		goto L581
	} else {
		goto L592
	}
L591:
	;
	v4590 = v4584
	goto L581
L592:
	;
	v4586 = v4570 + int32(1)
	if v4586 != v4564 {
		v4570 = v4586
		goto L590
	} else {
		goto L593
	}
L593:
	;
	goto L591
L594:
	;
	v4594 = F_have_relevant_joinclause(m, l0, v4393, v4547)
	mBase = m.M
	v4595 = m.ExcPending
	if v4595 != 0 {
		goto L13
	} else {
		goto L595
	}
L595:
	;
	if v4594 == int32(0) {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v4598 = F_have_join_order_restriction(m, l0, v4393, v4547)
	mBase = m.M
	v4599 = m.ExcPending
	if v4599 != 0 {
		goto L13
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	v4602 = F_make_join_rel(m, l0, v4393, v4547)
	mBase = m.M
	v4603 = m.ExcPending
	if v4603 != 0 {
		goto L13
	} else {
		goto L601
	}
L599:
	;
	if v4598 == int32(0) {
		goto L579
	} else {
		goto L600
	}
L600:
	;
	goto L598
L601:
	;
	goto L579
L602:
	;
	goto L578
L603:
	;
	goto L539
L604:
	;
	goto L534
L605:
	;
	m.G0 = v3829 + int32(16)
	v5070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v5074 = *(*int32)(unsafe.Add(mBase, uint32(v5070+v3789<<(uint(int32(2))%32))))
	if v5074 == int32(0) {
		goto L646
	} else {
		goto L647
	}
L606:
	;
	v4734 = *(*int32)(unsafe.Add(mBase, uint32(v3838)))
	if v4734 != 0 {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v4735 = int32(0)
	v4736 = *(*int32)(unsafe.Add(mBase, uint32(v4734)+4))
	if v4735 < v4736 {
		goto L610
	} else {
		goto L611
	}
L608:
	;
	goto L609
L609:
	;
	v5013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v5013 != 0 {
		goto L605
	} else {
		goto L641
	}
L610:
	;
	v4746 = v4735
	goto L613
L611:
	;
	goto L612
L612:
	;
	v4973 = *(*int32)(unsafe.Add(mBase, uint32(v3836)))
	if v4973 != 0 {
		goto L605
	} else {
		goto L640
	}
L613:
	;
	v4780 = *(*int32)(unsafe.Add(mBase, uint32(v3833+int32(4))))
	if v4780 == int32(0) {
		goto L615
	} else {
		goto L616
	}
L614:
	;
	goto L612
L615:
	;
	v4931 = v4746 + int32(1)
	v4932 = *(*int32)(unsafe.Add(mBase, uint32(v4734)+4))
	if v4931 < v4932 {
		v4746 = v4931
		goto L613
	} else {
		goto L639
	}
L616:
	;
	v4783 = *(*int32)(unsafe.Add(mBase, uint32(v4780)+4))
	if v4783 <= int32(0) {
		goto L615
	} else {
		goto L617
	}
L617:
	;
	v4786 = *(*int32)(unsafe.Add(mBase, uint32(v4734)+12))
	v4790 = *(*int32)(unsafe.Add(mBase, uint32(v4786+v4746<<(uint(int32(2))%32))))
	v4795 = int32(0)
	goto L618
L618:
	;
	v4831 = *(*int32)(unsafe.Add(mBase, uint32(v4780)+12))
	v4835 = *(*int32)(unsafe.Add(mBase, uint32(v4831+v4795<<(uint(int32(2))%32))))
	v4836 = *(*int32)(unsafe.Add(mBase, uint32(v4835)+8))
	v4837 = *(*int32)(unsafe.Add(mBase, uint32(v4790)+8))
	v4838 = int32(0)
	if v4836 == v4838 {
		v4879 = v4838
		goto L621
	} else {
		goto L622
	}
L619:
	;
	goto L615
L620:
	;
	if v4879 == int32(0) {
		goto L634
	} else {
		goto L635
	}
L621:
	;
	goto L620
L622:
	;
	if v4837 == int32(0) {
		v4879 = v4838
		goto L621
	} else {
		goto L623
	}
L623:
	;
	v4847 = *(*int32)(unsafe.Add(mBase, uint32(v4836)+4))
	v4848 = *(*int32)(unsafe.Add(mBase, uint32(v4837)+4))
	if v4847 < v4848 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v4850 = v4847
	goto L626
L625:
	;
	v4850 = v4848
	goto L626
L626:
	;
	if v4850 <= int32(1) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v4853 = int32(1)
	goto L629
L628:
	;
	v4853 = v4850
	goto L629
L629:
	;
	v4854 = int32(8)
	v4859 = int32(0)
	goto L630
L630:
	;
	v4866 = v4859 << (uint(int32(2)) % 32)
	v4868 = *(*int32)(unsafe.Add(mBase, uint32(v4837+v4854+v4866)))
	v4870 = *(*int32)(unsafe.Add(mBase, uint32(v4866+(v4836+v4854))))
	v4871 = v4868 & v4870
	v4873 = base.B2i32(v4871 != int32(0))
	if v4871 != 0 {
		v4879 = v4873
		goto L621
	} else {
		goto L632
	}
L631:
	;
	v4879 = v4873
	goto L621
L632:
	;
	v4875 = v4859 + int32(1)
	if v4875 != v4853 {
		v4859 = v4875
		goto L630
	} else {
		goto L633
	}
L633:
	;
	goto L631
L634:
	;
	v4885 = F_make_join_rel(m, l0, v4790, v4835)
	mBase = m.M
	v4886 = m.ExcPending
	if v4886 != 0 {
		goto L13
	} else {
		goto L637
	}
L635:
	;
	goto L636
L636:
	;
	v4888 = v4795 + int32(1)
	v4889 = *(*int32)(unsafe.Add(mBase, uint32(v4780)+4))
	if v4888 < v4889 {
		v4795 = v4888
		goto L618
	} else {
		goto L638
	}
L637:
	;
	goto L636
L638:
	;
	goto L619
L639:
	;
	goto L614
L640:
	;
	goto L609
L641:
	;
	v5014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v5014 != 0 {
		goto L605
	} else {
		goto L642
	}
L642:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5018 = m.ExcPending
	if v5018 != 0 {
		goto L13
	} else {
		goto L643
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3829))) = v3789
	F_errmsg_internal(m, int32(149603), v3829)
	mBase = m.M
	v5022 = m.ExcPending
	if v5022 != 0 {
		goto L13
	} else {
		goto L644
	}
L644:
	;
	F_errfinish(m, int32(495194), int32(255), int32(306749))
	mBase = m.M
	v5027 = m.ExcPending
	if v5027 != 0 {
		goto L13
	} else {
		goto L645
	}
L645:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L646:
	;
	v5230 = v3789 + int32(1)
	if v5230 <= v46 {
		v3789 = v5230
		goto L435
	} else {
		goto L670
	}
L647:
	;
	v5077 = int32(0)
	v5078 = *(*int32)(unsafe.Add(mBase, uint32(v5074)+4))
	if v5078 <= v5077 {
		goto L646
	} else {
		goto L648
	}
L648:
	;
	v5082 = v5077
	goto L649
L649:
	;
	v5120 = *(*int32)(unsafe.Add(mBase, uint32(v5074)+12))
	v5124 = *(*int32)(unsafe.Add(mBase, uint32(v5120+v5082<<(uint(int32(2))%32))))
	F_generate_partitionwise_join_paths(m, l0, v5124)
	mBase = m.M
	v5126 = m.ExcPending
	if v5126 != 0 {
		goto L13
	} else {
		goto L651
	}
L650:
	;
	goto L646
L651:
	;
	v5127 = *(*int32)(unsafe.Add(mBase, uint32(v5124)+8))
	v5128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v5129 = int32(0)
	v5136 = base.B2i32(v5127|v5128 == v5129)
	if v5127 == v5129 {
		v5175 = v5136
		goto L653
	} else {
		goto L654
	}
L652:
	;
	if v5175 == int32(0) {
		goto L664
	} else {
		goto L665
	}
L653:
	;
	goto L652
L654:
	;
	if v5128 == int32(0) {
		v5175 = v5136
		goto L653
	} else {
		goto L655
	}
L655:
	;
	v5142 = *(*int32)(unsafe.Add(mBase, uint32(v5127)+4))
	v5143 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+4))
	if v5142 != v5143 {
		v5175 = int32(0)
		goto L653
	} else {
		goto L656
	}
L656:
	;
	v5145 = int32(1)
	if v5142 <= v5145 {
		goto L657
	} else {
		goto L658
	}
L657:
	;
	v5148 = v5145
	goto L659
L658:
	;
	v5148 = v5142
	goto L659
L659:
	;
	v5149 = int32(8)
	v5154 = int32(0)
	goto L660
L660:
	;
	v5162 = v5154 << (uint(int32(2)) % 32)
	v5164 = *(*int32)(unsafe.Add(mBase, uint32(v5127+v5149+v5162)))
	v5166 = *(*int32)(unsafe.Add(mBase, uint32(v5162+(v5128+v5149))))
	v5167 = base.B2i32(v5164 == v5166)
	if v5166 != v5164 {
		v5175 = v5167
		goto L653
	} else {
		goto L662
	}
L661:
	;
	v5175 = v5167
	goto L653
L662:
	;
	v5170 = v5154 + int32(1)
	if v5170 != v5148 {
		v5154 = v5170
		goto L660
	} else {
		goto L663
	}
L663:
	;
	goto L661
L664:
	;
	F_generate_useful_gather_paths(m, l0, v5124, int32(0))
	mBase = m.M
	v5183 = m.ExcPending
	if v5183 != 0 {
		goto L13
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	F_set_cheapest(m, v5124)
	mBase = m.M
	v5185 = m.ExcPending
	if v5185 != 0 {
		goto L13
	} else {
		goto L668
	}
L667:
	;
	goto L666
L668:
	;
	v5187 = v5082 + int32(1)
	v5188 = *(*int32)(unsafe.Add(mBase, uint32(v5074)+4))
	if v5187 < v5188 {
		v5082 = v5187
		goto L649
	} else {
		goto L669
	}
L669:
	;
	goto L650
L670:
	;
	goto L436
L671:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5281 = m.ExcPending
	if v5281 != 0 {
		goto L13
	} else {
		goto L674
	}
L672:
	;
	goto L673
L673:
	;
	v5291 = *(*int32)(unsafe.Add(mBase, uint32(v5275)+12))
	v5292 = *(*int32)(unsafe.Add(mBase, uint32(v5291)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	m.G0 = v3774 + int32(16)
	v5300 = v5292
	goto L1
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3774))) = v46
	F_errmsg_internal(m, int32(149603), v3774)
	mBase = m.M
	v5285 = m.ExcPending
	if v5285 != 0 {
		goto L13
	} else {
		goto L675
	}
L675:
	;
	F_errfinish(m, int32(495231), int32(3533), int32(326714))
	mBase = m.M
	v5290 = m.ExcPending
	if v5290 != 0 {
		goto L13
	} else {
		goto L676
	}
L676:
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
	var v58 int32
	_ = v58
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
				v58 = v35
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
				if v39 != int32(1) {
					v58 = v35
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
					if v43 == int32(0) {
						v58 = int32(0)
					} else {
						v46 = int32(1)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v47 != int32(7) {
							v58 = v46
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
							if v50 != 0 {
								v58 = int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
								v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
								if v52 != int64(0) {
									v58 = v46
								} else {
									v58 = int32(0)
								}
							}
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
			if v43 == int32(0) {
				v58 = int32(0)
			} else {
				v46 = int32(1)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47 != int32(7) {
					v58 = v46
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
					if v50 != 0 {
						v58 = int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
						if v52 != int64(0) {
							v58 = v46
						} else {
							v58 = int32(0)
						}
					}
				}
			}
		}
		if v58 == int32(0) {
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 float64
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int64
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 float64
	_ = v450
	var v451 float64
	_ = v451
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 float64
	_ = v506
	var v507 float64
	_ = v507
	var v525 float64
	_ = v525
	var v535 float64
	_ = v535
	var v536 float64
	_ = v536
	var v538 float64
	_ = v538
	var v540 float64
	_ = v540
	var v541 float64
	_ = v541
	var v559 float64
	_ = v559
	var v569 float64
	_ = v569
	var v570 float64
	_ = v570
	var v572 float64
	_ = v572
	var v573 int32
	_ = v573
	var v574 float64
	_ = v574
	var v575 float64
	_ = v575
	var v579 float64
	_ = v579
	var v582 float64
	_ = v582
	var v584 float64
	_ = v584
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 float64
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 float64
	_ = v657
	var v658 float64
	_ = v658
	var v659 int32
	_ = v659
	var v660 int64
	_ = v660
	var v669 int32
	_ = v669
	var v678 int32
	_ = v678
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 float64
	_ = v713
	var v714 float64
	_ = v714
	var v732 float64
	_ = v732
	var v742 float64
	_ = v742
	var v743 float64
	_ = v743
	var v745 float64
	_ = v745
	var v747 float64
	_ = v747
	var v748 float64
	_ = v748
	var v766 float64
	_ = v766
	var v776 float64
	_ = v776
	var v777 float64
	_ = v777
	var v779 float64
	_ = v779
	var v780 int32
	_ = v780
	var v781 float64
	_ = v781
	var v782 float64
	_ = v782
	var v786 float64
	_ = v786
	var v789 float64
	_ = v789
	var v791 float64
	_ = v791
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 float64
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 float64
	_ = v834
	var v835 int64
	_ = v835
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 float64
	_ = v887
	var v888 float64
	_ = v888
	var v889 float64
	_ = v889
	var v908 float64
	_ = v908
	var v909 float64
	_ = v909
	var v917 float64
	_ = v917
	var v918 float64
	_ = v918
	var v920 float64
	_ = v920
	var v922 float64
	_ = v922
	var v924 float64
	_ = v924
	var v926 float64
	_ = v926
	var v927 float64
	_ = v927
	var v946 float64
	_ = v946
	var v947 float64
	_ = v947
	var v948 float64
	_ = v948
	var v955 float64
	_ = v955
	var v956 float64
	_ = v956
	var v958 float64
	_ = v958
	var v959 int32
	_ = v959
	var v960 float64
	_ = v960
	var v961 float64
	_ = v961
	var v964 float64
	_ = v964
	var v966 float64
	_ = v966
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1126 float64
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 float64
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 float64
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int64
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1156 int32
	_ = v1156
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 float64
	_ = v1185
	var v1211 float64
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1220 float64
	_ = v1220
	var v1221 float64
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int64
	_ = v1223
	var v1232 int32
	_ = v1232
	var v1247 int32
	_ = v1247
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 float64
	_ = v1276
	var v1277 float64
	_ = v1277
	var v1295 float64
	_ = v1295
	var v1305 float64
	_ = v1305
	var v1306 float64
	_ = v1306
	var v1308 float64
	_ = v1308
	var v1310 float64
	_ = v1310
	var v1311 float64
	_ = v1311
	var v1329 float64
	_ = v1329
	var v1339 float64
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 float64
	_ = v1341
	var v1343 float64
	_ = v1343
	var v1344 float64
	_ = v1344
	var v1348 float64
	_ = v1348
	var v1350 float64
	_ = v1350
	var v1352 float64
	_ = v1352
	var v1361 float64
	_ = v1361
	var v1366 float64
	_ = v1366
	var v1379 int32
	_ = v1379
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1433 int32
	_ = v1433
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1526 int32
	_ = v1526
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1655 int32
	_ = v1655
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1723 int32
	_ = v1723
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1741 int32
	_ = v1741
	var v1749 int32
	_ = v1749
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1798 int32
	_ = v1798
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1835 int32
	_ = v1835
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1873 int32
	_ = v1873
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1984 int32
	_ = v1984
	var v1993 int32
	_ = v1993
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2061 int32
	_ = v2061
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2094 int32
	_ = v2094
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2143 float64
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2150 int32
	_ = v2150
	var v2156 int32
	_ = v2156
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2199 int32
	_ = v2199
	var v2205 int32
	_ = v2205
	var v2214 int32
	_ = v2214
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2316 int32
	_ = v2316
	var v2333 int32
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2366 int32
	_ = v2366
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2398 int32
	_ = v2398
	var v2403 int32
	_ = v2403
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2451 int32
	_ = v2451
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2474 int32
	_ = v2474
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2500 int32
	_ = v2500
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 float64
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2584 int32
	_ = v2584
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2611 int32
	_ = v2611
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2637 int32
	_ = v2637
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2655 int32
	_ = v2655
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2671 int32
	_ = v2671
	var v2676 int32
	_ = v2676
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	v5 = int32(0)
	v18 = float64(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v34 == v5 {
		v54 = v5
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	if v2624 != 0 {
		goto L415
	} else {
		goto L416
	}
L2:
	;
	if v54 != 0 {
		v2596 = l0
		v2597 = l1
		v2598 = l2
		v2599 = l3
		v2611 = v30
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
	v54 = int32(1)
	goto L3
L7:
	;
	if v42 != int32(290) {
		v54 = v5
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
		v54 = v5
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
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v157 {
	case 0:
		goto L52
	case 1, 6, 7, 8:
		v2596 = l0
		v2597 = l1
		v2598 = l2
		v2599 = l3
		v2611 = v30
		goto L1
	default:
		goto L48
	case 3:
		goto L51
	case 4:
		goto L50
	case 5:
		goto L49
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
	if v65 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	return
L20:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L21:
	;
	F_add_paths_to_append_rel(m, l0, l1, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v76 = v5
	v77 = v5
	goto L25
L24:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L25:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v77<<(uint(int32(2))%32))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v103 != l2 {
		v147 = v76
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_add_paths_to_append_rel(m, l0, l1, v147)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L19
	} else {
		goto L46
	}
L27:
	;
	v152 = v77 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v152 < v153 {
		v76 = v147
		v77 = v152
		goto L25
	} else {
		goto L45
	}
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v107 = v105 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111+v107)))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v114 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+26)) = uint8(v117)
	goto L31
L30:
	;
	goto L31
L31:
	;
	F_set_rel_pathlist(m, l0, v110, v105, v113)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v121 = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v110)+32))
	if v123 == v121 {
		v143 = v121
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v143 != 0 {
		v147 = v76
		goto L27
	} else {
		goto L43
	}
L34:
	;
	goto L33
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v127 = v126
	goto L36
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if base.Ui32(int32(2)) <= base.Ui32(v131-int32(301)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v143 = int32(1)
	goto L34
L38:
	;
	if v131 != int32(290) {
		v143 = v121
		goto L34
	} else {
		goto L41
	}
L39:
	;
	v127 = v130 + int32(72)
	goto L36
L40:
	;
	goto L37
L41:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
	if v138 != 0 {
		v143 = v121
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v145 = F_lappend(m, v76, v110)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	v147 = v145
	goto L27
L45:
	;
	goto L26
L46:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L47:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v997 = m.G0
	v999 = v997 - int32(16)
	m.G0 = v999
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v1004 = F_TidQualFromRestrictInfoList(m, l0, v1001, l1, v999+int32(15))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L19
	} else {
		goto L174
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L19
	} else {
		goto L171
	}
L49:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v807 = F_palloc0(m, int32(72))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L19
	} else {
		goto L157
	}
L50:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v600 = F_palloc0(m, int32(72))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L19
	} else {
		goto L135
	}
L51:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+72)))
	if v255 != int32(1) {
		v372 = v5
		goto L86
	} else {
		goto L87
	}
L52:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v158 == int32(102) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	m.T0[v163].(func(*base.Module, int32, int32, int32))(m, l0, l1, v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v166 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v171 = F_palloc0(m, int32(72))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = int64(1460288880919)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+12)) = v176
	v178 = F_get_baserel_parampathinfo(m, l0, l1, v169)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+20)) = uint8(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v171)+16)) = v178
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+64)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v180
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+21)) = uint8(v183)
	F_cost_samplescan(m, v171, l0, l1, v178)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v191) <= base.Ui32(int32(1)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	F_add_path(m, l1, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L19
	} else {
		goto L85
	}
L62:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v194 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v246 = F_GetTsmRoutine(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L82
	}
L65:
	;
	if v241 == int32(1) {
		v251 = v171
		goto L61
	} else {
		goto L81
	}
L66:
	;
	v241 = int32(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v203 = int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v204 <= v203 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v207 = v203
	goto L71
L70:
	;
	v207 = v204
	goto L71
L71:
	;
	v210 = int32(0)
	v212 = v210
	v213 = v210
	goto L72
L72:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v194+int32(8)+v212<<(uint(int32(2))%32))))
	if v221 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v241 = v234
	goto L65
L74:
	;
	goto L73
L75:
	;
	v222 = int32(2)
	if v213 != 0 {
		v234 = v222
		goto L74
	} else {
		goto L78
	}
L76:
	;
	v227 = v213
	goto L77
L77:
	;
	v230 = v212 + int32(1)
	if v230 != v207 {
		v212 = v230
		v213 = v227
		goto L72
	} else {
		goto L80
	}
L78:
	;
	v223 = int32(1)
	if base.Ui32(v223) < base.Ui32(base.I32_popcnt(v221)) {
		v234 = v222
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v227 = v223
	goto L77
L80:
	;
	v234 = v227
	goto L74
L81:
	;
	goto L64
L82:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+9)))
	if v248 != 0 {
		v251 = v171
		goto L61
	} else {
		goto L83
	}
L83:
	;
	v249 = F_create_material_path(m, l1, v171)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L19
	} else {
		goto L84
	}
L84:
	;
	v251 = v249
	goto L61
L85:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L86:
	;
	v394 = F_palloc0(m, int32(72))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L19
	} else {
		goto L113
	}
L87:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v259 == int32(0) {
		v372 = v5
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v262 <= int32(0) {
		v372 = v5
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
	v273 = v5
	goto L91
L90:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v315 = m.G0
	v317 = v315 - int32(48)
	m.G0 = v317
	v326 = F_get_ordering_op_properties(m, int32(412), v317+int32(44), v317+int32(40), v317+int32(36))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L19
	} else {
		goto L100
	}
L91:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v265+v273<<(uint(int32(2))%32))))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	if v298 != int32(6) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v372 = int32(0)
	goto L86
L93:
	;
	v310 = v273 + int32(1)
	if v262 != v310 {
		v273 = v310
		goto L91
	} else {
		goto L98
	}
L94:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+8)))
	if v301 != v266 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v303 != v304 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v297)+28))
	if v306 == int32(0) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	goto L93
L98:
	;
	goto L92
L99:
	;
	v372 = v346
	goto L86
L100:
	;
	if v326 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317)+44))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v317)+40))
	v330 = F_exprCollation(m, v297)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L19
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L19
	} else {
		goto L110
	}
L104:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v317)+36))
	v334 = base.B2i32(v332 == int32(5))
	v335 = int32(0)
	v337 = F_make_pathkey_from_sortinfo(m, l0, v297, v328, v329, v330, v334, v334, v335, v313, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L19
	} else {
		goto L105
	}
L105:
	;
	if v337 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v317)+32)) = v337
	v344 = F_list_make1_impl(m, int32(1), v317+int32(12))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L19
	} else {
		goto L109
	}
L107:
	;
	v346 = int32(0)
	goto L108
L108:
	;
	m.G0 = v317 + int32(48)
	goto L99
L109:
	;
	v346 = v344
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = int32(412)
	F_errmsg_internal(m, int32(209605), v317+int32(16))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(494711), int32(1016), int32(21063))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = int64(1494648619287)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = v399
	v401 = F_get_baserel_parampathinfo(m, l0, l1, v254)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	v403 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v394)+20)) = uint8(v403)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+16)) = v401
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+64)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v394)+24)) = v403
	*(*uint8)(unsafe.Add(mBase, uint32(v394)+21)) = uint8(v406)
	v411 = m.G0
	v413 = v411 - int32(32)
	m.G0 = v413
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v415 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v401 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v429 = v415 + v416<<(uint(int32(2))%32)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+52))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v429 = v422 + v423<<(uint(int32(2))%32) - int32(4)
	goto L115
L119:
	;
	v435 = v401 + int32(8)
	goto L121
L120:
	;
	v435 = l1 + int32(16)
	goto L121
L121:
	;
	v436 = *(*float64)(unsafe.Add(mBase, uint32(v435)))
	*(*float64)(unsafe.Add(mBase, uint32(v394)+32)) = v436
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v430)+68))
	v440 = v413 + int32(24)
	v441 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v440))) = v441
	*(*int64)(unsafe.Add(mBase, uint32(v413)+16)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v413)+8)) = l0
	v448 = F_cost_qual_eval_walker(m, v438, v413+int32(8))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L19
	} else {
		goto L122
	}
L122:
	;
	v450 = *(*float64)(unsafe.Add(mBase, uint32(v440)))
	v451 = *(*float64)(unsafe.Add(mBase, uint32(v413)+16))
	if v401 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v570 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v572 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v574 = *(*float64)(unsafe.Add(mBase, uint32(v573)+24))
	v575 = *(*float64)(unsafe.Add(mBase, uint32(v573)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+40)) = int32(0)
	v579 = float64(0)
	v582 = base.F64_add(v575, base.F64_add(base.F64_add(base.F64_add(v451, v450), v579), v569))
	*(*float64)(unsafe.Add(mBase, uint32(v394)+48)) = v582
	v584 = *(*float64)(unsafe.Add(mBase, uint32(v394)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v394)+56)) = base.F64_add(v582, base.F64_add(base.F64_mul(v574, v584), base.F64_add(base.F64_mul(v570, base.F64_add(v559, v572)), v579)))
	m.G0 = v413 + int32(32)
	F_add_path(m, l1, v394)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L19
	} else {
		goto L134
	}
L124:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v401)+16))
	v453 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v440))) = v453
	*(*int64)(unsafe.Add(mBase, uint32(v413)+16)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v413)+8)) = l0
	if v452 == int32(0) {
		v525 = v18
		v535 = float64(0)
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v540 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v541 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v559 = v540
	v569 = v541
	goto L123
L127:
	;
	v536 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v538 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v559 = base.F64_add(v525, v536)
	v569 = base.F64_add(v535, v538)
	goto L123
L128:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v462 <= int32(0) {
		v525 = v18
		v535 = float64(0)
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v471 = int32(0)
	goto L130
L130:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v452)+12))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v493+v471<<(uint(int32(2))%32))))
	v500 = F_cost_qual_eval_walker(m, v497, v413+int32(8))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L19
	} else {
		goto L132
	}
L131:
	;
	v506 = *(*float64)(unsafe.Add(mBase, uint32(v413)+24))
	v507 = *(*float64)(unsafe.Add(mBase, uint32(v413)+16))
	v525 = v506
	v535 = v507
	goto L127
L132:
	;
	v503 = v471 + int32(1)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v503 < v504 {
		v471 = v503
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v600)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v600))) = int64(1503238553879)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+12)) = v605
	v607 = F_get_baserel_parampathinfo(m, l0, l1, v598)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	v609 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v600)+20)) = uint8(v609)
	*(*int32)(unsafe.Add(mBase, uint32(v600)+16)) = v607
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+64)) = v609
	*(*int32)(unsafe.Add(mBase, uint32(v600)+24)) = v609
	*(*uint8)(unsafe.Add(mBase, uint32(v600)+21)) = uint8(v612)
	v618 = m.G0
	v620 = v618 - int32(32)
	m.G0 = v620
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v622 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	if v607 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v636 = v622 + v623<<(uint(int32(2))%32)
	goto L137
L139:
	;
	goto L140
L140:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+52))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v636 = v629 + v630<<(uint(int32(2))%32) - int32(4)
	goto L137
L141:
	;
	v642 = v607 + int32(8)
	goto L143
L142:
	;
	v642 = l1 + int32(16)
	goto L143
L143:
	;
	v643 = *(*float64)(unsafe.Add(mBase, uint32(v642)))
	*(*float64)(unsafe.Add(mBase, uint32(v600)+32)) = v643
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v637)+76))
	v647 = v620 + int32(24)
	v648 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v647))) = v648
	*(*int64)(unsafe.Add(mBase, uint32(v620)+16)) = v648
	*(*int32)(unsafe.Add(mBase, uint32(v620)+8)) = l0
	v655 = F_cost_qual_eval_walker(m, v645, v620+int32(8))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	v657 = *(*float64)(unsafe.Add(mBase, uint32(v647)))
	v658 = *(*float64)(unsafe.Add(mBase, uint32(v620)+16))
	if v607 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v777 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v779 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v600)+12))
	v781 = *(*float64)(unsafe.Add(mBase, uint32(v780)+24))
	v782 = *(*float64)(unsafe.Add(mBase, uint32(v780)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+40)) = int32(0)
	v786 = float64(0)
	v789 = base.F64_add(v782, base.F64_add(base.F64_add(base.F64_add(v658, v657), v786), v776))
	*(*float64)(unsafe.Add(mBase, uint32(v600)+48)) = v789
	v791 = *(*float64)(unsafe.Add(mBase, uint32(v600)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v600)+56)) = base.F64_add(v789, base.F64_add(base.F64_mul(v781, v791), base.F64_add(base.F64_mul(v777, base.F64_add(v766, v779)), v786)))
	m.G0 = v620 + int32(32)
	F_add_path(m, l1, v600)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L19
	} else {
		goto L156
	}
L146:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v607)+16))
	v660 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v647))) = v660
	*(*int64)(unsafe.Add(mBase, uint32(v620)+16)) = v660
	*(*int32)(unsafe.Add(mBase, uint32(v620)+8)) = l0
	if v659 == int32(0) {
		v732 = v18
		v742 = float64(0)
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L148
L148:
	;
	v747 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v748 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v766 = v747
	v776 = v748
	goto L145
L149:
	;
	v743 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v745 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v766 = base.F64_add(v732, v743)
	v776 = base.F64_add(v742, v745)
	goto L145
L150:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v669 <= int32(0) {
		v732 = v18
		v742 = float64(0)
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v678 = int32(0)
	goto L152
L152:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v659)+12))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700+v678<<(uint(int32(2))%32))))
	v707 = F_cost_qual_eval_walker(m, v704, v620+int32(8))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L19
	} else {
		goto L154
	}
L153:
	;
	v713 = *(*float64)(unsafe.Add(mBase, uint32(v620)+24))
	v714 = *(*float64)(unsafe.Add(mBase, uint32(v620)+16))
	v732 = v713
	v742 = v714
	goto L149
L154:
	;
	v710 = v678 + int32(1)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v710 < v711 {
		v678 = v710
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v807))) = int64(1498943586583)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+12)) = v812
	v814 = F_get_baserel_parampathinfo(m, l0, l1, v805)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	v816 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v807)+20)) = uint8(v816)
	*(*int32)(unsafe.Add(mBase, uint32(v807)+16)) = v814
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+64)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(v807)+24)) = v816
	*(*uint8)(unsafe.Add(mBase, uint32(v807)+21)) = uint8(v819)
	v825 = m.G0
	v827 = v825 - int32(32)
	m.G0 = v827
	if v814 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v956 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v958 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v807)+12))
	v960 = *(*float64)(unsafe.Add(mBase, uint32(v959)+24))
	v961 = *(*float64)(unsafe.Add(mBase, uint32(v959)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+40)) = int32(0)
	v964 = float64(0)
	v966 = base.F64_add(v961, base.F64_add(v955, v964))
	*(*float64)(unsafe.Add(mBase, uint32(v807)+48)) = v966
	*(*float64)(unsafe.Add(mBase, uint32(v807)+56)) = base.F64_add(v966, base.F64_add(base.F64_mul(v960, v947), base.F64_add(base.F64_mul(v956, base.F64_add(v948, base.F64_add(v946, v958))), v964)))
	m.G0 = v827 + int32(32)
	F_add_path(m, l1, v807)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L19
	} else {
		goto L170
	}
L160:
	;
	v829 = *(*float64)(unsafe.Add(mBase, uint32(v814)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v807)+32)) = v829
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v814)+16))
	v832 = int32(0)
	v834 = *(*float64)(unsafe.Add(mBase, _consts[614]))
	v835 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v827)+24)) = v835
	*(*int64)(unsafe.Add(mBase, uint32(v827)+16)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v827)+8)) = l0
	if v831 == v832 {
		v908 = v18
		v909 = v829
		v917 = float64(0)
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v922 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v807)+32)) = v922
	v924 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v926 = *(*float64)(unsafe.Add(mBase, _consts[614]))
	v927 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v946 = v924
	v947 = v922
	v948 = v926
	v955 = v927
	goto L159
L163:
	;
	v918 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v920 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v946 = base.F64_add(v908, v918)
	v947 = v909
	v948 = v834
	v955 = base.F64_add(v917, v920)
	goto L159
L164:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v844 <= int32(0) {
		v908 = v18
		v909 = v829
		v917 = float64(0)
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v852 = v832
	goto L166
L166:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874+v852<<(uint(int32(2))%32))))
	v881 = F_cost_qual_eval_walker(m, v878, v827+int32(8))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L19
	} else {
		goto L168
	}
L167:
	;
	v887 = *(*float64)(unsafe.Add(mBase, uint32(v807)+32))
	v888 = *(*float64)(unsafe.Add(mBase, uint32(v827)+24))
	v889 = *(*float64)(unsafe.Add(mBase, uint32(v827)+16))
	v908 = v888
	v909 = v887
	v917 = v889
	goto L163
L168:
	;
	v884 = v852 + int32(1)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v884 < v885 {
		v852 = v884
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v2596 = l0
	v2597 = l1
	v2598 = l2
	v2599 = l3
	v2611 = v30
	goto L1
L171:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v986
	F_errmsg_internal(m, int32(488444), v30)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L19
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(495231), int32(527), int32(73988))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L19
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, _consts[615])))
	if v1004 != 0 {
		goto L179
	} else {
		goto L180
	}
L175:
	;
	m.G0 = v999 + int32(16)
	if v1433 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L176:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1422 = F_create_tidscan_path(m, l0, l1, v1004, v1421)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L19
	} else {
		goto L237
	}
L177:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+152)))
	if v1033&int32(1) == int32(0) {
		goto L189
	} else {
		goto L190
	}
L178:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1017 = F_create_tidscan_path(m, l0, l1, v1004, v1016)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L19
	} else {
		goto L185
	}
L179:
	;
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999)+15)))
	if v1007&int32(1) != 0 {
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v1007&int32(1) != 0 {
		v1031 = v5
		goto L177
	} else {
		goto L184
	}
L182:
	;
	if v1008&int32(1) != 0 {
		goto L176
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v1433 = v5
	goto L175
L185:
	;
	F_add_path(m, l1, v1017)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L19
	} else {
		goto L186
	}
L186:
	;
	v1021 = int32(1)
	if v1008&v1021 != 0 {
		v1433 = v1021
		goto L175
	} else {
		goto L187
	}
L187:
	;
	v1024 = int32(0)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, _consts[615])))
	if v1026&int32(1) == v1024 {
		v1433 = v1024
		goto L175
	} else {
		goto L188
	}
L188:
	;
	v1031 = v1024
	goto L177
L189:
	;
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v1407 == int32(1) {
		goto L231
	} else {
		goto L232
	}
L190:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v1038 == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	if v1041 <= int32(0) {
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v1051 = v1031
	v1052 = int32(0)
	goto L193
L193:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+12))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1072+v1051<<(uint(int32(2))%32))))
	v1077 = F_IsBinaryTidClause(m, v1076, l1)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L19
	} else {
		goto L196
	}
L194:
	;
	if v1089 == int32(0) {
		goto L189
	} else {
		goto L201
	}
L195:
	;
	v1091 = v1051 + int32(1)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	if v1091 < v1092 {
		v1051 = v1091
		v1052 = v1089
		goto L193
	} else {
		goto L200
	}
L196:
	;
	if v1077 == int32(0) {
		v1089 = v1052
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+4))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	if base.Ui32(int32(3)) < base.Ui32(v1082-int32(2799)) {
		v1089 = v1052
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v1087 = F_lappend(m, v1052, v1076)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L19
	} else {
		goto L199
	}
L199:
	;
	v1089 = v1087
	goto L195
L200:
	;
	goto L194
L201:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1098 = F_palloc0(m, int32(80))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L19
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1098))) = int64(1486058684702)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+12)) = v1103
	v1105 = F_get_baserel_parampathinfo(m, l0, l1, v1096)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L19
	} else {
		goto L203
	}
L203:
	;
	v1107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1098)+20)) = uint8(v1107)
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+16)) = v1105
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+72)) = v1089
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+64)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+24)) = v1107
	*(*uint8)(unsafe.Add(mBase, uint32(v1098)+21)) = uint8(v1110)
	v1117 = m.G0
	v1119 = v1117 - int32(48)
	m.G0 = v1119
	if v1105 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1125 = v1105 + int32(8)
	goto L206
L205:
	;
	v1125 = l1 + int32(16)
	goto L206
L206:
	;
	v1126 = *(*float64)(unsafe.Add(mBase, uint32(v1125)))
	*(*float64)(unsafe.Add(mBase, uint32(v1098)+32)) = v1126
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1129 = int32(0)
	v1131 = F_clauselist_selectivity(m, l0, v1089, v1128, v1129, v1129)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L19
	} else {
		goto L207
	}
L207:
	;
	v1133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v1135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1119)+40)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1119)+32)) = v1135
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+24)) = l0
	if v1089 == int32(0) {
		v1211 = v18
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_get_tablespace_page_costs(m, v1213, v1119+int32(16), v1119+int32(8))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L19
	} else {
		goto L215
	}
L209:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	if v1142 <= int32(0) {
		v1211 = v18
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1156 = v5
	goto L211
L211:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+12))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1172+v1156<<(uint(int32(2))%32))))
	v1179 = F_cost_qual_eval_walker(m, v1176, v1119+int32(24))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L19
	} else {
		goto L213
	}
L212:
	;
	v1185 = *(*float64)(unsafe.Add(mBase, uint32(v1119)+40))
	v1211 = v1185
	goto L208
L213:
	;
	v1182 = v1156 + int32(1)
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	if v1182 < v1183 {
		v1156 = v1182
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v1220 = *(*float64)(unsafe.Add(mBase, uint32(v1119)+8))
	v1221 = *(*float64)(unsafe.Add(mBase, uint32(v1119)+16))
	if v1105 != 0 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+12))
	v1341 = *(*float64)(unsafe.Add(mBase, uint32(v1340)+24))
	v1343 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v1344 = *(*float64)(unsafe.Add(mBase, uint32(v1340)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+40)) = int32(0)
	v1348 = float64(0)
	v1350 = base.F64_add(v1344, base.F64_add(base.F64_add(v1211, v1339), v1348))
	*(*float64)(unsafe.Add(mBase, uint32(v1098)+48)) = v1350
	v1352 = *(*float64)(unsafe.Add(mBase, uint32(v1098)+32))
	v1361 = base.F64_ceil(base.F64_mul(v1131, base.F64_convert_i32_u(v1134)))
	if base.F64_le(v1361, v1348) != 0 {
		goto L227
	} else {
		goto L228
	}
L217:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+16))
	v1223 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1119)+40)) = v1223
	*(*int64)(unsafe.Add(mBase, uint32(v1119)+32)) = v1223
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+24)) = l0
	if v1222 == int32(0) {
		v1295 = v18
		v1305 = float64(0)
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L219
L219:
	;
	v1310 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v1311 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v1329 = v1310
	v1339 = v1311
	goto L216
L220:
	;
	v1306 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v1308 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v1329 = base.F64_add(v1295, v1306)
	v1339 = base.F64_add(v1305, v1308)
	goto L216
L221:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+4))
	if v1232 <= int32(0) {
		v1295 = v18
		v1305 = float64(0)
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v1247 = int32(0)
	goto L223
L223:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+12))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1263+v1247<<(uint(int32(2))%32))))
	v1270 = F_cost_qual_eval_walker(m, v1267, v1119+int32(24))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L19
	} else {
		goto L225
	}
L224:
	;
	v1276 = *(*float64)(unsafe.Add(mBase, uint32(v1119)+40))
	v1277 = *(*float64)(unsafe.Add(mBase, uint32(v1119)+32))
	v1295 = v1276
	v1305 = v1277
	goto L220
L225:
	;
	v1273 = v1247 + int32(1)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+4))
	if v1273 < v1274 {
		v1247 = v1273
		goto L223
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v1366 = v1348
	goto L229
L228:
	;
	v1366 = base.F64_add(v1361, float64(-1))
	goto L229
L229:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1098)+56)) = base.F64_add(v1350, base.F64_add(base.F64_mul(v1341, v1352), base.F64_add(base.F64_mul(base.F64_sub(base.F64_add(v1329, v1343), v1211), base.F64_mul(v1131, v1133)), base.F64_add(base.F64_add(base.F64_mul(v1220, v1366), v1221), float64(0)))))
	m.G0 = v1119 + int32(48)
	F_add_path(m, l1, v1098)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L19
	} else {
		goto L230
	}
L230:
	;
	goto L189
L231:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v1413 = F_generate_implied_equalities_for_column(m, l0, l1, int32(827), int32(0), v1412)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L19
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	F_BuildParameterizedTidPaths(m, l0, l1, v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L19
	} else {
		goto L236
	}
L234:
	;
	F_BuildParameterizedTidPaths(m, l0, l1, v1413)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L19
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	v1433 = int32(0)
	goto L175
L237:
	;
	F_add_path(m, l1, v1422)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L19
	} else {
		goto L238
	}
L238:
	;
	v1433 = int32(1)
	goto L175
L239:
	;
	v1460 = F_create_seqscan_path(m, l0, l1, v996, int32(0))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L19
	} else {
		goto L242
	}
L240:
	;
	v2569 = l0
	v2570 = l1
	v2571 = l2
	v2572 = l3
	v2584 = v30
	goto L241
L241:
	;
	v2596 = v2569
	v2597 = v2570
	v2598 = v2571
	v2599 = v2572
	v2611 = v2584
	goto L1
L242:
	;
	F_add_path(m, l1, v1460)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L19
	} else {
		goto L243
	}
L243:
	;
	v1464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1464 != int32(1) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1584 = m.G0
	v1586 = v1584 - int32(416)
	m.G0 = v1586
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v1588 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L245:
	;
	if v996 != 0 {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	if v1469 != int32(-1) {
		v1526 = v1469
		goto L247
	} else {
		goto L248
	}
L247:
	;
	if v1526 < v1468 {
		goto L257
	} else {
		goto L258
	}
L248:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1476 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if base.B2i32(v1472 == int32(0))&base.B2i32(base.I64_extend_i32_u(v1478) < base.I64_extend_i32_s(v1476)) != 0 {
		goto L244
	} else {
		goto L249
	}
L249:
	;
	v1482 = int32(1)
	if v1476 <= v1482 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1485 = v1482
	goto L252
L251:
	;
	v1485 = v1476
	goto L252
L252:
	;
	v1491 = v1485
	v1492 = int32(1)
	goto L253
L253:
	;
	v1515 = v1491 * int32(3)
	if base.Ui32(v1478) < base.Ui32(v1515) {
		v1526 = v1492
		goto L247
	} else {
		goto L255
	}
L254:
	;
	v1526 = v1518
	goto L247
L255:
	;
	v1518 = v1492 + int32(1)
	if v1515 < int32(715827883) {
		v1491 = v1515
		v1492 = v1518
		goto L253
	} else {
		goto L256
	}
L256:
	;
	goto L254
L257:
	;
	v1549 = v1526
	goto L259
L258:
	;
	v1549 = v1468
	goto L259
L259:
	;
	if v1549 <= int32(0) {
		goto L244
	} else {
		goto L260
	}
L260:
	;
	v1553 = F_create_seqscan_path(m, l0, l1, int32(0), v1549)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L19
	} else {
		goto L261
	}
L261:
	;
	F_add_partial_path(m, l1, v1553)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L19
	} else {
		goto L262
	}
L262:
	;
	goto L244
L263:
	;
	m.G0 = v1586 + int32(416)
	v2569 = l0
	v2570 = l1
	v2571 = l2
	v2572 = l3
	v2584 = v30
	goto L241
L264:
	;
	v1591 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+404)) = v1591
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+400)) = v1591
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1588)+4))
	if v1591 < v1595 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1614 = v5
	v1618 = v5
	goto L268
L266:
	;
	v2094 = v5
	v2109 = int32(0)
	goto L267
L267:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v2112 = F_generate_bitmap_or_paths(m, l0, l1, v2110, int32(0))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L19
	} else {
		goto L333
	}
L268:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1588)+12))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1629+v1618<<(uint(int32(2))%32))))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+88))
	if v1634 != 0 {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+404))
	v2094 = v2061
	v2109 = v2080
	goto L267
L270:
	;
	v2077 = v1618 + int32(1)
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v1588)+4))
	if v2077 < v2078 {
		v1614 = v2061
		v1618 = v2077
		goto L268
	} else {
		goto L332
	}
L271:
	;
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633)+100)))
	if v1635 != int32(1) {
		v2061 = v1614
		goto L270
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v1643 = F__emscripten_memset_bulkmem(m, v1586+int32(268), base.I32_extend8_s(int32(0)), int32(132))
	mBase = m.M
	goto L275
L274:
	;
	goto L273
L275:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+96))
	if v1644 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	F_get_index_paths(m, l0, l1, v1633, v1586+int32(268), v1586+int32(404))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L19
	} else {
		goto L283
	}
L277:
	;
	v1647 = int32(0)
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+4))
	if v1648 <= v1647 {
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1655 = v1647
	goto L279
L279:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+12))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1678+v1655<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v1682, v1633, v1586+int32(268))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L19
	} else {
		goto L281
	}
L280:
	;
	goto L276
L281:
	;
	v1688 = v1655 + int32(1)
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+4))
	if v1688 < v1689 {
		v1655 = v1688
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	v1729 = F__emscripten_memset_bulkmem(m, v1586+int32(136), base.I32_extend8_s(int32(0)), int32(132))
	mBase = m.M
	goto L284
L284:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v1730 == int32(0) {
		v1798 = v1614
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1818 = F__emscripten_memset_bulkmem(m, v1586+int32(4), base.I32_extend8_s(int32(0)), int32(132))
	mBase = m.M
	goto L301
L286:
	;
	v1733 = int32(0)
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+4))
	if v1734 <= v1733 {
		v1798 = v1614
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1741 = v1733
	v1749 = v1614
	goto L288
L288:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+12))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1764+v1741<<(uint(int32(2))%32))))
	v1769 = F_join_clause_is_movable_to(m, v1768, l1)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L19
	} else {
		goto L290
	}
L289:
	;
	v1798 = v1781
	goto L285
L290:
	;
	if v1769 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+52))
	goto L294
L292:
	;
	v1781 = v1749
	goto L293
L293:
	;
	v1783 = v1741 + int32(1)
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+4))
	if v1783 < v1784 {
		v1741 = v1783
		v1749 = v1781
		goto L288
	} else {
		goto L300
	}
L294:
	;
	if v1771 != int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1774 = F_list_append_unique_ptr(m, v1749, v1768)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L19
	} else {
		goto L298
	}
L296:
	;
	v1776 = v1749
	goto L297
L297:
	;
	F_match_clause_to_index(m, l0, v1768, v1633, v1586+int32(136))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L19
	} else {
		goto L299
	}
L298:
	;
	v1776 = v1774
	goto L297
L299:
	;
	v1781 = v1776
	goto L293
L300:
	;
	goto L289
L301:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+12))
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819)+216)))
	if v1820 != int32(1) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586)+136)))
	if v1967 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L303:
	;
	v1823 = int32(0)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+40))
	if v1824 <= v1823 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1835 = v1823
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+412)) = v1835
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+408)) = v1633
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+12))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+104))
	v1861 = F_generate_implied_equalities_for_column(m, l0, v1856, int32(821), v1586+int32(408), v1860)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L19
	} else {
		goto L308
	}
L306:
	;
	goto L302
L307:
	;
	v1937 = v1835 + int32(1)
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+40))
	if v1937 < v1938 {
		v1835 = v1937
		goto L305
	} else {
		goto L315
	}
L308:
	;
	if v1861 == int32(0) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1865 = int32(0)
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+4))
	if v1866 <= v1865 {
		goto L307
	} else {
		goto L310
	}
L310:
	;
	v1873 = v1865
	goto L311
L311:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+12))
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1896+v1873<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v1900, v1633, v1586+int32(4))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L19
	} else {
		goto L313
	}
L312:
	;
	goto L307
L313:
	;
	v1906 = v1873 + int32(1)
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+4))
	if v1906 < v1907 {
		v1873 = v1906
		goto L311
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	goto L306
L316:
	;
	v1970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586)+4)))
	if v1970 != int32(1) {
		v2061 = v1798
		goto L270
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1973 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+408)) = v1973
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+40))
	if v1977 <= v1973 {
		v2061 = v1798
		goto L270
	} else {
		goto L320
	}
L319:
	;
	goto L318
L320:
	;
	v1984 = v1973
	v1993 = v1973
	goto L321
L321:
	;
	v2009 = v1984 << (uint(int32(2)) % 32)
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v1586+int32(140)+v2009)))
	if v2011 != 0 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v2061 = v1798
	goto L270
L323:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v2011)+4))
	v2013 = v2012
	goto L325
L324:
	;
	v2013 = int32(0)
	goto L325
L325:
	;
	v2022 = v2013 + v1993
	F_consider_index_join_outer_rels(m, l0, l1, v1633, v1586+int32(268), v1586+int32(136), v1586+int32(4), v1586+int32(400), v2011, v2022, v1586+int32(408))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L19
	} else {
		goto L326
	}
L326:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2009+(v1586+int32(8)))))
	if v2028 != 0 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+4))
	v2031 = v2029
	goto L329
L328:
	;
	v2031 = int32(0)
	goto L329
L329:
	;
	v2040 = v2031 + v2022
	F_consider_index_join_outer_rels(m, l0, l1, v1633, v1586+int32(268), v1586+int32(136), v1586+int32(4), v1586+int32(400), v2028, v2040, v1586+int32(408))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L19
	} else {
		goto L330
	}
L330:
	;
	v2046 = v1984 + int32(1)
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+40))
	if v2046 < v2047 {
		v1984 = v2046
		v1993 = v2040
		goto L321
	} else {
		goto L331
	}
L331:
	;
	goto L322
L332:
	;
	goto L269
L333:
	;
	v2114 = F_list_concat(m, v2109, v2112)
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L19
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+404)) = v2114
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v2118 = F_generate_bitmap_or_paths(m, l0, l1, v2094, v2117)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L19
	} else {
		goto L335
	}
L335:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+400))
	v2121 = F_list_concat(m, v2120, v2118)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L19
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+400)) = v2121
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+404))
	if v2124 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	if v2121 == int32(0) {
		goto L263
	} else {
		goto L366
	}
L338:
	;
	v2127 = F_choose_bitmap_and(m, l0, l1, v2124)
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L19
	} else {
		goto L339
	}
L339:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2132 = F_create_bitmap_heap_path(m, l0, l1, v2127, v2129, float64(1), int32(0))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L19
	} else {
		goto L340
	}
L340:
	;
	F_add_path(m, l1, v2132)
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L19
	} else {
		goto L341
	}
L341:
	;
	v2136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v2136 != int32(1) {
		goto L337
	} else {
		goto L342
	}
L342:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v2139 != 0 {
		goto L337
	} else {
		goto L343
	}
L343:
	;
	v2141 = int32(0)
	v2143 = F_compute_bitmap_pages(m, l0, l1, v2127, float64(1), v2141, v2141)
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L19
	} else {
		goto L344
	}
L344:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	if v2147 != int32(-1) {
		v2214 = v2147
		goto L346
	} else {
		goto L347
	}
L345:
	;
	goto L337
L346:
	;
	if v2214 < v2146 {
		goto L360
	} else {
		goto L361
	}
L347:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2150 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2159 = int32(0)
	if base.F64_ge(v2143, float64(0)) == v2159 {
		v2214 = v2159
		goto L346
	} else {
		goto L352
	}
L349:
	;
	if base.F64_ge(v2143, float64(0)) == int32(0) {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	if base.F64_lt(v2143, base.F64_convert_i32_s(v2156)) != 0 {
		goto L345
	} else {
		goto L351
	}
L351:
	;
	goto L348
L352:
	;
	v2164 = int32(1)
	v2166 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	if v2166 <= v2164 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v2169 = v2164
	goto L355
L354:
	;
	v2169 = v2166
	goto L355
L355:
	;
	v2177 = int32(1)
	v2179 = v2169
	goto L356
L356:
	;
	v2199 = v2179 * int32(3)
	if base.F64_ge(v2143, base.F64_convert_i32_u(v2199)) == int32(0) {
		v2214 = v2177
		goto L346
	} else {
		goto L358
	}
L357:
	;
	v2214 = v2205
	goto L346
L358:
	;
	v2205 = v2177 + int32(1)
	if v2199 < int32(715827883) {
		v2177 = v2205
		v2179 = v2199
		goto L356
	} else {
		goto L359
	}
L359:
	;
	goto L357
L360:
	;
	v2236 = v2214
	goto L362
L361:
	;
	v2236 = v2146
	goto L362
L362:
	;
	if v2236 <= int32(0) {
		goto L345
	} else {
		goto L363
	}
L363:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2241 = F_create_bitmap_heap_path(m, l0, l1, v2127, v2239, float64(1), v2236)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L19
	} else {
		goto L364
	}
L364:
	;
	F_add_partial_path(m, l1, v2241)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L19
	} else {
		goto L365
	}
L365:
	;
	goto L345
L366:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2121)+4))
	if v2301 <= int32(0) {
		goto L263
	} else {
		goto L367
	}
L367:
	;
	v2304 = int32(0)
	v2310 = v2304
	v2316 = v2304
	goto L368
L368:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2121)+12))
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2333+v2310<<(uint(int32(2))%32))))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+16))
	if v2338 != 0 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	if v2342 == int32(0) {
		goto L263
	} else {
		goto L375
	}
L370:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2338)+4))
	v2341 = v2339
	goto L372
L371:
	;
	v2341 = int32(0)
	goto L372
L372:
	;
	v2342 = F_list_append_unique(m, v2316, v2341)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L19
	} else {
		goto L373
	}
L373:
	;
	v2345 = v2310 + int32(1)
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2121)+4))
	if v2345 < v2346 {
		v2310 = v2345
		v2316 = v2342
		goto L368
	} else {
		goto L374
	}
L374:
	;
	goto L369
L375:
	;
	v2350 = int32(0)
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2342)+4))
	if v2351 <= v2350 {
		goto L263
	} else {
		goto L376
	}
L376:
	;
	v2366 = v2350
	goto L377
L377:
	;
	v2381 = int32(0)
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+400))
	if v2382 == v2381 {
		v2500 = v2381
		goto L379
	} else {
		goto L380
	}
L378:
	;
	goto L263
L379:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+404))
	v2519 = F_list_concat(m, v2500, v2518)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L19
	} else {
		goto L406
	}
L380:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2382)+4))
	if v2385 <= int32(0) {
		v2500 = v2381
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2342)+12))
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2388+v2366<<(uint(int32(2))%32))))
	v2398 = int32(0)
	v2403 = v2381
	goto L382
L382:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v2382)+12))
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v2421+v2398<<(uint(int32(2))%32))))
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v2425)+16))
	if v2426 != 0 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v2500 = v2486
	goto L379
L384:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2426)+4))
	v2429 = v2427
	goto L386
L385:
	;
	v2429 = int32(0)
	goto L386
L386:
	;
	v2430 = int32(0)
	if v2429 == v2430 {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	if v2483 != 0 {
		goto L401
	} else {
		goto L402
	}
L388:
	;
	v2483 = int32(1)
	goto L387
L389:
	;
	goto L390
L390:
	;
	if v2392 == int32(0) {
		v2474 = v2430
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2483 = v2474
	goto L387
L392:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+4))
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v2392)+4))
	if v2440 < v2439 {
		v2474 = v2430
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v2442 = int32(1)
	if v2439 <= v2442 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2445 = v2442
	goto L396
L395:
	;
	v2445 = v2439
	goto L396
L396:
	;
	v2446 = int32(8)
	v2451 = int32(0)
	goto L397
L397:
	;
	v2458 = v2451 << (uint(int32(2)) % 32)
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2429+v2446+v2458)))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2458+(v2392+v2446))))
	v2465 = v2460 & (v2462 ^ int32(-1))
	v2467 = base.B2i32(v2465 == int32(0))
	if v2465 != 0 {
		v2474 = v2467
		goto L391
	} else {
		goto L399
	}
L398:
	;
	v2474 = v2467
	goto L391
L399:
	;
	v2469 = v2451 + int32(1)
	if v2469 != v2445 {
		v2451 = v2469
		goto L397
	} else {
		goto L400
	}
L400:
	;
	goto L398
L401:
	;
	v2484 = F_lappend(m, v2403, v2425)
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L19
	} else {
		goto L404
	}
L402:
	;
	v2486 = v2403
	goto L403
L403:
	;
	v2488 = v2398 + int32(1)
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2382)+4))
	if v2488 < v2489 {
		v2398 = v2488
		v2403 = v2486
		goto L382
	} else {
		goto L405
	}
L404:
	;
	v2486 = v2484
	goto L403
L405:
	;
	goto L383
L406:
	;
	v2521 = F_choose_bitmap_and(m, l0, l1, v2519)
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L19
	} else {
		goto L407
	}
L407:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v2521)+16))
	if v2523 != 0 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v2523)+4))
	v2526 = v2524
	goto L410
L409:
	;
	v2526 = int32(0)
	goto L410
L410:
	;
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2528 = F_get_loop_count(m, l0, v2527, v2526)
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L19
	} else {
		goto L411
	}
L411:
	;
	v2531 = F_create_bitmap_heap_path(m, l0, l1, v2521, v2526, v2528, int32(0))
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L19
	} else {
		goto L412
	}
L412:
	;
	F_add_path(m, l1, v2531)
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L19
	} else {
		goto L413
	}
L413:
	;
	v2536 = v2366 + int32(1)
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v2342)+4))
	if v2536 < v2537 {
		v2366 = v2536
		goto L377
	} else {
		goto L414
	}
L414:
	;
	goto L378
L415:
	;
	m.T0[v2624].(func(*base.Module, int32, int32, int32, int32))(m, v2596, v2597, v2598, v2599)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L19
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v2597)+4))
	if v2627 != 0 {
		goto L419
	} else {
		goto L420
	}
L418:
	;
	goto L417
L419:
	;
	F_set_cheapest(m, v2597)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L19
	} else {
		goto L435
	}
L420:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2597)+8))
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2596)+52))
	v2630 = int32(0)
	v2637 = base.B2i32(v2628|v2629 == v2630)
	if v2628 == v2630 {
		v2676 = v2637
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v2676 != 0 {
		goto L419
	} else {
		goto L433
	}
L422:
	;
	goto L421
L423:
	;
	if v2629 == int32(0) {
		v2676 = v2637
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2628)+4))
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+4))
	if v2643 != v2644 {
		v2676 = int32(0)
		goto L422
	} else {
		goto L425
	}
L425:
	;
	v2646 = int32(1)
	if v2643 <= v2646 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2649 = v2646
	goto L428
L427:
	;
	v2649 = v2643
	goto L428
L428:
	;
	v2650 = int32(8)
	v2655 = int32(0)
	goto L429
L429:
	;
	v2663 = v2655 << (uint(int32(2)) % 32)
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2628+v2650+v2663)))
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2663+(v2629+v2650))))
	v2668 = base.B2i32(v2665 == v2667)
	if v2667 != v2665 {
		v2676 = v2668
		goto L422
	} else {
		goto L431
	}
L430:
	;
	v2676 = v2668
	goto L422
L431:
	;
	v2671 = v2655 + int32(1)
	if v2671 != v2649 {
		v2655 = v2671
		goto L429
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	F_generate_useful_gather_paths(m, v2596, v2597, int32(0))
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L19
	} else {
		goto L434
	}
L434:
	;
	goto L419
L435:
	;
	m.G0 = v2611 + int32(16)
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
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
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
	var v163 int32
	_ = v163
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
	var v206 int32
	_ = v206
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
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
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v661 int32
	_ = v661
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 float64
	_ = v770
	var v771 float64
	_ = v771
	var v772 float64
	_ = v772
	var v773 float64
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 float64
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
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
	var v871 int32
	_ = v871
	var v873 float64
	_ = v873
	var v875 float64
	_ = v875
	var v898 int32
	_ = v898
	var v906 float64
	_ = v906
	var v907 float64
	_ = v907
	var v908 float64
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 float64
	_ = v961
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1076 int32
	_ = v1076
	var v1085 int32
	_ = v1085
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1284 int32
	_ = v1284
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1313 int64
	_ = v1313
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1394 int64
	_ = v1394
	var v1396 int64
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 float64
	_ = v1408
	var v1410 float64
	_ = v1410
	var v1418 float64
	_ = v1418
	var v1422 float64
	_ = v1422
	var v1424 float64
	_ = v1424
	var v1426 float64
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 float64
	_ = v1487
	var v1490 float64
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1514 float64
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 float64
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 float64
	_ = v1524
	var v1525 int64
	_ = v1525
	var v1532 int32
	_ = v1532
	var v1538 int32
	_ = v1538
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1576 float64
	_ = v1576
	var v1577 float64
	_ = v1577
	var v1579 float64
	_ = v1579
	var v1603 float64
	_ = v1603
	var v1604 float64
	_ = v1604
	var v1608 float64
	_ = v1608
	var v1609 float64
	_ = v1609
	var v1611 float64
	_ = v1611
	var v1613 float64
	_ = v1613
	var v1615 float64
	_ = v1615
	var v1617 float64
	_ = v1617
	var v1618 float64
	_ = v1618
	var v1643 float64
	_ = v1643
	var v1644 float64
	_ = v1644
	var v1645 float64
	_ = v1645
	var v1647 float64
	_ = v1647
	var v1648 float64
	_ = v1648
	var v1649 float64
	_ = v1649
	var v1650 float64
	_ = v1650
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1764 int32
	_ = v1764
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1773 float64
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1814 int32
	_ = v1814
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1911 int32
	_ = v1911
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1942 float64
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1998 int32
	_ = v1998
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
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2056 float64
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2091 int32
	_ = v2091
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 float64
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2125 float64
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2173 int32
	_ = v2173
	var v2196 int32
	_ = v2196
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2247 int32
	_ = v2247
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2329 int32
	_ = v2329
	var v2353 int32
	_ = v2353
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2469 int32
	_ = v2469
	var v2472 float64
	_ = v2472
	var v2473 float64
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2507 int32
	_ = v2507
	var v2512 int32
	_ = v2512
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2562 int32
	_ = v2562
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2610 int32
	_ = v2610
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2639 int32
	_ = v2639
	var v2664 int32
	_ = v2664
	var v2668 int32
	_ = v2668
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
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2711 int32
	_ = v2711
	var v2714 int32
	_ = v2714
	var v2717 int32
	_ = v2717
	var v2724 int32
	_ = v2724
	var v2749 int32
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2784 int32
	_ = v2784
	var v2792 float64
	_ = v2792
	var v2793 float64
	_ = v2793
	var v2794 float64
	_ = v2794
	var v2803 int32
	_ = v2803
	var v2805 float64
	_ = v2805
	var v2809 int32
	_ = v2809
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2852 int32
	_ = v2852
	var v2859 float64
	_ = v2859
	var v2861 float64
	_ = v2861
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2869 int32
	_ = v2869
	var v2871 int32
	_ = v2871
	var v2878 float64
	_ = v2878
	var v2880 float64
	_ = v2880
	var v2884 int32
	_ = v2884
	var v2886 int32
	_ = v2886
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2897 int32
	_ = v2897
	var v2923 int32
	_ = v2923
	var v2930 float64
	_ = v2930
	var v2932 float64
	_ = v2932
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2969 int32
	_ = v2969
	var v2976 int32
	_ = v2976
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3018 int32
	_ = v3018
	var v3020 int32
	_ = v3020
	var v3022 int32
	_ = v3022
	var v3029 int32
	_ = v3029
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
	m.G0 = v3029 + int32(144)
	return
L2:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v61 == int32(1) {
		goto L12
	} else {
		goto L13
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
	v3029 = v31
	goto L1
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v3001 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3000)+32)) = v3001
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v3001
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v3001
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v3015 = F_create_append_path(m, v3001, l1, v3001, v3001, v3001, v3011, v3001, v3001, float64(-1))
	mBase = m.M
	v3016 = m.ExcPending
	if v3016 != 0 {
		goto L4
	} else {
		goto L578
	}
L11:
	;
	if v2784&int32(1) == int32(0) {
		v2976 = v31
		goto L10
	} else {
		goto L551
	}
L12:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v915 {
	case 0:
		goto L202
	case 1:
		goto L201
	default:
		goto L199
	case 3:
		goto L193
	case 4:
		goto L194
	case 5:
		goto L195
	case 6:
		goto L196
	case 7:
		goto L197
	case 8:
		goto L198
	}
L15:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _consts[611])))
	if v67 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
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
		goto L21
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v70 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v71 != int32(112) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74-v75<<(uint(int32(2))%32))))
	if v79 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+217)) = uint8(v80)
	goto L16
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v91 == int32(0) {
		v2976 = v31
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v94 <= int32(0) {
		v2784 = v5
		v2792 = v24
		v2793 = v24
		v2794 = v24
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v106 = v5
	v112 = v5
	v120 = v24
	v121 = v24
	v122 = v24
	goto L24
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v106<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v131 != l2 {
		v898 = v112
		v906 = v120
		v907 = v121
		v908 = v122
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v2784 = v898
	v2792 = v906
	v2793 = v907
	v2794 = v908
	goto L11
L26:
	;
	v912 = v106 + int32(1)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v912 < v913 {
		v106 = v912
		v112 = v898
		v120 = v906
		v121 = v907
		v122 = v908
		goto L24
	} else {
		goto L189
	}
L27:
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
		goto L28
	}
L28:
	;
	v141 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
	if v143 == v141 {
		v163 = v141
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v163 != 0 {
		v898 = v112
		v906 = v120
		v907 = v121
		v908 = v122
		goto L26
	} else {
		goto L39
	}
L30:
	;
	goto L29
L31:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v147 = v146
	goto L32
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if base.Ui32(int32(2)) <= base.Ui32(v151-int32(301)) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v163 = int32(1)
	goto L30
L34:
	;
	if v151 != int32(290) {
		v163 = v141
		goto L30
	} else {
		goto L37
	}
L35:
	;
	v147 = v150 + int32(72)
	goto L32
L36:
	;
	goto L33
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)+72))
	if v158 != 0 {
		v163 = v141
		goto L30
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v165 = F_relation_excluded_by_constraints(m, l0, v139, v138)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v165 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
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
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v190 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	F_add_path(m, v139, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_set_cheapest(m, v139)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v898 = v112
	v906 = v120
	v907 = v121
	v908 = v122
	goto L26
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+212)) = v300
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v327 = F_adjust_appendrel_attrs(m, l0, v323, int32(1), v31+int32(128))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L74
	}
L48:
	;
	v300 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v194 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v196 <= v194 {
		v300 = v194
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v203 = v194
	v206 = v194
	goto L52
L52:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v203<<(uint(int32(2))%32))))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v234 = int32(0)
	if v232 == v234 {
		v275 = v234
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v300 = v288
	goto L47
L54:
	;
	if v275 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L55:
	;
	goto L54
L56:
	;
	if v233 == int32(0) {
		v275 = v234
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v243 < v244 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v246 = v243
	goto L60
L59:
	;
	v246 = v244
	goto L60
L60:
	;
	if v246 <= int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v249 = int32(1)
	goto L63
L62:
	;
	v249 = v246
	goto L63
L63:
	;
	v250 = int32(8)
	v255 = int32(0)
	goto L64
L64:
	;
	v262 = v255 << (uint(int32(2)) % 32)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v233+v250+v262)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+(v232+v250))))
	v267 = v264 & v266
	v269 = base.B2i32(v267 != int32(0))
	if v267 != 0 {
		v275 = v269
		goto L55
	} else {
		goto L66
	}
L65:
	;
	v275 = v269
	goto L55
L66:
	;
	v271 = v255 + int32(1)
	if v271 != v249 {
		v255 = v271
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v284 = F_adjust_appendrel_attrs(m, l0, v231, int32(1), v31+int32(128))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	v288 = v206
	goto L70
L70:
	;
	v290 = v203 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v290 < v291 {
		v203 = v290
		v206 = v288
		goto L52
	} else {
		goto L73
	}
L71:
	;
	v286 = F_lappend(m, v206, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v288 = v286
	goto L70
L73:
	;
	goto L53
L74:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v329)+4)) = v327
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v331 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+216)) = uint8(v723)
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+217)))
	if v725 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L76:
	;
	v335 = int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v336 != 0 {
		v342 = v335
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	v346 = m.G0
	v348 = v346 - int32(16)
	m.G0 = v348
	*(*int32)(unsafe.Add(mBase, uint32(v348)+12)) = v345
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v139)+228))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v353 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L79:
	;
	if v342 == int32(0) {
		goto L75
	} else {
		goto L84
	}
L80:
	;
	goto L79
L81:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v337 != 0 {
		v342 = v335
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v338 != 0 {
		v342 = v335
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v342 = base.B2i32(v339 != int32(0))
	goto L80
L84:
	;
	goto L78
L85:
	;
	if int32(0) <= v410 {
		goto L96
	} else {
		goto L97
	}
L86:
	;
	v410 = base.I32_ctz(v396) | v397<<(uint(int32(5))%32)
	goto L85
L87:
	;
	v410 = int32(-2)
	goto L85
L88:
	;
	v363 = base.I32_div_s(int32(0), int32(32))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v364 <= v363 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v367 = v353 + int32(8)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367+v363<<(uint(int32(2))%32))))
	v374 = v371 & int32(-1)
	if v374 != 0 {
		v396 = v374
		v397 = v363
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v376 = v363 + int32(1)
	if v376 == v364 {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v379 = v376
	goto L92
L92:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v367+v379<<(uint(int32(2))%32))))
	if v386 != 0 {
		v396 = v386
		v397 = v379
		goto L86
	} else {
		goto L94
	}
L93:
	;
	goto L87
L94:
	;
	v388 = v379 + int32(1)
	if v388 != v364 {
		v379 = v388
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v420 = v410
	goto L99
L97:
	;
	goto L98
L98:
	;
	m.G0 = v348 + int32(16)
	goto L75
L99:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v442+v420<<(uint(int32(2))%32))))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+41)))
	if v447 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v605 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L102:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v446)+16))
	if v448 == int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v451 = int32(0)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v452 <= v451 {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v459 = v451
	goto L105
L105:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v483+v459<<(uint(int32(2))%32))))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+12)))
	if v488 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L101
L107:
	;
	v574 = v459 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v574 < v575 {
		v459 = v574
		goto L105
	} else {
		goto L134
	}
L108:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	v490 = int32(0)
	if v489 == v490 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v543 == int32(0) {
		goto L107
	} else {
		goto L123
	}
L110:
	;
	v543 = int32(1)
	goto L109
L111:
	;
	goto L112
L112:
	;
	if v352 == int32(0) {
		v534 = v490
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v543 = v534
	goto L109
L114:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	if v500 < v499 {
		v534 = v490
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v502 = int32(1)
	if v499 <= v502 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v505 = v502
	goto L118
L117:
	;
	v505 = v499
	goto L118
L118:
	;
	v506 = int32(8)
	v511 = int32(0)
	goto L119
L119:
	;
	v518 = v511 << (uint(int32(2)) % 32)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v489+v506+v518)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v518+(v352+v506))))
	v525 = v520 & (v522 ^ int32(-1))
	v527 = base.B2i32(v525 == int32(0))
	if v525 != 0 {
		v534 = v527
		goto L113
	} else {
		goto L121
	}
L120:
	;
	v534 = v527
	goto L113
L121:
	;
	v529 = v511 + int32(1)
	if v529 != v505 {
		v511 = v529
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	if v546 == int32(0) {
		goto L107
	} else {
		goto L124
	}
L124:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v550 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	v563 = F_bms_difference(m, v562, v352)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L131
	}
L126:
	;
	v556 = F_adjust_appendrel_attrs(m, l0, v549, int32(1), v348+int32(12))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v139)+224))
	v559 = F_adjust_appendrel_attrs_multilevel(m, l0, v549, v139, v558)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L130
	}
L129:
	;
	v561 = v556
	goto L125
L130:
	;
	v561 = v559
	goto L125
L131:
	;
	v565 = F_bms_add_members(m, v563, v351)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v487)+16))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v139)+68))
	F_add_child_eq_member(m, l0, v446, v420, v561, v565, v567, v487, v568, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	goto L107
L134:
	;
	goto L106
L135:
	;
	if int32(0) <= v661 {
		v420 = v661
		goto L99
	} else {
		goto L146
	}
L136:
	;
	v661 = base.I32_ctz(v647) | v648<<(uint(int32(5))%32)
	goto L135
L137:
	;
	v661 = int32(-2)
	goto L135
L138:
	;
	v612 = v420 + int32(1)
	v614 = base.I32_div_s(v612, int32(32))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v615 <= v614 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v618 = v605 + int32(8)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v614<<(uint(int32(2))%32))))
	v625 = v622 & (int32(-1) << (uint(v612) % 32))
	if v625 != 0 {
		v647 = v625
		v648 = v614
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v627 = v614 + int32(1)
	if v627 == v615 {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v630 = v627
	goto L142
L142:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v618+v630<<(uint(int32(2))%32))))
	if v637 != 0 {
		v647 = v637
		v648 = v630
		goto L136
	} else {
		goto L144
	}
L143:
	;
	goto L137
L144:
	;
	v639 = v630 + int32(1)
	if v639 != v615 {
		v630 = v639
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	goto L100
L147:
	;
	v728 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+217)) = uint8(v728)
	goto L149
L148:
	;
	goto L149
L149:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730)+82)))
	if v731 != int32(1) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_set_rel_size(m, l0, v139, v134, v138)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L4
	} else {
		goto L154
	}
L151:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v734 != int32(1) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	F_set_rel_consider_parallel(m, l0, v139, v138)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	goto L150
L154:
	;
	v741 = int32(0)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
	if v743 == v741 {
		v763 = v741
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if v763 != 0 {
		v898 = v112
		v906 = v120
		v907 = v121
		v908 = v122
		goto L26
	} else {
		goto L165
	}
L156:
	;
	goto L155
L157:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v743)+12))
	v747 = v746
	goto L158
L158:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	if base.Ui32(int32(2)) <= base.Ui32(v751-int32(301)) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v763 = int32(1)
	goto L156
L160:
	;
	if v751 != int32(290) {
		v763 = v741
		goto L156
	} else {
		goto L163
	}
L161:
	;
	v747 = v750 + int32(72)
	goto L158
L162:
	;
	goto L159
L163:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v750)+72))
	if v758 != 0 {
		v763 = v741
		goto L156
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+26)))
	if v765 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v768 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v768)
	goto L168
L167:
	;
	goto L168
L168:
	;
	v770 = *(*float64)(unsafe.Add(mBase, uint32(v139)+16))
	v771 = base.F64_add(v120, v770)
	v772 = *(*float64)(unsafe.Add(mBase, uint32(v139)+120))
	v773 = base.F64_add(v122, v772)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+32))
	v778 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v775), v770), v121)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v774)+4))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	v787 = int32(0)
	goto L169
L169:
	;
	v811 = int32(0)
	if v781 == v811 {
		v821 = v811
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v822 = int32(1)
	if v779 == int32(0) {
		v898 = v822
		v906 = v771
		v907 = v778
		v908 = v773
		goto L26
	} else {
		goto L174
	}
L172:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v781)+4))
	if v815 <= v787 {
		v821 = int32(0)
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v781)+12))
	v821 = v817 + v787<<(uint(int32(2))%32)
	goto L171
L174:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	if v825 <= v787 {
		v898 = v822
		v906 = v771
		v907 = v778
		v908 = v773
		goto L26
	} else {
		goto L175
	}
L175:
	;
	if v821 == int32(0) {
		v898 = v822
		v906 = v771
		v907 = v778
		v908 = v773
		goto L26
	} else {
		goto L176
	}
L176:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	v832 = v829 + v787<<(uint(int32(2))%32)
	if v832 == int32(0) {
		v898 = v822
		v906 = v771
		v907 = v778
		v908 = v773
		goto L26
	} else {
		goto L177
	}
L177:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	if v836 != int32(6) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v787 = v787 + int32(1)
	goto L169
L179:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v835)+4))
	if v839 != l2 {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v841 = int32(*(*int16)(unsafe.Add(mBase, uint32(v835)+8)))
	v842 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	if v845 != int32(6) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v871 = v89 + (v841-v842)<<(uint(int32(3))%32)
	v873 = *(*float64)(unsafe.Add(mBase, uint32(v139)+16))
	v875 = *(*float64)(unsafe.Add(mBase, uint32(v871)))
	*(*float64)(unsafe.Add(mBase, uint32(v871))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v868), v873), v875)
	goto L178
L182:
	;
	v862 = F_exprType(m, v844)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L186
	}
L183:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v844)+4))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v139)+68))
	if v848 != v849 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v139)+88))
	v852 = int32(*(*int16)(unsafe.Add(mBase, uint32(v844)+8)))
	v853 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139)+80)))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v851+(v852-v853)<<(uint(int32(2))%32))))
	if int32(0) < v858 {
		v868 = v858
		goto L181
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v864 = F_exprTypmod(m, v844)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	v866 = F_get_typavgwidth(m, v862, v864)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v868 = v866
	goto L181
L189:
	;
	goto L25
L190:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	F_pfree(m, v2220)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L4
	} else {
		goto L429
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+184)) = v2173
	v2196 = v2167
	goto L190
L192:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v2167 = v2162
	v2173 = v1284
	goto L191
L193:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2064 != 0 {
		goto L415
	} else {
		goto L416
	}
L194:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(4636737291354636288)
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L4
	} else {
		goto L413
	}
L195:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2036 != 0 {
		goto L406
	} else {
		goto L407
	}
L196:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+92)))
	if v1667 != int32(1) {
		goto L338
	} else {
		goto L339
	}
L197:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1471 != 0 {
		goto L312
	} else {
		goto L313
	}
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(4607182418800017408)
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L4
	} else {
		goto L306
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L4
	} else {
		goto L303
	}
L200:
	;
	v1306 = m.G0
	v1308 = v1306 - int32(32)
	m.G0 = v1308
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(4652007308841189376)
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v1313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1308)+24)) = v1313
	*(*int64)(unsafe.Add(mBase, uint32(v1308)+16)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v1308)+8)) = l0
	v1319 = v1308 + int32(16)
	if v1312 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L201:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = int32(0)
	v976 = F_copyObjectImpl(m, v973)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L4
	} else {
		goto L217
	}
L202:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	switch v916 - int32(102) {
	case 0:
		goto L200
	default:
		goto L203
	case 10:
		goto L204
	}
L203:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v942 != 0 {
		goto L208
	} else {
		goto L209
	}
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v922 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v921)+32)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v922
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v936 = F_create_append_path(m, v922, l1, v922, v922, v922, v932, v922, v922, float64(-1))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	F_add_path(m, l1, v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	v3029 = v31
	goto L1
L208:
	;
	v943 = m.G0
	v945 = v943 - int32(16)
	m.G0 = v945
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	F_check_index_predicates(m, l0, l1)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L4
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	F_check_index_predicates(m, l0, l1)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L4
	} else {
		goto L215
	}
L211:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v947)+4))
	v951 = F_GetTsmRoutine(m, v950)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v947)+8))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	m.T0[v956].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, v953, v945+int32(12), v945)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v945)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+116)) = v959
	v961 = *(*float64)(unsafe.Add(mBase, uint32(v945)))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v961
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L4
	} else {
		goto L214
	}
L214:
	;
	m.G0 = v945 + int32(16)
	v3029 = v31
	goto L1
L215:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	v3029 = v31
	goto L1
L217:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = int64(0)
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v976)+76))
	if v981 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	v986 = v982 + int32(1)
	goto L220
L219:
	;
	v986 = int32(1)
	goto L220
L220:
	;
	v987 = F_palloc0(m, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L4
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v987
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+133)) = uint8(v990)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v992 == int32(0) {
		v2196 = v5
		goto L190
	} else {
		goto L222
	}
L222:
	;
	v997 = F_subquery_is_pushdown_safe(m, v976, v976, v31+int32(128))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	if v997 == int32(0) {
		v2196 = v5
		goto L190
	} else {
		goto L224
	}
L224:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v1001 == int32(0) {
		v2167 = v5
		v2173 = v5
		goto L191
	} else {
		goto L225
	}
L225:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if v1004 <= int32(0) {
		v2167 = v5
		v2173 = v5
		goto L191
	} else {
		goto L226
	}
L226:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+132)))
	v1009 = int32(1)
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+133)))
	v1018 = v5
	v1024 = v5
	goto L227
L227:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1042+v1018<<(uint(int32(2))%32))))
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046)+10)))
	if v1047 == int32(1) {
		goto L233
	} else {
		goto L234
	}
L228:
	;
	goto L192
L229:
	;
	v1303 = v1018 + int32(1)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if v1303 < v1304 {
		v1018 = v1303
		v1024 = v1284
		goto L227
	} else {
		goto L285
	}
L230:
	;
	F_list_free(m, v1058)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L4
	} else {
		goto L283
	}
L231:
	;
	F_subquery_push_qual(m, v976, l3, l2, v1050)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L4
	} else {
		goto L282
	}
L232:
	;
	F_list_free(m, v1058)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L4
	} else {
		goto L281
	}
L233:
	;
	v1233 = F_lappend(m, v1024, v1046)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L4
	} else {
		goto L280
	}
L234:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+4))
	v1051 = F_contain_subplans(m, v1050)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L4
	} else {
		goto L235
	}
L235:
	;
	if v1051 != 0 {
		goto L233
	} else {
		goto L236
	}
L236:
	;
	if v1008&v1009 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1053 = F_contain_volatile_functions(m, v1046)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L4
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	if v1011&v1009 != 0 {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	if v1053 != 0 {
		goto L233
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v1055 = F_contain_leaked_vars(m, v1050)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1058 = F_pull_var_clause(m, v1050, int32(16))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L4
	} else {
		goto L247
	}
L245:
	;
	if v1055 != 0 {
		goto L233
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	if v1058 == int32(0) {
		goto L232
	} else {
		goto L248
	}
L248:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	if v1062 <= int32(0) {
		goto L232
	} else {
		goto L249
	}
L249:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+12))
	v1076 = int32(0)
	v1085 = int32(1)
	goto L250
L250:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1065+v1076<<(uint(int32(2))%32))))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1099)))
	if v1100 != int32(6) {
		goto L230
	} else {
		goto L252
	}
L251:
	;
	F_list_free(m, v1058)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L4
	} else {
		goto L260
	}
L252:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+4))
	if v1103 != l2 {
		goto L230
	} else {
		goto L253
	}
L253:
	;
	v1105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1099)+8)))
	if v1105 == int32(0) {
		goto L230
	} else {
		goto L254
	}
L254:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1105+v1007))))
	if v1109 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	if v1109&int32(23) != 0 {
		goto L230
	} else {
		goto L258
	}
L256:
	;
	v1113 = v1085
	goto L257
L257:
	;
	v1115 = v1076 + int32(1)
	if v1062 != v1115 {
		v1076 = v1115
		v1085 = v1113
		goto L250
	} else {
		goto L259
	}
L258:
	;
	v1113 = int32(2)
	goto L257
L259:
	;
	goto L251
L260:
	;
	if v1113 == int32(1) {
		goto L231
	} else {
		goto L261
	}
L261:
	;
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+37)))
	if v1121 != int32(1) {
		goto L233
	} else {
		goto L262
	}
L262:
	;
	v1124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+140)) = uint8(v1124)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1050)))
	if v1126 != int32(17) {
		goto L233
	} else {
		goto L263
	}
L263:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+28))
	if v1129 == int32(0) {
		goto L233
	} else {
		goto L264
	}
L264:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1132 != int32(2) {
		goto L233
	} else {
		goto L265
	}
L265:
	;
	F_set_opfuncid(m, v1050)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+8))
	v1138 = F_func_strict(m, v1137)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L4
	} else {
		goto L267
	}
L267:
	;
	if v1138 == int32(0) {
		goto L233
	} else {
		goto L268
	}
L268:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+28))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+12))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1144)))
	if v1145 != int32(6) {
		v1170 = v1143
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+140)))
	if v1200&int32(1) == int32(0) {
		v1284 = v1024
		goto L229
	} else {
		goto L279
	}
L270:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+4))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1172)))
	if v1173 != int32(6) {
		goto L233
	} else {
		goto L275
	}
L271:
	;
	v1148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1144)+8)))
	if v1148 <= int32(0) {
		v1170 = v1143
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v976)+76))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1148<<(uint(int32(2))%32)-int32(4))))
	v1159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1158)+8)))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+4))
	v1166 = F_find_window_run_conditions(m, v976, v1159, v1160, v1050, int32(1), v31+int32(140), v31+int32(124))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	if v1166 != 0 {
		goto L269
	} else {
		goto L274
	}
L274:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+28))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+12))
	v1170 = v1169
	goto L270
L275:
	;
	v1176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1172)+8)))
	if v1176 <= int32(0) {
		goto L233
	} else {
		goto L276
	}
L276:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v976)+76))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+12))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1180+v1176<<(uint(int32(2))%32)-int32(4))))
	v1187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1186)+8)))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	v1194 = F_find_window_run_conditions(m, v976, v1187, v1188, v1050, int32(0), v31+int32(140), v31+int32(124))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	if v1194 == int32(0) {
		goto L233
	} else {
		goto L278
	}
L278:
	;
	goto L269
L279:
	;
	goto L233
L280:
	;
	v1284 = v1233
	goto L229
L281:
	;
	goto L231
L282:
	;
	v1284 = v1024
	goto L229
L283:
	;
	v1272 = F_lappend(m, v1024, v1046)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	v1284 = v1272
	goto L229
L285:
	;
	goto L228
L286:
	;
	v1394 = *(*int64)(unsafe.Add(mBase, uint32(v1319)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+192)) = v1394
	v1396 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+200)) = v1396
	F_set_rel_width(m, l0, l1)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L4
	} else {
		goto L293
	}
L287:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+4))
	if v1322 <= int32(0) {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v1329 = v5
	goto L289
L289:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+12))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1353+v1329<<(uint(int32(2))%32))))
	v1360 = F_cost_qual_eval_walker(m, v1357, v1308+int32(8))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L4
	} else {
		goto L291
	}
L290:
	;
	goto L286
L291:
	;
	v1363 = v1329 + int32(1)
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+4))
	if v1363 < v1364 {
		v1329 = v1363
		goto L289
	} else {
		goto L292
	}
L292:
	;
	goto L290
L293:
	;
	m.G0 = v1308 + int32(32)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	m.T0[v1405].(func(*base.Module, int32, int32, int32))(m, l0, l1, v1403)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L4
	} else {
		goto L294
	}
L294:
	;
	v1408 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v1410 = float64(1e+100)
	if base.F64_gt(v1408, v1410) != 0 {
		v1422 = v1410
		goto L296
	} else {
		goto L297
	}
L295:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v1422
	v1424 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	if base.F64_lt(v1422, v1424) != 0 {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	goto L295
L297:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1408)&int64(9223372036854775807)) {
		v1422 = v1410
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1418 = float64(1)
	if base.F64_le(v1408, v1418) != 0 {
		v1422 = v1418
		goto L296
	} else {
		goto L299
	}
L299:
	;
	v1422 = base.F64_nearest(v1408)
	goto L296
L300:
	;
	v1426 = v1424
	goto L302
L301:
	;
	v1426 = v1422
	goto L302
L302:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v1426
	v3029 = v31
	goto L1
L303:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1432
	F_errmsg_internal(m, int32(488444), v31)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L4
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(495231), int32(453), int32(341680))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L4
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1448 = F_palloc0(m, int32(72))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L4
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1448))) = int64(1421634175255)
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+12)) = v1453
	v1455 = F_get_baserel_parampathinfo(m, l0, l1, v1446)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L4
	} else {
		goto L308
	}
L308:
	;
	v1457 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1448)+20)) = uint8(v1457)
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+16)) = v1455
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+64)) = v1457
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+24)) = v1457
	*(*uint8)(unsafe.Add(mBase, uint32(v1448)+21)) = uint8(v1460)
	F_cost_resultscan(m, v1448, l0, l1, v1455)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L4
	} else {
		goto L309
	}
L309:
	;
	F_add_path(m, l1, v1448)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L4
	} else {
		goto L310
	}
L310:
	;
	v3029 = v31
	goto L1
L311:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1485)))
	v1487 = *(*float64)(unsafe.Add(mBase, uint32(v1486)+112))
	if base.F64_lt(v1487, float64(0)) != 0 {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1485 = v1471 + v1472<<(uint(int32(2))%32)
	goto L311
L313:
	;
	goto L314
L314:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+52))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+12))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1485 = v1478 + v1479<<(uint(int32(2))%32) - int32(4)
	goto L311
L315:
	;
	v1490 = float64(1000)
	goto L317
L316:
	;
	v1490 = v1487
	goto L317
L317:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v1490
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L4
	} else {
		goto L318
	}
L318:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1496 = F_palloc0(m, int32(72))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L4
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1496))) = int64(1511828488471)
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+12)) = v1501
	v1503 = F_get_baserel_parampathinfo(m, l0, l1, v1494)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L4
	} else {
		goto L320
	}
L320:
	;
	v1505 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1496)+20)) = uint8(v1505)
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+16)) = v1503
	v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+64)) = v1505
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+24)) = v1505
	*(*uint8)(unsafe.Add(mBase, uint32(v1496)+21)) = uint8(v1508)
	v1514 = float64(0)
	v1515 = m.G0
	v1517 = v1515 - int32(32)
	m.G0 = v1517
	if v1503 != 0 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1648 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v1649 = float64(0)
	v1650 = base.F64_add(v1647, v1649)
	*(*float64)(unsafe.Add(mBase, uint32(v1496)+48)) = v1650
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+40)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v1496)+56)) = base.F64_add(v1650, base.F64_add(base.F64_mul(v1648, base.F64_add(v1644, base.F64_add(v1643, v1645))), v1649))
	m.G0 = v1517 + int32(32)
	F_add_path(m, l1, v1496)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L4
	} else {
		goto L332
	}
L322:
	;
	v1519 = *(*float64)(unsafe.Add(mBase, uint32(v1503)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1496)+32)) = v1519
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+16))
	v1522 = int32(0)
	v1524 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v1525 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1517)+24)) = v1525
	*(*int64)(unsafe.Add(mBase, uint32(v1517)+16)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+8)) = l0
	if v1521 == v1522 {
		v1603 = v1514
		v1604 = v24
		v1608 = v1524
		goto L325
	} else {
		goto L326
	}
L323:
	;
	goto L324
L324:
	;
	v1613 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1496)+32)) = v1613
	v1615 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v1617 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v1618 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v1643 = v1615
	v1644 = v1617
	v1645 = v1617
	v1647 = v1618
	goto L321
L325:
	;
	v1609 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v1611 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v1643 = base.F64_add(v1604, v1609)
	v1644 = v1524
	v1645 = v1608
	v1647 = base.F64_add(v1603, v1611)
	goto L321
L326:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+4))
	if v1532 <= int32(0) {
		v1603 = v1514
		v1604 = v24
		v1608 = v1524
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1538 = v1522
	goto L328
L328:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+12))
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1563+v1538<<(uint(int32(2))%32))))
	v1570 = F_cost_qual_eval_walker(m, v1567, v1517+int32(8))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L4
	} else {
		goto L330
	}
L329:
	;
	v1576 = *(*float64)(unsafe.Add(mBase, uint32(v1517)+24))
	v1577 = *(*float64)(unsafe.Add(mBase, uint32(v1517)+16))
	v1579 = *(*float64)(unsafe.Add(mBase, _consts[612]))
	v1603 = v1577
	v1604 = v1576
	v1608 = v1579
	goto L325
L330:
	;
	v1573 = v1538 + int32(1)
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+4))
	if v1573 < v1574 {
		v1538 = v1573
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	v3029 = v31
	goto L1
L333:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L4
	} else {
		goto L402
	}
L334:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L4
	} else {
		goto L399
	}
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L4
	} else {
		goto L396
	}
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L4
	} else {
		goto L393
	}
L337:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+4))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+48))
	if v1801 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L338:
	;
	v1674 = v1666
	v1675 = l0
	goto L341
L339:
	;
	goto L340
L340:
	;
	if v1666 == int32(0) {
		goto L336
	} else {
		goto L348
	}
L341:
	;
	if v1674 == int32(0) {
		goto L337
	} else {
		goto L343
	}
L342:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L4
	} else {
		goto L345
	}
L343:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+16))
	if v1702 != 0 {
		v1674 = v1674 - int32(1)
		v1675 = v1702
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1707
	F_errmsg_internal(m, int32(725713), v31+int32(112))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	F_errfinish(m, int32(495231), int32(2929), int32(74005))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	v1725 = v1666
	v1726 = l0
	goto L350
L349:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+348))
	if v1770 == int32(0) {
		goto L335
	} else {
		goto L357
	}
L350:
	;
	v1750 = v1725 - int32(1)
	if v1750 == int32(0) {
		goto L349
	} else {
		goto L352
	}
L351:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L4
	} else {
		goto L354
	}
L352:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+16))
	if v1753 != 0 {
		v1725 = v1750
		v1726 = v1753
		goto L350
	} else {
		goto L353
	}
L353:
	;
	goto L351
L354:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v1758
	F_errmsg_internal(m, int32(725713), v31+int32(48))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L4
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(495231), int32(3061), int32(74022))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L4
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	v1773 = *(*float64)(unsafe.Add(mBase, uint32(v1770)+32))
	F_set_cte_size_estimates(m, l0, l1, v1773)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L4
	} else {
		goto L358
	}
L358:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1778 = F_palloc0(m, int32(72))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1778))) = int64(1516123455767)
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+12)) = v1783
	v1785 = F_get_baserel_parampathinfo(m, l0, l1, v1776)
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	v1787 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1778)+20)) = uint8(v1787)
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+16)) = v1785
	v1790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+64)) = v1787
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+24)) = v1787
	*(*uint8)(unsafe.Add(mBase, uint32(v1778)+21)) = uint8(v1790)
	F_cost_ctescan(m, v1778, l0, l1, v1785)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	F_add_path(m, l1, v1778)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L4
	} else {
		goto L362
	}
L362:
	;
	v3029 = v31
	goto L1
L363:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+76))
	if v1917 != 0 {
		goto L382
	} else {
		goto L383
	}
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L4
	} else {
		goto L379
	}
L365:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1801)+4))
	if v1804 <= int32(0) {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1801)+12))
	v1814 = int32(0)
	goto L367
L367:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1808+v1814<<(uint(int32(2))%32))))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+4))
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1807))))
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842))))
	if v1846 == int32(0) {
		v1865 = v1845
		v1866 = v1846
		goto L370
	} else {
		goto L371
	}
L368:
	;
	goto L364
L369:
	;
	if v1866-v1865 == int32(0) {
		goto L363
	} else {
		goto L377
	}
L370:
	;
	goto L369
L371:
	;
	if v1845 != v1846 {
		v1865 = v1845
		v1866 = v1846
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1850 = v1842
	v1851 = v1807
	goto L373
L373:
	;
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1851)+1)))
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1850)+1)))
	if v1855 == int32(0) {
		v1865 = v1854
		v1866 = v1855
		goto L370
	} else {
		goto L375
	}
L374:
	;
	v1865 = v1854
	v1866 = v1855
	goto L370
L375:
	;
	v1858 = int32(1)
	if v1854 == v1855 {
		v1850 = v1850 + v1858
		v1851 = v1851 + v1858
		goto L373
	} else {
		goto L376
	}
L376:
	;
	goto L374
L377:
	;
	v1871 = v1814 + int32(1)
	if v1804 != v1871 {
		v1814 = v1871
		goto L367
	} else {
		goto L378
	}
L378:
	;
	goto L368
L379:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1905
	F_errmsg_internal(m, int32(725872), v31-int32(-64))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L4
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(495231), int32(2947), int32(74005))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L4
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+4))
	v1920 = v1918
	goto L384
L383:
	;
	v1920 = int32(0)
	goto L384
L384:
	;
	if v1920 <= v1814 {
		goto L334
	} else {
		goto L385
	}
L385:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+12))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1922+v1814<<(uint(int32(2))%32))))
	if v1926 <= int32(0) {
		goto L333
	} else {
		goto L386
	}
L386:
	;
	v1932 = v1926<<(uint(int32(2))%32) - int32(4)
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1933)+12))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+12))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1932+v1935)))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1933)+8))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+12))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1939+v1932)))
	v1942 = *(*float64)(unsafe.Add(mBase, uint32(v1941)+24))
	F_set_cte_size_estimates(m, l0, l1, v1942)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1937)+64))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1941)+44))
	v1947 = F_convert_subquery_pathkeys(m, l0, l1, v1945, v1946)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1951 = F_palloc0(m, int32(72))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1951))) = int64(1507533521175)
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+12)) = v1956
	v1958 = F_get_baserel_parampathinfo(m, l0, l1, v1949)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	v1960 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+20)) = uint8(v1960)
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+16)) = v1958
	v1963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+64)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+24)) = v1960
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+21)) = uint8(v1963)
	F_cost_ctescan(m, v1951, l0, l1, v1958)
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	F_add_path(m, l1, v1951)
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	v3029 = v31
	goto L1
L393:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1976
	F_errmsg_internal(m, int32(725713), v31+int32(16))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L4
	} else {
		goto L394
	}
L394:
	;
	F_errfinish(m, int32(495231), int32(3054), int32(74022))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L4
	} else {
		goto L395
	}
L395:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L396:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1992
	F_errmsg_internal(m, int32(725772), v31+int32(32))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L4
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(495231), int32(3065), int32(74022))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L4
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L399:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v2008
	F_errmsg_internal(m, int32(725739), v31+int32(96))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L4
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(495231), int32(2949), int32(74005))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L4
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v2024
	F_errmsg_internal(m, int32(725805), v31+int32(80))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L4
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(495231), int32(2952), int32(74005))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L4
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2050)))
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2051)+80))
	if v2052 != 0 {
		goto L409
	} else {
		goto L410
	}
L406:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2050 = v2036 + v2037<<(uint(int32(2))%32)
	goto L405
L407:
	;
	goto L408
L408:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+52))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+12))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2050 = v2043 + v2044<<(uint(int32(2))%32) - int32(4)
	goto L405
L409:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2052)+4))
	v2056 = base.F64_convert_i32_s(v2053)
	goto L411
L410:
	;
	v2056 = float64(0)
	goto L411
L411:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v2056
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L4
	} else {
		goto L412
	}
L412:
	;
	v3029 = v31
	goto L1
L413:
	;
	v3029 = v31
	goto L1
L414:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2078)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+68))
	if v2082 == int32(0) {
		goto L418
	} else {
		goto L419
	}
L415:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2078 = v2064 + v2065<<(uint(int32(2))%32)
	goto L414
L416:
	;
	goto L417
L417:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+52))
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2070)+12))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2078 = v2071 + v2072<<(uint(int32(2))%32) - int32(4)
	goto L414
L418:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L4
	} else {
		goto L428
	}
L419:
	;
	v2085 = int32(0)
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2082)+4))
	if v2086 <= v2085 {
		goto L418
	} else {
		goto L420
	}
L420:
	;
	v2091 = v2085
	goto L421
L421:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2082)+12))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2117+v2091<<(uint(int32(2))%32))))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2121)+4))
	v2123 = F_expression_returns_set_rows(m, l0, v2122)
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L4
	} else {
		goto L423
	}
L422:
	;
	goto L418
L423:
	;
	v2125 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	if base.F64_gt(v2123, v2125) != 0 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v2123
	goto L426
L425:
	;
	goto L426
L426:
	;
	v2129 = v2091 + int32(1)
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2082)+4))
	if v2129 < v2130 {
		v2091 = v2129
		goto L421
	} else {
		goto L427
	}
L427:
	;
	goto L422
L428:
	;
	v3029 = v31
	goto L1
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v2196
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v976)+144))
	if v2224 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972)+36)))
	if v2416 != 0 {
		v2473 = v24
		goto L467
	} else {
		goto L468
	}
L431:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v976)+120))
	if v2225 != 0 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v2226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+40)))
	if v2226 != int32(1) {
		goto L430
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2229)+4))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v2230, v2231, v31+int32(140))
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L4
	} else {
		goto L436
	}
L435:
	;
	goto L434
L436:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v2236 == int32(0) {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v2316 = F_bms_is_member(m, int32(7), v2315)
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L4
	} else {
		goto L444
	}
L438:
	;
	v2239 = int32(0)
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+4))
	if v2240 <= v2239 {
		goto L437
	} else {
		goto L439
	}
L439:
	;
	v2247 = v2239
	goto L440
L440:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+12))
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2271+v2247<<(uint(int32(2))%32))))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2275)+4))
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v2276, v2277, v31+int32(140))
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L4
	} else {
		goto L442
	}
L441:
	;
	goto L437
L442:
	;
	v2283 = v2247 + int32(1)
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+4))
	if v2283 < v2284 {
		v2247 = v2283
		goto L440
	} else {
		goto L443
	}
L443:
	;
	goto L441
L444:
	;
	if v2316 != 0 {
		goto L430
	} else {
		goto L445
	}
L445:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v976)+76))
	if v2318 == int32(0) {
		goto L430
	} else {
		goto L446
	}
L446:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2318)+4))
	if v2321 <= int32(0) {
		goto L430
	} else {
		goto L447
	}
L447:
	;
	v2329 = int32(0)
	goto L448
L448:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v2318)+12))
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2353+v2329<<(uint(int32(2))%32))))
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+16))
	if v2358 != 0 {
		goto L450
	} else {
		goto L451
	}
L449:
	;
	goto L430
L450:
	;
	v2385 = v2329 + int32(1)
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2318)+4))
	if v2385 < v2386 {
		v2329 = v2385
		goto L448
	} else {
		goto L466
	}
L451:
	;
	v2359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+26)))
	if v2359 != 0 {
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+4))
	v2361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2357)+8)))
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v2365 = F_bms_is_member(m, v2361+int32(7), v2364)
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L4
	} else {
		goto L453
	}
L453:
	;
	if v2365 != 0 {
		goto L450
	} else {
		goto L454
	}
L454:
	;
	v2367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+38)))
	if v2367 == int32(1) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2370 = F_expression_returns_set(m, v2360)
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L4
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	v2372 = F_contain_volatile_functions(m, v2360)
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L4
	} else {
		goto L460
	}
L458:
	;
	if v2370 != 0 {
		goto L450
	} else {
		goto L459
	}
L459:
	;
	goto L457
L460:
	;
	if v2372 != 0 {
		goto L450
	} else {
		goto L461
	}
L461:
	;
	v2374 = F_exprType(m, v2360)
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L4
	} else {
		goto L462
	}
L462:
	;
	v2376 = F_exprTypmod(m, v2360)
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L4
	} else {
		goto L463
	}
L463:
	;
	v2378 = F_exprCollation(m, v2360)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L4
	} else {
		goto L464
	}
L464:
	;
	v2380 = F_makeNullConst(m, v2374, v2376, v2378)
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L4
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2357)+4)) = v2380
	goto L450
L466:
	;
	goto L449
L467:
	;
	v2474 = int32(0)
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2478 = F_subquery_planner(m, v2475, v976, l0, v2474, v2473, v2474)
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L4
	} else {
		goto L491
	}
L468:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v972)+100))
	if v2417 != 0 {
		v2473 = v24
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v972)+108))
	if v2418 != 0 {
		v2473 = v24
		goto L467
	} else {
		goto L470
	}
L470:
	;
	v2419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+318)))
	if v2419 != 0 {
		v2473 = v24
		goto L467
	} else {
		goto L471
	}
L471:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v972)+120))
	if v2420 != 0 {
		v2473 = v24
		goto L467
	} else {
		goto L472
	}
L472:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v972)+124))
	if v2421 != 0 {
		v2473 = v24
		goto L467
	} else {
		goto L473
	}
L473:
	;
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2422 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	if v2469 == int32(2) {
		v2473 = v24
		goto L467
	} else {
		goto L490
	}
L475:
	;
	v2469 = int32(0)
	goto L474
L476:
	;
	goto L477
L477:
	;
	v2431 = int32(1)
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+4))
	if v2432 <= v2431 {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2435 = v2431
	goto L480
L479:
	;
	v2435 = v2432
	goto L480
L480:
	;
	v2438 = int32(0)
	v2440 = v2438
	v2441 = v2438
	goto L481
L481:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v2422+int32(8)+v2440<<(uint(int32(2))%32))))
	if v2449 != 0 {
		goto L484
	} else {
		goto L485
	}
L482:
	;
	v2469 = v2462
	goto L474
L483:
	;
	goto L482
L484:
	;
	v2450 = int32(2)
	if v2441 != 0 {
		v2462 = v2450
		goto L483
	} else {
		goto L487
	}
L485:
	;
	v2455 = v2441
	goto L486
L486:
	;
	v2458 = v2440 + int32(1)
	if v2458 != v2435 {
		v2440 = v2458
		v2441 = v2455
		goto L481
	} else {
		goto L489
	}
L487:
	;
	v2451 = int32(1)
	if base.Ui32(v2451) < base.Ui32(base.I32_popcnt(v2449)) {
		v2462 = v2450
		goto L483
	} else {
		goto L488
	}
L488:
	;
	v2455 = v2451
	goto L486
L489:
	;
	v2462 = v2455
	goto L483
L490:
	;
	v2472 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	v2473 = v2472
	goto L467
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v2478
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+144)) = v2481
	v2483 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2483
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v2488 = F_fetch_upper_rel(m, v2485, int32(7), v2483)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L4
	} else {
		goto L492
	}
L492:
	;
	v2490 = int32(0)
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v2488)+32))
	if v2492 == v2490 {
		v2512 = v2490
		goto L494
	} else {
		goto L495
	}
L493:
	;
	if v2512 != 0 {
		goto L503
	} else {
		goto L504
	}
L494:
	;
	goto L493
L495:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2492)+12))
	v2496 = v2495
	goto L496
L496:
	;
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2496)))
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(v2499)))
	if base.Ui32(int32(2)) <= base.Ui32(v2500-int32(301)) {
		goto L498
	} else {
		goto L499
	}
L497:
	;
	v2512 = int32(1)
	goto L494
L498:
	;
	if v2500 != int32(290) {
		v2512 = v2490
		goto L494
	} else {
		goto L501
	}
L499:
	;
	v2496 = v2499 + int32(72)
	goto L496
L500:
	;
	goto L497
L501:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2499)+72))
	if v2507 != 0 {
		v2512 = v2490
		goto L494
	} else {
		goto L502
	}
L502:
	;
	goto L500
L503:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2517 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2516)+32)) = v2517
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v2517
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v2517
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v2531 = F_create_append_path(m, v2517, l1, v2517, v2517, v2517, v2527, v2517, v2517, float64(-1))
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L4
	} else {
		goto L506
	}
L504:
	;
	goto L505
L505:
	;
	F_set_subquery_size_estimates(m, l0, l1)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L4
	} else {
		goto L509
	}
L506:
	;
	F_add_path(m, l1, v2531)
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L4
	} else {
		goto L507
	}
L507:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L4
	} else {
		goto L508
	}
L508:
	;
	v3029 = v31
	goto L1
L509:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2539)+4))
	if v2540 != 0 {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(v2540)+4))
	v2542 = v2541
	goto L512
L511:
	;
	v2542 = v2474
	goto L512
L512:
	;
	v2544 = *(*int32)(unsafe.Add(mBase, uint32(v976)+76))
	if v2544 != 0 {
		goto L514
	} else {
		goto L515
	}
L513:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2488)+32))
	if v2629 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L514:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(v2544)+4))
	v2547 = v2545
	goto L516
L515:
	;
	v2547 = int32(0)
	goto L516
L516:
	;
	if v2547 != v2542 {
		v2610 = int32(0)
		goto L513
	} else {
		goto L517
	}
L517:
	;
	if v2540 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2610 = int32(1)
	goto L513
L519:
	;
	goto L520
L520:
	;
	v2552 = int32(0)
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(v2540)+4))
	if v2552 < v2553 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v2557 = v2553
	goto L523
L522:
	;
	v2557 = v2552
	goto L523
L523:
	;
	v2562 = v2552
	goto L524
L524:
	;
	v2586 = base.B2i32(v2562 == v2557)
	if v2562 == v2557 {
		v2610 = v2586
		goto L513
	} else {
		goto L526
	}
L525:
	;
	v2610 = v2586
	goto L513
L526:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2540)+12))
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2587+v2562<<(uint(int32(2))%32))))
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v2591)))
	if v2592 != int32(6) {
		v2610 = v2586
		goto L513
	} else {
		goto L527
	}
L527:
	;
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(v2591)+4))
	if v2595 != l2 {
		v2610 = v2586
		goto L513
	} else {
		goto L528
	}
L528:
	;
	v2599 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2591)+8)))
	if v2562+int32(1) == v2599 {
		v2562 = v2599
		goto L524
	} else {
		goto L529
	}
L529:
	;
	goto L525
L530:
	;
	v2711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v2711 != int32(1) {
		v3029 = v31
		goto L1
	} else {
		goto L540
	}
L531:
	;
	v2632 = int32(0)
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+4))
	if v2633 <= v2632 {
		goto L530
	} else {
		goto L532
	}
L532:
	;
	v2639 = v2632
	goto L533
L533:
	;
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+12))
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2664+v2639<<(uint(int32(2))%32))))
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v2668)+64))
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2668)+12))
	v2671 = F_make_tlist_from_pathtarget(m, v2670)
	mBase = m.M
	v2672 = m.ExcPending
	if v2672 != 0 {
		goto L4
	} else {
		goto L535
	}
L534:
	;
	goto L530
L535:
	;
	v2673 = F_convert_subquery_pathkeys(m, l0, l1, v2669, v2671)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L4
	} else {
		goto L536
	}
L536:
	;
	v2675 = F_create_subqueryscan_path(m, l0, l1, v2668, v2610, v2673, v978)
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L4
	} else {
		goto L537
	}
L537:
	;
	F_add_path(m, l1, v2675)
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L4
	} else {
		goto L538
	}
L538:
	;
	v2680 = v2639 + int32(1)
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+4))
	if v2680 < v2681 {
		v2639 = v2680
		goto L533
	} else {
		goto L539
	}
L539:
	;
	goto L534
L540:
	;
	if v978 != 0 {
		v3029 = v31
		goto L1
	} else {
		goto L541
	}
L541:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v2488)+40))
	if v2714 == int32(0) {
		v3029 = v31
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2714)+4))
	if v2717 <= int32(0) {
		v3029 = v31
		goto L1
	} else {
		goto L543
	}
L543:
	;
	v2724 = int32(0)
	goto L544
L544:
	;
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2714)+12))
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v2749+v2724<<(uint(int32(2))%32))))
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v2753)+64))
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v2753)+12))
	v2756 = F_make_tlist_from_pathtarget(m, v2755)
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L4
	} else {
		goto L546
	}
L545:
	;
	v3029 = v31
	goto L1
L546:
	;
	v2758 = F_convert_subquery_pathkeys(m, l0, l1, v2754, v2756)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L4
	} else {
		goto L547
	}
L547:
	;
	v2761 = F_create_subqueryscan_path(m, l0, l1, v2753, v2610, v2758, int32(0))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L4
	} else {
		goto L548
	}
L548:
	;
	F_add_partial_path(m, l1, v2761)
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L4
	} else {
		goto L549
	}
L549:
	;
	v2766 = v2724 + int32(1)
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2714)+4))
	if v2766 < v2767 {
		v2724 = v2766
		goto L544
	} else {
		goto L550
	}
L550:
	;
	goto L545
L551:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v2792
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v2794
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v2805 = base.F64_nearest(base.F64_div(v2793, v2792))
	if base.F64_lt(base.F64_abs(v2805), float64(2.147483648e+09)) != 0 {
		goto L553
	} else {
		goto L554
	}
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2803)+32)) = v2811
	v2813 = int32(0)
	if v84 < v2813 {
		goto L556
	} else {
		goto L557
	}
L553:
	;
	v2809 = base.I32_trunc_f64_s(v2805)
	v2811 = v2809
	goto L552
L554:
	;
	goto L555
L555:
	;
	v2811 = int32(-2147483648)
	goto L552
L556:
	;
	F_pfree(m, v89)
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L4
	} else {
		goto L577
	}
L557:
	;
	v2816 = int32(1)
	v2817 = v84 + v2816
	if v83 != v82 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v2827 = int32(0)
	v2828 = v2813
	goto L561
L559:
	;
	v2897 = v2813
	goto L560
L560:
	;
	if v2817&v2816 == int32(0) {
		goto L556
	} else {
		goto L572
	}
L561:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v2859 = *(*float64)(unsafe.Add(mBase, uint32(v89+v2828<<(uint(int32(3))%32))))
	v2861 = base.F64_nearest(base.F64_div(v2859, v2792))
	if base.F64_lt(base.F64_abs(v2861), float64(2.147483648e+09)) != 0 {
		goto L564
	} else {
		goto L565
	}
L562:
	;
	v2897 = v2889
	goto L560
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2852+v2828<<(uint(int32(2))%32)))) = v2867
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v2871 = v2828 | int32(1)
	v2878 = *(*float64)(unsafe.Add(mBase, uint32(v89+v2871<<(uint(int32(3))%32))))
	v2880 = base.F64_nearest(base.F64_div(v2878, v2792))
	if base.F64_lt(base.F64_abs(v2880), float64(2.147483648e+09)) != 0 {
		goto L568
	} else {
		goto L569
	}
L564:
	;
	v2865 = base.I32_trunc_f64_s(v2861)
	v2867 = v2865
	goto L563
L565:
	;
	goto L566
L566:
	;
	v2867 = int32(-2147483648)
	goto L563
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2869+v2871<<(uint(int32(2))%32)))) = v2886
	v2888 = int32(2)
	v2889 = v2828 + v2888
	v2891 = v2827 + v2888
	if v2891 != v2817&int32(-2) {
		v2827 = v2891
		v2828 = v2889
		goto L561
	} else {
		goto L571
	}
L568:
	;
	v2884 = base.I32_trunc_f64_s(v2880)
	v2886 = v2884
	goto L567
L569:
	;
	goto L570
L570:
	;
	v2886 = int32(-2147483648)
	goto L567
L571:
	;
	goto L562
L572:
	;
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v2930 = *(*float64)(unsafe.Add(mBase, uint32(v89+v2897<<(uint(int32(3))%32))))
	v2932 = base.F64_nearest(base.F64_div(v2930, v2792))
	if base.F64_lt(base.F64_abs(v2932), float64(2.147483648e+09)) != 0 {
		goto L574
	} else {
		goto L575
	}
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2923+v2897<<(uint(int32(2))%32)))) = v2938
	goto L556
L574:
	;
	v2936 = base.I32_trunc_f64_s(v2932)
	v2938 = v2936
	goto L573
L575:
	;
	goto L576
L576:
	;
	v2938 = int32(-2147483648)
	goto L573
L577:
	;
	v3029 = v31
	goto L1
L578:
	;
	F_add_path(m, l1, v3015)
	mBase = m.M
	v3018 = m.ExcPending
	if v3018 != 0 {
		goto L4
	} else {
		goto L579
	}
L579:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v3020 = m.ExcPending
	if v3020 != 0 {
		goto L4
	} else {
		goto L580
	}
L580:
	;
	F_pfree(m, v89)
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L4
	} else {
		goto L581
	}
L581:
	;
	v3029 = v2976
	goto L1
}
