package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_Async_Notify(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
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
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
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
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v17 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L130
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L10
	} else {
		goto L126
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L10
	} else {
		goto L122
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[346])))
	if v21 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
	} else {
		goto L119
	}
L8:
	;
	if l0 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v26 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v26 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(637517), v11+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(480648), int32(602), int32(19581))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	if l0&int32(3) == int32(0) {
		v64 = l0
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v98 = v3
	goto L17
L17:
	;
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v98 = v97
	goto L17
L19:
	;
	v97 = v89 - l0
	goto L18
L20:
	;
	v68 = v64
	goto L29
L21:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v48 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v53 = l0
	goto L25
L25:
	;
	v57 = v53 + int32(1)
	if v57&int32(3) == int32(0) {
		v64 = v57
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v89 = v57
	goto L19
L27:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v62 != 0 {
		v53 = v57
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v77 = int32(-2139062144)
	if (int32(16843008)-v74|v74)&v77 == v77 {
		v68 = v68 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v83 = v68
	goto L32
L31:
	;
	goto L30
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 != 0 {
		v83 = v83 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v89 = v83
	goto L19
L34:
	;
	goto L33
L35:
	;
	if l1&int32(3) == int32(0) {
		v122 = l1
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v156 = v3
	goto L37
L37:
	;
	if v98 == int32(0) {
		goto L4
	} else {
		goto L55
	}
L38:
	;
	v156 = v155
	goto L37
L39:
	;
	v155 = v147 - l1
	goto L38
L40:
	;
	v126 = v122
	goto L49
L41:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v106 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v155 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	v111 = l1
	goto L45
L45:
	;
	v115 = v111 + int32(1)
	if v115&int32(3) == int32(0) {
		v122 = v115
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v147 = v115
	goto L39
L47:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v120 != 0 {
		v111 = v115
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v135 = int32(-2139062144)
	if (int32(16843008)-v132|v132)&v135 == v135 {
		v126 = v126 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v141 = v126
	goto L52
L51:
	;
	goto L50
L52:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v145 != 0 {
		v141 = v141 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v147 = v141
	goto L39
L54:
	;
	goto L53
L55:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v98) {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(int32(8000)) <= base.Ui32(v156) {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v163 = int32(4449520)
	v164 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v167
	v172 = F_palloc(m, v98+v156+int32(6))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+2)) = uint16(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v172))) = uint16(v98)
	v177 = v172 + int32(4)
	if (l0^v177)&int32(3) != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	if l1 != 0 {
		goto L81
	} else {
		goto L82
	}
L60:
	;
	goto L59
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v231)
	if v231&int32(255) == int32(0) {
		goto L60
	} else {
		goto L76
	}
L62:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v230 = l0
	v231 = v183
	v232 = v177
	goto L61
L63:
	;
	goto L64
L64:
	;
	if l0&int32(3) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v187 = l0
	v189 = v177
	goto L68
L66:
	;
	v201 = l0
	v203 = v177
	goto L67
L67:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v208 = int32(-2139062144)
	if (int32(16843008)-v205|v205)&v208 != v208 {
		v230 = v201
		v231 = v205
		v232 = v203
		goto L61
	} else {
		goto L72
	}
L68:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v190)
	if v190 == int32(0) {
		goto L60
	} else {
		goto L70
	}
L69:
	;
	v201 = v197
	v203 = v195
	goto L67
L70:
	;
	v194 = int32(1)
	v195 = v189 + v194
	v197 = v187 + v194
	if v197&int32(3) != 0 {
		v187 = v197
		v189 = v195
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v213 = v201
	v214 = v205
	v215 = v203
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v214
	v217 = int32(4)
	v218 = v215 + v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v221 = v213 + v217
	v225 = int32(-2139062144)
	if (v219|(int32(16843008)-v219))&v225 == v225 {
		v213 = v221
		v214 = v219
		v215 = v218
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v230 = v221
	v231 = v219
	v232 = v218
	goto L61
L75:
	;
	goto L74
L76:
	;
	v239 = v230
	v241 = v232
	goto L77
L77:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)) = uint8(v242)
	v244 = int32(1)
	if v242 != 0 {
		v239 = v239 + v244
		v241 = v241 + v244
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L60
L79:
	;
	goto L78
L80:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v333 != 0 {
		goto L107
	} else {
		goto L108
	}
L81:
	;
	v254 = v98 + v177 + int32(1)
	if (l1^v254)&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	goto L83
L83:
	;
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v177)+1)) = uint8(v330)
	goto L80
L84:
	;
	goto L80
L85:
	;
	goto L84
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v308)
	if v308&int32(255) == int32(0) {
		goto L85
	} else {
		goto L101
	}
L87:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v307 = l1
	v308 = v260
	v309 = v254
	goto L86
L88:
	;
	goto L89
L89:
	;
	if l1&int32(3) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v264 = l1
	v266 = v254
	goto L93
L91:
	;
	v278 = l1
	v280 = v254
	goto L92
L92:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v285 = int32(-2139062144)
	if (int32(16843008)-v282|v282)&v285 != v285 {
		v307 = v278
		v308 = v282
		v309 = v280
		goto L86
	} else {
		goto L97
	}
L93:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	*(*uint8)(unsafe.Add(mBase, uint32(v266))) = uint8(v267)
	if v267 == int32(0) {
		goto L85
	} else {
		goto L95
	}
L94:
	;
	v278 = v274
	v280 = v272
	goto L92
L95:
	;
	v271 = int32(1)
	v272 = v266 + v271
	v274 = v264 + v271
	if v274&int32(3) != 0 {
		v264 = v274
		v266 = v272
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v290 = v278
	v291 = v282
	v292 = v280
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291
	v294 = int32(4)
	v295 = v292 + v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v298 = v290 + v294
	v302 = int32(-2139062144)
	if (v296|(int32(16843008)-v296))&v302 == v302 {
		v290 = v298
		v291 = v296
		v292 = v295
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v307 = v298
	v308 = v296
	v309 = v295
	goto L86
L100:
	;
	goto L99
L101:
	;
	v316 = v307
	v318 = v309
	goto L102
L102:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)) = uint8(v319)
	v321 = int32(1)
	if v319 != 0 {
		v316 = v316 + v321
		v318 = v318 + v321
		goto L102
	} else {
		goto L104
	}
L103:
	;
	goto L85
L104:
	;
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v164
	m.G0 = v11 + int32(32)
	return
L106:
	;
	v357 = F_AsyncExistsPendingNotify(m, v172)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L10
	} else {
		goto L113
	}
L107:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v15 <= v334 {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v339 = F_MemoryContextAlloc(m, v337, int32(16))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v172
	v347 = F_list_make1_impl(m, int32(1), v11+int32(12))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = v347
	v352 = int32(4346472)
	v353 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+12)) = v353
	*(*int32)(unsafe.Add(mBase, _consts[347])) = v339
	goto L105
L113:
	;
	if v357 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_pfree(m, v172)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	F_AddEventToPendingNotifies(m, v172)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L118
	}
L117:
	;
	goto L105
L118:
	;
	goto L105
L119:
	;
	F_errmsg_internal(m, int32(211805), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(480648), int32(599), int32(19581))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(8583), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(480648), int32(611), int32(19581))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L10
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(314500), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(480648), int32(617), int32(19581))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L10
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(314403), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(480648), int32(622), int32(19581))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_Async_UnlistenAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[346])))
	if v7 != int32(1) {
		v28 = *(*int32)(unsafe.Add(mBase, _consts[348]))
		if v28 == int32(0) {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[349])))
			if v32 == int32(0) {
				m.G0 = v4 + int32(16)
				return
			} else {
				F_queue_listen(m, int32(2), int32(719562))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					m.G0 = v4 + int32(16)
					return
				}
			}
		} else {
			F_queue_listen(m, int32(2), int32(719562))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				m.G0 = v4 + int32(16)
				return
			}
		}
	} else {
		v12 = F_errstart(m, int32(14), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 == int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[348]))
				if v28 == int32(0) {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[349])))
					if v32 == int32(0) {
						m.G0 = v4 + int32(16)
						return
					} else {
						F_queue_listen(m, int32(2), int32(719562))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							m.G0 = v4 + int32(16)
							return
						}
					}
				} else {
					F_queue_listen(m, int32(2), int32(719562))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						m.G0 = v4 + int32(16)
						return
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[350]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v17
				F_errmsg_internal(m, int32(641411), v4)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(480648), int32(772), int32(293047))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _consts[348]))
						if v28 == int32(0) {
							v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[349])))
							if v32 == int32(0) {
								m.G0 = v4 + int32(16)
								return
							} else {
								F_queue_listen(m, int32(2), int32(719562))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									m.G0 = v4 + int32(16)
									return
								}
							}
						} else {
							F_queue_listen(m, int32(2), int32(719562))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								m.G0 = v4 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_Async_UnlistenOnExit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[346])))
	if v8 != int32(1) {
		m.Env.Pgmem_listen(m, int32(719562), int32(2))
		mBase = m.M
		v32 = *(*int32)(unsafe.Add(mBase, _consts[351]))
		F_list_free_deep(m, v32)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(0)
			F_asyncQueueUnregister(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	} else {
		v13 = F_errstart(m, int32(14), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			if v13 == int32(0) {
				m.Env.Pgmem_listen(m, int32(719562), int32(2))
				mBase = m.M
				v32 = *(*int32)(unsafe.Add(mBase, _consts[351]))
				F_list_free_deep(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(0)
					F_asyncQueueUnregister(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[350]))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v18
				F_errmsg_internal(m, int32(641265), v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(480648), int32(1205), int32(95468))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						m.Env.Pgmem_listen(m, int32(719562), int32(2))
						mBase = m.M
						v32 = *(*int32)(unsafe.Add(mBase, _consts[351]))
						F_list_free_deep(m, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[351])) = int32(0)
							F_asyncQueueUnregister(m)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								m.G0 = v5 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecAsyncAppendResponse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v9 != 0 {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
			if v10&int32(2) == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+136))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v19 + int32(1)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+132))
				*(*int32)(unsafe.Add(mBase, uint32(v23+v19<<(uint(int32(2))%32)))) = v9
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+148))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = F_bms_add_member(m, v28, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v30
					return
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v15 - int32(1)
				return
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v15 - int32(1)
			return
		}
	} else {
		return
	}
}
