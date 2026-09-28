package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 == int32(2275) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L28
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L24
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 != int32(1) {
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
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L20
	}
L6:
	;
	v15 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	if v15 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_deconstruct_array_builtin(m, l0, int32(2275), v7+int32(12), int32(0), l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = F_palloc(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v30 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_pfree(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L19
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v36<<(uint(int32(3))%32))))
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
	*(*int32)(unsafe.Add(mBase, uint32(v28+v36<<(uint(int32(2))%32)))) = v45
	v49 = v36 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v49 < v50 {
		v36 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	m.G0 = v7 + int32(16)
	return v28
L20:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_ArrayGetIntegerTypmods_0), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ArrayGetIntegerTypmods_1), int32(242), int32(_a_F_ArrayGetIntegerTypmods_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
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
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_ArrayGetIntegerTypmods_3), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ArrayGetIntegerTypmods_1), int32(247), int32(_a_F_ArrayGetIntegerTypmods_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_ArrayGetIntegerTypmods_4), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ArrayGetIntegerTypmods_1), int32(252), int32(_a_F_ArrayGetIntegerTypmods_2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
func F_array_agg_array_combine(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
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
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
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
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = v14 + int32(12)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == v2 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L32
	} else {
		goto L161
	}
L2:
	;
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v47 = v44
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v39
	v44 = v40
	goto L3
L5:
	;
	v36 = int32(0)
	if v17 == v36 {
		v44 = v36
		goto L3
	} else {
		goto L15
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	switch v22 - int32(435) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L5
	}
L7:
	;
	if v17 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = int32(1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v39 = v29
	v40 = int32(1)
	goto L4
L12:
	;
	v47 = int32(2)
	goto L2
L13:
	;
	goto L14
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+376))
	v39 = v34
	v40 = int32(2)
	goto L4
L15:
	;
	v39 = v36
	v40 = v2
	goto L4
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v48 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L32
	} else {
		goto L158
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v52 = v51
	goto L21
L20:
	;
	v52 = v2
	goto L21
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v53 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_u(v574)
L23:
	;
	if v52 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v56 != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v52 != 0 {
		v574 = v52
		goto L22
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
	v574 = int32(0)
	goto L22
L29:
	;
	v63 = int32(_a_F_array_agg_array_combine_0)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v56)+80))
	v69 = int32(0)
	v71 = F_initArrayResultArr(m, v68, v69, v66, v69)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	if v123 <= int32(0) {
		v574 = v52
		goto L22
	} else {
		goto L45
	}
L32:
	;
	return int64(0)
L33:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v75
	v77 = F_palloc(m, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v80 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v85 = base.I32_div_s(v81+int32(7), int32(8))
	v86 = F_palloc(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v92 = v77
	goto L37
L37:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v94 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v86
	if v85 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	base.MemoryCopy(m, v86, v89, v85)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v92 = v91
	goto L37
L42:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	base.MemoryCopy(m, v92, v95, v94)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+28)) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+48)) = v105
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v56)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+40)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v56)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+32)) = v109
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v56)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+56)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v56)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+64)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v56)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+72)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v56)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+80)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v56)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+84)) = v119
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v64
	v574 = v71
	goto L22
L45:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	if v126 == v127 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v213 = v212 + v209
	if base.B2i32(v209 < int32(0))^base.B2i32(v213 < v212) != 0 {
		goto L1
	} else {
		goto L66
	}
L47:
	;
	v156 = int32(1)
	goto L56
L48:
	;
	if v126 < int32(2) {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L32
	} else {
		goto L52
	}
L51:
	;
	v131 = int32(56)
	v135 = int32(32)
	goto L47
L52:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_array_agg_array_combine_1), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(1053), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L32
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v168 = v156 << (uint(int32(2)) % 32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v52+v135+v168)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+(v56+v135))))
	if v170 != v172 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L32
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v168+(v52+v131))))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168+(v56+v131))))
	if v175 != v177 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v180 = v156 + int32(1)
	if v126 != v180 {
		v156 = v180
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L46
L62:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_array_agg_array_combine_1), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L32
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(1061), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L32
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v219 = v123 + v218
	if base.B2i32(v123 < int32(0)) != base.B2i32(v219 < v218) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v222 = int32(_a_F_array_agg_array_combine_0)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v227 < v213 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v229 = int32(1)
	if v213&(v213-v229) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v244 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	v237 = v229 << (uint(int32(32)-base.I32_clz(v213)) % 32)
	goto L73
L72:
	;
	v237 = v213
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v240 = F_repalloc(m, v239, v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L32
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v240
	goto L70
L75:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v553 != 0 {
		goto L155
	} else {
		goto L156
	}
L76:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v424 = int32(0)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	if v425 <= v424 {
		goto L125
	} else {
		goto L126
	}
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v247 == int32(0) {
		goto L75
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v219 <= v401 {
		v420 = v244
		goto L76
	} else {
		goto L119
	}
L80:
	;
	v252 = int32(256)
	if v219 <= v252 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v255 = v252
	goto L83
L82:
	;
	v255 = v219
	goto L83
L83:
	;
	if v255&(v255-int32(1)) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v262 = int32(1) << (uint(int32(32)-base.I32_clz(v255)) % 32)
	goto L86
L85:
	;
	v262 = v255
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v262
	v267 = base.I32_div_s(v262+int32(7), int32(8))
	v268 = F_palloc(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L32
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v268
	v271 = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	if v274 <= v271 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v420 = v400
	goto L76
L89:
	;
	goto L88
L90:
	;
	v286 = base.I32_div_s(v271, int32(8))
	v287 = v268 + v286
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	goto L92
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v315)
	goto L89
L92:
	;
	v291 = v287
	v292 = v288
	v295 = v274
	v296 = int32(1)
	goto L95
L95:
	;
	v300 = v292 | v296
	v301 = int32(1)
	v302 = v295 - v301
	v304 = v296 << (uint(v301) % 32)
	if v304 == int32(256) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v316 != int32(1) {
		goto L91
	} else {
		goto L102
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v291))) = uint8(v300)
	if v302 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L98:
	;
	v314 = v291
	v315 = v300
	v316 = v304
	goto L99
L99:
	;
	if base.Ui32(int32(1)) < base.Ui32(v295) {
		v291 = v314
		v292 = v315
		v295 = v302
		v296 = v316
		goto L95
	} else {
		goto L101
	}
L100:
	;
	v310 = int32(1)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	v314 = v291 + v310
	v315 = v311
	v316 = v310
	goto L99
L101:
	;
	goto L96
L102:
	;
	goto L89
L119:
	;
	v403 = int32(1)
	if v219&(v219-v403) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v411 = v403 << (uint(int32(32)-base.I32_clz(v219)) % 32)
	goto L122
L121:
	;
	v411 = v219
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v411
	v416 = base.I32_div_s(v411+int32(7), int32(8))
	v417 = F_repalloc(m, v244, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L32
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v417
	v420 = v417
	goto L76
L124:
	;
	goto L75
L125:
	;
	goto L124
L126:
	;
	v435 = int32(1) << (uint(v422&int32(7)) % 32)
	v437 = base.I32_div_s(v422, int32(8))
	v438 = v420 + v437
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v423 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v532))) = uint8(v533)
	goto L125
L128:
	;
	v442 = v438
	v443 = v439
	v446 = v425
	v447 = v435
	goto L131
L129:
	;
	goto L130
L130:
	;
	v477 = base.I32_div_s(v424, int32(8))
	v478 = v423 + v477
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v480 = v438
	v481 = v439
	v483 = v478
	v484 = v425
	v485 = v435
	v486 = int32(1)
	v487 = v479
	goto L139
L131:
	;
	v451 = v443 | v447
	v452 = int32(1)
	v453 = v446 - v452
	v455 = v447 << (uint(v452) % 32)
	if v455 == int32(256) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v467 != int32(1) {
		v532 = v465
		v533 = v466
		goto L127
	} else {
		goto L138
	}
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v451)
	if v453 == int32(0) {
		goto L125
	} else {
		goto L136
	}
L134:
	;
	v465 = v442
	v466 = v451
	v467 = v455
	goto L135
L135:
	;
	if base.Ui32(int32(1)) < base.Ui32(v446) {
		v442 = v465
		v443 = v466
		v446 = v453
		v447 = v467
		goto L131
	} else {
		goto L137
	}
L136:
	;
	v461 = int32(1)
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	v465 = v442 + v461
	v466 = v462
	v467 = v461
	goto L135
L137:
	;
	goto L132
L138:
	;
	goto L125
L139:
	;
	if v486&v487 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v510 == int32(1) {
		goto L125
	} else {
		goto L154
	}
L141:
	;
	v494 = v481 | v485
	goto L143
L142:
	;
	v494 = v481 & (v485 ^ int32(-1))
	goto L143
L143:
	;
	v495 = int32(1)
	v496 = v484 - v495
	v498 = v485 << (uint(v495) % 32)
	if v498 == int32(256) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v480))) = uint8(v494)
	if v496 == int32(0) {
		goto L125
	} else {
		goto L147
	}
L145:
	;
	v508 = v480
	v509 = v494
	v510 = v498
	goto L146
L146:
	;
	v512 = v486 << (uint(int32(1)) % 32)
	if v512 == int32(256) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v504 = int32(1)
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+1)))
	v508 = v480 + v504
	v509 = v505
	v510 = v504
	goto L146
L148:
	;
	goto L140
L149:
	;
	if v496 == int32(0) {
		goto L148
	} else {
		goto L152
	}
L150:
	;
	v521 = v483
	v522 = v512
	v523 = v487
	goto L151
L151:
	;
	if base.Ui32(int32(1)) < base.Ui32(v484) {
		v480 = v508
		v481 = v509
		v483 = v521
		v484 = v496
		v485 = v510
		v486 = v522
		v487 = v523
		goto L139
	} else {
		goto L153
	}
L152:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
	v518 = int32(1)
	v521 = v483 + v518
	v522 = v518
	v523 = v517
	goto L151
L153:
	;
	goto L148
L154:
	;
	v532 = v508
	v533 = v509
	goto L127
L155:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	base.MemoryCopy(m, v554+v555, v557, v553)
	goto L157
L156:
	;
	goto L157
L157:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v559 + v560
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v563 + v564
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+32)) = v567 + v568
	*(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_combine[0])) = v223
	v574 = v52
	goto L22
L158:
	;
	F_errmsg_internal(m, int32(_a_F_array_agg_array_combine_4), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L32
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(986), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L32
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L32
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_agg_array_combine_5), v14)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L32
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_array_agg_array_combine_2), int32(1074), int32(_a_F_array_agg_array_combine_3))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L32
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_array_deserialize(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(1)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		v19 = v17 & v15
		if v19 != 0 {
			v20 = v15
		} else {
			v20 = int32(4)
		}
		if v17 == int32(1) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
			if v27 == int32(18) {
				v30 = int32(16)
			} else {
				v30 = int32(0)
			}
			if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v37 = int32(4)
			} else {
				v37 = v30
			}
			v48 = v37
		} else {
			v38 = int32(1)
			if v19 != 0 {
				v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v48
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + v20
		v54 = F_pq_getmsgint(m, v8, int32(4))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int64(0)
		} else {
			v57 = F_pq_getmsgint(m, v8, int32(4))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v60 = F_pq_getmsgint(m, v8, int32(4))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_deserialize[0]))
					v66 = F_initArrayResultArr(m, v57, v54, v64, int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						v71 = int32(1024)
						for {
							if v71 < v60 {
								v71 = v71 << (uint(int32(1)) % 32)
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v71
						v77 = F_palloc(m, v71)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v77
							v80 = F_pq_getmsgbytes(m, v8, v60)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								if v60 != 0 {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
									base.MemoryCopy(m, v82, v80, v60)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v60
								v86 = F_pq_getmsgint(m, v8, int32(4))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v86
									v90 = F_pq_getmsgint(m, v8, int32(4))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v90
										if int32(0) < v90 {
											v98 = base.I32_div_s(v90+int32(7), int32(8))
											v99 = F_palloc(m, v98)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v99
												v102 = F_pq_getmsgbytes(m, v8, v98)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int64(0)
												} else {
													if v98 == int32(0) {
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
														base.MemoryCopy(m, v106, v102, v98)
													}
													v113 = F_pq_getmsgint(m, v8, int32(4))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v113
														v117 = F_pq_getmsgint(m, v8, int32(4))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v117
															v121 = F_pq_getmsgbytes(m, v8, int32(24))
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int64(0)
															} else {
																v123 = *(*int64)(unsafe.Add(mBase, uint32(v121)+16))
																*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v123
																v125 = *(*int64)(unsafe.Add(mBase, uint32(v121)+8))
																*(*int64)(unsafe.Add(mBase, uint32(v66)+40)) = v125
																v127 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
																*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = v127
																v130 = F_pq_getmsgbytes(m, v8, int32(24))
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int64(0)
																} else {
																	v132 = *(*int64)(unsafe.Add(mBase, uint32(v130)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v66)+72)) = v132
																	v134 = *(*int64)(unsafe.Add(mBase, uint32(v130)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v66)+64)) = v134
																	v136 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
																	*(*int64)(unsafe.Add(mBase, uint32(v66)+56)) = v136
																	F_pq_getmsgend(m, v8)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v8 + int32(16)
																		return base.I64_extend_i32_u(v66)
																	}
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = int32(0)
											v113 = F_pq_getmsgint(m, v8, int32(4))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v113
												v117 = F_pq_getmsgint(m, v8, int32(4))
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v117
													v121 = F_pq_getmsgbytes(m, v8, int32(24))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int64(0)
													} else {
														v123 = *(*int64)(unsafe.Add(mBase, uint32(v121)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v123
														v125 = *(*int64)(unsafe.Add(mBase, uint32(v121)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v66)+40)) = v125
														v127 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
														*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = v127
														v130 = F_pq_getmsgbytes(m, v8, int32(24))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int64(0)
														} else {
															v132 = *(*int64)(unsafe.Add(mBase, uint32(v130)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v66)+72)) = v132
															v134 = *(*int64)(unsafe.Add(mBase, uint32(v130)+8))
															*(*int64)(unsafe.Add(mBase, uint32(v66)+64)) = v134
															v136 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
															*(*int64)(unsafe.Add(mBase, uint32(v66)+56)) = v136
															F_pq_getmsgend(m, v8)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return int64(0)
															} else {
																m.G0 = v8 + int32(16)
																return base.I64_extend_i32_u(v66)
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_palloc0(m, int32(52))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v132 = m.ExcPending
			if v132 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_array_create_iterator_0), int32(0))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_array_create_iterator_1), int32(_a_F_array_create_iterator_2), int32(_a_F_array_create_iterator_3))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 < l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_array_create_iterator_0), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_create_iterator_1), int32(_a_F_array_create_iterator_2), int32(_a_F_array_create_iterator_3))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v23 != 0 {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = l0 + v24<<(uint(int32(3))%32) + int32(16)
				} else {
					v31 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = l0 + int32(16)
				v36 = F_ArrayGetNItemsSafe(m, v33, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v36
					if l2 != 0 {
						v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v39)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
						*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v41)
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
						*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v43)
						v55 = v43
						switch v55 - int32(99) {
						case 0:
							v76 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
							*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
							if l1 != 0 {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v80 = int32(2)
								v84 = l1 << (uint(v80) % 32)
								v85 = v35 + v79<<(uint(v80)%32) - v84
								*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v89 = v87 << (uint(v80) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
								v94 = F_ArrayGetNItemsSafe(m, l1, v85)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
									v99 = F_palloc(m, v94<<(uint(int32(3))%32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
										v103 = F_palloc(m, v102)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v110 == int32(0) {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v120 = v110
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
											m.G0 = v11 + int32(16)
											return v14
										}
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v110 == int32(0) {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v120 = v110
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
								m.G0 = v11 + int32(16)
								return v14
							}
						case 1:
							v76 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
							*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
							if l1 != 0 {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v80 = int32(2)
								v84 = l1 << (uint(v80) % 32)
								v85 = v35 + v79<<(uint(v80)%32) - v84
								*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v89 = v87 << (uint(v80) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
								v94 = F_ArrayGetNItemsSafe(m, l1, v85)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
									v99 = F_palloc(m, v94<<(uint(int32(3))%32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
										v103 = F_palloc(m, v102)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v110 == int32(0) {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v120 = v110
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
											m.G0 = v11 + int32(16)
											return v14
										}
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v110 == int32(0) {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v120 = v110
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
								m.G0 = v11 + int32(16)
								return v14
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = base.I32_extend8_s(v55)
								F_errmsg_internal(m, int32(_a_F_array_create_iterator_4), v11)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_array_create_iterator_5), int32(322), int32(_a_F_array_create_iterator_6))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							v76 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
							*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
							if l1 != 0 {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v80 = int32(2)
								v84 = l1 << (uint(v80) % 32)
								v85 = v35 + v79<<(uint(v80)%32) - v84
								*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v89 = v87 << (uint(v80) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
								v94 = F_ArrayGetNItemsSafe(m, l1, v85)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
									v99 = F_palloc(m, v94<<(uint(int32(3))%32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
										v103 = F_palloc(m, v102)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v110 == int32(0) {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v120 = v110
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
											m.G0 = v11 + int32(16)
											return v14
										}
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v110 == int32(0) {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v120 = v110
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
								m.G0 = v11 + int32(16)
								return v14
							}
						case 16:
							v76 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
							*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
							if l1 != 0 {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v80 = int32(2)
								v84 = l1 << (uint(v80) % 32)
								v85 = v35 + v79<<(uint(v80)%32) - v84
								*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v89 = v87 << (uint(v80) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
								v94 = F_ArrayGetNItemsSafe(m, l1, v85)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
									v99 = F_palloc(m, v94<<(uint(int32(3))%32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
										v103 = F_palloc(m, v102)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v110 == int32(0) {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v120 = v110
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
											m.G0 = v11 + int32(16)
											return v14
										}
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v110 == int32(0) {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v120 = v110
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
								m.G0 = v11 + int32(16)
								return v14
							}
						}
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_get_typlenbyvalalign(m, v45, v14+int32(12), v14+int32(14), v14+int32(15))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
							v55 = v54
							switch v55 - int32(99) {
							case 0:
								v76 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
								if l1 != 0 {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v80 = int32(2)
									v84 = l1 << (uint(v80) % 32)
									v85 = v35 + v79<<(uint(v80)%32) - v84
									*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v89 = v87 << (uint(v80) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
									v94 = F_ArrayGetNItemsSafe(m, l1, v85)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
										v99 = F_palloc(m, v94<<(uint(int32(3))%32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
											v103 = F_palloc(m, v102)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
												v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v110 == int32(0) {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v120 = v110
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
												m.G0 = v11 + int32(16)
												return v14
											}
										}
									}
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v110 == int32(0) {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v120 = v110
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
									m.G0 = v11 + int32(16)
									return v14
								}
							case 1:
								v76 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
								if l1 != 0 {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v80 = int32(2)
									v84 = l1 << (uint(v80) % 32)
									v85 = v35 + v79<<(uint(v80)%32) - v84
									*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v89 = v87 << (uint(v80) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
									v94 = F_ArrayGetNItemsSafe(m, l1, v85)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
										v99 = F_palloc(m, v94<<(uint(int32(3))%32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
											v103 = F_palloc(m, v102)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
												v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v110 == int32(0) {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v120 = v110
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
												m.G0 = v11 + int32(16)
												return v14
											}
										}
									}
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v110 == int32(0) {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v120 = v110
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
									m.G0 = v11 + int32(16)
									return v14
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = base.I32_extend8_s(v55)
									F_errmsg_internal(m, int32(_a_F_array_create_iterator_4), v11)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_array_create_iterator_5), int32(322), int32(_a_F_array_create_iterator_6))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 6:
								v76 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
								if l1 != 0 {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v80 = int32(2)
									v84 = l1 << (uint(v80) % 32)
									v85 = v35 + v79<<(uint(v80)%32) - v84
									*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v89 = v87 << (uint(v80) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
									v94 = F_ArrayGetNItemsSafe(m, l1, v85)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
										v99 = F_palloc(m, v94<<(uint(int32(3))%32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
											v103 = F_palloc(m, v102)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
												v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v110 == int32(0) {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v120 = v110
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
												m.G0 = v11 + int32(16)
												return v14
											}
										}
									}
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v110 == int32(0) {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v120 = v110
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
									m.G0 = v11 + int32(16)
									return v14
								}
							case 16:
								v76 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v76)
								if l1 != 0 {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v80 = int32(2)
									v84 = l1 << (uint(v80) % 32)
									v85 = v35 + v79<<(uint(v80)%32) - v84
									*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v85
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v89 = v87 << (uint(v80) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35 + v89 + v89 - v84
									v94 = F_ArrayGetNItemsSafe(m, l1, v85)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v94
										v99 = F_palloc(m, v94<<(uint(int32(3))%32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v99
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
											v103 = F_palloc(m, v102)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v103
												v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v110 == int32(0) {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v120 = v110
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
												m.G0 = v11 + int32(16)
												return v14
											}
										}
									}
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v110 == int32(0) {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v120 = (v113<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v120 = v110
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l0 + v120
									m.G0 = v11 + int32(16)
									return v14
								}
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
	var v51 int64
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
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l4+v47<<(uint(int32(3))%32))))
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v75 int32
	_ = v75
	F_pushJsonbValue(m, l0, int32(4), int32(0))
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
	v23 = int32(1)
	v24 = l1 + v23
	v27 = v23
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_pushJsonbValue(m, l0, int32(5), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	if l2 == v24 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v57 = v27 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v57 <= v58 {
		v27 = v57
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l4+v39<<(uint(int32(3))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v39))))
	F_datum_to_jsonb_internal(m, v43, v45, l0, l7, l8, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_array_dim_to_jsonb(m, l0, v24, l2, l3, l4, l5, l6, l7, l8)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v49 + int32(1)
	goto L8
L13:
	;
	goto L8
L14:
	;
	goto L7
L15:
	;
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
					F_errfinish(m, int32(_a_F_array_exec_setup_1), int32(496), int32(_a_F_array_exec_setup_2))
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
							v39 = int32(1378)
						} else {
							v39 = int32(1379)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v39
						if v14 != 0 {
							v43 = int32(1380)
						} else {
							v43 = int32(1381)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v43
						if v14 != 0 {
							v47 = int32(1382)
						} else {
							v47 = int32(1383)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1384)
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
					F_errfinish(m, int32(_a_F_array_exec_setup_1), int32(491), int32(_a_F_array_exec_setup_2))
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
func F_array_fill_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32 {
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
	var v289 int32
	_ = v289
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
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int64
	_ = v432
	var v434 int32
	_ = v434
	var v443 int64
	_ = v443
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	v16 = m.G0
	v18 = v16 - int32(96)
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
	v718 = m.ExcPending
	if v718 != 0 {
		goto L15
	} else {
		goto L162
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L15
	} else {
		goto L157
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L15
	} else {
		goto L153
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L15
	} else {
		goto L148
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L15
	} else {
		goto L144
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L15
	} else {
		goto L140
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L15
	} else {
		goto L136
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
	v560 = m.ExcPending
	if v560 != 0 {
		goto L15
	} else {
		goto L131
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
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v176
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v176
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v176
	v322 = v18 - int32(-64)
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
		v289 = int32(0)
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
		v289 = v194
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
	v289 = v194
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
	if v289 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v305 = v289
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
	m.G0 = v18 + int32(96)
	return v539
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
	v539 = v331
	goto L73
L78:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+6)))
	v368 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366)+4)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+7)))
	switch v370 - int32(99) {
	case 0:
		v392 = int32(1)
		goto L86
	case 1:
		goto L89
	default:
		goto L88
	case 6:
		goto L90
	case 16:
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
	if l3 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v392 = int32(2)
	goto L86
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L15
	} else {
		goto L91
	}
L89:
	;
	v392 = int32(8)
	goto L86
L90:
	;
	v392 = int32(4)
	goto L86
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = base.I32_extend8_s(v370)
	F_errmsg_internal(m, int32(_a_F_array_fill_internal_0), v18+int32(32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L15
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_1), int32(322), int32(_a_F_array_fill_internal_2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L15
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	if v368 != int32(-1) {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	goto L96
L96:
	;
	v511 = base.I32_div_s(v324+int32(7), int32(8))
	v518 = (v511 + v173<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v519 = F_palloc0(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L15
	} else {
		goto L126
	}
L97:
	;
	v443 = base.I64_extend_i32_s((v434+v392-int32(1))&(int32(0)-v392)) * base.I64_extend_i32_s(v324)
	v447 = base.I32_wrap_i64(v443)
	if base.B2i32(base.I32_wrap_i64(int64(base.Ui64(v443)>>(uint(int64(32))%64))) != v447>>(uint(int32(31))%32))|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v447)) != 0 {
		goto L1
	} else {
		goto L114
	}
L98:
	;
	v428 = F_strlen(m, base.I32_wrap_i64(l2))
	mBase = m.M
	v432 = l2
	v434 = v428 + int32(1)
	goto L97
L99:
	;
	if v368 <= int32(0) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v400 = F_pg_detoast_datum(m, base.I32_wrap_i64(l2))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L15
	} else {
		goto L103
	}
L102:
	;
	v432 = l2
	v434 = v368
	goto L97
L103:
	;
	v402 = base.I64_extend_i32_u(v400)
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v403 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v407 = int32(18)
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
	if v409 == v407 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v420 = int32(1)
	if v403&v420 != 0 {
		v432 = v402
		v434 = int32(base.Ui32(v403) >> (uint(v420) % 32))
		goto L97
	} else {
		goto L113
	}
L107:
	;
	v412 = v407
	goto L109
L108:
	;
	v412 = int32(2)
	goto L109
L109:
	;
	if base.Ui32((v409-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v419 = int32(6)
	goto L112
L111:
	;
	v419 = v412
	goto L112
L112:
	;
	v432 = v402
	v434 = v419
	goto L97
L113:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v432 = v402
	v434 = int32(base.Ui32(v424) >> (uint(int32(2)) % 32))
	goto L97
L114:
	;
	v459 = (v173<<(uint(int32(3))%32) + int32(23)) & int32(120)
	v460 = v447 + v459
	v461 = F_palloc0(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L15
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461)+12)) = l4
	v464 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v461)+8)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v461)+4)) = v173
	v468 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = v460 << (uint(v468) % 32)
	v472 = v461 + int32(16)
	v474 = v173 << (uint(v468) % 32)
	v476 = base.B2i32(v474 == v464)
	if v476 == v464 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	base.MemoryCopy(m, v472, v323, v474)
	goto L118
L117:
	;
	goto L118
L118:
	;
	if v476 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	base.MemoryCopy(m, v474+v472, v322, v474)
	goto L121
L120:
	;
	goto L121
L121:
	;
	v493 = v461 + v459
	v494 = v464
	goto L122
L122:
	;
	v502 = F_ArrayCastAndSet(m, v432, v368, v367&int32(1), v392, v493)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L15
	} else {
		goto L124
	}
L123:
	;
	v539 = v461
	goto L73
L124:
	;
	v506 = v494 + int32(1)
	if v506 != v324 {
		v493 = v502 + v493
		v494 = v506
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v519)+8)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v173
	v524 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v518 << (uint(v524) % 32)
	v528 = v519 + int32(16)
	v530 = v173 << (uint(v524) % 32)
	v531 = int32(0)
	v532 = base.B2i32(v530 == v531)
	if v532 == v531 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	base.MemoryCopy(m, v528, v323, v530)
	goto L129
L128:
	;
	goto L129
L129:
	;
	if v530 == v531 {
		v539 = v519
		goto L73
	} else {
		goto L130
	}
L130:
	;
	base.MemoryCopy(m, v530+v528, v322, v530)
	v539 = v519
	goto L73
L131:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_3), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L15
	} else {
		goto L133
	}
L133:
	;
	v570 = F_errdetail(m, int32(_a_F_array_fill_internal_4), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L15
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_6), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L15
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L15
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_8), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L15
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_9), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L15
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L15
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v168
	F_errmsg(m, int32(_a_F_array_fill_internal_10), v18)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L15
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_11), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L15
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L15
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v168
	F_errmsg(m, int32(_a_F_array_fill_internal_12), v18+int32(16))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L15
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_13), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L15
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L15
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_3), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L15
	} else {
		goto L150
	}
L150:
	;
	v657 = F_errdetail(m, int32(_a_F_array_fill_internal_4), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L15
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_14), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L15
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L15
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_8), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L15
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_15), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L15
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L15
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_array_fill_internal_3), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L15
	} else {
		goto L159
	}
L159:
	;
	v708 = F_errdetail(m, int32(_a_F_array_fill_internal_16), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L15
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_17), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L15
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L15
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_fill_internal_18), v18+int32(48))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L15
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_array_fill_internal_5), int32(_a_F_array_fill_internal_19), int32(_a_F_array_fill_internal_7))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L15
	} else {
		goto L165
	}
L165:
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v2 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		F_pfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
func F_array_get_element(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int64
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int64
	_ = v430
	var v431 int64
	_ = v431
	var v432 int64
	_ = v432
	var v433 int64
	_ = v433
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v464 int64
	_ = v464
	v9 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	if v9 < l3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v19 + int32(16)
	return v464
L2:
	;
	if base.B2i32(l1 != v245)|base.B2i32(base.Ui32(v245-int32(7)) < base.Ui32(int32(-6))) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	v26 = base.I32_div_s(l3, l4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v26
	v244 = v19 + int32(12)
	v245 = int32(1)
	v246 = base.I32_wrap_i64(l0)
	v247 = v9
	v250 = v19 + int32(8)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v33 = base.I32_wrap_i64(l0)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v224 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L32
	} else {
		goto L37
	}
L7:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v37&int32(254) != int32(2) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
	goto L11
L9:
	;
	v110 = int32(0)
	v119 = l1 - int32(1)
	if v119 < v110 {
		v197 = v110
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v69 = v55
	goto L16
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	if base.B2i32(v44 != l1)|base.B2i32(base.Ui32(v44-int32(7)) < base.Ui32(int32(-6))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
	v55 = int32(0)
	if l1 <= v55 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v58)
	v464 = int64(0)
	goto L1
L15:
	;
	goto L10
L16:
	;
	v78 = v69 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2+v78)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v53+v78)))
	if v80 < v82 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v91)
	v464 = int64(0)
	goto L1
L18:
	;
	goto L17
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v54+v78)))
	if v85+v82 <= v80 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v89 = v69 + int32(1)
	if l1 != v89 {
		v69 = v89
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	F_deconstruct_expanded_array(m, v43)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	goto L22
L24:
	;
	v122 = int32(1)
	if v119 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v131 = v119
	v132 = v122
	v133 = v110
	v138 = v110
	goto L28
L26:
	;
	v174 = v119
	v175 = v122
	v176 = v110
	goto L27
L27:
	;
	v183 = v174 << (uint(int32(2)) % 32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2+v183)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v53)))
	v197 = (v185-v187)*v175 + v176
	goto L23
L28:
	;
	v139 = int32(2)
	v140 = v131 << (uint(v139) % 32)
	v142 = v140 - int32(4)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2+v142)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v53+v142)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v140+v54)))
	v150 = v149 * v132
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v140+l2)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v140+v53)))
	v159 = (v144-v146)*v150 + ((v153-v155)*v132 + v133)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v54+v142)))
	v162 = v161 * v150
	v164 = v131 - v139
	v166 = v138 + v139
	if v166 != l1&int32(-2) {
		v131 = v164
		v132 = v162
		v133 = v159
		v138 = v166
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if l1&int32(1) == int32(0) {
		v197 = v159
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v174 = v164
	v175 = v162
	v176 = v159
	goto L27
L32:
	;
	return int64(0)
L33:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	if v208 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v218)
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v207+v197<<(uint(int32(3))%32))))
	v464 = v223
	goto L1
L35:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v208))))
	if v212 != int32(1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v215 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v215)
	v464 = int64(0)
	goto L1
L37:
	;
	v227 = v224 + int32(16)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v230 = v228 << (uint(int32(3)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	if v233 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v234 = v227 + v230
	goto L40
L39:
	;
	v234 = int32(0)
	goto L40
L40:
	;
	if v233 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v239 = v233
	goto L43
L42:
	;
	v239 = (v230 + int32(23)) & int32(-8)
	goto L43
L43:
	;
	v244 = v227
	v245 = v228
	v246 = v224 + v239
	v247 = v234
	v250 = v227 + v228<<(uint(int32(2))%32)
	goto L2
L44:
	;
	v314 = int32(0)
	v323 = l1 - int32(1)
	if v323 < v314 {
		v401 = v314
		goto L57
	} else {
		goto L58
	}
L45:
	;
	v273 = v259
	goto L50
L46:
	;
	v259 = int32(0)
	if l1 <= v259 {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v262)
	v464 = int64(0)
	goto L1
L49:
	;
	goto L45
L50:
	;
	v282 = v273 << (uint(int32(2)) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2+v282)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v250+v282)))
	if v284 < v286 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v295 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v295)
	v464 = int64(0)
	goto L1
L52:
	;
	goto L51
L53:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v244+v282)))
	if v289+v286 <= v284 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v293 = v273 + int32(1)
	if l1 != v293 {
		v273 = v293
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	if v247 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	goto L56
L58:
	;
	v326 = int32(1)
	if v323 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v335 = v323
	v336 = v326
	v337 = v314
	v342 = v314
	goto L62
L60:
	;
	v378 = v323
	v379 = v326
	v380 = v314
	goto L61
L61:
	;
	v387 = v378 << (uint(int32(2)) % 32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l2+v387)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v387+v250)))
	v401 = (v389-v391)*v379 + v380
	goto L57
L62:
	;
	v343 = int32(2)
	v344 = v335 << (uint(v343) % 32)
	v346 = v344 - int32(4)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l2+v346)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v250+v346)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v344+v244)))
	v354 = v353 * v336
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v344+l2)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v344+v250)))
	v363 = (v348-v350)*v354 + ((v357-v359)*v336 + v337)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v244+v346)))
	v366 = v365 * v354
	v368 = v335 - v343
	v370 = v342 + v343
	if v370 != l1&int32(-2) {
		v335 = v368
		v336 = v366
		v337 = v363
		v342 = v370
		goto L62
	} else {
		goto L64
	}
L63:
	;
	if l1&int32(1) == int32(0) {
		v401 = v363
		goto L57
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v378 = v368
	v379 = v366
	v380 = v363
	goto L61
L66:
	;
	v421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v421)
	v424 = F_array_seek(m, v246, v421, v247, v401, l4, l6)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L32
	} else {
		goto L69
	}
L67:
	;
	v410 = base.I32_div_s(v401, int32(8))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247+v410))))
	if int32(base.Ui32(v412)>>(uint(v401&int32(7))%32))&int32(1) != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v418)
	v464 = int64(0)
	goto L1
L69:
	;
	if l5 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if base.I32_popcnt(l4) != int32(1) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v464 = base.I64_extend_i32_u(v424)
	goto L1
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L32
	} else {
		goto L79
	}
L74:
	;
	switch base.I32_ctz(l4) {
	case 0:
		goto L78
	case 1:
		goto L77
	case 2:
		goto L76
	case 3:
		goto L75
	default:
		goto L73
	}
L75:
	;
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
	v464 = v433
	goto L1
L76:
	;
	v432 = int64(*(*int32)(unsafe.Add(mBase, uint32(v424))))
	v464 = v432
	goto L1
L77:
	;
	v431 = int64(*(*int16)(unsafe.Add(mBase, uint32(v424))))
	v464 = v431
	goto L1
L78:
	;
	v430 = int64(*(*int8)(unsafe.Add(mBase, uint32(v424))))
	v464 = v430
	goto L1
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l4
	F_errmsg_internal(m, int32(_a_F_array_get_element_0), v19)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L32
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_array_get_element_1), int32(123), int32(_a_F_array_get_element_2))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L32
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_le(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v2 <= int32(0)))
	}
}
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v6 = int32(-1)
	v10 = F_array_get_element(m, base.I64_extend_i32_u(l0), int32(1), l1, v6, v6, int32(0), int32(105), l2)
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		return v10
	}
}
func F_array_subscript_fetch_old_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int64
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
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 != 0 {
		v29 = int32(1)
		v32 = int64(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+64)) = uint8(v29)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v32
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
		v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+9)))
		v27 = F_array_get_slice(m, v18, v19, v13+int32(12), v13+int32(36), v22, v23, v24, v25, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = int32(0)
			v32 = v27
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+64)) = uint8(v29)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v32
			return
		}
	}
}
func F_array_to_halfvec(m *base.Module, l0 int32) int64 {
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
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 float32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
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
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L69
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L61
	}
L4:
	;
	return int64(0)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v223 = m.ExcPending
	if v223 != 0 {
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
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v127 <= int32(0) {
		goto L21
	} else {
		goto L39
	}
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v101 <= int32(0) {
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
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v85+v75<<(uint(int32(3))%32))))
	v90 = F_DirectFunctionCall1Coll(m, int32(1459), int32(0), v89)
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
	v94 = F_Float4ToHalf(m, base.F32_reinterpret_i32(base.I32_wrap_i64(v90)))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v75<<(uint(int32(1))%32)))) = uint16(v94)
	v98 = v75 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v98 < v99 {
		v75 = v98
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v107 = int32(0)
	goto L35
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v119 = *(*float32)(unsafe.Add(mBase, uint32(v115+v107<<(uint(int32(3))%32))))
	v120 = F_Float4ToHalf(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	goto L21
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v107<<(uint(int32(1))%32)))) = uint16(v120)
	v124 = v107 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v124 < v125 {
		v107 = v124
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v133 = int32(0)
	goto L40
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v141+v133<<(uint(int32(3))%32))))
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
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(8)+v133<<(uint(int32(1))%32)))) = uint16(v147)
	v151 = v133 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v151 < v152 {
		v133 = v151
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
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v160<<(uint(int32(3))%32))))
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
	return base.I64_extend_i32_u(v56)
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
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_array_to_halfvec_0), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(456), int32(_a_F_array_to_halfvec_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_array_to_halfvec_3), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(461), int32(_a_F_array_to_halfvec_2))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
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
	v258 = m.ExcPending
	if v258 != 0 {
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
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(92), int32(_a_F_array_to_halfvec_5))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
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
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_array_to_halfvec_6), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_array_to_halfvec_1), int32(495), int32(_a_F_array_to_halfvec_2))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
func F_array_to_json_pretty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_initStringInfo(m, v7)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		F_array_to_json_internal(m, v10, v7, base.B2i32(v9 != int64(0)))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v21 = F_cstring_to_text_with_len(m, v19, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_u(v21)
			}
		}
	}
}
func F_array_to_text_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = l1 + int32(16)
	v25 = F_ArrayGetNItemsSafe(m, v22, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(80)
	return v301
L2:
	;
	return int32(0)
L3:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = F_palloc(m, int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_initStringInfo(m, v20-int32(-64))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(16)
	v301 = v32
	goto L1
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v42 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+6)))
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+4)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+7)))
	switch v86 - int32(99) {
	case 0:
		v106 = int32(1)
		goto L18
	case 1:
		goto L21
	default:
		goto L20
	case 6:
		goto L22
	case 16:
		goto L19
	}
L10:
	;
	F_get_type_io_data(m, v36, int32(1), v58+int32(4), v58+int32(6), v58+int32(7), v58+int32(8), v58+int32(12), v58+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L16
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v47 = F_MemoryContextAlloc(m, v45, int32(48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v56 == v36 {
		v82 = v42
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v36 ^ int32(-1)
	v58 = v52
	goto L10
L15:
	;
	v58 = v42
	goto L10
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	F_fmgr_info_cxt(m, v74, v58+int32(20), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v36
	v82 = v58
	goto L9
L18:
	;
	if int32(0) < v25 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v106 = int32(2)
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	v106 = int32(8)
	goto L18
L22:
	;
	v106 = int32(4)
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = base.I32_extend8_s(v86)
	F_errmsg_internal(m, int32(_a_F_array_to_text_internal_0), v20)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_array_to_text_internal_1), int32(322), int32(_a_F_array_to_text_internal_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v111 = v109 << (uint(int32(3)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v114 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v289 = v287 + int32(4)
	v290 = F_palloc(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L93
	}
L29:
	;
	v115 = v24 + v111
	goto L31
L30:
	;
	v115 = int32(0)
	goto L31
L31:
	;
	if v114 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v126 = v114
	goto L34
L33:
	;
	v126 = (v111 + int32(23)) & int32(-8)
	goto L34
L34:
	;
	v129 = int32(1)
	v134 = int32(0)
	v137 = v115
	v138 = l1 + v126
	v141 = v129
	v143 = v134
	v144 = v134
	goto L35
L35:
	;
	if v137 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	goto L28
L37:
	;
	v256 = int32(1)
	v258 = v141 << (uint(v256) % 32)
	v260 = base.B2i32(v258 == int32(256))
	if v258 == int32(256) {
		goto L83
	} else {
		goto L84
	}
L38:
	;
	v253 = v249
	v254 = int32(1)
	goto L37
L39:
	;
	if v83&v129 != 0 {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v141&v156 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if l3 == int32(0) {
		v253 = v138
		v254 = v143
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v143 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = l2
	F_appendStringInfo(m, v20-int32(-64), int32(_a_F_array_to_text_internal_3), v20+int32(48))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_appendStringInfoString(m, v20-int32(-64), l3)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L47
	}
L46:
	;
	v249 = v138
	goto L38
L47:
	;
	v249 = v138
	goto L38
L48:
	;
	v194 = F_OutputFunctionCall(m, v82+int32(20), v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L61
	}
L49:
	;
	if base.I32_popcnt(v84) != v129 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v193 = base.I64_extend_i32_u(v138)
	goto L48
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L58
	}
L53:
	;
	switch base.I32_ctz(v84) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	default:
		goto L52
	}
L54:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	v193 = v176
	goto L48
L55:
	;
	v175 = int64(*(*int32)(unsafe.Add(mBase, uint32(v138))))
	v193 = v175
	goto L48
L56:
	;
	v174 = int64(*(*int16)(unsafe.Add(mBase, uint32(v138))))
	v193 = v174
	goto L48
L57:
	;
	v173 = int64(*(*int8)(unsafe.Add(mBase, uint32(v138))))
	v193 = v173
	goto L48
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v84
	F_errmsg_internal(m, int32(_a_F_array_to_text_internal_4), v20+int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_array_to_text_internal_1), int32(123), int32(_a_F_array_to_text_internal_5))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
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
	if v143 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if int32(0) < v84 {
		v246 = v138 + v84
		goto L68
	} else {
		goto L69
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l2
	F_appendStringInfo(m, v20-int32(-64), int32(_a_F_array_to_text_internal_3), v20+int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	F_appendStringInfoString(m, v20-int32(-64), v194)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L67
	}
L66:
	;
	goto L62
L67:
	;
	goto L62
L68:
	;
	v249 = (v246 + (v106 - int32(1))) & (int32(0) - v106)
	goto L38
L69:
	;
	if v84 == int32(-1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v214 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v241 = F_strlen(m, v138)
	mBase = m.M
	v246 = v241 + v138 + int32(1)
	goto L68
L73:
	;
	v218 = int32(18)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v220 == v218 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v232 = int32(1)
	if v214&v232 != 0 {
		v246 = v138 + int32(base.Ui32(v214)>>(uint(v232)%32))
		goto L68
	} else {
		goto L82
	}
L76:
	;
	v223 = v218
	goto L78
L77:
	;
	v223 = int32(2)
	goto L78
L78:
	;
	if base.Ui32((v220-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v230 = int32(6)
	goto L81
L80:
	;
	v230 = v223
	goto L81
L81:
	;
	v246 = v138 + v230
	goto L68
L82:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v246 = v138 + int32(base.Ui32(v237)>>(uint(int32(2))%32))
	goto L68
L83:
	;
	v261 = v256
	goto L85
L84:
	;
	v261 = v258
	goto L85
L85:
	;
	if v137 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v262 = v261
	goto L88
L87:
	;
	v262 = v141
	goto L88
L88:
	;
	if v137 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v265 = v137 + v260
	goto L91
L90:
	;
	v265 = int32(0)
	goto L91
L91:
	;
	v267 = v144 + int32(1)
	if v267 != v25 {
		v137 = v265
		v138 = v253
		v141 = v262
		v143 = v254
		v144 = v267
		goto L35
	} else {
		goto L92
	}
L92:
	;
	goto L36
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v289 << (uint(int32(2)) % 32)
	if v287 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	base.MemoryCopy(m, v290+int32(4), v286, v287)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	F_pfree(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v301 = v290
	goto L1
}
func F_array_to_text_null(m *base.Module, l0 int32) int64 {
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
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v8 == v2 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v11 != int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v19 = F_pg_detoast_datum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = F_pg_detoast_datum_packed(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
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
							return int64(0)
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
									return int64(0)
								} else {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
									if v76 != 0 {
										v133 = v2
										v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v134)
										}
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										v78 = F_pg_detoast_datum_packed(m, v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int64(0)
										} else {
											v80 = F_pg_detoast_datum_packed(m, v78)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int64(0)
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
													return int64(0)
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
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v134)
														}
													} else {
														F_pfree(m, v80)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int64(0)
														} else {
															v133 = v114
															v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v134)
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
								if v76 != 0 {
									v133 = v2
									v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v134)
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									v78 = F_pg_detoast_datum_packed(m, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int64(0)
									} else {
										v80 = F_pg_detoast_datum_packed(m, v78)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int64(0)
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
												return int64(0)
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
														return int64(0)
													} else {
														return base.I64_extend_i32_u(v134)
													}
												} else {
													F_pfree(m, v80)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int64(0)
													} else {
														v133 = v114
														v134 = F_array_to_text_internal(m, l0, v19, v60, v133)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v134)
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
			return int64(0)
		}
	} else {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		return int64(0)
	}
}
func F_compute_array_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
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
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int64
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v334 int64
	_ = v334
	var v339 int64
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v391 int64
	_ = v391
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v426 int32
	_ = v426
	var v433 int64
	_ = v433
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v458 int64
	_ = v458
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v466 int64
	_ = v466
	var v467 int64
	_ = v467
	var v468 int64
	_ = v468
	var v469 int64
	_ = v469
	var v470 int64
	_ = v470
	var v471 int64
	_ = v471
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v477 int64
	_ = v477
	var v478 int64
	_ = v478
	var v479 int64
	_ = v479
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v486 int64
	_ = v486
	var v487 int64
	_ = v487
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	var v490 int64
	_ = v490
	var v491 int64
	_ = v491
	var v523 int64
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int64
	_ = v534
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v561 int64
	_ = v561
	var v567 int64
	_ = v567
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v578 int32
	_ = v578
	var v579 int64
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v604 int64
	_ = v604
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int64
	_ = v639
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v643 int32
	_ = v643
	var v644 int64
	_ = v644
	var v645 int64
	_ = v645
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 float64
	_ = v673
	var v680 int32
	_ = v680
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int64
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int64
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v725 float64
	_ = v725
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
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
	var v803 int32
	_ = v803
	var v817 int32
	_ = v817
	var v818 int64
	_ = v818
	var v819 int64
	_ = v819
	var v822 int64
	_ = v822
	var v823 int64
	_ = v823
	var v824 int64
	_ = v824
	var v825 int64
	_ = v825
	var v826 int64
	_ = v826
	var v827 int64
	_ = v827
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v830 int64
	_ = v830
	var v831 int64
	_ = v831
	var v832 int64
	_ = v832
	var v833 int64
	_ = v833
	var v834 int64
	_ = v834
	var v835 int64
	_ = v835
	var v836 int64
	_ = v836
	var v837 int64
	_ = v837
	var v838 int64
	_ = v838
	var v839 int64
	_ = v839
	var v840 int64
	_ = v840
	var v841 int64
	_ = v841
	var v842 int64
	_ = v842
	var v843 int64
	_ = v843
	var v844 int64
	_ = v844
	var v845 int64
	_ = v845
	var v846 int64
	_ = v846
	var v847 int64
	_ = v847
	var v848 int64
	_ = v848
	var v849 int64
	_ = v849
	var v850 int64
	_ = v850
	var v851 int64
	_ = v851
	var v852 int64
	_ = v852
	var v884 int64
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v984 int32
	_ = v984
	var v986 int64
	_ = v986
	var v987 int32
	_ = v987
	var v988 int64
	_ = v988
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1013 int64
	_ = v1013
	var v1024 int32
	_ = v1024
	var v1030 int32
	_ = v1030
	var v1044 int64
	_ = v1044
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int64
	_ = v1056
	var v1058 int64
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1080 int64
	_ = v1080
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	v5 = int32(0)
	v20 = int64(0)
	v26 = m.G0
	v28 = v26 - int32(176)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	m.T0[v33].(func(*base.Module, int32, int32, int32, float64))(m, l0, l1, l2, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v30
	*(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[0])) = v30
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = int32(1373)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = int32(1374)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+112)) = int64(103079215112)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v47
	v51 = v39 * int32(10)
	v56 = F_hash_create(m, int32(_a_F_compute_array_stats_0), base.I64_extend_i32_s(v51), v28+int32(104), int32(1224))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+64)) = int64(34359738372)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v61
	v66 = base.I32_div_s(v39*int32(_a_F_compute_array_stats_1), int32(7))
	v72 = F_hash_create(m, int32(_a_F_compute_array_stats_2), int64(64), v28+int32(56), int32(1064))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
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
	v426 = v5
	v433 = v20
	v437 = float64(0)
	goto L7
L6:
	;
	v88 = v5
	v90 = int32(1)
	v93 = v5
	v95 = v5
	v100 = v20
	goto L9
L7:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	if v438 == int32(0) {
		v451 = v5
		goto L68
	} else {
		goto L69
	}
L8:
	;
	v426 = v384
	v433 = v391
	v437 = base.F64_convert_i32_s(v379)
	goto L7
L9:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L64
	}
L11:
	;
	goto L10
L12:
	;
	v109 = m.T0[l1].(func(*base.Module, int32, int32, int32) int64)(m, l0, v95, v28+int32(55))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+55)))
	if v111 != 0 {
		v379 = v88
		v381 = v90
		v384 = v93
		v391 = v100
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v396 = v95 + int32(1)
	if l2 != v396 {
		v88 = v379
		v90 = v381
		v93 = v384
		v95 = v396
		v100 = v391
		goto L9
	} else {
		goto L63
	}
L15:
	;
	v112 = F_toast_raw_datum_size(m, v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(_a_F_compute_array_stats_3)) < base.Ui32(v112) {
		v379 = v88
		v381 = v90
		v384 = v93
		v391 = v100
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v117 = F_pg_detoast_datum(m, base.I32_wrap_i64(v109))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+14)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+16)))
	F_deconstruct_array(m, v117, v120, v121, v122, v28+int32(44), v28+int32(40), v28+int32(48))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	if v131 < v133 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v142 = v131
	v146 = v131
	v147 = v90
	v156 = v100
	goto L23
L21:
	;
	v324 = v131
	v325 = v90
	v334 = v100
	goto L22
L22:
	;
	v339 = v334 - v100
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+156)) = uint32(v339)
	v346 = F_hash_search(m, v72, v28+int32(156), int32(1), v28+int32(32))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L53
	}
L23:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v142))))
	if v163 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v324 = v295
	v325 = v296
	v334 = v305
	goto L22
L25:
	;
	v311 = v142 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	if v311 < v312 {
		v142 = v311
		v146 = v295
		v147 = v296
		v156 = v305
		goto L23
	} else {
		goto L52
	}
L26:
	;
	v295 = int32(1)
	v296 = v147
	v305 = v156
	goto L25
L27:
	;
	goto L28
L28:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v165+v142<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v169
	v176 = F_hash_search(m, v56, v28+int32(32), int32(1), v28+int32(31))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+31)))
	if v178 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v202 = v156 + int64(1)
	v203 = base.I64_rem_s(v202, base.I64_extend_i32_s(v66))
	if v203 != int64(0) {
		v295 = v146
		v296 = v147
		v305 = v202
		goto L25
	} else {
		goto L36
	}
L31:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	if v181 == v95 {
		v295 = v146
		v296 = v147
		v305 = v156
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+14)))
	v191 = F_datumCopy(m, v188, v189, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+16)) = v95
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v184 + int32(1)
	goto L30
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+16)) = v95
	v194 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v176)+12)) = v147 - v194
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = v191
	goto L30
L36:
	;
	v207 = v28 + int32(156)
	F_hash_seq_init(m, v207, v56)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v210 = F_hash_seq_search(m, v207)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v210 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v217 = v210
	goto L42
L40:
	;
	goto L41
L41:
	;
	v295 = v146
	v296 = v147 + int32(1)
	v305 = v202
	goto L25
L42:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	if v147 < v237+v238 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v256 = F_hash_seq_search(m, v28+int32(156))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L50
	}
L45:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v217)))
	v244 = F_hash_search(m, v56, v217, int32(2), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v244 == int32(0) {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[0]))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+12)))
	if v250 != 0 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, base.I32_wrap_i64(v241))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	if v256 != 0 {
		v217 = v256
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
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+32)))
	if v348 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	v355 = v351 + int32(1)
	goto L56
L55:
	;
	v355 = int32(1)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+4)) = v355
	if base.I64_extend_i32_u(v117) != v109 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v117)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	F_pfree(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	F_pfree(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v379 = v88 + v324
	v381 = v325
	v384 = v93 + int32(1)
	v391 = v334
	goto L14
L63:
	;
	goto L8
L64:
	;
	F_errmsg_internal(m, int32(_a_F_compute_array_stats_4), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_compute_array_stats_5), int32(711), int32(_a_F_compute_array_stats_6))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
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
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L158
	}
L68:
	;
	if v426 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	if v441 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v451 = int32(1)
	goto L68
L71:
	;
	goto L72
L72:
	;
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	if v445 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v451 = int32(2)
	goto L68
L74:
	;
	goto L75
L75:
	;
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)))
	if v449 != 0 {
		goto L67
	} else {
		goto L76
	}
L76:
	;
	v451 = int32(3)
	goto L68
L77:
	;
	m.G0 = v28 + int32(176)
	return
L78:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v456)+8))
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v456)+808))
	if v458 != int64(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v524 = base.I32_wrap_i64(v523)
	v525 = F_palloc_mul(m, int32(4), v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v456)+752))
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v456)+728))
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v456)+704))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v456)+680))
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v456)+656))
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v456)+632))
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v456)+608))
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v456)+584))
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v456)+560))
	v470 = *(*int64)(unsafe.Add(mBase, uint32(v456)+536))
	v471 = *(*int64)(unsafe.Add(mBase, uint32(v456)+512))
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v456)+488))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v456)+464))
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v456)+440))
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v456)+416))
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v456)+392))
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v456)+368))
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v456)+344))
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v456)+320))
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v456)+296))
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v456)+272))
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v456)+248))
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v456)+224))
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v456)+200))
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v456)+176))
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v456)+152))
	v487 = *(*int64)(unsafe.Add(mBase, uint32(v456)+128))
	v488 = *(*int64)(unsafe.Add(mBase, uint32(v456)+104))
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v456)+80))
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v456)+56))
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v456)+32))
	v523 = v461 + (v462 + (v463 + (v464 + (v465 + (v466 + (v467 + (v468 + (v469 + (v470 + (v471 + (v472 + (v473 + (v474 + (v475 + (v476 + (v477 + (v478 + (v479 + (v480 + (v481 + (v482 + (v483 + (v484 + (v485 + (v486 + (v487 + (v488 + (v489 + (v490 + (v491 + v457))))))))))))))))))))))))))))))
	goto L82
L81:
	;
	v523 = v457
	goto L82
L82:
	;
	goto L79
L83:
	;
	v528 = v28 + int32(156)
	F_hash_seq_init(m, v528, v56)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v534 = base.I64_div_s(v433*int64(9), base.I64_extend_i32_s(v66))
	v535 = int64(0)
	v536 = F_hash_seq_search(m, v528)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v612 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L100
	}
L86:
	;
	if v536 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v591 = int32(0)
	v604 = v535
	goto L85
L88:
	;
	goto L89
L89:
	;
	v547 = v536
	v548 = int32(0)
	v561 = v535
	goto L90
L90:
	;
	v567 = int64(*(*int32)(unsafe.Add(mBase, uint32(v547)+8)))
	if v534 < v567 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v591 = v578
	v604 = v579
	goto L85
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v525+v548<<(uint(int32(2))%32)))) = v547
	v573 = int64(*(*int32)(unsafe.Add(mBase, uint32(v547)+8)))
	if v573 < v561 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v578 = v548
	v579 = v561
	goto L94
L94:
	;
	v583 = F_hash_seq_search(m, v28+int32(156))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	v575 = v561
	goto L97
L96:
	;
	v575 = v573
	goto L97
L97:
	;
	v578 = v548 + int32(1)
	v579 = v575
	goto L94
L98:
	;
	if v583 != 0 {
		v547 = v583
		v548 = v578
		v561 = v579
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L91
L100:
	;
	if v612 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v524
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v51
	F_errmsg_internal(m, int32(_a_F_compute_array_stats_7), v28)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v51 < v591 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	F_errfinish(m, int32(_a_F_compute_array_stats_5), int32(492), int32(_a_F_compute_array_stats_8))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v647 = l0 + int32(52)
	if int32(0) <= v643 {
		goto L114
	} else {
		goto L115
	}
L107:
	;
	F_qsort_interruptible(m, v525, v591, int32(4), int32(1375), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v641 = v534 + int64(1)
	if v591 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v525+v51<<(uint(int32(2))%32)-int32(4))))
	v639 = int64(*(*int32)(unsafe.Add(mBase, uint32(v638)+8)))
	v643 = v51
	v644 = v604
	v645 = v639
	goto L106
L111:
	;
	v642 = v604
	goto L113
L112:
	;
	v642 = v641
	goto L113
L113:
	;
	v643 = v591
	v644 = v642
	v645 = v641
	goto L106
L114:
	;
	F_qsort_interruptible(m, v525, v643, int32(4), int32(1376), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	v803 = v451
	goto L116
L116:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v817)+8))
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v817)+808))
	if v819 != int64(0) {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v655 = int32(_a_F_compute_array_stats_9)
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1]))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1])) = v658
	v662 = F_palloc(m, v643<<(uint(int32(3))%32))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v665 = v643 + int32(3)
	v668 = F_palloc(m, v665<<(uint(int32(2))%32))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v643 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v747 = int32(2)
	v749 = v668 + v643<<(uint(v747)%32)
	*(*float32)(unsafe.Add(mBase, uint32(v749)+8)) = base.F32_demote_f64(base.F64_div(v437, v725))
	*(*float32)(unsafe.Add(mBase, uint32(v749)+4)) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v644), v725))
	*(*float32)(unsafe.Add(mBase, uint32(v749))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v645), v725))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_array_stats[1])) = v656
	v763 = int32(1)
	v764 = v451 << (uint(v763) % 32)
	v766 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v647+v764))) = uint16(v766)
	v770 = l0 + v451<<(uint(v747)%32)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+64)) = v771
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+124)) = v668
	*(*int32)(unsafe.Add(mBase, uint32(v770)+84)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v770)+164)) = v662
	*(*int32)(unsafe.Add(mBase, uint32(v770)+104)) = v665
	*(*int32)(unsafe.Add(mBase, uint32(v770)+144)) = v643
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+184)) = v779
	v782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v764)+204)) = uint16(v782)
	v784 = l0 + v451
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+214)) = uint8(v785)
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+219)) = uint8(v787)
	v803 = v451 + v763
	goto L116
L121:
	;
	v725 = base.F64_convert_i32_u(v426)
	goto L120
L122:
	;
	goto L123
L123:
	;
	v673 = base.F64_convert_i32_u(v426)
	v680 = int32(0)
	goto L124
L124:
	;
	v704 = v680 << (uint(int32(2)) % 32)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v525+v704)))
	v707 = *(*int64)(unsafe.Add(mBase, uint32(v706)))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v709 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+14)))
	v710 = F_datumCopy(m, v707, v708, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L126
	}
L125:
	;
	v725 = v673
	goto L120
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v662+v680<<(uint(int32(3))%32)))) = v710
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v706)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v704+v668))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v714), v673))
	v720 = v680 + int32(1)
	if v720 != v643 {
		v680 = v720
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v885 = base.I32_wrap_i64(v884)
	if v885 <= int32(0) {
		goto L77
	} else {
		goto L132
	}
L129:
	;
	v822 = *(*int64)(unsafe.Add(mBase, uint32(v817)+752))
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v817)+728))
	v824 = *(*int64)(unsafe.Add(mBase, uint32(v817)+704))
	v825 = *(*int64)(unsafe.Add(mBase, uint32(v817)+680))
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v817)+656))
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v817)+632))
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v817)+608))
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v817)+584))
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v817)+560))
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v817)+536))
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v817)+512))
	v833 = *(*int64)(unsafe.Add(mBase, uint32(v817)+488))
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v817)+464))
	v835 = *(*int64)(unsafe.Add(mBase, uint32(v817)+440))
	v836 = *(*int64)(unsafe.Add(mBase, uint32(v817)+416))
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v817)+392))
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v817)+368))
	v839 = *(*int64)(unsafe.Add(mBase, uint32(v817)+344))
	v840 = *(*int64)(unsafe.Add(mBase, uint32(v817)+320))
	v841 = *(*int64)(unsafe.Add(mBase, uint32(v817)+296))
	v842 = *(*int64)(unsafe.Add(mBase, uint32(v817)+272))
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v817)+248))
	v844 = *(*int64)(unsafe.Add(mBase, uint32(v817)+224))
	v845 = *(*int64)(unsafe.Add(mBase, uint32(v817)+200))
	v846 = *(*int64)(unsafe.Add(mBase, uint32(v817)+176))
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v817)+152))
	v848 = *(*int64)(unsafe.Add(mBase, uint32(v817)+128))
	v849 = *(*int64)(unsafe.Add(mBase, uint32(v817)+104))
	v850 = *(*int64)(unsafe.Add(mBase, uint32(v817)+80))
	v851 = *(*int64)(unsafe.Add(mBase, uint32(v817)+56))
	v852 = *(*int64)(unsafe.Add(mBase, uint32(v817)+32))
	v884 = v822 + (v823 + (v824 + (v825 + (v826 + (v827 + (v828 + (v829 + (v830 + (v831 + (v832 + (v833 + (v834 + (v835 + (v836 + (v837 + (v838 + (v839 + (v840 + (v841 + (v842 + (v843 + (v844 + (v845 + (v846 + (v847 + (v848 + (v849 + (v850 + (v851 + (v852 + v818))))))))))))))))))))))))))))))
	goto L131
L130:
	;
	v884 = v818
	goto L131
L131:
	;
	goto L128
L132:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v890 = F_palloc_mul(m, int32(4), v885)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v893 = v28 + int32(156)
	F_hash_seq_init(m, v893, v72)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v896 = F_hash_seq_search(m, v893)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v896 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v904 = int32(0)
	v905 = v896
	goto L139
L137:
	;
	goto L138
L138:
	;
	v959 = int32(2)
	if v888 <= v959 {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v890+v904<<(uint(int32(2))%32)))) = v905
	v932 = F_hash_seq_search(m, v28+int32(156))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L138
L141:
	;
	if v932 != 0 {
		v904 = v904 + int32(1)
		v905 = v932
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v962 = v959
	goto L145
L144:
	;
	v962 = v888
	goto L145
L145:
	;
	v963 = int32(0)
	F_qsort_interruptible(m, v890, v885, int32(4), int32(1377), v963)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v971 = v962 + int32(1)
	v974 = F_MemoryContextAlloc(m, v969, v971<<(uint(int32(2))%32))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v974+v962<<(uint(int32(2))%32)))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v433), base.F64_convert_i32_u(v426)))
	v984 = int32(1)
	v986 = base.I64_extend_i32_u(v962 - v984)
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	v988 = int64(*(*int32)(unsafe.Add(mBase, uint32(v987)+4)))
	v999 = v963
	v1001 = int32(0)
	v1013 = v986 * v988
	goto L148
L148:
	;
	if int64(0) < v1013 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1099 = int32(5)
	*(*uint16)(unsafe.Add(mBase, uint32(v647+v803<<(uint(int32(1))%32)))) = uint16(v1099)
	v1103 = l0 + v803<<(uint(int32(2))%32)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1103)+64)) = v1104
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1103)+124)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v1103)+84)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v1103)+104)) = v971
	goto L77
L150:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	*(*float32)(unsafe.Add(mBase, uint32(v974+v1001<<(uint(int32(2))%32)))) = base.F32_convert_i32_s(v1089)
	v1094 = v1001 + int32(1)
	if v1094 != v962 {
		v999 = v1066
		v1001 = v1094
		v1013 = v1080 - base.I64_extend_i32_u(v426-v984)
		goto L148
	} else {
		goto L157
	}
L151:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v890+v999<<(uint(int32(2))%32))))
	v1066 = v999
	v1067 = v1024
	v1080 = v1013
	goto L150
L152:
	;
	goto L153
L153:
	;
	v1030 = v999
	v1044 = v1013
	goto L154
L154:
	;
	v1051 = v1030 + int32(1)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v890+v1051<<(uint(int32(2))%32))))
	v1056 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1055)+4)))
	v1058 = v1056*v986 + v1044
	if v1058 <= int64(0) {
		v1030 = v1051
		v1044 = v1058
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v1066 = v1051
	v1067 = v1055
	v1080 = v1058
	goto L150
L156:
	;
	goto L155
L157:
	;
	goto L149
L158:
	;
	F_errmsg_internal(m, int32(_a_F_compute_array_stats_10), int32(0))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_compute_array_stats_5), int32(440), int32(_a_F_compute_array_stats_8))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(99)
	v13 = int32(1)
	if l2 <= int32(699) {
		switch l2 - int32(18) {
		case 0:
			v59 = int32(1)
			v60 = v12
			v61 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		case 1:
			v59 = int32(64)
			v60 = v12
			v61 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		case 2:
			v59 = int32(8)
			v60 = int32(100)
			v61 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		case 3:
			v59 = int32(2)
			v60 = int32(115)
			v61 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
				F_errmsg_internal(m, int32(_a_F_construct_array_builtin_0), v10)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_construct_array_builtin_1), int32(3469), int32(_a_F_construct_array_builtin_2))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 5, 8, 10:
			v59 = int32(4)
			v60 = int32(105)
			v61 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		case 7:
			v59 = int32(-1)
			v60 = int32(105)
			v61 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		case 9:
			v59 = int32(6)
			v60 = int32(115)
			v61 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		}
	} else {
		switch l2 - int32(700) {
		case 0:
			v59 = int32(4)
			v60 = int32(105)
			v61 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		case 1:
			v59 = int32(8)
			v60 = int32(100)
			v61 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v63 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
			v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v71
			}
		default:
			if l2 == int32(2206) {
				v59 = int32(4)
				v60 = int32(105)
				v61 = v13
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
				v63 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
				v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v71
				}
			} else {
				if l2 != int32(2275) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(_a_F_construct_array_builtin_0), v10)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_construct_array_builtin_1), int32(3469), int32(_a_F_construct_array_builtin_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v59 = int32(-2)
					v60 = v12
					v61 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
					v63 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
					v71 = F_construct_md_array(m, l0, int32(0), v63, v10+int32(12), v10+int32(8), l2, v59, v61, v60)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v71
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
	v3 = Fn14291(m, l0, int32(82))
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
