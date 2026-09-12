package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageBtreeConsolidate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var __phi38 int32
	_ = __phi38
	var v39 int32
	_ = v39
	var __phi39 int32
	_ = __phi39
	var v45 int32
	_ = v45
	var __phi45 int32
	_ = __phi45
	var v46 int32
	_ = v46
	var __phi46 int32
	_ = __phi46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v318 int32
	_ = v318
	var __phi318 int32
	_ = __phi318
	var v323 int32
	_ = v323
	var __phi323 int32
	_ = __phi323
	var v329 int32
	_ = v329
	var __phi329 int32
	_ = __phi329
	var v330 int32
	_ = v330
	var __phi330 int32
	_ = __phi330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v513 int32
	_ = v513
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v561 int32
	_ = v561
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(169)) < base.Ui32(v19) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = int32(1)
	v28 = l0 - v25 + v27
	v29 = v28 + v22
	v31 = v29 - v27
	if v31 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = v28 - int32(1)
	__phi38 = v31
	__phi39 = l1
	__phi45 = int32(0)
	__phi46 = v29
	v38 = __phi38
	v39 = __phi39
	v45 = __phi45
	v46 = __phi46
	goto L7
L5:
	;
	F_FreePageBtreeRemovePage(m, l0, v561)
	mBase = m.M
	goto L1
L6:
	;
	__phi318 = l1
	__phi323 = int32(0)
	__phi329 = v29
	__phi330 = v31
	v318 = __phi318
	v323 = __phi323
	v329 = __phi329
	v330 = __phi330
	goto L78
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+3))
	v61 = int32(0)
	v62 = v58
	goto L9
L8:
	;
	v107 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v46+v93<<(uint(int32(3))%32))+23))
	if v111 != 0 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	if base.Ui32(v62) <= base.Ui32(v61) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.Ui32(v58-int32(1)) <= base.Ui32(v93) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L10
L12:
	;
	v93 = v61
	goto L11
L13:
	;
	goto L14
L14:
	;
	v79 = int32(1)
	v80 = int32(base.Ui32(v61+v62) >> (uint(v79) % 32))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(11)+v80<<(uint(int32(3))%32))))
	v87 = base.B2i32(base.Ui32(v56) < base.Ui32(v86))
	if base.Ui32(v56) < base.Ui32(v86) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v88 = v61
	goto L17
L16:
	;
	v88 = v80 + v79
	goto L17
L17:
	;
	if base.Ui32(v56) < base.Ui32(v86) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v89 = v80
	goto L20
L19:
	;
	v89 = v62
	goto L20
L20:
	;
	if v56 != v86 {
		v61 = v88
		v62 = v89
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v93 = v80
	goto L11
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v99 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L8
L25:
	;
	v102 = int32(1)
	v104 = v99 + v28
	v106 = v104 - v102
	if v106 != 0 {
		__phi38 = v106
		__phi39 = v38
		__phi45 = v45 + v102
		__phi46 = v104
		v38 = __phi38
		v39 = __phi39
		v45 = __phi45
		v46 = __phi46
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L6
L27:
	;
	v114 = v35 + v111
	goto L29
L28:
	;
	v114 = v107
	goto L29
L29:
	;
	if v45 <= int32(0) {
		v211 = v114
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v211 == int32(0) {
		goto L6
	} else {
		goto L58
	}
L31:
	;
	v118 = v45 & int32(3)
	if v118 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.Ui32(v45) < base.Ui32(int32(4)) {
		v211 = v151
		goto L30
	} else {
		goto L42
	}
L33:
	;
	v150 = v45
	v151 = v114
	goto L32
L34:
	;
	goto L35
L35:
	;
	v123 = v45
	v124 = v114
	v126 = v107
	goto L36
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	if v139 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v150 = v144
	v151 = v142
	goto L32
L38:
	;
	v142 = v35 + v139
	goto L40
L39:
	;
	v142 = int32(0)
	goto L40
L40:
	;
	v143 = int32(1)
	v144 = v123 - v143
	v146 = v126 + v143
	if v146 != v118 {
		v123 = v144
		v124 = v142
		v126 = v146
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v170 = v150
	v171 = v151
	goto L43
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	if v186 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v211 = v201
	goto L30
L45:
	;
	v189 = v35 + v186
	goto L47
L46:
	;
	v189 = int32(0)
	goto L47
L47:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	if v190 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v193 = v35 + v190
	goto L50
L49:
	;
	v193 = int32(0)
	goto L50
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	if v194 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v197 = v35 + v194
	goto L53
L52:
	;
	v197 = int32(0)
	goto L53
L53:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	if v198 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v201 = v35 + v198
	goto L56
L55:
	;
	v201 = int32(0)
	goto L56
L56:
	;
	if base.Ui32(v170-int32(5)) < base.Ui32(int32(-2)) {
		v170 = v170 - int32(4)
		v171 = v201
		goto L43
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if base.Ui32(int32(510)) < base.Ui32(v228+v19) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v233 = v228 << (uint(int32(3)) % 32)
	v234 = int32(12)
	v235 = v211 + v234
	v237 = l1 + v234
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v238 == int32(-1729435864) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v233 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	if v233 != 0 {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v246 + v247
	v561 = v211
	goto L5
L64:
	;
	v244 = F__emscripten_memcpy_bulkmem(m, v237+v19<<(uint(int32(3))%32), v235, v233)
	mBase = m.M
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v257 = v255 + v256
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v257
	if v257 == int32(0) {
		v561 = v211
		goto L5
	} else {
		goto L71
	}
L68:
	;
	v253 = F__emscripten_memcpy_bulkmem(m, v237+v19<<(uint(int32(3))%32), v235, v233)
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v269 = int32(0)
	goto L72
L72:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v269<<(uint(int32(3))%32))))
	if v288 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v561 = v211
	goto L5
L74:
	;
	v291 = v35 + v288
	goto L76
L75:
	;
	v291 = int32(0)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = l1 - v28 + int32(1)
	v294 = v269 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v294) < base.Ui32(v295) {
		v269 = v294
		goto L72
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v329)+3))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v318)+12))
	v341 = int32(0)
	v342 = v336
	goto L80
L79:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v329+v373<<(uint(int32(3))%32))+7))
	if v389 != 0 {
		goto L98
	} else {
		goto L99
	}
L80:
	;
	if base.Ui32(v342) <= base.Ui32(v341) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v373 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	goto L81
L83:
	;
	v373 = v341
	goto L82
L84:
	;
	goto L85
L85:
	;
	v359 = int32(1)
	v360 = int32(base.Ui32(v341+v342) >> (uint(v359) % 32))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v329+int32(11)+v360<<(uint(int32(3))%32))))
	v367 = base.B2i32(base.Ui32(v337) < base.Ui32(v366))
	if base.Ui32(v337) < base.Ui32(v366) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v368 = v341
	goto L88
L87:
	;
	v368 = v360 + v359
	goto L88
L88:
	;
	if base.Ui32(v337) < base.Ui32(v366) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v369 = v360
	goto L91
L90:
	;
	v369 = v342
	goto L91
L91:
	;
	if v337 != v366 {
		v341 = v368
		v342 = v369
		goto L80
	} else {
		goto L92
	}
L92:
	;
	v373 = v360
	goto L82
L93:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	if v378 == int32(0) {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	goto L79
L96:
	;
	v381 = int32(1)
	v383 = v378 + v28
	v385 = v383 - v381
	if v385 != 0 {
		__phi318 = v330
		__phi323 = v323 + v381
		__phi329 = v383
		__phi330 = v385
		v318 = __phi318
		v323 = __phi323
		v329 = __phi329
		v330 = __phi330
		goto L78
	} else {
		goto L97
	}
L97:
	;
	goto L1
L98:
	;
	v392 = v35 + v389
	goto L100
L99:
	;
	v392 = int32(0)
	goto L100
L100:
	;
	if v323 <= int32(0) {
		v453 = v392
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v453 == int32(0) {
		goto L1
	} else {
		goto L119
	}
L102:
	;
	if v323&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v392+v397<<(uint(int32(3))%32))+8))
	if v401 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v407 = v392
	v408 = v323
	goto L105
L105:
	;
	if v323 == int32(1) {
		v453 = v407
		goto L101
	} else {
		goto L109
	}
L106:
	;
	v404 = v35 + v401
	goto L108
L107:
	;
	v404 = int32(0)
	goto L108
L108:
	;
	v407 = v404
	v408 = v323 - int32(1)
	goto L105
L109:
	;
	v413 = v407
	v414 = v408
	goto L110
L110:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v413+v429<<(uint(int32(3))%32))+8))
	if v433 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v453 = v444
	goto L101
L112:
	;
	v436 = v35 + v433
	goto L114
L113:
	;
	v436 = int32(0)
	goto L114
L114:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v436+v437<<(uint(int32(3))%32))+8))
	if v441 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v444 = v35 + v441
	goto L117
L116:
	;
	v444 = int32(0)
	goto L117
L117:
	;
	if base.Ui32(v414-int32(3)) < base.Ui32(int32(-2)) {
		v413 = v444
		v414 = v414 - int32(2)
		goto L110
	} else {
		goto L118
	}
L118:
	;
	goto L111
L119:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if base.Ui32(int32(510)) < base.Ui32(v471+v19) {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v476 = v19 << (uint(int32(3)) % 32)
	v477 = int32(12)
	v478 = l1 + v477
	v480 = v453 + v477
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v481 == int32(-1729435864) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v561 = l1
	goto L5
L122:
	;
	if v476 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	goto L124
L124:
	;
	if v476 != 0 {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v489 + v490
	goto L121
L126:
	;
	v487 = F__emscripten_memcpy_bulkmem(m, v480+v471<<(uint(int32(3))%32), v478, v476)
	mBase = m.M
	goto L128
L127:
	;
	goto L128
L128:
	;
	goto L125
L129:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v500 = v498 + v499
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v500
	if v500 == int32(0) {
		goto L121
	} else {
		goto L133
	}
L130:
	;
	v496 = F__emscripten_memcpy_bulkmem(m, v480+v471<<(uint(int32(3))%32), v478, v476)
	mBase = m.M
	goto L132
L131:
	;
	goto L132
L132:
	;
	goto L129
L133:
	;
	v513 = int32(0)
	goto L134
L134:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v453+int32(16)+v513<<(uint(int32(3))%32))))
	if v531 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L121
L136:
	;
	v534 = v35 + v531
	goto L138
L137:
	;
	v534 = int32(0)
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v453 - v28 + int32(1)
	v537 = v513 + int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if base.Ui32(v537) < base.Ui32(v538) {
		v513 = v537
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
}
func F_FreePageBtreeRemovePage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var __phi27 int32
	_ = __phi27
	var v28 int32
	_ = v28
	var __phi28 int32
	_ = __phi28
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v98 = v29 + int32(11)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v101 == int32(-1729435864) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
	return
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(1)
	v20 = l0 - v17 + v19
	v21 = v20 + v14
	v23 = v21 - v19
	if v23 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	__phi27 = l1
	__phi28 = v23
	__phi29 = v21
	v27 = __phi27
	v28 = __phi28
	v29 = __phi29
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+3))
	if base.Ui32(int32(1)) < base.Ui32(v39) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = int32(1)
	v46 = l0 - v43 + v45
	v49 = (v27 - v20) & int32(-4096)
	v50 = v46 + v49
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(8225038576)
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v53
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = v46 + v42 - v45
	goto L10
L9:
	;
	v59 = v53
	goto L10
L10:
	;
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = v59 - v46 + int32(1)
	goto L13
L12:
	;
	v64 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v64
	v67 = v49 | int32(1)
	if v59 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v67
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v70 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v74 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v77 = v74 + v20
	v79 = v77 - int32(1)
	if v79 != 0 {
		__phi27 = v28
		__phi28 = v79
		__phi29 = v77
		v27 = __phi27
		v28 = __phi28
		v29 = __phi29
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L6
L19:
	;
	v496 = v29 + int32(3)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v498 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v497 - v498
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v505 = l0 - v502 + v498
	v508 = (v27 - v20) & int32(-4096)
	v509 = v505 + v508
	*(*int64)(unsafe.Add(mBase, uint32(v509))) = int64(8225038576)
	v512 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v509)+8)) = v512
	if v501 != 0 {
		goto L143
	} else {
		goto L144
	}
L20:
	;
	v106 = v100
	v109 = v39
	goto L23
L21:
	;
	goto L22
L22:
	;
	v295 = v100
	v298 = v39
	goto L83
L23:
	;
	if base.Ui32(v109) <= base.Ui32(v106) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if base.Ui32(v39-int32(1)) <= base.Ui32(v132) {
		v486 = v132
		goto L19
	} else {
		goto L36
	}
L25:
	;
	goto L24
L26:
	;
	v132 = v106
	goto L25
L27:
	;
	goto L28
L28:
	;
	v119 = int32(1)
	v120 = int32(base.Ui32(v106+v109) >> (uint(v119) % 32))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v98+v120<<(uint(int32(3))%32))))
	v127 = base.B2i32(base.Ui32(v99) < base.Ui32(v126))
	if base.Ui32(v99) < base.Ui32(v126) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v128 = v106
	goto L31
L30:
	;
	v128 = v120 + v119
	goto L31
L31:
	;
	if base.Ui32(v99) < base.Ui32(v126) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v129 = v120
	goto L34
L33:
	;
	v129 = v109
	goto L34
L34:
	;
	if v99 != v126 {
		v106 = v128
		v109 = v129
		goto L23
	} else {
		goto L35
	}
L35:
	;
	v132 = v120
	goto L25
L36:
	;
	v139 = int32(3)
	v141 = v98 + v132<<(uint(v139)%32)
	v143 = v141 + int32(8)
	v148 = (v39 + (v132 ^ int32(-1))) << (uint(v139) % 32)
	if v141 == v143 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v486 = v132
	goto L19
L38:
	;
	goto L37
L39:
	;
	v152 = v141 + v148
	if base.Ui32(v143-v152) <= base.Ui32(int32(0)-v148<<(uint(int32(1))%32)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v159 = F___memcpy(m, v141, v143, v148)
	mBase = m.M
	goto L37
L41:
	;
	goto L42
L42:
	;
	v162 = (v141 ^ v143) & int32(3)
	if base.Ui32(v141) < base.Ui32(v143) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v264 == int32(0) {
		goto L38
	} else {
		goto L79
	}
L44:
	;
	if base.Ui32(v242) <= base.Ui32(int32(3)) {
		v263 = v241
		v264 = v242
		v265 = v243
		goto L43
	} else {
		goto L75
	}
L45:
	;
	if v162 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v162 != 0 {
		v224 = v148
		goto L58
	} else {
		goto L59
	}
L48:
	;
	v263 = v143
	v264 = v148
	v265 = v141
	goto L43
L49:
	;
	goto L50
L50:
	;
	if v141&int32(3) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v241 = v143
	v242 = v148
	v243 = v141
	goto L44
L52:
	;
	goto L53
L53:
	;
	v169 = v143
	v170 = v148
	v171 = v141
	goto L54
L54:
	;
	if v170 == int32(0) {
		goto L38
	} else {
		goto L56
	}
L55:
	;
	v241 = v178
	v242 = v180
	v243 = v182
	goto L44
L56:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v175)
	v177 = int32(1)
	v178 = v169 + v177
	v180 = v170 - v177
	v182 = v171 + v177
	if v182&int32(3) != 0 {
		v169 = v178
		v170 = v180
		v171 = v182
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v224 == int32(0) {
		goto L38
	} else {
		goto L71
	}
L59:
	;
	if v152&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v189 = v148
	goto L63
L61:
	;
	v204 = v148
	goto L62
L62:
	;
	if base.Ui32(v204) <= base.Ui32(int32(3)) {
		v224 = v204
		goto L58
	} else {
		goto L67
	}
L63:
	;
	if v189 == int32(0) {
		goto L38
	} else {
		goto L65
	}
L64:
	;
	v204 = v195
	goto L62
L65:
	;
	v195 = v189 - int32(1)
	v196 = v141 + v195
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v198)
	if v196&int32(3) != 0 {
		v189 = v195
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v211 = v204
	goto L68
L68:
	;
	v215 = v211 - int32(4)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v143+v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v141+v215))) = v218
	if base.Ui32(int32(3)) < base.Ui32(v215) {
		v211 = v215
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v224 = v215
	goto L58
L70:
	;
	goto L69
L71:
	;
	v231 = v224
	goto L72
L72:
	;
	v235 = v231 - int32(1)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v235))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141+v235))) = uint8(v238)
	if v235 != 0 {
		v231 = v235
		goto L72
	} else {
		goto L74
	}
L73:
	;
	goto L38
L74:
	;
	goto L73
L75:
	;
	v248 = v241
	v249 = v242
	v250 = v243
	goto L76
L76:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v252
	v254 = int32(4)
	v255 = v248 + v254
	v257 = v250 + v254
	v259 = v249 - v254
	if base.Ui32(int32(3)) < base.Ui32(v259) {
		v248 = v255
		v249 = v259
		v250 = v257
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v263 = v255
	v264 = v259
	v265 = v257
	goto L43
L78:
	;
	goto L77
L79:
	;
	v270 = v263
	v271 = v264
	v272 = v265
	goto L80
L80:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v274)
	v276 = int32(1)
	v281 = v271 - v276
	if v281 != 0 {
		v270 = v270 + v276
		v271 = v281
		v272 = v272 + v276
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L38
L82:
	;
	goto L81
L83:
	;
	if base.Ui32(v298) <= base.Ui32(v295) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if base.Ui32(v39-int32(1)) <= base.Ui32(v321) {
		v486 = v321
		goto L19
	} else {
		goto L96
	}
L85:
	;
	goto L84
L86:
	;
	v321 = v295
	goto L85
L87:
	;
	goto L88
L88:
	;
	v308 = int32(1)
	v309 = int32(base.Ui32(v295+v298) >> (uint(v308) % 32))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v98+v309<<(uint(int32(3))%32))))
	v316 = base.B2i32(base.Ui32(v99) < base.Ui32(v315))
	if base.Ui32(v99) < base.Ui32(v315) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v317 = v295
	goto L91
L90:
	;
	v317 = v309 + v308
	goto L91
L91:
	;
	if base.Ui32(v99) < base.Ui32(v315) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v318 = v309
	goto L94
L93:
	;
	v318 = v298
	goto L94
L94:
	;
	if v99 != v315 {
		v295 = v317
		v298 = v318
		goto L83
	} else {
		goto L95
	}
L95:
	;
	v321 = v309
	goto L85
L96:
	;
	v328 = int32(3)
	v330 = v98 + v321<<(uint(v328)%32)
	v332 = v330 + int32(8)
	v337 = (v39 + (v321 ^ int32(-1))) << (uint(v328) % 32)
	if v330 == v332 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v486 = v321
	goto L19
L98:
	;
	goto L97
L99:
	;
	v341 = v330 + v337
	if base.Ui32(v332-v341) <= base.Ui32(int32(0)-v337<<(uint(int32(1))%32)) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v348 = F___memcpy(m, v330, v332, v337)
	mBase = m.M
	goto L97
L101:
	;
	goto L102
L102:
	;
	v351 = (v330 ^ v332) & int32(3)
	if base.Ui32(v330) < base.Ui32(v332) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v453 == int32(0) {
		goto L98
	} else {
		goto L139
	}
L104:
	;
	if base.Ui32(v431) <= base.Ui32(int32(3)) {
		v452 = v430
		v453 = v431
		v454 = v432
		goto L103
	} else {
		goto L135
	}
L105:
	;
	if v351 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	if v351 != 0 {
		v413 = v337
		goto L118
	} else {
		goto L119
	}
L108:
	;
	v452 = v332
	v453 = v337
	v454 = v330
	goto L103
L109:
	;
	goto L110
L110:
	;
	if v330&int32(3) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v430 = v332
	v431 = v337
	v432 = v330
	goto L104
L112:
	;
	goto L113
L113:
	;
	v358 = v332
	v359 = v337
	v360 = v330
	goto L114
L114:
	;
	if v359 == int32(0) {
		goto L98
	} else {
		goto L116
	}
L115:
	;
	v430 = v367
	v431 = v369
	v432 = v371
	goto L104
L116:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	*(*uint8)(unsafe.Add(mBase, uint32(v360))) = uint8(v364)
	v366 = int32(1)
	v367 = v358 + v366
	v369 = v359 - v366
	v371 = v360 + v366
	if v371&int32(3) != 0 {
		v358 = v367
		v359 = v369
		v360 = v371
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if v413 == int32(0) {
		goto L98
	} else {
		goto L131
	}
L119:
	;
	if v341&int32(3) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v378 = v337
	goto L123
L121:
	;
	v393 = v337
	goto L122
L122:
	;
	if base.Ui32(v393) <= base.Ui32(int32(3)) {
		v413 = v393
		goto L118
	} else {
		goto L127
	}
L123:
	;
	if v378 == int32(0) {
		goto L98
	} else {
		goto L125
	}
L124:
	;
	v393 = v384
	goto L122
L125:
	;
	v384 = v378 - int32(1)
	v385 = v330 + v384
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332+v384))))
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v387)
	if v385&int32(3) != 0 {
		v378 = v384
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v400 = v393
	goto L128
L128:
	;
	v404 = v400 - int32(4)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v332+v404)))
	*(*int32)(unsafe.Add(mBase, uint32(v330+v404))) = v407
	if base.Ui32(int32(3)) < base.Ui32(v404) {
		v400 = v404
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v413 = v404
	goto L118
L130:
	;
	goto L129
L131:
	;
	v420 = v413
	goto L132
L132:
	;
	v424 = v420 - int32(1)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332+v424))))
	*(*uint8)(unsafe.Add(mBase, uint32(v330+v424))) = uint8(v427)
	if v424 != 0 {
		v420 = v424
		goto L132
	} else {
		goto L134
	}
L133:
	;
	goto L98
L134:
	;
	goto L133
L135:
	;
	v437 = v430
	v438 = v431
	v439 = v432
	goto L136
L136:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v441
	v443 = int32(4)
	v444 = v437 + v443
	v446 = v439 + v443
	v448 = v438 - v443
	if base.Ui32(int32(3)) < base.Ui32(v448) {
		v437 = v444
		v438 = v448
		v439 = v446
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v452 = v444
	v453 = v448
	v454 = v446
	goto L103
L138:
	;
	goto L137
L139:
	;
	v459 = v452
	v460 = v453
	v461 = v454
	goto L140
L140:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v463)
	v465 = int32(1)
	v470 = v460 - v465
	if v470 != 0 {
		v459 = v459 + v465
		v460 = v470
		v461 = v461 + v465
		goto L140
	} else {
		goto L142
	}
L141:
	;
	goto L98
L142:
	;
	goto L141
L143:
	;
	v518 = v501 + v505 - v498
	goto L145
L144:
	;
	v518 = v512
	goto L145
L145:
	;
	if v518 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v523 = v518 - v505 + int32(1)
	goto L148
L147:
	;
	v523 = int32(0)
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509)+12)) = v523
	v526 = v508 | int32(1)
	if v518 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+8)) = v526
	goto L151
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v526
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v529 + int32(1)
	if v486 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_FreePageBtreeConsolidate(m, l0, v28)
	mBase = m.M
	return
L153:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v534 = l0 - v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v29)+11))
	v544 = v28
	goto L154
L154:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v544)+8))
	if v549 == int32(0) {
		goto L152
	} else {
		goto L156
	}
L155:
	;
	goto L152
L156:
	;
	v552 = v549 + v534
	if v552 == int32(0) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	v556 = v552 + int32(12)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	v561 = v558
	v562 = int32(0)
	goto L158
L158:
	;
	if base.Ui32(v561) <= base.Ui32(v562) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	if base.Ui32(v586) < base.Ui32(v558) {
		goto L171
	} else {
		goto L172
	}
L160:
	;
	goto L159
L161:
	;
	v586 = v562
	goto L160
L162:
	;
	goto L163
L163:
	;
	v574 = int32(1)
	v575 = int32(base.Ui32(v561+v562) >> (uint(v574) % 32))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v556+v575<<(uint(int32(3))%32))))
	v582 = base.B2i32(base.Ui32(v535) < base.Ui32(v581))
	if base.Ui32(v535) < base.Ui32(v581) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v583 = v562
	goto L166
L165:
	;
	v583 = v575 + v574
	goto L166
L166:
	;
	if base.Ui32(v535) < base.Ui32(v581) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v584 = v575
	goto L169
L168:
	;
	v584 = v561
	goto L169
L169:
	;
	if v581 != v535 {
		v561 = v584
		v562 = v583
		goto L158
	} else {
		goto L170
	}
L170:
	;
	v586 = v575
	goto L160
L171:
	;
	v594 = int32(0)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v552+v586<<(uint(int32(3))%32))+16))
	if v598 != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v605 = int32(-1)
	goto L173
L173:
	;
	v606 = v605 + v586
	*(*int32)(unsafe.Add(mBase, uint32(v556+v606<<(uint(int32(3))%32)))) = v535
	if v606 == int32(0) {
		v544 = v552
		goto L154
	} else {
		goto L180
	}
L174:
	;
	v601 = v534 + v598
	goto L176
L175:
	;
	v601 = v594
	goto L176
L176:
	;
	if v601 != v544 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v603 = int32(-1)
	goto L179
L178:
	;
	v603 = v594
	goto L179
L179:
	;
	v605 = v603
	goto L173
L180:
	;
	goto L155
}
func F_FreePageManagerGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v7 = F_FreePageManagerGetInternal(m, l0, l1, l2)
	mBase = m.M
	v8 = F_FreePageBtreeCleanup(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v12) < base.Ui32(v8) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
	goto L5
L4:
	;
	goto L5
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v15 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+548))
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	return v7
L9:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v74
	goto L8
L10:
	;
	v22 = l0 + int32(36)
	v25 = int32(128)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = l0 - v55
	v60 = v56 + v18
	v61 = int32(0)
	goto L22
L13:
	;
	v31 = v25 - int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22+v31<<(uint(int32(2))%32))))
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v74 = int32(0)
	goto L9
L15:
	;
	v74 = v25
	goto L9
L16:
	;
	goto L17
L17:
	;
	v36 = int32(2)
	v37 = v25 - v36
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22+v37<<(uint(v36)%32))))
	if v41 != 0 {
		v74 = v31
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v43 = v25 - int32(3)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22+v43<<(uint(int32(2))%32))))
	if v47 != 0 {
		v74 = v37
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v49 = v25 - int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22+v49<<(uint(int32(2))%32))))
	if v53 != 0 {
		v74 = v43
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v49 != 0 {
		v25 = v49
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if base.Ui32(v61) < base.Ui32(v65) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v74 = v67
	goto L9
L24:
	;
	v67 = v65
	goto L26
L25:
	;
	v67 = v61
	goto L26
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v68 == int32(0) {
		v74 = v67
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v71 = v68 + v56
	if v71 != 0 {
		v60 = v71
		v61 = v67
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L23
}
func F_FreePageManagerPutInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
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
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var __phi535 int32
	_ = __phi535
	var v536 int32
	_ = v536
	var __phi536 int32
	_ = __phi536
	var v541 int32
	_ = v541
	var __phi541 int32
	_ = __phi541
	var v549 int32
	_ = v549
	var __phi549 int32
	_ = __phi549
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v748 int32
	_ = v748
	var v767 int32
	_ = v767
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1163 int32
	_ = v1163
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1295 int32
	_ = v1295
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1560 int32
	_ = v1560
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1639 int32
	_ = v1639
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1659 int32
	_ = v1659
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1716 int32
	_ = v1716
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1750 int32
	_ = v1750
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1954 int32
	_ = v1954
	var v1966 int32
	_ = v1966
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1991 int32
	_ = v1991
	var v1996 int32
	_ = v1996
	var v2003 int32
	_ = v2003
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2063 int32
	_ = v2063
	var v2068 int32
	_ = v2068
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2082 int32
	_ = v2082
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2246 int32
	_ = v2246
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2261 int32
	_ = v2261
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2281 int32
	_ = v2281
	var v2288 int32
	_ = v2288
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2316 int32
	_ = v2316
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2338 int32
	_ = v2338
	var v2352 int32
	_ = v2352
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2375 int32
	_ = v2375
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2452 int32
	_ = v2452
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2518 int32
	_ = v2518
	var v2522 int32
	_ = v2522
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2559 int32
	_ = v2559
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2568 int32
	_ = v2568
	var v2574 int32
	_ = v2574
	var v2581 int32
	_ = v2581
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2594 int32
	_ = v2594
	var v2601 int32
	_ = v2601
	var v2605 int32
	_ = v2605
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2651 int32
	_ = v2651
	var v2664 int32
	_ = v2664
	var v2668 int32
	_ = v2668
	var v2677 int32
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2695 int32
	_ = v2695
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2772 int32
	_ = v2772
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2809 int32
	_ = v2809
	var v2812 int32
	_ = v2812
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2838 int32
	_ = v2838
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2862 int32
	_ = v2862
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2871 int32
	_ = v2871
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2912 int32
	_ = v2912
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2928 int32
	_ = v2928
	var v2932 int32
	_ = v2932
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2962 int32
	_ = v2962
	var v2969 int32
	_ = v2969
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2984 int32
	_ = v2984
	var v2991 int32
	_ = v2991
	var v2995 int32
	_ = v2995
	var v2998 int32
	_ = v2998
	var v3004 int32
	_ = v3004
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3032 int32
	_ = v3032
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3061 int32
	_ = v3061
	var v3074 int32
	_ = v3074
	var v3078 int32
	_ = v3078
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3106 int32
	_ = v3106
	var v3114 int32
	_ = v3114
	var v3117 int32
	_ = v3117
	var v3121 int32
	_ = v3121
	var v3123 int32
	_ = v3123
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3165 int32
	_ = v3165
	var v3170 int32
	_ = v3170
	var v3174 int32
	_ = v3174
	var v3177 int32
	_ = v3177
	var v3179 int32
	_ = v3179
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3235 int32
	_ = v3235
	var v3241 int32
	_ = v3241
	var v3246 int32
	_ = v3246
	var v3249 int32
	_ = v3249
	var v3281 int32
	_ = v3281
	v5 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = l0 - v30 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v28 + int32(32)
	return v3281
L2:
	;
	v3281 = l2
	goto L1
L3:
	;
	v1980 = int32(1)
	v1981 = l0 - v1980
	v1991 = l1
	v1996 = v1966
	v2003 = int32(0)
	goto L518
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L510
	} else {
		goto L514
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L510
	} else {
		goto L511
	}
L6:
	;
	v318 = v28 + int32(16)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = int32(1)
	if v328 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	v40 = int32(129)
	if base.Ui32(v40) <= base.Ui32(l2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 == v74+v35 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v43 = v40
	goto L13
L12:
	;
	v43 = l2
	goto L13
L13:
	;
	v48 = v43<<(uint(int32(2))%32) + l0 + int32(32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v51 = l1 << (uint(int32(12)) % 32)
	v52 = v33 + v51
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(-364896016)
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v56
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = v49 + v33 - int32(1)
	goto L16
L15:
	;
	v62 = v56
	goto L16
L16:
	;
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = v62 - v33 + int32(1)
	goto L19
L18:
	;
	v67 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v67
	v70 = v51 | int32(1)
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v70
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3281 = v73
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2 + v35
	v81 = v33 + v74<<(uint(int32(12))%32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v82 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v74 == l1+l2 {
		goto L51
	} else {
		goto L52
	}
L26:
	;
	v87 = v33 + v82 - int32(1)
	goto L28
L27:
	;
	v87 = int32(0)
	goto L28
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v88 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v87 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v91 = v88 + v33
	if v91 == int32(1) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+7)) = v82
	goto L29
L32:
	;
	v109 = int32(129)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(v109) <= base.Ui32(v110) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = v96
	goto L32
L34:
	;
	goto L35
L35:
	;
	v98 = int32(129)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if base.Ui32(v98) <= base.Ui32(v99) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = v98
	goto L38
L37:
	;
	v102 = v99
	goto L38
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v102<<(uint(int32(2))%32)+l0)+32)) = v106
	goto L32
L39:
	;
	v113 = v109
	goto L41
L40:
	;
	v113 = v110
	goto L41
L41:
	;
	v118 = v113<<(uint(int32(2))%32) + l0 + int32(32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v122 = int32(1)
	v123 = l0 - v120 + v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v126 = v124 << (uint(int32(12)) % 32)
	v127 = v123 + v126
	*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = int32(-364896016)
	v131 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v131
	if v119 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v137 = v119 + v123 - v122
	goto L44
L43:
	;
	v137 = v131
	goto L44
L44:
	;
	if v137 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v142 = v137 - v123 + int32(1)
	goto L47
L46:
	;
	v142 = int32(0)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v142
	v145 = v126 | int32(1)
	if v137 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v145
	goto L50
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v145
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3281 = v148
	goto L1
L51:
	;
	v153 = v33 + v74<<(uint(int32(12))%32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	if v154 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v223 != 0 {
		goto L80
	} else {
		goto L81
	}
L54:
	;
	v159 = v33 + v154 - int32(1)
	goto L56
L55:
	;
	v159 = int32(0)
	goto L56
L56:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	if v160 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v159 != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v163 = v160 + v33
	if v163 == int32(1) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+7)) = v154
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v183 = v182 + l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v183
	v185 = int32(129)
	if base.Ui32(v185) <= base.Ui32(v183) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+12)) = v168
	goto L60
L62:
	;
	goto L63
L63:
	;
	v170 = int32(129)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if base.Ui32(v170) <= base.Ui32(v171) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v174 = v170
	goto L66
L65:
	;
	v174 = v171
	goto L66
L66:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v174<<(uint(int32(2))%32)+l0)+32)) = v178
	goto L60
L67:
	;
	v188 = v185
	goto L69
L68:
	;
	v188 = v183
	goto L69
L69:
	;
	v193 = v188<<(uint(int32(2))%32) + l0 + int32(32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = int32(1)
	v198 = l0 - v195 + v197
	v200 = l1 << (uint(int32(12)) % 32)
	v201 = v198 + v200
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = int32(-364896016)
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v205
	if v194 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v211 = v194 + v198 - v197
	goto L72
L71:
	;
	v211 = v205
	goto L72
L72:
	;
	if v211 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v216 = v211 - v198 + int32(1)
	goto L75
L74:
	;
	v216 = int32(0)
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v216
	v219 = v200 | int32(1)
	if v211 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+8)) = v219
	goto L78
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v219
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3281 = v222
	goto L1
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = int64(6860498728)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+16)) = v265
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = int64(0)
	v270 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v258 - v33 + v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v270
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	if v275 != 0 {
		goto L6
	} else {
		goto L94
	}
L80:
	;
	v224 = v223 + v33
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+11))
	if v225 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if l3 != 0 {
		v3281 = int32(0)
		goto L1
	} else {
		goto L92
	}
L83:
	;
	v230 = v33 + v225 - int32(1)
	goto L85
L84:
	;
	v230 = int32(0)
	goto L85
L85:
	;
	if v230 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v224)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = v231
	goto L88
L87:
	;
	goto L88
L88:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v234 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v233 - v234
	if v230 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v241 = v230 - v33 + v234
	goto L91
L90:
	;
	v241 = int32(0)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v241
	v258 = v224 - int32(1)
	goto L79
L92:
	;
	v249 = F_FreePageManagerGetInternal(m, l0, int32(1), v28+int32(16))
	mBase = m.M
	if v249 == int32(0) {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v258 = v33 + v252<<(uint(int32(12))%32)
	goto L79
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = l1
	v278 = int32(129)
	if base.Ui32(v278) <= base.Ui32(l2) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v281 = v278
	goto L97
L96:
	;
	v281 = l2
	goto L97
L97:
	;
	v286 = v281<<(uint(int32(2))%32) + l0 + int32(32)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = int32(1)
	v291 = l0 - v288 + v290
	v293 = l1 << (uint(int32(12)) % 32)
	v294 = v291 + v293
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = int32(-364896016)
	v298 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v298
	if v287 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v304 = v291 + v287 - v290
	goto L100
L99:
	;
	v304 = v298
	goto L100
L100:
	;
	if v304 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v309 = v304 - v291 + int32(1)
	goto L103
L102:
	;
	v309 = int32(0)
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+12)) = v309
	v312 = v293 | int32(1)
	if v304 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+8)) = v312
	goto L106
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = v312
	goto L2
L107:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v514 = v508 + int32(12) + v511<<(uint(int32(3))%32)
	if v511 != 0 {
		goto L155
	} else {
		goto L156
	}
L108:
	;
	v492 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v318)+8)) = uint8(v492)
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v492
	goto L107
L109:
	;
	v335 = int32(1)
	v336 = l0 - v329 + v335
	v339 = v336 + v328 - v335
	if v339 == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	if v343 == int32(430584521) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v346 = int32(1)
	v354 = v339
	v356 = v346
	goto L114
L112:
	;
	v428 = int32(2)
	v430 = v339
	goto L113
L113:
	;
	v437 = int32(0)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if base.Ui32(int32(509)) < base.Ui32(v439) {
		goto L136
	} else {
		goto L137
	}
L114:
	;
	v362 = v354 + int32(12)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v365 = int32(0)
	v368 = v363
	goto L116
L115:
	;
	v428 = v409 + int32(1)
	v430 = v419
	goto L113
L116:
	;
	if base.Ui32(v368) <= base.Ui32(v365) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if base.Ui32(v393) < base.Ui32(v363) {
		goto L129
	} else {
		goto L130
	}
L118:
	;
	goto L117
L119:
	;
	v393 = v365
	goto L118
L120:
	;
	goto L121
L121:
	;
	v379 = int32(1)
	v380 = int32(base.Ui32(v365+v368) >> (uint(v379) % 32))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v362+v380<<(uint(int32(3))%32))))
	v387 = base.B2i32(base.Ui32(l1) < base.Ui32(v386))
	if base.Ui32(l1) < base.Ui32(v386) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v388 = v365
	goto L124
L123:
	;
	v388 = v380 + v379
	goto L124
L124:
	;
	if base.Ui32(l1) < base.Ui32(v386) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v389 = v380
	goto L127
L126:
	;
	v389 = v368
	goto L127
L127:
	;
	if l1 != v386 {
		v365 = v388
		v368 = v389
		goto L116
	} else {
		goto L128
	}
L128:
	;
	v393 = v380
	goto L118
L129:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v362+v393<<(uint(int32(3))%32))))
	v403 = base.B2i32(v401 != l1)
	goto L131
L130:
	;
	v403 = int32(1)
	goto L131
L131:
	;
	if base.Ui32(int32(509)) < base.Ui32(v363) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v409 = v356 + int32(1)
	goto L134
L133:
	;
	v409 = int32(0)
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = v409
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v354+(v393-base.B2i32(v393 != int32(0))&v403)<<(uint(int32(3))%32))+16))
	v419 = v336 - v346 + v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if v420 == int32(430584521) {
		v354 = v419
		v356 = v409
		goto L114
	} else {
		goto L135
	}
L135:
	;
	goto L115
L136:
	;
	v442 = v428
	goto L138
L137:
	;
	v442 = v437
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = v442
	v445 = v430 + int32(12)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	v447 = v437
	v450 = v446
	goto L139
L139:
	;
	if base.Ui32(v450) <= base.Ui32(v447) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v430
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if base.Ui32(v475) < base.Ui32(v481) {
		goto L152
	} else {
		goto L153
	}
L141:
	;
	goto L140
L142:
	;
	v475 = v447
	goto L141
L143:
	;
	goto L144
L144:
	;
	v461 = int32(1)
	v462 = int32(base.Ui32(v447+v450) >> (uint(v461) % 32))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v445+v462<<(uint(int32(3))%32))))
	v469 = base.B2i32(base.Ui32(l1) < base.Ui32(v468))
	if base.Ui32(l1) < base.Ui32(v468) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v470 = v447
	goto L147
L146:
	;
	v470 = v462 + v461
	goto L147
L147:
	;
	if base.Ui32(l1) < base.Ui32(v468) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v471 = v462
	goto L150
L149:
	;
	v471 = v450
	goto L150
L150:
	;
	if l1 != v468 {
		v447 = v470
		v450 = v471
		goto L139
	} else {
		goto L151
	}
L151:
	;
	v475 = v462
	goto L141
L152:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v445+v475<<(uint(int32(3))%32))))
	v488 = base.B2i32(l1 == v486)
	goto L154
L153:
	;
	v488 = int32(0)
	goto L154
L154:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v318)+8)) = uint8(v488)
	goto L107
L155:
	;
	v518 = v514 - int32(8)
	goto L157
L156:
	;
	v518 = int32(0)
	goto L157
L157:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	if base.Ui32(v511) < base.Ui32(v519) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if v518 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L159:
	;
	v778 = v508
	v779 = v514
	v794 = v511
	goto L158
L160:
	;
	goto L161
L161:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v508)+8))
	if v521 == int32(0) {
		v748 = v5
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v748 != 0 {
		goto L215
	} else {
		goto L216
	}
L163:
	;
	v524 = v521 + v33
	v526 = v524 - int32(1)
	if v526 == int32(0) {
		v748 = v5
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v530 = v33 - int32(1)
	__phi535 = v526
	__phi536 = v508
	__phi541 = v524
	__phi549 = v5
	v535 = __phi535
	v536 = __phi536
	v541 = __phi541
	v549 = __phi549
	goto L165
L165:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v541)+3))
	v565 = int32(0)
	v566 = v560
	goto L167
L166:
	;
	v617 = int32(0)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v541+v602<<(uint(int32(3))%32))+23))
	if v621 != 0 {
		goto L185
	} else {
		goto L186
	}
L167:
	;
	if base.Ui32(v566) <= base.Ui32(v565) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	if base.Ui32(v560-int32(1)) <= base.Ui32(v602) {
		goto L180
	} else {
		goto L181
	}
L169:
	;
	goto L168
L170:
	;
	v602 = v565
	goto L169
L171:
	;
	goto L172
L172:
	;
	v588 = int32(1)
	v589 = int32(base.Ui32(v565+v566) >> (uint(v588) % 32))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v541+int32(11)+v589<<(uint(int32(3))%32))))
	v596 = base.B2i32(base.Ui32(v558) < base.Ui32(v595))
	if base.Ui32(v558) < base.Ui32(v595) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v597 = v565
	goto L175
L174:
	;
	v597 = v589 + v588
	goto L175
L175:
	;
	if base.Ui32(v558) < base.Ui32(v595) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v598 = v589
	goto L178
L177:
	;
	v598 = v566
	goto L178
L178:
	;
	if v558 != v595 {
		v565 = v597
		v566 = v598
		goto L167
	} else {
		goto L179
	}
L179:
	;
	v602 = v589
	goto L169
L180:
	;
	v608 = int32(0)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	if v609 == v608 {
		v748 = v608
		goto L162
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	goto L166
L183:
	;
	v612 = int32(1)
	v614 = v609 + v33
	v616 = v614 - v612
	if v616 != 0 {
		__phi535 = v616
		__phi536 = v535
		__phi541 = v614
		__phi549 = v549 + v612
		v535 = __phi535
		v536 = __phi536
		v541 = __phi541
		v549 = __phi549
		goto L165
	} else {
		goto L184
	}
L184:
	;
	v748 = v608
	goto L162
L185:
	;
	v624 = v530 + v621
	goto L187
L186:
	;
	v624 = v617
	goto L187
L187:
	;
	if v549 <= int32(0) {
		v748 = v624
		goto L162
	} else {
		goto L188
	}
L188:
	;
	v628 = v549 & int32(3)
	if v628 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	if base.Ui32(v549) < base.Ui32(int32(4)) {
		v748 = v674
		goto L162
	} else {
		goto L199
	}
L190:
	;
	v669 = v549
	v674 = v624
	goto L189
L191:
	;
	goto L192
L192:
	;
	v635 = v549
	v636 = v617
	v640 = v624
	goto L193
L193:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v640)+16))
	if v656 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v669 = v661
	v674 = v659
	goto L189
L195:
	;
	v659 = v530 + v656
	goto L197
L196:
	;
	v659 = int32(0)
	goto L197
L197:
	;
	v660 = int32(1)
	v661 = v635 - v660
	v663 = v636 + v660
	if v663 != v628 {
		v635 = v661
		v636 = v663
		v640 = v659
		goto L193
	} else {
		goto L198
	}
L198:
	;
	goto L194
L199:
	;
	v696 = v669
	v701 = v674
	goto L200
L200:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v701)+16))
	if v717 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v748 = v732
	goto L162
L202:
	;
	v720 = v530 + v717
	goto L204
L203:
	;
	v720 = int32(0)
	goto L204
L204:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+16))
	if v721 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v724 = v530 + v721
	goto L207
L206:
	;
	v724 = int32(0)
	goto L207
L207:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)+16))
	if v725 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v728 = v530 + v725
	goto L210
L209:
	;
	v728 = int32(0)
	goto L210
L210:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+16))
	if v729 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v732 = v530 + v729
	goto L213
L212:
	;
	v732 = int32(0)
	goto L213
L213:
	;
	if base.Ui32(v696-int32(5)) < base.Ui32(int32(-2)) {
		v696 = v696 - int32(4)
		v701 = v732
		goto L200
	} else {
		goto L214
	}
L214:
	;
	goto L201
L215:
	;
	v767 = v748 + int32(12)
	goto L217
L216:
	;
	v767 = int32(0)
	goto L217
L217:
	;
	v778 = v748
	v779 = v767
	v794 = int32(0)
	goto L158
L218:
	;
	if v779 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L219:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if base.Ui32(v797+v798) < base.Ui32(l1) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v801 = l1 + l2
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v801 - v797
	v804 = int32(0)
	if v779 == v804 {
		v855 = v804
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v858 = int32(1)
	v859 = l0 - v856 + v858
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v863 = v859 + v860<<(uint(int32(12))%32)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)+8))
	if v864 != 0 {
		goto L237
	} else {
		goto L238
	}
L222:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	if base.Ui32(v801) < base.Ui32(v808) {
		v855 = int32(0)
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v810 + (v808 - v797)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v816 = int32(1)
	v817 = l0 - v814 + v816
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	v821 = v817 + v818<<(uint(int32(12))%32)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+8))
	if v822 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v827 = v817 + v822 - v816
	goto L226
L225:
	;
	v827 = int32(0)
	goto L226
L226:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v821)+12))
	if v828 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	if v827 != 0 {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v831 = v817 + v828
	if v831 == int32(1) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v831)+7)) = v822
	goto L227
L230:
	;
	v855 = int32(1)
	goto L221
L231:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v821)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v827)+12)) = v836
	goto L230
L232:
	;
	goto L233
L233:
	;
	v838 = int32(129)
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v821)+4))
	if base.Ui32(v838) <= base.Ui32(v839) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v842 = v838
	goto L236
L235:
	;
	v842 = v839
	goto L236
L236:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v821)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v842<<(uint(int32(2))%32)+l0)+32)) = v846
	goto L230
L237:
	;
	v869 = v859 + v864 - v858
	goto L239
L238:
	;
	v869 = int32(0)
	goto L239
L239:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v863)+12))
	if v870 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	if v869 != 0 {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v873 = v859 + v870
	if v873 == int32(1) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v873)+7)) = v864
	goto L240
L243:
	;
	v891 = int32(129)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if base.Ui32(v891) <= base.Ui32(v892) {
		goto L250
	} else {
		goto L251
	}
L244:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v863)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v869)+12)) = v878
	goto L243
L245:
	;
	goto L246
L246:
	;
	v880 = int32(129)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	if base.Ui32(v880) <= base.Ui32(v881) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v884 = v880
	goto L249
L248:
	;
	v884 = v881
	goto L249
L249:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v863)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v884<<(uint(int32(2))%32)+l0)+32)) = v888
	goto L243
L250:
	;
	v895 = v891
	goto L252
L251:
	;
	v895 = v892
	goto L252
L252:
	;
	v900 = v895<<(uint(int32(2))%32) + l0 + int32(32)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v904 = int32(1)
	v905 = l0 - v902 + v904
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v908 = v906 << (uint(int32(12)) % 32)
	v909 = v905 + v908
	*(*int32)(unsafe.Add(mBase, uint32(v909)+4)) = v892
	*(*int32)(unsafe.Add(mBase, uint32(v909))) = int32(-364896016)
	v913 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v909)+8)) = v913
	if v901 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v919 = v901 + v905 - v904
	goto L255
L254:
	;
	v919 = v913
	goto L255
L255:
	;
	if v919 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v924 = v919 - v905 + int32(1)
	goto L258
L257:
	;
	v924 = int32(0)
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v909)+12)) = v924
	v927 = v908 | int32(1)
	if v919 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v919)+8)) = v927
	goto L261
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v900))) = v927
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v855 == int32(0) {
		v3281 = v930
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	if v943 == int32(1) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v3281 = v930
	goto L1
L264:
	;
	F_FreePageBtreeRemovePage(m, l0, v778)
	mBase = m.M
	goto L263
L265:
	;
	goto L266
L266:
	;
	v948 = v943 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v778)+4)) = v948
	if base.Ui32(v948) <= base.Ui32(v794) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_FreePageBtreeConsolidate(m, l0, v778)
	mBase = m.M
	goto L263
L268:
	;
	v952 = v778 + int32(12)
	v953 = int32(3)
	v955 = v952 + v794<<(uint(v953)%32)
	v961 = F_memmove(m, v955, v955+int32(8), (v948-v794)<<(uint(v953)%32))
	mBase = m.M
	if v794 != 0 {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v963 = l0 - v962
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v952)))
	v972 = v778
	goto L270
L270:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v972)+8))
	if v978 == int32(0) {
		goto L267
	} else {
		goto L272
	}
L271:
	;
	goto L267
L272:
	;
	v981 = v978 + v963
	if v981 == int32(0) {
		goto L267
	} else {
		goto L273
	}
L273:
	;
	v985 = v981 + int32(12)
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	v990 = int32(0)
	v991 = v987
	goto L274
L274:
	;
	if base.Ui32(v991) <= base.Ui32(v990) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	if base.Ui32(v1017) < base.Ui32(v987) {
		goto L287
	} else {
		goto L288
	}
L276:
	;
	goto L275
L277:
	;
	v1017 = v990
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1003 = int32(1)
	v1004 = int32(base.Ui32(v990+v991) >> (uint(v1003) % 32))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v985+v1004<<(uint(int32(3))%32))))
	v1011 = base.B2i32(base.Ui32(v964) < base.Ui32(v1010))
	if base.Ui32(v964) < base.Ui32(v1010) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1012 = v990
	goto L282
L281:
	;
	v1012 = v1004 + v1003
	goto L282
L282:
	;
	if base.Ui32(v964) < base.Ui32(v1010) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1013 = v1004
	goto L285
L284:
	;
	v1013 = v991
	goto L285
L285:
	;
	if v964 != v1010 {
		v990 = v1012
		v991 = v1013
		goto L274
	} else {
		goto L286
	}
L286:
	;
	v1017 = v1004
	goto L276
L287:
	;
	v1023 = int32(0)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v981+v1017<<(uint(int32(3))%32))+16))
	if v1027 != 0 {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	v1034 = int32(-1)
	goto L289
L289:
	;
	v1035 = v1034 + v1017
	*(*int32)(unsafe.Add(mBase, uint32(v985+v1035<<(uint(int32(3))%32)))) = v964
	if v1035 == int32(0) {
		v972 = v981
		goto L270
	} else {
		goto L296
	}
L290:
	;
	v1030 = v963 + v1027
	goto L292
L291:
	;
	v1030 = v1023
	goto L292
L292:
	;
	if v1030 != v972 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1032 = int32(-1)
	goto L295
L294:
	;
	v1032 = v1023
	goto L295
L295:
	;
	v1034 = v1032
	goto L289
L296:
	;
	goto L271
L297:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v1282 != 0 {
		goto L354
	} else {
		goto L355
	}
L298:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	if base.Ui32(l1+l2) < base.Ui32(v1072) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1077 = int32(1)
	v1078 = l0 - v1075 + v1077
	v1081 = v1078 + v1072<<(uint(int32(12))%32)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+8))
	if v1082 != 0 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1087 = v1078 + v1082 - v1077
	goto L302
L301:
	;
	v1087 = int32(0)
	goto L302
L302:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+12))
	if v1090 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1098 = v1072 - l1 + v1089
	if v1087 != 0 {
		goto L307
	} else {
		goto L308
	}
L304:
	;
	v1093 = v1078 + v1090
	if v1093 == int32(1) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+7)) = v1082
	goto L303
L306:
	;
	v1112 = int32(129)
	if base.Ui32(v1112) <= base.Ui32(v1098) {
		goto L313
	} else {
		goto L314
	}
L307:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1087)+12)) = v1099
	goto L306
L308:
	;
	goto L309
L309:
	;
	v1101 = int32(129)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	if base.Ui32(v1101) <= base.Ui32(v1102) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1105 = v1101
	goto L312
L311:
	;
	v1105 = v1102
	goto L312
L312:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1105<<(uint(int32(2))%32)+l0)+32)) = v1109
	goto L306
L313:
	;
	v1115 = v1112
	goto L315
L314:
	;
	v1115 = v1098
	goto L315
L315:
	;
	v1120 = v1115<<(uint(int32(2))%32) + l0 + int32(32)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1120)))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1124 = int32(1)
	v1125 = l0 - v1122 + v1124
	v1127 = l1 << (uint(int32(12)) % 32)
	v1128 = v1125 + v1127
	*(*int32)(unsafe.Add(mBase, uint32(v1128)+4)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v1128))) = int32(-364896016)
	v1132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1128)+8)) = v1132
	if v1121 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1138 = v1121 + v1125 - v1124
	goto L318
L317:
	;
	v1138 = v1132
	goto L318
L318:
	;
	if v1138 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1143 = v1138 - v1125 + int32(1)
	goto L321
L320:
	;
	v1143 = int32(0)
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1128)+12)) = v1143
	v1146 = v1127 | int32(1)
	if v1138 != 0 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+8)) = v1146
	goto L324
L323:
	;
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1120))) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v779)+4)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v779))) = l1
	if v794 != 0 {
		v3281 = v1098
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1152 = l0 - v1151
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	v1163 = v778
	goto L326
L326:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+8))
	if v1179 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v3281 = v1280
	goto L1
L328:
	;
	goto L327
L329:
	;
	v1182 = v1179 + v1152
	if v1182 == int32(0) {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1186 = v1182 + int32(12)
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+4))
	v1193 = int32(0)
	v1194 = v1188
	goto L331
L331:
	;
	if base.Ui32(v1194) <= base.Ui32(v1193) {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	if base.Ui32(v1230) < base.Ui32(v1188) {
		goto L344
	} else {
		goto L345
	}
L333:
	;
	goto L332
L334:
	;
	v1230 = v1193
	goto L333
L335:
	;
	goto L336
L336:
	;
	v1216 = int32(1)
	v1217 = int32(base.Ui32(v1193+v1194) >> (uint(v1216) % 32))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1186+v1217<<(uint(int32(3))%32))))
	v1224 = base.B2i32(base.Ui32(v1153) < base.Ui32(v1223))
	if base.Ui32(v1153) < base.Ui32(v1223) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1225 = v1193
	goto L339
L338:
	;
	v1225 = v1217 + v1216
	goto L339
L339:
	;
	if base.Ui32(v1153) < base.Ui32(v1223) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1226 = v1217
	goto L342
L341:
	;
	v1226 = v1194
	goto L342
L342:
	;
	if v1153 != v1223 {
		v1193 = v1225
		v1194 = v1226
		goto L331
	} else {
		goto L343
	}
L343:
	;
	v1230 = v1217
	goto L333
L344:
	;
	v1235 = int32(0)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1182+v1230<<(uint(int32(3))%32))+16))
	if v1239 != 0 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v1247 = int32(-1)
	goto L346
L346:
	;
	v1248 = v1247 + v1230
	*(*int32)(unsafe.Add(mBase, uint32(v1186+v1248<<(uint(int32(3))%32)))) = v1153
	if v1248 == int32(0) {
		v1163 = v1182
		goto L326
	} else {
		goto L353
	}
L347:
	;
	v1242 = v1152 + v1239
	goto L349
L348:
	;
	v1242 = v1235
	goto L349
L349:
	;
	if v1242 != v1163 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1244 = int32(-1)
	goto L352
L351:
	;
	v1244 = v1235
	goto L352
L352:
	;
	v1247 = v1244
	goto L346
L353:
	;
	goto L328
L354:
	;
	if l3 != 0 {
		v3281 = int32(0)
		goto L1
	} else {
		goto L357
	}
L355:
	;
	v1560 = v508
	v1569 = v511
	v1570 = v519
	goto L356
L356:
	;
	v1575 = v1560 + int32(12)
	v1576 = int32(3)
	v1578 = v1575 + v1569<<(uint(v1576)%32)
	v1580 = v1578 + int32(8)
	v1583 = (v1570 - v1569) << (uint(v1576) % 32)
	if v1580 == v1578 {
		goto L422
	} else {
		goto L423
	}
L357:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v1282) <= base.Ui32(v1284) {
		v1966 = v508
		goto L3
	} else {
		goto L358
	}
L358:
	;
	v1295 = int32(0)
	goto L359
L359:
	;
	v1316 = F_FreePageManagerGetInternal(m, l0, int32(1), v28+int32(12))
	mBase = m.M
	if v1316 == int32(0) {
		goto L4
	} else {
		goto L361
	}
L360:
	;
	v1355 = v28 + int32(16)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+12)) = int32(1)
	if v1365 == int32(0) {
		goto L373
	} else {
		goto L374
	}
L361:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1322 = int32(1)
	v1323 = l0 - v1320 + v1322
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v1326 = v1324 << (uint(int32(12)) % 32)
	v1327 = v1323 + v1326
	*(*int64)(unsafe.Add(mBase, uint32(v1327))) = int64(8225038576)
	v1330 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+8)) = v1330
	if v1319 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1336 = v1319 + v1323 - v1322
	goto L364
L363:
	;
	v1336 = v1330
	goto L364
L364:
	;
	if v1336 != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1341 = v1336 - v1323 + int32(1)
	goto L367
L366:
	;
	v1341 = int32(0)
	goto L367
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+12)) = v1341
	v1344 = v1326 | int32(1)
	if v1336 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+8)) = v1344
	goto L370
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1344
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1348 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1347 + v1348
	v1352 = v1295 + v1348
	if v1352 != v1282-v1284 {
		v1295 = v1352
		goto L359
	} else {
		goto L371
	}
L371:
	;
	goto L360
L372:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v1546 != 0 {
		v1966 = v1545
		goto L3
	} else {
		goto L420
	}
L373:
	;
	v1529 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1355)+8)) = uint8(v1529)
	*(*int32)(unsafe.Add(mBase, uint32(v1355))) = v1529
	goto L372
L374:
	;
	v1372 = int32(1)
	v1373 = l0 - v1366 + v1372
	v1376 = v1373 + v1365 - v1372
	if v1376 == int32(0) {
		goto L373
	} else {
		goto L375
	}
L375:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1376)))
	if v1380 == int32(430584521) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1383 = int32(1)
	v1391 = v1376
	v1393 = v1383
	goto L379
L377:
	;
	v1465 = int32(2)
	v1467 = v1376
	goto L378
L378:
	;
	v1474 = int32(0)
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+4))
	if base.Ui32(int32(509)) < base.Ui32(v1476) {
		goto L401
	} else {
		goto L402
	}
L379:
	;
	v1399 = v1391 + int32(12)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1391)+4))
	v1402 = int32(0)
	v1405 = v1400
	goto L381
L380:
	;
	v1465 = v1446 + int32(1)
	v1467 = v1456
	goto L378
L381:
	;
	if base.Ui32(v1405) <= base.Ui32(v1402) {
		goto L384
	} else {
		goto L385
	}
L382:
	;
	if base.Ui32(v1430) < base.Ui32(v1400) {
		goto L394
	} else {
		goto L395
	}
L383:
	;
	goto L382
L384:
	;
	v1430 = v1402
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1416 = int32(1)
	v1417 = int32(base.Ui32(v1402+v1405) >> (uint(v1416) % 32))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1399+v1417<<(uint(int32(3))%32))))
	v1424 = base.B2i32(base.Ui32(l1) < base.Ui32(v1423))
	if base.Ui32(l1) < base.Ui32(v1423) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1425 = v1402
	goto L389
L388:
	;
	v1425 = v1417 + v1416
	goto L389
L389:
	;
	if base.Ui32(l1) < base.Ui32(v1423) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1426 = v1417
	goto L392
L391:
	;
	v1426 = v1405
	goto L392
L392:
	;
	if l1 != v1423 {
		v1402 = v1425
		v1405 = v1426
		goto L381
	} else {
		goto L393
	}
L393:
	;
	v1430 = v1417
	goto L383
L394:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1399+v1430<<(uint(int32(3))%32))))
	v1440 = base.B2i32(v1438 != l1)
	goto L396
L395:
	;
	v1440 = int32(1)
	goto L396
L396:
	;
	if base.Ui32(int32(509)) < base.Ui32(v1400) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1446 = v1393 + int32(1)
	goto L399
L398:
	;
	v1446 = int32(0)
	goto L399
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+12)) = v1446
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1391+(v1430-base.B2i32(v1430 != int32(0))&v1440)<<(uint(int32(3))%32))+16))
	v1456 = v1373 - v1383 + v1455
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	if v1457 == int32(430584521) {
		v1391 = v1456
		v1393 = v1446
		goto L379
	} else {
		goto L400
	}
L400:
	;
	goto L380
L401:
	;
	v1479 = v1465
	goto L403
L402:
	;
	v1479 = v1474
	goto L403
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+12)) = v1479
	v1482 = v1467 + int32(12)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+4))
	v1484 = v1474
	v1487 = v1483
	goto L404
L404:
	;
	if base.Ui32(v1487) <= base.Ui32(v1484) {
		goto L407
	} else {
		goto L408
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+4)) = v1512
	*(*int32)(unsafe.Add(mBase, uint32(v1355))) = v1467
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+4))
	if base.Ui32(v1512) < base.Ui32(v1518) {
		goto L417
	} else {
		goto L418
	}
L406:
	;
	goto L405
L407:
	;
	v1512 = v1484
	goto L406
L408:
	;
	goto L409
L409:
	;
	v1498 = int32(1)
	v1499 = int32(base.Ui32(v1484+v1487) >> (uint(v1498) % 32))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1482+v1499<<(uint(int32(3))%32))))
	v1506 = base.B2i32(base.Ui32(l1) < base.Ui32(v1505))
	if base.Ui32(l1) < base.Ui32(v1505) {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1507 = v1484
	goto L412
L411:
	;
	v1507 = v1499 + v1498
	goto L412
L412:
	;
	if base.Ui32(l1) < base.Ui32(v1505) {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1508 = v1499
	goto L415
L414:
	;
	v1508 = v1487
	goto L415
L415:
	;
	if l1 != v1505 {
		v1484 = v1507
		v1487 = v1508
		goto L404
	} else {
		goto L416
	}
L416:
	;
	v1512 = v1499
	goto L406
L417:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1482+v1512<<(uint(int32(3))%32))))
	v1525 = base.B2i32(l1 == v1523)
	goto L419
L418:
	;
	v1525 = int32(0)
	goto L419
L419:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1355)+8)) = uint8(v1525)
	goto L372
L420:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+4))
	v1560 = v1545
	v1569 = v1547
	v1570 = v1548
	goto L356
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1578)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1578))) = l1
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+4)) = v1730 + int32(1)
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1569 == int32(0) {
		goto L467
	} else {
		goto L468
	}
L422:
	;
	goto L421
L423:
	;
	v1587 = v1580 + v1583
	if base.Ui32(v1578-v1587) <= base.Ui32(int32(0)-v1583<<(uint(int32(1))%32)) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1594 = F___memcpy(m, v1580, v1578, v1583)
	mBase = m.M
	goto L421
L425:
	;
	goto L426
L426:
	;
	v1597 = (v1580 ^ v1578) & int32(3)
	if base.Ui32(v1580) < base.Ui32(v1578) {
		goto L429
	} else {
		goto L430
	}
L427:
	;
	if v1699 == int32(0) {
		goto L422
	} else {
		goto L463
	}
L428:
	;
	if base.Ui32(v1677) <= base.Ui32(int32(3)) {
		v1698 = v1676
		v1699 = v1677
		v1700 = v1678
		goto L427
	} else {
		goto L459
	}
L429:
	;
	if v1597 != 0 {
		goto L432
	} else {
		goto L433
	}
L430:
	;
	goto L431
L431:
	;
	if v1597 != 0 {
		v1659 = v1583
		goto L442
	} else {
		goto L443
	}
L432:
	;
	v1698 = v1578
	v1699 = v1583
	v1700 = v1580
	goto L427
L433:
	;
	goto L434
L434:
	;
	if v1580&int32(3) == int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1676 = v1578
	v1677 = v1583
	v1678 = v1580
	goto L428
L436:
	;
	goto L437
L437:
	;
	v1604 = v1578
	v1605 = v1583
	v1606 = v1580
	goto L438
L438:
	;
	if v1605 == int32(0) {
		goto L422
	} else {
		goto L440
	}
L439:
	;
	v1676 = v1613
	v1677 = v1615
	v1678 = v1617
	goto L428
L440:
	;
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1606))) = uint8(v1610)
	v1612 = int32(1)
	v1613 = v1604 + v1612
	v1615 = v1605 - v1612
	v1617 = v1606 + v1612
	if v1617&int32(3) != 0 {
		v1604 = v1613
		v1605 = v1615
		v1606 = v1617
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	if v1659 == int32(0) {
		goto L422
	} else {
		goto L455
	}
L443:
	;
	if v1587&int32(3) != 0 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1624 = v1583
	goto L447
L445:
	;
	v1639 = v1583
	goto L446
L446:
	;
	if base.Ui32(v1639) <= base.Ui32(int32(3)) {
		v1659 = v1639
		goto L442
	} else {
		goto L451
	}
L447:
	;
	if v1624 == int32(0) {
		goto L422
	} else {
		goto L449
	}
L448:
	;
	v1639 = v1630
	goto L446
L449:
	;
	v1630 = v1624 - int32(1)
	v1631 = v1580 + v1630
	v1633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1578+v1630))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1631))) = uint8(v1633)
	if v1631&int32(3) != 0 {
		v1624 = v1630
		goto L447
	} else {
		goto L450
	}
L450:
	;
	goto L448
L451:
	;
	v1646 = v1639
	goto L452
L452:
	;
	v1650 = v1646 - int32(4)
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1578+v1650)))
	*(*int32)(unsafe.Add(mBase, uint32(v1580+v1650))) = v1653
	if base.Ui32(int32(3)) < base.Ui32(v1650) {
		v1646 = v1650
		goto L452
	} else {
		goto L454
	}
L453:
	;
	v1659 = v1650
	goto L442
L454:
	;
	goto L453
L455:
	;
	v1666 = v1659
	goto L456
L456:
	;
	v1670 = v1666 - int32(1)
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1578+v1670))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1580+v1670))) = uint8(v1673)
	if v1670 != 0 {
		v1666 = v1670
		goto L456
	} else {
		goto L458
	}
L457:
	;
	goto L422
L458:
	;
	goto L457
L459:
	;
	v1683 = v1676
	v1684 = v1677
	v1685 = v1678
	goto L460
L460:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1683)))
	*(*int32)(unsafe.Add(mBase, uint32(v1685))) = v1687
	v1689 = int32(4)
	v1690 = v1683 + v1689
	v1692 = v1685 + v1689
	v1694 = v1684 - v1689
	if base.Ui32(int32(3)) < base.Ui32(v1694) {
		v1683 = v1690
		v1684 = v1694
		v1685 = v1692
		goto L460
	} else {
		goto L462
	}
L461:
	;
	v1698 = v1690
	v1699 = v1694
	v1700 = v1692
	goto L427
L462:
	;
	goto L461
L463:
	;
	v1705 = v1698
	v1706 = v1699
	v1707 = v1700
	goto L464
L464:
	;
	v1709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1705))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1707))) = uint8(v1709)
	v1711 = int32(1)
	v1716 = v1706 - v1711
	if v1716 != 0 {
		v1705 = v1705 + v1711
		v1706 = v1716
		v1707 = v1707 + v1711
		goto L464
	} else {
		goto L466
	}
L465:
	;
	goto L422
L466:
	;
	goto L465
L467:
	;
	v1737 = l0 - v1734
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1575)))
	v1750 = v1560
	goto L470
L468:
	;
	v1870 = v1734
	goto L469
L469:
	;
	v1891 = int32(129)
	if base.Ui32(v1891) <= base.Ui32(l2) {
		goto L498
	} else {
		goto L499
	}
L470:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1750)+8))
	if v1764 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1870 = v1865
	goto L469
L472:
	;
	goto L471
L473:
	;
	v1767 = v1764 + v1737
	if v1767 == int32(0) {
		goto L472
	} else {
		goto L474
	}
L474:
	;
	v1771 = v1767 + int32(12)
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1767)+4))
	v1778 = int32(0)
	v1779 = v1773
	goto L475
L475:
	;
	if base.Ui32(v1779) <= base.Ui32(v1778) {
		goto L478
	} else {
		goto L479
	}
L476:
	;
	if base.Ui32(v1815) < base.Ui32(v1773) {
		goto L488
	} else {
		goto L489
	}
L477:
	;
	goto L476
L478:
	;
	v1815 = v1778
	goto L477
L479:
	;
	goto L480
L480:
	;
	v1801 = int32(1)
	v1802 = int32(base.Ui32(v1778+v1779) >> (uint(v1801) % 32))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1771+v1802<<(uint(int32(3))%32))))
	v1809 = base.B2i32(base.Ui32(v1738) < base.Ui32(v1808))
	if base.Ui32(v1738) < base.Ui32(v1808) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v1810 = v1778
	goto L483
L482:
	;
	v1810 = v1802 + v1801
	goto L483
L483:
	;
	if base.Ui32(v1738) < base.Ui32(v1808) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1811 = v1802
	goto L486
L485:
	;
	v1811 = v1779
	goto L486
L486:
	;
	if v1738 != v1808 {
		v1778 = v1810
		v1779 = v1811
		goto L475
	} else {
		goto L487
	}
L487:
	;
	v1815 = v1802
	goto L477
L488:
	;
	v1820 = int32(0)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1767+v1815<<(uint(int32(3))%32))+16))
	if v1824 != 0 {
		goto L491
	} else {
		goto L492
	}
L489:
	;
	v1832 = int32(-1)
	goto L490
L490:
	;
	v1833 = v1832 + v1815
	*(*int32)(unsafe.Add(mBase, uint32(v1771+v1833<<(uint(int32(3))%32)))) = v1738
	if v1833 == int32(0) {
		v1750 = v1767
		goto L470
	} else {
		goto L497
	}
L491:
	;
	v1827 = v1737 + v1824
	goto L493
L492:
	;
	v1827 = v1820
	goto L493
L493:
	;
	if v1827 != v1750 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v1829 = int32(-1)
	goto L496
L495:
	;
	v1829 = v1820
	goto L496
L496:
	;
	v1832 = v1829
	goto L490
L497:
	;
	goto L472
L498:
	;
	v1894 = v1891
	goto L500
L499:
	;
	v1894 = l2
	goto L500
L500:
	;
	v1899 = v1894<<(uint(int32(2))%32) + l0 + int32(32)
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1899)))
	v1902 = int32(1)
	v1903 = l0 - v1870 + v1902
	v1905 = l1 << (uint(int32(12)) % 32)
	v1906 = v1903 + v1905
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1906))) = int32(-364896016)
	v1910 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+8)) = v1910
	if v1900 != 0 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1916 = v1900 + v1903 - v1902
	goto L503
L502:
	;
	v1916 = v1910
	goto L503
L503:
	;
	if v1916 != 0 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v1921 = v1916 - v1903 + int32(1)
	goto L506
L505:
	;
	v1921 = int32(0)
	goto L506
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+12)) = v1921
	v1924 = v1905 | int32(1)
	if v1916 != 0 {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1916)+8)) = v1924
	goto L509
L508:
	;
	goto L509
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1899))) = v1924
	goto L2
L510:
	;
	return int32(0)
L511:
	;
	F_errmsg_internal(m, int32(82329), int32(0))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L510
	} else {
		goto L512
	}
L512:
	;
	F_errfinish(m, int32(490378), int32(1534), int32(307087))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L510
	} else {
		goto L513
	}
L513:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L514:
	;
	F_errmsg_internal(m, int32(82329), int32(0))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L510
	} else {
		goto L515
	}
L515:
	;
	F_errfinish(m, int32(490378), int32(1689), int32(307087))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L510
	} else {
		goto L516
	}
L516:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L517:
	;
	v3215 = int32(129)
	if base.Ui32(v3215) <= base.Ui32(l2) {
		goto L840
	} else {
		goto L841
	}
L518:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+8))
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2013 = int32(1)
	v2014 = l0 - v2011 + v2013
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2018 = v2014 + v2015 - v2013
	if v2015 != 0 {
		goto L520
	} else {
		goto L521
	}
L519:
	;
	v2871 = v2809 + int32(12)
	v2878 = int32(0)
	v2880 = v2867
	goto L753
L520:
	;
	v2020 = v2018
	goto L522
L521:
	;
	v2020 = int32(0)
	goto L522
L522:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+12))
	if v2021 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2026 = v2014 + v2021 - int32(1)
	goto L525
L524:
	;
	v2026 = int32(0)
	goto L525
L525:
	;
	if v2026 != 0 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2026)+8)) = v2027
	goto L528
L527:
	;
	goto L528
L528:
	;
	v2030 = v2020 + int32(12)
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2032 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2031 - v2032
	if v2026 != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2039 = v2026 - v2014 + v2032
	goto L531
L530:
	;
	v2039 = int32(0)
	goto L531
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2039
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1996)))
	*(*int32)(unsafe.Add(mBase, uint32(v2020))) = v2041
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+4))
	v2045 = int32(base.Ui32(v2043) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+4)) = v2045
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+8)) = v2047
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+4))
	v2050 = v2049 - v2045
	*(*int32)(unsafe.Add(mBase, uint32(v1996)+4)) = v2050
	v2053 = v1996 + int32(12)
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+4))
	v2056 = v2054 << (uint(int32(3)) % 32)
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v1996)))
	if v2057 == int32(-1729435864) {
		goto L533
	} else {
		goto L534
	}
L532:
	;
	v2140 = v2010 + (v33 - v1980)
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+12))
	if base.Ui32(v1991) < base.Ui32(v2141) {
		goto L551
	} else {
		goto L552
	}
L533:
	;
	if v2056 != 0 {
		goto L537
	} else {
		goto L538
	}
L534:
	;
	goto L535
L535:
	;
	if v2056 != 0 {
		goto L541
	} else {
		goto L542
	}
L536:
	;
	goto L532
L537:
	;
	v2063 = F__emscripten_memcpy_bulkmem(m, v2030, v2053+v2050<<(uint(int32(3))%32), v2056)
	mBase = m.M
	goto L539
L538:
	;
	goto L539
L539:
	;
	goto L536
L540:
	;
	if v2054 == int32(0) {
		goto L532
	} else {
		goto L544
	}
L541:
	;
	v2068 = F__emscripten_memcpy_bulkmem(m, v2030, v2053+v2050<<(uint(int32(3))%32), v2056)
	mBase = m.M
	goto L543
L542:
	;
	goto L543
L543:
	;
	goto L540
L544:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2075 = l0 - v2074
	v2082 = int32(0)
	goto L545
L545:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2020+int32(16)+v2082<<(uint(int32(3))%32))))
	if v2106 != 0 {
		goto L547
	} else {
		goto L548
	}
L546:
	;
	goto L532
L547:
	;
	v2109 = v2075 + v2106
	goto L549
L548:
	;
	v2109 = int32(0)
	goto L549
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+8)) = v2020 - v2075
	v2112 = v2082 + int32(1)
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+4))
	if base.Ui32(v2112) < base.Ui32(v2113) {
		v2082 = v2112
		goto L545
	} else {
		goto L550
	}
L550:
	;
	goto L546
L551:
	;
	v2143 = v1996
	goto L553
L552:
	;
	v2143 = v2020
	goto L553
L553:
	;
	v2145 = v2143 + int32(12)
	v2146 = int32(0)
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	if v2003 == v2146 {
		goto L555
	} else {
		goto L556
	}
L554:
	;
	if v2010 != 0 {
		goto L734
	} else {
		goto L735
	}
L555:
	;
	v2154 = v2146
	v2157 = v2147
	goto L558
L556:
	;
	goto L557
L557:
	;
	v2467 = v2146
	v2468 = v2147
	goto L646
L558:
	;
	if base.Ui32(v2157) <= base.Ui32(v2154) {
		goto L561
	} else {
		goto L562
	}
L559:
	;
	v2197 = int32(3)
	v2199 = v2145 + v2195<<(uint(v2197)%32)
	v2202 = v2145 + v2192<<(uint(v2197)%32)
	v2205 = (v2147 - v2192) << (uint(v2197) % 32)
	if v2199 == v2202 {
		goto L572
	} else {
		goto L573
	}
L560:
	;
	goto L559
L561:
	;
	v2192 = v2154
	v2195 = v2154 + int32(1)
	goto L560
L562:
	;
	goto L563
L563:
	;
	v2179 = int32(1)
	v2180 = int32(base.Ui32(v2154+v2157) >> (uint(v2179) % 32))
	v2182 = v2180 + v2179
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2145+v2180<<(uint(int32(3))%32))))
	v2187 = base.B2i32(base.Ui32(v1991) < base.Ui32(v2186))
	if base.Ui32(v1991) < base.Ui32(v2186) {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v2188 = v2154
	goto L566
L565:
	;
	v2188 = v2182
	goto L566
L566:
	;
	if base.Ui32(v1991) < base.Ui32(v2186) {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v2189 = v2180
	goto L569
L568:
	;
	v2189 = v2157
	goto L569
L569:
	;
	if v1991 != v2186 {
		v2154 = v2188
		v2157 = v2189
		goto L558
	} else {
		goto L570
	}
L570:
	;
	v2192 = v2180
	v2195 = v2182
	goto L560
L571:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2202)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v2202))) = v1991
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+4)) = v2352 + int32(1)
	if v1996 != v2143 {
		goto L554
	} else {
		goto L617
	}
L572:
	;
	goto L571
L573:
	;
	v2209 = v2199 + v2205
	if base.Ui32(v2202-v2209) <= base.Ui32(int32(0)-v2205<<(uint(int32(1))%32)) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	v2216 = F___memcpy(m, v2199, v2202, v2205)
	mBase = m.M
	goto L571
L575:
	;
	goto L576
L576:
	;
	v2219 = (v2199 ^ v2202) & int32(3)
	if base.Ui32(v2199) < base.Ui32(v2202) {
		goto L579
	} else {
		goto L580
	}
L577:
	;
	if v2321 == int32(0) {
		goto L572
	} else {
		goto L613
	}
L578:
	;
	if base.Ui32(v2299) <= base.Ui32(int32(3)) {
		v2320 = v2298
		v2321 = v2299
		v2322 = v2300
		goto L577
	} else {
		goto L609
	}
L579:
	;
	if v2219 != 0 {
		goto L582
	} else {
		goto L583
	}
L580:
	;
	goto L581
L581:
	;
	if v2219 != 0 {
		v2281 = v2205
		goto L592
	} else {
		goto L593
	}
L582:
	;
	v2320 = v2202
	v2321 = v2205
	v2322 = v2199
	goto L577
L583:
	;
	goto L584
L584:
	;
	if v2199&int32(3) == int32(0) {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v2298 = v2202
	v2299 = v2205
	v2300 = v2199
	goto L578
L586:
	;
	goto L587
L587:
	;
	v2226 = v2202
	v2227 = v2205
	v2228 = v2199
	goto L588
L588:
	;
	if v2227 == int32(0) {
		goto L572
	} else {
		goto L590
	}
L589:
	;
	v2298 = v2235
	v2299 = v2237
	v2300 = v2239
	goto L578
L590:
	;
	v2232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2228))) = uint8(v2232)
	v2234 = int32(1)
	v2235 = v2226 + v2234
	v2237 = v2227 - v2234
	v2239 = v2228 + v2234
	if v2239&int32(3) != 0 {
		v2226 = v2235
		v2227 = v2237
		v2228 = v2239
		goto L588
	} else {
		goto L591
	}
L591:
	;
	goto L589
L592:
	;
	if v2281 == int32(0) {
		goto L572
	} else {
		goto L605
	}
L593:
	;
	if v2209&int32(3) != 0 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v2246 = v2205
	goto L597
L595:
	;
	v2261 = v2205
	goto L596
L596:
	;
	if base.Ui32(v2261) <= base.Ui32(int32(3)) {
		v2281 = v2261
		goto L592
	} else {
		goto L601
	}
L597:
	;
	if v2246 == int32(0) {
		goto L572
	} else {
		goto L599
	}
L598:
	;
	v2261 = v2252
	goto L596
L599:
	;
	v2252 = v2246 - int32(1)
	v2253 = v2199 + v2252
	v2255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2202+v2252))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2253))) = uint8(v2255)
	if v2253&int32(3) != 0 {
		v2246 = v2252
		goto L597
	} else {
		goto L600
	}
L600:
	;
	goto L598
L601:
	;
	v2268 = v2261
	goto L602
L602:
	;
	v2272 = v2268 - int32(4)
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2202+v2272)))
	*(*int32)(unsafe.Add(mBase, uint32(v2199+v2272))) = v2275
	if base.Ui32(int32(3)) < base.Ui32(v2272) {
		v2268 = v2272
		goto L602
	} else {
		goto L604
	}
L603:
	;
	v2281 = v2272
	goto L592
L604:
	;
	goto L603
L605:
	;
	v2288 = v2281
	goto L606
L606:
	;
	v2292 = v2288 - int32(1)
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2202+v2292))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2199+v2292))) = uint8(v2295)
	if v2292 != 0 {
		v2288 = v2292
		goto L606
	} else {
		goto L608
	}
L607:
	;
	goto L572
L608:
	;
	goto L607
L609:
	;
	v2305 = v2298
	v2306 = v2299
	v2307 = v2300
	goto L610
L610:
	;
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v2305)))
	*(*int32)(unsafe.Add(mBase, uint32(v2307))) = v2309
	v2311 = int32(4)
	v2312 = v2305 + v2311
	v2314 = v2307 + v2311
	v2316 = v2306 - v2311
	if base.Ui32(int32(3)) < base.Ui32(v2316) {
		v2305 = v2312
		v2306 = v2316
		v2307 = v2314
		goto L610
	} else {
		goto L612
	}
L611:
	;
	v2320 = v2312
	v2321 = v2316
	v2322 = v2314
	goto L577
L612:
	;
	goto L611
L613:
	;
	v2327 = v2320
	v2328 = v2321
	v2329 = v2322
	goto L614
L614:
	;
	v2331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2327))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2329))) = uint8(v2331)
	v2333 = int32(1)
	v2338 = v2328 - v2333
	if v2338 != 0 {
		v2327 = v2327 + v2333
		v2328 = v2338
		v2329 = v2329 + v2333
		goto L614
	} else {
		goto L616
	}
L615:
	;
	goto L572
L616:
	;
	goto L615
L617:
	;
	if v2192 != 0 {
		goto L554
	} else {
		goto L618
	}
L618:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2360 = v1981 - v2357 + int32(1)
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2053)))
	v2375 = v1996
	goto L619
L619:
	;
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2375)+8))
	if v2387 == int32(0) {
		goto L554
	} else {
		goto L621
	}
L620:
	;
	goto L554
L621:
	;
	v2390 = v2387 + v2360
	if v2390 == int32(0) {
		goto L554
	} else {
		goto L622
	}
L622:
	;
	v2394 = v2390 + int32(12)
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2390)+4))
	v2401 = int32(0)
	v2402 = v2396
	goto L623
L623:
	;
	if base.Ui32(v2402) <= base.Ui32(v2401) {
		goto L626
	} else {
		goto L627
	}
L624:
	;
	if base.Ui32(v2438) < base.Ui32(v2396) {
		goto L636
	} else {
		goto L637
	}
L625:
	;
	goto L624
L626:
	;
	v2438 = v2401
	goto L625
L627:
	;
	goto L628
L628:
	;
	v2424 = int32(1)
	v2425 = int32(base.Ui32(v2401+v2402) >> (uint(v2424) % 32))
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2394+v2425<<(uint(int32(3))%32))))
	v2432 = base.B2i32(base.Ui32(v2361) < base.Ui32(v2431))
	if base.Ui32(v2361) < base.Ui32(v2431) {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	v2433 = v2401
	goto L631
L630:
	;
	v2433 = v2425 + v2424
	goto L631
L631:
	;
	if base.Ui32(v2361) < base.Ui32(v2431) {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v2434 = v2425
	goto L634
L633:
	;
	v2434 = v2402
	goto L634
L634:
	;
	if v2361 != v2431 {
		v2401 = v2433
		v2402 = v2434
		goto L623
	} else {
		goto L635
	}
L635:
	;
	v2438 = v2425
	goto L625
L636:
	;
	v2443 = int32(0)
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v2390+v2438<<(uint(int32(3))%32))+16))
	if v2447 != 0 {
		goto L639
	} else {
		goto L640
	}
L637:
	;
	v2455 = int32(-1)
	goto L638
L638:
	;
	v2456 = v2455 + v2438
	*(*int32)(unsafe.Add(mBase, uint32(v2394+v2456<<(uint(int32(3))%32)))) = v2361
	if v2456 == int32(0) {
		v2375 = v2390
		goto L619
	} else {
		goto L645
	}
L639:
	;
	v2450 = v2360 + v2447
	goto L641
L640:
	;
	v2450 = v2443
	goto L641
L641:
	;
	if v2450 != v2375 {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	v2452 = int32(-1)
	goto L644
L643:
	;
	v2452 = v2443
	goto L644
L644:
	;
	v2455 = v2452
	goto L638
L645:
	;
	goto L620
L646:
	;
	if base.Ui32(v2468) <= base.Ui32(v2467) {
		goto L649
	} else {
		goto L650
	}
L647:
	;
	v2510 = int32(3)
	v2512 = v2145 + v2508<<(uint(v2510)%32)
	v2515 = v2145 + v2506<<(uint(v2510)%32)
	v2518 = (v2147 - v2506) << (uint(v2510) % 32)
	if v2512 == v2515 {
		goto L660
	} else {
		goto L661
	}
L648:
	;
	goto L647
L649:
	;
	v2506 = v2467
	v2508 = v2467 + int32(1)
	goto L648
L650:
	;
	goto L651
L651:
	;
	v2492 = int32(1)
	v2493 = int32(base.Ui32(v2467+v2468) >> (uint(v2492) % 32))
	v2495 = v2493 + v2492
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2145+v2493<<(uint(int32(3))%32))))
	v2500 = base.B2i32(base.Ui32(v1991) < base.Ui32(v2499))
	if base.Ui32(v1991) < base.Ui32(v2499) {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v2501 = v2467
	goto L654
L653:
	;
	v2501 = v2495
	goto L654
L654:
	;
	if base.Ui32(v1991) < base.Ui32(v2499) {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v2502 = v2493
	goto L657
L656:
	;
	v2502 = v2468
	goto L657
L657:
	;
	if v1991 != v2499 {
		v2467 = v2501
		v2468 = v2502
		goto L646
	} else {
		goto L658
	}
L658:
	;
	v2506 = v2493
	v2508 = v2495
	goto L648
L659:
	;
	v2664 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2515)+4)) = v2003 - v33 + v2664
	*(*int32)(unsafe.Add(mBase, uint32(v2515))) = v1991
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+4)) = v2668 + v2664
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+8)) = v2143 - v33 + v2664
	if v1996 != v2143 {
		goto L554
	} else {
		goto L705
	}
L660:
	;
	goto L659
L661:
	;
	v2522 = v2512 + v2518
	if base.Ui32(v2515-v2522) <= base.Ui32(int32(0)-v2518<<(uint(int32(1))%32)) {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v2529 = F___memcpy(m, v2512, v2515, v2518)
	mBase = m.M
	goto L659
L663:
	;
	goto L664
L664:
	;
	v2532 = (v2512 ^ v2515) & int32(3)
	if base.Ui32(v2512) < base.Ui32(v2515) {
		goto L667
	} else {
		goto L668
	}
L665:
	;
	if v2634 == int32(0) {
		goto L660
	} else {
		goto L701
	}
L666:
	;
	if base.Ui32(v2612) <= base.Ui32(int32(3)) {
		v2633 = v2611
		v2634 = v2612
		v2635 = v2613
		goto L665
	} else {
		goto L697
	}
L667:
	;
	if v2532 != 0 {
		goto L670
	} else {
		goto L671
	}
L668:
	;
	goto L669
L669:
	;
	if v2532 != 0 {
		v2594 = v2518
		goto L680
	} else {
		goto L681
	}
L670:
	;
	v2633 = v2515
	v2634 = v2518
	v2635 = v2512
	goto L665
L671:
	;
	goto L672
L672:
	;
	if v2512&int32(3) == int32(0) {
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v2611 = v2515
	v2612 = v2518
	v2613 = v2512
	goto L666
L674:
	;
	goto L675
L675:
	;
	v2539 = v2515
	v2540 = v2518
	v2541 = v2512
	goto L676
L676:
	;
	if v2540 == int32(0) {
		goto L660
	} else {
		goto L678
	}
L677:
	;
	v2611 = v2548
	v2612 = v2550
	v2613 = v2552
	goto L666
L678:
	;
	v2545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2539))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2541))) = uint8(v2545)
	v2547 = int32(1)
	v2548 = v2539 + v2547
	v2550 = v2540 - v2547
	v2552 = v2541 + v2547
	if v2552&int32(3) != 0 {
		v2539 = v2548
		v2540 = v2550
		v2541 = v2552
		goto L676
	} else {
		goto L679
	}
L679:
	;
	goto L677
L680:
	;
	if v2594 == int32(0) {
		goto L660
	} else {
		goto L693
	}
L681:
	;
	if v2522&int32(3) != 0 {
		goto L682
	} else {
		goto L683
	}
L682:
	;
	v2559 = v2518
	goto L685
L683:
	;
	v2574 = v2518
	goto L684
L684:
	;
	if base.Ui32(v2574) <= base.Ui32(int32(3)) {
		v2594 = v2574
		goto L680
	} else {
		goto L689
	}
L685:
	;
	if v2559 == int32(0) {
		goto L660
	} else {
		goto L687
	}
L686:
	;
	v2574 = v2565
	goto L684
L687:
	;
	v2565 = v2559 - int32(1)
	v2566 = v2512 + v2565
	v2568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2515+v2565))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2566))) = uint8(v2568)
	if v2566&int32(3) != 0 {
		v2559 = v2565
		goto L685
	} else {
		goto L688
	}
L688:
	;
	goto L686
L689:
	;
	v2581 = v2574
	goto L690
L690:
	;
	v2585 = v2581 - int32(4)
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v2515+v2585)))
	*(*int32)(unsafe.Add(mBase, uint32(v2512+v2585))) = v2588
	if base.Ui32(int32(3)) < base.Ui32(v2585) {
		v2581 = v2585
		goto L690
	} else {
		goto L692
	}
L691:
	;
	v2594 = v2585
	goto L680
L692:
	;
	goto L691
L693:
	;
	v2601 = v2594
	goto L694
L694:
	;
	v2605 = v2601 - int32(1)
	v2608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2515+v2605))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2512+v2605))) = uint8(v2608)
	if v2605 != 0 {
		v2601 = v2605
		goto L694
	} else {
		goto L696
	}
L695:
	;
	goto L660
L696:
	;
	goto L695
L697:
	;
	v2618 = v2611
	v2619 = v2612
	v2620 = v2613
	goto L698
L698:
	;
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2618)))
	*(*int32)(unsafe.Add(mBase, uint32(v2620))) = v2622
	v2624 = int32(4)
	v2625 = v2618 + v2624
	v2627 = v2620 + v2624
	v2629 = v2619 - v2624
	if base.Ui32(int32(3)) < base.Ui32(v2629) {
		v2618 = v2625
		v2619 = v2629
		v2620 = v2627
		goto L698
	} else {
		goto L700
	}
L699:
	;
	v2633 = v2625
	v2634 = v2629
	v2635 = v2627
	goto L665
L700:
	;
	goto L699
L701:
	;
	v2640 = v2633
	v2641 = v2634
	v2642 = v2635
	goto L702
L702:
	;
	v2644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2640))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2642))) = uint8(v2644)
	v2646 = int32(1)
	v2651 = v2641 - v2646
	if v2651 != 0 {
		v2640 = v2640 + v2646
		v2641 = v2651
		v2642 = v2642 + v2646
		goto L702
	} else {
		goto L704
	}
L703:
	;
	goto L660
L704:
	;
	goto L703
L705:
	;
	if v2506 != 0 {
		goto L554
	} else {
		goto L706
	}
L706:
	;
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2680 = v1981 - v2677 + int32(1)
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(v2053)))
	v2695 = v1996
	goto L707
L707:
	;
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(v2695)+8))
	if v2707 == int32(0) {
		goto L554
	} else {
		goto L709
	}
L708:
	;
	goto L554
L709:
	;
	v2710 = v2707 + v2680
	if v2710 == int32(0) {
		goto L554
	} else {
		goto L710
	}
L710:
	;
	v2714 = v2710 + int32(12)
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2710)+4))
	v2721 = int32(0)
	v2722 = v2716
	goto L711
L711:
	;
	if base.Ui32(v2722) <= base.Ui32(v2721) {
		goto L714
	} else {
		goto L715
	}
L712:
	;
	if base.Ui32(v2758) < base.Ui32(v2716) {
		goto L724
	} else {
		goto L725
	}
L713:
	;
	goto L712
L714:
	;
	v2758 = v2721
	goto L713
L715:
	;
	goto L716
L716:
	;
	v2744 = int32(1)
	v2745 = int32(base.Ui32(v2721+v2722) >> (uint(v2744) % 32))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2714+v2745<<(uint(int32(3))%32))))
	v2752 = base.B2i32(base.Ui32(v2681) < base.Ui32(v2751))
	if base.Ui32(v2681) < base.Ui32(v2751) {
		goto L717
	} else {
		goto L718
	}
L717:
	;
	v2753 = v2721
	goto L719
L718:
	;
	v2753 = v2745 + v2744
	goto L719
L719:
	;
	if base.Ui32(v2681) < base.Ui32(v2751) {
		goto L720
	} else {
		goto L721
	}
L720:
	;
	v2754 = v2745
	goto L722
L721:
	;
	v2754 = v2722
	goto L722
L722:
	;
	if v2681 != v2751 {
		v2721 = v2753
		v2722 = v2754
		goto L711
	} else {
		goto L723
	}
L723:
	;
	v2758 = v2745
	goto L713
L724:
	;
	v2763 = int32(0)
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2710+v2758<<(uint(int32(3))%32))+16))
	if v2767 != 0 {
		goto L727
	} else {
		goto L728
	}
L725:
	;
	v2775 = int32(-1)
	goto L726
L726:
	;
	v2776 = v2775 + v2758
	*(*int32)(unsafe.Add(mBase, uint32(v2714+v2776<<(uint(int32(3))%32)))) = v2681
	if v2776 == int32(0) {
		v2695 = v2710
		goto L707
	} else {
		goto L733
	}
L727:
	;
	v2770 = v2680 + v2767
	goto L729
L728:
	;
	v2770 = v2763
	goto L729
L729:
	;
	if v2770 != v2695 {
		goto L730
	} else {
		goto L731
	}
L730:
	;
	v2772 = int32(-1)
	goto L732
L731:
	;
	v2772 = v2763
	goto L732
L732:
	;
	v2775 = v2772
	goto L726
L733:
	;
	goto L708
L734:
	;
	v2809 = v2140
	goto L736
L735:
	;
	v2809 = int32(0)
	goto L736
L736:
	;
	if v2809 == int32(0) {
		goto L737
	} else {
		goto L738
	}
L737:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2814 = int32(1)
	v2815 = l0 - v2812 + v2814
	v2816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2816 != 0 {
		goto L740
	} else {
		goto L741
	}
L738:
	;
	goto L739
L739:
	;
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2030)))
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+4))
	if base.Ui32(int32(509)) < base.Ui32(v2867) {
		v1991 = v2866
		v1996 = v2140
		v2003 = v2018
		goto L518
	} else {
		goto L752
	}
L740:
	;
	v2821 = v2815 + v2816 - v2814
	goto L742
L741:
	;
	v2821 = int32(0)
	goto L742
L742:
	;
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2821)+12))
	if v2822 != 0 {
		goto L743
	} else {
		goto L744
	}
L743:
	;
	v2827 = v2815 + v2822 - int32(1)
	goto L745
L744:
	;
	v2827 = int32(0)
	goto L745
L745:
	;
	if v2827 != 0 {
		goto L746
	} else {
		goto L747
	}
L746:
	;
	v2828 = *(*int32)(unsafe.Add(mBase, uint32(v2821)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2827)+8)) = v2828
	goto L748
L747:
	;
	goto L748
L748:
	;
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2831 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2830 - v2831
	if v2827 != 0 {
		goto L749
	} else {
		goto L750
	}
L749:
	;
	v2838 = v2827 - v2815 + v2831
	goto L751
L750:
	;
	v2838 = int32(0)
	goto L751
L751:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2838
	*(*int32)(unsafe.Add(mBase, uint32(v2821)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2821))) = int64(9020519113)
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+12))
	v2846 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2821)+16)) = v1996 - v33 + v2846
	*(*int32)(unsafe.Add(mBase, uint32(v2821)+12)) = v2844
	v2852 = v2821 - v33 + v2846
	*(*int32)(unsafe.Add(mBase, uint32(v1996)+8)) = v2852
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2821)+24)) = v2020 - v33 + v2846
	*(*int32)(unsafe.Add(mBase, uint32(v2821)+20)) = v2854
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+8)) = v2852
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2852
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2862 + v2846
	goto L517
L752:
	;
	goto L519
L753:
	;
	if base.Ui32(v2880) <= base.Ui32(v2878) {
		goto L756
	} else {
		goto L757
	}
L754:
	;
	v2920 = int32(3)
	v2922 = v2871 + v2918<<(uint(v2920)%32)
	v2925 = v2871 + v2916<<(uint(v2920)%32)
	v2928 = (v2867 - v2916) << (uint(v2920) % 32)
	if v2922 == v2925 {
		goto L767
	} else {
		goto L768
	}
L755:
	;
	goto L754
L756:
	;
	v2916 = v2878
	v2918 = v2878 + int32(1)
	goto L755
L757:
	;
	goto L758
L758:
	;
	v2902 = int32(1)
	v2903 = int32(base.Ui32(v2878+v2880) >> (uint(v2902) % 32))
	v2905 = v2903 + v2902
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2871+v2903<<(uint(int32(3))%32))))
	v2910 = base.B2i32(base.Ui32(v2866) < base.Ui32(v2909))
	if base.Ui32(v2866) < base.Ui32(v2909) {
		goto L759
	} else {
		goto L760
	}
L759:
	;
	v2911 = v2878
	goto L761
L760:
	;
	v2911 = v2905
	goto L761
L761:
	;
	if base.Ui32(v2866) < base.Ui32(v2909) {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	v2912 = v2903
	goto L764
L763:
	;
	v2912 = v2880
	goto L764
L764:
	;
	if v2866 != v2909 {
		v2878 = v2911
		v2880 = v2912
		goto L753
	} else {
		goto L765
	}
L765:
	;
	v2916 = v2903
	v2918 = v2905
	goto L755
L766:
	;
	v3074 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2925)+4)) = v2020 - v33 + v3074
	*(*int32)(unsafe.Add(mBase, uint32(v2925))) = v2866
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2809)+4)) = v3078 + v3074
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+8)) = v2809 - v33 + v3074
	if v2916 != 0 {
		goto L517
	} else {
		goto L812
	}
L767:
	;
	goto L766
L768:
	;
	v2932 = v2922 + v2928
	if base.Ui32(v2925-v2932) <= base.Ui32(int32(0)-v2928<<(uint(int32(1))%32)) {
		goto L769
	} else {
		goto L770
	}
L769:
	;
	v2939 = F___memcpy(m, v2922, v2925, v2928)
	mBase = m.M
	goto L766
L770:
	;
	goto L771
L771:
	;
	v2942 = (v2922 ^ v2925) & int32(3)
	if base.Ui32(v2922) < base.Ui32(v2925) {
		goto L774
	} else {
		goto L775
	}
L772:
	;
	if v3044 == int32(0) {
		goto L767
	} else {
		goto L808
	}
L773:
	;
	if base.Ui32(v3022) <= base.Ui32(int32(3)) {
		v3043 = v3021
		v3044 = v3022
		v3045 = v3023
		goto L772
	} else {
		goto L804
	}
L774:
	;
	if v2942 != 0 {
		goto L777
	} else {
		goto L778
	}
L775:
	;
	goto L776
L776:
	;
	if v2942 != 0 {
		v3004 = v2928
		goto L787
	} else {
		goto L788
	}
L777:
	;
	v3043 = v2925
	v3044 = v2928
	v3045 = v2922
	goto L772
L778:
	;
	goto L779
L779:
	;
	if v2922&int32(3) == int32(0) {
		goto L780
	} else {
		goto L781
	}
L780:
	;
	v3021 = v2925
	v3022 = v2928
	v3023 = v2922
	goto L773
L781:
	;
	goto L782
L782:
	;
	v2949 = v2925
	v2950 = v2928
	v2951 = v2922
	goto L783
L783:
	;
	if v2950 == int32(0) {
		goto L767
	} else {
		goto L785
	}
L784:
	;
	v3021 = v2958
	v3022 = v2960
	v3023 = v2962
	goto L773
L785:
	;
	v2955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2949))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2951))) = uint8(v2955)
	v2957 = int32(1)
	v2958 = v2949 + v2957
	v2960 = v2950 - v2957
	v2962 = v2951 + v2957
	if v2962&int32(3) != 0 {
		v2949 = v2958
		v2950 = v2960
		v2951 = v2962
		goto L783
	} else {
		goto L786
	}
L786:
	;
	goto L784
L787:
	;
	if v3004 == int32(0) {
		goto L767
	} else {
		goto L800
	}
L788:
	;
	if v2932&int32(3) != 0 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v2969 = v2928
	goto L792
L790:
	;
	v2984 = v2928
	goto L791
L791:
	;
	if base.Ui32(v2984) <= base.Ui32(int32(3)) {
		v3004 = v2984
		goto L787
	} else {
		goto L796
	}
L792:
	;
	if v2969 == int32(0) {
		goto L767
	} else {
		goto L794
	}
L793:
	;
	v2984 = v2975
	goto L791
L794:
	;
	v2975 = v2969 - int32(1)
	v2976 = v2922 + v2975
	v2978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2925+v2975))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2976))) = uint8(v2978)
	if v2976&int32(3) != 0 {
		v2969 = v2975
		goto L792
	} else {
		goto L795
	}
L795:
	;
	goto L793
L796:
	;
	v2991 = v2984
	goto L797
L797:
	;
	v2995 = v2991 - int32(4)
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(v2925+v2995)))
	*(*int32)(unsafe.Add(mBase, uint32(v2922+v2995))) = v2998
	if base.Ui32(int32(3)) < base.Ui32(v2995) {
		v2991 = v2995
		goto L797
	} else {
		goto L799
	}
L798:
	;
	v3004 = v2995
	goto L787
L799:
	;
	goto L798
L800:
	;
	v3011 = v3004
	goto L801
L801:
	;
	v3015 = v3011 - int32(1)
	v3018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2925+v3015))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2922+v3015))) = uint8(v3018)
	if v3015 != 0 {
		v3011 = v3015
		goto L801
	} else {
		goto L803
	}
L802:
	;
	goto L767
L803:
	;
	goto L802
L804:
	;
	v3028 = v3021
	v3029 = v3022
	v3030 = v3023
	goto L805
L805:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(v3028)))
	*(*int32)(unsafe.Add(mBase, uint32(v3030))) = v3032
	v3034 = int32(4)
	v3035 = v3028 + v3034
	v3037 = v3030 + v3034
	v3039 = v3029 - v3034
	if base.Ui32(int32(3)) < base.Ui32(v3039) {
		v3028 = v3035
		v3029 = v3039
		v3030 = v3037
		goto L805
	} else {
		goto L807
	}
L806:
	;
	v3043 = v3035
	v3044 = v3039
	v3045 = v3037
	goto L772
L807:
	;
	goto L806
L808:
	;
	v3050 = v3043
	v3051 = v3044
	v3052 = v3045
	goto L809
L809:
	;
	v3054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3050))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3052))) = uint8(v3054)
	v3056 = int32(1)
	v3061 = v3051 - v3056
	if v3061 != 0 {
		v3050 = v3050 + v3056
		v3051 = v3061
		v3052 = v3052 + v3056
		goto L809
	} else {
		goto L811
	}
L810:
	;
	goto L767
L811:
	;
	goto L810
L812:
	;
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3087 = l0 - v3086
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v2871)))
	v3106 = v2140
	goto L813
L813:
	;
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v3106)+8))
	if v3114 == int32(0) {
		goto L517
	} else {
		goto L815
	}
L814:
	;
	goto L517
L815:
	;
	v3117 = v3114 + v3087
	if v3117 == int32(0) {
		goto L517
	} else {
		goto L816
	}
L816:
	;
	v3121 = v3117 + int32(12)
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+4))
	v3128 = int32(0)
	v3129 = v3123
	goto L817
L817:
	;
	if base.Ui32(v3129) <= base.Ui32(v3128) {
		goto L820
	} else {
		goto L821
	}
L818:
	;
	if base.Ui32(v3165) < base.Ui32(v3123) {
		goto L830
	} else {
		goto L831
	}
L819:
	;
	goto L818
L820:
	;
	v3165 = v3128
	goto L819
L821:
	;
	goto L822
L822:
	;
	v3151 = int32(1)
	v3152 = int32(base.Ui32(v3128+v3129) >> (uint(v3151) % 32))
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v3121+v3152<<(uint(int32(3))%32))))
	v3159 = base.B2i32(base.Ui32(v3088) < base.Ui32(v3158))
	if base.Ui32(v3088) < base.Ui32(v3158) {
		goto L823
	} else {
		goto L824
	}
L823:
	;
	v3160 = v3128
	goto L825
L824:
	;
	v3160 = v3152 + v3151
	goto L825
L825:
	;
	if base.Ui32(v3088) < base.Ui32(v3158) {
		goto L826
	} else {
		goto L827
	}
L826:
	;
	v3161 = v3152
	goto L828
L827:
	;
	v3161 = v3129
	goto L828
L828:
	;
	if v3088 != v3158 {
		v3128 = v3160
		v3129 = v3161
		goto L817
	} else {
		goto L829
	}
L829:
	;
	v3165 = v3152
	goto L819
L830:
	;
	v3170 = int32(0)
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v3117+v3165<<(uint(int32(3))%32))+16))
	if v3174 != 0 {
		goto L833
	} else {
		goto L834
	}
L831:
	;
	v3182 = int32(-1)
	goto L832
L832:
	;
	v3183 = v3182 + v3165
	*(*int32)(unsafe.Add(mBase, uint32(v3121+v3183<<(uint(int32(3))%32)))) = v3088
	if v3183 == int32(0) {
		v3106 = v3117
		goto L813
	} else {
		goto L839
	}
L833:
	;
	v3177 = v3087 + v3174
	goto L835
L834:
	;
	v3177 = v3170
	goto L835
L835:
	;
	if v3177 != v3106 {
		goto L836
	} else {
		goto L837
	}
L836:
	;
	v3179 = int32(-1)
	goto L838
L837:
	;
	v3179 = v3170
	goto L838
L838:
	;
	v3182 = v3179
	goto L832
L839:
	;
	goto L814
L840:
	;
	v3218 = v3215
	goto L842
L841:
	;
	v3218 = l2
	goto L842
L842:
	;
	v3223 = v3218<<(uint(int32(2))%32) + l0 + int32(32)
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(v3223)))
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3227 = int32(1)
	v3228 = l0 - v3225 + v3227
	v3230 = l1 << (uint(int32(12)) % 32)
	v3231 = v3228 + v3230
	*(*int32)(unsafe.Add(mBase, uint32(v3231)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v3231))) = int32(-364896016)
	v3235 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3231)+8)) = v3235
	if v3224 != 0 {
		goto L843
	} else {
		goto L844
	}
L843:
	;
	v3241 = v3224 + v3228 - v3227
	goto L845
L844:
	;
	v3241 = v3235
	goto L845
L845:
	;
	if v3241 != 0 {
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v3246 = v3241 - v3228 + int32(1)
	goto L848
L847:
	;
	v3246 = int32(0)
	goto L848
L848:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3231)+12)) = v3246
	v3249 = v3230 | int32(1)
	if v3241 != 0 {
		goto L849
	} else {
		goto L850
	}
L849:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3241)+8)) = v3249
	goto L851
L850:
	;
	goto L851
L851:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3223))) = v3249
	goto L2
}
func F_PageGetExactFreeSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v4 = v2 - v3
	v5 = int32(0)
	if v5 < v4 {
		v8 = v4
	} else {
		v8 = v5
	}
	return v8
}
func F_PageIndexTupleDeleteNoCompact(m *base.Module, l0 int32, l1 int32) {
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
	var v22 int32
	_ = v22
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v17) < base.Ui32(int32(24)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L78
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L78
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L78
	} else {
		goto L79
	}
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if base.Ui32(v20) < base.Ui32(v17) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.Ui32(v22) < base.Ui32(v20) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v22) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if (v22+int32(7))&int32(32760) != v22 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v17 != int32(24) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(base.Ui32(v17+int32(262120)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v38 = int32(0)
	goto L11
L11:
	;
	v39 = int32(65535)
	v40 = v38 & v39
	if base.Ui32(v40) <= base.Ui32((l1-int32(1))&v39) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v47 = l0 + int32(24)
	v52 = v47 + l1<<(uint(int32(2))%32) - int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v55 = int32(base.Ui32(v53) >> (uint(int32(17)) % 32))
	v57 = v53 & int32(32767)
	if base.Ui32(v57) < base.Ui32(v20) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v22) < base.Ui32(v57+v55) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v57 != (v57+int32(7))&int32(65528) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(l1) < base.Ui32(v38&int32(65535)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v81 = (v55 + int32(7)) & int32(65528)
	if base.Ui32(v20) < base.Ui32(v57) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
	v78 = v17
	v79 = v40
	goto L16
L18:
	;
	goto L19
L19:
	;
	v74 = v17 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v74)
	v78 = v74
	v79 = v40 - int32(1)
	goto L16
L20:
	;
	v83 = l0 + v20
	v84 = v83 + v81
	v85 = v57 - v20
	if v84 == v83 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v232 = v78
	v233 = v20
	goto L22
L22:
	;
	v234 = v233 + v81
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v234)
	if base.Ui32(v232&int32(65535)) < base.Ui32(int32(25)) {
		goto L69
	} else {
		goto L70
	}
L23:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v232 = v231
	v233 = v230
	goto L22
L24:
	;
	goto L23
L25:
	;
	v89 = v84 + v85
	if base.Ui32(v83-v89) <= base.Ui32(int32(0)-v85<<(uint(int32(1))%32)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = F___memcpy(m, v84, v83, v85)
	mBase = m.M
	goto L23
L27:
	;
	goto L28
L28:
	;
	v99 = (v84 ^ v83) & int32(3)
	if base.Ui32(v84) < base.Ui32(v83) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v201 == int32(0) {
		goto L24
	} else {
		goto L65
	}
L30:
	;
	if base.Ui32(v179) <= base.Ui32(int32(3)) {
		v200 = v178
		v201 = v179
		v202 = v180
		goto L29
	} else {
		goto L61
	}
L31:
	;
	if v99 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v99 != 0 {
		v161 = v85
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v200 = v83
	v201 = v85
	v202 = v84
	goto L29
L35:
	;
	goto L36
L36:
	;
	if v84&int32(3) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v178 = v83
	v179 = v85
	v180 = v84
	goto L30
L38:
	;
	goto L39
L39:
	;
	v106 = v83
	v107 = v85
	v108 = v84
	goto L40
L40:
	;
	if v107 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L41:
	;
	v178 = v115
	v179 = v117
	v180 = v119
	goto L30
L42:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v112)
	v114 = int32(1)
	v115 = v106 + v114
	v117 = v107 - v114
	v119 = v108 + v114
	if v119&int32(3) != 0 {
		v106 = v115
		v107 = v117
		v108 = v119
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v161 == int32(0) {
		goto L24
	} else {
		goto L57
	}
L45:
	;
	if v89&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v126 = v85
	goto L49
L47:
	;
	v141 = v85
	goto L48
L48:
	;
	if base.Ui32(v141) <= base.Ui32(int32(3)) {
		v161 = v141
		goto L44
	} else {
		goto L53
	}
L49:
	;
	if v126 == int32(0) {
		goto L24
	} else {
		goto L51
	}
L50:
	;
	v141 = v132
	goto L48
L51:
	;
	v132 = v126 - int32(1)
	v133 = v84 + v132
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v135)
	if v133&int32(3) != 0 {
		v126 = v132
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v148 = v141
	goto L54
L54:
	;
	v152 = v148 - int32(4)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v83+v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v84+v152))) = v155
	if base.Ui32(int32(3)) < base.Ui32(v152) {
		v148 = v152
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v161 = v152
	goto L44
L56:
	;
	goto L55
L57:
	;
	v168 = v161
	goto L58
L58:
	;
	v172 = v168 - int32(1)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v84+v172))) = uint8(v175)
	if v172 != 0 {
		v168 = v172
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L24
L60:
	;
	goto L59
L61:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L62
L62:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189
	v191 = int32(4)
	v192 = v185 + v191
	v194 = v187 + v191
	v196 = v186 - v191
	if base.Ui32(int32(3)) < base.Ui32(v196) {
		v185 = v192
		v186 = v196
		v187 = v194
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v200 = v192
	v201 = v196
	v202 = v194
	goto L29
L64:
	;
	goto L63
L65:
	;
	v207 = v200
	v208 = v201
	v209 = v202
	goto L66
L66:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v209))) = uint8(v211)
	v213 = int32(1)
	v218 = v208 - v213
	if v218 != 0 {
		v207 = v207 + v213
		v208 = v218
		v209 = v209 + v213
		goto L66
	} else {
		goto L68
	}
L67:
	;
	goto L24
L68:
	;
	goto L67
L69:
	;
	m.G0 = v15 + int32(48)
	return
L70:
	;
	if v79 <= int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v245 = int32(1)
	goto L72
L72:
	;
	v259 = v245<<(uint(int32(2))%32) + v47 - int32(4)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if base.Ui32(v260) < base.Ui32(int32(131072)) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L69
L74:
	;
	if v245 != v79 {
		v245 = v245 + int32(1)
		goto L72
	} else {
		goto L77
	}
L75:
	;
	if base.Ui32(v57) < base.Ui32(v260&int32(32767)) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = (v260+v81)&int32(32767) | v260&int32(-32768)
	goto L74
L77:
	;
	goto L73
L78:
	;
	return
L79:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v300
	F_errmsg(m, int32(56420), v15)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(490368), int32(1314), int32(109754))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg_internal(m, int32(57607), v15+int32(32))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(490368), int32(1318), int32(109754))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L78
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L78
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v57
	F_errmsg(m, int32(56482), v15+int32(16))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L78
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(490368), int32(1330), int32(109754))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
