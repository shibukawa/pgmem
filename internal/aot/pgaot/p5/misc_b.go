package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BuildCallback_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
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
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 float32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 float64
	_ = v616
	var v618 float64
	_ = v618
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v30 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v34 = int32(_a_F_BuildCallback_1_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[0]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l5)+184))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[0])) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v44 = l5 + int32(48)
	v45 = F_HnswFormIndexValue(m, v28+int32(8), l2, l3, v42, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v28 + int32(16)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[0])) = v35
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l5)+184))
	F_MemoryContextReset(m, v695)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L5
	} else {
		goto L171
	}
L5:
	;
	return
L6:
	;
	if v45 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v50 = v33 + int32(76)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v52 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v79 = F_LWLockAcquire(m, v50, int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L19
	}
L9:
	;
	v56 = int32(18)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v58 == v56 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v69 = int32(1)
	if v52&v69 != 0 {
		v77 = int32(base.Ui32(v52) >> (uint(v69) % 32))
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v61 = v56
	goto L14
L13:
	;
	v61 = int32(2)
	goto L14
L14:
	;
	if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = int32(6)
	goto L17
L16:
	;
	v68 = v61
	goto L17
L17:
	;
	v77 = v68
	goto L8
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v77 = int32(base.Ui32(v73) >> (uint(int32(2)) % 32))
	goto L8
L19:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+92)))
	if v81 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v612 = base.AtomicRmwXchg32(m, v33, int32(0), int32(1))
	if v612 != 0 {
		goto L163
	} else {
		goto L164
	}
L21:
	;
	F_LWLockRelease(m, v50)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v91 = v33 + int32(52)
	v93 = F_LWLockAcquire(m, v91, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L27
	}
L24:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v88 = F_HnswInsertTupleOnDisk(m, l0, v44, v86, l1, int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	if v88 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L4
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	if v39 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(_a_F_BuildCallback_1_1)
	goto L30
L29:
	;
	v98 = int32(0)
	goto L30
L30:
	;
	v99 = F_add_size(m, v95, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	if base.Ui32(v101) <= base.Ui32(v99) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_LWLockRelease(m, v91)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(l5)+168))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l5)+176))
	v148 = l5 + int32(188)
	v149 = F_HnswInitElement(m, v39, l1, v144, v145, v146, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L53
	}
L35:
	;
	F_LWLockRelease(m, v50)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v108 = F_LWLockAcquire(m, v50, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+92)))
	if v110 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v115 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_LWLockRelease(m, v50)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L50
	}
L41:
	;
	if v115 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = base.I64_trunc_sat_f64_s(v117)
	F_errmsg(m, int32(_a_F_BuildCallback_1_2), v28)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_FlushPages(m, l5)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L49
	}
L45:
	;
	v125 = F_errdetail(m, int32(_a_F_BuildCallback_1_3), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errhint(m, int32(_a_F_BuildCallback_1_4), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_BuildCallback_1_5), int32(542), int32(_a_F_BuildCallback_1_6))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	goto L40
L50:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v142 = F_HnswInsertTupleOnDisk(m, l0, v44, v140, l1, int32(1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v142 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	goto L4
L53:
	;
	v151 = F_HnswAlloc(m, v148, v77)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_LWLockRelease(m, v91)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	if v77 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	base.MemoryCopy(m, v151, v155, v77)
	goto L58
L57:
	;
	goto L58
L58:
	;
	if v151 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v161 = v151 - v39 + int32(1)
	goto L61
L60:
	;
	v161 = int32(0)
	goto L61
L61:
	;
	if v39 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v162 = v161
	goto L64
L63:
	;
	v162 = v151
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+88)) = v162
	v165 = v149 + int32(92)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[1]))
	F_LWLockInitialize(m, v165, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v175 = v173 + int32(32)
	v177 = F_LWLockAcquire(m, v175, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_LWLockRelease(m, v175)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	v182 = v173 + int32(16)
	v184 = F_LWLockAcquire(m, v182, int32(1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	if v170 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v222 = int32(0)
	F_HnswFindElementNeighbors(m, v170, v149, v221, v222, v44, v171, v172, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L86
	}
L70:
	;
	F_LWLockRelease(m, v182)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L78
	}
L71:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+65)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+65)))
	if base.Ui32(v196) <= base.Ui32(v197) {
		v221 = v195
		goto L69
	} else {
		goto L77
	}
L72:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	if v186 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	if v192 == int32(0) {
		goto L70
	} else {
		goto L76
	}
L75:
	;
	v195 = v170 + v186 - int32(1)
	goto L71
L76:
	;
	v195 = v192
	goto L71
L77:
	;
	goto L70
L78:
	;
	v202 = int32(0)
	v204 = F_LWLockAcquire(m, v175, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v207 = F_LWLockAcquire(m, v182, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	F_LWLockRelease(m, v175)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	if v170 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	v221 = v213
	goto L69
L83:
	;
	goto L84
L84:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	if v214 == int32(0) {
		v221 = v202
		goto L69
	} else {
		goto L85
	}
L85:
	;
	v221 = v170 + v214 - int32(1)
	goto L69
L86:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	if v227 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v251 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L88:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v149)+72))
	v230 = int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227+v228-v230)))
	if v232 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v149)+88))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v149)+72))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v248 = v244
	v249 = v246
	goto L87
L91:
	;
	v237 = v227 + v232 - v230
	goto L93
L92:
	;
	v237 = int32(0)
	goto L93
L93:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v149)+88))
	if v238 == int32(0) {
		v248 = v222
		v249 = v237
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v248 = v238 + v227 - int32(1)
	v249 = v237
	goto L87
L95:
	;
	F_LWLockRelease(m, v182)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L5
	} else {
		goto L161
	}
L96:
	;
	v541 = v149 + int32(4)
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+64)))
	v545 = v543 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+64)) = uint8(v545)
	v549 = v298 + v543*int32(6)
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v541)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v549)+8)) = uint16(v550)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	*(*int32)(unsafe.Add(mBase, uint32(v549)+4)) = v552
	goto L159
L97:
	;
	v348 = base.AtomicRmwXchg32(m, v250, int32(0), int32(1))
	if v348 != 0 {
		goto L113
	} else {
		goto L114
	}
L98:
	;
	v260 = int32(0)
	goto L99
L99:
	;
	v285 = v249 + int32(8) + v260*int32(12)
	if v227 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L97
L101:
	;
	v303 = F_datumIsEqual(m, base.I64_extend_i32_u(v248), base.I64_extend_i32_u(v299), int32(0), int32(-1))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L5
	} else {
		goto L107
	}
L102:
	;
	v298 = v289
	v299 = v290 + v227 - int32(1)
	goto L101
L103:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v287 = v227 + v286
	v289 = v287 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+87))
	if v290 != 0 {
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+88))
	v298 = v292
	v299 = v293
	goto L101
L106:
	;
	v298 = v289
	v299 = int32(0)
	goto L101
L107:
	;
	if v303 == int32(0) {
		goto L97
	} else {
		goto L108
	}
L108:
	;
	v308 = v298 + int32(92)
	v310 = F_LWLockAcquire(m, v308, int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+64)))
	if v312 != int32(10) {
		goto L96
	} else {
		goto L110
	}
L110:
	;
	F_LWLockRelease(m, v308)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v318 = v260 + int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v318 < v319 {
		v260 = v318
		goto L99
	} else {
		goto L112
	}
L112:
	;
	goto L100
L113:
	;
	F_s_lock(m, v250, int32(_a_F_BuildCallback_1_7))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v352
	v356 = v149 - v227 + int32(1)
	if v227 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L115
L117:
	;
	v357 = v356
	goto L119
L118:
	;
	v357 = v149
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v357
	v359 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v250))), uint32(v359))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+65)))
	v364 = v362
	goto L120
L120:
	;
	v392 = v171 << (uint(base.B2i32(v364 == int32(0))) % 32)
	v393 = F_mul_size(m, int32(12), v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L122
	}
L121:
	;
	if v221 != 0 {
		goto L152
	} else {
		goto L153
	}
L122:
	;
	v395 = F_add_size(m, int32(8), v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v397 = F_palloc(m, v395)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v400 = F_LWLockAcquire(m, v165, int32(1))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	if v227 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	if v395 != 0 {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v420 = v409 + v227 - int32(1)
	goto L126
L128:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v149)+72))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v227+v402+v364<<(uint(int32(2))%32)-int32(1))))
	if v409 != 0 {
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v149)+72))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411+v364<<(uint(int32(2))%32))))
	v420 = v415
	goto L126
L131:
	;
	v420 = int32(0)
	goto L126
L132:
	;
	base.MemoryCopy(m, v397, v420, v395)
	goto L134
L133:
	;
	goto L134
L134:
	;
	F_LWLockRelease(m, v165)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if int32(0) < v424 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v432 = int32(0)
	goto L139
L137:
	;
	goto L138
L138:
	;
	if int32(0) < v364 {
		v364 = v364 - int32(1)
		goto L120
	} else {
		goto L151
	}
L139:
	;
	v457 = v397 + int32(8) + v432*int32(12)
	if v227 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L138
L141:
	;
	v493 = *(*float32)(unsafe.Add(mBase, uint32(v457)+4))
	v494 = int32(0)
	F_HnswUpdateConnection(m, v227, v490, v149, v493, v392, v494, v494, v44)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L5
	} else {
		goto L148
	}
L142:
	;
	v458 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v460 = v227 + v459
	v462 = v460 + int32(91)
	v464 = F_LWLockAcquire(m, v462, v458)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v481 = v479 + int32(92)
	v483 = F_LWLockAcquire(m, v481, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L147
	}
L145:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460)+71))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v227+v466+v364<<(uint(int32(2))%32)-int32(1))))
	if v473 == int32(0) {
		v490 = v458
		v491 = v462
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v490 = v227 + v473 - int32(1)
	v491 = v462
	goto L141
L147:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+72))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v485+v364<<(uint(int32(2))%32))))
	v490 = v489
	v491 = v481
	goto L141
L148:
	;
	F_LWLockRelease(m, v491)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L149
	}
L149:
	;
	v501 = v432 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if v501 < v502 {
		v432 = v501
		goto L139
	} else {
		goto L150
	}
L150:
	;
	goto L140
L151:
	;
	goto L121
L152:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+65)))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+65)))
	if base.Ui32(v533) <= base.Ui32(v534) {
		goto L95
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v227 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+48)) = v149
	goto L95
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+48)) = v356
	goto L95
L159:
	;
	F_LWLockRelease(m, v308)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	goto L95
L161:
	;
	F_LWLockRelease(m, v50)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	goto L20
L163:
	;
	F_s_lock(m, v33, int32(_a_F_BuildCallback_1_7))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L5
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v616 = *(*float64)(unsafe.Add(mBase, uint32(v33)+8))
	v618 = base.F64_add(v616, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v33)+8)) = v618
	v624 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[2]))
	if v624 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v665 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v33))), uint32(v665))
	goto L4
L168:
	;
	goto L167
L169:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildCallback_1[3])))
	if v628&int32(1) == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v633 = int32(_a_F_BuildCallback_1_8)
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4]))
	v636 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4])) = v635 + v636
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = v639 + v636
	v643 = int32(0)
	v645 = int32(_a_F_BuildCallback_1_9)
	v646 = base.AtomicRmwOr32(m, v643, v645, v643)
	*(*int64)(unsafe.Add(mBase, uint32(v624+int32(96))+232)) = base.I64_trunc_sat_f64_s(v618)
	v654 = base.AtomicRmwOr32(m, v643, v645, v643)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = v655 + v636
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4])) = v661 - v636
	goto L168
L171:
	;
	goto L3
}
func F_BuildSpeculativeIndexInfo(m *base.Module, l0 int32, l1 int32) {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+10)))
	v16 = F_palloc_mul(m, int32(4), v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v16
	v20 = F_palloc_mul(m, int32(4), v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v20
	v24 = F_palloc_mul(m, int32(2), v15)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v24
	if int32(0) < v15 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v32 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v11 + int32(16)
	return
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
	v41 = v32 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
	v46 = F_IndexAmTranslateCompareType(m, int32(3), v39, v44, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v49 = v32 << (uint(int32(1)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v50))) = uint16(v46)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53+v41)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56+v41)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59+v49))))
	v62 = F_get_opfamily_member(m, v55, v58, v58, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v41))) = v62
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67+v41)))
	if v69 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v72 = F_get_opcode(m, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v41))) = v72
	v78 = v32 + int32(1)
	if v78 != v15 {
		v32 = v78
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+v32<<(uint(int32(1))%32)))))
	v101 = v32 << (uint(int32(2)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101+v102)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105+v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v99
	F_errmsg_internal(m, int32(_a_F_BuildSpeculativeIndexInfo_0), v11)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_BuildSpeculativeIndexInfo_1), int32(2750), int32(_a_F_BuildSpeculativeIndexInfo_2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_begin_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_begin_prepare_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(1058)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
	v18 = int32(_a_F_begin_prepare_cb_wrapper_1)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_begin_prepare_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_begin_prepare_cb_wrapper[0])) = v8 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v8 + int32(16)
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+147)) = uint8(v28)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+164)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+152)) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v36 == v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_begin_prepare_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_begin_prepare_cb_wrapper_3), v8)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_begin_prepare_cb_wrapper_4), int32(1021), int32(_a_F_begin_prepare_cb_wrapper_5))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v36].(func(*base.Module, int32, int32))(m, v10, l1)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_begin_prepare_cb_wrapper[0])) = v59
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_beginmerge(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v13 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L29
	}
L2:
	;
	v16 = v13
	goto L4
L3:
	;
	v16 = v14
	goto L4
L4:
	;
	if int32(0) < v16 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v11 + int32(32)
	return
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v24<<(uint(int32(2))%32))))
	v33 = v11 + int32(8)
	v35 = F_LogicalTapeRead(m, v31, v33, int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	return
L11:
	;
	if v35 != int32(4) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v40].(func(*base.Module, int32, int32, int32, int32))(m, l0, v33, v31, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v113 = v24 + int32(1)
	if v113 != v16 {
		v24 = v113
		goto L8
	} else {
		goto L28
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v24
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_beginmerge[0]))
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v49 + int32(1)
	if v49 <= int32(0) {
		v88 = v49
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v97 = v44 + v88*int32(24)
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+16)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v102
	goto L15
L22:
	;
	v56 = v49
	goto L23
L23:
	;
	v65 = int32(1)
	v68 = int32(base.Ui32(v56-v65) >> (uint(v65) % 32))
	v71 = v44 + v68*int32(24)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = m.T0[v72].(func(*base.Module, int32, int32, int32) int32)(m, v11+int32(8), v71, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v88 = int32(0)
	goto L21
L25:
	;
	if int32(0) <= v73 {
		v88 = v56
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v79 = v44 + v56*int32(24)
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+16)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v84
	if v68 != 0 {
		v56 = v68
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L9
L29:
	;
	F_errmsg_internal(m, int32(_a_F_beginmerge_0), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_beginmerge_1), int32(3173), int32(_a_F_beginmerge_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bernoulli_samplescangetsamplesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 float32
	_ = v18
	var v21 int32
	_ = v21
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	v8 = float64(0.10000000149011612)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_estimate_expression_value(m, l0, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v13 != int32(7) {
			v37 = v8
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+32)))
			if v16 != 0 {
				v37 = v8
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
				v18 = base.F32_reinterpret_i32(v17)
				v21 = int32(0)
				if base.B2i32(base.F32_ge(v18, float32(0)) == v21)|base.B2i32(base.F32_le(v18, float32(100)) == v21)|base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v17&int32(2147483647))) != 0 {
					v37 = v8
				} else {
					v37 = base.F64_promote_f32(base.F32_div(v18, float32(100)))
				}
			}
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v39
		v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
		v42 = base.F64_mul(v37, v41)
		v43 = float64(1e+100)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v42)&int64(9223372036854775807)))|base.F64_gt(v42, v43) != 0 {
			v56 = v43
		} else {
			v52 = float64(1)
			if base.F64_le(v42, v52) != 0 {
				v56 = v52
			} else {
				v56 = base.F64_nearest(v42)
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v56
		return
	}
}
func F_bf_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = l2
	if l2 != 0 {
		base.MemoryCopy(m, v6+int32(4), l1, l2)
	} else {
	}
	v14 = v6 + int32(68)
	if l3 != 0 {
		if v8 == int32(0) {
			return int32(0)
		} else {
			base.MemoryCopy(m, v14, l3, v8)
			return int32(0)
		}
	} else {
		if v8 == int32(0) {
		} else {
			base.MemoryFill(m, v14, int32(0), v8)
		}
		return int32(0)
	}
}
func F_big5_to_euc_tw(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v17, v18, v19, int32(36), int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v198)
	m.G0 = v12 + int32(16)
	return base.I64_extend_i32_s(v189 - v16)
L4:
	;
	v189 = v16
	v190 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v16
	v29 = v15
	v30 = v19
	goto L7
L7:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if v37 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v189 = v186
	v190 = v181
	goto L3
L9:
	;
	if int32(0) < v182 {
		v28 = v186
		v29 = v181
		v30 = v182
		goto L7
	} else {
		goto L60
	}
L10:
	;
	v41 = F_pg_encoding_verifymbchar(m, int32(36), v28, v30)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v37 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L13:
	;
	if v41 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v14 != int64(0) {
		v189 = v28
		v190 = v29
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v55 = (v50 | v37<<(uint(int32(8))%32)) & int32(_a_F_big5_to_euc_tw_0)
	v57 = v12 + int32(15)
	if base.Ui32(v55) <= base.Ui32(int32(_a_F_big5_to_euc_tw_1)) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	switch v120 - int32(149) {
	case 0:
		goto L47
	case 1:
		goto L49
	default:
		goto L48
	}
L20:
	;
	v119 = v117 & int32(_a_F_big5_to_euc_tw_0)
	goto L19
L21:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v112)
	v117 = int32(63)
	goto L20
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v106)
	v117 = v105 | int32(-32640)
	goto L20
L23:
	;
	v100 = int32(246)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v100)
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
	v117 = v102 | int32(-32640)
	goto L20
L24:
	;
	v99 = int32(_a_F_big5_to_euc_tw_2)
	goto L23
L25:
	;
	v99 = int32(_a_F_big5_to_euc_tw_3)
	goto L23
L26:
	;
	v99 = int32(_a_F_big5_to_euc_tw_4)
	goto L23
L27:
	;
	v99 = int32(_a_F_big5_to_euc_tw_5)
	goto L23
L28:
	;
	v99 = int32(_a_F_big5_to_euc_tw_6)
	goto L23
L29:
	;
	v99 = int32(_a_F_big5_to_euc_tw_7)
	goto L23
L30:
	;
	v89 = F_BinarySearchRange(m, int32(_a_F_big5_to_euc_tw_8), int32(46), v55)
	mBase = m.M
	if v89 == int32(0) {
		goto L21
	} else {
		goto L45
	}
L31:
	;
	v105 = v84
	v106 = int32(149)
	goto L22
L32:
	;
	switch v55 - int32(_a_F_big5_to_euc_tw_9) {
	case 0:
		v72 = int32(_a_F_big5_to_euc_tw_10)
		goto L35
	case 1, 3:
		goto L39
	case 2:
		goto L38
	case 4:
		goto L37
	default:
		goto L40
	}
L33:
	;
	goto L34
L34:
	;
	switch v55 - int32(_a_F_big5_to_euc_tw_11) {
	case 0:
		v99 = int32(_a_F_big5_to_euc_tw_12)
		goto L23
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	case 5:
		goto L25
	case 6:
		goto L24
	default:
		goto L43
	}
L35:
	;
	v73 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v73)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+2)))
	v117 = v75 | int32(-32640)
	goto L20
L36:
	;
	v72 = int32(_a_F_big5_to_euc_tw_13)
	goto L35
L37:
	;
	v72 = int32(_a_F_big5_to_euc_tw_14)
	goto L35
L38:
	;
	v72 = int32(_a_F_big5_to_euc_tw_15)
	goto L35
L39:
	;
	v68 = F_BinarySearchRange(m, int32(_a_F_big5_to_euc_tw_16), int32(23), v55)
	mBase = m.M
	if v68 != 0 {
		v84 = v68
		goto L31
	} else {
		goto L42
	}
L40:
	;
	if v55 == int32(_a_F_big5_to_euc_tw_17) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L21
L43:
	;
	if v55 != int32(_a_F_big5_to_euc_tw_18) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v84 = int32(_a_F_big5_to_euc_tw_19)
	goto L31
L45:
	;
	v105 = v89
	v106 = int32(150)
	goto L22
L46:
	;
	v181 = v164
	v182 = v30 - v41
	v186 = v28 + v41
	goto L9
L47:
	;
	v155 = int32(8)
	v159 = v119<<(uint(v155)%32) | int32(base.Ui32(v119)>>(uint(v155)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v29))) = uint16(v159)
	v164 = v29 + int32(2)
	goto L46
L48:
	;
	if base.Ui32((v120+int32(10))&int32(255)) <= base.Ui32(int32(4)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)) = uint8(v119)
	v124 = int32(_a_F_big5_to_euc_tw_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v29))) = uint16(v124)
	v127 = int32(base.Ui32(v119) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)) = uint8(v127)
	v164 = v29 + int32(4)
	goto L46
L50:
	;
	v137 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)) = uint8(v119)
	v142 = int32(base.Ui32(v119) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)) = uint8(v142)
	v145 = v139 - int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)) = uint8(v145)
	v164 = v29 + int32(4)
	goto L46
L51:
	;
	goto L52
L52:
	;
	if v14 != int64(0) {
		v189 = v28
		v190 = v29
		goto L3
	} else {
		goto L53
	}
L53:
	;
	F_report_untranslatable_char(m, int32(36), int32(4), v28, v30)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	if v14 != int64(0) {
		v189 = v28
		v190 = v29
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v37)
	v175 = int32(1)
	v181 = v29 + v175
	v182 = v30 - v175
	v186 = v28 + v175
	goto L9
L58:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	goto L8
}
func F_binaryheap_replace_first(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(2) <= v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = l0 + int32(24)
	v17 = v10
	v20 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v23 = int32(1)
	v24 = v20 << (uint(v23) % 32)
	v26 = v24 | v23
	v28 = v24 + int32(2)
	if v28 < v17 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14+v20<<(uint(int32(3))%32)))) = l1
	goto L3
L6:
	;
	goto L5
L7:
	;
	v30 = int32(3)
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v14+v26<<(uint(v30)%32))))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v14+v28<<(uint(v30)%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = m.T0[v39].(func(*base.Module, int64, int64, int32) int32)(m, v33, v37, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v46 = v26
	v47 = v17
	goto L9
L9:
	;
	if v47 <= v26 {
		goto L6
	} else {
		goto L15
	}
L10:
	;
	return
L11:
	;
	if v40 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v28
	goto L14
L13:
	;
	v44 = v26
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = v44
	v47 = v45
	goto L9
L15:
	;
	v51 = v14 + v46<<(uint(int32(3))%32)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v55 = m.T0[v54].(func(*base.Module, int64, int64, int32) int32)(m, l1, v52, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if int32(0) <= v55 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+v20<<(uint(int32(3))%32)))) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = v64
	v20 = v46
	goto L4
}
func F_binaryheap_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
}
func F_bind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_bind(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_bind[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_bitcat(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = F_bit_catenate(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v10)
			}
		}
	}
}
func F_biteq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v15 == v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = int32(8)
	v19 = v8 + v18
	v21 = v13 + v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v23 = int32(2)
	v24 = int32(base.Ui32(v22) >> (uint(v23) % 32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v27 = int32(base.Ui32(v25) >> (uint(v23) % 32))
	if base.Ui32(v24) < base.Ui32(v27) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v99 = int64(0)
	goto L6
L6:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v100 != v8 {
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v29 = v24
	goto L9
L8:
	;
	v29 = v27
	goto L9
L9:
	;
	v31 = v29 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v31) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v99 = base.I64_extend_i32_u(base.B2i32(v93 == int32(0)))
	goto L6
L11:
	;
	v93 = int32(0)
	goto L10
L12:
	;
	v67 = v62
	v68 = v63
	v69 = v64
	goto L22
L13:
	;
	if (v19|v21)&int32(3) != 0 {
		v62 = v19
		v63 = v21
		v64 = v31
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v55 = v19
	v56 = v21
	v57 = v31
	goto L15
L15:
	;
	if v57 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v39 = v19
	v40 = v21
	v41 = v31
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v44 != v45 {
		v62 = v39
		v63 = v40
		v64 = v41
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v55 = v50
	v56 = v48
	v57 = v52
	goto L15
L19:
	;
	v47 = int32(4)
	v48 = v40 + v47
	v50 = v39 + v47
	v52 = v41 - v47
	if base.Ui32(int32(3)) < base.Ui32(v52) {
		v39 = v50
		v40 = v48
		v41 = v52
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v62 = v55
	v63 = v56
	v64 = v57
	goto L12
L22:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 == v73 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v93 = v72 - v73
	goto L10
L24:
	;
	v75 = int32(1)
	v80 = v69 - v75
	if v80 != 0 {
		v67 = v67 + v75
		v68 = v68 + v75
		v69 = v80
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L11
L28:
	;
	F_pfree(m, v8)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v104 != v13 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	F_pfree(m, v13)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v99
L35:
	;
	goto L34
}
func F_bitge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int64(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v99))
L38:
	;
	goto L37
}
func F_bitoverlay(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v12 = F_bit_overlay(m, v3, v8, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
func F_bitoverlay_no_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v13 = F_bit_overlay(m, v4, v9, v11, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v13)
			}
		}
	}
}
func F_blbuildempty(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_BloomInitMetapage(m, l0, int32(3))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_blbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 float64
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(_a_F_blbulkdelete_0)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = F_palloc0(m, int32(40))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v30 = l1
	goto L3
L3:
	;
	F_initBloomState(m, v20+int32(16), v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v30 = v26
	goto L3
L6:
	;
	v36 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v43 = int32(0)
	v48 = F_read_stream_begin_relation(m, int32(13), v42, v22, v43, int32(3), v20+int32(8), v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = v5
	v64 = int32(1)
	goto L12
L10:
	;
	v290 = v5
	goto L11
L11:
	;
	F_read_stream_end(m, v48)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L57
	}
L12:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v290 = v270
	goto L11
L14:
	;
	v74 = F_read_stream_next_buffer(m, v48, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_LockBufferInternal(m, v74, int32(3))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v79 = F_GenericXLogStart(m, v22)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v279 = v64 + int32(1)
	if v279 != v36 {
		v62 = v270
		v64 = v279
		goto L12
	} else {
		goto L56
	}
L18:
	;
	v98 = v82 + int32(24)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1180))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
	v101 = int32(1)
	v107 = v99 * ((v100+v101)&int32(_a_F_blbulkdelete_1) - v101)
	if v107 != 0 {
		goto L32
	} else {
		goto L33
	}
L19:
	;
	v82 = F_GenericXLogRegisterBuffer(m, v79, v74, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+14)))
	if v84 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+16)))
	v86 = v82 + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+2)))
	if v87&int32(2) == int32(0) {
		goto L18
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_UnlockReleaseBuffer(m, v74)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	F_pfree(m, v79)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v270 = v62
	goto L17
L27:
	;
	F_UnlockReleaseBuffer(m, v74)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L55
	}
L28:
	;
	F_pfree(m, v79)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L54
	}
L29:
	;
	v219 = v207 - v82
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)) = uint16(v219)
	F_GenericXLogFinish(m, v79)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L53
	}
L30:
	;
	if v149 == v147 {
		v232 = v62
		goto L28
	} else {
		goto L52
	}
L31:
	;
	if base.B2i32(base.Ui32(int32(_a_F_blbulkdelete_2)-v167*v170) < base.Ui32(v167))|base.B2i32(base.Ui32(int32(2003)) < base.Ui32(v62)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L32:
	;
	v109 = v98
	v114 = v98
	goto L35
L33:
	;
	goto L34
L34:
	;
	if v100 == int32(0) {
		v232 = v62
		goto L28
	} else {
		goto L47
	}
L35:
	;
	v126 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v109, l3)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+16)))
	v152 = v82 + v151
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	if v153 == int32(0) {
		goto L30
	} else {
		goto L46
	}
L37:
	;
	v149 = v109 + v148
	if base.Ui32(v149) < base.Ui32(v107+v98) {
		v109 = v149
		v114 = v147
		goto L35
	} else {
		goto L45
	}
L38:
	;
	if v126 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+16)))
	v129 = v82 + v128
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129))))
	v132 = v130 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v129))) = uint16(v132)
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v30)+16)) = base.F64_add(v134, float64(1))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1180))
	v147 = v114
	v148 = v138
	goto L37
L40:
	;
	goto L41
L41:
	;
	if v109 == v114 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1180))
	v147 = v114 + v145
	v148 = v145
	goto L37
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1180))
	if v140 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	base.MemoryCopy(m, v114, v109, v140)
	goto L42
L45:
	;
	goto L36
L46:
	;
	v160 = base.B2i32(v149 == v147)
	v165 = v147
	v167 = v148
	v170 = v153
	goto L31
L47:
	;
	v160 = int32(1)
	v165 = v98
	v167 = v99
	v170 = v100
	goto L31
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(1184)+v62<<(uint(int32(2))%32)))) = v64
	v194 = v62 + int32(1)
	goto L50
L49:
	;
	v194 = v62
	goto L50
L50:
	;
	if v160 == int32(0) {
		v207 = v165
		v211 = v194
		goto L29
	} else {
		goto L51
	}
L51:
	;
	v232 = v194
	goto L28
L52:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+2)))
	v200 = v198 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+2)) = uint16(v200)
	v207 = v147
	v211 = v62
	goto L29
L53:
	;
	v251 = v211
	goto L27
L54:
	;
	v251 = v232
	goto L27
L55:
	;
	v270 = v251
	goto L17
L56:
	;
	goto L13
L57:
	;
	v301 = F_ReadBuffer(m, v22, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_LockBufferInternal(m, v301, int32(3))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v306 = F_GenericXLogStart(m, v22)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v309 = F_GenericXLogRegisterBuffer(m, v306, v301, int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v312 = v290 << (uint(int32(2)) % 32)
	if v312 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v309+int32(168), v20+int32(1184), v312)
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v309)+30)) = uint16(v290)
	v319 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v309)+28)) = uint16(v319)
	F_GenericXLogFinish(m, v306)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	F_UnlockReleaseBuffer(m, v301)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	m.G0 = v20 + int32(_a_F_blbulkdelete_0)
	return v30
}
func F_blcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 float64
	_ = v23
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v19 = v15 + int32(8)
	base.MemoryFill(m, v19, int32(0), int32(72))
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v23
	F_genericcostestimate(m, l0, l1, l2, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = v29
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v31
		v33 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v33
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v35
		v37 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v37
		m.G0 = v15 + int32(80)
		return
	}
}
func F_boolgt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3) & base.B2i32(v5 != v3))
}
func F_boot_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_boot_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_boot_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_1(m, int32(_a_F_boot_yy_create_buffer_2))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_boot_yy_create_buffer[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_boot_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_1(m, int32(_a_F_boot_yy_create_buffer_2))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_boot_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_1(m, int32(_a_F_boot_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_1(m, int32(_a_F_boot_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_bpcharin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_strlen(m, v3)
	mBase = m.M
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_bpchar_input(m, v3, v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v7)
	}
}
func F_bpchartruelen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v5 = int32(-1)
	v7 = l1 - int32(1)
	if v5 <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v5
	goto L3
L2:
	;
	v10 = v7
	goto L3
L3:
	;
	v14 = l1
	goto L4
L4:
	;
	v18 = v14 - int32(1)
	if v18 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v14
L6:
	;
	return v10 + int32(1)
L7:
	;
	goto L8
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v18))))
	if v23 == int32(32) {
		v14 = v18
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_bpchartypmodout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(64))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if int32(5) <= v8 {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8 - int32(4)
			v21 = F_pg_snprintf(m, v10, int32(64), int32(_a_F_bpchartypmodout_0), v6)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v10)
			}
		} else {
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v23)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v10)
		}
	}
}
func F_brinbuildCallbackParallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v17 = v13 | v14<<(uint(v10)%32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	if base.Ui32(v18) <= base.Ui32(v17) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
		if base.Ui32(v17) <= base.Ui32(v18+v20-int32(1)) {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
			v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
				v32 = F_brin_form_tuple(m, v29, v18, v25, v11+int32(12))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+72))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					F_tuplesort_putbrintuple(m, v34, v32, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v38, float64(1))
						F_pfree(m, v32)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
							v45 = v44
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
							v47 = base.I32_rem_u_s(v17, v46)
							*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
							F_brin_memtuple_initialize(m, v45, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
								v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v45 = v25
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
				v47 = base.I32_rem_u_s(v17, v46)
				*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
				F_brin_memtuple_initialize(m, v45, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
					v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
		if v26 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
			v32 = F_brin_form_tuple(m, v29, v18, v25, v11+int32(12))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+72))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				F_tuplesort_putbrintuple(m, v34, v32, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v38 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
					*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v38, float64(1))
					F_pfree(m, v32)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
						v45 = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
						v47 = base.I32_rem_u_s(v17, v46)
						*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
						F_brin_memtuple_initialize(m, v45, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
							v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v45 = v25
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
			v47 = base.I32_rem_u_s(v17, v46)
			*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
			F_brin_memtuple_initialize(m, v45, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
				v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			}
		}
	}
}
func F_brinrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l1 == int32(0) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 <= int32(0) {
		} else {
			v12 = v8 * int32(56)
			if v12 == int32(0) {
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				base.MemoryCopy(m, v15, l1, v12)
			}
		}
	}
	return
}
func F_brinvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v269 int64
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v411 int64
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int64
	_ = v456
	var v459 int32
	_ = v459
	var v461 int64
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v555 int64
	_ = v555
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	v15 = int64(0)
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	v24 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	if int32(0) < v249 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	return int32(0)
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	v33 = F_get_opfamily_name(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L46
	}
L7:
	;
	v37 = base.I64_extend_i32_u(v32)
	v38 = int64(0)
	v40 = F_SearchSysCacheList(m, int32(4), int32(1), v37, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(1)
	v45 = int64(0)
	v47 = F_SearchSysCacheList(m, int32(5), v42, v37, v45, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v49 <= int32(0) {
		v237 = v42
		v247 = v15
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(0)
	v59 = v42
	v69 = v15
	goto L11
L11:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v47-int32(-64)+v57<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+72))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
	v77 = v75 + v76
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
	switch v78 - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L17
	}
L12:
	;
	v237 = v213
	v247 = v214
	goto L1
L13:
	;
	v216 = v57 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v216 < v217 {
		v57 = v216
		v59 = v213
		v69 = v214
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v209 = int64(*(*int16)(unsafe.Add(mBase, uint32(v77)+16)))
	v213 = v207
	v214 = int64(1)<<(uint(v209)%64) | v69
	goto L13
L15:
	;
	v177 = int32(0)
	v180 = F_errstart(m, int32(17), v177)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L39
	}
L16:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v167 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v167
	v170 = int32(1)
	v175 = F_check_amproc_signature(m, v166, v167, v170, v170, v170, v20+int32(208))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L37
	}
L17:
	;
	if base.Ui32(int32(_a_F_brinvalidate_0)) < base.Ui32((v78-int32(16))&int32(_a_F_brinvalidate_1)) {
		v207 = v59
		goto L14
	} else {
		goto L30
	}
L18:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v127 = F_check_amoptsproc_signature(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L28
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+256)) = int64(9796820404457)
	v118 = int32(3)
	v122 = F_check_amproc_signature(m, v111, int32(16), int32(1), v118, v118, v20+int32(256))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L26
	}
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+248)) = int64(98784250089)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+240)) = int64(9796820404457)
	v107 = F_check_amproc_signature(m, v96, int32(16), int32(1), int32(3), int32(4), v20+int32(240))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L24
	}
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v82 = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+232)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v20)+224)) = v82
	v88 = int32(4)
	v92 = F_check_amproc_signature(m, v81, int32(16), int32(1), v88, v88, v20+int32(224))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v92 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v207 = v59
	goto L14
L24:
	;
	if v107 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v207 = v59
	goto L14
L26:
	;
	if v122 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v207 = v59
	goto L14
L28:
	;
	if v127 == int32(0) {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v207 = v59
	goto L14
L30:
	;
	v137 = int32(0)
	v140 = F_errstart(m, int32(17), v137)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v140 == int32(0) {
		v213 = v137
		v214 = v69
		goto L13
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v148 = F_format_procedure(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+188)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+180)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_3), v20+int32(176))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(114), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v213 = v137
	v214 = v69
	goto L13
L37:
	;
	if v175 != 0 {
		v207 = v59
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L15
L39:
	;
	if v180 == int32(0) {
		v207 = v177
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v188 = F_format_procedure(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+204)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+196)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_6), v20+int32(192))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(130), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v207 = v177
	goto L14
L45:
	;
	goto L12
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_brinvalidate_7), v20)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(58), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v258 = int32(0)
	v260 = v237
	v269 = v15
	goto L52
L50:
	;
	v402 = v237
	v411 = v15
	goto L51
L51:
	;
	v415 = v30 + int32(8)
	v416 = F_identify_opfamily_groups(m, v40, v47)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L97
	}
L52:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(-64)+v258<<(uint(int32(2))%32))))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+72))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+22)))
	v278 = v276 + v277
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+16)))
	if base.Ui32((v279+int32(-64))&int32(_a_F_brinvalidate_1)) <= base.Ui32(int32(_a_F_brinvalidate_8)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v402 = v392
	v411 = v324
	goto L51
L54:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+18)))
	if v325 == int32(115) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v286 = int32(0)
	v289 = F_errstart(m, int32(17), v286)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	if v315 != v316 {
		v323 = v260
		v324 = v269
		goto L54
	} else {
		goto L64
	}
L58:
	;
	if v289 == int32(0) {
		v323 = v286
		v324 = v269
		goto L54
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v297 = F_format_operator(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v299 = int32(*(*int16)(unsafe.Add(mBase, uint32(v278)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v20)+164)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_9), v20+int32(160))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(152), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v323 = v286
	v324 = v269
	goto L54
L64:
	;
	v323 = v260
	v324 = int64(1)<<(uint(base.I64_extend_i32_u(v279))%64) | v269
	goto L54
L65:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v363 = F_check_amop_signature(m, v359, int32(16), v361, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L77
	}
L66:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v278)+28))
	if v328 == int32(0) {
		v358 = v323
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v331 = int32(0)
	v334 = F_errstart(m, int32(17), v331)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v334 == int32(0) {
		v358 = v331
		goto L65
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v342 = F_format_operator(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+152)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v20)+148)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_10), v20+int32(144))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(180), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v358 = v331
	goto L65
L76:
	;
	v394 = v258 + int32(1)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	if v394 < v395 {
		v258 = v394
		v260 = v392
		v269 = v324
		goto L52
	} else {
		goto L85
	}
L77:
	;
	if v363 != 0 {
		v392 = v358
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v365 = int32(0)
	v368 = F_errstart(m, int32(17), v365)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v368 == int32(0) {
		v392 = v365
		goto L76
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v376 = F_format_operator(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_11), v20+int32(128))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(193), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v392 = v365
	goto L76
L85:
	;
	goto L53
L86:
	;
	F_ReleaseCatCacheList(m, v876)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L2
	} else {
		goto L187
	}
L87:
	;
	v841 = int32(0)
	v844 = F_errstart(m, int32(17), v841)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L2
	} else {
		goto L182
	}
L88:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	if v821&int32(16) != 0 {
		v867 = v805
		v871 = v809
		v874 = v812
		v876 = v814
		v878 = v816
		goto L86
	} else {
		goto L181
	}
L89:
	;
	v781 = int32(0)
	v784 = F_errstart(m, int32(17), v781)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L2
	} else {
		goto L173
	}
L90:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	if v759&int32(8) != 0 {
		v804 = v742
		v805 = v20
		v809 = v747
		v810 = v415
		v812 = v40
		v814 = v47
		v816 = v24
		goto L88
	} else {
		goto L172
	}
L91:
	;
	v713 = int32(0)
	v716 = F_errstart(m, int32(17), v713)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L2
	} else {
		goto L163
	}
L92:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	if v692&int32(4) != 0 {
		v742 = v675
		v747 = v680
		goto L90
	} else {
		goto L162
	}
L93:
	;
	v648 = int32(0)
	v651 = F_errstart(m, int32(17), v648)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L2
	} else {
		goto L153
	}
L94:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	if v626&int32(2) != 0 {
		v675 = v609
		v680 = v614
		goto L92
	} else {
		goto L152
	}
L95:
	;
	v609 = v539 + int32(16)
	v614 = v540
	goto L94
L96:
	;
	v577 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L2
	} else {
		goto L139
	}
L97:
	;
	if v416 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v560 = int32(1)
	v562 = int32(0)
	goto L96
L99:
	;
	goto L100
L100:
	;
	v423 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v423 < v424 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v431 = int32(0)
	v432 = v423
	v433 = v402
	goto L104
L102:
	;
	v539 = v423
	v540 = v402
	goto L103
L103:
	;
	if v539 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L104:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445+v431<<(uint(int32(2))%32))))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	if v31 == v450 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v539 = v455
	v540 = v529
	goto L103
L106:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v452 == v31 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v455 = v432
	goto L108
L108:
	;
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v449)+16))
	if v456 == int64(0) {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v454 = v449
	goto L111
L110:
	;
	v454 = v432
	goto L111
L111:
	;
	v455 = v454
	goto L108
L112:
	;
	v532 = v431 + int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v532 < v533 {
		v431 = v532
		v432 = v455
		v433 = v529
		goto L104
	} else {
		goto L134
	}
L113:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v450 != v459 {
		v529 = v433
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v449)+8))
	if v461 == v411 {
		v494 = v433
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L115
L117:
	;
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v449)+16))
	if v496 == v247 {
		v529 = v494
		goto L112
	} else {
		goto L126
	}
L118:
	;
	v463 = int32(0)
	v466 = F_errstart(m, int32(17), v463)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	if v466 == int32(0) {
		v494 = v463
		goto L117
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v474 = F_format_type_be(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	v477 = F_format_type_be(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+124)) = v477
	*(*int32)(unsafe.Add(mBase, uint32(v20)+120)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_12), v20+int32(112))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(232), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v494 = v463
	goto L117
L126:
	;
	v498 = int32(0)
	v501 = F_errstart(m, int32(17), v498)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L2
	} else {
		goto L127
	}
L127:
	;
	if v501 == int32(0) {
		v529 = v498
		goto L112
	} else {
		goto L128
	}
L128:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L2
	} else {
		goto L129
	}
L129:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v509 = F_format_type_be(m, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	v512 = F_format_type_be(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v33
	F_errmsg(m, int32(_a_F_brinvalidate_13), v20+int32(96))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(242), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L2
	} else {
		goto L133
	}
L133:
	;
	v529 = v498
	goto L112
L134:
	;
	goto L105
L135:
	;
	v560 = int32(1)
	v562 = int32(0)
	goto L96
L136:
	;
	goto L137
L137:
	;
	v555 = *(*int64)(unsafe.Add(mBase, uint32(v539)+8))
	if v555 == v411 {
		goto L95
	} else {
		goto L138
	}
L138:
	;
	v560 = int32(0)
	v562 = v539
	goto L96
L139:
	;
	if v577 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L2
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v602 = v562 + int32(16)
	if v560 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v415
	F_errmsg(m, int32(_a_F_brinvalidate_14), v20+int32(80))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(253), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v596 = v562 + int32(16)
	if v560 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v609 = v596
	v614 = int32(0)
	goto L94
L147:
	;
	goto L148
L148:
	;
	v630 = v596
	v647 = int32(1)
	goto L93
L149:
	;
	v609 = v602
	v614 = int32(0)
	goto L94
L150:
	;
	goto L151
L151:
	;
	v630 = v602
	v647 = int32(1)
	goto L93
L152:
	;
	v630 = v609
	v647 = int32(0)
	goto L93
L153:
	;
	if v651 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L2
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if v647 == int32(0) {
		v675 = v630
		v680 = v648
		goto L92
	} else {
		goto L161
	}
L157:
	;
	v656 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v415
	F_errmsg(m, int32(_a_F_brinvalidate_15), v20-int32(-64))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	if v647 != 0 {
		v696 = v630
		v699 = v656
		goto L91
	} else {
		goto L160
	}
L160:
	;
	v675 = v630
	v680 = v648
	goto L92
L161:
	;
	v696 = v630
	v699 = int32(1)
	goto L91
L162:
	;
	v696 = v675
	v699 = int32(0)
	goto L91
L163:
	;
	if v716 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L2
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if v699 == int32(0) {
		v742 = v696
		v747 = v713
		goto L90
	} else {
		goto L171
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v415
	F_errmsg(m, int32(_a_F_brinvalidate_15), v20+int32(48))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	if v699 == int32(0) {
		v742 = v696
		v747 = v713
		goto L90
	} else {
		goto L170
	}
L170:
	;
	v763 = v696
	v764 = v20
	v769 = v415
	v771 = v40
	v773 = v47
	v775 = v24
	v780 = int32(1)
	goto L89
L171:
	;
	v763 = v696
	v764 = v20
	v769 = v415
	v771 = v40
	v773 = v47
	v775 = v24
	v780 = int32(1)
	goto L89
L172:
	;
	v763 = v742
	v764 = v20
	v769 = v415
	v771 = v40
	v773 = v47
	v775 = v24
	v780 = int32(0)
	goto L89
L173:
	;
	if v784 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L2
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	if v780 != 0 {
		v825 = v764
		v830 = v769
		v832 = v771
		v834 = v773
		v836 = v775
		goto L87
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v764)+40)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v764)+36)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v764)+32)) = v769
	F_errmsg(m, int32(_a_F_brinvalidate_15), v764+int32(32))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	goto L176
L180:
	;
	v804 = v763
	v805 = v764
	v809 = v781
	v810 = v769
	v812 = v771
	v814 = v773
	v816 = v775
	goto L88
L181:
	;
	v825 = v805
	v830 = v810
	v832 = v812
	v834 = v814
	v836 = v816
	goto L87
L182:
	;
	if v844 == int32(0) {
		v867 = v825
		v871 = v841
		v874 = v832
		v876 = v834
		v878 = v836
		goto L86
	} else {
		goto L183
	}
L183:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L2
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v825)+24)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v825)+20)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v825)+16)) = v830
	F_errmsg(m, int32(_a_F_brinvalidate_15), v825+int32(16))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L2
	} else {
		goto L186
	}
L186:
	;
	v867 = v825
	v871 = v841
	v874 = v832
	v876 = v834
	v878 = v836
	goto L86
L187:
	;
	F_ReleaseCatCacheList(m, v874)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	F_ReleaseCatCache(m, v878)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L2
	} else {
		goto L189
	}
L189:
	;
	m.G0 = v867 + int32(272)
	return v871
}
func F_bsearch_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v14 = l1
	v15 = l2
	goto L4
L4:
	;
	v24 = v14 + int32(base.Ui32(v15)>>(uint(int32(1))%32))*l3
	v25 = m.T0[l4].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, l5)
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	return int32(0)
L7:
	;
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v24
L9:
	;
	goto L10
L10:
	;
	v34 = base.B2i32(int32(0) < v25)
	if int32(0) < v25 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = l3 + v24
	goto L13
L12:
	;
	v35 = v14
	goto L13
L13:
	;
	v38 = int32(base.Ui32(v15-v34) >> (uint(int32(1)) % 32))
	if v38 != 0 {
		v14 = v35
		v15 = v38
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_btcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 float32
	_ = v140
	var v142 float32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v174 int32
	_ = v174
	var v175 float64
	_ = v175
	var v178 float64
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 float64
	_ = v185
	var v188 int32
	_ = v188
	var v195 float64
	_ = v195
	var v198 float64
	_ = v198
	var v199 int32
	_ = v199
	var v204 float64
	_ = v204
	var v205 int32
	_ = v205
	var v211 float64
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 float32
	_ = v231
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 float64
	_ = v236
	var v239 int32
	_ = v239
	var v242 float64
	_ = v242
	var v244 int32
	_ = v244
	var v246 float64
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 float64
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 float64
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
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
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 float64
	_ = v385
	var v386 int32
	_ = v386
	var v390 float64
	_ = v390
	var v391 float64
	_ = v391
	var v394 float64
	_ = v394
	var v422 float64
	_ = v422
	var v425 float64
	_ = v425
	var v426 float64
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v491 float64
	_ = v491
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 float64
	_ = v518
	var v519 float64
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v549 float64
	_ = v549
	var v550 float64
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 float64
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 float64
	_ = v602
	var v603 int32
	_ = v603
	var v607 float64
	_ = v607
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 float64
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 float64
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 float64
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 float64
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 float64
	_ = v731
	var v733 float64
	_ = v733
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v823 int32
	_ = v823
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 float64
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 float64
	_ = v877
	var v879 int32
	_ = v879
	var v883 float64
	_ = v883
	var v885 float64
	_ = v885
	var v886 float64
	_ = v886
	var v889 float64
	_ = v889
	var v916 float64
	_ = v916
	var v920 float64
	_ = v920
	var v928 int32
	_ = v928
	var v930 float64
	_ = v930
	var v931 float64
	_ = v931
	var v932 float64
	_ = v932
	var v933 float64
	_ = v933
	var v938 float64
	_ = v938
	var v939 float64
	_ = v939
	var v943 float64
	_ = v943
	var v946 float64
	_ = v946
	var v948 float64
	_ = v948
	var v950 float64
	_ = v950
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 float64
	_ = v960
	var v961 float64
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 float32
	_ = v981
	var v982 float64
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 float64
	_ = v986
	var v989 int32
	_ = v989
	var v992 float64
	_ = v992
	var v994 int32
	_ = v994
	var v996 float64
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 float64
	_ = v1005
	var v1011 float64
	_ = v1011
	var v1017 float64
	_ = v1017
	var v1020 float64
	_ = v1020
	v9 = int32(0)
	v25 = float64(0)
	v29 = m.G0
	v31 = v29 - int32(176)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	base.MemoryFill(m, v31-int32(-64), v9, int32(72))
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+56)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = v39
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v47 == v9 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717)+101)))
	if v735 == int32(1) {
		goto L144
	} else {
		goto L145
	}
L2:
	;
	v56 = l0
	v57 = l1
	v58 = l2
	v59 = l3
	v60 = l4
	v61 = l5
	v62 = l6
	v63 = l7
	v64 = v31
	v66 = v33
	v67 = v9
	v68 = v9
	v69 = v9
	v71 = v9
	v73 = v9
	v74 = v47
	v75 = v9
	v76 = v9
	v77 = v9
	v79 = v9
	v80 = v51
	v82 = v25
	goto L8
L3:
	;
	v707 = l0
	v708 = l1
	v709 = l2
	v710 = l3
	v711 = l4
	v712 = l5
	v713 = l6
	v714 = l7
	v715 = v31
	v717 = v33
	v718 = v9
	v720 = v9
	v722 = v9
	v724 = v9
	v728 = v9
	v730 = v9
	v731 = v55
	v733 = v25
	goto L1
L4:
	;
	v55 = float64(1)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v51 = float64(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if int32(0) < v52 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v55 = v51
	goto L3
L8:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v75<<(uint(int32(2))%32))))
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+14)))
	if v89 <= v69 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v707 = v56
	v708 = v57
	v709 = v58
	v710 = v59
	v711 = v60
	v712 = v61
	v713 = v62
	v714 = v63
	v715 = v64
	v717 = v66
	v718 = v686
	v720 = v537
	v722 = v690
	v724 = v692
	v728 = v545
	v730 = v698
	v731 = v699
	v733 = v550
	goto L1
L10:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v552 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L11:
	;
	v535 = v67
	v536 = v68
	v537 = v69
	v539 = v71
	v545 = v77
	v549 = v80
	v550 = v82
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v76 != 0 {
		v707 = v56
		v708 = v57
		v709 = v58
		v710 = v59
		v711 = v60
		v712 = v61
		v713 = v62
		v714 = v63
		v715 = v64
		v717 = v66
		v718 = v67
		v720 = v69
		v722 = v71
		v724 = v73
		v728 = v77
		v730 = v79
		v731 = v80
		v733 = v82
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v93 = v67 & int32(1)
	if v93 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v94 = int32(0)
	goto L17
L16:
	;
	v94 = v68
	goto L17
L17:
	;
	v95 = v93 + v69
	if v95 < v89 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v521 = int32(0)
	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+14)))
	if v506 == v522 {
		v535 = v521
		v536 = v505
		v537 = v506
		v539 = v508
		v545 = v514
		v549 = v518
		v550 = v519
		goto L10
	} else {
		goto L112
	}
L19:
	;
	v109 = v94
	v110 = v95
	v118 = v77
	v122 = v80
	v123 = v82
	goto L22
L20:
	;
	v477 = v94
	v478 = v95
	v480 = v71
	v486 = v77
	v491 = v82
	goto L21
L21:
	;
	v505 = v477
	v506 = v478
	v508 = v480
	v514 = v486
	v518 = v80
	v519 = v491
	goto L18
L22:
	;
	v126 = v64 + int32(32)
	F_examine_indexcol_variable(m, v56, v66, v110, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v477 = v109
	v478 = v110
	v480 = int32(1)
	v486 = v260
	v491 = v261
	goto L21
L24:
	;
	return
L25:
	;
	v130 = v64 + int32(31)
	v131 = int32(0)
	v132 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v131)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v136 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
	if v110 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L27:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+28)))
	if v174 != 0 {
		goto L41
	} else {
		goto L42
	}
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+22)))
	v139 = v137 + v138
	v140 = *(*float32)(unsafe.Add(mBase, uint32(v139)+8))
	v142 = *(*float32)(unsafe.Add(mBase, uint32(v139)+16))
	v169 = base.F64_promote_f32(v142)
	v170 = base.F64_promote_f32(v140)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	if v144 == int32(16) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v169 = float64(2)
	v170 = v132
	goto L27
L32:
	;
	goto L33
L33:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v148 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v155 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+84))
	if v151 != int32(5) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v169 = float64(-1)
	v170 = v132
	goto L27
L37:
	;
	v169 = float64(0)
	v170 = v132
	goto L27
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v158 != int32(6) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+8)))
	switch v162 - int32(_a_F_btcostestimate_0) {
	case 0:
		goto L40
	default:
		goto L37
	case 5:
		v169 = float64(-1)
		v170 = v132
		goto L27
	}
L40:
	;
	v169 = float64(1)
	v170 = v132
	goto L27
L41:
	;
	v175 = base.F64_neg(base.F64_sub(float64(1), v170))
	goto L43
L42:
	;
	v175 = v169
	goto L43
L43:
	;
	if base.F64_gt(v175, float64(0)) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v178 = F_clamp_row_est(m, v175)
	mBase = m.M
	v204 = v178
	goto L26
L45:
	;
	goto L46
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v179 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v182)
	v204 = float64(200)
	goto L26
L48:
	;
	goto L49
L49:
	;
	v185 = *(*float64)(unsafe.Add(mBase, uint32(v179)+128))
	if base.F64_le(v185, float64(0)) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v188)
	v204 = float64(200)
	goto L26
L51:
	;
	goto L52
L52:
	;
	if base.F64_lt(v175, float64(0)) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v195 = F_clamp_row_est(m, base.F64_mul(v185, base.F64_neg(v175)))
	mBase = m.M
	v204 = v195
	goto L26
L54:
	;
	goto L55
L55:
	;
	if base.F64_lt(v185, float64(200)) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v198 = F_clamp_row_est(m, v185)
	mBase = m.M
	v204 = v198
	goto L26
L57:
	;
	goto L58
L58:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v199)
	v204 = float64(200)
	goto L26
L59:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+31)))
	if v262 != 0 {
		goto L80
	} else {
		goto L81
	}
L60:
	;
	if v205 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v249 = v205
	v251 = v118
	v252 = v123
	goto L62
L62:
	;
	if v249 == int32(0) {
		v260 = v251
		v261 = v252
		goto L59
	} else {
		goto L78
	}
L63:
	;
	v260 = int32(1)
	v261 = v123
	goto L59
L64:
	;
	goto L65
L65:
	;
	v211 = float64(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v66)+52))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v217 = F_get_opfamily_member(m, v213, v215, v215, int32(1))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L24
	} else {
		goto L67
	}
L66:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
	v249 = v248
	v251 = int32(1)
	v252 = v246
	goto L62
L67:
	;
	if v217 == int32(0) {
		v246 = v211
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v222 = v64 + int32(140)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
	v226 = F_get_attstatsslot(m, v222, v223, int32(3), v217, int32(2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L24
	} else {
		goto L69
	}
L69:
	;
	if v226 == int32(0) {
		v246 = v211
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v64)+160))
	v231 = *(*float32)(unsafe.Add(mBase, uint32(v230)))
	v232 = base.F64_promote_f32(v231)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v66)+64))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v235 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v236 = base.F64_neg(v232)
	goto L73
L72:
	;
	v236 = v232
	goto L73
L73:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v66)+40))
	if int32(1) < v239 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v242 = base.F64_mul(v236, float64(0.75))
	goto L76
L75:
	;
	v242 = v236
	goto L76
L76:
	;
	F_free_attstatsslot(m, v222)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	v246 = v242
	goto L66
L78:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
	m.T0[v255].(func(*base.Module, int32))(m, v249)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	v260 = v251
	v261 = v252
	goto L59
L80:
	;
	goto L23
L81:
	;
	if v109 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	if v263 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v422 = v204
	goto L84
L84:
	;
	if v109 != 0 {
		goto L107
	} else {
		goto L108
	}
L85:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v264 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v380 = v109
	goto L87
L87:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+76))
	v383 = int32(0)
	v385 = F_clauselist_selectivity(m, v56, v380, v382, v383, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L24
	} else {
		goto L102
	}
L88:
	;
	v350 = F_list_concat(m, v333, v109)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L24
	} else {
		goto L101
	}
L89:
	;
	v333 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v268 = int32(0)
	v279 = v268
	v281 = v268
	goto L92
L92:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298+v279<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v64)+140)) = v302
	v308 = F_list_make1_impl(m, int32(1), v64+int32(24))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L24
	} else {
		goto L94
	}
L93:
	;
	v333 = v317
	goto L88
L94:
	;
	v311 = F_predicate_implied_by(m, v308, v109, int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	if v311 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v315 = F_list_concat(m, v281, v308)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L24
	} else {
		goto L99
	}
L97:
	;
	v317 = v281
	goto L98
L98:
	;
	v319 = v279 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v319 < v320 {
		v279 = v319
		v281 = v317
		goto L92
	} else {
		goto L100
	}
L99:
	;
	v317 = v315
	goto L98
L100:
	;
	goto L93
L101:
	;
	v380 = v350
	goto L87
L102:
	;
	if base.F64_lt(v385, float64(0.005)) != 0 {
		goto L80
	} else {
		goto L103
	}
L103:
	;
	v390 = base.F64_nearest(base.F64_mul(v204, v385))
	v391 = float64(1)
	if base.F64_gt(v390, v391) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v394 = v390
	goto L106
L105:
	;
	v394 = v391
	goto L106
L106:
	;
	v422 = v394
	goto L84
L107:
	;
	v425 = v422
	goto L109
L108:
	;
	v425 = base.F64_add(v422, float64(1))
	goto L109
L109:
	;
	v426 = base.F64_mul(v122, v425)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if base.F64_gt(v426, base.F64_convert_i32_u(v427)) != 0 {
		goto L80
	} else {
		goto L110
	}
L110:
	;
	v430 = int32(1)
	v431 = int32(0)
	v433 = v110 + v430
	v434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+14)))
	if v434 <= v433 {
		v505 = v431
		v506 = v433
		v508 = v430
		v514 = v260
		v518 = v426
		v519 = v261
		goto L18
	} else {
		goto L111
	}
L111:
	;
	v109 = v431
	v110 = v433
	v118 = v260
	v122 = v426
	v123 = v261
	goto L22
L112:
	;
	v707 = v56
	v708 = v57
	v709 = v58
	v710 = v59
	v711 = v60
	v712 = v61
	v713 = v62
	v714 = v63
	v715 = v64
	v717 = v66
	v718 = v521
	v720 = v506
	v722 = v508
	v724 = v73
	v728 = v514
	v730 = v79
	v731 = v518
	v733 = v519
	goto L1
L113:
	;
	v704 = v75 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v704 < v705 {
		v67 = v686
		v68 = v687
		v69 = v537
		v71 = v690
		v73 = v692
		v75 = v704
		v76 = v695
		v77 = v545
		v79 = v698
		v80 = v699
		v82 = v550
		goto L8
	} else {
		goto L142
	}
L114:
	;
	v686 = v535
	v687 = v536
	v690 = v539
	v692 = v73
	v695 = v76
	v698 = v79
	v699 = v549
	goto L113
L115:
	;
	goto L116
L116:
	;
	v555 = int32(0)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	if v556 <= v555 {
		v686 = v535
		v687 = v536
		v690 = v539
		v692 = v73
		v695 = v76
		v698 = v79
		v699 = v549
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v568 = v555
	v570 = v535
	v571 = v536
	v574 = v539
	v576 = v73
	v579 = v76
	v582 = v79
	v583 = v549
	goto L118
L118:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v552)+12))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587+v568<<(uint(int32(2))%32))))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	switch v593 - int32(17) {
	case 0:
		goto L122
	default:
		goto L123
	case 3:
		goto L125
	case 20:
		goto L126
	case 35:
		goto L124
	}
L119:
	;
	v686 = v652
	v687 = v670
	v690 = v654
	v692 = v659
	v695 = v655
	v698 = v656
	v699 = v657
	goto L113
L120:
	;
	v659 = F_lappend(m, v576, v591)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L24
	} else {
		goto L136
	}
L121:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	if v639 == int32(0) {
		v652 = v570
		v654 = v634
		v655 = v635
		v656 = v582
		v657 = v636
		goto L120
	} else {
		goto L134
	}
L122:
	;
	v634 = v574
	v635 = v579
	v636 = v583
	v638 = v592 + int32(4)
	goto L121
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L24
	} else {
		goto L131
	}
L124:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v592)+8))
	v613 = base.B2i32(v611 == int32(0))
	v652 = v613 | v570
	v654 = v574
	v655 = v579
	v656 = v613 | v582
	v657 = v583
	goto L120
L125:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v592)+28))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+12))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	v602 = F_estimate_array_length(m, v56, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L24
	} else {
		goto L127
	}
L126:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v592)+8))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	v634 = v574
	v635 = int32(1)
	v636 = v583
	v638 = v598
	goto L121
L127:
	;
	if base.F64_gt(v602, float64(1)) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v607 = base.F64_mul(v583, v602)
	goto L130
L129:
	;
	v607 = v583
	goto L130
L130:
	;
	v634 = int32(1)
	v635 = v579
	v636 = v607
	v638 = v592 + int32(4)
	goto L121
L131:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v620
	F_errmsg_internal(m, int32(_a_F_btcostestimate_1), v64+int32(16))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L24
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_btcostestimate_2), int32(_a_F_btcostestimate_3), int32(_a_F_btcostestimate_4))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L24
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v66)+52))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v642+v537<<(uint(int32(2))%32))))
	v647 = F_get_op_opfamily_strategy(m, v639, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L24
	} else {
		goto L135
	}
L135:
	;
	v652 = base.B2i32(v647 == int32(3)) | v570
	v654 = v634
	v655 = v635
	v656 = v582
	v657 = v636
	goto L120
L136:
	;
	if v652&int32(1)|v655 != 0 {
		v670 = v571
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v672 = v568 + int32(1)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	if v672 < v673 {
		v568 = v672
		v570 = v652
		v571 = v670
		v574 = v654
		v576 = v659
		v579 = v655
		v582 = v656
		v583 = v657
		goto L118
	} else {
		goto L141
	}
L138:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v66)+40))
	if v664-int32(1) <= v537 {
		v670 = v571
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v668 = F_lappend(m, v571, v591)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	v670 = v668
	goto L137
L141:
	;
	goto L119
L142:
	;
	goto L9
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v715)+128)) = int32(1)
	*(*float64)(unsafe.Add(mBase, uint32(v715)+120)) = v916
	*(*float64)(unsafe.Add(mBase, uint32(v715)+104)) = v920
	F_genericcostestimate(m, v707, v708, v709, v715-int32(-64))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L24
	} else {
		goto L172
	}
L144:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v717)+40))
	v742 = int32(1)
	if (v718^int32(-1)|base.B2i32(v720 != v741-v742)|v722|v730)&v742 == int32(0) {
		v916 = v731
		v920 = float64(1)
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v717)+88))
	if v752 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L146
L148:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	if v753 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v869 = v724
	goto L150
L150:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v717)+12))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)+76))
	v872 = int32(0)
	v874 = F_clauselist_selectivity(m, v707, v869, v871, v872, v872)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L24
	} else {
		goto L165
	}
L151:
	;
	v839 = F_list_concat(m, v823, v724)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L24
	} else {
		goto L164
	}
L152:
	;
	v823 = int32(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v757 = int32(0)
	v768 = v757
	v771 = v757
	goto L155
L155:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v787+v768<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v715)+12)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v715)+140)) = v791
	v797 = F_list_make1_impl(m, int32(1), v715+int32(12))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L24
	} else {
		goto L157
	}
L156:
	;
	v823 = v806
	goto L151
L157:
	;
	v800 = F_predicate_implied_by(m, v797, v724, int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L24
	} else {
		goto L158
	}
L158:
	;
	if v800 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v804 = F_list_concat(m, v771, v797)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L24
	} else {
		goto L162
	}
L160:
	;
	v806 = v771
	goto L161
L161:
	;
	v808 = v768 + int32(1)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	if v808 < v809 {
		v768 = v808
		v771 = v806
		goto L155
	} else {
		goto L163
	}
L162:
	;
	v806 = v804
	goto L161
L163:
	;
	goto L156
L164:
	;
	v869 = v839
	goto L150
L165:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v717)+12))
	v877 = *(*float64)(unsafe.Add(mBase, uint32(v876)+128))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v717)+16))
	v883 = base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v879), float64(0.3333333)))
	if base.F64_lt(v731, v883) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v885 = v731
	goto L168
L167:
	;
	v885 = v883
	goto L168
L168:
	;
	v886 = float64(1)
	if base.F64_gt(v885, v886) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v889 = v885
	goto L171
L170:
	;
	v889 = v886
	goto L171
L171:
	;
	v916 = v889
	v920 = base.F64_nearest(base.F64_div(base.F64_mul(v874, v877), v889))
	goto L143
L172:
	;
	v930 = *(*float64)(unsafe.Add(mBase, _c_F_btcostestimate[0]))
	v931 = *(*float64)(unsafe.Add(mBase, uint32(v715)+120))
	v932 = *(*float64)(unsafe.Add(mBase, uint32(v715)+64))
	v933 = *(*float64)(unsafe.Add(mBase, uint32(v717)+24))
	if base.F64_gt(v933, float64(1)) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v717)+32))
	if v728 != 0 {
		v1005 = v733
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v938 = *(*float64)(unsafe.Add(mBase, uint32(v715)+72))
	v948 = v932
	v950 = v938
	goto L173
L175:
	;
	goto L176
L176:
	;
	v939 = F_log(m, v933)
	mBase = m.M
	v943 = base.F64_mul(base.F64_ceil(base.F64_div(v939, float64(0.6931471805599453))), v930)
	v946 = *(*float64)(unsafe.Add(mBase, uint32(v715)+72))
	v948 = base.F64_add(v932, v943)
	v950 = base.F64_add(base.F64_mul(v931, v943), v946)
	goto L173
L177:
	;
	v1011 = base.F64_mul(v930, base.F64_mul(base.F64_convert_i32_s(v951+int32(1)), float64(50)))
	*(*float64)(unsafe.Add(mBase, uint32(v710))) = base.F64_add(v948, v1011)
	*(*float64)(unsafe.Add(mBase, uint32(v711))) = base.F64_add(base.F64_mul(v931, v1011), v950)
	v1017 = *(*float64)(unsafe.Add(mBase, uint32(v715)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v712))) = v1017
	*(*float64)(unsafe.Add(mBase, uint32(v713))) = v1005
	v1020 = *(*float64)(unsafe.Add(mBase, uint32(v715)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v714))) = v1020
	m.G0 = v715 + int32(176)
	return
L178:
	;
	F_examine_indexcol_variable(m, v707, v717, int32(0), v715+int32(32))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L24
	} else {
		goto L179
	}
L179:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v715)+40))
	if v957 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v960 = *(*float64)(unsafe.Add(mBase, uint32(v715)+88))
	v1005 = v960
	goto L177
L181:
	;
	goto L182
L182:
	;
	v961 = float64(0)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v717)+52))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v717)+56))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	v967 = F_get_opfamily_member(m, v963, v965, v965, int32(1))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L24
	} else {
		goto L184
	}
L183:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v715)+40))
	if v997 == int32(0) {
		v1005 = v996
		goto L177
	} else {
		goto L195
	}
L184:
	;
	if v967 == int32(0) {
		v996 = v961
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v972 = v715 + int32(140)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v715)+40))
	v976 = F_get_attstatsslot(m, v972, v973, int32(3), v967, int32(2))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L24
	} else {
		goto L186
	}
L186:
	;
	if v976 == int32(0) {
		v996 = v961
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v715)+160))
	v981 = *(*float32)(unsafe.Add(mBase, uint32(v980)))
	v982 = base.F64_promote_f32(v981)
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v717)+64))
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984))))
	if v985 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v986 = base.F64_neg(v982)
	goto L190
L189:
	;
	v986 = v982
	goto L190
L190:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v717)+40))
	if int32(1) < v989 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v992 = base.F64_mul(v986, float64(0.75))
	goto L193
L192:
	;
	v992 = v986
	goto L193
L193:
	;
	F_free_attstatsslot(m, v972)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L24
	} else {
		goto L194
	}
L194:
	;
	v996 = v992
	goto L183
L195:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v715)+44))
	m.T0[v1000].(func(*base.Module, int32))(m, v997)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L24
	} else {
		goto L196
	}
L196:
	;
	v1005 = v996
	goto L177
}
func F_btfloat4sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1433)
	return int64(0)
}
func F_btfloat8sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1434)
	return int64(0)
}
func F_btinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v580 int32
	_ = v580
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v906 int32
	_ = v906
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v995 int32
	_ = v995
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1120 int32
	_ = v1120
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1298 int32
	_ = v1298
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1390 int32
	_ = v1390
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	v9 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v33 = F_index_form_tuple(m, v32, l1, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v39
	v41 = m.G0
	v43 = v41 - int32(480)
	m.G0 = v43
	v45 = F__bt_mkscankey(m, l0, v33)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v47 = int32(0)
	if l5 == v47 {
		v56 = v9
		v57 = v47
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v33
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+6)))
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+84)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+76)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v43)+92)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v43)+72)) = (v59&int32(_a_F_btinsert_0) + int32(7)) & int32(_a_F_btinsert_1)
	goto L8
L5:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
	if v51 != 0 {
		v56 = v9
		v57 = int32(1)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v52
	v56 = int32(1)
	v57 = v52
	goto L4
L7:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if l5 != int32(3) {
		goto L217
	} else {
		goto L218
	}
L8:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v107 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v885 != int32(1) {
		v906 = v880
		goto L7
	} else {
		goto L213
	}
L10:
	;
	if v56 == int32(0) {
		v906 = v57
		goto L7
	} else {
		goto L47
	}
L11:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v221 = F__bt_search(m, l0, l4, v218, v43+int32(80), int32(3), int32(1))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L46
	}
L12:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	if v110 == int32(-1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v113 = F_ReadBuffer(m, l0, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v113
	v116 = F_ConditionalLockBuffer(m, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v116 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v186 != 0 {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	F__bt_checkpage(m, l0, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_ReleaseBuffer(m, v118)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L37
	}
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v121 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	F_UnlockReleaseBuffer(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L36
	}
L22:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+16)))
	v141 = v140 + v139
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v142 != 0 {
		goto L21
	} else {
		goto L26
	}
L23:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v121^int32(-1))<<(uint(int32(2))%32))))
	v139 = v131
	goto L22
L24:
	;
	goto L25
L25:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v139 = v133 + v121<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L26:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)))
	if v143&int32(21) != int32(1) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v148 = int32(4)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+14)))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+12)))
	v151 = v149 - v150
	if v151 <= v148 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if base.Ui32(v154-int32(4)) <= base.Ui32(v157) {
		goto L21
	} else {
		goto L32
	}
L29:
	;
	v154 = v148
	goto L31
L30:
	;
	v154 = v151
	goto L31
L31:
	;
	goto L28
L32:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+12)))
	if base.B2i32(base.Ui32(v159) < base.Ui32(int32(25)))|base.B2i32((v159+int32(_a_F_btinsert_2))&int32(_a_F_btinsert_3) == int32(0)) != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v172 = F__bt_compare(m, l0, v170, v139, int32(1))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if int32(0) < v172 {
		v225 = int32(0)
		goto L10
	} else {
		goto L35
	}
L35:
	;
	goto L21
L36:
	;
	goto L16
L37:
	;
	goto L16
L38:
	;
	v212 = v186
	goto L40
L39:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+64)) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+56)) = v190
	v194 = F_smgropen(m, v43+int32(56), v187)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+16)) = int32(-1)
	goto L11
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v194
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+72))
	if v198 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v212 = v210
	goto L40
L43:
	;
	v206 = v198
	goto L45
L44:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+76))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)+72))
	v206 = v204
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+72)) = v206 + int32(1)
	goto L42
L46:
	;
	v225 = v221
	goto L10
L47:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+408)) = int32(4)
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v233 < v232 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v252) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v237+(v233^int32(-1))<<(uint(int32(2))%32))))
	v251 = v243
	goto L48
L50:
	;
	goto L51
L51:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v251 = v245 + v233<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	v260 = int32(base.Ui32(v252+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L54
L53:
	;
	v260 = int32(0)
	goto L54
L54:
	;
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+16)))
	v266 = F__bt_binsrch_insert(m, l0, v43+int32(68))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v268 = int32(0)
	v275 = v268
	v280 = v268
	v281 = v268
	v283 = v251
	v286 = v251 + v261
	v290 = int32(1)
	v291 = v266
	v296 = v232
	v298 = v260
	goto L62
L56:
	;
	goto L9
L57:
	;
	if v225 == int32(0) {
		goto L8
	} else {
		goto L208
	}
L58:
	;
	F_XactLockTableWait(m, v668, l0, v33, int32(5))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L207
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L201
	}
L60:
	;
	if base.B2i32(v787 == int32(0))&base.B2i32(l5 == int32(3)) != 0 {
		goto L59
	} else {
		goto L198
	}
L61:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v43)+412))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v43)+416))
	if v666 != 0 {
		goto L159
	} else {
		goto L160
	}
L62:
	;
	v308 = v275
	v313 = v280
	v323 = v290
	v324 = v291
	v329 = v296
	goto L65
L63:
	;
	if l5 != int32(2) {
		goto L61
	} else {
		goto L154
	}
L64:
	;
	goto L63
L65:
	;
	v336 = v324 & int32(_a_F_btinsert_4)
	v339 = v283 + int32(20) + v336<<(uint(int32(2))%32)
	v340 = int32(0)
	v343 = v340
	v346 = v308
	v351 = v313
	v358 = v340
	v361 = v323
	v367 = v329
	goto L67
L66:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v563 == int32(0) {
		v787 = v559
		goto L60
	} else {
		goto L132
	}
L67:
	;
	v374 = v298 & int32(_a_F_btinsert_4)
	if base.Ui32(v336) <= base.Ui32(v374) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	if base.Ui32(v336) < base.Ui32(v374) {
		goto L129
	} else {
		goto L130
	}
L69:
	;
	goto L68
L70:
	;
	v542 = int32(1)
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v537)+4)))
	if v534 < v545&int32(4095)-v542 {
		v343 = v534 + v542
		v346 = v536
		v351 = v537
		v358 = v542
		v361 = v538
		v367 = v541
		goto L67
	} else {
		goto L128
	}
L71:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+404)) = uint16(v437)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+400)) = v439
	if l5 == int32(3) {
		goto L100
	} else {
		goto L101
	}
L72:
	;
	v433 = v343
	v434 = v361
	v435 = v432
	v436 = v431
	goto L71
L73:
	;
	v431 = v358
	v432 = v398
	goto L72
L74:
	;
	if v367|base.B2i32(l5 != int32(3)) != 0 {
		v880 = int32(1)
		goto L56
	} else {
		goto L97
	}
L75:
	;
	if v281 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	if v358 == int32(0) {
		v553 = v346
		v554 = v351
		v556 = v361
		v559 = v367
		goto L69
	} else {
		goto L96
	}
L78:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+88)))
	if v336 == v378 {
		goto L74
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v358 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L80
L82:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+7)))
	if v399&int32(32) == int32(0) {
		goto L73
	} else {
		goto L91
	}
L83:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v383 = int32(_a_F_btinsert_5)
	if v382&v383 == v383 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v395 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+399)) = uint8(v395)
	v397 = v346
	v398 = v351
	goto L82
L86:
	;
	v553 = v339
	v554 = v351
	v556 = v361
	v559 = v367
	goto L69
L87:
	;
	goto L88
L88:
	;
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+399)) = uint8(v387)
	v389 = F__bt_compare(m, l0, v228, v283, v336)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v389 != 0 {
		v787 = v367
		goto L60
	} else {
		goto L90
	}
L90:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v397 = v339
	v398 = v283 + v391&int32(_a_F_btinsert_6)
	goto L82
L91:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+5)))
	if v404&int32(32) == int32(0) {
		goto L73
	} else {
		goto L92
	}
L92:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+2)))
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398))))
	v414 = v409 + (v398 + v410<<(uint(int32(16))%32))
	v415 = int32(1)
	v416 = int32(0)
	if v358 == v416 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v433 = v416
	v434 = int32(1)
	v435 = v414
	v436 = v415
	goto L71
L94:
	;
	goto L95
L95:
	;
	v431 = v415
	v432 = v414 + v343*int32(6)
	goto L72
L96:
	;
	v534 = v343
	v536 = v346
	v537 = v351
	v538 = v361
	v541 = v367
	goto L70
L97:
	;
	goto L59
L98:
	;
	if v436&int32(1) == int32(0) {
		v553 = v397
		v554 = v398
		v556 = v434
		v559 = v524
		goto L69
	} else {
		goto L127
	}
L99:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+399)))
	if v490 != int32(1) {
		v524 = v367
		goto L98
	} else {
		goto L115
	}
L100:
	;
	v444 = v43 + int32(400)
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+2)))
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444))))
	v450 = int32(16)
	v452 = v448 | v449<<(uint(v450)%32)
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+2)))
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229))))
	v457 = v453 | v454<<(uint(v450)%32)
	if base.Ui32(v452) < base.Ui32(v457) {
		v468 = int32(-1)
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	v488 = F_table_index_fetch_tuple_check(m, l4, v43+int32(400), v43+int32(408), v43+int32(399))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L113
	}
L103:
	;
	if v468 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	goto L103
L105:
	;
	if base.Ui32(v457) < base.Ui32(v452) {
		v468 = int32(1)
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)))
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+4)))
	if base.Ui32(v462) < base.Ui32(v463) {
		v468 = int32(-1)
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v468 = base.B2i32(base.Ui32(v463) < base.Ui32(v462))
	goto L104
L108:
	;
	v524 = int32(1)
	goto L98
L109:
	;
	goto L110
L110:
	;
	v478 = F_table_index_fetch_tuple_check(m, l4, v43+int32(400), v43+int32(408), v43+int32(399))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v478 == int32(0) {
		goto L99
	} else {
		goto L112
	}
L112:
	;
	goto L61
L113:
	;
	if v488 != 0 {
		goto L64
	} else {
		goto L114
	}
L114:
	;
	goto L99
L115:
	;
	if v436&int32(1) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if v434&int32(1) == int32(0) {
		v534 = v433
		v536 = v397
		v537 = v398
		v538 = v434
		v541 = v367
		goto L70
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v281 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+4)))
	if v433 != v499&int32(4095)-int32(1) {
		v534 = v433
		v536 = v397
		v537 = v398
		v538 = v434
		v541 = v367
		goto L70
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v506 = v281
	goto L123
L122:
	;
	v506 = v505
	goto L123
L123:
	;
	v507 = F_BufferBeginSetHintBits(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	if v507 == int32(0) {
		v524 = v367
		goto L98
	} else {
		goto L125
	}
L125:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = v511 | int32(_a_F_btinsert_5)
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286)+12)))
	v517 = v515 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+12)) = uint16(v517)
	v519 = int32(1)
	F_BufferFinishSetHintBits(m, v506, v519, v519)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v524 = v367
	goto L98
L127:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+399)))
	v534 = v433
	v536 = v397
	v537 = v398
	v538 = (v529 | (v436 ^ int32(1))) & v434
	v541 = v524
	goto L70
L128:
	;
	v553 = v536
	v554 = v537
	v556 = v538
	v559 = v541
	goto L69
L129:
	;
	v308 = v553
	v313 = v554
	v323 = v556
	v324 = v324 + int32(1)
	v329 = v559
	goto L65
L130:
	;
	goto L131
L131:
	;
	goto L66
L132:
	;
	v567 = F__bt_compare(m, l0, v228, v283, int32(1))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	if v567 != 0 {
		v787 = v559
		goto L60
	} else {
		goto L134
	}
L134:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v570 = v569
	v580 = v281
	goto L136
L135:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v650 != 0 {
		goto L148
	} else {
		goto L149
	}
L136:
	;
	v602 = F__bt_relandgetbuf(m, l0, v580, v570, int32(1))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L145
	}
L138:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621)+16)))
	v623 = v622 + v621
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+12)))
	if v624&int32(20) == int32(0) {
		goto L135
	} else {
		goto L143
	}
L139:
	;
	if v602 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607+(v602^int32(-1))<<(uint(int32(2))%32))))
	v621 = v613
	goto L138
L141:
	;
	goto L142
L142:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v621 = v615 + v602<<(uint(int32(13))%32) + int32(-8192)
	goto L138
L143:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v629 != 0 {
		v570 = v629
		v580 = v602
		goto L136
	} else {
		goto L144
	}
L144:
	;
	goto L137
L145:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v634 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btinsert_7), v43+int32(16))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_btinsert_8), int32(757), int32(_a_F_btinsert_9))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
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
	v651 = int32(2)
	goto L150
L149:
	;
	v651 = int32(1)
	goto L150
L150:
	;
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v652) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v660 = int32(base.Ui32(v652+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L153
L152:
	;
	v660 = int32(0)
	goto L153
L153:
	;
	v275 = v553
	v280 = v554
	v281 = v602
	v283 = v621
	v286 = v623
	v290 = v556
	v291 = v651
	v296 = v559
	v298 = v660
	goto L62
L154:
	;
	if v281 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_UnlockReleaseBuffer(m, v281)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v880 = int32(0)
	goto L56
L158:
	;
	goto L157
L159:
	;
	v668 = v666
	goto L161
L160:
	;
	v668 = v667
	goto L161
L161:
	;
	if v668 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v281 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+404)) = uint16(v683)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+400)) = v685
	v691 = F_table_index_fetch_tuple_check(m, l4, v43+int32(400), int32(_a_F_btinsert_10), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L172
	}
L165:
	;
	F_UnlockReleaseBuffer(m, v281)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v671 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+84)) = uint8(v671)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v43)+444))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	F_UnlockReleaseBuffer(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v677 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v677
	if v673 == v677 {
		goto L58
	} else {
		goto L170
	}
L170:
	;
	F_SpeculativeInsertionWait(m, v668, v673)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	goto L57
L172:
	;
	if v691 == int32(0) {
		v787 = v367
		goto L60
	} else {
		goto L173
	}
L173:
	;
	v695 = int32(0)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v696 < v695 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	F_CheckForSerializableConflictIn(m, l0, v695, v715)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L178
	}
L175:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[2]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v700+(v696^int32(-1))*int32(56))+16))
	v715 = v706
	goto L174
L176:
	;
	goto L177
L177:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[3]))
	v709 = int32(56)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v708+v696*v709-v709)+16))
	v715 = v714
	goto L174
L178:
	;
	if v281 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	F_UnlockReleaseBuffer(m, v281)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	F_UnlockReleaseBuffer(m, v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L183
	}
L182:
	;
	goto L181
L183:
	;
	v723 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+84)) = uint8(v723)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v723
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v729 = v43 + int32(128)
	v731 = v43 + int32(96)
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+6)))
	if v723 <= base.I32_extend16_s(v734) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v745 = F_BuildIndexValueDescription(m, l0, v729, v731)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L188
	}
L185:
	;
	v738 = int32(8)
	goto L187
L186:
	;
	v738 = int32(16)
	goto L187
L187:
	;
	F_index_deform_tuple_internal(m, v727, v729, v731, v229+v738, v229+int32(8), int32(base.Ui32(v734)>>(uint(int32(15))%32)))
	mBase = m.M
	goto L184
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_errcode(m, int32(83906754))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v754 + int32(4)
	F_errmsg(m, int32(_a_F_btinsert_11), v43+int32(48))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	if v745 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v745
	v767 = F_errdetail(m, int32(_a_F_btinsert_12), v43+int32(32))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_errtableconstraint(m, l4, v769+int32(4))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L196
	}
L195:
	;
	goto L194
L196:
	;
	F_errfinish(m, int32(_a_F_btinsert_8), int32(676), int32(_a_F_btinsert_9))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	v793 = int32(1)
	if v281 == int32(0) {
		v880 = v793
		goto L56
	} else {
		goto L199
	}
L199:
	;
	F_UnlockReleaseBuffer(m, v281)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v880 = v793
	goto L56
L201:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v815 + int32(4)
	F_errmsg(m, int32(_a_F_btinsert_13), v43)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errhint(m, int32(_a_F_btinsert_14), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_errtableconstraint(m, l4, v826+int32(4))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_btinsert_8), int32(780), int32(_a_F_btinsert_9))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	goto L57
L208:
	;
	v844 = v225
	goto L209
L209:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v844)+8))
	F_pfree(m, v844)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L211
	}
L210:
	;
	goto L8
L211:
	;
	if v872 != 0 {
		v844 = v872
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v33
	v906 = v880
	goto L7
L214:
	;
	if v225 != 0 {
		goto L322
	} else {
		goto L323
	}
L215:
	;
	v1331 = v43 + int32(68)
	v1332 = F__bt_binsrch_insert(m, l0, v1331)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L315
	}
L216:
	;
	if v56 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L217:
	;
	v923 = int32(0)
	if v920 < v923 {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	goto L219
L219:
	;
	F_UnlockReleaseBuffer(m, v920)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L272
	}
L220:
	;
	F_CheckForSerializableConflictIn(m, l0, v923, v942)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L224
	}
L221:
	;
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[2]))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v927+(v920^int32(-1))*int32(56))+16))
	v942 = v933
	goto L220
L222:
	;
	goto L223
L223:
	;
	v935 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[3]))
	v936 = int32(56)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v935+v920*v936-v936)+16))
	v942 = v941
	goto L220
L224:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v946 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v964)+16)))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if base.Ui32(int32(2705)) <= base.Ui32(v966) {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v950+(v946^int32(-1))<<(uint(int32(2))%32))))
	v964 = v956
	goto L225
L227:
	;
	goto L228
L228:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v964 = v958 + v946<<(uint(int32(13))%32) + int32(-8192)
	goto L225
L229:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	F__bt_check_third_page(m, l0, l4, v969, v964, v970)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v973 = v964 + v965
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
	if v974 != 0 {
		goto L216
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	v975 = int32(4)
	v976 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v964)+14)))
	v977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v964)+12)))
	v978 = v976 - v977
	if v978 <= v975 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if base.Ui32(v984) <= base.Ui32(v981-int32(4)) {
		goto L215
	} else {
		goto L238
	}
L235:
	;
	v981 = v975
	goto L237
L236:
	;
	v981 = v978
	goto L237
L237:
	;
	goto L234
L238:
	;
	v986 = v964
	v995 = v973
	goto L239
L239:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+12)))
	if v1017&int32(64) != 0 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	goto L215
L241:
	;
	v1023 = int32(0)
	F__bt_delete_or_dedup_one_page(m, l0, l4, v43+int32(68), int32(1), v1023, v1023, v1023)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+84)))
	if v1039 != int32(1) {
		goto L250
	} else {
		goto L251
	}
L244:
	;
	v1028 = int32(4)
	v1029 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v986)+14)))
	v1030 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v986)+12)))
	v1031 = v1029 - v1030
	if v1031 <= v1028 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if base.Ui32(v1037) <= base.Ui32(v1034-int32(4)) {
		goto L215
	} else {
		goto L249
	}
L246:
	;
	v1034 = v1028
	goto L248
L247:
	;
	v1034 = v1031
	goto L248
L248:
	;
	goto L245
L249:
	;
	goto L243
L250:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v995)+4))
	if v1059 == int32(0) {
		goto L215
	} else {
		goto L257
	}
L251:
	;
	v1042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+88)))
	v1043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+86)))
	if base.Ui32(v1042) < base.Ui32(v1043) {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v1045 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v986)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1045) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1053 = int32(base.Ui32(v1045+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L255
L254:
	;
	v1053 = int32(0)
	goto L255
L255:
	;
	if base.Ui32(v1042) <= base.Ui32(v1053&int32(_a_F_btinsert_4)) {
		goto L215
	} else {
		goto L256
	}
L256:
	;
	goto L250
L257:
	;
	v1063 = F__bt_compare(m, l0, v945, v986, int32(1))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	if v1063 != 0 {
		goto L215
	} else {
		goto L259
	}
L259:
	;
	v1066 = Fn14349(m, int64(32))
	mBase = m.M
	goto L260
L260:
	;
	if base.Ui32(v1066) < base.Ui32(int32(42949673)) {
		goto L215
	} else {
		goto L261
	}
L261:
	;
	F__bt_stepright(m, l0, l4, v43+int32(68), v225)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v1073 < int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1092 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+16)))
	v1094 = int32(4)
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+14)))
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+12)))
	v1097 = v1095 - v1096
	if v1097 <= v1094 {
		goto L268
	} else {
		goto L269
	}
L264:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1077+(v1073^int32(-1))<<(uint(int32(2))%32))))
	v1091 = v1083
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v1091 = v1085 + v1073<<(uint(int32(13))%32) + int32(-8192)
	goto L263
L267:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if base.Ui32(v1100-int32(4)) < base.Ui32(v1103) {
		v986 = v1091
		v995 = v1092 + v1091
		goto L239
	} else {
		goto L271
	}
L268:
	;
	v1100 = v1094
	goto L270
L269:
	;
	v1100 = v1097
	goto L270
L270:
	;
	goto L267
L271:
	;
	goto L240
L272:
	;
	goto L214
L273:
	;
	v1283 = int32(4)
	v1284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+14)))
	v1285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+12)))
	v1286 = v1284 - v1285
	if v1286 <= v1283 {
		goto L310
	} else {
		goto L311
	}
L274:
	;
	v1252 = v964
	v1253 = l6
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+86)))
	v1110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+88)))
	v1112 = l6 | base.B2i32(base.Ui32(v1109) < base.Ui32(v1110))
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+84)))
	if base.B2i32(v1113 != int32(1))|base.B2i32(base.Ui32(v1110) < base.Ui32(v1109)) == int32(0) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v964)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1120) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v973)+4))
	if v1133 == int32(0) {
		v1252 = v964
		v1253 = v1112
		goto L273
	} else {
		goto L284
	}
L280:
	;
	v1128 = int32(base.Ui32(v1120+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L282
L281:
	;
	v1128 = int32(0)
	goto L282
L282:
	;
	if base.Ui32(v1110) <= base.Ui32(v1128&int32(_a_F_btinsert_4)) {
		v1252 = v964
		v1253 = v1112
		goto L273
	} else {
		goto L283
	}
L283:
	;
	goto L279
L284:
	;
	v1137 = F__bt_compare(m, l0, v945, v964, int32(1))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	if v1137 <= int32(0) {
		v1252 = v964
		v1253 = v1112
		goto L273
	} else {
		goto L286
	}
L286:
	;
	F__bt_stepright(m, l0, l4, v43+int32(68), v225)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if int32(0) <= v1145 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v1164 = v1163
	goto L292
L289:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v1163 = v1149 + v1145<<(uint(int32(13))%32) + int32(-8192)
	goto L288
L290:
	;
	goto L291
L291:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1156+(v1145^int32(-1))<<(uint(int32(2))%32))))
	v1163 = v1162
	goto L288
L292:
	;
	v1195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1164)+16)))
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+84)))
	if v1196 != int32(1) {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	v1252 = v1164
	v1253 = int32(1)
	goto L273
L294:
	;
	goto L293
L295:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1164+v1195)+4))
	if v1217 == int32(0) {
		goto L294
	} else {
		goto L302
	}
L296:
	;
	v1199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+88)))
	v1200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+86)))
	if base.Ui32(v1199) < base.Ui32(v1200) {
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1164)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1202) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1210 = int32(base.Ui32(v1202+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L300
L299:
	;
	v1210 = int32(0)
	goto L300
L300:
	;
	if base.Ui32(v1199) <= base.Ui32(v1210&int32(_a_F_btinsert_4)) {
		goto L294
	} else {
		goto L301
	}
L301:
	;
	goto L295
L302:
	;
	v1220 = int32(1)
	v1222 = F__bt_compare(m, l0, v945, v1164, v1220)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	if v1222 <= int32(0) {
		v1252 = v1164
		v1253 = v1220
		goto L273
	} else {
		goto L304
	}
L304:
	;
	F__bt_stepright(m, l0, l4, v43+int32(68), v225)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v1230 < int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1234+(v1230^int32(-1))<<(uint(int32(2))%32))))
	v1248 = v1240
	goto L308
L307:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v1248 = v1242 + v1230<<(uint(int32(13))%32) + int32(-8192)
	goto L308
L308:
	;
	v1164 = v1248
	goto L292
L309:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if base.Ui32(v1292) <= base.Ui32(v1289-int32(4)) {
		goto L215
	} else {
		goto L313
	}
L310:
	;
	v1289 = v1283
	goto L312
L311:
	;
	v1289 = v1286
	goto L312
L312:
	;
	goto L309
L313:
	;
	F__bt_delete_or_dedup_one_page(m, l0, l4, v43+int32(68), int32(0), v56, v1253, l6)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	goto L215
L315:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v43)+92))
	if v1334 == int32(-1) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1338 = int32(0)
	F__bt_delete_or_dedup_one_page(m, l0, l4, v1331, int32(1), v1338, v1338, v1338)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L319
	}
L317:
	;
	v1348 = v1332
	v1349 = v1334
	goto L318
L318:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	v1351 = int32(0)
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	F__bt_insertonpg(m, l0, l4, v45, v1350, v1351, v225, v33, v1352, v1348, v1349, v1351)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+92)) = int32(0)
	v1345 = F__bt_binsrch_insert(m, l0, v1331)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v43)+92))
	v1348 = v1345
	v1349 = v1347
	goto L318
L321:
	;
	goto L214
L322:
	;
	v1390 = v225
	goto L325
L323:
	;
	goto L324
L324:
	;
	F_pfree(m, v45)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L329
	}
L325:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+8))
	F_pfree(m, v1390)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L327
	}
L326:
	;
	goto L324
L327:
	;
	if v1418 != 0 {
		v1390 = v1418
		goto L325
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	m.G0 = v43 + int32(480)
	F_pfree(m, v33)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	return v906
}
func F_btint24cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(v4 < v3) - base.B2i32(v3 < v4))
}
func F_btint28cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(v5 < v4) - base.B2i32(v4 < v5))
}
func F_btint2fastcmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	return base.I32_extend16_s(base.I32_wrap_i64(l0)) - base.I32_extend16_s(base.I32_wrap_i64(l1))
}
func F_btint48cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(v5 < v4) - base.B2i32(v4 < v5))
}
func F_btoid8sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(208)
	return int64(0)
}
func F_btproperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v11 = base.B2i32(l2 == int32(7)) & base.B2i32(l1 != int32(0))
	if v11 != 0 {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v12)
	} else {
	}
	return v11
}
func F_btrecordcmp(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v2)
	}
}
func F_bttextcmp(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v48 = int32(1)
			if v17&v48 != 0 {
				v52 = v48
			} else {
				v52 = int32(4)
			}
			v54 = int32(1)
			if v16&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			if v16 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v16&v76 != 0 {
					v88 = int32(base.Ui32(v16)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v9+v52, v46, v14+v58, v88, v47)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v91 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v95 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_s(v89)
							}
						} else {
							return base.I64_extend_i32_s(v89)
						}
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v95 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(v89)
						}
					} else {
						return base.I64_extend_i32_s(v89)
					}
				}
			}
		}
	}
}
func F_bttidcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.I64_extend_i32_s(v27)
}
func F_btvarstrequalimage(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_btvarstrequalimage_0), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_btvarstrequalimage_1), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_btvarstrequalimage_2), int32(1337), int32(_a_F_btvarstrequalimage_3))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v27 = F_pg_newlocale_from_collation(m, v2)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
			return v29
		}
	}
}
func F_buildNSItemFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v17 = v13 << (uint(int32(5)) % 32)
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v18 = F_palloc0(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = int32(0)
	goto L6
L6:
	;
	v34 = int32(0)
	if l2 == v34 {
		v45 = v34
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 == int32(0) {
		v54 = v34
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v39 <= v27 {
		v45 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v45 = v41 + v27<<(uint(int32(2))%32)
	goto L8
L11:
	;
	if l4 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v48 <= v27 {
		v54 = v34
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v54 = v50 + v27<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v82 = v18 + v27<<(uint(int32(5))%32)
	v84 = v27 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)) = uint16(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = l1
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v65+v27<<(uint(int32(2))%32))))
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+28)) = uint16(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v94
	v27 = v84
	goto L6
L15:
	;
	v68 = F_palloc(m, int32(28))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L19
	}
L16:
	;
	v57 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v54 == v57)|(base.B2i32(v45 == v57)|base.B2i32(v61 <= v27)) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v65 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v68)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
	return v68
}
func F_build_backup_content(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v94 int32
	_ = v94
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int64
	_ = v171
	var v176 int64
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	v9 = m.G0
	v11 = v9 - int32(544)
	m.G0 = v11
	v14 = v11 + int32(336)
	F_initStringInfo(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[0]))
		v27 = F_pg_localtime(m, l0+int32(1056), v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = F_pg_strftime(m, v11+int32(416), int32(128), int32(_a_F_build_backup_content_0), v27)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v32
				v35 = int64(*(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[1])))
				v36 = base.I64_div_u_s(v31, v35)
				v38 = base.I64_div_u_s(int64(4294967296), v35)
				v39 = base.I64_div_u_s(v36, v38)
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+196)) = uint32(v39)
				v42 = v36 - v38*v39
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+200)) = uint32(v42)
				v45 = v11 + int32(352)
				v50 = F_pg_snprintf(m, v45, int32(64), int32(_a_F_build_backup_content_1), v11+int32(192))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+180)) = uint32(v52)
					v55 = int64(base.Ui64(v52) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+176)) = uint32(v55)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+184)) = v45
					F_appendStringInfo(m, v14, int32(_a_F_build_backup_content_2), v11+int32(176))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if l1 != 0 {
							v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v64
							v67 = int64(*(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[1])))
							v68 = base.I64_div_u_s(v63, v67)
							v70 = base.I64_div_u_s(int64(4294967296), v67)
							v71 = base.I64_div_u_s(v68, v70)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+164)) = uint32(v71)
							v74 = v68 - v70*v71
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+168)) = uint32(v74)
							v77 = v11 + int32(208)
							v82 = F_pg_snprintf(m, v77, int32(64), int32(_a_F_build_backup_content_1), v11+int32(160))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+148)) = uint32(v84)
								v87 = int64(base.Ui64(v84) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+144)) = uint32(v87)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+152)) = v77
								F_appendStringInfo(m, v14, int32(_a_F_build_backup_content_3), v11+int32(144))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1048))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+132)) = uint32(v99)
									v102 = int64(base.Ui64(v99) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+128)) = uint32(v102)
									v105 = v11 + int32(336)
									F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_4), v11+int32(128))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoString(m, v105, int32(_a_F_build_backup_content_5))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
											if v116 != 0 {
												v117 = int32(_a_F_build_backup_content_6)
											} else {
												v117 = int32(_a_F_build_backup_content_7)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v117
											F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_8), v11+int32(112))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(416)
												F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_9), v11+int32(96))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l0
													F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_10), v11+int32(80))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v138
														F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_11), v11-int32(-64))
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return int32(0)
														} else {
															if l1 != 0 {
																v146 = v11 + int32(208)
																v152 = *(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[0]))
																v153 = F_pg_localtime(m, l0+int32(1104), v152)
																mBase = m.M
																v154 = m.ExcPending
																if v154 != 0 {
																	return int32(0)
																} else {
																	v155 = F_pg_strftime(m, v146, int32(128), int32(_a_F_build_backup_content_0), v153)
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v146
																		F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_12), v11+int32(48))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return int32(0)
																		} else {
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v163
																			F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_13), v11+int32(32))
																			mBase = m.M
																			v169 = m.ExcPending
																			if v169 != 0 {
																				return int32(0)
																			} else {
																				v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																				if v171 != int64(0) {
																					*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v171)
																					v176 = int64(base.Ui64(v171) >> (uint(int64(32)) % 64))
																					*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v176)
																					v179 = v11 + int32(336)
																					F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_14), v11+int32(16))
																					mBase = m.M
																					v184 = m.ExcPending
																					if v184 != 0 {
																						return int32(0)
																					} else {
																						v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v185
																						F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_15), v11)
																						mBase = m.M
																						v189 = m.ExcPending
																						if v189 != 0 {
																							return int32(0)
																						} else {
																							v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																							m.G0 = v11 + int32(544)
																							return v191
																						}
																					}
																				} else {
																					v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																					m.G0 = v11 + int32(544)
																					return v191
																				}
																			}
																		}
																	}
																}
															} else {
																v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																if v171 != int64(0) {
																	*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v171)
																	v176 = int64(base.Ui64(v171) >> (uint(int64(32)) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v176)
																	v179 = v11 + int32(336)
																	F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_14), v11+int32(16))
																	mBase = m.M
																	v184 = m.ExcPending
																	if v184 != 0 {
																		return int32(0)
																	} else {
																		v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v185
																		F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_15), v11)
																		mBase = m.M
																		v189 = m.ExcPending
																		if v189 != 0 {
																			return int32(0)
																		} else {
																			v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																			m.G0 = v11 + int32(544)
																			return v191
																		}
																	}
																} else {
																	v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																	m.G0 = v11 + int32(544)
																	return v191
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
							v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1048))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+132)) = uint32(v99)
							v102 = int64(base.Ui64(v99) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+128)) = uint32(v102)
							v105 = v11 + int32(336)
							F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_4), v11+int32(128))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								F_appendStringInfoString(m, v105, int32(_a_F_build_backup_content_5))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
									if v116 != 0 {
										v117 = int32(_a_F_build_backup_content_6)
									} else {
										v117 = int32(_a_F_build_backup_content_7)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v117
									F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_8), v11+int32(112))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(416)
										F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_9), v11+int32(96))
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l0
											F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_10), v11+int32(80))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v138
												F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_11), v11-int32(-64))
												mBase = m.M
												v144 = m.ExcPending
												if v144 != 0 {
													return int32(0)
												} else {
													if l1 != 0 {
														v146 = v11 + int32(208)
														v152 = *(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[0]))
														v153 = F_pg_localtime(m, l0+int32(1104), v152)
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return int32(0)
														} else {
															v155 = F_pg_strftime(m, v146, int32(128), int32(_a_F_build_backup_content_0), v153)
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v146
																F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_12), v11+int32(48))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v163
																	F_appendStringInfo(m, v105, int32(_a_F_build_backup_content_13), v11+int32(32))
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return int32(0)
																	} else {
																		v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																		if v171 != int64(0) {
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v171)
																			v176 = int64(base.Ui64(v171) >> (uint(int64(32)) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v176)
																			v179 = v11 + int32(336)
																			F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_14), v11+int32(16))
																			mBase = m.M
																			v184 = m.ExcPending
																			if v184 != 0 {
																				return int32(0)
																			} else {
																				v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v185
																				F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_15), v11)
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return int32(0)
																				} else {
																					v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																					m.G0 = v11 + int32(544)
																					return v191
																				}
																			}
																		} else {
																			v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																			m.G0 = v11 + int32(544)
																			return v191
																		}
																	}
																}
															}
														}
													} else {
														v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
														if v171 != int64(0) {
															*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v171)
															v176 = int64(base.Ui64(v171) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v176)
															v179 = v11 + int32(336)
															F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_14), v11+int32(16))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return int32(0)
															} else {
																v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v185
																F_appendStringInfo(m, v179, int32(_a_F_build_backup_content_15), v11)
																mBase = m.M
																v189 = m.ExcPending
																if v189 != 0 {
																	return int32(0)
																} else {
																	v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
																	m.G0 = v11 + int32(544)
																	return v191
																}
															}
														} else {
															v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
															m.G0 = v11 + int32(544)
															return v191
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
func F_build_bound_expr(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4)+8)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if l3 != 0 {
		v17 = int32(4)
	} else {
		v17 = int32(5)
	}
	if l3 != 0 {
		v20 = int32(2)
	} else {
		v20 = int32(1)
	}
	if l2 != 0 {
		v21 = v17
	} else {
		v21 = v20
	}
	v22 = F_get_opfamily_member(m, l5, v14, v14, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 == int32(0) {
			return int32(0)
		} else {
			v34 = F_makeConst(m, v14, int32(-1), v11, v13, l1, int32(0), v12&int32(1))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = F_make_opclause(m, v22, l0, v34, l6)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					return v36
				}
			}
		}
	}
}
func F_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v132 int64
	_ = v132
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v15 = v13 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+82)))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_build_datatype_2))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L53
	}
L2:
	;
	v18 = F_palloc(m, int32(48))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_build_datatype_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L49
	}
L5:
	;
	return int32(0)
L6:
	;
	v24 = F_pstrdup(m, v15+int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	switch v29 - int32(98) {
	case 0, 3, 11, 16:
		goto L13
	case 1:
		goto L11
	case 2:
		goto L14
	default:
		goto L10
	case 14:
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v54
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v56)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v58)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v54 = int32(2)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_build_datatype_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	v54 = int32(1)
	goto L8
L12:
	;
	if v27 != int32(2249) {
		goto L9
	} else {
		goto L17
	}
L13:
	;
	v54 = int32(0)
	goto L8
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v33 = F_type_is_rowtype(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if v33 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L11
L18:
	;
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+79)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
	F_errmsg_internal(m, int32(_a_F_build_datatype_7), v11)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_build_datatype_4), int32(2038), int32(_a_F_build_datatype_5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v63 = l2
	goto L23
L22:
	;
	v63 = v62
	goto L23
L23:
	;
	if l2 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v64 = v63
	goto L26
L25:
	;
	v64 = v62
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	switch v66 - int32(98) {
	case 0:
		goto L30
	default:
		goto L28
	case 2:
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l1
	if v100 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v97)
	v100 = v54
	goto L27
L29:
	;
	v81 = int32(0)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+76)))
	if v82 != int32(_a_F_build_datatype_6) {
		v94 = v81
		v95 = v54
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	if v70 == v69 {
		v79 = v69
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v79)
	v100 = v54
	goto L27
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v73 != int32(_a_F_build_datatype_0) {
		v79 = v69
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
	v79 = base.B2i32(v76 != int32(112))
	goto L31
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v94)
	v100 = v95
	goto L27
L35:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
	if v85 == int32(112) {
		v94 = v81
		v95 = v54
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v89 = F_get_base_element_type(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v94 = base.B2i32(v89 != int32(0))
	v95 = v93
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v132
	m.G0 = v11 + int32(48)
	return v18
L39:
	;
	v127 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+28)) = v127
	v132 = v127
	goto L38
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v104 == int32(2249) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v108 = F_lookup_type_cache(m, v104, int32(_a_F_build_datatype_1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+13)))
	if v110 == int32(100) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+300))
	v115 = F_lookup_type_cache(m, v113, int32(256))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	v117 = v108
	goto L45
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+188))
	if v118 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v117 = v115
	goto L45
L47:
	;
	v121 = F_copyObjectImpl(m, l3)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v121
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v117)+192))
	v132 = v125
	goto L38
L49:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v15 + int32(4)
	F_errmsg(m, int32(_a_F_build_datatype_8), v11+int32(32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_build_datatype_4), int32(2007), int32(_a_F_build_datatype_5))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v166 = F_format_type_be(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v166
	F_errmsg(m, int32(_a_F_build_datatype_3), v11+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_build_datatype_4), int32(2089), int32(_a_F_build_datatype_5))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_build_dummy_expanded_header(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 == int32(0) {
		v9 = F_expanded_record_fetch_tupdesc(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = v9
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			if v12 == int32(0) {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v20 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_build_dummy_expanded_header_0), int32(0), int32(1024), int32(_a_F_build_dummy_expanded_header_1))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v20
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v26 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
						if v27 == v25 {
							v64 = v26
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								base.MemoryFill(m, v34, int32(0), int32(120))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v41 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
								v43 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
								*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
								*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
								v52 = v34 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
								v64 = v34
								*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
								*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
								v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
								return
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				}
			} else {
				F_MemoryContextReset(m, v12)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v26 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
						if v27 == v25 {
							v64 = v26
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								base.MemoryFill(m, v34, int32(0), int32(120))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v41 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
								v43 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
								*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
								*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
								v52 = v34 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
								v64 = v34
								*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
								*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
								v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
								return
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				}
			}
		}
	} else {
		v11 = v6
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		if v12 == int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_build_dummy_expanded_header_0), int32(0), int32(1024), int32(_a_F_build_dummy_expanded_header_1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v20
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v26 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
					if v27 == v25 {
						v64 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.MemoryFill(m, v34, int32(0), int32(120))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v41 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
						v43 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
						v52 = v34 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
						v64 = v34
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					}
				}
			}
		} else {
			F_MemoryContextReset(m, v12)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v26 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
					if v27 == v25 {
						v64 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v34 = F_MemoryContextAlloc(m, v29, v25*int32(9)+int32(120))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.MemoryFill(m, v34, int32(0), int32(120))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v41 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
						v43 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
						v52 = v34 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(3))%32)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
						v64 = v34
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					}
				}
			}
		}
	}
}
func F_build_implied_join_equality(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = F_copyObjectImpl(m, l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_copyObjectImpl(m, l4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_make_opclause(m, l1, v8, v12, l2)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = int32(0)
				v22 = F_make_restrictinfo(m, l0, v14, int32(1), v17, v17, v17, l6, l5, v17, v17)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
					if v24 != 0 {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
						if v53 != 0 {
							F_check_memoizable(m, v22)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								return v22
							}
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
							if v54 == int32(0) {
								F_check_memoizable(m, v22)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									return v22
								}
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
								if v57 != int32(17) {
									F_check_memoizable(m, v22)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										return v22
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
									if v60 == int32(0) {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
										if v63 != int32(2) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
											v69 = F_exprType(m, v68)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = F_op_hashjoinable(m, v66, v69)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													if v71 == int32(0) {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v75 = F_contain_volatile_functions(m, v22)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return int32(0)
														} else {
															if v75 != 0 {
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
															}
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
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
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						if v25 == int32(0) {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
							if v53 != 0 {
								F_check_memoizable(m, v22)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									return v22
								}
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
								if v54 == int32(0) {
									F_check_memoizable(m, v22)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										return v22
									}
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
									if v57 != int32(17) {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
										if v60 == int32(0) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
											if v63 != int32(2) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
												v69 = F_exprType(m, v68)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = F_op_hashjoinable(m, v66, v69)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														if v71 == int32(0) {
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
															}
														} else {
															v75 = F_contain_volatile_functions(m, v22)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return int32(0)
															} else {
																if v75 != 0 {
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																}
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
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
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if v28 != int32(17) {
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
								if v53 != 0 {
									F_check_memoizable(m, v22)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										return v22
									}
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
									if v54 == int32(0) {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
										if v57 != int32(17) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
											if v60 == int32(0) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
												if v63 != int32(2) {
													F_check_memoizable(m, v22)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v22
													}
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
													v69 = F_exprType(m, v68)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = F_op_hashjoinable(m, v66, v69)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int32(0)
														} else {
															if v71 == int32(0) {
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
																}
															} else {
																v75 = F_contain_volatile_functions(m, v22)
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return int32(0)
																} else {
																	if v75 != 0 {
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																	}
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
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
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
								if v31 == int32(0) {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
									if v53 != 0 {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
										if v54 == int32(0) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
											if v57 != int32(17) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
												if v60 == int32(0) {
													F_check_memoizable(m, v22)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v22
													}
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
													if v63 != int32(2) {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
														v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
														v69 = F_exprType(m, v68)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = F_op_hashjoinable(m, v66, v69)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																if v71 == int32(0) {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v75 = F_contain_volatile_functions(m, v22)
																	mBase = m.M
																	v76 = m.ExcPending
																	if v76 != 0 {
																		return int32(0)
																	} else {
																		if v75 != 0 {
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																		}
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
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
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
									if v34 != int32(2) {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
										if v53 != 0 {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
											if v54 == int32(0) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
												if v57 != int32(17) {
													F_check_memoizable(m, v22)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v22
													}
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
													if v60 == int32(0) {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
														if v63 != int32(2) {
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
															}
														} else {
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
															v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
															v69 = F_exprType(m, v68)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v71 = F_op_hashjoinable(m, v66, v69)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return int32(0)
																} else {
																	if v71 == int32(0) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v75 = F_contain_volatile_functions(m, v22)
																		mBase = m.M
																		v76 = m.ExcPending
																		if v76 != 0 {
																			return int32(0)
																		} else {
																			if v75 != 0 {
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																			}
																			F_check_memoizable(m, v22)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return int32(0)
																			} else {
																				return v22
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
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
										v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
										v40 = F_exprType(m, v39)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = F_op_mergejoinable(m, v37, v40)
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return int32(0)
											} else {
												if v42 == int32(0) {
													v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
													if v53 != 0 {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
														if v54 == int32(0) {
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
															}
														} else {
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
															if v57 != int32(17) {
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
																}
															} else {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
																if v60 == int32(0) {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																	if v63 != int32(2) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
																		v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
																		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
																		v69 = F_exprType(m, v68)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return int32(0)
																		} else {
																			v71 = F_op_hashjoinable(m, v66, v69)
																			mBase = m.M
																			v72 = m.ExcPending
																			if v72 != 0 {
																				return int32(0)
																			} else {
																				if v71 == int32(0) {
																					F_check_memoizable(m, v22)
																					mBase = m.M
																					v81 = m.ExcPending
																					if v81 != 0 {
																						return int32(0)
																					} else {
																						return v22
																					}
																				} else {
																					v75 = F_contain_volatile_functions(m, v22)
																					mBase = m.M
																					v76 = m.ExcPending
																					if v76 != 0 {
																						return int32(0)
																					} else {
																						if v75 != 0 {
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																						}
																						F_check_memoizable(m, v22)
																						mBase = m.M
																						v81 = m.ExcPending
																						if v81 != 0 {
																							return int32(0)
																						} else {
																							return v22
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
													v46 = F_contain_volatile_functions(m, v22)
													mBase = m.M
													v47 = m.ExcPending
													if v47 != 0 {
														return int32(0)
													} else {
														if v46 != 0 {
															v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
															if v53 != 0 {
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
																}
															} else {
																v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
																if v54 == int32(0) {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
																	if v57 != int32(17) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
																		if v60 == int32(0) {
																			F_check_memoizable(m, v22)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return int32(0)
																			} else {
																				return v22
																			}
																		} else {
																			v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																			if v63 != int32(2) {
																				F_check_memoizable(m, v22)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return int32(0)
																				} else {
																					return v22
																				}
																			} else {
																				v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
																				v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
																				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
																				v69 = F_exprType(m, v68)
																				mBase = m.M
																				v70 = m.ExcPending
																				if v70 != 0 {
																					return int32(0)
																				} else {
																					v71 = F_op_hashjoinable(m, v66, v69)
																					mBase = m.M
																					v72 = m.ExcPending
																					if v72 != 0 {
																						return int32(0)
																					} else {
																						if v71 == int32(0) {
																							F_check_memoizable(m, v22)
																							mBase = m.M
																							v81 = m.ExcPending
																							if v81 != 0 {
																								return int32(0)
																							} else {
																								return v22
																							}
																						} else {
																							v75 = F_contain_volatile_functions(m, v22)
																							mBase = m.M
																							v76 = m.ExcPending
																							if v76 != 0 {
																								return int32(0)
																							} else {
																								if v75 != 0 {
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																								}
																								F_check_memoizable(m, v22)
																								mBase = m.M
																								v81 = m.ExcPending
																								if v81 != 0 {
																									return int32(0)
																								} else {
																									return v22
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
															v48 = F_get_mergejoin_opfamilies(m, v37)
															mBase = m.M
															v49 = m.ExcPending
															if v49 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v48
																v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
																if v53 != 0 {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
																	if v54 == int32(0) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
																		if v57 != int32(17) {
																			F_check_memoizable(m, v22)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return int32(0)
																			} else {
																				return v22
																			}
																		} else {
																			v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
																			if v60 == int32(0) {
																				F_check_memoizable(m, v22)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return int32(0)
																				} else {
																					return v22
																				}
																			} else {
																				v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																				if v63 != int32(2) {
																					F_check_memoizable(m, v22)
																					mBase = m.M
																					v81 = m.ExcPending
																					if v81 != 0 {
																						return int32(0)
																					} else {
																						return v22
																					}
																				} else {
																					v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
																					v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
																					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
																					v69 = F_exprType(m, v68)
																					mBase = m.M
																					v70 = m.ExcPending
																					if v70 != 0 {
																						return int32(0)
																					} else {
																						v71 = F_op_hashjoinable(m, v66, v69)
																						mBase = m.M
																						v72 = m.ExcPending
																						if v72 != 0 {
																							return int32(0)
																						} else {
																							if v71 == int32(0) {
																								F_check_memoizable(m, v22)
																								mBase = m.M
																								v81 = m.ExcPending
																								if v81 != 0 {
																									return int32(0)
																								} else {
																									return v22
																								}
																							} else {
																								v75 = F_contain_volatile_functions(m, v22)
																								mBase = m.M
																								v76 = m.ExcPending
																								if v76 != 0 {
																									return int32(0)
																								} else {
																									if v75 != 0 {
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																									}
																									F_check_memoizable(m, v22)
																									mBase = m.M
																									v81 = m.ExcPending
																									if v81 != 0 {
																										return int32(0)
																									} else {
																										return v22
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
func F_build_joinrel_restrictlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
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
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v13 = F_bms_union(m, v11, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	if v17 == int32(0) {
		v222 = v6
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	if v223 == int32(0) {
		v429 = v222
		goto L60
	} else {
		goto L61
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		v222 = v6
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = v6
	v32 = v6
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v30<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = int32(0)
	if v38 == v40 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v222 = v208
	goto L3
L8:
	;
	v210 = v30 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v210 < v211 {
		v30 = v210
		v32 = v208
		goto L6
	} else {
		goto L59
	}
L9:
	;
	if v93 == int32(0) {
		v208 = v32
		goto L8
	} else {
		goto L23
	}
L10:
	;
	v93 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if v39 == int32(0) {
		v86 = v40
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v93 = v86
	goto L9
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v50 < v49 {
		v86 = v40
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(1)
	if v49 <= v52 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = v52
	goto L18
L17:
	;
	v55 = v49
	goto L18
L18:
	;
	v56 = int32(8)
	v61 = int32(0)
	goto L19
L19:
	;
	v68 = v61 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v38+v56+v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v39+v56+v68)))
	v75 = v70 & (v72 ^ int32(-1))
	v77 = base.B2i32(v75 == int32(0))
	if v75 != 0 {
		v86 = v77
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v86 = v77
	goto L13
L21:
	;
	v79 = v61 + int32(1)
	if v79 != v55 {
		v61 = v79
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+11)))
	if v96 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v206 = F_list_append_unique_ptr(m, v32, v37)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L58
	}
L25:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)))
	if v99 != int32(1) {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v103 = int32(0)
	if v102 == v103 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L27
L29:
	;
	if v156 == int32(0) {
		v208 = v32
		goto L8
	} else {
		goto L43
	}
L30:
	;
	v156 = int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v13 == int32(0) {
		v149 = v103
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v156 = v149
	goto L29
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v113 < v112 {
		v149 = v103
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v115 = int32(1)
	if v112 <= v115 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = v115
	goto L38
L37:
	;
	v118 = v112
	goto L38
L38:
	;
	v119 = int32(8)
	v124 = int32(0)
	goto L39
L39:
	;
	v131 = v124 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v102+v119+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v13+v119+v131)))
	v138 = v133 & (v135 ^ int32(-1))
	v140 = base.B2i32(v138 == int32(0))
	if v138 != 0 {
		v149 = v140
		goto L33
	} else {
		goto L41
	}
L40:
	;
	v149 = v140
	goto L33
L41:
	;
	v142 = v124 + int32(1)
	if v142 != v118 {
		v124 = v142
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	v160 = int32(0)
	if base.B2i32(v159 == v160)|base.B2i32(v13 == v160) != 0 {
		v205 = v160
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v205 != 0 {
		v208 = v32
		goto L8
	} else {
		goto L57
	}
L45:
	;
	goto L44
L46:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v170 < v171 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v173 = v170
	goto L49
L48:
	;
	v173 = v171
	goto L49
L49:
	;
	if v173 <= int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v176 = int32(1)
	goto L52
L51:
	;
	v176 = v173
	goto L52
L52:
	;
	v177 = int32(8)
	v182 = int32(0)
	goto L53
L53:
	;
	v189 = v182 << (uint(int32(2)) % 32)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13+v177+v189)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v159+v177+v189)))
	v194 = v191 & v193
	v196 = base.B2i32(v194 != int32(0))
	if v194 != 0 {
		v205 = v196
		goto L45
	} else {
		goto L55
	}
L54:
	;
	v205 = v196
	goto L45
L55:
	;
	v198 = v182 + int32(1)
	if v198 != v176 {
		v182 = v198
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L24
L58:
	;
	v208 = v206
	goto L8
L59:
	;
	goto L7
L60:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v432 = F_generate_join_implied_equalities(m, l0, v430, v431, l3, l4)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L117
	}
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v226 <= int32(0) {
		v429 = v222
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v237 = int32(0)
	v239 = v222
	goto L63
L63:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v237<<(uint(int32(2))%32))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v247 = int32(0)
	if v245 == v247 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v429 = v415
	goto L60
L65:
	;
	v417 = v237 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v417 < v418 {
		v237 = v417
		v239 = v415
		goto L63
	} else {
		goto L116
	}
L66:
	;
	if v300 == int32(0) {
		v415 = v239
		goto L65
	} else {
		goto L80
	}
L67:
	;
	v300 = int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if v246 == int32(0) {
		v293 = v247
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v300 = v293
	goto L66
L71:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v257 < v256 {
		v293 = v247
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v259 = int32(1)
	if v256 <= v259 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v262 = v259
	goto L75
L74:
	;
	v262 = v256
	goto L75
L75:
	;
	v263 = int32(8)
	v268 = int32(0)
	goto L76
L76:
	;
	v275 = v268 << (uint(int32(2)) % 32)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v245+v263+v275)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v246+v263+v275)))
	v282 = v277 & (v279 ^ int32(-1))
	v284 = base.B2i32(v282 == int32(0))
	if v282 != 0 {
		v293 = v284
		goto L70
	} else {
		goto L78
	}
L77:
	;
	v293 = v284
	goto L70
L78:
	;
	v286 = v268 + int32(1)
	if v286 != v262 {
		v268 = v286
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+11)))
	if v303 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v413 = F_list_append_unique_ptr(m, v239, v244)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L115
	}
L82:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+12)))
	if v306 != int32(1) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v310 = int32(0)
	if v309 == v310 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	if v363 == int32(0) {
		v415 = v239
		goto L65
	} else {
		goto L100
	}
L87:
	;
	v363 = int32(1)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v13 == int32(0) {
		v356 = v310
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v363 = v356
	goto L86
L91:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v320 < v319 {
		v356 = v310
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v322 = int32(1)
	if v319 <= v322 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v325 = v322
	goto L95
L94:
	;
	v325 = v319
	goto L95
L95:
	;
	v326 = int32(8)
	v331 = int32(0)
	goto L96
L96:
	;
	v338 = v331 << (uint(int32(2)) % 32)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v309+v326+v338)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v13+v326+v338)))
	v345 = v340 & (v342 ^ int32(-1))
	v347 = base.B2i32(v345 == int32(0))
	if v345 != 0 {
		v356 = v347
		goto L90
	} else {
		goto L98
	}
L97:
	;
	v356 = v347
	goto L90
L98:
	;
	v349 = v331 + int32(1)
	if v349 != v325 {
		v331 = v349
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v244)+36))
	v367 = int32(0)
	if base.B2i32(v366 == v367)|base.B2i32(v13 == v367) != 0 {
		v412 = v367
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v412 != 0 {
		v415 = v239
		goto L65
	} else {
		goto L114
	}
L102:
	;
	goto L101
L103:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v377 < v378 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v380 = v377
	goto L106
L105:
	;
	v380 = v378
	goto L106
L106:
	;
	if v380 <= int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v383 = int32(1)
	goto L109
L108:
	;
	v383 = v380
	goto L109
L109:
	;
	v384 = int32(8)
	v389 = int32(0)
	goto L110
L110:
	;
	v396 = v389 << (uint(int32(2)) % 32)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v13+v384+v396)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v366+v384+v396)))
	v401 = v398 & v400
	v403 = base.B2i32(v401 != int32(0))
	if v401 != 0 {
		v412 = v403
		goto L102
	} else {
		goto L112
	}
L111:
	;
	v412 = v403
	goto L102
L112:
	;
	v405 = v389 + int32(1)
	if v405 != v383 {
		v389 = v405
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	goto L81
L115:
	;
	v415 = v413
	goto L65
L116:
	;
	goto L64
L117:
	;
	v434 = F_list_concat(m, v429, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	return v434
}
func F_build_pertrans_for_aggref(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int64, l9 int32, l10 int32, l11 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	v10 = l9
	v13 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v13
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v13)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l3
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v29
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = v38
	goto L3
L2:
	;
	v39 = v13
	goto L3
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = v41
	goto L6
L5:
	;
	v42 = v13
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+49)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v48 = int32(0)
	F_build_aggregate_transfn_expr(m, l10, l11, v39, v46, l5, v47, l4, v48, v19+int32(12), v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v55 = l0 + int32(32)
	F_fmgr_info(m, l4, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58
	v61 = v45 + int32(1)
	v66 = F_palloc(m, v61<<(uint(int32(4))%32)+int32(24))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v55
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l1
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+16)) = uint8(v73)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+18)) = uint16(v61)
	F_get_typlenbyval(m, l5, l0+int32(188), l0+int32(191))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	if l6 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v91 = m.G0
	v93 = v91 - int32(16)
	m.G0 = v93
	v96 = F_palloc0(m, int32(28))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if l7 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = int64(9801115369471)
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v96
	v112 = F_list_make1_impl(m, int32(1), v93+int32(8))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v114 = int32(0)
	v116 = F_makeFuncExpr(m, l6, int32(17), v112, v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(8)))) = v116
	m.G0 = v93 + int32(16)
	v123 = l0 + int32(60)
	F_fmgr_info(m, l6, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v126
	v129 = F_palloc(m, int32(40))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+228)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v123
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = l1
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v136
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+16)) = uint8(v136)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v145 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+18)) = uint16(v145)
	goto L14
L20:
	;
	v151 = m.G0
	v153 = v151 - int32(16)
	m.G0 = v153
	v156 = F_palloc0(m, int32(28))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+50)))
	if v222 == int32(110) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = int64(77309411327)
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v156
	v168 = F_palloc0(m, int32(28))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v168)+8)) = int64(9801115369471)
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v168
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v180
	v185 = F_list_make2_impl(m, v153+int32(4), v153)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v187 = int32(0)
	v189 = F_makeFuncExpr(m, l7, int32(2281), v185, v187, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)))) = v189
	m.G0 = v153 + int32(16)
	v196 = l0 + int32(88)
	F_fmgr_info(m, l7, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v199
	v202 = F_palloc(m, int32(56))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+232)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v196
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = l1
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v209 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v209
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	*(*uint8)(unsafe.Add(mBase, uint32(v214)+16)) = uint8(v209)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v218 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+18)) = uint16(v218)
	goto L22
L29:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v390 != 0 {
		goto L68
	} else {
		goto L69
	}
L30:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v276 = F_ExecTypeFromTL(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L47
	}
L31:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	if v265 == int32(0) {
		v385 = v264
		goto L29
	} else {
		goto L46
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v250
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v253)
	if v250 <= int32(0) {
		v262 = v250
		v263 = v251
		v264 = v252
		goto L31
	} else {
		goto L45
	}
L33:
	;
	v243 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v245 != 0 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	v250 = v240
	v251 = v225
	v252 = v240
	v253 = v226 ^ int32(1)
	goto L32
L35:
	;
	if v225 == int32(0) {
		goto L33
	} else {
		goto L41
	}
L36:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+51)))
	if v226 != int32(1) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v233)
	v262 = v233
	v263 = v233
	v264 = v233
	goto L31
L39:
	;
	if v225 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L34
L42:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v247 = v246
	goto L44
L43:
	;
	v247 = v243
	goto L44
L44:
	;
	v250 = v247
	v251 = v245
	v252 = v243
	v253 = base.B2i32(int32(0) < v247)
	goto L32
L45:
	;
	v271 = v250
	v272 = v251
	v273 = v252
	v274 = int32(1)
	goto L30
L46:
	;
	v271 = v262
	v272 = v263
	v273 = v264
	v274 = int32(0)
	goto L30
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v276
	v280 = F_ExecInitExtraTupleSlot(m, l2, v276, int32(_a_F_build_pertrans_for_aggref_0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v280
	if v274 == int32(0) {
		v385 = v273
		goto L29
	} else {
		goto L49
	}
L49:
	;
	if v42 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v306 = F_palloc(m, v271<<(uint(int32(1))%32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L57
	}
L51:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l10+v39<<(uint(int32(2))%32))))
	F_get_typlenbyval(m, v290, l0+int32(186), l0+int32(190))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v273 <= int32(0) {
		goto L50
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v301 = F_ExecInitExtraTupleSlot(m, l2, v299, int32(_a_F_build_pertrans_for_aggref_0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v301
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v306
	v310 = v271 << (uint(int32(2)) % 32)
	v311 = F_palloc(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v311
	v314 = F_palloc(m, v310)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v314
	v317 = F_palloc(m, v271)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v317
	if v272 == int32(0) {
		v385 = v273
		goto L29
	} else {
		goto L61
	}
L61:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v322 <= int32(0) {
		v385 = v273
		goto L29
	} else {
		goto L62
	}
L62:
	;
	v333 = int32(0)
	goto L63
L63:
	;
	v343 = v333 << (uint(int32(2)) % 32)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343+v344)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v348 = F_get_sortgroupclause_tle(m, v346, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	v385 = v273
	goto L29
L65:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v350+v333<<(uint(int32(1))%32)))) = uint16(v354)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v343))) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	v361 = F_exprCollation(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v363+v343))) = v361
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v366+v333))) = uint8(v368)
	v371 = v333 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v371 < v372 {
		v333 = v371
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v393 = F_palloc(m, v385<<(uint(int32(2))%32))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v480 = int32(1)
	if v21 <= v480 {
		goto L86
	} else {
		goto L87
	}
L71:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v395 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v385 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v398 <= int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v409 = int32(0)
	goto L75
L75:
	;
	v419 = v409 << (uint(int32(2)) % 32)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421+v419)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v393+v419))) = v424
	v427 = v409 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v427 < v428 {
		v409 = v427
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L72
L77:
	;
	goto L76
L78:
	;
	F_pfree(m, v393)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L7
	} else {
		goto L85
	}
L79:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v449 = F_get_opcode(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L7
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v458 = F_execTuplesMatchPrepare(m, v455, v385, v456, v393, v457, l1)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	F_fmgr_info(m, v449, l0+int32(144))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v458
	goto L78
L85:
	;
	goto L70
L86:
	;
	v483 = v480
	goto L88
L87:
	;
	v483 = v21
	goto L88
L88:
	;
	v484 = F_palloc0_mul(m, int32(4), v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v484
	m.G0 = v19 + int32(16)
	return
}
func F_byteage(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v160 = int32(0)
	return base.I64_extend_i32_u(base.B2i32(v151 == v160)&base.B2i32(v75 <= v45) | base.B2i32(v160 < v151))
L60:
	;
	goto L59
}
func F_byteaoverlay(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v12 = F_bytea_overlay(m, v3, v8, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
func F_byteasend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
