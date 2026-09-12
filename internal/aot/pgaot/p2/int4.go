package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v9) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v9 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v12>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v29&v25 | v40<<(uint(v19)%64)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v51 != v52>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(387858), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(481059), int32(150), int32(533617))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
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
		v74 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v74
		}
	}
}
func F_int4_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 float64
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v386 int32
	_ = v386
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 - int32(-64)
	return v386
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v51-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int32(0)
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = int32(4)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v26&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v39 = int32(1)
	if v21&v39 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v39)%32)) - v39
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v35 = v24
	goto L10
L9:
	;
	v35 = base.B2i32(v26 == int32(18)) << (uint(v24) % 32)
	goto L10
L10:
	;
	if v26 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = v24
	goto L13
L12:
	;
	v38 = v35
	goto L13
L13:
	;
	v51 = v38
	goto L2
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v57 = F_cstring_to_text(m, int32(722455))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v63 = F_palloc0(m, v51<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v386 = v57
	goto L1
L19:
	;
	v69 = F_NUM_cache(m, v51, v11+int32(-36), v17, v11+int32(-37))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v71&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v308 = v63 + int32(4)
	F_NUM_processor(m, v69, v11+int32(-36), v308, v298, int32(0), v299, v301, int32(1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L3
	} else {
		goto L95
	}
L22:
	;
	v74 = F_int_to_roman(m, v15)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v71&int32(16384) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v298 = v74
	v299 = int32(0)
	v301 = v2
	goto L21
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v79
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = base.F64_convert_i32_s(v15)
	v84 = F_psprintf(m, int32(405038), v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v71&int32(2048) != 0 {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v86 != int32(43) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v298 = v84
	v299 = int32(0)
	v301 = v2
	goto L21
L31:
	;
	goto L32
L32:
	;
	v90 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v90)
	v298 = v84
	v299 = int32(0)
	v301 = v2
	goto L21
L33:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v122 = base.B2i32(v120 == int32(45))
	v123 = v117 + v122
	if v123&int32(3) == int32(0) {
		v147 = v123
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v100 = F_pow(m, float64(10), base.F64_convert_i32_s(v98))
	mBase = m.M
	if base.F64_lt(base.F64_abs(v100), float64(2.147483648e+09)) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	v115 = F_DirectFunctionCall1Coll(m, int32(1313), int32(0), v15)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L42
	}
L37:
	;
	v108 = F_DirectFunctionCall1Coll(m, int32(1313), int32(0), v106*v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L41
	}
L38:
	;
	v104 = base.I32_trunc_f64_s(v100)
	v106 = v104
	goto L37
L39:
	;
	goto L40
L40:
	;
	v106 = int32(-2147483648)
	goto L37
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v110 + v98
	v117 = v108
	goto L33
L42:
	;
	v117 = v115
	goto L33
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v181 != 0 {
		goto L60
	} else {
		goto L61
	}
L44:
	;
	v180 = v172 - v123
	goto L43
L45:
	;
	v151 = v147
	goto L54
L46:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v131 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v180 = int32(0)
	goto L43
L48:
	;
	goto L49
L49:
	;
	v136 = v123
	goto L50
L50:
	;
	v140 = v136 + int32(1)
	if v140&int32(3) == int32(0) {
		v147 = v140
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v172 = v140
	goto L44
L52:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v145 != 0 {
		v136 = v140
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v160 = int32(-2139062144)
	if (int32(16843008)-v157|v157)&v160 == v160 {
		v151 = v151 + int32(4)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v166 = v151
	goto L57
L56:
	;
	goto L55
L57:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v170 != 0 {
		v166 = v166 + int32(1)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v172 = v166
	goto L44
L59:
	;
	goto L58
L60:
	;
	v185 = F_palloc(m, v180+v181+int32(2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	v272 = v123
	goto L62
L62:
	;
	if v120 == int32(45) {
		goto L86
	} else {
		goto L87
	}
L63:
	;
	if (v123^v185)&int32(3) != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v261 = v185 + v180
	v262 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v262)
	v268 = F__emscripten_memset_bulkmem(m, v261+int32(1), base.I32_extend8_s(int32(48)), v181)
	mBase = m.M
	goto L85
L65:
	;
	goto L64
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v240)
	if v240&int32(255) == int32(0) {
		goto L65
	} else {
		goto L81
	}
L67:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v239 = v123
	v240 = v192
	v241 = v185
	goto L66
L68:
	;
	goto L69
L69:
	;
	if v123&int32(3) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v196 = v123
	v198 = v185
	goto L73
L71:
	;
	v210 = v123
	v212 = v185
	goto L72
L72:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v217 = int32(-2139062144)
	if (int32(16843008)-v214|v214)&v217 != v217 {
		v239 = v210
		v240 = v214
		v241 = v212
		goto L66
	} else {
		goto L77
	}
L73:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v199)
	if v199 == int32(0) {
		goto L65
	} else {
		goto L75
	}
L74:
	;
	v210 = v206
	v212 = v204
	goto L72
L75:
	;
	v203 = int32(1)
	v204 = v198 + v203
	v206 = v196 + v203
	if v206&int32(3) != 0 {
		v196 = v206
		v198 = v204
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v222 = v210
	v223 = v214
	v224 = v212
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v223
	v226 = int32(4)
	v227 = v224 + v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v230 = v222 + v226
	v234 = int32(-2139062144)
	if (v228|(int32(16843008)-v228))&v234 == v234 {
		v222 = v230
		v223 = v228
		v224 = v227
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v239 = v230
	v240 = v228
	v241 = v227
	goto L66
L80:
	;
	goto L79
L81:
	;
	v248 = v239
	v250 = v241
	goto L82
L82:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)) = uint8(v251)
	v253 = int32(1)
	if v251 != 0 {
		v248 = v248 + v253
		v250 = v250 + v253
		goto L82
	} else {
		goto L84
	}
L83:
	;
	goto L65
L84:
	;
	goto L83
L85:
	;
	v270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v261+v181)+1)) = uint8(v270)
	v272 = v185
	goto L62
L86:
	;
	v276 = int32(45)
	goto L88
L87:
	;
	v276 = int32(43)
	goto L88
L88:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v180 < v277 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v298 = v272
	v299 = v277 - v180
	v301 = v276
	goto L21
L90:
	;
	goto L91
L91:
	;
	v280 = int32(0)
	if v180 <= v277 {
		v298 = v272
		v299 = v280
		v301 = v276
		goto L21
	} else {
		goto L92
	}
L92:
	;
	v282 = v181 + v277
	v285 = F_palloc(m, v282+int32(2))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L3
	} else {
		goto L93
	}
L93:
	;
	v289 = v282 + int32(1)
	v291 = F__emscripten_memset_bulkmem(m, v285, base.I32_extend8_s(int32(35)), v289)
	mBase = m.M
	goto L94
L94:
	;
	v293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v291+v289))) = uint8(v293)
	v296 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v291+v277))) = uint8(v296)
	v298 = v285
	v299 = v280
	v301 = v276
	goto L21
L95:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	if v313 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_pfree(m, v69)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L3
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v308&int32(3) == int32(0) {
		v341 = v308
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L98
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v374<<(uint(int32(2))%32) + int32(16)
	v386 = v63
	goto L1
L101:
	;
	v374 = v366 - v308
	goto L100
L102:
	;
	v345 = v341
	goto L111
L103:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	if v325 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v374 = int32(0)
	goto L100
L105:
	;
	goto L106
L106:
	;
	v330 = v308
	goto L107
L107:
	;
	v334 = v330 + int32(1)
	if v334&int32(3) == int32(0) {
		v341 = v334
		goto L102
	} else {
		goto L109
	}
L108:
	;
	v366 = v334
	goto L101
L109:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v339 != 0 {
		v330 = v334
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v354 = int32(-2139062144)
	if (int32(16843008)-v351|v351)&v354 == v354 {
		v345 = v345 + int32(4)
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v360 = v345
	goto L114
L113:
	;
	goto L112
L114:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if v364 != 0 {
		v360 = v360 + int32(1)
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v366 = v360
	goto L101
L116:
	;
	goto L115
}
