package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HalfvecInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1486])) = int32(7530)
	*(*int32)(unsafe.Add(mBase, _consts[1487])) = int32(7531)
	*(*int32)(unsafe.Add(mBase, _consts[1488])) = int32(7532)
	*(*int32)(unsafe.Add(mBase, _consts[1489])) = int32(7533)
	return
}
func F_halfvec_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 float64
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 float64
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v23 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L131
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L127
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L124
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v26 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v29 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v30 != int32(701) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v35 = v26 - int32(1)
	v37 = v35 & int32(65535)
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v26&int32(65535) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v45 = v33
	goto L13
L13:
	;
	v47 = v21 + int32(8)
	v49 = v16 + int32(24)
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v49)))
	v52 = base.I32_extend16_s(v45)
	v54 = v52 + int32(1)
	v55 = F_mul_size(m, int32(4), v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v40 = int32(65535)
	if v35&v40 != v33&v40 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v45 = v35
	goto L13
L17:
	;
	goto L16
L18:
	;
	v57 = F_palloc(m, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v61 = F_Float8GetDatum(m, base.F64_add(v50, float64(1)))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v61
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v322 = F_construct_array(m, v57, v54, int32(701), int32(8), int32(0), int32(100))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L122
	}
L22:
	;
	v64 = int32(0)
	if v52 <= v64 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v52 <= int32(0) {
		goto L21
	} else {
		goto L74
	}
L25:
	;
	v67 = v64
	goto L26
L26:
	;
	v77 = int32(1)
	v78 = v67 + v77
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v49+v78<<(uint(int32(3))%32))))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v67<<(uint(v77)%32)))))
	v91 = v86 & int32(1023)
	v95 = v86 << (uint(int32(16)) % 32) & int32(-2147483648)
	v98 = int32(31)
	v99 = int32(base.Ui32(v86)>>(uint(int32(10))%32)) & v98
	if v99 != v98 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	goto L21
L28:
	;
	v179 = base.F64_add(v82, base.F64_promote_f32(base.F32_reinterpret_i32(v172|v171<<(uint(int32(13))%32))))
	if base.F64_eq(base.F64_abs(v179), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L71
	}
L29:
	;
	goto L28
L30:
	;
	v171 = v91
	v172 = v99<<(uint(int32(23))%32) + v95 + int32(939524096)
	goto L29
L31:
	;
	if v86&int32(512) != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v99 != 0 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v91 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v91 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v171 = int32(0)
	v172 = v95
	goto L29
L37:
	;
	v171 = int32(0)
	v172 = v95 | int32(2139095040)
	goto L29
L38:
	;
	goto L39
L39:
	;
	v171 = v91
	v172 = v95 | int32(2143289344)
	goto L29
L40:
	;
	v171 = v159 & int32(1022)
	v172 = v161 | v95
	goto L29
L41:
	;
	v159 = v91 << (uint(int32(1)) % 32)
	v161 = int32(939524096)
	goto L40
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(255)) < base.Ui32(v91) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v159 = v91 << (uint(int32(2)) % 32)
	v161 = int32(931135488)
	goto L40
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(127)) < base.Ui32(v91) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v159 = v91 << (uint(int32(3)) % 32)
	v161 = int32(922746880)
	goto L40
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(int32(63)) < base.Ui32(v91) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v159 = v91 << (uint(int32(4)) % 32)
	v161 = int32(914358272)
	goto L40
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(int32(31)) < base.Ui32(v91) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v159 = v91 << (uint(int32(5)) % 32)
	v161 = int32(905969664)
	goto L40
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(15)) < base.Ui32(v91) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v159 = v91 << (uint(int32(6)) % 32)
	v161 = int32(897581056)
	goto L40
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(int32(7)) < base.Ui32(v91) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v159 = v91 << (uint(int32(7)) % 32)
	v161 = int32(889192448)
	goto L40
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v159 = v91 << (uint(int32(8)) % 32)
	v161 = int32(880803840)
	goto L40
L63:
	;
	goto L64
L64:
	;
	v154 = base.B2i32(v91 == int32(1))
	if v91 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v155 = int32(1024)
	goto L67
L66:
	;
	v155 = v91 << (uint(int32(9)) % 32)
	goto L67
L67:
	;
	if v91 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v158 = int32(864026624)
	goto L70
L69:
	;
	v158 = int32(872415232)
	goto L70
L70:
	;
	v159 = v155
	v161 = v158
	goto L40
L71:
	;
	v186 = F_Float8GetDatum(m, v179)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57+v78<<(uint(int32(2))%32)))) = v186
	if v78 != v52 {
		v67 = v78
		goto L26
	} else {
		goto L73
	}
L73:
	;
	goto L27
L74:
	;
	v194 = int32(0)
	goto L75
L75:
	;
	v203 = int32(1)
	v204 = v194 + v203
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v194<<(uint(v203)%32)))))
	v216 = v211 & int32(1023)
	v220 = v211 << (uint(int32(16)) % 32) & int32(-2147483648)
	v223 = int32(31)
	v224 = int32(base.Ui32(v211)>>(uint(int32(10))%32)) & v223
	if v224 != v223 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L21
L77:
	;
	v304 = F_Float8GetDatum(m, base.F64_promote_f32(base.F32_reinterpret_i32(v297|v296<<(uint(int32(13))%32))))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L120
	}
L78:
	;
	goto L77
L79:
	;
	v296 = v216
	v297 = v224<<(uint(int32(23))%32) + v220 + int32(939524096)
	goto L78
L80:
	;
	if v211&int32(512) != 0 {
		goto L90
	} else {
		goto L91
	}
L81:
	;
	if v224 != 0 {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v216 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v216 != 0 {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v296 = int32(0)
	v297 = v220
	goto L78
L86:
	;
	v296 = int32(0)
	v297 = v220 | int32(2139095040)
	goto L78
L87:
	;
	goto L88
L88:
	;
	v296 = v216
	v297 = v220 | int32(2143289344)
	goto L78
L89:
	;
	v296 = v284 & int32(1022)
	v297 = v286 | v220
	goto L78
L90:
	;
	v284 = v216 << (uint(int32(1)) % 32)
	v286 = int32(939524096)
	goto L89
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(int32(255)) < base.Ui32(v216) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v284 = v216 << (uint(int32(2)) % 32)
	v286 = int32(931135488)
	goto L89
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(127)) < base.Ui32(v216) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v284 = v216 << (uint(int32(3)) % 32)
	v286 = int32(922746880)
	goto L89
L97:
	;
	goto L98
L98:
	;
	if base.Ui32(int32(63)) < base.Ui32(v216) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v284 = v216 << (uint(int32(4)) % 32)
	v286 = int32(914358272)
	goto L89
L100:
	;
	goto L101
L101:
	;
	if base.Ui32(int32(31)) < base.Ui32(v216) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v284 = v216 << (uint(int32(5)) % 32)
	v286 = int32(905969664)
	goto L89
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(15)) < base.Ui32(v216) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v284 = v216 << (uint(int32(6)) % 32)
	v286 = int32(897581056)
	goto L89
L106:
	;
	goto L107
L107:
	;
	if base.Ui32(int32(7)) < base.Ui32(v216) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v284 = v216 << (uint(int32(7)) % 32)
	v286 = int32(889192448)
	goto L89
L109:
	;
	goto L110
L110:
	;
	if base.Ui32(int32(3)) < base.Ui32(v216) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v284 = v216 << (uint(int32(8)) % 32)
	v286 = int32(880803840)
	goto L89
L112:
	;
	goto L113
L113:
	;
	v279 = base.B2i32(v216 == int32(1))
	if v216 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v280 = int32(1024)
	goto L116
L115:
	;
	v280 = v216 << (uint(int32(9)) % 32)
	goto L116
L116:
	;
	if v216 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v283 = int32(864026624)
	goto L119
L118:
	;
	v283 = int32(872415232)
	goto L119
L119:
	;
	v284 = v280
	v286 = v283
	goto L89
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57+v204<<(uint(int32(2))%32)))) = v304
	if v52 != v204 {
		v194 = v204
		goto L75
	} else {
		goto L121
	}
L121:
	;
	goto L76
L122:
	;
	F_pfree(m, v57)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	m.G0 = v13 + int32(32)
	return v322
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(293430)
	F_errmsg_internal(m, int32(26010), v13)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(512461), int32(173), int32(26501))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = base.I32_extend16_s(v35)
	F_errmsg(m, int32(477735), v13+int32(16))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(512461), int32(92), int32(294762))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 float32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 float32
	_ = v230
	var v237 int32
	_ = v237
	var v253 int32
	_ = v253
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v23 = base.B2i32(v21 < v22)
	if v21 < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v21 < v22 {
		goto L101
	} else {
		goto L102
	}
L5:
	;
	v24 = v21
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	if v24 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v27 = int32(8)
	v32 = int32(0)
	goto L9
L9:
	;
	v44 = v32 << (uint(int32(1)) % 32)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v27+v44))))
	v51 = v46 & int32(1023)
	v55 = v46 << (uint(int32(16)) % 32) & int32(-2147483648)
	v58 = int32(31)
	v59 = int32(base.Ui32(v46)>>(uint(int32(10))%32)) & v58
	if v59 != v58 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(1)
L11:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44+(v18+v27)))))
	v144 = v139 & int32(1023)
	v148 = v139 << (uint(int32(16)) % 32) & int32(-2147483648)
	v151 = int32(31)
	v152 = int32(base.Ui32(v139)>>(uint(int32(10))%32)) & v151
	if v152 != v151 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v137 = base.F32_reinterpret_i32(v132 | v131<<(uint(int32(13))%32))
	goto L11
L13:
	;
	v131 = v51
	v132 = v59<<(uint(int32(23))%32) + v55 + int32(939524096)
	goto L12
L14:
	;
	if v46&int32(512) != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v59 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v51 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v51 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v131 = int32(0)
	v132 = v55
	goto L12
L20:
	;
	v131 = int32(0)
	v132 = v55 | int32(2139095040)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v131 = v51
	v132 = v55 | int32(2143289344)
	goto L12
L23:
	;
	v131 = v119 & int32(1022)
	v132 = v121 | v55
	goto L12
L24:
	;
	v119 = v51 << (uint(int32(1)) % 32)
	v121 = int32(939524096)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v51) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v119 = v51 << (uint(int32(2)) % 32)
	v121 = int32(931135488)
	goto L23
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(127)) < base.Ui32(v51) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v119 = v51 << (uint(int32(3)) % 32)
	v121 = int32(922746880)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(63)) < base.Ui32(v51) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v119 = v51 << (uint(int32(4)) % 32)
	v121 = int32(914358272)
	goto L23
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(31)) < base.Ui32(v51) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = v51 << (uint(int32(5)) % 32)
	v121 = int32(905969664)
	goto L23
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(15)) < base.Ui32(v51) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v119 = v51 << (uint(int32(6)) % 32)
	v121 = int32(897581056)
	goto L23
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(7)) < base.Ui32(v51) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v119 = v51 << (uint(int32(7)) % 32)
	v121 = int32(889192448)
	goto L23
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(3)) < base.Ui32(v51) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v119 = v51 << (uint(int32(8)) % 32)
	v121 = int32(880803840)
	goto L23
L46:
	;
	goto L47
L47:
	;
	v114 = base.B2i32(v51 == int32(1))
	if v51 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v115 = int32(1024)
	goto L50
L49:
	;
	v115 = v51 << (uint(int32(9)) % 32)
	goto L50
L50:
	;
	if v51 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v118 = int32(864026624)
	goto L53
L52:
	;
	v118 = int32(872415232)
	goto L53
L53:
	;
	v119 = v115
	v121 = v118
	goto L23
L54:
	;
	if base.F32_gt(v137, v230)|base.F32_lt(v137, v230) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v230 = base.F32_reinterpret_i32(v225 | v224<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v224 = v144
	v225 = v152<<(uint(int32(23))%32) + v148 + int32(939524096)
	goto L55
L57:
	;
	if v139&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v152 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v144 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v144 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v224 = int32(0)
	v225 = v148
	goto L55
L63:
	;
	v224 = int32(0)
	v225 = v148 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v224 = v144
	v225 = v148 | int32(2143289344)
	goto L55
L66:
	;
	v224 = v212 & int32(1022)
	v225 = v214 | v148
	goto L55
L67:
	;
	v212 = v144 << (uint(int32(1)) % 32)
	v214 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v144) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v212 = v144 << (uint(int32(2)) % 32)
	v214 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v144) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v212 = v144 << (uint(int32(3)) % 32)
	v214 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v144) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v212 = v144 << (uint(int32(4)) % 32)
	v214 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v144) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v212 = v144 << (uint(int32(5)) % 32)
	v214 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v144) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v212 = v144 << (uint(int32(6)) % 32)
	v214 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v144) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v212 = v144 << (uint(int32(7)) % 32)
	v214 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v144) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v212 = v144 << (uint(int32(8)) % 32)
	v214 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v207 = base.B2i32(v144 == int32(1))
	if v144 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v208 = int32(1024)
	goto L93
L92:
	;
	v208 = v144 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v144 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v211 = int32(864026624)
	goto L96
L95:
	;
	v211 = int32(872415232)
	goto L96
L96:
	;
	v212 = v208
	v214 = v211
	goto L66
L97:
	;
	v237 = v32 + int32(1)
	if v24 != v237 {
		v32 = v237
		goto L9
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L10
L100:
	;
	goto L4
L101:
	;
	v253 = int32(1)
	goto L103
L102:
	;
	v253 = base.B2i32(v22 < v21)
	goto L103
L103:
	;
	return v253
}
func F_halfvec_to_vector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
		F_CheckDim_3(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
			if base.B2i32(v20 != int32(-1))&base.B2i32(v20 != v26) == int32(0) {
				v33 = F_mul_size(m, int32(4), v26)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = F_add_size(m, int32(8), v33)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = F_palloc0(m, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v26)
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v35 << (uint(int32(2)) % 32)
							v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
							if int32(0) < v43 {
								v46 = int32(8)
								v54 = int32(0)
								for {
									v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16+v46+v54<<(uint(int32(1))%32)))))
									v66 = v64 & int32(-2147483648)
									v68 = v64 & int32(1023)
									v71 = int32(31)
									v72 = int32(base.Ui32(v64)>>(uint(int32(10))%32)) & v71
									if v72 != v71 {
										if v72 != 0 {
											v143 = v72<<(uint(int32(23))%32) + v66 + int32(939524096)
											v144 = v68
										} else {
											if v68 != 0 {
												if v64&int32(512) != 0 {
													v132 = v68 << (uint(int32(1)) % 32)
													v134 = int32(939524096)
												} else {
													if base.Ui32(int32(255)) < base.Ui32(v68) {
														v132 = v68 << (uint(int32(2)) % 32)
														v134 = int32(931135488)
													} else {
														if base.Ui32(int32(127)) < base.Ui32(v68) {
															v132 = v68 << (uint(int32(3)) % 32)
															v134 = int32(922746880)
														} else {
															if base.Ui32(int32(63)) < base.Ui32(v68) {
																v132 = v68 << (uint(int32(4)) % 32)
																v134 = int32(914358272)
															} else {
																if base.Ui32(int32(31)) < base.Ui32(v68) {
																	v132 = v68 << (uint(int32(5)) % 32)
																	v134 = int32(905969664)
																} else {
																	if base.Ui32(int32(15)) < base.Ui32(v68) {
																		v132 = v68 << (uint(int32(6)) % 32)
																		v134 = int32(897581056)
																	} else {
																		if base.Ui32(int32(7)) < base.Ui32(v68) {
																			v132 = v68 << (uint(int32(7)) % 32)
																			v134 = int32(889192448)
																		} else {
																			if base.Ui32(int32(3)) < base.Ui32(v68) {
																				v132 = v68 << (uint(int32(8)) % 32)
																				v134 = int32(880803840)
																			} else {
																				v127 = base.B2i32(v68 == int32(1))
																				if v68 == int32(1) {
																					v128 = int32(1024)
																				} else {
																					v128 = v68 << (uint(int32(9)) % 32)
																				}
																				if v68 == int32(1) {
																					v131 = int32(864026624)
																				} else {
																					v131 = int32(872415232)
																				}
																				v132 = v128
																				v134 = v131
																			}
																		}
																	}
																}
															}
														}
													}
												}
												v143 = v134 | v66
												v144 = v132 & int32(1022)
											} else {
												v143 = v66
												v144 = int32(0)
											}
										}
									} else {
										if v68 == int32(0) {
											v143 = v66 | int32(2139095040)
											v144 = int32(0)
										} else {
											v143 = v66 | int32(2143289344)
											v144 = v68
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v37+v46+v54<<(uint(int32(2))%32)))) = v143 | v144<<(uint(int32(13))%32)
									v155 = v54 + int32(1)
									if v155 != v43 {
										v54 = v155
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							m.G0 = v13 + int32(16)
							return v37
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v20
						F_errmsg(m, int32(477735), v13)
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(506510), int32(88), int32(294762))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
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
}
func F_halfvec_typmod_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v19 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(575626), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(512461), int32(358), int32(285545))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
					if base.Ui32(int32(16001)) <= base.Ui32(v19) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(16000)
								F_errmsg(m, int32(489426), v5)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(512461), int32(363), int32(285545))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
						m.G0 = v5 + int32(16)
						return v19
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(226996), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512461), int32(353), int32(285545))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
}
func F_halfvec_vector_dims(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3)+4)))
		return v7
	}
}
