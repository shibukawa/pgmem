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
	v11 = Fn14332(m, l0, l1, int32(_a_F_uuid_abbrev_abort_0), int32(398), int32(_a_F_uuid_abbrev_abort_1), int32(_a_F_uuid_abbrev_abort_2), int32(391), int32(_a_F_uuid_abbrev_abort_3), int32(373), int32(_a_F_uuid_abbrev_abort_4))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_uuid_generate_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
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
	v473 = m.ExcPending
	if v473 != 0 {
		goto L43
	} else {
		goto L151
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L43
	} else {
		goto L138
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L43
	} else {
		goto L125
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L43
	} else {
		goto L112
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L43
	} else {
		goto L99
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L43
	} else {
		goto L86
	}
L7:
	;
	v306 = F_DirectFunctionCall1Coll(m, int32(3591), int32(0), base.I64_extend_i32_u(v10+int32(112)))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L43
	} else {
		goto L85
	}
L8:
	;
	v293 = v10 + int32(96)
	F_uuid_generate_random(m, v293)
	mBase = m.M
	F_uuid_unparse(m, v293, v10+int32(112))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	v134 = v10 + int32(96)
	F_uuid_generate_time(m, v134)
	mBase = m.M
	F_uuid_unparse(m, v134, v10+int32(112))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L43
	} else {
		goto L44
	}
L11:
	;
	v13 = v10 + int32(112)
	goto L15
L12:
	;
	goto L7
L13:
	;
	v130 = F_strlen(m, v119)
	mBase = m.M
	goto L12
L15:
	;
	goto L16
L16:
	;
	v20 = int32(36)
	if (v13^l2)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v123)
	goto L13
L18:
	;
	v104 = v99
	v105 = v100
	v106 = v101
	goto L39
L19:
	;
	if v94 == int32(0) {
		v119 = v92
		v120 = v93
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v92 = l2
	v93 = v13
	v94 = v20
	goto L19
L21:
	;
	goto L22
L22:
	;
	v24 = int32(0)
	if base.B2i32(l2&int32(3) == v24)|int32(0) == v24 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v60 == int32(0) {
		v119 = v57
		v120 = v58
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v36 = l2
	v37 = v13
	v38 = v20
	goto L27
L25:
	;
	goto L26
L26:
	;
	v57 = l2
	v58 = v13
	v59 = v20
	v60 = int32(1)
	goto L23
L27:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v40)
	if v40 == int32(0) {
		v99 = v36
		v100 = v37
		v101 = v38
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v57 = v51
	v58 = v45
	v59 = v47
	v60 = v49
	goto L23
L29:
	;
	v44 = int32(1)
	v45 = v37 + v44
	v47 = v38 - v44
	v48 = int32(0)
	v49 = base.B2i32(v47 != v48)
	v51 = v36 + v44
	if v51&int32(3) == v48 {
		v57 = v51
		v58 = v45
		v59 = v47
		v60 = v49
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v47 != 0 {
		v36 = v51
		v37 = v45
		v38 = v47
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.B2i32(v63 == int32(0))|base.B2i32(base.Ui32(v59) < base.Ui32(int32(4))) != 0 {
		v92 = v57
		v93 = v58
		v94 = v59
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v70 = v57
	v71 = v58
	v72 = v59
	goto L34
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 != v78 {
		v99 = v70
		v100 = v71
		v101 = v72
		goto L18
	} else {
		goto L36
	}
L35:
	;
	v92 = v86
	v93 = v84
	v94 = v88
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v75
	v83 = int32(4)
	v84 = v71 + v83
	v86 = v70 + v83
	v88 = v72 - v83
	if base.Ui32(int32(3)) < base.Ui32(v88) {
		v70 = v86
		v71 = v84
		v72 = v88
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v99 = v92
	v100 = v93
	v101 = v94
	goto L18
L39:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v108)
	if v108 == int32(0) {
		v119 = v104
		v120 = v105
		goto L17
	} else {
		goto L41
	}
L40:
	;
	v119 = v115
	v120 = v113
	goto L17
L41:
	;
	v112 = int32(1)
	v113 = v105 + v112
	v115 = v104 + v112
	v117 = v106 - v112
	if v117 != 0 {
		v104 = v115
		v105 = v113
		v106 = v117
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	return int64(0)
L44:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(int32(36) < l3) != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v149 = v10 - l3 + int32(148)
	if (l2^v149)&int32(3) != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v203)
	if v203&int32(255) == int32(0) {
		goto L47
	} else {
		goto L63
	}
L49:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v202 = l2
	v203 = v155
	v204 = v149
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
	v159 = l2
	v161 = v149
	goto L55
L53:
	;
	v173 = l2
	v175 = v149
	goto L54
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v180 = int32(-2139062144)
	if (int32(16843008)-v177|v177)&v180 != v180 {
		v202 = v173
		v203 = v177
		v204 = v175
		goto L48
	} else {
		goto L59
	}
L55:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v162)
	if v162 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v173 = v169
	v175 = v167
	goto L54
L57:
	;
	v166 = int32(1)
	v167 = v161 + v166
	v169 = v159 + v166
	if v169&int32(3) != 0 {
		v159 = v169
		v161 = v167
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v185 = v173
	v186 = v177
	v187 = v175
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v186
	v189 = int32(4)
	v190 = v187 + v189
	v192 = v185 + v189
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v197 = int32(-2139062144)
	if (int32(16843008)-v194|v194)&v197 == v197 {
		v185 = v192
		v186 = v194
		v187 = v190
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v202 = v192
	v203 = v194
	v204 = v190
	goto L48
L62:
	;
	goto L61
L63:
	;
	v211 = v202
	v213 = v204
	goto L64
L64:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)) = uint8(v214)
	v216 = int32(1)
	if v214 != 0 {
		v211 = v211 + v216
		v213 = v213 + v216
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
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+104)))
	v277 = v273&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+104)) = uint8(v277)
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+102)))
	v284 = v279&int32(_a_F_uuid_generate_internal_0) | l0<<(uint(int32(4))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+102)) = uint16(v284)
	F_uuid_unparse(m, v10+int32(96), v10+int32(112))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L43
	} else {
		goto L83
	}
L68:
	;
	v227 = F_pg_cryptohash_create(m, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L43
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v248 = F_pg_cryptohash_create(m, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L43
	} else {
		goto L77
	}
L71:
	;
	v229 = F_pg_cryptohash_init(m, v227)
	mBase = m.M
	if v229 < int32(0) {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v233 = F_pg_cryptohash_update(m, v227, l1, int32(16))
	mBase = m.M
	if v233 < int32(0) {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	v236 = F_pg_cryptohash_update(m, v227, l2, l3)
	mBase = m.M
	if v236 < int32(0) {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	v242 = F_pg_cryptohash_final(m, v227, v10+int32(96), int32(16))
	mBase = m.M
	if v242 < int32(0) {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_pg_cryptohash_free(m, v227)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L43
	} else {
		goto L76
	}
L76:
	;
	goto L67
L77:
	;
	v250 = F_pg_cryptohash_init(m, v248)
	mBase = m.M
	if v250 < int32(0) {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v254 = F_pg_cryptohash_update(m, v248, l1, int32(16))
	mBase = m.M
	if v254 < int32(0) {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v257 = F_pg_cryptohash_update(m, v248, l2, l3)
	mBase = m.M
	if v257 < int32(0) {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v263 = F_pg_cryptohash_final(m, v248, v10+int32(112), int32(20))
	mBase = m.M
	if v263 < int32(0) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_pg_cryptohash_free(m, v248)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L43
	} else {
		goto L82
	}
L82:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v10)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v268
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v10)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = v270
	goto L67
L83:
	;
	goto L7
L84:
	;
	goto L7
L85:
	;
	m.G0 = v10 + int32(160)
	return v306
L86:
	;
	if v227 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_uuid_generate_internal_1)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_2), v10)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L43
	} else {
		goto L97
	}
L88:
	;
	v330 = int32(_a_F_uuid_generate_internal_3)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v322 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v325 = int32(_a_F_uuid_generate_internal_4)
	goto L93
L92:
	;
	v325 = int32(_a_F_uuid_generate_internal_5)
	goto L93
L93:
	;
	if v322 == int32(2) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v328 = int32(_a_F_uuid_generate_internal_3)
	goto L96
L95:
	;
	v328 = v325
	goto L96
L96:
	;
	v330 = v328
	goto L87
L97:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(338), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
	if v227 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_uuid_generate_internal_1)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_8), v10+int32(16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L43
	} else {
		goto L110
	}
L101:
	;
	v360 = int32(_a_F_uuid_generate_internal_3)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v352 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v355 = int32(_a_F_uuid_generate_internal_4)
	goto L106
L105:
	;
	v355 = int32(_a_F_uuid_generate_internal_5)
	goto L106
L106:
	;
	if v352 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v358 = int32(_a_F_uuid_generate_internal_3)
	goto L109
L108:
	;
	v358 = v355
	goto L109
L109:
	;
	v360 = v358
	goto L100
L110:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(342), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
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
	if v227 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_uuid_generate_internal_1)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_9), v10+int32(32))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L43
	} else {
		goto L123
	}
L114:
	;
	v392 = int32(_a_F_uuid_generate_internal_3)
	goto L113
L115:
	;
	goto L116
L116:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v384 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v387 = int32(_a_F_uuid_generate_internal_4)
	goto L119
L118:
	;
	v387 = int32(_a_F_uuid_generate_internal_5)
	goto L119
L119:
	;
	if v384 == int32(2) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v390 = int32(_a_F_uuid_generate_internal_3)
	goto L122
L121:
	;
	v390 = v387
	goto L122
L122:
	;
	v392 = v390
	goto L113
L123:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(347), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	if v248 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a_F_uuid_generate_internal_10)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_2), v10+int32(48))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L43
	} else {
		goto L136
	}
L127:
	;
	v424 = int32(_a_F_uuid_generate_internal_3)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v416 == int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v419 = int32(_a_F_uuid_generate_internal_4)
	goto L132
L131:
	;
	v419 = int32(_a_F_uuid_generate_internal_5)
	goto L132
L132:
	;
	if v416 == int32(2) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v422 = int32(_a_F_uuid_generate_internal_3)
	goto L135
L134:
	;
	v422 = v419
	goto L135
L135:
	;
	v424 = v422
	goto L126
L136:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(357), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
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
	if v248 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(_a_F_uuid_generate_internal_10)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_8), v10-int32(-64))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L43
	} else {
		goto L149
	}
L140:
	;
	v456 = int32(_a_F_uuid_generate_internal_3)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v448 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v451 = int32(_a_F_uuid_generate_internal_4)
	goto L145
L144:
	;
	v451 = int32(_a_F_uuid_generate_internal_5)
	goto L145
L145:
	;
	if v448 == int32(2) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v454 = int32(_a_F_uuid_generate_internal_3)
	goto L148
L147:
	;
	v454 = v451
	goto L148
L148:
	;
	v456 = v454
	goto L139
L149:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(361), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
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
	if v248 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(_a_F_uuid_generate_internal_10)
	F_errmsg_internal(m, int32(_a_F_uuid_generate_internal_9), v10+int32(80))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L43
	} else {
		goto L162
	}
L153:
	;
	v488 = int32(_a_F_uuid_generate_internal_3)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v480 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v483 = int32(_a_F_uuid_generate_internal_4)
	goto L158
L157:
	;
	v483 = int32(_a_F_uuid_generate_internal_5)
	goto L158
L158:
	;
	if v480 == int32(2) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v486 = int32(_a_F_uuid_generate_internal_3)
	goto L161
L160:
	;
	v486 = v483
	goto L161
L161:
	;
	v488 = v486
	goto L152
L162:
	;
	F_errfinish(m, int32(_a_F_uuid_generate_internal_6), int32(364), int32(_a_F_uuid_generate_internal_7))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
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
func F_uuid_in(m *base.Module, l0 int32) int64 {
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	return base.I64_extend_i32_u(v15)
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
	F_errsave_finish(m, v23, int32(_a_F_uuid_in_2), int32(200), int32(_a_F_uuid_in_3))
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
func F_uuid_increment(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = base.I32_wrap_i64(l1)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v14
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v16 != int32(255) {
			v132 = v7 + int32(15)
			v133 = v16
			v135 = v133 + int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
			v137 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
			return base.I64_extend_i32_u(v7)
		} else {
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v21)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
			if v23 != int32(255) {
				v132 = v7 + int32(14)
				v133 = v23
				v135 = v133 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
				v137 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
				return base.I64_extend_i32_u(v7)
			} else {
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v28)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)))
				if v30 != int32(255) {
					v132 = v7 + int32(13)
					v133 = v30
					v135 = v133 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
					v137 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
					return base.I64_extend_i32_u(v7)
				} else {
					v35 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v35)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
					if v37 != int32(255) {
						v132 = v7 + int32(12)
						v133 = v37
						v135 = v133 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
						v137 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
						return base.I64_extend_i32_u(v7)
					} else {
						v42 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v42)
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)))
						if v44 != int32(255) {
							v132 = v7 + int32(11)
							v133 = v44
							v135 = v133 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
							v137 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
							return base.I64_extend_i32_u(v7)
						} else {
							v49 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v49)
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+10)))
							if v51 != int32(255) {
								v132 = v7 + int32(10)
								v133 = v51
								v135 = v133 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
								v137 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
								return base.I64_extend_i32_u(v7)
							} else {
								v56 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v7)+10)) = uint8(v56)
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
								if v58 != int32(255) {
									v132 = v7 + int32(9)
									v133 = v58
									v135 = v133 + int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
									v137 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
									return base.I64_extend_i32_u(v7)
								} else {
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)) = uint8(v63)
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)))
									if v65 != int32(255) {
										v132 = v7 + int32(8)
										v133 = v65
										v135 = v133 + int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
										v137 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
										return base.I64_extend_i32_u(v7)
									} else {
										v70 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v70)
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)))
										if v72 != int32(255) {
											v132 = v7 + int32(7)
											v133 = v72
											v135 = v133 + int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
											v137 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
											return base.I64_extend_i32_u(v7)
										} else {
											v77 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)) = uint8(v77)
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)))
											if v79 != int32(255) {
												v132 = v7 + int32(6)
												v133 = v79
												v135 = v133 + int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
												v137 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
												return base.I64_extend_i32_u(v7)
											} else {
												v84 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)) = uint8(v84)
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)))
												if v86 != int32(255) {
													v132 = v7 + int32(5)
													v133 = v86
													v135 = v133 + int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
													v137 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
													return base.I64_extend_i32_u(v7)
												} else {
													v91 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v91)
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
													if v93 != int32(255) {
														v132 = v7 + int32(4)
														v133 = v93
														v135 = v133 + int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
														v137 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
														return base.I64_extend_i32_u(v7)
													} else {
														v98 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v98)
														v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)))
														if v100 != int32(255) {
															v132 = v7 + int32(3)
															v133 = v100
															v135 = v133 + int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
															v137 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
															return base.I64_extend_i32_u(v7)
														} else {
															v105 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v105)
															v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)))
															if v107 != int32(255) {
																v132 = v7 + int32(2)
																v133 = v107
																v135 = v133 + int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
																v137 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
																return base.I64_extend_i32_u(v7)
															} else {
																v112 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v112)
																v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
																if v114 != int32(255) {
																	v132 = v7 + int32(1)
																	v133 = v114
																	v135 = v133 + int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
																	v137 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
																	return base.I64_extend_i32_u(v7)
																} else {
																	v119 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v119)
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
																	if v121 != int32(255) {
																		v132 = v7
																		v133 = v121
																		v135 = v133 + int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
																		v137 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v137)
																		return base.I64_extend_i32_u(v7)
																	} else {
																		v124 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v124)
																		F_pfree(m, v7)
																		mBase = m.M
																		v127 = m.ExcPending
																		if v127 != 0 {
																			return int64(0)
																		} else {
																			v128 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v128)
																			return int64(0)
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
func F_uuid_ns_x500(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14401(m, l0, int32(_a_F_uuid_ns_x500_0), int32(_a_F_uuid_ns_x500_1), int32(_a_F_uuid_ns_x500_2), int32(_a_F_uuid_ns_x500_3), int32(_a_F_uuid_ns_x500_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_uuid_sortsupport(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14333(m, l0, int32(1742), int32(1741), int32(1740))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
