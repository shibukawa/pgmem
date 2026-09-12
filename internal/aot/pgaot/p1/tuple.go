package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v10 = F_palloc(m, l0*int32(116)+int32(20))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(-4294965047)
	if int32(0) < l0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v10
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v32 = int32(100)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1+v26<<(uint(int32(2))%32))))
	goto L9
L7:
	;
	goto L5
L8:
	;
	F_populate_compact_attribute(m, v10, v26)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v40 = F__emscripten_memcpy_bulkmem(m, v10+int32(20)+v28<<(uint(int32(4))%32)+v26*v32, v38, v32)
	mBase = m.M
	goto L11
L11:
	;
	goto L8
L12:
	;
	v45 = v26 + int32(1)
	if v45 != l0 {
		v26 = v45
		goto L6
	} else {
		goto L13
	}
L13:
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
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
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
	v116 = m.ExcPending
	if v116 != 0 {
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
	v75 = v28
	goto L15
L15:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+14)))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v41 = v35
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v32)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L26
	}
L19:
	;
	v47 = v32 + v41<<(uint(int32(3))%32)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	if int32(0) < v41 {
		v41 = v41 - int32(1)
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(102)+v51<<(uint(int32(4))%32)+v41*int32(100)))))
	if v58 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	F_pfree(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
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
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = v73
	goto L15
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v83 = v79
	goto L30
L28:
	;
	v107 = v75
	goto L29
L29:
	;
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L9
	} else {
		goto L36
	}
L30:
	;
	v87 = v83 - int32(1)
	v90 = v80 + v87*int32(12)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	F_pfree(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v80)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L35
	}
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	F_pfree(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(int32(1)) < base.Ui32(v83) {
		v83 = v87
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v107 = v101
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	v2 = l1
	v6 = l5
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = v16 << (uint(int32(4)) % 32)
	v19 = l0 + v18
	v22 = v19 + v2*int32(100)
	v24 = v22 - int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v7
	v28 = v22 - int32(76)
	if l2 == v7 {
		if v28&int32(3) == int32(0) {
			if base.Ui32(v24+int32(68)) <= base.Ui32(v28) {
			} else {
				v40 = v2 * int32(100)
				v42 = v40 + l0 + v18
				v44 = v42 - int32(12)
				v46 = v42 - int32(72)
				if base.Ui32(v46) < base.Ui32(v44) {
					v48 = v44
				} else {
					v48 = v46
				}
				v58 = F__emscripten_memset_bulkmem(m, v28, base.I32_extend8_s(int32(0)), (v48-(v40+v19)+int32(75))&int32(-4)+int32(4))
				mBase = m.M
			}
		} else {
			v59 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v28))) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v59
		}
	} else {
		if l2 == v28 {
		} else {
			v77 = F_strncpy(m, v28, l2, int32(64))
			mBase = m.M
			v78 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v77)+63)) = uint8(v78)
		}
	}
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+86)) = v83
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+80)) = uint16(v6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+74)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = l4
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+94)) = uint16(v83)
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+92)) = uint8(v90)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+90)) = uint16(v83)
	v95 = F_SearchSysCache1(m, int32(82), l3)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		return
	} else {
		if v95 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l3
				F_errmsg_internal(m, int32(54674), v14)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					F_errfinish(m, int32(520533), int32(896), int32(12651))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+22)))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = l3
			v115 = v113 + v112
			v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v116)
			v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v118)
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v120)
			v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+129)))
			v123 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v123)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v122)
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v115)+144))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v126
			F_populate_compact_attribute(m, l0, v2-int32(1))
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return
			} else {
				F_ReleaseCatCache(m, v95)
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v238 int32
	_ = v238
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
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
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
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
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
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
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
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
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v733 int32
	_ = v733
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
	v55 = int32(4549024)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v58
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v56
	v103 = l3 - int32(2)
	switch v103 {
	case 0, 2:
		goto L25
	default:
		v474 = int32(0)
		goto L24
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
	switch v103 - int32(1) {
	case 0:
		goto L84
	case 1:
		goto L83
	default:
		goto L82
	}
L25:
	;
	v104 = m.G0
	v106 = v104 + int32(-64)
	m.G0 = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v62)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v112 = *(*int32)(unsafe.Add(mBase, _consts[663]))
	if v112 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v146 = v112
	goto L28
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	if v114 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v152 = F_hash_search(m, v146, v104+int32(-52), int32(1), v104+int32(-48))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L35
	}
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	v124 = F_AllocSetContextCreateInternal(m, v119, int32(67406), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v127 = v114
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+56)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v106)+32)) = int64(343597383684)
	v137 = F_hash_create(m, int32(415294), int32(64), v104+int32(-48), int32(1064))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[664])) = v124
	v127 = v124
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[663])) = v137
	F_CacheRegisterRelcacheCallback(m, int32(1016))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[663]))
	v146 = v144
	goto L28
L35:
	;
	v155 = v152 + int32(8)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+16)))
	if v156 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	m.G0 = v106 - int32(-64)
	F_check_relation_updatable(m, v155)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L80
	}
L37:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v195 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v177 = int32(4549024)
	v178 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v181 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v181
	v184 = v152 + int32(52)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
	if v185 == int32(0) {
		v193 = v184
		v194 = v178
		goto L37
	} else {
		goto L44
	}
L39:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+40)))
	if v159 != int32(1) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v163 = int32(4549024)
	v164 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v167
	v172 = F__emscripten_memset_bulkmem(m, v152, base.I32_extend8_s(int32(0)), int32(80))
	mBase = m.M
	goto L43
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+48)) = v62
	goto L36
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v173
	v193 = v172 + int32(52)
	v194 = v164
	goto L37
L44:
	;
	F_free_attrmap(m, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = int32(0)
	v193 = v184
	v194 = v178
	goto L37
L46:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+8)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v201 = F_pstrdup(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+48)) = v62
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+44)) = v329
	if v98 != 0 {
		goto L62
	} else {
		goto L63
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+12)) = v201
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v205 = F_pstrdup(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+16)) = v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+20)) = v208
	v212 = F_palloc(m, v208<<(uint(int32(2))%32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+24)) = v212
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v218 = F_palloc(m, v215<<(uint(int32(2))%32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+28)) = v218
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if int32(0) < v221 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v238 = v5
	goto L56
L54:
	;
	goto L55
L55:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+32)) = uint8(v296)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	v299 = F_bms_copy(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L60
	}
L56:
	;
	v251 = v238 << (uint(int32(2)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251+v252)))
	v255 = F_pstrdup(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v152)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v257+v251))) = v255
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v152)+28))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v262+v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v260+v251))) = v264
	v267 = v238 + int32(1)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v267 < v268 {
		v238 = v267
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+36)) = v299
	goto L48
L61:
	;
	F_logicalrep_rel_mark_updatable(m, v155)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L78
	}
L62:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v332 = F_make_attrmap(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v390 = F_make_attrmap(m, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L73
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v332
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	if v335 <= int32(0) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v338 = int32(0)
	v345 = v332
	v354 = v338
	v357 = v338
	goto L67
L67:
	;
	v367 = v354 << (uint(int32(1)) % 32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v370 = int32(*(*int16)(unsafe.Add(mBase, uint32(v367+v368))))
	if v370 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L61
L69:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v371+v370<<(uint(int32(1))%32)-int32(2)))))
	v379 = v377
	goto L71
L70:
	;
	v379 = int32(65535)
	goto L71
L71:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	*(*uint16)(unsafe.Add(mBase, uint32(v380+v367))) = uint16(v379)
	v384 = v357 + int32(1)
	v385 = base.I32_extend16_s(v384)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v385 < v387 {
		v345 = v386
		v354 = v385
		v357 = v384
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v390
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v397 = v395 << (uint(int32(1)) % 32)
	if v397 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L61
L75:
	;
	v398 = F__emscripten_memcpy_bulkmem(m, v393, v394, v397)
	mBase = m.M
	goto L77
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v194
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
	v431 = F_FindLogicalRepLocalIndex(m, v62, v31, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v433 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+40)) = uint8(v433)
	*(*int32)(unsafe.Add(mBase, uint32(v152)+60)) = v431
	goto L36
L80:
	;
	v474 = v155
	goto L24
L81:
	;
	m.G0 = v29 + int32(96)
	return
L82:
	;
	v507 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v507
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v507
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v507
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v474)+52))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_TargetPrivilegesCheck(m, v62, int64(2))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L89
	}
L83:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v474)+52))
	F_apply_handle_delete_internal(m, l0, v60, v96, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L88
	}
L84:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_InitConflictIndexes(m, v60)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v498, int64(1))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_ExecSimpleRelationInsert(m, v60, v495, v96)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L81
L88:
	;
	goto L81
L89:
	;
	v520 = F_table_slot_create(m, v62, v514+int32(104))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v513 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v551 = F_GetTupleTransactionInfo(m, v520, v29+int32(24), v29+int32(28), v29+int32(32))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L104
	}
L92:
	;
	F_slot_store_data(m, v520, v474, l2)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L100
	}
L93:
	;
	v522 = F_RelationFindReplTupleByIndex(m, v62, v513, v96, v520)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v526 = F_RelationFindReplTupleSeq(m, v62, v96, v520)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	if v522 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L91
L98:
	;
	if v526 != 0 {
		goto L91
	} else {
		goto L99
	}
L99:
	;
	goto L92
L100:
	;
	v531 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v531
	v541 = F_list_make1_impl(m, int32(1), v29+int32(4))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_ReportApplyConflict(m, v34, v60, int32(15), int32(3), v96, v520, v541)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L81
L103:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v580 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	if v551 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+28)))
	v557 = int32(*(*uint16)(unsafe.Add(mBase, _consts[169])))
	if v555 == v557 {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v561 = F_table_slot_create(m, v62, v34+int32(104))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_slot_store_data(m, v561, v474, l2)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v520
	v567 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v567
	v573 = int32(1)
	v575 = F_list_make1_impl(m, v573, v29)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_ReportApplyConflict(m, v34, v60, int32(15), v573, v96, v561, v575)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	goto L103
L111:
	;
	v583 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	v585 = v580
	goto L113
L113:
	;
	v586 = int32(4549024)
	v587 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v589
	F_slot_modify_data(m, v96, v520, v474, l2)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	v585 = v583
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v587
	v597 = int32(0)
	F_EvalPlanQualInit(m, v29+int32(44), v34, v597, v597, int32(-1), v597)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+131)))
	if v604 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	F_EvalPlanQualEnd(m, v29+int32(44))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L165
	}
L118:
	;
	if v81 != 0 {
		goto L128
	} else {
		goto L129
	}
L119:
	;
	v608 = F_ExecPartitionCheck(m, v60, v96, v34, int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	F_InitConflictIndexes(m, v60)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	if v608 == int32(0) {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v96
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v615, int64(4))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_ExecSimpleRelationUpdate(m, v60, v34, v29+int32(44), v520, v96)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L117
L127:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v642 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L128:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v62)+52))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v625 = F_convert_tuples_by_name(m, v623, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+32))
	m.T0[v631].(func(*base.Module, int32, int32))(m, l1, v96)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v625)+8))
	v628 = F_execute_attr_map_slot(m, v627, v96, l1)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v640 = v628
	goto L127
L133:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v635 <= v636 {
		v640 = l1
		goto L127
	} else {
		goto L134
	}
L134:
	;
	F_slot_getsomeattrs_int(m, l1, v635)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v640 = l1
	goto L127
L136:
	;
	v645 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	v647 = v642
	goto L138
L138:
	;
	v648 = int32(4549024)
	v649 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v647)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v651
	v653 = F_ExecFindPartition(m, v36, v32, v46, v640, v34)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L140
	}
L139:
	;
	v647 = v645
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v649
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653)+8))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+48))
	v659 = int32(*(*int8)(unsafe.Add(mBase, uint32(v658)+119)))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v658)+68))
	v661 = F_get_namespace_name(m, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v657)+48))
	F_CheckSubscriptionRelkind(m, v659, v661, v663+int32(4))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v520
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v669, int64(8))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_ExecSimpleRelationDelete(m, v60, v34, v29+int32(44), v520)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v677 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v680 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v682 = v677
	goto L147
L147:
	;
	v683 = int32(4549024)
	v684 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v682)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v653)+204))
	if v688 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v682 = v680
	goto L147
L149:
	;
	v693 = F_table_slot_create(m, v657, v34+int32(104))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	v695 = v688
	goto L151
L151:
	;
	v696 = F_ExecGetRootToChildMap(m, v653, v34)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	v695 = v693
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v684
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_InitConflictIndexes(m, v653)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L162
	}
L154:
	;
	if v696 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
	v699 = F_execute_attr_map_slot(m, v698, v640, v695)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v695)+8))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+32))
	m.T0[v702].(func(*base.Module, int32, int32))(m, v695, v640)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v711 = v699
	goto L153
L159:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v640)+12))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v707 = int32(*(*int16)(unsafe.Add(mBase, uint32(v640)+6)))
	if v706 <= v707 {
		v711 = v695
		goto L153
	} else {
		goto L160
	}
L160:
	;
	F_slot_getsomeattrs_int(m, v640, v706)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v711 = v695
	goto L153
L162:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v653)+8))
	F_TargetPrivilegesCheck(m, v718, int64(1))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_ExecSimpleRelationInsert(m, v653, v715, v711)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L117
L165:
	;
	goto L81
}
