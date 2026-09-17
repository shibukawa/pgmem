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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 float64
	_ = v242
	var v243 float64
	_ = v243
	var v247 float64
	_ = v247
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
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
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L152
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L148
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L144
	}
L7:
	;
	v34 = l1 ^ int32(1)
	v47 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	F_pfree(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L139
	}
L10:
	;
	v50 = int32(0)
	if l3 <= v50 {
		v464 = v50
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v469 = int32(0)
	if base.B2i32(l1 == v469)|base.B2i32(v464 < l3) == v469 {
		goto L123
	} else {
		goto L124
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v54 = int32(2)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v47<<(uint(v54)%32))))
	v58 = int32(4)
	v59 = v57 + v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v64 = int32(base.Ui32(v60)>>(uint(v54)%32)) - v58
	v75 = v50
	goto L14
L14:
	;
	v82 = l2 + v75<<(uint(int32(4))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v64 <= v84 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v464 = l3
	goto L12
L16:
	;
	v452 = v75 + int32(1)
	if v452 != l3 {
		v75 = v452
		goto L14
	} else {
		goto L122
	}
L17:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v59))))
	if v87 != int32(61) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v135 != 0 {
		goto L16
	} else {
		goto L32
	}
L20:
	;
	v135 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v96 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v97 = v59
	v98 = v90
	v99 = v84
	v100 = v96
	goto L27
L24:
	;
	v123 = v90
	v127 = int32(0)
	goto L25
L25:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v135 = v127 - v128
	goto L19
L26:
	;
	v123 = v118
	v127 = v120
	goto L25
L27:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if base.B2i32(v100 != v102)|base.B2i32(v102 == int32(0)) != 0 {
		v118 = v98
		v120 = v100
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v118 = v112
	v120 = int32(0)
	goto L26
L29:
	;
	v108 = v99 - int32(1)
	if v108 == int32(0) {
		v118 = v98
		v120 = v100
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v111 = int32(1)
	v112 = v98 + v111
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v113 != 0 {
		v97 = v97 + v111
		v98 = v112
		v99 = v108
		v100 = v113
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	if l1 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v136&int32(1) != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v139 = v64 - v84
	v140 = F_palloc(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v143 = v139 - int32(1)
	if v143 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	base.MemoryCopy(m, v140, v59+v145+int32(1), v143)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140+v143))) = uint8(v151)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+20))
	switch v154 {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	default:
		goto L44
	}
L41:
	;
	F_pfree(m, v140)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L121
	}
L42:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v431)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v430
	goto L41
L43:
	;
	if v419&int32(1) == int32(0) {
		goto L41
	} else {
		goto L119
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L116
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v140
	if l1 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L46:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v153)+24))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v282 != 0 {
		goto L80
	} else {
		goto L81
	}
L47:
	;
	v230 = v82 + int32(8)
	v231 = int32(0)
	v233 = F_parse_real(m, v140, v230, v231, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L68
	}
L48:
	;
	v184 = v82 + int32(8)
	v185 = int32(0)
	v187 = F_parse_int(m, v140, v184, v185, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L56
	}
L49:
	;
	v157 = F_strlen(m, v140)
	mBase = m.M
	v158 = F_parse_bool_with_len(m, v140, v157, v82+int32(8))
	mBase = m.M
	goto L50
L50:
	;
	if (v34|v158)&int32(1) != 0 {
		v419 = v158
		goto L43
	} else {
		goto L51
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v170
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_0), v18+int32(48))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1617), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	if (v34|v187)&int32(1) == int32(0) {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	if l1 == int32(0) {
		v419 = v187
		goto L43
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	if v197 <= v196 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v196 <= v199 {
		v419 = v187
		goto L43
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v140
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_3), v18+int32(80))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v153)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v217
	F_errdetail(m, int32(_a_F_parseRelOptionsInternal_4), v18-int32(-64))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1637), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	if (v34|v233)&int32(1) == int32(0) {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	if l1 == int32(0) {
		v419 = v233
		goto L43
	} else {
		goto L70
	}
L70:
	;
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v230)))
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v153)+32))
	if base.F64_lt(v242, v243) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v247 = *(*float64)(unsafe.Add(mBase, uint32(v153)+40))
	if base.F64_gt(v242, v247) == int32(0) {
		v419 = v233
		goto L43
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v140
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_3), v18+int32(128))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v153)+32))
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v153)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+120)) = v268
	*(*float64)(unsafe.Add(mBase, uint32(v18)+112)) = v267
	F_errdetail(m, int32(_a_F_parseRelOptionsInternal_5), v18+int32(112))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1657), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v291 = v281
	v292 = v282
	goto L83
L81:
	;
	goto L82
L82:
	;
	if l1 != 0 {
		goto L100
	} else {
		goto L101
	}
L83:
	;
	v300 = v140
	v301 = v292
	goto L86
L84:
	;
	goto L82
L85:
	;
	if v338 == int32(0) {
		goto L42
	} else {
		goto L98
	}
L86:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v304 == v305 {
		v327 = v304
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v338 = int32(0)
	goto L85
L88:
	;
	v329 = int32(1)
	if v327 != 0 {
		v300 = v300 + v329
		v301 = v301 + v329
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v304-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v315 = v304 | int32(32)
	goto L92
L91:
	;
	v315 = v304
	goto L92
L92:
	;
	if base.Ui32((v305-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v324 = v305 | int32(32)
	goto L95
L94:
	;
	v324 = v305
	goto L95
L95:
	;
	if v315 == v324 {
		v327 = v315
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v338 = v315 - v324
	goto L85
L97:
	;
	goto L87
L98:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	if v341 != 0 {
		v291 = v291 + int32(8)
		v292 = v341
		goto L83
	} else {
		goto L99
	}
L99:
	;
	goto L84
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v387
	F_pfree(m, v140)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v367
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_6), v18+int32(176))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v375 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v375
	F_errdetail_internal(m, int32(_a_F_parseRelOptionsInternal_7), v18+int32(160))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1681), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v464 = v75
	goto L12
L112:
	;
	v400 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v400)
	v464 = v75
	goto L12
L113:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v394 == int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	m.T0[v394].(func(*base.Module, int32))(m, v140)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	goto L112
L116:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v407
	F_errmsg_internal(m, int32(_a_F_parseRelOptionsInternal_8), v18+int32(16))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1703), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v426 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v426)
	F_pfree(m, v140)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v464 = v75
	goto L12
L121:
	;
	v464 = v75
	goto L12
L122:
	;
	goto L15
L123:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v475+v47<<(uint(int32(2))%32))))
	v480 = F_text_to_cstring(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v511 = v47 + int32(1)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v18)+184))
	if v511 < v512 {
		v47 = v511
		goto L10
	} else {
		goto L138
	}
L126:
	;
	v482 = int32(61)
	v483 = F___strchrnul(m, v480, v482)
	mBase = m.M
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	if v485 == v482 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v489 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v489 = v483
	goto L130
L129:
	;
	v489 = int32(0)
	goto L130
L130:
	;
	goto L127
L131:
	;
	v490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v489))) = uint8(v490)
	goto L133
L132:
	;
	goto L133
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v480
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_9), v18+int32(32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1488), int32(_a_F_parseRelOptionsInternal_10))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
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
	goto L11
L139:
	;
	if l0 != v20 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_pfree(m, v20)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	m.G0 = v18 + int32(192)
	return
L143:
	;
	goto L142
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v546
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_11), v18)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1601), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v564
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_12), v18+int32(96))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1629), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+148)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v585
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_13), v18+int32(144))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1649), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_rel_width(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v97 float64
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v238 int64
	_ = v238
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v272 int64
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v288 int64
	_ = v288
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v324 int64
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v349 int64
	_ = v349
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v356 int64
	_ = v356
	v3 = int32(0)
	v13 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v38
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v44 == int32(0) {
		v349 = v13
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v34 = v20 + v21<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v34 = v27 + v28<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v353 = int64(1073741823)
	if v353 <= v349 {
		goto L63
	} else {
		goto L64
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if int32(0) < v47 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = v18 + int32(16)
	v57 = v3
	v62 = v3
	v64 = v13
	goto L10
L8:
	;
	v192 = v3
	v194 = v13
	goto L9
L9:
	;
	if v192 == int32(0) {
		v349 = v194
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v57<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 != int32(6) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v192 = v173
	v194 = v175
	goto L9
L12:
	;
	v179 = v57 + int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v179 < v180 {
		v57 = v179
		v62 = v173
		v64 = v175
		goto L10
	} else {
		goto L37
	}
L13:
	;
	v143 = F_exprType(m, v71)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L33
	}
L14:
	;
	if v72 != int32(319) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v101 != v102 {
		goto L13
	} else {
		goto L21
	}
L17:
	;
	v77 = F_find_placeholder_info(m, l0, v71)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v79 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+24)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l0
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v82
	v88 = F_cost_qual_eval_walker(m, v80, v18+int32(8))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v91)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v91)+16)) = base.F64_add(v92, v93)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v96)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v96)+24)) = base.F64_add(v90, v97)
	v173 = v62
	v175 = v64 + v79
	goto L12
L21:
	;
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+8)))
	if v104 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v173 = int32(1)
	v175 = v64
	goto L12
L23:
	;
	goto L24
L24:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v111 = (v104 - v108) << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112)))
	if int32(0) < v114 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v173 = v62
	v175 = v64 + base.I64_extend_i32_u(v114)
	goto L12
L26:
	;
	goto L27
L27:
	;
	v119 = int32(0)
	if base.B2i32(v36 == v119)|base.B2i32(v104 <= v119) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	v136 = F_get_typavgwidth(m, v134, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L32
	}
L29:
	;
	v124 = F_get_attavgwidth(m, v36, v104)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v124 <= int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v128+v111))) = v124
	v173 = v62
	v175 = v64 + base.I64_extend_i32_u(v124)
	goto L12
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v111))) = v136
	v173 = v62
	v175 = v64 + base.I64_extend_i32_s(v136)
	goto L12
L33:
	;
	v145 = F_exprTypmod(m, v71)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v147 = F_get_typavgwidth(m, v143, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l0
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v150
	v156 = F_cost_qual_eval_walker(m, v71, v18+int32(8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v159)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v159)+16)) = base.F64_add(v160, v161)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v164)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v164)+24)) = base.F64_add(v158, v165)
	v173 = v62
	v175 = v64 + base.I64_extend_i32_s(v147)
	goto L12
L37:
	;
	goto L11
L38:
	;
	if v36 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v331 = int64(1073741823)
	if v331 <= v324 {
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v201 = int32(1)
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+82)))
	if v202 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v302 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v306 = F_get_relation_data_width(m, v36, v301-v302<<(uint(int32(2))%32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L18
	} else {
		goto L59
	}
L43:
	;
	v324 = int64(24)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v206 = int32(2)
	v209 = base.I32_extend16_s(v202 + int32(1))
	if v209 <= v206 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v212 = v206
	goto L48
L47:
	;
	v212 = v209
	goto L48
L48:
	;
	v214 = v212 - int32(1)
	v216 = v214 & int32(3)
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v219 = int64(24)
	if int32(5) <= v209 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v225 = v201
	v230 = int32(0)
	v238 = v219
	goto L52
L50:
	;
	v259 = v201
	v272 = v219
	goto L51
L51:
	;
	v275 = v259
	v280 = int32(0)
	v288 = v272
	goto L56
L52:
	;
	v243 = v218 + (v225-v217)<<(uint(int32(2))%32)
	v244 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243))))
	v246 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243)+4)))
	v248 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243)+8)))
	v250 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243)+12)))
	v251 = v238 + v244 + v246 + v248 + v250
	v252 = int32(4)
	v253 = v225 + v252
	v255 = v230 + v252
	if v255 != v214&int32(-4) {
		v225 = v253
		v230 = v255
		v238 = v251
		goto L52
	} else {
		goto L54
	}
L53:
	;
	if v216 == int32(0) {
		v324 = v251
		goto L39
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v259 = v253
	v272 = v251
	goto L51
L56:
	;
	v294 = int64(*(*int32)(unsafe.Add(mBase, uint32(v218+(v275-v217)<<(uint(int32(2))%32)))))
	v295 = v288 + v294
	v296 = int32(1)
	v299 = v280 + v296
	if v299 != v216 {
		v275 = v275 + v296
		v280 = v299
		v288 = v295
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v324 = v295
	goto L39
L58:
	;
	goto L57
L59:
	;
	v324 = base.I64_extend_i32_s(v306) + int64(24)
	goto L39
L60:
	;
	v334 = v331
	goto L62
L61:
	;
	v334 = v324
	goto L62
L62:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v326-v327<<(uint(int32(2))%32)))) = uint32(v334)
	v349 = v194 + v324
	goto L5
L63:
	;
	v356 = v353
	goto L65
L64:
	;
	v356 = v349
	goto L65
L65:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v352)+32)) = uint32(v356)
	m.G0 = v18 + int32(32)
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	if l1 == v7 {
		v585 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(48)
	return v585
L2:
	;
	if l0 == int32(0) {
		v243 = v7
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v249 <= int32(0) {
		v572 = v243
		goto L51
	} else {
		goto L52
	}
L4:
	;
	v26 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	F_deconstruct_array_builtin(m, v26, int32(25), v20+int32(44), int32(0), v20+int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v38 <= int32(0) {
		v243 = v7
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v50 = v38
	v51 = v7
	v52 = v7
	goto L9
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v51<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v63 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v243 = v223
	goto L3
L11:
	;
	v230 = v51 + int32(1)
	if v230 < v221 {
		v50 = v221
		v51 = v230
		v52 = v223
		goto L9
	} else {
		goto L50
	}
L12:
	;
	v66 = int32(4)
	v67 = v62 + v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v75 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_transformRelOptions[0]))
	v209 = F_accumArrayResult(m, v52, v62, int32(0), int32(25), v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L49
	}
L15:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v73+v75<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L14
L17:
	;
	v186 = v75 + int32(1)
	if v186 != v63 {
		v75 = v186
		goto L15
	} else {
		goto L48
	}
L18:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v130 = F_strlen(m, v129)
	mBase = m.M
	if int32(base.Ui32(v68)>>(uint(int32(2))%32))-v66 <= v130 {
		goto L17
	} else {
		goto L32
	}
L19:
	;
	if v96 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v96 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L17
L23:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v105 == int32(0))|base.B2i32(v105 != v108) != 0 {
		v126 = v105
		v127 = v108
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v126-v127 != 0 {
		goto L17
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	v111 = v96
	v112 = l2
	goto L27
L27:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v116 == int32(0) {
		v126 = v116
		v127 = v115
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v126 = v116
	v127 = v115
	goto L25
L29:
	;
	v119 = int32(1)
	if v116 == v115 {
		v111 = v111 + v119
		v112 = v112 + v119
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L18
L32:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v67))))
	if v133 != int32(61) {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	if v130 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v180 == int32(0) {
		v221 = v50
		v223 = v52
		goto L11
	} else {
		goto L47
	}
L35:
	;
	v180 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v141 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v142 = v67
	v143 = v129
	v144 = v130
	v145 = v141
	goto L42
L39:
	;
	v168 = v129
	v172 = int32(0)
	goto L40
L40:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v180 = v172 - v173
	goto L34
L41:
	;
	v168 = v163
	v172 = v165
	goto L40
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if base.B2i32(v145 != v147)|base.B2i32(v147 == int32(0)) != 0 {
		v163 = v143
		v165 = v145
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v163 = v157
	v165 = int32(0)
	goto L41
L44:
	;
	v153 = v144 - int32(1)
	if v153 == int32(0) {
		v163 = v143
		v165 = v145
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v156 = int32(1)
	v157 = v143 + v156
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v158 != 0 {
		v142 = v142 + v156
		v143 = v157
		v144 = v153
		v145 = v158
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L17
L48:
	;
	goto L16
L49:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v221 = v211
	v223 = v209
	goto L11
L50:
	;
	goto L10
L51:
	;
	if v572 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L52:
	;
	v261 = int32(0)
	v264 = v243
	goto L54
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L5
	} else {
		goto L129
	}
L54:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v261<<(uint(int32(2))%32))))
	if l5 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L5
	} else {
		goto L125
	}
L56:
	;
	goto L55
L57:
	;
	v504 = v261 + int32(1)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v504 < v505 {
		v261 = v504
		v264 = v497
		goto L54
	} else {
		goto L124
	}
L58:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	if v275 == int32(0) {
		v497 = v264
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v294 != 0 {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_transformRelOptions_0), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1242), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	if v398 != 0 {
		goto L95
	} else {
		goto L96
	}
L67:
	;
	if l3 == int32(0) {
		goto L53
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if l2 != 0 {
		v497 = v264
		goto L57
	} else {
		goto L94
	}
L70:
	;
	v297 = int32(0)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v298 == v297 {
		goto L53
	} else {
		goto L71
	}
L71:
	;
	v301 = v298
	v311 = v297
	goto L72
L72:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if base.B2i32(v320 == int32(0))|base.B2i32(v320 != v323) != 0 {
		v341 = v320
		v342 = v323
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if l2 == int32(0) {
		v497 = v264
		goto L57
	} else {
		goto L85
	}
L74:
	;
	if v341-v342 != 0 {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	goto L74
L76:
	;
	v326 = v294
	v327 = v301
	goto L77
L77:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+1)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+1)))
	if v331 == int32(0) {
		v341 = v331
		v342 = v330
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v341 = v331
	v342 = v330
	goto L75
L79:
	;
	v334 = int32(1)
	if v331 == v330 {
		v326 = v326 + v334
		v327 = v327 + v334
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v345 = v311 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l3+v345<<(uint(int32(2))%32))))
	if v349 != 0 {
		v301 = v349
		v311 = v345
		goto L72
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L73
L84:
	;
	goto L53
L85:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v354 == int32(0))|base.B2i32(v354 != v357) != 0 {
		v375 = v354
		v376 = v357
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v375-v376 == int32(0) {
		goto L66
	} else {
		goto L93
	}
L87:
	;
	goto L86
L88:
	;
	v360 = v294
	v361 = l2
	goto L89
L89:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+1)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	if v365 == int32(0) {
		v375 = v365
		v376 = v364
		goto L87
	} else {
		goto L91
	}
L90:
	;
	v375 = v365
	v376 = v364
	goto L87
L91:
	;
	v368 = int32(1)
	if v365 == v364 {
		v360 = v360 + v368
		v361 = v361 + v368
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v497 = v264
	goto L57
L94:
	;
	goto L66
L95:
	;
	v399 = F_defGetString(m, v274)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	v402 = int32(_a_F_transformRelOptions_3)
	goto L97
L97:
	;
	v403 = int32(61)
	v404 = F___strchrnul(m, v397, v403)
	mBase = m.M
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v406 == v403 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v402 = v399
	goto L97
L99:
	;
	if v410 != 0 {
		goto L56
	} else {
		goto L103
	}
L100:
	;
	v410 = v404
	goto L102
L101:
	;
	v410 = int32(0)
	goto L102
L102:
	;
	goto L99
L103:
	;
	if l4 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v461 = F_strlen(m, v397)
	mBase = m.M
	v462 = F_strlen(m, v402)
	mBase = m.M
	v463 = v461 + v462
	v466 = F_palloc(m, v463+int32(6))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L5
	} else {
		goto L121
	}
L105:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v413 != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v414 = int32(_a_F_transformRelOptions_4)
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformRelOptions[1])))
	if base.B2i32(v417 == int32(0))|base.B2i32(v417 != v420) != 0 {
		v438 = v417
		v439 = v420
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v438-v439 != 0 {
		goto L104
	} else {
		goto L114
	}
L108:
	;
	goto L107
L109:
	;
	v423 = v397
	v424 = v414
	goto L110
L110:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+1)))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	if v428 == int32(0) {
		v438 = v428
		v439 = v427
		goto L108
	} else {
		goto L112
	}
L111:
	;
	v438 = v428
	v439 = v427
	goto L108
L112:
	;
	v431 = int32(1)
	if v428 == v427 {
		v423 = v423 + v431
		v424 = v424 + v431
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v441 = F_defGetBoolean(m, v274)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	if v441 == int32(0) {
		v497 = v264
		goto L57
	} else {
		goto L116
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_transformRelOptions_5), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1320), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v466))) = v463<<(uint(int32(2))%32) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v397
	v478 = F_pg_sprintf(m, v466+int32(4), int32(_a_F_transformRelOptions_6), v20)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_transformRelOptions[0]))
	v484 = F_accumArrayResult(m, v264, v466, int32(0), int32(25), v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v497 = v484
	goto L57
L124:
	;
	v572 = v497
	goto L51
L125:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v397
	F_errmsg(m, int32(_a_F_transformRelOptions_7), v20+int32(16))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1306), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v549
	F_errmsg(m, int32(_a_F_transformRelOptions_8), v20+int32(32))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1276), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v585 = int32(0)
	goto L1
L134:
	;
	goto L135
L135:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_transformRelOptions[0]))
	v583 = F_makeArrayResult(m, v572, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v585 = v583
	goto L1
}
