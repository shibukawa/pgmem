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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
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
	v37 = int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v37
	v41 = F_palloc_mul(m, int32(24), v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v41
	v44 = int32(0)
	v50 = F_ParseTzFile(m, v13, v44, v21+int32(12), v21+int32(8), v44)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[0])) = v34
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L85
	}
L8:
	;
	if v50 < int32(0) {
		v375 = v18
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v58 = v50<<(uint(int32(4))%32) | int32(8)
	if int32(0) < v50 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v63 = v58
	v64 = v4
	goto L13
L11:
	;
	v89 = v58
	goto L12
L12:
	;
	v99 = F_guc_malloc(m, v89)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L20
	}
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v54+v64*int32(24))+4))
	if v76 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v89 = v83
	goto L12
L15:
	;
	v77 = F_strlen(m, v76)
	mBase = m.M
	v83 = (v77+int32(12))&int32(-8) + v63
	goto L17
L16:
	;
	v83 = v63
	goto L17
L17:
	;
	v85 = v64 + int32(1)
	if v85 != v50 {
		v63 = v83
		v64 = v85
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v99 != 0 {
		v375 = v99
		goto L7
	} else {
		goto L82
	}
L20:
	;
	if v99 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v89
	if v50 <= int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v113 = int32(0)
	v115 = v58
	goto L23
L23:
	;
	v124 = v99 + int32(8) + v113<<(uint(int32(4))%32)
	v127 = v54 + v113*int32(24)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	goto L28
L24:
	;
	goto L19
L25:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v248 != 0 {
		goto L57
	} else {
		goto L58
	}
L26:
	;
	v245 = F_strlen(m, v234)
	mBase = m.M
	goto L25
L28:
	;
	goto L29
L29:
	;
	v135 = int32(10)
	if (v124^v128)&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v238)
	goto L26
L31:
	;
	v219 = v214
	v220 = v215
	v221 = v216
	goto L52
L32:
	;
	if v209 == int32(0) {
		v234 = v207
		v235 = v208
		goto L30
	} else {
		goto L51
	}
L33:
	;
	v207 = v128
	v208 = v124
	v209 = v135
	goto L32
L34:
	;
	goto L35
L35:
	;
	v139 = int32(0)
	if base.B2i32(v128&int32(3) == v139)|int32(0) == v139 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v175 == int32(0) {
		v234 = v172
		v235 = v173
		goto L30
	} else {
		goto L45
	}
L37:
	;
	v151 = v128
	v152 = v124
	v153 = v135
	goto L40
L38:
	;
	goto L39
L39:
	;
	v172 = v128
	v173 = v124
	v174 = v135
	v175 = int32(1)
	goto L36
L40:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v155)
	if v155 == int32(0) {
		v214 = v151
		v215 = v152
		v216 = v153
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v172 = v166
	v173 = v160
	v174 = v162
	v175 = v164
	goto L36
L42:
	;
	v159 = int32(1)
	v160 = v152 + v159
	v162 = v153 - v159
	v163 = int32(0)
	v164 = base.B2i32(v162 != v163)
	v166 = v151 + v159
	if v166&int32(3) == v163 {
		v172 = v166
		v173 = v160
		v174 = v162
		v175 = v164
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v162 != 0 {
		v151 = v166
		v152 = v160
		v153 = v162
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if base.B2i32(v178 == int32(0))|base.B2i32(base.Ui32(v174) < base.Ui32(int32(4))) != 0 {
		v207 = v172
		v208 = v173
		v209 = v174
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v185 = v172
	v186 = v173
	v187 = v174
	goto L47
L47:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v193 = int32(-2139062144)
	if (int32(16843008)-v190|v190)&v193 != v193 {
		v214 = v185
		v215 = v186
		v216 = v187
		goto L31
	} else {
		goto L49
	}
L48:
	;
	v207 = v201
	v208 = v199
	v209 = v203
	goto L32
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v190
	v198 = int32(4)
	v199 = v186 + v198
	v201 = v185 + v198
	v203 = v187 - v198
	if base.Ui32(int32(3)) < base.Ui32(v203) {
		v185 = v201
		v186 = v199
		v187 = v203
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v214 = v207
	v215 = v208
	v216 = v209
	goto L31
L52:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v220))) = uint8(v223)
	if v223 == int32(0) {
		v234 = v219
		v235 = v220
		goto L30
	} else {
		goto L54
	}
L53:
	;
	v234 = v230
	v235 = v228
	goto L30
L54:
	;
	v227 = int32(1)
	v228 = v220 + v227
	v230 = v219 + v227
	v232 = v221 - v227
	if v232 != 0 {
		v219 = v230
		v220 = v228
		v221 = v232
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v348 = v113 + int32(1)
	if v348 != v50 {
		v113 = v348
		v115 = v345
		goto L23
	} else {
		goto L81
	}
L57:
	;
	v249 = v99 + v115
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = int32(0)
	v253 = v249 + int32(4)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if (v254^v253)&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	goto L59
L59:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+12)))
	v341 = v339 + int32(5)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+11)) = uint8(v341)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v343
	v345 = v115
	goto L56
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v115
	v330 = int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+11)) = uint8(v330)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v333 = F_strlen(m, v332)
	mBase = m.M
	v345 = (v333+int32(12))&int32(-8) + v115
	goto L56
L61:
	;
	goto L60
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v308)
	if v308&int32(255) == int32(0) {
		goto L61
	} else {
		goto L77
	}
L63:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	v307 = v254
	v308 = v260
	v309 = v253
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v254&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v264 = v254
	v266 = v253
	goto L69
L67:
	;
	v278 = v254
	v280 = v253
	goto L68
L68:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v285 = int32(-2139062144)
	if (int32(16843008)-v282|v282)&v285 != v285 {
		v307 = v278
		v308 = v282
		v309 = v280
		goto L62
	} else {
		goto L73
	}
L69:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	*(*uint8)(unsafe.Add(mBase, uint32(v266))) = uint8(v267)
	if v267 == int32(0) {
		goto L61
	} else {
		goto L71
	}
L70:
	;
	v278 = v274
	v280 = v272
	goto L68
L71:
	;
	v271 = int32(1)
	v272 = v266 + v271
	v274 = v264 + v271
	if v274&int32(3) != 0 {
		v264 = v274
		v266 = v272
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v290 = v278
	v291 = v282
	v292 = v280
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291
	v294 = int32(4)
	v295 = v292 + v294
	v297 = v290 + v294
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v302 = int32(-2139062144)
	if (int32(16843008)-v299|v299)&v302 == v302 {
		v290 = v297
		v291 = v299
		v292 = v295
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v307 = v297
	v308 = v299
	v309 = v295
	goto L62
L76:
	;
	goto L75
L77:
	;
	v316 = v307
	v318 = v309
	goto L78
L78:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)) = uint8(v319)
	v321 = int32(1)
	if v319 != 0 {
		v316 = v316 + v321
		v318 = v318 + v321
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
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[2])) = v364
	goto L83
L83:
	;
	v370 = F_format_elog_string(m, int32(_a_F_check_timezone_abbreviations_3), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone_abbreviations[3])) = v370
	v375 = int32(0)
	goto L7
L85:
	;
	m.G0 = v21 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v375
	return base.B2i32(v375 != int32(0))
}
