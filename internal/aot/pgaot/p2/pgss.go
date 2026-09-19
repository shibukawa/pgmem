package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pgss_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v247 int64
	_ = v247
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v268 int64
	_ = v268
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v279 int64
	_ = v279
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v328 int64
	_ = v328
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v389 int32
	_ = v389
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
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
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v530 int32
	_ = v530
	var v548 int32
	_ = v548
	var v549 int64
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v566 int64
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	v5 = int32(0)
	v19 = int64(0)
	v24 = m.G0
	v26 = v24 - int32(736)
	m.G0 = v26
	v29 = l0 + int32(16)
	v36 = v5
	v37 = v5
	v38 = v5
	v39 = v5
	v40 = v5
	v41 = v5
	v42 = v5
	v43 = v5
	v44 = int32(-1)
	v49 = v19
	v50 = v19
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	switch v44 - int32(1) {
	case 0:
		v446 = v36
		v447 = v37
		v448 = v38
		v449 = v39
		goto L9
	case 1:
		v136 = v36
		v137 = v37
		v138 = v40
		v139 = v41
		v140 = v49
		v141 = v50
		goto L11
	default:
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v548 = int32(m.ExcTag)
	v549 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v548 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v448
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v449
	v514 = int32(_a_F_pgss_planner_0)
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v516 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v446
	F_pg_re_throw(m)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L6
	} else {
		goto L52
	}
L8:
	;
	m.G0 = v26 + int32(736)
	return v490
L9:
	;
	if v447 != 0 {
		goto L7
	} else {
		goto L45
	}
L10:
	;
	v428 = int32(_a_F_pgss_planner_0)
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v430 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1]))
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0]))
	goto L41
L11:
	;
	if v137 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[3]))
	if int32(0) <= v57 {
		v426 = v36
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[4]))
	if v61 != int32(2) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v83 == int64(0) {
		v426 = v29
		goto L10
	} else {
		goto L23
	}
L15:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v61 != int32(1)) != 0 {
		v426 = v36
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l1 == int32(0) {
		v426 = v36
		goto L10
	} else {
		goto L21
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	if v70 != 0 {
		v426 = v36
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_planner[5])))
	if v72&int32(1) != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v426 = v36
	goto L10
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_planner[5])))
	if v78&int32(1) == int32(0) {
		v426 = v36
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	base.MemoryCopy(m, v26+int32(528), int32(_a_F_pgss_planner_1), int32(128))
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+392)) = v92
	v95 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+384)) = v95
	v98 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+376)) = v98
	v101 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+368)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v29
	v112 = int32(1)
	F___clock_gettime(m, v112, v26+int32(656))
	mBase = m.M
	v116 = int32(_a_F_pgss_planner_0)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v118 + v112
	v122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+664)))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v26)+656))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1]))
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0]))
	goto L24
L24:
	;
	v129 = v26 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v26 + int32(12)
	goto L27
L25:
	;
	v136 = v29
	v137 = int32(0)
	v138 = v127
	v139 = v125
	v140 = v122
	v141 = v123
	goto L11
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v26 + int32(176)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[10]))
	if v149 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v138
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v139
	v409 = int32(_a_F_pgss_planner_0)
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v411 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	F_pg_re_throw(m)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L6
	} else {
		goto L40
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v139
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v138
	v179 = int32(_a_F_pgss_planner_0)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	v182 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v181 - v182
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	F___clock_gettime(m, v182, v26+int32(672))
	mBase = m.M
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v26)+672))
	v199 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+680)))
	v201 = v26 + int32(400)
	base.MemoryFill(m, v201, int32(0), int32(128))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	v215 = v26 + int32(528)
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	v218 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[11]))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v215)))
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = v216 + (v218 - v219)
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v201)+8))
	v225 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[12]))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v215)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+8)) = v223 + (v225 - v226)
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v201)+16))
	v232 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[13]))
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v215)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+16)) = v230 + (v232 - v233)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v201)+24))
	v239 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[14]))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v215)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+24)) = v237 + (v239 - v240)
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v201)+32))
	v246 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[15]))
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v215)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+32)) = v244 + (v246 - v247)
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v201)+40))
	v253 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[16]))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v215)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+40)) = v251 + (v253 - v254)
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v201)+48))
	v260 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[17]))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v215)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+48)) = v258 + (v260 - v261)
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v201)+56))
	v267 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[18]))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v215)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+56)) = v265 + (v267 - v268)
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v201)+64))
	v274 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[19]))
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v215)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+64)) = v272 + (v274 - v275)
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v201)+72))
	v281 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[20]))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v215)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+72)) = v279 + (v281 - v282)
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v201)+80))
	v288 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[21]))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v215)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+80)) = v286 + (v288 - v289)
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v201)+88))
	v295 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[22]))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v215)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+88)) = v293 + (v295 - v296)
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v201)+96))
	v302 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[23]))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v215)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+96)) = v300 + (v302 - v303)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v201)+104))
	v309 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[24]))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v215)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+104)) = v307 + (v309 - v310)
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v201)+112))
	v316 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[25]))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v215)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+112)) = v314 + (v316 - v317)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v201)+120))
	v323 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[26]))
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v215)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+120)) = v321 + (v323 - v324)
	goto L37
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	v159 = m.T0[v149].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	v170 = F_standard_planner(m, l0, l2, l3)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v172 = v42
	v173 = v159
	v174 = v159
	goto L31
L36:
	;
	v172 = v170
	v173 = v43
	v174 = v170
	goto L31
L37:
	;
	v328 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+360)) = v328
	*(*int64)(unsafe.Add(mBase, uint32(v26)+352)) = v328
	*(*int64)(unsafe.Add(mBase, uint32(v26)+344)) = v328
	*(*int64)(unsafe.Add(mBase, uint32(v26)+336)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	v346 = v26 + int32(336)
	v348 = v26 + int32(368)
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v346)+16))
	v351 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[7]))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v348)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v346)+16)) = v349 + (v351 - v352)
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v346)))
	v358 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[9]))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v348)))
	*(*int64)(unsafe.Add(mBase, uint32(v346))) = v356 + (v358 - v359)
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v346)+8))
	v365 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[8]))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v348)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v346)+8)) = v363 + (v365 - v366)
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v346)+24))
	v372 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[6]))
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v348)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v346)+24)) = v370 + (v372 - v373)
	goto L38
L38:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v136
	v389 = int32(0)
	F_pgss_store(m, l1, v379, v378, v377, v389, base.F64_div(base.F64_convert_i64_s(v199-v140+(v198-v141)*int64(1000000000)), float64(1e+06)), int64(0), v201, v346, v389, v389, v389, v389)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v490 = v174
	goto L8
L40:
	;
	goto L3
L41:
	;
	v439 = v26 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v439)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v26 + int32(12)
	goto L44
L42:
	;
	v446 = v426
	v447 = int32(0)
	v448 = v437
	v449 = v435
	goto L9
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v26 + int32(16)
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[10]))
	if v455 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v449
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v448
	v483 = int32(_a_F_pgss_planner_0)
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v485 - int32(1)
	v490 = v478
	goto L8
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v446
	v465 = m.T0[v455].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+692)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v26)+688)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v26)+696)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+700)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+704)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+708)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v26)+712)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v26)+720)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v26)+732)) = v446
	v476 = F_standard_planner(m, l0, l2, l3)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	v478 = v465
	goto L46
L51:
	;
	v478 = v476
	goto L46
L52:
	;
	goto L5
L53:
	;
	v553 = int32(v549)
	m.G0 = v26
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	if v26+int32(12) == v559 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	m.ExcPending = 1
	goto L62
L55:
	;
	if v563 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	v563 = v561
	goto L58
L57:
	;
	v563 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v26)+732))
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v26)+720))
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v26)+712))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v26)+708))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v26)+704))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v26)+700))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v26)+696))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v26)+692))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v26)+688))
	v36 = v564
	v37 = v555
	v38 = v572
	v39 = v571
	v40 = v568
	v41 = v567
	v42 = v570
	v43 = v569
	v44 = v563
	v49 = v566
	v50 = v565
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v556, v555)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return int32(0)
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgss_shmem_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v166 int32
	_ = v166
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
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
	var v517 int32
	_ = v517
	v1 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(544)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[0]))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v18].(func(*base.Module))(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[1])) = v22
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2])) = v22
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[3]))
	v32 = F_LWLockAcquire(m, v28+int32(2688), v22)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v39 = F_ShmemInitStruct(m, int32(_a_F_pgss_shmem_startup_0), int32(56), v15+int32(543))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2])) = v39
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+543)))
	if v42 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[4]))
	if int32(0) < v46 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+508)) = int64(1855425871896)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[5]))
	v189 = F_ShmemInitHash(m, int32(_a_F_pgss_shmem_startup_1), v185, v185, v15+int32(492), int32(40))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L30
	}
L11:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[3]))
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+16)) = int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(v132)+8)) = int64(4621819117588971520)
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v127 + v53<<(uint(int32(7))%32)
	v138 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v132)+20)), uint32(v138))
	v141 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v132)+40)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v132)+32)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v132)+24)) = v141
	v150 = m.G0
	v151 = int32(16)
	v152 = v150 - v151
	m.G0 = v152
	F_gettimeofday(m, v152)
	mBase = m.M
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
	v156 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+8)))
	m.G0 = v152 + v151
	goto L29
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[6]))
	v53 = int32(214)
	v54 = v1
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L26
	}
L15:
	;
	v66 = v50 + v54*int32(68)
	v67 = int32(_a_F_pgss_shmem_startup_0)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[7])))
	if base.B2i32(v70 == int32(0))|base.B2i32(v70 != v73) != 0 {
		v91 = v70
		v92 = v73
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	if v91-v92 == int32(0) {
		goto L11
	} else {
		goto L24
	}
L18:
	;
	goto L17
L19:
	;
	v76 = v66
	v77 = v67
	goto L20
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v81 == int32(0) {
		v91 = v81
		v92 = v80
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v91 = v81
	v92 = v80
	goto L18
L22:
	;
	v84 = int32(1)
	if v81 == v80 {
		v76 = v76 + v84
		v77 = v77 + v84
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v66)+64))
	v99 = v54 + int32(1)
	if v99 != v46 {
		v53 = v96 + v53
		v54 = v99
		goto L15
	} else {
		goto L25
	}
L25:
	;
	goto L16
L26:
	;
	F_errmsg_internal(m, int32(_a_F_pgss_shmem_startup_2), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_startup_3), int32(605), int32(_a_F_pgss_shmem_startup_4))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v166)+48)) = v156 + v155*int64(1000000) - int64(946684800000000)
	goto L10
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[1])) = v189
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[3]))
	F_LWLockRelease(m, v193+int32(2688))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[8])))
	if v199 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_on_shmem_exit(m, int32(_a_F_pgss_shmem_startup_5), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+543)))
	if v206 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	m.G0 = v15 + int32(544)
	return
L37:
	;
	v207 = int32(_a_F_pgss_shmem_startup_6)
	v208 = F_unlink(m, v207)
	mBase = m.M
	v209 = int32(0)
	v214 = F_AllocateFile(m, v207, int32(_a_F_pgss_shmem_startup_7))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L41
	}
L38:
	;
	if v501 != 0 {
		goto L106
	} else {
		goto L107
	}
L39:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_startup_8), v494, int32(_a_F_pgss_shmem_startup_9))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L105
	}
L40:
	;
	v470 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L101
	}
L41:
	;
	if v214 == int32(0) {
		v457 = v209
		v459 = v209
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[9])))
	if v219 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v222 = F_FreeFile(m, v214)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v226 = F_AllocateFile(m, int32(_a_F_pgss_shmem_startup_10), int32(_a_F_pgss_shmem_startup_11))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L49
	}
L46:
	;
	goto L36
L47:
	;
	v441 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L97
	}
L48:
	;
	v413 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L93
	}
L49:
	;
	if v226 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[10]))
	if v231 != int32(44) {
		v402 = v209
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v237 = F_palloc(m, int32(2048))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	v234 = F_FreeFile(m, v214)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L36
L55:
	;
	v243 = F_fread(m, v15+int32(488), int32(4), int32(1), v226)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if v243 != int32(1) {
		v402 = v237
		goto L48
	} else {
		goto L57
	}
L57:
	;
	v251 = F_fread(m, v15+int32(480), int32(4), int32(1), v226)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if v251 != int32(1) {
		v402 = v237
		goto L48
	} else {
		goto L59
	}
L59:
	;
	v259 = F_fread(m, v15+int32(484), int32(4), int32(1), v226)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v259 != int32(1) {
		v402 = v237
		goto L48
	} else {
		goto L61
	}
L61:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v15)+488))
	if v263 != int32(539100168) {
		v430 = v237
		goto L47
	} else {
		goto L62
	}
L62:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v15)+480))
	if v266 != int32(1800) {
		v430 = v237
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v15)+484))
	if int32(0) < v269 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v278 = v237
	v280 = int32(2048)
	v285 = v1
	goto L67
L65:
	;
	v372 = v237
	goto L66
L66:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2]))
	v387 = F_fread(m, v382+int32(40), int32(16), int32(1), v226)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L88
	}
L67:
	;
	v291 = F_fread(m, v15+int32(48), int32(432), int32(1), v226)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	v372 = v310
	goto L66
L69:
	;
	if v291 != int32(1) {
		v402 = v278
		goto L48
	} else {
		goto L70
	}
L70:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	if base.Ui32(int32(35)) <= base.Ui32(v295) {
		v430 = v278
		goto L47
	} else {
		goto L71
	}
L71:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if v280 <= v298 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v300 = int32(1)
	v301 = v280 << (uint(v300) % 32)
	v303 = v298 + v300
	if v303 < v301 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v309 = v298
	v310 = v278
	v311 = v280
	goto L74
L74:
	;
	v312 = int32(1)
	v315 = F_fread(m, v310, v312, v309+v312, v226)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L79
	}
L75:
	;
	v305 = v301
	goto L77
L76:
	;
	v305 = v303
	goto L77
L77:
	;
	v306 = F_repalloc(m, v278, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	v309 = v308
	v310 = v306
	v311 = v305
	goto L74
L79:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if v315 != v317+int32(1) {
		v402 = v310
		goto L48
	} else {
		goto L80
	}
L80:
	;
	v322 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v317+v310))) = uint8(v322)
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v15)+72))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
	if v324 != int64(0)-v326 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2]))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+24))
	v332 = int32(1)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	v336 = F_fwrite(m, v310, v332, v333+v332, v214)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v366 = v285 + int32(1)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v15)+484))
	if v366 < v367 {
		v278 = v310
		v280 = v311
		v285 = v366
		goto L67
	} else {
		goto L87
	}
L84:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if v336 != v338+int32(1) {
		v457 = v226
		v459 = v310
		goto L40
	} else {
		goto L85
	}
L85:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_startup[2]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+24)) = v344 + v336
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	v351 = F_entry_alloc(m, v15+int32(48), v331, v338, v349, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.MemoryCopy(m, v351+int32(24), v15+int32(72), int32(368))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int64)(unsafe.Add(mBase, uint32(v351)+408)) = v357
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v351)+416)) = v359
	goto L83
L87:
	;
	goto L68
L88:
	;
	if v387 != int32(1) {
		v402 = v372
		goto L48
	} else {
		goto L89
	}
L89:
	;
	F_pfree(m, v372)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v393 = F_FreeFile(m, v226)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v395 = F_FreeFile(m, v214)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v398 = F_unlink(m, int32(_a_F_pgss_shmem_startup_10))
	mBase = m.M
	goto L36
L93:
	;
	if v413 == int32(0) {
		v499 = v226
		v501 = v402
		goto L38
	} else {
		goto L94
	}
L94:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_pgss_shmem_startup_10)
	F_errmsg(m, int32(_a_F_pgss_shmem_startup_12), v15+int32(16))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v483 = v226
	v485 = v402
	v494 = int32(701)
	goto L39
L97:
	;
	if v441 == int32(0) {
		v499 = v226
		v501 = v430
		goto L38
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(_a_F_pgss_shmem_startup_10)
	F_errmsg(m, int32(_a_F_pgss_shmem_startup_13), v15+int32(32))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	v483 = v226
	v485 = v430
	v494 = int32(707)
	goto L39
L101:
	;
	if v470 == int32(0) {
		v499 = v457
		v501 = v459
		goto L38
	} else {
		goto L102
	}
L102:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_pgss_shmem_startup_6)
	F_errmsg(m, int32(_a_F_pgss_shmem_startup_14), v15)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v483 = v457
	v485 = v459
	v494 = int32(713)
	goto L39
L105:
	;
	v499 = v483
	v501 = v485
	goto L38
L106:
	;
	F_pfree(m, v501)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v499 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v512 = F_FreeFile(m, v499)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if v214 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v514 = F_FreeFile(m, v214)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v517 = F_unlink(m, int32(_a_F_pgss_shmem_startup_10))
	mBase = m.M
	goto L36
L117:
	;
	goto L116
}
