package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateActiveSnapshotCommandId(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	v8 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[4]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
		if v15 != 0 {
			v17 = int32(1)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
			v17 = v16
		}
		if v17&int32(1) == int32(0) {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[26]))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v8
			return
		} else {
			if v8 == v6 {
				v37 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v8
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(259975), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errfinish(m, int32(495541), int32(762), int32(465979))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
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
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
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
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v750 int32
	_ = v750
	var v762 int32
	_ = v762
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	if base.Ui32(l7) <= base.Ui32(int32(41)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v762 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v742))) = uint8(v762)
	m.G0 = v25 + int32(48)
	return v750 - l0
L2:
	;
	v52 = l1
	v53 = l2
	v61 = l0
	goto L12
L3:
	;
	if int32(0) < l1 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v742 = l2
	v750 = l0
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = l7
	F_errmsg(m, int32(482880), v25+int32(32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(492911), int32(522), int32(314416))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v73 == int32(0) {
		v236 = v52
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v742 = v724
	v750 = v736
	goto L1
L14:
	;
	v737 = v728 - v725
	if int32(0) < v737 {
		v52 = v737
		v53 = v724
		v61 = v736
		goto L12
	} else {
		goto L199
	}
L15:
	;
	if l3 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L16:
	;
	v251 = int32(0)
	switch v234 - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L87
	case 2:
		goto L88
	case 3:
		goto L89
	default:
		v300 = v251
		goto L83
	}
L17:
	;
	if l8 != 0 {
		v742 = v53
		v750 = v61
		goto L1
	} else {
		goto L79
	}
L18:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	if int32(0) <= v76 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v52 < v100 {
		v236 = v52
		goto L17
	} else {
		goto L32
	}
L20:
	;
	v100 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v81 = v76 & int32(255)
	if v81&int32(224) == int32(192) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v100 = int32(2)
	goto L19
L24:
	;
	goto L25
L25:
	;
	if v81&int32(240) == int32(224) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = int32(3)
	goto L19
L27:
	;
	goto L28
L28:
	;
	if v81&int32(248) == int32(240) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v98 = int32(4)
	goto L31
L30:
	;
	v98 = int32(1)
	goto L31
L31:
	;
	v100 = v98
	goto L19
L32:
	;
	v102 = int32(0)
	switch v100 - int32(1) {
	case 0:
		goto L37
	case 1:
		goto L38
	case 2:
		goto L39
	case 3:
		goto L40
	default:
		v151 = v102
		goto L34
	}
L33:
	;
	if v151 == int32(0) {
		v236 = v52
		goto L17
	} else {
		goto L54
	}
L34:
	;
	goto L33
L35:
	;
	v151 = base.B2i32(base.Ui32(v143&int32(255)) < base.Ui32(int32(245)))
	goto L34
L36:
	;
	if base.I32_extend8_s(v138) < int32(-62) {
		v151 = v102
		goto L34
	} else {
		goto L53
	}
L37:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v138 = v137
	goto L36
L38:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	switch v112 - int32(224) {
	case 0:
		goto L47
	default:
		goto L43
	case 13:
		goto L46
	case 16:
		goto L45
	case 20:
		goto L44
	}
L39:
	;
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+2)))
	if int32(-65) < v108 {
		v151 = v102
		goto L34
	} else {
		goto L42
	}
L40:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+3)))
	if int32(-65) < v105 {
		v151 = v102
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L38
L43:
	;
	if v111 <= int32(-65) {
		v138 = v112
		goto L36
	} else {
		goto L52
	}
L44:
	;
	if int32(-113) < v111 {
		v151 = v102
		goto L34
	} else {
		goto L51
	}
L45:
	;
	if base.Ui32((v111-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v151 = v102
		goto L34
	} else {
		goto L50
	}
L46:
	;
	if int32(-97) < v111 {
		v151 = v102
		goto L34
	} else {
		goto L49
	}
L47:
	;
	v115 = int32(224)
	if base.Ui32(v115) <= base.Ui32((v111-int32(-64))&int32(255)) {
		v143 = v115
		goto L35
	} else {
		goto L48
	}
L48:
	;
	v151 = v102
	goto L34
L49:
	;
	v143 = int32(237)
	goto L35
L50:
	;
	v143 = int32(240)
	goto L35
L51:
	;
	v143 = int32(244)
	goto L35
L52:
	;
	v151 = v102
	goto L34
L53:
	;
	v143 = v138
	goto L35
L54:
	;
	v154 = int32(0)
	v155 = int32(1)
	switch v100 - v155 {
	case 0:
		goto L59
	case 1:
		v187 = v155
		v188 = v154
		v189 = v61
		v190 = v154
		goto L55
	case 2:
		goto L56
	case 3:
		goto L58
	default:
		goto L57
	}
L55:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v61))))
	v196 = v188 & int32(255)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v204 = v192 | (v190<<(uint(int32(24))%32) | v196<<(uint(int32(16))%32) | v200<<(uint(int32(8))%32))
	v205 = v61 + v100
	if l4 == int32(0) {
		goto L15
	} else {
		goto L63
	}
L56:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v187 = int32(2)
	v188 = v185
	v189 = v61 + int32(1)
	v190 = v154
	goto L55
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v187 = int32(3)
	v188 = v167
	v189 = v61 + int32(2)
	v190 = v168
	goto L55
L59:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v159)
	v161 = int32(1)
	v724 = v53 + v161
	v725 = v155
	v728 = v52
	v736 = v61 + v161
	goto L14
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v100
	F_errmsg_internal(m, int32(474721), v25)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(492911), int32(570), int32(314416))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	if base.Ui32(v52) <= base.Ui32(v100) {
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v209 = v52 - v100
	v210 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205))))
	if int32(0) <= v210 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v234 <= v209 {
		goto L16
	} else {
		goto L78
	}
L66:
	;
	v234 = int32(1)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v215 = v210 & int32(255)
	if v215&int32(224) == int32(192) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v234 = int32(2)
	goto L65
L70:
	;
	goto L71
L71:
	;
	if v215&int32(240) == int32(224) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v234 = int32(3)
	goto L65
L73:
	;
	goto L74
L74:
	;
	if v215&int32(248) == int32(240) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v232 = int32(4)
	goto L77
L76:
	;
	v232 = int32(1)
	goto L77
L77:
	;
	v234 = v232
	goto L65
L78:
	;
	v236 = v209
	goto L17
L79:
	;
	if v236 <= int32(0) {
		v742 = v53
		v750 = v61
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_report_invalid_encoding(m, int32(6), v61, v236)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	if v300 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L83:
	;
	goto L82
L84:
	;
	v300 = base.B2i32(base.Ui32(v292&int32(255)) < base.Ui32(int32(245)))
	goto L83
L85:
	;
	if base.I32_extend8_s(v287) < int32(-62) {
		v300 = v251
		goto L83
	} else {
		goto L102
	}
L86:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v287 = v286
	goto L85
L87:
	;
	v260 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205)+1)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	switch v261 - int32(224) {
	case 0:
		goto L96
	default:
		goto L92
	case 13:
		goto L95
	case 16:
		goto L94
	case 20:
		goto L93
	}
L88:
	;
	v257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205)+2)))
	if int32(-65) < v257 {
		v300 = v251
		goto L83
	} else {
		goto L91
	}
L89:
	;
	v254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205)+3)))
	if int32(-65) < v254 {
		v300 = v251
		goto L83
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L87
L92:
	;
	if v260 <= int32(-65) {
		v287 = v261
		goto L85
	} else {
		goto L101
	}
L93:
	;
	if int32(-113) < v260 {
		v300 = v251
		goto L83
	} else {
		goto L100
	}
L94:
	;
	if base.Ui32((v260-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v300 = v251
		goto L83
	} else {
		goto L99
	}
L95:
	;
	if int32(-97) < v260 {
		v300 = v251
		goto L83
	} else {
		goto L98
	}
L96:
	;
	v264 = int32(224)
	if base.Ui32(v264) <= base.Ui32((v260-int32(-64))&int32(255)) {
		v292 = v264
		goto L84
	} else {
		goto L97
	}
L97:
	;
	v300 = v251
	goto L83
L98:
	;
	v292 = int32(237)
	goto L84
L99:
	;
	v292 = int32(240)
	goto L84
L100:
	;
	v292 = int32(244)
	goto L84
L101:
	;
	v300 = v251
	goto L83
L102:
	;
	v292 = v287
	goto L84
L103:
	;
	if l8 != 0 {
		v742 = v53
		v750 = v61
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v234 < int32(2) {
		goto L15
	} else {
		goto L108
	}
L106:
	;
	F_report_invalid_encoding(m, int32(6), v205, v209)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	switch v234 - int32(2) {
	case 0:
		goto L110
	case 1:
		goto L113
	case 2:
		goto L112
	default:
		goto L111
	}
L109:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v354 | v353
	v362 = F_bsearch(m, v25+int32(40), l4, l5, int32(12), int32(1650))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L7
	} else {
		goto L117
	}
L110:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v352 = v205 + int32(1)
	v353 = v349 << (uint(int32(8)) % 32)
	goto L109
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L7
	} else {
		goto L114
	}
L112:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
	v352 = v205 + int32(3)
	v353 = v321<<(uint(int32(16))%32) | v324<<(uint(int32(24))%32) | v328<<(uint(int32(8))%32)
	goto L109
L113:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v352 = v205 + int32(2)
	v353 = v312<<(uint(int32(8))%32) | v315<<(uint(int32(16))%32)
	goto L109
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v234
	F_errmsg_internal(m, int32(474721), v25+int32(16))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(492911), int32(627), int32(314416))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	if v362 == int32(0) {
		goto L15
	} else {
		goto L118
	}
L118:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	if base.Ui32(int32(16777216)) <= base.Ui32(v366) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v370 = int32(base.Ui32(v366) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v370)
	v374 = v53 + int32(1)
	goto L121
L120:
	;
	v374 = v53
	goto L121
L121:
	;
	if v366&int32(16711680) != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v378 = int32(base.Ui32(v366) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v374))) = uint8(v378)
	v382 = v374 + int32(1)
	goto L124
L123:
	;
	v382 = v374
	goto L124
L124:
	;
	if v366&int32(65280) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v386 = int32(base.Ui32(v366) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v386)
	v390 = v382 + int32(1)
	goto L127
L126:
	;
	v390 = v382
	goto L127
L127:
	;
	if v366&int32(255) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v366)
	v396 = v390 + int32(1)
	goto L130
L129:
	;
	v396 = v390
	goto L130
L130:
	;
	v724 = v396
	v725 = v234
	v728 = v209
	v736 = v234 + v205
	goto L14
L131:
	;
	v724 = v722
	v725 = v100
	v728 = v52
	v736 = v205
	goto L14
L132:
	;
	if l6 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L133:
	;
	v404 = int32(0)
	switch v100 - int32(1) {
	case 0:
		goto L136
	case 1:
		goto L137
	case 2:
		goto L138
	case 3:
		goto L139
	default:
		v633 = v404
		goto L135
	}
L134:
	;
	if v646 == int32(0) {
		goto L132
	} else {
		goto L172
	}
L135:
	;
	v646 = v633
	goto L134
L136:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(v192) < base.Ui32(v606) {
		v633 = v404
		goto L135
	} else {
		goto L167
	}
L137:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if base.Ui32(v200) < base.Ui32(v561) {
		v633 = v404
		goto L135
	} else {
		goto L160
	}
L138:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)))
	if base.Ui32(v196) < base.Ui32(v496) {
		v633 = v404
		goto L135
	} else {
		goto L151
	}
L139:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	if base.Ui32(v190) < base.Ui32(v411) {
		v633 = v404
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+41)))
	if base.Ui32(v413) < base.Ui32(v190) {
		v633 = v404
		goto L135
	} else {
		goto L141
	}
L141:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+42)))
	if base.Ui32(v196) < base.Ui32(v415) {
		v633 = v404
		goto L135
	} else {
		goto L142
	}
L142:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+43)))
	if base.Ui32(v417) < base.Ui32(v196) {
		v633 = v404
		goto L135
	} else {
		goto L143
	}
L143:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+44)))
	if base.Ui32(v200) < base.Ui32(v419) {
		v633 = v404
		goto L135
	} else {
		goto L144
	}
L144:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if base.Ui32(v421) < base.Ui32(v200) {
		v633 = v404
		goto L135
	} else {
		goto L145
	}
L145:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if base.Ui32(v192) < base.Ui32(v423) {
		v633 = v404
		goto L135
	} else {
		goto L146
	}
L146:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+47)))
	if base.Ui32(v425) < base.Ui32(v192) {
		v633 = v404
		goto L135
	} else {
		goto L147
	}
L147:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v428 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v430 = int32(2)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v428+(v190-v411)<<(uint(v430)%32)+v427<<(uint(v430)%32))))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v428+(v196-v415)<<(uint(v430)%32)+v448<<(uint(v430)%32))))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v428+(v200-v419)<<(uint(v430)%32)+v452<<(uint(v430)%32))))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v428+(v192-v423)<<(uint(v430)%32)+v456<<(uint(v430)%32))))
	v646 = v460
	goto L134
L149:
	;
	goto L150
L150:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v463 = int32(1)
	v483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461+(v190-v411)<<(uint(v463)%32)+v427&int32(65535)<<(uint(v463)%32)))))
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461+(v196-v415)<<(uint(v463)%32)+v483<<(uint(v463)%32)))))
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461+(v200-v419)<<(uint(v463)%32)+v487<<(uint(v463)%32)))))
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461+(v192-v423)<<(uint(v463)%32)+v491<<(uint(v463)%32)))))
	v646 = v495
	goto L134
L151:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)))
	if base.Ui32(v498) < base.Ui32(v196) {
		v633 = v404
		goto L135
	} else {
		goto L152
	}
L152:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+30)))
	if base.Ui32(v200) < base.Ui32(v500) {
		v633 = v404
		goto L135
	} else {
		goto L153
	}
L153:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+31)))
	if base.Ui32(v502) < base.Ui32(v200) {
		v633 = v404
		goto L135
	} else {
		goto L154
	}
L154:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if base.Ui32(v192) < base.Ui32(v504) {
		v633 = v404
		goto L135
	} else {
		goto L155
	}
L155:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+33)))
	if base.Ui32(v506) < base.Ui32(v192) {
		v633 = v404
		goto L135
	} else {
		goto L156
	}
L156:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v509 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v511 = int32(2)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v509+(v196-v496)<<(uint(v511)%32)+v508<<(uint(v511)%32))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v509+(v200-v500)<<(uint(v511)%32)+v525<<(uint(v511)%32))))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v509+(v192-v504)<<(uint(v511)%32)+v529<<(uint(v511)%32))))
	v646 = v533
	goto L134
L158:
	;
	goto L159
L159:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v536 = int32(1)
	v552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v534+(v196-v496)<<(uint(v536)%32)+v508&int32(65535)<<(uint(v536)%32)))))
	v556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v534+(v200-v500)<<(uint(v536)%32)+v552<<(uint(v536)%32)))))
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v534+(v192-v504)<<(uint(v536)%32)+v556<<(uint(v536)%32)))))
	v646 = v560
	goto L134
L160:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if base.Ui32(v563) < base.Ui32(v200) {
		v633 = v404
		goto L135
	} else {
		goto L161
	}
L161:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)))
	if base.Ui32(v192) < base.Ui32(v565) {
		v633 = v404
		goto L135
	} else {
		goto L162
	}
L162:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)))
	if base.Ui32(v567) < base.Ui32(v192) {
		v633 = v404
		goto L135
	} else {
		goto L163
	}
L163:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v570 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v572 = int32(2)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v570+(v200-v561)<<(uint(v572)%32)+v569<<(uint(v572)%32))))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v570+(v192-v565)<<(uint(v572)%32)+v582<<(uint(v572)%32))))
	v646 = v586
	goto L134
L165:
	;
	goto L166
L166:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v589 = int32(1)
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587+(v200-v561)<<(uint(v589)%32)+v569&int32(65535)<<(uint(v589)%32)))))
	v605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587+(v192-v565)<<(uint(v589)%32)+v601<<(uint(v589)%32)))))
	v646 = v605
	goto L134
L167:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if base.Ui32(v608) < base.Ui32(v192) {
		v633 = v404
		goto L135
	} else {
		goto L168
	}
L168:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v610 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v612 = int32(2)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v610+(v192-v606)<<(uint(v612)%32)+v615<<(uint(v612)%32))))
	v646 = v619
	goto L134
L170:
	;
	goto L171
L171:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v622 = int32(1)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v620+(v192-v606)<<(uint(v622)%32)+v625<<(uint(v622)%32)))))
	v633 = v629
	goto L135
L172:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v646) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v652 = int32(base.Ui32(v646) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v652)
	v656 = v53 + int32(1)
	goto L175
L174:
	;
	v656 = v53
	goto L175
L175:
	;
	if v646&int32(16711680) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v660 = int32(base.Ui32(v646) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v656))) = uint8(v660)
	v664 = v656 + int32(1)
	goto L178
L177:
	;
	v664 = v656
	goto L178
L178:
	;
	if v646&int32(65280) != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v668 = int32(base.Ui32(v646) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v664))) = uint8(v668)
	v672 = v664 + int32(1)
	goto L181
L180:
	;
	v672 = v664
	goto L181
L181:
	;
	if v646&int32(255) == int32(0) {
		v722 = v672
		goto L131
	} else {
		goto L182
	}
L182:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v672))) = uint8(v646)
	v722 = v672 + int32(1)
	goto L131
L183:
	;
	if l8 != 0 {
		v742 = v53
		v750 = v61
		goto L1
	} else {
		goto L197
	}
L184:
	;
	v683 = m.T0[l6].(func(*base.Module, int32) int32)(m, v204)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	if v683 == int32(0) {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v683) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v690 = int32(base.Ui32(v683) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v690)
	v694 = v53 + int32(1)
	goto L189
L188:
	;
	v694 = v53
	goto L189
L189:
	;
	if v683&int32(16711680) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v698 = int32(base.Ui32(v683) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v694))) = uint8(v698)
	v702 = v694 + int32(1)
	goto L192
L191:
	;
	v702 = v694
	goto L192
L192:
	;
	if v683&int32(65280) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v706 = int32(base.Ui32(v683) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v702))) = uint8(v706)
	v710 = v702 + int32(1)
	goto L195
L194:
	;
	v710 = v702
	goto L195
L195:
	;
	if v683&int32(255) == int32(0) {
		v722 = v710
		goto L131
	} else {
		goto L196
	}
L196:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v683)
	v722 = v710 + int32(1)
	goto L131
L197:
	;
	F_report_untranslatable_char(m, int32(6), l7, v61, v52)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L7
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	goto L13
}
func F___uflow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(-1)
	v9 = F___toread(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v22 = v8
			m.G0 = v6 + int32(16)
			return v22
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v17 = m.T0[v16].(func(*base.Module, int32, int32, int32) int32)(m, l0, v6+int32(15), int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != int32(1) {
					v22 = v8
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
					v22 = v21
				}
				m.G0 = v6 + int32(16)
				return v22
			}
		}
	}
}
func F_uniq(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		if v7 != 0 {
			v8 = F_array_contains_nulls(m, v3)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				if v8 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(152677), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496196), int32(254), int32(231012))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
					v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
					v13 = F_ArrayGetNItems(m, v10, v3+int32(16))
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return int32(0)
					} else {
						if int32(2) <= v13 {
							v17 = F__int_unique(m, v3)
							mBase = m.M
							v18 = m.ExcPending
							if v18 != 0 {
								return int32(0)
							} else {
								v19 = v17
								return v19
							}
						} else {
							v19 = v3
							return v19
						}
					}
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			v13 = F_ArrayGetNItems(m, v10, v3+int32(16))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if int32(2) <= v13 {
					v17 = F__int_unique(m, v3)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = v17
						return v19
					}
				} else {
					v19 = v3
					return v19
				}
			}
		}
	}
}
func F_update_frameheadpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
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
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v320 int64
	_ = v320
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v341 int64
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
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
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int64
	_ = v392
	var v399 int32
	_ = v399
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v437 int32
	_ = v437
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+380)))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L22
	} else {
		goto L138
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L22
	} else {
		goto L135
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L22
	} else {
		goto L132
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = int32(4515712)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v30
	if v24&int32(32) != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v19 + int32(16)
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v27
	goto L6
L8:
	;
	v437 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+380)) = uint8(v437)
	goto L7
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v24&int32(512) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v24&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v24&int32(10240) == int32(0) {
		goto L7
	} else {
		goto L47
	}
L15:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v40
	goto L8
L16:
	;
	goto L17
L17:
	;
	if v24&int32(10) == int32(0) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+96))
	if v46 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(0)
	goto L8
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_tuplestore_select_read_pointer(m, v51, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v55 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	goto L32
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v58 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	if v59&int32(2) == int32(0) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v65 = int32(1)
	v67 = F_tuplestore_gettupleslot(m, v64, v65, v65, v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L22
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v67 == int32(0) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v88 == int32(0) {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	goto L8
L34:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
	if v91&int32(2) != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+96))
	if v95 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v98)+12)) = v88
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v102 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	F_MemoryContextReset(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v108 = int32(4515712)
	v109 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v111
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	v116 = m.T0[v115].(func(*base.Module, int32, int32, int32) int32)(m, v102, v98, v19+int32(14))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L22
	} else {
		goto L41
	}
L40:
	;
	goto L8
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v109
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	F_MemoryContextReset(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	if v116 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v125 = v123 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v125
	F_spool_tuples(m, l0, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v130 = int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v133 = F_tuplestore_gettupleslot(m, v129, v130, v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	if v133 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	goto L33
L47:
	;
	if v24&int32(4) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
	if v24&int32(2048) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v24&int32(2) != 0 {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	v148 = int64(0) - v144
	goto L53
L52:
	;
	v148 = v144
	goto L53
L53:
	;
	v149 = v141 + v148
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v149
	if int64(0) <= v149 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v149 <= v141+int64(1) {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	v163 = int64(0)
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v163
	goto L8
L57:
	;
	F_spool_tuples(m, l0, v149-int64(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v161 <= v160 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	v163 = v160
	goto L56
L60:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v25)+100))
	v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v167))))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+308)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_tuplestore_select_read_pointer(m, v170, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L22
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v24&int32(8) == int32(0) {
		goto L7
	} else {
		goto L98
	}
L63:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v174 != int64(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v191 = int32(1)
	v203 = v168 - v191
	goto L72
L65:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v177 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+4)))
	if v178&int32(2) == int32(0) {
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v184 = int32(1)
	v186 = F_tuplestore_gettupleslot(m, v183, v184, v184, v177)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L22
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v186 == int32(0) {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	goto L64
L72:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v222 == int32(0) {
		goto L8
	} else {
		goto L74
	}
L73:
	;
	goto L8
L74:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)))
	if v225&int32(2) != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v222)+6)))
	if v228 < v168 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_slot_getsomeattrs_int(m, v222, v168)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L22
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v233 = v203 << (uint(int32(2)) % 32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233+v234)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v222)+20))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v203))))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v241 = int32(*(*int16)(unsafe.Add(mBase, uint32(v240)+6)))
	if v241 < v168 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	F_slot_getsomeattrs_int(m, v240, v168)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L22
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)+20))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v203))))
	if v239&int32(1) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L82
L84:
	;
	v273 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v275 = v273 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v275
	F_spool_tuples(m, l0, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L22
	} else {
		goto L95
	}
L85:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267+v233)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v271 = F_FunctionCall5Coll(m, l0+int32(248), v266, v236, v269, v270, (base.B2i32(v24&int32(2048) == int32(0))^v169)&v191, (v169^v191)&v191)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L22
	} else {
		goto L93
	}
L86:
	;
	if v247&int32(1) == int32(0) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+309)))
	if v259 == int32(0) {
		goto L8
	} else {
		goto L91
	}
L89:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+309)))
	if v256 == int32(0) {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	goto L8
L91:
	;
	if v247&int32(1) == int32(0) {
		goto L84
	} else {
		goto L92
	}
L92:
	;
	goto L8
L93:
	;
	if v271 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	goto L84
L95:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v280 = int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v283 = F_tuplestore_gettupleslot(m, v279, v280, v280, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	if v283 != 0 {
		goto L72
	} else {
		goto L97
	}
L97:
	;
	goto L73
L98:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v290)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_tuplestore_select_read_pointer(m, v292, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	v296 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v296 != int64(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v313 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L101:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v299 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+4)))
	if v300&int32(2) == int32(0) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v306 = int32(1)
	v308 = F_tuplestore_gettupleslot(m, v305, v306, v306, v299)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L22
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	if v308 == int32(0) {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L100
L108:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+8))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	m.T0[v418].(func(*base.Module, int32))(m, v416)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L22
	} else {
		goto L131
	}
L109:
	;
	if v24&int32(2048) != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v320 = int64(0) - v291
	goto L112
L111:
	;
	v320 = v291
	goto L112
L112:
	;
	v324 = v313
	goto L113
L113:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+4)))
	if v338&int32(2) != 0 {
		goto L108
	} else {
		goto L115
	}
L114:
	;
	goto L108
L115:
	;
	v341 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if v289+v320 <= v341 {
		goto L108
	} else {
		goto L116
	}
L116:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+32))
	m.T0[v345].(func(*base.Module, int32, int32))(m, v343, v324)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L22
	} else {
		goto L117
	}
L117:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v350 = v348 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v350
	F_spool_tuples(m, l0, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L22
	} else {
		goto L118
	}
L118:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v355 = int32(1)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v358 = F_tuplestore_gettupleslot(m, v354, v355, v355, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L22
	} else {
		goto L119
	}
L119:
	;
	if v358 == int32(0) {
		goto L108
	} else {
		goto L120
	}
L120:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+96))
	if v363 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v399 != 0 {
		v324 = v399
		goto L113
	} else {
		goto L130
	}
L122:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	*(*int32)(unsafe.Add(mBase, uint32(v367)+8)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v367)+12)) = v366
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v371 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v367)+20))
	F_MemoryContextReset(m, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L22
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v377 = int32(4515712)
	v378 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v367)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v380
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v371)+20))
	v385 = m.T0[v384].(func(*base.Module, int32, int32, int32) int32)(m, v371, v367, v19+int32(15))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L22
	} else {
		goto L127
	}
L126:
	;
	goto L121
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v378
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v367)+20))
	F_MemoryContextReset(m, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	if v385 != 0 {
		goto L121
	} else {
		goto L129
	}
L129:
	;
	v392 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v392 + int64(1)
	goto L121
L130:
	;
	goto L114
L131:
	;
	goto L8
L132:
	;
	F_errmsg_internal(m, int32(364677), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L22
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(498681), int32(1578), int32(136619))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L22
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
	F_errmsg_internal(m, int32(364677), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L22
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(498681), int32(1659), int32(136619))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L22
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errmsg_internal(m, int32(364677), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L22
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(498681), int32(1735), int32(136619))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L22
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_update_frametailpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
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
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v331 int64
	_ = v331
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int64
	_ = v360
	var v362 int64
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
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v411 int32
	_ = v411
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v451 int32
	_ = v451
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+381)))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L12
	} else {
		goto L142
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L139
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L12
	} else {
		goto L136
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = int32(4515712)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v31
	if v25&int32(256) != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v20 + int32(16)
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v28
	goto L6
L8:
	;
	v451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+381)) = uint8(v451)
	goto L7
L9:
	;
	F_spool_tuples(m, l0, int64(-1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if v25&int32(1024) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v38
	goto L8
L14:
	;
	if v25&int32(4) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v25&int32(20480) == int32(0) {
		goto L7
	} else {
		goto L51
	}
L17:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v44 + int64(1)
	goto L8
L18:
	;
	goto L19
L19:
	;
	if v25&int32(10) == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+96))
	if v52 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_spool_tuples(m, l0, int64(-1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_tuplestore_select_read_pointer(m, v60, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v58
	goto L8
L25:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v64 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	goto L34
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v67 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)))
	if v68&int32(2) == int32(0) {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v74 = int32(1)
	v76 = F_tuplestore_gettupleslot(m, v73, v74, v74, v67)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v76 == int32(0) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v98 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L35:
	;
	goto L8
L36:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v101&int32(2) != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	if v104 <= v105 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v143 = v141 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v143
	F_spool_tuples(m, l0, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L48
	}
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+96))
	if v108 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v98
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v115 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	F_MemoryContextReset(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v121 = int32(4515712)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v124
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	v129 = m.T0[v128].(func(*base.Module, int32, int32, int32) int32)(m, v115, v111, v20+int32(14))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L45
	}
L44:
	;
	goto L38
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v122
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	F_MemoryContextReset(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	if v129 == int32(0) {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	goto L38
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v148 = int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v151 = F_tuplestore_gettupleslot(m, v147, v148, v148, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	if v151 != 0 {
		goto L34
	} else {
		goto L50
	}
L50:
	;
	goto L35
L51:
	;
	if v25&int32(4) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	if v25&int32(4096) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v25&int32(2) != 0 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	v166 = int64(0) - v162
	goto L57
L56:
	;
	v166 = v162
	goto L57
L57:
	;
	v167 = v159 + v166
	v169 = v167 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v169
	if int64(0) <= v169 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v169 <= v159+int64(1) {
		goto L8
	} else {
		goto L61
	}
L59:
	;
	v181 = int64(0)
	goto L60
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v181
	goto L8
L61:
	;
	F_spool_tuples(m, l0, v167)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v179 <= v178 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v181 = v178
	goto L60
L64:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v26)+100))
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v185))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+308)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_tuplestore_select_read_pointer(m, v188, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v25&int32(8) == int32(0) {
		goto L7
	} else {
		goto L102
	}
L67:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v192 != int64(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v215 = v186 - int32(1)
	goto L76
L69:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v195 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+4)))
	if v196&int32(2) == int32(0) {
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v202 = int32(1)
	v204 = F_tuplestore_gettupleslot(m, v201, v202, v202, v195)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	if v204 == int32(0) {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	goto L68
L76:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v235 == int32(0) {
		goto L8
	} else {
		goto L78
	}
L77:
	;
	goto L8
L78:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+4)))
	if v238&int32(2) != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v241 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+6)))
	if v241 < v186 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_slot_getsomeattrs_int(m, v235, v186)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v246 = v215 << (uint(int32(2)) % 32)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v246+v247)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250+v215))))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253)+6)))
	if v254 < v186 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	F_slot_getsomeattrs_int(m, v253, v186)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L12
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v215))))
	if v252&int32(1) == int32(0) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L86
L88:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v286 = v284 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v286
	F_spool_tuples(m, l0, v286)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L12
	} else {
		goto L99
	}
L89:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v276+v246)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v280 = F_FunctionCall5Coll(m, l0+int32(276), v275, v249, v278, v279, base.B2i32(v25&int32(4096) == int32(0))^v187, v187)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L12
	} else {
		goto L97
	}
L90:
	;
	if v260&int32(1) == int32(0) {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+309)))
	if v272 != 0 {
		goto L88
	} else {
		goto L95
	}
L93:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+309)))
	if v269 == int32(0) {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	goto L8
L95:
	;
	if v260&int32(1) != 0 {
		goto L88
	} else {
		goto L96
	}
L96:
	;
	goto L8
L97:
	;
	if v280 == int32(0) {
		goto L8
	} else {
		goto L98
	}
L98:
	;
	goto L88
L99:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v291 = int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v294 = F_tuplestore_gettupleslot(m, v290, v291, v291, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	if v294 != 0 {
		goto L76
	} else {
		goto L101
	}
L101:
	;
	goto L77
L102:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_tuplestore_select_read_pointer(m, v303, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v307 != int64(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v324 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v310 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+4)))
	if v311&int32(2) == int32(0) {
		goto L104
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v317 = int32(1)
	v319 = F_tuplestore_gettupleslot(m, v316, v317, v317, v310)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	if v319 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L104
L112:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	m.T0[v431].(func(*base.Module, int32))(m, v429)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L12
	} else {
		goto L135
	}
L113:
	;
	if v25&int32(4096) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v331 = int64(0) - v302
	goto L116
L115:
	;
	v331 = v302
	goto L116
L116:
	;
	v335 = v324
	goto L117
L117:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+4)))
	if v350&int32(2) != 0 {
		goto L112
	} else {
		goto L119
	}
L118:
	;
	goto L112
L119:
	;
	v353 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
	if v300+v331 < v353 {
		goto L112
	} else {
		goto L120
	}
L120:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+32))
	m.T0[v357].(func(*base.Module, int32, int32))(m, v355, v335)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
	;
	v360 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v362 = v360 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v362
	F_spool_tuples(m, l0, v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L122
	}
L122:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v367 = int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v370 = F_tuplestore_gettupleslot(m, v366, v367, v367, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L12
	} else {
		goto L123
	}
L123:
	;
	if v370 == int32(0) {
		goto L112
	} else {
		goto L124
	}
L124:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+96))
	if v375 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v411 != 0 {
		v335 = v411
		goto L117
	} else {
		goto L134
	}
L126:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v379)+12)) = v378
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v383 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	F_MemoryContextReset(m, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L12
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v389 = int32(4515712)
	v390 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v392
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v383)+20))
	v397 = m.T0[v396].(func(*base.Module, int32, int32, int32) int32)(m, v383, v379, v20+int32(15))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L12
	} else {
		goto L131
	}
L130:
	;
	goto L125
L131:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v390
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	F_MemoryContextReset(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	if v397 != 0 {
		goto L125
	} else {
		goto L133
	}
L133:
	;
	v404 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v404 + int64(1)
	goto L125
L134:
	;
	goto L118
L135:
	;
	goto L8
L136:
	;
	F_errmsg_internal(m, int32(364677), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(498681), int32(1831), int32(136528))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L12
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errmsg_internal(m, int32(364677), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(498681), int32(1913), int32(136528))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errmsg_internal(m, int32(364677), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(498681), int32(1989), int32(136528))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
