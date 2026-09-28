package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_setop_child_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 float64
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
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
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v377 int32
	_ = v377
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 float64
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 float64
	_ = v530
	var v531 int32
	_ = v531
	var v533 float64
	_ = v533
	var v534 int32
	_ = v534
	var v552 float64
	_ = v552
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l4 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_set_subquery_size_estimates(m, l0, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L32
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v26 = v24
	goto L7
L6:
	;
	v26 = int32(0)
	goto L7
L7:
	;
	if l3 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L29
	}
L9:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v109 != 0 {
		goto L25
	} else {
		goto L26
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v29 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v38 = v26
	v41 = v7
	goto L12
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v41<<(uint(int32(2))%32))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+26)))
	if v54 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	if v38 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v83 = v38
	goto L16
L16:
	;
	v87 = v41 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v87 < v88 {
		v38 = v83
		v41 = v87
		goto L12
	} else {
		goto L24
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	v68 = F_exprType(m, v62)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	F_add_child_eq_member(m, l0, v60, int32(-1), v62, v63, v67, v66, v68, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v74 = v38 + int32(4)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.Ui32(v74) < base.Ui32(v76+v77<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v82 = v74
	goto L23
L22:
	;
	v82 = int32(0)
	goto L23
L23:
	;
	v83 = v82
	goto L16
L24:
	;
	goto L13
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v114 = v110 - int32(1)
	goto L27
L26:
	;
	v114 = int32(-1)
	goto L27
L27:
	;
	v115 = F_bms_add_range(m, v107, int32(0), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+144)) = v115
	goto L4
L29:
	;
	F_errmsg_internal(m, int32(_a_F_build_setop_child_paths_0), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_build_setop_child_paths_1), int32(3069), int32(_a_F_build_setop_child_paths_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v153 = F_fetch_upper_rel(m, v150, int32(7), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+44))
	if v157 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_build_setop_child_paths[0]))
	if v431 != 0 {
		goto L127
	} else {
		goto L128
	}
L35:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v401 != 0 {
		goto L34
	} else {
		goto L123
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if int32(0) < v158 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v155&int32(1) == int32(0) {
		goto L34
	} else {
		goto L122
	}
L39:
	;
	v176 = v7
	goto L42
L40:
	;
	goto L41
L41:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v377 == int32(0) {
		goto L34
	} else {
		goto L121
	}
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v176<<(uint(int32(2))%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v153)+60))
	v184 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v153)+44))
	if v186 == v184 {
		v207 = v184
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L41
L44:
	;
	v357 = v176 + int32(1)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v357 < v358 {
		v176 = v357
		goto L42
	} else {
		goto L120
	}
L45:
	;
	if v207 != 0 {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	goto L45
L47:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v190 = v189
	goto L48
L48:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if base.Ui32(int32(2)) <= base.Ui32(v194-int32(303)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v207 = int32(1)
	goto L46
L50:
	;
	if v194 != int32(293) {
		v207 = v184
		goto L46
	} else {
		goto L53
	}
L51:
	;
	v190 = v193 + int32(72)
	goto L48
L52:
	;
	goto L49
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v193)+72))
	if v201 != 0 {
		v207 = v184
		goto L46
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	F_mark_dummy_rel(m, l1)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L18
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v210 = base.B2i32(v182 != v183)
	if v210 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L44
L59:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v182)+64))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v215 = F_make_tlist_from_pathtarget(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L18
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if l4 == int32(0) {
		goto L44
	} else {
		goto L66
	}
L62:
	;
	v217 = F_convert_subquery_pathkeys(m, l0, l1, v213, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	v220 = F_create_subqueryscan_path(m, l0, l1, v182, l2, v217, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	F_add_path(m, l1, v220)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v182)+64))
	v228 = v20 + int32(12)
	if v23 == v226 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	if v332 == v183 {
		goto L44
	} else {
		goto L115
	}
L68:
	;
	if v306 != 0 {
		v332 = v182
		goto L67
	} else {
		goto L100
	}
L69:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v294
	v306 = int32(1)
	goto L68
L70:
	;
	if v23 != 0 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v23 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(0)
	v306 = int32(1)
	goto L68
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(0)
	v306 = int32(1)
	goto L68
L75:
	;
	goto L76
L76:
	;
	if v226 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v246 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v246
	v306 = v246
	goto L68
L78:
	;
	goto L79
L79:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v250 = int32(0)
	if v250 < v249 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v253 = v249
	goto L82
L81:
	;
	v253 = v250
	goto L82
L82:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v259 = int32(0)
	goto L83
L83:
	;
	if v259 < v254 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v270 = v266 + v259<<(uint(int32(2))%32)
	goto L87
L86:
	;
	v270 = int32(0)
	goto L87
L87:
	;
	if v259 == v253 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v253
	v306 = base.B2i32(v270 == int32(0))
	goto L68
L89:
	;
	goto L90
L90:
	;
	v276 = base.B2i32(v270 == int32(0))
	if v270 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v259
	v306 = v276
	goto L68
L92:
	;
	goto L93
L93:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	if v280 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v259
	v306 = v276
	goto L68
L95:
	;
	goto L96
L96:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v280+v259<<(uint(int32(2))%32))))
	if v284 != v288 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v259
	v306 = int32(0)
	goto L68
L98:
	;
	v259 = v259 + int32(1)
	goto L83
L100:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_setop_child_paths[1])))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v310 = *(*float64)(unsafe.Add(mBase, uint32(v309)+320))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v210 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v321&int32(1) != 0 {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v321 = v308
	goto L101
L103:
	;
	goto L104
L104:
	;
	if v311 == int32(0) {
		goto L44
	} else {
		goto L105
	}
L105:
	;
	v316 = int32(1)
	if v308&v316 == int32(0) {
		goto L44
	} else {
		goto L106
	}
L106:
	;
	v321 = v316
	goto L101
L107:
	;
	v325 = v311
	goto L109
L108:
	;
	v325 = int32(0)
	goto L109
L109:
	;
	if v325 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v328 = F_create_sort_path(m, v153, v182, v23, v310)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L18
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v330 = F_create_incremental_sort_path(m, v309, v153, v182, v23, v311, v310)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L18
	} else {
		goto L114
	}
L113:
	;
	v332 = v328
	goto L67
L114:
	;
	v332 = v330
	goto L67
L115:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v332)+64))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	v341 = F_make_tlist_from_pathtarget(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L18
	} else {
		goto L116
	}
L116:
	;
	v343 = F_convert_subquery_pathkeys(m, l0, l1, v339, v341)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L18
	} else {
		goto L117
	}
L117:
	;
	v346 = F_create_subqueryscan_path(m, l0, l1, v332, l2, v343, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L18
	} else {
		goto L118
	}
L118:
	;
	F_add_path(m, l1, v346)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L18
	} else {
		goto L119
	}
L119:
	;
	goto L44
L120:
	;
	goto L43
L121:
	;
	goto L35
L122:
	;
	goto L35
L123:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v153)+52))
	if v402 == int32(0) {
		goto L34
	} else {
		goto L124
	}
L124:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v402)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v407 = int32(0)
	v409 = F_create_subqueryscan_path(m, l0, l1, v406, l2, v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	F_add_partial_path(m, l1, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L18
	} else {
		goto L126
	}
L126:
	;
	goto L34
L127:
	;
	v432 = int32(0)
	m.T0[v431].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v432, v432, l1, v432)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L18
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L18
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	if l5 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+100))
	if v441 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	goto L134
L134:
	;
	m.G0 = v20 + int32(16)
	return
L135:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v552
	goto L134
L136:
	;
	v450 = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v440)+76))
	if v453 == v450 {
		v528 = v450
		goto L143
	} else {
		goto L144
	}
L137:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v449 = *(*float64)(unsafe.Add(mBase, uint32(v448)+32))
	v552 = v449
	goto L135
L138:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v440)+108))
	if v442 != 0 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v440)+120))
	if v443 != 0 {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+334)))
	if v444 != 0 {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+36)))
	if v445 != int32(1) {
		goto L136
	} else {
		goto L142
	}
L142:
	;
	goto L137
L143:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v530 = *(*float64)(unsafe.Add(mBase, uint32(v529)+32))
	v531 = int32(0)
	v533 = F_estimate_num_groups(m, v439, v528, v530, v531, v531)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L18
	} else {
		goto L155
	}
L144:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if int32(0) < v456 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v459 = v450
	v461 = v450
	goto L148
L146:
	;
	v494 = v450
	goto L147
L147:
	;
	v528 = v494
	goto L143
L148:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v476+v461<<(uint(int32(2))%32))))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+26)))
	if v481&int32(1) == int32(0) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v494 = v489
	goto L147
L150:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	v487 = F_lappend(m, v459, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L18
	} else {
		goto L153
	}
L151:
	;
	v489 = v459
	goto L152
L152:
	;
	v491 = v461 + int32(1)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v491 < v492 {
		v459 = v489
		v461 = v491
		goto L148
	} else {
		goto L154
	}
L153:
	;
	v489 = v487
	goto L152
L154:
	;
	goto L149
L155:
	;
	v552 = v533
	goto L135
}
