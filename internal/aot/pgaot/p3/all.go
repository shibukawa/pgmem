package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAllSchemaPublicationRelations(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = F_GetPublicationSchemas(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v13 <= v12 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L8
L8:
	;
	v18 = v12
	v20 = int32(0)
	goto L9
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v22 = int32(2)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v18<<(uint(v22)%32))))
	v27 = F_GetSchemaPublicationRelations(m, v25, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	return v29
L11:
	;
	v29 = F_list_concat(m, v20, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = v18 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v32 < v33 {
		v18 = v32
		v20 = v29
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
}
func F_ReleaseAllPlanCacheRefsInOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v13 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L26
	}
L5:
	;
	v20 = l0 + int32(24)
	v22 = v18
	v23 = int32(0)
	v24 = v18
	goto L8
L6:
	;
	goto L7
L7:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
	if v72 != 0 {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v31 = v20 + v23<<(uint(int32(4))%32)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v32 == int32(_a_F_ReleaseAllPlanCacheRefsInOwner_0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v39 = v20 + v22<<(uint(int32(4))%32) - int32(16)
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v40
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v43
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v47 = v45 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v47)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAllPlanCacheRefsInOwner[0]))
	m.T0[v50].(func(*base.Module, int64))(m, v42)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v57 = v23
	v58 = v24
	goto L12
L12:
	;
	v60 = v57 + int32(1)
	v62 = v58 & int32(255)
	if v60 < v62 {
		v22 = v62
		v23 = v60
		v24 = v58
		goto L8
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v57 = v23 - int32(1)
	v58 = v55
	goto L12
L15:
	;
	goto L9
L16:
	;
	v75 = v72
	v76 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v117)
	m.G0 = v11 + int32(16)
	goto L1
L19:
	;
	v83 = v76 << (uint(int32(4)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	v85 = v83 + v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v86 == int32(_a_F_ReleaseAllPlanCacheRefsInOwner_0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	*(*int32)(unsafe.Add(mBase, uint32(v92+v83)+8)) = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v96 - int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAllPlanCacheRefsInOwner[0]))
	m.T0[v101].(func(*base.Module, int64))(m, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	v105 = v75
	goto L23
L23:
	;
	v107 = v76 + int32(1)
	if base.Ui32(v107) < base.Ui32(v105) {
		v75 = v105
		v76 = v107
		goto L19
	} else {
		goto L25
	}
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
	v105 = v104
	goto L23
L25:
	;
	goto L20
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAllPlanCacheRefsInOwner[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127
	F_errmsg_internal(m, int32(_a_F_ReleaseAllPlanCacheRefsInOwner_1), v11)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ReleaseAllPlanCacheRefsInOwner_2), int32(829), int32(_a_F_ReleaseAllPlanCacheRefsInOwner_3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_disable_all_timeouts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_disable_all_timeouts[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_disable_all_timeouts[1])) = v2
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[2])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[3])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[4])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[5])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[6])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[7])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[8])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[9])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[10])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[11])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[12])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[13])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[14])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[15])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[16])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[17])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[18])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[19])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[20])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[21])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[22])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[23])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[24])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[25])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[26])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[27])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[28])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[29])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[30])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[31])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[32])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[33])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[34])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[35])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[36])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[37])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[38])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[39])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[40])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[41])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[42])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[43])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[44])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[45])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[46])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, _c_F_disable_all_timeouts[47])) = uint8(v2)
	return
}
func F_show_all_settings(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int64
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int32
	_ = v289
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
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
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 float64
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 float64
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 float64
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 float64
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
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
	var v498 int64
	_ = v498
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v545 int32
	_ = v545
	var v546 int64
	_ = v546
	var v550 int32
	_ = v550
	var v561 int64
	_ = v561
	v10 = m.G0
	v12 = v10 - int32(480)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	goto L47
L4:
	;
	return int64(0)
L5:
	;
	v22 = int32(_a_F_show_all_settings_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[0])) = v25
	v28 = F_CreateTemplateTupleDesc(m, int32(17))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v28, int32(1), int32(_a_F_show_all_settings_1), int32(25), int32(-1), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v28, int32(2), int32(_a_F_show_all_settings_2), int32(25), int32(-1), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v28, int32(3), int32(_a_F_show_all_settings_3), int32(25), int32(-1), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v28, int32(4), int32(_a_F_show_all_settings_4), int32(25), int32(-1), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v28, int32(5), int32(_a_F_show_all_settings_5), int32(25), int32(-1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v28, int32(6), int32(_a_F_show_all_settings_6), int32(25), int32(-1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v28, int32(7), int32(_a_F_show_all_settings_7), int32(25), int32(-1), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v28, int32(8), int32(_a_F_show_all_settings_8), int32(25), int32(-1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v28, int32(9), int32(_a_F_show_all_settings_9), int32(25), int32(-1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v28, int32(10), int32(_a_F_show_all_settings_10), int32(25), int32(-1), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v28, int32(11), int32(_a_F_show_all_settings_11), int32(25), int32(-1), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v28, int32(12), int32(_a_F_show_all_settings_12), int32(1009), int32(-1), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v28, int32(13), int32(_a_F_show_all_settings_13), int32(25), int32(-1), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_TupleDescInitEntry(m, v28, int32(14), int32(_a_F_show_all_settings_14), int32(25), int32(-1), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v28, int32(15), int32(_a_F_show_all_settings_15), int32(25), int32(-1), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v28, int32(16), int32(_a_F_show_all_settings_16), int32(23), int32(-1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_TupleDescInitEntry(m, v28, int32(17), int32(_a_F_show_all_settings_17), int32(16), int32(-1), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v149 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v149 < v158 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v236 = F_TupleDescGetAttInMetadata(m, v28)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L43
	}
L25:
	;
	v162 = v28 + int32(28)
	v169 = v149
	v170 = v158
	v172 = v149
	goto L29
L26:
	;
	v226 = v149
	v233 = v158
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v226
	goto L24
L28:
	;
	v226 = v220
	v233 = v199
	goto L27
L29:
	;
	v178 = v162 + v158<<(uint(int32(3))%32) + v169*int32(100)
	v181 = v162 + v169<<(uint(int32(3))%32)
	if v158 != v170 {
		v199 = v170
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v220 = v158
	goto L28
L31:
	;
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v181)+2)))
	if v200 <= int32(0) {
		v220 = v169
		goto L28
	} else {
		goto L39
	}
L32:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+7)))
	if v183 != int32(118) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v199 = v169
	goto L31
L34:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+4)))
	if v186 != int32(1) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+6)))
	if v189&int32(6) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v181)+2)))
	if v192 <= int32(0) {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+90)))
	if v195 != int32(118) {
		v199 = v158
		goto L31
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+90)))
	if v203 == int32(118) {
		v220 = v169
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+5)))
	v212 = (v172 + v206 - int32(1)) & (int32(0) - v206)
	if int32(_a_F_show_all_settings_18) < v212 {
		v220 = v169
		goto L28
	} else {
		goto L41
	}
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v181))) = uint16(v212)
	v218 = v169 + int32(1)
	if v218 != v158 {
		v169 = v218
		v170 = v199
		v172 = v212 + v200
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v236
	v241 = F_get_guc_variables(m, v12+int32(220))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v241
	v244 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+220)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v244
	*(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[0])) = v23
	goto L3
L45:
	;
	m.G0 = v12 + int32(480)
	return v561
L46:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v308
	v311 = F_ShowGUCOption(m, v271, int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L61
	}
L47:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	v254 = base.I32_wrap_i64(v253)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	if v254 < v255 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	v261 = v254
	v267 = v253
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L60
	}
L51:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258+v261<<(uint(int32(2))%32))))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	if v272&int32(4) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v285 = v267
	goto L55
L54:
	;
	if v272&int32(1024) == int32(0) {
		goto L46
	} else {
		goto L56
	}
L55:
	;
	v287 = v285 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v287
	v289 = base.I32_wrap_i64(v287)
	if v289 < v255 {
		v261 = v289
		v267 = v287
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[1]))
	v282 = F_has_privs_of_role(m, v280, int32(3374))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if v282 != 0 {
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	v285 = v284
	goto L55
L59:
	;
	goto L52
L60:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+20)) = int32(2)
	v305 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v305)
	v561 = int64(0)
	goto L45
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+148)) = v311
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v315 = F_get_config_unit_name(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v315
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	v319 = int32(2)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v318<<(uint(v319)%32))+uint32(_c_F_show_all_settings[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327<<(uint(v319)%32))+uint32(_c_F_show_all_settings[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+168)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v271)+24))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332<<(uint(v319)%32))+uint32(_c_F_show_all_settings[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v271)+32))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v337<<(uint(v319)%32))+uint32(_c_F_show_all_settings[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v340
	switch v332 {
	case 0:
		goto L69
	case 1:
		goto L68
	case 2:
		goto L67
	case 3:
		goto L66
	case 4:
		goto L65
	default:
		goto L64
	}
L63:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v271)+32))
	if v504 != int32(3) {
		goto L108
	} else {
		goto L109
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = int32(0)
	v498 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+188)) = v498
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = v498
	goto L63
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v481 = F_config_enum_get_options(m, v271+int32(96), int32(_a_F_show_all_settings_19), int32(_a_F_show_all_settings_20), int32(_a_F_show_all_settings_21))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L102
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v271)+100))
	if v460 != 0 {
		goto L94
	} else {
		goto L95
	}
L67:
	;
	v408 = *(*float64)(unsafe.Add(mBase, uint32(v271)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+128)) = v408
	v411 = v12 + int32(224)
	v416 = F_pg_snprintf(m, v411, int32(256), int32(_a_F_show_all_settings_22), v12+int32(128))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L86
	}
L68:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v271)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v360
	v363 = v12 + int32(224)
	v368 = F_pg_snprintf(m, v363, int32(256), int32(_a_F_show_all_settings_23), v12-int32(-64))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L78
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = int64(0)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+100)))
	if v348 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v349 = int32(_a_F_show_all_settings_24)
	goto L72
L71:
	;
	v349 = int32(_a_F_show_all_settings_25)
	goto L72
L72:
	;
	v350 = F_pstrdup(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v350
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+116)))
	if v355 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v356 = int32(_a_F_show_all_settings_24)
	goto L76
L75:
	;
	v356 = int32(_a_F_show_all_settings_25)
	goto L76
L76:
	;
	v357 = F_pstrdup(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v357
	goto L63
L78:
	;
	v370 = F_pstrdup(m, v363)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v370
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v271)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v373
	v379 = F_pg_snprintf(m, v363, int32(256), int32(_a_F_show_all_settings_23), v12+int32(48))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v381 = F_pstrdup(m, v363)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v381
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v271)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v386
	v392 = F_pg_snprintf(m, v363, int32(256), int32(_a_F_show_all_settings_23), v12+int32(32))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v394 = F_pstrdup(m, v363)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v394
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v271)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v397
	v403 = F_pg_snprintf(m, v363, int32(256), int32(_a_F_show_all_settings_23), v12+int32(16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v405 = F_pstrdup(m, v363)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v405
	goto L63
L86:
	;
	v418 = F_pstrdup(m, v411)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v418
	v421 = *(*float64)(unsafe.Add(mBase, uint32(v271)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+112)) = v421
	v427 = F_pg_snprintf(m, v411, int32(256), int32(_a_F_show_all_settings_22), v12+int32(112))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v429 = F_pstrdup(m, v411)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v429
	v434 = *(*float64)(unsafe.Add(mBase, uint32(v271)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+96)) = v434
	v440 = F_pg_snprintf(m, v411, int32(256), int32(_a_F_show_all_settings_22), v12+int32(96))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v442 = F_pstrdup(m, v411)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v442
	v445 = *(*float64)(unsafe.Add(mBase, uint32(v271)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+80)) = v445
	v451 = F_pg_snprintf(m, v411, int32(256), int32(_a_F_show_all_settings_22), v12+int32(80))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v453 = F_pstrdup(m, v411)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v453
	goto L63
L94:
	;
	v461 = F_pstrdup(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	v464 = int32(0)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v464
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v271)+116))
	if v466 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v464 = v461
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = int32(0)
	goto L63
L99:
	;
	goto L100
L100:
	;
	v471 = F_pstrdup(m, v466)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v471
	goto L63
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = v481
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v271)+100))
	v485 = F_config_enum_lookup_by_value(m, v271, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v487 = F_pstrdup(m, v485)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v487
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v271)+120))
	v491 = F_config_enum_lookup_by_value(m, v271, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v493 = F_pstrdup(m, v491)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v493
	goto L63
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v530
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v271)+28))
	if v534&int32(2) != 0 {
		goto L114
	} else {
		goto L115
	}
L108:
	;
	v526 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v526
	v530 = v526
	goto L107
L109:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_show_all_settings[1]))
	v510 = F_has_privs_of_role(m, v508, int32(3374))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	if v510 == int32(0) {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v271)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v271)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v516
	v519 = v12 + int32(224)
	v522 = F_pg_snprintf(m, v519, int32(256), int32(_a_F_show_all_settings_23), v12)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v524 = F_pstrdup(m, v519)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v530 = v524
	goto L107
L114:
	;
	v537 = int32(_a_F_show_all_settings_26)
	goto L116
L115:
	;
	v537 = int32(_a_F_show_all_settings_27)
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v537
	v541 = F_BuildTupleFromCStrings(m, v257, v12+int32(144))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	v544 = F_HeapTupleHeaderGetDatum(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v546 + int64(1)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v550)+20)) = int32(1)
	v561 = v544
	goto L45
}
