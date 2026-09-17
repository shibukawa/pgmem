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
	var v9 float64
	_ = v9
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
	var v72 int32
	_ = v72
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
	var v148 int32
	_ = v148
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
	var v175 float64
	_ = v175
	var v181 int64
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 float64
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v328 float64
	_ = v328
	var v337 int32
	_ = v337
	var v346 int64
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	v5 = int32(0)
	v9 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a_F_DecodeTimeCommon_0)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0])) = v5
	v25 = F_strtox_2(m, l0, v12+int32(8), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0]))
	if v28 == int32(68) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v374
L3:
	;
	v374 = int32(-2)
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
		v374 = v32
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0])) = int32(0)
	v45 = F_strtol(m, v33+int32(1), v12+int32(8), int32(10))
	mBase = m.M
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0]))
	if v48 == int32(68) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v374 = int32(-2)
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
		v374 = v32
		goto L2
	case 12:
		goto L13
	default:
		goto L15
	}
L11:
	;
	if base.Ui32(int32(59)) < base.Ui32(v353) {
		goto L88
	} else {
		goto L89
	}
L12:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if v346 < int64(0) {
		goto L85
	} else {
		goto L86
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0])) = int32(0)
	v201 = F_strtol(m, v52+int32(1), v12+int32(8), int32(10))
	mBase = m.M
	goto L51
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v52
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v53 != 0 {
		v374 = v32
		goto L2
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(0)
	if l1 != int32(_a_F_DecodeTimeCommon_1) {
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
	v374 = int32(-2)
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
	v353 = v69
	goto L11
L21:
	;
	v74 = v52 + int32(1)
	v75 = int32(_a_F_DecodeTimeCommon_2)
	v79 = m.G0
	v81 = v79 - int32(32)
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v82
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[1])))
	if v90 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v175 = v9
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v175, float64(1e+06))))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui64(v181-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L48
	} else {
		goto L49
	}
L24:
	;
	v159 = F_strlen(m, v74)
	mBase = m.M
	if v158 != v159 {
		v374 = v32
		goto L2
	} else {
		goto L43
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
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[2])))
	if v94 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v74
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
	v158 = v98 - v74
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
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v125 == int32(0) {
		v148 = v74
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v158 = v148 - v74
	goto L24
L38:
	;
	v129 = v74
	v130 = v125
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(base.Ui32(v130)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v138)>>(uint(v130)%32))&int32(1) == int32(0) {
		v148 = v129
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v148 = v146
	goto L37
L41:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v146 = v129 + int32(1)
	if v144 != 0 {
		v129 = v146
		v130 = v144
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0])) = int32(0)
	v166 = F_strtod(m, v52, v12+int32(12))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v171 != 0 {
		v374 = v32
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0]))
	if v173 != 0 {
		v374 = v32
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v175 = v166
	goto L23
L48:
	;
	v374 = int32(-2)
	goto L2
L49:
	;
	goto L50
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = int64(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v190 = base.I32_wrap_i64(v181)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v189
	v353 = v190
	goto L11
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v201
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0]))
	if v204 == int32(68) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v374 = int32(-2)
	goto L2
L53:
	;
	goto L54
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v209 == int32(0) {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	if v209 != int32(46) {
		v374 = v32
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v216 = m.G0
	v218 = v216 - int32(16)
	m.G0 = v218
	*(*int32)(unsafe.Add(mBase, uint32(v218)+12)) = v208
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v221 == int32(0) {
		v328 = v9
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v218 + int32(16)
	if v337 != 0 {
		v374 = v32
		goto L2
	} else {
		goto L84
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v328, float64(1e+06))))
	v337 = int32(0)
	goto L57
L59:
	;
	v225 = v208 + int32(1)
	v226 = int32(_a_F_DecodeTimeCommon_2)
	v230 = m.G0
	v232 = v230 - int32(32)
	v233 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v232)+24)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v232))) = v233
	v241 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[1])))
	if v241 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v337 = int32(-1)
	goto L57
L61:
	;
	v310 = F_strlen(m, v225)
	mBase = m.M
	if v309 != v310 {
		goto L60
	} else {
		goto L80
	}
L62:
	;
	v309 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[2])))
	if v245 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v249 = v225
	goto L68
L66:
	;
	goto L67
L67:
	;
	v259 = v226
	v260 = v241
	goto L71
L68:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v255 == v241 {
		v249 = v249 + int32(1)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v309 = v249 - v225
	goto L61
L70:
	;
	goto L69
L71:
	;
	v267 = v232 + int32(base.Ui32(v260)>>(uint(int32(3))%32))&int32(28)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v268 | v269<<(uint(v260)%32)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	if v273 != 0 {
		v259 = v259 + v269
		v260 = v273
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v276 == int32(0) {
		v299 = v225
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v309 = v299 - v225
	goto L61
L75:
	;
	v280 = v225
	v281 = v276
	goto L76
L76:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v232+int32(base.Ui32(v281)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v289)>>(uint(v281)%32))&int32(1) == int32(0) {
		v299 = v280
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v299 = v297
	goto L74
L78:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v297 = v280 + int32(1)
	if v295 != 0 {
		v280 = v297
		v281 = v295
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0])) = int32(0)
	v317 = F_strtod(m, v208, v218+int32(12))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L44
	} else {
		goto L81
	}
L81:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v320 != 0 {
		goto L60
	} else {
		goto L82
	}
L82:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeCommon[0]))
	if v322 == int32(0) {
		v328 = v317
		goto L58
	} else {
		goto L83
	}
L83:
	;
	goto L60
L84:
	;
	goto L12
L85:
	;
	v374 = int32(-2)
	goto L2
L86:
	;
	goto L87
L87:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v353 = v350
	goto L11
L88:
	;
	v374 = int32(-2)
	goto L2
L89:
	;
	goto L90
L90:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v361 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v374 = int32(-2)
	goto L2
L92:
	;
	goto L93
L93:
	;
	if base.Ui32(int32(60)) < base.Ui32(v361) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v374 = int32(-2)
	goto L2
L95:
	;
	goto L96
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(int32(_a_F_DecodeTimeCommon_3)) < base.Ui32(v369) {
		v374 = int32(-2)
		goto L2
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v369
	v374 = int32(0)
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
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
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
				v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_time_recv[0])))
				v17 = v16 + v6
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_time_recv[1])))
				v19 = base.I64_rem_s(v17, v18)
				v23 = v17 - v19
			} else {
				v23 = v6
			}
			v24 = F_Int64GetDatum(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v24
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_time_recv_0), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_time_recv_1), int32(1601), int32(_a_F_time_recv_2))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
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
