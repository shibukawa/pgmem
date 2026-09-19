package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 == int32(2275) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L28
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L24
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != int32(1) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L20
	}
L6:
	;
	v16 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_deconstruct_array_builtin(m, l0, int32(2275), v8+int32(12), int32(0), l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = F_palloc(m, v26<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v31 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L19
	}
L15:
	;
	v40 = v37 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40)))
	v45 = F_pg_strtoint32(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v40))) = v45
	v49 = v37 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v49 < v50 {
		v37 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	m.G0 = v8 + int32(16)
	return v29
L20:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_ArrayGetIntegerTypmods_0), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ArrayGetIntegerTypmods_1), int32(242), int32(_a_F_ArrayGetIntegerTypmods_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_ArrayGetIntegerTypmods_3), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ArrayGetIntegerTypmods_1), int32(247), int32(_a_F_ArrayGetIntegerTypmods_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_ArrayGetIntegerTypmods_4), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ArrayGetIntegerTypmods_1), int32(252), int32(_a_F_ArrayGetIntegerTypmods_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_array_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = v15 + int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v48 = v45
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
	v45 = v41
	goto L2
L4:
	;
	v37 = int32(0)
	if v18 == v37 {
		v45 = v37
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	switch v23 - int32(429) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L4
	}
L6:
	;
	if v18 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v18 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v40 = v30
	v41 = int32(1)
	goto L3
L11:
	;
	v48 = int32(2)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
	v40 = v35
	v41 = int32(2)
	goto L3
L14:
	;
	v40 = v37
	v41 = v2
	goto L3
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v49 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L31
	} else {
		goto L156
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = v52
	goto L20
L19:
	;
	v53 = v2
	goto L20
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v54 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v15 + int32(16)
	return v574
L22:
	;
	if v53 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v57 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v53 != 0 {
		v574 = v53
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
	v574 = int32(0)
	goto L21
L28:
	;
	v64 = int32(_a_F_array_agg_array_combine_0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	v70 = int32(0)
	v72 = F_initArrayResultArr(m, v69, v70, v67, v70)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v124 <= int32(0) {
		v574 = v53
		goto L21
	} else {
		goto L44
	}
L31:
	;
	return int32(0)
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v76
	v78 = F_palloc(m, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v81 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v86 = base.I32_div_s(v82+int32(7), int32(8))
	v87 = F_palloc(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L31
	} else {
		goto L37
	}
L35:
	;
	v93 = v78
	goto L36
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v95 != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v87
	if v86 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	base.MemoryCopy(m, v87, v90, v86)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v93 = v92
	goto L36
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	base.MemoryCopy(m, v93, v96, v95)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+28)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v57)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+48)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v57)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+40)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v57)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+32)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v57)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+56)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v57)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+64)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v57)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+72)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+80)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+84)) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v65
	v574 = v72
	goto L21
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	if v127 == v128 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v214 = int32(_a_F_array_agg_array_combine_0)
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v217
	v219 = v130 + v131
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v220 < v219 {
		goto L65
	} else {
		goto L66
	}
L46:
	;
	v159 = int32(1)
	goto L55
L47:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v127 < int32(2) {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L31
	} else {
		goto L51
	}
L50:
	;
	v134 = int32(56)
	v138 = int32(32)
	goto L46
L51:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_array_agg_array_combine_1), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(1051), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v172 = v159 << (uint(int32(2)) % 32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v53+v138+v172)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v57+v138))))
	if v174 != v176 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L31
	} else {
		goto L61
	}
L57:
	;
	goto L56
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v53+v134))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v57+v134))))
	if v179 != v181 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v184 = v159 + int32(1)
	if v127 != v184 {
		v159 = v184
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L45
L61:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L31
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_array_agg_array_combine_1), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L31
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(1059), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L31
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
	v222 = int32(1)
	if v219&(v219-v222) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v237 != 0 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v230 = v222 << (uint(int32(32)-base.I32_clz(v219)) % 32)
	goto L70
L69:
	;
	v230 = v219
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v233 = F_repalloc(m, v232, v230)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v233
	goto L67
L72:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v240 = v238 + v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v241 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v553 != 0 {
		goto L153
	} else {
		goto L154
	}
L75:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v423 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v424 <= v423 {
		goto L123
	} else {
		goto L124
	}
L76:
	;
	v244 = int32(1)
	v246 = int32(256)
	v248 = v240 + v244
	if v248 <= v246 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	if v240 <= v397 {
		v418 = v241
		goto L75
	} else {
		goto L117
	}
L79:
	;
	v251 = v246
	goto L81
L80:
	;
	v251 = v248
	goto L81
L81:
	;
	if v251&(v251-int32(1)) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v258 = v244 << (uint(int32(32)-base.I32_clz(v251)) % 32)
	goto L84
L83:
	;
	v258 = v251
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v258
	v263 = base.I32_div_s(v258+int32(7), int32(8))
	v264 = F_palloc(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L31
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v264
	v267 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	if v270 <= v267 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v418 = v396
	goto L75
L87:
	;
	goto L86
L88:
	;
	v282 = base.I32_div_s(v267, int32(8))
	v283 = v264 + v282
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	goto L90
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v311)
	goto L87
L90:
	;
	v287 = v283
	v288 = v284
	v291 = v270
	v292 = int32(1)
	goto L93
L93:
	;
	v296 = v288 | v292
	v297 = int32(1)
	v298 = v291 - v297
	v300 = v292 << (uint(v297) % 32)
	if v300 == int32(256) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v312 != int32(1) {
		goto L89
	} else {
		goto L100
	}
L95:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v287))) = uint8(v296)
	if v298 == int32(0) {
		goto L87
	} else {
		goto L98
	}
L96:
	;
	v310 = v287
	v311 = v296
	v312 = v300
	goto L97
L97:
	;
	if base.Ui32(int32(1)) < base.Ui32(v291) {
		v287 = v310
		v288 = v311
		v291 = v298
		v292 = v312
		goto L93
	} else {
		goto L99
	}
L98:
	;
	v306 = int32(1)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	v310 = v287 + v306
	v311 = v307
	v312 = v306
	goto L97
L99:
	;
	goto L94
L100:
	;
	goto L87
L117:
	;
	v399 = int32(1)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v402 = v401 + v397
	if v402&(v402-v399) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v409 = v399 << (uint(int32(32)-base.I32_clz(v402)) % 32)
	goto L120
L119:
	;
	v409 = v402
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v409
	v414 = base.I32_div_s(v409+int32(7), int32(8))
	v415 = F_repalloc(m, v241, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L31
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v415
	v418 = v415
	goto L75
L122:
	;
	goto L74
L123:
	;
	goto L122
L124:
	;
	v434 = int32(1) << (uint(v421&int32(7)) % 32)
	v436 = base.I32_div_s(v421, int32(8))
	v437 = v418 + v436
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v422 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v531))) = uint8(v532)
	goto L123
L126:
	;
	v441 = v437
	v442 = v438
	v445 = v424
	v446 = v434
	goto L129
L127:
	;
	goto L128
L128:
	;
	v476 = base.I32_div_s(v423, int32(8))
	v477 = v422 + v476
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	v479 = v437
	v480 = v438
	v482 = v477
	v483 = v424
	v484 = v434
	v485 = int32(1)
	v486 = v478
	goto L137
L129:
	;
	v450 = v442 | v446
	v451 = int32(1)
	v452 = v445 - v451
	v454 = v446 << (uint(v451) % 32)
	if v454 == int32(256) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v466 != int32(1) {
		v531 = v464
		v532 = v465
		goto L125
	} else {
		goto L136
	}
L131:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v441))) = uint8(v450)
	if v452 == int32(0) {
		goto L123
	} else {
		goto L134
	}
L132:
	;
	v464 = v441
	v465 = v450
	v466 = v454
	goto L133
L133:
	;
	if base.Ui32(int32(1)) < base.Ui32(v445) {
		v441 = v464
		v442 = v465
		v445 = v452
		v446 = v466
		goto L129
	} else {
		goto L135
	}
L134:
	;
	v460 = int32(1)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	v464 = v441 + v460
	v465 = v461
	v466 = v460
	goto L133
L135:
	;
	goto L130
L136:
	;
	goto L123
L137:
	;
	if v485&v486 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v509 == int32(1) {
		goto L123
	} else {
		goto L152
	}
L139:
	;
	v493 = v480 | v484
	goto L141
L140:
	;
	v493 = v480 & (v484 ^ int32(-1))
	goto L141
L141:
	;
	v494 = int32(1)
	v495 = v483 - v494
	v497 = v484 << (uint(v494) % 32)
	if v497 == int32(256) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v493)
	if v495 == int32(0) {
		goto L123
	} else {
		goto L145
	}
L143:
	;
	v507 = v479
	v508 = v493
	v509 = v497
	goto L144
L144:
	;
	v511 = v485 << (uint(int32(1)) % 32)
	if v511 == int32(256) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v503 = int32(1)
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
	v507 = v479 + v503
	v508 = v504
	v509 = v503
	goto L144
L146:
	;
	goto L138
L147:
	;
	if v495 == int32(0) {
		goto L146
	} else {
		goto L150
	}
L148:
	;
	v520 = v482
	v521 = v511
	v522 = v486
	goto L149
L149:
	;
	if base.Ui32(int32(1)) < base.Ui32(v483) {
		v479 = v507
		v480 = v508
		v482 = v520
		v483 = v495
		v484 = v509
		v485 = v521
		v486 = v522
		goto L137
	} else {
		goto L151
	}
L150:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	v517 = int32(1)
	v520 = v482 + v517
	v521 = v517
	v522 = v516
	goto L149
L151:
	;
	goto L146
L152:
	;
	v531 = v507
	v532 = v508
	goto L125
L153:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	base.MemoryCopy(m, v554+v555, v557, v553)
	goto L155
L154:
	;
	goto L155
L155:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v559 + v560
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = v563 + v564
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+32)) = v567 + v568
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v215
	v574 = v53
	goto L21
L156:
	;
	F_errmsg_internal(m, int32(_a_F_array_agg_array_combine_4), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L31
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(985), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L31
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_array_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(1)
		v17 = v12 + v16
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		v22 = v20 & v16
		if v22 != 0 {
			v23 = v17
		} else {
			v23 = v12 + int32(4)
		}
		if v20 == int32(1) {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v29 == int32(18) {
				v32 = int32(16)
			} else {
				v32 = int32(0)
			}
			if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v39 = int32(4)
			} else {
				v39 = v32
			}
			v50 = v39
		} else {
			v40 = int32(1)
			if v22 != 0 {
				v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v23
		v56 = F_pq_getmsgint(m, v9, int32(4))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v59 = F_pq_getmsgint(m, v9, int32(4))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v62 = F_pq_getmsgint(m, v9, int32(4))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_deserialize[0]))
					v68 = F_initArrayResultArr(m, v59, v56, v66, int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v73 = int32(1024)
						for {
							if v73 < v62 {
								v73 = v73 << (uint(int32(1)) % 32)
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v73
						v80 = F_palloc(m, v73)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v80
							v83 = F_pq_getmsgbytes(m, v9, v62)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v62 != 0 {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
									base.MemoryCopy(m, v85, v83, v62)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v62
								v89 = F_pq_getmsgint(m, v9, int32(4))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v89
									v93 = F_pq_getmsgint(m, v9, int32(4))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v93
										if int32(0) < v93 {
											v101 = base.I32_div_s(v93+int32(7), int32(8))
											v102 = F_palloc(m, v101)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v102
												v105 = F_pq_getmsgbytes(m, v9, v101)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													if v101 == int32(0) {
													} else {
														v109 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
														base.MemoryCopy(m, v109, v105, v101)
													}
													v116 = F_pq_getmsgint(m, v9, int32(4))
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = v116
														v120 = F_pq_getmsgint(m, v9, int32(4))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v68)+28)) = v120
															v124 = F_pq_getmsgbytes(m, v9, int32(24))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																v126 = *(*int64)(unsafe.Add(mBase, uint32(v124)+16))
																*(*int64)(unsafe.Add(mBase, uint32(v68)+48)) = v126
																v128 = *(*int64)(unsafe.Add(mBase, uint32(v124)+8))
																*(*int64)(unsafe.Add(mBase, uint32(v68)+40)) = v128
																v130 = *(*int64)(unsafe.Add(mBase, uint32(v124)))
																*(*int64)(unsafe.Add(mBase, uint32(v68)+32)) = v130
																v133 = F_pq_getmsgbytes(m, v9, int32(24))
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int32(0)
																} else {
																	v135 = *(*int64)(unsafe.Add(mBase, uint32(v133)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v68)+72)) = v135
																	v137 = *(*int64)(unsafe.Add(mBase, uint32(v133)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v68)+64)) = v137
																	v139 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
																	*(*int64)(unsafe.Add(mBase, uint32(v68)+56)) = v139
																	F_pq_getmsgend(m, v9)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v9 + int32(16)
																		return v68
																	}
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = int32(0)
											v116 = F_pq_getmsgint(m, v9, int32(4))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = v116
												v120 = F_pq_getmsgint(m, v9, int32(4))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v68)+28)) = v120
													v124 = F_pq_getmsgbytes(m, v9, int32(24))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														v126 = *(*int64)(unsafe.Add(mBase, uint32(v124)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v68)+48)) = v126
														v128 = *(*int64)(unsafe.Add(mBase, uint32(v124)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v68)+40)) = v128
														v130 = *(*int64)(unsafe.Add(mBase, uint32(v124)))
														*(*int64)(unsafe.Add(mBase, uint32(v68)+32)) = v130
														v133 = F_pq_getmsgbytes(m, v9, int32(24))
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return int32(0)
														} else {
															v135 = *(*int64)(unsafe.Add(mBase, uint32(v133)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v68)+72)) = v135
															v137 = *(*int64)(unsafe.Add(mBase, uint32(v133)+8))
															*(*int64)(unsafe.Add(mBase, uint32(v68)+64)) = v137
															v139 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
															*(*int64)(unsafe.Add(mBase, uint32(v68)+56)) = v139
															F_pq_getmsgend(m, v9)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v68
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_array_create_iterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v9 = F_palloc0(m, int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_array_create_iterator_0), int32(0))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_array_create_iterator_1), int32(_a_F_array_create_iterator_2), int32(_a_F_array_create_iterator_3))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 < l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_array_create_iterator_0), int32(0))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_create_iterator_1), int32(_a_F_array_create_iterator_2), int32(_a_F_array_create_iterator_3))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v26 = l0 + v19<<(uint(int32(3))%32) + int32(16)
				} else {
					v26 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v30 = l0 + int32(16)
				v31 = F_ArrayGetNItemsSafe(m, v28, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v31
					if l2 != 0 {
						v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v34)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v36)
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v38)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
						if l1 != 0 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v51 = int32(2)
							v55 = l1 << (uint(v51) % 32)
							v56 = v30 + v50<<(uint(v51)%32) - v55
							*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v56
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v60 = v58 << (uint(v51) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v30 + v60 + v60 - v55
							v65 = F_ArrayGetNItemsSafe(m, l1, v56)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v65
								v70 = F_palloc(m, v65<<(uint(int32(2))%32))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v70
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v74 = F_palloc(m, v73)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v74
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v81 == int32(0) {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v91 = v81
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
										return v9
									}
								}
							}
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v81 == int32(0) {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v91 = v81
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
							return v9
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_get_typlenbyvalalign(m, v40, v9+int32(12), v9+int32(14), v9+int32(15))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
							if l1 != 0 {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v51 = int32(2)
								v55 = l1 << (uint(v51) % 32)
								v56 = v30 + v50<<(uint(v51)%32) - v55
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v56
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v60 = v58 << (uint(v51) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v30 + v60 + v60 - v55
								v65 = F_ArrayGetNItemsSafe(m, l1, v56)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v65
									v70 = F_palloc(m, v65<<(uint(int32(2))%32))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v70
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
										v74 = F_palloc(m, v73)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v74
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v81 == int32(0) {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v91 = v81
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
											return v9
										}
									}
								}
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v81 == int32(0) {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v91 = v81
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
								return v9
							}
						}
					}
				}
			}
		}
	}
}
func F_array_dim_to_json(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	F_appendStringInfoChar(m, l0, int32(91))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = l3 + l1<<(uint(int32(2))%32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if int32(0) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l9 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v25 = int32(_a_F_array_dim_to_json_0)
	goto L8
L7:
	;
	v25 = int32(_a_F_array_dim_to_json_1)
	goto L8
L8:
	;
	v26 = int32(1)
	v27 = l1 + v26
	v30 = v26
	goto L9
L9:
	;
	if int32(2) <= v30 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	F_appendStringInfoString(m, l0, v25)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l2 == v27 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v66 = v30 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v66 <= v67 {
		v30 = v66
		goto L9
	} else {
		goto L21
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4+v47<<(uint(int32(2))%32))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v47))))
	F_datum_to_json_internal(m, v51, v53, l0, l7, l8, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_array_dim_to_json(m, l0, v27, l2, l3, l4, l5, l6, l7, l8, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v57 + int32(1)
	goto L15
L20:
	;
	goto L15
L21:
	;
	goto L10
L22:
	;
	return
}
func F_array_dim_to_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v15 = F_pushJsonbValue(m, l0, int32(4), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
	v20 = l3 + l1<<(uint(int32(2))%32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if int32(0) < v21 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(1)
	v25 = l1 + v24
	v28 = v24
	goto L6
L4:
	;
	goto L5
L5:
	;
	v75 = F_pushJsonbValue(m, l0, int32(5), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	if l2 == v25 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v58 = v28 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v58 <= v59 {
		v28 = v58
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l4+v40<<(uint(int32(2))%32))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v40))))
	F_datum_to_jsonb_internal(m, v44, v46, l0, l7, l8, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_array_dim_to_jsonb(m, l0, v25, l2, l3, l4, l5, l6, l7, l8)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v50 + int32(1)
	goto L8
L13:
	;
	goto L8
L14:
	;
	goto L7
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75
	return
}
func F_array_exec_setup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v11 < int32(7) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v14 != v11 {
			v17 = v14
		} else {
			v17 = int32(0)
		}
		if v17 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_array_exec_setup_0), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_array_exec_setup_1), int32(495), int32(_a_F_array_exec_setup_2))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = F_palloc(m, int32(60))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v19
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v25 = F_get_typlen(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)) = uint16(v25)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_get_typlenbyvalalign(m, v28, v19+int32(6), v19+int32(8), v19+int32(9))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v14 != 0 {
							v39 = int32(1250)
						} else {
							v39 = int32(1251)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v39
						if v14 != 0 {
							v43 = int32(1252)
						} else {
							v43 = int32(1253)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v43
						if v14 != 0 {
							v47 = int32(1254)
						} else {
							v47 = int32(1255)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1256)
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v61
				F_errmsg(m, int32(_a_F_array_exec_setup_3), v9)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_array_exec_setup_1), int32(490), int32(_a_F_array_exec_setup_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_array_fill_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v421 int32
	_ = v421
	var v424 int64
	_ = v424
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 < int32(2) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L15
	} else {
		goto L156
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L15
	} else {
		goto L151
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L15
	} else {
		goto L147
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L15
	} else {
		goto L142
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L15
	} else {
		goto L138
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L15
	} else {
		goto L134
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L15
	} else {
		goto L130
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L15
	} else {
		goto L125
	}
L11:
	;
	if v156 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	v25 = l0 + int32(16)
	v26 = F_ArrayGetNItemsSafe(m, v20, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v135 = v20
	goto L14
L14:
	;
	v156 = v135
	v161 = (v135<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L11
L15:
	;
	return int32(0)
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v30 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = v25 + v31<<(uint(int32(3))%32)
	goto L19
L18:
	;
	v36 = int32(0)
	goto L19
L19:
	;
	if int32(8) <= v26 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v52 = v26
	v53 = v36
	goto L23
L21:
	;
	v76 = v26
	v77 = v36
	goto L22
L22:
	;
	if int32(0) < v76 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v61 != int32(255) {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v76 = v69
	v77 = v36 + int32(base.Ui32(v26-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L22
L25:
	;
	v69 = v52 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v52) {
		v52 = v69
		v53 = v53 + int32(1)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v95 = v76
	v96 = int32(1)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 != 0 {
		v156 = v128
		v161 = v30
		goto L11
	} else {
		goto L34
	}
L30:
	;
	if v96&v87 == int32(0) {
		goto L7
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v107 = int32(1)
	if v107 < v95 {
		v95 = v95 - v107
		v96 = v96 << (uint(v107) % 32)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v135 = v128
	goto L14
L35:
	;
	if l1 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	v173 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v168 < int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v168) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v173 = v168
	goto L35
L41:
	;
	v323 = l0 + v161
	v324 = F_ArrayGetNItemsSafe(m, v173, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L15
	} else {
		goto L71
	}
L42:
	;
	v176 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v176
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v176
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v176
	v322 = v18 + int32(48)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(2) <= v184 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v187 == int32(0) {
		v290 = int32(0)
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v292 {
		goto L64
	} else {
		goto L65
	}
L47:
	;
	v191 = l1 + int32(16)
	v192 = F_ArrayGetNItemsSafe(m, v184, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v194 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v200 = v191 + v195<<(uint(int32(3))%32)
	goto L51
L50:
	;
	v200 = int32(0)
	goto L51
L51:
	;
	if int32(8) <= v192 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v216 = v192
	v217 = v200
	goto L55
L53:
	;
	v240 = v192
	v241 = v200
	goto L54
L54:
	;
	if v240 <= int32(0) {
		v290 = v194
		goto L46
	} else {
		goto L59
	}
L55:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v225 != int32(255) {
		goto L3
	} else {
		goto L57
	}
L56:
	;
	v240 = v233
	v241 = v200 + int32(base.Ui32(v192-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L54
L57:
	;
	v233 = v216 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v216) {
		v216 = v233
		v217 = v217 + int32(1)
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v259 = v240
	v260 = int32(1)
	goto L60
L60:
	;
	if v260&v251 == int32(0) {
		goto L3
	} else {
		goto L62
	}
L61:
	;
	v290 = v194
	goto L46
L62:
	;
	v271 = int32(1)
	if v271 < v259 {
		v259 = v259 - v271
		v260 = v260 << (uint(v271) % 32)
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v297 = v295
	goto L66
L65:
	;
	v297 = int32(0)
	goto L66
L66:
	;
	if v297 != v173 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	if v290 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v305 = v290
	goto L70
L69:
	;
	v305 = (v292<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L70
L70:
	;
	v322 = l1 + v305
	goto L41
L71:
	;
	F_ArrayCheckBounds(m, v173, v323, v322)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L15
	} else {
		goto L72
	}
L72:
	;
	if v324 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	m.G0 = v18 + int32(80)
	return v521
L74:
	;
	v331 = F_palloc0(m, int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L15
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	if v339 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v331)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v331))) = int64(64)
	v521 = v331
	goto L73
L78:
	;
	if l3 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	F_get_typlenbyvalalign(m, l4, v355+int32(4), v355+int32(6), v355+int32(7))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L15
	} else {
		goto L85
	}
L80:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338)+20))
	v344 = F_MemoryContextAlloc(m, v342, int32(48))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L15
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	if v352 == l4 {
		v366 = v339
		goto L78
	} else {
		goto L84
	}
L83:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v346)+16)) = v344
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = int32(0)
	v355 = v349
	goto L79
L84:
	;
	v355 = v339
	goto L79
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = l4
	v366 = v355
	goto L78
L86:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+6)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+7)))
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366)+4)))
	if v371 != int32(-1) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	goto L88
L88:
	;
	v493 = base.I32_div_s(v324+int32(7), int32(8))
	v500 = (v493 + v173<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v501 = F_palloc0(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L15
	} else {
		goto L120
	}
L89:
	;
	switch v370 - int32(99) {
	case 0:
		v421 = v406
		goto L104
	case 1:
		goto L106
	default:
		goto L105
	case 6:
		goto L107
	}
L90:
	;
	v401 = F_strlen(m, l2)
	mBase = m.M
	v405 = l2
	v406 = v401 + int32(1)
	goto L89
L91:
	;
	if v371 <= int32(0) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v376 = F_pg_detoast_datum(m, l2)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L15
	} else {
		goto L95
	}
L94:
	;
	v405 = l2
	v406 = v371
	goto L89
L95:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v378 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+1)))
	if base.Ui32((v382-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v405 = v376
		v406 = int32(6)
		goto L89
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v394 = int32(1)
	if v378&v394 != 0 {
		v405 = v376
		v406 = int32(base.Ui32(v378) >> (uint(v394) % 32))
		goto L89
	} else {
		goto L103
	}
L99:
	;
	v389 = int32(18)
	if v382 == v389 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v393 = v389
	goto L102
L101:
	;
	v393 = int32(2)
	goto L102
L102:
	;
	v405 = v376
	v406 = v393
	goto L89
L103:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v405 = v376
	v406 = int32(base.Ui32(v398) >> (uint(int32(2)) % 32))
	goto L89
L104:
	;
	v424 = base.I64_extend_i32_s(v421) * base.I64_extend_i32_s(v324)
	v428 = base.I32_wrap_i64(v424)
	if base.B2i32(base.I32_wrap_i64(int64(base.Ui64(v424)>>(uint(int64(32))%64))) != v428>>(uint(int32(31))%32))|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v428)) != 0 {
		goto L1
	} else {
		goto L108
	}
L105:
	;
	v421 = (v406 + int32(1)) & int32(-2)
	goto L104
L106:
	;
	v421 = (v406 + int32(7)) & int32(-8)
	goto L104
L107:
	;
	v421 = (v406 + int32(3)) & int32(-4)
	goto L104
L108:
	;
	v440 = (v173<<(uint(int32(3))%32) + int32(23)) & int32(120)
	v441 = v428 + v440
	v442 = F_palloc0(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+12)) = l4
	v445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+8)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v442)+4)) = v173
	v449 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v441 << (uint(v449) % 32)
	v453 = v442 + int32(16)
	v455 = v173 << (uint(v449) % 32)
	v457 = base.B2i32(v455 == v445)
	if v457 == v445 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	base.MemoryCopy(m, v453, v323, v455)
	goto L112
L111:
	;
	goto L112
L112:
	;
	if v457 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	base.MemoryCopy(m, v455+v453, v322, v455)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v475 = v442 + v440
	v476 = v445
	goto L116
L116:
	;
	v484 = F_ArrayCastAndSet(m, v405, v371, v369&int32(1), base.I32_extend8_s(v370), v475)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L15
	} else {
		goto L118
	}
L117:
	;
	v521 = v442
	goto L73
L118:
	;
	v488 = v476 + int32(1)
	if v488 != v324 {
		v475 = v484 + v475
		v476 = v488
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v501)+8)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v173
	v506 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v500 << (uint(v506) % 32)
	v510 = v501 + int32(16)
	v512 = v173 << (uint(v506) % 32)
	v513 = int32(0)
	v514 = base.B2i32(v512 == v513)
	if v514 == v513 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	base.MemoryCopy(m, v510, v323, v512)
	goto L123
L122:
	;
	goto L123
L123:
	;
	if v512 == v513 {
		v521 = v501
		goto L73
	} else {
		goto L124
	}
L124:
	;
	base.MemoryCopy(m, v512+v510, v322, v512)
	v521 = v501
	goto L73
L125:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_0), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L15
	} else {
		goto L127
	}
L127:
	;
	F_errdetail(m, int32(_a_F_array_fill_internal_1), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_3), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L15
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L15
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_5), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_6), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L15
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L15
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v168
	F_errmsg(m, int32(_a_F_array_fill_internal_7), v18)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L15
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_8), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L15
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L15
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v168
	F_errmsg(m, int32(_a_F_array_fill_internal_9), v18+int32(16))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L15
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_10), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L15
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L15
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_0), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L15
	} else {
		goto L144
	}
L144:
	;
	F_errdetail(m, int32(_a_F_array_fill_internal_1), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L15
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_11), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L15
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L15
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_5), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L15
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_12), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L15
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L15
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_0), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L15
	} else {
		goto L153
	}
L153:
	;
	F_errdetail(m, int32(_a_F_array_fill_internal_13), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L15
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_14), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L15
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L15
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_fill_internal_15), v18+int32(32))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L15
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_2), int32(_a_F_array_fill_internal_16), int32(_a_F_array_fill_internal_4))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L15
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_free_iterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v2 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		F_pfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_pfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_array_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
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
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v434 int32
	_ = v434
	var v452 int32
	_ = v452
	v9 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	if v9 < l3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v452
L2:
	;
	v434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v434)
	v452 = int32(0)
	goto L1
L3:
	;
	if base.B2i32(l1 != v231)|base.B2i32(base.Ui32(v231-int32(7)) < base.Ui32(int32(-6))) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	v25 = base.I32_div_s(l3, l4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v25
	v229 = l0
	v230 = v18 + int32(12)
	v231 = int32(1)
	v232 = v9
	v235 = v18 + int32(8)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v31 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v209 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L32
	} else {
		goto L37
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v34&int32(254) != int32(2) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if base.B2i32(v40 != l1)|base.B2i32(base.Ui32(v40-int32(7)) < base.Ui32(int32(-6))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v98 = int32(0)
	v107 = l1 - int32(1)
	if v107 < v98 {
		v185 = v98
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v62 = v51
	goto L16
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v51 = int32(0)
	if l1 <= v51 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L2
L15:
	;
	goto L11
L16:
	;
	v70 = v62 << (uint(int32(2)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+v70)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v49+v70)))
	if v72 < v74 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L2
L18:
	;
	goto L17
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v50+v70)))
	if v77+v74 <= v72 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v81 = v62 + int32(1)
	if l1 != v81 {
		v62 = v81
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	F_deconstruct_expanded_array(m, v39)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	goto L22
L24:
	;
	v110 = int32(1)
	if v107 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v119 = v107
	v120 = v110
	v121 = v98
	v126 = v98
	goto L28
L26:
	;
	v162 = v107
	v163 = v110
	v164 = v98
	goto L27
L27:
	;
	v171 = v162 << (uint(int32(2)) % 32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l2+v171)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v49)))
	v185 = (v173-v175)*v163 + v164
	goto L23
L28:
	;
	v127 = int32(2)
	v128 = v119 << (uint(v127) % 32)
	v130 = v128 - int32(4)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2+v130)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v49+v130)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v128+v50)))
	v138 = v137 * v120
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128+l2)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v128+v49)))
	v147 = (v132-v134)*v138 + ((v141-v143)*v120 + v121)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v50+v130)))
	v150 = v149 * v138
	v152 = v119 - v127
	v154 = v126 + v127
	if v154 != l1&int32(-2) {
		v119 = v152
		v120 = v150
		v121 = v147
		v126 = v154
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if l1&int32(1) == int32(0) {
		v185 = v147
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v162 = v152
	v163 = v150
	v164 = v147
	goto L27
L32:
	;
	return int32(0)
L33:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	if v196 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v203)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v195+v185<<(uint(int32(2))%32))))
	v452 = v208
	goto L1
L35:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v185))))
	if v200 != int32(1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	v212 = v209 + int32(16)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v215 = v213 << (uint(int32(3)) % 32)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	if v218 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v219 = v212 + v215
	goto L40
L39:
	;
	v219 = int32(0)
	goto L40
L40:
	;
	if v218 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v224 = v218
	goto L43
L42:
	;
	v224 = (v215 + int32(23)) & int32(-8)
	goto L43
L43:
	;
	v229 = v209 + v224
	v230 = v212
	v231 = v213
	v232 = v219
	v235 = v212 + v213<<(uint(int32(2))%32)
	goto L3
L44:
	;
	v291 = int32(0)
	v300 = l1 - int32(1)
	if v300 < v291 {
		v378 = v291
		goto L57
	} else {
		goto L58
	}
L45:
	;
	v255 = v244
	goto L50
L46:
	;
	v244 = int32(0)
	if l1 <= v244 {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L2
L49:
	;
	goto L45
L50:
	;
	v263 = v255 << (uint(int32(2)) % 32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l2+v263)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v235+v263)))
	if v265 < v267 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L2
L52:
	;
	goto L51
L53:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v230+v263)))
	if v270+v267 <= v265 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v274 = v255 + int32(1)
	if l1 != v274 {
		v255 = v274
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	if v232 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	goto L56
L58:
	;
	v303 = int32(1)
	if v300 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v312 = v300
	v313 = v303
	v314 = v291
	v319 = v291
	goto L62
L60:
	;
	v355 = v300
	v356 = v303
	v357 = v291
	goto L61
L61:
	;
	v364 = v355 << (uint(int32(2)) % 32)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l2+v364)))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364+v235)))
	v378 = (v366-v368)*v356 + v357
	goto L57
L62:
	;
	v320 = int32(2)
	v321 = v312 << (uint(v320) % 32)
	v323 = v321 - int32(4)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l2+v323)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v235+v323)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v321+v230)))
	v331 = v330 * v313
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v321+l2)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v321+v235)))
	v340 = (v325-v327)*v331 + ((v334-v336)*v313 + v314)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v230+v323)))
	v343 = v342 * v331
	v345 = v312 - v320
	v347 = v319 + v320
	if v347 != l1&int32(-2) {
		v312 = v345
		v313 = v343
		v314 = v340
		v319 = v347
		goto L62
	} else {
		goto L64
	}
L63:
	;
	if l1&int32(1) == int32(0) {
		v378 = v340
		goto L57
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v355 = v345
	v356 = v343
	v357 = v340
	goto L61
L66:
	;
	v395 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v395)
	v398 = F_array_seek(m, v229, v395, v232, v378, l4, l6)
	mBase = m.M
	if l5 == v395 {
		v452 = v398
		goto L1
	} else {
		goto L69
	}
L67:
	;
	v387 = base.I32_div_s(v378, int32(8))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232+v387))))
	if int32(base.Ui32(v389)>>(uint(v378&int32(7))%32))&int32(1) != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	goto L2
L69:
	;
	switch l4 - int32(1) {
	case 0:
		goto L73
	case 1:
		goto L72
	default:
		goto L70
	case 3:
		goto L71
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L32
	} else {
		goto L74
	}
L71:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v452 = v405
	goto L1
L72:
	;
	v404 = int32(*(*int16)(unsafe.Add(mBase, uint32(v398))))
	v452 = v404
	goto L1
L73:
	;
	v403 = int32(*(*int8)(unsafe.Add(mBase, uint32(v398))))
	v452 = v403
	goto L1
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l4
	F_errmsg_internal(m, int32(_a_F_array_get_element_0), v18)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L32
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_array_get_element_1), int32(70), int32(_a_F_array_get_element_2))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L32
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = int32(-1)
	v9 = F_array_get_element(m, l0, int32(1), l1, v5, v5, int32(0), int32(105), l2)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_array_subscript_fetch_old_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 != 0 {
		v28 = int32(1)
		v31 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)) = uint8(v28)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v31
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
		v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
		v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
		v26 = F_array_get_slice(m, v17, v18, v12+int32(12), v12+int32(36), v21, v22, v23, v24, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = int32(0)
			v31 = v26
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)) = uint8(v28)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v31
			return
		}
	}
}
func F_array_to_halfvec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 float32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L69
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L61
	}
L4:
	;
	return int32(0)
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v15 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L57
	}
L9:
	;
	v20 = F_array_contains_nulls(m, v11)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_get_typlenbyvalalign(m, v22, v8+int32(30), v8+int32(29), v8+int32(28))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	if v20 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+30)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
	v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+28)))
	F_deconstruct_array(m, v11, v32, v33, v34, v8+int32(24), int32(0), v8+int32(20))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	F_CheckDim_1(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if base.B2i32(v18 != int32(-1))&base.B2i32(v47 != v18) != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v52 = F_mul_size(m, int32(2), v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v54 = F_add_size(m, int32(8), v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v56 = F_palloc0(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v54 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	switch v62 - int32(700) {
	case 0:
		goto L24
	case 1:
		goto L23
	default:
		goto L25
	}
L21:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	F_pfree(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L49
	}
L22:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v154 <= int32(0) {
		goto L21
	} else {
		goto L44
	}
L23:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v126 <= int32(0) {
		goto L21
	} else {
		goto L39
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v100 <= int32(0) {
		goto L21
	} else {
		goto L34
	}
L25:
	;
	if v62 == int32(23) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v62 != int32(1700) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v69 <= int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v75 = int32(0)
	goto L29
L29:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v75<<(uint(int32(2))%32))))
	v90 = F_DirectFunctionCall1Coll(m, int32(1319), int32(0), v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L21
L31:
	;
	v93 = F_Float4ToHalf(m, base.F32_reinterpret_i32(v90))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v75<<(uint(int32(1))%32)))) = uint16(v93)
	v97 = v75 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v97 < v98 {
		v75 = v97
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v106 = int32(0)
	goto L35
L35:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v118 = *(*float32)(unsafe.Add(mBase, uint32(v114+v106<<(uint(int32(2))%32))))
	v119 = F_Float4ToHalf(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	goto L21
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v106<<(uint(int32(1))%32)))) = uint16(v119)
	v123 = v106 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v123 < v124 {
		v106 = v123
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v132 = int32(0)
	goto L40
L40:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v132<<(uint(int32(2))%32))))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v144)))
	v147 = F_Float4ToHalf(m, base.F32_demote_f64(v145))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L21
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v132<<(uint(int32(1))%32)))) = uint16(v147)
	v151 = v132 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v151 < v152 {
		v132 = v151
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v160 = int32(0)
	goto L45
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v160<<(uint(int32(2))%32))))
	v174 = F_Float4ToHalf(m, base.F32_convert_i32_s(v172))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L21
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v160<<(uint(int32(1))%32)))) = uint16(v174)
	v178 = v160 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v178 < v179 {
		v160 = v178
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+4)))
	if int32(0) < v189 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v195 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	m.G0 = v8 + int32(32)
	return v56
L53:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v195<<(uint(int32(1))%32)))))
	F_CheckElement_1(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	v207 = v195 + int32(1)
	v208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+4)))
	if v207 < v208 {
		v195 = v207
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_array_to_halfvec_0), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(456), int32(_a_F_array_to_halfvec_2))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_array_to_halfvec_3), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(461), int32(_a_F_array_to_halfvec_2))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
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
	F_errcode(m, int32(130))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
	F_errmsg(m, int32(_a_F_array_to_halfvec_4), v8)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(92), int32(_a_F_array_to_halfvec_5))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_array_to_halfvec_6), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(495), int32(_a_F_array_to_halfvec_2))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_to_json_pretty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_makeStringInfo(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_array_to_json_internal(m, v4, v5, base.B2i32(v3 != int32(0)))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			v15 = F_cstring_to_text_with_len(m, v13, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_array_to_text_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v22 = l1 + int32(16)
	v23 = F_ArrayGetNItemsSafe(m, v20, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 - int32(-64)
	return v277
L2:
	;
	return int32(0)
L3:
	;
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = F_palloc(m, int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_initStringInfo(m, v16+int32(-16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(16)
	v277 = v30
	goto L1
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if int32(0) < v23 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	F_get_type_io_data(m, v34, int32(1), v56+int32(4), v56+int32(6), v56+int32(7), v56+int32(8), v56+int32(12), v56+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L16
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v45 = F_MemoryContextAlloc(m, v43, int32(48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v54 == v34 {
		v80 = v40
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v34 ^ int32(-1)
	v56 = v50
	goto L10
L15:
	;
	v56 = v40
	goto L10
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	F_fmgr_info_cxt(m, v72, v56+int32(20), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v34
	v80 = v56
	goto L9
L18:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	v84 = base.I32_extend16_s(v83)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v87 = v85 << (uint(int32(3)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v90 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v265 = v263 + int32(4)
	v266 = F_palloc(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L86
	}
L21:
	;
	v91 = v22 + v87
	goto L23
L22:
	;
	v91 = int32(0)
	goto L23
L23:
	;
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = v90
	goto L26
L25:
	;
	v98 = (v87 + int32(23)) & int32(-8)
	goto L26
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+6)))
	v101 = int32(1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+7)))
	v108 = int32(0)
	v111 = v91
	v112 = l1 + v98
	v115 = v101
	v117 = v108
	v120 = v108
	goto L27
L27:
	;
	if v111 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L20
L29:
	;
	v234 = int32(1)
	v236 = v115 << (uint(v234) % 32)
	v238 = base.B2i32(v236 == int32(256))
	if v236 == int32(256) {
		goto L76
	} else {
		goto L77
	}
L30:
	;
	v231 = v227
	v232 = int32(1)
	goto L29
L31:
	;
	if v100&v101 != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v115&v128 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if l3 == int32(0) {
		v231 = v112
		v232 = v117
		goto L29
	} else {
		goto L34
	}
L34:
	;
	if v117 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l2
	F_appendStringInfo(m, v16+int32(-16), int32(_a_F_array_to_text_internal_0), v16+int32(-32))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_appendStringInfoString(m, v16+int32(-16), l3)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L39
	}
L38:
	;
	v227 = v112
	goto L30
L39:
	;
	v227 = v112
	goto L30
L40:
	;
	v162 = F_OutputFunctionCall(m, v80+int32(20), v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L51
	}
L41:
	;
	switch v83 - v101 {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L44
	case 3:
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v161 = v112
	goto L40
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L48
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v161 = v147
	goto L40
L46:
	;
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112))))
	v161 = v146
	goto L40
L47:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	v161 = v145
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v84
	F_errmsg_internal(m, int32(_a_F_array_to_text_internal_1), v18)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_array_to_text_internal_2), int32(70), int32(_a_F_array_to_text_internal_3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v117 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if int32(0) < v84 {
		v214 = v112 + v84
		goto L58
	} else {
		goto L59
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l2
	F_appendStringInfo(m, v16+int32(-16), int32(_a_F_array_to_text_internal_0), v16+int32(-48))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_appendStringInfoString(m, v16+int32(-16), v162)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L57
	}
L56:
	;
	goto L52
L57:
	;
	goto L52
L58:
	;
	switch v105 - int32(99) {
	case 0:
		v227 = v214
		goto L30
	case 1:
		goto L74
	default:
		goto L73
	case 6:
		goto L75
	}
L59:
	;
	if v84 == int32(-1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v182 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v209 = F_strlen(m, v112)
	mBase = m.M
	v214 = v209 + v112 + int32(1)
	goto L58
L63:
	;
	v186 = int32(18)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	if v188 == v186 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v200 = int32(1)
	if v182&v200 != 0 {
		v214 = v112 + int32(base.Ui32(v182)>>(uint(v200)%32))
		goto L58
	} else {
		goto L72
	}
L66:
	;
	v191 = v186
	goto L68
L67:
	;
	v191 = int32(2)
	goto L68
L68:
	;
	if base.Ui32((v188-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v198 = int32(6)
	goto L71
L70:
	;
	v198 = v191
	goto L71
L71:
	;
	v214 = v112 + v198
	goto L58
L72:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v214 = v112 + int32(base.Ui32(v205)>>(uint(int32(2))%32))
	goto L58
L73:
	;
	v227 = (v214 + int32(1)) & int32(-2)
	goto L30
L74:
	;
	v227 = (v214 + int32(7)) & int32(-8)
	goto L30
L75:
	;
	v227 = (v214 + int32(3)) & int32(-4)
	goto L30
L76:
	;
	v239 = v234
	goto L78
L77:
	;
	v239 = v236
	goto L78
L78:
	;
	if v111 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v240 = v239
	goto L81
L80:
	;
	v240 = v115
	goto L81
L81:
	;
	if v111 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v243 = v111 + v238
	goto L84
L83:
	;
	v243 = int32(0)
	goto L84
L84:
	;
	v245 = v120 + int32(1)
	if v245 != v23 {
		v111 = v243
		v112 = v231
		v115 = v240
		v117 = v232
		v120 = v245
		goto L27
	} else {
		goto L85
	}
L85:
	;
	goto L28
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = v265 << (uint(int32(2)) % 32)
	if v263 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	base.MemoryCopy(m, v266+int32(4), v262, v263)
	goto L89
L88:
	;
	goto L89
L89:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_pfree(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v277 = v266
	goto L1
}
func F_array_to_text_null(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	v2 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == v2 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v11 != int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v19 = F_pg_detoast_datum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_pg_detoast_datum_packed(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
						if v28 == int32(1) {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
							if v34 == int32(18) {
								v37 = int32(16)
							} else {
								v37 = int32(0)
							}
							if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v44 = int32(4)
							} else {
								v44 = v37
							}
							v57 = v44
						} else {
							v45 = int32(1)
							if v28&v45 != 0 {
								v57 = int32(base.Ui32(v28)>>(uint(v45)%32)) - v45
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
								v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v60 = F_palloc(m, v57+int32(1))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v62 = int32(1)
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
								if v64&v62 != 0 {
									v67 = v62
								} else {
									v67 = int32(4)
								}
								base.MemoryCopy(m, v60, v26+v67, v57)
							} else {
							}
							v71 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v57+v60))) = uint8(v71)
							if v26 != v24 {
								F_pfree(m, v26)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
									if v76 != 0 {
										v133 = v2
										v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int32(0)
										} else {
											return v134
										}
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v78 = F_pg_detoast_datum_packed(m, v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v80 = F_pg_detoast_datum_packed(m, v78)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
												if v82 == int32(1) {
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
													if v88 == int32(18) {
														v91 = int32(16)
													} else {
														v91 = int32(0)
													}
													if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
														v98 = int32(4)
													} else {
														v98 = v91
													}
													v111 = v98
												} else {
													v99 = int32(1)
													if v82&v99 != 0 {
														v111 = int32(base.Ui32(v82)>>(uint(v99)%32)) - v99
													} else {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
														v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
													}
												}
												v114 = F_palloc(m, v111+int32(1))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													if v111 != 0 {
														v116 = int32(1)
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
														if v118&v116 != 0 {
															v121 = v116
														} else {
															v121 = int32(4)
														}
														base.MemoryCopy(m, v114, v80+v121, v111)
													} else {
													}
													v125 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v111+v114))) = uint8(v125)
													if v80 == v78 {
														v133 = v114
														v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															return v134
														}
													} else {
														F_pfree(m, v80)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															v133 = v114
															v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int32(0)
															} else {
																return v134
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
								if v76 != 0 {
									v133 = v2
									v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int32(0)
									} else {
										return v134
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v78 = F_pg_detoast_datum_packed(m, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v80 = F_pg_detoast_datum_packed(m, v78)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
											if v82 == int32(1) {
												v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
												if v88 == int32(18) {
													v91 = int32(16)
												} else {
													v91 = int32(0)
												}
												if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
													v98 = int32(4)
												} else {
													v98 = v91
												}
												v111 = v98
											} else {
												v99 = int32(1)
												if v82&v99 != 0 {
													v111 = int32(base.Ui32(v82)>>(uint(v99)%32)) - v99
												} else {
													v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
												}
											}
											v114 = F_palloc(m, v111+int32(1))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												if v111 != 0 {
													v116 = int32(1)
													v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
													if v118&v116 != 0 {
														v121 = v116
													} else {
														v121 = int32(4)
													}
													base.MemoryCopy(m, v114, v80+v121, v111)
												} else {
												}
												v125 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v111+v114))) = uint8(v125)
												if v80 == v78 {
													v133 = v114
													v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int32(0)
													} else {
														return v134
													}
												} else {
													F_pfree(m, v80)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														v133 = v114
														v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															return v134
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		}
	} else {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		return int32(0)
	}
}
func F_compute_array_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int64
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int64
	_ = v350
	var v357 int64
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int64
	_ = v409
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v454 int64
	_ = v454
	var v460 float64
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int64
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v584 int64
	_ = v584
	var v587 int64
	_ = v587
	var v591 int64
	_ = v591
	var v597 int64
	_ = v597
	var v599 int64
	_ = v599
	var v603 int64
	_ = v603
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v607 int64
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v634 int64
	_ = v634
	var v641 float64
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int64
	_ = v671
	var v672 int32
	_ = v672
	var v673 int64
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 float64
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v801 int32
	_ = v801
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v991 int32
	_ = v991
	var v993 int64
	_ = v993
	var v994 int32
	_ = v994
	var v995 int64
	_ = v995
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1022 int64
	_ = v1022
	var v1034 int32
	_ = v1034
	var v1041 int32
	_ = v1041
	var v1056 int64
	_ = v1056
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1071 int64
	_ = v1071
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1095 int64
	_ = v1095
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	v5 = int32(0)
	v22 = int64(0)
	v29 = m.G0
	v31 = v29 - int32(176)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	m.T0[v36].(func(*base.Module, int32, int32, int32, float64))(m, l0, l1, l2, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
	*(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[0])) = v33
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = int32(1245)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = int32(1246)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+124)) = int64(68719476740)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+148)) = v50
	v54 = v42 * int32(10)
	v58 = F_hash_create(m, int32(_a_F_compute_array_stats_0), v54, v31+int32(108), int32(1224))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+76)) = int64(34359738372)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v63
	v68 = base.I32_div_s(v42*int32(_a_F_compute_array_stats_1), int32(7))
	v74 = F_hash_create(m, int32(_a_F_compute_array_stats_2), int32(64), v31+int32(60), int32(1064))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l2 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v448 = v5
	v454 = v22
	v460 = float64(0)
	goto L7
L6:
	;
	v90 = v5
	v93 = int32(1)
	v97 = v5
	v101 = v5
	v103 = v22
	goto L9
L7:
	;
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	if v461 == int32(0) {
		v474 = v5
		goto L68
	} else {
		goto L69
	}
L8:
	;
	v448 = v403
	v454 = v409
	v460 = base.F64_convert_i32_s(v407)
	goto L7
L9:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L64
	}
L11:
	;
	goto L10
L12:
	;
	v114 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v90, v31+int32(59))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+59)))
	if v116 != 0 {
		v399 = v93
		v403 = v97
		v407 = v101
		v409 = v103
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v416 = v90 + int32(1)
	if l2 != v416 {
		v90 = v416
		v93 = v399
		v97 = v403
		v101 = v407
		v103 = v409
		goto L9
	} else {
		goto L63
	}
L15:
	;
	v117 = F_toast_raw_datum_size(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(_a_F_compute_array_stats_3)) < base.Ui32(v117) {
		v399 = v93
		v403 = v97
		v407 = v101
		v409 = v103
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v121 = F_pg_detoast_datum(m, v114)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+14)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)))
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+16)))
	F_deconstruct_array(m, v121, v124, v125, v126, v31+int32(48), v31+int32(44), v31+int32(52))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v135 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	if v135 < v137 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v145 = v135
	v152 = v93
	v159 = v135
	v161 = v103
	goto L23
L21:
	;
	v341 = v93
	v348 = v135
	v350 = v103
	goto L22
L22:
	;
	v357 = v350 - v103
	*(*uint32)(unsafe.Add(mBase, uint32(v31)+156)) = uint32(v357)
	v364 = F_hash_search(m, v74, v31+int32(156), int32(1), v31+int32(40))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L53
	}
L23:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v145))))
	if v170 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v341 = v309
	v348 = v316
	v350 = v318
	goto L22
L25:
	;
	v326 = v145 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	if v326 < v327 {
		v145 = v326
		v152 = v309
		v159 = v316
		v161 = v318
		goto L23
	} else {
		goto L52
	}
L26:
	;
	v309 = v152
	v316 = int32(1)
	v318 = v161
	goto L25
L27:
	;
	goto L28
L28:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v145<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v176
	v183 = F_hash_search(m, v58, v31+int32(40), int32(1), v31+int32(39))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+39)))
	if v185 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v209 = v161 + int64(1)
	v210 = base.I64_rem_s(v209, base.I64_extend_i32_s(v68))
	if v210 != int64(0) {
		v309 = v152
		v316 = v159
		v318 = v209
		goto L25
	} else {
		goto L36
	}
L31:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	if v188 == v90 {
		v309 = v152
		v316 = v159
		v318 = v161
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)))
	v197 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+14)))
	v198 = F_datumCopy(m, v195, v196, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v90
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v191 + int32(1)
	goto L30
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v90
	v201 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v152 - v201
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v198
	goto L30
L36:
	;
	v214 = v31 + int32(156)
	F_hash_seq_init(m, v214, v58)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v217 = F_hash_seq_search(m, v214)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v217 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v225 = v217
	goto L42
L40:
	;
	goto L41
L41:
	;
	v309 = v152 + int32(1)
	v316 = v159
	v318 = v209
	goto L25
L42:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v152 < v247+v248 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v265 = F_hash_seq_search(m, v31+int32(156))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L50
	}
L45:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v254 = F_hash_search(m, v58, v225, int32(2), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v254 == int32(0) {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[0]))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+12)))
	if v260 != 0 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v251)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	if v265 != 0 {
		v225 = v265
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	goto L24
L53:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+40)))
	if v366 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v373 = v369 + int32(1)
	goto L56
L55:
	;
	v373 = int32(1)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v373
	if v114 != v121 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v121)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	F_pfree(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	F_pfree(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v399 = v341
	v403 = v97 + int32(1)
	v407 = v348 + v101
	v409 = v350
	goto L14
L63:
	;
	goto L8
L64:
	;
	F_errmsg_internal(m, int32(_a_F_compute_array_stats_4), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_compute_array_stats_5), int32(695), int32(_a_F_compute_array_stats_6))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L154
	}
L68:
	;
	if v448 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	if v464 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v474 = int32(1)
	goto L68
L71:
	;
	goto L72
L72:
	;
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	if v468 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v474 = int32(2)
	goto L68
L74:
	;
	goto L75
L75:
	;
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)))
	if v472 != 0 {
		goto L67
	} else {
		goto L76
	}
L76:
	;
	v474 = int32(3)
	goto L68
L77:
	;
	m.G0 = v31 + int32(176)
	return
L78:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478)+412))
	if v480 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v546 = F_palloc(m, v543<<(uint(int32(2))%32))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v478)+376))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478)+364))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v478)+352))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v478)+340))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v478)+328))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v478)+316))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v478)+304))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v478)+292))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v478)+280))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v478)+268))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v478)+256))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v478)+244))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v478)+232))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v478)+220))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v478)+208))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v478)+196))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v478)+184))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v478)+172))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v478)+160))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v478)+148))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v478)+136))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v478)+124))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v478)+112))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v478)+100))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v478)+88))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v478)+76))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v478)+64))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v478)+52))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v478)+40))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v478)+28))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v478)+16))
	v543 = v481 + (v482 + (v483 + (v484 + (v485 + (v486 + (v487 + (v488 + (v489 + (v490 + (v491 + (v492 + (v493 + (v494 + (v495 + (v496 + (v497 + (v498 + (v499 + (v500 + (v501 + (v502 + (v503 + (v504 + (v505 + (v506 + (v507 + (v508 + (v509 + (v510 + (v511 + v479))))))))))))))))))))))))))))))
	goto L82
L81:
	;
	v543 = v479
	goto L82
L82:
	;
	goto L79
L83:
	;
	v549 = v31 + int32(156)
	F_hash_seq_init(m, v549, v58)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v555 = base.I64_div_s(v454*int64(9), base.I64_extend_i32_s(v68))
	v556 = F_hash_seq_search(m, v549)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v644 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L103
	}
L86:
	;
	if v556 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v618 = int32(0)
	v634 = v454
	v641 = float64(0)
	goto L85
L88:
	;
	goto L89
L89:
	;
	v568 = int32(0)
	v569 = v556
	v584 = v454
	v587 = v22
	goto L90
L90:
	;
	v591 = int64(*(*int32)(unsafe.Add(mBase, uint32(v569)+4)))
	if v555 < v591 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v618 = v604
	v634 = v605
	v641 = base.F64_convert_i64_u(v607)
	goto L85
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v546+v568<<(uint(int32(2))%32)))) = v569
	v597 = int64(*(*int32)(unsafe.Add(mBase, uint32(v569)+4)))
	if v597 < v587 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v604 = v568
	v605 = v584
	v607 = v587
	goto L94
L94:
	;
	v610 = F_hash_seq_search(m, v31+int32(156))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L101
	}
L95:
	;
	v599 = v587
	goto L97
L96:
	;
	v599 = v597
	goto L97
L97:
	;
	if v584 < v597 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v603 = v584
	goto L100
L99:
	;
	v603 = v597
	goto L100
L100:
	;
	v604 = v568 + int32(1)
	v605 = v603
	v607 = v599
	goto L94
L101:
	;
	if v610 != 0 {
		v568 = v604
		v569 = v610
		v584 = v605
		v587 = v607
		goto L90
	} else {
		goto L102
	}
L102:
	;
	goto L91
L103:
	;
	if v644 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v618
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v543
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v54
	F_errmsg_internal(m, int32(_a_F_compute_array_stats_7), v31)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v618 <= v54 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	F_errfinish(m, int32(_a_F_compute_array_stats_5), int32(494), int32(_a_F_compute_array_stats_8))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v675 = l0 + int32(52)
	if int32(0) < v672 {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v672 = v618
	v673 = v634
	goto L109
L111:
	;
	goto L112
L112:
	;
	F_qsort_interruptible(m, v546, v618, int32(4), int32(1247), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v546+v54<<(uint(int32(2))%32)-int32(4))))
	v671 = int64(*(*int32)(unsafe.Add(mBase, uint32(v670)+4)))
	v672 = v54
	v673 = v671
	goto L109
L114:
	;
	v678 = int32(0)
	F_qsort_interruptible(m, v546, v672, int32(4), int32(1248), v678)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	v801 = v474
	goto L116
L116:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)+4))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v820)+412))
	if v822 != 0 {
		goto L125
	} else {
		goto L126
	}
L117:
	;
	v684 = int32(_a_F_compute_array_stats_9)
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1]))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1])) = v687
	v689 = base.F64_convert_i32_u(v448)
	v692 = F_palloc(m, v672<<(uint(int32(2))%32))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v695 = v672 + int32(3)
	v698 = F_palloc(m, v695<<(uint(int32(2))%32))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v706 = v678
	goto L120
L120:
	;
	v729 = v706 << (uint(int32(2)) % 32)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v546+v729)))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)))
	v735 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+14)))
	v736 = F_datumCopy(m, v733, v734, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v748 = int32(2)
	v750 = v698 + v672<<(uint(v748)%32)
	*(*float32)(unsafe.Add(mBase, uint32(v750)+8)) = base.F32_demote_f64(base.F64_div(v460, v689))
	*(*float32)(unsafe.Add(mBase, uint32(v750)+4)) = base.F32_demote_f64(base.F64_div(v641, v689))
	*(*float32)(unsafe.Add(mBase, uint32(v750))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v673), v689))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1])) = v685
	v763 = int32(1)
	v764 = v474 << (uint(v763) % 32)
	v766 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v675+v764))) = uint16(v766)
	v770 = l0 + v474<<(uint(v748)%32)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+64)) = v771
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+124)) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v770)+84)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v770)+164)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v770)+104)) = v695
	*(*int32)(unsafe.Add(mBase, uint32(v770)+144)) = v672
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+184)) = v779
	v782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v764)+204)) = uint16(v782)
	v784 = l0 + v474
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+214)) = uint8(v785)
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+219)) = uint8(v787)
	v801 = v474 + v763
	goto L116
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692+v729))) = v736
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v698+v729))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v740), v689))
	v746 = v706 + int32(1)
	if v746 != v672 {
		v706 = v746
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	if v885 <= int32(0) {
		goto L77
	} else {
		goto L128
	}
L125:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v820)+376))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v820)+364))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v820)+352))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v820)+340))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v820)+328))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v820)+316))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v820)+304))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v820)+292))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v820)+280))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v820)+268))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v820)+256))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v820)+244))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v820)+232))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v820)+220))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v820)+208))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v820)+196))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v820)+184))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v820)+172))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v820)+160))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v820)+148))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v820)+136))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v820)+124))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v820)+112))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v820)+100))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v820)+88))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v820)+76))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v820)+64))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v820)+52))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v820)+40))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v820)+28))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v820)+16))
	v885 = v823 + (v824 + (v825 + (v826 + (v827 + (v828 + (v829 + (v830 + (v831 + (v832 + (v833 + (v834 + (v835 + (v836 + (v837 + (v838 + (v839 + (v840 + (v841 + (v842 + (v843 + (v844 + (v845 + (v846 + (v847 + (v848 + (v849 + (v850 + (v851 + (v852 + (v853 + v821))))))))))))))))))))))))))))))
	goto L127
L126:
	;
	v885 = v821
	goto L127
L127:
	;
	goto L124
L128:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v891 = F_palloc(m, v885<<(uint(int32(2))%32))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v894 = v31 + int32(156)
	F_hash_seq_init(m, v894, v74)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v897 = F_hash_seq_search(m, v894)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if v897 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v905 = v897
	v906 = int32(0)
	goto L135
L133:
	;
	goto L134
L134:
	;
	v966 = int32(2)
	if v888 <= v966 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891+v906<<(uint(int32(2))%32)))) = v905
	v936 = F_hash_seq_search(m, v31+int32(156))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	goto L134
L137:
	;
	if v936 != 0 {
		v905 = v936
		v906 = v906 + int32(1)
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v969 = v966
	goto L141
L140:
	;
	v969 = v888
	goto L141
L141:
	;
	v970 = int32(0)
	F_qsort_interruptible(m, v891, v885, int32(4), int32(1249), v970)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v978 = v969 + int32(1)
	v981 = F_MemoryContextAlloc(m, v976, v978<<(uint(int32(2))%32))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v981+v969<<(uint(int32(2))%32)))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v454), base.F64_convert_i32_u(v448)))
	v991 = int32(1)
	v993 = base.I64_extend_i32_u(v969 - v991)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v891)))
	v995 = int64(*(*int32)(unsafe.Add(mBase, uint32(v994)+4)))
	v1007 = v970
	v1009 = int32(0)
	v1022 = v993 * v995
	goto L144
L144:
	;
	if int64(0) < v1022 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v1115 = int32(5)
	*(*uint16)(unsafe.Add(mBase, uint32(v675+v801<<(uint(int32(1))%32)))) = uint16(v1115)
	v1119 = l0 + v801<<(uint(int32(2))%32)
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+64)) = v1120
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+124)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+84)) = v1122
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+104)) = v978
	goto L77
L146:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1079)))
	*(*float32)(unsafe.Add(mBase, uint32(v981+v1009<<(uint(int32(2))%32)))) = base.F32_convert_i32_s(v1105)
	v1110 = v1009 + int32(1)
	if v1110 != v969 {
		v1007 = v1080
		v1009 = v1110
		v1022 = v1095 - base.I64_extend_i32_u(v448-v991)
		goto L144
	} else {
		goto L153
	}
L147:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v891+v1007<<(uint(int32(2))%32))))
	v1079 = v1034
	v1080 = v1007
	v1095 = v1022
	goto L146
L148:
	;
	goto L149
L149:
	;
	v1041 = v1007
	v1056 = v1022
	goto L150
L150:
	;
	v1064 = v1041 + int32(1)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v891+v1064<<(uint(int32(2))%32))))
	v1069 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1068)+4)))
	v1071 = v1069*v993 + v1056
	if v1071 <= int64(0) {
		v1041 = v1064
		v1056 = v1071
		goto L150
	} else {
		goto L152
	}
L151:
	;
	v1079 = v1068
	v1080 = v1064
	v1095 = v1071
	goto L146
L152:
	;
	goto L151
L153:
	;
	goto L145
L154:
	;
	F_errmsg_internal(m, int32(_a_F_compute_array_stats_10), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_compute_array_stats_5), int32(440), int32(_a_F_compute_array_stats_8))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(99)
	v13 = int32(1)
	if l2 <= int32(699) {
		switch l2 - int32(18) {
		case 0:
			v60 = int32(1)
			v61 = v12
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 1:
			v60 = int32(64)
			v61 = v12
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 2:
			v60 = int32(8)
			v61 = int32(100)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 3:
			v60 = int32(2)
			v61 = int32(115)
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
				F_errmsg_internal(m, int32(_a_F_construct_array_builtin_0), v10)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_construct_array_builtin_1), int32(3463), int32(_a_F_construct_array_builtin_2))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 5, 8, 10:
			v60 = int32(4)
			v61 = int32(105)
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 7:
			v60 = int32(-1)
			v61 = int32(105)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 9:
			v60 = int32(6)
			v61 = int32(115)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		}
	} else {
		switch l2 - int32(700) {
		case 0:
			v60 = int32(4)
			v61 = int32(105)
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 1:
			v60 = int32(8)
			v61 = int32(100)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		default:
			if l2 == int32(2206) {
				v60 = int32(4)
				v61 = int32(105)
				v62 = v13
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
				v64 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
				v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v72
				}
			} else {
				if l2 != int32(2275) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(_a_F_construct_array_builtin_0), v10)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_construct_array_builtin_1), int32(3463), int32(_a_F_construct_array_builtin_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v60 = int32(-2)
					v61 = v12
					v62 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
					v64 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
					v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v72
					}
				}
			}
		}
	}
}
func F_get_array_type(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13922(m, l0, int32(82))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_parse_array(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_check_stack_depth(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v79
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = m.T0[v6].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v15 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v19 != int32(5) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v12 != 0 {
		v79 = v12
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v22 = int32(11)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v31 = F_json_lex(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	v26 = int32(6)
	goto L14
L13:
	;
	v26 = v22
	goto L14
L14:
	;
	if v19 == int32(12) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v29 = v22
	goto L17
L16:
	;
	v29 = v26
	goto L17
L17:
	;
	return v29
L18:
	;
	if v31 != 0 {
		v79 = v31
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v33 == int32(6) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v66 = F_json_lex(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v36 = F_parse_array_element(m, l0, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v36 != 0 {
		v79 = v36
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L24
L24:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v42 != int32(7) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v79 = v58
	goto L3
L26:
	;
	if v42 == int32(6) {
		goto L20
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v56 = F_json_lex(m, l0)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L36
	}
L29:
	;
	v47 = int32(11)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v51 = int32(7)
	goto L32
L31:
	;
	v51 = v47
	goto L32
L32:
	;
	if v42 == int32(12) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v54 = v47
	goto L35
L34:
	;
	v54 = v51
	goto L35
L35:
	;
	return v54
L36:
	;
	if v56 != 0 {
		v79 = v56
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v58 = F_parse_array_element(m, l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v58 == int32(0) {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	goto L25
L40:
	;
	if v66 != 0 {
		v79 = v66
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v68 - int32(1)
	if v5 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = m.T0[v5].(func(*base.Module, int32) int32)(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v79 = int32(0)
	goto L3
L45:
	;
	if v73 != 0 {
		v79 = v73
		goto L3
	} else {
		goto L46
	}
L46:
	;
	goto L44
}
