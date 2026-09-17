package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn13940(m, l0, l1, int32(_a_F_uuid_abbrev_abort_0), int32(381), int32(_a_F_uuid_abbrev_abort_1), int32(_a_F_uuid_abbrev_abort_2), int32(374), int32(_a_F_uuid_abbrev_abort_3), int32(356), int32(_a_F_uuid_abbrev_abort_4))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_uuid_generate_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	switch l0 {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L8
	case 3, 5:
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L43
	} else {
		goto L151
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L43
	} else {
		goto L138
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L43
	} else {
		goto L125
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L43
	} else {
		goto L112
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L43
	} else {
		goto L99
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L43
	} else {
		goto L86
	}
L7:
	;
	v304 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v9+int32(112))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L43
	} else {
		goto L85
	}
L8:
	;
	v292 = v9 + int32(96)
	F_uuid_generate_random(m, v292)
	mBase = m.M
	F_uuid_unparse(m, v292, v9+int32(112))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L43
	} else {
		goto L84
	}
L9:
	;
	if l0 == int32(3) {
		goto L68
	} else {
		goto L69
	}
L10:
	;
	v133 = v9 + int32(96)
	F_uuid_generate_time(m, v133)
	mBase = m.M
	F_uuid_unparse(m, v133, v9+int32(112))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L43
	} else {
		goto L44
	}
L11:
	;
	v12 = v9 + int32(112)
	goto L15
L12:
	;
	goto L7
L13:
	;
	v129 = F_strlen(m, v118)
	mBase = m.M
	goto L12
L15:
	;
	goto L16
L16:
	;
	v19 = int32(36)
	if (v12^l2)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v122)
	goto L13
L18:
	;
	v103 = v98
	v104 = v99
	v105 = v100
	goto L39
L19:
	;
	if v93 == int32(0) {
		v118 = v91
		v119 = v92
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v91 = l2
	v92 = v12
	v93 = v19
	goto L19
L21:
	;
	goto L22
L22:
	;
	v23 = int32(0)
	if base.B2i32(l2&int32(3) == v23)|int32(0) == v23 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v59 == int32(0) {
		v118 = v56
		v119 = v57
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v35 = l2
	v36 = v12
	v37 = v19
	goto L27
L25:
	;
	goto L26
L26:
	;
	v56 = l2
	v57 = v12
	v58 = v19
	v59 = int32(1)
	goto L23
L27:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v39)
	if v39 == int32(0) {
		v98 = v35
		v99 = v36
		v100 = v37
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v56 = v50
	v57 = v44
	v58 = v46
	v59 = v48
	goto L23
L29:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v37 - v43
	v47 = int32(0)
	v48 = base.B2i32(v46 != v47)
	v50 = v35 + v43
	if v50&int32(3) == v47 {
		v56 = v50
		v57 = v44
		v58 = v46
		v59 = v48
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v46 != 0 {
		v35 = v50
		v36 = v44
		v37 = v46
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.B2i32(v62 == int32(0))|base.B2i32(base.Ui32(v58) < base.Ui32(int32(4))) != 0 {
		v91 = v56
		v92 = v57
		v93 = v58
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v69 = v56
	v70 = v57
	v71 = v58
	goto L34
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v77 = int32(-2139062144)
	if (int32(16843008)-v74|v74)&v77 != v77 {
		v98 = v69
		v99 = v70
		v100 = v71
		goto L18
	} else {
		goto L36
	}
L35:
	;
	v91 = v85
	v92 = v83
	v93 = v87
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
	v82 = int32(4)
	v83 = v70 + v82
	v85 = v69 + v82
	v87 = v71 - v82
	if base.Ui32(int32(3)) < base.Ui32(v87) {
		v69 = v85
		v70 = v83
		v71 = v87
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v98 = v91
	v99 = v92
	v100 = v93
	goto L18
L39:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v107)
	if v107 == int32(0) {
		v118 = v103
		v119 = v104
		goto L17
	} else {
		goto L41
	}
L40:
	;
	v118 = v114
	v119 = v112
	goto L17
L41:
	;
	v111 = int32(1)
	v112 = v104 + v111
	v114 = v103 + v111
	v116 = v105 - v111
	if v116 != 0 {
		v103 = v114
		v104 = v112
		v105 = v116
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	return int32(0)
L44:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(int32(36) < l3) != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v148 = v9 - l3 + int32(148)
	if (l2^v148)&int32(3) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L7
L47:
	;
	goto L46
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v202)
	if v202&int32(255) == int32(0) {
		goto L47
	} else {
		goto L63
	}
L49:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v201 = l2
	v202 = v154
	v203 = v148
	goto L48
L50:
	;
	goto L51
L51:
	;
	if l2&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v158 = l2
	v160 = v148
	goto L55
L53:
	;
	v172 = l2
	v174 = v148
	goto L54
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v179 = int32(-2139062144)
	if (int32(16843008)-v176|v176)&v179 != v179 {
		v201 = v172
		v202 = v176
		v203 = v174
		goto L48
	} else {
		goto L59
	}
L55:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v161)
	if v161 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v172 = v168
	v174 = v166
	goto L54
L57:
	;
	v165 = int32(1)
	v166 = v160 + v165
	v168 = v158 + v165
	if v168&int32(3) != 0 {
		v158 = v168
		v160 = v166
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v184 = v172
	v185 = v176
	v186 = v174
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v185
	v188 = int32(4)
	v189 = v186 + v188
	v191 = v184 + v188
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v196 = int32(-2139062144)
	if (int32(16843008)-v193|v193)&v196 == v196 {
		v184 = v191
		v185 = v193
		v186 = v189
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v201 = v191
	v202 = v193
	v203 = v189
	goto L48
L62:
	;
	goto L61
L63:
	;
	v210 = v201
	v212 = v203
	goto L64
L64:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)) = uint8(v213)
	v215 = int32(1)
	if v213 != 0 {
		v210 = v210 + v215
		v212 = v212 + v215
		goto L64
	} else {
		goto L66
	}
L65:
	;
	goto L47
L66:
	;
	goto L65
L67:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+104)))
	v276 = v272&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+104)) = uint8(v276)
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+102)))
	v283 = v278&int32(_a_F_uuid_generate_internal_0) | l0<<(uint(int32(4))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+102)) = uint16(v283)
	F_uuid_unparse(m, v9+int32(96), v9+int32(112))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L43
	} else {
		goto L83
	}
L68:
	;
	v226 = F_pg_cryptohash_create(m, int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L43
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v247 = F_pg_cryptohash_create(m, int32(1))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L43
	} else {
		goto L77
	}
L71:
	;
	v228 = F_pg_cryptohash_init(m, v226)
	mBase = m.M
	if v228 < int32(0) {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v232 = F_pg_cryptohash_update(m, v226, l1, int32(16))
	mBase = m.M
	if v232 < int32(0) {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	v235 = F_pg_cryptohash_update(m, v226, l2, l3)
	mBase = m.M
	if v235 < int32(0) {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	v241 = F_pg_cryptohash_final(m, v226, v9+int32(96), int32(16))
	mBase = m.M
	if v241 < int32(0) {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_pg_cryptohash_free(m, v226)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L43
	} else {
		goto L76
	}
L76:
	;
	goto L67
L77:
	;
	v249 = F_pg_cryptohash_init(m, v247)
	mBase = m.M
	if v249 < int32(0) {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v253 = F_pg_cryptohash_update(m, v247, l1, int32(16))
	mBase = m.M
	if v253 < int32(0) {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v256 = F_pg_cryptohash_update(m, v247, l2, l3)
	mBase = m.M
	if v256 < int32(0) {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v262 = F_pg_cryptohash_final(m, v247, v9+int32(112), int32(20))
	mBase = m.M
	if v262 < int32(0) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_pg_cryptohash_free(m, v247)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L43
	} else {
		goto L82
	}
L82:
	;
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v9)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v9)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v269
	goto L67
L83:
	;
	goto L7
L84:
	;
	goto L7
L85:
	;
	m.G0 = v9 + int32(160)
	return v304
L86:
	;
	if v226 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_uuid_generate_internal_1)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_2), v9)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L43
	} else {
		goto L97
	}
L88:
	;
	v328 = int32(_a_F_uuid_generate_internal_3)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v320 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v323 = int32(_a_F_uuid_generate_internal_4)
	goto L93
L92:
	;
	v323 = int32(_a_F_uuid_generate_internal_5)
	goto L93
L93:
	;
	if v320 == int32(2) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v326 = int32(_a_F_uuid_generate_internal_3)
	goto L96
L95:
	;
	v326 = v323
	goto L96
L96:
	;
	v328 = v326
	goto L87
L97:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(338), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L43
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	if v226 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_uuid_generate_internal_1)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_8), v9+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L43
	} else {
		goto L110
	}
L101:
	;
	v358 = int32(_a_F_uuid_generate_internal_3)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v350 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v353 = int32(_a_F_uuid_generate_internal_4)
	goto L106
L105:
	;
	v353 = int32(_a_F_uuid_generate_internal_5)
	goto L106
L106:
	;
	if v350 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v356 = int32(_a_F_uuid_generate_internal_3)
	goto L109
L108:
	;
	v356 = v353
	goto L109
L109:
	;
	v358 = v356
	goto L100
L110:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(342), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L43
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	if v226 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(_a_F_uuid_generate_internal_1)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_9), v9+int32(32))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L43
	} else {
		goto L123
	}
L114:
	;
	v390 = int32(_a_F_uuid_generate_internal_3)
	goto L113
L115:
	;
	goto L116
L116:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v382 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v385 = int32(_a_F_uuid_generate_internal_4)
	goto L119
L118:
	;
	v385 = int32(_a_F_uuid_generate_internal_5)
	goto L119
L119:
	;
	if v382 == int32(2) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v388 = int32(_a_F_uuid_generate_internal_3)
	goto L122
L121:
	;
	v388 = v385
	goto L122
L122:
	;
	v390 = v388
	goto L113
L123:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(347), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L43
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	if v247 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(_a_F_uuid_generate_internal_10)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_2), v9+int32(48))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L43
	} else {
		goto L136
	}
L127:
	;
	v422 = int32(_a_F_uuid_generate_internal_3)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v414 == int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v417 = int32(_a_F_uuid_generate_internal_4)
	goto L132
L131:
	;
	v417 = int32(_a_F_uuid_generate_internal_5)
	goto L132
L132:
	;
	if v414 == int32(2) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v420 = int32(_a_F_uuid_generate_internal_3)
	goto L135
L134:
	;
	v420 = v417
	goto L135
L135:
	;
	v422 = v420
	goto L126
L136:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(357), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L43
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
	if v247 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = int32(_a_F_uuid_generate_internal_10)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_8), v9-int32(-64))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L43
	} else {
		goto L149
	}
L140:
	;
	v454 = int32(_a_F_uuid_generate_internal_3)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v446 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v449 = int32(_a_F_uuid_generate_internal_4)
	goto L145
L144:
	;
	v449 = int32(_a_F_uuid_generate_internal_5)
	goto L145
L145:
	;
	if v446 == int32(2) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v452 = int32(_a_F_uuid_generate_internal_3)
	goto L148
L147:
	;
	v452 = v449
	goto L148
L148:
	;
	v454 = v452
	goto L139
L149:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(361), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L43
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	if v247 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_uuid_generate_internal_10)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_9), v9+int32(80))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L43
	} else {
		goto L162
	}
L153:
	;
	v486 = int32(_a_F_uuid_generate_internal_3)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v478 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v481 = int32(_a_F_uuid_generate_internal_4)
	goto L158
L157:
	;
	v481 = int32(_a_F_uuid_generate_internal_5)
	goto L158
L158:
	;
	if v478 == int32(2) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v484 = int32(_a_F_uuid_generate_internal_3)
	goto L161
L160:
	;
	v484 = v481
	goto L161
L161:
	;
	v486 = v484
	goto L152
L162:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(364), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L43
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_uuid_generate_time(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_uuid_generate_time[0])))
	if v11 == int32(0) {
		v14 = int32(_a_F_uuid_generate_time_0)
		v16 = m.Env.Pgmem_random_bytes(m, v14, int32(6))
		mBase = m.M
		v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_uuid_generate_time[1])))
		v20 = int32(1)
		v21 = v19 | v20
		*(*uint8)(unsafe.Add(mBase, _c_F_uuid_generate_time[1])) = uint8(v21)
		v23 = int32(_a_F_uuid_generate_time_1)
		v25 = m.Env.Pgmem_random_bytes(m, v23, int32(2))
		mBase = m.M
		*(*uint8)(unsafe.Add(mBase, _c_F_uuid_generate_time[0])) = uint8(v20)
		v31 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_uuid_generate_time[2])))
		v33 = v31 & int32(_a_F_uuid_generate_time_2)
		*(*uint16)(unsafe.Add(mBase, _c_F_uuid_generate_time[2])) = uint16(v33)
	} else {
	}
	F_gettimeofday(m, v8)
	mBase = m.M
	v36 = int32(_a_F_uuid_generate_time_3)
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	v45 = v37*int64(10000000) + v40*int64(10) + int64(122192928000000000)
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_uuid_generate_time[3]))
	if base.Ui64(v47) < base.Ui64(v45) {
		v51 = v45
	} else {
		v51 = v47 + int64(1)
	}
	*(*int64)(unsafe.Add(mBase, _c_F_uuid_generate_time[3])) = v51
	v54 = int64(base.Ui64(v51) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v54)
	v57 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v57)
	v60 = int64(base.Ui64(v51) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v51)
	v64 = int64(base.Ui64(v51) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v64)
	v67 = int64(base.Ui64(v51) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v67)
	v70 = int64(base.Ui64(v51) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v70)
	v77 = int32(16)
	v78 = base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(56))%64)))&int32(15) | v77
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v78)
	v81 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_uuid_generate_time[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v81)
	v88 = int32(base.Ui32(v81)>>(uint(int32(8))%32))&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v88)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_uuid_generate_time[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = v91
	v94 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_uuid_generate_time[4])))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v94)
	m.G0 = v8 + v77
	return
}
func F_uuid_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = int32(0)
	v26 = v13 + base.B2i32(v19 == int32(123))
	goto L5
L3:
	;
	m.G0 = v11 + int32(16)
	return v15
L4:
	;
	v114 = F_errsave_start(m, v23)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L29
	}
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	if v19 == int32(123) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v39)
	v42 = v39 & int32(255)
	goto L9
L9:
	;
	if base.B2i32(base.Ui32(v42-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v42|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(base.Ui32(v39) >> (uint(int32(8)) % 32))
	goto L11
L11:
	;
	if base.B2i32(base.Ui32(v57-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v57|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v71)
	v79 = F_strtox_2(m, v11+int32(12), v71, int32(16), int64(4294967295))
	mBase = m.M
	v80 = base.I32_wrap_i64(v79)
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25+v15))) = uint8(v80)
	v83 = v26 + int32(2)
	if v25&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v88 = v26 + int32(3)
	goto L16
L15:
	;
	v88 = v83
	goto L16
L16:
	;
	if v25 != int32(15) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v91 = v88
	goto L19
L18:
	;
	v91 = v83
	goto L19
L19:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+2)))
	if v92 != int32(45) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v95 = v83
	goto L22
L21:
	;
	v95 = v91
	goto L22
L22:
	;
	v97 = v25 + int32(1)
	if v97 != int32(16) {
		v25 = v97
		v26 = v95
		goto L5
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v102 != int32(125) {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	v107 = v95
	goto L26
L26:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L27:
	;
	v107 = v95 + int32(1)
	goto L26
L28:
	;
	goto L4
L29:
	;
	if v114 == int32(0) {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_uuid_in_0)
	F_errmsg(m, int32(_a_F_uuid_in_1), v11)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errsave_finish(m, v23, int32(_a_F_uuid_in_2), int32(183), int32(_a_F_uuid_in_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L3
}
func F_uuid_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v14 != int32(255) {
			v130 = v14
			v131 = v6 + int32(15)
			v133 = v130 + int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
			v135 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
			return v6
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v19)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
			if v21 != int32(255) {
				v130 = v21
				v131 = v6 + int32(14)
				v133 = v130 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
				v135 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
				return v6
			} else {
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v26)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)))
				if v28 != int32(255) {
					v130 = v28
					v131 = v6 + int32(13)
					v133 = v130 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
					v135 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
					return v6
				} else {
					v33 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v33)
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
					if v35 != int32(255) {
						v130 = v35
						v131 = v6 + int32(12)
						v133 = v130 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
						v135 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
						return v6
					} else {
						v40 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)))
						if v42 != int32(255) {
							v130 = v42
							v131 = v6 + int32(11)
							v133 = v130 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
							v135 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
							return v6
						} else {
							v47 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v47)
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)))
							if v49 != int32(255) {
								v130 = v49
								v131 = v6 + int32(10)
								v133 = v130 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
								v135 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
								return v6
							} else {
								v54 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v54)
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
								if v56 != int32(255) {
									v130 = v56
									v131 = v6 + int32(9)
									v133 = v130 + int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
									v135 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
									return v6
								} else {
									v61 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v61)
									v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)))
									if v63 != int32(255) {
										v130 = v63
										v131 = v6 + int32(8)
										v133 = v130 + int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
										v135 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
										return v6
									} else {
										v68 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v68)
										v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
										if v70 != int32(255) {
											v130 = v70
											v131 = v6 + int32(7)
											v133 = v130 + int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
											v135 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
											return v6
										} else {
											v75 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v75)
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)))
											if v77 != int32(255) {
												v130 = v77
												v131 = v6 + int32(6)
												v133 = v130 + int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
												v135 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
												return v6
											} else {
												v82 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v82)
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
												if v84 != int32(255) {
													v130 = v84
													v131 = v6 + int32(5)
													v133 = v130 + int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
													v135 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
													return v6
												} else {
													v89 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v89)
													v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
													if v91 != int32(255) {
														v130 = v91
														v131 = v6 + int32(4)
														v133 = v130 + int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
														v135 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
														return v6
													} else {
														v96 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v96)
														v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
														if v98 != int32(255) {
															v130 = v98
															v131 = v6 + int32(3)
															v133 = v130 + int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
															v135 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
															return v6
														} else {
															v103 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v103)
															v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
															if v105 != int32(255) {
																v130 = v105
																v131 = v6 + int32(2)
																v133 = v130 + int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																v135 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																return v6
															} else {
																v110 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v110)
																v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
																if v112 != int32(255) {
																	v130 = v112
																	v131 = v6 + int32(1)
																	v133 = v130 + int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																	v135 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																	return v6
																} else {
																	v117 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v117)
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
																	if v119 != int32(255) {
																		v130 = v119
																		v131 = v6
																		v133 = v130 + int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																		v135 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																		return v6
																	} else {
																		v122 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v122)
																		F_pfree(m, v6)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			v126 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v126)
																			return int32(0)
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_uuid_ns_x500(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14016(m, l0, int32(_a_F_uuid_ns_x500_0), int32(_a_F_uuid_ns_x500_1), int32(_a_F_uuid_ns_x500_2), int32(_a_F_uuid_ns_x500_3), int32(_a_F_uuid_ns_x500_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_uuid_sortsupport(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13942(m, l0, int32(1529), int32(1528), int32(1527))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
