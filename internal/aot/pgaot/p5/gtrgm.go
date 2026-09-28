package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 float32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v280 float64
	_ = v280
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = int32(0)
	if v29 == v30 {
		v46 = v30
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v46&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v33 == int32(0) {
		v46 = v30
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v36 != int32(7) {
		v46 = v30
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v39 != int32(17) {
		v46 = v30
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
	v46 = v42 ^ int32(1)
	goto L4
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = F_get_fn_opclass_options(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v57 = int32(95)
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v61 = int32(base.Ui32(v59) >> (uint(int32(2)) % 32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v57 = v52<<(uint(int32(3))%32) - int32(1)
	goto L11
L13:
	;
	v170 = base.I32_wrap_i64(v27) & int32(_a_F_gtrgm_distance_0)
	v177 = int32(0)
	if base.B2i32(base.Ui32(int32(10)) < base.Ui32(v170))|base.B2i32(int32(1)<<(uint(v170)%32)&int32(1284) == v177) == v177 {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v137 = int32(4)
	v141 = F_generate_trgm(m, v22+v137, v61-v137)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if int32(base.Ui32(v67)>>(uint(int32(2))%32)) != v61 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v61) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v132 != 0 {
		goto L14
	} else {
		goto L35
	}
L18:
	;
	v132 = int32(0)
	goto L17
L19:
	;
	v106 = v101
	v107 = v102
	v108 = v103
	goto L29
L20:
	;
	if (v64|v22)&int32(3) != 0 {
		v101 = v64
		v102 = v22
		v103 = v61
		goto L19
	} else {
		goto L23
	}
L21:
	;
	v94 = v64
	v95 = v22
	v96 = v61
	goto L22
L22:
	;
	if v96 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L23:
	;
	v78 = v64
	v79 = v22
	v80 = v61
	goto L24
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v83 != v84 {
		v101 = v78
		v102 = v79
		v103 = v80
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v94 = v89
	v95 = v87
	v96 = v91
	goto L22
L26:
	;
	v86 = int32(4)
	v87 = v79 + v86
	v89 = v78 + v86
	v91 = v80 - v86
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		v78 = v89
		v79 = v87
		v80 = v91
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v101 = v94
	v102 = v95
	v103 = v96
	goto L19
L29:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v111 == v112 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v132 = v111 - v112
	goto L17
L31:
	;
	v114 = int32(1)
	v119 = v108 - v114
	if v119 != 0 {
		v106 = v106 + v114
		v107 = v107 + v114
		v108 = v119
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
	v166 = v64
	v167 = (v61 + int32(7)) & int32(2147483640)
	goto L13
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v148 = (v61 + int32(7)) & int32(2147483640)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v153 = F_MemoryContextAlloc(m, v144, v148+int32(base.Ui32(v149)>>(uint(int32(2))%32)))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v61 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	base.MemoryCopy(m, v153, v22, v61)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v158 = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
	if v158 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	base.MemoryCopy(m, v153+v148, v141, v158)
	goto L43
L42:
	;
	goto L43
L43:
	;
	if v64 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_pfree(m, v64)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v153
	v166 = v153
	v167 = v148
	goto L13
L47:
	;
	goto L46
L48:
	;
	v182 = v166 + v167
	v184 = base.B2i32(v170 != int32(2))
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+16)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v187)+12)))
	if v189&int32(1) != 0 {
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
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L64
	}
L51:
	;
	m.G0 = v18 + int32(16)
	return base.I64_reinterpret_f64(v280)
L52:
	;
	v192 = F_cnt_sml(m, v182, v62, v184)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)))
	if v200&int32(4) != 0 {
		v280 = float64(0)
		goto L51
	} else {
		goto L56
	}
L55:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18)+12)) = v192
	v280 = base.F64_sub(float64(1), base.F64_promote_f32(v192))
	goto L51
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v208 = int32(base.Ui32(v204)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v208) < base.Ui32(int32(3)) {
		v280 = float64(-1)
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v211 = int32(5)
	v215 = int32(1)
	v217 = base.I32_div_u_s(v208, int32(3))
	if base.Ui32(v217) <= base.Ui32(v215) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v220 = v215
	goto L60
L59:
	;
	v220 = v217
	goto L60
L60:
	;
	v221 = int32(0)
	v223 = v221
	v224 = v221
	goto L61
L61:
	;
	v238 = int32(3)
	v240 = v182 + v211 + v223*v238
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240))))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+2)))
	v246 = base.I32_rem_u_s(v241|v242<<(uint(int32(16))%32), v57)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v211+int32(base.Ui32(v246)>>(uint(v238)%32))))))
	v254 = int32(1)
	v256 = v224 + int32(base.Ui32(v250)>>(uint(v246&int32(7))%32))&v254
	v258 = v223 + v254
	if v258 != v220 {
		v223 = v258
		v224 = v256
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v280 = base.F64_sub(float64(1), base.F64_div(base.F64_convert_i32_u(v256), base.F64_convert_i32_u(v217)))
	goto L51
L63:
	;
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v170
	F_errmsg_internal(m, int32(_a_F_gtrgm_distance_1), v18)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_gtrgm_distance_2), int32(527), int32(_a_F_gtrgm_distance_3))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gtrgm_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_gtrgm_in_0), int32(61), int32(_a_F_gtrgm_in_1), int32(_a_F_gtrgm_in_2), int32(_a_F_gtrgm_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_gtrgm_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
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
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v322 int64
	_ = v322
	var v323 int32
	_ = v323
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v334 int64
	_ = v334
	var v338 int64
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v362 int64
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v381 int64
	_ = v381
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
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
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v538 int64
	_ = v538
	var v540 int32
	_ = v540
	var v543 int64
	_ = v543
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v551 int64
	_ = v551
	var v552 int32
	_ = v552
	var v555 int64
	_ = v555
	var v559 int64
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v583 int64
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v601 int64
	_ = v601
	var v603 int32
	_ = v603
	var v606 int64
	_ = v606
	var v607 int64
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v643 int64
	_ = v643
	var v645 int32
	_ = v645
	var v648 int64
	_ = v648
	var v649 int32
	_ = v649
	var v652 int64
	_ = v652
	var v653 int32
	_ = v653
	var v656 int64
	_ = v656
	var v657 int32
	_ = v657
	var v660 int64
	_ = v660
	var v664 int64
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v688 int64
	_ = v688
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v707 int64
	_ = v707
	var v709 int32
	_ = v709
	var v712 int64
	_ = v712
	var v713 int64
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v802 int64
	_ = v802
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int64
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int64
	_ = v824
	var v825 int32
	_ = v825
	var v826 int64
	_ = v826
	var v827 int32
	_ = v827
	var v828 int64
	_ = v828
	var v829 int32
	_ = v829
	var v830 int64
	_ = v830
	var v834 int64
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int64
	_ = v842
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int64
	_ = v849
	var v853 int32
	_ = v853
	var v854 int64
	_ = v854
	var v855 int64
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v863 int64
	_ = v863
	var v873 int64
	_ = v873
	var v874 int64
	_ = v874
	var v875 int32
	_ = v875
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v889 int64
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int64
	_ = v896
	var v897 int32
	_ = v897
	var v898 int64
	_ = v898
	var v899 int32
	_ = v899
	var v900 int64
	_ = v900
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v906 int64
	_ = v906
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v914 int64
	_ = v914
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int64
	_ = v921
	var v925 int32
	_ = v925
	var v926 int64
	_ = v926
	var v927 int64
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int64
	_ = v935
	var v945 int64
	_ = v945
	var v946 int64
	_ = v946
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v961 int64
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int64
	_ = v968
	var v969 int32
	_ = v969
	var v970 int64
	_ = v970
	var v971 int32
	_ = v971
	var v972 int64
	_ = v972
	var v973 int32
	_ = v973
	var v974 int64
	_ = v974
	var v978 int64
	_ = v978
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v986 int64
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int64
	_ = v993
	var v997 int32
	_ = v997
	var v998 int64
	_ = v998
	var v999 int64
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1007 int64
	_ = v1007
	var v1017 int64
	_ = v1017
	var v1034 int64
	_ = v1034
	var v1056 int64
	_ = v1056
	var v1065 int32
	_ = v1065
	var v1099 int64
	_ = v1099
	v2 = int32(0)
	v17 = int64(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v22 = base.I32_wrap_i64(v21)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v24 == v2 {
		v41 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v41&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v28 == int32(0) {
		v41 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v31 != int32(7) {
		v41 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v34 != int32(17) {
		v41 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+32)))
	v41 = v37 ^ int32(1)
	goto L2
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = F_get_fn_opclass_options(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v50 = int32(12)
	goto L9
L9:
	;
	v52 = v21 & int64(4294967295)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(0)
	v58 = v54 + int32(5)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v59&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	return int64(0)
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v50 = v49
	goto L9
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v22))) = base.F32_div(base.F32_convert_i32_s(v287+(base.I32_wrap_i64(v1099)^int32(-1))), base.F32_convert_i32_s(v287))
	return v52
L13:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v22))) = base.F32_convert_i32_s(v1065)
	return v52
L14:
	;
	v1065 = v50<<(uint(int32(3))%32) + (base.I32_wrap_i64(v1056) ^ int32(-1))
	goto L13
L15:
	;
	v1065 = v50<<(uint(int32(3))%32) + (base.I32_wrap_i64(v1034) ^ int32(-1))
	goto L13
L16:
	;
	v946 = int64(0)
	v947 = int32(0)
	if v50 == v947 {
		goto L178
	} else {
		goto L179
	}
L17:
	;
	v874 = int64(0)
	v875 = int32(0)
	if v50 == v875 {
		goto L163
	} else {
		goto L164
	}
L18:
	;
	v802 = int64(0)
	v803 = int32(0)
	if v50 == v803 {
		goto L148
	} else {
		goto L149
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v64 = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
	v68 = (v50 + int32(7)) & int32(-8)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v70 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v500 = int32(4)
	v501 = v59 & v500
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)))
	if v502&v500 != 0 {
		goto L100
	} else {
		goto L101
	}
L22:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)))
	if v283&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L23:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v143 = F_MemoryContextAlloc(m, v141, v64+v68)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L45
	}
L24:
	;
	v73 = v70 + v68
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if int32(base.Ui32(v74)>>(uint(int32(2))%32)) != v64 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v64) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v139 != 0 {
		goto L23
	} else {
		goto L44
	}
L27:
	;
	v139 = int32(0)
	goto L26
L28:
	;
	v113 = v108
	v114 = v109
	v115 = v110
	goto L38
L29:
	;
	if (v73|v53)&int32(3) != 0 {
		v108 = v73
		v109 = v53
		v110 = v64
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v101 = v73
	v102 = v53
	v103 = v64
	goto L31
L31:
	;
	if v103 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v85 = v73
	v86 = v53
	v87 = v64
	goto L33
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v90 != v91 {
		v108 = v85
		v109 = v86
		v110 = v87
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v101 = v96
	v102 = v94
	v103 = v98
	goto L31
L35:
	;
	v93 = int32(4)
	v94 = v86 + v93
	v96 = v85 + v93
	v98 = v87 - v93
	if base.Ui32(int32(3)) < base.Ui32(v98) {
		v85 = v96
		v86 = v94
		v87 = v98
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v108 = v101
	v109 = v102
	v110 = v103
	goto L28
L38:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v118 == v119 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v139 = v118 - v119
	goto L26
L40:
	;
	v121 = int32(1)
	v126 = v115 - v121
	if v126 != 0 {
		v113 = v113 + v121
		v114 = v114 + v121
		v115 = v126
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
	v266 = v70
	goto L22
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v149 = int32(base.Ui32(v145)>>(uint(int32(2))%32)) - int32(5)
	v150 = int32(3)
	v151 = base.I32_div_u_s(v149, v150)
	if v143&v150 != 0 {
		v174 = v50
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v181 = int32(3)
	v184 = v50<<(uint(v181)%32) - int32(1)
	v186 = base.I32_div_s(v184, int32(8))
	v187 = v143 + v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v190 = v188 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v190)
	if base.Ui32(v181) <= base.Ui32(v149) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	if v174 == int32(0) {
		goto L46
	} else {
		goto L55
	}
L48:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v50) {
		v174 = v50
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v50&int32(3) != 0 {
		v174 = v50
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v50 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v162 = v143 + v50
	v164 = v143 + int32(4)
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v166 = v162
	goto L54
L53:
	;
	v166 = v164
	goto L54
L54:
	;
	v174 = (v143^int32(-1)+v166)&int32(-4) + int32(4)
	goto L47
L55:
	;
	base.MemoryFill(m, v143, int32(0), v174)
	goto L46
L56:
	;
	v196 = int32(1)
	if base.Ui32(v151) <= base.Ui32(v196) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v64 != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v199 = v196
	goto L61
L60:
	;
	v199 = v151
	goto L61
L61:
	;
	v204 = int32(0)
	goto L62
L62:
	;
	v219 = int32(3)
	v221 = v53 + int32(5) + v204*v219
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)))
	v227 = base.I32_rem_u_s(v222|v223<<(uint(int32(16))%32), v184)
	v230 = v143 + int32(base.Ui32(v227)>>(uint(v219)%32))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v232 = int32(1)
	v236 = v231 | v232<<(uint(v227&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v236)
	v239 = v204 + v232
	if v239 != v199 {
		v204 = v239
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
	base.MemoryCopy(m, v143+v68, v53, v64)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if v70 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_pfree(m, v70)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = v143
	v266 = v143
	goto L22
L71:
	;
	goto L70
L72:
	;
	v287 = v50 << (uint(int32(3)) % 32)
	if int32(7) < v50 {
		goto L18
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v50 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L75:
	;
	if v50 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v1099 = v17
	goto L12
L77:
	;
	goto L78
L78:
	;
	v292 = int32(3)
	v293 = v50 & v292
	if base.Ui32(v292) <= base.Ui32(v50-int32(1)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v302 = v266
	v306 = int32(0)
	v317 = v17
	goto L82
L80:
	;
	v347 = v266
	v362 = v17
	goto L81
L81:
	;
	v366 = v347
	v368 = int32(0)
	v381 = v362
	goto L86
L82:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+3)))
	v322 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_gtrgm_penalty[0]))))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+2)))
	v326 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_gtrgm_penalty[0]))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v330 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v327)+uint32(_c_F_gtrgm_penalty[0]))))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v334 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v331)+uint32(_c_F_gtrgm_penalty[0]))))
	v338 = v322 + (v326 + (v330 + (v317 + v334)))
	v339 = int32(4)
	v340 = v302 + v339
	v342 = v306 + v339
	if v342 != v50&int32(-4) {
		v302 = v340
		v306 = v342
		v317 = v338
		goto L82
	} else {
		goto L84
	}
L83:
	;
	if v293 == int32(0) {
		v1099 = v338
		goto L12
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	v347 = v340
	v362 = v338
	goto L81
L86:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	v386 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v383)+uint32(_c_F_gtrgm_penalty[0]))))
	v387 = v381 + v386
	v388 = int32(1)
	v391 = v368 + v388
	if v391 != v293 {
		v366 = v366 + v388
		v368 = v391
		v381 = v387
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v1099 = v387
	goto L12
L88:
	;
	goto L87
L89:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v22))) = float32(0)
	return v52
L90:
	;
	goto L91
L91:
	;
	v398 = int32(0)
	if v50 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v22))) = base.F32_convert_i32_u(v479)
	return v52
L93:
	;
	v407 = v398
	v410 = v398
	v415 = int32(0)
	goto L96
L94:
	;
	v452 = v398
	v455 = v398
	goto L95
L95:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v455))))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455+v58))))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471^v473)+uint32(_c_F_gtrgm_penalty[0]))))
	v479 = v452 + v477
	goto L92
L96:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v410))))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v58))))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426^v428)+uint32(_c_F_gtrgm_penalty[0]))))
	v435 = v410 | int32(1)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v435))))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v266))))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437^v439)+uint32(_c_F_gtrgm_penalty[0]))))
	v444 = v407 + v432 + v443
	v445 = int32(2)
	v446 = v410 + v445
	v448 = v415 + v445
	if v448 != v50&int32(2147483646) {
		v407 = v444
		v410 = v446
		v415 = v448
		goto L96
	} else {
		goto L98
	}
L97:
	;
	if v50&int32(1) == int32(0) {
		v479 = v444
		goto L92
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v452 = v444
	v455 = v446
	goto L95
L100:
	;
	if v501 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	if v501 != 0 {
		goto L120
	} else {
		goto L121
	}
L103:
	;
	v1065 = v2
	goto L13
L104:
	;
	goto L105
L105:
	;
	v506 = v53 + int32(5)
	if int32(7) < v50 {
		goto L17
	} else {
		goto L106
	}
L106:
	;
	if v50 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v1056 = v17
	goto L14
L108:
	;
	goto L109
L109:
	;
	v513 = int32(3)
	v514 = v50 & v513
	if base.Ui32(v513) <= base.Ui32(v50-int32(1)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v523 = v506
	v527 = int32(0)
	v538 = v17
	goto L113
L111:
	;
	v568 = v506
	v583 = v17
	goto L112
L112:
	;
	v586 = v568
	v588 = v2
	v601 = v583
	goto L117
L113:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+3)))
	v543 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v540)+uint32(_c_F_gtrgm_penalty[0]))))
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+2)))
	v547 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v544)+uint32(_c_F_gtrgm_penalty[0]))))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+1)))
	v551 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v548)+uint32(_c_F_gtrgm_penalty[0]))))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	v555 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_gtrgm_penalty[0]))))
	v559 = v543 + (v547 + (v551 + (v538 + v555)))
	v560 = int32(4)
	v561 = v523 + v560
	v563 = v527 + v560
	if v563 != v50&int32(-4) {
		v523 = v561
		v527 = v563
		v538 = v559
		goto L113
	} else {
		goto L115
	}
L114:
	;
	if v514 == int32(0) {
		v1056 = v559
		goto L14
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v568 = v561
	v583 = v559
	goto L112
L117:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	v606 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v603)+uint32(_c_F_gtrgm_penalty[0]))))
	v607 = v601 + v606
	v608 = int32(1)
	v611 = v588 + v608
	if v611 != v514 {
		v586 = v586 + v608
		v588 = v611
		v601 = v607
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v1056 = v607
	goto L14
L119:
	;
	goto L118
L120:
	;
	if int32(7) < v50 {
		goto L16
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if v50 <= int32(0) {
		goto L137
	} else {
		goto L138
	}
L123:
	;
	if v50 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v1034 = v17
	goto L15
L125:
	;
	goto L126
L126:
	;
	v619 = int32(3)
	v620 = v50 & v619
	if base.Ui32(v619) <= base.Ui32(v50-int32(1)) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v630 = v2
	v632 = v58
	v643 = v17
	goto L130
L128:
	;
	v677 = v58
	v688 = v17
	goto L129
L129:
	;
	v692 = int32(0)
	v696 = v677
	v707 = v688
	goto L134
L130:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+3)))
	v648 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v645)+uint32(_c_F_gtrgm_penalty[0]))))
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+2)))
	v652 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_gtrgm_penalty[0]))))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+1)))
	v656 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v653)+uint32(_c_F_gtrgm_penalty[0]))))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	v660 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v657)+uint32(_c_F_gtrgm_penalty[0]))))
	v664 = v648 + (v652 + (v656 + (v643 + v660)))
	v665 = int32(4)
	v666 = v632 + v665
	v668 = v630 + v665
	if v668 != v50&int32(-4) {
		v630 = v668
		v632 = v666
		v643 = v664
		goto L130
	} else {
		goto L132
	}
L131:
	;
	if v620 == int32(0) {
		v1034 = v664
		goto L15
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v677 = v666
	v688 = v664
	goto L129
L134:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	v712 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v709)+uint32(_c_F_gtrgm_penalty[0]))))
	v713 = v707 + v712
	v714 = int32(1)
	v717 = v692 + v714
	if v717 != v620 {
		v692 = v717
		v696 = v696 + v714
		v707 = v713
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v1034 = v713
	goto L15
L136:
	;
	goto L135
L137:
	;
	v1065 = v2
	goto L13
L138:
	;
	goto L139
L139:
	;
	v722 = v53 + int32(5)
	v723 = int32(0)
	if v50 != int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v731 = v723
	v733 = v2
	v738 = v2
	goto L143
L141:
	;
	v776 = v723
	v778 = v2
	goto L142
L142:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776+v722))))
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776+v58))))
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794^v796)+uint32(_c_F_gtrgm_penalty[0]))))
	v1065 = v778 + v800
	goto L13
L143:
	;
	v749 = v731 | int32(1)
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722+v749))))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749+v58))))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751^v753)+uint32(_c_F_gtrgm_penalty[0]))))
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+v58))))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+v722))))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759^v761)+uint32(_c_F_gtrgm_penalty[0]))))
	v767 = v757 + (v733 + v765)
	v768 = int32(2)
	v769 = v731 + v768
	v771 = v738 + v768
	if v771 != v50&int32(2147483646) {
		v731 = v769
		v733 = v767
		v738 = v771
		goto L143
	} else {
		goto L145
	}
L144:
	;
	if v50&int32(1) == int32(0) {
		v1065 = v767
		goto L13
	} else {
		goto L146
	}
L145:
	;
	goto L144
L146:
	;
	v776 = v769
	v778 = v767
	goto L142
L147:
	;
	v1099 = v873
	goto L12
L148:
	;
	v873 = int64(0)
	goto L147
L149:
	;
	goto L150
L150:
	;
	v810 = v50 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v873 = v863
	goto L147
L152:
	;
	v815 = v266
	v817 = v802
	v820 = v803
	goto L155
L153:
	;
	v840 = v266
	v842 = v802
	goto L154
L154:
	;
	v847 = v840
	v848 = int32(0)
	v849 = v842
	goto L159
L155:
	;
	v821 = int32(4)
	v822 = v815 + v821
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+3)))
	v824 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v823)+uint32(_c_F_gtrgm_penalty[0]))))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+2)))
	v826 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v825)+uint32(_c_F_gtrgm_penalty[0]))))
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+1)))
	v828 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v827)+uint32(_c_F_gtrgm_penalty[0]))))
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815))))
	v830 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v829)+uint32(_c_F_gtrgm_penalty[0]))))
	v834 = v824 + (v826 + (v828 + (v817 + v830)))
	v836 = v820 + v821
	if v836 != v50&int32(-4) {
		v815 = v822
		v817 = v834
		v820 = v836
		goto L155
	} else {
		goto L157
	}
L156:
	;
	if v810 == int32(0) {
		v863 = v834
		goto L151
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	v840 = v822
	v842 = v834
	goto L154
L159:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	v854 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v853)+uint32(_c_F_gtrgm_penalty[0]))))
	v855 = v849 + v854
	v856 = int32(1)
	v859 = v848 + v856
	if v859 != v810 {
		v847 = v847 + v856
		v848 = v859
		v849 = v855
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v863 = v855
	goto L151
L161:
	;
	goto L160
L162:
	;
	v1056 = v945
	goto L14
L163:
	;
	v945 = int64(0)
	goto L162
L164:
	;
	goto L165
L165:
	;
	v882 = v50 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v945 = v935
	goto L162
L167:
	;
	v887 = v506
	v889 = v874
	v892 = v875
	goto L170
L168:
	;
	v912 = v506
	v914 = v874
	goto L169
L169:
	;
	v919 = v912
	v920 = int32(0)
	v921 = v914
	goto L174
L170:
	;
	v893 = int32(4)
	v894 = v887 + v893
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+3)))
	v896 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v895)+uint32(_c_F_gtrgm_penalty[0]))))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+2)))
	v898 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v897)+uint32(_c_F_gtrgm_penalty[0]))))
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+1)))
	v900 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v899)+uint32(_c_F_gtrgm_penalty[0]))))
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v902 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v901)+uint32(_c_F_gtrgm_penalty[0]))))
	v906 = v896 + (v898 + (v900 + (v889 + v902)))
	v908 = v892 + v893
	if v908 != v50&int32(-4) {
		v887 = v894
		v889 = v906
		v892 = v908
		goto L170
	} else {
		goto L172
	}
L171:
	;
	if v882 == int32(0) {
		v935 = v906
		goto L166
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v912 = v894
	v914 = v906
	goto L169
L174:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919))))
	v926 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v925)+uint32(_c_F_gtrgm_penalty[0]))))
	v927 = v921 + v926
	v928 = int32(1)
	v931 = v920 + v928
	if v931 != v882 {
		v919 = v919 + v928
		v920 = v931
		v921 = v927
		goto L174
	} else {
		goto L176
	}
L175:
	;
	v935 = v927
	goto L166
L176:
	;
	goto L175
L177:
	;
	v1034 = v1017
	goto L15
L178:
	;
	v1017 = int64(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v954 = v50 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v1017 = v1007
	goto L177
L182:
	;
	v959 = v58
	v961 = v946
	v964 = v947
	goto L185
L183:
	;
	v984 = v58
	v986 = v946
	goto L184
L184:
	;
	v991 = v984
	v992 = int32(0)
	v993 = v986
	goto L189
L185:
	;
	v965 = int32(4)
	v966 = v959 + v965
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+3)))
	v968 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v967)+uint32(_c_F_gtrgm_penalty[0]))))
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+2)))
	v970 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v969)+uint32(_c_F_gtrgm_penalty[0]))))
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+1)))
	v972 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v971)+uint32(_c_F_gtrgm_penalty[0]))))
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959))))
	v974 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_c_F_gtrgm_penalty[0]))))
	v978 = v968 + (v970 + (v972 + (v961 + v974)))
	v980 = v964 + v965
	if v980 != v50&int32(-4) {
		v959 = v966
		v961 = v978
		v964 = v980
		goto L185
	} else {
		goto L187
	}
L186:
	;
	if v954 == int32(0) {
		v1007 = v978
		goto L181
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v984 = v966
	v986 = v978
	goto L184
L189:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991))))
	v998 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v997)+uint32(_c_F_gtrgm_penalty[0]))))
	v999 = v993 + v998
	v1000 = int32(1)
	v1003 = v992 + v1000
	if v1003 != v954 {
		v991 = v991 + v1000
		v992 = v1003
		v993 = v999
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v1007 = v999
	goto L181
L191:
	;
	goto L190
}
func F_gtrgm_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
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
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+32)))
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
	return int64(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v18)))) = int32(base.Ui32(v297) >> (uint(int32(2)) % 32))
	return base.I64_extend_i32_u(v51)
L17:
	;
	v66 = int32(3)
	v67 = v48 & v66
	v86 = v2
	goto L18
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(8)+v86*int32(24))))
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
