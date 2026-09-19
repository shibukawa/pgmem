package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14012(m, l0, int32(_a_F_tts_buffer_is_current_xact_tuple_0), int32(796))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
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
			v10 = int32(_a_F_tts_heap_copy_minimal_tuple_0)
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_tts_heap_copy_minimal_tuple[0]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_tts_heap_copy_minimal_tuple[0])) = v13
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
				*(*int32)(unsafe.Add(mBase, _c_F_tts_heap_copy_minimal_tuple[0])) = v11
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
	var v52 int32
	_ = v52
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
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v196 int32
	_ = v196
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
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
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v555 int32
	_ = v555
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
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errfinish(m, int32(_a_F_tts_heap_getsomeattrs_0), int32(70), int32(_a_F_tts_heap_getsomeattrs_1))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L44
	} else {
		goto L156
	}
L5:
	;
	m.G0 = v19 + int32(48)
	return
L6:
	;
	if v342 < v27 {
		goto L113
	} else {
		goto L114
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v30 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v39 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v36&int32(8) != 0 {
		v342 = v31
		v343 = v35
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v39 = v35
	goto L7
L12:
	;
	v342 = v178 + int32(1)
	v343 = v338
	goto L6
L13:
	;
	v338 = v333 + v241
	goto L12
L14:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v333 = int32(base.Ui32(v329) >> (uint(int32(2)) % 32))
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v309
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v27)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v327 = v325 & int32(_a_F_tts_heap_getsomeattrs_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v327)
	goto L5
L16:
	;
	if v27 <= v31 {
		v309 = v39
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v27 <= v31 {
		v309 = v39
		goto L15
	} else {
		goto L66
	}
L19:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v46 = v22 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = v31
	v52 = v39
	goto L22
L20:
	;
	v342 = v51 + int32(1)
	v343 = v166
	goto L6
L21:
	;
	v166 = v162 + v95
	goto L20
L22:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v40))) = uint8(v67)
	v72 = v47 + int32(20) + v51<<(uint(int32(4))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v67 <= v73 {
		v95 = v73
		v96 = v67
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v162 = int32(base.Ui32(v158) >> (uint(int32(2)) % 32))
	goto L21
L24:
	;
	v97 = v95 + v46
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+6)))
	if v101 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+12)))
	v82 = (v52 + v76 - int32(1)) & (int32(0) - v76)
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v83 == int32(_a_F_tts_heap_getsomeattrs_3) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v95 = v93
	v96 = int32(0)
	goto L24
L27:
	;
	if v52 == v82 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v82
	v93 = v82
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v52
	v93 = v52
	goto L26
L31:
	;
	goto L32
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v46))))
	if v89 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = v52
	goto L35
L34:
	;
	v90 = v82
	goto L35
L35:
	;
	v95 = v90
	v96 = int32(1)
	goto L24
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v51<<(uint(int32(2))%32)))) = v120
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v122 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L37:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	switch v104 - int32(1) {
	case 0:
		goto L43
	case 1:
		goto L42
	default:
		goto L40
	case 3:
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v120 = v97
	goto L36
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v120 = v109
	goto L36
L42:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97))))
	v120 = v108
	goto L36
L43:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	v120 = v107
	goto L36
L44:
	;
	return
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = base.I32_extend16_s(v104)
	F_errmsg_internal(m, int32(_a_F_tts_heap_getsomeattrs_4), v19)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	goto L4
L47:
	;
	goto L23
L48:
	;
	if v122 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v154 = v95 + v122
	if v96 != 0 {
		v166 = v154
		goto L20
	} else {
		goto L64
	}
L51:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v127 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v150 = F_strlen(m, v97)
	mBase = m.M
	v166 = v150 + v95 + int32(1)
	goto L20
L54:
	;
	v131 = int32(18)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v133 == v131 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v127&int32(1) == int32(0) {
		goto L47
	} else {
		goto L63
	}
L57:
	;
	v136 = v131
	goto L59
L58:
	;
	v136 = int32(2)
	goto L59
L59:
	;
	if base.Ui32((v133-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v143 = int32(6)
	goto L62
L61:
	;
	v143 = v136
	goto L62
L62:
	;
	v162 = v143
	goto L21
L63:
	;
	v162 = int32(base.Ui32(v127) >> (uint(int32(1)) % 32))
	goto L21
L64:
	;
	v156 = v51 + int32(1)
	if v156 != v27 {
		v51 = v156
		v52 = v154
		goto L22
	} else {
		goto L65
	}
L65:
	;
	v309 = v154
	goto L15
L66:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v173 = v22 + v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v178 = v31
	v179 = v39
	goto L67
L67:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(23)+v178>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v196)>>(uint(v178&int32(7))%32))&int32(1) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v309 = v303
	goto L15
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v178<<(uint(int32(2))%32)))) = int32(0)
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v178+v40))) = uint8(v210)
	v342 = v178 + v210
	v343 = v179
	goto L6
L70:
	;
	goto L71
L71:
	;
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v178+v40))) = uint8(v214)
	v220 = v174 + int32(20) + v178<<(uint(int32(4))%32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v214 <= v221 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v244 = v241 + v173
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+6)))
	if v248 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L73:
	;
	v241 = v221
	v243 = v214
	goto L72
L74:
	;
	goto L75
L75:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+12)))
	v230 = (v179 + v224 - int32(1)) & (int32(0) - v224)
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	if v231 == int32(_a_F_tts_heap_getsomeattrs_3) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v179 == v230 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v230
	v241 = v230
	v243 = v214
	goto L72
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v179
	v241 = v179
	v243 = v214
	goto L72
L80:
	;
	goto L81
L81:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v173))))
	if v237 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v238 = v179
	goto L84
L83:
	;
	v238 = v230
	goto L84
L84:
	;
	v241 = v238
	v243 = int32(1)
	goto L72
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v178<<(uint(int32(2))%32)))) = v269
	v271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220)+4)))
	if v271 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	switch v251 - int32(1) {
	case 0:
		goto L92
	case 1:
		goto L91
	default:
		goto L89
	case 3:
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v269 = v244
	goto L85
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L44
	} else {
		goto L93
	}
L90:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v269 = v256
	goto L85
L91:
	;
	v255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244))))
	v269 = v255
	goto L85
L92:
	;
	v254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v244))))
	v269 = v254
	goto L85
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = base.I32_extend16_s(v251)
	F_errmsg_internal(m, int32(_a_F_tts_heap_getsomeattrs_4), v19+int32(32))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L44
	} else {
		goto L94
	}
L94:
	;
	goto L4
L95:
	;
	if v271 == int32(-1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v303 = v241 + v271
	if v243 != 0 {
		v338 = v303
		goto L12
	} else {
		goto L111
	}
L98:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v276 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v299 = F_strlen(m, v244)
	mBase = m.M
	v338 = v299 + v241 + int32(1)
	goto L12
L101:
	;
	v280 = int32(18)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	if v282 == v280 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	if v276&int32(1) == int32(0) {
		goto L14
	} else {
		goto L110
	}
L104:
	;
	v285 = v280
	goto L106
L105:
	;
	v285 = int32(2)
	goto L106
L106:
	;
	if base.Ui32((v282-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v292 = int32(6)
	goto L109
L108:
	;
	v292 = v285
	goto L109
L109:
	;
	v333 = v292
	goto L13
L110:
	;
	v333 = int32(base.Ui32(v276) >> (uint(int32(1)) % 32))
	goto L13
L111:
	;
	v305 = v178 + int32(1)
	if v305 != v27 {
		v178 = v305
		v179 = v303
		goto L67
	} else {
		goto L112
	}
L112:
	;
	goto L68
L113:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+22)))
	v365 = v361 + v364
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v369 = v342
	v370 = v343
	goto L116
L114:
	;
	v495 = v342
	v496 = v343
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v496
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v495)
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v514 = v512 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v514)
	goto L5
L116:
	;
	if v30 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v495 = v27
	v496 = v486
	goto L115
L118:
	;
	v492 = v369 + int32(1)
	if v492 != v27 {
		v369 = v492
		v370 = v486
		goto L116
	} else {
		goto L155
	}
L119:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v369+v366))) = uint8(v404)
	v408 = v358 + int32(20) + v369<<(uint(int32(4))%32)
	v409 = int32(*(*int16)(unsafe.Add(mBase, uint32(v408)+4)))
	v410 = int32(_a_F_tts_heap_getsomeattrs_3)
	v411 = v409 & v410
	if v411 == v410 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+int32(23)+v369>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v389)>>(uint(v369&int32(7))%32))&int32(1) != 0 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367+v369<<(uint(int32(2))%32)))) = int32(0)
	v401 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v369+v366))) = uint8(v401)
	v486 = v370
	goto L118
L122:
	;
	v425 = v423 + v365
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+6)))
	if v429 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L123:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v365))))
	if v415 != 0 {
		v423 = v370
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+12)))
	v423 = (v370 + v416 - int32(1)) & (int32(0) - v416)
	goto L122
L126:
	;
	goto L125
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367+v369<<(uint(int32(2))%32)))) = v447
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v408)+4)))
	if int32(0) < v449 {
		goto L137
	} else {
		goto L138
	}
L128:
	;
	switch v411 - int32(1) {
	case 0:
		goto L134
	case 1:
		goto L133
	default:
		goto L131
	case 3:
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v447 = v425
	goto L127
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L44
	} else {
		goto L135
	}
L132:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v447 = v436
	goto L127
L133:
	;
	v435 = int32(*(*int16)(unsafe.Add(mBase, uint32(v425))))
	v447 = v435
	goto L127
L134:
	;
	v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v425))))
	v447 = v434
	goto L127
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v409
	F_errmsg_internal(m, int32(_a_F_tts_heap_getsomeattrs_4), v19+int32(16))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L44
	} else {
		goto L136
	}
L136:
	;
	goto L4
L137:
	;
	v486 = v423 + v449
	goto L118
L138:
	;
	goto L139
L139:
	;
	if v449 == int32(-1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425))))
	if v455 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	v482 = F_strlen(m, v425)
	mBase = m.M
	v486 = v482 + v423 + int32(1)
	goto L118
L143:
	;
	v459 = int32(18)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+1)))
	if v461 == v459 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	if v455&int32(1) != 0 {
		goto L152
	} else {
		goto L153
	}
L146:
	;
	v464 = v459
	goto L148
L147:
	;
	v464 = int32(2)
	goto L148
L148:
	;
	if base.Ui32((v461-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v471 = int32(6)
	goto L151
L150:
	;
	v471 = v464
	goto L151
L151:
	;
	v486 = v423 + v471
	goto L118
L152:
	;
	v486 = int32(base.Ui32(v455)>>(uint(int32(1))%32)) + v423
	goto L118
L153:
	;
	goto L154
L154:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v486 = int32(base.Ui32(v478)>>(uint(int32(2))%32)) + v423
	goto L118
L155:
	;
	goto L117
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tts_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14011(m, l0, l1, l2, int32(_a_F_tts_heap_getsysattr_0), int32(369))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
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
			*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v47 + int32(24)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			if v61 != 0 {
				base.MemoryCopy(m, v47+int32(32), v41, v61)
			} else {
			}
			v65 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v65)
			v67 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v67
			*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v67
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
				*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v47 + int32(24)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v61 != 0 {
					base.MemoryCopy(m, v47+int32(32), v41, v61)
				} else {
				}
				v65 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v65)
				v67 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v67
				*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v67
				return v47
			}
		} else {
			v10 = int32(_a_F_tts_minimal_copy_heap_tuple_0)
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copy_heap_tuple[0]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copy_heap_tuple[0])) = v13
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
				*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copy_heap_tuple[0])) = v11
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
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v47 + int32(24)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					if v61 != 0 {
						base.MemoryCopy(m, v47+int32(32), v41, v61)
					} else {
					}
					v65 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v65)
					v67 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v67
					*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v67
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
	v4 = int32(_a_F_tts_minimal_copyslot_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copyslot[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copyslot[0])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copyslot[0])) = v5
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
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
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
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if int32(0) < v137 {
					v143 = int32(0)
					for {
						v147 = v143 << (uint(int32(2)) % 32)
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147)))
						*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v143))))
						*(*uint8)(unsafe.Add(mBase, uint32(v154+v143))) = uint8(v158)
						v161 = v143 + int32(1)
						v162 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v161 < v162 {
							v143 = v161
							continue
						} else {
							break
						}
						break
					}
					v167 = v162
				} else {
					v167 = v137
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v167)
				v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v172 = v170 & int32(_a_F_tts_virtual_copyslot_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v172)
				F_tts_virtual_materialize(m, l0)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
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
							v114 = v29 - v36
							v116 = v114 << (uint(int32(2)) % 32)
							if v116 != 0 {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								base.MemoryFill(m, v117+v36<<(uint(int32(2))%32), int32(0), v116)
							} else {
							}
							if v114 == int32(0) {
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								base.MemoryFill(m, v125+v36, int32(1), v114)
							}
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
							if v45 == int32(0) {
								v114 = v29 - v36
								v116 = v114 << (uint(int32(2)) % 32)
								if v116 != 0 {
									v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
									base.MemoryFill(m, v117+v36<<(uint(int32(2))%32), int32(0), v116)
								} else {
								}
								if v114 == int32(0) {
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
									base.MemoryFill(m, v125+v36, int32(1), v114)
								}
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
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if int32(0) < v137 {
						v143 = int32(0)
						for {
							v147 = v143 << (uint(int32(2)) % 32)
							v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147)))
							*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v143))))
							*(*uint8)(unsafe.Add(mBase, uint32(v154+v143))) = uint8(v158)
							v161 = v143 + int32(1)
							v162 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v161 < v162 {
								v143 = v161
								continue
							} else {
								break
							}
							break
						}
						v167 = v162
					} else {
						v167 = v137
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v167)
					v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v172 = v170 & int32(_a_F_tts_virtual_copyslot_0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v172)
					F_tts_virtual_materialize(m, l0)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
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
			v137 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if int32(0) < v137 {
				v143 = int32(0)
				for {
					v147 = v143 << (uint(int32(2)) % 32)
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147)))
					*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v143))))
					*(*uint8)(unsafe.Add(mBase, uint32(v154+v143))) = uint8(v158)
					v161 = v143 + int32(1)
					v162 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if v161 < v162 {
						v143 = v161
						continue
					} else {
						break
					}
					break
				}
				v167 = v162
			} else {
				v167 = v137
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v167)
			v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v172 = v170 & int32(_a_F_tts_virtual_copyslot_0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v172)
			F_tts_virtual_materialize(m, l0)
			mBase = m.M
			v175 = m.ExcPending
			if v175 != 0 {
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
						v114 = v29 - v36
						v116 = v114 << (uint(int32(2)) % 32)
						if v116 != 0 {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							base.MemoryFill(m, v117+v36<<(uint(int32(2))%32), int32(0), v116)
						} else {
						}
						if v114 == int32(0) {
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							base.MemoryFill(m, v125+v36, int32(1), v114)
						}
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						if v45 == int32(0) {
							v114 = v29 - v36
							v116 = v114 << (uint(int32(2)) % 32)
							if v116 != 0 {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								base.MemoryFill(m, v117+v36<<(uint(int32(2))%32), int32(0), v116)
							} else {
							}
							if v114 == int32(0) {
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								base.MemoryFill(m, v125+v36, int32(1), v114)
							}
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
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if int32(0) < v137 {
					v143 = int32(0)
					for {
						v147 = v143 << (uint(int32(2)) % 32)
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147)))
						*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v143))))
						*(*uint8)(unsafe.Add(mBase, uint32(v154+v143))) = uint8(v158)
						v161 = v143 + int32(1)
						v162 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v161 < v162 {
							v143 = v161
							continue
						} else {
							break
						}
						break
					}
					v167 = v162
				} else {
					v167 = v137
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v167)
				v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v172 = v170 & int32(_a_F_tts_virtual_copyslot_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v172)
				F_tts_virtual_materialize(m, l0)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
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
	v20 = v2
	v22 = v14
	v25 = v2
	goto L4
L4:
	;
	v30 = v13 + int32(20) + v25<<(uint(int32(4))%32)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+6)))
	if v31 != 0 {
		v107 = v20
		v109 = v22
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v107 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L6:
	;
	v113 = v25 + int32(1)
	if v113 < v109 {
		v20 = v107
		v22 = v109
		v25 = v113
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v25))))
	if v34 != 0 {
		v107 = v20
		v109 = v22
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v25<<(uint(int32(2))%32))))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+4)))
	if v40 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v107 = v102 + v103
	v109 = v22
	goto L6
L10:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v90 = int32(1)
	v94 = (v20 + v88 - v90) & (int32(0) - v88)
	if v43&v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if base.Ui32((v76-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v102 = v52
		v103 = int32(6)
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
	v69 = (v20 + v63 - int32(1)) & (v67 - v63)
	if v67 < v40 {
		v102 = v69
		v103 = v40
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	v52 = (v20 + v46 - int32(1)) & (int32(0) - v46)
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
	v107 = v59 + v52
	v109 = v62
	goto L6
L19:
	;
	v72 = F_strlen(m, v39)
	mBase = m.M
	v102 = v69
	v103 = v72 + int32(1)
	goto L9
L20:
	;
	v83 = int32(18)
	if v76 == v83 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = v83
	goto L23
L22:
	;
	v87 = int32(2)
	goto L23
L23:
	;
	v102 = v52
	v103 = v87
	goto L9
L24:
	;
	v102 = v94
	v103 = int32(base.Ui32(v43) >> (uint(int32(1)) % 32))
	goto L9
L25:
	;
	goto L26
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v102 = v94
	v103 = int32(base.Ui32(v99) >> (uint(int32(2)) % 32))
	goto L9
L27:
	;
	goto L5
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v118 = F_MemoryContextAlloc(m, v117, v107)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v118
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v123 = v121 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v123)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v125 <= int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v132 = v118
	v137 = int32(0)
	goto L31
L31:
	;
	v142 = v13 + int32(20) + v137<<(uint(int32(4))%32)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+6)))
	if v143 != 0 {
		v233 = v132
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L1
L33:
	;
	v239 = v137 + int32(1)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v239 < v240 {
		v132 = v233
		v137 = v239
		goto L31
	} else {
		goto L57
	}
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v137))))
	if v146 != 0 {
		v233 = v132
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v148 = v137 << (uint(int32(2)) % 32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148+v149)))
	v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+4)))
	if v152 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v226 != 0 {
		goto L54
	} else {
		goto L55
	}
L37:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+12)))
	v213 = int32(1)
	v217 = (v132 + v211 - v213) & (int32(0) - v211)
	if v155&v213 != 0 {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+12)))
	v193 = int32(1)
	v197 = (v132 + v191 - v193) & (int32(0) - v191)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if base.Ui32((v199-v193)&int32(255)) < base.Ui32(int32(3)) {
		v225 = v197
		v226 = int32(6)
		goto L36
	} else {
		goto L47
	}
L39:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 != int32(1) {
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+12)))
	v183 = int32(0)
	v185 = (v132 + v179 - int32(1)) & (v183 - v179)
	if v183 < v152 {
		v225 = v185
		v226 = v152
		goto L36
	} else {
		goto L46
	}
L42:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v158&int32(254) != int32(2) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v151)+2))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+12)))
	v170 = (v132 + v164 - int32(1)) & (int32(0) - v164)
	v171 = F_EOH_get_flat_size(m, v163)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	F_EOH_flatten_into(m, v163, v170, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v175+v148))) = v170
	v233 = v170 + v171
	goto L33
L46:
	;
	v188 = F_strlen(m, v151)
	mBase = m.M
	v225 = v185
	v226 = v188 + int32(1)
	goto L36
L47:
	;
	v206 = int32(18)
	if v199 == v206 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v210 = v206
	goto L50
L49:
	;
	v210 = int32(2)
	goto L50
L50:
	;
	v225 = v197
	v226 = v210
	goto L36
L51:
	;
	v225 = v217
	v226 = int32(base.Ui32(v155) >> (uint(int32(1)) % 32))
	goto L36
L52:
	;
	goto L53
L53:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v225 = v217
	v226 = int32(base.Ui32(v222) >> (uint(int32(2)) % 32))
	goto L36
L54:
	;
	base.MemoryCopy(m, v225, v151, v226)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v148))) = v225
	v233 = v225 + v226
	goto L33
L57:
	;
	goto L32
}
