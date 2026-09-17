package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_int4(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_int4(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_series_int8(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_int8(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_series_numeric(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_numeric(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_series_numeric_support(m *base.Module, l0 int32) int32 {
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
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
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 != int32(460) {
		v434 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(96)
	return v434
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v17 == int32(0) {
		v434 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v20 != int32(15) {
		v434 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = F_estimate_expression_value(m, v23, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = F_estimate_expression_value(m, v31, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(3) <= v37 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v43 = F_estimate_expression_value(m, v40, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v45 = int32(0)
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v46 == int32(7) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v45 = v43
	goto L10
L12:
	;
	v434 = v426
	goto L1
L13:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v82
	v85 = *(*int64)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v85
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v91 = F_pg_detoast_datum(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L31
	}
L14:
	;
	v74 = int32(0)
	v75 = int32(7)
	if base.B2i32(v46 != v75)|base.B2i32(v50 != v75) != 0 {
		v434 = v74
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v65 = int32(0)
	v66 = int32(7)
	if base.B2i32(v46 != v66)|base.B2i32(v50 != v66) != 0 {
		v434 = v65
		goto L1
	} else {
		goto L28
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
	v426 = v13
	goto L12
L17:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	if v49 != 0 {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v50 == int32(7) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
	if v53 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v45 == int32(0) {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v56 != int32(7) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+24)))
	if v59 != int32(1) {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	goto L16
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v71 == int32(7) {
		v80 = v65
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v434 = v65
	goto L1
L30:
	;
	v80 = v74
	goto L13
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v94 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
	if base.Ui32(int32(_a_F_generate_series_numeric_support_0)) < base.Ui32(v96) {
		v434 = v80
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	if base.Ui32(int32(_a_F_generate_series_numeric_support_0)) < base.Ui32(v99) {
		v434 = v80
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v45 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v103 = F_pg_detoast_datum(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v181 = v11 + int32(72)
	v182 = int32(_a_F_generate_series_numeric_support_7)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[3]))
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[4]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v191 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L38:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	if base.Ui32(int32(_a_F_generate_series_numeric_support_0)) < base.Ui32(v105) {
		v434 = v80
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v109 = v11 + int32(72)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	if int32(0) <= v117 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L37
L41:
	;
	v120 = int32(-8)
	goto L43
L42:
	;
	v120 = int32(-6)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(base.Ui32(int32(base.Ui32(v112)>>(uint(int32(2))%32))+v120) >> (uint(int32(1)) % 32))
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v125 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v141
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	v144 = int32(_a_F_generate_series_numeric_support_2)
	v145 = v143 & v144
	if v145 != v144 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v129 = v125 & int32(_a_F_generate_series_numeric_support_1)
	v141 = v129<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v129&int32(63)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+6)))
	v141 = v139
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v156
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v158 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	if v145 != int32(_a_F_generate_series_numeric_support_3) {
		v156 = v145
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v156 = v143 & int32(_a_F_generate_series_numeric_support_5)
	goto L48
L52:
	;
	v156 = v143 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_numeric_support_4)
	goto L48
L53:
	;
	v167 = int32(base.Ui32(v158)>>(uint(int32(7))%32)) & int32(63)
	goto L55
L54:
	;
	v167 = v158 & int32(_a_F_generate_series_numeric_support_6)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v167
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+4)))
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v170
	if v169 < v170 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v176 = int32(6)
	goto L58
L57:
	;
	v176 = int32(8)
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v103 + v176
	goto L40
L59:
	;
	if v227 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L60:
	;
	if v190 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	if v190 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	v227 = int32(0)
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v189 == int32(_a_F_generate_series_numeric_support_4) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v201 = int32(1)
	goto L68
L67:
	;
	v201 = int32(-1)
	goto L68
L68:
	;
	v227 = v201
	goto L59
L69:
	;
	if v202 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[5]))
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_numeric_support[6]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	if v202 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v207 = int32(-1)
	goto L74
L73:
	;
	v207 = int32(1)
	goto L74
L74:
	;
	v227 = v207
	goto L59
L75:
	;
	if v189 == int32(_a_F_generate_series_numeric_support_4) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	if v189 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v227 = int32(1)
	goto L59
L79:
	;
	goto L80
L80:
	;
	v217 = F_cmp_abs_common(m, v211, v191, v210, v209, v190, v208)
	mBase = m.M
	v227 = v217
	goto L59
L81:
	;
	v227 = int32(-1)
	goto L59
L82:
	;
	goto L83
L83:
	;
	v221 = F_cmp_abs_common(m, v209, v190, v208, v211, v191, v210)
	mBase = m.M
	v227 = v221
	goto L59
L84:
	;
	v426 = int32(0)
	goto L12
L85:
	;
	goto L86
L86:
	;
	v232 = v11 + int32(48)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	if int32(0) <= v240 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v303 = v11 + int32(24)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v311 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if int32(0) <= v311 {
		goto L107
	} else {
		goto L108
	}
L88:
	;
	v243 = int32(-8)
	goto L90
L89:
	;
	v243 = int32(-6)
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = int32(base.Ui32(int32(base.Ui32(v235)>>(uint(int32(2))%32))+v243) >> (uint(int32(1)) % 32))
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	if v248 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v264
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
	v267 = int32(_a_F_generate_series_numeric_support_2)
	v268 = v266 & v267
	if v268 != v267 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v252 = v248 & int32(_a_F_generate_series_numeric_support_1)
	v264 = v252<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v252&int32(63)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+6)))
	v264 = v262
	goto L91
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v279
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	if v281 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	if v268 != int32(_a_F_generate_series_numeric_support_3) {
		v279 = v268
		goto L95
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v279 = v266 & int32(_a_F_generate_series_numeric_support_5)
	goto L95
L99:
	;
	v279 = v266 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_numeric_support_4)
	goto L95
L100:
	;
	v290 = int32(base.Ui32(v281)>>(uint(int32(7))%32)) & int32(63)
	goto L102
L101:
	;
	v290 = v281 & int32(_a_F_generate_series_numeric_support_6)
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+12)) = v290
	v292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+4)))
	v293 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = v293
	if v292 < v293 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v299 = int32(6)
	goto L105
L104:
	;
	v299 = int32(8)
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = v91 + v299
	goto L87
L106:
	;
	v373 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v373
	F_sub_var(m, v303, v232, v11)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L125
	}
L107:
	;
	v314 = int32(-8)
	goto L109
L108:
	;
	v314 = int32(-6)
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = int32(base.Ui32(int32(base.Ui32(v306)>>(uint(int32(2))%32))+v314) >> (uint(int32(1)) % 32))
	v319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if v319 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v335
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v338 = int32(_a_F_generate_series_numeric_support_2)
	v339 = v337 & v338
	if v339 != v338 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v323 = v319 & int32(_a_F_generate_series_numeric_support_1)
	v335 = v323<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v323&int32(63)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+6)))
	v335 = v333
	goto L110
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = v350
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if v352 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	if v339 != int32(_a_F_generate_series_numeric_support_3) {
		v350 = v339
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v350 = v337 & int32(_a_F_generate_series_numeric_support_5)
	goto L114
L118:
	;
	v350 = v337 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_numeric_support_4)
	goto L114
L119:
	;
	v361 = int32(base.Ui32(v352)>>(uint(int32(7))%32)) & int32(63)
	goto L121
L120:
	;
	v361 = v352 & int32(_a_F_generate_series_numeric_support_6)
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+12)) = v361
	v363 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	v364 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v303)+16)) = v364
	if v363 < v364 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v370 = int32(6)
	goto L124
L123:
	;
	v370 = int32(8)
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+20)) = v94 + v370
	goto L106
L125:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v381 == v382 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if v45 != 0 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v419 = float64(0)
	goto L128
L128:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v421 == int32(0) {
		v426 = v13
		goto L12
	} else {
		goto L141
	}
L129:
	;
	v412 = F_numericvar_to_double_no_overflow(m, v11)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L140
	}
L130:
	;
	v386 = int32(0)
	F_div_var(m, v11, v11+int32(72), v11, v386, v386, v386)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v395 = v393 << (uint(int32(2)) % 32)
	if base.Ui32(int32(2147483644)) <= base.Ui32(v395) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L129
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(0)
	goto L129
L135:
	;
	goto L136
L136:
	;
	v405 = base.I32_div_s(v395+int32(7), int32(4))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v405 < v406 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v408 = v405
	goto L139
L138:
	;
	v408 = v406
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v408
	goto L129
L140:
	;
	v419 = base.F64_add(v412, float64(1))
	goto L128
L141:
	;
	F_pfree(m, v421)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	v426 = v13
	goto L12
}
