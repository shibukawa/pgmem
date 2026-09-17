package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_emscripten_builtin_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var __phi65 int32
	_ = __phi65
	var v68 int32
	_ = v68
	var __phi68 int32
	_ = __phi68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var __phi230 int32
	_ = __phi230
	var v233 int32
	_ = v233
	var __phi233 int32
	_ = __phi233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = l0 - int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(4))))
	v18 = v16 & int32(-8)
	v19 = v13 + v18
	if v16&int32(1) != 0 {
		v139 = v18
		v142 = v13
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(v19) <= base.Ui32(v142) {
		goto L1
	} else {
		goto L44
	}
L4:
	;
	if v16&int32(2) == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v27 = v13 - v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[0]))
	if base.Ui32(v27) < base.Ui32(v29) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = v18 + v26
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[1]))
	if v33 != v27 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v49 == int32(0) {
		v139 = v31
		v142 = v27
		goto L3
	} else {
		goto L29
	}
L8:
	;
	v100 = int32(0)
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v38
	v139 = v31
	v142 = v27
	goto L3
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if base.Ui32(v26) <= base.Ui32(int32(255)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v82 = int32(3)
	if v81&v82 != v82 {
		v139 = v31
		v142 = v27
		goto L3
	} else {
		goto L28
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v35 != v38 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v35 != v27 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v40 = int32(_a_F_emscripten_builtin_free_0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[2])) = v42 & base.I32_rotl(int32(-2), int32(base.Ui32(v26)>>(uint(int32(3))%32)))
	v139 = v31
	v142 = v27
	goto L3
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v51
	v100 = v35
	goto L7
L18:
	;
	goto L19
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v62 = v54
	v63 = v27 + int32(20)
	goto L22
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v57 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	__phi65 = v62
	__phi68 = v63
	v65 = __phi65
	v68 = __phi68
	goto L24
L23:
	;
	v62 = v57
	v63 = v27 + int32(16)
	goto L22
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v75 != 0 {
		__phi65 = v75
		__phi68 = v65 + int32(20)
		v65 = __phi65
		v68 = __phi68
		goto L24
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(0)
	v100 = v65
	goto L7
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v78 != 0 {
		__phi65 = v78
		__phi68 = v65 + int32(16)
		v65 = __phi65
		v68 = __phi68
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[3])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v81 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v31 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v31
	return
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v111 = v109 << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_emscripten_builtin_free[4])))
	if v112 == v27 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+24)) = v49
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v131 != 0 {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_emscripten_builtin_free[4]))) = v100
	if v100 != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v27 == v124 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v117 = int32(_a_F_emscripten_builtin_free_1)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[5])) = v119 & base.I32_rotl(int32(-2), v109)
	v139 = v31
	v142 = v27
	goto L3
L35:
	;
	if v100 == int32(0) {
		v139 = v31
		v142 = v27
		goto L3
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v100
	goto L35
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v100
	goto L35
L39:
	;
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v131)+24)) = v100
	goto L42
L41:
	;
	goto L42
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v134 == int32(0) {
		v139 = v31
		v142 = v27
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = v100
	v139 = v31
	v142 = v27
	goto L3
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v149&int32(1) == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v149&int32(2) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	if base.Ui32(v313) <= base.Ui32(int32(255)) {
		goto L93
	} else {
		goto L94
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v197 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v197+v142))) = v197
	if v142 != v181 {
		v313 = v197
		goto L46
	} else {
		goto L92
	}
L48:
	;
	if v214 == int32(0) {
		goto L47
	} else {
		goto L77
	}
L49:
	;
	v257 = int32(0)
	goto L48
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[6]))
	if v159 == v19 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v149 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v139 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v139+v142))) = v139
	v313 = v139
	goto L46
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[6])) = v142
	v163 = int32(_a_F_emscripten_builtin_free_2)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[7]))
	v166 = v165 + v139
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[7])) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v166 | int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[1]))
	if v142 != v172 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[1]))
	if v181 == v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[3])) = v175
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[1])) = v175
	return
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[1])) = v142
	v185 = int32(_a_F_emscripten_builtin_free_3)
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[3]))
	v188 = v187 + v139
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[3])) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v188 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v188+v142))) = v188
	return
L58:
	;
	goto L59
L59:
	;
	v197 = v149&int32(-8) + v139
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if base.Ui32(v149) <= base.Ui32(int32(255)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v201 == v198 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v198 != v19 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v203 = int32(_a_F_emscripten_builtin_free_0)
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[2])) = v205 & base.I32_rotl(int32(-2), int32(base.Ui32(v149)>>(uint(int32(3))%32)))
	goto L47
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v201
	goto L47
L66:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v216
	v257 = v198
	goto L48
L67:
	;
	goto L68
L68:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v219 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v227 = v219
	v228 = v19 + int32(20)
	goto L71
L70:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v222 == int32(0) {
		goto L49
	} else {
		goto L72
	}
L71:
	;
	__phi230 = v227
	__phi233 = v228
	v230 = __phi230
	v233 = __phi233
	goto L73
L72:
	;
	v227 = v222
	v228 = v19 + int32(16)
	goto L71
L73:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v240 != 0 {
		__phi230 = v240
		__phi233 = v230 + int32(20)
		v230 = __phi230
		v233 = __phi233
		goto L73
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = int32(0)
	v257 = v230
	goto L48
L75:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v243 != 0 {
		__phi230 = v243
		__phi233 = v230 + int32(16)
		v230 = __phi230
		v233 = __phi233
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v268 = v266 << (uint(int32(2)) % 32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+uint32(_c_F_emscripten_builtin_free[4])))
	if v269 == v19 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+24)) = v214
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v288 != 0 {
		goto L88
	} else {
		goto L89
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+uint32(_c_F_emscripten_builtin_free[4]))) = v257
	if v257 != 0 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	if v19 == v281 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v274 = int32(_a_F_emscripten_builtin_free_1)
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[5])) = v276 & base.I32_rotl(int32(-2), v266)
	goto L47
L83:
	;
	if v257 == int32(0) {
		goto L47
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+16)) = v257
	goto L83
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+20)) = v257
	goto L83
L87:
	;
	goto L78
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+16)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v288)+24)) = v257
	goto L90
L89:
	;
	goto L90
L90:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v291 == int32(0) {
		goto L47
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+20)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v257
	goto L47
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[3])) = v197
	return
L93:
	;
	v325 = v313 & int32(248)
	v327 = v325 + int32(_a_F_emscripten_builtin_free_4)
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[2]))
	v333 = int32(1) << (uint(int32(base.Ui32(v313)>>(uint(int32(3))%32))) % 32)
	if v329&v333 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(v313) <= base.Ui32(int32(16777215)) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+uint32(_c_F_emscripten_builtin_free[8]))) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v341)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v341
	return
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[2])) = v333 | v329
	v341 = v327
	goto L96
L98:
	;
	goto L99
L99:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v325)+uint32(_c_F_emscripten_builtin_free[8])))
	v341 = v340
	goto L96
L100:
	;
	v352 = base.I32_clz(int32(base.Ui32(v313) >> (uint(int32(8)) % 32)))
	v355 = int32(1)
	v363 = int32(base.Ui32(v313)>>(uint(int32(38)-v352)%32))&v355 | v352<<(uint(v355)%32) ^ int32(62)
	goto L102
L101:
	;
	v363 = int32(31)
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+28)) = v363
	*(*int64)(unsafe.Add(mBase, uint32(v142)+16)) = int64(0)
	v368 = v363 << (uint(int32(2)) % 32)
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[5]))
	v374 = int32(1) << (uint(v363) % 32)
	if v372&v374 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436+v142))) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v434+v142))) = v443
	v449 = int32(_a_F_emscripten_builtin_free_5)
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[9]))
	v453 = v451 - int32(1)
	if v453 != 0 {
		goto L116
	} else {
		goto L117
	}
L104:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v428)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v398)+8)) = v142
	v434 = int32(24)
	v435 = v398
	v436 = int32(8)
	v438 = v428
	v443 = int32(0)
	goto L103
L105:
	;
	v434 = v427
	v435 = v142
	v436 = v420
	v438 = v422
	v443 = v142
	goto L103
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[5])) = v372 | v374
	*(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_emscripten_builtin_free[4]))) = v142
	v420 = int32(24)
	v422 = v368 + int32(_a_F_emscripten_builtin_free_6)
	v427 = int32(8)
	goto L105
L107:
	;
	goto L108
L108:
	;
	if v363 != int32(31) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v391 = int32(25) - int32(base.Ui32(v363)>>(uint(int32(1))%32))
	goto L111
L110:
	;
	v391 = int32(0)
	goto L111
L111:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_emscripten_builtin_free[4])))
	v396 = v313 << (uint(v391) % 32)
	v398 = v393
	goto L112
L112:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v403&int32(-8) == v313 {
		goto L104
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+16)) = v142
	v420 = int32(24)
	v422 = v398
	v427 = int32(8)
	goto L105
L114:
	;
	v413 = v398 + int32(base.Ui32(v396)>>(uint(int32(29))%32))&int32(4)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+16))
	if v414 != 0 {
		v396 = v396 << (uint(int32(1)) % 32)
		v398 = v414
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v455 = v453
	goto L118
L117:
	;
	v455 = int32(-1)
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_free[9])) = v455
	goto L1
}
func F_emscripten_builtin_memalign(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var __phi173 int32
	_ = __phi173
	var v175 int32
	_ = v175
	var __phi175 int32
	_ = __phi175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var __phi333 int32
	_ = __phi333
	var v335 int32
	_ = v335
	var __phi335 int32
	_ = __phi335
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var __phi638 int32
	_ = __phi638
	var v640 int32
	_ = v640
	var __phi640 int32
	_ = __phi640
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var __phi798 int32
	_ = __phi798
	var v800 int32
	_ = v800
	var __phi800 int32
	_ = __phi800
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v825 int32
	_ = v825
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v881 int32
	_ = v881
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v994 int32
	_ = v994
	var v1031 int32
	_ = v1031
	if base.Ui32(l0) <= base.Ui32(int32(8)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = F_emscripten_builtin_malloc(m, l1)
	mBase = m.M
	return v10
L2:
	;
	goto L3
L3:
	;
	v12 = int32(16)
	if base.Ui32(l0) <= base.Ui32(v12) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v1031
L5:
	;
	if base.Ui32(int32(-64)-v32) <= base.Ui32(l1) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v16 = v12
	goto L8
L7:
	;
	v16 = l0
	goto L8
L8:
	;
	if v16&(v16-int32(1)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v32 = v16
	goto L5
L10:
	;
	goto L11
L11:
	;
	v24 = v12
	goto L12
L12:
	;
	if base.Ui32(v24) < base.Ui32(v16) {
		v24 = v24 << (uint(int32(1)) % 32)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v32 = v24
	goto L5
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[0])) = int32(48)
	v1031 = int32(0)
	goto L4
L16:
	;
	goto L17
L17:
	;
	v48 = int32(11)
	if base.Ui32(l1) < base.Ui32(v48) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = int32(16)
	goto L20
L19:
	;
	v54 = (l1 + v48) & int32(-8)
	goto L20
L20:
	;
	v58 = F_emscripten_builtin_malloc(m, v54+v32+int32(12))
	mBase = m.M
	if v58 == int32(0) {
		v1031 = int32(0)
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v62 = v58 - int32(8)
	if (v32-int32(1))&v58 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v554)+4))
	if v560&int32(3) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L23:
	;
	v554 = v62
	goto L22
L24:
	;
	goto L25
L25:
	;
	v69 = v58 - int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v76 = int32(0)
	v80 = (v32+v58-int32(1))&(v76-v32) - int32(8)
	if base.Ui32(v80-v62) <= base.Ui32(int32(15)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v32
	goto L28
L27:
	;
	v85 = v76
	goto L28
L28:
	;
	v86 = v80 + v85
	v87 = v86 - v62
	v88 = v70&int32(-8) - v87
	if v70&int32(3) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v93 + v87
	v554 = v86
	goto L22
L30:
	;
	goto L31
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v98 = int32(1)
	v101 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v88 | v97&v98 | v101
	v104 = v86 + v88
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v105 | v98
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v87 | v109&v98 | v101
	v116 = v62 + v87
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v117 | v98
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v129&v98 != 0 {
		v246 = v62
		v247 = v87
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v554 = v86
	goto L22
L33:
	;
	goto L32
L34:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v255&int32(2) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L35:
	;
	if v129&int32(2) == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v137 = v136 + v87
	v138 = v62 - v136
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1]))
	if v138 != v140 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v156 == int32(0) {
		v246 = v138
		v247 = v137
		goto L34
	} else {
		goto L59
	}
L38:
	;
	v208 = int32(0)
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v145
	v246 = v138
	v247 = v137
	goto L34
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if base.Ui32(v136) <= base.Ui32(int32(255)) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v189 = int32(3)
	if v188&v189 != v189 {
		v246 = v138
		v247 = v137
		goto L34
	} else {
		goto L58
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v142 != v145 {
		goto L39
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	if v138 != v142 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v147 = int32(_a_F_emscripten_builtin_memalign_0)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2])) = v149 & base.I32_rotl(int32(-2), int32(base.Ui32(v136)>>(uint(int32(3))%32)))
	v246 = v138
	v247 = v137
	goto L34
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v158
	v208 = v142
	goto L37
L48:
	;
	goto L49
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	if v161 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v169 = v161
	v170 = v138 + int32(20)
	goto L52
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v164 == int32(0) {
		goto L38
	} else {
		goto L53
	}
L52:
	;
	__phi173 = v170
	__phi175 = v169
	v173 = __phi173
	v175 = __phi175
	goto L54
L53:
	;
	v169 = v164
	v170 = v138 + int32(16)
	goto L52
L54:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v182 != 0 {
		__phi173 = v175 + int32(20)
		__phi175 = v182
		v173 = __phi173
		v175 = __phi175
		goto L54
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = int32(0)
	v208 = v175
	goto L37
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	if v185 != 0 {
		__phi173 = v175 + int32(16)
		__phi175 = v185
		v173 = __phi173
		v175 = __phi175
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v188 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v137 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v137
	goto L32
L59:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v218 = v216 << (uint(int32(2)) % 32)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+uint32(_c_F_emscripten_builtin_memalign[4])))
	if v219 == v138 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+24)) = v156
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v238 != 0 {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+uint32(_c_F_emscripten_builtin_memalign[4]))) = v208
	if v208 != 0 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v138 == v231 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v224 = int32(_a_F_emscripten_builtin_memalign_1)
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5])) = v226 & base.I32_rotl(int32(-2), v216)
	v246 = v138
	v247 = v137
	goto L34
L65:
	;
	if v208 == int32(0) {
		v246 = v138
		v247 = v137
		goto L34
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v208
	goto L65
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+20)) = v208
	goto L65
L69:
	;
	goto L60
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+16)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v238)+24)) = v208
	goto L72
L71:
	;
	goto L72
L72:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	if v241 == int32(0) {
		v246 = v138
		v247 = v137
		goto L34
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+20)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v241)+24)) = v208
	v246 = v138
	v247 = v137
	goto L34
L74:
	;
	if base.Ui32(v416) <= base.Ui32(int32(255)) {
		goto L121
	} else {
		goto L122
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v299 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v246+v299))) = v299
	if v246 != v283 {
		v416 = v299
		goto L74
	} else {
		goto L120
	}
L76:
	;
	if v316 == int32(0) {
		goto L75
	} else {
		goto L105
	}
L77:
	;
	v360 = int32(0)
	goto L76
L78:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[6]))
	if v261 == v116 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v255 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v247 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v246+v247))) = v247
	v416 = v247
	goto L74
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[6])) = v246
	v265 = int32(_a_F_emscripten_builtin_memalign_2)
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[7]))
	v268 = v267 + v247
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[7])) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v268 | int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1]))
	if v246 != v274 {
		goto L33
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1]))
	if v283 == v116 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v277
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1])) = v277
	goto L32
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1])) = v246
	v287 = int32(_a_F_emscripten_builtin_memalign_3)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3]))
	v290 = v289 + v247
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v290 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v246+v290))) = v290
	goto L32
L86:
	;
	goto L87
L87:
	;
	v299 = v255&int32(-8) + v247
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	if base.Ui32(v255) <= base.Ui32(int32(255)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	if v303 == v300 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	if v300 != v116 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v305 = int32(_a_F_emscripten_builtin_memalign_0)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2])) = v307 & base.I32_rotl(int32(-2), int32(base.Ui32(v255)>>(uint(int32(3))%32)))
	goto L75
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+12)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = v303
	goto L75
L94:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = v318
	v360 = v300
	goto L76
L95:
	;
	goto L96
L96:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	if v321 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v329 = v321
	v330 = v116 + int32(20)
	goto L99
L98:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	if v324 == int32(0) {
		goto L77
	} else {
		goto L100
	}
L99:
	;
	__phi333 = v330
	__phi335 = v329
	v333 = __phi333
	v335 = __phi335
	goto L101
L100:
	;
	v329 = v324
	v330 = v116 + int32(16)
	goto L99
L101:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v335)+20))
	if v342 != 0 {
		__phi333 = v335 + int32(20)
		__phi335 = v342
		v333 = __phi333
		v335 = __phi335
		goto L101
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = int32(0)
	v360 = v335
	goto L76
L103:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v335)+16))
	if v345 != 0 {
		__phi333 = v335 + int32(16)
		__phi335 = v345
		v333 = __phi333
		v335 = __phi335
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v116)+28))
	v370 = v368 << (uint(int32(2)) % 32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+uint32(_c_F_emscripten_builtin_memalign[4])))
	if v371 == v116 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+24)) = v316
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	if v390 != 0 {
		goto L116
	} else {
		goto L117
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+uint32(_c_F_emscripten_builtin_memalign[4]))) = v360
	if v360 != 0 {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v316)+16))
	if v116 == v383 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v376 = int32(_a_F_emscripten_builtin_memalign_1)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5])) = v378 & base.I32_rotl(int32(-2), v368)
	goto L75
L111:
	;
	if v360 == int32(0) {
		goto L75
	} else {
		goto L115
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+16)) = v360
	goto L111
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+20)) = v360
	goto L111
L115:
	;
	goto L106
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v390)+24)) = v360
	goto L118
L117:
	;
	goto L118
L118:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	if v393 == int32(0) {
		goto L75
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+20)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v393)+24)) = v360
	goto L75
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v299
	goto L32
L121:
	;
	v427 = v416 & int32(248)
	v429 = v427 + int32(_a_F_emscripten_builtin_memalign_4)
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2]))
	v435 = int32(1) << (uint(int32(base.Ui32(v416)>>(uint(int32(3))%32))) % 32)
	if v431&v435 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	goto L123
L123:
	;
	if base.Ui32(v416) <= base.Ui32(int32(16777215)) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427)+uint32(_c_F_emscripten_builtin_memalign[8]))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v443)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v443
	goto L32
L125:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2])) = v435 | v431
	v443 = v429
	goto L124
L126:
	;
	goto L127
L127:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v427)+uint32(_c_F_emscripten_builtin_memalign[8])))
	v443 = v442
	goto L124
L128:
	;
	v454 = base.I32_clz(int32(base.Ui32(v416) >> (uint(int32(8)) % 32)))
	v457 = int32(1)
	v465 = int32(base.Ui32(v416)>>(uint(int32(38)-v454)%32))&v457 | v454<<(uint(v457)%32) ^ int32(62)
	goto L130
L129:
	;
	v465 = int32(31)
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = v465
	*(*int64)(unsafe.Add(mBase, uint32(v246)+16)) = int64(0)
	v470 = v465 << (uint(int32(2)) % 32)
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5]))
	v476 = int32(1) << (uint(v465) % 32)
	if v474&v476 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v497)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v529)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v497)+8)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v529
	goto L33
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v246
	goto L32
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5])) = v474 | v476
	*(*int32)(unsafe.Add(mBase, uint32(v470)+uint32(_c_F_emscripten_builtin_memalign[4]))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v470 + int32(_a_F_emscripten_builtin_memalign_5)
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v465 != int32(31) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v492 = int32(25) - int32(base.Ui32(v465)>>(uint(int32(1))%32))
	goto L138
L137:
	;
	v492 = int32(0)
	goto L138
L138:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v470)+uint32(_c_F_emscripten_builtin_memalign[4])))
	v497 = v494
	v498 = v416 << (uint(v492) % 32)
	goto L139
L139:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v504&int32(-8) == v416 {
		goto L131
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+16)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v497
	goto L132
L141:
	;
	v514 = v497 + int32(base.Ui32(v498)>>(uint(int32(29))%32))&int32(4)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+16))
	if v515 != 0 {
		v497 = v515
		v498 = v498 << (uint(int32(1)) % 32)
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v1031 = v554 + int32(8)
	goto L4
L144:
	;
	v566 = v560 & int32(-8)
	if base.Ui32(v566) <= base.Ui32(v54+int32(16)) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v570 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v554)+4)) = v54 | v560&v570 | int32(2)
	v576 = v554 + v54
	v577 = v566 - v54
	*(*int32)(unsafe.Add(mBase, uint32(v576)+4)) = v577 | int32(3)
	v581 = v554 + v566
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+4)) = v582 | v570
	v593 = v576 + v577
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v594&v570 != 0 {
		v711 = v576
		v712 = v577
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L143
L147:
	;
	goto L146
L148:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	if v720&int32(2) == int32(0) {
		goto L192
	} else {
		goto L193
	}
L149:
	;
	if v594&int32(2) == int32(0) {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v602 = v601 + v577
	v603 = v576 - v601
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1]))
	if v603 != v605 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	if v621 == int32(0) {
		v711 = v603
		v712 = v602
		goto L148
	} else {
		goto L173
	}
L152:
	;
	v673 = int32(0)
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+12)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(v607)+8)) = v610
	v711 = v603
	v712 = v602
	goto L148
L154:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	if base.Ui32(v601) <= base.Ui32(int32(255)) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	v654 = int32(3)
	if v653&v654 != v654 {
		v711 = v603
		v712 = v602
		goto L148
	} else {
		goto L172
	}
L157:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v603)+8))
	if v607 != v610 {
		goto L153
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v603)+24))
	if v603 != v607 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v612 = int32(_a_F_emscripten_builtin_memalign_0)
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2])) = v614 & base.I32_rotl(int32(-2), int32(base.Ui32(v601)>>(uint(int32(3))%32)))
	v711 = v603
	v712 = v602
	goto L148
L161:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v603)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+12)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(v607)+8)) = v623
	v673 = v607
	goto L151
L162:
	;
	goto L163
L163:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v626 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v634 = v626
	v635 = v603 + int32(20)
	goto L166
L165:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v603)+16))
	if v629 == int32(0) {
		goto L152
	} else {
		goto L167
	}
L166:
	;
	__phi638 = v635
	__phi640 = v634
	v638 = __phi638
	v640 = __phi640
	goto L168
L167:
	;
	v634 = v629
	v635 = v603 + int32(16)
	goto L166
L168:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v640)+20))
	if v647 != 0 {
		__phi638 = v640 + int32(20)
		__phi640 = v647
		v638 = __phi638
		v640 = __phi640
		goto L168
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = int32(0)
	v673 = v640
	goto L151
L170:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v640)+16))
	if v650 != 0 {
		__phi638 = v640 + int32(16)
		__phi640 = v650
		v638 = __phi638
		v640 = __phi640
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v593)+4)) = v653 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v602 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v602
	goto L146
L173:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v683 = v681 << (uint(int32(2)) % 32)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+uint32(_c_F_emscripten_builtin_memalign[4])))
	if v684 == v603 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+24)) = v621
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v603)+16))
	if v703 != 0 {
		goto L184
	} else {
		goto L185
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683)+uint32(_c_F_emscripten_builtin_memalign[4]))) = v673
	if v673 != 0 {
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v621)+16))
	if v603 == v696 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v689 = int32(_a_F_emscripten_builtin_memalign_1)
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5])) = v691 & base.I32_rotl(int32(-2), v681)
	v711 = v603
	v712 = v602
	goto L148
L179:
	;
	if v673 == int32(0) {
		v711 = v603
		v712 = v602
		goto L148
	} else {
		goto L183
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621)+16)) = v673
	goto L179
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621)+20)) = v673
	goto L179
L183:
	;
	goto L174
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+16)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v703)+24)) = v673
	goto L186
L185:
	;
	goto L186
L186:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v706 == int32(0) {
		v711 = v603
		v712 = v602
		goto L148
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+20)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v706)+24)) = v673
	v711 = v603
	v712 = v602
	goto L148
L188:
	;
	if base.Ui32(v881) <= base.Ui32(int32(255)) {
		goto L235
	} else {
		goto L236
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711)+4)) = v764 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v711+v764))) = v764
	if v711 != v748 {
		v881 = v764
		goto L188
	} else {
		goto L234
	}
L190:
	;
	if v781 == int32(0) {
		goto L189
	} else {
		goto L219
	}
L191:
	;
	v825 = int32(0)
	goto L190
L192:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[6]))
	if v726 == v593 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v593)+4)) = v720 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v711)+4)) = v712 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v711+v712))) = v712
	v881 = v712
	goto L188
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[6])) = v711
	v730 = int32(_a_F_emscripten_builtin_memalign_2)
	v732 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[7]))
	v733 = v732 + v712
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[7])) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v711)+4)) = v733 | int32(1)
	v739 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1]))
	if v711 != v739 {
		goto L147
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1]))
	if v748 == v593 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v742 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v742
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1])) = v742
	goto L146
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[1])) = v711
	v752 = int32(_a_F_emscripten_builtin_memalign_3)
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3]))
	v755 = v754 + v712
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v755
	*(*int32)(unsafe.Add(mBase, uint32(v711)+4)) = v755 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v711+v755))) = v755
	goto L146
L200:
	;
	goto L201
L201:
	;
	v764 = v720&int32(-8) + v712
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	if base.Ui32(v720) <= base.Ui32(int32(255)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	if v768 == v765 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v593)+24))
	if v765 != v593 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v770 = int32(_a_F_emscripten_builtin_memalign_0)
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2])) = v772 & base.I32_rotl(int32(-2), int32(base.Ui32(v720)>>(uint(int32(3))%32)))
	goto L189
L206:
	;
	goto L207
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v768)+12)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = v768
	goto L189
L208:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v783)+12)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = v783
	v825 = v765
	goto L190
L209:
	;
	goto L210
L210:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v593)+20))
	if v786 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v794 = v786
	v795 = v593 + int32(20)
	goto L213
L212:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v593)+16))
	if v789 == int32(0) {
		goto L191
	} else {
		goto L214
	}
L213:
	;
	__phi798 = v795
	__phi800 = v794
	v798 = __phi798
	v800 = __phi800
	goto L215
L214:
	;
	v794 = v789
	v795 = v593 + int32(16)
	goto L213
L215:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v800)+20))
	if v807 != 0 {
		__phi798 = v800 + int32(20)
		__phi800 = v807
		v798 = __phi798
		v800 = __phi800
		goto L215
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v798))) = int32(0)
	v825 = v800
	goto L190
L217:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v800)+16))
	if v810 != 0 {
		__phi798 = v800 + int32(16)
		__phi800 = v810
		v798 = __phi798
		v800 = __phi800
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v593)+28))
	v835 = v833 << (uint(int32(2)) % 32)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+uint32(_c_F_emscripten_builtin_memalign[4])))
	if v836 == v593 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v825)+24)) = v781
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v593)+16))
	if v855 != 0 {
		goto L230
	} else {
		goto L231
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v835)+uint32(_c_F_emscripten_builtin_memalign[4]))) = v825
	if v825 != 0 {
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v781)+16))
	if v593 == v848 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v841 = int32(_a_F_emscripten_builtin_memalign_1)
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5])) = v843 & base.I32_rotl(int32(-2), v833)
	goto L189
L225:
	;
	if v825 == int32(0) {
		goto L189
	} else {
		goto L229
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v781)+16)) = v825
	goto L225
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v781)+20)) = v825
	goto L225
L229:
	;
	goto L220
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v825)+16)) = v855
	*(*int32)(unsafe.Add(mBase, uint32(v855)+24)) = v825
	goto L232
L231:
	;
	goto L232
L232:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v593)+20))
	if v858 == int32(0) {
		goto L189
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v825)+20)) = v858
	*(*int32)(unsafe.Add(mBase, uint32(v858)+24)) = v825
	goto L189
L234:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[3])) = v764
	goto L146
L235:
	;
	v892 = v881 & int32(248)
	v894 = v892 + int32(_a_F_emscripten_builtin_memalign_4)
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2]))
	v900 = int32(1) << (uint(int32(base.Ui32(v881)>>(uint(int32(3))%32))) % 32)
	if v896&v900 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	goto L237
L237:
	;
	if base.Ui32(v881) <= base.Ui32(int32(16777215)) {
		goto L242
	} else {
		goto L243
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892)+uint32(_c_F_emscripten_builtin_memalign[8]))) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v908)+12)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v711)+12)) = v894
	*(*int32)(unsafe.Add(mBase, uint32(v711)+8)) = v908
	goto L146
L239:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[2])) = v900 | v896
	v908 = v894
	goto L238
L240:
	;
	goto L241
L241:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v892)+uint32(_c_F_emscripten_builtin_memalign[8])))
	v908 = v907
	goto L238
L242:
	;
	v919 = base.I32_clz(int32(base.Ui32(v881) >> (uint(int32(8)) % 32)))
	v922 = int32(1)
	v930 = int32(base.Ui32(v881)>>(uint(int32(38)-v919)%32))&v922 | v919<<(uint(v922)%32) ^ int32(62)
	goto L244
L243:
	;
	v930 = int32(31)
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711)+28)) = v930
	*(*int64)(unsafe.Add(mBase, uint32(v711)+16)) = int64(0)
	v935 = v930 << (uint(int32(2)) % 32)
	v939 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5]))
	v941 = int32(1) << (uint(v930) % 32)
	if v939&v941 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v962)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v994)+12)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v962)+8)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v711)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v711)+12)) = v962
	*(*int32)(unsafe.Add(mBase, uint32(v711)+8)) = v994
	goto L147
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711)+12)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v711)+8)) = v711
	goto L146
L247:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_memalign[5])) = v939 | v941
	*(*int32)(unsafe.Add(mBase, uint32(v935)+uint32(_c_F_emscripten_builtin_memalign[4]))) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v711)+24)) = v935 + int32(_a_F_emscripten_builtin_memalign_5)
	goto L246
L248:
	;
	goto L249
L249:
	;
	if v930 != int32(31) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v957 = int32(25) - int32(base.Ui32(v930)>>(uint(int32(1))%32))
	goto L252
L251:
	;
	v957 = int32(0)
	goto L252
L252:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v935)+uint32(_c_F_emscripten_builtin_memalign[4])))
	v962 = v959
	v963 = v881 << (uint(v957) % 32)
	goto L253
L253:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	if v969&int32(-8) == v881 {
		goto L245
	} else {
		goto L255
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+16)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v711)+24)) = v962
	goto L246
L255:
	;
	v979 = v962 + int32(base.Ui32(v963)>>(uint(int32(29))%32))&int32(4)
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v979)+16))
	if v980 != 0 {
		v962 = v980
		v963 = v963 << (uint(int32(1)) % 32)
		goto L253
	} else {
		goto L256
	}
L256:
	;
	goto L254
}
