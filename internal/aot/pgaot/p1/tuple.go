package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v9 = F_palloc(m, l0*int32(116)+int32(20))
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
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(-4294965047)
	if int32(0) < l0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v9
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v28 = int32(100)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+v23<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v9+v24<<(uint(int32(4))%32)+v23*v28+int32(20), v36, v28)
	F_populate_compact_attribute(m, v9, v23)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v42 = v23 + int32(1)
	if v42 != l0 {
		v23 = v42
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_FreeTupleDesc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_pfree(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L37
	}
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v11 = v7
	goto L7
L5:
	;
	v28 = v6
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v15 = v11 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8+v15<<(uint(int32(3))%32))+4))
	F_pfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_pfree(m, v8)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	if base.Ui32(int32(1)) < base.Ui32(v11) {
		v11 = v15
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v26
	goto L6
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = v33 - int32(1)
	if int32(0) <= v35 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v73 = v28
	goto L15
L15:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+14)))
	if v77 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v39 = v35
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v32)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L26
	}
L19:
	;
	v45 = v32 + v39<<(uint(int32(3))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	if int32(0) < v39 {
		v39 = v39 - int32(1)
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v49<<(uint(int32(4))%32)+v39*int32(100))+102)))
	if v56 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	goto L20
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v73 = v71
	goto L15
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v81 = v77
	goto L30
L28:
	;
	v105 = v73
	goto L29
L29:
	;
	F_pfree(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L36
	}
L30:
	;
	v85 = v81 - int32(1)
	v88 = v78 + v85*int32(12)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	F_pfree(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v78)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L35
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(int32(1)) < base.Ui32(v81) {
		v81 = v85
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v105 = v99
	goto L29
L36:
	;
	goto L3
L37:
	;
	return
}
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v55 int64
	_ = v55
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	v2 = l1
	v6 = l5
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = v2 * int32(100)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = v18 << (uint(int32(4)) % 32)
	v22 = v17 + (l0 + v20)
	v24 = v22 - int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v7
	v28 = v22 - int32(76)
	if l2 == v7 {
		if v28&int32(3) == int32(0) {
			v38 = l0 + v17 + v20 - int32(76)
			v39 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v38)+56)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+40)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
		} else {
			v55 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28))) = v55
		}
	} else {
		if l2 == v28 {
		} else {
			v73 = F_strncpy(m, v28, l2, int32(64))
			mBase = m.M
			v74 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v73)+63)) = uint8(v74)
		}
	}
	v77 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+86)) = v77
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+80)) = uint16(v6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+74)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = l4
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+90)) = uint16(v77)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+94)) = uint16(v77)
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+92)) = uint8(v86)
	v89 = F_SearchSysCache1(m, int32(82), l3)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		return
	} else {
		if v89 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l3
				F_errmsg_internal(m, int32(_a_F_TupleDescInitEntry_0), v14)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_TupleDescInitEntry_1), int32(896), int32(_a_F_TupleDescInitEntry_2))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = l3
			v109 = v106 + v107
			v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v110)
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v112)
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v114)
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+129)))
			v117 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v117)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v116)
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v109)+144))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v120
			F_populate_compact_attribute(m, l0, v2-int32(1))
			mBase = m.M
			v125 = m.ExcPending
			if v125 != 0 {
				return
			} else {
				F_ReleaseCatCache(m, v89)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					m.G0 = v14 + int32(16)
					return
				}
			}
		}
	}
}
func F_apply_handle_tuple_routing(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v58 int32
	_ = v58
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int64
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v699 int32
	_ = v699
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_palloc0(m, int32(264))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(396)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(0)
	v46 = F_ExecSetupPartitionTupleRouting(m, v34, v33)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v49 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v52 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v54 = v49
	goto L6
L6:
	;
	v55 = int32(_a_F_apply_handle_tuple_routing_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v58
	v60 = F_ExecFindPartition(m, v36, v32, v46, l1, v34)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v54 = v52
	goto L6
L8:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63)+119)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v66 = F_get_namespace_name(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	F_CheckSubscriptionRelkind(m, v64, v66, v68+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+204))
	if v73 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v78 = F_table_slot_create(m, v62, v34+int32(104))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v80 = v73
	goto L13
L13:
	;
	v81 = F_ExecGetRootToChildMap(m, v60, v34)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v80 = v78
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v56
	if l3 != int32(3) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	if v81 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v84 = F_execute_attr_map_slot(m, v83, l1, v80)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+32))
	m.T0[v87].(func(*base.Module, int32, int32))(m, v80, l1)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v96 = v84
	v98 = v83
	goto L15
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+6)))
	if v91 <= v92 {
		v96 = v80
		v98 = v5
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_slot_getsomeattrs_int(m, v80, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v96 = v80
	v98 = v5
	goto L15
L24:
	;
	m.G0 = v29 + int32(96)
	return
L25:
	;
	v518 = F_GetTupleTransactionInfo(m, v480, v29+int32(24), v29+int32(28), v29+int32(32))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L102
	}
L26:
	;
	v103 = m.G0
	v105 = v103 + int32(-64)
	m.G0 = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v62)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[1]))
	if v111 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_InitConflictIndexes(m, v60)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L98
	}
L29:
	;
	v145 = v111
	goto L31
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2]))
	if v113 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v151 = F_hash_search(m, v145, v103+int32(-52), int32(1), v103+int32(-48))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[3]))
	v123 = F_AllocSetContextCreateInternal(m, v118, int32(_a_F_apply_handle_tuple_routing_1), int32(0), int32(_a_F_apply_handle_tuple_routing_2), int32(_a_F_apply_handle_tuple_routing_3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v126 = v113
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+56)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v105)+32)) = int64(343597383684)
	v136 = F_hash_create(m, int32(_a_F_apply_handle_tuple_routing_4), int32(64), v103+int32(-48), int32(1064))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2])) = v123
	v126 = v123
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[1])) = v136
	F_CacheRegisterRelcacheCallback(m, int32(1016))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[1]))
	v145 = v143
	goto L31
L38:
	;
	v154 = v151 + int32(8)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+16)))
	if v155 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	m.G0 = v105 - int32(-64)
	F_check_relation_updatable(m, v154)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L80
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+48)) = v62
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+44)) = v323
	if v98 != 0 {
		goto L65
	} else {
		goto L66
	}
L41:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v195 = F_pstrdup(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L52
	}
L42:
	;
	v175 = int32(_a_F_apply_handle_tuple_routing_0)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v179
	v182 = v151 + int32(52)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v151)+52))
	if v183 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+40)))
	if v158 != int32(1) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v162 = int32(_a_F_apply_handle_tuple_routing_0)
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v166
	base.MemoryFill(m, v151, int32(0), int32(80))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v171
	v190 = v151 + int32(52)
	v191 = v163
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+48)) = v62
	goto L39
L47:
	;
	F_free_attrmap(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v188 != 0 {
		v311 = v182
		v321 = v176
		goto L40
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = int32(0)
	goto L49
L51:
	;
	v190 = v182
	v191 = v176
	goto L41
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+12)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v199 = F_pstrdup(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+16)) = v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+20)) = v202
	v206 = F_palloc(m, v202<<(uint(int32(2))%32))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+24)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v212 = F_palloc(m, v209<<(uint(int32(2))%32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+28)) = v212
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if int32(0) < v215 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v235 = v5
	goto L59
L57:
	;
	goto L58
L58:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+32)) = uint8(v290)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	v293 = F_bms_copy(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L63
	}
L59:
	;
	v245 = v235 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245+v246)))
	v249 = F_pstrdup(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v151)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v251+v245))) = v249
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v151)+28))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256+v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v254+v245))) = v258
	v261 = v235 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v261 < v262 {
		v235 = v261
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+36)) = v293
	v311 = v190
	v321 = v191
	goto L40
L64:
	;
	F_logicalrep_rel_mark_updatable(m, v154)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L78
	}
L65:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v326 = F_make_attrmap(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v384 = F_make_attrmap(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L76
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v326
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v329 <= int32(0) {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v332 = int32(0)
	v344 = v326
	v350 = v332
	v351 = v332
	goto L70
L70:
	;
	v361 = v350 << (uint(int32(1)) % 32)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v361+v362))))
	if v364 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L64
L72:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365+v364<<(uint(int32(1))%32)-int32(2)))))
	v373 = v371
	goto L74
L73:
	;
	v373 = int32(_a_F_apply_handle_tuple_routing_5)
	goto L74
L74:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374+v361))) = uint16(v373)
	v378 = v351 + int32(1)
	v379 = base.I32_extend16_s(v378)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v379 < v381 {
		v344 = v380
		v350 = v379
		v351 = v378
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v389 = v387 << (uint(int32(1)) % 32)
	if v389 == int32(0) {
		goto L64
	} else {
		goto L77
	}
L77:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	base.MemoryCopy(m, v392, v393, v389)
	goto L64
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v321
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v151)+52))
	v426 = F_FindLogicalRepLocalIndex(m, v62, v31, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v428 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+40)) = uint8(v428)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+60)) = v426
	goto L39
L80:
	;
	if l3 != int32(2) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v154)+52))
	F_apply_handle_delete_internal(m, l0, v60, v96, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v467 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v467
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v154)+52))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_TargetPrivilegesCheck(m, v62, int64(2))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	goto L24
L85:
	;
	v480 = F_table_slot_create(m, v62, v474+int32(104))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v473 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	F_slot_store_data(m, v480, v154, l2)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L95
	}
L88:
	;
	v482 = F_RelationFindReplTupleByIndex(m, v62, v473, v96, v480)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v486 = F_RelationFindReplTupleSeq(m, v62, v96, v480)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	if v482 == int32(0) {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L25
L93:
	;
	if v486 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	goto L87
L95:
	;
	v491 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v491
	v499 = F_list_make1_impl(m, int32(1), v29+int32(4))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_ReportApplyConflict(m, v34, v60, int32(15), int32(3), v96, v480, v499)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L24
L98:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v506, int64(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_ExecSimpleRelationInsert(m, v60, v503, v96)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L24
L101:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v546 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	if v518 == int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+28)))
	v524 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[4])))
	if v522 == v524 {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v528 = F_table_slot_create(m, v62, v34+int32(104))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_slot_store_data(m, v528, v154, l2)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v480
	v534 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v534
	v538 = int32(1)
	v540 = F_list_make1_impl(m, v538, v29)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_ReportApplyConflict(m, v34, v60, int32(15), v538, v96, v528, v540)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L101
L109:
	;
	v549 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	v551 = v546
	goto L111
L111:
	;
	v552 = int32(_a_F_apply_handle_tuple_routing_0)
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v551)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v555
	F_slot_modify_data(m, v96, v480, v154, l2)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	v551 = v549
	goto L111
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v553
	v563 = int32(0)
	F_EvalPlanQualInit(m, v29+int32(44), v34, v563, v563, int32(-1), v563)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+131)))
	if v570 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	F_EvalPlanQualEnd(m, v29+int32(44))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L163
	}
L116:
	;
	if v81 != 0 {
		goto L126
	} else {
		goto L127
	}
L117:
	;
	v574 = F_ExecPartitionCheck(m, v60, v96, v34, int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	F_InitConflictIndexes(m, v60)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	if v574 == int32(0) {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v96
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v581, int64(4))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_ExecSimpleRelationUpdate(m, v60, v34, v29+int32(44), v480, v96)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L115
L125:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v608 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L126:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v62)+52))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v591 = F_convert_tuples_by_name(m, v589, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+32))
	m.T0[v597].(func(*base.Module, int32, int32))(m, l1, v96)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	v594 = F_execute_attr_map_slot(m, v593, v96, l1)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v606 = v594
	goto L125
L131:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	v602 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v601 <= v602 {
		v606 = l1
		goto L125
	} else {
		goto L132
	}
L132:
	;
	F_slot_getsomeattrs_int(m, l1, v601)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v606 = l1
	goto L125
L134:
	;
	v611 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	v613 = v608
	goto L136
L136:
	;
	v614 = int32(_a_F_apply_handle_tuple_routing_0)
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v617
	v619 = F_ExecFindPartition(m, v36, v32, v46, v606, v34)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L138
	}
L137:
	;
	v613 = v611
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v615
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)+48))
	v625 = int32(*(*int8)(unsafe.Add(mBase, uint32(v624)+119)))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v624)+68))
	v627 = F_get_namespace_name(m, v626)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v623)+48))
	F_CheckSubscriptionRelkind(m, v625, v627, v629+int32(4))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v480
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v635, int64(8))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_ExecSimpleRelationDelete(m, v60, v34, v29+int32(44), v480)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v643 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v646 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	v648 = v643
	goto L145
L145:
	;
	v649 = int32(_a_F_apply_handle_tuple_routing_0)
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v648)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v652
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v619)+204))
	if v654 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v648 = v646
	goto L145
L147:
	;
	v659 = F_table_slot_create(m, v623, v34+int32(104))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	v661 = v654
	goto L149
L149:
	;
	v662 = F_ExecGetRootToChildMap(m, v619, v34)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	v661 = v659
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v650
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_InitConflictIndexes(m, v619)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L160
	}
L152:
	;
	if v662 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v662)+8))
	v665 = F_execute_attr_map_slot(m, v664, v606, v661)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661)+8))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)+32))
	m.T0[v668].(func(*base.Module, int32, int32))(m, v661, v606)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v677 = v665
	goto L151
L157:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v606)+12))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	v673 = int32(*(*int16)(unsafe.Add(mBase, uint32(v606)+6)))
	if v672 <= v673 {
		v677 = v661
		goto L151
	} else {
		goto L158
	}
L158:
	;
	F_slot_getsomeattrs_int(m, v606, v672)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v677 = v661
	goto L151
L160:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	F_TargetPrivilegesCheck(m, v684, int64(1))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_ExecSimpleRelationInsert(m, v619, v681, v677)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	goto L115
L163:
	;
	goto L24
}
