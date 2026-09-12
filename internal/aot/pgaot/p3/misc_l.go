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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+616))
	if v21 != v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v30 = base.I32_div_s(v20-v27, int32(640))
	v32 = base.I32_rem_s(v30, int32(16))
	v37 = v24 + v32<<(uint(int32(7))%32) + int32(23296)
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
	v63 = int32(4449520)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v67
	v74 = F__emscripten_memset_bulkmem(m, v11+int32(12), base.I32_extend8_s(int32(0)), int32(1460))
	mBase = m.M
	goto L13
L7:
	;
	return
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[128]))
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
	v76 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
	v82 = F_pg_snprintf(m, v11+int32(12), int32(96), int32(461100), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v89 = F_pg_snprintf(m, v11+int32(108), int32(96), int32(211838), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+212)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+204)) = int64(4294967315)
	v99 = F_pg_sprintf(m, v11+int32(216), int32(153365), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v105 = F_pg_sprintf(m, v11+int32(1240), int32(267142), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+1336)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+1468)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v113 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v120 = v2
	v121 = v2
	goto L21
L19:
	;
	goto L20
L20:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v178 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+1340)) = v120
	if v121&int32(1) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L20
L23:
	;
	v167 = v120 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v167 < v168 {
		v120 = v167
		v121 = v165
		goto L21
	} else {
		goto L30
	}
L24:
	;
	v149 = v120 << (uint(int32(3)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v149+v150))) = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v149)+4))
	F_shm_mq_detach(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L29
	}
L25:
	;
	v130 = v120 << (uint(int32(3)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v133 = F_RegisterDynamicBackgroundWorker(m, v11+int32(12), v130+v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	if v133 == int32(0) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v138 = v137 + v130
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v140
	goto L28
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v142 + int32(1)
	v165 = int32(0)
	goto L23
L29:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v149)+4)) = int32(0)
	v165 = int32(1)
	goto L23
L30:
	;
	goto L22
L31:
	;
	v181 = F_palloc0(m, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v64
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v181
	goto L33
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
	var v21 int64
	_ = v21
	var v30 int32
	_ = v30
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
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
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int64
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v506 int32
	_ = v506
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v685 int32
	_ = v685
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(1744)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1372)) = v6
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(1360)))) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(1368)))) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v15)+1352)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1344)) = l0
	v30 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1348)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1340)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v30
	if l0 == v30 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(1744)
	return v685
L2:
	;
	v189 = F_pg_getaddrinfo_all(m, l1, v184, v15+int32(1340), v15+int32(1372))
	mBase = m.M
	if v189 != 0 {
		goto L38
	} else {
		goto L39
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+276)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = l3
	v46 = F_pg_snprintf(m, v15+int32(304), int32(1024), int32(448826), v15+int32(272))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v178 = F_pg_snprintf(m, v15+int32(1696), int32(32), int32(469640), v15+int32(288))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L37
	}
L6:
	;
	return int32(0)
L7:
	;
	v51 = v15 + int32(304)
	if v51&int32(3) == int32(0) {
		v75 = v51
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if base.Ui32(int32(108)) <= base.Ui32(v108) {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v108 = v100 - v51
	goto L8
L10:
	;
	v79 = v75
	goto L19
L11:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v59 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v108 = int32(0)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v64 = v51
	goto L15
L15:
	;
	v68 = v64 + int32(1)
	if v68&int32(3) == int32(0) {
		v75 = v68
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v100 = v68
	goto L9
L17:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v73 != 0 {
		v64 = v68
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v88 = int32(-2139062144)
	if (int32(16843008)-v85|v85)&v88 == v88 {
		v79 = v79 + int32(4)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v94 = v79
	goto L22
L21:
	;
	goto L20
L22:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v98 != 0 {
		v94 = v94 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v100 = v94
	goto L9
L24:
	;
	goto L23
L25:
	;
	v111 = int32(-1)
	v114 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v132 = v15 + int32(304)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+304)))
	if v133 == int32(64) {
		v184 = v132
		goto L2
	} else {
		goto L32
	}
L28:
	;
	if v114 == int32(0) {
		v685 = v111
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(107)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v15 + int32(304)
	F_errmsg(m, int32(636830), v15)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(477916), int32(459), int32(77727))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v685 = v111
	goto L1
L32:
	;
	v136 = m.G0
	v138 = v136 - int32(1040)
	m.G0 = v138
	v141 = v15 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v141
	v147 = F_pg_snprintf(m, v138+int32(16), int32(1024), int32(304396), v138)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_CreateLockFile(m, v138+int32(16), int32(1), l3, int32(0), v141)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	m.G0 = v138 + int32(1040)
	v159 = v15 + int32(304)
	v160 = F_unlink(m, v159)
	mBase = m.M
	v161 = int32(4348456)
	v163 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	v166 = F_pstrdup(m, v159)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v168 = F_lappend(m, v163, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[565])) = v168
	v184 = v132
	goto L2
L37:
	;
	v184 = v15 + int32(1696)
	goto L2
L38:
	;
	v574 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L6
	} else {
		goto L160
	}
L39:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v190 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v198 = v190
	v205 = v6
	goto L41
L41:
	;
	if base.B2i32(l0 == int32(1)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1344))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v551 == int32(1) {
		goto L149
	} else {
		goto L150
	}
L43:
	;
	goto L42
L44:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v198)+28))
	if v543 != 0 {
		v198 = v543
		v205 = v542
		goto L41
	} else {
		goto L146
	}
L45:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v209 == int32(1) {
		v542 = v205
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[566]))
	if v213 == int32(64) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v218 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v235 = v15 + int32(304)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	switch v238 - int32(1) {
	case 0:
		v273 = v238
		v274 = int32(25806)
		v276 = v235
		goto L56
	case 1:
		v259 = int32(530313)
		goto L57
	default:
		goto L58
	case 9:
		goto L59
	}
L52:
	;
	if v218 == int32(0) {
		v550 = v205
		goto L43
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(64)
	F_errmsg(m, int32(445042), v15+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(477916), int32(504), int32(77727))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v550 = v205
	goto L43
L56:
	;
	v279 = F_socket(m, v273, int32(1), int32(0))
	mBase = m.M
	if v279 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v265 = int32(0)
	v268 = F_pg_getnameinfo_all(m, v260, v261, v15+int32(1376), int32(255), v265, v265, int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L62
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v238
	v249 = F_pg_snprintf(m, v15+int32(1632), int32(64), int32(449394), v15+int32(224))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	v259 = int32(529497)
	goto L57
L60:
	;
	v251 = int32(1)
	v253 = v15 + int32(1632)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v254 == v251 {
		v273 = v251
		v274 = v253
		v276 = v235
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v259 = v253
	goto L57
L62:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v273 = v270
	v274 = v259
	v276 = v15 + int32(1376)
	goto L56
L63:
	;
	v284 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v304 == v302 {
		goto L71
	} else {
		goto L72
	}
L66:
	;
	if v284 == int32(0) {
		v542 = v205
		goto L44
	} else {
		goto L67
	}
L67:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v274
	F_errmsg(m, int32(285093), v15+int32(32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(477916), int32(547), int32(77727))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	v542 = v205
	goto L44
L71:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v312 = F_bind(m, v279, v310, v311)
	mBase = m.M
	if v312 < int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L71
L74:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v319 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v351 != int32(1) {
		goto L91
	} else {
		goto L92
	}
L77:
	;
	if v319 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L6
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v350 = F_close(m, v279)
	mBase = m.M
	v542 = v205
	goto L44
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v274
	F_errmsg(m, int32(285058), v15-int32(-64))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	if v316 == int32(3) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	if v332 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	F_errfinish(m, int32(477916), int32(624), int32(77727))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L6
	} else {
		goto L90
	}
L86:
	;
	v338 = int32(525571)
	goto L88
L87:
	;
	v338 = int32(535629)
	goto L88
L88:
	;
	F_errhint(m, v338, v15+int32(48))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	goto L80
L91:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v457 = int32(0)
	v461 = m.Env.X__syscall_listen(m, v279, v454<<(uint(int32(1))%32), v457, v457, v457, v457)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v461) {
		goto L122
	} else {
		goto L123
	}
L92:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v354 == int32(64) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v359 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v451 = F_close(m, v279)
	mBase = m.M
	v550 = v205
	goto L43
L95:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _consts[568]))
	v428 = F_chmod(m, v184, v427)
	mBase = m.M
	if v428 != int32(-1) {
		goto L91
	} else {
		goto L115
	}
L96:
	;
	v366 = F_strtox_2(m, v358, v15+int32(1740), int32(10), int64(4294967295))
	mBase = m.M
	goto L97
L97:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1740))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	if v372 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	F_errfinish(m, int32(477916), v420, int32(490235))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L114
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(44)
	v380 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v394 = m.Env.X__syscall_fchownat(m, int32(-100), v184, int32(-1), base.I32_wrap_i64(v366), int32(0))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v394) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	if v380 == int32(0) {
		goto L94
	} else {
		goto L103
	}
L103:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v385
	F_errmsg(m, int32(68491), v15+int32(160))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v420 = int32(755)
	goto L98
L105:
	;
	if v402 != int32(-1) {
		goto L95
	} else {
		goto L109
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v394
	v402 = int32(-1)
	goto L108
L107:
	;
	v402 = v394
	goto L108
L108:
	;
	goto L105
L109:
	;
	v407 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	if v407 == int32(0) {
		goto L94
	} else {
		goto L111
	}
L111:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v184
	F_errmsg(m, int32(286905), v15+int32(144))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	v420 = int32(765)
	goto L98
L114:
	;
	v424 = F_close(m, v279)
	mBase = m.M
	v550 = v205
	goto L43
L115:
	;
	v433 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	if v433 == int32(0) {
		goto L94
	} else {
		goto L117
	}
L117:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v184
	F_errmsg(m, int32(286862), v15+int32(128))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(477916), int32(776), int32(490235))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	goto L94
L121:
	;
	if v469 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v461
	v469 = int32(-1)
	goto L124
L123:
	;
	v469 = v461
	goto L124
L124:
	;
	goto L121
L125:
	;
	v474 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v494 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L6
	} else {
		goto L135
	}
L128:
	;
	if v474 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v490 = F_close(m, v279)
	mBase = m.M
	v542 = v205
	goto L44
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v274
	F_errmsg(m, int32(285018), v15+int32(80))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(477916), int32(652), int32(77727))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	if v491 == int32(1) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v523 = int32(4357996)
	v524 = *(*int32)(unsafe.Add(mBase, _consts[566]))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v524<<(uint(int32(2))%32)))) = v279
	v531 = *(*int32)(unsafe.Add(mBase, _consts[566]))
	v532 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[566])) = v531 + v532
	v542 = v205 + v532
	goto L44
L137:
	;
	F_errfinish(m, int32(477916), v519, int32(77727))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L6
	} else {
		goto L145
	}
L138:
	;
	if v494 == int32(0) {
		goto L136
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if v494 == int32(0) {
		goto L136
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v276
	F_errmsg(m, int32(661998), v15+int32(96))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	v519 = int32(660)
	goto L137
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v274
	F_errmsg(m, int32(449746), v15+int32(112))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	v519 = int32(665)
	goto L137
L145:
	;
	goto L136
L146:
	;
	v550 = v542
	goto L43
L147:
	;
	if v550 != 0 {
		goto L157
	} else {
		goto L158
	}
L148:
	;
	goto L147
L149:
	;
	if v552 == int32(0) {
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	if v552 == int32(0) {
		goto L148
	} else {
		goto L156
	}
L152:
	;
	v558 = v552
	goto L153
L153:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+28))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v558)+20))
	F_emscripten_builtin_free(m, v560)
	mBase = m.M
	F_emscripten_builtin_free(m, v558)
	mBase = m.M
	if v559 != 0 {
		v558 = v559
		goto L153
	} else {
		goto L155
	}
L154:
	;
	goto L148
L155:
	;
	goto L154
L156:
	;
	F_freeaddrinfo(m, v552)
	mBase = m.M
	goto L148
L157:
	;
	v570 = int32(0)
	goto L159
L158:
	;
	v570 = int32(-1)
	goto L159
L159:
	;
	v685 = v570
	goto L1
L160:
	;
	if l1 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v662 = int32(-1)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v663 == int32(0) {
		v685 = v662
		goto L1
	} else {
		goto L191
	}
L162:
	;
	F_errfinish(m, int32(477916), v658, int32(77727))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L6
	} else {
		goto L190
	}
L163:
	;
	if v574 == int32(0) {
		goto L161
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	if v574 == int32(0) {
		goto L161
	} else {
		goto L178
	}
L166:
	;
	v581 = int32(4031072)
	v583 = v189 + int32(1)
	if v583 == int32(0) {
		v603 = v581
		goto L168
	} else {
		goto L169
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v603 + base.B2i32(v605 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = l1
	F_errmsg(m, int32(191246), v15+int32(256))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L177
	}
L168:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	goto L167
L169:
	;
	v587 = v581
	v588 = v583
	goto L170
L170:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v589 == int32(0) {
		v603 = v587
		goto L168
	} else {
		goto L172
	}
L171:
	;
	v603 = v599
	goto L168
L172:
	;
	v593 = v587
	goto L173
L173:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+1)))
	if v597 != 0 {
		v593 = v593 + int32(1)
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v599 = v593 + int32(2)
	v601 = v588 + int32(1)
	if v601 != 0 {
		v587 = v599
		v588 = v601
		goto L170
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	goto L171
L177:
	;
	v658 = int32(478)
	goto L162
L178:
	;
	v622 = int32(4031072)
	v624 = v189 + int32(1)
	if v624 == int32(0) {
		v644 = v622
		goto L180
	} else {
		goto L181
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v644 + base.B2i32(v646 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v184
	F_errmsg(m, int32(191198), v15+int32(240))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L189
	}
L180:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644))))
	goto L179
L181:
	;
	v628 = v622
	v629 = v624
	goto L182
L182:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v630 == int32(0) {
		v644 = v628
		goto L180
	} else {
		goto L184
	}
L183:
	;
	v644 = v640
	goto L180
L184:
	;
	v634 = v628
	goto L185
L185:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+1)))
	if v638 != 0 {
		v634 = v634 + int32(1)
		goto L185
	} else {
		goto L187
	}
L186:
	;
	v640 = v634 + int32(2)
	v642 = v629 + int32(1)
	if v642 != 0 {
		v628 = v640
		v629 = v642
		goto L182
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	goto L183
L189:
	;
	v658 = int32(482)
	goto L162
L190:
	;
	goto L161
L191:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1344))
	if v666 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v685 = v662
	goto L1
L193:
	;
	goto L192
L194:
	;
	if v663 == int32(0) {
		goto L193
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if v663 == int32(0) {
		goto L193
	} else {
		goto L201
	}
L197:
	;
	v672 = v663
	goto L198
L198:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v672)+28))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v672)+20))
	F_emscripten_builtin_free(m, v674)
	mBase = m.M
	F_emscripten_builtin_free(m, v672)
	mBase = m.M
	if v673 != 0 {
		v672 = v673
		goto L198
	} else {
		goto L200
	}
L199:
	;
	goto L193
L200:
	;
	goto L199
L201:
	;
	F_freeaddrinfo(m, v663)
	mBase = m.M
	goto L193
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
	return base.B2i32(l0 != int32(0)) & base.B2i32(l0 != int32(4044232)) & base.B2i32(l0 != int32(4044256)) & base.B2i32(l0 != int32(4613748)) & base.B2i32(l0 != int32(4613772))
}
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	if l0 == int32(0) {
		v9 = F_palloc(m, int32(32))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(4294967768)
			v18 = v9 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v18
			v71 = v9
			v72 = v18
			v73 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v22 <= v21 {
			v24 = int32(1)
			v26 = int32(16)
			v28 = v21 + v24
			if v28 <= v26 {
				v31 = v26
			} else {
				v31 = v28
			}
			if v31&(v31-int32(1)) != 0 {
				v38 = v24 << (uint(int32(32)-base.I32_clz(v31)) % 32)
			} else {
				v38 = v31
			}
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v41 = l0 + int32(16)
			if v39 == v41 {
				v43 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v47 = F_MemoryContextAlloc(m, v43, v38<<(uint(int32(2))%32))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v52 = v50 << (uint(int32(2)) % 32)
						if v52 != 0 {
							v53 = F__emscripten_memcpy_bulkmem(m, v47, v41, v52)
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v66 = v62
						v68 = v66 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v71 = l0
						v72 = v70
						v73 = v68
						*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
						return v71
					}
				}
			} else {
				v57 = F_repalloc(m, v39, v38<<(uint(int32(2))%32))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v66 = v62
					v68 = v66 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v71 = l0
					v72 = v70
					v73 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
					return v71
				}
			}
		} else {
			v66 = v21
			v68 = v66 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v71 = l0
			v72 = v70
			v73 = v68
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	}
}
func F_latin2_to_win1250(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(9), int32(29))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(9), int32(29), int32(2188000), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
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
func F_ldexp(m *base.Module, l0 float64, l1 int32) float64 {
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 float64
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	if int32(1024) <= l1 {
		v6 = base.F64_mul(l0, float64(8.98846567431158e+307))
		if base.Ui32(l1) < base.Ui32(int32(2047)) {
			v35 = v6
			v36 = l1 - int32(1023)
		} else {
			v13 = int32(3069)
			if base.Ui32(v13) <= base.Ui32(l1) {
				v16 = v13
			} else {
				v16 = l1
			}
			v35 = base.F64_mul(v6, float64(8.98846567431158e+307))
			v36 = v16 - int32(2046)
		}
	} else {
		if int32(-1023) < l1 {
			v35 = l0
			v36 = l1
		} else {
			v22 = base.F64_mul(l0, float64(2.004168360008973e-292))
			if base.Ui32(int32(-1992)) < base.Ui32(l1) {
				v35 = v22
				v36 = l1 + int32(969)
			} else {
				v29 = int32(-2960)
				if base.Ui32(l1) <= base.Ui32(v29) {
					v32 = v29
				} else {
					v32 = l1
				}
				v35 = base.F64_mul(v22, float64(2.004168360008973e-292))
				v36 = v32 + int32(1938)
			}
		}
	}
	return base.F64_mul(v35, base.F64_reinterpret_i64(base.I64_extend_i32_u(v36+int32(1023))<<(uint(int64(52))%64)))
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(4))))
	if v9 == v2 {
		return int32(0)
	} else {
		v15 = v9 & int32(3)
		if base.Ui32(v9) < base.Ui32(int32(4)) {
			v49 = l0
			v50 = int32(0)
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
			v49 = v45
			v50 = v43
		}
		if v15 != 0 {
			v55 = v49
			v56 = v50
			v58 = v2
			for {
				v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55))))
				v64 = v56 + base.B2i32(int32(-65) < v61)
				v65 = int32(1)
				v68 = v58 + v65
				if v68 != v15 {
					v55 = v55 + v65
					v56 = v64
					v58 = v68
					continue
				} else {
					break
				}
				break
			}
			v71 = v64
		} else {
			v71 = v50
		}
		return v71
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
		goto L10
	} else {
		goto L11
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
	if l2 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L8:
	;
	v77 = int32(0)
	goto L7
L9:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L19
L10:
	;
	if (v14|v15)&int32(3) != 0 {
		v46 = v14
		v47 = v15
		v48 = v6
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v39 = v14
	v40 = v15
	v41 = v6
	goto L12
L12:
	;
	if v41 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v23 = v14
	v24 = v15
	v25 = v6
	goto L14
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L12
L16:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L9
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v77 = v56 - v57
	goto L7
L21:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L8
L25:
	;
	return v88
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if base.Ui32(v85) < base.Ui32(v84) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	if v77 != 0 {
		v88 = v77
		goto L25
	} else {
		goto L30
	}
L28:
	;
	if v77 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v80)
	goto L26
L30:
	;
	goto L26
L31:
	;
	v87 = int32(-1)
	goto L33
L32:
	;
	v87 = int32(1)
	goto L33
L33:
	;
	v88 = v87
	goto L25
}
func F_like_regex_support(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v6 - int32(458) {
	case 0:
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v10 != 0 {
			v22 = float64(0.005)
			*(*float64)(unsafe.Add(mBase, uint32(l0)+40)) = v22
			return l0
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v12 = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v18 = F_patternsel_common(m, v11, v12, v13, v14, v15, v16, l1, v12)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = v18
				*(*float64)(unsafe.Add(mBase, uint32(l0)+40)) = v22
				return l0
			}
		}
	default:
		v51 = v3
		return v51
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v25 != 0 {
			v51 = v3
			return v51
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v26 == int32(0) {
				v51 = v3
				return v51
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				switch v29 - int32(15) {
				case 0:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v49 = F_match_pattern_prefix(m, v44, v45, l1, v46, v47, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = v49
						return v51
					}
				default:
					v51 = v3
					return v51
				case 2:
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v39 = F_match_pattern_prefix(m, v34, v35, l1, v36, v37, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v39
					}
				}
			}
		}
	}
}
func F_likesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v2, v3, v4, v5, v6, v7, v4, v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(534646)
	goto L3
L1:
	;
	if v48-v49 != 0 {
		goto L15
	} else {
		goto L16
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
	v44 = v11
	v48 = int32(0)
	goto L7
L7:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	goto L1
L8:
	;
	v44 = v39
	v48 = v41
	goto L7
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v22 != v24 {
		v39 = v20
		v41 = v22
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v39 = v33
	v41 = int32(0)
	goto L8
L11:
	;
	if v24 == int32(0) {
		v39 = v20
		v41 = v22
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v29 = v21 - int32(1)
	if v29 == int32(0) {
		v39 = v20
		v41 = v22
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v32 = int32(1)
	v33 = v20 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v34 != 0 {
		v19 = v19 + v32
		v20 = v33
		v21 = v29
		v22 = v34
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v73 = l0
	goto L17
L16:
	;
	v58 = l0 + int32(8)
	v60 = v58
	goto L19
L17:
	;
	v74 = F_expand_dynamic_library_name(m, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	if v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	v70 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v62 == int32(47) {
		v70 = v60
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v60 = v60 + int32(1)
	goto L19
L26:
	;
	v71 = l0
	goto L28
L27:
	;
	v71 = v58
	goto L28
L28:
	;
	v73 = v71
	goto L17
L29:
	;
	return int32(0)
L30:
	;
	v78 = F_internal_load_library(m, v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if l3 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v78
	goto L34
L33:
	;
	goto L34
L34:
	;
	v81 = F_pgmem_dlsym(m, v78, l1)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	if l2 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_pfree(m, v74)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L29
	} else {
		goto L43
	}
L37:
	;
	if v81 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg(m, int32(679043), v9)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(476042), int32(134), int32(240806))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	m.G0 = v9 + int32(16)
	return v81
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
									F_errmsg_internal(m, int32(37763), v11+int32(16))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_errfinish(m, int32(479911), int32(1040), int32(231597))
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
								v41 = *(*int32)(unsafe.Add(mBase, _consts[499]))
								F_fmgr_info_cxt(m, v34, l0+int32(212), v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									if v24 != 0 {
										v47 = *(*int32)(unsafe.Add(mBase, _consts[499]))
										F_fmgr_info_cxt(m, v24, l0+int32(240), v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											if v23 != 0 {
												v53 = *(*int32)(unsafe.Add(mBase, _consts[499]))
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
											v53 = *(*int32)(unsafe.Add(mBase, _consts[499]))
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
				F_errmsg_internal(m, int32(48526), v11)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(479911), int32(1020), int32(231597))
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(-1)
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1043), v6+int32(8), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		m.G0 = v6 + int32(16)
		return v19
	}
}
func F_locate_var_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(-1)
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(901), v6+int32(8), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		m.G0 = v6 + int32(16)
		return v19
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
				F_errmsg_internal(m, int32(468503), v7)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(478432), int32(4554), int32(95241))
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
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_consts[848])))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			F_LockRefindAndRelease(m, v36, v10, l2, v37, int32(1))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
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
				v51 = v49 + int32(614242)
				v55 = base.F64_convert_i32_s(int32(base.Ui32(v51)>>(uint(int32(20))%32)) + v48)
				v57 = base.F64_mul(v55, float64(0.30102999566361177))
				v70 = base.F64_add(base.F64_reinterpret_i64(v46&int64(4294967295)|base.I64_extend_i32_u(v51&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
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
				v51 = v49 + int32(614242)
				v55 = base.F64_convert_i32_s(int32(base.Ui32(v51)>>(uint(int32(20))%32)) + v48)
				v57 = base.F64_mul(v55, float64(0.30102999566361177))
				v70 = base.F64_add(base.F64_reinterpret_i64(v46&int64(4294967295)|base.I64_extend_i32_u(v51&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
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
					v51 = v49 + int32(614242)
					v55 = base.F64_convert_i32_s(int32(base.Ui32(v51)>>(uint(int32(20))%32)) + v48)
					v57 = base.F64_mul(v55, float64(0.30102999566361177))
					v70 = base.F64_add(base.F64_reinterpret_i64(v46&int64(4294967295)|base.I64_extend_i32_u(v51&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	v18 = v16 & int32(240)
	if v18 == v3 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+56)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		F_ReorderBufferProcessXid(m, v23, v24, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v28 <= int32(0) {
				m.G0 = v12 + int32(16)
				return
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
				if v33 != v35 {
					m.G0 = v12 + int32(16)
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v37 != 0 {
						v38 = F_filter_by_origin_cb_wrapper(m, l0, v22)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							if v38 != 0 {
								m.G0 = v12 + int32(16)
								return
							} else {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
								if v40 == int32(1) {
									v43 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
									v44 = F_SnapBuildProcessChange(m, v21, v24, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										if v44 == int32(0) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
											if v48&int32(1) == int32(0) {
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
												if v53 != int32(2) {
													m.G0 = v12 + int32(16)
													return
												} else {
													v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
													v57 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
													if base.Ui64(v56) < base.Ui64(v57) {
														m.G0 = v12 + int32(16)
														return
													} else {
														v59 = int32(1)
														v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
														v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
														if v61 == v59 {
															if v60&int32(1) != 0 {
															} else {
																v66 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
															}
															m.G0 = v12 + int32(16)
															return
														} else {
															if v60&int32(1) == int32(0) {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
																if v72 != 0 {
																	v120 = v72
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
																	v123 = v121
																	v124 = v120
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v130 = v32 + int32(16)
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
																	F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
																		return
																	}
																} else {
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
																	v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
																	v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
																	mBase = m.M
																	v80 = m.ExcPending
																	if v80 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
																		v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
																		v87 = v79 + int32(72)
																		*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
																		*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
																		v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
																		*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
																		v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
																		v94 = v90 << (uint(int32(2)) % 32)
																		if v94 != 0 {
																			v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
																			mBase = m.M
																			v96 = v95
																		} else {
																			v96 = v87
																		}
																		F_pg_qsort(m, v96, v90, int32(4), int32(185))
																		mBase = m.M
																		v100 = m.ExcPending
																		if v100 != 0 {
																			return
																		} else {
																			v101 = int64(0)
																			*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
																			*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
																			v105 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
																			*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
																			*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
																			*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
																			*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
																			v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
																			v120 = v116
																			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
																			v123 = v121
																			v124 = v120
																			v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																			v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																			v130 = v32 + int32(16)
																			v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
																			v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
																			F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return
																			} else {
																				m.G0 = v12 + int32(16)
																				return
																			}
																		}
																	}
																}
															} else {
																v123 = v59
																v124 = v3
																v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																v130 = v32 + int32(16)
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
																F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(16)
																	return
																}
															}
														}
													}
												}
											} else {
												v59 = int32(1)
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
												if v61 == v59 {
													if v60&int32(1) != 0 {
													} else {
														v66 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
													}
													m.G0 = v12 + int32(16)
													return
												} else {
													if v60&int32(1) == int32(0) {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
														if v72 != 0 {
															v120 = v72
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
															v123 = v121
															v124 = v120
															v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v130 = v32 + int32(16)
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
															F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														} else {
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
															v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
															v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
																v87 = v79 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
																*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
																v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
																*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
																v94 = v90 << (uint(int32(2)) % 32)
																if v94 != 0 {
																	v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
																	mBase = m.M
																	v96 = v95
																} else {
																	v96 = v87
																}
																F_pg_qsort(m, v96, v90, int32(4), int32(185))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	v101 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
																	v105 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
																	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
																	v120 = v116
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
																	v123 = v121
																	v124 = v120
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v130 = v32 + int32(16)
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
																	F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
																		return
																	}
																}
															}
														}
													} else {
														v123 = v59
														v124 = v3
														v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v130 = v32 + int32(16)
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
														F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												}
											}
										}
									}
								} else {
									if v40 != 0 {
										v59 = int32(1)
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
										if v61 == v59 {
											if v60&int32(1) != 0 {
											} else {
												v66 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
											}
											m.G0 = v12 + int32(16)
											return
										} else {
											if v60&int32(1) == int32(0) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
												if v72 != 0 {
													v120 = v72
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
													v123 = v121
													v124 = v120
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
													v130 = v32 + int32(16)
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
													F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
													v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
														v87 = v79 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
														*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
														*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
														v94 = v90 << (uint(int32(2)) % 32)
														if v94 != 0 {
															v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
															mBase = m.M
															v96 = v95
														} else {
															v96 = v87
														}
														F_pg_qsort(m, v96, v90, int32(4), int32(185))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															v101 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
															*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
															v105 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
															*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
															*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
															v120 = v116
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
															v123 = v121
															v124 = v120
															v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v130 = v32 + int32(16)
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
															F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v123 = v59
												v124 = v3
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
												v130 = v32 + int32(16)
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
												F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										}
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
										if v53 != int32(2) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											v57 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
											if base.Ui64(v56) < base.Ui64(v57) {
												m.G0 = v12 + int32(16)
												return
											} else {
												v59 = int32(1)
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
												if v61 == v59 {
													if v60&int32(1) != 0 {
													} else {
														v66 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
													}
													m.G0 = v12 + int32(16)
													return
												} else {
													if v60&int32(1) == int32(0) {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
														if v72 != 0 {
															v120 = v72
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
															v123 = v121
															v124 = v120
															v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v130 = v32 + int32(16)
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
															F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														} else {
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
															v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
															v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
																v87 = v79 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
																*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
																v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
																*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
																v94 = v90 << (uint(int32(2)) % 32)
																if v94 != 0 {
																	v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
																	mBase = m.M
																	v96 = v95
																} else {
																	v96 = v87
																}
																F_pg_qsort(m, v96, v90, int32(4), int32(185))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	v101 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
																	v105 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
																	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
																	v120 = v116
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
																	v123 = v121
																	v124 = v120
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v130 = v32 + int32(16)
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
																	F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
																		return
																	}
																}
															}
														}
													} else {
														v123 = v59
														v124 = v3
														v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v130 = v32 + int32(16)
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
														F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
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
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
						if v40 == int32(1) {
							v43 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							v44 = F_SnapBuildProcessChange(m, v21, v24, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								if v44 == int32(0) {
									m.G0 = v12 + int32(16)
									return
								} else {
									v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
									if v48&int32(1) == int32(0) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
										if v53 != int32(2) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											v57 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
											if base.Ui64(v56) < base.Ui64(v57) {
												m.G0 = v12 + int32(16)
												return
											} else {
												v59 = int32(1)
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
												if v61 == v59 {
													if v60&int32(1) != 0 {
													} else {
														v66 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
													}
													m.G0 = v12 + int32(16)
													return
												} else {
													if v60&int32(1) == int32(0) {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
														if v72 != 0 {
															v120 = v72
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
															v123 = v121
															v124 = v120
															v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v130 = v32 + int32(16)
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
															F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														} else {
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
															v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
															v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
																v87 = v79 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
																*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
																v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
																*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
																v94 = v90 << (uint(int32(2)) % 32)
																if v94 != 0 {
																	v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
																	mBase = m.M
																	v96 = v95
																} else {
																	v96 = v87
																}
																F_pg_qsort(m, v96, v90, int32(4), int32(185))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	v101 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
																	v105 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
																	*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
																	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
																	*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
																	v120 = v116
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
																	v123 = v121
																	v124 = v120
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
																	v130 = v32 + int32(16)
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
																	F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
																		return
																	}
																}
															}
														}
													} else {
														v123 = v59
														v124 = v3
														v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
														v130 = v32 + int32(16)
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
														F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										v59 = int32(1)
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
										if v61 == v59 {
											if v60&int32(1) != 0 {
											} else {
												v66 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
											}
											m.G0 = v12 + int32(16)
											return
										} else {
											if v60&int32(1) == int32(0) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
												if v72 != 0 {
													v120 = v72
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
													v123 = v121
													v124 = v120
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
													v130 = v32 + int32(16)
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
													F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
													v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
														v87 = v79 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
														*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
														*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
														v94 = v90 << (uint(int32(2)) % 32)
														if v94 != 0 {
															v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
															mBase = m.M
															v96 = v95
														} else {
															v96 = v87
														}
														F_pg_qsort(m, v96, v90, int32(4), int32(185))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															v101 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
															*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
															v105 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
															*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
															*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
															v120 = v116
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
															v123 = v121
															v124 = v120
															v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v130 = v32 + int32(16)
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
															F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v123 = v59
												v124 = v3
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
												v130 = v32 + int32(16)
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
												F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						} else {
							if v40 != 0 {
								v59 = int32(1)
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
								if v61 == v59 {
									if v60&int32(1) != 0 {
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
									}
									m.G0 = v12 + int32(16)
									return
								} else {
									if v60&int32(1) == int32(0) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
										if v72 != 0 {
											v120 = v72
											v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
											v123 = v121
											v124 = v120
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
											v130 = v32 + int32(16)
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
											F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
											v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
												v87 = v79 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
												*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
												*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
												v94 = v90 << (uint(int32(2)) % 32)
												if v94 != 0 {
													v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
													mBase = m.M
													v96 = v95
												} else {
													v96 = v87
												}
												F_pg_qsort(m, v96, v90, int32(4), int32(185))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													v101 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
													*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
													v105 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
													*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
													v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
													*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
													v120 = v116
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
													v123 = v121
													v124 = v120
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
													v130 = v32 + int32(16)
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
													F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v123 = v59
										v124 = v3
										v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
										v130 = v32 + int32(16)
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
										F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v53 != int32(2) {
									m.G0 = v12 + int32(16)
									return
								} else {
									v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
									v57 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
									if base.Ui64(v56) < base.Ui64(v57) {
										m.G0 = v12 + int32(16)
										return
									} else {
										v59 = int32(1)
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
										if v61 == v59 {
											if v60&int32(1) != 0 {
											} else {
												v66 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+165)) = uint8(v66)
											}
											m.G0 = v12 + int32(16)
											return
										} else {
											if v60&int32(1) == int32(0) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
												if v72 != 0 {
													v120 = v72
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
													v123 = v121
													v124 = v120
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
													v130 = v32 + int32(16)
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
													F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
													v79 = F_MemoryContextAllocZero(m, v73, v74<<(uint(int32(2))%32)+int32(76))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(5)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v83
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
														v87 = v79 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v87
														*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v85
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
														*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v90
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
														v94 = v90 << (uint(int32(2)) % 32)
														if v94 != 0 {
															v95 = F__emscripten_memcpy_bulkmem(m, v87, v92, v94)
															mBase = m.M
															v96 = v95
														} else {
															v96 = v87
														}
														F_pg_qsort(m, v96, v90, int32(4), int32(185))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															v101 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v79)+64)) = v101
															*(*int64)(unsafe.Add(mBase, uint32(v79)+44)) = v101
															v105 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v101
															*(*int32)(unsafe.Add(mBase, uint32(v79)+27)) = v105
															*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v79
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v112 + int32(1)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
															v120 = v116
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
															v123 = v121
															v124 = v120
															v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
															v130 = v32 + int32(16)
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
															F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v123 = v59
												v124 = v3
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
												v130 = v32 + int32(16)
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
												F_ReorderBufferQueueMessage(m, v125, v24, v124, v126, v123&int32(1), v130, v131, v130+v132)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
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
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v18
			F_errmsg_internal(m, int32(56695), v12)
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return
			} else {
				F_errfinish(m, int32(480196), int32(597), int32(397557))
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
		v6 = int32(521551)
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
			F_errmsg_internal(m, int32(50599), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(479993), int32(92), int32(232334))
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
func F_ltrim1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(1)
		v11 = v6 + v10
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v16 = v14 & v10
		if v16 != 0 {
			v17 = v11
		} else {
			v17 = v6 + int32(4)
		}
		if v14 == int32(1) {
			v20 = int32(4)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v22&int32(254) == int32(2) {
				v31 = v20
			} else {
				v31 = base.B2i32(v22 == int32(18)) << (uint(v20) % 32)
			}
			if v22 == int32(1) {
				v34 = v20
			} else {
				v34 = v31
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v16 != 0 {
				v45 = int32(base.Ui32(v14)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = int32(1)
		v50 = F_dotrim(m, v17, v45, int32(708501), v47, v47, int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			return v50
		}
	}
}
func F_ltsWriteBlock(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
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
	v7 = m.G0
	v11 = (v7 - int32(12288)) & int32(-4096)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 < l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v13
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
		goto L7
	} else {
		goto L10
	}
L4:
	;
	v26 = F__emscripten_memset_bulkmem(m, v11+int32(4096), base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L6
L5:
	;
	goto L3
L6:
	;
	F_ltsWriteBlock(m, l0, v20, v11+int32(4096))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v31 < l1 {
		v20 = v31
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	if v40 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_BufFileWrite(m, v44, l2, int32(8192))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L18
	}
L14:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v48 == l1 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1 + int64(1)
	goto L17
L16:
	;
	goto L17
L17:
	;
	m.G0 = v7
	return
L18:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4080)) = l1
	F_errmsg(m, int32(371350), v11+int32(4080))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(479547), int32(267), int32(304332))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
