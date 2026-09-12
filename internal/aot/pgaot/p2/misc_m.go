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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
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
	v98 = v13
	goto L6
L6:
	;
	return v98
L7:
	;
	return int32(0)
L8:
	;
	v25 = int32(4549024)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v30
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v33 <= v32 {
		v89 = v32
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
	v98 = v89
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
		v89 = v54
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
	v61 = v54
	v63 = int32(1)
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v69 = v66 + v63*int32(56)
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
	v89 = v82
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v74
	v77 = int32(1)
	if base.Ui32(v61) <= base.Ui32(v77) {
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
	v80 = v61
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
	v82 = v61
	goto L26
L26:
	;
	v84 = v63 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v84 < v85 {
		v61 = v82
		v63 = v84
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
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
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
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l0 == v4 {
		v508 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L12
	} else {
		goto L97
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v508
L3:
	;
	switch l2 - int32(2) {
	case 0:
		goto L11
	case 1:
		goto L7
	case 2:
		goto L10
	case 3:
		goto L9
	default:
		goto L8
	}
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	if v63 < int32(0) {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v53 = v52 | v51
	if v53&int32(1) != 0 {
		goto L4
	} else {
		goto L16
	}
L6:
	;
	v49 = v46
	v50 = v47
	v51 = v4
	v52 = int32(0)
	goto L5
L7:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	v46 = v4
	v47 = v45
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v49 = v27
	v50 = v26
	v51 = v28
	v52 = v29
	goto L5
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	v46 = v24
	v47 = int32(0)
	goto L6
L11:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v49 = v4
	v50 = v4
	v51 = v22
	v52 = v23
	goto L5
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	F_errmsg_internal(m, int32(505005), v16)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(515472), int32(4997), int32(368867))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	if v50&int32(1) != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v49&int32(1) == int32(0) {
		v508 = v4
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L4
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	if v63 < v67 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v143 = int32(0)
	if v50&int32(1) == v143 {
		v227 = v143
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v70 = v63 + int32(1)
	if v67 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v93
	*(*int32)(unsafe.Add(mBase, _consts[415])) = v95
	if v93 <= v67 {
		goto L20
	} else {
		goto L34
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v75 = int32(8)
	if v70 <= v75 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v86 = v67 << (uint(int32(1)) % 32)
	if v86 < v70 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v78 = v75
	goto L28
L27:
	;
	v78 = v70
	goto L28
L28:
	;
	v81 = F_MemoryContextAlloc(m, v74, v78*int32(20))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v93 = v78
	v95 = v81
	goto L22
L30:
	;
	v88 = v70
	goto L32
L31:
	;
	v88 = v86
	goto L32
L32:
	;
	v91 = F_repalloc(m, v84, v88*int32(20))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v93 = v88
	v95 = v91
	goto L22
L34:
	;
	v104 = v67
	goto L35
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v118 = v115 + v104*int32(20)
	v119 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v119
	v126 = v104 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	if v126 < v128 {
		v104 = v126
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L20
L37:
	;
	goto L36
L38:
	;
	if v53&int32(1) == int32(0) {
		v318 = v143
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v152 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v157 = v150 + v152*int32(20) + int32(16)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v158 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v207 = int32(4549024)
	v208 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v211 = *(*int32)(unsafe.Add(mBase, _consts[169]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v211
	v214 = F_palloc0(m, int32(36))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L50
	}
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v161 <= int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v169 = int32(0)
	goto L43
L43:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v164+v169<<(uint(int32(2))%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v183 != l1 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L40
L45:
	;
	v192 = v169 + int32(1)
	if v161 != v192 {
		v169 = v192
		goto L43
	} else {
		goto L49
	}
L46:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v185 != int32(3) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
	if v188 != int32(1) {
		v227 = v182
		goto L38
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = l1
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v220 = F_lappend(m, v219, v214)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v220
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v208
	v227 = v214
	goto L38
L52:
	;
	v331 = int32(0)
	if v49&int32(1) == v331 {
		v416 = v331
		goto L66
	} else {
		goto L67
	}
L53:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v245 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v250 = v243 + v245*int32(20) + int32(16)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v300 = int32(4549024)
	v301 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v304 = *(*int32)(unsafe.Add(mBase, _consts[169]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v304
	v307 = F_palloc0(m, int32(36))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L64
	}
L55:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v254 <= int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v262 = int32(0)
	goto L57
L57:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v257+v262<<(uint(int32(2))%32))))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v276 != l1 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L54
L59:
	;
	v285 = v262 + int32(1)
	if v254 != v285 {
		v262 = v285
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v278 != int32(2) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+8)))
	if v281 != int32(1) {
		v318 = v275
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L58
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v307))) = l1
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v313 = F_lappend(m, v312, v307)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v313
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v301
	v318 = v307
	goto L52
L66:
	;
	v425 = int32(4549024)
	v426 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v429 = *(*int32)(unsafe.Add(mBase, _consts[169]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v429
	v431 = int32(4549076)
	v432 = *(*int32)(unsafe.Add(mBase, _consts[261]))
	v435 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	*(*int32)(unsafe.Add(mBase, _consts[261])) = v435
	if v52&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L67:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v339 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v344 = v337 + v339*int32(20) + int32(16)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v345 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v394 = int32(4549024)
	v395 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v398 = *(*int32)(unsafe.Add(mBase, _consts[169]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v398
	v401 = F_palloc0(m, int32(36))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L12
	} else {
		goto L78
	}
L69:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v348 <= int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	v356 = int32(0)
	goto L71
L71:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v351+v356<<(uint(int32(2))%32))))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v370 != l1 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L68
L73:
	;
	v379 = v356 + int32(1)
	if v348 != v379 {
		v356 = v379
		goto L71
	} else {
		goto L77
	}
L74:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v372 != int32(4) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+8)))
	if v375 != int32(1) {
		v416 = v369
		goto L66
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	goto L72
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v401)+4)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v401))) = l1
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v407 = F_lappend(m, v406, v401)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v407
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v395
	v416 = v401
	goto L66
L80:
	;
	if v51&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v318)+24))
	if v441 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v442 = int32(0)
	v445 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v446 = F_tuplestore_begin_heap(m, v442, v442, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L12
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+24)) = v446
	goto L80
L84:
	;
	if v49&int32(1) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v318)+28))
	if v453 != 0 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v454 = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v458 = F_tuplestore_begin_heap(m, v454, v454, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+28)) = v458
	goto L84
L88:
	;
	v474 = v50 & int32(1)
	if v474 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v416)+24))
	if v465 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v466 = int32(0)
	v469 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v470 = F_tuplestore_begin_heap(m, v466, v466, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L12
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+24)) = v470
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v426
	*(*int32)(unsafe.Add(mBase, _consts[261])) = v432
	v490 = F_palloc0(m, int32(20))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L12
	} else {
		goto L96
	}
L93:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	if v477 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v478 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v482 = F_tuplestore_begin_heap(m, v478, v478, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+28)) = v482
	goto L92
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+16)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v490)+12)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v490)+8)) = v227
	*(*uint8)(unsafe.Add(mBase, uint32(v490)+3)) = uint8(v474)
	v496 = int32(1)
	v497 = v51 & v496
	*(*uint8)(unsafe.Add(mBase, uint32(v490)+2)) = uint8(v497)
	v500 = v52 & v496
	*(*uint8)(unsafe.Add(mBase, uint32(v490)+1)) = uint8(v500)
	v503 = v49 & v496
	*(*uint8)(unsafe.Add(mBase, uint32(v490))) = uint8(v503)
	v508 = v490
	goto L2
L97:
	;
	F_errmsg_internal(m, int32(16881), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(515472), int32(5007), int32(368867))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	v15 = m.G0
	v17 = v15 - int32(224)
	m.G0 = v17
	v20 = int32(base.Ui32(l2) >> (uint(int32(17)) % 32))
	v23 = l0 + l1<<(uint(int32(2))%32)
	v25 = v23 + int32(40)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if base.Ui32(v20) < base.Ui32(v26) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(224)
	return v222
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v222 = v28 + v20<<(uint(int32(3))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = int32(0)
	if base.Ui32(int32(31)) < base.Ui32(l4) {
		v222 = v32
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if base.Ui32(v20) < base.Ui32(v51) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56))
	v51 = v26
	v52 = v38 + v26<<(uint(int32(3))%32) - int32(8)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v44 = F_mdopenfork(m, l0, l1, l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v44 == int32(0) {
		v222 = v32
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v51 = v50
	v52 = v44
	goto L6
L13:
	;
	v222 = v52
	goto L1
L14:
	;
	goto L15
L15:
	;
	v65 = v52
	v67 = v51
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v75 = F_FileSize(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L20
	}
L17:
	;
	v222 = v178
	goto L1
L18:
	;
	v178 = F__mdfd_openseg(m, l0, l1, v67, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L51
	}
L19:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L48
	}
L20:
	;
	if int64(0) <= v75 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = base.I32_wrap_i64(int64(base.Ui64(v75) >> (uint(int64(13)) % 64)))
	if base.Ui32(int32(131073)) <= base.Ui32(v81) {
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
	v144 = m.ExcPending
	if v144 != 0 {
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
	if v81 == int32(131072) {
		v177 = int32(0)
		goto L18
	} else {
		goto L34
	}
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
	if base.B2i32(l4&int32(8) != int32(0))&v87 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v91 = int32(64)
	if v81 == int32(131072) {
		v177 = v91
		goto L18
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v101 = F_palloc_aligned(m, int32(8192), int32(4096), int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	F_mdextend(m, l0, l1, v67<<(uint(int32(17))%32)-int32(1), v101, l3)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	F_pfree(m, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v177 = v91
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
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(44)
	v222 = int32(0)
	goto L1
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	F__mdfd_segpath(m, v17+int32(141), l0, l1, v67)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v17 + int32(141)
	F_errmsg(m, int32(162002), v17+int32(32))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(520146), int32(1848), int32(351840))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v149 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v147*int32(48))+32))
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v153
	F_errmsg(m, int32(310993), v17)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(520146), int32(1882), int32(161991))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	F_errmsg_internal(m, int32(351627), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(520146), int32(1794), int32(351840))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	if v178 == int32(0) {
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
	if v67 != v20 {
		v65 = v178
		v67 = v67 + int32(1)
		goto L16
	} else {
		goto L64
	}
L55:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v186 == int32(44) {
		v222 = int32(0)
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
	v193 = m.ExcPending
	if v193 != 0 {
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
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F__mdfd_segpath(m, v17+int32(58), l0, l1, v67)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v17 + int32(58)
	F_errmsg(m, int32(307912), v17+int32(16))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(520146), int32(1862), int32(351840))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	v7 = *(*int32)(unsafe.Add(mBase, _consts[143]))
	if v7 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[72]))
		v17 = F_AllocSetContextCreateInternal(m, v12, int32(67064), int32(0), int32(1024), int32(8192))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[143])) = v17
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
					v31 = F__emscripten_memcpy_bulkmem(m, v30, l2, v22)
					mBase = m.M
					v32 = v31
				} else {
					v32 = v30
				}
				F_pg_qsort(m, v32, l1, int32(8), int32(291))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v38 = v25 + int32(8)
					v40 = *(*int32)(unsafe.Add(mBase, _consts[144]))
					if v40 == int32(0) {
						v43 = int32(4146096)
						*(*int32)(unsafe.Add(mBase, _consts[145])) = v43
						*(*int32)(unsafe.Add(mBase, _consts[146])) = int32(0)
						v50 = v43
					} else {
						v50 = v40
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(4146096)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = v38
					*(*int32)(unsafe.Add(mBase, _consts[144])) = v38
					v57 = int32(4146104)
					v59 = *(*int32)(unsafe.Add(mBase, _consts[146]))
					v61 = v59 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[146])) = v61
					if base.Ui32(int32(257)) <= base.Ui32(v61) {
						v66 = *(*int32)(unsafe.Add(mBase, _consts[145]))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
						v72 = int32(4146104)
						v74 = *(*int32)(unsafe.Add(mBase, _consts[146]))
						*(*int32)(unsafe.Add(mBase, _consts[146])) = v74 - int32(1)
						F_pfree(m, v66-int32(8))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
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
				v31 = F__emscripten_memcpy_bulkmem(m, v30, l2, v22)
				mBase = m.M
				v32 = v31
			} else {
				v32 = v30
			}
			F_pg_qsort(m, v32, l1, int32(8), int32(291))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v38 = v25 + int32(8)
				v40 = *(*int32)(unsafe.Add(mBase, _consts[144]))
				if v40 == int32(0) {
					v43 = int32(4146096)
					*(*int32)(unsafe.Add(mBase, _consts[145])) = v43
					*(*int32)(unsafe.Add(mBase, _consts[146])) = int32(0)
					v50 = v43
				} else {
					v50 = v40
				}
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(4146096)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v50
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v38
				*(*int32)(unsafe.Add(mBase, _consts[144])) = v38
				v57 = int32(4146104)
				v59 = *(*int32)(unsafe.Add(mBase, _consts[146]))
				v61 = v59 + int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[146])) = v61
				if base.Ui32(int32(257)) <= base.Ui32(v61) {
					v66 = *(*int32)(unsafe.Add(mBase, _consts[145]))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
					*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
					v72 = int32(4146104)
					v74 = *(*int32)(unsafe.Add(mBase, _consts[146]))
					*(*int32)(unsafe.Add(mBase, _consts[146])) = v74 - int32(1)
					F_pfree(m, v66-int32(8))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = F_colname_is_unique(m, l0, l1, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = F_strlen(m, l0)
	mBase = m.M
	v23 = F_palloc(m, v20+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v84 = l0
	goto L5
L5:
	;
	m.G0 = v12 + int32(32)
	return v84
L6:
	;
	v28 = v20
	v31 = int32(0)
	goto L7
L7:
	;
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v84 = v35
	goto L5
L9:
	;
	v37 = v31 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v37
	v43 = F_pg_sprintf(m, v28+v35, int32(483962), v12+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	v34 = F__emscripten_memcpy_bulkmem(m, v23, l0, v28)
	mBase = m.M
	v35 = v34
	goto L12
L11:
	;
	v35 = v23
	goto L12
L12:
	;
	goto L9
L13:
	;
	v45 = F_strlen(m, v35)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v45) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = v28
	goto L17
L15:
	;
	v74 = v28
	goto L16
L16:
	;
	v80 = F_colname_is_unique(m, v35, l1, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L17:
	;
	v59 = F_pg_mbcliplen(m, l0, v51, v51-int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v74 = v59
	goto L16
L19:
	;
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v37
	v66 = F_pg_sprintf(m, v59+v62, int32(483962), v12)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v61 = F__emscripten_memcpy_bulkmem(m, v35, l0, v59)
	mBase = m.M
	v62 = v61
	goto L23
L22:
	;
	v62 = v35
	goto L23
L23:
	;
	goto L20
L24:
	;
	v68 = F_strlen(m, v62)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v68) {
		v51 = v59
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	if v80 == int32(0) {
		v28 = v74
		v31 = v37
		goto L7
	} else {
		goto L27
	}
L27:
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v81 int32
	_ = v81
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
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
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
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
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
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
	v24 = v2
	v26 = v2
	v27 = v2
	v29 = v2
	goto L4
L3:
	;
	m.G0 = v15 + int32(128)
	return v576
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = v33
	v39 = v26
	v40 = v27
	v42 = v29
	goto L6
L5:
	;
	v540 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v541 == v540 {
		v551 = v540
		goto L127
	} else {
		goto L128
	}
L6:
	;
	v46 = F_pg_mblen_cstr(m, v35)
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
	v535 = v46 + v530
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v535
	v35 = v535
	v39 = v532
	v40 = v533
	v42 = v534
	goto L6
L11:
	;
	v497 = F_t_isalnum_cstr(m, v51)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L116
	}
L12:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v261 = F_t_isalnum_cstr(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L63
	}
L13:
	;
	if v24 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L47
	}
L15:
	;
	v172 = v24
	goto L43
L16:
	;
	v159 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v110 + v159
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v162 - v159
	if int32(0) < v162 {
		goto L13
	} else {
		goto L42
	}
L17:
	;
	if v24 == int32(32) {
		goto L14
	} else {
		goto L41
	}
L18:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v111 == int32(32) {
		v530 = v110
		v532 = v39
		v533 = v40
		v534 = v42
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
	v148 = int32(33)
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
		v576 = v55
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v65 = int32(0)
	if v24 == v65 {
		v24 = v65
		v26 = v39
		v27 = v40
		v29 = v42
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v72 = v24
	goto L25
L25:
	;
	v81 = v72 - int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15+v81<<(uint(int32(2))%32))))
	switch v85 - int32(33) {
	case 0, 5:
		goto L27
	default:
		v24 = v72
		v26 = v39
		v27 = v40
		v29 = v42
		goto L4
	}
L26:
	;
	v24 = int32(0)
	v26 = v39
	v27 = v40
	v29 = v42
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v102 + int32(1)
	if v81 != 0 {
		v72 = v81
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	switch v111 - int32(38) {
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
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v146 != 0 {
		goto L8
	} else {
		goto L39
	}
L32:
	;
	v120 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v110 + v120
	if v24 == int32(0) {
		v148 = v122
		goto L17
	} else {
		goto L36
	}
L33:
	;
	if v111 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v111 != int32(124) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v122 != int32(124) {
		v148 = v122
		goto L17
	} else {
		goto L37
	}
L37:
	;
	v131 = F_palloc(m, int32(20))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v133 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v131)+12)) = uint16(v133)
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = int64(532575944707)
	*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v133
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v142 + int32(1)
	v26 = v39
	v27 = v40
	v29 = v42
	goto L4
L39:
	;
	if v24 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v576 = int32(0)
	goto L3
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v24<<(uint(int32(2))%32)))) = v148
	v24 = v24 + int32(1)
	v26 = v39
	v27 = v40
	v29 = v42
	goto L4
L42:
	;
	goto L8
L43:
	;
	v181 = v172 - int32(1)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v15+v181<<(uint(int32(2))%32))))
	v188 = F_palloc(m, int32(20))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v576 = int32(0)
	goto L3
L45:
	;
	v190 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v188)+12)) = uint16(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v190
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v188
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v200 + int32(1)
	if v181 != 0 {
		v172 = v181
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_errmsg_internal(m, int32(86678), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(516089), int32(252), int32(313572))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v576 = int32(0)
	goto L3
L51:
	;
	goto L52
L52:
	;
	v228 = v24
	goto L53
L53:
	;
	v237 = v228 - int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15+v237<<(uint(int32(2))%32))))
	v244 = F_palloc(m, int32(20))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v576 = int32(0)
	goto L3
L55:
	;
	v246 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v244)+12)) = uint16(v246)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+4)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v246
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v244
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v256 + int32(1)
	if v237 != 0 {
		v228 = v237
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	if v39 <= int32(65535) {
		goto L73
	} else {
		goto L74
	}
L58:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v530 = v303
	v532 = v46 + v39
	v533 = int32(0)
	v534 = v42
	goto L10
L59:
	;
	v530 = v263
	v532 = v39
	v533 = v40 | int32(1)
	v534 = v42
	goto L10
L60:
	;
	v530 = v263
	v532 = v39
	v533 = v40 | int32(2)
	v534 = v42
	goto L10
L61:
	;
	v530 = v263
	v532 = v39
	v533 = v40 | int32(4)
	v534 = v42
	goto L10
L62:
	;
	if v40&int32(65535) == int32(0) {
		goto L58
	} else {
		goto L67
	}
L63:
	;
	if v261 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	switch v264 - int32(37) {
	case 0:
		goto L61
	case 1, 2, 3, 4, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26:
		goto L57
	case 5:
		goto L59
	case 8:
		goto L62
	case 27:
		goto L60
	default:
		goto L65
	}
L65:
	;
	if v264 != int32(95) {
		goto L57
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v276 = F_errsave_start(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v276 == int32(0) {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(221451), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errsave_finish(m, v275, int32(516089), int32(102), int32(15957))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L8
L73:
	;
	v309 = F_ltree_crc32_sz(m, v42, v39)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v475 = int32(1)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v477 = F_errsave_start(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L111
	}
L76:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v314 = F_palloc(m, int32(20))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v314)+12)) = uint16(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v314)+4)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = int32(2)
	v320 = v312 - v311
	if int32(65536) <= v320 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v323 = int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v325 = F_errsave_start(m, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if int32(256) <= v39 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	if v325 == int32(0) {
		v576 = v323
		goto L3
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(351680), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errsave_finish(m, v324, int32(516089), int32(165), int32(15506))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v576 = v323
	goto L3
L86:
	;
	v347 = int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v349 = F_errsave_start(m, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v314)+10)) = uint16(v39)
	*(*uint16)(unsafe.Add(mBase, uint32(v314)+8)) = uint16(v320)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v314)+16)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v314
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v375 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v374 + v375
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v380 = v378 - v379
	v382 = v39 + v375
	v383 = v380 + v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v384 <= v383 {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	if v349 == int32(0) {
		v576 = v347
		goto L3
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(341862), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errsave_finish(m, v348, int32(516089), int32(169), int32(15506))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v576 = v347
	goto L3
L94:
	;
	v387 = v379
	v388 = v384
	goto L97
L95:
	;
	v411 = v378
	goto L96
L96:
	;
	if v39 != 0 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v399 = v388 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v399
	v401 = F_repalloc(m, v387, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	v411 = v404
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v401
	v404 = v401 + v380
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v406 <= v383 {
		v387 = v401
		v388 = v406
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v423 = v422 + v39
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v423
	v425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v423))) = uint8(v425)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v428 + int32(1)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v432 + v382
	if v24 == v425 {
		v24 = v425
		v26 = v39
		v27 = v40
		v29 = v42
		goto L4
	} else {
		goto L105
	}
L102:
	;
	v420 = F__emscripten_memcpy_bulkmem(m, v411, v42, v39)
	mBase = m.M
	goto L104
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	v441 = v24
	goto L106
L106:
	;
	v450 = v441 - int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v15+v450<<(uint(int32(2))%32))))
	switch v454 - int32(33) {
	case 0, 5:
		goto L108
	default:
		v24 = v441
		v26 = v39
		v27 = v40
		v29 = v42
		goto L4
	}
L107:
	;
	v24 = int32(0)
	v26 = v39
	v27 = v40
	v29 = v42
	goto L4
L108:
	;
	v459 = F_palloc(m, int32(20))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v461 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v459)+12)) = uint16(v461)
	*(*int32)(unsafe.Add(mBase, uint32(v459)+4)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v459)+8)) = v461
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+16)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v459
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v471 + int32(1)
	if v450 != 0 {
		v441 = v450
		goto L106
	} else {
		goto L110
	}
L110:
	;
	goto L107
L111:
	;
	if v477 == int32(0) {
		v576 = v475
		goto L3
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(341845), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errsave_finish(m, v476, int32(516089), int32(187), int32(163116))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v576 = v475
	goto L3
L116:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v497 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v510 = F_errsave_start(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L122
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
	v530 = v499
	v532 = v46
	v533 = int32(0)
	v534 = v499
	goto L10
L119:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	switch v500 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v530 = v499
		v532 = v39
		v533 = v40
		v534 = v42
		goto L10
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35:
		goto L117
	case 36:
		goto L118
	default:
		goto L120
	}
L120:
	;
	if v500 != int32(95) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	if v510 == int32(0) {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(221493), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errsave_finish(m, v509, int32(516089), int32(94), int32(15957))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L8
L127:
	;
	v552 = int32(1)
	v553 = F_errsave_start(m, v551)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L133
	}
L128:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	if v544 != int32(447) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v551 = v541
	goto L127
L130:
	;
	goto L131
L131:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+4)))
	if v547 == int32(0) {
		v551 = v541
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v576 = int32(1)
	goto L3
L133:
	;
	if v553 == int32(0) {
		v576 = v552
		goto L3
	} else {
		goto L134
	}
L134:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(221501), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errsave_finish(m, v551, int32(516089), int32(284), int32(313572))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v576 = v552
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
		m.T0[v7].(func(*base.Module, int32, int32, int32))(m, l0, int32(243295), int32(0))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v12 - int32(63) {
	case 0:
		v36 = int32(4)
		goto L1
	case 1:
		goto L2
	default:
		goto L3
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1+v36)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = v39
	goto L10
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_markRelsAsNulledBy(m, l0, v29, l2)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
	F_errmsg_internal(m, int32(504634), v9)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(519109), int32(1795), int32(27361))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
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
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_markRelsAsNulledBy(m, l0, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v36 = int32(36)
	goto L1
L10:
	;
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = v54 + v38<<(uint(int32(2))%32) - int32(4)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = F_bms_add_member(m, v60, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L19
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v48 = v47
	goto L14
L13:
	;
	v48 = int32(0)
	goto L14
L14:
	;
	if v48 < v38 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v51 = F_lappend(m, v41, int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v51
	v41 = v51
	goto L10
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v61
	m.G0 = v9 + int32(16)
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
			v29 = int32(4549024)
			v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
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
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
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
				v29 = int32(4549024)
				v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
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
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
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
					v29 = int32(4549024)
					v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
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
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
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
	var v7 int32
	_ = v7
	v7 = F__emscripten_memset_bulkmem(m, l0+int32(24), base.I32_extend8_s(int32(0)), int32(8168))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
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
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if base.Ui32(v12) < base.Ui32(v11) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
			F_errmsg_internal(m, int32(50909), v8)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errfinish(m, int32(518028), int32(82), int32(436332))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if base.Ui32(int32(8192)) < base.Ui32(v10) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
				F_errmsg_internal(m, int32(50909), v8)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errfinish(m, int32(518028), int32(82), int32(436332))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if base.Ui32(v11) < base.Ui32(int32(24)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
					F_errmsg_internal(m, int32(50909), v8)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(518028), int32(82), int32(436332))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if base.Ui32(v12) <= base.Ui32(v10) {
					v38 = F__emscripten_memset_bulkmem(m, l0+v11, base.I32_extend8_s(int32(0)), v12-v11)
					mBase = m.M
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg_internal(m, int32(50909), v8)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_errfinish(m, int32(518028), int32(82), int32(436332))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
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
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 float64
	_ = v319
	var v321 float64
	_ = v321
	var v323 float64
	_ = v323
	var v325 float64
	_ = v325
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
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
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
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
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
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
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v929 int32
	_ = v929
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
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
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1196 int32
	_ = v1196
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1326 int32
	_ = v1326
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
	if base.Ui32(v35) < base.Ui32(v34) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v38&int32(1) == int32(0) {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v38 = v37
	goto L6
L5:
	;
	v38 = int32(1)
	goto L6
L6:
	;
	goto L3
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v43 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v56 = int32(0)
	goto L9
L9:
	;
	v74 = v56 << (uint(int32(2)) % 32)
	v75 = l3 + int32(4) + v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v76 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v147 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v82 = int32(0)
	if v82 < v79 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v85 = v79
	goto L16
L15:
	;
	v85 = v82
	goto L16
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v92 = int32(0)
	goto L17
L17:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v86+v92<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v117 == l1 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L11
L19:
	;
	v120 = v92 + int32(1)
	if v120 != v85 {
		v92 = v120
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v1397 = v56 + int32(1)
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v1397 < v1398 {
		v56 = v1397
		goto L9
	} else {
		goto L327
	}
L22:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v74)))
	if base.Ui32(v152) <= base.Ui32(int32(16383)) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L32
	} else {
		goto L324
	}
L24:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1347 = F_lappend(m, v1346, v1326)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L32
	} else {
		goto L323
	}
L25:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	switch v167 - int32(15) {
	case 0:
		goto L42
	default:
		goto L39
	case 2:
		goto L43
	case 5:
		goto L41
	case 22:
		goto L40
	}
L26:
	;
	v164 = F_match_boolean_index_clause(m, l0, l1, v56, l2)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L32
	} else {
		goto L35
	}
L27:
	;
	if v152 == int32(424) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v160 = F_op_in_opfamily(m, int32(91), v152)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v152 == int32(2222) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	return
L33:
	;
	if v160 == int32(0) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	if v164 != 0 {
		v1326 = v164
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	if v1299 == int32(0) {
		goto L21
	} else {
		goto L322
	}
L38:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v1292 = F_get_index_clause_from_support(m, l0, l1, v1291, v370, v56, l2)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L32
	} else {
		goto L321
	}
L39:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L254
L40:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if v457 != int32(403) {
		goto L21
	} else {
		goto L116
	}
L41:
	;
	v402 = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+20)))
	if v404 != int32(1) {
		v1299 = v402
		goto L37
	} else {
		goto L100
	}
L42:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+28))
	if v359 == int32(0) {
		goto L21
	} else {
		goto L93
	}
L43:
	;
	v170 = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+28))
	if v172 == v170 {
		v1299 = v170
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v175 != int32(2) {
		v1299 = v170
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178+v74)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v74)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+68))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v191 = F_match_index_to_operand(m, v190, v56, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L32
	} else {
		goto L47
	}
L46:
	;
	v232 = F_match_index_to_operand(m, v185, v56, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L32
	} else {
		goto L64
	}
L47:
	;
	if v191 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v196 = F_bms_is_member(m, v189, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L32
	} else {
		goto L49
	}
L49:
	;
	if v196 != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v198 = F_contain_volatile_functions(m, v185)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L32
	} else {
		goto L51
	}
L51:
	;
	if v198 != 0 {
		goto L46
	} else {
		goto L52
	}
L52:
	;
	if v180 != v186 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_set_opfuncid(m, v171)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L32
	} else {
		goto L62
	}
L54:
	;
	v202 = v180
	goto L56
L55:
	;
	v202 = int32(0)
	goto L56
L56:
	;
	if v202 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v203 = F_op_in_opfamily(m, v187, v183)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L32
	} else {
		goto L58
	}
L58:
	;
	if v203 == int32(0) {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v208 = F_palloc0(m, int32(20))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L32
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = l1
	v218 = F_list_make1_impl(m, int32(1), v28+int32(24))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+16)) = v220
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+14)) = uint16(v56)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+12)) = uint8(v220)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v218
	v1299 = v208
	goto L37
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	v230 = F_get_index_clause_from_support(m, l0, l1, v228, int32(0), v56, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	v1299 = v230
	goto L37
L64:
	;
	if v232 == int32(0) {
		v1299 = v170
		goto L37
	} else {
		goto L65
	}
L65:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v237 = F_bms_is_member(m, v189, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L32
	} else {
		goto L66
	}
L66:
	;
	if v237 != 0 {
		v1299 = v170
		goto L37
	} else {
		goto L67
	}
L67:
	;
	v239 = F_contain_volatile_functions(m, v190)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L32
	} else {
		goto L68
	}
L68:
	;
	if v239 != 0 {
		v1299 = v170
		goto L37
	} else {
		goto L69
	}
L69:
	;
	if v180 != v186 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	F_set_opfuncid(m, v171)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L32
	} else {
		goto L91
	}
L71:
	;
	v243 = v180
	goto L73
L72:
	;
	v243 = int32(0)
	goto L73
L73:
	;
	if v243 != 0 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v244 = F_get_commutator(m, v187)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L32
	} else {
		goto L75
	}
L75:
	;
	if v244 == int32(0) {
		goto L70
	} else {
		goto L76
	}
L76:
	;
	v248 = F_op_in_opfamily(m, v244, v183)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L32
	} else {
		goto L77
	}
L77:
	;
	if v248 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L78:
	;
	v252 = m.G0
	v254 = v252 - int32(16)
	m.G0 = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v258 = F_palloc0(m, int32(36))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L32
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = int32(17)
	v263 = v258 + int32(8)
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v263))) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v266
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v256)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+16)) = v268
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v256)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+24)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v256)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+32)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v244
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v256)+28))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+12)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+8)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v281
	v287 = F_list_make2_impl(m, v254+int32(4), v254)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L32
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+28)) = v287
	v291 = F_palloc0(m, int32(168))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L32
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = int32(318)
	goto L83
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+4)) = v258
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+44)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+48)) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+100)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+104)) = v305
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+108)) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v310 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v296)+116)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v296)+112)) = v309
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v314 == v315 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v296 = F__emscripten_memcpy_bulkmem(m, v291, l1, int32(168))
	mBase = m.M
	goto L85
L85:
	;
	goto L82
L86:
	;
	v317 = v244
	goto L88
L87:
	;
	v317 = v310
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+124)) = v317
	v319 = *(*float64)(unsafe.Add(mBase, uint32(l1)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v296)+128)) = v319
	v321 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v296)+136)) = v321
	v323 = *(*float64)(unsafe.Add(mBase, uint32(l1)+152))
	*(*float64)(unsafe.Add(mBase, uint32(v296)+144)) = v323
	v325 = *(*float64)(unsafe.Add(mBase, uint32(l1)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+160)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v296)+152)) = v325
	m.G0 = v254 + int32(16)
	v333 = F_palloc0(m, int32(20))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L32
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v296
	v343 = F_list_make1_impl(m, int32(1), v28+int32(20))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L32
	} else {
		goto L90
	}
L90:
	;
	v345 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v333)+16)) = v345
	*(*uint16)(unsafe.Add(mBase, uint32(v333)+14)) = uint16(v56)
	*(*uint8)(unsafe.Add(mBase, uint32(v333)+12)) = uint8(v345)
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = v343
	v1299 = v333
	goto L37
L91:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	v356 = F_get_index_clause_from_support(m, l0, l1, v354, int32(1), v56, l2)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L32
	} else {
		goto L92
	}
L92:
	;
	v1299 = v356
	goto L37
L93:
	;
	v362 = int32(0)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v363 <= v362 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	v370 = v362
	goto L95
L95:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391+v370<<(uint(int32(2))%32))))
	v396 = F_match_index_to_operand(m, v395, v56, l2)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L32
	} else {
		goto L97
	}
L96:
	;
	goto L21
L97:
	;
	if v396 != 0 {
		goto L38
	} else {
		goto L98
	}
L98:
	;
	v399 = v370 + int32(1)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v399 < v400 {
		v370 = v399
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403)+28))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v411 = F_pull_varnos(m, l0, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L32
	} else {
		goto L101
	}
L101:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v413+v74)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v416+v74)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v403)+24))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+68))
	v423 = F_match_index_to_operand(m, v409, v56, l2)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L32
	} else {
		goto L102
	}
L102:
	;
	if v423 == int32(0) {
		v1299 = v402
		goto L37
	} else {
		goto L103
	}
L103:
	;
	v427 = F_bms_is_member(m, v422, v411)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L32
	} else {
		goto L104
	}
L104:
	;
	if v427 != 0 {
		v1299 = v402
		goto L37
	} else {
		goto L105
	}
L105:
	;
	v429 = F_contain_volatile_functions(m, v410)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L32
	} else {
		goto L106
	}
L106:
	;
	if v429 != 0 {
		v1299 = v402
		goto L37
	} else {
		goto L107
	}
L107:
	;
	if v415 != v419 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v433 = v415
	goto L110
L109:
	;
	v433 = int32(0)
	goto L110
L110:
	;
	if v433 != 0 {
		v1299 = v402
		goto L37
	} else {
		goto L111
	}
L111:
	;
	v434 = F_op_in_opfamily(m, v420, v418)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L32
	} else {
		goto L112
	}
L112:
	;
	if v434 == int32(0) {
		v1299 = v402
		goto L37
	} else {
		goto L113
	}
L113:
	;
	v439 = F_palloc0(m, int32(20))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L32
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = l1
	v449 = F_list_make1_impl(m, int32(1), v28+int32(28))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L32
	} else {
		goto L115
	}
L115:
	;
	v451 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v439)+16)) = v451
	*(*uint16)(unsafe.Add(mBase, uint32(v439)+14)) = uint16(v56)
	*(*uint8)(unsafe.Add(mBase, uint32(v439)+12)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v439)+8)) = v449
	v1299 = v439
	goto L37
L116:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v460+v74)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+68))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+12))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v465)+24))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v465)+20))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+12))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v475+v74)))
	if v477 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v465)+16))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v477 != v480 {
		goto L21
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v482 = F_match_index_to_operand(m, v474, v56, l2)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L32
	} else {
		goto L123
	}
L120:
	;
	goto L119
L121:
	;
	v510 = F_get_op_opfamily_strategy(m, v508, v462)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L32
	} else {
		goto L139
	}
L122:
	;
	v493 = F_match_index_to_operand(m, v471, v56, l2)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L32
	} else {
		goto L130
	}
L123:
	;
	if v482 == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v486 = F_pull_varnos(m, l0, v471)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L32
	} else {
		goto L125
	}
L125:
	;
	v488 = F_bms_is_member(m, v464, v486)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L32
	} else {
		goto L126
	}
L126:
	;
	if v488 != 0 {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v490 = F_contain_volatile_functions(m, v471)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L32
	} else {
		goto L128
	}
L128:
	;
	if v490 != 0 {
		goto L122
	} else {
		goto L129
	}
L129:
	;
	v508 = v468
	v509 = int32(1)
	goto L121
L130:
	;
	if v493 == int32(0) {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	v497 = F_pull_varnos(m, l0, v474)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L32
	} else {
		goto L132
	}
L132:
	;
	v499 = F_bms_is_member(m, v464, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L32
	} else {
		goto L133
	}
L133:
	;
	if v499 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	v501 = F_contain_volatile_functions(m, v474)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L32
	} else {
		goto L135
	}
L135:
	;
	if v501 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	v503 = F_get_commutator(m, v468)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L32
	} else {
		goto L137
	}
L137:
	;
	if v503 == int32(0) {
		goto L21
	} else {
		goto L138
	}
L138:
	;
	v508 = v503
	v509 = int32(0)
	goto L121
L139:
	;
	if base.Ui32(int32(5)) < base.Ui32(v510) {
		goto L21
	} else {
		goto L140
	}
L140:
	;
	if int32(1)<<(uint(v510)%32)&int32(54) == int32(0) {
		goto L21
	} else {
		goto L141
	}
L141:
	;
	v521 = F_palloc0(m, int32(20))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L32
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = int32(281)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v521)+14)) = uint16(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v521)+4)) = l1
	if v509 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v530 = int32(24)
	goto L145
L144:
	;
	v530 = int32(20)
	goto L145
L145:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v525+v530)))
	if v509 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v535 = int32(20)
	goto L148
L147:
	;
	v535 = int32(24)
	goto L148
L148:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v525+v535)))
	v539 = v56 << (uint(int32(2)) % 32)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v539+v540)))
	F_get_op_opfamily_properties(m, v508, v542, int32(0), v28+int32(124), v28+int32(120), v28+int32(116))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L32
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v56
	v557 = F_list_make1_impl(m, int32(471), v28+int32(84))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L32
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+16)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v508
	v565 = F_list_make1_impl(m, int32(472), v28+int32(80))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L32
	} else {
		goto L151
	}
L151:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v567+v539)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v569
	v575 = F_list_make1_impl(m, int32(472), v28+int32(76))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L32
	} else {
		goto L152
	}
L152:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v28)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v577
	v583 = F_list_make1_impl(m, int32(472), v28+int32(72))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L32
	} else {
		goto L153
	}
L153:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v585
	v591 = F_list_make1_impl(m, int32(472), v28+int32(68))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L32
	} else {
		goto L154
	}
L154:
	;
	v601 = int32(1)
	v607 = v575
	v608 = v565
	v609 = v591
	v612 = v583
	goto L155
L155:
	;
	if v537 != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	if v795 != 0 {
		goto L193
	} else {
		goto L194
	}
L157:
	;
	goto L156
L158:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v621 = v619
	goto L160
L159:
	;
	v621 = int32(0)
	goto L160
L160:
	;
	if v621 <= v601 {
		goto L157
	} else {
		goto L161
	}
L161:
	;
	v624 = v601 << (uint(int32(2)) % 32)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v532)+12))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v624+v625)))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v628+v624)))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+12))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v632+v624)))
	if v509 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v637 = F_get_commutator(m, v634)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L32
	} else {
		goto L165
	}
L163:
	;
	v641 = v634
	goto L164
L164:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+68))
	v644 = F_pull_varnos(m, l0, v627)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L32
	} else {
		goto L167
	}
L165:
	;
	if v637 == int32(0) {
		goto L157
	} else {
		goto L166
	}
L166:
	;
	v641 = v637
	goto L164
L167:
	;
	v646 = F_bms_is_member(m, v643, v644)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L32
	} else {
		goto L168
	}
L168:
	;
	if v646 != 0 {
		goto L157
	} else {
		goto L169
	}
L169:
	;
	v648 = F_contain_volatile_functions(m, v627)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L32
	} else {
		goto L170
	}
L170:
	;
	if v648 != 0 {
		goto L157
	} else {
		goto L171
	}
L171:
	;
	v650 = int32(0)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v650 < v651 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v658 = v650
	goto L175
L173:
	;
	v714 = v650
	v735 = v651
	goto L174
L174:
	;
	if v735 <= v714 {
		goto L157
	} else {
		goto L186
	}
L175:
	;
	v679 = F_match_index_to_operand(m, v630, v658, l2)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L32
	} else {
		goto L179
	}
L176:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v714 = v707
	v735 = v709
	goto L174
L177:
	;
	goto L176
L178:
	;
	v704 = v658 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v704 < v705 {
		v658 = v704
		goto L175
	} else {
		goto L185
	}
L179:
	;
	if v679 == int32(0) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v684 = v658 << (uint(int32(2)) % 32)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v684+v685)))
	v688 = F_get_op_opfamily_strategy(m, v641, v687)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L32
	} else {
		goto L181
	}
L181:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	if v688 != v690 {
		goto L178
	} else {
		goto L182
	}
L182:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v692+v684)))
	if v694 == int32(0) {
		v707 = v658
		goto L177
	} else {
		goto L183
	}
L183:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v525)+16))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)+12))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v698+v624)))
	if v694 == v700 {
		v707 = v658
		goto L177
	} else {
		goto L184
	}
L184:
	;
	goto L178
L185:
	;
	v707 = v704
	goto L177
L186:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v521)+16))
	v738 = F_lappend_int(m, v737, v714)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L32
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+16)) = v738
	v742 = v714 << (uint(int32(2)) % 32)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v742+v743)))
	F_get_op_opfamily_properties(m, v641, v745, int32(0), v28+int32(124), v28+int32(120), v28+int32(116))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L32
	} else {
		goto L188
	}
L188:
	;
	v757 = F_lappend_oid(m, v608, v641)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L32
	} else {
		goto L189
	}
L189:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v759+v742)))
	v762 = F_lappend_oid(m, v607, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L32
	} else {
		goto L190
	}
L190:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v28)+120))
	v765 = F_lappend_oid(m, v612, v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L32
	} else {
		goto L191
	}
L191:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	v768 = F_lappend_oid(m, v609, v767)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L32
	} else {
		goto L192
	}
L192:
	;
	v601 = v601 + int32(1)
	v607 = v762
	v608 = v757
	v609 = v768
	v612 = v765
	goto L155
L193:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	v798 = v796
	goto L195
L194:
	;
	v798 = int32(0)
	goto L195
L195:
	;
	v799 = base.B2i32(v601 != v798)
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+12)) = uint8(v799)
	if v509 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v1033
	v1038 = F_list_make1_impl(m, int32(1), v28+int32(44))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L32
	} else {
		goto L253
	}
L197:
	;
	if v798 == v601 {
		v929 = v608
		goto L200
	} else {
		goto L201
	}
L198:
	;
	if v601 != v798 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = l1
	v1032 = v28 + int32(92)
	goto L196
L200:
	;
	if int32(2) <= v601 {
		goto L240
	} else {
		goto L241
	}
L201:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	switch v808 - int32(1) {
	case 0:
		v828 = int32(2)
		goto L203
	case 1, 3:
		goto L202
	default:
		goto L205
	case 4:
		goto L204
	}
L202:
	;
	v906 = int32(0)
	if v608 == v906 {
		v914 = v906
		goto L234
	} else {
		goto L235
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v828
	v830 = int32(0)
	v834 = v830
	v836 = v830
	goto L209
L204:
	;
	v828 = int32(4)
	goto L203
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L32
	} else {
		goto L206
	}
L206:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v815
	F_errmsg_internal(m, int32(489434), v28+int32(32))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L32
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(518243), int32(3643), int32(380441))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L32
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	v857 = int32(0)
	if v607 == v857 {
		v867 = v857
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v868 = int32(0)
	if v612 == v868 {
		v879 = v868
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	if v861 <= v836 {
		v867 = int32(0)
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v607)+12))
	v867 = v863 + v836<<(uint(int32(2))%32)
	goto L211
L214:
	;
	if v609 == int32(0) {
		v929 = v868
		goto L200
	} else {
		goto L217
	}
L215:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if v873 <= v836 {
		v879 = int32(0)
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v612)+12))
	v879 = v875 + v836<<(uint(int32(2))%32)
	goto L214
L217:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	if v882 <= v836 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v929 = v834
	goto L200
L219:
	;
	goto L220
L220:
	;
	if v867 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v929 = v834
	goto L200
L222:
	;
	goto L223
L223:
	;
	if v879 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v929 = v834
	goto L200
L225:
	;
	goto L226
L226:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	v891 = v888 + v836<<(uint(int32(2))%32)
	if v891 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v929 = v834
	goto L200
L228:
	;
	goto L229
L229:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v891)))
	v897 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+124)))
	v898 = F_get_opfamily_member(m, v894, v895, v896, v897)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L32
	} else {
		goto L230
	}
L230:
	;
	if v898 == int32(0) {
		goto L23
	} else {
		goto L231
	}
L231:
	;
	v904 = F_lappend_oid(m, v834, v898)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L32
	} else {
		goto L232
	}
L232:
	;
	v834 = v904
	v836 = v836 + int32(1)
	goto L209
L233:
	;
	v929 = v914
	goto L200
L234:
	;
	goto L233
L235:
	;
	if v601 <= int32(0) {
		v914 = v906
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v601 < v911 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608)+4)) = v601
	goto L239
L238:
	;
	goto L239
L239:
	;
	v914 = v608
	goto L234
L240:
	;
	v943 = F_palloc0(m, int32(28))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L32
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+16)) = int32(0)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v929)+12))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v979)))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	v983 = F_copyObjectImpl(m, v982)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L32
	} else {
		goto L249
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = int32(37)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v943)+8)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v943)+4)) = v947
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	v951 = F_list_copy_head(m, v950, v601)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L32
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943)+12)) = v951
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v525)+16))
	v955 = F_list_copy_head(m, v954, v601)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L32
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943)+16)) = v955
	v958 = F_list_copy_head(m, v537, v601)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L32
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943)+20)) = v958
	v961 = F_list_copy_head(m, v532, v601)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L32
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943)+24)) = v961
	v965 = int32(0)
	v972 = F_make_restrictinfo(m, l0, v943, int32(1), v965, v965, v965, v965, v965, v965, v965)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L32
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v972
	v1032 = v28 + int32(92)
	goto L196
L249:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v532)+12))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)))
	v987 = F_copyObjectImpl(m, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L32
	} else {
		goto L250
	}
L250:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v525)+16))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)+12))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	v992 = F_make_opclause(m, v980, v983, v987, v991)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L32
	} else {
		goto L251
	}
L251:
	;
	v995 = int32(0)
	v1002 = F_make_restrictinfo(m, l0, v992, int32(1), v995, v995, v995, v995, v995, v995, v995)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L32
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v1002
	v1032 = v28 + int32(92)
	goto L196
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+8)) = v1038
	v1326 = v521
	goto L24
L254:
	;
	if v1041 != int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1044 = int32(0)
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+107)))
	if v1045 != int32(1) {
		v1299 = v1044
		goto L37
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+108)))
	if v1260 != int32(1) {
		goto L21
	} else {
		goto L314
	}
L258:
	;
	v1048 = int32(0)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+8))
	if v1050 == v1048 {
		v1224 = v1048
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1229 = F_make_SAOP_expr(m, v1173, v1136, v1175, v1174, v1174, v1180, v1179&int32(1))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L32
	} else {
		goto L310
	}
L260:
	;
	F_list_free(m, v1224)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L32
	} else {
		goto L309
	}
L261:
	;
	v1054 = int32(0)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+4))
	if v1055 <= v1054 {
		v1224 = v1054
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+68))
	v1060 = int32(0)
	v1074 = v1054
	v1079 = int32(1)
	v1080 = v1060
	v1081 = v1060
	v1083 = v1060
	v1086 = v1060
	v1087 = v1060
	goto L264
L263:
	;
	v1224 = v1196
	goto L260
L264:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+12))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1090+v1080<<(uint(int32(2))%32))))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)))
	if v1095 != int32(318) {
		v1196 = v1083
		goto L263
	} else {
		goto L266
	}
L265:
	;
	if v1136 != 0 {
		goto L259
	} else {
		goto L308
	}
L266:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+4))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	if v1099 != int32(17) {
		v1196 = v1083
		goto L263
	} else {
		goto L267
	}
L267:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+28))
	if v1102 == int32(0) {
		v1196 = v1083
		goto L263
	} else {
		goto L268
	}
L268:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+4))
	if v1105 != int32(2) {
		v1196 = v1083
		goto L263
	} else {
		goto L269
	}
L269:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+12))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1109)))
	v1112 = F_match_index_to_operand(m, v1111, v56, l2)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L32
	} else {
		goto L272
	}
L270:
	;
	if v1079&int32(1) != 0 {
		goto L287
	} else {
		goto L288
	}
L271:
	;
	v1121 = F_match_index_to_operand(m, v1110, v56, l2)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L32
	} else {
		goto L278
	}
L272:
	;
	if v1112 == int32(0) {
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+48))
	v1117 = F_bms_is_member(m, v1059, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L32
	} else {
		goto L274
	}
L274:
	;
	if v1117 != 0 {
		goto L271
	} else {
		goto L275
	}
L275:
	;
	v1119 = F_contain_volatile_functions(m, v1110)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L32
	} else {
		goto L276
	}
L276:
	;
	if v1119 != 0 {
		goto L271
	} else {
		goto L277
	}
L277:
	;
	v1134 = v1110
	v1135 = v1108
	v1136 = v1111
	goto L270
L278:
	;
	if v1121 == int32(0) {
		v1196 = v1083
		goto L263
	} else {
		goto L279
	}
L279:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+44))
	v1126 = F_bms_is_member(m, v1059, v1125)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L32
	} else {
		goto L280
	}
L280:
	;
	if v1126 != 0 {
		v1196 = v1083
		goto L263
	} else {
		goto L281
	}
L281:
	;
	v1128 = F_contain_volatile_functions(m, v1111)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L32
	} else {
		goto L282
	}
L282:
	;
	if v1128 != 0 {
		v1196 = v1083
		goto L263
	} else {
		goto L283
	}
L283:
	;
	v1130 = F_get_commutator(m, v1108)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L32
	} else {
		goto L284
	}
L284:
	;
	if v1130 == int32(0) {
		v1196 = v1083
		goto L263
	} else {
		goto L285
	}
L285:
	;
	v1134 = v1111
	v1135 = v1130
	v1136 = v1110
	goto L270
L286:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1134)))
	v1179 = base.B2i32(v1176 != int32(7)) | v1074
	v1180 = F_lappend(m, v1083, v1134)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L32
	} else {
		goto L306
	}
L287:
	;
	v1139 = F_exprType(m, v1134)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L32
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	if v1135 != v1081 {
		v1196 = v1083
		goto L263
	} else {
		goto L302
	}
L290:
	;
	v1141 = F_get_array_type(m, v1139)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L32
	} else {
		goto L291
	}
L291:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+24))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1144+v74)))
	if v1146 != v1143 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1149 = v1146
	goto L294
L293:
	;
	v1149 = int32(0)
	goto L294
L294:
	;
	if v1149 != 0 {
		v1196 = v1083
		goto L263
	} else {
		goto L295
	}
L295:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1150+v74)))
	v1153 = F_op_in_opfamily(m, v1135, v1152)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L32
	} else {
		goto L296
	}
L296:
	;
	if v1139 == int32(2249) {
		v1196 = v1083
		goto L263
	} else {
		goto L297
	}
L297:
	;
	if v1153 == int32(0) {
		v1196 = v1083
		goto L263
	} else {
		goto L298
	}
L298:
	;
	if v1141 == int32(0) {
		v1196 = v1083
		goto L263
	} else {
		goto L299
	}
L299:
	;
	v1161 = F_exprType(m, v1136)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L32
	} else {
		goto L300
	}
L300:
	;
	if v1161 != int32(2249) {
		v1173 = v1135
		v1174 = v1143
		v1175 = v1139
		goto L286
	} else {
		goto L301
	}
L301:
	;
	v1196 = v1083
	goto L263
L302:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+24))
	if v1086 != v1166 {
		v1196 = v1083
		goto L263
	} else {
		goto L303
	}
L303:
	;
	v1168 = F_exprType(m, v1134)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L32
	} else {
		goto L304
	}
L304:
	;
	if v1168 != v1087 {
		v1196 = v1083
		goto L263
	} else {
		goto L305
	}
L305:
	;
	v1173 = v1081
	v1174 = v1086
	v1175 = v1087
	goto L286
L306:
	;
	v1184 = v1080 + int32(1)
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+4))
	if v1184 < v1185 {
		v1074 = v1179
		v1079 = int32(0)
		v1080 = v1184
		v1081 = v1173
		v1083 = v1180
		v1086 = v1174
		v1087 = v1175
		goto L264
	} else {
		goto L307
	}
L307:
	;
	goto L265
L308:
	;
	v1196 = v1180
	goto L263
L309:
	;
	v1299 = v1044
	goto L37
L310:
	;
	v1232 = F_palloc0(m, int32(20))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L32
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1232))) = int32(281)
	v1238 = int32(0)
	v1245 = F_make_restrictinfo(m, l0, v1229, int32(1), v1238, v1238, v1238, v1238, v1238, v1238, v1238)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L32
	} else {
		goto L312
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v1245
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v1245
	v1252 = F_list_make1_impl(m, int32(1), v28+int32(12))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L32
	} else {
		goto L313
	}
L313:
	;
	v1254 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+16)) = v1254
	*(*uint16)(unsafe.Add(mBase, uint32(v1232)+14)) = uint16(v56)
	*(*uint8)(unsafe.Add(mBase, uint32(v1232)+12)) = uint8(v1254)
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+8)) = v1252
	v1299 = v1232
	goto L37
L314:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v1263 != int32(52) {
		goto L21
	} else {
		goto L315
	}
L315:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+12)))
	if v1266 != 0 {
		goto L21
	} else {
		goto L316
	}
L316:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v1268 = F_match_index_to_operand(m, v1267, v56, l2)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L32
	} else {
		goto L317
	}
L317:
	;
	if v1268 == int32(0) {
		goto L21
	} else {
		goto L318
	}
L318:
	;
	v1273 = F_palloc0(m, int32(20))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L32
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1273))) = int32(281)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = l1
	v1283 = F_list_make1_impl(m, int32(1), v28+int32(16))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L32
	} else {
		goto L320
	}
L320:
	;
	v1285 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+16)) = v1285
	*(*uint16)(unsafe.Add(mBase, uint32(v1273)+14)) = uint16(v56)
	*(*uint8)(unsafe.Add(mBase, uint32(v1273)+12)) = uint8(v1285)
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+8)) = v1283
	v1326 = v1273
	goto L24
L321:
	;
	v1299 = v1292
	goto L37
L322:
	;
	v1326 = v1299
	goto L24
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v1347
	v1350 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1350)
	goto L1
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v894
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v1359
	F_errmsg_internal(m, int32(42007), v28+int32(48))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L32
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(518243), int32(3657), int32(380441))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L32
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
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
	F_errmsg_internal(m, int32(456299), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(512720), int32(50), int32(238458))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
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
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
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
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
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
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l1
	v19 = v14 + int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1486]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 == v4 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v463 + int32(16)
	return v466
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v446
	v463 = v14
	v466 = l2
	goto L1
L3:
	;
	v463 = v439
	v466 = int32(-1)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v422
	v439 = v14
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(25)
	if v404 == int32(0) {
		v439 = v14
		goto L3
	} else {
		goto L85
	}
L6:
	;
	if v385&int32(255) != 0 {
		v404 = v386
		v406 = v388
		goto L5
	} else {
		goto L81
	}
L7:
	;
	v383 = v374 - int32(1)
	if v371 != 0 {
		v404 = v372
		v406 = v383
		goto L5
	} else {
		goto L80
	}
L8:
	;
	v190 = v4
	v191 = l0
	v193 = v20
	v194 = v4
	v195 = l2
	goto L51
L9:
	;
	v60 = v4
	v63 = v20
	v64 = int32(1)
	v65 = l2
	goto L23
L10:
	;
	v58 = F_strlen(m, v20)
	mBase = m.M
	v463 = v14
	v466 = v58
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
		v446 = v20
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v30 = l2
	v31 = l0
	v33 = v20
	goto L16
L16:
	;
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33))))
	if v41 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v51
	v463 = v14
	v466 = l2 - v30
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v41 & int32(57343)
	v47 = int32(1)
	v48 = v33 + v47
	v50 = v30 - v47
	if v50 != 0 {
		v30 = v50
		v31 = v31 + int32(4)
		v33 = v48
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
	v446 = v48
	goto L2
L22:
	;
	goto L8
L23:
	;
	if v64 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v75 = int32(base.Ui32(v73) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v75-int32(16)|(v60>>(uint(int32(26))%32)+v75)) {
		v371 = v60
		v372 = l0
		v374 = v63
		v376 = v65
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v116 = v63
	v118 = v65
	goto L38
L28:
	;
	v85 = v63 + int32(1)
	if v60&int32(33554432) == int32(0) {
		v109 = v85
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v110 = int32(1)
	v63 = v109
	v64 = v110
	v65 = v65 - v110
	goto L23
L30:
	;
	v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85))))
	if int32(-64) <= v90 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v404 = l0
	v406 = v63 - int32(1)
	goto L5
L32:
	;
	goto L33
L33:
	;
	v96 = v63 + int32(2)
	if v60&int32(524288) == int32(0) {
		v109 = v96
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	if int32(-64) <= v101 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v404 = l0
	v406 = v63 - int32(1)
	goto L5
L36:
	;
	goto L37
L37:
	;
	v109 = v63 + int32(3)
	goto L29
L38:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v116))))
	if v124 <= int32(0) {
		v158 = v124
		v161 = v116
		v163 = v118
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v179 = v158&int32(255) - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v179) {
		v385 = v158
		v386 = l0
		v388 = v161
		v390 = v163
		goto L6
	} else {
		goto L50
	}
L40:
	;
	if int32(0) < base.I32_extend8_s(v158) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	if v116&int32(3) != 0 {
		v158 = v124
		v161 = v116
		v163 = v118
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if (v129-int32(16843009)|v129)&int32(-2139062144) != 0 {
		v158 = v129
		v161 = v116
		v163 = v118
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v138 = v116
	v140 = v118
	goto L44
L44:
	;
	v146 = int32(4)
	v147 = v140 - v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v150 = v138 + v146
	if (v148|(v148-int32(16843009)))&int32(-2139062144) == int32(0) {
		v138 = v150
		v140 = v147
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v158 = v148
	v161 = v150
	v163 = v147
	goto L40
L46:
	;
	goto L45
L47:
	;
	v172 = int32(1)
	v116 = v161 + v172
	v118 = v163 - v172
	goto L38
L48:
	;
	goto L49
L49:
	;
	goto L39
L50:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179<<(uint(int32(2))%32))+uint32(_consts[1487])))
	v60 = v188
	v63 = v161 + int32(1)
	v64 = int32(0)
	v65 = v163
	goto L23
L51:
	;
	if v194 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v195 == int32(0) {
		v446 = v193
		goto L2
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v311 = int32(base.Ui32(v309) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v311-int32(16)|(v311+v190>>(uint(int32(26))%32))) {
		v371 = v190
		v372 = v191
		v374 = v193
		v376 = v195
		goto L7
	} else {
		goto L72
	}
L56:
	;
	v206 = v191
	v208 = v193
	v210 = v195
	goto L57
L57:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v217 = base.I32_extend8_s(v216)
	if v217 <= int32(0) {
		v286 = v217
		v287 = v206
		v289 = v208
		v290 = v216
		v291 = v210
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v298 = v290 - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v298) {
		v385 = v286
		v386 = v287
		v388 = v289
		v390 = v291
		goto L6
	} else {
		goto L71
	}
L59:
	;
	goto L58
L60:
	;
	if base.Ui32(v210) < base.Ui32(int32(5)) {
		v269 = v206
		v271 = v208
		v272 = v216
		v273 = v210
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v272
	v282 = int32(1)
	v283 = v271 + v282
	v285 = v273 - v282
	if v285 != 0 {
		v206 = v269 + int32(4)
		v208 = v283
		v210 = v285
		goto L57
	} else {
		goto L70
	}
L62:
	;
	if v208&int32(3) != 0 {
		v269 = v206
		v271 = v208
		v272 = v216
		v273 = v210
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v225 = v206
	v227 = v208
	v229 = v210
	goto L65
L64:
	;
	v264 = v259 & int32(255)
	if base.I32_extend8_s(v259) <= int32(0) {
		v286 = v259
		v287 = v260
		v289 = v261
		v290 = v264
		v291 = v262
		goto L59
	} else {
		goto L69
	}
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if (v235-int32(16843009)|v235)&int32(-2139062144) != 0 {
		v259 = v235
		v260 = v225
		v261 = v227
		v262 = v229
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v259 = v258
	v260 = v251
	v261 = v253
	v262 = v255
	goto L64
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v235 & int32(255)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v244
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v246
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v248
	v251 = v225 + int32(16)
	v252 = int32(4)
	v253 = v227 + v252
	v255 = v229 - v252
	if base.Ui32(v252) < base.Ui32(v255) {
		v225 = v251
		v227 = v253
		v229 = v255
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v269 = v260
	v271 = v261
	v272 = v264
	v273 = v262
	goto L61
L70:
	;
	v446 = v283
	goto L2
L71:
	;
	v301 = int32(1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v298<<(uint(int32(2))%32))+uint32(_consts[1487])))
	v190 = v307
	v191 = v287
	v193 = v289 + v301
	v194 = v301
	v195 = v291
	goto L51
L72:
	;
	v321 = v193 + int32(1)
	v326 = v309 - int32(128) | v190<<(uint(int32(6))%32)
	if int32(0) <= v326 {
		v351 = v326
		v355 = v321
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v191 = v191 + int32(4)
	v193 = v355
	v194 = int32(0)
	v195 = v195 - int32(1)
	goto L51
L74:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(25)
	v422 = v193 - int32(1)
	goto L4
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v351
	goto L73
L76:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	v331 = v329 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v331) {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v335 = v326 << (uint(int32(6)) % 32)
	v336 = v331 | v335
	v338 = v193 + int32(2)
	if int32(0) <= v335 {
		v351 = v336
		v355 = v338
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v343 = v341 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v343) {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v351 = v343 | v336<<(uint(int32(6))%32)
	v355 = v193 + int32(3)
	goto L75
L80:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	v385 = v384
	v386 = v372
	v388 = v383
	v390 = v376
	goto L6
L81:
	;
	if v386 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v386))) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v398
	goto L84
L83:
	;
	goto L84
L84:
	;
	v463 = v14
	v466 = l2 - v390
	goto L1
L85:
	;
	v422 = v406
	goto L4
}
func F_mcelem_array_contained_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v29 float64
	_ = v29
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v44 int32
	_ = v44
	var v48 float32
	_ = v48
	var v49 float32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 float32
	_ = v62
	var v66 float32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v83 float32
	_ = v83
	var v85 float32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v128 float32
	_ = v128
	var v130 float32
	_ = v130
	var v144 int32
	_ = v144
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
	var v158 float32
	_ = v158
	var v164 float32
	_ = v164
	var v165 float32
	_ = v165
	var v168 float32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v184 float32
	_ = v184
	var v186 float32
	_ = v186
	var v207 int32
	_ = v207
	var v216 float32
	_ = v216
	var v218 float32
	_ = v218
	var v237 int32
	_ = v237
	var v246 float32
	_ = v246
	var v248 float32
	_ = v248
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v277 float32
	_ = v277
	var v279 float32
	_ = v279
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v312 float32
	_ = v312
	var v314 float32
	_ = v314
	var v316 int32
	_ = v316
	var v331 float32
	_ = v331
	var v332 float32
	_ = v332
	var v335 float32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v353 float32
	_ = v353
	var v355 float32
	_ = v355
	var v358 int32
	_ = v358
	var v390 float32
	_ = v390
	var v392 float32
	_ = v392
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v409 float32
	_ = v409
	var v412 float32
	_ = v412
	var v415 float32
	_ = v415
	var v418 float32
	_ = v418
	var v419 float32
	_ = v419
	var v420 float32
	_ = v420
	var v431 float32
	_ = v431
	var v433 int32
	_ = v433
	var v447 float32
	_ = v447
	var v449 float32
	_ = v449
	var v466 float64
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 float64
	_ = v478
	var v486 float64
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
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
	var v511 int32
	_ = v511
	var __phi511 int32
	_ = __phi511
	var v525 int32
	_ = v525
	var __phi525 int32
	_ = __phi525
	var v527 int32
	_ = v527
	var __phi527 int32
	_ = __phi527
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 float32
	_ = v542
	var v547 float32
	_ = v547
	var v548 float32
	_ = v548
	var v553 int32
	_ = v553
	var v564 int32
	_ = v564
	var v582 int32
	_ = v582
	var v589 float32
	_ = v589
	var v592 int32
	_ = v592
	var v595 float32
	_ = v595
	var v601 float32
	_ = v601
	var v605 float32
	_ = v605
	var v609 int32
	_ = v609
	var v614 float32
	_ = v614
	var v618 float32
	_ = v618
	var v619 int32
	_ = v619
	var v623 float32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v667 float32
	_ = v667
	var v671 float32
	_ = v671
	var v673 int32
	_ = v673
	var v676 float32
	_ = v676
	var v711 int32
	_ = v711
	var v730 int32
	_ = v730
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v755 int32
	_ = v755
	var __phi755 int32
	_ = __phi755
	var v760 int32
	_ = v760
	var __phi760 int32
	_ = __phi760
	var v771 int32
	_ = v771
	var __phi771 int32
	_ = __phi771
	var v781 int32
	_ = v781
	var v784 float32
	_ = v784
	var v789 float32
	_ = v789
	var v790 float32
	_ = v790
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v810 int32
	_ = v810
	var v827 int32
	_ = v827
	var v835 float32
	_ = v835
	var v838 int32
	_ = v838
	var v841 float32
	_ = v841
	var v847 float32
	_ = v847
	var v851 float32
	_ = v851
	var v855 int32
	_ = v855
	var v860 float32
	_ = v860
	var v864 float32
	_ = v864
	var v865 int32
	_ = v865
	var v869 float32
	_ = v869
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v913 float32
	_ = v913
	var v917 float32
	_ = v917
	var v919 int32
	_ = v919
	var v922 float32
	_ = v922
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1014 float32
	_ = v1014
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1045 int32
	_ = v1045
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 float32
	_ = v1074
	var v1076 float32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 float32
	_ = v1083
	var v1085 float32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 float32
	_ = v1129
	var v1131 float32
	_ = v1131
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1201 int32
	_ = v1201
	var v1215 int32
	_ = v1215
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1241 float32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1262 float32
	_ = v1262
	var v1272 int32
	_ = v1272
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1313 int32
	_ = v1313
	var v1314 float32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1333 float32
	_ = v1333
	var v1336 float32
	_ = v1336
	var v1337 float32
	_ = v1337
	var v1345 int32
	_ = v1345
	var v1348 float32
	_ = v1348
	var v1349 float32
	_ = v1349
	var v1355 float32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1385 float32
	_ = v1385
	var v1394 int32
	_ = v1394
	var v1400 float32
	_ = v1400
	var v1402 float32
	_ = v1402
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1437 int32
	_ = v1437
	var v1444 float32
	_ = v1444
	var v1463 int32
	_ = v1463
	var v1465 float32
	_ = v1465
	var v1469 float32
	_ = v1469
	var v1472 float32
	_ = v1472
	var v1476 float32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1490 float32
	_ = v1490
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 float32
	_ = v1518
	var v1526 float32
	_ = v1526
	var v1554 float64
	_ = v1554
	v17 = int32(0)
	v29 = float64(0.005)
	if l2 == v17 {
		v1554 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v1554
L2:
	;
	if l3 != l1+int32(3) {
		v1554 = v29
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l6 == int32(0) {
		v1554 = v29
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l7 < int32(3) {
		v1554 = v29
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = int32(2)
	v41 = l2 + l1<<(uint(v39)%32)
	v42 = *(*float32)(unsafe.Add(mBase, uint32(v41)))
	v44 = l7 - int32(1)
	v48 = *(*float32)(unsafe.Add(mBase, uint32(l6+v44<<(uint(v39)%32))))
	v49 = *(*float32)(unsafe.Add(mBase, uint32(v41)+8))
	v52 = F_palloc(m, l5<<(uint(v39)%32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return float64(0)
L7:
	;
	if l5 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l1 <= v268 {
		v447 = v277
		v449 = v279
		goto L35
	} else {
		goto L36
	}
L9:
	;
	v268 = int32(0)
	v277 = v48
	v279 = float32(1)
	v283 = v17
	goto L8
L10:
	;
	goto L11
L11:
	;
	v62 = base.F32_mul(v42, float32(0.5))
	if base.F64_gt(base.F64_promote_f32(v62), float64(0.005)) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v66 = float32(0.005)
	goto L14
L13:
	;
	v66 = v62
	goto L14
L14:
	;
	v68 = l8 + int32(104)
	v74 = int32(0)
	v83 = v48
	v85 = float32(1)
	v89 = v17
	v91 = v17
	goto L15
L15:
	;
	if v91 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v268 = v237
	v277 = v246
	v279 = v248
	v283 = v252
	goto L8
L17:
	;
	v263 = v91 + int32(1)
	if v263 != l5 {
		v74 = v237
		v83 = v246
		v85 = v248
		v89 = v252
		v91 = v263
		goto L15
	} else {
		goto L34
	}
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l8)+32))
	v102 = l4 + v91<<(uint(int32(2))%32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102-int32(4))))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v107 = F_FunctionCall2Coll(m, v68, v99, v105, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l1 <= v74 {
		v175 = v74
		v184 = v83
		v186 = v85
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v107 == int32(0) {
		v237 = v74
		v246 = v83
		v248 = v85
		v252 = v89
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v237 = v207
	v246 = v216
	v248 = v218
	v252 = v89 + int32(1)
	goto L17
L24:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v52+v89<<(uint(int32(2))%32)))) = v66
	v207 = v175
	v216 = v184
	v218 = v186
	goto L23
L25:
	;
	v119 = v74
	v128 = v83
	v130 = v85
	goto L26
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l8)+32))
	v146 = v119 << (uint(int32(2)) % 32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0+v146)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l4+v91<<(uint(int32(2))%32))))
	v150 = F_FunctionCall2Coll(m, v68, v144, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L28
	}
L27:
	;
	v175 = l1
	v184 = v165
	v186 = v168
	goto L24
L28:
	;
	if int32(0) <= v150 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v150 != 0 {
		v175 = v119
		v184 = v128
		v186 = v130
		goto L24
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v164 = *(*float32)(unsafe.Add(mBase, uint32(l2+v146)))
	v165 = base.F32_sub(v128, v164)
	v168 = base.F32_mul(v130, base.F32_sub(float32(1), v164))
	v170 = v119 + int32(1)
	if v170 != l1 {
		v119 = v170
		v128 = v165
		v130 = v168
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v158 = *(*float32)(unsafe.Add(mBase, uint32(l2+v146)))
	*(*float32)(unsafe.Add(mBase, uint32(v52+v89<<(uint(int32(2))%32)))) = v158
	v207 = v119 + int32(1)
	v216 = base.F32_sub(v128, v158)
	v218 = v130
	goto L23
L33:
	;
	goto L27
L34:
	;
	goto L16
L35:
	;
	v466 = F_exp(m, base.F64_promote_f32(base.F32_neg(v447)))
	mBase = m.M
	v467 = l1 + v283
	if v467 <= int32(0) {
		v492 = v283
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v296 = (l1 - v268) & int32(3)
	if v296 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v268-l1) {
		v447 = v353
		v449 = v355
		goto L35
	} else {
		goto L44
	}
L38:
	;
	v353 = v277
	v355 = v279
	v358 = v268
	goto L37
L39:
	;
	goto L40
L40:
	;
	v308 = v268
	v312 = v277
	v314 = v279
	v316 = int32(0)
	goto L41
L41:
	;
	v331 = *(*float32)(unsafe.Add(mBase, uint32(l2+v308<<(uint(int32(2))%32))))
	v332 = base.F32_sub(v312, v331)
	v335 = base.F32_mul(v314, base.F32_sub(float32(1), v331))
	v336 = int32(1)
	v337 = v308 + v336
	v339 = v316 + v336
	if v339 != v296 {
		v308 = v337
		v312 = v332
		v314 = v335
		v316 = v339
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v353 = v332
	v355 = v335
	v358 = v337
	goto L37
L43:
	;
	goto L42
L44:
	;
	v390 = v353
	v392 = v355
	v395 = v358
	goto L45
L45:
	;
	v407 = v395 << (uint(int32(2)) % 32)
	v409 = *(*float32)(unsafe.Add(mBase, uint32(l2+v407)))
	v412 = *(*float32)(unsafe.Add(mBase, uint32(v407+(l2+int32(4)))))
	v415 = *(*float32)(unsafe.Add(mBase, uint32(v407+(l2+int32(8)))))
	v418 = *(*float32)(unsafe.Add(mBase, uint32(v407+(l2+int32(12)))))
	v419 = base.F32_sub(base.F32_sub(base.F32_sub(base.F32_sub(v390, v409), v412), v415), v418)
	v420 = float32(1)
	v431 = base.F32_mul(base.F32_mul(base.F32_mul(base.F32_mul(v392, base.F32_sub(v420, v409)), base.F32_sub(v420, v412)), base.F32_sub(v420, v415)), base.F32_sub(v420, v418))
	v433 = v395 + int32(4)
	if v433 != l1 {
		v390 = v419
		v392 = v431
		v395 = v433
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v447 = v419
	v449 = v431
	goto L35
L47:
	;
	goto L46
L48:
	;
	v497 = v492<<(uint(int32(2))%32) + int32(4)
	v498 = F_palloc(m, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L55
	}
L49:
	;
	v472 = base.I32_div_s(l1*int32(100), v467)
	if v283 <= v472 {
		v492 = v283
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_pg_qsort(m, v52, v283, int32(4), int32(1258))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v478 = base.F64_convert_i32_s(l1)
	v486 = base.F64_mul(base.F64_sub(base.F64_sqrt(base.F64_add(base.F64_mul(v478, v478), base.F64_mul(v478, float64(400)))), v478), float64(0.5))
	if base.F64_lt(base.F64_abs(v486), float64(2.147483648e+09)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v490 = base.I32_trunc_f64_s(v486)
	v492 = v490
	goto L48
L53:
	;
	goto L54
L54:
	;
	v492 = int32(-2147483648)
	goto L48
L55:
	;
	v500 = F_palloc(m, v497)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v498))) = int32(1065353216)
	if v492 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	F_pfree(m, v730)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L6
	} else {
		goto L85
	}
L58:
	;
	v711 = v498
	v730 = v500
	goto L57
L59:
	;
	goto L60
L60:
	;
	__phi511 = v498
	__phi525 = int32(1)
	__phi527 = v500
	v511 = __phi511
	v525 = __phi525
	v527 = __phi527
	goto L61
L61:
	;
	if v492 < v525 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v711 = v527
	v730 = v511
	goto L57
L63:
	;
	if v525 != v492 {
		__phi511 = v527
		__phi525 = v525 + int32(1)
		__phi527 = v511
		v511 = __phi511
		v525 = __phi525
		v527 = __phi527
		goto L61
	} else {
		goto L84
	}
L64:
	;
	v537 = v492
	goto L66
L65:
	;
	v537 = v525
	goto L66
L66:
	;
	if v537 < int32(0) {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v541 = v511 - int32(4)
	v542 = *(*float32)(unsafe.Add(mBase, uint32(v511)))
	v547 = *(*float32)(unsafe.Add(mBase, uint32(v52-int32(4)+v525<<(uint(int32(2))%32))))
	v548 = base.F32_sub(float32(1), v547)
	*(*float32)(unsafe.Add(mBase, uint32(v527))) = base.F32_add(base.F32_mul(v542, v548), float32(0))
	v553 = int32(1)
	if v537 != v553 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v564 = v553
	v582 = int32(0)
	goto L71
L69:
	;
	v635 = v553
	goto L70
L70:
	;
	if v537&v553 == int32(0) {
		goto L63
	} else {
		goto L80
	}
L71:
	;
	v589 = float32(0)
	v592 = v564 << (uint(int32(2)) % 32)
	v595 = *(*float32)(unsafe.Add(mBase, uint32(v592+v541)))
	if base.Ui32(v564) < base.Ui32(v525) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v635 = v628
	goto L70
L73:
	;
	v601 = *(*float32)(unsafe.Add(mBase, uint32(v511+v564<<(uint(int32(2))%32))))
	v605 = base.F32_add(base.F32_mul(v601, v548), float32(0))
	goto L75
L74:
	;
	v605 = v589
	goto L75
L75:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v527+v592))) = base.F32_add(base.F32_mul(v595, v547), v605)
	v609 = v564 + int32(1)
	if base.Ui32(v609) < base.Ui32(v525) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v614 = *(*float32)(unsafe.Add(mBase, uint32(v511+v609<<(uint(int32(2))%32))))
	v618 = base.F32_add(base.F32_mul(v614, v548), float32(0))
	goto L78
L77:
	;
	v618 = v589
	goto L78
L78:
	;
	v619 = int32(2)
	v623 = *(*float32)(unsafe.Add(mBase, uint32(v592+v511)))
	*(*float32)(unsafe.Add(mBase, uint32(v527+v609<<(uint(v619)%32)))) = base.F32_add(base.F32_mul(v623, v547), v618)
	v628 = v564 + v619
	v630 = v582 + v619
	if v630 != v537&int32(2147483646) {
		v564 = v628
		v582 = v630
		goto L71
	} else {
		goto L79
	}
L79:
	;
	goto L72
L80:
	;
	if base.Ui32(v635) < base.Ui32(v525) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v667 = *(*float32)(unsafe.Add(mBase, uint32(v511+v635<<(uint(int32(2))%32))))
	v671 = base.F32_add(base.F32_mul(v667, v548), float32(0))
	goto L83
L82:
	;
	v671 = float32(0)
	goto L83
L83:
	;
	v673 = v635 << (uint(int32(2)) % 32)
	v676 = *(*float32)(unsafe.Add(mBase, uint32(v673+v541)))
	*(*float32)(unsafe.Add(mBase, uint32(v527+v673))) = base.F32_add(base.F32_mul(v676, v547), v671)
	goto L63
L84:
	;
	goto L62
L85:
	;
	v741 = F_palloc(m, v497)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v743 = F_palloc(m, v497)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v741))) = int32(1065353216)
	if l1 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if base.F64_gt(base.F64_promote_f32(v447), float64(0.005)) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L89:
	;
	v974 = v741
	v976 = v743
	goto L88
L90:
	;
	goto L91
L91:
	;
	__phi755 = v741
	__phi760 = int32(1)
	__phi771 = v743
	v755 = __phi755
	v760 = __phi760
	v771 = __phi771
	goto L92
L92:
	;
	if v492 < v760 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v974 = v771
	v976 = v755
	goto L88
L94:
	;
	if l1 != v760 {
		__phi755 = v771
		__phi760 = v760 + int32(1)
		__phi771 = v755
		v755 = __phi755
		v760 = __phi760
		v771 = __phi771
		goto L92
	} else {
		goto L116
	}
L95:
	;
	v781 = v492
	goto L97
L96:
	;
	v781 = v760
	goto L97
L97:
	;
	if v781 < int32(0) {
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v784 = *(*float32)(unsafe.Add(mBase, uint32(v755)))
	v789 = *(*float32)(unsafe.Add(mBase, uint32(l2-int32(4)+v760<<(uint(int32(2))%32))))
	v790 = base.F32_sub(float32(1), v789)
	*(*float32)(unsafe.Add(mBase, uint32(v771))) = base.F32_add(base.F32_mul(v784, v790), float32(0))
	if v492 == int32(0) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v798 = v755 - int32(4)
	v799 = int32(1)
	if v781 != v799 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v810 = v799
	v827 = int32(0)
	goto L103
L101:
	;
	v881 = v799
	goto L102
L102:
	;
	if v781&v799 == int32(0) {
		goto L94
	} else {
		goto L112
	}
L103:
	;
	v835 = float32(0)
	v838 = v810 << (uint(int32(2)) % 32)
	v841 = *(*float32)(unsafe.Add(mBase, uint32(v798+v838)))
	if base.Ui32(v810) < base.Ui32(v760) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v881 = v874
	goto L102
L105:
	;
	v847 = *(*float32)(unsafe.Add(mBase, uint32(v755+v810<<(uint(int32(2))%32))))
	v851 = base.F32_add(base.F32_mul(v847, v790), float32(0))
	goto L107
L106:
	;
	v851 = v835
	goto L107
L107:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v771+v838))) = base.F32_add(base.F32_mul(v841, v789), v851)
	v855 = v810 + int32(1)
	if base.Ui32(v855) < base.Ui32(v760) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v860 = *(*float32)(unsafe.Add(mBase, uint32(v755+v855<<(uint(int32(2))%32))))
	v864 = base.F32_add(base.F32_mul(v860, v790), float32(0))
	goto L110
L109:
	;
	v864 = v835
	goto L110
L110:
	;
	v865 = int32(2)
	v869 = *(*float32)(unsafe.Add(mBase, uint32(v838+v755)))
	*(*float32)(unsafe.Add(mBase, uint32(v771+v855<<(uint(v865)%32)))) = base.F32_add(base.F32_mul(v869, v789), v864)
	v874 = v810 + v865
	v876 = v827 + v865
	if v876 != v781&int32(-2) {
		v810 = v874
		v827 = v876
		goto L103
	} else {
		goto L111
	}
L111:
	;
	goto L104
L112:
	;
	if base.Ui32(v881) < base.Ui32(v760) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v913 = *(*float32)(unsafe.Add(mBase, uint32(v755+v881<<(uint(int32(2))%32))))
	v917 = base.F32_add(base.F32_mul(v913, v790), float32(0))
	goto L115
L114:
	;
	v917 = float32(0)
	goto L115
L115:
	;
	v919 = v881 << (uint(int32(2)) % 32)
	v922 = *(*float32)(unsafe.Add(mBase, uint32(v919+v798)))
	*(*float32)(unsafe.Add(mBase, uint32(v771+v919))) = base.F32_add(base.F32_mul(v922, v789), v917)
	goto L94
L116:
	;
	goto L93
L117:
	;
	F_pfree(m, v1201)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L6
	} else {
		goto L140
	}
L118:
	;
	v1201 = v976
	v1215 = v974
	goto L117
L119:
	;
	goto L120
L120:
	;
	v990 = int32(0)
	if v990 <= v492 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v994 = int32(1)
	v996 = v492 + v994
	if v996 <= v994 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v1201 = v974
	v1215 = v976
	goto L117
L124:
	;
	v999 = v994
	goto L126
L125:
	;
	v999 = v996
	goto L126
L126:
	;
	v1003 = F__emscripten_memset_bulkmem(m, v976, base.I32_extend8_s(int32(0)), v999<<(uint(int32(2))%32))
	mBase = m.M
	goto L127
L127:
	;
	v1006 = v990
	v1007 = v996
	v1014 = base.F32_demote_f64(v466)
	goto L128
L128:
	;
	if v492 < v1006 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L123
L130:
	;
	v1162 = int32(1)
	v1165 = v1006 + v1162
	if v1165 != v999 {
		v1006 = v1165
		v1007 = v1007 - v1162
		v1014 = base.F32_mul(v1014, base.F32_div(v447, base.F32_convert_i32_u(v1165)))
		goto L128
	} else {
		goto L139
	}
L131:
	;
	v1036 = v1003 + v1006<<(uint(int32(2))%32)
	v1037 = int32(0)
	if v1006 != v492 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v1045 = v1037
	v1063 = int32(0)
	goto L135
L133:
	;
	v1096 = v1037
	goto L134
L134:
	;
	if v1007&int32(1) == int32(0) {
		goto L130
	} else {
		goto L138
	}
L135:
	;
	v1070 = int32(2)
	v1071 = v1045 << (uint(v1070) % 32)
	v1072 = v1036 + v1071
	v1074 = *(*float32)(unsafe.Add(mBase, uint32(v1071+v974)))
	v1076 = *(*float32)(unsafe.Add(mBase, uint32(v1072)))
	*(*float32)(unsafe.Add(mBase, uint32(v1072))) = base.F32_add(base.F32_mul(v1074, v1014), v1076)
	v1080 = v1071 | int32(4)
	v1081 = v1036 + v1080
	v1083 = *(*float32)(unsafe.Add(mBase, uint32(v1080+v974)))
	v1085 = *(*float32)(unsafe.Add(mBase, uint32(v1081)))
	*(*float32)(unsafe.Add(mBase, uint32(v1081))) = base.F32_add(base.F32_mul(v1083, v1014), v1085)
	v1089 = v1045 + v1070
	v1091 = v1063 + v1070
	if v1091 != v1007&int32(-2) {
		v1045 = v1089
		v1063 = v1091
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v1096 = v1089
	goto L134
L137:
	;
	goto L136
L138:
	;
	v1126 = v1096 << (uint(int32(2)) % 32)
	v1127 = v1036 + v1126
	v1129 = *(*float32)(unsafe.Add(mBase, uint32(v1126+v974)))
	v1131 = *(*float32)(unsafe.Add(mBase, uint32(v1127)))
	*(*float32)(unsafe.Add(mBase, uint32(v1127))) = base.F32_add(base.F32_mul(v1129, v1014), v1131)
	goto L130
L139:
	;
	goto L129
L140:
	;
	v1229 = F_palloc(m, v497)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	if int32(0) <= v492 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v1241 = base.F32_div(float32(1), base.F32_convert_i32_u(l7-int32(2)))
	v1242 = int32(1)
	v1244 = v492 + v1242
	if v1244 <= v1242 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v1490 = float32(0)
	goto L144
L144:
	;
	F_pfree(m, v711)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L6
	} else {
		goto L176
	}
L145:
	;
	v1247 = v1242
	goto L147
L146:
	;
	v1247 = v1244
	goto L147
L147:
	;
	v1248 = int32(0)
	v1254 = v1248
	v1262 = float32(0)
	v1272 = v1248
	goto L148
L148:
	;
	if v44 <= v1254 {
		v1360 = v1254
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1426 = int32(1)
	v1428 = v492 + v1426
	if v1428 <= v1426 {
		goto L167
	} else {
		goto L168
	}
L150:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v1229+v1272<<(uint(int32(2))%32)))) = v1400
	v1424 = v1272 + int32(1)
	if v1424 != v1247 {
		v1254 = v1394
		v1262 = v1402
		v1272 = v1424
		goto L148
	} else {
		goto L166
	}
L151:
	;
	v1385 = float32(0)
	if base.F32_gt(v1262, v1385) == int32(0) {
		v1394 = v1360
		v1400 = v1385
		v1402 = v1262
		goto L150
	} else {
		goto L165
	}
L152:
	;
	v1280 = v44 - v1254
	v1286 = v1254
	v1291 = int32(0)
	goto L155
L153:
	;
	if base.F32_gt(v1262, float32(0)) != 0 {
		goto L162
	} else {
		goto L163
	}
L154:
	;
	if v1291 <= int32(0) {
		v1360 = v1286
		goto L151
	} else {
		goto L160
	}
L155:
	;
	v1313 = l6 + v1286<<(uint(int32(2))%32)
	v1314 = *(*float32)(unsafe.Add(mBase, uint32(v1313)))
	if base.F32_le(v1314, base.F32_convert_i32_u(v1272)) == int32(0) {
		goto L154
	} else {
		goto L157
	}
L156:
	;
	if v1280 <= int32(0) {
		v1360 = v44
		goto L151
	} else {
		goto L159
	}
L157:
	;
	v1318 = int32(1)
	v1321 = v1291 + v1318
	if v1321 != v1280 {
		v1286 = v1286 + v1318
		v1291 = v1321
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v1345 = v44
	v1348 = float32(0)
	v1349 = base.F32_convert_i32_u(v1280 - int32(1))
	goto L153
L160:
	;
	v1333 = base.F32_convert_i32_u(v1291 - int32(1))
	v1336 = *(*float32)(unsafe.Add(mBase, uint32(v1313-int32(4))))
	v1337 = base.F32_sub(v1314, v1336)
	if base.F32_gt(v1337, float32(0)) == int32(0) {
		v1345 = v1286
		v1348 = v1337
		v1349 = v1333
		goto L153
	} else {
		goto L161
	}
L161:
	;
	v1345 = v1286
	v1348 = v1337
	v1349 = base.F32_add(base.F32_div(float32(0.5), v1337), v1333)
	goto L153
L162:
	;
	v1355 = base.F32_add(base.F32_div(float32(0.5), v1262), v1349)
	goto L164
L163:
	;
	v1355 = v1349
	goto L164
L164:
	;
	v1394 = v1345
	v1400 = base.F32_mul(v1241, v1355)
	v1402 = v1348
	goto L150
L165:
	;
	v1394 = v1360
	v1400 = base.F32_div(v1241, v1262)
	v1402 = v1262
	goto L150
L166:
	;
	goto L149
L167:
	;
	v1431 = v1426
	goto L169
L168:
	;
	v1431 = v1428
	goto L169
L169:
	;
	v1437 = int32(0)
	v1444 = float32(0)
	goto L170
L170:
	;
	v1463 = v1437 << (uint(int32(2)) % 32)
	v1465 = *(*float32)(unsafe.Add(mBase, uint32(v1215+v1463)))
	if base.F32_gt(v1465, float32(0)) != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v1490 = v1476
	goto L144
L172:
	;
	v1469 = *(*float32)(unsafe.Add(mBase, uint32(v1463+v1229)))
	v1472 = *(*float32)(unsafe.Add(mBase, uint32(v711+v1463)))
	v1476 = base.F32_add(v1444, base.F32_div(base.F32_mul(base.F32_mul(v1469, base.F32_demote_f64(base.F64_mul(v466, base.F64_promote_f32(v449)))), v1472), v1465))
	goto L174
L173:
	;
	v1476 = v1444
	goto L174
L174:
	;
	v1478 = v1437 + int32(1)
	if v1478 != v1431 {
		v1437 = v1478
		v1444 = v1476
		goto L170
	} else {
		goto L175
	}
L175:
	;
	goto L171
L176:
	;
	F_pfree(m, v1215)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	F_pfree(m, v1229)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_pfree(m, v52)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	v1518 = base.F32_mul(base.F32_sub(float32(1), v49), v1490)
	if base.F32_lt(v1518, float32(0)) != 0 {
		v1526 = float32(0)
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1554 = base.F64_promote_f32(v1526)
	goto L1
L181:
	;
	if base.F32_gt(v1518, float32(1)) == int32(0) {
		v1526 = v1518
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v1526 = float32(1)
	goto L180
}
func F_mcelem_array_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
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
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v119 int32
	_ = v119
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v10 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+11)))
	F_deconstruct_array(m, l0, v24, v25, v26, v21+int32(8), v21+int32(4), v21+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return float64(0)
	} else {
		v37 = int32(0)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
		if v38 <= v37 {
			v95 = v10
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
			F_qsort_arg(m, v103, v95, int32(4), int32(1259), l1)
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return float64(0)
			} else {
				if l8&int32(-2) == int32(2750) {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
					v113 = F_mcelem_array_contain_overlap_selec(m, l2, l3, l4, l5, v112, v95, l8, l1)
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return float64(0)
					} else {
						v138 = v113
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
						F_pfree(m, v139)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return float64(0)
						} else {
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
							F_pfree(m, v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return float64(0)
							} else {
								m.G0 = v21 + int32(16)
								return v138
							}
						}
					}
				} else {
					if l8 != int32(2752) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return float64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = l8
							F_errmsg_internal(m, int32(46700), v21)
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(514559), int32(494), int32(510895))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
						v118 = F_mcelem_array_contained_selec(m, l2, l3, l4, l5, v117, v95, l6, l7, l1)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return float64(0)
						} else {
							v138 = v118
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							F_pfree(m, v139)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return float64(0)
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
								F_pfree(m, v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return float64(0)
								} else {
									m.G0 = v21 + int32(16)
									return v138
								}
							}
						}
					}
				}
			}
		} else {
			v41 = v37
			v51 = v10
			v52 = v38
			v53 = v10
			for {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v41))))
				if v61 != 0 {
					v75 = v51
					v76 = v52
					v77 = int32(1)
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
					v64 = int32(2)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v63+v41<<(uint(v64)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v63+v51<<(uint(v64)%32)))) = v70
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v75 = v51 + int32(1)
					v76 = v74
					v77 = v53
				}
				v79 = v41 + int32(1)
				if v79 < v76 {
					v41 = v79
					v51 = v75
					v52 = v76
					v53 = v77
					continue
				} else {
					break
				}
				break
			}
			if base.B2i32(l8 == int32(2751))&v77 != 0 {
				v138 = float64(0)
				v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				F_pfree(m, v139)
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return float64(0)
				} else {
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
					F_pfree(m, v142)
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return float64(0)
					} else {
						m.G0 = v21 + int32(16)
						return v138
					}
				}
			} else {
				v95 = v75
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				F_qsort_arg(m, v103, v95, int32(4), int32(1259), l1)
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return float64(0)
				} else {
					if l8&int32(-2) == int32(2750) {
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
						v113 = F_mcelem_array_contain_overlap_selec(m, l2, l3, l4, l5, v112, v95, l8, l1)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return float64(0)
						} else {
							v138 = v113
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							F_pfree(m, v139)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return float64(0)
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
								F_pfree(m, v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return float64(0)
								} else {
									m.G0 = v21 + int32(16)
									return v138
								}
							}
						}
					} else {
						if l8 != int32(2752) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return float64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21))) = l8
								F_errmsg_internal(m, int32(46700), v21)
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(514559), int32(494), int32(510895))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							v118 = F_mcelem_array_contained_selec(m, l2, l3, l4, l5, v117, v95, l6, l7, l1)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return float64(0)
							} else {
								v138 = v118
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
								F_pfree(m, v139)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return float64(0)
								} else {
									v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
									F_pfree(m, v142)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return float64(0)
									} else {
										m.G0 = v21 + int32(16)
										return v138
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
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v73*int32(48))+32))
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
	F_errmsg(m, int32(311426), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(520146), int32(1470), int32(507827))
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
	v105 = *(*int32)(unsafe.Add(mBase, _consts[830]))
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
func F_mdmaxcombine(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(131072) - l2&int32(131071)
}
func F_mdprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	if base.Ui64(int64(4294967295)) < base.Ui64(base.I64_extend_i32_s(l3)+base.I64_extend_i32_u(l2)) {
		v108 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v108
L2:
	;
	if l3 <= int32(0) {
		v108 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = l2
	v22 = l3
	goto L4
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v108 = v40
	goto L1
L6:
	;
	v34 = int32(2)
	goto L8
L7:
	;
	v34 = int32(1)
	goto L8
L8:
	;
	v35 = F__mdfd_getseg(m, l0, l1, v21, int32(0), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v39 = int32(0)
	v40 = base.B2i32(v35 != v39)
	if v35 == v39 {
		v108 = v40
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v45 = v21 & int32(131071)
	v50 = int32(131072) - v45
	if base.Ui32(v22) < base.Ui32(v50) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = v22
	goto L14
L13:
	;
	v52 = v50
	goto L14
L14:
	;
	v56 = F_FileAccess(m, v43)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if int32(0) <= v56 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	v99 = v22 - v52
	if int32(0) < v99 {
		v21 = v21 + v52
		v22 = v99
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v72 = int32(4155324)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(167772180)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77+v43*int32(48))))
	v81 = F_posix_fadvise(m, v79, base.I64_extend_i32_u(v45<<(uint(int32(13))%32)), base.I64_extend_i32_u(v52<<(uint(int32(13))%32)), int32(3))
	mBase = m.M
	v83 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	if v81 == int32(27) {
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	goto L5
}
func F_mdunlink(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	if l1 == int32(-1) {
		v12 = l0 + int32(8)
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v13
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v15
		F_mdunlinkfork(m, v5+int32(-16), int32(0), l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v22
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v24
			F_mdunlinkfork(m, v5+int32(-32), int32(1), l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v31
				v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v33
				F_mdunlinkfork(m, v5+int32(-48), int32(2), l2)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v41 = int32(3)
					v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = v44
					F_mdunlinkfork(m, v7, v41, l2)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		v41 = l1
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v42
		v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v44
		F_mdunlinkfork(m, v7, v41, l2)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v6+int32(8), v10, v11, v12, int32(-1), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = v6 + int32(8)
	goto L6
L3:
	;
	v136 = F_unlink(m, l1)
	mBase = m.M
	m.G0 = v6 + int32(80)
	return v136
L4:
	;
	v133 = F_strlen(m, v122)
	mBase = m.M
	goto L3
L6:
	;
	goto L7
L7:
	;
	v27 = int32(1023)
	if (l1^v20)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v126)
	goto L4
L9:
	;
	v107 = v102
	v108 = v103
	v109 = v104
	goto L31
L10:
	;
	if v97 == int32(0) {
		v122 = v95
		v123 = v96
		goto L8
	} else {
		goto L30
	}
L11:
	;
	v95 = v20
	v96 = l1
	v97 = v27
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v20&int32(3) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v64 == int32(0) {
		v122 = v61
		v123 = v62
		goto L8
	} else {
		goto L23
	}
L15:
	;
	v61 = v20
	v62 = l1
	v63 = v27
	v64 = int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v40 = v20
	v41 = l1
	v42 = v27
	goto L18
L18:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v44)
	if v44 == int32(0) {
		v102 = v40
		v103 = v41
		v104 = v42
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v61 = v55
	v62 = v49
	v63 = v51
	v64 = v53
	goto L14
L20:
	;
	v48 = int32(1)
	v49 = v41 + v48
	v51 = v42 - v48
	v52 = int32(0)
	v53 = base.B2i32(v51 != v52)
	v55 = v40 + v48
	if v55&int32(3) == v52 {
		v61 = v55
		v62 = v49
		v63 = v51
		v64 = v53
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v51 != 0 {
		v40 = v55
		v41 = v49
		v42 = v51
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v67 == int32(0) {
		v95 = v61
		v96 = v62
		v97 = v63
		goto L10
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v63) < base.Ui32(int32(4)) {
		v95 = v61
		v96 = v62
		v97 = v63
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v73 = v61
	v74 = v62
	v75 = v63
	goto L26
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v81 = int32(-2139062144)
	if (int32(16843008)-v78|v78)&v81 != v81 {
		v102 = v73
		v103 = v74
		v104 = v75
		goto L9
	} else {
		goto L28
	}
L27:
	;
	v95 = v89
	v96 = v87
	v97 = v91
	goto L10
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v78
	v86 = int32(4)
	v87 = v74 + v86
	v89 = v73 + v86
	v91 = v75 - v86
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		v73 = v89
		v74 = v87
		v75 = v91
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v102 = v95
	v103 = v96
	v104 = v97
	goto L9
L31:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v111)
	if v111 == int32(0) {
		v122 = v107
		v123 = v108
		goto L8
	} else {
		goto L33
	}
L32:
	;
	v122 = v118
	v123 = v116
	goto L8
L33:
	;
	v115 = int32(1)
	v116 = v108 + v115
	v118 = v107 + v115
	v120 = v109 - v115
	if v120 != 0 {
		v107 = v118
		v108 = v116
		v109 = v120
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
}
func F_mdunlinkfork(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(224)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v9+int32(128), v13, v14, v15, v16, v2)
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
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v9 + int32(224)
	return
L4:
	;
	v119 = v9 + int32(200) | int32(4)
	v127 = int32(1)
	goto L30
L5:
	;
	v90 = F_do_truncate(m, v9+int32(128))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L26
	}
L6:
	;
	if v16 != int32(-1) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v16 != int32(-1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v2 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[296])))
	if v22&int32(1) == int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v60 = F_unlink(m, v9+int32(128))
	mBase = m.M
	if int32(0) <= v60 {
		goto L4
	} else {
		goto L17
	}
L12:
	;
	v31 = F_do_truncate(m, v9+int32(128))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+44)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v35
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+42)) = uint16(v2)
	v48 = F_RegisterSyncRequest(m, v9+int32(40), int32(2), int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v34
	if int32(0) <= v31 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	if v34 == int32(44) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v64 == int32(44) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v69 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v69 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v64
	goto L4
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(128)
	F_errmsg(m, int32(311032), v9+int32(32))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(520146), int32(401), int32(327758))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+44)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v94
	v104 = int32(1)
	v106 = F_RegisterSyncRequest(m, v9+int32(40), v104, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v93
	if int32(0) <= v90 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v93 == int32(44) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L4
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(128)
	v138 = F_pg_sprintf(m, v9+int32(40), int32(41608), v9+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v175 == int32(44) {
		goto L3
	} else {
		goto L43
	}
L32:
	;
	if base.B2i32(v16 != int32(-1)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v144 = F_do_truncate(m, v9+int32(40))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v171 = F_unlink(m, v9+int32(40))
	mBase = m.M
	if int32(0) <= v171 {
		v127 = v127 + int32(1)
		goto L30
	} else {
		goto L42
	}
L36:
	;
	if v144 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v149 == int32(44) {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+200)) = int64(0)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v154
	v156 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v9)+216)) = base.I64_extend_i32_u(v127)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+202)) = uint16(v2)
	v165 = F_RegisterSyncRequest(m, v9+int32(200), int32(2), int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L35
L42:
	;
	goto L31
L43:
	;
	v180 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v180 == int32(0) {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(40)
	F_errmsg(m, int32(311032), v9)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(520146), int32(460), int32(327758))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L3
}
func F_merge_map_updates(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
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
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 < v17 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L20
	} else {
		goto L24
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v20 = int32(8)
	v21 = l0 + v20
	v29 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v15 + int32(16)
	return
L6:
	;
	v38 = l1 + v20 + v29<<(uint(int32(3))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 <= v41 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v104 = v29 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v104 < v105 {
		v29 = v104
		goto L6
	} else {
		goto L19
	}
L9:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L10:
	;
	v48 = v41
	goto L11
L11:
	;
	v59 = v21 + v48<<(uint(int32(3))%32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 != v40 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v39
	goto L8
L13:
	;
	v63 = v48 + int32(1)
	if v42 != v63 {
		v48 = v63
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	goto L9
L17:
	;
	if int32(64) <= v42 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v84 = v21 + v42<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v39
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v87 + int32(1)
	goto L8
L19:
	;
	goto L7
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v40
	F_errmsg_internal(m, int32(49870), v15)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(515261), int32(401), int32(370933))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errmsg_internal(m, int32(248589), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(515261), int32(403), int32(370933))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = v14 + l3
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = int32(2)
	v20 = v17 + l3<<(uint(v18)%32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = v22 + l2
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v25 + l2<<(uint(v18)%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 < int32(0) {
		if v29 != int32(-1) {
			if v29 < int32(0) {
				v113 = int32(-1)
				if v21 < int32(0) {
					v130 = v113
				} else {
					if v16&int32(1) != 0 {
						v130 = v113
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v121 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
						v130 = v21
					}
				}
				return v130
			} else {
				if v24&int32(1) != 0 {
					v113 = int32(-1)
					if v21 < int32(0) {
						v130 = v113
					} else {
						if v16&int32(1) != 0 {
							v130 = v113
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v121 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
							v130 = v21
						}
					}
					return v130
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v106 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v104+l3))) = uint8(v106)
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v108+l2))) = uint8(v106)
					return v29
				}
			}
		} else {
			if v21 != int32(-1) {
				if v29 < int32(0) {
					v113 = int32(-1)
					if v21 < int32(0) {
						v130 = v113
					} else {
						if v16&int32(1) != 0 {
							v130 = v113
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v121 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
							v130 = v21
						}
					}
					return v130
				} else {
					if v24&int32(1) != 0 {
						v113 = int32(-1)
						if v21 < int32(0) {
							v130 = v113
						} else {
							if v16&int32(1) != 0 {
								v130 = v113
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v121 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
								v130 = v21
							}
						}
						return v130
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
						v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v106 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v104+l3))) = uint8(v106)
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*uint8)(unsafe.Add(mBase, uint32(v108+l2))) = uint8(v106)
						return v29
					}
				}
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v79
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v83 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v81+l2))) = uint8(v83)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v85+l3<<(uint(int32(2))%32)))) = v79
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v90+l3))) = uint8(v83)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v94 + v83
				return v79
			}
		}
	} else {
		if v21 < int32(0) {
			if v29 != int32(-1) {
				if v29 < int32(0) {
					v113 = int32(-1)
					if v21 < int32(0) {
						v130 = v113
					} else {
						if v16&int32(1) != 0 {
							v130 = v113
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v121 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
							v130 = v21
						}
					}
					return v130
				} else {
					if v24&int32(1) != 0 {
						v113 = int32(-1)
						if v21 < int32(0) {
							v130 = v113
						} else {
							if v16&int32(1) != 0 {
								v130 = v113
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v121 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
								v130 = v21
							}
						}
						return v130
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
						v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v106 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v104+l3))) = uint8(v106)
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*uint8)(unsafe.Add(mBase, uint32(v108+l2))) = uint8(v106)
						return v29
					}
				}
			} else {
				if v21 != int32(-1) {
					if v29 < int32(0) {
						v113 = int32(-1)
						if v21 < int32(0) {
							v130 = v113
						} else {
							if v16&int32(1) != 0 {
								v130 = v113
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v121 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
								v130 = v21
							}
						}
						return v130
					} else {
						if v24&int32(1) != 0 {
							v113 = int32(-1)
							if v21 < int32(0) {
								v130 = v113
							} else {
								if v16&int32(1) != 0 {
									v130 = v113
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v28))) = v21
									v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v121 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v119+l2))) = uint8(v121)
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*uint8)(unsafe.Add(mBase, uint32(v123+l3))) = uint8(v121)
									v130 = v21
								}
							}
							return v130
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
							v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v106 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v104+l3))) = uint8(v106)
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v108+l2))) = uint8(v106)
							return v29
						}
					}
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v79
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v83 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v81+l2))) = uint8(v83)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v85+l3<<(uint(int32(2))%32)))) = v79
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v90+l3))) = uint8(v83)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v94 + v83
					return v79
				}
			}
		} else {
			if v21 == v29 {
				return v29
			} else {
				v36 = int32(-1)
				if v24&int32(1) != 0 {
					v130 = v36
					return v130
				} else {
					if v16&int32(1) != 0 {
						v130 = v36
						return v130
					} else {
						if base.Ui32(v29) < base.Ui32(v21) {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v42)
							v45 = l3 << (uint(int32(2)) % 32)
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v45+v46))) = v29
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v49+l3))) = uint8(v42)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v42)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v55+v45))) = v21
							return v29
						} else {
							v59 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v59)
							v62 = l2 << (uint(int32(2)) % 32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v62+v63))) = v21
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v66+l2))) = uint8(v59)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v59)
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v72+v62))) = v29
							v130 = v21
							return v130
						}
					}
				}
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
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
	var v389 int32
	_ = v389
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
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v588 int32
	_ = v588
	var v589 int64
	_ = v589
	var v591 int64
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int64
	_ = v654
	var v656 int64
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v678 int64
	_ = v678
	var v680 int64
	_ = v680
	var v694 int32
	_ = v694
	var v710 int32
	_ = v710
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int64
	_ = v773
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
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
		goto L31
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v60 + v59
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v66 - base.I64_extend_i32_s(v59)
	if v53 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(0)
	goto L8
L17:
	;
	v153 = v60
	goto L16
L18:
	;
	goto L19
L19:
	;
	v73 = v53 & int32(7)
	if base.Ui32(v53) < base.Ui32(int32(8)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v73 == int32(0) {
		v153 = v120
		goto L16
	} else {
		goto L27
	}
L21:
	;
	v120 = v60
	goto L20
L22:
	;
	goto L23
L23:
	;
	v80 = v60
	v82 = int32(0)
	goto L24
L24:
	;
	v94 = v80 - int32(-8192)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1247]))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1248]))) = v80 + int32(7168)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1249]))) = v80 + int32(6144)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[61]))) = v80 + int32(5120)
	v106 = v80 + int32(3072)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v80 + int32(4096)
	v109 = v80 + int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v106
	v112 = v80 + int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v112
	v116 = v82 + int32(8)
	if v116 != v53&int32(2147483640) {
		v80 = v94
		v82 = v116
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v120 = v94
	goto L20
L26:
	;
	goto L25
L27:
	;
	v134 = v120
	v136 = int32(0)
	goto L28
L28:
	;
	v146 = v134 + int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v146
	v149 = v136 + int32(1)
	if v149 != v73 {
		v134 = v146
		v136 = v149
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v153 = v146
	goto L16
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v192
	v195 = F_GetMemoryChunkSpace(m, v192)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v199 = v197 - base.I64_extend_i32_u(v195)
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+160)) = uint32(v199)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v199 & int64(-4294967296)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1101])))
	if v205 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v239 = base.B2i32(v233 == int32(0))
	goto L41
L34:
	;
	v210 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v210 == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(base.Ui32(v214) >> (uint(int32(10)) % 32))
	F_errmsg_internal(m, int32(142735), v13+int32(-32))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(511995), int32(2094), int32(144430))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L171
	}
L40:
	;
	m.G0 = v15 - int32(-64)
	return
L41:
	;
	if v239&int32(1) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v740
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v742 != 0 {
		goto L155
	} else {
		goto L156
	}
L43:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v427 < v428 {
		goto L82
	} else {
		goto L83
	}
L44:
	;
	v254 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v254 < v255 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v259 = v254
	goto L48
L46:
	;
	goto L47
L47:
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
		goto L53
	}
L48:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v259<<(uint(int32(2))%32))))
	F_LogicalTapeClose(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_pfree(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v278 = v259 + int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v278 < v279 {
		v259 = v278
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L47
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v304
	v309 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+160)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v315 = base.I32_div_s(v310+v311-int32(1), v310)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v315 < v316 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v318 = v315
	goto L56
L55:
	;
	v318 = v316
	goto L56
L56:
	;
	v324 = base.I64_div_s(v309-base.I64_extend_i32_s(v318<<(uint(int32(13))%32)), base.I64_extend_i32_s(v310))
	v325 = int64(0)
	if v325 < v324 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v328 = v324
	goto L59
L58:
	;
	v328 = v325
	goto L59
L59:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1101])))
	if v330 != int32(1) {
		v359 = v310
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(0) < v359 {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v335 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v335 == int32(0) {
		v359 = v337
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v341 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-48)))) = v341
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(base.Ui64(v328) >> (uint(int64(10)) % 64))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v340
	F_errmsg_internal(m, int32(212354), v15)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(511995), int32(2142), int32(144430))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v359 = v357
	goto L60
L67:
	;
	v365 = int32(0)
	goto L70
L68:
	;
	v389 = v359
	goto L69
L69:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v399&int32(1) != 0 {
		goto L43
	} else {
		goto L74
	}
L70:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376+v365<<(uint(int32(2))%32))))
	F_LogicalTapeRewindForRead(m, v380, base.I32_wrap_i64(v328))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L72
	}
L71:
	;
	v389 = v385
	goto L69
L72:
	;
	v384 = v365 + int32(1)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v384 < v385 {
		v365 = v384
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v389 < v402 {
		goto L43
	} else {
		goto L75
	}
L75:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v404 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v405 != int32(-1) {
		goto L43
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
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
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(5)
	goto L40
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v453 + int32(1)
	F_beginmerge(m, l0)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L86
	}
L82:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v431 = F_LogicalTapeCreate(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v447 = base.I32_rem_s(v446, v427)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v445+v447<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v451
	v453 = v446
	goto L81
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v431
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v434+v435<<(uint(int32(2))%32)))) = v431
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v440 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v453 = v444
	goto L81
L86:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v459 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	goto L90
L88:
	;
	goto L89
L89:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(0)
	F_LogicalTapeWrite(m, v725, v13+int32(-16), int32(4))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L150
	}
L90:
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
		goto L92
	}
L91:
	;
	goto L89
L92:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v486 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v501 = F_LogicalTapeRead(m, v480, v13+int32(-16), int32(4))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L99
	}
L94:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v486) < base.Ui32(v489) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_pfree(m, v486)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v491) <= base.Ui32(v486) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = v493
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v486
	goto L93
L98:
	;
	goto L93
L99:
	;
	if v501 != int32(4) {
		goto L39
	} else {
		goto L100
	}
L100:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v505 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v710 {
		goto L90
	} else {
		goto L149
	}
L102:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v508].(func(*base.Module, int32, int32, int32, int32))(m, l0, v13+int32(-16), v480, v505)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v595 = v593 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v595
	if int32(0) < v595 {
		goto L126
	} else {
		goto L127
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v476
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v514 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v514 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v517 = int32(0)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v521) < base.Ui32(int32(2)) {
		v575 = v517
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v588 = v512 + v575<<(uint(int32(4))%32)
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v588))) = v589
	v591 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v588)+8)) = v591
	goto L101
L111:
	;
	v527 = int32(1)
	v529 = v517
	v530 = v517
	goto L112
L112:
	;
	v539 = v529 + int32(2)
	if base.Ui32(v521) <= base.Ui32(v539) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v575 = v553
	goto L110
L114:
	;
	v553 = v527
	goto L116
L115:
	;
	v541 = int32(4)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v548 = m.T0[v547].(func(*base.Module, int32, int32, int32) int32)(m, v512+v527<<(uint(v541)%32), v512+v539<<(uint(v541)%32), l0)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L117
	}
L116:
	;
	v556 = v512 + v553<<(uint(int32(4))%32)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v558 = m.T0[v557].(func(*base.Module, int32, int32, int32) int32)(m, v13+int32(-16), v556, l0)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L121
	}
L117:
	;
	if int32(0) < v548 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v552 = v539
	goto L120
L119:
	;
	v552 = v527
	goto L120
L120:
	;
	v553 = v552
	goto L116
L121:
	;
	if v558 <= int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v575 = v530
	goto L110
L123:
	;
	goto L124
L124:
	;
	v564 = v512 + v530<<(uint(int32(4))%32)
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v556)))
	*(*int64)(unsafe.Add(mBase, uint32(v564))) = v565
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v556)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v564)+8)) = v567
	v569 = int32(1)
	v570 = v553 << (uint(v569) % 32)
	v572 = v570 | v569
	if base.Ui32(v572) < base.Ui32(v521) {
		v527 = v572
		v529 = v570
		v530 = v553
		goto L112
	} else {
		goto L125
	}
L125:
	;
	goto L113
L126:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v603 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v603 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v694 - int32(1)
	goto L101
L129:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	v607 = v595
	goto L131
L131:
	;
	v608 = v595<<(uint(int32(4))%32) + v599
	v609 = int32(0)
	if base.Ui32(v607) < base.Ui32(int32(2)) {
		v664 = v609
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v607 = v606
	goto L131
L133:
	;
	v677 = v599 + v664<<(uint(int32(4))%32)
	v678 = *(*int64)(unsafe.Add(mBase, uint32(v608)))
	*(*int64)(unsafe.Add(mBase, uint32(v677))) = v678
	v680 = *(*int64)(unsafe.Add(mBase, uint32(v608)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v677)+8)) = v680
	goto L128
L134:
	;
	v618 = int32(1)
	v621 = v609
	v622 = v609
	goto L135
L135:
	;
	v628 = v621 + int32(2)
	if base.Ui32(v607) <= base.Ui32(v628) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v664 = v642
	goto L133
L137:
	;
	v642 = v618
	goto L139
L138:
	;
	v630 = int32(4)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v637 = m.T0[v636].(func(*base.Module, int32, int32, int32) int32)(m, v599+v618<<(uint(v630)%32), v599+v628<<(uint(v630)%32), l0)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L140
	}
L139:
	;
	v645 = v599 + v642<<(uint(int32(4))%32)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v647 = m.T0[v646].(func(*base.Module, int32, int32, int32) int32)(m, v608, v645, l0)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L144
	}
L140:
	;
	if int32(0) < v637 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v641 = v628
	goto L143
L142:
	;
	v641 = v618
	goto L143
L143:
	;
	v642 = v641
	goto L139
L144:
	;
	if v647 <= int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v664 = v622
	goto L133
L146:
	;
	goto L147
L147:
	;
	v653 = v599 + v622<<(uint(int32(4))%32)
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v645)))
	*(*int64)(unsafe.Add(mBase, uint32(v653))) = v654
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v645)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v653)+8)) = v656
	v658 = int32(1)
	v659 = v642 << (uint(v658) % 32)
	v661 = v659 | v658
	if base.Ui32(v661) < base.Ui32(v607) {
		v618 = v661
		v621 = v659
		v622 = v642
		goto L135
	} else {
		goto L148
	}
L148:
	;
	goto L136
L149:
	;
	goto L91
L150:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v735 = base.B2i32(v733 == int32(0))
	if v733 != 0 {
		v239 = v735
		goto L41
	} else {
		goto L151
	}
L151:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if int32(1) < v736 {
		v239 = v735
		goto L41
	} else {
		goto L152
	}
L152:
	;
	goto L42
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(4)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v784 <= int32(0) {
		goto L40
	} else {
		goto L166
	}
L154:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_pfree(m, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L4
	} else {
		goto L160
	}
L155:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v743 != int32(-1) {
		goto L154
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	F_LogicalTapeFreeze(m, v740, int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L4
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	goto L153
L160:
	;
	v752 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v752
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	F_LogicalTapeFreeze(m, v756, v13+int32(-16))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v742)))
	*(*int32)(unsafe.Add(mBase, uint32(v742))) = int32(1)
	if v761 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	F_s_lock(m, v742, int32(511995), int32(3034), int32(387191))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L4
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v742+v769<<(uint(int32(3))%32))+72)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v742))) = int32(0)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v742)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v742)+8)) = v777 + int32(1)
	goto L153
L165:
	;
	goto L164
L166:
	;
	v789 = int32(0)
	goto L167
L167:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v800+v789<<(uint(int32(2))%32))))
	F_LogicalTapeClose(m, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L169
	}
L168:
	;
	goto L40
L169:
	;
	v808 = v789 + int32(1)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v808 < v809 {
		v789 = v808
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	F_errmsg_internal(m, int32(387217), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(511995), int32(2862), int32(293155))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v4 = int32(1)
	if int32(20) < l0 {
		v35 = v4
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[703]))
		if base.Ui32(l0-int32(15)) <= base.Ui32(int32(1)) {
			if int32(22) <= v8 {
				v20 = int32(0)
				if l0 == int32(16) {
					v35 = v20
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _consts[363]))
					if v24 != int32(2) {
						v35 = v20
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[704])))
						if v28 != 0 {
							v35 = v20
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, _consts[705]))
							v35 = base.B2i32(l0 == int32(17)) | base.B2i32(v32 <= l0)
						}
					}
				}
			} else {
				v35 = v4
			}
		} else {
			if l0 == int32(20) {
				v20 = int32(0)
				if l0 == int32(16) {
					v35 = v20
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _consts[363]))
					if v24 != int32(2) {
						v35 = v20
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[704])))
						if v28 != 0 {
							v35 = v20
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, _consts[705]))
							v35 = base.B2i32(l0 == int32(17)) | base.B2i32(v32 <= l0)
						}
					}
				}
			} else {
				if v8 == int32(15) {
					v20 = int32(0)
					if l0 == int32(16) {
						v35 = v20
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, _consts[363]))
						if v24 != int32(2) {
							v35 = v20
						} else {
							v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[704])))
							if v28 != 0 {
								v35 = v20
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, _consts[705]))
								v35 = base.B2i32(l0 == int32(17)) | base.B2i32(v32 <= l0)
							}
						}
					}
				} else {
					if v8 <= l0 {
						v35 = v4
					} else {
						v20 = int32(0)
						if l0 == int32(16) {
							v35 = v20
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, _consts[363]))
							if v24 != int32(2) {
								v35 = v20
							} else {
								v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[704])))
								if v28 != 0 {
									v35 = v20
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, _consts[705]))
									v35 = base.B2i32(l0 == int32(17)) | base.B2i32(v32 <= l0)
								}
							}
						}
					}
				}
			}
		}
	}
	return v35
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
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
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l1 - int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0+v14<<(uint(int32(2))%32))+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if l2 != v20 {
		*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l2
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v23
	} else {
	}
	v37 = l3*int32(28) + v19
	v39 = v37 + int32(4)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v40 == int32(0) {
		v43 = int32(4)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v14<<(uint(int32(2))%32))))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		v59 = v50 + v51<<(uint(v43)%32) + v14*int32(100) + int32(88)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
		v62 = F_SearchSysCache4(m, v43, v49, v60, l2, base.I32_extend16_s(l3))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			if v62 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v91
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_errmsg_internal(m, int32(42007), v11)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518120), int32(2938), int32(251721))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v70 = F_SysCacheGetAttrNotNull(m, int32(4), v62, int32(7))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v62)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = F_get_opcode(m, v70)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_fmgr_info_cxt(m, v74, v39, v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return v39
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v39
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v86 int32
	_ = v86
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int64
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int64
	_ = v301
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
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
	if int32(4) <= v101 {
		goto L38
	} else {
		goto L39
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
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v33 < int32(0) {
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
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v38 = v36 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v38) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if int32(1)<<(uint(v38)%32)&int32(163841) == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v47 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v33*int32(24))+12)) = v56
	v60 = v56
	goto L16
L18:
	;
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v58
	v60 = v58
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+36)) = v48
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
	goto L11
L23:
	;
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v66
	goto L23
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v66
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v67
	goto L29
L28:
	;
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v73 - int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v78 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v84 = v18 + int32(8)
	if v77 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v77
	goto L30
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v77
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = v78
	goto L36
L35:
	;
	goto L36
L36:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v86 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+16)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v92
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	goto L10
L37:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v197 != 0 {
		goto L74
	} else {
		goto L75
	}
L38:
	;
	if int32(32) < v8 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L43
L41:
	;
	if base.Ui32(int32(32)) < base.Ui32(v101) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v115 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	F_cparc(m, l0, v115, l2, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115)+4)))
	if v128 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	if v162 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v133 = v131 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v133) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	if int32(1)<<(uint(v133)%32)&int32(163841) == int32(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v142 != 0 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v115)+36))
	if v143 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v155 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+20))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v147+v128*int32(24))+12)) = v151
	v155 = v151
	goto L53
L55:
	;
	goto L56
L56:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+32)) = v153
	v155 = v153
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v143
	goto L59
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v115)+32)) = int64(0)
	goto L48
L60:
	;
	if v161 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+20)) = v161
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+16)) = v161
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = v162
	goto L66
L65:
	;
	goto L66
L66:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v168 - int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v115)+24))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v115)+28))
	if v173 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v179 = v115 + int32(8)
	if v172 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v172
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = v172
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+28)) = v173
	goto L73
L72:
	;
	goto L73
L73:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v181 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = int32(0)
	v187 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v179)+16)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v179)+8)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v179))) = v187
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+16)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v115
	goto L47
L74:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_sortouts(m, l0, l1)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	F_sortouts(m, l0, l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	if v205 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v206 == int32(0) {
		v349 = v206
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v349 == int32(0) {
		goto L1
	} else {
		goto L134
	}
L82:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v209 == int32(0) {
		v349 = v206
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v215 = v206
	v216 = v209
	goto L84
L84:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v220 < v222 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v349 = v340
	goto L81
L86:
	;
	if v340 == int32(0) {
		v349 = v340
		goto L81
	} else {
		goto L132
	}
L87:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v340 = v215
	v341 = v338
	goto L86
L88:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	if v314 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L89:
	;
	if v222 < v220 {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215)+4)))
	v226 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216)+4)))
	if v225 < v226 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	if v226 < v225 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v229 < v230 {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	if v230 < v229 {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215)+4)))
	if v242 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v340 = v234
	v341 = v233
	goto L86
L96:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	if v276 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L97:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v247 = v245 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v247) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	if int32(1)<<(uint(v247)%32)&int32(163841) == int32(0) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v256 != 0 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v215)+36))
	if v257 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v269 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261+v242*int32(24))+12)) = v265
	v269 = v265
	goto L101
L103:
	;
	goto L104
L104:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+32)) = v267
	v269 = v267
	goto L101
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+36)) = v257
	goto L107
L106:
	;
	goto L107
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v215)+32)) = int64(0)
	goto L96
L108:
	;
	if v275 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v275
	goto L108
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v275
	goto L108
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+20)) = v276
	goto L114
L113:
	;
	goto L114
L114:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v282 - int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	if v287 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v293 = v215 + int32(8)
	if v286 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = v286
	goto L115
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+24)) = v286
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+28)) = v287
	goto L121
L120:
	;
	goto L121
L121:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v295 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = int32(0)
	v301 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v293)+16)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v293)+8)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = v301
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v215
	goto L95
L122:
	;
	if v313 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+20)) = v313
	goto L122
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+16)) = v313
	goto L122
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+20)) = v314
	goto L128
L127:
	;
	goto L128
L128:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+12)) = v320 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = l2
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v325
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v329 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+20)) = v215
	goto L131
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v215
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v332 + int32(1)
	v340 = v313
	v341 = v216
	goto L86
L132:
	;
	if v341 != 0 {
		v215 = v340
		v216 = v341
		goto L84
	} else {
		goto L133
	}
L133:
	;
	goto L85
L134:
	;
	v358 = v349
	goto L135
L135:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v358)+16))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)+20))
	if v364 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L1
L137:
	;
	if v363 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+20)) = v363
	goto L137
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v363
	goto L137
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+20)) = v364
	goto L143
L142:
	;
	goto L143
L143:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+12)) = v370 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+8)) = l2
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = v375
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v379 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+20)) = v358
	goto L146
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v358
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v382 + int32(1)
	if v363 != 0 {
		v358 = v363
		goto L135
	} else {
		goto L147
	}
L147:
	;
	goto L136
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
		F_errmsg_internal(m, int32(459052), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errfinish(m, int32(515581), int32(208), int32(329210))
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
