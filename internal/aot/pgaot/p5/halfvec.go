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
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 float64
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 float64
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
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
	v49 = v21 + int32(8)
	v51 = v16 + int32(24)
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	v54 = base.I32_extend16_s(v47)
	v56 = v54 + int32(1)
	v57 = F_mul_size(m, int32(4), v56)
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
	v59 = F_palloc(m, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v63 = F_Float8GetDatum(m, base.F64_add(v52, float64(1)))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v63
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v322 = F_construct_array(m, v59, v56, int32(701), int32(8), int32(0), int32(100))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L122
	}
L22:
	;
	v66 = int32(0)
	if v54 <= v66 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v54 <= int32(0) {
		goto L21
	} else {
		goto L74
	}
L25:
	;
	v69 = v66
	goto L26
L26:
	;
	v79 = int32(1)
	v80 = v69 + v79
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v51+v80<<(uint(int32(3))%32))))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+v69<<(uint(v79)%32)))))
	v93 = v88 & int32(1023)
	v97 = v88 << (uint(int32(16)) % 32) & int32(-2147483648)
	v100 = int32(31)
	v101 = int32(base.Ui32(v88)>>(uint(int32(10))%32)) & v100
	if v101 != v100 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	goto L21
L28:
	;
	v180 = base.F64_add(v84, base.F64_promote_f32(base.F32_reinterpret_i32(v174|v173<<(uint(int32(13))%32))))
	if base.F64_eq(base.F64_abs(v180), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L71
	}
L29:
	;
	goto L28
L30:
	;
	v173 = v93
	v174 = v101<<(uint(int32(23))%32) + v97 + int32(939524096)
	goto L29
L31:
	;
	if v88&int32(512) != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v101 != 0 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v93 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v93 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v173 = int32(0)
	v174 = v97
	goto L29
L37:
	;
	v173 = int32(0)
	v174 = v97 | int32(2139095040)
	goto L29
L38:
	;
	goto L39
L39:
	;
	v173 = v93
	v174 = v97 | int32(2143289344)
	goto L29
L40:
	;
	v173 = v161 & int32(1022)
	v174 = v163 | v97
	goto L29
L41:
	;
	v161 = v93 << (uint(int32(1)) % 32)
	v163 = int32(939524096)
	goto L40
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(255)) < base.Ui32(v93) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v161 = v93 << (uint(int32(2)) % 32)
	v163 = int32(931135488)
	goto L40
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(127)) < base.Ui32(v93) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v161 = v93 << (uint(int32(3)) % 32)
	v163 = int32(922746880)
	goto L40
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(int32(63)) < base.Ui32(v93) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v161 = v93 << (uint(int32(4)) % 32)
	v163 = int32(914358272)
	goto L40
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(int32(31)) < base.Ui32(v93) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v161 = v93 << (uint(int32(5)) % 32)
	v163 = int32(905969664)
	goto L40
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(15)) < base.Ui32(v93) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v161 = v93 << (uint(int32(6)) % 32)
	v163 = int32(897581056)
	goto L40
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(int32(7)) < base.Ui32(v93) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v161 = v93 << (uint(int32(7)) % 32)
	v163 = int32(889192448)
	goto L40
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(int32(3)) < base.Ui32(v93) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v161 = v93 << (uint(int32(8)) % 32)
	v163 = int32(880803840)
	goto L40
L63:
	;
	goto L64
L64:
	;
	v156 = base.B2i32(v93 == int32(1))
	if v93 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v157 = int32(1024)
	goto L67
L66:
	;
	v157 = v93 << (uint(int32(9)) % 32)
	goto L67
L67:
	;
	if v93 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v160 = int32(864026624)
	goto L70
L69:
	;
	v160 = int32(872415232)
	goto L70
L70:
	;
	v161 = v157
	v163 = v160
	goto L40
L71:
	;
	v187 = F_Float8GetDatum(m, v180)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+v80<<(uint(int32(2))%32)))) = v187
	if v80 != v54 {
		v69 = v80
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
	v204 = int32(1)
	v205 = v194 + v204
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+v194<<(uint(v204)%32)))))
	v217 = v212 & int32(1023)
	v221 = v212 << (uint(int32(16)) % 32) & int32(-2147483648)
	v224 = int32(31)
	v225 = int32(base.Ui32(v212)>>(uint(int32(10))%32)) & v224
	if v225 != v224 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L21
L77:
	;
	v304 = F_Float8GetDatum(m, base.F64_promote_f32(base.F32_reinterpret_i32(v298|v297<<(uint(int32(13))%32))))
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
	v297 = v217
	v298 = v225<<(uint(int32(23))%32) + v221 + int32(939524096)
	goto L78
L80:
	;
	if v212&int32(512) != 0 {
		goto L90
	} else {
		goto L91
	}
L81:
	;
	if v225 != 0 {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v217 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v217 != 0 {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v297 = int32(0)
	v298 = v221
	goto L78
L86:
	;
	v297 = int32(0)
	v298 = v221 | int32(2139095040)
	goto L78
L87:
	;
	goto L88
L88:
	;
	v297 = v217
	v298 = v221 | int32(2143289344)
	goto L78
L89:
	;
	v297 = v285 & int32(1022)
	v298 = v287 | v221
	goto L78
L90:
	;
	v285 = v217 << (uint(int32(1)) % 32)
	v287 = int32(939524096)
	goto L89
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(int32(255)) < base.Ui32(v217) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v285 = v217 << (uint(int32(2)) % 32)
	v287 = int32(931135488)
	goto L89
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(int32(127)) < base.Ui32(v217) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v285 = v217 << (uint(int32(3)) % 32)
	v287 = int32(922746880)
	goto L89
L97:
	;
	goto L98
L98:
	;
	if base.Ui32(int32(63)) < base.Ui32(v217) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v285 = v217 << (uint(int32(4)) % 32)
	v287 = int32(914358272)
	goto L89
L100:
	;
	goto L101
L101:
	;
	if base.Ui32(int32(31)) < base.Ui32(v217) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v285 = v217 << (uint(int32(5)) % 32)
	v287 = int32(905969664)
	goto L89
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(15)) < base.Ui32(v217) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v285 = v217 << (uint(int32(6)) % 32)
	v287 = int32(897581056)
	goto L89
L106:
	;
	goto L107
L107:
	;
	if base.Ui32(int32(7)) < base.Ui32(v217) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v285 = v217 << (uint(int32(7)) % 32)
	v287 = int32(889192448)
	goto L89
L109:
	;
	goto L110
L110:
	;
	if base.Ui32(int32(3)) < base.Ui32(v217) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v285 = v217 << (uint(int32(8)) % 32)
	v287 = int32(880803840)
	goto L89
L112:
	;
	goto L113
L113:
	;
	v280 = base.B2i32(v217 == int32(1))
	if v217 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v281 = int32(1024)
	goto L116
L115:
	;
	v281 = v217 << (uint(int32(9)) % 32)
	goto L116
L116:
	;
	if v217 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v284 = int32(864026624)
	goto L119
L118:
	;
	v284 = int32(872415232)
	goto L119
L119:
	;
	v285 = v281
	v287 = v284
	goto L89
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+v205<<(uint(int32(2))%32)))) = v304
	if v54 != v205 {
		v194 = v205
		goto L75
	} else {
		goto L121
	}
L121:
	;
	goto L76
L122:
	;
	F_pfree(m, v59)
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_halfvec_accum_1)
	F_errmsg_internal(m, int32(_a_F_halfvec_accum_2), v13)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_halfvec_accum_3), int32(173), int32(_a_F_halfvec_accum_4))
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
	F_errmsg(m, int32(_a_F_halfvec_accum_5), v13+int32(16))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_halfvec_accum_3), int32(92), int32(_a_F_halfvec_accum_6))
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return v21 | base.B2i32(v20 < v19)
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
	return int32(1)
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
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
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
							return v37
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v20
						F_errmsg(m, int32(_a_F_halfvec_to_vector_0), v13)
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_to_vector_1), int32(88), int32(_a_F_halfvec_to_vector_2))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn13934(m, l0, int32(_a_F_halfvec_typmod_in_0), int32(363), int32(_a_F_halfvec_typmod_in_1), int32(_a_F_halfvec_typmod_in_2), int32(_a_F_halfvec_typmod_in_3), int32(358), int32(_a_F_halfvec_typmod_in_4), int32(353), int32(_a_F_halfvec_typmod_in_5))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
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
