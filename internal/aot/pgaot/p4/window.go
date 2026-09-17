package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_window_frame_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	if l0&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if l0&int32(2) != 0 {
		v21 = int32(_a_F_get_window_frame_options_0)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v25 = l0 & int32(16)
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_appendStringInfoString(m, v9, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if l0&int32(4) != 0 {
		v21 = int32(_a_F_get_window_frame_options_18)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if l0&int32(8) == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v21 = int32(_a_F_get_window_frame_options_19)
	goto L5
L9:
	;
	return
L10:
	;
	goto L4
L11:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_get_window_frame_options_1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l0&int32(32) != 0 {
		v50 = int32(_a_F_get_window_frame_options_2)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	if v25 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	F_appendStringInfoString(m, v9, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L23
	}
L17:
	;
	if l0&int32(512) != 0 {
		v50 = int32(_a_F_get_window_frame_options_11)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if l0&int32(_a_F_get_window_frame_options_17) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_get_rule_expr(m, l1, l3, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if l0&int32(2048) != 0 {
		v50 = int32(_a_F_get_window_frame_options_14)
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if l0&int32(_a_F_get_window_frame_options_20) == int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v50 = int32(_a_F_get_window_frame_options_16)
	goto L16
L23:
	;
	goto L15
L24:
	;
	if l0&int32(_a_F_get_window_frame_options_3) != 0 {
		v93 = int32(_a_F_get_window_frame_options_4)
		goto L36
	} else {
		goto L37
	}
L25:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_get_window_frame_options_9))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if l0&int32(256) != 0 {
		v79 = int32(_a_F_get_window_frame_options_10)
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_appendStringInfoString(m, v9, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L34
	}
L28:
	;
	if l0&int32(1024) != 0 {
		v79 = int32(_a_F_get_window_frame_options_11)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if l0&int32(_a_F_get_window_frame_options_12) == int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_get_rule_expr(m, l2, l3, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	if l0&int32(_a_F_get_window_frame_options_13) != 0 {
		v79 = int32(_a_F_get_window_frame_options_14)
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if l0&int32(_a_F_get_window_frame_options_15) == int32(0) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v79 = int32(_a_F_get_window_frame_options_16)
	goto L27
L34:
	;
	goto L24
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v98 = v96 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100+v98))) = uint8(v102)
	goto L3
L36:
	;
	F_appendStringInfoString(m, v9, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L40
	}
L37:
	;
	if l0&int32(_a_F_get_window_frame_options_5) != 0 {
		v93 = int32(_a_F_get_window_frame_options_6)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if l0&int32(_a_F_get_window_frame_options_7) == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v93 = int32(_a_F_get_window_frame_options_8)
	goto L36
L40:
	;
	goto L35
}
func F_transformWindowDefinitions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	if l1 == v4 {
		v489 = v4
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L53
	} else {
		goto L137
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L53
	} else {
		goto L134
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L53
	} else {
		goto L129
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L53
	} else {
		goto L124
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L53
	} else {
		goto L119
	}
L6:
	;
	m.G0 = v16 + int32(112)
	return v489
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 <= int32(0) {
		v489 = v4
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = v4
	v34 = v4
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34<<(uint(int32(2))%32))))
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v45 == v41 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v489 = v477
	goto L6
L11:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v263 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L12:
	;
	v260 = int32(0)
	goto L11
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L53
	} else {
		goto L59
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L53
	} else {
		goto L54
	}
L15:
	;
	if v31 == int32(0) {
		goto L14
	} else {
		goto L37
	}
L16:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v118 == int32(0) {
		goto L12
	} else {
		goto L36
	}
L17:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v117 != 0 {
		goto L15
	} else {
		goto L35
	}
L18:
	;
	if v31 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v50 <= int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v61 = int32(0)
	goto L21
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53+v61<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L17
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if base.B2i32(v75 == int32(0))|base.B2i32(v75 != v78) != 0 {
		v96 = v75
		v97 = v78
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v102 = v61 + int32(1)
	if v50 != v102 {
		v61 = v102
		goto L21
	} else {
		goto L34
	}
L26:
	;
	if v96-v97 == int32(0) {
		goto L13
	} else {
		goto L33
	}
L27:
	;
	goto L26
L28:
	;
	v81 = v72
	v82 = v45
	goto L29
L29:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v86 == int32(0) {
		v96 = v86
		v97 = v85
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v96 = v86
	v97 = v85
	goto L27
L31:
	;
	v89 = int32(1)
	if v86 == v85 {
		v81 = v81 + v89
		v82 = v82 + v89
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L25
L34:
	;
	goto L22
L35:
	;
	goto L12
L36:
	;
	goto L14
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v123 <= int32(0) {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v134 = int32(0)
	goto L39
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v126+v134<<(uint(int32(2))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v145 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L14
L41:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.B2i32(v148 == int32(0))|base.B2i32(v148 != v151) != 0 {
		v169 = v148
		v170 = v151
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v175 = v134 + int32(1)
	if v123 != v175 {
		v134 = v175
		goto L39
	} else {
		goto L52
	}
L44:
	;
	if v169-v170 == int32(0) {
		v260 = v144
		goto L11
	} else {
		goto L51
	}
L45:
	;
	goto L44
L46:
	;
	v154 = v145
	v155 = v117
	goto L47
L47:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v159 == int32(0) {
		v169 = v159
		v170 = v158
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v169 = v159
	v170 = v158
	goto L45
L49:
	;
	v162 = int32(1)
	if v159 == v158 {
		v154 = v154 + v162
		v155 = v155 + v162
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L43
L52:
	;
	goto L40
L53:
	;
	return int32(0)
L54:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v199
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_0), v16+int32(80))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2806), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v221
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_4), v16+int32(96))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L53
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2793), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v318 = F_transformGroupClause(m, l0, v314, int32(0), l2, v306, int32(9), int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L53
	} else {
		goto L74
	}
L65:
	;
	v306 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v267 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v269 <= v267 {
		v306 = v267
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v277 = v267
	v278 = v267
	goto L69
L69:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v278<<(uint(int32(2))%32))))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v292 = F_findTargetlistEntrySQL99(m, l0, v290, l2, int32(10))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L53
	} else {
		goto L71
	}
L70:
	;
	v306 = v295
	goto L64
L71:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v295 = F_addTargetToSortList(m, l0, v292, v277, v294, v289)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L53
	} else {
		goto L72
	}
L72:
	;
	v298 = v278 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v298 < v299 {
		v277 = v295
		v278 = v298
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v321 = F_palloc0(m, int32(56))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L53
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = int32(108)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+4)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v327
	if v260 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+20)) = v406
	v410 = int32(0)
	if base.B2i32(v406&int32(2) == v410)|base.B2i32(v406&int32(_a_F_transformWindowDefinitions_5) == v410) == v410 {
		goto L102
	} else {
		goto L103
	}
L77:
	;
	if v318 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v321)+52)) = uint8(v400)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+16)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = v318
	v404 = v306
	goto L76
L80:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v330 = F_copyObjectImpl(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L53
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = v330
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if v306 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v321)+52)) = uint8(v362)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+16)) = v363
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	if v366 == int32(1058) {
		v404 = v363
		goto L76
	} else {
		goto L93
	}
L83:
	;
	v334 = int32(0)
	if v333 == v334 {
		v362 = v334
		v363 = v306
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v360 = F_copyObjectImpl(m, v333)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L53
	} else {
		goto L92
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L53
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L53
	} else {
		goto L88
	}
L88:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v344
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_6), v16+int32(48))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L53
	} else {
		goto L89
	}
L89:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L53
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2867), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L53
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v362 = int32(1)
	v363 = v360
	goto L82
L93:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v306|v369 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v371 != int32(1058) {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L53
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L53
	} else {
		goto L97
	}
L97:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v381
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_7), v16+int32(32))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L53
	} else {
		goto L98
	}
L98:
	;
	F_errhint(m, int32(_a_F_transformWindowDefinitions_8), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L53
	} else {
		goto L99
	}
L99:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L53
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2904), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L53
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
	if v404 == int32(0) {
		goto L3
	} else {
		goto L105
	}
L103:
	;
	v450 = v406
	goto L104
L104:
	;
	if v450&int32(8) != 0 {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if v421 != int32(1) {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v427 = F_get_sortgroupclause_expr(m, v425, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L53
	} else {
		goto L107
	}
L107:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	v436 = F_get_ordering_op_properties(m, v429, v16+int32(108), v16+int32(104), v16+int32(100))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L53
	} else {
		goto L108
	}
L108:
	;
	if v436 == int32(0) {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	v440 = F_exprCollation(m, v427)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L53
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+40)) = v440
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+16)))
	v445 = v443 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v321)+44)) = uint8(v445)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v321)+45)) = uint8(v447)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v321)+20))
	v450 = v449
	goto L104
L111:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v321)+16))
	if v454 == int32(0) {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	v462 = F_transformFrameOffset(m, l0, v450, v457, v458, v321+int32(32), v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L53
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+24)) = v462
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v321)+20))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	v471 = F_transformFrameOffset(m, l0, v465, v466, v467, v321+int32(36), v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L53
	} else {
		goto L116
	}
L116:
	;
	v474 = v34 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+48)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v321)+28)) = v471
	v477 = F_lappend(m, v31, v321)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L53
	} else {
		goto L117
	}
L117:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v474 < v479 {
		v31 = v477
		v34 = v474
		goto L9
	} else {
		goto L118
	}
L118:
	;
	goto L10
L119:
	;
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L53
	} else {
		goto L120
	}
L120:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v505
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_9), v16-int32(-64))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L53
	} else {
		goto L121
	}
L121:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L53
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2855), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L53
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
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L53
	} else {
		goto L125
	}
L125:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v527
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_7), v16+int32(16))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L53
	} else {
		goto L126
	}
L126:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L53
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2897), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L53
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
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L53
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_10), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L53
	} else {
		goto L131
	}
L131:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L53
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2924), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L53
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
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v565
	F_errmsg_internal(m, int32(_a_F_transformWindowDefinitions_11), v16)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L53
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2933), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L53
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(_a_F_transformWindowDefinitions_3))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L53
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(_a_F_transformWindowDefinitions_12), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L53
	} else {
		goto L139
	}
L139:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	F_parser_errposition(m, l0, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L53
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_transformWindowDefinitions_1), int32(2947), int32(_a_F_transformWindowDefinitions_2))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L53
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_window_cume_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_WinGetPartitionRowCount(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+176))
	goto L3
L3:
	;
	v15 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_WinSetMarkPosition(m, v7, v13)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	if v17 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(1)
	v29 = int32(0)
	goto L4
L7:
	;
	goto L8
L8:
	;
	v25 = F_WinRowsArePeers(m, v7, v13-int64(1), v13)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v29 = v25 ^ int32(1)
	goto L4
L10:
	;
	v33 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v29 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v75 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v68), base.F64_convert_i64_s(v8)))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	if v37 != int64(1) {
		v68 = v37
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)+176))
	goto L17
L16:
	;
	goto L15
L17:
	;
	v44 = v42 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v44
	if v8 <= v44 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v44
	goto L12
L19:
	;
	goto L20
L20:
	;
	v48 = v44
	goto L21
L21:
	;
	v55 = F_WinRowsArePeers(m, v7, v48-int64(1), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v68 = v61
	goto L12
L23:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	if v55 == int32(0) {
		v68 = v57
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v60 = int64(1)
	v61 = v57 + v60
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v61
	v64 = v48 + v60
	if v64 != v8 {
		v48 = v64
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	return v75
}
func F_window_lag_with_offset(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v8 + int32(15)
	v14 = F_WinGetFuncArgCurrent(m, v10, int32(1), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v18 == int32(0) {
			v21 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v23 == v21 {
				v71 = v21
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
				if v29 == int32(0) {
					v71 = v21
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v34 = v32 - int32(11)
					if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v34))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v34)%32))&int32(1) == int32(0))|int32(0) != 0 {
						v71 = v21
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F_window_lag_with_offset[0])))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v29+v49)))
						if v51 == int32(0) {
							v71 = v21
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
							if v54 <= int32(1) {
								v71 = v21
							} else {
								v56 = int32(1)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(4))))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								switch v62 - int32(7) {
								case 0:
									v71 = v56
								case 1:
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
									if v65 == int32(0) {
										v71 = v56
									} else {
										v71 = int32(0)
									}
								default:
									v71 = int32(0)
								}
							}
						}
					}
				}
			}
			v74 = F_WinGetFuncArgInPartition(m, v10, v21-v14, v71, v13, v8+int32(14))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v76 != int32(1) {
					v83 = v74
				} else {
					v80 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v80)
					v83 = int32(0)
				}
				m.G0 = v8 + int32(16)
				return v83
			}
		} else {
			v80 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v80)
			v83 = int32(0)
			m.G0 = v8 + int32(16)
			return v83
		}
	}
}
