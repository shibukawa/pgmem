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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v30
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = int32(-4)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
	F_errfinish(m, int32(_a_F_add_row_identity_var_1), int32(874), int32(_a_F_add_row_identity_var_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(323)
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
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v143
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
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
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
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v163
	goto L1
}
func F_make_row_comparison_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v20 = v19
	goto L3
L2:
	;
	v20 = v6
	goto L3
L3:
	;
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	m.G0 = v17 + int32(48)
	return v564
L5:
	;
	v557 = F_makeBoolExpr(m, int32(0), v61, l4)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L33
	} else {
		goto L135
	}
L6:
	;
	v554 = F_makeBoolExpr(m, int32(1), v61, l4)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L33
	} else {
		goto L134
	}
L7:
	;
	v482 = int32(0)
	if v61 == v482 {
		v527 = v482
		v529 = v482
		v530 = v482
		goto L124
	} else {
		goto L125
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L33
	} else {
		goto L118
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L33
	} else {
		goto L113
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L33
	} else {
		goto L107
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L33
	} else {
		goto L102
	}
L12:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v23 = v21
	goto L14
L13:
	;
	v23 = int32(0)
	goto L14
L14:
	;
	if v23 == v20 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v20 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L33
	} else {
		goto L97
	}
L18:
	;
	v32 = v6
	v34 = v6
	goto L19
L19:
	;
	v41 = int32(0)
	if l2 == v41 {
		v51 = v41
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v88 = F_palloc(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L33
	} else {
		goto L39
	}
L21:
	;
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v45 <= v32 {
		v51 = int32(0)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v51 = v47 + v32<<(uint(int32(2))%32)
	goto L21
L24:
	;
	goto L20
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60+v32<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v73 = F_make_op(m, l0, l1, v67, v71, v72, l4)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v20 != int32(1) {
		goto L24
	} else {
		goto L32
	}
L27:
	;
	v61 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v51 == int32(0))|base.B2i32(v57 <= v32) != 0 {
		v61 = v34
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v60 != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v61 = v34
	goto L26
L32:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v564 = v66
	goto L4
L33:
	;
	return int32(0)
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	if v77 != int32(16) {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v80 = F_expression_returns_set(m, v73)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	if v80 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v84 = F_lappend(m, v34, v73)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v32 = v32 + int32(1)
	v34 = v84
	goto L19
L39:
	;
	if v61 == int32(0) {
		v185 = v6
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v185 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v92 <= int32(0) {
		v185 = v6
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v98 = int32(0)
	v106 = v6
	goto L43
L43:
	;
	v111 = v98 << (uint(int32(2)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113+v111)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v117 = F_get_op_index_interpretation(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L33
	} else {
		goto L45
	}
L44:
	;
	v185 = v170
	goto L40
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88+v111))) = v117
	if v117 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v98 != 0 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v157 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v123 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v125 <= v123 {
		v157 = v123
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v131 = v123
	v133 = v123
	goto L51
L51:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v133<<(uint(int32(2))%32))))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v148 = F_bms_add_member(m, v131, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L33
	} else {
		goto L53
	}
L52:
	;
	v157 = v148
	goto L46
L53:
	;
	v151 = v133 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v151 < v152 {
		v131 = v148
		v133 = v151
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v168 = F_bms_int_members(m, v106, v157)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L33
	} else {
		goto L58
	}
L56:
	;
	v170 = v157
	goto L57
L57:
	;
	v172 = v98 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v172 < v173 {
		v98 = v172
		v106 = v170
		goto L43
	} else {
		goto L59
	}
L58:
	;
	v170 = v168
	goto L57
L59:
	;
	goto L44
L60:
	;
	if v245 < int32(0) {
		goto L8
	} else {
		goto L71
	}
L61:
	;
	v245 = base.I32_ctz(v231) | v232<<(uint(int32(5))%32)
	goto L60
L62:
	;
	v245 = int32(-2)
	goto L60
L63:
	;
	v198 = base.I32_div_s(int32(0), int32(32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v199 <= v198 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v202 = v185 + int32(8)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v198<<(uint(int32(2))%32))))
	v209 = v206 & int32(-1)
	if v209 != 0 {
		v231 = v209
		v232 = v198
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v211 = v198 + int32(1)
	if v211 == v199 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v214 = v211
	goto L67
L67:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202+v214<<(uint(int32(2))%32))))
	if v221 != 0 {
		v231 = v221
		v232 = v214
		goto L61
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v223 = v214 + int32(1)
	if v223 != v199 {
		v214 = v223
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	switch v245 - int32(3) {
	case 0:
		goto L5
	default:
		goto L72
	case 3:
		goto L6
	}
L72:
	;
	if v20 <= int32(0) {
		v480 = v6
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v263 = int32(0)
	v265 = v6
	goto L74
L74:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v88+v263<<(uint(int32(2))%32))))
	if v270 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L33
	} else {
		goto L91
	}
L76:
	;
	goto L75
L77:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v273 <= int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v276 = int32(0)
	if v276 < v273 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v279 = v273
	goto L81
L80:
	;
	v279 = v276
	goto L81
L81:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v287 = int32(0)
	goto L82
L82:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v280+v287<<(uint(int32(2))%32))))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if v245 != v300 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v305 == int32(0) {
		goto L76
	} else {
		goto L88
	}
L84:
	;
	v303 = v287 + int32(1)
	if v279 != v303 {
		v287 = v303
		goto L82
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	goto L76
L88:
	;
	v308 = F_lappend_oid(m, v265, v305)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L33
	} else {
		goto L89
	}
L89:
	;
	v311 = v263 + int32(1)
	if v311 != v20 {
		v263 = v311
		v265 = v308
		goto L74
	} else {
		goto L90
	}
L90:
	;
	v480 = v308
	goto L7
L91:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L33
	} else {
		goto L92
	}
L92:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v334+v335<<(uint(int32(2))%32)-int32(4))))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v342
	F_errmsg(m, int32(_a_F_make_row_comparison_op_0), v17+int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L33
	} else {
		goto L93
	}
L93:
	;
	F_errdetail(m, int32(_a_F_make_row_comparison_op_1), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L33
	} else {
		goto L94
	}
L94:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L33
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_2), int32(3002), int32(_a_F_make_row_comparison_op_3))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L33
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L33
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_make_row_comparison_op_4), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L33
	} else {
		goto L99
	}
L99:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L33
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_2), int32(2858), int32(_a_F_make_row_comparison_op_3))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L33
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L33
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_make_row_comparison_op_5), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L33
	} else {
		goto L104
	}
L104:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L33
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_2), int32(2868), int32(_a_F_make_row_comparison_op_3))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L33
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L33
	} else {
		goto L108
	}
L108:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v404 = F_format_type_be(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L33
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v404
	F_errmsg(m, int32(_a_F_make_row_comparison_op_6), v17+int32(32))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L33
	} else {
		goto L110
	}
L110:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L33
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_2), int32(2895), int32(_a_F_make_row_comparison_op_3))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L33
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L33
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_make_row_comparison_op_7), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L33
	} else {
		goto L115
	}
L115:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L33
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_2), int32(2900), int32(_a_F_make_row_comparison_op_3))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L33
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L33
	} else {
		goto L119
	}
L119:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v444+v445<<(uint(int32(2))%32)-int32(4))))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v452
	F_errmsg(m, int32(_a_F_make_row_comparison_op_0), v17)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L33
	} else {
		goto L120
	}
L120:
	;
	F_errhint(m, int32(_a_F_make_row_comparison_op_8), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L33
	} else {
		goto L121
	}
L121:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L33
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_make_row_comparison_op_2), int32(2961), int32(_a_F_make_row_comparison_op_3))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L33
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	v542 = F_palloc0(m, int32(28))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L33
	} else {
		goto L133
	}
L125:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v487 <= int32(0) {
		v527 = v482
		v529 = v482
		v530 = v482
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v491 = v482
	v493 = v482
	v494 = v482
	v496 = int32(0)
	goto L127
L127:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v505+v496<<(uint(int32(2))%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v511 = F_lappend_oid(m, v494, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L33
	} else {
		goto L129
	}
L128:
	;
	v527 = v516
	v529 = v521
	v530 = v511
	goto L124
L129:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509)+28))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v516 = F_lappend(m, v491, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L33
	} else {
		goto L130
	}
L130:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v509)+28))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	v521 = F_lappend(m, v493, v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L33
	} else {
		goto L131
	}
L131:
	;
	v524 = v496 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v524 < v525 {
		v491 = v516
		v493 = v521
		v494 = v511
		v496 = v524
		goto L127
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+24)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v542)+20)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v542)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+12)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v542)+8)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v542)+4)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = int32(37)
	v564 = v542
	goto L4
L134:
	;
	v564 = v554
	goto L4
L135:
	;
	v564 = v557
	goto L4
}
func F_row_to_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_makeStringInfo(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_composite_to_json(m, v2, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			v12 = F_cstring_to_text_with_len(m, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
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
	F_errfinish(m, int32(_a_F_transformRowExpr_2), int32(2206), int32(_a_F_transformRowExpr_3))
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
