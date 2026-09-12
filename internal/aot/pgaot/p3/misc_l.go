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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+616))
	if v21 != v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[150]))
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
	v63 = int32(4520272)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v67
	v74 = F__emscripten_memset_bulkmem(m, v11+int32(12), base.I32_extend8_s(int32(0)), int32(1460))
	mBase = m.M
	goto L13
L7:
	;
	return
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[137]))
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
	v76 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
	v82 = F_pg_snprintf(m, v11+int32(12), int32(96), int32(480694), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v89 = F_pg_snprintf(m, v11+int32(108), int32(96), int32(221425), int32(0))
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
	v99 = F_pg_sprintf(m, v11+int32(216), int32(162142), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v105 = F_pg_sprintf(m, v11+int32(1240), int32(279264), int32(0))
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
	v111 = *(*int32)(unsafe.Add(mBase, _consts[138]))
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int64
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
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
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v629 int32
	_ = v629
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
	return v629
L2:
	;
	v133 = F_pg_getaddrinfo_all(m, l1, v128, v15+int32(1340), v15+int32(1372))
	mBase = m.M
	if v133 != 0 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+276)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = l3
	v46 = F_pg_snprintf(m, v15+int32(304), int32(1024), int32(467827), v15+int32(272))
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
	v122 = F_pg_snprintf(m, v15+int32(1696), int32(32), int32(489370), v15+int32(288))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L20
	}
L6:
	;
	return int32(0)
L7:
	;
	v52 = F_strlen(m, v15+int32(304))
	mBase = m.M
	if base.Ui32(int32(108)) <= base.Ui32(v52) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v55 = int32(-1)
	v58 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v76 = v15 + int32(304)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+304)))
	if v77 == int32(64) {
		v128 = v76
		goto L2
	} else {
		goto L15
	}
L11:
	;
	if v58 == int32(0) {
		v629 = v55
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(107)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v15 + int32(304)
	F_errmsg(m, int32(675444), v15)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(498308), int32(459), int32(81656))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v629 = v55
	goto L1
L15:
	;
	v80 = m.G0
	v82 = v80 - int32(1040)
	m.G0 = v82
	v85 = v15 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v85
	v91 = F_pg_snprintf(m, v82+int32(16), int32(1024), int32(317999), v82)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_CreateLockFile(m, v82+int32(16), int32(1), l3, int32(0), v85)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v82 + int32(1040)
	v103 = v15 + int32(304)
	v104 = F_unlink(m, v103)
	mBase = m.M
	v105 = int32(4419208)
	v107 = *(*int32)(unsafe.Add(mBase, _consts[574]))
	v110 = F_pstrdup(m, v103)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v112 = F_lappend(m, v107, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[574])) = v112
	v128 = v76
	goto L2
L20:
	;
	v128 = v15 + int32(1696)
	goto L2
L21:
	;
	v518 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L143
	}
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v134 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v142 = v134
	v149 = v6
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
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1344))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v495 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L26:
	;
	goto L25
L27:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v142)+28))
	if v487 != 0 {
		v142 = v487
		v149 = v486
		goto L24
	} else {
		goto L129
	}
L28:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v153 == int32(1) {
		v486 = v149
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[575]))
	if v157 == int32(64) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v162 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v179 = v15 + int32(304)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	switch v182 - int32(1) {
	case 0:
		v217 = v182
		v218 = int32(27286)
		v220 = v179
		goto L39
	case 1:
		v203 = int32(558985)
		goto L40
	default:
		goto L41
	case 9:
		goto L42
	}
L35:
	;
	if v162 == int32(0) {
		v494 = v149
		goto L26
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(64)
	F_errmsg(m, int32(463765), v15+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(498308), int32(504), int32(81656))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v494 = v149
	goto L26
L39:
	;
	v223 = F_socket(m, v217, int32(1), int32(0))
	mBase = m.M
	if v223 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	v209 = int32(0)
	v212 = F_pg_getnameinfo_all(m, v204, v205, v15+int32(1376), int32(255), v209, v209, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v182
	v193 = F_pg_snprintf(m, v15+int32(1632), int32(64), int32(468395), v15+int32(224))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	v203 = int32(557687)
	goto L40
L43:
	;
	v195 = int32(1)
	v197 = v15 + int32(1632)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v198 == v195 {
		v217 = v195
		v218 = v197
		v220 = v179
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v203 = v197
	goto L40
L45:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v217 = v214
	v218 = v203
	v220 = v15 + int32(1376)
	goto L39
L46:
	;
	v228 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v246 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v248 == v246 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	if v228 == int32(0) {
		v486 = v149
		goto L27
	} else {
		goto L50
	}
L50:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v218
	F_errmsg(m, int32(297922), v15+int32(32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(498308), int32(547), int32(81656))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v486 = v149
	goto L27
L54:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	v256 = F_bind(m, v223, v254, v255)
	mBase = m.M
	if v256 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L57:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v263 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v295 != int32(1) {
		goto L74
	} else {
		goto L75
	}
L60:
	;
	if v263 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v294 = F_close(m, v223)
	mBase = m.M
	v486 = v149
	goto L27
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v218
	F_errmsg(m, int32(297887), v15-int32(-64))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	if v260 == int32(3) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	if v276 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	F_errfinish(m, int32(498308), int32(624), int32(81656))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L73
	}
L69:
	;
	v282 = int32(546858)
	goto L71
L70:
	;
	v282 = int32(573040)
	goto L71
L71:
	;
	F_errhint(m, v282, v15+int32(48))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	v398 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v401 = int32(0)
	v405 = m.Env.X__syscall_listen(m, v223, v398<<(uint(int32(1))%32), v401, v401, v401, v401)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v405) {
		goto L105
	} else {
		goto L106
	}
L75:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v298 == int32(64) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[576]))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v303 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v395 = F_close(m, v223)
	mBase = m.M
	v494 = v149
	goto L26
L78:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[577]))
	v372 = F_chmod(m, v128, v371)
	mBase = m.M
	if v372 != int32(-1) {
		goto L74
	} else {
		goto L98
	}
L79:
	;
	v310 = F_strtox_2(m, v302, v15+int32(1740), int32(10), int64(4294967295))
	mBase = m.M
	goto L80
L80:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1740))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v316 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_errfinish(m, int32(498308), v364, int32(511155))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L97
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(44)
	v324 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v338 = m.Env.X__syscall_fchownat(m, int32(-100), v128, int32(-1), base.I32_wrap_i64(v310), int32(0))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v338) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	if v324 == int32(0) {
		goto L77
	} else {
		goto L86
	}
L86:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _consts[576]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v329
	F_errmsg(m, int32(71460), v15+int32(160))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v364 = int32(755)
	goto L81
L88:
	;
	if v346 != int32(-1) {
		goto L78
	} else {
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v338
	v346 = int32(-1)
	goto L91
L90:
	;
	v346 = v338
	goto L91
L91:
	;
	goto L88
L92:
	;
	v351 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	if v351 == int32(0) {
		goto L77
	} else {
		goto L94
	}
L94:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v128
	F_errmsg(m, int32(299772), v15+int32(144))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v364 = int32(765)
	goto L81
L97:
	;
	v368 = F_close(m, v223)
	mBase = m.M
	v494 = v149
	goto L26
L98:
	;
	v377 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	if v377 == int32(0) {
		goto L77
	} else {
		goto L100
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v128
	F_errmsg(m, int32(299729), v15+int32(128))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(498308), int32(776), int32(511155))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	goto L77
L104:
	;
	if v413 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v405
	v413 = int32(-1)
	goto L107
L106:
	;
	v413 = v405
	goto L107
L107:
	;
	goto L104
L108:
	;
	v418 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L6
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v438 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L6
	} else {
		goto L118
	}
L111:
	;
	if v418 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L6
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v434 = F_close(m, v223)
	mBase = m.M
	v486 = v149
	goto L27
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v218
	F_errmsg(m, int32(297847), v15+int32(80))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(498308), int32(652), int32(81656))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	if v435 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v467 = int32(4428748)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[575]))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v468<<(uint(int32(2))%32)))) = v223
	v475 = *(*int32)(unsafe.Add(mBase, _consts[575]))
	v476 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[575])) = v475 + v476
	v486 = v149 + v476
	goto L27
L120:
	;
	F_errfinish(m, int32(498308), v463, int32(81656))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L128
	}
L121:
	;
	if v438 == int32(0) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if v438 == int32(0) {
		goto L119
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v220
	F_errmsg(m, int32(700862), v15+int32(96))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	v463 = int32(660)
	goto L120
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v218
	F_errmsg(m, int32(468747), v15+int32(112))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	v463 = int32(665)
	goto L120
L128:
	;
	goto L119
L129:
	;
	v494 = v486
	goto L26
L130:
	;
	if v494 != 0 {
		goto L140
	} else {
		goto L141
	}
L131:
	;
	goto L130
L132:
	;
	if v496 == int32(0) {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v496 == int32(0) {
		goto L131
	} else {
		goto L139
	}
L135:
	;
	v502 = v496
	goto L136
L136:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+28))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v502)+20))
	F_emscripten_builtin_free(m, v504)
	mBase = m.M
	F_emscripten_builtin_free(m, v502)
	mBase = m.M
	if v503 != 0 {
		v502 = v503
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
	F_freeaddrinfo(m, v496)
	mBase = m.M
	goto L131
L140:
	;
	v514 = int32(0)
	goto L142
L141:
	;
	v514 = int32(-1)
	goto L142
L142:
	;
	v629 = v514
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
	v606 = int32(-1)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1372))
	if v607 == int32(0) {
		v629 = v606
		goto L1
	} else {
		goto L174
	}
L145:
	;
	F_errfinish(m, int32(498308), v602, int32(81656))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L6
	} else {
		goto L173
	}
L146:
	;
	if v518 == int32(0) {
		goto L144
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if v518 == int32(0) {
		goto L144
	} else {
		goto L161
	}
L149:
	;
	v525 = int32(4088496)
	v527 = v133 + int32(1)
	if v527 == int32(0) {
		v547 = v525
		goto L151
	} else {
		goto L152
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v547 + base.B2i32(v549 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = l1
	F_errmsg(m, int32(200417), v15+int32(256))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L160
	}
L151:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	goto L150
L152:
	;
	v531 = v525
	v532 = v527
	goto L153
L153:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
	if v533 == int32(0) {
		v547 = v531
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v547 = v543
	goto L151
L155:
	;
	v537 = v531
	goto L156
L156:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+1)))
	if v541 != 0 {
		v537 = v537 + int32(1)
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v543 = v537 + int32(2)
	v545 = v532 + int32(1)
	if v545 != 0 {
		v531 = v543
		v532 = v545
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
	v602 = int32(478)
	goto L145
L161:
	;
	v566 = int32(4088496)
	v568 = v133 + int32(1)
	if v568 == int32(0) {
		v588 = v566
		goto L163
	} else {
		goto L164
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v588 + base.B2i32(v590 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v128
	F_errmsg(m, int32(200369), v15+int32(240))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L6
	} else {
		goto L172
	}
L163:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
	goto L162
L164:
	;
	v572 = v566
	v573 = v568
	goto L165
L165:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	if v574 == int32(0) {
		v588 = v572
		goto L163
	} else {
		goto L167
	}
L166:
	;
	v588 = v584
	goto L163
L167:
	;
	v578 = v572
	goto L168
L168:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+1)))
	if v582 != 0 {
		v578 = v578 + int32(1)
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v584 = v578 + int32(2)
	v586 = v573 + int32(1)
	if v586 != 0 {
		v572 = v584
		v573 = v586
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
	v602 = int32(482)
	goto L145
L173:
	;
	goto L144
L174:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1344))
	if v610 == int32(1) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v629 = v606
	goto L1
L176:
	;
	goto L175
L177:
	;
	if v607 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	if v607 == int32(0) {
		goto L176
	} else {
		goto L184
	}
L180:
	;
	v616 = v607
	goto L181
L181:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+28))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v616)+20))
	F_emscripten_builtin_free(m, v618)
	mBase = m.M
	F_emscripten_builtin_free(m, v616)
	mBase = m.M
	if v617 != 0 {
		v616 = v617
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
	F_freeaddrinfo(m, v607)
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
	return base.B2i32(l0 != int32(0)) & base.B2i32(l0 != int32(4101656)) & base.B2i32(l0 != int32(4101680)) & base.B2i32(l0 != int32(4685060)) & base.B2i32(l0 != int32(4685084))
}
func F__lca(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v15 = v9 + int32(16)
		v16 = F_ArrayGetNItems(m, v13, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v18 < int32(2) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v22 = F_array_contains_nulls(m, v9)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v22 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(153041), int32(0))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496992), int32(308), int32(508485))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
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
						v26 = F_palloc(m, v16<<(uint(int32(2))%32))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							if v16 <= int32(0) {
							} else {
								if v21 != 0 {
									v36 = v21
								} else {
									v36 = (v18<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v37 = v9 + v36
								if v16&int32(1) == int32(0) {
									v56 = v37
									v57 = v16
								} else {
									v43 = v16 - int32(1)
									v44 = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v26+v43<<(uint(v44)%32)))) = v37
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
									v56 = v37 + (int32(base.Ui32(v48)>>(uint(v44)%32))+int32(3))&int32(2147483644)
									v57 = v43
								}
								if v16 == int32(1) {
								} else {
									v61 = v56
									v64 = v57
									for {
										v67 = int32(1)
										v68 = v64 - v67
										v69 = int32(2)
										*(*int32)(unsafe.Add(mBase, uint32(v26+v68<<(uint(v69)%32)))) = v61
										v74 = v64 - v69
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v81 = int32(3)
										v83 = int32(2147483644)
										v85 = v61 + (int32(base.Ui32(v78)>>(uint(v69)%32))+v81)&v83
										*(*int32)(unsafe.Add(mBase, uint32(v26+v74<<(uint(v69)%32)))) = v85
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
										if base.Ui32(v67) < base.Ui32(v68) {
											v61 = v85 + (int32(base.Ui32(v87)>>(uint(v69)%32))+v81)&v83
											v64 = v74
											continue
										} else {
											break
										}
										break
									}
								}
							}
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v105 = F_ArrayGetNItems(m, v104, v15)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = F_lca_inner(m, v26, v105)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v26)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v111 != v9 {
											F_pfree(m, v9)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												if v107 != 0 {
													v118 = v107
												} else {
													v115 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
													v118 = int32(0)
												}
												return v118
											}
										} else {
											if v107 != 0 {
												v118 = v107
											} else {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v118 = int32(0)
											}
											return v118
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
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(313792), int32(0))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496992), int32(304), int32(508485))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
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
	v6 = F_DirectFunctionCall2Coll(m, int32(5584), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
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
		v22 = F_local2local(m, v6, v5, v10, int32(9), int32(29), int32(2239360), base.B2i32(v7 != int32(0)))
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
				v32 = int32(4)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v34&int32(254) == int32(2) {
					v43 = v32
				} else {
					v43 = base.B2i32(v34 == int32(18)) << (uint(v32) % 32)
				}
				if v34 == int32(1) {
					v46 = v32
				} else {
					v46 = v43
				}
				v57 = v46
			} else {
				v47 = int32(1)
				if v28 != 0 {
					v57 = int32(base.Ui32(v26)>>(uint(v47)%32)) - v47
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v58 = int32(1)
			v59 = v18 + v58
			if v20&v58 != 0 {
				v64 = v59
			} else {
				v64 = v18 + int32(4)
			}
			if v20 == int32(1) {
				v67 = int32(4)
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
				if v69&int32(254) == int32(2) {
					v78 = v67
				} else {
					v78 = base.B2i32(v69 == int32(18)) << (uint(v67) % 32)
				}
				if v69 == int32(1) {
					v81 = v67
				} else {
					v81 = v78
				}
				v94 = v81
			} else {
				v82 = int32(1)
				if v20&v82 != 0 {
					v94 = int32(base.Ui32(v20)>>(uint(v82)%32)) - v82
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v94 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v95 = F_varstr_levenshtein(m, v29, v57, v64, v94, v23, v22, v21)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				return v95
			}
		}
	}
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
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
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
	v232 = m.ExcPending
	if v232 != 0 {
		goto L27
	} else {
		goto L74
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L27
	} else {
		goto L71
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L27
	} else {
		goto L68
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L27
	} else {
		goto L65
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v33 < v36 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v81 <= int32(0) {
		goto L1
	} else {
		goto L24
	}
L10:
	;
	v81 = v43 + int32(1)
	goto L9
L11:
	;
	v43 = v33
	v44 = v36
	goto L14
L12:
	;
	goto L13
L13:
	;
	v69 = F_SystemAttributeByName(m, v32)
	mBase = m.M
	if v69 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v46 = int32(4)
	v51 = v31 + int32(20) + v44<<(uint(v46)%32) + v43*int32(100)
	v54 = F_namestrcmp(m, v51+v46, v32)
	mBase = m.M
	if v54 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+91)))
	if v57 != int32(1) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v61 = v43 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v61 < v62 {
		v43 = v61
		v44 = v62
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
	v81 = int32(-9)
	goto L9
L22:
	;
	goto L23
L23:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+74)))
	v81 = v73
	goto L9
L24:
	;
	v85 = v19 & int32(3)
	if v29 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v85 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v91 = F_bms_is_member(m, v81+int32(7), v90)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	if v91 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v97 = F_SPI_getvalue(m, v28, v31, v81)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v99 = F_SPI_getvalue(m, v29, v31, v81)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v97 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v97 == int32(0) {
		goto L25
	} else {
		goto L53
	}
L33:
	;
	F_pfree(m, v99)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L27
	} else {
		goto L52
	}
L34:
	;
	if v99 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if v99 == int32(0) {
		goto L32
	} else {
		goto L51
	}
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v104 == int32(0) {
		v123 = v103
		v124 = v104
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v129 = int32(0)
	v133 = F_strtox_2(m, v97, v129, int32(10), int64(4294967295))
	mBase = m.M
	goto L49
L40:
	;
	if v124-v123 == int32(0) {
		goto L33
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	if v103 != v104 {
		v123 = v103
		v124 = v104
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v108 = v97
	v109 = v99
	goto L44
L44:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v113 == int32(0) {
		v123 = v112
		v124 = v113
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v123 = v112
	v124 = v113
	goto L41
L46:
	;
	v116 = int32(1)
	if v112 == v113 {
		v108 = v108 + v116
		v109 = v109 + v116
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	goto L39
L49:
	;
	v135 = F_DirectFunctionCall1Coll(m, int32(2368), v129, base.I32_wrap_i64(v133))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	goto L36
L51:
	;
	goto L33
L52:
	;
	goto L32
L53:
	;
	F_pfree(m, v97)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L27
	} else {
		goto L54
	}
L54:
	;
	goto L25
L55:
	;
	m.G0 = v11 + int32(48)
	if v85 == int32(2) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v149 = F_SPI_getvalue(m, v28, v31, v81)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L27
	} else {
		goto L57
	}
L57:
	;
	if v149 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v154 = int32(0)
	v158 = F_strtox_2(m, v149, v154, int32(10), int64(4294967295))
	mBase = m.M
	goto L59
L59:
	;
	v160 = F_DirectFunctionCall1Coll(m, int32(2368), v154, base.I32_wrap_i64(v158))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L27
	} else {
		goto L60
	}
L60:
	;
	F_pfree(m, v149)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L27
	} else {
		goto L61
	}
L61:
	;
	goto L55
L62:
	;
	v170 = v29
	goto L64
L63:
	;
	v170 = v28
	goto L64
L64:
	;
	return v170
L65:
	;
	F_errmsg_internal(m, int32(225694), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L27
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(497248), int32(39), int32(409407))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L27
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v194
	F_errmsg_internal(m, int32(30472), v11)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(497248), int32(43), int32(409407))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v213
	F_errmsg_internal(m, int32(250906), v11+int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(497248), int32(55), int32(409407))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L27
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v234
	F_errmsg_internal(m, int32(72210), v11+int32(32))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L27
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(497248), int32(71), int32(409407))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L27
	} else {
		goto L76
	}
L76:
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
	v11 = int32(572057)
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
	F_errmsg(m, int32(717998), v9)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(496310), int32(134), int32(251798))
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
									F_errmsg_internal(m, int32(39969), v11+int32(16))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_errfinish(m, int32(500408), int32(1040), int32(242476))
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
								v41 = *(*int32)(unsafe.Add(mBase, _consts[508]))
								F_fmgr_info_cxt(m, v34, l0+int32(212), v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									if v24 != 0 {
										v47 = *(*int32)(unsafe.Add(mBase, _consts[508]))
										F_fmgr_info_cxt(m, v24, l0+int32(240), v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											if v23 != 0 {
												v53 = *(*int32)(unsafe.Add(mBase, _consts[508]))
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
											v53 = *(*int32)(unsafe.Add(mBase, _consts[508]))
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
				F_errmsg_internal(m, int32(50788), v11)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(500408), int32(1020), int32(242476))
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
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1044), v6+int32(8), int32(0))
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
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(902), v6+int32(8), int32(0))
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
				F_errmsg_internal(m, int32(488233), v7)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(498841), int32(4554), int32(100876))
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
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_consts[857])))
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
			F_errmsg_internal(m, int32(58994), v12)
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return
			} else {
				F_errfinish(m, int32(500693), int32(597), int32(414750))
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
		v6 = int32(542761)
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
			F_errmsg_internal(m, int32(52861), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(500490), int32(92), int32(243213))
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
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
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
				v73 = v21
			} else {
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)))
				if v27 != 0 {
					v28 = F__emscripten_memcpy_bulkmem(m, v21, v9+int32(10), v27)
					mBase = m.M
					v29 = v28
				} else {
					v29 = v21
				}
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)))
				v31 = v29 + v30
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
				if base.Ui32(v32) < base.Ui32(int32(2)) {
					v73 = v31
				} else {
					v45 = v9 + int32(8) + (v30+int32(9))&int32(131064)
					v46 = v31
					v48 = int32(1)
					for {
						v50 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v50)
						v53 = v46 + int32(1)
						v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
						if v56 != 0 {
							v57 = F__emscripten_memcpy_bulkmem(m, v53, v45+int32(2), v56)
							mBase = m.M
							v58 = v57
						} else {
							v58 = v53
						}
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
						v60 = v58 + v59
						v67 = v48 + int32(1)
						v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
						if base.Ui32(v67) < base.Ui32(v68) {
							v45 = v45 + (v59+int32(9))&int32(131064)
							v46 = v60
							v48 = v67
							continue
						} else {
							break
						}
						break
					}
					v73 = v60
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = (v73 - v18) << (uint(int32(2)) % 32)
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v81 != v9 {
				F_pfree(m, v9)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
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
		v50 = F_dotrim(m, v17, v45, int32(747780), v47, v47, int32(0))
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
	F_errmsg(m, int32(388000), v11+int32(4080))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(500012), int32(267), int32(317935))
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
