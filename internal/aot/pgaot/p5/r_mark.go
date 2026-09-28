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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5-int32(3) <= v4 {
		v29 = v2
		return v29
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v5-int32(1)))))
		if base.B2i32(v13 != int32(177))&base.B2i32(v13 != int32(105)) != 0 {
			v29 = v2
			return v29
		} else {
			v22 = F_find_among_b(m, l0, int32(_a_F_r_mark_lArI_0), int32(2), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v29 = base.B2i32(v22 != int32(0))
				return v29
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
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
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
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
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L7
L4:
	;
	return v513
L5:
	;
	if v141 != 0 {
		v513 = v2
		goto L4
	} else {
		goto L28
	}
L6:
	;
	v141 = v134
	goto L5
L7:
	;
	if v25 <= v26 {
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
	v43 = int32(1)
	v44 = v25 - v43
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27+v44))))
	v48 = v46 & int32(255)
	if base.B2i32(v44 == v26)|base.B2i32(int32(0) <= v46) != 0 {
		v106 = v48
		v110 = v43
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if int32(305) < v106 {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v55 = v48 & int32(63)
	v57 = v25 - int32(2)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v57))))
	v61 = v59 << (uint(int32(6)) % 32)
	if base.B2i32(v57 != v26)&base.B2i32(base.Ui32(v59) < base.Ui32(int32(192))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v106 = v61&int32(1984) | v55
	v110 = int32(2)
	goto L10
L13:
	;
	goto L14
L14:
	;
	v74 = v61&int32(4032) | v55
	v76 = v25 - int32(3)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v76))))
	if base.B2i32(v76 != v26)&base.B2i32(base.Ui32(v78) < base.Ui32(int32(224))) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v106 = v78<<(uint(int32(12))%32)&int32(_a_F_r_mark_sU_0) | v74
	v110 = int32(3)
	goto L10
L16:
	;
	goto L17
L17:
	;
	v96 = int32(4)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v27-v96))))
	v106 = v78<<(uint(int32(12))%32)&int32(_a_F_r_mark_sU_1) | v98&int32(7)<<(uint(int32(18))%32) | v74
	v110 = v96
	goto L10
L18:
	;
	v141 = v110
	goto L5
L19:
	;
	goto L20
L20:
	;
	v112 = v106 - int32(105)
	if v112 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v141 = v110
	goto L5
L22:
	;
	goto L23
L23:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v112)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_sU[0]))))
	if int32(base.Ui32(v118)>>(uint(v112&int32(7))%32))&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v141 = v110
	goto L5
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - v110
	goto L27
L27:
	;
	goto L8
L28:
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
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v502
	v513 = v507
	goto L4
L30:
	;
	v502 = v500
	v507 = int32(1)
	goto L29
L31:
	;
	v500 = v154 - v142 + v285
	goto L30
L32:
	;
	v293 = v144 - v142
	v294 = v290 + v293
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	if v294 <= v291 {
		goto L61
	} else {
		goto L62
	}
L33:
	;
	v290 = v142
	v291 = v145
	v292 = v143
	goto L32
L34:
	;
	goto L35
L35:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v144-int32(1)))))
	if v150 != int32(115) {
		v290 = v142
		v291 = v145
		v292 = v143
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v154 = v144 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L39
L37:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v284 == int32(0) {
		goto L31
	} else {
		goto L60
	}
L38:
	;
	v284 = v277
	goto L37
L39:
	;
	if v154 <= v169 {
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
	v186 = int32(1)
	v187 = v154 - v186
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v170+v187))))
	v191 = v189 & int32(255)
	if base.B2i32(v187 == v169)|base.B2i32(int32(0) <= v189) != 0 {
		v249 = v191
		v253 = v186
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if int32(305) < v249 {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v198 = v191 & int32(63)
	v200 = v154 - int32(2)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v200))))
	v204 = v202 << (uint(int32(6)) % 32)
	if base.B2i32(v200 != v169)&base.B2i32(base.Ui32(v202) < base.Ui32(int32(192))) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v249 = v204&int32(1984) | v198
	v253 = int32(2)
	goto L42
L45:
	;
	goto L46
L46:
	;
	v217 = v204&int32(4032) | v198
	v219 = v154 - int32(3)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v219))))
	if base.B2i32(v219 != v169)&base.B2i32(base.Ui32(v221) < base.Ui32(int32(224))) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v249 = v221<<(uint(int32(12))%32)&int32(_a_F_r_mark_sU_0) | v217
	v253 = int32(3)
	goto L42
L48:
	;
	goto L49
L49:
	;
	v239 = int32(4)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v170-v239))))
	v249 = v221<<(uint(int32(12))%32)&int32(_a_F_r_mark_sU_1) | v241&int32(7)<<(uint(int32(18))%32) | v217
	v253 = v239
	goto L42
L50:
	;
	v284 = v253
	goto L37
L51:
	;
	goto L52
L52:
	;
	v255 = v249 - int32(97)
	if v255 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v284 = v253
	goto L37
L54:
	;
	goto L55
L55:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v255)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_sU[1]))))
	if int32(base.Ui32(v261)>>(uint(v255&int32(7))%32))&int32(1) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v284 = v253
	goto L37
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v253
	goto L59
L59:
	;
	goto L40
L60:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v290 = v285
	v291 = v289
	v292 = v288
	goto L32
L61:
	;
	v306 = int32(0)
	goto L66
L62:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v292-int32(1)))))
	if v300 != int32(115) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v502 = v294 - int32(1)
	v507 = int32(0)
	goto L29
L64:
	;
	if v358 < int32(0) {
		v513 = v306
		goto L4
	} else {
		goto L83
	}
L66:
	;
	goto L67
L67:
	;
	goto L68
L68:
	;
	v313 = v294
	v315 = int32(1)
	goto L71
L70:
	;
	v358 = v340
	goto L64
L71:
	;
	if v313 <= v291 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v358 = int32(-1)
	goto L64
L74:
	;
	goto L75
L75:
	;
	v320 = v313 - int32(1)
	v322 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292+v320))))
	if base.B2i32(int32(0) <= v322)|base.B2i32(v320 <= v291) != 0 {
		v340 = v320
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v344 = int32(1)
	if v344 < v315 {
		v313 = v340
		v315 = v315 - v344
		goto L71
	} else {
		goto L82
	}
L77:
	;
	v328 = v320
	goto L78
L78:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v328))))
	if base.Ui32(int32(191)) < base.Ui32(v333) {
		v340 = v328
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v340 = v291
	goto L76
L80:
	;
	v337 = v328 - int32(1)
	if v291 < v337 {
		v328 = v337
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L72
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v358
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L84:
	;
	if v490 != 0 {
		v513 = v306
		goto L4
	} else {
		goto L107
	}
L85:
	;
	v490 = v483
	goto L84
L86:
	;
	if v358 <= v375 {
		v483 = int32(-1)
		goto L85
	} else {
		goto L88
	}
L87:
	;
	v483 = int32(0)
	goto L85
L88:
	;
	v392 = int32(1)
	v393 = v358 - v392
	v395 = int32(*(*int8)(unsafe.Add(mBase, uint32(v376+v393))))
	v397 = v395 & int32(255)
	if base.B2i32(v393 == v375)|base.B2i32(int32(0) <= v395) != 0 {
		v455 = v397
		v459 = v392
		goto L89
	} else {
		goto L90
	}
L89:
	;
	if int32(305) < v455 {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	v404 = v397 & int32(63)
	v406 = v358 - int32(2)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v406))))
	v410 = v408 << (uint(int32(6)) % 32)
	if base.B2i32(v406 != v375)&base.B2i32(base.Ui32(v408) < base.Ui32(int32(192))) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v455 = v410&int32(1984) | v404
	v459 = int32(2)
	goto L89
L92:
	;
	goto L93
L93:
	;
	v423 = v410&int32(4032) | v404
	v425 = v358 - int32(3)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v425))))
	if base.B2i32(v425 != v375)&base.B2i32(base.Ui32(v427) < base.Ui32(int32(224))) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v455 = v427<<(uint(int32(12))%32)&int32(_a_F_r_mark_sU_0) | v423
	v459 = int32(3)
	goto L89
L95:
	;
	goto L96
L96:
	;
	v445 = int32(4)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358+v376-v445))))
	v455 = v427<<(uint(int32(12))%32)&int32(_a_F_r_mark_sU_1) | v447&int32(7)<<(uint(int32(18))%32) | v423
	v459 = v445
	goto L89
L97:
	;
	v490 = v459
	goto L84
L98:
	;
	goto L99
L99:
	;
	v461 = v455 - int32(97)
	if v461 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v490 = v459
	goto L84
L101:
	;
	goto L102
L102:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v461)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_sU[1]))))
	if int32(base.Ui32(v467)>>(uint(v461&int32(7))%32))&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v490 = v459
	goto L84
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v358 - v459
	goto L106
L106:
	;
	goto L87
L107:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v500 = v491 + v293
	goto L30
}
func F_r_mark_sUn(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14363(m, l0, int32(4), int32(_a_F_r_mark_sUn_0), int32(110))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_r_mark_suffix_with_optional_n_consonant(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v2 = int32(110)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 <= v12 {
		v32 = v9
		v33 = v12
		v34 = v10
		v35 = v11 - v9
		v36 = v32 + v35
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
		if v36 <= v33 {
			v47 = int32(0)
			v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
			mBase = m.M
			if v49 < v47 {
				v80 = v47
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
				v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), int32(0))
				mBase = m.M
				if v57 != 0 {
					v80 = v47
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v67 = v58 + v35
					v69 = v67
					v74 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
					v80 = v74
				}
			}
		} else {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
			if v42 != v2 {
				v47 = int32(0)
				v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
				mBase = m.M
				if v49 < v47 {
					v80 = v47
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
					v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), int32(0))
					mBase = m.M
					if v57 != 0 {
						v80 = v47
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v67 = v58 + v35
						v69 = v67
						v74 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
						v80 = v74
					}
				}
			} else {
				v69 = v36 - int32(1)
				v74 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
				v80 = v74
			}
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11-int32(1)))))
		if v17 != v2 {
			v32 = v9
			v33 = v12
			v34 = v10
			v35 = v11 - v9
			v36 = v32 + v35
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
			if v36 <= v33 {
				v47 = int32(0)
				v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
				mBase = m.M
				if v49 < v47 {
					v80 = v47
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
					v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), int32(0))
					mBase = m.M
					if v57 != 0 {
						v80 = v47
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v67 = v58 + v35
						v69 = v67
						v74 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
						v80 = v74
					}
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
				if v42 != v2 {
					v47 = int32(0)
					v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
					mBase = m.M
					if v49 < v47 {
						v80 = v47
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
						v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), int32(0))
						mBase = m.M
						if v57 != 0 {
							v80 = v47
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v67 = v58 + v35
							v69 = v67
							v74 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
							v80 = v74
						}
					}
				} else {
					v69 = v36 - int32(1)
					v74 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
					v80 = v74
				}
			}
		} else {
			v20 = v11 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
			v25 = int32(0)
			v26 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), v25)
			mBase = m.M
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v26 == v25 {
				v67 = v20 - v9 + v27
				v69 = v67
				v74 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
				v80 = v74
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v32 = v27
				v33 = v31
				v34 = v30
				v35 = v11 - v9
				v36 = v32 + v35
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
				if v36 <= v33 {
					v47 = int32(0)
					v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
					mBase = m.M
					if v49 < v47 {
						v80 = v47
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
						v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), int32(0))
						mBase = m.M
						if v57 != 0 {
							v80 = v47
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v67 = v58 + v35
							v69 = v67
							v74 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
							v80 = v74
						}
					}
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
					if v42 != v2 {
						v47 = int32(0)
						v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
						mBase = m.M
						if v49 < v47 {
							v80 = v47
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
							v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_n_consonant_0), int32(97), int32(305), int32(0))
							mBase = m.M
							if v57 != 0 {
								v80 = v47
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v67 = v58 + v35
								v69 = v67
								v74 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
								v80 = v74
							}
						}
					} else {
						v69 = v36 - int32(1)
						v74 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
						v80 = v74
					}
				}
			}
		}
	}
	return v80
}
func F_r_mark_yUz(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14365(m, l0, int32(_a_F_r_mark_yUz_0), int32(122))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
