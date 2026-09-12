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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
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
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)))
	if v160&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L7:
	;
	v157 = F_strlen(m, v153)
	mBase = m.M
	v159 = v157 + (v154 - v43)
	goto L6
L8:
	;
	v153 = v40
	v154 = v43
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
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v150)
	v153 = v146
	v154 = v147
	goto L7
L12:
	;
	v131 = v126
	v132 = v127
	v133 = v128
	goto L34
L13:
	;
	if v121 == int32(0) {
		v146 = v119
		v147 = v120
		goto L11
	} else {
		goto L33
	}
L14:
	;
	v119 = v40
	v120 = v43
	v121 = v51
	goto L13
L15:
	;
	goto L16
L16:
	;
	v55 = int32(0)
	if v40&int32(3) == v55 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v88 == int32(0) {
		v146 = v85
		v147 = v86
		goto L11
	} else {
		goto L26
	}
L18:
	;
	v85 = v40
	v86 = v43
	v87 = v51
	v88 = base.B2i32(v51 != v55)
	goto L17
L19:
	;
	if v51 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v64 = v40
	v65 = v43
	v66 = v51
	goto L21
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v68)
	if v68 == int32(0) {
		v126 = v64
		v127 = v65
		v128 = v66
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v85 = v79
	v86 = v73
	v87 = v75
	v88 = v77
	goto L17
L23:
	;
	v72 = int32(1)
	v73 = v65 + v72
	v75 = v66 - v72
	v76 = int32(0)
	v77 = base.B2i32(v75 != v76)
	v79 = v64 + v72
	if v79&int32(3) == v76 {
		v85 = v79
		v86 = v73
		v87 = v75
		v88 = v77
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v75 != 0 {
		v64 = v79
		v65 = v73
		v66 = v75
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v91 == int32(0) {
		v119 = v85
		v120 = v86
		v121 = v87
		goto L13
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v87) < base.Ui32(int32(4)) {
		v119 = v85
		v120 = v86
		v121 = v87
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v97 = v85
	v98 = v86
	v99 = v87
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v105 = int32(-2139062144)
	if (int32(16843008)-v102|v102)&v105 != v105 {
		v126 = v97
		v127 = v98
		v128 = v99
		goto L12
	} else {
		goto L31
	}
L30:
	;
	v119 = v113
	v120 = v111
	v121 = v115
	goto L13
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v102
	v110 = int32(4)
	v111 = v98 + v110
	v113 = v97 + v110
	v115 = v99 - v110
	if base.Ui32(int32(3)) < base.Ui32(v115) {
		v97 = v113
		v98 = v111
		v99 = v115
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v126 = v119
	v127 = v120
	v128 = v121
	goto L12
L34:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
	if v135 == int32(0) {
		v146 = v131
		v147 = v132
		goto L11
	} else {
		goto L36
	}
L35:
	;
	v146 = v142
	v147 = v140
	goto L11
L36:
	;
	v139 = int32(1)
	v140 = v132 + v139
	v142 = v131 + v139
	v144 = v133 - v139
	if v144 != 0 {
		v131 = v142
		v132 = v140
		v133 = v144
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	m.G0 = v13 + int32(144)
	return v374
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v277
	if v277 == v15 {
		goto L74
	} else {
		goto L75
	}
L40:
	;
	if v159 < l2 {
		goto L55
	} else {
		goto L56
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v188 = F_strtox_2(m, v15, v13+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	goto L51
L42:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	switch v163 - int32(1) {
	case 0:
		goto L40
	case 1:
		goto L44
	default:
		goto L43
	}
L43:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	switch v168 - int32(1) {
	case 0:
		goto L41
	case 1:
		goto L46
	default:
		goto L47
	}
L44:
	;
	if v160&int32(6) != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+12)))
	if v180 != 0 {
		goto L40
	} else {
		goto L50
	}
L47:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
	if v171 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if base.Ui32(int32(10)) <= base.Ui32((v172-int32(48))&int32(255)) {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	goto L40
L50:
	;
	goto L41
L51:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	v276 = base.I32_wrap_i64(v188)
	v277 = v190
	goto L39
L52:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v276 = base.I32_wrap_i64(v224)
	v277 = v273 + v229
	goto L39
L53:
	;
	v374 = int32(-1)
	goto L38
L54:
	;
	F_errhint(m, int32(573522), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L58
	} else {
		goto L72
	}
L55:
	;
	v192 = F_errsave_start(m, l4)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v224 = F_strtox_2(m, v13+int32(131), v13+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	goto L64
L58:
	;
	return int32(0)
L59:
	;
	if v192 == int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v202
	F_errmsg(m, int32(414817), v13+int32(16))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errdetail(m, int32(582769), v13)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v261 = int32(2247)
	goto L54
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	v229 = v226 - (v13 + int32(131))
	if v229 <= int32(0) {
		goto L52
	} else {
		goto L65
	}
L65:
	;
	if l2 <= v229 {
		goto L52
	} else {
		goto L66
	}
L66:
	;
	v233 = F_errsave_start(m, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L58
	} else {
		goto L67
	}
L67:
	;
	if v233 == int32(0) {
		goto L53
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L58
	} else {
		goto L69
	}
L69:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v13 + int32(131)
	F_errmsg(m, int32(662981), v13+int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L58
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l2
	F_errdetail(m, int32(611872), v13+int32(32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L58
	} else {
		goto L71
	}
L71:
	;
	v261 = int32(2261)
	goto L54
L72:
	;
	F_errsave_finish(m, l4, int32(478903), v261, int32(270179))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L58
	} else {
		goto L73
	}
L73:
	;
	goto L53
L74:
	;
	v280 = int32(-1)
	v281 = F_errsave_start(m, l4)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L58
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v309 == int32(68) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	if v281 == int32(0) {
		v374 = v280
		goto L38
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L58
	} else {
		goto L79
	}
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v13 + int32(131)
	F_errmsg(m, int32(662981), v13-int32(-64))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L58
	} else {
		goto L80
	}
L80:
	;
	F_errdetail(m, int32(574133), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L58
	} else {
		goto L81
	}
L81:
	;
	F_errsave_finish(m, l4, int32(478903), int32(2271), int32(270179))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L58
	} else {
		goto L82
	}
L82:
	;
	v374 = v280
	goto L38
L83:
	;
	v312 = int32(-1)
	v313 = F_errsave_start(m, l4)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L58
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if l0 != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	if v313 == int32(0) {
		v374 = v312
		goto L38
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L58
	} else {
		goto L88
	}
L88:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v321
	F_errmsg(m, int32(386366), v13+int32(96))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L58
	} else {
		goto L89
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = int64(9223372034707292160)
	F_errdetail(m, int32(616602), v13+int32(80))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L58
	} else {
		goto L90
	}
L90:
	;
	F_errsave_finish(m, l4, int32(478903), int32(2279), int32(270179))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L58
	} else {
		goto L91
	}
L91:
	;
	v374 = v312
	goto L38
L92:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v340 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v372 = v277
	goto L94
L94:
	;
	v374 = v372 - v15
	goto L38
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v276
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v372 = v370
	goto L94
L96:
	;
	if v340 == v276 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v344 = int32(-1)
	v345 = F_errsave_start(m, l4)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L58
	} else {
		goto L98
	}
L98:
	;
	if v345 == int32(0) {
		v374 = v344
		goto L38
	} else {
		goto L99
	}
L99:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L58
	} else {
		goto L100
	}
L100:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v353
	F_errmsg(m, int32(316634), v13+int32(112))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L58
	} else {
		goto L101
	}
L101:
	;
	F_errdetail(m, int32(597111), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L58
	} else {
		goto L102
	}
L102:
	;
	F_errsave_finish(m, l4, int32(478903), int32(2176), int32(87245))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L58
	} else {
		goto L103
	}
L103:
	;
	v374 = v344
	goto L38
}
func F_get_from_clause(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v30 = int32(1)
	v32 = v4
	goto L4
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 == int32(63) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v247 = v32 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v247 < v248 {
		v30 = v242
		v32 = v247
		goto L4
	} else {
		goto L62
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44<<(uint(int32(2))%32)-int32(4))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+125)))
	if v51 != int32(1) {
		v242 = v30
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v30&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	v242 = int32(0)
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
	F_appendStringInfoString(m, v22, int32(708326))
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
	F_get_from_clause_item(m, v38, l0, l2)
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
	F_initStringInfo(m, v13)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v13
	F_get_from_clause_item(m, v38, l0, l2)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v72&int32(2) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_appendBinaryStringInfo(m, v22, v218, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L15
	} else {
		goto L60
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
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v80 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v120 = F_strlen(m, v116)
	mBase = m.M
	v127 = v120 + int32(1)
	goto L34
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != int32(10) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
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
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v93-int32(1)))))
	if v104 != int32(32) {
		goto L21
	} else {
		goto L30
	}
L29:
	;
	goto L21
L30:
	;
	v108 = v93 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v108
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+v100))) = uint8(v111)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v111 < v113 {
		v93 = v113
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	if v139 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v129 = int32(0)
	if v127 == v129 {
		v139 = v129
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v139 = v134
	goto L33
L36:
	;
	v133 = v127 - int32(1)
	v134 = v116 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v135 != int32(10) {
		v127 = v133
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v142 = v139 + int32(1)
	goto L40
L39:
	;
	v142 = v116
	goto L40
L40:
	;
	if v142&int32(3) == int32(0) {
		v166 = v142
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if base.Ui32(v199+v80) <= base.Ui32(v77) {
		goto L21
	} else {
		goto L58
	}
L42:
	;
	v199 = v191 - v142
	goto L41
L43:
	;
	v170 = v166
	goto L52
L44:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v150 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v199 = int32(0)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v155 = v142
	goto L48
L48:
	;
	v159 = v155 + int32(1)
	if v159&int32(3) == int32(0) {
		v166 = v159
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v191 = v159
	goto L42
L50:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v164 != 0 {
		v155 = v159
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v179 = int32(-2139062144)
	if (int32(16843008)-v176|v176)&v179 == v179 {
		v170 = v170 + int32(4)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v185 = v170
	goto L55
L54:
	;
	goto L53
L55:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v189 != 0 {
		v185 = v185 + int32(1)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v191 = v185
	goto L42
L57:
	;
	goto L56
L58:
	;
	F_appendContextKeyword(m, l2, int32(719562), int32(-8), int32(8), int32(4))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	goto L21
L60:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	F_pfree(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	goto L11
L62:
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
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
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L39
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
	if v48 == int32(0) {
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
	goto L4
L18:
	;
	if v60 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v69 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if v79 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if int32(0) < v26 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v98 = v91 + v26<<(uint(int32(2))%32)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v98 = v95 + int32(4)
	goto L22
L26:
	;
	F_appendStringInfoString(m, v17, int32(708326))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v105 = F_quote_identifier(m, v99)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v107 = F_format_type_with_typemod(m, v90, v89)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v105
	F_appendStringInfo(m, v17, int32(172130), v15+int32(16))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v88 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v26 = v26 + int32(1)
	goto L3
L34:
	;
	v118 = F_get_typcollation(m, v90)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v118 == v88 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v121 = F_generate_collation_name(m, v88)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v121
	F_appendStringInfo(m, v17, int32(189607), v15)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
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
	var v48 int32
	_ = v48
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
	var v99 int32
	_ = v99
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
	v48 = v41
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v57 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+22)) = uint16(v57)
	v60 = v48 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v60 < v61 {
		v48 = v60
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
	v99 = int32(0)
	goto L21
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v99<<(uint(int32(2))%32))))
	v108 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+22)) = uint16(v108)
	v111 = v99 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v111 < v112 {
		v99 = v111
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
