package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(57602), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(471063), int32(774), int32(195893))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		v28 = F_heap_getsysattr(m, v6, l1, l2)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			return v28
		}
	}
}
func F_tts_minimal_getsomeattrs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v259 int32
	_ = v259
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v514 int32
	_ = v514
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)))
	v25 = v23 & int32(2047)
	if l1 < v25 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = l1
	goto L3
L2:
	;
	v27 = v25
	goto L3
L3:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
	v30 = v28 & int32(1)
	v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	m.G0 = v19 + int32(48)
	return
L5:
	;
	if v467 < v27 {
		goto L154
	} else {
		goto L155
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v30 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v39 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v36&int32(8) != 0 {
		v467 = v31
		v469 = v35
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v39 = v35
	goto L6
L11:
	;
	v467 = v241 + int32(1)
	v469 = v463
	goto L5
L12:
	;
	v463 = v459 + v304
	goto L11
L13:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v459 = int32(base.Ui32(v455) >> (uint(int32(2)) % 32))
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v436
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v27)
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v453 = v451 & int32(65527)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v453)
	goto L4
L15:
	;
	if v27 <= v31 {
		v436 = v39
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v27 <= v31 {
		v436 = v39
		goto L14
	} else {
		goto L86
	}
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v46 = v22 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = v31
	v53 = v39
	goto L21
L19:
	;
	v467 = v51 + int32(1)
	v469 = v229
	goto L5
L20:
	;
	v229 = v225 + v95
	goto L19
L21:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v40))) = uint8(v67)
	v72 = v47 + int32(20) + v51<<(uint(int32(4))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v67 <= v73 {
		v95 = v73
		v96 = v67
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v225 = int32(base.Ui32(v221) >> (uint(int32(2)) % 32))
	goto L20
L23:
	;
	v97 = v95 + v46
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+6)))
	if v101 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+12)))
	v82 = (v53 + v76 - int32(1)) & (int32(0) - v76)
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v83 == int32(65535) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v95 = v93
	v96 = int32(0)
	goto L23
L26:
	;
	if v53 == v82 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v82
	v93 = v82
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v53
	v93 = v53
	goto L25
L30:
	;
	goto L31
L31:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v46))))
	if v89 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v90 = v53
	goto L34
L33:
	;
	v90 = v82
	goto L34
L34:
	;
	v95 = v90
	v96 = int32(1)
	goto L23
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v51<<(uint(int32(2))%32)))) = v125
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v127 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	switch v104 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L41
	default:
		goto L39
	case 3:
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v125 = v97
	goto L35
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v125 = v109
	goto L35
L41:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97))))
	v125 = v108
	goto L35
L42:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	v125 = v107
	goto L35
L43:
	;
	return
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = base.I32_extend16_s(v104)
	F_errmsg_internal(m, int32(460720), v19)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(310223), int32(70), int32(64535))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	goto L22
L48:
	;
	if v127 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v217 = v127 + v95
	if v96 != 0 {
		v229 = v217
		goto L19
	} else {
		goto L84
	}
L51:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v132 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v97&int32(3) == int32(0) {
		v180 = v97
		goto L69
	} else {
		goto L70
	}
L54:
	;
	v135 = int32(6)
	v137 = int32(18)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v139 == v137 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v132&int32(1) == int32(0) {
		goto L47
	} else {
		goto L66
	}
L57:
	;
	v142 = v137
	goto L59
L58:
	;
	v142 = int32(2)
	goto L59
L59:
	;
	if v139&int32(254) == int32(2) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v147 = v135
	goto L62
L61:
	;
	v147 = v142
	goto L62
L62:
	;
	if v139 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v150 = v135
	goto L65
L64:
	;
	v150 = v147
	goto L65
L65:
	;
	v225 = v150
	goto L20
L66:
	;
	v225 = int32(base.Ui32(v132) >> (uint(int32(1)) % 32))
	goto L20
L67:
	;
	v229 = v213 + v95 + int32(1)
	goto L19
L68:
	;
	v213 = v205 - v97
	goto L67
L69:
	;
	v184 = v180
	goto L78
L70:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v164 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v213 = int32(0)
	goto L67
L72:
	;
	goto L73
L73:
	;
	v169 = v97
	goto L74
L74:
	;
	v173 = v169 + int32(1)
	if v173&int32(3) == int32(0) {
		v180 = v173
		goto L69
	} else {
		goto L76
	}
L75:
	;
	v205 = v173
	goto L68
L76:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v178 != 0 {
		v169 = v173
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v193 = int32(-2139062144)
	if (int32(16843008)-v190|v190)&v193 == v193 {
		v184 = v184 + int32(4)
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v199 = v184
	goto L81
L80:
	;
	goto L79
L81:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v203 != 0 {
		v199 = v199 + int32(1)
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v205 = v199
	goto L68
L83:
	;
	goto L82
L84:
	;
	v219 = v51 + int32(1)
	if v219 != v27 {
		v51 = v219
		v53 = v217
		goto L21
	} else {
		goto L85
	}
L85:
	;
	v436 = v217
	goto L14
L86:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v236 = v22 + v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v241 = v31
	v243 = v39
	goto L87
L87:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(23)+v241>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v259)>>(uint(v241&int32(7))%32))&int32(1) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v436 = v429
	goto L14
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v241<<(uint(int32(2))%32)))) = int32(0)
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v241+v40))) = uint8(v273)
	v467 = v241 + v273
	v469 = v243
	goto L5
L90:
	;
	goto L91
L91:
	;
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241+v40))) = uint8(v277)
	v283 = v237 + int32(20) + v241<<(uint(int32(4))%32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v277 <= v284 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v307 = v304 + v236
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+6)))
	if v311 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L93:
	;
	v304 = v284
	v306 = v277
	goto L92
L94:
	;
	goto L95
L95:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+12)))
	v293 = (v243 + v287 - int32(1)) & (int32(0) - v287)
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+4)))
	if v294 == int32(65535) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v243 == v293 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v293
	v304 = v293
	v306 = v277
	goto L92
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v243
	v304 = v243
	v306 = v277
	goto L92
L100:
	;
	goto L101
L101:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v236))))
	if v300 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v301 = v243
	goto L104
L103:
	;
	v301 = v293
	goto L104
L104:
	;
	v304 = v301
	v306 = int32(1)
	goto L92
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v241<<(uint(int32(2))%32)))) = v337
	v339 = int32(*(*int16)(unsafe.Add(mBase, uint32(v283)+4)))
	if v339 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L106:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+4)))
	switch v314 - int32(1) {
	case 0:
		goto L112
	case 1:
		goto L111
	default:
		goto L109
	case 3:
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v337 = v307
	goto L105
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L43
	} else {
		goto L113
	}
L110:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v337 = v319
	goto L105
L111:
	;
	v318 = int32(*(*int16)(unsafe.Add(mBase, uint32(v307))))
	v337 = v318
	goto L105
L112:
	;
	v317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v307))))
	v337 = v317
	goto L105
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = base.I32_extend16_s(v314)
	F_errmsg_internal(m, int32(460720), v19+int32(32))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L43
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(310223), int32(70), int32(64535))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L43
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	if v339 == int32(-1) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	v429 = v339 + v304
	if v306 != 0 {
		v463 = v429
		goto L11
	} else {
		goto L152
	}
L119:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v344 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	if v307&int32(3) == int32(0) {
		v392 = v307
		goto L137
	} else {
		goto L138
	}
L122:
	;
	v347 = int32(6)
	v349 = int32(18)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)))
	if v351 == v349 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v344&int32(1) == int32(0) {
		goto L13
	} else {
		goto L134
	}
L125:
	;
	v354 = v349
	goto L127
L126:
	;
	v354 = int32(2)
	goto L127
L127:
	;
	if v351&int32(254) == int32(2) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v359 = v347
	goto L130
L129:
	;
	v359 = v354
	goto L130
L130:
	;
	if v351 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v362 = v347
	goto L133
L132:
	;
	v362 = v359
	goto L133
L133:
	;
	v459 = v362
	goto L12
L134:
	;
	v459 = int32(base.Ui32(v344) >> (uint(int32(1)) % 32))
	goto L12
L135:
	;
	v463 = v425 + v304 + int32(1)
	goto L11
L136:
	;
	v425 = v417 - v307
	goto L135
L137:
	;
	v396 = v392
	goto L146
L138:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v376 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v425 = int32(0)
	goto L135
L140:
	;
	goto L141
L141:
	;
	v381 = v307
	goto L142
L142:
	;
	v385 = v381 + int32(1)
	if v385&int32(3) == int32(0) {
		v392 = v385
		goto L137
	} else {
		goto L144
	}
L143:
	;
	v417 = v385
	goto L136
L144:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v390 != 0 {
		v381 = v385
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v405 = int32(-2139062144)
	if (int32(16843008)-v402|v402)&v405 == v405 {
		v396 = v396 + int32(4)
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v411 = v396
	goto L149
L148:
	;
	goto L147
L149:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v415 != 0 {
		v411 = v411 + int32(1)
		goto L149
	} else {
		goto L151
	}
L150:
	;
	v417 = v411
	goto L136
L151:
	;
	goto L150
L152:
	;
	v431 = v241 + int32(1)
	if v431 != v27 {
		v241 = v431
		v243 = v429
		goto L87
	} else {
		goto L153
	}
L153:
	;
	goto L88
L154:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+22)))
	v490 = v486 + v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v494 = v467
	v496 = v469
	goto L157
L155:
	;
	v683 = v467
	v685 = v469
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v685
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v683)
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v702 = v700 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v702)
	goto L4
L157:
	;
	if v30 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v683 = v27
	v685 = v675
	goto L156
L159:
	;
	v680 = v494 + int32(1)
	if v680 != v27 {
		v494 = v680
		v496 = v675
		goto L157
	} else {
		goto L217
	}
L160:
	;
	v529 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v494+v491))) = uint8(v529)
	v533 = v483 + int32(20) + v494<<(uint(int32(4))%32)
	v534 = int32(*(*int16)(unsafe.Add(mBase, uint32(v533)+4)))
	v535 = int32(65535)
	v536 = v534 & v535
	if v536 == v535 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486+int32(23)+v494>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v514)>>(uint(v494&int32(7))%32))&int32(1) != 0 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v492+v494<<(uint(int32(2))%32)))) = int32(0)
	v526 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v494+v491))) = uint8(v526)
	v675 = v496
	goto L159
L163:
	;
	v550 = v548 + v490
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+6)))
	if v554 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496+v490))))
	if v540 != 0 {
		v548 = v496
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+12)))
	v548 = (v496 + v541 - int32(1)) & (int32(0) - v541)
	goto L163
L167:
	;
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v492+v494<<(uint(int32(2))%32)))) = v577
	v579 = int32(*(*int16)(unsafe.Add(mBase, uint32(v533)+4)))
	if int32(0) < v579 {
		goto L179
	} else {
		goto L180
	}
L169:
	;
	switch v536 - int32(1) {
	case 0:
		goto L175
	case 1:
		goto L174
	default:
		goto L172
	case 3:
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v577 = v550
	goto L168
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L43
	} else {
		goto L176
	}
L173:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v577 = v561
	goto L168
L174:
	;
	v560 = int32(*(*int16)(unsafe.Add(mBase, uint32(v550))))
	v577 = v560
	goto L168
L175:
	;
	v559 = int32(*(*int8)(unsafe.Add(mBase, uint32(v550))))
	v577 = v559
	goto L168
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v534
	F_errmsg_internal(m, int32(460720), v19+int32(16))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L43
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(310223), int32(70), int32(64535))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L43
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v675 = v579 + v548
	goto L159
L180:
	;
	goto L181
L181:
	;
	if v579 == int32(-1) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	if v585 == int32(1) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L184
L184:
	;
	if v550&int32(3) == int32(0) {
		v637 = v550
		goto L202
	} else {
		goto L203
	}
L185:
	;
	v588 = int32(6)
	v590 = int32(18)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+1)))
	if v592 == v590 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	goto L187
L187:
	;
	if v585&int32(1) != 0 {
		goto L197
	} else {
		goto L198
	}
L188:
	;
	v595 = v590
	goto L190
L189:
	;
	v595 = int32(2)
	goto L190
L190:
	;
	if v592&int32(254) == int32(2) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v600 = v588
	goto L193
L192:
	;
	v600 = v595
	goto L193
L193:
	;
	if v592 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v603 = v588
	goto L196
L195:
	;
	v603 = v600
	goto L196
L196:
	;
	v675 = v603 + v548
	goto L159
L197:
	;
	v675 = int32(base.Ui32(v585)>>(uint(int32(1))%32)) + v548
	goto L159
L198:
	;
	goto L199
L199:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v675 = int32(base.Ui32(v610)>>(uint(int32(2))%32)) + v548
	goto L159
L200:
	;
	v675 = v670 + v548 + int32(1)
	goto L159
L201:
	;
	v670 = v662 - v550
	goto L200
L202:
	;
	v641 = v637
	goto L211
L203:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	if v621 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v670 = int32(0)
	goto L200
L205:
	;
	goto L206
L206:
	;
	v626 = v550
	goto L207
L207:
	;
	v630 = v626 + int32(1)
	if v630&int32(3) == int32(0) {
		v637 = v630
		goto L202
	} else {
		goto L209
	}
L208:
	;
	v662 = v630
	goto L201
L209:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	if v635 != 0 {
		v626 = v630
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	v650 = int32(-2139062144)
	if (int32(16843008)-v647|v647)&v650 == v650 {
		v641 = v641 + int32(4)
		goto L211
	} else {
		goto L213
	}
L212:
	;
	v656 = v641
	goto L214
L213:
	;
	goto L212
L214:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656))))
	if v660 != 0 {
		v656 = v656 + int32(1)
		goto L214
	} else {
		goto L216
	}
L215:
	;
	v662 = v656
	goto L201
L216:
	;
	goto L215
L217:
	;
	goto L158
}
func F_tts_minimal_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(57602), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(471063), int32(564), int32(195943))
				v21 = m.ExcPending
				if v21 != 0 {
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
func F_tts_minimal_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
	return
}
func F_tts_virtual_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_heap_form_minimal_tuple(m, v3, v4, v5, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
