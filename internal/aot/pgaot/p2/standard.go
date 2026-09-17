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
	var v30 int32
	_ = v30
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
	v30 = int32(0)
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32))))
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
	v60 = v30 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v60 < v61 {
		v30 = v60
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 float64
	_ = v92
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v174 float64
	_ = v174
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 float64
	_ = v261
	var v263 float64
	_ = v263
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v270 float64
	_ = v270
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v293 float64
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 float64
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v329 float64
	_ = v329
	var v332 int32
	_ = v332
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v336 float64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 float64
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v386 float64
	_ = v386
	var v389 int32
	_ = v389
	var v391 float64
	_ = v391
	var v392 float64
	_ = v392
	var v395 float64
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int64
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 float64
	_ = v569
	var v574 float64
	_ = v574
	var v578 int32
	_ = v578
	var v582 float64
	_ = v582
	var v587 float64
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 float64
	_ = v596
	var v601 float64
	_ = v601
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
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
	base.MemoryFill(m, v19+int32(8), v23, int32(74))
	if l1&int32(2048) == v23 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	v83 = int32(0)
	v85 = v80 & base.B2i32(v82 != v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+83)) = uint8(v85)
	if l1&int32(256) == v83 {
		v100 = v12
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+82)) = uint8(v75)
	v77 = int32(117)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+84)) = uint8(v77)
	v80 = int32(0)
	goto L3
L5:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[1])))
	if v38&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 != int32(1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	if v46 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[2]))
	if v48 <= int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[3]))
	if int32(0) <= v52 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v55 = m.G0
	v57 = v55 - int32(16)
	m.G0 = v57
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = int32(0)
	v61 = int32(_a_F_standard_planner_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+8)) = uint16(v61)
	v65 = F_max_parallel_hazard_walker(m, l0, v57+int32(8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57)+8)))
	m.G0 = v57 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+84)) = uint8(v67)
	v73 = base.B2i32(v67 != int32(117))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+82)) = uint8(v73)
	v80 = v73
	goto L3
L12:
	;
	v102 = int32(0)
	v105 = F_subquery_planner(m, v19, l0, v102, v102, v100, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v92 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[4]))
	if base.F64_ge(v92, float64(1)) != 0 {
		v100 = v12
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.F64_le(v92, float64(0)) == int32(0) {
		v100 = v92
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v100 = float64(1e-10)
	goto L12
L16:
	;
	v109 = F_fetch_upper_rel(m, v105, int32(7), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
	if base.F64_le(v100, float64(0)) != 0 {
		v203 = v111
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v211 = F_create_plan(m, v105, v203)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L49
	}
L19:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
	if v114 == int32(0) {
		v203 = v111
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		v203 = v111
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v124 = int32(0)
	v126 = v111
	v131 = v117
	goto L22
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v124<<(uint(int32(2))%32))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v139 != 0 {
		v193 = v126
		v194 = v131
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v203 = v193
	goto L18
L24:
	;
	v196 = v124 + int32(1)
	if v196 < v194 {
		v124 = v196
		v126 = v193
		v131 = v194
		goto L22
	} else {
		goto L48
	}
L25:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
	if v138 == v140 {
		v193 = v126
		v194 = v131
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v126)+40))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)+40))
	if v145 != v146 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v188 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L28:
	;
	if v145 < v146 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if base.F64_le(v100, float64(0))|base.F64_ge(v100, float64(1)) != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v151 = int32(-1)
	goto L33
L32:
	;
	v151 = int32(1)
	goto L33
L33:
	;
	v188 = v151
	goto L27
L34:
	;
	v188 = v184
	goto L27
L35:
	;
	v157 = int32(-1)
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v126)+56))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v138)+56))
	if base.F64_lt(v158, v159) != 0 {
		v184 = v157
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v126)+56))
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v126)+48))
	v174 = base.F64_add(base.F64_mul(v100, base.F64_sub(v170, v171)), v171)
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v138)+56))
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v138)+48))
	v179 = base.F64_add(base.F64_mul(v100, base.F64_sub(v175, v176)), v176)
	if base.F64_lt(v174, v179) != 0 {
		v184 = int32(-1)
		goto L34
	} else {
		goto L44
	}
L38:
	;
	if base.F64_gt(v158, v159) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v188 = int32(1)
	goto L27
L40:
	;
	goto L41
L41:
	;
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v126)+48))
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v138)+48))
	if base.F64_lt(v163, v164) != 0 {
		v184 = v157
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if base.F64_gt(v163, v164) != 0 {
		v184 = int32(1)
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v188 = int32(0)
	goto L27
L44:
	;
	v184 = base.F64_lt(v179, v174)
	goto L34
L45:
	;
	v191 = v126
	goto L47
L46:
	;
	v191 = v138
	goto L47
L47:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v193 = v191
	v194 = v192
	goto L24
L48:
	;
	goto L23
L49:
	;
	if l1&int32(2) == int32(0) {
		v221 = v211
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	if v223 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v217 = F_ExecSupportsBackwardScan(m, v211)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v217 != 0 {
		v221 = v211
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v219 = F_materialize_finished_plan(m, v211)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v221 = v219
	goto L50
L55:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	if v404 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L56:
	;
	v401 = v221
	goto L55
L57:
	;
	goto L58
L58:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)))
	if v226 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v401 = v221
	goto L55
L60:
	;
	goto L61
L61:
	;
	if v223 != int32(2) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v235 = F_palloc0(m, int32(88))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+60))
	if v231 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v401 = v221
	goto L55
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = int32(368)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v221)+44))
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v235)+80)) = uint8(v240)
	*(*int32)(unsafe.Add(mBase, uint32(v235)+72)) = v240
	v244 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v235)+56)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v235)+52)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v235)+48)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v235)+44)) = v239
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v235)+81)) = uint8(base.B2i32(v251 == int32(2)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v221)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v235)+60)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v221)+60)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v235)+76)) = int32(-1)
	v261 = *(*float64)(unsafe.Add(mBase, uint32(v221)+8))
	v263 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[5]))
	*(*float64)(unsafe.Add(mBase, uint32(v235)+8)) = base.F64_add(v261, v263)
	v267 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[6]))
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v221)+24))
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v221)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v235)+16)) = base.F64_add(base.F64_mul(v267, v268), base.F64_add(v263, v270))
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v221)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v235)+24)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v221)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v235)+36)) = uint16(v244)
	*(*int32)(unsafe.Add(mBase, uint32(v235)+32)) = v276
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v235)+60))
	v293 = float64(0)
	if v280 == v244 {
		v380 = v244
		v386 = v293
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v391 = *(*float64)(unsafe.Add(mBase, uint32(v221)+8))
	v392 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = base.F64_sub(v391, v392)
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v221)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v221)+16)) = base.F64_sub(v395, v392)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v398)+83)) = uint8(v399)
	v401 = v235
	goto L55
L67:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = v386
	v389 = v380 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(7)))) = uint8(v389)
	goto L66
L68:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v296 <= int32(0) {
		v380 = v244
		v386 = v293
		goto L67
	} else {
		goto L69
	}
L69:
	;
	if v296 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v353<<(uint(int32(2))%32))))
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v366)+64))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+38)))
	v380 = v371 ^ int32(1) | v355
	v386 = base.F64_add(v361, base.F64_add(v367, v368))
	goto L67
L71:
	;
	v353 = int32(0)
	v355 = v244
	v361 = v293
	goto L70
L72:
	;
	goto L73
L73:
	;
	v302 = int32(0)
	if v302 < v296 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v305 = v296
	goto L76
L75:
	;
	v305 = v302
	goto L76
L76:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v315 = int32(0)
	v317 = v244
	v322 = v244
	v323 = v293
	goto L77
L77:
	;
	v324 = int32(2)
	v326 = v310 + v315<<(uint(v324)%32)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v327)+56))
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v327)+64))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v332)+56))
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v332)+64))
	v336 = base.F64_add(base.F64_add(v323, base.F64_add(v328, v329)), base.F64_add(v333, v334))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+38)))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+38)))
	v342 = base.B2i32(v337&v338 == int32(0)) | v317
	v344 = v315 + v324
	v346 = v322 + v324
	if v346 != v305&int32(2147483646) {
		v315 = v344
		v317 = v342
		v322 = v346
		v323 = v336
		goto L77
	} else {
		goto L79
	}
L78:
	;
	if v305&int32(1) == int32(0) {
		v380 = v342
		v386 = v336
		goto L67
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v353 = v344
	v355 = v342
	v361 = v336
	goto L70
L81:
	;
	v467 = F_set_plan_references(m, v105, v401)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L95
	}
L82:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v413 = int32(0)
	goto L83
L83:
	;
	v423 = int32(0)
	if v408 == v423 {
		v433 = v423
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v407 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if v427 <= v413 {
		v433 = int32(0)
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v408)+12))
	v433 = v429 + v413<<(uint(int32(2))%32)
	goto L85
L88:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v441+v413<<(uint(int32(2))%32))))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	F_SS_finalize_plan(m, v448, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L94
	}
L89:
	;
	F_SS_finalize_plan(m, v105, v401)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L93
	}
L90:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if base.B2i32(v433 == int32(0))|base.B2i32(v438 <= v413) != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	if v441 != 0 {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	goto L81
L94:
	;
	v413 = v413 + int32(1)
	goto L83
L95:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v475 = int32(0)
	goto L96
L96:
	;
	v485 = int32(0)
	if v470 == v485 {
		v495 = v485
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v469 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v489 <= v475 {
		v495 = int32(0)
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v495 = v491 + v475<<(uint(int32(2))%32)
	goto L98
L101:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v503+v475<<(uint(int32(2))%32))))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v638 = F_set_plan_references(m, v636, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L126
	}
L102:
	;
	v506 = F_palloc0(m, int32(104))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L106
	}
L103:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if base.B2i32(v495 == int32(0))|base.B2i32(v500 <= v475) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v469)+12))
	if v503 != 0 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = int32(330)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+4)) = v510
	v512 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+8)) = v512
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+24)) = uint8(base.B2i32(v514 != int32(0)))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+25)) = uint8(v518)
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+26)) = uint8(v520)
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+80)))
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+27)) = uint8(v522)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+81)))
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+28)) = uint8(v524)
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+83)))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+36)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+29)) = uint8(v526)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+40)) = v529
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+44)) = v531
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v535 = F_bms_difference(m, v533, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+48)) = v535
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+52)) = v538
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+56)) = v540
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+60)) = v542
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+64)) = v544
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+68)) = v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+72)) = v548
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+76)) = v550
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+80)) = v552
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+84)) = v554
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+88)) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+92)) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+96)) = v560
	v565 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[7])))
	if v565 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v626 != 0 {
		goto L122
	} else {
		goto L123
	}
L109:
	;
	v569 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[8]))
	if base.F64_ge(v569, float64(0)) == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v574 = *(*float64)(unsafe.Add(mBase, uint32(v467)+16))
	if base.F64_gt(v574, v569) == int32(0) {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v578 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = v578
	v582 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[9]))
	if base.F64_ge(v582, float64(0)) == int32(0) {
		v594 = v578
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v596 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[10]))
	if base.F64_ge(v596, float64(0)) == int32(0) {
		v608 = v594
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v467)+16))
	if base.F64_gt(v587, v582) == int32(0) {
		v594 = v578
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v591 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = v591
	v594 = v591
	goto L112
L115:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[11])))
	if v610 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v601 = *(*float64)(unsafe.Add(mBase, uint32(v467)+16))
	if base.F64_gt(v601, v596) == int32(0) {
		v608 = v594
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v606 = v594 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = v606
	v608 = v606
	goto L115
L118:
	;
	v614 = v608 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = v614
	v616 = v614
	goto L120
L119:
	;
	v616 = v608
	goto L120
L120:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[12])))
	if v618 != int32(1) {
		goto L108
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = v616 | int32(16)
	goto L108
L122:
	;
	F_DestroyPartitionDirectory(m, v626)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	m.G0 = v16 + int32(16)
	return v506
L125:
	;
	goto L124
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v638
	v475 = v475 + int32(1)
	goto L96
}
