package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_r_mark_lArI(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5-int32(3) <= v4 {
		v28 = v2
		return v28
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v5-int32(1)))))
		if base.B2i32(v13 != int32(177))&base.B2i32(v13 != int32(105)) != 0 {
			v28 = v2
			return v28
		} else {
			v21 = F_find_among_b(m, l0, int32(4422000), int32(2))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.B2i32(v21 != int32(0))
				return v28
			}
		}
	}
}
func F_r_mark_sU(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v353 int32
	_ = v353
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	v2 = int32(0)
	v8 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v8 == v2 {
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L7
L4:
	;
	return v502
L5:
	;
	if v141 != 0 {
		v502 = v2
		goto L4
	} else {
		goto L29
	}
L6:
	;
	v141 = v134
	goto L5
L7:
	;
	if v29 <= v30 {
		v134 = int32(-1)
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v134 = int32(0)
	goto L6
L9:
	;
	v47 = int32(1)
	v48 = v29 - v47
	v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26+v48))))
	v52 = v50 & int32(255)
	if v48 == v30 {
		v107 = v52
		v108 = v47
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if int32(305) < v107 {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	if int32(0) <= v50 {
		v107 = v52
		v108 = v47
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v58 = v52 & int32(63)
	v60 = v29 - int32(2)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v60))))
	v64 = v62 << (uint(int32(6)) % 32)
	if base.B2i32(v60 != v30)&base.B2i32(base.Ui32(v62) < base.Ui32(int32(192))) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v107 = v64&int32(1984) | v58
	v108 = int32(2)
	goto L10
L14:
	;
	goto L15
L15:
	;
	v77 = v64&int32(4032) | v58
	v79 = v29 - int32(3)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v79))))
	if base.B2i32(v79 != v30)&base.B2i32(base.Ui32(v81) < base.Ui32(int32(224))) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v107 = v81<<(uint(int32(12))%32)&int32(61440) | v77
	v108 = int32(3)
	goto L10
L17:
	;
	goto L18
L18:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+(v26-int32(4))))))
	v107 = v81<<(uint(int32(12))%32)&int32(258048) | v99&int32(7)<<(uint(int32(18))%32) | v77
	v108 = int32(4)
	goto L10
L19:
	;
	v141 = v108
	goto L5
L20:
	;
	goto L21
L21:
	;
	v112 = v107 - int32(105)
	if v112 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v141 = v108
	goto L5
L23:
	;
	goto L24
L24:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v112)>>(uint(int32(3))%32)))+uint32(_consts[1446]))))
	if int32(base.Ui32(v118)>>(uint(v112&int32(7))%32))&int32(1) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v141 = v108
	goto L5
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29 - v108
	goto L28
L28:
	;
	goto L8
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v144 <= v145 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v495
	v502 = int32(1)
	goto L4
L31:
	;
	v495 = v154 - v142 + v285
	goto L30
L32:
	;
	v293 = v144 - v142
	v294 = v292 + v293
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	if v291 < v294 {
		goto L62
	} else {
		goto L63
	}
L33:
	;
	v290 = v143
	v291 = v145
	v292 = v142
	goto L32
L34:
	;
	goto L35
L35:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v144-int32(1)))))
	if v150 != int32(115) {
		v290 = v143
		v291 = v145
		v292 = v142
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v154 = v144 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L39
L37:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v284 == int32(0) {
		goto L31
	} else {
		goto L61
	}
L38:
	;
	v284 = v277
	goto L37
L39:
	;
	if v154 <= v173 {
		v277 = int32(-1)
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v277 = int32(0)
	goto L38
L41:
	;
	v190 = int32(1)
	v191 = v154 - v190
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v169+v191))))
	v195 = v193 & int32(255)
	if v191 == v173 {
		v250 = v195
		v251 = v190
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if int32(305) < v250 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	if int32(0) <= v193 {
		v250 = v195
		v251 = v190
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v201 = v195 & int32(63)
	v203 = v154 - int32(2)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v203))))
	v207 = v205 << (uint(int32(6)) % 32)
	if base.B2i32(v203 != v173)&base.B2i32(base.Ui32(v205) < base.Ui32(int32(192))) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v250 = v207&int32(1984) | v201
	v251 = int32(2)
	goto L42
L46:
	;
	goto L47
L47:
	;
	v220 = v207&int32(4032) | v201
	v222 = v154 - int32(3)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v222))))
	if base.B2i32(v222 != v173)&base.B2i32(base.Ui32(v224) < base.Ui32(int32(224))) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v250 = v224<<(uint(int32(12))%32)&int32(61440) | v220
	v251 = int32(3)
	goto L42
L49:
	;
	goto L50
L50:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+(v169-int32(4))))))
	v250 = v224<<(uint(int32(12))%32)&int32(258048) | v242&int32(7)<<(uint(int32(18))%32) | v220
	v251 = int32(4)
	goto L42
L51:
	;
	v284 = v251
	goto L37
L52:
	;
	goto L53
L53:
	;
	v255 = v250 - int32(97)
	if v255 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v284 = v251
	goto L37
L55:
	;
	goto L56
L56:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v255)>>(uint(int32(3))%32)))+uint32(_consts[1445]))))
	if int32(base.Ui32(v261)>>(uint(v255&int32(7))%32))&int32(1) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v284 = v251
	goto L37
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v251
	goto L60
L60:
	;
	goto L40
L61:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v290 = v288
	v291 = v289
	v292 = v285
	goto L32
L62:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v290-int32(1)))))
	if v300 == int32(115) {
		v502 = v2
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L68
L65:
	;
	goto L64
L66:
	;
	if v353 < int32(0) {
		v502 = v2
		goto L4
	} else {
		goto L86
	}
L68:
	;
	goto L69
L69:
	;
	goto L70
L70:
	;
	v309 = v294
	v311 = int32(1)
	goto L73
L72:
	;
	v353 = v335
	goto L66
L73:
	;
	if v309 <= v291 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L72
L75:
	;
	v353 = int32(-1)
	goto L66
L76:
	;
	goto L77
L77:
	;
	v316 = v309 - int32(1)
	v318 = int32(*(*int8)(unsafe.Add(mBase, uint32(v290+v316))))
	if int32(0) <= v318 {
		v335 = v316
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v339 = int32(1)
	if v339 < v311 {
		v309 = v335
		v311 = v311 - v339
		goto L73
	} else {
		goto L85
	}
L79:
	;
	if v316 <= v291 {
		v335 = v316
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v323 = v316
	goto L81
L81:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v323))))
	if base.Ui32(int32(191)) < base.Ui32(v328) {
		v335 = v323
		goto L78
	} else {
		goto L83
	}
L82:
	;
	v335 = v291
	goto L78
L83:
	;
	v332 = v323 - int32(1)
	if v291 < v332 {
		v323 = v332
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	goto L74
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v353
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L89
L87:
	;
	if v485 != 0 {
		v502 = v2
		goto L4
	} else {
		goto L111
	}
L88:
	;
	v485 = v478
	goto L87
L89:
	;
	if v353 <= v374 {
		v478 = int32(-1)
		goto L88
	} else {
		goto L91
	}
L90:
	;
	v478 = int32(0)
	goto L88
L91:
	;
	v391 = int32(1)
	v392 = v353 - v391
	v394 = int32(*(*int8)(unsafe.Add(mBase, uint32(v370+v392))))
	v396 = v394 & int32(255)
	if v392 == v374 {
		v451 = v396
		v452 = v391
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if int32(305) < v451 {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	if int32(0) <= v394 {
		v451 = v396
		v452 = v391
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v402 = v396 & int32(63)
	v404 = v353 - int32(2)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v404))))
	v408 = v406 << (uint(int32(6)) % 32)
	if base.B2i32(v404 != v374)&base.B2i32(base.Ui32(v406) < base.Ui32(int32(192))) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v451 = v408&int32(1984) | v402
	v452 = int32(2)
	goto L92
L96:
	;
	goto L97
L97:
	;
	v421 = v408&int32(4032) | v402
	v423 = v353 - int32(3)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v423))))
	if base.B2i32(v423 != v374)&base.B2i32(base.Ui32(v425) < base.Ui32(int32(224))) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v451 = v425<<(uint(int32(12))%32)&int32(61440) | v421
	v452 = int32(3)
	goto L92
L99:
	;
	goto L100
L100:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+(v370-int32(4))))))
	v451 = v425<<(uint(int32(12))%32)&int32(258048) | v443&int32(7)<<(uint(int32(18))%32) | v421
	v452 = int32(4)
	goto L92
L101:
	;
	v485 = v452
	goto L87
L102:
	;
	goto L103
L103:
	;
	v456 = v451 - int32(97)
	if v456 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v485 = v452
	goto L87
L105:
	;
	goto L106
L106:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v456)>>(uint(int32(3))%32)))+uint32(_consts[1445]))))
	if int32(base.Ui32(v462)>>(uint(v456&int32(7))%32))&int32(1) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v485 = v452
	goto L87
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v353 - v452
	goto L110
L110:
	;
	goto L90
L111:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v495 = v486 + v293
	goto L30
}
func F_r_mark_sUn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v27 = v2
		return v27
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8-int32(2) <= v7 {
			v27 = v2
			return v27
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v8-int32(1)))))
			if v16 != int32(110) {
				v27 = v2
				return v27
			} else {
				v21 = F_find_among_b(m, l0, int32(4421552), int32(4))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v27 = base.B2i32(v21 != int32(0))
					return v27
				}
			}
		}
	}
}
func F_r_mark_suffix_with_optional_n_consonant(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 <= v11 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v369
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v369 = int32(1)
	goto L1
L3:
	;
	v365 = v20 - v8 + v151
	goto L2
L4:
	;
	v159 = v10 - v8
	v160 = v158 + v159
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160
	if v157 < v160 {
		goto L34
	} else {
		goto L35
	}
L5:
	;
	v156 = v9
	v157 = v11
	v158 = v8
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v9-int32(1)))))
	if v16 != int32(110) {
		v156 = v9
		v157 = v11
		v158 = v8
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v20 = v10 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L11
L9:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v150 == int32(0) {
		goto L3
	} else {
		goto L33
	}
L10:
	;
	v150 = v143
	goto L9
L11:
	;
	if v20 <= v39 {
		v143 = int32(-1)
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v143 = int32(0)
	goto L10
L13:
	;
	v56 = int32(1)
	v57 = v20 - v56
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35+v57))))
	v61 = v59 & int32(255)
	if v57 == v39 {
		v116 = v61
		v117 = v56
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if int32(305) < v116 {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if int32(0) <= v59 {
		v116 = v61
		v117 = v56
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v67 = v61 & int32(63)
	v69 = v20 - int32(2)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v69))))
	v73 = v71 << (uint(int32(6)) % 32)
	if base.B2i32(v69 != v39)&base.B2i32(base.Ui32(v71) < base.Ui32(int32(192))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v116 = v73&int32(1984) | v67
	v117 = int32(2)
	goto L14
L18:
	;
	goto L19
L19:
	;
	v86 = v73&int32(4032) | v67
	v88 = v20 - int32(3)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v88))))
	if base.B2i32(v88 != v39)&base.B2i32(base.Ui32(v90) < base.Ui32(int32(224))) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v116 = v90<<(uint(int32(12))%32)&int32(61440) | v86
	v117 = int32(3)
	goto L14
L21:
	;
	goto L22
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+(v35-int32(4))))))
	v116 = v90<<(uint(int32(12))%32)&int32(258048) | v108&int32(7)<<(uint(int32(18))%32) | v86
	v117 = int32(4)
	goto L14
L23:
	;
	v150 = v117
	goto L9
L24:
	;
	goto L25
L25:
	;
	v121 = v116 - int32(97)
	if v121 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v150 = v117
	goto L9
L27:
	;
	goto L28
L28:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v121)>>(uint(int32(3))%32)))+uint32(_consts[1445]))))
	if int32(base.Ui32(v127)>>(uint(v121&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v150 = v117
	goto L9
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 - v117
	goto L32
L32:
	;
	goto L12
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = v154
	v157 = v155
	v158 = v151
	goto L4
L34:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v156-int32(1)))))
	if v167 == int32(110) {
		v369 = int32(0)
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v171 = int32(0)
	goto L40
L37:
	;
	goto L36
L38:
	;
	if v222 < int32(0) {
		v369 = v171
		goto L1
	} else {
		goto L58
	}
L40:
	;
	goto L41
L41:
	;
	goto L42
L42:
	;
	v178 = v160
	v180 = int32(1)
	goto L45
L44:
	;
	v222 = v204
	goto L38
L45:
	;
	if v178 <= v157 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v222 = int32(-1)
	goto L38
L48:
	;
	goto L49
L49:
	;
	v185 = v178 - int32(1)
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156+v185))))
	if int32(0) <= v187 {
		v204 = v185
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v208 = int32(1)
	if v208 < v180 {
		v178 = v204
		v180 = v180 - v208
		goto L45
	} else {
		goto L57
	}
L51:
	;
	if v185 <= v157 {
		v204 = v185
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v192 = v185
	goto L53
L53:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v192))))
	if base.Ui32(int32(191)) < base.Ui32(v197) {
		v204 = v192
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v204 = v157
	goto L50
L55:
	;
	v201 = v192 - int32(1)
	if v157 < v201 {
		v192 = v201
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L46
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L61
L59:
	;
	if v354 != 0 {
		v369 = v171
		goto L1
	} else {
		goto L83
	}
L60:
	;
	v354 = v347
	goto L59
L61:
	;
	if v222 <= v243 {
		v347 = int32(-1)
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v347 = int32(0)
	goto L60
L63:
	;
	v260 = int32(1)
	v261 = v222 - v260
	v263 = int32(*(*int8)(unsafe.Add(mBase, uint32(v239+v261))))
	v265 = v263 & int32(255)
	if v261 == v243 {
		v320 = v265
		v321 = v260
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if int32(305) < v320 {
		goto L73
	} else {
		goto L74
	}
L65:
	;
	if int32(0) <= v263 {
		v320 = v265
		v321 = v260
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v271 = v265 & int32(63)
	v273 = v222 - int32(2)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v273))))
	v277 = v275 << (uint(int32(6)) % 32)
	if base.B2i32(v273 != v243)&base.B2i32(base.Ui32(v275) < base.Ui32(int32(192))) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v320 = v277&int32(1984) | v271
	v321 = int32(2)
	goto L64
L68:
	;
	goto L69
L69:
	;
	v290 = v277&int32(4032) | v271
	v292 = v222 - int32(3)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v292))))
	if base.B2i32(v292 != v243)&base.B2i32(base.Ui32(v294) < base.Ui32(int32(224))) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v320 = v294<<(uint(int32(12))%32)&int32(61440) | v290
	v321 = int32(3)
	goto L64
L71:
	;
	goto L72
L72:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+(v239-int32(4))))))
	v320 = v294<<(uint(int32(12))%32)&int32(258048) | v312&int32(7)<<(uint(int32(18))%32) | v290
	v321 = int32(4)
	goto L64
L73:
	;
	v354 = v321
	goto L59
L74:
	;
	goto L75
L75:
	;
	v325 = v320 - int32(97)
	if v325 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v354 = v321
	goto L59
L77:
	;
	goto L78
L78:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v325)>>(uint(int32(3))%32)))+uint32(_consts[1445]))))
	if int32(base.Ui32(v331)>>(uint(v325&int32(7))%32))&int32(1) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v354 = v321
	goto L59
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222 - v321
	goto L82
L82:
	;
	goto L62
L83:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = v355 + v159
	goto L2
}
func F_r_mark_yUz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v100 = v2
		return v100
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = v7 - int32(1)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 <= v10 {
			v100 = v2
			return v100
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9))))
			if v14 != int32(122) {
				v100 = v2
				return v100
			} else {
				v19 = F_find_among_b(m, l0, int32(4421632), int32(4))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v100 = v2
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v33 <= v34 {
							v55 = v32
							v56 = v34
							v57 = v31
							v58 = v33 - v31
							v59 = v57 + v58
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
							if v56 < v59 {
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55-int32(1)))))
								if v66 == int32(121) {
									v95 = int32(0)
								} else {
									v70 = int32(0)
									v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
									mBase = m.M
									if v72 < v70 {
										v95 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
										v80 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), int32(0))
										mBase = m.M
										if v80 != 0 {
											v95 = v70
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v91 = v81 + v58
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
											v95 = int32(1)
										}
									}
								}
							} else {
								v70 = int32(0)
								v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
								mBase = m.M
								if v72 < v70 {
									v95 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
									v80 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), int32(0))
									mBase = m.M
									if v80 != 0 {
										v95 = v70
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v91 = v81 + v58
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
										v95 = int32(1)
									}
								}
							}
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v32-int32(1)))))
							if v39 != int32(121) {
								v55 = v32
								v56 = v34
								v57 = v31
								v58 = v33 - v31
								v59 = v57 + v58
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
								if v56 < v59 {
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55-int32(1)))))
									if v66 == int32(121) {
										v95 = int32(0)
									} else {
										v70 = int32(0)
										v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
										mBase = m.M
										if v72 < v70 {
											v95 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
											v80 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), int32(0))
											mBase = m.M
											if v80 != 0 {
												v95 = v70
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v91 = v81 + v58
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
												v95 = int32(1)
											}
										}
									}
								} else {
									v70 = int32(0)
									v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
									mBase = m.M
									if v72 < v70 {
										v95 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
										v80 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), int32(0))
										mBase = m.M
										if v80 != 0 {
											v95 = v70
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v91 = v81 + v58
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
											v95 = int32(1)
										}
									}
								}
							} else {
								v43 = v33 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
								v48 = int32(0)
								v49 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), v48)
								mBase = m.M
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v49 == v48 {
									v91 = v43 - v31 + v50
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
									v95 = int32(1)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v55 = v53
									v56 = v54
									v57 = v50
									v58 = v33 - v31
									v59 = v57 + v58
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
									if v56 < v59 {
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55-int32(1)))))
										if v66 == int32(121) {
											v95 = int32(0)
										} else {
											v70 = int32(0)
											v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
											mBase = m.M
											if v72 < v70 {
												v95 = v70
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
												v80 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), int32(0))
												mBase = m.M
												if v80 != 0 {
													v95 = v70
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v91 = v81 + v58
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
													v95 = int32(1)
												}
											}
										}
									} else {
										v70 = int32(0)
										v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
										mBase = m.M
										if v72 < v70 {
											v95 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
											v80 = F_in_grouping_b_U(m, l0, int32(2263840), int32(97), int32(305), int32(0))
											mBase = m.M
											if v80 != 0 {
												v95 = v70
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v91 = v81 + v58
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
												v95 = int32(1)
											}
										}
									}
								}
							}
						}
						v100 = v95
					}
					return v100
				}
			}
		}
	}
}
