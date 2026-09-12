package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_rel_type_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+72))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_parseRelOptionsInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
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
	var v243 float64
	_ = v243
	var v244 float64
	_ = v244
	var v248 float64
	_ = v248
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v550 int32
	_ = v550
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	v16 = m.G0
	v18 = v16 - int32(192)
	m.G0 = v18
	v20 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_deconstruct_array_builtin(m, v20, int32(25), v18+int32(188), int32(0), v18+int32(184))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+184))
	if int32(0) < v30 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = l1 ^ int32(1)
	v45 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	F_pfree(m, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L154
	}
L7:
	;
	if int32(0) < l3 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	goto L6
L9:
	;
	v567 = v45 + int32(1)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v18)+184))
	if v567 < v568 {
		v45 = v567
		goto L7
	} else {
		goto L153
	}
L10:
	;
	F_pfree(m, v149)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L152
	}
L11:
	;
	if v523&int32(1) == int32(0) {
		goto L10
	} else {
		goto L150
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L146
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L142
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L138
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v53 = int32(2)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v45<<(uint(v53)%32))))
	v57 = int32(4)
	v58 = v56 + v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v63 = int32(base.Ui32(v59)>>(uint(v53)%32)) - v57
	v70 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L125
	}
L18:
	;
	v82 = l2 + v70<<(uint(int32(4))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v63 <= v84 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v409 = v70 + int32(1)
	if v409 != l3 {
		v70 = v409
		goto L18
	} else {
		goto L124
	}
L21:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v58))))
	if v87 != int32(61) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v134 != 0 {
		goto L20
	} else {
		goto L37
	}
L24:
	;
	v134 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v96 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v97 = v58
	v98 = v90
	v99 = v84
	v100 = v96
	goto L31
L28:
	;
	v122 = v90
	v126 = int32(0)
	goto L29
L29:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v134 = v126 - v127
	goto L23
L30:
	;
	v122 = v117
	v126 = v119
	goto L29
L31:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 != v102 {
		v117 = v98
		v119 = v100
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v117 = v111
	v119 = int32(0)
	goto L30
L33:
	;
	if v102 == int32(0) {
		v117 = v98
		v119 = v100
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v107 = v99 - int32(1)
	if v107 == int32(0) {
		v117 = v98
		v119 = v100
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v110 = int32(1)
	v111 = v98 + v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v112 != 0 {
		v97 = v97 + v110
		v98 = v111
		v99 = v107
		v100 = v112
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	if l1 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v135&int32(1) != 0 {
		goto L14
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v138 = v63 - v84
	v139 = F_palloc(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v144 = int32(1)
	v147 = v138 - v144
	if v147 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149+v147))) = uint8(v151)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+20))
	switch v154 {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		goto L48
	}
L44:
	;
	v148 = F__emscripten_memcpy_bulkmem(m, v139, v58+v142+v144, v147)
	mBase = m.M
	v149 = v148
	goto L46
L45:
	;
	v149 = v139
	goto L46
L46:
	;
	goto L43
L47:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v405)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v404
	goto L10
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L121
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v149
	if l1 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L50:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v153)+24))
	v288 = v282
	goto L86
L51:
	;
	v231 = v82 + int32(8)
	v232 = int32(0)
	v234 = F_parse_real(m, v149, v231, v232, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L74
	}
L52:
	;
	v185 = v82 + int32(8)
	v186 = int32(0)
	v188 = F_parse_int(m, v149, v185, v186, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L62
	}
L53:
	;
	v157 = F_strlen(m, v149)
	mBase = m.M
	v158 = F_parse_bool_with_len(m, v149, v157, v82+int32(8))
	mBase = m.M
	goto L54
L54:
	;
	if (v34|v158)&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v523 = l1 | v158
	goto L11
L56:
	;
	goto L57
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v171
	F_errmsg(m, int32(209645), v18+int32(48))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(504839), int32(1617), int32(251778))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	if (v34|v188)&int32(1) == int32(0) {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	if l1 == int32(0) {
		v523 = v188
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	if v198 <= v197 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v197 <= v200 {
		v523 = v188
		goto L11
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v149
	F_errmsg(m, int32(716676), v18+int32(80))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v153)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v218
	F_errdetail(m, int32(682995), v18-int32(-64))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(504839), int32(1637), int32(251778))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	if (v34|v234)&int32(1) == int32(0) {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	if l1 == int32(0) {
		v523 = v234
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v231)))
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v153)+32))
	if base.F64_lt(v243, v244) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v248 = *(*float64)(unsafe.Add(mBase, uint32(v153)+40))
	if base.F64_gt(v243, v248) == int32(0) {
		v523 = v234
		goto L11
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v149
	F_errmsg(m, int32(716676), v18+int32(128))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v153)+32))
	v269 = *(*float64)(unsafe.Add(mBase, uint32(v153)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+120)) = v269
	*(*float64)(unsafe.Add(mBase, uint32(v18)+112)) = v268
	F_errdetail(m, int32(682526), v18+int32(112))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(504839), int32(1657), int32(251778))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v298 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if l1 != 0 {
		goto L105
	} else {
		goto L106
	}
L88:
	;
	v301 = v149
	v302 = v298
	goto L92
L89:
	;
	goto L90
L90:
	;
	goto L87
L91:
	;
	if v339 == int32(0) {
		goto L47
	} else {
		goto L104
	}
L92:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v305 == v306 {
		v328 = v305
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v339 = int32(0)
	goto L91
L94:
	;
	v330 = int32(1)
	if v328 != 0 {
		v301 = v301 + v330
		v302 = v302 + v330
		goto L92
	} else {
		goto L103
	}
L95:
	;
	if base.Ui32((v305-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v316 = v305 | int32(32)
	goto L98
L97:
	;
	v316 = v305
	goto L98
L98:
	;
	if base.Ui32((v306-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v325 = v306 | int32(32)
	goto L101
L100:
	;
	v325 = v306
	goto L101
L101:
	;
	if v316 == v325 {
		v328 = v316
		goto L94
	} else {
		goto L102
	}
L102:
	;
	v339 = v316 - v325
	goto L91
L103:
	;
	goto L93
L104:
	;
	v288 = v288 + int32(8)
	goto L86
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v372
	F_pfree(m, v149)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L116
	}
L108:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v352
	F_errmsg(m, int32(209687), v18+int32(176))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v360 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v360
	F_errdetail_internal(m, int32(210227), v18+int32(160))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_errfinish(m, int32(504839), int32(1681), int32(251778))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	goto L9
L117:
	;
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v385)
	goto L9
L118:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v379 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	m.T0[v379].(func(*base.Module, int32))(m, v149)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v392
	F_errmsg_internal(m, int32(486188), v18+int32(32))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(504839), int32(1703), int32(251778))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	goto L19
L125:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428+v45<<(uint(int32(2))%32))))
	v433 = F_text_to_cstring(m, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v435 = int32(61)
	v436 = F___strchrnul(m, v433, v435)
	mBase = m.M
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if v438 == v435 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v442 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v442 = v436
	goto L130
L129:
	;
	v442 = int32(0)
	goto L130
L130:
	;
	goto L127
L131:
	;
	v443 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v443)
	goto L133
L132:
	;
	goto L133
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v433
	F_errmsg(m, int32(715031), v18)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(504839), int32(1488), int32(319140))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v469
	F_errmsg(m, int32(424358), v18+int32(16))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(504839), int32(1601), int32(251778))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v489
	F_errmsg(m, int32(209603), v18+int32(96))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(504839), int32(1629), int32(251778))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+148)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v510
	F_errmsg(m, int32(209554), v18+int32(144))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(504839), int32(1649), int32(251778))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	v530 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v530)
	F_pfree(m, v149)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	goto L9
L152:
	;
	goto L9
L153:
	;
	goto L8
L154:
	;
	if l0 != v20 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_pfree(m, v20)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	m.G0 = v18 + int32(192)
	return
L158:
	;
	goto L157
}
func F_set_rel_width(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v141 int32
	_ = v141
	var v142 float64
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v162 int32
	_ = v162
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v254 int64
	_ = v254
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v287 int64
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v306 int64
	_ = v306
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v333 int64
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int64
	_ = v340
	var v343 int64
	_ = v343
	var v359 int64
	_ = v359
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v366 int64
	_ = v366
	v3 = int32(0)
	v14 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = v39
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v45 == int32(0) {
		v359 = v14
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v35 = v21 + v22<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v35 = v28 + v29<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v363 = int64(1073741823)
	if v363 <= v359 {
		goto L65
	} else {
		goto L66
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v195 == int32(0) {
		v359 = v198
		goto L5
	} else {
		goto L40
	}
L8:
	;
	v195 = v3
	v198 = v14
	goto L7
L9:
	;
	goto L10
L10:
	;
	v52 = v19 + int32(16)
	v54 = v19 + int32(24)
	v58 = v3
	v65 = v3
	v68 = v14
	goto L11
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v58<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v76 != int32(319) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v195 = v175
	v198 = v178
	goto L7
L13:
	;
	v182 = v58 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v182 < v183 {
		v58 = v182
		v65 = v175
		v68 = v178
		goto L11
	} else {
		goto L39
	}
L14:
	;
	v146 = F_exprType(m, v75)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L29
	} else {
		goto L35
	}
L15:
	;
	if v76 != int32(6) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v122 = F_find_placeholder_info(m, l0, v75)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L33
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v81 != v82 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+8)))
	if v84 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v175 = int32(1)
	v178 = v68
	goto L13
L21:
	;
	goto L22
L22:
	;
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v91 = (v84 - v88) << (uint(int32(2)) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91+v92)))
	if int32(0) < v94 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v175 = v65
	v178 = v68 + base.I64_extend_i32_u(v94)
	goto L13
L24:
	;
	goto L25
L25:
	;
	if v37 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v115 = F_get_typavgwidth(m, v113, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L32
	}
L27:
	;
	if v84 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v103 = F_get_attavgwidth(m, v37, v84)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	if v103 <= int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v107+v91))) = v103
	v175 = v65
	v178 = v68 + base.I64_extend_i32_u(v103)
	goto L13
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v91))) = v115
	v175 = v65
	v178 = v68 + base.I64_extend_i32_s(v115)
	goto L13
L33:
	;
	v124 = int64(*(*int32)(unsafe.Add(mBase, uint32(v122)+24)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l0
	v127 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v127
	v133 = F_cost_qual_eval_walker(m, v125, v19+int32(8))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v136)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v136)+16)) = base.F64_add(v137, v138)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v141)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v141)+24)) = base.F64_add(v135, v142)
	v175 = v65
	v178 = v68 + v124
	goto L13
L35:
	;
	v148 = F_exprTypmod(m, v75)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v150 = F_get_typavgwidth(m, v146, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l0
	v153 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v153
	v159 = F_cost_qual_eval_walker(m, v75, v19+int32(8))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v162)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v162)+16)) = base.F64_add(v163, v164)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v167)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v167)+24)) = base.F64_add(v161, v168)
	v175 = v65
	v178 = v68 + base.I64_extend_i32_s(v150)
	goto L13
L39:
	;
	goto L12
L40:
	;
	if v37 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v336 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v340 = int64(1073741823)
	if v340 <= v333 {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	if v220 == int32(0) {
		v333 = v287
		goto L41
	} else {
		goto L58
	}
L43:
	;
	v240 = int32(0)
	v242 = v205
	v254 = int64(24)
	goto L55
L44:
	;
	v205 = int32(1)
	v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+82)))
	if v206 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v231 = F_get_relation_data_width(m, v37, v226-v227<<(uint(int32(2))%32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L29
	} else {
		goto L54
	}
L47:
	;
	v333 = int64(24)
	goto L41
L48:
	;
	goto L49
L49:
	;
	v210 = int32(2)
	v213 = base.I32_extend16_s(v206 + int32(1))
	if v213 <= v210 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v216 = v210
	goto L52
L51:
	;
	v216 = v213
	goto L52
L52:
	;
	v218 = v216 - int32(1)
	v220 = v218 & int32(3)
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if int32(5) <= v213 {
		goto L43
	} else {
		goto L53
	}
L53:
	;
	v275 = v205
	v287 = int64(24)
	goto L42
L54:
	;
	v333 = base.I64_extend_i32_s(v231) + int64(24)
	goto L41
L55:
	;
	v259 = v222 + (v242-v221)<<(uint(int32(2))%32)
	v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v259))))
	v262 = int64(*(*int32)(unsafe.Add(mBase, uint32(v259)+4)))
	v264 = int64(*(*int32)(unsafe.Add(mBase, uint32(v259)+8)))
	v266 = int64(*(*int32)(unsafe.Add(mBase, uint32(v259)+12)))
	v267 = v254 + v260 + v262 + v264 + v266
	v268 = int32(4)
	v269 = v242 + v268
	v271 = v240 + v268
	if v271 != v218&int32(-4) {
		v240 = v271
		v242 = v269
		v254 = v267
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v275 = v269
	v287 = v267
	goto L42
L57:
	;
	goto L56
L58:
	;
	v294 = v275
	v295 = int32(0)
	v306 = v287
	goto L59
L59:
	;
	v312 = int64(*(*int32)(unsafe.Add(mBase, uint32(v222+(v294-v221)<<(uint(int32(2))%32)))))
	v313 = v306 + v312
	v314 = int32(1)
	v317 = v295 + v314
	if v317 != v220 {
		v294 = v294 + v314
		v295 = v317
		v306 = v313
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v333 = v313
	goto L41
L61:
	;
	goto L60
L62:
	;
	v343 = v340
	goto L64
L63:
	;
	v343 = v333
	goto L64
L64:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v335-v336<<(uint(int32(2))%32)))) = uint32(v343)
	v359 = v198 + v333
	goto L5
L65:
	;
	v366 = v363
	goto L67
L66:
	;
	v366 = v359
	goto L67
L67:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v362)+32)) = uint32(v366)
	m.G0 = v19 + int32(32)
	return
}
func F_transformRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	if l1 == v7 {
		v580 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(48)
	return v580
L2:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v247 <= int32(0) {
		v566 = v240
		goto L55
	} else {
		goto L56
	}
L4:
	;
	v240 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v26 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_deconstruct_array_builtin(m, v26, int32(25), v20+int32(44), int32(0), v20+int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v38 <= int32(0) {
		v240 = v7
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v51 = v7
	v53 = v7
	v55 = v38
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v53<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v63 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v240 = v220
	goto L3
L13:
	;
	v228 = v53 + int32(1)
	if v228 < v224 {
		v51 = v220
		v53 = v228
		v55 = v224
		goto L11
	} else {
		goto L54
	}
L14:
	;
	v66 = int32(4)
	v67 = v62 + v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v75 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v207 = F_accumArrayResult(m, v51, v62, int32(0), int32(25), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L53
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v73+v75<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if l2 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L16
L19:
	;
	v184 = v75 + int32(1)
	if v184 != v63 {
		v75 = v184
		goto L17
	} else {
		goto L52
	}
L20:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v129 = F_strlen(m, v128)
	mBase = m.M
	if int32(base.Ui32(v68)>>(uint(int32(2))%32))-v66 <= v129 {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	if v96 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v96 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L24:
	;
	goto L19
L25:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v106 == int32(0) {
		v125 = v105
		v126 = v106
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v126-v125 != 0 {
		goto L19
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	if v105 != v106 {
		v125 = v105
		v126 = v106
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v110 = v96
	v111 = l2
	goto L30
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v115 == int32(0) {
		v125 = v114
		v126 = v115
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v125 = v114
	v126 = v115
	goto L27
L32:
	;
	v118 = int32(1)
	if v114 == v115 {
		v110 = v110 + v118
		v111 = v111 + v118
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L20
L35:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v67))))
	if v132 != int32(61) {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	if v129 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v178 == int32(0) {
		v220 = v51
		v224 = v55
		goto L13
	} else {
		goto L51
	}
L38:
	;
	v178 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v140 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = v67
	v142 = v128
	v143 = v129
	v144 = v140
	goto L45
L42:
	;
	v166 = v128
	v170 = int32(0)
	goto L43
L43:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v178 = v170 - v171
	goto L37
L44:
	;
	v166 = v161
	v170 = v163
	goto L43
L45:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v144 != v146 {
		v161 = v142
		v163 = v144
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v161 = v155
	v163 = int32(0)
	goto L44
L47:
	;
	if v146 == int32(0) {
		v161 = v142
		v163 = v144
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v151 = v143 - int32(1)
	if v151 == int32(0) {
		v161 = v142
		v163 = v144
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v154 = int32(1)
	v155 = v142 + v154
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	if v156 != 0 {
		v141 = v141 + v154
		v142 = v155
		v143 = v151
		v144 = v156
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	goto L19
L52:
	;
	goto L18
L53:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v220 = v207
	v224 = v209
	goto L13
L54:
	;
	goto L12
L55:
	;
	if v566 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L56:
	;
	v261 = v240
	v262 = int32(0)
	goto L58
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L7
	} else {
		goto L136
	}
L58:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268+v262<<(uint(int32(2))%32))))
	if l5 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L7
	} else {
		goto L132
	}
L60:
	;
	goto L59
L61:
	;
	v499 = v262 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v499 < v500 {
		v261 = v491
		v262 = v499
		goto L58
	} else {
		goto L131
	}
L62:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	if v273 == int32(0) {
		v491 = v261
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v292 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(135444), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(504839), int32(1242), int32(141479))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	if v394 != 0 {
		goto L101
	} else {
		goto L102
	}
L71:
	;
	if l3 == int32(0) {
		goto L57
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if l2 != 0 {
		v491 = v261
		goto L61
	} else {
		goto L100
	}
L74:
	;
	v295 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v296 == v295 {
		goto L57
	} else {
		goto L75
	}
L75:
	;
	v299 = v296
	v307 = v295
	goto L76
L76:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v319 == int32(0) {
		v338 = v318
		v339 = v319
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if l2 == int32(0) {
		v491 = v261
		goto L61
	} else {
		goto L90
	}
L78:
	;
	if v339-v338 != 0 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	goto L78
L80:
	;
	if v318 != v319 {
		v338 = v318
		v339 = v319
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v323 = v292
	v324 = v299
	goto L82
L82:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+1)))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+1)))
	if v328 == int32(0) {
		v338 = v327
		v339 = v328
		goto L79
	} else {
		goto L84
	}
L83:
	;
	v338 = v327
	v339 = v328
	goto L79
L84:
	;
	v331 = int32(1)
	if v327 == v328 {
		v323 = v323 + v331
		v324 = v324 + v331
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v342 = v307 + int32(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l3+v342<<(uint(int32(2))%32))))
	if v346 != 0 {
		v299 = v346
		v307 = v342
		goto L76
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	goto L77
L89:
	;
	goto L57
L90:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v352 == int32(0) {
		v371 = v351
		v372 = v352
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v372-v371 == int32(0) {
		goto L70
	} else {
		goto L99
	}
L92:
	;
	goto L91
L93:
	;
	if v351 != v352 {
		v371 = v351
		v372 = v352
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v356 = v292
	v357 = l2
	goto L95
L95:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+1)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+1)))
	if v361 == int32(0) {
		v371 = v360
		v372 = v361
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v371 = v360
	v372 = v361
	goto L92
L97:
	;
	v364 = int32(1)
	if v360 == v361 {
		v356 = v356 + v364
		v357 = v357 + v364
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v491 = v261
	goto L61
L100:
	;
	goto L70
L101:
	;
	v395 = F_defGetString(m, v272)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L104
	}
L102:
	;
	v398 = int32(352272)
	goto L103
L103:
	;
	v399 = int32(61)
	v400 = F___strchrnul(m, v393, v399)
	mBase = m.M
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v402 == v399 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v398 = v395
	goto L103
L105:
	;
	if v406 != 0 {
		goto L60
	} else {
		goto L109
	}
L106:
	;
	v406 = v400
	goto L108
L107:
	;
	v406 = int32(0)
	goto L108
L108:
	;
	goto L105
L109:
	;
	if l4 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v456 = F_strlen(m, v393)
	mBase = m.M
	v457 = F_strlen(m, v398)
	mBase = m.M
	v458 = v456 + v457
	v461 = F_palloc(m, v458+int32(6))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L7
	} else {
		goto L128
	}
L111:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v409 != 0 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v410 = int32(177395)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, _consts[13])))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	if v414 == int32(0) {
		v433 = v413
		v434 = v414
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if v434-v433 != 0 {
		goto L110
	} else {
		goto L121
	}
L114:
	;
	goto L113
L115:
	;
	if v413 != v414 {
		v433 = v413
		v434 = v414
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v418 = v393
	v419 = v410
	goto L117
L117:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v423 == int32(0) {
		v433 = v422
		v434 = v423
		goto L114
	} else {
		goto L119
	}
L118:
	;
	v433 = v422
	v434 = v423
	goto L114
L119:
	;
	v426 = int32(1)
	if v422 == v423 {
		v418 = v418 + v426
		v419 = v419 + v426
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v436 = F_defGetBoolean(m, v272)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	if v436 == int32(0) {
		v491 = v261
		goto L61
	} else {
		goto L123
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(453664), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(504839), int32(1320), int32(141479))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = v458<<(uint(int32(2))%32) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v393
	v473 = F_pg_sprintf(m, v461+int32(4), int32(180447), v20)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v479 = F_accumArrayResult(m, v261, v461, int32(0), int32(25), v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	v491 = v479
	goto L61
L131:
	;
	v566 = v491
	goto L55
L132:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v393
	F_errmsg(m, int32(744738), v20+int32(16))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(504839), int32(1306), int32(141479))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L7
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v544
	F_errmsg(m, int32(735948), v20+int32(32))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(504839), int32(1276), int32(141479))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v580 = int32(0)
	goto L1
L141:
	;
	goto L142
L142:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v578 = F_makeArrayResult(m, v566, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	v580 = v578
	goto L1
}
