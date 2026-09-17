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
				F_errmsg(m, int32(_a_F_int4_mul_cash_0), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4_mul_cash_1), int32(150), int32(_a_F_int4_mul_cash_2))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 float64
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v261
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v48-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int32(0)
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v19 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v25 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if v19&v36 != 0 {
		v48 = int32(base.Ui32(v19)>>(uint(v36)%32)) - v36
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v28 = int32(16)
	goto L10
L9:
	;
	v28 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = int32(4)
	goto L13
L12:
	;
	v35 = v28
	goto L13
L13:
	;
	v48 = v35
	goto L2
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v54 = F_cstring_to_text(m, int32(_a_F_int4_to_char_0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v60 = F_palloc0(m, v48<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v261 = v54
	goto L1
L19:
	;
	v66 = F_NUM_cache(m, v48, v9+int32(-36), v15, v9+int32(-37))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v68&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v241 = v60 + int32(4)
	F_NUM_processor(m, v66, v9+int32(-36), v241, v232, int32(0), v237, v236, int32(1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L3
	} else {
		goto L76
	}
L22:
	;
	v71 = F_int_to_roman(m, v13)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v68&int32(_a_F_int4_to_char_1) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v232 = v71
	v236 = v2
	v237 = int32(0)
	goto L21
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = base.F64_convert_i32_s(v13)
	v82 = F_psprintf(m, int32(_a_F_int4_to_char_2), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v68&int32(2048) != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v84 != int32(43) {
		v232 = v82
		v236 = v2
		v237 = int32(0)
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v87 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v87)
	v232 = v82
	v236 = v2
	v237 = int32(0)
	goto L21
L31:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v113 = base.B2i32(v111 == int32(45))
	v114 = v109 + v113
	v115 = F_strlen(m, v114)
	mBase = m.M
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v116 != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v97 = F_pow(m, float64(10), base.F64_convert_i32_s(v95))
	mBase = m.M
	v100 = F_DirectFunctionCall1Coll(m, int32(1298), int32(0), v13*base.I32_trunc_sat_f64_s(v97))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v107 = F_DirectFunctionCall1Coll(m, int32(1298), int32(0), v13)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L36
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v95 + v102
	v109 = v100
	goto L31
L36:
	;
	v109 = v107
	goto L31
L37:
	;
	v120 = F_palloc(m, v115+v116+int32(2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	v206 = v114
	goto L39
L39:
	;
	if v111 == int32(45) {
		goto L65
	} else {
		goto L66
	}
L40:
	;
	if (v114^v120)&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v196 = v120 + v115
	v197 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v197)
	if v116 != 0 {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	goto L41
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v175)
	if v175&int32(255) == int32(0) {
		goto L42
	} else {
		goto L58
	}
L44:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v174 = v114
	v175 = v127
	v176 = v120
	goto L43
L45:
	;
	goto L46
L46:
	;
	if v114&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v131 = v114
	v133 = v120
	goto L50
L48:
	;
	v145 = v114
	v147 = v120
	goto L49
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v152 = int32(-2139062144)
	if (int32(16843008)-v149|v149)&v152 != v152 {
		v174 = v145
		v175 = v149
		v176 = v147
		goto L43
	} else {
		goto L54
	}
L50:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
	if v134 == int32(0) {
		goto L42
	} else {
		goto L52
	}
L51:
	;
	v145 = v141
	v147 = v139
	goto L49
L52:
	;
	v138 = int32(1)
	v139 = v133 + v138
	v141 = v131 + v138
	if v141&int32(3) != 0 {
		v131 = v141
		v133 = v139
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v157 = v145
	v158 = v149
	v159 = v147
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v158
	v161 = int32(4)
	v162 = v159 + v161
	v164 = v157 + v161
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v169 = int32(-2139062144)
	if (int32(16843008)-v166|v166)&v169 == v169 {
		v157 = v164
		v158 = v166
		v159 = v162
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v174 = v164
	v175 = v166
	v176 = v162
	goto L43
L57:
	;
	goto L56
L58:
	;
	v183 = v174
	v185 = v176
	goto L59
L59:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)) = uint8(v186)
	v188 = int32(1)
	if v186 != 0 {
		v183 = v183 + v188
		v185 = v185 + v188
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
	base.MemoryFill(m, v196+int32(1), int32(48), v116)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116+v196)+1)) = uint8(v204)
	v206 = v120
	goto L39
L65:
	;
	v210 = int32(45)
	goto L67
L66:
	;
	v210 = int32(43)
	goto L67
L67:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v115 < v211 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v232 = v206
	v236 = v210
	v237 = v211 - v115
	goto L21
L69:
	;
	goto L70
L70:
	;
	if v115 <= v211 {
		v232 = v206
		v236 = v210
		v237 = int32(0)
		goto L21
	} else {
		goto L71
	}
L71:
	;
	v216 = v116 + v211
	v219 = F_palloc(m, v216+int32(2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v222 = v216 + int32(1)
	if v222 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryFill(m, v219, int32(35), v222)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v219+v222))) = uint8(v226)
	v229 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v219+v211))) = uint8(v229)
	v232 = v219
	v236 = v210
	v237 = v226
	goto L21
L76:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+27)))
	if v246 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_pfree(m, v66)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L3
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v251 = F_strlen(m, v241)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v251<<(uint(int32(2))%32) + int32(16)
	v261 = v60
	goto L1
L80:
	;
	goto L79
}
