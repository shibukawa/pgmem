package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShowGUCOption(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 float64
	_ = v70
	var v76 int64
	_ = v76
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 float64
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 float64
	_ = v173
	var v177 float64
	_ = v177
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v207 float64
	_ = v207
	var v210 int32
	_ = v210
	var v215 float64
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	v11 = m.G0
	v13 = v11 - int32(320)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v16 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	case 3:
		goto L4
	case 4:
		goto L3
	default:
		v299 = int32(_a_F_ShowGUCOption_0)
		goto L1
	}
L1:
	;
	v307 = F_pstrdup(m, v299)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L9
	} else {
		goto L81
	}
L2:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v295 != 0 {
		goto L78
	} else {
		goto L79
	}
L3:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v235 != 0 {
		goto L63
	} else {
		goto L64
	}
L4:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v228 != 0 {
		goto L55
	} else {
		goto L56
	}
L5:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v124 != 0 {
		goto L35
	} else {
		goto L36
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v24 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v20 = m.T0[v17].(func(*base.Module) int32)(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v299 = v20
	goto L1
L11:
	;
	v25 = m.T0[v24].(func(*base.Module) int32)(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = base.I64_extend_i32_s(v28)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v29
	v31 = int32(_a_F_ShowGUCOption_1)
	v32 = int32(0)
	if base.B2i32(l1 == v32)|base.B2i32(v28 <= v32) != 0 {
		v108 = v31
		v115 = v29
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v299 = v25
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v115
	v119 = v13 - int32(-64)
	v122 = F_pg_snprintf(m, v119, int32(256), int32(_a_F_ShowGUCOption_2), v13)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L34
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = v37 & int32(2130706432)
	if v39 == int32(0) {
		v108 = v31
		v115 = v29
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v42 = int32(0)
	v44 = v13 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
	if v39&int32(251658240) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = int32(_a_F_ShowGUCOption_3)
	goto L20
L19:
	;
	v51 = int32(_a_F_ShowGUCOption_4)
	goto L20
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v55 = v42
	goto L24
L22:
	;
	goto L23
L23:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
	v108 = v104
	v115 = v105
	goto L15
L24:
	;
	v67 = v51 + v55<<(uint(int32(4))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v39 != v68 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L23
L26:
	;
	goto L25
L27:
	;
	v87 = v55 + int32(1)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v87<<(uint(int32(4))%32)))))
	if v91 != 0 {
		v55 = v87
		goto L24
	} else {
		goto L33
	}
L28:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	if base.F64_le(v70, float64(1)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v76 = base.I64_rem_s(v29, base.I64_trunc_sat_f64_s(v70))
	if v76 != int64(0) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(56)))) = base.I64_trunc_sat_f64_s(base.F64_nearest(base.F64_div(base.F64_convert_i64_u(v29), v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v67
	goto L26
L32:
	;
	goto L31
L33:
	;
	goto L26
L34:
	;
	v299 = v119
	goto L1
L35:
	;
	v125 = m.T0[v124].(func(*base.Module) int32)(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v127)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+56)) = v128
	v130 = int32(_a_F_ShowGUCOption_1)
	v131 = int32(0)
	if base.B2i32(l1 == v131)|base.B2i32(base.F64_gt(v128, float64(0)) == v131) != 0 {
		v210 = v130
		v215 = v128
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v299 = v125
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v210
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v215
	v221 = v13 - int32(-64)
	v226 = F_pg_snprintf(m, v221, int32(256), int32(_a_F_ShowGUCOption_5), v13+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L9
	} else {
		goto L54
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v140 = v138 & int32(2130706432)
	if v140 == int32(0) {
		v210 = v130
		v215 = v128
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v144 = v13 + int32(56)
	v145 = int32(0)
	v147 = v13 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v145
	if v140&int32(251658240) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v207 = *(*float64)(unsafe.Add(mBase, uint32(v13)+56))
	v210 = v206
	v215 = v207
	goto L39
L43:
	;
	v154 = int32(_a_F_ShowGUCOption_3)
	goto L45
L44:
	;
	v154 = int32(_a_F_ShowGUCOption_4)
	goto L45
L45:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v155 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v158 = v145
	goto L47
L47:
	;
	v170 = v154 + v158<<(uint(int32(4))%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v140 != v171 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L42
L49:
	;
	v191 = v158 + int32(1)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v191<<(uint(int32(4))%32)))))
	if v195 != 0 {
		v158 = v191
		goto L47
	} else {
		goto L53
	}
L50:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v170)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v144))) = base.F64_div(v128, v173)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v170
	v177 = *(*float64)(unsafe.Add(mBase, uint32(v144)))
	if base.F64_gt(v177, float64(0)) == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	if base.F64_le(base.F64_abs(base.F64_add(base.F64_div(base.F64_nearest(v177), v177), float64(-1))), float64(1e-08)) != 0 {
		goto L42
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	goto L48
L54:
	;
	v299 = v221
	goto L1
L55:
	;
	v229 = m.T0[v228].(func(*base.Module) int32)(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L9
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v232 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v299 = v229
	goto L1
L59:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v233 != 0 {
		v299 = v232
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v299 = int32(_a_F_ShowGUCOption_1)
	goto L1
L62:
	;
	goto L61
L63:
	;
	v236 = m.T0[v235].(func(*base.Module) int32)(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v240 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v299 = v236
	goto L1
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L9
	} else {
		goto L75
	}
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	if v243 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v246 == v239 {
		v299 = v243
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v249 = v240
	goto L71
L71:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	if v258 == int32(0) {
		goto L67
	} else {
		goto L73
	}
L72:
	;
	v299 = v258
	goto L1
L73:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v239 != v261 {
		v249 = v249 + int32(12)
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v239
	F_errmsg_internal(m, int32(_a_F_ShowGUCOption_6), v13+int32(32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ShowGUCOption_7), int32(3036), int32(_a_F_ShowGUCOption_8))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v296 = int32(_a_F_ShowGUCOption_9)
	goto L80
L79:
	;
	v296 = int32(_a_F_ShowGUCOption_10)
	goto L80
L80:
	;
	v299 = v296
	goto L1
L81:
	;
	m.G0 = v13 + int32(320)
	return v307
}
