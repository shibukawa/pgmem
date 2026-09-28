package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CompareCandidateDistances(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.F32_lt(v7, v9) != 0 {
		v23 = int32(1)
	} else {
		if base.F32_gt(v7, v9) != 0 {
			v23 = int32(-1)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if base.Ui32(v14) < base.Ui32(v15) {
				v23 = int32(1)
			} else {
				if base.Ui32(v15) < base.Ui32(v14) {
					v20 = int32(-1)
				} else {
					v20 = int32(0)
				}
				v23 = v20
			}
		}
	}
	return v23
}
func F_ComputeXidHorizons(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
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
	var v467 int32
	_ = v467
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int64
	_ = v522
	var v526 int64
	_ = v526
	var v529 int64
	_ = v529
	var v533 int64
	_ = v533
	var v536 int64
	_ = v536
	var v540 int64
	_ = v540
	var v543 int64
	_ = v543
	var v544 int32
	_ = v544
	var v548 int64
	_ = v548
	var v550 int32
	_ = v550
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v556 int64
	_ = v556
	var v558 int64
	_ = v558
	var v560 int32
	_ = v560
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int32
	_ = v570
	var v571 int64
	_ = v571
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v584 int32
	_ = v584
	v2 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[1])))
	if v27 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[2]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	v48 = F_LWLockAcquire(m, v44+int32(512), int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[4]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+308))
	v35 = base.B2i32(v33 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[1])) = uint8(v35)
	v37 = v35
	goto L4
L3:
	;
	v37 = v2
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[5]))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v52
	v54 = int32(3)
	v57 = base.I32_wrap_i64(v52) + int32(1)
	if base.Ui32(v57) <= base.Ui32(v54) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v60 = v54
	goto L9
L8:
	;
	v60 = v57
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v60
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[6]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	if v66 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v67 = v66
	goto L12
L11:
	;
	v67 = v60
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if int32(0) < v75 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[7]))
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[2]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[8]))
	v87 = v60
	v89 = v60
	v90 = v2
	v92 = v60
	goto L16
L14:
	;
	goto L15
L15:
	;
	if v37 != 0 {
		goto L59
	} else {
		goto L60
	}
L16:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v90))))
	v112 = v90 << (uint(int32(2)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(36)+v112)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v40)))
	v119 = v85 + v114*int32(768)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+52))
	if v120 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	goto L15
L18:
	;
	v194 = v90 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v194 < v195 {
		v87 = v188
		v89 = v190
		v90 = v194
		v92 = v192
		goto L16
	} else {
		goto L55
	}
L19:
	;
	v140 = base.B2i32(base.Ui32(v134) < base.Ui32(int32(3)))
	if base.B2i32(v140 == int32(0))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v87)) != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	if int32(base.Ui32(v120-v116)>>(uint(int32(31))%32)) != 0 {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	if v116 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v116 != 0 {
		v134 = v116
		goto L19
	} else {
		goto L31
	}
L24:
	;
	v134 = v120
	goto L19
L25:
	;
	goto L26
L26:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v120))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v116)) != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v120) < base.Ui32(v116) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v129 = v120
	goto L30
L29:
	;
	v129 = v116
	goto L30
L30:
	;
	v134 = v129
	goto L19
L31:
	;
	v188 = v87
	v190 = v89
	v192 = v92
	goto L18
L32:
	;
	v133 = v120
	goto L34
L33:
	;
	v133 = v116
	goto L34
L34:
	;
	v134 = v133
	goto L19
L35:
	;
	v146 = int32(base.Ui32(v87-v134) >> (uint(int32(31)) % 32))
	goto L37
L36:
	;
	v146 = base.B2i32(base.Ui32(v87) < base.Ui32(v134))
	goto L37
L37:
	;
	if v146 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v147 = v87
	goto L40
L39:
	;
	v147 = v134
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v147
	if v110&int32(18) != 0 {
		v188 = v147
		v190 = v89
		v192 = v92
		goto L18
	} else {
		goto L41
	}
L41:
	;
	if base.B2i32(v140 == int32(0))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v89)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v160 = int32(base.Ui32(v89-v134) >> (uint(int32(31)) % 32))
	goto L44
L43:
	;
	v160 = base.B2i32(base.Ui32(v89) < base.Ui32(v134))
	goto L44
L44:
	;
	if v160 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v161 = v89
	goto L47
L46:
	;
	v161 = v134
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v161
	v163 = int32(0)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	if base.B2i32(v81 == v163)|base.B2i32(v165 == v81)|(int32(base.Ui32(v110)>>(uint(int32(5))%32))|v37)&int32(1) == v163 {
		v188 = v147
		v190 = v161
		v192 = v92
		goto L18
	} else {
		goto L48
	}
L48:
	;
	if base.B2i32(v140 == int32(0))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v92)) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v185 = int32(base.Ui32(v92-v134) >> (uint(int32(31)) % 32))
	goto L51
L50:
	;
	v185 = base.B2i32(base.Ui32(v92) < base.Ui32(v134))
	goto L51
L51:
	;
	if v185 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v186 = v92
	goto L54
L53:
	;
	v186 = v134
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v186
	v188 = v147
	v190 = v161
	v192 = v186
	goto L18
L55:
	;
	goto L17
L56:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v356 != 0 {
		goto L108
	} else {
		goto L109
	}
L57:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v336 != 0 {
		goto L90
	} else {
		goto L91
	}
L58:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v335 = v333
	goto L57
L59:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v221 = int32(0)
	v224 = base.AtomicRmwOr32(m, v221, int32(_a_F_ComputeXidHorizons_0), v221)
	if v219 <= v220 {
		v287 = v221
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v327+int32(512))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L89
	}
L62:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v289+int32(512))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L70
	}
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[9]))
	v231 = v220
	goto L64
L64:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v231))))
	if v252 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v287 = int32(0)
	goto L62
L66:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[10]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v231<<(uint(int32(2))%32))))
	v287 = v260
	goto L62
L67:
	;
	goto L68
L68:
	;
	v262 = v231 + int32(1)
	if v262 != v219 {
		v231 = v262
		goto L64
	} else {
		goto L69
	}
L69:
	;
	goto L65
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v294 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v287 == int32(0) {
		goto L58
	} else {
		goto L74
	}
L72:
	;
	v308 = v287
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v310 == int32(0) {
		v335 = v287
		goto L57
	} else {
		goto L81
	}
L74:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v294))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v287)) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v306 = int32(base.Ui32(v294-v287) >> (uint(int32(31)) % 32))
	goto L77
L76:
	;
	v306 = base.B2i32(base.Ui32(v294) < base.Ui32(v287))
	goto L77
L77:
	;
	if v306 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v307 = v294
	goto L80
L79:
	;
	v307 = v287
	goto L80
L80:
	;
	v308 = v307
	goto L73
L81:
	;
	if v287 == int32(0) {
		v335 = v310
		goto L57
	} else {
		goto L82
	}
L82:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v310))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v287)) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v324 = int32(base.Ui32(v310-v287) >> (uint(int32(31)) % 32))
	goto L85
L84:
	;
	v324 = base.B2i32(base.Ui32(v310) < base.Ui32(v287))
	goto L85
L85:
	;
	if v324 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v325 = v310
	goto L88
L87:
	;
	v325 = v287
	goto L88
L88:
	;
	v335 = v325
	goto L57
L89:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v356 = v332
	goto L56
L90:
	;
	if v287 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v351 = v287
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v351
	v356 = v335
	goto L56
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v336
	v356 = v335
	goto L56
L94:
	;
	goto L95
L95:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v336))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v287)) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v349 = int32(base.Ui32(v336-v287) >> (uint(int32(31)) % 32))
	goto L98
L97:
	;
	v349 = base.B2i32(base.Ui32(v336) < base.Ui32(v287))
	goto L98
L98:
	;
	if v349 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v350 = v336
	goto L101
L100:
	;
	v350 = v287
	goto L101
L101:
	;
	v351 = v350
	goto L92
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v464
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v467 != 0 {
		goto L154
	} else {
		goto L155
	}
L103:
	;
	v447 = int32(0)
	if v435 == v447 {
		goto L141
	} else {
		goto L142
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v415
	if v416 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v436
	if v433 != 0 {
		goto L103
	} else {
		goto L137
	}
L106:
	;
	if v418 == int32(0) {
		goto L104
	} else {
		goto L130
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v356
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v415 = v356
	v416 = v414
	v418 = v413
	goto L106
L108:
	;
	if v375 == int32(0) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	v389 = v375
	goto L110
L110:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v390 == int32(0) {
		v406 = v375
		goto L118
	} else {
		goto L119
	}
L111:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v375))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v356)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v387 = int32(base.Ui32(v356-v375) >> (uint(int32(31)) % 32))
	goto L114
L113:
	;
	v387 = base.B2i32(base.Ui32(v356) < base.Ui32(v375))
	goto L114
L114:
	;
	if v387 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v388 = v356
	goto L117
L116:
	;
	v388 = v375
	goto L117
L117:
	;
	v389 = v388
	goto L110
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v406
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v389 == int32(0) {
		v433 = v406
		v435 = v409
		v436 = v409
		goto L105
	} else {
		goto L129
	}
L119:
	;
	if v375 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v406 = v390
	goto L118
L121:
	;
	goto L122
L122:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v390))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v375)) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v404 = int32(base.Ui32(v390-v375) >> (uint(int32(31)) % 32))
	goto L125
L124:
	;
	v404 = base.B2i32(base.Ui32(v390) < base.Ui32(v375))
	goto L125
L125:
	;
	if v404 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v405 = v390
	goto L128
L127:
	;
	v405 = v375
	goto L128
L128:
	;
	v406 = v405
	goto L118
L129:
	;
	v415 = v389
	v416 = v406
	v418 = v409
	goto L106
L130:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v415))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v418)) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v430 = int32(base.Ui32(v415-v418) >> (uint(int32(31)) % 32))
	goto L133
L132:
	;
	v430 = base.B2i32(base.Ui32(v415) < base.Ui32(v418))
	goto L133
L133:
	;
	if v430 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v431 = v415
	goto L136
L135:
	;
	v431 = v418
	goto L136
L136:
	;
	v433 = v416
	v435 = v418
	v436 = v431
	goto L105
L137:
	;
	v462 = int32(0)
	v463 = v436
	v464 = v435
	v465 = int32(1)
	goto L102
L138:
	;
	v443 = int32(0)
	v462 = v443
	v463 = v415
	v464 = v443
	v465 = int32(1)
	goto L102
L139:
	;
	goto L140
L140:
	;
	v462 = v416
	v463 = v415
	v464 = v416
	v465 = int32(0)
	goto L102
L141:
	;
	v462 = v433
	v463 = v436
	v464 = v433
	v465 = v447
	goto L102
L142:
	;
	goto L143
L143:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v433))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v435)) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v459 = int32(base.Ui32(v433-v435) >> (uint(int32(31)) % 32))
	goto L146
L145:
	;
	v459 = base.B2i32(base.Ui32(v433) < base.Ui32(v435))
	goto L146
L146:
	;
	if v459 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v460 = v433
	goto L149
L148:
	;
	v460 = v435
	goto L149
L149:
	;
	v462 = v433
	v463 = v436
	v464 = v460
	v465 = v447
	goto L102
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v519
	v522 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v526 = v522 + base.I64_extend_i32_s(v518-base.I32_wrap_i64(v522))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[11])) = v526
	v529 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v533 = v529 + base.I64_extend_i32_s(v464-base.I32_wrap_i64(v529))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[12])) = v533
	v536 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v540 = v536 + base.I64_extend_i32_s(v462-base.I32_wrap_i64(v536))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[13])) = v540
	v543 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v548 = v543 + base.I64_extend_i32_s(v544-base.I32_wrap_i64(v543))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[14])) = v548
	v550 = int32(_a_F_ComputeXidHorizons_1)
	v552 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[15]))
	if base.Ui64(v552) < base.Ui64(v526) {
		goto L186
	} else {
		goto L187
	}
L151:
	;
	if v465 != 0 {
		v518 = v505
		v519 = v504
		goto L150
	} else {
		goto L179
	}
L152:
	;
	v501 = int32(0)
	if v464 == v501 {
		v518 = v501
		v519 = v462
		goto L150
	} else {
		goto L178
	}
L153:
	;
	if v464 == int32(0) {
		v504 = v485
		v505 = v463
		goto L151
	} else {
		goto L168
	}
L154:
	;
	if v463 == int32(0) {
		v485 = v467
		goto L153
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if v463 == int32(0) {
		goto L152
	} else {
		goto L167
	}
L157:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v467))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v463)) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	if base.Ui32(v467) < base.Ui32(v463) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L160
L160:
	;
	if int32(base.Ui32(v467-v463)>>(uint(int32(31))%32)) != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v478 = v467
	goto L163
L162:
	;
	v478 = v463
	goto L163
L163:
	;
	v485 = v478
	goto L153
L164:
	;
	v482 = v467
	goto L166
L165:
	;
	v482 = v463
	goto L166
L166:
	;
	v485 = v482
	goto L153
L167:
	;
	v485 = v463
	goto L153
L168:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v485))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v464)) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if base.Ui32(v485) < base.Ui32(v464) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	if int32(base.Ui32(v485-v464)>>(uint(int32(31))%32)) != 0 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v496 = v485
	goto L174
L173:
	;
	v496 = v464
	goto L174
L174:
	;
	v504 = v496
	v505 = v463
	goto L151
L175:
	;
	v500 = v485
	goto L177
L176:
	;
	v500 = v464
	goto L177
L177:
	;
	v504 = v500
	v505 = v463
	goto L151
L178:
	;
	v504 = v464
	v505 = v501
	goto L151
L179:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v504))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v462)) != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v515 = int32(base.Ui32(v504-v462) >> (uint(int32(31)) % 32))
	goto L182
L181:
	;
	v515 = base.B2i32(base.Ui32(v504) < base.Ui32(v462))
	goto L182
L182:
	;
	if v515 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v516 = v504
	goto L185
L184:
	;
	v516 = v462
	goto L185
L185:
	;
	v518 = v505
	v519 = v516
	goto L150
L186:
	;
	v554 = v526
	goto L188
L187:
	;
	v554 = v552
	goto L188
L188:
	;
	if base.I32_wrap_i64(v552) != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v556 = v554
	goto L191
L190:
	;
	v556 = v526
	goto L191
L191:
	;
	if base.I32_wrap_i64(v526) != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v558 = v556
	goto L194
L193:
	;
	v558 = v552
	goto L194
L194:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[15])) = v558
	v560 = int32(_a_F_ComputeXidHorizons_2)
	v562 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[16]))
	if base.Ui64(v562) < base.Ui64(v533) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v564 = v533
	goto L197
L196:
	;
	v564 = v562
	goto L197
L197:
	;
	if base.I32_wrap_i64(v562) != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v566 = v564
	goto L200
L199:
	;
	v566 = v533
	goto L200
L200:
	;
	if base.I32_wrap_i64(v533) != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v568 = v566
	goto L203
L202:
	;
	v568 = v562
	goto L203
L203:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[16])) = v568
	v570 = int32(_a_F_ComputeXidHorizons_3)
	v571 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[18])) = v548
	if base.Ui64(v571) < base.Ui64(v540) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v576 = v540
	goto L206
L205:
	;
	v576 = v571
	goto L206
L206:
	;
	if base.I32_wrap_i64(v571) != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v578 = v576
	goto L209
L208:
	;
	v578 = v540
	goto L209
L209:
	;
	if base.I32_wrap_i64(v540) != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v580 = v578
	goto L212
L211:
	;
	v580 = v571
	goto L212
L212:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[17])) = v580
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[20])) = v584
	return
}
func F_ConditionVariableCancelSleep(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0]))
	if v7 != 0 {
		v10 = base.AtomicRmwXchg32(m, v7, int32(0), int32(1))
		if v10 != 0 {
			F_s_lock(m, v7, int32(_a_F_ConditionVariableCancelSleep_0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[2]))
				v21 = v16 + v18*int32(768)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+360))
				if v22 == int32(0) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+356))
					if v25 != 0 {
						v31 = v25
						*(*int32)(unsafe.Add(mBase, uint32(v16+v22*int32(768))+356)) = v31
						v36 = v22
						v37 = v31
						if v37 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v36
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							*(*int32)(unsafe.Add(mBase, uint32(v43+v37*int32(768))+360)) = v36
						}
						*(*int64)(unsafe.Add(mBase, uint32(v21)+356)) = int64(0)
					} else {
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+356))
					if v22 != int32(-1) {
						v31 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v16+v22*int32(768))+356)) = v31
						v36 = v22
						v37 = v31
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v26
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+360))
						v36 = v30
						v37 = v26
					}
					if v37 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v36
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						*(*int32)(unsafe.Add(mBase, uint32(v43+v37*int32(768))+360)) = v36
					}
					*(*int64)(unsafe.Add(mBase, uint32(v21)+356)) = int64(0)
				}
				v52 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v52))
				*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0])) = v52
				return
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[2]))
			v21 = v16 + v18*int32(768)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+360))
			if v22 == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+356))
				if v25 != 0 {
					v31 = v25
					*(*int32)(unsafe.Add(mBase, uint32(v16+v22*int32(768))+356)) = v31
					v36 = v22
					v37 = v31
					if v37 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v36
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						*(*int32)(unsafe.Add(mBase, uint32(v43+v37*int32(768))+360)) = v36
					}
					*(*int64)(unsafe.Add(mBase, uint32(v21)+356)) = int64(0)
				} else {
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+356))
				if v22 != int32(-1) {
					v31 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v16+v22*int32(768))+356)) = v31
					v36 = v22
					v37 = v31
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v26
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+360))
					v36 = v30
					v37 = v26
				}
				if v37 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v36
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					*(*int32)(unsafe.Add(mBase, uint32(v43+v37*int32(768))+360)) = v36
				}
				*(*int64)(unsafe.Add(mBase, uint32(v21)+356)) = int64(0)
			}
			v52 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v52))
			*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0])) = v52
			return
		}
	} else {
		return
	}
}
func F_ConditionalXactLockTableWait(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v18 = F_LockAcquireExtended(m, v7, int32(5), v3, int32(1), v3, l1)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v86
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		v86 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = F_LockRelease(m, v7, int32(5), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		v86 = int32(1)
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v33
	v40 = int32(0)
	v45 = F_LockAcquireExtended(m, v7, int32(5), v40, int32(1), v40, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v45 == int32(0) {
		v86 = v40
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = v33
	goto L11
L11:
	;
	v55 = F_LockRelease(m, v7, int32(5), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v86 = v57 ^ int32(1)
	goto L1
L13:
	;
	goto L12
L14:
	;
	v57 = F_TransactionIdIsInProgress(m, v49)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v57 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalXactLockTableWait[0]))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_pg_usleep(m, int32(1000))
	mBase = m.M
	v67 = F_SubTransGetTopmostTransaction(m, v49)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v67
	v75 = int32(0)
	v78 = F_LockAcquireExtended(m, v7, int32(5), v75, int32(1), v75, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v78 != 0 {
		v49 = v67
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L13
}
func F_CountOtherDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[0]))
	v36 = int32(0)
	goto L2
L1:
	;
	m.G0 = v17 + int32(48)
	return v391
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[1]))
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v379+int32(512))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L73
	}
L4:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	v52 = F_LWLockAcquire(m, v48+int32(512), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if int32(0) < v54 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v57 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[3]))
	v65 = v57
	v66 = v57
	v69 = v57
	goto L13
L11:
	;
	goto L12
L12:
	;
	goto L3
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(36)+v65<<(uint(int32(2))%32))))
	v82 = v59 + v79*int32(768)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v83 != l0 {
		v120 = v66
		v121 = v69
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v127+int32(512))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L23
	}
L15:
	;
	v123 = v65 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v123 < v124 {
		v65 = v123
		v66 = v120
		v69 = v121
		goto L13
	} else {
		goto L22
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[4]))
	if v82 == v86 {
		v120 = v66
		v121 = v69
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	if v88 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v91 = int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v92 + v91
	v120 = v66
	v121 = v91
	goto L15
L19:
	;
	goto L20
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[5]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v65))))
	v101 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v102 + v101
	if base.B2i32(v100&v101 == int32(0))|base.B2i32(int32(9) < v66) != 0 {
		v120 = v66
		v121 = v101
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17+v66<<(uint(int32(2))%32)))) = v116
	v120 = v66 + int32(1)
	v121 = v101
	goto L15
L22:
	;
	goto L14
L23:
	;
	if v121 == int32(0) {
		v391 = v121
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v134 = int32(0)
	if v134 < v120 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v140 = v134
	goto L28
L26:
	;
	goto L27
L27:
	;
	v174 = int32(0)
	v176 = m.G0
	v178 = v176 - int32(32)
	m.G0 = v178
	v182 = F_errstart(m, int32(14), v174)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L7
	} else {
		goto L31
	}
L28:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v17+v140<<(uint(int32(2))%32))))
	v156 = F_pgmem_kill(m, v154, int32(15))
	mBase = m.M
	v158 = v140 + int32(1)
	if v158 != v120 {
		v140 = v158
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	goto L29
L31:
	;
	if v182 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_CountOtherDBBackends_0), v178+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	v200 = F_LWLockAcquire(m, v196+int32(_a_F_CountOtherDBBackends_1), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	F_errfinish(m, int32(_a_F_CountOtherDBBackends_2), int32(1423), int32(_a_F_CountOtherDBBackends_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[6]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if int32(0) < v204 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	m.G0 = v178 + int32(32)
	F_pg_usleep(m, int32(_a_F_CountOtherDBBackends_4))
	mBase = m.M
	v374 = v36 + int32(1)
	if v374 != int32(50) {
		v36 = v374
		goto L2
	} else {
		goto L72
	}
L39:
	;
	v210 = v203
	v212 = v174
	v213 = v204
	v215 = v174
	goto L42
L40:
	;
	goto L41
L41:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v349+int32(_a_F_CountOtherDBBackends_1))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L71
	}
L42:
	;
	v223 = v210 + v212*int32(1488)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+16)))
	if v224 != int32(1) {
		v317 = v210
		v318 = v213
		v319 = v215
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v325+int32(_a_F_CountOtherDBBackends_1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L65
	}
L44:
	;
	v322 = v212 + int32(1)
	if v322 < v318 {
		v210 = v317
		v212 = v322
		v213 = v318
		v215 = v319
		goto L42
	} else {
		goto L64
	}
L45:
	;
	v228 = v223 + int32(16)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+208)))
	if v229&int32(4) == int32(0) {
		v317 = v210
		v318 = v213
		v319 = v215
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	v240 = F_LWLockAcquire(m, v236+int32(512), int32(1))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v242 = int32(0)
	if v234 == v242 {
		v280 = v242
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v309+int32(512))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L63
	}
L49:
	;
	if v280 == int32(0) {
		v307 = v215
		goto L48
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[0]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251 <= int32(0) {
		v280 = v242
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[3]))
	v260 = int32(0)
	goto L53
L53:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v250+int32(36)+v260<<(uint(int32(2))%32))))
	v271 = v258 + v268*int32(768)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	if v272 == v234 {
		v280 = v271
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v280 = int32(0)
	goto L50
L55:
	;
	v275 = v260 + int32(1)
	if v275 != v251 {
		v260 = v275
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	if v286 != l0 {
		v307 = v215
		goto L48
	} else {
		goto L58
	}
L58:
	;
	v288 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)) = uint8(v288)
	v293 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	if v293 == int32(0) {
		v307 = v288
		goto L48
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v234
	F_errmsg_internal(m, int32(_a_F_CountOtherDBBackends_5), v178)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_CountOtherDBBackends_2), int32(1449), int32(_a_F_CountOtherDBBackends_3))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v307 = v288
	goto L48
L63:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[6]))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v317 = v315
	v318 = v316
	v319 = v307
	goto L44
L64:
	;
	goto L43
L65:
	;
	if v319 == int32(0) {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[7])))
	if v334 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L38
L68:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v338+int32(28)))) = int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[9]))
	v347 = F_pgmem_kill(m, v345, int32(10))
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L38
L72:
	;
	v391 = v121
	goto L1
L73:
	;
	v391 = int32(0)
	goto L1
}
func F_coerce_null_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v15 = F_getBaseTypeAndTypmod(m, l0, v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v22 = F_makeConst(m, v15, v19, l2, l3, int64(0), int32(1), l4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if l0 != v15 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v26 = int32(0)
				v30 = F_coerce_to_domain(m, v22, v15, v25, l0, v26, int32(2), int32(-1), v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = v30
					m.G0 = v10 + int32(16)
					return v32
				}
			} else {
				v32 = v22
				m.G0 = v10 + int32(16)
				return v32
			}
		}
	}
}
func F_collations_agree_on_equality(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if base.B2i32(l1 == v3)|(base.B2i32(l0 == v3)|base.B2i32(l0 == l1)) != 0 {
		v48 = int32(1)
		m.G0 = v7 + int32(32)
		return v48 & int32(1)
	} else {
		v19 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(l0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_collations_agree_on_equality_0), v7)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_collations_agree_on_equality_1), int32(1286), int32(_a_F_collations_agree_on_equality_2))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26)+77)))
				F_ReleaseCatCache(m, v19)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v28 != int32(1) {
						v48 = int32(0)
						m.G0 = v7 + int32(32)
						return v48 & int32(1)
					} else {
						v36 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(l1))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
									F_errmsg_internal(m, int32(_a_F_collations_agree_on_equality_0), v7+int32(16))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_collations_agree_on_equality_1), int32(1286), int32(_a_F_collations_agree_on_equality_2))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v41)+77)))
								F_ReleaseCatCache(m, v36)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v48 = v43
									m.G0 = v7 + int32(32)
									return v48 & int32(1)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_colorTrgmInfoPenaltyCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F32_ne(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F32_lt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_combo_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v7, v8, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_combo_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v7, v8, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_combo_encrypt_len(m *base.Module, l0 int32, l1 int32) int32 {
	return l1 + int32(512)
}
func F_comp_ptrgm(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_comp_ptrgm[0]))
	v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v17 = v6
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v17 = base.B2i32(v11 < v10) - base.B2i32(v10 < v11)
		}
		return v17
	}
}
func F_comparePairs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v25 int32
	_ = v25
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v85
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	goto L4
L4:
	;
	if base.Ui32(v6) < base.Ui32(v5) {
		goto L28
	} else {
		goto L29
	}
L5:
	;
	if v71 != 0 {
		v85 = v71
		goto L1
	} else {
		goto L23
	}
L6:
	;
	v71 = int32(0)
	goto L5
L7:
	;
	v45 = v40
	v46 = v41
	v47 = v42
	goto L17
L8:
	;
	if (v8|v9)&int32(3) != 0 {
		v40 = v8
		v41 = v9
		v42 = v5
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v33 = v8
	v34 = v9
	v35 = v5
	goto L10
L10:
	;
	if v35 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v17 = v8
	v18 = v9
	v19 = v5
	goto L12
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != v23 {
		v40 = v17
		v41 = v18
		v42 = v19
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v33 = v28
	v34 = v26
	v35 = v30
	goto L10
L14:
	;
	v25 = int32(4)
	v26 = v18 + v25
	v28 = v17 + v25
	v30 = v19 - v25
	if base.Ui32(int32(3)) < base.Ui32(v30) {
		v17 = v28
		v18 = v26
		v19 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v40 = v33
	v41 = v34
	v42 = v35
	goto L7
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == v51 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v71 = v50 - v51
	goto L5
L19:
	;
	v53 = int32(1)
	v58 = v47 - v53
	if v58 != 0 {
		v45 = v45 + v53
		v46 = v46 + v53
		v47 = v58
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L6
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v73 == v74 {
		v85 = int32(0)
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v73 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(1)
	goto L27
L26:
	;
	v78 = int32(-1)
	goto L27
L27:
	;
	return v78
L28:
	;
	v83 = int32(1)
	goto L30
L29:
	;
	v83 = int32(-1)
	goto L30
L30:
	;
	v85 = v83
	goto L1
}
func F_compare_distances(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F64_lt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_compare_values(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_FunctionCall2Coll(m, v4, v5, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int64(0) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_FunctionCall2Coll(m, v14, v15, v16, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v23 = base.B2i32(v18 != int64(0))
				return v23
			}
		} else {
			v23 = int32(-1)
			return v23
		}
	}
}
func F_comparecost_3(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v8 float32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.F32_gt(v7, v8) != 0 {
		v10 = int32(1)
	} else {
		v10 = int32(-1)
	}
	if base.F32_ne(v7, v8) != 0 {
		v13 = v10
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_compatible_oper_opid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_oper(m, v4, l0, l1, l2, v4, int32(-1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			v58 = v4
			m.G0 = v10 + int32(16)
			if v58 != 0 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
				F_ReleaseCatCache(m, v58)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v69 = v65
					return v69
				}
			} else {
				v69 = v4
				return v69
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
			v23 = v21 + v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
			v25 = F_IsBinaryCoercible(m, l1, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					F_ReleaseCatCache(m, v15)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = F_op_signature_string(m, l0, l1, l2)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
									F_errmsg(m, int32(_a_F_compatible_oper_opid_0), v10)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, int32(0), int32(-1))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_compatible_oper_opid_1), int32(484), int32(_a_F_compatible_oper_opid_2))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					v30 = F_IsBinaryCoercible(m, l2, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							F_ReleaseCatCache(m, v15)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = F_op_signature_string(m, l0, l1, l2)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
											F_errmsg(m, int32(_a_F_compatible_oper_opid_0), v10)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, int32(0), int32(-1))
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_compatible_oper_opid_1), int32(484), int32(_a_F_compatible_oper_opid_2))
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
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
						} else {
							v58 = v15
							m.G0 = v10 + int32(16)
							if v58 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
								v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
								F_ReleaseCatCache(m, v58)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v69 = v65
									return v69
								}
							} else {
								v69 = v4
								return v69
							}
						}
					}
				}
			}
		}
	}
}
func F_composite_to_json(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
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
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v18 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v22 = F_lookup_rowtype_tupdesc(m, v20, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(base.Ui32(v24) >> (uint(int32(2)) % 32))
	F_appendStringInfoChar(m, l1, int32(123))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if int32(0) < v32 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_appendStringInfoChar(m, l1, int32(125))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L60
	}
L8:
	;
	v37 = int32(3)
	goto L10
L9:
	;
	v37 = int32(1)
	goto L10
L10:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = int32(_a_F_composite_to_json_0)
	goto L13
L12:
	;
	v40 = int32(_a_F_composite_to_json_1)
	goto L13
L13:
	;
	v46 = int32(0)
	v48 = v32
	v51 = int32(0)
	goto L14
L14:
	;
	v61 = v22 + v48<<(uint(int32(3))%32) + v46*int32(100)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+119)))
	if v62 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L7
L16:
	;
	v189 = v48
	v191 = v51
	v193 = v46 + int32(1)
	goto L18
L17:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v193 < v189 {
		v46 = v193
		v48 = v189
		v51 = v191
		goto L14
	} else {
		goto L59
	}
L19:
	;
	F_appendBinaryStringInfo(m, l1, v40, v37)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_escape_json(m, l1, v61+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	F_appendStringInfoChar(m, l1, int32(58))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v79 = v46 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+18)))
	if base.Ui32(v81&int32(2047)) <= base.Ui32(v46) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	if v157 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L26:
	;
	v87 = F_getmissingattr(m, v22, v79, v15+int32(27))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v89)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+20)))
	if v91&int32(1) == v89 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v156 = v87
	goto L25
L30:
	;
	v98 = v22 + int32(20) + v79<<(uint(int32(3))%32)
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98))))
	if int32(0) <= v99 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(base.Ui32(v46)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v138)>>(uint(v46&int32(7))%32))&int32(1) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v104 = v80 + v102 + v99
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v105 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v133 = F_nocachegetattr(m, v15+int32(28), v79, v22)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L48
	}
L36:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if base.I32_popcnt(v108) != int32(1) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v156 = base.I64_extend_i32_u(v104)
	goto L25
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L45
	}
L40:
	;
	switch base.I32_ctz(v108) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	default:
		goto L39
	}
L41:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	v156 = v116
	goto L25
L42:
	;
	v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v104))))
	v156 = v115
	goto L25
L43:
	;
	v114 = int64(*(*int16)(unsafe.Add(mBase, uint32(v104))))
	v156 = v114
	goto L25
L44:
	;
	v113 = int64(*(*int8)(unsafe.Add(mBase, uint32(v104))))
	v156 = v113
	goto L25
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v108
	F_errmsg_internal(m, int32(_a_F_composite_to_json_2), v15)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_composite_to_json_3), int32(123), int32(_a_F_composite_to_json_4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
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
	v156 = v133
	goto L25
L49:
	;
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v146)
	v156 = int64(0)
	goto L25
L50:
	;
	goto L51
L51:
	;
	v151 = F_nocachegetattr(m, v15+int32(28), v79, v22)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v156 = v151
	goto L25
L53:
	;
	v181 = int32(1)
	F_datum_to_json_internal(m, v156, v179&v181, l1, v178, v180, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L58
	}
L54:
	;
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v160
	v178 = v160
	v179 = int32(1)
	v180 = v160
	goto L53
L55:
	;
	goto L56
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(28))+68))
	F_json_categorize_type(m, v167, int32(0), v15+int32(20), v15+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v178 = v176
	v179 = v175
	v180 = v177
	goto L53
L58:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v189 = v187
	v191 = v181
	v193 = v79
	goto L18
L59:
	;
	goto L15
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) <= v210 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_DecrTupleDescRefCount(m, v22)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	m.G0 = v15 + int32(48)
	return
L64:
	;
	goto L63
}
func F_compress_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.Env.Pgmem_deflate_write(m, v7, l2, l3)
	mBase = m.M
	if v8 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-105)
L2:
	;
	goto L3
L3:
	;
	v14 = l1 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = m.Env.Pgmem_zstream_read(m, v22, v14, v15)
	mBase = m.M
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v32
L6:
	;
	return int32(-105)
L7:
	;
	goto L8
L8:
	;
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = F_pushf_write(m, l0, v14, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if int32(0) <= v32 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_computeDistance(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v34 float64
	_ = v34
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v96 float64
	_ = v96
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v131 float64
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 float64
	_ = v176
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v179 float64
	_ = v179
	var v181 float64
	_ = v181
	var v183 float64
	_ = v183
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l0 != 0 {
		v24 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), base.I64_extend_i32_u(l2), base.I64_extend_i32_u(l1+int32(16)))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return float64(0)
		} else {
			v186 = base.F64_reinterpret_i64(v24)
			m.G0 = v16 + int32(16)
			return v186
		}
	} else {
		v29 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		v30 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		if base.F64_le(v29, v30) == int32(0) {
			v91 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v92 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_le(v91, v92) == int32(0) {
				v148 = base.I64_extend_i32_u(l2)
				v152 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1+int32(16)))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return float64(0)
				} else {
					v157 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return float64(0)
					} else {
						v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v16))) = v159
						v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v161
						v165 = base.I64_extend_i32_u(v16)
						v166 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return float64(0)
						} else {
							v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							*(*float64)(unsafe.Add(mBase, uint32(v16))) = v168
							v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v170
							v174 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return float64(0)
							} else {
								v176 = base.F64_reinterpret_i64(v174)
								v177 = base.F64_reinterpret_i64(v166)
								v178 = base.F64_reinterpret_i64(v157)
								v179 = base.F64_reinterpret_i64(v152)
								if base.F64_lt(v178, v179) != 0 {
									v181 = v178
								} else {
									v181 = v179
								}
								if base.F64_lt(v177, v181) != 0 {
									v183 = v177
								} else {
									v183 = v181
								}
								if base.F64_lt(v176, v183) != 0 {
									v185 = v176
								} else {
									v185 = v183
								}
								v186 = v185
								m.G0 = v16 + int32(16)
								return v186
							}
						}
					}
				}
			} else {
				v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				if base.F64_ge(v91, v96) == int32(0) {
					v148 = base.I64_extend_i32_u(l2)
					v152 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1+int32(16)))
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return float64(0)
					} else {
						v157 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return float64(0)
						} else {
							v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v16))) = v159
							v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v161
							v165 = base.I64_extend_i32_u(v16)
							v166 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return float64(0)
							} else {
								v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v16))) = v168
								v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v170
								v174 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return float64(0)
								} else {
									v176 = base.F64_reinterpret_i64(v174)
									v177 = base.F64_reinterpret_i64(v166)
									v178 = base.F64_reinterpret_i64(v157)
									v179 = base.F64_reinterpret_i64(v152)
									if base.F64_lt(v178, v179) != 0 {
										v181 = v178
									} else {
										v181 = v179
									}
									if base.F64_lt(v177, v181) != 0 {
										v183 = v177
									} else {
										v183 = v181
									}
									if base.F64_lt(v176, v183) != 0 {
										v185 = v176
									} else {
										v185 = v183
									}
									v186 = v185
									m.G0 = v16 + int32(16)
									return v186
								}
							}
						}
					}
				} else {
					if base.F64_gt(v29, v30) != 0 {
						v102 = math.Float64frombits(uint64(0x7ff0000000000000))
						v104 = base.F64_sub(v29, v30)
						if base.F64_eq(base.F64_abs(v29), v102)|base.F64_ne(base.F64_abs(v104), v102)|base.F64_eq(base.F64_abs(v30), v102) != 0 {
							v186 = v104
							m.G0 = v16 + int32(16)
							return v186
						} else {
							v114 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return float64(0)
							} else {
								v186 = v114
								m.G0 = v16 + int32(16)
								return v186
							}
						}
					} else {
						v116 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						if base.F64_gt(v116, v29) != 0 {
							v119 = math.Float64frombits(uint64(0x7ff0000000000000))
							v121 = base.F64_sub(v116, v29)
							if base.F64_eq(base.F64_abs(v29), v119)|base.F64_ne(base.F64_abs(v121), v119)|base.F64_eq(base.F64_abs(v116), v119) != 0 {
								v186 = v121
								m.G0 = v16 + int32(16)
								return v186
							} else {
								v131 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return float64(0)
								} else {
									v186 = v131
									m.G0 = v16 + int32(16)
									return v186
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_computeDistance_1), int32(1259), int32(_a_F_computeDistance_2))
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
						}
					}
				}
			}
		} else {
			v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
			if base.F64_ge(v29, v34) == int32(0) {
				v91 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v92 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v91, v92) == int32(0) {
					v148 = base.I64_extend_i32_u(l2)
					v152 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1+int32(16)))
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return float64(0)
					} else {
						v157 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return float64(0)
						} else {
							v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v16))) = v159
							v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v161
							v165 = base.I64_extend_i32_u(v16)
							v166 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return float64(0)
							} else {
								v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v16))) = v168
								v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v170
								v174 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return float64(0)
								} else {
									v176 = base.F64_reinterpret_i64(v174)
									v177 = base.F64_reinterpret_i64(v166)
									v178 = base.F64_reinterpret_i64(v157)
									v179 = base.F64_reinterpret_i64(v152)
									if base.F64_lt(v178, v179) != 0 {
										v181 = v178
									} else {
										v181 = v179
									}
									if base.F64_lt(v177, v181) != 0 {
										v183 = v177
									} else {
										v183 = v181
									}
									if base.F64_lt(v176, v183) != 0 {
										v185 = v176
									} else {
										v185 = v183
									}
									v186 = v185
									m.G0 = v16 + int32(16)
									return v186
								}
							}
						}
					}
				} else {
					v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v91, v96) == int32(0) {
						v148 = base.I64_extend_i32_u(l2)
						v152 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1+int32(16)))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return float64(0)
						} else {
							v157 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, base.I64_extend_i32_u(l1))
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return float64(0)
							} else {
								v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v16))) = v159
								v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v161
								v165 = base.I64_extend_i32_u(v16)
								v166 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return float64(0)
								} else {
									v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
									*(*float64)(unsafe.Add(mBase, uint32(v16))) = v168
									v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v170
									v174 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v148, v165)
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return float64(0)
									} else {
										v176 = base.F64_reinterpret_i64(v174)
										v177 = base.F64_reinterpret_i64(v166)
										v178 = base.F64_reinterpret_i64(v157)
										v179 = base.F64_reinterpret_i64(v152)
										if base.F64_lt(v178, v179) != 0 {
											v181 = v178
										} else {
											v181 = v179
										}
										if base.F64_lt(v177, v181) != 0 {
											v183 = v177
										} else {
											v183 = v181
										}
										if base.F64_lt(v176, v183) != 0 {
											v185 = v176
										} else {
											v185 = v183
										}
										v186 = v185
										m.G0 = v16 + int32(16)
										return v186
									}
								}
							}
						}
					} else {
						if base.F64_gt(v29, v30) != 0 {
							v102 = math.Float64frombits(uint64(0x7ff0000000000000))
							v104 = base.F64_sub(v29, v30)
							if base.F64_eq(base.F64_abs(v29), v102)|base.F64_ne(base.F64_abs(v104), v102)|base.F64_eq(base.F64_abs(v30), v102) != 0 {
								v186 = v104
								m.G0 = v16 + int32(16)
								return v186
							} else {
								v114 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return float64(0)
								} else {
									v186 = v114
									m.G0 = v16 + int32(16)
									return v186
								}
							}
						} else {
							v116 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							if base.F64_gt(v116, v29) != 0 {
								v119 = math.Float64frombits(uint64(0x7ff0000000000000))
								v121 = base.F64_sub(v116, v29)
								if base.F64_eq(base.F64_abs(v29), v119)|base.F64_ne(base.F64_abs(v121), v119)|base.F64_eq(base.F64_abs(v116), v119) != 0 {
									v186 = v121
									m.G0 = v16 + int32(16)
									return v186
								} else {
									v131 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return float64(0)
									} else {
										v186 = v131
										m.G0 = v16 + int32(16)
										return v186
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(_a_F_computeDistance_1), int32(1259), int32(_a_F_computeDistance_2))
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
							}
						}
					}
				}
			} else {
				v38 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v38, v39) != 0 {
					v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v38, v41) != 0 {
						v186 = float64(0)
						m.G0 = v16 + int32(16)
						return v186
					} else {
						v43 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
						v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						if base.F64_gt(v43, v44) != 0 {
							v47 = math.Float64frombits(uint64(0x7ff0000000000000))
							v49 = base.F64_sub(v43, v44)
							if base.F64_eq(base.F64_abs(v43), v47)|base.F64_ne(base.F64_abs(v49), v47)|base.F64_eq(base.F64_abs(v44), v47) != 0 {
								v186 = v49
								m.G0 = v16 + int32(16)
								return v186
							} else {
								v59 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return float64(0)
								} else {
									v186 = v59
									m.G0 = v16 + int32(16)
									return v186
								}
							}
						} else {
							v61 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							if base.F64_gt(v61, v43) != 0 {
								v64 = math.Float64frombits(uint64(0x7ff0000000000000))
								v66 = base.F64_sub(v61, v43)
								if base.F64_eq(base.F64_abs(v43), v64)|base.F64_ne(base.F64_abs(v66), v64)|base.F64_eq(base.F64_abs(v61), v64) != 0 {
									v186 = v66
									m.G0 = v16 + int32(16)
									return v186
								} else {
									v76 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return float64(0)
									} else {
										v186 = v76
										m.G0 = v16 + int32(16)
										return v186
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(_a_F_computeDistance_1), int32(1248), int32(_a_F_computeDistance_2))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return float64(0)
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
				} else {
					v43 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
					v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					if base.F64_gt(v43, v44) != 0 {
						v47 = math.Float64frombits(uint64(0x7ff0000000000000))
						v49 = base.F64_sub(v43, v44)
						if base.F64_eq(base.F64_abs(v43), v47)|base.F64_ne(base.F64_abs(v49), v47)|base.F64_eq(base.F64_abs(v44), v47) != 0 {
							v186 = v49
							m.G0 = v16 + int32(16)
							return v186
						} else {
							v59 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return float64(0)
							} else {
								v186 = v59
								m.G0 = v16 + int32(16)
								return v186
							}
						}
					} else {
						v61 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
						if base.F64_gt(v61, v43) != 0 {
							v64 = math.Float64frombits(uint64(0x7ff0000000000000))
							v66 = base.F64_sub(v61, v43)
							if base.F64_eq(base.F64_abs(v43), v64)|base.F64_ne(base.F64_abs(v66), v64)|base.F64_eq(base.F64_abs(v61), v64) != 0 {
								v186 = v66
								m.G0 = v16 + int32(16)
								return v186
							} else {
								v76 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return float64(0)
								} else {
									v186 = v76
									m.G0 = v16 + int32(16)
									return v186
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_computeDistance_1), int32(1248), int32(_a_F_computeDistance_2))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return float64(0)
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
	}
}
func F_computeLeafRecompressWALData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v14 == v2)|base.B2i32(l0 == v14) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = v2
	v23 = v14
	goto L4
L2:
	;
	v37 = v2
	goto L3
L3:
	;
	v49 = F_palloc(m, v37<<(uint(int32(1))%32)+int32(_a_F_computeLeafRecompressWALData_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
	v33 = v22 + base.B2i32(v30 != int32(0))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v34 != l0 {
		v22 = v33
		v23 = v34
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v37 = v33
	goto L3
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v37)
	v53 = v49 + int32(2)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = int32(0)
	if base.B2i32(v54 == v55)|base.B2i32(l0 == v54) == v55 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = v53
	v63 = v54
	v66 = v2
	goto L12
L10:
	;
	v155 = v53
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v155 - v49
	m.G0 = v12 + int32(16)
	return
L12:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
	switch v70 {
	case 0:
		goto L18
	case 1:
		goto L16
	default:
		goto L17
	}
L13:
	;
	v155 = v147
	goto L11
L14:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v152 != l0 {
		v62 = v147
		v63 = v152
		v66 = v151
		goto L12
	} else {
		goto L37
	}
L15:
	;
	v147 = v142 + v62 + int32(2)
	v151 = v139 + v66
	goto L14
L16:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v134)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	v139 = v134
	v142 = int32(0)
	goto L15
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)))
	v78 = (v74 + int32(1)) & int32(_a_F_computeLeafRecompressWALData_1)
	v80 = v78 + int32(8)
	if v70 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v147 = v62
	v151 = v66 + int32(1)
	goto L14
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L34
	}
L20:
	;
	if v80 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	if base.Ui32(v83*int32(6)) <= base.Ui32(v80) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	if v70&int32(254) != int32(2) {
		goto L19
	} else {
		goto L30
	}
L24:
	;
	v88 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v88)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)) = uint16(v90)
	v93 = v90 * int32(6)
	if v93 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v101 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v101)
	v110 = v101
	goto L20
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	base.MemoryCopy(m, v62+int32(4), v96, v93)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v139 = int32(1)
	v142 = v93 + int32(2)
	goto L15
L30:
	;
	v110 = v70
	goto L20
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	base.MemoryCopy(m, v62+int32(2), v113, v80)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v139 = base.B2i32(v110 != int32(2))
	v142 = (v78 + int32(9)) & int32(_a_F_computeLeafRecompressWALData_2)
	goto L15
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v70
	F_errmsg_internal(m, int32(_a_F_computeLeafRecompressWALData_3), v12)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_computeLeafRecompressWALData_4), int32(955), int32(_a_F_computeLeafRecompressWALData_5))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	goto L13
}
func F_connect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_connect(m, l0, int32(_a_F_connect_0), int32(12), v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_connect[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_consider_index_join_outer_rels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v328 int32
	_ = v328
	var v350 int32
	_ = v350
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	v11 = int32(0)
	if l7 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v22 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = v11
	goto L4
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v39<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v55 = F_list_member(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v55 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v59 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v371 = v39 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v371 < v372 {
		v39 = v371
		goto L4
	} else {
		goto L85
	}
L11:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v54, l9)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L84
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v62 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v77 = int32(0)
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v77<<(uint(int32(2))%32))))
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L11
L16:
	;
	v328 = v77 + int32(1)
	if v328 != v62 {
		v77 = v328
		goto L14
	} else {
		goto L83
	}
L17:
	;
	if v184 != int32(3) {
		goto L16
	} else {
		goto L52
	}
L18:
	;
	v184 = base.B2i32(v90 != int32(0))
	goto L17
L19:
	;
	goto L20
L20:
	;
	if v90 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v184 = int32(2)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v107 < v108 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v110 = v107
	goto L26
L25:
	;
	v110 = v108
	goto L26
L26:
	;
	if v110 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v113 = int32(1)
	goto L29
L28:
	;
	v113 = v110
	goto L29
L29:
	;
	v114 = int32(8)
	v118 = int32(0)
	v120 = v118
	v121 = v118
	goto L32
L30:
	;
	v184 = int32(3)
	goto L17
L31:
	;
	v184 = v171
	goto L17
L32:
	;
	v131 = v121 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v54+v114+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v90+v114))))
	if v133&(v135^int32(-1)) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v108 < v107 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v157 = v121 + int32(1)
	if v157 != v113 {
		v120 = v155
		v121 = v157
		goto L32
	} else {
		goto L41
	}
L35:
	;
	if base.B2i32(v120 == int32(1))|v135&(v133^int32(-1)) != 0 {
		v171 = int32(3)
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v135&(v133^int32(-1)) == int32(0) {
		v155 = v120
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v155 = int32(2)
	goto L34
L39:
	;
	if v120 == int32(2) {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v155 = int32(1)
	goto L34
L41:
	;
	goto L33
L42:
	;
	if v155 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if v108 <= v107 {
		v171 = v155
		goto L31
	} else {
		goto L48
	}
L45:
	;
	v164 = int32(3)
	goto L47
L46:
	;
	v164 = int32(2)
	goto L47
L47:
	;
	v184 = v164
	goto L17
L48:
	;
	if v155 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v170 = int32(3)
	goto L51
L50:
	;
	v170 = int32(1)
	goto L51
L51:
	;
	v171 = v170
	goto L31
L52:
	;
	if v52 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v299 != 0 {
		goto L77
	} else {
		goto L78
	}
L54:
	;
	v189 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v190 <= v189 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v201 = v189
	v208 = v190
	goto L56
L56:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212+v201<<(uint(int32(2))%32))))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+60))
	if v52 == v218 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L53
L58:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	v221 = int32(0)
	if v220 == v221 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v276 = v208
	goto L60
L60:
	;
	v278 = v201 + int32(1)
	if v278 < v276 {
		v201 = v278
		v208 = v276
		goto L56
	} else {
		goto L76
	}
L61:
	;
	if v274 != 0 {
		goto L16
	} else {
		goto L75
	}
L62:
	;
	v274 = int32(1)
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v90 == int32(0) {
		v267 = v221
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v274 = v267
	goto L61
L66:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v231 < v230 {
		v267 = v221
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v233 = int32(1)
	if v230 <= v233 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v236 = v233
	goto L70
L69:
	;
	v236 = v230
	goto L70
L70:
	;
	v237 = int32(8)
	v242 = int32(0)
	goto L71
L71:
	;
	v249 = v242 << (uint(int32(2)) % 32)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v220+v237+v249)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v90+v237+v249)))
	v256 = v251 & (v253 ^ int32(-1))
	v258 = base.B2i32(v256 == int32(0))
	if v256 != 0 {
		v267 = v258
		goto L65
	} else {
		goto L73
	}
L72:
	;
	v267 = v258
	goto L65
L73:
	;
	v260 = v242 + int32(1)
	if v260 != v236 {
		v242 = v260
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v276 = v275
	goto L60
L76:
	;
	goto L57
L77:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v302 = v300
	goto L79
L78:
	;
	v302 = int32(0)
	goto L79
L79:
	;
	if l8*int32(10) <= v302 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v304 = F_bms_union(m, v54, v90)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v304, l9)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	goto L16
L83:
	;
	goto L15
L84:
	;
	goto L10
L85:
	;
	goto L5
}
func F_conv_utf8_to(m *base.Module, l0 int32) int32 {
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		v45 = l0
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_conv_utf8_to_0)) {
			v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(1984) | l0&int32(63)
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(16777215)) {
				v45 = int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(_a_F_conv_utf8_to_1) | (int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | l0&int32(63))
			} else {
				v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | (int32(base.Ui32(l0)>>(uint(int32(6))%32))&int32(_a_F_conv_utf8_to_2) | (int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(_a_F_conv_utf8_to_3) | l0&int32(63)))
			}
		}
	}
	if base.Ui32(v45-int32(1106)) <= base.Ui32(int32(_a_F_conv_utf8_to_4)) {
		v51 = v45 - int32(286)
		v52 = int32(_a_F_conv_utf8_to_0)
		v53 = v51 & v52
		v55 = base.I32_div_u_s(v53, int32(1260))
		v58 = int32(10)
		v59 = base.I32_div_u_s(v53, v58)
		v61 = base.I32_rem_u_s(v59, int32(126))
		return v55<<(uint(int32(16))%32) | (v61<<(uint(int32(8))%32) + (v51-v59*v58)&v52 + int32(_a_F_conv_utf8_to_5)) | int32(-2127560656)
	} else {
		if base.Ui32(v45-int32(_a_F_conv_utf8_to_6)) <= base.Ui32(int32(2109)) {
			v81 = v45 - int32(576)
			v82 = int32(_a_F_conv_utf8_to_0)
			v83 = v81 & v82
			v85 = base.I32_div_u_s(v83, int32(1260))
			v88 = int32(10)
			v89 = base.I32_div_u_s(v83, v88)
			v91 = base.I32_rem_u_s(v89, int32(126))
			return v85<<(uint(int32(16))%32) | (v91<<(uint(int32(8))%32) + (v81-v89*v88)&v82 + int32(_a_F_conv_utf8_to_5)) | int32(-2127560656)
		} else {
			if base.Ui32(v45-int32(_a_F_conv_utf8_to_7)) <= base.Ui32(int32(764)) {
				v111 = v45 - int32(878)
				v112 = int32(_a_F_conv_utf8_to_0)
				v113 = v111 & v112
				v115 = base.I32_div_u_s(v113, int32(1260))
				v120 = int32(10)
				v121 = base.I32_div_u_s(v113, v120)
				v123 = base.I32_rem_u_s(v121, int32(126))
				return v115<<(uint(int32(16))%32) + int32(2146828288) | (v123<<(uint(int32(8))%32) + (v111-v121*v120)&v112 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
			} else {
				if base.Ui32(v45-int32(_a_F_conv_utf8_to_8)) <= base.Ui32(int32(884)) {
					v143 = v45 - int32(887)
					v144 = int32(_a_F_conv_utf8_to_0)
					v145 = v143 & v144
					v147 = base.I32_div_u_s(v145, int32(1260))
					v152 = int32(10)
					v153 = base.I32_div_u_s(v145, v152)
					v155 = base.I32_rem_u_s(v153, int32(126))
					return v147<<(uint(int32(16))%32) + int32(2146828288) | (v155<<(uint(int32(8))%32) + (v143-v153*v152)&v144 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
				} else {
					if base.Ui32(v45-int32(_a_F_conv_utf8_to_9)) <= base.Ui32(int32(470)) {
						v175 = v45 - int32(889)
						v176 = int32(_a_F_conv_utf8_to_0)
						v177 = v175 & v176
						v179 = base.I32_div_u_s(v177, int32(1260))
						v184 = int32(10)
						v185 = base.I32_div_u_s(v177, v184)
						v187 = base.I32_rem_u_s(v185, int32(126))
						return v179<<(uint(int32(16))%32) + int32(2146828288) | (v187<<(uint(int32(8))%32) + (v175-v185*v184)&v176 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
					} else {
						if base.Ui32(v45-int32(_a_F_conv_utf8_to_10)) <= base.Ui32(int32(372)) {
							v207 = v45 - int32(894)
							v208 = int32(_a_F_conv_utf8_to_0)
							v209 = v207 & v208
							v211 = base.I32_div_u_s(v209, int32(1260))
							v216 = int32(10)
							v217 = base.I32_div_u_s(v209, v216)
							v219 = base.I32_rem_u_s(v217, int32(126))
							return v211<<(uint(int32(16))%32) + int32(2146828288) | (v219<<(uint(int32(8))%32) + (v207-v217*v216)&v208 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
						} else {
							if base.Ui32(v45-int32(_a_F_conv_utf8_to_11)) <= base.Ui32(int32(440)) {
								v239 = v45 - int32(900)
								v240 = int32(_a_F_conv_utf8_to_0)
								v241 = v239 & v240
								v243 = base.I32_div_u_s(v241, int32(1260))
								v248 = int32(10)
								v249 = base.I32_div_u_s(v241, v248)
								v251 = base.I32_rem_u_s(v249, int32(126))
								return v243<<(uint(int32(16))%32) + int32(2146828288) | (v251<<(uint(int32(8))%32) + (v239-v249*v248)&v240 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
							} else {
								if base.Ui32(v45-int32(_a_F_conv_utf8_to_12)) <= base.Ui32(int32(702)) {
									v271 = v45 - int32(911)
									v272 = int32(_a_F_conv_utf8_to_0)
									v273 = v271 & v272
									v275 = base.I32_div_u_s(v273, int32(1260))
									v280 = int32(10)
									v281 = base.I32_div_u_s(v273, v280)
									v283 = base.I32_rem_u_s(v281, int32(126))
									return v275<<(uint(int32(16))%32) + int32(2146828288) | (v283<<(uint(int32(8))%32) + (v271-v281*v280)&v272 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
								} else {
									if base.Ui32(v45-int32(_a_F_conv_utf8_to_13)) <= base.Ui32(int32(_a_F_conv_utf8_to_14)) {
										v303 = v45 - int32(_a_F_conv_utf8_to_15)
										v304 = int32(_a_F_conv_utf8_to_0)
										v305 = v303 & v304
										v307 = base.I32_div_u_s(v305, int32(_a_F_conv_utf8_to_16))
										v311 = base.I32_div_u_s(v305, int32(1260))
										v312 = int32(10)
										v313 = base.I32_rem_u_s(v311, v312)
										v320 = base.I32_div_u_s(v305, v312)
										v322 = base.I32_rem_u_s(v320, int32(126))
										return v307<<(uint(int32(24))%32) | v313<<(uint(int32(16))%32) - int32(2130706432) | (v322<<(uint(int32(8))%32) + (v303-v320*v312)&v304 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
									} else {
										if base.Ui32(v45-int32(_a_F_conv_utf8_to_18)) <= base.Ui32(int32(_a_F_conv_utf8_to_19)) {
											v342 = v45 - int32(_a_F_conv_utf8_to_20)
											v343 = int32(_a_F_conv_utf8_to_0)
											v344 = v342 & v343
											v346 = base.I32_div_u_s(v344, int32(_a_F_conv_utf8_to_16))
											v350 = base.I32_div_u_s(v344, int32(1260))
											v351 = int32(10)
											v352 = base.I32_rem_u_s(v350, v351)
											v359 = base.I32_div_u_s(v344, v351)
											v361 = base.I32_rem_u_s(v359, int32(126))
											return v346<<(uint(int32(24))%32) | v352<<(uint(int32(16))%32) - int32(2130706432) | (v361<<(uint(int32(8))%32) + (v342-v359*v351)&v343 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
										} else {
											if base.Ui32(v45-int32(_a_F_conv_utf8_to_21)) <= base.Ui32(int32(1029)) {
												v381 = v45 - int32(_a_F_conv_utf8_to_22)
												v382 = int32(_a_F_conv_utf8_to_0)
												v383 = v381 & v382
												v385 = base.I32_div_u_s(v383, int32(_a_F_conv_utf8_to_16))
												v389 = base.I32_div_u_s(v383, int32(1260))
												v390 = int32(10)
												v391 = base.I32_rem_u_s(v389, v390)
												v398 = base.I32_div_u_s(v383, v390)
												v400 = base.I32_rem_u_s(v398, int32(126))
												return v385<<(uint(int32(24))%32) | v391<<(uint(int32(16))%32) - int32(2130706432) | (v400<<(uint(int32(8))%32) + (v381-v398*v390)&v382 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
											} else {
												if base.Ui32(v45-int32(_a_F_conv_utf8_to_23)) <= base.Ui32(int32(25)) {
													v420 = v45 - int32(_a_F_conv_utf8_to_24)
													v421 = int32(_a_F_conv_utf8_to_0)
													v422 = v420 & v421
													v424 = base.I32_div_u_s(v422, int32(_a_F_conv_utf8_to_16))
													v428 = base.I32_div_u_s(v422, int32(1260))
													v429 = int32(10)
													v430 = base.I32_rem_u_s(v428, v429)
													v437 = base.I32_div_u_s(v422, v429)
													v439 = base.I32_rem_u_s(v437, int32(126))
													return v424<<(uint(int32(24))%32) | v430<<(uint(int32(16))%32) - int32(2130706432) | (v439<<(uint(int32(8))%32) + (v420-v437*v429)&v421 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
												} else {
													if base.Ui32(v45-int32(_a_F_conv_utf8_to_25)) <= base.Ui32(int32(_a_F_conv_utf8_to_26)) {
														v459 = v45 + int32(_a_F_conv_utf8_to_27)
														v460 = int32(10)
														v461 = base.I32_div_u_s(v459, v460)
														v463 = base.I32_rem_u_s(v461, int32(126))
														v473 = base.I32_div_u_s(v459, int32(_a_F_conv_utf8_to_16))
														v477 = base.I32_div_u_s(v459, int32(1260))
														v481 = base.I32_rem_u_s(v477&int32(_a_F_conv_utf8_to_0), v460)
														v493 = v463<<(uint(int32(8))%32) + (v459 - v461*v460) + int32(_a_F_conv_utf8_to_5) | (v473<<(uint(int32(24))%32) | v481<<(uint(int32(16))%32) - int32(2130706432)) | int32(_a_F_conv_utf8_to_17)
													} else {
														v493 = int32(0)
													}
													return v493
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
func F_convert_EXISTS_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v319 int32
	_ = v319
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(400)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v16 != 0 {
		v319 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(400)
	return v319
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = F_copyObjectImpl(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v22 = F_simplify_EXISTS_query(m, l0, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v22 == int32(0) {
		v319 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(0)
	v31 = F_contain_vars_of_level(m, v18, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v31 != 0 {
		v319 = v5
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v34 = F_contain_vars_of_level(m, v27, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if v34 == int32(0) {
		v319 = v5
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v38 = F_contain_volatile_functions(m, v27)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if v38 != 0 {
		v319 = v5
		goto L1
	} else {
		goto L12
	}
L12:
	;
	base.MemoryFill(m, v13+int32(8), int32(0), int32(392))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(269)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v18
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v27
	v52 = F_preprocess_relation_rtes(m, v13)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+60))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = int32(0)
	F_replace_empty_jointree(m, v52)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v61 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v63 = v62
	goto L17
L16:
	;
	v63 = int32(0)
	goto L17
L17:
	;
	F_OffsetVarNodes(m, v52, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_OffsetVarNodes(m, v55, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_IncrementVarSublevelsUp(m, v52, int32(-1), int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	F_IncrementVarSublevelsUp(m, v55, int32(-1), int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v76 = F_pull_varnos(m, l0, v55)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v76 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if int32(0) <= v134 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v134 = base.I32_ctz(v120) | v121<<(uint(int32(5))%32)
	goto L23
L25:
	;
	v134 = int32(-2)
	goto L23
L26:
	;
	v85 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v88 <= v85 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v91 = v76 + int32(8)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v98 = v95 & int32(-1)
	if v98 != 0 {
		v120 = v98
		v121 = v85
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v99 = int32(1)
	if v99 == v88 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v103 = v99
	goto L30
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91+v103<<(uint(int32(2))%32))))
	if v110 != 0 {
		v120 = v110
		v121 = v103
		goto L24
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	v112 = v103 + int32(1)
	if v112 != v88 {
		v103 = v112
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v141 = v134
	v145 = v5
	goto L37
L35:
	;
	v217 = v5
	goto L36
L36:
	;
	F_bms_free(m, v76)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L55
	}
L37:
	;
	if v141 <= v63 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v217 = v150
	goto L36
L39:
	;
	v148 = F_bms_add_member(m, v145, v141)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	v150 = v145
	goto L41
L41:
	;
	if v76 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v150 = v148
	goto L41
L43:
	;
	if int32(0) <= v206 {
		v141 = v206
		v145 = v150
		goto L37
	} else {
		goto L54
	}
L44:
	;
	v206 = base.I32_ctz(v192) | v193<<(uint(int32(5))%32)
	goto L43
L45:
	;
	v206 = int32(-2)
	goto L43
L46:
	;
	v157 = v141 + int32(1)
	v159 = int32(base.Ui32(v157) >> (uint(int32(5)) % 32))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v160 <= v159 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v163 = v76 + int32(8)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v159<<(uint(int32(2))%32))))
	v170 = v167 & (int32(-1) << (uint(v157) % 32))
	if v170 != 0 {
		v192 = v170
		v193 = v159
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v172 = v159 + int32(1)
	if v172 == v160 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v175 = v172
	goto L50
L50:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v163+v175<<(uint(int32(2))%32))))
	if v182 != 0 {
		v192 = v182
		v193 = v175
		goto L44
	} else {
		goto L52
	}
L51:
	;
	goto L45
L52:
	;
	v184 = v175 + int32(1)
	if v184 != v160 {
		v175 = v184
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L38
L55:
	;
	v221 = int32(0)
	if v217 == v221 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v275 == int32(0) {
		v319 = v221
		goto L1
	} else {
		goto L70
	}
L57:
	;
	v275 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if l3 == int32(0) {
		v268 = v221
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v275 = v268
	goto L56
L61:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v232 < v231 {
		v268 = v221
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v234 = int32(1)
	if v231 <= v234 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v237 = v234
	goto L65
L64:
	;
	v237 = v231
	goto L65
L65:
	;
	v238 = int32(8)
	v243 = int32(0)
	goto L66
L66:
	;
	v250 = v243 << (uint(int32(2)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v217+v238+v250)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l3+v238+v250)))
	v257 = v252 & (v254 ^ int32(-1))
	v259 = base.B2i32(v257 == int32(0))
	if v257 != 0 {
		v268 = v259
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v268 = v259
	goto L60
L68:
	;
	v261 = v243 + int32(1)
	if v261 != v237 {
		v243 = v261
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
	F_CombineRangeTables(m, v17+int32(52), v17+int32(56), v282, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v287 = F_palloc0(m, int32(40))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v289 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v287)+12)) = v289
	*(*uint8)(unsafe.Add(mBase, uint32(v287)+8)) = uint8(v289)
	if l2 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v295 = int32(5)
	goto L75
L74:
	;
	v295 = int32(4)
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = int32(64)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v52)+60))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if v300 == int32(0) {
		v308 = v299
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v309 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v287)+32)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v287)+28)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v287)+20)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v287)+16)) = v308
	v319 = v287
	goto L1
L77:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v303 != int32(1) {
		v308 = v299
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v308 = v307
	goto L76
}
func F_convert_combining_aggrefs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v3 {
		v79 = v3
		m.G0 = v7 + int32(16)
		return v79
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v11 == int32(9) {
			v15 = F_palloc0(m, int32(72))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(9)
				base.MemoryCopy(m, v15, l0, int32(72))
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v23
				v27 = F_copyObjectImpl(m, v15)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = int32(6)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
					if v38 == int32(2281) {
						v41 = int32(17)
					} else {
						v41 = v38
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v41
					v48 = int32(0)
					v50 = F_makeTargetEntry(m, v15, int32(1), v48, v48)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v50
						*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v50
						v57 = F_list_make1_impl(m, int32(1), v7+int32(8))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = int32(9)
							v79 = v27
							m.G0 = v7 + int32(16)
							return v79
						}
					}
				}
			}
		} else {
			v75 = F_expression_tree_mutator_impl(m, l0, int32(885), l1)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v79 = v75
				m.G0 = v7 + int32(16)
				return v79
			}
		}
	}
}
func F_copyScalarSubstructure(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v11 {
		case 0, 3:
			m.G0 = v7 + int32(16)
			return
		case 1:
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = F_MemoryContextAlloc(m, l1, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v15 != 0 {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v13, v16, v15)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
				m.G0 = v7 + int32(16)
				return
			}
		case 2:
			v58 = int32(_a_F_copyScalarSubstructure_0)
			v59 = *(*int32)(unsafe.Add(mBase, _c_F_copyScalarSubstructure[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_copyScalarSubstructure[0])) = l1
			v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
			v65 = F_datumCopy(m, v62, int32(0), int32(-1))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				v68 = F_pg_detoast_datum(m, base.I32_wrap_i64(v65))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v68
					*(*int32)(unsafe.Add(mBase, _c_F_copyScalarSubstructure[0])) = v59
					m.G0 = v7 + int32(16)
					return
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_copyScalarSubstructure_1), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_copyScalarSubstructure_2), int32(924), int32(_a_F_copyScalarSubstructure_3))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 32:
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v19 <= int32(1183) {
				if base.B2i32(v19 == int32(1114))|base.B2i32(base.Ui32(v19-int32(1082)) < base.Ui32(int32(2))) != 0 {
					m.G0 = v7 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v82
						F_errmsg_internal(m, int32(_a_F_copyScalarSubstructure_4), v7)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_copyScalarSubstructure_2), int32(920), int32(_a_F_copyScalarSubstructure_3))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
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
				if v19 == int32(1184) {
					m.G0 = v7 + int32(16)
					return
				} else {
					if v19 != int32(1266) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v82
							F_errmsg_internal(m, int32(_a_F_copyScalarSubstructure_4), v7)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_copyScalarSubstructure_2), int32(920), int32(_a_F_copyScalarSubstructure_3))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v33 = int32(_a_F_copyScalarSubstructure_0)
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_copyScalarSubstructure[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_copyScalarSubstructure[0])) = l1
						v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						v40 = F_datumCopy(m, v37, int32(0), int32(12))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v40
							*(*int32)(unsafe.Add(mBase, _c_F_copyScalarSubstructure[0])) = v34
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_copy_dest_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	return
}
func F_copydir(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v9 = m.G0
	v11 = v9 - int32(_a_F_copydir_0)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_copydir[0]))
	v15 = F_mkdir(m, l1, v14)
	mBase = m.M
	goto L1
L1:
	;
	if v15 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v18 = F_AllocateDir(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L58
	}
L5:
	;
	return
L6:
	;
	v20 = F_ReadDir(m, v18, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = v20
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_FreeDir(m, v18)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L33
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_copydir[1]))
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)))
	if v34 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v91 = F_ReadDir(m, v18, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L31
	}
L18:
	;
	v47 = v26 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l0
	v51 = v11 + int32(2112)
	v56 = F_pg_snprintf(m, v51, int32(2048), int32(_a_F_copydir_1), v11+int32(32))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L23
	}
L19:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v37 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v40 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+21)))
	if v43 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v66 = F_pg_snprintf(m, v11-int32(-64), int32(2048), int32(_a_F_copydir_1), v11+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v70 = F_get_dirent_type(m, v51, v26, int32(0), int32(21))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	F_copy_file(m, v11+int32(2112), v11-int32(-64))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L30
	}
L26:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L28
	}
L27:
	;
	switch v70 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L26
	default:
		goto L17
	}
L28:
	;
	F_copydir(m, v11+int32(2112), v11-int32(-64), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	goto L17
L31:
	;
	if v91 != 0 {
		v26 = v91
		goto L11
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copydir[2])))
	if v104 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v107 = F_AllocateDir(m, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v11 + int32(_a_F_copydir_0)
	return
L37:
	;
	v109 = F_ReadDir(m, v107, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v109 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v115 = v109
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_FreeDir(m, v107)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L56
	}
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+19)))
	if v119 != int32(46) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v151 = F_ReadDir(m, v107, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L54
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v115 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	v136 = v11 - int32(-64)
	v139 = F_pg_snprintf(m, v136, int32(2048), int32(_a_F_copydir_1), v11)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L50
	}
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+20)))
	if v122 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+20)))
	if v125 != int32(46) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+21)))
	if v128 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v143 = F_get_dirent_type(m, v136, v115, int32(0), int32(21))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v143 != int32(2) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	F_fsync_fname(m, v136, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	if v151 != 0 {
		v115 = v151
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	F_fsync_fname(m, l1, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	goto L36
L58:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l1
	F_errmsg(m, int32(_a_F_copydir_2), v11+int32(48))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_copydir_3), int32(59), int32(_a_F_copydir_4))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_core_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_core_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_core_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_2(m, int32(_a_F_core_yy_create_buffer_2))
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
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_core_yy_create_buffer[0]))
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
					*(*int32)(unsafe.Add(mBase, _c_F_core_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_2(m, int32(_a_F_core_yy_create_buffer_2))
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
func F_cost_sort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32, l5 float64, l6 int32, l7 float64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v31 int64
	_ = v31
	var v32 float64
	_ = v32
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v73 float64
	_ = v73
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v84 float64
	_ = v84
	var v88 float64
	_ = v88
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v102 float64
	_ = v102
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v112 float64
	_ = v112
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = v12 + int32(8)
	v20 = float64(2)
	if base.F64_lt(l3, v20) != 0 {
		v23 = v20
	} else {
		v23 = l3
	}
	v25 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[0]))
	v28 = base.F64_mul(v23, base.F64_add(base.F64_add(v25, v25), l5))
	v31 = base.I64_extend_i32_s(l6) << (uint(int64(10)) % 64)
	v32 = base.F64_convert_i64_s(v31)
	v39 = base.F64_convert_i32_u((l4+int32(7))&int32(-8) + int32(24))
	v41 = base.F64_mul(l3, v39)
	v45 = base.F64_lt(l7, v23) & base.F64_gt(l7, float64(0))
	if v45 != 0 {
		v46 = base.F64_mul(l7, v39)
	} else {
		v46 = v41
	}
	if base.F64_lt(v32, v46) != 0 {
		v48 = F_log(m, v23)
		mBase = m.M
		v49 = F_tuplesort_merge_order(m, v31)
		mBase = m.M
		v52 = base.F64_mul(base.F64_div(v48, float64(0.693147180559945)), v28)
		*(*float64)(unsafe.Add(mBase, uint32(v15))) = v52
		v56 = base.F64_ceil(base.F64_mul(v41, float64(0.0001220703125)))
		v58 = base.F64_div(v41, v32)
		v59 = base.F64_convert_i32_s(v49)
		if base.F64_gt(v58, v59) != 0 {
			v61 = F_log(m, v58)
			mBase = m.M
			v62 = F_log(m, v59)
			mBase = m.M
			v66 = base.F64_ceil(base.F64_div(v61, v62))
		} else {
			v66 = float64(1)
		}
		v69 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[1]))
		v73 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[2]))
		v96 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v56, v56), v66), base.F64_add(base.F64_mul(v69, float64(0.75)), base.F64_mul(v73, float64(0.25)))), v52)
	} else {
		if v45 != 0 {
			v79 = l7
		} else {
			v79 = v23
		}
		v80 = base.F64_add(v79, v79)
		if base.F64_gt(v23, v80)|base.F64_gt(v41, v32) != 0 {
			v84 = F_log(m, v80)
			mBase = m.M
			v96 = base.F64_mul(base.F64_div(v84, float64(0.693147180559945)), v28)
		} else {
			v88 = F_log(m, v23)
			mBase = m.M
			v96 = base.F64_mul(base.F64_div(v88, float64(0.693147180559945)), v28)
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(v15))) = v96
	v99 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[0]))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = base.F64_mul(v23, v99)
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_sort[3])))
	v106 = base.F64_add(l2, v102)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1 + (v105 ^ int32(1))
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v106, v112)
	m.G0 = v12 + int32(16)
	return
}
