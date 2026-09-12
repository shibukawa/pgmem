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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 float64
	_ = v69
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v85 float64
	_ = v85
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v124 int64
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 float64
	_ = v183
	var v187 float64
	_ = v187
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v217 float64
	_ = v217
	var v220 int32
	_ = v220
	var v225 float64
	_ = v225
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
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
		v306 = int32(545093)
		goto L1
	}
L1:
	;
	v314 = F_pstrdup(m, v306)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L9
	} else {
		goto L90
	}
L2:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v302 != 0 {
		goto L87
	} else {
		goto L88
	}
L3:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v247 != 0 {
		goto L73
	} else {
		goto L74
	}
L4:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v240 != 0 {
		goto L65
	} else {
		goto L66
	}
L5:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v135 != 0 {
		goto L44
	} else {
		goto L45
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
	v306 = v20
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
	v31 = int32(741336)
	if l1 == int32(0) {
		v117 = v31
		v124 = v29
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v306 = v25
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v124
	v131 = F_pg_snprintf(m, v13-int32(-64), int32(256), int32(175606), v13)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L9
	} else {
		goto L43
	}
L16:
	;
	if v28 <= int32(0) {
		v117 = v31
		v124 = v29
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = v36 & int32(2130706432)
	if v38 == int32(0) {
		v117 = v31
		v124 = v29
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v43 = int32(0)
	v45 = v13 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v43
	if v38&int32(251658240) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v52 = int32(1747136)
	goto L21
L20:
	;
	v52 = int32(1747552)
	goto L21
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = v43
	goto L25
L23:
	;
	goto L24
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
	v117 = v113
	v124 = v114
	goto L15
L25:
	;
	v66 = v52 + v54<<(uint(int32(4))%32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v38 != v67 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L24
L27:
	;
	goto L26
L28:
	;
	v96 = v54 + int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v96<<(uint(int32(4))%32)))))
	if v100 != 0 {
		v54 = v96
		goto L25
	} else {
		goto L42
	}
L29:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v66)+8))
	if base.F64_le(v69, float64(1)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.F64_lt(base.F64_abs(v69), float64(9.223372036854776e+18)) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v85 = base.F64_nearest(base.F64_div(base.F64_convert_i64_u(v29), v69))
	if base.F64_lt(base.F64_abs(v85), float64(9.223372036854776e+18)) != 0 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v80 = base.I64_rem_s(v29, v79)
	if v80 != int64(0) {
		goto L28
	} else {
		goto L37
	}
L34:
	;
	v77 = base.I64_trunc_f64_s(v69)
	v79 = v77
	goto L33
L35:
	;
	goto L36
L36:
	;
	v79 = int64(-9223372036854775807 - 1)
	goto L33
L37:
	;
	goto L32
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(56)))) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v66
	goto L27
L39:
	;
	v89 = base.I64_trunc_f64_s(v85)
	v91 = v89
	goto L38
L40:
	;
	goto L41
L41:
	;
	v91 = int64(-9223372036854775807 - 1)
	goto L38
L42:
	;
	goto L27
L43:
	;
	v306 = v13 - int32(-64)
	goto L1
L44:
	;
	v136 = m.T0[v135].(func(*base.Module) int32)(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v138)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+56)) = v139
	v141 = int32(741336)
	if l1 == int32(0) {
		v220 = v141
		v225 = v139
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v306 = v136
	goto L1
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v220
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v225
	v236 = F_pg_snprintf(m, v13-int32(-64), int32(256), int32(175601), v13+int32(16))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L64
	}
L49:
	;
	if base.F64_gt(v139, float64(0)) == int32(0) {
		v220 = v141
		v225 = v139
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v150 = v148 & int32(2130706432)
	if v150 == int32(0) {
		v220 = v141
		v225 = v139
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v154 = v13 + int32(56)
	v155 = int32(0)
	v157 = v13 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v155
	if v150&int32(251658240) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v13)+56))
	v220 = v216
	v225 = v217
	goto L48
L53:
	;
	v164 = int32(1747136)
	goto L55
L54:
	;
	v164 = int32(1747552)
	goto L55
L55:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v165 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v168 = v155
	goto L57
L57:
	;
	v180 = v164 + v168<<(uint(int32(4))%32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v150 != v181 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L52
L59:
	;
	v201 = v168 + int32(1)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v201<<(uint(int32(4))%32)))))
	if v205 != 0 {
		v168 = v201
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v180)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v154))) = base.F64_div(v139, v183)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v180
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v154)))
	if base.F64_gt(v187, float64(0)) == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	if base.F64_le(base.F64_abs(base.F64_add(base.F64_div(base.F64_nearest(v187), v187), float64(-1))), float64(1e-08)) != 0 {
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L58
L64:
	;
	v306 = v13 - int32(-64)
	goto L1
L65:
	;
	v241 = m.T0[v240].(func(*base.Module) int32)(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	if v244 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v306 = v241
	goto L1
L69:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v245 != 0 {
		v306 = v244
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v306 = int32(741336)
	goto L1
L72:
	;
	goto L71
L73:
	;
	v248 = m.T0[v247].(func(*base.Module) int32)(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v252 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v306 = v248
	goto L1
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L84
	}
L78:
	;
	v256 = v252
	goto L79
L79:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v265 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L80:
	;
	goto L77
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v268 == v251 {
		v306 = v265
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v271 = v256 + int32(12)
	if v271 != 0 {
		v256 = v271
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v251
	F_errmsg_internal(m, int32(181449), v13+int32(32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(499158), int32(3036), int32(344377))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v303 = int32(272805)
	goto L89
L88:
	;
	v303 = int32(338517)
	goto L89
L89:
	;
	v306 = v303
	goto L1
L90:
	;
	m.G0 = v13 + int32(320)
	return v314
}
