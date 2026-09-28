package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_row_identity_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v15 == l2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v33 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L12
	}
L5:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	v22 = v18 + int32(1)
	goto L7
L6:
	;
	v22 = int32(1)
	goto L7
L7:
	;
	v24 = F_pstrdup(m, l3)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v27 = F_makeTargetEntry(m, l1, base.I32_extend16_s(v22), v24, int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v30 = F_lappend(m, v29, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v30
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = int32(-4)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v122 = F_palloc0(m, int32(20))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L40
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(0)
	if v43 < v40 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v47 = v40
	goto L18
L17:
	;
	v47 = v43
	goto L18
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v50 = v43
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(v65 == int32(0))|base.B2i32(v65 != v68) != 0 {
		v86 = v65
		v87 = v68
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v93 = F_equal(m, v33, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L32
	}
L21:
	;
	if v86-v87 != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	v71 = l3
	v72 = v62
	goto L24
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v76
		v87 = v75
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v86 = v76
	v87 = v75
	goto L22
L26:
	;
	v79 = int32(1)
	if v76 == v75 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v90 = v50 + int32(1)
	if v47 != v90 {
		v50 = v90
		goto L19
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L20
L31:
	;
	goto L13
L32:
	;
	if v93 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v96 = F_bms_add_member(m, v95, l2)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v96
	goto L1
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
	F_errmsg_internal(m, int32(_a_F_add_row_identity_var_0), v12)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_add_row_identity_var_1), int32(925), int32(_a_F_add_row_identity_var_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(325)
	v126 = F_copyObjectImpl(m, v33)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v126
	v129 = F_exprType(m, v33)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v131 = F_exprTypmod(m, v33)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v133 = F_get_typavgwidth(m, v129, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v133
	v136 = F_pstrdup(m, l3)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v136
	v139 = F_bms_make_singleton(m, l2)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v139
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v143 = F_lappend(m, v142, v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v143
	if v143 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v148 = v146
	goto L50
L49:
	;
	v148 = int32(0)
	goto L50
L50:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+8)) = uint16(v148)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v150 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+4)))
	v155 = v151 + int32(1)
	goto L53
L52:
	;
	v155 = int32(1)
	goto L53
L53:
	;
	v157 = F_pstrdup(m, l3)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v160 = F_makeTargetEntry(m, v33, base.I32_extend16_s(v155), v157, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v163 = F_lappend(m, v162, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v163
	goto L1
}
func F_make_row_comparison_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
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
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v22 = v21
	goto L3
L2:
	;
	v22 = v6
	goto L3
L3:
	;
	if l3 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	m.G0 = v19 + int32(48)
	return v597
L5:
	;
	if v284 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L6:
	;
	v268 = int32(4)
	v271 = F_palloc_mul(m, v268, v22)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L30
	} else {
		goto L81
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L30
	} else {
		goto L76
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L30
	} else {
		goto L70
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L30
	} else {
		goto L65
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v25 = v23
	goto L12
L11:
	;
	v25 = int32(0)
	goto L12
L12:
	;
	if v25 == v22 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v22 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L30
	} else {
		goto L60
	}
L16:
	;
	v34 = v6
	v36 = v6
	goto L17
L17:
	;
	v45 = int32(0)
	if l2 == v45 {
		v55 = v45
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v91 = int32(4)
	v92 = v36 + v91
	v94 = F_palloc_mul(m, v91, v22)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L30
	} else {
		goto L36
	}
L19:
	;
	if l3 == int32(0) {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v49 <= v34 {
		v55 = int32(0)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v55 = v51 + v34<<(uint(int32(2))%32)
	goto L19
L22:
	;
	v58 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v55 == v58)|base.B2i32(v60 <= v34) == v58 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L18
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65+v34<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v78 = F_make_op(m, l0, l1, v72, v76, v77, l4)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v65 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v67 = int32(1)
	if v22 != v67 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v597 = v71
	goto L4
L30:
	;
	return int32(0)
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v82 != int32(16) {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v85 = F_expression_returns_set(m, v78)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v85 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v89 = F_lappend(m, v36, v78)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v34 = v34 + int32(1)
	v36 = v89
	goto L17
L36:
	;
	if v36 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v279 = v67
	v280 = int32(0)
	v283 = v92
	v284 = v6
	v285 = v94
	goto L5
L38:
	;
	goto L39
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v99 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v279 = int32(0)
	v280 = v36
	v283 = v92
	v284 = v6
	v285 = v94
	goto L5
L41:
	;
	goto L42
L42:
	;
	v106 = int32(0)
	v115 = v6
	goto L43
L43:
	;
	v121 = v106 << (uint(int32(2)) % 32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123+v121)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v127 = F_get_op_index_interpretation(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L30
	} else {
		goto L45
	}
L44:
	;
	v279 = int32(0)
	v280 = v36
	v283 = v92
	v284 = v184
	v285 = v94
	goto L5
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94+v121))) = v127
	if v127 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v106 != 0 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v169 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v133 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v135 <= v133 {
		v169 = v133
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v141 = v133
	v143 = v133
	goto L51
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v143<<(uint(int32(2))%32))))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v160 = F_bms_add_member(m, v141, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L30
	} else {
		goto L53
	}
L52:
	;
	v169 = v160
	goto L46
L53:
	;
	v163 = v143 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v163 < v164 {
		v141 = v160
		v143 = v163
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v182 = F_bms_int_members(m, v115, v169)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L30
	} else {
		goto L58
	}
L56:
	;
	v184 = v169
	goto L57
L57:
	;
	v186 = v106 + int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v186 < v187 {
		v106 = v186
		v115 = v184
		goto L43
	} else {
		goto L59
	}
L58:
	;
	v184 = v182
	goto L57
L59:
	;
	goto L44
L60:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L30
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_make_row_comparison_op_0), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_1), int32(2859), int32(_a_F_make_row_comparison_op_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L30
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L30
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_make_row_comparison_op_3), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L30
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_1), int32(2869), int32(_a_F_make_row_comparison_op_2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v234 = F_format_type_be(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L30
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v234
	F_errmsg(m, int32(_a_F_make_row_comparison_op_4), v19+int32(32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L30
	} else {
		goto L73
	}
L73:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L30
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_1), int32(2896), int32(_a_F_make_row_comparison_op_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L30
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L30
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_make_row_comparison_op_5), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L30
	} else {
		goto L78
	}
L78:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L30
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_1), int32(2901), int32(_a_F_make_row_comparison_op_2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L30
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v279 = int32(1)
	v280 = int32(0)
	v283 = v268
	v284 = v6
	v285 = v271
	goto L5
L82:
	;
	if v345 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L83:
	;
	v345 = base.I32_ctz(v331) | v332<<(uint(int32(5))%32)
	goto L82
L84:
	;
	v345 = int32(-2)
	goto L82
L85:
	;
	v296 = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v299 <= v296 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v302 = v284 + int32(8)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v309 = v306 & int32(-1)
	if v309 != 0 {
		v331 = v309
		v332 = v296
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v310 = int32(1)
	if v310 == v299 {
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v314 = v310
	goto L89
L89:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v302+v314<<(uint(int32(2))%32))))
	if v321 != 0 {
		v331 = v321
		v332 = v314
		goto L83
	} else {
		goto L91
	}
L90:
	;
	goto L84
L91:
	;
	v323 = v314 + int32(1)
	if v323 != v299 {
		v314 = v323
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L30
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	switch v345 - int32(3) {
	case 0:
		goto L102
	default:
		goto L104
	case 3:
		goto L103
	}
L96:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L30
	} else {
		goto L97
	}
L97:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v355+v356<<(uint(int32(2))%32)-int32(4))))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v363
	F_errmsg(m, int32(_a_F_make_row_comparison_op_6), v19)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L30
	} else {
		goto L98
	}
L98:
	;
	F_errhint(m, int32(_a_F_make_row_comparison_op_7), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L30
	} else {
		goto L99
	}
L99:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L30
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_1), int32(2962), int32(_a_F_make_row_comparison_op_2))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L30
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v590 = F_makeBoolExpr(m, int32(0), v280, l4)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L30
	} else {
		goto L141
	}
L103:
	;
	v587 = F_makeBoolExpr(m, int32(1), v280, l4)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L30
	} else {
		goto L140
	}
L104:
	;
	if v22 <= int32(0) {
		v511 = v6
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v513 = int32(0)
	if v279 != 0 {
		v558 = v513
		v560 = v513
		v561 = v513
		goto L130
	} else {
		goto L131
	}
L106:
	;
	v395 = int32(0)
	v398 = v6
	goto L107
L107:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v285+v395<<(uint(int32(2))%32))))
	if v403 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L30
	} else {
		goto L124
	}
L109:
	;
	goto L108
L110:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v406 <= int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v409 = int32(0)
	if v409 < v406 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v412 = v406
	goto L114
L113:
	;
	v412 = v409
	goto L114
L114:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v420 = int32(0)
	goto L115
L115:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v413+v420<<(uint(int32(2))%32))))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v345 != v435 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if v440 == int32(0) {
		goto L109
	} else {
		goto L121
	}
L117:
	;
	v438 = v420 + int32(1)
	if v412 != v438 {
		v420 = v438
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	goto L109
L121:
	;
	v443 = F_lappend_oid(m, v398, v440)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L30
	} else {
		goto L122
	}
L122:
	;
	v446 = v395 + int32(1)
	if v446 != v22 {
		v395 = v446
		v398 = v443
		goto L107
	} else {
		goto L123
	}
L123:
	;
	v511 = v443
	goto L105
L124:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L30
	} else {
		goto L125
	}
L125:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v471+v472<<(uint(int32(2))%32)-int32(4))))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v479
	F_errmsg(m, int32(_a_F_make_row_comparison_op_6), v19+int32(16))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L30
	} else {
		goto L126
	}
L126:
	;
	v488 = F_errdetail(m, int32(_a_F_make_row_comparison_op_8), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L30
	} else {
		goto L127
	}
L127:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L30
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_1), int32(3003), int32(_a_F_make_row_comparison_op_2))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L30
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v575 = F_palloc0(m, int32(28))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L30
	} else {
		goto L139
	}
L131:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v516 <= int32(0) {
		v558 = v513
		v560 = v513
		v561 = v513
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v520 = v513
	v522 = v513
	v523 = v513
	v525 = int32(0)
	goto L133
L133:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536+v525<<(uint(int32(2))%32))))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	v542 = F_lappend_oid(m, v523, v541)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L30
	} else {
		goto L135
	}
L134:
	;
	v558 = v547
	v560 = v552
	v561 = v542
	goto L130
L135:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540)+28))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+12))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v547 = F_lappend(m, v520, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L30
	} else {
		goto L136
	}
L136:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v540)+28))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+12))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	v552 = F_lappend(m, v522, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L30
	} else {
		goto L137
	}
L137:
	;
	v555 = v525 + int32(1)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v555 < v556 {
		v520 = v547
		v522 = v552
		v523 = v542
		v525 = v555
		goto L133
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575)+24)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v575)+20)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v575)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v575)+12)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v575)+8)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v575)+4)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = int32(37)
	v597 = v575
	goto L4
L140:
	;
	v597 = v587
	goto L4
L141:
	;
	v597 = v590
	goto L4
}
func F_row_to_json(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_initStringInfo(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_composite_to_json(m, v8, v6, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v18 = F_cstring_to_text_with_len(m, v16, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v18)
			}
		}
	}
}
func F_transformRowExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = F_palloc0(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(36)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v19 = F_transformExpressionList(m, l0, v17, v18, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v19
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if int32(1665) <= v22 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(8589936841)
	v30 = int32(1)
	v32 = v19
	goto L9
L8:
	;
	goto L7
L9:
	;
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v57
	m.G0 = v8 + int32(48)
	return v11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v37 = v35
	goto L13
L12:
	;
	v37 = int32(0)
	goto L13
L13:
	;
	if v30 <= v37 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30
	v41 = v8 + int32(32)
	v44 = F_pg_snprintf(m, v41, int32(16), int32(_a_F_transformRowExpr_0), v8)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L10
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v47 = F_pstrdup(m, v41)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v49 = F_makeString(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v51 = F_lappend(m, v46, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v51
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v30 = v30 + int32(1)
	v32 = v56
	goto L9
L21:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(1664)
	F_errmsg(m, int32(_a_F_transformRowExpr_1), v8+int32(16))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_transformRowExpr_2), int32(2207), int32(_a_F_transformRowExpr_3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
