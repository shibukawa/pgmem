package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(57712), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(471063), int32(796), int32(364436))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	return v144
L10:
	;
	v144 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v35 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v144 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v39 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v144 = v136
	goto L9
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v43 == int32(0) {
		v136 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v107 = int32(0)
	v109 = v39 - int32(1)
	goto L39
L20:
	;
	v48 = v43
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v53 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v136 = int32(0)
	goto L16
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v100 != 0 {
		v48 = v100
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v56 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v59 = int32(1)
	if v24 == v56 {
		v136 = v59
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v63 = v61 - int32(1)
	if v63 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v68 = int32(0)
	v70 = v63
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v76 = int32(2)
	v77 = base.I32_div_s(v70-v68, v76)
	v78 = v77 + v68
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74+v78<<(uint(v76)%32))))
	if v82 == v24 {
		v136 = v59
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v86 = F_TransactionIdPrecedes(m, v82, v24)
	mBase = m.M
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v87 = v78 + int32(1)
	goto L33
L32:
	;
	v87 = v68
	goto L33
L33:
	;
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v90 = v70
	goto L36
L35:
	;
	v90 = v78 - int32(1)
	goto L36
L36:
	;
	if v87 <= v90 {
		v68 = v87
		v70 = v90
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L22
L39:
	;
	v114 = int32(2)
	v115 = base.I32_div_s(v109-v107, v114)
	v116 = v115 + v107
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v114)%32))))
	v121 = base.B2i32(v120 == v24)
	if v120 == v24 {
		v136 = v121
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v136 = v121
	goto L16
L41:
	;
	v124 = base.B2i32(base.Ui32(v120) < base.Ui32(v24))
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v125 = v116 + int32(1)
	goto L44
L43:
	;
	v125 = v107
	goto L44
L44:
	;
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v128 = v109
	goto L47
L46:
	;
	v128 = v116 - int32(1)
	goto L47
L47:
	;
	if v125 <= v128 {
		v107 = v125
		v109 = v128
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
}
func F_tts_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 != 0 {
		v33 = v5
		v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			return v35
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v7&int32(4) != 0 {
			v33 = int32(0)
			v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v35
			}
		} else {
			v10 = int32(4442576)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v13
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v15
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = F_heap_form_tuple(m, v19, v20, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v22
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v29 = v27 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v29)
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
				v33 = v22
				v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_tts_heap_getsomeattrs(m *base.Module, l0 int32, l1 int32) {
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v436
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v685
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
func F_tts_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
					F_errfinish(m, int32(471063), int32(369), int32(195873))
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
func F_tts_minimal_copy_heap_tuple(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v5 != 0 {
		v41 = v5
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
		v47 = F_palloc(m, v44+int32(32))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v49)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = v44 + int32(8)
			v59 = v47 + int32(24)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v59
			v62 = v47 + int32(32)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			if v63 != 0 {
				v64 = F__emscripten_memcpy_bulkmem(m, v62, v41, v63)
				mBase = m.M
				v65 = v64
			} else {
				v65 = v62
			}
			v66 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v66)
			v68 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v65))) = v68
			*(*int64)(unsafe.Add(mBase, uint32(v59))) = v68
			return v47
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v7&int32(4) != 0 {
			v41 = int32(0)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			v47 = F_palloc(m, v44+int32(32))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v49)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = v44 + int32(8)
				v59 = v47 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v59
				v62 = v47 + int32(32)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v63 != 0 {
					v64 = F__emscripten_memcpy_bulkmem(m, v62, v41, v63)
					mBase = m.M
					v65 = v64
				} else {
					v65 = v62
				}
				v66 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v66)
				v68 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v65))) = v68
				*(*int64)(unsafe.Add(mBase, uint32(v59))) = v68
				return v47
			}
		} else {
			v10 = int32(4442576)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v13
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v15
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v23 = F_heap_form_minimal_tuple(m, v19, v20, v21, v15)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v23
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v30 = v28 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v33 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v23 - v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v32 + v33
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
				v41 = v23
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v47 = F_palloc(m, v44+int32(32))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v49)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = v44 + int32(8)
					v59 = v47 + int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v59
					v62 = v47 + int32(32)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					if v63 != 0 {
						v64 = F__emscripten_memcpy_bulkmem(m, v62, v41, v63)
						mBase = m.M
						v65 = v64
					} else {
						v65 = v62
					}
					v66 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v66)
					v68 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = v68
					*(*int64)(unsafe.Add(mBase, uint32(v59))) = v68
					return v47
				}
			}
		}
	}
}
func F_tts_minimal_copyslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = int32(4442576)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = v5
		v17 = F_ExecStoreMinimalTuple(m, v12, l0, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F_tts_virtual_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v3&int32(4) != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_pfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
			v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v14 = v11 & int32(-5)
			v15 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v15)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v22 = v14 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v22)
			return
		}
	} else {
		v14 = v3
		v15 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v15)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
		v22 = v14 | int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v22)
		return
	}
}
func F_tts_virtual_copyslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v7&int32(4) != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_pfree(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v18 = v15 & int32(-5)
			v19 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v19)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v19)
			v26 = v18 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v26)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
			if v29 <= v30 {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if int32(0) < v141 {
					v147 = int32(0)
					for {
						v151 = v147 << (uint(int32(2)) % 32)
						v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v151)))
						*(*int32)(unsafe.Add(mBase, uint32(v151+v152))) = v156
						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v147))))
						*(*uint8)(unsafe.Add(mBase, uint32(v158+v147))) = uint8(v162)
						v165 = v147 + int32(1)
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v165 < v166 {
							v147 = v165
							continue
						} else {
							break
						}
						break
					}
					v171 = v166
				} else {
					v171 = v141
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v171)
				v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v176 = v174 & int32(65533)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v176)
				F_tts_virtual_materialize(m, l0)
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return
				} else {
					return
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
				m.T0[v33].(func(*base.Module, int32, int32))(m, l1, v29)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
					if v29 <= v36 {
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
						if v42 == int32(0) {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v121 = int32(2)
							v125 = v29 - v36
							v128 = F___memset(m, v120+v36<<(uint(v121)%32), int32(0), v125<<(uint(v121)%32))
							mBase = m.M
							v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v132 = F___memset(m, v129+v36, int32(1), v125)
							mBase = m.M
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
							if v45 == int32(0) {
								v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v121 = int32(2)
								v125 = v29 - v36
								v128 = F___memset(m, v120+v36<<(uint(v121)%32), int32(0), v125<<(uint(v121)%32))
								mBase = m.M
								v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								v132 = F___memset(m, v129+v36, int32(1), v125)
								mBase = m.M
							} else {
								if v29 <= v36 {
								} else {
									v49 = int32(1)
									v50 = v36 + v49
									if (v29-v36)&v49 != 0 {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
										v60 = v45 + v36<<(uint(int32(3))%32)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v54+v36<<(uint(int32(2))%32)))) = v61
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
										v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
										v67 = v65 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v63+v36))) = uint8(v67)
										v69 = v50
									} else {
										v69 = v36
									}
									if v29 == v50 {
									} else {
										v73 = v69
										for {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
											v79 = int32(2)
											v82 = int32(3)
											v84 = v45 + v73<<(uint(v82)%32)
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(v79)%32)))) = v85
											v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
											v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
											v90 = int32(1)
											v91 = v89 ^ v90
											*(*uint8)(unsafe.Add(mBase, uint32(v87+v73))) = uint8(v91)
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
											v95 = v73 + v90
											v101 = v45 + v95<<(uint(v82)%32)
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v93+v95<<(uint(v79)%32)))) = v102
											v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
											v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
											v108 = v106 ^ v90
											*(*uint8)(unsafe.Add(mBase, uint32(v104+v95))) = uint8(v108)
											v111 = v73 + v79
											if v111 != v29 {
												v73 = v111
												continue
											} else {
												break
											}
											break
										}
									}
								}
							}
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
					}
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if int32(0) < v141 {
						v147 = int32(0)
						for {
							v151 = v147 << (uint(int32(2)) % 32)
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v151)))
							*(*int32)(unsafe.Add(mBase, uint32(v151+v152))) = v156
							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v147))))
							*(*uint8)(unsafe.Add(mBase, uint32(v158+v147))) = uint8(v162)
							v165 = v147 + int32(1)
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v165 < v166 {
								v147 = v165
								continue
							} else {
								break
							}
							break
						}
						v171 = v166
					} else {
						v171 = v141
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v171)
					v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v176 = v174 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v176)
					F_tts_virtual_materialize(m, l0)
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v18 = v7
		v19 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v19)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v19)
		v26 = v18 | int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
		v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
		if v29 <= v30 {
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if int32(0) < v141 {
				v147 = int32(0)
				for {
					v151 = v147 << (uint(int32(2)) % 32)
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v151)))
					*(*int32)(unsafe.Add(mBase, uint32(v151+v152))) = v156
					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v147))))
					*(*uint8)(unsafe.Add(mBase, uint32(v158+v147))) = uint8(v162)
					v165 = v147 + int32(1)
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if v165 < v166 {
						v147 = v165
						continue
					} else {
						break
					}
					break
				}
				v171 = v166
			} else {
				v171 = v141
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v171)
			v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v176 = v174 & int32(65533)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v176)
			F_tts_virtual_materialize(m, l0)
			mBase = m.M
			v179 = m.ExcPending
			if v179 != 0 {
				return
			} else {
				return
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
			m.T0[v33].(func(*base.Module, int32, int32))(m, l1, v29)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				if v29 <= v36 {
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
					if v42 == int32(0) {
						v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v121 = int32(2)
						v125 = v29 - v36
						v128 = F___memset(m, v120+v36<<(uint(v121)%32), int32(0), v125<<(uint(v121)%32))
						mBase = m.M
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v132 = F___memset(m, v129+v36, int32(1), v125)
						mBase = m.M
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						if v45 == int32(0) {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v121 = int32(2)
							v125 = v29 - v36
							v128 = F___memset(m, v120+v36<<(uint(v121)%32), int32(0), v125<<(uint(v121)%32))
							mBase = m.M
							v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v132 = F___memset(m, v129+v36, int32(1), v125)
							mBase = m.M
						} else {
							if v29 <= v36 {
							} else {
								v49 = int32(1)
								v50 = v36 + v49
								if (v29-v36)&v49 != 0 {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
									v60 = v45 + v36<<(uint(int32(3))%32)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v54+v36<<(uint(int32(2))%32)))) = v61
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
									v67 = v65 ^ int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v63+v36))) = uint8(v67)
									v69 = v50
								} else {
									v69 = v36
								}
								if v29 == v50 {
								} else {
									v73 = v69
									for {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
										v79 = int32(2)
										v82 = int32(3)
										v84 = v45 + v73<<(uint(v82)%32)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(v79)%32)))) = v85
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
										v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
										v90 = int32(1)
										v91 = v89 ^ v90
										*(*uint8)(unsafe.Add(mBase, uint32(v87+v73))) = uint8(v91)
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
										v95 = v73 + v90
										v101 = v45 + v95<<(uint(v82)%32)
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v93+v95<<(uint(v79)%32)))) = v102
										v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
										v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
										v108 = v106 ^ v90
										*(*uint8)(unsafe.Add(mBase, uint32(v104+v95))) = uint8(v108)
										v111 = v73 + v79
										if v111 != v29 {
											v73 = v111
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
				}
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if int32(0) < v141 {
					v147 = int32(0)
					for {
						v151 = v147 << (uint(int32(2)) % 32)
						v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v154+v151)))
						*(*int32)(unsafe.Add(mBase, uint32(v151+v152))) = v156
						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v147))))
						*(*uint8)(unsafe.Add(mBase, uint32(v158+v147))) = uint8(v162)
						v165 = v147 + int32(1)
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v165 < v166 {
							v147 = v165
							continue
						} else {
							break
						}
						break
					}
					v171 = v166
				} else {
					v171 = v141
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v171)
				v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v176 = v174 & int32(65533)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v176)
				F_tts_virtual_materialize(m, l0)
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_tts_virtual_materialize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	v2 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v10&int32(4) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v2
	v24 = v2
	v27 = v14
	goto L4
L4:
	;
	v30 = v13 + int32(20) + v24<<(uint(int32(4))%32)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+6)))
	if v31 != 0 {
		v166 = v22
		v168 = v27
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v166 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L6:
	;
	v170 = v24 + int32(1)
	if v170 < v168 {
		v22 = v166
		v24 = v170
		v27 = v168
		goto L4
	} else {
		goto L44
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v24))))
	if v34 != 0 {
		v166 = v22
		v168 = v27
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v24<<(uint(int32(2))%32))))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+4)))
	if v40 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v166 = v160 + v161
	v168 = v27
	goto L6
L10:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v148 = int32(1)
	v152 = (v22 + v146 - v148) & (int32(0) - v146)
	if v43&v148 != 0 {
		goto L41
	} else {
		goto L42
	}
L11:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if base.Ui32((v132-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v160 = v52
		v161 = int32(6)
		goto L9
	} else {
		goto L37
	}
L12:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v43 != int32(1) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v67 = int32(0)
	v69 = (v22 + v63 - int32(1)) & (v67 - v63)
	if v67 < v40 {
		v160 = v69
		v161 = v40
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v52 = (v22 + v46 - int32(1)) & (int32(0) - v46)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v53&int32(254) != int32(2) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+2))
	v59 = F_EOH_get_flat_size(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v166 = v59 + v52
	v168 = v62
	goto L6
L19:
	;
	if v39&int32(3) == int32(0) {
		v95 = v39
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v160 = v69
	v161 = v128 + int32(1)
	goto L9
L21:
	;
	v128 = v120 - v39
	goto L20
L22:
	;
	v99 = v95
	goto L31
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v128 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v84 = v39
	goto L27
L27:
	;
	v88 = v84 + int32(1)
	if v88&int32(3) == int32(0) {
		v95 = v88
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v120 = v88
	goto L21
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v93 != 0 {
		v84 = v88
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v108 = int32(-2139062144)
	if (int32(16843008)-v105|v105)&v108 == v108 {
		v99 = v99 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v114 = v99
	goto L34
L33:
	;
	goto L32
L34:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v118 != 0 {
		v114 = v114 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v120 = v114
	goto L21
L36:
	;
	goto L35
L37:
	;
	v139 = int32(18)
	if v132&int32(255) == v139 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v145 = v139
	goto L40
L39:
	;
	v145 = int32(2)
	goto L40
L40:
	;
	v160 = v52
	v161 = v145
	goto L9
L41:
	;
	v160 = v152
	v161 = int32(base.Ui32(v43) >> (uint(int32(1)) % 32))
	goto L9
L42:
	;
	goto L43
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v160 = v152
	v161 = int32(base.Ui32(v157) >> (uint(int32(2)) % 32))
	goto L9
L44:
	;
	goto L5
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v175 = F_MemoryContextAlloc(m, v174, v166)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v175
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v180 = v178 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v182 <= int32(0) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v191 = v175
	v193 = int32(0)
	goto L48
L48:
	;
	v199 = v13 + int32(20) + v193<<(uint(int32(4))%32)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+6)))
	if v200 != 0 {
		v351 = v191
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L1
L50:
	;
	v355 = v193 + int32(1)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v355 < v356 {
		v191 = v351
		v193 = v355
		goto L48
	} else {
		goto L92
	}
L51:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v193))))
	if v203 != 0 {
		v351 = v191
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v205 = v193 << (uint(int32(2)) % 32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205+v206)))
	v209 = int32(*(*int16)(unsafe.Add(mBase, uint32(v199)+4)))
	if v209 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v342 != 0 {
		goto L89
	} else {
		goto L90
	}
L54:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+12)))
	v328 = int32(1)
	v332 = (v191 + v326 - v328) & (int32(0) - v326)
	if v212&v328 != 0 {
		goto L85
	} else {
		goto L86
	}
L55:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+12)))
	v306 = int32(1)
	v310 = (v191 + v304 - v306) & (int32(0) - v304)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if base.Ui32((v312-v306)&int32(255)) < base.Ui32(int32(3)) {
		v340 = v310
		v342 = int32(6)
		goto L53
	} else {
		goto L81
	}
L56:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v212 != int32(1) {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+12)))
	v240 = int32(0)
	v242 = (v191 + v236 - int32(1)) & (v240 - v236)
	if v240 < v209 {
		v340 = v242
		v342 = v209
		goto L53
	} else {
		goto L63
	}
L59:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v215&int32(254) != int32(2) {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v208)+2))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+12)))
	v227 = (v191 + v221 - int32(1)) & (int32(0) - v221)
	v228 = F_EOH_get_flat_size(m, v220)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L17
	} else {
		goto L61
	}
L61:
	;
	F_EOH_flatten_into(m, v220, v227, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L17
	} else {
		goto L62
	}
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v232+v205))) = v227
	v351 = v227 + v228
	goto L50
L63:
	;
	if v208&int32(3) == int32(0) {
		v268 = v208
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v340 = v242
	v342 = v301 + int32(1)
	goto L53
L65:
	;
	v301 = v293 - v208
	goto L64
L66:
	;
	v272 = v268
	goto L75
L67:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v252 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v301 = int32(0)
	goto L64
L69:
	;
	goto L70
L70:
	;
	v257 = v208
	goto L71
L71:
	;
	v261 = v257 + int32(1)
	if v261&int32(3) == int32(0) {
		v268 = v261
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v293 = v261
	goto L65
L73:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	if v266 != 0 {
		v257 = v261
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v281 = int32(-2139062144)
	if (int32(16843008)-v278|v278)&v281 == v281 {
		v272 = v272 + int32(4)
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v287 = v272
	goto L78
L77:
	;
	goto L76
L78:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v291 != 0 {
		v287 = v287 + int32(1)
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v293 = v287
	goto L65
L80:
	;
	goto L79
L81:
	;
	v319 = int32(18)
	if v312&int32(255) == v319 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v325 = v319
	goto L84
L83:
	;
	v325 = int32(2)
	goto L84
L84:
	;
	v340 = v310
	v342 = v325
	goto L53
L85:
	;
	v340 = v332
	v342 = int32(base.Ui32(v212) >> (uint(int32(1)) % 32))
	goto L53
L86:
	;
	goto L87
L87:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v340 = v332
	v342 = int32(base.Ui32(v337) >> (uint(int32(2)) % 32))
	goto L53
L88:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v205))) = v344
	v351 = v344 + v342
	goto L50
L89:
	;
	v343 = F__emscripten_memcpy_bulkmem(m, v340, v208, v342)
	mBase = m.M
	v344 = v343
	goto L91
L90:
	;
	v344 = v340
	goto L91
L91:
	;
	goto L88
L92:
	;
	goto L49
}
