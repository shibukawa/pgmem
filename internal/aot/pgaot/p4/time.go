package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimeCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 float64
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 float64
	_ = v174
	var v177 float64
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 float64
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 float64
	_ = v331
	var v334 float64
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v353 int64
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	v5 = int32(0)
	v8 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(31744)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v5
	v25 = F_strtox_2(m, l0, v12+int32(8), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v28 == int32(68) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v381
L3:
	;
	v381 = int32(-2)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v32 = int32(-1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != int32(58) {
		v381 = v32
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v45 = F_strtol(m, v33+int32(1), v12+int32(8), int32(10))
	mBase = m.M
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v48 == int32(68) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v381 = int32(-2)
	goto L2
L9:
	;
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	switch v53 - int32(46) {
	case 0:
		goto L14
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v381 = v32
		goto L2
	case 12:
		goto L13
	default:
		goto L15
	}
L11:
	;
	if base.Ui32(int32(59)) < base.Ui32(v360) {
		goto L100
	} else {
		goto L101
	}
L12:
	;
	v353 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if v353 < int64(0) {
		goto L97
	} else {
		goto L98
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v205 = F_strtol(m, v52+int32(1), v12+int32(8), int32(10))
	mBase = m.M
	goto L57
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v52
	v73 = v52 + int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v53 != 0 {
		v381 = v32
		goto L2
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(0)
	if l1 != int32(6144) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui64(v60-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v381 = int32(-2)
	goto L2
L19:
	;
	goto L20
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v45
	v69 = base.I32_wrap_i64(v60)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v69
	v360 = v69
	goto L11
L21:
	;
	v75 = int32(552279)
	v79 = m.G0
	v81 = v79 - int32(32)
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v82
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[956])))
	if v90 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v174 = v8
	goto L23
L23:
	;
	v177 = base.F64_nearest(base.F64_mul(v174, float64(1e+06)))
	if base.F64_lt(base.F64_abs(v177), float64(2.147483648e+09)) != 0 {
		goto L51
	} else {
		goto L52
	}
L24:
	;
	v159 = F_strlen(m, v73)
	mBase = m.M
	if v158 != v159 {
		v381 = v32
		goto L2
	} else {
		goto L45
	}
L25:
	;
	v158 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _consts[957])))
	if v94 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v73
	goto L31
L29:
	;
	goto L30
L30:
	;
	v108 = v75
	v109 = v90
	goto L34
L31:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v104 == v90 {
		v98 = v98 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v158 = v98 - v73
	goto L24
L33:
	;
	goto L32
L34:
	;
	v116 = v81 + int32(base.Ui32(v109)>>(uint(int32(3))%32))&int32(28)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v118 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 | v118<<(uint(v109)%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v122 != 0 {
		v108 = v108 + v118
		v109 = v122
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v125 == int32(0) {
		v150 = v73
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v158 = v150 - v73
	goto L24
L38:
	;
	v129 = v73
	v130 = v125
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(base.Ui32(v130)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v138)>>(uint(v130)%32))&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v150 = v146
	goto L37
L41:
	;
	v150 = v129
	goto L37
L42:
	;
	goto L43
L43:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v146 = v129 + int32(1)
	if v144 != 0 {
		v129 = v146
		v130 = v144
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v166 = F_strtod(m, v52, v12+int32(12))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v171 != 0 {
		v381 = v32
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v173 != 0 {
		v381 = v32
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v174 = v166
	goto L23
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v183
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui64(v185-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v181 = base.I32_trunc_f64_s(v177)
	v183 = v181
	goto L50
L52:
	;
	goto L53
L53:
	;
	v183 = int32(-2147483648)
	goto L50
L54:
	;
	v381 = int32(-2)
	goto L2
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = int64(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v194 = base.I32_wrap_i64(v185)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v193
	v360 = v194
	goto L11
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v205
	v208 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v208 == int32(68) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v381 = int32(-2)
	goto L2
L59:
	;
	goto L60
L60:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v213 == int32(0) {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	if v213 != int32(46) {
		v381 = v32
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v220 = m.G0
	v222 = v220 - int32(16)
	m.G0 = v222
	*(*int32)(unsafe.Add(mBase, uint32(v222)+12)) = v212
	v226 = v212 + int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v227 == int32(0) {
		v331 = v8
		goto L64
	} else {
		goto L65
	}
L63:
	;
	m.G0 = v222 + int32(16)
	if v344 != 0 {
		v381 = v32
		goto L2
	} else {
		goto L96
	}
L64:
	;
	v334 = base.F64_nearest(base.F64_mul(v331, float64(1e+06)))
	if base.F64_lt(base.F64_abs(v334), float64(2.147483648e+09)) != 0 {
		goto L93
	} else {
		goto L94
	}
L65:
	;
	v230 = int32(552279)
	v234 = m.G0
	v236 = v234 - int32(32)
	v237 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = v237
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _consts[956])))
	if v245 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v344 = int32(-1)
	goto L63
L67:
	;
	v314 = F_strlen(m, v226)
	mBase = m.M
	if v313 != v314 {
		goto L66
	} else {
		goto L88
	}
L68:
	;
	v313 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, _consts[957])))
	if v249 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v253 = v226
	goto L74
L72:
	;
	goto L73
L73:
	;
	v263 = v230
	v264 = v245
	goto L77
L74:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v259 == v245 {
		v253 = v253 + int32(1)
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v313 = v253 - v226
	goto L67
L76:
	;
	goto L75
L77:
	;
	v271 = v236 + int32(base.Ui32(v264)>>(uint(int32(3))%32))&int32(28)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v273 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = v272 | v273<<(uint(v264)%32)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	if v277 != 0 {
		v263 = v263 + v273
		v264 = v277
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v280 == int32(0) {
		v305 = v226
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	v313 = v305 - v226
	goto L67
L81:
	;
	v284 = v226
	v285 = v280
	goto L82
L82:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v236+int32(base.Ui32(v285)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v293)>>(uint(v285)%32))&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v305 = v301
	goto L80
L84:
	;
	v305 = v284
	goto L80
L85:
	;
	goto L86
L86:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v301 = v284 + int32(1)
	if v299 != 0 {
		v284 = v301
		v285 = v299
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v321 = F_strtod(m, v212, v222+int32(12))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L46
	} else {
		goto L89
	}
L89:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v324 != 0 {
		goto L66
	} else {
		goto L90
	}
L90:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v326 == int32(0) {
		v331 = v321
		goto L64
	} else {
		goto L91
	}
L91:
	;
	goto L66
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)))) = v340
	v344 = int32(0)
	goto L63
L93:
	;
	v338 = base.I32_trunc_f64_s(v334)
	v340 = v338
	goto L92
L94:
	;
	goto L95
L95:
	;
	v340 = int32(-2147483648)
	goto L92
L96:
	;
	goto L12
L97:
	;
	v381 = int32(-2)
	goto L2
L98:
	;
	goto L99
L99:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v360 = v357
	goto L11
L100:
	;
	v381 = int32(-2)
	goto L2
L101:
	;
	goto L102
L102:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v368 < int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v381 = int32(-2)
	goto L2
L104:
	;
	goto L105
L105:
	;
	if base.Ui32(int32(60)) < base.Ui32(v368) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v381 = int32(-2)
	goto L2
L107:
	;
	goto L108
L108:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(int32(1000000)) < base.Ui32(v376) {
		v381 = int32(-2)
		goto L2
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v376
	v381 = int32(0)
	goto L2
}
func F_time_hash_extended(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hashint8extended(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_time_mi_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v7 - v5
		return v9
	}
}
func F_time_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pq_getmsgint64(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if base.Ui64(v6) < base.Ui64(int64(86400000001)) {
			if base.Ui32(v4) <= base.Ui32(int32(6)) {
				v15 = v4 << (uint(int32(3)) % 32)
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[1258])))
				v19 = v18 + v6
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[1259])))
				v23 = base.I64_rem_s(v19, v22)
				v27 = v19 - v23
			} else {
				v27 = v6
			}
			v28 = F_Int64GetDatum(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(402775), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498948), int32(1601), int32(36625))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
	}
}
