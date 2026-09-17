package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_union(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v462 int32
	_ = v462
	var v463 int64
	_ = v463
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v528 = F_ArrayGetNItemsSafe(m, v525, v516+int32(16))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L8
	} else {
		goto L122
	}
L2:
	;
	if v54 <= v206 {
		v291 = v205
		goto L71
	} else {
		goto L72
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L67
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L63
	}
L5:
	;
	v14 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return int32(0)
L9:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v19 = F_array_contains_nulls(m, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = l0 + int32(16)
	v24 = F_ArrayGetNItemsSafe(m, v21, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L17
	}
L14:
	;
	if v19 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = F_ArrayGetNItemsSafe(m, v35, v23)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L22
	}
L17:
	;
	if v24 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = F_ArrayGetNItemsSafe(m, v26, l1+int32(16))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if v29 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v32 = F_construct_empty_array(m, int32(23))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	return v32
L22:
	;
	if v36 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v40 = F_copy_intArrayType(m, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	v42 = int32(0)
	goto L25
L25:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v45 = l1 + int32(16)
	v46 = F_ArrayGetNItemsSafe(m, v43, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	v42 = v40
	goto L25
L27:
	;
	if v46 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v50 = F_copy_intArrayType(m, l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	v52 = v42
	goto L30
L30:
	;
	if v52 != 0 {
		v516 = v52
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v52 = v50
	goto L30
L32:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = F_ArrayGetNItemsSafe(m, v53, v23)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v57 = F_ArrayGetNItemsSafe(m, v56, v45)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v59 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = (v62<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L37
L36:
	;
	v69 = v59
	goto L37
L37:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v70 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v80 = (v73<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L40
L39:
	;
	v80 = v70
	goto L40
L40:
	;
	v81 = v54 + v57
	if int32(0) < v81 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v121 = l0 + v69
	v122 = l1 + v80
	v123 = v118 + v119
	v124 = int32(0)
	if base.B2i32(v124 < v54)&base.B2i32(v124 < v57) == v124 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v118 = v108
	v119 = (v111<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v120 = v110
	goto L41
L43:
	;
	v87 = v81<<(uint(int32(2))%32) + int32(24)
	v88 = F_palloc0(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v103 = F_construct_empty_array(m, int32(23))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = int32(23)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v87 << (uint(int32(2)) % 32)
	v108 = v88
	v110 = v88 + int32(8)
	goto L42
L47:
	;
	v106 = v103 + int32(8)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	if v107 != 0 {
		v118 = v103
		v119 = v107
		v120 = v106
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v108 = v103
	v110 = v106
	goto L42
L49:
	;
	v205 = v123
	v206 = int32(0)
	v209 = v124
	goto L2
L50:
	;
	goto L51
L51:
	;
	v134 = v123
	v135 = int32(0)
	v138 = v124
	goto L52
L52:
	;
	v146 = int32(2)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v121+v135<<(uint(v146)%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v122+v138<<(uint(v146)%32))))
	if v149 == v153 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v205 = v170
	v206 = v167
	v209 = v168
	goto L2
L54:
	;
	v170 = v134 + int32(4)
	if v54 <= v167 {
		v205 = v170
		v206 = v167
		v209 = v168
		goto L2
	} else {
		goto L61
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v149
	v156 = int32(1)
	v167 = v135 + v156
	v168 = v138 + v156
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v149 < v153 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v149
	v167 = v135 + int32(1)
	v168 = v138
	goto L54
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v153
	v167 = v135
	v168 = v138 + int32(1)
	goto L54
L61:
	;
	if v168 < v57 {
		v134 = v170
		v135 = v167
		v138 = v168
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L53
L63:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_inner_int_union_0), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_inner_int_union_1), int32(83), int32(_a_F_inner_int_union_2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_inner_int_union_0), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_inner_int_union_1), int32(84), int32(_a_F_inner_int_union_2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
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
	if v57 <= v209 {
		v377 = v291
		goto L84
	} else {
		goto L85
	}
L72:
	;
	v220 = (v54 - v206) & int32(3)
	if v220 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v206-v54) {
		v291 = v248
		goto L71
	} else {
		goto L80
	}
L74:
	;
	v248 = v205
	v250 = v206
	goto L73
L75:
	;
	goto L76
L76:
	;
	v224 = v205
	v226 = v206
	v230 = int32(0)
	goto L77
L77:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v121+v226<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v239
	v242 = v224 + int32(4)
	v243 = int32(1)
	v244 = v226 + v243
	v246 = v230 + v243
	if v246 != v220 {
		v224 = v242
		v226 = v244
		v230 = v246
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v248 = v242
	v250 = v244
	goto L73
L79:
	;
	goto L78
L80:
	;
	v263 = v248
	v265 = v250
	goto L81
L81:
	;
	v277 = v121 + v265<<(uint(int32(2))%32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v280
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+8)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v284
	v287 = v263 + int32(16)
	v289 = v265 + int32(4)
	if v289 != v54 {
		v263 = v287
		v265 = v289
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v291 = v287
	goto L71
L83:
	;
	goto L82
L84:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v389 != 0 {
		goto L97
	} else {
		goto L98
	}
L85:
	;
	v306 = (v57 - v209) & int32(3)
	if v306 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v209-v57) {
		v377 = v334
		goto L84
	} else {
		goto L93
	}
L87:
	;
	v334 = v291
	v335 = v209
	goto L86
L88:
	;
	goto L89
L89:
	;
	v310 = v291
	v311 = v209
	v312 = int32(0)
	goto L90
L90:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v122+v311<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v325
	v328 = v310 + int32(4)
	v329 = int32(1)
	v330 = v311 + v329
	v332 = v312 + v329
	if v332 != v306 {
		v310 = v328
		v311 = v330
		v312 = v332
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v334 = v328
	v335 = v330
	goto L86
L92:
	;
	goto L91
L93:
	;
	v349 = v334
	v350 = v335
	goto L94
L94:
	;
	v363 = v122 + v350<<(uint(int32(2))%32)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v364
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v363)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v368
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v363)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+12)) = v370
	v373 = v349 + int32(16)
	v375 = v350 + int32(4)
	if v375 != v57 {
		v349 = v373
		v350 = v375
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v377 = v373
	goto L84
L96:
	;
	goto L95
L97:
	;
	v397 = v389
	goto L99
L98:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v397 = (v390<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L99
L99:
	;
	v399 = v377 - (v397 + v118)
	v401 = v399 >> (uint(int32(2)) % 32)
	if v401 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v405 = F_construct_empty_array(m, int32(23))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v410 = F_ArrayGetNItemsSafe(m, v407, v118+int32(16))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L8
	} else {
		goto L104
	}
L103:
	;
	v516 = v405
	goto L1
L104:
	;
	if v410 == v401 {
		v516 = v118
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v413 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v421 = v413
	goto L108
L107:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v421 = (v414<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L108
L108:
	;
	v422 = v421 + v399
	v423 = F_repalloc(m, v118, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v422 << (uint(int32(2)) % 32)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	if v428 <= int32(0) {
		v516 = v423
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423)+16)) = v401
	if v428 == int32(1) {
		v516 = v423
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v435 = v423 + int32(16)
	v436 = int32(1)
	v437 = v428 - v436
	v438 = int32(7)
	v439 = v437 & v438
	if base.Ui32(v438) <= base.Ui32(v428-int32(2)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v448 = v436
	v450 = int32(0)
	goto L115
L113:
	;
	v478 = v436
	goto L114
L114:
	;
	v491 = v478
	v492 = int32(0)
	goto L119
L115:
	;
	v462 = v435 + v448<<(uint(int32(2))%32)
	v463 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v462)+24)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v462)+16)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v462)+8)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v462))) = v463
	v471 = int32(8)
	v472 = v448 + v471
	v474 = v450 + v471
	if v474 != v437&int32(-8) {
		v448 = v472
		v450 = v474
		goto L115
	} else {
		goto L117
	}
L116:
	;
	if v439 == int32(0) {
		v516 = v423
		goto L1
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	v478 = v472
	goto L114
L119:
	;
	v506 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v435+v491<<(uint(int32(2))%32)))) = v506
	v511 = v492 + v506
	if v511 != v439 {
		v491 = v491 + v506
		v492 = v511
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v516 = v423
	goto L1
L121:
	;
	goto L120
L122:
	;
	if int32(2) <= v528 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v532 = F__int_unique(m, v516)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L8
	} else {
		goto L126
	}
L124:
	;
	v534 = v516
	goto L125
L125:
	;
	return v534
L126:
	;
	v534 = v532
	goto L125
}
