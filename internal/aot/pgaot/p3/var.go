package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int64
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422+l2))) = v420
	return
L2:
	;
	v420 = v417
	v422 = int32(8)
	goto L1
L3:
	;
	v417 = int32(_a_F_add_var_0)
	goto L2
L4:
	;
	if v5 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if v5 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L7:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = int32(0)
	if base.B2i32(v19 < v16)&base.B2i32(v20 < v15) == v20 {
		v51 = v16
		v55 = v20
		goto L17
	} else {
		goto L18
	}
L10:
	;
	return
L11:
	;
	v417 = int32(0)
	goto L2
L12:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L66
	}
L13:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L10
	} else {
		goto L65
	}
L14:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v194 != 0 {
		goto L58
	} else {
		goto L59
	}
L15:
	;
	switch v193 {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L12
	}
L16:
	;
	v193 = v183
	goto L15
L17:
	;
	if base.B2i32(v18 <= int32(0))|base.B2i32(v19 <= v51) != 0 {
		v87 = v19
		v89 = v20
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v32 = v16
	v36 = v20
	goto L19
L19:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v36<<(uint(int32(1))%32)))))
	if v42 != 0 {
		v183 = int32(1)
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v51 = v46
	v55 = v44
	goto L17
L21:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v32 - v43
	if v46 <= v19 {
		v51 = v46
		v55 = v44
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if v44 < v15 {
		v32 = v46
		v36 = v44
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if v51 != v87 {
		v129 = v55
		v130 = v89
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v68 = v19
	v70 = v20
	goto L26
L26:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v70<<(uint(int32(1))%32)))))
	if v75 != 0 {
		v183 = int32(-1)
		goto L16
	} else {
		goto L28
	}
L27:
	;
	v87 = v79
	v89 = v77
	goto L24
L28:
	;
	v76 = int32(1)
	v77 = v70 + v76
	v79 = v68 - v76
	if v79 <= v51 {
		v87 = v79
		v89 = v77
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if v77 < v18 {
		v68 = v79
		v70 = v77
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	if v15 < v129 {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v98 = v55
	v99 = v89
	goto L33
L33:
	;
	if base.B2i32(v15 <= v98)|base.B2i32(v18 <= v99) != 0 {
		v129 = v98
		v130 = v99
		goto L31
	} else {
		goto L35
	}
L34:
	;
	if base.I32_extend16_s(v115) < base.I32_extend16_s(v113) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v104 = int32(1)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v98<<(uint(v104)%32)))))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99<<(uint(v104)%32)+v17))))
	if v113 == v115 {
		v98 = v98 + v104
		v99 = v99 + v104
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v122 = int32(1)
	goto L39
L38:
	;
	v122 = int32(-1)
	goto L39
L39:
	;
	v193 = v122
	goto L15
L40:
	;
	v133 = v129
	goto L42
L41:
	;
	v133 = v15
	goto L42
L42:
	;
	v140 = v129
	goto L43
L43:
	;
	if v133 == v140 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v183 = v166
	goto L16
L45:
	;
	if v18 < v130 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v166 = int32(1)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v140<<(uint(v166)%32)))))
	if v172 == int32(0) {
		v140 = v140 + v166
		goto L43
	} else {
		goto L57
	}
L48:
	;
	v145 = v130
	goto L50
L49:
	;
	v145 = v18
	goto L50
L50:
	;
	v153 = v130
	goto L51
L51:
	;
	if v145 == v153 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v183 = int32(-1)
	goto L16
L53:
	;
	v193 = int32(0)
	goto L15
L54:
	;
	goto L55
L55:
	;
	v157 = int32(1)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153<<(uint(v157)%32)+v17))))
	if v162 == int32(0) {
		v153 = v153 + v157
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	goto L44
L58:
	;
	F_pfree(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v199 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v199
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v204 < v203 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v206 = v203
	goto L64
L63:
	;
	v206 = v204
	goto L64
L64:
	;
	v420 = v206
	v422 = int32(12)
	goto L1
L65:
	;
	v417 = int32(0)
	goto L2
L66:
	;
	goto L3
L67:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v221 = int32(0)
	if base.B2i32(v220 < v217)&base.B2i32(v221 < v216) == v221 {
		v252 = v217
		v256 = v221
		goto L75
	} else {
		goto L76
	}
L68:
	;
	goto L69
L69:
	;
	F_add_abs(m, l0, l1, l2)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L10
	} else {
		goto L125
	}
L70:
	;
	F_sub_abs(m, l1, l0, l2)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L10
	} else {
		goto L124
	}
L71:
	;
	F_sub_abs(m, l0, l1, l2)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L10
	} else {
		goto L123
	}
L72:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v395 != 0 {
		goto L116
	} else {
		goto L117
	}
L73:
	;
	switch v394 {
	case 0:
		goto L72
	case 1:
		goto L71
	default:
		goto L70
	}
L74:
	;
	v394 = v384
	goto L73
L75:
	;
	if base.B2i32(v219 <= int32(0))|base.B2i32(v220 <= v252) != 0 {
		v288 = v220
		v290 = v221
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v233 = v217
	v237 = v221
	goto L77
L77:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215+v237<<(uint(int32(1))%32)))))
	if v243 != 0 {
		v384 = int32(1)
		goto L74
	} else {
		goto L79
	}
L78:
	;
	v252 = v247
	v256 = v245
	goto L75
L79:
	;
	v244 = int32(1)
	v245 = v237 + v244
	v247 = v233 - v244
	if v247 <= v220 {
		v252 = v247
		v256 = v245
		goto L75
	} else {
		goto L80
	}
L80:
	;
	if v245 < v216 {
		v233 = v247
		v237 = v245
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	if v252 != v288 {
		v330 = v256
		v331 = v290
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v269 = v220
	v271 = v221
	goto L84
L84:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218+v271<<(uint(int32(1))%32)))))
	if v276 != 0 {
		v384 = int32(-1)
		goto L74
	} else {
		goto L86
	}
L85:
	;
	v288 = v280
	v290 = v278
	goto L82
L86:
	;
	v277 = int32(1)
	v278 = v271 + v277
	v280 = v269 - v277
	if v280 <= v252 {
		v288 = v280
		v290 = v278
		goto L82
	} else {
		goto L87
	}
L87:
	;
	if v278 < v219 {
		v269 = v280
		v271 = v278
		goto L84
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	if v216 < v330 {
		goto L98
	} else {
		goto L99
	}
L90:
	;
	v299 = v256
	v300 = v290
	goto L91
L91:
	;
	if base.B2i32(v216 <= v299)|base.B2i32(v219 <= v300) != 0 {
		v330 = v299
		v331 = v300
		goto L89
	} else {
		goto L93
	}
L92:
	;
	if base.I32_extend16_s(v316) < base.I32_extend16_s(v314) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v305 = int32(1)
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215+v299<<(uint(v305)%32)))))
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300<<(uint(v305)%32)+v218))))
	if v314 == v316 {
		v299 = v299 + v305
		v300 = v300 + v305
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v323 = int32(1)
	goto L97
L96:
	;
	v323 = int32(-1)
	goto L97
L97:
	;
	v394 = v323
	goto L73
L98:
	;
	v334 = v330
	goto L100
L99:
	;
	v334 = v216
	goto L100
L100:
	;
	v341 = v330
	goto L101
L101:
	;
	if v334 == v341 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v384 = v367
	goto L74
L103:
	;
	if v219 < v331 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	v367 = int32(1)
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215+v341<<(uint(v367)%32)))))
	if v373 == int32(0) {
		v341 = v341 + v367
		goto L101
	} else {
		goto L115
	}
L106:
	;
	v346 = v331
	goto L108
L107:
	;
	v346 = v219
	goto L108
L108:
	;
	v354 = v331
	goto L109
L109:
	;
	if v346 == v354 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v384 = int32(-1)
	goto L74
L111:
	;
	v394 = int32(0)
	goto L73
L112:
	;
	goto L113
L113:
	;
	v358 = int32(1)
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354<<(uint(v358)%32)+v218))))
	if v363 == int32(0) {
		v354 = v354 + v358
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	goto L102
L116:
	;
	F_pfree(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L10
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v400 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v400
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v400
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v405 < v404 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L118
L120:
	;
	v407 = v404
	goto L122
L121:
	;
	v407 = v405
	goto L122
L122:
	;
	v420 = v407
	v422 = int32(12)
	goto L1
L123:
	;
	goto L3
L124:
	;
	v417 = int32(0)
	goto L2
L125:
	;
	goto L3
}
func F_pull_var_clause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
	v13 = F_pull_var_clause_walker(m, l0, v6+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		m.G0 = v6 + int32(16)
		return v17
	}
}
