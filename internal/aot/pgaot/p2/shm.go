package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F___shm_mapname(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	v5 = l0
	goto L1
L1:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v11 == int32(47) {
		v5 = v5 + int32(1)
		goto L1
	} else {
		goto L3
	}
L2:
	;
	goto L11
L3:
	;
	goto L2
L4:
	;
	v136 = int32(_a_F___shm_mapname_0)
	goto L38
L5:
	;
	if base.Ui32(v112) < base.Ui32(int32(256)) {
		goto L4
	} else {
		goto L35
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___shm_mapname[0])) = int32(28)
	return int32(0)
L7:
	;
	if v99 == v5 {
		goto L6
	} else {
		goto L30
	}
L8:
	;
	goto L7
L9:
	;
	v89 = v84
	goto L26
L10:
	;
	v84 = v76
	goto L9
L11:
	;
	if v5&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v22 = v5
	goto L17
L15:
	;
	v36 = v5
	goto L16
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 != v45 {
		v76 = v36
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if base.B2i32(v27 == int32(0))|base.B2i32(int32(47) == v27) != 0 {
		v99 = v22
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v36 = v33
	goto L16
L19:
	;
	v33 = v22 + int32(1)
	if v33&int32(3) != 0 {
		v22 = v33
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v51 = v36
	v53 = v42
	goto L22
L22:
	;
	v57 = v53 ^ int32(791621423)
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v76 = v51
		goto L10
	} else {
		goto L24
	}
L23:
	;
	v84 = v66
	goto L9
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v66 = v51 + int32(4)
	v70 = int32(-2139062144)
	if (v64|(int32(16843008)-v64))&v70 == v70 {
		v51 = v66
		v53 = v64
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v91 == int32(0) {
		v99 = v89
		goto L8
	} else {
		goto L28
	}
L27:
	;
	v99 = v89
	goto L8
L28:
	;
	if v91 != int32(47) {
		v89 = v89 + int32(1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v111 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v112 = v99 - v5
	if int32(2) < v112 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v115 != int32(46) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99-int32(1)))))
	if v120 != int32(46) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L6
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___shm_mapname[0])) = int32(37)
	return int32(0)
L36:
	;
	v308 = l1 + int32(9)
	v310 = v112 + int32(1)
	if base.Ui32(int32(512)) <= base.Ui32(v310) {
		goto L84
	} else {
		goto L85
	}
L38:
	;
	goto L39
L39:
	;
	v144 = l1 + int32(9)
	if (l1^v136)&int32(3) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if base.Ui32(v276) < base.Ui32(v144) {
		goto L77
	} else {
		goto L78
	}
L44:
	;
	if l1&int32(3) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v144) < base.Ui32(int32(4)) {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	v180 = v144 & int32(-4)
	if base.Ui32(v144) < base.Ui32(int32(64)) {
		v230 = v174
		v231 = v175
		goto L58
	} else {
		goto L59
	}
L48:
	;
	v174 = v136
	v175 = l1
	goto L47
L49:
	;
	goto L50
L50:
	;
	goto L52
L52:
	;
	goto L53
L53:
	;
	v157 = v136
	v158 = l1
	goto L54
L54:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v162)
	v164 = int32(1)
	v165 = v157 + v164
	v167 = v158 + v164
	if v167&int32(3) == int32(0) {
		v174 = v165
		v175 = v167
		goto L47
	} else {
		goto L56
	}
L55:
	;
	v174 = v165
	v175 = v167
	goto L47
L56:
	;
	if base.Ui32(v167) < base.Ui32(v144) {
		v157 = v165
		v158 = v167
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if base.Ui32(v180) <= base.Ui32(v231) {
		v275 = v230
		v276 = v231
		goto L43
	} else {
		goto L64
	}
L59:
	;
	v184 = v180 + int32(-64)
	if base.Ui32(v184) < base.Ui32(v175) {
		v230 = v174
		v231 = v175
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v187 = v174
	v188 = v175
	goto L61
L61:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v187)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+20)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v187)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+24)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v187)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+28)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v187)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+32)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v187)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+36)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v187)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+40)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v187)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+44)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+48)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+52)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v187)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+56)) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v187)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+60)) = v222
	v224 = int32(-64)
	v225 = v187 - v224
	v227 = v188 - v224
	if base.Ui32(v227) <= base.Ui32(v184) {
		v187 = v225
		v188 = v227
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v230 = v225
	v231 = v227
	goto L58
L63:
	;
	goto L62
L64:
	;
	v237 = v230
	v238 = v231
	goto L65
L65:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v242
	v244 = int32(4)
	v245 = v237 + v244
	v247 = v238 + v244
	if base.Ui32(v247) < base.Ui32(v180) {
		v237 = v245
		v238 = v247
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v275 = v245
	v276 = v247
	goto L43
L67:
	;
	goto L66
L68:
	;
	v275 = v136
	v276 = l1
	goto L43
L69:
	;
	goto L70
L70:
	;
	goto L72
L72:
	;
	goto L73
L73:
	;
	v256 = v136
	v257 = l1
	goto L74
L74:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	*(*uint8)(unsafe.Add(mBase, uint32(v257))) = uint8(v261)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)) = uint8(v263)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v257)+2)) = uint8(v265)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v257)+3)) = uint8(v267)
	v269 = int32(4)
	v270 = v256 + v269
	v272 = v257 + v269
	if base.Ui32(v272) <= base.Ui32(v144-int32(4)) {
		v256 = v270
		v257 = v272
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v275 = v270
	v276 = v272
	goto L43
L76:
	;
	goto L75
L77:
	;
	v282 = v275
	v283 = v276
	goto L80
L78:
	;
	goto L79
L79:
	;
	goto L36
L80:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v287)
	v289 = int32(1)
	v292 = v283 + v289
	if v292 != v144 {
		v282 = v282 + v289
		v283 = v292
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L79
L82:
	;
	goto L81
L83:
	;
	return l1
L84:
	;
	if v310 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v317 = v308 + v310
	if (v308^v5)&int32(3) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	base.MemoryCopy(m, v308, v5, v310)
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L83
L90:
	;
	if base.Ui32(v449) < base.Ui32(v317) {
		goto L124
	} else {
		goto L125
	}
L91:
	;
	if v308&int32(3) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L93
L93:
	;
	if base.Ui32(v317) < base.Ui32(int32(4)) {
		goto L115
	} else {
		goto L116
	}
L94:
	;
	v353 = v317 & int32(-4)
	if base.Ui32(v317) < base.Ui32(int32(64)) {
		v403 = v347
		v404 = v348
		goto L105
	} else {
		goto L106
	}
L95:
	;
	v347 = v5
	v348 = v308
	goto L94
L96:
	;
	goto L97
L97:
	;
	if v310 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v347 = v5
	v348 = v308
	goto L94
L99:
	;
	goto L100
L100:
	;
	v330 = v5
	v331 = v308
	goto L101
L101:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v335)
	v337 = int32(1)
	v338 = v330 + v337
	v340 = v331 + v337
	if v340&int32(3) == int32(0) {
		v347 = v338
		v348 = v340
		goto L94
	} else {
		goto L103
	}
L102:
	;
	v347 = v338
	v348 = v340
	goto L94
L103:
	;
	if base.Ui32(v340) < base.Ui32(v317) {
		v330 = v338
		v331 = v340
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	if base.Ui32(v353) <= base.Ui32(v404) {
		v448 = v403
		v449 = v404
		goto L90
	} else {
		goto L111
	}
L106:
	;
	v357 = v353 + int32(-64)
	if base.Ui32(v357) < base.Ui32(v348) {
		v403 = v347
		v404 = v348
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v360 = v347
	v361 = v348
	goto L108
L108:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+4)) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+8)) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+12)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+16)) = v373
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v360)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+20)) = v375
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v360)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+24)) = v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v360)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+28)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v360)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+32)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v360)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+36)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v360)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+40)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v360)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+44)) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v360)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+48)) = v389
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v360)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+52)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v360)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+56)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v360)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+60)) = v395
	v397 = int32(-64)
	v398 = v360 - v397
	v400 = v361 - v397
	if base.Ui32(v400) <= base.Ui32(v357) {
		v360 = v398
		v361 = v400
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v403 = v398
	v404 = v400
	goto L105
L110:
	;
	goto L109
L111:
	;
	v410 = v403
	v411 = v404
	goto L112
L112:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v415
	v417 = int32(4)
	v418 = v410 + v417
	v420 = v411 + v417
	if base.Ui32(v420) < base.Ui32(v353) {
		v410 = v418
		v411 = v420
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v448 = v418
	v449 = v420
	goto L90
L114:
	;
	goto L113
L115:
	;
	v448 = v5
	v449 = v308
	goto L90
L116:
	;
	goto L117
L117:
	;
	if base.Ui32(v310) < base.Ui32(int32(4)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v448 = v5
	v449 = v308
	goto L90
L119:
	;
	goto L120
L120:
	;
	v429 = v5
	v430 = v308
	goto L121
L121:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	*(*uint8)(unsafe.Add(mBase, uint32(v430))) = uint8(v434)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+1)) = uint8(v436)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+2)) = uint8(v438)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+3)) = uint8(v440)
	v442 = int32(4)
	v443 = v429 + v442
	v445 = v430 + v442
	if base.Ui32(v445) <= base.Ui32(v317-int32(4)) {
		v429 = v443
		v430 = v445
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v448 = v443
	v449 = v445
	goto L90
L123:
	;
	goto L122
L124:
	;
	v455 = v448
	v456 = v449
	goto L127
L125:
	;
	goto L126
L126:
	;
	goto L83
L127:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(v460)
	v462 = int32(1)
	v465 = v456 + v462
	if v465 != v317 {
		v455 = v455 + v462
		v456 = v465
		goto L127
	} else {
		goto L129
	}
L128:
	;
	goto L126
L129:
	;
	goto L128
}
func F_shm_mq_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	v3 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v3))
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v10 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)) = uint16(v10)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1&int32(-8) - int32(40)
	return l0
}
func F_shm_mq_send_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	v6 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v231
	m.G0 = v22 + int32(16)
	return v232
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v30 = base.I64_extend_i32_u(v29)
	v31 = int32(2)
	v39 = v6
	goto L5
L3:
	;
	v211 = v6
	goto L4
L4:
	;
	v231 = v211
	v232 = int32(0)
	goto L1
L5:
	;
	v51 = int64(0)
	v54 = base.AtomicRmwCmpxchg64(m, v24, int32(16), v51, v51)
	v58 = base.AtomicRmwCmpxchg64(m, v24, int32(24), v51, v51)
	v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+36)))
	if v60 != 0 {
		v231 = v39
		v232 = v31
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v211 = v199
	goto L4
L7:
	;
	v62 = v58 + v59
	v63 = v54 + v30 - v62
	v65 = base.I64_extend_i32_u(l1 - v39)
	if base.Ui64(v63) < base.Ui64(v65) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui32(v199) < base.Ui32(l1) {
		v39 = v199
		goto L5
	} else {
		goto L61
	}
L9:
	;
	v67 = v63
	goto L11
L10:
	;
	v67 = v65
	goto L11
L11:
	;
	if v67 == int64(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v70 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v179 = base.I32_wrap_i64(v67)
	v180 = base.I64_rem_u_s(v62, v30)
	v181 = base.I32_wrap_i64(v180)
	v182 = v29 - v181
	if base.Ui32(v179) < base.Ui32(v182) {
		goto L55
	} else {
		goto L56
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v105 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v106 = int64(0)
	v108 = int32(24)
	v109 = base.AtomicRmwCmpxchg64(m, v24, v108, v106, v106)
	v112 = base.AtomicRmwXchg64(m, v24, v108, v105+v109)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v115 = v113 + int32(20)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v103)
	v199 = v39
	goto L8
L19:
	;
	if v73 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v99 = F_shm_mq_wait_internal(m, v24, v24+int32(4), v73)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L24
	} else {
		goto L32
	}
L22:
	;
	v88 = base.AtomicRmwXchg32(m, v24, int32(0), int32(1))
	if v88 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v78 = F_GetBackgroundWorkerPid(m, v73, v22+int32(12))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	if base.Ui32(v78) < base.Ui32(int32(2)) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+36)) = uint8(v84)
	v231 = v39
	v232 = v31
	goto L1
L27:
	;
	F_s_lock(m, v24, int32(_a_F_shm_mq_send_bytes_0), int32(246), int32(_a_F_shm_mq_send_bytes_1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L24
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v95 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v24))), uint32(v95))
	if v94 != 0 {
		goto L18
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v231 = v39
	v232 = int32(1)
	goto L1
L32:
	;
	if v99 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+36)) = uint8(v101)
	v231 = v39
	v232 = v31
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	if l3 != 0 {
		goto L48
	} else {
		goto L49
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v119 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v122 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[0]))
	if v126 == v122 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v128 = m.G0
	v130 = v128 - int32(16)
	m.G0 = v130
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[1]))
	if v133 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v156 = F_pgmem_kill(m, v122, int32(23))
	mBase = m.M
	goto L35
L42:
	;
	m.G0 = v130 + int32(16)
	goto L34
L43:
	;
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+15)) = uint8(v136)
	goto L44
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[2]))
	v144 = F_write(m, v140, v130+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v144 {
		goto L42
	} else {
		goto L46
	}
L45:
	;
	goto L42
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[3]))
	if v148 == int32(27) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v231 = v39
	v232 = int32(1)
	goto L1
L49:
	;
	goto L50
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[4]))
	v167 = F_WaitLatch(m, v163, int32(33), int32(0), int32(134217764))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(0)
	goto L52
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_send_bytes[5]))
	if v174 == int32(0) {
		v199 = v39
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L24
	} else {
		goto L54
	}
L54:
	;
	v199 = v39
	goto L8
L55:
	;
	v184 = v179
	goto L57
L56:
	;
	v184 = v182
	goto L57
L57:
	;
	if v184 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+37)))
	base.MemoryCopy(m, v24+int32(38)+v185+v181, l2+v39, v184)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v190 + (v184+int32(7))&int32(-8)
	v199 = v184 + v39
	goto L8
L61:
	;
	goto L6
}
func F_shm_toc_create(m *base.Module, l0 int64, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = l0
	v5 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+8)), uint32(v5))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2 & int32(-32)
	return l1
}
func F_shm_toc_estimate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_mul_size(m, v3, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_add_size(m, int32(24), v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = F_add_size(m, v9, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return (v12 + int32(31)) & int32(-32)
			}
		}
	}
}
