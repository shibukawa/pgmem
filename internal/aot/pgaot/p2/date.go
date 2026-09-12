package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeDate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v90 int32
	_ = v90
	var v105 int32
	_ = v105
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v6
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(112)
	return v463
L2:
	;
	v26 = l0
	v31 = v25
	v33 = v6
	goto L5
L3:
	;
	v441 = l1
	goto L4
L4:
	;
	if v441&int32(-32801) != int32(14) {
		goto L94
	} else {
		goto L95
	}
L5:
	;
	v43 = v31 & int32(255)
	goto L7
L6:
	;
	v171 = l1
	v180 = v6
	v181 = v6
	goto L29
L7:
	;
	if base.B2i32(base.Ui32(v43-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v43|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = v26 + int32(1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v59 != 0 {
		v26 = v58
		v31 = v59
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+v33<<(uint(int32(2))%32)))) = v26
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.Ui32((v65-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v463 = int32(-1)
	goto L1
L12:
	;
	if v138&int32(255) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	v72 = v26
	goto L16
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(int32(25)) < base.Ui32((v65|int32(32)-int32(97))&int32(255)) {
		v137 = v26
		v138 = v65
		goto L12
	} else {
		goto L19
	}
L16:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if base.Ui32((v90-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v72 = v72 + int32(1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v137 = v72
	v138 = v90
	goto L12
L19:
	;
	v105 = v26
	goto L20
L20:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if base.Ui32((v123|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v105 = v105 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v137 = v105
	v138 = v123
	goto L12
L22:
	;
	goto L21
L23:
	;
	goto L6
L24:
	;
	v169 = v33 + int32(1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v154)
	v156 = int32(1)
	v157 = v33 + v156
	v159 = v137 + v156
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v160 == v154 {
		v169 = v157
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v33) < base.Ui32(int32(24)) {
		v26 = v159
		v31 = v160
		v33 = v157
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v169 = v157
	goto L23
L29:
	;
	v187 = v180 << (uint(int32(2)) % 32)
	v188 = v19 + v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v189))))
	if base.Ui32(int32(25)) < base.Ui32((v190|int32(32)-int32(97))&int32(255)) {
		v375 = v171
		v385 = v181
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v395 = v375
	v399 = int32(0)
	goto L83
L31:
	;
	v391 = v180 + int32(1)
	if v391 != v169 {
		v171 = v375
		v180 = v391
		v181 = v385
		goto L29
	} else {
		goto L82
	}
L32:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_consts[1064])))
	if v201 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_consts[1064]))) = v340
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+11)))
	if v352 == int32(8) {
		v375 = v171
		v385 = v181
		goto L31
	} else {
		goto L79
	}
L34:
	;
	goto L39
L35:
	;
	goto L36
L36:
	;
	v251 = int32(1662496)
	v257 = int32(1663632)
	goto L52
L37:
	;
	if v238-v239 == int32(0) {
		v340 = v201
		goto L33
	} else {
		goto L51
	}
L39:
	;
	goto L40
L40:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v208 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v209 = v189
	v210 = v201
	v211 = int32(10)
	v212 = v208
	goto L45
L42:
	;
	v234 = v201
	v238 = int32(0)
	goto L43
L43:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	goto L37
L44:
	;
	v234 = v229
	v238 = v231
	goto L43
L45:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v212 != v214 {
		v229 = v210
		v231 = v212
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v229 = v223
	v231 = int32(0)
	goto L44
L47:
	;
	if v214 == int32(0) {
		v229 = v210
		v231 = v212
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v219 = v211 - int32(1)
	if v219 == int32(0) {
		v229 = v210
		v231 = v212
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v222 = int32(1)
	v223 = v210 + v222
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v224 != 0 {
		v209 = v209 + v222
		v210 = v223
		v211 = v219
		v212 = v224
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	goto L36
L52:
	;
	v272 = v251 + (v257-v251)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v272))))
	v274 = v190 - v273
	if v274 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v463 = int32(-1)
	goto L1
L54:
	;
	goto L59
L55:
	;
	v324 = v274
	goto L56
L56:
	;
	v328 = base.B2i32(v324 < int32(0))
	if v324 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L57:
	;
	if v315 == int32(0) {
		v340 = v272
		goto L33
	} else {
		goto L71
	}
L59:
	;
	goto L60
L60:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v283 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v284 = v189
	v285 = v272
	v286 = int32(10)
	v287 = v283
	goto L65
L62:
	;
	v309 = v272
	v313 = int32(0)
	goto L63
L63:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v315 = v313 - v314
	goto L57
L64:
	;
	v309 = v304
	v313 = v306
	goto L63
L65:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v287 != v289 {
		v304 = v285
		v306 = v287
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v304 = v298
	v306 = int32(0)
	goto L64
L67:
	;
	if v289 == int32(0) {
		v304 = v285
		v306 = v287
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v294 = v286 - int32(1)
	if v294 == int32(0) {
		v304 = v285
		v306 = v287
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v297 = int32(1)
	v298 = v285 + v297
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v299 != 0 {
		v284 = v284 + v297
		v285 = v298
		v286 = v294
		v287 = v299
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	v324 = v315
	goto L56
L72:
	;
	v329 = v272 - int32(16)
	goto L74
L73:
	;
	v329 = v257
	goto L74
L74:
	;
	if v324 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v332 = v251
	goto L77
L76:
	;
	v332 = v272 + int32(16)
	goto L77
L77:
	;
	if base.Ui32(v332) <= base.Ui32(v329) {
		v251 = v332
		v257 = v329
		goto L52
	} else {
		goto L78
	}
L78:
	;
	goto L53
L79:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v356 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v356 << (uint(v352) % 32)
	v359 = int32(-1)
	if v352 != v356 {
		v463 = v359
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v355
	if v171&int32(2) != 0 {
		v463 = v359
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v366 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v365 | v366
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = int32(0)
	v375 = v171 | v366
	v385 = int32(1)
	goto L31
L82:
	;
	goto L30
L83:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v19+v399<<(uint(int32(2))%32))))
	if v413 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v441 = v434
	goto L4
L85:
	;
	v415 = F_strlen(m, v413)
	mBase = m.M
	if v415 <= int32(0) {
		v463 = int32(-1)
		goto L1
	} else {
		goto L88
	}
L86:
	;
	v434 = v395
	goto L87
L87:
	;
	v438 = v399 + int32(1)
	if v438 != v169 {
		v395 = v434
		v399 = v438
		goto L83
	} else {
		goto L93
	}
L88:
	;
	v422 = F_DecodeNumber(m, v415, v413, v385, v395, v19+int32(104), l4, v19+int32(108), l3)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	return int32(0)
L90:
	;
	if v422 != 0 {
		v463 = v422
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v427&v395 != 0 {
		v463 = int32(-1)
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v429 | v427
	v434 = v395 | v427
	goto L87
L93:
	;
	goto L84
L94:
	;
	v462 = int32(-1)
	goto L96
L95:
	;
	v462 = int32(0)
	goto L96
L96:
	;
	v463 = v462
	goto L1
}
func F_EncodeDateOnly(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	switch l1 - int32(1) {
	case 0, 3:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v6 {
			v11 = v6
		} else {
			v11 = int32(1) - v6
		}
		v12 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v11) {
		} else {
		}
		v26 = F_pg_ultoa_n(m, v11, l2)
		mBase = m.M
		if v12 <= v26 {
			v37 = l2 + v26
		} else {
			v29 = l2 + v12
			v31 = F_memmove(m, v29-v26, l2, v26)
			mBase = m.M
			v34 = F___memset(m, l2, int32(48), v12-v26)
			mBase = m.M
			v37 = v29
		}
		v38 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
		v41 = v37 + int32(1)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v43 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v42) {
			v57 = F_pg_ultoa_n(m, v42, v41)
			mBase = m.M
			if v43 <= v57 {
				v68 = v41 + v57
			} else {
				v60 = v37 + int32(3)
				v62 = F_memmove(m, v60-v57, v41, v57)
				mBase = m.M
				v65 = F___memset(m, v41, int32(48), v43-v57)
				mBase = m.M
				v68 = v60
			}
		} else {
			v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v41))) = uint16(v53)
			v68 = v37 + int32(3)
		}
		v69 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v68))) = uint8(v69)
		v72 = v68 + int32(1)
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v74 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v73) {
			v88 = F_pg_ultoa_n(m, v73, v72)
			mBase = m.M
			if v74 <= v88 {
				v99 = v72 + v88
			} else {
				v91 = v68 + int32(3)
				v93 = F_memmove(m, v91-v88, v72, v88)
				mBase = m.M
				v96 = F___memset(m, v72, int32(48), v74-v88)
				mBase = m.M
				v99 = v91
			}
		} else {
			v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v72))) = uint16(v84)
			v99 = v68 + int32(3)
		}
		v408 = v99
	case 1:
		v103 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
		v105 = base.B2i32(v103 == int32(1))
		if v103 == int32(1) {
			v106 = int32(12)
		} else {
			v106 = int32(16)
		}
		v108 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
		v109 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v108) {
			v123 = F_pg_ultoa_n(m, v108, l2)
			mBase = m.M
			if v109 <= v123 {
				v134 = l2 + v123
			} else {
				v126 = l2 + v109
				v128 = F_memmove(m, v126-v123, l2, v123)
				mBase = m.M
				v131 = F___memset(m, l2, int32(48), v109-v123)
				mBase = m.M
				v134 = v126
			}
		} else {
			v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v119)
			v134 = l2 + int32(2)
		}
		v135 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v135)
		v138 = v134 + int32(1)
		if v103 == int32(1) {
			v141 = int32(16)
		} else {
			v141 = int32(12)
		}
		v143 = *(*int32)(unsafe.Add(mBase, uint32(l0+v141)))
		v144 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v143) {
			v158 = F_pg_ultoa_n(m, v143, v138)
			mBase = m.M
			if v144 <= v158 {
				v169 = v138 + v158
			} else {
				v161 = v134 + int32(3)
				v163 = F_memmove(m, v161-v158, v138, v158)
				mBase = m.M
				v166 = F___memset(m, v138, int32(48), v144-v158)
				mBase = m.M
				v169 = v161
			}
		} else {
			v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v154)
			v169 = v134 + int32(3)
		}
		v170 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v170)
		v172 = int32(1)
		v173 = v169 + v172
		v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v174 {
			v179 = v174
		} else {
			v179 = v172 - v174
		}
		v180 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v179) {
		} else {
		}
		v194 = F_pg_ultoa_n(m, v179, v173)
		mBase = m.M
		if v180 <= v194 {
			v205 = v173 + v194
		} else {
			v197 = v169 + int32(5)
			v199 = F_memmove(m, v197-v194, v173, v194)
			mBase = m.M
			v202 = F___memset(m, v173, int32(48), v180-v194)
			mBase = m.M
			v205 = v197
		}
		v408 = v205
	case 2:
		v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v207 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v206) {
			v221 = F_pg_ultoa_n(m, v206, l2)
			mBase = m.M
			if v207 <= v221 {
				v232 = l2 + v221
			} else {
				v224 = l2 + v207
				v226 = F_memmove(m, v224-v221, l2, v221)
				mBase = m.M
				v229 = F___memset(m, l2, int32(48), v207-v221)
				mBase = m.M
				v232 = v224
			}
		} else {
			v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v217)
			v232 = l2 + int32(2)
		}
		v233 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v233)
		v236 = v232 + int32(1)
		v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v238 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v237) {
			v252 = F_pg_ultoa_n(m, v237, v236)
			mBase = m.M
			if v238 <= v252 {
				v263 = v236 + v252
			} else {
				v255 = v232 + int32(3)
				v257 = F_memmove(m, v255-v252, v236, v252)
				mBase = m.M
				v260 = F___memset(m, v236, int32(48), v238-v252)
				mBase = m.M
				v263 = v255
			}
		} else {
			v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v248)
			v263 = v232 + int32(3)
		}
		v264 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v264)
		v266 = int32(1)
		v267 = v263 + v266
		v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v268 {
			v273 = v268
		} else {
			v273 = v266 - v268
		}
		v274 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v273) {
		} else {
		}
		v288 = F_pg_ultoa_n(m, v273, v267)
		mBase = m.M
		if v274 <= v288 {
			v299 = v267 + v288
		} else {
			v291 = v263 + int32(5)
			v293 = F_memmove(m, v291-v288, v267, v288)
			mBase = m.M
			v296 = F___memset(m, v267, int32(48), v274-v288)
			mBase = m.M
			v299 = v291
		}
		v408 = v299
	default:
		v303 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
		v305 = base.B2i32(v303 == int32(1))
		if v303 == int32(1) {
			v306 = int32(12)
		} else {
			v306 = int32(16)
		}
		v308 = *(*int32)(unsafe.Add(mBase, uint32(l0+v306)))
		v309 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v308) {
			v323 = F_pg_ultoa_n(m, v308, l2)
			mBase = m.M
			if v309 <= v323 {
				v334 = l2 + v323
			} else {
				v326 = l2 + v309
				v328 = F_memmove(m, v326-v323, l2, v323)
				mBase = m.M
				v331 = F___memset(m, l2, int32(48), v309-v323)
				mBase = m.M
				v334 = v326
			}
		} else {
			v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v319)
			v334 = l2 + int32(2)
		}
		v335 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v334))) = uint8(v335)
		v338 = v334 + int32(1)
		if v303 == int32(1) {
			v341 = int32(16)
		} else {
			v341 = int32(12)
		}
		v343 = *(*int32)(unsafe.Add(mBase, uint32(l0+v341)))
		v344 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v343) {
			v358 = F_pg_ultoa_n(m, v343, v338)
			mBase = m.M
			if v344 <= v358 {
				v369 = v338 + v358
			} else {
				v361 = v334 + int32(3)
				v363 = F_memmove(m, v361-v358, v338, v358)
				mBase = m.M
				v366 = F___memset(m, v338, int32(48), v344-v358)
				mBase = m.M
				v369 = v361
			}
		} else {
			v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343<<(uint(int32(1))%32))+uint32(_consts[1065]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v338))) = uint16(v354)
			v369 = v334 + int32(3)
		}
		v370 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v370)
		v372 = int32(1)
		v373 = v369 + v372
		v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v374 {
			v379 = v374
		} else {
			v379 = v372 - v374
		}
		v380 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v379) {
		} else {
		}
		v394 = F_pg_ultoa_n(m, v379, v373)
		mBase = m.M
		if v380 <= v394 {
			v405 = v373 + v394
		} else {
			v397 = v369 + int32(5)
			v399 = F_memmove(m, v397-v394, v373, v394)
			mBase = m.M
			v402 = F___memset(m, v373, int32(48), v380-v394)
			mBase = m.M
			v405 = v397
		}
		v408 = v405
	}
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v409 <= int32(0) {
		v413 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1067])))
		*(*uint8)(unsafe.Add(mBase, uint32(v408)+2)) = uint8(v413)
		v416 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1068])))
		*(*uint16)(unsafe.Add(mBase, uint32(v408))) = uint16(v416)
		v420 = v408 + int32(3)
	} else {
		v420 = v408
	}
	v421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v421)
	return
}
func F_date_cmp_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[1063]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return v65
}
func F_date_eq_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
	} else {
		if v6 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
		} else {
			if int32(106751982) < v6 {
				v24 = int32(0)
			} else {
				v18 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
			}
		}
	}
	return v24
}
func F_date_ge_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[1063]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return int32(base.Ui32(v65^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_date_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	v9 = m.G0
	v11 = v9 - int32(448)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_ParseDateTime(m, v14, v11+int32(32), int32(129), v11+int32(288), v11+int32(176), v11+int32(388))
	mBase = m.M
	if v24 == int32(0) {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+388))
		v42 = F_DecodeDateTime(m, v11+int32(288), v11+int32(176), v31, v11+int32(392), v11+int32(400), v11+int32(444), v11+int32(396), v11+int32(24))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			if v42 == int32(0) {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+392))
				switch v58 - int32(2) {
				case 0:
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v11)+420))
					if v124 <= int32(-4713) {
						if v124 != int32(-4713) {
							v141 = int32(0)
							v142 = F_errsave_start(m, v13)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								if v142 == int32(0) {
									v214 = v141
									m.G0 = v11 + int32(448)
									return v214
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
										F_errmsg(m, int32(729558), v11+int32(16))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												v214 = v141
												m.G0 = v11 + int32(448)
												return v214
											}
										}
									}
								}
							}
						} else {
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
							if v129 <= int32(10) {
								v141 = int32(0)
								v142 = F_errsave_start(m, v13)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int32(0)
								} else {
									if v142 == int32(0) {
										v214 = v141
										m.G0 = v11 + int32(448)
										return v214
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
											F_errmsg(m, int32(729558), v11+int32(16))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v214 = v141
													m.G0 = v11 + int32(448)
													return v214
												}
											}
										}
									}
								}
							} else {
								v160 = v129
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
								v166 = base.B2i32(int32(2) < v160)
								if int32(2) < v160 {
									v167 = int32(4800)
								} else {
									v167 = int32(4799)
								}
								v168 = v167 + v124
								v173 = base.I32_div_s(v168, int32(4))
								v176 = base.I32_div_s(v168, int32(-100))
								v179 = base.I32_div_s(v168, int32(400))
								if int32(2) < v160 {
									v183 = int32(1)
								} else {
									v183 = int32(13)
								}
								v188 = base.I32_div_s((v183+v160)*int32(7834), int32(256))
								v191 = v161 + v168*int32(365) + v173 + v176 + v179 + v188 - int32(32167)
								if base.Ui32(int32(2147483494)) <= base.Ui32(v191) {
									v194 = int32(0)
									v195 = F_errsave_start(m, v13)
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										if v195 == int32(0) {
											v214 = v194
											m.G0 = v11 + int32(448)
											return v214
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
												F_errmsg(m, int32(729558), v11)
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v13, int32(500133), int32(176), int32(280451))
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int32(0)
													} else {
														v214 = v194
														m.G0 = v11 + int32(448)
														return v214
													}
												}
											}
										}
									}
								} else {
									v214 = v191 - int32(2451545)
									m.G0 = v11 + int32(448)
									return v214
								}
							}
						}
					} else {
						if v124 <= int32(5874897) {
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
							v160 = v134
							v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
							v166 = base.B2i32(int32(2) < v160)
							if int32(2) < v160 {
								v167 = int32(4800)
							} else {
								v167 = int32(4799)
							}
							v168 = v167 + v124
							v173 = base.I32_div_s(v168, int32(4))
							v176 = base.I32_div_s(v168, int32(-100))
							v179 = base.I32_div_s(v168, int32(400))
							if int32(2) < v160 {
								v183 = int32(1)
							} else {
								v183 = int32(13)
							}
							v188 = base.I32_div_s((v183+v160)*int32(7834), int32(256))
							v191 = v161 + v168*int32(365) + v173 + v176 + v179 + v188 - int32(32167)
							if base.Ui32(int32(2147483494)) <= base.Ui32(v191) {
								v194 = int32(0)
								v195 = F_errsave_start(m, v13)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									if v195 == int32(0) {
										v214 = v194
										m.G0 = v11 + int32(448)
										return v214
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
											F_errmsg(m, int32(729558), v11)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v13, int32(500133), int32(176), int32(280451))
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int32(0)
												} else {
													v214 = v194
													m.G0 = v11 + int32(448)
													return v214
												}
											}
										}
									}
								}
							} else {
								v214 = v191 - int32(2451545)
								m.G0 = v11 + int32(448)
								return v214
							}
						} else {
							if v124 != int32(5874898) {
								v141 = int32(0)
								v142 = F_errsave_start(m, v13)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int32(0)
								} else {
									if v142 == int32(0) {
										v214 = v141
										m.G0 = v11 + int32(448)
										return v214
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
											F_errmsg(m, int32(729558), v11+int32(16))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v214 = v141
													m.G0 = v11 + int32(448)
													return v214
												}
											}
										}
									}
								}
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
								if v137 < int32(6) {
									v160 = v137
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
									v166 = base.B2i32(int32(2) < v160)
									if int32(2) < v160 {
										v167 = int32(4800)
									} else {
										v167 = int32(4799)
									}
									v168 = v167 + v124
									v173 = base.I32_div_s(v168, int32(4))
									v176 = base.I32_div_s(v168, int32(-100))
									v179 = base.I32_div_s(v168, int32(400))
									if int32(2) < v160 {
										v183 = int32(1)
									} else {
										v183 = int32(13)
									}
									v188 = base.I32_div_s((v183+v160)*int32(7834), int32(256))
									v191 = v161 + v168*int32(365) + v173 + v176 + v179 + v188 - int32(32167)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v191) {
										v194 = int32(0)
										v195 = F_errsave_start(m, v13)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											if v195 == int32(0) {
												v214 = v194
												m.G0 = v11 + int32(448)
												return v214
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
													F_errmsg(m, int32(729558), v11)
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(500133), int32(176), int32(280451))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v214 = v194
															m.G0 = v11 + int32(448)
															return v214
														}
													}
												}
											}
										}
									} else {
										v214 = v191 - int32(2451545)
										m.G0 = v11 + int32(448)
										return v214
									}
								} else {
									v141 = int32(0)
									v142 = F_errsave_start(m, v13)
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										if v142 == int32(0) {
											v214 = v141
											m.G0 = v11 + int32(448)
											return v214
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
												F_errmsg(m, int32(729558), v11+int32(16))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v214 = v141
														m.G0 = v11 + int32(448)
														return v214
													}
												}
											}
										}
									}
								}
							}
						}
					}
				default:
					F_DateTimeParseError(m, int32(-1), v11+int32(24), v14, int32(358309), v13)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
						v214 = int32(0)
						m.G0 = v11 + int32(448)
						return v214
					}
				case 7:
					v214 = int32(-2147483648)
					m.G0 = v11 + int32(448)
					return v214
				case 8:
					v214 = int32(2147483647)
					m.G0 = v11 + int32(448)
					return v214
				case 9:
					v72 = v11 + int32(400)
					v73 = m.G0
					v75 = v73 - int32(16)
					m.G0 = v75
					*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = int64(0)
					v81 = F_pg_gmtime(m, v75+int32(8))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						if v81 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(294654), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497332), int32(2176), int32(377185))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v98
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v100
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v102
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v104
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v106
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v98 + int32(1900)
							*(*int32)(unsafe.Add(mBase, uint32(v72))) = v108
							*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v100 + int32(1)
							m.G0 = v75 + int32(16)
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v11)+420))
							if v124 <= int32(-4713) {
								if v124 != int32(-4713) {
									v141 = int32(0)
									v142 = F_errsave_start(m, v13)
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										if v142 == int32(0) {
											v214 = v141
											m.G0 = v11 + int32(448)
											return v214
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
												F_errmsg(m, int32(729558), v11+int32(16))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v214 = v141
														m.G0 = v11 + int32(448)
														return v214
													}
												}
											}
										}
									}
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
									if v129 <= int32(10) {
										v141 = int32(0)
										v142 = F_errsave_start(m, v13)
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return int32(0)
										} else {
											if v142 == int32(0) {
												v214 = v141
												m.G0 = v11 + int32(448)
												return v214
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
													F_errmsg(m, int32(729558), v11+int32(16))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															v214 = v141
															m.G0 = v11 + int32(448)
															return v214
														}
													}
												}
											}
										}
									} else {
										v160 = v129
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
										v166 = base.B2i32(int32(2) < v160)
										if int32(2) < v160 {
											v167 = int32(4800)
										} else {
											v167 = int32(4799)
										}
										v168 = v167 + v124
										v173 = base.I32_div_s(v168, int32(4))
										v176 = base.I32_div_s(v168, int32(-100))
										v179 = base.I32_div_s(v168, int32(400))
										if int32(2) < v160 {
											v183 = int32(1)
										} else {
											v183 = int32(13)
										}
										v188 = base.I32_div_s((v183+v160)*int32(7834), int32(256))
										v191 = v161 + v168*int32(365) + v173 + v176 + v179 + v188 - int32(32167)
										if base.Ui32(int32(2147483494)) <= base.Ui32(v191) {
											v194 = int32(0)
											v195 = F_errsave_start(m, v13)
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												if v195 == int32(0) {
													v214 = v194
													m.G0 = v11 + int32(448)
													return v214
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
														F_errmsg(m, int32(729558), v11)
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v13, int32(500133), int32(176), int32(280451))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return int32(0)
															} else {
																v214 = v194
																m.G0 = v11 + int32(448)
																return v214
															}
														}
													}
												}
											}
										} else {
											v214 = v191 - int32(2451545)
											m.G0 = v11 + int32(448)
											return v214
										}
									}
								}
							} else {
								if v124 <= int32(5874897) {
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
									v160 = v134
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
									v166 = base.B2i32(int32(2) < v160)
									if int32(2) < v160 {
										v167 = int32(4800)
									} else {
										v167 = int32(4799)
									}
									v168 = v167 + v124
									v173 = base.I32_div_s(v168, int32(4))
									v176 = base.I32_div_s(v168, int32(-100))
									v179 = base.I32_div_s(v168, int32(400))
									if int32(2) < v160 {
										v183 = int32(1)
									} else {
										v183 = int32(13)
									}
									v188 = base.I32_div_s((v183+v160)*int32(7834), int32(256))
									v191 = v161 + v168*int32(365) + v173 + v176 + v179 + v188 - int32(32167)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v191) {
										v194 = int32(0)
										v195 = F_errsave_start(m, v13)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											if v195 == int32(0) {
												v214 = v194
												m.G0 = v11 + int32(448)
												return v214
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
													F_errmsg(m, int32(729558), v11)
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(500133), int32(176), int32(280451))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v214 = v194
															m.G0 = v11 + int32(448)
															return v214
														}
													}
												}
											}
										}
									} else {
										v214 = v191 - int32(2451545)
										m.G0 = v11 + int32(448)
										return v214
									}
								} else {
									if v124 != int32(5874898) {
										v141 = int32(0)
										v142 = F_errsave_start(m, v13)
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return int32(0)
										} else {
											if v142 == int32(0) {
												v214 = v141
												m.G0 = v11 + int32(448)
												return v214
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
													F_errmsg(m, int32(729558), v11+int32(16))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															v214 = v141
															m.G0 = v11 + int32(448)
															return v214
														}
													}
												}
											}
										}
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+416))
										if v137 < int32(6) {
											v160 = v137
											v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+412))
											v166 = base.B2i32(int32(2) < v160)
											if int32(2) < v160 {
												v167 = int32(4800)
											} else {
												v167 = int32(4799)
											}
											v168 = v167 + v124
											v173 = base.I32_div_s(v168, int32(4))
											v176 = base.I32_div_s(v168, int32(-100))
											v179 = base.I32_div_s(v168, int32(400))
											if int32(2) < v160 {
												v183 = int32(1)
											} else {
												v183 = int32(13)
											}
											v188 = base.I32_div_s((v183+v160)*int32(7834), int32(256))
											v191 = v161 + v168*int32(365) + v173 + v176 + v179 + v188 - int32(32167)
											if base.Ui32(int32(2147483494)) <= base.Ui32(v191) {
												v194 = int32(0)
												v195 = F_errsave_start(m, v13)
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													if v195 == int32(0) {
														v214 = v194
														m.G0 = v11 + int32(448)
														return v214
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
															F_errmsg(m, int32(729558), v11)
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v13, int32(500133), int32(176), int32(280451))
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
																	return int32(0)
																} else {
																	v214 = v194
																	m.G0 = v11 + int32(448)
																	return v214
																}
															}
														}
													}
												}
											} else {
												v214 = v191 - int32(2451545)
												m.G0 = v11 + int32(448)
												return v214
											}
										} else {
											v141 = int32(0)
											v142 = F_errsave_start(m, v13)
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return int32(0)
											} else {
												if v142 == int32(0) {
													v214 = v141
													m.G0 = v11 + int32(448)
													return v214
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
														F_errmsg(m, int32(729558), v11+int32(16))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v13, int32(500133), int32(168), int32(280451))
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return int32(0)
															} else {
																v214 = v141
																m.G0 = v11 + int32(448)
																return v214
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v48 = v42
				F_DateTimeParseError(m, v48, v11+int32(24), v14, int32(358309), v13)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
					v214 = int32(0)
					m.G0 = v11 + int32(448)
					return v214
				}
			}
		}
	} else {
		v48 = v24
		F_DateTimeParseError(m, v48, v11+int32(24), v14, int32(358309), v13)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v54 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
			v214 = int32(0)
			m.G0 = v11 + int32(448)
			return v214
		}
	}
}
func F_date_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4 < v3 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_date_lt_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
		} else {
			if int32(106751982) < v6 {
				return base.B2i32(v4 == int64(9223372036854775807))
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				return int32(base.Ui32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4)) >> (uint(int32(31)) % 32))
			}
		}
	}
}
