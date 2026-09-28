package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_from_char_parse_int_len(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int64
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = v16
	v24 = int32(0)
	goto L1
L1:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v28-int32(9)))&base.B2i32(v28 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v42 = v24 + v16
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
	v45 = v14 + int32(131)
	v47 = l2 + int32(1)
	if v47 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v38 = int32(1)
	v23 = v23 + v38
	v24 = v24 + v38
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v166&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L7:
	;
	v163 = F_strlen(m, v159)
	mBase = m.M
	v165 = v163 + (v160 - v45)
	goto L6
L8:
	;
	v159 = v42
	v160 = v45
	goto L7
L9:
	;
	goto L10
L10:
	;
	v53 = v47 - int32(1)
	if (v45^v42)&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v156)
	v159 = v152
	v160 = v153
	goto L7
L12:
	;
	v137 = v132
	v138 = v133
	v139 = v134
	goto L33
L13:
	;
	if v127 == int32(0) {
		v152 = v125
		v153 = v126
		goto L11
	} else {
		goto L32
	}
L14:
	;
	v125 = v42
	v126 = v45
	v127 = v53
	goto L13
L15:
	;
	goto L16
L16:
	;
	v57 = int32(0)
	if base.B2i32(v42&int32(3) == v57)|base.B2i32(v53 == v57) == v57 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v93 == int32(0) {
		v152 = v90
		v153 = v91
		goto L11
	} else {
		goto L26
	}
L18:
	;
	v69 = v42
	v70 = v45
	v71 = v53
	goto L21
L19:
	;
	goto L20
L20:
	;
	v90 = v42
	v91 = v45
	v92 = v53
	v93 = base.B2i32(v53 != v57)
	goto L17
L21:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v73)
	if v73 == int32(0) {
		v132 = v69
		v133 = v70
		v134 = v71
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v90 = v84
	v91 = v78
	v92 = v80
	v93 = v82
	goto L17
L23:
	;
	v77 = int32(1)
	v78 = v70 + v77
	v80 = v71 - v77
	v81 = int32(0)
	v82 = base.B2i32(v80 != v81)
	v84 = v69 + v77
	if v84&int32(3) == v81 {
		v90 = v84
		v91 = v78
		v92 = v80
		v93 = v82
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v80 != 0 {
		v69 = v84
		v70 = v78
		v71 = v80
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if base.B2i32(v96 == int32(0))|base.B2i32(base.Ui32(v92) < base.Ui32(int32(4))) != 0 {
		v125 = v90
		v126 = v91
		v127 = v92
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v103 = v90
	v104 = v91
	v105 = v92
	goto L28
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v111 = int32(-2139062144)
	if (int32(16843008)-v108|v108)&v111 != v111 {
		v132 = v103
		v133 = v104
		v134 = v105
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v125 = v119
	v126 = v117
	v127 = v121
	goto L13
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v108
	v116 = int32(4)
	v117 = v104 + v116
	v119 = v103 + v116
	v121 = v105 - v116
	if base.Ui32(int32(3)) < base.Ui32(v121) {
		v103 = v119
		v104 = v117
		v105 = v121
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v132 = v125
	v133 = v126
	v134 = v127
	goto L12
L33:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v141)
	if v141 == int32(0) {
		v152 = v137
		v153 = v138
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v152 = v148
	v153 = v146
	goto L11
L35:
	;
	v145 = int32(1)
	v146 = v138 + v145
	v148 = v137 + v145
	v150 = v139 - v145
	if v150 != 0 {
		v137 = v148
		v138 = v146
		v139 = v150
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	m.G0 = v14 + int32(144)
	return v388
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v286
	if v286 == v16 {
		goto L72
	} else {
		goto L73
	}
L39:
	;
	if base.Ui32(v165) < base.Ui32(l2) {
		goto L54
	} else {
		goto L55
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_from_char_parse_int_len[0])) = int32(0)
	v194 = F_strtox_2(m, v16, v14+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	goto L50
L41:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v169 - int32(1) {
	case 0:
		goto L39
	case 1:
		goto L43
	default:
		goto L42
	}
L42:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	switch v174 - int32(1) {
	case 0:
		goto L40
	case 1:
		goto L45
	default:
		goto L46
	}
L43:
	;
	if v166&int32(6) != 0 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+12)))
	if v186 != 0 {
		goto L39
	} else {
		goto L49
	}
L46:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if v177 != 0 {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if base.Ui32(int32(10)) <= base.Ui32((v178-int32(48))&int32(255)) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	goto L39
L49:
	;
	goto L40
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v283 = base.I32_wrap_i64(v194)
	v286 = v196
	goto L38
L51:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v283 = base.I32_wrap_i64(v230)
	v286 = v280 + v233
	goto L38
L52:
	;
	v388 = int32(-1)
	goto L37
L53:
	;
	F_errhint(m, int32(_a_F_from_char_parse_int_len_0), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L57
	} else {
		goto L70
	}
L54:
	;
	v198 = F_errsave_start(m, l4)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_from_char_parse_int_len[0])) = int32(0)
	v225 = v14 + int32(131)
	v230 = F_strtox_2(m, v225, v14+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	goto L63
L57:
	;
	return int32(0)
L58:
	;
	if v198 == int32(0) {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v208
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_1), v14+int32(16))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l2
	v218 = F_errdetail(m, int32(_a_F_from_char_parse_int_len_2), v14)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v265 = int32(2221)
	goto L53
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v233 = v232 - v225
	if base.B2i32(v232 == v225)|base.B2i32(base.Ui32(l2) <= base.Ui32(v233)) != 0 {
		goto L51
	} else {
		goto L64
	}
L64:
	;
	v237 = F_errsave_start(m, l4)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L57
	} else {
		goto L65
	}
L65:
	;
	if v237 == int32(0) {
		goto L52
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v225
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_3), v14+int32(48))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L57
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
	v258 = F_errdetail(m, int32(_a_F_from_char_parse_int_len_4), v14+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v265 = int32(2234)
	goto L53
L70:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), v265, int32(_a_F_from_char_parse_int_len_6))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L57
	} else {
		goto L71
	}
L71:
	;
	goto L52
L72:
	;
	v289 = int32(-1)
	v290 = F_errsave_start(m, l4)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L57
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_from_char_parse_int_len[0]))
	if v318 == int32(68) {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	if v290 == int32(0) {
		v388 = v289
		goto L37
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L57
	} else {
		goto L77
	}
L77:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v14 + int32(131)
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_3), v14-int32(-64))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L57
	} else {
		goto L78
	}
L78:
	;
	v310 = F_errdetail(m, int32(_a_F_from_char_parse_int_len_7), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), int32(2244), int32(_a_F_from_char_parse_int_len_6))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L57
	} else {
		goto L80
	}
L80:
	;
	v388 = v289
	goto L37
L81:
	;
	v321 = int32(-1)
	v322 = F_errsave_start(m, l4)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L57
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if l0 != 0 {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	if v322 == int32(0) {
		v388 = v321
		goto L37
	} else {
		goto L85
	}
L85:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L57
	} else {
		goto L86
	}
L86:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v330
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_8), v14+int32(96))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L57
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = int64(9223372034707292160)
	v342 = F_errdetail(m, int32(_a_F_from_char_parse_int_len_9), v14+int32(80))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L57
	} else {
		goto L88
	}
L88:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), int32(2252), int32(_a_F_from_char_parse_int_len_6))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L57
	} else {
		goto L89
	}
L89:
	;
	v388 = v321
	goto L37
L90:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v350 = int32(0)
	if base.B2i32(v349 == v350)|base.B2i32(v349 == v283) == v350 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v384 = v286
	goto L92
L92:
	;
	v388 = v384 - v16
	goto L37
L93:
	;
	v356 = int32(-1)
	v357 = F_errsave_start(m, l4)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L57
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v283
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v384 = v382
	goto L92
L96:
	;
	if v357 == int32(0) {
		v388 = v356
		goto L37
	} else {
		goto L97
	}
L97:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L57
	} else {
		goto L98
	}
L98:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v365
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_10), v14+int32(112))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L57
	} else {
		goto L99
	}
L99:
	;
	v374 = F_errdetail(m, int32(_a_F_from_char_parse_int_len_11), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L57
	} else {
		goto L100
	}
L100:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), int32(2151), int32(_a_F_from_char_parse_int_len_12))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L57
	} else {
		goto L101
	}
L101:
	;
	v388 = v356
	goto L37
}
func F_get_from_clause(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v31 = int32(1)
	v33 = v4
	goto L4
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v33<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 == int32(63) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v195 = v33 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v195 < v196 {
		v31 = v189
		v33 = v195
		goto L4
	} else {
		goto L45
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32)-int32(4))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+125)))
	if v53 != int32(1) {
		v189 = v31
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	v189 = int32(0)
	goto L6
L12:
	;
	F_appendContextKeyword(m, l2, l1, int32(-8), int32(8), int32(2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_from_clause_0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L18
	}
L15:
	;
	return
L16:
	;
	F_get_from_clause_item(m, v40, l0, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_initStringInfo(m, v14)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v14
	F_get_from_clause_item(m, v40, l0, l2)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v72&int32(2) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	F_appendBinaryStringInfo(m, v24, v164, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L43
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v77 < int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v80 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v121 = F_strlen(m, v117)
	mBase = m.M
	v128 = v121 + int32(1)
	goto L34
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != int32(10) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v87 <= int32(0) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v93 = v87
	goto L28
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v93-int32(1)))))
	if v105 != int32(32) {
		goto L21
	} else {
		goto L30
	}
L29:
	;
	goto L21
L30:
	;
	v109 = v93 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v109
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109+v101))) = uint8(v112)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v112 < v114 {
		v93 = v114
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	if v140 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v130 = int32(0)
	if v128 == v130 {
		v140 = v130
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v140 = v135
	goto L33
L36:
	;
	v134 = v128 - int32(1)
	v135 = v117 + v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v136 != int32(10) {
		v128 = v134
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v143 = v140 + int32(1)
	goto L40
L39:
	;
	v143 = v117
	goto L40
L40:
	;
	v144 = F_strlen(m, v143)
	mBase = m.M
	if base.Ui32(v144+v80) <= base.Ui32(v77) {
		goto L21
	} else {
		goto L41
	}
L41:
	;
	F_appendContextKeyword(m, l2, int32(_a_F_get_from_clause_1), int32(-8), int32(8), int32(4))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_pfree(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	goto L11
L45:
	;
	goto L5
}
func F_get_from_clause_coldeflist(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v26 = int32(0)
	goto L3
L3:
	;
	v38 = int32(0)
	if v24 == v38 {
		v48 = v38
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	v49 = int32(0)
	if v23 == v49 {
		v60 = v49
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v42 <= v26 {
		v48 = int32(0)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v48 = v44 + v26<<(uint(int32(2))%32)
	goto L5
L8:
	;
	if v22 == int32(0) {
		v69 = v49
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v54 <= v26 {
		v60 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v60 = v56 + v26<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v70 = int32(0)
	if v21 == v70 {
		v79 = v70
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v63 <= v26 {
		v69 = v49
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v69 = v65 + v26<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v80 = int32(0)
	if base.B2i32(v48 == v80)|base.B2i32(v60 == v80)|(base.B2i32(v69 == v80)|base.B2i32(v79 == v80)) == v80 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v73 <= v26 {
		v79 = v70
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v79 = v75 + v26<<(uint(int32(2))%32)
	goto L14
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	goto L4
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if int32(0) < v26 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v103 = v96 + v26<<(uint(int32(2))%32)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v103 = v100 + int32(4)
	goto L20
L24:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_coldeflist_0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v110 = F_quote_identifier(m, v104)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v112 = F_format_type_with_typemod(m, v95, v94)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v110
	F_appendStringInfo(m, v17, int32(_a_F_get_from_clause_coldeflist_1), v15+int32(16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v93 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v26 = v26 + int32(1)
	goto L3
L32:
	;
	v123 = F_get_typcollation(m, v95)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v123 == v93 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v126 = F_generate_collation_name(m, v93)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v126
	F_appendStringInfo(m, v17, int32(_a_F_get_from_clause_coldeflist_2), v15)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	m.G0 = v15 + int32(32)
	return
}
func F_pull_from_mbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v10
	v12 = v9 - v10
	if l2 < v12 {
		v14 = l2
	} else {
		v14 = v12
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + v14
	return v14
}
func F_transformFromClause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v89 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v3
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v33 = F_transformFromClauseItem(m, l0, v28, v10+int32(12), v10+int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	F_checkNameSpaceConflicts(m, v35, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v36 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v71 = F_lappend(m, v70, v33)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L15
	}
L10:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v42 <= v41 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v47 = v41
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v47<<(uint(int32(2))%32))))
	v57 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+22)) = uint16(v57)
	v60 = v47 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v60 < v61 {
		v47 = v60
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v75 = F_list_concat(m, v74, v36)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v75
	v79 = v22 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v79 < v80 {
		v22 = v79
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L5
L18:
	;
	m.G0 = v10 + int32(16)
	return
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v92 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v98 = int32(0)
	goto L21
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v98<<(uint(int32(2))%32))))
	v108 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+22)) = uint16(v108)
	v111 = v98 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v111 < v112 {
		v98 = v111
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	goto L22
}
