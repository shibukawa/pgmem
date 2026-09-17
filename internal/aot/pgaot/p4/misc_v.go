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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v289 int64
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int64
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v411 int64
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v456 int32
	_ = v456
	var v458 int64
	_ = v458
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v504 int32
	_ = v504
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
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
	F_LWLockRelease(m, v344)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L11
	} else {
		goto L109
	}
L2:
	;
	m.G0 = v12 + int32(128)
	return v504
L3:
	;
	v504 = int32(1)
	goto L2
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v18
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)) = uint8(v20)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[0]))
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v109
	v113 = int32(0)
	if v14 < v113 {
		v131 = v113
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
	v48 = v12 + int32(108)
	v52 = l1 ^ int32(1)
	v55 = F_LockAcquireExtended(m, v48, int32(5), v46, v52, v46, v46)
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
		v504 = v46
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v61 = F_LockRelease(m, v48, int32(5), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v63 = int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)))
	if v64 != v63 {
		v504 = v63
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L18
L18:
	;
	v76 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)) = uint8(v76)
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v78
	v82 = F_TwoPhaseGetXidByVirtualXID(m, v12, v12+int32(96))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L11
	} else {
		goto L20
	}
L19:
	;
	goto L3
L20:
	;
	v85 = base.B2i32(v82 == int32(0))
	if v82 == int32(0) {
		v504 = v85
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v82
	v94 = v12 + int32(108)
	v96 = int32(0)
	v99 = F_LockAcquireExtended(m, v94, int32(5), v96, v52, v96, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	if v99 == int32(0) {
		v504 = v85
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v105 = F_LockRelease(m, v94, int32(5), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)))
	if v107 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	if v131 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[1]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	if base.Ui32(v120) <= base.Ui32(v14) {
		v131 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = v122 + v14*int32(640)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+44))
	if v127 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v128 = v125
	goto L32
L31:
	;
	v128 = int32(0)
	goto L32
L32:
	;
	v131 = v128
	goto L27
L33:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v134
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v136)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[0]))
	if v139 == v136 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v223 = v131 + int32(584)
	v225 = F_LWLockAcquire(m, v223, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L11
	} else {
		goto L51
	}
L36:
	;
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v142
	v149 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(24), v12+int32(127))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	if v149 == int32(0) {
		v504 = int32(1)
		goto L2
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v149
	v158 = int32(0)
	v160 = v12 + int32(108)
	v164 = l1 ^ int32(1)
	v167 = F_LockAcquireExtended(m, v160, int32(5), v158, v164, v158, v158)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	if v167 == int32(0) {
		v504 = v158
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v173 = F_LockRelease(m, v160, int32(5), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v175 = int32(1)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v176 != v175 {
		v504 = v175
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L43
L43:
	;
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v188)
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v190
	v196 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(16), v12+int32(127))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	goto L3
L45:
	;
	v199 = base.B2i32(v196 == int32(0))
	if v196 == int32(0) {
		v504 = v199
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v196
	v208 = v12 + int32(108)
	v210 = int32(0)
	v213 = F_LockAcquireExtended(m, v208, int32(5), v210, v164, v210, v210)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	if v213 == int32(0) {
		v504 = v199
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v219 = F_LockRelease(m, v208, int32(5), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v221 != 0 {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v131)+52))
	if v14 == v227 {
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
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v131)+612))
	if v229 == v109 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_LWLockRelease(m, v223)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L11
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v233 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v233
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v235)
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[0]))
	if v238 == v235 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v241
	v248 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(72), v12+int32(127))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	if v248 == int32(0) {
		v504 = int32(1)
		goto L2
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v248
	v257 = int32(0)
	v259 = v12 + int32(108)
	v263 = l1 ^ int32(1)
	v266 = F_LockAcquireExtended(m, v259, int32(5), v257, v263, v257, v257)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	if v266 == int32(0) {
		v504 = v257
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v272 = F_LockRelease(m, v259, int32(5), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v274 = int32(1)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v275 != v274 {
		v504 = v274
		goto L2
	} else {
		goto L64
	}
L64:
	;
	goto L65
L65:
	;
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v287)
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v289
	v295 = F_TwoPhaseGetXidByVirtualXID(m, v12-int32(-64), v12+int32(127))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L11
	} else {
		goto L67
	}
L66:
	;
	goto L3
L67:
	;
	v298 = base.B2i32(v295 == int32(0))
	if v295 == int32(0) {
		v504 = v298
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v295
	v307 = v12 + int32(108)
	v309 = int32(0)
	v312 = F_LockAcquireExtended(m, v307, int32(5), v309, v263, v309, v309)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	if v312 == int32(0) {
		v504 = v298
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v318 = F_LockRelease(m, v307, int32(5), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v320 != 0 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	goto L66
L73:
	;
	F_LWLockRelease(m, v223)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L11
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v326 = int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+608)))
	if v327 == v326 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v504 = int32(0)
	goto L2
L77:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[2]))
	v333 = v12 + int32(80)
	v334 = F_get_hash_value(m, v331, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L11
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v131)+36))
	F_LWLockRelease(m, v223)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L88
	}
L80:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[3]))
	v344 = v337 + v334&int32(15)<<(uint(int32(7))%32) + int32(_a_F_VirtualXactLock_0)
	v346 = F_LWLockAcquire(m, v344, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	v350 = F_SetupLockInTable(m, int32(_a_F_VirtualXactLock_1), v131, v333, v334, int32(7))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	if v350 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+128))
	v356 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v354)+128)) = v355 + v356
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v354)+116))
	v361 = v359 + v356
	*(*int32)(unsafe.Add(mBase, uint32(v354)+116)) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v354)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+16)) = v363 | int32(128)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v354)+72))
	if v367 == v361 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+20)) = v369 & int32(-129)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v350)+12)) = v373 | int32(128)
	F_LWLockRelease(m, v344)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	v379 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+608)) = uint8(v379)
	goto L79
L88:
	;
	v389 = v12 + int32(80)
	v391 = int32(0)
	v395 = F_LockAcquireExtended(m, v389, int32(5), v391, v391, v391, v391)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	v399 = F_LockRelease(m, v389, int32(5), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	v401 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v401
	v403 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v403)
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_VirtualXactLock[0]))
	if v406 == v403 {
		v504 = v326
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v385 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v411
	v417 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(56), v12+int32(127))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L95
	}
L93:
	;
	v421 = v385
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v421
	v427 = int32(0)
	v429 = v12 + int32(108)
	v435 = F_LockAcquireExtended(m, v429, int32(5), v427, v427, v427, v427)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L11
	} else {
		goto L97
	}
L95:
	;
	if v417 == int32(0) {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	v421 = v417
	goto L94
L97:
	;
	if v435 == int32(0) {
		v504 = v427
		goto L2
	} else {
		goto L98
	}
L98:
	;
	v441 = F_LockRelease(m, v429, int32(5), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	v443 = int32(1)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v444 != v443 {
		v504 = v443
		goto L2
	} else {
		goto L100
	}
L100:
	;
	goto L101
L101:
	;
	v456 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v456)
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v458
	v464 = F_TwoPhaseGetXidByVirtualXID(m, v12+int32(48), v12+int32(127))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L11
	} else {
		goto L103
	}
L102:
	;
	goto L3
L103:
	;
	v467 = base.B2i32(v464 == int32(0))
	if v464 == int32(0) {
		v504 = v467
		goto L2
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v464
	v476 = v12 + int32(108)
	v478 = int32(0)
	v482 = F_LockAcquireExtended(m, v476, int32(5), v478, v478, v478, v478)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	if v482 == int32(0) {
		v504 = v467
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v488 = F_LockRelease(m, v476, int32(5), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	if v490 != 0 {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	goto L102
L109:
	;
	F_LWLockRelease(m, v223)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(_a_F_VirtualXactLock_2))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_VirtualXactLock_3), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(_a_F_VirtualXactLock_4)
	F_errhint(m, int32(_a_F_VirtualXactLock_5), v12+int32(32))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_VirtualXactLock_6), int32(_a_F_VirtualXactLock_7), int32(_a_F_VirtualXactLock_8))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
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
func F_varbit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v17 <= int32(0) {
			v61 = v13
			m.G0 = v10 + int32(16)
			return v61
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v20 <= v17 {
				v61 = v13
				m.G0 = v10 + int32(16)
				return v61
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16777346))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
							F_errmsg(m, int32(_a_F_varbit_0), v10)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_varbit_1), int32(758), int32(_a_F_varbit_2))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
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
					v27 = int32(8)
					v28 = base.I32_div_s(v17+int32(7), v27)
					v30 = v28 + v27
					v31 = F_palloc(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v30 << (uint(int32(2)) % 32)
						v38 = v30 & int32(1073741823)
						v40 = v38 - int32(8)
						if v40 != 0 {
							v41 = int32(8)
							base.MemoryCopy(m, v31+v41, v13+v41, v40)
						} else {
						}
						v50 = v30<<(uint(int32(3))%32) - v17 + int32(-64)
						if v50 <= int32(0) {
							v61 = v31
						} else {
							v55 = v31 + v38 - int32(1)
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
							v59 = v56 & (int32(255) << (uint(v50) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
							v61 = v31
						}
						m.G0 = v10 + int32(16)
						return v61
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	return v229
L2:
	;
	if int32(0) < v14 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v29 = v15 + int32(1)
	v30 = F_strlen(m, v29)
	mBase = m.M
	if int32(536870911) <= v30 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v27 = F_strlen(m, v25)
	mBase = m.M
	v56 = v25
	v58 = int32(1)
	v59 = v27
	goto L2
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
	v33 = F_errsave_start(m, v13)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v56 = v29
	v58 = v2
	v59 = v30 << (uint(int32(2)) % 32)
	goto L2
L12:
	;
	return int32(0)
L13:
	;
	if v33 == int32(0) {
		v229 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(2147483640)
	F_errmsg(m, int32(_a_F_varbit_in_0), v9+int32(-16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, v13, int32(_a_F_varbit_in_1), int32(500), int32(_a_F_varbit_in_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v229 = v2
	goto L1
L18:
	;
	v175 = v56
	v176 = int32(128)
	v178 = v79
	goto L55
L19:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v101 == int32(0) {
		v229 = v70
		goto L1
	} else {
		goto L35
	}
L20:
	;
	v83 = F_errsave_start(m, v13)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L30
	}
L21:
	;
	if v14 < v59 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v63 = v59
	goto L23
L23:
	;
	v66 = int32(8)
	v67 = base.I32_div_s(v59+int32(7), v66)
	v69 = v67 + v66
	v70 = F_palloc0(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	v63 = v14
	goto L23
L25:
	;
	if v59 < v63 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v73 = v59
	goto L28
L27:
	;
	v73 = v63
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v69 << (uint(int32(2)) % 32)
	v79 = v70 + int32(8)
	if v58 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	if v83 == int32(0) {
		v229 = v2
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
	F_errmsg(m, int32(_a_F_varbit_in_3), v9+int32(-32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, v13, int32(_a_F_varbit_in_1), int32(514), int32(_a_F_varbit_in_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v229 = v2
	goto L1
L35:
	;
	v105 = v56
	v107 = v101
	v108 = v79
	v111 = int32(0)
	goto L36
L36:
	;
	v114 = v107 - int32(48)
	if base.Ui32(v114&int32(255)) < base.Ui32(int32(10)) {
		v135 = v114
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v229 = v70
	goto L1
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v168)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	if v172 != 0 {
		v105 = v105 + int32(1)
		v107 = v172
		v108 = v170
		v111 = v169
		goto L36
	} else {
		goto L54
	}
L39:
	;
	v168 = v135 << (uint(int32(4)) % 32)
	v169 = int32(1)
	v170 = v108
	goto L38
L40:
	;
	v143 = int32(0)
	v144 = F_errsave_start(m, v13)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L48
	}
L41:
	;
	if v111 == int32(0) {
		goto L39
	} else {
		goto L47
	}
L42:
	;
	if base.Ui32((v107-int32(65))&int32(255)) <= base.Ui32(int32(5)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v135 = v107 - int32(55)
	goto L41
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(5)) < base.Ui32((v107-int32(97))&int32(255)) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v135 = v107 - int32(87)
	goto L41
L47:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v168 = v138 | v135
	v169 = int32(0)
	v170 = v108 + int32(1)
	goto L38
L48:
	;
	if v144 == int32(0) {
		v229 = v143
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v151 = F_pg_mblen_cstr(m, v105)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v151
	F_errmsg(m, int32(_a_F_varbit_in_4), v9+int32(-48))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	F_errsave_finish(m, v13, int32(_a_F_varbit_in_1), int32(561), int32(_a_F_varbit_in_2))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v229 = v143
	goto L1
L54:
	;
	goto L37
L55:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	switch v183 - int32(48) {
	case 0:
		goto L58
	case 1:
		goto L59
	default:
		goto L57
	}
L56:
	;
	if v183 == int32(0) {
		v229 = v70
		goto L1
	} else {
		goto L63
	}
L57:
	;
	goto L56
L58:
	;
	v189 = int32(1)
	v194 = int32(base.Ui32(v176)>>(uint(v189)%32)) & int32(127)
	if v194 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v187 = v186 | v176
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v187)
	goto L58
L60:
	;
	v196 = v194
	goto L62
L61:
	;
	v196 = int32(-128)
	goto L62
L62:
	;
	v175 = v175 + v189
	v176 = v196
	v178 = v178 + base.B2i32(v194 == int32(0))
	goto L55
L63:
	;
	v202 = int32(0)
	v203 = F_errsave_start(m, v13)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	if v203 == int32(0) {
		v229 = v202
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	v210 = F_pg_mblen_cstr(m, v175)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v210
	F_errmsg(m, int32(_a_F_varbit_in_5), v11)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_errsave_finish(m, v13, int32(_a_F_varbit_in_1), int32(536), int32(_a_F_varbit_in_2))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	v229 = v202
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v35 = v2
		return v35
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v35 = v2
			return v35
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v35 = v2
				return v35
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v21 = F_exprTypmod(m, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if base.B2i32(int32(0) < v17)&(base.B2i32(v17 < v21)|base.B2i32(v21 <= int32(0))) != 0 {
						v35 = v2
						return v35
					} else {
						v30 = F_relabel_to_typmod(m, v20, v17)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v35 = v30
							return v35
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
	F_errmsg(m, int32(_a_F_varchar_input_0), v10)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, l3, int32(_a_F_varchar_input_1), int32(476), int32(_a_F_varchar_input_2))
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_strlen(m, v3)
	mBase = m.M
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_varchar_input(m, v3, v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
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
	var v22 int32
	_ = v22
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
				if l1 == int32(1042) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1539)
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
												*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1540)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1541)
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
					if l1 == int32(19) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1535)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1536)
						v22 = int32(1)
						if v8&v22 != 0 {
							v42 = v22
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
													*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1540)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1541)
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
				}
			} else {
				v28 = base.B2i32(l1 != int32(19))
				if l1 != int32(19) {
					v29 = int32(1537)
				} else {
					v29 = int32(1538)
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
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1540)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1541)
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
				F_errmsg(m, int32(_a_F_varstr_sortsupport_0), int32(0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_F_varstr_sortsupport_1), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_varstr_sortsupport_2), int32(1648), int32(_a_F_varstr_sortsupport_3))
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
func F_verifyNotNullPKCompatible(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
	v11 = v9 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)))
	if v12 == int32(110) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+106)))
		if v15 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
					F_errmsg(m, int32(_a_F_verifyNotNullPKCompatible_0), v7+int32(32))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
						v54 = F_get_rel_name(m, v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(_a_F_verifyNotNullPKCompatible_1)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v11 + int32(4)
							F_errdetail(m, int32(_a_F_verifyNotNullPKCompatible_2), v7+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_verifyNotNullPKCompatible_3)
								F_errhint(m, int32(_a_F_verifyNotNullPKCompatible_4), v7)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_verifyNotNullPKCompatible_5), int32(_a_F_verifyNotNullPKCompatible_6), int32(_a_F_verifyNotNullPKCompatible_7))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = l1
						F_errmsg(m, int32(_a_F_verifyNotNullPKCompatible_0), v7+int32(80))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
							v92 = F_get_rel_name(m, v91)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = int32(_a_F_verifyNotNullPKCompatible_8)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v11 + int32(4)
								F_errdetail(m, int32(_a_F_verifyNotNullPKCompatible_2), v7-int32(-64))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_verifyNotNullPKCompatible_9)
									F_errhint(m, int32(_a_F_verifyNotNullPKCompatible_10), v7+int32(48))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_verifyNotNullPKCompatible_5), int32(_a_F_verifyNotNullPKCompatible_11), int32(_a_F_verifyNotNullPKCompatible_7))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
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
			} else {
				m.G0 = v7 + int32(112)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = v28
			F_errmsg_internal(m, int32(_a_F_verifyNotNullPKCompatible_12), v7+int32(96))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_verifyNotNullPKCompatible_5), int32(_a_F_verifyNotNullPKCompatible_13), int32(_a_F_verifyNotNullPKCompatible_7))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
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
func F_view_has_instead_trigger(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	switch l1 - int32(2) {
	case 0:
		goto L7
	case 1:
		goto L8
	case 2:
		goto L6
	case 3:
		goto L9
	default:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L34
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v91
L3:
	;
	v91 = int32(0)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L34
	} else {
		goto L35
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v34 <= int32(0) {
		v91 = int32(1)
		goto L2
	} else {
		goto L17
	}
L6:
	;
	if v11 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L7:
	;
	if v11 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L8:
	;
	if v11 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v91 = int32(1)
	goto L2
L11:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v17 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v91 = int32(1)
	goto L2
L13:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v23 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v91 = int32(1)
	goto L2
L15:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v29 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v91 = int32(1)
	goto L2
L17:
	;
	v37 = int32(0)
	if v37 < v34 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v40 = v34
	goto L20
L19:
	;
	v40 = v37
	goto L20
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v47 = int32(0)
	goto L21
L21:
	;
	v48 = int32(2)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41+v47<<(uint(v48)%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	switch v52 - v48 {
	case 0:
		goto L25
	case 1:
		goto L26
	case 2:
		goto L24
	default:
		goto L1
	case 5:
		goto L23
	}
L22:
	;
	v91 = v70
	goto L2
L23:
	;
	v70 = int32(1)
	v72 = v47 + v70
	if v72 != v40 {
		v47 = v72
		goto L21
	} else {
		goto L33
	}
L24:
	;
	v63 = int32(0)
	if v11 == v63 {
		v91 = v63
		goto L2
	} else {
		goto L31
	}
L25:
	;
	v59 = int32(0)
	if v11 == v59 {
		v91 = v59
		goto L2
	} else {
		goto L29
	}
L26:
	;
	v55 = int32(0)
	if v11 == v55 {
		v91 = v55
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v58 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v91 = v55
	goto L2
L29:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v62 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v91 = v59
	goto L2
L31:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v66 == int32(0) {
		v91 = v63
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	goto L22
L34:
	;
	return int32(0)
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg_internal(m, int32(_a_F_view_has_instead_trigger_0), v9)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_view_has_instead_trigger_1), int32(2568), int32(_a_F_view_has_instead_trigger_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
	F_errmsg_internal(m, int32(_a_F_view_has_instead_trigger_3), v9+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_view_has_instead_trigger_1), int32(2562), int32(_a_F_view_has_instead_trigger_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	v6 = base.I32_div_u_s(l1, int32(_a_F_visibilitymap_pin_0))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v7 != 0 {
		if v7 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_pin[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11+(v7^int32(-1))<<(uint(int32(6))%32))+16))
			v26 = v17
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_pin[1]))
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
