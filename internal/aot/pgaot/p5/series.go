package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_int4(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_int4(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_generate_series_int8(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_int8(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_generate_series_numeric(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_numeric(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_generate_series_numeric_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 float64
	_ = v412
	var v413 int32
	_ = v413
	var v419 float64
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v432 int64
	_ = v432
	var v440 int64
	_ = v440
	var v448 int64
	_ = v448
	v8 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = base.I32_wrap_i64(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 != int32(468) {
		v448 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(96)
	return v448
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v19 == int32(0) {
		v448 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(15) {
		v448 = v8
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = F_estimate_expression_value(m, v25, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int64(0)
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = F_estimate_expression_value(m, v33, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if int32(3) <= v38 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = F_estimate_expression_value(m, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v46 = int32(0)
	goto L10
L10:
	;
	v48 = v14 & int64(4294967295)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v49 == int32(7) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v46 = v44
	goto L10
L12:
	;
	v448 = v440
	goto L1
L13:
	;
	v440 = v432
	goto L12
L14:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v82
	v85 = *(*int64)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v85
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v91 = F_pg_detoast_datum(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L32
	}
L15:
	;
	v76 = int32(7)
	if base.B2i32(v49 != v76)|base.B2i32(v53 != v76) != 0 {
		v448 = v8
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v68 = int32(7)
	if base.B2i32(v49 != v68)|base.B2i32(v53 != v68) != 0 {
		v448 = v8
		goto L1
	} else {
		goto L29
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(0)
	v432 = v48
	goto L13
L18:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+32)))
	if v52 != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v53 == int32(7) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+32)))
	if v56 != 0 {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v59 != int32(7) {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+32)))
	if v62 != int32(1) {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	goto L17
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v73 == int32(7) {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v448 = v8
	goto L1
L31:
	;
	goto L14
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v94 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
	if base.Ui32(int32(_a_F_generate_series_numeric_support_0)) < base.Ui32(v96) {
		v440 = v8
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	if base.Ui32(int32(_a_F_generate_series_numeric_support_0)) < base.Ui32(v99) {
		v440 = v8
		goto L12
	} else {
		goto L35
	}
L35:
	;
	if v46 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v103 = F_pg_detoast_datum(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v181 = v12 + int32(72)
	v182 = int32(_a_F_generate_series_numeric_support_7)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[3]))
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[4]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v191 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L39:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	if base.Ui32(int32(_a_F_generate_series_numeric_support_0)) < base.Ui32(v105) {
		v440 = v8
		goto L12
	} else {
		goto L40
	}
L40:
	;
	v109 = v12 + int32(72)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	if int32(0) <= v117 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L38
L42:
	;
	v120 = int32(-8)
	goto L44
L43:
	;
	v120 = int32(-6)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(base.Ui32(int32(base.Ui32(v112)>>(uint(int32(2))%32))+v120) >> (uint(int32(1)) % 32))
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v125 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v141
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	v144 = int32(_a_F_generate_series_numeric_support_2)
	v145 = v143 & v144
	if v145 != v144 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v129 = v125 & int32(_a_F_generate_series_numeric_support_1)
	v141 = v129<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v129&int32(63)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+6)))
	v141 = v139
	goto L45
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v156
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v158 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	if v145 != int32(_a_F_generate_series_numeric_support_3) {
		v156 = v145
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v156 = v143 & int32(_a_F_generate_series_numeric_support_5)
	goto L49
L53:
	;
	v156 = v143 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_numeric_support_4)
	goto L49
L54:
	;
	v167 = int32(base.Ui32(v158)>>(uint(int32(7))%32)) & int32(63)
	goto L56
L55:
	;
	v167 = v158 & int32(_a_F_generate_series_numeric_support_6)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v167
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v170
	if v169 < v170 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v176 = int32(6)
	goto L59
L58:
	;
	v176 = int32(8)
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v103 + v176
	goto L41
L60:
	;
	if v227 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L61:
	;
	if v190 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	if v190 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v227 = int32(0)
	goto L60
L65:
	;
	goto L66
L66:
	;
	if v189 == int32(_a_F_generate_series_numeric_support_4) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v201 = int32(1)
	goto L69
L68:
	;
	v201 = int32(-1)
	goto L69
L69:
	;
	v227 = v201
	goto L60
L70:
	;
	if v202 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[5]))
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[6]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	if v202 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v207 = int32(-1)
	goto L75
L74:
	;
	v207 = int32(1)
	goto L75
L75:
	;
	v227 = v207
	goto L60
L76:
	;
	if v189 == int32(_a_F_generate_series_numeric_support_4) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	if v189 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v227 = int32(1)
	goto L60
L80:
	;
	goto L81
L81:
	;
	v217 = F_cmp_abs_common(m, v211, v191, v210, v209, v190, v208)
	mBase = m.M
	v227 = v217
	goto L60
L82:
	;
	v227 = int32(-1)
	goto L60
L83:
	;
	goto L84
L84:
	;
	v221 = F_cmp_abs_common(m, v209, v190, v208, v211, v191, v210)
	mBase = m.M
	v227 = v221
	goto L60
L85:
	;
	v432 = int64(0)
	goto L13
L86:
	;
	goto L87
L87:
	;
	v232 = v12 + int32(48)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	if int32(0) <= v240 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v303 = v12 + int32(24)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v311 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if int32(0) <= v311 {
		goto L108
	} else {
		goto L109
	}
L89:
	;
	v243 = int32(-8)
	goto L91
L90:
	;
	v243 = int32(-6)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = int32(base.Ui32(int32(base.Ui32(v235)>>(uint(int32(2))%32))+v243) >> (uint(int32(1)) % 32))
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	if v248 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v264
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
	v267 = int32(_a_F_generate_series_numeric_support_2)
	v268 = v266 & v267
	if v268 != v267 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v252 = v248 & int32(_a_F_generate_series_numeric_support_1)
	v264 = v252<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v252&int32(63)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+6)))
	v264 = v262
	goto L92
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v279
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	if v281 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	if v268 != int32(_a_F_generate_series_numeric_support_3) {
		v279 = v268
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v279 = v266 & int32(_a_F_generate_series_numeric_support_5)
	goto L96
L100:
	;
	v279 = v266 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_numeric_support_4)
	goto L96
L101:
	;
	v290 = int32(base.Ui32(v281)>>(uint(int32(7))%32)) & int32(63)
	goto L103
L102:
	;
	v290 = v281 & int32(_a_F_generate_series_numeric_support_6)
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+12)) = v290
	v292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	v293 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = v293
	if v292 < v293 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v299 = int32(6)
	goto L106
L105:
	;
	v299 = int32(8)
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = v91 + v299
	goto L88
L107:
	;
	v373 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v373
	F_sub_var(m, v303, v232, v12)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L126
	}
L108:
	;
	v314 = int32(-8)
	goto L110
L109:
	;
	v314 = int32(-6)
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = int32(base.Ui32(int32(base.Ui32(v306)>>(uint(int32(2))%32))+v314) >> (uint(int32(1)) % 32))
	v319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if v319 < int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v335
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v338 = int32(_a_F_generate_series_numeric_support_2)
	v339 = v337 & v338
	if v339 != v338 {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	v323 = v319 & int32(_a_F_generate_series_numeric_support_1)
	v335 = v323<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v323&int32(63)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+6)))
	v335 = v333
	goto L111
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = v350
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if v352 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	if v339 != int32(_a_F_generate_series_numeric_support_3) {
		v350 = v339
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v350 = v337 & int32(_a_F_generate_series_numeric_support_5)
	goto L115
L119:
	;
	v350 = v337 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_numeric_support_4)
	goto L115
L120:
	;
	v361 = int32(base.Ui32(v352)>>(uint(int32(7))%32)) & int32(63)
	goto L122
L121:
	;
	v361 = v352 & int32(_a_F_generate_series_numeric_support_6)
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+12)) = v361
	v363 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	v364 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v303)+16)) = v364
	if v363 < v364 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v370 = int32(6)
	goto L125
L124:
	;
	v370 = int32(8)
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+20)) = v94 + v370
	goto L107
L126:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v381 == v382 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v46 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v419 = float64(0)
	goto L129
L129:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v421 == int32(0) {
		v432 = v48
		goto L13
	} else {
		goto L142
	}
L130:
	;
	v412 = F_numericvar_to_double_no_overflow(m, v12)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L141
	}
L131:
	;
	v386 = int32(0)
	F_div_var(m, v12, v12+int32(72), v12, v386, v386, v386)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v395 = v393 << (uint(int32(2)) % 32)
	if base.Ui32(int32(2147483644)) <= base.Ui32(v395) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L130
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(0)
	goto L130
L136:
	;
	goto L137
L137:
	;
	v405 = base.I32_div_s(v395+int32(7), int32(4))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v405 < v406 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v408 = v405
	goto L140
L139:
	;
	v408 = v406
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v408
	goto L130
L141:
	;
	v419 = base.F64_add(v412, float64(1))
	goto L129
L142:
	;
	F_pfree(m, v421)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	v432 = v48
	goto L13
}
