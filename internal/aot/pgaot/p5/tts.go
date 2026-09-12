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
	F_errmsg(m, int32(61654), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(505602), int32(796), int32(392026))
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[38]))
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[39]))
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[40]))
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
	v105 = *(*int32)(unsafe.Add(mBase, _consts[41]))
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
			v10 = int32(4536272)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[10])) = v13
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
				*(*int32)(unsafe.Add(mBase, _consts[10])) = v11
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
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v402 int32
	_ = v402
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
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
	if v355 < v27 {
		goto L120
	} else {
		goto L121
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
		v355 = v31
		v357 = v35
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
	v355 = v185 + int32(1)
	v357 = v351
	goto L5
L12:
	;
	v351 = v347 + v248
	goto L11
L13:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v347 = int32(base.Ui32(v343) >> (uint(int32(2)) % 32))
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v324
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v27)
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v341 = v339 & int32(65527)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v341)
	goto L4
L15:
	;
	if v27 <= v31 {
		v324 = v39
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
		v324 = v39
		goto L14
	} else {
		goto L69
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
	v355 = v51 + int32(1)
	v357 = v173
	goto L5
L20:
	;
	v173 = v169 + v95
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
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v169 = int32(base.Ui32(v165) >> (uint(int32(2)) % 32))
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
	F_errmsg_internal(m, int32(493552), v19)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(333866), int32(70), int32(69245))
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
	v161 = v127 + v95
	if v96 != 0 {
		v173 = v161
		goto L19
	} else {
		goto L67
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
	v157 = F_strlen(m, v97)
	mBase = m.M
	v173 = v157 + v95 + int32(1)
	goto L19
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
	v169 = v150
	goto L20
L66:
	;
	v169 = int32(base.Ui32(v132) >> (uint(int32(1)) % 32))
	goto L20
L67:
	;
	v163 = v51 + int32(1)
	if v163 != v27 {
		v51 = v163
		v53 = v161
		goto L21
	} else {
		goto L68
	}
L68:
	;
	v324 = v161
	goto L14
L69:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v180 = v22 + v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v185 = v31
	v187 = v39
	goto L70
L70:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(23)+v185>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v203)>>(uint(v185&int32(7))%32))&int32(1) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v324 = v317
	goto L14
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v185<<(uint(int32(2))%32)))) = int32(0)
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v40))) = uint8(v217)
	v355 = v185 + v217
	v357 = v187
	goto L5
L73:
	;
	goto L74
L74:
	;
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v40))) = uint8(v221)
	v227 = v181 + int32(20) + v185<<(uint(int32(4))%32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v221 <= v228 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v251 = v248 + v180
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+6)))
	if v255 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L76:
	;
	v248 = v228
	v250 = v221
	goto L75
L77:
	;
	goto L78
L78:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+12)))
	v237 = (v187 + v231 - int32(1)) & (int32(0) - v231)
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227)+4)))
	if v238 == int32(65535) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if v187 == v237 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v237
	v248 = v237
	v250 = v221
	goto L75
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v187
	v248 = v187
	v250 = v221
	goto L75
L83:
	;
	goto L84
L84:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v180))))
	if v244 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v245 = v187
	goto L87
L86:
	;
	v245 = v237
	goto L87
L87:
	;
	v248 = v245
	v250 = int32(1)
	goto L75
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v185<<(uint(int32(2))%32)))) = v281
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v227)+4)))
	if v283 <= int32(0) {
		goto L99
	} else {
		goto L100
	}
L89:
	;
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227)+4)))
	switch v258 - int32(1) {
	case 0:
		goto L95
	case 1:
		goto L94
	default:
		goto L92
	case 3:
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v281 = v251
	goto L88
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L43
	} else {
		goto L96
	}
L93:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v281 = v263
	goto L88
L94:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v251))))
	v281 = v262
	goto L88
L95:
	;
	v261 = int32(*(*int8)(unsafe.Add(mBase, uint32(v251))))
	v281 = v261
	goto L88
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = base.I32_extend16_s(v258)
	F_errmsg_internal(m, int32(493552), v19+int32(32))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L43
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(333866), int32(70), int32(69245))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L43
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	if v283 == int32(-1) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v317 = v283 + v248
	if v250 != 0 {
		v351 = v317
		goto L11
	} else {
		goto L118
	}
L102:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v288 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v313 = F_strlen(m, v251)
	mBase = m.M
	v351 = v313 + v248 + int32(1)
	goto L11
L105:
	;
	v291 = int32(6)
	v293 = int32(18)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	if v295 == v293 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	if v288&int32(1) == int32(0) {
		goto L13
	} else {
		goto L117
	}
L108:
	;
	v298 = v293
	goto L110
L109:
	;
	v298 = int32(2)
	goto L110
L110:
	;
	if v295&int32(254) == int32(2) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v303 = v291
	goto L113
L112:
	;
	v303 = v298
	goto L113
L113:
	;
	if v295 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v306 = v291
	goto L116
L115:
	;
	v306 = v303
	goto L116
L116:
	;
	v347 = v306
	goto L12
L117:
	;
	v347 = int32(base.Ui32(v288) >> (uint(int32(1)) % 32))
	goto L12
L118:
	;
	v319 = v185 + int32(1)
	if v319 != v27 {
		v185 = v319
		v187 = v317
		goto L70
	} else {
		goto L119
	}
L119:
	;
	goto L71
L120:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+22)))
	v378 = v374 + v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v382 = v355
	v384 = v357
	goto L123
L121:
	;
	v515 = v355
	v517 = v357
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v517
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v515)
	v532 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v534 = v532 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v534)
	goto L4
L123:
	;
	if v30 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v515 = v27
	v517 = v507
	goto L122
L125:
	;
	v512 = v382 + int32(1)
	if v512 != v27 {
		v382 = v512
		v384 = v507
		goto L123
	} else {
		goto L166
	}
L126:
	;
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v382+v379))) = uint8(v417)
	v421 = v371 + int32(20) + v382<<(uint(int32(4))%32)
	v422 = int32(*(*int16)(unsafe.Add(mBase, uint32(v421)+4)))
	v423 = int32(65535)
	v424 = v422 & v423
	if v424 == v423 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+int32(23)+v382>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v402)>>(uint(v382&int32(7))%32))&int32(1) != 0 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380+v382<<(uint(int32(2))%32)))) = int32(0)
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v382+v379))) = uint8(v414)
	v507 = v384
	goto L125
L129:
	;
	v438 = v436 + v378
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421)+6)))
	if v442 == int32(1) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+v378))))
	if v428 != 0 {
		v436 = v384
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421)+12)))
	v436 = (v384 + v429 - int32(1)) & (int32(0) - v429)
	goto L129
L133:
	;
	goto L132
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380+v382<<(uint(int32(2))%32)))) = v465
	v467 = int32(*(*int16)(unsafe.Add(mBase, uint32(v421)+4)))
	if int32(0) < v467 {
		goto L145
	} else {
		goto L146
	}
L135:
	;
	switch v424 - int32(1) {
	case 0:
		goto L141
	case 1:
		goto L140
	default:
		goto L138
	case 3:
		goto L139
	}
L136:
	;
	goto L137
L137:
	;
	v465 = v438
	goto L134
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L43
	} else {
		goto L142
	}
L139:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	v465 = v449
	goto L134
L140:
	;
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v438))))
	v465 = v448
	goto L134
L141:
	;
	v447 = int32(*(*int8)(unsafe.Add(mBase, uint32(v438))))
	v465 = v447
	goto L134
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v422
	F_errmsg_internal(m, int32(493552), v19+int32(16))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L43
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(333866), int32(70), int32(69245))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L43
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v507 = v467 + v436
	goto L125
L146:
	;
	goto L147
L147:
	;
	if v467 == int32(-1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v473 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v502 = F_strlen(m, v438)
	mBase = m.M
	v507 = v502 + v436 + int32(1)
	goto L125
L151:
	;
	v476 = int32(6)
	v478 = int32(18)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
	if v480 == v478 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	if v473&int32(1) != 0 {
		goto L163
	} else {
		goto L164
	}
L154:
	;
	v483 = v478
	goto L156
L155:
	;
	v483 = int32(2)
	goto L156
L156:
	;
	if v480&int32(254) == int32(2) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v488 = v476
	goto L159
L158:
	;
	v488 = v483
	goto L159
L159:
	;
	if v480 == int32(1) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v491 = v476
	goto L162
L161:
	;
	v491 = v488
	goto L162
L162:
	;
	v507 = v491 + v436
	goto L125
L163:
	;
	v507 = int32(base.Ui32(v473)>>(uint(int32(1))%32)) + v436
	goto L125
L164:
	;
	goto L165
L165:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	v507 = int32(base.Ui32(v498)>>(uint(int32(2))%32)) + v436
	goto L125
L166:
	;
	goto L124
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
				F_errmsg(m, int32(61544), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(505602), int32(369), int32(210341))
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
			v10 = int32(4536272)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[10])) = v13
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
				*(*int32)(unsafe.Add(mBase, _consts[10])) = v11
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
	v4 = int32(4536272)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[10])) = v5
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
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
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
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
		v110 = v22
		v112 = v27
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v110 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L6:
	;
	v114 = v24 + int32(1)
	if v114 < v112 {
		v22 = v110
		v24 = v114
		v27 = v112
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v24))))
	if v34 != 0 {
		v110 = v22
		v112 = v27
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
	v110 = v104 + v105
	v112 = v27
	goto L6
L10:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v92 = int32(1)
	v96 = (v22 + v90 - v92) & (int32(0) - v90)
	if v43&v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if base.Ui32((v76-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v104 = v52
		v105 = int32(6)
		goto L9
	} else {
		goto L20
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
		v104 = v69
		v105 = v40
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
	v110 = v59 + v52
	v112 = v62
	goto L6
L19:
	;
	v72 = F_strlen(m, v39)
	mBase = m.M
	v104 = v69
	v105 = v72 + int32(1)
	goto L9
L20:
	;
	v83 = int32(18)
	if v76&int32(255) == v83 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v89 = v83
	goto L23
L22:
	;
	v89 = int32(2)
	goto L23
L23:
	;
	v104 = v52
	v105 = v89
	goto L9
L24:
	;
	v104 = v96
	v105 = int32(base.Ui32(v43) >> (uint(int32(1)) % 32))
	goto L9
L25:
	;
	goto L26
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v104 = v96
	v105 = int32(base.Ui32(v101) >> (uint(int32(2)) % 32))
	goto L9
L27:
	;
	goto L5
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v119 = F_MemoryContextAlloc(m, v118, v110)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v119
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v124 = v122 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v126 <= int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v135 = v119
	v137 = int32(0)
	goto L31
L31:
	;
	v143 = v13 + int32(20) + v137<<(uint(int32(4))%32)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+6)))
	if v144 != 0 {
		v239 = v135
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L1
L33:
	;
	v243 = v137 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v243 < v244 {
		v135 = v239
		v137 = v243
		goto L31
	} else {
		goto L58
	}
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v137))))
	if v147 != 0 {
		v239 = v135
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v149 = v137 << (uint(int32(2)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149+v150)))
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+4)))
	if v153 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v230 != 0 {
		goto L55
	} else {
		goto L56
	}
L37:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+12)))
	v216 = int32(1)
	v220 = (v135 + v214 - v216) & (int32(0) - v214)
	if v156&v216 != 0 {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+12)))
	v194 = int32(1)
	v198 = (v135 + v192 - v194) & (int32(0) - v192)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	if base.Ui32((v200-v194)&int32(255)) < base.Ui32(int32(3)) {
		v228 = v198
		v230 = int32(6)
		goto L36
	} else {
		goto L47
	}
L39:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v156 != int32(1) {
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+12)))
	v184 = int32(0)
	v186 = (v135 + v180 - int32(1)) & (v184 - v180)
	if v184 < v153 {
		v228 = v186
		v230 = v153
		goto L36
	} else {
		goto L46
	}
L42:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	if v159&int32(254) != int32(2) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v152)+2))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+12)))
	v171 = (v135 + v165 - int32(1)) & (int32(0) - v165)
	v172 = F_EOH_get_flat_size(m, v164)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	F_EOH_flatten_into(m, v164, v171, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v176+v149))) = v171
	v239 = v171 + v172
	goto L33
L46:
	;
	v189 = F_strlen(m, v152)
	mBase = m.M
	v228 = v186
	v230 = v189 + int32(1)
	goto L36
L47:
	;
	v207 = int32(18)
	if v200&int32(255) == v207 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v213 = v207
	goto L50
L49:
	;
	v213 = int32(2)
	goto L50
L50:
	;
	v228 = v198
	v230 = v213
	goto L36
L51:
	;
	v228 = v220
	v230 = int32(base.Ui32(v156) >> (uint(int32(1)) % 32))
	goto L36
L52:
	;
	goto L53
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v228 = v220
	v230 = int32(base.Ui32(v225) >> (uint(int32(2)) % 32))
	goto L36
L54:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v233+v149))) = v232
	v239 = v232 + v230
	goto L33
L55:
	;
	v231 = F__emscripten_memcpy_bulkmem(m, v228, v152, v230)
	mBase = m.M
	v232 = v231
	goto L57
L56:
	;
	v232 = v228
	goto L57
L57:
	;
	goto L54
L58:
	;
	goto L32
}
