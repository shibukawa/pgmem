package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HalfvecInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_HalfvecInit[0])) = int32(_a_F_HalfvecInit_0)
	*(*int32)(unsafe.Add(mBase, _c_F_HalfvecInit[1])) = int32(_a_F_HalfvecInit_1)
	*(*int32)(unsafe.Add(mBase, _c_F_HalfvecInit[2])) = int32(_a_F_HalfvecInit_2)
	*(*int32)(unsafe.Add(mBase, _c_F_HalfvecInit[3])) = int32(_a_F_HalfvecInit_3)
	return
}
func F_halfvec_accum(m *base.Module, l0 int32) int64 {
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
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 float64
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 float64
	_ = v176
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L127
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L123
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L120
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
	v37 = v35 & int32(_a_F_halfvec_accum_0)
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = int32(_a_F_halfvec_accum_0)
	if v35&v38 != v33&v38 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v47 = v33
	goto L13
L13:
	;
	v48 = int32(8)
	v49 = v21 + v48
	v51 = v16 + int32(24)
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	v54 = base.I32_extend16_s(v47)
	v56 = v54 + int32(1)
	v57 = F_palloc_mul(m, v48, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v46 = v26 & v38
	goto L16
L15:
	;
	v46 = int32(0)
	goto L16
L16:
	;
	if v46 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v47 = v35
	goto L13
L18:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v57))) = base.F64_add(v52, float64(1))
	if v37 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v312 = F_construct_array(m, v57, v56, int32(701), int32(8), int32(1), int32(100))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L118
	}
L20:
	;
	v62 = int32(0)
	if v54 <= v62 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v54 <= int32(0) {
		goto L19
	} else {
		goto L71
	}
L23:
	;
	v65 = v62
	goto L24
L24:
	;
	v75 = int32(1)
	v76 = v65 + v75
	v78 = v76 << (uint(int32(3)) % 32)
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v51+v78)))
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+v65<<(uint(v75)%32)))))
	v89 = v84 & int32(1023)
	v93 = v84 << (uint(int32(16)) % 32) & int32(-2147483648)
	v96 = int32(31)
	v97 = int32(base.Ui32(v84)>>(uint(int32(10))%32)) & v96
	if v97 != v96 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	goto L19
L26:
	;
	v176 = base.F64_add(v80, base.F64_promote_f32(base.F32_reinterpret_i32(v170|v169<<(uint(int32(13))%32))))
	if base.F64_eq(base.F64_abs(v176), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L69
	}
L27:
	;
	goto L26
L28:
	;
	v169 = v89
	v170 = v97<<(uint(int32(23))%32) + v93 + int32(939524096)
	goto L27
L29:
	;
	if v84&int32(512) != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if v97 != 0 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v89 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v89 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v169 = int32(0)
	v170 = v93
	goto L27
L35:
	;
	v169 = int32(0)
	v170 = v93 | int32(2139095040)
	goto L27
L36:
	;
	goto L37
L37:
	;
	v169 = v89
	v170 = v93 | int32(2143289344)
	goto L27
L38:
	;
	v169 = v157 & int32(1022)
	v170 = v159 | v93
	goto L27
L39:
	;
	v157 = v89 << (uint(int32(1)) % 32)
	v159 = int32(939524096)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(255)) < base.Ui32(v89) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v157 = v89 << (uint(int32(2)) % 32)
	v159 = int32(931135488)
	goto L38
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(127)) < base.Ui32(v89) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v157 = v89 << (uint(int32(3)) % 32)
	v159 = int32(922746880)
	goto L38
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(int32(63)) < base.Ui32(v89) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v157 = v89 << (uint(int32(4)) % 32)
	v159 = int32(914358272)
	goto L38
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(int32(31)) < base.Ui32(v89) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v157 = v89 << (uint(int32(5)) % 32)
	v159 = int32(905969664)
	goto L38
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(15)) < base.Ui32(v89) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v157 = v89 << (uint(int32(6)) % 32)
	v159 = int32(897581056)
	goto L38
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(int32(7)) < base.Ui32(v89) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v157 = v89 << (uint(int32(7)) % 32)
	v159 = int32(889192448)
	goto L38
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(int32(3)) < base.Ui32(v89) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v157 = v89 << (uint(int32(8)) % 32)
	v159 = int32(880803840)
	goto L38
L61:
	;
	goto L62
L62:
	;
	v152 = base.B2i32(v89 == int32(1))
	if v89 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v153 = int32(1024)
	goto L65
L64:
	;
	v153 = v89 << (uint(int32(9)) % 32)
	goto L65
L65:
	;
	if v89 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v156 = int32(864026624)
	goto L68
L67:
	;
	v156 = int32(872415232)
	goto L68
L68:
	;
	v157 = v153
	v159 = v156
	goto L38
L69:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v57+v78))) = v176
	if v54 != v76 {
		v65 = v76
		goto L24
	} else {
		goto L70
	}
L70:
	;
	goto L25
L71:
	;
	v186 = int32(0)
	goto L72
L72:
	;
	v196 = int32(1)
	v197 = v186 + v196
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+v186<<(uint(v196)%32)))))
	v209 = v204 & int32(1023)
	v213 = v204 << (uint(int32(16)) % 32) & int32(-2147483648)
	v216 = int32(31)
	v217 = int32(base.Ui32(v204)>>(uint(int32(10))%32)) & v216
	if v217 != v216 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	goto L19
L74:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v57+v197<<(uint(int32(3))%32)))) = base.F64_promote_f32(base.F32_reinterpret_i32(v290 | v289<<(uint(int32(13))%32)))
	if v54 != v197 {
		v186 = v197
		goto L72
	} else {
		goto L117
	}
L75:
	;
	goto L74
L76:
	;
	v289 = v209
	v290 = v217<<(uint(int32(23))%32) + v213 + int32(939524096)
	goto L75
L77:
	;
	if v204&int32(512) != 0 {
		goto L87
	} else {
		goto L88
	}
L78:
	;
	if v217 != 0 {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v209 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v209 != 0 {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v289 = int32(0)
	v290 = v213
	goto L75
L83:
	;
	v289 = int32(0)
	v290 = v213 | int32(2139095040)
	goto L75
L84:
	;
	goto L85
L85:
	;
	v289 = v209
	v290 = v213 | int32(2143289344)
	goto L75
L86:
	;
	v289 = v277 & int32(1022)
	v290 = v279 | v213
	goto L75
L87:
	;
	v277 = v209 << (uint(int32(1)) % 32)
	v279 = int32(939524096)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(int32(255)) < base.Ui32(v209) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v277 = v209 << (uint(int32(2)) % 32)
	v279 = int32(931135488)
	goto L86
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(int32(127)) < base.Ui32(v209) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v277 = v209 << (uint(int32(3)) % 32)
	v279 = int32(922746880)
	goto L86
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(63)) < base.Ui32(v209) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v277 = v209 << (uint(int32(4)) % 32)
	v279 = int32(914358272)
	goto L86
L97:
	;
	goto L98
L98:
	;
	if base.Ui32(int32(31)) < base.Ui32(v209) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v277 = v209 << (uint(int32(5)) % 32)
	v279 = int32(905969664)
	goto L86
L100:
	;
	goto L101
L101:
	;
	if base.Ui32(int32(15)) < base.Ui32(v209) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v277 = v209 << (uint(int32(6)) % 32)
	v279 = int32(897581056)
	goto L86
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(7)) < base.Ui32(v209) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v277 = v209 << (uint(int32(7)) % 32)
	v279 = int32(889192448)
	goto L86
L106:
	;
	goto L107
L107:
	;
	if base.Ui32(int32(3)) < base.Ui32(v209) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v277 = v209 << (uint(int32(8)) % 32)
	v279 = int32(880803840)
	goto L86
L109:
	;
	goto L110
L110:
	;
	v272 = base.B2i32(v209 == int32(1))
	if v209 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v273 = int32(1024)
	goto L113
L112:
	;
	v273 = v209 << (uint(int32(9)) % 32)
	goto L113
L113:
	;
	if v209 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v276 = int32(864026624)
	goto L116
L115:
	;
	v276 = int32(872415232)
	goto L116
L116:
	;
	v277 = v273
	v279 = v276
	goto L86
L117:
	;
	goto L73
L118:
	;
	F_pfree(m, v57)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	m.G0 = v13 + int32(32)
	return base.I64_extend_i32_u(v312)
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_halfvec_accum_1)
	F_errmsg_internal(m, int32(_a_F_halfvec_accum_2), v13)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_halfvec_accum_3), int32(173), int32(_a_F_halfvec_accum_4))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = base.I32_extend16_s(v35)
	F_errmsg(m, int32(_a_F_halfvec_accum_5), v13+int32(16))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_halfvec_accum_3), int32(92), int32(_a_F_halfvec_accum_6))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
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
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 float32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 float32
	_ = v225
	var v232 int32
	_ = v232
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = base.B2i32(v19 < v20)
	if v19 < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v19 < v20 {
		goto L101
	} else {
		goto L102
	}
L5:
	;
	v22 = v19
	goto L7
L6:
	;
	v22 = v20
	goto L7
L7:
	;
	if v22 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = v30 << (uint(int32(1)) % 32)
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v25+v41))))
	v48 = v43 & int32(1023)
	v52 = v43 << (uint(int32(16)) % 32) & int32(-2147483648)
	v55 = int32(31)
	v56 = int32(base.Ui32(v43)>>(uint(int32(10))%32)) & v55
	if v56 != v55 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int64(1)
L11:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v25+v41))))
	v140 = v135 & int32(1023)
	v144 = v135 << (uint(int32(16)) % 32) & int32(-2147483648)
	v147 = int32(31)
	v148 = int32(base.Ui32(v135)>>(uint(int32(10))%32)) & v147
	if v148 != v147 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v133 = base.F32_reinterpret_i32(v129 | v128<<(uint(int32(13))%32))
	goto L11
L13:
	;
	v128 = v48
	v129 = v56<<(uint(int32(23))%32) + v52 + int32(939524096)
	goto L12
L14:
	;
	if v43&int32(512) != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v56 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v48 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v128 = int32(0)
	v129 = v52
	goto L12
L20:
	;
	v128 = int32(0)
	v129 = v52 | int32(2139095040)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v128 = v48
	v129 = v52 | int32(2143289344)
	goto L12
L23:
	;
	v128 = v116 & int32(1022)
	v129 = v118 | v52
	goto L12
L24:
	;
	v116 = v48 << (uint(int32(1)) % 32)
	v118 = int32(939524096)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v48) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = v48 << (uint(int32(2)) % 32)
	v118 = int32(931135488)
	goto L23
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(127)) < base.Ui32(v48) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v116 = v48 << (uint(int32(3)) % 32)
	v118 = int32(922746880)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(63)) < base.Ui32(v48) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = v48 << (uint(int32(4)) % 32)
	v118 = int32(914358272)
	goto L23
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(31)) < base.Ui32(v48) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v116 = v48 << (uint(int32(5)) % 32)
	v118 = int32(905969664)
	goto L23
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(15)) < base.Ui32(v48) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v116 = v48 << (uint(int32(6)) % 32)
	v118 = int32(897581056)
	goto L23
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(7)) < base.Ui32(v48) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v116 = v48 << (uint(int32(7)) % 32)
	v118 = int32(889192448)
	goto L23
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(3)) < base.Ui32(v48) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v116 = v48 << (uint(int32(8)) % 32)
	v118 = int32(880803840)
	goto L23
L46:
	;
	goto L47
L47:
	;
	v111 = base.B2i32(v48 == int32(1))
	if v48 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v112 = int32(1024)
	goto L50
L49:
	;
	v112 = v48 << (uint(int32(9)) % 32)
	goto L50
L50:
	;
	if v48 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v115 = int32(864026624)
	goto L53
L52:
	;
	v115 = int32(872415232)
	goto L53
L53:
	;
	v116 = v112
	v118 = v115
	goto L23
L54:
	;
	if base.F32_gt(v133, v225)|base.F32_lt(v133, v225) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v225 = base.F32_reinterpret_i32(v221 | v220<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v220 = v140
	v221 = v148<<(uint(int32(23))%32) + v144 + int32(939524096)
	goto L55
L57:
	;
	if v135&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v148 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v140 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v140 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v220 = int32(0)
	v221 = v144
	goto L55
L63:
	;
	v220 = int32(0)
	v221 = v144 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v220 = v140
	v221 = v144 | int32(2143289344)
	goto L55
L66:
	;
	v220 = v208 & int32(1022)
	v221 = v210 | v144
	goto L55
L67:
	;
	v208 = v140 << (uint(int32(1)) % 32)
	v210 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v140) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v208 = v140 << (uint(int32(2)) % 32)
	v210 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v140) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v208 = v140 << (uint(int32(3)) % 32)
	v210 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v140) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v208 = v140 << (uint(int32(4)) % 32)
	v210 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v140) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v208 = v140 << (uint(int32(5)) % 32)
	v210 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v140) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v208 = v140 << (uint(int32(6)) % 32)
	v210 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v140) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v208 = v140 << (uint(int32(7)) % 32)
	v210 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v208 = v140 << (uint(int32(8)) % 32)
	v210 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v203 = base.B2i32(v140 == int32(1))
	if v140 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v204 = int32(1024)
	goto L93
L92:
	;
	v204 = v140 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v140 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v207 = int32(864026624)
	goto L96
L95:
	;
	v207 = int32(872415232)
	goto L96
L96:
	;
	v208 = v204
	v210 = v207
	goto L66
L97:
	;
	v232 = v30 + int32(1)
	if v22 != v232 {
		v30 = v232
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
	return int64(1)
L102:
	;
	goto L103
L103:
	;
	return base.I64_extend_i32_u(base.B2i32(v20 < v19))
}
func F_halfvec_to_vector(m *base.Module, l0 int32) int64 {
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
	var v53 int32
	_ = v53
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
	var v154 int32
	_ = v154
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
		F_CheckDim_3(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
			if base.B2i32(v20 != int32(-1))&base.B2i32(v20 != v26) == int32(0) {
				v33 = F_mul_size(m, int32(4), v26)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = F_add_size(m, int32(8), v33)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						v37 = F_palloc0(m, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v26)
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v35 << (uint(int32(2)) % 32)
							v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
							if int32(0) < v43 {
								v46 = int32(8)
								v53 = int32(0)
								for {
									v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16+v46+v53<<(uint(int32(1))%32)))))
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
									*(*int32)(unsafe.Add(mBase, uint32(v37+v46+v53<<(uint(int32(2))%32)))) = v143 | v144<<(uint(int32(13))%32)
									v154 = v53 + int32(1)
									if v154 != v43 {
										v53 = v154
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							m.G0 = v13 + int32(16)
							return base.I64_extend_i32_u(v37)
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v20
						F_errmsg(m, int32(_a_F_halfvec_to_vector_0), v13)
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_to_vector_1), int32(88), int32(_a_F_halfvec_to_vector_2))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int64(0)
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
func F_halfvec_typmod_in(m *base.Module, l0 int32) int64 {
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn14301(m, l0, int32(_a_F_halfvec_typmod_in_0), int32(363), int32(_a_F_halfvec_typmod_in_1), int32(_a_F_halfvec_typmod_in_2), int32(_a_F_halfvec_typmod_in_3), int32(358), int32(_a_F_halfvec_typmod_in_4), int32(353), int32(_a_F_halfvec_typmod_in_5))
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		return v11
	}
}
func F_halfvec_vector_dims(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = int64(*(*int16)(unsafe.Add(mBase, uint32(v3)+4)))
		return v7
	}
}
