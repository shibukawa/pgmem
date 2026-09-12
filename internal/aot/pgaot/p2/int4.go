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
				F_errmsg(m, int32(401737), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497640), int32(150), int32(558481))
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
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
	return v274
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
	v57 = F_cstring_to_text(m, int32(756936))
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
	v274 = v57
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
	v252 = v63 + int32(4)
	F_NUM_processor(m, v69, v11+int32(-36), v252, v242, int32(0), v243, v245, int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L78
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
	v242 = v74
	v243 = int32(0)
	v245 = v2
	goto L21
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v79
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = base.F64_convert_i32_s(v15)
	v84 = F_psprintf(m, int32(419646), v13)
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
	v242 = v84
	v243 = int32(0)
	v245 = v2
	goto L21
L31:
	;
	goto L32
L32:
	;
	v90 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v90)
	v242 = v84
	v243 = int32(0)
	v245 = v2
	goto L21
L33:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v122 = base.B2i32(v120 == int32(45))
	v123 = v117 + v122
	v124 = F_strlen(m, v123)
	mBase = m.M
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v125 != 0 {
		goto L43
	} else {
		goto L44
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
	v115 = F_DirectFunctionCall1Coll(m, int32(1314), int32(0), v15)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L42
	}
L37:
	;
	v108 = F_DirectFunctionCall1Coll(m, int32(1314), int32(0), v106*v15)
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
	v129 = F_palloc(m, v124+v125+int32(2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v216 = v123
	goto L45
L45:
	;
	if v120 == int32(45) {
		goto L69
	} else {
		goto L70
	}
L46:
	;
	if (v123^v129)&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v205 = v129 + v124
	v206 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v206)
	v212 = F__emscripten_memset_bulkmem(m, v205+int32(1), base.I32_extend8_s(int32(48)), v125)
	mBase = m.M
	goto L68
L48:
	;
	goto L47
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v184)
	if v184&int32(255) == int32(0) {
		goto L48
	} else {
		goto L64
	}
L50:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v183 = v123
	v184 = v136
	v185 = v129
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v123&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v140 = v123
	v142 = v129
	goto L56
L54:
	;
	v154 = v123
	v156 = v129
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 != v161 {
		v183 = v154
		v184 = v158
		v185 = v156
		goto L49
	} else {
		goto L60
	}
L56:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v143)
	if v143 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L57:
	;
	v154 = v150
	v156 = v148
	goto L55
L58:
	;
	v147 = int32(1)
	v148 = v142 + v147
	v150 = v140 + v147
	if v150&int32(3) != 0 {
		v140 = v150
		v142 = v148
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v166 = v154
	v167 = v158
	v168 = v156
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v167
	v170 = int32(4)
	v171 = v168 + v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v174 = v166 + v170
	v178 = int32(-2139062144)
	if (v172|(int32(16843008)-v172))&v178 == v178 {
		v166 = v174
		v167 = v172
		v168 = v171
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v183 = v174
	v184 = v172
	v185 = v171
	goto L49
L63:
	;
	goto L62
L64:
	;
	v192 = v183
	v194 = v185
	goto L65
L65:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)) = uint8(v195)
	v197 = int32(1)
	if v195 != 0 {
		v192 = v192 + v197
		v194 = v194 + v197
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L48
L67:
	;
	goto L66
L68:
	;
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v205+v125)+1)) = uint8(v214)
	v216 = v129
	goto L45
L69:
	;
	v220 = int32(45)
	goto L71
L70:
	;
	v220 = int32(43)
	goto L71
L71:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v124 < v221 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v242 = v216
	v243 = v221 - v124
	v245 = v220
	goto L21
L73:
	;
	goto L74
L74:
	;
	v224 = int32(0)
	if v124 <= v221 {
		v242 = v216
		v243 = v224
		v245 = v220
		goto L21
	} else {
		goto L75
	}
L75:
	;
	v226 = v125 + v221
	v229 = F_palloc(m, v226+int32(2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	v233 = v226 + int32(1)
	v235 = F__emscripten_memset_bulkmem(m, v229, base.I32_extend8_s(int32(35)), v233)
	mBase = m.M
	goto L77
L77:
	;
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v233))) = uint8(v237)
	v240 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v221))) = uint8(v240)
	v242 = v229
	v243 = v224
	v245 = v220
	goto L21
L78:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	if v257 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_pfree(m, v69)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v262 = F_strlen(m, v252)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v262<<(uint(int32(2))%32) + int32(16)
	v274 = v63
	goto L1
L82:
	;
	goto L81
}
