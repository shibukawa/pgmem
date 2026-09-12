package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_timezone_abbreviations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v384 int32
	_ = v384
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v18 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(14111), v18, int32(1024), int32(8192))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v33 = int32(4515712)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(128)
	v40 = F_palloc(m, int32(3072))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v40
	v43 = int32(0)
	v49 = F_ParseTzFile(m, v13, v43, v21+int32(12), v21+int32(8), v43)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v34
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L89
	}
L8:
	;
	if v49 < int32(0) {
		v371 = v18
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v57 = v49<<(uint(int32(4))%32) | int32(8)
	if int32(0) < v49 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = v57
	v63 = v4
	goto L13
L11:
	;
	v88 = v57
	goto L12
L12:
	;
	v98 = F_guc_malloc(m, v88)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L20
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v53+v63*int32(24))+4))
	if v75 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v88 = v82
	goto L12
L15:
	;
	v76 = F_strlen(m, v75)
	mBase = m.M
	v82 = (v76+int32(12))&int32(-8) + v62
	goto L17
L16:
	;
	v82 = v62
	goto L17
L17:
	;
	v84 = v63 + int32(1)
	if v84 != v49 {
		v62 = v82
		v63 = v84
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v98 != 0 {
		v371 = v98
		goto L7
	} else {
		goto L86
	}
L20:
	;
	if v98 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v88
	if v49 <= int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v112 = int32(0)
	v115 = v57
	goto L23
L23:
	;
	v123 = v98 + int32(8) + v112<<(uint(int32(4))%32)
	v126 = v53 + v112*int32(24)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	goto L28
L24:
	;
	goto L19
L25:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v243 != 0 {
		goto L58
	} else {
		goto L59
	}
L26:
	;
	v240 = F_strlen(m, v229)
	mBase = m.M
	goto L25
L28:
	;
	goto L29
L29:
	;
	v134 = int32(10)
	if (v123^v127)&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v233)
	goto L26
L31:
	;
	v214 = v209
	v215 = v210
	v216 = v211
	goto L53
L32:
	;
	if v204 == int32(0) {
		v229 = v202
		v230 = v203
		goto L30
	} else {
		goto L52
	}
L33:
	;
	v202 = v127
	v203 = v123
	v204 = v134
	goto L32
L34:
	;
	goto L35
L35:
	;
	if v127&int32(3) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v171 == int32(0) {
		v229 = v168
		v230 = v169
		goto L30
	} else {
		goto L45
	}
L37:
	;
	v168 = v127
	v169 = v123
	v170 = v134
	v171 = int32(1)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v147 = v127
	v148 = v123
	v149 = v134
	goto L40
L40:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v151)
	if v151 == int32(0) {
		v209 = v147
		v210 = v148
		v211 = v149
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v168 = v162
	v169 = v156
	v170 = v158
	v171 = v160
	goto L36
L42:
	;
	v155 = int32(1)
	v156 = v148 + v155
	v158 = v149 - v155
	v159 = int32(0)
	v160 = base.B2i32(v158 != v159)
	v162 = v147 + v155
	if v162&int32(3) == v159 {
		v168 = v162
		v169 = v156
		v170 = v158
		v171 = v160
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v158 != 0 {
		v147 = v162
		v148 = v156
		v149 = v158
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v174 == int32(0) {
		v202 = v168
		v203 = v169
		v204 = v170
		goto L32
	} else {
		goto L46
	}
L46:
	;
	if base.Ui32(v170) < base.Ui32(int32(4)) {
		v202 = v168
		v203 = v169
		v204 = v170
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v180 = v168
	v181 = v169
	v182 = v170
	goto L48
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v188 = int32(-2139062144)
	if (int32(16843008)-v185|v185)&v188 != v188 {
		v209 = v180
		v210 = v181
		v211 = v182
		goto L31
	} else {
		goto L50
	}
L49:
	;
	v202 = v196
	v203 = v194
	v204 = v198
	goto L32
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v185
	v193 = int32(4)
	v194 = v181 + v193
	v196 = v180 + v193
	v198 = v182 - v193
	if base.Ui32(int32(3)) < base.Ui32(v198) {
		v180 = v196
		v181 = v194
		v182 = v198
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v209 = v202
	v210 = v203
	v211 = v204
	goto L31
L53:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v218)
	if v218 == int32(0) {
		v229 = v214
		v230 = v215
		goto L30
	} else {
		goto L55
	}
L54:
	;
	v229 = v225
	v230 = v223
	goto L30
L55:
	;
	v222 = int32(1)
	v223 = v215 + v222
	v225 = v214 + v222
	v227 = v216 - v222
	if v227 != 0 {
		v214 = v225
		v215 = v223
		v216 = v227
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v344 = v112 + int32(1)
	if v344 != v49 {
		v112 = v344
		v115 = v341
		goto L23
	} else {
		goto L85
	}
L58:
	;
	v244 = v98 + v115
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = int32(0)
	v248 = v244 + int32(4)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if (v249^v248)&int32(3) != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	goto L60
L60:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+12)))
	if v336 != 0 {
		goto L82
	} else {
		goto L83
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v115
	v325 = int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+11)) = uint8(v325)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v328 = F_strlen(m, v327)
	mBase = m.M
	v341 = (v328+int32(12))&int32(-8) + v115
	goto L57
L62:
	;
	goto L61
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v303)
	if v303&int32(255) == int32(0) {
		goto L62
	} else {
		goto L78
	}
L64:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v302 = v249
	v303 = v255
	v304 = v248
	goto L63
L65:
	;
	goto L66
L66:
	;
	if v249&int32(3) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v259 = v249
	v261 = v248
	goto L70
L68:
	;
	v273 = v249
	v275 = v248
	goto L69
L69:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v280 = int32(-2139062144)
	if (int32(16843008)-v277|v277)&v280 != v280 {
		v302 = v273
		v303 = v277
		v304 = v275
		goto L63
	} else {
		goto L74
	}
L70:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v262)
	if v262 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L71:
	;
	v273 = v269
	v275 = v267
	goto L69
L72:
	;
	v266 = int32(1)
	v267 = v261 + v266
	v269 = v259 + v266
	if v269&int32(3) != 0 {
		v259 = v269
		v261 = v267
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v285 = v273
	v286 = v277
	v287 = v275
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v286
	v289 = int32(4)
	v290 = v287 + v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v293 = v285 + v289
	v297 = int32(-2139062144)
	if (v291|(int32(16843008)-v291))&v297 == v297 {
		v285 = v293
		v286 = v291
		v287 = v290
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v302 = v293
	v303 = v291
	v304 = v290
	goto L63
L77:
	;
	goto L76
L78:
	;
	v311 = v302
	v313 = v304
	goto L79
L79:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)) = uint8(v314)
	v316 = int32(1)
	if v314 != 0 {
		v311 = v311 + v316
		v313 = v313 + v316
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L62
L81:
	;
	goto L80
L82:
	;
	v337 = int32(6)
	goto L84
L83:
	;
	v337 = int32(5)
	goto L84
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+11)) = uint8(v337)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v339
	v341 = v115
	goto L57
L85:
	;
	goto L24
L86:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v360
	goto L87
L87:
	;
	v366 = F_format_elog_string(m, int32(13904), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v366
	v371 = int32(0)
	goto L7
L89:
	;
	m.G0 = v21 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v371
	return base.B2i32(v371 != int32(0))
}
