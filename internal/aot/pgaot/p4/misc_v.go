package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_VirtualXactLock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int64
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int64
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v476 int32
	_ = v476
	var v478 int64
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v526 int32
	_ = v526
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_LWLockRelease(m, v356)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L11
	} else {
		goto L109
	}
L2:
	;
	m.G0 = v12 + int32(128)
	return v526
L3:
	;
	v526 = int32(1)
	goto L2
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v18
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)) = uint8(v20)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v23 == v20 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v14
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v113
	v117 = int32(0)
	if v14 < v117 {
		v135 = v117
		goto L27
	} else {
		goto L28
	}
L7:
	;
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v28
	v34 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(8), v12+int32(96))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v40 = v17
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v40
	v46 = int32(0)
	v52 = l1 ^ int32(1)
	v55 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v46, v52, v46, v46)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	if v34 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v40 = v34
	goto L10
L14:
	;
	if v55 == int32(0) {
		v526 = v46
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v63 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v65 = int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)))
	if v66 != v65 {
		v526 = v65
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L18
L18:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)) = uint8(v78)
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v80
	v84 = F_TwoPhaseGetXidByVirtualXID(m, v12, v12+int32(96))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L20
	}
L19:
	;
	goto L3
L20:
	;
	v87 = base.B2i32(v84 == int32(0))
	if v84 == int32(0) {
		v526 = v87
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v84
	v98 = int32(0)
	v101 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v98, v52, v98, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	if v101 == int32(0) {
		v526 = v87
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v109 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)))
	if v111 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	if v135 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	if base.Ui32(v124) <= base.Ui32(v14) {
		v135 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v129 = v126 + v14*int32(640)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	if v131 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v132 = v129
	goto L32
L31:
	;
	v132 = int32(0)
	goto L32
L32:
	;
	v135 = v132
	goto L27
L33:
	;
	v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v138
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v140)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v143 == v140 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v231 = v135 + int32(584)
	v233 = F_LWLockAcquire(m, v231, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L11
	} else {
		goto L51
	}
L36:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v146
	v153 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(24), v12+int32(127))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	if v153 == int32(0) {
		v526 = int32(1)
		goto L2
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v153
	v162 = int32(0)
	v168 = l1 ^ int32(1)
	v171 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v162, v168, v162, v162)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	if v171 == int32(0) {
		v526 = v162
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v179 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v181 = int32(1)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v182 != v181 {
		v526 = v181
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L43
L43:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v194)
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v196
	v202 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(16), v12+int32(127))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	goto L3
L45:
	;
	v205 = base.B2i32(v202 == int32(0))
	if v202 == int32(0) {
		v526 = v205
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v202
	v216 = int32(0)
	v219 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v216, v168, v216, v216)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	if v219 == int32(0) {
		v526 = v205
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v227 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v229 != 0 {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v135)+52))
	if v14 == v235 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if l1 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L53:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v135)+612))
	if v237 == v113 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_LWLockRelease(m, v231)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v241
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v243)
	v246 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v246 == v243 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v249
	v256 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(72), v12+int32(127))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	if v256 == int32(0) {
		v526 = int32(1)
		goto L2
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v256
	v265 = int32(0)
	v271 = l1 ^ int32(1)
	v274 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v265, v271, v265, v265)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	if v274 == int32(0) {
		v526 = v265
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v282 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v284 = int32(1)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v285 != v284 {
		v526 = v284
		goto L2
	} else {
		goto L64
	}
L64:
	;
	goto L65
L65:
	;
	v297 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v297)
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v299
	v305 = F_TwoPhaseGetXidByVirtualXID(m, v12-int32(-64), v12+int32(127))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L67
	}
L66:
	;
	goto L3
L67:
	;
	v308 = base.B2i32(v305 == int32(0))
	if v305 == int32(0) {
		v526 = v308
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v305
	v319 = int32(0)
	v322 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v319, v271, v319, v319)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	if v322 == int32(0) {
		v526 = v308
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v330 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v332 != 0 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	goto L66
L73:
	;
	F_LWLockRelease(m, v231)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v338 = int32(1)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+608)))
	if v339 == v338 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v526 = int32(0)
	goto L2
L77:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[1104]))
	v346 = F_get_hash_value(m, v343, v12+int32(80))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L11
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v135)+36))
	F_LWLockRelease(m, v231)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L11
	} else {
		goto L88
	}
L80:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v356 = v349 + v346&int32(15)<<(uint(int32(7))%32) + int32(23296)
	v358 = F_LWLockAcquire(m, v356, int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	v364 = F_SetupLockInTable(m, int32(1586680), v135, v12+int32(80), v346, int32(7))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	if v364 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+128))
	v370 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v368)+128)) = v369 + v370
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+16)) = v373 | int32(128)
	v378 = v368 + int32(116)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v381 = v379 + v370
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v368)+72))
	if v383 == v381 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v368)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+20)) = v385 & int32(-129)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+12)) = v389 | int32(128)
	F_LWLockRelease(m, v356)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	v395 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+608)) = uint8(v395)
	goto L79
L88:
	;
	v407 = int32(0)
	v411 = F_LockAcquireExtended(m, v12+int32(80), int32(5), v407, v407, v407, v407)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	v417 = F_LockRelease(m, v12+int32(80), int32(5), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	v419 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v419
	v421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v421)
	v424 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v424 == v421 {
		v526 = v338
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v401 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v429
	v435 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(56), v12+int32(127))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L11
	} else {
		goto L95
	}
L93:
	;
	v439 = v401
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v439
	v445 = int32(0)
	v453 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v445, v445, v445, v445)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L11
	} else {
		goto L97
	}
L95:
	;
	if v435 == int32(0) {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	v439 = v435
	goto L94
L97:
	;
	if v453 == int32(0) {
		v526 = v445
		goto L2
	} else {
		goto L98
	}
L98:
	;
	v461 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	v463 = int32(1)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v464 != v463 {
		v526 = v463
		goto L2
	} else {
		goto L100
	}
L100:
	;
	goto L101
L101:
	;
	v476 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v476)
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v478
	v484 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(48), v12+int32(127))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L11
	} else {
		goto L103
	}
L102:
	;
	goto L3
L103:
	;
	v487 = base.B2i32(v484 == int32(0))
	if v484 == int32(0) {
		v526 = v487
		goto L2
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v484
	v498 = int32(0)
	v502 = F_LockAcquireExtended(m, v12+int32(108), int32(5), v498, v498, v498, v498)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	if v502 == int32(0) {
		v526 = v487
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v510 = F_LockRelease(m, v12+int32(108), int32(5), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v512 != 0 {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	goto L102
L109:
	;
	F_LWLockRelease(m, v231)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(12976), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(243557)
	F_errhint(m, int32(625496), v12+int32(32))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(474605), int32(4789), int32(302114))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vac_close_indexes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v4 = l0
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pfree(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	v8 = v4 - int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1+v8<<(uint(int32(2))%32))))
	F_relation_close(m, v12, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	if v8 != 0 {
		v4 = v8
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L3
}
func F_varbit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v15 <= int32(0) {
			v60 = v11
			m.G0 = v8 + int32(16)
			return v60
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			if v18 <= v15 {
				v60 = v11
				m.G0 = v8 + int32(16)
				return v60
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16777346))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
							F_errmsg(m, int32(637405), v8)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(470441), int32(758), int32(95981))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v25 = int32(8)
					v26 = base.I32_div_s(v15+int32(7), v25)
					v28 = v26 + v25
					v29 = F_palloc(m, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v28 << (uint(int32(2)) % 32)
						v35 = int32(8)
						v40 = v28 & int32(1073741823)
						v42 = v40 - v35
						if v42 != 0 {
							v43 = F__emscripten_memcpy_bulkmem(m, v29+v35, v11+v35, v42)
							mBase = m.M
						} else {
						}
						v49 = v28<<(uint(int32(3))%32) - v15 + int32(-64)
						if v49 <= int32(0) {
							v60 = v29
						} else {
							v54 = v29 + v40 - int32(1)
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
							v58 = v55 & (int32(255) << (uint(v49) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
							v60 = v29
						}
						m.G0 = v8 + int32(16)
						return v60
					}
				}
			}
		}
	}
}
func F_varbit_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	switch v16 - int32(88) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v25 = v15
		goto L4
	case 10:
		goto L5
	default:
		goto L6
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v340
L2:
	;
	if int32(0) < v13 {
		goto L55
	} else {
		goto L56
	}
L3:
	;
	v85 = v15 + int32(1)
	if v85&int32(3) == int32(0) {
		v109 = v85
		goto L28
	} else {
		goto L29
	}
L4:
	;
	if v25&int32(3) == int32(0) {
		v50 = v25
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v25 = v15 + int32(1)
	goto L4
L6:
	;
	if v16 == int32(120) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v16 != int32(66) {
		v25 = v15
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v169 = v25
	v170 = int32(1)
	v171 = v83
	goto L2
L10:
	;
	v83 = v75 - v25
	goto L9
L11:
	;
	v54 = v50
	goto L20
L12:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v83 = int32(0)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = v25
	goto L16
L16:
	;
	v43 = v39 + int32(1)
	if v43&int32(3) == int32(0) {
		v50 = v43
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v75 = v43
	goto L10
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 != 0 {
		v39 = v43
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 == v63 {
		v54 = v54 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v69 = v54
	goto L23
L22:
	;
	goto L21
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 != 0 {
		v69 = v69 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v75 = v69
	goto L10
L25:
	;
	goto L24
L26:
	;
	if int32(536870911) <= v142 {
		goto L43
	} else {
		goto L44
	}
L27:
	;
	v142 = v134 - v85
	goto L26
L28:
	;
	v113 = v109
	goto L37
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v93 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v142 = int32(0)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v98 = v85
	goto L33
L33:
	;
	v102 = v98 + int32(1)
	if v102&int32(3) == int32(0) {
		v109 = v102
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v134 = v102
	goto L27
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v107 != 0 {
		v98 = v102
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v122 = int32(-2139062144)
	if (int32(16843008)-v119|v119)&v122 == v122 {
		v113 = v113 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v128 = v113
	goto L40
L39:
	;
	goto L38
L40:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 != 0 {
		v128 = v128 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v134 = v128
	goto L27
L42:
	;
	goto L41
L43:
	;
	v145 = F_errsave_start(m, v14)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v169 = v85
	v170 = v2
	v171 = v142 << (uint(int32(2)) % 32)
	goto L2
L46:
	;
	return int32(0)
L47:
	;
	if v145 == int32(0) {
		v340 = v2
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(2147483640)
	F_errmsg(m, int32(638247), v9+int32(-16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, v14, int32(470441), int32(500), int32(265571))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v340 = v2
	goto L1
L52:
	;
	v287 = v191
	v288 = int32(128)
	v290 = v169
	goto L89
L53:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v213 == int32(0) {
		v340 = v182
		goto L1
	} else {
		goto L69
	}
L54:
	;
	v195 = F_errsave_start(m, v14)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L46
	} else {
		goto L64
	}
L55:
	;
	if v13 < v171 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	v175 = v171
	goto L57
L57:
	;
	v178 = int32(8)
	v179 = base.I32_div_s(v171+int32(7), v178)
	v181 = v179 + v178
	v182 = F_palloc0(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L46
	} else {
		goto L59
	}
L58:
	;
	v175 = v13
	goto L57
L59:
	;
	if v171 < v175 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v185 = v171
	goto L62
L61:
	;
	v185 = v175
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v181 << (uint(int32(2)) % 32)
	v191 = v182 + int32(8)
	if v170 == int32(0) {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	goto L52
L64:
	;
	if v195 == int32(0) {
		v340 = v2
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L46
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v13
	F_errmsg(m, int32(637405), v9+int32(-32))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L46
	} else {
		goto L67
	}
L67:
	;
	F_errsave_finish(m, v14, int32(470441), int32(514), int32(265571))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L46
	} else {
		goto L68
	}
L68:
	;
	v340 = v2
	goto L1
L69:
	;
	v217 = v191
	v218 = v213
	v220 = v169
	v224 = int32(0)
	goto L70
L70:
	;
	v226 = v218 - int32(48)
	if base.Ui32(v226&int32(255)) < base.Ui32(int32(10)) {
		v247 = v226
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v340 = v182
	goto L1
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v280)
	v285 = v220 + int32(1)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v286 != 0 {
		v217 = v282
		v218 = v286
		v220 = v285
		v224 = v281
		goto L70
	} else {
		goto L88
	}
L73:
	;
	v280 = v247 << (uint(int32(4)) % 32)
	v281 = int32(1)
	v282 = v217
	goto L72
L74:
	;
	v255 = int32(0)
	v256 = F_errsave_start(m, v14)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L46
	} else {
		goto L82
	}
L75:
	;
	if v224 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L76:
	;
	if base.Ui32((v218-int32(65))&int32(255)) <= base.Ui32(int32(5)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v247 = v218 - int32(55)
	goto L75
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(5)) < base.Ui32((v218-int32(97))&int32(255)) {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v247 = v218 - int32(87)
	goto L75
L81:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v280 = v250 | v247
	v281 = int32(0)
	v282 = v217 + int32(1)
	goto L72
L82:
	;
	if v256 == int32(0) {
		v340 = v255
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L46
	} else {
		goto L84
	}
L84:
	;
	v263 = F_pg_mblen_cstr(m, v220)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L46
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v263
	F_errmsg(m, int32(95827), v9+int32(-48))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L46
	} else {
		goto L86
	}
L86:
	;
	F_errsave_finish(m, v14, int32(470441), int32(561), int32(265571))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L46
	} else {
		goto L87
	}
L87:
	;
	v340 = v255
	goto L1
L88:
	;
	goto L71
L89:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	switch v295 - int32(48) {
	case 0:
		goto L92
	case 1:
		goto L93
	default:
		goto L91
	}
L90:
	;
	if v295 == int32(0) {
		v340 = v182
		goto L1
	} else {
		goto L97
	}
L91:
	;
	goto L90
L92:
	;
	v303 = v288 & int32(255)
	v307 = base.B2i32(base.Ui32(v303) < base.Ui32(int32(2)))
	if base.Ui32(v303) < base.Ui32(int32(2)) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v299 = v298 | v288
	*(*uint8)(unsafe.Add(mBase, uint32(v287))) = uint8(v299)
	goto L92
L94:
	;
	v308 = int32(-128)
	goto L96
L95:
	;
	v308 = int32(base.Ui32(v303) >> (uint(int32(1)) % 32))
	goto L96
L96:
	;
	v287 = v287 + v307
	v288 = v308
	v290 = v290 + int32(1)
	goto L89
L97:
	;
	v314 = int32(0)
	v315 = F_errsave_start(m, v14)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L46
	} else {
		goto L98
	}
L98:
	;
	if v315 == int32(0) {
		v340 = v314
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v322 = F_pg_mblen_cstr(m, v290)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L46
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v322
	F_errmsg(m, int32(95766), v11)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L46
	} else {
		goto L102
	}
L102:
	;
	F_errsave_finish(m, v14, int32(470441), int32(536), int32(265571))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L46
	} else {
		goto L103
	}
L103:
	;
	v340 = v314
	goto L1
}
func F_varbit_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v32 = v2
		return v32
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v32 = v2
			return v32
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v32 = v2
				return v32
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v19 = F_exprTypmod(m, v18)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if int32(0) < v17 {
						if v19 <= int32(0) {
							v32 = v2
							return v32
						} else {
							if v17 < v19 {
								v32 = v2
								return v32
							} else {
								v28 = F_relabel_to_typmod(m, v18, v17)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									v32 = v28
									return v32
								}
							}
						}
					} else {
						v28 = F_relabel_to_typmod(m, v18, v17)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v32 = v28
							return v32
						}
					}
				}
			}
		}
	}
}
func F_varchar_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 < int32(4) {
		v61 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v71
L2:
	;
	v67 = F_cstring_to_text_with_len(m, l0, v61)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L20
	}
L3:
	;
	v15 = l2 - int32(4)
	if base.Ui32(l1) <= base.Ui32(v15) {
		v61 = l1
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = F_pg_mbcharcliplen(m, l0, l1, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v61 = v17
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	if base.Ui32(l1) <= base.Ui32(v17) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v24 = v17
	goto L9
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v24))))
	if v30 == int32(32) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v36 = int32(0)
	v37 = F_errsave_start(m, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	v34 = v24 + int32(1)
	if l1 != v34 {
		v24 = v34
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	if v37 == int32(0) {
		v71 = v36
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
	F_errmsg(m, int32(637450), v10)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, l3, int32(472712), int32(476), int32(61733))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v71 = v36
	goto L1
L20:
	;
	v71 = v67
	goto L1
}
func F_varcharin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3&int32(3) == int32(0) {
		v27 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = F_varchar_input(m, v3, v60, v61, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v60 = v52 - v3
	goto L1
L3:
	;
	v31 = v27
	goto L12
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v16 = v3
	goto L8
L8:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v52 = v20
	goto L2
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v46 = v31
	goto L15
L14:
	;
	goto L13
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v52 = v46
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	return v63
}
func F_varstr_sortsupport(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	if l2 != 0 {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		v9 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
			if v11 == int32(1) {
				if l1 != int32(19) {
					if l1 != int32(1042) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1554)
						v37 = int32(1)
						if v8&v37 == int32(0) {
							return
						} else {
							v42 = v37
							v45 = int32(0)
							v46 = v42
							v48 = F_palloc(m, int32(104))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_palloc(m, int32(1024))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v53 = int32(1024)
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v53
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = v51
									v57 = F_palloc(m, v53)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+96)) = v45
										*(*int64)(unsafe.Add(mBase, uint32(v48)+20)) = int64(4294967295)
										*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(-4294966272)
										*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v57
										*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = l1
										*(*uint8)(unsafe.Add(mBase, uint32(v48)+29)) = uint8(v11)
										v67 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v48)+28)) = uint8(v67)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
										if v46 == int32(0) {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v48)+88)) = int64(4596373779694328218)
											F_initHyperLogLog(m, v48+int32(40), int32(10))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_initHyperLogLog(m, v48-int32(-64), int32(10))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1555)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1556)
													v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v88
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(116)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1550)
						v20 = int32(1)
						if v8&v20 != 0 {
							v42 = v20
							v45 = int32(0)
							v46 = v42
							v48 = F_palloc(m, int32(104))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_palloc(m, int32(1024))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v53 = int32(1024)
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v53
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = v51
									v57 = F_palloc(m, v53)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+96)) = v45
										*(*int64)(unsafe.Add(mBase, uint32(v48)+20)) = int64(4294967295)
										*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(-4294966272)
										*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v57
										*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = l1
										*(*uint8)(unsafe.Add(mBase, uint32(v48)+29)) = uint8(v11)
										v67 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v48)+28)) = uint8(v67)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
										if v46 == int32(0) {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v48)+88)) = int64(4596373779694328218)
											F_initHyperLogLog(m, v48+int32(40), int32(10))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_initHyperLogLog(m, v48-int32(-64), int32(10))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1555)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1556)
													v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v88
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(116)
													return
												}
											}
										}
									}
								}
							}
						} else {
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1551)
					return
				}
			} else {
				v28 = base.B2i32(l1 != int32(19))
				if l1 != int32(19) {
					v29 = int32(1552)
				} else {
					v29 = int32(1553)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+12)))
				v45 = v9
				v46 = v28 & v32 & v8
				v48 = F_palloc(m, int32(104))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = F_palloc(m, int32(1024))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = int32(1024)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v48))) = v51
						v57 = F_palloc(m, v53)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+96)) = v45
							*(*int64)(unsafe.Add(mBase, uint32(v48)+20)) = int64(4294967295)
							*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(-4294966272)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = l1
							*(*uint8)(unsafe.Add(mBase, uint32(v48)+29)) = uint8(v11)
							v67 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v48)+28)) = uint8(v67)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
							if v46 == int32(0) {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v48)+88)) = int64(4596373779694328218)
								F_initHyperLogLog(m, v48+int32(40), int32(10))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									F_initHyperLogLog(m, v48-int32(-64), int32(10))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1555)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1556)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v88
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(116)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v99 = m.ExcPending
		if v99 != 0 {
			return
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return
			} else {
				F_errmsg(m, int32(233060), int32(0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					F_errhint(m, int32(534951), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						F_errfinish(m, int32(476976), int32(1648), int32(98766))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_view_has_instead_trigger(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	switch l1 - int32(2) {
	case 0:
		goto L4
	case 1:
		goto L3
	case 2:
		goto L5
	case 3:
		goto L8
	default:
		goto L7
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v109
L2:
	;
	v109 = int32(0)
	goto L1
L3:
	;
	if v11 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L4:
	;
	if v11 == int32(0) {
		goto L2
	} else {
		goto L38
	}
L5:
	;
	if v11 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L29
	} else {
		goto L33
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v109 = int32(1)
	goto L1
L10:
	;
	goto L11
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v18 <= int32(0) {
		v109 = int32(1)
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v21 = int32(0)
	if v21 < v18 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v24 = v18
	goto L15
L14:
	;
	v24 = v21
	goto L15
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v31 = int32(0)
	goto L16
L16:
	;
	v32 = int32(2)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25+v31<<(uint(v32)%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	switch v36 - v32 {
	case 0:
		goto L20
	case 1:
		goto L21
	case 2:
		goto L19
	default:
		goto L6
	case 5:
		goto L18
	}
L17:
	;
	v109 = v54
	goto L1
L18:
	;
	v54 = int32(1)
	v56 = v31 + v54
	if v56 != v24 {
		v31 = v56
		goto L16
	} else {
		goto L28
	}
L19:
	;
	v47 = int32(0)
	if v11 == v47 {
		v109 = v47
		goto L1
	} else {
		goto L26
	}
L20:
	;
	v43 = int32(0)
	if v11 == v43 {
		v109 = v43
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v39 = int32(0)
	if v11 == v39 {
		v109 = v39
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v42 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v109 = v39
	goto L1
L24:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v46 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v109 = v43
	goto L1
L26:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v50 == int32(0) {
		v109 = v47
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	goto L17
L29:
	;
	return int32(0)
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg_internal(m, int32(464019), v9)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(472506), int32(2568), int32(213048))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v77
	F_errmsg_internal(m, int32(463990), v9+int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(472506), int32(2562), int32(213048))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v91 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v109 = int32(1)
	goto L1
L38:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v97 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v109 = int32(1)
	goto L1
L40:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v103 == int32(0) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v109 = int32(1)
	goto L1
}
func F_visibilitymap_pin(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = base.I32_div_u_s(l1, int32(32672))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v7 != 0 {
		if v7 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11+(v7^int32(-1))<<(uint(int32(6))%32))+16))
			v26 = v17
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+v7<<(uint(int32(6))%32)+int32(-64))+16))
			v26 = v25
		}
		if v26 == v6 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			F_ReleaseBuffer(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v32 = F_vm_readbuf(m, l0, v6, int32(1))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
					return
				}
			}
		}
	} else {
		v32 = F_vm_readbuf(m, l0, v6, int32(1))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
			return
		}
	}
}
