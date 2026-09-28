package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_mul_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v39 int64
	_ = v39
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = int64(63)
	v11 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	v18 = int64(32)
	v19 = int64(base.Ui64(v11) >> (uint(v18) % 64))
	v21 = int64(base.Ui64(v8) >> (uint(v18) % 64))
	v24 = int64(4294967295)
	v25 = v11 & v24
	v27 = v8 & v24
	v28 = v25 * v27
	v32 = int64(base.Ui64(v28)>>(uint(v18)%64)) + v25*v21
	v39 = v27*v19 + v32&v24
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8*(v11>>(uint(v9)%64)) + v8>>(uint(v9)%64)*v11 + v19*v21 + int64(base.Ui64(v32)>>(uint(v18)%64)) + int64(base.Ui64(v39)>>(uint(v18)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v28&v24 | v39<<(uint(v18)%64)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v50 != v51>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int4_mul_cash_0), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4_mul_cash_1), int32(151), int32(_a_F_int4_mul_cash_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v6 + int32(16)
		return v51
	}
}
func F_int4_to_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 float64
	_ = v102
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return base.I64_extend_i32_u(v270)
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v49-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int64(0)
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v20 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v26 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v37 = int32(1)
	if v20&v37 != 0 {
		v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v29 = int32(16)
	goto L10
L9:
	;
	v29 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = int32(4)
	goto L13
L12:
	;
	v36 = v29
	goto L13
L13:
	;
	v49 = v36
	goto L2
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v55 = F_cstring_to_text(m, int32(_a_F_int4_to_char_0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v57 = base.I32_wrap_i64(v14)
	v62 = F_palloc0(m, v49<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v270 = v55
	goto L1
L19:
	;
	v68 = F_NUM_cache(m, v49, v10+int32(-36), v16, v10+int32(-37))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v70&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v250 = v62 + int32(4)
	F_NUM_processor(m, v68, v10+int32(-36), v250, v240, int32(0), v246, v244, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L76
	}
L22:
	;
	v73 = F_int_to_roman(m, v57)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v70&int32(_a_F_int4_to_char_1) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v240 = v73
	v244 = v2
	v246 = int32(0)
	goto L21
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v78
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = base.F64_convert_i32_s(v57)
	v84 = F_psprintf(m, int32(_a_F_int4_to_char_2), v12)
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
	if v70&int32(2048) != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v86 != int32(43) {
		v240 = v84
		v244 = v2
		v246 = int32(0)
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v89 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v89)
	v240 = v84
	v244 = v2
	v246 = int32(0)
	goto L21
L31:
	;
	v95 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v102 = F_pow(m, float64(10), base.F64_convert_i32_s(v100))
	mBase = m.M
	v104 = F_DirectFunctionCall1Coll(m, int32(1437), v95, base.I64_reinterpret_f64(v102))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	v112 = v14
	goto L33
L33:
	;
	v116 = F_DirectFunctionCall1Coll(m, int32(1438), int32(0), base.I64_extend32_s(v112))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	v106 = F_DirectFunctionCall2Coll(m, int32(1436), v95, base.I64_extend32_s(v14), v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v100 + v108
	v112 = v106
	goto L33
L36:
	;
	v118 = base.I32_wrap_i64(v116)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v121 = base.B2i32(v119 == int32(45))
	v122 = v118 + v121
	v123 = F_strlen(m, v122)
	mBase = m.M
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v124 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v128 = F_palloc(m, v123+v124+int32(2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	v214 = v122
	goto L39
L39:
	;
	if v119 == int32(45) {
		goto L65
	} else {
		goto L66
	}
L40:
	;
	if (v122^v128)&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v204 = v128 + v123
	v205 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v205)
	if v124 != 0 {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	goto L41
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v183)
	if v183&int32(255) == int32(0) {
		goto L42
	} else {
		goto L58
	}
L44:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v182 = v122
	v183 = v135
	v184 = v128
	goto L43
L45:
	;
	goto L46
L46:
	;
	if v122&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v139 = v122
	v141 = v128
	goto L50
L48:
	;
	v153 = v122
	v155 = v128
	goto L49
L49:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v160 = int32(-2139062144)
	if (int32(16843008)-v157|v157)&v160 != v160 {
		v182 = v153
		v183 = v157
		v184 = v155
		goto L43
	} else {
		goto L54
	}
L50:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v142)
	if v142 == int32(0) {
		goto L42
	} else {
		goto L52
	}
L51:
	;
	v153 = v149
	v155 = v147
	goto L49
L52:
	;
	v146 = int32(1)
	v147 = v141 + v146
	v149 = v139 + v146
	if v149&int32(3) != 0 {
		v139 = v149
		v141 = v147
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v165 = v153
	v166 = v157
	v167 = v155
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v166
	v169 = int32(4)
	v170 = v167 + v169
	v172 = v165 + v169
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v177 = int32(-2139062144)
	if (int32(16843008)-v174|v174)&v177 == v177 {
		v165 = v172
		v166 = v174
		v167 = v170
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v182 = v172
	v183 = v174
	v184 = v170
	goto L43
L57:
	;
	goto L56
L58:
	;
	v191 = v182
	v193 = v184
	goto L59
L59:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)) = uint8(v194)
	v196 = int32(1)
	if v194 != 0 {
		v191 = v191 + v196
		v193 = v193 + v196
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L42
L61:
	;
	goto L60
L62:
	;
	base.MemoryFill(m, v204+int32(1), int32(48), v124)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124+v204)+1)) = uint8(v212)
	v214 = v128
	goto L39
L65:
	;
	v218 = int32(45)
	goto L67
L66:
	;
	v218 = int32(43)
	goto L67
L67:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if base.Ui32(v123) < base.Ui32(v219) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v240 = v214
	v244 = v218
	v246 = v219 - v123
	goto L21
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v123) <= base.Ui32(v219) {
		v240 = v214
		v244 = v218
		v246 = int32(0)
		goto L21
	} else {
		goto L71
	}
L71:
	;
	v224 = v124 + v219
	v227 = F_palloc(m, v224+int32(2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v230 = v224 + int32(1)
	if v230 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryFill(m, v227, int32(35), v230)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v227+v230))) = uint8(v234)
	v237 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v227+v219))) = uint8(v237)
	v240 = v227
	v244 = v218
	v246 = v234
	goto L21
L76:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+27)))
	if v255 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_pfree(m, v68)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v260 = F_strlen(m, v250)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v260<<(uint(int32(2))%32) + int32(16)
	v270 = v62
	goto L1
L80:
	;
	goto L79
}
