package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MJEvalOuterValues(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(2)
L2:
	;
	goto L3
L3:
	;
	v13 = int32(2)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v14&v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v97 = v13
	goto L6
L6:
	;
	return v97
L7:
	;
	return int32(0)
L8:
	;
	v25 = int32(_a_F_MJEvalOuterValues_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_MJEvalOuterValues[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MJEvalOuterValues[0])) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v30
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v33 <= v32 {
		v88 = v32
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MJEvalOuterValues[0])) = v26
	v97 = v88
	goto L6
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v41 = m.T0[v40].(func(*base.Module, int32, int32, int32) int32)(m, v37, v19, v36+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v41
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+16)))
	if v45 != int32(1) {
		v54 = int32(0)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v55 < int32(2) {
		v88 = v54
		goto L9
	} else {
		goto L17
	}
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+29)))
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v54 = int32(1)
	goto L12
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+131)))
	if v49 == int32(1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v54 = int32(2)
	goto L12
L17:
	;
	v60 = v54
	v62 = int32(1)
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v69 = v66 + v62*int32(56)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v74 = m.T0[v73].(func(*base.Module, int32, int32, int32) int32)(m, v70, v19, v69+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L20
	}
L19:
	;
	v88 = v82
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v74
	v77 = int32(1)
	if base.Ui32(v60) <= base.Ui32(v77) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = v77
	goto L23
L22:
	;
	v80 = v60
	goto L23
L23:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+16)))
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = v80
	goto L26
L25:
	;
	v82 = v60
	goto L26
L26:
	;
	v84 = v62 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v84 < v85 {
		v60 = v82
		v62 = v84
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
}
func F_MakeTransitionCaptureState(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
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
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l0 == v4 {
		v504 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L94
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v504
L3:
	;
	switch l2 - int32(2) {
	case 0:
		goto L10
	case 1:
		goto L6
	case 2:
		goto L9
	case 3:
		goto L8
	default:
		goto L7
	}
L4:
	;
	v53 = v52 | v51
	v54 = int32(1)
	if v53&v54|v49&v54|v50&v54 == int32(0) {
		v504 = v4
		goto L2
	} else {
		goto L15
	}
L5:
	;
	v49 = v47
	v50 = v46
	v51 = v4
	v52 = int32(0)
	goto L4
L6:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	v46 = v4
	v47 = v45
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v49 = v26
	v50 = v27
	v51 = v28
	v52 = v29
	goto L4
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	v46 = v24
	v47 = int32(0)
	goto L5
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v49 = v4
	v50 = v4
	v51 = v22
	v52 = v23
	goto L4
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	F_errmsg_internal(m, int32(_a_F_MakeTransitionCaptureState_0), v16)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_MakeTransitionCaptureState_1), int32(_a_F_MakeTransitionCaptureState_2), int32(_a_F_MakeTransitionCaptureState_3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[0]))
	if v65 < int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[1]))
	if v65 < v69 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v145 = int32(0)
	if v49&int32(1) == v145 {
		v227 = v145
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v72 = v65 + int32(1)
	if v69 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[1])) = v95
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[2])) = v97
	if v95 <= v69 {
		goto L17
	} else {
		goto L31
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[3]))
	v77 = int32(8)
	if v72 <= v77 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[2]))
	v88 = v69 << (uint(int32(1)) % 32)
	if v88 < v72 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v80 = v77
	goto L25
L24:
	;
	v80 = v72
	goto L25
L25:
	;
	v83 = F_MemoryContextAlloc(m, v76, v80*int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v95 = v80
	v97 = v83
	goto L19
L27:
	;
	v90 = v72
	goto L29
L28:
	;
	v90 = v88
	goto L29
L29:
	;
	v93 = F_repalloc(m, v86, v90*int32(20))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v95 = v90
	v97 = v93
	goto L19
L31:
	;
	v106 = v69
	goto L32
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[2]))
	v120 = v117 + v106*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = int32(0)
	v123 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v120)+8)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v120))) = v123
	v128 = v106 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[1]))
	if v128 < v130 {
		v106 = v128
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L17
L34:
	;
	goto L33
L35:
	;
	if v53&int32(1) == int32(0) {
		v316 = v145
		goto L49
	} else {
		goto L50
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[2]))
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[0]))
	v157 = v152 + v154*int32(20)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	if v158 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v207 = int32(_a_F_MakeTransitionCaptureState_4)
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4]))
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v211
	v214 = F_palloc0(m, int32(36))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L47
	}
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v161 <= int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v169 = int32(0)
	goto L40
L40:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v164+v169<<(uint(int32(2))%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v183 != l1 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L37
L42:
	;
	v192 = v169 + int32(1)
	if v161 != v192 {
		v169 = v192
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v185 != int32(3) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
	if v188 != int32(1) {
		v227 = v182
		goto L35
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = l1
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v220 = F_lappend(m, v219, v214)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = v220
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v208
	v227 = v214
	goto L35
L49:
	;
	v329 = int32(0)
	if v50&int32(1) == v329 {
		v412 = v329
		goto L63
	} else {
		goto L64
	}
L50:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[2]))
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[0]))
	v248 = v243 + v245*int32(20)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	if v249 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v298 = int32(_a_F_MakeTransitionCaptureState_4)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4]))
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v302
	v305 = F_palloc0(m, int32(36))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L61
	}
L52:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v252 <= int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v260 = int32(0)
	goto L54
L54:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v255+v260<<(uint(int32(2))%32))))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v274 != l1 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L51
L56:
	;
	v283 = v260 + int32(1)
	if v252 != v283 {
		v260 = v283
		goto L54
	} else {
		goto L60
	}
L57:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v276 != int32(2) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+8)))
	if v279 != int32(1) {
		v316 = v273
		goto L49
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	goto L55
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = l1
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	v311 = F_lappend(m, v310, v305)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v311
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v299
	v316 = v305
	goto L49
L63:
	;
	v421 = int32(_a_F_MakeTransitionCaptureState_4)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4]))
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v425
	v427 = int32(_a_F_MakeTransitionCaptureState_5)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[6]))
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[6])) = v431
	if v52&int32(1) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[2]))
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[0]))
	v340 = v335 + v337*int32(20)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+16))
	if v341 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v390 = int32(_a_F_MakeTransitionCaptureState_4)
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4]))
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v394
	v397 = F_palloc0(m, int32(36))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L11
	} else {
		goto L75
	}
L66:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v344 <= int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v352 = int32(0)
	goto L68
L68:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v347+v352<<(uint(int32(2))%32))))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if v366 != l1 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L65
L70:
	;
	v375 = v352 + int32(1)
	if v344 != v375 {
		v352 = v375
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v368 != int32(4) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+8)))
	if v371 != int32(1) {
		v412 = v365
		goto L63
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	goto L69
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+4)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = l1
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v340)+16))
	v403 = F_lappend(m, v402, v397)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+16)) = v403
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v391
	v412 = v397
	goto L63
L77:
	;
	if v51&int32(1) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v316)+24))
	if v437 != 0 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v438 = int32(0)
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[8]))
	v442 = F_tuplestore_begin_heap(m, v438, v438, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+24)) = v442
	goto L77
L81:
	;
	if v50&int32(1) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v316)+28))
	if v449 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v450 = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[8]))
	v454 = F_tuplestore_begin_heap(m, v450, v450, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+28)) = v454
	goto L81
L85:
	;
	v470 = v49 & int32(1)
	if v470 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v412)+24))
	if v461 != 0 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v462 = int32(0)
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[8]))
	v466 = F_tuplestore_begin_heap(m, v462, v462, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v412)+24)) = v466
	goto L85
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[4])) = v422
	*(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[6])) = v428
	v486 = F_palloc0(m, int32(20))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L11
	} else {
		goto L93
	}
L90:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	if v473 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v474 = int32(0)
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTransitionCaptureState[8]))
	v478 = F_tuplestore_begin_heap(m, v474, v474, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+28)) = v478
	goto L89
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v486)+16)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v486)+12)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v486)+8)) = v227
	*(*uint8)(unsafe.Add(mBase, uint32(v486)+3)) = uint8(v470)
	v492 = int32(1)
	v493 = v51 & v492
	*(*uint8)(unsafe.Add(mBase, uint32(v486)+2)) = uint8(v493)
	v496 = v52 & v492
	*(*uint8)(unsafe.Add(mBase, uint32(v486)+1)) = uint8(v496)
	v499 = v50 & v492
	*(*uint8)(unsafe.Add(mBase, uint32(v486))) = uint8(v499)
	v504 = v486
	goto L2
L94:
	;
	F_errmsg_internal(m, int32(_a_F_MakeTransitionCaptureState_6), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_MakeTransitionCaptureState_1), int32(_a_F_MakeTransitionCaptureState_7), int32(_a_F_MakeTransitionCaptureState_3))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_MetaphAdd(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	if l1 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = F_strlen(m, l1)
		mBase = m.M
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v11 <= v8+v9 {
			v14 = v9 + int32(10)
			v16 = F_repalloc(m, v7, v11+v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19 + v14
				v22 = v16
				v24 = F_strlen(m, v22)
				mBase = m.M
				v26 = F_strcpy(m, v24+v22, l1)
				mBase = m.M
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27 + v9
				return
			}
		} else {
			v22 = v7
			v24 = F_strlen(m, v22)
			mBase = m.M
			v26 = F_strcpy(m, v24+v22, l1)
			mBase = m.M
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27 + v9
			return
		}
	} else {
		return
	}
}
func F__mdfd_getseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	v19 = int32(base.Ui32(l2) >> (uint(int32(17)) % 32))
	v22 = l0 + l1<<(uint(int32(2))%32)
	v24 = v22 + int32(40)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if base.Ui32(v19) < base.Ui32(v25) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(128)
	return v216
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v216 = v27 + v19<<(uint(int32(3))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v31 = int32(0)
	if base.Ui32(int32(31)) < base.Ui32(l4) {
		v216 = v31
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if base.Ui32(v19) < base.Ui32(v50) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56))
	v50 = v25
	v51 = v37 + v25<<(uint(int32(3))%32) - int32(8)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v43 = F_mdopenfork(m, l0, l1, l4)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v43 == int32(0) {
		v216 = v31
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v50 = v49
	v51 = v43
	goto L6
L13:
	;
	v216 = v51
	goto L1
L14:
	;
	goto L15
L15:
	;
	v64 = v51
	v66 = v50
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v73 = F_FileSize(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L20
	}
L17:
	;
	v216 = v174
	goto L1
L18:
	;
	v174 = F__mdfd_openseg(m, l0, l1, v66, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L51
	}
L19:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L48
	}
L20:
	;
	if int64(0) <= v73 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v79 = base.I32_wrap_i64(int64(base.Ui64(v73) >> (uint(int64(13)) % 64)))
	if base.Ui32(int32(_a_F__mdfd_getseg_0)) <= base.Ui32(v79) {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L43
	}
L24:
	;
	if l4&int32(4) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v79 == int32(_a_F__mdfd_getseg_1) {
		v173 = int32(0)
		goto L18
	} else {
		goto L34
	}
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__mdfd_getseg[0])))
	if base.B2i32(l4&int32(8) != int32(0))&v85 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v89 = int32(64)
	if v79 == int32(_a_F__mdfd_getseg_1) {
		v173 = v89
		goto L18
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v99 = F_palloc_aligned(m, int32(_a_F__mdfd_getseg_2), int32(_a_F__mdfd_getseg_3), int32(4))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	F_mdextend(m, l0, l1, v66<<(uint(int32(17))%32)-int32(1), v99, l3)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	F_pfree(m, v99)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v173 = v89
	goto L18
L34:
	;
	if l4&int32(2) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__mdfd_getseg[1])) = int32(44)
	v216 = int32(0)
	goto L1
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v121 = v16 + int32(45)
	F__mdfd_segpath(m, v121, l0, l1, v66)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v121
	F_errmsg(m, int32(_a_F__mdfd_getseg_4), v16+int32(32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F__mdfd_getseg_5), int32(1848), int32(_a_F__mdfd_getseg_6))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v145 = *(*int32)(unsafe.Add(mBase, _c_F__mdfd_getseg[2]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v143*int32(48))+32))
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v149
	F_errmsg(m, int32(_a_F__mdfd_getseg_7), v16)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F__mdfd_getseg_5), int32(1882), int32(_a_F__mdfd_getseg_8))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(_a_F__mdfd_getseg_9), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F__mdfd_getseg_5), int32(1794), int32(_a_F__mdfd_getseg_6))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v174 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if l4&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v66 != v19 {
		v64 = v174
		v66 = v66 + int32(1)
		goto L16
	} else {
		goto L64
	}
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F__mdfd_getseg[1]))
	if v182 == int32(44) {
		v216 = int32(0)
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v193 = v16 + int32(45)
	F__mdfd_segpath(m, v193, l0, l1, v66)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v193
	F_errmsg(m, int32(_a_F__mdfd_getseg_10), v16+int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F__mdfd_getseg_5), int32(1862), int32(_a_F__mdfd_getseg_6))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	goto L17
}
func F_mXactCachePut(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[0]))
	if v7 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[1]))
		v17 = F_AllocSetContextCreateInternal(m, v12, int32(_a_F_mXactCachePut_0), int32(0), int32(1024), int32(_a_F_mXactCachePut_1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[0])) = v17
			v20 = v17
			v22 = l1 << (uint(int32(3)) % 32)
			v25 = F_MemoryContextAlloc(m, v20, v22+int32(16))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
				v30 = v25 + int32(16)
				if v22 != 0 {
					base.MemoryCopy(m, v30, l2, v22)
				} else {
				}
				F_pg_qsort(m, v30, l1, int32(8), int32(291))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[2]))
					if v37 == int32(0) {
						v40 = int32(_a_F_mXactCachePut_2)
						*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[3])) = v40
						*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4])) = int32(0)
						v47 = v40
					} else {
						v47 = v37
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(_a_F_mXactCachePut_2)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v47
					v52 = v25 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = v52
					*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[2])) = v52
					v56 = int32(_a_F_mXactCachePut_3)
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4]))
					v60 = v58 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4])) = v60
					if base.Ui32(int32(257)) <= base.Ui32(v60) {
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[3]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v67
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = v69
						v71 = int32(_a_F_mXactCachePut_3)
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4])) = v73 - int32(1)
						F_pfree(m, v65-int32(8))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			}
		}
	} else {
		v20 = v7
		v22 = l1 << (uint(int32(3)) % 32)
		v25 = F_MemoryContextAlloc(m, v20, v22+int32(16))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
			v30 = v25 + int32(16)
			if v22 != 0 {
				base.MemoryCopy(m, v30, l2, v22)
			} else {
			}
			F_pg_qsort(m, v30, l1, int32(8), int32(291))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[2]))
				if v37 == int32(0) {
					v40 = int32(_a_F_mXactCachePut_2)
					*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[3])) = v40
					*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4])) = int32(0)
					v47 = v40
				} else {
					v47 = v37
				}
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(_a_F_mXactCachePut_2)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v47
				v52 = v25 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = v52
				*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[2])) = v52
				v56 = int32(_a_F_mXactCachePut_3)
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4]))
				v60 = v58 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4])) = v60
				if base.Ui32(int32(257)) <= base.Ui32(v60) {
					v65 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[3]))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v67
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					*(*int32)(unsafe.Add(mBase, uint32(v67))) = v69
					v71 = int32(_a_F_mXactCachePut_3)
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4]))
					*(*int32)(unsafe.Add(mBase, _c_F_mXactCachePut[4])) = v73 - int32(1)
					F_pfree(m, v65-int32(8))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	}
}
func F_makeAlias(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v5 = F_palloc0(m, int32(12))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(2)
		v11 = F_pstrdup(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v11
			return v5
		}
	}
}
func F_makeInteger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(465)
		return v4
	}
}
func F_make_colname_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = F_colname_is_unique(m, l0, l1, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = F_strlen(m, l0)
	mBase = m.M
	v21 = F_palloc(m, v18+int32(16))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v74 = l0
	goto L5
L5:
	;
	m.G0 = v10 + int32(32)
	return v74
L6:
	;
	v26 = v18
	v29 = int32(0)
	goto L7
L7:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v74 = v21
	goto L5
L9:
	;
	base.MemoryCopy(m, v21, l0, v26)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v32 = v29 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v32
	v38 = F_pg_sprintf(m, v26+v21, int32(_a_F_make_colname_unique_0), v10+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v40 = F_strlen(m, v21)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v40) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v26
	goto L16
L14:
	;
	v66 = v26
	goto L15
L15:
	;
	v70 = F_colname_is_unique(m, v21, l1, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L16:
	;
	v52 = F_pg_mbcliplen(m, l0, v46, v46-int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v66 = v52
	goto L15
L18:
	;
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	base.MemoryCopy(m, v21, l0, v52)
	goto L21
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
	v58 = F_pg_sprintf(m, v52+v21, int32(_a_F_make_colname_unique_0), v10)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v60 = F_strlen(m, v21)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v60) {
		v46 = v52
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	if v70 == int32(0) {
		v26 = v66
		v29 = v32
		goto L7
	} else {
		goto L25
	}
L25:
	;
	goto L8
}
func F_makepol_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
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
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
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
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = v2
	v26 = v2
	v27 = v2
	v30 = v2
	goto L4
L3:
	;
	m.G0 = v15 + int32(128)
	return v551
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = v33
	v39 = v26
	v40 = v27
	v43 = v30
	goto L6
L5:
	;
	v519 = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v520 == v519 {
		v530 = v519
		goto L125
	} else {
		goto L126
	}
L6:
	;
	v46 = F_pg_mblen_cstr(m, v36)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	goto L7
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v48 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L12
	case 2:
		goto L18
	default:
		goto L8
	}
L10:
	;
	v515 = v510 + v46
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v515
	v36 = v515
	v39 = v512
	v40 = v513
	v43 = v514
	goto L6
L11:
	;
	v481 = F_t_isalnum_cstr(m, v51)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L114
	}
L12:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v260 = F_t_isalnum_cstr(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L62
	}
L13:
	;
	if v22 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L46
	}
L15:
	;
	v175 = v22
	goto L42
L16:
	;
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v112 + v162
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v165 - v162
	if int32(0) < v165 {
		goto L13
	} else {
		goto L41
	}
L17:
	;
	if v22 == int32(32) {
		goto L14
	} else {
		goto L40
	}
L18:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v113 == int32(32) {
		v510 = v112
		v512 = v39
		v513 = v40
		v514 = v43
		goto L10
	} else {
		goto L30
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	switch v52 - int32(33) {
	case 0:
		goto L20
	default:
		goto L11
	case 7:
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v51 + int32(1)
	v152 = int32(33)
	goto L17
L21:
	;
	v55 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v51 + v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v59 + v55
	v63 = F_makepol_2(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v63 != 0 {
		v551 = v55
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v65 = int32(0)
	if v22 == v65 {
		v22 = v65
		v26 = v39
		v27 = v40
		v30 = v43
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v72 = v22
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15+v72<<(uint(int32(2))%32)-int32(4))))
	switch v85 - int32(33) {
	case 0, 5:
		goto L27
	default:
		v22 = v72
		v26 = v39
		v27 = v40
		v30 = v43
		goto L4
	}
L26:
	;
	v22 = int32(0)
	v26 = v39
	v27 = v40
	v30 = v43
	goto L4
L27:
	;
	v90 = F_palloc(m, int32(20))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v92 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+12)) = uint16(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v92
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v90
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v102 + v103
	v107 = v72 - v103
	if v107 != 0 {
		v72 = v107
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	switch v113 - int32(38) {
	case 0:
		goto L32
	case 1, 2:
		goto L8
	case 3:
		goto L16
	default:
		goto L33
	}
L31:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v149 != 0 {
		goto L8
	} else {
		goto L38
	}
L32:
	;
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v122
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v112 + v122
	if base.B2i32(v22 == int32(0))|base.B2i32(v124 != int32(124)) != 0 {
		v152 = v124
		goto L17
	} else {
		goto L36
	}
L33:
	;
	if v113 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v113 != int32(124) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v134 = F_palloc(m, int32(20))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+12)) = uint16(v136)
	*(*int64)(unsafe.Add(mBase, uint32(v134))) = int64(532575944707)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v136
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+16)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v134
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v145 + int32(1)
	v26 = v39
	v27 = v40
	v30 = v43
	goto L4
L38:
	;
	if v22 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	v551 = int32(0)
	goto L3
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v22<<(uint(int32(2))%32)))) = v152
	v22 = v22 + int32(1)
	v26 = v39
	v27 = v40
	v30 = v43
	goto L4
L41:
	;
	goto L8
L42:
	;
	v184 = v175 - int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v15+v184<<(uint(int32(2))%32))))
	v191 = F_palloc(m, int32(20))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v551 = int32(0)
	goto L3
L44:
	;
	v193 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+12)) = uint16(v193)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v193
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+16)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v191
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v203 + int32(1)
	if v184 != 0 {
		v175 = v184
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_errmsg_internal(m, int32(_a_F_makepol_2_0), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_makepol_2_1), int32(252), int32(_a_F_makepol_2_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
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
	v551 = int32(0)
	goto L3
L50:
	;
	goto L51
L51:
	;
	v227 = v22
	goto L52
L52:
	;
	v236 = v227 - int32(1)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v15+v236<<(uint(int32(2))%32))))
	v243 = F_palloc(m, int32(20))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v551 = int32(0)
	goto L3
L54:
	;
	v245 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v243)+12)) = uint16(v245)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v245
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v243
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v255 + int32(1)
	if v236 != 0 {
		v227 = v236
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	if v39 <= int32(_a_F_makepol_2_3) {
		goto L72
	} else {
		goto L73
	}
L57:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v510 = v298
	v512 = v46 + v39
	v513 = int32(0)
	v514 = v43
	goto L10
L58:
	;
	v510 = v262
	v512 = v39
	v513 = v40 | int32(1)
	v514 = v43
	goto L10
L59:
	;
	v510 = v262
	v512 = v39
	v513 = v40 | int32(2)
	v514 = v43
	goto L10
L60:
	;
	v510 = v262
	v512 = v39
	v513 = v40 | int32(4)
	v514 = v43
	goto L10
L61:
	;
	if v40&int32(_a_F_makepol_2_3) == int32(0) {
		goto L57
	} else {
		goto L66
	}
L62:
	;
	if v260 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	switch v263 - int32(37) {
	case 0:
		goto L60
	case 1, 2, 3, 4, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26:
		goto L56
	case 5:
		goto L58
	case 8:
		goto L61
	case 27:
		goto L59
	default:
		goto L64
	}
L64:
	;
	if v263 != int32(95) {
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v275 = F_errsave_start(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v275 == int32(0) {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_makepol_2_4), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errsave_finish(m, v274, int32(_a_F_makepol_2_1), int32(102), int32(_a_F_makepol_2_5))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L8
L72:
	;
	v304 = F_ltree_crc32_sz(m, v43, v39)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v463 = int32(1)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v465 = F_errsave_start(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L109
	}
L75:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v309 = F_palloc(m, int32(20))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v309)+12)) = uint16(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v309)+4)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = int32(2)
	v315 = v307 - v306
	if int32(_a_F_makepol_2_6) <= v315 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v318 = int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v320 = F_errsave_start(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if int32(256) <= v39 {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	if v320 == int32(0) {
		v551 = v318
		goto L3
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_makepol_2_7), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errsave_finish(m, v319, int32(_a_F_makepol_2_1), int32(165), int32(_a_F_makepol_2_8))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v551 = v318
	goto L3
L85:
	;
	v338 = int32(1)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v340 = F_errsave_start(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v309)+10)) = uint16(v39)
	*(*uint16)(unsafe.Add(mBase, uint32(v309)+8)) = uint16(v315)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v309)+16)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v309
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v362 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v361 + v362
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v367 = v365 - v366
	v369 = v39 + v362
	v370 = v367 + v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v371 <= v370 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	if v340 == int32(0) {
		v551 = v338
		goto L3
	} else {
		goto L89
	}
L89:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_makepol_2_9), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errsave_finish(m, v339, int32(_a_F_makepol_2_1), int32(169), int32(_a_F_makepol_2_8))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v551 = v338
	goto L3
L93:
	;
	v374 = v371
	v375 = v366
	goto L96
L94:
	;
	v398 = v365
	goto L95
L95:
	;
	if v39 != 0 {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v386 = v374 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v386
	v388 = F_repalloc(m, v375, v386)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	v398 = v391
	goto L95
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v388
	v391 = v388 + v367
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v393 <= v370 {
		v374 = v393
		v375 = v388
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	base.MemoryCopy(m, v398, v43, v39)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v409 = v408 + v39
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v409
	v411 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v411)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v414 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v418 + v369
	if v22 == v411 {
		v22 = v411
		v26 = v39
		v27 = v40
		v30 = v43
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v427 = v22
	goto L104
L104:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v15+v427<<(uint(int32(2))%32)-int32(4))))
	switch v440 - int32(33) {
	case 0, 5:
		goto L106
	default:
		v22 = v427
		v26 = v39
		v27 = v40
		v30 = v43
		goto L4
	}
L105:
	;
	v22 = int32(0)
	v26 = v39
	v27 = v40
	v30 = v43
	goto L4
L106:
	;
	v445 = F_palloc(m, int32(20))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v447 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v445)+12)) = uint16(v447)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v447
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+16)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v445
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v458 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v457 + v458
	v462 = v427 - v458
	if v462 != 0 {
		v427 = v462
		goto L104
	} else {
		goto L108
	}
L108:
	;
	goto L105
L109:
	;
	if v465 == int32(0) {
		v551 = v463
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_makepol_2_10), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errsave_finish(m, v464, int32(_a_F_makepol_2_1), int32(187), int32(_a_F_makepol_2_11))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v551 = v463
	goto L3
L114:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v481 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v494 = F_errsave_start(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L120
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
	v510 = v483
	v512 = v46
	v513 = int32(0)
	v514 = v483
	goto L10
L117:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	switch v484 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v510 = v483
		v512 = v39
		v513 = v40
		v514 = v43
		goto L10
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35:
		goto L115
	case 36:
		goto L116
	default:
		goto L118
	}
L118:
	;
	if v484 != int32(95) {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	goto L116
L120:
	;
	if v494 == int32(0) {
		goto L8
	} else {
		goto L121
	}
L121:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_makepol_2_12), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errsave_finish(m, v493, int32(_a_F_makepol_2_1), int32(94), int32(_a_F_makepol_2_5))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L8
L125:
	;
	v531 = int32(1)
	v532 = F_errsave_start(m, v530)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L131
	}
L126:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	if v523 != int32(447) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v530 = v520
	goto L125
L128:
	;
	goto L129
L129:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+4)))
	if v526 == int32(0) {
		v530 = v520
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v551 = int32(1)
	goto L3
L131:
	;
	if v532 == int32(0) {
		v551 = v531
		goto L3
	} else {
		goto L132
	}
L132:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_makepol_2_13), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errsave_finish(m, v530, int32(_a_F_makepol_2_1), int32(284), int32(_a_F_makepol_2_2))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v551 = v531
	goto L3
}
func F_maketree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = F_palloc(m, int32(12))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = l0
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v13 != int32(2) {
				return v4
			} else {
				v18 = F_maketree(m, l0+int32(12))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v18
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					if v21 == int32(1) {
						return v4
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v28 = F_maketree(m, l0+v24*int32(12))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = v28
							return v4
						}
					}
				}
			}
		}
	}
}
func F_manifest_process_version(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	if l1 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		m.T0[v7].(func(*base.Module, int32, int32, int32))(m, l0, int32(_a_F_manifest_process_version_0), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_manifest_process_wal_range(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_palloc(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = l3
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
		v14 = F_lappend(m, v13, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v14
			return
		}
	}
}
func F_markRelsAsNulledBy(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v11 - int32(63) {
	case 0:
		v35 = int32(4)
		goto L1
	case 1:
		goto L2
	default:
		goto L3
	}
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v35)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = v38
	goto L10
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_markRelsAsNulledBy(m, l0, v28, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
	F_errmsg_internal(m, int32(_a_F_markRelsAsNulledBy_0), v8)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(_a_F_markRelsAsNulledBy_1), int32(1795), int32(_a_F_markRelsAsNulledBy_2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_markRelsAsNulledBy(m, l0, v31, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v35 = int32(36)
	goto L1
L10:
	;
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v57 = v52 + v37<<(uint(int32(2))%32) - int32(4)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = F_bms_add_member(m, v58, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v46 = v44
	goto L14
L13:
	;
	v46 = int32(0)
	goto L14
L14:
	;
	if v46 < v37 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = F_lappend(m, v40, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L11
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v49
	v40 = v49
	goto L10
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
	m.G0 = v8 + int32(16)
	return
}
func F_mark_dummy_rel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4 == int32(0) {
		v27 = F_GetMemoryChunkContext(m, l0)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = int32(_a_F_mark_dummy_rel_0)
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0])) = v27
			v33 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v47 = F_create_append_path(m, v33, l0, v33, v33, v33, v43, v33, v33, float64(-1))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_add_path(m, l0, v47)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_set_cheapest(m, l0)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0])) = v30
						return
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		v9 = v7
		for {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui32(int32(2)) <= base.Ui32(v12-int32(301)) {
				break
			} else {
				v9 = v11 + int32(72)
				continue
			}
			break
		}
		if v12 != int32(290) {
			v27 = F_GetMemoryChunkContext(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = int32(_a_F_mark_dummy_rel_0)
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0])) = v27
				v33 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v47 = F_create_append_path(m, v33, l0, v33, v33, v33, v43, v33, v33, float64(-1))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_add_path(m, l0, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_set_cheapest(m, l0)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0])) = v30
							return
						}
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
			if v19 == int32(0) {
				return
			} else {
				v27 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = int32(_a_F_mark_dummy_rel_0)
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0])) = v27
					v33 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					v47 = F_create_append_path(m, v33, l0, v33, v33, v33, v43, v33, v33, float64(-1))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_add_path(m, l0, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_set_cheapest(m, l0)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_mark_dummy_rel[0])) = v30
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_mask_page_content(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	base.MemoryFill(m, l0+int32(24), v4, int32(_a_F_mask_page_content_0))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	return
}
func F_mask_unused_space(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(base.Ui32(v11) < base.Ui32(v10))|base.B2i32(base.Ui32(int32(_a_F_mask_unused_space_0)) < base.Ui32(v13))|(base.B2i32(base.Ui32(v10) < base.Ui32(int32(24)))|base.B2i32(base.Ui32(v13) < base.Ui32(v11))) == int32(0) {
		v24 = v11 - v10
		if v24 != 0 {
			base.MemoryFill(m, l0+v10, int32(0), v24)
		} else {
		}
		m.G0 = v8 + int32(16)
		return
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
			F_errmsg_internal(m, int32(_a_F_mask_unused_space_1), v8)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_mask_unused_space_2), int32(82), int32(_a_F_mask_unused_space_3))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
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
func F_match_clause_to_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
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
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 float64
	_ = v316
	var v318 float64
	_ = v318
	var v320 float64
	_ = v320
	var v322 float64
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
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
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1171 int32
	_ = v1171
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1279 int32
	_ = v1279
	var v1306 int32
	_ = v1306
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	v26 = m.G0
	v28 = v26 - int32(128)
	m.G0 = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v28 + int32(128)
	return
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
	if base.Ui32(v33) < base.Ui32(v32) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v37&int32(1) == int32(0) {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v37 = v35
	goto L6
L5:
	;
	v37 = int32(1)
	goto L6
L6:
	;
	goto L3
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v42 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v57 = int32(0)
	goto L9
L9:
	;
	v73 = v57 << (uint(int32(2)) % 32)
	v74 = l3 + int32(4) + v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v75 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v146 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v81 = int32(0)
	if v81 < v78 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v84 = v78
	goto L16
L15:
	;
	v84 = v81
	goto L16
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v91 = int32(0)
	goto L17
L17:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v85+v91<<(uint(int32(2))%32))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v116 == l1 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L11
L19:
	;
	v119 = v91 + int32(1)
	if v119 != v84 {
		v91 = v119
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v1375 = v57 + int32(1)
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v1375 < v1376 {
		v57 = v1375
		goto L9
	} else {
		goto L317
	}
L22:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v73)))
	if base.Ui32(v151) <= base.Ui32(int32(_a_F_match_clause_to_index_0)) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L31
	} else {
		goto L314
	}
L24:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v1325 = F_lappend(m, v1324, v1306)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L31
	} else {
		goto L313
	}
L25:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	switch v167 - int32(15) {
	case 0:
		goto L41
	default:
		goto L38
	case 2:
		goto L42
	case 5:
		goto L40
	case 22:
		goto L39
	}
L26:
	;
	v164 = F_match_boolean_index_clause(m, l0, l1, v57, l2)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L31
	} else {
		goto L34
	}
L27:
	;
	if base.B2i32(v151 == int32(424))|base.B2i32(v151 == int32(2222)) != 0 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v160 = F_op_in_opfamily(m, int32(91), v151)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L25
L31:
	;
	return
L32:
	;
	if v160 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	if v164 != 0 {
		v1306 = v164
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L25
L36:
	;
	if v1279 == int32(0) {
		goto L21
	} else {
		goto L312
	}
L37:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	v1270 = F_get_index_clause_from_support(m, l0, l1, v1269, v367, v57, l2)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L31
	} else {
		goto L311
	}
L38:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L246
L39:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if v455 != int32(403) {
		goto L21
	} else {
		goto L110
	}
L40:
	;
	v399 = int32(0)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+20)))
	if v401 != int32(1) {
		v1279 = v399
		goto L36
	} else {
		goto L95
	}
L41:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+28))
	if v356 == int32(0) {
		goto L21
	} else {
		goto L88
	}
L42:
	;
	v170 = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+28))
	if v172 == v170 {
		v1279 = v170
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v175 != int32(2) {
		v1279 = v170
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178+v73)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v73)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+68))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v191 = F_match_index_to_operand(m, v190, v57, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L31
	} else {
		goto L46
	}
L45:
	;
	v232 = F_match_index_to_operand(m, v189, v57, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L31
	} else {
		goto L63
	}
L46:
	;
	if v191 == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v196 = F_bms_is_member(m, v187, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	if v196 != 0 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v198 = F_contain_volatile_functions(m, v189)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	if v198 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	if v180 != v184 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	F_set_opfuncid(m, v171)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L31
	} else {
		goto L61
	}
L53:
	;
	v202 = v180
	goto L55
L54:
	;
	v202 = int32(0)
	goto L55
L55:
	;
	if v202 != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v203 = F_op_in_opfamily(m, v185, v183)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	if v203 == int32(0) {
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v208 = F_palloc0(m, int32(20))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L31
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = l1
	v218 = F_list_make1_impl(m, int32(1), v28+int32(24))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+16)) = v220
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+14)) = uint16(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+12)) = uint8(v220)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v218
	v1279 = v208
	goto L36
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	v230 = F_get_index_clause_from_support(m, l0, l1, v228, int32(0), v57, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L31
	} else {
		goto L62
	}
L62:
	;
	v1279 = v230
	goto L36
L63:
	;
	if v232 == int32(0) {
		v1279 = v170
		goto L36
	} else {
		goto L64
	}
L64:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v237 = F_bms_is_member(m, v187, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	if v237 != 0 {
		v1279 = v170
		goto L36
	} else {
		goto L66
	}
L66:
	;
	v239 = F_contain_volatile_functions(m, v190)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	if v239 != 0 {
		v1279 = v170
		goto L36
	} else {
		goto L68
	}
L68:
	;
	if v180 != v184 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	F_set_opfuncid(m, v171)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L31
	} else {
		goto L86
	}
L70:
	;
	v243 = v180
	goto L72
L71:
	;
	v243 = int32(0)
	goto L72
L72:
	;
	if v243 != 0 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v244 = F_get_commutator(m, v185)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L31
	} else {
		goto L74
	}
L74:
	;
	if v244 == int32(0) {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v248 = F_op_in_opfamily(m, v244, v183)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L31
	} else {
		goto L76
	}
L76:
	;
	if v248 == int32(0) {
		goto L69
	} else {
		goto L77
	}
L77:
	;
	v252 = m.G0
	v254 = v252 - int32(16)
	m.G0 = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v258 = F_palloc0(m, int32(36))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = int32(17)
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v256)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+16)) = v266
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v256)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+24)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v256)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+32)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v244
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256)+28))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+12)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+8)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v279
	v285 = F_list_make2_impl(m, v254+int32(4), v254)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L31
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+28)) = v285
	v289 = F_palloc0(m, int32(168))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L31
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = int32(318)
	base.MemoryCopy(m, v289, l1, int32(168))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v258
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+44)) = v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+48)) = v298
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+100)) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+104)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+108)) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v307 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+116)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v289)+112)) = v306
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v311 == v312 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v314 = v244
	goto L83
L82:
	;
	v314 = v307
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289)+124)) = v314
	v316 = *(*float64)(unsafe.Add(mBase, uint32(l1)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v289)+128)) = v316
	v318 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v289)+136)) = v318
	v320 = *(*float64)(unsafe.Add(mBase, uint32(l1)+152))
	*(*float64)(unsafe.Add(mBase, uint32(v289)+144)) = v320
	v322 = *(*float64)(unsafe.Add(mBase, uint32(l1)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v289)+160)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v289)+152)) = v322
	m.G0 = v254 + int32(16)
	v330 = F_palloc0(m, int32(20))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L31
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v289
	v340 = F_list_make1_impl(m, int32(1), v28+int32(20))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L31
	} else {
		goto L85
	}
L85:
	;
	v342 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+16)) = v342
	*(*uint16)(unsafe.Add(mBase, uint32(v330)+14)) = uint16(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v330)+12)) = uint8(v342)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+8)) = v340
	v1279 = v330
	goto L36
L86:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	v353 = F_get_index_clause_from_support(m, l0, l1, v351, int32(1), v57, l2)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L31
	} else {
		goto L87
	}
L87:
	;
	v1279 = v353
	goto L36
L88:
	;
	v359 = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v360 <= v359 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	v367 = v359
	goto L90
L90:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388+v367<<(uint(int32(2))%32))))
	v393 = F_match_index_to_operand(m, v392, v57, l2)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L31
	} else {
		goto L92
	}
L91:
	;
	goto L21
L92:
	;
	if v393 != 0 {
		goto L37
	} else {
		goto L93
	}
L93:
	;
	v396 = v367 + int32(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v396 < v397 {
		v367 = v396
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400)+28))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v408 = F_pull_varnos(m, l0, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L31
	} else {
		goto L96
	}
L96:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v410+v73)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v413+v73)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v400)+24))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+68))
	v420 = F_match_index_to_operand(m, v406, v57, l2)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L31
	} else {
		goto L97
	}
L97:
	;
	if v420 == int32(0) {
		v1279 = v399
		goto L36
	} else {
		goto L98
	}
L98:
	;
	v424 = F_bms_is_member(m, v419, v408)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L31
	} else {
		goto L99
	}
L99:
	;
	if v424 != 0 {
		v1279 = v399
		goto L36
	} else {
		goto L100
	}
L100:
	;
	v426 = F_contain_volatile_functions(m, v407)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L31
	} else {
		goto L101
	}
L101:
	;
	if v412 != v416 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v430 = v412
	goto L104
L103:
	;
	v430 = int32(0)
	goto L104
L104:
	;
	if v426|v430 != 0 {
		v1279 = v399
		goto L36
	} else {
		goto L105
	}
L105:
	;
	v432 = F_op_in_opfamily(m, v417, v415)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L31
	} else {
		goto L106
	}
L106:
	;
	if v432 == int32(0) {
		v1279 = v399
		goto L36
	} else {
		goto L107
	}
L107:
	;
	v437 = F_palloc0(m, int32(20))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L31
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = l1
	v447 = F_list_make1_impl(m, int32(1), v28+int32(28))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L31
	} else {
		goto L109
	}
L109:
	;
	v449 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+16)) = v449
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+14)) = uint16(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+12)) = uint8(v449)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+8)) = v447
	v1279 = v437
	goto L36
L110:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v458+v73)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+68))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+24))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+12))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v463)+20))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v473+v73)))
	if v475 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v463)+16))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	if v475 != v478 {
		goto L21
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v480 = F_match_index_to_operand(m, v472, v57, l2)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L31
	} else {
		goto L117
	}
L114:
	;
	goto L113
L115:
	;
	v508 = F_get_op_opfamily_strategy(m, v506, v460)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L31
	} else {
		goto L133
	}
L116:
	;
	v491 = F_match_index_to_operand(m, v469, v57, l2)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L31
	} else {
		goto L124
	}
L117:
	;
	if v480 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v484 = F_pull_varnos(m, l0, v469)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L31
	} else {
		goto L119
	}
L119:
	;
	v486 = F_bms_is_member(m, v462, v484)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L31
	} else {
		goto L120
	}
L120:
	;
	if v486 != 0 {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v488 = F_contain_volatile_functions(m, v469)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L31
	} else {
		goto L122
	}
L122:
	;
	if v488 != 0 {
		goto L116
	} else {
		goto L123
	}
L123:
	;
	v506 = v466
	v507 = int32(1)
	goto L115
L124:
	;
	if v491 == int32(0) {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	v495 = F_pull_varnos(m, l0, v472)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L31
	} else {
		goto L126
	}
L126:
	;
	v497 = F_bms_is_member(m, v462, v495)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L31
	} else {
		goto L127
	}
L127:
	;
	if v497 != 0 {
		goto L21
	} else {
		goto L128
	}
L128:
	;
	v499 = F_contain_volatile_functions(m, v472)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L31
	} else {
		goto L129
	}
L129:
	;
	if v499 != 0 {
		goto L21
	} else {
		goto L130
	}
L130:
	;
	v501 = F_get_commutator(m, v466)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L31
	} else {
		goto L131
	}
L131:
	;
	if v501 == int32(0) {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	v506 = v501
	v507 = int32(0)
	goto L115
L133:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v508))|base.B2i32(int32(1)<<(uint(v508)%32)&int32(54) == int32(0)) != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	v520 = F_palloc0(m, int32(20))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L31
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(281)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v520)+14)) = uint16(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v520)+4)) = l1
	if v507 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v529 = int32(24)
	goto L138
L137:
	;
	v529 = int32(20)
	goto L138
L138:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v524+v529)))
	if v507 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v534 = int32(20)
	goto L141
L140:
	;
	v534 = int32(24)
	goto L141
L141:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v524+v534)))
	v538 = v57 << (uint(int32(2)) % 32)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v538+v539)))
	F_get_op_opfamily_properties(m, v506, v541, int32(0), v28+int32(124), v28+int32(120), v28+int32(116))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L31
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v57
	v556 = F_list_make1_impl(m, int32(471), v28+int32(84))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L31
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+16)) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v506
	v564 = F_list_make1_impl(m, int32(472), v28+int32(80))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L31
	} else {
		goto L144
	}
L144:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v566+v538)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v568
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v568
	v574 = F_list_make1_impl(m, int32(472), v28+int32(76))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L31
	} else {
		goto L145
	}
L145:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v28)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v576
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v576
	v582 = F_list_make1_impl(m, int32(472), v28+int32(72))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L31
	} else {
		goto L146
	}
L146:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v584
	v590 = F_list_make1_impl(m, int32(472), v28+int32(68))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L31
	} else {
		goto L147
	}
L147:
	;
	v598 = v564
	v601 = int32(1)
	v603 = v574
	v605 = v590
	v607 = v582
	goto L148
L148:
	;
	if v536 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	if v768 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L150:
	;
	goto L149
L151:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v620 = v618
	goto L153
L152:
	;
	v620 = int32(0)
	goto L153
L153:
	;
	if v620 <= v601 {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v623 = v601 << (uint(int32(2)) % 32)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v531)+12))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v623+v624)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v627+v623)))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v631+v623)))
	if v507 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v636 = F_get_commutator(m, v633)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L31
	} else {
		goto L158
	}
L156:
	;
	v640 = v633
	goto L157
L157:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+68))
	v643 = F_pull_varnos(m, l0, v626)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L31
	} else {
		goto L160
	}
L158:
	;
	if v636 == int32(0) {
		goto L150
	} else {
		goto L159
	}
L159:
	;
	v640 = v636
	goto L157
L160:
	;
	v645 = F_bms_is_member(m, v642, v643)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L31
	} else {
		goto L161
	}
L161:
	;
	if v645 != 0 {
		goto L150
	} else {
		goto L162
	}
L162:
	;
	v647 = F_contain_volatile_functions(m, v626)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L31
	} else {
		goto L163
	}
L163:
	;
	if v647 != 0 {
		goto L150
	} else {
		goto L164
	}
L164:
	;
	v649 = int32(0)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v650 <= v649 {
		goto L150
	} else {
		goto L165
	}
L165:
	;
	v657 = v649
	goto L166
L166:
	;
	v678 = F_match_index_to_operand(m, v629, v657, l2)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L31
	} else {
		goto L170
	}
L167:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v708 <= v706 {
		goto L150
	} else {
		goto L177
	}
L168:
	;
	goto L167
L169:
	;
	v703 = v657 + int32(1)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v703 < v704 {
		v657 = v703
		goto L166
	} else {
		goto L176
	}
L170:
	;
	if v678 == int32(0) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v683 = v657 << (uint(int32(2)) % 32)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v683+v684)))
	v687 = F_get_op_opfamily_strategy(m, v640, v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L31
	} else {
		goto L172
	}
L172:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	if v687 != v689 {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v691+v683)))
	if v693 == int32(0) {
		v706 = v657
		goto L168
	} else {
		goto L174
	}
L174:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)+12))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v697+v623)))
	if v693 == v699 {
		v706 = v657
		goto L168
	} else {
		goto L175
	}
L175:
	;
	goto L169
L176:
	;
	v706 = v703
	goto L168
L177:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v520)+16))
	v711 = F_lappend_int(m, v710, v706)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L31
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+16)) = v711
	v715 = v706 << (uint(int32(2)) % 32)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v715+v716)))
	F_get_op_opfamily_properties(m, v640, v718, int32(0), v28+int32(124), v28+int32(120), v28+int32(116))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L31
	} else {
		goto L179
	}
L179:
	;
	v730 = F_lappend_oid(m, v598, v640)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L31
	} else {
		goto L180
	}
L180:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v732+v715)))
	v735 = F_lappend_oid(m, v603, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L31
	} else {
		goto L181
	}
L181:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v28)+120))
	v738 = F_lappend_oid(m, v607, v737)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L31
	} else {
		goto L182
	}
L182:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	v741 = F_lappend_oid(m, v605, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L31
	} else {
		goto L183
	}
L183:
	;
	v598 = v730
	v601 = v601 + int32(1)
	v603 = v735
	v605 = v741
	v607 = v738
	goto L148
L184:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(92))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v1010
	v1015 = F_list_make1_impl(m, int32(1), v28+int32(44))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L31
	} else {
		goto L245
	}
L185:
	;
	if int32(2) <= v601 {
		goto L232
	} else {
		goto L233
	}
L186:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	switch v786 - int32(1) {
	case 0:
		v806 = int32(2)
		goto L195
	case 1, 3:
		goto L194
	default:
		goto L197
	case 4:
		goto L196
	}
L187:
	;
	v771 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+12)) = uint8(v771)
	goto L186
L188:
	;
	goto L189
L189:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	v774 = base.B2i32(v601 != v773)
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+12)) = uint8(v774)
	v776 = int32(0)
	if base.B2i32(v507 == v776)|v774 == v776 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = l1
	goto L184
L191:
	;
	goto L192
L192:
	;
	if v773 == v601 {
		v900 = v598
		goto L185
	} else {
		goto L193
	}
L193:
	;
	goto L186
L194:
	;
	v884 = int32(0)
	if base.B2i32(v598 == v884)|base.B2i32(v601 <= v884) != 0 {
		goto L226
	} else {
		goto L227
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v806
	v808 = int32(0)
	v812 = v808
	v814 = v808
	goto L201
L196:
	;
	v806 = int32(4)
	goto L195
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L31
	} else {
		goto L198
	}
L198:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v793
	F_errmsg_internal(m, int32(_a_F_match_clause_to_index_1), v28+int32(32))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L31
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_match_clause_to_index_2), int32(3643), int32(_a_F_match_clause_to_index_3))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L31
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	v835 = int32(0)
	if v603 == v835 {
		v845 = v835
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v846 = int32(0)
	if v607 == v846 {
		v857 = v846
		goto L206
	} else {
		goto L207
	}
L204:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	if v839 <= v814 {
		v845 = int32(0)
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v845 = v841 + v814<<(uint(int32(2))%32)
	goto L203
L206:
	;
	if v605 == int32(0) {
		v900 = v846
		goto L185
	} else {
		goto L209
	}
L207:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	if v851 <= v814 {
		v857 = int32(0)
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v607)+12))
	v857 = v853 + v814<<(uint(int32(2))%32)
	goto L206
L209:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v860 <= v814 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v900 = v812
	goto L185
L211:
	;
	goto L212
L212:
	;
	if v845 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v900 = v812
	goto L185
L214:
	;
	goto L215
L215:
	;
	if v857 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v900 = v812
	goto L185
L217:
	;
	goto L218
L218:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	if v866 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v900 = v812
	goto L185
L220:
	;
	goto L221
L221:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v857)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v866+v814<<(uint(int32(2))%32))))
	v875 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+124)))
	v876 = F_get_opfamily_member(m, v869, v870, v874, v875)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L31
	} else {
		goto L222
	}
L222:
	;
	if v876 == int32(0) {
		goto L23
	} else {
		goto L223
	}
L223:
	;
	v882 = F_lappend_oid(m, v812, v876)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L31
	} else {
		goto L224
	}
L224:
	;
	v812 = v882
	v814 = v814 + int32(1)
	goto L201
L225:
	;
	v900 = v894
	goto L185
L226:
	;
	v894 = int32(0)
	goto L228
L227:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	if v601 < v891 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	goto L225
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+4)) = v601
	goto L231
L230:
	;
	goto L231
L231:
	;
	v894 = v598
	goto L228
L232:
	;
	v923 = F_palloc0(m, int32(28))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L31
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+16)) = int32(0)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v900)+12))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v957)))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	v961 = F_copyObjectImpl(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L31
	} else {
		goto L241
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923))) = int32(37)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v923)+8)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v923)+4)) = v927
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	v931 = F_list_copy_head(m, v930, v601)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L31
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923)+12)) = v931
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	v935 = F_list_copy_head(m, v934, v601)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L31
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923)+16)) = v935
	v938 = F_list_copy_head(m, v536, v601)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L31
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923)+20)) = v938
	v941 = F_list_copy_head(m, v531, v601)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L31
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923)+24)) = v941
	v945 = int32(0)
	v952 = F_make_restrictinfo(m, l0, v923, int32(1), v945, v945, v945, v945, v945, v945, v945)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L31
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v952
	goto L184
L241:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v531)+12))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	v965 = F_copyObjectImpl(m, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L31
	} else {
		goto L242
	}
L242:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v967)+12))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	v970 = F_make_opclause(m, v958, v961, v965, v969)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L31
	} else {
		goto L243
	}
L243:
	;
	v973 = int32(0)
	v980 = F_make_restrictinfo(m, l0, v970, int32(1), v973, v973, v973, v973, v973, v973, v973)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L31
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v980
	goto L184
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+8)) = v1015
	v1306 = v520
	goto L24
L246:
	;
	if v1018 != int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1021 = int32(0)
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+107)))
	if v1022 != int32(1) {
		v1279 = v1021
		goto L36
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+108)))
	if v1238 != int32(1) {
		goto L21
	} else {
		goto L304
	}
L250:
	;
	v1025 = int32(0)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+8))
	if v1027 == v1025 {
		v1202 = v1025
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1207 = F_make_SAOP_expr(m, v1151, v1112, v1153, v1152, v1152, v1158, v1157&int32(1))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L31
	} else {
		goto L300
	}
L252:
	;
	F_list_free(m, v1202)
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L31
	} else {
		goto L299
	}
L253:
	;
	v1031 = int32(0)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+4))
	if v1032 <= v1031 {
		v1202 = v1031
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+68))
	v1037 = int32(0)
	v1047 = int32(1)
	v1054 = v1031
	v1056 = v1037
	v1057 = v1037
	v1058 = v1037
	v1060 = v1037
	v1062 = v1037
	goto L256
L255:
	;
	v1202 = v1171
	goto L252
L256:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+12))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1067+v1062<<(uint(int32(2))%32))))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)))
	if v1072 != int32(318) {
		v1171 = v1056
		goto L255
	} else {
		goto L258
	}
L257:
	;
	if v1112 != 0 {
		goto L251
	} else {
		goto L298
	}
L258:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+4))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	if v1076 != int32(17) {
		v1171 = v1056
		goto L255
	} else {
		goto L259
	}
L259:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+28))
	if v1079 == int32(0) {
		v1171 = v1056
		goto L255
	} else {
		goto L260
	}
L260:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+4))
	if v1082 != int32(2) {
		v1171 = v1056
		goto L255
	} else {
		goto L261
	}
L261:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+4))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	v1089 = F_match_index_to_operand(m, v1088, v57, l2)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L31
	} else {
		goto L264
	}
L262:
	;
	if v1047&int32(1) != 0 {
		goto L279
	} else {
		goto L280
	}
L263:
	;
	v1098 = F_match_index_to_operand(m, v1087, v57, l2)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L31
	} else {
		goto L270
	}
L264:
	;
	if v1089 == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+48))
	v1094 = F_bms_is_member(m, v1036, v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L31
	} else {
		goto L266
	}
L266:
	;
	if v1094 != 0 {
		goto L263
	} else {
		goto L267
	}
L267:
	;
	v1096 = F_contain_volatile_functions(m, v1087)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L31
	} else {
		goto L268
	}
L268:
	;
	if v1096 != 0 {
		goto L263
	} else {
		goto L269
	}
L269:
	;
	v1111 = v1087
	v1112 = v1088
	v1113 = v1085
	goto L262
L270:
	;
	if v1098 == int32(0) {
		v1171 = v1056
		goto L255
	} else {
		goto L271
	}
L271:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+44))
	v1103 = F_bms_is_member(m, v1036, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L31
	} else {
		goto L272
	}
L272:
	;
	if v1103 != 0 {
		v1171 = v1056
		goto L255
	} else {
		goto L273
	}
L273:
	;
	v1105 = F_contain_volatile_functions(m, v1088)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L31
	} else {
		goto L274
	}
L274:
	;
	if v1105 != 0 {
		v1171 = v1056
		goto L255
	} else {
		goto L275
	}
L275:
	;
	v1107 = F_get_commutator(m, v1085)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L31
	} else {
		goto L276
	}
L276:
	;
	if v1107 == int32(0) {
		v1171 = v1056
		goto L255
	} else {
		goto L277
	}
L277:
	;
	v1111 = v1088
	v1112 = v1087
	v1113 = v1107
	goto L262
L278:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1157 = base.B2i32(v1154 != int32(7)) | v1060
	v1158 = F_lappend(m, v1056, v1111)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L31
	} else {
		goto L296
	}
L279:
	;
	v1116 = F_exprType(m, v1111)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L31
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	if v1054 != v1113 {
		v1171 = v1056
		goto L255
	} else {
		goto L292
	}
L282:
	;
	v1118 = F_get_array_type(m, v1116)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L31
	} else {
		goto L283
	}
L283:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1120+v73)))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+24))
	if v1122 != v1124 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1126 = v1122
	goto L286
L285:
	;
	v1126 = int32(0)
	goto L286
L286:
	;
	if v1126 != 0 {
		v1171 = v1056
		goto L255
	} else {
		goto L287
	}
L287:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1129+v73)))
	v1132 = F_op_in_opfamily(m, v1113, v1131)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L31
	} else {
		goto L288
	}
L288:
	;
	if base.B2i32(v1118 == int32(0))|(base.B2i32(v1132 == int32(0))|base.B2i32(v1116 == int32(2249))) != 0 {
		v1171 = v1056
		goto L255
	} else {
		goto L289
	}
L289:
	;
	v1140 = F_exprType(m, v1112)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L31
	} else {
		goto L290
	}
L290:
	;
	if v1140 != int32(2249) {
		v1151 = v1113
		v1152 = v1124
		v1153 = v1116
		goto L278
	} else {
		goto L291
	}
L291:
	;
	v1171 = v1056
	goto L255
L292:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+24))
	if v1057 != v1145 {
		v1171 = v1056
		goto L255
	} else {
		goto L293
	}
L293:
	;
	v1147 = F_exprType(m, v1111)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L31
	} else {
		goto L294
	}
L294:
	;
	if v1147 != v1058 {
		v1171 = v1056
		goto L255
	} else {
		goto L295
	}
L295:
	;
	v1151 = v1054
	v1152 = v1057
	v1153 = v1058
	goto L278
L296:
	;
	v1162 = v1062 + int32(1)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+4))
	if v1162 < v1163 {
		v1047 = int32(0)
		v1054 = v1151
		v1056 = v1158
		v1057 = v1152
		v1058 = v1153
		v1060 = v1157
		v1062 = v1162
		goto L256
	} else {
		goto L297
	}
L297:
	;
	goto L257
L298:
	;
	v1171 = v1158
	goto L255
L299:
	;
	v1279 = v1021
	goto L36
L300:
	;
	v1210 = F_palloc0(m, int32(20))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L31
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1210)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1210))) = int32(281)
	v1216 = int32(0)
	v1223 = F_make_restrictinfo(m, l0, v1207, int32(1), v1216, v1216, v1216, v1216, v1216, v1216, v1216)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L31
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v1223
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v1223
	v1230 = F_list_make1_impl(m, int32(1), v28+int32(12))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L31
	} else {
		goto L303
	}
L303:
	;
	v1232 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1210)+16)) = v1232
	*(*uint16)(unsafe.Add(mBase, uint32(v1210)+14)) = uint16(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v1210)+12)) = uint8(v1232)
	*(*int32)(unsafe.Add(mBase, uint32(v1210)+8)) = v1230
	v1279 = v1210
	goto L36
L304:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v1241 != int32(52) {
		goto L21
	} else {
		goto L305
	}
L305:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+12)))
	if v1244 != 0 {
		goto L21
	} else {
		goto L306
	}
L306:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v1246 = F_match_index_to_operand(m, v1245, v57, l2)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L31
	} else {
		goto L307
	}
L307:
	;
	if v1246 == int32(0) {
		goto L21
	} else {
		goto L308
	}
L308:
	;
	v1251 = F_palloc0(m, int32(20))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L31
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1251))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = l1
	v1261 = F_list_make1_impl(m, int32(1), v28+int32(16))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L31
	} else {
		goto L310
	}
L310:
	;
	v1263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+16)) = v1263
	*(*uint16)(unsafe.Add(mBase, uint32(v1251)+14)) = uint16(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v1251)+12)) = uint8(v1263)
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+8)) = v1261
	v1306 = v1251
	goto L24
L311:
	;
	v1279 = v1270
	goto L36
L312:
	;
	v1306 = v1279
	goto L24
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v1325
	v1328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1328)
	goto L1
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v870
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v874
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v869
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v1337
	F_errmsg_internal(m, int32(_a_F_match_clause_to_index_4), v28+int32(48))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L31
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_match_clause_to_index_2), int32(3657), int32(_a_F_match_clause_to_index_3))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L31
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L317:
	;
	goto L10
}
func F_mbms_add_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = int32(0)
	if v3 <= l0|l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L15
	}
L4:
	;
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v22 = v19 + l0<<(uint(int32(2))%32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = F_bms_add_member(m, v23, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L14
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = v10
	goto L8
L7:
	;
	v12 = int32(0)
	goto L8
L8:
	;
	if v12 <= l0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v15 = F_lappend(m, v9, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L5
L12:
	;
	return int32(0)
L13:
	;
	v9 = v15
	goto L4
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
	return v9
L15:
	;
	F_errmsg_internal(m, int32(_a_F_mbms_add_member_0), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_mbms_add_member_1), int32(50), int32(_a_F_mbms_add_member_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mbstowcs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l1
	v18 = v13 + int32(12)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_mbstowcs[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 == v4 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v444 + int32(16)
	return v446
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v428
	v444 = v13
	v446 = l2
	goto L1
L3:
	;
	v444 = v422
	v446 = int32(-1)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v406
	v422 = v13
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_mbstowcs[1])) = int32(25)
	if v393 == int32(0) {
		v422 = v13
		goto L3
	} else {
		goto L84
	}
L6:
	;
	if v371&int32(255) != 0 {
		v391 = v374
		v393 = v376
		goto L5
	} else {
		goto L80
	}
L7:
	;
	v369 = v361 - int32(1)
	if v358 != 0 {
		v391 = v369
		v393 = v363
		goto L5
	} else {
		goto L79
	}
L8:
	;
	v183 = v4
	v184 = l2
	v186 = v19
	v187 = v4
	v188 = l0
	goto L50
L9:
	;
	v58 = v4
	v59 = l2
	v61 = v19
	v62 = int32(1)
	goto L23
L10:
	;
	v56 = F_strlen(m, v19)
	mBase = m.M
	v444 = v13
	v446 = v56
	goto L1
L11:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L22
	}
L14:
	;
	if l2 == int32(0) {
		v428 = v19
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v29 = l2
	v32 = v19
	v34 = l0
	goto L16
L16:
	;
	v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
	if v39 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v49
	v444 = v13
	v446 = l2 - v29
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v39 & int32(_a_F_mbstowcs_0)
	v45 = int32(1)
	v46 = v32 + v45
	v48 = v29 - v45
	if v48 != 0 {
		v29 = v48
		v32 = v46
		v34 = v34 + int32(4)
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v428 = v46
	goto L2
L22:
	;
	goto L8
L23:
	;
	if v62 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v72 = int32(base.Ui32(v70) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v72-int32(16)|(v58>>(uint(int32(26))%32)+v72)) {
		v358 = v58
		v359 = v59
		v361 = v61
		v363 = l0
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v111 = v59
	v113 = v61
	goto L38
L28:
	;
	v82 = v61 + int32(1)
	if v58&int32(33554432) == int32(0) {
		v106 = v82
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v107 = int32(1)
	v59 = v59 - v107
	v61 = v106
	v62 = v107
	goto L23
L30:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82))))
	if int32(-64) <= v87 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v391 = v61 - int32(1)
	v393 = l0
	goto L5
L32:
	;
	goto L33
L33:
	;
	v93 = v61 + int32(2)
	if v58&int32(_a_F_mbstowcs_1) == int32(0) {
		v106 = v93
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v93))))
	if int32(-64) <= v98 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v391 = v61 - int32(1)
	v393 = l0
	goto L5
L36:
	;
	goto L37
L37:
	;
	v106 = v61 + int32(3)
	goto L29
L38:
	;
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113))))
	if v113&int32(3)|base.B2i32(v122 <= int32(0)) != 0 {
		v154 = v122
		v155 = v111
		v157 = v113
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v174 = v154&int32(255) - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v174) {
		v371 = v154
		v372 = v155
		v374 = v157
		v376 = l0
		goto L6
	} else {
		goto L49
	}
L40:
	;
	if int32(0) < base.I32_extend8_s(v154) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if (v126-int32(16843009)|v126)&int32(-2139062144) != 0 {
		v154 = v126
		v155 = v111
		v157 = v113
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v133 = v111
	v135 = v113
	goto L43
L43:
	;
	v142 = int32(4)
	v143 = v133 - v142
	v145 = v135 + v142
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if (v146-int32(16843009)|v146)&int32(-2139062144) == int32(0) {
		v133 = v143
		v135 = v145
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v154 = v146
	v155 = v143
	v157 = v145
	goto L40
L45:
	;
	goto L44
L46:
	;
	v167 = int32(1)
	v111 = v155 - v167
	v113 = v157 + v167
	goto L38
L47:
	;
	goto L48
L48:
	;
	goto L39
L49:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174<<(uint(int32(2))%32))+uint32(_c_F_mbstowcs[2])))
	v58 = v181
	v59 = v155
	v61 = v157 + int32(1)
	v62 = int32(0)
	goto L23
L50:
	;
	if v187 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if v184 == int32(0) {
		v428 = v186
		goto L2
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v300 = int32(base.Ui32(v298) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v300-int32(16)|(v300+v183>>(uint(int32(26))%32))) {
		v358 = v183
		v359 = v184
		v361 = v186
		v363 = v188
		goto L7
	} else {
		goto L71
	}
L55:
	;
	v198 = v184
	v200 = v186
	v202 = v188
	goto L56
L56:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v208 = base.I32_extend8_s(v207)
	if v208 <= int32(0) {
		v278 = v208
		v279 = v198
		v281 = v200
		v282 = v207
		v283 = v202
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v289 = v282 - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v289) {
		v371 = v278
		v372 = v279
		v374 = v281
		v376 = v283
		goto L6
	} else {
		goto L70
	}
L58:
	;
	goto L57
L59:
	;
	if v200&int32(3)|base.B2i32(base.Ui32(v198) < base.Ui32(int32(5))) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v219 = v198
	v221 = v200
	v223 = v202
	goto L64
L61:
	;
	v262 = v198
	v264 = v200
	v265 = v207
	v266 = v202
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = v265
	v274 = int32(1)
	v275 = v264 + v274
	v277 = v262 - v274
	if v277 != 0 {
		v198 = v277
		v200 = v275
		v202 = v266 + int32(4)
		goto L56
	} else {
		goto L69
	}
L63:
	;
	v257 = v252 & int32(255)
	if base.I32_extend8_s(v252) <= int32(0) {
		v278 = v252
		v279 = v253
		v281 = v254
		v282 = v257
		v283 = v255
		goto L58
	} else {
		goto L68
	}
L64:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if (v228-int32(16843009)|v228)&int32(-2139062144) != 0 {
		v252 = v228
		v253 = v219
		v254 = v221
		v255 = v223
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v252 = v251
	v253 = v248
	v254 = v246
	v255 = v244
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v228 & int32(255)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v237
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v239
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v241
	v244 = v223 + int32(16)
	v245 = int32(4)
	v246 = v221 + v245
	v248 = v219 - v245
	if base.Ui32(v245) < base.Ui32(v248) {
		v219 = v248
		v221 = v246
		v223 = v244
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v262 = v253
	v264 = v254
	v265 = v257
	v266 = v255
	goto L62
L69:
	;
	v428 = v275
	goto L2
L70:
	;
	v292 = int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v289<<(uint(int32(2))%32))+uint32(_c_F_mbstowcs[2])))
	v183 = v296
	v184 = v279
	v186 = v281 + v292
	v187 = v292
	v188 = v283
	goto L50
L71:
	;
	v310 = v186 + int32(1)
	v315 = v298 - int32(128) | v183<<(uint(int32(6))%32)
	if int32(0) <= v315 {
		v340 = v315
		v343 = v310
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v184 = v184 - int32(1)
	v186 = v343
	v187 = int32(0)
	v188 = v188 + int32(4)
	goto L50
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_mbstowcs[1])) = int32(25)
	v406 = v186 - int32(1)
	goto L4
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v340
	goto L72
L75:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	v320 = v318 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v320) {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v324 = v315 << (uint(int32(6)) % 32)
	v325 = v320 | v324
	v327 = v186 + int32(2)
	if int32(0) <= v324 {
		v340 = v325
		v343 = v327
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v332 = v330 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v332) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v340 = v332 | v325<<(uint(int32(6))%32)
	v343 = v186 + int32(3)
	goto L74
L79:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	v371 = v370
	v372 = v359
	v374 = v369
	v376 = v363
	goto L6
L80:
	;
	if v376 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v383 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v376))) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v383
	goto L83
L82:
	;
	goto L83
L83:
	;
	v444 = v13
	v446 = l2 - v372
	goto L1
L84:
	;
	v406 = v391
	goto L4
}
func F_mcelem_array_contained_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 float32
	_ = v46
	var v48 int32
	_ = v48
	var v52 float32
	_ = v52
	var v53 float32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 float32
	_ = v66
	var v70 float32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v85 float32
	_ = v85
	var v87 float32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v130 float32
	_ = v130
	var v132 float32
	_ = v132
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 float32
	_ = v162
	var v168 float32
	_ = v168
	var v169 float32
	_ = v169
	var v172 float32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 float32
	_ = v186
	var v188 float32
	_ = v188
	var v211 int32
	_ = v211
	var v218 float32
	_ = v218
	var v220 float32
	_ = v220
	var v241 int32
	_ = v241
	var v248 float32
	_ = v248
	var v250 float32
	_ = v250
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v279 float32
	_ = v279
	var v281 float32
	_ = v281
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 float32
	_ = v314
	var v316 float32
	_ = v316
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 float32
	_ = v337
	var v338 float32
	_ = v338
	var v341 float32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v355 float32
	_ = v355
	var v357 float32
	_ = v357
	var v380 int32
	_ = v380
	var v386 float32
	_ = v386
	var v388 float32
	_ = v388
	var v406 int32
	_ = v406
	var v407 float32
	_ = v407
	var v409 float32
	_ = v409
	var v411 float32
	_ = v411
	var v413 float32
	_ = v413
	var v414 float32
	_ = v414
	var v415 float32
	_ = v415
	var v426 float32
	_ = v426
	var v428 int32
	_ = v428
	var v440 float32
	_ = v440
	var v442 float32
	_ = v442
	var v461 float64
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 float64
	_ = v473
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var __phi497 int32
	_ = __phi497
	var v500 int32
	_ = v500
	var __phi500 int32
	_ = __phi500
	var v505 int32
	_ = v505
	var __phi505 int32
	_ = __phi505
	var v526 int32
	_ = v526
	var v529 float32
	_ = v529
	var v536 float32
	_ = v536
	var v537 float32
	_ = v537
	var v546 int32
	_ = v546
	var v572 int32
	_ = v572
	var v577 float32
	_ = v577
	var v581 float32
	_ = v581
	var v586 float32
	_ = v586
	var v631 int32
	_ = v631
	var v641 int32
	_ = v641
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var __phi662 int32
	_ = __phi662
	var v665 int32
	_ = v665
	var __phi665 int32
	_ = __phi665
	var v666 int32
	_ = v666
	var __phi666 int32
	_ = __phi666
	var v691 int32
	_ = v691
	var v694 float32
	_ = v694
	var v701 float32
	_ = v701
	var v702 float32
	_ = v702
	var v713 int32
	_ = v713
	var v739 int32
	_ = v739
	var v744 float32
	_ = v744
	var v748 float32
	_ = v748
	var v753 float32
	_ = v753
	var v790 int32
	_ = v790
	var v798 int32
	_ = v798
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v842 int32
	_ = v842
	var v848 float32
	_ = v848
	var v854 int32
	_ = v854
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v895 int32
	_ = v895
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 float32
	_ = v908
	var v910 float32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 float32
	_ = v917
	var v919 float32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 float32
	_ = v961
	var v963 float32
	_ = v963
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1073 float32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1093 float32
	_ = v1093
	var v1102 int32
	_ = v1102
	var v1112 int32
	_ = v1112
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1147 int32
	_ = v1147
	var v1148 float32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1167 float32
	_ = v1167
	var v1170 float32
	_ = v1170
	var v1171 float32
	_ = v1171
	var v1179 int32
	_ = v1179
	var v1181 float32
	_ = v1181
	var v1183 float32
	_ = v1183
	var v1187 float32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1217 float32
	_ = v1217
	var v1226 int32
	_ = v1226
	var v1233 float32
	_ = v1233
	var v1234 float32
	_ = v1234
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1275 float32
	_ = v1275
	var v1295 int32
	_ = v1295
	var v1297 float32
	_ = v1297
	var v1301 float32
	_ = v1301
	var v1304 float32
	_ = v1304
	var v1308 float32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1321 float32
	_ = v1321
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1350 float32
	_ = v1350
	var v1358 float32
	_ = v1358
	var v1388 float64
	_ = v1388
	v17 = int32(0)
	v32 = int32(3)
	if base.B2i32(l2 == v17)|base.B2i32(l3 != l1+v32)|(base.B2i32(l6 == v17)|base.B2i32(l7 < v32)) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v1388 = float64(0.005)
	goto L3
L2:
	;
	v43 = int32(2)
	v45 = l2 + l1<<(uint(v43)%32)
	v46 = *(*float32)(unsafe.Add(mBase, uint32(v45)))
	v48 = l7 - int32(1)
	v52 = *(*float32)(unsafe.Add(mBase, uint32(l6+v48<<(uint(v43)%32))))
	v53 = *(*float32)(unsafe.Add(mBase, uint32(v45)+8))
	v56 = F_palloc(m, l5<<(uint(v43)%32))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v1388
L4:
	;
	return float64(0)
L5:
	;
	if l5 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l1 <= v272 {
		v440 = v279
		v442 = v281
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v272 = int32(0)
	v279 = v52
	v281 = float32(1)
	v285 = v17
	goto L6
L8:
	;
	goto L9
L9:
	;
	v66 = base.F32_mul(v46, float32(0.5))
	if base.F64_gt(base.F64_promote_f32(v66), float64(0.005)) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v70 = float32(0.005)
	goto L12
L11:
	;
	v70 = v66
	goto L12
L12:
	;
	v72 = l8 + int32(104)
	v78 = int32(0)
	v85 = v52
	v87 = float32(1)
	v91 = v17
	v92 = v17
	goto L13
L13:
	;
	if v92 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v272 = v241
	v279 = v248
	v281 = v250
	v285 = v254
	goto L6
L15:
	;
	v267 = v92 + int32(1)
	if v267 != l5 {
		v78 = v241
		v85 = v248
		v87 = v250
		v91 = v254
		v92 = v267
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l8)+32))
	v106 = l4 + v92<<(uint(int32(2))%32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106-int32(4))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v111 = F_FunctionCall2Coll(m, v72, v103, v109, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if l1 <= v78 {
		v179 = v78
		v186 = v85
		v188 = v87
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v111 == int32(0) {
		v241 = v78
		v248 = v85
		v250 = v87
		v254 = v91
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v241 = v211
	v248 = v218
	v250 = v220
	v254 = v91 + int32(1)
	goto L15
L22:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v56+v91<<(uint(int32(2))%32)))) = v70
	v211 = v179
	v218 = v186
	v220 = v188
	goto L21
L23:
	;
	v123 = v78
	v130 = v85
	v132 = v87
	goto L24
L24:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l8)+32))
	v150 = v123 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0+v150)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l4+v92<<(uint(int32(2))%32))))
	v154 = F_FunctionCall2Coll(m, v72, v148, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	v179 = l1
	v186 = v169
	v188 = v172
	goto L22
L26:
	;
	if int32(0) <= v154 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v154 != 0 {
		v179 = v123
		v186 = v130
		v188 = v132
		goto L22
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v168 = *(*float32)(unsafe.Add(mBase, uint32(l2+v150)))
	v169 = base.F32_sub(v130, v168)
	v172 = base.F32_mul(v132, base.F32_sub(float32(1), v168))
	v174 = v123 + int32(1)
	if v174 != l1 {
		v123 = v174
		v130 = v169
		v132 = v172
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v162 = *(*float32)(unsafe.Add(mBase, uint32(l2+v150)))
	*(*float32)(unsafe.Add(mBase, uint32(v56+v91<<(uint(int32(2))%32)))) = v162
	v211 = v123 + int32(1)
	v218 = base.F32_sub(v130, v162)
	v220 = v132
	goto L21
L31:
	;
	goto L25
L32:
	;
	goto L14
L33:
	;
	v461 = F_exp(m, base.F64_promote_f32(base.F32_neg(v440)))
	mBase = m.M
	v462 = l1 + v285
	if v462 <= int32(0) {
		v483 = v285
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v300 = (l1 - v272) & int32(3)
	if v300 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v272-l1) {
		v440 = v355
		v442 = v357
		goto L33
	} else {
		goto L42
	}
L36:
	;
	v349 = v272
	v355 = v279
	v357 = v281
	goto L35
L37:
	;
	goto L38
L38:
	;
	v308 = v272
	v312 = int32(0)
	v314 = v279
	v316 = v281
	goto L39
L39:
	;
	v332 = int32(1)
	v333 = v308 + v332
	v337 = *(*float32)(unsafe.Add(mBase, uint32(l2+v308<<(uint(int32(2))%32))))
	v338 = base.F32_sub(v314, v337)
	v341 = base.F32_mul(v316, base.F32_sub(float32(1), v337))
	v343 = v312 + v332
	if v343 != v300 {
		v308 = v333
		v312 = v343
		v314 = v338
		v316 = v341
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v349 = v333
	v355 = v338
	v357 = v341
	goto L35
L41:
	;
	goto L40
L42:
	;
	v380 = v349
	v386 = v355
	v388 = v357
	goto L43
L43:
	;
	v406 = l2 + v380<<(uint(int32(2))%32)
	v407 = *(*float32)(unsafe.Add(mBase, uint32(v406)))
	v409 = *(*float32)(unsafe.Add(mBase, uint32(v406)+4))
	v411 = *(*float32)(unsafe.Add(mBase, uint32(v406)+8))
	v413 = *(*float32)(unsafe.Add(mBase, uint32(v406)+12))
	v414 = base.F32_sub(base.F32_sub(base.F32_sub(base.F32_sub(v386, v407), v409), v411), v413)
	v415 = float32(1)
	v426 = base.F32_mul(base.F32_mul(base.F32_mul(base.F32_mul(v388, base.F32_sub(v415, v407)), base.F32_sub(v415, v409)), base.F32_sub(v415, v411)), base.F32_sub(v415, v413))
	v428 = v380 + int32(4)
	if v428 != l1 {
		v380 = v428
		v386 = v414
		v388 = v426
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v440 = v414
	v442 = v426
	goto L33
L45:
	;
	goto L44
L46:
	;
	v488 = v483<<(uint(int32(2))%32) + int32(4)
	v489 = F_palloc(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L50
	}
L47:
	;
	v467 = base.I32_div_s(l1*int32(100), v462)
	if v285 <= v467 {
		v483 = v285
		goto L46
	} else {
		goto L48
	}
L48:
	;
	F_pg_qsort(m, v56, v285, int32(4), int32(1242))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v473 = base.F64_convert_i32_s(l1)
	v483 = base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_sub(base.F64_sqrt(base.F64_add(base.F64_mul(v473, v473), base.F64_mul(v473, float64(400)))), v473), float64(0.5)))
	goto L46
L50:
	;
	v491 = F_palloc(m, v488)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = int32(1065353216)
	if v483 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	F_pfree(m, v631)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L71
	}
L53:
	;
	v631 = v491
	v641 = v489
	goto L52
L54:
	;
	goto L55
L55:
	;
	__phi497 = int32(1)
	__phi500 = v489
	__phi505 = v491
	v497 = __phi497
	v500 = __phi500
	v505 = __phi505
	goto L56
L56:
	;
	if v483 < v497 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v631 = v500
	v641 = v505
	goto L52
L58:
	;
	v526 = v483
	goto L60
L59:
	;
	v526 = v497
	goto L60
L60:
	;
	if int32(0) <= v526 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v529 = *(*float32)(unsafe.Add(mBase, uint32(v500)))
	v536 = *(*float32)(unsafe.Add(mBase, uint32(v56+v497<<(uint(int32(2))%32)-int32(4))))
	v537 = base.F32_sub(float32(1), v536)
	*(*float32)(unsafe.Add(mBase, uint32(v505))) = base.F32_add(base.F32_mul(v529, v537), float32(0))
	v546 = int32(1)
	goto L64
L62:
	;
	goto L63
L63:
	;
	if v497 != v483 {
		__phi497 = v497 + int32(1)
		__phi500 = v505
		__phi505 = v500
		v497 = __phi497
		v500 = __phi500
		v505 = __phi505
		goto L56
	} else {
		goto L70
	}
L64:
	;
	v572 = v546 << (uint(int32(2)) % 32)
	v577 = *(*float32)(unsafe.Add(mBase, uint32(v572+v500-int32(4))))
	if base.Ui32(v546) < base.Ui32(v497) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L63
L66:
	;
	v581 = *(*float32)(unsafe.Add(mBase, uint32(v572+v500)))
	v586 = base.F32_add(base.F32_mul(v581, v537), float32(0))
	goto L68
L67:
	;
	v586 = float32(0)
	goto L68
L68:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v505+v572))) = base.F32_add(base.F32_mul(v577, v536), v586)
	if v546 != v526 {
		v546 = v546 + int32(1)
		goto L64
	} else {
		goto L69
	}
L69:
	;
	goto L65
L70:
	;
	goto L57
L71:
	;
	v653 = F_palloc(m, v488)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v655 = F_palloc(m, v488)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = int32(1065353216)
	if l1 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if base.F64_gt(base.F64_promote_f32(v440), float64(0.005)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L75:
	;
	v790 = v655
	v798 = v653
	goto L74
L76:
	;
	goto L77
L77:
	;
	__phi662 = v655
	__phi665 = v653
	__phi666 = int32(1)
	v662 = __phi662
	v665 = __phi665
	v666 = __phi666
	goto L78
L78:
	;
	if v483 < v666 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v790 = v665
	v798 = v662
	goto L74
L80:
	;
	if l1 != v666 {
		__phi662 = v665
		__phi665 = v662
		__phi666 = v666 + int32(1)
		v662 = __phi662
		v665 = __phi665
		v666 = __phi666
		goto L78
	} else {
		goto L92
	}
L81:
	;
	v691 = v483
	goto L83
L82:
	;
	v691 = v666
	goto L83
L83:
	;
	if v691 < int32(0) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v694 = *(*float32)(unsafe.Add(mBase, uint32(v665)))
	v701 = *(*float32)(unsafe.Add(mBase, uint32(l2+v666<<(uint(int32(2))%32)-int32(4))))
	v702 = base.F32_sub(float32(1), v701)
	*(*float32)(unsafe.Add(mBase, uint32(v662))) = base.F32_add(base.F32_mul(v694, v702), float32(0))
	if v483 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v713 = int32(1)
	goto L86
L86:
	;
	v739 = v713 << (uint(int32(2)) % 32)
	v744 = *(*float32)(unsafe.Add(mBase, uint32(v665+v739-int32(4))))
	if base.Ui32(v713) < base.Ui32(v666) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L80
L88:
	;
	v748 = *(*float32)(unsafe.Add(mBase, uint32(v665+v739)))
	v753 = base.F32_add(base.F32_mul(v748, v702), float32(0))
	goto L90
L89:
	;
	v753 = float32(0)
	goto L90
L90:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v662+v739))) = base.F32_add(base.F32_mul(v744, v701), v753)
	if v713 != v691 {
		v713 = v713 + int32(1)
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	goto L79
L93:
	;
	F_pfree(m, v1033)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L4
	} else {
		goto L118
	}
L94:
	;
	v1033 = v790
	v1038 = v798
	goto L93
L95:
	;
	goto L96
L96:
	;
	v823 = int32(0)
	if v823 <= v483 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v826 = int32(1)
	v828 = v483 + v826
	if v828 <= v826 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v1033 = v798
	v1038 = v790
	goto L93
L100:
	;
	v831 = v826
	goto L102
L101:
	;
	v831 = v828
	goto L102
L102:
	;
	v833 = v831 << (uint(int32(2)) % 32)
	if v833 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	base.MemoryFill(m, v790, int32(0), v833)
	goto L105
L104:
	;
	goto L105
L105:
	;
	v842 = v828
	v848 = base.F32_demote_f64(v461)
	v854 = v823
	goto L106
L106:
	;
	if v483 < v854 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L99
L108:
	;
	v994 = int32(1)
	v997 = v854 + v994
	if v997 != v831 {
		v842 = v842 - v994
		v848 = base.F32_mul(v848, base.F32_div(v440, base.F32_convert_i32_u(v997)))
		v854 = v997
		goto L106
	} else {
		goto L117
	}
L109:
	;
	v868 = v790 + v854<<(uint(int32(2))%32)
	v869 = int32(0)
	if v483 != v854 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v879 = v869
	v895 = int32(0)
	goto L113
L111:
	;
	v932 = v869
	goto L112
L112:
	;
	v958 = v932 << (uint(int32(2)) % 32)
	v959 = v868 + v958
	v961 = *(*float32)(unsafe.Add(mBase, uint32(v958+v798)))
	v963 = *(*float32)(unsafe.Add(mBase, uint32(v959)))
	*(*float32)(unsafe.Add(mBase, uint32(v959))) = base.F32_add(base.F32_mul(v961, v848), v963)
	goto L108
L113:
	;
	v904 = int32(2)
	v905 = v879 << (uint(v904) % 32)
	v906 = v868 + v905
	v908 = *(*float32)(unsafe.Add(mBase, uint32(v905+v798)))
	v910 = *(*float32)(unsafe.Add(mBase, uint32(v906)))
	*(*float32)(unsafe.Add(mBase, uint32(v906))) = base.F32_add(base.F32_mul(v908, v848), v910)
	v914 = v905 | int32(4)
	v915 = v868 + v914
	v917 = *(*float32)(unsafe.Add(mBase, uint32(v914+v798)))
	v919 = *(*float32)(unsafe.Add(mBase, uint32(v915)))
	*(*float32)(unsafe.Add(mBase, uint32(v915))) = base.F32_add(base.F32_mul(v917, v848), v919)
	v923 = v879 + v904
	v925 = v895 + v904
	if v925 != v842&int32(-2) {
		v879 = v923
		v895 = v925
		goto L113
	} else {
		goto L115
	}
L114:
	;
	if v842&int32(1) == int32(0) {
		goto L108
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v932 = v923
	goto L112
L117:
	;
	goto L107
L118:
	;
	v1061 = F_palloc(m, v488)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	if int32(0) <= v483 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v1073 = base.F32_div(float32(1), base.F32_convert_i32_u(l7-int32(2)))
	v1074 = int32(1)
	v1076 = v483 + v1074
	if v1076 <= v1074 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v1321 = float32(0)
	goto L122
L122:
	;
	F_pfree(m, v641)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L4
	} else {
		goto L154
	}
L123:
	;
	v1079 = v1074
	goto L125
L124:
	;
	v1079 = v1076
	goto L125
L125:
	;
	v1080 = int32(0)
	v1086 = v1080
	v1093 = float32(0)
	v1102 = v1080
	goto L126
L126:
	;
	if v48 <= v1086 {
		v1192 = v1086
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v1258 = int32(1)
	v1260 = v483 + v1258
	if v1260 <= v1258 {
		goto L145
	} else {
		goto L146
	}
L128:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v1061+v1102<<(uint(int32(2))%32)))) = v1234
	v1256 = v1102 + int32(1)
	if v1256 != v1079 {
		v1086 = v1226
		v1093 = v1233
		v1102 = v1256
		goto L126
	} else {
		goto L144
	}
L129:
	;
	v1217 = float32(0)
	if base.F32_gt(v1093, v1217) == int32(0) {
		v1226 = v1192
		v1233 = v1093
		v1234 = v1217
		goto L128
	} else {
		goto L143
	}
L130:
	;
	v1112 = v48 - v1086
	v1120 = v1086
	v1121 = int32(0)
	goto L133
L131:
	;
	if base.F32_gt(v1093, float32(0)) != 0 {
		goto L140
	} else {
		goto L141
	}
L132:
	;
	if v1121 <= int32(0) {
		v1192 = v1120
		goto L129
	} else {
		goto L138
	}
L133:
	;
	v1147 = l6 + v1120<<(uint(int32(2))%32)
	v1148 = *(*float32)(unsafe.Add(mBase, uint32(v1147)))
	if base.F32_le(v1148, base.F32_convert_i32_u(v1102)) == int32(0) {
		goto L132
	} else {
		goto L135
	}
L134:
	;
	if v1112 <= int32(0) {
		v1192 = v48
		goto L129
	} else {
		goto L137
	}
L135:
	;
	v1152 = int32(1)
	v1155 = v1121 + v1152
	if v1155 != v1112 {
		v1120 = v1120 + v1152
		v1121 = v1155
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v1179 = v48
	v1181 = float32(0)
	v1183 = base.F32_convert_i32_u(v1112 - int32(1))
	goto L131
L138:
	;
	v1167 = base.F32_convert_i32_u(v1121 - int32(1))
	v1170 = *(*float32)(unsafe.Add(mBase, uint32(v1147-int32(4))))
	v1171 = base.F32_sub(v1148, v1170)
	if base.F32_gt(v1171, float32(0)) == int32(0) {
		v1179 = v1120
		v1181 = v1171
		v1183 = v1167
		goto L131
	} else {
		goto L139
	}
L139:
	;
	v1179 = v1120
	v1181 = v1171
	v1183 = base.F32_add(base.F32_div(float32(0.5), v1171), v1167)
	goto L131
L140:
	;
	v1187 = base.F32_add(base.F32_div(float32(0.5), v1093), v1183)
	goto L142
L141:
	;
	v1187 = v1183
	goto L142
L142:
	;
	v1226 = v1179
	v1233 = v1181
	v1234 = base.F32_mul(v1073, v1187)
	goto L128
L143:
	;
	v1226 = v1192
	v1233 = v1093
	v1234 = base.F32_div(v1073, v1093)
	goto L128
L144:
	;
	goto L127
L145:
	;
	v1263 = v1258
	goto L147
L146:
	;
	v1263 = v1260
	goto L147
L147:
	;
	v1269 = int32(0)
	v1275 = float32(0)
	goto L148
L148:
	;
	v1295 = v1269 << (uint(int32(2)) % 32)
	v1297 = *(*float32)(unsafe.Add(mBase, uint32(v1038+v1295)))
	if base.F32_gt(v1297, float32(0)) != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v1321 = v1308
	goto L122
L150:
	;
	v1301 = *(*float32)(unsafe.Add(mBase, uint32(v1295+v1061)))
	v1304 = *(*float32)(unsafe.Add(mBase, uint32(v1295+v641)))
	v1308 = base.F32_add(v1275, base.F32_div(base.F32_mul(base.F32_mul(v1301, base.F32_demote_f64(base.F64_mul(v461, base.F64_promote_f32(v442)))), v1304), v1297))
	goto L152
L151:
	;
	v1308 = v1275
	goto L152
L152:
	;
	v1310 = v1269 + int32(1)
	if v1310 != v1263 {
		v1269 = v1310
		v1275 = v1308
		goto L148
	} else {
		goto L153
	}
L153:
	;
	goto L149
L154:
	;
	F_pfree(m, v1038)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	F_pfree(m, v1061)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	F_pfree(m, v56)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v1350 = base.F32_mul(base.F32_sub(float32(1), v53), v1321)
	if base.F32_lt(v1350, float32(0)) != 0 {
		v1358 = float32(0)
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1388 = base.F64_promote_f32(v1358)
	goto L3
L159:
	;
	if base.F32_gt(v1350, float32(1)) == int32(0) {
		v1358 = v1350
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v1358 = float32(1)
	goto L158
}
func F_mcelem_array_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v10 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+11)))
	F_deconstruct_array(m, l0, v20, v21, v22, v17+int32(8), v17+int32(4), v17+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return float64(0)
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
		if int32(0) < v33 {
			v37 = int32(0)
			v47 = v10
			v48 = v33
			v49 = v10
			for {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v37))))
				if v53 != 0 {
					v67 = v47
					v68 = v48
					v69 = int32(1)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					v56 = int32(2)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v55+v37<<(uint(v56)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(v56)%32)))) = v62
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					v67 = v47 + int32(1)
					v68 = v66
					v69 = v49
				}
				v71 = v37 + int32(1)
				if v71 < v68 {
					v37 = v71
					v47 = v67
					v48 = v68
					v49 = v69
					continue
				} else {
					break
				}
				break
			}
			if base.B2i32(l8 == int32(2751))&v69 != 0 {
				v122 = float64(0)
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				F_pfree(m, v123)
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return float64(0)
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					F_pfree(m, v126)
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return float64(0)
					} else {
						m.G0 = v17 + int32(16)
						return v122
					}
				}
			} else {
				v87 = v67
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				F_qsort_arg(m, v91, v87, int32(4), int32(1243), l1)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return float64(0)
				} else {
					if l8&int32(-2) == int32(2750) {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						v101 = F_mcelem_array_contain_overlap_selec(m, l2, l3, l4, l5, v100, v87, l8, l1)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return float64(0)
						} else {
							v122 = v101
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							F_pfree(m, v123)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return float64(0)
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
								F_pfree(m, v126)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return float64(0)
								} else {
									m.G0 = v17 + int32(16)
									return v122
								}
							}
						}
					} else {
						if l8 != int32(2752) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return float64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17))) = l8
								F_errmsg_internal(m, int32(_a_F_mcelem_array_selec_0), v17)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_mcelem_array_selec_1), int32(494), int32(_a_F_mcelem_array_selec_2))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							v106 = F_mcelem_array_contained_selec(m, l2, l3, l4, l5, v105, v87, l6, l7, l1)
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return float64(0)
							} else {
								v122 = v106
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								F_pfree(m, v123)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return float64(0)
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									F_pfree(m, v126)
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return float64(0)
									} else {
										m.G0 = v17 + int32(16)
										return v122
									}
								}
							}
						}
					}
				}
			}
		} else {
			v87 = v10
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			F_qsort_arg(m, v91, v87, int32(4), int32(1243), l1)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return float64(0)
			} else {
				if l8&int32(-2) == int32(2750) {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					v101 = F_mcelem_array_contain_overlap_selec(m, l2, l3, l4, l5, v100, v87, l8, l1)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return float64(0)
					} else {
						v122 = v101
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						F_pfree(m, v123)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return float64(0)
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							F_pfree(m, v126)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return float64(0)
							} else {
								m.G0 = v17 + int32(16)
								return v122
							}
						}
					}
				} else {
					if l8 != int32(2752) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return float64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v17))) = l8
							F_errmsg_internal(m, int32(_a_F_mcelem_array_selec_0), v17)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_mcelem_array_selec_1), int32(494), int32(_a_F_mcelem_array_selec_2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						v106 = F_mcelem_array_contained_selec(m, l2, l3, l4, l5, v105, v87, l6, l7, l1)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return float64(0)
						} else {
							v122 = v106
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							F_pfree(m, v123)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return float64(0)
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
								F_pfree(m, v126)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return float64(0)
								} else {
									m.G0 = v17 + int32(16)
									return v122
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_mcv_combine_selectivities(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64) float64 {
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v28 float64
	_ = v28
	v5 = float64(0)
	v7 = base.F64_sub(l0, l2)
	if base.F64_lt(v7, v5) != 0 {
		v15 = v5
	} else {
		if base.F64_gt(v7, float64(1)) == int32(0) {
			v15 = v7
		} else {
			v15 = float64(1)
		}
	}
	v17 = base.F64_sub(float64(1), l3)
	if base.F64_lt(v17, v15) != 0 {
		v19 = v17
	} else {
		v19 = v15
	}
	v20 = base.F64_add(l1, v19)
	if base.F64_lt(v20, float64(0)) != 0 {
		v28 = v5
	} else {
		if base.F64_gt(v20, float64(1)) == int32(0) {
			v28 = v20
		} else {
			v28 = float64(1)
		}
	}
	return v28
}
func F_mdc_free_2(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	m.T0[v2].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_mdc_init_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = F_pgp_load_digest(m, int32(2), v6+int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v11 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
			v20 = int32(0)
		} else {
			v20 = v11
		}
		m.G0 = v6 + int32(16)
		return v20
	}
}
func F_mdc_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	m.T0[v5].(func(*base.Module, int32, int32, int32))(m, l1, l2, l3)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_pushf_write(m, l0, l2, l3)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_mdfiletagmatches(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	return base.B2i32(v3 == v4)
}
func F_mdimmedsync(m *base.Module, l0 int32, l1 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_mdnblocks(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = l0 + l1<<(uint(int32(2))%32) + int32(40)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v23 = v20
	goto L3
L3:
	;
	v32 = F__mdfd_openseg(m, l0, l1, v23, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	if int32(0) < v23 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v32 != 0 {
		v23 = v23 + int32(1)
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v40 = l0 + l1<<(uint(int32(2))%32) + int32(56)
	v44 = v23
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v11 + int32(16)
	return
L10:
	;
	v50 = v44 - int32(1)
	v52 = v50 << (uint(int32(3)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v57 = F_FileSync(m, v55, int32(167772179))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	if v20 < v44 {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if int32(0) <= v57 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mdimmedsync[0])))
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v67 = F_errstart(m, v65, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v65 = int32(21)
	goto L18
L17:
	;
	v65 = int32(23)
	goto L18
L18:
	;
	goto L15
L19:
	;
	if v67 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_mdimmedsync[1]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v73*int32(48))+32))
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
	F_errmsg(m, int32(_a_F_mdimmedsync_0), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_mdimmedsync_1), int32(1470), int32(_a_F_mdimmedsync_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	F_FileClose(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(int32(1)) < base.Ui32(v44) {
		v44 = v50
		goto L10
	} else {
		goto L42
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v50 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v50
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v112
	goto L29
L31:
	;
	if v93 <= int32(0) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v93 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	F_pfree(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v112 = int32(0)
	goto L30
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_mdimmedsync[2]))
	v106 = F_MemoryContextAlloc(m, v105, v52)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v50 <= v93 {
		goto L29
	} else {
		goto L40
	}
L39:
	;
	v112 = v106
	goto L30
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v110 = F_repalloc(m, v109, v52)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v112 = v110
	goto L30
L42:
	;
	goto L11
}
func F_mdunlink(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	if l1 == int32(-1) {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v11
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v13
		F_mdunlinkfork(m, v5+int32(-16), int32(0), l2)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v20
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v22
			F_mdunlinkfork(m, v5+int32(-32), int32(1), l2)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v29
				v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v31
				F_mdunlinkfork(m, v5+int32(-48), int32(2), l2)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v39 = int32(3)
					v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v40
					v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = v42
					F_mdunlinkfork(m, v7, v39, l2)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		v39 = l1
		v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v40
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v42
		F_mdunlinkfork(m, v7, v39, l2)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			m.G0 = v7 - int32(-64)
			return
		}
	}
}
func F_mdunlinkfiletag(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = v7 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v10, v11, v12, v13, int32(-1), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L6
L3:
	;
	v139 = F_unlink(m, l1)
	mBase = m.M
	m.G0 = v7 + int32(80)
	return v139
L4:
	;
	v136 = F_strlen(m, v125)
	mBase = m.M
	goto L3
L6:
	;
	goto L7
L7:
	;
	v26 = int32(1023)
	if (l1^v10)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
	goto L4
L9:
	;
	v110 = v105
	v111 = v106
	v112 = v107
	goto L30
L10:
	;
	if v100 == int32(0) {
		v125 = v98
		v126 = v99
		goto L8
	} else {
		goto L29
	}
L11:
	;
	v98 = v10
	v99 = l1
	v100 = v26
	goto L10
L12:
	;
	goto L13
L13:
	;
	v30 = int32(0)
	if base.B2i32(v10&int32(3) == v30)|int32(0) == v30 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v66 == int32(0) {
		v125 = v63
		v126 = v64
		goto L8
	} else {
		goto L23
	}
L15:
	;
	v42 = v10
	v43 = l1
	v44 = v26
	goto L18
L16:
	;
	goto L17
L17:
	;
	v63 = v10
	v64 = l1
	v65 = v26
	v66 = int32(1)
	goto L14
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v46)
	if v46 == int32(0) {
		v105 = v42
		v106 = v43
		v107 = v44
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v63 = v57
	v64 = v51
	v65 = v53
	v66 = v55
	goto L14
L20:
	;
	v50 = int32(1)
	v51 = v43 + v50
	v53 = v44 - v50
	v54 = int32(0)
	v55 = base.B2i32(v53 != v54)
	v57 = v42 + v50
	if v57&int32(3) == v54 {
		v63 = v57
		v64 = v51
		v65 = v53
		v66 = v55
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v53 != 0 {
		v42 = v57
		v43 = v51
		v44 = v53
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if base.B2i32(v69 == int32(0))|base.B2i32(base.Ui32(v65) < base.Ui32(int32(4))) != 0 {
		v98 = v63
		v99 = v64
		v100 = v65
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v76 = v63
	v77 = v64
	v78 = v65
	goto L25
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v84 = int32(-2139062144)
	if (int32(16843008)-v81|v81)&v84 != v84 {
		v105 = v76
		v106 = v77
		v107 = v78
		goto L9
	} else {
		goto L27
	}
L26:
	;
	v98 = v92
	v99 = v90
	v100 = v94
	goto L10
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v81
	v89 = int32(4)
	v90 = v77 + v89
	v92 = v76 + v89
	v94 = v78 - v89
	if base.Ui32(int32(3)) < base.Ui32(v94) {
		v76 = v92
		v77 = v90
		v78 = v94
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v105 = v98
	v106 = v99
	v107 = v100
	goto L9
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v114)
	if v114 == int32(0) {
		v125 = v110
		v126 = v111
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v125 = v121
	v126 = v119
	goto L8
L32:
	;
	v118 = int32(1)
	v119 = v111 + v118
	v121 = v110 + v118
	v123 = v112 - v118
	if v123 != 0 {
		v110 = v121
		v111 = v119
		v112 = v123
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
}
func F_mdunlinkfork(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v10+int32(128), v14, v15, v16, v17, v2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = base.B2i32(v17 != int32(-1))
	if l2|(v21|v2) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v10 + int32(224)
	return
L4:
	;
	v119 = v10 + int32(200) | int32(4)
	v125 = int32(1)
	goto L29
L5:
	;
	v90 = F_do_truncate(m, v10+int32(128))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mdunlinkfork[0])))
	if v27&int32(1) == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v17 != int32(-1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v61 = v10 + int32(128)
	v62 = F_unlink(m, v61)
	mBase = m.M
	if int32(0) <= v62 {
		goto L4
	} else {
		goto L16
	}
L11:
	;
	v34 = F_do_truncate(m, v10+int32(128))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1]))
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v40
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v38
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+42)) = uint16(v2)
	v51 = F_RegisterSyncRequest(m, v10+int32(40), int32(2), int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1])) = v37
	if int32(0) <= v34 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	if v37 == int32(44) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1]))
	if v66 == int32(44) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v71 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v71 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1])) = v66
	goto L4
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v61
	F_errmsg(m, int32(_a_F_mdunlinkfork_0), v10+int32(32))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_mdunlinkfork_1), int32(401), int32(_a_F_mdunlinkfork_2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1]))
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v94
	v104 = int32(1)
	v106 = F_RegisterSyncRequest(m, v10+int32(40), v104, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1])) = v93
	if int32(0) <= v90 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v93 == int32(44) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L4
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(128)
	v135 = v10 + int32(40)
	v139 = F_pg_sprintf(m, v135, int32(_a_F_mdunlinkfork_3), v10+int32(16))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1]))
	if v174 == int32(44) {
		goto L3
	} else {
		goto L42
	}
L31:
	;
	if base.B2i32(v17 != int32(-1)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v143 = F_do_truncate(m, v135)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v169 = v10 + int32(40)
	v170 = F_unlink(m, v169)
	mBase = m.M
	if int32(0) <= v170 {
		v125 = v125 + int32(1)
		goto L29
	} else {
		goto L41
	}
L35:
	;
	if v143 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_mdunlinkfork[1]))
	if v148 == int32(44) {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+200)) = int64(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = v155
	*(*int64)(unsafe.Add(mBase, uint32(v10)+216)) = base.I64_extend_i32_u(v125)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+202)) = uint16(v2)
	v164 = F_RegisterSyncRequest(m, v10+int32(200), int32(2), int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L34
L41:
	;
	goto L30
L42:
	;
	v179 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v179 == int32(0) {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v169
	F_errmsg(m, int32(_a_F_mdunlinkfork_0), v10)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_mdunlinkfork_1), int32(460), int32(_a_F_mdunlinkfork_2))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L3
}
func F_merge_matching_partitions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = v14 + l3
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = v17 + l2
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = int32(2)
	v23 = v20 + l2<<(uint(v21)%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = v25 + l3<<(uint(v21)%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if int32(0) <= v24|v29 {
		if v29 == v24 {
			return v24
		} else {
			if v16|v19 != 0 {
				v128 = int32(-1)
				return v128
			} else {
				if base.Ui32(v24) < base.Ui32(v29) {
					v38 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v38)
					v41 = l3 << (uint(int32(2)) % 32)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v41+v42))) = v24
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v45+l3))) = uint8(v38)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v38)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v51+v41))) = v29
					return v24
				} else {
					v55 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v55)
					v58 = l2 << (uint(int32(2)) % 32)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v58+v59))) = v29
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v62+l2))) = uint8(v55)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v55)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v68+v58))) = v24
					v128 = v29
					return v128
				}
			}
		}
	} else {
		if v29&v24 == int32(-1) {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v74
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v78 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v76+l2))) = uint8(v78)
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v80+l3<<(uint(int32(2))%32)))) = v74
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*uint8)(unsafe.Add(mBase, uint32(v85+l3))) = uint8(v78)
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v89 + v78
			return v74
		} else {
			v96 = int32(0)
			if v19&int32(1)|base.B2i32(v24 < v96) == v96 {
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v24
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v104 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v102+l3))) = uint8(v104)
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v106+l2))) = uint8(v104)
				return v24
			} else {
				if v16&int32(1)|base.B2i32(v29 < int32(0)) != 0 {
					v128 = int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v29
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v120 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v118+l2))) = uint8(v120)
					v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v122+l3))) = uint8(v120)
					v128 = v29
				}
				return v128
			}
		}
	}
}
func F_mergeruns(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
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
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v565 int64
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v586 int32
	_ = v586
	var v587 int64
	_ = v587
	var v589 int64
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v675 int32
	_ = v675
	var v676 int64
	_ = v676
	var v678 int64
	_ = v678
	var v692 int32
	_ = v692
	var v708 int32
	_ = v708
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int64
	_ = v771
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_MemoryContextResetOnly(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v23
	goto L1
L4:
	;
	return
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v39 = F_GetMemoryChunkSpace(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v41 + base.I64_extend_i32_u(v39)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_pfree(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v50 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)) = uint8(v185)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v192 = F_MemoryContextAlloc(m, v189, v187<<(uint(int32(4))%32))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L28
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if base.Ui32(v53) <= base.Ui32(int32(2147483646)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+148)) = int64(0)
	goto L8
L12:
	;
	v59 = v53<<(uint(int32(10))%32) + int32(1024)
	v60 = F_palloc(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+148)) = int64(0)
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v59 + v60
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v66 - base.I64_extend_i32_s(v59)
	if v53 == int32(0) {
		v154 = v60
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(0)
	goto L8
L17:
	;
	v73 = v53 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v53) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v81 = int32(0)
	v82 = v60
	goto L21
L19:
	;
	v123 = v60
	goto L20
L20:
	;
	v134 = int32(0)
	v136 = v123
	goto L25
L21:
	;
	v92 = v82 - int32(-8192)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_mergeruns[0]))) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_mergeruns[1]))) = v82 + int32(_a_F_mergeruns_0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_mergeruns[2]))) = v82 + int32(_a_F_mergeruns_1)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_mergeruns[3]))) = v82 + int32(_a_F_mergeruns_2)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+3072)) = v82 + int32(_a_F_mergeruns_3)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+2048)) = v82 + int32(3072)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+1024)) = v82 + int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v82 + int32(1024)
	v116 = v81 + int32(8)
	if v116 != v53&int32(2147483640) {
		v81 = v116
		v82 = v92
		goto L21
	} else {
		goto L23
	}
L22:
	;
	if v73 == int32(0) {
		v154 = v92
		goto L16
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v123 = v92
	goto L20
L25:
	;
	v146 = v136 + int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v146
	v149 = v134 + int32(1)
	if v149 != v73 {
		v134 = v149
		v136 = v146
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v154 = v146
	goto L16
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v192
	v195 = F_GetMemoryChunkSpace(m, v192)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v199 = v197 - base.I64_extend_i32_u(v195)
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+160)) = uint32(v199)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v199 & int64(-4294967296)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mergeruns[4])))
	if v205 != int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v239 = base.B2i32(v233 == int32(0))
	goto L38
L31:
	;
	v210 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v210 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(base.Ui32(v214) >> (uint(int32(10)) % 32))
	F_errmsg_internal(m, int32(_a_F_mergeruns_4), v13+int32(-32))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_mergeruns_5), int32(2094), int32(_a_F_mergeruns_6))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L168
	}
L37:
	;
	m.G0 = v15 - int32(-64)
	return
L38:
	;
	if v239&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v738
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v740 != 0 {
		goto L152
	} else {
		goto L153
	}
L40:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v427 < v428 {
		goto L79
	} else {
		goto L80
	}
L41:
	;
	v254 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v254 < v255 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v259 = v254
	goto L45
L43:
	;
	goto L44
L44:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v298
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v300
	v304 = F_palloc0(m, v298<<(uint(int32(2))%32))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L50
	}
L45:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v259<<(uint(int32(2))%32))))
	F_LogicalTapeClose(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_pfree(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	v278 = v259 + int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v278 < v279 {
		v259 = v278
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L44
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v304
	v309 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+160)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v315 = base.I32_div_s(v310+v311-int32(1), v310)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v315 < v316 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v318 = v315
	goto L53
L52:
	;
	v318 = v316
	goto L53
L53:
	;
	v324 = base.I64_div_s(v309-base.I64_extend_i32_s(v318<<(uint(int32(13))%32)), base.I64_extend_i32_s(v310))
	v325 = int64(0)
	if v325 < v324 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v328 = v324
	goto L56
L55:
	;
	v328 = v325
	goto L56
L56:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mergeruns[4])))
	if v330 != int32(1) {
		v359 = v310
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if int32(0) < v359 {
		goto L64
	} else {
		goto L65
	}
L58:
	;
	v335 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v335 == int32(0) {
		v359 = v337
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v341 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-48)))) = v341
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(base.Ui64(v328) >> (uint(int64(10)) % 64))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v340
	F_errmsg_internal(m, int32(_a_F_mergeruns_7), v15)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_mergeruns_5), int32(2142), int32(_a_F_mergeruns_6))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v359 = v357
	goto L57
L64:
	;
	v365 = int32(0)
	goto L67
L65:
	;
	v390 = v359
	goto L66
L66:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v399&int32(1) != 0 {
		goto L40
	} else {
		goto L71
	}
L67:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376+v365<<(uint(int32(2))%32))))
	F_LogicalTapeRewindForRead(m, v380, base.I32_wrap_i64(v328))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	v390 = v385
	goto L66
L69:
	;
	v384 = v365 + int32(1)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v384 < v385 {
		v365 = v384
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v390 < v402 {
		goto L40
	} else {
		goto L72
	}
L72:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v404 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v405 != int32(-1) {
		goto L40
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v409 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+40)) = uint8(v409)
	F_beginmerge(m, l0)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(5)
	goto L37
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v453 + int32(1)
	F_beginmerge(m, l0)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L83
	}
L79:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v431 = F_LogicalTapeCreate(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v447 = base.I32_rem_s(v446, v427)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v445+v447<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v451
	v453 = v446
	goto L78
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v431
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v434+v435<<(uint(int32(2))%32)))) = v431
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v440 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v453 = v444
	goto L78
L83:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v459 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	goto L87
L85:
	;
	goto L86
L86:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(0)
	F_LogicalTapeWrite(m, v723, v13+int32(-16), int32(4))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L4
	} else {
		goto L147
	}
L87:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v474+v476<<(uint(int32(2))%32))))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v482].(func(*base.Module, int32, int32, int32))(m, l0, v481, v475)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L89
	}
L88:
	;
	goto L86
L89:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v486 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v499 = v13 + int32(-16)
	v501 = F_LogicalTapeRead(m, v480, v499, int32(4))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L96
	}
L91:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v486) < base.Ui32(v489) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_pfree(m, v486)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v491) <= base.Ui32(v486) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = v493
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v486
	goto L90
L95:
	;
	goto L90
L96:
	;
	if v501 != int32(4) {
		goto L36
	} else {
		goto L97
	}
L97:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v505 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v708 {
		goto L87
	} else {
		goto L146
	}
L99:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v506].(func(*base.Module, int32, int32, int32, int32))(m, l0, v499, v480, v505)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v593 = v591 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v593
	if int32(0) < v593 {
		goto L123
	} else {
		goto L124
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v476
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_mergeruns[5]))
	if v512 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v515 = int32(0)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v519) < base.Ui32(int32(2)) {
		v573 = v515
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	v586 = v510 + v573<<(uint(int32(4))%32)
	v587 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v586)+8)) = v587
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v586))) = v589
	goto L98
L108:
	;
	v524 = int32(1)
	v525 = v515
	v528 = v515
	goto L109
L109:
	;
	v537 = v528 + int32(2)
	if base.Ui32(v519) <= base.Ui32(v537) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v573 = v551
	goto L107
L111:
	;
	v551 = v524
	goto L113
L112:
	;
	v539 = int32(4)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v546 = m.T0[v545].(func(*base.Module, int32, int32, int32) int32)(m, v510+v524<<(uint(v539)%32), v510+v537<<(uint(v539)%32), l0)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L114
	}
L113:
	;
	v554 = v510 + v551<<(uint(int32(4))%32)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v556 = m.T0[v555].(func(*base.Module, int32, int32, int32) int32)(m, v13+int32(-16), v554, l0)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L118
	}
L114:
	;
	if int32(0) < v546 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v550 = v537
	goto L117
L116:
	;
	v550 = v524
	goto L117
L117:
	;
	v551 = v550
	goto L113
L118:
	;
	if v556 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v573 = v525
	goto L107
L120:
	;
	goto L121
L121:
	;
	v562 = v510 + v525<<(uint(int32(4))%32)
	v563 = *(*int64)(unsafe.Add(mBase, uint32(v554)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+8)) = v563
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v554)))
	*(*int64)(unsafe.Add(mBase, uint32(v562))) = v565
	v567 = int32(1)
	v568 = v551 << (uint(v567) % 32)
	v570 = v568 | v567
	if base.Ui32(v570) < base.Ui32(v519) {
		v524 = v570
		v525 = v551
		v528 = v568
		goto L109
	} else {
		goto L122
	}
L122:
	;
	goto L110
L123:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_mergeruns[5]))
	if v601 != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v692 - int32(1)
	goto L98
L126:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	v605 = v593
	goto L128
L128:
	;
	v606 = v593<<(uint(int32(4))%32) + v597
	v607 = int32(0)
	if base.Ui32(v605) < base.Ui32(int32(2)) {
		v662 = v607
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v605 = v604
	goto L128
L130:
	;
	v675 = v597 + v662<<(uint(int32(4))%32)
	v676 = *(*int64)(unsafe.Add(mBase, uint32(v606)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v675)+8)) = v676
	v678 = *(*int64)(unsafe.Add(mBase, uint32(v606)))
	*(*int64)(unsafe.Add(mBase, uint32(v675))) = v678
	goto L125
L131:
	;
	v615 = int32(1)
	v619 = v607
	v621 = v607
	goto L132
L132:
	;
	v626 = v621 + int32(2)
	if base.Ui32(v605) <= base.Ui32(v626) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v662 = v640
	goto L130
L134:
	;
	v640 = v615
	goto L136
L135:
	;
	v628 = int32(4)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v635 = m.T0[v634].(func(*base.Module, int32, int32, int32) int32)(m, v597+v615<<(uint(v628)%32), v597+v626<<(uint(v628)%32), l0)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L137
	}
L136:
	;
	v643 = v597 + v640<<(uint(int32(4))%32)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v645 = m.T0[v644].(func(*base.Module, int32, int32, int32) int32)(m, v606, v643, l0)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L4
	} else {
		goto L141
	}
L137:
	;
	if int32(0) < v635 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v639 = v626
	goto L140
L139:
	;
	v639 = v615
	goto L140
L140:
	;
	v640 = v639
	goto L136
L141:
	;
	if v645 <= int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v662 = v619
	goto L130
L143:
	;
	goto L144
L144:
	;
	v651 = v597 + v619<<(uint(int32(4))%32)
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v643)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+8)) = v652
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v643)))
	*(*int64)(unsafe.Add(mBase, uint32(v651))) = v654
	v656 = int32(1)
	v657 = v640 << (uint(v656) % 32)
	v659 = v657 | v656
	if base.Ui32(v659) < base.Ui32(v605) {
		v615 = v659
		v619 = v640
		v621 = v657
		goto L132
	} else {
		goto L145
	}
L145:
	;
	goto L133
L146:
	;
	goto L88
L147:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v733 = base.B2i32(v731 == int32(0))
	if v731 != 0 {
		v239 = v733
		goto L38
	} else {
		goto L148
	}
L148:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if int32(1) < v734 {
		v239 = v733
		goto L38
	} else {
		goto L149
	}
L149:
	;
	goto L39
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(4)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v781 <= int32(0) {
		goto L37
	} else {
		goto L163
	}
L151:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_pfree(m, v747)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L4
	} else {
		goto L157
	}
L152:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v741 != int32(-1) {
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	F_LogicalTapeFreeze(m, v738, int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L156
	}
L155:
	;
	goto L154
L156:
	;
	goto L150
L157:
	;
	v750 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v750
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	F_LogicalTapeFreeze(m, v754, v13+int32(-16))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = int32(1)
	if v759 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_s_lock(m, v740, int32(_a_F_mergeruns_5), int32(3034), int32(_a_F_mergeruns_8))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v771 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v740+v767<<(uint(int32(3))%32))+72)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = int32(0)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v740)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v740)+8)) = v775 + int32(1)
	goto L150
L162:
	;
	goto L161
L163:
	;
	v786 = int32(0)
	goto L164
L164:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v797+v786<<(uint(int32(2))%32))))
	F_LogicalTapeClose(m, v801)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	goto L37
L166:
	;
	v805 = v786 + int32(1)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v805 < v806 {
		v786 = v805
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	F_errmsg_internal(m, int32(_a_F_mergeruns_9), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_mergeruns_5), int32(2862), int32(_a_F_mergeruns_10))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_message_level_is_interesting(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v4 = int32(1)
	if int32(20) < l0 {
		v38 = v4
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[0]))
		if base.Ui32(l0-int32(15)) <= base.Ui32(int32(1)) {
			if int32(22) <= v8 {
				v21 = int32(0)
				if l0 == int32(16) {
					v38 = v21
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[1]))
					if v25 != int32(2) {
						v38 = v21
					} else {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_message_level_is_interesting[2])))
						if v29&int32(1) != 0 {
							v38 = v21
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[3]))
							v38 = base.B2i32(l0 == int32(17)) | base.B2i32(v35 <= l0)
						}
					}
				}
			} else {
				v38 = v4
			}
		} else {
			if base.B2i32(l0 == int32(20))|base.B2i32(v8 == int32(15)) != 0 {
				v21 = int32(0)
				if l0 == int32(16) {
					v38 = v21
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[1]))
					if v25 != int32(2) {
						v38 = v21
					} else {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_message_level_is_interesting[2])))
						if v29&int32(1) != 0 {
							v38 = v21
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[3]))
							v38 = base.B2i32(l0 == int32(17)) | base.B2i32(v35 <= l0)
						}
					}
				}
			} else {
				if v8 <= l0 {
					v38 = v4
				} else {
					v21 = int32(0)
					if l0 == int32(16) {
						v38 = v21
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[1]))
						if v25 != int32(2) {
							v38 = v21
						} else {
							v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_message_level_is_interesting[2])))
							if v29&int32(1) != 0 {
								v38 = v21
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, _c_F_message_level_is_interesting[3]))
								v38 = base.B2i32(l0 == int32(17)) | base.B2i32(v35 <= l0)
							}
						}
					}
				}
			}
		}
	}
	return v38
}
func F_minmax_multi_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l1 - int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0+v14<<(uint(int32(2))%32))+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v21 = v19 + int32(28)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if l2 != v22 {
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v21))) = l2
	} else {
	}
	v39 = v19 + l3*int32(28)
	v41 = v39 + int32(4)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v42 == int32(0) {
		v45 = int32(4)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+208))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v14<<(uint(int32(2))%32))))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
		v61 = v52 + v53<<(uint(v45)%32) + v14*int32(100) + int32(88)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
		v64 = F_SearchSysCache4(m, v45, v51, v62, l2, base.I32_extend16_s(l3))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int32(0)
		} else {
			if v64 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v92
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_errmsg_internal(m, int32(_a_F_minmax_multi_get_strategy_procinfo_0), v11)
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_minmax_multi_get_strategy_procinfo_1), int32(2938), int32(_a_F_minmax_multi_get_strategy_procinfo_2))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v72 = F_SysCacheGetAttrNotNull(m, int32(4), v64, int32(7))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v64)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = F_get_opcode(m, v72)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_fmgr_info_cxt(m, v76, v41, v78)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return v41
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v41
	}
}
func F_minmax_qp_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+124))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+76))
	v10 = F_make_pathkeys_for_sortclauses(m, l0, v8, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v10
		return
	}
}
func F_movedb_failure_callback(m *base.Module, l0 int32, l1 int32) {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = F_GetDatabasePath(m, v3, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = F_rmtree(m, v5)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_pfree(m, v5)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_moveouts(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L5
L3:
	;
	goto L4
L4:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v104 = int32(32)
	if base.B2i32(int32(4) <= v101)&(base.B2i32(v104 < v8)|base.B2i32(base.Ui32(v104) < base.Ui32(v101))) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_createarc(m, l0, v21, v22, l2, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v32 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v67 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v37 = v35 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v37))|base.B2i32(int32(1)<<(uint(v37)%32)&int32(_a_F_moveouts_0) == int32(0)) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v47 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v48 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v32*int32(24))+12)) = v56
	v60 = v56
	goto L15
L17:
	;
	goto L18
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v58
	v60 = v58
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+36)) = v48
	goto L21
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
	goto L11
L22:
	;
	if v66 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v66
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v66
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v67
	goto L28
L27:
	;
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v73 - int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v78 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v77 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v77
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v77
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = v78
	goto L35
L34:
	;
	goto L35
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v84 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v91 = v18 + int32(8)
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v92
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	goto L10
L36:
	;
	goto L39
L37:
	;
	goto L38
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_moveouts[0]))
	if v201 != 0 {
		goto L69
	} else {
		goto L70
	}
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	F_cparc(m, l0, v119, l2, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+4)))
	if v131 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L39
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	if v166 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v136 = v134 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v136))|base.B2i32(int32(1)<<(uint(v136)%32)&int32(_a_F_moveouts_0) == int32(0)) != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v146 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v119)+36))
	if v147 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v159 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v151+v131*int32(24))+12)) = v155
	v159 = v155
	goto L48
L50:
	;
	goto L51
L51:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+32)) = v157
	v159 = v157
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+36)) = v147
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = int64(0)
	goto L44
L55:
	;
	if v165 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+20)) = v165
	goto L55
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+16)) = v165
	goto L55
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v166
	goto L61
L60:
	;
	goto L61
L61:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = v172 - int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	if v177 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v176 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+16)) = v176
	goto L62
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+24)) = v176
	goto L62
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+28)) = v177
	goto L68
L67:
	;
	goto L68
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v183 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(0)
	v190 = v119 + int32(8)
	v191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+16)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v191
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v119
	goto L43
L69:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_sortouts(m, l0, l1)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_sortouts(m, l0, l2)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	if v209 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v210 == int32(0) {
		v351 = v210
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v351 == int32(0) {
		goto L1
	} else {
		goto L128
	}
L77:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v213 == int32(0) {
		v351 = v210
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v217 = v210
	v220 = v213
	goto L79
L79:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v224 < v226 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v351 = v343
	goto L76
L81:
	;
	if v343 == int32(0) {
		v351 = v343
		goto L76
	} else {
		goto L126
	}
L82:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v343 = v217
	v345 = v342
	goto L81
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	if v318 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L84:
	;
	if v226 < v224 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217)+4)))
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220)+4)))
	if v229 < v230 {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	if v230 < v229 {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v233 < v234 {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	if v234 < v233 {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v245 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v343 = v238
	v345 = v237
	goto L81
L91:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	if v280 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L92:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v250 = v248 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v250))|base.B2i32(int32(1)<<(uint(v250)%32)&int32(_a_F_moveouts_0) == int32(0)) != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v260 != 0 {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v217)+36))
	if v261 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v273 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v217)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v245*int32(24))+12)) = v269
	v273 = v269
	goto L95
L97:
	;
	goto L98
L98:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v217)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+32)) = v271
	v273 = v271
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+36)) = v261
	goto L101
L100:
	;
	goto L101
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v217)+32)) = int64(0)
	goto L91
L102:
	;
	if v279 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+20)) = v279
	goto L102
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v279
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v280
	goto L108
L107:
	;
	goto L108
L108:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v286 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	if v291 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v290 != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v290
	goto L109
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v290
	goto L109
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v291
	goto L115
L114:
	;
	goto L115
L115:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v297 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(0)
	v304 = v217 + int32(8)
	v305 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v304)+16)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304)+8)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v217
	goto L90
L116:
	;
	if v317 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+20)) = v317
	goto L116
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+16)) = v317
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+20)) = v318
	goto L122
L121:
	;
	goto L122
L122:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+12)) = v324 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = l2
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = v329
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v333 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+20)) = v217
	goto L125
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v217
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v336 + int32(1)
	v343 = v317
	v345 = v220
	goto L81
L126:
	;
	if v345 != 0 {
		v217 = v343
		v220 = v345
		goto L79
	} else {
		goto L127
	}
L127:
	;
	goto L80
L128:
	;
	v360 = v351
	goto L129
L129:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v360)+20))
	if v368 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L1
L131:
	;
	if v367 != 0 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = v367
	goto L131
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+16)) = v367
	goto L131
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+20)) = v368
	goto L137
L136:
	;
	goto L137
L137:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v374 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+8)) = l2
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = v379
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v383 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = v360
	goto L140
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v360
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v386 + int32(1)
	if v367 != 0 {
		v360 = v367
		goto L129
	} else {
		goto L141
	}
L141:
	;
	goto L130
}
func F_mq_putmessage_noblock(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_errmsg_internal(m, int32(_a_F_mq_putmessage_noblock_0), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errfinish(m, int32(_a_F_mq_putmessage_noblock_1), int32(208), int32(_a_F_mq_putmessage_noblock_2))
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_mxid_age(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v2 != 0 {
			v9 = v3 - v2
		} else {
			v9 = int32(2147483647)
		}
		return v9
	}
}
