package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 float32
	_ = v28
	var v29 float32
	_ = v29
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_inner_int_union(m, v11, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		F_rt__int_size(m, v14, v7+int32(12))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			F_rt__int_size(m, v22, v7+int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v28 = *(*float32)(unsafe.Add(mBase, uint32(v7)+12))
				v29 = *(*float32)(unsafe.Add(mBase, uint32(v7)+8))
				*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v9)))) = base.F32_sub(v28, v29)
				F_pfree(m, v14)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					m.G0 = v7 + int32(16)
					return v9 & int64(4294967295)
				}
			}
		}
	}
}
func F_g_int_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v70 float32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 float32
	_ = v90
	var v91 float32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 float32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 float32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v129 float32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 float32
	_ = v150
	var v151 float32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 float32
	_ = v156
	var v157 int32
	_ = v157
	var v158 float32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v184 float32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 float32
	_ = v323
	var v324 float32
	_ = v324
	var v326 float32
	_ = v326
	var v327 float32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 float32
	_ = v409
	var v410 float32
	_ = v410
	var v413 float32
	_ = v413
	var v414 float32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 float32
	_ = v431
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 float32
	_ = v444
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = base.I32_wrap_i64(v26)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = (v29 + int32(_a_F_g_int_picksplit_0)) & int32(_a_F_g_int_picksplit_1)
	v37 = v33<<(uint(int32(1))%32) + int32(4)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	v43 = F_palloc(m, v37)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v43
	if base.Ui32(int32(2)) <= base.Ui32(v33) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v49 = v28 + int32(8)
	v50 = int32(1)
	v52 = v50
	v54 = v50
	v57 = v2
	v60 = v2
	v70 = float32(0)
	goto L7
L5:
	;
	v192 = v43
	v195 = v2
	v198 = v2
	goto L6
L6:
	;
	v213 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v213
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v219 = v28 + int32(8)
	v221 = int32(_a_F_g_int_picksplit_1)
	v229 = base.B2i32(v195&v221 == v213) | base.B2i32(v198&v221 == v213)
	if v229 != 0 {
		goto L46
	} else {
		goto L47
	}
L7:
	;
	v75 = v49 + v54*int32(24)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	v78 = F_inner_int_union(m, v76, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v192 = v189
	v195 = v171
	v198 = v174
	goto L6
L9:
	;
	F_rt__int_size(m, v78, v24+int32(20))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v84 = F_inner_int_inter(m, v76, v77)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_rt__int_size(m, v84, v24+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v90 = *(*float32)(unsafe.Add(mBase, uint32(v24)+16))
	v91 = *(*float32)(unsafe.Add(mBase, uint32(v24)+20))
	F_pfree(m, v78)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_pfree(m, v84)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v96 = int32(1)
	v97 = v54 + v96
	v98 = base.F32_sub(v91, v90)
	v102 = (base.F32_gt(v98, v70) | v52) & v96
	if v102 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v103 = v97
	goto L17
L16:
	;
	v103 = v60
	goto L17
L17:
	;
	if v102 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v104 = v98
	goto L20
L19:
	;
	v104 = v70
	goto L20
L20:
	;
	if v102 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v105 = v54
	goto L23
L22:
	;
	v105 = v57
	goto L23
L23:
	;
	v107 = v54 + int32(2)
	if base.Ui32(v107&int32(_a_F_g_int_picksplit_1)) <= base.Ui32(v33) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v111 = v107
	v116 = v105
	v119 = v103
	v129 = v104
	goto L27
L25:
	;
	v171 = v105
	v174 = v103
	v184 = v104
	goto L26
L26:
	;
	if v97 != v33 {
		v52 = int32(0)
		v54 = v97
		v57 = v171
		v60 = v174
		v70 = v184
		goto L7
	} else {
		goto L45
	}
L27:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v49+v111&int32(_a_F_g_int_picksplit_1)*int32(24))))
	v138 = F_inner_int_union(m, v76, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v171 = v160
	v174 = v159
	v184 = v158
	goto L26
L29:
	;
	F_rt__int_size(m, v138, v24+int32(20))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v144 = F_inner_int_inter(m, v76, v137)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_rt__int_size(m, v144, v24+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v150 = *(*float32)(unsafe.Add(mBase, uint32(v24)+16))
	v151 = *(*float32)(unsafe.Add(mBase, uint32(v24)+20))
	F_pfree(m, v138)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v144)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v156 = base.F32_sub(v151, v150)
	v157 = base.F32_gt(v156, v129)
	if v157 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v158 = v156
	goto L37
L36:
	;
	v158 = v129
	goto L37
L37:
	;
	if v157 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v159 = v111
	goto L40
L39:
	;
	v159 = v119
	goto L40
L40:
	;
	if v157 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v160 = v54
	goto L43
L42:
	;
	v160 = v116
	goto L43
L43:
	;
	v162 = v111 + int32(1)
	if base.Ui32(v162&int32(_a_F_g_int_picksplit_1)) <= base.Ui32(v33) {
		v111 = v162
		v116 = v160
		v119 = v159
		v129 = v158
		goto L27
	} else {
		goto L44
	}
L44:
	;
	goto L28
L45:
	;
	goto L8
L46:
	;
	v230 = int32(1)
	goto L48
L47:
	;
	v230 = v195
	goto L48
L48:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v219+v230&int32(_a_F_g_int_picksplit_1)*int32(24))))
	v237 = F_copy_intArrayType(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_rt__int_size(m, v237, v24+int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v229 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v244 = int32(2)
	goto L53
L52:
	;
	v244 = v198
	goto L53
L53:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v219+v244&int32(_a_F_g_int_picksplit_1)*int32(24))))
	v251 = F_copy_intArrayType(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_rt__int_size(m, v251, v24+int32(8))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v258 = int32(_a_F_g_int_picksplit_1)
	v259 = v29 + v258
	v261 = v259 & v258
	v262 = F_palloc_mul(m, int32(8), v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v29&int32(_a_F_g_int_picksplit_1) == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	F_pfree(m, v262)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L96
	}
L58:
	;
	F_pg_qsort(m, v262, v261, int32(8), int32(_a_F_g_int_picksplit_2))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v272 = int32(1)
	v274 = v272
	v278 = v272
	goto L62
L61:
	;
	v465 = v192
	v469 = v217
	v471 = v251
	v472 = v237
	goto L57
L62:
	;
	v297 = v262 + v274<<(uint(int32(3))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v297-int32(8)))) = uint16(v278)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v219+v274*int32(24))))
	v305 = F_inner_int_union(m, v237, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	F_pg_qsort(m, v262, v261, int32(8), int32(_a_F_g_int_picksplit_2))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L71
	}
L64:
	;
	F_rt__int_size(m, v305, v24+int32(28))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v305)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v313 = F_inner_int_union(m, v251, v304)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_rt__int_size(m, v313, v24+int32(24))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v313)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v323 = *(*float32)(unsafe.Add(mBase, uint32(v24)+28))
	v324 = *(*float32)(unsafe.Add(mBase, uint32(v24)+12))
	v326 = *(*float32)(unsafe.Add(mBase, uint32(v24)+24))
	v327 = *(*float32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v297-int32(4)))) = base.F32_abs(base.F32_sub(base.F32_sub(v323, v324), base.F32_sub(v326, v327)))
	v333 = v278 + int32(1)
	v334 = int32(_a_F_g_int_picksplit_1)
	v335 = v333 & v334
	if base.Ui32(v335) <= base.Ui32(v259&v334) {
		v274 = v335
		v278 = v333
		goto L62
	} else {
		goto L70
	}
L70:
	;
	goto L63
L71:
	;
	v343 = int32(1)
	if base.Ui32(v261) <= base.Ui32(v343) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v346 = v343
	goto L74
L73:
	;
	v346 = v261
	goto L74
L74:
	;
	v350 = int32(0)
	v352 = v192
	v356 = v217
	v358 = v251
	v359 = v237
	goto L75
L75:
	;
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262+v350<<(uint(int32(3))%32)))))
	if v230&int32(_a_F_g_int_picksplit_1) == v374 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v465 = v453
	v469 = v456
	v471 = v457
	v472 = v458
	goto L57
L77:
	;
	v461 = v350 + int32(1)
	if v461 != v346 {
		v350 = v461
		v352 = v453
		v356 = v456
		v358 = v457
		v359 = v458
		goto L75
	} else {
		goto L95
	}
L78:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v356))) = uint16(v230)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v377 + int32(1)
	v453 = v352
	v456 = v356 + int32(2)
	v457 = v358
	v458 = v359
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v244&int32(_a_F_g_int_picksplit_1) == v374 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v352))) = uint16(v244)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v387 + int32(1)
	v453 = v352 + int32(2)
	v456 = v356
	v457 = v358
	v458 = v359
	goto L77
L82:
	;
	goto L83
L83:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v219+v374*int32(24))))
	v397 = F_inner_int_union(m, v359, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v399 = F_inner_int_union(m, v358, v396)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_rt__int_size(m, v397, v24+int32(28))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_rt__int_size(m, v399, v24+int32(24))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v409 = *(*float32)(unsafe.Add(mBase, uint32(v24)+28))
	v410 = *(*float32)(unsafe.Add(mBase, uint32(v24)+12))
	v413 = *(*float32)(unsafe.Add(mBase, uint32(v24)+24))
	v414 = *(*float32)(unsafe.Add(mBase, uint32(v24)+8))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	v419 = v417 - v418
	if base.F64_lt(base.F64_promote_f32(base.F32_sub(v409, v410)), base.F64_add(base.F64_promote_f32(base.F32_sub(v413, v414)), base.F64_mul(base.F64_convert_i32_s(v419*v419*v419), float64(-0.01)))) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_pfree(m, v359)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_pfree(m, v358)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	F_pfree(m, v399)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v431 = *(*float32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*float32)(unsafe.Add(mBase, uint32(v24)+12)) = v431
	*(*uint16)(unsafe.Add(mBase, uint32(v356))) = uint16(v374)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v434 + int32(1)
	v453 = v352
	v456 = v356 + int32(2)
	v457 = v358
	v458 = v397
	goto L77
L93:
	;
	F_pfree(m, v397)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v444 = *(*float32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*float32)(unsafe.Add(mBase, uint32(v24)+8)) = v444
	*(*uint16)(unsafe.Add(mBase, uint32(v352))) = uint16(v374)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v447 + int32(1)
	v453 = v352 + int32(2)
	v456 = v356
	v457 = v399
	v458 = v359
	goto L77
L95:
	;
	goto L76
L96:
	;
	v486 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v469))) = uint16(v486)
	*(*uint16)(unsafe.Add(mBase, uint32(v465))) = uint16(v486)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = base.I64_extend_i32_u(v471)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = base.I64_extend_i32_u(v472)
	m.G0 = v24 + int32(32)
	return v26 & int64(4294967295)
}
