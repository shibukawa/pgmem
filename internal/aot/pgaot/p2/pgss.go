package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int64
	_ = v87
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
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
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	var v361 int64
	_ = v361
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v386 int64
	_ = v386
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v406 int32
	_ = v406
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v548 int32
	_ = v548
	var v567 int32
	_ = v567
	var v568 int64
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int64
	_ = v584
	var v585 int64
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	v6 = int32(0)
	v21 = int64(0)
	v26 = m.G0
	v28 = v26 - int32(752)
	m.G0 = v28
	v31 = l0 + int32(16)
	v39 = v6
	v40 = v6
	v41 = v6
	v42 = v6
	v43 = v6
	v44 = v6
	v45 = v6
	v46 = v6
	v47 = int32(-1)
	v53 = v21
	v54 = v21
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
	switch v47 - int32(1) {
	case 0:
		v463 = v39
		v464 = v40
		v465 = v41
		v466 = v42
		goto L9
	case 1:
		v143 = v39
		v144 = v40
		v145 = v43
		v146 = v44
		v147 = v53
		v148 = v54
		goto L11
	default:
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v567 = int32(m.ExcTag)
	v568 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v567 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v465
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v466
	v532 = int32(_a_F_pgss_planner_0)
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v534 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v463
	F_pg_re_throw(m)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L6
	} else {
		goto L52
	}
L8:
	;
	m.G0 = v28 + int32(752)
	return v507
L9:
	;
	if v464 != 0 {
		goto L7
	} else {
		goto L45
	}
L10:
	;
	v445 = int32(_a_F_pgss_planner_0)
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v447 + int32(1)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1]))
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0]))
	goto L41
L11:
	;
	if v144 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[3]))
	if int32(0) <= v61 {
		v443 = v39
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[4]))
	if v65 != int32(2) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v87 == int64(0) {
		v443 = v31
		goto L10
	} else {
		goto L23
	}
L15:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v65 != int32(1)) != 0 {
		v443 = v39
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
		v443 = v39
		goto L10
	} else {
		goto L21
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	if v74 != 0 {
		v443 = v39
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_planner[5])))
	if v76&int32(1) != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v443 = v39
	goto L10
L21:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_planner[5])))
	if v82&int32(1) == int32(0) {
		v443 = v39
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	base.MemoryCopy(m, v28+int32(544), int32(_a_F_pgss_planner_1), int32(128))
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+408)) = v96
	v99 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+400)) = v99
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+392)) = v102
	v105 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+384)) = v105
	v108 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+376)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v31
	v119 = int32(1)
	F___clock_gettime(m, v119, v28+int32(672))
	mBase = m.M
	v123 = int32(_a_F_pgss_planner_0)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v125 + v119
	v129 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+680)))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v28)+672))
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1]))
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0]))
	goto L24
L24:
	;
	v136 = v28 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v28 + int32(12)
	goto L27
L25:
	;
	v143 = v31
	v144 = int32(0)
	v145 = v134
	v146 = v132
	v147 = v129
	v148 = v130
	goto L11
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v28 + int32(176)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[11]))
	if v156 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v145
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v146
	v426 = int32(_a_F_pgss_planner_0)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v428 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	F_pg_re_throw(m)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L40
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v146
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v145
	v186 = int32(_a_F_pgss_planner_0)
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	v189 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v188 - v189
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	F___clock_gettime(m, v189, v28+int32(688))
	mBase = m.M
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v28)+688))
	v206 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+696)))
	v208 = v28 + int32(416)
	base.MemoryFill(m, v208, int32(0), int32(128))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	v222 = v28 + int32(544)
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
	v225 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[12]))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v222)))
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = v223 + (v225 - v226)
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v208)+8))
	v232 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[13]))
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v222)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+8)) = v230 + (v232 - v233)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v208)+16))
	v239 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[14]))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v222)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+16)) = v237 + (v239 - v240)
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v208)+24))
	v246 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[15]))
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v222)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+24)) = v244 + (v246 - v247)
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v208)+32))
	v253 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[16]))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v222)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+32)) = v251 + (v253 - v254)
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v208)+40))
	v260 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[17]))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v222)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+40)) = v258 + (v260 - v261)
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v208)+48))
	v267 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[18]))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+48)) = v265 + (v267 - v268)
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v208)+56))
	v274 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[19]))
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v222)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+56)) = v272 + (v274 - v275)
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v208)+64))
	v281 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[20]))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v222)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+64)) = v279 + (v281 - v282)
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v208)+72))
	v288 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[21]))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v222)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+72)) = v286 + (v288 - v289)
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v208)+80))
	v295 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[22]))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v222)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+80)) = v293 + (v295 - v296)
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v208)+88))
	v302 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[23]))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v222)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+88)) = v300 + (v302 - v303)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v208)+96))
	v309 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[24]))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v222)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+96)) = v307 + (v309 - v310)
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v208)+104))
	v316 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[25]))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v222)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+104)) = v314 + (v316 - v317)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v208)+112))
	v323 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[26]))
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v222)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+112)) = v321 + (v323 - v324)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v208)+120))
	v330 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[27]))
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v222)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+120)) = v328 + (v330 - v331)
	goto L37
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	v166 = m.T0[v156].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	v177 = F_standard_planner(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v179 = v45
	v180 = v166
	v181 = v166
	goto L31
L36:
	;
	v179 = v177
	v180 = v46
	v181 = v177
	goto L31
L37:
	;
	v335 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+368)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v28)+360)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v28)+352)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v28)+344)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v28)+336)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	v355 = v28 + int32(336)
	v357 = v28 + int32(376)
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v355)+16))
	v360 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[8]))
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v357)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v355)+16)) = v358 + (v360 - v361)
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v355)))
	v367 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[10]))
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v357)))
	*(*int64)(unsafe.Add(mBase, uint32(v355))) = v365 + (v367 - v368)
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v355)+8))
	v374 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[9]))
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v357)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v355)+8)) = v372 + (v374 - v375)
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v355)+24))
	v381 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[7]))
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v357)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v355)+24)) = v379 + (v381 - v382)
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v355)+32))
	v388 = *(*int64)(unsafe.Add(mBase, _c_F_pgss_planner[6]))
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v357)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v355)+32)) = v386 + (v388 - v389)
	goto L38
L38:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v143
	v406 = int32(0)
	F_pgss_store(m, l1, v396, v395, v394, v406, base.F64_div(base.F64_convert_i64_s(v206-v147+(v205-v148)*int64(1000000000)), float64(1e+06)), int64(0), v208, v355, v406, v406, v406, v406, v393)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v507 = v181
	goto L8
L40:
	;
	goto L3
L41:
	;
	v456 = v28 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = v28 + int32(12)
	goto L44
L42:
	;
	v463 = v443
	v464 = int32(0)
	v465 = v454
	v466 = v452
	goto L9
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v28 + int32(16)
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[11]))
	if v472 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[1])) = v466
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[0])) = v465
	v500 = int32(_a_F_pgss_planner_0)
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_planner[2])) = v502 - int32(1)
	v507 = v495
	goto L8
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v463
	v482 = m.T0[v472].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+708)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v28)+704)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v28)+712)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+716)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+720)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+724)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v28)+728)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v28)+736)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v28)+748)) = v463
	v493 = F_standard_planner(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	v495 = v482
	goto L46
L51:
	;
	v495 = v493
	goto L46
L52:
	;
	goto L5
L53:
	;
	v572 = int32(v568)
	m.G0 = v28
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	if v28+int32(12) == v578 {
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
	if v582 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	v582 = v580
	goto L58
L57:
	;
	v582 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v28)+748))
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v28)+736))
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v28)+728))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v28)+724))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v28)+720))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v28)+716))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v28)+712))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v28)+708))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v28)+704))
	v39 = v583
	v40 = v574
	v41 = v591
	v42 = v590
	v43 = v587
	v44 = v586
	v45 = v589
	v46 = v588
	v47 = v582
	v53 = v585
	v54 = v584
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v575, v574)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
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
