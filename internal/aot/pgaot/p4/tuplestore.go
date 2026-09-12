package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetTuplestoreDestReceiverParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	return
}
func F_build_tuplestore_recursively(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v208 int64
	_ = v208
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	v26 = m.G0
	v28 = v26 - int32(320)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l12)))
	if base.B2i32(int32(0) < l9)&base.B2i32(l9 < l7) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L7
	} else {
		goto L114
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L7
	} else {
		goto L107
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L100
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L95
	}
L5:
	;
	m.G0 = v28 + int32(320)
	return
L6:
	;
	F_initStringInfo(m, v28+int32(304))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v39 = F_quote_literal_cstr(m, l5)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l11 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if l10 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+220)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+216)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+212)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+204)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+200)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28)+196)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+192)) = l0
	F_appendStringInfo(m, v28+int32(304), int32(202453), v28+int32(192))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+172)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+168)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+164)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+156)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+152)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28)+148)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = l0
	F_appendStringInfo(m, v28+int32(304), int32(201924), v28+int32(144))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	v76 = int32(0)
	goto L10
L15:
	;
	v76 = int32(4)
	goto L10
L16:
	;
	v81 = v76 | int32(16)
	goto L18
L17:
	;
	v81 = v76 + int32(12)
	goto L18
L18:
	;
	v82 = F_palloc(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	if l7 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v86
	v96 = F_pg_sprintf(m, v28+int32(292), int32(498496), v28+int32(128))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v128 = l7
	goto L22
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v28)+304))
	v132 = F_SPI_execute(m, v129, int32(1), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L36
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v28 + int32(292)
	if l10 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = l5
	goto L26
L25:
	;
	goto L26
L26:
	;
	if l11 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v102 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v102
	v112 = F_pg_sprintf(m, v28+int32(280), int32(498496), v28+int32(112))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v122 = F_BuildTupleFromCStrings(m, l12, v82)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L34
	}
L30:
	;
	if l10 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = int32(16)
	goto L33
L32:
	;
	v116 = int32(12)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+v116))) = v28 + int32(280)
	goto L29
L34:
	;
	F_tuplestore_puttuple(m, l13, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v128 = int32(1)
	goto L22
L36:
	;
	if v132 != int32(5) {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if v137 == int64(0) {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v143 <= int32(1) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v146 = int32(4)
	v148 = v142 + v143<<(uint(v146)%32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+96))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v153 = v30 + v150<<(uint(v146)%32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+96))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+88))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+88))
	if v155 != v156 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if base.B2i32(v154 != v149)&base.B2i32(int32(0) <= v154) != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v148)+196))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153)+196))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v153)+188))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v148)+188))
	if v164 != v165 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	if base.B2i32(v163 != v162)&base.B2i32(int32(0) <= v163) != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_initStringInfo(m, v28+int32(264))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	F_initStringInfo(m, v28+int32(248))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	F_initStringInfo(m, v28+int32(232))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v208 = int64(0)
	goto L47
L47:
	;
	F_appendStringInfoString(m, v28+int32(264), l6)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L49
	}
L48:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v28)+264))
	if v334 != 0 {
		goto L85
	} else {
		goto L86
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = l4
	F_appendStringInfo(m, v28+int32(248), int32(179749), v28-int32(-64))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224+base.I32_wrap_i64(v208)<<(uint(int32(2))%32))))
	v231 = F_SPI_getvalue(m, v229, v142, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	v234 = F_SPI_getvalue(m, v229, v142, int32(2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v128
	v242 = F_pg_sprintf(m, v28+int32(292), int32(498496), v28+int32(48))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	if v231 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l4
	F_appendStringInfo(m, v28+int32(232), int32(179749), v28+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v28)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v28 + int32(292)
	if l10 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	v256 = F_strstr(m, v254, v255)
	mBase = m.M
	if v256 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l4
	F_appendStringInfo(m, v28+int32(264), int32(179892), v28+int32(16))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v266
	goto L62
L61:
	;
	goto L62
L62:
	;
	if l11 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v292 = F_BuildTupleFromCStrings(m, l12, v82)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L69
	}
L64:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v275 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v275
	v283 = F_pg_sprintf(m, v28+int32(280), int32(498496), v28)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	if l10 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v28 + int32(280)
	goto L63
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v28 + int32(280)
	goto L63
L69:
	;
	F_tuplestore_puttuple(m, l13, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_pfree(m, v292)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	if v231 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_build_tuplestore_recursively(m, l0, l1, l2, l3, l4, v231, v266, v128+int32(1), l8, l9, l10, l11, l12, l13)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v234 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	F_pfree(m, v231)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	F_pfree(m, v234)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v305 = v28 + int32(264)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v306))) = uint8(v307)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+12)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = v307
	goto L81
L80:
	;
	goto L79
L81:
	;
	v314 = v28 + int32(248)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v316 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v315))) = uint8(v316)
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v314)+4)) = v316
	goto L82
L82:
	;
	v323 = v28 + int32(232)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v324))) = uint8(v325)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+12)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v325
	goto L83
L83:
	;
	v332 = v208 + int64(1)
	if v332 != v137 {
		v208 = v332
		goto L47
	} else {
		goto L84
	}
L84:
	;
	goto L48
L85:
	;
	F_pfree(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v339 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+264)) = int32(0)
	goto L87
L89:
	;
	F_pfree(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	if v344 == int32(0) {
		goto L5
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = int32(0)
	goto L91
L93:
	;
	F_pfree(m, v344)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	goto L5
L95:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(16950), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_errdetail(m, int32(605999), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(512270), int32(1481), int32(177858))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(376826), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	v408 = F_format_type_with_typemod(m, v156, v149)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v410 = F_format_type_with_typemod(m, v155, v154)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v408
	F_errdetail(m, int32(618991), v28+int32(96))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(512270), int32(1498), int32(177858))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(376826), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v435 = F_format_type_with_typemod(m, v165, v162)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	v437 = F_format_type_with_typemod(m, v164, v163)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v435
	F_errdetail(m, int32(618923), v28+int32(80))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(512270), int32(1511), int32(177858))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(456715), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(512270), int32(1340), int32(20387))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v2 = l1
	v7 = F_palloc0(m, int32(120))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		v13 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+10)) = uint16(v13)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)) = uint8(v2)
		if l0 != 0 {
			v18 = int32(12)
		} else {
			v18 = int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v20
		v26 = base.I64_extend_i32_s(l2) << (uint(int64(10)) % 64)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v26
		v30 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v35 = F_GenerationContextCreate(m, v30, int32(168430), v20, int32(8192), int32(8388608))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
			v39 = *(*int32)(unsafe.Add(mBase, _consts[182]))
			v40 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+76)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v39
			v43 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+88)) = uint8(v43)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = int32(4096)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v40
			v50 = F_palloc(m, int32(16384))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v50
				v53 = F_GetMemoryChunkSpace(m, v50)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = int32(8)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = int64(4294967296)
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v59 - base.I64_extend_i32_u(v53)
					v64 = F_palloc(m, int32(192))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = v18
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
						v69 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)) = uint8(v69)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = int32(1867)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(1868)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = int32(1869)
						return v7
					}
				}
			}
		}
	}
}
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_tuplestore_gettuple(m, l0, l1, v8+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if l2 == int32(0) {
				v25 = v16
				v26 = v12
				v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return base.B2i32(v12 != int32(0))
				}
			} else {
				if v16&int32(1) != 0 {
					v25 = v16
					v26 = v12
					v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return base.B2i32(v12 != int32(0))
					}
				} else {
					v22 = F_heap_copy_minimal_tuple(m, v12, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = int32(1)
						v26 = v22
						v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return base.B2i32(v12 != int32(0))
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
			m.T0[v32].(func(*base.Module, int32))(m, l3)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return base.B2i32(v12 != int32(0))
			}
		}
	}
}
func F_tuplestore_rescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v7 = v3 + v4*int32(24)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		v53 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v53
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v53)
		return
	case 1:
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v11)
		return
	case 2:
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v15)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = F_BufFileSeek(m, v17, v15, int64(0), v15)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 == int32(0) {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errmsg(m, int32(395725), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							F_errfinish(m, int32(510749), int32(1308), int32(290282))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(360994), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errfinish(m, int32(510749), int32(1311), int32(290282))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_tuplestore_skiptuples(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v98 int32
	_ = v98
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 <= int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v98
L2:
	;
	v98 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v98 = int32(0)
	goto L1
L6:
	;
	v17 = l1
	goto L9
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v49 = v45 + v46*int32(24)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if l2 != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v25 = F_tuplestore_gettuple(m, l0, l2, v10+int32(15))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v98 = int32(1)
	goto L1
L11:
	;
	return int32(0)
L12:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_pfree(m, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v40 = int64(1)
	if base.Ui64(v40) < base.Ui64(v17) {
		v17 = v17 - v40
		goto L9
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L10
L23:
	;
	if v50&int32(1) != 0 {
		v98 = v4
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v50&int32(1) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if l1 <= base.I64_extend_i32_s(v53-v54) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v54 + base.I32_wrap_i64(l1)
	v98 = int32(1)
	goto L1
L28:
	;
	goto L29
L29:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v62)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v53
	v98 = v4
	goto L1
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v76 < base.I64_extend_i32_s(v77-v78) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v76 = l1
	v77 = v69
	goto L30
L32:
	;
	goto L33
L33:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v70
	v76 = l1 - int64(1)
	v77 = v70
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v77 - base.I32_wrap_i64(v76)
	v98 = int32(1)
	goto L1
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v78
	goto L5
}
