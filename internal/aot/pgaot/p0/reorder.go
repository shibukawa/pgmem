package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferCleanupSerializedTXNs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v5 = m.G0
	v7 = v5 - int32(2208)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_ReorderBufferCleanupSerializedTXNs_0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l0
	v13 = v7 + int32(48)
	v17 = F_pg_sprintf(m, v13, int32(_a_F_ReorderBufferCleanupSerializedTXNs_1), v7+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = F___fstatat(m, int32(-100), v13, v7+int32(2112), int32(256))
	mBase = m.M
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	m.G0 = v7 + int32(2208)
	return
L5:
	;
	if v23 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+2116))
	if v26&int32(_a_F_ReorderBufferCleanupSerializedTXNs_2) != int32(_a_F_ReorderBufferCleanupSerializedTXNs_3) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v32 = v7 + int32(48)
	v33 = F_AllocateDir(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v36 = F_ReadDirExtended(m, v33, v32, int32(17))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = v36
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_FreeDir(m, v33)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+19)))
	if v42 != int32(120) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v70 = F_ReadDirExtended(m, v33, v7+int32(48), int32(17))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	if v45 != int32(105) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+21)))
	if v48 != int32(100) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v40 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_ReorderBufferCleanupSerializedTXNs_0)
	v58 = v7 + int32(48)
	v63 = F_pg_snprintf(m, v58, int32(2060), int32(_a_F_ReorderBufferCleanupSerializedTXNs_4), v7+int32(16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v65 = F_unlink(m, v58)
	mBase = m.M
	if v65 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	goto L17
L23:
	;
	if v70 != 0 {
		v40 = v70
		goto L15
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	goto L4
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_ReorderBufferCleanupSerializedTXNs_0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
	F_errmsg(m, int32(_a_F_ReorderBufferCleanupSerializedTXNs_5), v7)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferCleanupSerializedTXNs_6), int32(_a_F_ReorderBufferCleanupSerializedTXNs_7), int32(_a_F_ReorderBufferCleanupSerializedTXNs_8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferCommitChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v14 == v6)|base.B2i32(l2 != v14) == v6 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v21 != 0 {
			v43 = v21
			*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = l4
			*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = l3
			F_ReorderBufferAssignChild(m, l0, l1, l2, int64(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		} else {
			m.G0 = v11 + int32(16)
			return
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = F_hash_search(m, v22, v11+int32(12), int32(0), v11+int32(11))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
			if v30 == int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
				m.G0 = v11 + int32(16)
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v37
				if v38 == int32(0) {
					m.G0 = v11 + int32(16)
					return
				} else {
					v43 = v38
					*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = l4
					*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = l3
					F_ReorderBufferAssignChild(m, l0, l1, l2, int64(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_ReorderBufferRestoreChanges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v125 int32
	_ = v125
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v154 int64
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int64
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(1104)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v24 = l1 + int32(128)
	if base.B2i32(v20 == v5)|base.B2i32(v24 == v20) == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = v20
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	v76 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreChanges[0])))
	v77 = base.I64_div_u_s(v74, v76)
	v91 = v5
	goto L13
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v47
	F_ReorderBufferFreeChange(m, l0, v33-int32(52), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if v45 != v24 {
		v33 = v45
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L6
	} else {
		goto L135
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L6
	} else {
		goto L131
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L127
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L123
	}
L13:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui64(v95) <= base.Ui64(v77) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	m.G0 = v18 + int32(1104)
	return v550
L15:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreChanges[1]))
	if v98 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v550 = v91
	goto L17
L17:
	;
	goto L14
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v101 != int32(-1) {
		v175 = v101
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	if base.Ui32(v541) < base.Ui32(int32(_a_F_ReorderBufferRestoreChanges_0)) {
		v91 = v541
		goto L13
	} else {
		goto L122
	}
L23:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v178 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreChanges[0]))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	if v106 == int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v111 = base.I64_div_u_s(v109, base.I64_extend_i32_s(v105))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v111
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreChanges[0]))
	v115 = v114
	v116 = v111
	goto L27
L26:
	;
	v115 = v105
	v116 = v106
	goto L27
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v119 = v116 * base.I64_extend_i32_s(v115)
	*(*uint32)(unsafe.Add(mBase, uint32(v18-int32(-64)))) = uint32(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(_a_F_ReorderBufferRestoreChanges_1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v117
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreChanges[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v125 + int32(24)
	v130 = int64(base.Ui64(v119) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+60)) = uint32(v130)
	v133 = v18 + int32(80)
	v138 = F_pg_snprintf(m, v133, int32(1024), int32(_a_F_ReorderBufferRestoreChanges_2), v18+int32(48))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v141 = F_PathNameOpenFile(m, v133, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v141
	if int32(0) <= v141 {
		v175 = v141
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreChanges[3]))
	if v149 == int32(44) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v154 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v154 + int64(1)
	v541 = v91
	goto L22
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v18 + int32(80)
	F_errmsg(m, int32(_a_F_ReorderBufferRestoreChanges_3), v18)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferRestoreChanges_4), int32(_a_F_ReorderBufferRestoreChanges_5), int32(_a_F_ReorderBufferRestoreChanges_6))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v198
	v207 = F_FileReadV(m, v197, v18+int32(80), int32(1), v199, int32(167772203))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L46
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v192
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v197 = v196
	v198 = v192
	goto L38
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v183 = F_MemoryContextAlloc(m, v181, int32(72))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(int32(71)) < base.Ui32(v178) {
		v197 = v175
		v198 = v185
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v192 = v183
	goto L39
L44:
	;
	v189 = F_repalloc(m, v185, int32(72))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v192 = v189
	goto L39
L46:
	;
	if v207 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_FileClose(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v207 < int32(0) {
		goto L12
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v216 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v216 + int64(1)
	v541 = v91
	goto L22
L51:
	;
	if v207 != int32(72) {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v226 = v224 + int64(72)
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v231 = v229 + int32(72)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v232 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v250 = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v247 - v250
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v246 + v250
	v260 = F_FileReadV(m, v249, v18+int32(80), int32(1), v248, int32(167772203))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L61
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v241
	v244 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v246 = v241
	v247 = v245
	v248 = v244
	goto L53
L55:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v236 = F_MemoryContextAlloc(m, v235, v231)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v231) <= base.Ui32(v232) {
		v246 = v228
		v247 = v229
		v248 = v226
		goto L53
	} else {
		goto L59
	}
L58:
	;
	v241 = v236
	goto L54
L59:
	;
	v239 = F_repalloc(m, v228, v231)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v241 = v239
	goto L54
L61:
	;
	if v260 < int32(0) {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	if v260 != v264-int32(72) {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v268 + base.I64_extend_i32_u(v260)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v275 = F_MemoryContextAlloc(m, v273, int32(64))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v277 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v275)+56)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275)+48)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275)+40)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275)+32)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275)+24)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275)+16)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275)+8)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v275))) = v277
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v272)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v275))) = v293
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v272)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+8)) = v295
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v272)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+16)) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v272)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+24)) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v272)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+32)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v272)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+40)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v272)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+48)) = v305
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v272)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+56)) = v307
	v310 = v272 + int32(72)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	switch v311 {
	case 0, 1, 2, 8:
		goto L70
	case 3:
		goto L69
	case 4:
		goto L68
	case 5:
		goto L67
	default:
		goto L65
	case 11:
		goto L66
	}
L65:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v440 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L66:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v429 = F_MemoryContextAlloc(m, v425, v426<<(uint(int32(2))%32))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L6
	} else {
		goto L93
	}
L67:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v272)+96))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v272)+88))
	v410 = (v404+v405)<<(uint(int32(2))%32) + int32(72)
	v411 = F_MemoryContextAllocZero(m, v403, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L89
	}
L68:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v396 = v394 << (uint(int32(4)) % 32)
	v397 = F_MemoryContextAlloc(m, v393, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L6
	} else {
		goto L87
	}
L69:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v272)+72))
	v374 = F_MemoryContextAlloc(m, v372, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L6
	} else {
		goto L81
	}
L70:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	if v312 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v272)+72))
	v317 = F_MemoryContextAlloc(m, v313, v314+int32(24))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	v339 = v310
	goto L73
L73:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v275)+40))
	if v342 == int32(0) {
		goto L65
	} else {
		goto L78
	}
L74:
	;
	v319 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = v317 + v319
	*(*int32)(unsafe.Add(mBase, uint32(v275)+36)) = v317
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v310)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v310)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v317)+8)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v310)))
	*(*int64)(unsafe.Add(mBase, uint32(v317))) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v329)+16)) = v329 + v319
	v334 = v272 + int32(92)
	if v314 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+16))
	base.MemoryCopy(m, v336, v334, v314)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v339 = v334 + v314
	goto L73
L78:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v349 = F_MemoryContextAlloc(m, v345, v346+int32(24))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v351 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v349)+16)) = v349 + v351
	*(*int32)(unsafe.Add(mBase, uint32(v275)+40)) = v349
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v339)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+16)) = v355
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v339)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v349)+8)) = v357
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v339)))
	*(*int64)(unsafe.Add(mBase, uint32(v349))) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v275)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+16)) = v361 + v351
	if v346 == int32(0) {
		goto L65
	} else {
		goto L80
	}
L80:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v275)+40))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	base.MemoryCopy(m, v368, v339+int32(20), v346)
	goto L65
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+20)) = v374
	v378 = v272 + int32(76)
	if v373 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	base.MemoryCopy(m, v374, v378, v373)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v380 = v378 + v373
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+24)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v384 = F_MemoryContextAlloc(m, v383, v381)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+28)) = v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	if v387 == int32(0) {
		goto L65
	} else {
		goto L86
	}
L86:
	;
	base.MemoryCopy(m, v384, v380+int32(4), v387)
	goto L65
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+24)) = v397
	if v396 == int32(0) {
		goto L65
	} else {
		goto L88
	}
L88:
	;
	base.MemoryCopy(m, v397, v310, v396)
	goto L65
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+20)) = v411
	if v410 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	base.MemoryCopy(m, v411, v310, v410)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v416 = v411 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v411)+12)) = v416
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v411)+30)) = uint8(v418)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v411)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v411)+20)) = v416 + v420<<(uint(int32(2))%32)
	goto L65
L93:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v433 = v431 << (uint(int32(2)) % 32)
	if v433 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	base.MemoryCopy(m, v429, v310, v433)
	goto L96
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+28)) = v429
	goto L65
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v24
	goto L99
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+56)) = v24
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+52)) = v446
	v449 = v275 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v446)+4)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v449
	v452 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = v452 + int64(1)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	switch v459 {
	case 0, 1, 2, 8:
		goto L107
	case 3:
		goto L106
	case 4:
		goto L105
	case 5:
		goto L104
	default:
		v498 = int32(64)
		goto L102
	case 11:
		goto L103
	}
L100:
	;
	v541 = v91 + int32(1)
	goto L22
L101:
	;
	if v503 == int32(0) {
		goto L100
	} else {
		goto L112
	}
L102:
	;
	v503 = v498
	goto L101
L103:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v498 = v492<<(uint(int32(2))%32) - int32(-64)
	goto L102
L104:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+24))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v484)+16))
	v503 = (v485+v486)<<(uint(int32(2))%32) + int32(136)
	goto L101
L105:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v503 = v479<<(uint(int32(4))%32) - int32(-64)
	goto L101
L106:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v474 = F_strlen(m, v473)
	mBase = m.M
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v503 = v474 + v475 + int32(73)
	goto L101
L107:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v275)+40))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	if v461 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v466 = v462 + int32(84)
	goto L110
L109:
	;
	v466 = int32(64)
	goto L110
L110:
	;
	if v460 == int32(0) {
		v498 = v466
		goto L102
	} else {
		goto L111
	}
L111:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v503 = v466 + v469 + int32(20)
	goto L101
L112:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	if v506 == int32(7) {
		goto L100
	} else {
		goto L113
	}
L113:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v509)+216)) = v510 + v503
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509)+40))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v514 + v503
	if v513 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v517 = v513
	goto L116
L115:
	;
	v517 = v509
	goto L116
L116:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v517)+220)) = v518 + v503
	if v510 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_remove(m, v521, v509+int32(204))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L6
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v526, v509+int32(204))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L6
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	goto L100
L122:
	;
	v550 = v541
	goto L17
L123:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_ReorderBufferRestoreChanges_7), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferRestoreChanges_4), int32(_a_F_ReorderBufferRestoreChanges_8), int32(_a_F_ReorderBufferRestoreChanges_6))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v207
	F_errmsg(m, int32(_a_F_ReorderBufferRestoreChanges_9), v18+int32(32))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferRestoreChanges_4), int32(_a_F_ReorderBufferRestoreChanges_10), int32(_a_F_ReorderBufferRestoreChanges_6))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_ReorderBufferRestoreChanges_7), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferRestoreChanges_4), int32(_a_F_ReorderBufferRestoreChanges_11), int32(_a_F_ReorderBufferRestoreChanges_6))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v611 - int32(72)
	F_errmsg(m, int32(_a_F_ReorderBufferRestoreChanges_9), v18+int32(16))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferRestoreChanges_4), int32(_a_F_ReorderBufferRestoreChanges_12), int32(_a_F_ReorderBufferRestoreChanges_6))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
