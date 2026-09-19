package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
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
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
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
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 float32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v271 float64
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = int32(0)
	if v27 == v28 {
		v44 = v28
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v44&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v31 == int32(0) {
		v44 = v28
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 != int32(7) {
		v44 = v28
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 != int32(17) {
		v44 = v28
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	v44 = v40 ^ int32(1)
	goto L4
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = F_get_fn_opclass_options(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v55 = int32(95)
	goto L11
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v58 = int32(base.Ui32(v56) >> (uint(int32(2)) % 32))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v61 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v55 = v50<<(uint(int32(3))%32) - int32(1)
	goto L11
L13:
	;
	v170 = int32(0)
	if base.B2i32(int32(1)<<(uint(v25)%32)&int32(1284) == v170)|base.B2i32(base.Ui32(int32(10)) < base.Ui32(v25)) == v170 {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v134 = int32(4)
	v138 = F_generate_trgm(m, v20+v134, v58-v134)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if int32(base.Ui32(v64)>>(uint(int32(2))%32)) != v58 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v58) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v129 != 0 {
		goto L14
	} else {
		goto L35
	}
L18:
	;
	v129 = int32(0)
	goto L17
L19:
	;
	v103 = v98
	v104 = v99
	v105 = v100
	goto L29
L20:
	;
	if (v61|v20)&int32(3) != 0 {
		v98 = v61
		v99 = v20
		v100 = v58
		goto L19
	} else {
		goto L23
	}
L21:
	;
	v91 = v61
	v92 = v20
	v93 = v58
	goto L22
L22:
	;
	if v93 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L23:
	;
	v75 = v61
	v76 = v20
	v77 = v58
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v80 != v81 {
		v98 = v75
		v99 = v76
		v100 = v77
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v91 = v86
	v92 = v84
	v93 = v88
	goto L22
L26:
	;
	v83 = int32(4)
	v84 = v76 + v83
	v86 = v75 + v83
	v88 = v77 - v83
	if base.Ui32(int32(3)) < base.Ui32(v88) {
		v75 = v86
		v76 = v84
		v77 = v88
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v98 = v91
	v99 = v92
	v100 = v93
	goto L19
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 == v109 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v129 = v108 - v109
	goto L17
L31:
	;
	v111 = int32(1)
	v116 = v105 - v111
	if v116 != 0 {
		v103 = v103 + v111
		v104 = v104 + v111
		v105 = v116
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L18
L35:
	;
	v163 = v61
	v164 = (v58 + int32(7)) & int32(2147483640)
	goto L13
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	v145 = (v58 + int32(7)) & int32(2147483640)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v150 = F_MemoryContextAlloc(m, v141, v145+int32(base.Ui32(v146)>>(uint(int32(2))%32)))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v58 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	base.MemoryCopy(m, v150, v20, v58)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v155 = int32(base.Ui32(v153) >> (uint(int32(2)) % 32))
	if v155 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	base.MemoryCopy(m, v150+v145, v138, v155)
	goto L43
L42:
	;
	goto L43
L43:
	;
	if v61 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_pfree(m, v61)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v150
	v163 = v150
	v164 = v145
	goto L13
L47:
	;
	goto L46
L48:
	;
	v177 = v163 + v164
	v179 = base.B2i32(v25 != int32(2))
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v179)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v182)+12)))
	if v184&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L65
	}
L51:
	;
	v272 = F_Float8GetDatum(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L64
	}
L52:
	;
	v187 = F_cnt_sml(m, v177, v59, v179)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
	if v195&int32(4) != 0 {
		v271 = float64(0)
		goto L51
	} else {
		goto L56
	}
L55:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v16)+12)) = v187
	v271 = base.F64_sub(float64(1), base.F64_promote_f32(v187))
	goto L51
L56:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v203 = int32(base.Ui32(v199)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v203) < base.Ui32(int32(3)) {
		v271 = float64(-1)
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v206 = int32(5)
	v210 = int32(1)
	v212 = base.I32_div_u_s(v203, int32(3))
	if base.Ui32(v212) <= base.Ui32(v210) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v215 = v210
	goto L60
L59:
	;
	v215 = v212
	goto L60
L60:
	;
	v216 = int32(0)
	v218 = v216
	v219 = v216
	goto L61
L61:
	;
	v231 = int32(3)
	v233 = v177 + v206 + v218*v231
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233))))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+2)))
	v239 = base.I32_rem_u_s(v234|v235<<(uint(int32(16))%32), v55)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v206+int32(base.Ui32(v239)>>(uint(v231)%32))))))
	v247 = int32(1)
	v249 = v219 + int32(base.Ui32(v243)>>(uint(v239&int32(7))%32))&v247
	v251 = v218 + v247
	if v251 != v215 {
		v218 = v251
		v219 = v249
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v271 = base.F64_sub(float64(1), base.F64_div(base.F64_convert_i32_u(v249), base.F64_convert_i32_u(v212)))
	goto L51
L63:
	;
	goto L62
L64:
	;
	m.G0 = v16 + int32(16)
	return v272
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v25
	F_errmsg_internal(m, int32(_a_F_gtrgm_distance_0), v16)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_gtrgm_distance_1), int32(525), int32(_a_F_gtrgm_distance_2))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gtrgm_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_gtrgm_in_0), int32(61), int32(_a_F_gtrgm_in_1), int32(_a_F_gtrgm_in_2), int32(_a_F_gtrgm_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gtrgm_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v322 int64
	_ = v322
	var v323 int32
	_ = v323
	var v326 int64
	_ = v326
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v354 int64
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v529 int64
	_ = v529
	var v530 int32
	_ = v530
	var v533 int64
	_ = v533
	var v534 int32
	_ = v534
	var v537 int64
	_ = v537
	var v538 int32
	_ = v538
	var v541 int64
	_ = v541
	var v545 int64
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v569 int64
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v586 int64
	_ = v586
	var v587 int32
	_ = v587
	var v590 int64
	_ = v590
	var v591 int64
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v627 int64
	_ = v627
	var v628 int32
	_ = v628
	var v631 int64
	_ = v631
	var v632 int32
	_ = v632
	var v635 int64
	_ = v635
	var v636 int32
	_ = v636
	var v639 int64
	_ = v639
	var v640 int32
	_ = v640
	var v643 int64
	_ = v643
	var v647 int64
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v671 int64
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v689 int64
	_ = v689
	var v690 int32
	_ = v690
	var v693 int64
	_ = v693
	var v694 int64
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v785 int64
	_ = v785
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int64
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v832 int64
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int64
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int64
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int64
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int64
	_ = v870
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int64
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int64
	_ = v890
	var v891 int64
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int64
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v910 int64
	_ = v910
	var v911 int32
	_ = v911
	var v912 int64
	_ = v912
	var v913 int32
	_ = v913
	var v914 int64
	_ = v914
	var v915 int32
	_ = v915
	var v916 int64
	_ = v916
	var v917 int32
	_ = v917
	var v918 int64
	_ = v918
	var v922 int64
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v933 int64
	_ = v933
	var v938 int64
	_ = v938
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v965 int64
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v985 int64
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v997 int64
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int64
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int64
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int64
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int64
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int64
	_ = v1043
	var v1044 int64
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1054 int64
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1063 int64
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int64
	_ = v1071
	var v1075 int64
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1086 int64
	_ = v1086
	var v1091 int64
	_ = v1091
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1118 int64
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1138 int64
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1150 int64
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int64
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1166 int64
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1176 int64
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int64
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int64
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1207 int64
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1216 int64
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int64
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int64
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int64
	_ = v1224
	var v1228 int64
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1239 int64
	_ = v1239
	var v1256 int64
	_ = v1256
	var v1277 int64
	_ = v1277
	var v1285 int32
	_ = v1285
	var v1318 int64
	_ = v1318
	v2 = int32(0)
	v17 = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == v2 {
		v39 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v26 == int32(0) {
		v39 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(7) {
		v39 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 != int32(17) {
		v39 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v39 = v35 ^ int32(1)
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_fn_opclass_options(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = int32(12)
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v54 = v50 + int32(5)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if v55&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = v47
	goto L9
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = base.F32_div(base.F32_convert_i32_s(v280+(base.I32_wrap_i64(v1318)^int32(-1))), base.F32_convert_i32_s(v280))
	return v18
L13:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = base.F32_convert_i32_s(v1285)
	return v18
L14:
	;
	v1285 = v48<<(uint(v494)%32) + (base.I32_wrap_i64(v1277) ^ int32(-1))
	goto L13
L15:
	;
	v1285 = v48<<(uint(v597)%32) + (base.I32_wrap_i64(v1256) ^ int32(-1))
	goto L13
L16:
	;
	v1091 = int64(0)
	if base.B2i32(v54 != (v50+int32(8))&int32(-4))|base.B2i32(v48 < int32(4)) != 0 {
		v1170 = v54
		v1171 = v48
		v1176 = v1091
		goto L194
	} else {
		goto L195
	}
L17:
	;
	v938 = int64(0)
	if base.B2i32(v493 != (v49+int32(8))&int32(-4))|base.B2i32(v48 < int32(4)) != 0 {
		v1017 = v493
		v1018 = v48
		v1023 = v938
		goto L168
	} else {
		goto L169
	}
L18:
	;
	v785 = int64(0)
	if base.B2i32(v260 != (v260+int32(3))&int32(-4))|base.B2i32(v48 < int32(4)) != 0 {
		v864 = v260
		v865 = v48
		v870 = v785
		goto L142
	} else {
		goto L143
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v60 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	v64 = (v48 + int32(7)) & int32(-8)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v66 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v487 = int32(4)
	v488 = v55 & v487
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+4)))
	if v489&v487 != 0 {
		goto L98
	} else {
		goto L99
	}
L22:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+4)))
	if v276&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v139 = F_MemoryContextAlloc(m, v137, v60+v64)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L45
	}
L24:
	;
	v69 = v66 + v64
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if int32(base.Ui32(v70)>>(uint(int32(2))%32)) != v60 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v60) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v135 != 0 {
		goto L23
	} else {
		goto L44
	}
L27:
	;
	v135 = int32(0)
	goto L26
L28:
	;
	v109 = v104
	v110 = v105
	v111 = v106
	goto L38
L29:
	;
	if (v69|v49)&int32(3) != 0 {
		v104 = v69
		v105 = v49
		v106 = v60
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v97 = v69
	v98 = v49
	v99 = v60
	goto L31
L31:
	;
	if v99 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v81 = v69
	v82 = v49
	v83 = v60
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v86 != v87 {
		v104 = v81
		v105 = v82
		v106 = v83
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v97 = v92
	v98 = v90
	v99 = v94
	goto L31
L35:
	;
	v89 = int32(4)
	v90 = v82 + v89
	v92 = v81 + v89
	v94 = v83 - v89
	if base.Ui32(int32(3)) < base.Ui32(v94) {
		v81 = v92
		v82 = v90
		v83 = v94
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v104 = v97
	v105 = v98
	v106 = v99
	goto L28
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 == v115 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v135 = v114 - v115
	goto L26
L40:
	;
	v117 = int32(1)
	v122 = v111 - v117
	if v122 != 0 {
		v109 = v109 + v117
		v110 = v110 + v117
		v111 = v122
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
	v260 = v66
	goto L22
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v145 = int32(base.Ui32(v141)>>(uint(int32(2))%32)) - int32(5)
	v146 = int32(3)
	v147 = base.I32_div_u_s(v145, v146)
	if v139&v146 != 0 {
		v170 = v48
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v177 = int32(3)
	v180 = v48<<(uint(v177)%32) - int32(1)
	v182 = base.I32_div_s(v180, int32(8))
	v183 = v139 + v182
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v186 = v184 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v186)
	if base.Ui32(v177) <= base.Ui32(v145) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	if v170 == int32(0) {
		goto L46
	} else {
		goto L55
	}
L48:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v48) {
		v170 = v48
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v48&int32(3) != 0 {
		v170 = v48
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v48 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v158 = v139 + v48
	v160 = v139 + int32(4)
	if base.Ui32(v160) < base.Ui32(v158) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v162 = v158
	goto L54
L53:
	;
	v162 = v160
	goto L54
L54:
	;
	v170 = (v139^int32(-1)+v162)&int32(-4) + int32(4)
	goto L47
L55:
	;
	base.MemoryFill(m, v139, int32(0), v170)
	goto L46
L56:
	;
	v192 = int32(1)
	if base.Ui32(v147) <= base.Ui32(v192) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v60 != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v195 = v192
	goto L61
L60:
	;
	v195 = v147
	goto L61
L61:
	;
	v200 = int32(0)
	goto L62
L62:
	;
	v214 = int32(3)
	v216 = v49 + int32(5) + v200*v214
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+2)))
	v222 = base.I32_rem_u_s(v217|v218<<(uint(int32(16))%32), v180)
	v225 = v139 + int32(base.Ui32(v222)>>(uint(v214)%32))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v227 = int32(1)
	v231 = v226 | v227<<(uint(v222&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v231)
	v234 = v200 + v227
	if v234 != v195 {
		v200 = v234
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L58
L64:
	;
	goto L63
L65:
	;
	base.MemoryCopy(m, v139+v64, v49, v60)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if v66 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_pfree(m, v66)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+16)) = v139
	v260 = v139
	goto L22
L71:
	;
	goto L70
L72:
	;
	v279 = int32(3)
	v280 = v48 << (uint(v279) % 32)
	if v279 < v48 {
		goto L18
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v48 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L75:
	;
	if v48 == int32(0) {
		v1318 = v17
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v285 = int32(3)
	v286 = v48 & v285
	if base.Ui32(v285) <= base.Ui32(v48-int32(1)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v295 = v260
	v299 = int32(0)
	v310 = v17
	goto L80
L78:
	;
	v339 = v260
	v354 = v17
	goto L79
L79:
	;
	v357 = v339
	v359 = int32(0)
	v372 = v354
	goto L84
L80:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+3)))
	v314 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v311)+uint32(_c_F_gtrgm_penalty[0]))))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)))
	v318 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v315)+uint32(_c_F_gtrgm_penalty[0]))))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	v322 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_gtrgm_penalty[0]))))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v326 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_gtrgm_penalty[0]))))
	v330 = v314 + (v318 + (v322 + (v310 + v326)))
	v331 = int32(4)
	v332 = v295 + v331
	v334 = v299 + v331
	if v334 != v48&int32(-4) {
		v295 = v332
		v299 = v334
		v310 = v330
		goto L80
	} else {
		goto L82
	}
L81:
	;
	if v286 == int32(0) {
		v1318 = v330
		goto L12
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v339 = v332
	v354 = v330
	goto L79
L84:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v376 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v373)+uint32(_c_F_gtrgm_penalty[0]))))
	v377 = v372 + v376
	v378 = int32(1)
	v381 = v359 + v378
	if v381 != v286 {
		v357 = v357 + v378
		v359 = v381
		v372 = v377
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v1318 = v377
	goto L12
L86:
	;
	goto L85
L87:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = float32(0)
	return v18
L88:
	;
	goto L89
L89:
	;
	v388 = int32(0)
	if v48 != int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = base.F32_convert_i32_u(v467)
	return v18
L91:
	;
	v397 = v388
	v400 = v388
	v406 = int32(0)
	goto L94
L92:
	;
	v441 = v388
	v444 = v388
	goto L93
L93:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v444))))
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444+v54))))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459^v461)+uint32(_c_F_gtrgm_penalty[0]))))
	v467 = v441 + v465
	goto L90
L94:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v400))))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v54))))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415^v417)+uint32(_c_F_gtrgm_penalty[0]))))
	v424 = v400 | int32(1)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v424))))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v260))))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426^v428)+uint32(_c_F_gtrgm_penalty[0]))))
	v433 = v397 + v421 + v432
	v434 = int32(2)
	v435 = v400 + v434
	v437 = v406 + v434
	if v437 != v48&int32(2147483646) {
		v397 = v433
		v400 = v435
		v406 = v437
		goto L94
	} else {
		goto L96
	}
L95:
	;
	if v48&int32(1) == int32(0) {
		v467 = v433
		goto L90
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	v441 = v433
	v444 = v435
	goto L93
L98:
	;
	if v488 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	if v488 != 0 {
		goto L116
	} else {
		goto L117
	}
L101:
	;
	v1285 = v2
	goto L13
L102:
	;
	goto L103
L103:
	;
	v493 = v49 + int32(5)
	v494 = int32(3)
	if v494 < v48 {
		goto L17
	} else {
		goto L104
	}
L104:
	;
	if v48 == int32(0) {
		v1277 = v17
		goto L14
	} else {
		goto L105
	}
L105:
	;
	v500 = int32(3)
	v501 = v48 & v500
	if base.Ui32(v500) <= base.Ui32(v48-int32(1)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v510 = v493
	v514 = int32(0)
	v525 = v17
	goto L109
L107:
	;
	v554 = v493
	v569 = v17
	goto L108
L108:
	;
	v571 = v554
	v573 = v2
	v586 = v569
	goto L113
L109:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+3)))
	v529 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v526)+uint32(_c_F_gtrgm_penalty[0]))))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+2)))
	v533 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v530)+uint32(_c_F_gtrgm_penalty[0]))))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	v537 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v534)+uint32(_c_F_gtrgm_penalty[0]))))
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	v541 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v538)+uint32(_c_F_gtrgm_penalty[0]))))
	v545 = v529 + (v533 + (v537 + (v525 + v541)))
	v546 = int32(4)
	v547 = v510 + v546
	v549 = v514 + v546
	if v549 != v48&int32(-4) {
		v510 = v547
		v514 = v549
		v525 = v545
		goto L109
	} else {
		goto L111
	}
L110:
	;
	if v501 == int32(0) {
		v1277 = v545
		goto L14
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v554 = v547
	v569 = v545
	goto L108
L113:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	v590 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v587)+uint32(_c_F_gtrgm_penalty[0]))))
	v591 = v586 + v590
	v592 = int32(1)
	v595 = v573 + v592
	if v595 != v501 {
		v571 = v571 + v592
		v573 = v595
		v586 = v591
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v1277 = v591
	goto L14
L115:
	;
	goto L114
L116:
	;
	v597 = int32(3)
	if v597 < v48 {
		goto L16
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v48 <= int32(0) {
		goto L131
	} else {
		goto L132
	}
L119:
	;
	if v48 == int32(0) {
		v1256 = v17
		goto L15
	} else {
		goto L120
	}
L120:
	;
	v603 = int32(3)
	v604 = v48 & v603
	if base.Ui32(v603) <= base.Ui32(v48-int32(1)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v614 = v2
	v616 = v54
	v627 = v17
	goto L124
L122:
	;
	v660 = v54
	v671 = v17
	goto L123
L123:
	;
	v674 = int32(0)
	v678 = v660
	v689 = v671
	goto L128
L124:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616)+3)))
	v631 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v628)+uint32(_c_F_gtrgm_penalty[0]))))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616)+2)))
	v635 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_gtrgm_penalty[0]))))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616)+1)))
	v639 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v636)+uint32(_c_F_gtrgm_penalty[0]))))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	v643 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v640)+uint32(_c_F_gtrgm_penalty[0]))))
	v647 = v631 + (v635 + (v639 + (v627 + v643)))
	v648 = int32(4)
	v649 = v616 + v648
	v651 = v614 + v648
	if v651 != v48&int32(-4) {
		v614 = v651
		v616 = v649
		v627 = v647
		goto L124
	} else {
		goto L126
	}
L125:
	;
	if v604 == int32(0) {
		v1256 = v647
		goto L15
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v660 = v649
	v671 = v647
	goto L123
L128:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678))))
	v693 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v690)+uint32(_c_F_gtrgm_penalty[0]))))
	v694 = v689 + v693
	v695 = int32(1)
	v698 = v674 + v695
	if v698 != v604 {
		v674 = v698
		v678 = v678 + v695
		v689 = v694
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v1256 = v694
	goto L15
L130:
	;
	goto L129
L131:
	;
	v1285 = v2
	goto L13
L132:
	;
	goto L133
L133:
	;
	v703 = v49 + int32(5)
	v704 = int32(0)
	if v48 != int32(1) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v712 = v704
	v714 = v2
	v720 = v2
	goto L137
L135:
	;
	v756 = v704
	v758 = v2
	goto L136
L136:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756+v703))))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756+v54))))
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773^v775)+uint32(_c_F_gtrgm_penalty[0]))))
	v1285 = v758 + v779
	goto L13
L137:
	;
	v729 = v712 | int32(1)
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703+v729))))
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729+v54))))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731^v733)+uint32(_c_F_gtrgm_penalty[0]))))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712+v54))))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712+v703))))
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739^v741)+uint32(_c_F_gtrgm_penalty[0]))))
	v747 = v737 + (v714 + v745)
	v748 = int32(2)
	v749 = v712 + v748
	v751 = v720 + v748
	if v751 != v48&int32(2147483646) {
		v712 = v749
		v714 = v747
		v720 = v751
		goto L137
	} else {
		goto L139
	}
L138:
	;
	if v48&int32(1) == int32(0) {
		v1285 = v747
		goto L13
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v756 = v749
	v758 = v747
	goto L136
L141:
	;
	v1318 = v933
	goto L12
L142:
	;
	if v865 == int32(0) {
		v933 = v870
		goto L154
	} else {
		goto L155
	}
L143:
	;
	v795 = v48 - int32(4)
	v799 = int32(base.Ui32(v795)>>(uint(int32(2))%32)) + int32(1)
	v801 = v799 & int32(3)
	if base.Ui32(int32(12)) <= base.Ui32(v795) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v806 = v260
	v807 = v48
	v810 = int32(0)
	v812 = v785
	goto L147
L145:
	;
	v838 = v260
	v839 = v48
	v844 = v785
	goto L146
L146:
	;
	v846 = v838
	v847 = v839
	v848 = int32(0)
	v852 = v844
	goto L151
L147:
	;
	v813 = int32(16)
	v814 = v807 - v813
	v816 = v806 + v813
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v806)+12))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v806)+8))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v806)+4))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	v832 = base.I64_extend_i32_u(base.I32_popcnt(v817)) + (base.I64_extend_i32_u(base.I32_popcnt(v820)) + (base.I64_extend_i32_u(base.I32_popcnt(v823)) + (v812 + base.I64_extend_i32_u(base.I32_popcnt(v826)))))
	v834 = v810 + int32(4)
	if v834 != v799&int32(2147483644) {
		v806 = v816
		v807 = v814
		v810 = v834
		v812 = v832
		goto L147
	} else {
		goto L149
	}
L148:
	;
	if v801 == int32(0) {
		v864 = v816
		v865 = v814
		v870 = v832
		goto L142
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	v838 = v816
	v839 = v814
	v844 = v832
	goto L146
L151:
	;
	v853 = int32(4)
	v854 = v847 - v853
	v856 = v846 + v853
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v846)))
	v860 = v852 + base.I64_extend_i32_u(base.I32_popcnt(v857))
	v862 = v848 + int32(1)
	if v862 != v801 {
		v846 = v856
		v847 = v854
		v848 = v862
		v852 = v860
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v864 = v856
	v865 = v854
	v870 = v860
	goto L142
L153:
	;
	goto L152
L154:
	;
	goto L141
L155:
	;
	v874 = v865 & int32(3)
	if v874 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	if base.Ui32(v865) < base.Ui32(int32(4)) {
		v933 = v901
		goto L154
	} else {
		goto L163
	}
L157:
	;
	v895 = v864
	v897 = v865
	v901 = v870
	goto L156
L158:
	;
	goto L159
L159:
	;
	v878 = v864
	v880 = v865
	v882 = int32(0)
	v884 = v870
	goto L160
L160:
	;
	v885 = int32(1)
	v886 = v878 + v885
	v888 = v880 - v885
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878))))
	v890 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v889)+uint32(_c_F_gtrgm_penalty[0]))))
	v891 = v884 + v890
	v893 = v882 + v885
	if v893 != v874 {
		v878 = v886
		v880 = v888
		v882 = v893
		v884 = v891
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v895 = v886
	v897 = v888
	v901 = v891
	goto L156
L162:
	;
	goto L161
L163:
	;
	v904 = v895
	v906 = v897
	v910 = v901
	goto L164
L164:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+3)))
	v912 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v911)+uint32(_c_F_gtrgm_penalty[0]))))
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+2)))
	v914 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v913)+uint32(_c_F_gtrgm_penalty[0]))))
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+1)))
	v916 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v915)+uint32(_c_F_gtrgm_penalty[0]))))
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904))))
	v918 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v917)+uint32(_c_F_gtrgm_penalty[0]))))
	v922 = v912 + (v914 + (v916 + (v910 + v918)))
	v923 = int32(4)
	v926 = v906 - v923
	if v926 != 0 {
		v904 = v904 + v923
		v906 = v926
		v910 = v922
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v933 = v922
	goto L154
L166:
	;
	goto L165
L167:
	;
	v1277 = v1086
	goto L14
L168:
	;
	if v1018 == int32(0) {
		v1086 = v1023
		goto L180
	} else {
		goto L181
	}
L169:
	;
	v948 = v48 - int32(4)
	v952 = int32(base.Ui32(v948)>>(uint(int32(2))%32)) + int32(1)
	v954 = v952 & int32(3)
	if base.Ui32(int32(12)) <= base.Ui32(v948) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v959 = v493
	v960 = v48
	v963 = int32(0)
	v965 = v938
	goto L173
L171:
	;
	v991 = v493
	v992 = v48
	v997 = v938
	goto L172
L172:
	;
	v999 = v991
	v1000 = v992
	v1001 = int32(0)
	v1005 = v997
	goto L177
L173:
	;
	v966 = int32(16)
	v967 = v960 - v966
	v969 = v959 + v966
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v959)+12))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v959)+8))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	v985 = base.I64_extend_i32_u(base.I32_popcnt(v970)) + (base.I64_extend_i32_u(base.I32_popcnt(v973)) + (base.I64_extend_i32_u(base.I32_popcnt(v976)) + (v965 + base.I64_extend_i32_u(base.I32_popcnt(v979)))))
	v987 = v963 + int32(4)
	if v987 != v952&int32(2147483644) {
		v959 = v969
		v960 = v967
		v963 = v987
		v965 = v985
		goto L173
	} else {
		goto L175
	}
L174:
	;
	if v954 == int32(0) {
		v1017 = v969
		v1018 = v967
		v1023 = v985
		goto L168
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v991 = v969
	v992 = v967
	v997 = v985
	goto L172
L177:
	;
	v1006 = int32(4)
	v1007 = v1000 - v1006
	v1009 = v999 + v1006
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v999)))
	v1013 = v1005 + base.I64_extend_i32_u(base.I32_popcnt(v1010))
	v1015 = v1001 + int32(1)
	if v1015 != v954 {
		v999 = v1009
		v1000 = v1007
		v1001 = v1015
		v1005 = v1013
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v1017 = v1009
	v1018 = v1007
	v1023 = v1013
	goto L168
L179:
	;
	goto L178
L180:
	;
	goto L167
L181:
	;
	v1027 = v1018 & int32(3)
	if v1027 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	if base.Ui32(v1018) < base.Ui32(int32(4)) {
		v1086 = v1054
		goto L180
	} else {
		goto L189
	}
L183:
	;
	v1048 = v1017
	v1050 = v1018
	v1054 = v1023
	goto L182
L184:
	;
	goto L185
L185:
	;
	v1031 = v1017
	v1033 = v1018
	v1035 = int32(0)
	v1037 = v1023
	goto L186
L186:
	;
	v1038 = int32(1)
	v1039 = v1031 + v1038
	v1041 = v1033 - v1038
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031))))
	v1043 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1042)+uint32(_c_F_gtrgm_penalty[0]))))
	v1044 = v1037 + v1043
	v1046 = v1035 + v1038
	if v1046 != v1027 {
		v1031 = v1039
		v1033 = v1041
		v1035 = v1046
		v1037 = v1044
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v1048 = v1039
	v1050 = v1041
	v1054 = v1044
	goto L182
L188:
	;
	goto L187
L189:
	;
	v1057 = v1048
	v1059 = v1050
	v1063 = v1054
	goto L190
L190:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+3)))
	v1065 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1064)+uint32(_c_F_gtrgm_penalty[0]))))
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+2)))
	v1067 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1066)+uint32(_c_F_gtrgm_penalty[0]))))
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+1)))
	v1069 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1068)+uint32(_c_F_gtrgm_penalty[0]))))
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057))))
	v1071 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1070)+uint32(_c_F_gtrgm_penalty[0]))))
	v1075 = v1065 + (v1067 + (v1069 + (v1063 + v1071)))
	v1076 = int32(4)
	v1079 = v1059 - v1076
	if v1079 != 0 {
		v1057 = v1057 + v1076
		v1059 = v1079
		v1063 = v1075
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v1086 = v1075
	goto L180
L192:
	;
	goto L191
L193:
	;
	v1256 = v1239
	goto L15
L194:
	;
	if v1171 == int32(0) {
		v1239 = v1176
		goto L206
	} else {
		goto L207
	}
L195:
	;
	v1101 = v48 - int32(4)
	v1105 = int32(base.Ui32(v1101)>>(uint(int32(2))%32)) + int32(1)
	v1107 = v1105 & int32(3)
	if base.Ui32(int32(12)) <= base.Ui32(v1101) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1112 = v54
	v1113 = v48
	v1116 = int32(0)
	v1118 = v1091
	goto L199
L197:
	;
	v1144 = v54
	v1145 = v48
	v1150 = v1091
	goto L198
L198:
	;
	v1152 = v1144
	v1153 = v1145
	v1154 = int32(0)
	v1158 = v1150
	goto L203
L199:
	;
	v1119 = int32(16)
	v1120 = v1113 - v1119
	v1122 = v1112 + v1119
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+12))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+8))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	v1138 = base.I64_extend_i32_u(base.I32_popcnt(v1123)) + (base.I64_extend_i32_u(base.I32_popcnt(v1126)) + (base.I64_extend_i32_u(base.I32_popcnt(v1129)) + (v1118 + base.I64_extend_i32_u(base.I32_popcnt(v1132)))))
	v1140 = v1116 + int32(4)
	if v1140 != v1105&int32(2147483644) {
		v1112 = v1122
		v1113 = v1120
		v1116 = v1140
		v1118 = v1138
		goto L199
	} else {
		goto L201
	}
L200:
	;
	if v1107 == int32(0) {
		v1170 = v1122
		v1171 = v1120
		v1176 = v1138
		goto L194
	} else {
		goto L202
	}
L201:
	;
	goto L200
L202:
	;
	v1144 = v1122
	v1145 = v1120
	v1150 = v1138
	goto L198
L203:
	;
	v1159 = int32(4)
	v1160 = v1153 - v1159
	v1162 = v1152 + v1159
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	v1166 = v1158 + base.I64_extend_i32_u(base.I32_popcnt(v1163))
	v1168 = v1154 + int32(1)
	if v1168 != v1107 {
		v1152 = v1162
		v1153 = v1160
		v1154 = v1168
		v1158 = v1166
		goto L203
	} else {
		goto L205
	}
L204:
	;
	v1170 = v1162
	v1171 = v1160
	v1176 = v1166
	goto L194
L205:
	;
	goto L204
L206:
	;
	goto L193
L207:
	;
	v1180 = v1171 & int32(3)
	if v1180 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if base.Ui32(v1171) < base.Ui32(int32(4)) {
		v1239 = v1207
		goto L206
	} else {
		goto L215
	}
L209:
	;
	v1201 = v1170
	v1203 = v1171
	v1207 = v1176
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1184 = v1170
	v1186 = v1171
	v1188 = int32(0)
	v1190 = v1176
	goto L212
L212:
	;
	v1191 = int32(1)
	v1192 = v1184 + v1191
	v1194 = v1186 - v1191
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184))))
	v1196 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1195)+uint32(_c_F_gtrgm_penalty[0]))))
	v1197 = v1190 + v1196
	v1199 = v1188 + v1191
	if v1199 != v1180 {
		v1184 = v1192
		v1186 = v1194
		v1188 = v1199
		v1190 = v1197
		goto L212
	} else {
		goto L214
	}
L213:
	;
	v1201 = v1192
	v1203 = v1194
	v1207 = v1197
	goto L208
L214:
	;
	goto L213
L215:
	;
	v1210 = v1201
	v1212 = v1203
	v1216 = v1207
	goto L216
L216:
	;
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+3)))
	v1218 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+uint32(_c_F_gtrgm_penalty[0]))))
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+2)))
	v1220 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_c_F_gtrgm_penalty[0]))))
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+1)))
	v1222 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1221)+uint32(_c_F_gtrgm_penalty[0]))))
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210))))
	v1224 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1223)+uint32(_c_F_gtrgm_penalty[0]))))
	v1228 = v1218 + (v1220 + (v1222 + (v1216 + v1224)))
	v1229 = int32(4)
	v1232 = v1212 - v1229
	if v1232 != 0 {
		v1210 = v1210 + v1229
		v1212 = v1232
		v1216 = v1228
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v1239 = v1228
	goto L206
L218:
	;
	goto L217
}
func F_gtrgm_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == v2 {
		v39 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v26 == int32(0) {
		v39 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(7) {
		v39 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 != int32(17) {
		v39 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v39 = v35 ^ int32(1)
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_fn_opclass_options(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = int32(12)
	goto L9
L9:
	;
	v50 = v48 + int32(5)
	v51 = F_palloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = v47
	goto L9
L12:
	;
	v53 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v53)
	v56 = v50 << (uint(v53) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
	v59 = v51 + int32(5)
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryFill(m, v59, int32(0), v48)
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v20 <= int32(0) {
		v297 = v56
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v297) >> (uint(int32(2)) % 32))
	return v51
L17:
	;
	v66 = int32(3)
	v67 = v48 & v66
	v86 = v2
	goto L18
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v86<<(uint(int32(4))%32))))
	v96 = v94 + int32(5)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v97&int32(2) != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v297 = v56
	goto L16
L20:
	;
	v285 = v86 + int32(1)
	if v285 != v20 {
		v86 = v285
		goto L18
	} else {
		goto L41
	}
L21:
	;
	v262 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v262
	v265 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v265)
	v297 = v262
	goto L16
L22:
	;
	if v97&int32(4) != 0 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if base.Ui32(int32(base.Ui32(v209)>>(uint(int32(2))%32))-int32(5)) < base.Ui32(int32(3)) {
		goto L20
	} else {
		goto L37
	}
L25:
	;
	if v48 <= int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v104 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v109 = v104
	v113 = v104
	goto L30
L28:
	;
	v163 = v104
	goto L29
L29:
	;
	v180 = v163
	v186 = v104
	goto L34
L30:
	;
	v126 = v109 + v59
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v96))))
	v130 = v127 | v129
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v130)
	v133 = v109 | int32(1)
	v134 = v59 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v96))))
	v138 = v135 | v137
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v138)
	v141 = v109 | int32(2)
	v142 = v59 + v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v96))))
	v146 = v143 | v145
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v146)
	v149 = v109 | int32(3)
	v150 = v59 + v149
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v96))))
	v154 = v151 | v153
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v154)
	v156 = int32(4)
	v157 = v109 + v156
	v159 = v113 + v156
	if v159 != v48&int32(2147483644) {
		v109 = v157
		v113 = v159
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v67 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v163 = v157
	goto L29
L34:
	;
	v197 = v180 + v59
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v96))))
	v201 = v198 | v200
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v201)
	v203 = int32(1)
	v206 = v186 + v203
	if v206 != v67 {
		v180 = v180 + v203
		v186 = v206
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L20
L36:
	;
	goto L35
L37:
	;
	v216 = int32(0)
	goto L38
L38:
	;
	v233 = int32(3)
	v235 = v96 + v216*v233
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+2)))
	v241 = base.I32_rem_u_s(v236|v237<<(uint(int32(16))%32), v48<<(uint(v66)%32)-int32(1))
	v244 = v59 + int32(base.Ui32(v241)>>(uint(v233)%32))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v246 = int32(1)
	v250 = v245 | v246<<(uint(v241&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v244))) = uint8(v250)
	v253 = v216 + v246
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v260 = base.I32_div_u_s(int32(base.Ui32(v254)>>(uint(int32(2))%32))-int32(5), v233)
	if base.Ui32(v253) < base.Ui32(v260) {
		v216 = v253
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L20
L40:
	;
	goto L39
L41:
	;
	goto L19
}
