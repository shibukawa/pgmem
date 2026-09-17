package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LaunchParallelWorkers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(1472)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v13 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(1472)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+616))
	if v21 != v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[1]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[2]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v30 = base.I32_div_s(v20-v27, int32(640))
	v32 = base.I32_rem_s(v30, int32(16))
	v37 = v24 + v32<<(uint(int32(7))%32) + int32(_a_F_LaunchParallelWorkers_0)
	v39 = F_LWLockAcquire(m, v37, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v63 = int32(_a_F_LaunchParallelWorkers_1)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[3]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[3])) = v67
	v70 = v11 + int32(12)
	base.MemoryFill(m, v70, int32(0), int32(1460))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
	v79 = F_pg_snprintf(m, v70, int32(96), int32(_a_F_LaunchParallelWorkers_2), v11)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L13
	}
L7:
	;
	return
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+616)) = v42
	v45 = v42 + int32(620)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+624))
	if v46 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v45
	v50 = v45
	goto L11
L10:
	;
	v50 = v46
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+628)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v42)+632)) = v50
	v54 = v42 + int32(628)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v42)+624)) = v54
	F_LWLockRelease(m, v37)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v86 = F_pg_snprintf(m, v11+int32(108), int32(96), int32(_a_F_LaunchParallelWorkers_3), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+212)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+204)) = int64(4294967315)
	v96 = F_pg_sprintf(m, v11+int32(216), int32(_a_F_LaunchParallelWorkers_4), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v102 = F_pg_sprintf(m, v11+int32(1240), int32(_a_F_LaunchParallelWorkers_5), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+1336)) = v105
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+1468)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v110 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v118 = v2
	v120 = v2
	goto L20
L18:
	;
	goto L19
L19:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v175 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+1340)) = v118
	if v120&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L19
L22:
	;
	v164 = v118 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v164 < v165 {
		v118 = v164
		v120 = v162
		goto L20
	} else {
		goto L28
	}
L23:
	;
	v146 = v118 << (uint(int32(3)) % 32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v146+v147))) = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151+v146)+4))
	F_shm_mq_detach(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L27
	}
L24:
	;
	v127 = v118 << (uint(int32(3)) % 32)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v130 = F_RegisterDynamicBackgroundWorker(m, v11+int32(12), v127+v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	if v130 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v135 = v134 + v127
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v139 + int32(1)
	v162 = int32(0)
	goto L22
L27:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v146)+4)) = int32(0)
	v162 = int32(1)
	goto L22
L28:
	;
	goto L21
L29:
	;
	v178 = F_palloc0(m, v175)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LaunchParallelWorkers[3])) = v64
	goto L1
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v178
	goto L31
}
func F_ListenServerPort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int64
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
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
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v608 int32
	_ = v608
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(1744)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1372)) = v6
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+1352)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v15)+1360)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1368)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1344)) = l0
	v26 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1348)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1340)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v26
	if l0 == v26 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(1744)
	return v608
L2:
	;
	v116 = F_pg_getaddrinfo_all(m, l1, v110, v15+int32(1340), v15+int32(1372))
	mBase = m.M
	if v116 != 0 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+276)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = l3
	v37 = v15 + int32(304)
	v42 = F_pg_snprintf(m, v37, int32(1024), int32(_a_F_ListenServerPort_0), v15+int32(272))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+288)) = l2
	v101 = v15 + int32(1696)
	v106 = F_pg_snprintf(m, v101, int32(32), int32(_a_F_ListenServerPort_1), v15+int32(288))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L20
	}
L6:
	;
	return int32(0)
L7:
	;
	v46 = F_strlen(m, v37)
	mBase = m.M
	if base.Ui32(int32(108)) <= base.Ui32(v46) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = int32(-1)
	v52 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v68 = v15 + int32(304)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+304)))
	if v69 == int32(64) {
		v110 = v68
		goto L2
	} else {
		goto L15
	}
L11:
	;
	if v52 == int32(0) {
		v608 = v49
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(107)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
	F_errmsg(m, int32(_a_F_ListenServerPort_2), v15)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), int32(459), int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v608 = v49
	goto L1
L15:
	;
	v72 = m.G0
	v74 = v72 - int32(1040)
	m.G0 = v74
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v68
	v78 = v74 + int32(16)
	v81 = F_pg_snprintf(m, v78, int32(1024), int32(_a_F_ListenServerPort_5), v74)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_CreateLockFile(m, v78, int32(1), l3, int32(0), v68)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v74 + int32(1040)
	v90 = F_unlink(m, v68)
	mBase = m.M
	v91 = int32(_a_F_ListenServerPort_6)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[0]))
	v94 = F_pstrdup(m, v68)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v96 = F_lappend(m, v93, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[0])) = v96
	v110 = v68
	goto L2
L20:
	;
	v110 = v101
	goto L2
L21:
	;
	v497 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L6
	} else {
		goto L143
	}
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v117 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v125 = v117
	v132 = v6
	goto L24
L24:
	;
	if base.B2i32(l0 == int32(1)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1344))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v474 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L26:
	;
	goto L25
L27:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v125)+28))
	if v466 != 0 {
		v125 = v466
		v132 = v464
		goto L24
	} else {
		goto L129
	}
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v136 == int32(1) {
		v464 = v132
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[1]))
	if v140 == int32(64) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v145 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v162 = v15 + int32(304)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	switch v165 - int32(1) {
	case 0:
		v196 = v165
		v197 = int32(_a_F_ListenServerPort_7)
		v199 = v162
		goto L39
	case 1:
		v184 = int32(_a_F_ListenServerPort_8)
		goto L40
	default:
		goto L41
	case 9:
		goto L42
	}
L35:
	;
	if v145 == int32(0) {
		v472 = v132
		goto L26
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(64)
	F_errmsg(m, int32(_a_F_ListenServerPort_9), v15+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), int32(504), int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v472 = v132
	goto L26
L39:
	;
	v202 = F_socket(m, v196, int32(1), int32(0))
	mBase = m.M
	if v202 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v188 = v15 + int32(1376)
	v190 = int32(0)
	v193 = F_pg_getnameinfo_all(m, v185, v186, v188, int32(255), v190, v190, int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v165
	v171 = v15 + int32(1632)
	v176 = F_pg_snprintf(m, v171, int32(64), int32(_a_F_ListenServerPort_10), v15+int32(224))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	v184 = int32(_a_F_ListenServerPort_11)
	goto L40
L43:
	;
	v178 = int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v179 == v178 {
		v196 = v178
		v197 = v171
		v199 = v162
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v184 = v171
	goto L40
L45:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v196 = v195
	v197 = v184
	v199 = v188
	goto L39
L46:
	;
	v207 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v225 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v227 == v225 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	if v207 == int32(0) {
		v464 = v132
		goto L27
	} else {
		goto L50
	}
L50:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v197
	F_errmsg(m, int32(_a_F_ListenServerPort_12), v15+int32(32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), int32(547), int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v464 = v132
	goto L27
L54:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v235 = F_bind(m, v202, v233, v234)
	mBase = m.M
	if v235 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L57:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[2]))
	v242 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v274 != int32(1) {
		goto L74
	} else {
		goto L75
	}
L60:
	;
	if v242 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v273 = F_close(m, v202)
	mBase = m.M
	v464 = v132
	goto L27
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v197
	F_errmsg(m, int32(_a_F_ListenServerPort_13), v15-int32(-64))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	if v239 == int32(3) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	if v255 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), int32(624), int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L73
	}
L69:
	;
	v261 = int32(_a_F_ListenServerPort_14)
	goto L71
L70:
	;
	v261 = int32(_a_F_ListenServerPort_15)
	goto L71
L71:
	;
	F_errhint(m, v261, v15+int32(48))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	goto L63
L74:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[3]))
	v380 = int32(0)
	v384 = m.Env.X__syscall_listen(m, v202, v377<<(uint(int32(1))%32), v380, v380, v380, v380)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v384) {
		goto L105
	} else {
		goto L106
	}
L75:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v277 == int32(64) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[4]))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v282 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v374 = F_close(m, v202)
	mBase = m.M
	v472 = v132
	goto L26
L78:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[5]))
	v351 = F_chmod(m, v110, v350)
	mBase = m.M
	if v351 != int32(-1) {
		goto L74
	} else {
		goto L98
	}
L79:
	;
	v289 = F_strtox_2(m, v281, v15+int32(1740), int32(10), int64(4294967295))
	mBase = m.M
	goto L80
L80:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1740))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v295 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), v343, int32(_a_F_ListenServerPort_16))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L97
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[2])) = int32(44)
	v303 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v317 = m.Env.X__syscall_fchownat(m, int32(-100), v110, int32(-1), base.I32_wrap_i64(v289), int32(0))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v317) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	if v303 == int32(0) {
		goto L77
	} else {
		goto L86
	}
L86:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v308
	F_errmsg(m, int32(_a_F_ListenServerPort_17), v15+int32(160))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v343 = int32(755)
	goto L81
L88:
	;
	if v325 != int32(-1) {
		goto L78
	} else {
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[2])) = int32(0) - v317
	v325 = int32(-1)
	goto L91
L90:
	;
	v325 = v317
	goto L91
L91:
	;
	goto L88
L92:
	;
	v330 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	if v330 == int32(0) {
		goto L77
	} else {
		goto L94
	}
L94:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v110
	F_errmsg(m, int32(_a_F_ListenServerPort_18), v15+int32(144))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v343 = int32(765)
	goto L81
L97:
	;
	v347 = F_close(m, v202)
	mBase = m.M
	v472 = v132
	goto L26
L98:
	;
	v356 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	if v356 == int32(0) {
		goto L77
	} else {
		goto L100
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v110
	F_errmsg(m, int32(_a_F_ListenServerPort_19), v15+int32(128))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), int32(776), int32(_a_F_ListenServerPort_16))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	goto L77
L104:
	;
	if v392 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[2])) = int32(0) - v384
	v392 = int32(-1)
	goto L107
L106:
	;
	v392 = v384
	goto L107
L107:
	;
	goto L104
L108:
	;
	v397 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L6
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v417 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L118
	}
L111:
	;
	if v397 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v413 = F_close(m, v202)
	mBase = m.M
	v464 = v132
	goto L27
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v197
	F_errmsg(m, int32(_a_F_ListenServerPort_20), v15+int32(80))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), int32(652), int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	if v414 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v446 = int32(_a_F_ListenServerPort_21)
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v447<<(uint(int32(2))%32)))) = v202
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[1]))
	v455 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ListenServerPort[1])) = v454 + v455
	v464 = v132 + v455
	goto L27
L120:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), v442, int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L6
	} else {
		goto L128
	}
L121:
	;
	if v417 == int32(0) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if v417 == int32(0) {
		goto L119
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v199
	F_errmsg(m, int32(_a_F_ListenServerPort_22), v15+int32(96))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	v442 = int32(660)
	goto L120
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v197
	F_errmsg(m, int32(_a_F_ListenServerPort_23), v15+int32(112))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	v442 = int32(665)
	goto L120
L128:
	;
	goto L119
L129:
	;
	v472 = v464
	goto L26
L130:
	;
	if v472 != 0 {
		goto L140
	} else {
		goto L141
	}
L131:
	;
	goto L130
L132:
	;
	if v475 == int32(0) {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v475 == int32(0) {
		goto L131
	} else {
		goto L139
	}
L135:
	;
	v481 = v475
	goto L136
L136:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+28))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v481)+20))
	F_emscripten_builtin_free(m, v483)
	mBase = m.M
	F_emscripten_builtin_free(m, v481)
	mBase = m.M
	if v482 != 0 {
		v481 = v482
		goto L136
	} else {
		goto L138
	}
L137:
	;
	goto L131
L138:
	;
	goto L137
L139:
	;
	F_freeaddrinfo(m, v475)
	mBase = m.M
	goto L131
L140:
	;
	v493 = int32(0)
	goto L142
L141:
	;
	v493 = int32(-1)
	goto L142
L142:
	;
	v608 = v493
	goto L1
L143:
	;
	if l1 != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v585 = int32(-1)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v586 == int32(0) {
		v608 = v585
		goto L1
	} else {
		goto L174
	}
L145:
	;
	F_errfinish(m, int32(_a_F_ListenServerPort_3), v581, int32(_a_F_ListenServerPort_4))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L173
	}
L146:
	;
	if v497 == int32(0) {
		goto L144
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if v497 == int32(0) {
		goto L144
	} else {
		goto L161
	}
L149:
	;
	v504 = int32(_a_F_ListenServerPort_24)
	v506 = v116 + int32(1)
	if v506 == int32(0) {
		v526 = v504
		goto L151
	} else {
		goto L152
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v526 + base.B2i32(v528 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = l1
	F_errmsg(m, int32(_a_F_ListenServerPort_25), v15+int32(256))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L6
	} else {
		goto L160
	}
L151:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	goto L150
L152:
	;
	v510 = v504
	v511 = v506
	goto L153
L153:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	if v512 == int32(0) {
		v526 = v510
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v526 = v522
	goto L151
L155:
	;
	v516 = v510
	goto L156
L156:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
	if v520 != 0 {
		v516 = v516 + int32(1)
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v522 = v516 + int32(2)
	v524 = v511 + int32(1)
	if v524 != 0 {
		v510 = v522
		v511 = v524
		goto L153
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	goto L154
L160:
	;
	v581 = int32(478)
	goto L145
L161:
	;
	v545 = int32(_a_F_ListenServerPort_24)
	v547 = v116 + int32(1)
	if v547 == int32(0) {
		v567 = v545
		goto L163
	} else {
		goto L164
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v567 + base.B2i32(v569 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v110
	F_errmsg(m, int32(_a_F_ListenServerPort_26), v15+int32(240))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L6
	} else {
		goto L172
	}
L163:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	goto L162
L164:
	;
	v551 = v545
	v552 = v547
	goto L165
L165:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v553 == int32(0) {
		v567 = v551
		goto L163
	} else {
		goto L167
	}
L166:
	;
	v567 = v563
	goto L163
L167:
	;
	v557 = v551
	goto L168
L168:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	if v561 != 0 {
		v557 = v557 + int32(1)
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v563 = v557 + int32(2)
	v565 = v552 + int32(1)
	if v565 != 0 {
		v551 = v563
		v552 = v565
		goto L165
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	goto L166
L172:
	;
	v581 = int32(482)
	goto L145
L173:
	;
	goto L144
L174:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1344))
	if v589 == int32(1) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v608 = v585
	goto L1
L176:
	;
	goto L175
L177:
	;
	if v586 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	if v586 == int32(0) {
		goto L176
	} else {
		goto L184
	}
L180:
	;
	v595 = v586
	goto L181
L181:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+28))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v595)+20))
	F_emscripten_builtin_free(m, v597)
	mBase = m.M
	F_emscripten_builtin_free(m, v595)
	mBase = m.M
	if v596 != 0 {
		v595 = v596
		goto L181
	} else {
		goto L183
	}
L182:
	;
	goto L176
L183:
	;
	goto L182
L184:
	;
	F_freeaddrinfo(m, v586)
	mBase = m.M
	goto L176
}
func F_LockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = int32(0)
	v7 = F_LockAcquireExtended(m, l0, l1, l2, l3, v5, v5)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_LookupOperWithArgs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v10 != 0 {
		v11 = F_LookupTypeNameOid(m, v10, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			if v9 != 0 {
				v16 = F_LookupTypeNameOid(m, v9, l1)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = v16
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v20 = F_LookupOperName(m, v19, v15, v18, l1)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v20
					}
				}
			} else {
				v18 = v3
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v20 = F_LookupOperName(m, v19, v15, v18, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			}
		}
	} else {
		v15 = v3
		if v9 != 0 {
			v16 = F_LookupTypeNameOid(m, v9, l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = v16
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v20 = F_LookupOperName(m, v19, v15, v18, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			}
		} else {
			v18 = v3
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = F_LookupOperName(m, v19, v15, v18, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v20
			}
		}
	}
}
func F___loc_is_allocated(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 != int32(0)) & base.B2i32(l0 != int32(_a_F___loc_is_allocated_0)) & base.B2i32(l0 != int32(_a_F___loc_is_allocated_1)) & base.B2i32(l0 != int32(_a_F___loc_is_allocated_2)) & base.B2i32(l0 != int32(_a_F___loc_is_allocated_3))
}
func F__lca(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v16 = v10 + int32(16)
		v17 = F_ArrayGetNItemsSafe(m, v14, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if v19 < int32(2) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				v23 = F_array_contains_nulls(m, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F__lca_0), int32(0))
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F__lca_1), int32(308), int32(_a_F__lca_2))
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v27 = F_palloc(m, v17<<(uint(int32(2))%32))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							if v17 <= int32(0) {
							} else {
								if v22 != 0 {
									v37 = v22
								} else {
									v37 = (v19<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v38 = v10 + v37
								if v17&int32(1) == int32(0) {
									v57 = v38
									v58 = v17
								} else {
									v44 = v17 - int32(1)
									v45 = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v27+v44<<(uint(v45)%32)))) = v38
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
									v57 = v38 + (int32(base.Ui32(v49)>>(uint(v45)%32))+int32(3))&int32(2147483644)
									v58 = v44
								}
								if v17 == int32(1) {
								} else {
									v62 = v57
									v63 = v58
									for {
										v69 = int32(2)
										*(*int32)(unsafe.Add(mBase, uint32(v27+v63<<(uint(v69)%32)-int32(4)))) = v62
										v76 = v63 - v69
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
										v83 = int32(3)
										v85 = int32(2147483644)
										v87 = v62 + (int32(base.Ui32(v80)>>(uint(v69)%32))+v83)&v85
										*(*int32)(unsafe.Add(mBase, uint32(v27+v76<<(uint(v69)%32)))) = v87
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
										if v69 < v63 {
											v62 = v87 + (int32(base.Ui32(v89)>>(uint(v69)%32))+v83)&v85
											v63 = v76
											continue
										} else {
											break
										}
										break
									}
								}
							}
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v108 = F_ArrayGetNItemsSafe(m, v107, v16)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								v110 = F_lca_inner(m, v27, v108)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v27)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v114 != v10 {
											F_pfree(m, v10)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												if v110 != 0 {
													v121 = v110
												} else {
													v118 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
													v121 = int32(0)
												}
												return v121
											}
										} else {
											if v110 != 0 {
												v121 = v110
											} else {
												v118 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
												v121 = int32(0)
											}
											return v121
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F__lca_3), int32(0))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F__lca_1), int32(304), int32(_a_F__lca_2))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
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
	}
}
func F__lt_q_rregex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__lt_q_rregex_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13937(m, l0, l1, int64(4294967768))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_latin2_to_win1250(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13936(m, l0, int32(_a_F_latin2_to_win1250_0), int32(29), int32(9))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_latin2mic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v4 = l3
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v49)
	return v47 - l0
L2:
	;
	v42 = l1
	v47 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v12 = l1
	v13 = l2
	v17 = l0
	goto L5
L5:
	;
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17))))
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v42 = v36
	v47 = v34
	goto L1
L7:
	;
	if l5 != 0 {
		v42 = v12
		v47 = v17
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v19 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_report_invalid_encoding(m, l4, v17, v13)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v4)
	v31 = v12 + int32(1)
	goto L15
L14:
	;
	v31 = v12
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v19)
	v33 = int32(1)
	v34 = v17 + v33
	v36 = v31 + v33
	if v33 < v13 {
		v12 = v36
		v13 = v13 - v33
		v17 = v34
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L6
}
func F_len_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(4))))
	if v9 == v2 {
		return int32(0)
	} else {
		v15 = v9 & int32(3)
		if base.Ui32(v9) < base.Ui32(int32(4)) {
			v51 = l0
			v52 = int32(0)
			v57 = v51
			v58 = v52
			v62 = v2
			for {
				v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
				v66 = v58 + base.B2i32(int32(-65) < v63)
				v67 = int32(1)
				v70 = v62 + v67
				if v70 != v15 {
					v57 = v57 + v67
					v58 = v66
					v62 = v70
					continue
				} else {
					break
				}
				break
			}
			v73 = v66
		} else {
			v22 = l0
			v23 = int32(0)
			v26 = v2
			for {
				v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
				v29 = int32(-65)
				v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+1)))
				v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+2)))
				v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+3)))
				v43 = v23 + base.B2i32(v29 < v28) + base.B2i32(v29 < v32) + base.B2i32(v29 < v36) + base.B2i32(v29 < v40)
				v44 = int32(4)
				v45 = v22 + v44
				v47 = v26 + v44
				if v47 != v9&int32(-4) {
					v22 = v45
					v23 = v43
					v26 = v47
					continue
				} else {
					break
				}
				break
			}
			if v15 == int32(0) {
				v73 = v43
			} else {
				v51 = v45
				v52 = v43
				v57 = v51
				v58 = v52
				v62 = v2
				for {
					v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
					v66 = v58 + base.B2i32(int32(-65) < v63)
					v67 = int32(1)
					v70 = v62 + v67
					if v70 != v15 {
						v57 = v57 + v67
						v58 = v66
						v62 = v70
						continue
					} else {
						break
					}
					break
				}
				v73 = v66
			}
		}
		return v73
	}
}
func F_lengthCompareJsonbPair(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v7 < v6 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v6) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v12 = int32(1)
	goto L6
L5:
	;
	v12 = int32(-1)
	goto L6
L6:
	;
	return v12
L7:
	;
	return v91
L8:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if base.Ui32(v88) < base.Ui32(v87) {
		goto L31
	} else {
		goto L32
	}
L9:
	;
	v78 = int32(0)
	if v77|base.B2i32(l2 == v78) == v78 {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v77 = int32(0)
	goto L9
L11:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L21
L12:
	;
	if (v14|v15)&int32(3) != 0 {
		v46 = v14
		v47 = v15
		v48 = v6
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v39 = v14
	v40 = v15
	v41 = v6
	goto L14
L14:
	;
	if v41 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v23 = v14
	v24 = v15
	v25 = v6
	goto L16
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L14
L18:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L11
L21:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v77 = v56 - v57
	goto L9
L23:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L10
L27:
	;
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v83)
	goto L8
L28:
	;
	goto L29
L29:
	;
	if v77 != 0 {
		v91 = v77
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L8
L31:
	;
	v90 = int32(-1)
	goto L33
L32:
	;
	v90 = int32(1)
	goto L33
L33:
	;
	v91 = v90
	goto L7
}
func F_levenshtein_with_costs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = v11 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v28 = v26 & int32(1)
			if v28 != 0 {
				v29 = v16
			} else {
				v29 = v11 + int32(4)
			}
			if v26 == int32(1) {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v35 == int32(18) {
					v38 = int32(16)
				} else {
					v38 = int32(0)
				}
				if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v45 = int32(4)
				} else {
					v45 = v38
				}
				v56 = v45
			} else {
				v46 = int32(1)
				if v28 != 0 {
					v56 = int32(base.Ui32(v26)>>(uint(v46)%32)) - v46
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v57 = int32(1)
			v58 = v18 + v57
			if v20&v57 != 0 {
				v63 = v58
			} else {
				v63 = v18 + int32(4)
			}
			if v20 == int32(1) {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
				if v69 == int32(18) {
					v72 = int32(16)
				} else {
					v72 = int32(0)
				}
				if base.Ui32((v69-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v79 = int32(4)
				} else {
					v79 = v72
				}
				v92 = v79
			} else {
				v80 = int32(1)
				if v20&v80 != 0 {
					v92 = int32(base.Ui32(v20)>>(uint(v80)%32)) - v80
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v92 = int32(base.Ui32(v86)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v93 = F_varstr_levenshtein(m, v29, v56, v63, v92, v23, v22, v21)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				return v93
			}
		}
	}
}
func F_like_regex_support(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v5 - int32(458) {
	case 0:
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v8 != 0 {
			v21 = float64(0.005)
			*(*float64)(unsafe.Add(mBase, uint32(l0)+40)) = v21
			return l0
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v11 = int32(0)
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v17 = F_patternsel_common(m, v10, v11, v12, v13, v14, v15, l1, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = v17
				*(*float64)(unsafe.Add(mBase, uint32(l0)+40)) = v21
				return l0
			}
		}
	default:
		v51 = v3
		return v51
	case 3:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v24 != 0 {
			v51 = v3
			return v51
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v25 == int32(0) {
				v51 = v3
				return v51
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				switch v28 - int32(15) {
				case 0:
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v48 = F_match_pattern_prefix(m, v43, v44, l1, v45, v46, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v51 = v48
						return v51
					}
				default:
					v51 = v3
					return v51
				case 2:
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v38 = F_match_pattern_prefix(m, v33, v34, l1, v35, v36, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						return v38
					}
				}
			}
		}
	}
}
func F_likesel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13988(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_lo_manage(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L27
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L27
	} else {
		goto L70
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L27
	} else {
		goto L67
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L27
	} else {
		goto L64
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v16 != int32(442) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19&int32(4) == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v25 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v33 < v35 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v75 <= int32(0) {
		goto L1
	} else {
		goto L24
	}
L10:
	;
	v75 = v41 + int32(1)
	goto L9
L11:
	;
	v40 = v35
	v41 = v33
	goto L14
L12:
	;
	goto L13
L13:
	;
	v64 = F_SystemAttributeByName(m, v32)
	mBase = m.M
	if v64 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v47 = v31 + v40<<(uint(int32(4))%32) + v41*int32(100)
	v50 = F_namestrcmp(m, v47+int32(24), v32)
	mBase = m.M
	if v50 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+111)))
	if v53 != int32(1) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v57 = v41 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v57 < v58 {
		v40 = v58
		v41 = v57
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	goto L15
L21:
	;
	v75 = int32(-9)
	goto L9
L22:
	;
	goto L23
L23:
	;
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+74)))
	v75 = v68
	goto L9
L24:
	;
	v79 = v19 & int32(3)
	if v29 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v79 != int32(1) {
		goto L54
	} else {
		goto L55
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v85 = F_bms_is_member(m, v75+int32(7), v84)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	if v85 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v91 = F_SPI_getvalue(m, v28, v31, v75)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v93 = F_SPI_getvalue(m, v29, v31, v75)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v91 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v91 == int32(0) {
		goto L25
	} else {
		goto L52
	}
L33:
	;
	F_pfree(m, v93)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L27
	} else {
		goto L51
	}
L34:
	;
	if v93 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if v93 == int32(0) {
		goto L32
	} else {
		goto L50
	}
L37:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if base.B2i32(v97 == int32(0))|base.B2i32(v97 != v100) != 0 {
		v118 = v97
		v119 = v100
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v124 = int32(0)
	v128 = F_strtox_2(m, v91, v124, int32(10), int64(4294967295))
	mBase = m.M
	goto L48
L40:
	;
	if v118-v119 == int32(0) {
		goto L33
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v103 = v91
	v104 = v93
	goto L43
L43:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v108 == int32(0) {
		v118 = v108
		v119 = v107
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v118 = v108
	v119 = v107
	goto L41
L45:
	;
	v111 = int32(1)
	if v108 == v107 {
		v103 = v103 + v111
		v104 = v104 + v111
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L39
L48:
	;
	v130 = F_DirectFunctionCall1Coll(m, int32(2352), v124, base.I32_wrap_i64(v128))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L27
	} else {
		goto L49
	}
L49:
	;
	goto L36
L50:
	;
	goto L33
L51:
	;
	goto L32
L52:
	;
	F_pfree(m, v91)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L27
	} else {
		goto L53
	}
L53:
	;
	goto L25
L54:
	;
	m.G0 = v11 + int32(48)
	if v79 == int32(2) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v144 = F_SPI_getvalue(m, v28, v31, v75)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L27
	} else {
		goto L56
	}
L56:
	;
	if v144 == int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v149 = int32(0)
	v153 = F_strtox_2(m, v144, v149, int32(10), int64(4294967295))
	mBase = m.M
	goto L58
L58:
	;
	v155 = F_DirectFunctionCall1Coll(m, int32(2352), v149, base.I32_wrap_i64(v153))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L27
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v144)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L27
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	v165 = v29
	goto L63
L62:
	;
	v165 = v28
	goto L63
L63:
	;
	return v165
L64:
	;
	F_errmsg_internal(m, int32(_a_F_lo_manage_0), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L27
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_lo_manage_1), int32(39), int32(_a_F_lo_manage_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L27
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
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v185
	F_errmsg_internal(m, int32(_a_F_lo_manage_3), v11)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L27
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_lo_manage_1), int32(43), int32(_a_F_lo_manage_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v200
	F_errmsg_internal(m, int32(_a_F_lo_manage_4), v11+int32(16))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_lo_manage_1), int32(55), int32(_a_F_lo_manage_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v217
	F_errmsg_internal(m, int32(_a_F_lo_manage_5), v11+int32(32))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L27
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_lo_manage_1), int32(71), int32(_a_F_lo_manage_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L27
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_load_external_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(_a_F_load_external_function_0)
	goto L3
L1:
	;
	if v49-v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	goto L4
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = l0
	v20 = v11
	v21 = int32(8)
	v22 = v18
	goto L9
L6:
	;
	v45 = v11
	v49 = int32(0)
	goto L7
L7:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L1
L8:
	;
	v45 = v40
	v49 = v42
	goto L7
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v22 != v24)|base.B2i32(v24 == int32(0)) != 0 {
		v40 = v20
		v42 = v22
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v40 = v34
	v42 = int32(0)
	goto L8
L11:
	;
	v30 = v21 - int32(1)
	if v30 == int32(0) {
		v40 = v20
		v42 = v22
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v33 = int32(1)
	v34 = v20 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v35 != 0 {
		v19 = v19 + v33
		v20 = v34
		v21 = v30
		v22 = v35
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v64 = l0
	goto L16
L15:
	;
	v59 = l0 + int32(8)
	v61 = Fn13878(m, v59, int32(47))
	mBase = m.M
	goto L17
L16:
	;
	v65 = F_expand_dynamic_library_name(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v62 = l0
	goto L20
L19:
	;
	v62 = v59
	goto L20
L20:
	;
	v64 = v62
	goto L16
L21:
	;
	return int32(0)
L22:
	;
	v69 = F_internal_load_library(m, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if l3 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v69
	goto L26
L25:
	;
	goto L26
L26:
	;
	v72 = F_pgmem_dlsym(m, v69, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v74 = int32(0)
	if v72|base.B2i32(l2 == v74) == v74 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L21
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_pfree(m, v65)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L21
	} else {
		goto L35
	}
L31:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg(m, int32(_a_F_load_external_function_1), v9)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_load_external_function_2), int32(134), int32(_a_F_load_external_function_3))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	m.G0 = v9 + int32(16)
	return v72
}
func F_load_rangetype_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_SearchSysCache1(m, int32(55), v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v19 = v17 + v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			F_ReleaseCatCache(m, v15)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = F_get_opclass_family(m, v25)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = F_get_opclass_input_type(m, v25)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v28
						v34 = F_get_opfamily_proc(m, v28, v30, v30, int32(1))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							if v34 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v28
									*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v30
									*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v30
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(1)
									F_errmsg_internal(m, int32(_a_F_load_rangetype_info_0), v11+int32(16))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_load_rangetype_info_1), int32(1040), int32(_a_F_load_rangetype_info_2))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_load_rangetype_info[0]))
								F_fmgr_info_cxt(m, v34, l0+int32(212), v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									if v24 != 0 {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_load_rangetype_info[0]))
										F_fmgr_info_cxt(m, v24, l0+int32(240), v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											if v23 != 0 {
												v53 = *(*int32)(unsafe.Add(mBase, _c_F_load_rangetype_info[0]))
												F_fmgr_info_cxt(m, v23, l0+int32(268), v53)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return
												} else {
													v57 = F_lookup_type_cache(m, v20, int32(0))
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v57
														m.G0 = v11 + int32(32)
														return
													}
												}
											} else {
												v57 = F_lookup_type_cache(m, v20, int32(0))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v57
													m.G0 = v11 + int32(32)
													return
												}
											}
										}
									} else {
										if v23 != 0 {
											v53 = *(*int32)(unsafe.Add(mBase, _c_F_load_rangetype_info[0]))
											F_fmgr_info_cxt(m, v23, l0+int32(268), v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												v57 = F_lookup_type_cache(m, v20, int32(0))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v57
													m.G0 = v11 + int32(32)
													return
												}
											}
										} else {
											v57 = F_lookup_type_cache(m, v20, int32(0))
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v57
												m.G0 = v11 + int32(32)
												return
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
				F_errmsg_internal(m, int32(_a_F_load_rangetype_info_3), v11)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_load_rangetype_info_1), int32(1020), int32(_a_F_load_rangetype_info_2))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
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
func F_locate_agg_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13978(m, l0, l1, int32(1044), int32(-1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_locate_var_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13978(m, l0, l1, int32(902), int32(-1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_lock_twophase_postcommit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_TwoPhaseGetDummyProc(m, l0, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
		if base.Ui32((v12-int32(3))&int32(255)) <= base.Ui32(int32(253)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
				F_errmsg_internal(m, int32(_a_F_lock_twophase_postcommit_0), v7)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_lock_twophase_postcommit_1), int32(_a_F_lock_twophase_postcommit_2), int32(_a_F_lock_twophase_postcommit_3))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_lock_twophase_postcommit[0])))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			F_LockRefindAndRelease(m, v34, v10, l2, v35, int32(1))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_log10(m *base.Module, l0 float64) float64 {
	var v12 int64
	_ = v12
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v113 float64
	_ = v113
	var v125 float64
	_ = v125
	v12 = base.I64_reinterpret_f64(l0)
	if v12 <= int64(4503599627370495) {
		if base.F64_eq(l0, float64(0)) != 0 {
			return base.F64_div(float64(-1), base.F64_mul(l0, l0))
		} else {
			if int64(0) <= v12 {
				v42 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(1.8014398509481984e+16)))
				v46 = v42
				v48 = int32(-1077)
				v49 = base.I32_wrap_i64(int64(base.Ui64(v42) >> (uint(int64(32)) % 64)))
				v51 = v49 + int32(_a_F_log10_0)
				v55 = base.F64_convert_i32_s(int32(base.Ui32(v51)>>(uint(int32(20))%32)) + v48)
				v57 = base.F64_mul(v55, float64(0.30102999566361177))
				v70 = base.F64_add(base.F64_reinterpret_i64(v46&int64(4294967295)|base.I64_extend_i32_u(v51&int32(_a_F_log10_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
				v73 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
				v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v70, v73)) & int64(-4294967296))
				v79 = float64(0.4342944818781689)
				v80 = base.F64_mul(v78, v79)
				v81 = base.F64_add(v57, v80)
				v86 = base.F64_div(v70, base.F64_add(v70, float64(2)))
				v87 = base.F64_mul(v86, v86)
				v88 = base.F64_mul(v87, v87)
				v113 = base.F64_add(base.F64_mul(v86, base.F64_add(v73, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v87, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v70, v78), v73))
				v125 = base.F64_add(v81, base.F64_add(base.F64_add(v80, base.F64_sub(v57, v81)), base.F64_add(base.F64_mul(v113, v79), base.F64_add(base.F64_mul(v55, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v113, v78), float64(2.5082946711645275e-11))))))
				return v125
			} else {
				return base.F64_div(base.F64_sub(l0, l0), float64(0))
			}
		}
	} else {
		if base.Ui64(int64(9218868437227405311)) < base.Ui64(v12) {
			v125 = l0
			return v125
		} else {
			v29 = int32(-1023)
			v31 = int64(base.Ui64(v12) >> (uint(int64(32)) % 64))
			if v31 != int64(1072693248) {
				v46 = v12
				v48 = v29
				v49 = base.I32_wrap_i64(v31)
				v51 = v49 + int32(_a_F_log10_0)
				v55 = base.F64_convert_i32_s(int32(base.Ui32(v51)>>(uint(int32(20))%32)) + v48)
				v57 = base.F64_mul(v55, float64(0.30102999566361177))
				v70 = base.F64_add(base.F64_reinterpret_i64(v46&int64(4294967295)|base.I64_extend_i32_u(v51&int32(_a_F_log10_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
				v73 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
				v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v70, v73)) & int64(-4294967296))
				v79 = float64(0.4342944818781689)
				v80 = base.F64_mul(v78, v79)
				v81 = base.F64_add(v57, v80)
				v86 = base.F64_div(v70, base.F64_add(v70, float64(2)))
				v87 = base.F64_mul(v86, v86)
				v88 = base.F64_mul(v87, v87)
				v113 = base.F64_add(base.F64_mul(v86, base.F64_add(v73, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v87, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v70, v78), v73))
				v125 = base.F64_add(v81, base.F64_add(base.F64_add(v80, base.F64_sub(v57, v81)), base.F64_add(base.F64_mul(v113, v79), base.F64_add(base.F64_mul(v55, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v113, v78), float64(2.5082946711645275e-11))))))
				return v125
			} else {
				if base.I32_wrap_i64(v12) != 0 {
					v46 = v12
					v48 = v29
					v49 = int32(1072693248)
					v51 = v49 + int32(_a_F_log10_0)
					v55 = base.F64_convert_i32_s(int32(base.Ui32(v51)>>(uint(int32(20))%32)) + v48)
					v57 = base.F64_mul(v55, float64(0.30102999566361177))
					v70 = base.F64_add(base.F64_reinterpret_i64(v46&int64(4294967295)|base.I64_extend_i32_u(v51&int32(_a_F_log10_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
					v73 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
					v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v70, v73)) & int64(-4294967296))
					v79 = float64(0.4342944818781689)
					v80 = base.F64_mul(v78, v79)
					v81 = base.F64_add(v57, v80)
					v86 = base.F64_div(v70, base.F64_add(v70, float64(2)))
					v87 = base.F64_mul(v86, v86)
					v88 = base.F64_mul(v87, v87)
					v113 = base.F64_add(base.F64_mul(v86, base.F64_add(v73, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v87, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v70, v78), v73))
					v125 = base.F64_add(v81, base.F64_add(base.F64_add(v80, base.F64_sub(v57, v81)), base.F64_add(base.F64_mul(v113, v79), base.F64_add(base.F64_mul(v55, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v113, v78), float64(2.5082946711645275e-11))))))
					return v125
				} else {
					return float64(0)
				}
			}
		}
	}
}
func F_logicalmsg_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+48)))
	v19 = v17 & int32(240)
	if v19 == v3 {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+56)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		F_ReorderBufferProcessXid(m, v24, v25, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v29 <= int32(0) {
				m.G0 = v13 + int32(16)
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
				if v34 != v36 {
					m.G0 = v13 + int32(16)
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v38 != 0 {
						v39 = F_filter_by_origin_cb_wrapper(m, l0, v23)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							if v39 != 0 {
								m.G0 = v13 + int32(16)
								return
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
								if v41 == int32(1) {
									v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
									v45 = F_SnapBuildProcessChange(m, v22, v25, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										if v45 == int32(0) {
											m.G0 = v13 + int32(16)
											return
										} else {
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
											if v49 != 0 {
												v56 = int32(1)
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
												v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
												if v58 == v56 {
													if v57&int32(1) != 0 {
													} else {
														v63 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v63)
													}
													m.G0 = v13 + int32(16)
													return
												} else {
													if v57&int32(1) == int32(0) {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
														if v69 == int32(0) {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
															v78 = F_MemoryContextAllocZero(m, v72, v73<<(uint(int32(2))%32)+int32(76))
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(5)
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
																v86 = v78 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v86
																*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v84
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v89
																v92 = v89 << (uint(int32(2)) % 32)
																if v92 != 0 {
																	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
																	base.MemoryCopy(m, v86, v93, v92)
																} else {
																}
																F_pg_qsort(m, v86, v89, int32(4), int32(185))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	v99 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
																	*(*int64)(unsafe.Add(mBase, uint32(v78)+44)) = v99
																	v103 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v103
																	*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v99
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+27)) = v103
																	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v78
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v110 + int32(1)
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
																	v115 = v114
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
																	v120 = v115
																	v122 = v119
																	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v129 = v33 + int32(16)
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
																	F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(16)
																		return
																	}
																}
															}
														} else {
															v115 = v69
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
															v120 = v115
															v122 = v119
															v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v129 = v33 + int32(16)
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
															F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																m.G0 = v13 + int32(16)
																return
															}
														}
													} else {
														v120 = v3
														v122 = v56
														v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v129 = v33 + int32(16)
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
														F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															m.G0 = v13 + int32(16)
															return
														}
													}
												}
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												if v50 != int32(2) {
													m.G0 = v13 + int32(16)
													return
												} else {
													v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													v54 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
													if base.Ui64(v53) < base.Ui64(v54) {
														m.G0 = v13 + int32(16)
														return
													} else {
														v56 = int32(1)
														v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
														v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
														if v58 == v56 {
															if v57&int32(1) != 0 {
															} else {
																v63 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v63)
															}
															m.G0 = v13 + int32(16)
															return
														} else {
															if v57&int32(1) == int32(0) {
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
																if v69 == int32(0) {
																	v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
																	v78 = F_MemoryContextAllocZero(m, v72, v73<<(uint(int32(2))%32)+int32(76))
																	mBase = m.M
																	v79 = m.ExcPending
																	if v79 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(5)
																		v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
																		v86 = v78 + int32(72)
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v86
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v84
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v89
																		v92 = v89 << (uint(int32(2)) % 32)
																		if v92 != 0 {
																			v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
																			base.MemoryCopy(m, v86, v93, v92)
																		} else {
																		}
																		F_pg_qsort(m, v86, v89, int32(4), int32(185))
																		mBase = m.M
																		v98 = m.ExcPending
																		if v98 != 0 {
																			return
																		} else {
																			v99 = int64(0)
																			*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
																			*(*int64)(unsafe.Add(mBase, uint32(v78)+44)) = v99
																			v103 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v103
																			*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v99
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+27)) = v103
																			*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v78
																			v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v110 + int32(1)
																			v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
																			v115 = v114
																			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
																			v120 = v115
																			v122 = v119
																			v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																			v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																			v129 = v33 + int32(16)
																			v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
																			v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
																			F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return
																			} else {
																				m.G0 = v13 + int32(16)
																				return
																			}
																		}
																	}
																} else {
																	v115 = v69
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
																	v120 = v115
																	v122 = v119
																	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v129 = v33 + int32(16)
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
																	F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(16)
																		return
																	}
																}
															} else {
																v120 = v3
																v122 = v56
																v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																v129 = v33 + int32(16)
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
																F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(16)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
									if v50 != int32(2) {
										m.G0 = v13 + int32(16)
										return
									} else {
										v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										v54 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
										if base.Ui64(v53) < base.Ui64(v54) {
											m.G0 = v13 + int32(16)
											return
										} else {
											v56 = int32(1)
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
											v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
											if v58 == v56 {
												if v57&int32(1) != 0 {
												} else {
													v63 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v63)
												}
												m.G0 = v13 + int32(16)
												return
											} else {
												if v57&int32(1) == int32(0) {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
													if v69 == int32(0) {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
														v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
														v78 = F_MemoryContextAllocZero(m, v72, v73<<(uint(int32(2))%32)+int32(76))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(5)
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
															v86 = v78 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v86
															*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v84
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v89
															v92 = v89 << (uint(int32(2)) % 32)
															if v92 != 0 {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
																base.MemoryCopy(m, v86, v93, v92)
															} else {
															}
															F_pg_qsort(m, v86, v89, int32(4), int32(185))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																v99 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
																*(*int64)(unsafe.Add(mBase, uint32(v78)+44)) = v99
																v103 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v103
																*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v99
																*(*int32)(unsafe.Add(mBase, uint32(v78)+27)) = v103
																*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v78
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v110 + int32(1)
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
																v115 = v114
																v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
																v120 = v115
																v122 = v119
																v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																v129 = v33 + int32(16)
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
																F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(16)
																	return
																}
															}
														}
													} else {
														v115 = v69
														v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
														v120 = v115
														v122 = v119
														v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v129 = v33 + int32(16)
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
														F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															m.G0 = v13 + int32(16)
															return
														}
													}
												} else {
													v120 = v3
													v122 = v56
													v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
													v129 = v33 + int32(16)
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
													F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return
													} else {
														m.G0 = v13 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
						if v41 == int32(1) {
							v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							v45 = F_SnapBuildProcessChange(m, v22, v25, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								if v45 == int32(0) {
									m.G0 = v13 + int32(16)
									return
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
									if v49 != 0 {
										v56 = int32(1)
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
										if v58 == v56 {
											if v57&int32(1) != 0 {
											} else {
												v63 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v63)
											}
											m.G0 = v13 + int32(16)
											return
										} else {
											if v57&int32(1) == int32(0) {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
												if v69 == int32(0) {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
													v78 = F_MemoryContextAllocZero(m, v72, v73<<(uint(int32(2))%32)+int32(76))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(5)
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
														v86 = v78 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v86
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v84
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v89
														v92 = v89 << (uint(int32(2)) % 32)
														if v92 != 0 {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
															base.MemoryCopy(m, v86, v93, v92)
														} else {
														}
														F_pg_qsort(m, v86, v89, int32(4), int32(185))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															v99 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
															*(*int64)(unsafe.Add(mBase, uint32(v78)+44)) = v99
															v103 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v103
															*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v99
															*(*int32)(unsafe.Add(mBase, uint32(v78)+27)) = v103
															*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v78
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v110 + int32(1)
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
															v115 = v114
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
															v120 = v115
															v122 = v119
															v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v129 = v33 + int32(16)
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
															F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																m.G0 = v13 + int32(16)
																return
															}
														}
													}
												} else {
													v115 = v69
													v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
													v120 = v115
													v122 = v119
													v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
													v129 = v33 + int32(16)
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
													F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return
													} else {
														m.G0 = v13 + int32(16)
														return
													}
												}
											} else {
												v120 = v3
												v122 = v56
												v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
												v129 = v33 + int32(16)
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
												F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													m.G0 = v13 + int32(16)
													return
												}
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
										if v50 != int32(2) {
											m.G0 = v13 + int32(16)
											return
										} else {
											v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											v54 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
											if base.Ui64(v53) < base.Ui64(v54) {
												m.G0 = v13 + int32(16)
												return
											} else {
												v56 = int32(1)
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
												v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
												if v58 == v56 {
													if v57&int32(1) != 0 {
													} else {
														v63 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v63)
													}
													m.G0 = v13 + int32(16)
													return
												} else {
													if v57&int32(1) == int32(0) {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
														if v69 == int32(0) {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
															v78 = F_MemoryContextAllocZero(m, v72, v73<<(uint(int32(2))%32)+int32(76))
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(5)
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
																v86 = v78 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v86
																*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v84
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v89
																v92 = v89 << (uint(int32(2)) % 32)
																if v92 != 0 {
																	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
																	base.MemoryCopy(m, v86, v93, v92)
																} else {
																}
																F_pg_qsort(m, v86, v89, int32(4), int32(185))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	v99 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
																	*(*int64)(unsafe.Add(mBase, uint32(v78)+44)) = v99
																	v103 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v103
																	*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v99
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+27)) = v103
																	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v78
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v110 + int32(1)
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
																	v115 = v114
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
																	v120 = v115
																	v122 = v119
																	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v129 = v33 + int32(16)
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
																	F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(16)
																		return
																	}
																}
															}
														} else {
															v115 = v69
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
															v120 = v115
															v122 = v119
															v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v129 = v33 + int32(16)
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
															F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																m.G0 = v13 + int32(16)
																return
															}
														}
													} else {
														v120 = v3
														v122 = v56
														v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v129 = v33 + int32(16)
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
														F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															m.G0 = v13 + int32(16)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
							if v50 != int32(2) {
								m.G0 = v13 + int32(16)
								return
							} else {
								v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
								v54 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
								if base.Ui64(v53) < base.Ui64(v54) {
									m.G0 = v13 + int32(16)
									return
								} else {
									v56 = int32(1)
									v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
									if v58 == v56 {
										if v57&int32(1) != 0 {
										} else {
											v63 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v63)
										}
										m.G0 = v13 + int32(16)
										return
									} else {
										if v57&int32(1) == int32(0) {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
											if v69 == int32(0) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
												v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
												v78 = F_MemoryContextAllocZero(m, v72, v73<<(uint(int32(2))%32)+int32(76))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(5)
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
													v86 = v78 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v84
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v89
													v92 = v89 << (uint(int32(2)) % 32)
													if v92 != 0 {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
														base.MemoryCopy(m, v86, v93, v92)
													} else {
													}
													F_pg_qsort(m, v86, v89, int32(4), int32(185))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														v99 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
														*(*int64)(unsafe.Add(mBase, uint32(v78)+44)) = v99
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v103
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v99
														*(*int32)(unsafe.Add(mBase, uint32(v78)+27)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v78
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v110 + int32(1)
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
														v115 = v114
														v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
														v120 = v115
														v122 = v119
														v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v129 = v33 + int32(16)
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
														F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															m.G0 = v13 + int32(16)
															return
														}
													}
												}
											} else {
												v115 = v69
												v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
												v120 = v115
												v122 = v119
												v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
												v129 = v33 + int32(16)
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
												F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													m.G0 = v13 + int32(16)
													return
												}
											}
										} else {
											v120 = v3
											v122 = v56
											v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
											v129 = v33 + int32(16)
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
											F_ReorderBufferQueueMessage(m, v124, v25, v120, v125, v122&int32(1), v129, v130, v129+v131)
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return
											} else {
												m.G0 = v13 + int32(16)
												return
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v19
			F_errmsg_internal(m, int32(_a_F_logicalmsg_decode_0), v13)
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_logicalmsg_decode_1), int32(597), int32(_a_F_logicalmsg_decode_2))
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
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
func F_logicalmsg_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	if base.Ui32(l0) < base.Ui32(int32(16)) {
		v6 = int32(_a_F_logicalmsg_identify_0)
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_logicalmsg_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+48)))
	v10 = v8 & int32(240)
	if v10 != 0 {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v10
			F_errmsg_internal(m, int32(_a_F_logicalmsg_redo_0), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_logicalmsg_redo_1), int32(92), int32(_a_F_logicalmsg_redo_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ltree2text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v18 = F_palloc(m, int32(base.Ui32(v13)>>(uint(int32(2))%32))+int32(4))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = v18 + int32(4)
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			if v22 == int32(0) {
				v69 = v21
			} else {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)))
				if v25 != 0 {
					base.MemoryCopy(m, v21, v9+int32(10), v25)
				} else {
				}
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)))
				v30 = v21 + v29
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
				if base.Ui32(v31) < base.Ui32(int32(2)) {
					v69 = v30
				} else {
					v43 = v30
					v45 = v9 + int32(8) + (v29+int32(9))&int32(_a_F_ltree2text_0)
					v48 = int32(1)
					for {
						v49 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v49)
						v52 = v43 + int32(1)
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
						if v53 != 0 {
							base.MemoryCopy(m, v52, v45+int32(2), v53)
						} else {
						}
						v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
						v58 = v52 + v57
						v65 = v48 + int32(1)
						v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
						if base.Ui32(v65) < base.Ui32(v66) {
							v43 = v58
							v45 = v45 + (v57+int32(9))&int32(_a_F_ltree2text_0)
							v48 = v65
							continue
						} else {
							break
						}
						break
					}
					v69 = v58
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = (v69 - v18) << (uint(int32(2)) % 32)
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v79 != v9 {
				F_pfree(m, v9)
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					return v18
				}
			} else {
				return v18
			}
		}
	}
}
func F_ltrim1(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13856(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ltsWriteBlock(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v8 = m.G0
	v12 = (v8 - int32(_a_F_ltsWriteBlock_0)) & int32(-4096)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 < l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = v14
	goto L4
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = F_BufFileSeekBlock(m, v39, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	v24 = v12 + int32(_a_F_ltsWriteBlock_1)
	base.MemoryFill(m, v24, int32(0), int32(_a_F_ltsWriteBlock_2))
	F_ltsWriteBlock(m, l0, v22, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v30 < l1 {
		v22 = v30
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_BufFileWrite(m, v44, l2, int32(_a_F_ltsWriteBlock_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v48 == l1 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1 + int64(1)
	goto L16
L15:
	;
	goto L16
L16:
	;
	m.G0 = v8
	return
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4080)) = l1
	F_errmsg(m, int32(_a_F_ltsWriteBlock_3), v12+int32(4080))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_ltsWriteBlock_4), int32(267), int32(_a_F_ltsWriteBlock_5))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
