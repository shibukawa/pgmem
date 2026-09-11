package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_record_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v9 != v5 {
		if v9 == int32(0) {
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
			if v14 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v13
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v13
			}
			if v13 == int32(0) {
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v19
			}
		}
		if v5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v5
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v26
			if v26 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v4
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v4
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(0)
		}
	} else {
	}
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v39 != 0 {
		F_DeleteExpandedObject(m, v39+int32(12))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
		return
	}
}
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = l1
	v9 = F_palloc0(m, int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(6)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+40)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v13
		return v9
	}
}
func F_set_var_from_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v237 int64
	_ = v237
	var v249 int64
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v20 - int32(43) {
	case 0:
		goto L3
	default:
		v28 = l1
		v29 = v6
		goto L1
	case 2:
		goto L2
	}
L1:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v32 = base.B2i32(v30 == int32(46))
	v33 = v28 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if base.Ui32(int32(9)) < base.Ui32((v34-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v28 = l1 + int32(1)
	v29 = int32(16384)
	goto L1
L3:
	;
	v28 = l1 + int32(1)
	v29 = v6
	goto L1
L4:
	;
	m.G0 = v18 + int32(16)
	return v590
L5:
	;
	v570 = int32(0)
	v571 = F_errsave_start(m, l4)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L24
	} else {
		goto L108
	}
L6:
	;
	if v33&int32(3) == int32(0) {
		v64 = v33
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v100 = F_palloc(m, v97+int32(8))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v97 = v89 - v33
	goto L7
L9:
	;
	v68 = v64
	goto L18
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v48 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v97 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v53 = v33
	goto L14
L14:
	;
	v57 = v53 + int32(1)
	if v57&int32(3) == int32(0) {
		v64 = v57
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v89 = v57
	goto L8
L16:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v62 != 0 {
		v53 = v57
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v77 = int32(-2139062144)
	if (int32(16843008)-v74|v74)&v77 == v77 {
		v68 = v68 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v83 = v68
	goto L21
L20:
	;
	goto L19
L21:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 != 0 {
		v83 = v83 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v89 = v83
	goto L8
L23:
	;
	goto L22
L24:
	;
	return int32(0)
L25:
	;
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v104
	v106 = int32(4)
	v107 = int32(-1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v108 == v104 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v193 = v185 + v100
	v194 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v193))) = uint16(v194)
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)) = uint8(v194)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v198|int32(32) != int32(101) {
		goto L46
	} else {
		goto L47
	}
L27:
	;
	v183 = v33
	v185 = v106
	v186 = v107
	v188 = v6
	goto L26
L28:
	;
	goto L29
L29:
	;
	v112 = v108
	v116 = v33
	v117 = v32
	v118 = v106
	v119 = v107
	v121 = v6
	goto L30
L30:
	;
	v127 = v112 - int32(48)
	if base.Ui32(v127&int32(255)) <= base.Ui32(int32(9)) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v183 = v171
	v185 = v173
	v186 = v174
	v188 = v175
	goto L26
L32:
	;
	if v170&int32(255) != 0 {
		v112 = v170
		v116 = v171
		v117 = v172
		v118 = v173
		v119 = v174
		v121 = v175
		goto L30
	} else {
		goto L43
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v118+v100))) = uint8(v127)
	v134 = int32(1)
	v145 = v116 + v134
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v170 = v146
	v171 = v145
	v172 = v117
	v173 = v118 + v134
	v174 = v119 + (v117^int32(-1))&v134
	v175 = v121 + v117&v134
	goto L32
L34:
	;
	goto L35
L35:
	;
	v148 = v112 & int32(255)
	if v148 != int32(95) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v148 != int32(46) {
		v183 = v116
		v185 = v118
		v186 = v119
		v188 = v121
		goto L26
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v162 = v116 + int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if base.Ui32(int32(9)) < base.Ui32((v163-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L42
	}
L39:
	;
	if v117&int32(1) != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v155 = int32(1)
	v157 = v116 + v155
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v158 != int32(95) {
		v170 = v158
		v171 = v157
		v172 = v155
		v173 = v118
		v174 = v119
		v175 = v121
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L5
L42:
	;
	v170 = v163
	v171 = v162
	v172 = v117
	v173 = v118
	v174 = v119
	v175 = v121
	goto L32
L43:
	;
	goto L31
L44:
	;
	v538 = int32(0)
	v539 = F_errsave_start(m, l4)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L24
	} else {
		goto L103
	}
L45:
	;
	if int32(0) <= v294 {
		goto L70
	} else {
		goto L71
	}
L46:
	;
	v294 = v186
	v295 = v183
	v296 = v188
	goto L45
L47:
	;
	goto L48
L48:
	;
	v203 = int32(0)
	v205 = v183 + int32(1)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	switch v206 - int32(43) {
	case 0:
		goto L51
	default:
		v214 = v203
		v215 = v205
		goto L49
	case 2:
		goto L50
	}
L49:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if base.Ui32(int32(9)) < base.Ui32((v216-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L52
	}
L50:
	;
	v214 = int32(1)
	v215 = v183 + int32(2)
	goto L49
L51:
	;
	v214 = v203
	v215 = v183 + int32(2)
	goto L49
L52:
	;
	v224 = v216
	v232 = v215
	v237 = int64(0)
	goto L53
L53:
	;
	if base.Ui32((v224-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v214 != 0 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	goto L54
L56:
	;
	if v268&int32(255) != 0 {
		v224 = v268
		v232 = v269
		v237 = v270
		goto L53
	} else {
		goto L63
	}
L57:
	;
	v249 = v237*int64(10) + base.I64_extend_i32_u(v224)&int64(15)
	if int64(1073741823) < v249 {
		goto L44
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v224&int32(255) != int32(95) {
		v274 = v232
		v275 = v237
		goto L55
	} else {
		goto L61
	}
L60:
	;
	v253 = v232 + int32(1)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v268 = v254
	v269 = v253
	v270 = v249
	goto L56
L61:
	;
	v260 = v232 + int32(1)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if base.Ui32(int32(9)) < base.Ui32((v261-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v268 = v261
	v269 = v260
	v270 = v237
	goto L56
L63:
	;
	v274 = v269
	v275 = v270
	goto L55
L64:
	;
	v278 = int64(0) - v275
	goto L66
L65:
	;
	v278 = v275
	goto L66
L66:
	;
	v279 = base.I32_wrap_i64(v278)
	v280 = v188 - v279
	v281 = int32(0)
	if v281 < v280 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v284 = v280
	goto L69
L68:
	;
	v284 = v281
	goto L69
L69:
	;
	v294 = v279 + v186
	v295 = v274
	v296 = v284
	goto L45
L70:
	;
	v303 = int32(4)
	v306 = base.I32_div_s(v294+v303, v303)
	v311 = v306 - int32(1)
	goto L72
L71:
	;
	v311 = v294 >> (uint(int32(2)) % 32)
	goto L72
L72:
	;
	v312 = int32(2)
	v314 = v311<<(uint(v312)%32) - v294
	v315 = v185 + v314
	v317 = v315 + v312
	v319 = base.I32_div_s(v317, int32(4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v320 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_pfree(m, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L24
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v327 = F_palloc(m, v319<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L24
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v327
	v330 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v330)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v319
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v311
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v338 = v336 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v338
	if v317 < int32(4) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_pfree(m, v100)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L24
	} else {
		goto L87
	}
L79:
	;
	v342 = int32(1)
	v343 = v342 - v314
	if v319&v342 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v346 = v343 + v100
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	v348 = int32(10)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+2)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+3)))
	v359 = ((v347*v348+v350)*v348+v354)*v348 + v358
	*(*uint16)(unsafe.Add(mBase, uint32(v336)+2)) = uint16(v359)
	v367 = v336 + int32(4)
	v368 = int32(5) - v314
	v369 = v319 - int32(1)
	goto L82
L81:
	;
	v367 = v338
	v368 = v343
	v369 = v319
	goto L82
L82:
	;
	if base.Ui32(v315-int32(2)) < base.Ui32(int32(4)) {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v379 = v367
	v380 = v368
	v381 = v369
	goto L84
L84:
	;
	v389 = v380 + v100
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	v391 = int32(10)
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+2)))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+3)))
	v402 = ((v390*v391+v393)*v391+v397)*v391 + v401
	*(*uint16)(unsafe.Add(mBase, uint32(v379))) = uint16(v402)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+7)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+6)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+5)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+4)))
	v416 = v404 + (v405+(v406+v407*v391)*v391)*v391
	*(*uint16)(unsafe.Add(mBase, uint32(v379)+2)) = uint16(v416)
	v422 = int32(2)
	if v422 < v381 {
		v379 = v379 + int32(4)
		v380 = v380 + int32(8)
		v381 = v381 - v422
		goto L84
	} else {
		goto L86
	}
L85:
	;
	goto L78
L86:
	;
	goto L85
L87:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v444 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v520
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v295
	v590 = int32(1)
	goto L4
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v520 = int32(0)
	v524 = v506
	goto L88
L90:
	;
	v451 = v444
	v455 = v443
	goto L94
L91:
	;
	goto L92
L92:
	;
	if v444 != 0 {
		v520 = v444
		v524 = v443
		goto L88
	} else {
		goto L102
	}
L93:
	;
	v479 = v451
	goto L98
L94:
	;
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v455))))
	if v465 != 0 {
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v506 = v443 + v444<<(uint(int32(1))%32)
	goto L89
L96:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v467 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v466 - v467
	if v467 < v451 {
		v451 = v451 - v467
		v455 = v455 + int32(2)
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v455-int32(2)+v479<<(uint(int32(1))%32)))))
	if v496 != 0 {
		v520 = v479
		v524 = v455
		goto L88
	} else {
		goto L100
	}
L99:
	;
	v506 = v455
	goto L89
L100:
	;
	v497 = int32(1)
	if v497 < v479 {
		v479 = v479 - v497
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v506 = v443
	goto L89
L103:
	;
	if v539 == int32(0) {
		v590 = v538
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L24
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(103371), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L24
	} else {
		goto L106
	}
L106:
	;
	F_errsave_finish(m, l4, int32(469336), int32(7320), int32(194288))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L24
	} else {
		goto L107
	}
L107:
	;
	v590 = v538
	goto L4
L108:
	;
	if v571 == int32(0) {
		v590 = v570
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L24
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(461065)
	F_errmsg(m, int32(674710), v18)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L24
	} else {
		goto L111
	}
L111:
	;
	F_errsave_finish(m, l4, int32(469336), int32(7326), int32(194288))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L24
	} else {
		goto L112
	}
L112:
	;
	v590 = v570
	goto L4
}
