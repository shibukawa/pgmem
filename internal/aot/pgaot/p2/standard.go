package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_standard_ExecutorFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v8 = int32(_a_F_standard_ExecutorFinish_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorFinish[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorFinish[0])) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_InstrStartNode(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+148))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+128)))
	if v70&int32(32) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = int32(0)
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v29<<(uint(int32(2))%32))))
	goto L11
L10:
	;
	goto L6
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+152))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v60 = v29 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v60 < v61 {
		v29 = v60
		goto L9
	} else {
		goto L26
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	F_MemoryContextReset(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	F_ExecReScan(m, v36)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v52 = m.T0[v51].(func(*base.Module, int32) int32)(m, v36)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v54&int32(2) == int32(0) {
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L12
L25:
	;
	goto L24
L26:
	;
	goto L10
L27:
	;
	F_AfterTriggerEndQuery(m, v11)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v77 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	F_InstrStopNode(m, v77, float64(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorFinish[0])) = v9
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+136)) = uint8(v83)
	return
L34:
	;
	goto L33
}
func F_standard_planner(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 float64
	_ = v91
	var v99 float64
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 float64
	_ = v117
	var v123 float64
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v191 float64
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v275 float64
	_ = v275
	var v277 float64
	_ = v277
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v288 float64
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v307 float64
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 float64
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 float64
	_ = v342
	var v343 float64
	_ = v343
	var v346 int32
	_ = v346
	var v347 float64
	_ = v347
	var v348 float64
	_ = v348
	var v350 float64
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 float64
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 float64
	_ = v381
	var v382 float64
	_ = v382
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v400 float64
	_ = v400
	var v403 int32
	_ = v403
	var v405 float64
	_ = v405
	var v406 float64
	_ = v406
	var v409 float64
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int64
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v584 float64
	_ = v584
	var v589 float64
	_ = v589
	var v593 int32
	_ = v593
	var v597 float64
	_ = v597
	var v602 float64
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 float64
	_ = v611
	var v616 float64
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	v12 = float64(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = F_palloc0(m, int32(92))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(266)
	v33 = F__emscripten_memset_bulkmem(m, v19+int32(8), base.I32_extend8_s(v23), int32(74))
	mBase = m.M
	goto L3
L3:
	;
	if l1&int32(2048) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	v82 = int32(0)
	v84 = v79 & base.B2i32(v81 != v82)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+83)) = uint8(v84)
	if l1&int32(256) == v82 {
		v99 = v12
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+82)) = uint8(v74)
	v76 = int32(117)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+84)) = uint8(v76)
	v79 = int32(0)
	goto L4
L6:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[1])))
	if v39 != int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 != int32(1) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	if v45 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[2]))
	if v47 <= int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[3]))
	if int32(0) <= v51 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v54 = m.G0
	v56 = v54 - int32(16)
	m.G0 = v56
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
	v60 = int32(_a_F_standard_planner_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+8)) = uint16(v60)
	v64 = F_max_parallel_hazard_walker(m, l0, v56+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+8)))
	m.G0 = v56 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+84)) = uint8(v66)
	v72 = base.B2i32(v66 != int32(117))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+82)) = uint8(v72)
	v79 = v72
	goto L4
L13:
	;
	v101 = int32(0)
	v104 = F_subquery_planner(m, v19, l0, v101, v101, v99, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v91 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[4]))
	if base.F64_ge(v91, float64(1)) != 0 {
		v99 = v12
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if base.F64_le(v91, float64(0)) == int32(0) {
		v99 = v91
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v99 = float64(1e-10)
	goto L13
L17:
	;
	v108 = F_fetch_upper_rel(m, v104, int32(7), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+48))
	if base.F64_le(v99, float64(0)) != 0 {
		v219 = v110
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v225 = F_create_plan(m, v104, v219)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L53
	}
L20:
	;
	if base.F64_ge(v99, float64(1)) == int32(0) {
		v123 = v99
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v108)+32))
	if v125 == int32(0) {
		v219 = v110
		goto L19
	} else {
		goto L24
	}
L22:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v110)+32))
	if base.F64_gt(v117, float64(0)) == int32(0) {
		v123 = v99
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v123 = base.F64_div(v99, v117)
	goto L21
L24:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v128 <= int32(0) {
		v219 = v110
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v135 = int32(0)
	v138 = v128
	v139 = v110
	goto L26
L26:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v135<<(uint(int32(2))%32))))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	if v150 != 0 {
		v207 = v138
		v208 = v139
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v219 = v208
	goto L19
L28:
	;
	v210 = v135 + int32(1)
	if v210 < v207 {
		v135 = v210
		v138 = v207
		v139 = v208
		goto L26
	} else {
		goto L52
	}
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v108)+48))
	if v149 == v151 {
		v207 = v138
		v208 = v139
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v139)+40))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	if v157 != v158 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v202 <= int32(0) {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	if v157 < v158 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if base.F64_le(v123, float64(0))|base.F64_ge(v123, float64(1)) != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v163 = int32(-1)
	goto L37
L36:
	;
	v163 = int32(1)
	goto L37
L37:
	;
	v202 = v163
	goto L31
L38:
	;
	v202 = v197
	goto L31
L39:
	;
	v169 = int32(-1)
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v139)+56))
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v149)+56))
	if base.F64_lt(v170, v171) != 0 {
		v197 = v169
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v139)+56))
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v139)+48))
	v186 = base.F64_add(base.F64_mul(v123, base.F64_sub(v182, v183)), v183)
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v149)+56))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v149)+48))
	v191 = base.F64_add(base.F64_mul(v123, base.F64_sub(v187, v188)), v188)
	if base.F64_lt(v186, v191) != 0 {
		v197 = int32(-1)
		goto L38
	} else {
		goto L48
	}
L42:
	;
	if base.F64_gt(v170, v171) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v202 = int32(1)
	goto L31
L44:
	;
	goto L45
L45:
	;
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v139)+48))
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v149)+48))
	if base.F64_lt(v175, v176) != 0 {
		v197 = v169
		goto L38
	} else {
		goto L46
	}
L46:
	;
	if base.F64_gt(v175, v176) != 0 {
		v197 = int32(1)
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v202 = int32(0)
	goto L31
L48:
	;
	v197 = base.F64_lt(v191, v186)
	goto L38
L49:
	;
	v205 = v139
	goto L51
L50:
	;
	v205 = v149
	goto L51
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v207 = v206
	v208 = v205
	goto L28
L52:
	;
	goto L27
L53:
	;
	if l1&int32(2) == int32(0) {
		v235 = v225
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	if v237 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v231 = F_ExecSupportsBackwardScan(m, v225)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v231 != 0 {
		v235 = v225
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v233 = F_materialize_finished_plan(m, v225)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v235 = v233
	goto L54
L59:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	if v418 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L60:
	;
	v415 = v235
	goto L59
L61:
	;
	goto L62
L62:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+37)))
	if v240 != int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v415 = v235
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v237 != int32(2) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v249 = F_palloc0(m, int32(88))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v235)+60))
	if v245 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v415 = v235
	goto L59
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = int32(368)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v235)+44))
	v254 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v249)+80)) = uint8(v254)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+72)) = v254
	v258 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+56)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v249)+52)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v249)+48)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v249)+44)) = v253
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v249)+81)) = uint8(base.B2i32(v265 == int32(2)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v235)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+60)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v235)+60)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v249)+76)) = int32(-1)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v235)+8))
	v277 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[5]))
	*(*float64)(unsafe.Add(mBase, uint32(v249)+8)) = base.F64_add(v275, v277)
	v281 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[6]))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v235)+24))
	v284 = *(*float64)(unsafe.Add(mBase, uint32(v235)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v249)+16)) = base.F64_add(base.F64_mul(v281, v282), base.F64_add(v277, v284))
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v235)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v249)+24)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v235)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+36)) = uint16(v258)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+32)) = v290
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v249)+60))
	v307 = float64(0)
	if v294 == v258 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v405 = *(*float64)(unsafe.Add(mBase, uint32(v235)+8))
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v235)+8)) = base.F64_sub(v405, v406)
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v235)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v235)+16)) = base.F64_sub(v409, v406)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v413 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v412)+83)) = uint8(v413)
	v415 = v249
	goto L59
L71:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = v400
	v403 = v393 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(7)))) = uint8(v403)
	goto L70
L72:
	;
	v393 = v258
	v400 = v307
	goto L71
L73:
	;
	goto L74
L74:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v310 <= int32(0) {
		v393 = v258
		v400 = v307
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v313 = int32(0)
	if v313 < v310 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v316 = v310
	goto L78
L77:
	;
	v316 = v313
	goto L78
L78:
	;
	v317 = int32(1)
	if v310 == v317 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v316&v317 == int32(0) {
		v393 = v366
		v400 = v373
		goto L71
	} else {
		goto L86
	}
L80:
	;
	v365 = int32(0)
	v366 = v258
	v373 = v307
	goto L79
L81:
	;
	goto L82
L82:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v329 = int32(0)
	v330 = v258
	v333 = v258
	v337 = v307
	goto L83
L83:
	;
	v338 = int32(2)
	v340 = v324 + v329<<(uint(v338)%32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v342 = *(*float64)(unsafe.Add(mBase, uint32(v341)+56))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v341)+64))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v346)+56))
	v348 = *(*float64)(unsafe.Add(mBase, uint32(v346)+64))
	v350 = base.F64_add(base.F64_add(v337, base.F64_add(v342, v343)), base.F64_add(v347, v348))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+38)))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+38)))
	v356 = v351&v352 ^ int32(1) | v330
	v358 = v329 + v338
	v360 = v333 + v338
	if v360 != v316&int32(2147483646) {
		v329 = v358
		v330 = v356
		v333 = v360
		v337 = v350
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v365 = v358
	v366 = v356
	v373 = v350
	goto L79
L85:
	;
	goto L84
L86:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376+v365<<(uint(int32(2))%32))))
	v381 = *(*float64)(unsafe.Add(mBase, uint32(v380)+56))
	v382 = *(*float64)(unsafe.Add(mBase, uint32(v380)+64))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+38)))
	v393 = v385 ^ int32(1) | v366
	v400 = base.F64_add(v373, base.F64_add(v381, v382))
	goto L71
L87:
	;
	v480 = F_set_plan_references(m, v104, v415)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L102
	}
L88:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v427 = int32(0)
	goto L89
L89:
	;
	v437 = int32(0)
	if v422 == v437 {
		v447 = v437
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v421 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	if v441 <= v427 {
		v447 = int32(0)
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v422)+12))
	v447 = v443 + v427<<(uint(int32(2))%32)
	goto L91
L94:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	F_SS_finalize_plan(m, v461, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L101
	}
L95:
	;
	F_SS_finalize_plan(m, v104, v415)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L100
	}
L96:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v450 <= v427 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v447 == int32(0) {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v421)+12))
	v457 = v454 + v427<<(uint(int32(2))%32)
	if v457 != 0 {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	goto L95
L100:
	;
	goto L87
L101:
	;
	v427 = v427 + int32(1)
	goto L89
L102:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v488 = int32(0)
	goto L103
L103:
	;
	v498 = int32(0)
	if v483 == v498 {
		v508 = v498
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v482 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v502 <= v488 {
		v508 = int32(0)
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	v508 = v504 + v488<<(uint(int32(2))%32)
	goto L105
L108:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v650 = F_set_plan_references(m, v648, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L134
	}
L109:
	;
	v521 = F_palloc0(m, int32(104))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L114
	}
L110:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	if v511 <= v488 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	if v508 == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v482)+12))
	v518 = v515 + v488<<(uint(int32(2))%32)
	if v518 != 0 {
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = int32(330)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+4)) = v525
	v527 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+8)) = v527
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+24)) = uint8(base.B2i32(v529 != int32(0)))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+25)) = uint8(v533)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+26)) = uint8(v535)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+80)))
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+27)) = uint8(v537)
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+81)))
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+28)) = uint8(v539)
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+83)))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+36)) = v480
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+29)) = uint8(v541)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+40)) = v544
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+44)) = v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v550 = F_bms_difference(m, v548, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+48)) = v550
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+52)) = v553
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+56)) = v555
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+60)) = v557
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+64)) = v559
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+68)) = v561
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+72)) = v563
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+76)) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+80)) = v567
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+84)) = v569
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+88)) = v571
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+92)) = v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+96)) = v575
	v580 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[7])))
	if v580 != int32(1) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v641 != 0 {
		goto L130
	} else {
		goto L131
	}
L117:
	;
	v584 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[8]))
	if base.F64_ge(v584, float64(0)) == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v589 = *(*float64)(unsafe.Add(mBase, uint32(v480)+16))
	if base.F64_gt(v589, v584) == int32(0) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v593 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v593
	v597 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[9]))
	if base.F64_ge(v597, float64(0)) == int32(0) {
		v609 = v593
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v611 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[10]))
	if base.F64_ge(v611, float64(0)) == int32(0) {
		v623 = v609
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v602 = *(*float64)(unsafe.Add(mBase, uint32(v480)+16))
	if base.F64_gt(v602, v597) == int32(0) {
		v609 = v593
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v606 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v606
	v609 = v606
	goto L120
L123:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[11])))
	if v625 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v616 = *(*float64)(unsafe.Add(mBase, uint32(v480)+16))
	if base.F64_gt(v616, v611) == int32(0) {
		v623 = v609
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v621 = v609 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v621
	v623 = v621
	goto L123
L126:
	;
	v629 = v623 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v629
	v631 = v629
	goto L128
L127:
	;
	v631 = v623
	goto L128
L128:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[12])))
	if v633 != int32(1) {
		goto L116
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v631 | int32(16)
	goto L116
L130:
	;
	F_DestroyPartitionDirectory(m, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	m.G0 = v16 + int32(16)
	return v521
L133:
	;
	goto L132
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v508))) = v650
	v488 = v488 + int32(1)
	goto L103
}
