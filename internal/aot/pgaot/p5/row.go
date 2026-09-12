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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v18 != 0 {
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
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = v19 + int32(1)
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
	v121 = F_palloc0(m, int32(20))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L41
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
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v66 == int32(0) {
		v85 = v65
		v86 = v66
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v92 = F_equal(m, v33, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L33
	}
L21:
	;
	if v86-v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	goto L21
L23:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v70 = l3
	v71 = v62
	goto L25
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v85 = v74
	v86 = v75
	goto L22
L27:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v89 = v50 + int32(1)
	if v47 != v89 {
		v50 = v89
		goto L19
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L20
L32:
	;
	goto L13
L33:
	;
	if v92 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v95 = F_bms_add_member(m, v94, l2)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v95
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
	F_errmsg_internal(m, int32(742695), v12)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(516161), int32(874), int32(239148))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = int32(323)
	v125 = F_copyObjectImpl(m, v33)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v125
	v128 = F_exprType(m, v33)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v130 = F_exprTypmod(m, v33)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v132 = F_get_typavgwidth(m, v128, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v132
	v135 = F_pstrdup(m, l3)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v135
	v138 = F_bms_make_singleton(m, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v142 = F_lappend(m, v141, v121)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v142
	if v142 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v147 = v145
	goto L51
L50:
	;
	v147 = int32(0)
	goto L51
L51:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+8)) = uint16(v147)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v150 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+4)))
	v154 = v151 + int32(1)
	goto L54
L53:
	;
	v154 = int32(1)
	goto L54
L54:
	;
	v156 = F_pstrdup(m, l3)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v159 = F_makeTargetEntry(m, v33, base.I32_extend16_s(v154), v156, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v162 = F_lappend(m, v161, v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v162
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
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v69 int32
	_ = v69
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
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
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
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
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v17 + int32(48)
	return v560
L5:
	;
	if v257 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v22 = v21
	goto L8
L7:
	;
	v22 = v6
	goto L8
L8:
	;
	if v22 == v20 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v20 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L32
	} else {
		goto L76
	}
L12:
	;
	v29 = v6
	v31 = v6
	goto L17
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L32
	} else {
		goto L71
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L32
	} else {
		goto L65
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L32
	} else {
		goto L60
	}
L17:
	;
	v38 = int32(0)
	if l2 == v38 {
		v48 = v38
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v84 = F_palloc(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L32
	} else {
		goto L38
	}
L19:
	;
	if l3 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v42 <= v29 {
		v48 = int32(0)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v48 = v44 + v29<<(uint(int32(2))%32)
	goto L19
L22:
	;
	goto L18
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v69 = F_make_op(m, l0, l1, v66, v67, v68, l4)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	if v20 != int32(1) {
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v60 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v52 <= v29 {
		v60 = v31
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v48 == int32(0) {
		v60 = v31
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v59 = v56 + v29<<(uint(int32(2))%32)
	if v59 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v60 = v31
	goto L24
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v560 = v65
	goto L4
L32:
	;
	return int32(0)
L33:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v73 != int32(16) {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	v76 = F_expression_returns_set(m, v69)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if v76 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	v80 = F_lappend(m, v31, v69)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v29 = v29 + int32(1)
	v31 = v80
	goto L17
L38:
	;
	if v60 == int32(0) {
		v257 = v6
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v89 <= v88 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v257 = v6
	goto L5
L41:
	;
	goto L42
L42:
	;
	v100 = v88
	v101 = v6
	goto L43
L43:
	;
	v107 = v100 << (uint(int32(2)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109+v107)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v113 = F_get_op_index_interpretation(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L32
	} else {
		goto L45
	}
L44:
	;
	v257 = v166
	goto L5
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84+v107))) = v113
	if v113 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v100 != 0 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v153 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v121 <= v119 {
		v153 = v119
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v127 = v119
	v129 = v119
	goto L51
L51:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v129<<(uint(int32(2))%32))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v144 = F_bms_add_member(m, v127, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L32
	} else {
		goto L53
	}
L52:
	;
	v153 = v144
	goto L46
L53:
	;
	v147 = v129 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v147 < v148 {
		v127 = v144
		v129 = v147
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v164 = F_bms_int_members(m, v101, v153)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L32
	} else {
		goto L58
	}
L56:
	;
	v166 = v153
	goto L57
L57:
	;
	v168 = v100 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v168 < v169 {
		v100 = v168
		v101 = v166
		goto L43
	} else {
		goto L59
	}
L58:
	;
	v166 = v164
	goto L57
L59:
	;
	goto L44
L60:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(113399), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L32
	} else {
		goto L62
	}
L62:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(514888), int32(2900), int32(245268))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L32
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L32
	} else {
		goto L66
	}
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v197 = F_format_type_be(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L32
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v197
	F_errmsg(m, int32(197713), v17+int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L32
	} else {
		goto L68
	}
L68:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L32
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(514888), int32(2895), int32(245268))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L32
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L32
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(333961), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L32
	} else {
		goto L73
	}
L73:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L32
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(514888), int32(2868), int32(245268))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L32
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L32
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(152754), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L32
	} else {
		goto L78
	}
L78:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L32
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(514888), int32(2858), int32(245268))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L32
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
	if v318 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L82:
	;
	v318 = base.I32_ctz(v304) | v305<<(uint(int32(5))%32)
	goto L81
L83:
	;
	v318 = int32(-2)
	goto L81
L84:
	;
	v271 = base.I32_div_s(int32(0), int32(32))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v272 <= v271 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v275 = v257 + int32(8)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v271<<(uint(int32(2))%32))))
	v282 = v279 & int32(-1)
	if v282 != 0 {
		v304 = v282
		v305 = v271
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v284 = v271 + int32(1)
	if v284 == v272 {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v287 = v284
	goto L88
L88:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v275+v287<<(uint(int32(2))%32))))
	if v294 != 0 {
		v304 = v294
		v305 = v287
		goto L82
	} else {
		goto L90
	}
L89:
	;
	goto L83
L90:
	;
	v296 = v287 + int32(1)
	if v296 != v272 {
		v287 = v296
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L32
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	switch v318 - int32(3) {
	case 0:
		goto L103
	default:
		goto L104
	case 3:
		goto L101
	}
L95:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L32
	} else {
		goto L96
	}
L96:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v328+v329<<(uint(int32(2))%32)-int32(4))))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v336
	F_errmsg(m, int32(190098), v17)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L32
	} else {
		goto L97
	}
L97:
	;
	F_errhint(m, int32(623880), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L32
	} else {
		goto L98
	}
L98:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L32
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(514888), int32(2961), int32(245268))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L32
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v553 = F_makeBoolExpr(m, int32(1), v60, l4)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L32
	} else {
		goto L142
	}
L102:
	;
	v481 = int32(0)
	if v60 == v481 {
		v526 = v481
		v529 = v481
		v534 = v481
		goto L130
	} else {
		goto L131
	}
L103:
	;
	v465 = F_makeBoolExpr(m, int32(0), v60, l4)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L32
	} else {
		goto L129
	}
L104:
	;
	if v20 <= int32(0) {
		v479 = v6
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v366 = int32(0)
	v369 = v6
	goto L106
L106:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v84+v366<<(uint(int32(2))%32))))
	if v374 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L32
	} else {
		goto L123
	}
L108:
	;
	goto L107
L109:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v377 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v380 = int32(0)
	if v380 < v377 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v383 = v377
	goto L113
L112:
	;
	v383 = v380
	goto L113
L113:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v391 = int32(0)
	goto L114
L114:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v384+v391<<(uint(int32(2))%32))))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v318 != v404 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if v409 == int32(0) {
		goto L108
	} else {
		goto L120
	}
L116:
	;
	v407 = v391 + int32(1)
	if v383 != v407 {
		v391 = v407
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	goto L108
L120:
	;
	v412 = F_lappend_oid(m, v369, v409)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L32
	} else {
		goto L121
	}
L121:
	;
	v415 = v366 + int32(1)
	if v415 != v20 {
		v366 = v415
		v369 = v412
		goto L106
	} else {
		goto L122
	}
L122:
	;
	v479 = v412
	goto L102
L123:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L32
	} else {
		goto L124
	}
L124:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v438+v439<<(uint(int32(2))%32)-int32(4))))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v446
	F_errmsg(m, int32(190098), v17+int32(16))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L32
	} else {
		goto L125
	}
L125:
	;
	F_errdetail(m, int32(620387), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L32
	} else {
		goto L126
	}
L126:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L32
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(514888), int32(3002), int32(245268))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L32
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v560 = v465
	goto L4
L130:
	;
	v541 = F_palloc0(m, int32(28))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L32
	} else {
		goto L141
	}
L131:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v486 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v526 = v481
	v529 = v481
	v534 = v481
	goto L130
L133:
	;
	goto L134
L134:
	;
	v490 = v481
	v493 = v481
	v495 = int32(0)
	v498 = v481
	goto L135
L135:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v504+v495<<(uint(int32(2))%32))))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	v510 = F_lappend_oid(m, v493, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L32
	} else {
		goto L137
	}
L136:
	;
	v526 = v515
	v529 = v510
	v534 = v520
	goto L130
L137:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508)+28))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+12))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v515 = F_lappend(m, v490, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L32
	} else {
		goto L138
	}
L138:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v508)+28))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+12))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v520 = F_lappend(m, v498, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L32
	} else {
		goto L139
	}
L139:
	;
	v523 = v495 + int32(1)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v523 < v524 {
		v490 = v515
		v493 = v510
		v495 = v523
		v498 = v520
		goto L135
	} else {
		goto L140
	}
L140:
	;
	goto L136
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v541)+24)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v541)+20)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v541)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v541)+12)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v541)+8)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v541)+4)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = int32(37)
	v560 = v541
	goto L4
L142:
	;
	v560 = v553
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = F_palloc0(m, int32(24))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(36)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v20 = F_transformExpressionList(m, l0, v18, v19, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v20
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(1665) <= v23 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(8589936841)
	v31 = int32(1)
	v33 = v20
	goto L9
L8:
	;
	goto L7
L9:
	;
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v60
	m.G0 = v9 + int32(48)
	return v12
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = v37
	goto L13
L12:
	;
	v38 = int32(0)
	goto L13
L13:
	;
	if v31 <= v38 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
	v45 = F_pg_snprintf(m, v9+int32(32), int32(16), int32(483836), v9)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v50 = F_pstrdup(m, v9+int32(32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v52 = F_makeString(m, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v54 = F_lappend(m, v47, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v54
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v31 = v31 + int32(1)
	v33 = v59
	goto L9
L21:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(1664)
	F_errmsg(m, int32(176527), v9+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(514888), int32(2206), int32(216149))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
