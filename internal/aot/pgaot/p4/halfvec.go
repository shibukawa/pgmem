package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_halfvec_mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 float32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 float32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v423 int32
	_ = v423
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	if v27 == v28 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L139
	}
L5:
	;
	m.G0 = v17 + int32(16)
	return v37
L6:
	;
	v363 = int32(0)
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v364 <= v363 {
		goto L5
	} else {
		goto L127
	}
L7:
	;
	v33 = F_mul_size(m, int32(2), base.I32_extend16_s(v27))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L123
	}
L10:
	;
	v35 = F_add_size(m, int32(8), v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v37 = F_palloc0(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v35 << (uint(int32(2)) % 32)
	v43 = int32(0)
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v44 <= v43 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(8)
	v48 = v20 + v47
	v50 = v25 + v47
	v52 = v37 + v47
	v53 = v43
	goto L14
L14:
	;
	v68 = v53 << (uint(int32(1)) % 32)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+v68))))
	v75 = v70 & int32(1023)
	v79 = v70 << (uint(int32(16)) % 32) & int32(-2147483648)
	v82 = int32(31)
	v83 = int32(base.Ui32(v70)>>(uint(int32(10))%32)) & v82
	if v83 != v82 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	goto L6
L16:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+v50))))
	v167 = v162 & int32(1023)
	v171 = v162 << (uint(int32(16)) % 32) & int32(-2147483648)
	v174 = int32(31)
	v175 = int32(base.Ui32(v162)>>(uint(int32(10))%32)) & v174
	if v175 != v174 {
		goto L63
	} else {
		goto L64
	}
L17:
	;
	goto L16
L18:
	;
	v155 = v75
	v156 = v83<<(uint(int32(23))%32) + v79 + int32(939524096)
	goto L17
L19:
	;
	if v70&int32(512) != 0 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	if v83 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v75 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v75 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v155 = int32(0)
	v156 = v79
	goto L17
L25:
	;
	v155 = int32(0)
	v156 = v79 | int32(2139095040)
	goto L17
L26:
	;
	goto L27
L27:
	;
	v155 = v75
	v156 = v79 | int32(2143289344)
	goto L17
L28:
	;
	v155 = v143 & int32(1022)
	v156 = v145 | v79
	goto L17
L29:
	;
	v143 = v75 << (uint(int32(1)) % 32)
	v145 = int32(939524096)
	goto L28
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(255)) < base.Ui32(v75) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v143 = v75 << (uint(int32(2)) % 32)
	v145 = int32(931135488)
	goto L28
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(int32(127)) < base.Ui32(v75) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v143 = v75 << (uint(int32(3)) % 32)
	v145 = int32(922746880)
	goto L28
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(int32(63)) < base.Ui32(v75) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v143 = v75 << (uint(int32(4)) % 32)
	v145 = int32(914358272)
	goto L28
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(int32(31)) < base.Ui32(v75) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v143 = v75 << (uint(int32(5)) % 32)
	v145 = int32(905969664)
	goto L28
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(15)) < base.Ui32(v75) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v143 = v75 << (uint(int32(6)) % 32)
	v145 = int32(897581056)
	goto L28
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(7)) < base.Ui32(v75) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v143 = v75 << (uint(int32(7)) % 32)
	v145 = int32(889192448)
	goto L28
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(int32(3)) < base.Ui32(v75) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v143 = v75 << (uint(int32(8)) % 32)
	v145 = int32(880803840)
	goto L28
L51:
	;
	goto L52
L52:
	;
	v138 = base.B2i32(v75 == int32(1))
	if v75 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v139 = int32(1024)
	goto L55
L54:
	;
	v139 = v75 << (uint(int32(9)) % 32)
	goto L55
L55:
	;
	if v75 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v142 = int32(864026624)
	goto L58
L57:
	;
	v142 = int32(872415232)
	goto L58
L58:
	;
	v143 = v139
	v145 = v142
	goto L28
L59:
	;
	v253 = base.F32_mul(base.F32_reinterpret_i32(v156|v155<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v248|v247<<(uint(int32(13))%32)))
	v254 = base.I32_reinterpret_f32(v253)
	v256 = int32(base.Ui32(v254) >> (uint(int32(16)) % 32))
	v260 = base.F32_abs(v253)
	if base.F32_eq(v260, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v339 = v256 & int32(_a_F_halfvec_mul_0)
		goto L102
	} else {
		goto L103
	}
L60:
	;
	goto L59
L61:
	;
	v247 = v167
	v248 = v175<<(uint(int32(23))%32) + v171 + int32(939524096)
	goto L60
L62:
	;
	if v162&int32(512) != 0 {
		goto L72
	} else {
		goto L73
	}
L63:
	;
	if v175 != 0 {
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v167 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v167 != 0 {
		goto L62
	} else {
		goto L67
	}
L67:
	;
	v247 = int32(0)
	v248 = v171
	goto L60
L68:
	;
	v247 = int32(0)
	v248 = v171 | int32(2139095040)
	goto L60
L69:
	;
	goto L70
L70:
	;
	v247 = v167
	v248 = v171 | int32(2143289344)
	goto L60
L71:
	;
	v247 = v235 & int32(1022)
	v248 = v237 | v171
	goto L60
L72:
	;
	v235 = v167 << (uint(int32(1)) % 32)
	v237 = int32(939524096)
	goto L71
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(255)) < base.Ui32(v167) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v235 = v167 << (uint(int32(2)) % 32)
	v237 = int32(931135488)
	goto L71
L76:
	;
	goto L77
L77:
	;
	if base.Ui32(int32(127)) < base.Ui32(v167) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v235 = v167 << (uint(int32(3)) % 32)
	v237 = int32(922746880)
	goto L71
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(int32(63)) < base.Ui32(v167) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v235 = v167 << (uint(int32(4)) % 32)
	v237 = int32(914358272)
	goto L71
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(int32(31)) < base.Ui32(v167) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v235 = v167 << (uint(int32(5)) % 32)
	v237 = int32(905969664)
	goto L71
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(int32(15)) < base.Ui32(v167) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v235 = v167 << (uint(int32(6)) % 32)
	v237 = int32(897581056)
	goto L71
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(int32(7)) < base.Ui32(v167) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v235 = v167 << (uint(int32(7)) % 32)
	v237 = int32(889192448)
	goto L71
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(int32(3)) < base.Ui32(v167) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v235 = v167 << (uint(int32(8)) % 32)
	v237 = int32(880803840)
	goto L71
L94:
	;
	goto L95
L95:
	;
	v230 = base.B2i32(v167 == int32(1))
	if v167 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v231 = int32(1024)
	goto L98
L97:
	;
	v231 = v167 << (uint(int32(9)) % 32)
	goto L98
L98:
	;
	if v167 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v234 = int32(864026624)
	goto L101
L100:
	;
	v234 = int32(872415232)
	goto L101
L101:
	;
	v235 = v231
	v237 = v234
	goto L71
L102:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v68+v52))) = uint16(v339)
	v342 = v53 + int32(1)
	if v44 != v342 {
		v53 = v342
		goto L14
	} else {
		goto L122
	}
L103:
	;
	v264 = v256 & int32(_a_F_halfvec_mul_1)
	v266 = v254 & int32(_a_F_halfvec_mul_2)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v260)) {
		v339 = v264 | int32(base.Ui32(v266)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_mul_3)
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v278 = int32(base.Ui32(v254)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v278) < base.Ui32(int32(99)) {
		v339 = v264
		goto L102
	} else {
		goto L105
	}
L105:
	;
	if base.Ui32(v278) <= base.Ui32(int32(112)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v290 = int32(1)<<(uint(v278-int32(90))%32) + int32(base.Ui32(v266)>>(uint(int32(113)-v278)%32))
	v292 = v290 | v254
	v293 = v290
	goto L108
L107:
	;
	v292 = v254
	v293 = v266
	goto L108
L108:
	;
	v295 = int32(base.Ui32(v293) >> (uint(int32(13)) % 32))
	v300 = int32(1)
	v304 = int32(3)
	v305 = int32(base.Ui32(v293)>>(uint(int32(12))%32)) & v304
	if base.B2i32(v305 != v304)&(base.B2i32(v292&int32(4095) == int32(0))|base.B2i32(v305 != v300)) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v316 = v295
	goto L111
L110:
	;
	v316 = v295 + v300
	goto L111
L111:
	;
	v318 = base.B2i32(v316 == int32(1024))
	if v316 == int32(1024) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v319 = int32(-126)
	goto L114
L113:
	;
	v319 = int32(-127)
	goto L114
L114:
	;
	v320 = v319 + v278
	if int32(16) <= v320 {
		v339 = v264 | int32(_a_F_halfvec_mul_4)
		goto L102
	} else {
		goto L115
	}
L115:
	;
	if int32(-15) < v320 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v330 = v320<<(uint(int32(10))%32) + int32(_a_F_halfvec_mul_5) | v264
	goto L118
L117:
	;
	v330 = v264
	goto L118
L118:
	;
	if v316 == int32(1024) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v332 = int32(0)
	goto L121
L120:
	;
	v332 = v316
	goto L121
L121:
	;
	v339 = v330 | v332
	goto L102
L122:
	;
	goto L15
L123:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v351
	F_errmsg(m, int32(_a_F_halfvec_mul_6), v17)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_halfvec_mul_7), int32(80), int32(_a_F_halfvec_mul_8))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
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
	v367 = v363
	goto L128
L128:
	;
	v382 = v367 << (uint(int32(1)) % 32)
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v382))))
	v386 = v384 & int32(_a_F_halfvec_mul_9)
	if v386 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L5
L130:
	;
	v402 = v367 + int32(1)
	if v402 != v364 {
		v367 = v402
		goto L128
	} else {
		goto L138
	}
L131:
	;
	if v386 != int32(_a_F_halfvec_mul_4) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382+v48))))
	if v392&int32(_a_F_halfvec_mul_9) == int32(0) {
		goto L130
	} else {
		goto L136
	}
L134:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382+v50))))
	if v398&int32(_a_F_halfvec_mul_9) != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	goto L130
L138:
	;
	goto L129
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_sub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 float32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 float32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	if v27 == v28 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v17 + int32(16)
	return v37
L5:
	;
	v363 = int32(0)
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v364 <= v363 {
		goto L4
	} else {
		goto L126
	}
L6:
	;
	v33 = F_mul_size(m, int32(2), base.I32_extend16_s(v27))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L122
	}
L9:
	;
	v35 = F_add_size(m, int32(8), v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v37 = F_palloc0(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v35 << (uint(int32(2)) % 32)
	v43 = int32(0)
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v44 <= v43 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v47 = int32(8)
	v52 = v37 + v47
	v53 = v43
	goto L13
L13:
	;
	v68 = v53 << (uint(int32(1)) % 32)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v47+v68))))
	v75 = v70 & int32(1023)
	v79 = v70 << (uint(int32(16)) % 32) & int32(-2147483648)
	v82 = int32(31)
	v83 = int32(base.Ui32(v70)>>(uint(int32(10))%32)) & v82
	if v83 != v82 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	goto L5
L15:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+(v25+v47)))))
	v167 = v162 & int32(1023)
	v171 = v162 << (uint(int32(16)) % 32) & int32(-2147483648)
	v174 = int32(31)
	v175 = int32(base.Ui32(v162)>>(uint(int32(10))%32)) & v174
	if v175 != v174 {
		goto L62
	} else {
		goto L63
	}
L16:
	;
	goto L15
L17:
	;
	v155 = v75
	v156 = v83<<(uint(int32(23))%32) + v79 + int32(939524096)
	goto L16
L18:
	;
	if v70&int32(512) != 0 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	if v83 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v75 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v75 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v155 = int32(0)
	v156 = v79
	goto L16
L24:
	;
	v155 = int32(0)
	v156 = v79 | int32(2139095040)
	goto L16
L25:
	;
	goto L26
L26:
	;
	v155 = v75
	v156 = v79 | int32(2143289344)
	goto L16
L27:
	;
	v155 = v143 & int32(1022)
	v156 = v145 | v79
	goto L16
L28:
	;
	v143 = v75 << (uint(int32(1)) % 32)
	v145 = int32(939524096)
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(int32(255)) < base.Ui32(v75) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v143 = v75 << (uint(int32(2)) % 32)
	v145 = int32(931135488)
	goto L27
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(int32(127)) < base.Ui32(v75) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v143 = v75 << (uint(int32(3)) % 32)
	v145 = int32(922746880)
	goto L27
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(int32(63)) < base.Ui32(v75) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v143 = v75 << (uint(int32(4)) % 32)
	v145 = int32(914358272)
	goto L27
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(int32(31)) < base.Ui32(v75) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v143 = v75 << (uint(int32(5)) % 32)
	v145 = int32(905969664)
	goto L27
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(15)) < base.Ui32(v75) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v143 = v75 << (uint(int32(6)) % 32)
	v145 = int32(897581056)
	goto L27
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(7)) < base.Ui32(v75) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v143 = v75 << (uint(int32(7)) % 32)
	v145 = int32(889192448)
	goto L27
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(int32(3)) < base.Ui32(v75) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v143 = v75 << (uint(int32(8)) % 32)
	v145 = int32(880803840)
	goto L27
L50:
	;
	goto L51
L51:
	;
	v138 = base.B2i32(v75 == int32(1))
	if v75 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v139 = int32(1024)
	goto L54
L53:
	;
	v139 = v75 << (uint(int32(9)) % 32)
	goto L54
L54:
	;
	if v75 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v142 = int32(864026624)
	goto L57
L56:
	;
	v142 = int32(872415232)
	goto L57
L57:
	;
	v143 = v139
	v145 = v142
	goto L27
L58:
	;
	v253 = base.F32_sub(base.F32_reinterpret_i32(v156|v155<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v248|v247<<(uint(int32(13))%32)))
	v254 = base.I32_reinterpret_f32(v253)
	v256 = int32(base.Ui32(v254) >> (uint(int32(16)) % 32))
	v260 = base.F32_abs(v253)
	if base.F32_eq(v260, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v339 = v256 & int32(_a_F_halfvec_sub_0)
		goto L101
	} else {
		goto L102
	}
L59:
	;
	goto L58
L60:
	;
	v247 = v167
	v248 = v175<<(uint(int32(23))%32) + v171 + int32(939524096)
	goto L59
L61:
	;
	if v162&int32(512) != 0 {
		goto L71
	} else {
		goto L72
	}
L62:
	;
	if v175 != 0 {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v167 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v167 != 0 {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v247 = int32(0)
	v248 = v171
	goto L59
L67:
	;
	v247 = int32(0)
	v248 = v171 | int32(2139095040)
	goto L59
L68:
	;
	goto L69
L69:
	;
	v247 = v167
	v248 = v171 | int32(2143289344)
	goto L59
L70:
	;
	v247 = v235 & int32(1022)
	v248 = v237 | v171
	goto L59
L71:
	;
	v235 = v167 << (uint(int32(1)) % 32)
	v237 = int32(939524096)
	goto L70
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(255)) < base.Ui32(v167) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v235 = v167 << (uint(int32(2)) % 32)
	v237 = int32(931135488)
	goto L70
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(127)) < base.Ui32(v167) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v235 = v167 << (uint(int32(3)) % 32)
	v237 = int32(922746880)
	goto L70
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(63)) < base.Ui32(v167) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v235 = v167 << (uint(int32(4)) % 32)
	v237 = int32(914358272)
	goto L70
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(31)) < base.Ui32(v167) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v235 = v167 << (uint(int32(5)) % 32)
	v237 = int32(905969664)
	goto L70
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(15)) < base.Ui32(v167) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v235 = v167 << (uint(int32(6)) % 32)
	v237 = int32(897581056)
	goto L70
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(7)) < base.Ui32(v167) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v235 = v167 << (uint(int32(7)) % 32)
	v237 = int32(889192448)
	goto L70
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(int32(3)) < base.Ui32(v167) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v235 = v167 << (uint(int32(8)) % 32)
	v237 = int32(880803840)
	goto L70
L93:
	;
	goto L94
L94:
	;
	v230 = base.B2i32(v167 == int32(1))
	if v167 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v231 = int32(1024)
	goto L97
L96:
	;
	v231 = v167 << (uint(int32(9)) % 32)
	goto L97
L97:
	;
	if v167 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v234 = int32(864026624)
	goto L100
L99:
	;
	v234 = int32(872415232)
	goto L100
L100:
	;
	v235 = v231
	v237 = v234
	goto L70
L101:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v68+v52))) = uint16(v339)
	v342 = v53 + int32(1)
	if v44 != v342 {
		v53 = v342
		goto L13
	} else {
		goto L121
	}
L102:
	;
	v264 = v256 & int32(_a_F_halfvec_sub_1)
	v266 = v254 & int32(_a_F_halfvec_sub_2)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v260)) {
		v339 = v264 | int32(base.Ui32(v266)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_sub_3)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v278 = int32(base.Ui32(v254)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v278) < base.Ui32(int32(99)) {
		v339 = v264
		goto L101
	} else {
		goto L104
	}
L104:
	;
	if base.Ui32(v278) <= base.Ui32(int32(112)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v290 = int32(1)<<(uint(v278-int32(90))%32) + int32(base.Ui32(v266)>>(uint(int32(113)-v278)%32))
	v292 = v290 | v254
	v293 = v290
	goto L107
L106:
	;
	v292 = v254
	v293 = v266
	goto L107
L107:
	;
	v295 = int32(base.Ui32(v293) >> (uint(int32(13)) % 32))
	v300 = int32(1)
	v304 = int32(3)
	v305 = int32(base.Ui32(v293)>>(uint(int32(12))%32)) & v304
	if base.B2i32(v305 != v304)&(base.B2i32(v292&int32(4095) == int32(0))|base.B2i32(v305 != v300)) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v316 = v295
	goto L110
L109:
	;
	v316 = v295 + v300
	goto L110
L110:
	;
	v318 = base.B2i32(v316 == int32(1024))
	if v316 == int32(1024) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v319 = int32(-126)
	goto L113
L112:
	;
	v319 = int32(-127)
	goto L113
L113:
	;
	v320 = v319 + v278
	if int32(16) <= v320 {
		v339 = v264 | int32(_a_F_halfvec_sub_4)
		goto L101
	} else {
		goto L114
	}
L114:
	;
	if int32(-15) < v320 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v330 = v320<<(uint(int32(10))%32) + int32(_a_F_halfvec_sub_5) | v264
	goto L117
L116:
	;
	v330 = v264
	goto L117
L117:
	;
	if v316 == int32(1024) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v332 = int32(0)
	goto L120
L119:
	;
	v332 = v316
	goto L120
L120:
	;
	v339 = v330 | v332
	goto L101
L121:
	;
	goto L14
L122:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v351
	F_errmsg(m, int32(_a_F_halfvec_sub_6), v17)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_halfvec_sub_7), int32(80), int32(_a_F_halfvec_sub_8))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v367 = v363
	goto L127
L127:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v367<<(uint(int32(1))%32)))))
	if v384&int32(_a_F_halfvec_sub_9) != int32(_a_F_halfvec_sub_4) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L133
	}
L129:
	;
	v390 = v367 + int32(1)
	if v364 != v390 {
		v367 = v390
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	goto L4
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_to_float4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+4)))
		v12 = F_mul_size(m, int32(4), v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_palloc(m, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+4)))
				if int32(0) < v16 {
					v22 = int32(0)
					for {
						v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(8)+v22<<(uint(int32(1))%32)))))
						v37 = v32 & int32(1023)
						v41 = v32 << (uint(int32(16)) % 32) & int32(-2147483648)
						v44 = int32(31)
						v45 = int32(base.Ui32(v32)>>(uint(int32(10))%32)) & v44
						if v45 != v44 {
							if v45 != 0 {
								v117 = v37
								v118 = v45<<(uint(int32(23))%32) + v41 + int32(939524096)
							} else {
								if v37 != 0 {
									if v32&int32(512) != 0 {
										v105 = v37 << (uint(int32(1)) % 32)
										v107 = int32(939524096)
									} else {
										if base.Ui32(int32(255)) < base.Ui32(v37) {
											v105 = v37 << (uint(int32(2)) % 32)
											v107 = int32(931135488)
										} else {
											if base.Ui32(int32(127)) < base.Ui32(v37) {
												v105 = v37 << (uint(int32(3)) % 32)
												v107 = int32(922746880)
											} else {
												if base.Ui32(int32(63)) < base.Ui32(v37) {
													v105 = v37 << (uint(int32(4)) % 32)
													v107 = int32(914358272)
												} else {
													if base.Ui32(int32(31)) < base.Ui32(v37) {
														v105 = v37 << (uint(int32(5)) % 32)
														v107 = int32(905969664)
													} else {
														if base.Ui32(int32(15)) < base.Ui32(v37) {
															v105 = v37 << (uint(int32(6)) % 32)
															v107 = int32(897581056)
														} else {
															if base.Ui32(int32(7)) < base.Ui32(v37) {
																v105 = v37 << (uint(int32(7)) % 32)
																v107 = int32(889192448)
															} else {
																if base.Ui32(int32(3)) < base.Ui32(v37) {
																	v105 = v37 << (uint(int32(8)) % 32)
																	v107 = int32(880803840)
																} else {
																	v100 = base.B2i32(v37 == int32(1))
																	if v37 == int32(1) {
																		v101 = int32(1024)
																	} else {
																		v101 = v37 << (uint(int32(9)) % 32)
																	}
																	if v37 == int32(1) {
																		v104 = int32(864026624)
																	} else {
																		v104 = int32(872415232)
																	}
																	v105 = v101
																	v107 = v104
																}
															}
														}
													}
												}
											}
										}
									}
									v117 = v105 & int32(1022)
									v118 = v107 | v41
								} else {
									v117 = int32(0)
									v118 = v41
								}
							}
						} else {
							if v37 == int32(0) {
								v117 = int32(0)
								v118 = v41 | int32(2139095040)
							} else {
								v117 = v37
								v118 = v41 | int32(2143289344)
							}
						}
						*(*float32)(unsafe.Add(mBase, uint32(v14+v22<<(uint(int32(2))%32)))) = base.F32_reinterpret_i32(v118 | v117<<(uint(int32(13))%32))
						v125 = v22 + int32(1)
						if v125 != v16 {
							v22 = v125
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v135 = F_construct_array(m, v14, v16, int32(700), int32(4), int32(1), int32(105))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v14)
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						return v135
					}
				}
			}
		}
	}
}
