package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_from_char_parse_int_len(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = v15
	v23 = int32(0)
	goto L1
L1:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v26-int32(9)))&base.B2i32(v26 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v40 = v23 + v15
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
	v43 = v13 + int32(131)
	v45 = l2 + int32(1)
	if v45 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v36 = int32(1)
	v22 = v22 + v36
	v23 = v23 + v36
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)))
	if v164&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L7:
	;
	v161 = F_strlen(m, v157)
	mBase = m.M
	v163 = v161 + (v158 - v43)
	goto L6
L8:
	;
	v157 = v40
	v158 = v43
	goto L7
L9:
	;
	goto L10
L10:
	;
	v51 = v45 - int32(1)
	if (v43^v40)&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v154)
	v157 = v150
	v158 = v151
	goto L7
L12:
	;
	v135 = v130
	v136 = v131
	v137 = v132
	goto L33
L13:
	;
	if v125 == int32(0) {
		v150 = v123
		v151 = v124
		goto L11
	} else {
		goto L32
	}
L14:
	;
	v123 = v40
	v124 = v43
	v125 = v51
	goto L13
L15:
	;
	goto L16
L16:
	;
	v55 = int32(0)
	if base.B2i32(v40&int32(3) == v55)|base.B2i32(v51 == v55) == v55 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v91 == int32(0) {
		v150 = v88
		v151 = v89
		goto L11
	} else {
		goto L26
	}
L18:
	;
	v67 = v40
	v68 = v43
	v69 = v51
	goto L21
L19:
	;
	goto L20
L20:
	;
	v88 = v40
	v89 = v43
	v90 = v51
	v91 = base.B2i32(v51 != v55)
	goto L17
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	*(*uint8)(unsafe.Add(mBase, uint32(v68))) = uint8(v71)
	if v71 == int32(0) {
		v130 = v67
		v131 = v68
		v132 = v69
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v88 = v82
	v89 = v76
	v90 = v78
	v91 = v80
	goto L17
L23:
	;
	v75 = int32(1)
	v76 = v68 + v75
	v78 = v69 - v75
	v79 = int32(0)
	v80 = base.B2i32(v78 != v79)
	v82 = v67 + v75
	if v82&int32(3) == v79 {
		v88 = v82
		v89 = v76
		v90 = v78
		v91 = v80
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v78 != 0 {
		v67 = v82
		v68 = v76
		v69 = v78
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if base.B2i32(v94 == int32(0))|base.B2i32(base.Ui32(v90) < base.Ui32(int32(4))) != 0 {
		v123 = v88
		v124 = v89
		v125 = v90
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v101 = v88
	v102 = v89
	v103 = v90
	goto L28
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v109 = int32(-2139062144)
	if (int32(16843008)-v106|v106)&v109 != v109 {
		v130 = v101
		v131 = v102
		v132 = v103
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v123 = v117
	v124 = v115
	v125 = v119
	goto L13
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v106
	v114 = int32(4)
	v115 = v102 + v114
	v117 = v101 + v114
	v119 = v103 - v114
	if base.Ui32(int32(3)) < base.Ui32(v119) {
		v101 = v117
		v102 = v115
		v103 = v119
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v130 = v123
	v131 = v124
	v132 = v125
	goto L12
L33:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v139)
	if v139 == int32(0) {
		v150 = v135
		v151 = v136
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v150 = v146
	v151 = v144
	goto L11
L35:
	;
	v143 = int32(1)
	v144 = v136 + v143
	v146 = v135 + v143
	v148 = v137 - v143
	if v148 != 0 {
		v135 = v146
		v136 = v144
		v137 = v148
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	m.G0 = v13 + int32(144)
	return v384
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v282
	if v282 == v15 {
		goto L72
	} else {
		goto L73
	}
L39:
	;
	if v163 < l2 {
		goto L54
	} else {
		goto L55
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_from_char_parse_int_len[0])) = int32(0)
	v192 = F_strtox_2(m, v15, v13+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	goto L50
L41:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	switch v167 - int32(1) {
	case 0:
		goto L39
	case 1:
		goto L43
	default:
		goto L42
	}
L42:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	switch v172 - int32(1) {
	case 0:
		goto L40
	case 1:
		goto L45
	default:
		goto L46
	}
L43:
	;
	if v164&int32(6) != 0 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+12)))
	if v184 != 0 {
		goto L39
	} else {
		goto L49
	}
L46:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
	if v175 != 0 {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if base.Ui32(int32(10)) <= base.Ui32((v176-int32(48))&int32(255)) {
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
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	v280 = base.I32_wrap_i64(v192)
	v282 = v194
	goto L38
L51:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v280 = base.I32_wrap_i64(v228)
	v282 = v277 + v231
	goto L38
L52:
	;
	v384 = int32(-1)
	goto L37
L53:
	;
	F_errhint(m, int32(_a_F_from_char_parse_int_len_0), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L57
	} else {
		goto L70
	}
L54:
	;
	v196 = F_errsave_start(m, l4)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
	v223 = v13 + int32(131)
	v228 = F_strtox_2(m, v223, v13+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	goto L63
L57:
	;
	return int32(0)
L58:
	;
	if v196 == int32(0) {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v206
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_1), v13+int32(16))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errdetail(m, int32(_a_F_from_char_parse_int_len_2), v13)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v263 = int32(2247)
	goto L53
L63:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	v231 = v230 - v223
	if base.B2i32(v231 <= int32(0))|base.B2i32(l2 <= v231) != 0 {
		goto L51
	} else {
		goto L64
	}
L64:
	;
	v236 = F_errsave_start(m, l4)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L57
	} else {
		goto L65
	}
L65:
	;
	if v236 == int32(0) {
		goto L52
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v223
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_3), v13+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L57
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l2
	F_errdetail(m, int32(_a_F_from_char_parse_int_len_4), v13+int32(32))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v263 = int32(2261)
	goto L53
L70:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), v263, int32(_a_F_from_char_parse_int_len_6))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L57
	} else {
		goto L71
	}
L71:
	;
	goto L52
L72:
	;
	v285 = int32(-1)
	v286 = F_errsave_start(m, l4)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L57
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_from_char_parse_int_len[0]))
	if v314 == int32(68) {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	if v286 == int32(0) {
		v384 = v285
		goto L37
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L57
	} else {
		goto L77
	}
L77:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v13 + int32(131)
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_3), v13-int32(-64))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L57
	} else {
		goto L78
	}
L78:
	;
	F_errdetail(m, int32(_a_F_from_char_parse_int_len_7), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), int32(2271), int32(_a_F_from_char_parse_int_len_6))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L57
	} else {
		goto L80
	}
L80:
	;
	v384 = v285
	goto L37
L81:
	;
	v317 = int32(-1)
	v318 = F_errsave_start(m, l4)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
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
	if v318 == int32(0) {
		v384 = v317
		goto L37
	} else {
		goto L85
	}
L85:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L57
	} else {
		goto L86
	}
L86:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v326
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_8), v13+int32(96))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L57
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = int64(9223372034707292160)
	F_errdetail(m, int32(_a_F_from_char_parse_int_len_9), v13+int32(80))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L57
	} else {
		goto L88
	}
L88:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), int32(2279), int32(_a_F_from_char_parse_int_len_6))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L57
	} else {
		goto L89
	}
L89:
	;
	v384 = v317
	goto L37
L90:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = int32(0)
	if base.B2i32(v345 == v346)|base.B2i32(v345 == v280) == v346 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v380 = v282
	goto L92
L92:
	;
	v384 = v380 - v15
	goto L37
L93:
	;
	v352 = int32(-1)
	v353 = F_errsave_start(m, l4)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L57
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v280
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v380 = v378
	goto L92
L96:
	;
	if v353 == int32(0) {
		v384 = v352
		goto L37
	} else {
		goto L97
	}
L97:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L57
	} else {
		goto L98
	}
L98:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v361
	F_errmsg(m, int32(_a_F_from_char_parse_int_len_10), v13+int32(112))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L57
	} else {
		goto L99
	}
L99:
	;
	F_errdetail(m, int32(_a_F_from_char_parse_int_len_11), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L57
	} else {
		goto L100
	}
L100:
	;
	F_errsave_finish(m, l4, int32(_a_F_from_char_parse_int_len_5), int32(2176), int32(_a_F_from_char_parse_int_len_12))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L57
	} else {
		goto L101
	}
L101:
	;
	v384 = v352
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
