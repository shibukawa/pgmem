package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14009(m, l0, l1, l2, int32(_a_F_tts_buffer_heap_getsysattr_0), int32(774))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
	F_errfinish(m, int32(_a_F_tts_minimal_getsomeattrs_0), int32(70), int32(_a_F_tts_minimal_getsomeattrs_1))
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v309
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v27)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v327 = v325 & int32(_a_F_tts_minimal_getsomeattrs_2)
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
	if v83 == int32(_a_F_tts_minimal_getsomeattrs_3) {
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
	F_errmsg_internal(m, int32(_a_F_tts_minimal_getsomeattrs_4), v19)
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
	if v231 == int32(_a_F_tts_minimal_getsomeattrs_3) {
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
	F_errmsg_internal(m, int32(_a_F_tts_minimal_getsomeattrs_4), v19+int32(32))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v496
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
	v410 = int32(_a_F_tts_minimal_getsomeattrs_3)
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
	F_errmsg_internal(m, int32(_a_F_tts_minimal_getsomeattrs_4), v19+int32(16))
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
			F_errmsg(m, int32(_a_F_tts_minimal_getsysattr_0), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_tts_minimal_getsysattr_1), int32(564), int32(_a_F_tts_minimal_getsysattr_2))
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
