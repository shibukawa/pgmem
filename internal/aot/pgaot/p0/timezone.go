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
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
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
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[0]))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(_a_F_check_timezone_abbreviations_0), v18, int32(1024), int32(_a_F_check_timezone_abbreviations_1))
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
	v33 = int32(_a_F_check_timezone_abbreviations_2)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[0])) = v29
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
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[0])) = v34
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L85
	}
L8:
	;
	if v49 < int32(0) {
		v374 = v18
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
		v374 = v98
		goto L7
	} else {
		goto L82
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
	v114 = v57
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
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v247 != 0 {
		goto L57
	} else {
		goto L58
	}
L26:
	;
	v244 = F_strlen(m, v233)
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
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v237)
	goto L26
L31:
	;
	v218 = v213
	v219 = v214
	v220 = v215
	goto L52
L32:
	;
	if v208 == int32(0) {
		v233 = v206
		v234 = v207
		goto L30
	} else {
		goto L51
	}
L33:
	;
	v206 = v127
	v207 = v123
	v208 = v134
	goto L32
L34:
	;
	goto L35
L35:
	;
	v138 = int32(0)
	if base.B2i32(v127&int32(3) == v138)|int32(0) == v138 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v174 == int32(0) {
		v233 = v171
		v234 = v172
		goto L30
	} else {
		goto L45
	}
L37:
	;
	v150 = v127
	v151 = v123
	v152 = v134
	goto L40
L38:
	;
	goto L39
L39:
	;
	v171 = v127
	v172 = v123
	v173 = v134
	v174 = int32(1)
	goto L36
L40:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v154)
	if v154 == int32(0) {
		v213 = v150
		v214 = v151
		v215 = v152
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v171 = v165
	v172 = v159
	v173 = v161
	v174 = v163
	goto L36
L42:
	;
	v158 = int32(1)
	v159 = v151 + v158
	v161 = v152 - v158
	v162 = int32(0)
	v163 = base.B2i32(v161 != v162)
	v165 = v150 + v158
	if v165&int32(3) == v162 {
		v171 = v165
		v172 = v159
		v173 = v161
		v174 = v163
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v161 != 0 {
		v150 = v165
		v151 = v159
		v152 = v161
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if base.B2i32(v177 == int32(0))|base.B2i32(base.Ui32(v173) < base.Ui32(int32(4))) != 0 {
		v206 = v171
		v207 = v172
		v208 = v173
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v184 = v171
	v185 = v172
	v186 = v173
	goto L47
L47:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v192 = int32(-2139062144)
	if (int32(16843008)-v189|v189)&v192 != v192 {
		v213 = v184
		v214 = v185
		v215 = v186
		goto L31
	} else {
		goto L49
	}
L48:
	;
	v206 = v200
	v207 = v198
	v208 = v202
	goto L32
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v189
	v197 = int32(4)
	v198 = v185 + v197
	v200 = v184 + v197
	v202 = v186 - v197
	if base.Ui32(int32(3)) < base.Ui32(v202) {
		v184 = v200
		v185 = v198
		v186 = v202
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v213 = v206
	v214 = v207
	v215 = v208
	goto L31
L52:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v222)
	if v222 == int32(0) {
		v233 = v218
		v234 = v219
		goto L30
	} else {
		goto L54
	}
L53:
	;
	v233 = v229
	v234 = v227
	goto L30
L54:
	;
	v226 = int32(1)
	v227 = v219 + v226
	v229 = v218 + v226
	v231 = v220 - v226
	if v231 != 0 {
		v218 = v229
		v219 = v227
		v220 = v231
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v347 = v112 + int32(1)
	if v347 != v49 {
		v112 = v347
		v114 = v344
		goto L23
	} else {
		goto L81
	}
L57:
	;
	v248 = v98 + v114
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = int32(0)
	v252 = v248 + int32(4)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if (v253^v252)&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	goto L59
L59:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+12)))
	v340 = v338 + int32(5)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+11)) = uint8(v340)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v342
	v344 = v114
	goto L56
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v114
	v329 = int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+11)) = uint8(v329)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v332 = F_strlen(m, v331)
	mBase = m.M
	v344 = (v332+int32(12))&int32(-8) + v114
	goto L56
L61:
	;
	goto L60
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v307)
	if v307&int32(255) == int32(0) {
		goto L61
	} else {
		goto L77
	}
L63:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v306 = v253
	v307 = v259
	v308 = v252
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v253&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v263 = v253
	v265 = v252
	goto L69
L67:
	;
	v277 = v253
	v279 = v252
	goto L68
L68:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v284 = int32(-2139062144)
	if (int32(16843008)-v281|v281)&v284 != v284 {
		v306 = v277
		v307 = v281
		v308 = v279
		goto L62
	} else {
		goto L73
	}
L69:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v266)
	if v266 == int32(0) {
		goto L61
	} else {
		goto L71
	}
L70:
	;
	v277 = v273
	v279 = v271
	goto L68
L71:
	;
	v270 = int32(1)
	v271 = v265 + v270
	v273 = v263 + v270
	if v273&int32(3) != 0 {
		v263 = v273
		v265 = v271
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v289 = v277
	v290 = v281
	v291 = v279
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v290
	v293 = int32(4)
	v294 = v291 + v293
	v296 = v289 + v293
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v301 = int32(-2139062144)
	if (int32(16843008)-v298|v298)&v301 == v301 {
		v289 = v296
		v290 = v298
		v291 = v294
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v306 = v296
	v307 = v298
	v308 = v294
	goto L62
L76:
	;
	goto L75
L77:
	;
	v315 = v306
	v317 = v308
	goto L78
L78:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)) = uint8(v318)
	v320 = int32(1)
	if v318 != 0 {
		v315 = v315 + v320
		v317 = v317 + v320
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L61
L80:
	;
	goto L79
L81:
	;
	goto L24
L82:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[2])) = v363
	goto L83
L83:
	;
	v369 = F_format_elog_string(m, int32(_a_F_check_timezone_abbreviations_3), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[3])) = v369
	v374 = int32(0)
	goto L7
L85:
	;
	m.G0 = v21 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v374
	return base.B2i32(v374 != int32(0))
}
