package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_PathNameCreateTemporaryFile(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[0]))
	F_ResourceOwnerEnlarge(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[1]))
		v18 = F_PathNameOpenFilePerm(m, l0, int32(578), v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 <= int32(0) {
				if l1 == int32(0) {
					m.G0 = v7 + int32(16)
					return v18
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(_a_F_PathNameCreateTemporaryFile_0), v7)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_PathNameCreateTemporaryFile_1), int32(1873), int32(_a_F_PathNameCreateTemporaryFile_2))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
			} else {
				v40 = v18 * int32(48)
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[2]))
				v43 = v40 + v42
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)))
				v46 = v44 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)) = uint16(v46)
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[0]))
				F_ResourceOwnerRemember(m, v49, base.I64_extend_i32_u(v18), int32(_a_F_PathNameCreateTemporaryFile_3))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[2]))
					v56 = v55 + v40
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v58
					v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
					v62 = v60 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v62)
					v65 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_PathNameCreateTemporaryFile[3])) = uint8(v65)
					m.G0 = v7 + int32(16)
					return v18
				}
			}
		}
	}
}
func F_add_path(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
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
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 float64
	_ = v291
	var v292 float64
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 float64
	_ = v323
	var v324 float64
	_ = v324
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 float64
	_ = v342
	var v343 float64
	_ = v343
	var v351 float64
	_ = v351
	var v352 float64
	_ = v352
	var v364 int32
	_ = v364
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v479 int32
	_ = v479
	var v482 float64
	_ = v482
	var v483 float64
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v609 int32
	_ = v609
	var v612 float64
	_ = v612
	var v613 float64
	_ = v613
	var v615 float64
	_ = v615
	var v616 float64
	_ = v616
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 float64
	_ = v634
	var v635 float64
	_ = v635
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_add_path[0]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v21 = v20
	goto L8
L7:
	;
	v21 = v3
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v682 != int32(283) {
		goto L273
	} else {
		goto L274
	}
L10:
	;
	v29 = v3
	v30 = v22
	v32 = v3
	goto L13
L11:
	;
	v675 = v3
	v678 = int32(0)
	goto L12
L12:
	;
	v679 = F_list_insert_nth(m, v678, v675, l1)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L4
	} else {
		goto L272
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v29 < v35 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v675 = v661
	v678 = v664
	goto L12
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v29<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v37 != v43 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v661 = v32
	goto L17
L17:
	;
	goto L14
L18:
	;
	if v649 != 0 {
		v29 = v648 + int32(1)
		v30 = v649
		v32 = v650
		goto L13
	} else {
		goto L271
	}
L19:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v630 <= v631 {
		goto L264
	} else {
		goto L265
	}
L20:
	;
	v106 = int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v107 != 0 {
		goto L48
	} else {
		goto L49
	}
L21:
	;
	if v37 < v43 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v42)+56))
	if base.F64_gt(v49, base.F64_mul(v50, float64(1.01))) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v48 = int32(1)
	goto L26
L25:
	;
	v48 = int32(2)
	goto L26
L26:
	;
	v105 = v48
	goto L20
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v55 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	if base.F64_lt(base.F64_mul(v49, float64(1.01)), v50) != 0 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v65 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.F64_gt(v65, base.F64_mul(v66, float64(1.01))) != 0 {
		v626 = int32(1)
		goto L19
	} else {
		goto L36
	}
L31:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+24)))
	if v58 != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+25)))
	if v60 == int32(1) {
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v105 = int32(2)
	goto L20
L35:
	;
	v105 = int32(2)
	goto L20
L36:
	;
	v105 = int32(2)
	goto L20
L37:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v75 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(v94, base.F64_mul(v95, float64(1.01))) != 0 {
		v105 = int32(2)
		goto L20
	} else {
		goto L47
	}
L40:
	;
	v84 = int32(1)
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(v86, base.F64_mul(v87, float64(1.01))) == int32(0) {
		v105 = v84
		goto L20
	} else {
		goto L46
	}
L41:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
	if v78 != 0 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v80 = int32(1)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+25)))
	if v81 != v80 {
		v105 = v80
		goto L20
	} else {
		goto L45
	}
L44:
	;
	v105 = int32(1)
	goto L20
L45:
	;
	goto L40
L46:
	;
	v626 = v84
	goto L19
L47:
	;
	v105 = base.F64_gt(v95, base.F64_mul(v94, float64(1.01)))
	goto L20
L48:
	;
	v110 = int32(0)
	goto L50
L49:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	v110 = v109
	goto L50
L50:
	;
	if v21 == v110 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v170 == int32(3) {
		v626 = v106
		goto L19
	} else {
		goto L75
	}
L52:
	;
	v170 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v119 = int32(0)
	goto L57
L55:
	;
	if v159 != 0 {
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v154 = int32(0)
	if v141 != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v123 = int32(0)
	if v21 == v123 {
		v133 = v123
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v170 = int32(3)
	goto L51
L59:
	;
	if v110 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v127 <= v119 {
		v133 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v133 = v129 + v119<<(uint(int32(2))%32)
	goto L59
L62:
	;
	v139 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if base.B2i32(v133 == v139)|base.B2i32(v141 == v139) != 0 {
		goto L56
	} else {
		goto L67
	}
L63:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v119 < v134 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v136 = int32(0)
	v159 = base.B2i32(v133 == v136)
	v161 = v136
	goto L55
L66:
	;
	goto L65
L67:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141+v119<<(uint(int32(2))%32))))
	if v149 == v151 {
		v119 = v119 + int32(1)
		goto L57
	} else {
		goto L68
	}
L68:
	;
	goto L58
L69:
	;
	v158 = int32(2)
	goto L71
L70:
	;
	v158 = v154
	goto L71
L71:
	;
	v159 = base.B2i32(v133 == v154)
	v161 = v158
	goto L55
L72:
	;
	v163 = v161
	goto L74
L73:
	;
	v163 = int32(1)
	goto L74
L74:
	;
	v170 = v163
	goto L51
L75:
	;
	switch v105 - int32(1) {
	case 0:
		goto L80
	case 1:
		goto L78
	default:
		goto L81
	}
L76:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	v626 = base.B2i32(base.Ui32(v623) < base.Ui32(v622))
	goto L19
L77:
	;
	v615 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v616 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_ge(v615, v616) == int32(0) {
		v626 = v106
		goto L19
	} else {
		goto L262
	}
L78:
	;
	if v170 == int32(1) {
		v626 = v106
		goto L19
	} else {
		goto L218
	}
L79:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v496 = F_list_delete_nth_cell(m, v495, v29)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L215
	}
L80:
	;
	if v170 == int32(2) {
		v626 = v106
		goto L19
	} else {
		goto L170
	}
L81:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v175 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v178 = v176
	goto L84
L83:
	;
	v178 = int32(0)
	goto L84
L84:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v179 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v182 = v180
	goto L87
L86:
	;
	v182 = int32(0)
	goto L87
L87:
	;
	if v178 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	switch v170 - int32(1) {
	case 0:
		goto L125
	case 1:
		goto L124
	default:
		goto L123
	}
L89:
	;
	v276 = base.B2i32(v182 != int32(0))
	goto L88
L90:
	;
	goto L91
L91:
	;
	if v182 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v276 = int32(2)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v199 < v200 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v202 = v199
	goto L97
L96:
	;
	v202 = v200
	goto L97
L97:
	;
	if v202 <= int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v205 = int32(1)
	goto L100
L99:
	;
	v205 = v202
	goto L100
L100:
	;
	v206 = int32(8)
	v210 = int32(0)
	v212 = v210
	v213 = v210
	goto L103
L101:
	;
	v276 = int32(3)
	goto L88
L102:
	;
	v276 = v263
	goto L88
L103:
	;
	v223 = v213 << (uint(int32(2)) % 32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v178+v206+v223)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223+(v182+v206))))
	if v225&(v227^int32(-1)) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	if v200 < v199 {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v249 = v213 + int32(1)
	if v249 != v205 {
		v212 = v247
		v213 = v249
		goto L103
	} else {
		goto L112
	}
L106:
	;
	if base.B2i32(v212 == int32(1))|v227&(v225^int32(-1)) != 0 {
		v263 = int32(3)
		goto L102
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v227&(v225^int32(-1)) == int32(0) {
		v247 = v212
		goto L105
	} else {
		goto L110
	}
L109:
	;
	v247 = int32(2)
	goto L105
L110:
	;
	if v212 == int32(2) {
		goto L101
	} else {
		goto L111
	}
L111:
	;
	v247 = int32(1)
	goto L105
L112:
	;
	goto L104
L113:
	;
	if v247 == int32(1) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	if v200 <= v199 {
		v263 = v247
		goto L102
	} else {
		goto L119
	}
L116:
	;
	v256 = int32(3)
	goto L118
L117:
	;
	v256 = int32(2)
	goto L118
L118:
	;
	v276 = v256
	goto L88
L119:
	;
	if v247 == int32(2) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v262 = int32(3)
	goto L122
L121:
	;
	v262 = int32(1)
	goto L122
L122:
	;
	v263 = v262
	goto L102
L123:
	;
	switch v276 {
	case 0:
		goto L132
	case 1:
		goto L131
	case 2:
		goto L77
	default:
		v626 = v106
		goto L19
	}
L124:
	;
	if v276&int32(-3) != 0 {
		v626 = v106
		goto L19
	} else {
		goto L129
	}
L125:
	;
	if base.Ui32(int32(1)) < base.Ui32(v276) {
		v626 = v106
		goto L19
	} else {
		goto L126
	}
L126:
	;
	v281 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_le(v281, v282) == int32(0) {
		v626 = v106
		goto L19
	} else {
		goto L127
	}
L127:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v287) <= base.Ui32(v286) {
		goto L79
	} else {
		goto L128
	}
L128:
	;
	v626 = v106
	goto L19
L129:
	;
	v291 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_ge(v291, v292) != 0 {
		goto L76
	} else {
		goto L130
	}
L130:
	;
	v626 = v106
	goto L19
L131:
	;
	v367 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_le(v367, v368) == int32(0) {
		v626 = v106
		goto L19
	} else {
		goto L168
	}
L132:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v295) < base.Ui32(v294) {
		goto L79
	} else {
		goto L133
	}
L133:
	;
	v297 = int32(0)
	if base.Ui32(v294) < base.Ui32(v295) {
		v626 = v297
		goto L19
	} else {
		goto L134
	}
L134:
	;
	v299 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_lt(v299, v300) != 0 {
		goto L79
	} else {
		goto L135
	}
L135:
	;
	if base.F64_gt(v299, v300) != 0 {
		v626 = v297
		goto L19
	} else {
		goto L136
	}
L136:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v303 != v304 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v364 == int32(1) {
		goto L79
	} else {
		goto L167
	}
L138:
	;
	if v303 < v304 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v310 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v311 = *(*float64)(unsafe.Add(mBase, uint32(v42)+56))
	if base.F64_gt(v310, base.F64_mul(v311, float64(1.0000000001))) != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	v309 = int32(1)
	goto L143
L142:
	;
	v309 = int32(2)
	goto L143
L143:
	;
	v364 = v309
	goto L137
L144:
	;
	v364 = int32(2)
	goto L137
L145:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v316 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	if base.F64_lt(base.F64_mul(v310, float64(1.0000000001)), v311) != 0 {
		goto L156
	} else {
		goto L157
	}
L148:
	;
	v323 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	v324 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.F64_gt(v323, base.F64_mul(v324, float64(1.0000000001))) == int32(0) {
		goto L144
	} else {
		goto L154
	}
L149:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+24)))
	if v319 != 0 {
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+25)))
	if v320 != int32(1) {
		goto L144
	} else {
		goto L153
	}
L152:
	;
	goto L144
L153:
	;
	goto L148
L154:
	;
	v364 = int32(3)
	goto L137
L155:
	;
	v364 = int32(1)
	goto L137
L156:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v335 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v351 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v352 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(v351, base.F64_mul(v352, float64(1.0000000001))) != 0 {
		v364 = int32(2)
		goto L137
	} else {
		goto L166
	}
L159:
	;
	v342 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(v342, base.F64_mul(v343, float64(1.0000000001))) == int32(0) {
		goto L155
	} else {
		goto L165
	}
L160:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+24)))
	if v338 != 0 {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+25)))
	if v339 != int32(1) {
		goto L155
	} else {
		goto L164
	}
L163:
	;
	goto L155
L164:
	;
	goto L159
L165:
	;
	v364 = int32(3)
	goto L137
L166:
	;
	v364 = base.F64_gt(v352, base.F64_mul(v351, float64(1.0000000001)))
	goto L137
L167:
	;
	v626 = v297
	goto L19
L168:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v373) <= base.Ui32(v372) {
		goto L79
	} else {
		goto L169
	}
L169:
	;
	v626 = v106
	goto L19
L170:
	;
	v377 = int32(0)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v379 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v381 = v380
	goto L173
L172:
	;
	v381 = v377
	goto L173
L173:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v382 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v384 = v383
	goto L176
L175:
	;
	v384 = v377
	goto L176
L176:
	;
	v385 = int32(1)
	if v381 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	if base.Ui32(int32(1)) < base.Ui32(v479) {
		v626 = v385
		goto L19
	} else {
		goto L212
	}
L178:
	;
	v479 = base.B2i32(v384 != int32(0))
	goto L177
L179:
	;
	goto L180
L180:
	;
	if v384 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v479 = int32(2)
	goto L177
L182:
	;
	goto L183
L183:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v402 < v403 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v405 = v402
	goto L186
L185:
	;
	v405 = v403
	goto L186
L186:
	;
	if v405 <= int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v408 = int32(1)
	goto L189
L188:
	;
	v408 = v405
	goto L189
L189:
	;
	v409 = int32(8)
	v413 = int32(0)
	v415 = v413
	v416 = v413
	goto L192
L190:
	;
	v479 = int32(3)
	goto L177
L191:
	;
	v479 = v466
	goto L177
L192:
	;
	v426 = v416 << (uint(int32(2)) % 32)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v381+v409+v426)))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426+(v384+v409))))
	if v428&(v430^int32(-1)) != 0 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	if v403 < v402 {
		goto L202
	} else {
		goto L203
	}
L194:
	;
	v452 = v416 + int32(1)
	if v452 != v408 {
		v415 = v450
		v416 = v452
		goto L192
	} else {
		goto L201
	}
L195:
	;
	if base.B2i32(v415 == int32(1))|v430&(v428^int32(-1)) != 0 {
		v466 = int32(3)
		goto L191
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	if v430&(v428^int32(-1)) == int32(0) {
		v450 = v415
		goto L194
	} else {
		goto L199
	}
L198:
	;
	v450 = int32(2)
	goto L194
L199:
	;
	if v415 == int32(2) {
		goto L190
	} else {
		goto L200
	}
L200:
	;
	v450 = int32(1)
	goto L194
L201:
	;
	goto L193
L202:
	;
	if v450 == int32(1) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	if v403 <= v402 {
		v466 = v450
		goto L191
	} else {
		goto L208
	}
L205:
	;
	v459 = int32(3)
	goto L207
L206:
	;
	v459 = int32(2)
	goto L207
L207:
	;
	v479 = v459
	goto L177
L208:
	;
	if v450 == int32(2) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v465 = int32(3)
	goto L211
L210:
	;
	v465 = int32(1)
	goto L211
L211:
	;
	v466 = v465
	goto L191
L212:
	;
	v482 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v483 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_le(v482, v483) == int32(0) {
		v626 = v385
		goto L19
	} else {
		goto L213
	}
L213:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v487) < base.Ui32(v488) {
		v626 = v385
		goto L19
	} else {
		goto L214
	}
L214:
	;
	goto L79
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v496
	v500 = v29 - int32(1)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v501 == int32(283) {
		v648 = v500
		v649 = v496
		v650 = v32
		goto L18
	} else {
		goto L216
	}
L216:
	;
	F_pfree(m, v42)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	v648 = v500
	v649 = v496
	v650 = v32
	goto L18
L218:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v508 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	v511 = v509
	goto L221
L220:
	;
	v511 = int32(0)
	goto L221
L221:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v512 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v515 = v513
	goto L224
L223:
	;
	v515 = int32(0)
	goto L224
L224:
	;
	if v511 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	if v609&int32(-3) != 0 {
		v626 = v106
		goto L19
	} else {
		goto L260
	}
L226:
	;
	v609 = base.B2i32(v515 != int32(0))
	goto L225
L227:
	;
	goto L228
L228:
	;
	if v515 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v609 = int32(2)
	goto L225
L230:
	;
	goto L231
L231:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	if v532 < v533 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v535 = v532
	goto L234
L233:
	;
	v535 = v533
	goto L234
L234:
	;
	if v535 <= int32(1) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v538 = int32(1)
	goto L237
L236:
	;
	v538 = v535
	goto L237
L237:
	;
	v539 = int32(8)
	v543 = int32(0)
	v545 = v543
	v546 = v543
	goto L240
L238:
	;
	v609 = int32(3)
	goto L225
L239:
	;
	v609 = v596
	goto L225
L240:
	;
	v556 = v546 << (uint(int32(2)) % 32)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v511+v539+v556)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556+(v515+v539))))
	if v558&(v560^int32(-1)) != 0 {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	if v533 < v532 {
		goto L250
	} else {
		goto L251
	}
L242:
	;
	v582 = v546 + int32(1)
	if v582 != v538 {
		v545 = v580
		v546 = v582
		goto L240
	} else {
		goto L249
	}
L243:
	;
	if base.B2i32(v545 == int32(1))|v560&(v558^int32(-1)) != 0 {
		v596 = int32(3)
		goto L239
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if v560&(v558^int32(-1)) == int32(0) {
		v580 = v545
		goto L242
	} else {
		goto L247
	}
L246:
	;
	v580 = int32(2)
	goto L242
L247:
	;
	if v545 == int32(2) {
		goto L238
	} else {
		goto L248
	}
L248:
	;
	v580 = int32(1)
	goto L242
L249:
	;
	goto L241
L250:
	;
	if v580 == int32(1) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L252
L252:
	;
	if v533 <= v532 {
		v596 = v580
		goto L239
	} else {
		goto L256
	}
L253:
	;
	v589 = int32(3)
	goto L255
L254:
	;
	v589 = int32(2)
	goto L255
L255:
	;
	v609 = v589
	goto L225
L256:
	;
	if v580 == int32(2) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v595 = int32(3)
	goto L259
L258:
	;
	v595 = int32(1)
	goto L259
L259:
	;
	v596 = v595
	goto L239
L260:
	;
	v612 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v613 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_ge(v612, v613) != 0 {
		goto L76
	} else {
		goto L261
	}
L261:
	;
	v626 = v106
	goto L19
L262:
	;
	goto L76
L263:
	;
	if v626 == int32(0) {
		goto L9
	} else {
		goto L270
	}
L264:
	;
	if v630 != v631 {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	if v626 == int32(0) {
		goto L9
	} else {
		goto L269
	}
L267:
	;
	v634 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v635 = *(*float64)(unsafe.Add(mBase, uint32(v42)+56))
	if base.F64_ge(v634, v635) == int32(0) {
		goto L263
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v648 = v29
	v649 = v30
	v650 = v29 + int32(1)
	goto L18
L270:
	;
	v648 = v29
	v649 = v30
	v650 = v32
	goto L18
L271:
	;
	v661 = v650
	goto L17
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v679
	return
L273:
	;
	F_pfree(m, l1)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	return
L276:
	;
	goto L275
}
func F_add_path_precheck(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 float64
	_ = v48
	var v56 float64
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	v7 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v11 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	if l5 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = int32(0)
	goto L9
L8:
	;
	v22 = l4
	goto L9
L9:
	;
	if l5 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v25 = int32(25)
	goto L12
L11:
	;
	v25 = int32(24)
	goto L12
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v25))))
	v35 = v7
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v35<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if l1 != v43 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return v192
L15:
	;
	if v27 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	if v43 <= l1 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v42)+56))
	if base.F64_le(l3, base.F64_mul(v48, float64(1.01))) == int32(0) {
		goto L15
	} else {
		goto L20
	}
L19:
	;
	return int32(1)
L20:
	;
	return int32(1)
L21:
	;
	goto L14
L22:
	;
	v186 = int32(1)
	v188 = v35 + v186
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v188 < v189 {
		v35 = v188
		goto L13
	} else {
		goto L70
	}
L23:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(l2, base.F64_mul(v56, float64(1.01))) == int32(0) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v63 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v66 = int32(0)
	goto L29
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	v66 = v65
	goto L29
L29:
	;
	if v22 == v66 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v126&int32(-3) != 0 {
		goto L22
	} else {
		goto L54
	}
L31:
	;
	v126 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v75 = int32(0)
	goto L36
L34:
	;
	if v115 != 0 {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	v110 = int32(0)
	if v97 != 0 {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v79 = int32(0)
	if v22 == v79 {
		v89 = v79
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v126 = int32(3)
	goto L30
L38:
	;
	if v66 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v83 <= v75 {
		v89 = int32(0)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v89 = v85 + v75<<(uint(int32(2))%32)
	goto L38
L41:
	;
	v95 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if base.B2i32(v89 == v95)|base.B2i32(v97 == v95) != 0 {
		goto L35
	} else {
		goto L46
	}
L42:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v75 < v90 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v92 = int32(0)
	v115 = base.B2i32(v89 == v92)
	v117 = v92
	goto L34
L45:
	;
	goto L44
L46:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97+v75<<(uint(int32(2))%32))))
	if v105 == v107 {
		v75 = v75 + int32(1)
		goto L36
	} else {
		goto L47
	}
L47:
	;
	goto L37
L48:
	;
	v114 = int32(2)
	goto L50
L49:
	;
	v114 = v110
	goto L50
L50:
	;
	v115 = base.B2i32(v89 == v110)
	v117 = v114
	goto L34
L51:
	;
	v119 = v117
	goto L53
L52:
	;
	v119 = int32(1)
	goto L53
L53:
	;
	v126 = v119
	goto L30
L54:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v129 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v132 = v130
	goto L57
L56:
	;
	v132 = int32(0)
	goto L57
L57:
	;
	v133 = int32(0)
	if base.B2i32(l5 == v133)|base.B2i32(v132 == v133) != 0 {
		v179 = base.B2i32(l5|v132 == v133)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v179 != 0 {
		v192 = int32(0)
		goto L21
	} else {
		goto L69
	}
L59:
	;
	goto L58
L60:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v147 != v148 {
		v179 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v150 = int32(1)
	if v147 <= v150 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v153 = v150
	goto L64
L63:
	;
	v153 = v147
	goto L64
L64:
	;
	v154 = int32(8)
	v159 = int32(0)
	goto L65
L65:
	;
	v167 = v159 << (uint(int32(2)) % 32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l5+v154+v167)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v132+v154+v167)))
	v172 = base.B2i32(v169 == v171)
	if v169 != v171 {
		v179 = v172
		goto L59
	} else {
		goto L67
	}
L66:
	;
	v179 = v172
	goto L59
L67:
	;
	v175 = v159 + int32(1)
	if v175 != v153 {
		v159 = v175
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	goto L22
L70:
	;
	v192 = v186
	goto L21
}
func F_check_search_path(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v98 int64
	_ = v98
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v116 int64
	_ = v116
	var v121 int64
	_ = v121
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v191 int64
	_ = v191
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_check_search_path[0]))
	if v15 == v4 {
		v256 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v298
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_search_path[1])) = v206
	v298 = int32(1)
	goto L1
L3:
	;
	v260 = F_pstrdup(m, v13)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L68
	}
L4:
	;
	F_spcache_init(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_check_search_path[2]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_check_search_path[1]))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v13
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_check_search_path[3]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)))
	v77 = (int64(base.Ui64(v72)>>(uint(int64(23))%64)) ^ v72) * int64(2388976653695081527)
	v81 = int64(-8645972361240307355)
	v84 = (v77 ^ int64(base.Ui64(v77)>>(uint(int64(47))%64)) ^ v81) * v81
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v86 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 != v23 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(v33 == int32(0))|base.B2i32(v33 != v36) != 0 {
		v54 = v33
		v55 = v36
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v54-v55 != 0 {
		goto L7
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v39 = v30
	v40 = v13
	goto L13
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v44 == int32(0) {
		v54 = v44
		v55 = v43
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v54 = v44
	v55 = v43
	goto L11
L15:
	;
	v47 = int32(1)
	if v44 == v43 {
		v39 = v39 + v47
		v40 = v40 + v47
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v298 = int32(1)
	goto L1
L18:
	;
	v199 = v65 & base.I32_wrap_i64(int64(base.Ui64(v191)>>(uint(int64(47))%64))^v191-int64(base.Ui64(v191)>>(uint(int64(32))%64)))
	v202 = v64 + v199*int32(24)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+22)))
	if v203 == int32(0) {
		v256 = v23
		goto L3
	} else {
		goto L53
	}
L19:
	;
	v191 = (base.I64_extend_i32_s(v177-v85) + int64(base.Ui64(v179)>>(uint(int64(23))%64)) ^ v179) * int64(2388976653695081527)
	goto L18
L20:
	;
	v177 = v85
	v179 = v84
	goto L19
L21:
	;
	goto L22
L22:
	;
	v89 = v85
	v91 = v84
	v94 = v86
	goto L23
L23:
	;
	v98 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v98 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v177 = v175
	v179 = v174
	goto L19
L25:
	;
	v168 = (v163 ^ int64(base.Ui64(v163)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v174 = (v91 ^ int64(base.Ui64(v168)>>(uint(int64(47))%64)) ^ v168) * int64(-8645972361240307355)
	v175 = v89 + v162
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v176 != 0 {
		v89 = v175
		v91 = v174
		v94 = v176
		goto L23
	} else {
		goto L52
	}
L26:
	;
	v162 = v156
	v163 = base.I64_extend8_s(base.I64_extend_i32_u(v94)) | v157
	goto L25
L27:
	;
	v156 = int32(1)
	v157 = int64(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+2)))
	if v105 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v156 = v152
	v157 = v98<<(uint(int64(8))%64) | v153
	goto L26
L31:
	;
	v152 = int32(2)
	v153 = int64(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
	if v110 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v111 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v111 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	v152 = int32(3)
	v153 = v105 << (uint(int64(16)) % 64)
	goto L30
L37:
	;
	v147 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v89))))
	v162 = v145
	v163 = v146 | v147
	goto L25
L38:
	;
	v145 = int32(4)
	v146 = int64(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v116 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+5)))
	if v116 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v145 = v138
	v146 = v139 | v111<<(uint(int64(32))%64)
	goto L37
L42:
	;
	v138 = int32(5)
	v139 = int64(0)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v121 = int64(*(*int8)(unsafe.Add(mBase, uint32(v89)+6)))
	if v121 == int64(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v138 = v133
	v139 = v132 | v116<<(uint(int64(40))%64)
	goto L41
L46:
	;
	v132 = int64(0)
	v133 = int32(6)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+7)))
	if v126 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	v162 = int32(8)
	v163 = v128
	goto L25
L50:
	;
	goto L51
L51:
	;
	v132 = v121 << (uint(int64(48)) % 64)
	v133 = int32(7)
	goto L45
L52:
	;
	goto L24
L53:
	;
	v206 = v202
	v209 = v199
	goto L54
L54:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v23 == v214 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v256 = v23
	goto L3
L56:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(v219 == int32(0))|base.B2i32(v219 != v222) != 0 {
		v240 = v219
		v241 = v222
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	v247 = (v209 + int32(1)) & v65
	v250 = v64 + v247*int32(24)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+22)))
	if v251 != 0 {
		v206 = v250
		v209 = v247
		goto L54
	} else {
		goto L67
	}
L59:
	;
	if v240-v241 == int32(0) {
		goto L2
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v225 = v216
	v226 = v13
	goto L62
L62:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	if v230 == int32(0) {
		v240 = v230
		v241 = v229
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v240 = v230
	v241 = v229
	goto L60
L64:
	;
	v233 = int32(1)
	if v230 == v229 {
		v225 = v225 + v233
		v226 = v226 + v233
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L58
L67:
	;
	goto L55
L68:
	;
	v265 = F_SplitIdentifierString(m, v260, int32(44), v11+int32(8))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	if v265 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_check_search_path[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_search_path[5])) = v271
	goto L73
L71:
	;
	goto L72
L72:
	;
	F_pfree(m, v260)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L77
	}
L73:
	;
	v277 = F_format_elog_string(m, int32(_a_F_check_search_path_0), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_search_path[6])) = v277
	F_pfree(m, v260)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_list_free(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v298 = int32(0)
	goto L1
L77:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_list_free(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v290 = int32(1)
	if v15 == int32(0) {
		v298 = v290
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v293 = F_spcache_insert(m, v13, v256)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v298 = v290
	goto L1
}
func F_create_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 float64) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 float64
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 float64
	_ = v332
	var v334 int32
	_ = v334
	var v336 float64
	_ = v336
	var v338 float64
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	v7 = l6
	v16 = F_palloc0(m, int32(96))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(293)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(338)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v22
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v191
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v194)
	if v7 != 0 {
		goto L32
	} else {
		goto L33
	}
L4:
	;
	if l4 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32|base.B2i32(v29 == int32(0)) != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = F_get_baserel_parampathinfo(m, l0, l1, l4)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v191 = v36
	goto L3
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v38 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v176 = int32(0)
	goto L10
L10:
	;
	v191 = v176
	goto L3
L11:
	;
	v176 = v156
	goto L10
L12:
	;
	v134 = F_palloc0(m, int32(24))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L30
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v41 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v56 = int32(0)
	goto L15
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v56<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v64 = int32(0)
	if base.B2i32(v63 == v64)|base.B2i32(l4 == v64) != 0 {
		v110 = base.B2i32(v63|l4 == v64)
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L12
L17:
	;
	if v110 != 0 {
		v156 = v62
		goto L11
	} else {
		goto L28
	}
L18:
	;
	goto L17
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v78 != v79 {
		v110 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v81 = int32(1)
	if v78 <= v81 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = v81
	goto L23
L22:
	;
	v84 = v78
	goto L23
L23:
	;
	v85 = int32(8)
	v90 = int32(0)
	goto L24
L24:
	;
	v98 = v90 << (uint(int32(2)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v63+v85+v98)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l4+v85+v98)))
	v103 = base.B2i32(v100 == v102)
	if v100 != v102 {
		v110 = v103
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v110 = v103
	goto L18
L26:
	;
	v106 = v90 + int32(1)
	if v106 != v84 {
		v90 = v106
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v116 = v56 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v116 < v117 {
		v56 = v116
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	v136 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v134)+8)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(281)
	*(*int64)(unsafe.Add(mBase, uint32(v134)+16)) = v136
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v144 = F_lappend(m, v143, v134)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v144
	v156 = v134
	goto L11
L32:
	;
	F_list_sort(m, v29, int32(928))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v29 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	F_list_sort(m, v201, int32(929))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v207 = v205
	goto L39
L38:
	;
	v207 = int32(0)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v210 = F_list_concat(m, v29, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v210
	v213 = float64(-1)
	if l0 == int32(0) {
		v272 = v213
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+80)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	if v274 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v218 = int32(0)
	if base.B2i32(v216 == v218)|base.B2i32(v217 == v218) != 0 {
		v264 = base.B2i32(v216|v217 == v218)
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v264 == int32(0) {
		v272 = v213
		goto L41
	} else {
		goto L54
	}
L44:
	;
	goto L43
L45:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v232 != v233 {
		v264 = int32(0)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v235 = int32(1)
	if v232 <= v235 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v238 = v235
	goto L49
L48:
	;
	v238 = v232
	goto L49
L49:
	;
	v239 = int32(8)
	v244 = int32(0)
	goto L50
L50:
	;
	v252 = v244 << (uint(int32(2)) % 32)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v216+v239+v252)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v217+v239+v252)))
	v257 = base.B2i32(v254 == v256)
	if v254 != v256 {
		v264 = v257
		goto L44
	} else {
		goto L52
	}
L51:
	;
	v264 = v257
	goto L44
L52:
	;
	v260 = v244 + int32(1)
	if v260 != v238 {
		v244 = v260
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v271 = *(*float64)(unsafe.Add(mBase, uint32(l0)+320))
	v272 = v271
	goto L41
L55:
	;
	if base.F64_ge(l7, float64(0)) != 0 {
		goto L74
	} else {
		goto L75
	}
L56:
	;
	F_cost_append(m, v16, l0)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L73
	}
L57:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if int32(0) < v277 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v282 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v325 != int32(1) {
		goto L56
	} else {
		goto L67
	}
L61:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)))
	if v295 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L60
L63:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v296+v282<<(uint(int32(2))%32))))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+21)))
	v303 = v301
	goto L65
L64:
	;
	v303 = int32(0)
	goto L65
L65:
	;
	v304 = int32(1)
	v305 = v303 & v304
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v305)
	v308 = v282 + v304
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v308 < v309 {
		v282 = v308
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+20)))
	if v7 == v330 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v329)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v342
	goto L55
L69:
	;
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v329)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v329)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v334
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v329)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = v336
	v338 = *(*float64)(unsafe.Add(mBase, uint32(v329)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = v338
	goto L68
L70:
	;
	goto L71
L71:
	;
	F_cost_append(m, v16, l0)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	goto L55
L74:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = l7
	goto L76
L75:
	;
	goto L76
L76:
	;
	return v16
}
func F_create_gather_merge_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v134 float64
	_ = v134
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v17 = F_palloc0(m, int32(80))
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
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(299)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if l4 == v23 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v76 = int32(1)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v32 = int32(0)
	goto L8
L7:
	;
	v76 = v68
	goto L3
L8:
	;
	v36 = int32(0)
	if l4 == v36 {
		v46 = v36
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v68 = int32(0)
	goto L7
L10:
	;
	if v23 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v40 <= v32 {
		v46 = int32(0)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v46 = v42 + v32<<(uint(int32(2))%32)
	goto L10
L13:
	;
	v52 = base.B2i32(v46 == int32(0))
	if v46 == int32(0) {
		v68 = v52
		goto L7
	} else {
		goto L18
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v32 < v47 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v76 = base.B2i32(v46 == int32(0))
	goto L3
L17:
	;
	goto L16
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v55 == int32(0) {
		v68 = v52
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+v55)))
	if v62 == v64 {
		v32 = v32 + int32(1)
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(373)
	v81 = F_get_baserel_parampathinfo(m, l0, l1, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
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
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L33
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = l2
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v81
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v87
	if l3 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = l3
	goto L27
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v91 = v90
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v91
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v95 = float64(0)
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	if l5 != 0 {
		v104 = l5
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v104)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+32)) = v105
	v108 = *(*float64)(unsafe.Add(mBase, _c_F_create_gather_merge_path[0]))
	v110 = *(*float64)(unsafe.Add(mBase, _c_F_create_gather_merge_path[1]))
	v112 = *(*float64)(unsafe.Add(mBase, _c_F_create_gather_merge_path[2]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+40))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v114 + base.B2i32(v115&int64(32768) == int64(0))
	v122 = base.F64_add(v112, v112)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v126 = base.F64_add(base.F64_convert_i32_s(v123), float64(1))
	v128 = F_log(m, v126)
	mBase = m.M
	v130 = base.F64_div(v128, float64(0.693147180559945))
	v132 = float64(0)
	v134 = base.F64_add(v110, base.F64_add(base.F64_mul(base.F64_mul(v122, v126), v130), v132))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = base.F64_add(base.F64_add(v94, v95), v134)
	*(*float64)(unsafe.Add(mBase, uint32(v17)+56)) = base.F64_add(base.F64_add(v97, v95), base.F64_add(v134, base.F64_add(base.F64_mul(base.F64_mul(v105, v108), float64(1.05)), base.F64_add(base.F64_mul(v112, v105), base.F64_add(base.F64_mul(base.F64_mul(v105, v122), v130), v132)))))
	return v17
L29:
	;
	if v81 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v104 = v81 + int32(8)
	goto L28
L31:
	;
	goto L32
L32:
	;
	v104 = l1 + int32(16)
	goto L28
L33:
	;
	F_errmsg_internal(m, int32(_a_F_create_gather_merge_path_0), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_create_gather_merge_path_1), int32(1834), int32(_a_F_create_gather_merge_path_2))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v60 float64
	_ = v60
	v8 = F_palloc0(m, int32(80))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1571958030641)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v17
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v21 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v26 = v24
		} else {
			v26 = int32(0)
		}
		v27 = int32(1)
		v28 = v26 & v27
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)) = uint8(v28)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v30
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v35 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v36 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
		v41 = *(*int32)(unsafe.Add(mBase, _c_F_create_sort_path[0]))
		v43 = m.G0
		v44 = int32(16)
		v45 = v43 - v44
		m.G0 = v45
		F_cost_tuplesort(m, v45+int32(8), v45, v36, v38, float64(0), v41, l3)
		mBase = m.M
		v50 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v36
		v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_sort_path[1])))
		v54 = base.F64_add(v35, v50)
		*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = v54
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v34 + (v53 ^ v27)
		v60 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
		*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(v54, v60)
		m.G0 = v45 + v44
		return v8
	}
}
func F_path_area(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v139 float64
	_ = v139
	var v140 int32
	_ = v140
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v146 float64
	_ = v146
	var v153 float64
	_ = v153
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v165 float64
	_ = v165
	var v166 int32
	_ = v166
	var v169 float64
	_ = v169
	var v175 float64
	_ = v175
	var v176 int32
	_ = v176
	var v177 float64
	_ = v177
	v2 = float64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v153 = base.F64_abs(v146)
	v155 = base.F64_mul(v153, float64(0.5))
	v157 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v155), v157)|base.F64_eq(v153, v157) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v24 = v11 + int32(16)
	v26 = int32(0)
	v28 = v2
	v31 = v16
	goto L9
L3:
	;
	return int64(0)
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v16 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
	return int64(0)
L8:
	;
	v146 = v2
	goto L1
L9:
	;
	v35 = int32(4)
	v37 = v24 + v26<<(uint(v35)%32)
	v38 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
	v40 = v26 + int32(1)
	v41 = base.I32_rem_s(v40, v31)
	v44 = v24 + v41<<(uint(v35)%32)
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = base.F64_mul(v38, v45)
	v48 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v46), v48)|base.F64_eq(base.F64_abs(v38), v48)|base.F64_eq(base.F64_abs(v45), v48) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v146 = v141
	goto L1
L11:
	;
	v76 = math.Float64frombits(uint64(0x7ff0000000000000))
	v78 = base.F64_add(v28, v74)
	if base.F64_eq(base.F64_abs(v28), v76)|base.F64_ne(base.F64_abs(v78), v76)|base.F64_eq(base.F64_abs(v74), v76) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v61 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v63 = float64(0)
	if base.F64_eq(v38, v63)|base.F64_ne(v46, v63)|base.F64_eq(v45, v63) != 0 {
		v74 = v46
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v74 = v61
	goto L11
L16:
	;
	v72 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v74 = v72
	goto L11
L18:
	;
	v90 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	v92 = v78
	goto L20
L20:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v37)+8))
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v44)))
	v95 = base.F64_mul(v93, v94)
	v97 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v95), v97)|base.F64_eq(base.F64_abs(v93), v97)|base.F64_eq(base.F64_abs(v94), v97) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v92 = v90
	goto L20
L22:
	;
	v125 = math.Float64frombits(uint64(0x7ff0000000000000))
	v127 = base.F64_sub(v92, v123)
	if base.F64_eq(base.F64_abs(v92), v125)|base.F64_ne(base.F64_abs(v127), v125)|base.F64_eq(base.F64_abs(v123), v125) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v110 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v112 = float64(0)
	if base.F64_eq(v93, v112)|base.F64_ne(v95, v112)|base.F64_eq(v94, v112) != 0 {
		v123 = v95
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v123 = v110
	goto L22
L27:
	;
	v121 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v123 = v121
	goto L22
L29:
	;
	v139 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	v141 = v127
	goto L31
L31:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v40 < v142 {
		v26 = v40
		v28 = v141
		v31 = v142
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v141 = v139
	goto L31
L33:
	;
	goto L10
L34:
	;
	v165 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v169 = float64(0)
	if base.F64_eq(v146, v169)|base.F64_ne(v155, v169) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	return base.I64_reinterpret_f64(v165)
L38:
	;
	v177 = v155
	goto L40
L39:
	;
	v175 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L41
	}
L40:
	;
	return base.I64_reinterpret_f64(v177)
L41:
	;
	v177 = v175
	goto L40
}
func F_path_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 float64
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v96 float64
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v117 int32
	_ = v117
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v129 float64
	_ = v129
	var v131 float64
	_ = v131
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v142 float64
	_ = v142
	var v143 int32
	_ = v143
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v180 float64
	_ = v180
	var v182 float64
	_ = v182
	var v184 float64
	_ = v184
	var v186 float64
	_ = v186
	var v188 float64
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v195 float64
	_ = v195
	var v197 int32
	_ = v197
	var v198 float64
	_ = v198
	var v200 float64
	_ = v200
	var v205 float64
	_ = v205
	var v206 int32
	_ = v206
	var v213 float64
	_ = v213
	var v219 float64
	_ = v219
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v240 float64
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v258 float64
	_ = v258
	var v261 int32
	_ = v261
	var v280 int32
	_ = v280
	var v301 int64
	_ = v301
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(0) < v30 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v20 - int32(-64)
	return v301
L5:
	;
	v301 = base.I64_reinterpret_f64(v258)
	goto L4
L6:
	;
	v33 = int32(16)
	v34 = v28 + v33
	v43 = v30
	v46 = v2
	v48 = v2
	v54 = float64(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v280)
	v301 = int64(0)
	goto L4
L9:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v250 != 0 {
		goto L5
	} else {
		goto L49
	}
L11:
	;
	v261 = v48 + int32(1)
	if v261 < v247 {
		v43 = v247
		v46 = v250
		v48 = v261
		v54 = v258
		goto L9
	} else {
		goto L48
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v58 == int32(0) {
		v247 = v43
		v250 = v46
		v258 = v54
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v61 = v48
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v62 <= int32(0) {
		v247 = v43
		v250 = v46
		v258 = v54
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v61 = v43
	goto L14
L16:
	;
	v65 = int32(4)
	v67 = v23 + v33 + v48<<(uint(v65)%32)
	v70 = v23 + v61<<(uint(v65)%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v71 == int32(0) {
		v117 = v46
		v119 = v54
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v120 < int32(2) {
		v232 = v117
		v240 = v119
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+32)) = v74
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+40)) = v76
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = v78
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = v80
	v84 = v28 + v62<<(uint(int32(4))%32)
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v84)))
	*(*float64)(unsafe.Add(mBase, uint32(v20))) = v85
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v84)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+8)) = v87
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+16)) = v89
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v28)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+24)) = v91
	v96 = F_lseg_closept_lseg(m, int32(0), v18+int32(-32), v20)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v46 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v117 = int32(1)
	v119 = v96
	goto L17
L21:
	;
	v100 = int32(1)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v96)&int64(9223372036854775807)) {
		v117 = v100
		v119 = v54
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if base.F64_lt(v96, v54) == int32(0) {
		v117 = v100
		v119 = v54
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v247 = v242
	v250 = v232
	v258 = v240
	goto L11
L26:
	;
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+32)) = v123
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+40)) = v125
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = v127
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = v129
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
	*(*float64)(unsafe.Add(mBase, uint32(v20))) = v131
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v28+int32(24))))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+8)) = v133
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v28)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+16)) = v135
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v28)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+24)) = v137
	v142 = F_lseg_closept_lseg(m, int32(0), v18+int32(-32), v20)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v117 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v160 = int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v162 < int32(3) {
		v232 = v160
		v240 = v159
		goto L25
	} else {
		goto L34
	}
L29:
	;
	v159 = v142
	goto L28
L30:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v142)&int64(9223372036854775807)) {
		v159 = v119
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v119)&int64(9223372036854775807)) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if base.F64_lt(v142, v119) == int32(0) {
		v159 = v119
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v169 = int32(2)
	v180 = v159
	goto L35
L35:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+32)) = v182
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+40)) = v184
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+48)) = v186
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+56)) = v188
	v191 = v169 << (uint(int32(4)) % 32)
	v192 = v28 + v191
	v193 = *(*float64)(unsafe.Add(mBase, uint32(v192)))
	*(*float64)(unsafe.Add(mBase, uint32(v20))) = v193
	v195 = *(*float64)(unsafe.Add(mBase, uint32(v192)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+8)) = v195
	v197 = v191 + v34
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v197)))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+16)) = v198
	v200 = *(*float64)(unsafe.Add(mBase, uint32(v197)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+24)) = v200
	v205 = F_lseg_closept_lseg(m, int32(0), v18+int32(-32), v20)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v232 = v160
	v240 = v220
	goto L25
L37:
	;
	if base.Ui64(base.I64_reinterpret_f64(v205)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if base.F64_lt(v205, v180) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v220 = v180
	goto L40
L40:
	;
	v222 = v169 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v222 < v223 {
		v169 = v222
		v180 = v220
		goto L35
	} else {
		goto L47
	}
L41:
	;
	v213 = v205
	goto L43
L42:
	;
	v213 = v180
	goto L43
L43:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v180)&int64(9223372036854775807)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v219 = v205
	goto L46
L45:
	;
	v219 = v213
	goto L46
L46:
	;
	v220 = v219
	goto L40
L47:
	;
	goto L36
L48:
	;
	goto L10
L49:
	;
	goto L8
}
func F_path_is_prefix_of_path(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	v3 = F_strlen(m, l0)
	mBase = m.M
	if v3 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v48 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v10 = l0
	v11 = l1
	v12 = v3
	v13 = v9
	goto L9
L6:
	;
	v36 = l1
	v40 = int32(0)
	goto L7
L7:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v48 = v40 - v41
	goto L1
L8:
	;
	v36 = v31
	v40 = v33
	goto L7
L9:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if base.B2i32(v13 != v15)|base.B2i32(v15 == int32(0)) != 0 {
		v31 = v11
		v33 = v13
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v31 = v25
	v33 = int32(0)
	goto L8
L11:
	;
	v21 = v12 - int32(1)
	if v21 == int32(0) {
		v31 = v11
		v33 = v13
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(1)
	v25 = v11 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v26 != 0 {
		v10 = v10 + v24
		v11 = v25
		v12 = v21
		v13 = v26
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v58 = int32(0)
	goto L16
L15:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+l1))))
	v58 = base.B2i32(v51 == int32(47)) | base.B2i32(v51 == int32(0))
	goto L16
L16:
	;
	return v58
}
func F_path_mul_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if int32(0) < v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return base.I64_extend_i32_u(v7)
L6:
	;
	v25 = v7 + int32(16) + v18<<(uint(int32(4))%32)
	F_point_mul_point(m, v25, v25, v14)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v29 = v18 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v29 < v30 {
		v18 = v29
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
