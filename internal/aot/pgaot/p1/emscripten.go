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
	var v66 int32
	_ = v66
	var __phi66 int32
	_ = __phi66
	var v69 int32
	_ = v69
	var __phi69 int32
	_ = __phi69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
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
	var v231 int32
	_ = v231
	var __phi231 int32
	_ = __phi231
	var v234 int32
	_ = v234
	var __phi234 int32
	_ = __phi234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
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
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
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
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1503]))
	if base.Ui32(v27) < base.Ui32(v29) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = v18 + v26
	v33 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
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
	v99 = int32(0)
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
	v40 = int32(4747256)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v42 & base.I32_rotl(int32(-2), int32(base.Ui32(v26)>>(uint(int32(3))%32)))
	v139 = v31
	v142 = v27
	goto L3
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v51
	v99 = v35
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
	__phi66 = v62
	__phi69 = v63
	v66 = __phi66
	v69 = __phi69
	goto L24
L23:
	;
	v62 = v57
	v63 = v27 + int32(16)
	goto L22
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v75 != 0 {
		__phi66 = v75
		__phi69 = v66 + int32(20)
		v66 = __phi66
		v69 = __phi69
		goto L24
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(0)
	v99 = v66
	goto L7
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v78 != 0 {
		__phi66 = v78
		__phi69 = v66 + int32(16)
		v66 = __phi66
		v69 = __phi69
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v81 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v31 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v31
	return
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v111 = v109 << (uint(int32(2)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[1507])))
	if v114 == v27 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = v49
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v131 != 0 {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[1507]))) = v99
	if v99 != 0 {
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
	v117 = int32(4747260)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v119 & base.I32_rotl(int32(-2), v109)
	v139 = v31
	v142 = v27
	goto L3
L35:
	;
	if v99 == int32(0) {
		v139 = v31
		v142 = v27
		goto L3
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v99
	goto L35
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v99
	goto L35
L39:
	;
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v131)+24)) = v99
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
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = v99
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
	v256 = int32(0)
	goto L48
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[1509]))
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
	*(*int32)(unsafe.Add(mBase, _consts[1509])) = v142
	v163 = int32(4747268)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[1510]))
	v166 = v165 + v139
	*(*int32)(unsafe.Add(mBase, _consts[1510])) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v166 | int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
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
	v181 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
	if v181 == v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v175
	*(*int32)(unsafe.Add(mBase, _consts[1504])) = v175
	return
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1504])) = v142
	v185 = int32(4747264)
	v187 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	v188 = v187 + v139
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v188
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
	v203 = int32(4747256)
	v205 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v205 & base.I32_rotl(int32(-2), int32(base.Ui32(v149)>>(uint(int32(3))%32)))
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
	v256 = v198
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
	__phi231 = v227
	__phi234 = v228
	v231 = __phi231
	v234 = __phi234
	goto L73
L72:
	;
	v227 = v222
	v228 = v19 + int32(16)
	goto L71
L73:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	if v240 != 0 {
		__phi231 = v240
		__phi234 = v231 + int32(20)
		v231 = __phi231
		v234 = __phi234
		goto L73
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(0)
	v256 = v231
	goto L48
L75:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if v243 != 0 {
		__phi231 = v243
		__phi234 = v231 + int32(16)
		v231 = __phi231
		v234 = __phi234
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
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v268)+uint32(_consts[1507])))
	if v271 == v19 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+24)) = v214
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v288 != 0 {
		goto L88
	} else {
		goto L89
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+uint32(_consts[1507]))) = v256
	if v256 != 0 {
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
	v274 = int32(4747260)
	v276 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v276 & base.I32_rotl(int32(-2), v266)
	goto L47
L83:
	;
	if v256 == int32(0) {
		goto L47
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+16)) = v256
	goto L83
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+20)) = v256
	goto L83
L87:
	;
	goto L78
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+16)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v288)+24)) = v256
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
	*(*int32)(unsafe.Add(mBase, uint32(v256)+20)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v256
	goto L47
L92:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v197
	return
L93:
	;
	v325 = v313 & int32(-8)
	v327 = v325 + int32(4747296)
	v329 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v325)+uint32(_consts[1511]))) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v341)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v341
	return
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v333 | v329
	v341 = v327
	goto L96
L98:
	;
	goto L99
L99:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v325)+uint32(_consts[1511])))
	v341 = v340
	goto L96
L100:
	;
	v352 = base.I32_clz(int32(base.Ui32(v313) >> (uint(int32(8)) % 32)))
	v355 = int32(1)
	v362 = int32(base.Ui32(v313)>>(uint(int32(38)-v352)%32))&v355 - v352<<(uint(v355)%32) + int32(62)
	goto L102
L101:
	;
	v362 = int32(31)
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+28)) = v362
	*(*int64)(unsafe.Add(mBase, uint32(v142)+16)) = int64(0)
	v367 = v362 << (uint(int32(2)) % 32)
	v371 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	v373 = int32(1) << (uint(v362) % 32)
	if v371&v373 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434+v142))) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v433+v142))) = v442
	v448 = int32(4747288)
	v450 = *(*int32)(unsafe.Add(mBase, _consts[1512]))
	v452 = v450 - int32(1)
	if v452 != 0 {
		goto L116
	} else {
		goto L117
	}
L104:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v427)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v398)+8)) = v142
	v433 = int32(24)
	v434 = int32(8)
	v435 = v398
	v438 = v427
	v442 = int32(0)
	goto L103
L105:
	;
	v433 = v426
	v434 = v418
	v435 = v142
	v438 = v422
	v442 = v142
	goto L103
L106:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v371 | v373
	*(*int32)(unsafe.Add(mBase, uint32(v367)+uint32(_consts[1507]))) = v142
	v418 = int32(24)
	v422 = v367 + int32(4747560)
	v426 = int32(8)
	goto L105
L107:
	;
	goto L108
L108:
	;
	if v362 != int32(31) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v390 = int32(25) - int32(base.Ui32(v362)>>(uint(int32(1))%32))
	goto L111
L110:
	;
	v390 = int32(0)
	goto L111
L111:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v367)+uint32(_consts[1507])))
	v394 = v313 << (uint(v390) % 32)
	v398 = v392
	goto L112
L112:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v402&int32(-8) == v313 {
		goto L104
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v412)+16)) = v142
	v418 = int32(24)
	v422 = v398
	v426 = int32(8)
	goto L105
L114:
	;
	v412 = v398 + int32(base.Ui32(v394)>>(uint(int32(29))%32))&int32(4)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+16))
	if v413 != 0 {
		v394 = v394 << (uint(int32(1)) % 32)
		v398 = v413
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v454 = v452
	goto L118
L117:
	;
	v454 = int32(-1)
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1512])) = v454
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
	var v174 int32
	_ = v174
	var __phi174 int32
	_ = __phi174
	var v176 int32
	_ = v176
	var __phi176 int32
	_ = __phi176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
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
	var v334 int32
	_ = v334
	var __phi334 int32
	_ = __phi334
	var v336 int32
	_ = v336
	var __phi336 int32
	_ = __phi336
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
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
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v528 int32
	_ = v528
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var __phi638 int32
	_ = __phi638
	var v640 int32
	_ = v640
	var __phi640 int32
	_ = __phi640
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var __phi798 int32
	_ = __phi798
	var v800 int32
	_ = v800
	var __phi800 int32
	_ = __phi800
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v823 int32
	_ = v823
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v992 int32
	_ = v992
	var v1029 int32
	_ = v1029
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
	return v1029
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
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(48)
	v1029 = int32(0)
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
		v1029 = int32(0)
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
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	if v559&int32(3) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L23:
	;
	v553 = v62
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
	v553 = v86
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
	v553 = v86
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
	v140 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
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
	v207 = int32(0)
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
	v147 = int32(4747256)
	v149 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v149 & base.I32_rotl(int32(-2), int32(base.Ui32(v136)>>(uint(int32(3))%32)))
	v246 = v138
	v247 = v137
	goto L34
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v158
	v207 = v142
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
	__phi174 = v169
	__phi176 = v170
	v174 = __phi174
	v176 = __phi176
	goto L54
L53:
	;
	v169 = v164
	v170 = v138 + int32(16)
	goto L52
L54:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	if v182 != 0 {
		__phi174 = v182
		__phi176 = v174 + int32(20)
		v174 = __phi174
		v176 = __phi176
		goto L54
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = int32(0)
	v207 = v174
	goto L37
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	if v185 != 0 {
		__phi174 = v185
		__phi176 = v174 + int32(16)
		v174 = __phi174
		v176 = __phi176
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v188 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v137 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v137
	goto L32
L59:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v218 = v216 << (uint(int32(2)) % 32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v218)+uint32(_consts[1507])))
	if v221 == v138 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v156
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v238 != 0 {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+uint32(_consts[1507]))) = v207
	if v207 != 0 {
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
	v224 = int32(4747260)
	v226 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v226 & base.I32_rotl(int32(-2), v216)
	v246 = v138
	v247 = v137
	goto L34
L65:
	;
	if v207 == int32(0) {
		v246 = v138
		v247 = v137
		goto L34
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v207
	goto L65
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+20)) = v207
	goto L65
L69:
	;
	goto L60
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+16)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v238)+24)) = v207
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
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v241)+24)) = v207
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
	v359 = int32(0)
	goto L76
L78:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[1509]))
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
	*(*int32)(unsafe.Add(mBase, _consts[1509])) = v246
	v265 = int32(4747268)
	v267 = *(*int32)(unsafe.Add(mBase, _consts[1510]))
	v268 = v267 + v247
	*(*int32)(unsafe.Add(mBase, _consts[1510])) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v268 | int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
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
	v283 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
	if v283 == v116 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v277
	*(*int32)(unsafe.Add(mBase, _consts[1504])) = v277
	goto L32
L85:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1504])) = v246
	v287 = int32(4747264)
	v289 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	v290 = v289 + v247
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v290
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
	v305 = int32(4747256)
	v307 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v307 & base.I32_rotl(int32(-2), int32(base.Ui32(v255)>>(uint(int32(3))%32)))
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
	v359 = v300
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
	__phi334 = v329
	__phi336 = v330
	v334 = __phi334
	v336 = __phi336
	goto L101
L100:
	;
	v329 = v324
	v330 = v116 + int32(16)
	goto L99
L101:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v334)+20))
	if v342 != 0 {
		__phi334 = v342
		__phi336 = v334 + int32(20)
		v334 = __phi334
		v336 = __phi336
		goto L101
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336))) = int32(0)
	v359 = v334
	goto L76
L103:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v334)+16))
	if v345 != 0 {
		__phi334 = v345
		__phi336 = v334 + int32(16)
		v334 = __phi334
		v336 = __phi336
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
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v370)+uint32(_consts[1507])))
	if v373 == v116 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+24)) = v316
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	if v390 != 0 {
		goto L116
	} else {
		goto L117
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+uint32(_consts[1507]))) = v359
	if v359 != 0 {
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
	v376 = int32(4747260)
	v378 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v378 & base.I32_rotl(int32(-2), v368)
	goto L75
L111:
	;
	if v359 == int32(0) {
		goto L75
	} else {
		goto L115
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+16)) = v359
	goto L111
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+20)) = v359
	goto L111
L115:
	;
	goto L106
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+16)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v390)+24)) = v359
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
	*(*int32)(unsafe.Add(mBase, uint32(v359)+20)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v393)+24)) = v359
	goto L75
L120:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v299
	goto L32
L121:
	;
	v427 = v416 & int32(-8)
	v429 = v427 + int32(4747296)
	v431 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v427)+uint32(_consts[1511]))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v443)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v443
	goto L32
L125:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v435 | v431
	v443 = v429
	goto L124
L126:
	;
	goto L127
L127:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v427)+uint32(_consts[1511])))
	v443 = v442
	goto L124
L128:
	;
	v454 = base.I32_clz(int32(base.Ui32(v416) >> (uint(int32(8)) % 32)))
	v457 = int32(1)
	v464 = int32(base.Ui32(v416)>>(uint(int32(38)-v454)%32))&v457 - v454<<(uint(v457)%32) + int32(62)
	goto L130
L129:
	;
	v464 = int32(31)
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = v464
	*(*int64)(unsafe.Add(mBase, uint32(v246)+16)) = int64(0)
	v469 = v464 << (uint(int32(2)) % 32)
	v473 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	v475 = int32(1) << (uint(v464) % 32)
	if v473&v475 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v499)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v499)+8)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v528
	goto L33
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v246
	goto L32
L133:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v475 | v473
	*(*int32)(unsafe.Add(mBase, uint32(v469)+uint32(_consts[1507]))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v469 + int32(4747560)
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v464 != int32(31) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v491 = int32(25) - int32(base.Ui32(v464)>>(uint(int32(1))%32))
	goto L138
L137:
	;
	v491 = int32(0)
	goto L138
L138:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v469)+uint32(_consts[1507])))
	v496 = v416 << (uint(v491) % 32)
	v499 = v493
	goto L139
L139:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v503&int32(-8) == v416 {
		goto L131
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+16)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v499
	goto L132
L141:
	;
	v513 = v499 + int32(base.Ui32(v496)>>(uint(int32(29))%32))&int32(4)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	if v514 != 0 {
		v496 = v496 << (uint(int32(1)) % 32)
		v499 = v514
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v1029 = v553 + int32(8)
	goto L4
L144:
	;
	v565 = v559 & int32(-8)
	if base.Ui32(v565) <= base.Ui32(v54+int32(16)) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v569 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v54 | v559&v569 | int32(2)
	v575 = v553 + v54
	v576 = v565 - v54
	*(*int32)(unsafe.Add(mBase, uint32(v575)+4)) = v576 | int32(3)
	v580 = v553 + v565
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+4)) = v581 | v569
	v592 = v575 + v576
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v593&v569 != 0 {
		v710 = v575
		v711 = v576
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
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	if v719&int32(2) == int32(0) {
		goto L192
	} else {
		goto L193
	}
L149:
	;
	if v593&int32(2) == int32(0) {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	v601 = v600 + v576
	v602 = v575 - v600
	v604 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
	if v602 != v604 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	if v620 == int32(0) {
		v710 = v602
		v711 = v601
		goto L148
	} else {
		goto L173
	}
L152:
	;
	v671 = int32(0)
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v609)+12)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v606)+8)) = v609
	v710 = v602
	v711 = v601
	goto L148
L154:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	if base.Ui32(v600) <= base.Ui32(int32(255)) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	v653 = int32(3)
	if v652&v653 != v653 {
		v710 = v602
		v711 = v601
		goto L148
	} else {
		goto L172
	}
L157:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v602)+8))
	if v606 != v609 {
		goto L153
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v602)+24))
	if v602 != v606 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v611 = int32(4747256)
	v613 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v613 & base.I32_rotl(int32(-2), int32(base.Ui32(v600)>>(uint(int32(3))%32)))
	v710 = v602
	v711 = v601
	goto L148
L161:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v602)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v622)+12)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v606)+8)) = v622
	v671 = v606
	goto L151
L162:
	;
	goto L163
L163:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v602)+20))
	if v625 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v633 = v625
	v634 = v602 + int32(20)
	goto L166
L165:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v602)+16))
	if v628 == int32(0) {
		goto L152
	} else {
		goto L167
	}
L166:
	;
	__phi638 = v633
	__phi640 = v634
	v638 = __phi638
	v640 = __phi640
	goto L168
L167:
	;
	v633 = v628
	v634 = v602 + int32(16)
	goto L166
L168:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v638)+20))
	if v646 != 0 {
		__phi638 = v646
		__phi640 = v638 + int32(20)
		v638 = __phi638
		v640 = __phi640
		goto L168
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = int32(0)
	v671 = v638
	goto L151
L170:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v638)+16))
	if v649 != 0 {
		__phi638 = v649
		__phi640 = v638 + int32(16)
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
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v601
	*(*int32)(unsafe.Add(mBase, uint32(v592)+4)) = v652 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v602)+4)) = v601 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = v601
	goto L146
L173:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v602)+28))
	v682 = v680 << (uint(int32(2)) % 32)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v682)+uint32(_consts[1507])))
	if v685 == v602 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+24)) = v620
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v602)+16))
	if v702 != 0 {
		goto L184
	} else {
		goto L185
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v682)+uint32(_consts[1507]))) = v671
	if v671 != 0 {
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v620)+16))
	if v602 == v695 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v688 = int32(4747260)
	v690 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v690 & base.I32_rotl(int32(-2), v680)
	v710 = v602
	v711 = v601
	goto L148
L179:
	;
	if v671 == int32(0) {
		v710 = v602
		v711 = v601
		goto L148
	} else {
		goto L183
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620)+16)) = v671
	goto L179
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620)+20)) = v671
	goto L179
L183:
	;
	goto L174
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+16)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v702)+24)) = v671
	goto L186
L185:
	;
	goto L186
L186:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v602)+20))
	if v705 == int32(0) {
		v710 = v602
		v711 = v601
		goto L148
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+20)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v705)+24)) = v671
	v710 = v602
	v711 = v601
	goto L148
L188:
	;
	if base.Ui32(v880) <= base.Ui32(int32(255)) {
		goto L235
	} else {
		goto L236
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v763 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v710+v763))) = v763
	if v710 != v747 {
		v880 = v763
		goto L188
	} else {
		goto L234
	}
L190:
	;
	if v780 == int32(0) {
		goto L189
	} else {
		goto L219
	}
L191:
	;
	v823 = int32(0)
	goto L190
L192:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[1509]))
	if v725 == v592 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v592)+4)) = v719 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v711 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v710+v711))) = v711
	v880 = v711
	goto L188
L195:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1509])) = v710
	v729 = int32(4747268)
	v731 = *(*int32)(unsafe.Add(mBase, _consts[1510]))
	v732 = v731 + v711
	*(*int32)(unsafe.Add(mBase, _consts[1510])) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v732 | int32(1)
	v738 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
	if v710 != v738 {
		goto L147
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _consts[1504]))
	if v747 == v592 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v741 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v741
	*(*int32)(unsafe.Add(mBase, _consts[1504])) = v741
	goto L146
L199:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1504])) = v710
	v751 = int32(4747264)
	v753 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	v754 = v753 + v711
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v754
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v754 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v710+v754))) = v754
	goto L146
L200:
	;
	goto L201
L201:
	;
	v763 = v719&int32(-8) + v711
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v592)+12))
	if base.Ui32(v719) <= base.Ui32(int32(255)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v592)+8))
	if v767 == v764 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v592)+24))
	if v764 != v592 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v769 = int32(4747256)
	v771 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v771 & base.I32_rotl(int32(-2), int32(base.Ui32(v719)>>(uint(int32(3))%32)))
	goto L189
L206:
	;
	goto L207
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+12)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v764)+8)) = v767
	goto L189
L208:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v592)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v782)+12)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v764)+8)) = v782
	v823 = v764
	goto L190
L209:
	;
	goto L210
L210:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v592)+20))
	if v785 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v793 = v785
	v794 = v592 + int32(20)
	goto L213
L212:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v592)+16))
	if v788 == int32(0) {
		goto L191
	} else {
		goto L214
	}
L213:
	;
	__phi798 = v793
	__phi800 = v794
	v798 = __phi798
	v800 = __phi800
	goto L215
L214:
	;
	v793 = v788
	v794 = v592 + int32(16)
	goto L213
L215:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v798)+20))
	if v806 != 0 {
		__phi798 = v806
		__phi800 = v798 + int32(20)
		v798 = __phi798
		v800 = __phi800
		goto L215
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v800))) = int32(0)
	v823 = v798
	goto L190
L217:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v798)+16))
	if v809 != 0 {
		__phi798 = v809
		__phi800 = v798 + int32(16)
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
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v592)+28))
	v834 = v832 << (uint(int32(2)) % 32)
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v834)+uint32(_consts[1507])))
	if v837 == v592 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823)+24)) = v780
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v592)+16))
	if v854 != 0 {
		goto L230
	} else {
		goto L231
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v834)+uint32(_consts[1507]))) = v823
	if v823 != 0 {
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v780)+16))
	if v592 == v847 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v840 = int32(4747260)
	v842 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v842 & base.I32_rotl(int32(-2), v832)
	goto L189
L225:
	;
	if v823 == int32(0) {
		goto L189
	} else {
		goto L229
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v780)+16)) = v823
	goto L225
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v780)+20)) = v823
	goto L225
L229:
	;
	goto L220
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823)+16)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v854)+24)) = v823
	goto L232
L231:
	;
	goto L232
L232:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v592)+20))
	if v857 == int32(0) {
		goto L189
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823)+20)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v857)+24)) = v823
	goto L189
L234:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v763
	goto L146
L235:
	;
	v891 = v880 & int32(-8)
	v893 = v891 + int32(4747296)
	v895 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	v899 = int32(1) << (uint(int32(base.Ui32(v880)>>(uint(int32(3))%32))) % 32)
	if v895&v899 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	goto L237
L237:
	;
	if base.Ui32(v880) <= base.Ui32(int32(16777215)) {
		goto L242
	} else {
		goto L243
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891)+uint32(_consts[1511]))) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v907)+12)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v710)+12)) = v893
	*(*int32)(unsafe.Add(mBase, uint32(v710)+8)) = v907
	goto L146
L239:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v899 | v895
	v907 = v893
	goto L238
L240:
	;
	goto L241
L241:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v891)+uint32(_consts[1511])))
	v907 = v906
	goto L238
L242:
	;
	v918 = base.I32_clz(int32(base.Ui32(v880) >> (uint(int32(8)) % 32)))
	v921 = int32(1)
	v928 = int32(base.Ui32(v880)>>(uint(int32(38)-v918)%32))&v921 - v918<<(uint(v921)%32) + int32(62)
	goto L244
L243:
	;
	v928 = int32(31)
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+28)) = v928
	*(*int64)(unsafe.Add(mBase, uint32(v710)+16)) = int64(0)
	v933 = v928 << (uint(int32(2)) % 32)
	v937 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	v939 = int32(1) << (uint(v928) % 32)
	if v937&v939 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v963)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v992)+12)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v963)+8)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v710)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v710)+12)) = v963
	*(*int32)(unsafe.Add(mBase, uint32(v710)+8)) = v992
	goto L147
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+12)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v710)+8)) = v710
	goto L146
L247:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = v939 | v937
	*(*int32)(unsafe.Add(mBase, uint32(v933)+uint32(_consts[1507]))) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v710)+24)) = v933 + int32(4747560)
	goto L246
L248:
	;
	goto L249
L249:
	;
	if v928 != int32(31) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v955 = int32(25) - int32(base.Ui32(v928)>>(uint(int32(1))%32))
	goto L252
L251:
	;
	v955 = int32(0)
	goto L252
L252:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v933)+uint32(_consts[1507])))
	v960 = v880 << (uint(v955) % 32)
	v963 = v957
	goto L253
L253:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v963)+4))
	if v967&int32(-8) == v880 {
		goto L245
	} else {
		goto L255
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v977)+16)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v710)+24)) = v963
	goto L246
L255:
	;
	v977 = v963 + int32(base.Ui32(v960)>>(uint(int32(29))%32))&int32(4)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)+16))
	if v978 != 0 {
		v960 = v960 << (uint(int32(1)) % 32)
		v963 = v978
		goto L253
	} else {
		goto L256
	}
L256:
	;
	goto L254
}
