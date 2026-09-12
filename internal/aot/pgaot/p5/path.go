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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_ResourceOwnerEnlarge(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[644]))
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
							F_errmsg(m, int32(297715), v7)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499883), int32(1889), int32(389722))
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
				v42 = *(*int32)(unsafe.Add(mBase, _consts[551]))
				v44 = int32(4)
				v45 = v40 + v42 + v44
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
				v48 = v46 | v44
				*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v48)
				v51 = *(*int32)(unsafe.Add(mBase, _consts[10]))
				F_ResourceOwnerRemember(m, v51, v18, int32(1628448))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, _consts[551]))
					v57 = v56 + v40
					v59 = *(*int32)(unsafe.Add(mBase, _consts[10]))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v59
					v62 = v57 + int32(4)
					v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
					v65 = v63 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
					v68 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[645])) = uint8(v68)
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
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v274 int32
	_ = v274
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 float64
	_ = v289
	var v290 float64
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 float64
	_ = v321
	var v322 float64
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 float64
	_ = v340
	var v341 float64
	_ = v341
	var v349 float64
	_ = v349
	var v350 float64
	_ = v350
	var v362 int32
	_ = v362
	var v365 float64
	_ = v365
	var v366 float64
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
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
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v478 int32
	_ = v478
	var v481 float64
	_ = v481
	var v482 float64
	_ = v482
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
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
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
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 float64
	_ = v633
	var v634 float64
	_ = v634
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v681 != int32(280) {
		goto L274
	} else {
		goto L275
	}
L10:
	;
	v29 = v3
	v30 = v22
	v32 = v3
	goto L13
L11:
	;
	v674 = v3
	v677 = int32(0)
	goto L12
L12:
	;
	v678 = F_list_insert_nth(m, v677, v674, l1)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L273
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
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v674 = v660
	v677 = v663
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
	v660 = v32
	goto L17
L17:
	;
	goto L14
L18:
	;
	if v648 != 0 {
		v29 = v647 + int32(1)
		v30 = v648
		v32 = v649
		goto L13
	} else {
		goto L272
	}
L19:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v629 <= v630 {
		goto L265
	} else {
		goto L266
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
	if v167 == int32(3) {
		v626 = v106
		goto L19
	} else {
		goto L73
	}
L52:
	;
	v167 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v118 = int32(0)
	goto L57
L55:
	;
	if v157 != 0 {
		goto L70
	} else {
		goto L71
	}
L56:
	;
	v151 = int32(0)
	v157 = base.B2i32(v131 == v151)
	v159 = base.B2i32(v140 != v151) << (uint(int32(1)) % 32)
	goto L55
L57:
	;
	v121 = int32(0)
	if v21 == v121 {
		v131 = v121
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v167 = int32(3)
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
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v125 <= v118 {
		v131 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v131 = v127 + v118<<(uint(int32(2))%32)
	goto L59
L62:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v140 = v137 + v118<<(uint(int32(2))%32)
	if v131 == int32(0) {
		goto L56
	} else {
		goto L67
	}
L63:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v118 < v132 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v134 = int32(0)
	v157 = base.B2i32(v131 == v134)
	v159 = v134
	goto L55
L66:
	;
	goto L65
L67:
	;
	if v140 == int32(0) {
		goto L56
	} else {
		goto L68
	}
L68:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v147 == v148 {
		v118 = v118 + int32(1)
		goto L57
	} else {
		goto L69
	}
L69:
	;
	goto L58
L70:
	;
	v161 = v159
	goto L72
L71:
	;
	v161 = int32(1)
	goto L72
L72:
	;
	v167 = v161
	goto L51
L73:
	;
	switch v105 - int32(1) {
	case 0:
		goto L78
	case 1:
		goto L76
	default:
		goto L79
	}
L74:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	v626 = base.B2i32(base.Ui32(v622) < base.Ui32(v621))
	goto L19
L75:
	;
	v615 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v616 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_ge(v615, v616) == int32(0) {
		v626 = v106
		goto L19
	} else {
		goto L263
	}
L76:
	;
	if v167 == int32(1) {
		v626 = v106
		goto L19
	} else {
		goto L218
	}
L77:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v495 = F_list_delete_nth_cell(m, v494, v29)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L215
	}
L78:
	;
	if v167 == int32(2) {
		v626 = v106
		goto L19
	} else {
		goto L169
	}
L79:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v172 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v175 = v173
	goto L82
L81:
	;
	v175 = int32(0)
	goto L82
L82:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v176 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v179 = v177
	goto L85
L84:
	;
	v179 = int32(0)
	goto L85
L85:
	;
	if v175 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	switch v167 - int32(1) {
	case 0:
		goto L124
	case 1:
		goto L123
	default:
		goto L122
	}
L87:
	;
	v274 = base.B2i32(v179 != int32(0))
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v179 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v274 = int32(2)
	goto L86
L91:
	;
	goto L92
L92:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v196 < v197 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v199 = v196
	goto L95
L94:
	;
	v199 = v197
	goto L95
L95:
	;
	if v199 <= int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v202 = int32(1)
	goto L98
L97:
	;
	v202 = v199
	goto L98
L98:
	;
	v203 = int32(8)
	v207 = int32(0)
	v209 = v207
	v210 = v207
	goto L101
L99:
	;
	v274 = int32(3)
	goto L86
L100:
	;
	v274 = v260
	goto L86
L101:
	;
	v220 = v210 << (uint(int32(2)) % 32)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v175+v203+v220)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220+(v179+v203))))
	if v222&(v224^int32(-1)) != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	if v197 < v196 {
		goto L112
	} else {
		goto L113
	}
L103:
	;
	v246 = v210 + int32(1)
	if v246 != v202 {
		v209 = v243
		v210 = v246
		goto L101
	} else {
		goto L111
	}
L104:
	;
	v230 = int32(3)
	if v209 == int32(1) {
		v260 = v230
		goto L100
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v224&(v222^int32(-1)) == int32(0) {
		v243 = v209
		goto L103
	} else {
		goto L109
	}
L107:
	;
	if v224&(v222^int32(-1)) != 0 {
		v260 = v230
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v243 = int32(2)
	goto L103
L109:
	;
	if v209 == int32(2) {
		goto L99
	} else {
		goto L110
	}
L110:
	;
	v243 = int32(1)
	goto L103
L111:
	;
	goto L102
L112:
	;
	if v243 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if v197 <= v196 {
		v260 = v243
		goto L100
	} else {
		goto L118
	}
L115:
	;
	v253 = int32(3)
	goto L117
L116:
	;
	v253 = int32(2)
	goto L117
L117:
	;
	v274 = v253
	goto L86
L118:
	;
	if v243 == int32(2) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v259 = int32(3)
	goto L121
L120:
	;
	v259 = int32(1)
	goto L121
L121:
	;
	v260 = v259
	goto L100
L122:
	;
	switch v274 {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2:
		goto L75
	default:
		v626 = v106
		goto L19
	}
L123:
	;
	if v274&int32(-3) != 0 {
		v626 = v106
		goto L19
	} else {
		goto L128
	}
L124:
	;
	if base.Ui32(int32(1)) < base.Ui32(v274) {
		v626 = v106
		goto L19
	} else {
		goto L125
	}
L125:
	;
	v279 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_le(v279, v280) == int32(0) {
		v626 = v106
		goto L19
	} else {
		goto L126
	}
L126:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v285) <= base.Ui32(v284) {
		goto L77
	} else {
		goto L127
	}
L127:
	;
	v626 = v106
	goto L19
L128:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v290 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_ge(v289, v290) != 0 {
		goto L74
	} else {
		goto L129
	}
L129:
	;
	v626 = v106
	goto L19
L130:
	;
	v365 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v366 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_le(v365, v366) == int32(0) {
		v626 = v106
		goto L19
	} else {
		goto L167
	}
L131:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v293) < base.Ui32(v292) {
		goto L77
	} else {
		goto L132
	}
L132:
	;
	v295 = int32(0)
	if base.Ui32(v292) < base.Ui32(v293) {
		v626 = v295
		goto L19
	} else {
		goto L133
	}
L133:
	;
	v297 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_lt(v297, v298) != 0 {
		goto L77
	} else {
		goto L134
	}
L134:
	;
	if base.F64_gt(v297, v298) != 0 {
		v626 = v295
		goto L19
	} else {
		goto L135
	}
L135:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v301 != v302 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v362 == int32(1) {
		goto L77
	} else {
		goto L166
	}
L137:
	;
	if v301 < v302 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v308 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v42)+56))
	if base.F64_gt(v308, base.F64_mul(v309, float64(1.0000000001))) != 0 {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	v307 = int32(1)
	goto L142
L141:
	;
	v307 = int32(2)
	goto L142
L142:
	;
	v362 = v307
	goto L136
L143:
	;
	v362 = int32(2)
	goto L136
L144:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v314 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L146
L146:
	;
	if base.F64_lt(base.F64_mul(v308, float64(1.0000000001)), v309) != 0 {
		goto L155
	} else {
		goto L156
	}
L147:
	;
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	v322 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.F64_gt(v321, base.F64_mul(v322, float64(1.0000000001))) == int32(0) {
		goto L143
	} else {
		goto L153
	}
L148:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+24)))
	if v317 != 0 {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+25)))
	if v318 != int32(1) {
		goto L143
	} else {
		goto L152
	}
L151:
	;
	goto L143
L152:
	;
	goto L147
L153:
	;
	v362 = int32(3)
	goto L136
L154:
	;
	v362 = int32(1)
	goto L136
L155:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v333 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	goto L157
L157:
	;
	v349 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(v349, base.F64_mul(v350, float64(1.0000000001))) != 0 {
		v362 = int32(2)
		goto L136
	} else {
		goto L165
	}
L158:
	;
	v340 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v42)+48))
	if base.F64_gt(v340, base.F64_mul(v341, float64(1.0000000001))) == int32(0) {
		goto L154
	} else {
		goto L164
	}
L159:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+24)))
	if v336 != 0 {
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+25)))
	if v337 != int32(1) {
		goto L154
	} else {
		goto L163
	}
L162:
	;
	goto L154
L163:
	;
	goto L158
L164:
	;
	v362 = int32(3)
	goto L136
L165:
	;
	v362 = base.F64_gt(v350, base.F64_mul(v349, float64(1.0000000001)))
	goto L136
L166:
	;
	v626 = v295
	goto L19
L167:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v371) <= base.Ui32(v370) {
		goto L77
	} else {
		goto L168
	}
L168:
	;
	v626 = v106
	goto L19
L169:
	;
	v375 = int32(0)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v377 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	v379 = v378
	goto L172
L171:
	;
	v379 = v375
	goto L172
L172:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v380 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v382 = v381
	goto L175
L174:
	;
	v382 = v375
	goto L175
L175:
	;
	v383 = int32(1)
	if v379 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	if base.Ui32(int32(1)) < base.Ui32(v478) {
		v626 = v383
		goto L19
	} else {
		goto L212
	}
L177:
	;
	v478 = base.B2i32(v382 != int32(0))
	goto L176
L178:
	;
	goto L179
L179:
	;
	if v382 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v478 = int32(2)
	goto L176
L181:
	;
	goto L182
L182:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v400 < v401 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v403 = v400
	goto L185
L184:
	;
	v403 = v401
	goto L185
L185:
	;
	if v403 <= int32(1) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v406 = int32(1)
	goto L188
L187:
	;
	v406 = v403
	goto L188
L188:
	;
	v407 = int32(8)
	v411 = int32(0)
	v413 = v411
	v414 = v411
	goto L191
L189:
	;
	v478 = int32(3)
	goto L176
L190:
	;
	v478 = v464
	goto L176
L191:
	;
	v424 = v414 << (uint(int32(2)) % 32)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v379+v407+v424)))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424+(v382+v407))))
	if v426&(v428^int32(-1)) != 0 {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	if v401 < v400 {
		goto L202
	} else {
		goto L203
	}
L193:
	;
	v450 = v414 + int32(1)
	if v450 != v406 {
		v413 = v447
		v414 = v450
		goto L191
	} else {
		goto L201
	}
L194:
	;
	v434 = int32(3)
	if v413 == int32(1) {
		v464 = v434
		goto L190
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if v428&(v426^int32(-1)) == int32(0) {
		v447 = v413
		goto L193
	} else {
		goto L199
	}
L197:
	;
	if v428&(v426^int32(-1)) != 0 {
		v464 = v434
		goto L190
	} else {
		goto L198
	}
L198:
	;
	v447 = int32(2)
	goto L193
L199:
	;
	if v413 == int32(2) {
		goto L189
	} else {
		goto L200
	}
L200:
	;
	v447 = int32(1)
	goto L193
L201:
	;
	goto L192
L202:
	;
	if v447 == int32(1) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	if v401 <= v400 {
		v464 = v447
		goto L190
	} else {
		goto L208
	}
L205:
	;
	v457 = int32(3)
	goto L207
L206:
	;
	v457 = int32(2)
	goto L207
L207:
	;
	v478 = v457
	goto L176
L208:
	;
	if v447 == int32(2) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v463 = int32(3)
	goto L211
L210:
	;
	v463 = int32(1)
	goto L211
L211:
	;
	v464 = v463
	goto L190
L212:
	;
	v481 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v482 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_le(v481, v482) == int32(0) {
		v626 = v383
		goto L19
	} else {
		goto L213
	}
L213:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+21)))
	if base.Ui32(v486) < base.Ui32(v487) {
		v626 = v383
		goto L19
	} else {
		goto L214
	}
L214:
	;
	goto L77
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v495
	v499 = v29 - int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v500 == int32(280) {
		v647 = v499
		v648 = v495
		v649 = v32
		goto L18
	} else {
		goto L216
	}
L216:
	;
	F_pfree(m, v42)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	v647 = v499
	v648 = v495
	v649 = v32
	goto L18
L218:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v507 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v510 = v508
	goto L221
L220:
	;
	v510 = int32(0)
	goto L221
L221:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v511 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v514 = v512
	goto L224
L223:
	;
	v514 = int32(0)
	goto L224
L224:
	;
	if v510 == int32(0) {
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
		goto L261
	}
L226:
	;
	v609 = base.B2i32(v514 != int32(0))
	goto L225
L227:
	;
	goto L228
L228:
	;
	if v514 == int32(0) {
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
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v531 < v532 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v534 = v531
	goto L234
L233:
	;
	v534 = v532
	goto L234
L234:
	;
	if v534 <= int32(1) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v537 = int32(1)
	goto L237
L236:
	;
	v537 = v534
	goto L237
L237:
	;
	v538 = int32(8)
	v542 = int32(0)
	v544 = v542
	v545 = v542
	goto L240
L238:
	;
	v609 = int32(3)
	goto L225
L239:
	;
	v609 = v595
	goto L225
L240:
	;
	v555 = v545 << (uint(int32(2)) % 32)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v510+v538+v555)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+(v514+v538))))
	if v557&(v559^int32(-1)) != 0 {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	if v532 < v531 {
		goto L251
	} else {
		goto L252
	}
L242:
	;
	v581 = v545 + int32(1)
	if v581 != v537 {
		v544 = v578
		v545 = v581
		goto L240
	} else {
		goto L250
	}
L243:
	;
	v565 = int32(3)
	if v544 == int32(1) {
		v595 = v565
		goto L239
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if v559&(v557^int32(-1)) == int32(0) {
		v578 = v544
		goto L242
	} else {
		goto L248
	}
L246:
	;
	if v559&(v557^int32(-1)) != 0 {
		v595 = v565
		goto L239
	} else {
		goto L247
	}
L247:
	;
	v578 = int32(2)
	goto L242
L248:
	;
	if v544 == int32(2) {
		goto L238
	} else {
		goto L249
	}
L249:
	;
	v578 = int32(1)
	goto L242
L250:
	;
	goto L241
L251:
	;
	if v578 == int32(1) {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	goto L253
L253:
	;
	if v532 <= v531 {
		v595 = v578
		goto L239
	} else {
		goto L257
	}
L254:
	;
	v588 = int32(3)
	goto L256
L255:
	;
	v588 = int32(2)
	goto L256
L256:
	;
	v609 = v588
	goto L225
L257:
	;
	if v578 == int32(2) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v594 = int32(3)
	goto L260
L259:
	;
	v594 = int32(1)
	goto L260
L260:
	;
	v595 = v594
	goto L239
L261:
	;
	v612 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v613 = *(*float64)(unsafe.Add(mBase, uint32(v42)+32))
	if base.F64_ge(v612, v613) != 0 {
		goto L74
	} else {
		goto L262
	}
L262:
	;
	v626 = v106
	goto L19
L263:
	;
	goto L74
L264:
	;
	if v626 == int32(0) {
		goto L9
	} else {
		goto L271
	}
L265:
	;
	if v630 != v629 {
		goto L264
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	if v626 == int32(0) {
		goto L9
	} else {
		goto L270
	}
L268:
	;
	v633 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v634 = *(*float64)(unsafe.Add(mBase, uint32(v42)+56))
	if base.F64_ge(v633, v634) == int32(0) {
		goto L264
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	v647 = v29
	v648 = v30
	v649 = v29 + int32(1)
	goto L18
L271:
	;
	v647 = v29
	v648 = v30
	v649 = v32
	goto L18
L272:
	;
	v660 = v649
	goto L17
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v678
	return
L274:
	;
	F_pfree(m, l1)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	return
L277:
	;
	goto L276
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	v7 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
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
	return v186
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
	v180 = int32(1)
	v182 = v35 + v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v182 < v183 {
		v35 = v182
		goto L13
	} else {
		goto L69
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
	v62 = int32(0)
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
	v65 = v62
	goto L29
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	v65 = v64
	goto L29
L29:
	;
	if v22 == v65 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v122&int32(-3) != 0 {
		goto L22
	} else {
		goto L52
	}
L31:
	;
	v122 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v73 = int32(0)
	goto L36
L34:
	;
	if v112 != 0 {
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v106 = int32(0)
	v112 = base.B2i32(v86 == v106)
	v114 = base.B2i32(v95 != v106) << (uint(int32(1)) % 32)
	goto L34
L36:
	;
	v76 = int32(0)
	if v22 == v76 {
		v86 = v76
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v122 = int32(3)
	goto L30
L38:
	;
	if v65 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v80 <= v73 {
		v86 = int32(0)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v86 = v82 + v73<<(uint(int32(2))%32)
	goto L38
L41:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v95 = v92 + v73<<(uint(int32(2))%32)
	if v86 == int32(0) {
		goto L35
	} else {
		goto L46
	}
L42:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v73 < v87 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v89 = int32(0)
	v112 = base.B2i32(v86 == v89)
	v114 = v89
	goto L34
L45:
	;
	goto L44
L46:
	;
	if v95 == int32(0) {
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v102 == v103 {
		v73 = v73 + int32(1)
		goto L36
	} else {
		goto L48
	}
L48:
	;
	goto L37
L49:
	;
	v116 = v114
	goto L51
L50:
	;
	v116 = int32(1)
	goto L51
L51:
	;
	v122 = v116
	goto L30
L52:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v125 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v127 = v126
	goto L55
L54:
	;
	v127 = v62
	goto L55
L55:
	;
	v128 = int32(0)
	v135 = base.B2i32(l5|v127 == v128)
	if l5 == v128 {
		v174 = v135
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v174 != 0 {
		v186 = v62
		goto L21
	} else {
		goto L68
	}
L57:
	;
	goto L56
L58:
	;
	if v127 == int32(0) {
		v174 = v135
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v141 != v142 {
		v174 = int32(0)
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v144 = int32(1)
	if v141 <= v144 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v147 = v144
	goto L63
L62:
	;
	v147 = v141
	goto L63
L63:
	;
	v148 = int32(8)
	v153 = int32(0)
	goto L64
L64:
	;
	v161 = v153 << (uint(int32(2)) % 32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l5+v148+v161)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+(v127+v148))))
	v166 = base.B2i32(v163 == v165)
	if v165 != v163 {
		v174 = v166
		goto L57
	} else {
		goto L66
	}
L65:
	;
	v174 = v166
	goto L57
L66:
	;
	v169 = v153 + int32(1)
	if v169 != v147 {
		v153 = v169
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	goto L22
L69:
	;
	v186 = v180
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int64
	_ = v71
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v166 int64
	_ = v166
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v189 int64
	_ = v189
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[252]))
	if v15 == v4 {
		v253 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v295
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[253])) = v204
	v295 = int32(1)
	goto L1
L3:
	;
	v257 = F_pstrdup(m, v13)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L65
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v25 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v13
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v59
	v62 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v71 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)))
	v76 = (int64(base.Ui64(v71)>>(uint(int64(23))%64)) ^ v71) * int64(2388976653695081527)
	v80 = int64(-8645972361240307355)
	v83 = (v76 ^ int64(base.Ui64(v76)>>(uint(int64(47))%64)) ^ v80) * v80
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 == int32(0) {
		goto L21
	} else {
		goto L22
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
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v34 == int32(0) {
		v53 = v33
		v54 = v34
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v54-v53 != 0 {
		goto L7
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v33 != v34 {
		v53 = v33
		v54 = v34
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = v30
	v39 = v13
	goto L14
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v42
		v54 = v43
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v53 = v42
	v54 = v43
	goto L11
L16:
	;
	v46 = int32(1)
	if v42 == v43 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v295 = int32(1)
	goto L1
L19:
	;
	v197 = v64 & base.I32_wrap_i64(int64(base.Ui64(v189)>>(uint(int64(47))%64))^v189-int64(base.Ui64(v189)>>(uint(int64(32))%64)))
	v200 = v63 + v197*int32(24)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+22)))
	if v201 == int32(0) {
		v253 = v23
		goto L3
	} else {
		goto L49
	}
L20:
	;
	v189 = (base.I64_extend_i32_s(v175-v84) + int64(base.Ui64(v178)>>(uint(int64(23))%64)) ^ v178) * int64(2388976653695081527)
	goto L19
L21:
	;
	v175 = v84
	v178 = v83
	goto L20
L22:
	;
	goto L23
L23:
	;
	v88 = v84
	v91 = v83
	v94 = v85
	goto L24
L24:
	;
	v96 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v96 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v175 = v173
	v178 = v172
	goto L20
L26:
	;
	v166 = (v161 ^ int64(base.Ui64(v161)>>(uint(int64(23))%64))) * int64(2388976653695081527)
	v172 = (v91 ^ int64(base.Ui64(v166)>>(uint(int64(47))%64)) ^ v166) * int64(-8645972361240307355)
	v173 = v88 + v160
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v174 != 0 {
		v88 = v173
		v91 = v172
		v94 = v174
		goto L24
	} else {
		goto L48
	}
L27:
	;
	v153 = int32(1)
	v154 = int64(0)
	goto L29
L28:
	;
	v101 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+2)))
	if v101 == int64(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v160 = v153
	v161 = v154 | base.I64_extend8_s(base.I64_extend_i32_u(v94))
	goto L26
L30:
	;
	v147 = int32(2)
	v148 = int64(0)
	goto L32
L31:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+3)))
	if v105 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v153 = v147
	v154 = v148 | v96<<(uint(int64(8))%64)
	goto L29
L33:
	;
	v107 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+4)))
	if v107 == int64(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v147 = int32(3)
	v148 = v101 << (uint(int64(16)) % 64)
	goto L32
L36:
	;
	v140 = int32(4)
	v141 = int64(0)
	goto L38
L37:
	;
	v114 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+5)))
	if v114 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v142 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v88))))
	v160 = v140
	v161 = v141 | v142
	goto L26
L39:
	;
	v135 = int32(5)
	v136 = int64(0)
	goto L41
L40:
	;
	v119 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+6)))
	if v119 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v140 = v135
	v141 = v107<<(uint(int64(32))%64) | v136
	goto L38
L42:
	;
	v129 = int32(6)
	v130 = int64(0)
	goto L44
L43:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+7)))
	if v123 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v135 = v129
	v136 = v130 | v114<<(uint(int64(40))%64)
	goto L41
L45:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	v160 = int32(8)
	v161 = v125
	goto L26
L46:
	;
	goto L47
L47:
	;
	v129 = int32(7)
	v130 = v119 << (uint(int64(48)) % 64)
	goto L44
L48:
	;
	goto L25
L49:
	;
	v204 = v200
	v207 = v197
	goto L50
L50:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v23 == v212 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v253 = v23
	goto L3
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v218 == int32(0) {
		v237 = v217
		v238 = v218
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v244 = (v207 + int32(1)) & v64
	v247 = v63 + v244*int32(24)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+22)))
	if v248 != 0 {
		v204 = v247
		v207 = v244
		goto L50
	} else {
		goto L64
	}
L55:
	;
	if v238-v237 == int32(0) {
		goto L2
	} else {
		goto L63
	}
L56:
	;
	goto L55
L57:
	;
	if v217 != v218 {
		v237 = v217
		v238 = v218
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v222 = v214
	v223 = v13
	goto L59
L59:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if v227 == int32(0) {
		v237 = v226
		v238 = v227
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v237 = v226
	v238 = v227
	goto L56
L61:
	;
	v230 = int32(1)
	if v226 == v227 {
		v222 = v222 + v230
		v223 = v223 + v230
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L54
L64:
	;
	goto L51
L65:
	;
	v262 = F_SplitIdentifierString(m, v257, int32(44), v11+int32(8))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	if v262 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, _consts[87])) = v268
	goto L70
L68:
	;
	goto L69
L69:
	;
	F_pfree(m, v257)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L74
	}
L70:
	;
	v274 = F_format_elog_string(m, int32(646303), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v274
	F_pfree(m, v257)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_list_free(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	v295 = int32(0)
	goto L1
L74:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_list_free(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v287 = int32(1)
	if v15 == int32(0) {
		v295 = v287
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v290 = F_spcache_insert(m, v13, v253)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v295 = v287
	goto L1
}
func F_create_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
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
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 float64
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 float64
	_ = v326
	var v328 float64
	_ = v328
	var v330 float64
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v351 int32
	_ = v351
	v8 = l7
	v16 = F_palloc0(m, int32(88))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1434519077154)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v23
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v186
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v189)
	if v8 != 0 {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	v32 = int32(0)
	if l5 == v32 {
		v171 = v32
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v29 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v30 = F_get_baserel_parampathinfo(m, l0, l1, l5)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v186 = v30
	goto L3
L9:
	;
	v186 = v171
	goto L3
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v35 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v171 = v153
	goto L9
L12:
	;
	v130 = F_palloc0(m, int32(24))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L31
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v52 = int32(0)
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v52<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v61 = int32(0)
	v68 = base.B2i32(v60|l5 == v61)
	if v60 == v61 {
		v107 = v68
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L12
L17:
	;
	if v107 != 0 {
		v153 = v59
		goto L11
	} else {
		goto L29
	}
L18:
	;
	goto L17
L19:
	;
	if l5 == int32(0) {
		v107 = v68
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v74 != v75 {
		v107 = int32(0)
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v77 = int32(1)
	if v74 <= v77 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v80 = v77
	goto L24
L23:
	;
	v80 = v74
	goto L24
L24:
	;
	v81 = int32(8)
	v86 = int32(0)
	goto L25
L25:
	;
	v94 = v86 << (uint(int32(2)) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v60+v81+v94)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+(l5+v81))))
	v99 = base.B2i32(v96 == v98)
	if v98 != v96 {
		v107 = v99
		goto L18
	} else {
		goto L27
	}
L26:
	;
	v107 = v99
	goto L18
L27:
	;
	v102 = v86 + int32(1)
	if v102 != v80 {
		v86 = v102
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v112 = v52 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v112 < v113 {
		v52 = v112
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	v132 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v130)+8)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = int32(278)
	*(*int64)(unsafe.Add(mBase, uint32(v130)+16)) = v132
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v140 = F_lappend(m, v139, v130)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v140
	v153 = v130
	goto L11
L33:
	;
	F_list_sort(m, l2, int32(879))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	F_list_sort(m, l3, int32(880))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v201 = v199
	goto L40
L39:
	;
	v201 = int32(0)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v201
	v203 = F_list_concat(m, l2, l3)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v203
	v206 = float64(-1)
	if l0 == int32(0) {
		v264 = v206
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+80)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	if v266 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v211 = int32(0)
	v218 = base.B2i32(v209|v210 == v211)
	if v209 == v211 {
		v257 = v218
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v257 == int32(0) {
		v264 = v206
		goto L42
	} else {
		goto L56
	}
L45:
	;
	goto L44
L46:
	;
	if v210 == int32(0) {
		v257 = v218
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v224 != v225 {
		v257 = int32(0)
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v227 = int32(1)
	if v224 <= v227 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v230 = v227
	goto L51
L50:
	;
	v230 = v224
	goto L51
L51:
	;
	v231 = int32(8)
	v236 = int32(0)
	goto L52
L52:
	;
	v244 = v236 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v209+v231+v244)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v210+v231))))
	v249 = base.B2i32(v246 == v248)
	if v248 != v246 {
		v257 = v249
		goto L45
	} else {
		goto L54
	}
L53:
	;
	v257 = v249
	goto L45
L54:
	;
	v252 = v236 + int32(1)
	if v252 != v230 {
		v236 = v252
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v263 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v264 = v263
	goto L42
L57:
	;
	if base.F64_ge(l8, float64(0)) != 0 {
		goto L76
	} else {
		goto L77
	}
L58:
	;
	F_cost_append(m, v16)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L75
	}
L59:
	;
	v269 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v269 < v270 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v274 = v269
	goto L63
L61:
	;
	goto L62
L62:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v319 != int32(1) {
		goto L58
	} else {
		goto L69
	}
L63:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)))
	if v287 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L62
L65:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v274<<(uint(int32(2))%32))))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+21)))
	v297 = v295
	goto L67
L66:
	;
	v297 = int32(0)
	goto L67
L67:
	;
	v298 = int32(1)
	v299 = v297 & v298
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v299)
	v302 = v274 + v298
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v302 < v303 {
		v274 = v302
		goto L63
	} else {
		goto L68
	}
L68:
	;
	goto L64
L69:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+20)))
	if v8 == v324 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v323)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v334
	goto L57
L71:
	;
	v326 = *(*float64)(unsafe.Add(mBase, uint32(v323)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v326
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v323)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = v328
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v323)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = v330
	goto L70
L72:
	;
	goto L73
L73:
	;
	F_cost_append(m, v16)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	goto L57
L76:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = l8
	goto L78
L77:
	;
	goto L78
L78:
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
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
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
	var v93 int32
	_ = v93
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
	var v114 int32
	_ = v114
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
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
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
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(297)
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
	v58 = v55 + v32<<(uint(int32(2))%32)
	if v58 == int32(0) {
		v68 = v52
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v63 == v64 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(369)
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
	v151 = m.ExcPending
	if v151 != 0 {
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
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v91 = v90
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
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
	v108 = *(*float64)(unsafe.Add(mBase, _consts[397]))
	v110 = *(*float64)(unsafe.Add(mBase, _consts[398]))
	v112 = *(*float64)(unsafe.Add(mBase, _consts[387]))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _consts[399])))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v93 + (v114 ^ int32(1))
	v119 = base.F64_add(v112, v112)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v123 = base.F64_add(base.F64_convert_i32_s(v120), float64(1))
	v125 = F_log(m, v123)
	mBase = m.M
	v127 = base.F64_div(v125, float64(0.693147180559945))
	v129 = float64(0)
	v131 = base.F64_add(v110, base.F64_add(base.F64_mul(base.F64_mul(v119, v123), v127), v129))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = base.F64_add(base.F64_add(v94, v95), v131)
	*(*float64)(unsafe.Add(mBase, uint32(v17)+56)) = base.F64_add(base.F64_add(v97, v95), base.F64_add(v131, base.F64_add(base.F64_mul(base.F64_mul(v105, v108), float64(1.05)), base.F64_add(base.F64_mul(v112, v105), base.F64_add(base.F64_mul(base.F64_mul(v105, v119), v127), v129)))))
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
	F_errmsg_internal(m, int32(441574), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(499586), int32(2118), int32(321770))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	var v16 int32
	_ = v16
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
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1554778161455)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
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
		v41 = *(*int32)(unsafe.Add(mBase, _consts[20]))
		v43 = m.G0
		v44 = int32(16)
		v45 = v43 - v44
		m.G0 = v45
		F_cost_tuplesort(m, v45+int32(8), v45, v36, v38, float64(0), v41, l3)
		mBase = m.M
		v50 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v36
		v53 = int32(*(*uint8)(unsafe.Add(mBase, _consts[400])))
		v54 = base.F64_add(v35, v50)
		*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = v54
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v34 + (v53 ^ v27)
		v60 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
		*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(v54, v60)
		m.G0 = v45 + v44
		return v8
	}
}
func F_path_area(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v95 float64
	_ = v95
	var v105 float64
	_ = v105
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	v2 = float64(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L45
	}
L2:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L44
	}
L3:
	;
	v115 = base.F64_abs(v105)
	v117 = base.F64_mul(v115, float64(0.5))
	v119 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v117), v119)&base.F64_ne(v115, v119) != 0 {
		goto L2
	} else {
		goto L41
	}
L4:
	;
	v26 = v13 + int32(16)
	v28 = int32(0)
	v29 = v2
	goto L11
L5:
	;
	return int32(0)
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v18 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
	return int32(0)
L10:
	;
	v105 = v2
	goto L3
L11:
	;
	v41 = v26 + v28<<(uint(int32(4))%32)
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v41)))
	v44 = v28 + int32(1)
	if v44 != v18 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v105 = v95
	goto L3
L13:
	;
	if base.F64_ne(v52, float64(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v47 = v44
	goto L16
L15:
	;
	v47 = int32(0)
	goto L16
L16:
	;
	v50 = v26 + v47<<(uint(int32(4))%32)
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v50)+8))
	v52 = base.F64_mul(v42, v51)
	v53 = base.F64_abs(v52)
	if base.F64_ne(v53, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if base.F64_eq(base.F64_abs(v42), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	v68 = base.F64_add(v29, v52)
	v69 = base.F64_abs(v68)
	if base.F64_ne(v69, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if base.F64_eq(v42, float64(0)) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if base.F64_ne(v51, float64(0)) != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v41)+8))
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v50)))
	v79 = base.F64_mul(v77, v78)
	v80 = base.F64_abs(v79)
	if base.F64_ne(v80, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	if base.F64_eq(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.F64_ne(v53, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	if base.F64_ne(v79, float64(0)) != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if base.F64_eq(base.F64_abs(v77), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if base.F64_ne(base.F64_abs(v78), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v95 = base.F64_sub(v68, v79)
	if base.F64_ne(base.F64_abs(v95), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if base.F64_eq(v77, float64(0)) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.F64_ne(v78, float64(0)) != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v44 != v18 {
		v28 = v44
		v29 = v95
		goto L11
	} else {
		goto L40
	}
L37:
	;
	if base.F64_eq(v69, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if base.F64_ne(v80, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	goto L12
L41:
	;
	v124 = float64(0)
	if base.F64_eq(v117, v124)&base.F64_ne(v105, v124) != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v129 = F_Float8GetDatum(m, v117)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	return v129
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_path_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
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
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v167 float64
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v187 float64
	_ = v187
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v193 float64
	_ = v193
	var v195 float64
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v208 float64
	_ = v208
	var v213 float64
	_ = v213
	var v214 int32
	_ = v214
	var v221 float64
	_ = v221
	var v227 float64
	_ = v227
	var v228 float64
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v248 float64
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 float64
	_ = v265
	var v268 int32
	_ = v268
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v29 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v19 - int32(-64)
	return v309
L5:
	;
	v291 = F_Float8GetDatum(m, v265)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L50
	}
L6:
	;
	v32 = int32(16)
	v33 = v27 + v32
	v35 = v22 + v32
	v43 = v29
	v44 = v2
	v48 = v2
	v54 = float64(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v288 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v288)
	v309 = int32(0)
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
	if v255&int32(1) != 0 {
		goto L5
	} else {
		goto L49
	}
L11:
	;
	v268 = v48 + int32(1)
	if v268 < v254 {
		v43 = v254
		v44 = v255
		v48 = v268
		v54 = v265
		goto L9
	} else {
		goto L48
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v58 == int32(0) {
		v254 = v43
		v255 = v44
		v265 = v54
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
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v62 <= int32(0) {
		v254 = v43
		v255 = v44
		v265 = v54
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
	v67 = v35 + v48<<(uint(v65)%32)
	v72 = v61<<(uint(v65)%32) + v35 - int32(16)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v73 == int32(0) {
		v123 = v44
		v124 = v54
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v126 < int32(2) {
		v238 = v123
		v248 = v124
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v76
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+40)) = v78
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v80
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = v82
	v88 = v62<<(uint(int32(4))%32) + v33 - int32(16)
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	*(*float64)(unsafe.Add(mBase, uint32(v19))) = v89
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v88)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v91
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v93
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v95
	v100 = F_lseg_closept_lseg(m, int32(0), v17+int32(-32), v19)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v44&int32(1) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v123 = int32(1)
	v124 = v100
	goto L17
L21:
	;
	v106 = int32(1)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v100)&int64(9223372036854775807)) {
		v123 = v106
		v124 = v54
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
	if base.F64_gt(v54, v100) == int32(0) {
		v123 = v106
		v124 = v54
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v254 = v250
	v255 = v238
	v265 = v248
	goto L11
L26:
	;
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v129
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+40)) = v131
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v133
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = v135
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v19))) = v137
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v139
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v27+int32(32))))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v141
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v27+int32(40))))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v143
	v148 = F_lseg_closept_lseg(m, int32(0), v17+int32(-32), v19)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v123&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v170 < int32(3) {
		v238 = int32(1)
		v248 = v167
		goto L25
	} else {
		goto L34
	}
L29:
	;
	v167 = v148
	goto L28
L30:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v148)&int64(9223372036854775807)) {
		v167 = v124
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v124)&int64(9223372036854775807)) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if base.F64_gt(v124, v148) == int32(0) {
		v167 = v124
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v176 = int32(2)
	v187 = v167
	goto L35
L35:
	;
	v189 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v189
	v191 = *(*float64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+40)) = v191
	v193 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v193
	v195 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = v195
	v199 = v33 + v176<<(uint(int32(4))%32)
	v201 = v199 - int32(16)
	v202 = *(*float64)(unsafe.Add(mBase, uint32(v201)))
	*(*float64)(unsafe.Add(mBase, uint32(v19))) = v202
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v201)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v204
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v199)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v206
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v199)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v208
	v213 = F_lseg_closept_lseg(m, int32(0), v17+int32(-32), v19)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v238 = v229
	v248 = v228
	goto L25
L37:
	;
	if base.Ui64(base.I64_reinterpret_f64(v213)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if base.F64_gt(v187, v213) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v228 = v187
	goto L40
L40:
	;
	v229 = int32(1)
	v231 = v176 + v229
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v231 < v232 {
		v176 = v231
		v187 = v228
		goto L35
	} else {
		goto L47
	}
L41:
	;
	v221 = v213
	goto L43
L42:
	;
	v221 = v187
	goto L43
L43:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v187)&int64(9223372036854775807)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v227 = v213
	goto L46
L45:
	;
	v227 = v221
	goto L46
L46:
	;
	v228 = v227
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
L50:
	;
	v309 = v291
	goto L4
}
func F_path_is_prefix_of_path(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v16 int32
	_ = v16
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
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	v3 = int32(0)
	v4 = F_strlen(m, l0)
	mBase = m.M
	if v4 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v48 != 0 {
		goto L15
	} else {
		goto L16
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
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v11 = l0
	v12 = l1
	v13 = v4
	v14 = v10
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != v16 {
		v31 = v12
		v33 = v14
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
	if v16 == int32(0) {
		v31 = v12
		v33 = v14
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v21 = v13 - int32(1)
	if v21 == int32(0) {
		v31 = v12
		v33 = v14
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v24 = int32(1)
	v25 = v12 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v26 != 0 {
		v11 = v11 + v24
		v12 = v25
		v13 = v21
		v14 = v26
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v57 = v3
	goto L17
L16:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+l1))))
	v57 = base.B2i32(v50 == int32(47)) | base.B2i32(v50 == int32(0))
	goto L17
L17:
	;
	return v57
}
func F_path_mul_pt(m *base.Module, l0 int32) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v7
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
